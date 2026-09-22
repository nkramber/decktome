package profile

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

//go:embed bands.json
var bandsJSON []byte

// Band is the range one feature must sit in. A nil High is no ceiling.
type Band struct {
	Low  float64  `json:"low"`
	High *float64 `json:"high"`
}

// Holds reports whether v sits in the band.
func (b Band) Holds(v float64) bool {
	if v < b.Low {
		return false
	}
	return b.High == nil || v <= *b.High
}

// Bands is the loaded band table: one map of feature to band per
// Commander bracket, and one per 60-card power step.
type Bands struct {
	VerifiedAt string
	commander  map[int32]map[string]Band
	sixty      map[mtgv1.SixtyStep]map[string]Band
	// The fixing floors are keyed by the count of deck colors, so they
	// sit outside the feature tables (F-33, D-799).
	fixCommander map[int32]map[int]float64
	fixSixty     map[mtgv1.SixtyStep]map[int]float64
}

var sixtyWords = map[string]mtgv1.SixtyStep{
	"casual":     mtgv1.SixtyStep_SIXTY_STEP_CASUAL,
	"fnm":        mtgv1.SixtyStep_SIXTY_STEP_FNM,
	"tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
}

// LoadBands parses the embedded band table. A bad file fails loudly, and
// a band with a ceiling under its floor is a bad file.
func LoadBands() (*Bands, error) {
	var f struct {
		VerifiedAt string                     `json:"verified_at"`
		Commander  map[string]map[string]Band `json:"commander"`
		Sixty      map[string]map[string]Band `json:"sixty"`
		Fixing     struct {
			Commander map[string]map[string]float64 `json:"commander"`
			Sixty     map[string]map[string]float64 `json:"sixty"`
		} `json:"fixing_land"`
	}
	if err := json.Unmarshal(bandsJSON, &f); err != nil {
		return nil, fmt.Errorf("bands.json: %w", err)
	}
	if f.VerifiedAt == "" {
		return nil, fmt.Errorf("bands.json: verified_at is missing")
	}
	out := &Bands{
		VerifiedAt:   f.VerifiedAt,
		commander:    map[int32]map[string]Band{},
		sixty:        map[mtgv1.SixtyStep]map[string]Band{},
		fixCommander: map[int32]map[int]float64{},
		fixSixty:     map[mtgv1.SixtyStep]map[int]float64{},
	}
	for k, table := range f.Commander {
		n, err := strconv.Atoi(k)
		if err != nil {
			return nil, fmt.Errorf("bands.json: bad bracket %q", k)
		}
		if err := checkTable(table); err != nil {
			return nil, fmt.Errorf("bands.json: bracket %s: %w", k, err)
		}
		out.commander[int32(n)] = table
	}
	for k, table := range f.Sixty {
		step, ok := sixtyWords[k]
		if !ok {
			return nil, fmt.Errorf("bands.json: bad power step %q", k)
		}
		if err := checkTable(table); err != nil {
			return nil, fmt.Errorf("bands.json: step %s: %w", k, err)
		}
		out.sixty[step] = table
	}
	for n := int32(1); n <= 5; n++ {
		if _, ok := out.commander[n]; !ok {
			return nil, fmt.Errorf("bands.json: bracket %d is missing", n)
		}
	}
	for k, row := range f.Fixing.Commander {
		n, err := strconv.Atoi(k)
		if _, ok := out.commander[int32(n)]; err != nil || !ok {
			return nil, fmt.Errorf("bands.json: fixing_land: bad bracket %q", k)
		}
		floors, err := fixingRow(row)
		if err != nil {
			return nil, fmt.Errorf("bands.json: fixing_land: bracket %s: %w", k, err)
		}
		out.fixCommander[int32(n)] = floors
	}
	for k, row := range f.Fixing.Sixty {
		step, ok := sixtyWords[k]
		if !ok {
			return nil, fmt.Errorf("bands.json: fixing_land: bad power step %q", k)
		}
		floors, err := fixingRow(row)
		if err != nil {
			return nil, fmt.Errorf("bands.json: fixing_land: step %s: %w", k, err)
		}
		out.fixSixty[step] = floors
	}
	return out, nil
}

// fixingRow parses the floors of one power, keyed by a color count of 2
// to 5. A deck of one color needs no fixing, so it holds no key.
func fixingRow(row map[string]float64) (map[int]float64, error) {
	out := map[int]float64{}
	for k, v := range row {
		n, err := strconv.Atoi(k)
		if err != nil || n < 2 || n > 5 {
			return nil, fmt.Errorf("bad color count %q", k)
		}
		if v < 0 {
			return nil, fmt.Errorf("color count %s: floor %g is under zero", k, v)
		}
		out[n] = v
	}
	return out, nil
}

func checkTable(table map[string]Band) error {
	for key, b := range table {
		if b.High != nil && *b.High < b.Low {
			return fmt.Errorf("%s: high %g is under low %g", key, *b.High, b.Low)
		}
	}
	return nil
}

