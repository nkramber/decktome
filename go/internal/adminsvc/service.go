// Package adminsvc is the admin screen of the owner (D-1076). Each call
// needs the custom claim admin: true on the token, and every other
// caller gets PermissionDenied.
package adminsvc

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/access"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/mail"
)

// SignInURL is the page the approval email names.
const SignInURL = "https://decktome.com/sign-in"

// Requests is the store of the requests for beta access (D-1075).
type Requests interface {
	List(ctx context.Context, state string) ([]access.Request, error)
	Get(ctx context.Context, email string) (access.Request, error)
	Decide(ctx context.Context, email, state string, now time.Time) error
}

// Invite puts one email on the invite list (D-420).
type Invite func(ctx context.Context, email string) error

// Server answers the admin calls.
type Server struct {
	requests Requests
	invite   Invite
	mailer   mail.Sender
	log      *slog.Logger
	now      func() time.Time
}

// Option tunes the server.
type Option func(*Server)

// WithMailer sends the approval email (D-1077). A server with no mailer
// approves and sends nothing.
func WithMailer(m mail.Sender) Option { return func(s *Server) { s.mailer = m } }

// WithClock replaces time.Now (tests).
func WithClock(now func() time.Time) Option { return func(s *Server) { s.now = now } }

// New builds the server.
func New(requests Requests, invite Invite, log *slog.Logger, opts ...Option) *Server {
	s := &Server{requests: requests, invite: invite, log: log, now: time.Now}
	for _, o := range opts {
		o(s)
	}
	return s
}

var errNotAdmin = errors.New("this screen is open to the admin alone")

func admin(ctx context.Context) error {
	if !auth.IsAdmin(ctx) {
		return connect.NewError(connect.CodePermissionDenied, errNotAdmin)
	}
	return nil
}

var states = map[mtgv1.AccessStatus]string{
	mtgv1.AccessStatus_ACCESS_STATUS_UNSPECIFIED: access.Pending,
	mtgv1.AccessStatus_ACCESS_STATUS_PENDING:     access.Pending,
	mtgv1.AccessStatus_ACCESS_STATUS_APPROVED:    access.Approved,
	mtgv1.AccessStatus_ACCESS_STATUS_DISMISSED:   access.Dismissed,
}

func statusOf(state string) mtgv1.AccessStatus {
	switch state {
	case access.Pending:
		return mtgv1.AccessStatus_ACCESS_STATUS_PENDING
	case access.Approved:
		return mtgv1.AccessStatus_ACCESS_STATUS_APPROVED
	case access.Dismissed:
		return mtgv1.AccessStatus_ACCESS_STATUS_DISMISSED
	}
	return mtgv1.AccessStatus_ACCESS_STATUS_UNSPECIFIED
}

// ListAccessRequests reads the requests in one state, the newest first.
func (s *Server) ListAccessRequests(ctx context.Context, req *connect.Request[mtgv1.ListAccessRequestsRequest]) (*connect.Response[mtgv1.ListAccessRequestsResponse], error) {
	if err := admin(ctx); err != nil {
		return nil, err
	}
	state, ok := states[req.Msg.GetStatus()]
	if !ok {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("unknown state"))
	}
	list, err := s.requests.List(ctx, state)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	out := make([]*mtgv1.AccessRequest, 0, len(list))
	for _, r := range list {
		out = append(out, &mtgv1.AccessRequest{
			Email:     r.Email,
			Note:      r.Note,
			Status:    statusOf(r.Status),
			Count:     r.Count,
			CreatedAt: timestamppb.New(r.CreatedAt),
			UpdatedAt: timestamppb.New(r.UpdatedAt),
		})
	}
	return connect.NewResponse(&mtgv1.ListAccessRequestsResponse{Requests: out}), nil
}

// ApproveAccessRequest puts the email on the invite list, marks the
// request approved, and sends the approval email (D-1076, D-1077). A
// failed send keeps the approval, and the answer says so.
func (s *Server) ApproveAccessRequest(ctx context.Context, req *connect.Request[mtgv1.ApproveAccessRequestRequest]) (*connect.Response[mtgv1.ApproveAccessRequestResponse], error) {
	if err := admin(ctx); err != nil {
		return nil, err
	}
	r, err := s.find(ctx, req.Msg.GetEmail())
	if err != nil {
		return nil, err
	}
	if err := s.invite(ctx, r.Email); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if err := s.requests.Decide(ctx, r.Email, access.Approved, s.now()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	sent := false
	if s.mailer != nil {
		mctx, cancel := context.WithTimeout(ctx, mail.Timeout)
		defer cancel()
		if err := s.mailer.Send(mctx, approval(r.Email)); err != nil {
			s.log.Warn("admin: the approval email did not send", "err", err)
		} else {
			sent = true
		}
	}
	return connect.NewResponse(&mtgv1.ApproveAccessRequestResponse{EmailSent: sent}), nil
}

// DismissAccessRequest marks the request dismissed. The email stays off
// the invite list, and no email goes out.
func (s *Server) DismissAccessRequest(ctx context.Context, req *connect.Request[mtgv1.DismissAccessRequestRequest]) (*connect.Response[mtgv1.DismissAccessRequestResponse], error) {
	if err := admin(ctx); err != nil {
		return nil, err
	}
	r, err := s.find(ctx, req.Msg.GetEmail())
	if err != nil {
		return nil, err
	}
	if err := s.requests.Decide(ctx, r.Email, access.Dismissed, s.now()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.DismissAccessRequestResponse{}), nil
}

func (s *Server) find(ctx context.Context, email string) (access.Request, error) {
	r, err := s.requests.Get(ctx, email)
	if errors.Is(err, access.ErrNotFound) {
		return access.Request{}, connect.NewError(connect.CodeNotFound, err)
	}
	if err != nil {
		return access.Request{}, connect.NewError(connect.CodeInternal, err)
	}
	return r, nil
}

// approval is the email of an approved request (D-1077).
func approval(to string) mail.Message {
	return mail.Message{
		To:      to,
		Subject: "Your Deck Tome beta access is ready",
		Text: "Your email is now on the Deck Tome beta list.\n\n" +
			"Create your account with this email address at " + SignInURL + "\n\n" +
			"You get this email because you asked for beta access on Deck Tome.",
	}
}
