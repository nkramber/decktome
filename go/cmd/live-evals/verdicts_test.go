package main

import (
	"strings"
	"testing"
	"time"

	"github.com/nkramber/decktome/go/internal/feedback"
)

// D-1149: the key of a verdict reads back, and a deck id is never a
// verdict.
func TestVerdictKeyReadsBack(t *testing.T) {
	id, ok := verdictOf(verdictKey("Ab12Cd34Ef56Gh78Ij90"))
	if !ok || id != "Ab12Cd34Ef56Gh78Ij90" {
		t.Fatalf("verdictOf = %q, %v", id, ok)
	}
	for _, key := range []string{"Ab12Cd34Ef56Gh78Ij90", "v-", ""} {
		if _, ok := verdictOf(key); ok {
			t.Fatalf("verdictOf(%q) read a verdict", key)
		}
	}
}

func TestVerdictRowNamesTheThumbsDown(t *testing.T) {
	at := time.Date(2026, 10, 4, 16, 0, 0, 0, time.UTC)
	row := verdictRow(feedback.Pending{UID: "u1", ID: "f1", Kind: "card", SessionID: "s1", DeckID: "d1", CreatedAt: at}, "user")
	want := pendingRow{UID: "u1", Deck: "v-f1", Session: "s1", Who: "user", Kind: "thumbs-down", Verdict: "f1",
		Target: "card", TargetDeck: "d1", CreatedAt: at}
	if row != want {
		t.Fatalf("row = %+v, want %+v", row, want)
	}
}

// The summary names the reasons, and it prints no text of the reader.
func TestDescribeVerdictPrintsNoNote(t *testing.T) {
	p := feedback.Pending{ID: "f1", Kind: "deck", DeckID: "d1", CreatedAt: time.Date(2026, 10, 4, 16, 0, 0, 0, time.UTC)}
	got := describeVerdict(1, p, "user", feedback.Item{Reasons: []string{"off_theme"}, Text: "too many artifacts"})
	for _, want := range []string{"thumbs down on a deck", "off_theme", "deck d1", "note of 18 characters"} {
		if !strings.Contains(got, want) {
			t.Fatalf("summary misses %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, "artifacts") {
		t.Fatalf("summary printed the note:\n%s", got)
	}
}

// An empty deck id or session id matches no verdict, so a bundle with
// no deck never holds the verdicts of other readers.
func TestVerdictsOfAnEmptyIDMatchesNothing(t *testing.T) {
	all := []feedback.Item{{ID: "a"}, {ID: "b", DeckID: "d1"}, {ID: "c", SessionID: "s1"}}
	if got := verdictsOf(all, "", ""); len(got) != 0 {
		t.Fatalf("verdictsOf with no id = %+v", got)
	}
	if got := verdictsOf(all, "", "s1"); len(got) != 1 || got[0].ID != "c" {
		t.Fatalf("verdictsOf of a session = %+v", got)
	}
}
