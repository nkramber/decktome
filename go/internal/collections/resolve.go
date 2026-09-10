package collections

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// Resolve joins parsed rows against the card index.
// Join order: Scryfall ID, then set code plus collector number,
// then exact name. Non-English rows are reported, never imported (D-23).
// A token, emblem, or art card row is reported as NOT_PLAYABLE before
// the name fallback, so it never counts as the real card.
//
// Rows with the same printing, finish, and condition merge into one
// entry. Quantities add up, so ownership counts do not change. A merged
// entry stops at maxQuantity: the row that would push it over is
// reported as BAD_ROW, the same reason the parser gives one row over the
// cap. The row then counts in the report and not in the collection, so
// resolved plus unresolved rows still equals the input rows.
func Resolve(rows []Row, idx *cards.Index) ([]*mtgv1.CollectionEntry, []*mtgv1.UnresolvedRow) {
	var entries []*mtgv1.CollectionEntry
	var bad []*mtgv1.UnresolvedRow
	merged := map[string]*mtgv1.CollectionEntry{}
	for _, row := range rows {
		if row.Language != "" && !strings.EqualFold(row.Language, "en") {
			bad = append(bad, unresolved(row.Line, row.Raw, mtgv1.UnresolvedReason_UNRESOLVED_REASON_NON_ENGLISH))
			continue
		}
		card, byName, reason := resolveOne(row, idx)
		if card == nil {
			bad = append(bad, unresolved(row.Line, row.Raw, reason))
			continue
		}
		e := buildEntry(row, card, byName, idx)
		key := strings.Join([]string{e.ScryfallId, e.SetCode, e.CollectorNumber, e.Finish.String(), e.Condition.String()}, "|")
		if prev, dup := merged[key]; dup {
			if int(prev.Quantity)+int(e.Quantity) > maxQuantity {
				bad = append(bad, unresolved(row.Line, row.Raw, mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW))
				continue
			}
			prev.Quantity += e.Quantity
			continue
		}
		merged[key] = e
		entries = append(entries, e)
	}
	return entries, bad
}

// buildEntry fills an entry from a row and its card. When the match
// came from the name alone, the input id and number point at a printing
// the index does not know. The entry then takes the card's
// default printing id, and keeps the input set code and number only when
// they match that printing. The proto has no field for this fact. A
// caller can detect it: the entry's scryfall_id differs from the input
// row's Scryfall ID.
func buildEntry(row Row, card *mtgv1.Card, byName bool, idx *cards.Index) *mtgv1.CollectionEntry {
	e := &mtgv1.CollectionEntry{
		ScryfallId:      row.ScryfallID,
		OracleId:        card.OracleId,
		Name:            card.Name,
		SetCode:         row.SetCode,
		SetName:         row.SetName,
		CollectorNumber: row.Collector,
		Quantity:        int32(row.Quantity), //nolint:gosec // parser caps at maxQuantity
		Finish:          row.Finish,
		Condition:       row.Condition,
		Rarity:          row.Rarity,
		Language:        row.Language,
	}
	if byName {
		dp := card.GetDefaultPrinting()
		e.ScryfallId = dp.GetScryfallId()
		if !strings.EqualFold(row.SetCode, dp.GetSetCode()) || !strings.EqualFold(row.Collector, dp.GetCollectorNumber()) {
			e.SetCode = dp.GetSetCode()
			e.SetName = dp.GetSetName()
			e.CollectorNumber = dp.GetCollectorNumber()
		}
	}
	// A format with no Scryfall id column resolves on the set and the
	// number, and the entry would carry no printing at all (PR-34). The
	// binder reads that field for the art and the price of what the
	// reader owns (D-299), OwnedPrintings drops an entry without one,
	// and the binder tile keys itself on it. So the entry takes the
	// printing the pair named.
	if e.ScryfallId == "" && e.SetCode != "" && e.CollectorNumber != "" {
		if p, ok := idx.PrintingBySetCollector(e.SetCode, e.CollectorNumber); ok {
			e.ScryfallId = p.GetScryfallId()
			if e.SetName == "" {
				e.SetName = p.GetSetName()
			}
			if e.Rarity == "" {
				e.Rarity = p.GetRarity()
			}
		}
	}
	return e
}

