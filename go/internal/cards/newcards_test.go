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

// TestNewCardsMarker writes and reads the marker, and PendingNewCards
// finds each marker with no ended pass, oldest first.
func TestNewCardsMarker(t *testing.T) {
	ctx := context.Background()
	store := DirStore{Root: t.TempDir()}
	writeSnapshot(t, store, "20260823T090000", "")
	writeSnapshot(t, store, "20260824T090000", "")
	if pending, err := PendingNewCards(ctx, store); err != nil || len(pending) != 0 {
		t.Fatalf("no marker yet: %+v err=%v", pending, err)
	}
	old := NewCardsRecord{SnapshotAsOf: "2026-08-23T09:00:00Z", Cards: []string{"a"}}
	if err := WriteNewCards(ctx, store, "20260823T090000", old); err != nil {
		t.Fatal(err)
	}
	rec := NewCardsRecord{SnapshotAsOf: "2026-08-24T09:00:00Z", Cards: []string{"b", "c"}}
	if err := WriteNewCards(ctx, store, "20260824T090000", rec); err != nil {
		t.Fatal(err)
	}
	pending, err := PendingNewCards(ctx, store)
	if err != nil || len(pending) != 2 {
		t.Fatalf("PendingNewCards = %+v err=%v, want two markers", pending, err)
	}
	if pending[0].Version != "20260823T090000" || pending[1].Version != "20260824T090000" || !slices.Equal(pending[1].Record.Cards, rec.Cards) {
		t.Fatalf("PendingNewCards = %+v, want the oldest first", pending)
	}
	version, got := pending[1].Version, pending[1].Record
	got.Pass, got.Decks = "2026-08-24T10:00:00Z", 3
	if err := WriteNewCards(ctx, store, version, got); err != nil {
		t.Fatal(err)
	}
	if again, ok, err := ReadNewCards(ctx, store, version); err != nil || !ok || again.Pass == "" || again.Decks != 3 {
		t.Fatalf("ReadNewCards = %+v ok=%v err=%v", again, ok, err)
	}
	// The newer marker ended, and the older one still waits for its pass.
	if pending, err = PendingNewCards(ctx, store); err != nil || len(pending) != 1 || pending[0].Version != "20260823T090000" {
		t.Fatalf("after the newer pass: %+v err=%v, want the older marker alone", pending, err)
	}
}
