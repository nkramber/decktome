package generate

import "testing"

// TestLintSummaryCatchesTheTwoClaimsThatReachedAUser is F-26. Gate run 14
// of 2026-08-26 told one user that Grist, the Hunger Tide can not lead a
// deck, which the rules contradict, and it named a card's color identity
// as a rule. Both passed the gate and the linter.
func TestLintSummaryCatchesTheTwoClaimsThatReachedAUser(t *testing.T) {
	claims := []string{
		"Grist, the Hunger Tide can not lead a deck, so it sits in the 99.",
		"Every card fits Karlov's color identity.",
		"Sol Ring is banned in this format.",
		"Mana Crypt is not legal here.",
		"The rules allow two commanders when both have partner.",
		"You may not play more than one copy.",
		"Rhystic Study counts as a Game Changer.",
		"Commander is a singleton format.",
	}
	for _, c := range claims {
		if got := LintSummary(c); len(got) == 0 {
			t.Errorf("no claim found in %q", c)
		}
	}
}

// TestLintSummaryPassesAPlainPlan covers the false positive. A summary
// that says what the deck does states no rule, and it must go out clean.
func TestLintSummaryPassesAPlainPlan(t *testing.T) {
	clean := []string{
		"This deck gains life and turns it into damage. Karlov grows on every gain, and the drain effects close the game. It gives up speed for resilience.",
		"A tempo deck. Cheap threats land early, and the counterspells protect them. The sideboard answers graveyard decks.",
		"The deck ramps into large creatures, then wins with one big attack. It is slow against direct damage.",
		"An artifact deck built around cheap permanents and the payoffs that count them.",
	}
	for _, s := range clean {
		if got := LintSummary(s); len(got) > 0 {
			t.Errorf("a plain plan was flagged: %q -> %v", s, got)
		}
	}
}

// TestPluralReadsCorrectly covers the finding text, which the user reads.
func TestPluralReadsCorrectly(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want string
	}{
		{1, "1 basic land"}, {2, "2 basic lands"}, {0, "0 basic lands"},
	} {
		if got := plural(tc.n, "basic land"); got != tc.want {
			t.Errorf("plural(%d) = %q, want %q", tc.n, got, tc.want)
		}
	}
}
