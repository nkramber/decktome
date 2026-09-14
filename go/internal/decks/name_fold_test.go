package decks

import "testing"

// TestFilterKeepReadsANameWithoutItsAccent is D-716: the deck search
// finds a commander whatever the accent the reader typed.
func TestFilterKeepReadsANameWithoutItsAccent(t *testing.T) {
	row := storedDeck{Name: "Mill", CommanderNames: []string{"Gríma, Saruman's Footman"}}
	f := Filter{Query: "grima"}
	if !f.keep(row) {
		t.Error("the query without the accent drops the deck")
	}
}
