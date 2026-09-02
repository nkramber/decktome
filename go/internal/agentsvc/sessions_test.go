package agentsvc

import (
	"context"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// stockedStore holds n sessions, the newest first by updated_at.
func stockedStore(n int) *fakeStore {
	store := newFakeStore()
	for i := 0; i < n; i++ {
		id := "s-" + strings.Repeat("x", i+1)
		store.sessions[id] = &mtgv1.Session{
			Id: id, Status: mtgv1.SessionStatus_SESSION_STATUS_ASKING,
			Turns:     []*mtgv1.Turn{{UserMessage: "request " + id}},
			UpdatedAt: timestamppb.New(timestamppb.Now().AsTime().Add(-time.Duration(i) * time.Minute)),
		}
		store.versions[id] = 1
	}
	return store
}

// TestListSessionsPagesNewestFirst is D-433. The list reads summaries,
// and a token of another reader is refused.
func TestListSessionsPagesNewestFirst(t *testing.T) {
	store := stockedStore(5)
	client, _ := testServerOpts(t, store, nil)
	first, err := client.ListSessions(context.Background(), connect.NewRequest(&mtgv1.ListSessionsRequest{PageSize: 2}))
	if err != nil {
		t.Fatal(err)
	}
	got := first.Msg.GetSessions()
	if len(got) != 2 || got[0].GetId() != "s-x" || got[0].GetFirstMessage() != "request s-x" {
		t.Fatalf("first page = %v, want the two newest with their first message", got)
	}
	if first.Msg.GetNextPageToken() == "" {
		t.Fatal("a first page of five returns a token")
	}
	seen, token := len(got), first.Msg.GetNextPageToken()
	for token != "" {
		res, err := client.ListSessions(context.Background(), connect.NewRequest(&mtgv1.ListSessionsRequest{PageSize: 2, PageToken: token}))
		if err != nil {
			t.Fatal(err)
		}
		seen += len(res.Msg.GetSessions())
		token = res.Msg.GetNextPageToken()
	}
	if seen != 5 {
		t.Errorf("read %d sessions over the pages, want 5", seen)
	}
	_, err = client.ListSessions(context.Background(), connect.NewRequest(&mtgv1.ListSessionsRequest{PageToken: "nonsense"}))
	if connect.CodeOf(err) != connect.CodeInvalidArgument {
		t.Errorf("a bad token gave %v, want InvalidArgument", connect.CodeOf(err))
	}
}

// TestUpdateAndDeleteSession: a rename writes the name and keeps the
// turns, and a delete removes the session and refuses a second time.
func TestUpdateAndDeleteSession(t *testing.T) {
	store := stockedStore(1)
	client, _ := testServerOpts(t, store, nil)
	res, err := client.UpdateSession(context.Background(), connect.NewRequest(&mtgv1.UpdateSessionRequest{SessionId: "s-x", Name: "  Elves  "}))
	if err != nil {
		t.Fatal(err)
	}
	if res.Msg.GetSession().GetName() != "Elves" || store.sessions["s-x"].GetName() != "Elves" {
		t.Errorf("rename gave %q, store holds %q", res.Msg.GetSession().GetName(), store.sessions["s-x"].GetName())
	}
	if len(store.sessions["s-x"].GetTurns()) != 1 {
		t.Error("a rename lost a turn")
	}
	for _, tc := range []struct {
		name, id, newName string
		want              connect.Code
	}{
		{"empty name", "s-x", " ", connect.CodeInvalidArgument},
		{"long name", "s-x", strings.Repeat("x", MaxSessionNameBytes+1), connect.CodeInvalidArgument},
		{"bad id", "../x", "ok", connect.CodeInvalidArgument},
		{"missing", "missing", "ok", connect.CodeNotFound},
	} {
		_, err := client.UpdateSession(context.Background(), connect.NewRequest(&mtgv1.UpdateSessionRequest{SessionId: tc.id, Name: tc.newName}))
		if connect.CodeOf(err) != tc.want {
			t.Errorf("%s: code = %v, want %v", tc.name, connect.CodeOf(err), tc.want)
		}
	}
	if _, err := client.DeleteSession(context.Background(), connect.NewRequest(&mtgv1.DeleteSessionRequest{SessionId: "s-x"})); err != nil {
		t.Fatal(err)
	}
	if _, ok := store.sessions["s-x"]; ok {
		t.Error("the session is still stored")
	}
	_, err = client.DeleteSession(context.Background(), connect.NewRequest(&mtgv1.DeleteSessionRequest{SessionId: "s-x"}))
	if connect.CodeOf(err) != connect.CodeNotFound {
		t.Errorf("a second delete gave %v, want NotFound", connect.CodeOf(err))
	}
}

// TestDeleteSessionTakesItsDecks is D-456: a chat and its decks are one
// thing, and a deck of another chat stays.
func TestDeleteSessionTakesItsDecks(t *testing.T) {
	store := stockedStore(1)
	store.sessions["s-x"].DeckIds = []string{"deck-a", "deck-b"}
	ds := &fakeDeckStore{put: []*mtgv1.Deck{{Id: "deck-a"}, {Id: "deck-b"}, {Id: "deck-other"}}}
	client, _ := testServerOpts(t, store, []Option{WithDeckStore(ds)})
	if _, err := client.DeleteSession(context.Background(), connect.NewRequest(&mtgv1.DeleteSessionRequest{SessionId: "s-x"})); err != nil {
		t.Fatal(err)
	}
	if len(ds.put) != 1 || ds.put[0].GetId() != "deck-other" {
		t.Errorf("decks left = %v, want deck-other alone", ds.put)
	}
	if _, ok := store.sessions["s-x"]; ok {
		t.Error("the session is still stored")
	}
}
