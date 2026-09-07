// Command bracket-gate runs the PR-14A bracket prompts through the
// generator and writes the gate document.
//
// The bars come from the roadmap (PR-14A). Every deck passes the block
// checks, every deck sits in band with no content violation after its
// repair turns, and the judge agrees with the bracket in eight of ten.
//
// CAUTION: this calls a real provider and it costs money. BRACKET_GATE=1
// is required, so it can not run by accident. `make bracket-gate` writes
// to BRACKET_GATE_OUT and refuses a file that already holds a verdict: a
// rerun must never overwrite a scored document (D-65). A -dry run calls
// no provider and needs no guard. The exit code is 1 on a FAIL verdict,
// and the document is written first.
//
// Usage:
//
//	BRACKET_GATE=1 CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go run ./cmd/bracket-gate
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/rules"
)

//go:embed prompts.json
var promptsJSON []byte

type prompt struct {
	ID        int    `json:"id"`
	Bracket   int32  `json:"bracket"`
	Theme     string `json:"theme"`
	Commander string `json:"commander"`
	Plan      string `json:"plan"`
}

type result struct {
	prompt       prompt
	deck         *mtgv1.Deck
	notes        []string
	repaired     bool
	repairReason string
	poolSize     int
	judged       *generate.BracketJudgement
	judgeErr     error
	err          error
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	only := flag.String("only", "", "run these prompt ids only, comma separated")
	dry := flag.Bool("dry", false, "build every shortlist and stop before the provider calls")
	noJudge := flag.Bool("no-judge", false, "skip the judge lane, which costs one judge call a deck")
	rejudge := flag.String("rejudge", "", "judge the decks of this gate document, and build nothing")
	runOut := flag.String("run-out", "", "write the run header and the rows as JSONL here (PR-15)")
	flag.Parse()
	if *rejudge != "" {
		return runRejudge(*rejudge, *runOut)
	}

	var file struct {
		Prompts []prompt `json:"prompts"`
	}
	if err := json.Unmarshal(promptsJSON, &file); err != nil {
		return fmt.Errorf("prompts: %w", err)
	}
	prompts, err := selectPrompts(file.Prompts, *only)
	if err != nil {
		return err
	}
	if !*dry {
		if err := gatekit.SpendGuard("BRACKET_GATE"); err != nil {
			return err
		}
	}
	// The run file is never overwritten (D-65), and the check runs before
	// the first provider call.
	if err := gatekit.RefuseExisting(*runOut); err != nil {
		return err
	}
	run := evalrun.New("bracket", evalrun.RunID(*runOut))
	run.Header.Only = *only
	run.Header.Prompts["generate"] = generate.PromptVersion
	run.LowerIsBetter("blocks", "off_band", "content_violations", "judge_error", "repaired")
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(context.Background(), quiet)
	if err != nil {
		return err
	}
	run.SetSnapshot(idx.AsOf)
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	rcfg, err := rules.Load()
	if err != nil {
		return err
	}
	var b *generate.Builder
	var acc *llm.Accumulator
	var client *llm.Client
	if !*dry {
		client, err = llm.NewFromEnv(gatekit.Env, quiet)
		if err != nil {
			return err
		}
		run.SetRoles(client.Config(), llm.RoleGenerate, llm.RoleRepair, llm.RoleJudge)
		prices, err := llm.LoadPrices()
		if err != nil {
			return err
		}
		acc = llm.NewAccumulator(prices)
		prof, err := gatekit.Profiler(idx, rcfg, quiet)
		if err != nil {
			return err
		}
		// The quality model grades every deck the gate builds, and the
		// summary names the tier (PR-14B). No stored model grades nothing.
		scorer, err := gatekit.Scorer(context.Background())
		if err != nil {
			return err
		}
		b = generate.NewBuilder(client, rcfg, idx, quiet, generate.WithProfiler(prof), generate.WithScorer(scorer))
	}

	start := time.Now()
	var results []result
	for _, p := range prompts {
		r := build(context.Background(), b, cb, idx, p, acc, *dry)
		if !*dry && !*noJudge && r.deck != nil {
			r.judged, r.judgeErr = generate.JudgeBracket(context.Background(), client, r.deck, idx, acc)
			if r.judgeErr != nil {
				fmt.Fprintf(os.Stderr, "  judge %d failed: %v\n", p.ID, r.judgeErr)
			}
		}
		results = append(results, r)
		fmt.Fprintf(os.Stderr, "  %2d. bracket %d %-28s pool %3d  %s\n", p.ID, p.Bracket, p.Commander, r.poolSize, status(r))
	}
	if *dry {
		fmt.Fprintf(os.Stderr, "\ndry run: %d shortlists built, no provider call ran\n", len(results))
		return nil
	}
	pass := report(os.Stdout, results, acc, idx, time.Since(start), run)
	if err := evalrun.WriteFile(*runOut, run); err != nil {
		return err
	}
	if !pass {
		return errGateFailed
	}
	return nil
}

