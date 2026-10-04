package deckreads

import (
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
)

// These tests need the local emulator, so CI skips them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/deckreads -count=1

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

// TestReadRoundTrip is D-1107 and D-1113: a user with no read has none,
// a read matches its deck and text for that user alone, a read stays
// after many newer reads, and each document holds the field of the TTL
// policy.
func TestReadRoundTrip(t *testing.T) {
	r := emulatorRepo(t)
	ctx := t.Context()
	uid := "u-reads-" + time.Now().Format("150405.000000")
	now := time.Now().UTC().Truncate(time.Millisecond)

	if ok, err := r.Has(ctx, uid, 42, "Deck\n4 Lightning Bolt\n", now); err != nil || ok {
		t.Fatalf("no read: ok = %v, err = %v", ok, err)
	}
	if err := r.Add(ctx, uid, 42, "Deck\n4 Lightning Bolt\n", now); err != nil {
		t.Fatal(err)
	}
	for i := range 25 {
		if err := r.Add(ctx, uid, int64(100+i), "Deck\n4 Shock\n", now.Add(time.Second)); err != nil {
			t.Fatal(err)
		}
	}
	for _, tc := range []struct {
		uid  string
		deck int64
		text string
		want bool
	}{
		{uid, 42, "Deck\n4 Lightning Bolt\n", true},
		{uid, 124, "Deck\n4 Shock\n", true},
		{uid, 42, "Deck\n4 Shock\n", false},
		{uid + "-other", 42, "Deck\n4 Lightning Bolt\n", false},
	} {
		ok, err := r.Has(ctx, tc.uid, tc.deck, tc.text, now.Add(time.Minute))
		if err != nil || ok != tc.want {
			t.Errorf("%s %d: ok = %v, err = %v, want %v", tc.uid, tc.deck, ok, err, tc.want)
		}
	}
	if ok, err := r.Has(ctx, uid, 42, "Deck\n4 Lightning Bolt\n", now.Add(Life)); err != nil || ok {
		t.Errorf("after Life: ok = %v, err = %v", ok, err)
	}
	snap, err := r.client.Collection(Collection).Doc(Key(uid, 42, "Deck\n4 Lightning Bolt\n")).Get(ctx)
	if err != nil {
		t.Fatal(err)
	}
	at, err := snap.DataAt(TTLField)
	if got, ok := at.(time.Time); err != nil || !ok || !got.Equal(now.Add(Life)) {
		t.Errorf("%s = %v (%v), want %v", TTLField, at, err, now.Add(Life))
	}
}
