package candidates

import (
	"encoding/json"
	"fmt"
	"slices"
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
	partner                  mtgv1.PartnerKind
	// digital marks a card whose only printing is digital (D-306).
	digital bool
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
		// The loader derives this from the type line for real cards. The
		// fixture derives it the same way, so a commander test sees what
		// production sees.
		legendary, creature := false, false
		for _, ty := range super {
			legendary = legendary || ty == "Legendary"
		}
		for _, ty := range card {
			creature = creature || ty == "Creature"
		}
		protoCards = append(protoCards, &mtgv1.Card{
			OracleId: c.id, Name: c.name, TypeLine: c.typeLine, OracleText: c.text,
			ColorIdentity: c.identity, Keywords: c.keywords, Subtypes: c.subtypes,
			Supertypes: super, CardTypes: card, ManaValue: c.mv, EdhrecRank: c.rank,
			Legalities:     map[string]mtgv1.LegalityStatus{"commander": st, "standard": legal},
			GameChanger:    c.gameChanger,
			CanBeCommander: legendary && creature,
			Partner:        c.partner,
		})
		if c.digital {
			protoCards[len(protoCards)-1].DefaultPrinting = &mtgv1.Printing{Digital: true}
		}
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
	b, _ := New()
	cases := []struct {
		theme string
		want  []string
	}{
		{"Build me a Lifegain / Aristocrats deck, with cats!", []string{"lifegain", "aristocrats", "cats"}},
		// A fragment of "+1/+1" is not a word: the needle "1" matched
		// every card with a digit.
		{"+1/+1 counters", []string{"counters"}},
		// Two tokens that name a hyphenated row join.
		{"go wide", []string{"go-wide"}},
		{"go wide tokens", []string{"go-wide", "tokens"}},
		{"extra turns", []string{"extra-turns"}},
		{"go-wide", []string{"go-wide"}},
		// A short token that names no row goes.
		{"ub mill", []string{"mill"}},
		{"2024 zombies", []string{"zombies"}},
		// A letter outside ASCII is part of the word.
		{"nazgûl tribal", []string{"nazgûl", "tribal"}},
	}
	for _, c := range cases {
		got := b.themes.words(c.theme)
		if strings.Join(got, " ") != strings.Join(c.want, " ") {
			t.Errorf("words(%q) = %v, want %v", c.theme, got, c.want)
		}
	}
	idx := fixture(t, testCards())
	m := b.themes.match("cats zzzz", idx.Tags())
	if !contains(m.Subtypes, "Cat") {
		t.Errorf("cats should map to subtype Cat: %+v", m)
	}
	if !contains(m.Text, "zzzz") {
		t.Errorf("unknown word handling: %+v", m)
	}
	list, _ := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: "cats"})
	if _, ok := find(list.Candidates, "Ajani's Pridemate"); !ok {
		t.Errorf("subtype signal should list the Cat: %v", names(list.Candidates))
	}
}

// TestThemeUnmatchedReportsDeadWords covers the Unmatched field. The
// generic branch set hit unconditionally, so the field stayed empty even
// for a word that no card carried.
func TestThemeUnmatchedReportsDeadWords(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	list, err := b.Build(idx, Request{Format: cmdr, Theme: "cats zzzz lifegain"})
	if err != nil {
		t.Fatal(err)
	}
	if !contains(list.Theme.Unmatched, "zzzz") {
		t.Errorf("zzzz fired on no card, want it unmatched: %+v", list.Theme.Unmatched)
	}
	for _, w := range []string{"cats", "lifegain"} {
		if contains(list.Theme.Unmatched, w) {
			t.Errorf("%s fired on a card, must not be unmatched: %+v", w, list.Theme.Unmatched)
		}
	}
}

