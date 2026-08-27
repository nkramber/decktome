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
	list, err := cb.Build(idx, candidates.Request{
		Format:             mtgv1.FormatId_FORMAT_ID_COMMANDER,
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
		Format:       mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Plan:         fmt.Sprintf("a %s deck led by %s, at bracket 3", *theme, cmdr.GetName()),
		Pool:         pool,
		Commanders:   []string{cmdr.GetOracleId()},
		PoolRule:     mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Roles:        generate.Roles(list),
		Limits:       "Exactly 100 cards, the commander included. So list exactly 99 cards. One copy of each name, basic lands excepted. Every card must fit the commander's color identity.",
		Targets:      map[string]int{"land": 36, "ramp": 10, "draw": 10, "removal": 8, "wipe": 3, "threat": 12, "synergy": 20},
		LegalityAsOf: idx.AsOf.Format("2006-01-02"),
	}
	res, err := b.Build(context.Background(), req, acc)
	if err != nil {
		return err
	}
	report(pool, res, acc)
	return nil
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
