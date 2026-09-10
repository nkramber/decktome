package collections

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// Detection reads the format of an uploaded file from the file itself
// (PR-34, D-647). The reader drops a file and never names the app it
// came from.
//
// A CSV names itself in its header row, and every importer of this repo
// reads columns by name and never by position. So the header is the
// signature: a set of column names that one format holds and no other
// does.
//
// Two formats that share a header shape are a tie, and a tie never
// guesses. Detect answers false, and the caller asks the reader.

// signature names one format by the header columns that identify it.
// Any one of the column sets is enough, so a format with two shapes
// carries two sets.
type signature struct {
	source mtgv1.ImportSource
	// name is the format in a reader's words, for the message a tie or a
	// miss writes.
	name string
	// columns holds the alternatives. Every name of one alternative must
	// be in the header for that alternative to match.
	columns [][]string
}

// signatures are the CSV formats this app reads. A new platform is one
// entry here and one parser (PR-34).
//
// The ManaBox alternatives mirror its own import rules: a Scryfall ID
// alone, or a name with a set code or a set name. The guide at
// https://www.manabox.app/guides/collection/import-export/ names them,
// read 2026-09-09.
var signatures = []signature{
	{
		source: mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV,
		name:   "ManaBox",
		columns: [][]string{
			{"Scryfall ID"},
			{"Name", "Set code"},
			{"Name", "Set name"},
		},
	},
	{
		source:  mtgv1.ImportSource_IMPORT_SOURCE_MOXFIELD_CSV,
		name:    "Moxfield",
		columns: moxfieldColumns,
	},
}

// requiredAlternatives are the ManaBox key columns, read from the
// signature table so the parser and the detector can not disagree.
func requiredAlternatives() [][]string {
	for _, s := range signatures {
		if s.source == mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV {
			return s.columns
		}
	}
	return nil
}

// SourceName writes a format in a reader's words.
func SourceName(s mtgv1.ImportSource) string {
	for _, sig := range signatures {
		if sig.source == s {
			return sig.name
		}
	}
	if s == mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT {
		return "an Arena list"
	}
	return "an unknown format"
}

// Formats lists every format this app reads, in a reader's words. The
// upload screen shows it, so a reader knows what to export.
func Formats() []string {
	out := make([]string, 0, len(signatures)+1)
	for _, s := range signatures {
		out = append(out, s.name)
	}
	return append(out, "an Arena list")
}

// ErrAmbiguous reports a header that two formats could have written.
// Detection never guesses between them (D-647).
type ErrAmbiguous struct{ Names []string }

func (e *ErrAmbiguous) Error() string {
	return fmt.Sprintf("the header fits %s alike, so the format is not readable from the file", strings.Join(e.Names, " and "))
}

// ErrUnknownFormat reports a file no format claims.
type ErrUnknownFormat struct{}

func (e *ErrUnknownFormat) Error() string {
	return "the file matches no format this app reads: " + strings.Join(Formats(), ", ")
}

// Detect names the format of an uploaded file. It answers the source and
// no error when exactly one format claims the file.
//
// It reads the first line alone, so a file of any size costs the same.
func Detect(content []byte) (mtgv1.ImportSource, error) {
	header, ok := headerOf(content)
	if ok {
		var hits []signature
		for _, s := range signatures {
			if matchesHeader(s, header) {
				hits = append(hits, s)
			}
		}
		switch len(hits) {
		case 1:
			return hits[0].source, nil
		case 0:
			// A CSV header no format claims may still be an Arena list
			// whose first line held a comma, so the read goes on.
		default:
			names := make([]string, 0, len(hits))
			for _, h := range hits {
				names = append(names, h.name)
			}
			return mtgv1.ImportSource_IMPORT_SOURCE_UNSPECIFIED, &ErrAmbiguous{Names: names}
		}
	}
	// An Arena list has no header. Its first content line reads as a
	// count and a card name, and parseArenaLine is the one reader of
	// that shape, so the detector and the parser can not disagree.
	if line, ok := firstContentLine(content); ok {
		if _, ok := parseArenaLine(line); ok {
			return mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT, nil
		}
	}
	return mtgv1.ImportSource_IMPORT_SOURCE_UNSPECIFIED, &ErrUnknownFormat{}
}

// matchesHeader reports whether a header holds every column of any one
// alternative of the signature.
func matchesHeader(s signature, header map[string]bool) bool {
	for _, alt := range s.columns {
		all := true
		for _, c := range alt {
			if !header[c] {
				all = false
				break
			}
		}
		if all && len(alt) > 0 {
			return true
		}
	}
	return false
}

// headerOf reads the first CSV record as a set of column names. It
// answers false when the first line is not a CSV record, and when it
// holds one field alone: a single-column file names no format, and an
// Arena list would read as one.
func headerOf(content []byte) (map[string]bool, bool) {
	cr := csv.NewReader(bytes.NewReader(content))
	cr.FieldsPerRecord = -1
	rec, err := cr.Read()
	if err != nil || len(rec) < 2 {
		return nil, false
	}
	// A UTF-8 BOM before the first cell is part of that cell for
	// encoding/csv (the ManaBox parser strips it the same way).
	rec[0] = strings.TrimPrefix(rec[0], "\uFEFF")
	out := make(map[string]bool, len(rec))
	for _, h := range rec {
		out[strings.TrimSpace(h)] = true
	}
	return out, true
}

// firstContentLine reads the first line that holds something. An Arena
// export starts with a blank line or a "Deck" heading often enough that
// the first line alone is not a fair read.
func firstContentLine(content []byte) (string, bool) {
	sc := bufio.NewScanner(bytes.NewReader(content))
	sc.Buffer(make([]byte, 0, 64<<10), 1<<20)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.EqualFold(line, "Deck") || strings.EqualFold(line, "Sideboard") {
			continue
		}
		return line, true
	}
	return "", false
}

// Parse reads an uploaded file with the parser of one source. It is the
// one place that maps a format onto its reader, so a caller adds no
// switch of its own.
func Parse(source mtgv1.ImportSource, content []byte) ([]Row, []*mtgv1.UnresolvedRow, error) {
	switch source {
	case mtgv1.ImportSource_IMPORT_SOURCE_MANABOX_CSV:
		return ParseManaBoxCSV(bytes.NewReader(content))
	case mtgv1.ImportSource_IMPORT_SOURCE_ARENA_TEXT:
		return ParseArenaText(bytes.NewReader(content))
	case mtgv1.ImportSource_IMPORT_SOURCE_MOXFIELD_CSV:
		return ParseMoxfieldCSV(bytes.NewReader(content))
	}
	return nil, nil, fmt.Errorf("collections: no parser for %s", source)
}