// doctorCards is the lifegain fixture plus a Doctor and a companion. The
// Doctor carries no partner kind, so canPair must read its creature
// types (CR 702.124m). The type line follows the snapshot.
func doctorCards() []tc {
	R, U := mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_U
	return append(testCards(),
		tc{id: "tenth", name: "The Tenth Doctor", typeLine: "Legendary Creature — Time Lord Doctor",
			text: "Whenever you gain life, draw a card.", identity: []mtgv1.Color{U, R}, mv: 5, rank: 800,
			subtypes: []string{"Time", "Lord", "Doctor"}, tags: []string{"lifegain"}},
		tc{id: "clara", name: "Clara Oswald", typeLine: "Legendary Creature — Human Advisor",
			text: "Doctor's companion", identity: nil, mv: 2, rank: 900,
			subtypes: []string{"Human", "Advisor"}, partner: mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION},
	)
}

// TestDoctorPairIsOffered checks that a Doctor pairs with a companion
// (D-154).
func TestDoctorPairIsOffered(t *testing.T) {
	b, _ := New()
	R, U := mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_U
	idx := fixture(t, doctorCards())
	if doctor, ok := idx.ByName("The Tenth Doctor"); !ok || !canPair(doctor) {
		t.Fatal("a Doctor must pass canPair")
	}
	pool, err := b.CommanderPool(idx, Request{Format: cmdr, Theme: "lifegain", Colors: []mtgv1.Color{U, R}, WantPair: true})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range pool {
		if c.Partner != nil && c.DisplayName() == "The Tenth Doctor + Clara Oswald" {
			return
		}
	}
	var got []string
	for _, c := range pool {
		got = append(got, c.DisplayName())
	}
	t.Errorf("no Doctor pair offered: %v", got)
}

// TestCommanderPairsSkipExcludedIds checks that an excluded commander is
// out of the pairs, as it is out of the 99.
func TestCommanderPairsSkipExcludedIds(t *testing.T) {
	b, _ := New()
	R, U := mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_U
	idx := fixture(t, doctorCards())
	pool, err := b.CommanderPool(idx, Request{Format: cmdr, Theme: "lifegain", Colors: []mtgv1.Color{U, R},
		WantPair: true, CommanderOracleIDs: []string{"tenth"}})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range pool {
		if c.Partner != nil && (c.Card.OracleId == "tenth" || c.Partner.OracleId == "tenth") {
			t.Errorf("an excluded commander is in a pair: %s", c.DisplayName())
		}
	}
}

