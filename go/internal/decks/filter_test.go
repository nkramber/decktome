package decks

import (
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// The filter runs in Go over the flat rows (PR-17), so it needs no
// emulator. These tests cover the rules the listing applies.

func mark(b bool) *bool { return &b }

func TestFilterKeep(t *testing.T) {
	row := storedDeck{
		Name:           "Elf Ball",
		SessionID:      "s1",
		FormatID:       int64(mtgv1.FormatId_FORMAT_ID_COMMANDER),
		Favorite:       true,
		PowerBracket:   3,
		CommanderNames: []string{"Marwyn, the Nurturer"},
	}
	for _, tc := range []struct {
		name string
		f    Filter
		want bool
	}{
		{"the empty filter keeps every deck", Filter{}, true},
		{"the same format keeps it", Filter{Format: mtgv1.FormatId_FORMAT_ID_COMMANDER}, true},
		{"another format drops it", Filter{Format: mtgv1.FormatId_FORMAT_ID_MODERN}, false},
		{"the favorite mark keeps it", Filter{Favorite: mark(true)}, true},
		{"the cleared mark drops it", Filter{Favorite: mark(false)}, false},
		{"a name match keeps it", Filter{Query: "elf"}, true},
		{"the match ignores case", Filter{Query: "ELF BALL"}, true},
		{"a commander match keeps it", Filter{Query: "marwyn"}, true},
		{"no match drops it", Filter{Query: "goblin"}, false},
		{"the same bracket keeps it", Filter{PowerBracket: 3}, true},
		{"another bracket drops it", Filter{PowerBracket: 4}, false},
		{"a sixty step drops a Commander deck", Filter{PowerSixtyStep: mtgv1.SixtyStep_SIXTY_STEP_FNM}, false},
		{"the same session keeps it", Filter{SessionID: "s1"}, true},
		{"another session drops it", Filter{SessionID: "s2"}, false},
		{"every part must pass", Filter{Format: mtgv1.FormatId_FORMAT_ID_MODERN, Query: "elf"}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.f.keep(row); got != tc.want {
				t.Errorf("keep = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFilterKeepSixtyStep(t *testing.T) {
	row := storedDeck{Name: "Goblin storm", FormatID: int64(mtgv1.FormatId_FORMAT_ID_MODERN), PowerSixtyStep: int64(mtgv1.SixtyStep_SIXTY_STEP_FNM)}
	if !(Filter{PowerSixtyStep: mtgv1.SixtyStep_SIXTY_STEP_FNM}).keep(row) {
		t.Error("the same step dropped the deck")
	}
	if (Filter{PowerSixtyStep: mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT}).keep(row) {
		t.Error("another step kept the deck")
	}
	if (Filter{PowerBracket: 3}).keep(row) {
		t.Error("a bracket kept a 60-card deck")
	}
	if !(Filter{}).keep(row) {
		t.Error("the empty filter dropped the deck")
	}
}

func TestFilterKeepOldRow(t *testing.T) {
	// A deck written before PR-17 holds no favorite and no commander
	// names. It must still list, and a favorite filter must not keep it.
	old := storedDeck{Name: "Old deck", FormatID: int64(mtgv1.FormatId_FORMAT_ID_COMMANDER)}
	if !(Filter{}).keep(old) {
		t.Error("the empty filter dropped an older deck")
	}
	if (Filter{Favorite: mark(true)}).keep(old) {
		t.Error("the favorite filter kept a deck with no mark")
	}
	if !(Filter{Favorite: mark(false)}).keep(old) {
		t.Error("the cleared filter dropped a deck with no mark")
	}
	if (Filter{Query: "old deck's commander"}).keep(old) {
		t.Error("a query matched a deck with no such text")
	}
	if (Filter{PowerBracket: 3}).keep(old) {
		t.Error("a bracket filter kept a deck with no power")
	}
}

func TestCardCount(t *testing.T) {
	d := &mtgv1.Deck{
		Cards:     []*mtgv1.DeckCard{{Count: 4}, {Count: 1}, {Count: 30}},
		Sideboard: []*mtgv1.DeckCard{{Count: 15}},
		Upgrades:  []*mtgv1.DeckCard{{Count: 2}},
	}
	if got := CardCount(d); got != 35 {
		t.Errorf("CardCount = %d, want 35: the main deck alone", got)
	}
	if got := CardCount(&mtgv1.Deck{}); got != 0 {
		t.Errorf("CardCount of an empty deck = %d", got)
	}
}

func TestCommanderNames(t *testing.T) {
	d := &mtgv1.Deck{
		CommanderOracleIds: []string{"o-marwyn", "o-absent"},
		Cards: []*mtgv1.DeckCard{
			{OracleId: "o-marwyn", Name: "Marwyn, the Nurturer"},
			{OracleId: "o-llanowar", Name: "Llanowar Elves"},
		},
	}
	got := CommanderNames(d)
	if len(got) != 1 || got[0] != "Marwyn, the Nurturer" {
		t.Errorf("CommanderNames = %v, want the one name the card list holds", got)
	}
	if CommanderNames(&mtgv1.Deck{}) != nil {
		t.Error("a deck with no commander gives no names")
	}
}

func TestStoredPower(t *testing.T) {
	if got := storedPower(storedDeck{PowerBracket: 3}); got.GetBracket() != 3 {
		t.Errorf("bracket = %v", got)
	}
	if got := storedPower(storedDeck{PowerSixtyStep: int64(mtgv1.SixtyStep_SIXTY_STEP_FNM)}); got.GetSixtyStep() != mtgv1.SixtyStep_SIXTY_STEP_FNM {
		t.Errorf("sixty step = %v", got)
	}
	if got := storedPower(storedDeck{}); got != nil {
		t.Errorf("a deck with no power gives %v, want nil", got)
	}
}

func TestToStoredAndBack(t *testing.T) {
	d := &mtgv1.Deck{
		Id:                 "d1",
		Name:               "Elf Ball",
		SessionId:          "s1",
		Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
		Power:              &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: 3}},
		Favorite:           true,
		LegalityAsOf:       "2026-08-24",
		BuyCostUsd:         12.5,
		CommanderOracleIds: []string{"o-marwyn"},
		Cards:              []*mtgv1.DeckCard{{OracleId: "o-marwyn", Name: "Marwyn, the Nurturer", Count: 1}, {Count: 99}},
	}
	view := storedToProto("d1", toStored(d, nil))
	if view.GetName() != "Elf Ball" || view.GetFavorite() != true {
		t.Errorf("view = %v", view)
	}
	if view.GetCardCount() != 100 {
		t.Errorf("card count = %d, want 100", view.GetCardCount())
	}
	if view.GetPower().GetBracket() != 3 {
		t.Errorf("power = %v", view.GetPower())
	}
	if len(view.GetCommanderOracleIds()) != 1 {
		t.Errorf("commanders = %v", view.GetCommanderOracleIds())
	}
	if len(view.GetCards()) != 0 {
		t.Error("the list view carries no cards")
	}
}