// For returns the band table of a deck: the bracket's for Commander, and
// the power step's for a 60-card format. A Commander deck with no
// bracket reads bracket 3, the middle of the default of D-93. A 60-card
// deck with no step reads the casual step. The returned bracket is the
// one the table is for, 0 for a 60-card deck.
func (b *Bands) For(format mtgv1.FormatId, power *mtgv1.PowerLevel) (table map[string]Band, bracket int32) {
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		bracket = power.GetBracket()
		if _, ok := b.commander[bracket]; !ok {
			bracket = DefaultBracket
		}
		return b.commander[bracket], bracket
	}
	step := power.GetSixtyStep()
	if _, ok := b.sixty[step]; !ok {
		step = mtgv1.SixtyStep_SIXTY_STEP_CASUAL
	}
	return b.sixty[step], 0
}

// FixingFloor answers the fixing lands a deck of the given count of
// colors must hold, and 0 when its power holds no floor (F-33, D-799). A
// fixing land is a land of LandClassOf below LandOther. The power reads
// as For reads it.
func (b *Bands) FixingFloor(format mtgv1.FormatId, power *mtgv1.PowerLevel, colors int) float64 {
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		_, bracket := b.For(format, power)
		return b.fixCommander[bracket][colors]
	}
	step := power.GetSixtyStep()
	if _, ok := b.sixty[step]; !ok {
		step = mtgv1.SixtyStep_SIXTY_STEP_CASUAL
	}
	return b.fixSixty[step][colors]
}

// DefaultBracket is the bracket a Commander deck with no bracket reads
// its bands at. D-93 names bracket 2 to 3 as the default, and 3 is the
// one whose bands hold the wider deck.
const DefaultBracket = 3

// Lines writes the bands a generator must build to, one line per
// feature the prompt can act on. The role counts go out as job targets
// already, so this names the rest: the curve, the mana base shape, the
// tutor and fast-mana caps, and the three numbers the simulation reads.
//
// Every feature that can raise an off-band finding reaches the model,
// here or in the job targets (F-78). Before this the prompt named five
// of the fifteen features `bands.json` holds, and 12 of the 26 off-band
// findings of deck gate run 16 and bracket gate run 1 were on two
// features no prompt ever stated.
func (b *Bands) Lines(format mtgv1.FormatId, power *mtgv1.PowerLevel) []string {
	table, _ := b.For(format, power)
	var out []string
	for _, key := range promptKeys {
		band, ok := table[key]
		if !ok {
			continue
		}
		if line := derivedLine(key, band, format); line != "" {
			out = append(out, line)
			continue
		}
		line := fmt.Sprintf("- %s: %s", promptWords[key], bandWords(key, band))
		// A power floor names the mark the shortlist gives each card that
		// counts toward it, so the model counts what the check counts
		// (D-704).
		if mark := powerMarks[key]; mark != "" && band.Low > 0 {
			line += fmt.Sprintf(", and the shortlist marks each one %q", mark)
		}
		out = append(out, line)
	}
	if line := b.fixingLine(format, power); line != "" {
		out = append(out, line)
	}
	return out
}

// fixingLine names the fixing floor of each count of deck colors. A
// 60-card prompt reads it before the model picks the colors, so the line
// names every count (F-33, D-799). It is "" when the power holds no
// floor.
func (b *Bands) fixingLine(format mtgv1.FormatId, power *mtgv1.PowerLevel) string {
	var parts []string
	for n := 2; n <= 5; n++ {
		if v := b.FixingFloor(format, power, n); v > 0 {
			parts = append(parts, fmt.Sprintf("%s in a %s-color deck", num(v), colorCountWords[n]))
		}
	}
	if len(parts) == 0 {
		return ""
	}
	list := parts[0]
	if len(parts) > 1 {
		list = strings.Join(parts[:len(parts)-1], ", ") + ", and " + parts[len(parts)-1]
	}
	return "- fixing lands, lands that make two or more of the deck's colors: at least " + list + ". " +
		"A dual land, a land that makes any color, and a fetch land that finds two of the deck's colors each count as one. " +
		"A basic land counts as none."
}

var colorCountWords = map[int]string{2: "two", 3: "three", 4: "four", 5: "five"}

// Words writes one band as the phrase the job block reads, or "" when
// the format and the power carry no band for the key (F-78, Part 2).
func (b *Bands) Words(format mtgv1.FormatId, power *mtgv1.PowerLevel, key string) string {
	table, _ := b.For(format, power)
	band, ok := table[key]
	if !ok {
		return ""
	}
	return bandWords(key, band)
}

// PowerFloors answers each power feature whose band holds a floor above
// zero, with the floor: the tutors, the fast mana, and the Game Changers
// of brackets 4 and 5 (D-704). Every other deck reads none.
func (b *Bands) PowerFloors(format mtgv1.FormatId, power *mtgv1.PowerLevel) map[string]float64 {
	table, _ := b.For(format, power)
	out := map[string]float64{}
	for _, key := range PowerKeys {
		if band, ok := table[key]; ok && band.Low > 0 {
			out[key] = band.Low
		}
	}
	return out
}

