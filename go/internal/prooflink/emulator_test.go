package prooflink

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
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/prooflink -count=1

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

func code(t *testing.T) string {
	t.Helper()
	c, err := NewCode()
	if err != nil {
		t.Fatal(err)
	}
	return c
}

// TestLinkRoundTrip is D-1081, D-1082, and D-1119: a link works one
// time, a second take names the uid alone, a send within the cooldown
// stores nothing, a new send removes the old link, and an expired link
// refuses.
func TestLinkRoundTrip(t *testing.T) {
	client := emulatorClient(t)
	ctx := context.Background()
	repo := NewRepo(client)
	uid := "u-" + time.Now().UTC().Format("150405.000000")
	t0 := time.Now().UTC().Truncate(time.Millisecond)
	link := func(at time.Time) Link {
		return Link{UID: uid, Email: "ann@example.com", CreatedAt: at, ExpiresAt: at.Add(Life)}
	}

	first := code(t)
	if err := repo.Put(ctx, first, link(t0)); err != nil {
		t.Fatalf("first put: %v", err)
	}
	if err := repo.Put(ctx, code(t), link(t0.Add(time.Second))); !errors.Is(err, ErrTooSoon) {
		t.Fatalf("a send within the cooldown: %v, want ErrTooSoon", err)
	}
	second := code(t)
	if err := repo.Put(ctx, second, link(t0.Add(Cooldown))); err != nil {
		t.Fatalf("second put: %v", err)
	}
	if _, err := repo.Take(ctx, first, t0.Add(Cooldown)); !errors.Is(err, ErrUnknown) {
		t.Fatalf("the replaced link: %v, want ErrUnknown", err)
	}
	got, err := repo.Take(ctx, second, t0.Add(2*Cooldown))
	if err != nil || got.UID != uid || got.Email != "ann@example.com" {
		t.Fatalf("take: %+v, %v", got, err)
	}
	again, err := repo.Take(ctx, second, t0.Add(3*Cooldown))
	if !errors.Is(err, ErrUsed) || again.UID != uid || again.Email != "" || !again.UsedAt.Equal(t0.Add(2*Cooldown)) {
		t.Fatalf("a second take: %+v, %v, want ErrUsed with the uid alone", again, err)
	}
	snap, err := client.Collection(Collection).Doc(Key(second)).Get(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := snap.DataAt("email"); err == nil {
		t.Fatalf("the used link keeps the email: %v", snap.Data())
	}
	if _, err := repo.Take(ctx, second, t0.Add(Cooldown+Life)); !errors.Is(err, ErrUsed) {
		t.Fatalf("a used link after its expiry: %v, want ErrUsed", err)
	}
	if _, err := repo.Take(ctx, second, t0.Add(Cooldown+Life)); !errors.Is(err, ErrUnknown) {
		t.Fatalf("the expired used link stays: %v, want ErrUnknown", err)
	}

	late := code(t)
	at := t0.Add(time.Hour)
	if err := repo.Put(ctx, late, link(at)); err != nil {
		t.Fatalf("late put: %v", err)
	}
	if _, err := repo.Take(ctx, late, at.Add(Life)); !errors.Is(err, ErrExpired) {
		t.Fatalf("an expired link: %v, want ErrExpired", err)
	}
	if _, err := repo.Take(ctx, late, at); !errors.Is(err, ErrUnknown) {
		t.Fatalf("the expired link stays: %v, want ErrUnknown", err)
	}
}
