// Command deck-gate runs the PR-8 golden prompts through the generator
// and writes the gate document.
//
// The bars come from the roadmap. Every returned deck passes the block
// checks, and no invented name reaches the user.
//
// CAUTION: this calls a real provider and it costs money. DECK_GATE=1 is
// required, so it can not run by accident. Ask the owner before every
// run. `make deck-gate` writes to DECK_GATE_OUT and refuses a file that
// already holds a verdict: a rerun must never overwrite a scored document
// (D-65). A -dry run calls no provider and needs no guard.
//
// Usage:
//
//	DECK_GATE=1 CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go run ./cmd/deck-gate -collection internal/collections/testdata/manabox_collection.csv
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/precons"
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
	// Precon names a preconstructed deck the build must keep a share of
	// (D-218, D-247).
	Precon string  `json:"precon"`
	Budget float64 `json:"budget"`
	Plan   string  `json:"plan"`
}

type result struct {
	prompt       prompt
	deck         *mtgv1.Deck
	notes        []string
	repaired     bool
	repairReason string
	poolSize     int
	judged       *generate.Judgement
	// judgeErr is the judge lane's failure. A deck with one has no
	// verdict on F-26, so it can not count as a pass on that bar (T-17).
	judgeErr error
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
			id, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
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

	// A dry run calls no provider, so it needs no guard.
	if !*dry {
		if err := gatekit.SpendGuard("DECK_GATE"); err != nil {
			return err
		}
	}
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(context.Background(), quiet)
	if err != nil {
		return err
	}
	owned := map[string]int32{}
	if *collPath != "" {
		owned, _, err = gatekit.LoadOwned(*collPath, idx)
		if err != nil {
			return err
		}
	}
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	preconSet, err := precons.Load(idx)
	if err != nil {
		return fmt.Errorf("precons: %w", err)
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
		client, err = llm.NewFromEnv(gatekit.Env, quiet)
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
		r := build(context.Background(), b, cb, idx, owned, p, acc, *dry, preconSet)
		// The judge lane is the real check for F-26, and the deterministic
		// net can not read the truth of a rules claim (D-229).
		if !*dry && !*noJudge && r.deck != nil && r.deck.GetSummary() != "" {
			j, err := generate.JudgeSummary(context.Background(), client, p.Name, r.deck.GetSummary(), acc)
			if err != nil {
				r.judgeErr = err
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
	owned map[string]int32, p prompt, acc *llm.Accumulator, dry bool, preconSet *precons.Set) result {
	out := result{prompt: p}
	format := gatekit.FormatID(p.Format)
	if format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		out.err = fmt.Errorf("unknown format %q", p.Format)
		return out
	}
	var commanders []*mtgv1.Card
	var commanderIDs []string
	// A prompt with no commander delegates the pick, as a user who says
	// "you pick" does. The generator must choose one (D-232).
	if p.Commander == "" && format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		pool, err := cb.CommanderPool(idx, candidates.Request{
			Format: format, Theme: p.Theme, Colors: gatekit.Colors(p.Colors),
			PoolRule: gatekit.PoolRuleID(p.Pool), Owned: ownedFor(p, owned), Bracket: p.Bracket,
		})
		if err != nil {
			out.err = fmt.Errorf("commander pool: %w", err)
			return out
		}
		if len(pool) == 0 {
			out.err = fmt.Errorf("the library holds no commander for %q", p.Theme)
			return out
		}
		commanders = append(commanders, pool[0].Card)
		commanderIDs = append(commanderIDs, pool[0].Card.GetOracleId())
	}
	if p.Commander != "" {
		c, ok := idx.ByName(p.Commander)
		if !ok {
			out.err = fmt.Errorf("no card named %q", p.Commander)
			return out
		}
		commanders = append(commanders, c)
		commanderIDs = append(commanderIDs, c.GetOracleId())
	}
	colors := gatekit.Colors(p.Colors)
	if len(commanders) > 0 {
		colors = commanders[0].GetColorIdentity()
	}
	poolRule := gatekit.PoolRuleID(p.Pool)
	own := ownedFor(p, owned)
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
	// The build reads the ids and states the cards in its own prompt, so
	// the gate adds nothing to the plan text by hand (D-242).
	var lockedIDs []string
	for _, name := range p.Locked {
		c, ok := idx.ByName(name)
		if !ok {
			out.err = fmt.Errorf("no card named %q to lock", name)
			return out
		}
		commanders = append(commanders, c)
		lockedIDs = append(lockedIDs, c.GetOracleId())
	}
	buyList := poolRule == mtgv1.PoolRule_POOL_RULE_ANY_CARD || p.Budget > 0
	// The shortlist leaves basic lands out on purpose (D-225).
	always := append([]*mtgv1.Card(nil), commanders...)
	always = append(always, generate.BasicLands(idx.ByName, colors)...)
	// The share rule of D-218 measures the precon's own cards, so every
	// one must be nameable. The shortlist ranks by theme and holds only
	// some of them. This must run before the pool is built: it did not,
	// and 27 of 93 cards reached the model, which then refused an
	// impossible instruction (D-248).
	var preconName string
	var preconIDs []string
	var preconLands int
	if p.Precon != "" {
		pc, ok := preconSet.Get(p.Precon)
		if !ok {
			out.err = fmt.Errorf("no precon named %q", p.Precon)
			return out
		}
		preconName, preconIDs, preconLands = pc.Name, pc.OracleIDs, pc.Lands
		for _, id := range preconIDs {
			if c, ok := idx.ByOracleID(id); ok {
				always = append(always, c)
			}
		}
	}
	pool := generate.FromList(list, always, buyList)
	out.poolSize = pool.Size()
	if dry {
		return out
	}
	plan := p.Plan
	res, err := b.Build(ctx, generate.Request{
		SessionID:       fmt.Sprintf("gate-%d", p.ID),
		Format:          format,
		Power:           power(p),
		Plan:            plan,
		Pool:            pool,
		Commanders:      commanderIDs,
		Locked:          lockedIDs,
		Precon:          preconName,
		PreconOracleIDs: preconIDs,
		PreconLands:     preconLands,
		PoolRule:        poolRule,
		OracleCounts:    own,
		Roles:           generate.Roles(list),
		Targets:         generate.TargetsFor(format, power(p)),
		Limits:          generate.LimitsFor(format),
		LegalityAsOf:    idx.AsOf.Format("2006-01-02"),
		BudgetUSD:       p.Budget,
	}, acc)
	if err != nil {
		out.err = err
		return out
	}
	out.deck, out.notes, out.repaired, out.repairReason = res.Deck, res.Notes, res.Repaired, res.RepairReason
	return out
}

// ownedFor is the collection a prompt reads, empty when it wants none.
func ownedFor(p prompt, owned map[string]int32) map[string]int32 {
	if p.Collection {
		return owned
	}
	return map[string]int32{}
}
