package smspva

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type fixture struct {
	Key      string          `json:"key"`
	Params   json.RawMessage `json:"params"`
	Response json.RawMessage `json:"response"`
}

func fixtures(t *testing.T) []fixture {
	t.Helper()
	raw, err := os.ReadFile("testdata/cases.json")
	if err != nil {
		t.Fatal(err)
	}
	var cases []fixture
	if err = json.Unmarshal(raw, &cases); err != nil {
		t.Fatal(err)
	}
	return cases
}
func makeClient(t *testing.T, options Options) *Client {
	t.Helper()
	client, err := New(options)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(client.CloseIdleConnections)
	return client
}
func jsonEqual(a, b []byte) bool {
	x, e := decodeObject(a)
	if e != nil {
		return false
	}
	y, e := decodeObject(b)
	return e == nil && reflect.DeepEqual(x, y)
}

const secret = "go-local-key-+&?=/"

func TestAll56NamedMethodsWire(t *testing.T) {
	cases := fixtures(t)
	if len(cases) != 56 || len(namedCalls) != 56 {
		t.Fatal("Expected all56namedmethods")
	}
	var current fixture
	var calls int
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		op := operations[current.Key]
		if r.Method != op.Method {
			t.Errorf("verb %s: %s", current.Key, r.Method)
		}
		params, _ := decodeObject(current.Params)
		path := op.Path
		expected := url.Values{}
		if op.Action != "" {
			expected.Set("method", op.Action)
		}
		for _, p := range op.Params {
			value, ok := params[p.Name]
			if !ok || value == nil || (p.OmitEmpty && value == "") {
				continue
			}
			text := fmt.Sprint(value)
			if p.In == "path" {
				path = strings.Replace(path, "{"+p.Name+"}", url.PathEscape(text), 1)
			} else if p.In == "query" {
				expected.Set(p.Name, text)
			}
		}
		if op.Auth != "none" {
			if op.Family == "activation" {
				if r.Header.Get("apikey") != secret {
					t.Error("Missing headerkey")
				}
			} else {
				expected.Set("apikey", secret)
			}
		} else if r.Header.Get("apikey") != "" {
			t.Error("Publicheaderkeyleak")
		}
		if r.URL.EscapedPath() != path {
			t.Errorf("path %s: %s != %s", current.Key, r.URL.EscapedPath(), path)
		}
		if !reflect.DeepEqual(r.URL.Query(), expected) {
			t.Errorf("query %s: got %v want %v", current.Key, r.URL.Query(), expected)
		}
		if r.Header.Get("Cookie") != "" {
			t.Error("Cookies must notbe sent")
		}
		http.SetCookie(w, &http.Cookie{Name: "session", Value: "not-for-next-call"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(current.Response)
	}))
	defer upstream.Close()
	client := makeClient(t, Options{APIKey: secret, ActivationBaseURL: upstream.URL, RentalBaseURL: upstream.URL, EmailBaseURL: upstream.URL})
	for _, item := range cases {
		current = item
		t.Run(item.Key, func(t *testing.T) {
			response, err := namedCalls[item.Key](context.Background(), client, item.Params)
			if err != nil {
				t.Fatal(err)
			}
			if !jsonEqual(response.Body, item.Response) {
				t.Fatalf("Responsechanged: %s", response.Body)
			}
		})
	}
	if calls != 56 {
		t.Fatalf("Expected56calls,got%d", calls)
	}
}

func TestPublicOptionalAuthAndEncoding(t *testing.T) {
	var received []*http.Request
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = append(received, r)
		if strings.HasPrefix(r.URL.Path, "/activation") {
			w.Write([]byte(`{"statusCode":200,"data":{}}`))
		} else {
			w.Write([]byte(`{"status":1,"data":[]}`))
		}
	}))
	defer upstream.Close()
	client := makeClient(t, Options{APIKey: secret, ActivationBaseURL: upstream.URL, RentalBaseURL: upstream.URL, EmailBaseURL: upstream.URL})
	for _, pcode := range []*string{nil, Ptr("")} {
		_, err := client.Email.GetData(context.Background(), EmailGetDataParams{Pcode: pcode})
		if err != nil {
			t.Fatal(err)
		}
		r := received[len(received)-1]
		if r.URL.Query().Has("pcode") || r.URL.Query().Has("apikey") {
			t.Fatal("Omittedpublicparamsleaked")
		}
	}
	value := "A +&?=/中文"
	_, err := client.Activation.GetOperators(context.Background(), ActivationGetOperatorsParams{Country: &value})
	if err != nil {
		t.Fatal(err)
	}
	if received[len(received)-1].URL.Path != "/activation/operators/"+value {
		t.Fatal("Pathencodingfailed")
	}
	var optional string
	for key, op := range operations {
		if op.Auth == "optional" {
			optional = key
			break
		}
	}
	if optional == "" {
		t.Fatal("Expectedoptionalauth")
	}
	var item fixture
	for _, f := range fixtures(t) {
		if f.Key == optional {
			item = f
			break
		}
	}
	params, _ := decodeObject(item.Params)
	if _, err := client.call(context.Background(), optional, params, WithoutAPIKey()); err != nil {
		t.Fatal(err)
	}
	if received[len(received)-1].URL.Query().Has("apikey") {
		t.Fatal("Optionalkeynotomitted")
	}
	public := makeClient(t, Options{EmailBaseURL: upstream.URL})
	if _, err := public.Email.GetProviders(context.Background()); err != nil {
		t.Fatal(err)
	}
	n := len(received)
	if _, err := public.Activation.GetBalance(context.Background()); err == nil {
		t.Fatal("Missingkeyaccepted")
	}
	if len(received) != n {
		t.Fatal("ValidationperformedHTTP")
	}
}

