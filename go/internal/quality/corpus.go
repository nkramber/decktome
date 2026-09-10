package quality

import (
	"math"
	"slices"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/profile"
)

// The corpus constants. RateAlpha is the smoothing weight of the prior,
// PairMinCards is the list count a card needs before its pairs count,
// PairMinTogether is the count a pair needs, PairMinLift is the lift a
// stored pair needs, and MaxPairs caps the stored pairs.
const (
	RateAlpha       = 5.0
	PairMinCards    = 5
	PairMinTogether = 3
	PairMinLift     = 1.5
	MaxPairs        = 200000
	TopCardCount    = 12
)

// The corpus reads two groups of lists, and each one answers a different
// question about a deck (PR-29, F-53).
//
// The top group is the great and the good lists. It answers "how close
// is this deck to a tournament list". It read alone until PR-29, and
// that is why a precon and a broken copy of it barely separated: both
// are far from a cEDH list, so the two strongest features moved little
// between them.
//
// The casual group is the typical and the baseline lists: the average
// decks of EDHREC, the MTGGoldfish lists, and the precons. It answers
// "how close is it to a deck people actually build". A precon's pairs
// sit in those lists and its broken copy's do not, so the two separate.

// tierWeight is the placement weight of a list in the top group: a great
// list counts twice a good one.
func tierWeight(tier string) float64 {
	switch tier {
	case meta.TierGreat:
		return 2
	case meta.TierGood:
		return 1
	default:
		return 0
	}
}

// casualTierWeight is the placement weight in the casual group. The two
// rungs count alike: neither a typical list nor a precon is the better
// deck, and they are two views of the same casual table.
func casualTierWeight(tier string) float64 {
	switch tier {
	case meta.TierTypical, meta.TierBaseline:
		return 1
	default:
		return 0
	}
}

// nonlandIDs lists the distinct nonland Oracle ids of a deck.
func nonlandIDs(deck *mtgv1.Deck, idx *cards.Index) []string {
	var out []string
	seen := map[string]bool{}
	for _, dc := range deck.GetCards() {
		c, ok := idx.ByOracleID(dc.GetOracleId())
		if !ok || seen[c.GetOracleId()] || slices.Contains(c.GetCardTypes(), "Land") {
			continue
		}
		seen[c.GetOracleId()] = true
		out = append(out, c.GetOracleId())
	}
	sort.Strings(out)
	return out
}

// group is the two tables one set of lists gives: the smoothed card
// rates and the pair lifts.
type group struct {
	rates map[string]float64
	prior float64
	pairs map[string]float64
	// weight is the placement weight per card, and names the card names.
	// The top cards of the prompt read both.
	weight map[string]float64
	names  map[string]string
}

// buildGroup reads every training list the weight function keeps, and it
// answers the card rates and the pair lifts of that group.
//
// One reader serves the top group and the casual group (PR-29), so the
// two can not drift apart. Before PR-29 this was the body of
// buildCorpus, and it read the top group alone.
func buildGroup(train []*Resolved, idx *cards.Index, weightOf func(string) float64) group {
	g := group{rates: map[string]float64{}, pairs: map[string]float64{},
		weight: map[string]float64{}, names: map[string]string{}}
	var total float64
	present := map[string]int{}
	var lists [][]string
	var listWeights []float64
	for _, r := range train {
		w := weightOf(r.List.Tier)
		if w == 0 {
			continue
		}
		total += w
		ids := nonlandIDs(r.Deck, idx)
		for _, id := range ids {
			g.weight[id] += w
			present[id]++
			if c, ok := idx.ByOracleID(id); ok {
				g.names[id] = c.GetName()
			}
		}
		lists = append(lists, ids)
		listWeights = append(listWeights, w)
	}
	if len(g.weight) == 0 || total == 0 {
		return g
	}
	prior := 1.0 / float64(len(g.weight))
	g.prior = RateAlpha * prior / (total + RateAlpha)
	for id, w := range g.weight {
		g.rates[id] = (w + RateAlpha*prior) / (total + RateAlpha)
	}

	// The pairs: the lift of two cards over chance, among the cards
	// enough lists hold.
	type pairKey [2]string
	together := map[pairKey]float64{}
	for li, ids := range lists {
		w := listWeights[li]
		kept := ids[:0:0]
		for _, id := range ids {
			if present[id] >= PairMinCards {
				kept = append(kept, id)
			}
		}
		for i := 0; i < len(kept); i++ {
			for j := i + 1; j < len(kept); j++ {
				together[pairKey{kept[i], kept[j]}] += w
			}
		}
	}
	type lifted struct {
		key  string
		lift float64
		w    float64
	}
	var out []lifted
	for k, w := range together {
		if w < PairMinTogether {
			continue
		}
		lift := (w * total) / (g.weight[k[0]] * g.weight[k[1]])
		if lift < PairMinLift {
			continue
		}
		out = append(out, lifted{PairKey(k[0], k[1]), math.Log(lift), w})
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i].lift*out[i].w, out[j].lift*out[j].w
		if a != b {
			return a > b
		}
		return out[i].key < out[j].key
	})
	if len(out) > MaxPairs {
		out = out[:MaxPairs]
	}
	g.pairs = make(map[string]float64, len(out))
	for _, p := range out {
		g.pairs[p.key] = round4(p.lift)
	}
	return g
}

