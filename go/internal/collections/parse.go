// Package collections imports and stores user card collections (PR-4).
//
// Import rules: parse by header name, never by column position (F-2).
// Every input row either resolves to a card or comes back in the report
// with a reason. Nothing is dropped in silence (D-23, PR-4 gate).
package collections

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

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

var finishByValue = map[string]mtgv1.Finish{
	"":       mtgv1.Finish_FINISH_NORMAL,
	"normal": mtgv1.Finish_FINISH_NORMAL,
	"foil":   mtgv1.Finish_FINISH_FOIL,
	"etched": mtgv1.Finish_FINISH_ETCHED,
}

var conditionByValue = map[string]mtgv1.Condition{
	"":             mtgv1.Condition_CONDITION_NEAR_MINT,
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
	line := 1
	for {
		rec, err := cr.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		line++
		if err != nil {
			bad = append(bad, unresolved(line, strings.Join(rec, ","), mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW))
			continue
		}
		qty := 1
		if q := get(rec, "Quantity"); q != "" {
			qty, err = strconv.Atoi(q)
			if err != nil || qty <= 0 {
				bad = append(bad, unresolved(line, strings.Join(rec, ","), mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW))
				continue
			}
		}
		row := Row{
			Line:       line,
			Raw:        strings.Join(rec, ","),
			Name:       get(rec, "Name"),
			SetCode:    get(rec, "Set code"),
			SetName:    get(rec, "Set name"),
			Collector:  get(rec, "Collector number"),
			Quantity:   qty,
			Finish:     finishByValue[strings.ToLower(get(rec, "Foil"))],
			Condition:  conditionByValue[strings.ToLower(get(rec, "Condition"))],
			Language:   get(rec, "Language"),
			Rarity:     strings.ToLower(get(rec, "Rarity")),
			ScryfallID: strings.ToLower(get(rec, "Scryfall ID")),
		}
		if row.Finish == mtgv1.Finish_FINISH_UNSPECIFIED {
			row.Finish = mtgv1.Finish_FINISH_NORMAL
		}
		if row.Condition == mtgv1.Condition_CONDITION_UNSPECIFIED {
			row.Condition = mtgv1.Condition_CONDITION_NEAR_MINT
		}
		rows = append(rows, row)
	}
	return rows, bad, nil
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
	if err != nil || qty <= 0 {
		return Row{}, false
	}
	rest := fields[1:]
	row := Row{Quantity: qty, Finish: mtgv1.Finish_FINISH_NORMAL, Condition: mtgv1.Condition_CONDITION_NEAR_MINT, Language: "en"}
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
	if len(raw) > 200 {
		raw = raw[:200]
	}
	return &mtgv1.UnresolvedRow{Line: int32(line), Raw: raw, Reason: reason}
}
