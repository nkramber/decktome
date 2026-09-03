// Command generate-probe builds one deck with the real generate role and
// reports what the model did with the shortlist.
//
// It answers one question the fake provider can not: does the model copy
// card names exactly from the shortlist (F-13)? It writes no document and
// it changes no state.
//
// CAUTION: this calls a real provider and it costs money. One Commander
// deck is a few cents. GENERATE_PROBE=1 is required, so it can not run
// by accident. A -dry run calls no provider and needs no guard.
//
// Usage:
//
//	GENERATE_PROBE=1 CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go run ./cmd/generate-probe -theme lifegain -commander "Karlov of the Ghost Council"
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	theme := flag.String("theme", "lifegain", "the deck theme")
	cmdrName := flag.String("commander", "Karlov of the Ghost Council", "the commander")
	limit := flag.Int("limit", 300, "how many shortlist cards the model may name")
	dry := flag.Bool("dry", false, "build the shortlist and stop before the provider call")
	repeat := flag.Int("repeat", 1, "how many times to build, on one session id, to measure the prompt cache")
	format := flag.String("format", "commander", "commander, standard, or modern")
	flag.Parse()

	if !*dry {
		if err := gatekit.SpendGuard("GENERATE_PROBE"); err != nil {
			return err
		}
	}
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(context.Background(), quiet)
	if err != nil {
		return err
	}
	cmdr, ok := idx.ByName(*cmdrName)
	if !ok {
		return fmt.Errorf("no card named %q in the snapshot", *cmdrName)
	}
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	fid := gatekit.FormatID(*format)
	if fid == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return fmt.Errorf("unknown format %q: give commander, standard, or modern", *format)
	}
	list, err := cb.Build(idx, candidates.Request{
		Format:             fid,
		Colors:             cmdr.GetColorIdentity(),
		Theme:              *theme,
		CommanderOracleIDs: []string{cmdr.GetOracleId()},
		PoolRule:           mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Bracket:            3,
		Limits:             candidates.Limits{Total: *limit},
	})
	if err != nil {
		return fmt.Errorf("candidates: %w", err)
	}
	// Only Commander has a command zone (D-233).
	var cmdrIDs []string
	always := generate.BasicLands(idx.ByName, cmdr.GetColorIdentity())
	if fid == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		cmdrIDs = []string{cmdr.GetOracleId()}
		always = append([]*mtgv1.Card{cmdr}, always...)
	}
	pool := generate.FromList(list, always, false)
	fmt.Printf("snapshot: %d cards, %d paper-printing swaps (D-221)\n", idx.Len(), idx.PaperSwaps())
	if *dry {
		fmt.Printf("shortlist: %d names. No provider call ran.\n", pool.Size())
		return nil
	}

	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return fmt.Errorf("llm client: %w", err)
	}
	rcfg, err := rules.Load()
	if err != nil {
		return err
	}
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
	b := generate.NewBuilder(client, rcfg, idx, quiet, generate.WithProfiler(prof), generate.WithScorer(scorer))
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)
	req := generate.Request{
		SessionID:    "probe-1",
		Format:       fid,
		Plan:         fmt.Sprintf("a %s deck led by %s, at bracket 3", *theme, cmdr.GetName()),
		Pool:         pool,
		Commanders:   cmdrIDs,
		PoolRule:     mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Roles:        generate.Roles(list),
		Limits:       generate.LimitsFor(fid),
		Targets:      generate.TargetsFor(fid, nil),
		LegalityAsOf: idx.AsOf.Format("2006-01-02"),
	}
	// The cache lever of the roadmap: one key per session, the stable text
	// first and the session text last. A repeat run reads the same prefix,
	// so call 1 writes the cache and the rest read it.
	for i := 1; i <= *repeat; i++ {
		before := acc.Report()
		res, err := b.Build(context.Background(), req, acc)
		if err != nil {
			return err
		}
		after := acc.Report()
		if *repeat > 1 {
			reportCall(i, before, after)
			continue
		}
		report(pool, res, acc)
	}
	if *repeat > 1 {
		reportCache(acc, *repeat)
	}
	return nil
}

// reportCall prints what one call of a repeat run cost, and how much of
// its input the provider served from the cache.
func reportCall(n int, before, after llm.Report) {
	// Tokens is nil until an attempt reports usage, so the first call
	// compares against an empty report and not a nil pointer.
	var b, a llm.Usage
	if before.Tokens != nil {
		b = *before.Tokens
	}
	if after.Tokens != nil {
		a = *after.Tokens
	}
	in := a.InputTokens - b.InputTokens
	cached := a.CachedInputTokens - b.CachedInputTokens
	out := a.OutputTokens - b.OutputTokens
	// The delta of two priced reports. A nil on either side is unpriced,
	// never $0 (M-1).
	cost := "unpriced"
	if after.CostUSD != nil {
		c := *after.CostUSD
		if before.CostUSD != nil {
			c -= *before.CostUSD
		}
		cost = fmt.Sprintf("$%.5f", c)
	}
	share := 0.0
	if in > 0 {
		share = 100 * float64(cached) / float64(in)
	}
	fmt.Printf("call %d: input %6d (cached %6d, %5.1f%%)  output %6d  cost %s\n",
		n, in, cached, share, out, cost)
}

// reportCache states what the cache saved over the run, against the price
// table. The roadmap says a cache read costs a tenth of a fresh read, and
// that is true of the input alone.
func reportCache(acc *llm.Accumulator, calls int) {
	rep := acc.Report()
	var t llm.Usage
	if rep.Tokens != nil {
		t = *rep.Tokens
	}
	fmt.Printf("\ntotals over %d calls: input %d, cached %d, output %d, reasoning %d\n",
		calls, t.InputTokens, t.CachedInputTokens, t.OutputTokens, t.ReasoningTokens)
	fmt.Printf("reported cost: %s\n", gatekit.CostWord(rep))
}

func report(pool *generate.Pool, res *generate.Result, acc *llm.Accumulator) {
	d := res.Deck
	fmt.Printf("shortlist: %d names\n", pool.Size())
	fmt.Printf("repair turn ran: %v\n", res.Repaired)
	fmt.Printf("cards listed: %d entries, %d with counts. Sideboard: %d cards.\n",
		len(d.GetCards()), gatekit.CountCards(d), gatekit.CountSideboard(d))
	fmt.Printf("notes (names that missed twice): %d\n", len(res.Notes))
	for _, n := range res.Notes {
		fmt.Printf("  - %s\n", n)
	}
	fmt.Printf("\nsummary: %s\n", d.GetSummary())
	if claims := generate.LintSummary(d.GetSummary()); len(claims) > 0 {
		fmt.Printf("summary rules claims (F-26): %v\n", claims)
	}
	fmt.Printf("\nfindings:\n")
	for _, f := range d.GetValidation().GetFindings() {
		fmt.Printf("  [%s] %s: %s\n", strings.TrimPrefix(f.GetSeverity().String(), "SEVERITY_"), f.GetCode(), f.GetMessage())
	}
	if acc != nil {
		raw, _ := json.Marshal(acc.Report())
		fmt.Printf("\nusage: %s\n", raw)
	}
}
