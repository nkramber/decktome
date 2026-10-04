package main

import (
	"strings"
	"testing"
	"time"

	"github.com/nkramber/decktome/go/internal/feedback"
)

// D-1156: the key of a note reads back, a note key is no verdict key,
// and the mark finds the feedback id of both.
func TestNoteKeyReadsBack(t *testing.T) {
	id, ok := noteOf(noteKey("Ab12Cd34Ef56Gh78Ij90"))
	if !ok || id != "Ab12Cd34Ef56Gh78Ij90" {
		t.Fatalf("noteOf = %q, %v", id, ok)
	}
	for _, key := range []string{"Ab12Cd34Ef56Gh78Ij90", "n-", "", "v-f1"} {
		if _, ok := noteOf(key); ok {
			t.Fatalf("noteOf(%q) read a note", key)
		}
	}
	if _, ok := verdictOf(noteKey("f1")); ok {
		t.Fatal("a note key reads as a verdict")
	}
	for key, want := range map[string]string{"n-f1": "f1", "v-f2": "f2"} {
		if got, ok := feedbackOf(key); !ok || got != want {
			t.Fatalf("feedbackOf(%q) = %q, %v", key, got, ok)
		}
	}
	if _, ok := feedbackOf("deck1"); ok {
		t.Fatal("a deck id reads as feedback")
	}
}

func TestNoteRowNamesTheNote(t *testing.T) {
	at := time.Date(2026, 10, 4, 16, 0, 0, 0, time.UTC)
	row := noteRow(feedback.Pending{UID: "u1", ID: "f1", Kind: "general", SessionID: "s1", DeckID: "d1", Screen: "deck", CreatedAt: at}, "user")
	want := pendingRow{UID: "u1", Deck: "n-f1", Session: "s1", Who: "user", Kind: "note", Note: "f1", Screen: "deck",
		TargetDeck: "d1", CreatedAt: at}
	if row != want {
		t.Fatalf("row = %+v, want %+v", row, want)
	}
}

// The summary names the screen and the length, and no text of the reader.
func TestDescribeNotePrintsNoText(t *testing.T) {
	p := feedback.Pending{ID: "f1", Screen: "collection", CreatedAt: time.Date(2026, 10, 4, 16, 0, 0, 0, time.UTC)}
	got := describeNote(1, p, "user", feedback.Item{Text: "please add Spanish"})
	for _, want := range []string{"note on the collection screen", "note f1", "18 characters"} {
		if !strings.Contains(got, want) {
			t.Fatalf("summary %q lacks %q", got, want)
		}
	}
	if strings.Contains(got, "Spanish") {
		t.Fatalf("summary %q prints the text of the reader", got)
	}
}
