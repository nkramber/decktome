package quality

import (
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
)

// TestIsCheapReadsKarstensExamples runs the cheap rule over every card
// Karsten's article of 2022-07-29 names, with the Oracle text of the card
// snapshot of 2026-09-04.
func TestIsCheapReadsKarstensExamples(t *testing.T) {
	type face struct{ typeLine, text string }
	cases := []struct {
		name     string
		mv       float64
		typeLine string
		text     string
		faces    []face
		want     bool
	}{
		{"Brainstorm", 1, "Instant", "Draw three cards, then put two cards from your hand on top of your library in any order.", nil, true},
		{"Faithless Looting", 1, "Sorcery", "Draw two cards, then discard two cards.\nFlashback {2}{R} (You may cast this card from your graveyard for its flashback cost. Then exile it.)", nil, true},
		{"Deadly Dispute", 2, "Instant", "As an additional cost to cast this spell, sacrifice an artifact or creature.\nDraw two cards and create a Treasure token. (It's an artifact with \"{T}, Sacrifice this token: Add one mana of any color.\")", nil, true},
		{"Omen of the Sea", 2, "Enchantment", "Flash (You may cast this spell any time you could cast an instant.)\nWhen this enchantment enters, scry 2, then draw a card.\n{2}{U}, Sacrifice this enchantment: Scry 2. (Look at the top two cards of your library, then put any number of them on the bottom and the rest on top in any order.)", nil, true},
		{"Growth Spiral", 2, "Instant", "Draw a card. You may put a land card from your hand onto the battlefield.", nil, true},
		{"Drannith Stinger", 2, "Creature \u2014 Human Wizard", "Whenever you cycle another card, this creature deals 1 damage to each opponent.\nCycling {1} ({1}, Discard this card: Draw a card.)", nil, true},
		{"Expressive Iteration", 2, "Sorcery", "Look at the top three cards of your library. Put one of them into your hand, put one of them on the bottom of your library, and exile one of them. You may play the exiled card this turn.", nil, true},
		{"Manamorphose", 2, "Instant", "Add two mana in any combination of colors.\nDraw a card.", nil, true},
		{"Ice-Fang Coatl", 2, "Snow Creature \u2014 Snake", "Flash\nFlying\nWhen this creature enters, draw a card.\nThis creature has deathtouch as long as you control at least three other snow permanents.", nil, true},
		{"Llanowar Elves", 1, "Creature \u2014 Elf Druid", "{T}: Add {G}.", nil, true},
		{"Skirk Prospector", 1, "Creature \u2014 Goblin", "Sacrifice a Goblin: Add {R}.", nil, true},
		{"Springleaf Drum", 1, "Artifact", "{T}, Tap an untapped creature you control: Add one mana of any color.", nil, true},
		{"Lotus Cobra", 2, "Creature \u2014 Snake", "Landfall \u2014 Whenever a land you control enters, add one mana of any color.", nil, true},
		{"Dark Ritual", 1, "Instant", "Add {B}{B}{B}.", nil, true},
		{"Sylvan Scrying", 2, "Sorcery", "Search your library for a land card, reveal it, put it into your hand, then shuffle.", nil, true},
		{"Wolfwillow Haven", 2, "Enchantment \u2014 Aura", "Enchant land\nWhenever enchanted land is tapped for mana, its controller adds an additional {G}.\n{4}{G}, Sacrifice this Aura: Create a 2/2 green Wolf creature token. Activate only during your turn.", nil, true},
		{"Aether Vial", 1, "Artifact", "At the beginning of your upkeep, you may put a charge counter on this artifact.\n{T}: You may put a creature card with mana value equal to the number of charge counters on this artifact from your hand onto the battlefield.", nil, true},
		{"Augur of Bolas", 2, "Creature \u2014 Merfolk Wizard", "When this creature enters, look at the top three cards of your library. You may reveal an instant or sorcery card from among them and put it into your hand. Put the rest on the bottom of your library in any order.", nil, false},
		{"Trail of Crumbs", 2, "Enchantment", "When this enchantment enters, create a Food token. (It's an artifact with \"{2}, {T}, Sacrifice this token: You gain 3 life.\")\nWhenever you sacrifice a Food, you may pay {1}. If you do, look at the top two cards of your library. You may reveal a permanent card from among them and put it into your hand. Put the rest on the bottom of your library in any order.", nil, false},
		{"Esper Sentinel", 1, "Artifact Creature \u2014 Human Soldier", "Whenever an opponent casts their first noncreature spell each turn, draw a card unless that player pays {X}, where X is this creature's power.", nil, false},
		{"Fateful Absence", 2, "Instant", "Destroy target creature or planeswalker. Its controller investigates. (Create a Clue token. It's an artifact with \"{2}, Sacrifice this token: Draw a card.\")", nil, false},
		{"Ledger Shredder", 2, "Creature \u2014 Bird Advisor", "Flying\nWhenever a player casts their second spell each turn, this creature connives. (Draw a card, then discard a card. If you discarded a nonland card, put a +1/+1 counter on this creature.)", nil, false},
		{"Edgewall Innkeeper", 1, "Creature \u2014 Human Peasant", "Whenever you cast a creature spell that has an Adventure, draw a card. (It doesn't need to have gone on the adventure first.)", nil, false},
		{"Improbable Alliance", 2, "Enchantment", "Whenever you draw your second card each turn, create a 1/1 blue Faerie creature token with flying.\n{4}{U}{R}: Draw a card, then discard a card.", nil, false},
		{"Ox of Agonas", 5, "Creature \u2014 Ox", "When this creature enters, discard your hand, then draw three cards.\nEscape\u2014{R}{R}, Exile eight other cards from your graveyard. (You may cast this card from your graveyard for its escape cost.)\nThis creature escapes with a +1/+1 counter on it.", nil, false},
		{"Bloodtithe Harvester", 2, "Creature \u2014 Vampire", "When this creature enters, create a Blood token. (It's an artifact with \"{1}, {T}, Discard a card, Sacrifice this token: Draw a card.\")\n{T}, Sacrifice this creature: Target creature gets -X/-X until end of turn, where X is twice the number of Blood tokens you control. Activate only as a sorcery.", nil, false},
		{"Shark Typhoon", 6, "Enchantment", "Whenever you cast a noncreature spell, create an X/X blue Shark creature token with flying, where X is that spell's mana value.\nCycling {X}{1}{U} ({X}{1}{U}, Discard this card: Draw a card.)\nWhen you cycle this card, create an X/X blue Shark creature token with flying.", nil, false},
		{"Hydroid Krasis", 2, "Creature \u2014 Jellyfish Hydra Beast", "When you cast this spell, you gain half X life and draw half X cards. Round down each time.\nFlying, trample\nThis creature enters with X +1/+1 counters on it.", nil, false},
		{"Ravenous Squirrel", 1, "Creature \u2014 Squirrel", "Whenever you sacrifice an artifact or creature, put a +1/+1 counter on this creature.\n{1}{B}{G}, Sacrifice an artifact or creature: You gain 1 life and draw a card.", nil, false},
		{"Ranger Class", 2, "Enchantment \u2014 Class", "(Gain the next level as a sorcery to add its ability.)\nWhen this Class enters, create a 2/2 green Wolf creature token.\n{1}{G}: Level 2\nWhenever you attack, put a +1/+1 counter on target attacking creature.\n{3}{G}: Level 3\nYou may look at the top card of your library any time.\nYou may cast creature spells from the top of your library.", nil, false},
		{"Tangled Florahedron", 2, "Creature \u2014 Elemental // Land", "", []face{{"Creature \u2014 Elemental", "{T}: Add {G}."}, {"Land", "This land enters tapped.\n{T}: Add {G}."}}, false},
		{"Shambling Ghast", 1, "Creature \u2014 Zombie", "When this creature dies, choose one \u2014\n\u2022 Target creature an opponent controls gets -1/-1 until end of turn.\n\u2022 Create a Treasure token. (It's an artifact with \"{T}, Sacrifice this token: Add one mana of any color.\")", nil, false},
		{"Crop Rotation", 1, "Instant", "As an additional cost to cast this spell, sacrifice a land.\nSearch your library for a land card, put that card onto the battlefield, then shuffle.", nil, false},
	}
	for _, c := range cases {
		card := &mtgv1.Card{Name: c.name, ManaValue: c.mv, TypeLine: c.typeLine, OracleText: c.text}
		for _, f := range c.faces {
			card.Faces = append(card.Faces, &mtgv1.CardFace{TypeLine: f.typeLine, OracleText: f.text})
		}
		if got := isCheap(card); got != c.want {
			t.Errorf("%s: cheap = %v, and the article says %v", c.name, got, c.want)
		}
	}
}

