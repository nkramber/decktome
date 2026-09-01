package generate

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// setCard is a pool card with its paper sets.
func setCard(id, name string, codes ...string) *mtgv1.Card {
	return &mtgv1.Card{OracleId: id, Name: name, SetCodes: codes}
}

// TestMarkOutsideSetsMarksTheCardsTheSetsLack is D-383: the build sets
// the mark once, and the deck screen reads it.
func TestMarkOutsideSetsMarksTheCardsTheSetsLack(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{
		setCard("o-in", "Hobbit Hole", "hob"),
		setCard("o-eternal", "Smaug the Impenetrable", "hoc"),
		setCard("o-out", "Sol Ring", "m19", "cmm"),
		// A basic land is never marked: no set limit filters it (D-378).
		{OracleId: "o-plains", Name: "Plains", Supertypes: []string{"Basic"}, CardTypes: []string{"Land"}},
	}, nil)
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-in", Name: "Hobbit Hole", Count: 1},
		{OracleId: "o-eternal", Name: "Smaug the Impenetrable", Count: 1},
		{OracleId: "o-out", Name: "Sol Ring", Count: 1},
		{OracleId: "o-plains", Name: "Plains", Count: 30},
	}}
	req := Request{Pool: pool, SetCodes: []string{"hob", "hoc"}}
	markOutsideSets(deck, req, nil)

	want := map[string]bool{"o-in": false, "o-eternal": false, "o-out": true, "o-plains": false}
	for _, dc := range deck.GetCards() {
		if dc.GetOutsideRequestedSets() != want[dc.GetOracleId()] {
			t.Errorf("%s outside = %v, want %v", dc.GetOracleId(), dc.GetOutsideRequestedSets(), want[dc.GetOracleId()])
		}
	}
	var found *mtgv1.Finding
	for _, f := range deck.GetValidation().GetFindings() {
		if f.GetCode() == CodeOutsideSet {
			found = f
		}
	}
	if found == nil {
		t.Fatal("no outside_requested_set finding")
	}
	if found.GetSeverity() != mtgv1.Severity_SEVERITY_WARN {
		t.Errorf("severity = %v, want WARN: a set limit is a build rule and not a rule of the game", found.GetSeverity())
	}
	if !strings.Contains(found.GetMessage(), "Sol Ring") {
		t.Errorf("the finding does not name the card: %s", found.GetMessage())
	}
}

// TestMarkOutsideSetsDoesNothingWithNoLimit: a deck built with no set
// limit carries no mark and no finding.
func TestMarkOutsideSetsDoesNothingWithNoLimit(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{setCard("o-out", "Sol Ring", "m19")}, nil)
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-out", Name: "Sol Ring", Count: 1}}}
	markOutsideSets(deck, Request{Pool: pool}, nil)
	if deck.GetCards()[0].GetOutsideRequestedSets() {
		t.Error("a deck with no set limit must carry no outside mark")
	}
	if len(deck.GetValidation().GetFindings()) != 0 {
		t.Errorf("findings = %v, want none", deck.GetValidation().GetFindings())
	}
}

// TestMarkOutsideSetsReadsTheSideboard: a 60-card deck carries a
// sideboard, and its cards follow the same rule.
func TestMarkOutsideSetsReadsTheSideboard(t *testing.T) {
	pool := NewPool([]*mtgv1.Card{setCard("o-side", "Rest in Peace", "m19")}, nil)
	deck := &mtgv1.Deck{Sideboard: []*mtgv1.DeckCard{{OracleId: "o-side", Name: "Rest in Peace", Count: 2}}}
	markOutsideSets(deck, Request{Pool: pool, SetCodes: []string{"hob"}}, nil)
	if !deck.GetSideboard()[0].GetOutsideRequestedSets() {
		t.Error("a sideboard card outside the sets must carry the mark")
	}
}

// trimPool holds one commander, one locked card, a mana base of two
// copies, a basic land, and three ordinary cards of different cost.
func trimPool() *Pool {
	return NewPool([]*mtgv1.Card{
		{OracleId: "o-cmd", Name: "Thranduil, the Elvenking", ManaValue: 4},
		{OracleId: "o-lock", Name: "Sol Ring", ManaValue: 1},
		{OracleId: "o-cheap", Name: "Elvish Mystic", ManaValue: 1},
		{OracleId: "o-mid", Name: "Elven Chorus", ManaValue: 4},
		{OracleId: "o-dear", Name: "Bag End Banquet", ManaValue: 7},
		{OracleId: "o-pair", Name: "Mirkwood", ManaValue: 0},
		{OracleId: "o-plains", Name: "Plains", Supertypes: []string{"Basic"}, CardTypes: []string{"Land"}},
	}, nil)
}

