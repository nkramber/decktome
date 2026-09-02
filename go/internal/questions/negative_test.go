package questions

import (
	"io"
	"log/slog"
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

// An owned pool rule needs a collection (D-371). Session
// WJbs7FP2csZCULi4SVJu asked for "only from the 'Hobbit' set" with no
// collection at all. The classifier read the word "only" as ownership,
// and the empty rule left no card and no commander to build with.
func TestOwnedPoolRuleNeedsACollection(t *testing.T) {
	newAgent := func(hasCollection bool) (*Agent, *State) {
		a := &Agent{log: slog.New(slog.NewTextHandler(io.Discard, nil))}
		st := NewState(false)
		st.Ctx.HasCollection = hasCollection
		return a, st
	}

	t.Run("owned_only with no collection reads as any card", func(t *testing.T) {
		a, st := newAgent(false)
		a.apply(st, classifyOut{PoolRule: "owned_only"}, nil, "build only from the Hobbit set")
		if got := st.Slots.GetPoolRule(); got != mtgv1.PoolRule_POOL_RULE_ANY_CARD {
			t.Errorf("pool rule = %v, want ANY_CARD", got)
		}
	})

	t.Run("owned_first with no collection reads as any card", func(t *testing.T) {
		a, st := newAgent(false)
		a.apply(st, classifyOut{PoolRule: "owned_first"}, nil, "only artifacts")
		if got := st.Slots.GetPoolRule(); got != mtgv1.PoolRule_POOL_RULE_ANY_CARD {
			t.Errorf("pool rule = %v, want ANY_CARD", got)
		}
	})

	t.Run("a collection keeps the rule the user named", func(t *testing.T) {
		a, st := newAgent(true)
		a.apply(st, classifyOut{PoolRule: "owned_only"}, nil, "only cards I own")
		if got := st.Slots.GetPoolRule(); got != mtgv1.PoolRule_POOL_RULE_OWNED_ONLY {
			t.Errorf("pool rule = %v, want OWNED_ONLY", got)
		}
	})

	t.Run("the key still closes, so the question does not come back", func(t *testing.T) {
		a, st := newAgent(false)
		a.apply(st, classifyOut{PoolRule: "owned_only"}, nil, "only from one set")
		if st.Slots.GetSlotStates()["pool_rule"] == mtgv1.SlotState_SLOT_STATE_ASKED {
			t.Error("the pool question is still out after the rule was read")
		}
	})
}

// TestDeclineCarriesTheFormatDefault is D-406. Session
// ERodNKNOrcwWEFsy1gMa declined the format through the "You decide"
// control, the slot went to SKIPPED with no format, and no power row or
// commander row could fire. The structured path applies the same rules
// as the classifier path now.
func TestDeclineCarriesTheFormatDefault(t *testing.T) {
	slots := &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{
		"format":         mtgv1.SlotState_SLOT_STATE_ASKED,
		"commander_pick": mtgv1.SlotState_SLOT_STATE_ASKED,
	}}
	snap := Snapshot{
		Version: SnapshotVersion,
		Asks: []Ask{
			{QuestionID: "q1-format", Key: "format"},
			{QuestionID: "q5-commander", Key: "commander_pick"},
		},
	}
	snap.Ctx.Outstanding = map[string]string{"format": "format", "commander_pick": "commander"}
	st := Restore("s1", slots, snap)

	if key, ok := st.Decline("q1-format"); !ok || key != "format" {
		t.Fatalf("decline = %q, %v", key, ok)
	}
	// The slot records that the user did not choose, and the corpus
	// default routes the rows that follow.
	if st.Slots.GetSlotStates()["format"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("format state = %v, want SKIPPED", st.Slots.GetSlotStates()["format"])
	}
	if st.Slots.GetFormat().GetId() != DefaultFormat || st.Ctx.Format != DefaultFormat {
		t.Errorf("format = %v (ctx %v), want the corpus default %v", st.Slots.GetFormat().GetId(), st.Ctx.Format, DefaultFormat)
	}
	// A declined pick delegates the commander too (D-147).
	if _, ok := st.Decline("q5-commander"); !ok {
		t.Fatal("the pick did not decline")
	}
	if st.Slots.GetSlotStates()["commander"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("commander state = %v, want SKIPPED with the pick", st.Slots.GetSlotStates()["commander"])
	}
}
