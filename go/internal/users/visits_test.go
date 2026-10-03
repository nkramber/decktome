package users

import (
	"context"
	"errors"
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
	}, func() time.Time { return now })
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
