package candidates

import (
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// typalCards are the shapes of F-144 and F-148. Three lands make mana for
// a creature type the player chooses, one land reads the commander's
// types, one land names the type, and one land names no type. The other
// cards are nonlands. Adaptive Automaton chooses a type, Coat of Arms
// rewards a shared type, Kilnmouth Dragon and Knowledge Exploitation
// reward one named type, and Sidewinder Sliver is a real Sliver.
func typalCards() []tc {
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
		{id: "coat", name: "Coat of Arms", typeLine: "Artifact",
			text: "Each creature gets +1/+1 for each other creature on the battlefield that shares at least one creature type with it.",
			mv:   5, rank: 850, tags: []string{"typal-share"}},
		{id: "kilnmouth", name: "Kilnmouth Dragon", typeLine: "Creature — Dragon",
			text:     "Amplify 3 (As this creature enters, put three +1/+1 counters on it for each Dragon card you reveal in your hand.) Flying {T}: This creature deals damage equal to the number of +1/+1 counters on it to any target.",
			identity: []mtgv1.Color{R}, mv: 7, rank: 950, tags: []string{"typal-share"}, subtypes: []string{"Dragon"}},
		{id: "exploitation", name: "Knowledge Exploitation", typeLine: "Kindred Sorcery — Rogue",
			text:     "Prowl {3}{U} (You may cast this for its prowl cost if you dealt combat damage to a player this turn with a Rogue.) Search target opponent's library for an instant or sorcery card. You may cast that card without paying its mana cost. Then that player shuffles.",
			identity: []mtgv1.Color{U}, mv: 7, rank: 960, tags: []string{"typal-share"}, subtypes: []string{"Rogue"}},
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
	idx := fixture(t, typalCards())
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

// TestTypalLandSignalCountsOnALandAlone is D-760. The land signals and
// the card signals read two different card types, so a nonland reads the
// card tag and no land tag.
func TestTypalLandSignalCountsOnALandAlone(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalCards())
	m := b.themes.match("slivers", idx.Tags())
	automaton, ok := idx.ByName("Adaptive Automaton")
	if !ok {
		t.Fatal("the fixture has no Adaptive Automaton")
	}
	if signalOf(m, automaton, "land-tag:typal-choose") {
		t.Error("Adaptive Automaton is no land, and it read the typal land tag")
	}
	cavern, _ := idx.ByName("Cavern of Souls")
	if signalOf(m, cavern, "card-tag:typal-choose") {
		t.Error("Cavern of Souls is a land, and it read the typal card tag")
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
	idx := fixture(t, typalCards())
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
	idx := fixture(t, typalCards())
	m := b.themes.match("slivers", idx.Tags())
	// A land with the tag, the commander needle, and the subtype word.
	both := &mtgv1.Card{
		OracleId: "both", Name: "Every Signal", CardTypes: []string{"Land"},
		OracleText: "choose a creature type. Spend this mana only to cast a Sliver spell that shares a creature type with your commander.",
	}
	m.tagged["typal-choose"] = map[string]bool{"both": true}
	score, signals := m.score(both)
	if want := weightTypal / scoreCap; score != want {
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
	idx := fixture(t, typalCards())
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
	cards := append(typalCards(), tc{
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

// TestTypalCardSignalReachesTheGenericPayoffs is the nonland half of
// F-144. The tag typal-choose holds 88 nonlands of the snapshot of
// 2026-09-04, and no type row reads one of them. Door of Destinies names
// no creature type, so the subtype and the payoff needles of the row miss
// it.
func TestTypalCardSignalReachesTheGenericPayoffs(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalCards())
	m := b.themes.match("slivers", idx.Tags())
	automaton, _ := idx.ByName("Adaptive Automaton")
	score, signals := m.score(automaton)
	if !slices.Contains(signals, "card-tag:typal-choose") {
		t.Errorf("Adaptive Automaton reads signals %v, want the typal card tag", signals)
	}
	if want := weightTypal / scoreCap; score != want {
		t.Errorf("Adaptive Automaton scores %.3f, want %.3f", score, want)
	}
}

// TestNoncreatureTypalSignalSkipsACreature is F-144. The tag typal-share
// holds 70 Commander-legal nonlands of the snapshot of 2026-09-04. Thirty
// of them are no creature and reward any creature type, such as Coat of
// Arms. The other 40 are creatures of one named type, such as Kilnmouth
// Dragon, and they reward that type alone.
func TestNoncreatureTypalSignalSkipsACreature(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalCards())
	m := b.themes.match("slivers", idx.Tags())
	coat, ok := idx.ByName("Coat of Arms")
	if !ok {
		t.Fatal("the fixture has no Coat of Arms")
	}
	if !signalOf(m, coat, "noncreature-tag:typal-share") {
		t.Error("Coat of Arms is no creature, and it read no typal share tag")
	}
	for _, name := range []string{"Kilnmouth Dragon", "Knowledge Exploitation"} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("the fixture has no %s", name)
		}
		if score, signals := m.score(c); score != 0 {
			t.Errorf("%s scores %.3f on slivers with signals %v, want 0", name, score, signals)
		}
	}
}

// TestTypalCardSignalCountsOnce is F-144. The two tags state one fact
// about one card, so a card that holds both weighs as one payoff text.
func TestTypalCardSignalCountsOnce(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalCards())
	m := b.themes.match("slivers", idx.Tags())
	both := &mtgv1.Card{
		OracleId: "both", Name: "Every Card Signal", CardTypes: []string{"Artifact"},
		OracleText: "As this artifact enters, choose a creature type.",
	}
	m.tagged["typal-choose"] = map[string]bool{"both": true}
	m.tagged["typal-share"] = map[string]bool{"both": true}
	score, signals := m.score(both)
	if want := weightTypal / scoreCap; score != want {
		t.Errorf("a card with both typal card tags scores %.3f, want %.3f (signals %v)", score, want, signals)
	}
}

// TestNoTypeRowCarriesNoTypalCardSignal is D-759 and F-144. A row with no
// subtype names no creature type, so a card that rewards a chosen type
// rewards it no more than any other deck.
func TestNoTypeRowCarriesNoTypalCardSignal(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalCards())
	m := b.themes.match("lifegain", idx.Tags())
	if len(m.CardSlugs) != 0 || len(m.NoncreatureSlugs) != 0 {
		t.Errorf("the lifegain row carries typal card signals: %v %v", m.CardSlugs, m.NoncreatureSlugs)
	}
	for _, name := range []string{"Adaptive Automaton", "Coat of Arms"} {
		c, _ := idx.ByName(name)
		if score, signals := m.score(c); score != 0 {
			t.Errorf("%s scores %.3f on lifegain with signals %v, want 0", name, score, signals)
		}
	}
}

