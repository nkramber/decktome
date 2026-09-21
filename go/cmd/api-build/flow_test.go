package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
)

// fakeAPI stands in for the deployed API. It answers the calls of the
// flow and records what the command sent, so a test reads the whole
// client path without a network and without a provider.
type fakeAPI struct {
	mtgv1connect.UnimplementedAgentServiceHandler
	mtgv1connect.UnimplementedCollectionServiceHandler
	mtgv1connect.UnimplementedDeckServiceHandler
	mtgv1connect.UnimplementedHealthServiceHandler

	mu sync.Mutex
	// bearers holds the Authorization header of every call.
	bearers []string
	// imported holds the bytes of the uploaded collection.
	imported []byte
	// answers holds the answers of every chat turn after the first.
	answers [][]*mtgv1.Answer
	// firstMessage is the prompt of turn one.
	firstMessage string
	poolRule     mtgv1.PoolRule
	turns        int
	deleted      []string
	// questionTurns is how many turns ask questions before the deck.
	questionTurns int
	// coldChecks is how many health checks answer "none" before the
	// snapshot loads. It stands for a cold Cloud Run instance.
	coldChecks int
	// checks counts the health checks the run made.
	checks int
}

func (f *fakeAPI) Check(_ context.Context, r *connect.Request[mtgv1.CheckRequest]) (*connect.Response[mtgv1.CheckResponse], error) {
	f.record(r.Header())
	f.mu.Lock()
	f.checks++
	cold := f.checks <= f.coldChecks
	f.mu.Unlock()
	if cold {
		return connect.NewResponse(&mtgv1.CheckResponse{Status: "ok", CardSnapshot: "none", CardSnapshotAgeHours: -1}), nil
	}
	return connect.NewResponse(&mtgv1.CheckResponse{Status: "ok", CardSnapshot: "2026-09-20"}), nil
}

func (f *fakeAPI) record(h http.Header) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.bearers = append(f.bearers, h.Get("Authorization"))
}

func (f *fakeAPI) ImportCollection(_ context.Context, r *connect.Request[mtgv1.ImportCollectionRequest]) (*connect.Response[mtgv1.ImportCollectionResponse], error) {
	f.record(r.Header())
	f.mu.Lock()
	f.imported = r.Msg.GetContent()
	f.mu.Unlock()
	if r.Msg.GetSource() != mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV {
		return nil, connect.NewError(connect.CodeInvalidArgument, errBadSource)
	}
	return connect.NewResponse(&mtgv1.ImportCollectionResponse{
		Collection: &mtgv1.Collection{
			Id:        "coll-1",
			Name:      r.Msg.GetName(),
			CardCount: 12,
			Summary:   &mtgv1.CollectionSummary{RowCount: 9},
		},
		Report: &mtgv1.ImportReport{ResolvedCount: 9},
	}), nil
}

func (f *fakeAPI) DeleteCollection(_ context.Context, r *connect.Request[mtgv1.DeleteCollectionRequest]) (*connect.Response[mtgv1.DeleteCollectionResponse], error) {
	f.record(r.Header())
	f.mu.Lock()
	f.deleted = append(f.deleted, "collection:"+r.Msg.GetCollectionId())
	f.mu.Unlock()
	return connect.NewResponse(&mtgv1.DeleteCollectionResponse{}), nil
}

func (f *fakeAPI) Chat(_ context.Context, r *connect.Request[mtgv1.ChatRequest], s *connect.ServerStream[mtgv1.ChatResponse]) error {
	f.record(r.Header())
	f.mu.Lock()
	f.turns++
	turn := f.turns
	if turn == 1 {
		f.firstMessage = r.Msg.GetMessage()
		f.poolRule = r.Msg.GetPoolRule()
	} else {
		f.answers = append(f.answers, r.Msg.GetAnswers())
	}
	f.mu.Unlock()

	if turn == 1 {
		if err := s.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_SessionStarted{SessionStarted: "sess-1"}}); err != nil {
			return err
		}
	}
	if turn <= f.questionTurns {
		return s.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Question{Question: &mtgv1.Question{
			Id:      "q" + string(rune('0'+turn)),
			Slot:    []string{"format", "commander"}[turn-1],
			Text:    "which one?",
			Options: []string{"Commander", "Modern"},
			Closed:  turn == 1,
		}}})
	}
	return s.Send(&mtgv1.ChatResponse{Event: &mtgv1.ChatResponse_Deck{Deck: &mtgv1.Deck{
		Id:   "deck-1",
		Name: "the built deck",
	}}})
}

