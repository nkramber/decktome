package agentsvc

import (
	"context"
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/precons"
	"github.com/nkramber/decktome/go/internal/questions"
)

// fakeTable hands out one precon table.
type fakeTable struct{ tbl *precons.Table }

func (f fakeTable) Table() *precons.Table { return f.tbl }

// hobbitPrecon is a product over the cards of setIndex: Smaug leads it,
// and it holds Sol Ring, two Hobbit cards, and a Plains. The printing
// ids follow setIndex, so a collection can hold it whole.
func hobbitPrecon() *precons.Table {
	card := func(name, oracle string) meta.PreconCard {
		return meta.PreconCard{Name: name, Count: 1, OracleID: oracle, ScryfallID: oracle + "-p", SetCode: "hob", Number: "1"}
	}
	return precons.NewTable("v1", []meta.Precon{{
		Name: "Hobbit Precon", Code: "HOB", Type: "Commander Deck", ReleaseDate: "2026-08-14",
		Commanders: []meta.PreconCard{card("Smaug the Impenetrable", "o-smaug")},
		Cards: []meta.PreconCard{
			card("Sol Ring", "o-solring"), card("Hobbit Card aa", "o-in-aa"), card("Hobbit Card ba", "o-in-ba"),
			card("Plains", "o-plains"),
		},
	}})
}

// preconServer wires a server with the set index, a fake generator, the
// precon table, and a collection.
func preconServer(t *testing.T, fd *fakeDecks, cols fakeCollections) *Server {
	t.Helper()
	s := setServer(t, fd, 80)
	WithPreconTable(fakeTable{hobbitPrecon()})(s)
	WithCollections(cols)(s)
	return s
}

// TestBuildExcludesThePreconsCards is the gate of PR-24 (D-408): a build
// for a collection that holds the precon whole uses none of its cards.
// The surplus Sol Ring stays, the other three leave the pool, and the
// generator hears which cards are out.
func TestBuildExcludesThePreconsCards(t *testing.T) {
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	owned := map[string]int32{"o-smaug": 1, "o-solring": 2, "o-in-aa": 1, "o-in-ba": 1, "o-karlov": 1}
	s := preconServer(t, fd, fakeCollections{counts: owned})
	session, st := setSession(nil, "Karlov of the Ghost Council")
	session.Slots.PoolRule = mtgv1.PoolRule_POOL_RULE_ANY_CARD
	session.Slots.ExcludePreconKeys = []string{"HobbitPrecon_HOB"}
	if _, err := s.buildDeck(context.Background(), "u1", session, st, owned, nil, nil); err != nil {
		t.Fatal(err)
	}
	if got := fd.got.ExcludedOracleIDs; !slices.Equal(got, []string{"o-in-aa", "o-in-ba", "o-smaug"}) {
		t.Errorf("excluded = %v, want the three cards with no copy to spare", got)
	}
	if fd.got.OracleCounts["o-solring"] != 1 {
		t.Errorf("Sol Ring count = %d, want 1: the surplus copy stays usable", fd.got.OracleCounts["o-solring"])
	}
	if owned["o-solring"] != 2 {
		t.Error("the exclusion must not change the turn's own map")
	}
	for _, name := range fd.got.Pool.Names() {
		c, _ := fd.got.Pool.Card(name)
		if id := c.GetOracleId(); id == "o-in-aa" || id == "o-in-ba" || id == "o-smaug" {
			t.Errorf("the pool holds %s, a card of the excluded precon", name)
		}
	}
	var solRing bool
	for _, name := range fd.got.Pool.Names() {
		if name == "Sol Ring" {
			solRing = true
		}
	}
	if !solRing {
		t.Error("the pool lost Sol Ring, which has a copy to spare")
	}
}

