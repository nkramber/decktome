package main

import (
	"os"
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/questions"
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

// TestTheGateRunsTheD351Net is the fix of gate run 29. The gate ran
// without the net that agentsvc.Chat runs on every turn, so it measured
// an agent production does not have. All 10 dead ends of that run were
// turns the net heals.
//
// The check reads the source of both loops, because the turn loop needs
// a provider and this test must stay free.
func TestTheGateRunsTheD351Net(t *testing.T) {
	gate, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(gate), "st.CloseStalled()") {
		t.Fatal("the gate turn loop does not run CloseStalled, so it measures an agent production does not have (D-351)")
	}
	service, err := os.ReadFile("../../internal/agentsvc/service.go")
	if err != nil {
		t.Fatal(err)
	}
	// Both loops must guard the net on the same condition, or the gate
	// and the product disagree about which turn is stalled.
	const guard = "if !res.Ready && len(res.Questions) == 0 {"
	if !strings.Contains(string(service), guard) {
		t.Fatalf("agentsvc no longer guards the net with %q. Change the gate with it.", guard)
	}
	if !strings.Contains(string(gate), "if !turn.Ready && len(turn.Questions) == 0 {") {
		t.Error("the gate guards the net on another condition than agentsvc does")
	}
}

// TestADeadEndSurvivesTheNet is what the check now means. A dead end is
// a stalled last turn the net could not heal.
func TestADeadEndSurvivesTheNet(t *testing.T) {
	for _, tc := range []struct {
		name   string
		closed []stall
		turns  int
		want   bool
	}{
		{"the net healed the last turn", []stall{{Turn: 3}}, 3, false},
		{"the net healed an earlier turn", []stall{{Turn: 2}}, 3, true},
		{"the net closed nothing", nil, 3, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := !closedOnTurn(tc.closed, tc.turns); got != tc.want {
				t.Errorf("dead end = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestTheCommanderIsAnsweredByAnyOfItsRows is the fix of gate run 30.
// Four rows ask for the commander, and each refinement row carries its
// own key. A session that answered the pick row leaves the plain
// "commander" key empty for good. Run 30 called 6 such conversations
// premature, and every one had settled its commander.
func TestTheCommanderIsAnsweredByAnyOfItsRows(t *testing.T) {
	st := questions.NewState(false)
	st.Slots.Format = &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}
	req := required(st, false)
	var commander []string
	for _, keys := range req {
		if len(keys) > 1 {
			commander = keys
		}
	}
	if len(commander) != 4 {
		t.Fatalf("the commander requirement names %v, want its four keys", commander)
	}
	for _, key := range []string{"commander", "commander_pick", "named_card_role", "commander_illegal"} {
		if !slices.Contains(commander, key) {
			t.Errorf("the commander requirement does not name %q", key)
		}
	}
	// A 60-card session needs no commander at all.
	sixty := questions.NewState(false)
	sixty.Slots.Format = &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN}
	for _, keys := range required(sixty, false) {
		if len(keys) > 1 {
			t.Errorf("a 60-card session must need no commander: %v", keys)
		}
	}
}
