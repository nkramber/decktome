// Command deck-gate runs the PR-8 golden prompts through the generator
// and writes the gate document.
//
// The bars come from the roadmap. Every returned deck passes the block
// checks, and no invented name reaches the user.
//
// CAUTION: this calls a real provider and it costs money. DECK_GATE=1 is
// required, so it can not run by accident. `make deck-gate` writes to
// DECK_GATE_OUT and refuses a file that already holds a verdict: a rerun
// must never overwrite a scored document (D-65). A -dry run calls no
// provider and needs no guard. The exit code is 1 on a FAIL verdict, and
// the document is written first.
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
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
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
	// CollectionFile names a ManaBox export beside the file of the
	// -collection flag, for a prompt that needs another binder. Empty
	// reads the flag's file.
	CollectionFile string `json:"collection_file"`
	// ExcludePrecons names the precons the deck must use no card of
	// (D-408). The binder must hold each one whole, or the prompt fails.
	ExcludePrecons []string `json:"exclude_precons"`
	Locked         []string `json:"locked"`
	// Precon names a preconstructed deck the build must keep a share of
	// (D-218, D-247).
	Precon string  `json:"precon"`
	Budget float64 `json:"budget"`
	// Sets names the sets the reader wrote, in their own words. The gate
	// resolves each one to a set family through the index (D-376), so a
	// prompt reads the way a reader speaks and never as a code list.
	Sets []string `json:"sets"`
	// OutsideMana allows the mana fill of D-382. The reader answers the
	// mana row with a yes, and the gate says so here.
	OutsideMana bool   `json:"outside_mana"`
	Plan        string `json:"plan"`
}

