package generate

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
	"github.com/nkramber/decktome/go/internal/spellbook"
)

// emptyBook answers every content check with no combo and no flag.
type emptyBook struct{ calls int }

func (e *emptyBook) EstimateBracket(context.Context, []string, []string) (*spellbook.Result, error) {
	e.calls++
	return &spellbook.Result{BracketTag: "C"}, nil
}

func importCards() source {
	legal := map[string]mtgv1.LegalityStatus{
		"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL, "modern": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL,
	}
	w := []mtgv1.Color{mtgv1.Color_COLOR_W}
	return source{
		"o-lead":   {OracleId: "o-lead", Name: "Leader", CanBeCommander: true, ColorIdentity: w, TypeLine: "Legendary Creature", Legalities: legal, PriceUsd: 2},
		"o-tithe":  {OracleId: "o-tithe", Name: "Smothering Tithe", GameChanger: true, ColorIdentity: w, TypeLine: "Enchantment", Legalities: legal, PriceUsd: 20},
		"o-plains": {OracleId: "o-plains", Name: "Plains", TypeLine: "Basic Land - Plains", Legalities: legal},
	}
}

func importBuilder(t *testing.T, book profile.Classifier, steps ...llm.Step) (*Builder, *llm.Script) {
	t.Helper()
	sc := llm.NewScript(steps...)
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	prof, err := profile.New(cfg, func() *cards.TagIndex { return nil }, book)
	if err != nil {
		t.Fatal(err)
	}
	prof.SetHands(50)
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewBuilder(c, cfg, importCards(), log, WithProfiler(prof)), sc
}

func importDeck(format mtgv1.FormatId) *mtgv1.Deck {
	d := &mtgv1.Deck{
		Id: "d-1", Imported: true, Format: &mtgv1.Format{Id: format},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-tithe", Name: "Smothering Tithe", Count: 1},
			{OracleId: "o-plains", Name: "Plains", Count: 98},
		},
	}
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		d.CommanderOracleIds = []string{"o-lead"}
	}
	return d
}

func judgeStep(bracket, why string) llm.Step {
	raw, _ := json.Marshal(map[string]string{"bracket": bracket, "why": why})
	return llm.Step{Output: raw}
}

// TestReadImportKeepsTheFloor is D-850: a judge answer under the floor
// of the rules names a bracket that forbids a card, so the floor wins.
// One Game Changer puts the floor at 3.
func TestReadImportKeepsTheFloor(t *testing.T) {
	b, sc := importBuilder(t, &emptyBook{}, judgeStep("2", "a casual lifegain deck"))
	deck := importDeck(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	b.ReadImport(context.Background(), deck, nil, &llm.Accumulator{})
	if got := deck.GetPower().GetBracket(); got != 3 || deck.GetBracketEstimated() {
		t.Fatalf("bracket = %d, estimated = %v, want 3 and false", got, deck.GetBracketEstimated())
	}
	if len(sc.Calls) != 1 {
		t.Errorf("judge calls = %d", len(sc.Calls))
	}
	if deck.GetProfile().GetBracket() != 3 || deck.GetValidation() == nil {
		t.Errorf("profile = %v, validation = %v", deck.GetProfile(), deck.GetValidation())
	}
	if !strings.HasPrefix(deck.GetSummary(), "a casual lifegain deck") {
		t.Errorf("summary = %q", deck.GetSummary())
	}
	if in := sc.Calls[0].Input; !strings.Contains(in, "Combos:") {
		t.Errorf("the judge read no content check: %q", in)
	}
}

// TestReadImportTakesAHigherJudge reads a judge answer over the floor as
// the bracket.
func TestReadImportTakesAHigherJudge(t *testing.T) {
	b, _ := importBuilder(t, &emptyBook{}, judgeStep("4", "fast mana"))
	deck := importDeck(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	b.ReadImport(context.Background(), deck, nil, &llm.Accumulator{})
	if got := deck.GetPower().GetBracket(); got != 4 {
		t.Errorf("bracket = %d, want 4", got)
	}
}

// TestReadImportFloorOnJudgeFailure is D-854: with no judge answer the
// floor stands, and the deck marks it as an estimate.
func TestReadImportFloorOnJudgeFailure(t *testing.T) {
	b, _ := importBuilder(t, &emptyBook{})
	deck := importDeck(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	b.ReadImport(context.Background(), deck, nil, &llm.Accumulator{})
	if got := deck.GetPower().GetBracket(); got != 3 || !deck.GetBracketEstimated() {
		t.Errorf("bracket = %d, estimated = %v, want 3 and true", got, deck.GetBracketEstimated())
	}
}

// TestReadImportSixtyCards is D-859: a 60-card import calls no model and
// holds no power step and no profile until PR-71. The rules still read
// it (D-846).
func TestReadImportSixtyCards(t *testing.T) {
	book := &emptyBook{}
	b, sc := importBuilder(t, book)
	deck := importDeck(mtgv1.FormatId_FORMAT_ID_MODERN)
	b.ReadImport(context.Background(), deck, nil, &llm.Accumulator{})
	if len(sc.Calls) != 0 || book.calls != 0 {
		t.Errorf("model calls = %d, content checks = %d", len(sc.Calls), book.calls)
	}
	if deck.GetPower() != nil || deck.GetProfile() != nil || deck.GetBracketEstimated() {
		t.Errorf("power = %v, profile = %v", deck.GetPower(), deck.GetProfile())
	}
	if deck.GetValidation() == nil || deck.GetValidation().GetPassed() {
		t.Errorf("a 99-card Modern list passed validation: %v", deck.GetValidation())
	}
}

// TestReadImportMarksOwnedCards is D-849: the picked collection marks the
// owned commander and prices the rest.
func TestReadImportMarksOwnedCards(t *testing.T) {
	b, _ := importBuilder(t, &emptyBook{}, judgeStep("3", "x"))
	deck := importDeck(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	b.ReadImport(context.Background(), deck, map[string]int32{"o-lead": 1}, &llm.Accumulator{})
	if c := deck.GetCommanders(); len(c) != 1 || !c[0].GetOwned() {
		t.Fatalf("commanders = %v", c)
	}
	if deck.GetBuyCostUsd() != 20 {
		t.Errorf("buy cost = %v, want 20", deck.GetBuyCostUsd())
	}
}
