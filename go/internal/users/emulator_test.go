package users

import (
	"context"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
)

// The record talks to Firestore, and the write and the read back can not
// be proven without it. These tests need the local emulator, so CI skips
// them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/users -count=1

func emulatorRepo(t *testing.T) *Repo {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the record against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return NewRepo(client)
}

func uniqueUID() string { return "u-" + time.Now().UTC().Format("150405.000000000") }

// TestNoteCountsAndKeepsTheFirstDate is the shape of D-638. The counters
// rise, the creation date holds its first value, and the last seen time
// follows the newest creation.
func TestNoteCountsAndKeepsTheFirstDate(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	first := time.Date(2026, 9, 9, 10, 0, 0, 0, time.UTC)
	later := first.Add(2 * time.Hour)

	if err := r.Note(ctx, uid, "reader@example.com", DecksCreated, first); err != nil {
		t.Fatalf("first note: %v", err)
	}
	for _, c := range []Counter{DecksCreated, DeckRevisions, CollectionsUpload, SessionsStarted, FeedbackDown} {
		if err := r.Note(ctx, uid, "reader@example.com", c, later); err != nil {
			t.Fatalf("note %s: %v", c, err)
		}
	}

	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if rec.DecksCreated != 2 {
		t.Errorf("decks = %d, want 2", rec.DecksCreated)
	}
	if rec.DeckRevisions != 1 || rec.CollectionsUpload != 1 || rec.SessionsStarted != 1 {
		t.Errorf("counts = %d/%d/%d, want 1/1/1", rec.DeckRevisions, rec.CollectionsUpload, rec.SessionsStarted)
	}
	if rec.FeedbackDown != 1 || rec.FeedbackUp != 0 {
		t.Errorf("feedback = %d down and %d up, want 1 and 0", rec.FeedbackDown, rec.FeedbackUp)
	}
	// The creation date is the first write, and no later one moves it.
	if !rec.CreatedAt.Equal(first) {
		t.Errorf("created_at = %v, want the first write %v", rec.CreatedAt, first)
	}
	if !rec.LastSeenAt.Equal(later) {
		t.Errorf("last_seen_at = %v, want the newest creation %v", rec.LastSeenAt, later)
	}
	if rec.Email != "reader@example.com" {
		t.Errorf("email = %q", rec.Email)
	}
	if rec.Schema != schemaVersion {
		t.Errorf("schema = %d, want %d", rec.Schema, schemaVersion)
	}
}

// TestAnEmptyEmailKeepsTheOne holds a rule the sign-in path needs: a
// turn with no verified address must never erase the address on record.
func TestAnEmptyEmailKeepsTheOne(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	at := time.Date(2026, 9, 9, 10, 0, 0, 0, time.UTC)

	if err := r.Note(ctx, uid, "reader@example.com", DecksCreated, at); err != nil {
		t.Fatalf("note: %v", err)
	}
	if err := r.Note(ctx, uid, "", DecksCreated, at.Add(time.Hour)); err != nil {
		t.Fatalf("second note: %v", err)
	}
	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if rec.Email != "reader@example.com" {
		t.Errorf("email = %q, and an empty one erased it", rec.Email)
	}
}

// TestSeedNeverLowersACount holds the backfill rule of D-638. A creation
// that lands between the read and the write must survive the seed.
func TestSeedNeverLowersACount(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	at := time.Date(2026, 9, 9, 10, 0, 0, 0, time.UTC)

	for range 3 {
		if err := r.Note(ctx, uid, "", DecksCreated, at); err != nil {
			t.Fatalf("note: %v", err)
		}
	}
	// The backfill counted 1, and the record already holds 3.
	err := r.Seed(ctx, uid, "reader@example.com", map[Counter]int64{
		DecksCreated: 1, CollectionsUpload: 5,
	}, at.Add(-time.Hour), at)
	if err != nil {
		t.Fatalf("Seed: %v", err)
	}
	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if rec.DecksCreated != 3 {
		t.Errorf("decks = %d, and the seed lowered a live count", rec.DecksCreated)
	}
	if rec.CollectionsUpload != 5 {
		t.Errorf("collections = %d, want the seeded 5", rec.CollectionsUpload)
	}
	// The record already had a creation date, so the seed leaves it.
	if !rec.CreatedAt.Equal(at) {
		t.Errorf("created_at = %v, want the first write %v", rec.CreatedAt, at)
	}
}

// TestNoteRefusesAFieldThatIsNoCounter keeps a typo out of the record.
func TestNoteRefusesAFieldThatIsNoCounter(t *testing.T) {
	r := emulatorRepo(t)
	if err := r.Note(context.Background(), uniqueUID(), "", Counter("email"), time.Now()); err == nil {
		t.Fatal("a field that is no counter wrote to the record")
	}
}
