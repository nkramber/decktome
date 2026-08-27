package questions

import (
	"regexp"
	"strings"
	"unicode"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
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
	// skip holds the names the agent already offered, so a user who
	// answers "none" gets three new names (D-73).
	Commanders(theme string, skip []string) []string
	// OwnedThemeCount is the count of on-theme cards in the collection.
	// Zero drops the clause.
	OwnedThemeCount(theme string) int
}

// SlotAware lets a hint source read the slots as they stand inside the
// turn. The gate runner and agentsvc build the hint source before the
// turn runs, so a value the classifier has just filled is invisible to
// it. Conversation 23 of the batch run offered Jodah, the Unifier, a
// five-color commander, for a red-green deck, because the color slot
// filled in the same turn (D-124, the same class as D-82).
type SlotAware interface {
	UseSlots(format mtgv1.FormatId, colors []mtgv1.Color, pool mtgv1.PoolRule)
}

// PairAware lets a hint source know the user asked for a two-commander
// pair. Probe 73 writes "A Commander deck with a Background commander
// pair", and every run before D-154 answered it with three single
// legends, one of them off color.
type PairAware interface {
	UseWantPair(want, background bool)
}

// FactSource answers the planner facts that come from the card index and
// the collection. The classifier can not answer them, and each one gates
// a catalog row.
//
// agentsvc read them once before the turn, so a one-message answer such
// as "Modern red burn deck from my library" filled the format and the
// theme inside the turn, and the thin-theme count was never taken. The
// user got the plain pool question, and the {n} count of D-67 never
// showed. The agent reads them again after the classify call (M-6).
type FactSource interface {
	// WeakCommanderPool reports whether an owned mode holds no on-theme
	// commander (D-63).
	WeakCommanderPool(theme string) bool
	// ThinTheme reports whether the collection holds fewer than 30
	// on-theme cards, and the count (D-63).
	ThinTheme(theme string) (thin bool, count int)
}

// CommanderChecker reports whether a named card can lead a deck. A hint
// source that holds the card index implements it.
//
// known is false when the index does not hold the name. An unknown name
// is not proof of anything, and the agent claims nothing about it.
type CommanderChecker interface {
	CanLead(name string) (canLead, known bool)
}

// IdentityChecker reports whether a named card fits a set of colors. A
// hint source that holds the card index implements it.
//
// The pick row keeps the names on the table until the user refuses them
// (D-80, D-123). When the colors arrive after the offer, those names were
// never checked again. Probe 73 offered Jaheira, Friend of the Forest,
// which is mono-green, on turn 1 with no colors named. The user answered
// "Red and white" on turn 2, and the same three names went out on turns 2
// and 3 (D-153).
//
// known is false when the index does not hold the name. An unknown name
// is not proof of anything, and the agent drops nothing on it.
type IdentityChecker interface {
	FitsColors(name string, colors []mtgv1.Color) (fits, known bool)
}

var placeholder = regexp.MustCompile(`\{[a-z_]+\}`)

// resolve fills a row's placeholders, drops every sentence that still
// holds one, and falls back to the row's placeholder-free wording when
// what is left can not stand as a question.
//
// offered holds the commander names the resolved text names. The caller
// records them, so the next round of the pick row offers three others.
func resolve(row Row, st *State, h Hints) (text string, opts []string, offered []string) {
	filled, offered := substitute(row.Text, st, h)
	text, keptFirst := dropUnresolved(filled)
	// A trailing clause alone reads as a dangling question: drop the first
	// sentence of "{card} supports two plans... Which one do you want?" and
	// nothing says what "one" is. The first sentence carries the subject,
	// so the survivor only stands when the first sentence survived.
	if !usable(text) || !keptFirst {
		return strings.TrimSpace(row.Fallback), row.Options, nil
	}
	return text, pickOptions(row, offered), offered
}

// noneOption is the answer that asks for three other commanders (D-73).
const noneOption = "None, name three more"

// pickOptions builds the option list of a row that names commanders. The
// row carries no options of its own, because the names change every time.
//
// The pick row goes out as written (D-131), so the ask role no longer
// supplies its options. The names are the question, and the owner scored
// a replacement that dropped them as worse than the row (M-5 item 31).
func pickOptions(row Row, offered []string) []string {
	if len(offered) == 0 || len(row.Options) > 0 || len(commanderKeysIn(row.Text)) == 0 {
		return row.Options
	}
	out := make([]string, 0, len(offered)+1)
	out = append(out, offered...)
	return append(out, noneOption)
}

// vagueThemes are the theme values that name no archetype. They come
// from the competitive row (D-79), and they read as nonsense inside a
// statement: "the best deck under budget is strongest in white and
// blue". Gate run 3 of 2026-08-25 replaced three color questions for
// this reason.
var vagueThemes = []string{"best deck", "tier-one", "tier one", "strongest", "competitive", "under budget"}

