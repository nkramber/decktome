// Package decksvc serves DeckService: Validate, and the reads of the
// decks a build kept (D-245). Export is a later step.
package decksvc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/auth"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/decks"
	"github.com/nkramber/mtg-deck-builder/go/internal/gzstore"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// CollectionSource gives the owned count per Oracle id for one collection
// of one user (D-37). A missing collection returns a NotFound status.
type CollectionSource interface {
	OracleCounts(ctx context.Context, userID, collectionID string) (map[string]int32, error)
}

// Option configures the server.
type Option func(*Server)

// DeckSource reads the decks a build kept (D-245). List answers from the
// flat fields and carries no cards.
type DeckSource interface {
	Get(ctx context.Context, uid, id string) (*mtgv1.Deck, error)
	List(ctx context.Context, uid string, limit int) ([]*mtgv1.Deck, error)
}

// WithDecks wires the deck store. Without it GetDeck and ListDecks
// answer Unimplemented.
func WithDecks(src DeckSource) Option {
	return func(s *Server) { s.decks = src }
}

// WithUser wires the caller's identity. Without it every read that needs
// a user answers Unauthenticated, and Validate passes an empty user to
// the collection source.
func WithUser(userFn auth.UserFunc) Option {
	return func(s *Server) { s.userFn = userFn }
}

// listLimit caps one ListDecks answer. The request carries no paging
// field, so the cap keeps one response inside a sane size.
const listLimit = 100

// WithCollections wires the ownership check. Without it, Validate refuses
// a request that names a collection.
func WithCollections(src CollectionSource) Option {
	return func(s *Server) { s.collections = src }
}

// Server answers DeckService requests.
type Server struct {
	decks DeckSource
	mtgv1connect.UnimplementedDeckServiceHandler
	cfg         *rules.Config
	index       cardsvc.IndexSource
	collections CollectionSource
	userFn      auth.UserFunc
}

// New wires the service.
func New(cfg *rules.Config, index cardsvc.IndexSource, opts ...Option) *Server {
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
	errNoDeckStore    = errors.New("no deck store is wired")
	errNoDeckID       = errors.New("give a deck id")
	errNoUser         = errors.New("no user in the request context")
	errBadDeckID      = fmt.Errorf("deck_id: %w", gzstore.ErrBadID)
	errBadCollection  = fmt.Errorf("collection_id: %w", gzstore.ErrBadID)
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
		if !gzstore.ValidID(req.Msg.CollectionId) {
			return nil, connect.NewError(connect.CodeInvalidArgument, errBadCollection)
		}
		if s.collections == nil {
			return nil, connect.NewError(connect.CodeUnavailable, errNoCollections)
		}
		var err error
		counts, err = s.collections.OracleCounts(ctx, s.user(ctx), req.Msg.CollectionId)
		if err != nil {
			// Only a missing document is NotFound. A store failure is
			// Internal, so a client does not read an outage as a bad id.
			if status.Code(err) == codes.NotFound {
				return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("collection %q: %w", req.Msg.CollectionId, err))
			}
			return nil, connect.NewError(connect.CodeInternal, err)
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

// GetDeck reads one of the caller's decks (D-245).
func (s *Server) GetDeck(ctx context.Context, req *connect.Request[mtgv1.GetDeckRequest]) (*connect.Response[mtgv1.GetDeckResponse], error) {
	if s.decks == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errNoDeckStore)
	}
	id := strings.TrimSpace(req.Msg.GetDeckId())
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoDeckID)
	}
	if !gzstore.ValidID(id) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadDeckID)
	}
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	d, err := s.decks.Get(ctx, uid, id)
	if errors.Is(err, decks.ErrNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.GetDeckResponse{Deck: d}), nil
}

// ListDecks reads the caller's decks, newest first (D-245). The list
// carries the flat fields of each deck and no cards: GetDeck reads one
// deck whole.
func (s *Server) ListDecks(ctx context.Context, _ *connect.Request[mtgv1.ListDecksRequest]) (*connect.Response[mtgv1.ListDecksResponse], error) {
	if s.decks == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errNoDeckStore)
	}
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	list, err := s.decks.List(ctx, uid, listLimit)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.ListDecksResponse{Decks: list}), nil
}

// user reads the caller's id, or empty when no source is wired.
func (s *Server) user(ctx context.Context) string {
	if s.userFn == nil {
		return ""
	}
	return s.userFn(ctx)
}
