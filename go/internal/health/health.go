// Package health implements the HealthService. It reports the build
// version and the loaded card snapshot, so a probe can tell a stale
// snapshot from a dead process.
package health

import (
	"context"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// IndexSource gives the current card index, or nil before the first load.
// cardsvc.Server satisfies it.
type IndexSource interface {
	Current() *cards.Index
}

// Server answers HealthService requests.
type Server struct {
	mtgv1connect.UnimplementedHealthServiceHandler
	version string
	source  IndexSource
	now     func() time.Time
}

// New returns a Server that reports the build version and the snapshot
// state of source. A nil source reports "none".
func New(version string, source IndexSource) *Server {
	return &Server{version: version, source: source, now: time.Now}
}

// Check reports that the service is up, with the snapshot state.
func (s *Server) Check(_ context.Context, _ *connect.Request[mtgv1.CheckRequest]) (*connect.Response[mtgv1.CheckResponse], error) {
	return connect.NewResponse(s.Status()), nil
}

// Status builds the CheckResponse. The plain /healthz route uses the
// same values, so both probes agree.
func (s *Server) Status() *mtgv1.CheckResponse {
	res := &mtgv1.CheckResponse{Status: "ok", Version: s.version, CardSnapshot: "none", CardSnapshotAgeHours: -1}
	if s.source == nil {
		return res
	}
	if idx := s.source.Current(); idx != nil {
		res.CardSnapshot = idx.AsOf.UTC().Format(time.RFC3339)
		res.CardSnapshotAgeHours = s.now().Sub(idx.AsOf).Hours()
	}
	return res
}
