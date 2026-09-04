package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
)

// run14 mirrors the lines cmd/deck-gate writes: a clean deck and the
// empty deck of F-34.
const run14 = `# PR-8 deck gate

Run date: 2026-09-04. Card snapshot: 2026-09-03.

Verdict: FAIL. 23 of 24 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 2 |
| Prompt version | 12 |
| Calls | 62 |
| Cost | $2.5432 |
| Time | 2081 seconds |

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Commander: Karlov of the Ghost Council.

Grade: bad, score 0.12, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $120.00 to buy, $140.00 the whole deck.

**Summary:** a plain summary.

- JUDGE [true]: "a claim". why
- [WARN] ` + "`curve_summary`" + `: a curve
- [INFO] ` + "`not_owned`" + `: a card

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Grade: bad, score 0.00, model 20260903T210658Z.

Cards: 0 main, 0 sideboard. Repair turn: yes, for profile_off_band, two_card_combo. Block findings: 1.

Cost: $0.00 to buy, $0.00 the whole deck.

**Summary:** No legal repaired deck.

- NOTE: an invented name
- JUDGE [unknown]: "No legal repaired deck". why
- [BLOCK] ` + "`deck_size`" + `: deck has 1 cards
- [WARN] ` + "`land_count`" + `: 0 lands
`

const run14b = `# PR-8 deck gate

Run date: 2026-09-04. Card snapshot: 2026-09-03.

Verdict: PASS. 1 of 1 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

| Prompt version | 12 |
| Calls | 4 |
| Cost | $0.1342 |
| Time | 119 seconds |

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Grade: typical, score 0.53, model 20260903T210658Z.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $5.00 to buy, $50.00 the whole deck.

**Summary:** a deck.

- JUDGE [true]: "a claim". why
`

// TestImportReadsTheDocument: the header and one row per bar per deck.
func TestImportReadsTheDocument(t *testing.T) {
	run, err := importDeckGate(strings.NewReader(run14), "pr8-deck-gate-run14", "pr8-deck-gate-run14.md")
	if err != nil {
		t.Fatal(err)
	}
	h := run.Header
	if h.Suite != "decks" || h.Date != "2026-09-04" || h.Snapshot != "2026-09-03" || h.Verdict != "FAIL" || h.Prompts["generate"] != 12 || h.Calls != 62 || *h.CostUSD != 2.5432 || h.Seconds != 2081 {
		t.Errorf("header = %+v", h)
	}
	if h.Versions["quality_model"] != "20260903T210658Z" || !strings.HasPrefix(h.Note, "imported from pr8-deck-gate-run14.md") {
		t.Errorf("versions = %v, note %q", h.Versions, h.Note)
	}
	got := map[string]evalrun.Row{}
	for _, r := range run.Rows {
		got[r.Item+"/"+r.Metric] = r
	}
	for key, want := range map[string]float64{
		"1/built": 1, "1/blocks": 0, "1/invented_names": 0, "1/false_rules": 0, "1/warnings": 1, "1/cards": 99, "1/buy_cost": 120, "1/grade": 0.12, "1/pool": 303,
		"17/built": 1, "17/blocks": 1, "17/invented_names": 1, "17/repaired": 1, "17/cards": 0, "17/warnings": 1,
	} {
		if got[key].Value != want {
			t.Errorf("%s = %v, want %v", key, got[key].Value, want)
		}
	}
	if got["17/blocks"].Detail != "deck_size" || got["17/repaired"].Detail != "profile_off_band, two_card_combo" || got["1/grade"].Detail != "bad" {
		t.Errorf("details: blocks %q, repaired %q, grade %q", got["17/blocks"].Detail, got["17/repaired"].Detail, got["1/grade"].Detail)
	}
	if got["1/blocks"].Kind != evalrun.KindGate || got["1/cards"].Kind != evalrun.KindInfo {
		t.Error("blocks gates and cards informs")
	}
	if _, err := importDeckGate(strings.NewReader("# nothing\n"), "x", "x.md"); err == nil {
		t.Error("a document with no deck imported")
	}
}

