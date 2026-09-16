package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// The PR-53 tests of the wincon role (F-138, D-726). The role reads the
// curated finisher count of M-17: the parent tags with no child tag, and
// the evasive creatures of power 5 or more.

func finisherCards() []tc {
	return append(testCards(),
		tc{id: "artist", name: "Bloodletting Artist", typeLine: "Creature — Vampire", rank: 250, mv: 2, identity: []mtgv1.Color{B},
			text: "Whenever this creature or another creature dies, target player loses 1 life and you gain 1 life.",
			tags: []string{"blood-artist-ability", "lifegain"}},
		tc{id: "drainer", name: "Exsanguinate", typeLine: "Sorcery", rank: 260, mv: 2, identity: []mtgv1.Color{B},
			text: "Each opponent loses X life. You gain life equal to the life lost this way.",
			tags: []string{"drain-life", "removal", "lifegain"}},
		tc{id: "pilferer", name: "Ragavan, Nimble Pilferer", typeLine: "Creature — Monkey Pirate", rank: 40, mv: 1, identity: []mtgv1.Color{B},
			text: "Whenever this creature deals combat damage to a player, that player mills a card.",
			tags: []string{"mill-any", "lifegain"}},
		tc{id: "flier", name: "Grave Dragon", typeLine: "Creature — Dragon", rank: 420, mv: 6, identity: []mtgv1.Color{B},
			text: "Flying", keywords: []string{"Flying"}, tags: []string{"evasion", "lifegain"}},
		tc{id: "smallflier", name: "Gloom Bat", typeLine: "Creature — Bat", rank: 430, mv: 3, identity: []mtgv1.Color{B},
			text: "Flying", keywords: []string{"Flying"}, tags: []string{"evasion", "lifegain"}},
	)
}

// finisherIndex builds the fixture and gives the two fliers their power.
func finisherIndex(t *testing.T) *cards.Index {
	t.Helper()
	idx := fixture(t, finisherCards())
	for name, power := range map[string]string{"Grave Dragon": "5", "Gloom Bat": "2"} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("no card named %q", name)
		}
		c.Power = power
	}
	return idx
}

// TestAFinisherTagTakesTheWinconRole is D-726. The role read the tag
// alternate-win-condition alone, so a drain spell read as removal and the
// shortlist rarely held a win route.
func TestAFinisherTagTakesTheWinconRole(t *testing.T) {
	b, _ := New()
	idx := finisherIndex(t)
	list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{B}, Theme: "lifegain", Bracket: 3})
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct {
		name string
		want mtgv1.CardRole
	}{
		{"Bloodletting Artist", mtgv1.CardRole_CARD_ROLE_WINCON},
		{"Exsanguinate", mtgv1.CardRole_CARD_ROLE_WINCON},
		{"Grave Dragon", mtgv1.CardRole_CARD_ROLE_WINCON},
		{"Gloom Bat", mtgv1.CardRole_CARD_ROLE_SYNERGY},
		{"Ragavan, Nimble Pilferer", mtgv1.CardRole_CARD_ROLE_SYNERGY},
	} {
		c, ok := find(list.Candidates, tt.name)
		if !ok {
			t.Errorf("the shortlist holds no %s", tt.name)
			continue
		}
		if c.Role != tt.want {
			t.Errorf("%s reads the role %s, want %s", tt.name, c.Role, tt.want)
		}
	}
}

// TestASurplusFinisherKeepsItsRole is D-741. The role went to every card
// of the curated count, and the wincon cap of 15 then dropped finishers
// that the list held before. The target count takes the role now, and
// every other finisher keeps the role it earned.
func TestASurplusFinisherKeepsItsRole(t *testing.T) {
	b, _ := New()
	idx := finisherIndex(t)
	req := Request{Format: cmdr, Colors: []mtgv1.Color{B}, Theme: "lifegain", Bracket: 5}
	list, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	// Bracket 5 asks for one finisher, and the fixture holds three.
	var wincons, others []string
	for _, c := range list.Candidates {
		if !finisherNames[c.Card.GetName()] {
			continue
		}
		if c.Role == mtgv1.CardRole_CARD_ROLE_WINCON {
			wincons = append(wincons, c.Card.GetName())
			continue
		}
		others = append(others, c.Card.GetName())
	}
	if len(wincons) != 1 {
		t.Errorf("bracket 5 reads %v as wincon, want one card", wincons)
	}
	if len(others) != 2 {
		t.Errorf("the surplus finishers read %v, want the other two of the three", others)
	}
	// No finisher leaves the list, whatever the wincon cap holds.
	if len(wincons)+len(others) != len(finisherNames) {
		t.Errorf("the list holds %d of the %d finishers", len(wincons)+len(others), len(finisherNames))
	}
}

// finisherNames are the fixture cards of the curated count.
var finisherNames = map[string]bool{"Bloodletting Artist": true, "Exsanguinate": true, "Grave Dragon": true}

