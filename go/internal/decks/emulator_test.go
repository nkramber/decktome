package decks

import (
	"errors"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/protobuf/types/known/timestamppb"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// The store talks to Firestore, and no unit test can reach that. A wrong
// collection path, a bad `firestore` struct tag, or a query without its
// index would pass every green test in the repo, which is the lesson the
// session store learned on 2026-08-25.
//
// These tests need the local emulator, so CI skips them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/decks -count=1

func emulatorRepo(t *testing.T) (*Repo, func()) {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the store against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	return NewRepo(client), func() { _ = client.Close() }
}

func sampleDeck(id string, at time.Time) *mtgv1.Deck {
	return &mtgv1.Deck{
		Id:                 id,
		Name:               "lifegain Commander",
		SessionId:          "s-1",
		Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Summary:            "a lifegain deck",
		CommanderOracleIds: []string{"o-karlov"},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-welcome", Name: "Ajani's Welcome", Count: 1, PriceUsd: 0.25},
		},
		BuyCostUsd:   0.25,
		LegalityAsOf: "2026-08-24",
		CreatedAt:    timestamppb.New(at),
		Validation:   &mtgv1.ValidationResult{},
	}
}

func TestEmulatorDeckRoundTrip(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()

	id := repo.NewID("u1")
	if id == "" {
		t.Fatal("NewID gave an empty id")
	}
	want := sampleDeck(id, time.Now().UTC().Truncate(time.Second))
	if err := repo.Put(ctx, "u1", want); err != nil {
		t.Fatalf("put: %v", err)
	}
	got, err := repo.Get(ctx, "u1", id)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.GetName() != want.GetName() || got.GetBuyCostUsd() != want.GetBuyCostUsd() {
		t.Errorf("deck = %+v", got)
	}
	if len(got.GetCards()) != 1 || got.GetCards()[0].GetName() != "Ajani's Welcome" {
		t.Errorf("cards = %+v", got.GetCards())
	}
	// A deck belongs to its user. Another user must not read it.
	if _, err := repo.Get(ctx, "u2", id); !errors.Is(err, ErrNotFound) {
		t.Errorf("another user read the deck: %v", err)
	}
	// An unknown id is not found, and not an empty deck.
	if _, err := repo.Get(ctx, "u1", "nonesuch"); !errors.Is(err, ErrNotFound) {
		t.Errorf("an unknown id gave %v, want ErrNotFound", err)
	}
}

func TestEmulatorListIsNewestFirst(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()

	uid := "u-list-" + repo.NewID("seed")
	base := time.Now().UTC().Truncate(time.Second)
	var ids []string
	for i := 0; i < 3; i++ {
		id := repo.NewID(uid)
		ids = append(ids, id)
		d := sampleDeck(id, base.Add(time.Duration(i)*time.Minute))
		if err := repo.Put(ctx, uid, d); err != nil {
			t.Fatalf("put %d: %v", i, err)
		}
	}
	got, err := repo.List(ctx, uid, 0)
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("list gave %d decks, want 3", len(got))
	}
	// Newest first, so the last one written comes back first.
	if got[0].GetId() != ids[2] {
		t.Errorf("first deck = %q, want the newest %q", got[0].GetId(), ids[2])
	}
	if limited, err := repo.List(ctx, uid, 2); err != nil || len(limited) != 2 {
		t.Errorf("limited list = %d decks, err %v, want 2", len(limited), err)
	}
	// A user with no decks gets an empty list and no error.
	if empty, err := repo.List(ctx, "u-none", 0); err != nil || len(empty) != 0 {
		t.Errorf("empty list = %d decks, err %v", len(empty), err)
	}
}

func TestPutRefusesADeckWithNoID(t *testing.T) {
	repo := &Repo{}
	if err := repo.Put(t.Context(), "u1", &mtgv1.Deck{}); err == nil {
		t.Error("a deck with no id was accepted")
	}
	if err := repo.Put(t.Context(), "", &mtgv1.Deck{Id: "d1"}); err == nil {
		t.Error("a deck with no user was accepted")
	}
}
