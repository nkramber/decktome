// Command eval is the harness of PR-15 over the run files the gates
// write. It compares two runs of one suite and names the flips, it
// checks every suite of the baselines file against its newest run, it
// records a baseline, and it imports a deck gate document written before
// the run files existed.
//
// Every mode is free: no mode calls a provider.
//
// Usage:
//
//	go run ./cmd/eval compare -base a.jsonl[,b.jsonl] -next c.jsonl [-margin 0]
//	go run ./cmd/eval check [-dir ../docs/reference/eval] [-margin 0]
//	go run ./cmd/eval baseline -suite decks -run a.jsonl[,b.jsonl] [-dir ...] [-force]
//	go run ./cmd/eval import -doc ../docs/reference/pr8-deck-gate-run14.md -out ../docs/reference/eval/pr8-deck-gate-run14.jsonl
//	go run ./cmd/eval sweep -cap 5 -dry
//	EVAL_SWEEP=1 go run ./cmd/eval sweep -cap 5 [-suites questions,decks] [-continue]
//
// The sweep is the one paid mode. It runs the Makefile targets, each
// under its own guard, and it stops at the cap or at a FAIL.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
)

// The exit codes.
const (
	exitPass         = 0
	exitFail         = 1
	exitFault        = 2
	exitNotEvaluated = 3
)

// defaultDir is the run directory, from the module root.
const defaultDir = "../docs/reference/eval"

func main() {
	code, err := run(os.Args[1:], os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if code == exitPass {
			code = exitFault
		}
	}
	os.Exit(code)
}

