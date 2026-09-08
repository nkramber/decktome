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

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// Allowlist is the invite list. It is the same interface the auth
// interceptor reads, so both answer from one cached list (D-420).
type Allowlist interface {
	Allowed(ctx context.Context, email string) (bool, error)
}

// Server answers CheckInvite. A nil list allows every email, which is
// local mode with no cloud dependency (guardrail 9).
type Server struct {
	list Allowlist
}

// New builds the server over a list. Pass nil for a server with no list.
func New(list Allowlist) *Server { return &Server{list: list} }

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
