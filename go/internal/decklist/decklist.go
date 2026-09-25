// Package decklist reads a deck list that a user brings: a text file that
// Archidekt exports, or an Arena list (PR-70, D-845). A collection upload
// reads the same lines as copies the user owns. This package keeps the
// sections, so the commander and the sideboard stay apart from the deck.
package decklist

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/collections"
)

// Section names the part of a deck a line belongs to.
type Section int

// The sections. Main is the deck itself, the 99 of a Commander list.
const (
	Main Section = iota
	Commander
	Sideboard
	Companion
)

// maxLines caps one list. A Commander deck is 100 cards, and a list of
// many thousand lines is a paste of the wrong file.
const maxLines = 1000

// ErrTooManyLines refuses a list over maxLines. The whole list fails, so
// no line past the cap goes missing with no report (D-846).
var ErrTooManyLines = fmt.Errorf("a deck list takes at most %d lines", maxLines)

// maxCards caps the copies of one list. A Commander deck is 100 cards,
// and a 60-card deck with its sideboard is 75. A list over the cap asks
// the profile for millions of copies, so the whole list fails (REV-006).
const maxCards = 250

// ErrTooManyCards refuses a list whose copies add up past maxCards.
var ErrTooManyCards = fmt.Errorf("a deck list takes at most %d cards", maxCards)

// maxRawBytes bounds the echoed line text in a report, as a collection
// upload bounds it.
const maxRawBytes = 200

// Line is one card line of a list.
type Line struct {
	collections.Row
	Section Section
}

// List is a parsed deck list.
type List struct {
	Lines []Line
	// Marked says the list names its commander: the Archidekt category
	// Commander, or an Arena Commander section (D-847).
	Marked bool
	// Name is the name an Arena About section gives, or empty.
	Name string
	// Bad holds each line that reads as no card line.
	Bad []*mtgv1.UnresolvedRow
}

// category is the Archidekt category at the end of a line:
// "1x Sol Ring (c21) 263 [Ramp]".
var category = regexp.MustCompile(`\s*\[([^\]]*)\]\s*$`)

// modifier is an Archidekt flag inside a category: "Commander{top}".
var modifier = regexp.MustCompile(`\{[^}]*\}`)

// headers maps a section header of an Arena list to its section. A
// header can start with "//": the Arena export of D-860 opens with
// "// COMMANDER".
var headers = map[string]Section{
	"commander": Commander,
	"deck":      Main,
	"sideboard": Sideboard,
	"companion": Companion,
}

// Parse reads a list. It fails on a read error, on a list over
// maxLines, and on a list over maxCards. A line that reads as no card goes to Bad, and the import
// reports it (D-846).
func Parse(r io.Reader) (*List, error) {
	out := &List{}
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 64*1024), 64*1024)
	section := Main
	about := false
	n, copies := 0, 0
	for sc.Scan() {
		n++
		if n > maxLines {
			return nil, ErrTooManyLines
		}
		raw := sc.Text()
		text := strings.TrimSpace(strings.TrimPrefix(raw, "\ufeff"))
		if text == "" {
			// A blank line ends the commander section of the export of
			// D-860, and no Deck header follows it.
			if section == Commander || section == Companion {
				section = Main
			}
			about = false
			continue
		}
		head := strings.ToLower(strings.TrimSpace(strings.TrimPrefix(text, "//")))
		if s, ok := headers[head]; ok {
			section, about = s, false
			continue
		}
		if head == "about" {
			about = true
			continue
		}
		if about {
			if name, ok := strings.CutPrefix(text, "Name "); ok {
				out.Name = strings.TrimSpace(name)
			}
			continue
		}
		if strings.HasPrefix(text, "//") {
			continue
		}
		line, ok := parseLine(text)
		if !ok {
			out.Bad = append(out.Bad, &mtgv1.UnresolvedRow{
				Line: int32(n), Raw: truncate(raw), Reason: mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW,
			})
			continue
		}
		if copies += line.Quantity; copies > maxCards {
			return nil, ErrTooManyCards
		}
		line.Line, line.Raw = n, truncate(raw)
		if line.Section == Main {
			line.Section = section
		}
		if line.Section == Commander {
			out.Marked = true
		}
		out.Lines = append(out.Lines, line)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// parseLine reads one card line. It takes the Archidekt parts off first:
// the category at the end and the "x" after the count. The rest is an
// Arena line. The category Commander marks the commander (D-847), and
// every other category is dropped (D-848).
func parseLine(text string) (Line, bool) {
	line := Line{Section: Main}
	if m := category.FindStringSubmatch(text); m != nil {
		text = strings.TrimSpace(text[:len(text)-len(m[0])])
		for _, c := range strings.Split(m[1], ",") {
			if strings.EqualFold(strings.TrimSpace(modifier.ReplaceAllString(c, "")), "commander") {
				line.Section = Commander
			}
		}
	}
	if count, rest, ok := strings.Cut(text, " "); ok {
		if trimmed := strings.TrimRight(count, "xX"); trimmed != count && trimmed != "" {
			text = trimmed + " " + rest
		}
	}
	row, ok := collections.ParseLine(text)
	if !ok {
		return Line{}, false
	}
	line.Row = row
	return line, true
}

// Entry is one resolved card of a list, with its copies summed over the
// printings of its section.
type Entry struct {
	Card    *mtgv1.Card
	Count   int
	Section Section
}

// Resolve matches each line to a card of the index, as a collection
// upload does. It sums the copies of one card within one section, in the
// order of the list. A line that matches no card goes to the second
// answer.
func Resolve(l *List, idx *cards.Index) ([]Entry, []*mtgv1.UnresolvedRow) {
	var out []Entry
	bad := append([]*mtgv1.UnresolvedRow(nil), l.Bad...)
	at := map[Section]map[string]int{}
	for _, sec := range []Section{Commander, Main, Sideboard, Companion} {
		at[sec] = map[string]int{}
		for _, line := range l.Lines {
			if line.Section != sec {
				continue
			}
			// One line at a time: the collection resolver merges rows by
			// printing, and a deck sums copies by card.
			entries, missed := collections.Resolve([]collections.Row{line.Row}, idx)
			bad = append(bad, missed...)
			for _, e := range entries {
				card, ok := idx.ByOracleID(e.GetOracleId())
				if !ok {
					continue
				}
				if i, seen := at[sec][e.GetOracleId()]; seen {
					out[i].Count += int(e.GetQuantity())
					continue
				}
				at[sec][e.GetOracleId()] = len(out)
				out = append(out, Entry{Card: card, Count: int(e.GetQuantity()), Section: sec})
			}
		}
	}
	return out, bad
}

func truncate(raw string) string {
	if len(raw) > maxRawBytes {
		cut := maxRawBytes
		for cut > 0 && raw[cut]&0xC0 == 0x80 {
			cut--
		}
		raw = raw[:cut]
	}
	return strings.ToValidUTF8(raw, "\uFFFD")
}
