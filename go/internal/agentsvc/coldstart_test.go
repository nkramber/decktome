package agentsvc

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/cardsvc"
)

// loadingIndex is the card server of a cold start. It installs idx after
// a short load, and a read waits for the first index up to 5 seconds.
func loadingIndex(idx *cards.Index) *cardsvc.Server {
	s := cardsvc.New()
	s.SetIndexWait(5 * time.Second)
	go func() {
		time.Sleep(30 * time.Millisecond)
		s.Swap(idx)
	}()
	return s
}

// TestChatWaitsForTheFirstIndex is F-176. The API listens before the
// card snapshot loads, and a turn of the first seconds after a cold start
// read Unavailable at once. It waits for the first index now, as a card
// RPC does (D-952).
func TestChatWaitsForTheFirstIndex(t *testing.T) {
	store := newFakeStore()
	idx := cards.NewIndex(nil, nil, nil, time.Unix(1000, 0).UTC())
	client, _ := testServerOpts(t, store, []Option{WithCandidates(loadingIndex(idx), nil)}, firstTurn(t)...)
	got := chat(t, client, &mtgv1.ChatRequest{Message: "build me a lifegain commander deck for 50 dollars"})
	if got.started == "" || store.sessions[got.started] == nil {
		t.Fatalf("the turn during the load stored no session: %v", got.order)
	}
}

// TestAChatRefusedBeforeTheIndexChangesNothing is D-954. The web sends a
// turn again on the loading refusal alone, so that refusal must come
// before any state change of the turn, and it must carry its header.
func TestAChatRefusedBeforeTheIndexChangesNothing(t *testing.T) {
	store := newFakeStore()
	src := cardsvc.New()
	src.SetIndexWait(0)
	client, _ := testServerOpts(t, store, []Option{WithCandidates(src, nil)}, firstTurn(t)...)
	stream, err := client.Chat(context.Background(), connect.NewRequest(&mtgv1.ChatRequest{Message: "Build me an anime-themed commander deck from my collection."}))
	if err != nil {
		t.Fatalf("chat: %v", err)
	}
	defer func() { _ = stream.Close() }()
	for stream.Receive() {
		t.Errorf("a refused turn sent an event: %v", stream.Msg().GetEvent())
	}
	if !cardsvc.Loading(stream.Err()) {
		t.Errorf("err = %v, want the loading refusal", stream.Err())
	}
	if len(store.sessions) != 0 {
		t.Errorf("a refused turn stored %d sessions", len(store.sessions))
	}
}

// TestImportDeckWaitsForTheFirstIndex is F-176: an import of a deck list
// in the first seconds after a cold start waits for the first index.
func TestImportDeckWaitsForTheFirstIndex(t *testing.T) {
	cb, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	fd, ds := &fakeDecks{}, &fakeDeckStore{}
	opts := []Option{WithDecks(fd), WithCandidates(loadingIndex(importCardIndex()), cb), WithDeckStore(ds), WithUsers(&fakeNoter{})}
	client, _ := testServerOpts(t, newFakeStore(), opts)
	if res := importList(t, client, &mtgv1.ImportDeckRequest{Text: markedList, Name: "Living Weapon"}); res.GetDeck() == nil {
		t.Fatal("the import during the load kept no deck")
	}
}
