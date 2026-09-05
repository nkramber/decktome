package precons

import (
	"context"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
)

func card(name, oracle, scryfall string, n int) meta.PreconCard {
	return meta.PreconCard{Name: name, Count: n, OracleID: oracle, ScryfallID: scryfall, SetCode: "tst", Number: scryfall}
}

// tableRows is a small table: a Commander deck, its Collector's
// Edition with the same cards on other printings, two theme decks that
// share a name and hold different cards, and a one-word product.
func tableRows() []meta.Precon {
	return []meta.Precon{
		{Name: "Avengers Assemble", Code: "MSC", Type: "Commander Deck", ReleaseDate: "2026-06-26",
			Commanders: []meta.PreconCard{card("Captain America, Team Leader", "o-cap", "s-cap", 1)},
			Cards: []meta.PreconCard{
				card("Sol Ring", "o-sol", "s-sol", 1), card("Avenge", "o-avenge", "s-avenge", 1),
				card("Plains", "o-plains", "s-plains-a", 10), card("Plains", "o-plains", "s-plains-b", 8),
			}},
		{Name: "Avengers Assemble Collector's Edition", Code: "MSC", Type: "Commander Deck", ReleaseDate: "2026-06-26",
			Commanders: []meta.PreconCard{card("Captain America, Team Leader", "o-cap", "s-cap-ce", 1)},
			Cards: []meta.PreconCard{
				card("Sol Ring", "o-sol", "s-sol-ce", 1), card("Avenge", "o-avenge", "s-avenge-ce", 1),
				card("Plains", "o-plains", "s-plains-ce", 18),
			}},
		{Name: "Deck A", Code: "AAA", Type: "Theme Deck", ReleaseDate: "1999-01-01",
			Cards: []meta.PreconCard{card("Grizzly Bears", "o-bears", "s-bears", 4)}},
		{Name: "Deck A", Code: "BBB", Type: "Theme Deck", ReleaseDate: "2000-01-01",
			Cards: []meta.PreconCard{card("Llanowar Elves", "o-elves", "s-elves", 4)}},
		{Name: "Migraine", Code: "TMP", Type: "Theme Deck", ReleaseDate: "1997-10-14",
			Cards: []meta.PreconCard{card("Hymn to Tourach", "o-hymn", "s-hymn", 4)}},
	}
}

func TestNewTableIndexesTheRows(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	if tbl.Len() != 5 {
		t.Fatalf("len = %d, want 5", tbl.Len())
	}
	p, ok := tbl.Get("AvengersAssemble_MSC")
	if !ok {
		t.Fatal("the key of the deck is the MTGJSON file name")
	}
	counts := p.Counts()
	if counts["o-plains"] != 18 || counts["o-cap"] != 1 || counts["o-sol"] != 1 {
		t.Errorf("counts = %v: two printings of Plains merge, and the commander counts", counts)
	}
	ce, _ := tbl.Get("AvengersAssembleCollector'sEdition_MSC")
	if !p.SameCards(ce) {
		t.Error("a deck and its Collector's Edition hold the same cards")
	}
	a, _ := tbl.Get("DeckA_AAA")
	if a.SameCards(p) {
		t.Error("two different products must not read as the same cards")
	}
}

// TestResolveReadsTheProductName is D-496: the classifier hands over the
// reader's words for the product, and the table settles them.
func TestResolveReadsTheProductName(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	for _, tc := range []struct {
		phrase string
		want   []string
	}{
		{"Avengers Assemble", []string{"Avengers Assemble"}},
		{"my Avengers Assemble precon", []string{"Avengers Assemble"}},
		{"the avengers assemble commander deck", []string{"Avengers Assemble"}},
		{"Avengers Assemble Collector's Edition", []string{"Avengers Assemble Collector's Edition"}},
		{"Migraine", []string{"Migraine"}},
		{"my Migraine deck", []string{"Migraine"}},
	} {
		m := tbl.Resolve(tc.phrase)
		if !m.OK() {
			t.Errorf("%q: no product, options %v", tc.phrase, m.Options)
			continue
		}
		if got := m.Names(); !slices.Equal(got, tc.want) {
			t.Errorf("%q: names = %v, want %v", tc.phrase, got, tc.want)
		}
	}
}

// TestResolveAsksAboutTwoProductsWithOneName is the "Deck A" case: MTGJSON
// holds 21 names twice or more, with different cards each time.
func TestResolveAsksAboutTwoProductsWithOneName(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	m := tbl.Resolve("Deck A")
	if m.OK() {
		t.Fatalf("two products with different cards must not resolve: %v", m.Names())
	}
	if !slices.Equal(m.Options, []string{"Deck A (AAA)", "Deck A (BBB)"}) {
		t.Errorf("options = %v", m.Options)
	}
}

