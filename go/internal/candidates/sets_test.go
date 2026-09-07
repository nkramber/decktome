package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// setFixture holds one card per set shape: two in the Hobbit family,
// one in a core set, and a ramp card outside every named set.
func setFixture(t *testing.T) *cards.Index {
	t.Helper()
	return fixture(t, []tc{
		{id: "plains", name: "Plains", typeLine: "Basic Land — Plains", rank: 1, sets: []string{"m19"}},
		{id: "insoul", name: "Soul Warden", typeLine: "Creature — Human Cleric",
			text:     "Whenever another creature enters, you gain 1 life.",
			identity: []mtgv1.Color{W}, mv: 1, rank: 200, tags: []string{"lifegain"}, sets: []string{"hob"}},
		{id: "inangel", name: "Resplendent Angel", typeLine: "Creature — Angel",
			text: "Flying.", identity: []mtgv1.Color{W}, mv: 3, rank: 400,
			tags: []string{"lifegain"}, keywords: []string{"Flying"}, sets: []string{"hoc"}},
		// No theme signal and no staple role: the cut drops it without a
		// set limit, and D-379 keeps it with one.
		{id: "inplain", name: "Hobbit Farmer", typeLine: "Creature — Halfling",
			identity: []mtgv1.Color{W}, mv: 2, rank: 900, sets: []string{"hob"}},
		{id: "outsoul", name: "Suture Priest", typeLine: "Creature — Cleric",
			text:     "Whenever another creature enters, you gain 1 life.",
			identity: []mtgv1.Color{W}, mv: 2, rank: 210, tags: []string{"lifegain"}, sets: []string{"m19"}},
		{id: "outrock", name: "Sol Ring", typeLine: "Artifact", text: "{T}: Add {C}{C}.",
			mv: 1, rank: 1, tags: []string{"ramp", "mana-rock"}, sets: []string{"m19"}},
		{id: "outrock2", name: "Arcane Signet", typeLine: "Artifact", text: "{T}: Add one mana.",
			mv: 2, rank: 2, tags: []string{"ramp", "mana-rock"}, sets: []string{"m19"}},
		{id: "incmdr", name: "Thranduil, the Elvenking", typeLine: "Legendary Creature — Elf Noble",
			text:     "Whenever you gain life, draw a card.",
			identity: []mtgv1.Color{W}, mv: 4, rank: 50, tags: []string{"lifegain"}, sets: []string{"hob"}},
		{id: "outcmdr", name: "Karlov of the Ghost Council", typeLine: "Legendary Creature — Spirit Advisor",
			text:     "Whenever you gain life, put counters on Karlov.",
			identity: []mtgv1.Color{W}, mv: 2, rank: 40, tags: []string{"lifegain"}, sets: []string{"m19"}},
	})
}

func idSet(cs []Candidate) map[string]bool {
	out := map[string]bool{}
	for _, c := range cs {
		out[c.Card.GetOracleId()] = true
	}
	return out
}

// TestBuildKeepsOnlyTheNamedSets is the whole point of PR-17B: a deck
// asked for one set holds cards of that set alone (D-373).
func TestBuildKeepsOnlyTheNamedSets(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	list, err := b.Build(setFixture(t), Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain",
		Colors: []mtgv1.Color{W}, SetCodes: []string{"hob", "hoc"},
		PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
	})
	if err != nil {
		t.Fatal(err)
	}
	got := idSet(list.Candidates)
	for _, id := range []string{"insoul", "inangel", "inplain"} {
		if !got[id] {
			t.Errorf("the shortlist lost %q, which the sets hold", id)
		}
	}
	for _, id := range []string{"outsoul", "outrock", "outrock2", "outcmdr"} {
		if got[id] {
			t.Errorf("the shortlist holds %q, which the sets do not", id)
		}
	}
	if list.Stats.InSet != len(got) {
		t.Errorf("InSet = %d, want %d", list.Stats.InSet, len(got))
	}
	if list.Stats.Outside != 0 {
		t.Errorf("Outside = %d, want 0 with no fill allowed", list.Stats.Outside)
	}
}

