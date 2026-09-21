package candidates

import "github.com/nkramber/decktome/go/internal/cards"

// StarvedTheme reads whether an exclusion took the theme out of the
// pool (F-37). list is the shortlist of req, which holds the exclusion,
// and whole is the owned map before it. The theme is starved when list
// is thin on the theme and the whole library is not. It returns the
// owned theme count of the whole library. A library thin on the theme
// before the exclusion is the pool question of D-63, and it reads false.
func (b *Builder) StarvedTheme(idx *cards.Index, req Request, list *List, whole map[string]int32) (int, bool, error) {
	if len(req.ExcludeOracleIDs) == 0 || !list.Stats.ThinTheme {
		return 0, false, nil
	}
	req.Owned, req.ExcludeOracleIDs = whole, nil
	full, err := b.Build(idx, req)
	if err != nil {
		return 0, false, err
	}
	n := full.Stats.OnThemeOwned
	return n, n >= ThinThemeFloor, nil
}
