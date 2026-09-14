package cardname

import "testing"

// TestExactKeepsTheAccent: the exact key folds the case, the outer
// spaces, and a curly quote, and nothing else.
func TestExactKeepsTheAccent(t *testing.T) {
	for _, tc := range []struct{ in, want string }{
		{"  Commander’s Sphere ", "commander's sphere"},
		{"Gríma Wormtongue", "gríma wormtongue"},
	} {
		if got := Exact(tc.in); got != tc.want {
			t.Errorf("Exact(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// TestFoldReadsWhatAKeyboardTypes is D-716. Each row pairs a card name
// of the snapshot with the name a reader types without the special mark.
func TestFoldReadsWhatAKeyboardTypes(t *testing.T) {
	for _, tc := range []struct{ name, typed string }{
		{"Gríma, Saruman's Footman", "Grima, Saruman’s Footman"},
		{"Barad-dûr", "  BARAD-DUR "},
		{"Altaïr Ibn-La'Ahad", "Altair Ibn-La'Ahad"},
		{"Araña, Heart of the Spider", "Arana, Heart of the Spider"},
		{"Bespoke Bō", "Bespoke Bo"},
		{"Ratonhnhaké꞉ton", "Ratonhnhake:ton"},
		{"Human—Time Lord Meta-Crisis", "Human-Time Lord Meta-Crisis"},
		{"The Ultimate Nightmare of Wizards of the Coast® Customer Service", "The Ultimate Nightmare of Wizards of the Coast Customer Service"},
		{"Aeon Chronicler", "Æon Chronicler"},
		{"Gríma Wormtongue", "Ｇｒｉｍａ Wormtongue"},
	} {
		if Fold(tc.name) != Fold(tc.typed) {
			t.Errorf("Fold(%q) = %q, and Fold(%q) = %q, want one key", tc.name, Fold(tc.name), tc.typed, Fold(tc.typed))
		}
	}
}

// TestFoldKeepsTwoCardsApart: the fold reads marks, not words, so a near
// name stays a different card (F-13).
func TestFoldKeepsTwoCardsApart(t *testing.T) {
	if Fold("Ajani's Pridemate") == Fold("Ajani's Welcome") {
		t.Error("two different cards share a folded key")
	}
	if Fold("Fire // Ice") != "fire // ice" {
		t.Errorf("Fold(Fire // Ice) = %q, want the name with its slashes", Fold("Fire // Ice"))
	}
}
