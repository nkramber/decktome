// Command tune-check decides whether one automated iteration may be kept
// (D-133). It reads the eval summary of the new run, and the summary of
// the run before it, and it exits non-zero when the new run is worse.
//
// It costs nothing and calls no provider. The loop driver runs it after
// every iteration, and it reverts the working tree when this command
// fails.
//
// With -changes, the fixer's commits are judged one by one. Each commit
// declares the catalog rows it means to move, and every question whose
// verdict changed is charged to the commit that owns its row. The loop
// then keeps the commits that helped and drops the ones that hurt, and
// it measures the kept rows again before they become the baseline
// (D-181). Exit 4 says so, and the keep, drop, and remeasure lines on
// stdout say which.
//
// With -lessons, the command appends what it learned to a file the next
// fixer reads (D-182). With -merge, it folds a partial run into a full
// one. With -fixer, it writes the summary the fixer's checkout may hold:
// the same file with every holdout verdict removed (T-8). With
// -print-noise, it prints the noise margin the loop reads (T-7).
//
// Exit 2 is a tool fault: an unreadable file, a partial summary, or two
// commits that own one row. The loop stops on it and rejects nothing
// (T-12).
//
// Usage:
//
//	go run ./cmd/tune-check -next ../.local/tune/iter-03.json \
//	  -prev ../.local/tune/iter-02.json -target 0.05 \
//	  -changes ../.local/tune/iter-03-changes.json \
//	  -lessons ../docs/reference/autotune-lessons.md
//
//	go run ./cmd/tune-check -merge -prev base.json -next part.json \
//	  -out merged.json -out-doc merged.md
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/nkramber/mtg-deck-builder/go/internal/tune"
)

func main() {
	next := flag.String("next", "", "the eval summary of the new run")
	prev := flag.String("prev", "", "the eval summary of the run before it (optional)")
	target := flag.Float64("target", 0.05, "stop the loop once the bad-question ratio is at or under this")
	agree := flag.Bool("agree", false, "compare two summaries of one run, question by question, instead of deciding")
	changes := flag.String("changes", "", "the fixer's changes, one commit each, as JSON (optional)")
	noise := flag.Int("noise", tune.DefaultNoise, "bad questions two runs of the same code may differ by")
	lessons := flag.String("lessons", "", "append what this decision taught to this file (optional)")
	label := flag.String("label", "", "the iteration label, for the lessons file")
	merge := flag.Bool("merge", false, "fold a partial run (-next) into a full one (-prev) and write -out")
	prevDoc := flag.String("prev-doc", "", "the gate document of -prev, for -merge of a summary from before D-181")
	nextDoc := flag.String("next-doc", "", "the gate document of -next, for -merge of a summary from before D-181")
	out := flag.String("out", "", "where -merge writes the merged summary")
	outDoc := flag.String("out-doc", "", "where -merge writes the merged report the fixer reads (optional)")
	fixer := flag.Bool("fixer", false, "write -next to -out with every holdout verdict removed")
	printNoise := flag.Bool("print-noise", false, "print the default noise margin and exit")
	flag.Parse()
	switch {
	case *printNoise:
		fmt.Println(tune.DefaultNoise)
		return
	case *fixer:
		if err := fixerCopy(*next, *out); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(exitFault)
		}
		return
	case *agree:
		if err := agreement(*next, *prev, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(exitFault)
		}
		return
	case *merge:
		if err := mergeRuns(*prev, *prevDoc, *next, *nextDoc, *out, *outDoc, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(exitFault)
		}
		return
	}
	code, err := run(*next, *prev, *changes, *target, *noise, *lessons, *label, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFault)
	}
	os.Exit(code)
}

// The exit codes the loop driver reads.
const (
	exitAccept = 0
	// exitReject says the iteration made things worse. The driver reverts.
	exitReject = 1
	// exitFault says the tool could not decide. The driver stops.
	exitFault = 2
	// exitDone says the run is good enough and the loop may stop.
	exitDone = 3
	// exitPartial says some changes are kept and some are dropped. The
	// driver keeps the commits on the keep lines, drops the rest, and
	// runs the gate again on the remeasure conversations (D-181).
	exitPartial = 4
)

