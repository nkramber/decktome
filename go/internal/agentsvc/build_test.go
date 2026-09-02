package agentsvc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/precons"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
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
	got   generate.Request
	res   *generate.Result
	err   error
	runs  int
	block bool
	// started and release gate one build, so a test can act while a
	// build is in flight. started closes once.
	started   chan struct{}
	startOnce sync.Once
	release   chan struct{}
	// record makes the fake spend one generate call on the accumulator,
	// as the real generator does, so a test can read the session total
	// (D-447).
	record bool
}

func (f *fakeDecks) Build(ctx context.Context, req generate.Request, acc *llm.Accumulator) (*generate.Result, error) {
	f.runs++
	f.got = req
	if f.record && acc != nil {
		acc.Record(llm.RoleGenerate, "fake-generate", &llm.Usage{InputTokens: 15000, OutputTokens: 3000}, 0)
	}
	if f.started != nil {
		f.startOnce.Do(func() { close(f.started) })
	}
	if f.release != nil {
		<-f.release
		// A provider call on a cancelled context fails, as the real one
		// does.
		if err := ctx.Err(); err != nil {
			return nil, err
		}
	}
	// block waits for the caller's deadline, which is how a slow provider
	// looks from here.
	if f.block {
		<-ctx.Done()
		return nil, ctx.Err()
	}
	// The real builder stamps the deck with the id the caller reserved.
	// A fake that skips it hides the store, and storeDeck's own guard
	// then makes a failing test look green (D-245).
	if f.res != nil && f.res.Deck != nil {
		f.res.Deck.Id = req.DeckID
	}
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

// TestReadySessionStreamsTheDeck: the question workflow reports READY,
// and the deck reaches the user on the same turn.
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
	fd := &fakeDecks{err: errors.New("the model is down")}
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

// TestReadyWithoutAGeneratorSaysSo covers a deployment that wires no
// generator: the turn says so and builds nothing.
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

// TestBuildTimeoutEndsTheTurnCleanly is D-235. The llm client caps each
// call at three minutes, so a generate and a repair together can hold the
// stream for six. A build past its limit must end the turn with a status
// line, and never lose the turn or hang the stream.
func TestBuildTimeoutEndsTheTurnCleanly(t *testing.T) {
	store := newFakeStore()
	fd := &fakeDecks{block: true}
	opts := append(buildOpts(t, fd), WithBuildTimeout(50*time.Millisecond))
	client, _ := testServerOpts(t, store, opts, readySteps(t)...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})

	if fd.runs != 1 {
		t.Fatalf("build runs = %d, want 1", fd.runs)
	}
	if second.deck != nil {
		t.Error("a timed-out build sent a deck")
	}
	if second.failure != nil {
		t.Errorf("a timed-out build ended the turn: %v", second.failure)
	}
	var said bool
	for _, s := range second.statuses {
		if strings.Contains(s, "time limit") {
			said = true
		}
	}
	if !said {
		t.Errorf("the turn did not say the build ran past its limit: %v", second.statuses)
	}
	// The turn is still stored, so the questions are not lost.
	if len(store.sessions[first.started].GetTurns()) != 2 {
		t.Error("the turn was lost")
	}
}

// fakeDeckStore records what the build kept.
type fakeDeckStore struct {
	n    int
	put  []*mtgv1.Deck
	fail error
}

func (f *fakeDeckStore) NewID(string) string {
	f.n++
	return fmt.Sprintf("deck-%d", f.n)
}

func (f *fakeDeckStore) Get(_ context.Context, _ string, id string) (*mtgv1.Deck, error) {
	for _, d := range f.put {
		if d.GetId() == id {
			return d, nil
		}
	}
	return nil, fmt.Errorf("deck %s not found", id)
}

func (f *fakeDeckStore) Put(ctx context.Context, _ string, d *mtgv1.Deck) error {
	// A write on a cancelled context fails, as Firestore does.
	if err := ctx.Err(); err != nil {
		return err
	}
	if f.fail != nil {
		return f.fail
	}
	f.put = append(f.put, d)
	return nil
}

