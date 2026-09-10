// Package collections imports and stores user card collections.
//
// Import rules: parse by header name, never by column position. Every
// input row either resolves to a card or comes back in the report with
// a reason. Nothing is dropped in silence (D-23).
package collections

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// Row is one parsed input row before card resolution.
type Row struct {
	Line       int
	Raw        string
	Name       string
	SetCode    string
	SetName    string
	Collector  string
	Quantity   int
	Finish     mtgv1.Finish
	Condition  mtgv1.Condition
	Language   string
	Rarity     string
	ScryfallID string
}

// maxQuantity caps one row. ManaBox has no documented cap. A row above
// this is a typo or an attack, and int32 stays safe.
const maxQuantity = 10_000

// maxRawBytes bounds the echoed row text in a report.
const maxRawBytes = 200

// Known ManaBox value vocabularies. A real export uses Foil =
// normal|foil|etched and Condition = near_mint. The ManaBox guide
// (https://www.manabox.app/guides/collection/import-export/) lists the
// required columns (Name plus Set code or Set name) and gives no value
// vocabulary for Foil or Condition.
//
// The condition names below beyond near_mint follow the ManaBox app's
// condition picker (mint, near_mint, excellent, good, light_played,
// played, poor). They are unverified against an export. An unknown
// value is reported as UNKNOWN_VALUE, never defaulted in silence.
// An empty cell keeps the default: normal, near_mint.
var finishByValue = map[string]mtgv1.Finish{
	"normal": mtgv1.Finish_FINISH_NORMAL,
	"foil":   mtgv1.Finish_FINISH_FOIL,
	"etched": mtgv1.Finish_FINISH_ETCHED,
}

var conditionByValue = map[string]mtgv1.Condition{
	"mint":         mtgv1.Condition_CONDITION_MINT,
	"near_mint":    mtgv1.Condition_CONDITION_NEAR_MINT,
	"excellent":    mtgv1.Condition_CONDITION_EXCELLENT,
	"good":         mtgv1.Condition_CONDITION_GOOD,
	"light_played": mtgv1.Condition_CONDITION_LIGHT_PLAYED,
	"played":       mtgv1.Condition_CONDITION_PLAYED,
	"poor":         mtgv1.Condition_CONDITION_POOR,
}

// csvRows walks a CSV that names its columns in a header row. It reads
// every record, calls build for each one, and reports the records it
// could not read. Nothing is dropped in silence (D-23).
//
// One walker serves every CSV format (PR-34). A format brings its key
// columns and a row builder, and never a copy of this loop.
//
// format names the file in an error, alts are the key column sets, and
// build reads one record. build answers a reason other than UNSPECIFIED
// to reject the row.
func csvRows(r io.Reader, format string, alts [][]string,
	build func(get func(string) string) (Row, mtgv1.UnresolvedReason),
) ([]Row, []*mtgv1.UnresolvedRow, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	header, err := cr.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("%s csv: no header: %w", format, err)
	}
	// A UTF-8 BOM before the first header cell is part of the cell
	// for encoding/csv. Strip it.
	if len(header) > 0 {
		header[0] = strings.TrimPrefix(header[0], "\uFEFF")
	}
	col := map[string]int{}
	for i, h := range header {
		col[strings.TrimSpace(h)] = i
	}
	ok := false
	for _, alt := range alts {
		found := true
		for _, c := range alt {
			if _, has := col[c]; !has {
				found = false
				break
			}
		}
		if found {
			ok = true
			break
		}
	}
	if !ok {
		return nil, nil, fmt.Errorf("%s csv: header %v has no usable key columns", format, header)
	}

	var rows []Row
	var bad []*mtgv1.UnresolvedRow
	for {
		rec, err := cr.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			// A broken record has no fields to echo. ParseError carries
			// the line, and the report names it.
			line := 0
			var pe *csv.ParseError
			if errors.As(err, &pe) {
				line = pe.Line
			}
			bad = append(bad, unresolved(line, fmt.Sprintf("unparseable record at line %d", line), mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW))
			continue
		}
		// FieldPos gives the physical line of the record. A quoted
		// field can span lines, so a record counter is not enough.
		line, _ := cr.FieldPos(0)
		raw := strings.Join(rec, ",")
		get := func(name string) string {
			i, has := col[name]
			if !has || i >= len(rec) {
				return ""
			}
			return strings.TrimSpace(rec[i])
		}
		row, reason := build(get)
		if reason != mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNSPECIFIED {
			bad = append(bad, unresolved(line, raw, reason))
			continue
		}
		row.Line, row.Raw = line, raw
		rows = append(rows, row)
	}
	return rows, bad, nil
}

// parseQuantity reads a count cell. An empty cell means one.
func parseQuantity(v string) (int, bool) {
	if v == "" {
		return 1, true
	}
	q, err := strconv.Atoi(v)
	if err != nil || q < 1 || q > maxQuantity {
		return 0, false
	}
	return q, true
}