// TestHouseFormatOffersPaperCardsOnly is D-306. A card with no paper
// printing never reaches the shortlist, and a banned card stays allowed.
func TestHouseFormatOffersPaperCardsOnly(t *testing.T) {
	b, _ := New()
	list := append(testCards(),
		tc{id: "alchemy", name: "A-Digital Rock", typeLine: "Artifact", text: "{T}: Add {W}.",
			identity: []mtgv1.Color{W}, mv: 2, rank: 40, tags: []string{"ramp"}, digital: true},
		tc{id: "bannedrock", name: "Banned Rock", typeLine: "Artifact", text: "{T}: Add {W}.",
			identity: []mtgv1.Color{W}, mv: 2, rank: 41, tags: []string{"ramp"}, commanderBanned: true},
	)
	idx := fixture(t, list)
	house, err := b.Build(idx, Request{Format: mtgv1.FormatId_FORMAT_ID_HOUSE, Colors: []mtgv1.Color{W}, Theme: "lifegain"})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := find(house.Candidates, "A-Digital Rock"); ok {
		t.Errorf("house offers a card with no paper printing: %v", names(house.Candidates))
	}
	if _, ok := find(house.Candidates, "Banned Rock"); !ok {
		t.Errorf("house dropped a banned card: %v", names(house.Candidates))
	}
	commander, _ := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: "lifegain"})
	if _, ok := find(commander.Candidates, "Banned Rock"); ok {
		t.Error("Commander offers a banned card")
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

// TestExileIsRemoval checks that a blink clause is not removal: the text
// fallback must read the whole "exile target" clause.
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

// commanderCards holds five legends and one non-legend. Three of the
// legends carry a lifegain signal, and one of those three is white-black.
// D-148 needs that one: a commander must hold every color the user named,
// so a white-black request may not take the mono-white or mono-black
// legend beside it.
func commanderCards() []tc {
	return []tc{
		{id: "heliod", name: "Heliod, Sun-Crowned", typeLine: "Legendary Enchantment Creature — God",
			text:     "Whenever you gain life, put a +1/+1 counter on target creature you control.",
			identity: []mtgv1.Color{W}, mv: 3, rank: 120, tags: []string{"lifegain"}},
		{id: "vito", name: "Vito, Thorn of the Dusk Rose", typeLine: "Legendary Creature — Vampire Cleric",
			text:     "Whenever you gain life, each opponent loses that much life.",
			identity: []mtgv1.Color{B}, mv: 3, rank: 200, tags: []string{"lifegain"}},
		// A legend with no theme signal, but a staple role and a famous
		// rank. The 99-card list keeps it. A commander list must not.
		{id: "landlegend", name: "Ojer Stand-In", typeLine: "Legendary Creature — God",
			text: "{T}: Add {W}.", identity: []mtgv1.Color{W}, mv: 2, rank: 1, tags: []string{"ramp"}},
		// The one white-black lifegain legend. The type line and the
		// identity follow the card snapshot.
		{id: "karlov", name: "Karlov of the Ghost Council", typeLine: "Legendary Creature — Spirit Advisor",
			text:     "Whenever you gain life, put two +1/+1 counters on Karlov of the Ghost Council.",
			identity: []mtgv1.Color{W, B}, mv: 2, rank: 300, tags: []string{"lifegain"}},
		{id: "offcolorlegend", name: "Green Legend", typeLine: "Legendary Creature — Elf",
			text: "Whenever you gain life, draw a card.", identity: []mtgv1.Color{G}, mv: 2, rank: 500, tags: []string{"lifegain"}},
		{id: "notlegend", name: "Soul Warden", typeLine: "Creature — Human Cleric",
			text:     "Whenever another creature enters, you gain 1 life.",
			identity: []mtgv1.Color{W}, mv: 1, rank: 5, tags: []string{"lifegain"}},
	}
}

// TestCommandersRankByTheme is D-94. The helper read the 99-card
// shortlist, which capByRole emits one role bucket at a time with lands
// first. It was never ordered by score, so the first legends it met won.
func TestCommandersRankByTheme(t *testing.T) {
	idx := fixture(t, commanderCards())
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	got, err := b.Commanders(idx, Request{Format: cmdr, Theme: "lifegain"}, 3)
	if err != nil {
		t.Fatalf("commanders: %v", err)
	}
	list := names(got)
	for _, unwanted := range []string{"Ojer Stand-In", "Soul Warden"} {
		if contains(list, unwanted) {
			t.Errorf("%q reached the commander list: %v", unwanted, list)
		}
	}
	if len(list) == 0 || list[0] != "Heliod, Sun-Crowned" {
		t.Errorf("commanders = %v, want the best lifegain legend first", list)
	}
}

// TestCommandersRespectColorIdentity keeps an illegal suggestion out. Run
// 1 of the gate offered a white-black deck three commanders outside its
// identity.
//
// D-148 tightened the test in the other direction as well. A commander
// must hold every color the user named, so the mono-white and mono-black
// legends leave a white-black list. Only Karlov is both.
func TestCommandersRespectColorIdentity(t *testing.T) {
	idx := fixture(t, commanderCards())
	b, _ := New()
	got, err := b.Commanders(idx, Request{Format: cmdr, Theme: "lifegain", Colors: []mtgv1.Color{W, B}}, 5)
	if err != nil {
		t.Fatal(err)
	}
	if contains(names(got), "Green Legend") {
		t.Errorf("a green commander reached a white-black deck: %v", names(got))
	}
	for _, partial := range []string{"Heliod, Sun-Crowned", "Vito, Thorn of the Dusk Rose"} {
		if contains(names(got), partial) {
			t.Errorf("%q holds one of the two colors, and it reached a white-black list: %v",
				partial, names(got))
		}
	}
	if len(got) != 1 || names(got)[0] != "Karlov of the Ghost Council" {
		t.Errorf("commanders = %v, want the one white-black lifegain legend", names(got))
	}
}

// TestCommanderPoolTakesEveryNamedColor is D-148. A commander must hold
// every named color: a colorless commander makes a deck that can play no
// colored card.
func TestCommanderPoolTakesEveryNamedColor(t *testing.T) {
	allowed := map[mtgv1.Color]bool{W: true, B: true}
	cases := []struct {
		name     string
		identity []mtgv1.Color
		want     bool
	}{
		{"both colors", []mtgv1.Color{W, B}, true},
		{"one of the two", []mtgv1.Color{W}, false},
		{"colorless", nil, false},
		{"both colors and colorless", []mtgv1.Color{W, B, mtgv1.Color_COLOR_C}, true},
	}
	for _, c := range cases {
		if got := identityCovers(c.identity, allowed); got != c.want {
			t.Errorf("%s: identityCovers = %v, want %v", c.name, got, c.want)
		}
	}
	// A single named color still admits the mono-colored commander.
	if !identityCovers([]mtgv1.Color{W}, map[mtgv1.Color]bool{W: true}) {
		t.Error("a mono-white request lost its mono-white commander")
	}
}

// TestCommanderPoolIsTheWeakPoolSignal is D-63. An owned mode with no
// on-theme commander is exactly what the weak-pool row asks about, and
// the count answers it with no threshold to invent.
func TestCommanderPoolIsTheWeakPoolSignal(t *testing.T) {
	idx := fixture(t, commanderCards())
	b, _ := New()
	// The collection holds one lifegain legend.
	rich, err := b.CommanderPool(idx, Request{
		Format: cmdr, Theme: "lifegain",
		PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		Owned:    map[string]int32{"vito": 1},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(rich) != 1 || rich[0].Card.Name != "Vito, Thorn of the Dusk Rose" {
		t.Errorf("owned-only pool = %v, want Vito alone", names(rich))
	}
	// The collection holds none.
	weak, err := b.CommanderPool(idx, Request{
		Format: cmdr, Theme: "lifegain",
		PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		Owned:    map[string]int32{"notlegend": 1},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(weak) != 0 {
		t.Errorf("owned-only pool = %v, want none: that is the weak-pool signal", names(weak))
	}
}

// TestCommanderPoolOwnedFirst offers what the user already has first.
func TestCommanderPoolOwnedFirst(t *testing.T) {
	idx := fixture(t, commanderCards())
	b, _ := New()
	got, err := b.CommanderPool(idx, Request{
		Format: cmdr, Theme: "lifegain",
		PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
		// Vito is owned, Heliod ranks higher but is not.
		Owned: map[string]int32{"vito": 1},
	})
	if err != nil {
		t.Fatal(err)
	}
	// The ranking wins, and the owned commander keeps its mark (D-297).
	if len(got) < 2 || got[0].Card.Name == "Vito, Thorn of the Dusk Rose" {
		t.Errorf("owned-first pool = %v, want the ranking and not the owned commander first", names(got))
	}
	for _, c := range got {
		if c.Card.Name == "Vito, Thorn of the Dusk Rose" && c.Owned == 0 {
			t.Error("the owned commander lost its mark")
		}
	}
}

// TestCommanderPoolUnsetRuleRanksOnQuality is D-293. The offer goes out
// before the pool question, and a collection must not turn it into the
// legends the user happens to own.
func TestCommanderPoolUnsetRuleRanksOnQuality(t *testing.T) {
	idx := fixture(t, commanderCards())
	b, _ := New()
	got, err := b.CommanderPool(idx, Request{
		Format: cmdr, Theme: "lifegain",
		Owned: map[string]int32{"vito": 1},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) < 2 || got[0].Card.Name == "Vito, Thorn of the Dusk Rose" {
		t.Errorf("unset-rule pool = %v, want the ranking and not the owned commander first", names(got))
	}
	for _, c := range got {
		if c.Card.Name == "Vito, Thorn of the Dusk Rose" && c.Owned == 0 {
			t.Error("the owned commander lost its owned count")
		}
	}
}

// TestCommanderQualitySnapshot is the D-94 regression gate. It needs the
// local snapshot, like TestThemeSlugsExist, so CI skips it.
//
// The old helper read the 99-card shortlist, which capByRole emits one
// role bucket at a time with lands first. A blink request came back with
// three Ojer modal double-faced cards at 0.16 on the theme, while 85
// on-theme blink commanders existed. The bar below is what that bug
// could not clear.
//
//	CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go test ./internal/candidates -run TestCommanderQualitySnapshot -v
func TestCommanderQualitySnapshot(t *testing.T) {
	idx := snapshotIndex(t)
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	// Themes the tag data covers well. A narrow theme such as extra turns
	// holds few legendary creatures that say the words, and a short list
	// is the honest answer there.
	cases := []struct {
		theme  string
		colors []mtgv1.Color
	}{
		{"sacrifice", nil},
		{"lifegain", []mtgv1.Color{W, B}},
		{"blink", []mtgv1.Color{W, mtgv1.Color_COLOR_U}},
		{"mill", []mtgv1.Color{mtgv1.Color_COLOR_U, B}},
		{"tokens", []mtgv1.Color{W, B}},
	}
	for _, tcse := range cases {
		t.Run(tcse.theme, func(t *testing.T) {
			req := Request{Format: cmdr, Theme: tcse.theme, Colors: tcse.colors}
			got, err := b.Commanders(idx, req, 3)
			if err != nil {
				t.Fatalf("commanders: %v", err)
			}
			if len(got) != 3 {
				t.Fatalf("%d commanders for a well covered theme: %v", len(got), names(got))
			}
			match := b.themes.match(tcse.theme, idx.Tags())
			for i, c := range got {
				score, _ := match.score(c.Card)
				if score <= 0 {
					t.Errorf("%s carries no theme signal", c.Card.GetName())
				}
				// The lead suggestion must be a real payoff.
				if i == 0 && score < 0.5 {
					t.Errorf("the first commander %s scores %.2f, want 0.50 or more", c.Card.GetName(), score)
				}
				if len(tcse.colors) > 0 && !IdentityFits(c.Card.ColorIdentity, colorSetOf(tcse.colors)) {
					t.Errorf("%s sits outside the requested color identity", c.Card.GetName())
				}
			}
		})
	}
}

// TestSingular covers the plural rule the generic theme branch applies to
// a subtype. Faeries and zombies end in -ie, not -y.
func TestSingular(t *testing.T) {
	cases := []struct{ in, want string }{
		{"faeries", "faerie"},
		{"zombies", "zombie"},
		{"counters", "counter"},
		{"armies", "army"},
		{"harpies", "harpy"},
		{"elves", "elf"},
		{"cats", "cat"},
		{"boss", "boss"},
		{"elf", "elf"},
	}
	for _, c := range cases {
		if got := singular(c.in); got != c.want {
			t.Errorf("singular(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

// TestTotalCapDropsTheLowestScore covers the cut under Total. The role
// caps sum to more than Total, and the old cut took the first Total cards
// in role order, so the synergy tail went and a weak land stayed.
func TestTotalCapDropsTheLowestScore(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	req := Request{Format: cmdr, Colors: []mtgv1.Color{W, B}, Theme: "lifegain"}
	full, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if len(full.Candidates) < 4 {
		t.Fatalf("too few candidates to cut: %v", names(full.Candidates))
	}
	total := len(full.Candidates) - 2
	req.Limits = Limits{Total: total}
	cut, err := b.Build(idx, req)
	if err != nil {
		t.Fatal(err)
	}
	if len(cut.Candidates) != total {
		t.Fatalf("total cap: got %d, want %d", len(cut.Candidates), total)
	}
	kept := map[string]Candidate{}
	for _, c := range cut.Candidates {
		kept[c.Card.Name] = c
	}
	// Every dropped card scores at or under every kept card.
	minKept := 1.0
	for _, c := range cut.Candidates {
		minKept = min(minKept, c.Score)
	}
	for _, c := range full.Candidates {
		if _, ok := kept[c.Card.Name]; ok {
			continue
		}
		if c.Score > minKept {
			t.Errorf("%s (%.3f) was dropped and a card at %.3f was kept", c.Card.Name, c.Score, minKept)
		}
	}
	// The role order survives the cut.
	last := -1
	for _, c := range cut.Candidates {
		pos := slices.Index(roleOrder, c.Role)
		if pos < last {
			t.Errorf("role order broken at %s", c.Card.Name)
		}
		last = pos
	}
}

// TestTextFallbacksOnlyWithoutTags covers the role fallbacks. With tags
// loaded, an untagged card that says "draw a card" is not a draw staple.
// Without tags, the text is all there is, and the fallbacks stand in.
func TestTextFallbacksOnlyWithoutTags(t *testing.T) {
	draw := &mtgv1.Card{OracleId: "d", Name: "Untagged Cantrip", OracleText: "Draw a card.", CardTypes: []string{"Instant"}}
	tagged := map[string]map[string]bool{"draw": {"other": true}}
	cases := []struct {
		name     string
		roleTags map[string]map[string]bool
		useText  bool
		want     mtgv1.CardRole
	}{
		{"tags loaded, fallbacks off", tagged, false, mtgv1.CardRole_CARD_ROLE_OTHER},
		{"no tags, fallbacks on", map[string]map[string]bool{}, true, mtgv1.CardRole_CARD_ROLE_DRAW},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _ := assignRole(draw, c.roleTags, false, c.useText)
			if got != c.want {
				t.Errorf("role = %s, want %s", got, c.want)
			}
		})
	}
	// Build reads the tag index: an index with no tags turns the
	// fallbacks on, and the fixture with tags keeps them off.
	b, _ := New()
	noTags := cards.NewIndex([]*mtgv1.Card{{
		OracleId: "d", Name: "Untagged Cantrip", OracleText: "Draw a card.", CardTypes: []string{"Instant"},
		Legalities: map[string]mtgv1.LegalityStatus{"commander": legal},
	}}, nil, nil, time.Time{})
	list, err := b.Build(noTags, Request{Format: cmdr, Theme: "lifegain"})
	if err != nil {
		t.Fatal(err)
	}
	if c, ok := find(list.Candidates, "Untagged Cantrip"); !ok || c.Role != mtgv1.CardRole_CARD_ROLE_DRAW {
		t.Errorf("without tags the cantrip must be a draw staple: %v", names(list.Candidates))
	}
	withTags := fixture(t, []tc{
		{id: "d", name: "Untagged Cantrip", typeLine: "Instant", text: "Draw a card.", rank: 10},
		{id: "t", name: "Tagged Draw", typeLine: "Instant", text: "Draw two cards.", rank: 11, tags: []string{"draw"}},
	})
	list, err = b.Build(withTags, Request{Format: cmdr, Theme: "lifegain"})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := find(list.Candidates, "Untagged Cantrip"); ok {
		t.Error("with tags loaded an untagged cantrip is not a staple")
	}
	if c, ok := find(list.Candidates, "Tagged Draw"); !ok || c.Role != mtgv1.CardRole_CARD_ROLE_DRAW {
		t.Error("the tagged draw card must keep its role")
	}
}
