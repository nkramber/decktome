package profile

import "testing"

func TestSourcesNeededReadsTheTableAndClamps(t *testing.T) {
	cases := []struct {
		size, generic, pips, want int
	}{
		{99, 0, 1, 19}, {99, 1, 1, 19}, {99, 2, 1, 18}, {99, 1, 2, 28}, {99, 0, 3, 36}, {99, 1, 4, 36},
		{60, 0, 1, 14}, {60, 1, 2, 18}, {60, 2, 2, 16}, {60, 0, 4, 24},
		// Past the table: more generic mana reads the last row, more pips
		// read four, and no pip needs nothing.
		{99, 9, 1, 14}, {99, 7, 2, 20}, {99, 6, 3, 26}, {99, 3, 4, 36}, {99, 0, 6, 39}, {99, 2, 0, 0},
		{99, -1, 1, 19},
	}
	for _, tc := range cases {
		if got := sourcesNeeded(tc.size, tc.generic, tc.pips); got != tc.want {
			t.Errorf("sourcesNeeded(%d, %d, %d) = %d, want %d", tc.size, tc.generic, tc.pips, got, tc.want)
		}
	}
}

func TestRequirementCoversMostOfTheDeck(t *testing.T) {
	if got := requirement(nil); got != 0 {
		t.Errorf("no needs give %d", got)
	}
	// Nine cheap needs and one greedy card: the requirement is the
	// cheap one, so one card does not set the bar.
	needs := []int{19, 19, 19, 19, 19, 19, 19, 19, 19, 36}
	if got := requirement(needs); got != 19 {
		t.Errorf("requirement %d, want 19", got)
	}
	// Three greedy cards of ten read as greedy.
	needs = []int{19, 19, 19, 19, 19, 19, 19, 36, 36, 36}
	if got := requirement(needs); got != 36 {
		t.Errorf("requirement %d, want 36", got)
	}
}
