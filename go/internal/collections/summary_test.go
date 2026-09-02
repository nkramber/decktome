package collections

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// fakeCards answers the card types of a card.
type fakeCards map[string][]string

func (f fakeCards) ByOracleID(id string) (*mtgv1.Card, bool) {
	types, ok := f[id]
	if !ok {
		return nil, false
	}
	return &mtgv1.Card{OracleId: id, CardTypes: types}, true
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
	cards := fakeCards{"o-bolt": {"Instant"}, "o-jace": {"Planeswalker"}}
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
	if got.GetByType()["Instant"] != 6 || got.GetByType()["Planeswalker"] != 1 {
		t.Errorf("by type = %v, want 6 instants and 1 planeswalker", got.GetByType())
	}
}

// TestSummarizeCountsACardOfTwoTypesUnderEach: the type filter offers
// both, so an artifact creature answers to Artifact and to Creature.
func TestSummarizeCountsACardOfTwoTypesUnderEach(t *testing.T) {
	got := Summarize([]*mtgv1.CollectionEntry{entry("p1", "o-golem", "m10", "rare", 2)},
		fakeCards{"o-golem": {"Artifact", "Creature"}})
	if got.GetByType()["Artifact"] != 2 || got.GetByType()["Creature"] != 2 {
		t.Errorf("by type = %v, want 2 under each", got.GetByType())
	}
}

// TestSummarizeListsEverySet is D-398. The binder's set filter offers
// each set the collection holds, largest first, and a tie reads the
// same on every run.
func TestSummarizeListsEverySet(t *testing.T) {
	var entries []*mtgv1.CollectionEntry
	for i := 0; i < 30; i++ {
		entries = append(entries, entry("p", "o", string(rune('a'+i)), "common", int32(i+1)))
	}
	entries = append(entries, entry("p", "o", "aaa", "common", 30))
	got := Summarize(entries, nil)
	if len(got.GetSets()) != 31 {
		t.Fatalf("sets = %d, want every one of 31", len(got.GetSets()))
	}
	// The set code breaks the tie at 30 between "aaa" and the last rune.
	if got.GetSets()[0].GetSetCode() != "aaa" || got.GetSets()[0].GetCount() != 30 {
		t.Errorf("first set = %v, want aaa at 30", got.GetSets()[0])
	}
	if got.GetSets()[30].GetCount() != 1 {
		t.Errorf("last set = %v, want the smallest", got.GetSets()[30])
	}
}

// TestSummarizeWithNoCardIndex: every count but the types reads the
// entries alone, because the rarity and the set name ride on every row.
func TestSummarizeWithNoCardIndex(t *testing.T) {
	got := Summarize([]*mtgv1.CollectionEntry{entry("p1", "o1", "lea", "rare", 2)}, nil)
	if got.GetRowCount() != 1 || got.GetUniqueCards() != 1 || got.GetByRarity()["rare"] != 2 {
		t.Errorf("summary = %v", got)
	}
	if len(got.GetByType()) != 0 {
		t.Errorf("by type = %v, want none with no card index", got.GetByType())
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
