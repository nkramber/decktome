package access

import (
	"strings"
	"testing"
)

// TestKeyIsOneIdForOneAddress is D-1075: two spellings of one address
// are one record, and the id holds no slash.
func TestKeyIsOneIdForOneAddress(t *testing.T) {
	if Key(" Ann@Example.com ") != Key("ann@example.com") {
		t.Error("two spellings of one address gave two keys")
	}
	if Key("a/b@example.com") == Key("ab@example.com") || strings.Contains(Key("a/b@example.com"), "/") {
		t.Error("a key must be distinct and hold no slash")
	}
}

func TestCleanRefusesABadRequest(t *testing.T) {
	email, note, err := Clean(" Ann@Example.com ", "  a deck of elves  ")
	if err != nil || email != "ann@example.com" || note != "a deck of elves" {
		t.Fatalf("Clean = %q, %q, %v", email, note, err)
	}
	for _, bad := range []string{"", "ann", "@example.com", "ann@", "ann smith@example.com", strings.Repeat("a", 250) + "@example.com"} {
		if _, _, err := Clean(bad, ""); err == nil {
			t.Errorf("Clean(%q) passed", bad)
		}
	}
	if _, _, err := Clean("ann@example.com", strings.Repeat("é", MaxNoteRunes)); err != nil {
		t.Errorf("a note of %d runes failed: %v", MaxNoteRunes, err)
	}
	if _, _, err := Clean("ann@example.com", strings.Repeat("é", MaxNoteRunes+1)); err == nil {
		t.Error("a note over the cap passed")
	}
}
