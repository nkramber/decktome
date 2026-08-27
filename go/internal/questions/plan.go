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
	Format mtgv1.FormatId `json:"format"`
	// Filled is keyed by Row.StateKey, not by the proto slot name. A
	// skipped key counts as filled: the user declined and a default
	// applies.
	Filled map[string]bool `json:"filled"`
	// Asked marks a row id the agent already used. The gate forbids a
	// repeat.
	Asked map[string]bool `json:"asked"`
	// Outstanding maps a state key whose question is out with no answer
	// onto the slot that key informs.
	// No other row may ask that key while it is out. Probe 33 asked the
	// theme through the competitive row, got no answer, and asked it
	// again through the general row one turn later. To the user that
	// reads as the same question in other words (D-126).
	//
	// The commander rows are unaffected. Each one carries its own key, so
	// the pick row still follows the base row (D-71).
	Outstanding map[string]string `json:"outstanding"`
	// Words is every word the user has written so far, lowercased. The
	// word-routing rules read it.
	Words string `json:"words"`
	// Frozen marks a session whose build run has started (D-68). A frozen
	// session asks nothing.
	Frozen bool `json:"frozen"`
	// OfferChanged says the commanders on the table differ from the ones
	// the pick row named last. Only a row with RepeatOnChange reads it
	// (D-163).
	OfferChanged bool `json:"offer_changed"`
	// BadFormatChanged says the unsupported format the user named differs
	// from the one the decline row named last. It is the same rule for
	// the format rows that OfferChanged is for the pick row (D-210).
	BadFormatChanged bool `json:"bad_format_changed"`

	// OutOfScope marks a request for something other than a Magic deck.
	// Nothing else is worth asking until it is settled (D-99).
	OutOfScope    bool `json:"out_of_scope"`
	HasCollection bool `json:"has_collection"`
	OwnedMode     bool `json:"owned_mode"`
	ThinTheme     bool `json:"thin_theme"`
	CommanderSet  bool `json:"commander_set"`
	NamedCard     bool `json:"named_card"`
	// LockedCard marks a named card that is not the commander. The locked
	// row asks about these, and only these (D-70).
	LockedCard        bool `json:"locked_card"`
	Suggested         bool `json:"suggested"`
	CommanderNotOwned bool `json:"commander_not_owned"`
	WeakCommanderPool bool `json:"weak_commander_pool"`
	PowerCompetitive  bool `json:"power_competitive"`
	// PowerInferred says the agent filled the power step itself, because
	// the user asked for a strong deck and named no step. A step the user
	// named is not inferred, and a question about it repeats the answer
	// the user already gave (D-209).
	PowerInferred   bool `json:"power_inferred"`
	BuyList         bool `json:"buy_list"`
	BudgetAmbiguous bool `json:"budget_ambiguous"`
	HouseFormat     bool `json:"house_format"`
	TwoPlans        bool `json:"two_plans"`
	AfterBuild      bool `json:"after_build"`
	// TwoDecks marks a request for more than one deck. The app builds one
	// at a time, and it says so before it asks anything else (D-112).
	TwoDecks bool `json:"two_decks"`
	// UnsupportedFormat marks a format this app does not build, such as
	// Brawl. State holds the name and the nearest format (D-112).
	UnsupportedFormat bool `json:"unsupported_format"`
	// NoNearFormat marks an unsupported format with no nearest format to
	// offer. Historic and Timeless are the two (D-146).
	NoNearFormat bool `json:"no_near_format"`
	// Precon marks a request to upgrade a preconstructed deck (D-113).
	Precon bool `json:"precon"`
	// WantPair marks a request for a two-commander pair (D-154).
	WantPair bool `json:"want_pair"`
	// WantBackground narrows that to a pair that holds a Background.
	WantBackground bool `json:"want_background"`
	// CommanderIllegal marks a named commander that can not lead a deck,
	// such as Lightning Bolt (D-129).
	CommanderIllegal bool `json:"commander_illegal"`
	// Theme is the theme slot in the user's words.
	Theme string `json:"theme"`
}

