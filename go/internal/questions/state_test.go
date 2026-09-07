package questions

import (
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// stalledState is a session with two questions out and one slot filled.
// asked is the turn the two questions went out, and turn is the turn
// that runs now.
func stalledState(asked, turn int) *State {
	s := &State{Slots: &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{
		"colors": mtgv1.SlotState_SLOT_STATE_ASKED,
		"budget": mtgv1.SlotState_SLOT_STATE_ASKED,
		"format": mtgv1.SlotState_SLOT_STATE_FILLED,
	}}}
	s.Ctx.Filled, s.Ctx.Asked = map[string]bool{}, map[string]bool{}
	s.Ctx.Outstanding = map[string]string{"colors": "colors", "budget": "budget"}
	s.Turn = turn
	s.Asks = []Ask{
		{QuestionID: "q1", RowID: "colors", Key: "colors", Slot: "colors", Turn: asked},
		{QuestionID: "q2", RowID: "budget", Key: "budget", Slot: "budget", Turn: asked},
	}
	return s
}

// TestCloseStalled is the way out of a dead conversation (D-351). A
// question that has been out for the whole grace period closes with no
// value, and the session can build.
func TestCloseStalled(t *testing.T) {
	s := stalledState(1, 3)

	closed, waiting := s.CloseStalled()

	if !slices.Equal(closed, []string{"budget", "colors"}) {
		t.Errorf("closed = %v, want [budget colors]", closed)
	}
	if len(waiting) != 0 {
		t.Errorf("waiting = %v, want none", waiting)
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
	if c, _ := s.CloseStalled(); len(c) != 0 {
		t.Error("a second call closed something")
	}
}

// TestCloseStalledWaitsOneTurn is D-386. The reader answers on the turn
// after the question goes out. A net that closes the question on that
// same turn takes the answer away.
//
// Gate run 30 ended 32 of 107 conversations on turn 2 this way, and 31
// of them lost an answer the reader had given. Run 29 ran the same
// conversations without the net and played them out.
func TestCloseStalledWaitsOneTurn(t *testing.T) {
	for _, tc := range []struct {
		name         string
		asked, turn  int
		wantClosed   bool
		wantWaitings bool
	}{
		{"the question went out this turn", 2, 2, false, true},
		{"the reader replies on the next turn", 1, 2, false, true},
		{"the reader said nothing for a whole turn", 1, 3, true, false},
		{"the question has been out for ages", 1, 6, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := stalledState(tc.asked, tc.turn)
			closed, waiting := s.CloseStalled()
			if got := len(closed) > 0; got != tc.wantClosed {
				t.Errorf("closed = %v, want %v", closed, tc.wantClosed)
			}
			if got := len(waiting) > 0; got != tc.wantWaitings {
				t.Errorf("waiting = %v, want %v", waiting, tc.wantWaitings)
			}
			// A question the net holds back stays out, so the reader can
			// still answer it.
			if !tc.wantClosed && !s.Outstanding() {
				t.Error("the net closed a question it should have held back")
			}
		})
	}
}

// TestCloseStalledWithNoRecords closes the key. A session restored
// without its M-4 records can not tell the age of a question, and a
// stall for good is the worse failure (D-351).
func TestCloseStalledWithNoRecords(t *testing.T) {
	s := stalledState(1, 2)
	s.Asks = nil
	closed, waiting := s.CloseStalled()
	if len(closed) != 2 || len(waiting) != 0 {
		t.Errorf("closed = %v, waiting = %v, want both keys closed", closed, waiting)
	}
}

// TestCloseStalledReadsTheNewestQuestion: a key the pick row asked again
// takes the age of the newest question, not the first one.
func TestCloseStalledReadsTheNewestQuestion(t *testing.T) {
	s := stalledState(1, 3)
	s.Asks = append(s.Asks, Ask{QuestionID: "q3", RowID: "colors", Key: "colors", Slot: "colors", Turn: 3})
	closed, waiting := s.CloseStalled()
	if !slices.Equal(closed, []string{"budget"}) {
		t.Errorf("closed = %v, want [budget] alone", closed)
	}
	if !slices.Equal(waiting, []string{"colors"}) {
		t.Errorf("waiting = %v, want [colors]", waiting)
	}
}

// TestTheNetStillRescuesADeadConversation is the case D-351 exists for.
// Session 0EqqY19J6A4BxCVgsmAE answered "No" to two yes-or-no questions,
// and both keys stayed in the asked state. Every later turn ended with
// no question, no message, and no deck.
//
// The grace period of D-386 delays the rescue by one turn. It must not
// remove it.
func TestTheNetStillRescuesADeadConversation(t *testing.T) {
	s := stalledState(1, 1)
	// Turn 1 sent the questions. The net holds back.
	if closed, waiting := s.CloseStalled(); len(closed) > 0 || len(waiting) != 2 {
		t.Fatalf("turn 1: closed %v, waiting %v, want the net to hold back", closed, waiting)
	}
	// Turn 2 is the reader's first reply, and it answers nothing. The
	// net still holds back, because the reader may answer next turn.
	s.Turn = 2
	if closed, waiting := s.CloseStalled(); len(closed) > 0 || len(waiting) != 2 {
		t.Fatalf("turn 2: closed %v, waiting %v, want the net to hold back", closed, waiting)
	}
	if !s.Outstanding() {
		t.Fatal("turn 2 closed a question the reader could still answer")
	}
	// Turn 3 answers nothing either. The conversation can not move, so
	// the net closes both keys and the session builds.
	s.Turn = 3
	closed, waiting := s.CloseStalled()
	if !slices.Equal(closed, []string{"budget", "colors"}) {
		t.Fatalf("turn 3: closed %v, want both keys", closed)
	}
	if len(waiting) != 0 {
		t.Errorf("turn 3: waiting %v, want none", waiting)
	}
	if s.Outstanding() {
		t.Error("a question is still out after the rescue")
	}
}

// TestTheNetKeepsTheAnswerTheReaderGave is conversation 56 of gate run
// 30. The reader answered nothing the classifier could read on turn 2,
// and the net closed the format, the colors, and the budget. Turn 3
// would have answered all three, and it never ran (D-386).
func TestTheNetKeepsTheAnswerTheReaderGave(t *testing.T) {
	s := stalledState(1, 2)
	closed, _ := s.CloseStalled()
	if len(closed) > 0 {
		t.Fatalf("the net closed %v on the reader's first reply", closed)
	}
	// Turn 3 arrives, and the reader answers. The keys close on a value.
	s.Turn = 3
	s.Close("colors")
	s.Close("budget")
	if s.Outstanding() {
		t.Error("a question is still out after the reader answered")
	}
	for _, key := range []string{"colors", "budget"} {
		if s.Slots.GetSlotStates()[key] != mtgv1.SlotState_SLOT_STATE_FILLED {
			t.Errorf("%s = %v, want FILLED and not SKIPPED", key, s.Slots.GetSlotStates()[key])
		}
	}
}
