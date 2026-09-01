package collectionsvc

import (
	"context"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// stocked is a server whose store holds one collection of n entries.
func stocked(t *testing.T, n int) (*Server, *fakeRepo) {
	t.Helper()
	var entries []*mtgv1.CollectionEntry
	for i := 0; i < n; i++ {
		entries = append(entries, &mtgv1.CollectionEntry{
			ScryfallId: "p" + string(rune('a'+i%26)) + string(rune('a'+i/26)),
			OracleId:   "o" + string(rune('a'+i%26)),
			Name:       "Card", SetCode: "blb", SetName: "Bloomburrow",
			Rarity: "common", Quantity: 1,
		})
	}
	repo := newFakeRepo()
	repo.stored["col-1"] = &mtgv1.Collection{
		Id: "col-1", Name: "Binder", Entries: entries,
		CardCount: int32(n), ContentHash: "hash-1",
	}
	return newServer(repo, cards.NewIndex(nil, nil, nil, time.Time{})), repo
}

// TestGetCollectionOmitsTheEntries is D-392. The binder head reads the
// summary and no entry, so it moves no megabyte.
func TestGetCollectionOmitsTheEntries(t *testing.T) {
	s, _ := stocked(t, 250)
	res, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{
		CollectionId: "col-1", EntriesOmitted: true,
	}))
	if err != nil {
		t.Fatal(err)
	}
	col := res.Msg.GetCollection()
	if len(col.GetEntries()) != 0 {
		t.Errorf("the head carries %d entries, want none", len(col.GetEntries()))
	}
	if col.GetSummary().GetRowCount() != 250 {
		t.Errorf("row count = %d, want 250", col.GetSummary().GetRowCount())
	}
	if res.Msg.GetNextPageToken() != "" {
		t.Error("a head read returns no page token")
	}
}

// TestGetCollectionPagesTheEntries is D-392. The grid reads a page at a
// time, and the last page returns no token.
func TestGetCollectionPagesTheEntries(t *testing.T) {
	s, _ := stocked(t, 250)
	first, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{
		CollectionId: "col-1", PageSize: 100,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if n := len(first.Msg.GetCollection().GetEntries()); n != 100 {
		t.Fatalf("first page holds %d entries, want 100", n)
	}
	if first.Msg.GetNextPageToken() == "" {
		t.Fatal("the first page of 250 returns a token")
	}
	// Walk to the end, and count every entry exactly once.
	seen, token, pages := 0, first.Msg.GetNextPageToken(), 1
	seen += len(first.Msg.GetCollection().GetEntries())
	for token != "" && pages < 10 {
		res, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{
			CollectionId: "col-1", PageSize: 100, PageToken: token,
		}))
		if err != nil {
			t.Fatal(err)
		}
		seen += len(res.Msg.GetCollection().GetEntries())
		token = res.Msg.GetNextPageToken()
		pages++
	}
	if seen != 250 || pages != 3 {
		t.Errorf("read %d entries over %d pages, want 250 over 3", seen, pages)
	}
}

// TestGetCollectionBoundsThePageSize: zero takes the default, and a
// request over the cap takes the cap.
func TestGetCollectionBoundsThePageSize(t *testing.T) {
	s, _ := stocked(t, MaxPageSize+50)
	for _, tc := range []struct {
		name string
		size int32
		want int
	}{
		{"zero takes the default", 0, DefaultPageSize},
		{"a negative takes the default", -5, DefaultPageSize},
		{"over the cap takes the cap", MaxPageSize + 500, MaxPageSize},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{
				CollectionId: "col-1", PageSize: tc.size,
			}))
			if err != nil {
				t.Fatal(err)
			}
			if n := len(res.Msg.GetCollection().GetEntries()); n != tc.want {
				t.Errorf("page holds %d entries, want %d", n, tc.want)
			}
		})
	}
}

// TestGetCollectionRefusesABadPageToken: a token of another collection,
// or one a reader typed, names a row this collection does not hold.
func TestGetCollectionRefusesABadPageToken(t *testing.T) {
	s, _ := stocked(t, 10)
	for _, token := range []string{"nonsense", "-1", "500"} {
		_, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{
			CollectionId: "col-1", PageToken: token,
		}))
		if connect.CodeOf(err) != connect.CodeInvalidArgument {
			t.Errorf("token %q gave %v, want InvalidArgument", token, connect.CodeOf(err))
		}
	}
}

