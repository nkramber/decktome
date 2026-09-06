package allowlist

import (
	"context"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
)

// The list talks to Firestore, and the merge, the array union, and the
// missing document can not be proven without it. These tests need the
// local emulator, so CI skips them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/allowlist -count=1

func emulatorClient(t *testing.T) *firestore.Client {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the list against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return client
}

// TestListRoundTrip is D-420: no document is an empty list, an invite
// creates the document, a repeated invite changes nothing, and a removal
// takes the email off. The reads go through a fresh list each time, so
// the cache of a minute stays out of the way.
func TestListRoundTrip(t *testing.T) {
	client := emulatorClient(t)
	ctx := context.Background()
	if _, err := client.Collection(Collection).Doc(Doc).Delete(ctx); err != nil {
		t.Fatal(err)
	}
	if ok, err := FromFirestore(client).Allowed(ctx, "ann@example.com"); err != nil || ok {
		t.Fatalf("with no document: allowed = %v, %v, want false and no error", ok, err)
	}
	stamp := time.Now().UTC().Format("150405.000000")
	ann, bob := "Ann+"+stamp+"@Example.com", "bob+"+stamp+"@example.com"
	for _, email := range []string{ann, ann, bob} {
		if err := Add(ctx, client, email); err != nil {
			t.Fatal(err)
		}
	}
	snap, err := client.Collection(Collection).Doc(Doc).Get(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var doc struct {
		Emails []string `firestore:"emails"`
	}
	if err := snap.DataTo(&doc); err != nil {
		t.Fatal(err)
	}
	if len(doc.Emails) != 2 {
		t.Errorf("emails = %v, want the two invites once each", doc.Emails)
	}
	l := FromFirestore(client)
	if ok, err := l.Allowed(ctx, "ANN+"+stamp+"@example.com"); err != nil || !ok {
		t.Errorf("the invite in another case: %v, %v", ok, err)
	}
	if err := Remove(ctx, client, bob); err != nil {
		t.Fatal(err)
	}
	if ok, _ := FromFirestore(client).Allowed(ctx, bob); ok {
		t.Error("a removed email was allowed")
	}
	if err := Add(ctx, client, "no-at-sign"); err == nil {
		t.Error("an invite with no email address must fail")
	}
}
