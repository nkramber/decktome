package quality

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync/atomic"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/meta"
)

// Scorer grades decks with the loaded model. The model swaps when the
// worker fits a new one, so the scorer holds it behind an atomic.
type Scorer struct {
	model atomic.Pointer[Model]
}

// NewScorer makes a scorer. A nil model grades nothing until Swap.
func NewScorer(m *Model) *Scorer {
	s := &Scorer{}
	if m != nil {
		s.model.Store(m)
	}
	return s
}

// Swap installs a model.
func (s *Scorer) Swap(m *Model) { s.model.Store(m) }

// Model answers the loaded model, or nil.
func (s *Scorer) Model() *Model {
	if s == nil {
		return nil
	}
	return s.model.Load()
}

// Version names the loaded model, or "".
func (s *Scorer) Version() string {
	if m := s.Model(); m != nil {
		return m.Version
	}
	return ""
}

// ReasonCount is how many reasons a grade names.
const ReasonCount = 3

// Score grades a deck. It answers nil when no model covers the format.
func (s *Scorer) Score(in Input) *mtgv1.DeckQuality {
	m := s.Model()
	fm := m.Format(in.Deck.GetFormat().GetId())
	if fm == nil {
		return nil
	}
	features := Features(in, fm)
	z := fm.vector(features)
	graded, score := fm.grade(z)
	rr := ruleChecks(in)
	p := fm.tierProbabilities(z, graded, rr)
	out := &mtgv1.DeckQuality{ModelVersion: m.Version, Score: round4(score)}
	for k, v := range p {
		out.Probabilities = append(out.Probabilities, &mtgv1.TierProbability{Tier: fm.Tiers[k], Probability: round4(v)})
	}
	out.Tier = fm.Tiers[argmax(p)]
	out.Reasons = reasons(fm, z, rr)
	return out
}

// CardRate answers the inclusion rate of a card in a format, 0 with no
// model.
func (s *Scorer) CardRate(f mtgv1.FormatId, oracleID string) float64 {
	fm := s.Model().Format(f)
	if fm == nil {
		return 0
	}
	return fm.Rate(oracleID)
}

// MetaBoost answers the shortlist signal of a format: the card's rate
// against the best rate of the format, in [0, 1]. Nil with no model,
// so the shortlist ranks as before (candidates.Request.MetaBoost).
func (s *Scorer) MetaBoost(f mtgv1.FormatId) func(oracleID string) float64 {
	fm := s.Model().Format(f)
	if fm == nil || len(fm.CardRates) == 0 {
		return nil
	}
	top := 0.0
	for _, r := range fm.CardRates {
		top = math.Max(top, r)
	}
	if top == 0 {
		return nil
	}
	return func(oracleID string) float64 { return fm.Rate(oracleID) / top }
}

// CommanderSignal answers the bracket 5 signal of a commander, for the
// pool that ranks a bracket 4 or 5 request (OQ-48). Nil with no model.
func (s *Scorer) CommanderSignal() func(ids ...string) float64 {
	fm := s.Model().Format(mtgv1.FormatId_FORMAT_ID_COMMANDER)
	if fm == nil || len(fm.Commanders) == 0 {
		return nil
	}
	return func(ids ...string) float64 {
		sig, _ := fm.Commander(ids...)
		return sig.CEDH
	}
}

// CardQualities answers the per-format rows of one card, for GetCards.
func (s *Scorer) CardQualities(oracleID string) []*mtgv1.CardQuality {
	m := s.Model()
	if m == nil {
		return nil
	}
	var out []*mtgv1.CardQuality
	for _, f := range []mtgv1.FormatId{mtgv1.FormatId_FORMAT_ID_COMMANDER, mtgv1.FormatId_FORMAT_ID_STANDARD, mtgv1.FormatId_FORMAT_ID_MODERN} {
		fm := m.Format(f)
		if fm == nil {
			continue
		}
		row := &mtgv1.CardQuality{Format: f, Inclusion: round4(fm.Rate(oracleID))}
		if sig, ok := fm.Commander(oracleID); ok {
			row.CedhSignal = sig.CEDH
		}
		out = append(out, row)
	}
	return out
}

