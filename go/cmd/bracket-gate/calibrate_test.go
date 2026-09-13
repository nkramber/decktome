package main

import (
	"bytes"
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/collections"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/precons"
)

// TestCalibrationReadsBack is M-15: the calibration document reads back
// through readDecks with the bracket, the commander, and the cards of each
// deck. The cEDH pick keeps the newest whole top-finish list of one
// commander, from a field of 32 or more.
func TestCalibrationReadsBack(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "o-cap", Name: "Captain America, Team Leader"},
		{OracleId: "o-kinnan", Name: "Kinnan, Bonder Prodigy"},
		{OracleId: "o-tymna", Name: "Tymna the Weaver"},
		{OracleId: "o-kraum", Name: "Kraum, Ludevic's Opus"},
		{OracleId: "o-signet", Name: "Arcane Signet"},
		{OracleId: "o-island", Name: "Island"},
	}, nil, nil, time.Date(2026, 9, 4, 0, 0, 0, 0, time.UTC))
	decklists := []precons.Decklist{{
		Slug: "test", Name: "Test Precon", Commander: "Captain America, Team Leader",
		Rows: []collections.Row{{Name: "Arcane Signet", Quantity: 1}, {Name: "Island", Quantity: 98}},
	}}
	whole := []meta.Card{{Name: "Arcane Signet", Count: 1}, {Name: "Island", Count: 98}}
	top := func(id, date string, players int, commanders ...string) meta.List {
		return meta.List{Source: meta.SourceTopdeck, ID: id, Format: "commander", Date: date, Tier: "great",
			Placement: 1, Players: players, Commanders: commanders, Cards: whole}
	}
	lists := []meta.List{
		top("older", "2026-09-01", 64, "Kinnan, Bonder Prodigy"),
		top("newest", "2026-09-06", 64, "Kinnan, Bonder Prodigy"),
		top("partners", "2026-09-06", 64, "Tymna the Weaver", "Kraum, Ludevic's Opus"),
		top("small field", "2026-09-06", 16, "Tymna the Weaver"),
		{Source: meta.SourceTopdeck, ID: "unknown card", Format: "commander", Date: "2026-09-06", Tier: "great",
			Placement: 1, Players: 64, Commanders: []string{"Tymna the Weaver"},
			Cards: []meta.Card{{Name: "No Such Card", Count: 1}, {Name: "Island", Count: 98}}},
	}
	decks, err := calibrationDecks(idx, decklists, lists, 1)
	if err != nil {
		t.Fatal(err)
	}
	var b bytes.Buffer
	writeCalibration(&b, idx.AsOf, decks)
	back, err := readDecks(bytes.NewReader(b.Bytes()), idx)
	if err != nil {
		t.Fatalf("read back: %v\n%s", err, b.String())
	}
	if len(back) != 2 {
		t.Fatalf("decks = %d, want the precon and one cEDH list:\n%s", len(back), b.String())
	}
	for i, want := range []struct {
		bracket   int32
		commander string
		theme     string
	}{
		{2, "Captain America, Team Leader", "precon Test Precon"},
		{5, "Kinnan, Bonder Prodigy", "cedh 2026-09-06, place 1 of 64"},
	} {
		r := back[i]
		if r.prompt.Bracket != want.bracket || r.prompt.Commander != want.commander || r.prompt.Theme != want.theme {
			t.Errorf("deck %d = bracket %d, %q, %q, want %+v", i+1, r.prompt.Bracket, r.prompt.Commander, r.prompt.Theme, want)
		}
		n := int32(0)
		for _, c := range r.deck.GetCards() {
			n += c.GetCount()
			if c.GetName() == want.commander {
				t.Errorf("deck %d lists its commander in the 99", i+1)
			}
		}
		if n != 99 {
			t.Errorf("deck %d holds %d cards, want 99", i+1, n)
		}
	}
	if strings.Contains(b.String(), "older") || strings.Contains(b.String(), "unknown card") {
		t.Errorf("the document names a list id:\n%s", b.String())
	}

	// A thin set is an error: fewer qualifying cEDH lists than the lane
	// wants, or no precon at all.
	if _, err := calibrationDecks(idx, decklists, lists, 2); err == nil || !strings.Contains(err.Error(), "1 cEDH lists qualify, and the lane wants 2") {
		t.Errorf("a short cEDH set: err = %v, want an error", err)
	}
	if _, err := calibrationDecks(idx, nil, lists, 1); err == nil {
		t.Error("an empty precon set: no error")
	}
}
