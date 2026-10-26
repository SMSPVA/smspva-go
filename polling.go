package smspva

import (
	"context"
	"time"
)

// WaitForSMSOptions uses a saved order expiry, never a new relative timeout on each poll.
type WaitForSMSOptions struct {
	OrderID    int64
	ExpiresAt  time.Time
	Interval   time.Duration
	PartnerKey *string
}

// WaitForSMS polls an existing activation. A newly received SMS may charge the account.
// It never purchases, retries failed requests, or cancels an order.
func (r *Activation) WaitForSMS(ctx context.Context, options WaitForSMSOptions) (Response, error) {
	interval := options.Interval
	if interval == 0 {
		interval = 5 * time.Second
	}
	if interval < 5*time.Second {
		return Response{}, &ValidationError{"Interval must be at least five seconds in this helper"}
	}
	if options.ExpiresAt.IsZero() {
		return Response{}, &ValidationError{"A saved order expiry is required"}
	}
	if ctx == nil {
		return Response{}, &ValidationError{"Context cannot be nil"}
	}
	// Convert a persisted wall-clock expiry once to a monotonic deadline.
	remaining := time.Until(options.ExpiresAt)
	if remaining <= 0 {
		return Response{}, &PollingTimeoutError{}
	}
	deadline := time.Now().Add(remaining)
	bounded, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	params := ActivationReceiveSmsParams{OrderID: Ptr(options.OrderID), PartnerKey: options.PartnerKey}
	for time.Until(deadline) > 0 {
		result, err := r.ReceiveSms(bounded, params)
		if err != nil {
			return Response{}, err
		}
		if !IsPending(result) {
			return result, nil
		}
		left := time.Until(deadline)
		if left <= 0 {
			return Response{}, &PollingTimeoutError{}
		}
		delay := interval
		if left < delay {
			delay = left
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			kind := "cancelled"
			if ctx.Err() == context.DeadlineExceeded {
				kind = "timeout"
			}
			return Response{}, &TransportError{Family: "activation", Operation: "activation.receiveSms", Kind: kind, cause: ctx.Err()}
		case <-timer.C:
		}
		if left <= interval {
			return Response{}, &PollingTimeoutError{}
		}
	}
	return Response{}, &PollingTimeoutError{}
}
