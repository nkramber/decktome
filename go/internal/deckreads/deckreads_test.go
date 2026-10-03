package deckreads

import (
	"testing"
	"time"
)

var t0 = time.Date(2026, 10, 3, 12, 0, 0, 0, time.UTC)

// TestMatchesNeedsTheDeckAndTheText is D-1107: a link stays only after a
// read of the same deck and the same text, less than Life ago.
func TestMatchesNeedsTheDeckAndTheText(t *testing.T) {
	reads := Keep(nil, &Read{DeckID: 42, TextHash: Hash("Deck\n4 Lightning Bolt\n"), ReadAt: t0}, t0)
	for _, tc := range []struct {
		name string
		deck int64
		text string
		at   time.Time
		want bool
	}{
		{"the read", 42, "Deck\n4 Lightning Bolt\n", t0.Add(time.Minute), true},
		{"another deck", 43, "Deck\n4 Lightning Bolt\n", t0.Add(time.Minute), false},
		{"another text", 42, "Deck\n4 Shock\n", t0.Add(time.Minute), false},
		{"an old read", 42, "Deck\n4 Lightning Bolt\n", t0.Add(Life), false},
	} {
		if got := Matches(reads, tc.deck, tc.text, tc.at); got != tc.want {
			t.Errorf("%s: Matches = %v, want %v", tc.name, got, tc.want)
		}
	}
}

// TestKeepDropsOldReadsAndBoundsTheRest: an old read goes, and a document
// keeps the newest MaxReads.
func TestKeepDropsOldReadsAndBoundsTheRest(t *testing.T) {
	reads := []Read{{DeckID: 1, ReadAt: t0.Add(-Life)}}
	for i := range MaxReads + 5 {
		reads = Keep(reads, &Read{DeckID: int64(100 + i), ReadAt: t0}, t0)
	}
	if len(reads) != MaxReads {
		t.Fatalf("len = %d, want %d", len(reads), MaxReads)
	}
	if reads[0].DeckID != 105 || reads[len(reads)-1].DeckID != int64(100+MaxReads+4) {
		t.Errorf("kept %d to %d", reads[0].DeckID, reads[len(reads)-1].DeckID)
	}
}
