package questions

import (
	"context"
	"io"
	"log/slog"
	"os"
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
)

func spreadCand(name string, ids ...mtgv1.Color) candidates.Candidate {
	return candidates.Candidate{Card: &mtgv1.Card{Name: name, ColorIdentity: ids}}
}

func candNames(cs []candidates.Candidate) []string {
	var out []string
	for _, c := range cs {
		out = append(out, c.DisplayName())
	}
	return out
}

// TestSpreadColorsKeepsBothKinds is D-1151. Session
// wBrsxouAndrjDXEJ8dDw skipped the colors, and both offers held
// mono-colored names alone. Each offer of three now holds a
// mono-colored and a multicolor name.
func TestSpreadColorsKeepsBothKinds(t *testing.T) {
	W, U, B, R := mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B, mtgv1.Color_COLOR_R
	monoOnly := []candidates.Candidate{
		spreadCand("A", W), spreadCand("B", U), spreadCand("C", B),
		spreadCand("D", R), spreadCand("E", W, U), spreadCand("F", B, R),
	}
	got := candNames(spreadColors(monoOnly, 3))
	want := []string{"A", "B", "E", "C", "D", "F"}
	if !slices.Equal(got, want) {
		t.Errorf("mono-colored first three: got %v, want %v", got, want)
	}

	multiOnly := []candidates.Candidate{
		spreadCand("A", W, U), spreadCand("B", U, B), spreadCand("C", B, R), spreadCand("D", R),
	}
	got = candNames(spreadColors(multiOnly, 3))
	want = []string{"A", "B", "D", "C"}
	if !slices.Equal(got, want) {
		t.Errorf("multicolor first three: got %v, want %v", got, want)
	}

	// A pair counts the union of its two identities, so two mono-colored
	// partners of different colors make a multicolor choice (D-154).
	pair := spreadCand("P", W)
	pair.Partner = &mtgv1.Card{Name: "Q", ColorIdentity: []mtgv1.Color{U}}
	withPair := []candidates.Candidate{spreadCand("A", W), pair, spreadCand("C", B), spreadCand("D", W, B)}
	got = candNames(spreadColors(withPair, 3))
	want = []string{"A", "P + Q", "C", "D"}
	if !slices.Equal(got, want) {
		t.Errorf("a pair is multicolor: got %v, want %v", got, want)
	}

	// A list with one kind alone stays as it is.
	same := []candidates.Candidate{spreadCand("A", W), spreadCand("B", U), spreadCand("C", B), spreadCand("D", R)}
	if got := candNames(spreadColors(same, 3)); !slices.Equal(got, []string{"A", "B", "C", "D"}) {
		t.Errorf("a list of one kind changed: %v", got)
	}
}

// TestCommandersSpreadWithNoColors runs the offer of D-1151 on the card
// snapshot. With no theme and no colors, each offer of three holds a
// mono-colored and a multicolor name. Named colors keep the old offer.
//
//	CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall go test ./internal/questions -run TestCommandersSpread
func TestCommandersSpreadWithNoColors(t *testing.T) {
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		t.Skip("set CARDS_SNAPSHOT_DIR to run the snapshot tests")
	}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, quiet)
	if err != nil {
		t.Fatalf("snapshot: %v", err)
	}
	b, err := candidates.New()
	if err != nil {
		t.Fatal(err)
	}
	h := &CandidateHints{Index: idx, Builder: b, Bracket: 3}
	var seen []string
	for round := 0; round < 3; round++ {
		offer := h.Commanders("", seen)
		if len(offer) != 3 {
			t.Fatalf("round %d offered %v, want three names", round, offer)
		}
		mono, multi := 0, 0
		for _, n := range offer {
			var ids []mtgv1.Color
			for _, part := range strings.Split(n, " + ") {
				c, ok := idx.ByName(part)
				if !ok {
					t.Fatalf("round %d: %q is not in the index", round, part)
				}
				ids = append(ids, c.GetColorIdentity()...)
			}
			if len(candidates.ColorSet(ids)) == 1 {
				mono++
			} else {
				multi++
			}
		}
		if mono == 0 || multi == 0 {
			t.Errorf("round %d offered %v: %d mono-colored, %d multicolor", round, offer, mono, multi)
		}
		seen = append(seen, offer...)
	}
}