// usefulTheme reports whether a theme can stand as the subject of a
// statement about colors.
func usefulTheme(theme string) bool {
	theme = strings.ToLower(strings.TrimSpace(theme))
	return theme != "" && !anyWord(theme, vagueThemes)
}

// substitute replaces the placeholders whose values exist. It returns the
// commander names it placed in the text.
func substitute(text string, st *State, h Hints) (string, []string) {
	theme := st.Slots.GetTheme()
	rep := map[string]string{}
	var offered []string
	if theme != "" {
		rep["{theme}"] = theme
	}
	if len(st.NamedCards) > 0 {
		rep["{card}"] = st.NamedCards[0]
	}
	// {locked} names the cards to keep, and never the commander (D-70).
	if s := englishList(st.LockedCards()); s != "" {
		rep["{locked}"] = s
	}
	// The illegal-commander row names the card the user asked for (D-129).
	if s := strings.TrimSpace(st.IllegalCommander); s != "" {
		rep["{bad_commander}"] = s
	}
	// The precon row names the deck the user wants to upgrade (D-113).
	if s := strings.TrimSpace(st.PreconName); s != "" {
		rep["{precon}"] = s
	}
	// The unsupported-format row names what the user asked for, and the
	// nearest format this app builds (D-112).
	if s := strings.TrimSpace(st.UnsupportedFormatName); s != "" {
		rep["{bad_format}"] = s
	}
	if s := strings.TrimSpace(st.NearestFormat); s != "" {
		rep["{near_format}"] = s
	}
	// Ask the hint source only for a value the row names. An eager call
	// runs a whole PR-6 build for a row that holds no placeholder, and it
	// fills the hint cache before the colors are known. The gate run of
	// 2026-08-25 offered three commanders outside the deck's colors for
	// exactly that reason.
	if h != nil {
		// No catalog row names {colors} since D-108: the colors row stated
		// which colors a theme is strongest in, and the claim was wrong in
		// gate run 13. The path stays live because PR-8 reads the same
		// source for the default color answer the corpus names.
		if strings.Contains(text, "{colors}") && usefulTheme(theme) {
			if c := strings.TrimSpace(h.ThemeColors(theme)); c != "" && strings.Count(c, ",") < 2 {
				rep["{colors}"] = c
			}
		}
		if names, keys := commanderKeysIn(text), 0; len(names) > 0 {
			// The names on the table stay on the table. A new set comes
			// only after the user asks for one (D-73).
			list := st.CurrentOffer
			if len(list) < len(names) {
				// Keep every name that still stands, and top the list up.
				// A full replacement would swap names the user could still
				// choose, which D-80 and D-123 forbid. The list is short of
				// names either because the user refused them all, which
				// empties it, or because D-153 dropped the ones the colors
				// exclude.
				for _, fresh := range h.Commanders(theme, st.OfferedCommanders) {
					if len(list) >= len(names) {
						break
					}
					if strings.TrimSpace(fresh) != "" && !hasName(list, fresh) {
						list = append(list, fresh)
					}
				}
			}
			for _, key := range names {
				if keys < len(list) && strings.TrimSpace(list[keys]) != "" {
					rep[key] = list[keys]
					offered = append(offered, list[keys])
				}
				keys++
			}
		}
		if strings.Contains(text, "{n}") {
			if n := h.OwnedThemeCount(theme); n > 0 {
				rep["{n}"] = itoa(n)
			}
		}
	}
	for k, v := range rep {
		text = strings.ReplaceAll(text, k, v)
	}
	return text, offered
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

// MaxRewordOverlap is how much of a replacement may repeat the row it
// replaces. Above it, the replacement is a reword, not a new question.
//
// Gate run 4 of 2026-08-25 measured a median word overlap of 0.67 over
// eight replacements. Seven of the eight only added the user's colors,
// format, or card name, which the ask role adds anyway. The score prompt
// already says that is not a fault (D-88).
const MaxRewordOverlap = 0.6

// MinRowCoverage and MaxBorrowed catch the other shape of a reword: a
// truncation. Overlap is symmetric, so a replacement that deletes half
// the row scores low and passes, although it says strictly less.
//
// Item 8 of the M-5 sheet is the case. The row read "What do people play
// at your event? I tune the 15 sideboard cards to it." The replacement
// read "What do people play at your Modern event?" It dropped the
// sentence that says what the answer is for. It scored 0.44 overlap,
// well under the reword bar, and the owner scored it worse than the row
// (D-103).
//
// A truncation borrows nearly all its words from the row and keeps only
// a part of it. A genuinely different question borrows few words, so it
// passes both tests.
const (
	MinRowCoverage = 0.6
	MaxBorrowed    = 0.8
)

// nearCopy reports whether a replacement says no more than the row it
// replaces. It catches two shapes: a reword that repeats most of the row,
// and a truncation that drops most of it.
func nearCopy(replacement, row string) bool {
	return overlap(replacement, row) > MaxRewordOverlap || truncation(replacement, row)
}

// truncation reports whether a replacement is the row with parts removed.
func truncation(replacement, row string) bool {
	borrowed, covered := share(replacement, row)
	return borrowed >= MaxBorrowed && covered < MinRowCoverage
}

// share returns two numbers. borrowed is the part of the replacement that
// comes from the row, and covered is the part of the row the replacement
// keeps.
func share(replacement, row string) (borrowed, covered float64) {
	wa, wb := words(replacement), words(row)
	if len(wa) == 0 || len(wb) == 0 {
		return 0, 0
	}
	both := 0
	for w := range wa {
		if wb[w] {
			both++
		}
	}
	return float64(both) / float64(len(wa)), float64(both) / float64(len(wb))
}

// Overlap is the share of words two texts have in common, from 0 to 1.
// cmd/m5-report reads it to tell a replacement that copies the row word
// for word from one that says something else (D-116).
func Overlap(a, b string) float64 { return overlap(a, b) }

// overlap is the share of words the two texts have in common, from 0 to 1.
func overlap(a, b string) float64 {
	wa, wb := words(a), words(b)
	if len(wa) == 0 || len(wb) == 0 {
		return 0
	}
	both := 0
	for w := range wa {
		if wb[w] {
			both++
		}
	}
	either := len(wa) + len(wb) - both
	if either == 0 {
		return 0
	}
	return float64(both) / float64(either)
}

// words splits a text into its lowercase word set. Punctuation is
// dropped, so "plan?" and "plan." are one word.
func words(text string) map[string]bool {
	out := map[string]bool{}
	for _, w := range strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
	}) {
		out[w] = true
	}
	return out
}

