package users

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"
)

// TestVisitsWriteOnceInTheTTL: the interceptor records every verified
// call, so one write serves VisitTTL for each user, and a failed write
// keeps nothing (D-1093).
func TestVisitsWriteOnceInTheTTL(t *testing.T) {
	now := time.Date(2026, 10, 3, 19, 25, 0, 0, time.UTC)
	writes := map[string][]time.Time{}
	var fail error
	v := NewVisits(func(_ context.Context, uid, email string, at time.Time) error {
		if email != "ann@example.com" {
			t.Errorf("email = %q", email)
		}
		writes[uid] = append(writes[uid], at)
		return fail
	}, func() time.Time { return now }, nil)
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		v.Seen(ctx, "u1", "ann@example.com")
	}
	if len(writes["u1"]) != 1 {
		t.Errorf("three calls wrote %d times, want 1", len(writes["u1"]))
	}
	now = now.Add(VisitTTL - time.Second)
	v.Seen(ctx, "u1", "ann@example.com")
	if len(writes["u1"]) != 1 {
		t.Error("a call inside the TTL wrote again")
	}
	now = now.Add(time.Second)
	v.Seen(ctx, "u1", "ann@example.com")
	if got := writes["u1"]; len(got) != 2 || !got[1].Equal(now) {
		t.Errorf("a call after the TTL wrote %v, want a second write at %v", got, now)
	}
	fail = errors.New("store down")
	v.Seen(ctx, "u2", "ann@example.com")
	fail = nil
	v.Seen(ctx, "u2", "ann@example.com")
	if len(writes["u2"]) != 2 {
		t.Errorf("a failed write was kept: %d writes, want 2", len(writes["u2"]))
	}
	v.Seen(ctx, "", "ann@example.com")
	if len(writes[""]) != 0 {
		t.Error("a call with no user wrote a record")
	}
}

// TestVisitsOfOnePageWriteOnce: the home page sends three calls at once,
// and only the first one writes. The others find the slot claimed.
func TestVisitsOfOnePageWriteOnce(t *testing.T) {
	now := time.Date(2026, 10, 3, 19, 25, 0, 0, time.UTC)
	release := make(chan struct{})
	var mu sync.Mutex
	writes := 0
	v := NewVisits(func(context.Context, string, string, time.Time) error {
		mu.Lock()
		writes++
		first := writes == 1
		mu.Unlock()
		// Only the first write waits, so a second write ends the test
		// with a count and not a hang.
		if first {
			<-release
		}
		return nil
	}, func() time.Time { return now }, nil)
	var wg sync.WaitGroup
	wg.Go(func() { v.Seen(context.Background(), "u1", "ann@example.com") })
	// The first call holds its write open, so the slot is claimed.
	for {
		mu.Lock()
		started := writes == 1
		mu.Unlock()
		if started {
			break
		}
		time.Sleep(time.Millisecond)
	}
	v.Seen(context.Background(), "u1", "ann@example.com")
	v.Seen(context.Background(), "u1", "ann@example.com")
	close(release)
	wg.Wait()
	if writes != 1 {
		t.Errorf("three calls of one page wrote %d times, want 1", writes)
	}
}

// TestAFailedVisitIsLoggedWithNoEmail: a write that fails leaves a trace
// with the uid, and the address stays out of the log.
func TestAFailedVisitIsLoggedWithNoEmail(t *testing.T) {
	var buf bytes.Buffer
	v := NewVisits(func(context.Context, string, string, time.Time) error {
		return errors.New("permission denied")
	}, nil, slog.New(slog.NewTextHandler(&buf, nil)))
	v.Seen(context.Background(), "u-9", "ann@example.com")
	out := buf.String()
	if !strings.Contains(out, "user visit not recorded") || !strings.Contains(out, "u-9") || !strings.Contains(out, "permission denied") {
		t.Errorf("log = %q, want the message, the uid, and the error", out)
	}
	if strings.Contains(out, "ann@example.com") {
		t.Errorf("log = %q, and it holds the email", out)
	}
}
