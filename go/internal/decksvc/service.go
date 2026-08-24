// Package decksvc serves DeckService. PR-5 wires Validate. Get, List,
// and Export arrive with PR-8 and PR-13.
package decksvc

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// IndexSource hands out the current card index.
type IndexSource interface {
	Current() *cards.Index
}

// Server answers DeckService requests.
type Server struct {
	mtgv1connect.UnimplementedDeckServiceHandler
	cfg   *rules.Config
	index IndexSource
}

// New wires the service.
func New(cfg *rules.Config, index IndexSource) *Server {
	return &Server{cfg: cfg, index: index}
}

var (
	errNoIndex = errors.New("card database not loaded yet")
	errNoDeck  = errors.New("deck is required")
)

// Validate runs the rules engine on a deck (guardrail 1).
// Ownership context is not wired yet: validation runs pool-rule ANY_CARD
// until the agent (PR-8) passes the session's collection.
func (s *Server) Validate(_ context.Context, req *connect.Request[mtgv1.ValidateRequest]) (*connect.Response[mtgv1.ValidateResponse], error) {
	if req.Msg.Deck == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoDeck)
	}
	idx := s.index.Current()
	if idx == nil {
		return nil, connect.NewError(connect.CodeUnavailable, errNoIndex)
	}
	res := s.cfg.Validate(rules.Input{
		Deck:     req.Msg.Deck,
		PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Cards:    idx,
	})
	return connect.NewResponse(&mtgv1.ValidateResponse{Result: res}), nil
}