// ParseManaBoxCSV reads a ManaBox export. It returns the parsed rows and
// the rows it could not parse. It fails only on a broken header.
func ParseManaBoxCSV(r io.Reader) ([]Row, []*mtgv1.UnresolvedRow, error) {
	// The key columns come from the signature table of detect.go, so the
	// parser and the detector can not disagree about what a ManaBox file
	// looks like (D-647).
	return csvRows(r, "manabox", requiredAlternatives(), func(get func(string) string) (Row, mtgv1.UnresolvedReason) {
		qty, ok := parseQuantity(get("Quantity"))
		if !ok {
			return Row{}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW
		}
		finish, ok := parseFinish(get("Foil"))
		if !ok {
			return Row{}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE
		}
		condition, ok := parseCondition(get("Condition"))
		if !ok {
			return Row{}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE
		}
		return Row{
			Name:       get("Name"),
			SetCode:    get("Set code"),
			SetName:    get("Set name"),
			Collector:  get("Collector number"),
			Quantity:   qty,
			Finish:     finish,
			Condition:  condition,
			Language:   get("Language"),
			Rarity:     strings.ToLower(get("Rarity")),
			ScryfallID: strings.ToLower(get("Scryfall ID")),
		}, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNSPECIFIED
	})
}

// parseFinish maps a Foil cell. An empty cell means normal.
func parseFinish(v string) (mtgv1.Finish, bool) {
	if v == "" {
		return mtgv1.Finish_FINISH_NORMAL, true
	}
	f, ok := finishByValue[strings.ToLower(v)]
	return f, ok
}

// parseCondition maps a Condition cell. An empty cell means near mint.
func parseCondition(v string) (mtgv1.Condition, bool) {
	if v == "" {
		return mtgv1.Condition_CONDITION_NEAR_MINT, true
	}
	c, ok := conditionByValue[strings.ToLower(v)]
	return c, ok
}

// ParseArenaText reads the Arena deck/list format: "4 Lightning Bolt (STA) 42".
// The set and number are optional: "4 Lightning Bolt" is a valid line.
func ParseArenaText(r io.Reader) ([]Row, []*mtgv1.UnresolvedRow, error) {
	var rows []Row
	var bad []*mtgv1.UnresolvedRow
	sc := newLineScanner(r)
	line := 0
	for sc.Scan() {
		line++
		text := strings.TrimSpace(sc.Text())
		if text == "" || strings.HasPrefix(text, "//") {
			continue
		}
		// Section headers in deck exports ("Deck", "Sideboard", "Commander").
		lower := strings.ToLower(text)
		if lower == "deck" || lower == "sideboard" || lower == "commander" || lower == "about" {
			continue
		}
		row, ok := parseArenaLine(text)
		if !ok {
			bad = append(bad, unresolved(line, text, mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW))
			continue
		}
		row.Line = line
		// The raw line names the row a later step could not resolve
		// (D-246).
		row.Raw = text
		rows = append(rows, row)
	}
	return rows, bad, sc.Err()
}

func parseArenaLine(text string) (Row, bool) {
	fields := strings.Fields(text)
	if len(fields) < 2 {
		return Row{}, false
	}
	qty, err := strconv.Atoi(fields[0])
	if err != nil || qty < 1 || qty > maxQuantity {
		return Row{}, false
	}
	rest := fields[1:]
	row := Row{Quantity: qty, Finish: mtgv1.Finish_FINISH_NORMAL, Condition: mtgv1.Condition_CONDITION_NEAR_MINT, Language: "en"}
	// A finish marker trails the line: "*F*" is foil and "*E*" is etched.
	// It sits after the collector number, so it comes off before the set
	// and number pair is read, or the whole line becomes the card name
	// (D-246).
	if len(rest) > 0 {
		if f, ok := arenaFinish(rest[len(rest)-1]); ok {
			row.Finish = f
			rest = rest[:len(rest)-1]
		}
	}
	if len(rest) == 0 {
		return Row{}, false
	}
	// Trailing "(SET) 123" pair, when present.
	if len(rest) >= 2 {
		maybeSet := rest[len(rest)-2]
		if strings.HasPrefix(maybeSet, "(") && strings.HasSuffix(maybeSet, ")") {
			row.SetCode = strings.Trim(maybeSet, "()")
			row.Collector = rest[len(rest)-1]
			rest = rest[:len(rest)-2]
		}
	}
	row.Name = strings.Join(rest, " ")
	return row, row.Name != ""
}

func unresolved(line int, raw string, reason mtgv1.UnresolvedReason) *mtgv1.UnresolvedRow {
	return &mtgv1.UnresolvedRow{Line: int32(line), Raw: truncateRaw(raw), Reason: reason}
}

// truncateRaw cuts raw to maxRawBytes on a rune boundary, so the
// protobuf string stays valid UTF-8. Invalid input bytes are
// replaced too.
func truncateRaw(raw string) string {
	if len(raw) > maxRawBytes {
		cut := maxRawBytes
		for cut > 0 && !utf8.RuneStart(raw[cut]) {
			cut--
		}
		raw = raw[:cut]
	}
	return strings.ToValidUTF8(raw, "\uFFFD")
}

// arenaFinish reads a trailing finish marker of a deck export. Anything
// else is part of the card name.
func arenaFinish(field string) (mtgv1.Finish, bool) {
	switch strings.ToUpper(field) {
	case "*F*":
		return mtgv1.Finish_FINISH_FOIL, true
	case "*E*":
		return mtgv1.Finish_FINISH_ETCHED, true
	}
	return mtgv1.Finish_FINISH_UNSPECIFIED, false
}
