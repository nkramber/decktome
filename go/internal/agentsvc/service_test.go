package agentsvc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/precons"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
)

// fakeStore is an in-memory Store. It keeps the two halves apart, the
// same way Firestore does, so a test can prove the private state moves.
type fakeStore struct {
	mu       sync.Mutex
	sessions map[string]*mtgv1.Session
	states   map[string]questions.Snapshot
	versions map[string]int64
	ids      int
	putErr   error
	// onGetState runs after each GetState, so a test can slip a write in
	// between the read and the Put, the way an overlapping turn does.
	onGetState func()
	// onPut runs before each Put's version check, with the session about
	// to be written.
	onPut func(*mtgv1.Session)
}

func newFakeStore() *fakeStore {
	return &fakeStore{
		sessions: map[string]*mtgv1.Session{},
		states:   map[string]questions.Snapshot{},
		versions: map[string]int64{},
	}
}

func (f *fakeStore) NewID(string) string {
	f.ids++
	return "sess-" + string(rune('0'+f.ids))
}

func (f *fakeStore) Put(ctx context.Context, _ string, s *mtgv1.Session, snap questions.Snapshot, expected int64) error {
	// A write on a cancelled context fails, as Firestore does.
	if err := ctx.Err(); err != nil {
		return err
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.putErr != nil {
		return f.putErr
	}
	if f.onPut != nil {
		f.onPut(s)
	}
	if f.versions[s.GetId()] != expected {
		return sessions.ErrConflict
	}
	// A copy, as Firestore stores one. The service goes on editing its
	// own session after the Put.
	f.sessions[s.GetId()] = proto.Clone(s).(*mtgv1.Session)
	f.states[s.GetId()] = snap
	f.versions[s.GetId()] = expected + 1
	return nil
}

func (f *fakeStore) Get(_ context.Context, _, id string) (*mtgv1.Session, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.get(id)
}

func (f *fakeStore) get(id string) (*mtgv1.Session, error) {
	s, ok := f.sessions[id]
	if !ok {
		return nil, sessions.ErrNotFound
	}
	// A copy, as Firestore decodes one. A turn that edits the session
	// in place must not reach the store without a Put.
	return proto.Clone(s).(*mtgv1.Session), nil
}

func (f *fakeStore) GetState(_ context.Context, _, id string) (*mtgv1.Session, questions.Snapshot, int64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	s, err := f.get(id)
	if err != nil {
		return nil, questions.Snapshot{}, 0, err
	}
	snap, version := f.states[id], f.versions[id]
	if f.onGetState != nil {
		f.onGetState()
	}
	return s, snap, version, nil
}

// deckIDs, status, and versionOf read one stored session under the lock,
// for a test that polls while a turn still runs.
func (f *fakeStore) deckIDs(id string) []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.sessions[id].GetDeckIds()
}

func (f *fakeStore) status(id string) mtgv1.SessionStatus {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.sessions[id].GetStatus()
}

func (f *fakeStore) versionOf(id string) int64 {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.versions[id]
}

