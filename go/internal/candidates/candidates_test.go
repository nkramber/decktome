package candidates

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

const (
	legal  = mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL
	banned = mtgv1.LegalityStatus_LEGALITY_STATUS_BANNED
	cmdr   = mtgv1.FormatId_FORMAT_ID_COMMANDER
)

var (
	W = mtgv1.Color_COLOR_W
	B = mtgv1.Color_COLOR_B
	G = mtgv1.Color_COLOR_G
)

type tc struct {
	id, name, typeLine, text string
	identity                 []mtgv1.Color
	keywords, subtypes       []string
	mv                       float64
	rank                     int32
	tags                     []string
	commanderBanned          bool
	gameChanger              bool
}

// fixture builds an index with a tag file from the test cards.
func fixture(t *testing.T, list []tc) *cards.Index {
	t.Helper()
	var protoCards []*mtgv1.Card
	tagCards := map[string][]string{}
	for _, c := range list {
		types := strings.Split(strings.Split(c.typeLine, " — ")[0], " ")
		var super, card []string
		for _, ty := range types {
			if ty == "Basic" || ty == "Legendary" {
				super = append(super, ty)
			} else {
				card = append(card, ty)
			}
		}
		st := mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL
		if c.commanderBanned {
			st = banned
		}
		protoCards = append(protoCards, &mtgv1.Card{
			OracleId: c.id, Name: c.name, TypeLine: c.typeLine, OracleText: c.text,
			ColorIdentity: c.identity, Keywords: c.keywords, Subtypes: c.subtypes,
			Supertypes: super, CardTypes: card, ManaValue: c.mv, EdhrecRank: c.rank,
			Legalities:  map[string]mtgv1.LegalityStatus{"commander": st, "standard": legal},
			GameChanger: c.gameChanger,
		})
		for _, tg := range c.tags {
			tagCards[tg] = append(tagCards[tg], c.id)
		}
	}
	var sb strings.Builder
	n := 0
	for slug, ids := range tagCards {
		n++
		var tg []map[string]string
		for _, id := range ids {
			tg = append(tg, map[string]string{"oracle_id": id})
		}
		line, _ := json.Marshal(map[string]any{"id": fmt.Sprintf("t%d", n), "slug": slug, "taggings": tg})
		sb.Write(line)
		sb.WriteString("\n")
	}
	tags, err := cards.LoadTags(strings.NewReader(sb.String()), "tags")
	if err != nil {
		t.Fatal(err)
	}
	return cards.NewIndex(protoCards, nil, tags, time.Date(2026, 8, 24, 0, 0, 0, 0, time.UTC))
}

