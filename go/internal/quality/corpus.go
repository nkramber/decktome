package quality

import (
	"math"
	"slices"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
	"github.com/nkramber/mtg-deck-builder/go/internal/profile"
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

// tierWeight is the placement weight of a list in the rates and the
// pairs: a great list counts twice a good one.
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

// buildCorpus fills the rates, the pairs, the top cards, and the shape
// of a format model from the great and good lists of the training
// split. The profiles are the ones the fit measured.
func buildCorpus(fm *FormatModel, train []*Resolved, profiles map[string]*mtgv1.DeckProfile, idx *cards.Index) {
	var total float64
	weight := map[string]float64{}
	present := map[string]int{}
	names := map[string]string{}
	var lists [][]string
	var listWeights []float64
	var shapeLands, shapeMV float64
	shapeLists := 0
	for _, r := range train {
		w := tierWeight(r.List.Tier)
		if w == 0 {
			continue
		}
		total += w
		ids := nonlandIDs(r.Deck, idx)
		for _, id := range ids {
			weight[id] += w
			present[id]++
			if c, ok := idx.ByOracleID(id); ok {
				names[id] = c.GetName()
			}
		}
		lists = append(lists, ids)
		listWeights = append(listWeights, w)
		if r.List.Tier == meta.TierGreat {
			if p := profiles[r.List.Key()]; p != nil {
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
		}
	}
	fm.CardRates = map[string]float64{}
	if len(weight) == 0 || total == 0 {
		fm.RatePrior = 0
		fm.Pairs = map[string]float64{}
		return
	}
	fm.RatePrior = RateAlpha * (1.0 / float64(len(weight))) / (total + RateAlpha)
	prior := 1.0 / float64(len(weight))
	for id, w := range weight {
		fm.CardRates[id] = (w + RateAlpha*prior) / (total + RateAlpha)
	}
	if shapeLists > 0 {
		fm.Shape = Shape{Lists: shapeLists, Lands: round2(shapeLands / float64(shapeLists)), AvgManaValue: round2(shapeMV / float64(shapeLists))}
	}

	// The top cards of the great lists, by weight.
	type ranked struct {
		id string
		w  float64
	}
	var top []ranked
	for id, w := range weight {
		top = append(top, ranked{id, w})
	}
	sort.Slice(top, func(i, j int) bool {
		if top[i].w != top[j].w {
			return top[i].w > top[j].w
		}
		return names[top[i].id] < names[top[j].id]
	})
	fm.TopCards = nil
	for i := 0; i < len(top) && i < TopCardCount; i++ {
		fm.TopCards = append(fm.TopCards, names[top[i].id])
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
		lift := (w * total) / (weight[k[0]] * weight[k[1]])
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
	fm.Pairs = make(map[string]float64, len(out))
	for _, p := range out {
		fm.Pairs[p.key] = round4(p.lift)
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
