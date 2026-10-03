package main

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"log/slog"
	"maps"
	"slices"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/stale"
)

// memDecks is the deck store of the test. Mark counts its writes, and
// it fails for the deck id in fail.
type memDecks struct {
	decks  map[string]*mtgv1.Deck
	writes int
	fail   string
}

func (m *memDecks) Scan(_ context.Context, fn func(uid string, d *mtgv1.Deck) error) error {
	for _, id := range slices.Sorted(maps.Keys(m.decks)) {
		d := m.decks[id]
		if err := fn("u1", &mtgv1.Deck{Id: d.GetId(), Format: d.GetFormat(), Cards: d.GetCards(),
			Stale: d.GetStale(), StaleOracleIds: d.GetStaleOracleIds(), RerunCase: d.GetRerunCase(), StaleReason: d.GetStaleReason()}); err != nil {
			return err
		}
	}
	return nil
}

func (m *memDecks) Mark(_ context.Context, _, id string, fn func(d *mtgv1.Deck) bool) (bool, error) {
	if id == m.fail {
		return false, errors.New("firestore down")
	}
	if !fn(m.decks[id]) {
		return false, nil
	}
	m.writes++
	return true, nil
}

// putVersion stores one finalized version with the given oracle lines.
func putVersion(t *testing.T, store cards.Store, version string, lines ...string) {
	t.Helper()
	ctx := context.Background()
	w, err := store.Create(ctx, version, "oracle_cards.jsonl.gz")
	if err != nil {
		t.Fatal(err)
	}
	gz := gzip.NewWriter(w)
	for _, l := range lines {
		if _, err := io.WriteString(gz, l+"\n"); err != nil {
			t.Fatal(err)
		}
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	if err := store.Finalize(ctx, version); err != nil {
		t.Fatal(err)
	}
}

// TestStalePass runs the pass on the newest diff, records its end in the
// marker, and opens no deck store on the next run (I-1).
func TestStalePass(t *testing.T) {
	ctx := context.Background()
	store := &cards.DirStore{Root: t.TempDir()}
	const v1, v2 = "20260930T090000", "20260930T100000"
	putVersion(t, store, v1, `{"oracle_id":"a","legalities":{"modern":"legal"}}`)
	putVersion(t, store, v2, `{"oracle_id":"a","legalities":{"modern":"banned"}}`)
	if err := store.WriteLegalityDiff(ctx, v2, cards.LegalityDiffRecord{SnapshotAsOf: "2026-09-30T10:00:00Z", ChangedCards: 1}); err != nil {
		t.Fatal(err)
	}
	mem := &memDecks{decks: map[string]*mtgv1.Deck{"d1": {
		Id:     "d1",
		Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN},
		Cards:  []*mtgv1.DeckCard{{OracleId: "a", Name: "Card A", Count: 4, Role: mtgv1.CardRole_CARD_ROLE_THREAT}},
	}}}
	opened := 0
	var sent []string
	open := func(context.Context) (stale.Store, notifyStale, func(), error) {
		opened++
		return mem, recordPush(&sent), func() {}, nil
	}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	at := time.Date(2026, 9, 30, 10, 5, 0, 0, time.UTC)
	now := func() time.Time { return at }

	if err := stalePass(ctx, store, open, log, now); err != nil {
		t.Fatal(err)
	}
	if d := mem.decks["d1"]; !d.GetStale() || d.GetRerunCase() != mtgv1.RerunCase_RERUN_CASE_REBUILD {
		t.Fatalf("d1 = stale %v case %v", d.GetStale(), d.GetRerunCase())
	}
	rec, ok, err := store.ReadLegalityDiff(ctx, v2)
	if err != nil || !ok {
		t.Fatalf("marker: ok %v err %v", ok, err)
	}
	if rec.StalePass != "2026-09-30T10:05:00Z" || rec.StaleDecks != 1 || rec.ChangedCards != 1 {
		t.Fatalf("marker = %+v", rec)
	}
	if !slices.Equal(sent, []string{"u1:d1"}) {
		t.Fatalf("push = %v, want u1:d1", sent)
	}

	if err := stalePass(ctx, store, open, log, now); err != nil {
		t.Fatal(err)
	}
	if opened != 1 || mem.writes != 1 || len(sent) != 1 {
		t.Fatalf("an ended pass must not run again: opened %d, writes %d, push %v", opened, mem.writes, sent)
	}
}

// recordPush writes each push as "uid:id,id" into sent.
func recordPush(sent *[]string) notifyStale {
	return func(_ context.Context, uid string, decks []*mtgv1.Deck) {
		ids := make([]string, 0, len(decks))
		for _, d := range decks {
			ids = append(ids, d.GetId())
		}
		*sent = append(*sent, uid+":"+strings.Join(ids, ","))
	}
}

// TestStalePassPushesAfterAFailure: the deck that a failed pass marked
// gets its push, and the marker keeps no end, so the pass runs again.
func TestStalePassPushesAfterAFailure(t *testing.T) {
	ctx := context.Background()
	store := &cards.DirStore{Root: t.TempDir()}
	const v1, v2 = "20260930T090000", "20260930T100000"
	putVersion(t, store, v1, `{"oracle_id":"a","legalities":{"modern":"legal"}}`)
	putVersion(t, store, v2, `{"oracle_id":"a","legalities":{"modern":"banned"}}`)
	if err := store.WriteLegalityDiff(ctx, v2, cards.LegalityDiffRecord{SnapshotAsOf: "2026-09-30T10:00:00Z", ChangedCards: 1}); err != nil {
		t.Fatal(err)
	}
	deck := func(id string) *mtgv1.Deck {
		return &mtgv1.Deck{Id: id, Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN},
			Cards: []*mtgv1.DeckCard{{OracleId: "a", Name: "Card A", Count: 4, Role: mtgv1.CardRole_CARD_ROLE_THREAT}}}
	}
	mem := &memDecks{decks: map[string]*mtgv1.Deck{"d1": deck("d1"), "d2": deck("d2")}, fail: "d2"}
	var sent []string
	open := func(context.Context) (stale.Store, notifyStale, func(), error) {
		return mem, recordPush(&sent), func() {}, nil
	}
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	if err := stalePass(ctx, store, open, log, time.Now); err == nil {
		t.Fatal("a failed mark must fail the pass")
	}
	if !slices.Equal(sent, []string{"u1:d1"}) {
		t.Fatalf("push = %v, want u1:d1", sent)
	}
	if rec, _, _ := store.ReadLegalityDiff(ctx, v2); rec.StalePass != "" {
		t.Fatalf("a failed pass must leave the marker open, %+v", rec)
	}
}

// TestStalePassNoDiff opens no deck store when no version carries a diff.
func TestStalePassNoDiff(t *testing.T) {
	store := &cards.DirStore{Root: t.TempDir()}
	putVersion(t, store, "20260930T090000", `{"oracle_id":"a","legalities":{"modern":"legal"}}`)
	open := func(context.Context) (stale.Store, notifyStale, func(), error) {
		t.Fatal("no diff, so no deck store")
		return nil, nil, nil, nil
	}
	if err := stalePass(context.Background(), store, open, slog.New(slog.NewTextHandler(io.Discard, nil)), time.Now); err != nil {
		t.Fatal(err)
	}
}
