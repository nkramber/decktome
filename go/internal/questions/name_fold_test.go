package questions

import (
	"slices"
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// The D-716 tests of the commander row. The reader wrote "Grima as
// commander", and the row offered no card, because both cards that
// carry the name spell it "Gríma".

func grimaIndex() *cards.Index {
	legal := map[string]mtgv1.LegalityStatus{"commander": mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL}
	card := func(name, oracle string, rank int32) *mtgv1.Card {
		return &mtgv1.Card{
			Name: name, OracleId: oracle, EdhrecRank: rank,
			TypeLine: "Legendary Creature — Human Advisor", CanBeCommander: true, Legalities: legal,
		}
	}
	return cards.NewIndex([]*mtgv1.Card{
		card("Gríma Wormtongue", "o-wormtongue", 1),
		card("Gríma, Saruman's Footman", "o-footman", 2),
	}, nil, nil, time.Time{})
}

// TestResolveCommanderReadsANameWithoutItsAccent: "Grima" offers both
// cards that carry the name.
func TestResolveCommanderReadsANameWithoutItsAccent(t *testing.T) {
	h := &CandidateHints{Index: grimaIndex()}
	got := h.ResolveCommander("Grima")
	want := []string{"Gríma Wormtongue", "Gríma, Saruman's Footman"}
	if !slices.Equal(got, want) {
		t.Errorf("ResolveCommander(Grima) = %v, want %v", got, want)
	}
}

// TestCanLeadReadsANameWithoutItsAccent: the full name the reader typed
// again is a commander the index knows, so the row closes.
func TestCanLeadReadsANameWithoutItsAccent(t *testing.T) {
	h := &CandidateHints{Index: grimaIndex()}
	if canLead, known := h.CanLead("Grima, Saruman's Footman"); !known || !canLead {
		t.Errorf("CanLead = %v, %v, want a known commander", canLead, known)
	}
}

// TestSameCardReadsANameWithoutItsAccent: the short name and the full
// name merge, whatever the accent.
func TestSameCardReadsANameWithoutItsAccent(t *testing.T) {
	if !sameCard("Grima", "Gríma, Saruman's Footman") {
		t.Error("the short name without its accent reads as another card")
	}
}