// TestTheShortlistPinsTheFinishersOfTheBracket is D-726. A role cap can
// not drop the win route, and the pin stops at the target of the bracket.
func TestTheShortlistPinsTheFinishersOfTheBracket(t *testing.T) {
	b, _ := New()
	idx := finisherIndex(t)
	for _, tt := range []struct {
		bracket int32
		want    int
	}{{3, 3}, {5, 1}} {
		req := Request{Format: cmdr, Colors: []mtgv1.Color{B}, Theme: "lifegain", Bracket: tt.bracket,
			Limits: Limits{PerRole: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_WINCON: 0}}}
		list, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		pinned := 0
		for _, c := range list.Candidates {
			if c.Role == mtgv1.CardRole_CARD_ROLE_WINCON && c.Pinned {
				pinned++
			}
		}
		if pinned != tt.want {
			t.Errorf("bracket %d pins %d finishers, want %d: %v", tt.bracket, pinned, tt.want, names(list.Candidates))
		}
		if got := FinisherTarget(tt.bracket); got != tt.want {
			t.Errorf("bracket %d reads a finisher target of %d, want %d", tt.bracket, got, tt.want)
		}
	}
}

// TestAnOwnedModePinsTheOwnedFinishers is D-742. The promotion ran over
// the whole pool, and an owned mode drops the unowned cards after it. So
// a promotion landed on a card the mode drops, and the shortlist then
// read under the target of its bracket.
func TestAnOwnedModePinsTheOwnedFinishers(t *testing.T) {
	b, _ := New()
	idx := finisherIndex(t)
	// The reader owns the two weaker finishers, and not the best one.
	owned := map[string]int32{"drainer": 1, "flier": 1}
	for _, c := range testCards() {
		owned[c.id] = 1
	}
	delete(owned, "artist")
	for _, rule := range []mtgv1.PoolRule{
		mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
	} {
		req := Request{Format: cmdr, Colors: []mtgv1.Color{B}, Theme: "lifegain", Bracket: 5,
			PoolRule: rule, Owned: owned}
		list, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		var pinned []string
		for _, c := range list.Candidates {
			if c.Role == mtgv1.CardRole_CARD_ROLE_WINCON && c.Pinned {
				pinned = append(pinned, c.Card.GetName())
				if c.Owned == 0 {
					t.Errorf("%s: the pin took the unowned %s", rule, c.Card.GetName())
				}
			}
		}
		if len(pinned) != FinisherTarget(req.Bracket) {
			t.Errorf("%s pins %v, want %d finisher of the collection", rule, pinned, FinisherTarget(req.Bracket))
		}
	}
	// Any-card reads the best finisher of the pool, owned or not.
	list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{B}, Theme: "lifegain",
		Bracket: 5, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD, Owned: owned})
	if err != nil {
		t.Fatal(err)
	}
	c, ok := find(list.Candidates, "Bloodletting Artist")
	if !ok || c.Role != mtgv1.CardRole_CARD_ROLE_WINCON || !c.Pinned {
		t.Errorf("any-card reads Bloodletting Artist as %v, want the pinned wincon", c.Role)
	}
}

// TestAPrePinnedWinconCountsTowardTheTarget is the Gitar review of #182.
// `roles.go` names a wincon from the tag alternate-win-condition, and
// `pinPower` pins a finisher that reaches the keep rate. So a card can
// read wincon and pinned before the promotion. The promotion stepped over
// it and counted none of it, and it then added the whole target on top.
// Each extra pinned card skips its role cap and adds to the total.
func TestAPrePinnedWinconCountsTowardTheTarget(t *testing.T) {
	// The reader owns every card, so each pool rule can promote.
	mk := func(id string, role mtgv1.CardRole, pinned bool) Candidate {
		return Candidate{Card: &mtgv1.Card{OracleId: id, Name: id}, Role: role, Pinned: pinned, Owned: 1}
	}
	build := func() []Candidate {
		return []Candidate{
			mk("prepinned", mtgv1.CardRole_CARD_ROLE_WINCON, true),
			mk("a", mtgv1.CardRole_CARD_ROLE_THREAT, false),
			mk("b", mtgv1.CardRole_CARD_ROLE_SYNERGY, false),
			mk("c", mtgv1.CardRole_CARD_ROLE_REMOVAL, false),
			mk("d", mtgv1.CardRole_CARD_ROLE_SYNERGY, false),
		}
	}
	finishers := map[string]bool{"prepinned": true, "a": true, "b": true, "c": true, "d": true}
	pinnedWincons := func(cs []Candidate) []string {
		var out []string
		for _, c := range cs {
			if c.Role == mtgv1.CardRole_CARD_ROLE_WINCON && c.Pinned {
				out = append(out, c.Card.GetName())
			}
		}
		return out
	}
	for _, mode := range []mtgv1.PoolRule{
		mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
		mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
	} {
		cs := build()
		req := Request{Format: cmdr, Bracket: 4}
		promoteFinishers(cs, req, mode, FinisherTarget(4), finishers)
		got := pinnedWincons(cs)
		if len(got) != FinisherTarget(4) {
			t.Errorf("%s pins %v, want %d of them", mode, got, FinisherTarget(4))
		}
		// The card that already read wincon keeps the role and the pin.
		if got[0] != "prepinned" {
			t.Errorf("%s dropped the card that already read wincon: %v", mode, got)
		}
	}
	// A pinned card of another role never counts toward the target.
	cs := build()
	cs[0] = mk("prepinned", mtgv1.CardRole_CARD_ROLE_RAMP, true)
	promoteFinishers(cs, Request{Format: cmdr, Bracket: 4}, mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		FinisherTarget(4), finishers)
	if got := pinnedWincons(cs); len(got) != FinisherTarget(4) {
		t.Errorf("a pinned ramp card moved the count: %v", got)
	}
}
