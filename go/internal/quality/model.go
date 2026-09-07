// Package quality is the deck quality model of PR-14B (D-413 to
// D-417). It reads a deck and answers a tier of the ladder of D-414
// and the three strongest reasons in words.
//
// The model is one scorer per format, fitted in Go over the published
// lists the meta package stores, by ordinal regression over a feature
// vector. The features are what a person reads a deck by: card quality
// as the inclusion rate of a card in the top lists of its format,
// synergy as the lift of a card pair over chance, the shape the bracket
// profile of PR-14A measures, and the role counts. A bad list is
// synthetic: the engine breaks one axis of a real list, so each defect
// carries its own label.
//
// The scorer runs on the deterministic side, free at run time, and no
// prompt holds a weight. The code keeps the final say (guardrail 3).
package quality

import (
	"encoding/json"
	"fmt"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/meta"
)

// The feature keys, in the order a reader scans a deck: the cards, the
// pairs, the mana base, the curve, the jobs, the power signals, and the
// commander.
const (
	KeyCardRate       = "card_rate"
	KeyUnseenShare    = "unseen_share"
	KeySynergy        = "synergy"
	KeyLand           = "land"
	KeyAvgManaValue   = "avg_mana_value"
	KeyColorSources   = "color_sources"
	KeyTappedShare    = "tapped_share"
	KeyManaTurnFour   = "mana_turn_four"
	KeyHandsTwoToFour = "hands_two_to_four_lands"
	KeyCurveLow       = "curve_low"
	KeyCurveHigh      = "curve_high"
	KeyRamp           = "ramp"
	KeyDraw           = "draw"
	KeyRemoval        = "removal"
	KeyWipe           = "wipe"
	KeyInteraction    = "interaction"
	KeyEmptyRoles     = "empty_roles"
	KeyFastMana       = "fast_mana"
	KeyGameChanger    = "game_changer"
	KeyPlaysetShare   = "playset_share"
	KeySingletonShare = "singleton_share"
	KeySourceSpread   = "source_spread"
	KeyCEDHSignal     = "cedh_signal"
	KeyHighBracket    = "high_bracket_share"
	KeyCommanderDecks = "commander_decks"
)

// Keys lists every feature the fit measures. A feature with no spread
// in a format's lists leaves the format's model.
var Keys = []string{
	KeyCardRate, KeyUnseenShare, KeySynergy,
	KeyLand, KeyAvgManaValue, KeyColorSources, KeyTappedShare, KeyManaTurnFour, KeyHandsTwoToFour,
	KeyCurveLow, KeyCurveHigh,
	KeyRamp, KeyDraw, KeyRemoval, KeyWipe, KeyInteraction, KeyEmptyRoles,
	KeyFastMana, KeyGameChanger, KeyPlaysetShare, KeySingletonShare, KeySourceSpread,
	KeyCEDHSignal, KeyHighBracket, KeyCommanderDecks,
}

// Model is the fitted scorer of every format, stored beside the card
// snapshot as meta/model/<version>/quality.json.gz.
type Model struct {
	// Version is the fit time, for example 20260902T150000Z.
	Version string `json:"version"`
	// SnapshotAsOf is the card snapshot the fit resolved names against.
	SnapshotAsOf string `json:"snapshot_as_of"`
	// Formats holds one scorer per format word (meta.FormatCommander,
	// meta.FormatStandard, meta.FormatModern).
	Formats map[string]*FormatModel `json:"formats"`
}

// FormatModel is the scorer of one format.
type FormatModel struct {
	Format string `json:"format"`
	// Tiers is the ladder of the format, worst first, the rungs the
	// lists held. A 60-card format has no typical rung (D-414).
	Tiers []string `json:"tiers"`
	// TrainCounts and HoldoutCounts count the lists per tier.
	TrainCounts   map[string]int `json:"train_counts"`
	HoldoutCounts map[string]int `json:"holdout_counts"`
	// Keys, Means, Stds, and Weights run in one order. A feature is
	// standardized before its weight applies.
	Keys    []string  `json:"keys"`
	Means   []float64 `json:"means"`
	Stds    []float64 `json:"stds"`
	Weights []float64 `json:"weights"`
	// Thresholds are the cut points between the rungs, ascending, one
	// fewer than the tiers.
	Thresholds []float64 `json:"thresholds"`
	// DefectWeights and DefectBias are the defect detector: a logistic
	// model over the same standardized features, fitted on every real
	// list against the broken copies of the precons (D-485). One linear
	// ladder can not rank a top list over a precon and a precon over its
	// broken copy at once, so the detector reads the second question.
	DefectWeights []float64 `json:"defect_weights,omitempty"`
	DefectBias    float64   `json:"defect_bias,omitempty"`
	// DefectThreshold is the probability at which the detector flags a
	// deck, the cut that best separates the precons from their copies
	// on the training rows. Zero reads as one half.
	DefectThreshold float64 `json:"defect_threshold,omitempty"`
	// Holdout is the separation measured on the held-out lists.
	Holdout Holdout `json:"holdout"`
	// CardRates is the smoothed inclusion rate per Oracle id in the
	// great and good lists of the training split, for every card one
	// of them held. RatePrior is the rate of a card none held.
	CardRates map[string]float64 `json:"card_rates"`
	RatePrior float64            `json:"rate_prior"`
	// Pairs holds the log lift of a card pair over chance, keyed by the
	// two Oracle ids in order with a bar between, for the pairs that
	// lift. A pair not here lifts nothing.
	Pairs map[string]float64 `json:"pairs"`
	// TopCards names the nonland cards the great lists hold most, for
	// the prompt.
	TopCards []string `json:"top_cards"`
	// Shape is the mean shape of the great lists, for the prompt.
	Shape Shape `json:"shape"`
	// Commanders holds the commander signals by Oracle id, Commander
	// alone. A pair is keyed by both ids in order with a plus between.
	Commanders map[string]CommanderSignal `json:"commanders,omitempty"`
}

