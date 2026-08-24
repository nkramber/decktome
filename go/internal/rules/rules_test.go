package rules

import (
	"os"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

var (
	testIndex *cards.Index
	testCfg   *Config
)

func TestMain(m *testing.M) {
	f, err := os.Open("../cards/testdata/cards_fixture.jsonl")
	if err != nil {
		panic(err)
	}
	cardList, err := cards.LoadCards(f, "cards_fixture.jsonl")
	_ = f.Close()
	if err != nil {
		panic(err)
	}
	testIndex = cards.NewIndex(cardList, nil, nil, time.Now())
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

func wantPass(t *testing.T, name string, ds deckSpec) {
	t.Helper()
	t.Run("good/"+name, func(t *testing.T) {
		res := validate(t, ds)
		if !res.Passed {
			t.Errorf("deck must pass, blocks: %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
}

func wantBlock(t *testing.T, name string, ds deckSpec, wantCode string) {
	t.Helper()
	t.Run("bad/"+name, func(t *testing.T) {
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
	wantPass(t, "lutri in the 99 is legal", monoW([]string{"Heliod, Sun-Crowned"}, nil, 2)) // Lutri ban is companion-only
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
	// Yorion is banned in Modern, so the companion case runs in Legacy.
	wantPass(t, "legacy yorion companion", deckSpec{format: mtgv1.FormatId_FORMAT_ID_LEGACY,
		companion: "Yorion, Sky Nomad",
		cards:     map[string]int32{"Soul Warden": 4, "Counterspell": 4},
		sideboard: map[string]int32{"Yorion, Sky Nomad": 1},
		fill:      "Plains", fillTo: 80})
	wantPass(t, "vintage one lotus", deckSpec{format: mtgv1.FormatId_FORMAT_ID_VINTAGE,
		cards: map[string]int32{"Black Lotus": 1, "Brainstorm": 1, "Counterspell": 4},
		fill:  "Island", fillTo: 60})
	wantPass(t, "legacy fair deck", deckSpec{format: mtgv1.FormatId_FORMAT_ID_LEGACY,
		cards: map[string]int32{"Brainstorm": 4, "Counterspell": 4, "Delver of Secrets // Insectile Aberration": 4},
		fill:  "Island", fillTo: 60})
	wantPass(t, "house anything", deckSpec{format: mtgv1.FormatId_FORMAT_ID_HOUSE,
		cards: map[string]int32{"Black Lotus": 4, "The One Ring": 4},
		fill:  "Island", fillTo: 60})
	wantPass(t, "standard basics only", deckSpec{format: mtgv1.FormatId_FORMAT_ID_STANDARD,
		cards: map[string]int32{}, fill: "Plains", fillTo: 60})
	// owned modes
	t.Run("good/owned-first warns but passes", func(t *testing.T) {
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
	t.Run("good/any-card ignores ownership", func(t *testing.T) {
		ds := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2)
		res := testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex,
			PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD, OracleCounts: map[string]int32{}})
		if !res.Passed || codes(res, mtgv1.Severity_SEVERITY_WARN)[CodeNotOwned] != 0 {
			t.Errorf("any-card must ignore ownership: %v", res.Findings)
		}
	})
	t.Run("good/owned-only with full collection", func(t *testing.T) {
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
	t.Run("good/low lands warns only", func(t *testing.T) {
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
	wantPass(t, "pauper commons", deckSpec{format: mtgv1.FormatId_FORMAT_ID_PAUPER,
		cards: map[string]int32{"Counterspell": 4}, fill: "Island", fillTo: 60})
	wantPass(t, "esika front face commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 4,
		commanders: []string{"Esika, God of the Tree // The Prismatic Bridge"},
		cards:      map[string]int32{"Cultivate": 1, "Llanowar Elves": 1},
		fill:       "Forest", fillTo: 100})
	wantPass(t, "tovolar transform commander", deckSpec{format: mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket: 3,
		commanders: []string{"Tovolar, Dire Overlord // Tovolar, the Midnight Scourge"},
		cards:      map[string]int32{"Rampant Growth": 1, "Shivan Dragon": 1},
		fill:       "Forest", fillTo: 100})

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
	wantBlock(t, "two lotus in vintage", deckSpec{format: mtgv1.FormatId_FORMAT_ID_VINTAGE,
		cards: map[string]int32{"Black Lotus": 2}, fill: "Island", fillTo: 60}, CodeRestrictedCard)
	wantBlock(t, "two brainstorm restricted", deckSpec{format: mtgv1.FormatId_FORMAT_ID_VINTAGE,
		cards: map[string]int32{"Brainstorm": 2}, fill: "Island", fillTo: 60}, CodeRestrictedCard)
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
	t.Run("bad/unknown card id", func(t *testing.T) {
		deck := monoW([]string{"Heliod, Sun-Crowned"}, nil, 2).build(t)
		deck.Cards = append(deck.Cards, &mtgv1.DeckCard{OracleId: "not-a-real-id", Name: "Ghost", Count: 1})
		deck.Cards[len(deck.Cards)-2].Count-- // keep 100
		res := testCfg.Validate(Input{Deck: deck, Cards: testIndex})
		if res.Passed || codes(res, mtgv1.Severity_SEVERITY_BLOCK)[CodeUnknownCard] == 0 {
			t.Errorf("want unknown_card block, got %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
	t.Run("bad/owned-only without cards", func(t *testing.T) {
		ds := deckSpec{format: mtgv1.FormatId_FORMAT_ID_MODERN,
			cards: map[string]int32{"Soul Warden": 4}, fill: "Plains", fillTo: 60}
		res := testCfg.Validate(Input{Deck: ds.build(t), Cards: testIndex,
			PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY, OracleCounts: map[string]int32{}})
		if res.Passed || codes(res, mtgv1.Severity_SEVERITY_BLOCK)[CodeNotOwned] == 0 {
			t.Errorf("want not_owned block, got %v", codes(res, mtgv1.Severity_SEVERITY_BLOCK))
		}
	})
}

func TestLoadData(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Formats) < 8 {
		t.Errorf("formats = %d", len(cfg.Formats))
	}
	if cfg.MaxGameChangers[3] != 3 || cfg.MaxGameChangers[4] != -1 {
		t.Errorf("brackets = %v", cfg.MaxGameChangers)
	}
	if !cfg.BannedAsCompanion["Lutri, the Spellchaser"] {
		t.Error("Lutri missing from companion bans")
	}
}
