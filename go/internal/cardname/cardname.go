// Package cardname holds the two keys a card name matches on (D-716).
// The exact key reads the name as written. The folded key also reads a
// letter without its accent, and a mark as the one a keyboard types, so
// a reader who types "Grima" reaches "Gríma Wormtongue".
package cardname

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// Exact is the exact key: lower case, with the outer spaces removed and
// a curly quote read as the straight one a card name holds.
func Exact(s string) string {
	return strings.ToLower(strings.TrimSpace(quotes.Replace(s)))
}

// Fold is the folded key: the exact key with each accent removed and
// each mark a keyboard lacks read as the one it types. A compatibility
// form, such as a full-width letter, reads as its plain letter. A lookup
// reads the folded key only when the exact key misses, and only when no
// two card names share the folded key (guardrail 4, D-716).
func Fold(s string) string {
	var b strings.Builder
	for _, r := range norm.NFKD.String(marks.Replace(s)) {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b.WriteRune(r)
	}
	return Exact(b.String())
}

// quotes reads a curly quote as the straight one.
var quotes = strings.NewReplacer("’", "'", "‘", "'", "“", "\"", "”", "\"")

// marks reads each mark that no decomposition removes as the mark a
// keyboard types. Card names hold three: the dash of "Human—Time Lord
// Meta-Crisis", the colon of "Ratonhnhaké꞉ton", and the registered sign
// of an Unglued card. Æ is the old spelling of Aether.
var marks = strings.NewReplacer(
	"—", "-", // em dash
	"–", "-", // en dash
	"꞉", ":", // modifier letter colon
	"®", "", // registered sign
	"Æ", "Ae", // Æ
	"æ", "ae", // æ
)
