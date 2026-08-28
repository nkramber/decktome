package generate

import (
	"fmt"
	"regexp"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// F-26 says the model invents claims about the game, and that no
// deterministic check caught them. Gate run 14 of 2026-08-26 told one
// user "Grist, the Hunger Tide can not lead a deck", which the rules
// contradict, and it asked a question the rules allow no answer to. The
// eval lane found both. The linter found neither.
//
// The deck summary is prose, so it has more room for this than a
// question ever had. The rules below read the SHAPE of a rules claim,
// and they can not read its truth.
//
// CAUTION: this is a net, not a gate. A summary that states a false rule
// in words these patterns do not hold still reaches the user. The judge
// lane of PR-7B is the real check for F-26, and it costs money to run.

// claimShapes are the phrasings a rules claim takes. Each one is a claim
// about what the game allows, and the summary has no business making one.
var claimShapes = []*regexp.Regexp{
	// "can not lead a deck", "cannot be your commander", "can be your commander"
	regexp.MustCompile(`(?i)\b(can|could|may)( ?n[o']?t)?\s+(be\s+(your|the|a)\s+commander|lead\s+(a|the|your)\s+deck)`),
	// "is banned", "are banned in", "is restricted"
	regexp.MustCompile(`(?i)\b(is|are|was|were)\s+(banned|restricted|illegal|unbanned)\b`),
	// "is legal in", "is not legal in"
	regexp.MustCompile(`(?i)\b(is|are)( ?n[o']?t)?\s+legal\b`),
	// "the rules allow", "the format forbids", "the ban list"
	regexp.MustCompile(`(?i)\b(the\s+rules?|the\s+format|the\s+ban\s+list)\s+\w+`),
	// "you may play", "you cannot run", "players may not"
	regexp.MustCompile(`(?i)\b(you|players?)\s+(may|can|must)( ?n[o']?t)?\s+(play|run|use|include)\b`),
	// D-144's shape: a card's color identity, stated as a rule.
	regexp.MustCompile(`(?i)['\x{2019}]s\s+color\s+identity`),
	// "counts as a Game Changer", "counts toward the bracket"
	regexp.MustCompile(`(?i)\bcounts?\s+(as|toward|towards)\s+`),
	// "singleton format", "the copy limit is"
	regexp.MustCompile(`(?i)\b(singleton|the\s+copy\s+limit|the\s+deck\s+size)\b`),
}

// CodeSummaryRulesClaim is the finding a summary gets when it states a
// rule of the game.
const CodeSummaryRulesClaim = "summary_rules_claim"

// LintSummary reports the rules claims a summary appears to make. It
// returns the matched text of each one, in the order they appear.
func LintSummary(text string) []string {
	var out []string
	seen := map[string]bool{}
	for _, re := range claimShapes {
		for _, m := range re.FindAllString(text, -1) {
			m = strings.TrimSpace(m)
			if m == "" || seen[strings.ToLower(m)] {
				continue
			}
			seen[strings.ToLower(m)] = true
			out = append(out, m)
		}
	}
	return out
}

// lintSummaryInto adds one warning per rules claim in the deck summary.
// It is a warning and never a block: a shape is not a falsehood, and a
// block would refuse a correct deck over a phrase. The user still sees
// the finding, and the judge lane decides the truth.
func lintSummaryInto(deck *mtgv1.Deck) {
	claims := LintSummary(deck.GetSummary())
	if len(claims) == 0 {
		return
	}
	for _, c := range claims {
		addFinding(deck, CodeSummaryRulesClaim, mtgv1.Severity_SEVERITY_WARN,
			fmt.Sprintf("the summary states a rule of the game, %q. The engine reports the rules, and the summary must not.", c))
	}
}