// errGateFailed is the exit reason after a FAIL verdict. The document is
// written before it, so the record stays complete.
var errGateFailed = errors.New("bracket gate failed: read the verdict line of the document")

// selectPrompts keeps the prompts -only names, or all of them.
func selectPrompts(all []prompt, only string) ([]prompt, error) {
	ids, err := gatekit.ParseIDs(only)
	if err != nil {
		return nil, fmt.Errorf("bracket-gate: %w", err)
	}
	if ids == nil {
		return all, nil
	}
	want := map[int]bool{}
	for _, id := range ids {
		want[id] = true
	}
	var kept []prompt
	for _, p := range all {
		if want[p.ID] {
			kept = append(kept, p)
		}
	}
	if len(kept) == 0 {
		return nil, fmt.Errorf("-only %q matches no prompt", only)
	}
	return kept, nil
}

func status(r result) string {
	switch {
	case r.err != nil:
		return "ERROR: " + r.err.Error()
	case r.deck == nil:
		return "no deck"
	default:
		judged := "unjudged"
		if r.judged != nil {
			judged = fmt.Sprintf("judged %d", r.judged.Bracket)
		}
		return fmt.Sprintf("%d cards, %d blocks, %d off band, %d content, %s, repaired %v",
			gatekit.CountCards(r.deck), len(gatekit.BlockFindings(r.deck)), len(offBand(r.deck)), len(contentFindings(r.deck)), judged, r.repaired)
	}
}

// build runs one prompt: the commander by name, the shortlist in its
// identity, and the build at the bracket. Every prompt is any-card with
// no collection, so the profile reads the bracket and nothing else.
func build(ctx context.Context, b *generate.Builder, cb *candidates.Builder, idx *cards.Index,
	p prompt, acc *llm.Accumulator, dry bool) result {
	out := result{prompt: p}
	format := mtgv1.FormatId_FORMAT_ID_COMMANDER
	c, ok := idx.ByName(p.Commander)
	if !ok {
		out.err = fmt.Errorf("no card named %q", p.Commander)
		return out
	}
	commanderIDs := []string{c.GetOracleId()}
	colors := c.GetColorIdentity()
	list, err := cb.Build(idx, candidates.Request{
		Format:             format,
		Colors:             colors,
		Theme:              p.Theme,
		CommanderOracleIDs: commanderIDs,
		PoolRule:           mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Bracket:            p.Bracket,
	})
	if err != nil {
		out.err = fmt.Errorf("candidates: %w", err)
		return out
	}
	always := append([]*mtgv1.Card{c}, generate.BasicLands(idx.ByName, colors)...)
	pool := generate.FromList(list, always, true)
	out.poolSize = pool.Size()
	if dry {
		return out
	}
	power := gatekit.PowerLevel(p.Bracket, "")
	res, err := b.Build(ctx, generate.Request{
		SessionID:    fmt.Sprintf("bracket-gate-%d", p.ID),
		Format:       format,
		Power:        power,
		Plan:         p.Plan + fmt.Sprintf("\nTheme: %s\nCommander bracket: %d", p.Theme, p.Bracket),
		Pool:         pool,
		Commanders:   commanderIDs,
		PoolRule:     mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Roles:        generate.Roles(list),
		Targets:      generate.TargetsFor(format, power),
		Limits:       generate.LimitsFor(format),
		LegalityAsOf: idx.AsOf.Format("2006-01-02"),
	}, acc)
	if err != nil {
		out.err = err
		return out
	}
	out.deck, out.notes, out.repaired, out.repairReason = res.Deck, res.Notes, res.Repaired, res.RepairReason
	return out
}
