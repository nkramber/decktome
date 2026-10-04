package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nkramber/decktome/go/internal/decks"
	"github.com/nkramber/decktome/go/internal/feedback"
	"github.com/nkramber/decktome/go/internal/sessions"
)

// A thumbs down starts a live eval, as a new deck does (D-1149). The
// loop keys each item by the field `deck` of its row. A verdict takes the
// key verdictPrefix + its id. A Firestore id holds no "-", so no deck id
// starts with the prefix.
const verdictPrefix = "v-"

// thumbsDown is the kind of a verdict row.
const thumbsDown = "thumbs-down"

func verdictKey(id string) string { return verdictPrefix + id }

// verdictOf answers the verdict id of a key, and false for a deck id.
func verdictOf(key string) (string, bool) {
	if id, ok := strings.CutPrefix(key, verdictPrefix); ok && id != "" {
		return id, true
	}
	return "", false
}

func verdictRow(p feedback.Pending, who string) pendingRow {
	return pendingRow{UID: p.UID, Deck: verdictKey(p.ID), Session: p.SessionID, Who: who, Kind: thumbsDown,
		Verdict: p.ID, Target: p.Kind, TargetDeck: p.DeckID, CreatedAt: p.CreatedAt}
}

// verdictView is the file verdict.json of a bundle: what the reader
// gave a thumbs down to, and why.
type verdictView struct {
	ID           string    `json:"id"`
	Target       string    `json:"target"`
	Reasons      []string  `json:"reasons"`
	Text         string    `json:"text"`
	QuestionText string    `json:"question_text,omitempty"`
	AnswerText   string    `json:"answer_text,omitempty"`
	DeckID       string    `json:"deck_id,omitempty"`
	OracleID     string    `json:"oracle_id,omitempty"`
	SessionID    string    `json:"session_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// verdictBundle writes the bundle of one thumbs down, or of one general
// note when kind is note (D-1156). A live deck gives
// the whole deck bundle. A verdict with no deck, or on a deck that the
// reader deleted, gives the session part alone. The snapshot of the
// verdict holds the object as the reader saw it (D-635), so it goes in
// too, and it stands in for a session that the reader deleted.
func (st store) verdictBundle(ctx context.Context, uid, id, dir, nonce, kind string) error {
	v, err := st.feedback.Get(ctx, uid, id)
	if err != nil {
		return fmt.Errorf("verdict: %w", err)
	}
	live := false
	if v.DeckID != "" {
		switch err := st.bundle(ctx, uid, v.DeckID, dir, nonce); {
		case err == nil:
			live = true
		case !errors.Is(err, decks.ErrNotFound):
			return err
		}
	}
	if !live {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return err
		}
		if err := writeReadme(dir, nonce); err != nil {
			return err
		}
		meta := map[string]any{"user": short(uid), "kind": kind, "session": v.SessionID}
		if v.SessionID != "" {
			if err := st.writeSession(ctx, uid, v.SessionID, "", dir, nonce, meta); err != nil {
				if !errors.Is(err, sessions.ErrNotFound) {
					return err
				}
				meta["session_error"] = "the reader deleted the session"
			}
		}
		if err := writeJSON(filepath.Join(dir, "meta.json"), meta); err != nil {
			return err
		}
	}
	if v.Session != nil {
		if err := writeProto(filepath.Join(dir, "verdict-session.json"), v.Session); err != nil {
			return err
		}
		// The replay reads session.json. A deleted session leaves the
		// snapshot alone.
		if _, err := os.Stat(filepath.Join(dir, "session.json")); errors.Is(err, os.ErrNotExist) {
			if err := writeProto(filepath.Join(dir, "session.json"), v.Session); err != nil {
				return err
			}
		}
	}
	if v.Deck != nil {
		if err := writeProto(filepath.Join(dir, "verdict-deck.json"), v.Deck); err != nil {
			return err
		}
	}
	if v.Import != nil {
		if err := writeProto(filepath.Join(dir, "verdict-import.json"), v.Import); err != nil {
			return err
		}
	}
	if kind == note {
		return writeJSON(filepath.Join(dir, "note.json"), noteView{ID: v.ID, Screen: v.Screen, Text: v.Text,
			DeckID: v.DeckID, SessionID: v.SessionID, CreatedAt: v.CreatedAt})
	}
	return writeJSON(filepath.Join(dir, "verdict.json"), verdictView{ID: v.ID, Target: v.Kind, Reasons: v.Reasons,
		Text: v.Text, QuestionText: v.QuestionText, AnswerText: v.AnswerText, DeckID: v.DeckID, OracleID: v.OracleID,
		SessionID: v.SessionID, CreatedAt: v.CreatedAt})
}

// describeVerdict is the summary of one thumbs down. It prints the
// reasons and no text of the reader.
func describeVerdict(n int, p feedback.Pending, who string, v feedback.Item) string {
	var b strings.Builder
	fmt.Fprintf(&b, "\n[%d] thumbs down on a %s, %s account, %s\n", n, orNone(p.Kind), who, p.CreatedAt.UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "    verdict %s, deck %s, session %s\n", p.ID, orNone(p.DeckID), orNone(p.SessionID))
	fmt.Fprintf(&b, "    reasons: %s\n", orNone(strings.Join(v.Reasons, ", ")))
	if v.Text != "" {
		fmt.Fprintf(&b, "    the reader wrote a note of %d characters\n", len([]rune(v.Text)))
	}
	return b.String()
}
