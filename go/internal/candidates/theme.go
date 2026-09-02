package candidates

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
)

//go:embed themes.json
var themesJSON []byte

// themeRow splits signals into payoffs (cards that reward the theme) and
// enablers (cards that merely do the thing). Payoffs weigh more (D-62).
type themeRow struct {
	PayoffSlugs []string `json:"payoff_slugs"`
	PayoffText  []string `json:"payoff_text"`
	Slugs       []string `json:"slugs"`
	Keywords    []string `json:"keywords"`
	Text        []string `json:"text"`
	Subtype     string   `json:"subtype"`
}

type themeTable struct {
	VerifiedAt string              `json:"verified_at"`
	Themes     map[string]themeRow `json:"themes"`
	Roles      map[string][]string `json:"roles"`
}

func loadThemes() (*themeTable, error) {
	var t themeTable
	if err := json.Unmarshal(themesJSON, &t); err != nil {
		return nil, fmt.Errorf("candidates: themes.json: %w", err)
	}
	if t.VerifiedAt == "" {
		return nil, fmt.Errorf("candidates: themes.json has no verified_at")
	}
	return &t, nil
}

// Signal weights (D-62). A payoff tag is the strongest signal. An enabler
// keyword is weak: hundreds of cards have lifelink, few reward lifegain.
// Text is weaker than a tag of the same kind: a needle can match by
// accident. The cap is scoreCap, so payoff plus enabler is a full score.
const (
	weightPayoffTag  = 1.5
	weightPayoffText = 1.2
	weightTag        = 1.0
	weightSubtype    = 0.8
	weightKeyword    = 0.5
	weightText       = 0.4
	scoreCap         = 2.5
)

// ThemeMatch is the resolved theme, for logs and the gate doc.
type ThemeMatch struct {
	// Words are the normalized theme words in input order.
	Words []string
	// PayoffSlugs and Slugs are the Tagger slugs that exist for the words.
	PayoffSlugs []string
	Slugs       []string
	// PayoffText lists the payoff needles.
	PayoffText []string
	// Keywords, Subtypes, and Text are the other signals in use.
	Keywords []string
	Subtypes []string
	Text     []string
	// Unmatched lists words that fired on no card of the pool. Build
	// fills it after the scan: a needle is a guess until a card holds it.
	Unmatched []string

	tagged map[string]map[string]bool // slug -> oracle ids
	// wordOf maps a signal, as score emits it, to the theme word that
	// produced it. Build reads it to fill Unmatched.
	wordOf map[string]string
	// generic lists the text needles the generic rule made from words
	// the table does not know. pruneNoisy reads it (D-411).
	generic []string
}

// noisyShare is the share of the card database above which a generic
// text needle discriminates nothing (D-411). A needle that sits in the
// text of one card in ten ranks the pool on chance, and "you" sat in
// the text of most commanders: the offer for "the best deck you can"
// was the three most popular legends whose text held "you" and "can".
const noisyShare = 0.10

// matchIn turns the user's words into signals and drops every generic
// text needle the card database makes noise of (D-411). Build and
// CommanderPool read this, and match alone stays for the tests that
// read the raw needles.
func (t *themeTable) matchIn(theme string, idx *cards.Index) ThemeMatch {
	m := t.match(theme, idx.Tags())
	m.pruneNoisy(idx.All())
	return m
}

// pruneNoisy removes each generic text needle that more than noisyShare
// of the cards hold. One pass lowers each text once, and every needle
// reads that one string.
func (m *ThemeMatch) pruneNoisy(all []*mtgv1.Card) {
	if len(m.generic) == 0 || len(all) == 0 {
		return
	}
	hits := make(map[string]int, len(m.generic))
	for _, c := range all {
		text := strings.ToLower(c.GetOracleText())
		for _, n := range m.generic {
			if strings.Contains(text, n) {
				hits[n]++
			}
		}
	}
	limit := int(noisyShare * float64(len(all)))
	for _, n := range m.generic {
		if hits[n] > limit {
			m.Text = slices.DeleteFunc(m.Text, func(x string) bool { return x == n })
		}
	}
}

// firedSignals is the set of signals that fired on at least one card.
type firedSignals map[string]bool

func (f firedSignals) mark(signals []string) {
	for _, s := range signals {
		f[s] = true
	}
}

// unmatchedWords returns the theme words, in input order, that no fired
// signal traces back to. A word with no needle at all is unmatched too.
func (m ThemeMatch) unmatchedWords(fired firedSignals) []string {
	hit := map[string]bool{}
	for s := range fired {
		if w, ok := m.wordOf[s]; ok {
			hit[w] = true
		}
	}
	var out []string
	for _, w := range m.Words {
		if !hit[w] {
			out = append(out, w)
		}
	}
	return out
}

