package smspva

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

//go:embed contracts.json
var contractBytes []byte

type schema struct {
	Type      string            `json:"type"`
	Enum      []json.RawMessage `json:"enum"`
	Minimum   *float64          `json:"minimum"`
	Maximum   *float64          `json:"maximum"`
	MinLength *int              `json:"minLength"`
	MaxLength *int              `json:"maxLength"`
	Pattern   string            `json:"pattern"`
}
type parameter struct {
	Name      string `json:"name"`
	In        string `json:"in"`
	Required  bool   `json:"required"`
	OmitEmpty bool   `json:"omitEmpty"`
	Schema    schema `json:"schema"`
}
type operation struct {
	Family   string      `json:"family"`
	Path     string      `json:"path"`
	Method   string      `json:"method"`
	BaseURL  string      `json:"baseUrl"`
	Auth     string      `json:"auth"`
	Action   string      `json:"action"`
	Params   []parameter `json:"params"`
	Pending  bool        `json:"pending"`
	NoStock  bool        `json:"noStock"`
	Stateful bool        `json:"stateful"`
}

var operations = func() map[string]operation {
	var doc struct {
		Operations map[string]operation `json:"operations"`
	}
	if err := json.Unmarshal(contractBytes, &doc); err != nil {
		panic("Invalid bundled SMSPVA contract")
	}
	return doc.Operations
}()

// Options configures one account. Transport is optional; custom transports must not retry paid GETs.
type Options struct {
	APIKey                                         string
	Timeout                                        time.Duration
	ActivationBaseURL, RentalBaseURL, EmailBaseURL string
	Transport                                      http.RoundTripper
}
type Client struct {
	Activation *Activation
	Rental     *Rental
	Email      *Email
	apiKey     string
	timeout    time.Duration
	bases      map[string]string
	http       *http.Client
}
type requestOptions struct {
	timeout time.Duration
	useKey  bool
}
type RequestOption func(*requestOptions) error

func WithTimeout(timeout time.Duration) RequestOption {
	return func(o *requestOptions) error {
		if timeout <= 0 {
			return &ValidationError{"timeout must be positive"}
		}
		o.timeout = timeout
		return nil
	}
}
func WithoutAPIKey() RequestOption {
	return func(o *requestOptions) error { o.useKey = false; return nil }
}
func Ptr[T any](value T) *T { return &value }

func (c *Client) String() string   { return "SMSPVA client" }
func (c *Client) GoString() string { return "SMSPVA client" }

// Response retains the original JSON bytes. Decode into a caller-owned typed value when needed.
type Response struct {
	HTTPStatus int
	Body       json.RawMessage
}

func (r Response) Decode(target any) error { return json.Unmarshal(r.Body, target) }
func (r Response) MarshalJSON() ([]byte, error) {
	if !json.Valid(r.Body) {
		return nil, errors.New("Invalid response JSON")
	}
	return r.Body, nil
}
func IsPending(r Response) bool {
	if r.HTTPStatus != 202 {
		return false
	}
	body, err := decodeObject(r.Body)
	if err != nil {
		return false
	}
	status, ok := body["statusCode"].(json.Number)
	if !ok || status != "202" {
		return false
	}
	e, ok := body["error"].(map[string]any)
	return ok && e["type"] == "SMS_NOT_YET_RECEIVED"
}
func IsNoStock(r Response) bool {
	if r.HTTPStatus != 200 {
		return false
	}
	body, err := decodeObject(r.Body)
	if err != nil {
		return false
	}
	status, ok := body["status"].(json.Number)
	return ok && status == "2"
}

