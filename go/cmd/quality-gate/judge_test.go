package main

import (
	"strings"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
)

const judgeDoc = `# PR-8 deck gate

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

**Summary:** Karlov leads.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade.

- 10 Plains | land | Basic white sources.
- 1 Sol Ring | ramp | Mana.

### 2. a 60-card deck

Format: Standard. Theme: none. Pool: any_card. Shortlist: 90 names.

Commander: Karlov of the Ghost Council.

The quality model grades this deck good: the cards pair well.

- 4 Plains | land | Basics.

### 3. no grade

Format: Modern. Theme: none. Pool: any_card. Shortlist: 90 names.

- 4 Plains | land | Basics.
`

func TestReadDeckGate(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "plains", Name: "Plains"},
		{OracleId: "sol", Name: "Sol Ring"},
		{OracleId: "karlov", Name: "Karlov of the Ghost Council", CanBeCommander: true},
	}, nil, nil, time.Now())
	decks, err := readDeckGate(strings.NewReader(judgeDoc), idx, map[int][]string{1: {"Karlov of the Ghost Council"}})
	if err != nil {
		t.Fatal(err)
	}
	if len(decks) != 2 {
		t.Fatalf("decks = %d, want 2 (the deck with no grade is out)", len(decks))
	}
	first := decks[0]
	if first.grade != meta.TierBad || first.format != mtgv1.FormatId_FORMAT_ID_COMMANDER || len(first.deck.GetCards()) != 2 || first.deck.GetCards()[0].GetCount() != 10 {
		t.Errorf("first = %+v", first)
	}
	if len(first.deck.GetCommanderOracleIds()) != 1 || first.deck.GetCommanderOracleIds()[0] != "karlov" {
		t.Errorf("first commander from the prompts = %v", first.deck.GetCommanderOracleIds())
	}
	second := decks[1]
	if second.grade != meta.TierGood || len(second.deck.GetCommanderOracleIds()) != 1 {
		t.Errorf("second = %+v", second)
	}
	if _, err := readDeckGate(strings.NewReader("### 1. x\n\nThe quality model grades this deck good: x.\n\n- 1 Nope | land | x.\n"), idx, nil); err == nil {
		t.Errorf("an unknown card must fail the read")
	}
}

func TestTierOfWords(t *testing.T) {
	tests := map[string]string{
		"great": meta.TierGreat, "at the precon baseline": meta.TierBaseline, "below the precon baseline": meta.TierBad, "typical": meta.TierTypical, "nonsense": "",
	}
	for in, want := range tests {
		if got := tierOfWords(in); got != want {
			t.Errorf("%q: %q, want %q", in, got, want)
		}
	}
}
