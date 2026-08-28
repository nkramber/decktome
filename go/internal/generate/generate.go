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
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
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
	// reads the nonbasic ones, and basic lands swap free (A-5 of the
	// 2026-08-28 audit).
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
	// no cap. The pool drops every card above it.
	MaxManaValue float64
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
	// A reader of deck gate run 8 could not tell a miss from a budget.
	Repaired     bool
	RepairReason string
}

// Builder runs the generate and repair calls.
type Builder struct {
	llm   *llm.Client
	rules *rules.Config
	cards rules.CardSource
	log   *slog.Logger
}

// NewBuilder makes a builder. The card source answers the rules engine,
// and it never widens the pool the model may name.
func NewBuilder(c *llm.Client, cfg *rules.Config, cards rules.CardSource, log *slog.Logger) *Builder {
	if log == nil {
		log = slog.Default()
	}
	return &Builder{llm: c, rules: cfg, cards: cards, log: log}
}

// Build writes one deck, checks it, and repairs it once when the check
// refuses it. The deck reaches the caller with its ValidationResult
// attached, whatever the verdict: the user sees the referee's answer
// (roadmap PR-8).
func (b *Builder) Build(ctx context.Context, req Request, acc *llm.Accumulator) (*Result, error) {
	if req.Pool == nil || req.Pool.Size() == 0 {
		return nil, fmt.Errorf("generate: the shortlist is empty")
	}
	out, err := b.call(ctx, llm.RoleGenerate, generateInstructions, b.input(req, nil, nil), req.SessionID, acc)
	if err != nil {
		return nil, err
	}
	res := b.assemble(req, out)
	// One repair turn covers every refusal: a name the shortlist does not
	// hold, a block finding from the engine, and the two warnings that
	// buy a repair on their own. The repair turn reads every one of them,
	// or it cannot fix what it was not told (G-1 of the 2026-08-28 audit).
	if findings := repairable(res.deck.GetValidation()); len(res.misses) > 0 || len(findings) > 0 {
		b.log.Info("the deck was refused, so one repair turn runs",
			"session", req.SessionID, "misses", len(res.misses), "findings", len(findings))
		out2, err := b.call(ctx, llm.RoleRepair, repairInstructions,
			b.input(req, res.misses, findings), req.SessionID, acc)
		if err != nil {
			return nil, err
		}
		reason := repairReason(res.misses, findings)
		res = b.assemble(req, out2)
		res.repaired = true
		res.repairReason = reason
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

// assemble normalizes the model's list, builds the deck, and validates it.
func (b *Builder) assemble(req Request, out *deckOut) pass {
	main := Normalize(req.Pool, out.Cards)
	side := Normalize(req.Pool, out.Sideboard)
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
	// The model can not count its own list reliably, so a small precon
	// shortfall is closed here (D-250). It runs before the engine, so
	// every finding describes the deck the user gets (G-5 of the
	// 2026-08-28 audit).
	swapped := 0
	if req.Precon != "" {
		swapped = swapBackPrecon(deck, req)
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
		checkPreconShare(deck, req)
	}
	if padded > 0 {
		addFinding(deck, CodeBasicsAdded, mtgv1.Severity_SEVERITY_INFO,
			fmt.Sprintf("the list was %s short, so the builder added %s", plural(padded, "card"), plural(padded, "basic land")))
	}
	// The price is a daily estimate and not a rule, so going over budget
	// warns and never blocks (D-236). It does buy the repair turn (D-244).
	if req.BudgetUSD > 0 {
		cost, what := BuyCost(deck), "the cards you must buy"
		if req.BudgetWholeDeck {
			cost, what = DeckCost(deck), "the whole deck"
		}
		if cost > req.BudgetUSD {
			addFinding(deck, CodeOverBudget, mtgv1.Severity_SEVERITY_WARN,
				fmt.Sprintf("%s cost about $%.2f, and the budget is $%.2f", what, cost, req.BudgetUSD))
		}
	}
	// A card the user said to keep must be in the deck. A deck without
	// it answers the user's own instruction with silence (D-242).
	if missing := missingLocked(deck, req); len(missing) > 0 {
		addFinding(deck, CodeLockedCardMissing, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("the deck does not hold %s, which you asked to keep", strings.Join(missing, ", ")))
	}
	if req.ThinCommanderPool {
		addFinding(deck, CodeThinCommanderPool, mtgv1.Severity_SEVERITY_WARN,
			"your library holds no commander for this theme, so the deck was built without one from it")
	}
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
	deck.BuyCostUsd = BuyCost(deck)
	// The summary is prose, and F-26 lives there. The net reads the shape
	// of a rules claim and never its truth.
	lintSummaryInto(deck)
	return pass{deck: deck, misses: append(main.Misses, side.Misses...)}
}

// addFinding appends one finding and keeps Passed true to its meaning.
// The engine set Passed before the builder's own checks ran, and a BLOCK
// added after that shipped under passed = true (G-2 of the 2026-08-28
// audit).
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

// blocking lists the findings that stop a deck.
func blocking(v *mtgv1.ValidationResult) []*mtgv1.Finding {
	var out []*mtgv1.Finding
	for _, f := range v.GetFindings() {
		if f.GetSeverity() == mtgv1.Severity_SEVERITY_BLOCK {
			out = append(out, f)
		}
	}
	return out
}

// repairable lists the findings that buy the repair turn: every BLOCK,
// and the two warnings that do so by decision. Over budget is D-244,
// and a short precon share is D-248. The repair input carries each one.
func repairable(v *mtgv1.ValidationResult) []*mtgv1.Finding {
	var out []*mtgv1.Finding
	for _, f := range v.GetFindings() {
		switch {
		case f.GetSeverity() == mtgv1.Severity_SEVERITY_BLOCK:
			out = append(out, f)
		case f.GetCode() == CodeOverBudget, f.GetCode() == CodePreconShare, f.GetCode() == CodeRevisionOverManaValue:
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
