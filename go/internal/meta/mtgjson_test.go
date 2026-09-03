package meta

import "testing"

func TestParseDeckList(t *testing.T) {
	version, entries, err := ParseDeckList(fixture(t, "mtgjson_decklist.json"))
	if err != nil {
		t.Fatal(err)
	}
	if version != "5.3.0+20260902" {
		t.Errorf("version = %s", version)
	}
	if len(entries) != 8 || entries[0].FileName == "" || entries[0].Type != "Commander Deck" {
		t.Errorf("entries = %+v", entries)
	}
}

func TestKeepPrecon(t *testing.T) {
	tests := []struct {
		deckType string
		format   string
		ok       bool
	}{
		{"Commander Deck", FormatCommander, true},
		{"Challenger Deck", FormatSixty, true},
		{"Pioneer Challenger Deck", FormatSixty, true},
		{"Theme Deck", FormatSixty, true},
		{"Brawl Deck", "", true},
		{"Jumpstart", "", false},
		{"Secret Lair Drop", "", false},
		{"Welcome Deck", "", false},
		{"Bundle Land Pack", "", false},
	}
	for _, tt := range tests {
		format, ok := KeepPrecon(tt.deckType)
		if format != tt.format || ok != tt.ok {
			t.Errorf("%s: %q, %v, want %q, %v", tt.deckType, format, ok, tt.format, tt.ok)
		}
	}
}

// TestParsePreconChallenger reads a 60-card product. Two printings of
// one name merge into one row of the list.
func TestParsePreconChallenger(t *testing.T) {
	p, err := ParsePrecon(fixture(t, "mtgjson_challenger.json"))
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != "Rakdos Vampires" || p.Type != "Challenger Deck" || p.ReleaseDate != "2022-04-01" || p.Code != "Q07" {
		t.Errorf("precon = %+v", p)
	}
	if len(p.Cards) != 6 || p.Cards[0].Name != "Voldaren Bloodcaster // Bloodbat Summoner" || p.Cards[0].Count != 4 {
		t.Errorf("cards = %+v", p.Cards)
	}
	if p.Cards[0].ScryfallID == "" || p.Cards[0].OracleID == "" || p.Cards[0].SetCode != "VOW" {
		t.Errorf("first card carries no ids: %+v", p.Cards[0])
	}
	if p.Key() != "RakdosVampires_Q07" {
		t.Errorf("key = %s", p.Key())
	}
	l := p.List()
	if l == nil || l.Format != FormatSixty || l.Tier != TierBaseline || l.Date != "2022-04-01" {
		t.Fatalf("list = %+v", l)
	}
	if len(l.Cards) != 5 {
		t.Fatalf("list cards = %d, want 5 after the merge", len(l.Cards))
	}
	for _, c := range l.Cards {
		if c.Name == "Bloodtithe Harvester" && c.Count != 4 {
			t.Errorf("harvester count = %d, want 4", c.Count)
		}
	}
	if len(l.Sideboard) != 2 || l.Sideboard[1].Count != 3 {
		t.Errorf("sideboard = %+v", l.Sideboard)
	}
}

func TestParsePreconCommander(t *testing.T) {
	p, err := ParsePrecon(fixture(t, "mtgjson_commander.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Commanders) != 1 || p.Commanders[0].Name != "Anowon, the Ruin Thief" {
		t.Fatalf("commanders = %+v", p.Commanders)
	}
	l := p.List()
	if l == nil || l.Format != FormatCommander || len(l.Commanders) != 1 || l.Source != SourceMTGJSON {
		t.Fatalf("list = %+v", l)
	}
}

func TestParsePreconEmpty(t *testing.T) {
	if _, err := ParsePrecon([]byte(`{"data":{"name":"x","mainBoard":[]}}`)); err == nil {
		t.Fatal("a deck with no main board must fail")
	}
}

func TestPreconListBrawl(t *testing.T) {
	p := &Precon{Name: "x", Type: "Brawl Deck", Cards: []PreconCard{{Name: "a", Count: 1}}}
	if p.List() != nil {
		t.Fatal("a Brawl deck serves no fit")
	}
}
