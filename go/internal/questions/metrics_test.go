package questions

import (
	"context"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// TestCoverageCountsTheSources is the M-4 report (F-17): how many
// questions came from the catalog, how many the model invented, and
// whether an invented question filled its slot.
func TestCoverageCountsTheSources(t *testing.T) {
	first := commanderClassify()
	// The commander row fits poorly, so the model replaces it. The bracket
	// row fits, so the catalog wins.
	score := scoreStep(t,
		scored{RowID: "commander", Fit: 0.10, CustomText: "Which legend do you enjoy playing?", Reason: "test"},
		scored{RowID: "power_commander", Fit: 0.90, Reason: "fits"})
	// Turn 2 answers the invented question and leaves the bracket open.
	second := classifyOut{Format: "unknown", PoolRule: "unknown"}
	second.CommanderNames = []string{"Karlov of the Ghost Council"}

	a, _ := testAgent(t,
		classifyStep(t, first), score, askStep(t),
		classifyStep(t, second), fits(t), askStep(t))
	st := NewState(false)

	res, err := a.Turn(context.Background(), st, "lifegain commander deck", nil)
	if err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if res.Coverage.Asked != 2 || res.Coverage.Invented != 1 || res.Coverage.Catalog != 1 {
		t.Fatalf("coverage after turn 1 = %+v", res.Coverage)
	}
	if res.Coverage.CatalogOnly != 0 {
		t.Errorf("a session with an invented question counts as catalog-only")
	}
	if got := res.Coverage.MedianFit(); got != 0.5 {
		t.Errorf("median fit = %v, want 0.5", got)
	}

	res, err = a.Turn(context.Background(), st, "Karlov", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if res.Coverage.InventedFilled != 1 {
		t.Errorf("the invented question filled no slot: %+v", res.Coverage)
	}
	if res.Coverage.CatalogFilled != 0 {
		t.Errorf("the bracket question is still open, but the report calls it filled")
	}
	if n := res.Coverage.InventedByRow["commander"]; n != 1 {
		t.Errorf("invented by row = %v, want one for the commander row", res.Coverage.InventedByRow)
	}
	for _, ask := range st.Asks {
		if ask.QuestionID == "" || ask.Slot == "" || ask.Turn == 0 {
			t.Errorf("incomplete M-4 record %+v", ask)
		}
		if ask.Threshold != DefaultFitThreshold {
			t.Errorf("record %q carries threshold %v", ask.QuestionID, ask.Threshold)
		}
	}
}

// TestCoverageOverSessions is the shape the weekly report reads.
func TestCoverageOverSessions(t *testing.T) {
	var c Coverage
	c.Add([]Ask{{RowID: "commander", Fit: 0.9}, {RowID: "power_commander", Fit: 0.8}})
	c.Add([]Ask{{RowID: "colors", Fit: 0.2, Invented: true, Filled: true}})
	if c.Sessions != 2 || c.Asked != 3 || c.Catalog != 2 || c.Invented != 1 {
		t.Fatalf("coverage = %+v", c)
	}
	if c.CatalogOnly != 1 {
		t.Errorf("catalog-only sessions = %d, want 1", c.CatalogOnly)
	}
	if c.InventedFilled != 1 || c.CatalogFilled != 0 {
		t.Errorf("filled counts = %+v", c)
	}
	if got := c.MedianFit(); got != 0.8 {
		t.Errorf("median fit = %v, want 0.8", got)
	}
}

// TestCommanderFillsTheColorSlot keeps the card-pool question reachable.
// The pool row waits for the format, the colors, and the theme (D-67).
// A user who names a commander never answers a color question, so the
// commander must fill that slot.
func TestCommanderFillsTheColorSlot(t *testing.T) {
	out := classifyOut{Format: "commander", Theme: "lifegain", PoolRule: "unknown"}
	out.CommanderNames = []string{"Karlov of the Ghost Council"}
	a, _ := testAgent(t, classifyStep(t, out), fits(t, "power_commander", "pool"), askStep(t))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "karlov lifegain deck from my library", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if !st.Ctx.Filled["colors"] {
		t.Fatal("a set commander left the color slot open")
	}
	if question(res.Questions, "pool_rule") == nil {
		t.Errorf("the card-pool question did not fire: %+v", res.Questions)
	}
	if question(res.Questions, "colors") != nil {
		t.Error("the agent asked for colors although a commander is set")
	}
}

// TestMissingCommander is the trigger of the "Commander not owned" row.
// It must stay quiet on every uncertain case: no collection, no name, or
// a name the card database does not know.
func TestMissingCommander(t *testing.T) {
	karlov := &mtgv1.Card{OracleId: "oracle-karlov", Name: "Karlov of the Ghost Council"}
	oloro := &mtgv1.Card{OracleId: "oracle-oloro", Name: "Oloro, Ageless Ascetic"}
	idx := cards.NewIndex([]*mtgv1.Card{karlov, oloro}, nil, nil, time.Now())

	cases := []struct {
		name  string
		hints *CandidateHints
		names []string
		want  bool
	}{
		{"not in the collection", &CandidateHints{Index: idx, Owned: map[string]int32{"oracle-oloro": 1}},
			[]string{"Karlov of the Ghost Council"}, true},
		{"in the collection", &CandidateHints{Index: idx, Owned: map[string]int32{"oracle-karlov": 1}},
			[]string{"Karlov of the Ghost Council"}, false},
		{"no collection", &CandidateHints{Index: idx}, []string{"Karlov of the Ghost Council"}, false},
		{"no name", &CandidateHints{Index: idx, Owned: map[string]int32{"oracle-oloro": 1}}, nil, false},
		{"unknown name", &CandidateHints{Index: idx, Owned: map[string]int32{"oracle-oloro": 1}},
			[]string{"Not A Real Card"}, false},
		{"no index", &CandidateHints{Owned: map[string]int32{"oracle-oloro": 1}},
			[]string{"Karlov of the Ghost Council"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.hints.MissingCommander(tc.names); got != tc.want {
				t.Errorf("MissingCommander = %v, want %v", got, tc.want)
			}
		})
	}
	var nilHints *CandidateHints
	if nilHints.MissingCommander([]string{"Karlov of the Ghost Council"}) {
		t.Error("a nil hint source claimed a missing commander")
	}
	if thin, n := nilHints.ThinTheme("lifegain"); thin || n != 0 {
		t.Errorf("a nil hint source reported a thin theme: %v %d", thin, n)
	}
}

// countingHints records what the resolver asked for.
type countingHints struct {
	stubHints
	commanderCalls int
	colorCalls     int
}

func (c *countingHints) ThemeColors(theme string) string {
	c.colorCalls++
	return c.stubHints.ThemeColors(theme)
}

func (c *countingHints) Commanders(theme string, skip []string) []string {
	c.commanderCalls++
	return c.stubHints.Commanders(theme, skip)
}

// TestHintsAreAskedOnlyForWhatARowNames is the first half of the
// commander defect the gate run of 2026-08-25 found. An eager call runs a
// whole PR-6 build for a row that holds no placeholder.
func TestHintsAreAskedOnlyForWhatARowNames(t *testing.T) {
	c := load(t)
	cases := []struct {
		row                        string
		wantCommanders, wantColors int
	}{
		{"format", 0, 0},
		{"power_commander", 0, 0},
		{"commander", 0, 0},
		{"colors", 0, 1},
		{"commander_pick", 1, 0},
	}
	for _, tc := range cases {
		t.Run(tc.row, func(t *testing.T) {
			row, ok := c.Row(tc.row)
			if !ok {
				t.Fatalf("no row %q", tc.row)
			}
			h := &countingHints{stubHints: stubHints{colors: "white and black",
				commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}}}
			resolve(row, themedState("lifegain"), h)
			if h.commanderCalls != tc.wantCommanders {
				t.Errorf("row %q asked for commanders %d times, want %d", tc.row, h.commanderCalls, tc.wantCommanders)
			}
			if h.colorCalls != tc.wantColors {
				t.Errorf("row %q asked for theme colors %d times, want %d", tc.row, h.colorCalls, tc.wantColors)
			}
		})
	}
}

// TestHintCacheKeyCarriesTheColors is the second half. One key per theme
// kept the colorless commander list after the user named their colors.
// Conversation 24 of the gate run offered Lotho, Corrupt Shirriff (white
// and black), Zhao, the Moon Slayer (red), and Troyan, Gutsy Explorer
// (green and blue) for a white-blue deck. All three are outside the
// color identity.
func TestHintCacheKeyCarriesTheColors(t *testing.T) {
	none := &CandidateHints{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER}
	white := &CandidateHints{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Colors: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U}}
	if none.key("stax") == white.key("stax") {
		t.Errorf("one cache key serves two color sets: %q", none.key("stax"))
	}
	owned := &CandidateHints{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Colors: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U},
		Pool:   mtgv1.PoolRule_POOL_RULE_OWNED_ONLY}
	if owned.key("stax") == white.key("stax") {
		t.Errorf("one cache key serves two pool rules: %q", owned.key("stax"))
	}
}
