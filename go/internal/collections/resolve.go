package collections

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// Resolve joins parsed rows against the card index.
// Join order (F-2): Scryfall ID, then set code plus collector number,
// then exact name. Non-English rows are reported, never imported (D-23).
func Resolve(rows []Row, idx *cards.Index) ([]*mtgv1.CollectionEntry, []*mtgv1.UnresolvedRow) {
	var entries []*mtgv1.CollectionEntry
	var bad []*mtgv1.UnresolvedRow
	for _, row := range rows {
		if row.Language != "" && !strings.EqualFold(row.Language, "en") {
			bad = append(bad, unresolved(row.Line, row.Raw, mtgv1.UnresolvedReason_UNRESOLVED_REASON_NON_ENGLISH))
			continue
		}
		card, ok := resolveOne(row, idx)
		if !ok {
			bad = append(bad, unresolved(row.Line, row.Raw, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_CARD))
			continue
		}
		entries = append(entries, &mtgv1.CollectionEntry{
			ScryfallId:      row.ScryfallID,
			OracleId:        card.OracleId,
			Name:            card.Name,
			SetCode:         row.SetCode,
			CollectorNumber: row.Collector,
			Quantity:        int32(row.Quantity),
			Finish:          row.Finish,
			Condition:       row.Condition,
			Rarity:          row.Rarity,
		})
	}
	return entries, bad
}

func resolveOne(row Row, idx *cards.Index) (*mtgv1.Card, bool) {
	if row.ScryfallID != "" {
		if c, ok := idx.ByPrintingID(row.ScryfallID); ok {
			return c, true
		}
	}
	if row.SetCode != "" && row.Collector != "" {
		if c, ok := idx.BySetCollector(row.SetCode, row.Collector); ok {
			return c, true
		}
	}
	if row.Name != "" {
		if c, ok := idx.ByName(row.Name); ok {
			return c, true
		}
	}
	return nil, false
}

// ContentHash detects an identical re-upload (D-16).
func ContentHash(content []byte) string {
	sum := sha256.Sum256(content)
	return hex.EncodeToString(sum[:])
}

// OracleCounts sums owned copies per Oracle id across printings.
func OracleCounts(entries []*mtgv1.CollectionEntry) map[string]int32 {
	out := map[string]int32{}
	for _, e := range entries {
		out[e.OracleId] += e.Quantity
	}
	return out
}

// CardCount sums every copy.
func CardCount(entries []*mtgv1.CollectionEntry) int32 {
	var n int32
	for _, e := range entries {
		n += e.Quantity
	}
	return n
}

// SortEntries orders entries by name, then set, then number. Stable
// storage order keeps the content hash meaningful across imports.
func SortEntries(entries []*mtgv1.CollectionEntry) {
	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		if a.Name != b.Name {
			return a.Name < b.Name
		}
		if a.SetCode != b.SetCode {
			return a.SetCode < b.SetCode
		}
		return a.CollectorNumber < b.CollectorNumber
	})
}
