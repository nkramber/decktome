package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/profile"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// Request is one build. The caller fills it from a session whose slots
// are all answered, which is the point the agent service hands over
// (roadmap PR-8).
type Request struct {
	SessionID string
	Format    mtgv1.FormatId
	// HouseRules is what the user means by "anything goes", in the
	// user's own words. The deck's Format carries it (D-3).
	HouseRules string
	Power      *mtgv1.PowerLevel
	// Plan is the deck the user asked for, in the user's own terms.
	Plan string
	// Pool holds every card the model may name. A card outside it is a
	// miss, whatever the card index knows (F-13).
	Pool *Pool
	// Commanders are the commander oracle ids, one or two. Empty outside
	// Commander. The session picks them, and the model does not.
	Commanders []string
	// Locked are the oracle ids of cards the user said the deck must
	// keep. The classifier names them, and the deck must hold every one
	// (D-70, D-242).
	Locked []string
	// PoolRule decides whether ownership is a finding or a mark (D-37).
	PoolRule mtgv1.PoolRule
	// SetCodes are the paper sets the reader limited the deck to, a whole
	// set family (D-376). Empty means no limit. The pool already holds
	// the cards the limit allows, so this is what marks the exceptions:
	// a card the reader named (D-381), a mana card the fill took from
	// outside (D-382), and a basic land the sets do not print (D-378).
	SetCodes []string
	// OracleCounts is the owned count per oracle id, nil with no
	// collection.
	OracleCounts map[string]int32
	// Targets is the wanted count per job, from the corpus role guide.
	Targets map[string]int
	// Roles is the job word per oracle id, from the shortlist. The model
	// reads the job it was given and does not invent one.
	Roles map[string]string
	// Limits is the deck-building limits block the prompt reads.
	Limits string
	// Precon names the precon the user asked to upgrade, and it is empty
	// otherwise. PreconOracleIDs holds its cards (D-218). The share rule
	// reads the nonbasic ones, and basic lands swap free (D-218).
	Precon          string
	PreconOracleIDs []string
	// PreconLands is how many lands the precon runs, copies included. The
	// upgrade prompt names it so the mana base survives the rebuild
	// (D-251).
	PreconLands int
	// LegalityAsOf is the card-snapshot date the deck is checked against.
	LegalityAsOf string
	// DeckID is the id the store reserved. A deck carries its own id, so
	// the caller reserves one before the build (D-245).
	DeckID string
	// OnPhase reports where the build stands, so the chat can light a
	// step (D-435). Nil reports nothing. The builder calls it before the
	// generate call, before each check, and before the repair turn.
	OnPhase func(mtgv1.BuildPhase)
	// Name is what the user sees the deck called.
	Name string
	// Now stamps the deck. Tests give a fixed clock.
	Now func() time.Time
	// BudgetUSD is what the user allowed, 0 when they named no number.
	// BudgetWholeDeck says the cap covers every card and not the cards the
	// user must buy, which the budget-scope row asks (D-77, D-238).
	BudgetUSD       float64
	BudgetWholeDeck bool
	// ThinCommanderPool says the library holds no commander for the theme
	// in an owned mode. The retired weak-pool row asked about this and
	// could never reach the user who needed it, so the deck reports it
	// (D-232).
	ThinCommanderPool bool
	// Revision is set when the user asked for a change after a build
	// (PR-12B, D-283). The model gets the base deck and the brief, and
	// keeps every card the brief does not touch.
	Revision *Revision
}

// Revision is the brief of one change request against a built deck.
type Revision struct {
	// BaseDeckID is the deck the user read. The new deck records it.
	BaseDeckID string
	// Base is the card list the user read, with counts and jobs.
	Base []*mtgv1.DeckCard
	// Instructions restate the request for the generator, in short lines.
	Instructions []string
	// Remove names the cards the user wants out. The pool never holds
	// them, and a deck that keeps one is refused.
	Remove []string
	// Keep names the cards the user wants in. A deck without one is
	// refused.
	Keep []string
	// MaxManaValue caps the mana value of every nonland card. Zero means
	// no cap. The pool drops every card above it, the Keep names and the
	// Exempt ids excepted.
	MaxManaValue float64
	// Exempt lists the oracle ids the cap never drops: the locked cards
	// and the commanders. A locked card the pool dropped blocks as
	// missing, and the model can not put back what it can not name
	// (D-242).
	Exempt []string
	// SwapBasics is how many basic lands the deck must replace with
	// nonbasic lands, and LandKinds says what kind. Zero means no land
	// swap. A deck that adds fewer is refused (D-448).
	SwapBasics int
	LandKinds  string
}

