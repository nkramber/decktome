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

// TestJudgeBracketReadsTheSpellbookCombos is F-126: the bracket judge
// reads the combos of the stored content check, and a deck with no combo
// or no check says so, so the judge names no combo from memory (D-790).
func TestJudgeBracketReadsTheSpellbookCombos(t *testing.T) {
	cards := source{"o-leader": {OracleId: "o-leader", Name: "A Test Commander"}}
	deckWith := func(c *mtgv1.ContentCheck) *mtgv1.Deck {
		return &mtgv1.Deck{CommanderOracleIds: []string{"o-leader"}, Profile: &mtgv1.DeckProfile{Content: c}}
	}
	for _, tc := range []struct {
		name string
		deck *mtgv1.Deck
		want []string
	}{
		{"two combos", deckWith(&mtgv1.ContentCheck{Checked: true, Combos: []*mtgv1.ComboHit{
			{Cards: []string{"Card A", "Card B"}, TwoCard: true, Speed: 4},
			{Cards: []string{"Card C", "Card D", "Card E"}, Speed: 2},
		}}), []string{
			"\nCombos:\n- Card A + Card B (two cards, four mana or less)\n",
			"- Card C + Card D + Card E (3 cards, not a two-card combo, eight mana or less)\n",
		}},
		{"no combo", deckWith(&mtgv1.ContentCheck{Checked: true}), []string{"\nCombos:\nCommander Spellbook finds no combo in this deck.\n"}},
		{"unchecked", deckWith(&mtgv1.ContentCheck{Error: "timeout"}), []string{"The combo check did not run for this deck: timeout.\n"}},
		{"no profile", &mtgv1.Deck{CommanderOracleIds: []string{"o-leader"}}, []string{"The combo check did not run for this deck.\n"}},
	} {
		c, sc := bracketClient(t, `{"bracket": "3", "why": "a combo"}`)
		if _, err := JudgeBracket(context.Background(), c, tc.deck, cards, nil); err != nil {
			t.Fatal(err)
		}
		in := sc.Calls[0].Input
		for _, want := range tc.want {
			if !strings.Contains(in, want) {
				t.Errorf("%s: judge input lacks %q:\n%s", tc.name, want, in)
			}
		}
		if !strings.Contains(sc.Calls[0].Instructions, "Count only a listed combo") {
			t.Errorf("%s: the instructions do not bind the judge to the listed combos", tc.name)
		}
	}
	if plain := DeckText(deckWith(&mtgv1.ContentCheck{Checked: true}), cards); strings.Contains(plain, "Combos:") {
		t.Errorf("DeckText lists combos, and the tier and plan judges read it:\n%s", plain)
	}
	for speed, want := range map[int32]string{6: "no mana past the cards", 5: "no mana past the cards", 3: "six mana or less", 1: "more than eight mana"} {
		if got := speedWord(speed); got != want {
			t.Errorf("speedWord(%d) = %q, want %q", speed, got, want)
		}
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

// TestJudgeSixtyStep is D-866: the step judge answers one of three
// steps from its schema, and it refuses any other word. The format word
// heads the input, and the card facts and the rules text follow each
// name (D-872).
func TestJudgeSixtyStep(t *testing.T) {
	cards := source{
		"o-bolt": {OracleId: "o-bolt", Name: "Lightning Bolt", ManaCost: "{R}", TypeLine: "Instant", OracleText: "Lightning Bolt deals 3\ndamage to any target."},
		"o-dfc": {OracleId: "o-dfc", Name: "Front // Back", TypeLine: "Instant // Land", Faces: []*mtgv1.CardFace{
			{ManaCost: "{1}{R}", OracleText: "Front deals 2 damage."}, {OracleText: "{T}: Add {R}."},
		}},
	}
	deck := &mtgv1.Deck{
		Cards:     []*mtgv1.DeckCard{{OracleId: "o-bolt", Name: "Lightning Bolt", Count: 4}},
		Sideboard: []*mtgv1.DeckCard{{OracleId: "o-dfc", Name: "Front // Back", Count: 1}},
	}
	for word, want := range map[string]mtgv1.SixtyStep{
		"casual":     mtgv1.SixtyStep_SIXTY_STEP_CASUAL,
		"fnm":        mtgv1.SixtyStep_SIXTY_STEP_FNM,
		"tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
	} {
		c, sc := bracketClient(t, `{"step": "`+word+`", "why": "w"}`)
		j, err := JudgeSixtyStep(context.Background(), c, deck, "Pioneer", cards, nil)
		if err != nil {
			t.Fatal(err)
		}
		if j.Step != want || j.Why != "w" {
			t.Errorf("%s: step = %v, why = %q", word, j.Step, j.Why)
		}
		in := sc.Calls[0].Input
		for _, want := range []string{
			"Format: Pioneer\n\nCards:\n",
			"4 Lightning Bolt | {R} | Instant | Lightning Bolt deals 3 damage to any target.\n",
			"\nSideboard:\n1 Front // Back | {1}{R} | Instant // Land | Front deals 2 damage. // {T}: Add {R}.\n",
		} {
			if !strings.Contains(in, want) {
				t.Errorf("judge input lacks %q:\n%s", want, in)
			}
		}
		if sc.Calls[0].SchemaName != "sixty_step" {
			t.Errorf("schema = %q", sc.Calls[0].SchemaName)
		}
	}
	c, _ := bracketClient(t, `{"step": "competitive", "why": "w"}`)
	if _, err := JudgeSixtyStep(context.Background(), c, deck, "Modern", cards, nil); err == nil {
		t.Error("a step outside the schema passed")
	}
}