// Mark is the word the shortlist writes beside a card that counts toward
// a power floor, and "" for every other feature (D-704).
func Mark(key string) string { return powerMarks[key] }

var powerMarks = map[string]string{
	KeyTutor:       "tutor",
	KeyFastMana:    "fast mana",
	KeyGameChanger: "Game Changer",
	KeyFinisher:    "finisher",
}

// promptKeys are the features the prompt names, in order.
var promptKeys = []string{
	KeyAvgManaValue, KeyTappedLand, KeyColorlessLand, KeyTutor, KeyFastMana, KeyGameChanger,
	KeyFinisher, KeyColorSources, KeyManaTurnFour, KeyHandsTwoToFourLands, KeyCommanderTurnOverMV,
}

var promptWords = map[string]string{
	KeyAvgManaValue:  "average mana value of the nonland cards",
	KeyTappedLand:    "lands that enter tapped",
	KeyColorlessLand: "nonbasic lands that make only colorless mana",
	KeyTutor:         "tutors, cards that search the library for a card",
	KeyFastMana:      "fast mana, nonland mana producers of mana value one or less",
	KeyGameChanger:   "Game Changers, cards on the official Game Changers list",
	KeyFinisher:      "finishers, cards that can win the game, and evasive creatures of power 5 or more",
}

// derivedLine writes the four features a count can not state. Each one
// is a measured number, so the line names the number and the cards that
// move it. It returns "" for a key that reads as a plain band.
//
// The source counts come from the Karsten tables of `karsten.go`. Frank
// Karsten, "How Many Sources Do You Need to Consistently Cast Your
// Spells? A 2022 Update", TCGplayer, 2022-08-02, read 2026-09-02.
func derivedLine(key string, band Band, format mtgv1.FormatId) string {
	switch key {
	case KeyColorSources:
		one, two := 13, 18
		if format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
			one, two = 19, 26
		}
		return fmt.Sprintf("- color sources: every color of the deck needs its own. A card with one colored pip "+
			"in its cost wants about %d sources of that color, and a card with two pips wants about %d. "+
			"A land that makes the color counts as one source, and a mana rock or a mana creature as three fourths.", one, two)
	case KeyManaTurnFour:
		return fmt.Sprintf("- mana on turn four: the deck must reach %s mana on turn four, as a mean over ten thousand "+
			"opening hands. More lands raise it, lands that enter untapped raise it, and ramp of mana value two or "+
			"less raises it.", num(band.Low))
	case KeyHandsTwoToFourLands:
		return fmt.Sprintf("- opening hands: at least %s of first seven-card hands must hold two to four lands. "+
			"The land count alone decides this, so keep the land count inside its band.", percentWords(band.Low))
	case KeyCommanderTurnOverMV:
		if band.High == nil {
			return ""
		}
		return fmt.Sprintf("- the commander comes down on time: the mean turn it is castable must be no more than "+
			"%s of a turn past its mana value. Ramp and lands that enter untapped decide this.", num(*band.High))
	}
	return ""
}

// percentWords writes a share as a percentage, so a prompt line reads
// "70 percent" and never "0.7".
func percentWords(v float64) string { return num(math.Round(v*100)) + " percent" }

// bandWords writes a band as a phrase: "2.5 to 3.5", "at most 9", or
// "4 or more".
func bandWords(key string, b Band) string {
	if b.High == nil {
		if b.Low == 0 {
			return "no limit"
		}
		return fmt.Sprintf("%s or more", num(b.Low))
	}
	if b.Low == 0 && !isRatio(key) {
		return fmt.Sprintf("at most %s", num(*b.High))
	}
	return fmt.Sprintf("%s to %s", num(b.Low), num(*b.High))
}

func isRatio(key string) bool { return key == KeyAvgManaValue || key == KeyColorSources }

// num writes a band number without a trailing zero.
func num(v float64) string {
	s := strconv.FormatFloat(v, 'f', 2, 64)
	s = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	if s == "" || s == "-" {
		return "0"
	}
	return s
}

// Midpoints returns the middle of each role band, rounded, for the
// job targets of the generator. Only the role keys are returned.
func (b *Bands) Midpoints(format mtgv1.FormatId, power *mtgv1.PowerLevel) map[string]int {
	table, _ := b.For(format, power)
	out := map[string]int{}
	for _, key := range roleKeys {
		band, ok := table[key]
		if !ok || band.High == nil {
			continue
		}
		out[key] = int((band.Low + *band.High) / 2)
	}
	return out
}

// roleKeys are the features that are job targets of the prompt.
var roleKeys = []string{KeyLand, KeyRamp, KeyDraw, KeyRemoval, KeyWipe, KeyInteraction}

// Keys lists every feature key of a table, sorted, for a report.
func (b *Bands) Keys(format mtgv1.FormatId, power *mtgv1.PowerLevel) []string {
	table, _ := b.For(format, power)
	keys := make([]string, 0, len(table))
	for k := range table {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
