package generate

import (
	"context"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
	"github.com/nkramber/decktome/go/internal/spellbook"
)

// scoredCard is one shortlist card and its score.
type scoredCard struct {
	card  *mtgv1.Card
	score float64
}

// scoredPool builds a pool through FromList, the way the build does, so
// each shortlist card carries its score and each always card has none.
func scoredPool(always []*mtgv1.Card, cs ...scoredCard) *Pool {
	l := &candidates.List{}
	for _, c := range cs {
		l.Candidates = append(l.Candidates, candidates.Candidate{Card: c.card, Score: c.score})
	}
	return FromList(l, always, false)
}

// TestCutForbiddenPicksTheCard is PR-45a: one cut breaks each forbidden
// combo, the card in the most combos first, then a card outside the
// precon, then the card with the lowest shortlist score. A commander, a
// locked card, and a card a revision keeps never leave. Each cut names the
// content it held.
func TestCutForbiddenPicksTheCard(t *testing.T) {
	a, b, c, d, e := card("o-a", "Card A"), card("o-b", "Card B"), card("o-c", "Card C"), card("o-d", "Card D"), card("o-e", "Card E")
	// The scores disagree with the alphabet, so a rule by name or by
	// place can not pass.
	pool := scoredPool(nil, scoredCard{b, 0.5}, scoredCard{e, 0.4}, scoredCard{c, 0.3}, scoredCard{d, 0.2}, scoredCard{a, 0.1})
	// Card A is an always card here, so it has no score.
	given := scoredPool([]*mtgv1.Card{a}, scoredCard{b, 0.5})
	pair := [][]string{{"Card A", "Card B"}}
	cases := []struct {
		name string
		pool *Pool
		req  Request
		f    profile.Forbidden
		want string
	}{
		{"the card in two combos", pool, Request{}, profile.Forbidden{Combos: [][]string{{"Card A", "Card B"}, {"Card B", "Card C"}}}, "Card B (the combo Card A + Card B)"},
		{"the lower shortlist score", pool, Request{}, profile.Forbidden{Combos: pair}, "Card A (the combo Card A + Card B)"},
		{"a card with no score stays", given, Request{}, profile.Forbidden{Combos: pair}, "Card B (the combo Card A + Card B)"},
		{"a card outside the precon", pool, Request{PreconOracleIDs: []string{"o-a"}}, profile.Forbidden{Combos: pair}, "Card B (the combo Card A + Card B)"},
		{"a locked card stays", pool, Request{Locked: []string{"o-a"}}, profile.Forbidden{Combos: pair}, "Card B (the combo Card A + Card B)"},
		{"a locked pair stays", pool, Request{Locked: []string{"o-a", "o-b"}}, profile.Forbidden{Combos: pair}, ""},
		{"a revision keep stays", pool, Request{Revision: &Revision{Keep: []string{"Card A"}}}, profile.Forbidden{Combos: pair}, "Card B (the combo Card A + Card B)"},
		{"mass land denial", pool, Request{}, profile.Forbidden{MassLandDenial: []string{"Card C"}}, "Card C (mass land denial)"},
		{"extra turns over the cap", pool, Request{}, profile.Forbidden{ExtraTurns: []string{"Card D", "Card E"}, ExtraTurnCap: 1}, "Card D (an extra-turn card past the limit of 1)"},
		{"extra turns with no cap", pool, Request{}, profile.Forbidden{ExtraTurns: []string{"Card D", "Card E"}}, "Card E (an extra-turn card),Card D (an extra-turn card)"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := &mtgv1.Deck{CommanderOracleIds: []string{"o-cmd"}}
			for _, n := range tc.pool.Names() {
				pc, _ := tc.pool.Card(n)
				d.Cards = append(d.Cards, &mtgv1.DeckCard{OracleId: pc.GetOracleId(), Name: n, Count: 1})
			}
			before := len(d.GetCards())
			tc.req.Pool = tc.pool
			cut := cutForbidden(d, tc.req, tc.f)
			var got []string
			for _, c := range cut {
				got = append(got, c.card.GetName()+" ("+c.why+")")
			}
			if strings.Join(got, ",") != tc.want {
				t.Fatalf("cut %v, want %q", got, tc.want)
			}
			if len(d.GetCards()) != before-len(cut) {
				t.Errorf("deck holds %d cards after %d cuts from %d", len(d.GetCards()), len(cut), before)
			}
		})
	}
}

// comboClassifier reads a two-card combo whenever both of its cards are in
// the list it gets, the way the endpoint would.
type comboClassifier struct {
	pair  [2]string
	speed int
}

func (c *comboClassifier) EstimateBracket(_ context.Context, _ []string, main []string) (*spellbook.Result, error) {
	has := map[string]bool{}
	for _, n := range main {
		has[n] = true
	}
	res := &spellbook.Result{}
	if has[c.pair[0]] && has[c.pair[1]] {
		res.Combos = []spellbook.ClassifiedCombo{{
			Combo: spellbook.ComboRef{ID: "1", Uses: []spellbook.ComboUse{
				{Card: spellbook.CardRef{Name: c.pair[0]}}, {Card: spellbook.CardRef{Name: c.pair[1]}}}},
			Relevant: true, DefinitelyTwoCard: true, Speed: c.speed,
		}}
	}
	return res, nil
}

