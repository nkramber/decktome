package candidates

import (
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestCommanderRateLeadsABracketFiveList is D-839 (F-166). A bracket 5
// Commander request with a commander rate keeps an off-theme card that the
// lists of its commander play, and ranks it over a theme card they never
// play. Below bracket 5, and in another format, the rate reads nothing.
func TestCommanderRateLeadsABracketFiveList(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	rate := fixedRates(map[string]float64{"vanilla": 0.9})
	build := func(format mtgv1.FormatId, bracket int32, boost float64) *List {
		t.Helper()
		list, err := b.Build(idx, Request{Format: format, Colors: []mtgv1.Color{W, G}, Theme: "lifegain", Bracket: bracket,
			CommanderRate: rate, ThemeBoost: boost})
		if err != nil {
			t.Fatal(err)
		}
		return list
	}

	list := build(cmdr, 5, 0)
	bears, ok := find(list.Candidates, "Grizzly Bears")
	if !ok {
		t.Fatalf("bracket 5 dropped the card its commander lists play: %v", names(list.Candidates))
	}
	if !slices.Contains(bears.Signals, "commander rate") || !bears.Themed {
		t.Errorf("Grizzly Bears signals %v, themed %v; want the commander rate signal and themed", bears.Signals, bears.Themed)
	}
	want := 0.9*0.7 + popularity(bears.Card, maxRankOf(idx))*0.3
	if bears.Score != want {
		t.Errorf("Grizzly Bears score %v, want %v", bears.Score, want)
	}
	warden, ok := find(list.Candidates, "Soul Warden")
	if !ok {
		t.Fatalf("the theme card left the list: %v", names(list.Candidates))
	}
	if warden.Score >= bears.Score {
		t.Errorf("Soul Warden %v ranks over Grizzly Bears %v, and the rate must lead", warden.Score, bears.Score)
	}

	// The boost scales the theme: a larger boost gives the theme card a
	// larger score, and the rate card keeps its own.
	more, _ := find(build(cmdr, 5, 0.5).Candidates, "Soul Warden")
	if more.Score <= warden.Score {
		t.Errorf("boost 0.5 scores Soul Warden %v, the default %v; want more", more.Score, warden.Score)
	}

	for _, tt := range []struct {
		name    string
		format  mtgv1.FormatId
		bracket int32
	}{
		{"bracket 4", cmdr, 4},
		{"no bracket", cmdr, 0},
		{"Standard", mtgv1.FormatId_FORMAT_ID_STANDARD, 5},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if _, ok := find(build(tt.format, tt.bracket, 0).Candidates, "Grizzly Bears"); ok {
				t.Errorf("%s read the commander rate", tt.name)
			}
		})
	}
}

// TestCommanderRateDefaults is D-839: no rate reads none, and a zero boost
// reads DefaultThemeBoost.
func TestCommanderRateDefaults(t *testing.T) {
	if f, _ := commanderRateOf(Request{Format: cmdr, Bracket: 5}); f != nil {
		t.Error("a request with no rate reads one")
	}
	rate := fixedRates(nil)
	if _, boost := commanderRateOf(Request{Format: cmdr, Bracket: 5, CommanderRate: rate}); boost != DefaultThemeBoost {
		t.Errorf("boost %v, want the default %v", boost, DefaultThemeBoost)
	}
	if _, boost := commanderRateOf(Request{Format: cmdr, Bracket: 5, CommanderRate: rate, ThemeBoost: 0.3}); boost != 0.3 {
		t.Errorf("boost %v, want the override 0.3", boost)
	}
}
