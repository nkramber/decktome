package llm

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Class answers one question about a failure: can a retry help, and how?
// The Client handles each class differently (connector-syncer lesson 2).
type Class int

const (
	// ClassTerminal means no retry helps. Bad request, auth, unknown model.
	ClassTerminal Class = iota
	// ClassTransient means a retry after a backoff can help. 429, 5xx, timeouts.
	ClassTransient
	// ClassTruncation means the output hit the token cap. One immediate retry
	// with a higher cap can help. The same cap gives the same truncation.
	ClassTruncation
	// ClassRefusal means the model declined. A retry gives the same answer.
	ClassRefusal
	// ClassSchema means the output did not satisfy the schema. The output
	// is sampled, so the Client retries once. A second miss is terminal.
	ClassSchema
	// ClassBudget means the attempt or time budget ran out.
	ClassBudget
)

func (c Class) String() string {
	switch c {
	case ClassTerminal:
		return "terminal"
	case ClassTransient:
		return "transient"
	case ClassTruncation:
		return "truncation"
	case ClassRefusal:
		return "refusal"
	case ClassSchema:
		return "schema"
	case ClassBudget:
		return "budget"
	}
	return fmt.Sprintf("class(%d)", int(c))
}

// Error is every failure the package returns.
type Error struct {
	Class    Class
	Provider string
	Model    string
	// Status is the HTTP status when a provider answered, else 0.
	Status int
	// RetryAfter is the provider's Retry-After hint, or 0. The Client
	// waits at least this long before a transient retry (L-6).
	RetryAfter time.Duration
	Err        error
}

func (e *Error) Error() string {
	if e.Status != 0 {
		return fmt.Sprintf("llm %s/%s: %s (http %d): %v", e.Provider, e.Model, e.Class, e.Status, e.Err)
	}
	return fmt.Sprintf("llm %s/%s: %s: %v", e.Provider, e.Model, e.Class, e.Err)
}

func (e *Error) Unwrap() error { return e.Err }

// ClassOf reads the retry class of err. An error the package did not make
// is terminal: an unknown failure must not loop.
func ClassOf(err error) Class {
	var e *Error
	if errors.As(err, &e) {
		return e.Class
	}
	return ClassTerminal
}

func newErr(class Class, provider, model string, status int, err error) *Error {
	return &Error{Class: class, Provider: provider, Model: model, Status: status, Err: err}
}

// retryAfter reads the Retry-After header of resp as a delay. The header
// is seconds or an HTTP-date. A missing or unreadable header is 0.
func retryAfter(resp *http.Response, now time.Time) time.Duration {
	if resp == nil {
		return 0
	}
	v := resp.Header.Get("Retry-After")
	if v == "" {
		return 0
	}
	if secs, err := strconv.ParseFloat(v, 64); err == nil {
		if secs <= 0 {
			return 0
		}
		return time.Duration(secs * float64(time.Second))
	}
	if t, err := http.ParseTime(v); err == nil {
		if d := t.Sub(now); d > 0 {
			return d
		}
	}
	return 0
}

// withRetryAfter returns e with the Retry-After hint of resp set.
func (e *Error) withRetryAfter(resp *http.Response) *Error {
	e.RetryAfter = retryAfter(resp, time.Now())
	return e
}
