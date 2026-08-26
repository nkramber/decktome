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
	// WantPair asks the candidate builder for two-commander pairs. The
	// user asked for partners or a Background (D-154).
	WantPair bool
	// WantBackground narrows that to pairs that hold a Background.
	WantBackground bool
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

// CanLead reports whether a named card can lead a deck. Probe 41 asks
// for "Commander deck with Lightning Bolt as my commander", and every run
// before 2026-08-26 accepted it in silence (D-129).
//
// It answers "not known" for any legendary card it can not confirm. The
// engine reads a type line, and a type line is wrong about a whole class
// of commander. Grist, the Hunger Tide is "Legendary Planeswalker" and it
// is a legal commander, because a characteristic-defining ability makes
// it a creature card everywhere except the battlefield (Scryfall ruling,
// 2021-06-18). Gate run 14 told a user that Grist can not lead a deck.
// The gate passed and the linter found nothing, and the claim was false.
//
// A confident "no" therefore needs a card that is not legendary at all.
// Lightning Bolt is an instant and Sol Ring is not legendary, so both
// still answer (D-140).
func (h *CandidateHints) CanLead(name string) (canLead, known bool) {
	if h == nil || h.Index == nil {
		return false, false
	}
	card, ok := h.Index.ByName(strings.TrimSpace(name))
	if !ok {
		return false, false
	}
	if card.GetCanBeCommander() || card.GetIsBackground() {
		return true, true
	}
	if strings.Contains(strings.ToLower(card.GetTypeLine()), "legendary") {
		// A legendary card the engine can not confirm. Say nothing.
		return false, false
	}
	return false, true
}

// FitsColors reports whether the named card holds every color named and
// no other. It is the same test CommanderPool applies when it builds the
// pool (D-148), so a name that survives here would be offered again.
//
// An empty color list fits everything: the user has named no colors, so
// nothing is out of them (D-153).
func (h *CandidateHints) FitsColors(name string, colors []mtgv1.Color) (fits, known bool) {
	if h == nil || h.Index == nil {
		return false, false
	}
	card, ok := h.Index.ByName(strings.TrimSpace(name))
	if !ok {
		return false, false
	}
	if len(colors) == 0 {
		return true, true
	}
	return candidates.IdentityMatches(card.GetColorIdentity(), colors), true
}

// UseSlots takes the slot values as they stand inside the turn. The cache
// key carries the format, the colors, and the pool rule, so an answer
// computed under other values stays keyed to those values (D-82).
func (h *CandidateHints) UseSlots(format mtgv1.FormatId, colors []mtgv1.Color, pool mtgv1.PoolRule) {
	if h == nil {
		return
	}
	if format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		h.Format = format
	}
	if len(colors) > 0 {
		h.Colors = colors
	}
	if pool != mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		h.Pool = pool
	}
}

// UseWantPair records that the user asked for a two-commander pair. The
// pool offers pairs on its own when too few singles fit the colors, so
// this only adds the case the words ask for (D-154).
func (h *CandidateHints) UseWantPair(want, background bool) {
	if h == nil || !want {
		return
	}
	h.WantPair = true
	if background {
		h.WantBackground = true
	}
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
	cols, err := h.Builder.ThemeColors(h.Index, format, theme, 100, 0.15)
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
	if h.WantPair {
		cacheKey += "\x00pair"
	}
	if h.WantBackground {
		cacheKey += "\x00background"
	}
	if v, ok := h.commanders[cacheKey]; ok {
		return v
	}
	req := candidates.Request{
		Format:         mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Theme:          theme,
		Colors:         h.Colors,
		PoolRule:       h.Pool,
		Owned:          h.Owned,
		WantPair:       h.WantPair,
		WantBackground: h.WantBackground,
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
		// A pair reads "A + B". The pick row offers it as one choice,
		// because the user chooses a pair and not half of one (D-154).
		name := c.DisplayName()
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
	// The {n} clause of the thin-theme question reads this count, so the
	// source keeps it with the answer (M-6).
	h.OnThemeOwned = h.thinCount[theme]
	return h.thin[theme], h.thinCount[theme]
}
