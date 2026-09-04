package questions

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// Hints supply the values the catalog rows name in {braces}. PR-6 answers
// the commander names from the card index. A nil Hints means no value, and the
// clause that needs one is dropped.
//
// A brace that reaches the model comes back as a question aimed at the
// user, so every placeholder is filled or its clause is dropped (D-82).
type Hints interface {
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
// turn runs, so a value the classifier has just filled would otherwise
// be invisible to it, and the offer could name a commander outside the
// colors (D-124, the same class as D-82).
type SlotAware interface {
	UseSlots(format mtgv1.FormatId, colors []mtgv1.Color, pool mtgv1.PoolRule)
}

// SetAware lets a hint source read the set limit as it stands inside the
// turn, so the commander offer stays inside the sets (D-437).
type SetAware interface {
	UseSets(codes []string)
}

// PairAware lets a hint source know the user asked for a two-commander
// pair. Three single legends do not answer that request (D-154).
type PairAware interface {
	UseWantPair(want, background bool)
}

// FactSource answers the planner facts that come from the card index and
// the collection. The classifier can not answer them, and each one gates
// a catalog row.
//
// agentsvc reads them once before the turn, so a one-message answer such
// as "Modern red burn deck from my library" fills the format and the
// theme inside the turn. The agent reads them again after the classify
// call, so the thin-theme count is taken (D-198).
type FactSource interface {
	// ThinTheme reports whether the collection holds fewer than 30
	// on-theme cards, and the count (D-63).
	ThinTheme(theme string) (thin bool, count int)
}

// SetResolver maps the words a reader wrote onto a set family. A hint
// source that holds the card index implements it (D-376).
type SetResolver interface {
	// ResolveSet returns the family codes and the set names for a phrase.
	// ok is false when the phrase names no set or names two base sets,
	// and options then holds the names the question offers.
	ResolveSet(phrase string) (codes, names, options []string, ok bool)
	// ResolveSetGroup returns every family a franchise word names, as
	// codes and names, for a group request such as "sets with Marvel
	// characters" (D-525). ok is false when the word names no set.
	ResolveSetGroup(phrase string) (codes, names []string, ok bool)
}

// PreconMatch is the answer to ResolvePrecon.
type PreconMatch struct {
	// Products are the products the phrase named, by key and by name. A
	// deck and its Collector's Edition both answer, and the exclusion
	// counts them once.
	Products []PreconRef
	// Partial names the products of Products the collection does not
	// hold whole, when a collection is attached (D-497).
	Partial []string
	// Options are the names the question offers when OK is false.
	Options []string
	// OK is false when the phrase names no product, or names products
	// with different cards.
	OK bool
}

// PreconResolver maps the words a reader wrote onto the products of the
// precon table (D-496). A hint source that holds the table implements it.
type PreconResolver interface {
	ResolvePrecon(phrase string) PreconMatch
}

// OwnedPreconSource lists the precon products the collection holds whole
// (D-408). ok is false with no table or no collection, and the turn then
// says it excluded nothing.
type OwnedPreconSource interface {
	OwnedPrecons() (owned []PreconRef, ok bool)
}

// ManaSource counts the mana cards a set family offers against the count
// the deck wants (D-382). A hint source that holds the card index
// implements it.
type ManaSource interface {
	// ThinSetMana reports whether the sets hold fewer mana cards than the
	// deck wants, with the two counts. want is zero when the source can
	// not answer, and the clause that names the counts then drops.
	ThinSetMana(codes []string, format mtgv1.FormatId, colors []mtgv1.Color, power *mtgv1.PowerLevel) (thin bool, have, want int)
}

// CommanderChecker reports whether a named card can lead a deck. A hint
// source that holds the card index implements it.
//
// known is false when the index does not hold the name. An unknown name
// is not proof of anything, and the agent claims nothing about it.
type CommanderChecker interface {
	CanLead(name string) (canLead, known bool)
}

// FormatChecker reports whether a named card can only be played in
// Commander. A hint source that holds the card index implements it.
//
// A legendary creature is not proof of the format: many of them are
// legal in Standard or Modern too. A card that can lead a deck and is
// legal in neither leaves Commander as the one format this app builds
// for it (D-388).
type FormatChecker interface {
	// OnlyCommander returns whether the card can only be played in
	// Commander here. known is false when the index does not hold the
	// name, and the caller then claims nothing.
	OnlyCommander(name string) (only, known bool)
}

// IdentityChecker reports whether a named card fits a set of colors. A
// hint source that holds the card index implements it.
//
// The pick row keeps the names on the table until the user refuses them
// (D-80, D-123). When the colors arrive after the offer, those names
// must be checked again, or an off-color name stays on the table
// (D-153).
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
	return text, rowOptions(row, st, offered), offered
}

// rowOptions builds the option list of one resolved row. The set row
// offers the sets its text names, so the reader picks one instead of
// typing a name the app already listed (D-376). Every other row goes
// through pickOptions.
func rowOptions(row Row, st *State, offered []string) []string {
	if row.StateKey() == SlotSetUnresolved && len(st.SetOptions) > 0 {
		return append(append([]string(nil), st.SetOptions...), everySetOption)
	}
	// A reader who asked for a 60-card deck can not mean Commander, and
	// Commander is a 100-card format. The row offered it anyway, and
	// eval run 31 called the question inaccurate (D-388).
	if row.Slot == "format" && sixtyCardRequest(st.Ctx.Words) {
		return withoutCommander(row.Options)
	}
	return pickOptions(row, offered)
}

