package quality

import (
	"math"
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/profile"
)

// landModel is a Commander model of the land features alone, with the
// spread of the stored model of the raw land count.
func landModel() *FormatModel {
	return &FormatModel{
		Keys:    []string{KeyLand, KeyLandNeed},
		Means:   []float64{30.62, -3},
		Stds:    []float64{5.52, 3},
		Weights: []float64{-0.301, -0.3},
	}
}

func bracketPower(b int32) *mtgv1.PowerLevel {
	return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: b}}
}

// landDeck is a bracket 3 Commander deck of the given lands, filled to
// 99 with three-drops and one cantrip.
func (w *commanderWorld) landDeck(lands int32) *mtgv1.Deck {
	spells := append(append([]*mtgv1.Card{}, w.threes[:98-lands]...), w.cantrip)
	d := w.deck(mtgv1.FormatId_FORMAT_ID_COMMANDER, w.commander, spells,
		map[*mtgv1.Card]int32{w.plains: lands / 2, w.island: lands - lands/2})
	d.Power = bracketPower(3)
	return d
}

func landReasons(fm *FormatModel, in Input) []string {
	var out []string
	for _, r := range reasons(fm, fm.vector(Features(in, fm)), ruleChecks(in), in.Profile) {
		if strings.Contains(r, "land") {
			out = append(out, r)
		}
	}
	return out
}

// TestReasonsSkipACountInsideItsBand reads a bracket 3 deck of 38 lands,
// inside the band of 34 to 38: the model names no land reason. A deck of
// 40 lands sits off the band, and the model names it (D-1123).
func TestReasonsSkipACountInsideItsBand(t *testing.T) {
	w := newCommanderWorld(t)
	fm := landModel()

	in := w.input(w.landDeck(38))
	var row *mtgv1.ProfileFeature
	for _, f := range in.Profile.GetFeatures() {
		if f.GetKey() == profile.KeyLand {
			row = f
		}
	}
	if row == nil || row.GetValue() != 38 || row.GetOffBand() || !row.GetHasHigh() {
		t.Fatalf("the land row = %v, want 38 inside its band", row)
	}
	if got := landReasons(fm, in); len(got) != 0 {
		t.Errorf("a deck inside its band names %q", got)
	}

	off := w.input(w.landDeck(40))
	got := landReasons(fm, off)
	if len(got) == 0 || !strings.Contains(strings.Join(got, "|"), "the land count sits above the norm of the format") {
		t.Errorf("a deck off its band names %q", got)
	}
}

// TestLandNeedReasonReadsItsOwnSign reads a deck under its land need
// whose value sits above the mean of the format: the words say fewer
// lands than the curve needs, not more (D-1123).
func TestLandNeedReasonReadsItsOwnSign(t *testing.T) {
	fm := &FormatModel{Keys: []string{KeyLandNeed}, Means: []float64{-3}, Stds: []float64{3}, Weights: []float64{-0.3}}
	z := []float64{(-1 - -3) / 3.0}
	got := reasons(fm, z, RuleRead{}, nil)
	if len(got) != 1 || !strings.HasPrefix(got[0], "the deck holds fewer lands than its curve needs") {
		t.Errorf("reasons = %q", got)
	}
	z = []float64{(2 - -3) / 3.0}
	got = reasons(fm, z, RuleRead{}, nil)
	if len(got) != 1 || !strings.HasPrefix(got[0], "the deck holds more lands than its curve needs") {
		t.Errorf("reasons = %q", got)
	}
}

// TestLandNeedFeature reads the lands against Karsten's need of each
// deck size, and keeps the raw count for a stored model (D-1123).
func TestLandNeedFeature(t *testing.T) {
	if !slices.Contains(Keys, KeyLandNeed) || slices.Contains(Keys, KeyLand) {
		t.Errorf("the fit keys must hold the land need and not the raw count")
	}
	w := newCommanderWorld(t)
	in := w.input(w.landDeck(38))
	f := Features(in, nil)
	rr := ruleChecks(in)
	want := 38 - (karstenLandBase + karstenLandPerMV*rr.AvgManaValue - karstenLandPerCheap*float64(rr.Cheap))
	if f[KeyLand] != 38 || math.Abs(f[KeyLandNeed]-want) > 1e-9 {
		t.Errorf("land %v, land need %v, want 38 and %v", f[KeyLand], f[KeyLandNeed], want)
	}

	sixty := w.deck(mtgv1.FormatId_FORMAT_ID_MODERN, nil, append(append([]*mtgv1.Card{}, w.threes[:35]...), w.cantrip),
		map[*mtgv1.Card]int32{w.plains: 12, w.island: 12})
	sin := w.input(sixty)
	var mv float64
	for _, r := range sin.Profile.GetFeatures() {
		if r.GetKey() == profile.KeyAvgManaValue {
			mv = r.GetValue()
		}
	}
	sf := Features(sin, nil)
	want = 24 - (karstenSixtyLandBase + karstenSixtyLandPerMV*mv - karstenLandPerCheap*1)
	if math.Abs(sf[KeyLandNeed]-want) > 1e-9 {
		t.Errorf("60-card land need %v, want %v", sf[KeyLandNeed], want)
	}
}

// TestShapeLinesFollowTheBracket reads the shape line of each bracket:
// brackets 1 to 3 and no bracket read the typical lists, and brackets 4
// and 5 the top lists with their cards. A model with no typical shape
// writes nothing for bracket 3 (D-1123).
func TestShapeLinesFollowTheBracket(t *testing.T) {
	fm := &FormatModel{
		Shape:        Shape{Lists: 3176, Lands: 27.39, AvgManaValue: 2.18},
		TypicalShape: Shape{Lists: 865, Lands: 36.6, AvgManaValue: 3.1},
		TopCards:     []string{"Sol Ring"},
	}
	s := NewScorer(&Model{Version: "v", Formats: map[string]*FormatModel{meta.FormatCommander: fm}})
	cmd := mtgv1.FormatId_FORMAT_ID_COMMANDER
	for _, b := range []int32{0, 1, 3} {
		power := bracketPower(b)
		if b == 0 {
			power = nil
		}
		got := strings.Join(s.ShapeLines(cmd, power), "|")
		if !strings.Contains(got, "The average Commander decks hold 37 lands at an average mana value of 3.10, over 865 lists.") ||
			strings.Contains(got, "Sol Ring") || strings.Contains(got, "top") {
			t.Errorf("bracket %d: %q", b, got)
		}
	}
	for _, b := range []int32{4, 5} {
		got := s.ShapeLines(cmd, bracketPower(b))
		if len(got) != 2 || !strings.Contains(got[0], "top commander lists hold 27 lands") || !strings.Contains(got[1], "Sol Ring") {
			t.Errorf("bracket %d: %q", b, got)
		}
	}
	fm.TypicalShape = Shape{}
	if got := s.ShapeLines(cmd, bracketPower(3)); got != nil {
		t.Errorf("a model with no typical shape writes %q for bracket 3", got)
	}
}