// TestSetLimitLiftsTheThemeCut is D-379. Without a set limit a card with
// no theme signal and no staple role is dropped. With one it stays, and
// the theme still ranks it below every on-theme card.
func TestSetLimitLiftsTheThemeCut(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := setFixture(t)
	base := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain",
		Colors: []mtgv1.Color{W}, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
	}
	open, err := b.Build(idx, base)
	if err != nil {
		t.Fatal(err)
	}
	if idSet(open.Candidates)["inplain"] {
		t.Fatal("with no set limit the theme cut must drop a card with no signal and no staple role")
	}
	base.SetCodes = []string{"hob", "hoc"}
	limited, err := b.Build(idx, base)
	if err != nil {
		t.Fatal(err)
	}
	if !idSet(limited.Candidates)["inplain"] {
		t.Fatal("with a set limit the theme must rank and not cut (D-379)")
	}
	// The theme still leads: an on-theme card outscores the one with no
	// signal, because the staple penalty halves a no-signal score.
	var themed, plain float64
	for _, c := range limited.Candidates {
		switch c.Card.GetOracleId() {
		case "insoul":
			themed = c.Score
		case "inplain":
			plain = c.Score
		}
	}
	if themed <= plain {
		t.Errorf("on-theme score %v does not beat the no-signal score %v", themed, plain)
	}
}

// TestOutsideRolesFillTheManaBase is D-382: a set family short of mana
// cards takes them from the whole database, up to the wanted count.
func TestOutsideRolesFillTheManaBase(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	list, err := b.Build(setFixture(t), Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain",
		Colors: []mtgv1.Color{W}, SetCodes: []string{"hob", "hoc"},
		PoolRule:     mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		OutsideRoles: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_RAMP: 2},
	})
	if err != nil {
		t.Fatal(err)
	}
	got := idSet(list.Candidates)
	if !got["outrock"] || !got["outrock2"] {
		t.Errorf("the ramp fill did not reach the list: %v", got)
	}
	// The fill covers the named roles alone. A lifegain creature outside
	// the sets is not a mana card.
	if got["outsoul"] {
		t.Error("the fill took a card whose role it does not name")
	}
	if got["outcmdr"] {
		t.Error("the fill took a commander from outside the sets")
	}
	if list.Stats.Outside != 2 {
		t.Errorf("Outside = %d, want 2", list.Stats.Outside)
	}
	for _, c := range list.Candidates {
		if c.Card.GetOracleId() == "outrock" && !c.Outside {
			t.Error("a filled card must carry the outside mark (D-383)")
		}
		if c.Card.GetOracleId() == "insoul" && c.Outside {
			t.Error("a card the sets hold must not carry the outside mark")
		}
	}
}

// TestOutsideFillStopsAtTheWantedCount: a role the sets already fill
// takes nothing from outside (D-382).
func TestOutsideFillStopsAtTheWantedCount(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	list, err := b.Build(setFixture(t), Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain",
		Colors: []mtgv1.Color{W}, SetCodes: []string{"hob", "hoc", "m19"},
		PoolRule:     mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		OutsideRoles: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_RAMP: 2},
	})
	if err != nil {
		t.Fatal(err)
	}
	if list.Stats.Outside != 0 {
		t.Errorf("Outside = %d, want 0 when the sets hold every card", list.Stats.Outside)
	}
}

