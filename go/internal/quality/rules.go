package quality

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/profile"
	"github.com/nkramber/decktome/go/internal/rules"
)

// The land need of a 99-card deck: 31.42 + 3.13 × the average mana value
// of the spells − 0.28 × the cheap card draw or mana ramp spells (Frank
// Karsten, "How Many Lands Do You Need in Your Deck? An Updated
// Analysis", TCGplayer, 2022-07-29).
const (
	karstenLandBase     = 31.42
	karstenLandPerMV    = 3.13
	karstenLandPerCheap = 0.28
)

// The cuts of the rules checks (D-677, D-678). A Commander deck is
// flagged when its lands sit RuleLandShortfall or more under Karsten's
// need, when its spells average over RuleCurveCeiling, or when its worst
// color holds under RuleColorFloor of the sources its pips need.
const (
	RuleLandShortfall = 10.0
	RuleCurveCeiling  = 4.4
	RuleColorFloor    = 0.75
)

// RuleRead is the rules read of one deck. A check says flagged or not
// flagged and carries no probability, so it tips no deck to bad (F-115).
type RuleRead struct {
	// Checked is true for a Commander deck with a profile. The checks
	// read Commander alone (D-667).
	Checked bool
	// Lands, Curve, and Colors say which checks flag the deck.
	Lands, Curve, Colors bool
	LandCount            float64
	// Need is Karsten's land need of the deck, and Cheap the cheap draw
	// and ramp spells it counts.
	Need         float64
	Cheap        int
	AvgManaValue float64
	// ColorSources is the worst color's share of the sources it needs,
	// and HasColors says the profile measured one. A deck with no color
	// row reads no colors check.
	ColorSources float64
	HasColors    bool
	// WorstColor, WorstHave, and WorstNeed name the worst color of a deck
	// the colors check flags.
	WorstColor mtgv1.Color
	WorstHave  float64
	WorstNeed  int
}

// Flagged says whether a check flags the deck.
func (r RuleRead) Flagged() bool { return r.Lands || r.Curve || r.Colors }

// Shortfall is the lands the deck holds under Karsten's need.
func (r RuleRead) Shortfall() float64 { return r.Need - r.LandCount }

// ruleChecks reads a deck against the rules checks. A deck of another
// format, or one with no profile, reads no check.
func ruleChecks(in Input) RuleRead {
	if in.Deck.GetFormat().GetId() != mtgv1.FormatId_FORMAT_ID_COMMANDER || in.Profile == nil {
		return RuleRead{}
	}
	r := RuleRead{Checked: true}
	for _, f := range in.Profile.GetFeatures() {
		switch f.GetKey() {
		case profile.KeyLand:
			r.LandCount = f.GetValue()
		case profile.KeyAvgManaValue:
			r.AvgManaValue = f.GetValue()
		case profile.KeyColorSources:
			r.ColorSources, r.HasColors = f.GetValue(), true
		}
	}
	r.Cheap = cheapCount(in.Deck, in.Cards)
	r.Need = karstenLandBase + karstenLandPerMV*r.AvgManaValue - karstenLandPerCheap*float64(r.Cheap)
	r.Lands = r.Shortfall() >= RuleLandShortfall
	r.Curve = r.AvgManaValue > RuleCurveCeiling
	r.Colors = r.HasColors && r.ColorSources < RuleColorFloor
	if r.Colors && in.Cards != nil {
		worst := math.Inf(1)
		for _, s := range profile.ColorSources(in.Deck, in.Cards) {
			if ratio := s.Ratio(); ratio < worst {
				worst = ratio
				r.WorstColor, r.WorstHave, r.WorstNeed = s.Color, s.Have, s.Need
			}
		}
	}
	return r
}

// reasons names the checks that flag the deck, in the words the summary
// carries.
func (r RuleRead) reasons() []string {
	var out []string
	if r.Lands {
		out = append(out, fmt.Sprintf("the deck holds %.0f lands, and its curve needs about %.0f", r.LandCount, r.Need))
	}
	if r.Curve {
		out = append(out, fmt.Sprintf("the spells average %.2f mana, over the ceiling of %.1f", r.AvgManaValue, RuleCurveCeiling))
	}
	if r.Colors {
		word := colorWords[r.WorstColor]
		if word == "" {
			out = append(out, fmt.Sprintf("the worst color holds %.0f%% of the sources its spells need", r.ColorSources*100))
		} else {
			out = append(out, fmt.Sprintf("the %s sources cover %s of the %d its spells need", word, trimZero(r.WorstHave), r.WorstNeed))
		}
	}
	return out
}

