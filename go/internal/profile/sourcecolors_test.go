package profile

import (
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestSourceColorsOfCountsOnlyAReliableTapAbility is F-103. A card counts
// as a color source only through a tap mana ability that sacrifices
// nothing, and never for a color its own cost needs.
func TestSourceColorsOfCountsOnlyAReliableTapAbility(t *testing.T) {
	U, R, G := mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G
	all := []mtgv1.Color{W, U, B, R, G}
	cases := []struct {
		name string
		c    *mtgv1.Card
		want []mtgv1.Color
	}{
		{"a signet", card(spec{name: "Orzhov Signet", cost: "{2}", types: []string{"Artifact"},
			produced: []mtgv1.Color{W, B}, text: "{1}, {T}: Add {W}{B}."}), []mtgv1.Color{W, B}},
		{"a mana creature of its own color", card(spec{name: "Elvish Mystic", cost: "{G}", types: []string{"Creature"},
			produced: []mtgv1.Color{G}, text: "{T}: Add {G}."}), nil},
		{"a mana creature of every color", card(spec{name: "Birds of Paradise", cost: "{G}", types: []string{"Creature"},
			produced: all, text: "Flying\n{T}: Add one mana of any color."}), []mtgv1.Color{W, U, B, R}},
		{"a ritual", card(spec{name: "Dark Ritual", cost: "{B}", types: []string{"Instant"},
			produced: []mtgv1.Color{B}, text: "Add {B}{B}{B}."}), nil},
		{"a sacrifice altar", card(spec{name: "Phyrexian Altar", cost: "{3}", types: []string{"Artifact"},
			produced: all, text: "Sacrifice a creature: Add one mana of any color."}), nil},
		{"a Treasure maker", card(spec{name: "Pitiless Plunderer", cost: "{3}{B}", types: []string{"Creature"},
			produced: all, text: "Whenever another creature you control dies, create a Treasure token. " +
				"(It's an artifact with \"{T}, Sacrifice this token: Add one mana of any color.\")"}), nil},
		{"a granted ability", card(spec{name: "Elven Chorus", cost: "{3}{G}", types: []string{"Enchantment"},
			produced: all, text: "Creatures you control have \"{T}: Add one mana of any color.\""}), []mtgv1.Color{W, U, B, R}},
	}
	for _, tc := range cases {
		got := sourceColorsOf(tc.c)
		slices.Sort(got)
		want := slices.Clone(tc.want)
		slices.Sort(want)
		if !slices.Equal(got, want) {
			t.Errorf("%s: %v, want %v", tc.name, got, want)
		}
	}
}