// deckOfSize builds a Commander deck whose cards sum to n copies.
func deckOfSize(n int) *mtgv1.Deck {
	d := &mtgv1.Deck{CommanderOracleIds: []string{"o-cmd"}, Cards: []*mtgv1.DeckCard{
		{OracleId: "o-lock", Name: "Sol Ring", Count: 1},
		{OracleId: "o-cheap", Name: "Elvish Mystic", Count: 1},
		{OracleId: "o-mid", Name: "Elven Chorus", Count: 1},
		{OracleId: "o-dear", Name: "Bag End Banquet", Count: 1},
		{OracleId: "o-pair", Name: "Mirkwood", Count: 2},
	}}
	// Basic lands make up the rest, in one entry.
	have := 6
	d.Cards = append(d.Cards, &mtgv1.DeckCard{
		OracleId: "o-plains", Name: "Plains", Count: int32(n - have)})
	return d
}

// TestTrimToSizeCutsTheDearestCard is D-391, the mirror of D-225. A deck
// one or two cards over is blocked whole by the engine, and the reader
// gets nothing. Revise gate run 3 saw the model remove six cards and add
// seven under a mana cap.
func TestTrimToSizeCutsTheDearestCard(t *testing.T) {
	deck := deckOfSize(100) // one card over the 99 the format wants
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: trimPool(),
		Commanders: []string{"o-cmd"}, Locked: []string{"o-lock"}}
	got := trimToSize(deck, req)
	if len(got) != 1 || got[0] != "Bag End Banquet" {
		t.Fatalf("trimmed %v, want the dearest card", got)
	}
	total := len(deck.GetCommanderOracleIds())
	for _, c := range deck.GetCards() {
		total += int(c.GetCount())
	}
	if total != 100 {
		t.Errorf("the deck holds %d cards, want the format size", total)
	}
}

// TestTrimToSizeKeepsWhatTheReaderNamed: a commander, a locked card, a
// basic land, and any entry of more than one copy all survive a trim.
func TestTrimToSizeKeepsWhatTheReaderNamed(t *testing.T) {
	deck := deckOfSize(101) // two over
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: trimPool(),
		Commanders: []string{"o-cmd"}, Locked: []string{"o-lock"}}
	got := trimToSize(deck, req)
	if len(got) != 2 {
		t.Fatalf("trimmed %v, want two cards", got)
	}
	for _, name := range got {
		if name == "Sol Ring" {
			t.Error("the trim cut a locked card (D-242)")
		}
		if name == "Plains" {
			t.Error("the trim cut a basic land, which is the mana base")
		}
		if name == "Mirkwood" {
			t.Error("the trim cut an entry of more than one copy")
		}
	}
}

// TestTrimToSizeLeavesARealFailureAlone is the bound. Three cards over
// is a different deck, and it stays a block finding for the reader to
// see, the way D-225 leaves a large shortfall alone.
func TestTrimToSizeLeavesARealFailureAlone(t *testing.T) {
	for _, n := range []int{99, 103} {
		deck := deckOfSize(n)
		before := len(deck.GetCards())
		req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Pool: trimPool(),
			Commanders: []string{"o-cmd"}, Locked: []string{"o-lock"}}
		if got := trimToSize(deck, req); len(got) != 0 {
			t.Errorf("size %d: trimmed %v, want none", n, got)
		}
		if len(deck.GetCards()) != before {
			t.Errorf("size %d: the trim changed the deck", n)
		}
	}
}

// TestTrimToSizeSkipsAFormatWithNoExactSize: a 60-card format states a
// minimum, so a larger deck is legal and nothing is cut.
func TestTrimToSizeSkipsAFormatWithNoExactSize(t *testing.T) {
	deck := deckOfSize(100)
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_UNSPECIFIED, Pool: trimPool()}
	if got := trimToSize(deck, req); len(got) != 0 {
		t.Errorf("trimmed %v, want none for a format with no exact size", got)
	}
}
