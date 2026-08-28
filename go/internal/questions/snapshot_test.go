package questions

import (
	"encoding/json"
	"reflect"
	"testing"

	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// fullState is a session with a value in every exported State field.
// TestSnapshotRoundTrip refuses a zero field, so a new field must be
// added here, and the round trip then proves the snapshot carries it.
func fullState() *State {
	st := NewState(true)
	st.SessionID = "sess-1"
	st.Ctx.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
	st.Ctx.Filled["format"], st.Ctx.Asked["theme"] = true, true
	st.Ctx.Outstanding["theme"] = "theme"
	st.Ctx.Words, st.Ctx.Theme = "karlov lifegain", "lifegain"
	st.Ctx.CommanderSet, st.Ctx.NamedCard = true, true
	st.Ctx.UnsupportedFormat, st.Ctx.Precon, st.Ctx.CommanderIllegal = true, true, true
	st.Slots.Format = &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}
	st.Slots.SlotStates["format"] = mtgv1.SlotState_SLOT_STATE_FILLED
	st.NamedCards = []string{"Karlov of the Ghost Council", "Sanguine Bond"}
	st.CommanderNames = []string{"Karlov of the Ghost Council"}
	st.LockedNames = []string{"Sanguine Bond"}
	st.OfferedCommanders = []string{"Vito, Thorn of the Dusk Rose"}
	st.CurrentOffer = []string{"Oloro, Ageless Ascetic"}
	st.OfferAsked = []string{"Oloro, Ageless Ascetic"}
	st.UnsupportedFormatName, st.NearestFormat, st.UnsupportedFormatAsked = "Brawl", "Commander", "Brawl"
	st.PreconName, st.IllegalCommander = "Atraxa, Praetors' Voice", "Lightning Bolt"
	st.AskCount, st.Turn = 2, 2
	st.Messages = []string{"karlov lifegain", "keep sanguine bond"}
	st.Asks = []Ask{{QuestionID: "q1-theme", RowID: "theme", Slot: "theme", Key: "theme", Fit: 0.9, Threshold: 0.35, Turn: 1}}
	return st
}

// TestSnapshotRoundTrip proves the private state survives a store cycle
// (D-74). Without it, a resumed conversation repeats a question.
//
// It reflects over every exported State field. Five fields never reached
// the snapshot before the audit of 2026-08-28, and production restores
// from the snapshot on every turn, so the D-210 guard and the D-113
// precon name only worked in the gate harness (Q-1). A field this test
// does not fill fails it, and a field the snapshot drops fails it too.
func TestSnapshotRoundTrip(t *testing.T) {
	st := fullState()
	v := reflect.ValueOf(*st)
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if !f.IsExported() || f.Name == "Slots" || f.Name == "SessionID" {
			continue
		}
		if v.Field(i).IsZero() {
			t.Fatalf("State.%s has no value in fullState, so this test can not prove the snapshot carries it", f.Name)
		}
	}
	raw, err := json.Marshal(st.Snapshot())
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var snap Snapshot
	if err := json.Unmarshal(raw, &snap); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if snap.Version != SnapshotVersion {
		t.Errorf("version = %d, want %d", snap.Version, SnapshotVersion)
	}
	back := Restore("sess-1", proto.Clone(st.Slots).(*mtgv1.Slots), snap)
	w := reflect.ValueOf(*back)
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if !f.IsExported() || f.Name == "Slots" {
			continue
		}
		if !reflect.DeepEqual(v.Field(i).Interface(), w.Field(i).Interface()) {
			t.Errorf("State.%s did not survive the snapshot:\n got %+v\nwant %+v", f.Name, w.Field(i).Interface(), v.Field(i).Interface())
		}
	}
	if !proto.Equal(back.Slots, st.Slots) {
		t.Errorf("slots did not survive: %v", back.Slots)
	}
}

// TestRestoreReadsAnOldSnapshot opens a version 1 snapshot. The version
// 2 fields restore empty, and the rest survive.
func TestRestoreReadsAnOldSnapshot(t *testing.T) {
	raw := `{"version":1,"ctx":{"format":1,"filled":{"format":true},"asked":{},"outstanding":{},"has_collection":true},
		"named_cards":["Karlov of the Ghost Council"],"commander_names":["Karlov of the Ghost Council"],
		"locked_names":[],"offered_commanders":[],"current_offer":[],"offer_asked":[],"ask_count":1,"turn":1,"asks":[]}`
	var snap Snapshot
	if err := json.Unmarshal([]byte(raw), &snap); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	st := Restore("sess-3", nil, snap)
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER || len(st.CommanderNames) != 1 || st.Turn != 1 {
		t.Errorf("a version 1 snapshot lost its fields: %+v", st)
	}
	if st.UnsupportedFormatName != "" || st.PreconName != "" || len(st.Messages) != 0 {
		t.Errorf("a version 1 snapshot invented version 2 fields: %+v", st)
	}
	if st.Ctx.Outstanding == nil {
		t.Error("restore left the outstanding map nil")
	}
}
