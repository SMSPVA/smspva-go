# SMSPVA Go SDK

Official Go client for [SMSPVA](https://smspva.com/), a service for receiving verification SMS, renting phone numbers and renting private email addresses. Covers all 56 methods in Activation V2 (20), Rental (22) and Private Email (14). Requires Go 1.22+ and uses only the standard library. This is an alpha release. Alternative API is excluded.

## Installation

```sh
go get github.com/SMSPVA/smspva-go@v0.1.0-alpha.1
```

[Sign in to your SMSPVA account](https://smspva.com/signin.html), open **Profile and ApiKey**, and set **`SMSPVA_API_KEY`** in your server environment. Keep it out of source control. The example reads your balance and reference lists; it does not create an order.

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    smspva "github.com/SMSPVA/smspva-go"
)

func main() {
    client, err := smspva.New(smspva.Options{APIKey: os.Getenv("SMSPVA_API_KEY")})
    if err != nil { log.Fatal(err) }
    defer client.CloseIdleConnections()
    ctx := context.Background()

    balance, err := client.Activation.GetBalance(ctx)
    if err != nil { log.Fatal(err) }
    fmt.Println(string(balance.Body))
    countries, err := client.Rental.GetCountries(ctx)
    if err != nil { log.Fatal(err) }
    fmt.Println(string(countries.Body))
    providers, err := client.Email.GetProviders(ctx)
    if err != nil { log.Fatal(err) }
    fmt.Println(string(providers.Body))
}
```

## Requests and responses

Every method takes a context, a typed parameter struct when it has parameters, and optional request options. Parameter fields are pointers; nil omits a field and `Ptr(value)` supplies one. Required fields are checked before making a request. Integer fields use `int64`; strings, enums and bounds follow the original documentation. [METHODS.md](https://github.com/SMSPVA/smspva-go/blob/main/METHODS.md) lists all methods.

```go
services, err := client.Email.GetData(ctx, smspva.EmailGetDataParams{
    Pcode: smspva.Ptr("CODE_FROM_GET_PROVIDERS"),
})
if err != nil { return err }
fmt.Println(string(services.Body))
```

Country, service, provider, price and availability methods fetch live API data; the SDK has no fixed lists. Public methods work without a key and never send the configured key. Omitted or empty Email `pcode` is not sent. Optional-key Rental calls can use `WithoutAPIKey()`; required-key methods still require and send it. Activation sends `apikey` as a header; Rental and Email use query `apikey`.

`Response.Body` is the original JSON bytes and `HTTPStatus` is the actual HTTP status. Prices, IDs, timestamps and envelope fields are not normalized. Use `response.Decode(&yourStruct)` to read selected fields into your own types, or `json.Decoder.UseNumber` if decoding arbitrary numbers into maps. Generated response DTOs are not included.

Receiving an existing order's new SMS can charge your account:

```go
sms, err := client.Activation.ReceiveSms(ctx, smspva.ActivationReceiveSmsParams{
    OrderID: smspva.Ptr(int64(12345)),
})
if err != nil { return err }
if smspva.IsPending(sms) {
    fmt.Println("SMS has not arrived yet")
} else {
    fmt.Println(string(sms.Body))
}
```

`IsPending` recognizes HTTP 202 / `SMS_NOT_YET_RECEIVED`. Email creation `status=2` is also an ordinary result; use `IsNoStock`. For Rental multi-service counts, send `dtype` and `dcount` explicitly; backend defaults are unreliable according to the source documentation.

## Errors, timeout and cancellation

Use `errors.As` for `*smspva.APIError`, `*smspva.ProtocolError`, `*smspva.TransportError`, `*smspva.ValidationError` or `*smspva.PollingTimeoutError`. `APIError` retains the original `HTTPStatus`, `Code` and JSON `Body`. Codes decoded as numbers use `json.Number`.

The default whole-request timeout is 30 seconds, including reading the response body. Configure `Options.Timeout` or per-call `WithTimeout(duration)`. A shorter context deadline wins. Cancelling the context aborts the request; context sentinel errors can be checked using `errors.Is`. `TransportError.Kind` is `network`, `timeout` or `cancelled`; `OutcomeUnknown` flags uncertain paid/state-changing operations. Error messages omit credentials and URLs. Raw bodies can contain private message data; redact your own logs.

The default transport uses fresh HTTP/1.1 connections with keep-alives disabled, without environment proxies, cookie storage or redirects. This avoids the standard [Go transport's automatic replay of idempotent requests on reused connections](https://pkg.go.dev/net/http#Transport). Several paid API methods use GET. The SDK makes no automatic retries; inspect existing orders after an uncertain paid result before repeating it. An injected `Options.Transport` is the caller's responsibility: do not install a retrying transport or a transport that adds credentials to public requests.

Base URLs can be overridden in `Options` for controlled proxies or tests. Keys, URLs and headers are not method parameters.

## Waiting for SMS

```go
result, err := client.Activation.WaitForSMS(ctx, smspva.WaitForSMSOptions{
    OrderID: 12345,
    ExpiresAt: savedOrderExpiry,
})
```

Save an absolute expiry when the order is obtained and pass that same `time.Time` each time. The helper converts it once to a monotonic deadline, waits at least five seconds between pending requests, and never creates or cancels orders. It stops on a terminal API error, context cancellation, network failure or expiry. If an in-flight request times out at the deadline, its transport error retains the uncertain outcome; expiry while sleeping returns `PollingTimeoutError`.

## Development

From the repository root:

```sh
node scripts/generate.mjs
gofmt -w resources.go named_calls_test.go
go test ./...
go vet ./...
```

Generation uses the bundled `contracts.json` and `testdata/cases.json`. The generated source, embedded contracts and test fixtures are included, so an extracted module builds and tests on its own. Tests use loopback HTTP and cover every named method, auth and encoding, original responses, pending/no-stock, validation, redirects/cookies, a disconnected paid GET without replay, timeouts, cancellation and saved deadlines. No production API is called.

Tests do not validate paid production behavior.

## License

MIT. See [LICENSE](https://github.com/SMSPVA/smspva-go/blob/main/LICENSE).

## Documentation and support

Read the [SMSPVA API documentation](https://docs.smspva.com/) and the [method index](https://github.com/SMSPVA/smspva-go/blob/main/METHODS.md) for parameters, response codes and workflows. SMSPVA supports [phone number rentals](https://smspva.com/rent) and [private email rentals](https://smspva.com/rent-mail.html); availability and prices come from the live API.

For account or billing questions, visit the [SMSPVA help center](https://smspva.com/help.html). Report SDK bugs through [GitHub issues](https://github.com/SMSPVA/smspva-go/issues), without API keys or private message content.