// TestCommanderPoolKeepsTheNamedSets is the second gate line of PR-17B:
// a commander offer for one set names commanders of that set.
func TestCommanderPoolKeepsTheNamedSets(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	pool, err := b.CommanderPool(setFixture(t), Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain",
		SetCodes: []string{"hob", "hoc"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(pool) == 0 {
		t.Fatal("the commander pool is empty")
	}
	for _, c := range pool {
		if c.Card.GetOracleId() == "outcmdr" {
			t.Error("the pool offers a commander the sets do not hold")
		}
	}
	if pool[0].Card.GetName() != "Thranduil, the Elvenking" {
		t.Errorf("first commander = %q, want Thranduil, the Elvenking", pool[0].Card.GetName())
	}
}

// TestCommanderPoolFillsFromTheSetsOnly is D-367 under a set limit: the
// unthemed fill must not reach past the named sets.
func TestCommanderPoolFillsFromTheSetsOnly(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	// A theme the tag table does not know leaves the pool empty, so the
	// unthemed fill of D-367 runs.
	pool, err := b.CommanderPool(setFixture(t), Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "zzzunknownzzz",
		SetCodes: []string{"hob", "hoc"},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range pool {
		if c.Card.GetOracleId() == "outcmdr" {
			t.Fatal("the unthemed fill offered a commander outside the named sets")
		}
	}
}

// TestCountInSetsCountsTheNonbasics is the number the floor of D-380
// reads. A basic land is out of it: it repeats without limit and no set
// limit filters it.
func TestCountInSetsCountsTheNonbasics(t *testing.T) {
	idx := setFixture(t)
	got := CountInSets(idx, Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Colors: []mtgv1.Color{W},
		SetCodes: []string{"hob", "hoc"},
	})
	// insoul, inangel, inplain, incmdr. Plains is basic, and the m19
	// cards are outside the sets.
	if got != 4 {
		t.Errorf("CountInSets = %d, want 4", got)
	}
	if n := CountInSets(idx, Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER}); n != 0 {
		t.Errorf("CountInSets with no set limit = %d, want 0", n)
	}
}

// TestSetFloorPerFormat is D-380.
func TestSetFloorPerFormat(t *testing.T) {
	if got := SetFloor(mtgv1.FormatId_FORMAT_ID_COMMANDER); got != 70 {
		t.Errorf("Commander floor = %d, want 70", got)
	}
	for _, f := range []mtgv1.FormatId{
		mtgv1.FormatId_FORMAT_ID_MODERN, mtgv1.FormatId_FORMAT_ID_STANDARD,
		mtgv1.FormatId_FORMAT_ID_HOUSE,
	} {
		if got := SetFloor(f); got != 35 {
			t.Errorf("%s floor = %d, want 35", f, got)
		}
	}
}

// TestSetLimitDropsTheRoleCaps is the second half of D-379. The role
// caps shape a shortlist drawn from the whole database. Drawn from one
// set family they only lose cards: the "other" cap of 10 cut 40 of the
// 128 cards the Hobbit family offers in black-red, and a Commander deck
// needs 99.
func TestSetLimitDropsTheRoleCaps(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	// Twenty cards of one role, over the "other" cap of 10.
	list := []tc{{id: "cmd", name: "Thranduil, the Elvenking", typeLine: "Legendary Creature — Elf",
		identity: []mtgv1.Color{W}, rank: 10, sets: []string{"hob"}}}
	for i := 0; i < 20; i++ {
		list = append(list, tc{
			id: "o" + string(rune('a'+i)), name: "Hobbit Card " + string(rune('a'+i)),
			typeLine: "Enchantment", identity: []mtgv1.Color{W}, mv: 2,
			rank: int32(100 + i), sets: []string{"hob"},
		})
	}
	idx := fixture(t, list)
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "hobbits",
		Colors: []mtgv1.Color{W}, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
	}
	// With no set limit the "other" cap holds.
	open, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if n := countRole(open.Candidates, mtgv1.CardRole_CARD_ROLE_OTHER); n > DefaultLimits.PerRole[mtgv1.CardRole_CARD_ROLE_OTHER] {
		t.Fatalf("with no set limit the other role holds %d cards, over the cap", n)
	}
	req.SetCodes = []string{"hob", "hoc"}
	limited, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if got := len(limited.Candidates); got != 21 {
		t.Errorf("the set-limited shortlist holds %d cards, want every one of the 21 in the set", got)
	}
	if limited.Stats.InSet != 21 {
		t.Errorf("InSet = %d, want 21", limited.Stats.InSet)
	}
	// A caller that names its own caps keeps them.
	req.Limits = Limits{PerRole: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_OTHER: 3}}
	capped, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if n := countRole(capped.Candidates, mtgv1.CardRole_CARD_ROLE_OTHER); n != 3 {
		t.Errorf("an explicit cap of 3 kept %d cards", n)
	}
}

func countRole(cs []Candidate, want mtgv1.CardRole) int {
	n := 0
	for _, c := range cs {
		if c.Role == want {
			n++
		}
	}
	return n
}

// TestCommanderPoolOffersPaperCardsOnly is D-306. The unthemed fill of
// D-367 always tested this, and the themed half never did, so a
// digital-only legend could lead an offer.
func TestCommanderPoolOffersPaperCardsOnly(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, []tc{
		{id: "paper", name: "Thranduil, the Elvenking", typeLine: "Legendary Creature — Elf",
			text: "Whenever you gain life, draw a card.", identity: []mtgv1.Color{W},
			mv: 4, rank: 90, tags: []string{"lifegain"}, sets: []string{"hob"}},
		{id: "digital", name: "Alchemy Legend", typeLine: "Legendary Creature — Spirit",
			text: "Whenever you gain life, draw a card.", identity: []mtgv1.Color{W},
			mv: 3, rank: 1, tags: []string{"lifegain"}, digital: true},
	})
	pool, err := b.CommanderPool(idx, Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain",
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range pool {
		if c.Card.GetOracleId() == "digital" {
			t.Fatal("the commander pool offered a card with no paper printing (D-306)")
		}
	}
	if len(pool) == 0 || pool[0].Card.GetOracleId() != "paper" {
		t.Errorf("the paper commander is not the offer: %v", pool)
	}
}
