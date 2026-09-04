package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestRunRefusesAnExistingOutput is D-65. The refusal comes before any
// read of the gate document and before any provider call, so the test
// needs neither.
func TestRunRefusesAnExistingOutput(t *testing.T) {
	t.Setenv("QUESTIONS_EVAL", "1")
	dir := t.TempDir()
	taken := filepath.Join(dir, "taken")
	if err := os.WriteFile(taken, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	free := filepath.Join(dir, "free")
	cases := []struct{ name, out, json string }{
		{"the eval document exists", taken, free},
		{"the summary exists", free, taken},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := run(filepath.Join(dir, "no-gate.md"), tc.out, tc.json, "", 0.5, 0, 3)
			if err == nil || !strings.Contains(err.Error(), "D-65") {
				t.Errorf("err = %v, want a D-65 refusal", err)
			}
		})
	}
}
