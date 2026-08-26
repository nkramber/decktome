package questions

import (
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// State is one session's question state. Slots hold the values the deck
// generator reads. Ctx holds what the planner needs: which keys are
// closed, which rows were asked, and the facts the rows trigger on.
//
// slot_states carries both proto slot names and the advisory keys of the
// refinement rows (jank, budget_scope, house_format_limits,
// named_card_role, power_confirm). The map is free-form by design.
type State struct {
	Slots *mtgv1.Slots
	Ctx   Context
	// SessionID is the conversation id. It goes out as the provider cache
	// key, so every call of one session routes together.
	SessionID string
	// NamedCards are the cards the user named, newest first. The {card}
	// placeholder reads the first one.
	NamedCards []string
	// CommanderNames are the cards the user named as the commander.
	CommanderNames []string
	// LockedNames are the cards the user named to keep, oldest first. A
	// name that is also a commander name is not a locked card (D-70).
	LockedNames []string
	// OfferedCommanders are the names the agent has retired. A retired
	// name never comes back (D-73).
	OfferedCommanders []string
	// UnsupportedFormatName is the format the user asked for that this app
	// does not build, such as "Brawl". NearestFormat is the format the app
	// builds that is closest to it (D-112).
	UnsupportedFormatName string
	NearestFormat         string
	// PreconName is the precon the user wants to upgrade, named by its
	// commander (D-113).
	PreconName string
	// IllegalCommander is a card the user named as the commander that can
	// not lead a deck (D-129).
	IllegalCommander string
	// CurrentOffer are the names on the table now. The pick row repeats
	// with these same names until the user asks for others. The gate run
	// of 2026-08-25 named three others every turn, which read as if the
	// agent had ignored the answer.
	CurrentOffer []string
	// OfferAsked are the names the pick row sent last. The row asks again
	// only when CurrentOffer differs from these, so a turn that changes
	// nothing gets no second copy of the same question (D-163).
	OfferAsked []string
	// AskCount counts the questions the session has sent. It makes the
	// question id unique, which a repeated row would otherwise break.
	AskCount int
	// Turn counts the turns the session has run, from 1.
	Turn int
	// Asks are the M-4 records, oldest first.
	Asks []Ask
}

// SetOffer records the commander names now on the table.
func (s *State) SetOffer(names []string) {
	if len(names) == 0 {
		return
	}
	s.CurrentOffer = append([]string(nil), names...)
}

// RecordAskedOffer keeps the names a repeat-on-change row just sent. An
// empty list changes nothing, which is the rule SetOffer follows.
func (s *State) RecordAskedOffer(names []string) {
	if len(names) == 0 {
		return
	}
	s.OfferAsked = append([]string(nil), names...)
}

// OfferChanged reports whether the names on the table differ from the
// ones the pick row sent last. A refusal empties the table, and the
// color check of D-153 drops the names the colors exclude. Both cases
// give the user three names they have not seen.
func (s *State) OfferChanged() bool {
	if len(s.CurrentOffer) != len(s.OfferAsked) {
		return true
	}
	for i, name := range s.CurrentOffer {
		if !sameCard(name, s.OfferAsked[i]) {
			return true
		}
	}
	return false
}

// RetireOffer drops the names on the table and never offers them again.
// The agent calls it when the user asks for other names (D-73).
func (s *State) RetireOffer() {
	for _, name := range s.CurrentOffer {
		if name = strings.TrimSpace(name); name != "" && !hasName(s.OfferedCommanders, name) {
			s.OfferedCommanders = append(s.OfferedCommanders, name)
		}
	}
	s.CurrentOffer = nil
}

// commanderKeys are the state keys of the rows that ask for the commander.
// A named commander closes all of them: the pick row and the role row ask
// the same thing in other words, and the illegal row asked for a
// replacement that has now arrived (H-6).
var commanderKeys = []string{"commander", "commander_pick", "named_card_role", "commander_illegal"}

// notOwnedRowLive gates the fact behind the not-owned row (OQ-36, D-207).
const notOwnedRowLive = false

// RefreshFacts reads the planner facts a FactSource answers, from the
// slots as they stand now. agentsvc calls it before the turn, and the
// agent calls it again after the classify call (M-6).
func RefreshFacts(s *State, src FactSource) {
	if src == nil || !s.Ctx.HasCollection {
		return
	}
	// The named commander is not in the collection (corpus section 11).
	// The not-owned row stays dormant until the owner answers OQ-36
	// (D-207). It fired 14 times in gate run 20260826-212512-000 at fit
	// 0.05, and the agent replaced 13 of them with an invented question.
	// The key fix of D-197 stays, so the row is live the day the owner
	// flips this constant.
	if notOwnedRowLive {
		s.Ctx.CommanderNotOwned = src.MissingCommander(s.CommanderNames)
	}
	// No owned commander fits the theme (D-63, D-94). The count answers
	// it, so no threshold is invented.
	if !s.Ctx.CommanderSet {
		s.Ctx.WeakCommanderPool = src.WeakCommanderPool(s.Slots.GetTheme())
	}
	// PR-6 counts the on-theme owned cards (D-63). The count needs the
	// format and the theme, and it only matters while the pool key is
	// open.
	if s.Ctx.Filled["pool_rule"] || s.Slots.GetTheme() == "" ||
		s.Slots.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return
	}
	thin, _ := src.ThinTheme(s.Slots.GetTheme())
	s.Ctx.ThinTheme = thin
}