// classifyJSON writes one classify answer that satisfies the schema.
func classifyJSON(t *testing.T, fields map[string]any) llm.Step {
	t.Helper()
	out := map[string]any{
		"format": "unknown", "theme": "", "colors": []string{},
		"commander_names": []string{}, "locked_names": []string{}, "named_cards": []string{}, "set_names": []string{}, "set_groups": []string{}, "precon_names": []string{},
		"power": "", "pool_rule": "unknown", "budget_usd": 0.0, "budget_scope": "unknown",
		"house_rules": "", "closed_keys": []string{}, "declined_keys": []string{},
		"facts": map[string]bool{
			"named_card": false, "buy_list": false,
			"house_format": false, "budget_ambiguous": false,
			"power_competitive": false, "wants_suggestion": false, "out_of_scope": false, "exclude_precons": false,
		},
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

func scoreJSON(t *testing.T, rowIDs ...string) llm.Step {
	t.Helper()
	scores := make([]map[string]any, 0, len(rowIDs))
	for _, id := range rowIDs {
		scores = append(scores, map[string]any{"row_id": id, "fit": 0.9, "custom_text": "", "reason": "fits"})
	}
	raw, err := json.Marshal(map[string]any{"scores": scores})
	if err != nil {
		t.Fatal(err)
	}
	return llm.Step{Output: raw, Usage: &llm.Usage{InputTokens: 250, OutputTokens: 30}}
}

func askJSON(t *testing.T) llm.Step {
	t.Helper()
	raw, err := json.Marshal(map[string]any{"questions": []any{}})
	if err != nil {
		t.Fatal(err)
	}
	return llm.Step{Output: raw, Usage: &llm.Usage{InputTokens: 300, OutputTokens: 40}}
}

// firstTurn is the script of a Commander opening: the format, the theme,
// and the colors arrive at once, so the agent asks the commander and the
// bracket.
func firstTurn(t *testing.T) []llm.Step {
	t.Helper()
	return []llm.Step{
		classifyJSON(t, map[string]any{
			"format": "commander", "theme": "lifegain",
			"colors": []string{"W", "B"}, "pool_rule": "any_card",
			// The session holds no collection, so every card must be
			// bought and the budget row fires. A named cap closes it, and
			// this test measures the commander rows (D-168). The message
			// names the cap, or the budget never applies (D-537).
			"budget_usd": 50,
		}),
		scoreJSON(t, "commander", "power_commander"),
		askJSON(t),
	}
}

func fakeClient(t *testing.T, steps ...llm.Step) (*llm.Client, *llm.Script) {
	t.Helper()
	roles := map[llm.Role]llm.RoleSpec{}
	for _, r := range llm.Roles {
		roles[r] = llm.RoleSpec{Provider: llm.FakeName, Model: "fake-" + string(r), MaxOutputTokens: 1024}
	}
	sc := llm.NewScript(steps...)
	c, err := llm.New(&llm.Config{VerifiedAt: "2026-08-24", Roles: roles}, []llm.Provider{sc}, llm.WithoutJitter())
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	return c, sc
}

// testServer starts the real Connect handler, because Chat streams and a
// direct call can not make a ServerStream.
func testServer(t *testing.T, store Store, steps ...llm.Step) (mtgv1connect.AgentServiceClient, *llm.Script) {
	t.Helper()
	return testServerOpts(t, store, nil, steps...)
}

// testServerOpts builds the server with extra options, for the build
// wiring.
func testServerOpts(t *testing.T, store Store, extra []Option, steps ...llm.Step) (mtgv1connect.AgentServiceClient, *llm.Script) {
	t.Helper()
	cat, err := questions.Load()
	if err != nil {
		t.Fatalf("catalog: %v", err)
	}
	client, sc := fakeClient(t, steps...)
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	opts := append([]Option{
		WithLogger(quiet), WithClock(func() time.Time { return time.Unix(1000, 0).UTC() }),
	}, extra...)
	srv, err := New(cat, client, store, func(context.Context) string { return "u1" }, opts...)
	if err != nil {
		t.Fatalf("server: %v", err)
	}
	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewAgentServiceHandler(srv))
	httpSrv := httptest.NewServer(mux)
	t.Cleanup(httpSrv.Close)
	return mtgv1connect.NewAgentServiceClient(httpSrv.Client(), httpSrv.URL), sc
}

// events drains one Chat stream.
type events struct {
	started   string
	questions []*mtgv1.Question
	slots     *mtgv1.Slots
	statuses  []string
	usage     *mtgv1.Usage
	failure   *mtgv1.AgentError
	deck      *mtgv1.Deck
	texts     []string
	order     []string
}

func chat(t *testing.T, c mtgv1connect.AgentServiceClient, req *mtgv1.ChatRequest) events {
	t.Helper()
	stream, err := c.Chat(context.Background(), connect.NewRequest(req))
	if err != nil {
		t.Fatalf("chat: %v", err)
	}
	defer func() { _ = stream.Close() }()
	var got events
	for stream.Receive() {
		switch e := stream.Msg().GetEvent().(type) {
		case *mtgv1.ChatResponse_SessionStarted:
			got.started, got.order = e.SessionStarted, append(got.order, "started")
		case *mtgv1.ChatResponse_Question:
			got.questions, got.order = append(got.questions, e.Question), append(got.order, "question")
		case *mtgv1.ChatResponse_Slots:
			got.slots, got.order = e.Slots, append(got.order, "slots")
		case *mtgv1.ChatResponse_Status:
			got.statuses, got.order = append(got.statuses, e.Status), append(got.order, "status")
		case *mtgv1.ChatResponse_Usage:
			got.usage, got.order = e.Usage, append(got.order, "usage")
		case *mtgv1.ChatResponse_Failure:
			got.failure, got.order = e.Failure, append(got.order, "failure")
		case *mtgv1.ChatResponse_Deck:
			got.deck, got.order = e.Deck, append(got.order, "deck")
		case *mtgv1.ChatResponse_TextDelta:
			got.texts, got.order = append(got.texts, e.TextDelta), append(got.order, "text")
		}
	}
	if err := stream.Err(); err != nil {
		t.Fatalf("stream: %v", err)
	}
	return got
}

func TestChatStartsASessionAndAsks(t *testing.T) {
	store := newFakeStore()
	client, _ := testServer(t, store, firstTurn(t)...)
	got := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})

	if got.started == "" || got.order[0] != "started" {
		t.Fatalf("the first event is not the session id: %v", got.order)
	}
	if len(got.questions) == 0 {
		t.Fatal("no question reached the user")
	}
	if got.slots.GetTheme() != "lifegain" {
		t.Errorf("theme = %q", got.slots.GetTheme())
	}
	if got.usage.GetCalls() != 3 {
		t.Errorf("usage calls = %d, want 3", got.usage.GetCalls())
	}
	if got.usage.GetInputTokens() != 950 || got.usage.GetOutputTokens() != 130 {
		t.Errorf("usage tokens = %+v", got.usage)
	}
	stored := store.sessions[got.started]
	if stored == nil {
		t.Fatal("the session was not stored")
	}
	if len(stored.GetTurns()) != 1 || stored.GetTurns()[0].GetUserMessage() == "" {
		t.Errorf("stored turns = %+v", stored.GetTurns())
	}
	if stored.GetStatus() != mtgv1.SessionStatus_SESSION_STATUS_ASKING {
		t.Errorf("status = %v", stored.GetStatus())
	}
	snap := store.states[got.started]
	if len(snap.Ctx.Asked) != len(got.questions) {
		t.Errorf("the private state holds %d asked rows, the turn asked %d", len(snap.Ctx.Asked), len(got.questions))
	}
}

