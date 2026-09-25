// Command feedback-triage reads a harvest and turns each thumbs down
// into a test case (PR-28b, D-643).
//
// The reason keys of the dialog name the class with no model call, and
// the judge role reads only the verdicts the keys can not place. So -dry
// is free and deterministic, and a live run spends a few cents on each
// verdict the keys could not place. -dry still writes its cases under
// -apply, the way a dry deck gate still writes its shortlists (D-521).
//
// -apply writes each case into the file that owns it, and each owner
// question into docs/owner-questions.md. The owner reads the change on
// the pull request, which is the accept step (D-642).
//
// -manifest names the file that hands those cases to the fix cycle of
// PR-28c: the gate that owns each one, its id, and the reader's own
// words (D-645). It needs no -apply, so a plan reads the same list the
// live cycle acts on.
//
// CAUTION: a live run calls a real provider and it costs money.
// FEEDBACK_TRIAGE=1 is required, so it can not run by accident. A -dry
// run calls no provider and needs no guard.
//
// Usage:
//
//	go run ./cmd/feedback-triage -root .. -dry
//	FEEDBACK_TRIAGE=1 go run ./cmd/feedback-triage -root .. \
//	  -out ../docs/reference/pr28b-triage-2026-09-09.md -apply
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/harvest"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/triage"
)

// ownerQuestions is the decision queue a class that meets an owner
// decision writes to (D-558).
const ownerQuestions = "docs/owner-questions.md"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	root := flag.String("root", "..", "the repo root the triage reads and writes under")
	in := flag.String("in", "", "a harvest JSONL file. Empty reads every harvest no live -apply triage read")
	out := flag.String("out", "", "write the triage document here")
	dry := flag.Bool("dry", false, "call no model. A verdict the reason keys can not place keeps its need and writes no case")
	apply := flag.Bool("apply", false, "write each case into the file that owns it")
	manifest := flag.String("manifest", "", "write the case manifest here, for the fix cycle of PR-28c")
	flag.Parse()

	if !*dry {
		if err := gatekit.SpendGuard("FEEDBACK_TRIAGE"); err != nil {
			return err
		}
	}
	// The document is never overwritten (D-65), and the check runs before
	// the first provider call.
	if err := gatekit.RefuseExisting(*out); err != nil {
		return err
	}

	paths := []string{*in}
	if *in == "" {
		pending, err := triage.PendingHarvests(*root)
		if err != nil {
			return err
		}
		if len(pending) == 0 {
			return fmt.Errorf("no harvest under %s that a triage did not apply. Run make feedback-harvest first", filepath.Join(*root, harvest.Dir))
		}
		paths = pending
	}
	// from names the harvest of each record, so a live run marks a
	// harvest read only when each of its verdicts applied (REV-082).
	var recs []harvest.Record
	var from []string
	for _, path := range paths {
		part, err := triage.ReadHarvest(path)
		if err != nil {
			return err
		}
		// A verdict that an earlier live run applied writes no case twice.
		fresh, err := triage.Unapplied(*root, part)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "harvest      %s, %d verdict(s), %d not applied\n", path, len(part), len(fresh))
		recs = append(recs, fresh...)
		for range fresh {
			from = append(from, path)
		}
	}

	quiet := gatekit.Quiet()
	// The card index names a commander and a card. It is free and local,
	// and a triage without it still runs: every case then names the gap.
	var namer triage.Namer
	if idx, err := gatekit.LoadSnapshot(context.Background(), quiet); err == nil {
		namer = func(id string) (string, bool) {
			c, ok := idx.ByOracleID(id)
			if !ok {
				return "", false
			}
			return c.GetName(), true
		}
	} else {
		fmt.Fprintf(os.Stderr, "cards        none: %v. Every case names its gap.\n", err)
	}

	var judger triage.Judger
	var acc *llm.Accumulator
	if !*dry {
		client, err := llm.NewFromEnv(gatekit.Env, quiet)
		if err != nil {
			return err
		}
		prices, err := llm.LoadPrices()
		if err != nil {
			return err
		}
		acc = llm.NewAccumulator(prices)
		judger = triage.LiveJudger(client, acc)
	}

	nextID := func(target string) (int, error) {
		if triage.IsFixtureTarget(target) {
			return triage.NextFixtureID(filepath.Join(*root, target))
		}
		return triage.NextID(filepath.Join(*root, target))
	}
	results := triage.Run(context.Background(), recs, judger, namer, nextID)

	if *apply {
		applyErr := applyCases(*root, results)
		// A dry run writes no case for a verdict that needs the judge, so
		// only a live run records what it applied. A failed verdict or a
		// failed write keeps its verdict pending, and the writes that landed
		// stay on record (REV-082).
		if !*dry {
			if err := triage.Settle(*root, recs, from, results); err != nil {
				return errors.Join(applyErr, err)
			}
		}
		if applyErr != nil {
			return applyErr
		}
	}
	// The manifest names the cases a gate measures, whether or not this
	// run wrote them into a gate file. So the plan of a dry cycle reads
	// the same list the live one acts on.
	if *manifest != "" {
		m := triage.ManifestOf(strings.Join(paths, ", "), time.Now().UTC(), results)
		if err := triage.WriteManifest(*manifest, m); err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "manifest     %s, %d case(s) a gate measures\n", *manifest, len(m.Cases))
	}

	cost := "nothing, no model ran"
	if acc != nil {
		cost = gatekit.CostWord(acc.Report())
	}
	var w = os.Stdout
	if *out != "" {
		f, err := os.Create(*out) // #nosec G304 -- the operator names the file.
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()
		w = f
	}
	pass := triage.Report(w, results, time.Now().UTC(), *dry, cost)
	if *out != "" {
		fmt.Fprintf(os.Stderr, "wrote %s\n", *out)
	}
	fmt.Fprintf(os.Stderr, "kinds        %s\n", triage.Kinds(results))
	if !pass {
		return fmt.Errorf("the triage verdict is FAIL. Read the document")
	}
	return nil
}

// applyCases writes every case into the file that owns it, and every
// owner question into the decision queue.
func applyCases(root string, results []triage.Result) error {
	for i, r := range results {
		if !r.NeedsWrite() || r.Route.Owner {
			continue
		}
		path := filepath.Join(root, r.Case.Target)
		// A parser fixture is a file of its own in a folder (D-888).
		if triage.IsFixtureTarget(r.Case.Target) {
			written, err := triage.WriteFixture(path, r.Case.ID, r.Case.Body)
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(root, written)
			if err != nil {
				return err
			}
			results[i].Applied = filepath.ToSlash(rel)
			continue
		}
		if err := triage.Append(path, r.Case.Body); err != nil {
			return err
		}
		results[i].Applied = r.Case.Target
	}
	owners := triage.Owners(results)
	if len(owners) == 0 {
		return nil
	}
	path := filepath.Join(root, ownerQuestions)
	next, err := triage.NextOQ(path)
	if err != nil {
		return err
	}
	// Each question marks its own result as it lands, so a failed write
	// leaves the earlier ones on record (REV-082).
	for i, r := range results {
		if r.Err != nil || !r.Route.Owner {
			continue
		}
		question, whyYou, blocks := triage.OwnerRow(r.Route)
		number := fmt.Sprintf("OQ-%d", next)
		if err := triage.AppendOwnerQuestion(path, number, question, whyYou, blocks); err != nil {
			return err
		}
		results[i].Applied = ownerQuestions
		next++
	}
	return nil
}
