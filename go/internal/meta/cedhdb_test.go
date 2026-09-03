package meta

import "testing"

func TestParseCEDHDB(t *testing.T) {
	entries, err := ParseCEDHDB(fixture(t, "cedhdb.html"))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 3 {
		t.Fatalf("entries = %d, want 3", len(entries))
	}
	first := entries[0]
	if first.Title != "Rog Reyhan Oops" || first.Section != SectionBrew {
		t.Errorf("first = %q in %q", first.Title, first.Section)
	}
	if len(first.Commanders) != 2 || first.Commanders[0] != "Reyhan, Last of the Abzan" || first.Commanders[1] != "Rograkh, Son of Rohgahh" {
		t.Errorf("first commanders = %v", first.Commanders)
	}
	if first.Date != "2026-08-19" {
		t.Errorf("first date = %q", first.Date)
	}
	if len(first.Links) != 1 || first.Links[0] != "https://moxfield.com/decks/VaR9P-HceECgmgm55DC7ow" {
		t.Errorf("first links = %v", first.Links)
	}
	if first.Competitive() {
		t.Errorf("a brew is not competitive")
	}
	second := entries[1]
	if second.Title != "Kefka Midrange" || len(second.Links) != 2 {
		t.Errorf("second = %q with %d links", second.Title, len(second.Links))
	}
}

func TestParseCEDHDBEmpty(t *testing.T) {
	if _, err := ParseCEDHDB([]byte("<html><body></body></html>")); err == nil {
		t.Fatal("a page with no entry must be an error")
	}
}