// TestChatResumesWithoutRepeating is the point of the private state
// document (D-74): the second turn knows what the first one asked.
func TestChatResumesWithoutRepeating(t *testing.T) {
	store := newFakeStore()
	// A typed slot closes on its value, not by name (D-83).
	steps := append(firstTurn(t),
		classifyJSON(t, map[string]any{
			"power":           "bracket 3",
			"commander_names": []string{"Karlov of the Ghost Council"},
		}),
		scoreJSON(t),
		askJSON(t))
	client, _ := testServer(t, store, steps...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	second := chat(t, client, &mtgv1.ChatRequest{
		SessionId: first.started,
		Message:   "Karlov, bracket 3",
	})
	if second.started != "" {
		t.Errorf("a resumed session started a new id %q", second.started)
	}
	asked := map[string]bool{}
	for _, q := range first.questions {
		asked[q.GetSlot()+q.GetText()] = true
	}
	for _, q := range second.questions {
		if asked[q.GetSlot()+q.GetText()] {
			t.Errorf("turn 2 repeated %q", q.GetText())
		}
	}
	stored := store.sessions[first.started]
	if len(stored.GetTurns()) != 2 {
		t.Errorf("stored turns = %d, want 2", len(stored.GetTurns()))
	}
	if stored.GetStatus() != mtgv1.SessionStatus_SESSION_STATUS_READY {
		t.Errorf("status = %v, want READY once every slot is closed", stored.GetStatus())
	}
	// Turn 1 costs three calls. Turn 2 costs one: the classify call closes
	// the last two keys, the plan is then empty, and the agent scores and
	// phrases nothing.
	if stored.GetUsage().GetCalls() != 4 {
		t.Errorf("session usage calls = %d, want 4 over the two turns", stored.GetUsage().GetCalls())
	}
}

// TestAnswersReachTheClassifier proves an option index becomes words the
// classifier can map, which is what the UI sends.
func TestAnswersReachTheClassifier(t *testing.T) {
	store := newFakeStore()
	steps := append(firstTurn(t),
		classifyJSON(t, map[string]any{"power": "bracket 3"}),
		scoreJSON(t),
		askJSON(t))
	client, sc := testServer(t, store, steps...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	var bracket *mtgv1.Question
	for _, q := range first.questions {
		if q.GetSlot() == "power" {
			bracket = q
		}
	}
	if bracket == nil || len(bracket.GetOptions()) < 3 {
		t.Fatalf("no bracket question with options: %+v", first.questions)
	}
	chat(t, client, &mtgv1.ChatRequest{
		SessionId: first.started,
		Answers:   []*mtgv1.Answer{{QuestionId: bracket.GetId(), OptionIndex: proto.Int32(2)}},
	})
	last := sc.Calls[3].Input
	if !strings.Contains(last, bracket.GetOptions()[2]) {
		t.Errorf("the chosen option did not reach the classifier: %s", last)
	}
	if !strings.Contains(last, "bracket") && !strings.Contains(last, "Bracket") {
		t.Errorf("the question text did not reach the classifier: %s", last)
	}
}

// TestModelFailureEndsTheTurn keeps a provider outage off the wire as a
// transport error. The UI reads a failure event instead.
func TestModelFailureEndsTheTurn(t *testing.T) {
	store := newFakeStore()
	client, _ := testServer(t, store, llm.Step{Err: errors.New("provider down")})
	got := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain deck"})
	if got.failure == nil {
		t.Fatal("no failure event")
	}
	if !strings.HasPrefix(got.failure.GetCode(), "llm_") {
		t.Errorf("failure code = %q", got.failure.GetCode())
	}
	if len(got.questions) != 0 {
		t.Errorf("a failed turn asked %d questions", len(got.questions))
	}
	if len(store.sessions) != 1 {
		t.Error("a failed turn stored no session, so the user loses the id")
	}
}

func TestGetSession(t *testing.T) {
	store := newFakeStore()
	client, _ := testServer(t, store, firstTurn(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})

	res, err := client.GetSession(context.Background(), connect.NewRequest(&mtgv1.GetSessionRequest{SessionId: first.started}))
	if err != nil {
		t.Fatalf("get session: %v", err)
	}
	if res.Msg.GetSession().GetId() != first.started {
		t.Errorf("session id = %q", res.Msg.GetSession().GetId())
	}
	if res.Msg.GetSession().GetSlots().GetTheme() != "lifegain" {
		t.Errorf("slots did not survive: %+v", res.Msg.GetSession().GetSlots())
	}

	_, err = client.GetSession(context.Background(), connect.NewRequest(&mtgv1.GetSessionRequest{SessionId: "nope"}))
	if connect.CodeOf(err) != connect.CodeNotFound {
		t.Errorf("unknown session gave %v, want NotFound", connect.CodeOf(err))
	}
}

func TestEmptyRequestIsRefused(t *testing.T) {
	client, _ := testServer(t, newFakeStore())
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{}))
	if err == nil {
		for stream.Receive() {
		}
		err = stream.Err()
		_ = stream.Close()
	}
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("an empty request gave %v, want InvalidArgument", connect.CodeOf(err))
	}
}