// ShapeLines writes the format shape for the generate prompt: the
// mean land count and mana value of the great lists, and the cards
// they hold most (roadmap PR-14B). Nil with no model.
func (s *Scorer) ShapeLines(f mtgv1.FormatId) []string {
	fm := s.Model().Format(f)
	if fm == nil || fm.Shape.Lists == 0 {
		return nil
	}
	lines := []string{
		fmt.Sprintf("- The top %s lists hold %.0f lands at an average mana value of %.2f, over %d lists.",
			strings.ToLower(formatWord(f)), fm.Shape.Lands, fm.Shape.AvgManaValue, fm.Shape.Lists),
	}
	if len(fm.TopCards) > 0 {
		lines = append(lines, "- The cards they hold most: "+strings.Join(fm.TopCards, ", ")+".")
	}
	return lines
}

func formatWord(f mtgv1.FormatId) string {
	switch f {
	case mtgv1.FormatId_FORMAT_ID_COMMANDER:
		return "Commander"
	case mtgv1.FormatId_FORMAT_ID_STANDARD:
		return "Standard"
	case mtgv1.FormatId_FORMAT_ID_MODERN:
		return "Modern"
	default:
		return f.String()
	}
}

// phrase is the words of one feature: what a high value says, and what
// a low value says, of the deck against the norm of its format.
type phrase struct {
	high, low string
}

var phrases = map[string]phrase{
	KeyCardRate:       {"the cards are ones the top lists of the format play", "few of the cards are ones the top lists of the format play"},
	KeyUnseenShare:    {"many of the cards appear in no top list", "nearly every card appears in a top list"},
	KeySynergy:        {"the cards pair the way the top lists pair them", "the cards pair in ways the top lists do not"},
	KeyLand:           {"the land count sits above the norm of the format", "the land count sits below the norm of the format"},
	KeyAvgManaValue:   {"the curve sits high for the format", "the curve sits low for the format"},
	KeyColorSources:   {"the color sources cover the pips", "the color sources fall short of the pips"},
	KeyTappedShare:    {"many lands enter tapped", "few lands enter tapped"},
	KeyManaTurnFour:   {"the deck makes more mana on turn four than the norm", "the deck makes less mana on turn four than the norm"},
	KeyHandsTwoToFour: {"more opening hands hold two to four lands than the norm", "fewer opening hands hold two to four lands than the norm"},
	KeyCurveLow:       {"the deck holds many cheap spells", "the deck holds few cheap spells"},
	KeyCurveHigh:      {"the deck holds many five-drops and up", "the deck holds few five-drops and up"},
	KeyRamp:           {"the deck ramps more than the norm", "the deck ramps less than the norm"},
	KeyDraw:           {"the deck draws more than the norm", "the deck draws less than the norm"},
	KeyRemoval:        {"the deck holds more removal than the norm", "the deck holds less removal than the norm"},
	KeyWipe:           {"the deck holds more wipes than the norm", "the deck holds fewer wipes than the norm"},
	KeyInteraction:    {"the deck holds more interaction than the norm", "the deck holds less interaction than the norm"},
	KeyEmptyRoles:     {"a job the format fills is empty", "every job the format fills is filled"},
	KeyFastMana:       {"the deck holds fast mana", "the deck holds no fast mana"},
	KeyGameChanger:    {"the deck holds Game Changers", "the deck holds no Game Changer"},
	KeyPlaysetShare:   {"the deck runs its spells as playsets", "the deck runs few playsets"},
	KeySingletonShare: {"the deck runs many spells as one copy", "the deck runs few spells as one copy"},
	KeySourceSpread:   {"the mana base serves one color far better than another", "the mana base serves the colors evenly"},
	KeyCEDHSignal:     {"the commander places in cEDH events", "the commander does not place in cEDH events"},
	KeyHighBracket:    {"the commander leads high-bracket decks", "the commander leads low-bracket decks"},
	KeyCommanderDecks: {"the commander is a popular one", "few decks lead with the commander"},
}

