package candidates

import (
	"slices"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The texts are the Oracle texts of the snapshot of 2026-09-04 (F-229).
const (
	jodahText  = "Legendary creatures you control get +X/+X, where X is the number of legendary creatures you control.\nWhenever you cast a legendary spell from your hand, exile cards from the top of your library until you exile a legendary nonland card with lesser mana value. You may cast that card without paying its mana cost. Put the rest on the bottom of your library in a random order."
	jhoiraText = "Whenever you cast a historic spell, draw a card. (Artifacts, legendaries, and Sagas are historic.)"
	sythisText = "Whenever you cast an enchantment spell, you gain 1 life and draw a card."
)

// spellCard is a card for the parse tests, with the type facts that the
// loader derives from the type line.
func spellCard(name string, super, types, sub []string, colors ...mtgv1.Color) *mtgv1.Card {
	return &mtgv1.Card{Name: name, Supertypes: super, CardTypes: types, Subtypes: sub, Colors: colors}
}

// TestCommanderSpellsReadTheCastTrigger is F-229. The cast trigger of a
// commander names the spells that drive it, and each such spell fits the
// class. A phrase with a word that names no kind of spell, such as
// "kicked" or the clause of a spell with no kind, gives no class.
func TestCommanderSpellsReadTheCastTrigger(t *testing.T) {
	subtypes := map[string]bool{"Dog": true, "Cat": true, "Saga": true, "Aura": true, "Equipment": true, "Vehicle": true}
	isSub := func(s string) bool { return subtypes[s] }
	isamaru := spellCard("Isamaru, Hound of Konda", []string{"Legendary"}, []string{"Creature"}, []string{"Dog"}, W)
	bears := spellCard("Grizzly Bears", nil, []string{"Creature"}, []string{"Bear"}, G)
	thopter := spellCard("Ornithopter", nil, []string{"Artifact", "Creature"}, []string{"Thopter"})
	benalia := spellCard("History of Benalia", nil, []string{"Enchantment"}, []string{"Saga"}, W)
	armor := spellCard("Ethereal Armor", nil, []string{"Enchantment"}, []string{"Aura"}, W)
	bolt := spellCard("Lightning Bolt", nil, []string{"Instant"}, nil, R)
	tower := spellCard("Minamo, School at Water's Edge", []string{"Legendary"}, []string{"Land"}, nil)
	for _, tt := range []struct {
		text    string
		phrases []string
		fit     []*mtgv1.Card
		miss    []*mtgv1.Card
	}{
		{jodahText, []string{"legendary"}, []*mtgv1.Card{isamaru}, []*mtgv1.Card{bears, thopter, tower}},
		{jhoiraText, []string{"historic"}, []*mtgv1.Card{isamaru, thopter, benalia}, []*mtgv1.Card{bears, armor, bolt}},
		{sythisText, []string{"enchantment"}, []*mtgv1.Card{benalia, armor}, []*mtgv1.Card{bears, thopter}},
		{"Whenever you cast your first instant spell each turn, if Kalamax is tapped, copy that spell.", []string{"instant"}, []*mtgv1.Card{bolt}, []*mtgv1.Card{armor}},
		{"Whenever you cast a noncreature spell, Stature gets +1/+1.", []string{"noncreature"}, []*mtgv1.Card{bolt, armor}, []*mtgv1.Card{bears, thopter, tower}},
		{"Whenever you cast a Dog spell, create a 1/1 green Cat creature token.\nWhenever you cast a Cat spell, create a 1/1 white Dog creature token.", []string{"dog", "cat"}, []*mtgv1.Card{isamaru}, []*mtgv1.Card{bears}},
		{"Whenever you cast an Aura, Equipment, or Vehicle spell, draw a card.", []string{"aura, equipment, or vehicle"}, []*mtgv1.Card{armor}, []*mtgv1.Card{benalia}},
		{"Whenever you cast a white or red spell, scry 1.", []string{"white or red"}, []*mtgv1.Card{isamaru, bolt}, []*mtgv1.Card{bears, thopter}},
		{"Whenever you cast a kicked spell, you may remove two +1/+1 counters from Verazol.", nil, nil, nil},
		{"Whenever you cast a spell, if the amount of mana spent to cast that spell was less than its mana value, draw a card.", nil, nil, nil},
	} {
		classes := castSpells(tt.text, isSub)
		var got []string
		for _, c := range classes {
			got = append(got, c.Phrase)
		}
		if !slices.Equal(got, tt.phrases) {
			t.Errorf("%q: phrases %q, want %q", tt.text, got, tt.phrases)
			continue
		}
		fits := func(c *mtgv1.Card) bool {
			for _, cl := range classes {
				if cl.fits(c) {
					return true
				}
			}
			return false
		}
		for _, c := range tt.fit {
			if !fits(c) {
				t.Errorf("%q: %s must fit", tt.phrases, c.GetName())
			}
		}
		for _, c := range tt.miss {
			if fits(c) {
				t.Errorf("%q: %s must not fit", tt.phrases, c.GetName())
			}
		}
	}
}

// TestCommanderSpellsJoinTheShortlist is F-229. A reader who gives no
// theme and picks a commander whose cast trigger counts a kind of spell
// gets those spells on the shortlist, marked on theme. Before the fix the
// shortlist held them through a staple role alone, so a card of no staple
// role never reached it: 20 of 570 owned legendary cards for Jodah. The
// class adds no theme word, so no word is unmatched and no theme is thin.
func TestCommanderSpellsJoinTheShortlist(t *testing.T) {
	b, err := New()
	if err != nil {
		t.Fatal(err)
	}
	wubrg := []mtgv1.Color{W, U, B, R, G}
	pool := []tc{
		{id: "jodah", name: "Jodah, the Unifier", typeLine: "Legendary Creature — Human Wizard", text: jodahText, identity: wubrg, mv: 5, rank: 900, subtypes: []string{"Human", "Wizard"}},
		{id: "jhoira", name: "Jhoira, Weatherlight Captain", typeLine: "Legendary Creature — Human Artificer", text: jhoiraText, identity: []mtgv1.Color{U, R}, mv: 4, rank: 900, subtypes: []string{"Human", "Artificer"}},
		{id: "sythis", name: "Sythis, Harvest's Hand", typeLine: "Legendary Enchantment Creature — Nymph", text: sythisText, identity: []mtgv1.Color{G, W}, mv: 2, rank: 900, subtypes: []string{"Nymph"}},
		{id: "isamaru", name: "Isamaru, Hound of Konda", typeLine: "Legendary Creature — Dog", identity: []mtgv1.Color{W}, mv: 1, rank: 9000, subtypes: []string{"Dog"}},
		{id: "thopter", name: "Ornithopter", typeLine: "Artifact Creature — Thopter", text: "Flying", identity: nil, mv: 0, rank: 9000, keywords: []string{"Flying"}, subtypes: []string{"Thopter"}},
		{id: "armor", name: "Ethereal Armor", typeLine: "Enchantment — Aura", text: "Enchant creature\nEnchanted creature gets +1/+1 for each enchantment you control and has first strike.", identity: []mtgv1.Color{W}, mv: 1, rank: 9000, subtypes: []string{"Aura"}},
		{id: "bears", name: "Grizzly Bears", typeLine: "Creature — Bear", identity: []mtgv1.Color{G}, mv: 2, rank: 9000, subtypes: []string{"Bear"}},
		{id: "giant", name: "Hill Giant", typeLine: "Creature — Giant", identity: []mtgv1.Color{R}, mv: 4, rank: 9000, subtypes: []string{"Giant"}},
	}
	idx := fixture(t, pool)
	for _, tt := range []struct {
		commander string
		colors    []mtgv1.Color
		in        string
		out       string
	}{
		{"jodah", wubrg, "Isamaru, Hound of Konda", "Grizzly Bears"},
		{"jhoira", []mtgv1.Color{U, R}, "Ornithopter", "Hill Giant"},
		{"sythis", []mtgv1.Color{G, W}, "Ethereal Armor", "Grizzly Bears"},
	} {
		req := Request{Format: cmdr, Colors: tt.colors, Bracket: 3,
			Owned: map[string]int32{"isamaru": 1, "thopter": 1, "armor": 1, "bears": 1, "giant": 1}, PoolRule: mtgv1.PoolRule_POOL_RULE_OWNED_ONLY}
		without, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := find(without.Candidates, tt.in); ok {
			t.Errorf("%s: %s is listed with no commander, so the test proves nothing", tt.commander, tt.in)
		}
		req.CommanderOracleIDs = []string{tt.commander}
		with, err := b.Build(idx, req)
		if err != nil {
			t.Fatal(err)
		}
		if c, ok := find(with.Candidates, tt.in); !ok || !c.Themed || !c.OnTheme {
			t.Errorf("%s: the cast trigger must pull %s in on theme: %+v, %v", tt.commander, tt.in, c, ok)
		}
		if _, ok := find(with.Candidates, tt.out); ok {
			t.Errorf("%s: %s fits no cast trigger and no staple role, so it must stay out", tt.commander, tt.out)
		}
		if len(with.Theme.Words) > 0 || len(with.Theme.Unmatched) > 0 || with.Stats.ThinTheme {
			t.Errorf("%s: words %v, unmatched %v, thin %v, want none", tt.commander, with.Theme.Words, with.Theme.Unmatched, with.Stats.ThinTheme)
		}
		if len(with.Theme.CommanderSpells) == 0 {
			t.Errorf("%s: the match names no commander spell", tt.commander)
		}
	}
}