// SetCommander records a commander the user named. It closes every
// commander row, and it drops the name from the locked list.
func (s *State) SetCommander(name string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return
	}
	s.CommanderNames = mergeName(s.CommanderNames, name)
	s.Ctx.CommanderSet = true
	// A legal commander settles the question the illegal one raised.
	s.Ctx.CommanderIllegal, s.IllegalCommander = false, ""
	for _, key := range commanderKeys {
		s.Close(key)
	}
	// The commander fills the color slot: its color identity is the
	// deck's color identity (corpus section 11). The agent must not ask
	// for colors, and the card-pool question must not wait for an answer
	// that never comes (D-67). PR-8 reads the identity from the card.
	s.Close("colors")
	s.Refresh()
}

// Reopen undoes a closed key, so its row may ask again. The slot state
// goes back to unspecified, which is what the planner reads.
func (s *State) Reopen(key string) {
	delete(s.Ctx.Filled, key)
	delete(s.Ctx.Outstanding, key)
	delete(s.Slots.SlotStates, key)
}

// ClearCommander drops the commander the user chose and reopens every
// commander row. The user asked for another one, and a closed key would
// otherwise leave the agent with nothing to ask (D-130).
//
// The color slot stays closed. The user gave the colors, or the old
// commander did, and a new commander comes from inside them.
func (s *State) ClearCommander() {
	s.CommanderNames = nil
	s.Ctx.CommanderSet = false
	s.Slots.CommanderOracleIds = nil
	// The role row stays closed. The user has just said that this card
	// does not lead the deck, and reopening the row would ask about the
	// card they replaced.
	for _, key := range []string{"commander", "commander_pick"} {
		s.Reopen(key)
	}
	s.RetireOffer()
	s.Ctx.Suggested = true
	s.Refresh()
}

// AddLocked records a card the user wants in the deck, and not as the
// commander. The role of that card is then settled, so the role question
// closes with it.
//
// Gate runs 10 to 13 of 2026-08-25 asked conversation 27 "Do you want
// Grist, the Hunger Tide as your commander, or as one card in the 99?"
// The first message of that conversation reads "but not as my commander".
// The row fired in all four runs (D-70).
func (s *State) AddLocked(name string) {
	if name = strings.TrimSpace(name); name == "" {
		return
	}
	s.LockedNames = mergeName(s.LockedNames, name)
	s.Close("named_card_role")
	s.Refresh()
}

// LockedCards are the named cards that are not the commander. The locked
// row asks about these, and only these. The live run of 2026-08-24 asked
// the user to keep or cut a list that held only their commander (D-70).
func (s *State) LockedCards() []string {
	var out []string
	for _, name := range s.LockedNames {
		if !hasName(s.CommanderNames, name) {
			out = append(out, name)
		}
	}
	return out
}

