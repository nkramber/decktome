// Command revise-gate runs the PR-12B revision prompts and writes the
// gate document. Each base deck is built once, then each revision runs
// from that base: the revise call, the revised build, and the checks.
//
// The bars come from the roadmap entry PR-12B. Every revised deck keeps
// the untouched cards, holds every limit of the brief, and passes the
// engine. An unclear request gets a question and no build.
//
// CAUTION: this calls a real provider and it costs money. REVISE_GATE=1
// is required, so it can not run by accident. `make revise-gate` writes
// to REVISE_GATE_OUT and refuses a file that already holds a verdict
// (D-65). A -dry run loads the prompts and the snapshot, calls no
// provider, and needs no guard. The exit code is 1 on a FAIL verdict,
// and the document is written first.
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/revise"
	"github.com/nkramber/decktome/go/internal/rules"
)

//go:embed prompts.json
var promptsJSON []byte

type base struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Format    string   `json:"format"`
	Power     string   `json:"power"`
	Bracket   int32    `json:"bracket"`
	Colors    []string `json:"colors"`
	Commander string   `json:"commander"`
	Theme     string   `json:"theme"`
	Plan      string   `json:"plan"`
	// Sets names the sets the reader wrote, in their own words. The gate
	// resolves each one to a set family through the index (D-376). A
	// revision reads the same slots as the build, so a set-limited base
	// proves the filter survives a revision (D-373).
	Sets      []string   `json:"sets"`
	Revisions []revision `json:"revisions"`
}

type revision struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	// Cap is the mana value the message names. The check holds the deck
	// to it whatever the brief read.
	Cap float64 `json:"cap"`
	// RemoveHighest fills {{HIGHEST}} with the base deck's dearest
	// nonland card, and the check wants it gone.
	RemoveHighest bool `json:"remove_highest"`
	// Unclear says the right outcome is a question, not a build. A cap
	// on an unclear prompt is the clear part of a mixed message, and the
	// brief must still hold it (D-284).
	Unclear bool `json:"unclear"`
	// Answer is what the gate replies to the question an unclear request
	// earns. With one, a second turn runs: the answer with the first
	// message as the prior, then the rebuild and every check of a clear
	// request (D-448). The gate never played that turn before, and the
	// product failed on it.
	Answer string `json:"answer"`
	// SwapBasics says the message, or the answer, asks for basic lands
	// replaced. The brief must count them, and the revised deck must
	// hold that many more nonbasic lands.
	SwapBasics bool `json:"swap_basics"`
	// Note is for the reader of the prompt file.
	Note string `json:"note"`
}

// keepBar is the share of untouched base names a revised deck must keep.
// A revision that asks for swaps moves some, so the bar is not 100.
const keepBar = 0.8

type outcome struct {
	base base
	rev  revision
	// answered marks the second turn of an unclear request: the gate's
	// answer to the question, and the rebuild it earns.
	answered bool
	message  string
	brief    *revise.Brief
	deck     *mtgv1.Deck
	// repair names what bought the repair turn, or is empty when none
	// ran, so a reader sees whether the model got a second try.
	repair   string
	note     string
	kept     float64
	blocks   []string
	failures []string
	err      error
}