func testCards() []tc {
	return []tc{
		{id: "plains", name: "Plains", typeLine: "Basic Land — Plains", identity: nil, rank: 1},
		{id: "soulwarden", name: "Soul Warden", typeLine: "Creature — Human Cleric", text: "Whenever another creature enters, you gain 1 life.", identity: []mtgv1.Color{W}, mv: 1, rank: 200, tags: []string{"lifegain"}},
		{id: "pridemate", name: "Ajani's Pridemate", typeLine: "Creature — Cat Soldier", text: "Whenever you gain life, put a +1/+1 counter on Ajani's Pridemate.", identity: []mtgv1.Color{W}, mv: 2, rank: 300, tags: []string{"lifegain"}, subtypes: []string{"Cat", "Soldier"}},
		{id: "lifelinker", name: "Lone Rider", typeLine: "Creature — Human Knight", text: "First strike, lifelink", identity: []mtgv1.Color{W}, keywords: []string{"Lifelink", "First strike"}, mv: 2, rank: 5000},
		{id: "angel", name: "Resplendent Angel", typeLine: "Creature — Angel", text: "Flying. At the beginning of each end step, if you gained 5 or more life this turn, create a 4/4 white Angel creature token.", identity: []mtgv1.Color{W}, mv: 3, rank: 400, tags: []string{"lifegain"}, keywords: []string{"Flying"}, subtypes: []string{"Angel"}},
		{id: "bigangel", name: "Archangel of Thune", typeLine: "Creature — Angel", text: "Flying, lifelink. Whenever you gain life, put a +1/+1 counter on each creature you control.", identity: []mtgv1.Color{W}, mv: 5, rank: 150, tags: []string{"lifegain"}, keywords: []string{"Flying", "Lifelink"}, subtypes: []string{"Angel"}},
		{id: "solring", name: "Sol Ring", typeLine: "Artifact", text: "{T}: Add {C}{C}.", identity: nil, mv: 1, rank: 1, tags: []string{"ramp", "mana-rock"}},
		{id: "swords", name: "Swords to Plowshares", typeLine: "Instant", text: "Exile target creature. Its controller gains life equal to its power.", identity: []mtgv1.Color{W}, mv: 1, rank: 10, tags: []string{"removal"}},
		{id: "wrath", name: "Wrath of God", typeLine: "Sorcery", text: "Destroy all creatures. They can't be regenerated.", identity: []mtgv1.Color{W}, mv: 4, rank: 500, tags: []string{"sweeper", "removal"}},
		{id: "draw", name: "Well of Lost Dreams", typeLine: "Artifact", text: "Whenever you gain life, you may pay {X}. If you do, draw X cards.", identity: nil, mv: 4, rank: 900, tags: []string{"draw", "lifegain"}},
		{id: "counter", name: "Counterspell", typeLine: "Instant", text: "Counter target spell.", identity: []mtgv1.Color{mtgv1.Color_COLOR_U}, mv: 2, rank: 20, tags: []string{"counterspell"}},
		{id: "bloodartist", name: "Blood Artist", typeLine: "Creature — Vampire", text: "Whenever Blood Artist or another creature dies, target player loses 1 life and you gain 1 life.", identity: []mtgv1.Color{B}, mv: 2, rank: 250, tags: []string{"death-trigger", "lifegain"}},
		{id: "banned", name: "Griselbrand", typeLine: "Legendary Creature — Demon", text: "Flying, lifelink", identity: []mtgv1.Color{B}, keywords: []string{"Flying", "Lifelink"}, mv: 8, rank: 700, commanderBanned: true},
		{id: "gc", name: "Smothering Tithe", typeLine: "Enchantment", text: "Whenever an opponent draws a card, that player may pay {2}. If they don't, you create a Treasure token.", identity: []mtgv1.Color{W}, mv: 4, rank: 30, gameChanger: true, tags: []string{"ramp"}},
		{id: "offcolor", name: "Llanowar Elves", typeLine: "Creature — Elf Druid", text: "{T}: Add {G}.", identity: []mtgv1.Color{G}, mv: 1, rank: 100, tags: []string{"ramp", "mana-dork"}},
		{id: "vanilla", name: "Grizzly Bears", typeLine: "Creature — Bear", text: "", identity: []mtgv1.Color{G}, mv: 2, rank: 20000},
		{id: "cmdr", name: "Heliod, Sun-Crowned", typeLine: "Legendary Enchantment Creature — God", text: "Whenever you gain life, put a +1/+1 counter on target creature or enchantment you control.", identity: []mtgv1.Color{W}, mv: 3, rank: 120, tags: []string{"lifegain"}},
		{id: "land", name: "Command Tower", typeLine: "Land", text: "{T}: Add one mana of any color in your commander's color identity.", identity: nil, rank: 2, tags: []string{"utility-land"}},
	}
}

func names(cs []Candidate) []string {
	var out []string
	for _, c := range cs {
		out = append(out, c.Card.Name)
	}
	return out
}

func find(cs []Candidate, name string) (Candidate, bool) {
	for _, c := range cs {
		if c.Card.Name == name {
			return c, true
		}
	}
	return Candidate{}, false
}

func TestBuildAnyCardLifegain(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, testCards())
	list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "Build me a lifegain deck", CommanderOracleIDs: []string{"cmdr"}})
	if err != nil {
		t.Fatal(err)
	}
	got := names(list.Candidates)
	for _, want := range []string{"Soul Warden", "Ajani's Pridemate", "Archangel of Thune", "Sol Ring", "Swords to Plowshares", "Wrath of God", "Well of Lost Dreams", "Blood Artist", "Command Tower", "Lone Rider"} {
		if _, ok := find(list.Candidates, want); !ok {
			t.Errorf("missing %s in %v", want, got)
		}
	}
	for _, bad := range []string{"Plains", "Heliod, Sun-Crowned", "Griselbrand", "Llanowar Elves", "Grizzly Bears", "Counterspell"} {
		if _, ok := find(list.Candidates, bad); ok {
			t.Errorf("must not list %s", bad)
		}
	}
	if list.Upgrades != nil {
		t.Errorf("any-card mode has no upgrades, got %v", names(list.Upgrades))
	}
	// Roles.
	roles := map[string]mtgv1.CardRole{
		"Command Tower": mtgv1.CardRole_CARD_ROLE_LAND, "Sol Ring": mtgv1.CardRole_CARD_ROLE_RAMP,
		"Swords to Plowshares": mtgv1.CardRole_CARD_ROLE_REMOVAL, "Wrath of God": mtgv1.CardRole_CARD_ROLE_WIPE,
		"Well of Lost Dreams": mtgv1.CardRole_CARD_ROLE_DRAW, "Archangel of Thune": mtgv1.CardRole_CARD_ROLE_THREAT,
		"Soul Warden": mtgv1.CardRole_CARD_ROLE_SYNERGY,
	}
	for name, want := range roles {
		c, _ := find(list.Candidates, name)
		if c.Role != want {
			t.Errorf("%s role = %s, want %s", name, c.Role, want)
		}
	}
	// Theme signals.
	sw, _ := find(list.Candidates, "Soul Warden")
	if !contains(sw.Signals, "tag:lifegain") {
		t.Errorf("Soul Warden signals = %v", sw.Signals)
	}
	lr, _ := find(list.Candidates, "Lone Rider")
	if !contains(lr.Signals, "keyword:Lifelink") {
		t.Errorf("Lone Rider signals = %v", lr.Signals)
	}
	// Order inside a role: on-theme, popular first.
	if got[0] != "Command Tower" {
		t.Errorf("first candidate = %s, want the land group first", got[0])
	}
	if list.Stats.Pool == 0 || list.Stats.OnTheme < 5 || list.Stats.Returned != len(got) {
		t.Errorf("stats = %+v", list.Stats)
	}
	if list.Theme.Describe() == "no theme signal" {
		t.Error("theme should resolve")
	}
}

