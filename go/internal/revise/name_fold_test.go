package revise

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestKeepNamesReadsANameWithoutItsAccent is D-716: a card the deck
// holds takes the deck's own spelling, whatever the accent.
func TestKeepNamesReadsANameWithoutItsAccent(t *testing.T) {
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-footman", Name: "Gríma, Saruman's Footman", Count: 1}}}
	got := keepNames([]string{"grima, saruman's footman"}, deck)
	if len(got) != 1 || got[0] != "Gríma, Saruman's Footman" {
		t.Errorf("keep = %v, want the deck's spelling", got)
	}
	if got := onlyInDeck([]string{"Grima, Saruman's Footman"}, deck); len(got) != 1 {
		t.Errorf("remove = %v, want the deck card", got)
	}
}
