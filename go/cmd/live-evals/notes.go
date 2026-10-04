package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/nkramber/decktome/go/internal/feedback"
)

// A general note of the "Leave feedback" button starts a live eval, as a
// thumbs down does (D-1156). A note takes the key notePrefix + its id. A
// Firestore id holds no "-", so no deck id starts with the prefix.
const notePrefix = "n-"

// note is the kind of a note row.
const note = "note"

func noteKey(id string) string { return notePrefix + id }

// noteOf answers the note id of a key, and false for any other key.
func noteOf(key string) (string, bool) {
	if id, ok := strings.CutPrefix(key, notePrefix); ok && id != "" {
		return id, true
	}
	return "", false
}

// feedbackOf answers the feedback id of a verdict key or a note key.
func feedbackOf(key string) (string, bool) {
	if id, ok := verdictOf(key); ok {
		return id, true
	}
	return noteOf(key)
}

func noteRow(p feedback.Pending, who string) pendingRow {
	return pendingRow{UID: p.UID, Deck: noteKey(p.ID), Session: p.SessionID, Who: who, Kind: note,
		Note: p.ID, Screen: p.Screen, TargetDeck: p.DeckID, CreatedAt: p.CreatedAt}
}

// noteView is the file note.json of a bundle: the screen and the text of
// the note.
type noteView struct {
	ID        string    `json:"id"`
	Screen    string    `json:"screen"`
	Text      string    `json:"text"`
	DeckID    string    `json:"deck_id,omitempty"`
	SessionID string    `json:"session_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// describeNote is the summary of one note. It prints the screen and the
// length of the text, and no text of the reader.
func describeNote(n int, p feedback.Pending, who string, v feedback.Item) string {
	var b strings.Builder
	fmt.Fprintf(&b, "\n[%d] note on the %s screen, %s account, %s\n", n, orNone(p.Screen), who, p.CreatedAt.UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "    note %s, deck %s, session %s\n", p.ID, orNone(p.DeckID), orNone(p.SessionID))
	fmt.Fprintf(&b, "    the reader wrote %d characters\n", len([]rune(v.Text)))
	return b.String()
}