// Result is one finished build.
type Result struct {
	Deck *mtgv1.Deck
	// Notes are the user-visible lines for a name that missed twice. The
	// model gets one repair turn, and a second miss becomes a note
	// (roadmap PR-8).
	Notes []string
	// Repaired says the repair turn ran. RepairReason says why, as a short
	// line the gate document prints: the miss count and the finding codes.
	Repaired     bool
	RepairReason string
}

// Builder runs the generate and repair calls.
type Builder struct {
	llm      *llm.Client
	rules    *rules.Config
	cards    rules.CardSource
	profiler *profile.Profiler
	log      *slog.Logger
}

// BuilderOption configures a Builder.
type BuilderOption func(*Builder)

// WithProfiler wires the bracket profile (PR-14A). Every built deck then
// carries its profile, and an off-band feature or a content violation
// buys the repair turn. Without it a deck carries no profile.
func WithProfiler(p *profile.Profiler) BuilderOption {
	return func(b *Builder) { b.profiler = p }
}

// NewBuilder makes a builder. The card source answers the rules engine,
// and it never widens the pool the model may name.
func NewBuilder(c *llm.Client, cfg *rules.Config, cards rules.CardSource, log *slog.Logger, opts ...BuilderOption) *Builder {
	if log == nil {
		log = slog.Default()
	}
	b := &Builder{llm: c, rules: cfg, cards: cards, log: log}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

// Build writes one deck, checks it, and repairs it once when the check
// refuses it. The deck reaches the caller with its ValidationResult
// attached, whatever the verdict: the user sees the referee's answer
// (roadmap PR-8).
func (b *Builder) Build(ctx context.Context, req Request, acc *llm.Accumulator) (*Result, error) {
	if req.Pool == nil || req.Pool.Size() == 0 {
		return nil, fmt.Errorf("generate: the shortlist is empty")
	}
	req.phase(mtgv1.BuildPhase_BUILD_PHASE_BUILDING)
	cut := b.cutShortlist(ctx, &req)
	out, err := b.call(ctx, llm.RoleGenerate, generateInstructions, b.input(req, nil, nil), req.SessionID, acc)
	if err != nil {
		return nil, err
	}
	req.phase(mtgv1.BuildPhase_BUILD_PHASE_CHECKING)
	res := b.assemble(ctx, req, out)
	// One repair turn covers every refusal: a name the shortlist does not
	// hold, a block finding from the engine, and the warnings that buy a
	// repair on their own. The repair turn reads every one of them, or
	// it cannot fix what it was not told (D-244, D-248).
	//
	// A profile finding alone earns one more pass when the first repair
	// left the deck off band. That is the bounded second pass of PR-14A,
	// and the count is MaxRepairs.
	for turn := 1; turn <= MaxRepairs; turn++ {
		findings := repairable(res.deck.GetValidation())
		if len(res.misses) == 0 && len(findings) == 0 {
			break
		}
		if turn > 1 && (len(res.misses) > 0 || !profileOnly(findings)) {
			break
		}
		b.log.Info("the deck was refused, so a repair turn runs",
			"session", req.SessionID, "turn", turn, "misses", len(res.misses), "findings", len(findings))
		req.phase(mtgv1.BuildPhase_BUILD_PHASE_REPAIRING)
		out2, err := b.call(ctx, llm.RoleRepair, repairInstructions,
			b.input(req, res.misses, findings), req.SessionID, acc)
		if err != nil {
			return nil, err
		}
		reason := repairReason(res.misses, findings)
		req.phase(mtgv1.BuildPhase_BUILD_PHASE_CHECKING)
		res = b.assemble(ctx, req, out2)
		res.repaired = true
		if res.repairReason != "" {
			reason = res.repairReason + "; then " + reason
		}
		res.repairReason = reason
	}
	if len(cut) > 0 {
		addFinding(res.deck, CodeShortlistCut, mtgv1.Severity_SEVERITY_INFO,
			fmt.Sprintf("bracket %d does not allow %s, so the shortlist left %s out: %s",
				req.Power.GetBracket(), these(len(cut)), these(len(cut)), strings.Join(cut, ", ")))
	}
	final := &Result{Deck: res.deck, Repaired: res.repaired, RepairReason: res.repairReason}
	// A name that missed twice never reaches the deck, and the user reads
	// why it is absent.
	for _, m := range res.misses {
		final.Notes = append(final.Notes, MissNote(m))
	}
	return final, nil
}

// pass is one generate or repair turn, assembled and checked.
type pass struct {
	deck         *mtgv1.Deck
	misses       []Miss
	repaired     bool
	repairReason string
}

// CodeShortlistCut reports the cards Commander Spellbook flagged for the
// bracket before the build, which the shortlist then left out (D-468).
// It is an INFO: the deck is what the bracket allows, and the reader
// should know what it never saw.
const CodeShortlistCut = "shortlist_cut"

// cutShortlist asks the profiler which shortlist cards the bracket
// forbids, and drops them from the pool (D-468). A locked card and a
// commander stay: the user named them (D-242). A failed call drops
// nothing and is logged, and the check after the build still runs.
func (b *Builder) cutShortlist(ctx context.Context, req *Request) []string {
	if b.profiler == nil || req.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER || req.Power.GetBracket() == 0 {
		return nil
	}
	keep := map[string]bool{}
	for _, id := range append(append([]string(nil), req.Commanders...), req.Locked...) {
		keep[id] = true
	}
	var names, commanders []string
	for _, id := range req.Commanders {
		if c, ok := b.cards.ByOracleID(id); ok {
			commanders = append(commanders, c.GetName())
		}
	}
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok || keep[c.GetOracleId()] || candidates.IsBasicLand(c) {
			continue
		}
		names = append(names, c.GetName())
	}
	drop, err := b.profiler.CutShortlist(ctx, req.Format, req.Power.GetBracket(), commanders, names)
	if err != nil {
		b.log.Warn("the shortlist was not read for the bracket, so nothing left it", "session", req.SessionID, "err", err)
		return nil
	}
	if len(drop) == 0 {
		return nil
	}
	// Only a card that was on the pool and not kept has left, and only
	// those are reported.
	gone := map[string]bool{}
	var removed []string
	for _, n := range drop {
		if c, ok := req.Pool.Card(n); ok && !keep[c.GetOracleId()] {
			gone[c.GetOracleId()] = true
			removed = append(removed, c.GetName())
		}
	}
	if len(removed) == 0 {
		return nil
	}
	req.Pool = req.Pool.Filter(func(c *mtgv1.Card) bool { return !gone[c.GetOracleId()] })
	return removed
}