func (o outcome) pass() bool { return o.err == nil && len(o.failures) == 0 }

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	only := flag.String("only", "", "run these base ids only, comma separated")
	dry := flag.Bool("dry", false, "load the prompts and the snapshot and stop before the provider calls")
	runOut := flag.String("run-out", "", "write the run header and the rows as JSONL here (PR-15)")
	flag.Parse()

	var file struct {
		Bases []base `json:"bases"`
	}
	if err := json.Unmarshal(promptsJSON, &file); err != nil {
		return fmt.Errorf("prompts: %w", err)
	}
	bases, err := selectBases(file.Bases, *only)
	if err != nil {
		return err
	}
	file.Bases = bases
	// A dry run calls no provider, so it needs no guard.
	if !*dry {
		if err := gatekit.SpendGuard("REVISE_GATE"); err != nil {
			return err
		}
	}
	// The run file is never overwritten (D-65), and the check runs before
	// the first provider call.
	if err := gatekit.RefuseExisting(*runOut); err != nil {
		return err
	}
	run := evalrun.New("revise", evalrun.RunID(*runOut))
	run.Header.Only = *only
	run.Header.Prompts["generate"] = generate.PromptVersion
	run.Header.Prompts["revise"] = revise.PromptVersion
	run.LowerIsBetter("blocks", "repaired")
	quiet := gatekit.Quiet()
	ctx := context.Background()
	idx, err := gatekit.LoadSnapshot(ctx, quiet)
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
	if *dry {
		revisions := 0
		for _, bs := range file.Bases {
			if bs.Commander != "" {
				if _, ok := idx.ByName(bs.Commander); !ok {
					return fmt.Errorf("base %d: no card named %q", bs.ID, bs.Commander)
				}
			}
			revisions += len(bs.Revisions)
		}
		fmt.Fprintf(os.Stderr, "dry run: %d bases and %d revisions read, no provider call ran\n", len(file.Bases), revisions)
		return nil
	}
	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return err
	}
	run.SetRoles(client.Config(), llm.RoleGenerate, llm.RoleRepair, llm.RoleRevise)
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)
	prof, err := gatekit.Profiler(idx, rcfg, quiet)
	if err != nil {
		return err
	}
	// The quality model grades every deck the gate builds, and the
	// summary names the tier (PR-14B). No stored model grades nothing.
	scorer, err := gatekit.Scorer(ctx)
	if err != nil {
		return err
	}
	b := generate.NewBuilder(client, rcfg, idx, quiet, generate.WithProfiler(prof), generate.WithScorer(scorer))

	start := time.Now()
	var outcomes []outcome
	for _, bs := range file.Bases {
		fmt.Fprintf(os.Stderr, "base %d. %s\n", bs.ID, bs.Name)
		baseDeck, pool, list, commanderIDs, err := buildBase(ctx, b, cb, idx, bs, acc)
		if err != nil {
			for _, r := range bs.Revisions {
				outcomes = append(outcomes, outcome{base: bs, rev: r, err: fmt.Errorf("base build: %w", err)})
			}
			fmt.Fprintf(os.Stderr, "  base failed: %v\n", err)
			continue
		}
		fmt.Fprintf(os.Stderr, "  %d cards, %d blocks\n", gatekit.CountCards(baseDeck), len(blocks(baseDeck)))
		for _, r := range bs.Revisions {
			for _, o := range runRevision(ctx, client, b, idx, bs, r, baseDeck, pool, list, commanderIDs, acc) {
				outcomes = append(outcomes, o)
				turn := ""
				if o.answered {
					turn = " (the answer)"
				}
				fmt.Fprintf(os.Stderr, "  %d%s. %s\n", r.ID, turn, status(o))
			}
		}
	}
	pass := report(os.Stdout, outcomes, acc, idx, time.Since(start), run)
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
var errGateFailed = errors.New("revise gate failed: read the verdict line of the document")

// selectBases keeps the bases -only names, or all of them when -only is
// empty. A list that names no base is an error: a run of zero revisions
// can not pass a gate.
func selectBases(all []base, only string) ([]base, error) {
	ids, err := gatekit.ParseIDs(only)
	if err != nil {
		return nil, fmt.Errorf("revise-gate: %w", err)
	}
	if ids == nil {
		return all, nil
	}
	want := map[int]bool{}
	for _, id := range ids {
		want[id] = true
	}
	var kept []base
	for _, b := range all {
		if want[b.ID] {
			kept = append(kept, b)
		}
	}
	if len(kept) == 0 {
		return nil, fmt.Errorf("-only %q matches no base", only)
	}
	return kept, nil
}

func status(o outcome) string {
	switch {
	case o.err != nil:
		return "ERROR: " + o.err.Error()
	case len(o.failures) > 0:
		return "FAIL: " + strings.Join(o.failures, "; ")
	case o.deck == nil:
		return "PASS, no build"
	default:
		return fmt.Sprintf("PASS, kept %.0f%%", o.kept*100)
	}
}

func formatOf(bs base) mtgv1.FormatId { return gatekit.FormatID(bs.Format) }

func powerOf(bs base) *mtgv1.PowerLevel { return gatekit.PowerLevel(bs.Bracket, bs.Power) }

func powerWord(bs base) string {
	if bs.Bracket > 0 {
		return fmt.Sprintf("bracket %d", bs.Bracket)
	}
	return strings.ToLower(bs.Power)
}

