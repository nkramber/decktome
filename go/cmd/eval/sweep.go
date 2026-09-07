package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/gatekit"
)

// The sweep runs the paid suites in the order of the eval list, under a
// cost cap, through the Makefile targets and their guards (PR-15, slice
// 6). It runs on the owner's machine, on the owner's word: EVAL_SWEEP=1
// and -cap. A result that is not good stops the sweep, as the owner's
// rule of 2026-09-03 says, unless -continue.

// step is one suite of the sweep.
type step struct {
	suite  string
	target string
	// prefix is the document family, and the run number counts from it.
	prefix string
	// docVar and runVar name the document and the run file to make.
	docVar, runVar string
	// fallbackCost is the cost of the last measured run, for the plan
	// when no run file of the suite exists yet (the hand-off of
	// 2026-09-04).
	fallbackCost float64
	// input names the step whose document this one reads, or "".
	input string
	// inputVar names the Makefile variable that takes that document.
	inputVar string
}

var sweepSteps = []step{
	{suite: "questions", target: "questions-gate", prefix: "pr7-question-gate-run", docVar: "GATE_OUT", runVar: "GATE_RUN", fallbackCost: 0.18},
	{suite: "question-eval", target: "questions-eval", prefix: "pr7-question-eval-run", docVar: "EVAL_OUT", runVar: "EVAL_ROWS", fallbackCost: 0.10, input: "questions", inputVar: "EVAL_RUN"},
	{suite: "decks", target: "deck-gate", prefix: "pr8-deck-gate-run", docVar: "DECK_GATE_OUT", runVar: "DECK_GATE_RUN", fallbackCost: 2.60},
	{suite: "tier-judge", target: "quality-judge", prefix: "pr14b-quality-judge-run", docVar: "QUALITY_JUDGE_OUT", runVar: "QUALITY_JUDGE_RUN", fallbackCost: 0.32, input: "decks", inputVar: "QUALITY_JUDGE_IN"},
	{suite: "revise", target: "revise-gate", prefix: "pr12b-revise-gate-run", docVar: "REVISE_GATE_OUT", runVar: "REVISE_GATE_RUN", fallbackCost: 1.23},
}

// planned is one step with its names and its estimate.
type planned struct {
	step
	number   int
	doc, run string
	estimate float64
	vars     []string
}

// stepRunner runs one Makefile target with its variables from the repo
// root. The tests inject one that writes a run file.
type stepRunner func(root, target string, vars []string) error