// guard checks the model's phrasing before it reaches the user. A phrasing
// that fails goes back to the resolved catalog text (the live run of
// 2026-08-24 sent out a compound question the ask role invented).
func guard(rowID, phrased, resolved string) string {
	p := strings.TrimSpace(phrased)
	switch {
	case p == "":
	case strings.ContainsAny(p, "{}"):
	case strings.Count(p, "?") != 1:
	case !strings.Contains(p, "?"):
	case len(p) > 4*len(resolved)+120:
	case namesUnsupportedFormat(rowID, p):
	case possessiveIdentity.MatchString(p):
	default:
		return p
	}
	return resolved
}

// The guard also refuses a phrasing that names one card's color identity,
// which possessiveIdentity reads. Gate runs 14, 15, and 16 each sent one,
// and the model changed the preposition every time D-144 caught it
// (D-151).

// namesUnsupportedFormat reports whether a phrasing names a format this
// app does not build. The ask role fits a question to the user's words,
// and it fitted the format the agent had just declined.
//
// Gate run 15 asked "What should the Oathbreaker deck focus on?" one line
// under "I do not build Oathbreaker", and it did the same for Historic.
// Both went out in the same turn, so the agent contradicted itself inside
// one message. The linter refuses the shape, and it failed run 15 (D-150).
//
// A row that exists to decline such a format is exempt, because naming it
// is the whole job of that row. Those rows carry `fixed` and never reach
// the ask role today, so the exemption is a guard against a later change
// (D-112, D-117).
func namesUnsupportedFormat(rowID, phrased string) bool {
	if declinesFormat(rowID) {
		return false
	}
	low := strings.ToLower(phrased)
	for _, u := range unsupported {
		if strings.Contains(low, u.phrase) {
			return true
		}
	}
	return false
}

// commanderKeysIn lists the commander placeholders a row holds, in the
// order the values fill them.
func commanderKeysIn(text string) []string {
	var out []string
	for _, key := range []string{"{a}", "{b}", "{c}"} {
		if strings.Contains(text, key) {
			out = append(out, key)
		}
	}
	return out
}

// englishList reads a list as English: "one, two, and three".
func englishList(items []string) string {
	var kept []string
	for _, s := range items {
		if s = strings.TrimSpace(s); s != "" {
			kept = append(kept, s)
		}
	}
	switch len(kept) {
	case 0:
		return ""
	case 1:
		return kept[0]
	case 2:
		return kept[0] + " and " + kept[1]
	default:
		return strings.Join(kept[:len(kept)-1], ", ") + ", and " + kept[len(kept)-1]
	}
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
