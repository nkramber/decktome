package precons

import (
	"context"
	"io"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

func loadSet(t *testing.T) *Set {
	t.Helper()
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		t.Skip("set CARDS_SNAPSHOT_DIR to resolve the precon lists")
	}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	idx, err := cards.LoadIndex(context.Background(), cards.DirStore{Root: dir}, quiet)
	if err != nil {
		t.Fatalf("cards: %v", err)
	}
	s, err := Load(idx)
	if err != nil {
		t.Fatalf("precons: %v", err)
	}
	return s
}

// TestEveryPreconResolves is the bar for the share rule of D-218. A
// precon the card index can not fully answer gives a wrong share, so a
// list with any unresolved row is a defect and not a warning.
func TestEveryPreconResolves(t *testing.T) {
	s := loadSet(t)
	if len(s.All()) == 0 {
		t.Fatal("no precon decklists are embedded")
	}
	for _, p := range s.All() {
		if p.Unresolved != 0 {
			t.Errorf("%s: %d rows did not resolve", p.Slug, p.Unresolved)
		}
		// A Commander precon is 100 cards, the commander included.
		if p.Cards != 100 {
			t.Errorf("%s: the list holds %d cards, want 100", p.Slug, p.Cards)
		}
		// Singleton, so the distinct count is the card count, basics
		// excepted. A list far under that has lost cards.
		if len(p.OracleIDs) < 60 {
			t.Errorf("%s: only %d distinct cards resolved", p.Slug, len(p.OracleIDs))
		}
	}
}

// TestFindReadsTheUsersWords is D-247. The classifier reports a card name
// for "upgrade my Avengers Assemble precon", not a product, so the words
// are what name the precon (D-240).
func TestFindReadsTheUsersWords(t *testing.T) {
	s := loadSet(t)
	for _, tc := range []struct{ words, want string }{
		{"upgrade my avengers assemble precon", "avengers-assemble"},
		{"I want to improve my Goblin Storm deck", "goblin-storm"},
		{"AVENGERS ASSEMBLE, but better", "avengers-assemble"},
		// The product names, not the file slugs (G-12).
		{"upgrade my riders of rohan precon", "lotr-riders-of-rohan"},
		{"make my Tricky Terrain deck better", "tricky-terrain-collectors-edition"},
		// A short phrase whose every word is in the name.
		{"riders of rohan", "lotr-riders-of-rohan"},
		{"rohan riders", "lotr-riders-of-rohan"},
	} {
		got, ok := s.Find(tc.words)
		if !ok || got.Slug != tc.want {
			t.Errorf("Find(%q) = %v, want %s", tc.words, got, tc.want)
		}
	}
	// One title word alone names nothing.
	for _, words := range []string{"build me a lifegain deck", "power", "of", "Rohan", "storm"} {
		if p, ok := s.Find(words); ok {
			t.Errorf("Find(%q) = %s, want no precon", words, p.Slug)
		}
	}
}

// TestSideboardIsNotPartOfTheHundred is D-247: a sideboard is not one of
// the hundred the share rule of D-218 measures.
func TestSideboardIsNotPartOfTheHundred(t *testing.T) {
	s := loadSet(t)
	for _, tc := range []struct {
		slug string
		side int
	}{
		{"from-cute-to-brute", 5},
		{"tricky-terrain-collectors-edition", 1},
		{"goblin-storm", 0},
	} {
		p, ok := s.Get(tc.slug)
		if !ok {
			t.Fatalf("no precon %q", tc.slug)
		}
		if p.Cards != 100 {
			t.Errorf("%s: deck holds %d cards, want 100", tc.slug, p.Cards)
		}
		if p.Sideboard != tc.side {
			t.Errorf("%s: sideboard holds %d cards, want %d", tc.slug, p.Sideboard, tc.side)
		}
	}
}

