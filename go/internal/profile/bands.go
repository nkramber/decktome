package profile

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
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
	}
	if err := json.Unmarshal(bandsJSON, &f); err != nil {
		return nil, fmt.Errorf("bands.json: %w", err)
	}
	if f.VerifiedAt == "" {
		return nil, fmt.Errorf("bands.json: verified_at is missing")
	}
	out := &Bands{
		VerifiedAt: f.VerifiedAt,
		commander:  map[int32]map[string]Band{},
		sixty:      map[mtgv1.SixtyStep]map[string]Band{},
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

// DefaultBracket is the bracket a Commander deck with no bracket reads
// its bands at. D-93 names bracket 2 to 3 as the default, and 3 is the
// one whose bands hold the wider deck.
const DefaultBracket = 3

// Lines writes the bands a generator must build to, one line per
// feature the prompt can act on. The role counts go out as job targets
// already, so this names the rest: the curve, the mana base shape, and
// the tutor and fast-mana caps.
func (b *Bands) Lines(format mtgv1.FormatId, power *mtgv1.PowerLevel) []string {
	table, _ := b.For(format, power)
	var out []string
	for _, key := range promptKeys {
		band, ok := table[key]
		if !ok {
			continue
		}
		out = append(out, fmt.Sprintf("- %s: %s", promptWords[key], bandWords(key, band)))
	}
	return out
}

// promptKeys are the features the prompt names, in order.
var promptKeys = []string{KeyAvgManaValue, KeyTappedLand, KeyColorlessLand, KeyTutor, KeyFastMana}

var promptWords = map[string]string{
	KeyAvgManaValue:  "average mana value of the nonland cards",
	KeyTappedLand:    "lands that enter tapped",
	KeyColorlessLand: "nonbasic lands that make only colorless mana",
	KeyTutor:         "tutors, cards that search the library for a card",
	KeyFastMana:      "fast mana, nonland mana producers of mana value one or less",
}

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
