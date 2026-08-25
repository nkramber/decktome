package main

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

func load(t *testing.T) gateFile {
	t.Helper()
	var f gateFile
	if err := json.Unmarshal(conversationsJSON, &f); err != nil {
		t.Fatalf("conversations.json: %v", err)
	}
	return f
}

func TestConversationsFile(t *testing.T) {
	f := load(t)
	if f.VerifiedAt == "" {
		t.Error("conversations.json has no verified_at")
	}
	if len(f.Conversations) < questions.MinGateSize {
		t.Errorf("%d conversations, the gate needs %d", len(f.Conversations), questions.MinGateSize)
	}
	seenID := map[int]bool{}
	seenName := map[string]bool{}
	for i, c := range f.Conversations {
		switch {
		case c.ID != i+1:
			t.Errorf("conversation %d has id %d, want %d", i+1, c.ID, i+1)
		case seenID[c.ID]:
			t.Errorf("duplicate id %d", c.ID)
		case c.Name == "":
			t.Errorf("conversation %d has no name", c.ID)
		case seenName[c.Name]:
			t.Errorf("duplicate name %q", c.Name)
		case len(c.Messages) == 0:
			t.Errorf("conversation %q has no message", c.Name)
		case len(c.Messages) > 4:
			t.Errorf("conversation %q has %d messages, the gate allows four turns", c.Name, len(c.Messages))
		}
		seenID[c.ID], seenName[c.Name] = true, true
		for j, m := range c.Messages {
			if strings.TrimSpace(m) == "" {
				t.Errorf("conversation %q message %d is empty", c.Name, j+1)
			}
		}
	}
}

// TestRunNeedsApproval keeps the command from spending money by accident.
func TestRunNeedsApproval(t *testing.T) {
	t.Setenv("QUESTIONS_GATE", "")
	err := run("", 0, io.Discard)
	if err == nil || !strings.Contains(err.Error(), "QUESTIONS_GATE=1") {
		t.Errorf("err = %v, want a refusal without QUESTIONS_GATE=1", err)
	}
}
