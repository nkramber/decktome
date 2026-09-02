package profile

import (
	"context"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

// TestTutorSlugsExist checks the tag slugs the profile reads against a
// real snapshot, the way the theme slugs are checked (make
// themes-check). It skips without CARDS_SNAPSHOT_DIR.
func TestTutorSlugsExist(t *testing.T) {
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		t.Skip("set CARDS_SNAPSHOT_DIR to check the tag slugs against a snapshot")
	}
	idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, quietLog())
	if err != nil || idx == nil {
		t.Fatalf("snapshot: %v", err)
	}
	tags := idx.Tags()
	for _, slug := range []string{tutorSlug, tutorLandSlug} {
		if !tags.Has(slug) {
			t.Errorf("the snapshot has no tag %q", slug)
		}
	}
	// The tutor set less the land searches must still hold the tutors
	// the bracket text means, and drop the ramp.
	if n := len(tags.Resolve(tutorSlug)) - len(tags.Resolve(tutorLandSlug)); n < 100 {
		t.Errorf("tutor less tutor-land resolves to %d cards, want 100 or more", n)
	}
}

// quietLog is a logger that writes nothing, for a snapshot load.
func quietLog() *slog.Logger { return slog.New(slog.NewTextHandler(io.Discard, nil)) }
