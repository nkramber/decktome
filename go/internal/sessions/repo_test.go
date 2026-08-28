package sessions

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

// The store itself needs the Firestore emulator, which no test in this
// repo starts yet. These tests cover the encoding both halves go through,
// which is where a silent data loss would happen.

func sampleSession() *mtgv1.Session {
	return &mtgv1.Session{
		Id:           "sess-1",
		CollectionId: "col-1",
		Status:       mtgv1.SessionStatus_SESSION_STATUS_ASKING,
		CreatedAt:    timestamppb.New(timestamppb.Now().AsTime()),
		Slots: &mtgv1.Slots{
			Theme:  "lifegain",
			Colors: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B},
			Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
			SlotStates: map[string]mtgv1.SlotState{
				"theme":     mtgv1.SlotState_SLOT_STATE_FILLED,
				"commander": mtgv1.SlotState_SLOT_STATE_ASKED,
			},
		},
		Turns: []*mtgv1.Turn{{
			UserMessage: "build me a lifegain deck",
			Questions:   []*mtgv1.Question{{Id: "q1-commander", Slot: "commander", Text: "Do you have a commander in mind?"}},
		}},
		Usage: &mtgv1.Usage{Calls: 3, InputTokens: 950, OutputTokens: 130, Priced: true},
	}
}

func TestSessionRoundTrip(t *testing.T) {
	want := sampleSession()
	payload, err := gzProto(want)
	if err != nil {
		t.Fatalf("encode: %v", err)
	}
	var got mtgv1.Session
	if err := ungzProto(payload, &got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !proto.Equal(want, &got) {
		t.Errorf("session did not survive:\n got %v\nwant %v", &got, want)
	}
}

func TestStateRoundTrip(t *testing.T) {
	st := questions.NewState(true)
	st.SessionID = "sess-1"
	st.Slots.Theme = "lifegain"
	st.AddLocked("Sanguine Bond")
	st.SetCommander("Karlov of the Ghost Council")
	st.SetOffer([]string{"Oloro, Ageless Ascetic"})
	st.RetireOffer()
	st.MarkAsked("commander", "commander", "commander")
	st.Ctx.Words = "build me a lifegain deck"

	payload, err := gzJSON(st.Snapshot())
	if err != nil {
		t.Fatalf("encode: %v", err)
	}
	var got questions.Snapshot
	if err := ungzJSON(payload, &got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	back := questions.Restore("sess-1", st.Slots, got)
	if !back.Ctx.Asked["commander"] || back.Ctx.Words != st.Ctx.Words {
		t.Errorf("the planner context did not survive: %+v", back.Ctx)
	}
	if len(back.CommanderNames) != 1 || len(back.LockedNames) != 1 || len(back.OfferedCommanders) != 1 {
		t.Errorf("the card lists did not survive: %+v", back)
	}
}

// TestEmptyPayloadOpens covers a session stored before the private state
// existed. It must read as an empty snapshot, not as an error.
func TestEmptyPayloadOpens(t *testing.T) {
	var snap questions.Snapshot
	if err := ungzJSON(nil, &snap); err != nil {
		t.Fatalf("an empty payload failed: %v", err)
	}
	if snap.Version != 0 {
		t.Errorf("version = %d, want 0", snap.Version)
	}
}

func TestPutRefusesASessionWithNoID(t *testing.T) {
	r := NewRepo(nil)
	err := r.Put(t.Context(), "u1", &mtgv1.Session{}, questions.Snapshot{}, 0)
	if err == nil || !strings.Contains(err.Error(), "needs an id") {
		t.Errorf("err = %v, want a missing id", err)
	}
}
