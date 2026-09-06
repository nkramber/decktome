package questions

import (
	"context"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// TestCoverageCountsTheSources is the M-4 report (F-17): how many
// questions came from the catalog, how many the model invented, and
// whether an invented question filled its slot.
func TestCoverageCountsTheSources(t *testing.T) {
	first := commanderClassify()
	// The session holds no collection, so a buy list exists and the
	// budget row fires. A named cap closes it, and this test counts the
	// commander rows alone (D-168).
	first.BudgetUSD = 50
	// The bracket row fits poorly, so the model replaces it. The commander
	// row is fixed and fits, so the catalog wins.
	score := scoreStep(t,
		scored{RowID: "power_commander", Fit: 0.10, CustomText: "How strong a table do you play at?", Reason: "test"},
		scored{RowID: "commander", Fit: 0.90, Reason: "fits"})
	// Turn 2 answers the invented question and leaves the commander open.
	second := classifyOut{Format: "unknown", PoolRule: "unknown", Power: "bracket 3"}

	a, _ := testAgent(t,
		classifyStep(t, first), score, askStep(t),
		classifyStep(t, second), fits(t), askStep(t))
	st := NewState(false)

	res, err := a.Turn(context.Background(), st, "lifegain commander deck, 50 dollars", nil)
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

	res, err = a.Turn(context.Background(), st, "bracket 3", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if res.Coverage.InventedFilled != 1 {
		t.Errorf("the invented question filled no slot: %+v", res.Coverage)
	}
	if res.Coverage.CatalogFilled != 0 {
		t.Errorf("the commander question is still open, but the report calls it filled")
	}
	if n := res.Coverage.InventedByRow["power_commander"]; n != 1 {
		t.Errorf("invented by row = %v, want one for the bracket row", res.Coverage.InventedByRow)
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

// TestNilHintsAreSafe keeps the nil-source guard. A deployment with no
// card index answers every fact source call with nil, and no fact may
// claim anything from it.
//
// The MissingCommander case left with D-226, which retired the not-owned
// row and its fact.
func TestNilHintsAreSafe(t *testing.T) {
	var nilHints *CandidateHints
	if thin, n := nilHints.ThinTheme("lifegain"); thin || n != 0 {
		t.Errorf("a nil hint source reported a thin theme: %v %d", thin, n)
	}
	if nilHints.OwnedThemeCount("lifegain") != 0 {
		t.Error("a nil hint source counted owned theme cards")
	}
}

// TestHintsAreAskedOnlyForWhatARowNames is the first half of D-82. An
// eager call runs a whole PR-6 build for a row that holds no
// placeholder.
func TestHintsAreAskedOnlyForWhatARowNames(t *testing.T) {
	c := load(t)
	cases := []struct {
		row            string
		wantCommanders int
	}{
		{"format", 0},
		{"power_commander", 0},
		{"commander", 0},
		{"colors", 0},
		{"commander_pick", 1},
	}
	for _, tc := range cases {
		t.Run(tc.row, func(t *testing.T) {
			row, ok := c.Row(tc.row)
			if !ok {
				t.Fatalf("no row %q", tc.row)
			}
			h := &fakeHints{commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}}
			resolve(row, themedState("lifegain"), h)
			if h.commanderCalls != tc.wantCommanders {
				t.Errorf("row %q asked for commanders %d times, want %d", tc.row, h.commanderCalls, tc.wantCommanders)
			}
		})
	}
}

// TestHintCacheKeyCarriesTheColors is the second half of D-82. One key
// per theme keeps the colorless commander list after the user names
// their colors, and every offered name can sit outside the color
// identity.
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

// TestThinThemeCacheCarriesTheColors is D-82 for ThinTheme. A cache keyed
// by the theme alone serves a stale count, and it must not overwrite
// OnThemeOwned as a side effect. The cache keys on the same key as every
// other hint, and the count reaches the {n} clause through the cache.
func TestThinThemeCacheCarriesTheColors(t *testing.T) {
	h := &CandidateHints{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Colors: []mtgv1.Color{mtgv1.Color_COLOR_W}, OnThemeOwned: 40}
	h.cacheThin(h.key("lifegain"), true, 12)
	if thin, n := h.ThinTheme("lifegain"); !thin || n != 12 {
		t.Errorf("a cached answer was not served: %v %d", thin, n)
	}
	if n := h.OwnedThemeCount("lifegain"); n != 12 {
		t.Errorf("owned count = %d, want the 12 the count measured", n)
	}
	if h.OnThemeOwned != 40 {
		t.Errorf("OnThemeOwned = %d, the cache overwrote the caller's value", h.OnThemeOwned)
	}
	// Another color set is another key. With no builder the miss answers
	// nothing, which proves the white answer was not served for green.
	h.Colors = []mtgv1.Color{mtgv1.Color_COLOR_G}
	if thin, n := h.ThinTheme("lifegain"); thin || n != 0 {
		t.Errorf("the white answer was served for green: %v %d", thin, n)
	}
	if n := h.OwnedThemeCount("lifegain"); n != 40 {
		t.Errorf("owned count under green = %d, want the caller's 40", n)
	}
}