// Holdout is the gate measure of a format (roadmap PR-14B): a great
// list scores above a precon in 90 percent of the pairs, and a precon
// above a synthetic bad deck in 95 percent. DefectAccuracy is the
// detector's own accuracy over the held-out precons and their copies.
type Holdout struct {
	Lists             int       `json:"lists"`
	GreatOverBaseline PairShare `json:"great_over_baseline"`
	BaselineOverBad   PairShare `json:"baseline_over_bad"`
	Accuracy          float64   `json:"accuracy"`
	DefectAccuracy    float64   `json:"defect_accuracy"`
	Confusion         [][]int   `json:"confusion"`
	// BadByDefect splits BaselineOverBad by the broken axis.
	BadByDefect map[string]PairShare `json:"bad_by_defect,omitempty"`
}

// PairShare counts ordered pairs and the ones the scorer ordered right.
type PairShare struct {
	Pairs int `json:"pairs"`
	Wins  int `json:"wins"`
}

// Share is wins over pairs, or 0 with no pair.
func (p PairShare) Share() float64 {
	if p.Pairs == 0 {
		return 0
	}
	return float64(p.Wins) / float64(p.Pairs)
}

// Shape is the mean shape of the great lists of a format.
type Shape struct {
	Lists        int     `json:"lists"`
	Lands        float64 `json:"lands"`
	AvgManaValue float64 `json:"avg_mana_value"`
}

// CommanderSignal is the power read of one commander.
type CommanderSignal struct {
	// CEDH is the bracket 5 signal, in [0, 1] (meta.Commander.CEDHSignal).
	CEDH float64 `json:"cedh"`
	// HighBracket is the share of the commander's bracketed EDHREC decks
	// at bracket 4 or 5.
	HighBracket float64 `json:"high_bracket"`
	// Decks is the EDHREC deck count.
	Decks int `json:"decks"`
}

// Decode reads a stored model.
func Decode(data []byte) (*Model, error) {
	var m Model
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("quality: model: %w", err)
	}
	if m.Version == "" || len(m.Formats) == 0 {
		return nil, fmt.Errorf("quality: model carries no version or no format")
	}
	for word, fm := range m.Formats {
		if err := fm.check(); err != nil {
			return nil, fmt.Errorf("quality: model %s, %s: %w", m.Version, word, err)
		}
	}
	return &m, nil
}

// Encode writes a model.
func (m *Model) Encode() ([]byte, error) { return json.Marshal(m) }

// Format answers the scorer of a format id, or nil.
func (m *Model) Format(f mtgv1.FormatId) *FormatModel {
	if m == nil {
		return nil
	}
	return m.Formats[meta.FormatWord(f)]
}

func (fm *FormatModel) check() error {
	n := len(fm.Keys)
	if n == 0 || len(fm.Means) != n || len(fm.Stds) != n || len(fm.Weights) != n {
		return fmt.Errorf("the keys, means, stds, and weights differ in length")
	}
	if len(fm.DefectWeights) != 0 && len(fm.DefectWeights) != n {
		return fmt.Errorf("the defect weights differ in length from the keys")
	}
	if len(fm.Tiers) < 2 || len(fm.Thresholds) != len(fm.Tiers)-1 {
		return fmt.Errorf("%d tiers and %d thresholds", len(fm.Tiers), len(fm.Thresholds))
	}
	for i := 1; i < len(fm.Thresholds); i++ {
		if fm.Thresholds[i] < fm.Thresholds[i-1] {
			return fmt.Errorf("the thresholds are not ascending")
		}
	}
	return nil
}

// Rate answers the inclusion rate of a card, the prior for one no top
// list held.
func (fm *FormatModel) Rate(oracleID string) float64 {
	if r, ok := fm.CardRates[oracleID]; ok {
		return r
	}
	return fm.RatePrior
}

// PairKey is the key of two Oracle ids in Pairs.
func PairKey(a, b string) string {
	if b < a {
		a, b = b, a
	}
	return a + "|" + b
}

// CommanderKey is the key of a commander or a pair in Commanders.
func CommanderKey(ids ...string) string {
	if len(ids) == 2 && ids[1] < ids[0] {
		ids = []string{ids[1], ids[0]}
	}
	out := ""
	for i, id := range ids {
		if i > 0 {
			out += "+"
		}
		out += id
	}
	return out
}

// Commander answers the signal of a commander or a pair. A pair with
// no row of its own reads the stronger partner.
func (fm *FormatModel) Commander(ids ...string) (CommanderSignal, bool) {
	if fm == nil || len(ids) == 0 {
		return CommanderSignal{}, false
	}
	if s, ok := fm.Commanders[CommanderKey(ids...)]; ok {
		return s, true
	}
	var best CommanderSignal
	found := false
	for _, id := range ids {
		s, ok := fm.Commanders[id]
		if !ok {
			continue
		}
		found = true
		if s.CEDH > best.CEDH || (s.CEDH == best.CEDH && s.HighBracket > best.HighBracket) {
			best = s
		}
	}
	return best, found
}
