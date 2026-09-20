package candidates

import (
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// typeNameCards are the shapes of the second half of F-144. Three cards
// name a creature type in their text and carry no type of their own: a
// token maker, a card that turns a permanent into the type, and a card
// that punishes the type. One card is a real Zombie, and one names no
// type at all.
func typeNameCards() []tc {
	return []tc{
		{id: "army", name: "Army of the Damned", typeLine: "Sorcery",
			text:     "Create thirteen tapped 2/2 black Zombie creature tokens. Flashback {7}{B}{B}{B}",
			identity: []mtgv1.Color{B}, mv: 7, rank: 100},
		{id: "mantle", name: "Nim Deathmantle", typeLine: "Artifact — Equipment",
			text: "Equipped creature gets +2/+2, has intimidate, and is a black Zombie in addition to its other colors and types.",
			mv:   2, rank: 200},
		{id: "slayer", name: "Undead Slayer", typeLine: "Creature — Human Cleric",
			text:     "{W}: Exile target attacking or blocking Zombie.",
			identity: []mtgv1.Color{W}, mv: 2, rank: 300, subtypes: []string{"Human", "Cleric"}},
		{id: "ghoul", name: "Butcher Ghoul", typeLine: "Creature — Zombie",
			text: "Undying", identity: []mtgv1.Color{B}, mv: 2, rank: 400,
			tags: []string{"typal-zombie"}, subtypes: []string{"Zombie"}},
		{id: "whisper", name: "Night's Whisper", typeLine: "Sorcery",
			text:     "You draw two cards and you lose 2 life.",
			identity: []mtgv1.Color{B}, mv: 2, rank: 500, tags: []string{"card-draw"}},
	}
}

// TestTypeRowReadsATypeNameInText is the second half of F-144 (D-768,
// D-769). A type row read a card that names its type with no signal, so
// a Zombie token maker reached no zombie shortlist. The row carries the
// type word as a text needle now.
func TestTypeRowReadsATypeNameInText(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typeNameCards())
	m := b.themes.match("zombies", idx.Tags())
	for _, name := range []string{"Army of the Damned", "Nim Deathmantle"} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("the fixture has no %q", name)
		}
		if !signalOf(m, c, "type-text:zombie") {
			_, signals := m.score(c)
			t.Errorf("%q names a Zombie and reads signals %v, want type-text:zombie", name, signals)
		}
	}
	whisper, _ := idx.ByName("Night's Whisper")
	if score, signals := m.score(whisper); score != 0 {
		t.Errorf("Night's Whisper names no type and scores %.2f with signals %v, want 0", score, signals)
	}
}

// TestTypeNameWeighsAsText is D-769. The needle names a type and proves
// no payoff, so it weighs as text and it never outranks a real creature
// of the type.
func TestTypeNameWeighsAsText(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typeNameCards())
	m := b.themes.match("zombies", idx.Tags())
	army, _ := idx.ByName("Army of the Damned")
	ghoul, _ := idx.ByName("Butcher Ghoul")
	armyScore, _ := m.score(army)
	ghoulScore, _ := m.score(ghoul)
	if want := weightText / scoreCap; armyScore != want {
		t.Errorf("Army of the Damned scores %.3f, want %.3f", armyScore, want)
	}
	if ghoulScore <= armyScore {
		t.Errorf("Butcher Ghoul is a Zombie and scores %.3f, want more than the %.3f of a card that names one",
			ghoulScore, armyScore)
	}
}

// TestTypeRowAndItsSingularReadTheSameText is D-731 and D-769. The
// generic rule gave a singular type word the type word as a text needle,
// and the plural read a type row that carried no such needle. So
// "zombie" and "zombies" read two different pools. Both read the needle
// now.
func TestTypeRowAndItsSingularReadTheSameText(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	rows := typeRows(t, b)
	// The generic rule reads a type word only when the card database
	// holds its typal tag (D-759, D-761). So the fixture carries one
	// tagged card of each type.
	list := typeNameCards()
	for name, word := range rows {
		list = append(list, tc{
			id: "typal-" + word, name: "Fixture " + name, typeLine: "Creature — " + title(word),
			text: "Other " + title(word) + "s you control get +1/+1.",
			mv:   3, rank: 1000, tags: []string{"typal-" + word}, subtypes: []string{title(word)},
		})
	}
	idx := fixture(t, list)
	for name, word := range rows {
		row := b.themes.match(name, idx.Tags())
		if !slices.Contains(row.TypeText, word) {
			t.Errorf("the %q row reads type needles %v, want %q", name, row.TypeText, word)
		}
		single := b.themes.match(word, idx.Tags())
		if !slices.Contains(single.TypeText, word) {
			t.Errorf("the generic word %q reads type needles %v, want %q", word, single.TypeText, word)
		}
		// A type word takes the word-boundary needle and no text needle,
		// because a text needle matches by substring (D-769).
		if slices.Contains(single.Text, word) {
			t.Errorf("the generic word %q reads the substring needle %q", word, word)
		}
	}
}