// TestSplitSideboardReadsTheHeader covers the split without the snapshot.
func TestSplitSideboardReadsTheHeader(t *testing.T) {
	const text = "// COMMANDER\n1 Zada, Hedron Grinder (SLD) 2406 *F*\n\n1 Sol Ring (SLD) 2417\n\n// SIDEBOARD\n2 Omo, Queen of Vesuva (M3C) 149\n"
	deck, side := splitSideboard(text)
	if strings.Contains(deck, "Omo") {
		t.Error("the sideboard reached the deck")
	}
	if !strings.Contains(deck, "Sol Ring") {
		t.Error("the deck lost a card")
	}
	if side != 2 {
		t.Errorf("sideboard = %d, want 2", side)
	}
	// A file with no sideboard keeps every line.
	whole, none := splitSideboard("1 Sol Ring (SLD) 2417\n")
	if none != 0 || !strings.Contains(whole, "Sol Ring") {
		t.Errorf("a file with no sideboard split wrong: %q %d", whole, none)
	}
	// The bare header the shared parser reads is a header here too.
	bare, n := splitSideboard("1 Sol Ring (SLD) 2417\n\nSideboard\n3 Omo, Queen of Vesuva (M3C) 149\n")
	if n != 3 || strings.Contains(bare, "Omo") {
		t.Errorf("the bare Sideboard header was not read: %q %d", bare, n)
	}
}

// TestSplitSideboardStopsAtTheNextSection: a Commander section after the
// sideboard is not sideboard.
func TestSplitSideboardStopsAtTheNextSection(t *testing.T) {
	const text = "1 Sol Ring (SLD) 2417\n\n// SIDEBOARD\n2 Omo, Queen of Vesuva (M3C) 149\n\n// COMMANDER\n1 Zada, Hedron Grinder (SLD) 2406\n"
	deck, side := splitSideboard(text)
	if side != 2 {
		t.Errorf("sideboard = %d, want 2: the later Commander section is not sideboard", side)
	}
	if strings.Contains(deck, "Zada") || strings.Contains(deck, "Omo") {
		t.Error("a later section reached the deck")
	}
}

// TestTitlesAreProductNames: a slug is a file name, and the user reads
// the product name.
func TestTitlesAreProductNames(t *testing.T) {
	for _, tc := range []struct{ slug, want string }{
		{"lotr-riders-of-rohan", "Riders of Rohan"},
		{"tricky-terrain-collectors-edition", "Tricky Terrain"},
		{"from-cute-to-brute", "From Cute to Brute"},
		// Unverified product names keep the slug in Title Case.
		{"ff-cloud", "Ff Cloud"},
		{"lorwyn-blight-curse", "Lorwyn Blight Curse"},
	} {
		if got := titleOf(tc.slug); got != tc.want {
			t.Errorf("titleOf(%q) = %q, want %q", tc.slug, got, tc.want)
		}
	}
}

// TestWordsOfMatchesAShortForm covers the phrase rule without the
// snapshot. One title word alone names nothing.
func TestWordsOfMatchesAShortForm(t *testing.T) {
	title := strings.Fields("riders of rohan")
	for _, tc := range []struct {
		phrase string
		want   bool
	}{
		{"riders of rohan", true},
		{"rohan", false},
		{"of", false},
		{"rohan rohan", false},
		{"rohan, riders", true},
		{"upgrade my riders of rohan precon", false},
		{"", false},
	} {
		if got := wordsOf(strings.Fields(tc.phrase), title); got != tc.want {
			t.Errorf("wordsOf(%q) = %v, want %v", tc.phrase, got, tc.want)
		}
	}
}

// TestUnresolvedListsTheLostLists: a precon the index can not answer
// must not feed the share rule, and the loader says which ones those
// are. A one-card index answers no precon, so every one is listed.
func TestUnresolvedListsTheLostLists(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "o-solring", Name: "Sol Ring", TypeLine: "Artifact"},
	}, nil, nil, time.Unix(0, 0).UTC())
	s, err := Load(idx)
	if err != nil {
		t.Fatalf("precons: %v", err)
	}
	if got, all := len(s.Unresolved()), len(s.All()); got != all || all == 0 {
		t.Errorf("unresolved = %d of %d, want every one", got, all)
	}
	for _, p := range s.Unresolved() {
		if p.Unresolved == 0 {
			t.Errorf("%s is listed with no unresolved row", p.Slug)
		}
	}
}

// TestNothingIsUnresolvedAgainstTheSnapshot: every embedded list resolves.
func TestNothingIsUnresolvedAgainstTheSnapshot(t *testing.T) {
	s := loadSet(t)
	if got := s.Unresolved(); len(got) != 0 {
		t.Errorf("unresolved precons = %v, want none", got)
	}
}
