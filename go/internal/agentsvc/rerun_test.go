package agentsvc

import (
	"context"
	"slices"
	"strings"
	"testing"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/gen/mtg/v1/mtgv1connect"
	"github.com/nkramber/decktome/go/internal/generate"
)

// staleBuild builds one deck in a session, then marks the stored deck
// stale as the stale pass would (I-1). It returns the client, the session
// id, and the stores.
func staleBuild(t *testing.T, rerun mtgv1.RerunCase, staleIDs ...string) (mtgv1connect.AgentServiceClient, string, *fakeStore, *fakeDeckStore, *fakeDecks) {
	t.Helper()
	store := newFakeStore()
	ds := &fakeDeckStore{}
	base := &mtgv1.Deck{
		Name: "lifegain Commander", Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 1, Role: mtgv1.CardRole_CARD_ROLE_SYNERGY},
			{OracleId: "o-plains", Name: "Plains", Count: 30, Role: mtgv1.CardRole_CARD_ROLE_LAND},
		},
	}
	fd := &fakeDecks{res: &generate.Result{Deck: base}}
	steps := builtSteps(t)
	if slices.Contains(staleIDs, "o-karlov") {
		steps = append(steps, classifyJSON(t, nil))
	}
	client, _ := testServerOpts(t, store, append(buildOpts(t, fd), WithDeckStore(ds)), steps...)
	first := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	second := chat(t, client, &mtgv1.ChatRequest{SessionId: first.started, Message: "Karlov, bracket 3"})
	if second.deck == nil || fd.runs != 1 || len(ds.put) != 1 {
		t.Fatalf("no first deck: runs=%d order=%v", fd.runs, second.order)
	}
	d := ds.put[0]
	d.CommanderOracleIds = []string{"o-karlov"}
	d.Commanders = []*mtgv1.DeckCard{{OracleId: "o-karlov", Name: "Karlov of the Ghost Council", Count: 1}}
	d.Stale, d.StaleOracleIds, d.RerunCase = true, staleIDs, rerun
	d.StaleReason = "the reason of the banner"
	return client, first.started, store, ds, fd
}

// TestRerunPatchRevisesTheStaleDeck is D-1019: a patch is a revision of
// the stale deck with its stale cards in Remove, and no revise call.
func TestRerunPatchRevisesTheStaleDeck(t *testing.T) {
	client, sid, store, _, fd := staleBuild(t, mtgv1.RerunCase_RERUN_CASE_PATCH, "o-welcome")
	fd.res = &generate.Result{Deck: &mtgv1.Deck{
		Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-inspiring", Name: "Inspiring Overseer", Count: 1},
			{OracleId: "o-plains", Name: "Plains", Count: 30},
		},
	}}
	got := chat(t, client, &mtgv1.ChatRequest{SessionId: sid, RerunDeckId: "deck-1"})
	if fd.runs != 2 {
		t.Fatalf("build runs = %d, want 2", fd.runs)
	}
	rev := fd.got.Revision
	if rev == nil || rev.BaseDeckID != "deck-1" || !slices.Equal(rev.Remove, []string{"Ajani's Welcome"}) {
		t.Fatalf("revision = %+v", rev)
	}
	if got.deck == nil || got.deck.GetRevisedFromDeckId() != "deck-1" {
		t.Fatalf("no patched deck: %v", got.order)
	}
	if !slices.Contains(got.statuses, "the reason of the banner") {
		t.Errorf("statuses = %v", got.statuses)
	}
	turns := store.sessions[sid].GetTurns()
	last := turns[len(turns)-1]
	if last.GetUserMessage() != rerunText || !strings.Contains(last.GetRevisionBrief(), "Ajani's Welcome") {
		t.Errorf("turn = %q, brief %q", last.GetUserMessage(), last.GetRevisionBrief())
	}
}

