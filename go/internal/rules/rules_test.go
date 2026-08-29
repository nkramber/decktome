package rules

import (
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

var (
	testIndex *cards.Index
	testCfg   *Config
)

// loadFixture reads one Scryfall JSONL fixture file.
func loadFixture(path string) []*mtgv1.Card {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	cardList, err := cards.LoadCards(f, path)
	_ = f.Close()
	if err != nil {
		panic(err)
	}
	return cardList
}

// TestMain loads the shared cards fixture plus the rules-only rows in
// testdata/cards_extra.jsonl (Scryfall API rows).
func TestMain(m *testing.M) {
	cardList := loadFixture("../cards/testdata/cards_fixture.jsonl")
	cardList = append(cardList, loadFixture("testdata/cards_extra.jsonl")...)
	testIndex = cards.NewIndex(cardList, nil, nil, time.Now())
	var err error
	testCfg, err = Load()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

// oid resolves a fixture card name to its Oracle id.
func oid(t *testing.T, name string) string {
	t.Helper()
	c, ok := testIndex.ByName(name)
	if !ok {
		t.Fatalf("fixture card missing: %q", name)
	}
	return c.OracleId
}

// deckSpec builds decks compactly. Counts of zero are invalid on purpose.
type deckSpec struct {
	format     mtgv1.FormatId
	bracket    int32
	commanders []string
	companion  string
	cards      map[string]int32
	sideboard  map[string]int32
	fill       string // basic land name to fill the main deck with
	fillTo     int32
}

func (ds deckSpec) build(t *testing.T) *mtgv1.Deck {
	t.Helper()
	deck := &mtgv1.Deck{Format: &mtgv1.Format{Id: ds.format}}
	if ds.bracket > 0 {
		deck.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: ds.bracket}}
	}
	var n int32
	for _, name := range ds.commanders {
		deck.CommanderOracleIds = append(deck.CommanderOracleIds, oid(t, name))
		n++
	}
	if ds.companion != "" {
		deck.CompanionOracleId = oid(t, ds.companion)
	}
	for name, count := range ds.cards {
		if count <= 0 {
			continue
		}
		deck.Cards = append(deck.Cards, &mtgv1.DeckCard{OracleId: oid(t, name), Name: name, Count: count})
		n += count
	}
	for name, count := range ds.sideboard {
		deck.Sideboard = append(deck.Sideboard, &mtgv1.DeckCard{OracleId: oid(t, name), Name: name, Count: count})
	}
	if ds.fill != "" && n < ds.fillTo {
		deck.Cards = append(deck.Cards, &mtgv1.DeckCard{OracleId: oid(t, ds.fill), Name: ds.fill, Count: ds.fillTo - n})
	}
	return deck
}

func validate(t *testing.T, ds deckSpec) *mtgv1.ValidationResult {
	t.Helper()
	return testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex})
}

func codes(res *mtgv1.ValidationResult, sev mtgv1.Severity) map[string]int {
	out := map[string]int{}
	for _, f := range res.Findings {
		if f.Severity == sev {
			out[f.Code]++
		}
	}
	return out
}

// goldenGood and goldenBad count the golden subtests that ran.
var goldenGood, goldenBad int

// goldenRun runs one golden subtest and counts it by its prefix.
func goldenRun(t *testing.T, name string, fn func(t *testing.T)) {
	t.Helper()
	switch {
	case strings.HasPrefix(name, "good/"):
		goldenGood++
	case strings.HasPrefix(name, "bad/"):
		goldenBad++
	}
	t.Run(name, fn)
}

