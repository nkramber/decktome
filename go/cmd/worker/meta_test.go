package main

import (
	"testing"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

func TestTopCommanders(t *testing.T) {
	idx := cards.NewIndex([]*mtgv1.Card{
		{OracleId: "a", Name: "A", CanBeCommander: true, EdhrecRank: 30},
		{OracleId: "b", Name: "B", CanBeCommander: true, EdhrecRank: 2},
		{OracleId: "c", Name: "C", CanBeCommander: false, EdhrecRank: 1},
		{OracleId: "d", Name: "D", CanBeCommander: true},
		{OracleId: "e", Name: "E", CanBeCommander: true, EdhrecRank: 10},
	}, nil, nil, time.Now())
	got := topCommanders(idx, 2)
	if len(got) != 2 || got[0] != "B" || got[1] != "E" {
		t.Fatalf("top = %v, want B then E: a non-commander and an unranked legend stay out", got)
	}
}