// TestResolveOffersNearNamesForAnUnknownPhrase feeds the row that asks
// which product the reader means.
func TestResolveOffersNearNamesForAnUnknownPhrase(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	m := tbl.Resolve("the avengers deck")
	if m.OK() {
		t.Fatalf("a partial name must not resolve: %v", m.Names())
	}
	if !slices.Equal(m.Options, []string{"Avengers Assemble", "Avengers Assemble Collector's Edition"}) {
		t.Errorf("options = %v", m.Options)
	}
	if m := tbl.Resolve("my precon"); m.OK() || len(m.Options) != 0 {
		t.Errorf("filler words alone name nothing: %v %v", m.Names(), m.Options)
	}
	if m := tbl.Resolve(""); m.OK() {
		t.Error("an empty phrase names nothing")
	}
}

// TestOwnedWholeReadsEveryPrinting is D-408: the collection holds every
// printing of the product with its count, and 99 of 100 is not whole.
func TestOwnedWholeReadsEveryPrinting(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	whole := map[string]int32{"s-cap": 1, "s-sol": 1, "s-avenge": 1, "s-plains-a": 10, "s-plains-b": 8}
	owned := tbl.Owned(whole, nil)
	if len(owned) != 1 || owned[0].Name != "Avengers Assemble" {
		t.Fatalf("owned = %v, want Avengers Assemble alone", names(owned))
	}
	short := map[string]int32{"s-cap": 1, "s-sol": 1, "s-avenge": 1, "s-plains-a": 10, "s-plains-b": 7}
	if got := tbl.Owned(short, nil); len(got) != 0 {
		t.Errorf("99 of 100 reads as owned: %v", names(got))
	}
	// A Sol Ring of another set does not make the reader the owner.
	other := map[string]int32{"s-cap": 1, "s-sol-other": 1, "s-avenge": 1, "s-plains-a": 10, "s-plains-b": 8}
	if got := tbl.Owned(other, nil); len(got) != 0 {
		t.Errorf("another printing reads as the product's: %v", names(got))
	}
	if got := tbl.Owned(nil, nil); got != nil {
		t.Errorf("no collection owns nothing: %v", names(got))
	}
}

// TestOwnedWholeIgnoresBasicLands is D-523: a binder that holds every
// nonbasic printing owns the product, whatever basic lands it lacks. A
// nonbasic card still counts, and a test that calls every card basic
// owns nothing.
func TestOwnedWholeIgnoresBasicLands(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	basic := func(oracleID string) bool { return oracleID == "o-plains" }
	noBasics := map[string]int32{"s-cap": 1, "s-sol": 1, "s-avenge": 1}
	if got := tbl.Owned(noBasics, basic); len(got) != 1 || got[0].Name != "Avengers Assemble" {
		t.Errorf("a binder with no basic lands owns the product: %v", names(got))
	}
	if got := tbl.Owned(noBasics, nil); len(got) != 0 {
		t.Errorf("with no basic test every printing counts: %v", names(got))
	}
	noSol := map[string]int32{"s-cap": 1, "s-avenge": 1, "s-plains-a": 10, "s-plains-b": 8}
	if got := tbl.Owned(noSol, basic); len(got) != 0 {
		t.Errorf("a missing nonbasic card still reads as not owned: %v", names(got))
	}
	all := func(string) bool { return true }
	if got := tbl.Owned(noBasics, all); len(got) != 0 {
		t.Errorf("a product with no nonbasic printing is never owned: %v", names(got))
	}
}

// TestExcludeLeavesTheSurplusCopies is D-408: two copies with one in the
// precon leave one usable, and a card with no copy left leaves the pool.
func TestExcludeLeavesTheSurplusCopies(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	p, _ := tbl.Get("AvengersAssemble_MSC")
	owned := map[string]int32{"o-sol": 2, "o-cap": 1, "o-avenge": 1, "o-plains": 30, "o-bears": 4}
	isBasic := func(id string) bool { return id == "o-plains" }
	reduced, excluded := Exclude([]*Product{p}, owned, isBasic)
	if reduced["o-sol"] != 1 {
		t.Errorf("Sol Ring left = %d, want 1", reduced["o-sol"])
	}
	if _, ok := reduced["o-cap"]; ok {
		t.Error("the commander has no copy left, so it leaves the counts")
	}
	if reduced["o-plains"] != 30 || reduced["o-bears"] != 4 {
		t.Errorf("a basic land and a card outside the product keep their counts: %v", reduced)
	}
	if !slices.Equal(excluded, []string{"o-avenge", "o-cap"}) {
		t.Errorf("excluded = %v, want the two cards with no copy left, sorted", excluded)
	}
	if owned["o-sol"] != 2 {
		t.Error("Exclude must not change the caller's map")
	}
}