var colorWords = map[mtgv1.Color]string{
	mtgv1.Color_COLOR_W: "white", mtgv1.Color_COLOR_U: "blue", mtgv1.Color_COLOR_B: "black",
	mtgv1.Color_COLOR_R: "red", mtgv1.Color_COLOR_G: "green",
}

// trimZero writes a source count with no trailing zeros: 9, 8.75.
func trimZero(v float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", v), "0"), ".")
}

// tierProbabilities answers the probabilities the tier reads. The rules
// set the tier of a Commander deck they check: a flagged deck grades
// bad, and every other deck takes the tier of the ladder alone. The
// detector keeps the score and sets no tier there (D-678). A deck of
// another format reads the probabilities of the grade.
func (fm *FormatModel) tierProbabilities(z, graded []float64, r RuleRead) []float64 {
	if !r.Checked {
		return graded
	}
	if r.Flagged() {
		out := make([]float64, len(graded))
		bad := slices.Index(fm.Tiers, meta.TierBad)
		if bad < 0 {
			bad = 0
		}
		if len(out) > 0 {
			out[bad] = 1
		}
		return out
	}
	return levelProbabilities(fm.Weights, fm.Thresholds, z)
}

// argmax answers the index of the largest value, the first on a tie.
func argmax(p []float64) int {
	best := 0
	for k, v := range p {
		if v > p[best] {
			best = k
		}
	}
	return best
}

// cyclesForOne matches a cycling cost of one mana.
var cyclesForOne = regexp.MustCompile(`cycling \{[1wubrgc]\}`)

// isCheap is Karsten's cheap card draw or cheap mana ramp spell, by the
// text rules of his article of 2022-07-29. A card with a land face
// counts as a land, as his formula counts it.
func isCheap(c *mtgv1.Card) bool {
	if c.GetManaValue() > 2 {
		return false
	}
	text := c.GetOracleText()
	typeLine := c.GetTypeLine()
	if faces := c.GetFaces(); len(faces) > 1 {
		typeLine = faces[0].GetTypeLine()
		if text == "" {
			parts := make([]string, 0, len(faces))
			for _, f := range faces {
				parts = append(parts, f.GetOracleText())
			}
			text = strings.Join(parts, "\n")
		}
	}
	if strings.Contains(typeLine, "Land") {
		return false
	}
	for _, f := range c.GetFaces() {
		if strings.Contains(f.GetTypeLine(), "Land") {
			return false
		}
	}
	t := strings.ToLower(text)
	all := func(words ...string) bool {
		for _, w := range words {
			if !strings.Contains(t, w) {
				return false
			}
		}
		return true
	}
	some := func(words ...string) bool {
		for _, w := range words {
			if strings.Contains(t, w) {
				return true
			}
		}
		return false
	}
	creature := strings.Contains(typeLine, "Creature")
	// Cheap card draw.
	if some("draw a card", "draw two cards", "draw three cards", "draws cards", "draws two cards", "draws three cards") &&
		!some("{4}", "blood token", "investigate") && (!creature || all("when", "enters")) {
		return true
	}
	if !creature && all("look", "library", "put", "your hand") && !some("pay ", "pays") {
		return true
	}
	if cyclesForOne.MatchString(t) {
		return true
	}
	// Cheap mana ramp.
	if strings.Contains(t, "add ") && !some("add its ability", "add a lore counter") && (!creature || !strings.Contains(t, "dies")) {
		return true
	}
	if all("search", "your library") && some("land", "basic") && !strings.Contains(t, "sacrifice") {
		return true
	}
	if all("enchanted land is tapped", "adds an additional") {
		return true
	}
	return all("put a creature card with", "from your hand onto the battlefield")
}

// cheapCount counts the cheap draw and ramp spells of a deck, each copy
// once.
func cheapCount(deck *mtgv1.Deck, src rules.CardSource) int {
	if src == nil {
		return 0
	}
	n := 0
	for _, dc := range deck.GetCards() {
		c, ok := src.ByOracleID(dc.GetOracleId())
		if !ok || slices.Contains(c.GetCardTypes(), "Land") {
			continue
		}
		if isCheap(c) {
			n += int(dc.GetCount())
		}
	}
	return n
}