// TestTheDeckIsKeptAndRecorded is D-245: the built deck is stored and
// its id lands on the session.
func TestTheDeckIsKeptAndRecorded(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	deck := &mtgv1.Deck{Summary: "a deck", Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}}
	opts := append(buildOpts(t, fd), WithDeckStore(ds))
	client, _ := testServerOpts(t, store, opts, readySteps(t)...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})

	if len(ds.put) != 1 {
		t.Fatalf("decks kept = %d, want 1", len(ds.put))
	}
	// The build reserves the id before it runs, because a deck carries
	// its own id.
	if fd.got.DeckID == "" {
		t.Error("the build was given no deck id")
	}
	// The id must reach the stored session, or the next turn cannot know
	// a deck exists.
	got := store.sessions[first.started].GetDeckIds()
	if len(got) != 1 || got[0] != fd.got.DeckID {
		t.Errorf("session deck ids = %v, want the built deck's id %q", got, fd.got.DeckID)
	}
}

// TestAStoreFailureKeepsTheDeck covers the failure path. The user is
// already reading the deck, so a store that refuses it must not take it
// away.
func TestAStoreFailureKeepsTheDeck(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{fail: context.DeadlineExceeded}
	deck := &mtgv1.Deck{Summary: "a deck", Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}}
	opts := append(buildOpts(t, fd), WithDeckStore(ds))
	client, _ := testServerOpts(t, store, opts, readySteps(t)...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})

	if ds.n != 1 {
		t.Fatalf("the store reserved %d ids, want 1: the failure path was never reached", ds.n)
	}
	if second.deck == nil {
		t.Error("a store failure took the deck away from the user")
	}
	if second.failure != nil {
		t.Errorf("a store failure ended the turn: %v", second.failure)
	}
	if ids := store.sessions[first.started].GetDeckIds(); len(ids) != 0 {
		t.Errorf("a deck that was not stored was recorded on the session: %v", ids)
	}
}

// TestNoDeckStoreStillBuilds keeps a deployment without a store working.
func TestNoDeckStoreStillBuilds(t *testing.T) {
	store := newFakeStore()
	deck := &mtgv1.Deck{Summary: "a deck", Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}}
	client, _ := testServerOpts(t, store, buildOpts(t, fd), readySteps(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil {
		t.Error("a server with no deck store sent no deck")
	}
}

// TestDeckColorsFollowEveryCommander: the basics follow every
// commander, so a partner pair gets both halves.
func TestDeckColorsFollowEveryCommander(t *testing.T) {
	w := &mtgv1.Card{ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W}}
	ug := &mtgv1.Card{ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_G}}
	for _, tc := range []struct {
		name       string
		format     mtgv1.FormatId
		chosen     []mtgv1.Color
		commanders []*mtgv1.Card
		want       []mtgv1.Color
	}{
		{"a partner pair is the union", mtgv1.FormatId_FORMAT_ID_COMMANDER, nil, []*mtgv1.Card{w, ug},
			[]mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_G}},
		{"one commander is its identity", mtgv1.FormatId_FORMAT_ID_COMMANDER,
			[]mtgv1.Color{mtgv1.Color_COLOR_R}, []*mtgv1.Card{ug},
			[]mtgv1.Color{mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_G}},
		{"a 60-card session keeps its colors", mtgv1.FormatId_FORMAT_ID_MODERN,
			[]mtgv1.Color{mtgv1.Color_COLOR_R}, nil, []mtgv1.Color{mtgv1.Color_COLOR_R}},
		{"a 60-card session with no colors gets every basic", mtgv1.FormatId_FORMAT_ID_MODERN,
			nil, nil, generate.AllColors},
		{"a Commander session with nothing gets none", mtgv1.FormatId_FORMAT_ID_COMMANDER, nil, nil, nil},
	} {
		got := deckColors(tc.format, tc.chosen, tc.commanders)
		if fmt.Sprint(got) != fmt.Sprint(tc.want) {
			t.Errorf("%s: colors = %v, want %v", tc.name, got, tc.want)
		}
	}
}

// buildServer makes a server with the build wired over the given index,
// for tests that call buildDeck without a stream.
func buildServer(t *testing.T, fd *fakeDecks, idx *cards.Index, extra ...Option) *Server {
	t.Helper()
	cat, err := questions.Load()
	if err != nil {
		t.Fatalf("catalog: %v", err)
	}
	client, _ := fakeClient(t)
	cb, err := candidates.New()
	if err != nil {
		t.Fatalf("candidates: %v", err)
	}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	opts := append([]Option{WithLogger(quiet), WithDecks(fd), WithCandidates(fixedIndex{idx}, cb)}, extra...)
	srv, err := New(cat, client, newFakeStore(), func(context.Context) string { return "u1" }, opts...)
	if err != nil {
		t.Fatalf("server: %v", err)
	}
	return srv
}