func TestBuildBracketDropsGameChangers(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	for _, tt := range []struct {
		bracket int32
		want    bool
	}{{0, true}, {1, false}, {2, false}, {3, true}, {4, true}} {
		list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: "lifegain", Bracket: tt.bracket})
		if err != nil {
			t.Fatal(err)
		}
		_, ok := find(list.Candidates, "Smothering Tithe")
		if ok != tt.want {
			t.Errorf("bracket %d: Smothering Tithe listed = %v, want %v", tt.bracket, ok, tt.want)
		}
	}
}

func TestBuildOwnedModes(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	owned := map[string]int32{"soulwarden": 1, "solring": 1, "swords": 2, "land": 1, "lifelinker": 1}
	req := Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "lifegain", Owned: owned}

	// Default with a collection is owned-first.
	list, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range list.Candidates {
		if c.Owned == 0 {
			t.Errorf("owned-first main list has unowned %s", c.Card.Name)
		}
	}
	if len(list.Upgrades) == 0 {
		t.Fatal("owned-first must list upgrades")
	}
	for _, c := range list.Upgrades {
		if c.Owned != 0 {
			t.Errorf("upgrade %s is owned", c.Card.Name)
		}
	}
	if _, ok := find(list.Upgrades, "Archangel of Thune"); !ok {
		t.Errorf("upgrades = %v, want Archangel of Thune", names(list.Upgrades))
	}
	// An unowned card weaker than the weakest owned card of its role is
	// not an upgrade: Lone Rider (owned, synergy, weak) sets the floor,
	// Soul Warden is owned. Blood Artist beats Lone Rider, so it is one.
	if _, ok := find(list.Upgrades, "Blood Artist"); !ok {
		t.Errorf("upgrades = %v, want Blood Artist", names(list.Upgrades))
	}

	req.PoolRule = mtgv1.PoolRule_POOL_RULE_OWNED_ONLY
	list, _ = b.Build(idx, req)
	if len(list.Upgrades) != 0 || len(list.Candidates) != 5 {
		t.Errorf("owned-only: %d candidates, %d upgrades", len(list.Candidates), len(list.Upgrades))
	}

	req.PoolRule = mtgv1.PoolRule_POOL_RULE_ANY_CARD
	list, _ = b.Build(idx, req)
	sw, _ := find(list.Candidates, "Soul Warden")
	if sw.Owned != 1 {
		t.Error("any-card keeps the owned mark as information (D-37)")
	}
	if _, ok := find(list.Candidates, "Archangel of Thune"); !ok {
		t.Error("any-card lists unowned cards")
	}

	req.Owned = nil
	req.PoolRule = mtgv1.PoolRule_POOL_RULE_OWNED_ONLY
	if _, err := b.Build(idx, req); err == nil {
		t.Error("owned-only without a collection must fail")
	}
}

func TestBuildLimits(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	list, err := b.Build(idx, Request{Format: cmdr, Theme: "lifegain", Limits: Limits{Total: 3}})
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Candidates) != 3 {
		t.Errorf("total cap: got %d", len(list.Candidates))
	}
	list, _ = b.Build(idx, Request{Format: cmdr, Theme: "lifegain", Limits: Limits{PerRole: map[mtgv1.CardRole]int{mtgv1.CardRole_CARD_ROLE_SYNERGY: 1}}})
	n := 0
	for _, c := range list.Candidates {
		if c.Role == mtgv1.CardRole_CARD_ROLE_SYNERGY {
			n++
		}
	}
	if n != 1 {
		t.Errorf("per-role cap: %d synergy cards", n)
	}
}

