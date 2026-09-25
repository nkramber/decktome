package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// colorlessCards holds a colorless lifegain legend beside a white one, a
// colorless lifegain artifact, and a white lifegain card.
func colorlessCards() []tc {
	return append(commanderCards(),
		tc{id: "kozlegend", name: "Colorless Legend", typeLine: "Legendary Creature — Eldrazi",
			text: "Whenever you gain life, draw a card.", mv: 4, rank: 50, tags: []string{"lifegain"}},
		tc{id: "lifestone", name: "Colorless Lifestone", typeLine: "Artifact",
			text: "Whenever you gain life, scry 1.", mv: 2, rank: 60, tags: []string{"lifegain"}},
	)
}

// TestAColorlessCommanderGetsAColorlessShortlist is REV-017 of the review
// of 2026-09-24. A commander with no color identity sent an empty color
// set, which read as every color, so the model got cards of every color
// and the engine blocked the deck.
func TestAColorlessCommanderGetsAColorlessShortlist(t *testing.T) {
	idx := fixture(t, colorlessCards())
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	list, err := b.Build(idx, Request{Format: cmdr, Theme: "lifegain", Colorless: true, CommanderOracleIDs: []string{"kozlegend"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Candidates) == 0 {
		t.Fatal("the colorless shortlist is empty, so the test reads nothing")
	}
	for _, c := range list.Candidates {
		if len(c.Card.GetColorIdentity()) > 0 {
			t.Errorf("%s of identity %v reached a colorless shortlist", c.Card.GetName(), c.Card.GetColorIdentity())
		}
	}
}

// TestTheOfferHoldsAColorlessCommanderOnRequestAlone is the answer of the
// owner to REV-017: the app supports a colorless commander, and it offers
// one only when the user asked for a colorless deck (D-915).
func TestTheOfferHoldsAColorlessCommanderOnRequestAlone(t *testing.T) {
	idx := fixture(t, colorlessCards())
	b, _ := New()
	for _, colorless := range []bool{false, true} {
		got, err := b.CommanderPool(idx, Request{Format: cmdr, Theme: "lifegain", Colorless: colorless})
		if err != nil {
			t.Fatal(err)
		}
		offered := contains(names(got), "Colorless Legend")
		if offered != colorless {
			t.Errorf("colorless request %v: the offer %v holds the colorless legend %v", colorless, names(got), offered)
		}
		if colorless {
			for _, c := range got {
				if len(c.Card.GetColorIdentity()) > 0 {
					t.Errorf("a colorless request offered %s", c.Card.GetName())
				}
			}
		}
	}
}

var _ = mtgv1.Color_COLOR_W
