package questions

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// suggestionTurnOne is a Commander request with every slot but the
// commander filled, and a classifier that reads a request for names.
func suggestionTurnOne() classifyOut {
	out := commanderClassify()
	out.Power, out.PoolRule = "bracket 5", "owned_only"
	out.Facts.WantsSuggestion = true
	return out
}

func suggestionHints() *fakeHints {
	return &fakeHints{commanders: []string{"Karlov of the Ghost Council", "Oloro, Ageless Ascetic", "Ayli, Eternal Pilgrim"}}
}

// askedCommanderRow reports which commander row the turn asked: the row
// that lets the reader name one, or the pick row with three names.
func askedCommanderRow(qs []*mtgv1.Question) string {
	for _, q := range qs {
		if q.GetSlot() != "commander" {
			continue
		}
		if strings.HasSuffix(q.GetId(), "-commander_pick") {
			return "commander_pick"
		}
		return "commander"
	}
	return ""
}

// TestAFormatWordAsksForNoCommanderNames is D-1220. The reader of deck
// v-8k83YGt2Dn6IddKn3CBT wrote "the best possible commander deck", and the
// classifier set wants_suggestion. The pick row asked twice, and the row
// that lets the reader name a commander never asked. The reader wrote
// "Never asked me if I want to name my own commander or not".
//
// The class is each request that names the format, or no commander, and
// asks for no names. "For Commander" also passed the superlative rule of
// D-167 before, so "the best deck for Commander" skipped both commander
// keys, and no commander question asked at all.
func TestAFormatWordAsksForNoCommanderNames(t *testing.T) {
	for _, msg := range []string{
		"Build the best possible commander deck from only the Lord of the Rings and Hobbit sets.",
		"Build the best deck for Commander from my collection, lifegain in white and black.",
		"I want a strong EDH deck from my cards, lifegain in white and black.",
	} {
		t.Run(msg, func(t *testing.T) {
			a, _ := testAgentHints(t, suggestionHints(),
				classifyStep(t, suggestionTurnOne()), fits(t, "commander"), askStep(t))
			st := NewState(true)
			res, err := a.Turn(context.Background(), st, msg, nil)
			if err != nil {
				t.Fatalf("turn 1: %v", err)
			}
			if st.Ctx.Suggested {
				t.Error("a request with no words for names set Suggested")
			}
			if got := askedCommanderRow(res.Questions); got != "commander" {
				t.Errorf("turn 1 asked commander row %q, want the row that lets the reader name one", got)
			}
		})
	}
}

// TestWordsForNamesStillGetTheOffer guards the other side of D-1220. A
// message that asks for names in words, or answers a commander question,
// keeps the fact, and the pick row asks.
func TestWordsForNamesStillGetTheOffer(t *testing.T) {
	for _, msg := range []string{
		"A lifegain commander deck in white and black. Suggest a commander.",
		"A lifegain deck for Commander in white and black. I have no commander in mind.",
		"Lifegain in white and black for Commander. Which commander should I use?",
	} {
		t.Run(msg, func(t *testing.T) {
			a, _ := testAgentHints(t, suggestionHints(),
				classifyStep(t, suggestionTurnOne()), fits(t, "commander_pick"), askStep(t))
			st := NewState(true)
			res, err := a.Turn(context.Background(), st, msg, nil)
			if err != nil {
				t.Fatalf("turn 1: %v", err)
			}
			if !st.Ctx.Suggested {
				t.Error("a request for names in words did not set Suggested")
			}
			if got := askedCommanderRow(res.Questions); got != "commander_pick" {
				t.Errorf("turn 1 asked commander row %q, want the pick row", got)
			}
		})
	}
}

// TestSuggestOneAnswersTheCommanderRow is the path of D-290 under D-1220.
// The commander row asks first, and its option "Suggest one" carries no
// commander word. The fact counts, because a commander question is out.
func TestSuggestOneAnswersTheCommanderRow(t *testing.T) {
	first := suggestionTurnOne()
	first.Facts.WantsSuggestion = false
	second := suggestionTurnOne()
	a, _ := testAgentHints(t, suggestionHints(),
		classifyStep(t, first), classifyStep(t, second))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "Build the best possible commander deck.", nil)
	if err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if got := askedCommanderRow(res.Questions); got != "commander" {
		t.Fatalf("turn 1 asked commander row %q, want the open row", got)
	}
	res, err = a.Turn(context.Background(), st, "Suggest one", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	if got := askedCommanderRow(res.Questions); got != "commander_pick" {
		t.Errorf("turn 2 asked commander row %q, want the pick row after \"Suggest one\"", got)
	}
}

// TestFormatWordsAreNotTheCard is the word side of D-1220. A format
// preposition or verb before "Commander" names the format, and an article
// or a possessive names the card.
func TestFormatWordsAreNotTheCard(t *testing.T) {
	for msg, want := range map[string]bool{
		"the best deck for commander":            false,
		"a lifegain deck in commander":           false,
		"i want to play commander with my cards": false,
		"the best possible commander deck":       false,
		"a deck for my commander":                true,
		"pick a commander for me":                true,
		"which commander should i use":           true,
	} {
		if got := namesCommander(msg); got != want {
			t.Errorf("namesCommander(%q) = %v, want %v", msg, got, want)
		}
	}
	if !mentionsCommander("i have no commander in mind") {
		t.Error("mentionsCommander skipped a negated commander word")
	}
	if delegatesCommander("build the best deck for commander from my collection") {
		t.Error("a superlative about the deck delegated the commander choice")
	}
	if !delegatesCommander("buy the best lifegain commander") {
		t.Error("the superlative of D-167 no longer delegates the commander choice")
	}
}
