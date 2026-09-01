package collections

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// fakeCards answers the color identity of a card.
type fakeCards map[string][]mtgv1.Color

func (f fakeCards) ByOracleID(id string) (*mtgv1.Card, bool) {
	c, ok := f[id]
	if !ok {
		return nil, false
	}
	return &mtgv1.Card{OracleId: id, ColorIdentity: c}, true
}

func entry(scry, oracle, set, rarity string, q int32, opts ...func(*mtgv1.CollectionEntry)) *mtgv1.CollectionEntry {
	e := &mtgv1.CollectionEntry{
		ScryfallId: scry, OracleId: oracle, SetCode: set, SetName: set + " name",
		Rarity: rarity, Quantity: q,
	}
	for _, o := range opts {
		o(e)
	}
	return e
}

func foil(e *mtgv1.CollectionEntry) { e.Finish = mtgv1.Finish_FINISH_FOIL }

// TestSummarizeCountsCardsAndNotRows is D-392. The head shows cards,
// copies included, and RowCount is the one count of rows.
func TestSummarizeCountsCardsAndNotRows(t *testing.T) {
	entries := []*mtgv1.CollectionEntry{
		entry("p1", "o-bolt", "lea", "common", 4),
		entry("p2", "o-bolt", "m10", "common", 2),
		entry("p3", "o-jace", "wwk", "mythic", 1),
	}
	cards := fakeCards{"o-bolt": {mtgv1.Color_COLOR_R}, "o-jace": {mtgv1.Color_COLOR_U}}
	got := Summarize(entries, cards)

	if got.GetRowCount() != 3 {
		t.Errorf("row count = %d, want 3", got.GetRowCount())
	}
	// Two rows of one card is one unique card.
	if got.GetUniqueCards() != 2 {
		t.Errorf("unique cards = %d, want 2", got.GetUniqueCards())
	}
	if got.GetByRarity()["common"] != 6 || got.GetByRarity()["mythic"] != 1 {
		t.Errorf("by rarity = %v, want 6 common and 1 mythic", got.GetByRarity())
	}
	if got.GetByColor()["COLOR_R"] != 6 || got.GetByColor()["COLOR_U"] != 1 {
		t.Errorf("by color = %v, want 6 red and 1 blue", got.GetByColor())
	}
}

// TestSummarizeRanksTheSets: the head shows the sets the collection
// holds most of, largest first, and a tie reads the same on every run.
func TestSummarizeRanksTheSets(t *testing.T) {
	got := Summarize([]*mtgv1.CollectionEntry{
		entry("p1", "o1", "lea", "common", 1),
		entry("p2", "o2", "m10", "common", 5),
		entry("p3", "o3", "aaa", "common", 1),
	}, nil)
	if len(got.GetTopSets()) != 3 {
		t.Fatalf("top sets = %v", got.GetTopSets())
	}
	if got.GetTopSets()[0].GetSetCode() != "m10" {
		t.Errorf("first set = %q, want the largest", got.GetTopSets()[0].GetSetCode())
	}
	// The set code breaks the tie between lea and aaa.
	if got.GetTopSets()[1].GetSetCode() != "aaa" {
		t.Errorf("second set = %q, want the tie broken by code", got.GetTopSets()[1].GetSetCode())
	}
}

// TestSummarizeWithNoCardIndex: every count but the colors reads the
// entries alone, because the rarity and the set name ride on every row.
func TestSummarizeWithNoCardIndex(t *testing.T) {
	got := Summarize([]*mtgv1.CollectionEntry{entry("p1", "o1", "lea", "rare", 2)}, nil)
	if got.GetRowCount() != 1 || got.GetUniqueCards() != 1 || got.GetByRarity()["rare"] != 2 {
		t.Errorf("summary = %v", got)
	}
	if len(got.GetByColor()) != 0 {
		t.Errorf("by color = %v, want none with no card index", got.GetByColor())
	}
}

