package generate

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
)

// TestThemeLeftOutNamesTheSwap is D-1198. Replay v3-14 played Yargle,
// Glutton of Urborg, a card of the set fill, and left out Chandra, Torch
// of Defiance on theme. The finding names both, and it buys the repair
// turn. A land on theme, a revision, and an upgrade read no finding.
func TestThemeLeftOutNamesTheSwap(t *testing.T) {
	card := func(id, name, typeLine string) *mtgv1.Card {
		return &mtgv1.Card{OracleId: id, Name: name, TypeLine: typeLine}
	}
	l := &candidates.List{Candidates: []candidates.Candidate{
		{Card: card("o-chandra", "Chandra, Torch of Defiance", "Legendary Planeswalker — Chandra"), Role: mtgv1.CardRole_CARD_ROLE_THREAT, Themed: true, OnTheme: true},
		{Card: card("o-sanctum", "Theorist's Sanctum", "Land — Island"), Role: mtgv1.CardRole_CARD_ROLE_LAND, Themed: true, OnTheme: true},
		{Card: card("o-yargle", "Yargle, Glutton of Urborg", "Legendary Creature — Frog Spirit"), Role: mtgv1.CardRole_CARD_ROLE_OTHER, SetFill: true},
		{Card: card("o-sol", "Sol Ring", "Artifact"), Role: mtgv1.CardRole_CARD_ROLE_RAMP},
	}}
	req := testRequest()
	req.Pool = FromList(l, nil, true)
	req.Themed = Themed(l)
	req.SetFill = SetFill(l)
	deckOf := func(ids ...string) *mtgv1.Deck {
		d := &mtgv1.Deck{Validation: &mtgv1.ValidationResult{Passed: true}}
		for _, c := range l.Candidates {
			for _, id := range ids {
				if c.Card.GetOracleId() == id {
					d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: id, Name: c.Card.GetName(), Count: 1})
				}
			}
		}
		return d
	}

	d := deckOf("o-yargle", "o-sol")
	checkThemeLeftOut(d, req)
	got := repairable(d.GetValidation())
	if len(got) != 1 || got[0].GetCode() != CodeThemeLeftOut {
		t.Fatalf("findings %v, want one %s that buys the repair", d.GetValidation().GetFindings(), CodeThemeLeftOut)
	}
	msg := got[0].GetMessage()
	if !strings.Contains(msg, "Chandra, Torch of Defiance") || !strings.Contains(msg, "Yargle, Glutton of Urborg") || strings.Contains(msg, "Theorist's Sanctum") {
		t.Errorf("message %q must name Chandra and Yargle, and not the land", msg)
	}

	for name, tc := range map[string]struct {
		deck *mtgv1.Deck
		edit func(*Request)
	}{
		"no set fill in the deck":  {deckOf("o-sol"), nil},
		"every card on theme held": {deckOf("o-chandra", "o-yargle"), nil},
		"a revision":               {deckOf("o-yargle"), func(r *Request) { r.Revision = &Revision{} }},
		"an upgrade":               {deckOf("o-yargle"), func(r *Request) { r.Precon = "Test Precon" }},
	} {
		r := req
		if tc.edit != nil {
			tc.edit(&r)
		}
		checkThemeLeftOut(tc.deck, r)
		if n := len(tc.deck.GetValidation().GetFindings()); n != 0 {
			t.Errorf("%s: %d findings, want none", name, n)
		}
	}
	if !strings.Contains(repairInstructions, "leaves out cards on theme") {
		t.Error("the repair prompt does not say how to fix the finding")
	}
}
