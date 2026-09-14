package questions

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestTheRoleRowWaitsForTheWhichCardRow is F-136. The reader wrote "Grima
// as commander", and two cards carry the name. The which-card row waits
// for the power, and the classifier still reported a named card, so the
// role row asked whether a card with no name should lead the deck.
func TestTheRoleRowWaitsForTheWhichCardRow(t *testing.T) {
	h := &CandidateHints{Index: grimaIndex()}
	out := classifyOut{Format: "commander", Theme: "opponent milling cards", PoolRule: "unknown",
		CommanderNames: []string{"Grima"}}
	out.Facts.NamedCard = true
	a, _ := testAgentHints(t, h, classifyStep(t, out),
		fits(t, "named_card_role", "commander_unresolved", "power_commander", "colors"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "Build a deck focused around opponent milling cards. Grima as commander", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if len(st.CommanderOptions) != 2 {
		t.Fatalf("the which-card row offers %v, want the two Gríma cards", st.CommanderOptions)
	}
	if st.Slots.GetSlotStates()["named_card_role"] == mtgv1.SlotState_SLOT_STATE_ASKED {
		t.Error("the role row asked while the which-card row holds the name")
	}
	for _, q := range res.Questions {
		if strings.Contains(q.GetText(), "one card in the 99") {
			t.Errorf("the turn asked %q", q.GetText())
		}
	}
}
