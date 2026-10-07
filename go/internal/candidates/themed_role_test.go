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

// TestCommanderRateIsNotOnTheme is D-1195. At bracket 5 the rate of the
// commander lists alone makes a card Themed (D-839), and the land cap
// reads that. The theme mark reads OnTheme, and the theme alone sets it.
// So Grizzly Bears, which the lists play and the lifegain theme does not
// match, is Themed and not OnTheme. Soul Warden is both.
func TestCommanderRateIsNotOnTheme(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	rate := fixedRates(map[string]float64{"vanilla": 0.9})
	for _, bracket := range []int32{5, 3} {
		list, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W, G}, Theme: "lifegain", Bracket: bracket, CommanderRate: rate})
		if err != nil {
			t.Fatal(err)
		}
		if bears, ok := find(list.Candidates, "Grizzly Bears"); ok && bears.OnTheme {
			t.Errorf("bracket %d: Grizzly Bears reads on theme from the commander rate alone", bracket)
		}
		warden, ok := find(list.Candidates, "Soul Warden")
		if !ok || !warden.OnTheme || !warden.Themed {
			t.Errorf("bracket %d: Soul Warden reads on theme %v, themed %v; want both", bracket, warden.OnTheme, warden.Themed)
		}
	}
}

// TestThemedPlaneswalkerIsAThreat is D-1197. Three replays of D-1196 left
// out Chandra, Torch of Defiance as "ramp", Garruk, Veiled Butcher as
// "removal", and Ajani Unrelenting as "wipe": a tag marks one loyalty
// ability, and the job word set each beside the staples of that job. A
// planeswalker with a lead now reads threat. An alternate win stays a
// wincon. With no lead, the quality fit keeps the role of the tag.
func TestThemedPlaneswalkerIsAThreat(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	R := mtgv1.Color_COLOR_R
	idx := fixture(t, []tc{
		{id: "chandra", name: "Chandra, Torch of Defiance", typeLine: "Legendary Planeswalker — Chandra", text: "+1: Add {R}{R}.", identity: []mtgv1.Color{R}, mv: 4, rank: 50, tags: []string{"ramp"}},
		{id: "garruk", name: "Garruk, Veiled Butcher", typeLine: "Legendary Planeswalker — Garruk", text: "+2: Up to one target creature gets -4/-1 until your next turn.", identity: []mtgv1.Color{B}, mv: 5, rank: 9000, tags: []string{"removal"}},
		{id: "ajani", name: "Ajani Unrelenting", typeLine: "Legendary Planeswalker — Ajani", text: "−7: Destroy all creatures.", identity: []mtgv1.Color{R, W}, mv: 6, rank: 9000, tags: []string{"sweeper"}},
		{id: "walkerwin", name: "Test Walker of Victory", typeLine: "Legendary Planeswalker — Test", text: "−10: You win the game.", identity: []mtgv1.Color{R}, mv: 5, rank: 9000, tags: []string{"alternate-win-condition"}},
	})
	roleTags := b.themes.roleSets(idx.Tags())
	for _, tt := range []struct {
		id      string
		themed  mtgv1.CardRole
		offList mtgv1.CardRole
	}{
		{"chandra", mtgv1.CardRole_CARD_ROLE_THREAT, mtgv1.CardRole_CARD_ROLE_RAMP},
		{"garruk", mtgv1.CardRole_CARD_ROLE_THREAT, mtgv1.CardRole_CARD_ROLE_REMOVAL},
		{"ajani", mtgv1.CardRole_CARD_ROLE_THREAT, mtgv1.CardRole_CARD_ROLE_WIPE},
		{"walkerwin", mtgv1.CardRole_CARD_ROLE_WINCON, mtgv1.CardRole_CARD_ROLE_WINCON},
	} {
		c, _ := idx.ByOracleID(tt.id)
		if got, _ := assignRole(c, roleTags, true, false); got != tt.themed {
			t.Errorf("%s with a lead: role %s, want %s", c.Name, RoleName(got), RoleName(tt.themed))
		}
		if got, _ := assignRole(c, roleTags, false, false); got != tt.offList {
			t.Errorf("%s with no lead: role %s, want %s", c.Name, RoleName(got), RoleName(tt.offList))
		}
	}
}

// TestSetFillMarksTheCardOfTheSetLimit is D-1198. A set limit lets a card
// with no lead and no staple role onto the list (D-379). SetFill marks
// that card alone: a card on theme and a staple read false, and a list
// with no set limit holds no such card.
func TestSetFillMarksTheCardOfTheSetLimit(t *testing.T) {
	b, _ := New()
	idx := fixture(t, testCards())
	all, err := b.Build(idx, Request{Format: cmdr, Colors: []mtgv1.Color{W, G}, Theme: "lifegain"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range all.Candidates {
		if c.SetFill {
			t.Errorf("no set limit: %s reads set fill", c.Card.GetName())
		}
	}
	set := fixture(t, []tc{
		{id: "insoul", name: "Soul Warden", typeLine: "Creature — Human Cleric",
			text:     "Whenever another creature enters, you gain 1 life.",
			identity: []mtgv1.Color{W}, mv: 1, rank: 200, tags: []string{"lifegain"}, sets: []string{"hob"}},
		{id: "inplain", name: "Hobbit Farmer", typeLine: "Creature — Halfling",
			identity: []mtgv1.Color{W}, mv: 2, rank: 900, sets: []string{"hob"}},
		{id: "inrock", name: "Mind Stone", typeLine: "Artifact", text: "{T}: Add {C}.",
			mv: 2, rank: 3, tags: []string{"ramp", "mana-rock"}, sets: []string{"hob"}},
	})
	limited, err := b.Build(set, Request{Format: cmdr, Colors: []mtgv1.Color{W}, Theme: "lifegain",
		SetCodes: []string{"hob"}, PoolRule: mtgv1.PoolRule_POOL_RULE_ANY_CARD})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]bool{"insoul": false, "inplain": true, "inrock": false}
	for _, c := range limited.Candidates {
		id := c.Card.GetOracleId()
		if w, ok := want[id]; ok {
			if c.SetFill != w {
				t.Errorf("set limit: %s set fill %v, want %v", c.Card.GetName(), c.SetFill, w)
			}
			delete(want, id)
		}
	}
	for id := range want {
		t.Errorf("set limit: the list lost %q", id)
	}
}
