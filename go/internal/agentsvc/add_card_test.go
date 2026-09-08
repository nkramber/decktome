package agentsvc

import (
	"io"
	"log/slog"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// The F-80 and D-614 tests of the service. F-80: the reader asked to
// include The Arkenstone and the revision changed nothing. D-614: a
// reader who limited the deck to sets reads a printing of those sets.

func keepIndex() *cards.Index {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	return cards.NewIndex([]*mtgv1.Card{
		{OracleId: "o-ark", Name: "The Arkenstone", TypeLine: "Legendary Artifact", Legalities: legal},
		{OracleId: "o-sol", Name: "Sol Ring", TypeLine: "Artifact", Legalities: legal},
		{OracleId: "o-plains", Name: "Plains", TypeLine: "Basic Land — Plains",
			CardTypes: []string{"Land"}, Supertypes: []string{"Basic"}, Legalities: legal},
	}, nil, nil, time.Unix(1000, 0).UTC())
}

func TestKeepableAddsACardTheReaderOwns(t *testing.T) {
	s := &Server{index: fixedIndex{keepIndex()}}
	base := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-sol", Name: "Sol Ring", Count: 1}}}
	owned := map[string]int32{"o-ark": 1, "o-sol": 1}

	keep, refused := s.keepable(base, []string{"The Arkenstone"}, owned, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
	if len(keep) != 1 || keep[0] != "The Arkenstone" {
		t.Errorf("keep = %v, want the card the reader owns", keep)
	}
	if len(refused) != 0 {
		t.Errorf("refused %v, and the reader owns the card", refused)
	}
}

// A name no card carries is refused with a word. The engine blocks a
// deck that lacks a kept card, and no repair turn writes a card that
// does not exist.
func TestKeepableRefusesANameNoCardCarries(t *testing.T) {
	s := &Server{index: fixedIndex{keepIndex()}}
	base := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-sol", Name: "Sol Ring", Count: 1}}}

	keep, refused := s.keepable(base, []string{"The Silmaril"}, nil, mtgv1.PoolRule_POOL_RULE_ANY_CARD)
	if len(keep) != 0 {
		t.Errorf("keep = %v, want nothing", keep)
	}
	if len(refused) != 1 {
		t.Fatalf("refused = %v, want one line", refused)
	}
}

// A reader who asked for their own cards alone gets no card they do not
// own, whatever they name. That promise is what F-76 was about.
func TestKeepableRefusesAnUnownedCardUnderOwnedOnly(t *testing.T) {
	s := &Server{index: fixedIndex{keepIndex()}}
	base := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-sol", Name: "Sol Ring", Count: 1}}}

	keep, refused := s.keepable(base, []string{"The Arkenstone"}, map[string]int32{}, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
	if len(keep) != 0 {
		t.Errorf("keep = %v, and the reader owns no copy", keep)
	}
	if len(refused) != 1 {
		t.Fatalf("refused = %v, want one line", refused)
	}
	// The same card passes when the pool buys cards.
	keep, refused = s.keepable(base, []string{"The Arkenstone"}, map[string]int32{}, mtgv1.PoolRule_POOL_RULE_OWNED_FIRST)
	if len(keep) != 1 || len(refused) != 0 {
		t.Errorf("owned-first keep = %v, refused = %v", keep, refused)
	}
}

// A card already in the deck always passes: it is there to keep.
func TestKeepableKeepsACardOfTheDeck(t *testing.T) {
	s := &Server{index: fixedIndex{keepIndex()}}
	base := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-sol", Name: "Sol Ring", Count: 1}}}

	keep, refused := s.keepable(base, []string{"Sol Ring"}, map[string]int32{}, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
	if len(keep) != 1 || len(refused) != 0 {
		t.Errorf("keep = %v, refused = %v, and the deck holds the card", keep, refused)
	}
}

// A basic land is always available, also under owned-only (D-37).
func TestKeepableTakesABasicLandUnderOwnedOnly(t *testing.T) {
	s := &Server{index: fixedIndex{keepIndex()}}
	base := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-sol", Name: "Sol Ring", Count: 1}}}

	keep, refused := s.keepable(base, []string{"Plains"}, map[string]int32{}, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
	if len(keep) != 1 || len(refused) != 0 {
		t.Errorf("keep = %v, refused = %v, and a basic land is always available", keep, refused)
	}
}

// TestOwnedPrintingPrefersTheNamedSets is D-614. The reader limited the
// deck to the Hobbit and Lord of the Rings sets, and they own Sol Ring
// in a Secret Lair printing worth more than the Lord of the Rings one.
// The deck shows the printing of the sets they asked for.
func TestOwnedPrintingPrefersTheNamedSets(t *testing.T) {
	art := &mtgv1.ImageUris{Normal: "https://example.test/art.jpg"}
	idx := cards.NewIndex(
		[]*mtgv1.Card{{OracleId: "o-sol", Name: "Sol Ring", TypeLine: "Artifact"}},
		[]cards.Printing{
			{ScryfallID: "p-sld", OracleID: "o-sol", Name: "Sol Ring", SetCode: "SLD",
				CollectorNumber: "1", PriceUSD: 90, ImageUris: art},
			{ScryfallID: "p-ltr", OracleID: "o-sol", Name: "Sol Ring", SetCode: "LTR",
				CollectorNumber: "2", PriceUSD: 4, ImageUris: art},
			{ScryfallID: "p-hob", OracleID: "o-sol", Name: "Sol Ring", SetCode: "HOB",
				CollectorNumber: "3", PriceUSD: 7, ImageUris: art},
		},
		nil, time.Unix(1000, 0).UTC())
	coll := fakeCollections{printings: map[string][]string{"o-sol": {"p-sld", "p-ltr", "p-hob"}}}
	s := &Server{index: fixedIndex{idx}, collections: coll, log: slog.New(slog.NewTextHandler(io.Discard, nil))}
	session := &mtgv1.Session{Id: "s-1", CollectionId: "c-1"}
	deck := func() *mtgv1.Deck {
		return &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
			{OracleId: "o-sol", Name: "Sol Ring", Count: 1, OwnedCount: 3, Owned: true},
		}}
	}

	// With a set limit the priciest printing of those sets wins.
	limited := deck()
	s.markOwnedPrintings(t.Context(), "u-1", session, idx, limited, []string{"hob", "hoc", "ltr", "ltc"})
	if got := limited.GetCards()[0].GetOwnedPrinting().GetSetCode(); got != "HOB" {
		t.Errorf("the printing is from %q, want the dearest of the named sets", got)
	}

	// With no set limit the priciest printing of the collection wins.
	open := deck()
	s.markOwnedPrintings(t.Context(), "u-1", session, idx, open, nil)
	if got := open.GetCards()[0].GetOwnedPrinting().GetSetCode(); got != "SLD" {
		t.Errorf("the printing is from %q, want the dearest of the collection", got)
	}

	// A reader who owns no copy from the named sets keeps the dearest of
	// the collection: an empty art is worse than an art of another set.
	other := deck()
	s.markOwnedPrintings(t.Context(), "u-1", session, idx, other, []string{"blb"})
	if got := other.GetCards()[0].GetOwnedPrinting().GetSetCode(); got != "SLD" {
		t.Errorf("the printing is from %q, want the dearest of the collection", got)
	}
}
