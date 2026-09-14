package gatekit

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
)

// TestMovedCardsCountsWhatTheRateBringsIn is the count of the gate dry run
// (D-706): a card the rate list holds and the list with no rate does not,
// the upgrades included.
func TestMovedCardsCountsWhatTheRateBringsIn(t *testing.T) {
	c := func(id string) candidates.Candidate { return candidates.Candidate{Card: &mtgv1.Card{OracleId: id}} }
	with := &candidates.List{Candidates: []candidates.Candidate{c("a"), c("b"), c("tutor")}, Upgrades: []candidates.Candidate{c("u")}}
	without := &candidates.List{Candidates: []candidates.Candidate{c("a"), c("b"), c("theme")}, Upgrades: []candidates.Candidate{c("u")}}
	if got := movedCards(with, without); got != 1 {
		t.Errorf("moved %d, want 1", got)
	}
	if got := len(listCards(with)); got != 4 {
		t.Errorf("cards %d, want 4", got)
	}
	if got := (RateShift{NoRate: true}).String(); got != "no stored model, so no top-list rate" {
		t.Errorf("no rate reads %q", got)
	}
}