// buildBase builds the deck the revisions start from, the way the
// product does: the shortlist, the basics, and the commander.
func buildBase(ctx context.Context, b *generate.Builder, cb *candidates.Builder, idx *cards.Index, bs base,
	acc *llm.Accumulator) (*mtgv1.Deck, *generate.Pool, *candidates.List, []string, error) {
	format := formatOf(bs)
	if format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return nil, nil, nil, nil, fmt.Errorf("unknown format %q", bs.Format)
	}
	var commanders []*mtgv1.Card
	var commanderIDs []string
	if bs.Commander != "" {
		c, ok := idx.ByName(bs.Commander)
		if !ok {
			return nil, nil, nil, nil, fmt.Errorf("no card named %q", bs.Commander)
		}
		commanders = append(commanders, c)
		commanderIDs = append(commanderIDs, c.GetOracleId())
	}
	colors := gatekit.Colors(bs.Colors)
	if len(commanders) > 0 {
		colors = commanders[0].GetColorIdentity()
	}
	setCodes, err := resolveSets(idx, bs.Sets)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	list, err := cb.Build(idx, candidates.Request{
		Format: format, Colors: colors, Theme: bs.Theme, CommanderOracleIDs: commanderIDs,
		PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD, Bracket: bs.Bracket,
		SetCodes: setCodes,
	})
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("candidates: %w", err)
	}
	always := append([]*mtgv1.Card(nil), commanders...)
	always = append(always, generate.BasicLands(idx.ByName, colors)...)
	pool := generate.FromList(list, always, true)
	res, err := b.Build(ctx, request(bs, idx, pool, list, commanderIDs, nil), acc)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	res.Deck.Id = fmt.Sprintf("base-%d", bs.ID)
	res.Deck.Name = bs.Name
	return res.Deck, pool, list, commanderIDs, nil
}

func request(bs base, idx *cards.Index, pool *generate.Pool, list *candidates.List, commanderIDs []string, rev *generate.Revision) generate.Request {
	format := formatOf(bs)
	return generate.Request{
		SessionID:    fmt.Sprintf("revise-gate-%d", bs.ID),
		Format:       format,
		Power:        powerOf(bs),
		Plan:         bs.Plan,
		Pool:         pool,
		Commanders:   commanderIDs,
		PoolRule:     mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		Roles:        generate.Roles(list),
		Targets:      generate.TargetsFor(format, powerOf(bs)),
		Limits:       generate.LimitsFor(format),
		LegalityAsOf: idx.AsOf.Format("2006-01-02"),
		Revision:     rev,
		SetCodes:     setCodesOf(idx, bs),
	}
}

// setCodesOf resolves the base's set names. A name this snapshot can not
// settle is a gate error, and build() reports it before this runs.
func setCodesOf(idx *cards.Index, bs base) []string {
	codes, _ := resolveSets(idx, bs.Sets)
	return codes
}

// resolveSets maps the set names of a base onto one set family, the way
// the chat does (D-376). A name this snapshot can not settle is an error
// here: the gate proves the filter, and the chat asks the reader.
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

// highestNonland is the base deck's dearest nonland card by mana value.
func highestNonland(deck *mtgv1.Deck, idx *cards.Index) string {
	best, bestMV := "", -1.0
	for _, dc := range deck.GetCards() {
		c, ok := idx.ByOracleID(dc.GetOracleId())
		if !ok || isLand(c) {
			continue
		}
		if c.GetManaValue() > bestMV {
			best, bestMV = c.GetName(), c.GetManaValue()
		}
	}
	return best
}

func isLand(c *mtgv1.Card) bool {
	for _, t := range c.GetCardTypes() {
		if t == "Land" {
			return true
		}
	}
	return false
}

