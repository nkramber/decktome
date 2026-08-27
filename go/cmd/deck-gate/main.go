// Command deck-gate runs the PR-8 golden prompts through the generator
// and writes the gate document.
//
// The bars come from the roadmap. Every returned deck passes the block
// checks, and no invented name reaches the user.
//
// CAUTION: this calls a real provider and it costs money. Ask the owner
// before every run, and write to a new DECK_GATE_OUT: a rerun must never
// overwrite a scored document (D-65).
//
// Usage:
//
//	CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go run ./cmd/deck-gate -collection internal/collections/testdata/manabox_collection.csv
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/collections"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

//go:embed prompts.json
var promptsJSON []byte

type prompt struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Format     string   `json:"format"`
	Theme      string   `json:"theme"`
	Colors     []string `json:"colors"`
	Commander  string   `json:"commander"`
	Bracket    int32    `json:"bracket"`
	Power      string   `json:"power"`
	Pool       string   `json:"pool"`
	Collection bool     `json:"collection"`
	Locked     []string `json:"locked"`
	Budget     float64  `json:"budget"`
	Plan       string   `json:"plan"`
}

type result struct {
	prompt   prompt
	deck     *mtgv1.Deck
	notes    []string
	repaired bool
	poolSize int
	judged   *generate.Judgement
	err      error
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	collPath := flag.String("collection", "", "a ManaBox CSV for the owned modes")
	only := flag.String("only", "", "run these prompt ids only, comma separated")
	dry := flag.Bool("dry", false, "build every shortlist and stop before the provider calls")
	noJudge := flag.Bool("no-judge", false, "skip the F-26 judge lane, which costs about $0.0034 a deck")
	flag.Parse()

	var file struct {
		Prompts []prompt `json:"prompts"`
	}
	if err := json.Unmarshal(promptsJSON, &file); err != nil {
		return fmt.Errorf("prompts: %w", err)
	}
	if *only != "" {
		want := map[int]bool{}
		for _, s := range strings.Split(*only, ",") {
			var id int
			if _, err := fmt.Sscanf(strings.TrimSpace(s), "%d", &id); err != nil {
				return fmt.Errorf("-only takes prompt ids: %w", err)
			}
			want[id] = true
		}
		var kept []prompt
		for _, p := range file.Prompts {
			if want[p.ID] {
				kept = append(kept, p)
			}
		}
		file.Prompts = kept
	}

	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		return fmt.Errorf("set CARDS_SNAPSHOT_DIR")
	}
	idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, quiet)
	if err != nil {
		return fmt.Errorf("cards: %w", err)
	}
	owned := map[string]int32{}
	if *collPath != "" {
		owned, err = loadOwned(*collPath, idx)
		if err != nil {
			return err
		}
	}
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	rcfg, err := rules.Load()
	if err != nil {
		return err
	}
	// A dry run builds every shortlist and calls no provider, so it needs
	// no API key. It is the free check before a paid run.
	var b *generate.Builder
	var acc *llm.Accumulator
	var client *llm.Client
	if !*dry {
		env := func(k string) string {
			if k == llm.EnvRequireKeys {
				return "1"
			}
			return os.Getenv(k)
		}
		client, err = llm.NewFromEnv(env, quiet)
		if err != nil {
			return err
		}
		prices, err := llm.LoadPrices()
		if err != nil {
			return err
		}
		acc = llm.NewAccumulator(prices)
		b = generate.NewBuilder(client, rcfg, idx, quiet)
	}

	start := time.Now()
	var results []result
	for _, p := range file.Prompts {
		r := build(context.Background(), b, cb, idx, owned, p, acc, *dry)
		// The judge lane is the real check for F-26, and the deterministic
		// net can not read the truth of a rules claim (D-229).
		if !*dry && !*noJudge && r.deck != nil && r.deck.GetSummary() != "" {
			j, err := generate.JudgeSummary(context.Background(), client, p.Name, r.deck.GetSummary(), acc)
			if err != nil {
				fmt.Fprintf(os.Stderr, "  judge %d failed: %v\n", p.ID, err)
			} else {
				r.judged = j
			}
		}
		results = append(results, r)
		fmt.Fprintf(os.Stderr, "  %2d. %-38s pool %3d  %s\n", p.ID, p.Name, r.poolSize, status(r))
	}
	if *dry {
		fmt.Fprintf(os.Stderr, "\ndry run: %d shortlists built, no provider call ran\n", len(results))
		return nil
	}
	report(os.Stdout, results, acc, idx, time.Since(start))
	return nil
}