func run(nextPath, prevPath, changesPath string, target float64, noise int, lessonsPath, label string, w io.Writer) (int, error) {
	if nextPath == "" {
		return 0, fmt.Errorf("give -next, an eval summary")
	}
	next, err := readWhole(nextPath)
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
		p, err := readWhole(prevPath)
		if err != nil {
			return 0, fmt.Errorf("-prev %s: %w", prevPath, err)
		}
		prev = p
	}
	var changes []tune.Change
	if changesPath != "" {
		raw, err := os.ReadFile(changesPath) // #nosec G304 -- the caller names the file.
		if err != nil {
			return 0, fmt.Errorf("-changes %s: %w", changesPath, err)
		}
		if err := json.Unmarshal(raw, &changes); err != nil {
			return 0, fmt.Errorf("-changes %s: %w", changesPath, err)
		}
	}
	d, paired, attr, err := tune.Decide(prev, next, changes, noise)
	if err != nil {
		return 0, err
	}
	for _, r := range d.Reasons {
		switch {
		case d.Accept:
			_, _ = fmt.Fprintf(w, "accept: %s\n", r)
		case d.Partial:
			_, _ = fmt.Fprintf(w, "partial: %s\n", r)
		default:
			_, _ = fmt.Fprintf(w, "reject: %s\n", r)
		}
	}
	if prev != nil {
		_, _ = fmt.Fprintf(w, "paired: %d shared, %d for, %d against, %d judge flips, %d churn, %d errored conversations\n",
			paired.Shared, paired.For(), paired.Against(), len(paired.JudgeFlips), len(attr.Churn), paired.Errored)
	}
	for _, cv := range attr.Changes {
		verb := "drop"
		if cv.Keep {
			verb = "keep"
		}
		_, _ = fmt.Fprintf(w, "%s: %s %s (%s)\n", verb, cv.Change.Commit, cv.Change.Subject, cv.Reason)
	}
	if len(d.Remeasure) > 0 {
		_, _ = fmt.Fprintf(w, "remeasure: %s\n", joinInts(d.Remeasure))
	}
	if lessonsPath != "" {
		if err := appendLessons(lessonsPath, label, prev, next, d, paired, attr); err != nil {
			return 0, err
		}
	}
	switch {
	case d.Partial:
		return exitPartial, nil
	case !d.Accept:
		return exitReject, nil
	}
	// The target reads the holdout ratio when there is one, as the docs
	// say (T-14).
	if scored, split := next.Scored(); next.Judged > 0 && scored <= target {
		_, _ = fmt.Fprintf(w, "done: the %s ratio is %.1f%%, at or under the target of %.1f%%\n",
			split, scored*100, target*100)
		return exitDone, nil
	}
	return exitAccept, nil
}

func joinInts(ns []int) string {
	parts := make([]string, len(ns))
	for i, n := range ns {
		parts[i] = strconv.Itoa(n)
	}
	return strings.Join(parts, ",")
}

