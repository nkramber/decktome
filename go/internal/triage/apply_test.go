package triage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/feedback"
)

// TestTheClassKeysMatchTheStore pins the class table against the reason
// keys the dialog sends. A key the store gained with no class would fall
// through the triage without a word, which is how the chat kind of D-594
// went unclassed for three days.
func TestTheClassKeysMatchTheStore(t *testing.T) {
	kinds := map[mtgv1.FeedbackKind]string{
		mtgv1.FeedbackKind_FEEDBACK_KIND_QUESTION: "question",
		mtgv1.FeedbackKind_FEEDBACK_KIND_SUMMARY:  "summary",
		mtgv1.FeedbackKind_FEEDBACK_KIND_CARD:     "card",
		mtgv1.FeedbackKind_FEEDBACK_KIND_DECK:     "deck",
		mtgv1.FeedbackKind_FEEDBACK_KIND_CHAT:     "chat",
	}
	seen := map[string]bool{}
	for enum, word := range kinds {
		for _, key := range feedback.Reasons(enum) {
			seen[word+"/"+key] = true
			if _, ok := ClassOf(word, key); !ok {
				t.Errorf("the store offers %s/%s and no class names it", word, key)
			}
		}
	}
	for _, c := range classes {
		if !seen[c.Kind+"/"+c.Reason] {
			t.Errorf("class %s names %s/%s, and the store offers no such reason", c.ID, c.Kind, c.Reason)
		}
	}
}

// TestAppendKeepsTheRestOfTheFile is the whole reason the writer is a
// text insert. conversations.json holds 108 conversations, and a
// re-encode would put every one of them into the diff (D-642).
func TestAppendKeepsTheRestOfTheFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "prompts.json")
	before := `{
  "note": "the golden prompts",
  "prompts": [
    {"id": 1, "name": "one"},
    {"id": 7, "name": "seven"}
  ]
}
`
	if err := os.WriteFile(path, []byte(before), 0o600); err != nil {
		t.Fatal(err)
	}
	next, err := NextID(path)
	if err != nil {
		t.Fatal(err)
	}
	if next != 8 {
		t.Errorf("next id = %d, want 8, one past the highest", next)
	}
	body, err := json.Marshal(map[string]any{"id": next, "name": "eight"})
	if err != nil {
		t.Fatal(err)
	}
	if err := Append(path, body); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	got := string(raw)
	// Every line that was there is still there, character for character.
	for _, line := range []string{`"note": "the golden prompts"`, `{"id": 1, "name": "one"}`, `{"id": 7, "name": "seven"}`} {
		if !strings.Contains(got, line) {
			t.Errorf("the append rewrote the file: %q is gone\n%s", line, got)
		}
	}
	// The file still parses, and it holds the new case.
	var file struct {
		Prompts []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"prompts"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatalf("the file no longer parses: %v\n%s", err, got)
	}
	if n := len(file.Prompts); n != 3 {
		t.Fatalf("%d prompts, want 3", n)
	}
	if file.Prompts[2].ID != 8 || file.Prompts[2].Name != "eight" {
		t.Errorf("the last prompt is %+v, want the new case", file.Prompts[2])
	}
}

