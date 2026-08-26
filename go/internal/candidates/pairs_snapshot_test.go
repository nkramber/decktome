package candidates

import (
	"context"
	"io"
	"log/slog"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// snapshotIndex loads the local card snapshot. Every card fact in this
// file is checked against it and never against memory.
func snapshotIndex(t *testing.T) *cards.Index {
	t.Helper()
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	idx, err := cards.LoadIndex(context.Background(),
		cards.DirStore{Root: "../../../.local/gcs/mtg-local-cards/scryfall"}, quiet)
	if err != nil || idx == nil {
		t.Skipf("no card snapshot: %v", err)
	}
	return idx
}

// TestPairsReachFourColors is D-154. WUBR, WBRG, and UBRG hold exactly
// one legal single commander each, so a four-color request could offer no
// choice at all. A pair carries the union of two color identities.
func TestPairsReachFourColors(t *testing.T) {
	idx := snapshotIndex(t)
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	four := []mtgv1.Color{
		mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B,
		mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G,
	}
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "value", Colors: four}
	pool, err := b.CommanderPool(idx, req)
	if err != nil {
		t.Fatalf("commander pool: %v", err)
	}
	var pairs, singles int
	for _, c := range pool {
		if c.Partner != nil {
			pairs++
			// Every pair must hold all four colors and no fifth.
			union := append(append([]mtgv1.Color(nil), c.Card.ColorIdentity...), c.Partner.ColorIdentity...)
			if !IdentityMatches(union, four) {
				t.Errorf("%s does not match the four colors: %v + %v",
					c.DisplayName(), c.Card.ColorIdentity, c.Partner.ColorIdentity)
			}
		} else {
			singles++
		}
	}
	t.Logf("UBRG pool: %d singles, %d pairs", singles, pairs)
	if pairs == 0 {
		t.Error("a four-color request found no pair, and one legal single commander exists")
	}
	if len(pool) < 2 {
		t.Errorf("a four-color request offers %d commanders, and the pick row holds three", len(pool))
	}
}

// TestBackgroundPairMatchesTheColors is probe 73 of gate run 16. The user
// asked for a Background pair and red-white. A mono-red leader with
// "choose a Background" plus a white Background is a red-white deck.
//
// WantBackground is what makes the Background appear. A Background holds
// no theme signal of its own, so a pair that holds one loses on score to
// a pair of two themed legends. The user named a Background, so a pair
// without one answers a different question (D-154).
func TestBackgroundPairMatchesTheColors(t *testing.T) {
	idx := snapshotIndex(t)
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	rw := []mtgv1.Color{mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_W}
	req := Request{
		Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "aggro",
		Colors: rw, WantPair: true, WantBackground: true,
	}
	pool, err := b.CommanderPool(idx, req)
	if err != nil {
		t.Fatalf("commander pool: %v", err)
	}
	var withBackground int
	for _, c := range pool {
		if c.Partner == nil {
			continue
		}
		union := append(append([]mtgv1.Color(nil), c.Card.ColorIdentity...), c.Partner.ColorIdentity...)
		if !IdentityMatches(union, rw) {
			t.Errorf("%s is not red-white: %v + %v",
				c.DisplayName(), c.Card.ColorIdentity, c.Partner.ColorIdentity)
		}
		// A Background never leads. It is not a commander on its own.
		if c.Card.GetIsBackground() {
			t.Errorf("a Background leads the pair: %s", c.DisplayName())
		}
		if c.Partner.GetIsBackground() {
			withBackground++
		}
		if !strings.Contains(c.DisplayName(), " + ") {
			t.Errorf("a pair does not read as a pair: %q", c.DisplayName())
		}
	}
	if withBackground == 0 {
		t.Error("the user asked for a Background pair and the pool offered none")
	}
	t.Logf("red-white pool: %d entries, %d with a Background", len(pool), withBackground)
}

// TestPairsStayOutOfTheDefaultPool keeps D-148 as it stands for the
// common case. One, two, and three colors each hold 47 to 322 single
// commanders, so a pair earns no place there unless the user asks.
func TestPairsStayOutOfTheDefaultPool(t *testing.T) {
	idx := snapshotIndex(t)
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	two := []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_B}
	req := Request{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER, Theme: "lifegain", Colors: two}
	pool, err := b.CommanderPool(idx, req)
	if err != nil {
		t.Fatalf("commander pool: %v", err)
	}
	if len(pool) < commanderNames {
		t.Fatalf("a two-color request found %d commanders, too few to judge", len(pool))
	}
	for _, c := range pool[:commanderNames] {
		if c.Partner != nil {
			t.Errorf("a pair reached a deep single pool the user did not ask about: %s", c.DisplayName())
		}
	}
}