// TestNoTableExcludesNothing: a session that excludes a precon before
// the meta job stored a table builds as before, and the log says so.
func TestNoTableExcludesNothing(t *testing.T) {
	fd := &fakeDecks{res: &generate.Result{Deck: &mtgv1.Deck{Validation: &mtgv1.ValidationResult{}}}}
	s := setServer(t, fd, 80)
	session, st := setSession(nil, "Karlov of the Ghost Council")
	session.Slots.ExcludePreconKeys = []string{"HobbitPrecon_HOB"}
	if _, err := s.buildDeck(context.Background(), "u1", session, st, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if len(fd.got.ExcludedOracleIDs) != 0 {
		t.Errorf("excluded = %v, want none with no table", fd.got.ExcludedOracleIDs)
	}
}

// TestOwnedPreconsReadThePrintings is D-408 through the turn's hints: a
// collection with every printing of the product owns it, and one with 99
// of 100 does not. A binder that lacks the basic lands alone owns it
// (D-523).
func TestOwnedPreconsReadThePrintings(t *testing.T) {
	whole := map[string]int32{"o-smaug-p": 1, "o-solring-p": 1, "o-in-aa-p": 1, "o-in-ba-p": 1, "o-plains-p": 1}
	noBasics := map[string]int32{"o-smaug-p": 1, "o-solring-p": 1, "o-in-aa-p": 1, "o-in-ba-p": 1}
	short := map[string]int32{"o-smaug-p": 1, "o-solring-p": 1, "o-in-aa-p": 1, "o-plains-p": 1}
	for _, tc := range []struct {
		name      string
		printings map[string]int32
		want      int
	}{
		{"whole", whole, 1},
		{"the basic lands missing", noBasics, 1},
		{"one nonbasic printing short", short, 0},
		{"no collection rows", map[string]int32{}, 0},
	} {
		fd := &fakeDecks{}
		s := preconServer(t, fd, fakeCollections{printingCounts: tc.printings})
		session, st := setSession(nil, "")
		session.CollectionId = "c1"
		h := s.hints(session, st, nil)
		s.withPrecons(context.Background(), "u1", session, h)
		owned, ok := h.OwnedPrecons()
		if !ok {
			t.Fatalf("%s: the hints have no table or no collection", tc.name)
		}
		if len(owned) != tc.want {
			t.Errorf("%s: owned = %v, want %d products", tc.name, owned, tc.want)
		}
		m := h.ResolvePrecon("my Hobbit Precon deck")
		if !m.OK || m.Products[0].Key != "HobbitPrecon_HOB" {
			t.Fatalf("%s: the phrase did not resolve: %+v", tc.name, m)
		}
		if partial := len(m.Partial) > 0; partial != (tc.want == 0) {
			t.Errorf("%s: partial = %v, want it when the collection does not hold the product whole", tc.name, m.Partial)
		}
	}
	// No collection on the session: the check is unavailable, and a
	// named product is neither whole nor partial.
	fd := &fakeDecks{}
	s := preconServer(t, fd, fakeCollections{})
	session, st := setSession(nil, "")
	h := s.hints(session, st, nil)
	s.withPrecons(context.Background(), "u1", session, h)
	if _, ok := h.OwnedPrecons(); ok {
		t.Error("with no collection the owned check must answer not ok")
	}
	if m := h.ResolvePrecon("Hobbit Precon"); !m.OK || len(m.Partial) != 0 {
		t.Errorf("with no collection a named product is excluded with no partial note: %+v", m)
	}
}

// TestPreconNotesNameTheProducts is D-390 for the exclusion: the reader
// hears what left, what left although the collection lacks a card of it,
// and when nothing left.
func TestPreconNotesNameTheProducts(t *testing.T) {
	for _, tc := range []struct {
		res  questions.Result
		want []string
	}{
		{questions.Result{PreconsApplied: []string{"Avengers Assemble"}},
			[]string{"I will use no card of Avengers Assemble"}},
		{questions.Result{PreconsApplied: []string{"Avengers Assemble", "Blight Curse"}, PreconsPartial: []string{"Blight Curse"}},
			[]string{"I will use no card of Avengers Assemble and Blight Curse",
				"Your collection does not hold all of Blight Curse, so I excluded its cards anyway"}},
		{questions.Result{PreconsApplied: []string{"A", "B"}, PreconsPartial: []string{"A", "B"}},
			[]string{"I will use no card of A and B",
				"Your collection does not hold all of A and B, so I excluded their cards anyway"}},
		{questions.Result{PreconsNone: true},
			[]string{"Your collection holds no whole precon, so I excluded nothing"}},
		{questions.Result{PreconsUnavailable: true},
			[]string{"I have no precon table loaded yet, so I excluded no precon"}},
		{questions.Result{}, nil},
	} {
		if got := preconNotes(tc.res); !slices.Equal(got, tc.want) {
			t.Errorf("preconNotes(%+v) = %q, want %q", tc.res, got, tc.want)
		}
	}
}
