package collections

import (
	"errors"
	"strings"
	"testing"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// The reader drops a file and never names the app it came from (D-647).
// Every case below is a file the app receives, and the detector must
// name its format or refuse to guess.

func TestDetectNamesTheFormat(t *testing.T) {
	for name, tc := range map[string]struct {
		content string
		want    mtgv1.ImportSource
	}{
		"a ManaBox export by set code": {
			"Name,Set code,Quantity,Foil\nSol Ring,C21,1,normal\n",
			mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		},
		"a ManaBox export by set name": {
			"Name,Set name,Quantity\nSol Ring,Commander 2021,1\n",
			mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		},
		"a ManaBox export by Scryfall id alone": {
			"Scryfall ID,Quantity\n0004ebd0-dfd6-4276-b435-40d8d5f2f6e9,1\n",
			mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		},
		"a ManaBox export with a byte order mark": {
			"\uFEFFName,Set code,Quantity\nSol Ring,C21,1\n",
			mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		},
		"an Arena list": {
			"4 Lightning Bolt (STA) 42\n1 Sol Ring (C21) 263\n",
			mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT,
		},
		"an Arena list with no set": {
			"4 Lightning Bolt\n",
			mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT,
		},
		"an Arena list under its headings": {
			"Deck\n\n4 Lightning Bolt (STA) 42\n\nSideboard\n2 Negate (M21) 55\n",
			mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT,
		},
		"a Moxfield export": {
			`"Count","Tradelist Count","Name","Edition","Condition","Language","Foil"` + "\n" +
				`"1","1","Sol Ring","c21","Near Mint","English",""` + "\n",
			mtgv1.ImportSource_IMPORT_SOURCE_MOXFIELD_CSV,
		},
	} {
		t.Run(name, func(t *testing.T) {
			got, err := Detect([]byte(tc.content))
			if err != nil {
				t.Fatalf("Detect = %v", err)
			}
			if got != tc.want {
				t.Errorf("Detect = %s, want %s", got, tc.want)
			}
			// A detected format must parse: a detector that names a
			// format the parser refuses is worse than no detector.
			if _, _, err := Parse(got, []byte(tc.content)); err != nil {
				t.Errorf("the detected format does not parse: %v", err)
			}
		})
	}
}

// TestDetectRefusesAFileItDoesNotKnow keeps a wrong guess out. A file
// parsed as the wrong format reports every row as bad, and the reader
// reads that as a broken collection rather than a wrong upload.
func TestDetectRefusesAFileItDoesNotKnow(t *testing.T) {
	for name, content := range map[string]string{
		"a spreadsheet of something else": "First,Last,Email\nAnn,Lee,ann@example.com\n",
		"an empty file":                   "",
		"prose":                           "Here are the cards I own, more or less.\n",
	} {
		t.Run(name, func(t *testing.T) {
			got, err := Detect([]byte(content))
			if err == nil {
				t.Fatalf("Detect named %s for a file it does not read", got)
			}
			var unknown *ErrUnknownFormat
			if !errors.As(err, &unknown) {
				t.Fatalf("err = %v, want the unknown-format refusal", err)
			}
			// The message names what the app does read, so the reader
			// knows what to export.
			for _, f := range Formats() {
				if !strings.Contains(err.Error(), f) {
					t.Errorf("the refusal does not name %s: %v", f, err)
				}
			}
		})
	}
}

// TestDetectNeverGuessesBetweenTwoFormats holds the tie rule of D-647.
// A header two formats could have written is not a header that names
// one.
func TestDetectNeverGuessesBetweenTwoFormats(t *testing.T) {
	// A second format that claims the same columns as ManaBox. The table
	// is restored before the test ends, so no other test reads it.
	old := signatures
	t.Cleanup(func() { signatures = old })
	signatures = append(append([]signature{}, old...), signature{
		source: mtgv1.ImportSource_IMPORT_SOURCE_UNSPECIFIED,
		name:   "Twinbox",
		columns: [][]string{
			{"Name", "Set code"},
		},
	})
	_, err := Detect([]byte("Name,Set code,Quantity\nSol Ring,C21,1\n"))
	if err == nil {
		t.Fatal("a header two formats claim was read as one of them")
	}
	var tie *ErrAmbiguous
	if !errors.As(err, &tie) {
		t.Fatalf("err = %v, want the tie refusal", err)
	}
	if len(tie.Names) != 2 {
		t.Errorf("the tie names %v, want both formats", tie.Names)
	}
	for _, want := range []string{"ManaBox", "Twinbox"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("the tie message does not name %s: %v", want, err)
		}
	}
}

// TestTheDetectorAndTheParserReadOneTable holds the rule that made
// requiredAlternatives a function. A parser that accepts a header the
// detector never names is a file the reader can not upload, and the
// other way round is a file that detects and then fails every row.
func TestTheDetectorAndTheParserReadOneTable(t *testing.T) {
	alts := requiredAlternatives()
	if len(alts) == 0 {
		t.Fatal("the ManaBox signature holds no key columns")
	}
	for _, alt := range alts {
		header := strings.Join(alt, ",") + ",Quantity\n"
		row := strings.Repeat("x,", len(alt)) + "1\n"
		content := []byte(header + row)
		got, err := Detect(content)
		if err != nil {
			t.Errorf("the detector refuses a header the parser takes, %v: %v", alt, err)
			continue
		}
		if got != mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV {
			t.Errorf("header %v detected as %s", alt, got)
		}
		if _, _, err := ParseManaBoxCSV(strings.NewReader(string(content))); err != nil {
			t.Errorf("the parser refuses a header the detector takes, %v: %v", alt, err)
		}
	}
}

// TestParseNamesNoParserForAnUnknownSource keeps the one mapping honest.
// A source with no parser must say so and never read as an empty file.
func TestParseNamesNoParserForAnUnknownSource(t *testing.T) {
	_, _, err := Parse(mtgv1.ImportSource_IMPORT_SOURCE_UNSPECIFIED, []byte("anything"))
	if err == nil {
		t.Fatal("an unspecified source parsed")
	}
	if !strings.Contains(err.Error(), "no parser") {
		t.Errorf("err = %v, want the missing parser named", err)
	}
}

// TestFormatsNamesEveryReaderFacingFormat keeps the upload screen and
// the detector in step. A format the app reads and never names is one a
// reader never knows to export.
func TestFormatsNamesEveryReaderFacingFormat(t *testing.T) {
	got := Formats()
	if len(got) != len(signatures)+1 {
		t.Errorf("Formats lists %d, want one per signature plus the Arena list", len(got))
	}
	for _, s := range signatures {
		if SourceName(s.source) != s.name {
			t.Errorf("SourceName(%s) = %q, want %q", s.source, SourceName(s.source), s.name)
		}
	}
	if SourceName(mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT) == "an unknown format" {
		t.Error("the Arena list has no reader-facing name")
	}
}