// TestGenericRuleCarriesTheTypalCardSignal is D-761 and F-144. A singular
// creature type keeps the generic rule and reads no type row, and the
// question gate records "dinosaur" and not "dinosaurs". The card signals
// follow the land signals onto that rule.
func TestGenericRuleCarriesTheTypalCardSignal(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typalCards())
	for _, theme := range []string{"slivers", "sliver"} {
		m := b.themes.match(theme, idx.Tags())
		automaton, _ := idx.ByName("Adaptive Automaton")
		if !signalOf(m, automaton, "card-tag:typal-choose") {
			t.Errorf("the %q theme reads no typal card tag on Adaptive Automaton", theme)
		}
		coat, _ := idx.ByName("Coat of Arms")
		if !signalOf(m, coat, "noncreature-tag:typal-share") {
			t.Errorf("the %q theme reads no typal share tag on Coat of Arms", theme)
		}
	}
}

// f144Cards are generic typal payoffs of the snapshot of 2026-09-04. No
// type row reads one of them, because each one names no creature type.
// The first five carry typal-choose, and the last three carry
// typal-share and are no creature.
var f144Cards = []string{
	"Door of Destinies", "Herald's Horn", "Icon of Ancestry",
	"Urza's Incubator", "Vanquisher's Banner",
	"Coat of Arms", "Shared Animosity", "Descendants' Path",
}

// TestTypalCardsReachATypalShortlist is F-144 on the snapshot. Prompt 4
// of the deck gate is a dinosaur Commander deck at bracket 2, led by
// Gishath, Sun's Avatar. The test needs a snapshot, and `make
// themes-check` sets it.
func TestTypalCardsReachATypalShortlist(t *testing.T) {
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
		for _, name := range f144Cards {
			c, ok := in[name]
			if !ok {
				t.Errorf("the %q shortlist lost %q (F-144)", theme, name)
				continue
			}
			if !c.Themed {
				t.Errorf("%q of the %q shortlist reads no theme signal", name, theme)
			}
		}
	}
}

// TestTypalCardSignalReadsNoTypeOfItsOwn is F-144 on the snapshot. A
// creature of one named type rewards that type alone, so a rabbits
// shortlist must read no signal on a dragon of the tag typal-share.
func TestTypalCardSignalReadsNoTypeOfItsOwn(t *testing.T) {
	idx := snapshotIndex(t)
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	m := b.themes.match("rabbits", idx.Tags())
	for _, name := range []string{"Kilnmouth Dragon", "Hunting Velociraptor", "Knowledge Exploitation"} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("the snapshot has no %s", name)
		}
		if score, signals := m.score(c); score != 0 {
			t.Errorf("%s scores %.3f on rabbits with signals %v, want 0", name, score, signals)
		}
	}
	// The generic payoffs of the same tag still read the theme.
	for _, name := range []string{"Coat of Arms", "Shared Animosity"} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("the snapshot has no %s", name)
		}
		if !signalOf(m, c, "noncreature-tag:typal-share") {
			t.Errorf("%s reads no typal share tag on rabbits", name)
		}
	}
}
