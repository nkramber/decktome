package questions

import (
	"context"
	"io"
	"log/slog"
	"slices"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// The F-75 tests. The reader wrote "Aragorn as commander", no card
// carries that name alone, and the build dropped the name and picked its
// own commander (D-232). The row asks which card the reader means now
// (D-606), and it offers an escape (D-607).

// aragornIndex holds the four cards that carry the name, one card that
// does not, and one card that can not lead a deck.
func aragornIndex() *cards.Index {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	card := func(name, oracle string, rank int32, gameChanger bool) *mtgv1.Card {
		return &mtgv1.Card{
			Name: name, OracleId: oracle, EdhrecRank: rank, GameChanger: gameChanger,
			TypeLine: "Legendary Creature — Human", CanBeCommander: true, Legalities: legal,
		}
	}
	return cards.NewIndex([]*mtgv1.Card{
		card("Aragorn, King of Gondor", "o-king", 1, false),
		card("Aragorn, Company Leader", "o-company", 2, false),
		card("Aragorn, the Uniter", "o-uniter", 3, true),
		card("Aragorn, Hornburg Hero", "o-hornburg", 4, false),
		card("Thranduil, Elvenking", "o-thranduil", 5, false),
		{Name: "Andúril, Flame of the West", TypeLine: "Legendary Artifact — Equipment", Legalities: legal},
	}, nil, nil, time.Time{})
}

func TestResolveCommanderNamesTheBestThreeCards(t *testing.T) {
	h := &CandidateHints{Index: aragornIndex()}
	got := h.ResolveCommander("Aragorn")
	want := []string{"Aragorn, King of Gondor", "Aragorn, Company Leader", "Aragorn, the Uniter"}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("ResolveCommander(Aragorn) = %v, want the three most popular %v", got, want)
	}
}

// A name no card holds answers nothing, and the row then says so.
func TestResolveCommanderAnswersNothingForAnUnknownName(t *testing.T) {
	h := &CandidateHints{Index: aragornIndex()}
	if got := h.ResolveCommander("Gandolf"); len(got) != 0 {
		t.Errorf("ResolveCommander(Gandolf) = %v, want nothing", got)
	}
	// A card that can not lead a deck is not an answer to this question.
	if got := h.ResolveCommander("Andúril"); len(got) != 0 {
		t.Errorf("ResolveCommander(Andúril) = %v, want nothing", got)
	}
	// The test reads whole words, so a fragment finds nothing.
	if got := h.ResolveCommander("Ara"); len(got) != 0 {
		t.Errorf("ResolveCommander(Ara) = %v, want nothing", got)
	}
}

// The words after the comma name the card too.
func TestResolveCommanderReadsTheWordsAfterTheComma(t *testing.T) {
	h := &CandidateHints{Index: aragornIndex()}
	got := h.ResolveCommander("King of Gondor")
	if len(got) != 1 || got[0] != "Aragorn, King of Gondor" {
		t.Errorf("ResolveCommander(King of Gondor) = %v, want the one card", got)
	}
}

// A game changer leaves a bracket 1 or 2 list, the rule CommanderPool
// applies to every offer.
func TestResolveCommanderDropsAGameChangerAtBracketTwo(t *testing.T) {
	h := &CandidateHints{Index: aragornIndex(), Bracket: 2}
	for _, name := range h.ResolveCommander("Aragorn") {
		if name == "Aragorn, the Uniter" {
			t.Error("a game changer reached a bracket 2 list")
		}
	}
}

// A bracket 4 or 5 request ranks on the cEDH signal of the quality
// model, and not on popularity (OQ-48, PR-14B).
func TestResolveCommanderRanksOnTheSignalAtBracketFive(t *testing.T) {
	h := &CandidateHints{Index: aragornIndex(), Bracket: 5,
		CommanderSignal: func(ids ...string) float64 {
			if len(ids) > 0 && ids[0] == "o-hornburg" {
				return 1
			}
			return 0
		}}
	got := h.ResolveCommander("Aragorn")
	if len(got) == 0 || got[0] != "Aragorn, Hornburg Hero" {
		t.Errorf("ResolveCommander at bracket 5 = %v, want the strongest card first", got)
	}
}

