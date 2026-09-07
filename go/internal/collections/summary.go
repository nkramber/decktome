package collections

import (
	"sort"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The binder head reads a collection's shape and never its rows (D-392).
// The import computes that shape once. Before this, the head asked for
// every entry and counted them in the browser, and the owner's 2,657
// rows moved about a megabyte on every page load.

// ArtCards is how many cards the binder head shows the art of. The head
// reads no entry, so the summary carries the ids (D-392).
const ArtCards = 10

// artRank orders the rarities the art strip prefers. The rarest come
// first, so the strip shows the collection at its best.
var artRank = map[string]int{"mythic": 0, "rare": 1, "uncommon": 2, "common": 3}

// CardSource resolves an Oracle id to a card. The type counts read the
// card types, which a collection entry does not carry (D-398).
type CardSource interface {
	ByOracleID(id string) (*mtgv1.Card, bool)
}

// Summarize counts the shape of a collection (D-392). cards may be nil,
// and the type counts are then empty: every other count reads the
// entries alone, because the rarity and the set name ride on every row.
//
// Every count is of cards and not of rows, copies included. A row of
// four Lightning Bolt counts four times. RowCount is the exception, and
// it names what it counts.
func Summarize(entries []*mtgv1.CollectionEntry, cards CardSource) *mtgv1.CollectionSummary {
	out := &mtgv1.CollectionSummary{
		RowCount: int32(len(entries)), //nolint:gosec // an import caps the rows far under int32
		ByRarity: map[string]int32{},
		ByType:   map[string]int32{},
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
		// A card of two types counts once under each, which is what the
		// type filter reads (D-398).
		for _, t := range c.GetCardTypes() {
			out.ByType[t] += q
		}
	}
	out.UniqueCards = int32(len(oracle)) //nolint:gosec // one entry per row, and the rows are capped
	sets := make([]*mtgv1.SetCount, 0, len(bySet))
	for k, n := range bySet {
		sets = append(sets, &mtgv1.SetCount{SetCode: k.code, SetName: k.name, Count: n})
	}
	// Largest first, and the set code breaks every tie, so two runs of
	// one collection give one answer. The list holds every set: the
	// binder's set filter offers each one (D-398).
	sort.SliceStable(sets, func(i, j int) bool {
		if sets[i].GetCount() != sets[j].GetCount() {
			return sets[i].GetCount() > sets[j].GetCount()
		}
		return sets[i].GetSetCode() < sets[j].GetSetCode()
	})
	out.Sets = sets
	out.ArtOracleIds = artIDs(entries)
	return out
}

// artIDs picks the cards the head shows the art of, the rarest first
// (D-392). A card the reader owns in two printings appears once, and
// the name breaks every tie, so two runs of one collection give one
// strip.
func artIDs(entries []*mtgv1.CollectionEntry) []string {
	seen := map[string]bool{}
	picked := make([]*mtgv1.CollectionEntry, 0, len(entries))
	for _, e := range entries {
		id := e.GetOracleId()
		if id == "" || seen[id] {
			continue
		}
		seen[id] = true
		picked = append(picked, e)
	}
	sort.SliceStable(picked, func(i, j int) bool {
		ri, ok := artRank[picked[i].GetRarity()]
		if !ok {
			ri = len(artRank)
		}
		rj, ok := artRank[picked[j].GetRarity()]
		if !ok {
			rj = len(artRank)
		}
		if ri != rj {
			return ri < rj
		}
		return picked[i].GetName() < picked[j].GetName()
	})
	if len(picked) > ArtCards {
		picked = picked[:ArtCards]
	}
	out := make([]string, 0, len(picked))
	for _, e := range picked {
		out = append(out, e.GetOracleId())
	}
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
