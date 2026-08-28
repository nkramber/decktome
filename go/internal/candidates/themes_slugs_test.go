package candidates

import (
	"sort"
	"testing"
)

// TestThemeSlugsExist guards defect A of the PR-6 gate (2026-08-24).
// themes.json held 16 slugs that Scryfall Tagger does not have. The
// matcher drops an unknown slug without a message, so a theme silently
// loses its payoff half. The test needs a local snapshot. It skips
// without one, and `make themes-check` runs it with the snapshot.
func TestThemeSlugsExist(t *testing.T) {
	idx := snapshotIndex(t)
	tbl, err := loadThemes()
	if err != nil {
		t.Fatalf("load themes: %v", err)
	}
	tags := idx.Tags()
	names := make([]string, 0, len(tbl.Themes))
	for name := range tbl.Themes {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		row := tbl.Themes[name]
		for _, slug := range append(append([]string(nil), row.PayoffSlugs...), row.Slugs...) {
			if !tags.Has(slug) {
				t.Errorf("theme %q names slug %q, which the snapshot does not have", name, slug)
			}
		}
	}
	roles := make([]string, 0, len(tbl.Roles))
	for role := range tbl.Roles {
		roles = append(roles, role)
	}
	sort.Strings(roles)
	for _, role := range roles {
		for _, slug := range tbl.Roles[role] {
			if !tags.Has(slug) {
				t.Errorf("role %q names slug %q, which the snapshot does not have", role, slug)
			}
		}
	}
}