// TestAPartialCommanderNameAsksWhichCard is F-75 whole. The name reaches
// no build until the reader picks a card.
func TestAPartialCommanderNameAsksWhichCard(t *testing.T) {
	h := &CandidateHints{Index: aragornIndex()}
	// The row offers the best three at the power the reader picked, so it
	// waits for the power (D-606, D-630). The reader names it here.
	out := classifyOut{Format: "commander", Theme: "humans", PoolRule: "unknown",
		Power: "bracket 4", CommanderNames: []string{"Aragorn"}}
	a, _ := testAgentHints(t, h, classifyStep(t, out),
		fits(t, "commander_unresolved", "power_commander"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "Build me an Aragorn deck.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if st.UnresolvedCommander != "Aragorn" {
		t.Fatalf("the unresolved name is %q, want Aragorn", st.UnresolvedCommander)
	}
	if st.Ctx.CommanderSet {
		t.Error("a name the card index does not hold became the commander")
	}
	if hasName(st.NamedCards, "Aragorn") {
		t.Error("the name stayed in the named cards, so the role row can ask about it")
	}
	if len(st.CommanderOptions) != 3 {
		t.Errorf("the row offers %d cards, want 3", len(st.CommanderOptions))
	}
	if res.Ready {
		t.Error("the session was ready with the commander still unsettled")
	}
	var asked *mtgv1.Question
	for _, q := range res.Questions {
		if strings.Contains(q.GetText(), "Aragorn") {
			asked = q
		}
	}
	if asked == nil {
		t.Fatalf("no question named the card: %v", res.Questions)
	}
	if n := len(asked.GetOptions()); n != 4 {
		t.Errorf("the question offers %d options, want the three cards and %q", n, NoneOfTheseOption)
	}
}

// A name no card holds says so, and it offers no card to pick.
func TestACommanderNameNoCardHoldsSaysSo(t *testing.T) {
	h := &CandidateHints{Index: aragornIndex()}
	out := classifyOut{Format: "commander", Theme: "wizards", PoolRule: "unknown",
		CommanderNames: []string{"Gandolf"}}
	a, _ := testAgentHints(t, h, classifyStep(t, out),
		fits(t, "commander_unknown", "power_commander"), askStep(t))
	st := NewState(false)
	res, err := a.Turn(context.Background(), st, "Build me a Gandolf deck.", nil)
	if err != nil {
		t.Fatalf("turn: %v", err)
	}
	if !st.Ctx.CommanderNoMatch {
		t.Fatal("the state does not mark a name no card holds")
	}
	if len(st.CommanderOptions) != 0 {
		t.Errorf("the row offers %v, want nothing", st.CommanderOptions)
	}
	for _, q := range res.Questions {
		if strings.Contains(q.GetText(), "Gandolf") {
			return
		}
	}
	t.Errorf("no question named the word the reader wrote: %v", res.Questions)
}

// TestPickingACardSetsTheCommanderWithNoModel is D-597 for this row. The
// option index names one card exactly, and the classifier is not in the
// path.
func TestPickingACardSetsTheCommanderWithNoModel(t *testing.T) {
	a, st := commanderRowAsked(t)
	st.OptionAnswers = []OptionAnswer{{QuestionID: "q1-commander_unresolved", Index: 1}}
	a.applyOptionAnswers(st)
	if !hasName(st.CommanderNames, "Aragorn, Company Leader") {
		t.Fatalf("the commander is %v, want the card the reader picked", st.CommanderNames)
	}
	if st.Ctx.CommanderUnresolved || st.UnresolvedCommander != "" {
		t.Error("the row still asks after the reader answered it")
	}
	if st.Slots.GetSlotStates()[SlotCommanderUnresolved] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Errorf("the key is %v, want filled", st.Slots.GetSlotStates()[SlotCommanderUnresolved])
	}
}

// "None of these" drops the name, and the app then offers its own three
// commanders (D-607).
func TestNoneOfTheseDropsTheName(t *testing.T) {
	a, st := commanderRowAsked(t)
	st.OptionAnswers = []OptionAnswer{{QuestionID: "q1-commander_unresolved", Index: 3}}
	a.applyOptionAnswers(st)
	if st.Ctx.CommanderSet {
		t.Error("the escape option set a commander")
	}
	if st.Ctx.CommanderUnresolved || st.UnresolvedCommander != "" {
		t.Error("the row still asks after the reader refused every card")
	}
	if st.Slots.GetSlotStates()[SlotCommanderUnresolved] != mtgv1.SlotState_SLOT_STATE_FILLED {
		t.Error("the key stayed open, so the row asks the same question again")
	}
}

