package allowlist

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestAllowedReadsTheListOnceAMinute is D-420: the first call reads the
// document, the calls inside the minute read the cache, and the call
// after it reads again. An email is one invite whatever its case, and an
// empty email is never on the list.
func TestAllowedReadsTheListOnceAMinute(t *testing.T) {
	reads := 0
	emails := []string{"Ann@Example.com", " bob@example.com "}
	now := time.Unix(1000, 0)
	l := New(func(context.Context) ([]string, error) {
		reads++
		return emails, nil
	}).WithClock(func() time.Time { return now })
	ctx := context.Background()
	for _, email := range []string{"ann@example.com", "ANN@example.com", "bob@example.com"} {
		if ok, err := l.Allowed(ctx, email); err != nil || !ok {
			t.Errorf("%q: allowed = %v, %v, want true", email, ok, err)
		}
	}
	if ok, _ := l.Allowed(ctx, "cy@example.com"); ok {
		t.Error("an email off the list was allowed")
	}
	if ok, _ := l.Allowed(ctx, ""); ok {
		t.Error("an empty email was allowed")
	}
	if reads != 1 {
		t.Errorf("reads = %d, want one inside the minute", reads)
	}
	emails = append(emails, "cy@example.com")
	now = now.Add(TTL)
	if ok, _ := l.Allowed(ctx, "cy@example.com"); !ok || reads != 2 {
		t.Errorf("after the minute: allowed = %v, reads = %d, want true and 2", ok, reads)
	}
}

// TestAllowedAnswersTheReadError: a list that can not be read answers
// the error and allows nobody.
func TestAllowedAnswersTheReadError(t *testing.T) {
	l := New(func(context.Context) ([]string, error) { return nil, errors.New("down") })
	if ok, err := l.Allowed(context.Background(), "ann@example.com"); ok || err == nil {
		t.Errorf("allowed = %v, err = %v, want false and the error", ok, err)
	}
}

func TestNormalize(t *testing.T) {
	if got := Normalize("  Ann@Example.COM "); got != "ann@example.com" {
		t.Errorf("Normalize = %q", got)
	}
}