func makeRunner(root, target string, vars []string) error {
	cmd := exec.Command("make", append([]string{target}, vars...)...) // #nosec G204 -- the target and the variables are this program's own.
	cmd.Dir = root
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// sweep plans the steps, prints the plan, and runs them unless dry.
func sweep(w io.Writer, root string, suites []string, capUSD float64, dry, keepGoing bool, runner stepRunner) (int, error) {
	if capUSD <= 0 {
		return exitFault, errors.New("give -cap, the most the sweep may spend in USD")
	}
	if root == "" {
		var err error
		if root, err = repoRoot(); err != nil {
			return exitFault, err
		}
	}
	docs := filepath.Join(root, "docs", "reference")
	evalDir := filepath.Join(docs, "eval")
	headers, err := readHeaders(evalDir)
	if err != nil {
		return exitFault, err
	}
	wanted := map[string]bool{}
	for _, s := range suites {
		wanted[s] = true
	}
	var plan []planned
	total := 0.0
	byName := map[string]*planned{}
	for _, st := range sweepSteps {
		if len(wanted) > 0 && !wanted[st.suite] {
			continue
		}
		n, err := nextRunNumber(docs, st.prefix)
		if err != nil {
			return exitFault, err
		}
		p := planned{step: st, number: n}
		p.doc = fmt.Sprintf("docs/reference/%s%d.md", st.prefix, n)
		p.run = fmt.Sprintf("docs/reference/eval/%s%d.jsonl", st.prefix, n)
		p.estimate = lastCost(headers, st.suite, st.fallbackCost)
		p.vars = []string{st.docVar + "=" + p.doc, st.runVar + "=" + p.run}
		if st.input != "" {
			in := byName[st.input]
			inDoc := ""
			if in != nil {
				inDoc = in.doc
			} else if inDoc, err = newestDoc(docs, stepByName(st.input).prefix); err != nil {
				return exitFault, err
			}
			p.vars = append(p.vars, st.inputVar+"="+inDoc)
		}
		if st.suite == "question-eval" {
			p.vars = append(p.vars, fmt.Sprintf("EVAL_JSON=.local/tune/run%d.json", n))
		}
		total += p.estimate
		plan = append(plan, p)
		byName[st.suite] = &plan[len(plan)-1]
	}
	pf := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	pf("# Eval sweep\n\nCap $%.2f. The estimate of a step is the cost of its last run file, or the last measured run when none exists.\n\n", capUSD)
	pf("| Step | Target | Document | Estimate | Running total |\n|---|---|---|---|---|\n")
	running := 0.0
	for _, p := range plan {
		running += p.estimate
		pf("| %s | `make %s` | `%s` | $%.2f | $%.2f |\n", p.suite, p.target, p.doc, p.estimate, running)
	}
	pf("\n")
	if total > capUSD {
		pf("The plan estimates $%.2f, over the cap of $%.2f. The sweep stops before the step that crosses it.\n\n", total, capUSD)
	}
	if dry {
		pf("Dry run: nothing ran, and nothing was spent.\n")
		return exitPass, nil
	}
	if err := gatekit.SpendGuard("EVAL_SWEEP"); err != nil {
		return exitFault, err
	}
	spent := 0.0
	code := exitPass
	for i, p := range plan {
		if spent+p.estimate > capUSD {
			pf("Stopped before %s: $%.2f spent, and the step estimates $%.2f against a cap of $%.2f.\n", p.suite, spent, p.estimate, capUSD)
			break
		}
		pf("## %s\n\n`make %s %s`\n\n", p.suite, p.target, strings.Join(p.vars, " "))
		runErr := runner(root, p.target, p.vars)
		rec, readErr := evalrun.ReadFile(filepath.Join(root, p.run))
		cost := 0.0
		if readErr == nil && rec.Header.CostUSD != nil {
			cost = *rec.Header.CostUSD
		}
		spent += cost
		switch {
		case readErr != nil:
			pf("The step wrote no run file: %v. Spent so far $%.2f.\n\n", readErr, spent)
			code = exitFail
		case rec.Header.Verdict == evalrun.VerdictFail:
			pf("Verdict FAIL, $%.4f, %d calls. Spent so far $%.2f.\n\n", cost, rec.Header.Calls, spent)
			code = exitFail
		case runErr != nil:
			pf("The target exited with an error: %v. The document reads %s, $%.4f. Spent so far $%.2f.\n\n", runErr, orUnnamed(rec.Header.Verdict), cost, spent)
			code = exitFail
		default:
			pf("Verdict %s, $%.4f, %d calls. Spent so far $%.2f.\n\n", orUnnamed(rec.Header.Verdict), cost, rec.Header.Calls, spent)
		}
		if code == exitFail && !keepGoing {
			if i+1 < len(plan) {
				pf("The sweep stops here: a result that is not good stops it, and -continue goes on.\n")
			}
			break
		}
		if spent > capUSD {
			pf("The cap of $%.2f is spent.\n", capUSD)
			break
		}
	}
	pf("\nSpent $%.2f of the cap of $%.2f.\n", spent, capUSD)
	return code, nil
}

func stepByName(suite string) step {
	for _, s := range sweepSteps {
		if s.suite == suite {
			return s
		}
	}
	return step{}
}

// repoRoot is the git root, or the parent of the working directory.
func repoRoot() (string, error) {
	if out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err == nil {
		return strings.TrimSpace(string(out)), nil
	}
	return filepath.Abs("..")
}

var runNumber = regexp.MustCompile(`run(\d+)[a-z]?\.md$`)

// nextRunNumber reads the documents of one family and answers the
// number after the highest, so a rerun's letter never collides.
func nextRunNumber(docs, prefix string) (int, error) {
	entries, err := os.ReadDir(docs)
	if err != nil {
		return 0, err
	}
	best := 0
	for _, e := range entries {
		name := e.Name()
		if !strings.HasPrefix(name, prefix) {
			continue
		}
		m := runNumber.FindStringSubmatch(name)
		if m == nil {
			continue
		}
		if n, _ := strconv.Atoi(m[1]); n > best {
			best = n
		}
	}
	return best + 1, nil
}

// newestDoc names the newest document of a family, relative to the
// repo root, for a step whose input step did not run in this sweep.
func newestDoc(docs, prefix string) (string, error) {
	n, err := nextRunNumber(docs, prefix)
	if err != nil {
		return "", err
	}
	if n == 1 {
		return "", fmt.Errorf("no document of %s exists to read", prefix)
	}
	// A rerun with a letter sits beside its run, and the plain run is
	// the whole document.
	return fmt.Sprintf("docs/reference/%s%d.md", prefix, n-1), nil
}

// lastCost is the cost of the newest full run file of a suite, or the
// fallback. A rerun of one prompt holds fewer items than a full run,
// and its cost is no estimate of the next full run, so the newest run
// among the ones with the most items answers.
func lastCost(headers []fileHeader, suite string, fallback float64) float64 {
	most := 0
	for _, fh := range headers {
		if fh.header.Suite == suite && fh.header.CostUSD != nil && fh.items > most {
			most = fh.items
		}
	}
	best := ""
	var bestHeader evalrun.Header
	for _, fh := range headers {
		if fh.header.Suite != suite || fh.header.CostUSD == nil || fh.items < most {
			continue
		}
		if best == "" || runLess(bestHeader, fh.header) {
			best, bestHeader = fh.name, fh.header
		}
	}
	if best == "" {
		return fallback
	}
	return *bestHeader.CostUSD
}
