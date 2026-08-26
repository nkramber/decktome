// Command tune-check decides whether one automated iteration may be kept
// (D-133). It reads the eval summary of the new run, and the summary of
// the run before it, and it exits non-zero when the new run is worse.
//
// It costs nothing and calls no provider. The loop driver runs it after
// every iteration, and it reverts the working tree when this command
// fails.
//
// Usage:
//
//	go run ./cmd/tune-check -next ../.local/tune/iter-03.json \
//	  -prev ../.local/tune/iter-02.json -target 0.05
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/nkramber/mtg-deck-builder/go/internal/tune"
)

func main() {
	next := flag.String("next", "", "the eval summary of the new run")
	prev := flag.String("prev", "", "the eval summary of the run before it (optional)")
	target := flag.Float64("target", 0.05, "stop the loop once the bad-question ratio is at or under this")
	agree := flag.Bool("agree", false, "compare two summaries of one run, question by question, instead of deciding")
	flag.Parse()
	if *agree {
		if err := agreement(*next, *prev, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		return
	}
	code, err := run(*next, *prev, *target, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	os.Exit(code)
}

// The exit codes the loop driver reads.
const (
	exitAccept = 0
	// exitReject says the iteration made things worse. The driver reverts.
	exitReject = 1
	// exitDone says the run is good enough and the loop may stop.
	exitDone = 3
)

func run(nextPath, prevPath string, target float64, w *os.File) (int, error) {
	if nextPath == "" {
		return 0, fmt.Errorf("give -next, an eval summary")
	}
	next, err := read(nextPath)
	if err != nil {
		return 0, err
	}
	// A -prev that was named and can not be read is an error, and never a
	// first run. The loop passed a relative path into a subshell that had
	// changed directory, so the file was missing, the comparison was
	// skipped in silence, and an iteration that raised the holdout ratio
	// was accepted and committed (D-171).
	var prev *tune.Summary
	if prevPath != "" {
		p, err := read(prevPath)
		if err != nil {
			return 0, fmt.Errorf("-prev %s: %w", prevPath, err)
		}
		prev = p
	}
	d := tune.Compare(prev, next)
	for _, r := range d.Reasons {
		if d.Accept {
			_, _ = fmt.Fprintf(w, "accept: %s\n", r)
		} else {
			_, _ = fmt.Fprintf(w, "reject: %s\n", r)
		}
	}
	if !d.Accept {
		return exitReject, nil
	}
	if next.Judged > 0 && next.Ratio <= target {
		_, _ = fmt.Fprintf(w, "done: the ratio is %.1f%%, at or under the target of %.1f%%\n",
			next.Ratio*100, target*100)
		return exitDone, nil
	}
	return exitAccept, nil
}

// agreement measures one eval model against another on the same run. It
// answers the question the owner must not guess at: how gently does a
// model score work its own model produced (OQ-26)?
//
// The eval role runs on the model that also writes the questions. That is
// the owner's call for cost, and this is how the cost of that call gets
// measured.
func agreement(aPath, bPath string, w io.Writer) error {
	a, err := read(aPath)
	if err != nil {
		return err
	}
	b, err := read(bPath)
	if err != nil {
		return err
	}
	index := map[string]tune.Verdict{}
	for _, v := range b.Verdicts {
		index[verdictKey(v)] = v
	}
	var both, same, aBad, bBad, bothBad int
	for _, v := range a.Verdicts {
		other, ok := index[verdictKey(v)]
		if !ok {
			continue
		}
		both++
		if v.Warranted == other.Warranted {
			same++
		}
		switch {
		case v.Bad() && other.Bad():
			bothBad++
			aBad++
			bBad++
		case v.Bad():
			aBad++
		case other.Bad():
			bBad++
		}
	}
	if both == 0 {
		return fmt.Errorf("the two summaries share no question")
	}
	_, _ = fmt.Fprintf(w, "%s (%s) against %s (%s)\n", a.Run, a.Model, b.Run, b.Model)
	_, _ = fmt.Fprintf(w, "shared questions: %d\n", both)
	_, _ = fmt.Fprintf(w, "the same verdict: %d (%.0f%%)\n", same, float64(same)/float64(both)*100)
	_, _ = fmt.Fprintf(w, "%s refused %d, %s refused %d, both refused %d\n", a.Model, aBad, b.Model, bBad, bothBad)
	if bBad > aBad {
		_, _ = fmt.Fprintf(w, "WARNING: %s is the gentler judge by %d questions. Read that as a floor on the real ratio.\n",
			a.Model, bBad-aBad)
	}
	return nil
}

func verdictKey(v tune.Verdict) string {
	return fmt.Sprintf("%s|%d|%s", v.Conversation, v.Turn, v.Row)
}

func read(path string) (*tune.Summary, error) {
	raw, err := os.ReadFile(path) // #nosec G304 -- the caller names the summary.
	if err != nil {
		return nil, err
	}
	var s tune.Summary
	if err := json.Unmarshal(raw, &s); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &s, nil
}