// TestExcludeCountsTwinProductsOnce: a deck and its Collector's Edition
// both answer a phrase, and the reader owns one product, not two.
func TestExcludeCountsTwinProductsOnce(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	m := tbl.Resolve("Avengers Assemble")
	ce, _ := tbl.Get("AvengersAssembleCollector'sEdition_MSC")
	products := append([]*Product(nil), m.Products...)
	products = append(products, ce)
	reduced, excluded := Exclude(products, map[string]int32{"o-sol": 2}, nil)
	if reduced["o-sol"] != 1 {
		t.Errorf("Sol Ring left = %d, want 1: the twins count once", reduced["o-sol"])
	}
	if !slices.Contains(excluded, "o-plains") {
		t.Error("with no isBasic, every card with no copy left is excluded")
	}
}

// TestExcludeWithNoCollectionExcludesEveryCard is the any-card pool.
func TestExcludeWithNoCollectionExcludesEveryCard(t *testing.T) {
	tbl := NewTable("v1", tableRows())
	p, _ := tbl.Get("Migraine_TMP")
	reduced, excluded := Exclude([]*Product{p}, nil, nil)
	if len(reduced) != 0 || !slices.Equal(excluded, []string{"o-hymn"}) {
		t.Errorf("reduced %v excluded %v", reduced, excluded)
	}
}

func names(ps []*Product) []string {
	var out []string
	for _, p := range ps {
		out = append(out, p.Name)
	}
	return out
}

// loadTable reads the stored precon table of the local meta store, which
// is the parent folder of the card snapshot.
func loadTable(t *testing.T) *Table {
	t.Helper()
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		t.Skip("set CARDS_SNAPSHOT_DIR to read the precon table")
	}
	store := meta.DirObjects{Root: filepath.Dir(dir)}
	ctx := context.Background()
	version, err := meta.LatestPreconsVersion(ctx, store)
	if err != nil {
		t.Fatalf("precon table version: %v", err)
	}
	if version == "" {
		t.Skip("the local meta store holds no precon table, run make meta-refresh")
	}
	rows, err := meta.ReadPrecons(ctx, store, version)
	if err != nil {
		t.Fatalf("precon table: %v", err)
	}
	return NewTable(version, rows)
}

// TestEmbeddedListsMatchTheTable is the cross-check of D-498. The nine
// embedded lists are the owner-verified ground truth, and every one must
// name the same cards as its MTGJSON row. A difference is a defect in one
// of the two sources, and the test names the cards.
func TestEmbeddedListsMatchTheTable(t *testing.T) {
	idx := loadIndex(t)
	set, err := Load(idx)
	if err != nil {
		t.Fatalf("precons: %v", err)
	}
	tbl := loadTable(t)
	nameOf := func(id string) string {
		if c, ok := idx.ByOracleID(id); ok {
			return c.GetName()
		}
		return id
	}
	for _, p := range set.All() {
		m := tbl.Resolve(p.Name)
		var row *Product
		if m.OK() {
			row = m.Products[0]
		} else {
			// The file name is a slug and not a verified product name for
			// two of the nine, so the cards find the row. A product the
			// table lacks is a finding for the owner, not a card error.
			row = byCards(tbl, p.OracleIDs)
			if row == nil {
				t.Errorf("%s: the table holds no product named %q and none with its cards", p.Slug, p.Name)
				continue
			}
			t.Logf("%s: the table names it %q (%s), not %q", p.Slug, row.Name, row.Code, p.Name)
		}
		want := map[string]bool{}
		for id := range row.counts {
			want[id] = true
		}
		var missing, extra []string
		for _, id := range p.OracleIDs {
			if !want[id] {
				extra = append(extra, id)
			}
			delete(want, id)
		}
		for id := range want {
			missing = append(missing, id)
		}
		for i := range missing {
			missing[i] = nameOf(missing[i])
		}
		for i := range extra {
			extra[i] = nameOf(extra[i])
		}
		sort.Strings(missing)
		sort.Strings(extra)
		if len(missing) > 0 || len(extra) > 0 {
			t.Errorf("%s against %s (%s): %d cards only in the table %v, %d cards only in the file %v",
				p.Slug, row.Name, row.Code, len(missing), missing, len(extra), extra)
		}
		if !strings.EqualFold(row.Name, p.Name) {
			t.Logf("%s: the file says %q and the table %q", p.Slug, p.Name, row.Name)
		}
	}
}

// byCards finds the product that shares the most cards with a list, when
// it shares at least nine in ten of them.
func byCards(tbl *Table, oracleIDs []string) *Product {
	var best *Product
	bestShared := 0
	for _, row := range tbl.All() {
		shared := 0
		for _, id := range oracleIDs {
			if _, ok := row.counts[id]; ok {
				shared++
			}
		}
		if shared > bestShared {
			best, bestShared = row, shared
		}
	}
	if best == nil || bestShared*10 < len(oracleIDs)*9 {
		return nil
	}
	return best
}
