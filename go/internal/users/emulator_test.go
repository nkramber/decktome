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
	if !rec.LastCreationAt.Equal(later) {
		t.Errorf("last_creation_at = %v, want the newest deck or chat %v", rec.LastCreationAt, later)
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
	// REV-034 of the review of 2026-09-24: a user imported 5 lists and deleted 2, and
	// the backfill counts the 3 that remain.
	for range 5 {
		if err := r.Note(ctx, uid, "", DecksImported, at); err != nil {
			t.Fatalf("note: %v", err)
		}
	}
	// The backfill counted 1, and the record already holds 3.
	err := r.Seed(ctx, uid, "reader@example.com", map[Counter]int64{
		DecksCreated: 1, CollectionsUpload: 5, DecksImported: 3,
	}, at.Add(-time.Hour), at, time.Time{})
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
	if rec.DecksImported != 5 {
		t.Errorf("imports = %d, and the seed lowered a live count", rec.DecksImported)
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

// TestDeactivateKeepsTheRecordAndTheFirstTime is D-941: a close writes
// the mark alone, a second close keeps the first time, and a later note
// keeps the mark.
func TestDeactivateKeepsTheRecordAndTheFirstTime(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	made := time.Date(2026, 9, 9, 10, 0, 0, 0, time.UTC)
	if err := r.Note(ctx, uid, "reader@example.com", DecksCreated, made); err != nil {
		t.Fatal(err)
	}
	if closed, err := r.Deactivated(ctx, uid); err != nil || closed {
		t.Fatalf("an open account reads closed %v, %v", closed, err)
	}
	first := made.Add(time.Hour)
	if at, err := r.Deactivate(ctx, uid, first); err != nil || !at.Equal(first) {
		t.Fatalf("Deactivate = %v, %v", at, err)
	}
	if at, err := r.Deactivate(ctx, uid, first.Add(time.Hour)); err != nil || !at.Equal(first) {
		t.Errorf("a second close moved the time to %v, %v", at, err)
	}
	if err := r.Note(ctx, uid, "reader@example.com", DeckRevisions, first.Add(2*time.Hour)); err != nil {
		t.Fatal(err)
	}
	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatal(err)
	}
	if rec.DeactivatedAt == nil || !rec.DeactivatedAt.Equal(first) || rec.DecksCreated != 1 || rec.Email != "reader@example.com" {
		t.Errorf("the record after the close: %+v", rec)
	}
	if closed, err := r.Deactivated(ctx, "u-no-record-"+uid); err != nil || closed {
		t.Errorf("a uid with no record reads closed %v, %v", closed, err)
	}
}

// TestTouchMakesTheRecordOfAUserWhoMadeNothing is D-1092: a verified
// call alone writes the record, with zero counts and no creation.
func TestTouchMakesTheRecordOfAUserWhoMadeNothing(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	first := time.Date(2026, 10, 3, 19, 25, 0, 0, time.UTC)
	later := first.Add(VisitTTL)

	if err := r.Touch(ctx, uid, "reader@example.com", first); err != nil {
		t.Fatalf("first touch: %v", err)
	}
	if err := r.Touch(ctx, uid, "", later); err != nil {
		t.Fatalf("second touch: %v", err)
	}
	// A late write of an earlier call never moves the time back.
	if err := r.Touch(ctx, uid, "", first); err != nil {
		t.Fatalf("late touch: %v", err)
	}
	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !rec.CreatedAt.Equal(first) || !rec.LastSeenAt.Equal(later) {
		t.Errorf("created_at %v and last_seen_at %v, want %v and %v", rec.CreatedAt, rec.LastSeenAt, first, later)
	}
	if !rec.LastCreationAt.IsZero() {
		t.Errorf("last_creation_at = %v, and a call made nothing", rec.LastCreationAt)
	}
	if rec.Email != "reader@example.com" || rec.Schema != schemaVersion {
		t.Errorf("email %q and schema %d", rec.Email, rec.Schema)
	}
	if rec.DecksCreated != 0 || rec.SessionsStarted != 0 {
		t.Errorf("counts %d/%d, want zero", rec.DecksCreated, rec.SessionsStarted)
	}
}

// TestAnUploadOrAVerdictMovesOnlyTheLastSeenTime is D-1093: the decks
// and the chats move last_creation_at, and every creation moves
// last_seen_at.
func TestAnUploadOrAVerdictMovesOnlyTheLastSeenTime(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	deck := time.Date(2026, 10, 3, 10, 0, 0, 0, time.UTC)

	if err := r.Note(ctx, uid, "reader@example.com", DecksImported, deck); err != nil {
		t.Fatalf("import: %v", err)
	}
	for i, c := range []Counter{CollectionsUpload, FeedbackUp, FeedbackDown} {
		if err := r.Note(ctx, uid, "reader@example.com", c, deck.Add(time.Duration(i+1)*time.Hour)); err != nil {
			t.Fatalf("note %s: %v", c, err)
		}
	}
	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !rec.LastCreationAt.Equal(deck) {
		t.Errorf("last_creation_at = %v, want the import %v", rec.LastCreationAt, deck)
	}
	if want := deck.Add(3 * time.Hour); !rec.LastSeenAt.Equal(want) {
		t.Errorf("last_seen_at = %v, want the newest verdict %v", rec.LastSeenAt, want)
	}
}

// TestTouchMovesARecordOfVersionOne is D-1094: in version 1,
// last_seen_at held the newest creation, so the first call copies it to
// last_creation_at before it moves last_seen_at.
func TestTouchMovesARecordOfVersionOne(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	made := time.Date(2026, 9, 20, 10, 0, 0, 0, time.UTC)
	visit := time.Date(2026, 10, 3, 19, 25, 0, 0, time.UTC)

	if _, err := r.doc(uid).Set(ctx, map[string]any{
		"schema": int64(1), "email": "reader@example.com",
		"created_at": made, "last_seen_at": made,
		string(DecksCreated): int64(3),
	}); err != nil {
		t.Fatalf("seed a version 1 record: %v", err)
	}
	if err := r.Touch(ctx, uid, "reader@example.com", visit); err != nil {
		t.Fatalf("touch: %v", err)
	}
	if err := r.Touch(ctx, uid, "reader@example.com", visit.Add(VisitTTL)); err != nil {
		t.Fatalf("second touch: %v", err)
	}
	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !rec.LastCreationAt.Equal(made) {
		t.Errorf("last_creation_at = %v, want the version 1 last_seen_at %v", rec.LastCreationAt, made)
	}
	if want := visit.Add(VisitTTL); !rec.LastSeenAt.Equal(want) {
		t.Errorf("last_seen_at = %v, want %v", rec.LastSeenAt, want)
	}
	if rec.Schema != schemaVersion || rec.DecksCreated != 3 || !rec.CreatedAt.Equal(made) {
		t.Errorf("schema %d, decks %d, created %v", rec.Schema, rec.DecksCreated, rec.CreatedAt)
	}
}

// TestSeedSetsTheNewestDeckOrChat is D-1093 in the backfill: a seed
// writes the newest deck or chat it found, and never moves it back.
func TestSeedSetsTheNewestDeckOrChat(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	made := time.Date(2026, 9, 20, 10, 0, 0, 0, time.UTC)
	upload := made.Add(time.Hour)

	fresh := uniqueUID()
	if err := r.Seed(ctx, fresh, "", map[Counter]int64{DecksCreated: 2, CollectionsUpload: 1}, made, upload, made); err != nil {
		t.Fatalf("Seed: %v", err)
	}
	rec, err := r.Get(ctx, fresh)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !rec.LastCreationAt.Equal(made) || !rec.LastSeenAt.Equal(upload) {
		t.Errorf("last_creation_at %v and last_seen_at %v, want %v and %v", rec.LastCreationAt, rec.LastSeenAt, made, upload)
	}

	// A record of version 1 holds a newer creation than the backfill
	// found, so the copy of D-1094 wins.
	old := uniqueUID()
	newer := made.Add(48 * time.Hour)
	if _, err := r.doc(old).Set(ctx, map[string]any{
		"schema": int64(1), "created_at": made, "last_seen_at": newer,
	}); err != nil {
		t.Fatalf("seed a version 1 record: %v", err)
	}
	if err := r.Seed(ctx, old, "", map[Counter]int64{DecksCreated: 1}, made, made, made); err != nil {
		t.Fatalf("Seed: %v", err)
	}
	if rec, err = r.Get(ctx, old); err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !rec.LastCreationAt.Equal(newer) {
		t.Errorf("last_creation_at = %v, want the copied %v", rec.LastCreationAt, newer)
	}
}

// TestALateNoteNeverMovesATimeBack is P2-1 of the review of #276: two
// creations can commit in the other order of their times, and each time
// keeps the newest one (D-1093).
func TestALateNoteNeverMovesATimeBack(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	newer := time.Date(2026, 10, 3, 20, 0, 0, 0, time.UTC)
	older := newer.Add(-time.Minute)

	if err := r.Note(ctx, uid, "", SessionsStarted, newer); err != nil {
		t.Fatalf("newer note: %v", err)
	}
	for _, c := range []Counter{DecksCreated, CollectionsUpload} {
		if err := r.Note(ctx, uid, "", c, older); err != nil {
			t.Fatalf("older note %s: %v", c, err)
		}
	}
	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !rec.LastSeenAt.Equal(newer) || !rec.LastCreationAt.Equal(newer) {
		t.Errorf("last_seen_at %v and last_creation_at %v, want both %v", rec.LastSeenAt, rec.LastCreationAt, newer)
	}
	if !rec.CreatedAt.Equal(newer) {
		t.Errorf("created_at = %v, want the first write %v", rec.CreatedAt, newer)
	}
	if rec.SessionsStarted != 1 || rec.DecksCreated != 1 || rec.CollectionsUpload != 1 {
		t.Errorf("counts %d/%d/%d, want 1/1/1", rec.SessionsStarted, rec.DecksCreated, rec.CollectionsUpload)
	}
}

// TestNoteMovesARecordOfVersionOne is D-1094 on the creation path: an
// upload on a version 1 record keeps the old newest creation.
func TestNoteMovesARecordOfVersionOne(t *testing.T) {
	r := emulatorRepo(t)
	ctx := context.Background()
	uid := uniqueUID()
	made := time.Date(2026, 9, 20, 10, 0, 0, 0, time.UTC)
	upload := time.Date(2026, 10, 3, 19, 25, 0, 0, time.UTC)

	if _, err := r.doc(uid).Set(ctx, map[string]any{
		"schema": int64(1), "created_at": made, "last_seen_at": made,
	}); err != nil {
		t.Fatalf("seed a version 1 record: %v", err)
	}
	if err := r.Note(ctx, uid, "", CollectionsUpload, upload); err != nil {
		t.Fatalf("note: %v", err)
	}
	rec, err := r.Get(ctx, uid)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !rec.LastCreationAt.Equal(made) || !rec.LastSeenAt.Equal(upload) {
		t.Errorf("last_creation_at %v and last_seen_at %v, want %v and %v", rec.LastCreationAt, rec.LastSeenAt, made, upload)
	}
	if rec.Schema != schemaVersion || rec.CollectionsUpload != 1 {
		t.Errorf("schema %d, uploads %d", rec.Schema, rec.CollectionsUpload)
	}
}
