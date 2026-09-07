package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/nkramber/decktome/go/internal/evalrun"
)

// TestFailStillWritesTheRunFile: the run file lands before the verdict
// error reaches the exit code, because the compare reads a FAIL too
// (PR-15).
func TestFailStillWritesTheRunFile(t *testing.T) {
	out := filepath.Join(t.TempDir(), "run.jsonl")
	rec := evalrun.New("questions", "run-test")
	fail := errors.New("gate failed")
	if err := writeRunThen(out, rec, fail); !errors.Is(err, fail) {
		t.Fatalf("err = %v, want the verdict error", err)
	}
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("the run file is missing: %v", err)
	}
	if err := writeRunThen("", rec, nil); err != nil {
		t.Fatalf("no run file, no error: %v", err)
	}
}
