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
// refinement rows (jank, table_tolerance, budget_scope, acquisition,
// house_format_limits, named_card_role). The map is free-form by design.
type State struct {
	Slots *mtgv1.Slots
	Ctx   Context
	// NamedCards are the cards the user named, newest first. The {card}
	// placeholder reads the first one.
	NamedCards []string
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

// Close marks a key filled and records it in the proto.
func (s *State) Close(key string) {
	s.Ctx.Filled[key] = true
	s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_FILLED
}

// Skip marks a key the user declined. A default applies, and the agent
// does not ask again.
func (s *State) Skip(key string) {
	s.Ctx.Filled[key] = true
	s.Slots.SlotStates[key] = mtgv1.SlotState_SLOT_STATE_SKIPPED
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