// TestUpdateCollectionRenames is the rename of the collections list.
func TestUpdateCollectionRenames(t *testing.T) {
	s, repo := stocked(t, 3)
	res, err := s.UpdateCollection(context.Background(), connect.NewRequest(&mtgv1.UpdateCollectionRequest{
		CollectionId: "col-1", Name: "  My binder  ",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if got := res.Msg.GetCollection().GetName(); got != "My binder" {
		t.Errorf("name = %q, want the trimmed name", got)
	}
	if got := repo.stored["col-1"].GetName(); got != "My binder" {
		t.Errorf("the store holds %q", got)
	}
	// The answer is a head: a rename needs no entry.
	if len(res.Msg.GetCollection().GetEntries()) != 0 {
		t.Error("a rename answered with entries")
	}
}

// TestUpdateCollectionCodes: an empty or over-long name is refused, and
// a missing collection is NotFound.
func TestUpdateCollectionCodes(t *testing.T) {
	s, _ := stocked(t, 1)
	for _, tc := range []struct {
		name, id, newName string
		want              connect.Code
	}{
		{"empty name", "col-1", "   ", connect.CodeInvalidArgument},
		{"long name", "col-1", strings.Repeat("x", MaxNameBytes+1), connect.CodeInvalidArgument},
		{"empty id", "", "ok", connect.CodeInvalidArgument},
		{"path id", "../other", "ok", connect.CodeInvalidArgument},
		{"missing", "missing", "ok", connect.CodeNotFound},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.UpdateCollection(context.Background(), connect.NewRequest(&mtgv1.UpdateCollectionRequest{
				CollectionId: tc.id, Name: tc.newName,
			}))
			if connect.CodeOf(err) != tc.want {
				t.Errorf("code = %v, want %v", connect.CodeOf(err), tc.want)
			}
		})
	}
}

// oneRowCSV is a ManaBox file of one row, so a diff has a known shape.
func oneRowCSV(qty string) string {
	return "Name,Set code,Collector number,Quantity,Scryfall ID,Foil,Condition,Language\n" +
		"Pawpatch Recruit,BLB,187," + qty + "," + scryfallID + ",foil,near_mint,en\n"
}

// importedServer imports one file and returns the server and the id.
func importedServer(t *testing.T, csv string) (*Server, string) {
	t.Helper()
	s := newServer(newFakeRepo(), testIndex())
	res, err := s.ImportCollection(context.Background(),
		importReq("Binder", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, csv))
	if err != nil {
		t.Fatal(err)
	}
	return s, res.Msg.GetCollection().GetId()
}

// TestTheSameFileDiffsEmpty is the PR-18 gate line.
func TestTheSameFileDiffsEmpty(t *testing.T) {
	s, id := importedServer(t, oneRowCSV("3"))
	res, err := s.DiffCollections(context.Background(), connect.NewRequest(&mtgv1.DiffCollectionsRequest{
		CollectionId: id,
		Source:       mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		Content:      []byte(oneRowCSV("3")),
	}))
	if err != nil {
		t.Fatal(err)
	}
	d := res.Msg.GetDiff()
	if !d.GetIdentical() {
		t.Fatalf("the same file is not identical: %+v", d)
	}
	if len(d.GetAdded())+len(d.GetRemoved())+len(d.GetChanged()) != 0 {
		t.Errorf("diff = %+v, want nothing", d)
	}
}

// TestADiffReadsAQuantityChange is D-393: a row both sides hold with
// another quantity reads as changed, not as an add beside a remove.
func TestADiffReadsAQuantityChange(t *testing.T) {
	s, id := importedServer(t, oneRowCSV("3"))
	res, err := s.DiffCollections(context.Background(), connect.NewRequest(&mtgv1.DiffCollectionsRequest{
		CollectionId: id,
		Source:       mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		Content:      []byte(oneRowCSV("5")),
	}))
	if err != nil {
		t.Fatal(err)
	}
	d := res.Msg.GetDiff()
	if len(d.GetChanged()) != 1 {
		t.Fatalf("changed = %+v, want one row", d.GetChanged())
	}
	if d.GetChanged()[0].GetFrom() != 3 || d.GetChanged()[0].GetTo() != 5 {
		t.Errorf("change = %d to %d, want 3 to 5", d.GetChanged()[0].GetFrom(), d.GetChanged()[0].GetTo())
	}
	if len(d.GetAdded()) != 0 || len(d.GetRemoved()) != 0 {
		t.Error("a quantity change is not an add beside a remove")
	}
	// The diff stores nothing: the stored collection still holds three.
	got, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{CollectionId: id}))
	if err != nil {
		t.Fatal(err)
	}
	if q := got.Msg.GetCollection().GetEntries()[0].GetQuantity(); q != 3 {
		t.Errorf("the store holds %d after a diff, want 3", q)
	}
}

// TestReplaceKeepsTheIdAndTheName is D-393. Every deck and chat that
// names the collection still works, and a replacement of the cards does
// not rename the reader's binder.
func TestReplaceKeepsTheIdAndTheName(t *testing.T) {
	s, id := importedServer(t, oneRowCSV("3"))
	res, err := s.ImportCollection(context.Background(), connect.NewRequest(&mtgv1.ImportCollectionRequest{
		Name:                "A name the reader did not choose",
		Source:              mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		Content:             []byte(oneRowCSV("5")),
		ReplaceCollectionId: id,
	}))
	if err != nil {
		t.Fatal(err)
	}
	col := res.Msg.GetCollection()
	if col.GetId() != id {
		t.Errorf("id = %q, want the collection it replaced (%q)", col.GetId(), id)
	}
	if col.GetName() != "Binder" {
		t.Errorf("name = %q, want the reader's own name", col.GetName())
	}
	// The cards are the new ones.
	got, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{CollectionId: id}))
	if err != nil {
		t.Fatal(err)
	}
	if q := got.Msg.GetCollection().GetEntries()[0].GetQuantity(); q != 5 {
		t.Errorf("the store holds %d after a replace, want 5", q)
	}
}

