package collections

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// binderRows are three decorated rows, as GetCollection fills them
// before the filter reads them (D-396).
func binderRows() []*mtgv1.CollectionEntry {
	return []*mtgv1.CollectionEntry{
		{Name: "Lightning Bolt", SetCode: "lea", Quantity: 4, Colors: []mtgv1.Color{mtgv1.Color_COLOR_R}, CardTypes: []string{"Instant"}, PriceUsd: 2},
		{Name: "Jace, the Mind Sculptor", SetCode: "wwk", Quantity: 1, Colors: []mtgv1.Color{mtgv1.Color_COLOR_U}, CardTypes: []string{"Planeswalker"}, PriceUsd: 90},
		{Name: "Path to Exile", SetCode: "lea", Quantity: 3, Colors: []mtgv1.Color{mtgv1.Color_COLOR_W}, CardTypes: []string{"Instant"}, PriceUsd: 5},
	}
}

func names(rows []*mtgv1.CollectionEntry) []string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, r.GetName())
	}
	return out
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TestFilterKeepsWhatEveryChoiceMatches is D-398. Each field narrows
// the rows, and the choices run together.
func TestFilterKeepsWhatEveryChoiceMatches(t *testing.T) {
	rows := binderRows()
	for _, tc := range []struct {
		name string
		f    *mtgv1.BinderFilter
		want []string
	}{
		{"no choice keeps every row, in name order", &mtgv1.BinderFilter{}, []string{"Jace, the Mind Sculptor", "Lightning Bolt", "Path to Exile"}},
		{"the name, whatever the case", &mtgv1.BinderFilter{Query: "  BOLT "}, []string{"Lightning Bolt"}},
		{"the set, whatever the case", &mtgv1.BinderFilter{SetCode: "LEA"}, []string{"Lightning Bolt", "Path to Exile"}},
		{"one color", &mtgv1.BinderFilter{Color: mtgv1.Color_COLOR_U}, []string{"Jace, the Mind Sculptor"}},
		{"one type", &mtgv1.BinderFilter{CardType: "Instant"}, []string{"Lightning Bolt", "Path to Exile"}},
		{"one count", &mtgv1.BinderFilter{Quantity: 3}, []string{"Path to Exile"}},
		{"a count or more", &mtgv1.BinderFilter{Quantity: 3, QuantityOrMore: true}, []string{"Lightning Bolt", "Path to Exile"}},
		{"every choice together", &mtgv1.BinderFilter{Query: "path", SetCode: "lea", Color: mtgv1.Color_COLOR_W, CardType: "Instant", Quantity: 3}, []string{"Path to Exile"}},
		{"a choice that matches nothing", &mtgv1.BinderFilter{SetCode: "lea", Color: mtgv1.Color_COLOR_U}, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := names(Apply(rows, FilterOf(tc.f), mtgv1.BinderSort_BINDER_SORT_UNSPECIFIED))
			if !equal(got, tc.want) {
				t.Errorf("kept %v, want %v", got, tc.want)
			}
		})
	}
}

// TestFilterReadsAGoldCardUnderBothColors: a card of two colors answers
// to each, and a card of none answers to colorless alone.
func TestFilterReadsAGoldCardUnderBothColors(t *testing.T) {
	rows := []*mtgv1.CollectionEntry{
		{Name: "Boros Charm", Colors: []mtgv1.Color{mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W}},
		{Name: "Sol Ring", CardTypes: []string{"Artifact"}},
	}
	for _, c := range []mtgv1.Color{mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W} {
		if got := names(Apply(rows, FilterOf(&mtgv1.BinderFilter{Color: c}), 0)); !equal(got, []string{"Boros Charm"}) {
			t.Errorf("color %v kept %v", c, got)
		}
	}
	if got := names(Apply(rows, FilterOf(&mtgv1.BinderFilter{Color: mtgv1.Color_COLOR_U}), 0)); len(got) != 0 {
		t.Errorf("blue kept %v", got)
	}
	if got := names(Apply(rows, FilterOf(&mtgv1.BinderFilter{Colorless: true}), 0)); !equal(got, []string{"Sol Ring"}) {
		t.Errorf("colorless kept %v", got)
	}
	// Colorless wins over a color, so the two never contradict.
	if got := names(Apply(rows, FilterOf(&mtgv1.BinderFilter{Colorless: true, Color: mtgv1.Color_COLOR_R}), 0)); !equal(got, []string{"Sol Ring"}) {
		t.Errorf("colorless beside red kept %v", got)
	}
}

// TestSortBinderOrders: each sort, and the name breaks every tie.
func TestSortBinderOrders(t *testing.T) {
	rows := binderRows()
	for _, tc := range []struct {
		by   mtgv1.BinderSort
		want []string
	}{
		{mtgv1.BinderSort_BINDER_SORT_UNSPECIFIED, []string{"Jace, the Mind Sculptor", "Lightning Bolt", "Path to Exile"}},
		{mtgv1.BinderSort_BINDER_SORT_NAME, []string{"Jace, the Mind Sculptor", "Lightning Bolt", "Path to Exile"}},
		{mtgv1.BinderSort_BINDER_SORT_COUNT, []string{"Lightning Bolt", "Path to Exile", "Jace, the Mind Sculptor"}},
		{mtgv1.BinderSort_BINDER_SORT_SET, []string{"Lightning Bolt", "Path to Exile", "Jace, the Mind Sculptor"}},
		{mtgv1.BinderSort_BINDER_SORT_PRICE, []string{"Jace, the Mind Sculptor", "Path to Exile", "Lightning Bolt"}},
	} {
		got := names(Apply(rows, Filter{}, tc.by))
		if !equal(got, tc.want) {
			t.Errorf("sort %v gave %v, want %v", tc.by, got, tc.want)
		}
	}
	// An unpriced row sorts last under the price sort.
	withUnpriced := append(binderRows(), &mtgv1.CollectionEntry{Name: "Aaa Unpriced"})
	got := names(Apply(withUnpriced, Filter{}, mtgv1.BinderSort_BINDER_SORT_PRICE))
	if got[len(got)-1] != "Aaa Unpriced" {
		t.Errorf("price sort gave %v, want the unpriced row last", got)
	}
}

// TestApplyChangesNoInput: the stored order of the entries is the order
// the pages read, and a sort must not move it.
func TestApplyChangesNoInput(t *testing.T) {
	rows := binderRows()
	before := names(rows)
	Apply(rows, Filter{}, mtgv1.BinderSort_BINDER_SORT_COUNT)
	if got := names(rows); !equal(got, before) {
		t.Errorf("the input moved: %v", got)
	}
}
