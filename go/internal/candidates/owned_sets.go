package candidates

import (
	"math"
	"strings"

	"github.com/nkramber/decktome/go/internal/cards"
)

// OwnedInSets keeps the owned copies of a deck limited to sets (D-1041).
// A copy counts when its printing is in one of the sets. A reader who
// limits the deck to the Lord of the Rings sets owns Ghost Quarter for it
// only through a copy from those sets: a Secret Lair copy is not one.
//
// printings maps each Scryfall id to the copies the reader holds. A
// basic land keeps its whole count, because a basic land never leaves
// the pool (D-37). A copy with no printing, or a printing the snapshot
// does not hold, does not count. owned is the whole count map, and it
// answers unchanged when no set limit applies.
func OwnedInSets(idx *cards.Index, owned, printings map[string]int32, codes []string) map[string]int32 {
	inSet := cards.CodeSet(codes)
	if idx == nil || len(inSet) == 0 || owned == nil {
		return owned
	}
	basic := BasicLandByOracle(idx)
	out := map[string]int32{}
	for id, n := range owned {
		if basic(id) {
			out[id] = n
		}
	}
	for id, n := range printings {
		p, ok := idx.Printing(id)
		if !ok || n <= 0 || !inSet[strings.ToLower(p.GetSetCode())] {
			continue
		}
		c, ok := idx.ByPrintingID(id)
		if !ok || basic(c.GetOracleId()) {
			continue
		}
		out[c.GetOracleId()] = int32(min(int64(out[c.GetOracleId()])+int64(n), math.MaxInt32))
	}
	return out
}