// match turns the user's words into signals. Each word maps through the
// table, or through the generic rule when the table has no row.
func (t *themeTable) match(theme string, tags *cards.TagIndex) ThemeMatch {
	var m ThemeMatch
	m.tagged = map[string]map[string]bool{}
	m.wordOf = map[string]string{}
	seenSlug := map[string]bool{}
	// The current word, so each signal remembers where it came from. The
	// first word that produced a signal keeps it.
	word := ""
	trace := func(signal string) {
		if _, ok := m.wordOf[signal]; !ok {
			m.wordOf[signal] = word
		}
	}
	addSlug := func(slug string, payoff bool) bool {
		if seenSlug[slug] || !tags.Has(slug) {
			return false
		}
		seenSlug[slug] = true
		set := map[string]bool{}
		for _, id := range tags.Resolve(slug) {
			set[id] = true
		}
		m.tagged[slug] = set
		if payoff {
			m.PayoffSlugs = append(m.PayoffSlugs, slug)
			trace("payoff:" + slug)
		} else {
			m.Slugs = append(m.Slugs, slug)
			trace("tag:" + slug)
		}
		return true
	}
	addNeedle := func(list *[]string, kind, n string) {
		*list = appendUnique(*list, n)
		trace(kind + ":" + n)
	}
	for _, w := range t.words(theme) {
		m.Words = append(m.Words, w)
		word = w
		if row, ok := t.Themes[w]; ok {
			for _, s := range row.PayoffSlugs {
				addSlug(s, true)
			}
			for _, n := range row.PayoffText {
				addNeedle(&m.PayoffText, "payoff-text", strings.ToLower(n))
			}
			for _, s := range row.Slugs {
				addSlug(s, false)
			}
			for _, k := range row.Keywords {
				addNeedle(&m.Keywords, "keyword", k)
			}
			if row.Subtype != "" {
				addNeedle(&m.Subtypes, "subtype", row.Subtype)
			}
			for _, n := range row.Text {
				addNeedle(&m.Text, "text", strings.ToLower(n))
			}
		} else {
			// Generic rule: payoff slugs w-matters and synergy-w, enabler
			// slugs w and typal-w, then keyword, subtype, and text.
			for _, s := range []string{w + "-matters", "synergy-" + w, "typal-" + w, "typal-" + singular(w)} {
				addSlug(s, true)
			}
			addSlug(w, false)
			addNeedle(&m.Keywords, "keyword", title(w))
			addNeedle(&m.Subtypes, "subtype", title(singular(w)))
			addNeedle(&m.Text, "text", w)
			m.generic = append(m.generic, w)
		}
	}
	// Build fills Unmatched after it scans the pool.
	return m
}

// score returns the theme fit of a card and the signals that fired.
func (m ThemeMatch) score(c *mtgv1.Card) (float64, []string) {
	var score float64
	var signals []string
	text := strings.ToLower(c.OracleText)
	// Tags in one hierarchy overlap (lifegain, repeatable-lifegain,
	// drain-life), so a kind counts once no matter how many slugs fire.
	hit := false
	for _, slug := range m.PayoffSlugs {
		if m.tagged[slug][c.OracleId] {
			hit = true
			signals = append(signals, "payoff:"+slug)
		}
	}
	if hit {
		score += weightPayoffTag
	}
	hit = false
	for _, n := range m.PayoffText {
		if n != "" && strings.Contains(text, n) {
			hit = true
			signals = append(signals, "payoff-text:"+n)
		}
	}
	if hit {
		score += weightPayoffText
	}
	hit = false
	for _, slug := range m.Slugs {
		if m.tagged[slug][c.OracleId] {
			hit = true
			signals = append(signals, "tag:"+slug)
		}
	}
	if hit {
		score += weightTag
	}
	// A keyword that fired also covers its own name as a text needle:
	// "lifelink" the keyword and "lifelink" the word are one fact.
	covered := map[string]bool{}
	for _, k := range m.Keywords {
		if hasKeywordFold(c.Keywords, k) {
			score += weightKeyword
			signals = append(signals, "keyword:"+k)
			covered[strings.ToLower(k)] = true
		}
	}
	for _, st := range m.Subtypes {
		if slices.Contains(c.Subtypes, st) {
			score += weightSubtype
			signals = append(signals, "subtype:"+st)
		}
	}
	for _, n := range m.Text {
		if n == "" || covered[n] || !strings.Contains(text, n) {
			continue
		}
		score += weightText
		signals = append(signals, "text:"+n)
	}
	if score > scoreCap {
		score = scoreCap
	}
	return score / scoreCap, signals
}

// roleSets resolves the role slugs to Oracle-id sets.
func (t *themeTable) roleSets(tags *cards.TagIndex) map[string]map[string]bool {
	out := map[string]map[string]bool{}
	for role, slugs := range t.Roles {
		set := map[string]bool{}
		for _, s := range slugs {
			for _, id := range tags.Resolve(s) {
				set[id] = true
			}
		}
		out[role] = set
	}
	return out
}

// minWordLen is the shortest token that can be a theme word. A shorter
// token is a fragment: "+1/+1 counters" splits into "1", "1", and
// "counters", and the needle "1" matches every card with a digit.
const minWordLen = 3