// TestTypeNeedleMatchesAWholeWord is D-769. A text needle matches by
// substring, and a short type sits in a common word. Over the snapshot
// of 2026-09-04 the substring "cat" reads 165 Commander-legal cards and
// the word reads 55, and "angel" reads 165 against 85. So the type
// needle matches on a word boundary, as the land needle does (D-759).
func TestTypeNeedleMatchesAWholeWord(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	idx := fixture(t, typeNeedleCards())
	m := b.themes.match("cats", idx.Tags())
	for _, name := range []string{"Accomplished Automaton", "Sudden Indication"} {
		c, ok := idx.ByName(name)
		if !ok {
			t.Fatalf("the fixture has no %q", name)
		}
		if score, signals := m.score(c); score != 0 {
			t.Errorf("%q holds no Cat and scores %.2f with signals %v, want 0", name, score, signals)
		}
	}
	feline, _ := idx.ByName("Feline Sovereign")
	if !signalOf(m, feline, "type-text:cat") {
		_, signals := m.score(feline)
		t.Errorf("Feline Sovereign names a Cat and reads signals %v, want type-text:cat", signals)
	}
}

// typeNeedleCards are the word-boundary shapes of D-769. Two cards hold
// "cat" inside a longer word, and one names a Cat.
func typeNeedleCards() []tc {
	return []tc{
		{id: "automaton", name: "Accomplished Automaton", typeLine: "Artifact Creature — Construct",
			text: "This creature cannot be blocked by an Automaton.", mv: 7, rank: 100,
			subtypes: []string{"Construct"}},
		{id: "indication", name: "Sudden Indication", typeLine: "Instant",
			text: "Indicate a card. Its controller reveals it.", identity: []mtgv1.Color{U},
			mv: 2, rank: 200},
		{id: "feline", name: "Feline Sovereign", typeLine: "Creature — Cat",
			text:     "Other Cats you control get +1/+1 and have trample.",
			identity: []mtgv1.Color{G}, mv: 4, rank: 300,
			tags: []string{"typal-cat"}, subtypes: []string{"Cat"}},
	}
}

// typeRows maps the name of every type row to its lowercase subtype.
func typeRows(t *testing.T, b *Builder) map[string]string {
	t.Helper()
	out := map[string]string{}
	for name, row := range b.themes.Themes {
		if row.Subtype != "" {
			out[name] = strings.ToLower(row.Subtype)
		}
	}
	if len(out) == 0 {
		t.Fatal("the table holds no type row")
	}
	return out
}

// typeNameSnapshotCards are Zombie payoffs of the snapshot of 2026-09-04
// that name the type and carry no Zombie type of their own. No zombie
// row read one of them before D-769. A signal alone buys no slot, and
// the shortlist is capped, so each card here entered the replay of
// docs/reference/pr58-type-name-2026-09-20.md.
var typeNameSnapshotCards = []string{
	"Necrotic Hex", "Overseer of the Damned", "Nim Deathmantle",
	"Transmogrant Altar", "Crowded Crypt",
}

// TestTypeNameCardsReachATypalShortlist is the second half of F-144 on
// the snapshot. The test needs a snapshot, and `make themes-check` sets
// it.
func TestTypeNameCardsReachATypalShortlist(t *testing.T) {
	idx := snapshotIndex(t)
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	cmdr, ok := idx.ByName("Gisa, the Hellraiser")
	if !ok {
		t.Fatal("the snapshot has no Gisa, the Hellraiser")
	}
	// The gate prompt writes the plural, and the question gate records
	// the singular. Both read the same pool now (D-731, D-769).
	for _, theme := range []string{"zombies", "zombie"} {
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
		for _, name := range typeNameSnapshotCards {
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