func (f *fakeAPI) GetSession(_ context.Context, r *connect.Request[mtgv1.GetSessionRequest]) (*connect.Response[mtgv1.GetSessionResponse], error) {
	f.record(r.Header())
	return connect.NewResponse(&mtgv1.GetSessionResponse{Session: &mtgv1.Session{
		Id:    r.Msg.GetSessionId(),
		Turns: []*mtgv1.Turn{{UserMessage: "one"}, {UserMessage: "two"}},
	}}), nil
}

func (f *fakeAPI) DeleteSession(_ context.Context, r *connect.Request[mtgv1.DeleteSessionRequest]) (*connect.Response[mtgv1.DeleteSessionResponse], error) {
	f.record(r.Header())
	f.mu.Lock()
	f.deleted = append(f.deleted, "session:"+r.Msg.GetSessionId())
	f.mu.Unlock()
	return connect.NewResponse(&mtgv1.DeleteSessionResponse{}), nil
}

func (f *fakeAPI) GetDeck(_ context.Context, r *connect.Request[mtgv1.GetDeckRequest]) (*connect.Response[mtgv1.GetDeckResponse], error) {
	f.record(r.Header())
	if r.Msg.GetDeckId() != "deck-1" {
		return nil, connect.NewError(connect.CodeNotFound, errNoDeck)
	}
	return connect.NewResponse(&mtgv1.GetDeckResponse{Deck: &mtgv1.Deck{
		Id:    "deck-1",
		Name:  "the built deck",
		Cards: []*mtgv1.DeckCard{{Count: 99}},
	}}), nil
}

func (f *fakeAPI) ListDecks(_ context.Context, r *connect.Request[mtgv1.ListDecksRequest]) (*connect.Response[mtgv1.ListDecksResponse], error) {
	f.record(r.Header())
	if r.Msg.GetSessionId() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errNoSession)
	}
	return connect.NewResponse(&mtgv1.ListDecksResponse{Decks: []*mtgv1.Deck{{Id: "deck-1"}}}), nil
}

func (f *fakeAPI) DeleteDeck(_ context.Context, r *connect.Request[mtgv1.DeleteDeckRequest]) (*connect.Response[mtgv1.DeleteDeckResponse], error) {
	f.record(r.Header())
	f.mu.Lock()
	f.deleted = append(f.deleted, "deck:"+r.Msg.GetDeckId())
	f.mu.Unlock()
	return connect.NewResponse(&mtgv1.DeleteDeckResponse{}), nil
}

// errors of the fake, named so the test reads them.
var (
	errBadSource = &staticError{"the source is not a ManaBox CSV"}
	errNoDeck    = &staticError{"no such deck"}
	errNoSession = &staticError{"no session id on the deck list"}
)

type staticError struct{ s string }

func (e *staticError) Error() string { return e.s }

// serve starts the fake API and the sign-in endpoint on one server.
func serve(t *testing.T, f *fakeAPI) (base, signin string) {
	t.Helper()
	mux := http.NewServeMux()
	mux.Handle(mtgv1connect.NewAgentServiceHandler(f))
	mux.Handle(mtgv1connect.NewCollectionServiceHandler(f))
	mux.Handle(mtgv1connect.NewDeckServiceHandler(f))
	mux.Handle(mtgv1connect.NewHealthServiceHandler(f))
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("key") != "test-key" {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]string{"message": "API_KEY_INVALID"}})
			return
		}
		var in map[string]any
		_ = json.NewDecoder(r.Body).Decode(&in)
		if in["password"] != "right" {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]any{"error": map[string]string{"message": "INVALID_LOGIN_CREDENTIALS"}})
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"idToken": "token-1", "localId": "uid-1", "email": in["email"], "expiresIn": "3600",
		})
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv.URL, srv.URL + "/signin"
}

