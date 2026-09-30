// Package pushsvc serves PushService (PR-26, D-1004, D-1005). It keeps
// the devices of the caller that take a web push. The send itself runs
// in the agent service, when a build ends after the user left.
package pushsvc

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/decktome/go/internal/auth"
	"github.com/nkramber/decktome/go/internal/push"
)

// Store keeps the devices of each user.
type Store interface {
	Add(ctx context.Context, uid, id string, at time.Time) error
	Remove(ctx context.Context, uid, id string) error
	Has(ctx context.Context, uid, id string) (bool, error)
}

// Server answers PushService requests.
type Server struct {
	mtgv1connect.UnimplementedPushServiceHandler
	store  Store
	userFn auth.UserFunc
	now    func() time.Time
}

// New wires the service.
func New(store Store, userFn auth.UserFunc) *Server {
	return &Server{store: store, userFn: userFn, now: time.Now}
}

var (
	errNoUser = errors.New("no user in the request context")
	errBadID  = errors.New("installation_id: name the Firebase Installation ID of this browser")
)

func (s *Server) caller(ctx context.Context, id string) (string, error) {
	uid := s.userFn(ctx)
	if uid == "" {
		return "", connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	if !push.ValidID(id) {
		return "", connect.NewError(connect.CodeInvalidArgument, errBadID)
	}
	return uid, nil
}

// RegisterDevice implements PushService.
func (s *Server) RegisterDevice(ctx context.Context, req *connect.Request[mtgv1.RegisterDeviceRequest]) (*connect.Response[mtgv1.RegisterDeviceResponse], error) {
	uid, err := s.caller(ctx, req.Msg.GetInstallationId())
	if err != nil {
		return nil, err
	}
	if err := s.store.Add(ctx, uid, req.Msg.GetInstallationId(), s.now()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.RegisterDeviceResponse{}), nil
}

// UnregisterDevice implements PushService.
func (s *Server) UnregisterDevice(ctx context.Context, req *connect.Request[mtgv1.UnregisterDeviceRequest]) (*connect.Response[mtgv1.UnregisterDeviceResponse], error) {
	uid, err := s.caller(ctx, req.Msg.GetInstallationId())
	if err != nil {
		return nil, err
	}
	if err := s.store.Remove(ctx, uid, req.Msg.GetInstallationId()); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.UnregisterDeviceResponse{}), nil
}

// GetDevice implements PushService.
func (s *Server) GetDevice(ctx context.Context, req *connect.Request[mtgv1.GetDeviceRequest]) (*connect.Response[mtgv1.GetDeviceResponse], error) {
	uid, err := s.caller(ctx, req.Msg.GetInstallationId())
	if err != nil {
		return nil, err
	}
	ok, err := s.store.Has(ctx, uid, req.Msg.GetInstallationId())
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.GetDeviceResponse{Registered: ok}), nil
}
