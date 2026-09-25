package rules

import (
	"os"
	"regexp"
	"slices"
	"testing"

	"github.com/nkramber/decktome/go/internal/cards"
)

// TestDiffFormatsAreTheFormatsOfTheApp holds the legality keys of the M-2
// marker equal to the formats of the app (REV-059). A new format in
// formats.json must join the marker too.
func TestDiffFormatsAreTheFormatsOfTheApp(t *testing.T) {
	raw, err := os.ReadFile("formats.json")
	if err != nil {
		t.Fatal(err)
	}
	var keys []string
	for _, m := range regexp.MustCompile(`"scryfall_key":\s*"([a-z]+)"`).FindAllStringSubmatch(string(raw), -1) {
		if !slices.Contains(keys, m[1]) {
			keys = append(keys, m[1])
		}
	}
	slices.Sort(keys)
	got := slices.Sorted(slices.Values(cards.DiffFormats))
	if !slices.Equal(got, keys) {
		t.Errorf("cards.DiffFormats = %v, want the scryfall keys of formats.json %v", got, keys)
	}
}