// TestSummarizeBoundsTheSetList: a collection holds cards of hundreds of
// sets, and the head shows a short list.
func TestSummarizeBoundsTheSetList(t *testing.T) {
	var entries []*mtgv1.CollectionEntry
	for i := 0; i < TopSets+5; i++ {
		entries = append(entries, entry("p", "o", string(rune('a'+i)), "common", int32(i+1)))
	}
	if got := Summarize(entries, nil); len(got.GetTopSets()) != TopSets {
		t.Errorf("top sets = %d, want %d", len(got.GetTopSets()), TopSets)
	}
}

// TestDiffKeysOnPrintingFinishAndCondition is D-393. That is the row a
// ManaBox export writes, so a foil and a normal copy of one card are two
// rows.
func TestDiffKeysOnPrintingFinishAndCondition(t *testing.T) {
	stored := []*mtgv1.CollectionEntry{
		entry("p1", "o-bolt", "lea", "common", 4),
		entry("p2", "o-jace", "wwk", "mythic", 1),
	}
	uploaded := []*mtgv1.CollectionEntry{
		// The same printing in foil is another row, not a change.
		entry("p1", "o-bolt", "lea", "common", 4, foil),
		entry("p1", "o-bolt", "lea", "common", 2),
		entry("p3", "o-path", "m10", "uncommon", 3),
	}
	got := Diff(stored, uploaded)

	if len(got.GetAdded()) != 2 {
		t.Errorf("added = %d, want the foil row and the new card", len(got.GetAdded()))
	}
	if got.GetAddedCards() != 7 {
		t.Errorf("added cards = %d, want 4 foil and 3 new", got.GetAddedCards())
	}
	if len(got.GetChanged()) != 1 || got.GetChanged()[0].GetFrom() != 4 || got.GetChanged()[0].GetTo() != 2 {
		t.Errorf("changed = %v, want one row from 4 to 2", got.GetChanged())
	}
	// A fall of two moves two cards.
	if got.GetChangedCards() != 2 {
		t.Errorf("changed cards = %d, want 2", got.GetChangedCards())
	}
	if len(got.GetRemoved()) != 1 || got.GetRemovedCards() != 1 {
		t.Errorf("removed = %v (%d cards), want the Jace row", got.GetRemoved(), got.GetRemovedCards())
	}
	if got.GetIdentical() {
		t.Error("a diff with changes is not identical")
	}
}

// TestDiffOfOneFileWithItselfIsEmpty is the PR-18 gate line: the same
// file uploaded twice diffs empty.
func TestDiffOfOneFileWithItselfIsEmpty(t *testing.T) {
	entries := []*mtgv1.CollectionEntry{
		entry("p1", "o-bolt", "lea", "common", 4, foil),
		entry("p2", "o-jace", "wwk", "mythic", 1),
	}
	got := Diff(entries, entries)
	if !got.GetIdentical() {
		t.Fatalf("a file against itself is not identical: %v", got)
	}
	if len(got.GetAdded())+len(got.GetRemoved())+len(got.GetChanged()) != 0 {
		t.Errorf("diff = %v, want nothing", got)
	}
	if got.GetAddedCards()+got.GetRemovedCards()+got.GetChangedCards() != 0 {
		t.Error("an identical diff moves no card")
	}
}

// TestDiffOfAnEmptyCollection: a first upload is every row added.
func TestDiffOfAnEmptyCollection(t *testing.T) {
	uploaded := []*mtgv1.CollectionEntry{entry("p1", "o1", "lea", "common", 3)}
	got := Diff(nil, uploaded)
	if len(got.GetAdded()) != 1 || got.GetAddedCards() != 3 || got.GetIdentical() {
		t.Errorf("diff = %v", got)
	}
	// The other way round is every row removed.
	back := Diff(uploaded, nil)
	if len(back.GetRemoved()) != 1 || back.GetRemovedCards() != 3 {
		t.Errorf("diff = %v", back)
	}
}
