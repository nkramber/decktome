// Command sixty-gate is the gate of the power judge of a 60-card import
// (PR-71). It reads labeled lists of the local meta store, asks the
// judge for the step of each, and writes a gate document to stdout.
//
// Casual reads the Theme, Intro, and Planeswalker decks of the precon
// table, FNM the MTGGoldfish user decks that a model of another family
// labeled fnm (D-875), and tournament a top 8 finish in an MTGO Challenge
// or an MTGTop8 RCQ (D-863). The gate passes when 80
// percent of the test reads name the label, and no read sits two steps
// off (D-868). A FAIL verdict exits 2.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/gcpenv"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/meta"
)

// passRate is the share of the test reads that must name the label
// (D-868).
const passRate = 0.80

// errFail is the verdict FAIL. main exits 2 on it, so a script tells a
// failed gate from a crash.
var errFail = errors.New("sixty-gate: the verdict is FAIL")

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if errors.Is(err, errFail) {
			os.Exit(2)
		}
		os.Exit(1)
	}
}

func run(w io.Writer) error {
	dry := flag.Bool("dry", false, "print the split and stop before the provider calls")
	split := flag.String("split", "test", "judge the test split, which the verdict reads, or the dev split")
	reads := flag.Int("reads", 3, "judge reads of each list")
	flag.Parse()
	if *split != "test" && *split != "dev" {
		return fmt.Errorf("sixty-gate: -split %q, want test or dev", *split)
	}
	if *reads < 1 {
		return fmt.Errorf("sixty-gate: -reads %d, want 1 or more", *reads)
	}
	if !*dry {
		if err := gatekit.SpendGuard("SIXTY_GATE"); err != nil {
			return err
		}
	}
	ctx := context.Background()
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(ctx, quiet)
	if err != nil {
		return err
	}
	store, err := gcpenv.MetaStore(ctx, "", nil)
	if err != nil {
		return err
	}
	version, err := meta.LatestPreconsVersion(ctx, store)
	if err != nil {
		return err
	}
	precons, err := meta.ReadPrecons(ctx, store, version)
	if err != nil {
		return err
	}
	modern, err := meta.AllLists(ctx, store, "modern")
	if err != nil {
		return err
	}
	standard, err := meta.AllLists(ctx, store, "standard")
	if err != nil {
		return err
	}
	fnm, err := fnmKeys()
	if err != nil {
		return err
	}
	set, err := pick(idx, precons, modern, standard, fnm)
	if err != nil {
		return err
	}
	head := header{snapshot: idx.AsOf.Format("2006-01-02"), precons: version, split: *split, reads: *reads}
	if *dry {
		writeSplit(w, head, set)
		return nil
	}

	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return err
	}
	spec := client.Config().Roles[llm.RoleJudge]
	head.judge = spec.Provider + " " + spec.Model
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)
	start := time.Now()
	var results []result
	for _, l := range set.lists {
		if l.dev != (*split == "dev") {
			continue
		}
		r := result{list: l}
		for range *reads {
			j, err := generate.JudgeSixtyStep(ctx, client, l.deck, l.format, idx, acc)
			if err != nil {
				r.steps = append(r.steps, mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED)
				if r.err == "" {
					r.err = err.Error()
				}
				continue
			}
			r.steps = append(r.steps, j.Step)
			if r.why == "" {
				r.why = j.Why
			}
		}
		results = append(results, r)
	}
	head.took = time.Since(start).Round(time.Second)
	head.cost = acc.Report()
	t := tally(results)
	writeReport(w, head, set, results, t)
	if *split == "test" && !t.pass() {
		return errFail
	}
	return nil
}

// header is what the document says about the run.
type header struct {
	snapshot string
	precons  string
	split    string
	reads    int
	judge    string
	took     time.Duration
	cost     llm.Report
}

// result is the reads of one list. An unspecified step is a judge error,
// and err keeps the first one.
type result struct {
	list  calList
	steps []mtgv1.SixtyStep
	why   string
	err   string
}

// rungTally counts the reads of one rung.
type rungTally struct {
	lists, reads, exact, one, two, errs int
}

// totals is the tally of each rung and of the whole run.
type totals struct {
	rungs map[mtgv1.SixtyStep]*rungTally
	all   rungTally
	// grid counts the reads of each label as each step, errors under
	// the unspecified step.
	grid map[mtgv1.SixtyStep]map[mtgv1.SixtyStep]int
}

func tally(results []result) totals {
	t := totals{rungs: map[mtgv1.SixtyStep]*rungTally{}, grid: map[mtgv1.SixtyStep]map[mtgv1.SixtyStep]int{}}
	for _, s := range steps {
		t.rungs[s] = &rungTally{}
		t.grid[s] = map[mtgv1.SixtyStep]int{}
	}
	for _, r := range results {
		rt := t.rungs[r.list.label]
		rt.lists++
		t.all.lists++
		for _, s := range r.steps {
			t.grid[r.list.label][s]++
			for _, c := range []*rungTally{rt, &t.all} {
				c.reads++
				switch d := distance(r.list.label, s); {
				case s == mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED:
					c.errs++
				case d == 0:
					c.exact++
				case d == 1:
					c.one++
				default:
					c.two++
				}
			}
		}
	}
	return t
}

