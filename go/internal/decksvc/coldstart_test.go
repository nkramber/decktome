package decksvc

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/cardsvc"
	"github.com/nkramber/decktome/go/internal/rules"
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

// TestValidateWaitsForTheFirstIndex is F-176: a check of a deck in the
// first seconds after a cold start waits for the first index (D-952).
func TestValidateWaitsForTheFirstIndex(t *testing.T) {
	cfg, err := rules.Load()
	if err != nil {
		t.Fatal(err)
	}
	idx := loadIndex(t)
	s := New(cfg, loadingIndex(idx))
	if _, err := s.Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{Deck: modernDeck(t, idx)})); err != nil {
		t.Errorf("Validate during the load: %v", err)
	}
	cold := cardsvc.New()
	cold.SetIndexWait(0)
	_, err = New(cfg, cold).Validate(context.Background(), connect.NewRequest(&mtgv1.ValidateRequest{Deck: &mtgv1.Deck{}}))
	if !cardsvc.Loading(err) {
		t.Errorf("Validate with no index = %v, want the loading refusal", err)
	}
}

// TestTheSharedDeckWaitsForTheFirstIndex is F-176: a shared deck opened
// in the first seconds after a cold start waits for the first index.
func TestTheSharedDeckWaitsForTheFirstIndex(t *testing.T) {
	ctx := context.Background()
	f := &fakeDecks{decks: map[string]*mtgv1.Deck{"d1": sharedDeckFixture()}}
	warm := shareServer(t, f, "u1")
	res, err := warm.ShareDeck(ctx, connect.NewRequest(&mtgv1.ShareDeckRequest{DeckId: "d1"}))
	if err != nil {
		t.Fatal(err)
	}
	cold := New(nil, loadingIndex(warm.index.Current()), WithDecks(f))
	got, err := cold.GetSharedDeck(ctx, connect.NewRequest(&mtgv1.GetSharedDeckRequest{Token: res.Msg.Token}))
	if err != nil {
		t.Fatalf("shared read during the load: %v", err)
	}
	if len(got.Msg.GetDeck().GetCards()) == 0 {
		t.Error("the shared deck holds no card")
	}
}