func status(r result) string {
	switch {
	case r.err != nil:
		return "ERROR: " + r.err.Error()
	case r.deck == nil:
		return "no deck"
	default:
		return fmt.Sprintf("%d cards, %d blocks, %d notes, repaired %v",
			countCards(r.deck), len(blocks(r.deck)), len(r.notes), r.repaired)
	}
}

func build(ctx context.Context, b *generate.Builder, cb *candidates.Builder, idx *cards.Index,
	owned map[string]int32, p prompt, acc *llm.Accumulator, dry bool) result {
	out := result{prompt: p}
	format := formatID(p.Format)
	if format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		out.err = fmt.Errorf("unknown format %q", p.Format)
		return out
	}
	var commanders []*mtgv1.Card
	var commanderIDs []string
	if p.Commander != "" {
		c, ok := idx.ByName(p.Commander)
		if !ok {
			out.err = fmt.Errorf("no card named %q", p.Commander)
			return out
		}
		commanders = append(commanders, c)
		commanderIDs = append(commanderIDs, c.GetOracleId())
	}
	colors := colorList(p.Colors)
	if len(commanders) > 0 {
		colors = commanders[0].GetColorIdentity()
	}
	poolRule := poolRuleID(p.Pool)
	own := map[string]int32{}
	if p.Collection {
		own = owned
	}
	list, err := cb.Build(idx, candidates.Request{
		Format:             format,
		Colors:             colors,
		Theme:              p.Theme,
		CommanderOracleIDs: commanderIDs,
		PoolRule:           poolRule,
		Owned:              own,
		Bracket:            p.Bracket,
	})
	if err != nil {
		out.err = fmt.Errorf("candidates: %w", err)
		return out
	}
	// A locked card must be nameable, or the deck can not hold it (D-70).
	for _, name := range p.Locked {
		if c, ok := idx.ByName(name); ok {
			commanders = append(commanders, c)
		}
	}
	buyList := poolRule == mtgv1.PoolRule_POOL_RULE_ANY_CARD || p.Budget > 0
	// The shortlist leaves basic lands out on purpose (D-225).
	always := append([]*mtgv1.Card(nil), commanders...)
	always = append(always, generate.BasicLands(idx.ByName, colors)...)
	pool := generate.FromList(list, always, buyList)
	out.poolSize = pool.Size()
	if dry {
		return out
	}
	plan := p.Plan
	for _, name := range p.Locked {
		plan += fmt.Sprintf("\n%s must be in the deck.", name)
	}
	res, err := b.Build(ctx, generate.Request{
		SessionID:    fmt.Sprintf("gate-%d", p.ID),
		Format:       format,
		Power:        power(p),
		Plan:         plan,
		Pool:         pool,
		Commanders:   commanderIDs,
		PoolRule:     poolRule,
		OracleCounts: own,
		Roles:        generate.Roles(list),
		Targets:      generate.TargetsFor(format, power(p)),
		Limits:       generate.LimitsFor(format),
		LegalityAsOf: idx.AsOf.Format("2006-01-02"),
	}, acc)
	if err != nil {
		out.err = err
		return out
	}
	out.deck, out.notes, out.repaired = res.Deck, res.Notes, res.Repaired
	return out
}

func loadOwned(path string, idx *cards.Index) (map[string]int32, error) {
	f, err := os.Open(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	rows, _, err := collections.ParseManaBoxCSV(f)
	if err != nil {
		return nil, fmt.Errorf("collection: %w", err)
	}
	entries, _ := collections.Resolve(rows, idx)
	return collections.OracleCounts(entries), nil
}
