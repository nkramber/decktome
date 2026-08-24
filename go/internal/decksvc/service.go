// Package decksvc serves DeckService. PR-5 wires Validate. Get, List,
// and Export arrive with PR-8 and PR-13.
package decksvc

import (
	"context"
	"errors"
	"fmt"

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

// CollectionSource gives the owned count per Oracle id for one collection
// of one user (D-37). A missing collection returns an error.
type CollectionSource interface {
	OracleCounts(ctx context.Context, userID, collectionID string) (map[string]int32, error)
}

// UserFunc reads the caller's user id from the request context.
type UserFunc func(ctx context.Context) string

// Option configures the server.
type Option func(*Server)

// WithCollections wires the ownership check. Without it, Validate refuses
// a request that names a collection.
func WithCollections(src CollectionSource, userFn UserFunc) Option {
	return func(s *Server) {
		s.collections = src
		s.userFn = userFn
	}
}

// Server answers DeckService requests.
type Server struct {
	mtgv1connect.UnimplementedDeckServiceHandler
	cfg         *rules.Config
	index       IndexSource
	collections CollectionSource
	userFn      UserFunc
}

// New wires the service.
func New(cfg *rules.Config, index IndexSource, opts ...Option) *Server {
	s := &Server{cfg: cfg, index: index}
	for _, o := range opts {
		o(s)
	}
	return s
}

var (
	errNoIndex        = errors.New("card database not loaded yet")
	errNoDeck         = errors.New("deck is required")
	errNoCollectionID = errors.New("an owned pool rule needs collection_id")
	errNoCollections  = errors.New("collections are not wired on this server")
)

// Validate runs the rules engine on a deck (guardrail 1).
//
// Pool rule default (deck_service.proto): UNSPECIFIED means ANY_CARD when
// collection_id is empty, else OWNED_FIRST. An owned rule without a
// collection_id is an invalid argument.
func (s *Server) Validate(ctx context.Context, req *connect.Request[mtgv1.ValidateRequest]) (*connect.Response[mtgv1.ValidateResponse], error) {
	if req.Msg.Deck == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoDeck)
	}
	idx := s.index.Current()
	if idx == nil {
		return nil, connect.NewError(connect.CodeUnavailable, errNoIndex)
	}
	pool := resolvePoolRule(req.Msg.PoolRule, req.Msg.CollectionId)
	var counts map[string]int32
	if req.Msg.CollectionId != "" {
		if s.collections == nil {
			return nil, connect.NewError(connect.CodeUnavailable, errNoCollections)
		}
		var userID string
		if s.userFn != nil {
			userID = s.userFn(ctx)
		}
		var err error
		counts, err = s.collections.OracleCounts(ctx, userID, req.Msg.CollectionId)
		if err != nil {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("collection %q: %w", req.Msg.CollectionId, err))
		}
	} else if pool != mtgv1.PoolRule_POOL_RULE_ANY_CARD {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoCollectionID)
	}
	res := s.cfg.Validate(rules.Input{
		Deck:         req.Msg.Deck,
		PoolRule:     pool,
		OracleCounts: counts,
		Cards:        idx,
	})
	if !idx.AsOf.IsZero() {
		res.LegalityAsOf = idx.AsOf.UTC().Format("2006-01-02")
	}
	return connect.NewResponse(&mtgv1.ValidateResponse{Result: res}), nil
}

// resolvePoolRule applies the D-37 default.
func resolvePoolRule(pr mtgv1.PoolRule, collectionID string) mtgv1.PoolRule {
	if pr != mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		return pr
	}
	if collectionID == "" {
		return mtgv1.PoolRule_POOL_RULE_ANY_CARD
	}
	return mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
}