// runRevision is one revision from the base: the revise call, then the
// build with the brief, then the checks against the bars. An unclear
// request with an answer gets a second turn, so the list holds two
// outcomes: the question, and the rebuild the answer earns (D-448).
func runRevision(ctx context.Context, client *llm.Client, b *generate.Builder, idx *cards.Index, bs base, r revision,
	baseDeck *mtgv1.Deck, basePool *generate.Pool, list *candidates.List, commanderIDs []string, acc *llm.Accumulator) []outcome {
	o := outcome{base: bs, rev: r}
	highest := highestNonland(baseDeck, idx)
	o.message = strings.ReplaceAll(r.Message, "{{HIGHEST}}", highest)
	brief, err := call(ctx, client, bs, o.message, "", baseDeck, idx, acc)
	if err != nil {
		o.err = err
		return []outcome{o}
	}
	o.brief = brief
	if !r.Unclear {
		return []outcome{rebuild(ctx, b, idx, bs, r, o, highest, baseDeck, basePool, list, commanderIDs, acc)}
	}
	if brief.Question == "" {
		o.failures = append(o.failures, "an unclear request got no question")
	}
	if brief.Question == "" && !brief.Acts() && len(brief.Declined) > 0 {
		// A decline with a reason is the other right answer to a land
		// upgrade in a one-color deck (D-284).
		o.failures = o.failures[:0]
	}
	o.note = revise.DeclineNote(brief)
	if brief.Question != "" {
		o.note = "Question: " + brief.Question
	}
	// The clear part of a mixed message must be in the brief even
	// when the unclear part earns a question.
	if r.Cap > 0 && brief.MaxManaValue != r.Cap {
		o.failures = append(o.failures, fmt.Sprintf("the brief read the cap as %g, and the message says %g", brief.MaxManaValue, r.Cap))
	}
	if r.Answer == "" || brief.Question == "" {
		return []outcome{o}
	}
	// The second turn: the answer, in the shape the chat sends it, with
	// the first message as the prior (agentsvc.withAnswers).
	a := outcome{base: bs, rev: r, answered: true}
	a.message = "Q: " + brief.Question + "\nA: " + r.Answer
	brief2, err := call(ctx, client, bs, a.message, o.message, baseDeck, idx, acc)
	if err != nil {
		a.err = err
		return []outcome{o, a}
	}
	a.brief = brief2
	return []outcome{o, rebuild(ctx, b, idx, bs, r, a, highest, baseDeck, basePool, list, commanderIDs, acc)}
}

func call(ctx context.Context, client *llm.Client, bs base, message, prior string, baseDeck *mtgv1.Deck, idx *cards.Index, acc *llm.Accumulator) (*revise.Brief, error) {
	return revise.Call(ctx, client, revise.Input{
		SessionID: fmt.Sprintf("revise-gate-%d", bs.ID), Message: message, Prior: prior, Deck: baseDeck,
		Format: generate.FormatWord(formatOf(bs)), Power: powerWord(bs), Cards: idx,
	}, acc)
}

// rebuild is the clear-request half of a revision: the brief must act,
// the build runs with it, and the checks hold the deck to the message.
func rebuild(ctx context.Context, b *generate.Builder, idx *cards.Index, bs base, r revision, o outcome, highest string,
	baseDeck *mtgv1.Deck, basePool *generate.Pool, list *candidates.List, commanderIDs []string, acc *llm.Accumulator) outcome {
	brief := o.brief
	if brief.Question != "" {
		o.failures = append(o.failures, "a clear request got a question: "+brief.Question)
		o.note = "Question: " + brief.Question
		return o
	}
	if !brief.Acts() {
		o.failures = append(o.failures, "a clear request was declined: "+revise.DeclineNote(brief))
		o.note = revise.DeclineNote(brief)
		return o
	}
	if r.Cap > 0 && brief.MaxManaValue != r.Cap {
		o.failures = append(o.failures, fmt.Sprintf("the brief read the cap as %g, and the message says %g", brief.MaxManaValue, r.Cap))
	}
	if r.RemoveHighest && !contains(brief.Remove, highest) {
		o.failures = append(o.failures, fmt.Sprintf("the brief did not list %q to remove", highest))
	}
	if r.SwapBasics && brief.SwapBasics == 0 {
		o.failures = append(o.failures, "the brief counted no basic lands to replace")
	}
	rev := &generate.Revision{
		BaseDeckID: baseDeck.GetId(), Base: baseDeck.GetCards(), Instructions: brief.Changes,
		Remove: brief.Remove, Keep: brief.Keep, MaxManaValue: brief.MaxManaValue,
		SwapBasics: brief.SwapBasics, LandKinds: brief.LandKinds,
	}
	// The base deck's cards join the pool, the way agentsvc does it, and
	// the brief filters it.
	var always []*mtgv1.Card
	for _, dc := range baseDeck.GetCards() {
		if c, ok := idx.ByOracleID(dc.GetOracleId()); ok {
			always = append(always, c)
		}
	}
	pool := generate.FromList(list, append(always, poolCards(basePool)...), true)
	pool = pool.Filter(func(c *mtgv1.Card) bool { return generate.AllowedByRevision(rev, c) })
	generate.FitSwapBasics(rev, pool)
	res, err := b.Build(ctx, request(bs, idx, pool, list, commanderIDs, rev), acc)
	if err != nil {
		o.err = err
		return o
	}
	o.deck = res.Deck
	if res.Repaired {
		o.repair = res.RepairReason
	}
	o.note = revise.Note(brief, revise.DiffDecks(baseDeck, res.Deck))
	o.blocks = append(o.blocks, blocks(res.Deck)...)
	if len(o.blocks) > 0 {
		o.failures = append(o.failures, "block findings: "+strings.Join(o.blocks, ", "))
	}
	// The checks hold to the message, not only to the brief.
	if r.Cap > 0 {
		for _, dc := range res.Deck.GetCards() {
			if c, ok := idx.ByOracleID(dc.GetOracleId()); ok && !isLand(c) && c.GetManaValue() > r.Cap {
				o.failures = append(o.failures, fmt.Sprintf("%s (%g) is over the cap of %g", c.GetName(), c.GetManaValue(), r.Cap))
			}
		}
	}
	if r.RemoveHighest && hasName(res.Deck, highest) {
		o.failures = append(o.failures, fmt.Sprintf("%q is still in the deck", highest))
	}
	if r.SwapBasics {
		rise := nonbasicLands(res.Deck, idx) - nonbasicLands(baseDeck, idx)
		if rise < rev.SwapBasics || rise == 0 {
			o.failures = append(o.failures, fmt.Sprintf("the deck holds %d more nonbasic lands, and the brief asked for %d", rise, rev.SwapBasics))
		}
	}
	o.kept = keptShare(baseDeck, res.Deck, rev, idx)
	if o.kept < keepBar {
		o.failures = append(o.failures, fmt.Sprintf("kept %.0f%% of the untouched cards, the bar is %.0f%%", o.kept*100, keepBar*100))
	}
	return o
}

