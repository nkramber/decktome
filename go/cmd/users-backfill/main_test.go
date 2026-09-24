package main

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/gzstore"
)

func packed(t *testing.T, d *mtgv1.Deck) []byte {
	t.Helper()
	b, err := gzstore.MarshalProto(d)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// TestDeckKind is D-861: the backfill counted no revision, because it
// read a flat field that no stored deck held. The flat fields answer
// now, and the packed deck answers for a document written before them.
func TestDeckKind(t *testing.T) {
	for _, tc := range []struct {
		name     string
		data     map[string]any
		rev, imp bool
	}{
		{"flat first build", map[string]any{"revised_from_deck_id": "", "imported": false}, false, false},
		{"flat revision", map[string]any{"revised_from_deck_id": "d1", "imported": false}, true, false},
		{"flat import", map[string]any{"revised_from_deck_id": "", "imported": true}, false, true},
		{"old revision", map[string]any{"deck_gz": packed(t, &mtgv1.Deck{RevisedFromDeckId: "d1"})}, true, false},
		{"old first build", map[string]any{"deck_gz": packed(t, &mtgv1.Deck{Name: "x"})}, false, false},
		{"no payload", map[string]any{}, false, false},
	} {
		rev, imp := deckKind(tc.data)
		if rev != tc.rev || imp != tc.imp {
			t.Errorf("%s: revision %v, import %v, want %v and %v", tc.name, rev, imp, tc.rev, tc.imp)
		}
	}
}