// TestThinCommanderPoolRunsNoBuild is D-232. A Commander session that
// delegated the commander, over a library with none for the theme, runs
// no build: no model call is spent, and the error names the reason.
func TestThinCommanderPoolRunsNoBuild(t *testing.T) {
	// No legendary creature at all, so the commander pool is empty.
	welcome := &mtgv1.Card{
		OracleId: "o-welcome", Name: "Ajani's Welcome", TypeLine: "Enchantment",
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W},
		Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	idx := cards.NewIndex([]*mtgv1.Card{welcome}, nil, nil, time.Unix(1000, 0).UTC())
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	srv := buildServer(t, fd, idx)
	session := &mtgv1.Session{Id: "s-1", Slots: &mtgv1.Slots{
		Format:   &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Theme:    "lifegain",
		Colors:   []mtgv1.Color{mtgv1.Color_COLOR_W},
		PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
	}}
	st := &questions.State{Slots: session.GetSlots()}
	res, err := srv.buildDeck(context.Background(), "u1", session, st, nil, nil, nil)
	if !errors.Is(err, ErrThinCommanderPool) {
		t.Fatalf("err = %v, want ErrThinCommanderPool", err)
	}
	if res != nil || fd.runs != 0 {
		t.Errorf("a build ran: res %v, runs %d", res, fd.runs)
	}
}

// TestUnresolvedPreconKeepsNoShare pins the unresolved-precon rule. A
// precon with rows the index could not answer gives a wrong share, so
// the build runs without it. A one-card index resolves no precon, so
// every one is such a list here.
func TestUnresolvedPreconKeepsNoShare(t *testing.T) {
	karlov := &mtgv1.Card{
		OracleId: "o-karlov", Name: "Karlov of the Ghost Council",
		TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true,
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B},
		Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	idx := cards.NewIndex([]*mtgv1.Card{karlov}, nil, nil, time.Unix(1000, 0).UTC())
	set, err := precons.Load(idx)
	if err != nil {
		t.Fatalf("precons: %v", err)
	}
	if len(set.Unresolved()) == 0 {
		t.Fatal("the one-card index resolved a precon, so the test proves nothing")
	}
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	srv := buildServer(t, fd, idx, WithPreconSource(&fakePrecons{set: set}))
	session := &mtgv1.Session{Id: "s-1", Slots: &mtgv1.Slots{
		Format:   &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Theme:    "lifegain",
		PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD,
	}}
	st := &questions.State{Slots: session.GetSlots(), CommanderNames: []string{"Karlov of the Ghost Council"}}
	st.Ctx.Precon = true
	st.Ctx.Words = "upgrade my avengers assemble precon"
	if _, err := srv.buildDeck(context.Background(), "u1", session, st, nil, nil, nil); err != nil {
		t.Fatalf("build: %v", err)
	}
	if fd.runs != 1 {
		t.Fatalf("build runs = %d, want 1", fd.runs)
	}
	if fd.got.Precon != "" || len(fd.got.PreconOracleIDs) != 0 {
		t.Errorf("an unresolved precon reached the share rule: %q with %d cards", fd.got.Precon, len(fd.got.PreconOracleIDs))
	}
}

// TestBuildCopiesTheHouseRules: the house-rules slot reaches the build
// request, and the deck's Format carries it from there (D-3).
func TestBuildCopiesTheHouseRules(t *testing.T) {
	welcome := &mtgv1.Card{
		OracleId: "o-welcome", Name: "Ajani's Welcome", TypeLine: "Enchantment",
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W},
		Legalities:    map[string]mtgv1.LegalityStatus{"modern": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	idx := cards.NewIndex([]*mtgv1.Card{welcome}, nil, nil, time.Unix(1000, 0).UTC())
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	srv := buildServer(t, fd, idx)
	session := &mtgv1.Session{Id: "s-1", Slots: &mtgv1.Slots{
		Format:     &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN},
		Theme:      "lifegain",
		Colors:     []mtgv1.Color{mtgv1.Color_COLOR_W},
		PoolRule:   mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		HouseRules: "any card, no ban list",
	}}
	st := &questions.State{Slots: session.GetSlots()}
	if _, err := srv.buildDeck(context.Background(), "u1", session, st, nil, nil, nil); err != nil {
		t.Fatalf("build: %v", err)
	}
	if fd.runs != 1 {
		t.Fatalf("build runs = %d, want 1", fd.runs)
	}
	if fd.got.HouseRules != "any card, no ban list" {
		t.Errorf("the build request carries house rules %q, want the slot value", fd.got.HouseRules)
	}
}