// A commander the card index knows answers the row, whatever words
// brought it. Without this the row asks again with the deck already led.
func TestAKnownCommanderAnswersTheRow(t *testing.T) {
	h := &CandidateHints{Index: aragornIndex()}
	a := &Agent{cat: load(t), threshold: DefaultFitThreshold, hints: h,
		log: slog.New(slog.NewTextHandler(io.Discard, nil))}
	st := NewState(false)
	a.applyNames(st, classifyOut{CommanderNames: []string{"Aragorn"}})
	if !st.Ctx.CommanderUnresolved {
		t.Fatal("the first turn did not raise the row")
	}
	a.applyNames(st, classifyOut{CommanderNames: []string{"Aragorn, King of Gondor"}})
	if st.Ctx.CommanderUnresolved || st.UnresolvedCommander != "" {
		t.Error("the row still asks with the commander settled")
	}
	if !hasName(st.CommanderNames, "Aragorn, King of Gondor") {
		t.Errorf("the commander is %v, want the card the reader named", st.CommanderNames)
	}
}

// commanderRowAsked is a session with the row of F-75 out, and the three
// cards on the table.
func commanderRowAsked(t *testing.T) (*Agent, *State) {
	t.Helper()
	a := &Agent{cat: load(t), threshold: DefaultFitThreshold,
		log: slog.New(slog.NewTextHandler(io.Discard, nil))}
	st := NewState(false)
	st.CommanderUnresolved("Aragorn", []string{
		"Aragorn, King of Gondor", "Aragorn, Company Leader", "Aragorn, the Uniter"})
	asked(st, "q1-commander_unresolved", "commander_unresolved", "commander", SlotCommanderUnresolved)
	return a, st
}

// TestTheOfferWaitsForThePower is D-630. The owner read session
// XnY2uVuQkAL0VN9AeNCj, where one turn asked the power bracket and
// offered three commanders at once. The offer ranks on the bracket: a
// game changer leaves a bracket 1 or 2 list, and a bracket 4 or 5 list
// ranks on the cEDH signal (PR-14B, OQ-48). So an offer made before the
// power answer ranks for no power at all.
func TestTheOfferWaitsForThePower(t *testing.T) {
	cat := load(t)
	for _, id := range []string{"commander_pick", "commander_unresolved"} {
		row, ok := cat.Row(id)
		if !ok {
			t.Fatalf("no row %q", id)
		}
		if !slices.Contains(row.When.Requires, "power") {
			t.Errorf("row %q does not wait for the power, so its offer ranks for none", id)
		}
	}

	// The row that says a name matched nothing offers no card, so it
	// needs no power and it asks at once.
	unknown, _ := cat.Row("commander_unknown")
	if slices.Contains(unknown.When.Requires, "power") {
		t.Error("the row that offers no card waits for the power")
	}

	// A reader who named a commander this app can not settle must not be
	// asked "which commander do you want" while that name waits. The
	// plain row stands down until the name resolves.
	plain, _ := cat.Row("commander")
	if plain.When.CommanderUnresolved == nil || *plain.When.CommanderUnresolved {
		t.Error("the plain commander row asks while an unresolved name waits")
	}

	// The power row waits for nothing of the commander, or the two rows
	// wait for each other and neither goes out.
	power, _ := cat.Row("power_commander")
	for _, k := range power.When.Requires {
		if strings.HasPrefix(k, "commander") {
			t.Errorf("the power row waits for %q, and the commander offer waits for the power", k)
		}
	}
}