func TestBuildDeterministic(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	req := Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "lifegain tokens"}
	a, _ := b.Build(idx, req)
	c, _ := b.Build(idx, req)
	if strings.Join(names(a.Candidates), ",") != strings.Join(names(c.Candidates), ",") {
		t.Error("two builds differ")
	}
}

func TestBuildNoThemeKeepsStaples(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: ""})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Sol Ring", "Swords to Plowshares", "Wrath of God", "Command Tower"} {
		if _, ok := find(list.Candidates, want); !ok {
			t.Errorf("staple %s missing", want)
		}
	}
	if _, ok := find(list.Candidates, "Ajani's Pridemate"); ok {
		t.Error("a synergy card needs a theme")
	}
}

func TestThemeWords(t *testing.T) {
	got := words("Build me a Lifegain / Aristocrats deck, with cats!")
	want := []string{"lifegain", "aristocrats", "cats"}
	if strings.Join(got, " ") != strings.Join(want, " ") {
		t.Errorf("words = %v, want %v", got, want)
	}
	b, _ := New()
	idx := fixture(t, testCards())
	m := b.themes.match("cats zzzz", idx.Tags())
	if !contains(m.Subtypes, "Cat") {
		t.Errorf("cats should map to subtype Cat: %+v", m)
	}
	if !contains(m.Unmatched, "zzzz") && !contains(m.Text, "zzzz") {
		t.Errorf("unknown word handling: %+v", m)
	}
	list, _ := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: "cats"})
	if _, ok := find(list.Candidates, "Ajani's Pridemate"); !ok {
		t.Errorf("subtype signal should list the Cat: %v", names(list.Candidates))
	}
}

func TestBuildNilIndex(t *testing.T) {
	b, _ := New()
	if _, err := b.Build(nil, Request{}); err == nil {
		t.Error("nil index must fail")
	}
}

func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}

func TestPayoffOutranksEnabler(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: "lifegain"})
	if err != nil {
		t.Fatal(err)
	}
	pride, _ := find(list.Candidates, "Ajani's Pridemate")
	rider, _ := find(list.Candidates, "Lone Rider")
	if pride.Score <= rider.Score {
		t.Errorf("payoff Pridemate %.2f must beat enabler Lone Rider %.2f", pride.Score, rider.Score)
	}
	if !contains(pride.Signals, "payoff-text:whenever you gain life") {
		t.Errorf("Pridemate signals = %v", pride.Signals)
	}
	// B: the Lifelink keyword covers the "lifelink" text needle, so a
	// lifelinker gets one signal for that fact, not two.
	for _, sig := range rider.Signals {
		if sig == "text:lifelink" {
			t.Errorf("Lone Rider double counts lifelink: %v", rider.Signals)
		}
	}
	if !contains(rider.Signals, "keyword:Lifelink") {
		t.Errorf("Lone Rider signals = %v", rider.Signals)
	}
}

func TestThinTheme(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	owned := map[string]int32{"soulwarden": 1, "solring": 1}
	list, _ := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: "lifegain", Owned: owned})
	if !list.Stats.ThinTheme || list.Stats.OnThemeOwned != 1 {
		t.Errorf("owned-first with one on-theme card: stats = %+v", list.Stats)
	}
	list, _ = b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: "lifegain", Owned: owned, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD})
	if list.Stats.ThinTheme {
		t.Error("any-card mode never reports a thin theme")
	}
}

// TestExileIsRemoval guards defect C of the PR-6 gate (2026-08-24). The
// text fallback called every "exile target" clause removal, so blink
// spells got role removal and the signal text:destroy target.
func TestExileIsRemoval(t *testing.T) {
	cases := []struct {
		name string
		text string
		want bool
	}{
		{"blink", "exile target creature you control, then return it to the battlefield under its owner's control.", false},
		{"blink another", "{4}{w}, {t}: exile another target creature you control, then return it to the battlefield under its owner's control.", false},
		{"blink up to one", "when this creature enters, exile up to one other target creature or artifact you control. return it to the battlefield under its owner's control at the beginning of the next end step.", false},
		{"exile removal", "exile target creature. its controller gains 2 life.", true},
		{"exile opponent creature", "exile target creature an opponent controls.", true},
		{"no exile", "destroy all creatures.", false},
	}
	for _, c := range cases {
		if got := exileIsRemoval(c.text); got != c.want {
			t.Errorf("%s: exileIsRemoval = %v, want %v", c.name, got, c.want)
		}
	}
}