// TestAfterBuildAsksNothingMore is D-256: the variance row is retired,
// so a session that holds a deck and has every slot settled asks nothing
// at all.
func TestAfterBuildAsksNothingMore(t *testing.T) {
	store := newFakeStore()
	client, _ := testServer(t, store, firstTurn(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	for _, q := range first.questions {
		if q.GetSlot() == "plan_variant" {
			t.Errorf("the retired variance row fired: %q", q.GetText())
		}
	}
	// A stored deck must not raise it either.
	store.sessions[first.started].DeckIds = []string{"deck-1"}
	steps := []llm.Step{
		classifyJSON(t, map[string]any{
			"power":           "bracket 3",
			"commander_names": []string{"Karlov of the Ghost Council"},
		}),
		scoreJSON(t),
		askJSON(t),
	}
	client2, _ := testServer(t, store, steps...)
	second := chat(t, client2, &mtgv1.ChatRequest{SessionId: first.started, Message: "give me another version"})
	if len(second.questions) != 0 {
		t.Errorf("a settled session with a deck asked %d questions: %v", len(second.questions), second.questions)
	}
}

// TestChatStaleVersionIsAborted: a turn
// reads the session, another write lands, and the turn's Put must fail
// with CodeAborted and store nothing.
func TestChatStaleVersionIsAborted(t *testing.T) {
	store := newFakeStore()
	steps := append(firstTurn(t),
		classifyJSON(t, map[string]any{"power": "bracket 3"}),
		scoreJSON(t),
		askJSON(t))
	client, _ := testServer(t, store, steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	before := len(store.sessions[first.started].GetTurns())

	// The overlapping turn wins the write between this turn's read and
	// its Put.
	store.onGetState = func() {
		store.onGetState = nil
		store.versions[first.started]++
	}
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{
		SessionId: first.started, Message: "bracket 3",
	}))
	if err == nil {
		for stream.Receive() {
		}
		err = stream.Err()
		_ = stream.Close()
	}
	if connect.CodeOf(err) != connect.CodeAborted {
		t.Fatalf("a stale put gave %v, want Aborted: %v", connect.CodeOf(err), err)
	}
	if !strings.Contains(err.Error(), "busy") {
		t.Errorf("the message does not tell the client to retry: %v", err)
	}
	if got := len(store.sessions[first.started].GetTurns()); got != before {
		t.Errorf("the loser's turn was stored: %d turns, want %d", got, before)
	}
}

// TestChatSecondTurnAdvancesTheVersion proves a normal resume passes the
// version it read, so the store accepts it.
func TestChatSecondTurnAdvancesTheVersion(t *testing.T) {
	store := newFakeStore()
	steps := append(firstTurn(t),
		classifyJSON(t, map[string]any{"power": "bracket 3"}),
		scoreJSON(t),
		askJSON(t))
	client, _ := testServer(t, store, steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	if store.versions[first.started] != 1 {
		t.Fatalf("version after turn 1 = %d, want 1", store.versions[first.started])
	}
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "bracket 3"})
	if store.versions[first.started] != 2 {
		t.Errorf("version after turn 2 = %d, want 2", store.versions[first.started])
	}
}

// TestAddUsageKeepsPricedOnAnEmptyTurn: a frozen turn makes no
// call and reports no cost. That is not an unpriced call.
func TestAddUsageKeepsPricedOnAnEmptyTurn(t *testing.T) {
	cost := 0.01
	total := addUsage(nil, llm.Report{Calls: 3, CostUSD: &cost})
	if !total.GetPriced() || total.GetCostUsd() != cost {
		t.Fatalf("a priced turn gave %+v", total)
	}
	total = addUsage(total, llm.Report{})
	if !total.GetPriced() {
		t.Error("a turn with no calls flipped priced to false")
	}
	if total.GetCalls() != 3 {
		t.Errorf("calls = %d, want 3", total.GetCalls())
	}
	// A turn with calls and no cost is unpriced, as before.
	total = addUsage(total, llm.Report{Calls: 1})
	if total.GetPriced() {
		t.Error("a turn with calls and no cost left priced true")
	}
}

