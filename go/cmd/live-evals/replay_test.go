package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

func replaySession() *mtgv1.Session {
	return &mtgv1.Session{Turns: []*mtgv1.Turn{
		{UserMessage: "Build a Ms. Marvel deck | heroic", Questions: []*mtgv1.Question{
			{Id: "q1", Text: "What power?", Options: []string{"Bracket 2", "Bracket 3"}},
			{Id: "q2", Text: "Budget?"},
			{Id: "q3", Text: "Theme?"},
		}},
		{Answers: []*mtgv1.Answer{
			{QuestionId: "q1", OptionIndex: proto.Int32(1)},
			{QuestionId: "q2", Text: " no budget "},
			{QuestionId: "q3", Declined: true},
		}},
		{UserMessage: "Fewer counterspells, please."},
	}}
}

func TestReplayTurnsJoinsAnswersAndMessages(t *testing.T) {
	got := replayTurns(replaySession())
	want := []string{
		"Build a Ms. Marvel deck | heroic",
		"Bracket 3. no budget. Skip that question.",
		"Fewer counterspells, please.",
	}
	if len(got) != len(want) {
		t.Fatalf("turns = %q", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("turn %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestReplayInputWritesTheProbeFiles(t *testing.T) {
	bundle, out := t.TempDir(), filepath.Join(t.TempDir(), "replay")
	if err := writeProto(filepath.Join(bundle, "session.json"), replaySession()); err != nil {
		t.Fatal(err)
	}
	col := &mtgv1.Collection{Entries: []*mtgv1.CollectionEntry{{OracleId: "o1", Quantity: 2}}}
	if err := writeProto(filepath.Join(bundle, "collection.json"), col); err != nil {
		t.Fatal(err)
	}
	if err := replayInput(bundle, out); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(out, "messages.json"))
	if err != nil {
		t.Fatal(err)
	}
	var msgs []string
	if err := json.Unmarshal(b, &msgs); err != nil || len(msgs) != 3 {
		t.Fatalf("messages.json = %s (%v)", b, err)
	}
	cb, err := os.ReadFile(filepath.Join(out, "collection.json"))
	if err != nil {
		t.Fatal(err)
	}
	var back mtgv1.Collection
	if err := protojson.Unmarshal(cb, &back); err != nil || back.GetEntries()[0].GetQuantity() != 2 {
		t.Fatalf("collection.json = %s (%v)", cb, err)
	}
}

func TestReplayInputWithNoCollection(t *testing.T) {
	bundle, out := t.TempDir(), filepath.Join(t.TempDir(), "replay")
	if err := writeProto(filepath.Join(bundle, "session.json"), replaySession()); err != nil {
		t.Fatal(err)
	}
	if err := replayInput(bundle, out); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(out, "collection.json")); !os.IsNotExist(err) {
		t.Fatalf("a collection file was written: %v", err)
	}
}

func TestReplayInputRefusesAnEmptySession(t *testing.T) {
	bundle := t.TempDir()
	if err := writeProto(filepath.Join(bundle, "session.json"), &mtgv1.Session{}); err != nil {
		t.Fatal(err)
	}
	if err := replayInput(bundle, t.TempDir()); err == nil {
		t.Fatal("an empty session was accepted")
	}
}
