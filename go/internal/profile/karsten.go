package profile

import (
	"fmt"
	"sort"
	"strings"
)

// The Karsten source tables. Frank Karsten, "How Many Sources Do You
// Need to Consistently Cast Your Spells? A 2022 Update", TCGplayer,
// 2022-08-02, read 2026-09-02. A row is a cost shape: the generic part
// and the colored pips of one color, so "1CC" is a three-mana spell
// with two pips of the color. The number is how many sources of that
// color the deck needs to cast the spell on curve with 89 percent plus
// the mana value in consistency. The 60-card column assumes 25 lands,
// and the 99-card column 41 lands, the free mulligan, and the draw on
// turn one.

var sources60 = map[string]int{
	"C": 14, "1C": 13, "2C": 12, "3C": 10, "4C": 9, "5C": 9,
	"CC": 21, "1CC": 18, "2CC": 16, "3CC": 15, "4CC": 13, "5CC": 12,
	"CCC": 23, "1CCC": 21, "2CCC": 19, "3CCC": 17, "4CCC": 16,
	"CCCC": 24, "1CCCC": 22,
}

var sources99 = map[string]int{
	"C": 19, "1C": 19, "2C": 18, "3C": 16, "4C": 15, "5C": 14,
	"CC": 30, "1CC": 28, "2CC": 26, "3CC": 23, "4CC": 22, "5CC": 20,
	"CCC": 36, "1CCC": 33, "2CCC": 30, "3CCC": 28, "4CCC": 26,
	"CCCC": 39, "1CCCC": 36,
}

// maxGeneric is the largest generic part the table holds per pip count.
var maxGeneric = map[int]int{1: 5, 2: 5, 3: 4, 4: 1}

// sourcesNeeded is the table value for a spell with generic mana and
// pips of one color. A shape past the table reads the nearest row: more
// pips than four read as four, and more generic mana than the row holds
// reads as the last row, which asks the least.
func sourcesNeeded(deckSize, generic, pips int) int {
	if pips <= 0 {
		return 0
	}
	if pips > 4 {
		pips = 4
	}
	if generic < 0 {
		generic = 0
	}
	if m := maxGeneric[pips]; generic > m {
		generic = m
	}
	key := strings.Repeat("C", pips)
	if generic > 0 {
		key = fmt.Sprintf("%d%s", generic, key)
	}
	if deckSize >= 99 {
		return sources99[key]
	}
	return sources60[key]
}

// RockSource is what one mana rock or dork counts as a source of a
// color it makes. Karsten counts a rock as three-fourths of a source.
const RockSource = 0.75

// SourcePercentile is the share of a color's spells the requirement
// covers. The requirement of one color is the source count that casts
// this share of the deck's spells of that color on curve, so one greedy
// card does not set the bar for the whole deck.
const SourcePercentile = 0.8

// requirement is the source count that covers SourcePercentile of the
// needs, each need repeated by its count. Zero needs give zero.
func requirement(needs []int) int {
	if len(needs) == 0 {
		return 0
	}
	sorted := append([]int(nil), needs...)
	sort.Ints(sorted)
	i := int(float64(len(sorted)-1) * SourcePercentile)
	return sorted[i]
}
