package ratelimit

import (
	"context"
	"fmt"
	"testing"
	"time"

	"connectrpc.com/connect"
)

// TestSixtyPerMinute is the PR-21 gate line: the 61st call in a minute
// is refused, and the next minute starts fresh.
func TestSixtyPerMinute(t *testing.T) {
	now := time.Date(2026, 9, 3, 12, 0, 0, 0, time.UTC)
	l := New(60, time.Minute).WithClock(func() time.Time { return now })
	for i := 1; i <= 60; i++ {
		if !l.Allow("203.0.113.5") {
			t.Fatalf("call %d refused", i)
		}
	}
	if l.Allow("203.0.113.5") {
		t.Fatal("the 61st call passed")
	}
	if !l.Allow("203.0.113.6") {
		t.Fatal("another address shares the bucket")
	}
	now = now.Add(time.Minute)
	if !l.Allow("203.0.113.5") {
		t.Fatal("the next minute did not start fresh")
	}
	if !New(0, time.Minute).Allow("x") {
		t.Fatal("no limit means every call passes")
	}
}

// TestClientAddressReadsTheForwardedHeader is the CAUTION of the plan:
// behind the proxy the remote address is the proxy, so the forwarded
// address is the key, and the proxy's own address is not. The caller
// writes the left part of the header, so the key is the rightmost client
// address (REV-010).
func TestClientAddressReadsTheForwardedHeader(t *testing.T) {
	for _, tc := range []struct{ forwarded, remote, want string }{
		{"203.0.113.5, 10.0.0.2", "10.0.0.1:443", "203.0.113.5"},
		{" 203.0.113.6 ", "10.0.0.1:443", "203.0.113.6"},
		{"198.51.100.1, 203.0.113.5", "169.254.1.1:443", "203.0.113.5"},
		{"198.51.100.1, 203.0.113.5, 35.191.3.4", "169.254.1.1:443", "203.0.113.5"},
		{"198.51.100.1, 203.0.113.5, 130.211.0.9", "169.254.1.1:443", "203.0.113.5"},
		{"not-an-address, 203.0.113.5", "169.254.1.1:443", "203.0.113.5"},
		{"203.0.113.5, not-an-address", "169.254.1.1:443", "203.0.113.5"},
		{"::ffff:203.0.113.5", "169.254.1.1:443", "203.0.113.5"},
		{"2001:db8::7, 2001:4860::1", "169.254.1.1:443", "2001:4860::1"},
		{"10.0.0.2, 35.191.3.4", "198.51.100.7:52011", "198.51.100.7"},
		{"", "198.51.100.7:52011", "198.51.100.7"},
		{"", "198.51.100.7", "198.51.100.7"},
		{"", "[2001:db8::1]:443", "2001:db8::1"},
	} {
		if got := ClientAddress(tc.forwarded, tc.remote); got != tc.want {
			t.Errorf("ClientAddress(%q, %q) = %q, want %q", tc.forwarded, tc.remote, got, tc.want)
		}
	}
}

// TestASpoofedFirstAddressGetsNoNewBucket is REV-010 of the review of
// 2026-09-24. A caller sent a new first address with each call, and each
// call got a new bucket. The eleventh check of an invite in one minute
// must fail.
func TestASpoofedFirstAddressGetsNoNewBucket(t *testing.T) {
	now := time.Date(2026, 9, 24, 12, 0, 0, 0, time.UTC)
	l := New(10, time.Minute).WithClock(func() time.Time { return now })
	allowed := 0
	for i := range 11 {
		key := ClientAddress(fmt.Sprintf("198.51.100.%d, 203.0.113.5", i+1), "169.254.1.1:443")
		if l.Allow(key) {
			allowed++
		}
	}
	if allowed != 10 {
		t.Errorf("allowed %d of 11 calls with a new first address each, want 10", allowed)
	}
}

// TestInterceptorLimitsTheNamedProcedures: two forwarded addresses get
// two buckets, and a procedure outside the list never counts.
func TestInterceptorLimitsTheNamedProcedures(t *testing.T) {
	now := time.Date(2026, 9, 3, 12, 0, 0, 0, time.UTC)
	l := New(2, time.Minute).WithClock(func() time.Time { return now })
	in := l.Interceptor("/mtg.v1.DeckService/GetSharedDeck")
	calls := 0
	next := in.WrapUnary(func(context.Context, connect.AnyRequest) (connect.AnyResponse, error) {
		calls++
		return nil, nil
	})
	call := func(procedure, forwarded string) error {
		req := connect.NewRequest(&struct{}{})
		req.Header().Set("X-Forwarded-For", forwarded)
		_, err := next(context.Background(), &anyRequest{Request: req, spec: connect.Spec{Procedure: procedure}, peer: connect.Peer{Addr: "10.0.0.1:443"}})
		return err
	}
	for i := 0; i < 2; i++ {
		if err := call("/mtg.v1.DeckService/GetSharedDeck", "203.0.113.5"); err != nil {
			t.Fatalf("call %d: %v", i+1, err)
		}
	}
	err := call("/mtg.v1.DeckService/GetSharedDeck", "203.0.113.5")
	if connect.CodeOf(err) != connect.CodeResourceExhausted {
		t.Fatalf("third call: %v, want ResourceExhausted", err)
	}
	if err := call("/mtg.v1.DeckService/GetSharedDeck", "203.0.113.9"); err != nil {
		t.Fatalf("another forwarded address shares the bucket: %v", err)
	}
	for i := 0; i < 5; i++ {
		if err := call("/mtg.v1.DeckService/GetDeck", "203.0.113.5"); err != nil {
			t.Fatalf("a procedure outside the list was limited: %v", err)
		}
	}
	if calls != 8 {
		t.Errorf("calls = %d, want 8", calls)
	}
}

// anyRequest gives a request the spec and the peer a server sees.
type anyRequest struct {
	*connect.Request[struct{}]
	spec connect.Spec
	peer connect.Peer
}

func (r *anyRequest) Spec() connect.Spec { return r.spec }
func (r *anyRequest) Peer() connect.Peer { return r.peer }
