package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nkramber/decktome/go/internal/triage"
)

// TestApplyWritesNoCaseThatNoGateRuns holds D-841. The question gate
// refuses a conversation with no message, so the apply step must not
// append one to the gate file.
func TestApplyWritesNoCaseThatNoGateRuns(t *testing.T) {
	root := t.TempDir()
	results := []triage.Result{{Case: triage.Case{
		Class: "Q1", Artifact: triage.AConversation, Target: triage.TargetConversations,
		ID: 112, Body: []byte(`{"id":112,"name":"Q1: asked again","messages":null}`),
		NoRun: "the question gate refuses a conversation with no message",
	}}}
	if err := applyCases(root, results); err != nil {
		t.Fatal(err)
	}
	if results[0].Applied != "" {
		t.Errorf("the case reads applied to %q", results[0].Applied)
	}
	if _, err := os.Stat(filepath.Join(root, triage.TargetConversations)); !os.IsNotExist(err) {
		t.Errorf("the gate file exists or can not be read: %v", err)
	}
}