// Refresh recomputes the facts that come from the card lists. A name that
// becomes the commander must stop the locked question in the same turn.
func (s *State) Refresh() {
	s.Ctx.LockedCard = len(s.LockedCards()) > 0
}

// hasName reports whether list holds the same card as name.
func hasName(list []string, name string) bool {
	for _, n := range list {
		if sameCard(n, name) {
			return true
		}
	}
	return false
}

// normName reads a card name for comparison, without case or edge space.
func normName(s string) string { return strings.ToLower(strings.TrimSpace(s)) }

// baseName is the part of a card name before the first comma. For
// "Grist, the Hunger Tide" it is "grist".
func baseName(s string) string {
	if i := strings.Index(s, ","); i >= 0 {
		return strings.TrimSpace(s[:i])
	}
	return s
}

// sameCard reports whether two strings name one card. A user writes the
// full name once and the short name afterwards, and the classifier
// reports both.
//
// Only a bare first name merges with a full name. Two names that both
// hold a comma stay apart, because "Toph, Hardheaded Teacher" and "Toph,
// the Blind Bandit" are two different cards.
//
// Gate runs 10 to 12 of 2026-08-25 asked "Must the deck keep Grist, the
// Hunger Tide and Grist ...", because the two forms both reached the
// locked list.
func sameCard(a, b string) bool {
	x, y := normName(a), normName(b)
	switch {
	case x == "" || y == "":
		return false
	case x == y:
		return true
	case !strings.Contains(x, ",") && baseName(y) == x:
		return true
	case !strings.Contains(y, ",") && baseName(x) == y:
		return true
	}
	return false
}

// mergeName adds a card name to a list. It keeps the longer form when the
// list already holds the same card, because the full name is the one the
// user can act on.
func mergeName(list []string, name string) []string {
	name = strings.TrimSpace(name)
	if name == "" {
		return list
	}
	for i, n := range list {
		if sameCard(n, name) {
			if len(name) > len(n) {
				list[i] = name
			}
			return list
		}
	}
	return append(list, name)
}

// mergeFront adds a card name to a newest-first list. The {card}
// placeholder reads the first entry, so a repeated name moves to the
// front instead of growing the list.
func mergeFront(list []string, name string) []string {
	name = strings.TrimSpace(name)
	if name == "" {
		return list
	}
	var out []string
	for _, n := range list {
		if sameCard(n, name) {
			if len(n) > len(name) {
				name = n
			}
			continue
		}
		out = append(out, n)
	}
	return append([]string{name}, out...)
}

// NewState makes an empty state for a new session.
func NewState(hasCollection bool) *State {
	return &State{
		Slots: &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{}},
		Ctx: Context{
			Filled:      map[string]bool{},
			Asked:       map[string]bool{},
			Outstanding: map[string]string{},
			// A user with no collection never gets the pool question. The
			// default is any-card (D-37).
			HasCollection: hasCollection,
		},
	}
}

// Close marks a key filled and records it in the proto. It also closes
// the newest open M-4 record for that key, which is how the report knows
// whether an invented question filled its slot.
func (s *State) Close(key string) {
	s.Ctx.Filled[key] = true
	delete(s.Ctx.Outstanding, key)
	s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_FILLED
	s.fillAsk(key)
}

// Skip marks a key the user declined. A default applies, and the agent
// does not ask again. The M-4 record counts a decline as filled: the
// question did its work, and the slot is closed.
func (s *State) Skip(key string) {
	s.Ctx.Filled[key] = true
	delete(s.Ctx.Outstanding, key)
	s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_SKIPPED
	s.fillAsk(key)
}

// fillAsk closes the newest open M-4 record for one key.
func (s *State) fillAsk(key string) {
	for i := len(s.Asks) - 1; i >= 0; i-- {
		if s.Asks[i].Key == key && !s.Asks[i].Filled {
			s.Asks[i].Filled = true
			return
		}
	}
}

// MarkAsked records a question the agent sent. The gate forbids a repeat.
func (s *State) MarkAsked(rowID, key, slot string) {
	s.Ctx.Asked[rowID] = true
	if s.Slots.SlotStates[key] == mtgv1.SlotState_SLOT_STATE_UNSPECIFIED {
		s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_ASKED
	}
	if !s.Ctx.Filled[key] {
		if s.Ctx.Outstanding == nil {
			s.Ctx.Outstanding = map[string]string{}
		}
		s.Ctx.Outstanding[key] = slot
	}
}

