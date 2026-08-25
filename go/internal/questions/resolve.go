package questions

import (
	"regexp"
	"strings"
)

// Hints supply the values the catalog rows name in {braces}. PR-6 answers
// the first two from the card index. A nil Hints means no value, and the
// clause that needs one is dropped.
//
// The live run of 2026-08-24 showed why this exists: the agent sent
// "{theme} is strongest in {colors}" to the model with both braces intact,
// and the model turned the agent's own statement into a second question
// aimed back at the user.
type Hints interface {
	// ThemeColors reads like "white and black". An empty string drops the
	// clause.
	ThemeColors(theme string) string
	// Commanders names up to three commanders for the theme and colors.
	Commanders(theme string) []string
	// OwnedThemeCount is the count of on-theme cards in the collection.
	// Zero drops the clause.
	OwnedThemeCount(theme string) int
}

var placeholder = regexp.MustCompile(`\{[a-z_]+\}`)

// resolve fills a row's placeholders, drops every sentence that still
// holds one, and falls back to the row's placeholder-free wording when
// what is left can not stand as a question.
func resolve(row Row, st *State, h Hints) (string, []string) {
	text, keptFirst := dropUnresolved(substitute(row.Text, st, h))
	// A trailing clause alone reads as a dangling question: drop the first
	// sentence of "{card} supports two plans... Which one do you want?" and
	// nothing says what "one" is. The first sentence carries the subject,
	// so the survivor only stands when the first sentence survived.
	if !usable(text) || !keptFirst {
		return strings.TrimSpace(row.Fallback), row.Options
	}
	return text, row.Options
}

// substitute replaces the placeholders whose values exist.
func substitute(text string, st *State, h Hints) string {
	theme := st.Slots.GetTheme()
	rep := map[string]string{}
	if theme != "" {
		rep["{theme}"] = theme
	}
	if len(st.NamedCards) > 0 {
		rep["{card}"] = st.NamedCards[0]
	}
	if h != nil {
		if c := strings.TrimSpace(h.ThemeColors(theme)); c != "" {
			rep["{colors}"] = c
		}
		names := h.Commanders(theme)
		for i, key := range []string{"{a}", "{b}", "{c}"} {
			if i < len(names) && strings.TrimSpace(names[i]) != "" {
				rep[key] = names[i]
			}
		}
		if n := h.OwnedThemeCount(theme); n > 0 {
			rep["{n}"] = itoa(n)
		}
	}
	for k, v := range rep {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}

// dropUnresolved removes each sentence that still holds a placeholder. A
// short question survives a missing clause: "Any color preference?" is
// better than a question about braces.
func dropUnresolved(text string) (kept string, keptFirst bool) {
	if !placeholder.MatchString(text) {
		return text, true
	}
	var keep []string
	for i, s := range splitSentences(text) {
		if placeholder.MatchString(s) {
			continue
		}
		if i == 0 {
			keptFirst = true
		}
		keep = append(keep, s)
	}
	return strings.TrimSpace(strings.Join(keep, " ")), keptFirst
}

// usable reports whether a resolved text can go out as one question.
func usable(text string) bool {
	text = strings.TrimSpace(text)
	return text != "" && !placeholder.MatchString(text) && strings.Contains(text, "?") && len(text) >= 12
}

// splitSentences cuts on sentence ends and keeps the punctuation.
func splitSentences(text string) []string {
	var out []string
	start := 0
	for i, r := range text {
		if r == '.' || r == '?' || r == '!' {
			out = append(out, strings.TrimSpace(text[start:i+1]))
			start = i + 1
		}
	}
	if rest := strings.TrimSpace(text[start:]); rest != "" {
		out = append(out, rest)
	}
	return out
}

// guard checks the model's phrasing before it reaches the user. A phrasing
// that fails goes back to the resolved catalog text (the live run of
// 2026-08-24 sent out a compound question the ask role invented).
func guard(phrased, resolved string) string {
	p := strings.TrimSpace(phrased)
	switch {
	case p == "":
	case strings.ContainsAny(p, "{}"):
	case strings.Count(p, "?") != 1:
	case !strings.Contains(p, "?"):
	case len(p) > 4*len(resolved)+120:
	default:
		return p
	}
	return resolved
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b [20]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}
