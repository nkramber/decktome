package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// pinCards are the power cards of powerCards and an on-theme ramp card.
func pinCards() []tc {
	return append(powerCards(), tc{id: "stone", name: "Lifegain Stone", typeLine: "Artifact",
		text: "Whenever you gain life, add {W}.", identity: []mtgv1.Color{W}, mv: 2, rank: 5000, tags: []string{"ramp", "lifegain"}})
}

// pinRequest is a bracket 5 lifegain list in which Sol Ring alone has a
// top-list rate.
func pinRequest() Request {
	return Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "lifegain", Bracket: 5,
		MetaBoost: fixedRates(map[string]float64{"solring": 0.9})}
}

// TestPinnedPowerCardSkipsItsRoleCap is D-710. A bracket 5 list caps ramp
// at one card, and an on-theme ramp card takes that place. Sol Ring is
// fast mana at the keep rate: with no pin the cap drops it, and with a pin
// it stays beside the on-theme card.
func TestPinnedPowerCardSkipsItsRoleCap(t *testing.T) {
	b, _ := New()
	idx := fixture(t, pinCards())
	req := pinRequest()
	req.Limits = Limits{PerRole: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_RAMP: 1}}
	for _, pin := range []bool{false, true} {
		req.PowerRate = PowerRate{Weight: 0.1, Keep: 0.5, Pin: pin}
		list, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		sol, hasSol := find(list.Candidates, "Sol Ring")
		_, hasStone := find(list.Candidates, "Lifegain Stone")
		if hasSol != pin || sol.Pinned != pin || !hasStone {
			t.Errorf("pin %v: Sol Ring listed %v and pinned %v, Lifegain Stone listed %v", pin, hasSol, sol.Pinned, hasStone)
		}
	}
}

// TestPinnedPowerCardAddsToTheTotal is F-132 and D-712. A pinned card takes
// no place under the total, so the list with the pin holds each card of the
// list with no pin, the total of unpinned cards, and the pinned card beside
// them. The first pin took a place, and the trim dropped fixing lands.
func TestPinnedPowerCardAddsToTheTotal(t *testing.T) {
	b, _ := New()
	idx := fixture(t, pinCards())
	req := pinRequest()
	req.Limits = Limits{Total: 3}
	req.PowerRate = PowerRate{Weight: 0.1, Keep: 0.5}
	without, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	req.PowerRate.Pin = true
	with, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if len(without.Candidates) != 3 {
		t.Fatalf("the list with no pin holds %v, want 3 cards", names(without.Candidates))
	}
	for _, c := range without.Candidates {
		if c.Card.GetName() == "Sol Ring" {
			continue
		}
		if _, ok := find(with.Candidates, c.Card.GetName()); !ok {
			t.Errorf("the pin dropped %s: with the pin %v, with no pin %v", c.Card.GetName(), names(with.Candidates), names(without.Candidates))
		}
	}
	unpinned := 0
	for _, c := range with.Candidates {
		if !c.Pinned {
			unpinned++
		}
	}
	if sol, ok := find(with.Candidates, "Sol Ring"); !ok || !sol.Pinned || unpinned != 3 {
		t.Errorf("with the pin the list holds %v: Sol Ring pinned %v, unpinned %d, want Sol Ring pinned beside 3 unpinned cards",
			names(with.Candidates), ok && sol.Pinned, unpinned)
	}
}
