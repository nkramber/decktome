package collections

import (
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/firestore"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// The collection store against the local Firestore emulator (PR-18).
// The fake of collectionsvc proves the service, and only the emulator
// proves the store: a field update, a read of one document, and the
// schema fallback all live in Firestore behaviour.
//
// These tests need the local emulator, so CI skips them:
//
//	firebase emulators:start --only firestore --project mtg-local
//	FIRESTORE_EMULATOR_HOST=127.0.0.1:8281 go test ./internal/collections -count=1
//
// CAUTION: a t.Cleanup can not delete from Firestore. Go cancels
// t.Context before a cleanup runs, so the delete fails and the next run
// reads the leftovers. Each test takes a fresh user id instead.

func emulatorRepo(t *testing.T) (*Repo, string) {
	t.Helper()
	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		t.Skip("set FIRESTORE_EMULATOR_HOST to run the store against the emulator")
	}
	client, err := firestore.NewClient(t.Context(), "mtg-local")
	if err != nil {
		t.Fatalf("firestore: %v", err)
	}
	t.Cleanup(func() { _ = client.Close() })
	return NewRepo(client), fmt.Sprintf("u-%s", t.Name())
}

func sampleCollection(name string) *mtgv1.Collection {
	entries := []*mtgv1.CollectionEntry{
		{ScryfallId: "p1", OracleId: "o-bolt", Name: "Lightning Bolt", SetCode: "lea", SetName: "Alpha",
			CollectorNumber: "1", Quantity: 4, Rarity: "common", Finish: mtgv1.Finish_FINISH_NORMAL},
		{ScryfallId: "p1", OracleId: "o-bolt", Name: "Lightning Bolt", SetCode: "lea", SetName: "Alpha",
			CollectorNumber: "1", Quantity: 1, Rarity: "common", Finish: mtgv1.Finish_FINISH_FOIL},
		{ScryfallId: "p2", OracleId: "o-jace", Name: "Jace, the Mind Sculptor", SetCode: "wwk", SetName: "Worldwake",
			CollectorNumber: "31", Quantity: 1, Rarity: "mythic", Finish: mtgv1.Finish_FINISH_NORMAL},
	}
	return &mtgv1.Collection{
		Name: name, Source: mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		ContentHash: "hash-" + name, Entries: entries, CardCount: CardCount(entries),
		Summary: Summarize(entries, nil),
	}
}

// TestEmulatorGetHeadReadsNoEntry is D-392. The binder head reads the
// stored summary, and the answer carries no row.
func TestEmulatorGetHeadReadsNoEntry(t *testing.T) {
	repo, uid := emulatorRepo(t)
	id, err := repo.Put(t.Context(), uid, sampleCollection("binder"))
	if err != nil {
		t.Fatal(err)
	}
	head, err := repo.GetHead(t.Context(), uid, id)
	if err != nil {
		t.Fatal(err)
	}
	if len(head.GetEntries()) != 0 {
		t.Errorf("the head carries %d entries, want none", len(head.GetEntries()))
	}
	sum := head.GetSummary()
	if sum.GetRowCount() != 3 || sum.GetUniqueCards() != 2 {
		t.Errorf("summary = %v, want 3 rows over 2 cards", sum)
	}
	if sum.GetByRarity()["common"] != 5 || sum.GetByRarity()["mythic"] != 1 {
		t.Errorf("by rarity = %v", sum.GetByRarity())
	}
	// A full read carries the same summary beside its rows.
	full, err := repo.Get(t.Context(), uid, id)
	if err != nil {
		t.Fatal(err)
	}
	if len(full.GetEntries()) != 3 || full.GetSummary().GetRowCount() != 3 {
		t.Errorf("a full read = %d entries, summary %v", len(full.GetEntries()), full.GetSummary())
	}
}