// withoutCommander drops Commander from an option list. An empty result
// never reaches the reader: the caller keeps the original list then.
func withoutCommander(options []string) []string {
	out := make([]string, 0, len(options))
	for _, o := range options {
		if strings.EqualFold(strings.TrimSpace(o), "commander") {
			continue
		}
		out = append(out, o)
	}
	if len(out) == 0 {
		return options
	}
	return out
}

// everySetOption drops the set limit. A reader who meant no set at all
// needs a way to say so, and the row is closed to free text.
const everySetOption = "Use every set"

// noneOption is the answer that asks for three other commanders (D-73).
const noneOption = "None, name three more"

// pickOptions builds the option list of a row that names commanders. The
// row carries no options of its own, because the names change every time.
//
// The pick row goes out as written (D-131), so the ask role no longer
// supplies its options. The names are the question, and a replacement
// that drops them is worse than the row (D-66).
func pickOptions(row Row, offered []string) []string {
	if len(offered) == 0 || len(row.Options) > 0 || len(commanderKeysIn(row.Text)) == 0 {
		return row.Options
	}
	out := make([]string, 0, len(offered)+1)
	out = append(out, offered...)
	return append(out, noneOption)
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
	// The set row names the phrase this app could not settle, and the
	// sets it may have meant (D-376).
	if v := strings.TrimSpace(st.UnresolvedSet); v != "" {
		rep["{bad_set}"] = v
	}
	// The row names every set the phrase could mean. The count is not
	// fixed, so one placeholder carries the whole list: three fixed
	// placeholders dropped the sentence whenever a phrase named two sets,
	// and the question then named none of them.
	if s := orList(st.SetOptions); s != "" {
		rep["{set_options}"] = s
	}
	// The mana row names the sets the reader chose (D-382).
	if s := englishList(st.SetNames); s != "" {
		rep["{sets}"] = s
	}
	// The precon row names the phrase the table could not settle, and
	// the products it may have meant (D-496).
	if v := strings.TrimSpace(st.UnresolvedPrecon); v != "" {
		rep["{bad_precon}"] = v
	}
	if s := orList(st.PreconOptions); s != "" {
		rep["{precon_options}"] = s
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
	// fills the hint cache before the colors are known (D-82).
	if h != nil {
		// No catalog row names {colors} since D-108: the colors row stated
		// which colors a theme is strongest in, and the claim was wrong.
		// The hint that answered it is gone, so a {colors} clause drops.
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
				rep["{n}"] = strconv.Itoa(n)
			}
		}
		// The mana row names the count the sets hold and the count the
		// deck wants (D-382). A source that can not count drops both, and
		// the fallback carries the question.
		if ms, ok := h.(ManaSource); ok && strings.Contains(text, "{n_mana}") {
			if _, have, want := ms.ThinSetMana(st.Slots.GetSetCodes(), st.Slots.GetFormat().GetId(),
				st.Slots.GetColors(), st.Slots.GetPower()); want > 0 {
				rep["{n_mana}"] = strconv.Itoa(have)
				rep["{m_mana}"] = strconv.Itoa(want)
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
// A replacement that only adds the user's colors, format, or card name
// repeats the row, and the ask role adds those words anyway. The score
// prompt already says that is not a fault (D-88).
const MaxRewordOverlap = 0.6

// MinRowCoverage and MaxBorrowed catch the other shape of a reword: a
// truncation. Overlap is symmetric, so a replacement that deletes half
// the row scores low and passes, although it says strictly less.
//
// A row that reads "What do people play at your event? I tune the 15
// sideboard cards to it." loses its point in the replacement "What do
// people play at your Modern event?". The overlap is low, and the
// replacement says strictly less (D-103).
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
// that fails goes back to the resolved catalog text, so a compound
// question the ask role invents never goes out (D-116).
func guard(rowID, phrased, resolved string) string {
	p := strings.TrimSpace(phrased)
	switch {
	case p == "":
	case strings.ContainsAny(p, "{}"):
	case strings.Count(p, "?") != 1:
	case len(p) > 4*len(resolved)+120:
	case namesUnsupportedFormat(rowID, p):
	case possessiveIdentity.MatchString(p):
	default:
		return p
	}
	return resolved
}

// The guard also refuses a phrasing that names one card's color identity,
// which possessiveIdentity reads. The model keeps the shape and changes
// the preposition, so the rule reads the shape (D-151).

// namesUnsupportedFormat reports whether a phrasing names a format this
// app does not build. The ask role fits a question to the user's words,
// and it fitted the format the agent had just declined.
//
// "What should the Oathbreaker deck focus on?" one line under "I do not
// build Oathbreaker" contradicts the agent inside one message. The
// linter refuses the shape (D-150).
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
// orList joins the choices of a question with "or", because a reader
// picks one of them. englishList joins with "and", for the things a deck
// holds together (D-518).
func orList(items []string) string {
	return joinList(items, "or")
}

func englishList(items []string) string {
	return joinList(items, "and")
}

func joinList(items []string, word string) string {
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
		return kept[0] + " " + word + " " + kept[1]
	default:
		return strings.Join(kept[:len(kept)-1], ", ") + ", " + word + " " + kept[len(kept)-1]
	}
}
