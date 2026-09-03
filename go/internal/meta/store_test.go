package meta

import (
	"context"
	"testing"
)

func TestMergeListsAndAll(t *testing.T) {
	ctx := context.Background()
	s := DirObjects{Root: t.TempDir()}
	lists := []List{
		{Source: SourceMTGO, ID: "a/1", Format: FormatModern, Date: "2026-09-02", Tier: TierGreat, Cards: []Card{{Name: "Plains", Count: 4}}},
		{Source: SourceMTGO, ID: "a/2", Format: FormatModern, Date: "2026-09-02", Tier: TierGood, Cards: []Card{{Name: "Island", Count: 4}}},
		{Source: SourceMTGO, ID: "b/1", Format: FormatModern, Date: "2026-08-30", Tier: TierGood, Cards: []Card{{Name: "Swamp", Count: 4}}},
		{Source: SourceMTGJSON, ID: "p", Format: FormatSixty, Date: "2020-01-01", Tier: TierBaseline, Cards: []Card{{Name: "Forest", Count: 4}}},
		{Source: SourceMTGO, ID: "nodate", Format: FormatModern, Tier: TierGood},
	}
	added, err := MergeLists(ctx, s, lists)
	if err != nil {
		t.Fatal(err)
	}
	if added != 4 {
		t.Fatalf("added = %d, want 4 (the list with no date is out)", added)
	}
	// The same lists again change nothing, and a changed one replaces.
	lists[0].Tier = TierGood
	added, err = MergeLists(ctx, s, lists[:1])
	if err != nil {
		t.Fatal(err)
	}
	if added != 0 {
		t.Fatalf("second merge added %d", added)
	}
	modern, err := AllLists(ctx, s, FormatModern)
	if err != nil {
		t.Fatal(err)
	}
	if len(modern) != 3 {
		t.Fatalf("modern = %d, want 3", len(modern))
	}
	for _, l := range modern {
		if l.ID == "a/1" && l.Tier != TierGood {
			t.Errorf("a/1 tier = %s, want the replaced value", l.Tier)
		}
	}
	names, err := s.List(ctx, ListsPrefix)
	if err != nil {
		t.Fatal(err)
	}
	if len(names) != 3 {
		t.Errorf("objects = %v, want two modern months and one sixty", names)
	}
}

func TestRawRoundTrip(t *testing.T) {
	ctx := context.Background()
	s := DirObjects{Root: t.TempDir()}
	ok, err := HasRaw(ctx, s, SourceMTGO, "x/y")
	if err != nil || ok {
		t.Fatalf("has before put = %v, %v", ok, err)
	}
	if err := PutRaw(ctx, s, SourceMTGO, "x/y", []byte("<html>")); err != nil {
		t.Fatal(err)
	}
	page, ok, err := GetRaw(ctx, s, SourceMTGO, "x/y")
	if err != nil || !ok || string(page) != "<html>" {
		t.Fatalf("get = %q, %v, %v", page, ok, err)
	}
}

func TestModelVersions(t *testing.T) {
	ctx := context.Background()
	s := DirObjects{Root: t.TempDir()}
	v, err := LatestModel(ctx, s)
	if err != nil || v != "" {
		t.Fatalf("empty store: %q, %v", v, err)
	}
	if err := WriteModel(ctx, s, "20260901T000000Z", []byte(`{"a":1}`)); err != nil {
		t.Fatal(err)
	}
	if err := WriteModel(ctx, s, "20260902T000000Z", []byte(`{"a":2}`)); err != nil {
		t.Fatal(err)
	}
	// A version with no marker is mid-write.
	if err := s.Put(ctx, ModelPrefix+"20260903T000000Z/"+ModelFile, []byte("x")); err != nil {
		t.Fatal(err)
	}
	v, err = LatestModel(ctx, s)
	if err != nil || v != "20260902T000000Z" {
		t.Fatalf("latest = %q, %v", v, err)
	}
	data, err := ReadModel(ctx, s, v)
	if err != nil || string(data) != `{"a":2}` {
		t.Fatalf("read = %s, %v", data, err)
	}
}

func TestPreconsAndCommanders(t *testing.T) {
	ctx := context.Background()
	s := DirObjects{Root: t.TempDir()}
	table := []Precon{{Name: "B", Code: "X", Type: "Commander Deck", ReleaseDate: "2020-01-01", Cards: []PreconCard{{Name: "Sol Ring", Count: 1}}},
		{Name: "A", Code: "X", Type: "Commander Deck", ReleaseDate: "2020-01-01"}}
	SortPrecons(table)
	if table[0].Name != "A" {
		t.Errorf("sort by name inside a date")
	}
	if err := WritePrecons(ctx, s, "5.3.0+20260902", table); err != nil {
		t.Fatal(err)
	}
	v, err := LatestPreconsVersion(ctx, s)
	if err != nil || v != "5.3.0+20260902" {
		t.Fatalf("latest precons = %q, %v", v, err)
	}
	got, err := ReadPrecons(ctx, s, v)
	if err != nil || len(got) != 2 || got[1].Cards[0].Name != "Sol Ring" {
		t.Fatalf("precons = %+v, %v", got, err)
	}
	reads := []Commander{{Name: "K", Slug: "k", NumDecks: 5, BracketCounts: map[int]int{5: 4, 3: 1}}}
	if err := WriteCommanders(ctx, s, "2026-09-02", reads); err != nil {
		t.Fatal(err)
	}
	if err := WriteCommanders(ctx, s, "2026-09-02", []Commander{{Name: "J", Slug: "j"}}); err != nil {
		t.Fatal(err)
	}
	day, err := LatestCommandersDay(ctx, s)
	if err != nil || day != "2026-09-02" {
		t.Fatalf("day = %q, %v", day, err)
	}
	cs, err := ReadCommanders(ctx, s, day)
	if err != nil || len(cs) != 2 || cs[1].Slug != "k" || cs[1].HighBracketShare() != 0.8 {
		t.Fatalf("commanders = %+v, %v", cs, err)
	}
}

func TestDirObjectsRejectsEscape(t *testing.T) {
	s := DirObjects{Root: t.TempDir()}
	if err := s.Put(context.Background(), "../x", []byte("x")); err == nil {
		t.Fatal("a name that leaves the root must fail")
	}
}
