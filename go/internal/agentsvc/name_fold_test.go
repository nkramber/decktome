package agentsvc

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestKeepableReadsADeckCardWithoutItsAccent is D-716: a card the deck
// already holds stays, whatever the accent the reader typed.
func TestKeepableReadsADeckCardWithoutItsAccent(t *testing.T) {
	s := &Server{index: fixedIndex{keepIndex()}}
	base := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-footman", Name: "Gríma, Saruman's Footman", Count: 1}}}
	keep, refused := s.keepable(base, []string{"Grima, Saruman's Footman"}, map[string]int32{}, mtgv1.PoolRule_POOL_RULE_OWNED_ONLY)
	if len(keep) != 1 || len(refused) != 0 {
		t.Errorf("keep = %v, refused = %v, want the deck card kept", keep, refused)
	}
}