func wantPass(t *testing.T, name string, ds deckSpec) {
	t.Helper()
	goldenRun(t, "good/"+name, func(t *testing.T) {
		res := validate(t, ds)
		if !res.Passed {
			t.Errorf("deck must pass, blocks: %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
}

func wantBlock(t *testing.T, name string, ds deckSpec, wantCode string) {
	t.Helper()
	goldenRun(t, "bad/"+name, func(t *testing.T) {
		res := validate(t, ds)
		if res.Passed {
			t.Fatalf("deck must fail with %s, but passed", wantCode)
		}
		if codes(res, mtgv1.Severity_SEVERITY_BLOCK)[wantCode] == 0 {
			t.Errorf("want block %s, got %v", wantCode, codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
}

// commander99 is a legal mono-white core: commander + cards + Plains fill.
func monoW(commanders []string, extra map[string]int32, bracket int32) deckSpec {
	cardsMap := map[string]int32{
		"Soul Warden": 1, "Ajani's Pridemate": 1, "Ajani's Welcome": 1,
		"Swords to Plowshares": 1, "Path to Exile": 1, "Wrath of God": 1,
		"Day of Judgment": 1, "Mind Stone": 1, "Arcane Signet": 1,
		"Serra Angel": 1, "Evolving Wilds": 1, "Command Tower": 1,
	}
	for k, v := range extra {
		cardsMap[k] = v
	}
	return deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: bracket,
		commanders: commanders, cards: cardsMap, fill: "Plains", fillTo: 100}
}

// TestGoldenDecks is the PR-5 gate: 30 known-good and 30 known-bad decks.
func TestGoldenDecks(t *testing.T) {
	// ---- 30 good decks ----
	wantPass(t, "heliod lifegain b2", monoW([]string{"Heliod, Sun-Crowned"}, nil, 2))
	wantPass(t, "heliod no bracket", monoW([]string{"Heliod, Sun-Crowned"}, nil, 0))
	wantPass(t, "heliod b1", monoW([]string{"Heliod, Sun-Crowned"}, nil, 1))
	wantPass(t, "heliod b3 three changers", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"Smothering Tithe": 1, "The One Ring": 1, "Ancient Tomb": 1}, 3))
	wantPass(t, "heliod b4 three changers unlimited", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"Smothering Tithe": 1, "The One Ring": 1, "Ancient Tomb": 1}, 4))
	wantPass(t, "heliod b5", monoW([]string{"Heliod, Sun-Crowned"}, map[string]int32{"Smothering Tithe": 1}, 5))
	wantPass(t, "wilson background pair", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Wilson, Refined Grizzly", "Raised by Giants"},
		cards:      map[string]int32{"Cultivate": 1, "Rampant Growth": 1, "Llanowar Elves": 1, "Beast Within": 1, "Harmonize": 1},
		fill:       "Forest", fillTo: 100})
	wantPass(t, "doctor and companion pair", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"The Tenth Doctor", "Clara Oswald"},
		cards:      map[string]int32{"Counterspell": 1, "Negate": 1, "Divination": 1},
		fill:       "Island", fillTo: 100})
	// The Lutri ban is companion-only: in the 99 under a UR commander it is legal.
	wantPass(t, "lutri in the 99 is legal", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Niv-Mizzet, Parun"},
		cards:      map[string]int32{"Lutri, the Spellchaser": 1, "Counterspell": 1, "Negate": 1},
		fill:       "Island", fillTo: 100})
	wantPass(t, "relentless rats many copies", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Felothar the Steadfast"},
		cards:      map[string]int32{"Relentless Rats": 30, "Murder": 1, "Doom Blade": 1},
		fill:       "Swamp", fillTo: 100})
	wantPass(t, "snow basics", monoW([]string{"Heliod, Sun-Crowned"}, map[string]int32{"Snow-Covered Plains": 10}, 2))
	for i, n := range []int32{60, 61, 75} {
		wantPass(t, "modern lifegain size "+string(rune('a'+i)), deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
			cards: map[string]int32{"Soul Warden": 4, "Ajani's Pridemate": 4, "Ajani's Welcome": 4, "Path to Exile": 4},
			fill:  "Plains", fillTo: n})
	}
	wantPass(t, "modern with sideboard 15", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards:     map[string]int32{"Soul Warden": 4, "Path to Exile": 4},
		sideboard: map[string]int32{"Negate": 4, "Counterspell": 4, "Wrath of God": 4, "Day of Judgment": 3},
		fill:      "Plains", fillTo: 60})
	// Yorion is banned in Modern, so the companion case uses Kaheera. It
	// is legal in Modern, checked against the card snapshot (D-155 dropped
	// Legacy, which used to host this case).
	wantPass(t, "modern kaheera companion", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		companion: "Kaheera, the Orphanguard",
		cards:     map[string]int32{"Soul Warden": 4},
		sideboard: map[string]int32{"Kaheera, the Orphanguard": 1},
		fill:      "Plains", fillTo: 60})
	wantPass(t, "modern fair deck", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards: map[string]int32{"Counterspell": 4, "Delver of Secrets // Insectile Aberration": 4},
		fill:  "Island", fillTo: 60})
	wantPass(t, "house anything", deckSpec{format: mtgv1.FormatId_FORMAT_ID_HOUSE,
		cards: map[string]int32{"Black Lotus": 4, "The One Ring": 4},
		fill:  "Island", fillTo: 60})
	wantPass(t, "standard basics only", deckSpec{format: mtgv1.FormatId_FORMAT_ID_STANDARD,
		cards: map[string]int32{}, fill: "Plains", fillTo: 60})
	// owned modes
	goldenRun(t, "good/owned-first warns but passes", func(t *testing.T) {
		ds := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2)
		res := testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex,
			PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_FIRST, OracleCounts: map[string]int32{}})
		if !res.Passed {
			t.Errorf("owned-first must warn, not block: %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
		if codes(res, mtgv1.Severity_SEVERITY_WARN)[CodeNotOwned] == 0 {
			t.Error("want not_owned warns")
		}
	})
	goldenRun(t, "good/any-card ignores ownership", func(t *testing.T) {
		ds := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2)
		res := testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex,
			PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD, OracleCounts: map[string]int32{}})
		if !res.Passed || codes(res, mtgv1.Severity_SEVERITY_WARN)[CodeNotOwned] != 0 {
			t.Errorf("any-card must ignore ownership: %v", res.Findings)
		}
	})
	goldenRun(t, "good/owned-only with full collection", func(t *testing.T) {
		ds := deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
			cards: map[string]int32{"Soul Warden": 4}, fill: "Plains", fillTo: 60}
		res := testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex,
			PoolRule:     mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
			OracleCounts: map[string]int32{oid(t, "Soul Warden"): 4}})
		if !res.Passed {
			t.Errorf("blocks: %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
	// land-count warn is not a failure
	goldenRun(t, "good/low lands warns only", func(t *testing.T) {
		ds := deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
			cards: map[string]int32{"Soul Warden": 4, "Counterspell": 4, "Ajani's Pridemate": 4, "Path to Exile": 4,
				"Negate": 4, "Divination": 4, "Murder": 4, "Doom Blade": 4, "Serra Angel": 4, "Shivan Dragon": 4,
				"Wrath of God": 4, "Day of Judgment": 4, "Harmonize": 4, "Beast Within": 4, "Cultivate": 4},
			fill: "Plains", fillTo: 61} // 1 land only
		res := validate(t, ds)
		if !res.Passed {
			t.Errorf("blocks: %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
		if codes(res, mtgv1.Severity_SEVERITY_WARN)[CodeLandCount] == 0 {
			t.Error("want land_count warn")
		}
	})
	wantPass(t, "modern commons", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards: map[string]int32{"Counterspell": 4}, fill: "Island", fillTo: 60})
	wantPass(t, "esika front face commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 4,
		commanders: []string{"Esika, God of the Tree // The Prismatic Bridge"},
		cards:      map[string]int32{"Cultivate": 1, "Llanowar Elves": 1},
		fill:       "Forest", fillTo: 100})
	wantPass(t, "tovolar transform commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 3,
		commanders: []string{"Tovolar, Dire Overlord // Tovolar, the Midnight Scourge"},
		cards:      map[string]int32{"Rampant Growth": 1, "Shivan Dragon": 1},
		fill:       "Forest", fillTo: 100})
	wantPass(t, "partner with pair pir and toothy", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Pir, Imaginative Rascal", "Toothy, Imaginary Friend"},
		cards:      map[string]int32{"Counterspell": 1, "Cultivate": 1},
		fill:       "Forest", fillTo: 100})
	wantPass(t, "friends forever pair", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Cecily, Haunted Mage", "Bjorna, Nightfall Alchemist"},
		cards:      map[string]int32{"Counterspell": 1, "Murder": 1},
		fill:       "Island", fillTo: 100})
	wantPass(t, "partner pair thrasios and tymna", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 3,
		commanders: []string{"Thrasios, Triton Hero", "Tymna the Weaver"},
		cards:      map[string]int32{"Counterspell": 1, "Swords to Plowshares": 1, "Murder": 1},
		fill:       "Island", fillTo: 100})
	wantPass(t, "partner survivors pair", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Abby, Merciless Soldier", "Ellie, Brick Master"},
		cards:      map[string]int32{"Cultivate": 1, "Shivan Dragon": 1},
		fill:       "Mountain", fillTo: 100})
	wantPass(t, "doctor pair reversed", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Clara Oswald", "The Tenth Doctor"},
		cards:      map[string]int32{"Counterspell": 1, "Negate": 1, "Divination": 1},
		fill:       "Island", fillTo: 100})
	wantPass(t, "vehicle commander weatherlight", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Weatherlight"},
		cards:      map[string]int32{"Sol Ring": 1, "Mind Stone": 1},
		fill:       "Wastes", fillTo: 100})
	wantPass(t, "spacecraft commander dawnsire", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Dawnsire, Sunstar Dreadnought"},
		cards:      map[string]int32{"Sol Ring": 1, "Everflowing Chalice": 1},
		fill:       "Wastes", fillTo: 100})
	wantPass(t, "nazgul nine copies", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Felothar the Steadfast"},
		cards:      map[string]int32{"Nazgûl": 9, "Murder": 1},
		fill:       "Swamp", fillTo: 100})
	wantPass(t, "seven dwarves seven copies", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards: map[string]int32{"Seven Dwarves": 7, "Skullcrack": 4},
		fill:  "Mountain", fillTo: 60})
	wantPass(t, "game changer commander in bracket 4", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 4,
		commanders: []string{"Tergrid, God of Fright // Tergrid's Lantern"},
		cards:      map[string]int32{"Murder": 1, "Doom Blade": 1},
		fill:       "Swamp", fillTo: 100})
	wantPass(t, "wastes fill colorless commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"The Reaper, King No More"},
		cards:      map[string]int32{"Sol Ring": 1},
		fill:       "Wastes", fillTo: 100})
	goldenRun(t, "good/house format reports info only", func(t *testing.T) {
		ds := deckSpec{format: mtgv1.FormatId_FORMAT_ID_HOUSE,
			cards: map[string]int32{"Dockside Extortionist": 4}, fill: "Mountain", fillTo: 60}
		res := validate(t, ds)
		if !res.Passed || codes(res, mtgv1.Severity_SEVERITY_INFO)[CodeHouseRules] == 0 {
			t.Errorf("want pass with house_rules_limited info, got %v", res.Findings)
		}
		if res.Format != mtgv1.FormatId_FORMAT_ID_HOUSE {
			t.Errorf("format = %v", res.Format)
		}
	})
	goldenRun(t, "good/owned-only aggregates per oracle id", func(t *testing.T) {
		ds := deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
			cards: map[string]int32{"Soul Warden": 3}, sideboard: map[string]int32{"Soul Warden": 1},
			fill: "Plains", fillTo: 60}
		res := testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex,
			PoolRule:     mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
			OracleCounts: map[string]int32{oid(t, "Soul Warden"): 4}})
		if !res.Passed {
			t.Errorf("blocks: %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
	goldenRun(t, "good/bracket info names the verified date", func(t *testing.T) {
		res := validate(t, monoW([]string{"Heliod, Sun-Crowned"}, nil, 2))
		var found bool
		for _, f := range res.Findings {
			if f.Code == CodeBracketProse && strings.Contains(f.Message, testCfg.VerifiedAt["brackets.json"]) &&
				strings.Contains(f.Message, "8+") {
				found = true
			}
		}
		if !found {
			t.Errorf("want bracket info with the verified date and expected turns, got %v", res.Findings)
		}
	})

	// Grist, the Hunger Tide is a creature outside the battlefield by a
	// characteristic-defining ability, so it leads a deck (D-140).
	wantPass(t, "grist leads by its creature CDA", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Grist, the Hunger Tide"},
		fill:       "Swamp", fillTo: 100})

	// ---- 30 bad decks ----
	bad99 := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2)
	bad99.fillTo = 99
	wantBlock(t, "commander 99 cards", bad99, CodeDeckSize)
	bad101 := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2)
	bad101.fillTo = 101
	wantBlock(t, "commander 101 cards", bad101, CodeDeckSize)
	wantBlock(t, "modern 59 cards", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards: map[string]int32{"Soul Warden": 4}, fill: "Plains", fillTo: 59}, CodeDeckSize)
	wantBlock(t, "duplicate nonbasic in commander", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"Soul Warden": 2}, 2), CodeCopyLimit)
	wantBlock(t, "five copies in modern", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards: map[string]int32{"Soul Warden": 5}, fill: "Plains", fillTo: 60}, CodeCopyLimit)
	wantBlock(t, "dockside banned in commander", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"Dockside Extortionist": 1}, 2), CodeBannedCard)
	wantBlock(t, "golos banned in commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Golos, Tireless Pilgrim"}, fill: "Plains", fillTo: 100}, CodeBannedCard)
	wantBlock(t, "one ring banned in modern", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards: map[string]int32{"The One Ring": 1}, fill: "Plains", fillTo: 60}, CodeBannedCard)
	wantBlock(t, "lotus banned in commander", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"Black Lotus": 1}, 2), CodeBannedCard)
	wantBlock(t, "brainstorm not legal in modern", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards: map[string]int32{"Brainstorm": 4}, fill: "Island", fillTo: 60}, CodeNotLegal)
	wantBlock(t, "no commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		fill: "Plains", fillTo: 100}, CodeNoCommander)
	wantBlock(t, "serra angel not legendary", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Serra Angel"}, fill: "Plains", fillTo: 100}, CodeBadCommander)
	wantBlock(t, "sol ring not a creature", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Sol Ring"}, fill: "Plains", fillTo: 100}, CodeBadCommander)
	wantBlock(t, "three commanders", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Heliod, Sun-Crowned", "Serra Angel", "Thrasios, Triton Hero"},
		fill:       "Plains", fillTo: 100}, CodeBadCommander)
	wantBlock(t, "partner with non-partner", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Thrasios, Triton Hero", "Heliod, Sun-Crowned"},
		fill:       "Island", fillTo: 100}, CodeBadPartner)
	wantBlock(t, "partner-with wrong partner", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Halana and Alena, Partners", "Thrasios, Triton Hero"},
		fill:       "Forest", fillTo: 100}, CodeBadPartner)
	wantBlock(t, "two backgrounds", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Raised by Giants", "Raised by Giants"},
		fill:       "Forest", fillTo: 100}, CodeCopyLimit)
	wantBlock(t, "off color pridemate under green", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Wilson, Refined Grizzly"},
		cards:      map[string]int32{"Ajani's Pridemate": 1},
		fill:       "Forest", fillTo: 100}, CodeOffColor)
	wantBlock(t, "off color murder under heliod", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"Murder": 1}, 2), CodeOffColor)
	wantBlock(t, "one changer in bracket 1", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"Smothering Tithe": 1}, 1), CodeGameChangers)
	wantBlock(t, "one changer in bracket 2", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"The One Ring": 1}, 2), CodeGameChangers)
	wantBlock(t, "four changers in bracket 3", monoW([]string{"Heliod, Sun-Crowned"},
		map[string]int32{"Smothering Tithe": 1, "The One Ring": 1, "Ancient Tomb": 1, "Rhystic Study": 1}, 3), CodeGameChangers)
	wantBlock(t, "lutri banned as companion", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Heliod, Sun-Crowned"}, companion: "Lutri, the Spellchaser",
		fill: "Plains", fillTo: 100}, CodeCompanionBanned)
	wantBlock(t, "non-companion as companion", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		companion: "Serra Angel",
		cards:     map[string]int32{"Soul Warden": 4}, fill: "Plains", fillTo: 60}, CodeBadCompanion)
	wantBlock(t, "sideboard 16", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards:     map[string]int32{"Soul Warden": 4},
		sideboard: map[string]int32{"Negate": 4, "Counterspell": 4, "Wrath of God": 4, "Day of Judgment": 4},
		fill:      "Plains", fillTo: 60}, CodeSideboardSize)
	wantBlock(t, "commander sideboard not allowed", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Heliod, Sun-Crowned"},
		sideboard:  map[string]int32{"Negate": 1},
		fill:       "Plains", fillTo: 100}, CodeSideboardSize)
	wantBlock(t, "banned card in sideboard", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards:     map[string]int32{"Soul Warden": 4},
		sideboard: map[string]int32{"The One Ring": 1},
		fill:      "Plains", fillTo: 60}, CodeBannedCard)
	goldenRun(t, "bad/unknown card id", func(t *testing.T) {
		deck := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2).build(t)
		deck.Cards = append(deck.Cards, &mtgv1.DeckCard{OracleId: "not-a-real-id", Name: "Ghost", Count: 1})
		deck.Cards[len(deck.Cards)-2].Count-- // keep 100
		res := testCfg.Validate(Input{Deck: deck, Cards: testIndex})
		if res.Passed || codes(res, mtgv1.Severity_SEVERITY_BLOCK)[CodeUnknownCard] == 0 {
			t.Errorf("want unknown_card block, got %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
	wantBlock(t, "partner with alone with plain partner", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Pir, Imaginative Rascal", "Thrasios, Triton Hero"},
		fill:       "Forest", fillTo: 100}, CodeBadPartner)
	wantBlock(t, "survivors with plain partner", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Abby, Merciless Soldier", "Thrasios, Triton Hero"},
		fill:       "Forest", fillTo: 100}, CodeBadPartner)
	wantBlock(t, "survivors with father and son", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Abby, Merciless Soldier", "Kratos, Stoic Father"},
		fill:       "Mountain", fillTo: 100}, CodeBadPartner)
	wantBlock(t, "friends forever with plain partner", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Cecily, Haunted Mage", "Thrasios, Triton Hero"},
		fill:       "Island", fillTo: 100}, CodeBadPartner)
	wantBlock(t, "lone background", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Raised by Giants"},
		fill:       "Forest", fillTo: 100}, CodeBadCommander)
	wantBlock(t, "background with plain partner", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Thrasios, Triton Hero", "Raised by Giants"},
		fill:       "Forest", fillTo: 100}, CodeBadCommander)
	wantBlock(t, "doctor companion with human doctor", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Clara Oswald", "Moonstone, Harsh Mistress"},
		fill:       "Island", fillTo: 100}, CodeBadPartner)
	wantBlock(t, "doctor companion with time lord rogue", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Clara Oswald", "The Master, Formed Anew"},
		fill:       "Island", fillTo: 100}, CodeBadPartner)
	wantBlock(t, "back face legendary bloodline keeper", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Bloodline Keeper // Lord of Lineage"},
		fill:       "Swamp", fillTo: 100}, CodeBadCommander)
	wantBlock(t, "nazgul ten copies", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Felothar the Steadfast"},
		cards:      map[string]int32{"Nazgûl": 10},
		fill:       "Swamp", fillTo: 100}, CodeCopyLimit)
	wantBlock(t, "seven dwarves eight copies", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		cards: map[string]int32{"Seven Dwarves": 8},
		fill:  "Mountain", fillTo: 60}, CodeCopyLimit)
	wantBlock(t, "game changer commander in bracket 2", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Tergrid, God of Fright // Tergrid's Lantern"},
		fill:       "Swamp", fillTo: 100}, CodeGameChangers)
	wantBlock(t, "bracket 6 unknown", monoW([]string{"Heliod, Sun-Crowned"}, nil, 6), CodeUnknownBracket)
	wantBlock(t, "unknown format", deckSpec{format: mtgv1.FormatId_FORMAT_ID_UNSPECIFIED,
		fill: "Plains", fillTo: 60}, CodeUnknownFormat)
	wantBlock(t, "companion off color in commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Heliod, Sun-Crowned"}, companion: "Yorion, Sky Nomad",
		fill: "Plains", fillTo: 100}, CodeOffColor)
	wantBlock(t, "lurrus companion banned in modern", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		companion: "Lurrus of the Dream-Den",
		cards:     map[string]int32{"Soul Warden": 4}, fill: "Plains", fillTo: 60}, CodeBannedCard)
	wantBlock(t, "companion not in sideboard", deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
		companion: "Kaheera, the Orphanguard",
		cards:     map[string]int32{"Soul Warden": 4}, fill: "Plains", fillTo: 60}, CodeCompanionNotSide)
	wantBlock(t, "companion breaks singleton in commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Niv-Mizzet, Parun"}, companion: "Jegantha, the Wellspring",
		cards: map[string]int32{"Jegantha, the Wellspring": 1},
		fill:  "Island", fillTo: 100}, CodeCopyLimit)
	// A legendary planeswalker with no creature CDA is not a commander
	// (D-140).
	wantBlock(t, "planeswalker without the creature CDA", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Kaya, Ghost Haunter"},
		fill:       "Swamp", fillTo: 100}, CodeBadCommander)
	goldenRun(t, "bad/nil cards gives a block finding", func(t *testing.T) {
		deck := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2).build(t)
		res := testCfg.Validate(Input{Deck: deck})
		if res.Passed || codes(res, mtgv1.Severity_SEVERITY_BLOCK)[CodeNoCardData] == 0 {
			t.Errorf("want no_card_data block, got %v", res.Findings)
		}
		res = testCfg.Validate(Input{Cards: testIndex})
		if res.Passed {
			t.Error("nil deck must not pass")
		}
	})
	goldenRun(t, "bad/owned-only split across main and side", func(t *testing.T) {
		ds := deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
			cards: map[string]int32{"Soul Warden": 3}, sideboard: map[string]int32{"Soul Warden": 1},
			fill: "Plains", fillTo: 60}
		res := testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex,
			PoolRule:     mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
			OracleCounts: map[string]int32{oid(t, "Soul Warden"): 3}})
		if res.Passed || codes(res, mtgv1.Severity_SEVERITY_BLOCK)[CodeNotOwned] != 1 {
			t.Errorf("want one not_owned block, got %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
	goldenRun(t, "bad/owned-only without cards", func(t *testing.T) {
		ds := deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
			cards: map[string]int32{"Soul Warden": 4}, fill: "Plains", fillTo: 60}
		res := testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex,
			PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY, OracleCounts: map[string]int32{}})
		if res.Passed || codes(res, mtgv1.Severity_SEVERITY_BLOCK)[CodeNotOwned] == 0 {
			t.Errorf("want not_owned block, got %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
}

// TestNegativeCountBlocks: a count under one is a bad_count, and no sum
// reads it.
func TestNegativeCountBlocks(t *testing.T) {
	deck := monoW([]string{"Heliod, Sun-Crowned"}, nil, 0).build(t)
	deck.Cards = append(deck.Cards,
		&mtgv1.DeckCard{OracleId: oid(t, "Plains"), Name: "Plains", Count: 1},
		&mtgv1.DeckCard{OracleId: oid(t, "Serra Angel"), Name: "Serra Angel", Count: -1},
	)
	res := testCfg.Validate(Input{Deck: deck, Cards: testIndex})
	if res.Passed {
		t.Fatal("a negative count must block")
	}
	blocks := codes(res, mtgv1.Severity_SEVERITY_BLOCK)
	if blocks[CodeBadCount] != 1 || blocks[CodeDeckSize] != 1 {
		t.Errorf("blocks = %v, want one bad_count and one deck_size", blocks)
	}
	if n := mainDeckCount(deck); n != 101 {
		t.Errorf("mainDeckCount = %d, want 101: the negative entry must not count", n)
	}
	// A zero count in the sideboard is a bad_count too.
	deck = deckSpec{format: mtgv1.FormatId_FORMAT_ID_STANDARD, fill: "Plains", fillTo: 60}.build(t)
	deck.Sideboard = append(deck.Sideboard, &mtgv1.DeckCard{OracleId: oid(t, "Plains"), Name: "Plains", Count: 0})
	res = testCfg.Validate(Input{Deck: deck, Cards: testIndex})
	if codes(res, mtgv1.Severity_SEVERITY_BLOCK)[CodeBadCount] != 1 {
		t.Errorf("sideboard zero count: %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
	}
}

// TestIsDoctorIsExported keeps the helper the candidate builder relies on.
func TestIsDoctorIsExported(t *testing.T) {
	c, ok := testIndex.ByName("The Tenth Doctor")
	if !ok {
		t.Fatal("fixture card missing: The Tenth Doctor")
	}
	if !IsDoctor(c) {
		t.Errorf("The Tenth Doctor must be a Doctor: %v", c.Subtypes)
	}
	if m, ok := testIndex.ByName("Moonstone, Harsh Mistress"); ok && IsDoctor(m) {
		t.Error("a Human Doctor Villain is not a Time Lord Doctor")
	}
}

func TestLoadData(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	// Four formats since D-155: Commander, Standard, Modern, and HOUSE.
	if len(cfg.Formats) < 4 {
		t.Errorf("formats = %d", len(cfg.Formats))
	}
	if cfg.Brackets[3].MaxGameChangers != 3 || cfg.Brackets[4].MaxGameChangers != -1 {
		t.Errorf("brackets = %v", cfg.Brackets)
	}
	if !slices.Contains(cfg.BannedAsCompanion["Lutri, the Spellchaser"], "commander") {
		t.Error("Lutri missing from companion bans for commander")
	}
	if len(cfg.BannedAsCompanion["Lutri, the Spellchaser"]) != 1 {
		t.Errorf("Lutri ban formats = %v, want commander only", cfg.BannedAsCompanion["Lutri, the Spellchaser"])
	}
	for _, name := range []string{"formats.json", "brackets.json", "companion_bans.json"} {
		if _, err := time.Parse("2006-01-02", cfg.VerifiedAt[name]); err != nil {
			t.Errorf("%s verified_at = %q: %v", name, cfg.VerifiedAt[name], err)
		}
	}
	if cfg.Brackets[1].ExpectedTurns != "9+" || cfg.Brackets[5].ExpectedTurns != "any" {
		t.Errorf("brackets = %v", cfg.Brackets)
	}
}

// TestGoldenCounts is the PR-5 gate size: at least 30 good and 30 bad
// decks. It reads the counters that TestGoldenDecks filled. The set holds
// 31 good and 30 bad decks (D-140).
func TestGoldenCounts(t *testing.T) {
	if goldenGood == 0 && goldenBad == 0 {
		t.Skip("TestGoldenDecks did not run")
	}
	if goldenGood < 31 || goldenBad < 30 {
		t.Errorf("golden gate: %d good, %d bad, want at least 31 and 30", goldenGood, goldenBad)
	}
	t.Logf("golden gate: %d good, %d bad", goldenGood, goldenBad)
}

// TestSecondCommanderIsCheckedAfterAnUnknownFirst: an unknown first id
// must not hide a second commander that can not lead.
func TestSecondCommanderIsCheckedAfterAnUnknownFirst(t *testing.T) {
	deck := deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2, fill: "Plains", fillTo: 100}.build(t)
	deck.CommanderOracleIds = []string{"not-an-id", oid(t, "Serra Angel")}
	res := testCfg.Validate(Input{Deck: deck, Cards: testIndex})
	blocks := codes(res, mtgv1.Severity_SEVERITY_BLOCK)
	if blocks[CodeUnknownCard] == 0 || blocks[CodeBadCommander] == 0 {
		t.Errorf("want unknown_card and bad_commander, got %v", blocks)
	}
	if blocks[CodeBadPartner] != 0 {
		t.Errorf("a pair with an unknown card must not be judged: %v", blocks)
	}
}

// TestDuplicateCommanderIsABadPartner: the same Oracle id twice is not a
// pair (CR 702.124f).
func TestDuplicateCommanderIsABadPartner(t *testing.T) {
	deck := deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 2,
		commanders: []string{"Thrasios, Triton Hero", "Thrasios, Triton Hero"},
		fill:       "Island", fillTo: 100}.build(t)
	res := testCfg.Validate(Input{Deck: deck, Cards: testIndex})
	var found bool
	for _, f := range res.Findings {
		if f.Code == CodeBadPartner && strings.Contains(f.Message, "duplicate commander") {
			found = true
		}
	}
	if !found {
		t.Errorf("want a duplicate commander bad_partner, got %v", res.Findings)
	}
	c, _ := testIndex.ByName("Thrasios, Triton Hero")
	if ValidPair(c, c) {
		t.Error("ValidPair(a, a) must be false")
	}
}

// TestZeroCountRowIsNotOffColor: a zero-count row is a bad_count and
// nothing else reads it.
func TestZeroCountRowIsNotOffColor(t *testing.T) {
	deck := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2).build(t)
	deck.Cards = append(deck.Cards, &mtgv1.DeckCard{OracleId: oid(t, "Murder"), Name: "Murder", Count: 0})
	res := testCfg.Validate(Input{Deck: deck, Cards: testIndex})
	blocks := codes(res, mtgv1.Severity_SEVERITY_BLOCK)
	if blocks[CodeBadCount] != 1 || blocks[CodeOffColor] != 0 {
		t.Errorf("want one bad_count and no off_color, got %v", blocks)
	}
}