// TestAppendRefusesAFileItCanNotRead keeps a half-written gate file off
// the disk. A file that does not end with an array and a document is one
// the writer does not understand.
func TestAppendRefusesAFileItCanNotRead(t *testing.T) {
	path := filepath.Join(t.TempDir(), "broken.json")
	if err := os.WriteFile(path, []byte(`{"prompts": []`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := Append(path, json.RawMessage(`{"id":1}`)); err == nil {
		t.Error("a file the writer can not read was written anyway")
	}
	raw, _ := os.ReadFile(path)
	if string(raw) != `{"prompts": []` {
		t.Errorf("the refused file changed: %s", raw)
	}
}

// TestTheOwnerRowJoinsTheDecisionQueue writes the artifact of a class
// that meets an owner decision (D-558).
func TestTheOwnerRowJoinsTheDecisionQueue(t *testing.T) {
	path := filepath.Join(t.TempDir(), "owner-questions.md")
	before := `# Questions only the owner can answer

## Waits on a decision

| # | Question | Why only you | What it blocks |
|---|---|---|---|
| OQ-79 | an older one | because | a thing |

## How to answer
`
	if err := os.WriteFile(path, []byte(before), 0o600); err != nil {
		t.Fatal(err)
	}
	next, err := NextOQ(path)
	if err != nil {
		t.Fatal(err)
	}
	if next != 80 {
		t.Errorf("next question = %d, want 80", next)
	}
	// The reader's words may hold a pipe and a newline, and neither one
	// may break the table.
	err = AppendOwnerQuestion(path, "OQ-80", "A reader disagrees\nwith D-37 | really", "It changes a rule", "One verdict")
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	got := string(raw)
	if !strings.Contains(got, "| OQ-80 | A reader disagrees with D-37 / really | It changes a rule | One verdict |") {
		t.Errorf("the row is missing or broken:\n%s", got)
	}
	if !strings.Contains(got, "| OQ-79 | an older one") {
		t.Error("the older question is gone")
	}
	// The new row sits at the head of the table, so the owner reads it
	// first.
	if strings.Index(got, "OQ-80") > strings.Index(got, "OQ-79") {
		t.Error("the new row sits under the old ones")
	}
}

// TestAppendJoinsAnEmptyArray covers the file somebody makes later. An
// empty array takes no leading comma, and a comma would leave the file
// unparseable.
func TestAppendJoinsAnEmptyArray(t *testing.T) {
	for name, before := range map[string]string{
		"an array on two lines": "{\n  \"prompts\": [\n  ]\n}\n",
		"an array on one line":  "{\n  \"prompts\": []\n}\n",
	} {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "prompts.json")
			if err := os.WriteFile(path, []byte(before), 0o600); err != nil {
				t.Fatal(err)
			}
			next, err := NextID(path)
			if err != nil {
				t.Fatal(err)
			}
			if next != 1 {
				t.Errorf("next id = %d, want 1 on an empty file", next)
			}
			body, err := json.Marshal(map[string]any{"id": next, "name": "the first case"})
			if err != nil {
				t.Fatal(err)
			}
			if err := Append(path, body); err != nil {
				t.Fatal(err)
			}
			raw, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var file struct {
				Prompts []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"prompts"`
			}
			if err := json.Unmarshal(raw, &file); err != nil {
				t.Fatalf("the file no longer parses: %v\n%s", err, raw)
			}
			if len(file.Prompts) != 1 || file.Prompts[0].ID != 1 || file.Prompts[0].Name != "the first case" {
				t.Errorf("prompts = %+v, want the one case", file.Prompts)
			}
		})
	}
}

// TestAppendNeverWritesAFileThatStoppedParsing is the last guard of the
// text insert. The writer edits text and never re-encodes, so it can not
// lean on the encoder to keep the file valid.
func TestAppendNeverWritesAFileThatStoppedParsing(t *testing.T) {
	path := filepath.Join(t.TempDir(), "prompts.json")
	// The tail reads like a gate file and the head does not close, so the
	// insert lands and the result can not parse.
	before := "{\n  \"prompts\": [\n    {\"id\": 1\n  ]\n}\n"
	if err := os.WriteFile(path, []byte(before), 0o600); err != nil {
		t.Fatal(err)
	}
	err := Append(path, json.RawMessage(`{"id":2}`))
	if err == nil {
		t.Fatal("a case that breaks the file was written anyway")
	}
	if !strings.Contains(err.Error(), "unparseable") {
		t.Errorf("error = %v, want the parse refusal", err)
	}
	raw, _ := os.ReadFile(path)
	if string(raw) != before {
		t.Errorf("the refused file changed:\n%s", raw)
	}
}
