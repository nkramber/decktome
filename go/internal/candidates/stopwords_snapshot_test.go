package candidates

import (
	"maps"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestFillerReachesNoThemeSnapshot is D-1188 on the real card pool. A
// filler word that matches no card asks the theme row about it, and one
// that matches card text widens the shortlist past the theme. So each
// phrasing must leave no word unmatched, and it must put the same count
// of cards on theme as its theme words alone.
func TestFillerReachesNoThemeSnapshot(t *testing.T) {
	idx := snapshotIndex(t)
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	onTheme := func(theme string) (int, []string) {
		list, err := b.Build(idx, Request{
			Format:   mtgv1.FormatId_FORMAT_ID_COMMANDER,
			Theme:    theme,
			PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		})
		if err != nil {
			t.Fatalf("build %q: %v", theme, err)
		}
		return list.Stats.OnTheme, list.Theme.Unmatched
	}
	all := maps.Clone(fillerThemes)
	maps.Copy(all, inflectedFillerThemes)
	for theme, bare := range all {
		got, missed := onTheme(theme)
		if len(missed) > 0 {
			t.Errorf("%q leaves %v unmatched, and the theme row would ask", theme, missed)
		}
		if want, _ := onTheme(bare); got != want {
			t.Errorf("%q puts %d cards on theme, want %d as %q does", theme, got, want, bare)
		}
	}
}

// inflectedFillerThemes are the themes of F-230, each with the theme words
// that must survive. The first is the prompt of a live eval. On the code
// before F-230 each one asked the theme row about the -s, -ing, or -ed form
// of a stop word, on the snapshot of 2026-09-04.
var inflectedFillerThemes = map[string]string{
	"making lots of treasure tokens and using them to win": "treasure tokens using win",
	"building around dragons":                              "dragons",
	"a deck that wanted lots of elves":                     "elves",
	"a deck that makes lots of treasure":                   "treasure",
	"builds around treasure tokens":                        "treasure tokens",
	"needing lots of zombies":                              "zombies",
	"focusing on lifegain":                                 "lifegain",
	"liked goblins":                                        "goblins",
	"giving my creatures counters":                         "creatures counters",
}

// TestStopWordsNameNoSubtypeSnapshot reads every subtype of a card that
// Commander permits. A stop word that is a subtype would drop a typal
// theme (D-1188).
func TestStopWordsNameNoSubtypeSnapshot(t *testing.T) {
	idx := snapshotIndex(t)
	sub := map[string]bool{}
	for _, c := range idx.All() {
		if c.GetLegalities()["commander"] != mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL {
			continue
		}
		lines := []string{c.GetTypeLine()}
		for _, f := range c.GetFaces() {
			lines = append(lines, f.GetTypeLine())
		}
		for _, l := range lines {
			if _, after, ok := strings.Cut(l, "—"); ok {
				for _, w := range strings.Fields(strings.ToLower(after)) {
					sub[w] = true
				}
			}
		}
	}
	for _, list := range []map[string]bool{englishStopWords, requestFiller} {
		for w := range list {
			if sub[w] || sub[singular(w)] {
				t.Errorf("stop word %q is a subtype", w)
			}
		}
	}
}
