package questions

import (
	"fmt"
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
)

// The kinds of gap that ask the gap question, in the order the checks
// run (D-1027).
const (
	GapUnowned = "unowned"
	GapTheme   = "theme"
	GapColors  = "colors"
)

// Gap is the first of the three checks of D-1027 that holds. Kind is
// empty when no check holds.
type Gap struct {
	Kind string
	// Names are the named cards the collection holds no copy of.
	Names []string
	// Theme is the theme the count of GapTheme reads.
	Theme string
	// Have is the owned count, and Want the floor it falls under.
	Have, Want int
}

// Sentence is the first sentence of the gap question. It says why the
// collection can not meet the request. A named card stays in the deck
// under each answer, and the buy list names it (D-1031).
func (g Gap) Sentence() string {
	switch g.Kind {
	case GapUnowned:
		return fmt.Sprintf("Your collection holds no copy of %s, so the buy list names it in each case.", englishList(g.Names))
	case GapTheme:
		return fmt.Sprintf("Your collection holds %d %s cards, and I want %d or more.", g.Have, g.Theme, g.Want)
	case GapColors:
		return fmt.Sprintf("Your collection holds %d cards for this deck, and a deck needs %d or more.", g.Have, g.Want)
	}
	return ""
}

// GapSource runs the three checks of the gap question (D-1027). A hint
// source that holds the card index and the collection implements it.
type GapSource interface {
	PoolGap(theme string, named, commanders []string) Gap
}

// refreshGap reads the gap fact. The gap question asks only when the user
// chose the collection, because that choice means owned cards alone, and
// the question covers a gap (D-1011). An answer to the question closes
// its key, so the checks stop.
func refreshGap(s *State, src FactSource) {
	gs, ok := src.(GapSource)
	if !ok || !s.Ctx.PoolFromReader || s.Ctx.Filled["pool_gap"] ||
		s.Slots.GetPoolRule() != mtgv1.PoolRule_POOL_RULE_OWNED_ONLY ||
		s.Slots.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		s.Ctx.PoolGap, s.Ctx.PoolGapReason = false, ""
		return
	}
	g := gs.PoolGap(s.Slots.GetTheme(), s.NamedCards, s.CommanderNames)
	s.Ctx.PoolGap, s.Ctx.PoolGapReason = g.Kind != "", g.Sentence()
}

// PoolGap runs the three checks of D-1027 in order, and returns the first
// that holds:
//
//  1. A card the request names that the collection holds no copy of. A
//     basic land is always available, and a name the index does not
//     hold is the commander row's job (F-75).
//  2. Fewer owned on-theme cards than ThemeFloor (D-1032).
//  3. Fewer owned nonbasic cards in the deck colors than SetFloor
//     (D-380, D-1032).
//
// A commander the user named sets the colors when the colors slot is
// empty, as the build does, so the counts read the deck that will be
// built. The answer is cached by every value it reads.
func (h *CandidateHints) PoolGap(theme string, named, commanders []string) Gap {
	if h == nil || h.Index == nil || len(h.Owned) == 0 {
		return Gap{}
	}
	key := h.key(theme) + "|" + strings.Join(named, ",") + "|" + strings.Join(commanders, ",")
	if g, ok := h.gaps[key]; ok {
		return g
	}
	g := h.poolGap(theme, named, commanders)
	if h.gaps == nil {
		h.gaps = map[string]Gap{}
	}
	h.gaps[key] = g
	return g
}

func (h *CandidateHints) poolGap(theme string, named, commanders []string) Gap {
	var missing []string
	for _, name := range append(slices.Clone(commanders), named...) {
		card, ok := h.Index.ByName(strings.TrimSpace(name))
		if !ok || candidates.IsBasicLand(card) || h.Owned[card.GetOracleId()] > 0 || hasName(missing, card.GetName()) {
			continue
		}
		missing = append(missing, card.GetName())
	}
	if len(missing) > 0 {
		return Gap{Kind: GapUnowned, Names: missing}
	}

	req := candidates.Request{
		Format:   h.Format,
		Theme:    theme,
		Colors:   h.Colors,
		PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
		Owned:    h.Owned,
		SetCodes: h.SetCodes,
	}
	if len(req.Colors) == 0 && len(commanders) > 0 {
		req.Colors, req.Colorless = h.commanderColors(commanders)
	}
	if strings.TrimSpace(theme) != "" && h.Builder != nil {
		list, err := h.Builder.Build(h.Index, req)
		if err != nil {
			h.warn("pool gap", err)
		} else if want := candidates.ThemeFloor(h.Format); list.Stats.OnThemeOwned < want {
			return Gap{Kind: GapTheme, Theme: theme, Have: list.Stats.OnThemeOwned, Want: want}
		}
	}
	if have, want := candidates.CountOwned(h.Index, req), candidates.SetFloor(h.Format); have < want {
		return Gap{Kind: GapColors, Have: have, Want: want}
	}
	return Gap{}
}

// commanderColors is the union of the color identities of the named
// commanders. A commander set with no color reads as colorless. A name
// the index does not hold leaves the colors open.
func (h *CandidateHints) commanderColors(commanders []string) (colors []mtgv1.Color, colorless bool) {
	for _, name := range commanders {
		card, ok := h.Index.ByName(strings.TrimSpace(name))
		if !ok {
			return nil, false
		}
		for _, c := range card.GetColorIdentity() {
			if !slices.Contains(colors, c) {
				colors = append(colors, c)
			}
		}
	}
	return colors, len(colors) == 0
}
