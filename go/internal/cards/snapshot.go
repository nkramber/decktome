package cards

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// maxLine bounds one JSONL line. The longest real card lines are far under
// this. A bound keeps a corrupt file from eating memory.
const maxLine = 4 << 20

// readLines runs fn over every line of a JSONL stream, gzip or plain.
func readLines(r io.Reader, name string, fn func(line []byte) error) error {
	if strings.HasSuffix(name, ".gz") {
		gz, err := gzip.NewReader(r)
		if err != nil {
			return fmt.Errorf("open %s: %w", name, err)
		}
		defer func() { _ = gz.Close() }()
		r = gz
	}
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 1<<20), maxLine)
	n := 0
	for sc.Scan() {
		n++
		if len(sc.Bytes()) == 0 {
			continue
		}
		if err := fn(sc.Bytes()); err != nil {
			return fmt.Errorf("%s line %d: %w", name, n, err)
		}
	}
	return sc.Err()
}

// skipLayouts are non-playable layouts, dropped at load time.
var skipLayouts = map[string]bool{
	"token": true, "double_faced_token": true, "emblem": true,
	"art_series": true, "vanguard": true, "scheme": true, "planar": true,
}

// LoadCards parses an oracle-cards JSONL stream.
// Non-playable layouts (tokens, art series) are dropped here, once.
func LoadCards(r io.Reader, name string) ([]*mtgv1.Card, error) {
	var out []*mtgv1.Card
	err := readLines(r, name, func(line []byte) error {
		c, err := parseCard(line)
		if err != nil {
			return err
		}
		if skipLayouts[c.Layout] {
			return nil
		}
		out = append(out, c)
		return nil
	})
	return out, err
}

// LoadPrintings parses a default-cards JSONL stream into printing rows.
func LoadPrintings(r io.Reader, name string) ([]Printing, error) {
	var out []Printing
	err := readLines(r, name, func(line []byte) error {
		p, err := parsePrinting(line)
		if err != nil {
			return err
		}
		if p.ScryfallID != "" && p.OracleID != "" {
			out = append(out, p)
		}
		return nil
	})
	return out, err
}