// TestShortlistFollowsTheCommanderIdentity is D-289. A user who says "any
// colors are fine" leaves the color slot empty, and the shortlist must
// still hold the commander's identity and nothing outside it.
func TestShortlistFollowsTheCommanderIdentity(t *testing.T) {
	store := newFakeStore()
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	// A green lifegain card, so the theme ranks it and only the color
	// identity can keep it out.
	elves := &mtgv1.Card{
		OracleId: "o-elves", Name: "Llanowar Elves", TypeLine: "Creature — Elf Druid",
		OracleText: "Lifelink. Whenever you gain life, put a +1/+1 counter on Llanowar Elves.", Keywords: []string{"Lifelink"},
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_G}, CardTypes: []string{"Creature"}, EdhrecRank: 10,
		Legalities: map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	karlov := &mtgv1.Card{
		OracleId: "o-karlov", Name: "Karlov of the Ghost Council",
		TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true,
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B}, CardTypes: []string{"Creature"},
		Legalities: map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	welcome := &mtgv1.Card{
		OracleId: "o-welcome", Name: "Ajani's Welcome", TypeLine: "Enchantment",
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W}, CardTypes: []string{"Enchantment"},
		Legalities: map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	idx := cards.NewIndex([]*mtgv1.Card{karlov, welcome, elves}, nil, nil, time.Unix(1000, 0).UTC())
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	steps := []llm.Step{
		classifyJSON(t, map[string]any{
			"format": "commander", "theme": "lifegain", "colors": []string{}, "pool_rule": "any_card", "budget_usd": 50,
		}),
		scoreJSON(t, "commander", "power_commander"),
		askJSON(t),
		classifyJSON(t, map[string]any{"power": "bracket 3", "commander_names": []string{"Karlov of the Ghost Council"}}),
	}
	client, _ := testServerOpts(t, store, []Option{WithDecks(fd), WithCandidates(fixedIndex{idx}, cb)}, steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck, any colors are fine"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil || fd.got.Pool == nil {
		t.Fatalf("no build: %v", second.order)
	}
	if _, ok := fd.got.Pool.ByOracleID("o-elves"); ok {
		t.Error("a green card reached the pool of a W/B commander")
	}
	if _, ok := fd.got.Pool.ByOracleID("o-karlov"); !ok {
		t.Error("the commander left the pool")
	}
}

// fakeCollections answers the owned counts and the owned printings of
// one collection. countsErr is what OracleCounts fails with, and reads
// counts how often it was called.
type fakeCollections struct {
	counts    map[string]int32
	printings map[string][]string
	countsErr error
	reads     *int
}

func (f fakeCollections) OracleCounts(context.Context, string, string) (map[string]int32, error) {
	if f.reads != nil {
		*f.reads++
	}
	return f.counts, f.countsErr
}

func (f fakeCollections) OwnedPrintings(context.Context, string, string) (map[string][]string, error) {
	return f.printings, nil
}