// buildCorpus fills the rates, the pairs, the top cards, and the shape
// of a format model from the training split. The profiles are the ones
// the fit measured.
//
// It reads two groups (PR-29). The top group gives CardRates and Pairs,
// which every caller of the model has read since PR-14B. The casual
// group gives CasualRates and CasualPairs, and nothing but the two
// casual features reads those.
//
// The top cards and the shape stay the great lists' own: they go in the
// build prompt as the cards and the shape to aim at, and a precon is not
// what a reader aims at (D-474).
func buildCorpus(fm *FormatModel, train []*Resolved, profiles map[string]*mtgv1.DeckProfile, idx *cards.Index) {
	top := buildGroup(train, idx, tierWeight)
	fm.CardRates, fm.RatePrior, fm.Pairs = top.rates, top.prior, top.pairs

	casual := buildGroup(train, idx, casualTierWeight)
	fm.CasualRates, fm.CasualRatePrior, fm.CasualPairs = casual.rates, casual.prior, casual.pairs

	// The shape of the great lists alone.
	var shapeLands, shapeMV float64
	shapeLists := 0
	for _, r := range train {
		if r.List.Tier != meta.TierGreat {
			continue
		}
		p := profiles[r.List.Key()]
		if p == nil {
			continue
		}
		shapeLists++
		for _, row := range p.GetFeatures() {
			switch row.GetKey() {
			case profile.KeyLand:
				shapeLands += row.GetValue()
			case profile.KeyAvgManaValue:
				shapeMV += row.GetValue()
			}
		}
	}
	if shapeLists > 0 {
		fm.Shape = Shape{Lists: shapeLists, Lands: round2(shapeLands / float64(shapeLists)), AvgManaValue: round2(shapeMV / float64(shapeLists))}
	}

	// The top cards of the top group, by weight.
	type ranked struct {
		id string
		w  float64
	}
	var ranks []ranked
	for id, w := range top.weight {
		ranks = append(ranks, ranked{id, w})
	}
	sort.Slice(ranks, func(i, j int) bool {
		if ranks[i].w != ranks[j].w {
			return ranks[i].w > ranks[j].w
		}
		return top.names[ranks[i].id] < top.names[ranks[j].id]
	})
	fm.TopCards = nil
	for i := 0; i < len(ranks) && i < TopCardCount; i++ {
		fm.TopCards = append(fm.TopCards, top.names[ranks[i].id])
	}
}

// commanderSignals fills the commander table from the day's reads,
// resolved by name.
func commanderSignals(fm *FormatModel, reads []meta.Commander, idx *cards.Index) {
	fm.Commanders = map[string]CommanderSignal{}
	for _, c := range reads {
		sig := CommanderSignal{CEDH: round4(c.CEDHSignal()), HighBracket: round4(c.HighBracketShare()), Decks: c.NumDecks}
		if sig.CEDH == 0 && sig.HighBracket == 0 && sig.Decks == 0 {
			continue
		}
		var ids []string
		for _, name := range splitPair(c.Name) {
			card, ok := idx.ByName(name)
			if !ok {
				ids = nil
				break
			}
			ids = append(ids, card.GetOracleId())
		}
		if len(ids) == 0 {
			continue
		}
		fm.Commanders[CommanderKey(ids...)] = sig
	}
}

// splitPair reads "A + B" as two names.
func splitPair(name string) []string {
	var out []string
	for _, part := range strings.Split(name, " + ") {
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func round2(v float64) float64 { return math.Round(v*100) / 100 }
func round4(v float64) float64 { return math.Round(v*10000) / 10000 }
