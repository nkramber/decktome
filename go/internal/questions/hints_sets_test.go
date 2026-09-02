package questions

import (
	"context"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// TestCommandersWithNoThemeInsideTheSets is D-437. Session
// DrNPxSaisYj2QVlUkTvB declined the theme under a Hobbit set limit, the
// pick row went out with no name, and the build chose a commander. An
// empty theme now answers the popular commanders of the sets.
//
//	CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall go test ./internal/questions -run TestCommandersWithNoTheme
func TestCommandersWithNoThemeInsideTheSets(t *testing.T) {
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
	h := &CandidateHints{Index: idx, Builder: b, SetCodes: []string{"hob", "hoc"}}
	names := h.Commanders("", nil)
	if len(names) != 3 {
		t.Fatalf("an empty theme offered %v, want three names", names)
	}
	sets := cards.CodeSet(h.SetCodes)
	for _, n := range names {
		c, ok := idx.ByName(n)
		if !ok || !cards.InSets(c, sets) {
			t.Errorf("%q is not a commander of the named sets", n)
		}
	}
	// A second offer under other sets is not the cached first one.
	h.UseSets([]string{"ltr"})
	other := h.Commanders("", nil)
	if len(other) == 3 && other[0] == names[0] && other[1] == names[1] {
		t.Errorf("the offer ignored a new set limit: %v", other)
	}
}
