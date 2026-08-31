package questions

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

func TestBareNegative(t *testing.T) {
	for _, yes := range []string{"No", "no.", "NOPE", "none", "No preference", "Doesn't matter", "doesn’t matter", "not really", "n/a", "any", "  no  "} {
		if !BareNegative(yes) {
			t.Errorf("%q is a bare negative", yes)
		}
	}
	for _, no := range []string{"", "no red", "none of those", "no more than $50", "not blue", "no, make it green", "nothing under three mana"} {
		if BareNegative(no) {
			t.Errorf("%q names something, so it is not a bare negative", no)
		}
	}
}

func TestInvitesYesNo(t *testing.T) {
	for _, yes := range []string{
		"Do you have a color preference?",
		"Do you have a budget for cards to buy?",
		"Any house rules at your table?",
		"Is there a card you want in the deck?",
	} {
		if !InvitesYesNo(yes) {
			t.Errorf("%q invites a yes or a no", yes)
		}
	}
	for _, no := range []string{
		"Which commander do you want? Name one, or I suggest three.",
		"Which format would you like: Commander, Standard, or Modern?",
		"Which power bracket should the Commander deck target?",
		"Do you have a color preference",
	} {
		if InvitesYesNo(no) {
			t.Errorf("%q does not take a bare yes or no", no)
		}
	}
}

// The shape that killed session 0EqqY19J6A4BxCVgsmAE on 2026-08-31.
func TestDeclineNegative(t *testing.T) {
	// The state comes through the real restore path, so the test uses
	// the maps a live session has.
	newState := func() *State {
		slots := &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{
			"colors":         mtgv1.SlotState_SLOT_STATE_ASKED,
			"budget":         mtgv1.SlotState_SLOT_STATE_ASKED,
			"format":         mtgv1.SlotState_SLOT_STATE_FILLED,
			"commander_pick": mtgv1.SlotState_SLOT_STATE_ASKED,
		}}
		snap := Snapshot{
			Version: SnapshotVersion,
			Asks: []Ask{
				{QuestionID: "q2-colors", Key: "colors"},
				{QuestionID: "q3-budget", Key: "budget"},
				{QuestionID: "q5-commander", Key: "commander_pick"},
			},
		}
		snap.Ctx.Outstanding = map[string]string{"colors": "colors", "budget": "budget", "commander_pick": "commander"}
		return Restore("s1", slots, snap)
	}

	t.Run("a bare no closes a yes-or-no question", func(t *testing.T) {
		s := newState()
		key, ok := s.DeclineNegative("q2-colors", "Do you have a color preference?", "No")
		if !ok || key != "colors" {
			t.Fatalf("closed = %q, %v", key, ok)
		}
		if s.Slots.GetSlotStates()["colors"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
			t.Errorf("colors = %v, want SKIPPED", s.Slots.GetSlotStates()["colors"])
		}
	})

	t.Run("both of the dead session's questions close", func(t *testing.T) {
		s := newState()
		s.DeclineNegative("q2-colors", "Do you have a color preference?", "No")
		s.DeclineNegative("q3-budget", "Do you have a budget for cards to buy?", "No")
		if s.Outstanding() && s.Slots.GetSlotStates()["commander_pick"] != mtgv1.SlotState_SLOT_STATE_ASKED {
			t.Error("a question is out that should have closed")
		}
		for _, key := range []string{"colors", "budget"} {
			if s.Slots.GetSlotStates()[key] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
				t.Errorf("%s = %v, want SKIPPED", key, s.Slots.GetSlotStates()[key])
			}
		}
	})

	t.Run("a commander row never closes on a negative", func(t *testing.T) {
		s := newState()
		if _, ok := s.DeclineNegative("q5-commander", "Do you have a commander in mind?", "None"); ok {
			t.Error("a commander row closed on a bare negative (D-120)")
		}
	})

	t.Run("a question that names no yes or no closes nothing", func(t *testing.T) {
		s := newState()
		if _, ok := s.DeclineNegative("q2-colors", "Which colors do you want?", "No"); ok {
			t.Error("a which-question closed on a bare negative")
		}
	})

	t.Run("an answer that names a value closes nothing", func(t *testing.T) {
		s := newState()
		if _, ok := s.DeclineNegative("q3-budget", "Do you have a budget for cards to buy?", "No more than $50"); ok {
			t.Error("a budget closed with no value read")
		}
	})

	t.Run("a key whose question is not out closes nothing", func(t *testing.T) {
		s := newState()
		s.Slots.SlotStates["colors"] = mtgv1.SlotState_SLOT_STATE_FILLED
		if _, ok := s.DeclineNegative("q2-colors", "Do you have a color preference?", "No"); ok {
			t.Error("a filled key closed again")
		}
	})
}
