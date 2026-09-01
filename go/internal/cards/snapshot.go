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

// nopCloser wraps a reader that needs no close, so maybeGunzip returns
// one type for both cases.
type nopCloser struct{ io.Reader }

func (nopCloser) Close() error { return nil }

// maybeGunzip wraps a stream in a gzip reader when the name ends in .gz.
// The caller closes the result.
func maybeGunzip(r io.Reader, name string) (io.ReadCloser, error) {
	if !strings.HasSuffix(name, ".gz") {
		return nopCloser{r}, nil
	}
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", name, err)
	}
	return gz, nil
}

// readLines runs fn over every line of a JSONL stream, gzip or plain.
func readLines(r io.Reader, name string, fn func(line []byte) error) error {
	rc, err := maybeGunzip(r, name)
	if err != nil {
		return err
	}
	defer func() { _ = rc.Close() }()
	r = rc
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

// SkipLayouts are non-playable layouts, dropped at load time. A front
// card is the art-only front of a reversible card. It carries the name
// of a real card, so it must never answer a name lookup. Checked against
// the snapshot of 2026-08-24: every other layout there is playable.
var SkipLayouts = map[string]bool{
	"token": true, "double_faced_token": true, "emblem": true,
	"art_series": true, "vanguard": true, "scheme": true, "planar": true,
	"front_card": true,
}

// LoadStats counts what LoadCardsStats dropped.
type LoadStats struct {
	// NoOracleID counts playable cards with no Oracle id on the card or
	// its faces (C-8). They can not enter the index.
	NoOracleID int
	// Skipped counts non-playable layouts.
	Skipped int
}

// LoadCards parses an oracle-cards JSONL stream.
// Non-playable layouts (tokens, art series) are dropped here, once.
func LoadCards(r io.Reader, name string) ([]*mtgv1.Card, error) {
	out, _, err := LoadCardsStats(r, name)
	return out, err
}

// LoadCardsStats is LoadCards with drop counts for the caller's log.
func LoadCardsStats(r io.Reader, name string) ([]*mtgv1.Card, LoadStats, error) {
	var out []*mtgv1.Card
	var stats LoadStats
	err := readLines(r, name, func(line []byte) error {
		c, err := parseCard(line)
		if err != nil {
			return err
		}
		if SkipLayouts[c.Layout] {
			stats.Skipped++
			return nil
		}
		if c.OracleId == "" {
			stats.NoOracleID++
			return nil
		}
		out = append(out, c)
		return nil
	})
	return out, stats, err
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