// nonbasicLands sums the copies of every nonbasic land in a deck.
func nonbasicLands(d *mtgv1.Deck, idx *cards.Index) int {
	n := 0
	for _, dc := range d.GetCards() {
		if c, ok := idx.ByOracleID(dc.GetOracleId()); ok && isLand(c) && !candidates.IsBasicLand(c) {
			n += int(dc.GetCount())
		}
	}
	return n
}

// poolCards lists the cards of a pool by name lookup.
func poolCards(p *generate.Pool) []*mtgv1.Card {
	var out []*mtgv1.Card
	for _, name := range p.Names() {
		if c, ok := p.Card(name); ok {
			out = append(out, c)
		}
	}
	return out
}

// keptShare is the share of untouched base names the revised deck still
// holds. Untouched means not removed by the brief and not over its cap.
func keptShare(baseDeck, revised *mtgv1.Deck, rev *generate.Revision, idx *cards.Index) float64 {
	untouched, kept := 0, 0
	for _, dc := range baseDeck.GetCards() {
		c, ok := idx.ByOracleID(dc.GetOracleId())
		if !ok || !generate.AllowedByRevision(rev, c) {
			continue
		}
		untouched++
		if hasName(revised, dc.GetName()) {
			kept++
		}
	}
	if untouched == 0 {
		return 1
	}
	return float64(kept) / float64(untouched)
}

func hasName(d *mtgv1.Deck, name string) bool {
	for _, c := range d.GetCards() {
		if strings.EqualFold(c.GetName(), name) {
			return true
		}
	}
	return false
}

func contains(list []string, name string) bool {
	for _, n := range list {
		if strings.EqualFold(n, name) {
			return true
		}
	}
	return false
}

// blocks names the codes of the findings that stop the deck.
func blocks(d *mtgv1.Deck) []string {
	var out []string
	for _, f := range gatekit.BlockFindings(d) {
		out = append(out, f.GetCode())
	}
	return out
}

