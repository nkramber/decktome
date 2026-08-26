package questions

import "testing"

// TestCatalogIsClean runs the linter over the rows themselves. It needs no
// gate run and no model call, so CI holds the line (D-115).
func TestCatalogIsClean(t *testing.T) {
	for _, f := range LintCatalog(load(t)) {
		t.Errorf("catalog row %q: %s (%q)", f.RowID, f.Detail, f.Text)
	}
}

// TestLintFindsTheKnownDefects replays the questions that the owner's
// M-5 scoring of 2026-08-25 marked wrong. Each case is a real line from a
// gate document.
func TestLintFindsTheKnownDefects(t *testing.T) {
	cases := []struct {
		name     string
		messages []string
		q        LintQuestion
		want     string
	}{
		{
			name:     "the user named the format in the first message",
			messages: []string{"A land destruction Commander deck."},
			q: LintQuestion{Turn: 1, RowID: "format", Slot: "format",
				Text: "What format would you like: Commander, Standard, Modern, Pioneer, or Pauper?"},
			want: "format_already_named",
		},
		{
			name:     "a deck as a gift has no table",
			messages: []string{"I want to build a deck as a gift for my brother. He likes zombies.", "Commander, black. He is new to the game."},
			q: LintQuestion{Turn: 2, RowID: "power_commander", Slot: "power",
				Text: "Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power."},
			want: "presumes_a_table",
		},
		{
			name:     "the locked row repeated the card name",
			messages: []string{"Build around Grist, the Hunger Tide, but not as my commander."},
			q: LintQuestion{Turn: 1, RowID: "locked", Slot: "locked",
				Text: "Must the deck keep Grist, the Hunger Tide and Grist, or may I cut a card that does not fit the plan?"},
			want: "stuttered_card_name",
		},
		{
			name:     "the colors row stated a fact",
			messages: []string{"I want a 60-card deck, and anything goes at our table.", "Any card, no ban list. Call it Modern. A dragon deck."},
			q: LintQuestion{Turn: 2, RowID: "colors", Slot: "colors",
				Text: "Any color preference? A a dragon deck deck is strongest in blue and red."},
			want: "states_a_fact",
		},
		{
			name:     "the format row offered Brawl",
			messages: []string{"I want a Brawl deck for Arena."},
			q: LintQuestion{Turn: 1, RowID: "format", Slot: "format",
				Text: "Which format would you like for Arena: Brawl or something else?"},
			want: "offers_an_unsupported_format",
		},
		{
			name:     "a brace reached the user",
			messages: []string{"Build me a deck."},
			q:        LintQuestion{Turn: 1, RowID: "colors", Slot: "colors", Text: "Any color preference for your {theme} deck?"},
			want:     "holds_a_placeholder",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			found := false
			for _, f := range LintConversation(tc.messages, []LintQuestion{tc.q}) {
				if f.Rule == tc.want {
					found = true
				}
			}
			if !found {
				t.Errorf("the linter missed %q in %q", tc.want, tc.q.Text)
			}
		})
	}
}

// TestLintAllowsWhatTheUserRaised keeps the linter from firing on a
// question that repeats the user's own words.
func TestLintAllowsWhatTheUserRaised(t *testing.T) {
	cases := []struct {
		name     string
		messages []string
		q        LintQuestion
	}{
		{
			name:     "the user named a playgroup",
			messages: []string{"Make a mill deck for my playgroup, budget 100."},
			q: LintQuestion{Turn: 1, RowID: "power_commander", Slot: "power",
				Text: "Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power."},
		},
		{
			name:     "the user named no format",
			messages: []string{"Build me a lifegain deck. I have a collection."},
			q: LintQuestion{Turn: 1, RowID: "format", Slot: "format",
				Text: "Which format: Commander, Standard, Modern, or something else?"},
		},
		{
			name:     "the unsupported-format row may name the format",
			messages: []string{"I want a Brawl deck for Arena."},
			q: LintQuestion{Turn: 1, RowID: "format_unsupported", Slot: "format",
				Text: "I do not build Brawl. The nearest format I build is Commander. Shall I use that?"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// A clean question produces no finding at all. An earlier
			// version of this test allowed one through, and the live smoke
			// before gate 14 caught what it missed.
			for _, f := range LintConversation(tc.messages, []LintQuestion{tc.q}) {
				t.Errorf("the linter fired %q on a clean question: %s", f.Rule, f.Detail)
			}
		})
	}
}