// csvFixture writes a small ManaBox file the fake reads back.
func csvFixture(t *testing.T) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "collection.csv")
	body := "Name,Set code,Quantity\nSol Ring,C21,1\n"
	if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	return p
}

func testOptions(base, csv string) options {
	return options{
		base:     base,
		csv:      csv,
		name:     "check",
		prompt:   "Build me a lifegain Commander deck.",
		poolRule: "owned-first",
		maxTurns: 8,
		timeout:  30 * time.Second,
		ready:    30 * time.Second,
	}
}

func goodCredentials(signin string) credentials {
	return credentials{endpoint: signin, apiKey: "test-key", email: "check@example.com", password: "right"}
}

// The whole flow of D-778: the sign-in, the import, the chat with an
// answer to each question, and the read-back of the stored deck.
func TestBuildWalksTheWholeFlow(t *testing.T) {
	f := &fakeAPI{questionTurns: 2}
	base, signin := serve(t, f)
	csv := csvFixture(t)

	if err := build(context.Background(), testOptions(base, csv), goodCredentials(signin)); err != nil {
		t.Fatalf("build: %v", err)
	}

	if f.turns != 3 {
		t.Errorf("the chat took %d turns, want 3", f.turns)
	}
	if f.firstMessage != "Build me a lifegain Commander deck." {
		t.Errorf("the first message was %q", f.firstMessage)
	}
	if f.poolRule != mtgv1.PoolRule_POOL_RULE_OWNED_FIRST {
		t.Errorf("the pool rule was %v", f.poolRule)
	}
	if !strings.Contains(string(f.imported), "Sol Ring") {
		t.Error("the collection file did not reach the import")
	}
	if len(f.answers) != 2 {
		t.Fatalf("the run sent %d turns of answers, want 2", len(f.answers))
	}
	for i, as := range f.answers {
		if len(as) != 1 {
			t.Fatalf("turn %d sent %d answers, want 1", i+2, len(as))
		}
		if as[0].GetQuestionId() == "" {
			t.Errorf("turn %d answered no question id", i+2)
		}
	}
	// The first question is closed, so it takes an option and never a
	// decline (D-295).
	if f.answers[0][0].OptionIndex == nil {
		t.Error("the closed question took no option")
	}
	if len(f.deleted) != 0 {
		t.Errorf("the run deleted %v with no -cleanup flag (D-780)", f.deleted)
	}
}

// Every call carries the id token, the streamed Chat call included.
func TestBuildSendsTheTokenOnEveryCall(t *testing.T) {
	f := &fakeAPI{questionTurns: 1}
	base, signin := serve(t, f)

	if err := build(context.Background(), testOptions(base, csvFixture(t)), goodCredentials(signin)); err != nil {
		t.Fatalf("build: %v", err)
	}
	if len(f.bearers) < 5 {
		t.Fatalf("the fake saw %d calls, want the import, two chats, and the three reads", len(f.bearers))
	}
	for i, got := range f.bearers {
		if got != "Bearer token-1" {
			t.Errorf("call %d carried %q", i, got)
		}
	}
}

func TestBuildCleanupDeletesWhatItMade(t *testing.T) {
	f := &fakeAPI{}
	base, signin := serve(t, f)
	o := testOptions(base, csvFixture(t))
	o.cleanup = true

	if err := build(context.Background(), o, goodCredentials(signin)); err != nil {
		t.Fatalf("build: %v", err)
	}
	want := []string{"deck:deck-1", "session:sess-1", "collection:coll-1"}
	if len(f.deleted) != len(want) {
		t.Fatalf("the run deleted %v, want %v", f.deleted, want)
	}
	for i, w := range want {
		if f.deleted[i] != w {
			t.Errorf("delete %d was %q, want %q", i, f.deleted[i], w)
		}
	}
}

