package gatekit

import (
	"fmt"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/profile"
)

// RateShift is what the top-list rate of the app moves in one shortlist
// (F-129, D-706): the cards the rate brings in, and the power cards of
// the list without the rate and with it.
type RateShift struct {
	// NoRate reports a gate with no stored model, so no rate moves the
	// list.
	NoRate  bool
	Moved   int
	Cards   int
	Without Power
	With    Power
}

// Power counts the cards a high bracket reads (D-704).
type Power struct{ Tutors, FastMana, GameChangers int }

// ShiftOfRate builds the shortlist of req again with no top-list rate, and
// compares it with list, the shortlist req built. A dry run prints it, so
// a reader sees what the rate moves before any paid run.
func ShiftOfRate(cb *candidates.Builder, idx *cards.Index, req candidates.Request, list *candidates.List) (RateShift, error) {
	if req.MetaBoost == nil {
		return RateShift{NoRate: true, Cards: len(listCards(list))}, nil
	}
	req.MetaBoost = nil
	without, err := cb.Build(idx, req)
	if err != nil {
		return RateShift{}, fmt.Errorf("the shortlist with no rate: %w", err)
	}
	tags := idx.Tags()
	return RateShift{
		Moved:   movedCards(list, without),
		Cards:   len(listCards(list)),
		Without: powerOf(without, tags),
		With:    powerOf(list, tags),
	}, nil
}

// String writes the shift for the progress line of a dry run.
func (s RateShift) String() string {
	if s.NoRate {
		return "no stored model, so no top-list rate"
	}
	return fmt.Sprintf("the rate brought in %d of %d cards: tutors %d to %d, fast mana %d to %d, Game Changers %d to %d",
		s.Moved, s.Cards, s.Without.Tutors, s.With.Tutors, s.Without.FastMana, s.With.FastMana,
		s.Without.GameChangers, s.With.GameChangers)
}

// listCards is every card of a shortlist, the upgrades included.
func listCards(l *candidates.List) []*mtgv1.Card {
	if l == nil {
		return nil
	}
	var out []*mtgv1.Card
	for _, cs := range [][]candidates.Candidate{l.Candidates, l.Upgrades} {
		for _, c := range cs {
			if c.Card != nil {
				out = append(out, c.Card)
			}
		}
	}
	return out
}

// movedCards counts the cards that with holds and without does not.
func movedCards(with, without *candidates.List) int {
	held := map[string]bool{}
	for _, c := range listCards(without) {
		held[c.GetOracleId()] = true
	}
	n := 0
	for _, c := range listCards(with) {
		if !held[c.GetOracleId()] {
			n++
		}
	}
	return n
}

func powerOf(l *candidates.List, tags *cards.TagIndex) Power {
	tutors, fast, changers := profile.PowerCards(listCards(l), tags)
	return Power{Tutors: tutors, FastMana: fast, GameChangers: changers}
}