// TestAnswerTextWinsOverOptionZero pins the Answer contract. The proto
// says an unset option_index means free text, so 0 is option 0. A client
// that sends text and option 0 gets its text, not the first option.
func TestAnswerTextWinsOverOptionZero(t *testing.T) {
	session := &mtgv1.Session{Turns: []*mtgv1.Turn{{Questions: []*mtgv1.Question{{
		Id: "q-power", Text: "Which bracket?", Options: []string{"Bracket 1", "Bracket 2"},
	}}}}}
	got := withAnswers("", []*mtgv1.Answer{{QuestionId: "q-power", OptionIndex: proto.Int32(0), Text: "somewhere near 3"}}, session)
	if got != "Q: Which bracket?\nA: somewhere near 3" {
		t.Errorf("text with option 0 gave %q", got)
	}
	got = withAnswers("", []*mtgv1.Answer{{QuestionId: "q-power", OptionIndex: proto.Int32(0)}}, session)
	if got != "Q: Which bracket?\nA: Bracket 1" {
		t.Errorf("option 0 with no text gave %q", got)
	}
	got = withAnswers("", []*mtgv1.Answer{{QuestionId: "q-power"}}, session)
	if got != "" {
		t.Errorf("free text with no text gave %q, want nothing", got)
	}
}

// TestBuildStoresThePostTurnState: the second write of a building turn
// must carry the turn's asked rows and the built status, or the next
// turn repeats a question and the session never says BUILT.
func TestBuildStoresThePostTurnState(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	deck := &mtgv1.Deck{Summary: "a deck", Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}}
	opts := append(buildOpts(t, fd), WithDeckStore(ds))
	client, _ := testServerOpts(t, store, opts, readySteps(t)...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	askedBefore := len(store.states[first.started].Ctx.Asked)
	if askedBefore == 0 {
		t.Fatal("turn 1 asked nothing, the test proves nothing")
	}
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil {
		t.Fatalf("no deck streamed: %v", second.order)
	}
	snap := store.states[first.started]
	if len(snap.Ctx.Asked) < askedBefore {
		t.Errorf("the stored state lost the asked rows: %d, had %d", len(snap.Ctx.Asked), askedBefore)
	}
	if len(snap.CommanderNames) == 0 {
		t.Errorf("the stored state lost the commander: %+v", snap)
	}
	if got := store.sessions[first.started].GetStatus(); got != mtgv1.SessionStatus_SESSION_STATUS_BUILT {
		t.Errorf("status = %v, want BUILT", got)
	}
	if store.versions[first.started] != 3 {
		t.Errorf("version = %d, want 3: the turn and the deck id", store.versions[first.started])
	}
}

// TestBuildFailureKeepsReady: a build that kept no deck leaves the
// session READY.
func TestBuildFailureKeepsReady(t *testing.T) {
	store := newFakeStore()
	fd := &fakeDecks{err: errors.New("the model is down")}
	client, _ := testServerOpts(t, store, buildOpts(t, fd), readySteps(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if got := store.sessions[first.started].GetStatus(); got != mtgv1.SessionStatus_SESSION_STATUS_READY {
		t.Errorf("status = %v, want READY", got)
	}
}

// TestMessageCap: a message or an answer past MaxMessageBytes, too many
// answers, or answers that together pass the cap are refused before any
// model call.
func TestMessageCap(t *testing.T) {
	long := strings.Repeat("x", MaxMessageBytes+1)
	tests := []struct {
		name string
		req  *mtgv1.ChatRequest
		want connect.Code
	}{
		{name: "long message", req: &mtgv1.ChatRequest{Message: long}, want: connect.CodeInvalidArgument},
		{name: "long answer", req: &mtgv1.ChatRequest{Answers: []*mtgv1.Answer{{QuestionId: "q", Text: long}}}, want: connect.CodeInvalidArgument},
		{name: "at the cap", req: &mtgv1.ChatRequest{Message: strings.Repeat("x", MaxMessageBytes)}, want: 0},
		{name: "too many answers", req: &mtgv1.ChatRequest{Answers: manyAnswers(MaxAnswers + 1)}, want: connect.CodeInvalidArgument},
		{name: "answers at the count cap", req: &mtgv1.ChatRequest{Answers: manyAnswers(MaxAnswers)}, want: 0},
		{name: "answers long together", req: &mtgv1.ChatRequest{Answers: []*mtgv1.Answer{
			{QuestionId: "a", Text: strings.Repeat("x", MaxMessageBytes/2+1)},
			{QuestionId: "b", Text: strings.Repeat("x", MaxMessageBytes/2+1)},
		}}, want: connect.CodeInvalidArgument},
		{name: "path session id", req: &mtgv1.ChatRequest{Message: "hi", SessionId: "../x"}, want: connect.CodeInvalidArgument},
		{name: "path collection id", req: &mtgv1.ChatRequest{Message: "hi", CollectionId: "a/b"}, want: connect.CodeInvalidArgument},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, sc := testServer(t, newFakeStore(), firstTurn(t)...)
			stream, err := client.Chat(context.Background(), connect.NewRequest(tt.req))
			if err == nil {
				for stream.Receive() {
				}
				err = stream.Err()
				_ = stream.Close()
			}
			var got connect.Code
			if err != nil {
				got = connect.CodeOf(err)
			}
			if got != tt.want {
				t.Errorf("code = %v, want %v: %v", got, tt.want, err)
			}
			if tt.want != 0 && len(sc.Calls) != 0 {
				t.Errorf("a refused turn made %d model calls", len(sc.Calls))
			}
		})
	}
}

// gate is a provider that waits until the test releases it.
type gate struct {
	entered chan struct{}
	release chan struct{}
	steps   []llm.Step
	mu      sync.Mutex
}

