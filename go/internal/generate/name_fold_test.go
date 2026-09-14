package generate

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The D-716 tests of the build. A name without its accent names the pool
// card, and a revision that keeps it reads the deck card.

func TestPoolCardReadsANameWithoutItsAccent(t *testing.T) {
	p := NewPool([]*mtgv1.Card{{OracleId: "o-footman", Name: "Gríma, Saruman's Footman"}}, nil)
	c, ok := p.Card("Grima, Saruman's Footman")
	if !ok || c.GetOracleId() != "o-footman" {
		t.Errorf("Card = %q, %v, want the pool card", c.GetOracleId(), ok)
	}
}

// TestCheckRevisionReadsAKeptNameWithoutItsAccent: the reader kept the
// card in their own words, and the deck holds it under its real name.
func TestCheckRevisionReadsAKeptNameWithoutItsAccent(t *testing.T) {
	cards := cardMap{"o-footman": {OracleId: "o-footman", Name: "Gríma, Saruman's Footman", CardTypes: []string{"Creature"}}}
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{{OracleId: "o-footman", Name: "Gríma, Saruman's Footman", Count: 1}}}
	rev := &Revision{Keep: []string{"Grima, Saruman's Footman"}}
	for _, f := range CheckRevision(deck, rev, cards) {
		if f.GetCode() == CodeRevisionKeptMissing {
			t.Errorf("the kept card reads as missing: %s", f.GetMessage())
		}
	}
}
