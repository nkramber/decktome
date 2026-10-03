package access

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
)

// These tests need the local emulator, so CI skips them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/access -count=1

func emulatorClient(t *testing.T) *firestore.Client {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the store against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return client
}

// TestRequestRoundTrip is D-1075: the first request makes a pending
// record, a repeat counts and keeps the state, the list reads the
// newest first, and a decision moves the record.
func TestRequestRoundTrip(t *testing.T) {
	client := emulatorClient(t)
	ctx := context.Background()
	repo := NewRepo(client)
	stamp := time.Now().UTC().Format("150405.000000")
	ann, bob := "ann+"+stamp+"@example.com", "bob+"+stamp+"@example.com"
	t0 := time.Now().UTC().Truncate(time.Millisecond)
	if created, err := repo.Put(ctx, ann, "elves", t0); err != nil || !created {
		t.Fatalf("first put: created = %v, %v", created, err)
	}
	if created, err := repo.Put(ctx, ann, "", t0.Add(time.Minute)); err != nil || created {
		t.Fatalf("repeat put: created = %v, %v", created, err)
	}
	if _, err := repo.Put(ctx, bob, "", t0.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	got, err := repo.Get(ctx, ann)
	if err != nil || got.Count != 2 || got.Note != "elves" || got.Status != Pending {
		t.Fatalf("get = %+v, %v", got, err)
	}
	list, err := repo.List(ctx, Pending)
	if err != nil {
		t.Fatal(err)
	}
	var order []string
	for _, r := range list {
		if r.Email == ann || r.Email == bob {
			order = append(order, r.Email)
		}
	}
	if len(order) != 2 || order[0] != bob {
		t.Fatalf("pending order = %v, want bob then ann", order)
	}
	if err := repo.Decide(ctx, ann, Dismissed, t0.Add(2*time.Hour)); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.Put(ctx, ann, "again", t0.Add(3*time.Hour)); err != nil {
		t.Fatal(err)
	}
	if got, _ := repo.Get(ctx, ann); got.Status != Dismissed || got.Note != "again" {
		t.Errorf("after a repeat of a dismissed request: %+v", got)
	}
	if err := repo.Decide(ctx, "nobody+"+stamp+"@example.com", Approved, t0); !errors.Is(err, ErrNotFound) {
		t.Errorf("decide of no request = %v, want ErrNotFound", err)
	}
	if _, err := repo.Get(ctx, "nobody+"+stamp+"@example.com"); !errors.Is(err, ErrNotFound) {
		t.Errorf("get of no request = %v, want ErrNotFound", err)
	}
}
