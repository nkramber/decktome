package cards

import (
	"context"
	"slices"
	"testing"
)

// TestNewlyLegal is D-1091: a card is new when it turns legal in a
// format of the app for the first time. A preview card is not_legal in
// each format, so it is new at its release and not at its preview.
func TestNewlyLegal(t *testing.T) {
	notLegal := map[string]string{"commander": "not_legal", "modern": "not_legal", "standard": "not_legal"}
	released := map[string]string{"commander": "legal", "modern": "legal", "standard": "legal"}
	legacyOnly := map[string]string{"commander": "banned", "modern": "banned", "legacy": "legal"}
	unbanned := map[string]string{"commander": "legal", "modern": "banned", "legacy": "legal"}
	tests := []struct {
		name          string
		before, after map[string]map[string]string
		want          []string
	}{
		{"a preview is not new", nil, map[string]map[string]string{"p": notLegal}, nil},
		{"a release of a preview is new", map[string]map[string]string{"p": notLegal}, map[string]map[string]string{"p": released}, []string{"p"}},
		{"a card with no preview is new", nil, map[string]map[string]string{"s": released}, []string{"s"}},
		{"a reprint is not new", map[string]map[string]string{"r": released}, map[string]map[string]string{"r": released}, nil},
		{"an unban is not new", map[string]map[string]string{"u": released}, map[string]map[string]string{"u": unbanned}, nil},
		{"legal in another format alone is not new", nil, map[string]map[string]string{"l": legacyOnly}, nil},
		// A Commander unban of a card banned in Modern and absent from
		// Standard: the card was legal in Legacy, so it is not new.
		{"an unban from banned in each app format is not new", map[string]map[string]string{"l": legacyOnly}, map[string]map[string]string{"l": unbanned}, nil},
		{"sorted", nil, map[string]map[string]string{"b": released, "a": released}, []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewlyLegal(tt.before, tt.after); !slices.Equal(got, tt.want) {
				t.Fatalf("NewlyLegal = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestCompareVersions reads both versions once and agrees with
// LegalityDiff on the count.
func TestCompareVersions(t *testing.T) {
	ctx := context.Background()
	store := DirStore{Root: t.TempDir()}
	writeSnapshot(t, store, "20260823T090000",
		cardLine("a", map[string]string{"modern": "legal"})+"\n"+
			cardLine("p", map[string]string{"modern": "not_legal"})+"\n")
	writeSnapshot(t, store, "20260824T090000",
		cardLine("a", map[string]string{"modern": "banned"})+"\n"+
			cardLine("p", map[string]string{"modern": "legal"})+"\n")
	changed, fresh, err := CompareVersions(ctx, store, "20260823T090000", "20260824T090000")
	if err != nil {
		t.Fatal(err)
	}
	if changed != 2 || !slices.Equal(fresh, []string{"p"}) {
		t.Fatalf("changed=%d new=%v, want 2 and [p]", changed, fresh)
	}
	want, err := LegalityDiff(ctx, store, "20260823T090000", "20260824T090000")
	if err != nil || want != changed {
		t.Fatalf("LegalityDiff = %d (%v), CompareVersions = %d", want, err, changed)
	}
}

// TestNewCardsMarker writes and reads the marker, and LatestNewCards
// finds the newest version that carries one.
func TestNewCardsMarker(t *testing.T) {
	ctx := context.Background()
	store := DirStore{Root: t.TempDir()}
	writeSnapshot(t, store, "20260823T090000", "")
	writeSnapshot(t, store, "20260824T090000", "")
	if _, _, ok, err := LatestNewCards(ctx, store); err != nil || ok {
		t.Fatalf("no marker yet: ok=%v err=%v", ok, err)
	}
	old := NewCardsRecord{SnapshotAsOf: "2026-08-23T09:00:00Z", Cards: []string{"a"}}
	if err := WriteNewCards(ctx, store, "20260823T090000", old); err != nil {
		t.Fatal(err)
	}
	rec := NewCardsRecord{SnapshotAsOf: "2026-08-24T09:00:00Z", Cards: []string{"b", "c"}}
	if err := WriteNewCards(ctx, store, "20260824T090000", rec); err != nil {
		t.Fatal(err)
	}
	version, got, ok, err := LatestNewCards(ctx, store)
	if err != nil || !ok {
		t.Fatalf("LatestNewCards: ok=%v err=%v", ok, err)
	}
	if version != "20260824T090000" || !slices.Equal(got.Cards, rec.Cards) || got.Pass != "" {
		t.Fatalf("LatestNewCards = %s %+v", version, got)
	}
	got.Pass, got.Decks = "2026-08-24T10:00:00Z", 3
	if err := WriteNewCards(ctx, store, version, got); err != nil {
		t.Fatal(err)
	}
	if again, ok, err := ReadNewCards(ctx, store, version); err != nil || !ok || again.Pass == "" || again.Decks != 3 {
		t.Fatalf("ReadNewCards = %+v ok=%v err=%v", again, ok, err)
	}
}
