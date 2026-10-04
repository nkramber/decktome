package candidates

import (
	"fmt"
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// f212Cards is a mono-blue voltron pool with the cards the theme words of
// F-212 name: hand size, cantrips, targeting spells, and staples.
func f212Cards() []tc {
	U := mtgv1.Color_COLOR_U
	u := []mtgv1.Color{U}
	return []tc{
		{id: "kamala", name: "Kamala Test Commander", typeLine: "Legendary Creature — Human Hero",
			text:     "You have no maximum hand size. Whenever you cast a spell that targets a creature you control, draw a card.",
			identity: u, mv: 3, rank: 900},
		{id: "boots", name: "Swiftfoot Boots", typeLine: "Artifact — Equipment", subtypes: []string{"Equipment"},
			text: "Equipped creature has hexproof and haste. Equip {1}", mv: 2, rank: 40, tags: []string{"protection", "synergy-equipment"}},
		{id: "jitte", name: "Umezawa's Jitte", typeLine: "Legendary Artifact — Equipment", subtypes: []string{"Equipment"},
			text: "Whenever equipped creature deals combat damage, put two charge counters on Umezawa's Jitte.", mv: 2, rank: 300,
			tags: []string{"removal", "synergy-equipment"}},
		{id: "divedown", name: "Dive Down", typeLine: "Instant", text: "Target creature you control gets +0/+3 and gains hexproof until end of turn.",
			identity: u, mv: 1, rank: 700, tags: []string{"protection"}},
		{id: "counter", name: "Counterspell", typeLine: "Instant", text: "Counter target spell.", identity: u, mv: 2, rank: 20, tags: []string{"counterspell"}},
		{id: "hellkite", name: "Sweeping Hellkite", typeLine: "Artifact Creature — Construct", text: "Flying. Deals 1 damage to each creature your opponents control.",
			mv: 6, rank: 800, tags: []string{"sweeper", "evasion"}},
		{id: "tower", name: "Reliquary Tower", typeLine: "Land", text: "You have no maximum hand size. {T}: Add {C}.", mv: 0, rank: 60,
			tags: []string{"hand-size-increase"}},
		{id: "maro", name: "Hand Maro", typeLine: "Creature — Spirit", text: "Hand Maro's power is equal to the number of cards in your hand.",
			identity: u, mv: 4, rank: 4000, tags: []string{"maro"}},
		{id: "opt", name: "Opt", typeLine: "Instant", text: "Scry 1. Draw a card.", identity: u, mv: 1, rank: 90, tags: []string{"cantrip", "draw"}},
		{id: "flyer", name: "Unblockable Drake", typeLine: "Creature — Drake", text: "Unblockable Drake can't be blocked.", identity: u, mv: 3, rank: 5000,
			tags: []string{"evasion"}},
		{id: "aura", name: "Arcane Flight", typeLine: "Enchantment — Aura", subtypes: []string{"Aura"},
			text: "Enchant creature you control. It gets +1/+1 and has flying.", identity: u, mv: 1, rank: 6000},
		{id: "shoreup", name: "Shore Up", typeLine: "Instant", text: "Target creature you control gets +1/+1 until end of turn. Untap it.",
			identity: u, mv: 1, rank: 6500},
		{id: "zz", name: "Zz Widget", typeLine: "Artifact", text: "Zz.", mv: 1, rank: 7000, tags: []string{"zzwidget"}},
	}
}

// TestF212ThemeWordsLeaveNoDeadWord is F-212 and D-1116. "handsize" and
// "cantrips" found no row, and "evasion" read as unmatched because the
// word voltron brought the tag evasion first. A word is unmatched only
// when its own signals fire on no card.
func TestF212ThemeWordsLeaveNoDeadWord(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, f212Cards())
	cases := map[string][]string{
		"handsize matters, heavy on cantrips": {"handsize", "cantrips"},
		"hand size":                           {"hand-size"},
		"voltron, evasion":                    {"voltron", "evasion"},
		"voltron, evasion, protection, handsize matters, heavy on cantrips": {"voltron", "evasion", "protection", "handsize", "cantrips"},
		// A word with no row reads the singular slug too.
		"zzwidgets": {"zzwidgets"},
	}
	for theme, words := range cases {
		list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{mtgv1.Color_COLOR_U}, Theme: theme})
		if err != nil {
			t.Fatal(err)
		}
		if !slices.Equal(list.Theme.Words, words) {
			t.Errorf("%q: words %v, want %v", theme, list.Theme.Words, words)
		}
		if len(list.Theme.Unmatched) > 0 {
			t.Errorf("%q left words unmatched: %s", theme, list.Theme.Describe())
		}
	}
	list, err := b.Build(idx, Request{Format: cmdr, Theme: "voltron zzzz"})
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(list.Theme.Unmatched, []string{"zzzz"}) {
		t.Errorf("unmatched %v, want [zzzz]", list.Theme.Unmatched)
	}
	hand, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{mtgv1.Color_COLOR_U}, Theme: "handsize"})
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Reliquary Tower", "Hand Maro"} {
		if c, ok := find(hand.Candidates, name); !ok || !c.Themed {
			t.Errorf("handsize must list %s on theme: %+v, %v", name, c, ok)
		}
	}
}

