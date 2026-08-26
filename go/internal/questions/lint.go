package questions

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// The question linter (D-115). Every rule here reads text alone, so it
// costs no model call and it can not drift between runs.
//
// The rules come from the owner's M-5 scoring of items 1 to 32 on
// 2026-08-25. Measured over the 793 questions of gate runs 10 to 13, they
// find a defect in 61 questions, and in 159 when the presumed-table rule
// counts. The linter can not judge tone or relevance. The owner's scoring
// still owns that.

// LintQuestion is one question as it went out.
type LintQuestion struct {
	// Turn counts from 1. The linter reads the messages before it.
	Turn  int
	RowID string
	Slot  string
	Text  string
}

// Finding is one defect the linter found.
type Finding struct {
	Rule   string
	Turn   int
	RowID  string
	Text   string
	Detail string
}

// String reads one finding as a line for a document.
func (f Finding) String() string {
	return fmt.Sprintf("turn %d, row `%s`, %s: %s", f.Turn, f.RowID, f.Rule, f.Detail)
}

// tableWords are the words a question uses for a group of players. A
// question may use one only after the user has.
var tableWords = []string{"your table", "your playgroup", "your group", "your event", "your pod"}

// tableEvidence are the words that let a question name a table. The user
// must have written one of them first.
var tableEvidence = []string{
	"table", "playgroup", "group", "pod", "event", "shop", "store",
	"fnm", "lgs", "rcq", "tournament", "friends", "we play", "our",
}

// illegalOffers are questions that offer an answer the rules forbid. The
// color identity of a commander is the color identity of the deck, and
// nothing may offer to leave it. Gate run 14 asked "Do you want to use
// any colors beyond Grist's color identity?" (D-144).
var illegalOffers = []string{
	"beyond the color identity", "beyond its color identity",
	"colors beyond", "colors outside", "outside the color identity",
	"outside its color identity", "additional colors beyond",
}

// possessiveIdentity finds a clause that names one card's color identity,
// such as "within Grist's color identity".
//
// D-144 refused this shape as a list of prepositions: "colors beyond",
// "outside the color identity". Gate run 14 said "beyond Grist's color
// identity", and the rule caught it. Runs 15 and 16 then said "within
// Grist's color identity", and the rule did not. The model kept the
// shape and changed the preposition, so the rule now reads the shape.
//
// No legitimate case exists. Every row that asks about the colors
// carries `commander_set: false`, so it fires only when no commander is
// settled. A card's color identity therefore settles nothing, and the
// question presumes that the card leads the deck. Conversation 3 is the
// case: the user wrote "Build around Grist, the Hunger Tide", which
// names a card and no role at all (D-118). The format was still open in
// the same turn, and color identity is a Commander term (D-151).
var possessiveIdentity = regexp.MustCompile(`(?i)['\x{2019}]s\s+color\s+identity`)

// factClaims are the shapes of a claim about the game. A question states
// no fact: another step owns that, and a wrong claim costs trust (D-108).
var factClaims = []string{"strongest in", "is strongest", "are strongest", "strongest for", "best in"}

// stutterStops are words that repeat in normal English. A repeat of one
// of them is not a stuttered card name.
var stutterStops = map[string]bool{
	"the": true, "and": true, "that": true, "this": true, "your": true,
	"with": true, "from": true, "into": true, "them": true, "deck": true,
}

// stutteredName finds a capitalized name the text repeats after "and",
// such as "Grist, the Hunger Tide and Grist". Gate runs 10 to 12 sent
// that question three times, because the short name and the full name
// both reached the locked list.
//
// Go regular expressions hold no back reference, so this scans the words.
func stutteredName(text string) (string, bool) {
	if i := strings.Index(text, "?"); i >= 0 {
		text = text[:i]
	}
	fields := strings.Fields(text)
	seen := map[string]bool{}
	for i, raw := range fields {
		word := strings.Trim(raw, ",.:!?;\"'")
		if word == "" {
			continue
		}
		prev := ""
		if i > 0 {
			prev = strings.Trim(fields[i-1], ",.:!?;\"'")
		}
		if strings.EqualFold(prev, "and") && seen[word] && nameLike(word) {
			return word, true
		}
		seen[word] = true
	}
	return "", false
}

// nameLike reports whether a word can be part of a card name.
func nameLike(word string) bool {
	if len(word) < 4 || stutterStops[strings.ToLower(word)] {
		return false
	}
	r := []rune(word)[0]
	return unicode.IsUpper(r)
}

