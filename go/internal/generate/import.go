package generate

import (
	"context"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/quality"
	"github.com/nkramber/decktome/go/internal/rules"
)

// ReadImport reads a deck that a user brought, as the build reads a deck
// it made (PR-70). The deck carries its cards, its commanders, and its
// format. owned holds the copies of the collection the user picked, and
// nil means no owned mark (D-849).
//
// A Commander deck gets the bracket of the judge. The bracket never sits
// under the floor of its rules (D-850). When the judge fails, the floor
// stands and the deck says so (D-854). A 60-card deck gets the step of
// the judge with no guard (D-865), and no step when the judge fails
// (D-864). The profile and the grade read a deck that holds a power
// (D-869).
//
// Each card stays as the user wrote it, and each finding of the rules
// engine stays on the deck (D-846). The summary is the reason of the
// judge and the sentence of the grade (D-855).
func (b *Builder) ReadImport(ctx context.Context, deck *mtgv1.Deck, owned map[string]int32, acc *llm.Accumulator) {
	commander := deck.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_COMMANDER
	for _, c := range append(append([]*mtgv1.DeckCard(nil), deck.GetCards()...), deck.GetSideboard()...) {
		if card, ok := b.cards.ByOracleID(c.GetOracleId()); ok {
			c.PriceUsd = card.GetPriceUsd()
		}
		c.OwnedCount = owned[c.GetOracleId()]
	}
	markOwned(deck.GetCards())
	markOwned(deck.GetSideboard())
	var why string
	if commander {
		why = b.importBracket(ctx, deck, acc)
	} else {
		why = b.importStep(ctx, deck, acc)
	}
	// Ownership is information on an import: the user owns the list and
	// asks for no build from the collection.
	deck.Validation = b.rules.Validate(rules.Input{
		Deck:         deck,
		PoolRule:     mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		OracleCounts: owned,
		Cards:        b.cards,
	})
	// The profile and the grade read the deck at its power, as for a
	// generated deck, so a deck with no power gets neither (D-869).
	read := deck.GetPower().GetBracket() > 0 || deck.GetPower().GetSixtyStep() != mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED
	if read && b.profiler != nil {
		prof, findings, _ := b.profiler.ReadForbidden(ctx, deck, b.cards)
		deck.Profile = prof
		for _, f := range findings {
			addFinding(deck, f.GetCode(), f.GetSeverity(), f.GetMessage())
		}
		if b.scorer != nil {
			deck.Quality = b.scorer.Score(quality.Input{Deck: deck, Profile: prof, Cards: b.cards})
		}
	}
	deck.Commanders = commanderCards(deck.GetCommanderOracleIds(), b.cards, owned)
	deck.BuyCostUsd = BuyCostWith(deck, b.cards, owned)
	deck.Summary = strings.TrimSpace(why + "\n\n" + quality.Summary(deck.GetQuality()))
}

// importBracket sets the bracket of a Commander import and answers the
// reason of the judge. The floor is the lowest bracket whose rules the
// deck passes. A judge answer under it names a bracket that forbids a
// card of the deck, so the floor wins (D-850). With no judge answer the
// floor stands as an estimate (D-854). With neither, the deck holds no
// bracket, and the next open asks again.
func (b *Builder) importBracket(ctx context.Context, deck *mtgv1.Deck, acc *llm.Accumulator) string {
	var floor int32
	if b.profiler != nil {
		if f, err := b.profiler.Floor(ctx, deck, b.cards); err == nil {
			floor = f
		} else {
			b.log.WarnContext(ctx, "import floor failed", "deck", deck.GetId(), "err", err)
		}
	}
	// The judge counts the combos of the content check and none that it
	// recalls (F-126, D-790), so the profile at the floor goes first.
	if floor > 0 {
		deck.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: floor}}
		deck.Profile, _, _ = b.profiler.ReadForbidden(ctx, deck, b.cards)
	}
	bracket, why := floor, ""
	deck.BracketEstimated = true
	if b.llm != nil {
		j, err := JudgeBracket(ctx, b.llm, deck, b.cards, acc)
		if err == nil {
			bracket, why = max(j.Bracket, floor), j.Why
			deck.BracketEstimated = false
		} else {
			b.log.WarnContext(ctx, "import judge failed", "deck", deck.GetId(), "err", err)
		}
	}
	deck.Power = nil
	if bracket > 0 {
		deck.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: bracket}}
	}
	return why
}

// importStep sets the power step of a 60-card import and answers the
// reason of the judge. No rule guards the step (D-865). With no judge
// answer the deck holds no step, and the next open asks again (D-864).
func (b *Builder) importStep(ctx context.Context, deck *mtgv1.Deck, acc *llm.Accumulator) string {
	deck.Power = nil
	if b.llm == nil {
		return ""
	}
	j, err := JudgeSixtyStep(ctx, b.llm, deck, SixtyFormatWord(deck.GetFormat().GetId()), b.cards, acc)
	if err != nil {
		b.log.WarnContext(ctx, "import step judge failed", "deck", deck.GetId(), "err", err)
		return ""
	}
	deck.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: j.Step}}
	return j.Why
}

// NeedsPowerRead says an imported deck asks the judge again at the next
// open: a Commander bracket that is the floor alone (D-854), or a 60-card
// deck with no step (D-864). A 60-card deck imported before PR-71 holds
// no step, so it gets its read too.
func NeedsPowerRead(deck *mtgv1.Deck) bool {
	if !deck.GetImported() {
		return false
	}
	if deck.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		return deck.GetBracketEstimated()
	}
	return deck.GetPower().GetSixtyStep() == mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED
}
