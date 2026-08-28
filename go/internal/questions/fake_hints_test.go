package questions

import (
	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
)

// fakeHints is the one hint source the tests share. Each field answers
// one interface the agent reads, and a zero field answers nothing, so a
// test sets only what it measures.
//
// Five fakes did this work before the audit of 2026-08-28, one per
// interface. A change to Hints then touched all five.
type fakeHints struct {
	// commanders are the names the source offers first. second replaces
	// them once the first set is retired, which is what the real source
	// does (D-73). With no second set, the retired names are filtered
	// out of the first.
	commanders []string
	second     []string
	// identity maps a name onto its color identity. An unknown name
	// answers "not known", as the index does (D-153).
	identity map[string][]mtgv1.Color
	// owned is the on-theme count for the {n} clause. thin and count are
	// the FactSource answers (D-63).
	owned int
	thin  bool
	count int
	// saw are the colors the agent handed over inside the turn (D-124).
	saw []mtgv1.Color
	// commanderCalls counts how often the resolver asked for names.
	commanderCalls int
}

func (f *fakeHints) Commanders(_ string, skip []string) []string {
	f.commanderCalls++
	if len(skip) > 0 && f.second != nil {
		return f.second
	}
	var out []string
	for _, name := range f.commanders {
		if hasName(skip, name) {
			continue
		}
		out = append(out, name)
		if len(out) == 3 {
			break
		}
	}
	return out
}

// OwnedThemeCount answers the count ThinTheme measured when no owned
// count is set, which is what the real source does (M-6).
func (f *fakeHints) OwnedThemeCount(string) int {
	if f.owned != 0 {
		return f.owned
	}
	return f.count
}

func (f *fakeHints) ThinTheme(string) (bool, int) { return f.thin, f.count }

func (f *fakeHints) FitsColors(name string, colors []mtgv1.Color) (bool, bool) {
	id, ok := f.identity[name]
	if !ok {
		return false, false
	}
	return candidates.IdentityMatches(id, colors), true
}

func (f *fakeHints) UseSlots(_ mtgv1.FormatId, colors []mtgv1.Color, _ mtgv1.PoolRule) {
	f.saw = colors
}