// TestF212Roles is D-1120. A protection permanent takes the role
// protection, a protection spell and a counterspell take interaction, a
// creature with the sweeper tag is no wipe, and equipment with the
// removal tag is no removal.
func TestF212Roles(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, f212Cards())
	roleTags := b.themes.roleSets(idx.Tags())
	want := map[string]mtgv1.CardRole{
		"boots":    mtgv1.CardRole_CARD_ROLE_PROTECTION,
		"divedown": mtgv1.CardRole_CARD_ROLE_INTERACTION,
		"counter":  mtgv1.CardRole_CARD_ROLE_INTERACTION,
		"hellkite": mtgv1.CardRole_CARD_ROLE_THREAT,
		"jitte":    mtgv1.CardRole_CARD_ROLE_SYNERGY,
	}
	for id, role := range want {
		c, _ := idx.ByOracleID(id)
		if got, _ := assignRole(c, roleTags, true, false); got != role {
			t.Errorf("%s: role %s, want %s", c.Name, RoleName(got), RoleName(role))
		}
	}
	// Off theme, the sweeper creature and the equipment fill no staple role.
	for _, id := range []string{"hellkite", "jitte"} {
		c, _ := idx.ByOracleID(id)
		if got, _ := assignRole(c, roleTags, false, false); got != mtgv1.CardRole_CARD_ROLE_OTHER {
			t.Errorf("%s off theme: role %s, want other", c.Name, RoleName(got))
		}
	}
}

// TestF212AvoidRanksLower is D-1122. "less artifacts" names the cards
// that reward artifacts, so it lowers the artifact payoff and leaves
// Swiftfoot Boots, an artifact with no artifact payoff, at its score.
// "no artifacts" is a hard request, and it lowers every artifact.
func TestF212AvoidRanksLower(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	pool := append(f212Cards(), tc{id: "thopterist", name: "Thopter Payoff", typeLine: "Creature — Vedalken Artificer",
		text: "Whenever an artifact enters the battlefield under your control, create a 1/1 Thopter.", identity: []mtgv1.Color{mtgv1.Color_COLOR_U},
		mv: 2, rank: 3000, tags: []string{"synergy-artifact"}})
	idx := fixture(t, pool)
	req := Request{Format: cmdr, Colors: []mtgv1.Color{mtgv1.Color_COLOR_U}, Theme: "voltron, artifacts"}
	base, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if base.Stats.Avoided != 0 {
		t.Errorf("avoided %d with no slot, want 0", base.Stats.Avoided)
	}
	build := func(avoid string) *List {
		req.Avoid = avoid
		l, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		return l
	}
	moved := func(l *List, name string) bool {
		before, ok1 := find(base.Candidates, name)
		after, ok2 := find(l.Candidates, name)
		if !ok1 || !ok2 {
			t.Fatalf("%s must stay listed: %v, %v", name, ok1, ok2)
		}
		if after.Avoided != (after.Score < before.Score) {
			t.Errorf("%s: avoided %v, score %.3f to %.3f", name, after.Avoided, before.Score, after.Score)
		}
		return after.Avoided
	}
	soft := build("less artifacts")
	if !moved(soft, "Thopter Payoff") {
		t.Error(`"less artifacts" must lower the artifact payoff`)
	}
	for _, name := range []string{"Swiftfoot Boots", "Umezawa's Jitte", "Counterspell"} {
		if moved(soft, name) {
			t.Errorf(`"less artifacts" lowered %s, which rewards no artifact`, name)
		}
	}
	hard := build("no artifacts")
	for _, c := range base.Candidates {
		if slices.Contains(c.Card.CardTypes, "Artifact") && !moved(hard, c.Card.Name) {
			t.Errorf(`"no artifacts" must lower the artifact %s`, c.Card.Name)
		}
	}
	if moved(hard, "Counterspell") {
		t.Error(`"no artifacts" lowered Counterspell`)
	}
	if hard.Stats.Avoided == 0 || soft.Stats.Avoided >= hard.Stats.Avoided {
		t.Errorf("avoided %d soft and %d hard, want fewer soft", soft.Stats.Avoided, hard.Stats.Avoided)
	}
}

