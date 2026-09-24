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

// A parser fixture is a file of its own in the folder of its page
// (D-888), and the result names that file.
func TestApplyWritesAParserFixture(t *testing.T) {
	root := t.TempDir()
	results := []triage.Result{{Case: triage.Case{
		Class: "I1", Artifact: triage.AParseFixture, Target: triage.TargetDeckReports,
		ID: 3, Body: []byte(`"hello\nworld\n"`),
	}}}
	if err := applyCases(root, results); err != nil {
		t.Fatal(err)
	}
	want := triage.TargetDeckReports + "/3.txt"
	if results[0].Applied != want {
		t.Errorf("applied = %q, want %q", results[0].Applied, want)
	}
	got, err := os.ReadFile(filepath.Join(root, want))
	if err != nil || string(got) != "hello\nworld\n" {
		t.Errorf("fixture = %q, %v", got, err)
	}
}