// Plan returns the questions to ask this turn, in ask order, at most
// MaxPerTurn. It never returns two rows that inform one proto slot, and
// never a row the session already asked.
func (c *Catalog) Plan(ctx Context) []Row {
	if ctx.Frozen {
		return nil
	}
	var out []Row
	usedKey, usedSlot := map[string]bool{}, map[string]bool{}
	// An out-of-scope request gets one question and no others. Asking the
	// format beside it reads as if the agent had not heard the request.
	// Gate run 11 asked "Which Yu-Gi-Oh format would you like?" (D-99).
	if ctx.OutOfScope && !ctx.Filled["scope"] && !ctx.Asked["out_of_scope"] {
		if row, ok := c.Row("out_of_scope"); ok {
			return []Row{row}
		}
	}
	// A two-deck request also gets one question and no others. Every other
	// row belongs to one deck, so the agent settles which deck first.
	// Probe 50 of gate runs 11 to 13 chose a deck silently, and it never
	// said that the app builds one at a time (D-112).
	if ctx.TwoDecks && !ctx.Filled["deck_count"] && !ctx.Asked["one_deck"] {
		if row, ok := c.Row("one_deck"); ok {
			return []Row{row}
		}
	}
	for _, r := range c.Rows {
		if len(out) >= MaxPerTurn {
			break
		}
		key := r.StateKey()
		// One question per proto slot per turn. Two rows that inform one
		// slot read as a contradiction in the same message.
		if ctx.Filled[key] || (ctx.Asked[r.ID] && !r.asksAgain(ctx)) || usedKey[key] || usedSlot[r.Slot] {
			continue
		}
		// Another row already asked this key, and no answer came back.
		if _, out := ctx.Outstanding[key]; out && !ctx.Asked[r.ID] {
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

// asksAgain reports whether a row the session already asked may ask a
// second time.
//
// A plain repeat row always may. A row that narrows the repeat asks again
// only when its content changed. The pick row names three commanders, so
// a repeat with the same three names is the same question in the same
// words. Conversations 1, 77, and 90 of gate run 18 each got one, and the
// eval refused every one of them as a duplicate. The user answered some
// other slot, and the agent read that as a reason to ask again.
//
// This is the D-158 rule for another row: the question is out, it is
// recorded as asked with no answer, and the gate reports it. Silence
// beats the same sentence twice (D-163).
func (r Row) asksAgain(ctx Context) bool {
	switch {
	case !r.Repeat:
		return false
	case r.RepeatOnChange:
		return ctx.contentChanged(r.Slot)
	}
	return true
}

// contentChanged reports whether the content of a repeat-on-change row
// differs from the content that row sent last. The pick row names three
// commanders, and the two decline rows name one format. A slot with no
// signal never repeats, which is the safe answer: silence beats the same
// sentence twice (D-163, D-210).
func (c Context) contentChanged(slot string) bool {
	switch slot {
	case "commander":
		return c.OfferChanged
	case "format":
		return c.BadFormatChanged
	}
	return false
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
	// A trigger word reads through anyPhrase, which refuses a negated
	// match. Probe 38 of gate run 11 wrote "no proxies", and the plain
	// substring test fired the house-rules row on it (D-111).
	if len(w.Words) > 0 && !anyPhrase(ctx.Words, w.Words) {
		return false
	}
	// A row may wait for the answer to another row. The commander row
	// waits for the role question: whether the named card leads the deck
	// decides whether a commander question is needed at all. Probe 39
	// asked both, one turn apart (D-128).
	for _, k := range w.NotOutstanding {
		if _, out := ctx.Outstanding[k]; out {
			return false
		}
	}
	facts := []struct {
		want *bool
		have bool
	}{
		{w.OutOfScope, ctx.OutOfScope},
		{w.PowerCompetitive, ctx.PowerCompetitive},
		{w.PowerInferred, ctx.PowerInferred},
		{w.NamedCard, ctx.NamedCard},
		{w.LockedCard, ctx.LockedCard},
		{w.Suggested, ctx.Suggested},
		{w.OwnedMode, ctx.OwnedMode},
		{w.CommanderNotOwned, ctx.CommanderNotOwned},
		{w.WeakCommanderPool, ctx.WeakCommanderPool},
		{w.CommanderSet, ctx.CommanderSet},
		{w.HasCollection, ctx.HasCollection},
		{w.ThinTheme, ctx.ThinTheme},
		{w.BuyList, ctx.BuyList},
		{w.BudgetAmbiguous, ctx.BudgetAmbiguous},
		{w.HouseFormat, ctx.HouseFormat},
		{w.TwoPlans, ctx.TwoPlans},
		{w.AfterBuild, ctx.AfterBuild},
		{w.TwoDecks, ctx.TwoDecks},
		{w.UnsupportedFormat, ctx.UnsupportedFormat},
		{w.NoNearFormat, ctx.NoNearFormat},
		{w.Precon, ctx.Precon},
		{w.CommanderIllegal, ctx.CommanderIllegal},
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
//
// HOUSE joined the list with D-155. It builds 60 cards by default
// (formats.json), and it was absent here, so no power row could fire for
// it at all: the 60-card rows need this test and the Commander rows need
// format "commander". A user who named no format and asked for no ban
// list would have finished with no power level.
func sixtyCard(f mtgv1.FormatId) bool {
	switch f {
	case mtgv1.FormatId_FORMAT_ID_STANDARD,
		mtgv1.FormatId_FORMAT_ID_MODERN,
		mtgv1.FormatId_FORMAT_ID_HOUSE:
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
	// "Proxy" left this list. A user who proxies has no budget, and the
	// word says nothing about which cards are legal (D-111). "Whatever"
	// left it too. "Whatever is winning" and "whatever you think is best"
	// are not house rules, and they fired the row in three runs.
	case anyPhrase(t, []string{"anything goes", "kitchen table", "no ban list"}):
		return "house_rules"
	case anyPhrase(t, []string{"janky", "jank", "silly", "meme", "for laughs"}):
		return "power"
	case anyPhrase(t, competitiveSigns):
		return "power"
	}
	return ""
}
