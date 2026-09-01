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
