package questions

import (
	"context"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// delegationTurnOne is the first turn of both D-670 delegation tests. The
// reader asks for artifacts and lifegain, and the turn asks the bracket,
// the colors, and the budget.
func delegationTurnOne() classifyOut {
	var out classifyOut
	out.Format, out.Theme, out.PoolRule = "commander", "artifacts and lifegain", "any_card"
	return out
}

// TestADelegationDeclinesOnlyTheSlotItNames walks D-670 through the agent.
// "White and blue. You pick the commander." hands over the commander
// alone, and the classifier declines the bracket and the budget with it.
// Both questions stay out, so the reader can still answer them.
func TestADelegationDeclinesOnlyTheSlotItNames(t *testing.T) {
	first := delegationTurnOne()
	second := first
	second.Colors = []string{"W", "U"}
	second.DeclinedKeys = []string{"power", "budget"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "power_commander", "colors", "budget"), askStep(t),
		classifyStep(t, second), fits(t, "power_commander", "budget"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Commander deck that does artifacts and lifegain.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	for _, k := range []string{"power", "budget"} {
		if st.Slots.GetSlotStates()[k] != mtgv1.SlotState_SLOT_STATE_ASKED {
			t.Fatalf("turn 1 did not ask the %s: %v", k, st.Slots.GetSlotStates())
		}
	}
	if _, err := a.Turn(context.Background(), st, "White and blue. You pick the commander.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	states := st.Slots.GetSlotStates()
	for _, k := range []string{"power", "budget"} {
		if states[k] == mtgv1.SlotState_SLOT_STATE_SKIPPED {
			t.Errorf("a delegation of the commander declined the %s", k)
		}
	}
	if states["commander"] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
		t.Errorf("commander = %v, want skipped: the reader handed it over", states["commander"])
	}
	if st.Ready(a.cat) {
		t.Error("the session calls itself ready before the reader named the bracket and the budget")
	}
}

// TestABareDelegationDeclinesEveryOpenKey guards the other side of D-670.
// "You decide." names no slot, so it hands back every question that is
// out (D-93).
func TestABareDelegationDeclinesEveryOpenKey(t *testing.T) {
	first := delegationTurnOne()
	second := first
	second.DeclinedKeys = []string{"power", "budget"}
	a, _ := testAgent(t,
		classifyStep(t, first), fits(t, "power_commander", "colors", "budget"), askStep(t),
		classifyStep(t, second), fits(t, "colors", "commander"), askStep(t))
	st := NewState(false)
	if _, err := a.Turn(context.Background(), st, "A Commander deck that does artifacts and lifegain.", nil); err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if _, err := a.Turn(context.Background(), st, "You decide.", nil); err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	for _, k := range []string{"power", "budget"} {
		if got := st.Slots.GetSlotStates()[k]; got != mtgv1.SlotState_SLOT_STATE_SKIPPED {
			t.Errorf("%s = %v, want skipped: a bare delegation hands back every open key", k, got)
		}
	}
}
