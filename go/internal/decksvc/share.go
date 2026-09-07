package decksvc

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"regexp"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/export"
)

// The share link (D-315). A token is 32 random bytes in URL-safe base64,
// 43 characters, shown once. The store keeps the hex SHA-256 of it, so a
// stolen store gives no link. A public read hashes the token it got and
// looks the hash up.

// tokenBytes is the size of a token before encoding.
const tokenBytes = 32

var tokenShape = regexp.MustCompile(`^[A-Za-z0-9_-]{43}$`)

// newToken makes a token and its hash.
func newToken() (token, hash string, err error) {
	raw := make([]byte, tokenBytes)
	if _, err := rand.Read(raw); err != nil {
		return "", "", err
	}
	token = base64.RawURLEncoding.EncodeToString(raw)
	return token, hashToken(token), nil
}

// hashToken is the store key of a token.
func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

var errNoToken = errors.New("no share token was given")

// ShareDeck makes a share link for one of the caller's decks and answers
// the token once (D-315). A deck with a link gets a new one.
func (s *Server) ShareDeck(ctx context.Context, req *connect.Request[mtgv1.ShareDeckRequest]) (*connect.Response[mtgv1.ShareDeckResponse], error) {
	uid, id, err := s.deckRef(ctx, req.Msg.GetDeckId())
	if err != nil {
		return nil, err
	}
	token, hash, err := newToken()
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if err := s.decks.Share(ctx, uid, id, hash); err != nil {
		return nil, storeError(err)
	}
	return connect.NewResponse(&mtgv1.ShareDeckResponse{Token: token}), nil
}

// RevokeShare ends the link of one of the caller's decks.
func (s *Server) RevokeShare(ctx context.Context, req *connect.Request[mtgv1.RevokeShareRequest]) (*connect.Response[mtgv1.RevokeShareResponse], error) {
	uid, id, err := s.deckRef(ctx, req.Msg.GetDeckId())
	if err != nil {
		return nil, err
	}
	if err := s.decks.Revoke(ctx, uid, id); err != nil {
		return nil, storeError(err)
	}
	return connect.NewResponse(&mtgv1.RevokeShareResponse{}), nil
}

// sharedRead answers the deck a token opens. A token of the wrong shape,
// an unknown one, a revoked one, and a deleted deck all read NotFound,
// so the answer says nothing about which.
func (s *Server) sharedRead(ctx context.Context, token string) (*mtgv1.Deck, error) {
	if s.decks == nil {
		return nil, connect.NewError(connect.CodeUnimplemented, errNoDeckStore)
	}
	if token == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoToken)
	}
	if !tokenShape.MatchString(token) {
		return nil, connect.NewError(connect.CodeNotFound, decks.ErrNotFound)
	}
	uid, id, err := s.decks.LookupShare(ctx, hashToken(token))
	if err != nil {
		return nil, storeError(err)
	}
	d, err := s.decks.Get(ctx, uid, id)
	if err != nil {
		return nil, storeError(err)
	}
	return d, nil
}

// GetSharedDeck reads a shared deck by its token, with no sign-in and no
// user field in the answer (D-315, guardrail 13).
func (s *Server) GetSharedDeck(ctx context.Context, req *connect.Request[mtgv1.GetSharedDeckRequest]) (*connect.Response[mtgv1.GetSharedDeckResponse], error) {
	d, err := s.sharedRead(ctx, req.Msg.GetToken())
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&mtgv1.GetSharedDeckResponse{Deck: sharedDeck(d, s.lookup())}), nil
}

// ExportSharedDeck renders a shared deck as Arena text with the default
// paper printings: the public copy holds no owned printing (D-315).
func (s *Server) ExportSharedDeck(ctx context.Context, req *connect.Request[mtgv1.ExportSharedDeckRequest]) (*connect.Response[mtgv1.ExportSharedDeckResponse], error) {
	d, err := s.sharedRead(ctx, req.Msg.GetToken())
	if err != nil {
		return nil, err
	}
	text, name, err := export.Render(publicDeck(d), s.lookup(), mtgv1.ExportFormat_EXPORT_FORMAT_ARENA_TEXT)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	return connect.NewResponse(&mtgv1.ExportSharedDeckResponse{Text: text, FileName: name}), nil
}

// lookup is the card index, or nothing before a snapshot loads.
func (s *Server) lookup() export.Lookup {
	if s.index != nil {
		if idx := s.index.Current(); idx != nil {
			return idx
		}
	}
	return noCards{}
}

// storeError maps a store error onto a Connect code.
func storeError(err error) error {
	if errors.Is(err, decks.ErrNotFound) {
		return connect.NewError(connect.CodeNotFound, err)
	}
	return connect.NewError(connect.CodeInternal, err)
}

// sharedDeck is the public copy of a deck: the name, the format, the
// power, the summary, and the cards by role with their card data, and no
// user field (guardrail 13).
func sharedDeck(d *mtgv1.Deck, cards export.Lookup) *mtgv1.SharedDeck {
	return &mtgv1.SharedDeck{
		Name:               d.GetName(),
		Format:             d.GetFormat(),
		Power:              d.GetPower(),
		Summary:            d.GetSummary(),
		CommanderOracleIds: d.GetCommanderOracleIds(),
		Cards:              sharedCards(d.GetCards(), cards),
		Sideboard:          sharedCards(d.GetSideboard(), cards),
		LegalityAsOf:       d.GetLegalityAsOf(),
		CardCount:          d.GetCardCount(),
	}
}

func sharedCards(list []*mtgv1.DeckCard, cards export.Lookup) []*mtgv1.SharedCard {
	out := make([]*mtgv1.SharedCard, 0, len(list))
	for _, dc := range list {
		sc := &mtgv1.SharedCard{OracleId: dc.GetOracleId(), Name: dc.GetName(), Count: dc.GetCount(), Role: dc.GetRole(), Reason: dc.GetReason()}
		if c, ok := cards.ByOracleID(dc.GetOracleId()); ok && c != nil {
			// The index card is shared, so the copy is a clone.
			sc.Card = proto.Clone(c).(*mtgv1.Card) //nolint:errcheck,forcetypeassert // Clone of a Card is a Card
		}
		out = append(out, sc)
	}
	return out
}

// publicDeck is a copy of a deck with every user field cleared, for the
// public export: no session, no owned mark, no owned printing.
func publicDeck(d *mtgv1.Deck) *mtgv1.Deck {
	out := proto.Clone(d).(*mtgv1.Deck) //nolint:errcheck,forcetypeassert // Clone of a Deck is a Deck
	out.SessionId = ""
	out.Favorite = false
	out.RevisedFromDeckId = ""
	out.Validation = nil
	for _, list := range [][]*mtgv1.DeckCard{out.Cards, out.Sideboard, out.Upgrades} {
		for _, dc := range list {
			dc.Owned = false
			dc.OwnedCount = 0
			dc.OwnedPrinting = nil
		}
	}
	return out
}
