package collections

import (
	"slices"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// Filter is what the binder grid asks for (D-398). Every field runs on
// the server over every row, so a match on a later page still shows.
// The colors, the card types, and the price ride on the entry, and
// GetCollection fills them from the card index of the day before it
// filters (D-396).
type Filter struct {
	// Query keeps a row whose card name holds this text, without regard
	// to case. Empty keeps them all.
	Query string
	// SetCode keeps one set. Empty keeps them all.
	SetCode string
	// Color keeps the cards of this color, and a card of two colors
	// answers to both. COLOR_UNSPECIFIED keeps them all.
	Color mtgv1.Color
	// Colorless keeps the cards of no color. It wins over Color.
	Colorless bool
	// CardType keeps one card type, as the type line names it. Empty
	// keeps them all.
	CardType string
	// Quantity keeps the rows of this many copies, or of this many or
	// more when QuantityOrMore is set. Zero keeps them all.
	Quantity       int32
	QuantityOrMore bool
}

// FilterOf reads the proto shape. It trims and lowers the query once,
// so Keep reads no text twice.
func FilterOf(f *mtgv1.BinderFilter) Filter {
	return Filter{
		Query:          strings.ToLower(strings.TrimSpace(f.GetQuery())),
		SetCode:        strings.ToLower(strings.TrimSpace(f.GetSetCode())),
		Color:          f.GetColor(),
		Colorless:      f.GetColorless(),
		CardType:       strings.TrimSpace(f.GetCardType()),
		Quantity:       f.GetQuantity(),
		QuantityOrMore: f.GetQuantityOrMore(),
	}
}

// Keep reports whether one row passes the filter. A row the card index
// does not know carries no color and no type, so a color or a type
// filter keeps no such row.
func (f Filter) Keep(e *mtgv1.CollectionEntry) bool {
	if f.SetCode != "" && !strings.EqualFold(e.GetSetCode(), f.SetCode) {
		return false
	}
	switch {
	case f.Colorless:
		if len(e.GetColors()) != 0 {
			return false
		}
	case f.Color != mtgv1.Color_COLOR_UNSPECIFIED:
		if !slices.Contains(e.GetColors(), f.Color) {
			return false
		}
	}
	if f.CardType != "" && !slices.Contains(e.GetCardTypes(), f.CardType) {
		return false
	}
	if f.Quantity > 0 {
		q := e.GetQuantity()
		if f.QuantityOrMore && q < f.Quantity {
			return false
		}
		if !f.QuantityOrMore && q != f.Quantity {
			return false
		}
	}
	if f.Query != "" && !strings.Contains(strings.ToLower(e.GetName()), f.Query) {
		return false
	}
	return true
}

// Apply keeps the rows the filter passes and orders them. It returns a
// new slice, so the stored order of the entries never changes.
func Apply(entries []*mtgv1.CollectionEntry, f Filter, by mtgv1.BinderSort) []*mtgv1.CollectionEntry {
	out := make([]*mtgv1.CollectionEntry, 0, len(entries))
	for _, e := range entries {
		if f.Keep(e) {
			out = append(out, e)
		}
	}
	SortBinder(out, by)
	return out
}

// SortBinder orders the rows of a binder. The name breaks every tie,
// and the stored order of SortEntries breaks a tie on the name, so one
// binder reads the same on every visit.
func SortBinder(entries []*mtgv1.CollectionEntry, by mtgv1.BinderSort) {
	sort.SliceStable(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		switch by {
		case mtgv1.BinderSort_BINDER_SORT_COUNT:
			if a.GetQuantity() != b.GetQuantity() {
				return a.GetQuantity() > b.GetQuantity()
			}
		case mtgv1.BinderSort_BINDER_SORT_SET:
			if a.GetSetCode() != b.GetSetCode() {
				return a.GetSetCode() < b.GetSetCode()
			}
		case mtgv1.BinderSort_BINDER_SORT_PRICE:
			// Dearest first, because a reader who sorts by price looks
			// for the cards that carry the value. An unpriced row is a
			// zero, so it sorts last.
			if a.GetPriceUsd() != b.GetPriceUsd() {
				return a.GetPriceUsd() > b.GetPriceUsd()
			}
		case mtgv1.BinderSort_BINDER_SORT_UNSPECIFIED, mtgv1.BinderSort_BINDER_SORT_NAME:
		}
		return a.GetName() < b.GetName()
	})
}