// TestDiffCodes: the same id and upload rules the import applies.
func TestDiffCodes(t *testing.T) {
	s, id := importedServer(t, oneRowCSV("3"))
	for _, tc := range []struct {
		name    string
		id      string
		content string
		source  mtgv1.ImportSource
		want    connect.Code
	}{
		{"empty id", "", oneRowCSV("1"), mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, connect.CodeInvalidArgument},
		{"path id", "../other", oneRowCSV("1"), mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, connect.CodeInvalidArgument},
		{"empty body", id, "", mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, connect.CodeInvalidArgument},
		{"no source", id, oneRowCSV("1"), mtgv1.ImportSource_IMPORT_SOURCE_UNSPECIFIED, connect.CodeInvalidArgument},
		{"missing collection", "missing", oneRowCSV("1"), mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV, connect.CodeNotFound},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.DiffCollections(context.Background(), connect.NewRequest(&mtgv1.DiffCollectionsRequest{
				CollectionId: tc.id, Source: tc.source, Content: []byte(tc.content),
			}))
			if connect.CodeOf(err) != tc.want {
				t.Errorf("code = %v, want %v", connect.CodeOf(err), tc.want)
			}
		})
	}
}

// TestGetCollectionFillsTheDisplayFields is D-396. No document stores
// the colors, the card types, or the price. The index of the day fills
// a page, so the price is never stale.
func TestGetCollectionFillsTheDisplayFields(t *testing.T) {
	c := &mtgv1.Card{
		OracleId:  "o-bolt",
		Name:      "Lightning Bolt",
		Colors:    []mtgv1.Color{mtgv1.Color_COLOR_R},
		CardTypes: []string{"Instant"},
		PriceUsd:  1.25,
	}
	printings := []cards.Printing{
		{ScryfallID: "p-1", OracleID: "o-bolt", Name: c.Name, SetCode: "lea", CollectorNumber: "161", Layout: "normal", PriceUSD: 42.5},
		{ScryfallID: "p-2", OracleID: "o-bolt", Name: c.Name, SetCode: "m10", CollectorNumber: "146", Layout: "normal"},
	}
	repo := newFakeRepo()
	repo.stored["col-1"] = &mtgv1.Collection{
		Id: "col-1", Name: "Binder", CardCount: 3, ContentHash: "hash-1",
		Entries: []*mtgv1.CollectionEntry{
			{ScryfallId: "p-1", OracleId: "o-bolt", Name: c.Name, SetCode: "lea", Quantity: 1},
			{ScryfallId: "p-2", OracleId: "o-bolt", Name: c.Name, SetCode: "m10", Quantity: 1},
			{ScryfallId: "p-9", OracleId: "o-unknown", Name: "Nothing", SetCode: "zzz", Quantity: 1},
		},
	}
	s := newServer(repo, cards.NewIndex([]*mtgv1.Card{c}, printings, nil, time.Now()))

	res, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{CollectionId: "col-1"}))
	if err != nil {
		t.Fatal(err)
	}
	got := res.Msg.GetCollection().GetEntries()
	if len(got) != 3 {
		t.Fatalf("read %d entries, want 3", len(got))
	}
	if len(got[0].GetColors()) != 1 || got[0].GetColors()[0] != mtgv1.Color_COLOR_R {
		t.Errorf("colors = %v, want red", got[0].GetColors())
	}
	if len(got[0].GetCardTypes()) != 1 || got[0].GetCardTypes()[0] != "Instant" {
		t.Errorf("card types = %v, want Instant", got[0].GetCardTypes())
	}
	// The reader owns this printing, so its price is the one that counts.
	if got[0].GetPriceUsd() != 42.5 {
		t.Errorf("price of the Alpha printing = %v, want 42.5", got[0].GetPriceUsd())
	}
	// A printing with no price falls back to the card price (D-231).
	if got[1].GetPriceUsd() != 1.25 {
		t.Errorf("price of the unpriced printing = %v, want the card price 1.25", got[1].GetPriceUsd())
	}
	// A row the index does not know keeps its empty fields.
	if len(got[2].GetColors()) != 0 || got[2].GetPriceUsd() != 0 {
		t.Errorf("an unknown row = %v", got[2])
	}
	// The head reads no entry, so it fills nothing (D-392).
	head, err := s.GetCollection(context.Background(), connect.NewRequest(&mtgv1.GetCollectionRequest{
		CollectionId: "col-1", EntriesOmitted: true,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if len(head.Msg.GetCollection().GetEntries()) != 0 {
		t.Error("the head carries entries")
	}
}
