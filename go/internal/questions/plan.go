package questions

import (
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// MaxPerTurn is the ceiling from the corpus: never more than three
// questions in one turn.
const MaxPerTurn = 3

// Context is what the planner knows when it picks the next questions.
// The caller fills it from the session slots and from PR-6's stats.
type Context struct {
	// Format is the format slot, when it is filled.
	Format mtgv1.FormatId
	// Filled is keyed by Row.StateKey, not by the proto slot name. A
	// skipped key counts as filled: the user declined and a default
	// applies.
	Filled map[string]bool
	// Asked marks a row id the agent already used. The gate forbids a
	// repeat.
	Asked map[string]bool
	// Words is every word the user has written so far, lowercased. The
	// word-routing rules read it.
	Words string
	// Frozen marks a session whose build run has started (D-68). A frozen
	// session asks nothing.
	Frozen bool

	HasCollection     bool
	OwnedMode         bool
	ThinTheme         bool
	CommanderSet      bool
	NamedCard         bool
	Suggested         bool
	CommanderNotOwned bool
	WeakCommanderPool bool
	PowerCompetitive  bool
	BuyList           bool
	BudgetAmbiguous   bool
	Deadline          bool
	HouseFormat       bool
	TwoPlans          bool
	AfterBuild        bool
	// Theme is the theme slot in the user's words. The salt list reads it.
	Theme string
}

// saltyThemes are the archetypes the corpus marks as salt (section 7).
// A Commander table may refuse them, so the agent asks first.
var saltyThemes = []string{"mill", "land destruction", "stax", "extra turn", "prison"}

// Plan returns the questions to ask this turn, in ask order, at most
// MaxPerTurn. It never returns two rows that inform one proto slot, and
// never a row the session already asked.
func (c *Catalog) Plan(ctx Context) []Row {
	if ctx.Frozen {
		return nil
	}
	var out []Row
	usedKey, usedSlot := map[string]bool{}, map[string]bool{}
	for _, r := range c.Rows {
		if len(out) >= MaxPerTurn {
			break
		}
		key := r.StateKey()
		// One question per proto slot per turn. Two rows that inform one
		// slot read as a contradiction in the same message.
		if ctx.Filled[key] || ctx.Asked[r.ID] || usedKey[key] || usedSlot[r.Slot] {
			continue
		}
		if !r.When.matches(ctx) {
			continue
		}
		out = append(out, r)
		usedKey[key], usedSlot[r.Slot] = true, true
	}
	return out
}

// matches reports whether every trigger of a row holds.
func (w When) matches(ctx Context) bool {
	for _, slot := range w.Requires {
		if !ctx.Filled[slot] {
			return false
		}
	}
	switch w.Format {
	case "commander":
		if ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
			return false
		}
	case "sixty":
		if !sixtyCard(ctx.Format) {
			return false
		}
	case "":
	default:
		return false
	}
	if len(w.Words) > 0 && !anyWord(ctx.Words, w.Words) {
		return false
	}
	if w.SaltyTheme != nil && *w.SaltyTheme != anyWord(strings.ToLower(ctx.Theme), saltyThemes) {
		return false
	}
	facts := []struct {
		want *bool
		have bool
	}{
		{w.PowerCompetitive, ctx.PowerCompetitive},
		{w.NamedCard, ctx.NamedCard},
		{w.Suggested, ctx.Suggested},
		{w.OwnedMode, ctx.OwnedMode},
		{w.CommanderNotOwned, ctx.CommanderNotOwned},
		{w.WeakCommanderPool, ctx.WeakCommanderPool},
		{w.CommanderSet, ctx.CommanderSet},
		{w.HasCollection, ctx.HasCollection},
		{w.ThinTheme, ctx.ThinTheme},
		{w.BuyList, ctx.BuyList},
		{w.BudgetAmbiguous, ctx.BudgetAmbiguous},
		{w.Deadline, ctx.Deadline},
		{w.HouseFormat, ctx.HouseFormat},
		{w.TwoPlans, ctx.TwoPlans},
		{w.AfterBuild, ctx.AfterBuild},
	}
	for _, f := range facts {
		if f.want != nil && *f.want != f.have {
			return false
		}
	}
	return true
}

// sixtyCard reports whether a format builds a 60-card deck. Commander and
// an empty format do not.
func sixtyCard(f mtgv1.FormatId) bool {
	switch f {
	case mtgv1.FormatId_FORMAT_ID_STANDARD, mtgv1.FormatId_FORMAT_ID_PIONEER,
		mtgv1.FormatId_FORMAT_ID_MODERN, mtgv1.FormatId_FORMAT_ID_LEGACY,
		mtgv1.FormatId_FORMAT_ID_VINTAGE, mtgv1.FormatId_FORMAT_ID_PAUPER:
		return true
	}
	return false
}

func anyWord(text string, words []string) bool {
	for _, w := range words {
		if strings.Contains(text, w) {
			return true
		}
	}
	return false
}

// Route maps a user phrase to the slot it belongs to (corpus section 11).
// It returns an empty string when no rule fires. The rules exist because
// the dogfood runs of 2026-08-24 showed three phrase families with no
// home: house rules, power, and jank.
func Route(text string) string {
	t := strings.ToLower(text)
	switch {
	case anyWord(t, []string{"anything goes", "kitchen table", "proxy", "proxies", "no ban list", "whatever"}):
		return "house_rules"
	case anyWord(t, []string{"janky", "jank", "silly", "meme", "for laughs"}):
		return "power"
	case anyWord(t, []string{"strongest", "competitive", "serious", "best deck", "win the event"}):
		return "power"
	}
	return ""
}
