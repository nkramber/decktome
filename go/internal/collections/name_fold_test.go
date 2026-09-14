package collections

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestFilterReadsANameWithoutItsAccent is D-716: the binder search finds
// a card whatever the accent or the quote the reader typed.
func TestFilterReadsANameWithoutItsAccent(t *testing.T) {
	rows := []*mtgv1.CollectionEntry{
		{Name: "Gríma, Saruman's Footman", Quantity: 1},
		{Name: "Lightning Bolt", Quantity: 4},
	}
	for _, q := range []string{"grima", "GRÍMA", "Saruman’s"} {
		got := names(Apply(rows, FilterOf(&mtgv1.BinderFilter{Query: q}), mtgv1.BinderSort_BINDER_SORT_UNSPECIFIED))
		if !equal(got, []string{"Gríma, Saruman's Footman"}) {
			t.Errorf("query %q kept %v", q, got)
		}
	}
}
