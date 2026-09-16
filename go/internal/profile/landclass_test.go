package profile

import (
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The PR-52 tests of the land readers (F-139, F-147, D-733, D-735). Each
// land text is the Oracle text of the card.

func landSpec(name, text string, subtypes []string, produced ...mtgv1.Color) *mtgv1.Card {
	return card(spec{name: name, types: []string{"Land"}, subtypes: subtypes, produced: produced, text: text})
}

// TestLandClassOfReadsRealLands is D-733. The class follows the lands of
// real lists: an untapped dual, a fetch land, a land that enters untapped
// on a condition, a land with a condition on its mana, a tapped dual, and
// each other land.
func TestLandClassOfReadsRealLands(t *testing.T) {
	U, R, G := mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G
	all := []mtgv1.Color{W, U, B, R, G}
	withColorless := append([]mtgv1.Color{C}, all...)
	ub := map[mtgv1.Color]bool{U: true, B: true}
	grave := landSpec("Watery Grave", "({T}: Add {U} or {B}.)\nAs this land enters, you may pay 2 life. If you don't, it enters tapped.", []string{"Island", "Swamp"}, U, B)
	for _, tt := range []struct {
		card *mtgv1.Card
		want LandClass
	}{
		{grave, LandUntappedDual},
		{landSpec("Underground River", "{T}: Add {C}.\n{T}: Add {U} or {B}. This land deals 1 damage to you.", nil, B, C, U), LandUntappedDual},
		{landSpec("City of Brass", "Whenever this land becomes tapped, it deals 1 damage to you.\n{T}: Add one mana of any color.", nil, all...), LandUntappedDual},
		{landSpec("Polluted Delta", "{T}, Pay 1 life, Sacrifice this land: Search your library for an Island or Swamp card, put it onto the battlefield, then shuffle.", nil), LandFetch},
		{landSpec("Prismatic Vista", "{T}, Pay 1 life, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield, then shuffle.", nil), LandFetch},
		{landSpec("Sunken Hollow", "({T}: Add {U} or {B}.)\nThis land enters tapped unless you control two or more basic lands.", []string{"Island", "Swamp"}, U, B), LandUntappedOnCondition},
		{landSpec("Choked Estuary", "As this land enters, you may reveal an Island or Swamp card from your hand. If you don't, this land enters tapped.\n{T}: Add {U} or {B}.", nil, B, U), LandUntappedOnCondition},
		{landSpec("Fabled Passage", "{T}, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle. Then if you control four or more lands, untap that land.", nil), LandUntappedOnCondition},
		{landSpec("Exotic Orchard", "{T}: Add one mana of any color that a land an opponent controls could produce.", nil, all...), LandManaOnCondition},
		{landSpec("Plaza of Heroes", "{T}: Add {C}.\n{T}: Add one mana of any color. Spend this mana only to cast a legendary spell.\n{T}: Add one mana of any color among legendary permanents you control.\n{3}, {T}, Exile this land: Target legendary creature gains hexproof and indestructible until end of turn.", nil, withColorless...), LandManaOnCondition},
		{landSpec("Spire of Industry", "{T}: Add {C}.\n{T}, Pay 1 life: Add one mana of any color. Activate only if you control an artifact.", nil, withColorless...), LandManaOnCondition},
		{landSpec("Thriving Isle", "This land enters tapped. As it enters, choose a color other than blue.\n{T}: Add {U} or one mana of the chosen color.", nil, all...), LandTappedDual},
		{landSpec("Path of Ancestry", "This land enters tapped.\n{T}: Add one mana of any color in your commander's color identity. When that mana is spent to cast a creature spell that shares a creature type with your commander, scry 1. (Look at the top card of your library. You may put that card on the bottom.)", nil, all...), LandTappedDual},
		{landSpec("Evolving Wilds", "{T}, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle.", nil), LandTappedDual},
		{landSpec("Mana Confluence", "{T}, Pay 1 life: Add one mana of any color.", nil, all...), LandUntappedDual},
		{landSpec("Forbidden Orchard", "{T}: Add one mana of any color.\nWhenever you tap this land for mana, target opponent creates a 1/1 colorless Spirit creature token.", nil, all...), LandUntappedDual},
		{landSpec("Elven Passage", "{T}, Pay 1 life, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle. You may behold an Elf. If you do, untap that land. (To behold an Elf, choose an Elf you control or reveal an Elf card from your hand.)", nil), LandTappedDual},
		{landSpec("Demolition Field", "{T}: Add {C}.\n{2}, {T}, Sacrifice this land: Destroy target nonbasic land an opponent controls. That land's controller may search their library for a basic land card, put it onto the battlefield, then shuffle. You may search your library for a basic land card, put it onto the battlefield, then shuffle.", nil, C), LandOther},
		{landSpec("Myriad Landscape", "This land enters tapped.\n{T}: Add {C}.\n{2}, {T}, Sacrifice this land: Search your library for up to two basic land cards that share a land type, put them onto the battlefield tapped, then shuffle.", nil, C), LandOther},
		{landSpec("Secluded Glen", "As this land enters, you may reveal a Faerie card from your hand. If you don't, this land enters tapped.\n{T}: Add {U} or {B}.", nil, B, U), LandTappedDual},
		{landSpec("Temple of the Dragon Queen", "As this land enters, you may reveal a Dragon card from your hand. This land enters tapped unless you revealed a Dragon card this way or you control a Dragon.\nAs this land enters, choose a color.\n{T}: Add one mana of the chosen color.", nil, all...), LandTappedDual},
		{landSpec("Murky Sewer", "This land enters tapped unless a player has 13 or less life.\n{T}: Add {U} or {B}.", nil, B, U), LandTappedDual},
		{landSpec("Starting Town", "This land enters tapped unless it's your first, second, or third turn of the game.\n{T}: Add {C}.\n{T}, Pay 1 life: Add one mana of any color.", nil, withColorless...), LandUntappedOnCondition},
		{landSpec("Hidden Lair", "{T}: Add {C}.\n{T}: Add {U} or {B}. Activate only if this land entered this turn or if you control a basic land.", nil, B, C, U), LandManaOnCondition},
		{landSpec("Opal Palace", "{T}: Add {C}.\n{1}, {T}: Add one mana of any color in your commander's color identity. If you spend this mana to cast your commander, it enters with a number of additional +1/+1 counters on it equal to the number of times it's been cast from the command zone this game.", nil, withColorless...), LandManaOnCondition},
		{landSpec("Cascading Cataracts", "Indestructible\n{T}: Add {C}.\n{5}, {T}: Add five mana in any combination of colors.", nil, withColorless...), LandOther},
		{landSpec("Baldur's Gate", "{T}: Add {C}.\n{2}, {T}: Add X mana of any one color, where X is the number of other Gates you control.", nil, withColorless...), LandOther},
		{landSpec("Survivors' Encampment", "{T}: Add {C}.\n{T}, Tap an untapped creature you control: Add one mana of any color.", nil, withColorless...), LandOther},
		{landSpec("Lazotep Quarry", "{T}: Add {C}.\n{T}, Sacrifice a creature: Add one mana of any color.\n{X}{2}, {T}, Sacrifice a Desert: Exile target creature card with mana value X from your graveyard. Create a token that's a copy of it, except it's a 4/4 black Zombie. Activate only as a sorcery.", nil, withColorless...), LandOther},
		{landSpec("Gemstone Mine", "This land enters with three mining counters on it.\n{T}, Remove a mining counter from this land: Add one mana of any color. If there are no mining counters on this land, sacrifice it.", nil, all...), LandOther},
		{landSpec("Glimmervoid", "At the beginning of the end step, if you control no artifacts, sacrifice this land.\n{T}: Add one mana of any color.", nil, all...), LandOther},
		{landSpec("Lotus Vale", "If this land would enter, sacrifice two untapped lands instead. If you do, put this land onto the battlefield. If you don't, put it into its owner's graveyard.\n{T}: Add three mana of any one color.", nil, all...), LandOther},
		{landSpec("Hobbit Hole", "{T}, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle.\nHalflingcycling {4} ({4}, Discard this card: Search your library for a Halfling card, reveal it, put it into your hand, then shuffle.)", nil), LandTappedDual},
		{landSpec("Ash Barrens", "{T}: Add {C}.\nBasic landcycling {1} ({1}, Discard this card: Search your library for a basic land card, reveal it, put it into your hand, then shuffle.)", nil, C), LandOther},
		{landSpec("Forgotten Monument", "{T}: Add {C}.\nOther Caves you control have \"{T}, Pay 1 life: Add one mana of any color.\"", nil, withColorless...), LandOther},
		{landSpec("Meteor Crater", "{T}: Choose a color of a permanent you control. Add one mana of that color.", nil, all...), LandOther},
		{fetch, LandOther},
		{landSpec("Takenuma, Abandoned Mire", "{T}: Add {B}.\nChannel — {3}{B}, Discard this card: Mill three cards, then return a creature or planeswalker card from your graveyard to your hand. This ability costs {1} less to activate for each legendary creature you control.", nil, B), LandOther},
		{swamp, LandOther},
		{knight, LandOther},
	} {
		if got := LandClassOf(tt.card, ub); got != tt.want {
			t.Errorf("%s in blue-black reads class %d, want %d", tt.card.GetName(), got, tt.want)
		}
	}
	// With no deck color every color counts.
	for _, tt := range []struct {
		card *mtgv1.Card
		want LandClass
	}{
		{landSpec("Brightclimb Pathway // Grimclimb Pathway", "{T}: Add {W}.\n//\n{T}: Add {B}.", nil, B, W), LandUntappedDual},
		{landSpec("Horizon Canopy", "{T}, Pay 1 life: Add {G} or {W}.\n{1}, {T}, Sacrifice this land: Draw a card.", nil, G, W), LandUntappedDual},
		{landSpec("Grove of the Burnwillows", "{T}: Add {C}.\n{T}: Add {R} or {G}. Each opponent gains 1 life.", nil, C, G, R), LandUntappedDual},
		{landSpec("Fetid Heath", "{T}: Add {C}.\n{W/B}, {T}: Add {W}{W}, {W}{B}, or {B}{B}.", nil, B, C, W), LandManaOnCondition},
		{landSpec("Rith's Grove", "When this land enters, sacrifice it unless you return a non-Lair land you control to its owner's hand.\n{T}: Add {R}, {G}, or {W}.", nil, G, R, W), LandOther},
		{landSpec("Shineshadow Snarl", "As this land enters, you may reveal a Plains or Swamp card from your hand. If you don't, this land enters tapped.\n{T}: Add {W} or {B}.", nil, B, W), LandUntappedOnCondition},
		{landSpec("Rush of Inspiration // Crackling Falls", "Draw two cards. Then discard a card at random unless you pay {E}{E} (two energy counters).\n//\nThis land enters tapped.\n{T}: Add {U} or {R}.", nil, R, U), LandTappedDual},
		{landSpec("Paliano, the High City", "Reveal this card as you draft it. The player to your right chooses a color, you choose another color, then the player to your left chooses a third color.\n{T}: Add one mana of any color chosen as you drafted cards named Paliano, the High City.", nil, all...), LandManaOnCondition},
		{landSpec("Mount Doom", "{T}, Pay 1 life: Add {B} or {R}.\n{1}{B}{R}, {T}: Mount Doom deals 1 damage to each opponent.\n{5}{B}{R}, {T}, Sacrifice Mount Doom and a legendary artifact: Choose up to two creatures, then destroy the rest. Activate only as a sorcery.", nil, B, R), LandUntappedDual},
	} {
		if got := LandClassOf(tt.card, nil); got != tt.want {
			t.Errorf("%s with no deck color reads class %d, want %d", tt.card.GetName(), got, tt.want)
		}
	}
	// A deck of one color holds no dual.
	if got := LandClassOf(grave, map[mtgv1.Color]bool{U: true}); got != LandOther {
		t.Errorf("Watery Grave in mono-blue reads class %d, want %d", got, LandOther)
	}
	// With no deck color every color counts, so Marsh Flats finds two.
	if got := LandClassOf(fetch, nil); got != LandFetch {
		t.Errorf("Marsh Flats with no deck color reads class %d, want %d", got, LandFetch)
	}
}

// TestAFetchLandCountsTheColorsItCanFind is F-147 and D-735. The profile
// counted a fetch land as a source of every deck color. Marsh Flats
// searches for a Plains or Swamp card, so in a blue-black deck of basic
// lands it finds no blue land. Watery Grave carries the Swamp type, so it
// gives Marsh Flats a blue land to find.
func TestAFetchLandCountsTheColorsItCanFind(t *testing.T) {
	U := mtgv1.Color_COLOR_U
	island := card(spec{name: "Island", types: []string{"Land"}, supers: []string{"Basic"}, subtypes: []string{"Island"}, produced: []mtgv1.Color{U}})
	grave := landSpec("Watery Grave", "({T}: Add {U} or {B}.)\nAs this land enters, you may pay 2 life. If you don't, it enters tapped.", []string{"Island", "Swamp"}, U, B)
	passage := landSpec("Fabled Passage", "{T}, Sacrifice this land: Search your library for a basic land card, put it onto the battlefield tapped, then shuffle. Then if you control four or more lands, untap that land.", nil)
	blueSpell := card(spec{name: "Blue Spell", cost: "{1}{U}", mv: 2, types: []string{"Creature"}, colors: []mtgv1.Color{U}})
	blackSpell := card(spec{name: "Black Spell", cost: "{1}{B}", mv: 2, types: []string{"Creature"}, colors: []mtgv1.Color{B}})
	src := source{}
	for _, c := range []*mtgv1.Card{island, swamp, fetch, grave, passage, blueSpell, blackSpell} {
		src[c.GetOracleId()] = c
	}
	have := func(extra ...row) map[mtgv1.Color]float64 {
		rows := append([]row{{island, 10, 0}, {swamp, 10, 0}, {blueSpell, 20, 0}, {blackSpell, 20, 0}}, extra...)
		deck := &mtgv1.Deck{Format: &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_MODERN}}
		for _, r := range rows {
			deck.Cards = append(deck.Cards, &mtgv1.DeckCard{OracleId: r.c.GetOracleId(), Name: r.c.GetName(), Count: r.n, Role: r.role})
		}
		out := map[mtgv1.Color]float64{}
		for _, s := range ColorSources(deck, src) {
			out[s.Color] = s.Have
		}
		return out
	}
	if got := have(row{fetch, 1, 0}); got[U] != 10 || got[B] != 11 {
		t.Errorf("Marsh Flats beside basic lands reads blue %v and black %v, want 10 and 11", got[U], got[B])
	}
	if got := have(row{fetch, 1, 0}, row{grave, 1, 0}); got[U] != 12 || got[B] != 12 {
		t.Errorf("Marsh Flats beside Watery Grave reads blue %v and black %v, want 12 and 12", got[U], got[B])
	}
	if got := have(row{passage, 1, 0}); got[U] != 11 || got[B] != 11 {
		t.Errorf("Fabled Passage reads blue %v and black %v, want 11 and 11", got[U], got[B])
	}
}
