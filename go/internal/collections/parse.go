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

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
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

// manaboxColumns maps the header names this parser reads. Unknown
// columns are ignored. The minimum viable sets mirror the ManaBox rules:
// a Scryfall ID alone, or a name plus set code or set name.
var requiredAlternatives = [][]string{
	{"Scryfall ID"},
	{"Name", "Set code"},
	{"Name", "Set name"},
}

// ParseManaBoxCSV reads a ManaBox export. It returns the parsed rows and
// the rows it could not parse. It fails only on a broken header.
func ParseManaBoxCSV(r io.Reader) ([]Row, []*mtgv1.UnresolvedRow, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	header, err := cr.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("manabox csv: no header: %w", err)
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
	for _, alt := range requiredAlternatives {
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
		return nil, nil, fmt.Errorf("manabox csv: header %v has no usable key columns", header)
	}
	get := func(rec []string, name string) string {
		i, has := col[name]
		if !has || i >= len(rec) {
			return ""
		}
		return strings.TrimSpace(rec[i])
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
		qty := 1
		if q := get(rec, "Quantity"); q != "" {
			qty, err = strconv.Atoi(q)
			if err != nil || qty < 1 || qty > maxQuantity {
				bad = append(bad, unresolved(line, raw, mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW))
				continue
			}
		}
		finish, ok := parseFinish(get(rec, "Foil"))
		if !ok {
			bad = append(bad, unresolved(line, raw, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE))
			continue
		}
		condition, ok := parseCondition(get(rec, "Condition"))
		if !ok {
			bad = append(bad, unresolved(line, raw, mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE))
			continue
		}
		rows = append(rows, Row{
			Line:       line,
			Raw:        raw,
			Name:       get(rec, "Name"),
			SetCode:    get(rec, "Set code"),
			SetName:    get(rec, "Set name"),
			Collector:  get(rec, "Collector number"),
			Quantity:   qty,
			Finish:     finish,
			Condition:  condition,
			Language:   get(rec, "Language"),
			Rarity:     strings.ToLower(get(rec, "Rarity")),
			ScryfallID: strings.ToLower(get(rec, "Scryfall ID")),
		})
	}
	return rows, bad, nil
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
