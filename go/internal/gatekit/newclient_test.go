package gatekit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nkramber/decktome/go/internal/livespend"
)

func TestNewClientWritesTheSpendOfALiveEval(t *testing.T) {
	dir := t.TempDir()
	t.Setenv(livespend.EnvBundle, dir)
	t.Setenv(livespend.EnvBudget, "3.00")
	t.Setenv("OPENAI_API_KEY", "sk-test-not-called")
	t.Setenv("ANTHROPIC_API_KEY", "sk-test-not-called")
	client, done, err := NewClient("chat-probe", Quiet())
	if err != nil {
		t.Fatal(err)
	}
	if client.Meter() == nil {
		t.Fatal("the client has no meter")
	}
	done()
	raw, err := os.ReadFile(filepath.Join(dir, livespend.SpendFile))
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
	if len(lines) != 2 || !strings.Contains(lines[0], `"event":"start"`) || !strings.Contains(lines[1], `"measured":true`) {
		t.Fatalf("spend file = %q, want a start line and a measured end line", raw)
	}
	// The run reads as measured, so the next paid run can start.
	if _, done2, err := NewClient("revise-gate", Quiet()); err != nil {
		t.Fatalf("the second run: %v", err)
	} else {
		done2()
	}
}

func TestNewClientRefusesAfterAnUnmeasuredRun(t *testing.T) {
	dir := t.TempDir()
	t.Setenv(livespend.EnvBundle, dir)
	t.Setenv(livespend.EnvBudget, "3.00")
	t.Setenv("OPENAI_API_KEY", "sk-test-not-called")
	t.Setenv("ANTHROPIC_API_KEY", "sk-test-not-called")
	// A start line with no end line: the run was killed.
	line := `{"id":"0123456789abcdef","event":"start","target":"chat-probe"}` + "\n"
	if err := os.WriteFile(filepath.Join(dir, livespend.SpendFile), []byte(line), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, _, err := NewClient("chat-probe", Quiet()); err == nil || !strings.Contains(err.Error(), "budget") {
		t.Fatalf("err = %v, want a refusal for the budget", err)
	}
}
