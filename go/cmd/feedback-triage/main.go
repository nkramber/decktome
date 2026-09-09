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
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
	in := flag.String("in", "", "a harvest JSONL file. Empty reads the newest one")
	out := flag.String("out", "", "write the triage document here")
	dry := flag.Bool("dry", false, "call no model. A verdict the reason keys can not place keeps its need and writes no case")
	apply := flag.Bool("apply", false, "write each case into the file that owns it")
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

	path := *in
	if path == "" {
		p, err := triage.NewestHarvest(*root)
		if err != nil {
			return err
		}
		if p == "" {
			return fmt.Errorf("no harvest under %s. Run make feedback-harvest first", filepath.Join(*root, harvest.Dir))
		}
		path = p
	}
	recs, err := triage.ReadHarvest(path)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "harvest      %s, %d verdict(s)\n", path, len(recs))

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

	nextID := func(target string) (int, error) { return triage.NextID(filepath.Join(*root, target)) }
	results := triage.Run(context.Background(), recs, judger, namer, nextID)

	if *apply {
		if err := applyCases(*root, results); err != nil {
			return err
		}
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
		if r.Err != nil || len(r.Case.Body) == 0 || r.Case.Target == "" {
			continue
		}
		path := filepath.Join(root, r.Case.Target)
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
	for _, o := range owners {
		question, whyYou, blocks := triage.OwnerRow(o)
		number := fmt.Sprintf("OQ-%d", next)
		if err := triage.AppendOwnerQuestion(path, number, question, whyYou, blocks); err != nil {
			return err
		}
		next++
	}
	for i, r := range results {
		if r.Err == nil && r.Route.Owner {
			results[i].Applied = ownerQuestions
		}
	}
	return nil
}
