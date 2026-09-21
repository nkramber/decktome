package generate

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
)

func bracketClient(t *testing.T, output string) (*llm.Client, *llm.Script) {
	t.Helper()
	sc := llm.NewScript(llm.Step{Output: json.RawMessage(output)})
	c, err := llm.New(fakeConfig(), []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	return c, sc
}

// TestJudgeBracketReadsTheGameChangerFlags is F-123: the bracket judge
// reads a mark on each card that the card data flags, the commander
// included, and the tier and plan judges read the list unmarked.
func TestJudgeBracketReadsTheGameChangerFlags(t *testing.T) {
	ramp := mtgv1.CardRole_CARD_ROLE_RAMP
	cards := source{
		"o-kinnan": {OracleId: "o-kinnan", Name: "Kinnan, Bonder Prodigy"},
		"o-vault":  {OracleId: "o-vault", Name: "Mana Vault", GameChanger: true},
		"o-ring":   {OracleId: "o-ring", Name: "Sol Ring"},
		"o-oracle": {OracleId: "o-oracle", Name: "Thassa's Oracle", GameChanger: true},
	}
	deck := &mtgv1.Deck{
		CommanderOracleIds: []string{"o-kinnan"},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-vault", Name: "Mana Vault", Count: 1, Role: ramp},
			{OracleId: "o-ring", Name: "Sol Ring", Count: 1, Role: ramp},
			{OracleId: "o-oracle", Name: "Thassa's Oracle", Count: 1},
		},
	}
	c, sc := bracketClient(t, `{"bracket": "5", "why": "fast mana and a combo"}`)
	j, err := JudgeBracket(context.Background(), c, deck, cards, nil)
	if err != nil {
		t.Fatal(err)
	}
	if j.Bracket != 5 || len(sc.Calls) != 1 {
		t.Fatalf("bracket = %d, calls = %d, want 5 and one call", j.Bracket, len(sc.Calls))
	}
	in := sc.Calls[0].Input
	for _, want := range []string{
		"Commander: Kinnan, Bonder Prodigy\n",
		"1 Mana Vault (" + roleWord(ramp) + ", Game Changer)\n",
		"1 Sol Ring (" + roleWord(ramp) + ")\n",
		"1 Thassa's Oracle (Game Changer)\n",
	} {
		if !strings.Contains(in, want) {
			t.Errorf("judge input lacks %q:\n%s", want, in)
		}
	}
	if plain := DeckText(deck, cards); strings.Contains(plain, "Game Changer") {
		t.Errorf("DeckText marks a Game Changer, and the tier and plan judges read it:\n%s", plain)
	}

	// A flagged commander carries the mark on its own line. The card is a
	// test fixture and names no real card.
	leader := source{"o-leader": {OracleId: "o-leader", Name: "A Flagged Commander", GameChanger: true}}
	c2, sc2 := bracketClient(t, `{"bracket": "4", "why": "a Game Changer leads"}`)
	if _, err := JudgeBracket(context.Background(), c2, &mtgv1.Deck{CommanderOracleIds: []string{"o-leader"}}, leader, nil); err != nil {
		t.Fatal(err)
	}
	if in := sc2.Calls[0].Input; !strings.Contains(in, "Commander: A Flagged Commander (Game Changer)\n") {
		t.Errorf("a flagged commander carries no mark:\n%s", in)
	}
}

// TestJudgeSummaryReadsTheCardFacts is D-789: with a deck, the summary
// judge reads the cost and the type of each card from the card data, and
// with none it reads the summary alone. The card is a test fixture.
func TestJudgeSummaryReadsTheCardFacts(t *testing.T) {
	cards := source{"o-lead": {OracleId: "o-lead", Name: "A Six-Mana Leader", ManaCost: "{4}{W}{B}", TypeLine: "Legendary Creature"}}
	deck := &mtgv1.Deck{CommanderOracleIds: []string{"o-lead"}}
	answer := `{"claims": [], "verdict": "clean"}`
	c, sc := bracketClient(t, answer)
	if _, err := JudgeSummary(context.Background(), c, "a deck", "A commander that costs six mana.", deck, cards, nil); err != nil {
		t.Fatal(err)
	}
	if in := sc.Calls[0].Input; !strings.Contains(in, "Commander: A Six-Mana Leader | {4}{W}{B} | Legendary Creature\n") {
		t.Errorf("judge input lacks the facts:\n%s", in)
	}
	if !strings.Contains(sc.Calls[0].Instructions, "not against your memory of the card") {
		t.Error("the instructions do not tell the judge to read the facts")
	}
	c2, sc2 := bracketClient(t, answer)
	if _, err := JudgeSummary(context.Background(), c2, "a deck", "A summary.", nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if in := sc2.Calls[0].Input; strings.Contains(in, "Commander:") {
		t.Errorf("a nil deck sends a card list:\n%s", in)
	}
}