type result struct {
	prompt       prompt
	deck         *mtgv1.Deck
	notes        []string
	repaired     bool
	repairReason string
	poolSize     int
	// setCodes is the family the prompt's set names resolved to, and
	// inSet is how many pool cards the sets hold (D-376, D-380).
	setCodes []string
	inSet    int
	outside  int
	// products names the precons the prompt excluded, excluded holds the
	// Oracle ids the exclusion took out of the pool, and spare counts
	// the product cards that stayed usable on a spare copy (D-408).
	products []string
	excluded map[string]bool
	spare    int
	judged   *generate.Judgement
	// judgeErr is the judge lane's failure. A deck with one has no
	// verdict on F-26, so it can not count as a pass on that bar (T-17).
	// An empty summary counts as one: the judge has nothing to read.
	judgeErr error
	// plan is the plan judge's read of the deck (PR-15). Its rows are
	// information, and a failure of the lane moves no verdict.
	plan    *generate.PlanJudgement
	planErr error
	err     error
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
	noJudge := flag.Bool("no-judge", false, "skip the F-26 judge lane, which costs one judge call a deck")
	runOut := flag.String("run-out", "", "write the run header and the rows as JSONL here (PR-15)")
	flag.Parse()

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
	file.Prompts = prompts

	// A dry run calls no provider, so it needs no guard.
	if !*dry {
		if err := gatekit.SpendGuard("DECK_GATE"); err != nil {
			return err
		}
	}
	// The run file is never overwritten (D-65), and the check runs before
	// the first provider call.
	if err := gatekit.RefuseExisting(*runOut); err != nil {
		return err
	}
	run := evalrun.New("decks", evalrun.RunID(*runOut))
	run.Header.Only = *only
	run.Header.Prompts["generate"] = generate.PromptVersion
	run.Header.Prompts["plan_rubric"] = generate.PlanRubricVersion
	run.LowerIsBetter("blocks", "invented_names", "false_rules", "judge_error", "excluded_in_deck", "warnings", "repaired", "buy_cost", "deck_cost")
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(context.Background(), quiet)
	if err != nil {
		return err
	}
	run.SetSnapshot(idx.AsOf)
	binders, err := loadBinders(*collPath, file.Prompts, idx)
	if err != nil {
		return err
	}
	// The precon table serves the exclusion prompts alone, so a run with
	// none needs no meta store (D-407).
	var preconTbl *precons.Table
	for _, p := range file.Prompts {
		if len(p.ExcludePrecons) == 0 {
			continue
		}
		if preconTbl, err = gatekit.PreconTable(context.Background()); err != nil {
			return fmt.Errorf("precon table: %w", err)
		}
		if preconTbl == nil {
			return fmt.Errorf("prompt %d excludes a precon, and the meta store holds no precon table: run make meta-refresh", p.ID)
		}
		run.Header.Versions["precons"] = preconTbl.Version
		break
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
	for _, p := range file.Prompts {
		r := build(context.Background(), b, cb, idx, binders, p, acc, *dry, preconSet, preconTbl)
		// The judge lane is the real check for F-26, and the deterministic
		// net can not read the truth of a rules claim (D-229).
		if !*dry && !*noJudge && r.deck != nil {
			r.judged, r.judgeErr = judge(context.Background(), client, p.Name, r.deck, acc)
			if r.judgeErr != nil {
				fmt.Fprintf(os.Stderr, "  judge %d failed: %v\n", p.ID, r.judgeErr)
			}
			r.plan, r.planErr = generate.JudgePlan(context.Background(), client, p.Plan, r.deck, idx, acc)
			if r.planErr != nil {
				fmt.Fprintf(os.Stderr, "  plan judge %d failed: %v\n", p.ID, r.planErr)
			}
		}
		results = append(results, r)
		fmt.Fprintf(os.Stderr, "  %2d. %-38s pool %3d%s%s  %s\n", p.ID, p.Name, r.poolSize, setWord(r), excludeWord(r), status(r))
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
var errGateFailed = errors.New("deck gate failed: read the verdict line of the document")

// selectPrompts keeps the prompts -only names, or all of them when -only
// is empty. A list that names no prompt is an error: a run of zero
// prompts can not pass a gate.
func selectPrompts(all []prompt, only string) ([]prompt, error) {
	ids, err := gatekit.ParseIDs(only)
	if err != nil {
		return nil, fmt.Errorf("deck-gate: %w", err)
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

// errEmptySummary is the judge error of a deck with no summary. The F-26
// bar needs a verdict on the summary, and an empty one gives none.
var errEmptySummary = errors.New("the deck has no summary to judge")

// judge runs the F-26 judge lane on one deck. An empty summary is a judge
// error and not a pass.
func judge(ctx context.Context, client *llm.Client, name string, d *mtgv1.Deck, acc *llm.Accumulator) (*generate.Judgement, error) {
	if d.GetSummary() == "" {
		return nil, errEmptySummary
	}
	return generate.JudgeSummary(ctx, client, name, d.GetSummary(), acc)
}

// setWord names the set limit of one result, for the progress line and
// the gate document. An unlimited prompt reads as nothing (D-373).
func setWord(r result) string {
	if len(r.setCodes) == 0 {
		return ""
	}
	return fmt.Sprintf("  sets %s in %d out %d", strings.Join(r.setCodes, ","), r.inSet, r.outside)
}

func status(r result) string {
	switch {
	case r.err != nil:
		return "ERROR: " + r.err.Error()
	case r.deck == nil:
		return "no deck"
	default:
		return fmt.Sprintf("%d cards, %d blocks, %d notes, repaired %v",
			gatekit.CountCards(r.deck), len(gatekit.BlockFindings(r.deck)), len(r.notes), r.repaired)
	}
}

func build(ctx context.Context, b *generate.Builder, cb *candidates.Builder, idx *cards.Index,
	binders map[string]*gatekit.Collection, p prompt, acc *llm.Accumulator, dry bool, preconSet *precons.Set, tbl *precons.Table) result {
	out := result{prompt: p}
	format := gatekit.FormatID(p.Format)
	if format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		out.err = fmt.Errorf("unknown format %q", p.Format)
		return out
	}
	// The set names resolve to a set family, the way the chat does
	// (D-376). A name this snapshot can not settle fails the prompt: the
	// gate proves the filter, and the chat asks the reader.
	setCodes, err := resolveSets(idx, p.Sets)
	if err != nil {
		out.err = err
		return out
	}
	out.setCodes = setCodes
	binder := binderFor(p, binders)
	own := binder.Oracle
	// The precon exclusion runs before both pools, the way the chat runs
	// it (D-408). The products' copies leave the owned counts, and a
	// card with no copy left leaves the commander pool and the
	// shortlist. A basic land never leaves (D-37).
	var excludedIDs []string
	if len(p.ExcludePrecons) > 0 {
		products, names, err := resolvePrecons(tbl, p.ExcludePrecons, binder.Printings)
		if err != nil {
			out.err = err
			return out
		}
		own, excludedIDs = precons.Exclude(products, own, func(id string) bool {
			c, ok := idx.ByOracleID(id)
			return ok && candidates.IsBasicLand(c)
		})
		out.products = names
		out.excluded = make(map[string]bool, len(excludedIDs))
		for _, id := range excludedIDs {
			out.excluded[id] = true
		}
		out.spare = spareCards(products, out.excluded, idx)
	}
	var commanders []*mtgv1.Card
	var commanderIDs []string
	// A prompt with no commander delegates the pick, as a user who says
	// "you pick" does. The generator must choose one (D-232).
	if p.Commander == "" && format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		pool, err := cb.CommanderPool(idx, candidates.Request{
			Format: format, Theme: p.Theme, Colors: gatekit.Colors(p.Colors),
			PoolRule: gatekit.PoolRuleID(p.Pool), Owned: own, Bracket: p.Bracket,
			SetCodes: setCodes, ExcludeOracleIDs: excludedIDs,
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
	// The viability floor of D-380. A family too thin for the format in
	// these colors builds nothing, and the reason names the counts.
	if len(setCodes) > 0 {
		have := candidates.CountInSets(idx, candidates.Request{Format: format, Colors: colors, SetCodes: setCodes})
		if want := candidates.SetFloor(format); have < want {
			out.err = fmt.Errorf("the sets hold %d cards in these colors, and this format needs about %d", have, want)
			return out
		}
	}
	var outsideRoles map[mtgv1.CardRole]int
	if len(setCodes) > 0 && p.OutsideMana {
		outsideRoles = manaRoles(generate.TargetsFor(format, gatekit.PowerLevel(p.Bracket, p.Power)))
	}
	list, err := cb.Build(idx, candidates.Request{
		Format:             format,
		Colors:             colors,
		Theme:              p.Theme,
		CommanderOracleIDs: commanderIDs,
		PoolRule:           poolRule,
		Owned:              own,
		Bracket:            p.Bracket,
		SetCodes:           setCodes,
		OutsideRoles:       outsideRoles,
		ExcludeOracleIDs:   excludedIDs,
	})
	if err != nil {
		out.err = fmt.Errorf("candidates: %w", err)
		return out
	}
	out.inSet, out.outside = list.Stats.InSet, list.Stats.Outside
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
		SessionID:         fmt.Sprintf("gate-%d", p.ID),
		Format:            format,
		Power:             gatekit.PowerLevel(p.Bracket, p.Power),
		Plan:              plan,
		Pool:              pool,
		Commanders:        commanderIDs,
		Locked:            lockedIDs,
		Precon:            preconName,
		PreconOracleIDs:   preconIDs,
		PreconLands:       preconLands,
		PoolRule:          poolRule,
		OracleCounts:      own,
		ExcludedOracleIDs: excludedIDs,
		Roles:             generate.Roles(list),
		Targets:           generate.TargetsFor(format, gatekit.PowerLevel(p.Bracket, p.Power)),
		Limits:            generate.LimitsFor(format),
		LegalityAsOf:      idx.AsOf.Format("2006-01-02"),
		BudgetUSD:         p.Budget,
		SetCodes:          setCodes,
	}, acc)
	if err != nil {
		out.err = err
		return out
	}
	out.deck, out.notes, out.repaired, out.repairReason = res.Deck, res.Notes, res.Repaired, res.RepairReason
	return out
}

// resolveSets maps the set names of a prompt onto one set family, the
// way the chat does (D-376). A name this snapshot can not settle is an
// error here: the gate proves the filter, and the chat asks the reader.
func resolveSets(idx *cards.Index, names []string) ([]string, error) {
	if len(names) == 0 {
		return nil, nil
	}
	tbl := idx.Sets()
	seen := map[string]bool{}
	var out []string
	for _, name := range names {
		res := tbl.Resolve(name)
		if res.Kind != cards.ResolveOne {
			return nil, fmt.Errorf("the set name %q resolves to %d sets, not one", name, len(res.Candidates))
		}
		for _, c := range res.Codes {
			if !seen[c] {
				seen[c] = true
				out = append(out, c)
			}
		}
	}
	sort.Strings(out)
	return out, nil
}

// manaRoles are the roles a set-limited deck may fill from outside the
// named sets (D-382). It mirrors the same function in agentsvc, because
// the gate builds a request the way the service does.
func manaRoles(targets map[string]int) map[mtgv1.CardRole]int {
	out := map[mtgv1.CardRole]int{}
	if n := targets["ramp"]; n > 0 {
		out[mtgv1.CardRole_CARD_ROLE_RAMP] = n
	}
	if n := targets["land"]; n > 0 {
		out[mtgv1.CardRole_CARD_ROLE_LAND] = n
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

// loadBinders reads the collection of the -collection flag under the
// empty key, and each collection_file a prompt names beside it. A prompt
// with a collection_file and no flag is an error: the file has no place
// to sit beside.
func loadBinders(collPath string, prompts []prompt, idx *cards.Index) (map[string]*gatekit.Collection, error) {
	binders := map[string]*gatekit.Collection{}
	if collPath != "" {
		c, err := gatekit.LoadCollection(collPath, idx)
		if err != nil {
			return nil, err
		}
		binders[""] = c
	}
	for _, p := range prompts {
		if p.CollectionFile == "" {
			continue
		}
		if _, ok := binders[p.CollectionFile]; ok {
			continue
		}
		if collPath == "" {
			return nil, fmt.Errorf("prompt %d names collection file %q, and -collection is empty", p.ID, p.CollectionFile)
		}
		c, err := gatekit.LoadCollection(filepath.Join(filepath.Dir(collPath), p.CollectionFile), idx)
		if err != nil {
			return nil, fmt.Errorf("prompt %d: %w", p.ID, err)
		}
		binders[p.CollectionFile] = c
	}
	return binders, nil
}

// binderFor is the collection a prompt reads, empty when it wants none.
func binderFor(p prompt, binders map[string]*gatekit.Collection) *gatekit.Collection {
	if c, ok := binders[p.CollectionFile]; ok && p.Collection {
		return c
	}
	return &gatekit.Collection{Oracle: map[string]int32{}}
}

// resolvePrecons maps the product names of a prompt onto the table, the
// way the chat does, and checks that the binder holds one product of
// each name whole (D-408). A deck and its Collector's Edition answer one
// name together, and the exclusion counts them once. A name that names
// no product fails the prompt: the gate proves the exclusion, and the
// chat asks the reader.
func resolvePrecons(tbl *precons.Table, names []string, printings map[string]int32) ([]*precons.Product, []string, error) {
	var products []*precons.Product
	var found []string
	for _, name := range names {
		m := tbl.Resolve(name)
		if !m.OK() {
			return nil, nil, fmt.Errorf("the precon name %q names no product, near %v", name, m.Options)
		}
		whole := false
		for _, p := range m.Products {
			if p.OwnedWhole(printings) {
				whole = true
			}
		}
		if !whole {
			return nil, nil, fmt.Errorf("the collection does not hold %s whole", strings.Join(m.Names(), " or "))
		}
		products = append(products, m.Products...)
		found = append(found, m.Names()...)
	}
	return products, found, nil
}

// spareCards counts the nonbasic cards of the products the exclusion
// left in the pool, because the binder holds a copy to spare (D-408).
// The gate document names the count, so a reader can tell a spare copy
// from a card that slipped through.
func spareCards(products []*precons.Product, excluded map[string]bool, idx *cards.Index) int {
	seen := map[string]bool{}
	for _, p := range products {
		for id := range p.Counts() {
			if seen[id] || excluded[id] {
				continue
			}
			if c, ok := idx.ByOracleID(id); ok && candidates.IsBasicLand(c) {
				continue
			}
			seen[id] = true
		}
	}
	return len(seen)
}

// excludeWord names the precon exclusion of one result, for the
// progress line. A prompt with none reads as nothing.
func excludeWord(r result) string {
	if len(r.products) == 0 {
		return ""
	}
	return fmt.Sprintf("  excludes %d cards of %s, %d spare", len(r.excluded), strings.Join(r.products, ", "), r.spare)
}
