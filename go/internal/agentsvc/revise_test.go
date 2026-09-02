package agentsvc

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// reviseJSON scripts one answer of the revise role.
func reviseJSON(t *testing.T, fields map[string]any) llm.Step {
	t.Helper()
	out := map[string]any{
		"changes": []string{}, "remove": []string{}, "keep": []string{},
		"max_mana_value": 0.0, "swap_basics": 0, "land_kinds": "", "question": "", "declined": []any{},
	}
	for k, v := range fields {
		out[k] = v
	}
	raw, err := json.Marshal(out)
	if err != nil {
		t.Fatal(err)
	}
	return llm.Step{Output: raw, Usage: &llm.Usage{InputTokens: 400, OutputTokens: 60}}
}

// builtSteps scripts the two turns to the first deck: the opening, then
// the answer that closes the last slots. The second turn asks nothing,
// so it spends one classify step only.
func builtSteps(t *testing.T) []llm.Step {
	t.Helper()
	return append(firstTurn(t), classifyJSON(t, map[string]any{
		"power":           "bracket 3",
		"commander_names": []string{"Karlov of the Ghost Council"},
	}))
}

// TestRevisionTurnKeepsTheBaseAndReportsTheDiff is D-283. A message after
// a build with no slot change revises the deck the user read.
func TestRevisionTurnKeepsTheBaseAndReportsTheDiff(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	base := &mtgv1.Deck{
		Name: "lifegain Commander", Summary: "a lifegain deck", Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 1}, {OracleId: "o-plains", Name: "Plains", Count: 30}},
	}
	fd := &fakeDecks{res: &generate.Result{Deck: base}}
	steps := append(builtSteps(t),
		classifyJSON(t, nil),
		reviseJSON(t, map[string]any{
			"changes":        []string{"Cut two Plains and add a card that draws cards"},
			"max_mana_value": 5.0,
			"declined":       []map[string]string{{"request": "Replace some lands with better options", "reason": "for a casual mono-white deck, all basic lands is fine"}},
		}))
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil || fd.runs != 1 {
		t.Fatalf("no first deck: runs=%d order=%v", fd.runs, second.order)
	}
	revised := &mtgv1.Deck{
		Name: "x", Summary: "a leaner lifegain deck", Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 1}, {OracleId: "o-plains", Name: "Plains", Count: 28}, {OracleId: "o-karlov", Name: "Karlov of the Ghost Council", Count: 1}},
	}
	fd.res = &generate.Result{Deck: revised}

	third := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Replace some lands with better options. Also no 6 or 7 mana cards."})
	if fd.runs != 2 {
		t.Fatalf("build runs = %d, want 2", fd.runs)
	}
	rev := fd.got.Revision
	if rev == nil {
		t.Fatal("the build got no revision brief")
	}
	if rev.BaseDeckID != "deck-1" || len(rev.Base) != 2 || rev.MaxManaValue != 5 {
		t.Errorf("brief = %+v", rev)
	}
	if third.deck == nil {
		t.Fatalf("no revised deck: %v", third.order)
	}
	if third.deck.GetRevisedFromDeckId() != "deck-1" {
		t.Errorf("revised_from = %q", third.deck.GetRevisedFromDeckId())
	}
	if third.deck.GetName() != "lifegain Commander" {
		t.Errorf("the revision renamed the deck: %q", third.deck.GetName())
	}
	note := strings.Join(third.texts, " ")
	for _, want := range []string{"I added 1 Karlov of the Ghost Council.", "I changed the count of Plains: 30 to 28.", "I did not replace some lands with better options: for a casual mono-white deck, all basic lands is fine."} {
		if !strings.Contains(note, want) {
			t.Errorf("note lacks %q: %q", want, note)
		}
	}
	if third.deck.GetRevisionNote() != note {
		t.Errorf("deck note = %q", third.deck.GetRevisionNote())
	}
	if !strings.Contains(strings.Join(third.statuses, "|"), "revising the deck") {
		t.Errorf("statuses = %v", third.statuses)
	}
	// The reply is stored on the turn, and the deck id is recorded.
	s := store.sessions[first.started]
	if got := s.GetTurns()[len(s.GetTurns())-1].GetAgentMessage(); got != note {
		t.Errorf("stored agent_message = %q", got)
	}
	if len(s.GetDeckIds()) != 2 {
		t.Errorf("deck ids = %v", s.GetDeckIds())
	}
}

