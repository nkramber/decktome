package generate

import (
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/profile"
)

// CodeBracketCut reports the cards the build cut after the content check,
// because the bracket forbids them (PR-45a, D-702). It is an INFO: the
// deck the reader gets no longer holds them.
const CodeBracketCut = "bracket_cut"

// bracketCut is one card the bracket cut, and the content it held.
type bracketCut struct {
	card *mtgv1.DeckCard
	why  string
}

// cutForbidden removes the cards a deck's bracket forbids, and returns
// each one with the content it held, in cut order (PR-45a, D-702). A commander, a locked card, and a
// card a revision keeps never leave. One cut breaks each combo: the card
// in the most forbidden combos first, then a card outside the precon, and
// then the card with the lowest shortlist score. A combo whose every card
// stays keeps its warning.
func cutForbidden(deck *mtgv1.Deck, req Request, f profile.Forbidden) []bracketCut {
	if f.Empty() {
		return nil
	}
	byName := make(map[string]*mtgv1.DeckCard, len(deck.GetCards()))
	for _, dc := range deck.GetCards() {
		byName[cutKey(dc.GetName())] = dc
	}
	kept := map[string]bool{}
	for _, id := range deck.GetCommanderOracleIds() {
		kept[id] = true
	}
	for _, id := range req.Locked {
		kept[id] = true
	}
	keptNames := map[string]bool{}
	if req.Revision != nil {
		for _, n := range req.Revision.Keep {
			keptNames[cutKey(n)] = true
		}
	}
	precon := make(map[string]bool, len(req.PreconOracleIDs))
	for _, id := range req.PreconOracleIDs {
		precon[id] = true
	}
	cuttable := func(name string) (*mtgv1.DeckCard, bool) {
		dc, ok := byName[cutKey(name)]
		if !ok || kept[dc.GetOracleId()] || keptNames[cutKey(dc.GetName())] {
			return nil, false
		}
		return dc, true
	}
	var out []bracketCut
	gone := map[string]bool{}
	remove := func(dc *mtgv1.DeckCard, why string) {
		if !gone[dc.GetOracleId()] {
			gone[dc.GetOracleId()] = true
			out = append(out, bracketCut{card: dc, why: why})
		}
	}

	combos := append([][]string(nil), f.Combos...)
	sort.SliceStable(combos, func(i, j int) bool {
		return strings.Join(combos[i], " + ") < strings.Join(combos[j], " + ")
	})
	inCombos := map[string]int{}
	for _, c := range combos {
		for _, n := range c {
			inCombos[cutKey(n)]++
		}
	}
	for _, c := range combos {
		var best *mtgv1.DeckCard
		broken := false
		for _, n := range c {
			if dc, ok := byName[cutKey(n)]; ok && gone[dc.GetOracleId()] {
				broken = true
				break
			}
			if dc, ok := cuttable(n); ok && (best == nil || cutsBefore(dc, best, inCombos, precon, req.Pool)) {
				best = dc
			}
		}
		if !broken && best != nil {
			remove(best, "the combo "+strings.Join(c, " + "))
		}
	}
	for _, n := range f.MassLandDenial {
		if dc, ok := cuttable(n); ok {
			remove(dc, "mass land denial")
		}
	}
	if len(f.ExtraTurns) > 0 {
		var turns []*mtgv1.DeckCard
		for _, n := range f.ExtraTurns {
			if dc, ok := byName[cutKey(n)]; ok && !gone[dc.GetOracleId()] {
				turns = append(turns, dc)
			}
		}
		// The cards that must stay fill the cap first, then the cards with
		// the highest shortlist score.
		sort.SliceStable(turns, func(i, j int) bool {
			_, ci := cuttable(turns[i].GetName())
			_, cj := cuttable(turns[j].GetName())
			if ci != cj {
				return !ci
			}
			return scoresAbove(req.Pool, turns[i], turns[j])
		})
		why := "an extra-turn card"
		if f.ExtraTurnCap > 0 {
			why = fmt.Sprintf("an extra-turn card past the limit of %d", f.ExtraTurnCap)
		}
		for i, dc := range turns {
			if i < f.ExtraTurnCap {
				continue
			}
			if _, ok := cuttable(dc.GetName()); ok {
				remove(dc, why)
			}
		}
	}
	if len(out) > 0 {
		cards := make([]*mtgv1.DeckCard, 0, len(deck.GetCards()))
		for _, dc := range deck.GetCards() {
			if !gone[dc.GetOracleId()] {
				cards = append(cards, dc)
			}
		}
		deck.Cards = cards
	}
	return out
}

// cutsBefore orders two cards of one combo for the cut: the card in more
// forbidden combos, then a card outside the precon, then the card with the
// lower shortlist score. The name breaks a tie.
func cutsBefore(a, b *mtgv1.DeckCard, inCombos map[string]int, precon map[string]bool, pool *Pool) bool {
	ka, kb := cutKey(a.GetName()), cutKey(b.GetName())
	if inCombos[ka] != inCombos[kb] {
		return inCombos[ka] > inCombos[kb]
	}
	if pa, pb := precon[a.GetOracleId()], precon[b.GetOracleId()]; pa != pb {
		return !pa
	}
	if above, below := scoresAbove(pool, b, a), scoresAbove(pool, a, b); above != below {
		return above
	}
	return ka > kb
}

// scoresAbove reports whether card a ranks above card b on the shortlist
// score. The shortlist groups its cards by role and the pool sorts its
// names by the alphabet, so the score is the only rank of a card. A card
// with no score, one the pool took as given, ranks above every scored
// card.
func scoresAbove(pool *Pool, a, b *mtgv1.DeckCard) bool {
	sa, oka := shortlistScore(pool, a)
	sb, okb := shortlistScore(pool, b)
	if oka != okb {
		return !oka
	}
	return oka && sa > sb
}

// shortlistScore is the shortlist score of a deck card, by its full name
// or by its front face.
func shortlistScore(pool *Pool, dc *mtgv1.DeckCard) (float64, bool) {
	if pool == nil {
		return 0, false
	}
	if s, ok := pool.Score(dc.GetName()); ok {
		return s, true
	}
	front, _, _ := strings.Cut(dc.GetName(), " // ")
	return pool.Score(front)
}

// cutKey matches a card name from the endpoint to a deck card name: the
// folded front face, because a double-faced card can come back under
// either name.
func cutKey(name string) string {
	front, _, _ := strings.Cut(name, " // ")
	return candidates.FoldName(front)
}