// TestOwnedCardShowsThePriciestOwnedPrinting is D-299.
func TestOwnedCardShowsThePriciestOwnedPrinting(t *testing.T) {
	store := newFakeStore()
	deck := &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}, Cards: []*mtgv1.DeckCard{
		{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 1, Owned: true, OwnedCount: 2},
		{OracleId: "o-karlov", Name: "Karlov of the Ghost Council", Count: 1, Owned: false},
	}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}}
	karlov := &mtgv1.Card{
		OracleId: "o-karlov", Name: "Karlov of the Ghost Council", TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true,
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B},
		Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	welcome := &mtgv1.Card{
		OracleId: "o-welcome", Name: "Ajani's Welcome", TypeLine: "Enchantment",
		ColorIdentity:   []mtgv1.Color{mtgv1.Color_COLOR_W},
		Legalities:      map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
		DefaultPrinting: &mtgv1.Printing{ScryfallId: "p-default"},
	}
	printings := []cards.Printing{
		{ScryfallID: "p-cheap", OracleID: "o-welcome", Name: "Ajani's Welcome", ImageUris: &mtgv1.ImageUris{Normal: "https://x/cheap.jpg"}, PriceUSD: 0.5},
		{ScryfallID: "p-dear", OracleID: "o-welcome", Name: "Ajani's Welcome", ImageUris: &mtgv1.ImageUris{Normal: "https://x/dear.jpg"}, PriceUSD: 12},
		{ScryfallID: "p-noimg", OracleID: "o-welcome", Name: "Ajani's Welcome", PriceUSD: 99},
	}
	idx := cards.NewIndex([]*mtgv1.Card{karlov, welcome}, printings, nil, time.Unix(1000, 0).UTC())
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	cols := fakeCollections{
		counts:    map[string]int32{"o-welcome": 2},
		printings: map[string][]string{"o-welcome": {"p-cheap", "p-dear", "p-noimg"}},
	}
	client, _ := testServerOpts(t, store, []Option{WithDecks(fd), WithCandidates(fixedIndex{idx}, cb), WithCollections(cols)}, readySteps(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck", CollectionId: "c1"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil {
		t.Fatalf("no deck: %v", second.order)
	}
	got := second.deck.GetCards()[0].GetOwnedPrinting()
	if got.GetScryfallId() != "p-dear" || got.GetImageUris().GetNormal() != "https://x/dear.jpg" || got.GetPriceUsd() != 12 {
		t.Errorf("owned printing = %v, want the dear one with an image", got)
	}
	if second.deck.GetCards()[1].GetOwnedPrinting() != nil {
		t.Error("an unowned card got an owned printing")
	}
}

// TestTurnDuringABuildIsRefused is D-303. A second message while a
// build runs must not start a second billed build.
func TestTurnDuringABuildIsRefused(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	deck := &mtgv1.Deck{Summary: "a deck", Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}, started: make(chan struct{}), release: make(chan struct{})}
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), readySteps(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})

	done := make(chan events, 1)
	go func() {
		done <- chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	}()
	select {
	case <-fd.started:
	case <-time.After(5 * time.Second):
		t.Fatal("the build never started")
	}
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{SessionId: first.started, Message: "another one"}))
	if err == nil {
		for stream.Receive() {
		}
		err = stream.Err()
		_ = stream.Close()
	}
	if connect.CodeOf(err) != connect.CodeAborted || !strings.Contains(err.Error(), "a build is in progress") {
		t.Errorf("a turn during a build gave %v, want Aborted with the reason: %v", connect.CodeOf(err), err)
	}
	close(fd.release)
	second := <-done
	if second.deck == nil || fd.runs != 1 {
		t.Errorf("the first build did not finish alone: deck %v, runs %d", second.deck != nil, fd.runs)
	}
	// The marker is gone, so the next turn is admitted.
	stream, err = client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{SessionId: first.started, Message: "thanks"}))
	if err == nil {
		for stream.Receive() {
		}
		err = stream.Err()
		_ = stream.Close()
	}
	if connect.CodeOf(err) == connect.CodeAborted {
		t.Errorf("the build marker outlived the build: %v", err)
	}
}

// TestDisconnectMidBuildStillStoresTheDeck is D-303. The client leaves
// after the paid call started, and the deck is stored and recorded.
func TestDisconnectMidBuildStillStoresTheDeck(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	deck := &mtgv1.Deck{Summary: "a deck", Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}, started: make(chan struct{}), release: make(chan struct{})}
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), readySteps(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		defer close(done)
		stream, err := client.Chat(ctx, connect.NewRequest(&mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"}))
		if err != nil {
			return
		}
		for stream.Receive() {
		}
		_ = stream.Close()
	}()
	select {
	case <-fd.started:
	case <-time.After(5 * time.Second):
		close(fd.release)
		t.Fatal("the build never started")
	}
	cancel()
	<-done
	close(fd.release)
	deadline := time.After(5 * time.Second)
	for len(store.deckIDs(first.started)) == 0 {
		select {
		case <-deadline:
			t.Fatalf("the deck id never reached the session: decks kept %d", len(ds.put))
		case <-time.After(20 * time.Millisecond):
		}
	}
	if got := store.status(first.started); got != mtgv1.SessionStatus_SESSION_STATUS_BUILT {
		t.Errorf("status = %v, want BUILT", got)
	}
}

// TestDeckIDWriteRetriesAfterAConflict is D-303. Another write lands
// between the turn's Put and the deck id write, and the id is appended
// to the current session instead of lost.
func TestDeckIDWriteRetriesAfterAConflict(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	deck := &mtgv1.Deck{Summary: "a deck", Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}}
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), readySteps(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	store.onPut = func(s *mtgv1.Session) {
		if len(s.GetDeckIds()) > 0 {
			store.onPut = nil
			store.versions[s.GetId()]++
		}
	}
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil {
		t.Fatalf("no deck: %v", second.order)
	}
	if ids := store.deckIDs(first.started); len(ids) != 1 || ids[0] != fd.got.DeckID {
		t.Errorf("deck ids = %v, want the built deck after the retry", ids)
	}
	if store.status(first.started) != mtgv1.SessionStatus_SESSION_STATUS_BUILT {
		t.Error("the retry lost the built status")
	}
	if got := store.versionOf(first.started); got != 4 {
		t.Errorf("version = %d, want 4: turn, bump, and the retried deck id", got)
	}
}