// MaxRepairs is the most repair turns one build runs. The first covers
// every refusal, and the second runs only for a profile finding the
// first left behind (PR-14A).
const MaxRepairs = 2

// profileOnly reports whether every finding is one of the profile's.
func profileOnly(findings []*mtgv1.Finding) bool {
	for _, f := range findings {
		switch f.GetCode() {
		case profile.CodeOffBand, profile.CodeMassLandDenial, profile.CodeExtraTurns, profile.CodeTwoCardCombo:
		default:
			return false
		}
	}
	return len(findings) > 0
}

// assemble normalizes the model's list, builds the deck, and validates it.
func (b *Builder) assemble(ctx context.Context, req Request, out *deckOut) pass {
	main := Normalize(req.Pool, out.Cards)
	side := Normalize(req.Pool, out.Sideboard)
	// The copy limit spans both lists, so the owned flag reads the count
	// of an oracle id across both (D-37).
	markOwned(append(append([]*mtgv1.DeckCard(nil), main.Cards...), side.Cards...))
	// Only Commander has a command zone. A 60-card session whose
	// classifier reported a commander name would otherwise build a deck
	// with one, and the engine refuses it as not legal in the format
	// (D-233).
	commanders := req.Commanders
	if req.Format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		commanders = nil
	}
	now := req.Now
	if now == nil {
		now = time.Now
	}
	deck := &mtgv1.Deck{
		Id:                 req.DeckID,
		Name:               req.Name,
		CreatedAt:          timestamppb.New(now().UTC()),
		Format:             &mtgv1.Format{Id: req.Format, HouseRules: req.HouseRules},
		Power:              req.Power,
		Summary:            strings.TrimSpace(out.Summary),
		CommanderOracleIds: commanders,
		Cards:              main.Cards,
		Sideboard:          side.Cards,
		SessionId:          req.SessionID,
		LegalityAsOf:       req.LegalityAsOf,
	}
	// A deck one or two cards short is a counting slip, and a basic land
	// is always a legal answer. The builder finishes the list before the
	// engine reads it, rather than return a deck the engine must refuse
	// (D-225).
	padded := padWithBasics(deck, req)
	// The same counting slip, the other way. A deck one or two cards
	// over is blocked whole by the engine, and the reader gets nothing
	// (D-391).
	trimmed := trimToSize(deck, req)
	// The model can not count its own list reliably, so a small precon
	// shortfall is closed here (D-250). It runs before the engine, so
	// every finding describes the deck the user gets.
	swapped := 0
	if req.Precon != "" {
		swapped = swapBackPrecon(deck, req, b.cards)
	}
	deck.Validation = b.rules.Validate(rules.Input{
		Deck:         deck,
		PoolRule:     req.PoolRule,
		OracleCounts: req.OracleCounts,
		Cards:        b.cards,
	})
	if swapped > 0 {
		addFinding(deck, CodePreconCardsRestored, mtgv1.Severity_SEVERITY_INFO,
			fmt.Sprintf("the deck was %s short of the %s precon share, so the builder put %s back",
				plural(swapped, "card"), req.Precon, plural(swapped, "card")))
	}
	// The precon share is a build rule and not a rules-engine rule, so
	// it is added here (D-218).
	if req.Precon != "" {
		checkPreconShare(deck, req, b.cards)
	}
	// The bracket profile reads the finished deck. Its findings are
	// warnings, and they buy the repair turn (PR-14A).
	if b.profiler != nil {
		prof, findings := b.profiler.Read(ctx, deck, b.cards)
		deck.Profile = prof
		for _, f := range findings {
			addFinding(deck, f.GetCode(), f.GetSeverity(), f.GetMessage())
		}
	}
	if padded > 0 {
		addFinding(deck, CodeBasicsAdded, mtgv1.Severity_SEVERITY_INFO,
			fmt.Sprintf("the list was %s short, so the builder added %s", plural(padded, "card"), plural(padded, "basic land")))
	}
	if len(trimmed) > 0 {
		addFinding(deck, CodeCardsTrimmed, mtgv1.Severity_SEVERITY_INFO,
			fmt.Sprintf("the list was %s over, so the builder cut %s: %s",
				plural(len(trimmed), "card"), these(len(trimmed)), strings.Join(trimmed, ", ")))
	}
	// The price is a daily estimate and not a rule, so going over budget
	// warns and never blocks (D-236). It does buy the repair turn (D-244).
	if req.BudgetUSD > 0 {
		cost, what := BuyCostWith(deck, b.cards, req.OracleCounts), "the cards you must buy"
		if req.BudgetWholeDeck {
			cost, what = DeckCostWith(deck, b.cards), "the whole deck"
		}
		if cost > req.BudgetUSD {
			addFinding(deck, CodeOverBudget, mtgv1.Severity_SEVERITY_WARN,
				fmt.Sprintf("%s cost about $%.2f, and the budget is $%.2f", what, cost, req.BudgetUSD))
		}
	}
	// A card the user said to keep must be in the deck. A deck without
	// it answers the user's own instruction with silence (D-242).
	if missing := missingLocked(deck, req, b.cards); len(missing) > 0 {
		addFinding(deck, CodeLockedCardMissing, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("the deck does not hold %s, which you asked to keep", strings.Join(missing, ", ")))
	}
	if req.ThinCommanderPool {
		addFinding(deck, CodeThinCommanderPool, mtgv1.Severity_SEVERITY_WARN,
			"your library holds no commander for this theme, so the deck was built without one from it")
	}
	// The set limit is a build rule and not a rule of the game, so it is
	// marked here and never blocks (D-373, D-383).
	markOutsideSets(deck, req, b.cards)
	// A revision must do what the brief says. The pool already dropped
	// the removed cards and the cards over the cap, so these fire only
	// on a kept card the model left out, or a pool the brief could not
	// filter (PR-12B).
	if req.Revision != nil {
		deck.RevisedFromDeckId = req.Revision.BaseDeckID
		for _, f := range CheckRevision(deck, req.Revision, b.cards) {
			addFinding(deck, f.GetCode(), f.GetSeverity(), f.GetMessage())
		}
	}
	// The deck carries what it costs, so a reader needs no card index to
	// see it (D-245).
	deck.BuyCostUsd = BuyCostWith(deck, b.cards, req.OracleCounts)
	// The summary is prose, and F-26 lives there. The net reads the shape
	// of a rules claim and never its truth.
	lintSummaryInto(deck)
	return pass{deck: deck, misses: append(main.Misses, side.Misses...)}
}

