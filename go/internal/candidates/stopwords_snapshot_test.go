package candidates

import (
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
	for theme, bare := range fillerThemes {
		got, missed := onTheme(theme)
		if len(missed) > 0 {
			t.Errorf("%q leaves %v unmatched, and the theme row would ask", theme, missed)
		}
		if want, _ := onTheme(bare); got != want {
			t.Errorf("%q puts %d cards on theme, want %d as %q does", theme, got, want, bare)
		}
	}
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