func (g *gate) Name() string { return llm.FakeName }

func (g *gate) Complete(ctx context.Context, call llm.Call) (llm.Response, error) {
	g.mu.Lock()
	first := len(g.steps) > 0
	var st llm.Step
	if first {
		st, g.steps = g.steps[0], g.steps[1:]
	}
	g.mu.Unlock()
	select {
	case g.entered <- struct{}{}:
	default:
	}
	select {
	case <-g.release:
	case <-ctx.Done():
		return llm.Response{}, ctx.Err()
	}
	if !first {
		return llm.Response{}, errors.New("script exhausted")
	}
	return llm.Response{Output: st.Output, Model: call.Model, Usage: st.Usage}, nil
}

// TestConcurrencyCap: the turn past the limit answers
// ResourceExhausted at once, and the running turn is not disturbed.
func TestConcurrencyCap(t *testing.T) {
	cat, err := questions.Load()
	if err != nil {
		t.Fatal(err)
	}
	roles := map[llm.Role]llm.RoleSpec{}
	for _, r := range llm.Roles {
		roles[r] = llm.RoleSpec{Provider: llm.FakeName, Model: "fake-" + string(r), MaxOutputTokens: 1024}
	}
	g := &gate{entered: make(chan struct{}, 1), release: make(chan struct{}), steps: firstTurn(t)}
	client, err := llm.New(&llm.Config{VerifiedAt: "2026-08-24", Roles: roles}, []llm.Provider{g}, llm.WithoutJitter())
	if err != nil {
		t.Fatal(err)
	}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	srv, err := New(cat, client, newFakeStore(), func(context.Context) string { return "u1" },
		WithLogger(quiet), WithChatLimit(1))
	if err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewAgentServiceHandler(srv))
	httpSrv := httptest.NewServer(mux)
	t.Cleanup(httpSrv.Close)
	c := mtgv1connect.NewAgentServiceClient(httpSrv.Client(), httpSrv.URL)

	firstDone := make(chan error, 1)
	go func() {
		stream, err := c.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"}))
		if err != nil {
			firstDone <- err
			return
		}
		for stream.Receive() {
		}
		firstDone <- stream.Err()
		_ = stream.Close()
	}()
	select {
	case <-g.entered:
	case <-time.After(5 * time.Second):
		t.Fatal("the first turn never reached the provider")
	}

	stream, err := c.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: "another"}))
	if err == nil {
		for stream.Receive() {
		}
		err = stream.Err()
		_ = stream.Close()
	}
	if connect.CodeOf(err) != connect.CodeResourceExhausted {
		t.Errorf("the second turn gave %v, want ResourceExhausted: %v", connect.CodeOf(err), err)
	}
	if err != nil && !strings.Contains(err.Error(), "again") {
		t.Errorf("the message does not tell the user to retry: %v", err)
	}
	close(g.release)
	if err := <-firstDone; err != nil {
		t.Errorf("the first turn failed: %v", err)
	}
	// The token is back, so a third turn is admitted.
	stream, err = c.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: "third"}))
	if err == nil {
		for stream.Receive() {
		}
		err = stream.Err()
		_ = stream.Close()
	}
	if connect.CodeOf(err) == connect.CodeResourceExhausted {
		t.Error("the gate did not release after the first turn")
	}
}

// TestFailureMessages covers the failure codes the UI acts on.
func TestFailureMessages(t *testing.T) {
	mk := func(class llm.Class, status int, msg string) error {
		return &llm.Error{Class: class, Provider: "p", Model: "m", Status: status, Err: errors.New(msg)}
	}
	tests := []struct {
		name      string
		err       error
		wantCode  string
		wantRetry bool
		wantText  string
	}{
		{name: "schema says retry", err: mk(llm.ClassSchema, 0, "bad shape"), wantCode: "llm_schema", wantRetry: true, wantText: "retry usually succeeds"},
		{name: "auth is the operator's", err: mk(llm.ClassTerminal, 401, "bad key"), wantCode: "llm_terminal", wantText: "operator must fix"},
		{name: "forbidden is the operator's", err: mk(llm.ClassTerminal, 403, "no"), wantCode: "llm_terminal", wantText: "operator must fix"},
		{name: "no provider is the operator's", err: mk(llm.ClassTerminal, 0, `no provider "x" wired`), wantCode: "llm_terminal", wantText: "operator must fix"},
		{name: "bad request is the user's", err: mk(llm.ClassTerminal, 400, "too long"), wantCode: "llm_terminal", wantText: "does not succeed on a retry"},
		{name: "transient", err: mk(llm.ClassTransient, 503, "down"), wantCode: "llm_transient", wantRetry: true, wantText: "Send the message again"},
		{name: "refusal", err: mk(llm.ClassRefusal, 0, "no"), wantCode: "llm_refusal", wantText: "other words"},
		{name: "plain error is terminal", err: errors.New("x"), wantCode: "llm_terminal", wantText: "does not succeed on a retry"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := failure(tt.err)
			if got.GetCode() != tt.wantCode || got.GetRetryable() != tt.wantRetry {
				t.Errorf("code %q retryable %v, want %q %v", got.GetCode(), got.GetRetryable(), tt.wantCode, tt.wantRetry)
			}
			if !strings.Contains(got.GetMessage(), tt.wantText) {
				t.Errorf("message %q does not say %q", got.GetMessage(), tt.wantText)
			}
		})
	}
}