// commanderWorld is a Commander pool: a white-blue commander, a colorless
// commander, spells at three and at six mana, colorless artifacts, a
// cheap draw spell, and basics.
type commanderWorld struct {
	idx       *cards.Index
	prof      *profile.Profiler
	commander *mtgv1.Card
	colorless *mtgv1.Card
	threes    []*mtgv1.Card
	sixes     []*mtgv1.Card
	grays     []*mtgv1.Card
	cantrip   *mtgv1.Card
	plains    *mtgv1.Card
	island    *mtgv1.Card
	wastes    *mtgv1.Card
}

func newCommanderWorld(t *testing.T) *commanderWorld {
	t.Helper()
	w := &commanderWorld{}
	w.commander = spell("Test Commander", 4, W)
	w.commander.ManaCost, w.commander.Supertypes = "{2}{W}{U}", []string{"Legendary"}
	w.commander.Colors, w.commander.ColorIdentity = []mtgv1.Color{W, U}, []mtgv1.Color{W, U}
	w.colorless = &mtgv1.Card{OracleId: "oid-colorless-commander", Name: "Colorless Commander", ManaValue: 4, ManaCost: "{4}",
		CardTypes: []string{"Artifact", "Creature"}, Supertypes: []string{"Legendary"}}
	all := []*mtgv1.Card{w.commander, w.colorless}
	for i := 0; i < 90; i++ {
		color := W
		if i%2 == 1 {
			color = U
		}
		c := spell(strings.Repeat("Three ", 1)+string(rune('A'+i%26))+string(rune('a'+i/26)), 3, color)
		w.threes = append(w.threes, c)
		all = append(all, c)
	}
	for i := 0; i < 70; i++ {
		color := W
		if i%2 == 1 {
			color = U
		}
		c := spell("Six "+string(rune('A'+i%26))+string(rune('a'+i/26)), 6, color)
		w.sixes = append(w.sixes, c)
		all = append(all, c)
	}
	for i := 0; i < 70; i++ {
		c := &mtgv1.Card{OracleId: "oid-gray-" + string(rune('a'+i%26)) + string(rune('a'+i/26)), Name: "Gray " + string(rune('A'+i%26)) + string(rune('a'+i/26)),
			ManaValue: 3, ManaCost: "{3}", CardTypes: []string{"Artifact"}}
		w.grays = append(w.grays, c)
		all = append(all, c)
	}
	w.cantrip = &mtgv1.Card{OracleId: "oid-test-cantrip", Name: "Test Cantrip", ManaValue: 1, ManaCost: "{U}", CardTypes: []string{"Instant"},
		TypeLine: "Instant", OracleText: "Draw a card.", Colors: []mtgv1.Color{U}, ColorIdentity: []mtgv1.Color{U}}
	w.plains, w.island = land("Plains", true, W), land("Island", true, U)
	w.wastes = land("Wastes", true, mtgv1.Color_COLOR_C)
	all = append(all, w.cantrip, w.plains, w.island, w.wastes)
	w.idx = cards.NewIndex(all, nil, nil, time.Date(2026, 9, 11, 0, 0, 0, 0, time.UTC))
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	if w.prof, err = profile.New(cfg, nil, nil); err != nil {
		t.Fatal(err)
	}
	return w
}

