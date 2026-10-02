package candidates

import (
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// TestOwnedInSetsKeepsTheBasicLands is D-1041 with D-37. A basic land
// never leaves the pool, so a Forest from any set stays owned. A copy of
// a set the reader did not name leaves, and a missing printing counts
// for nothing.
func TestOwnedInSetsKeepsTheBasicLands(t *testing.T) {
	idx := cards.NewIndex(
		[]*mtgv1.Card{
			{OracleId: "o-forest", Name: "Forest", Supertypes: []string{"Basic"}, CardTypes: []string{"Land"}},
			{OracleId: "o-ring", Name: "The One Ring", CardTypes: []string{"Artifact"}},
			{OracleId: "o-gq", Name: "Ghost Quarter", CardTypes: []string{"Land"}},
		},
		[]cards.Printing{
			{ScryfallID: "p-forest", OracleID: "o-forest", Name: "Forest", SetCode: "M21", CollectorNumber: "1"},
			{ScryfallID: "p-ring", OracleID: "o-ring", Name: "The One Ring", SetCode: "LTR", CollectorNumber: "2"},
			{ScryfallID: "p-gq", OracleID: "o-gq", Name: "Ghost Quarter", SetCode: "SLD", CollectorNumber: "3"},
		},
		nil, time.Unix(1000, 0).UTC())
	owned := map[string]int32{"o-forest": 20, "o-ring": 1, "o-gq": 4}
	printings := map[string]int32{"p-forest": 20, "p-ring": 1, "p-gq": 4, "p-gone": 2}

	got := OwnedInSets(idx, owned, printings, []string{"LTR", "ltc"})
	want := map[string]int32{"o-forest": 20, "o-ring": 1}
	if len(got) != len(want) || got["o-forest"] != 20 || got["o-ring"] != 1 {
		t.Errorf("owned in the named sets = %v, want %v", got, want)
	}
	if same := OwnedInSets(idx, owned, printings, nil); len(same) != 3 {
		t.Errorf("no set limit changed the counts: %v", same)
	}
}