// words normalizes the theme: lowercase, split on space and punctuation,
// stop words dropped, order kept, duplicates dropped. Two tokens that
// name a hyphenated row join first, so "go wide" finds the go-wide row.
// Then a token shorter than minWordLen or made of digits goes.
func (t *themeTable) words(theme string) []string {
	// A letter of any script is part of a word, so a non-ASCII theme word
	// stays whole.
	f := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-'
	}
	var tokens []string
	for _, w := range strings.FieldsFunc(strings.ToLower(theme), f) {
		if w = strings.Trim(w, "-"); w != "" {
			tokens = append(tokens, w)
		}
	}
	tokens = t.joinRows(tokens)
	var out []string
	seen := map[string]bool{}
	for _, w := range tokens {
		if utf8.RuneCountInString(w) < minWordLen || allDigits(w) || stopWords[w] || seen[w] {
			continue
		}
		seen[w] = true
		out = append(out, w)
	}
	return out
}

// joinRows joins two neighbor tokens when the table has a row under
// their hyphenated form. The joined word replaces both tokens.
func (t *themeTable) joinRows(tokens []string) []string {
	var out []string
	for i := 0; i < len(tokens); i++ {
		if i+1 < len(tokens) {
			joined := tokens[i] + "-" + tokens[i+1]
			if _, ok := t.Themes[joined]; ok {
				out = append(out, joined)
				i++
				continue
			}
		}
		out = append(out, tokens[i])
	}
	return out
}

// allDigits reports whether a token is a number.
func allDigits(w string) bool {
	for _, r := range w {
		if r < '0' || r > '9' {
			return false
		}
	}
	return w != ""
}

// stopWords are the words of a request that name no theme. The second
// group is the words of "build me the best deck you can", which reached
// the generic rule and became text needles before D-411.
var stopWords = map[string]bool{
	"a": true, "an": true, "the": true, "and": true, "or": true, "of": true, "with": true, "deck": true,
	"build": true, "me": true, "my": true, "for": true, "in": true, "on": true, "to": true, "some": true,
	"commander": true, "edh": true, "please": true, "want": true, "i": true, "that": true, "this": true,
	"cards": true, "card": true, "fun": true, "good": true, "strong": true, "casual": true,
	"you": true, "your": true, "can": true, "could": true, "would": true, "should": true, "make": true,
	"give": true, "need": true, "like": true, "best": true, "possible": true, "great": true, "really": true,
	"very": true, "just": true, "something": true, "anything": true, "decks": true, "play": true,
	"powerful": true, "competitive": true, "optimal": true, "optimized": true, "strongest": true,
}

// singularIE lists plurals in -ies whose singular ends in -ie. The rule
// below turns -ies into -y, which is right for armies and harpies and
// wrong for these creature types.
var singularIE = map[string]bool{
	"zombies": true, "faeries": true, "pixies": true, "genies": true,
}

// singular gives the singular of a theme word, for a subtype or a slug.
func singular(w string) string {
	switch {
	case singularIE[w]:
		return strings.TrimSuffix(w, "s")
	case strings.HasSuffix(w, "ies"):
		return strings.TrimSuffix(w, "ies") + "y"
	case strings.HasSuffix(w, "ves"):
		return strings.TrimSuffix(w, "ves") + "f"
	case strings.HasSuffix(w, "s") && !strings.HasSuffix(w, "ss"):
		return strings.TrimSuffix(w, "s")
	}
	return w
}

func title(w string) string {
	if w == "" {
		return w
	}
	return strings.ToUpper(w[:1]) + w[1:]
}

func appendUnique(list []string, v string) []string {
	if v == "" || slices.Contains(list, v) {
		return list
	}
	return append(list, v)
}

func hasKeywordFold(have []string, want string) bool {
	for _, k := range have {
		if strings.EqualFold(k, want) {
			return true
		}
	}
	return false
}

// Describe renders the match for the gate doc.
func (m ThemeMatch) Describe() string {
	var parts []string
	if len(m.PayoffSlugs) > 0 {
		s := append([]string(nil), m.PayoffSlugs...)
		sort.Strings(s)
		parts = append(parts, "payoff tags "+strings.Join(s, ", "))
	}
	if len(m.Slugs) > 0 {
		s := append([]string(nil), m.Slugs...)
		sort.Strings(s)
		parts = append(parts, "tags "+strings.Join(s, ", "))
	}
	if len(m.Keywords) > 0 {
		parts = append(parts, "keywords "+strings.Join(m.Keywords, ", "))
	}
	if len(m.Subtypes) > 0 {
		parts = append(parts, "subtypes "+strings.Join(m.Subtypes, ", "))
	}
	if len(m.Unmatched) > 0 {
		parts = append(parts, "unmatched "+strings.Join(m.Unmatched, ", "))
	}
	if len(parts) == 0 {
		return "no theme signal"
	}
	return strings.Join(parts, ". ")
}
