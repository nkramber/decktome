// Package invitesvc answers whether one email may use the app, before
// the person has an account (D-592).
//
// The create-account form asks this first. Without it the browser makes
// the Firebase account, the API then refuses every call, and the person
// holds an account the project never invited (F-59).
package invitesvc

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/access"
	"github.com/nkramber/decktome/go/internal/notify"
	"github.com/nkramber/decktome/go/internal/ratelimit"
)

// Allowlist is the invite list. It is the same interface the auth
// interceptor reads, so both answer from one cached list (D-420).
type Allowlist interface {
	Allowed(ctx context.Context, email string) (bool, error)
}

// Requests is the store of the requests for beta access (D-1075).
type Requests interface {
	Put(ctx context.Context, email, note string, now time.Time) (created bool, err error)
}

// Notifier sends the owner a notice off the request path (D-892).
type Notifier interface {
	Notify(n notify.Notice)
}

// NoticesPerHour caps the notices of new requests across all callers.
// The limit per client address bounds one caller, and this cap bounds a
// caller with many addresses. A request past the cap stays stored.
const NoticesPerHour = 20

// Server answers CheckInvite and RequestAccess. A nil list allows every
// email, which is local mode with no cloud dependency (guardrail 9).
type Server struct {
	list     Allowlist
	requests Requests
	notifier Notifier
	notices  *ratelimit.Limiter
	now      func() time.Time
}

// Option tunes the server.
type Option func(*Server)

// WithRequests stores the requests for beta access (D-1075). A server
// with no store answers RequestAccess with Unimplemented.
func WithRequests(r Requests) Option { return func(s *Server) { s.requests = r } }

// WithNotifier sends the owner a notice of each new request (D-1075).
func WithNotifier(n Notifier) Option { return func(s *Server) { s.notifier = n } }

// WithClock replaces time.Now (tests).
func WithClock(now func() time.Time) Option { return func(s *Server) { s.now = now } }

// New builds the server over a list. Pass nil for a server with no list.
func New(list Allowlist, opts ...Option) *Server {
	s := &Server{list: list, now: time.Now}
	for _, o := range opts {
		o(s)
	}
	s.notices = ratelimit.New(NoticesPerHour, time.Hour).WithClock(s.now)
	return s
}

// CheckInvite answers whether the invite list holds the email.
//
// The answer is a bare yes or no. It says nothing about an account, so
// it tells a caller no more than the sign-in form already does.
//
// A list that can not be read answers Unavailable, and never a false.
// A false would send a person away who is on the list.
func (s *Server) CheckInvite(ctx context.Context, req *connect.Request[mtgv1.CheckInviteRequest]) (*connect.Response[mtgv1.CheckInviteResponse], error) {
	if s.list == nil {
		return connect.NewResponse(&mtgv1.CheckInviteResponse{Allowed: true}), nil
	}
	ok, err := s.list.Allowed(ctx, req.Msg.GetEmail())
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, errListUnreadable)
	}
	return connect.NewResponse(&mtgv1.CheckInviteResponse{Allowed: ok}), nil
}

// errListUnreadable is the same sentence the interceptor answers, so the
// two paths read alike (D-314).
var errListUnreadable = errors.New("the invite list could not be read")

// RequestAccess stores one request for beta access (D-1075). An email on
// the invite list stores nothing and answers already_invited, because
// the person can create an account now. The check answers no more than
// CheckInvite does.
//
// A repeat request updates its record. The answer is the same for a new
// request, a repeat, and a dismissed request, so it tells the caller
// nothing about the decision of the owner.
func (s *Server) RequestAccess(ctx context.Context, req *connect.Request[mtgv1.RequestAccessRequest]) (*connect.Response[mtgv1.RequestAccessResponse], error) {
	email, note, err := access.Clean(req.Msg.GetEmail(), req.Msg.GetNote())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	if s.list != nil {
		ok, err := s.list.Allowed(ctx, email)
		if err != nil {
			return nil, connect.NewError(connect.CodeUnavailable, errListUnreadable)
		}
		if ok {
			return connect.NewResponse(&mtgv1.RequestAccessResponse{AlreadyInvited: true}), nil
		}
	}
	if s.requests == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errors.New("this server takes no requests for access"))
	}
	now := s.now()
	created, err := s.requests.Put(ctx, email, note, now)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if created && s.notifier != nil && s.notices.Allow("access") {
		s.notifier.Notify(noticeOf(email, note, now))
	}
	return connect.NewResponse(&mtgv1.RequestAccessResponse{}), nil
}

// noticeOf writes the notice of one new request. The admin screen holds
// the decision, so the notice names no action.
func noticeOf(email, note string, now time.Time) notify.Notice {
	var b strings.Builder
	fmt.Fprintf(&b, "email: %s\n", email)
	if note != "" {
		fmt.Fprintf(&b, "note: %s\n", note)
	}
	fmt.Fprintf(&b, "time: %s", now.UTC().Format("2006-01-02 15:04 UTC"))
	return notify.Notice{Title: "decktome: request for beta access", Message: b.String()}
}