// TestCheckNamesTheFlip is the gate line of slice 2: run 14 with 14b as
// the baseline, and run 14 alone as the newer run, names deck 17 as the
// one flip. The reverse, 14 as the base and 14b the next, reads better.
func TestCheckNamesTheFlip(t *testing.T) {
	dir := t.TempDir()
	write := func(name, doc string) {
		t.Helper()
		run, err := importDeckGate(strings.NewReader(doc), strings.TrimSuffix(name, ".jsonl"), name)
		if err != nil {
			t.Fatal(err)
		}
		if err := evalrun.WriteFile(filepath.Join(dir, name), run); err != nil {
			t.Fatal(err)
		}
	}
	write("pr8-deck-gate-run14.jsonl", run14)
	write("pr8-deck-gate-run14b.jsonl", run14b)

	var out bytes.Buffer
	if err := setBaseline(&out, dir, "decks", []string{"pr8-deck-gate-run14.jsonl"}, false); err == nil {
		t.Error("a FAIL run became a baseline without -force")
	}
	if err := setBaseline(&out, dir, "decks", []string{"pr8-deck-gate-run14.jsonl", "pr8-deck-gate-run14b.jsonl"}, false); err != nil {
		t.Fatalf("14 with 14b is a passing baseline: %v", err)
	}
	bl, err := readBaselines(dir)
	if err != nil || len(bl["decks"]) != 2 {
		t.Fatalf("baselines = %v, %v", bl, err)
	}

	// No run newer than the baseline: the check passes on the baseline
	// alone, and a run older than the baseline is history it never reads.
	write("pr8-deck-gate-run13.jsonl", strings.Replace(run14, "2026-09-04", "2026-09-03", 1))
	out.Reset()
	code, err := check(&out, dir, 0, false)
	if err != nil || code != exitPass || !strings.Contains(out.String(), "stands alone") {
		t.Errorf("check with no newer run: code %d, err %v:\n%s", code, err, out.String())
	}

	// A partial run since the baseline is listed and never compared, so
	// the suite stays NOT EVALUATED and the exit code stays green.
	part := evalrun.New("decks", "pr8-deck-gate-run14c")
	part.Header.Date, part.Header.Only, part.Header.Verdict = "2026-09-05", "17", "PASS"
	part.Gate("17", "blocks", 0, "")
	if err := evalrun.WriteFile(filepath.Join(dir, "pr8-deck-gate-run14c.jsonl"), part); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	code, err = check(&out, dir, 0, false)
	if err != nil || code != exitPass || !strings.Contains(out.String(), "stands alone") ||
		!strings.Contains(out.String(), "Partial runs since the baseline, not compared:\n\n- `pr8-deck-gate-run14c` of 2026-09-05 over `17`, 1 items, PASS.\n") {
		t.Errorf("check with a partial run: code %d, err %v:\n%s", code, err, out.String())
	}

	// A newer run that is run 14 alone reads the empty deck as the flip.
	// The information rows fold into a count until -info lists them.
	write("pr8-deck-gate-run15.jsonl", strings.Replace(run14, "2026-09-04", "2026-09-05", 1))
	out.Reset()
	code, err = check(&out, dir, 0, false)
	if err != nil || code != exitFail {
		t.Errorf("check: code %d, err %v", code, err)
	}
	doc := out.String()
	if !strings.Contains(doc, "| 17 | `blocks` | gate | 0 | 1 | WORSE, deck_size |") {
		t.Errorf("the check did not name deck 17:\n%s", doc)
	}
	if strings.Contains(doc, "\n| 1 | `") {
		t.Errorf("deck 1 did not move and must not appear:\n%s", doc)
	}
	if strings.Contains(doc, "`cards` | info") || !strings.Contains(doc, "information rows moved, 3 of them worse. Pass -info") {
		t.Errorf("the information rows must fold into a count:\n%s", doc)
	}
	out.Reset()
	if _, err := check(&out, dir, 0, true); err != nil || !strings.Contains(out.String(), "| 17 | `cards` | info | 99 | 0 | WORSE |") {
		t.Errorf("-info lists the information rows:\n%s", out.String())
	}

	// compare in the better direction, through the file mode.
	out.Reset()
	code, err = compareFiles(&out, []string{filepath.Join(dir, "pr8-deck-gate-run15.jsonl")}, filepath.Join(dir, "pr8-deck-gate-run14b.jsonl"), 0, false)
	if err != nil || code != exitPass || !strings.Contains(out.String(), "| 17 | `blocks` | gate | 1 | 0 | better, deck_size |") {
		t.Errorf("compare: code %d, err %v:\n%s", code, err, out.String())
	}
	if !strings.Contains(out.String(), "Items of the baseline the run lacks: 1.") {
		t.Errorf("the -only rerun lacks deck 1:\n%s", out.String())
	}
}

