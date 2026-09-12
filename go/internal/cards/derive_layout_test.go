package cards

import (
	"slices"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestTypesFollowTheLayoutRules is F-120. The index merged the types of
// every face, so Legion's Landing counted as a land in a deck. A card has
// only its front face in the library and hand (CR 712.8a, 710.2, 715.2,
// 722.2a). A split card keeps both halves (CR 709.4c), and a modal
// double-faced card keeps its land face (CR 712.12, D-688).
func TestTypesFollowTheLayoutRules(t *testing.T) {
	cases := []struct {
		name, layout        string
		faces               []string
		supers, types, subs []string
	}{
		{"Legion's Landing // Adanto, the First Fort", "transform",
			[]string{"Legendary Enchantment", "Legendary Land"},
			[]string{"Legendary"}, []string{"Enchantment"}, nil},
		{"Westvale Abbey // Ormendahl, Profane Prince", "transform",
			[]string{"Land", "Legendary Creature — Demon"},
			nil, []string{"Land"}, nil},
		{"Azusa's Many Journeys // Likeness of the Seeker", "transform",
			[]string{"Enchantment — Saga", "Enchantment Creature — Human Monk"},
			nil, []string{"Enchantment"}, []string{"Saga"}},
		{"Bonecrusher Giant // Stomp", "adventure",
			[]string{"Creature — Giant", "Instant — Adventure"},
			nil, []string{"Creature"}, []string{"Giant"}},
		{"Adventurous Eater // Have a Bite", "prepare",
			[]string{"Creature — Human Warlock", "Sorcery"},
			nil, []string{"Creature"}, []string{"Human", "Warlock"}},
		{"Budoka Gardener // Dokai, Weaver of Life", "flip",
			[]string{"Creature — Human Monk", "Legendary Creature — Human Monk"},
			nil, []string{"Creature"}, []string{"Human", "Monk"}},
		// The controls: both halves of a split card, and both faces of a
		// modal double-faced card, keep their types.
		{"Fire // Ice", "split",
			[]string{"Instant", "Instant"},
			nil, []string{"Instant"}, nil},
		{"Fell the Profane // Fell Mire", "modal_dfc",
			[]string{"Instant", "Land"},
			nil, []string{"Instant", "Land"}, nil},
	}
	for _, tc := range cases {
		c := &mtgv1.Card{Name: tc.name, Layout: tc.layout, TypeLine: strings.Join(tc.faces, " // ")}
		for _, f := range tc.faces {
			c.Faces = append(c.Faces, &mtgv1.CardFace{TypeLine: f})
		}
		derive(c)
		if !slices.Equal(c.Supertypes, tc.supers) || !slices.Equal(c.CardTypes, tc.types) || !slices.Equal(c.Subtypes, tc.subs) {
			t.Errorf("%s: supertypes %v, types %v, subtypes %v; want %v, %v, %v",
				tc.name, c.Supertypes, c.CardTypes, c.Subtypes, tc.supers, tc.types, tc.subs)
		}
	}
}