// TestCollectionCountsAreReadOncePerTurn: the hints and the build share
// one read.
func TestCollectionCountsAreReadOncePerTurn(t *testing.T) {
	store := newFakeStore()
	deck := &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}}
	reads := 0
	cols := fakeCollections{counts: map[string]int32{"o-welcome": 1}, reads: &reads}
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithCollections(cols)), readySteps(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck", CollectionId: "c1"})
	if reads != 1 {
		t.Errorf("the first turn read the counts %d times, want 1", reads)
	}
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil {
		t.Fatalf("no deck: %v", second.order)
	}
	if reads != 2 {
		t.Errorf("the building turn read the counts %d times in total, want 2", reads)
	}
	if fd.got.OracleCounts["o-welcome"] != 1 {
		t.Error("the build did not get the counts")
	}
}

// TestNewSessionWithAMissingCollectionIsNotFound covers a new session
// that names a collection the store does not hold.
func TestNewSessionWithAMissingCollectionIsNotFound(t *testing.T) {
	cols := fakeCollections{countsErr: status.Error(codes.NotFound, "no such document")}
	client, sc := testServerOpts(t, newFakeStore(), []Option{WithCollections(cols)}, firstTurn(t)...)
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: "build me a deck", CollectionId: "gone"}))
	if err == nil {
		for stream.Receive() {
		}
		err = stream.Err()
		_ = stream.Close()
	}
	if connect.CodeOf(err) != connect.CodeNotFound {
		t.Errorf("code = %v, want NotFound: %v", connect.CodeOf(err), err)
	}
	if len(sc.Calls) != 0 {
		t.Errorf("a refused turn made %d model calls", len(sc.Calls))
	}
}

// TestRevisionCapKeepsTheCommanderAndTheLockedCards: a mana value cap
// never drops a commander or a locked card from the pool, because the
// engine blocks a deck without them (D-242).
func TestRevisionCapKeepsTheCommanderAndTheLockedCards(t *testing.T) {
	karlov := &mtgv1.Card{
		OracleId: "o-karlov", Name: "Karlov of the Ghost Council", ManaValue: 9,
		TypeLine: "Legendary Creature — Spirit Advisor", CanBeCommander: true, CardTypes: []string{"Creature"},
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B},
		Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	locked := &mtgv1.Card{
		OracleId: "o-locked", Name: "Locked Thing", ManaValue: 6, TypeLine: "Creature", CardTypes: []string{"Creature"},
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W},
		Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	other := &mtgv1.Card{
		OracleId: "o-other", Name: "Other Thing", ManaValue: 6, TypeLine: "Creature", CardTypes: []string{"Creature"},
		ColorIdentity: []mtgv1.Color{mtgv1.Color_COLOR_W},
		Legalities:    map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL},
	}
	idx := cards.NewIndex([]*mtgv1.Card{karlov, locked, other}, nil, nil, time.Unix(1000, 0).UTC())
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	srv := buildServer(t, fd, idx)
	session := &mtgv1.Session{Id: "s-1", Slots: &mtgv1.Slots{
		Format:     &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Theme:      "lifegain",
		PoolRule:   mtgv1.PoolRule_POOL_RULE_ANY_CARD,
		SlotStates: map[string]mtgv1.SlotState{},
	}}
	st := questions.NewState(false)
	st.Slots = session.GetSlots()
	st.SetCommander("Karlov of the Ghost Council")
	st.AddLocked("Locked Thing")
	rev := &generate.Revision{Base: []*mtgv1.DeckCard{{OracleId: "o-other", Name: "Other Thing", Count: 1}}, MaxManaValue: 4}
	if _, err := srv.buildDeckFrom(context.Background(), "u1", session, st, nil, nil, rev, nil); err != nil {
		t.Fatalf("build: %v", err)
	}
	got := fd.got.Revision
	if got == nil || len(got.Exempt) != 2 || got.Exempt[0] != "o-karlov" || got.Exempt[1] != "o-locked" {
		t.Fatalf("exempt = %v, want the commander and the locked card", got)
	}
	for _, id := range []string{"o-karlov", "o-locked"} {
		if _, ok := fd.got.Pool.ByOracleID(id); !ok {
			t.Errorf("%s left the pool over the cap", id)
		}
	}
	if _, ok := fd.got.Pool.ByOracleID("o-other"); ok {
		t.Error("a base card over the cap stayed in the pool")
	}
}

