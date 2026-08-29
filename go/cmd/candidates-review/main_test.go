package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestRunToRefusesAnExistingFile is D-65. The refusal comes before the
// snapshot is read, so the test needs none.
func TestRunToRefusesAnExistingFile(t *testing.T) {
	t.Setenv("CARDS_SNAPSHOT_DIR", "")
	taken := filepath.Join(t.TempDir(), "review.md")
	if err := os.WriteFile(taken, []byte("scored"), 0o600); err != nil {
		t.Fatal(err)
	}
	err := runTo("", 40, taken)
	if err == nil || !strings.Contains(err.Error(), "D-65") {
		t.Errorf("err = %v, want a D-65 refusal", err)
	}
	if raw, _ := os.ReadFile(taken); string(raw) != "scored" {
		t.Error("the existing file was changed")
	}
}