func New(options Options) (*Client, error) {
	if strings.ContainsAny(options.APIKey, "\r\n") {
		return nil, &ValidationError{"API key must not contain line breaks"}
	}
	timeout := options.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	if timeout < 0 {
		return nil, &ValidationError{"timeout must be positive"}
	}
	bases := map[string]string{}
	for family, base := range map[string]string{"activation": options.ActivationBaseURL, "rental": options.RentalBaseURL, "email": options.EmailBaseURL} {
		if base == "" {
			continue
		}
		u, err := url.Parse(base)
		if err != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
			return nil, &ValidationError{"Base URL must use HTTP(S) without credentials, query or fragment"}
		}
		bases[family] = strings.TrimRight(base, "/")
	}
	transport := options.Transport
	if transport == nil {
		transport = &http.Transport{Proxy: nil, DisableKeepAlives: true, ForceAttemptHTTP2: false, TLSHandshakeTimeout: 10 * time.Second, DialContext: (&net.Dialer{Timeout: 30 * time.Second}).DialContext}
	}
	c := &Client{apiKey: options.APIKey, timeout: timeout, bases: bases, http: &http.Client{Transport: transport, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}}
	c.Activation = &Activation{c}
	c.Rental = &Rental{c}
	c.Email = &Email{c}
	return c, nil
}
func (c *Client) CloseIdleConnections() { c.http.CloseIdleConnections() }

func decodeObject(raw []byte) (map[string]any, error) {
	var body map[string]any
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(&body); err != nil || body == nil {
		return nil, errors.New("invalid JSON object")
	}
	var trailing any
	if decoder.Decode(&trailing) != io.EOF {
		return nil, errors.New("trailing JSON content")
	}
	return body, nil
}
func validate(p parameter, value any) error {
	fail := func() error { return &ValidationError{"Invalid value for parameter: " + p.Name} }
	s := p.Schema
	switch s.Type {
	case "string":
		text, ok := value.(string)
		if !ok {
			return fail()
		}
		n := utf8.RuneCountInString(text)
		if s.MinLength != nil && n < *s.MinLength || s.MaxLength != nil && n > *s.MaxLength {
			return fail()
		}
		if s.Pattern != "" {
			matched, err := regexp.MatchString(s.Pattern, text)
			if err != nil || !matched {
				return fail()
			}
		}
	case "integer", "number":
		n, ok := value.(json.Number)
		if !ok {
			return fail()
		}
		v, err := n.Float64()
		if err != nil || math.IsInf(v, 0) || math.IsNaN(v) {
			return fail()
		}
		if s.Type == "integer" {
			if _, err = n.Int64(); err != nil {
				return fail()
			}
		}
		if s.Minimum != nil && v < *s.Minimum || s.Maximum != nil && v > *s.Maximum {
			return fail()
		}
	case "boolean":
		if _, ok := value.(bool); !ok {
			return fail()
		}
	default:
		return fail()
	}
	if len(s.Enum) > 0 {
		encoded, _ := json.Marshal(value)
		found := false
		for _, item := range s.Enum {
			if bytes.Equal(encoded, item) {
				found = true
				break
			}
		}
		if !found {
			return fail()
		}
	}
	return nil
}

