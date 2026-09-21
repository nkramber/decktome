package generate

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
)

func planClient(t *testing.T, output string) *llm.Client {
	t.Helper()
	sc := llm.NewScript(llm.Step{Output: json.RawMessage(output)})
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	return c
}

// TestJudgePlanReadsTheFourGrades: the four words parse, the values
// read no as 0, partly as 0.5, and yes as 1, and the score is the mean.
func TestJudgePlanReadsTheFourGrades(t *testing.T) {
	c := planClient(t, `{
		"plan_coherent": {"grade": "yes", "why": "one plan"},
		"theme_fit": {"grade": "partly", "why": "half the theme"},
		"useful_as_built": {"grade": "no", "why": "no lands"},
		"summary_honest": {"grade": "yes", "why": "plain"}
	}`)
	deck := &mtgv1.Deck{Summary: "a deck", Cards: []*mtgv1.DeckCard{{Name: "Sol Ring", Count: 1}}}
	j, err := JudgePlan(context.Background(), c, "a lifegain deck", deck, source{}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if j.ThemeFit.Value() != 0.5 || j.UsefulAsBuilt.Value() != 0 || j.Grade("summary_honest").Why != "plain" {
		t.Errorf("judgement = %+v", j)
	}
	if j.Score() != 0.625 {
		t.Errorf("score = %v, want the mean of 1, 0.5, 0, 1", j.Score())
	}
	if len(PlanFields) != 4 || PlanFields[0] != "plan_coherent" {
		t.Errorf("fields = %v", PlanFields)
	}
}

// TestJudgePlanRefusesAWordOutsideTheScale: a grade the rubric does not
// hold is an error and never a zero. The client checks the schema and
// retries once, so the script serves the bad answer twice.
func TestJudgePlanRefusesAWordOutsideTheScale(t *testing.T) {
	bad := `{
		"plan_coherent": {"grade": "excellent", "why": "x"},
		"theme_fit": {"grade": "yes", "why": "x"},
		"useful_as_built": {"grade": "yes", "why": "x"},
		"summary_honest": {"grade": "yes", "why": "x"}
	}`
	sc := llm.NewScript(llm.Step{Output: json.RawMessage(bad)}, llm.Step{Output: json.RawMessage(bad)})
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	_, err = JudgePlan(context.Background(), c, "a deck", &mtgv1.Deck{}, source{}, nil, nil)
	if err == nil || !strings.Contains(err.Error(), "judge plan") {
		t.Errorf("err = %v, want a judge plan error", err)
	}
	if j := (PlanJudgement{PlanCoherent: PlanGrade{Grade: "excellent"}}); j.PlanCoherent.Value() != 0 {
		t.Error("a word outside the scale reads as 0, never more")
	}
}

// TestPlanReasonEmpty: a blank reason and the word "placeholder" count as
// no reason, and a sentence counts as one.
func TestPlanReasonEmpty(t *testing.T) {
	for _, why := range []string{"", "  ", "placeholder", "Placeholder"} {
		if !(PlanGrade{Grade: "yes", Why: why}).ReasonEmpty() {
			t.Errorf("%q must read as no reason", why)
		}
	}
	if (PlanGrade{Grade: "no", Why: "the deck has no lands"}).ReasonEmpty() {
		t.Error("a sentence is a reason")
	}
	j := PlanJudgement{
		PlanCoherent:  PlanGrade{Grade: "yes", Why: "one plan"},
		ThemeFit:      PlanGrade{Grade: "yes", Why: "placeholder"},
		UsefulAsBuilt: PlanGrade{Grade: "yes", Why: ""},
		SummaryHonest: PlanGrade{Grade: "yes", Why: "plain"},
	}
	if n := j.EmptyReasons(); n != 2 {
		t.Errorf("EmptyReasons = %d, want 2", n)
	}
}

// TestJudgePlanReadsTheCardFacts is F-39: the plan judge reads the mana
// cost and the type line of each card, and the color identity of the
// commander, from the card data. The cards are test fixtures and name no
// real card.
func TestJudgePlanReadsTheCardFacts(t *testing.T) {
	removal := mtgv1.CardRole_CARD_ROLE_REMOVAL
	w, b := mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B
	cards := source{
		"o-lead":  {OracleId: "o-lead", Name: "A Two-Color Leader", ManaCost: "{2}{W}{B}", TypeLine: "Legendary Creature — Human Noble", ColorIdentity: []mtgv1.Color{b, w}},
		"o-bolt":  {OracleId: "o-bolt", Name: "A Cheap Removal Spell", ManaCost: "{B}", TypeLine: "Instant"},
		"o-land":  {OracleId: "o-land", Name: "A Dual Land", TypeLine: "Land — Plains Swamp"},
		"o-faces": {OracleId: "o-faces", Name: "A Front // A Back", TypeLine: "Sorcery // Land", Faces: []*mtgv1.CardFace{{ManaCost: "{1}{W}"}, {}}},
	}
	deck := &mtgv1.Deck{
		Summary:            "a deck",
		CommanderOracleIds: []string{"o-lead"},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-bolt", Name: "A Cheap Removal Spell", Count: 1, Role: removal},
			{OracleId: "o-land", Name: "A Dual Land", Count: 1},
			{OracleId: "o-faces", Name: "A Front // A Back", Count: 1},
			{OracleId: "o-gone", Name: "A Card Outside The Data", Count: 1},
		},
	}
	sc := llm.NewScript(llm.Step{Output: json.RawMessage(`{
		"plan_coherent": {"grade": "yes", "why": "one plan"},
		"theme_fit": {"grade": "yes", "why": "the theme"},
		"useful_as_built": {"grade": "yes", "why": "plays"},
		"summary_honest": {"grade": "yes", "why": "plain"}
	}`)})
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := JudgePlan(context.Background(), c, "a deck", deck, cards, nil, nil); err != nil {
		t.Fatal(err)
	}
	in := sc.Calls[0].Input
	for _, want := range []string{
		"Commander: A Two-Color Leader | {2}{W}{B} | Legendary Creature — Human Noble\n",
		"Color identity: WB\n",
		"1 A Cheap Removal Spell | {B} | Instant (" + roleWord(removal) + ")\n",
		"1 A Dual Land | Land — Plains Swamp\n",
		"1 A Front // A Back | {1}{W} | Sorcery // Land\n",
		"1 A Card Outside The Data\n",
	} {
		if !strings.Contains(in, want) {
			t.Errorf("judge input lacks %q:\n%s", want, in)
		}
	}
	if !strings.Contains(sc.Calls[0].Instructions, "not from your memory of the cards") {
		t.Error("the instructions do not tell the judge to read the facts")
	}
	if plain := DeckText(deck, cards); strings.Contains(plain, "{B}") || strings.Contains(plain, "Color identity") {
		t.Errorf("DeckText carries the facts, and the tier judge reads it:\n%s", plain)
	}

	// A colorless commander reads the word, and a deck with no commander
	// names no identity.
	cards["o-grey"] = &mtgv1.Card{OracleId: "o-grey", Name: "A Colorless Leader", ManaCost: "{7}", TypeLine: "Legendary Artifact Creature"}
	if got := planDeckText(&mtgv1.Deck{CommanderOracleIds: []string{"o-grey"}}, cards, nil); !strings.Contains(got, "Color identity: colorless\n") {
		t.Errorf("colorless commander:\n%s", got)
	}
	if got := planDeckText(&mtgv1.Deck{Cards: deck.GetCards()}, cards, nil); strings.Contains(got, "Color identity") {
		t.Errorf("a deck with no commander names an identity:\n%s", got)
	}
}