// deck builds a deck of one copy of each spell and the named basics.
func (w *commanderWorld) deck(format mtgv1.FormatId, commander *mtgv1.Card, spells []*mtgv1.Card, basics map[*mtgv1.Card]int32) *mtgv1.Deck {
	d := &mtgv1.Deck{Format: &mtgv1.Format{Id: format}}
	if commander != nil {
		d.CommanderOracleIds = []string{commander.GetOracleId()}
	}
	for _, s := range spells {
		d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: s.GetOracleId(), Name: s.GetName(), Count: 1})
	}
	for b, n := range basics {
		d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: b.GetOracleId(), Name: b.GetName(), Count: n})
	}
	return d
}

func (w *commanderWorld) input(d *mtgv1.Deck) Input {
	return Input{Deck: d, Profile: w.prof.Measure(d, w.idx), Cards: w.idx}
}

func (w *commanderWorld) sound() *mtgv1.Deck {
	return w.deck(mtgv1.FormatId_FORMAT_ID_COMMANDER, w.commander, append(append([]*mtgv1.Card{}, w.threes[:62]...), w.cantrip),
		map[*mtgv1.Card]int32{w.plains: 18, w.island: 18})
}

func (w *commanderWorld) short() *mtgv1.Deck {
	return w.deck(mtgv1.FormatId_FORMAT_ID_COMMANDER, w.commander, append(append([]*mtgv1.Card{}, w.threes[:74]...), w.cantrip),
		map[*mtgv1.Card]int32{w.plains: 12, w.island: 12})
}

