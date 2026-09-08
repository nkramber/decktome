package revise

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestKeepNamesHoldsACardTheDeckLacks is F-80, found on the deployed app
// on 2026-09-08. The reader wrote "Include The Arkenstone from my
// collection", and the revision changed nothing. Session
// IDS5oQE3D0XDWs6o1FnR.
//
// One filter served both lists, and it dropped every name the deck did
// not hold. A card the reader asks to add is never in the deck, which is
// the whole reason they asked, so the brief could not carry one.
func TestKeepNamesHoldsACardTheDeckLacks(t *testing.T) {
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-sol", Name: "Sol Ring", Count: 1},
		{OracleId: "o-plains", Name: "Plains", Count: 10},
	}}
	got := keepNames([]string{"The Arkenstone", "sol ring", "Sol Ring", "  "}, deck)
	if len(got) != 2 {
		t.Fatalf("keep = %v, want the added card and the deck card once", got)
	}
	// The card the deck lacks stays in the reader's words.
	if got[0] != "The Arkenstone" {
		t.Errorf("keep[0] = %q, want the card the reader asked to add", got[0])
	}
	// A card the deck holds takes the deck's own spelling.
	if got[1] != "Sol Ring" {
		t.Errorf("keep[1] = %q, want the deck's spelling", got[1])
	}
}

// A name to remove still has to be in the deck. Nothing removes a card
// the deck never held, and the model can invent a name.
func TestRemoveStillReadsTheDeckAlone(t *testing.T) {
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-sol", Name: "Sol Ring", Count: 1}}}
	if got := onlyInDeck([]string{"The Arkenstone", "Sol Ring"}, deck); len(got) != 1 || got[0] != "Sol Ring" {
		t.Errorf("remove = %v, want the deck card alone", got)
	}
}

// TestDiffEmptyReadsAnUnchangedDeck is F-81. A revision that changes no
// card stored a new version, and the compare screen then read "the two
// decks hold the same cards" beside a v2.
func TestDiffEmptyReadsAnUnchangedDeck(t *testing.T) {
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-sol", Name: "Sol Ring", Count: 1},
		{OracleId: "o-plains", Name: "Plains", Count: 10},
	}}
	same := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-plains", Name: "Plains", Count: 10},
		{OracleId: "o-sol", Name: "Sol Ring", Count: 1},
	}}
	if !DiffDecks(deck, same).Empty() {
		t.Error("two decks of the same cards read as a change")
	}
	moved := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-sol", Name: "Sol Ring", Count: 1},
		{OracleId: "o-plains", Name: "Plains", Count: 11},
	}}
	if DiffDecks(deck, moved).Empty() {
		t.Error("a deck with one more land read as no change")
	}
	added := &mtgv1.Deck{Cards: append(append([]*mtgv1.DeckCard(nil), deck.GetCards()...),
		&mtgv1.DeckCard{OracleId: "o-ark", Name: "The Arkenstone", Count: 1})}
	if DiffDecks(deck, added).Empty() {
		t.Error("a deck with one more card read as no change")
	}
}