// appendLessons writes what one decision taught, in the form the next
// fixer reads. A rejected change is named with its hypothesis and the
// questions it made worse, so the same idea is not tried twice (D-182).
func appendLessons(path, label string, prev, next *tune.Summary, d tune.Decision, p tune.Paired, a tune.Attribution) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600) // #nosec G304 -- the caller names the file.
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	var b strings.Builder
	w := func(format string, args ...any) { b.WriteString(fmt.Sprintf(format, args...)) }
	outcome := "rejected"
	switch {
	case d.Accept:
		outcome = "accepted"
	case d.Partial:
		outcome = "partly kept"
	}
	if label == "" {
		label = next.Run
	}
	w("\n## %s, %s (%s)\n\n", label, outcome, time.Now().UTC().Format("2006-01-02"))
	if prev != nil {
		pr, split := prev.Scored()
		nr, _ := next.Scored()
		w("The %s ratio moved from %.1f%% to %.1f%%. ", split, pr*100, nr*100)
		w("Paired: %d questions got better, %d got worse, and %d verdicts flipped on identical text.\n\n",
			p.For(), p.Against(), len(p.JudgeFlips))
	}
	for _, r := range d.Reasons {
		w("- %s\n", r)
	}
	if len(a.Changes) > 0 {
		w("\n| Change | Rows | Verdict | Why |\n|---|---|---|---|\n")
		for _, cv := range a.Changes {
			verdict := "dropped"
			if cv.Keep {
				verdict = "kept"
			}
			w("| %s %s | %s | %s | %s |\n", short(cv.Change.Commit), cv.Change.Subject,
				strings.Join(cv.Change.Rows, ", "), verdict, cv.Reason)
		}
		for _, cv := range a.Changes {
			if cv.Keep && len(cv.Against) == 0 {
				continue
			}
			w("\n### %s: %s\n\n", short(cv.Change.Commit), cv.Change.Subject)
			if cv.Change.Hypothesis != "" {
				w("Hypothesis: %s\n\n", cv.Change.Hypothesis)
			}
			if !cv.Keep {
				w("Do not try this idea again in the same form. It made these questions worse:\n\n")
			} else {
				w("Kept, and these questions still got worse on its rows:\n\n")
			}
			for _, x := range cv.Against {
				w("- %s, turn %d, row `%s`: %q. %s\n", x.Conversation, x.Turn, x.Row, x.After.Text, x.After.Reason)
			}
		}
	}
	if n := len(a.Unattributed); n > 0 {
		w("\n%d questions moved on rows no change declared. Declare every row a change can touch.\n", n)
		rows := map[string]int{}
		for _, u := range a.Unattributed {
			rows[u.Row]++
		}
		names := make([]string, 0, len(rows))
		for r := range rows {
			names = append(names, r)
		}
		sort.Strings(names)
		for _, r := range names {
			w("- `%s`: %d\n", r, rows[r])
		}
	}
	_, err = f.WriteString(b.String())
	return err
}

func short(sha string) string {
	if len(sha) > 7 {
		return sha[:7]
	}
	return sha
}

// mergeRuns folds a partial run into a full one and writes the merged
// summary, and a short report the fixer reads (D-181). A summary from
// before D-181 carries no per-conversation counts. Name its gate document
// with -prev-doc, and the counts are read from there.
func mergeRuns(basePath, baseDoc, partPath, partDoc, outPath, outDoc string, w io.Writer) error {
	for name, val := range map[string]string{"-prev": basePath, "-next": partPath, "-out": outPath} {
		if val == "" {
			return fmt.Errorf("-merge needs %s", name)
		}
	}
	for _, p := range []string{outPath, outDoc} {
		if p == "" {
			continue
		}
		if _, err := os.Stat(p); err == nil {
			return fmt.Errorf("%s exists, and a result is never overwritten (D-65)", p)
		}
	}
	base, err := readWhole(basePath)
	if err != nil {
		return err
	}
	part, err := readWhole(partPath)
	if err != nil {
		return err
	}
	if err := fillCounts(base, baseDoc); err != nil {
		return err
	}
	if err := fillCounts(part, partDoc); err != nil {
		return err
	}
	merged, err := tune.Merge(base, part)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(outPath), 0o750); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(merged, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(outPath, append(raw, '\n'), 0o600); err != nil {
		return err
	}
	if outDoc != "" {
		if err := os.WriteFile(outDoc, []byte(mergedReport(base, part, merged)), 0o600); err != nil {
			return err
		}
	}
	scored, split := merged.Scored()
	_, _ = fmt.Fprintf(w, "merged %d conversations of %s into %s: %d judged, %.1f%% on the %s\n",
		len(part.Conversations), part.Run, base.Run, merged.Judged, scored*100, split)
	return nil
}

func fillCounts(s *tune.Summary, doc string) error {
	if len(s.Conversations) > 0 || doc == "" {
		return nil
	}
	run, err := tune.ReadRun(doc)
	if err != nil {
		return err
	}
	s.Conversations = tune.CountConversations(run)
	return nil
}