// TestF212StapleFloorKeepsCounterspell is D-1121. Twenty-five on-theme
// protection spells fill the interaction cap, and Counterspell has no
// theme signal. The floor of the role keeps it beside the cap and past
// the total cut. With no floor, the cap drops it.
func TestF212StapleFloorKeepsCounterspell(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	u := []mtgv1.Color{mtgv1.Color_COLOR_U}
	pool := f212Cards()
	for i := range 25 {
		pool = append(pool, tc{id: fmt.Sprintf("ward%d", i), name: fmt.Sprintf("Ward Spell %02d", i), typeLine: "Instant",
			text: "Target creature you control gains hexproof until end of turn.", identity: u, mv: 1, rank: int32(10000 + i),
			tags: []string{"protection", "heroic"}})
	}
	idx := fixture(t, pool)
	interaction := func(l *List) int {
		n := 0
		for _, c := range l.Candidates {
			if c.Role == mtgv1.CardRole_CARD_ROLE_INTERACTION {
				n++
			}
		}
		return n
	}
	req := Request{Format: cmdr, Colors: u, Theme: "heroic"}
	list, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	c, ok := find(list.Candidates, "Counterspell")
	if !ok || !c.Floor || c.Themed {
		t.Errorf("Counterspell must reach the list through the floor: %+v, %v", c, ok)
	}
	// A floor card joins after the cap, and takes no place of an on-theme
	// card. The role holds at most its cap plus its floor.
	limit := DefaultLimits.PerRole[mtgv1.CardRole_CARD_ROLE_INTERACTION]
	floor := stapleFloors(req)[mtgv1.CardRole_CARD_ROLE_INTERACTION]
	if n := interaction(list); n <= limit || n > limit+floor {
		t.Errorf("interaction holds %d cards, want over the cap %d and at most %d", n, limit, limit+floor)
	}
	req.Limits = Limits{Total: 10}
	cut, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := find(cut.Candidates, "Counterspell"); !ok {
		t.Errorf("the total cut dropped the floor: %v", names(cut.Candidates))
	}
	req.Limits = Limits{floor: map[mtgv1.CardRole]int{}}
	off, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := find(off.Candidates, "Counterspell"); ok {
		t.Error("with no floor, the cap must drop Counterspell, or the test proves nothing")
	}
}

// TestF212CommanderTriggerAddsHeroic is F-212. A commander that draws for
// each spell that targets its side pulls the targeting spells and auras
// in, though the words name only voltron. The row adds no word, so it is
// never unmatched.
func TestF212CommanderTriggerAddsHeroic(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, f212Cards())
	req := Request{Format: cmdr, Colors: []mtgv1.Color{mtgv1.Color_COLOR_U}, Theme: "voltron"}
	without, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	req.CommanderOracleIDs = []string{"kamala"}
	with, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(with.Theme.CommanderRows, []string{"heroic"}) || slices.Contains(with.Theme.Words, "heroic") {
		t.Errorf("commander rows %v, words %v, want [heroic] and no heroic word", with.Theme.CommanderRows, with.Theme.Words)
	}
	if len(with.Theme.Unmatched) > 0 {
		t.Errorf("unmatched %v, want none", with.Theme.Unmatched)
	}
	for _, name := range []string{"Arcane Flight", "Shore Up"} {
		if _, ok := find(without.Candidates, name); ok {
			t.Errorf("%s is listed with no commander, so the test proves nothing", name)
		}
		if c, ok := find(with.Candidates, name); !ok || !c.Themed {
			t.Errorf("the heroic commander must pull %s in: %+v, %v", name, c, ok)
		}
	}
}
