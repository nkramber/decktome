package collections

import (
	"sort"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// The binder head reads a collection's shape and never its rows (D-392).
// The import computes that shape once. Before this, the head asked for
// every entry and counted them in the browser, and the owner's 2,657
// rows moved about a megabyte on every page load.

// TopSets is how many sets the summary names. The head shows a short
// list, and a collection can hold cards of hundreds of sets.
const TopSets = 6

// CardSource resolves an Oracle id to a card. The color counts read the
// color identity, which a collection entry does not carry.
type CardSource interface {
	ByOracleID(id string) (*mtgv1.Card, bool)
}

// Summarize counts the shape of a collection (D-392). cards may be nil,
// and the color counts are then empty: every other count reads the
// entries alone, because the rarity and the set name ride on every row.
//
// Every count is of cards and not of rows, copies included. A row of
// four Lightning Bolt counts four times. RowCount is the exception, and
// it names what it counts.
func Summarize(entries []*mtgv1.CollectionEntry, cards CardSource) *mtgv1.CollectionSummary {
	out := &mtgv1.CollectionSummary{
		RowCount: int32(len(entries)), //nolint:gosec // an import caps the rows far under int32
		ByRarity: map[string]int32{},
		ByColor:  map[string]int32{},
	}
	oracle := map[string]bool{}
	type setKey struct{ code, name string }
	bySet := map[setKey]int32{}
	for _, e := range entries {
		q := e.GetQuantity()
		if e.GetOracleId() != "" {
			oracle[e.GetOracleId()] = true
		}
		if r := e.GetRarity(); r != "" {
			out.ByRarity[r] += q
		}
		if e.GetSetCode() != "" || e.GetSetName() != "" {
			bySet[setKey{code: e.GetSetCode(), name: e.GetSetName()}] += q
		}
		if cards == nil {
			continue
		}
		c, ok := cards.ByOracleID(e.GetOracleId())
		if !ok {
			continue
		}
		// A card of two colors counts once under each. A colorless card
		// counts under none, which is what a color filter reads.
		for _, col := range c.GetColorIdentity() {
			out.ByColor[col.String()] += q
		}
	}
	out.UniqueCards = int32(len(oracle)) //nolint:gosec // one entry per row, and the rows are capped
	sets := make([]*mtgv1.SetCount, 0, len(bySet))
	for k, n := range bySet {
		sets = append(sets, &mtgv1.SetCount{SetCode: k.code, SetName: k.name, Count: n})
	}
	// Largest first, and the set code breaks every tie, so two runs of
	// one collection give one answer.
	sort.SliceStable(sets, func(i, j int) bool {
		if sets[i].GetCount() != sets[j].GetCount() {
			return sets[i].GetCount() > sets[j].GetCount()
		}
		return sets[i].GetSetCode() < sets[j].GetSetCode()
	})
	if len(sets) > TopSets {
		sets = sets[:TopSets]
	}
	out.TopSets = sets
	return out
}

// rowKey is what makes two entries the same row (D-393). A ManaBox
// export writes one row per printing, finish, and condition, so a foil
// and a normal copy of one card are two rows.
type rowKey struct {
	scryfallID string
	finish     mtgv1.Finish
	condition  mtgv1.Condition
}

func keyOf(e *mtgv1.CollectionEntry) rowKey {
	return rowKey{
		scryfallID: e.GetScryfallId(),
		finish:     e.GetFinish(),
		condition:  e.GetCondition(),
	}
}

// Diff compares a stored collection with an uploaded one (D-393). It
// stores nothing: a reader reads it and then chooses whether to replace.
//
// The order of every list is the order of the side it came from, so two
// runs of one pair read the same.
func Diff(stored, uploaded []*mtgv1.CollectionEntry) *mtgv1.CollectionDiff {
	have := make(map[rowKey]*mtgv1.CollectionEntry, len(stored))
	for _, e := range stored {
		have[keyOf(e)] = e
	}
	seen := make(map[rowKey]bool, len(uploaded))
	out := &mtgv1.CollectionDiff{}
	for _, e := range uploaded {
		k := keyOf(e)
		seen[k] = true
		old, ok := have[k]
		if !ok {
			out.Added = append(out.Added, e)
			out.AddedCards += e.GetQuantity()
			continue
		}
		if old.GetQuantity() == e.GetQuantity() {
			continue
		}
		out.Changed = append(out.Changed, &mtgv1.QuantityChange{
			Entry: e, From: old.GetQuantity(), To: e.GetQuantity(),
		})
		// The count is what the change moves, in either direction. A row
		// that falls from 4 to 1 moves three cards.
		if d := e.GetQuantity() - old.GetQuantity(); d < 0 {
			out.ChangedCards -= d
		} else {
			out.ChangedCards += d
		}
	}
	for _, e := range stored {
		if seen[keyOf(e)] {
			continue
		}
		out.Removed = append(out.Removed, e)
		out.RemovedCards += e.GetQuantity()
	}
	out.Identical = len(out.Added) == 0 && len(out.Removed) == 0 && len(out.Changed) == 0
	return out
}
