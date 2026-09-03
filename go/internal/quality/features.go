package quality

import (
	"math"
	"slices"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/profile"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// Input is one deck to measure: the deck, its profile, and the card
// source that resolves its ids.
type Input struct {
	Deck    *mtgv1.Deck
	Profile *mtgv1.DeckProfile
	Cards   rules.CardSource
}

// Features measures every feature of a deck against a format model.
// The corpus features read the model's rates and pairs, and the shape
// features read the profile rows.
func Features(in Input, fm *FormatModel) map[string]float64 {
	out := map[string]float64{}
	rows := map[string]float64{}
	for _, r := range in.Profile.GetFeatures() {
		rows[r.GetKey()] = r.GetValue()
	}
	commander := in.Deck.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_COMMANDER

	type nonland struct {
		id    string
		count float64
		mv    float64
	}
	var spells []nonland
	var copies, playsets, singletons, low, high, lands float64
	sources := map[mtgv1.Color]float64{}
	pips := map[mtgv1.Color]bool{}
	for _, dc := range in.Deck.GetCards() {
		c, ok := in.Cards.ByOracleID(dc.GetOracleId())
		if !ok || dc.GetCount() <= 0 {
			continue
		}
		n := float64(dc.GetCount())
		if slices.Contains(c.GetCardTypes(), "Land") {
			lands += n
			for _, col := range c.GetProducedMana() {
				if col != mtgv1.Color_COLOR_C {
					sources[col] += n
				}
			}
			continue
		}
		for _, col := range c.GetColors() {
			pips[col] = true
		}
		spells = append(spells, nonland{id: c.GetOracleId(), count: n, mv: c.GetManaValue()})
		copies += n
		if dc.GetCount() >= 4 {
			playsets += n
		}
		if dc.GetCount() == 1 {
			singletons += n
		}
		if c.GetManaValue() <= 2 {
			low += n
		}
		if c.GetManaValue() >= 5 {
			high += n
		}
	}

	// The corpus features. A card the top lists never held reads the
	// prior, and a pair they never held together lifts nothing.
	var rateSum, unseen, pairSum float64
	var pairs int
	if fm != nil {
		for _, s := range spells {
			r := fm.Rate(s.id)
			rateSum += r * s.count
			if _, ok := fm.CardRates[s.id]; !ok {
				unseen += s.count
			}
		}
		for i := 0; i < len(spells); i++ {
			for j := i + 1; j < len(spells); j++ {
				pairs++
				pairSum += fm.Pairs[PairKey(spells[i].id, spells[j].id)]
			}
		}
	}
	if copies > 0 {
		out[KeyCardRate] = rateSum / copies
		out[KeyUnseenShare] = unseen / copies
		out[KeyCurveLow] = low / copies
		out[KeyCurveHigh] = high / copies
		if !commander {
			out[KeyPlaysetShare] = playsets / copies
			out[KeySingletonShare] = singletons / copies
		}
	}
	if pairs > 0 {
		out[KeySynergy] = pairSum / float64(pairs)
	}
	if commander {
		out[KeyPlaysetShare] = 0
		out[KeySingletonShare] = 0
	}
	// The source spread: the gap between the best and the worst served
	// color of the deck's spells, as a share of the lands. A mana base
	// of one color's basics under two-color spells reads high (D-485).
	if lands > 0 && len(pips) > 1 {
		best, worst := 0.0, math.Inf(1)
		for col := range pips {
			best = math.Max(best, sources[col])
			worst = math.Min(worst, sources[col])
		}
		out[KeySourceSpread] = (best - worst) / lands
	}

	// The shape and the jobs, from the profile.
	out[KeyLand] = rows[profile.KeyLand]
	out[KeyAvgManaValue] = rows[profile.KeyAvgManaValue]
	out[KeyColorSources] = rows[profile.KeyColorSources]
	if lands := rows[profile.KeyLand]; lands > 0 {
		out[KeyTappedShare] = rows[profile.KeyTappedLand] / lands
	}
	out[KeyManaTurnFour] = rows[profile.KeyManaTurnFour]
	out[KeyHandsTwoToFour] = rows[profile.KeyHandsTwoToFourLands]
	roleKeys := [][2]string{
		{KeyRamp, profile.KeyRamp}, {KeyDraw, profile.KeyDraw}, {KeyRemoval, profile.KeyRemoval},
		{KeyWipe, profile.KeyWipe}, {KeyInteraction, profile.KeyInteraction},
	}
	empty := 0.0
	for _, rk := range roleKeys {
		v := rows[rk[1]]
		out[rk[0]] = v
		if v == 0 && rk[0] != KeyWipe {
			empty++
		}
	}
	out[KeyEmptyRoles] = empty
	out[KeyFastMana] = rows[profile.KeyFastMana]
	out[KeyGameChanger] = rows[profile.KeyGameChanger]

	// The commander's own signal.
	if commander {
		var sig CommanderSignal
		if fm != nil {
			sig, _ = fm.Commander(in.Deck.GetCommanderOracleIds()...)
		}
		out[KeyCEDHSignal] = sig.CEDH
		out[KeyHighBracket] = sig.HighBracket
		out[KeyCommanderDecks] = math.Log1p(float64(sig.Decks))
	}
	return out
}

// vector reads the model's keys out of a feature map, standardized.
func (fm *FormatModel) vector(features map[string]float64) []float64 {
	z := make([]float64, len(fm.Keys))
	for i, k := range fm.Keys {
		if fm.Stds[i] == 0 {
			continue
		}
		z[i] = (features[k] - fm.Means[i]) / fm.Stds[i]
	}
	return z
}
