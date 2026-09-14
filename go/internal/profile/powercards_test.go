package profile

import (
	"context"
	"maps"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
)

// TestPowerCardsReadTheProfileRules is the counter of the gate dry run
// (D-706). Fast mana is a nonland, noncreature mana source of mana value
// one or less, so a mana creature does not count, and a Game Changer
// counts by its flag.
func TestPowerCardsReadTheProfileRules(t *testing.T) {
	sol := &mtgv1.Card{OracleId: "o-sol", Name: "Sol Ring", CardTypes: []string{"Artifact"}, ManaValue: 1,
		ProducedMana: []mtgv1.Color{mtgv1.Color_COLOR_C}}
	elves := &mtgv1.Card{OracleId: "o-elves", Name: "Llanowar Elves", CardTypes: []string{"Creature"}, ManaValue: 1,
		ProducedMana: []mtgv1.Color{mtgv1.Color_COLOR_G}}
	study := &mtgv1.Card{OracleId: "o-study", Name: "Rhystic Study", CardTypes: []string{"Enchantment"}, ManaValue: 3,
		GameChanger: true}
	tutors, fast, changers := PowerCards([]*mtgv1.Card{sol, elves, study}, nil)
	if tutors != 0 || fast != 1 || changers != 1 {
		t.Errorf("tutors %d, fast mana %d, Game Changers %d; want 0, 1, 1", tutors, fast, changers)
	}
}

// TestPowerOfNamesEveryFloorACardCounts is D-704: a card counts toward
// each power floor it meets. A tutor with the Game Changer flag counts
// toward two, and a land search is ramp and no tutor.
func TestPowerOfNamesEveryFloorACardCounts(t *testing.T) {
	demonic := &mtgv1.Card{OracleId: "o-demonic", Name: "Demonic Tutor", CardTypes: []string{"Sorcery"}, ManaValue: 2,
		GameChanger: true}
	cultivate := &mtgv1.Card{OracleId: "o-cultivate", Name: "Cultivate", CardTypes: []string{"Sorcery"}, ManaValue: 3}
	sol := &mtgv1.Card{OracleId: "o-sol", Name: "Sol Ring", CardTypes: []string{"Artifact"}, ManaValue: 1,
		ProducedMana: []mtgv1.Color{mtgv1.Color_COLOR_C}}
	tags, err := cards.LoadTags(strings.NewReader(
		`{"id":"t1","slug":"tutor","taggings":[{"oracle_id":"o-demonic"},{"oracle_id":"o-cultivate"}]}`+"\n"+
			`{"id":"t2","slug":"tutor-land","taggings":[{"oracle_id":"o-cultivate"}]}`+"\n"), "tags")
	if err != nil {
		t.Fatal(err)
	}
	of := PowerOf(tags)
	for _, tt := range []struct {
		c    *mtgv1.Card
		want string
	}{
		{demonic, "tutor,game_changer"},
		{cultivate, ""},
		{sol, "fast_mana"},
	} {
		if got := strings.Join(of(tt.c), ","); got != tt.want {
			t.Errorf("%s counts toward %q, want %q", tt.c.GetName(), got, tt.want)
		}
	}
}

// TestPowerFloorsFollowTheBracket is D-704: brackets 4 and 5 hold floors
// for tutors, fast mana, and Game Changers, and no other deck holds one.
func TestPowerFloorsFollowTheBracket(t *testing.T) {
	b, err := LoadBands()
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct {
		name   string
		format mtgv1.FormatId
		power  *mtgv1.PowerLevel
		want   map[string]float64
	}{
		{"bracket 3", mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(3), map[string]float64{}},
		{"no bracket", mtgv1.FormatId_FORMAT_ID_COMMANDER, nil, map[string]float64{}},
		{"bracket 4", mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(4), map[string]float64{KeyTutor: 2, KeyFastMana: 3, KeyGameChanger: 4}},
		{"bracket 5", mtgv1.FormatId_FORMAT_ID_COMMANDER, bracket(5), map[string]float64{KeyTutor: 4, KeyFastMana: 6, KeyGameChanger: 8}},
		{"60-card", mtgv1.FormatId_FORMAT_ID_MODERN, step(mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT), map[string]float64{}},
	} {
		if got := b.PowerFloors(tt.format, tt.power); !maps.Equal(got, tt.want) {
			t.Errorf("%s: floors %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestGameChangerFloorWarnsWhereNoLimitHolds is D-704. Brackets 4 and 5
// set no Game Changer limit, so a deck under the floor of the band warns.
// Bracket 3 keeps its limit, which the rules engine reports, so the
// profile adds no finding there.
func TestGameChangerFloorWarnsWhereNoLimitHolds(t *testing.T) {
	p := newProfiler(t, nil)
	rows := append(shaped(), row{tutor, 1, mtgv1.CardRole_CARD_ROLE_OTHER})
	for _, tt := range []struct {
		bracket int32
		want    string
	}{
		{3, ""},
		{4, "the Game Changer count is 1, and bracket 4 wants 4 or more (Demonic Tutor)"},
		{5, "the Game Changer count is 1, and bracket 5 wants 8 or more (Demonic Tutor)"},
	} {
		_, findings := p.Read(context.Background(), deckOf(tt.bracket, rows...), source(testCards))
		got := ""
		for _, f := range findings {
			if strings.HasPrefix(f.GetMessage(), "the Game Changer count") {
				got = f.GetMessage()
			}
		}
		if got != tt.want {
			t.Errorf("bracket %d: finding %q, want %q", tt.bracket, got, tt.want)
		}
	}
}
