package questions

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// supportedFormats is the whole list the app builds (D-155). HOUSE is not
// here: no menu offers it, and it is reached only by a user who names no
// format and asks for no ban list.
var supportedFormats = []mtgv1.FormatId{
	mtgv1.FormatId_FORMAT_ID_COMMANDER,
	mtgv1.FormatId_FORMAT_ID_STANDARD,
	mtgv1.FormatId_FORMAT_ID_MODERN,
}

// droppedFormats are the four D-155 removed. Every one must decline.
var droppedFormats = []string{"pioneer", "legacy", "vintage", "pauper"}

// TestNoRowOffersADroppedFormat is the invariant under D-155. A question
// that names a format the app refuses contradicts the row that declines
// it, which is the D-150 defect one level up.
//
// The rows that exist to decline a format are exempt, because naming the
// format is their whole job.
func TestNoRowOffersADroppedFormat(t *testing.T) {
	c := load(t)
	for _, r := range c.Rows {
		if declinesFormat(r.ID) {
			continue
		}
		texts := append([]string{r.Text, r.Fallback}, r.Options...)
		for _, text := range texts {
			low := strings.ToLower(text)
			for _, f := range droppedFormats {
				if strings.Contains(low, f) {
					t.Errorf("row %q names %q, which the app does not build: %q", r.ID, f, text)
				}
			}
		}
	}
}

// TestDroppedFormatsResolveToNothing keeps the classifier and the word
// rules from filling the format slot with a format the app refuses. A
// filled slot would skip the decline row entirely.
func TestDroppedFormatsResolveToNothing(t *testing.T) {
	for _, f := range droppedFormats {
		if got, ok := formatIDs[slotWord(f)]; ok && got != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
			t.Errorf("the classifier resolved %q to %v, want no answer", f, got)
		}
		if id, ok := FormatFromWords("i want a " + f + " deck"); ok {
			t.Errorf("the word rule read %q as %v, want no answer", f, id)
		}
		name, near, ok := UnsupportedFormat("i want a " + f + " deck")
		if !ok {
			t.Errorf("%q is not declined by the unsupported-format rule", f)
			continue
		}
		if near != "" {
			t.Errorf("%q offers %q as a substitute, and D-155 names none", f, near)
		}
		if !strings.EqualFold(name, f) {
			t.Errorf("%q declined under the name %q", f, name)
		}
	}
}

// TestSupportedFormatsAllBuild checks that every format the app offers has
// construction rules and a 60-card answer. A format in the menu with no
// rules behind it is the dead-path class of D-76 and D-77.
func TestSupportedFormatsAllBuild(t *testing.T) {
	raw, err := os.ReadFile("../rules/formats.json")
	if err != nil {
		t.Skipf("formats.json not readable: %v", err)
	}
	var cfg struct {
		Formats map[string]map[string]any `json:"formats"`
	}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		t.Fatal(err)
	}
	for _, f := range append(supportedFormats, mtgv1.FormatId_FORMAT_ID_HOUSE) {
		if _, ok := cfg.Formats[f.String()]; !ok {
			t.Errorf("%v is offered and formats.json holds no rules for it", f)
		}
	}
	for name := range cfg.Formats {
		if name == mtgv1.FormatId_FORMAT_ID_HOUSE.String() {
			continue
		}
		var found bool
		for _, f := range supportedFormats {
			if f.String() == name {
				found = true
			}
		}
		if !found {
			t.Errorf("formats.json holds rules for %s, which nothing can select", name)
		}
	}
	// Every 60-card format must answer sixtyCard, or no power row fires.
	for _, f := range []mtgv1.FormatId{
		mtgv1.FormatId_FORMAT_ID_STANDARD,
		mtgv1.FormatId_FORMAT_ID_MODERN,
		mtgv1.FormatId_FORMAT_ID_HOUSE,
	} {
		if !sixtyCard(f) {
			t.Errorf("%v builds 60 cards and sixtyCard says otherwise, so no power row fires", f)
		}
	}
	if sixtyCard(mtgv1.FormatId_FORMAT_ID_COMMANDER) {
		t.Error("Commander is not a 60-card format")
	}
}