func TestPendingNoStockAndErrors(t *testing.T) {
	status := 202
	body := `{"statusCode":202,"error":{"type":"SMS_NOT_YET_RECEIVED"}}`
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(status); w.Write([]byte(body)) }))
	defer upstream.Close()
	client := makeClient(t, Options{APIKey: secret, ActivationBaseURL: upstream.URL, EmailBaseURL: upstream.URL})
	response, err := client.Activation.ReceiveSms(context.Background(), ActivationReceiveSmsParams{OrderID: Ptr(int64(123))})
	if err != nil || !IsPending(response) {
		t.Fatalf("Pendingnotreturned: %v", err)
	}
	status = 200
	body = `{"status":2}`
	var create fixture
	for _, f := range fixtures(t) {
		if f.Key == "email.createOrder" {
			create = f
		}
	}
	response, err = namedCalls[create.Key](context.Background(), client, create.Params)
	if err != nil || !IsNoStock(response) {
		t.Fatalf("Nostocknotreturned: %v", err)
	}
	for _, code := range []int{200, 402, 410, 503} {
		status = code
		body = `{"statusCode":0,"error":{"type":"ORIGINAL_CODE"}}`
		_, err = client.Activation.GetBalance(context.Background())
		var api *APIError
		if !errors.As(err, &api) || api.HTTPStatus != code || api.Code != "ORIGINAL_CODE" || string(api.Body) != body {
			t.Fatalf("LostAPIerror: %#v", err)
		}
	}
	for _, invalid := range []string{`<html>error</html>`, `{}`, `[]`, `{"statusCode":"200"}`, `{"statusCode":true}`, `{"statusCode":200} extra`, `{"statusCode":200.5}`} {
		status = 200
		body = invalid
		_, err = client.Activation.GetBalance(context.Background())
		var protocol *ProtocolError
		if !errors.As(err, &protocol) {
			t.Fatalf("Expectedprotocolfor%s,got%v", body, err)
		}
	}
	if IsPending(Response{202, []byte(`{"statusCode":"202","error":{"type":"SMS_NOT_YET_RECEIVED"}}`)}) || IsNoStock(Response{200, []byte(`{"status":"2"}`)}) {
		t.Fatal("Helpersacceptedstringcodes")
	}
}

func TestInputValidationBeforeRequest(t *testing.T) {
	var calls int
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls++ }))
	defer upstream.Close()
	client := makeClient(t, Options{APIKey: secret, ActivationBaseURL: upstream.URL, RentalBaseURL: upstream.URL, EmailBaseURL: upstream.URL})
	invalid := []struct {
		key   string
		value any
	}{
		{"activation.getOperators", map[string]any{}}, {"activation.receiveSms", map[string]any{"orderid": "123"}},
		{"email.getProviders", map[string]any{"apikey": "evil"}}, {"activation.getOperators", map[string]any{"country": ".."}},
		{"activation.getBalance", map[string]any{"unknown": true}}, {"rental.getMultiServiceCount", map[string]any{}},
	}
	for _, item := range invalid {
		_, err := client.call(context.Background(), item.key, item.value)
		var validation *ValidationError
		if !errors.As(err, &validation) {
			t.Errorf("Expectedvalidation:%s", item.key)
		}
	}
	if calls != 0 {
		t.Fatal("ValidationcalledAPI")
	}
	if _, err := New(Options{APIKey: "x\r\nsecret"}); err == nil {
		t.Fatal("Badkeyaccepted")
	}
	if _, err := New(Options{EmailBaseURL: "https://user:secret@example.com"}); err == nil {
		t.Fatal("CredentialURLaccepted")
	}
	if strings.Contains(fmt.Sprintf("%+v %#v", client, client), secret) {
		t.Fatal("Clientprintingexposedkey")
	}
}