// TestRevisionQuestionEndsTheTurn is D-284: an unclear request gets one
// question and no build.
func TestRevisionQuestionEndsTheTurn(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}, Cards: []*mtgv1.DeckCard{{OracleId: "o-plains", Name: "Plains", Count: 30}}}}}
	steps := append(builtSteps(t),
		classifyJSON(t, nil),
		reviseJSON(t, map[string]any{"question": "Do you mean faster mana, utility lands, or more colors?"}))
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	third := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Better lands please"})
	if fd.runs != 1 {
		t.Errorf("a question ran a build: runs=%d", fd.runs)
	}
	if len(third.questions) != 1 || third.questions[0].GetText() != "Do you mean faster mana, utility lands, or more colors?" || !third.questions[0].GetInvented() {
		t.Fatalf("questions = %v", third.questions)
	}
	if third.deck != nil || len(third.texts) != 0 {
		t.Errorf("a question turn sent a deck or prose: %v", third.order)
	}
	s := store.sessions[first.started]
	last := s.GetTurns()[len(s.GetTurns())-1]
	if len(last.GetQuestions()) != 1 || last.GetAgentMessage() == "" {
		t.Errorf("the question was not stored on the turn: %v", last)
	}
}

// TestRevisionDeclineIsNotSilent is D-284: a request with no change gets
// a reason, and the deck stays.
func TestRevisionDeclineIsNotSilent(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}, Cards: []*mtgv1.DeckCard{{OracleId: "o-plains", Name: "Plains", Count: 30}}}}}
	steps := append(builtSteps(t),
		classifyJSON(t, nil),
		reviseJSON(t, map[string]any{"declined": []map[string]string{{"request": "Replace the lands", "reason": "all basic lands is fine here"}}}))
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	third := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Replace the lands"})
	if fd.runs != 1 || third.deck != nil {
		t.Errorf("a decline built: runs=%d deck=%v", fd.runs, third.deck != nil)
	}
	if len(third.texts) != 1 || !strings.Contains(third.texts[0], "I did not replace the lands: all basic lands is fine here. The deck stays as it was.") {
		t.Errorf("texts = %v", third.texts)
	}
	if third.usage == nil {
		t.Error("the turn sent no usage")
	}
}

// TestSlotChangeAfterBuildRebuilds is D-241: a changed setting means a
// full build, and the user reads why.
func TestSlotChangeAfterBuildRebuilds(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}, Cards: []*mtgv1.DeckCard{{OracleId: "o-plains", Name: "Plains", Count: 30}}}}}
	steps := append(builtSteps(t), classifyJSON(t, map[string]any{"power": "bracket 2"}))
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	third := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Make it bracket 2"})
	if fd.runs != 2 || fd.got.Revision != nil {
		t.Errorf("runs=%d revision=%v", fd.runs, fd.got.Revision)
	}
	if !strings.Contains(strings.Join(third.statuses, "|"), "a deck setting changed") {
		t.Errorf("statuses = %v", third.statuses)
	}
	if third.deck == nil || third.deck.GetRevisedFromDeckId() != "" {
		t.Errorf("deck = %v", third.deck)
	}
}

