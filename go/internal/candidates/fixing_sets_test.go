package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// fixingSetFixture holds the land shapes of the Hobbit family of the
// snapshot of 2026-09-04: a dual land for each of three color pairs, two
// lands that fetch a basic land, a land of one color, and a dual land
// outside the family.
func fixingSetFixture(t *testing.T) *cards.Index {
	t.Helper()
	return fixture(t, []tc{
		{id: "gu", name: "Elvenking's Halls", typeLine: "Land", text: "{T}: Add {G} or {U}.",
			identity: []mtgv1.Color{G, U}, produced: []mtgv1.Color{G, U}, sets: []string{"hob"}},
		{id: "rw", name: "Iron Hills", typeLine: "Land", text: "{T}: Add {R} or {W}.",
			identity: []mtgv1.Color{R, W}, produced: []mtgv1.Color{R, W}, sets: []string{"hob"}},
		{id: "br", name: "Goblin-town", typeLine: "Land", text: "{T}: Add {B} or {R}.",
			identity: []mtgv1.Color{B, R}, produced: []mtgv1.Color{B, R}, sets: []string{"hob"}},
		{id: "fetch", name: "Hobbit Hole", typeLine: "Land",
			text: "{T}, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle.",
			sets: []string{"hob"}},
		{id: "mono", name: "Rivendell", typeLine: "Legendary Land", text: "{T}: Add {U}.",
			identity: []mtgv1.Color{U}, produced: []mtgv1.Color{U}, sets: []string{"hoc"}},
		{id: "outside", name: "Hallowed Fountain", typeLine: "Land", text: "{T}: Add {W} or {U}.",
			identity: []mtgv1.Color{W, U}, produced: []mtgv1.Color{W, U}, sets: []string{"m19"}},
		{id: "forest", name: "Forest", typeLine: "Basic Land — Forest", text: "({T}: Add {G}.)",
			identity: []mtgv1.Color{G}, produced: []mtgv1.Color{G}, sets: []string{"hob"}},
	})
}

// TestCountFixingInSetsReadsTheDeckColors is D-1227. The mana row of
// D-382 counted ramp cards and nonbasic lands, and no fixing land. A
// four-color commander of the Hobbit family passed that count with 5
// fixing lands against a floor of 17.
func TestCountFixingInSetsReadsTheDeckColors(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixingSetFixture(t)
	for _, c := range []struct {
		name   string
		format mtgv1.FormatId
		colors []mtgv1.Color
		want   int
	}{
		// Aragorn, the Uniter: the green-blue and red-white lands, and the
		// fetch land. Goblin-town is outside the identity.
		{"four colors", cmdr, []mtgv1.Color{W, U, R, G}, 3},
		// Gríma, Saruman's Footman: the fetch land alone.
		{"blue and black", cmdr, []mtgv1.Color{U, B}, 1},
		// A 60-card deck reads the same lands.
		{"a 60-card deck", mtgv1.FormatId_FORMAT_ID_STANDARD, []mtgv1.Color{B, R}, 2},
		// A deck of one color needs no fixing, and no color names none.
		{"one color", cmdr, []mtgv1.Color{U}, 0},
		{"no colors", cmdr, nil, 0},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := b.CountFixingInSets(idx, Request{Format: c.format, Colors: c.colors, SetCodes: []string{"hob", "hoc"}})
			if got != c.want {
				t.Errorf("CountFixingInSets = %d, want %d", got, c.want)
			}
		})
	}
	if n := b.CountFixingInSets(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W, U}}); n != 0 {
		t.Errorf("CountFixingInSets with no set limit = %d, want 0", n)
	}
}