// TestRuleChecks reads each check on a Commander deck, a colorless deck,
// and a deck of another format (D-667, D-678).
func TestRuleChecks(t *testing.T) {
	w := newCommanderWorld(t)

	sound := ruleChecks(w.input(w.sound()))
	if !sound.Checked || sound.Flagged() {
		t.Errorf("a sound deck: %+v", sound)
	}
	if sound.Cheap != 1 || sound.LandCount != 36 || !sound.HasColors {
		t.Errorf("a sound deck reads %d cheap spells, %.0f lands, colors %v, want 1, 36, true", sound.Cheap, sound.LandCount, sound.HasColors)
	}
	if want := karstenLandBase + karstenLandPerMV*sound.AvgManaValue - karstenLandPerCheap; sound.Need != want {
		t.Errorf("need = %.3f, want %.3f", sound.Need, want)
	}

	short := ruleChecks(w.input(w.short()))
	if !short.Lands || short.Shortfall() < RuleLandShortfall {
		t.Errorf("a deck of 24 lands: %+v, shortfall %.1f", short, short.Shortfall())
	}

	curve := ruleChecks(w.input(w.deck(mtgv1.FormatId_FORMAT_ID_COMMANDER, w.commander, w.sixes[:63],
		map[*mtgv1.Card]int32{w.plains: 18, w.island: 18})))
	if !curve.Curve || curve.AvgManaValue <= RuleCurveCeiling {
		t.Errorf("a deck of six-drops: %+v", curve)
	}

	colors := ruleChecks(w.input(w.deck(mtgv1.FormatId_FORMAT_ID_COMMANDER, w.commander, append(append([]*mtgv1.Card{}, w.threes[:62]...), w.cantrip),
		map[*mtgv1.Card]int32{w.plains: 36})))
	if !colors.Colors || colors.Lands || colors.WorstColor != U || colors.WorstHave != 0 {
		t.Errorf("a deck of Plains under blue spells: %+v", colors)
	}
	if got := colors.reasons(); len(got) != 1 || !strings.Contains(got[0], "blue sources cover 0 of the") {
		t.Errorf("reasons = %q", got)
	}

	gray := ruleChecks(w.input(w.deck(mtgv1.FormatId_FORMAT_ID_COMMANDER, w.colorless, w.grays[:63],
		map[*mtgv1.Card]int32{w.wastes: 36})))
	if !gray.Checked || gray.HasColors || gray.Colors {
		t.Errorf("a colorless deck reads a colors check: %+v", gray)
	}

	standard := ruleChecks(w.input(w.deck(mtgv1.FormatId_FORMAT_ID_STANDARD, nil, w.threes[:36],
		map[*mtgv1.Card]int32{w.plains: 12, w.island: 12})))
	if standard.Checked || standard.Flagged() {
		t.Errorf("a Standard deck reads the checks: %+v", standard)
	}
}

