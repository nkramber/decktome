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

func sixtyJudge(step, why string) llm.Step {
	raw, _ := json.Marshal(map[string]string{"step": step, "why": why})
	return llm.Step{Output: raw}
}

// TestReadImportSixtyStep is D-865 and D-869: a 60-card import takes the
// step of the judge with no guard, and the profile and the grade then
// read the deck at that step. The judge reads the format and the
// sideboard.
func TestReadImportSixtyStep(t *testing.T) {
	b, sc := importBuilder(t, &emptyBook{}, sixtyJudge("fnm", "a tuned burn list"))
	deck := importDeck(mtgv1.FormatId_FORMAT_ID_MODERN)
	deck.Sideboard = []*mtgv1.DeckCard{{OracleId: "o-tithe", Name: "Smothering Tithe", Count: 1}}
	b.ReadImport(context.Background(), deck, nil, &llm.Accumulator{})
	if got := deck.GetPower().GetSixtyStep(); got != mtgv1.SixtyStep_SIXTY_STEP_FNM {
		t.Fatalf("step = %v, want FNM", got)
	}
	if len(sc.Calls) != 1 || deck.GetBracketEstimated() || NeedsPowerRead(deck) {
		t.Errorf("calls = %d, estimated = %v, needs a read = %v", len(sc.Calls), deck.GetBracketEstimated(), NeedsPowerRead(deck))
	}
	if deck.GetProfile() == nil || deck.GetValidation() == nil {
		t.Errorf("profile = %v, validation = %v", deck.GetProfile(), deck.GetValidation())
	}
	if !strings.HasPrefix(deck.GetSummary(), "a tuned burn list") {
		t.Errorf("summary = %q", deck.GetSummary())
	}
	in := sc.Calls[0].Input
	for _, want := range []string{"Format: Modern\n", "98 Plains | Basic Land - Plains\n", "\nSideboard:\n1 Smothering Tithe | Enchantment\n"} {
		if !strings.Contains(in, want) {
			t.Errorf("judge input lacks %q:\n%s", want, in)
		}
	}
}

// TestReadImportSixtyNoStepOnJudgeFailure is D-864: with no judge answer
// the deck holds no step, no profile, and no grade, and the next open
// asks again. The rules still read it (D-846).
func TestReadImportSixtyNoStepOnJudgeFailure(t *testing.T) {
	b, _ := importBuilder(t, &emptyBook{})
	deck := importDeck(mtgv1.FormatId_FORMAT_ID_MODERN)
	b.ReadImport(context.Background(), deck, nil, &llm.Accumulator{})
	if deck.GetPower() != nil || deck.GetProfile() != nil || deck.GetQuality() != nil || deck.GetBracketEstimated() {
		t.Errorf("power = %v, profile = %v, quality = %v", deck.GetPower(), deck.GetProfile(), deck.GetQuality())
	}
	if !NeedsPowerRead(deck) {
		t.Error("a 60-card import with no step asks no new read")
	}
	if deck.GetValidation() == nil || deck.GetValidation().GetPassed() {
		t.Errorf("a 99-card Modern list passed validation: %v", deck.GetValidation())
	}
}

// TestReadImportHouseFormatWord is D-857: the house format names no
// legality to the judge.
func TestReadImportHouseFormatWord(t *testing.T) {
	b, sc := importBuilder(t, &emptyBook{}, sixtyJudge("casual", "x"))
	deck := importDeck(mtgv1.FormatId_FORMAT_ID_HOUSE)
	b.ReadImport(context.Background(), deck, nil, &llm.Accumulator{})
	if in := sc.Calls[0].Input; !strings.HasPrefix(in, "Format: house rules, any card and no ban list\n") {
		t.Errorf("judge input = %q", in)
	}
}

// TestNeedsPowerRead is D-854 and D-864: a floor estimate and a 60-card
// deck with no step ask again, and a deck the app built never does.
func TestNeedsPowerRead(t *testing.T) {
	step := &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: mtgv1.SixtyStep_SIXTY_STEP_CASUAL}}
	bracket := &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}}
	commander, modern := &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER}, &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN}
	for _, c := range []struct {
		name string
		deck *mtgv1.Deck
		want bool
	}{
		{"built deck", &mtgv1.Deck{Format: modern}, false},
		{"commander read", &mtgv1.Deck{Imported: true, Format: commander, Power: bracket}, false},
		{"commander floor", &mtgv1.Deck{Imported: true, Format: commander, Power: bracket, BracketEstimated: true}, true},
		{"sixty read", &mtgv1.Deck{Imported: true, Format: modern, Power: step}, false},
		{"sixty with no step", &mtgv1.Deck{Imported: true, Format: modern}, true},
	} {
		if got := NeedsPowerRead(c.deck); got != c.want {
			t.Errorf("%s: needs a read = %v, want %v", c.name, got, c.want)
		}
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
