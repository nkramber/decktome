// Command case-check reads a gate run file and says whether each case of
// a triage manifest passed (PR-28c, D-645).
//
// The fix cycle runs it twice. Before the fixer it must report every
// case as a failure, because a case is the fault a reader met and a case
// that already passes never measured it. After the fixer it must report
// every one as a pass.
//
// It calls no model and it costs nothing. It reads the run file the gate
// wrote, so it reads the same bars the gate's own verdict reads.
//
// Usage:
//
//	go run ./cmd/case-check -manifest .local/tune/cases.json \
//	  -run docs/reference/eval/pr7-question-gate-cases.jsonl -gate questions
//	go run ./cmd/case-check -manifest ... -run ... -gate questions -want fail
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/triage"
)

// exitFault is the code of a fault in the tool, apart from a case that
// did not read the way the caller wanted. The cycle stops on a fault and
// judges nothing (T-12).
const exitFault = 2

func main() {
	code, err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(code)
}

func run() (int, error) {
	manifest := flag.String("manifest", "", "the manifest the triage wrote")
	runFile := flag.String("run", "", "the run file the gate wrote")
	gate := flag.String("gate", "", "read the cases of this gate alone: questions, decks, or brackets")
	want := flag.String("want", "pass", "what every case must read: pass or fail")
	flag.Parse()
	if *manifest == "" || *runFile == "" {
		return exitFault, fmt.Errorf("give -manifest and -run")
	}
	if *want != "pass" && *want != "fail" {
		return exitFault, fmt.Errorf("-want reads pass or fail, and not %q", *want)
	}

	m, err := triage.ReadManifest(*manifest)
	if err != nil {
		return exitFault, err
	}
	rec, err := evalrun.ReadFile(*runFile)
	if err != nil {
		return exitFault, fmt.Errorf("case-check: %s: %w", *runFile, err)
	}
	return check(os.Stdout, m, rec, *gate, *want, *runFile)
}

// check writes one row per case and answers the exit code. It is the
// whole tool, apart from the flags and the two file reads.
func check(w io.Writer, m triage.Manifest, rec *evalrun.Run, gate, want, runFile string) (int, error) {
	var cases []triage.Entry
	for _, c := range m.Cases {
		if gate == "" || c.Gate == gate {
			cases = append(cases, c)
		}
	}
	if len(cases) == 0 {
		return exitFault, fmt.Errorf("case-check: the manifest names no case of gate %q", gate)
	}
	// A run of the whole suite answers a case as well, but the cycle
	// spends on a partial run of the case ids alone. A run that holds no
	// row for a case measured something else, and that is a fault and
	// never a failure.
	lower := map[string]bool{}
	for _, k := range rec.Header.Lower {
		lower[k] = true
	}

	bad, missing := 0, 0
	_, _ = fmt.Fprintf(w, "| Case | Class | Read | Why |\n|---|---|---|---|\n")
	for _, c := range cases {
		fails, rows := failuresOf(rec, c.ID, lower)
		read := "pass"
		if len(fails) > 0 {
			read = "fail"
		}
		why := "every bar of the case is good"
		if rows == 0 {
			read, why = "no row", fmt.Sprintf("the run holds no gate row for item %d", c.ID)
			missing++
		} else if len(fails) > 0 {
			why = strings.Join(fails, "; ")
		}
		_, _ = fmt.Fprintf(w, "| %d | %s | %s | %s |\n", c.ID, c.Class, read, why)
		if rows > 0 && read != want {
			bad++
		}
	}
	_, _ = fmt.Fprintln(w)
	if missing > 0 {
		return exitFault, fmt.Errorf("case-check: %d case(s) have no row in %s, so the run measured something else", missing, runFile)
	}
	if bad > 0 {
		_, _ = fmt.Fprintf(w, "%d of %d cases did not read %s.\n", bad, len(cases), want)
		return 1, nil
	}
	_, _ = fmt.Fprintf(w, "Every one of the %d cases reads %s.\n", len(cases), want)
	return 0, nil
}

// failuresOf names every gate bar of one item that is not at its good
// value, and it counts the gate rows it read. A metric the run marks
// lower-is-better is good at zero, and every other gate bar is good at
// one: the gates write them as flags.
func failuresOf(rec *evalrun.Run, id int, lower map[string]bool) ([]string, int) {
	item := strconv.Itoa(id)
	var fails []string
	rows := 0
	for _, r := range rec.Rows {
		if r.Item != item || r.Kind != evalrun.KindGate {
			continue
		}
		rows++
		good := r.Value >= 1
		if lower[r.Metric] {
			good = r.Value == 0
		}
		if good {
			continue
		}
		detail := r.Detail
		if detail == "" {
			detail = fmt.Sprintf("%g", r.Value)
		}
		fails = append(fails, r.Metric+": "+detail)
	}
	sort.Strings(fails)
	return fails, rows
}
