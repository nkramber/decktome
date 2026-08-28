package agentsvc

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
	"github.com/nkramber/mtg-deck-builder/go/internal/sessions"
)

// fakeStore is an in-memory Store. It keeps the two halves apart, the
// same way Firestore does, so a test can prove the private state moves.
type fakeStore struct {
	sessions map[string]*mtgv1.Session
	states   map[string]questions.Snapshot
	versions map[string]int64
	ids      int
	putErr   error
	// onGetState runs after each GetState, so a test can slip a write in
	// between the read and the Put, the way an overlapping turn does.
	onGetState func()
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

func (f *fakeStore) Put(_ context.Context, _ string, s *mtgv1.Session, snap questions.Snapshot, expected int64) error {
	if f.putErr != nil {
		return f.putErr
	}
	if f.versions[s.GetId()] != expected {
		return sessions.ErrConflict
	}
	f.sessions[s.GetId()] = s
	f.states[s.GetId()] = snap
	f.versions[s.GetId()] = expected + 1
	return nil
}

func (f *fakeStore) Get(_ context.Context, _, id string) (*mtgv1.Session, error) {
	s, ok := f.sessions[id]
	if !ok {
		return nil, sessions.ErrNotFound
	}
	// A copy, as Firestore decodes one. A turn that edits the session
	// in place must not reach the store without a Put.
	return proto.Clone(s).(*mtgv1.Session), nil
}

func (f *fakeStore) GetState(ctx context.Context, uid, id string) (*mtgv1.Session, questions.Snapshot, int64, error) {
	s, err := f.Get(ctx, uid, id)
	if err != nil {
		return nil, questions.Snapshot{}, 0, err
	}
	snap, version := f.states[id], f.versions[id]
	if f.onGetState != nil {
		f.onGetState()
	}
	return s, snap, version, nil
}

// classifyJSON writes one classify answer that satisfies the schema.
func classifyJSON(t *testing.T, fields map[string]any) llm.Step {
	t.Helper()
	out := map[string]any{
		"format": "unknown", "theme": "", "colors": []string{},
		"commander_names": []string{}, "locked_names": []string{}, "named_cards": []string{},
		"power": "", "pool_rule": "unknown", "budget_usd": 0.0, "budget_scope": "unknown",
		"closed_keys": []string{}, "declined_keys": []string{},
		"facts": map[string]bool{
			"named_card": false, "buy_list": false,
			"house_format": false, "two_plans": false, "budget_ambiguous": false,
			"power_competitive": false, "wants_suggestion": false, "out_of_scope": false,
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
			// this test measures the commander rows (D-168).
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

// testServerOpts builds the server with extra options, for the wiring
// that PR-8 adds.
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
	got := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})

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
	// A typed slot closes on its value, never by name (D-83).
	steps := append(firstTurn(t),
		classifyJSON(t, map[string]any{
			"power":           "bracket 3",
			"commander_names": []string{"Karlov of the Ghost Council"},
		}),
		scoreJSON(t),
		askJSON(t))
	client, _ := testServer(t, store, steps...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
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
// classifier can map, which is what the UI sends (PR-12).
func TestAnswersReachTheClassifier(t *testing.T) {
	store := newFakeStore()
	steps := append(firstTurn(t),
		classifyJSON(t, map[string]any{"power": "bracket 3"}),
		scoreJSON(t),
		askJSON(t))
	client, sc := testServer(t, store, steps...)

	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
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
		Answers:   []*mtgv1.Answer{{QuestionId: bracket.GetId(), OptionIndex: 2}},
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
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})

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

// TestAfterBuildAsksNothingMore is what the variance row used to test.
// The row is retired with PR-9 (D-256), so a session that holds a deck
// and has every slot settled asks nothing at all: it builds again.
func TestAfterBuildAsksNothingMore(t *testing.T) {
	store := newFakeStore()
	client, _ := testServer(t, store, firstTurn(t)...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
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
	for _, q := range second.questions {
		if q.GetSlot() == "plan_variant" {
			t.Errorf("the retired variance row fired on a stored deck: %q", q.GetText())
		}
	}
}

// TestChatStaleVersionIsAborted is the H-7 guard at the service. A turn
// reads the session, another write lands, and the turn's Put must fail
// with CodeAborted and store nothing.
func TestChatStaleVersionIsAborted(t *testing.T) {
	store := newFakeStore()
	steps := append(firstTurn(t),
		classifyJSON(t, map[string]any{"power": "bracket 3"}),
		scoreJSON(t),
		askJSON(t))
	client, _ := testServer(t, store, steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
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
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck"})
	if store.versions[first.started] != 1 {
		t.Fatalf("version after turn 1 = %d, want 1", store.versions[first.started])
	}
	chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "bracket 3"})
	if store.versions[first.started] != 2 {
		t.Errorf("version after turn 2 = %d, want 2", store.versions[first.started])
	}
}

// TestAddUsageKeepsPricedOnAnEmptyTurn is M-8. A frozen turn makes no
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
// says -1 means free text, so 0 is option 0. A client that sends text and
// leaves option_index at 0 gets its text, not the first option.
func TestAnswerTextWinsOverOptionZero(t *testing.T) {
	session := &mtgv1.Session{Turns: []*mtgv1.Turn{{Questions: []*mtgv1.Question{{
		Id: "q-power", Text: "Which bracket?", Options: []string{"Bracket 1", "Bracket 2"},
	}}}}}
	got := withAnswers("", []*mtgv1.Answer{{QuestionId: "q-power", OptionIndex: 0, Text: "somewhere near 3"}}, session)
	if got != "Which bracket? somewhere near 3" {
		t.Errorf("text with option 0 gave %q", got)
	}
	got = withAnswers("", []*mtgv1.Answer{{QuestionId: "q-power", OptionIndex: 0}}, session)
	if got != "Which bracket? Bracket 1" {
		t.Errorf("option 0 with no text gave %q", got)
	}
	got = withAnswers("", []*mtgv1.Answer{{QuestionId: "q-power", OptionIndex: -1}}, session)
	if got != "" {
		t.Errorf("free text with no text gave %q, want nothing", got)
	}
}
