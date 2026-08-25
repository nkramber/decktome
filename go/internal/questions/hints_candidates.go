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
}

// ThemeColors names the colors a theme is strongest in, for example
// "white and black".
func (h *CandidateHints) ThemeColors(theme string) string {
	if h == nil || h.Index == nil || h.Builder == nil || strings.TrimSpace(theme) == "" {
		return ""
	}
	if v, ok := h.colors[theme]; ok {
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
	h.colors[theme] = out
	return out
}

// Commanders names up to three commanders for the theme.
func (h *CandidateHints) Commanders(theme string) []string {
	if h == nil || h.Index == nil || h.Builder == nil || strings.TrimSpace(theme) == "" {
		return nil
	}
	if v, ok := h.commanders[theme]; ok {
		return v
	}
	req := candidates.Request{
		Format:   mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Theme:    theme,
		Colors:   h.Colors,
		PoolRule: h.Pool,
		Owned:    h.Owned,
	}
	list, err := h.Builder.Commanders(h.Index, req, 3)
	if err != nil {
		h.warn("commanders", err)
		return nil
	}
	var names []string
	for _, c := range list {
		names = append(names, c.Card.GetName())
	}
	if h.commanders == nil {
		h.commanders = map[string][]string{}
	}
	h.commanders[theme] = names
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

// colorWords reads a color list as English: "white and black".
func colorWords(cols []mtgv1.Color) string {
	var words []string
	for _, c := range cols {
		if n, ok := colorNames[c]; ok {
			words = append(words, n)
		}
	}
	switch len(words) {
	case 0:
		return ""
	case 1:
		return words[0]
	case 2:
		return words[0] + " and " + words[1]
	default:
		return strings.Join(words[:len(words)-1], ", ") + ", and " + words[len(words)-1]
	}
}
