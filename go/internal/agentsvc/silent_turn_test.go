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
	var told bool
	for _, s := range got.statuses {
		if strings.Contains(s, "did not read your last answer") {
			told = true
		}
	}
	if len(got.questions) == 0 && got.deck == nil && !told {
		t.Errorf("a turn that asks nothing and builds nothing must name the state: %q", got.statuses)
	}
}
