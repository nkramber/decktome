package meta

import (
	"strings"
	"testing"
)

// TestParseGoldfishListing reads the Modern user-deck listing of
// 2026-09-03: thirty decks a page, and the total the page names.
func TestParseGoldfishListing(t *testing.T) {
	ids, more := ParseGoldfishListing(readTestdata(t, "mtggoldfish-custom-modern.html"))
	if len(ids) != 3 || ids[0] != "7938665" {
		t.Errorf("ids = %v", ids)
	}
	if !more {
		t.Error("the listing names 14,675 decks, so more pages wait")
	}
	if ids, more := ParseGoldfishListing([]byte("<html></html>")); len(ids) != 0 || more {
		t.Error("an empty page names no deck and no later page")
	}
}

// TestParseGoldfishDeck reads a user deck page of 2026-09-03 as a typical
// Modern list, with the sideboard and the date.
func TestParseGoldfishDeck(t *testing.T) {
	page := readTestdata(t, "mtggoldfish-deck-7938202.html")
	l, err := ParseGoldfishDeck("7938202", page)
	if err != nil {
		t.Fatal(err)
	}
	if l.Source != SourceGoldfish || l.ID != "7938202" || l.Format != FormatModern || l.Tier != TierTypical {
		t.Errorf("list = %+v", l)
	}
	if l.Event != "Torrential Ojer" || l.Date != "2026-09-03" {
		t.Errorf("event %q date %q", l.Event, l.Date)
	}
	if sumCopies(l.Cards) != 60 || sumCopies(l.Sideboard) != 15 {
		t.Errorf("%d main copies, %d side copies", sumCopies(l.Cards), sumCopies(l.Sideboard))
	}
	// A tournament deck page, and a format the model does not cover,
	// answer nil.
	tournament := []byte(strings.Replace(string(page), "User Submitted Deck", "Tournament Deck", 1))
	if l, err := ParseGoldfishDeck("1", tournament); err != nil || l != nil {
		t.Errorf("a tournament deck must answer nil: %v %v", l, err)
	}
	pauper := []byte(strings.Replace(string(page), "Format: Modern", "Format: Pauper", 1))
	if l, err := ParseGoldfishDeck("1", pauper); err != nil || l != nil {
		t.Errorf("a Pauper deck must answer nil: %v %v", l, err)
	}
	if _, err := ParseGoldfishDeck("1", []byte("<html></html>")); err == nil {
		t.Error("a page with no information block must fail")
	}
}