// pass is the gate of D-868. A judge error is not a read of the label,
// so it counts against the rate.
func (t totals) pass() bool {
	if t.all.reads == 0 {
		return false
	}
	return float64(t.all.exact) >= passRate*float64(t.all.reads) && t.all.two == 0
}

var steps = []mtgv1.SixtyStep{
	mtgv1.SixtyStep_SIXTY_STEP_CASUAL, mtgv1.SixtyStep_SIXTY_STEP_FNM, mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
}

func distance(a, b mtgv1.SixtyStep) int {
	d := int(a) - int(b)
	if d < 0 {
		d = -d
	}
	return d
}

func writeSplit(w io.Writer, h header, set picked) {
	_, _ = fmt.Fprintf(w, "# The calibration split of the 60-card power judge (PR-71)\n\n")
	_, _ = fmt.Fprintf(w, "Card snapshot: %s. Precon table: %s. No provider call.\n\n", h.snapshot, h.precons)
	writeSkipped(w, set)
	_, _ = fmt.Fprintf(w, "| Rung | Split | Date | Format | List |\n|---|---|---|---|---|\n")
	for _, l := range set.lists {
		sp := "test"
		if l.dev {
			sp = "dev"
		}
		_, _ = fmt.Fprintf(w, "| %s | %s | %s | %s | %s |\n", stepWord(l.label), sp, l.date, l.format, cell(l.id+", "+l.name))
	}
}

func writeSkipped(w io.Writer, set picked) {
	for _, s := range steps {
		if n := set.skipped[s]; n > 0 {
			_, _ = fmt.Fprintf(w, "The %s rung skipped %d newer lists with a card that the snapshot does not hold.\n\n", stepWord(s), n)
		}
	}
}

func writeReport(w io.Writer, h header, set picked, results []result, t totals) {
	_, _ = fmt.Fprintf(w, "# The power judge of a 60-card import: the %s split (PR-71)\n\n", h.split)
	_, _ = fmt.Fprintf(w, "Date: %s. Card snapshot: %s. Precon table: %s.\n\n", time.Now().UTC().Format("2006-01-02"), h.snapshot, h.precons)
	_, _ = fmt.Fprintf(w, "Judge: %s, prompt version %d. Reads of each list: %d.\n\n", h.judge, generate.SixtyJudgeVersion, h.reads)
	cost := "unknown"
	if h.cost.CostUSD != nil {
		cost = fmt.Sprintf("$%.4f", *h.cost.CostUSD)
	}
	_, _ = fmt.Fprintf(w, "Cost: %s over %d calls, in %s.\n\n", cost, h.cost.Calls, h.took)
	if h.split == "test" {
		verdict := "FAIL"
		if t.pass() {
			verdict = "PASS"
		}
		_, _ = fmt.Fprintf(w, "Verdict: %s\n\n", verdict)
		_, _ = fmt.Fprintf(w, "The gate of D-868: %d of %d reads name the label, and %.0f percent is the floor. %d reads sit two steps off, and 0 is the cap.\n\n",
			t.all.exact, t.all.reads, passRate*100, t.all.two)
	}
	writeSkipped(w, set)
	_, _ = fmt.Fprintf(w, "| Rung | Lists | Reads | Label | One step off | Two steps off | Errors |\n|---|---|---|---|---|---|---|\n")
	for _, s := range steps {
		r := t.rungs[s]
		_, _ = fmt.Fprintf(w, "| %s | %d | %d | %d | %d | %d | %d |\n", stepWord(s), r.lists, r.reads, r.exact, r.one, r.two, r.errs)
	}
	_, _ = fmt.Fprintf(w, "| all | %d | %d | %d | %d | %d | %d |\n\n", t.all.lists, t.all.reads, t.all.exact, t.all.one, t.all.two, t.all.errs)
	_, _ = fmt.Fprintf(w, "| Label | Read casual | Read fnm | Read tournament | Error |\n|---|---|---|---|---|\n")
	for _, s := range steps {
		g := t.grid[s]
		_, _ = fmt.Fprintf(w, "| %s | %d | %d | %d | %d |\n", stepWord(s),
			g[mtgv1.SixtyStep_SIXTY_STEP_CASUAL], g[mtgv1.SixtyStep_SIXTY_STEP_FNM], g[mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT], g[mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED])
	}
	_, _ = fmt.Fprintf(w, "\n| Rung | Date | Format | List | Reads | Reason of the first read |\n|---|---|---|---|---|---|\n")
	for _, r := range results {
		words := make([]string, 0, len(r.steps))
		for _, s := range r.steps {
			words = append(words, stepWord(s))
		}
		_, _ = fmt.Fprintf(w, "| %s | %s | %s | %s | %s | %s |\n", stepWord(r.list.label), r.list.date, r.list.format,
			cell(r.list.id+", "+r.list.name), strings.Join(words, " "), cell(reason(r)))
	}
}

// reason is the reason of the first read, or the first error when no
// read answered.
func reason(r result) string {
	if r.why == "" && r.err != "" {
		return "error: " + r.err
	}
	return r.why
}

// cell keeps a table cell on one line and out of the column bars.
func cell(s string) string {
	s = strings.ReplaceAll(s, "|", "/")
	return strings.Join(strings.Fields(s), " ")
}
