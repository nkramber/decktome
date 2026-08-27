package agentsvc

import (
	"context"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// fixedIndex is an IndexSource over a hand-built snapshot.
type fixedIndex struct{ idx *cards.Index }

func (f fixedIndex) Current() *cards.Index { return f.idx }

// buildOpts wires the generator, a card index, and a candidates builder.
// Without all three the build returns early and the test proves nothing.
func buildOpts(t *testing.T, fd *fakeDecks) []Option {
	t.Helper()
	karlov := &mtgv1.Card{
		OracleId: "o-karlov", Name: "Karlov of the Ghost Council",
		TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true,
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B},
		Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	welcome := &mtgv1.Card{
		OracleId: "o-welcome", Name: "Ajani's Welcome", TypeLine: "Enchantment",
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W},
		Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	idx := cards.NewIndex([]*mtgv1.Card{karlov, welcome}, nil, nil, time.Unix(1000, 0).UTC())
	cb, err := candidates.New()
	if err != nil {
		t.Fatalf("candidates: %v", err)
	}
	return []Option{WithDecks(fd), WithCandidates(fixedIndex{idx}, cb)}
}

// fakeDecks stands in for the generator. It records the request, so the
// test can read what the session handed over.
type fakeDecks struct {
	got  generate.Request
	res  *generate.Result
	err  error
	runs int
}

func (f *fakeDecks) Build(_ context.Context, req generate.Request, _ *llm.Accumulator) (*generate.Result, error) {
	f.runs++
	f.got = req
	return f.res, f.err
}

// readySteps drives a session to READY. It is the recipe of
// TestChatResumesWithoutRepeating: the second turn closes the last slots.
func readySteps(t *testing.T) []llm.Step {
	t.Helper()
	return append(firstTurn(t),
		classifyJSON(t, map[string]any{
			"power":           "bracket 3",
			"commander_names": []string{"Karlov of the Ghost Council"},
		}),
		scoreJSON(t),
		askJSON(t))
}

// TestReadySessionStreamsTheDeck is the PR-8 hand-over. The question
// workflow reports READY, and the deck reaches the user on the same turn.
func TestReadySessionStreamsTheDeck(t *testing.T) {
	store := newFakeStore()
	deck := &mtgv1.Deck{
		Summary:    "a lifegain deck",
		Validation: &mtgv1.ValidationResult{},
		Cards:      []*mtgv1.DeckCard{{Name: "Ajani's Welcome", Count: 1}},
	}
	fd := &fakeDecks{res: &generate.Result{Deck: deck, Notes: []string{"I could not place \"Nonesuch\"."}}}
	client, _ := testServerOpts(t, store, buildOpts(t, fd), readySteps(t)...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	if fd.runs != 0 {
		t.Fatal("the build ran before the session was ready")
	}
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})

	if fd.runs != 1 {
		t.Fatalf("build runs = %d, want 1", fd.runs)
	}
	if second.deck == nil {
		t.Fatalf("no deck reached the user: %v", second.order)
	}
	if second.deck.GetSummary() != "a lifegain deck" {
		t.Errorf("summary = %q", second.deck.GetSummary())
	}
	// A name that missed twice reaches the user as prose.
	if len(second.texts) != 1 {
		t.Errorf("texts = %v, want the one note", second.texts)
	}
	// The session, not the model, decides the session id and the format.
	if fd.got.SessionID != first.started {
		t.Errorf("session id = %q, want %q", fd.got.SessionID, first.started)
	}
	if fd.got.Pool == nil {
		t.Error("the build got no shortlist")
	}
}

// TestBuildFailureKeepsTheTurn covers the failure path. The questions are
// already stored and already sent, so a build error must not lose them.
func TestBuildFailureKeepsTheTurn(t *testing.T) {
	store := newFakeStore()
	fd := &fakeDecks{err: context.DeadlineExceeded}
	client, _ := testServerOpts(t, store, buildOpts(t, fd), readySteps(t)...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})

	if fd.runs != 1 {
		t.Fatalf("build runs = %d, want 1: the failure path was never reached", fd.runs)
	}
	if second.deck != nil {
		t.Error("a failed build sent a deck")
	}
	if second.failure != nil {
		t.Errorf("a failed build ended the turn: %v", second.failure)
	}
	// The session still reached READY and the turn is stored.
	if got := store.sessions[first.started].GetStatus(); got != mtgv1.SessionStatus_SESSION_STATUS_READY {
		t.Errorf("status = %v, want READY", got)
	}
	if len(store.sessions[first.started].GetTurns()) != 2 {
		t.Error("the turn was lost")
	}
}

// TestReadyWithoutAGeneratorSaysSo keeps the PR-7 behavior for a
// deployment that wires no generator.
func TestReadyWithoutAGeneratorSaysSo(t *testing.T) {
	store := newFakeStore()
	client, _ := testServer(t, store, readySteps(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck != nil {
		t.Error("a server with no generator sent a deck")
	}
	var said bool
	for _, s := range second.statuses {
		if s != "" {
			said = true
		}
	}
	if !said {
		t.Errorf("the turn said nothing about the missing generator: %v", second.statuses)
	}
}
