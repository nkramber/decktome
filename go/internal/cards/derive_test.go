package cards

import (
	"os"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// extraByName loads the rules-only fixture rows (Scryfall API, 2026-08-24).
func extraByName(t *testing.T) map[string]*mtgv1.Card {
	t.Helper()
	out := map[string]*mtgv1.Card{}
	for _, path := range []string{"testdata/cards_fixture.jsonl", "../rules/testdata/cards_extra.jsonl"} {
		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		list, err := LoadCards(f, path)
		_ = f.Close()
		if err != nil {
			t.Fatal(err)
		}
		for _, c := range list {
			out[c.Name] = c
		}
	}
	return out
}

func TestPartnerTextFromFixture(t *testing.T) {
	byName := extraByName(t)
	cases := []struct {
		name string
		kind mtgv1.PartnerKind
		text string
	}{
		{"Thrasios, Triton Hero", mtgv1.PartnerKind_PARTNER_KIND_PARTNER, ""},
		{"Abby, Merciless Soldier", mtgv1.PartnerKind_PARTNER_KIND_PARTNER, "Survivors"},
		{"Kratos, Stoic Father", mtgv1.PartnerKind_PARTNER_KIND_PARTNER, "Father & son"},
		{"Donatello, the Brains", mtgv1.PartnerKind_PARTNER_KIND_PARTNER, "Character select"},
		{"April O'Neil, Live on the Scene", mtgv1.PartnerKind_PARTNER_KIND_PARTNER, "Character select"},
		{"Cecily, Haunted Mage", mtgv1.PartnerKind_PARTNER_KIND_FRIENDS_FOREVER, ""},
		{"Pir, Imaginative Rascal", mtgv1.PartnerKind_PARTNER_KIND_WITH, ""},
		{"Clara Oswald", mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION, ""},
		{"Wilson, Refined Grizzly", mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND, ""},
		{"Heliod, Sun-Crowned", mtgv1.PartnerKind_PARTNER_KIND_NONE, ""},
	}
	for _, c := range cases {
		card, ok := byName[c.name]
		if !ok {
			t.Errorf("%s: missing from the fixtures", c.name)
			continue
		}
		if card.Partner != c.kind || card.PartnerText != c.text {
			t.Errorf("%s: partner %v text %q, want %v %q", c.name, card.Partner, card.PartnerText, c.kind, c.text)
		}
	}
	if pir := byName["Pir, Imaginative Rascal"]; pir.PartnerWithName != "Toothy, Imaginary Friend" {
		t.Errorf("Pir partner-with = %q", pir.PartnerWithName)
	}
}

func TestCopyLimitsFromFixture(t *testing.T) {
	byName := extraByName(t)
	cases := []struct {
		name string
		any  bool
		max  int32
	}{
		{"Relentless Rats", true, 0},
		{"Shadowborn Apostle", true, 0},
		{"Seven Dwarves", false, 7},
		{"Nazgûl", false, 9},
		{"Soul Warden", false, 0},
	}
	for _, c := range cases {
		card, ok := byName[c.name]
		if !ok {
			t.Errorf("%s: missing from the fixtures", c.name)
			continue
		}
		if card.AnyCountInDeck != c.any || card.MaxCopiesOverride != c.max {
			t.Errorf("%s: any %v max %d, want %v %d", c.name, card.AnyCountInDeck, card.MaxCopiesOverride, c.any, c.max)
		}
	}
}

func TestCanBeCommanderFromFixture(t *testing.T) {
	byName := extraByName(t)
	cases := map[string]bool{
		"Heliod, Sun-Crowned":           true,
		"Weatherlight":                  true, // legendary Vehicle, CR 903.3 (2026-08-07)
		"Shorikai, Genesis Engine":      true,
		"Dawnsire, Sunstar Dreadnought": true, // legendary Spacecraft
		"Tovolar, Dire Overlord // Tovolar, the Midnight Scourge": true,
		"Esika, God of the Tree // The Prismatic Bridge":          true,
		"Bloodline Keeper // Lord of Lineage":                     false, // back face only (CR 712.8a)
		"Serra Angel":                                             false,
		"Sol Ring":                                                false,
		"Sword of the Animist":                                    false, // legendary Equipment, no P/T box
		"Raised by Giants":                                        false, // a Background is not a commander by itself
		"Grist, the Hunger Tide":                                  false, // known gap: no Scryfall signal for the CDA
	}
	for name, want := range cases {
		card, ok := byName[name]
		if !ok {
			t.Errorf("%s: missing from the fixtures", name)
			continue
		}
		if card.CanBeCommander != want {
			t.Errorf("%s: can_be_commander %v, want %v", name, card.CanBeCommander, want)
		}
	}
}

func TestDeriveTextRules(t *testing.T) {
	mk := func(typeLine, text string) *mtgv1.Card {
		c := &mtgv1.Card{TypeLine: typeLine, OracleText: text}
		derive(c)
		return c
	}
	if c := mk("Creature — Rat", "Search for any number of cards named Rat."); c.AnyCountInDeck {
		t.Error("any-number text must be anchored to the deck sentence")
	}
	if c := mk("Creature — Dwarf", "A deck can have up to 12 cards named Dozen."); c.MaxCopiesOverride != 12 {
		t.Errorf("digit limit = %d", c.MaxCopiesOverride)
	}
	if c := mk("Legendary Creature — Human", "Partner—Survivors"); c.PartnerText != "Survivors" {
		t.Errorf("partner text without reminder = %q", c.PartnerText)
	}
	if c := mk("Legendary Artifact", "This card can't be your commander."); c.CanBeCommander {
		t.Error("a negation must not make a card eligible")
	}
	if c := mk("Legendary Planeswalker — Test", "Test can be your commander."); !c.CanBeCommander {
		t.Error("the text permission must make a card eligible")
	}
	if c := mk("Legendary Artifact — Vehicle", "Crew 1"); c.CanBeCommander {
		t.Error("a Vehicle with no power/toughness box is not eligible")
	}
}
