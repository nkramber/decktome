// Package decksvc serves DeckService: Validate, the reads of the decks a
// build kept (D-245), and ExportDeck (D-15).
package decksvc

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/auth"
	"github.com/nkramber/mtg-deck-builder/go/internal/cardsvc"
	"github.com/nkramber/mtg-deck-builder/go/internal/decks"
	"github.com/nkramber/mtg-deck-builder/go/internal/export"
	"github.com/nkramber/mtg-deck-builder/go/internal/gzstore"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
)

// CollectionSource gives the owned count per Oracle id for one collection
// of one user (D-37). A missing collection returns a NotFound status.
type CollectionSource interface {
	OracleCounts(ctx context.Context, userID, collectionID string) (map[string]int32, error)
}

// Option configures the server.
type Option func(*Server)

// DeckSource reads and writes the decks a build kept (D-245). List
// answers from the flat fields and carries no cards. Update writes the
// two fields the user owns, and Delete removes a deck for good (PR-17).
type DeckSource interface {
	Get(ctx context.Context, uid, id string) (*mtgv1.Deck, error)
	List(ctx context.Context, uid string, f decks.Filter, scan int) ([]*mtgv1.Deck, error)
	Update(ctx context.Context, uid, id string, name *string, favorite *bool) (*mtgv1.Deck, error)
	Delete(ctx context.Context, uid, id string) error
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

// maxNameBytes caps a deck name, so one document keeps a sane size.
const maxNameBytes = 200

const (
	// defaultPageSize serves a request that names no size.
	defaultPageSize = 24
	// maxPageSize caps one page, so one response stays a sane size.
	maxPageSize = 100
	// maxScan caps the rows one listing reads. The filter runs in Go, so
	// a listing reads the newest rows and keeps the ones that pass. A
	// user with more decks than this needs a search index (PR-17).
	maxScan = 500
)

// WithCollections wires the ownership check. Without it, Validate refuses
// a request that names a collection.
func WithCollections(src CollectionSource) Option {
	return func(s *Server) { s.collections = src }
}

// SessionSource is the chat store the delete reads. A deck and the chat
// that built it are one thing, so a deck delete takes the chat and every
// deck of the chat with it (D-456).
type SessionSource interface {
	Get(ctx context.Context, uid, id string) (*mtgv1.Session, error)
	Delete(ctx context.Context, uid, id string) error
}

// WithSessions wires the chat store for the delete.
func WithSessions(src SessionSource) Option {
	return func(s *Server) { s.sessions = src }
}

// Server answers DeckService requests.
type Server struct {
	decks DeckSource
	mtgv1connect.UnimplementedDeckServiceHandler
	cfg         *rules.Config
	index       cardsvc.IndexSource
	collections CollectionSource
	sessions    SessionSource
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
	errEmptyName      = errors.New("name: a deck name needs a character that is not a space")
	errLongName       = fmt.Errorf("name: a deck name takes at most %d bytes", maxNameBytes)
	errNoUpdate       = errors.New("give a name or a favorite mark to write")
	errBadBracket     = errors.New("power_bracket: a Commander bracket is 1 to 5, or 0 for every bracket")
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
// deck whole. The filter and the page come from the request (PR-17).
func (s *Server) ListDecks(ctx context.Context, req *connect.Request[mtgv1.ListDecksRequest]) (*connect.Response[mtgv1.ListDecksResponse], error) {
	if s.decks == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errNoDeckStore)
	}
	uid := s.user(ctx)
	if uid == "" {
		return nil, connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	filter := decks.Filter{
		Format:         req.Msg.GetFormat(),
		Favorite:       req.Msg.Favorite,
		Query:          strings.TrimSpace(req.Msg.GetQuery()),
		PowerBracket:   req.Msg.GetPowerBracket(),
		PowerSixtyStep: req.Msg.GetPowerSixtyStep(),
		SessionID:      req.Msg.GetSessionId(),
	}
	if b := filter.PowerBracket; b < 0 || b > 5 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadBracket)
	}
	offset, err := decodePageToken(req.Msg.GetPageToken(), filter)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	size := pageSize(req.Msg.GetPageSize())
	list, err := s.decks.List(ctx, uid, filter, maxScan)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	page, next := slicePage(list, offset, size, filter)
	return connect.NewResponse(&mtgv1.ListDecksResponse{Decks: page, NextPageToken: next}), nil
}

// pageSize applies the default and the cap.
func pageSize(n int32) int {
	switch {
	case n <= 0:
		return defaultPageSize
	case n > maxPageSize:
		return maxPageSize
	default:
		return int(n)
	}
}

// slicePage cuts one page and names the next token. An offset past the
// end gives an empty page and no token.
func slicePage(list []*mtgv1.Deck, offset, size int, f decks.Filter) ([]*mtgv1.Deck, string) {
	if offset >= len(list) {
		return nil, ""
	}
	end := offset + size
	if end >= len(list) {
		return list[offset:], ""
	}
	return list[offset:end], encodePageToken(end, f)
}

// The page token carries the offset and a fingerprint of the filter. A
// token of another filter is an invalid argument, because its offset
// counts a different list.
var errBadPageToken = errors.New("page_token: this token belongs to another filter or another listing")

func filterFingerprint(f decks.Filter) string {
	fav := "unset"
	if f.Favorite != nil {
		fav = strconv.FormatBool(*f.Favorite)
	}
	sum := sha256.Sum256([]byte(fmt.Sprintf("%d\x00%s\x00%s\x00%d\x00%d\x00%s", f.Format, fav, strings.ToLower(f.Query), f.PowerBracket, f.PowerSixtyStep, f.SessionID)))
	return hex.EncodeToString(sum[:6])
}