// LintConversation checks the questions of one conversation against the
// messages that came before each of them.
func LintConversation(messages []string, qs []LintQuestion) []Finding {
	var out []Finding
	add := func(q LintQuestion, rule, detail string) {
		out = append(out, Finding{Rule: rule, Turn: q.Turn, RowID: q.RowID, Text: q.Text, Detail: detail})
	}
	for _, q := range qs {
		prior := priorWords(messages, q.Turn)
		lower := strings.ToLower(q.Text)

		// The user named the format and the agent asked for it anyway.
		if q.Slot == "format" && !declinesFormat(q.RowID) {
			if id, ok := FormatFromWords(prior); ok {
				add(q, "format_already_named", fmt.Sprintf("the user named %s before this question", id.String()))
			}
		}
		// The question presumes a group of players the user never named.
		for _, w := range tableWords {
			if strings.Contains(lower, w) && !anyPlain(strings.ToLower(prior), tableEvidence) {
				add(q, "presumes_a_table", fmt.Sprintf("the question says %q and the user named no table", w))
				break
			}
		}
		// The question offers a format this app can not build. The row that
		// declines such a format is exempt: naming it is the whole job of
		// that row (D-112).
		for _, u := range unsupported {
			if declinesFormat(q.RowID) {
				break
			}
			if strings.Contains(lower, u.phrase) {
				add(q, "offers_an_unsupported_format", fmt.Sprintf("the question names %s", u.display))
				break
			}
		}
		// The question states a fact about the game.
		for _, c := range factClaims {
			if strings.Contains(lower, c) {
				add(q, "states_a_fact", fmt.Sprintf("the question claims %q", c))
				break
			}
		}
		// The question offers an answer the rules forbid.
		for _, o := range illegalOffers {
			if strings.Contains(lower, o) {
				add(q, "offers_an_illegal_answer",
					fmt.Sprintf("the question says %q, and a commander's color identity is the deck's", o))
				break
			}
		}
		// The question names one card's color identity.
		if m := possessiveIdentity.FindString(q.Text); m != "" {
			add(q, "names_a_card_color_identity",
				fmt.Sprintf("the question says %q, and no commander is settled when it goes out", strings.TrimSpace(m)))
		}
		if name, ok := stutteredName(q.Text); ok {
			add(q, "stuttered_card_name", fmt.Sprintf("the question repeats %q", name))
		}
		if n := strings.Count(q.Text, "?"); n > 1 {
			add(q, "two_questions_in_one", fmt.Sprintf("the question holds %d question marks", n))
		}
		if strings.ContainsAny(q.Text, "{}") {
			add(q, "holds_a_placeholder", "a brace reached the user")
		}
	}
	return out
}

// LintCatalog checks the rows themselves. It runs offline, in CI, and it
// needs no gate run.
func LintCatalog(c *Catalog) []Finding {
	var out []Finding
	for _, r := range c.Rows {
		for _, text := range []string{r.Text, r.Fallback} {
			if strings.TrimSpace(text) == "" {
				continue
			}
			lower := strings.ToLower(text)
			for _, w := range tableWords {
				// A row whose every trigger word names a table or an event
				// may name one back. The store-format row fires only after
				// the user writes FNM, an LGS, a store, an RCQ, a
				// tournament, or an event.
				if strings.Contains(lower, w) && !triggersProveATable(r) {
					out = append(out, Finding{Rule: "presumes_a_table", RowID: r.ID, Text: text,
						Detail: fmt.Sprintf("the row says %q, and a deck can be a gift (D-109)", w)})
					break
				}
			}
			for _, cl := range factClaims {
				if strings.Contains(lower, cl) {
					out = append(out, Finding{Rule: "states_a_fact", RowID: r.ID, Text: text,
						Detail: fmt.Sprintf("the row claims %q (D-108)", cl)})
					break
				}
			}
			if strings.Count(text, "?") > 1 {
				out = append(out, Finding{Rule: "two_questions_in_one", RowID: r.ID, Text: text,
					Detail: "the row holds more than one question mark"})
			}
		}
	}
	return out
}

// triggersProveATable reports whether every trigger word of a row is
// itself evidence that the user named a table or an event.
func triggersProveATable(r Row) bool {
	if len(r.When.Words) == 0 {
		return false
	}
	for _, w := range r.When.Words {
		if !anyPlain(strings.ToLower(w), tableEvidence) {
			return false
		}
	}
	return true
}

// priorWords joins every message the user sent up to and including the
// turn that carried the question.
func priorWords(messages []string, turn int) string {
	if turn <= 0 {
		return ""
	}
	if turn > len(messages) {
		turn = len(messages)
	}
	return strings.Join(messages[:turn], " ")
}

// anyPlain is a plain substring test. It reads a user's own words for
// evidence, so a negation does not matter: "we have no table" still shows
// that the user raised the subject.
func anyPlain(text string, words []string) bool {
	for _, w := range words {
		if strings.Contains(text, w) {
			return true
		}
	}
	return false
}

// declinesFormat reports whether a row exists to decline an unsupported
// format. Such a row must name that format, so both linter rules exempt
// it (D-112). format_unsupported_open is the variant that offers no
// substitute, which Historic and Timeless use (D-146).
func declinesFormat(rowID string) bool {
	return rowID == "format_unsupported" || rowID == "format_unsupported_open"
}