func run(args []string, w io.Writer) (int, error) {
	if len(args) == 0 {
		return exitFault, errors.New("give a mode: compare, check, baseline, import, or sweep")
	}
	mode, rest := args[0], args[1:]
	fs := flag.NewFlagSet("eval "+mode, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	base := fs.String("base", "", "the baseline run files, comma separated, read as one")
	next := fs.String("next", "", "the run file to read against the baseline")
	dir := fs.String("dir", defaultDir, "the run directory")
	margin := fs.Float64("margin", 0, "how far a gate row may move against its sense before it is a regression")
	info := fs.Bool("info", false, "list the information rows that moved, not only their count")
	suite := fs.String("suite", "", "the suite of the baseline")
	runs := fs.String("run", "", "the run files of the baseline, comma separated, relative to -dir")
	force := fs.Bool("force", false, "record a baseline that fails its own bars")
	doc := fs.String("doc", "", "the gate document to import")
	out := fs.String("out", "", "where the import writes the run file")
	capUSD := fs.Float64("cap", 0, "sweep: the most the sweep may spend, in USD")
	suites := fs.String("suites", "", "sweep: the suites to run, comma separated, default every one")
	dry := fs.Bool("dry", false, "sweep: print the plan and run nothing")
	keepGoing := fs.Bool("continue", false, "sweep: go on after a FAIL")
	root := fs.String("root", "", "sweep: the repo root, default the git root")
	if err := fs.Parse(rest); err != nil {
		return exitFault, fmt.Errorf("eval %s: %w", mode, err)
	}
	switch mode {
	case "compare":
		return compareFiles(w, splitList(*base), *next, *margin, *info)
	case "check":
		return check(w, *dir, *margin, *info)
	case "baseline":
		return exitPass, setBaseline(w, *dir, *suite, splitList(*runs), *force)
	case "import":
		return exitPass, importDoc(w, *doc, *out)
	case "sweep":
		return sweep(w, *root, splitList(*suites), *capUSD, *dry, *keepGoing, makeRunner)
	}
	return exitFault, fmt.Errorf("unknown mode %q: give compare, check, baseline, import, or sweep", mode)
}

func splitList(s string) []string {
	var out []string
	for _, p := range strings.Split(s, ",") {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// readRuns reads the files and merges them into one run.
func readRuns(paths []string) (*evalrun.Run, error) {
	var runs []*evalrun.Run
	for _, p := range paths {
		r, err := evalrun.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
		if len(runs) > 0 && runs[0].Header.Suite != r.Header.Suite {
			return nil, fmt.Errorf("%s is suite %s, and the run before it is suite %s", p, r.Header.Suite, runs[0].Header.Suite)
		}
		runs = append(runs, r)
	}
	return evalrun.Merge(runs...), nil
}

// compareFiles reads next against base and prints the comparison.
func compareFiles(w io.Writer, basePaths []string, nextPath string, margin float64, info bool) (int, error) {
	if nextPath == "" {
		return exitFault, errors.New("give -next, a run file")
	}
	next, err := evalrun.ReadFile(nextPath)
	if err != nil {
		return exitFault, err
	}
	var base *evalrun.Run
	if len(basePaths) > 0 {
		if base, err = readRuns(basePaths); err != nil {
			return exitFault, err
		}
		if base.Header.Suite != next.Header.Suite {
			return exitFault, fmt.Errorf("the base is suite %s and the next is suite %s: a compare reads one suite", base.Header.Suite, next.Header.Suite)
		}
	}
	c := evalrun.Compare(base, next, margin)
	writeComparison(w, c, info)
	return exitCode(c.Verdict), nil
}

func exitCode(verdict string) int {
	switch verdict {
	case evalrun.VerdictPass:
		return exitPass
	case evalrun.VerdictFail:
		return exitFail
	}
	return exitNotEvaluated
}

// writeComparison prints one comparison as Markdown. The gate flips are
// the table. The information rows fold into one count, because the cost
// rows move with the daily prices on every run, and info lists them.
func writeComparison(w io.Writer, c evalrun.Comparison, info bool) {
	p := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	p("## Suite `%s`: %s\n\n", c.Next.Suite, c.Verdict)
	if c.Base.Suite == "" {
		p("Run `%s` of %s, with no baseline. %s.\n\n", orUnnamed(c.Next.RunID), c.Next.Date, strings.Join(c.Reasons, ". "))
	} else {
		p("Baseline `%s` of %s against run `%s` of %s. %s.\n\n", orUnnamed(c.Base.RunID), c.Base.Date, orUnnamed(c.Next.RunID), c.Next.Date, strings.Join(c.Reasons, ". "))
	}
	if len(c.Epoch) > 0 {
		p("The epoch moved:\n\n")
		for _, note := range c.Epoch {
			p("- %s\n", note)
		}
		p("\n")
	}
	var gate, infos []evalrun.Flip
	infoWorse := 0
	for _, f := range c.Flips {
		if f.Kind == evalrun.KindGate {
			gate = append(gate, f)
			continue
		}
		infos = append(infos, f)
		if f.Worse {
			infoWorse++
		}
	}
	listed := gate
	if info {
		listed = c.Flips
	}
	if len(listed) > 0 {
		p("| Item | Metric | Kind | Before | After | Read |\n|---|---|---|---|---|---|\n")
		for _, f := range listed {
			read := "better"
			if f.Worse {
				read = "WORSE"
			}
			detail := f.AfterDetail
			if detail == "" {
				detail = f.BeforeDetail
			}
			if detail != "" {
				read += ", " + cell(detail)
			}
			p("| %s | `%s` | %s | %s | %s | %s |\n", f.Item, f.Metric, f.Kind, number(f.Before), number(f.After), read)
		}
		p("\n")
	} else if c.Base.Suite != "" && len(c.Flips) == 0 {
		p("No row moved.\n\n")
	}
	if len(infos) > 0 && !info {
		p("%d information rows moved, %d of them worse. Pass -info to list them.\n\n", len(infos), infoWorse)
	}
	if len(c.Gone) > 0 {
		p("Items of the baseline the run lacks: %s.\n\n", strings.Join(c.Gone, ", "))
	}
	if len(c.New) > 0 {
		p("Items new in the run: %s.\n\n", strings.Join(c.New, ", "))
	}
}

func number(v float64) string {
	if v == float64(int64(v)) {
		return fmt.Sprintf("%d", int64(v))
	}
	return fmt.Sprintf("%.4g", v)
}

func cell(s string) string {
	s = strings.ReplaceAll(strings.ReplaceAll(s, "|", "/"), "\n", " ")
	if len(s) > 80 {
		s = s[:77] + "..."
	}
	return s
}

func orUnnamed(s string) string {
	if s == "" {
		return "unnamed"
	}
	return s
}

// check reads every suite of the baselines file against the newest run
// of that suite in the directory. It is the Tier 0 job: free, over the
// committed files alone.
func check(w io.Writer, dir string, margin float64, info bool) (int, error) {
	bl, err := readBaselines(dir)
	if err != nil {
		return exitFault, err
	}
	p := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	p("# Eval check\n\n")
	if len(bl) == 0 {
		p("No baseline is recorded in %s, so there is nothing to check.\n", filepath.Join(dir, baselinesFile))
		return exitPass, nil
	}
	headers, err := readHeaders(dir)
	if err != nil {
		return exitFault, err
	}
	code := exitPass
	for _, suite := range sortedSuites(bl) {
		files := bl[suite]
		base, err := readRuns(prefixed(dir, files))
		if err != nil {
			return exitFault, fmt.Errorf("baseline of %s: %w", suite, err)
		}
		newest := newestRun(headers, suite, files)
		if newest == "" {
			// Nothing was compared, so the suite reads NOT EVALUATED and never
			// PASS. The exit code stays green: nothing regressed.
			p("## Suite `%s`: NOT EVALUATED\n\nThe baseline `%s` stands alone, with no newer run.\n\n", suite, base.Header.RunID)
			continue
		}
		next, err := evalrun.ReadFile(filepath.Join(dir, newest))
		if err != nil {
			return exitFault, err
		}
		c := evalrun.Compare(base, next, margin)
		writeComparison(w, c, info)
		if c.Verdict == evalrun.VerdictFail {
			code = exitFail
		}
	}
	return code, nil
}

func prefixed(dir string, files []string) []string {
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, filepath.Join(dir, f))
	}
	return out
}

// importDoc reads a deck gate document into a run file.
func importDoc(w io.Writer, doc, out string) error {
	if doc == "" || out == "" {
		return errors.New("give -doc, a deck gate document, and -out, the run file to write")
	}
	if err := gatekit.RefuseExisting(out); err != nil {
		return err
	}
	f, err := os.Open(doc) // #nosec G304 -- the operator names the file.
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	run, err := importDeckGate(f, evalrun.RunID(doc), filepath.Base(doc))
	if err != nil {
		return fmt.Errorf("import %s: %w", doc, err)
	}
	if err := evalrun.WriteFile(out, run); err != nil {
		return err
	}
	_, _ = fmt.Fprintf(w, "imported %s: %d rows of %d items, verdict %s, into %s\n", filepath.Base(doc), len(run.Rows), countItems(run), run.Header.Verdict, out)
	return nil
}

func countItems(r *evalrun.Run) int {
	seen := map[string]bool{}
	for _, row := range r.Rows {
		seen[row.Item] = true
	}
	return len(seen)
}
