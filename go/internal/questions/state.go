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
// refinement rows (jank, table_tolerance, budget_scope,
// house_format_limits, named_card_role). The map is free-form by design.
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
	// CurrentOffer are the names on the table now. The pick row repeats
	// with these same names until the user asks for others. The gate run
	// of 2026-08-25 named three others every turn, which read as if the
	// agent had ignored the answer.
	CurrentOffer []string
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
// A named commander closes all three: the pick row and the role row ask
// the same thing in other words.
var commanderKeys = []string{"commander", "commander_pick", "named_card_role"}

// SetCommander records a commander the user named. It closes every
// commander row, and it drops the name from the locked list.
func (s *State) SetCommander(name string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return
	}
	if !hasName(s.CommanderNames, name) {
		s.CommanderNames = append(s.CommanderNames, name)
	}
	s.Ctx.CommanderSet = true
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

// AddLocked records a card the user wants in the deck.
func (s *State) AddLocked(name string) {
	name = strings.TrimSpace(name)
	if name == "" || hasName(s.LockedNames, name) {
		return
	}
	s.LockedNames = append(s.LockedNames, name)
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

// hasName reports whether list holds name, ignoring case and space.
func hasName(list []string, name string) bool {
	name = strings.ToLower(strings.TrimSpace(name))
	for _, n := range list {
		if strings.ToLower(strings.TrimSpace(n)) == name {
			return true
		}
	}
	return false
}

// NewState makes an empty state for a new session.
func NewState(hasCollection bool) *State {
	return &State{
		Slots: &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{}},
		Ctx: Context{
			Filled: map[string]bool{},
			Asked:  map[string]bool{},
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
	s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_FILLED
	s.fillAsk(key)
}

// Skip marks a key the user declined. A default applies, and the agent
// does not ask again. The M-4 record counts a decline as filled: the
// question did its work, and the slot is closed.
func (s *State) Skip(key string) {
	s.Ctx.Filled[key] = true
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
func (s *State) MarkAsked(rowID, key string) {
	s.Ctx.Asked[rowID] = true
	if s.Slots.SlotStates[key] == mtgv1.SlotState_SLOT_STATE_UNSPECIFIED {
		s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_ASKED
	}
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
	"pioneer":   mtgv1.FormatId_FORMAT_ID_PIONEER,
	"modern":    mtgv1.FormatId_FORMAT_ID_MODERN,
	"legacy":    mtgv1.FormatId_FORMAT_ID_LEGACY,
	"vintage":   mtgv1.FormatId_FORMAT_ID_VINTAGE,
	"pauper":    mtgv1.FormatId_FORMAT_ID_PAUPER,
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

var sixtySteps = map[string]mtgv1.SixtyStep{
	"casual":     mtgv1.SixtyStep_SIXTY_STEP_CASUAL,
	"fnm":        mtgv1.SixtyStep_SIXTY_STEP_FNM,
	"tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
}
