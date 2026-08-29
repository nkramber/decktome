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
// index passes every unit test, so these tests run the paths against
// the emulator.
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
	got, err := repo.List(ctx, uid, Filter{}, 0)
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
	if limited, err := repo.List(ctx, uid, Filter{}, 2); err != nil || len(limited) != 2 {
		t.Errorf("limited list = %d decks, err %v, want 2", len(limited), err)
	}
	// A user with no decks gets an empty list and no error.
	if empty, err := repo.List(ctx, "u-none", Filter{}, 0); err != nil || len(empty) != 0 {
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

// PR-17 added Update, Delete, and five flat fields. A wrong struct tag or
// a transaction that reads the wrong document passes every unit test, so
// these paths run against the emulator too.

func TestEmulatorUpdateDeck(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()

	// A fresh user per run, the convention of the test above. A cleanup
	// can not delete: t.Context is cancelled before a cleanup runs.
	uid := "u-update-" + repo.NewID("seed")
	id := repo.NewID(uid)
	created := time.Now().UTC().Truncate(time.Second)
	d := sampleDeck(id, created)
	if err := repo.Put(ctx, uid, d); err != nil {
		t.Fatalf("put: %v", err)
	}

	name := "Karlov, renamed"
	yes := true
	got, err := repo.Update(ctx, uid, id, &name, &yes)
	if err != nil {
		t.Fatalf("update: %v", err)
	}
	if got.GetName() != name || !got.GetFavorite() {
		t.Errorf("update returned name %q favorite %v", got.GetName(), got.GetFavorite())
	}

	// The packed proto and the flat fields must agree, or the list and
	// the deck page show two different names.
	whole, err := repo.Get(ctx, uid, id)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if whole.GetName() != name || !whole.GetFavorite() {
		t.Errorf("get returned name %q favorite %v", whole.GetName(), whole.GetFavorite())
	}
	list, err := repo.List(ctx, uid, Filter{}, 0)
	if err != nil || len(list) != 1 {
		t.Fatalf("list = %v, err %v", list, err)
	}
	if list[0].GetName() != name || !list[0].GetFavorite() {
		t.Errorf("list view name %q favorite %v", list[0].GetName(), list[0].GetFavorite())
	}
	if list[0].GetCardCount() != 1 {
		t.Errorf("card count = %d, want 1", list[0].GetCardCount())
	}
	if len(list[0].GetCommanderOracleIds()) != 1 {
		t.Errorf("commanders = %v", list[0].GetCommanderOracleIds())
	}
	// A rename must not move the deck to the top of the listing.
	if at := list[0].GetCreatedAt().AsTime().UTC().Truncate(time.Second); !at.Equal(created) {
		t.Errorf("created_at = %v, want %v", at, created)
	}

	// One field alone leaves the other as it is.
	no := false
	if _, err := repo.Update(ctx, uid, id, nil, &no); err != nil {
		t.Fatalf("update mark: %v", err)
	}
	after, err := repo.Get(ctx, uid, id)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if after.GetName() != name || after.GetFavorite() {
		t.Errorf("after the mark write: name %q favorite %v", after.GetName(), after.GetFavorite())
	}

	if _, err := repo.Update(ctx, uid, "no-such-deck", &name, nil); !errors.Is(err, ErrNotFound) {
		t.Errorf("update of an unknown deck = %v, want ErrNotFound", err)
	}
}

func TestEmulatorDeleteDeck(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()

	uid := "u-delete-" + repo.NewID("seed")
	id := repo.NewID(uid)
	if err := repo.Put(ctx, uid, sampleDeck(id, time.Now().UTC())); err != nil {
		t.Fatalf("put: %v", err)
	}
	if err := repo.Delete(ctx, uid, id); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if _, err := repo.Get(ctx, uid, id); !errors.Is(err, ErrNotFound) {
		t.Errorf("get after delete = %v, want ErrNotFound", err)
	}
	if err := repo.Delete(ctx, uid, id); !errors.Is(err, ErrNotFound) {
		t.Errorf("second delete = %v, want ErrNotFound", err)
	}
}

func TestEmulatorListFilter(t *testing.T) {
	repo, done := emulatorRepo(t)
	defer done()
	ctx := t.Context()

	uid := "u-filter-" + repo.NewID("seed")
	now := time.Now().UTC()
	starred := repo.NewID(uid)
	plain := repo.NewID(uid)
	a := sampleDeck(starred, now)
	a.Name = "lifegain pile"
	a.Favorite = true
	// The commander sits in the card list, so its name reaches the flat
	// field a commander search reads.
	a.Cards = append(a.Cards, &mtgv1.DeckCard{OracleId: "o-karlov", Name: "Karlov of the Ghost Council", Count: 1})
	b := sampleDeck(plain, now.Add(-time.Minute))
	b.Name = "Goblin storm"
	b.CommanderOracleIds = nil
	b.Format = &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN}
	for _, d := range []*mtgv1.Deck{a, b} {
		if err := repo.Put(ctx, uid, d); err != nil {
			t.Fatalf("put: %v", err)
		}
	}
	yes := true
	for _, tc := range []struct {
		name string
		f    Filter
		want []string
	}{
		{"no filter, newest first", Filter{}, []string{starred, plain}},
		{"the favorites", Filter{Favorite: &yes}, []string{starred}},
		{"one format", Filter{Format: mtgv1.FormatId_FORMAT_ID_MODERN}, []string{plain}},
		{"a name search", Filter{Query: "goblin"}, []string{plain}},
		{"a commander search", Filter{Query: "ghost council"}, []string{starred}},
		{"a search reads the name, not the card list", Filter{Query: "ajani"}, nil},
		{"no match", Filter{Query: "zzz"}, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := repo.List(ctx, uid, tc.f, 0)
			if err != nil {
				t.Fatalf("list: %v", err)
			}
			var ids []string
			for _, d := range got {
				ids = append(ids, d.GetId())
			}
			if len(ids) != len(tc.want) {
				t.Fatalf("ids = %v, want %v", ids, tc.want)
			}
			for i := range ids {
				if ids[i] != tc.want[i] {
					t.Fatalf("ids = %v, want %v", ids, tc.want)
				}
			}
		})
	}
}