// TestNewestRunReadsTheDateThenTheNumber: run 10 is newer than run 9 on
// the same day, and a later day beats both.
func TestNewestRunReadsTheDateThenTheNumber(t *testing.T) {
	headers := []fileHeader{
		{name: "run9.jsonl", header: evalrun.Header{Suite: "decks", RunID: "pr8-deck-gate-run9", Date: "2026-09-01"}},
		{name: "run10.jsonl", header: evalrun.Header{Suite: "decks", RunID: "pr8-deck-gate-run10", Date: "2026-09-01"}},
		{name: "run8.jsonl", header: evalrun.Header{Suite: "decks", RunID: "pr8-deck-gate-run8", Date: "2026-09-02"}},
		{name: "other.jsonl", header: evalrun.Header{Suite: "bracket", RunID: "x", Date: "2026-09-09"}},
		{name: "run11.jsonl", header: evalrun.Header{Suite: "decks", RunID: "pr8-deck-gate-run11", Date: "2026-09-03", Only: "17"}},
	}
	if got := newestRun(headers, "decks", nil); got != "run8.jsonl" {
		t.Errorf("newest = %s, want the later whole run: the partial run 11 never stands for the suite", got)
	}
	if got := partialRuns(headers, "decks", []string{"run8.jsonl"}); len(got) != 1 || got[0].name != "run11.jsonl" {
		t.Errorf("partial runs since run 8 = %v, want run 11", got)
	}
	if got := partialRuns(headers, "decks", []string{"run11.jsonl"}); len(got) != 0 {
		t.Errorf("a partial run in the baseline is not since it: %v", got)
	}
	if got := newestRun(headers, "decks", []string{"run8.jsonl"}); got != "" {
		t.Errorf("newest past the baseline = %q, want none: runs 9 and 10 are older than the baseline of 2026-09-02", got)
	}
	if got := newestRun(headers, "decks", []string{"run9.jsonl"}); got != "run8.jsonl" {
		t.Errorf("newest past run 9 = %s, want the later day", got)
	}
	if got := newestRun(headers, "revise", nil); got != "" {
		t.Errorf("a suite with no run = %q", got)
	}
}

// TestRunRefusesAnUnknownMode pins the entry point.
func TestRunRefusesAnUnknownMode(t *testing.T) {
	var out bytes.Buffer
	if code, err := run([]string{"frobnicate"}, &out); err == nil || code != exitFault {
		t.Errorf("code %d, err %v", code, err)
	}
	if code, err := run(nil, &out); err == nil || code != exitFault {
		t.Errorf("code %d, err %v", code, err)
	}
	dir := t.TempDir()
	if code, err := run([]string{"check", "-dir", dir}, &out); err != nil || code != exitPass || !strings.Contains(out.String(), "No baseline") {
		t.Errorf("an empty directory checks clean: code %d, err %v:\n%s", code, err, out.String())
	}
	if _, err := os.Stat(filepath.Join(dir, baselinesFile)); err == nil {
		t.Error("a check wrote a baselines file")
	}
}