// RetireOutstanding drops every key whose question is out with no
// answer. A changed format makes those questions meaningless, and the
// one-row-per-key rule must not block the questions that replace them
// (D-125, D-126).
//
// A retired key leaves the asked state. Ready reads that state, and the
// only other ways out of it are an answer and a decline, so a retired
// question kept the session from ever reporting ready, and it offered
// the dead key to the classifier every turn (H-5). Ctx.Asked keeps the
// row id, so the same row does not ask again. A replacing row may.
func (s *State) RetireOutstanding() {
	for key := range s.Ctx.Outstanding {
		if s.Slots.GetSlotStates()[key] == mtgv1.SlotState_SLOT_STATE_ASKED {
			delete(s.Slots.SlotStates, key)
		}
	}
	s.Ctx.Outstanding = map[string]string{}
}

// Freeze stops every question. A build run has started, so the slot set is
// the deck's record (D-68). A later change starts a new run.
func (s *State) Freeze() { s.Ctx.Frozen = true }

// AddWords keeps every word the user has written. The routing rules and
// the word triggers read it.
func (s *State) AddWords(text string) {
	s.Ctx.Words = strings.TrimSpace(s.Ctx.Words + " " + strings.ToLower(text))
}

// Ready reports whether the session can build. Two things must hold: the
// planner has nothing left to ask, and no question is still out. The
// no-repeat rule empties the plan as soon as a question goes out, so the
// plan alone would call an unanswered session complete.
func (s *State) Ready(c *Catalog) bool {
	if len(c.Plan(s.Ctx)) > 0 {
		return false
	}
	return !s.Outstanding()
}

// Outstanding reports whether a question is out with no answer mapped yet.
func (s *State) Outstanding() bool {
	for _, state := range s.Slots.GetSlotStates() {
		if state == mtgv1.SlotState_SLOT_STATE_ASKED {
			return true
		}
	}
	return false
}

// slotWord normalizes a value the classifier wrote in free text. The
// format and pool-rule fields carry no schema enum, because an enum
// suppressed the field: the model answered "unknown" for a message that
// named the format outright (D-92). Go owns the vocabulary instead.
func slotWord(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	return strings.ReplaceAll(strings.ReplaceAll(s, "-", "_"), " ", "_")
}

// DefaultFormat is the format a declined format slot takes. The corpus
// names it under "default answers": Commander, the most played format in
// 2026. The slot keeps the SKIPPED state, so the deck generator can tell
// a default from a choice (D-98).
const DefaultFormat = mtgv1.FormatId_FORMAT_ID_COMMANDER

// formatIDs maps the classifier's format word to the proto enum.
var formatIDs = map[string]mtgv1.FormatId{
	"commander": mtgv1.FormatId_FORMAT_ID_COMMANDER,
	"standard":  mtgv1.FormatId_FORMAT_ID_STANDARD,
	"modern":    mtgv1.FormatId_FORMAT_ID_MODERN,
	"house":     mtgv1.FormatId_FORMAT_ID_HOUSE,
}

var poolRules = map[string]mtgv1.PoolRule{
	"owned_first": mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
	"owned_only":  mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
	"any_card":    mtgv1.PoolRule_POOL_RULE_ANY_CARD,
}

var colorIDs = map[string]mtgv1.Color{
	"W": mtgv1.Color_COLOR_W, "U": mtgv1.Color_COLOR_U, "B": mtgv1.Color_COLOR_B,
	"R": mtgv1.Color_COLOR_R, "G": mtgv1.Color_COLOR_G,
}

// cedhBracket is the bracket cEDH names. Bracket 5 is competitive
// Commander, and only the ban list limits it (corpus section 2.3).
const cedhBracket = 5

var sixtySteps = map[string]mtgv1.SixtyStep{
	"casual":     mtgv1.SixtyStep_SIXTY_STEP_CASUAL,
	"fnm":        mtgv1.SixtyStep_SIXTY_STEP_FNM,
	"tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
}