// TestRevisionUnlocksARemovedCard is D-301. The classify call reads a
// named card as a card to keep, the brief removes it, and the lock must
// go with it.
func TestRevisionUnlocksARemovedCard(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	base := &mtgv1.Deck{Name: "lifegain Commander", Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 1}, {OracleId: "o-plains", Name: "Plains", Count: 30}}}
	fd := &fakeDecks{res: &generate.Result{Deck: base}}
	steps := append(builtSteps(t),
		classifyJSON(t, map[string]any{"locked_names": []string{"Ajani's Welcome"}}),
		reviseJSON(t, map[string]any{"changes": []string{"Replace Ajani's Welcome with a card the user does not own"}, "remove": []string{"Ajani's Welcome"}}))
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	fd.res = &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}, Cards: []*mtgv1.DeckCard{{OracleId: "o-plains", Name: "Plains", Count: 31}}}}
	third := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Replace Ajani's Welcome with a card I do not own"})
	if third.deck == nil {
		t.Fatalf("no revised deck: %v", third.order)
	}
	if len(fd.got.Locked) != 0 {
		t.Errorf("the removed card stayed locked: %v", fd.got.Locked)
	}
	if _, ok := fd.got.Pool.ByOracleID("o-welcome"); ok {
		t.Error("the removed card stayed in the pool")
	}
	// The stored state dropped the lock too.
	snap := store.states[first.started]
	for _, n := range snap.LockedNames {
		if n == "Ajani's Welcome" {
			t.Error("the stored state still locks the removed card")
		}
	}
}

// TestRevisionSwapsBasicsAndKeepsTheBrief is D-448 and D-449. The swap
// count of the brief reaches the generator, fitted to the nonbasic
// lands the pool offers, and the turn stores the brief.
func TestRevisionSwapsBasicsAndKeepsTheBrief(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	base := &mtgv1.Deck{Name: "lifegain Commander", Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 1}, {OracleId: "o-plains", Name: "Plains", Count: 30}}}
	fd := &fakeDecks{res: &generate.Result{Deck: base}}
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	karlov := &mtgv1.Card{OracleId: "o-karlov", Name: "Karlov of the Ghost Council", TypeLine: "Legendary Creature — Spirit Advisor",
		CanBeCommander: true, ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B}, Legalities: legal}
	welcome := &mtgv1.Card{OracleId: "o-welcome", Name: "Ajani's Welcome", TypeLine: "Enchantment", ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W}, Legalities: legal}
	plains := &mtgv1.Card{OracleId: "o-plains", Name: "Plains", TypeLine: "Basic Land — Plains", CardTypes: []string{"Land"}, Supertypes: []string{"Basic"}, Legalities: legal}
	tower := &mtgv1.Card{OracleId: "o-tower", Name: "Command Tower", TypeLine: "Land", CardTypes: []string{"Land"}, Legalities: legal}
	idx := cards.NewIndex([]*mtgv1.Card{karlov, welcome, plains, tower}, nil, nil, time.Unix(1000, 0).UTC())
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	opts := []Option{WithDecks(fd), WithCandidates(fixedIndex{idx}, cb), WithDeckStore(ds)}
	steps := append(builtSteps(t),
		classifyJSON(t, nil),
		reviseJSON(t, map[string]any{
			"changes": []string{"Replace basic lands with dual lands"}, "swap_basics": 12, "land_kinds": "dual lands that enter untapped",
		}))
	client, _ := testServerOpts(t, store, opts, steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	fd.res = &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}, Cards: []*mtgv1.DeckCard{{OracleId: "o-plains", Name: "Plains", Count: 29}, {OracleId: "o-tower", Name: "Command Tower", Count: 1}}}}
	third := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Add better lands instead of the basics"})
	if third.deck == nil {
		t.Fatalf("no revised deck: %v", third.order)
	}
	rev := fd.got.Revision
	if rev == nil {
		t.Fatal("the generator got no revision")
	}
	// The pool offers one nonbasic land the base does not hold, so the
	// ask of 12 fits to 1.
	if rev.SwapBasics != 1 || rev.LandKinds != "dual lands that enter untapped" {
		t.Errorf("revision swap = %d %q, want 1 and the kinds", rev.SwapBasics, rev.LandKinds)
	}
	session := store.sessions[first.started]
	turns := session.GetTurns()
	last := turns[len(turns)-1]
	if !strings.Contains(last.GetRevisionBrief(), `"swap_basics":12`) {
		t.Errorf("the turn did not keep the brief: %q", last.GetRevisionBrief())
	}
	for _, turn := range turns[:len(turns)-1] {
		if turn.GetRevisionBrief() != "" {
			t.Errorf("a turn with no revise call holds a brief: %q", turn.GetRevisionBrief())
		}
	}
}