// TestBuildCutsAForbiddenCombo is PR-45a end to end: a bracket 3 build
// that holds a fast two-card combo loses one card of it, the deck says so,
// and the warning is gone. A combo whose cards are both locked stays, with
// its warning.
func TestBuildCutsAForbiddenCombo(t *testing.T) {
	out := step(t, deckOut{Summary: "a deck", Cards: []Entry{
		{Name: "Ajani's Welcome", Count: 1, Role: "synergy", Reason: "gains life"},
		{Name: "Ajani's Pridemate", Count: 1, Role: "threat", Reason: "grows on each gain"},
	}})
	base := testPool()
	welcome, _ := base.Card("Ajani's Welcome")
	pridemate, _ := base.Card("Ajani's Pridemate")
	solRing, _ := base.Card("Sol Ring")
	karlov, _ := base.Card("Karlov of the Ghost Council")
	for _, lockedPair := range []bool{false, true} {
		b, _, _ := testBuilder(t, out, out, out)
		cfg, err := rules.Load()
		if err != nil {
			t.Fatal(err)
		}
		prof, err := profile.New(cfg, nil, &comboClassifier{pair: [2]string{"Ajani's Welcome", "Ajani's Pridemate"}, speed: 5})
		if err != nil {
			t.Fatal(err)
		}
		prof.SetHands(50)
		b.profiler = prof
		req := testRequest()
		// Ajani's Pridemate scores lower, and the alphabet would cut
		// Ajani's Welcome.
		req.Pool = scoredPool([]*mtgv1.Card{karlov}, scoredCard{welcome, 0.9}, scoredCard{pridemate, 0.5}, scoredCard{solRing, 0.3})
		req.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
		req.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}}
		req.Limits = "Exactly 100 cards."
		req.Commanders = []string{"o-karlov"}
		if lockedPair {
			req.Locked = []string{"o-welcome", "o-pridemate"}
		}
		got, err := b.Build(context.Background(), req, nil)
		if err != nil {
			t.Fatalf("build: %v", err)
		}
		var cutMsg string
		var comboWarn bool
		for _, f := range got.Deck.GetValidation().GetFindings() {
			switch f.GetCode() {
			case CodeBracketCut:
				cutMsg = f.GetMessage()
			case profile.CodeTwoCardCombo:
				comboWarn = true
			}
		}
		held := map[string]bool{}
		for _, c := range got.Deck.GetCards() {
			held[c.GetName()] = true
		}
		if lockedPair {
			if cutMsg != "" || !comboWarn || !held["Ajani's Pridemate"] || !held["Ajani's Welcome"] {
				t.Errorf("locked pair: cut %q, warning %v, held %v", cutMsg, comboWarn, held)
			}
			continue
		}
		if held["Ajani's Pridemate"] || !held["Ajani's Welcome"] {
			t.Errorf("held %v, want Ajani's Pridemate cut and Ajani's Welcome kept", held)
		}
		if !strings.Contains(cutMsg, "to hold bracket 3, the builder cut 1 card: Ajani's Pridemate (the combo Ajani's Welcome + Ajani's Pridemate)") {
			t.Errorf("cut finding %q", cutMsg)
		}
		if comboWarn {
			t.Error("the combo warning stays after the cut")
		}
	}
}

// TestCheckPreconShareExceptsTheBracketCut is D-695: a precon card the
// bracket cut leaves the base of the share, so the cut never counts
// against the 85 percent.
func TestCheckPreconShareExceptsTheBracketCut(t *testing.T) {
	_, src, _ := testBuilder(t)
	req := testRequest()
	req.Precon = "Test Precon"
	req.PreconOracleIDs = []string{"o-welcome", "o-pridemate", "o-solring"}
	deck := &mtgv1.Deck{Cards: []*mtgv1.DeckCard{
		{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 1},
		{OracleId: "o-solring", Name: "Sol Ring", Count: 1},
	}}
	if PreconKeepCount(3) > 2 && !checkPreconShare(deck, req, src) {
		t.Error("without the cut, 2 of 3 precon cards must warn")
	}
	deck.Validation = nil
	if checkPreconShareExcept(deck, req, src, map[string]bool{"o-pridemate": true}) {
		t.Errorf("a cut precon card counted against the share: %v", deck.GetValidation().GetFindings())
	}
}

// TestPoolScoreFollowsTheShortlist is the rank source of PR-45a: FromList
// records each shortlist score, Names still sorts by the alphabet, a card
// the pool took as given has no score, and Filter keeps the scores.
func TestPoolScoreFollowsTheShortlist(t *testing.T) {
	a, b, k := card("o-a", "Card A"), card("o-b", "Card B"), card("o-k", "Kept Card")
	pool := scoredPool([]*mtgv1.Card{k}, scoredCard{b, 0.7}, scoredCard{a, 0.2})
	if got := strings.Join(pool.Names(), ","); got != "Card A,Card B,Kept Card" {
		t.Errorf("names %q, want the alphabet", got)
	}
	if s, ok := pool.Score("Card A"); !ok || s != 0.2 {
		t.Errorf("score of Card A = %v, %v, want 0.2", s, ok)
	}
	if _, ok := pool.Score("Kept Card"); ok {
		t.Error("a card the pool took as given carries a score")
	}
	filtered := pool.Filter(func(c *mtgv1.Card) bool { return c.GetOracleId() != "o-a" })
	if s, ok := filtered.Score("Card B"); !ok || s != 0.7 {
		t.Errorf("filtered score of Card B = %v, %v, want 0.7", s, ok)
	}
	if _, ok := filtered.Score("Card A"); ok {
		t.Error("a card the filter dropped keeps a score")
	}
}
