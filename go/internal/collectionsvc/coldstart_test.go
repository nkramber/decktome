package collectionsvc

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
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

func user1(context.Context) string { return "user1" }

// TestTheUploadWaitsForTheFirstIndex is F-176: a collection upload in the
// first seconds after a cold start waits for the first index (D-952).
func TestTheUploadWaitsForTheFirstIndex(t *testing.T) {
	s := New(newFakeRepo(), loadingIndex(testIndex()), user1)
	if _, err := s.ImportCollection(context.Background(), importReq("Binder", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV)); err != nil {
		t.Errorf("upload during the load: %v", err)
	}
	cold := cardsvc.New()
	cold.SetIndexWait(0)
	_, err := New(newFakeRepo(), cold, user1).ImportCollection(context.Background(), importReq("Binder", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, goodCSV))
	if !cardsvc.Loading(err) {
		t.Errorf("upload with no index = %v, want the loading refusal", err)
	}
}

// TestTheBinderWaitsForTheFirstIndex is F-176: a page of the binder read
// in the first seconds after a cold start waits for the first index.
func TestTheBinderWaitsForTheFirstIndex(t *testing.T) {
	_, repo := stocked(t, 3)
	s := New(repo, loadingIndex(cards.NewIndex(nil, nil, nil, time.Time{})), user1)
	if _, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{CollectionId: "col-1"})); err != nil {
		t.Errorf("page read during the load: %v", err)
	}
}
