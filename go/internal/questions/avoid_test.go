package questions

import (
	"context"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestAvoidReachesTheSlot is D-1122. The first outside user wrote that
// the deck held too many artifacts, and the classifier dropped the words.
// The classifier now names the thing in avoid, the slot holds it, and the
// theme stays the plan.
func TestAvoidReachesTheSlot(t *testing.T) {
	out := classifyOut{Format: "commander", Theme: "voltron", PoolRule: "unknown", Avoid: "artifacts"}
	a, _ := testAgentHints(t, &fakeHints{}, classifyStep(t, out),
		fits(t, "power_commander", "colors"), askStep(t))
	st := NewState(true)
	if _, err := a.Turn(context.Background(), st, "A voltron deck. There are too many artifacts.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if got := st.Slots.GetAvoid(); got != "artifacts" {
		t.Errorf("avoid = %q, want artifacts", got)
	}
	if got := st.Slots.GetSlotStates()[SlotAvoid]; got != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("slot state of avoid = %v, want filled", got)
	}
	if got := st.Slots.GetTheme(); got != "voltron" {
		t.Errorf("theme = %q, want voltron", got)
	}
}

// TestAvoidKeepsAHardRequest is D-1122. "No artifacts" asks for none of
// the thing, so the slot keeps "no", and the build lowers every artifact.
// "Less artifacts" from an earlier message stays a soft part of its own.
func TestAvoidKeepsAHardRequest(t *testing.T) {
	out := classifyOut{Format: "commander", Theme: "voltron", PoolRule: "unknown", Avoid: "no artifacts"}
	a, _ := testAgentHints(t, &fakeHints{}, classifyStep(t, out),
		fits(t, "power_commander", "colors"), askStep(t))
	st := NewState(true)
	applyAvoid(st, "artifacts")
	if _, err := a.Turn(context.Background(), st, "A voltron deck, and I don't want any artifacts.", nil); err != nil {
		t.Fatalf("turn: %v", err)
	}
	if got, want := st.Slots.GetAvoid(), "artifacts; no artifacts"; got != want {
		t.Errorf("avoid = %q, want %q", got, want)
	}
}

// TestAvoidAddsAndNeverRepeats is D-1122. A later message adds to the
// slot, and a thing the slot holds in any case is not added twice. An
// empty answer keeps the slot.
func TestAvoidAddsAndNeverRepeats(t *testing.T) {
	st := NewState(true)
	applyAvoid(st, "artifacts")
	applyAvoid(st, " Artifacts ")
	applyAvoid(st, "")
	applyAvoid(st, "creatures with flying")
	if got, want := st.Slots.GetAvoid(), "artifacts; creatures with flying"; got != want {
		t.Errorf("avoid = %q, want %q", got, want)
	}
}

// TestSlotsKeyReadsTheSettingsAlone is D-1118 beside D-1122. The fill
// states and the avoid slot change no setting, and a changed theme does.
func TestSlotsKeyReadsTheSettingsAlone(t *testing.T) {
	base := &mtgv1.Slots{Theme: "voltron", Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}}
	same := &mtgv1.Slots{Theme: "voltron", Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Avoid: "artifacts", SlotStates: map[string]mtgv1.SlotState{"theme": mtgv1.SlotState_SLOT_STATE_FILLED}}
	if SlotsKey(base) != SlotsKey(same) {
		t.Error("a fill state or the avoid slot changed the key")
	}
	other := &mtgv1.Slots{Theme: "voltron, cantrips", Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}}
	if SlotsKey(base) == SlotsKey(other) {
		t.Error("a new theme kept the key")
	}
	st := NewState(true)
	st.Slots = base
	if _, ok := st.SlotsChangedSinceBuild(); ok {
		t.Error("a session with no build reads a record")
	}
	st.MarkBuilt()
	st.Slots = other
	if changed, ok := st.SlotsChangedSinceBuild(); !ok || !changed {
		t.Errorf("changed=%v ok=%v after a new theme, want true true", changed, ok)
	}
}
