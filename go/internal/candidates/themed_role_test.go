package candidates

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// TestThemedCardKeepsAStapleRole is the candidates half of D-1190. A card
// that the theme matches can still take a staple role, so its job word
// says nothing of the theme. The shortlist marks it from Themed, and this
// test proves that the flag holds for each input of the class: a lifegain
// draw card, a Zombie that removes, and a planeswalker card of a new set
// that ramps. A staple with no theme signal stays unthemed.
func TestThemedCardKeepsAStapleRole(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	pool := append(testCards(),
		tc{id: "skinrender", name: "Skinrender", typeLine: "Creature — Phyrexian Zombie", text: "When Skinrender enters, put three -1/-1 counters on target creature.", identity: []mtgv1.Color{B}, mv: 4, rank: 800, tags: []string{"removal"}, subtypes: []string{"Phyrexian", "Zombie"}},
		tc{id: "gravecrawler", name: "Gravecrawler", typeLine: "Creature — Zombie", text: "Gravecrawler can't block. You may cast Gravecrawler from your graveyard as long as you control a Zombie.", identity: []mtgv1.Color{B}, mv: 1, rank: 900, subtypes: []string{"Zombie"}},
		tc{id: "pyromancer", name: "Way of the Pyromancer", typeLine: "Legendary Enchantment", text: "When Way of the Pyromancer enters, empower Jace 2. (Put two loyalty counters on a Jace token you control. If you don't control one, first create a blue Jace planeswalker token with \"[−1]: Surveil 1\" and \"[−3]: Draw a card.\")\nPlaneswalkers you control have \"[+1]: Add {R}.\"", identity: []mtgv1.Color{R}, mv: 2, rank: 9000, tags: []string{"ramp", "synergy-planeswalker"}},
	)
	idx := fixture(t, pool)
	for _, tt := range []struct {
		theme  string
		colors []mtgv1.Color
		card   string
		role   mtgv1.CardRole
	}{
		{"lifegain", []mtgv1.Color{W, B}, "Well of Lost Dreams", mtgv1.CardRole_CARD_ROLE_DRAW},
		{"zombies", []mtgv1.Color{B}, "Skinrender", mtgv1.CardRole_CARD_ROLE_REMOVAL},
		{"planeswalkers", []mtgv1.Color{R}, "Way of the Pyromancer", mtgv1.CardRole_CARD_ROLE_RAMP},
	} {
		list, err := b.Build(idx, Request{Format: cmdr, Colors: tt.colors, Theme: tt.theme, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD})
		if err != nil {
			t.Fatal(err)
		}
		c, ok := find(list.Candidates, tt.card)
		if !ok {
			t.Errorf("%s: %s is not on the shortlist %v", tt.theme, tt.card, names(list.Candidates))
			continue
		}
		if !c.Themed || c.Role != tt.role {
			t.Errorf("%s: %s reads themed %v and role %v, want themed and %v", tt.theme, tt.card, c.Themed, c.Role, tt.role)
		}
		if s, ok := find(list.Candidates, "Sol Ring"); ok && s.Themed {
			t.Errorf("%s: Sol Ring reads themed", tt.theme)
		}
	}
}