// addFinding appends one finding and keeps Passed true to its meaning: a
// BLOCK added after the engine ran clears it.
func addFinding(deck *mtgv1.Deck, code string, sev mtgv1.Severity, msg string) {
	v := deck.GetValidation()
	if v == nil {
		v = &mtgv1.ValidationResult{Passed: true}
		deck.Validation = v
	}
	v.Findings = append(v.Findings, &mtgv1.Finding{Code: code, Severity: sev, Message: msg})
	if sev == mtgv1.Severity_SEVERITY_BLOCK {
		v.Passed = false
	}
}

// call runs one model turn and reads its answer.
func (b *Builder) call(ctx context.Context, role llm.Role, instructions, input, key string, acc *llm.Accumulator) (*deckOut, error) {
	res, err := b.llm.Complete(ctx, role, llm.Request{
		Instructions: instructions,
		Input:        input,
		SchemaName:   "deck",
		Schema:       json.RawMessage(deckSchema),
		// One cache key per session, so a repair turn and a re-roll read
		// the same prefix (roadmap PR-8).
		CacheKey: key,
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("generate: %s: %w", role, err)
	}
	var out deckOut
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("generate: %s output: %w", role, err)
	}
	return &out, nil
}

// repairable lists the findings that buy the repair turn: every BLOCK,
// and the warnings that do so by decision. Over budget is D-244, a
// short precon share is D-248, and a profile finding is PR-14A. The
// repair input carries each one.
func repairable(v *mtgv1.ValidationResult) []*mtgv1.Finding {
	var out []*mtgv1.Finding
	for _, f := range v.GetFindings() {
		switch {
		case f.GetSeverity() == mtgv1.Severity_SEVERITY_BLOCK:
			out = append(out, f)
		case f.GetCode() == CodeOverBudget, f.GetCode() == CodePreconShare, f.GetCode() == CodeRevisionOverManaValue:
			out = append(out, f)
		case f.GetCode() == profile.CodeOffBand, f.GetCode() == profile.CodeMassLandDenial,
			f.GetCode() == profile.CodeExtraTurns, f.GetCode() == profile.CodeTwoCardCombo:
			out = append(out, f)
		}
	}
	return out
}

// repairReason names what bought the repair turn: the miss count and
// the codes of the findings the model was shown.
func repairReason(misses []Miss, findings []*mtgv1.Finding) string {
	var parts []string
	if len(misses) > 0 {
		parts = append(parts, fmt.Sprintf("%d missed names", len(misses)))
	}
	for _, f := range findings {
		parts = append(parts, f.GetCode())
	}
	return strings.Join(parts, ", ")
}

// phase reports a build step to the caller that asked for it (D-435).
func (r Request) phase(p mtgv1.BuildPhase) {
	if r.OnPhase != nil {
		r.OnPhase(p)
	}
}