// mergedReport writes the merged run in the shape the fixer reads: the
// counters, the rows, and every refused question outside the holdout.
func mergedReport(base, part *tune.Summary, s tune.Summary) string {
	var b strings.Builder
	p := func(format string, a ...any) { b.WriteString(fmt.Sprintf(format, a...)) }
	p("# PR-7 question eval, merged\n\n")
	p("Run: `%s`. It folds %d conversations of `%s` into `%s` (D-181). Eval model: `%s`.\n\n",
		s.Run, len(part.Conversations), part.Run, base.Run, s.Model)
	p("Scored %d questions. %d were not warranted, and %d are unsure.\n\n", s.Judged, s.Bad, s.Unsure)
	scored, split := s.Scored()
	p("**Bad-question ratio: %.1f%% on the %s.**\n\n", scored*100, split)
	if s.HoldoutJudged > 0 {
		p("The tune split holds %d questions at %.1f%%, and the holdout holds %d at %.1f%%.\n\n",
			s.TuneJudged, s.TuneRatio*100, s.HoldoutJudged, s.HoldoutRatio*100)
	}
	p("## The counters the ratio can not see\n\n| Counter | Value |\n|---|---|\n")
	p("| Questions asked | %d |\n| Questions that closed a slot | %d |\n| Invented by the model | %d |\n",
		s.Metrics.Questions, s.Metrics.CatalogFilled, s.Metrics.Invented)
	p("| Catalog-only conversations | %d |\n| Premature sessions | %d |\n| Linter findings | %d |\n\n",
		s.Metrics.CatalogOnly, s.Metrics.Premature, s.Metrics.LintFindings)
	if len(s.ByRow) > 0 {
		p("## Where the bad questions came from\n\n| Row | Bad questions |\n|---|---|\n")
		for _, row := range s.TopRows(0) {
			p("| `%s` | %d |\n", row, s.ByRow[row])
		}
		p("\n")
	}
	p("## Every question the eval refused\n\n")
	if s.HoldoutJudged > 0 {
		p("The holdout conversations are left out on purpose. %d of their questions failed, and the fixer may not read which (D-134).\n\n", s.HoldoutBad)
	}
	n := 0
	for _, v := range s.Verdicts {
		if !v.Bad() || v.Holdout {
			continue
		}
		n++
		p("### %s, turn %d, row `%s`\n\n**Asked:** %s\n\n", v.Conversation, v.Turn, v.Row, v.Text)
		if len(v.Faults) > 0 {
			p("Faults: %s. ", strings.Join(v.Faults, ", "))
		}
		p("Catalog action: %s.\n\n%s\n\n", orNone(v.CatalogAction), v.Reason)
	}
	if n == 0 {
		p("None.\n\n")
	}
	p("## Run\n\n- Cost: $%.4f.\n", s.CostUSD)
	return b.String()
}

func orNone(s string) string {
	if s == "" {
		return "none"
	}
	return s
}

// agreement measures one eval model against another on the same run. It
// answers the question the owner must not guess at: how gently does a
// model score work its own model produced (OQ-26)?
//
// The eval role runs on the model that also writes the questions. That is
// the owner's call for cost, and this is how the cost of that call gets
// measured. It also measures one judge against itself: two scorings of
// one document show the judge's own noise (D-183).
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
		index[tune.PairKey(v)] = v
	}
	var both, same, aBad, bBad, bothBad int
	for _, v := range a.Verdicts {
		other, ok := index[tune.PairKey(v)]
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
	_, _ = fmt.Fprintf(w, "verdicts that differ: %d, which is the noise floor for one judge on one document\n", both-same)
	if bBad > aBad {
		_, _ = fmt.Fprintf(w, "WARNING: %s is the gentler judge by %d questions. Read that as a floor on the real ratio.\n",
			a.Model, bBad-aBad)
	}
	return nil
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

// readWhole reads a summary and refuses a partial one. A budget-cut eval
// scores fewer questions, and a smaller count read as an improvement
// once (T-2).
func readWhole(path string) (*tune.Summary, error) {
	s, err := read(path)
	if err != nil {
		return nil, err
	}
	if why, skipped := s.Skipped(); skipped {
		return nil, fmt.Errorf("%s is a partial summary (%s), and a partial run is no baseline and no candidate", path, why)
	}
	return s, nil
}

// fixerCopy writes the summary with every holdout verdict removed. The
// counts stay, so the fixer reads the holdout ratio and never which
// questions made it (T-8, D-134).
func fixerCopy(in, out string) error {
	if in == "" || out == "" {
		return fmt.Errorf("-fixer needs -next and -out")
	}
	s, err := read(in)
	if err != nil {
		return err
	}
	kept := make([]tune.Verdict, 0, len(s.Verdicts))
	for _, v := range s.Verdicts {
		if !v.Holdout {
			kept = append(kept, v)
		}
	}
	s.Verdicts = kept
	raw, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(out, append(raw, '\n'), 0o600)
}