func (c *Client) call(ctx context.Context, key string, input any, options ...RequestOption) (Response, error) {
	op, ok := operations[key]
	if !ok {
		return Response{}, &ValidationError{"Unknown SDK operation"}
	}
	opts := requestOptions{timeout: c.timeout, useKey: true}
	for _, option := range options {
		if option == nil {
			return Response{}, &ValidationError{"Request option cannot be nil"}
		}
		if err := option(&opts); err != nil {
			return Response{}, err
		}
	}
	params := map[string]any{}
	if input != nil {
		raw, err := json.Marshal(input)
		if err != nil {
			return Response{}, &ValidationError{"Invalid parameters"}
		}
		params, err = decodeObject(raw)
		if err != nil {
			return Response{}, &ValidationError{"Parameters must be an object"}
		}
	}
	allowed := map[string]bool{}
	for _, p := range op.Params {
		allowed[p.Name] = true
	}
	for name := range params {
		if !allowed[name] {
			return Response{}, &ValidationError{"Unexpected request parameter"}
		}
	}
	query := url.Values{}
	headers := http.Header{"Accept": []string{"application/json"}}
	path := op.Path
	if op.Action != "" {
		query.Set("method", op.Action)
	}
	for _, p := range op.Params {
		value, present := params[p.Name]
		if !present || value == nil || (p.OmitEmpty && value == "") {
			if p.Required {
				return Response{}, &ValidationError{"Missing required parameter: " + p.Name}
			}
			continue
		}
		if err := validate(p, value); err != nil {
			return Response{}, err
		}
		text := fmt.Sprint(value)
		switch p.In {
		case "path":
			if text == "." || text == ".." {
				return Response{}, &ValidationError{"Invalid path parameter: " + p.Name}
			}
			path = strings.Replace(path, "{"+p.Name+"}", url.PathEscape(text), 1)
		case "header":
			if strings.ContainsAny(text, "\r\n") {
				return Response{}, &ValidationError{"Invalid header parameter: " + p.Name}
			}
			headers.Set(p.Name, text)
		default:
			query.Set(p.Name, text)
		}
	}
	if op.Auth == "required" && c.apiKey == "" {
		return Response{}, &ValidationError{"Set APIKey for this operation"}
	}
	if c.apiKey != "" && op.Auth != "none" && (op.Auth == "required" || opts.useKey) {
		if op.Family == "activation" {
			headers.Set("apikey", c.apiKey)
		} else {
			query.Set("apikey", c.apiKey)
		}
	}
	base := c.bases[op.Family]
	if base == "" {
		base = op.BaseURL
	}
	address := strings.TrimRight(base, "/") + path
	if len(query) > 0 {
		address += "?" + query.Encode()
	}
	if ctx == nil {
		return Response{}, &ValidationError{"Context cannot be nil"}
	}
	if err := ctx.Err(); err != nil {
		kind := "cancelled"
		if err == context.DeadlineExceeded {
			kind = "timeout"
		}
		return Response{}, &TransportError{Family: op.Family, Operation: key, Kind: kind, cause: err}
	}
	ctx, cancel := context.WithTimeout(ctx, opts.timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, op.Method, address, nil)
	if err != nil {
		return Response{}, &ValidationError{"Could not construct API request"}
	}
	req.Header = headers
	req.Close = true
	response, err := c.http.Do(req)
	if err != nil {
		return Response{}, transportFailure(ctx, err, op, key)
	}
	defer response.Body.Close()
	raw, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{}, transportFailure(ctx, err, op, key)
	}
	body, err := decodeObject(raw)
	if err != nil {
		return Response{}, &ProtocolError{op.Family, key, response.StatusCode, op.Stateful}
	}
	var status any
	if op.Family == "activation" {
		status = body["statusCode"]
	} else {
		status = body["status"]
	}
	apiCode := any(response.StatusCode)
	if n, ok := status.(json.Number); ok {
		apiCode = n
	}
	if id, ok := body["errorid"].(json.Number); ok {
		apiCode = id
	}
	if e, ok := body["error"].(map[string]any); ok {
		if code, ok := e["type"].(string); ok {
			apiCode = code
		}
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return Response{}, &APIError{op.Family, key, response.StatusCode, apiCode, raw}
	}
	numeric, ok := status.(json.Number)
	if !ok {
		return Response{}, &ProtocolError{op.Family, key, response.StatusCode, op.Stateful}
	}
	if _, err := numeric.Int64(); err != nil {
		return Response{}, &ProtocolError{op.Family, key, response.StatusCode, op.Stateful}
	}
	result := Response{response.StatusCode, raw}
	if op.Pending && IsPending(result) || op.NoStock && IsNoStock(result) {
		return result, nil
	}
	success := numeric == "1"
	if op.Family == "activation" {
		_, hasError := body["error"]
		success = numeric == "200" && (!hasError || body["error"] == nil)
	}
	if !success {
		return Response{}, &APIError{op.Family, key, response.StatusCode, apiCode, raw}
	}
	return result, nil
}
func transportFailure(ctx context.Context, err error, op operation, key string) error {
	kind := "network"
	var cause error
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		kind = "timeout"
		cause = context.DeadlineExceeded
	} else if errors.Is(ctx.Err(), context.Canceled) {
		kind = "cancelled"
		cause = context.Canceled
	} else {
		var ne net.Error
		if errors.As(err, &ne) && ne.Timeout() {
			kind = "timeout"
			cause = context.DeadlineExceeded
		}
	}
	return &TransportError{op.Family, key, kind, op.Stateful, cause}
}
