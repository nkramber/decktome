package candidates

import (
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// typalLandCards are the shapes of F-148. Three lands make mana for a
// creature type the player chooses, one land reads the commander's types,
// one land names the type, and one land names no type. The last two cards
// are nonlands: one carries the typal tag, and one is a real dinosaur.
func typalLandCards() []tc {
	return []tc{
		{id: "cavern", name: "Cavern of Souls", typeLine: "Land",
			text:     "As this land enters, choose a creature type. {T}: Add {C}. {T}: Add one mana of any color. Spend this mana only to cast a creature spell of the chosen type, and that spell can't be countered.",
			produced: []mtgv1.Color{W, G}, rank: 40, tags: []string{"typal-choose"}},
		{id: "courtyard", name: "Secluded Courtyard", typeLine: "Land",
			text:     "As this land enters, choose a creature type. {T}: Add {C}. {T}: Add one mana of any color. Spend this mana only to cast a creature spell of the chosen type.",
			produced: []mtgv1.Color{W, G}, rank: 500, tags: []string{"typal-choose"}},
		{id: "city", name: "Three Tree City", typeLine: "Legendary Land",
			text:     "As Three Tree City enters, choose a creature type. {T}: Add {C}. {2}, {T}: Choose a color. Add an amount of mana of that color equal to the number of creatures you control of the chosen type.",
			produced: []mtgv1.Color{W, G}, rank: 600, tags: []string{"typal-choose"}},
		{id: "path", name: "Path of Ancestry", typeLine: "Land",
			text:     "This land enters tapped. {T}: Add one mana of any color in your commander's color identity. When that mana is spent to cast a creature spell that shares a creature type with your commander, scry 1.",
			produced: []mtgv1.Color{W, G}, rank: 300},
		{id: "hive", name: "Sliver Hive", typeLine: "Land",
			text:     "{T}: Add {C}. {T}: Add one mana of any color. Spend this mana only to cast a Sliver spell. {5}, {T}: Create a 1/1 colorless Sliver creature token.",
			produced: []mtgv1.Color{W, G}, rank: 700},
		{id: "tower", name: "Command Tower", typeLine: "Land",
			text:     "{T}: Add one mana of any color in your commander's color identity.",
			produced: []mtgv1.Color{W, G}, rank: 2, tags: []string{"utility-land"}},
		{id: "automaton", name: "Adaptive Automaton", typeLine: "Artifact Creature — Shapeshifter",
			text: "As this creature enters, choose a creature type. This creature is the chosen type in addition to its other types. Other creatures you control of the chosen type get +1/+1.",
			mv:   3, rank: 800, tags: []string{"typal-choose"}, subtypes: []string{"Shapeshifter"}},
		{id: "sliver", name: "Sidewinder Sliver", typeLine: "Creature — Sliver",
			text: "All Sliver creatures have flanking.", identity: []mtgv1.Color{W}, mv: 1, rank: 900,
			tags: []string{"typal-sliver"}, subtypes: []string{"Sliver"}},
	}
}

// signalOf reports whether the card carries one named signal of the match.
func signalOf(m ThemeMatch, c *mtgv1.Card, want string) bool {
	_, signals := m.score(c)
	return slices.Contains(signals, want)
}

// TestTypalLandSignalReachesTheTypeLands is D-759, D-760, and F-148. A
// typal shortlist lost every land that makes mana for its type, because
// no signal of the type row reads such a land. The type row carries the
// typal land signals now: the tag, the needle of Path of Ancestry, and
// its own subtype word.
func TestTypalLandSignalReachesTheTypeLands(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalLandCards())
	m := b.themes.match("slivers", idx.Tags())
	if len(m.LandSlugs) == 0 || len(m.LandText) == 0 {
		t.Fatalf("the slivers row carries no typal land signal: slugs %v text %v", m.LandSlugs, m.LandText)
	}
	if !slices.Contains(m.LandText, "sliver") {
		t.Errorf("land needles %v hold no subtype word", m.LandText)
	}
	for _, name := range []string{"Cavern of Souls", "Secluded Courtyard", "Three Tree City", "Path of Ancestry", "Sliver Hive"} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("the fixture has no %q", name)
		}
		if score, signals := m.score(c); score <= 0 {
			t.Errorf("%s scores %.2f with signals %v, want a typal land signal", name, score, signals)
		}
	}
	tower, _ := idx.ByName("Command Tower")
	if score, signals := m.score(tower); score != 0 {
		t.Errorf("Command Tower makes mana for every type and names none: score %.2f, signals %v", score, signals)
	}
}