func TestNoRedirectCookiesOrAutomaticReplay(t *testing.T) {
	var count atomic.Int32
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := count.Add(1)
		if n == 1 {
			http.SetCookie(w, &http.Cookie{Name: "session", Value: "secret"})
			w.Write([]byte(`{"statusCode":200,"data":{}}`))
			return
		}
		if r.Header.Get("Cookie") != "" {
			t.Error("Cookiesinherited")
		}
		if n == 2 {
			w.Header().Set("Location", "/should-not-follow")
			w.WriteHeader(302)
			return
		}
		conn, _, err := w.(http.Hijacker).Hijack()
		if err != nil {
			t.Error(err)
			return
		}
		conn.Close()
	}))
	defer upstream.Close()
	client := makeClient(t, Options{APIKey: secret, ActivationBaseURL: upstream.URL})
	if _, err := client.Activation.GetBalance(context.Background()); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Activation.GetBalance(context.Background()); err == nil {
		t.Fatal("Redirectaccepted")
	}
	if count.Load() != 2 {
		t.Fatal("Redirectfollowed")
	}
	_, err := client.Activation.ReceiveSms(context.Background(), ActivationReceiveSmsParams{OrderID: Ptr(int64(123))})
	var transport *TransportError
	if !errors.As(err, &transport) || !transport.OutcomeUnknown {
		t.Fatal("Lostuncertainoutcome")
	}
	if count.Load() != 3 {
		t.Fatal("Requestreplayed")
	}
	if strings.Contains(err.Error(), secret) || strings.Contains(err.Error(), upstream.URL) {
		t.Fatal("Unsafeerror")
	}
}

func TestTimeoutCancellationAndSavedPollingDeadline(t *testing.T) {
	var calls atomic.Int32
	delay := atomic.Bool{}
	delay.Store(true)
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		if delay.Load() {
			<-r.Context().Done()
			return
		}
		w.WriteHeader(202)
		w.Write([]byte(`{"statusCode":202,"error":{"type":"SMS_NOT_YET_RECEIVED"}}`))
	}))
	defer upstream.Close()
	client := makeClient(t, Options{APIKey: secret, ActivationBaseURL: upstream.URL})
	_, err := client.Activation.ReceiveSms(context.Background(), ActivationReceiveSmsParams{OrderID: Ptr(int64(123))}, WithTimeout(20*time.Millisecond))
	var transport *TransportError
	if !errors.As(err, &transport) || transport.Kind != "timeout" || !transport.OutcomeUnknown {
		t.Fatalf("Losttimeout:%v", err)
	}
	if calls.Load() != 1 {
		t.Fatal("Timeoutretried")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = client.Activation.GetBalance(ctx)
	if !errors.Is(err, context.Canceled) || calls.Load() != 1 {
		t.Fatal("Precancelledrequestsent")
	}
	delay.Store(false)
	before := calls.Load()
	_, err = client.Activation.WaitForSMS(context.Background(), WaitForSMSOptions{OrderID: 123, ExpiresAt: time.Now().Add(40 * time.Millisecond)})
	var expired *PollingTimeoutError
	if !errors.As(err, &expired) || calls.Load()-before != 1 {
		t.Fatalf("Deadlinecausedextraattempt:%v count%d", err, calls.Load()-before)
	}
	before = calls.Load()
	_, err = client.Activation.WaitForSMS(context.Background(), WaitForSMSOptions{OrderID: 123, ExpiresAt: time.Now().Add(-time.Second)})
	if !errors.As(err, &expired) || calls.Load() != before {
		t.Fatal("Expiredordershouldnotpoll")
	}
	_, err = client.Activation.WaitForSMS(context.Background(), WaitForSMSOptions{OrderID: 123, ExpiresAt: time.Now().Add(time.Minute), Interval: time.Millisecond})
	var validation *ValidationError
	if !errors.As(err, &validation) {
		t.Fatal("Invalidpollintervalaccepted")
	}
	deadlineContext, stop := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer stop()
	before = calls.Load()
	_, err = client.Activation.GetBalance(deadlineContext)
	if !errors.As(err, &transport) || transport.Kind != "timeout" || transport.OutcomeUnknown || !errors.Is(err, context.DeadlineExceeded) || calls.Load() != before {
		t.Fatal("Already-expired context lost timeout classification")
	}
	waitContext, stopWait := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer stopWait()
	_, err = client.Activation.WaitForSMS(waitContext, WaitForSMSOptions{OrderID: 123, ExpiresAt: time.Now().Add(time.Minute)})
	if !errors.As(err, &transport) || transport.Kind != "timeout" || transport.OutcomeUnknown || !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal("Caller deadline during polling sleep lost timeout classification")
	}
}

func TestResponseRawNumberAndStringTypes(t *testing.T) {
	body := []byte(`{"statusCode":200,"data":{"id":9007199254740993,"price":"0.10"}}`)
	response := Response{200, body}
	encoded, err := json.Marshal(response)
	if err != nil || !bytes.Equal(encoded, body) {
		t.Fatal("RawJSONwaschanged")
	}
	var parsed struct {
		Data struct {
			ID    int64  `json:"id"`
			Price string `json:"price"`
		} `json:"data"`
	}
	if err = response.Decode(&parsed); err != nil || parsed.Data.ID != 9007199254740993 || parsed.Data.Price != "0.10" {
		t.Fatal("Typesnotpreserved")
	}
}
