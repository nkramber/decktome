// Package ratelimit bounds the public reads of the API per client
// address (D-315). Behind Firebase Hosting and Cloud Run the remote
// address is the proxy, so the limiter reads the client address from
// X-Forwarded-For first, and every visitor gets a bucket of their own.
package ratelimit

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"time"

	"connectrpc.com/connect"
)

// Limiter counts the calls of each client address in a fixed window.
type Limiter struct {
	limit  int
	window time.Duration
	now    func() time.Time
	mu     sync.Mutex
	starts map[string]time.Time
	counts map[string]int
	sweep  time.Time
}

// New makes a limiter of limit calls per window. Zero or less means no
// limit at all.
func New(limit int, window time.Duration) *Limiter {
	return &Limiter{limit: limit, window: window, now: time.Now, starts: map[string]time.Time{}, counts: map[string]int{}}
}

// WithClock replaces the clock, for the tests.
func (l *Limiter) WithClock(now func() time.Time) *Limiter {
	l.now = now
	return l
}

// Allow counts one call of a key and says whether it stays under the
// limit. The window starts at the first call and restarts when it ends.
func (l *Limiter) Allow(key string) bool {
	if l.limit <= 0 {
		return true
	}
	now := l.now()
	l.mu.Lock()
	defer l.mu.Unlock()
	if now.Sub(l.sweep) > l.window {
		for k, start := range l.starts {
			if now.Sub(start) > l.window {
				delete(l.starts, k)
				delete(l.counts, k)
			}
		}
		l.sweep = now
	}
	start, ok := l.starts[key]
	if !ok || now.Sub(start) >= l.window {
		l.starts[key] = now
		l.counts[key] = 1
		return true
	}
	l.counts[key]++
	return l.counts[key] <= l.limit
}

// ClientAddress reads the client behind a proxy: the first address of
// X-Forwarded-For, else the host of the remote address, else the remote
// address as it came.
func ClientAddress(forwardedFor, remoteAddr string) string {
	if first, _, _ := strings.Cut(forwardedFor, ","); strings.TrimSpace(first) != "" {
		return strings.TrimSpace(first)
	}
	if host, _, err := net.SplitHostPort(remoteAddr); err == nil {
		return host
	}
	return strings.TrimSpace(remoteAddr)
}

var errTooMany = errors.New("too many calls from this address, try again in a minute")

// Interceptor limits the named procedures. Every other procedure passes
// untouched. A refused call answers ResourceExhausted.
func (l *Limiter) Interceptor(procedures ...string) connect.Interceptor {
	limited := make(map[string]bool, len(procedures))
	for _, p := range procedures {
		limited[p] = true
	}
	return &interceptor{limiter: l, limited: limited}
}

type interceptor struct {
	limiter *Limiter
	limited map[string]bool
}

func (i *interceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		if req.Spec().IsClient || !i.limited[req.Spec().Procedure] {
			return next(ctx, req)
		}
		key := ClientAddress(req.Header().Get("X-Forwarded-For"), req.Peer().Addr)
		if !i.limiter.Allow(key) {
			return nil, connect.NewError(connect.CodeResourceExhausted, errTooMany)
		}
		return next(ctx, req)
	}
}

func (i *interceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (i *interceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		if !i.limited[conn.Spec().Procedure] {
			return next(ctx, conn)
		}
		key := ClientAddress(conn.RequestHeader().Get("X-Forwarded-For"), conn.Peer().Addr)
		if !i.limiter.Allow(key) {
			return connect.NewError(connect.CodeResourceExhausted, errTooMany)
		}
		return next(ctx, conn)
	}
}
