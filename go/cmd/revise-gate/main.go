// Command revise-gate runs the PR-12B revision prompts and writes the
// gate document. Each base deck is built once, then each revision runs
// from that base: the revise call, the revised build, and the checks.
//
// The bars come from the roadmap entry PR-12B. Every revised deck keeps
// the untouched cards, holds every limit of the brief, and passes the
// engine. An unclear request gets a question and no build.
//
// CAUTION: this calls a real provider and it costs money. REVISE_GATE=1
// is required, so it can not run by accident. Ask the owner before every
// run. `make revise-gate` writes to REVISE_GATE_OUT and refuses a file
// that already holds a verdict (D-65).
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/revise"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

//go:embed prompts.json
var promptsJSON []byte

type base struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Format    string     `json:"format"`
	Power     string     `json:"power"`
	Bracket   int32      `json:"bracket"`
	Colors    []string   `json:"colors"`
	Commander string     `json:"commander"`
	Theme     string     `json:"theme"`
	Plan      string     `json:"plan"`
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
	// Note is for the reader of the prompt file.
	Note string `json:"note"`
}

// keepBar is the share of untouched base names a revised deck must keep.
// A revision that asks for swaps moves some, so the bar is not 100.
const keepBar = 0.8

type outcome struct {
	base     base
	rev      revision
	message  string
	brief    *revise.Brief
	deck     *mtgv1.Deck
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
	collPath := flag.String("collection", "", "a ManaBox CSV, unused today and kept for parity")
	flag.Parse()
	_ = collPath

	var file struct {
		Bases []base `json:"bases"`
	}
	if err := json.Unmarshal(promptsJSON, &file); err != nil {
		return fmt.Errorf("prompts: %w", err)
	}
	if err := gatekit.SpendGuard("REVISE_GATE"); err != nil {
		return err
	}
	quiet := gatekit.Quiet()
	ctx := context.Background()
	idx, err := gatekit.LoadSnapshot(ctx, quiet)
	if err != nil {
		return err
	}
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	rcfg, err := rules.Load()
	if err != nil {
		return err
	}
	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return err
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)
	b := generate.NewBuilder(client, rcfg, idx, quiet)

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
		fmt.Fprintf(os.Stderr, "  %d cards, %d blocks\n", countCards(baseDeck), len(blocks(baseDeck)))
		for _, r := range bs.Revisions {
			o := runRevision(ctx, client, b, idx, bs, r, baseDeck, pool, list, commanderIDs, acc)
			outcomes = append(outcomes, o)
			fmt.Fprintf(os.Stderr, "  %d. %s\n", r.ID, status(o))
		}
	}
	report(os.Stdout, outcomes, acc, idx, time.Since(start))
	return nil
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

func powerOf(bs base) *mtgv1.PowerLevel {
	if bs.Bracket > 0 {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: bs.Bracket}}
	}
	steps := map[string]mtgv1.SixtyStep{
		"casual": mtgv1.SixtyStep_SIXTY_STEP_CASUAL, "fnm": mtgv1.SixtyStep_SIXTY_STEP_FNM, "tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
	}
	if s, ok := steps[strings.ToLower(bs.Power)]; ok {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: s}}
	}
	return nil
}

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
	list, err := cb.Build(idx, candidates.Request{
		Format: format, Colors: colors, Theme: bs.Theme, CommanderOracleIDs: commanderIDs,
		PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD, Bracket: bs.Bracket,
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
	}
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
// build with the brief, then the checks against the bars.
func runRevision(ctx context.Context, client *llm.Client, b *generate.Builder, idx *cards.Index, bs base, r revision,
	baseDeck *mtgv1.Deck, basePool *generate.Pool, list *candidates.List, commanderIDs []string, acc *llm.Accumulator) outcome {
	o := outcome{base: bs, rev: r}
	highest := highestNonland(baseDeck, idx)
	o.message = strings.ReplaceAll(r.Message, "{{HIGHEST}}", highest)
	brief, err := revise.Call(ctx, client, revise.Input{
		SessionID: fmt.Sprintf("revise-gate-%d", bs.ID), Message: o.message, Deck: baseDeck,
		Format: generate.FormatWord(formatOf(bs)), Power: powerWord(bs), Cards: idx,
	}, acc)
	if err != nil {
		o.err = err
		return o
	}
	o.brief = brief
	if r.Unclear {
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
		return o
	}
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
	rev := &generate.Revision{
		BaseDeckID: baseDeck.GetId(), Base: baseDeck.GetCards(), Instructions: brief.Changes,
		Remove: brief.Remove, Keep: brief.Keep, MaxManaValue: brief.MaxManaValue,
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
	res, err := b.Build(ctx, request(bs, idx, pool, list, commanderIDs, rev), acc)
	if err != nil {
		o.err = err
		return o
	}
	o.deck = res.Deck
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
	o.kept = keptShare(baseDeck, res.Deck, rev, idx)
	if o.kept < keepBar {
		o.failures = append(o.failures, fmt.Sprintf("kept %.0f%% of the untouched cards, the bar is %.0f%%", o.kept*100, keepBar*100))
	}
	return o
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

func countCards(d *mtgv1.Deck) int {
	n := 0
	for _, c := range d.GetCards() {
		n += int(c.GetCount())
	}
	return n
}

func blocks(d *mtgv1.Deck) []string {
	var out []string
	for _, f := range d.GetValidation().GetFindings() {
		if f.GetSeverity() == mtgv1.Severity_SEVERITY_BLOCK {
			out = append(out, f.GetCode())
		}
	}
	return out
}

func report(w io.Writer, outcomes []outcome, acc *llm.Accumulator, idx *cards.Index, took time.Duration) {
	// The document is the record, and a write error on stdout ends the
	// process in any case.
	pf := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	passed := 0
	for _, o := range outcomes {
		if o.pass() {
			passed++
		}
	}
	verdict := "PASS"
	if passed < len(outcomes) {
		verdict = "FAIL"
	}
	pf("# PR-12B revise gate\n\nRun date: %s. Snapshot: %s.\n\n", time.Now().UTC().Format("2006-01-02"), idx.AsOf.Format("2006-01-02"))
	pf("Verdict: %s. %d of %d revisions met their bar. The bars: an unclear request gets a question or a decline with a reason, a clear request gets a revised deck with no block finding, the deck holds every cap and every removal the message names, and it keeps at least %.0f percent of the untouched cards.\n\n",
		verdict, passed, len(outcomes), keepBar*100)
	r := acc.Report()
	cost := "unpriced"
	if r.CostUSD != nil {
		cost = fmt.Sprintf("$%.4f", *r.CostUSD)
	}
	pf("Cost: %d calls, %s. Time: %s.\n\n", r.Calls, cost, took.Round(time.Second))
	pf("| Base | Revision | Message | Outcome | Kept | Result |\n|---|---|---|---|---|---|\n")
	for _, o := range outcomes {
		out := "no build"
		if o.deck != nil {
			out = fmt.Sprintf("%d cards", countCards(o.deck))
		}
		if o.brief != nil && o.brief.Question != "" {
			out = "question"
		}
		kept := ""
		if o.deck != nil {
			kept = fmt.Sprintf("%.0f%%", o.kept*100)
		}
		pf("| %d | %d | %s | %s | %s | %s |\n", o.base.ID, o.rev.ID, cell(o.message), out, kept, cell(status(o)))
	}
	for _, o := range outcomes {
		pf("\n## Revision %d, base %d: %s\n\n", o.rev.ID, o.base.ID, o.base.Name)
		pf("**The user:** %s\n\n", o.message)
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
