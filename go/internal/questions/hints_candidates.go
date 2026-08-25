package questions

import (
	"context"
	"log/slog"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// CandidateHints answers the placeholder values from the card index, so
// the color question states a fact and the commander question names real
// cards (PR-6 feeds PR-7). A zero value answers nothing, which drops the
// clause that needs a value.
type CandidateHints struct {
	Index   *cards.Index
	Builder *candidates.Builder
	Format  mtgv1.FormatId
	Colors  []mtgv1.Color
	Owned   map[string]int32
	Pool    mtgv1.PoolRule
	// OnThemeOwned is PR-6's count for the {n} clause of the thin-theme
	// question. The caller reads it from candidates.Stats.
	OnThemeOwned int
	Log          *slog.Logger

	colors     map[string]string
	commanders map[string][]string
	// The thin-theme count runs the whole PR-6 build, so it runs once per
	// theme and the answer is kept.
	thinDone  map[string]bool
	thin      map[string]bool
	thinCount map[string]int
}

// ThemeColors names the colors a theme is strongest in, for example
// "white and black".
func (h *CandidateHints) ThemeColors(theme string) string {
	if h == nil || h.Index == nil || h.Builder == nil || strings.TrimSpace(theme) == "" {
		return ""
	}
	if v, ok := h.colors[h.key(theme)]; ok {
		return v
	}
	format := h.Format
	if format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	}
	cols, err := h.Builder.ThemeColors(h.Index, format, theme, 100, 0.25)
	if err != nil {
		h.warn("theme colors", err)
		return ""
	}
	out := colorWords(cols)
	if h.colors == nil {
		h.colors = map[string]string{}
	}
	h.colors[h.key(theme)] = out
	return out
}

// Commanders names up to three commanders for the theme. It never names
// one the agent already offered, so a user who answers "none" sees three
// others (D-73).
func (h *CandidateHints) Commanders(theme string, skip []string) []string {
	if h == nil || h.Index == nil || h.Builder == nil || strings.TrimSpace(theme) == "" {
		return nil
	}
	cacheKey := h.key(theme) + "\x00" + strings.Join(skip, "\x00")
	if v, ok := h.commanders[cacheKey]; ok {
		return v
	}
	req := candidates.Request{
		Format:   mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Theme:    theme,
		Colors:   h.Colors,
		PoolRule: h.Pool,
		Owned:    h.Owned,
	}
	list, err := h.Builder.Commanders(h.Index, req, 3+len(skip))
	if err != nil {
		h.warn("commanders", err)
		return nil
	}
	seen := map[string]bool{}
	for _, s := range skip {
		seen[strings.ToLower(strings.TrimSpace(s))] = true
	}
	var names []string
	for _, c := range list {
		name := c.Card.GetName()
		if seen[strings.ToLower(strings.TrimSpace(name))] {
			continue
		}
		names = append(names, name)
		if len(names) == 3 {
			break
		}
	}
	if h.commanders == nil {
		h.commanders = map[string][]string{}
	}
	h.commanders[cacheKey] = names
	return names
}

// OwnedThemeCount is the on-theme owned count PR-6 reported.
func (h *CandidateHints) OwnedThemeCount(string) int {
	if h == nil {
		return 0
	}
	return h.OnThemeOwned
}

func (h *CandidateHints) warn(what string, err error) {
	log := h.Log
	if log == nil {
		log = slog.Default()
	}
	log.WarnContext(context.Background(), "questions: hint failed", "hint", what, "error", err)
}

var colorNames = map[mtgv1.Color]string{
	mtgv1.Color_COLOR_W: "white", mtgv1.Color_COLOR_U: "blue", mtgv1.Color_COLOR_B: "black",
	mtgv1.Color_COLOR_R: "red", mtgv1.Color_COLOR_G: "green",
}

// key is the cache key of one hint. It carries every value the answer
// depends on. A key of the theme alone kept a colorless commander list
// after the user named their colors (the gate run of 2026-08-25).
func (h *CandidateHints) key(theme string) string {
	var b strings.Builder
	b.WriteString(theme)
	b.WriteString("|")
	b.WriteString(h.Format.String())
	b.WriteString("|")
	for _, c := range h.Colors {
		b.WriteString(c.String())
	}
	b.WriteString("|")
	b.WriteString(h.Pool.String())
	return b.String()
}

// colorWords reads a color list as English: "white and black".
func colorWords(cols []mtgv1.Color) string {
	var words []string
	for _, c := range cols {
		if n, ok := colorNames[c]; ok {
			words = append(words, n)
		}
	}
	return englishList(words)
}

// MissingCommander reports whether the collection holds none of the named
// commanders. It answers false when no name is given, when no collection
// is loaded, or when the index does not know the name. A card the index
// can not resolve is not proof that the user does not own it.
func (h *CandidateHints) MissingCommander(names []string) bool {
	if h == nil || h.Index == nil || len(h.Owned) == 0 || len(names) == 0 {
		return false
	}
	for _, name := range names {
		card, ok := h.Index.ByName(strings.TrimSpace(name))
		if !ok {
			return false
		}
		if h.Owned[card.GetOracleId()] > 0 {
			return false
		}
	}
	return true
}

// WeakCommanderPool reports whether an owned mode holds no on-theme
// commander (D-63). PR-6 answers it with a count, so there is no
// threshold to invent: an empty pool is a weak pool.
func (h *CandidateHints) WeakCommanderPool(theme string) bool {
	if h == nil || h.Index == nil || h.Builder == nil || len(h.Owned) == 0 || strings.TrimSpace(theme) == "" {
		return false
	}
	if h.Pool == mtgv1.PoolRule_POOL_RULE_ANY_CARD {
		return false
	}
	pool, err := h.Builder.CommanderPool(h.Index, candidates.Request{
		Format:   mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Theme:    theme,
		Colors:   h.Colors,
		PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		Owned:    h.Owned,
	})
	if err != nil {
		h.warn("weak commander pool", err)
		return false
	}
	return len(pool) == 0
}

// ThinTheme runs the PR-6 count and reports whether the collection holds
// fewer than 30 on-theme cards (D-63). The second value is the count, for
// the {n} clause of the thin-theme question.
//
// The mode is owned-first when the user has not answered the pool
// question, because that is the default with a collection (D-37).
func (h *CandidateHints) ThinTheme(theme string) (bool, int) {
	if h == nil || h.Index == nil || h.Builder == nil || len(h.Owned) == 0 || strings.TrimSpace(theme) == "" {
		return false, 0
	}
	if h.thinDone[theme] {
		return h.thin[theme], h.thinCount[theme]
	}
	pool := h.Pool
	if pool == mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		pool = mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
	}
	list, err := h.Builder.Build(h.Index, candidates.Request{
		Format:   h.Format,
		Theme:    theme,
		Colors:   h.Colors,
		PoolRule: pool,
		Owned:    h.Owned,
	})
	if err != nil {
		h.warn("thin theme", err)
		return false, 0
	}
	if h.thinDone == nil {
		h.thinDone, h.thin, h.thinCount = map[string]bool{}, map[string]bool{}, map[string]int{}
	}
	h.thinDone[theme] = true
	h.thin[theme], h.thinCount[theme] = list.Stats.ThinTheme, list.Stats.OnThemeOwned
	return h.thin[theme], h.thinCount[theme]
}
