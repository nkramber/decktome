package feedback

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
)

// The store talks to Firestore, and the write and the read back can not
// be proven without it. These tests need the local emulator, so CI skips
// them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/feedback -count=1

func emulatorRepo(t *testing.T) *Repo {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the store against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return NewRepo(client)
}

// TestFeedbackRoundTrip is D-558: a verdict reads back whole under its
// user, with the reasons, the prompt versions, and the time. Another
// user's read of the same id is ErrNotFound, as is an unknown id.
func TestFeedbackRoundTrip(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := "u-" + time.Now().UTC().Format("150405.000000")
	at := time.Date(2026, 9, 6, 15, 30, 0, 0, time.UTC)
	item := Item{
		Kind:      "card",
		Verdict:   "down",
		DeckID:    "d1",
		OracleID:  "o-sol",
		Reasons:   []string{"off_theme", "wrong_power"},
		Text:      "Sol Ring is banned at my table.",
		Prompts:   map[string]int64{"questions": 18, "generate": 12},
		CreatedAt: at,
	}
	id, err := r.Add(ctx, uid, item)
	if err != nil || id == "" {
		t.Fatalf("Add = %q, %v", id, err)
	}
	got, err := r.Get(ctx, uid, id)
	if err != nil {
		t.Fatal(err)
	}
	if got.Kind != "card" || got.Verdict != "down" || got.DeckID != "d1" || got.OracleID != "o-sol" || got.Text != item.Text {
		t.Errorf("Get = %+v, want the item back", got)
	}
	if len(got.Reasons) != 2 || got.Reasons[0] != "off_theme" || got.Reasons[1] != "wrong_power" {
		t.Errorf("reasons = %v", got.Reasons)
	}
	if got.Prompts["questions"] != 18 || got.Prompts["generate"] != 12 {
		t.Errorf("prompts = %v", got.Prompts)
	}
	if !got.CreatedAt.Equal(at) {
		t.Errorf("created_at = %v, want %v", got.CreatedAt, at)
	}
	if _, err := r.Get(ctx, uid+"-other", id); !errors.Is(err, ErrNotFound) {
		t.Errorf("another user's read = %v, want ErrNotFound", err)
	}
	if _, err := r.Get(ctx, uid, "nope"); !errors.Is(err, ErrNotFound) {
		t.Errorf("an unknown id = %v, want ErrNotFound", err)
	}
	// A thumbs up carries no reason and no text, and it reads back so.
	up, err := r.Add(ctx, uid, Item{Kind: "deck", Verdict: "up", DeckID: "d1", Prompts: item.Prompts, CreatedAt: at})
	if err != nil {
		t.Fatal(err)
	}
	if got, err := r.Get(ctx, uid, up); err != nil || len(got.Reasons) != 0 || got.Text != "" {
		t.Errorf("the thumbs up = %+v, %v", got, err)
	}
}
