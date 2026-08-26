package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/tune"
)

func write(t *testing.T, dir, name string, s tune.Summary) string {
	t.Helper()
	path := filepath.Join(dir, name)
	raw, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, raw, 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

// TestExitCodes covers the three answers the loop driver reads: keep the
// iteration, revert it, or stop because the run is good enough.
func TestExitCodes(t *testing.T) {
	dir := t.TempDir()
	good := tune.Summary{Judged: 100, Bad: 4, Ratio: 0.04,
		Metrics: tune.Metrics{Questions: 200, CatalogFilled: 150}}
	better := tune.Summary{Judged: 100, Bad: 10, Ratio: 0.10,
		Metrics: tune.Metrics{Questions: 200, CatalogFilled: 150}}
	worse := tune.Summary{Judged: 100, Bad: 30, Ratio: 0.30,
		Metrics: tune.Metrics{Questions: 200, CatalogFilled: 150}}
	prev := tune.Summary{Judged: 100, Bad: 20, Ratio: 0.20,
		Metrics: tune.Metrics{Questions: 200, CatalogFilled: 150}}

	prevPath := write(t, dir, "prev.json", prev)
	cases := []struct {
		name string
		next tune.Summary
		prev string
		want int
	}{
		{"an improvement is kept", better, prevPath, exitAccept},
		{"a regression is reverted", worse, prevPath, exitReject},
		{"the target ends the loop", good, prevPath, exitDone},
		{"the first run has no previous", better, "", exitAccept},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path := write(t, dir, "next.json", tc.next)
			code, err := run(path, tc.prev, 0.05, os.Stdout)
			if err != nil {
				t.Fatalf("run: %v", err)
			}
			if code != tc.want {
				t.Errorf("exit = %d, want %d", code, tc.want)
			}
		})
	}
}

// TestOmittedPreviousIsAFirstRun covers the real first iteration. The
// driver names no previous run, so there is nothing to compare.
func TestOmittedPreviousIsAFirstRun(t *testing.T) {
	dir := t.TempDir()
	path := write(t, dir, "next.json", tune.Summary{Judged: 10, Bad: 5, Ratio: 0.50,
		Metrics: tune.Metrics{Questions: 20, CatalogFilled: 15}})
	code, err := run(path, "", 0.05, os.Stdout)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if code != exitAccept {
		t.Errorf("exit = %d, want accept", code)
	}
}

// TestNamedPreviousMustExist replaces a test that read a named and
// missing file as a first run. That reading hid a real defect: the loop
// passed a relative path into a subshell that had changed directory, the
// file was not there, and the comparison was skipped in silence. An
// iteration that raised the holdout ratio was accepted and committed
// (D-171).
//
// A path the owner gave and a path the owner omitted mean different
// things. Only the second is a first run.
func TestNamedPreviousMustExist(t *testing.T) {
	dir := t.TempDir()
	path := write(t, dir, "next.json", tune.Summary{Judged: 10, Bad: 5, Ratio: 0.50,
		Metrics: tune.Metrics{Questions: 20, CatalogFilled: 15}})
	if _, err := run(path, filepath.Join(dir, "nope.json"), 0.05, os.Stdout); err == nil {
		t.Error("a named baseline that is missing was read as a first run")
	}
}
