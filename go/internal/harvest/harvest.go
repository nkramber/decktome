// Package harvest turns the feedback of the deployed app into a dated
// document and a JSONL file (PR-28a). The triage of PR-28b reads the
// JSONL file, and a person reads the document.
//
// The watermark comes from the files themselves. The newest time any
// harvest wrote is the floor of the next one, so the documents are the
// only record and a rerun never repeats an item (D-65). A deleted file
// moves the floor back, and the next harvest writes the items again.
//
// Every file here holds what a reader wrote, and it commits with the
// repository (D-642). No file holds an email, and the guard tests refuse
// the join that would write one (D-559, D-638).
package harvest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/nkramber/decktome/go/internal/feedback"
)

// Dir is where the harvest writes, under the repo root.
const Dir = "docs/reference/feedback"

// Record is one line of the JSONL file. It names the verdict, the words
// of the reader, and the object the verdict names as it stood at that
// moment (D-635).
//
// It holds the uid and never an email (D-559). No proto of this repo
// carries an email but the invite service, so the snapshot can hold
// none.
type Record struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UID       string    `json:"uid"`
	Kind      string    `json:"kind"`
	Verdict   string    `json:"verdict"`
	Reasons   []string  `json:"reasons,omitempty"`
	Text      string    `json:"text,omitempty"`
	SessionID string    `json:"session_id,omitempty"`
	// QuestionID is the id the session gave the question, and it reads
	// "q<n>-<row>", so the triage of PR-28b names the catalog row a
	// verdict is about with no model call (agent.go).
	QuestionID string           `json:"question_id,omitempty"`
	Question   string           `json:"question_text,omitempty"`
	Answer     string           `json:"answer_text,omitempty"`
	DeckID     string           `json:"deck_id,omitempty"`
	OracleID   string           `json:"oracle_id,omitempty"`
	Prompts    map[string]int64 `json:"prompts,omitempty"`
	// Context reads "snapshot" when the verdict carries the object it
	// names, and "absent" when it does not. A verdict written before
	// D-635 reads absent, and so does one whose kind names no object.
	Context string `json:"context"`
	// Session and Deck hold the snapshot as protojson. The triage of
	// PR-28b reads the deck list and the turns from here.
	Session json.RawMessage `json:"session,omitempty"`
	Deck    json.RawMessage `json:"deck,omitempty"`
}

// RecordOf reads one stored verdict as a record.
func RecordOf(item feedback.Item) (Record, error) {
	r := Record{
		ID:         item.ID,
		CreatedAt:  item.CreatedAt.UTC(),
		UID:        item.UID,
		Kind:       item.Kind,
		Verdict:    item.Verdict,
		Reasons:    item.Reasons,
		Text:       item.Text,
		SessionID:  item.SessionID,
		QuestionID: item.QuestionID,
		Question:   item.QuestionText,
		Answer:     item.AnswerText,
		DeckID:     item.DeckID,
		OracleID:   item.OracleID,
		Prompts:    item.Prompts,
		Context:    "absent",
	}
	m := protojson.MarshalOptions{}
	if item.Session != nil {
		raw, err := m.Marshal(item.Session)
		if err != nil {
			return Record{}, fmt.Errorf("harvest %s: session: %w", item.ID, err)
		}
		r.Session, r.Context = raw, "snapshot"
	}
	if item.Deck != nil {
		raw, err := m.Marshal(item.Deck)
		if err != nil {
			return Record{}, fmt.Errorf("harvest %s: deck: %w", item.ID, err)
		}
		r.Deck, r.Context = raw, "snapshot"
	}
	return r, nil
}

// Watermark reads the newest CreatedAt any harvest under root wrote. It
// answers the zero time when no harvest exists, and the caller then
// reads every verdict the store holds.
func Watermark(root string) (time.Time, error) {
	var newest time.Time
	paths, err := filepath.Glob(filepath.Join(root, Dir, "*.jsonl"))
	if err != nil {
		return time.Time{}, err
	}
	for _, p := range paths {
		raw, err := os.ReadFile(p)
		if err != nil {
			return time.Time{}, fmt.Errorf("harvest: %s: %w", p, err)
		}
		for _, line := range strings.Split(string(raw), "\n") {
			if strings.TrimSpace(line) == "" {
				continue
			}
			var rec Record
			// A line this build can not read must never move the floor
			// back, because the next harvest would write its item again.
			if err := json.Unmarshal([]byte(line), &rec); err != nil {
				return time.Time{}, fmt.Errorf("harvest: %s: %w", p, err)
			}
			if rec.CreatedAt.After(newest) {
				newest = rec.CreatedAt.UTC()
			}
		}
	}
	return newest, nil
}
