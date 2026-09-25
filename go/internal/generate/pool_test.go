package generate

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestAColorlessIdentityGetsWastes is REV-017 of the review of
// 2026-09-24. A colorless commander got no basic land, so no mana base
// could fill the deck.
func TestAColorlessIdentityGetsWastes(t *testing.T) {
	cards := map[string]*mtgv1.Card{
		"Wastes": {OracleId: "o-wastes", Name: "Wastes"},
		"Plains": {OracleId: "o-plains", Name: "Plains"},
	}
	find := func(name string) (*mtgv1.Card, bool) { c, ok := cards[name]; return c, ok }
	if got := BasicLands(find, nil); len(got) != 1 || got[0].GetName() != "Wastes" {
		t.Errorf("colorless basics = %v, want Wastes", got)
	}
	if got := BasicLands(find, []mtgv1.Color{mtgv1.Color_COLOR_W}); len(got) != 1 || got[0].GetName() != "Plains" {
		t.Errorf("white basics = %v, want Plains", got)
	}
}
