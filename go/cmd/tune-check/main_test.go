package main

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
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
			code, err := run(path, tc.prev, "", 0.05, tune.DefaultNoise, "", "", os.Stdout)
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
	code, err := run(path, "", "", 0.05, tune.DefaultNoise, "", "", os.Stdout)
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if code != exitAccept {
		t.Errorf("exit = %d, want accept", code)
	}
}

// TestNamedPreviousMustExist is D-171. A named baseline that is missing
// is an error and not a first run: a skipped comparison in silence
// accepts a regression. A path that was given and a path that was
// omitted mean different things, and only the second is a first run.
func TestNamedPreviousMustExist(t *testing.T) {
	dir := t.TempDir()
	path := write(t, dir, "next.json", tune.Summary{Judged: 10, Bad: 5, Ratio: 0.50,
		Metrics: tune.Metrics{Questions: 20, CatalogFilled: 15}})
	if _, err := run(path, filepath.Join(dir, "nope.json"), "", 0.05, tune.DefaultNoise, "", "", os.Stdout); err == nil {
		t.Error("a named baseline that is missing was read as a first run")
	}
}

func verdict(conv string, row, text, warranted string, holdout bool) tune.Verdict {
	return tune.Verdict{Conversation: conv, Turn: 1, Row: row, Text: text, Warranted: warranted, Holdout: holdout}
}

func summarize(vs ...tune.Verdict) tune.Summary {
	return tune.Summarize("r", "m", 0.01, tune.Metrics{Questions: 10, CatalogFilled: 8}, vs)
}

// TestPartialSummaryIsAFault is T-2. A budget-cut eval is refused as
// -prev and as -next, and the loop reads the error as a stop.
func TestPartialSummaryIsAFault(t *testing.T) {
	dir := t.TempDir()
	whole := write(t, dir, "whole.json", summarize(verdict("1. a", "meta", "q", "yes", false)))
	part := write(t, dir, "part.json", tune.Summary{Partial: true, StoppedReason: "the budget of $0.50 stopped the run"})
	cases := []struct{ name, next, prev string }{
		{"partial next", part, whole},
		{"partial prev", whole, part},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := run(tc.next, tc.prev, "", 0.05, tune.DefaultNoise, "", "", io.Discard)
			if err == nil || !strings.Contains(err.Error(), "partial") {
				t.Errorf("err = %v, want a refusal of the partial summary", err)
			}
		})
	}
	if err := mergeRuns(whole, "", part, "", filepath.Join(dir, "m.json"), "", io.Discard); err == nil {
		t.Error("-merge folded a partial run")
	}
}