// TestTypalLandSignalCountsOnALandAlone is D-760. The tag typal-choose
// holds 88 nonlands of the snapshot of 2026-09-04, such as Adaptive
// Automaton. The owner kept this pull request to the lands of F-148, so
// the signal reads the card type first.
func TestTypalLandSignalCountsOnALandAlone(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalLandCards())
	m := b.themes.match("slivers", idx.Tags())
	automaton, ok := idx.ByName("Adaptive Automaton")
	if !ok {
		t.Fatal("the fixture has no Adaptive Automaton")
	}
	if signalOf(m, automaton, "land-tag:typal-choose") {
		t.Error("Adaptive Automaton is no land, and it read the typal land tag")
	}
	// The row's own signals still read a real creature of the type.
	sliver, _ := idx.ByName("Sidewinder Sliver")
	if score, _ := m.score(sliver); score <= 0 {
		t.Errorf("Sidewinder Sliver scores %.2f, want a subtype signal", score)
	}
}

// TestNoTypeRowCarriesNoTypalLandSignal is D-759. A row with no subtype
// names no creature type, so a land that makes mana for a chosen type
// rewards it no more than any other deck.
func TestNoTypeRowCarriesNoTypalLandSignal(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalLandCards())
	m := b.themes.match("lifegain", idx.Tags())
	if len(m.LandSlugs) != 0 || len(m.LandText) != 0 {
		t.Errorf("the lifegain row carries typal land signals: slugs %v text %v", m.LandSlugs, m.LandText)
	}
	cavern, _ := idx.ByName("Cavern of Souls")
	if score, signals := m.score(cavern); score != 0 {
		t.Errorf("Cavern of Souls scores %.2f on lifegain with signals %v, want 0", score, signals)
	}
}

// TestTypalLandSignalCountsOnce is D-759. The tag and the needle state
// one fact about one card, so a land that holds both weighs as one payoff
// text and no more.
func TestTypalLandSignalCountsOnce(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalLandCards())
	m := b.themes.match("slivers", idx.Tags())
	// A land with the tag, the commander needle, and the subtype word.
	both := &mtgv1.Card{
		OracleId: "both", Name: "Every Signal", CardTypes: []string{"Land"},
		OracleText: "choose a creature type. Spend this mana only to cast a Sliver spell that shares a creature type with your commander.",
	}
	m.tagged["typal-choose"] = map[string]bool{"both": true}
	score, signals := m.score(both)
	if want := weightTypalLand / scoreCap; score != want {
		t.Errorf("a land with every typal land signal scores %.3f, want %.3f (signals %v)", score, want, signals)
	}
}

// TestGenericRuleCarriesTheTypalLandSignal is D-731 and D-761. A singular
// creature type keeps the generic rule and reads no type row, and the
// question gate records "dinosaur" and not "dinosaurs". A typal slug that
// the card database holds proves the word names a creature type.
func TestGenericRuleCarriesTheTypalLandSignal(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalLandCards())
	// "sliver" names no row of its own (D-731), and typal-sliver exists.
	if name, ok := b.themes.rowOf("sliver"); ok {
		t.Fatalf("rowOf(%q) = %q, and D-731 wants the generic rule", "sliver", name)
	}
	m := b.themes.match("sliver", idx.Tags())
	if len(m.LandSlugs) == 0 || len(m.LandText) == 0 {
		t.Fatalf("the generic rule carries no typal land signal: slugs %v text %v", m.LandSlugs, m.LandText)
	}
	for _, name := range []string{"Cavern of Souls", "Path of Ancestry", "Sliver Hive"} {
		c, _ := idx.ByName(name)
		if score, signals := m.score(c); score <= 0 {
			t.Errorf("%s scores %.2f on the singular theme with signals %v", name, score, signals)
		}
	}
	// A word with no typal slug names no creature type, so it takes none.
	if n := b.themes.match("zzzz", idx.Tags()); len(n.LandSlugs) != 0 || len(n.LandText) != 0 {
		t.Errorf("a word that names no creature type carries land signals: slugs %v text %v", n.LandSlugs, n.LandText)
	}
}

// f148Lands are the five lands that left the dinosaur typal shortlist of
// the deck gate on 2026-09-15. Four make mana for a creature type the
// player chooses, and Path of Ancestry reads the commander's types.
var f148Lands = []string{
	"Cavern of Souls", "Secluded Courtyard", "Unclaimed Territory",
	"Path of Ancestry", "Three Tree City",
}

