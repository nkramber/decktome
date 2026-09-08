package agentsvc

import (
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/questions"
)

// emptyClassify answers every field empty, so the classifier reads no
// value out of the message. It is the model that leaves a slot asked.
func emptyClassify(t *testing.T) []llm.Step {
	t.Helper()
	return []llm.Step{classifyJSON(t, map[string]any{})}
}

// TestATurnIsNeverSilent is D-598. Session oUZMC0F2vHe7GGl24LIP read a
// turn that asked nothing, built nothing, and said nothing. The net of
// D-351 holds an open question for StallGrace turns, so the session was
// not dead, and the reader had no way to know that. A screen that never
// changes reads as an app that broke.
func TestATurnIsNeverSilent(t *testing.T) {
	store := newFakeStore()
	// A question is out, and it went out this turn, so the net still
	// waits on it.
	snap := questions.Snapshot{Version: questions.SnapshotVersion}
	// The planner has nothing left to ask, and one question is still
	// out. That pair is the shape of the turn that says nothing.
	cat, err := questions.Load()
	if err != nil {
		t.Fatal(err)
	}
	snap.Ctx.Filled = map[string]bool{}
	snap.Ctx.Asked = map[string]bool{}
	for _, row := range cat.Rows {
		snap.Ctx.Filled[row.StateKey()] = true
		snap.Ctx.Asked[row.ID] = true
	}
	snap.Asks = []questions.Ask{{QuestionID: "q1-power_commander", RowID: "power_commander", Slot: "power", Key: "power", Turn: 1}}
	snap.Turn = 1
	store.sessions["s1"] = &mtgv1.Session{
		Id: "s1",
		Slots: &mtgv1.Slots{
			Format:     &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
			SlotStates: map[string]mtgv1.SlotState{"power": mtgv1.SlotState_SLOT_STATE_ASKED},
		},
		Status: mtgv1.SessionStatus_SESSION_STATUS_ASKING,
		Turns: []*mtgv1.Turn{{Questions: []*mtgv1.Question{
			{Id: "q1-power_commander", Text: "Which power bracket should the deck target?"},
		}}},
	}
	store.states["s1"] = snap

	client, _ := testServerOpts(t, store, nil, emptyClassify(t)...)
	got := chat(t, client, &mtgv1.ChatRequest{SessionId: "s1", Message: "4 optimized"})

	// The turn asked nothing and built nothing. It must still say
	// something, or the reader waits on a screen that never changes.
	silent := len(got.questions) == 0 && got.deck == nil && len(got.statuses) == 0 && len(got.texts) == 0
	if silent {
		t.Fatal("the turn asked nothing, built nothing, and said nothing: the reader reads a broken app")
	}
	// The turn moves: it asks the question again (D-599), it builds, or
	// it names the state. Which of the three is the turn's business. The
	// invariant is that it never ends with nothing at all.
	t.Logf("the turn answered with questions=%d deck=%v statuses=%q", len(got.questions), got.deck != nil, got.statuses)
}

// TestASecondStallSaysSoAndGoesOn is D-599. The question goes out once
// more, and a reader who answers it the same way twice must not wait on
// a chat that says nothing. The net skips the key, and the turn names it.
func TestASecondStallSaysSoAndGoesOn(t *testing.T) {
	store := newFakeStore()
	cat, err := questions.Load()
	if err != nil {
		t.Fatal(err)
	}
	snap := questions.Snapshot{Version: questions.SnapshotVersion}
	snap.Ctx.Filled = map[string]bool{}
	snap.Ctx.Asked = map[string]bool{}
	// The key was re-asked already, so the net is the only way on.
	snap.Ctx.Reasked = map[string]bool{"power": true}
	for _, row := range cat.Rows {
		snap.Ctx.Filled[row.StateKey()] = true
		snap.Ctx.Asked[row.ID] = true
	}
	snap.Asks = []questions.Ask{{QuestionID: "q1-power_commander", RowID: "power_commander", Slot: "power", Key: "power", Turn: 1}}
	snap.Turn = 4
	store.sessions["s1"] = &mtgv1.Session{
		Id: "s1",
		Slots: &mtgv1.Slots{
			Format:     &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
			SlotStates: map[string]mtgv1.SlotState{"power": mtgv1.SlotState_SLOT_STATE_ASKED},
		},
		Status: mtgv1.SessionStatus_SESSION_STATUS_ASKING,
		Turns:  []*mtgv1.Turn{{Questions: []*mtgv1.Question{{Id: "q1-power_commander", Text: "Which power bracket?"}}}},
	}
	store.states["s1"] = snap

	client, _ := testServerOpts(t, store, nil, emptyClassify(t)...)
	got := chat(t, client, &mtgv1.ChatRequest{SessionId: "s1", Message: "the strongest one"})

	if len(got.questions) == 0 && got.deck == nil && len(got.statuses) == 0 && len(got.texts) == 0 {
		t.Fatal("a second stall said nothing at all")
	}
	var told bool
	for _, s := range got.statuses {
		if strings.Contains(s, "did not read an answer to every question") {
			told = true
		}
	}
	if !told {
		t.Errorf("the net skipped a key and did not say so: %q", got.statuses)
	}
}
