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

// TestOwnedInSetsCountsOnlyTheNamedSetCopies is D-1052. A reader who owns
// a card from a named set and from another set owns, for this deck, the
// copies of the named set alone. The count the deck shows is that number.
func TestOwnedInSetsCountsOnlyTheNamedSetCopies(t *testing.T) {
	idx := cards.NewIndex(
		[]*mtgv1.Card{
			{OracleId: "o-sol", Name: "Sol Ring", CardTypes: []string{"Artifact"}},
			{OracleId: "o-plains", Name: "Plains", Supertypes: []string{"Basic"}, CardTypes: []string{"Land"}},
		},
		[]cards.Printing{
			{ScryfallID: "p-sol-ltc", OracleID: "o-sol", Name: "Sol Ring", SetCode: "LTC", CollectorNumber: "1"},
			{ScryfallID: "p-sol-c21", OracleID: "o-sol", Name: "Sol Ring", SetCode: "C21", CollectorNumber: "2"},
			{ScryfallID: "p-plains-ltr", OracleID: "o-plains", Name: "Plains", SetCode: "LTR", CollectorNumber: "3"},
			{ScryfallID: "p-plains-m21", OracleID: "o-plains", Name: "Plains", SetCode: "M21", CollectorNumber: "4"},
		},
		nil, time.Unix(1000, 0).UTC())
	owned := map[string]int32{"o-sol": 4, "o-plains": 30}
	printings := map[string]int32{"p-sol-ltc": 1, "p-sol-c21": 3, "p-plains-ltr": 10, "p-plains-m21": 20}

	got := OwnedInSets(idx, owned, printings, []string{"ltr", "ltc"})
	if got["o-sol"] != 1 {
		t.Errorf("Sol Ring = %d, want the 1 copy of LTC and not the 3 of C21", got["o-sol"])
	}
	if got["o-plains"] != 30 {
		t.Errorf("Plains = %d, want all 30: a basic land keeps its whole count", got["o-plains"])
	}
}