func TestBuildStopsAtMaxTurns(t *testing.T) {
	f := &fakeAPI{questionTurns: 99}
	base, signin := serve(t, f)
	o := testOptions(base, csvFixture(t))
	o.maxTurns = 2

	err := build(context.Background(), o, goodCredentials(signin))
	if err == nil {
		t.Fatal("a run that never builds a deck reported success")
	}
	if !strings.Contains(err.Error(), "no deck after 2 turns") {
		t.Errorf("the error does not name the turn limit: %v", err)
	}
}

func TestBuildReportsASignInRefusal(t *testing.T) {
	f := &fakeAPI{}
	base, signin := serve(t, f)
	cred := goodCredentials(signin)
	cred.password = "wrong"

	err := build(context.Background(), testOptions(base, csvFixture(t)), cred)
	if err == nil {
		t.Fatal("a wrong password signed in")
	}
	if !strings.Contains(err.Error(), "INVALID_LOGIN_CREDENTIALS") {
		t.Errorf("the error does not name the refusal: %v", err)
	}
	if strings.Contains(err.Error(), "wrong") {
		t.Error("the error carries the password")
	}
	if f.turns != 0 {
		t.Error("a refused sign-in still called the API")
	}
}

func TestBuildNeedsAnAccount(t *testing.T) {
	f := &fakeAPI{}
	base, signin := serve(t, f)
	cred := goodCredentials(signin)
	cred.email = ""

	err := build(context.Background(), testOptions(base, csvFixture(t)), cred)
	if err == nil || !strings.Contains(err.Error(), "API_BUILD_EMAIL") {
		t.Errorf("a run with no account did not say what to set: %v", err)
	}
}

// The invite gate is the refusal a new account reads (D-314). The run
// must name the command that fixes it.
func TestExplainNamesTheInviteCommand(t *testing.T) {
	err := connect.NewError(connect.CodePermissionDenied, errNoDeck)
	err.Meta().Set("Deck-Tome-Refusal", "not-invited")
	got := explain("the import", err)
	if !strings.Contains(got.Error(), "make allow") {
		t.Errorf("the refusal does not name the invite command: %v", got)
	}
	plain := explain("the import", errNoDeck)
	if strings.Contains(plain.Error(), "make allow") {
		t.Errorf("an ordinary error read as an invite refusal: %v", plain)
	}
}

// A Cloud Run instance scales to zero. A cold one answers an import
// with "card database not loaded yet", so the run waits for the
// snapshot before it imports.
func TestBuildWaitsForAColdInstance(t *testing.T) {
	old := readyPoll
	readyPoll = time.Millisecond
	t.Cleanup(func() { readyPoll = old })

	f := &fakeAPI{coldChecks: 2}
	base, signin := serve(t, f)

	if err := build(context.Background(), testOptions(base, csvFixture(t)), goodCredentials(signin)); err != nil {
		t.Fatalf("build: %v", err)
	}
	if f.checks < 3 {
		t.Errorf("the run made %d health checks, want at least 3", f.checks)
	}
	if f.imported == nil {
		t.Error("the run imported nothing after the wait")
	}
}

func TestBuildGivesUpOnAnInstanceThatNeverLoads(t *testing.T) {
	f := &fakeAPI{coldChecks: 1000}
	base, signin := serve(t, f)
	old := readyPoll
	readyPoll = time.Millisecond
	t.Cleanup(func() { readyPoll = old })

	o := testOptions(base, csvFixture(t))
	o.ready = 10 * time.Millisecond

	err := build(context.Background(), o, goodCredentials(signin))
	if err == nil {
		t.Fatal("a run against an API with no snapshot reported success")
	}
	if !strings.Contains(err.Error(), "no card snapshot") {
		t.Errorf("the error does not name the snapshot: %v", err)
	}
	if f.imported != nil {
		t.Error("the run imported before the API was ready")
	}
}