// resolveOne returns the card and whether only the name matched. A nil
// card carries the reason the row failed.
func resolveOne(row Row, idx *cards.Index) (card *mtgv1.Card, byName bool, reason mtgv1.UnresolvedReason) {
	if row.ScryfallID != "" {
		if c, ok := idx.ByPrintingID(row.ScryfallID); ok {
			return c, false, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNSPECIFIED
		}
	}
	if row.SetCode != "" && row.Collector != "" {
		if c, ok := idx.BySetCollector(row.SetCode, row.Collector); ok {
			return c, false, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNSPECIFIED
		}
	}
	// A known non-playable printing must not fall through to the name.
	if _, np := idx.NonPlayablePrinting(row.ScryfallID, row.SetCode, row.Collector); np {
		return nil, false, mtgv1.UnresolvedReason_UNRESOLVED_REASON_NOT_PLAYABLE
	}
	if row.Name != "" {
		if c, ok := idx.ByName(row.Name); ok {
			return c, true, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNSPECIFIED
		}
	}
	return nil, false, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_CARD
}

// ReasonCounts counts unresolved rows per reason name.
func ReasonCounts(bad []*mtgv1.UnresolvedRow) map[string]int32 {
	if len(bad) == 0 {
		return nil
	}
	out := map[string]int32{}
	for _, b := range bad {
		out[b.Reason.String()]++
	}
	return out
}

// ContentHash detects an identical re-upload (D-16).
func ContentHash(content []byte) string {
	sum := sha256.Sum256(content)
	return hex.EncodeToString(sum[:])
}

// OwnedPrintings lists the printing ids per Oracle id, each once, from
// the entries of one collection (D-299).
func OwnedPrintings(entries []*mtgv1.CollectionEntry) map[string][]string {
	out := map[string][]string{}
	seen := map[string]bool{}
	for _, e := range entries {
		if e.GetOracleId() == "" || e.GetScryfallId() == "" || e.GetQuantity() <= 0 {
			continue
		}
		key := e.GetOracleId() + "|" + e.GetScryfallId()
		if seen[key] {
			continue
		}
		seen[key] = true
		out[e.GetOracleId()] = append(out[e.GetOracleId()], e.GetScryfallId())
	}
	return out
}

// PrintingCounts sums owned copies per Scryfall id across finishes and
// conditions. The precon ownership check reads it: a reader owns a precon
// when the collection holds every printing of it with its count (D-408).
func PrintingCounts(entries []*mtgv1.CollectionEntry) map[string]int32 {
	out := map[string]int32{}
	for _, e := range entries {
		if e.GetScryfallId() == "" || e.GetQuantity() <= 0 {
			continue
		}
		out[e.ScryfallId] = addSaturate(out[e.ScryfallId], e.Quantity)
	}
	return out
}

// OracleCounts sums owned copies per Oracle id across printings. The
// sum saturates at the int32 maximum, it never wraps.
func OracleCounts(entries []*mtgv1.CollectionEntry) map[string]int32 {
	out := map[string]int32{}
	for _, e := range entries {
		out[e.OracleId] = addSaturate(out[e.OracleId], e.Quantity)
	}
	return out
}

// CardCount sums every copy. The sum saturates at the int32 maximum.
func CardCount(entries []*mtgv1.CollectionEntry) int32 {
	var n int32
	for _, e := range entries {
		n = addSaturate(n, e.Quantity)
	}
	return n
}

// addSaturate adds two counts and stops at the int32 maximum. A stored
// collection can hold entries the current cap did not bound.
func addSaturate(a, b int32) int32 {
	sum := int64(a) + int64(b)
	if sum > math.MaxInt32 {
		return math.MaxInt32
	}
	return int32(sum)
}

// SortEntries orders entries by name, set, number, finish, and
// condition. Two rows of one printing differ by finish or condition, so
// both join the key, and the sort is stable for the rest. One upload
// then stores one order every time. ContentHash hashes the upload
// bytes, not this order.
func SortEntries(entries []*mtgv1.CollectionEntry) {
	sort.SliceStable(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		if a.Name != b.Name {
			return a.Name < b.Name
		}
		if a.SetCode != b.SetCode {
			return a.SetCode < b.SetCode
		}
		if a.CollectorNumber != b.CollectorNumber {
			return a.CollectorNumber < b.CollectorNumber
		}
		if a.Finish != b.Finish {
			return a.Finish < b.Finish
		}
		return a.Condition < b.Condition
	})
}