// TestSweepPlansUnderTheCap: the plan numbers each document after the
// highest run of its family, reads the estimate from the last run file,
// stops before the step that crosses the cap, and stops on a FAIL.
func TestSweepPlansUnderTheCap(t *testing.T) {
	root := t.TempDir()
	docs := filepath.Join(root, "docs", "reference")
	evalDir := filepath.Join(docs, "eval")
	if err := os.MkdirAll(evalDir, 0o750); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"pr7-question-gate-run33.md", "pr8-deck-gate-run14.md", "pr8-deck-gate-run14b.md", "pr12b-revise-gate-run8.md"} {
		if err := os.WriteFile(filepath.Join(docs, name), []byte("# doc\n"), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	// The decks suite has a run file with a cost, so its estimate reads
	// it. Its one-prompt rerun is newer and cheaper, and the estimate
	// reads the full run.
	decks := evalrun.New("decks", "pr8-deck-gate-run14")
	cost := 2.54
	decks.Header.CostUSD = &cost
	decks.Header.Verdict = "FAIL"
	decks.Gate("1", "blocks", 0, "")
	decks.Gate("2", "blocks", 0, "")
	if err := evalrun.WriteFile(filepath.Join(evalDir, "pr8-deck-gate-run14.jsonl"), decks); err != nil {
		t.Fatal(err)
	}
	rerun := evalrun.New("decks", "pr8-deck-gate-run14b")
	small := 0.13
	rerun.Header.CostUSD = &small
	rerun.Header.Verdict = "PASS"
	rerun.Gate("2", "blocks", 0, "")
	if err := evalrun.WriteFile(filepath.Join(evalDir, "pr8-deck-gate-run14b.jsonl"), rerun); err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	code, err := sweep(&out, root, nil, 3, true, false, nil)
	if err != nil || code != exitPass {
		t.Fatalf("dry sweep: code %d, err %v:\n%s", code, err, out.String())
	}
	plan := out.String()
	for _, want := range []string{
		"| questions | `make questions-gate` | `docs/reference/pr7-question-gate-run34.md` | $0.18 |",
		"| question-eval | `make questions-eval` | `docs/reference/pr7-question-eval-run1.md` | $0.10 |",
		"| decks | `make deck-gate` | `docs/reference/pr8-deck-gate-run15.md` | $2.54 |",
		"| revise | `make revise-gate` | `docs/reference/pr12b-revise-gate-run9.md` | $1.23 |",
		"over the cap of $3.00",
		"Dry run: nothing ran",
	} {
		if !strings.Contains(plan, want) {
			t.Errorf("the plan lacks %q:\n%s", want, plan)
		}
	}

	// A run: the fake runner writes a run file per step. The eval step
	// reads the question gate document the sweep names, the deck step
	// fails and stops the sweep, and the spend counts what ran.
	t.Setenv("EVAL_SWEEP", "1")
	var ran []string
	runner := func(gotRoot, target string, vars []string) error {
		ran = append(ran, target+" "+strings.Join(vars, " "))
		var docPath, runPath string
		for _, v := range vars {
			k, val, _ := strings.Cut(v, "=")
			switch k {
			case "GATE_OUT", "EVAL_OUT", "DECK_GATE_OUT", "REVISE_GATE_OUT", "QUALITY_JUDGE_OUT":
				docPath = val
			case "GATE_RUN", "EVAL_ROWS", "DECK_GATE_RUN", "REVISE_GATE_RUN", "QUALITY_JUDGE_RUN":
				runPath = val
			}
		}
		suite := map[string]string{"questions-gate": "questions", "questions-eval": "question-eval", "deck-gate": "decks", "revise-gate": "revise"}[target]
		rec := evalrun.New(suite, evalrun.RunID(runPath))
		c := map[string]float64{"questions": 0.2, "question-eval": 0.1, "decks": 2.5}[suite]
		rec.Header.CostUSD = &c
		rec.Header.Verdict = "PASS"
		if suite == "decks" {
			rec.Header.Verdict = "FAIL"
		}
		rec.Gate("1", "x", 0, "")
		if err := os.WriteFile(filepath.Join(gotRoot, docPath), []byte("# doc\n"), 0o600); err != nil {
			return err
		}
		return evalrun.WriteFile(filepath.Join(gotRoot, runPath), rec)
	}
	out.Reset()
	code, err = sweep(&out, root, []string{"questions", "question-eval", "decks", "revise"}, 5, false, false, runner)
	if err != nil || code != exitFail {
		t.Fatalf("sweep: code %d, err %v:\n%s", code, err, out.String())
	}
	if len(ran) != 3 || !strings.HasPrefix(ran[2], "deck-gate ") {
		t.Errorf("ran %v, want three steps that stop at the deck gate FAIL", ran)
	}
	if !strings.Contains(ran[1], "EVAL_RUN=docs/reference/pr7-question-gate-run34.md") || !strings.Contains(ran[1], "EVAL_JSON=.local/tune/run1.json") {
		t.Errorf("the eval step reads the gate document of this sweep: %s", ran[1])
	}
	if !strings.Contains(out.String(), "Spent $2.80 of the cap of $5.00.") || !strings.Contains(out.String(), "The sweep stops here") {
		t.Errorf("the spend and the stop:\n%s", out.String())
	}

	// Without the guard nothing runs.
	t.Setenv("EVAL_SWEEP", "")
	out.Reset()
	if code, err := sweep(&out, root, nil, 5, false, false, runner); err == nil || code != exitFault {
		t.Errorf("the sweep ran without EVAL_SWEEP=1: code %d, err %v", code, err)
	}
	if _, err := sweep(&out, root, nil, 0, true, false, nil); err == nil {
		t.Error("a sweep with no cap planned")
	}
}
