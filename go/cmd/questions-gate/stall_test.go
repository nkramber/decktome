package main

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// slots builds one slot map for the table below.
func slots(states map[string]mtgv1.SlotState) *mtgv1.Slots {
	return &mtgv1.Slots{SlotStates: states}
}

const (
	stAsked   = mtgv1.SlotState_SLOT_STATE_ASKED
	stFilled  = mtgv1.SlotState_SLOT_STATE_FILLED
	stSkipped = mtgv1.SlotState_SLOT_STATE_SKIPPED
)

// A stalled turn moved nothing at all (D-357). Every other shape must
// pass, or the gate would fail a healthy conversation.
func TestStalledTurn(t *testing.T) {
	openColors := []string{"colors"}
	for _, tc := range []struct {
		name      string
		questions int
		ready     bool
		before    *mtgv1.Slots
		after     *mtgv1.Slots
		open      []string
		want      bool
	}{
		{
			name:   "the dead turn: nothing asked, nothing moved, a question out",
			before: slots(map[string]mtgv1.SlotState{"colors": stAsked, "format": stFilled}),
			after:  slots(map[string]mtgv1.SlotState{"colors": stAsked, "format": stFilled}),
			open:   openColors,
			want:   true,
		},
		{
			name:      "a turn that asks something is never stalled",
			questions: 1,
			before:    slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			after:     slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			open:      openColors,
		},
		{
			name:   "a turn that reports ready is never stalled",
			ready:  true,
			before: slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			after:  slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			open:   openColors,
		},
		{
			name:   "a turn that closes the open key moved",
			before: slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			after:  slots(map[string]mtgv1.SlotState{"colors": stFilled}),
			open:   nil,
		},
		{
			name:   "a turn that skips the open key moved",
			before: slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			after:  slots(map[string]mtgv1.SlotState{"colors": stSkipped}),
			open:   nil,
		},
		{
			name:   "a turn that fills another slot from the message moved",
			before: slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			after:  slots(map[string]mtgv1.SlotState{"colors": stAsked, "power": stFilled}),
			open:   openColors,
		},
		{
			name:   "no question is out, so an idle turn is not a dead end",
			before: slots(map[string]mtgv1.SlotState{"format": stFilled}),
			after:  slots(map[string]mtgv1.SlotState{"format": stFilled}),
			open:   nil,
		},
		{
			name:   "a value changed with no state change still counts as movement",
			before: &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{"colors": stAsked, "theme": stFilled}, Theme: "elves"},
			after:  &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{"colors": stAsked, "theme": stFilled}, Theme: "elfball"},
			open:   openColors,
		},
		{
			name:   "a power value written with no state change counts as movement",
			before: &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{"colors": stAsked}},
			after:  &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{"colors": stAsked}, Power: &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}}},
			open:   openColors,
		},
		{
			name:      "a turn that asks and moves nothing else is not stalled",
			questions: 2,
			before:    slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			after:     slots(map[string]mtgv1.SlotState{"colors": stAsked}),
			open:      openColors,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := stalledTurn(tc.questions, tc.ready, tc.before, tc.after, tc.open); got != tc.want {
				t.Errorf("stalledTurn = %v, want %v", got, tc.want)
			}
		})
	}
}

// The shape of session 0EqqY19J6A4BxCVgsmAE, turn by turn.
func TestStalledTurnReadsTheDeadSession(t *testing.T) {
	// Turn 2 asked two questions, so it moved.
	before2 := slots(map[string]mtgv1.SlotState{"format": stFilled, "colors": stAsked, "budget": stAsked})
	after2 := slots(map[string]mtgv1.SlotState{"format": stFilled, "colors": stAsked, "budget": stAsked, "power": stAsked, "commander": stAsked})
	if stalledTurn(2, false, before2, after2, []string{"budget", "colors", "commander", "power"}) {
		t.Error("turn 2 asked two questions, so it is not stalled")
	}
	// Turn 3 answered the power and the commander, and both closed.
	before3 := after2
	after3 := slots(map[string]mtgv1.SlotState{"format": stFilled, "colors": stAsked, "budget": stAsked, "power": stFilled, "commander": stSkipped})
	if stalledTurn(0, false, before3, after3, []string{"budget", "colors"}) {
		t.Error("turn 3 closed two keys, so it is not stalled")
	}
	// Turn 4 would have moved nothing: the two answered keys never closed.
	if !stalledTurn(0, false, after3, after3, []string{"budget", "colors"}) {
		t.Error("a turn that moves nothing with a question out is the dead end")
	}
}
