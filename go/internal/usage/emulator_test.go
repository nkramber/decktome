package usage

import (
	"context"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
)

// The ledger talks to Firestore, and the increment and the merge can not
// be proven without it. These tests need the local emulator, so CI skips
// them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/usage -count=1

func emulatorRepo(t *testing.T) *Repo {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the ledger against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return NewRepo(client)
}

// TestLedgerSumsTheTurns is D-421: two turns sum, a month with no
// document reads zero, and another user's month stays apart.
func TestLedgerSumsTheTurns(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := "u-" + time.Now().UTC().Format("150405.000000")
	if spent, err := r.Spent(ctx, uid, "2026-09"); err != nil || spent != 0 {
		t.Fatalf("empty month = %v, %v", spent, err)
	}
	at := time.Date(2026, 9, 6, 12, 0, 0, 0, time.UTC)
	if err := r.Add(ctx, uid, "2026-09", 0.10, 4, at); err != nil {
		t.Fatal(err)
	}
	if err := r.Add(ctx, uid, "2026-09", 0.25, 6, at.Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	spent, err := r.Spent(ctx, uid, "2026-09")
	if err != nil || spent < 0.349 || spent > 0.351 {
		t.Errorf("spent = %v, %v, want 0.35", spent, err)
	}
	if other, _ := r.Spent(ctx, uid+"-other", "2026-09"); other != 0 {
		t.Errorf("another user's month = %v, want 0", other)
	}
	if next, _ := r.Spent(ctx, uid, "2026-10"); next != 0 {
		t.Errorf("the next month = %v, want 0", next)
	}
}