// TestTheNetReleasesTheOffer locks D-631. The offer waits for the power
// (D-630), and the net of D-351 closes a question the reader never
// answers. Close and Skip both fill a key, and CloseStalled does not.
// So without this rule the offer waits for a power that never arrives,
// and the deck reaches the reader with no commander asked. Gate run 43
// found it on the conversation "terse: the user answers with a number".
func TestTheNetReleasesTheOffer(t *testing.T) {
	st := NewState(false)
	st.Turn = 1
	st.MarkAsked("power_commander", "power", "power")

	offer := When{Requires: []string{"power"}}
	if offer.matches(st.Ctx) {
		t.Fatal("the offer went out while the power question was still open")
	}

	st.Turn = 1 + StallGrace
	closed, _ := st.CloseStalled()
	if !slices.Contains(closed, "power") {
		t.Fatalf("the net closed %v, and the power question is not among them", closed)
	}
	if !st.Ctx.Skipped["power"] {
		t.Error("the net closed the power question and did not record it")
	}
	if !offer.matches(st.Ctx) {
		t.Error("the offer still waits for a power the reader never gives")
	}

	// A re-ask reopens the key, so the offer waits again.
	st2 := NewState(false)
	st2.Turn = 1
	st2.MarkAsked("power_commander", "power", "power")
	st2.Skip("power")
	if !offer.matches(st2.Ctx) {
		t.Error("a declined power holds the offer back")
	}
	st2.Reopen("power")
	if offer.matches(st2.Ctx) {
		t.Error("the offer goes out after the power reopened")
	}
}

// TestTheOfferWaitsForThePreferences locks D-669. The offer ranks on the
// theme and the colors: the pool ranks on theme fit, and the color check
// drops a name the colors exclude (D-153). An offer beside a question
// about either one ignores the answer that question asks for.
func TestTheOfferWaitsForThePreferences(t *testing.T) {
	cat := load(t)
	pick, ok := cat.Row("commander_pick")
	if !ok {
		t.Fatal("no row commander_pick")
	}
	for _, slot := range []string{"power", "theme", "colors"} {
		if !slices.Contains(pick.When.Requires, slot) {
			t.Errorf("the pick row does not wait for %q, so its offer ignores that answer", slot)
		}
	}
	// The preference rows wait for nothing of the commander, or the rows
	// wait for each other and none of them goes out.
	for _, id := range []string{"theme", "colors"} {
		row, ok := cat.Row(id)
		if !ok {
			t.Fatalf("no row %q", id)
		}
		for _, k := range append(slices.Clone(row.When.Requires), row.When.NotOutstanding...) {
			if strings.HasPrefix(k, "commander") {
				t.Errorf("row %q waits for %q, and the commander offer waits for row %q", id, k, id)
			}
		}
	}
}

// TestNoOfferBesideAPreferenceQuestion walks D-669 through the agent. The
// reader names the format and the bracket and no theme, builds from an
// owned-only pool, and the classifier reads a request for a suggestion.
// The first turn asks the theme and the colors and offers nothing. The
// next turn answers both, and the offer goes out alone.
//
// The pool is owned-only on purpose. An any-card pool asks the budget,
// which takes the third place of the turn (MaxPerTurn) and holds the
// offer back with no rule at all.
func TestNoOfferBesideAPreferenceQuestion(t *testing.T) {
	var first classifyOut
	first.Format, first.Power, first.PoolRule = "commander", "bracket 5", "owned_only"
	first.Facts.WantsSuggestion = true
	second := commanderClassify()
	second.Power, second.PoolRule = "bracket 5", "owned_only"
	h := &fakeHints{commanders: []string{"Vivi Ornitier", "Hapatra, Vizier of Poisons", "Ghalta, Primal Hunger"}}
	a, _ := testAgentHints(t, h,
		classifyStep(t, first), fits(t, "theme", "colors"), askStep(t),
		classifyStep(t, second))
	st := NewState(true)
	res, err := a.Turn(context.Background(), st, "Build a bracket-5 commander deck", nil)
	if err != nil {
		t.Fatalf("turn 1: %v", err)
	}
	if question(res.Questions, "theme") == nil || question(res.Questions, "colors") == nil {
		t.Fatalf("turn 1 did not ask the theme and the colors, it asked %d questions", len(res.Questions))
	}
	if q := question(res.Questions, "commander"); q != nil {
		t.Fatalf("turn 1 offered commanders beside the preference questions: %q", q.GetText())
	}
	res, err = a.Turn(context.Background(), st, "Lifegain, in white and black.", nil)
	if err != nil {
		t.Fatalf("turn 2: %v", err)
	}
	q := question(res.Questions, "commander")
	if q == nil {
		t.Fatal("turn 2 did not offer a commander after both answers")
	}
	for _, name := range h.commanders {
		if !strings.Contains(q.GetText(), name) {
			t.Errorf("the offer does not name %q: %q", name, q.GetText())
		}
	}
	for _, slot := range []string{"theme", "colors"} {
		if question(res.Questions, slot) != nil {
			t.Errorf("turn 2 asked the %s again beside the offer", slot)
		}
	}
}