func encodePageToken(offset int, f decks.Filter) string {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%d.%s", offset, filterFingerprint(f))))
}

func decodePageToken(token string, f decks.Filter) (int, error) {
	if token == "" {
		return 0, nil
	}
	raw, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return 0, errBadPageToken
	}
	offsetText, fingerprint, ok := strings.Cut(string(raw), ".")
	if !ok || fingerprint != filterFingerprint(f) {
		return 0, errBadPageToken
	}
	offset, err := strconv.Atoi(offsetText)
	if err != nil || offset < 0 {
		return 0, errBadPageToken
	}
	return offset, nil
}

// UpdateDeck writes the name and the favorite mark (PR-17). An unset
// field stays as it is, and an empty name is an invalid argument.
func (s *Server) UpdateDeck(ctx context.Context, req *connect.Request[mtgv1.UpdateDeckRequest]) (*connect.Response[mtgv1.UpdateDeckResponse], error) {
	uid, id, err := s.deckRef(ctx, req.Msg.GetDeckId())
	if err != nil {
		return nil, err
	}
	name := req.Msg.Name
	if name != nil {
		trimmed := strings.TrimSpace(*name)
		if trimmed == "" {
			return nil, connect.NewError(connect.CodeInvalidArgument, errEmptyName)
		}
		if len(trimmed) > maxNameBytes {
			return nil, connect.NewError(connect.CodeInvalidArgument, errLongName)
		}
		name = &trimmed
	}
	if name == nil && req.Msg.Favorite == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoUpdate)
	}
	d, err := s.decks.Update(ctx, uid, id, name, req.Msg.Favorite)
	if errors.Is(err, decks.ErrNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&mtgv1.UpdateDeckResponse{Deck: d}), nil
}

// DeleteDeck removes one deck for good (PR-17), and the chat that built
// it with every deck of that chat (D-456). A second delete of the same
// id answers NotFound. The chat goes first: a deck that outlives a
// failed chat delete is visible and can be deleted again, and a chat
// that outlives its decks is reachable from nowhere.
func (s *Server) DeleteDeck(ctx context.Context, req *connect.Request[mtgv1.DeleteDeckRequest]) (*connect.Response[mtgv1.DeleteDeckResponse], error) {
	uid, id, err := s.deckRef(ctx, req.Msg.GetDeckId())
	if err != nil {
		return nil, err
	}
	deck, err := s.decks.Get(ctx, uid, id)
	if errors.Is(err, decks.ErrNotFound) {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	ids := []string{id}
	if s.sessions != nil && deck.GetSessionId() != "" {
		siblings, err := s.deleteChat(ctx, uid, deck.GetSessionId())
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		ids = append(ids, siblings...)
	}
	seen := map[string]bool{}
	for _, did := range ids {
		if seen[did] || !gzstore.ValidID(did) {
			continue
		}
		seen[did] = true
		err := s.decks.Delete(ctx, uid, did)
		// A sibling that is already gone is no failure.
		if err != nil && !errors.Is(err, decks.ErrNotFound) {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}
	return connect.NewResponse(&mtgv1.DeleteDeckResponse{}), nil
}

// deleteChat removes the chat and returns the ids of its decks. A chat
// that is already gone returns nothing and no error.
func (s *Server) deleteChat(ctx context.Context, uid, sessionID string) ([]string, error) {
	if !gzstore.ValidID(sessionID) {
		return nil, nil
	}
	session, err := s.sessions.Get(ctx, uid, sessionID)
	if errors.Is(err, sessions.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if err := s.sessions.Delete(ctx, uid, sessionID); err != nil && !errors.Is(err, sessions.ErrNotFound) {
		return nil, err
	}
	return session.GetDeckIds(), nil
}

// deckRef checks the store, the id, and the caller. Every write of one
// deck starts here.
func (s *Server) deckRef(ctx context.Context, rawID string) (uid, id string, err error) {
	if s.decks == nil {
		return "", "", connect.NewError(connect.CodeUnimplemented, errNoDeckStore)
	}
	id = strings.TrimSpace(rawID)
	if id == "" {
		return "", "", connect.NewError(connect.CodeInvalidArgument, errNoDeckID)
	}
	if !gzstore.ValidID(id) {
		return "", "", connect.NewError(connect.CodeInvalidArgument, errBadDeckID)
	}
	uid = s.user(ctx)
	if uid == "" {
		return "", "", connect.NewError(connect.CodeUnauthenticated, errNoUser)
	}
	return uid, id, nil
}

// ExportDeck renders one of the caller's decks as text (D-15). The card
// index gives the names and the printings. Without an index the text
// falls back to the names the deck stored, so an export never waits on
// a snapshot load.
func (s *Server) ExportDeck(ctx context.Context, req *connect.Request[mtgv1.ExportDeckRequest]) (*connect.Response[mtgv1.ExportDeckResponse], error) {
	deck, err := s.GetDeck(ctx, connect.NewRequest(&mtgv1.GetDeckRequest{DeckId: req.Msg.GetDeckId()}))
	if err != nil {
		return nil, err
	}
	var lookup export.Lookup = noCards{}
	if idx := s.index.Current(); idx != nil {
		lookup = idx
	}
	text, name, err := export.Render(deck.Msg.GetDeck(), lookup, req.Msg.GetFormat())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	return connect.NewResponse(&mtgv1.ExportDeckResponse{Text: text, FileName: name}), nil
}

type noCards struct{}

func (noCards) ByOracleID(string) (*mtgv1.Card, bool) { return nil, false }

// user reads the caller's id, or empty when no source is wired.
func (s *Server) user(ctx context.Context) string {
	if s.userFn == nil {
		return ""
	}
	return s.userFn(ctx)
}