// TestCommanderTierReadsTheRules: the detector of this model flags every
// deck, and its ladder answers typical. A Commander deck the rules pass
// takes the ladder's tier, and the detector keeps the score (F-115,
// D-678). A deck the rules flag grades bad and names the check. A deck
// of another format keeps the grade of the detector (D-667).
func TestCommanderTierReadsTheRules(t *testing.T) {
	w := newCommanderWorld(t)
	ladder := func(word string) *FormatModel {
		return &FormatModel{Format: word, Tiers: []string{meta.TierBad, meta.TierBaseline, meta.TierTypical, meta.TierGood, meta.TierGreat},
			Thresholds: []float64{-3, -2, 2, 3}, DefectBias: 5, DefectThreshold: 0.5}
	}
	s := NewScorer(&Model{Version: "test", Formats: map[string]*FormatModel{
		meta.FormatCommander: ladder(meta.FormatCommander), meta.FormatStandard: ladder(meta.FormatStandard),
	}})

	sound := s.Score(w.input(w.sound()))
	if sound.GetTier() != meta.TierTypical {
		t.Errorf("a Commander deck the rules pass grades %s, want typical", sound.GetTier())
	}
	if sound.GetScore() >= DefectFloor {
		t.Errorf("the detector keeps the score: %.3f, want under %.2f", sound.GetScore(), DefectFloor)
	}
	if len(sound.GetReasons()) > ReasonCount {
		t.Errorf("reasons = %q", sound.GetReasons())
	}
	e := s.Explain(w.input(w.sound()))
	if e.Tier != meta.TierTypical || !e.Flagged || !e.Rules.Checked || e.Rules.Flagged() {
		t.Errorf("explain = %+v", e)
	}

	short := s.Score(w.input(w.short()))
	if short.GetTier() != meta.TierBad {
		t.Errorf("a deck of 24 lands grades %s, want bad", short.GetTier())
	}
	if p := short.GetProbabilities(); len(p) != 5 || p[0].GetProbability() != 1 {
		t.Errorf("probabilities = %v", p)
	}
	if len(short.GetReasons()) == 0 || !strings.Contains(short.GetReasons()[0], "holds 24 lands") {
		t.Errorf("reasons = %q", short.GetReasons())
	}
	if !strings.Contains(Summary(short), "below the precon baseline") {
		t.Errorf("summary = %q", Summary(short))
	}
	flagged := s.Explain(w.input(w.short()))
	if flagged.Tier != meta.TierBad || !flagged.Rules.Flagged() || argmax(flagged.Ladder) != 2 {
		t.Errorf("explain of a flagged deck: tier %s, ladder %v, want bad over a typical ladder", flagged.Tier, flagged.Ladder)
	}

	standard := s.Score(w.input(w.deck(mtgv1.FormatId_FORMAT_ID_STANDARD, nil, w.threes[:36],
		map[*mtgv1.Card]int32{w.plains: 12, w.island: 12})))
	if standard.GetTier() != meta.TierBad {
		t.Errorf("a Standard deck the detector flags grades %s, want bad", standard.GetTier())
	}
}