// TestBuildNamesTheMissingWiring: each missing piece has its own error.
func TestBuildNamesTheMissingWiring(t *testing.T) {
	cat, err := questions.Load()
	if err != nil {
		t.Fatal(err)
	}
	client, _ := fakeClient(t)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	fd := &fakeDecks{}
	session := &mtgv1.Session{Id: "s-1", Slots: &mtgv1.Slots{Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN}}}
	st := &questions.State{Slots: session.GetSlots()}
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name string
		opts []Option
		want error
	}{
		{"no index source", []Option{WithDecks(fd)}, errNoIndexSource},
		{"no builder", []Option{WithDecks(fd), WithCandidates(fixedIndex{}, nil)}, errNoCandidateBuilder},
		{"no index loaded", []Option{WithDecks(fd), WithCandidates(fixedIndex{}, cb)}, errNoIndexLoaded},
	} {
		srv, err := New(cat, client, newFakeStore(), func(context.Context) string { return "u1" }, append([]Option{WithLogger(quiet)}, tc.opts...)...)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := srv.buildDeck(context.Background(), "u1", session, st, nil, nil, nil); !errors.Is(err, tc.want) {
			t.Errorf("%s: err = %v, want %v", tc.name, err, tc.want)
		}
	}
}

// TestSessionSpendHoldsTheBuild is D-447. Session vAvg4eteJhmuPEuJwBul
// showed 9 calls and $0.0023 after a build that costs about $0.06 on
// its own: the turn summed its report before the build and never
// again. The streamed total and the stored total hold the build now.
func TestSessionSpendHoldsTheBuild(t *testing.T) {
	store := newFakeStore()
	deck := &mtgv1.Deck{Summary: "a lifegain deck", Validation: &mtgv1.ValidationResult{}, Cards: []*mtgv1.DeckCard{{Name: "Ajani's Welcome", Count: 1}}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}, record: true}
	// The post-build write of the session runs with a deck store, as in
	// production, and that write is the one that carries the build.
	ds := &fakeDeckStore{}
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), readySteps(t)...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	afterQuestions := first.usage.GetCalls()
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil {
		t.Fatalf("no deck reached the user: %v", second.order)
	}
	// The second turn spends its own question calls and the one build.
	questionCalls := second.usage.GetCalls() - afterQuestions
	if questionCalls < 2 {
		t.Fatalf("the second turn reports %d calls over the first, want the question calls and the build", questionCalls)
	}
	if second.usage.GetInputTokens() < first.usage.GetInputTokens()+15000 {
		t.Errorf("the streamed total holds %d input tokens, want the build's 15000 on top of %d", second.usage.GetInputTokens(), first.usage.GetInputTokens())
	}
	stored := store.sessions[first.started]
	if stored.GetUsage().GetCalls() != second.usage.GetCalls() || stored.GetUsage().GetInputTokens() != second.usage.GetInputTokens() {
		t.Errorf("the stored total (%d calls, %d in) differs from the streamed one (%d calls, %d in)",
			stored.GetUsage().GetCalls(), stored.GetUsage().GetInputTokens(), second.usage.GetCalls(), second.usage.GetInputTokens())
	}
}
