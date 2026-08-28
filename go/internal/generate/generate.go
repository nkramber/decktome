package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

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
	Power     *mtgv1.PowerLevel
	// Plan is the deck the user asked for, in the user's own terms.
	Plan string
	// Pool holds every card the model may name. A card outside it is a
	// miss, whatever the card index knows (F-13).
	Pool *Pool
	// Commanders are the commander oracle ids, one or two. Empty outside
	// Commander. The session picks them, and the model does not.
	Commanders []string
	// Locked are the oracle ids of cards the user said the deck must
	// keep. The locked row asks for them, and the deck must hold every
	// one (D-70, D-242).
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
	// otherwise. PreconOracleIDs holds its cards (D-218).
	Precon          string
	PreconOracleIDs []string
	// LegalityAsOf is the card-snapshot date the deck is checked against.
	LegalityAsOf string
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
}

// Result is one finished build.
type Result struct {
	Deck *mtgv1.Deck
	// Notes are the user-visible lines for a name that missed twice. The
	// model gets one repair turn, and a second miss becomes a note
	// (roadmap PR-8).
	Notes []string
	// Repaired says the repair turn ran.
	Repaired bool
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
	// One repair turn covers both refusals: a name the shortlist does not
	// hold, and a block finding from the engine.
	if misses, blocks := res.misses, blocking(res.deck.GetValidation()); len(misses) > 0 || len(blocks) > 0 {
		b.log.Info("the deck was refused, so one repair turn runs",
			"session", req.SessionID, "misses", len(misses), "blocks", len(blocks))
		out2, err := b.call(ctx, llm.RoleRepair, repairInstructions,
			b.input(req, misses, blocks), req.SessionID, acc)
		if err != nil {
			return nil, err
		}
		res = b.assemble(req, out2)
		res.repaired = true
	}
	final := &Result{Deck: res.deck, Repaired: res.repaired}
	// A name that missed twice never reaches the deck, and the user reads
	// why it is absent.
	for _, m := range res.misses {
		final.Notes = append(final.Notes, MissNote(m))
	}
	return final, nil
}

// pass is one generate or repair turn, assembled and checked.
type pass struct {
	deck     *mtgv1.Deck
	misses   []Miss
	repaired bool
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
	deck := &mtgv1.Deck{
		Format:             &mtgv1.Format{Id: req.Format},
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
	deck.Validation = b.rules.Validate(rules.Input{
		Deck:         deck,
		PoolRule:     req.PoolRule,
		OracleCounts: req.OracleCounts,
		Cards:        b.cards,
	})
	// The precon share is a build rule and not a rules-engine rule, so it
	// is added here (D-218).
	if req.Precon != "" {
		checkPreconShare(deck, req)
	}
	if padded > 0 {
		deck.Validation.Findings = append(deck.GetValidation().GetFindings(), &mtgv1.Finding{
			Code:     CodeBasicsAdded,
			Severity: mtgv1.Severity_SEVERITY_INFO,
			Message:  fmt.Sprintf("the list was %s short, so the builder added %s", plural(padded, "card"), plural(padded, "basic land")),
		})
	}
	// The price is a daily estimate and not a rule, so going over budget
	// warns and never blocks (D-236).
	if req.BudgetUSD > 0 {
		cost, what := BuyCost(deck), "the cards you must buy"
		if req.BudgetWholeDeck {
			cost, what = DeckCost(deck), "the whole deck"
		}
		if cost > req.BudgetUSD {
			deck.Validation.Findings = append(deck.GetValidation().GetFindings(), &mtgv1.Finding{
				Code:     CodeOverBudget,
				Severity: mtgv1.Severity_SEVERITY_WARN,
				Message: fmt.Sprintf("%s cost about $%.2f, and the budget is $%.2f",
					what, cost, req.BudgetUSD),
			})
		}
	}
	// A card the user said to keep must be in the deck. The locked row
	// asks for it, and a deck without it answers the user's own
	// instruction with silence (D-242).
	if missing := missingLocked(deck, req); len(missing) > 0 {
		deck.Validation.Findings = append(deck.GetValidation().GetFindings(), &mtgv1.Finding{
			Code:     CodeLockedCardMissing,
			Severity: mtgv1.Severity_SEVERITY_BLOCK,
			Message: fmt.Sprintf("the deck does not hold %s, which you asked to keep",
				strings.Join(missing, ", ")),
		})
	}
	if req.ThinCommanderPool {
		deck.Validation.Findings = append(deck.GetValidation().GetFindings(), &mtgv1.Finding{
			Code:     CodeThinCommanderPool,
			Severity: mtgv1.Severity_SEVERITY_WARN,
			Message:  "your library holds no commander for this theme, so the deck was built without one from it",
		})
	}
	// The summary is prose, and F-26 lives there. The net reads the shape
	// of a rules claim and never its truth.
	lintSummaryInto(deck)
	return pass{deck: deck, misses: append(main.Misses, side.Misses...)}
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
