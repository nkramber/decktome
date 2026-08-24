// Package health implements the HealthService used to prove the proto pipeline.
package health

import (
	"context"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
)

// Server answers HealthService requests.
type Server struct {
	mtgv1connect.UnimplementedHealthServiceHandler
	version string
}

// New returns a Server that reports the given build version.
func New(version string) *Server {
	return &Server{version: version}
}

// Check reports that the service is up.
func (s *Server) Check(_ context.Context, _ *connect.Request[mtgv1.CheckRequest]) (*connect.Response[mtgv1.CheckResponse], error) {
	return connect.NewResponse(&mtgv1.CheckResponse{Status: "ok", Version: s.version}), nil
}
