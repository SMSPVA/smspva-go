package smspva

import (
	"encoding/json"
	"fmt"
)

type ValidationError struct{ Message string }

func (e *ValidationError) Error() string { return e.Message }

type APIError struct {
	Family     string
	Operation  string
	HTTPStatus int
	Code       any
	Body       json.RawMessage
}

func (e *APIError) Error() string {
	return fmt.Sprintf("SMSPVA %s request failed (HTTP %d)", e.Family, e.HTTPStatus)
}

type ProtocolError struct {
	Family, Operation string
	HTTPStatus        int
	OutcomeUnknown    bool
}

func (e *ProtocolError) Error() string {
	return fmt.Sprintf("SMSPVA %s returned an invalid response (HTTP %d)", e.Family, e.HTTPStatus)
}

type TransportError struct {
	Family, Operation, Kind string
	OutcomeUnknown          bool
	cause                   error
}

func (e *TransportError) Error() string {
	return fmt.Sprintf("SMSPVA %s request %s; no automatic retry", e.Family, e.Kind)
}

// Only context sentinel errors are unwrapped; raw transport URLs and credentials are not exposed.
func (e *TransportError) Unwrap() error { return e.cause }

type PollingTimeoutError struct{}

func (e *PollingTimeoutError) Error() string { return "No SMS result before the saved order deadline" }