// reasons names the ReasonCount strongest signals, the strongest first.
// The strength is the weight times the standardized value, and the
// words follow the value's side of the norm and say which way the
// signal moved the grade.
func reasons(fm *FormatModel, z []float64, r RuleRead) []string {
	// A deck the rules flag names the checks, and a deck the rules check
	// and pass names the ladder, because the ladder set its tier (D-678).
	if r.Flagged() {
		out := r.reasons()
		if len(out) > ReasonCount {
			out = out[:ReasonCount]
		}
		return out
	}
	type contribution struct {
		key string
		c   float64
		z   float64
	}
	var all []contribution
	// A deck of another format that the detector reads as broken names
	// the detector's signals, because they are what lowered the grade
	// (D-485). The detector's weight raises the defect, so its sign turns.
	weights := fm.Weights
	sign := 1.0
	if !r.Checked && fm.flagged(z) {
		weights = fm.DefectWeights
		sign = -1
	}
	for i, k := range fm.Keys {
		if _, ok := phrases[k]; !ok || z[i] == 0 {
			continue
		}
		all = append(all, contribution{k, sign * weights[i] * z[i], z[i]})
	}
	sort.SliceStable(all, func(i, j int) bool { return math.Abs(all[i].c) > math.Abs(all[j].c) })
	var out []string
	for _, c := range all {
		if len(out) >= ReasonCount {
			break
		}
		if math.Abs(c.c) < 1e-6 {
			continue
		}
		p := phrases[c.key]
		words := p.high
		if c.z < 0 {
			words = p.low
		}
		verb := "raises"
		if c.c < 0 {
			verb = "lowers"
		}
		out = append(out, fmt.Sprintf("%s, and that %s the grade", words, verb))
	}
	return out
}

// Summary writes the one sentence the deck summary and the revision
// note carry: the tier and the reasons.
func Summary(q *mtgv1.DeckQuality) string {
	if q == nil || q.GetTier() == "" {
		return ""
	}
	word := TierWord(q.GetTier())
	// The frame is named: one ladder per format, its top the tournament
	// lists, so a bracket 2 reader reads a power read (D-487).
	if len(q.GetReasons()) == 0 {
		return fmt.Sprintf("The quality model grades this deck %s against the top lists of the format.", word)
	}
	parts := make([]string, 0, len(q.GetReasons()))
	for i, r := range q.GetReasons() {
		if i > 0 && r != "" {
			r = strings.ToUpper(r[:1]) + r[1:]
		}
		parts = append(parts, r)
	}
	return fmt.Sprintf("The quality model grades this deck %s against the top lists of the format: %s.", word, strings.Join(parts, ". "))
}

// TierWord is the tier as a grade word.
func TierWord(tier string) string {
	switch tier {
	case meta.TierGreat:
		return "great"
	case meta.TierGood:
		return "good"
	case meta.TierTypical:
		return "typical"
	case meta.TierBaseline:
		return "at the precon baseline"
	case meta.TierBad:
		return "below the precon baseline"
	default:
		return tier
	}
}

// TierDrop writes the sentence of a revision that lowered the grade,
// and "" when the grade held, rose, or is missing on either side.
func TierDrop(before, after *mtgv1.DeckQuality) string {
	if before == nil || after == nil {
		return ""
	}
	b, a := meta.TierLevel(before.GetTier()), meta.TierLevel(after.GetTier())
	if b < 0 || a < 0 || a >= b {
		return ""
	}
	return fmt.Sprintf("This change lowers the deck's grade from %s to %s.", TierWord(before.GetTier()), TierWord(after.GetTier()))
}

// Contribution is one feature's part in a grade: the raw value, the
// standardized value, and the products with the ladder's weight and
// the detector's weight.
type Contribution struct {
	Key    string
	Value  float64
	Z      float64
	Ladder float64
	Defect float64
}

// Explanation is the read of one deck for a gate document: the grade,
// the detector's probability and cut, the rules read, and every
// contribution.
type Explanation struct {
	Tier          string
	Score         float64
	Defect        float64
	Threshold     float64
	Flagged       bool
	Rules         RuleRead
	Ladder        []float64
	Contributions []Contribution
}

// Explain grades a deck and answers how. Nil when no model covers the
// format.
func (s *Scorer) Explain(in Input) *Explanation {
	m := s.Model()
	fm := m.Format(in.Deck.GetFormat().GetId())
	if fm == nil {
		return nil
	}
	features := Features(in, fm)
	z := fm.vector(features)
	graded, score := fm.grade(z)
	rr := ruleChecks(in)
	p := fm.tierProbabilities(z, graded, rr)
	out := &Explanation{Tier: fm.Tiers[argmax(p)], Score: round4(score), Defect: round4(fm.defect(z)), Threshold: fm.DefectThreshold, Flagged: fm.flagged(z), Rules: rr, Ladder: p}
	for i, k := range fm.Keys {
		c := Contribution{Key: k, Value: round4(features[k]), Z: round4(z[i]), Ladder: round4(fm.Weights[i] * z[i])}
		if len(fm.DefectWeights) == len(z) {
			c.Defect = round4(fm.DefectWeights[i] * z[i])
		}
		out.Contributions = append(out.Contributions, c)
	}
	return out
}