// TestPartialKeepExitsFour covers the exit the driver reads as "keep
// some, drop some, and measure the kept rows again" (D-181).
func TestPartialKeepExitsFour(t *testing.T) {
	dir := t.TempDir()
	prev := write(t, dir, "prev.json", summarize(
		verdict("1. a", "meta", "old", "no", false), verdict("2. b", "power", "old", "yes", false)))
	next := write(t, dir, "next.json", summarize(
		verdict("1. a", "meta", "new", "yes", false), verdict("2. b", "power", "new", "no", false)))
	changes := []tune.Change{
		{Commit: "aaa", Subject: "meta", Rows: []string{"meta"}},
		{Commit: "bbb", Subject: "power", Rows: []string{"power"}},
	}
	raw, _ := json.Marshal(changes)
	changesPath := filepath.Join(dir, "changes.json")
	if err := os.WriteFile(changesPath, raw, 0o600); err != nil {
		t.Fatal(err)
	}
	var out strings.Builder
	code, err := run(next, prev, changesPath, 0.05, tune.DefaultNoise, filepath.Join(dir, "lessons.md"), "it", &out)
	if err != nil {
		t.Fatal(err)
	}
	if code != exitPartial {
		t.Errorf("exit = %d, want %d: %s", code, exitPartial, out.String())
	}
	for _, line := range []string{"keep: aaa", "drop: bbb", "remeasure: 1"} {
		if !strings.Contains(out.String(), line) {
			t.Errorf("stdout lacks %q:\n%s", line, out.String())
		}
	}
	lessons, err := os.ReadFile(filepath.Join(dir, "lessons.md"))
	if err != nil || !strings.Contains(string(lessons), "partly kept") {
		t.Errorf("lessons = %q, %v", lessons, err)
	}
	// Two commits that own one row are a fault, not a verdict.
	raw, _ = json.Marshal([]tune.Change{{Commit: "aaa", Rows: []string{"meta"}}, {Commit: "bbb", Rows: []string{"meta"}}})
	if err := os.WriteFile(changesPath, raw, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := run(next, prev, changesPath, 0.05, tune.DefaultNoise, "", "", io.Discard); err == nil {
		t.Error("two owners of one row gave a verdict")
	}
}

// TestTargetReadsTheHoldout is T-14. The whole-set ratio may sit under
// the target while the holdout does not.
func TestTargetReadsTheHoldout(t *testing.T) {
	dir := t.TempDir()
	var vs []tune.Verdict
	for i := 0; i < 20; i++ {
		vs = append(vs, verdict("1. a", "meta", "q", "yes", false))
	}
	vs = append(vs, verdict("3. h", "meta", "q", "no", true), verdict("3. h", "power", "q", "yes", true))
	next := write(t, dir, "next.json", summarize(vs...))
	code, err := run(next, "", "", 0.10, tune.DefaultNoise, "", "", io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	if code != exitAccept {
		t.Errorf("exit = %d, want accept: the holdout is at 50%% while the whole set is under the target", code)
	}
}

// TestMergeWritesTheFoldedRun covers -merge end to end.
func TestMergeWritesTheFoldedRun(t *testing.T) {
	dir := t.TempDir()
	base := summarize(verdict("1. a", "meta", "old", "no", false), verdict("2. b", "colors", "c", "yes", false))
	base.Conversations = []tune.ConvCount{{Name: "1. a", Questions: 1, Filled: 1}, {Name: "2. b", Questions: 1, Filled: 1}}
	part := summarize(verdict("1. a", "meta", "new", "yes", false))
	part.Conversations = []tune.ConvCount{{Name: "1. a", Questions: 1, Filled: 1, CatalogFilled: 1}}
	basePath := write(t, dir, "base.json", base)
	partPath := write(t, dir, "part.json", part)
	out := filepath.Join(dir, "merged.json")
	doc := filepath.Join(dir, "merged.md")
	var w strings.Builder
	if err := mergeRuns(basePath, "", partPath, "", out, doc, &w); err != nil {
		t.Fatal(err)
	}
	merged, err := read(out)
	if err != nil {
		t.Fatal(err)
	}
	if merged.Judged != 2 || merged.Bad != 0 || len(merged.Conversations) != 2 {
		t.Errorf("merged = judged %d bad %d convs %d", merged.Judged, merged.Bad, len(merged.Conversations))
	}
	if _, err := os.Stat(doc); err != nil {
		t.Errorf("no merged report: %v", err)
	}
	// A result is never overwritten (D-65).
	if err := mergeRuns(basePath, "", partPath, "", out, "", io.Discard); err == nil {
		t.Error("-merge wrote over an existing summary")
	}
}

// TestAgreeComparesTwoJudges covers -agree.
func TestAgreeComparesTwoJudges(t *testing.T) {
	dir := t.TempDir()
	a := write(t, dir, "a.json", summarize(verdict("1. a", "meta", "q", "no", false), verdict("2. b", "meta", "q", "yes", false)))
	b := write(t, dir, "b.json", summarize(verdict("1. a", "meta", "q", "no", false), verdict("2. b", "meta", "q", "no", false)))
	var w strings.Builder
	if err := agreement(a, b, &w); err != nil {
		t.Fatal(err)
	}
	for _, line := range []string{"shared questions: 2", "the same verdict: 1 (50%)", "verdicts that differ: 1", "WARNING"} {
		if !strings.Contains(w.String(), line) {
			t.Errorf("output lacks %q:\n%s", line, w.String())
		}
	}
	none := write(t, dir, "none.json", summarize(verdict("9. z", "meta", "q", "no", false)))
	if err := agreement(a, none, io.Discard); err == nil {
		t.Error("two summaries with no shared question agreed")
	}
}

// TestFixerCopyDropsTheHoldout is T-8. The fixer's file keeps the
// counts and loses every holdout verdict.
func TestFixerCopyDropsTheHoldout(t *testing.T) {
	dir := t.TempDir()
	in := write(t, dir, "in.json", summarize(verdict("1. a", "meta", "q", "no", false), verdict("3. h", "meta", "q", "no", true)))
	out := filepath.Join(dir, "in.fixer.json")
	if err := fixerCopy(in, out); err != nil {
		t.Fatal(err)
	}
	got, err := read(out)
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Verdicts) != 1 || got.Verdicts[0].Holdout {
		t.Errorf("verdicts = %+v, want the tune verdict alone", got.Verdicts)
	}
	if got.HoldoutJudged != 1 || got.HoldoutBad != 1 {
		t.Errorf("the holdout counts were lost: judged %d bad %d", got.HoldoutJudged, got.HoldoutBad)
	}
}
