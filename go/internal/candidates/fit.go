package candidates

import (
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// FitFloor is the least theme score of a new card that fits a deck
// (D-1091). One subtype or one card type of the theme reaches it, and so
// does one tag. A keyword alone or a text needle alone does not, because
// hundreds of cards hold each one. A score is a share of scoreCap.
const FitFloor = weightSubtype / scoreCap

// deckThemeShare is the least share of the nonland cards of a deck that
// fits a theme, for DeckTheme to name it (D-1091). Below it the theme
// is noise of a few cards, and not the plan of the deck.
const deckThemeShare = 0.25

// ThemeScores returns the theme score of each card, in the order of cs,
// for a theme in the user's words (D-1091). It resolves the theme the way
// Build does, so a new card scores as it would score in a build.
func (b *Builder) ThemeScores(theme string, idx *cards.Index, cs []*mtgv1.Card) []float64 {
	m := b.themes.matchIn(theme, idx)
	out := make([]float64, len(cs))
	if m.Empty() {
		return out
	}
	for i, c := range cs {
		out[i], _ = m.score(c)
	}
	return out
}

// DeckTheme names the theme that the most cards of a deck fit, for a
// deck that has no theme in words (D-1091). It reads each row of the
// theme table, and each creature subtype of the deck. The row with the
// highest sum of scores wins, and the name comes first on a tie. It
// returns "" when no row fits deckThemeShare of the nonland cards.
func (b *Builder) DeckTheme(idx *cards.Index, deck []*mtgv1.Card) string {
	var spells []*mtgv1.Card
	words := map[string]bool{}
	for _, c := range deck {
		if isLand(c) {
			continue
		}
		spells = append(spells, c)
		if isCreatureCard(c) {
			for _, st := range c.GetSubtypes() {
				words[strings.ToLower(st)] = true
			}
		}
	}
	if len(spells) == 0 {
		return ""
	}
	for row := range b.themes.Themes {
		words[row] = true
	}
	ordered := make([]string, 0, len(words))
	for w := range words {
		ordered = append(ordered, w)
	}
	sort.Strings(ordered)
	best, bestSum := "", 0.0
	need := deckThemeShare * float64(len(spells))
	for _, w := range ordered {
		m := b.themes.match(w, idx.Tags())
		if m.Empty() {
			continue
		}
		var sum float64
		var fit int
		for _, c := range spells {
			if s, _ := m.score(c); s > 0 {
				sum += s
				fit++
			}
		}
		if float64(fit) >= need && sum > bestSum {
			best, bestSum = w, sum
		}
	}
	return best
}

func isLand(c *mtgv1.Card) bool {
	for _, t := range c.GetCardTypes() {
		if t == "Land" {
			return true
		}
	}
	return false
}