// TestPlanDeckTextMarksTheSets is D-788: a request that names sets marks
// each card in or outside them from the card data, the commander included.
// A reprint reads in the sets. The cards are test fixtures.
func TestPlanDeckTextMarksTheSets(t *testing.T) {
	cards := source{
		"o-lead":    {OracleId: "o-lead", Name: "A Set Leader", SetCodes: []string{"hob"}},
		"o-reprint": {OracleId: "o-reprint", Name: "An Old Reprint", SetCodes: []string{"5dn", "hoc"}},
		"o-other":   {OracleId: "o-other", Name: "A Card From Elsewhere", SetCodes: []string{"m21"}},
	}
	deck := &mtgv1.Deck{
		CommanderOracleIds: []string{"o-lead"},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-reprint", Name: "An Old Reprint", Count: 1},
			{OracleId: "o-other", Name: "A Card From Elsewhere", Count: 1},
		},
	}
	got := planDeckText(deck, cards, []string{"hob", "hoc"})
	for _, want := range []string{
		"Commander: A Set Leader (in the sets)\n",
		"Sets of the request: hob, hoc\n",
		"1 An Old Reprint (in the sets)\n",
		"1 A Card From Elsewhere (outside the sets)\n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("deck text lacks %q:\n%s", want, got)
		}
	}
	if plain := planDeckText(deck, cards, nil); strings.Contains(plain, "sets") {
		t.Errorf("a request with no sets marks them:\n%s", plain)
	}
}
