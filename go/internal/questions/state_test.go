package questions

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// CloseStalled is the way out of a dead conversation (D-351).
func TestCloseStalled(t *testing.T) {
	s := &State{Slots: &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{
		"colors": mtgv1.SlotState_SLOT_STATE_ASKED,
		"budget": mtgv1.SlotState_SLOT_STATE_ASKED,
		"format": mtgv1.SlotState_SLOT_STATE_FILLED,
	}}}
	s.Ctx.Outstanding = map[string]string{"colors": "colors", "budget": "budget"}

	got := s.CloseStalled()

	if len(got) != 2 || got[0] != "budget" || got[1] != "colors" {
		t.Errorf("closed = %v, want [budget colors]", got)
	}
	if s.Outstanding() {
		t.Error("a question is still out")
	}
	for _, key := range []string{"colors", "budget"} {
		if s.Slots.GetSlotStates()[key] != mtgv1.SlotState_SLOT_STATE_SKIPPED {
			t.Errorf("%s = %v, want SKIPPED", key, s.Slots.GetSlotStates()[key])
		}
	}
	if s.Slots.GetSlotStates()["format"] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Error("a filled slot changed")
	}
	if len(s.CloseStalled()) != 0 {
		t.Error("a second call closed something")
	}
}