// report writes the gate document and returns the verdict. A run of zero
// revisions fails: there is nothing to pass.
func report(w io.Writer, outcomes []outcome, acc *llm.Accumulator, idx *cards.Index, took time.Duration, run *evalrun.Run) bool {
	// The document is the record, and a write error on stdout ends the
	// process in any case.
	pf := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	passed := 0
	for _, o := range outcomes {
		if o.pass() {
			passed++
		}
	}
	pass := len(outcomes) > 0 && passed == len(outcomes)
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	recordRows(run, outcomes)
	run.Finish(acc.Report(), took, verdict)
	pf("# PR-12B revise gate\n\nRun date: %s. Snapshot: %s.\n\n", time.Now().UTC().Format("2006-01-02"), idx.AsOf.Format("2006-01-02"))
	pf("Verdict: %s. %d of %d turns met their bar. The bars: an unclear request gets a question or a decline with a reason, an answer to that question gets a rebuild, a clear request gets a revised deck with no block finding, the deck holds every cap, every removal, and every land swap the message names, and it keeps at least %.0f percent of the untouched cards.\n\n",
		verdict, passed, len(outcomes), keepBar*100)
	r := acc.Report()
	pf("Cost: %d calls, %s. Time: %s.\n\n", r.Calls, gatekit.CostWord(r), took.Round(time.Second))
	pf("| Base | Revision | Message | Outcome | Kept | Result |\n|---|---|---|---|---|---|\n")
	for _, o := range outcomes {
		out := "no build"
		if o.deck != nil {
			out = fmt.Sprintf("%d cards", gatekit.CountCards(o.deck))
		}
		if o.brief != nil && o.brief.Question != "" {
			out = "question"
		}
		kept := ""
		if o.deck != nil {
			kept = fmt.Sprintf("%.0f%%", o.kept*100)
		}
		msg := cell(o.message)
		if o.answered {
			msg = "(the answer) " + msg
		}
		pf("| %d | %d | %s | %s | %s | %s |\n", o.base.ID, o.rev.ID, msg, out, kept, cell(status(o)))
	}
	for _, o := range outcomes {
		turn := ""
		if o.answered {
			turn = ", the answer"
		}
		pf("\n## Revision %d%s, base %d: %s\n\n", o.rev.ID, turn, o.base.ID, o.base.Name)
		pf("**The user:** %s\n\n", cell(o.message))
		if o.err != nil {
			pf("ERROR: %v\n", o.err)
			continue
		}
		if o.brief != nil {
			raw, _ := json.MarshalIndent(o.brief, "", "  ")
			pf("The brief:\n\n```json\n%s\n```\n\n", raw)
		}
		pf("**The reply:** %s\n\n", o.note)
		if o.deck != nil {
			pf("Findings: %s\n\n", findings(o.deck))
			if o.repair != "" {
				pf("Repair turn: %s.\n\n", o.repair)
			}
			pf("The deck:\n\n")
			names := make([]string, 0, len(o.deck.GetCards()))
			for _, c := range o.deck.GetCards() {
				names = append(names, fmt.Sprintf("%d %s", c.GetCount(), c.GetName()))
			}
			sort.Strings(names)
			for _, n := range names {
				pf("- %s\n", n)
			}
		}
		if len(o.failures) > 0 {
			pf("\nFailures: %s\n", strings.Join(o.failures, "; "))
		}
	}
	pf("\n")
	run.Markdown(w)
	return pass
}

// recordRows writes one gate row per turn into the run, and the
// information rows beside it (PR-15). A turn's item is base.revision,
// with "a" on the answer turn of an unclear request.
func recordRows(run *evalrun.Run, outcomes []outcome) {
	for _, o := range outcomes {
		item := fmt.Sprintf("%d.%d", o.base.ID, o.rev.ID)
		if o.answered {
			item += "a"
		}
		if o.err != nil {
			run.Gate(item, "pass", 0, o.err.Error())
			continue
		}
		passed := 0.0
		if o.pass() {
			passed = 1
		}
		run.Gate(item, "pass", passed, strings.Join(o.failures, "; "))
		if o.deck != nil {
			run.Info(item, "kept", o.kept, "")
			run.Info(item, "blocks", float64(len(o.blocks)), strings.Join(o.blocks, ", "))
			run.Info(item, "cards", float64(gatekit.CountCards(o.deck)), "")
		}
		repaired := 0.0
		if o.repair != "" {
			repaired = 1
		}
		run.Info(item, "repaired", repaired, o.repair)
		question := 0.0
		if o.brief != nil && o.brief.Question != "" {
			question = 1
		}
		run.Info(item, "question", question, "")
	}
}

func findings(d *mtgv1.Deck) string {
	var out []string
	for _, f := range d.GetValidation().GetFindings() {
		out = append(out, fmt.Sprintf("%s (%s)", f.GetCode(), strings.TrimPrefix(f.GetSeverity().String(), "SEVERITY_")))
	}
	if len(out) == 0 {
		return "none"
	}
	return strings.Join(out, ", ")
}

func cell(s string) string { return strings.ReplaceAll(strings.ReplaceAll(s, "|", "/"), "\n", " ") }