// TestTypalLandsReachATypalShortlist is F-148 on the snapshot. Prompt 4
// of the deck gate is a dinosaur Commander deck at bracket 2, led by
// Gishath, Sun's Avatar. Before D-759 the five lands ranked 19, 25, and
// 31 in the mana order, and the land cap kept 40 lands of rank 13 and
// under. The test needs a snapshot, and `make themes-check` sets it.
func TestTypalLandsReachATypalShortlist(t *testing.T) {
	idx := snapshotIndex(t)
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	cmdr, ok := idx.ByName("Gishath, Sun's Avatar")
	if !ok {
		t.Fatal("the snapshot has no Gishath, Sun's Avatar")
	}
	// The gate prompt writes the plural, and the question gate records the
	// singular. A singular type word keeps the generic rule (D-731, D-761).
	for _, theme := range []string{"dinosaurs", "dinosaur"} {
		list, err := b.Build(idx, Request{
			Format:             mtgv1.FormatId_FORMAT_ID_COMMANDER,
			Colors:             cmdr.GetColorIdentity(),
			Theme:              theme,
			CommanderOracleIDs: []string{cmdr.GetOracleId()},
			PoolRule:           mtgv1.PoolRule_POOL_RULE_ANY_CARD,
			Bracket:            2,
		})
		if err != nil {
			t.Fatalf("build %q: %v", theme, err)
		}
		in := map[string]Candidate{}
		for _, c := range list.Candidates {
			in[c.Card.GetName()] = c
		}
		for _, name := range f148Lands {
			c, ok := in[name]
			if !ok {
				t.Errorf("the %q shortlist lost %q (F-148)", theme, name)
				continue
			}
			if !c.Themed {
				t.Errorf("%q of the %q shortlist reads no theme signal, so the land cap can drop it", name, theme)
			}
		}
	}
}

// TestLandNeedleMatchesAWholeWord answers the Gitar finding of #190. A
// creature type is a part of a common word of land text: "bat" sits in
// "battlefield", "orc" in "sorcery", and "mount" in "mountain". Those
// three read 218 of the 1,194 Commander-legal lands of the snapshot of
// 2026-09-04, and each one would take a place of the theme half.
func TestLandNeedleMatchesAWholeWord(t *testing.T) {
	no := []struct{ text, needle string }{
		{"put it onto the battlefield tapped", "bat"},
		{"activate only as a sorcery", "orc"},
		{"search your library for a mountain card", "mount"},
		{"whenever a creature you control dies", "rat"},
		{"a 1/1 shapeshifter creature token with changeling", "angel"},
	}
	for _, c := range no {
		if containsWord(c.text, c.needle) {
			t.Errorf("containsWord(%q, %q) = true, want false", c.text, c.needle)
		}
	}
	yes := []struct{ text, needle string }{
		{"spend this mana only to cast a sliver spell", "sliver"},
		{"create a 2/2 black zombie creature token", "zombie"},
		{"zombies you control get +1/+0", "zombie"},
		{"this land becomes a 2/5 red and green dinosaur creature", "dinosaur"},
		{"a creature spell that shares a creature type with your commander", "shares a creature type with your commander"},
		{"(it's every creature type.)", "every creature type"},
		{"bat tokens you control", "bat"},
	}
	for _, c := range yes {
		if !containsWord(c.text, c.needle) {
			t.Errorf("containsWord(%q, %q) = false, want true", c.text, c.needle)
		}
	}
}

// TestChangelingLandReadsEveryTypalTheme answers the Gitar finding of
// #190. A changeling permanent is every creature type, so a changeling
// land rewards every typal deck. The word-boundary rule drops the
// accidental hit of "angel" inside "changeling", and the needle "every
// creature type" keeps the land on purpose.
func TestChangelingLandReadsEveryTypalTheme(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	cards := append(typalLandCards(), tc{
		id: "countryside", name: "Abundant Countryside", typeLine: "Land",
		text:     "{T}: Add {C}. {T}: Add one mana of any color. Spend this mana only to cast a creature spell. {6}, {T}: Create a 1/1 colorless Shapeshifter creature token with changeling. (It's every creature type.)",
		produced: []mtgv1.Color{W, G}, rank: 1000,
	})
	idx := fixture(t, cards)
	countryside, ok := idx.ByName("Abundant Countryside")
	if !ok {
		t.Fatal("the fixture has no Abundant Countryside")
	}
	for _, theme := range []string{"slivers", "dinosaurs", "sliver"} {
		m := b.themes.match(theme, idx.Tags())
		if score, signals := m.score(countryside); score <= 0 {
			t.Errorf("a changeling land scores %.2f on %q with signals %v", score, theme, signals)
		}
	}
}
