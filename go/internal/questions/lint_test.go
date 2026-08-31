package questions

import "testing"

// TestCatalogIsClean runs the linter over the rows themselves. It needs no
// model call, so CI holds the line (D-115).
func TestCatalogIsClean(t *testing.T) {
	for _, f := range LintCatalog(load(t)) {
		t.Errorf("catalog row %q: %s (%q)", f.RowID, f.Detail, f.Text)
	}
}

// TestLintFindsTheKnownDefects replays questions the M-5 scoring marked
// wrong (D-66, D-115). Each case is a real line from a gate document.
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
			name:     "the role row repeated the card name",
			messages: []string{"Build around Grist, the Hunger Tide."},
			q: LintQuestion{Turn: 1, RowID: "named_card_role", Slot: "commander",
				Text: "Do you want Grist, the Hunger Tide and Grist as your commander, or as one card in the 99?"},
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
			// A clean question produces no finding at all. A test that
			// allows one through misses a real defect.
			for _, f := range LintConversation(tc.messages, []LintQuestion{tc.q}) {
				t.Errorf("the linter fired %q on a clean question: %s", f.Rule, f.Detail)
			}
		})
	}
}

// TestLintCatchesAnIllegalOffer is D-144. "Do you want to use any colors
// beyond Grist's color identity?" offers an answer the rules forbid: in
// Commander the color identity of the commander is the color identity
// of the deck.
func TestLintCatchesAnIllegalOffer(t *testing.T) {
	bad := []string{
		"Do you want to use any colors beyond Grist's color identity?",
		"Should the deck use colors outside its color identity?",
		"Would you like additional colors beyond the commander's identity?",
	}
	for _, text := range bad {
		q := LintQuestion{Turn: 1, RowID: "colors", Slot: "colors", Text: text}
		var found bool
		for _, f := range LintConversation([]string{"Build around Grist, the Hunger Tide."}, []LintQuestion{q}) {
			if f.Rule == "offers_an_illegal_answer" {
				found = true
			}
		}
		if !found {
			t.Errorf("the linter missed an illegal offer: %q", text)
		}
	}
	// A normal color question is untouched.
	ok := LintQuestion{Turn: 1, RowID: "colors", Slot: "colors", Text: "Any color preference?"}
	for _, f := range LintConversation([]string{"Build me a deck."}, []LintQuestion{ok}) {
		t.Errorf("the linter fired on a clean question: %s", f.Detail)
	}
}

// An acronym the reader has never seen explains itself the first time
// (D-374). The catalog spells FNM out, and the rule catches a rewording
// that drops it.
func TestLintUnexplainedAcronym(t *testing.T) {
	rules := func(fs []Finding) []string {
		var out []string
		for _, f := range fs {
			out = append(out, f.Rule)
		}
		return out
	}
	msgs := []string{"a 60-card deck"}

	t.Run("a bare FNM is a finding", func(t *testing.T) {
		got := LintConversation(msgs, []LintQuestion{{Turn: 1, RowID: "power_sixty", Slot: "power", Text: "How strong: casual, FNM level, or tournament-meta?"}})
		if !contains(rules(got), "unexplained_acronym") {
			t.Errorf("findings = %v, want unexplained_acronym", rules(got))
		}
	})

	t.Run("the spelled-out form passes", func(t *testing.T) {
		got := LintConversation(msgs, []LintQuestion{{Turn: 1, RowID: "power_sixty", Slot: "power", Text: "How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?"}})
		if contains(rules(got), "unexplained_acronym") {
			t.Errorf("the spelled-out question was flagged: %v", rules(got))
		}
	})

	t.Run("a reader who wrote FNM first hears it back", func(t *testing.T) {
		got := LintConversation([]string{"a deck for FNM"}, []LintQuestion{{Turn: 1, RowID: "power_sixty", Slot: "power", Text: "How strong: casual, FNM level, or tournament-meta?"}})
		if contains(rules(got), "unexplained_acronym") {
			t.Errorf("the user used the acronym first, so the question may: %v", rules(got))
		}
	})
}
