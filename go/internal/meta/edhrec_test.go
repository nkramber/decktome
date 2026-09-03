package meta

import "testing"

func TestEDHRECSlug(t *testing.T) {
	tests := []struct {
		names []string
		want  string
	}{
		{[]string{"Kinnan, Bonder Prodigy"}, "kinnan-bonder-prodigy"},
		{[]string{"Y'shtola, Night's Blessed"}, "yshtola-nights-blessed"},
		{[]string{"Sauron, the Dark Lord"}, "sauron-the-dark-lord"},
		{[]string{"Jace, Vryn's Prodigy // Jace, Telepath Unbound"}, "jace-vryns-prodigy"},
		{[]string{"Thrasios, Triton Hero", "Tymna the Weaver"}, "thrasios-triton-hero-tymna-the-weaver"},
		{[]string{"Atraxa, Praetors' Voice"}, "atraxa-praetors-voice"},
		{[]string{"Bartolomé del Presidio"}, "bartolome-del-presidio"},
		{[]string{"Nazgûl"}, "nazgul"},
	}
	for _, tt := range tests {
		if got := EDHRECSlug(tt.names...); got != tt.want {
			t.Errorf("%v: %q, want %q", tt.names, got, tt.want)
		}
	}
}

func TestParseEDHRECCommander(t *testing.T) {
	c, err := ParseEDHRECCommander("kinnan-bonder-prodigy", fixture(t, "edhrec_commander.json"))
	if err != nil {
		t.Fatal(err)
	}
	if c.Name != "Kinnan, Bonder Prodigy" || c.NumDecks != 21106 || c.Rank != 58 {
		t.Errorf("commander = %+v", c)
	}
	if c.BracketCounts[5] != 4279 || c.BracketCounts[1] != 19 {
		t.Errorf("bracket counts = %v", c.BracketCounts)
	}
	share := c.HighBracketShare()
	if share < 0.88 || share > 0.89 {
		t.Errorf("high bracket share = %.3f, want (1276+4279)/6256", share)
	}
}

func TestParseEDHRECAverageDeck(t *testing.T) {
	l, err := ParseEDHRECAverageDeck("kinnan-bonder-prodigy", fixture(t, "edhrec_average.json"))
	if err != nil {
		t.Fatal(err)
	}
	if l.Tier != TierTypical || l.Format != FormatCommander || l.ID != "kinnan-bonder-prodigy" {
		t.Errorf("list = %+v", l)
	}
	if len(l.Commanders) != 1 || l.Commanders[0] != "Kinnan, Bonder Prodigy" {
		t.Errorf("commanders = %v", l.Commanders)
	}
	// Seven type blocks, four rows each but the one planeswalker.
	if len(l.Cards) != 25 || l.Cards[0].Name != "Arcane Signet" {
		t.Errorf("cards = %d, first %+v", len(l.Cards), l.Cards[0])
	}
}

func TestParseEDHRECTop(t *testing.T) {
	cs, err := ParseEDHRECTop(fixture(t, "edhrec_top.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(cs) != 3 || cs[0].Slug != "yshtola-nights-blessed" || cs[0].Rank != 1 || cs[0].NumDecks != 2471 {
		t.Errorf("top = %+v", cs)
	}
	if EDHRECSlug(cs[0].Name) != cs[0].Slug || EDHRECSlug(cs[1].Name) != cs[1].Slug {
		t.Errorf("the slug rule disagrees with the site: %q vs %q", EDHRECSlug(cs[0].Name), cs[0].Slug)
	}
}
