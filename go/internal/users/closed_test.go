package users

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestClosedCacheReadsOnceAMinute: the interceptor asks for every
// request, so a read serves one minute for each user, and a failed read
// keeps nothing (D-941).
func TestClosedCacheReadsOnceAMinute(t *testing.T) {
	now := time.Date(2026, 9, 25, 12, 0, 0, 0, time.UTC)
	reads := map[string]int{}
	closed := map[string]bool{"u1": false}
	var fail error
	c := NewClosedCache(func(_ context.Context, uid string) (bool, error) {
		reads[uid]++
		return closed[uid], fail
	}, func() time.Time { return now })
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		if got, err := c.Closed(ctx, "u1"); err != nil || got {
			t.Fatalf("read %d: %v %v", i, got, err)
		}
	}
	if reads["u1"] != 1 {
		t.Errorf("three calls read the store %d times, want 1", reads["u1"])
	}
	closed["u1"] = true
	now = now.Add(ClosedTTL)
	if got, _ := c.Closed(ctx, "u1"); !got {
		t.Error("a close was not read after the TTL")
	}
	fail = errors.New("store down")
	if _, err := c.Closed(ctx, "u2"); err == nil {
		t.Error("a failed read answered no error")
	}
	fail = nil
	if _, err := c.Closed(ctx, "u2"); err != nil || reads["u2"] != 2 {
		t.Errorf("a failed read was kept: reads %d, err %v", reads["u2"], err)
	}
}