// TestRerunRebuildBuildsFromTheConversation is D-1020: a rebuild with a
// legal commander builds again from the slots, with no revision and no
// classify call.
func TestRerunRebuildBuildsFromTheConversation(t *testing.T) {
	client, sid, store, ds, fd := staleBuild(t, mtgv1.RerunCase_RERUN_CASE_REBUILD, "o-welcome")
	fd.res = &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{},
		Cards: []*mtgv1.DeckCard{{OracleId: "o-plains", Name: "Plains", Count: 31}}}}
	got := chat(t, client, &mtgv1.ChatRequest{SessionId: sid, RerunDeckId: "deck-1"})
	if fd.runs != 2 || fd.got.Revision != nil {
		t.Fatalf("runs = %d, revision %+v", fd.runs, fd.got.Revision)
	}
	if got.deck == nil || len(ds.put) != 2 || len(store.sessions[sid].GetDeckIds()) != 2 {
		t.Fatalf("no rebuilt deck: %v, decks %d", got.order, len(ds.put))
	}
	if got.usage == nil {
		t.Errorf("the rerun sent no usage: %v", got.order)
	}
}

// TestRerunBannedCommanderClearsTheCommander is D-1021: a ban of the
// commander builds nothing yet. The turn clears the commander, so the
// question flow asks for a new one.
func TestRerunBannedCommanderClearsTheCommander(t *testing.T) {
	client, sid, store, _, fd := staleBuild(t, mtgv1.RerunCase_RERUN_CASE_REBUILD, "o-karlov")
	got := chat(t, client, &mtgv1.ChatRequest{SessionId: sid, RerunDeckId: "deck-1"})
	if fd.runs != 1 {
		t.Fatalf("a banned commander must not build before the pick, runs = %d", fd.runs)
	}
	if !slices.Contains(got.statuses, "the reason of the banner") {
		t.Errorf("statuses = %v", got.statuses)
	}
	s := store.sessions[sid]
	if ids := s.GetSlots().GetCommanderOracleIds(); len(ids) > 0 {
		t.Errorf("the banned commander stayed in the slots: %v", ids)
	}
	turns := s.GetTurns()
	if last := turns[len(turns)-1]; last.GetUserMessage() != rerunText {
		t.Errorf("turn = %q", last.GetUserMessage())
	}
}

// TestRerunRefusals refuses a rerun of the wrong shape, and a rerun of a
// deck that is legal.
func TestRerunRefusals(t *testing.T) {
	client, sid, _, ds, _ := staleBuild(t, mtgv1.RerunCase_RERUN_CASE_PATCH, "o-welcome")
	cases := []struct {
		name string
		req  *mtgv1.ChatRequest
		code connect.Code
	}{
		{"no session", &mtgv1.ChatRequest{RerunDeckId: "deck-1"}, connect.CodeInvalidArgument},
		{"a message too", &mtgv1.ChatRequest{SessionId: sid, RerunDeckId: "deck-1", Message: "hi"}, connect.CodeInvalidArgument},
		{"a bad id", &mtgv1.ChatRequest{SessionId: sid, RerunDeckId: "a/b"}, connect.CodeInvalidArgument},
		{"a deck of another session", &mtgv1.ChatRequest{SessionId: sid, RerunDeckId: "deck-9"}, connect.CodeInvalidArgument},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if code := chatCode(t, client, c.req); code != c.code {
				t.Fatalf("code = %v, want %v", code, c.code)
			}
		})
	}
	ds.put[0].Stale = false
	if code := chatCode(t, client, &mtgv1.ChatRequest{SessionId: sid, RerunDeckId: "deck-1"}); code != connect.CodeFailedPrecondition {
		t.Fatalf("a legal deck: code = %v", code)
	}
}

// chatCode sends one turn and returns the code of its error, or 0.
func chatCode(t *testing.T, c mtgv1connect.AgentServiceClient, req *mtgv1.ChatRequest) connect.Code {
	t.Helper()
	stream, err := c.Chat(context.Background(), connect.NewRequest(req))
	if err != nil {
		return connect.CodeOf(err)
	}
	for stream.Receive() {
	}
	if err := stream.Err(); err != nil {
		return connect.CodeOf(err)
	}
	return 0
}