// TestStoreErrorCodes: a session past the document limit is
// refused with a message, not bricked as Internal.
func TestStoreErrorCodes(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want connect.Code
		text string
	}{
		{name: "too large", err: fmt.Errorf("store: %w", sessions.ErrTooLarge), want: connect.CodeResourceExhausted, text: "start a new session"},
		{name: "not found", err: sessions.ErrNotFound, want: connect.CodeNotFound},
		{name: "conflict", err: sessions.ErrConflict, want: connect.CodeAborted, text: "busy"},
		{name: "other", err: errors.New("x"), want: connect.CodeInternal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := storeError(tt.err)
			if connect.CodeOf(got) != tt.want {
				t.Errorf("code = %v, want %v", connect.CodeOf(got), tt.want)
			}
			if !strings.Contains(got.Error(), tt.text) {
				t.Errorf("message %q does not say %q", got.Error(), tt.text)
			}
		})
	}
}

type fakePrecons struct{ set *precons.Set }

func (f *fakePrecons) Current() *precons.Set { return f.set }

// TestPreconSourceIsLateBound: the set the source holds now is the set
// the build reads.
func TestPreconSourceIsLateBound(t *testing.T) {
	cat, err := questions.Load()
	if err != nil {
		t.Fatal(err)
	}
	client, _ := fakeClient(t)
	src := &fakePrecons{}
	srv, err := New(cat, client, newFakeStore(), func(context.Context) string { return "u1" }, WithPreconSource(src))
	if err != nil {
		t.Fatal(err)
	}
	if srv.preconSet() != nil {
		t.Error("a source with no set gave a set")
	}
	src.set = &precons.Set{}
	if srv.preconSet() != src.set {
		t.Error("the late set did not reach the service")
	}
	// Without a source there is no set.
	srv, err = New(cat, client, newFakeStore(), func(context.Context) string { return "u1" })
	if err != nil {
		t.Fatal(err)
	}
	if srv.preconSet() != nil {
		t.Error("a server with no source gave a set")
	}
}

// TestCardOptionsCarryOracleIds is D-287: an option that names a card
// carries its id, a non-card option an empty string, and a question with
// no card option keeps the field empty.
func TestCardOptionsCarryOracleIds(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "o-karlov", Name: "Karlov of the Ghost Council"},
		{OracleId: "o-giada", Name: "Giada, Font of Hope"},
	}, nil, nil, time.Unix(1000, 0).UTC())
	offer := &mtgv1.Question{Options: []string{"Karlov of the Ghost Council", "giada, font of hope", "None, name three more"}}
	format := &mtgv1.Question{Options: []string{"Commander", "Modern"}}
	cardOptions([]*mtgv1.Question{offer, format}, idx)
	if got := offer.GetOptionOracleIds(); len(got) != 3 || got[0] != "o-karlov" || got[1] != "o-giada" || got[2] != "" {
		t.Errorf("offer ids = %v", got)
	}
	if len(format.GetOptionOracleIds()) != 0 {
		t.Errorf("format ids = %v", format.GetOptionOracleIds())
	}
	cardOptions([]*mtgv1.Question{offer}, nil)
}

// manyAnswers makes n short answers.
func manyAnswers(n int) []*mtgv1.Answer {
	out := make([]*mtgv1.Answer, n)
	for i := range out {
		out[i] = &mtgv1.Answer{QuestionId: fmt.Sprintf("q%d", i), Text: "yes"}
	}
	return out
}

// TestGetSessionRefusesAPathID: an id with a slash addresses another
// document path, so it is refused before the store sees it.
func TestGetSessionRefusesAPathID(t *testing.T) {
	client, _ := testServer(t, newFakeStore())
	for _, id := range []string{"a/b", "..", ".", strings.Repeat("x", 1501)} {
		_, err := client.GetSession(context.Background(), connect.NewRequest(&mtgv1.GetSessionRequest{SessionId: id}))
		if connect.CodeOf(err) != connect.CodeInvalidArgument {
			t.Errorf("id %q gave %v, want InvalidArgument", id, connect.CodeOf(err))
		}
	}
}

// TestSessionStartedFollowsTheRequest: session_started is sent when the
// request names no session, and not when it resumes one.
func TestSessionStartedFollowsTheRequest(t *testing.T) {
	store := newFakeStore()
	client, _ := testServer(t, store, append(firstTurn(t), classifyJSON(t, nil), scoreJSON(t), askJSON(t))...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	if first.started == "" {
		t.Fatal("the first turn sent no session_started")
	}
	// A stored session with no created_at still resumes without a
	// second session_started.
	store.sessions[first.started].CreatedAt = nil
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "bracket 3"})
	if second.started != "" {
		t.Errorf("a resumed session sent session_started %q", second.started)
	}
}