// TestEmulatorRenameKeepsEverythingElse is the rename of the collections
// list. It writes one field, so it rewrites no entry and it keeps the
// import date: a rename must not move the collection up a list ordered
// by date.
func TestEmulatorRenameKeepsEverythingElse(t *testing.T) {
	repo, uid := emulatorRepo(t)
	id, err := repo.Put(t.Context(), uid, sampleCollection("before"))
	if err != nil {
		t.Fatal(err)
	}
	was, err := repo.GetHead(t.Context(), uid, id)
	if err != nil {
		t.Fatal(err)
	}
	got, err := repo.Rename(t.Context(), uid, id, "after")
	if err != nil {
		t.Fatal(err)
	}
	if got.GetName() != "after" {
		t.Errorf("name = %q, want after", got.GetName())
	}
	if got.GetImportedAt().AsTime() != was.GetImportedAt().AsTime() {
		t.Errorf("the import date moved: %v to %v", was.GetImportedAt().AsTime(), got.GetImportedAt().AsTime())
	}
	// The entries survive a rename untouched.
	full, err := repo.Get(t.Context(), uid, id)
	if err != nil {
		t.Fatal(err)
	}
	if len(full.GetEntries()) != 3 || full.GetCardCount() != 6 {
		t.Errorf("after a rename: %d entries, %d cards", len(full.GetEntries()), full.GetCardCount())
	}
	if _, err := repo.Rename(t.Context(), uid, "missing", "x"); err == nil {
		t.Error("a rename of a missing collection must fail")
	}
}

// TestEmulatorReplaceKeepsTheIdAndTheName is D-393. Every deck and chat
// that names the collection still works, and a replacement of the cards
// does not rename the reader's binder.
func TestEmulatorReplaceKeepsTheIdAndTheName(t *testing.T) {
	repo, uid := emulatorRepo(t)
	id, err := repo.Put(t.Context(), uid, sampleCollection("my binder"))
	if err != nil {
		t.Fatal(err)
	}
	next := sampleCollection("a name the reader did not choose")
	next.Entries = next.Entries[:1]
	next.Entries[0].Quantity = 9
	next.CardCount = CardCount(next.Entries)
	next.Summary = Summarize(next.Entries, nil)
	next.ContentHash = "hash-two"

	head, err := repo.Replace(t.Context(), uid, id, next)
	if err != nil {
		t.Fatal(err)
	}
	if head.GetId() != id {
		t.Errorf("id = %q, want the collection it replaced (%q)", head.GetId(), id)
	}
	if head.GetName() != "my binder" {
		t.Errorf("name = %q, want the reader's own name", head.GetName())
	}
	if head.GetSummary().GetRowCount() != 1 {
		t.Errorf("summary row count = %d, want the new rows", head.GetSummary().GetRowCount())
	}
	full, err := repo.Get(t.Context(), uid, id)
	if err != nil {
		t.Fatal(err)
	}
	if len(full.GetEntries()) != 1 || full.GetEntries()[0].GetQuantity() != 9 {
		t.Errorf("the cards did not replace: %v", full.GetEntries())
	}
	if full.GetContentHash() != "hash-two" {
		t.Errorf("content hash = %q, want the new file's", full.GetContentHash())
	}
	// The owned counts follow the new entries, so a build reads them.
	counts, err := repo.OracleCounts(t.Context(), uid, id)
	if err != nil {
		t.Fatal(err)
	}
	if counts["o-bolt"] != 9 || counts["o-jace"] != 0 {
		t.Errorf("owned counts = %v, want 9 Bolt and no Jace", counts)
	}
}

// TestEmulatorHeadOfAnOldDocument is the schema fallback of D-392. A
// document stored before schema version 2 holds no summary, and the head
// computes one from its entries that one time.
func TestEmulatorHeadOfAnOldDocument(t *testing.T) {
	repo, uid := emulatorRepo(t)
	id, err := repo.Put(t.Context(), uid, sampleCollection("old"))
	if err != nil {
		t.Fatal(err)
	}
	// Strip the stored summary, which is what a version 1 document holds.
	if _, err := repo.doc(uid, id).Update(t.Context(), []firestore.Update{
		{Path: "summary_gz", Value: nil}, {Path: "schema_version", Value: 1},
	}); err != nil {
		t.Fatal(err)
	}
	head, err := repo.GetHead(t.Context(), uid, id)
	if err != nil {
		t.Fatal(err)
	}
	if head.GetSummary().GetRowCount() != 3 {
		t.Errorf("row count = %d, want the summary computed from the entries", head.GetSummary().GetRowCount())
	}
	if len(head.GetEntries()) != 0 {
		t.Error("the head still carries no entry")
	}
}
