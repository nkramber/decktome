// Command generate-probe builds one deck with the real generate role and
// reports what the model did with the shortlist.
//
// It answers one question the fake provider can not: does the model copy
// card names exactly from the shortlist (F-13)? It writes no document and
// it changes no state.
//
// CAUTION: this calls a real provider and it costs money. One Commander
// deck is a few cents.
//
// Usage:
//
//	CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go run ./cmd/generate-probe -theme lifegain -commander "Karlov of the Ghost Council"
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
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

	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		return fmt.Errorf("set CARDS_SNAPSHOT_DIR")
	}
	idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, quiet)
	if err != nil {
		return fmt.Errorf("cards: %w", err)
	}
	cmdr, ok := idx.ByName(*cmdrName)
	if !ok {
		return fmt.Errorf("no card named %q in the snapshot", *cmdrName)
	}
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	fid := mtgv1.FormatId_FORMAT_ID_COMMANDER
	switch *format {
	case "standard":
		fid = mtgv1.FormatId_FORMAT_ID_STANDARD
	case "modern":
		fid = mtgv1.FormatId_FORMAT_ID_MODERN
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
	always := append([]*mtgv1.Card{cmdr}, generate.BasicLands(idx.ByName, cmdr.GetColorIdentity())...)
	pool := generate.FromList(list, always, false)
	fmt.Printf("snapshot: %d cards, %d paper-printing swaps (D-221)\n", idx.Len(), idx.PaperSwaps())
	if *dry {
		fmt.Printf("shortlist: %d names. No provider call ran.\n", pool.Size())
		return nil
	}

	env := func(k string) string {
		if k == llm.EnvRequireKeys {
			return "1"
		}
		return os.Getenv(k)
	}
	client, err := llm.NewFromEnv(env, quiet)
	if err != nil {
		return fmt.Errorf("llm client: %w", err)
	}
	rcfg, err := rules.Load()
	if err != nil {
		return err
	}
	b := generate.NewBuilder(client, rcfg, idx, quiet)
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
		Commanders:   []string{cmdr.GetOracleId()},
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
	cost := 0.0
	if after.CostUSD != nil {
		cost = *after.CostUSD
		if before.CostUSD != nil {
			cost -= *before.CostUSD
		}
	}
	share := 0.0
	if in > 0 {
		share = 100 * float64(cached) / float64(in)
	}
	fmt.Printf("call %d: input %6d (cached %6d, %5.1f%%)  output %6d  cost $%.5f\n",
		n, in, cached, share, out, cost)
}

// reportCache states what the cache saved over the run, against the price
// table. The roadmap says a cache read costs a tenth of a fresh read, and
// that is true of the input alone.
func reportCache(acc *llm.Accumulator, calls int) {
	rep := acc.Report()
	t := rep.Tokens
	cost := 0.0
	if rep.CostUSD != nil {
		cost = *rep.CostUSD
	}
	fmt.Printf("\ntotals over %d calls: input %d, cached %d, output %d, reasoning %d\n",
		calls, t.InputTokens, t.CachedInputTokens, t.OutputTokens, t.ReasoningTokens)
	fmt.Printf("reported cost: $%.5f\n", cost)
}

func report(pool *generate.Pool, res *generate.Result, acc *llm.Accumulator) {
	d := res.Deck
	total := 0
	for _, c := range d.GetCards() {
		total += int(c.GetCount())
	}
	fmt.Printf("shortlist: %d names\n", pool.Size())
	fmt.Printf("repair turn ran: %v\n", res.Repaired)
	fmt.Printf("cards listed: %d entries, %d with counts\n", len(d.GetCards()), total)
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