// TestRevisionQuestionSetsAsking: a session that asked a revision
// question is ASKING again, not BUILT.
func TestRevisionQuestionSetsAsking(t *testing.T) {
	store := newFakeStore()
	ds := &fakeDeckStore{}
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}, Cards: []*mtgv1.DeckCard{{OracleId: "o-plains", Name: "Plains", Count: 30}}}}}
	steps := append(builtSteps(t),
		classifyJSON(t, nil),
		reviseJSON(t, map[string]any{"question": "Which lands?"}))
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Better lands please"})
	if got := store.status(first.started); got != mtgv1.SessionStatus_SESSION_STATUS_ASKING {
		t.Errorf("status = %v, want ASKING", got)
	}
}

func TestSlotsChangedGuardsNil(t *testing.T) {
	if slotsChanged(nil, nil) {
		t.Error("two nil slot sets differ")
	}
	if !slotsChanged(nil, &mtgv1.Slots{Theme: "x"}) {
		t.Error("nil and a theme do not differ")
	}
	if slotsChanged(nil, &mtgv1.Slots{SlotStates: map[string]mtgv1.SlotState{"theme": mtgv1.SlotState_SLOT_STATE_ASKED}}) {
		t.Error("a fill state counted as a setting")
	}
}

func TestOperatorFaultGuardsNilErr(t *testing.T) {
	if operatorFault(&llm.Error{Class: llm.ClassTerminal, Status: 400}) {
		t.Error("a 400 with no inner error was the operator's")
	}
	if !operatorFault(&llm.Error{Class: llm.ClassTerminal, Status: 401}) {
		t.Error("a 401 with no inner error was not the operator's")
	}
}

// A collection the user deleted must not end a turn (D-347). The chat
// forgets it, builds from the whole card database, and says so once.
func TestChatWithADeletedCollection(t *testing.T) {
	store := newFakeStore()
	store.sessions["s1"] = &mtgv1.Session{
		Id:           "s1",
		CollectionId: "gone",
		Slots:        &mtgv1.Slots{PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY},
		Status:       mtgv1.SessionStatus_SESSION_STATUS_ASKING,
	}
	store.states["s1"] = questions.Snapshot{Version: questions.SnapshotVersion}
	missing := fakeCollections{countsErr: status.Error(codes.NotFound, "no such collection")}
	client, _ := testServerOpts(t, store, []Option{WithCollections(missing)}, firstTurn(t)...)

	got := chat(t, client, &mtgv1.ChatRequest{SessionId: "s1", Message: "swap the removal"})

	var warned bool
	for _, s := range got.statuses {
		if strings.Contains(s, "whole card database") {
			warned = true
		}
	}
	if !warned {
		t.Errorf("the chat said nothing about the deleted collection: %q", got.statuses)
	}
	stored := store.sessions["s1"]
	if stored.GetCollectionId() != "" {
		t.Errorf("the session still names the collection %q", stored.GetCollectionId())
	}
	if stored.GetSlots().GetPoolRule() != mtgv1.PoolRule_POOL_RULE_ANY_CARD {
		t.Errorf("pool rule = %v, want ANY_CARD", stored.GetSlots().GetPoolRule())
	}
}

// TestChatRefusesATurnWithNoIndex is D-405. A turn with no card index
// resolves no name, offers no commander, and builds nothing, so it
// waits with Unavailable as ImportCollection does. Before this the turn
// ran, and the status line claimed a commander it never had.
func TestChatRefusesATurnWithNoIndex(t *testing.T) {
	client, _ := testServerOpts(t, newFakeStore(), []Option{WithCandidates(fixedIndex{}, nil)})
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: "Build the best possible deck you can"}))
	if err != nil {
		t.Fatalf("chat: %v", err)
	}
	defer func() { _ = stream.Close() }()
	for stream.Receive() {
		t.Errorf("a turn with no index sent an event: %v", stream.Msg().GetEvent())
	}
	if got := connect.CodeOf(stream.Err()); got != connect.CodeUnavailable {
		t.Errorf("code = %v, want Unavailable", got)
	}
}

// List, Rename, and Delete serve the sessions list of D-433.
func (f *fakeStore) List(_ context.Context, _ string) ([]*mtgv1.SessionSummary, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([]*mtgv1.SessionSummary, 0, len(f.sessions))
	for _, s := range f.sessions {
		out = append(out, sessions.Summarize(s))
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i].GetUpdatedAt().AsTime(), out[j].GetUpdatedAt().AsTime()
		if !a.Equal(b) {
			return a.After(b)
		}
		return out[i].GetId() > out[j].GetId()
	})
	return out, nil
}

func (f *fakeStore) Rename(_ context.Context, _, id, name string) (*mtgv1.SessionSummary, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	s, ok := f.sessions[id]
	if !ok {
		return nil, sessions.ErrNotFound
	}
	s.Name = name
	f.versions[id]++
	return sessions.Summarize(s), nil
}

func (f *fakeStore) Delete(_ context.Context, _, id string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.sessions[id]; !ok {
		return sessions.ErrNotFound
	}
	delete(f.sessions, id)
	delete(f.states, id)
	delete(f.versions, id)
	return nil
}
