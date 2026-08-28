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

// Check reports the service state with the snapshot state. The status
// is "starting" until the first snapshot loads and "ok" after (L-14).
func (s *Server) Check(_ context.Context, _ *connect.Request[mtgv1.CheckRequest]) (*connect.Response[mtgv1.CheckResponse], error) {
	return connect.NewResponse(s.Status()), nil
}

// Ready reports whether a card index is loaded. The RPCs that need one
// refuse until then, so a router must not send traffic before it.
func (s *Server) Ready() bool {
	return s.source != nil && s.source.Current() != nil
}

// Status builds the CheckResponse. The plain /healthz and /readyz routes
// use the same values, so every probe agrees.
func (s *Server) Status() *mtgv1.CheckResponse {
	res := &mtgv1.CheckResponse{Status: "starting", Version: s.version, CardSnapshot: "none", CardSnapshotAgeHours: -1}
	if s.source == nil {
		return res
	}
	if idx := s.source.Current(); idx != nil {
		res.Status = "ok"
		res.CardSnapshot = idx.AsOf.UTC().Format(time.RFC3339)
		res.CardSnapshotAgeHours = s.now().Sub(idx.AsOf).Hours()
	}
	return res
}
