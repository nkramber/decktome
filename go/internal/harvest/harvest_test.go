package harvest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/feedback"
)

func at(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

// TestAVerdictWithNoTargetStillHarvests is a gate item of PR-28a. Both
// real verdicts of 2026-09-09 named a session and a deck the store no
// longer held, and both predate the snapshot of D-635. A harvest that
// drops them loses the only feedback this app has.
func TestAVerdictWithNoTargetStillHarvests(t *testing.T) {
	old := feedback.Item{
		ID: "f-old", Kind: "card", Verdict: "down", UID: "u1",
		DeckID: "u8FV7fc98qzNvRfsuJ5q", OracleID: "o-aragorn",
		Text:      "I specifically asked for Aragorn",
		CreatedAt: at("2026-09-08T03:50:00Z"),
	}
	rec, err := RecordOf(old)
	if err != nil {
		t.Fatalf("RecordOf: %v", err)
	}
	if rec.Context != "absent" {
		t.Errorf("context = %q, want absent", rec.Context)
	}
	if rec.Text == "" || rec.DeckID == "" {
		t.Error("the harvest dropped the words or the deck id of a verdict with no snapshot")
	}
	doc := Document(at("2026-09-09T00:00:00Z"), []Record{rec})
	if !strings.Contains(doc, "absent") {
		t.Error("the document does not say the context is absent")
	}
}

// TestNoFileHoldsAnEmail is a gate item of PR-28a (D-559). No proto of
// this repo carries an email but the invite service, so the guard holds
// by shape. This reads the written bytes so a later field can not break
// it quietly.
func TestNoFileHoldsAnEmail(t *testing.T) {
	root := t.TempDir()
	item := feedback.Item{
		ID: "f1", Kind: "question", Verdict: "down", UID: "KyjxInSjjfgihUhWLLeD4anjBz93",
		SessionID: "s1", QuestionText: "Which set do you mean?",
		Reasons: []string{"already_answered"}, CreatedAt: at("2026-09-08T03:45:00Z"),
		Session: &mtgv1.Session{Id: "s1"},
	}
	rec, err := RecordOf(item)
	if err != nil {
		t.Fatalf("RecordOf: %v", err)
	}
	doc, jsonl, err := Write(root, at("2026-09-09T00:00:00Z"), []Record{rec})
	if err != nil {
		t.Fatalf("Write: %v", err)
	}
	for _, p := range []string{doc, jsonl} {
		raw, err := os.ReadFile(p)
		if err != nil {
			t.Fatalf("read %s: %v", p, err)
		}
		if strings.Contains(string(raw), "@") {
			t.Errorf("%s holds an at sign, so it may hold an email", filepath.Base(p))
		}
		if !strings.Contains(string(raw), "KyjxInSjjfgihUhWLLeD4anjBz93") {
			t.Errorf("%s lost the uid, and the loop needs it", filepath.Base(p))
		}
	}
}

// TestTheWatermarkComesFromTheFiles locks the rule that the documents
// are the only record. A second harvest reads nothing new, and a
// harvest after a newer verdict reads that one alone.
func TestTheWatermarkComesFromTheFiles(t *testing.T) {
	root := t.TempDir()
	if mark, err := Watermark(root); err != nil || !mark.IsZero() {
		t.Fatalf("an empty root reads %v, %v, want the zero time", mark, err)
	}
	first := Record{ID: "f1", CreatedAt: at("2026-09-08T03:45:00Z"), Kind: "question", Verdict: "down", Context: "absent"}
	if _, _, err := Write(root, at("2026-09-09T00:00:00Z"), []Record{first}); err != nil {
		t.Fatalf("Write: %v", err)
	}
	mark, err := Watermark(root)
	if err != nil {
		t.Fatalf("Watermark: %v", err)
	}
	if !mark.Equal(first.CreatedAt) {
		t.Errorf("watermark = %v, want %v", mark, first.CreatedAt)
	}
}

// TestASecondHarvestOfADayTakesItsOwnName holds D-65: a run never
// writes over a document that exists.
func TestASecondHarvestOfADayTakesItsOwnName(t *testing.T) {
	root := t.TempDir()
	day := at("2026-09-09T00:00:00Z")
	rec := Record{ID: "f1", CreatedAt: at("2026-09-08T03:45:00Z"), Kind: "deck", Verdict: "up", Context: "absent"}
	doc1, run1, err := Write(root, day, []Record{rec})
	if err != nil {
		t.Fatalf("first: %v", err)
	}
	rec2 := Record{ID: "f2", CreatedAt: at("2026-09-09T04:00:00Z"), Kind: "deck", Verdict: "up", Context: "absent"}
	doc2, run2, err := Write(root, day, []Record{rec2})
	if err != nil {
		t.Fatalf("second: %v", err)
	}
	if doc1 == doc2 || run1 == run2 {
		t.Fatal("the second harvest of a day wrote over the first")
	}
	raw, err := os.ReadFile(run1)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"f1"`) {
		t.Error("the first harvest lost its item")
	}
}

// TestAnEmptyHarvestWritesNothing keeps an empty document out of the
// directory, because it tells a later watermark nothing.
func TestAnEmptyHarvestWritesNothing(t *testing.T) {
	root := t.TempDir()
	doc, jsonl, err := Write(root, at("2026-09-09T00:00:00Z"), nil)
	if err != nil || doc != "" || jsonl != "" {
		t.Fatalf("Write of nothing = %q, %q, %v", doc, jsonl, err)
	}
	if entries, _ := os.ReadDir(filepath.Join(root, Dir)); len(entries) != 0 {
		t.Error("an empty harvest made a file")
	}
}

// TestTheSnapshotReachesTheJSONL proves the triage of PR-28b can read
// the deck list, which is the whole reason for D-635.
func TestTheSnapshotReachesTheJSONL(t *testing.T) {
	item := feedback.Item{
		ID: "f1", Kind: "card", Verdict: "down", UID: "u1", DeckID: "d1",
		CreatedAt: at("2026-09-09T04:00:00Z"),
		Deck: &mtgv1.Deck{
			Id:    "d1",
			Cards: []*mtgv1.DeckCard{{OracleId: "o-sol"}, {OracleId: "o-karlov"}},
		},
	}
	rec, err := RecordOf(item)
	if err != nil {
		t.Fatalf("RecordOf: %v", err)
	}
	if rec.Context != "snapshot" {
		t.Fatalf("context = %q, want snapshot", rec.Context)
	}
	var deck map[string]any
	if err := json.Unmarshal(rec.Deck, &deck); err != nil {
		t.Fatalf("the deck snapshot is not JSON: %v", err)
	}
	cards, _ := deck["cards"].([]any)
	if len(cards) != 2 {
		t.Errorf("the JSONL holds %d cards, want 2", len(cards))
	}
}
