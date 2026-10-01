package agentsvc

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/questions"
)

// gatedProvider holds its first call open until the test opens the gate,
// the way a slow classify call holds the first turn (F-186).
type gatedProvider struct {
	*llm.Script
	once    sync.Once
	entered chan struct{}
	gate    chan struct{}
}

func (g *gatedProvider) Complete(ctx context.Context, call llm.Call) (llm.Response, error) {
	g.once.Do(func() {
		close(g.entered)
		<-g.gate
	})
	return g.Script.Complete(ctx, call)
}

// TestReloadDuringFirstTurnFindsTheSession is F-186. The web moves to the
// session address on session_started, so a reload reads GetSession while
// the first turn still runs. That read must find the session and report
// the turn, and the end of the turn must store its questions.
func TestReloadDuringFirstTurnFindsTheSession(t *testing.T) {
	cat, err := questions.Load()
	if err != nil {
		t.Fatalf("catalog: %v", err)
	}
	roles := map[llm.Role]llm.RoleSpec{}
	for _, r := range llm.Roles {
		roles[r] = llm.RoleSpec{Provider: llm.FakeName, Model: "fake-" + string(r), MaxOutputTokens: 1024}
	}
	p := &gatedProvider{Script: llm.NewScript(firstTurn(t)...), entered: make(chan struct{}), gate: make(chan struct{})}
	lc, err := llm.New(&llm.Config{VerifiedAt: "2026-08-24", Roles: roles}, []llm.Provider{p}, llm.WithoutJitter())
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	store := newFakeStore()
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	srv, err := New(cat, lc, store, func(context.Context) string { return "u1" },
		WithLogger(quiet), WithClock(func() time.Time { return time.Unix(1000, 0).UTC() }))
	if err != nil {
		t.Fatalf("server: %v", err)
	}
	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewAgentServiceHandler(srv))
	httpSrv := httptest.NewServer(mux)
	t.Cleanup(httpSrv.Close)
	client := mtgv1connect.NewAgentServiceClient(httpSrv.Client(), httpSrv.URL)

	const message = "build me a lifegain commander deck for 50 dollars"
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: message}))
	if err != nil {
		t.Fatalf("chat: %v", err)
	}
	defer func() { _ = stream.Close() }()
	if !stream.Receive() {
		t.Fatalf("no first event: %v", stream.Err())
	}
	id := stream.Msg().GetSessionStarted()
	if id == "" {
		t.Fatalf("the first event is not the session id: %v", stream.Msg())
	}
	select {
	case <-p.entered:
	case <-time.After(5 * time.Second):
		t.Fatal("the turn never reached the model")
	}

	res, err := client.GetSession(context.Background(), connect.NewRequest(&mtgv1.GetSessionRequest{SessionId: id}))
	if err != nil {
		close(p.gate)
		t.Fatalf("a reload during the first turn: %v", err)
	}
	if !res.Msg.GetBuilding() {
		t.Error("a reload during the first turn reads no running turn, so the page never polls")
	}
	turns := res.Msg.GetSession().GetTurns()
	if len(turns) != 1 || turns[0].GetUserMessage() != message {
		t.Errorf("the stored turns before the reply = %+v, want the user line alone", turns)
	}
	close(p.gate)
	for stream.Receive() {
	}
	if err := stream.Err(); err != nil {
		t.Fatalf("stream: %v", err)
	}

	res, err = client.GetSession(context.Background(), connect.NewRequest(&mtgv1.GetSessionRequest{SessionId: id}))
	if err != nil {
		t.Fatalf("get session after the turn: %v", err)
	}
	if res.Msg.GetBuilding() {
		t.Error("the turn ended, and the session still reads as running")
	}
	turns = res.Msg.GetSession().GetTurns()
	if len(turns) != 1 || len(turns[0].GetQuestions()) == 0 {
		t.Errorf("the stored turns after the reply = %+v, want one turn with its questions", turns)
	}
}

// TestAFirstTurnBuildKeepsItsLease is F-186 with a build. A first turn
// takes the lease before its model call. When the same turn builds, it
// takes that lease again with its own token, so the build holds the
// whole limit of a build. A second token would refuse the turn.
func TestAFirstTurnBuildKeepsItsLease(t *testing.T) {
	store := newFakeStore()
	var tick atomic.Int64
	clock := func() time.Time { return time.Unix(1000+tick.Add(1), 0).UTC() }
	deck := &mtgv1.Deck{Summary: "a lifegain deck", Validation: &mtgv1.ValidationResult{}}
	fd := &fakeDecks{res: &generate.Result{Deck: deck}, started: make(chan struct{}), release: make(chan struct{})}
	steps := []llm.Step{
		classifyJSON(t, map[string]any{
			"format": "commander", "theme": "lifegain",
			"colors": []string{"W", "B"}, "pool_rule": "any_card",
			"budget_usd": 50, "power": "bracket 3",
			"commander_names": []string{"Karlov of the Ghost Council"},
		}),
		scoreJSON(t),
		askJSON(t),
	}
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithClock(clock)), steps...)
	var first fakeLease
	store.onPut = func(s *mtgv1.Session) {
		store.onPut = nil
		first = store.leases[s.GetId()]
	}

	type result struct {
		got events
		err error
	}
	done := make(chan result, 1)
	go func() {
		got, err := chatE(client, &mtgv1.ChatRequest{Message: "build me a bracket 3 lifegain commander deck with Karlov for 50 dollars"})
		done <- result{got, err}
	}()
	select {
	case <-fd.started:
	case r := <-done:
		t.Fatalf("the first turn ended before a build: %v, failure %v, err %v", r.got.order, r.got.failure, r.err)
	case <-time.After(5 * time.Second):
		t.Fatal("no build started")
	}
	store.mu.Lock()
	var during fakeLease
	for _, l := range store.leases {
		during = l
	}
	store.mu.Unlock()
	close(fd.release)
	r := <-done
	if r.err != nil {
		t.Fatalf("chat: %v", r.err)
	}
	got := r.got

	if got.deck == nil {
		t.Fatalf("no deck: %v, failure %v", got.order, got.failure)
	}
	if first.token == "" {
		t.Fatal("the first write held no lease")
	}
	if during.token != first.token {
		t.Error("the build took a second lease, not the lease of the turn")
	}
	if !during.until.After(first.until) {
		t.Errorf("the build lease ends at %v, the turn lease at %v: the build did not take it again", during.until, first.until)
	}
	if leased, _ := store.Leased(context.Background(), "u1", got.started, clock()); leased {
		t.Error("the lease outlived the turn")
	}
}
