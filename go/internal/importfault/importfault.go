// Package importfault reads a file the app could not read, for a report
// of kind IMPORT (D-882 to D-889). The collection upload and the deck
// import mark such a file with the error detail UnreadableFile, and the
// page then offers the report form. The server reads the file again
// when the form comes, so the fault it keeps is its own and not the
// client's word (D-596).
//
// Read parses and never resolves. A row the resolver refuses, such as an
// unknown card, shows no form (D-887). So does the one BAD_ROW the
// resolver writes, a merge over the quantity cap: Read can not see it,
// and a report on it reads as no fault.
package importfault

import (
	"bytes"
	"errors"
	"slices"
	"strings"

	"connectrpc.com/connect"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/collections"
	"github.com/nkramber/decktome/go/internal/decklist"
)

const (
	// MaxRows caps the rows of one report (D-885).
	MaxRows = 500
	// MaxRowBytes caps the source text of one row and of the header. A
	// row of a ManaBox export is about 110 bytes, so a cut row is rare.
	// The cap of a parser report is 200 bytes, which cuts too many rows
	// for a fixture of D-888.
	MaxRowBytes = 1024
	// MaxContentBytes is the upload cap of a collection, the larger of
	// the two pages.
	MaxContentBytes = 5 << 20
)

var (
	// ErrNoFault reports a file that reads with no fault, so there is
	// nothing to report.
	ErrNoFault = errors.New("the file reads with no fault")
	// ErrNoCardLine reports a deck list with no line that reads as a
	// card. The deck import answers it, marked as unreadable.
	ErrNoCardLine = errors.New("no line of the list reads as a card")
	errBadPage    = errors.New("import_page: name the collection or the deck page")
	errEmpty      = errors.New("import_content: give the file")
	errTooLarge   = errors.New("import_content: the file is larger than 5 MiB")
)

// Unreadable wraps the error of a file the app could not read. The
// detail tells the page to offer the report form (D-887).
func Unreadable(err error) *connect.Error {
	ce := connect.NewError(connect.CodeInvalidArgument, err)
	if d, derr := connect.NewErrorDetail(&mtgv1.UnreadableFile{}); derr == nil {
		ce.AddDetail(d)
	}
	return ce
}

// IsUnreadable reports whether err carries the UnreadableFile detail.
func IsUnreadable(err error) bool {
	var ce *connect.Error
	if !errors.As(err, &ce) {
		return false
	}
	for _, d := range ce.Details() {
		if v, derr := d.Value(); derr == nil {
			if _, ok := v.(*mtgv1.UnreadableFile); ok {
				return true
			}
		}
	}
	return false
}

// Read reads content as the page reads it, and answers the fault or
// ErrNoFault. Another error is a bad argument.
func Read(page mtgv1.ImportPage, content []byte) (*mtgv1.ImportFault, error) {
	switch {
	case len(content) == 0:
		return nil, errEmpty
	case len(content) > MaxContentBytes:
		return nil, errTooLarge
	}
	var readErr error
	var bad []*mtgv1.UnresolvedRow
	switch page {
	case mtgv1.ImportPage_IMPORT_PAGE_COLLECTION:
		bad, readErr = readCollection(content)
	case mtgv1.ImportPage_IMPORT_PAGE_DECK:
		bad, readErr = readDeck(content)
	default:
		return nil, errBadPage
	}
	if readErr == nil && len(bad) == 0 {
		return nil, ErrNoFault
	}
	lines := splitLines(content)
	fault := &mtgv1.ImportFault{
		Page:      page,
		Header:    cut(lines[0]),
		ByteCount: int64(len(content)),
		RowCount:  int32(countRows(lines)), //nolint:gosec // bounded by MaxContentBytes
	}
	if readErr != nil {
		// Every row of a file that does not read at all is a row that
		// does not parse.
		fault.Error = readErr.Error()
		bad = everyRow(lines)
	} else {
		bad = sourceRows(bad, lines)
	}
	fault.BadRowCount = int32(len(bad)) //nolint:gosec // bounded by MaxContentBytes
	if len(bad) > MaxRows {
		bad = bad[:MaxRows]
	}
	fault.Rows = bad
	return fault, nil
}

// readCollection reads a collection file the way the upload reads it,
// up to the resolver.
func readCollection(content []byte) ([]*mtgv1.UnresolvedRow, error) {
	source, err := collections.Detect(content)
	if err != nil {
		return nil, err
	}
	_, bad, err := collections.Parse(source, content)
	if err != nil {
		return nil, err
	}
	return parseFaults(bad), nil
}

// readDeck reads a deck list the way the deck import reads it, up to the
// resolver.
func readDeck(content []byte) ([]*mtgv1.UnresolvedRow, error) {
	list, err := decklist.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	if len(list.Lines) == 0 {
		return nil, ErrNoCardLine
	}
	return parseFaults(list.Bad), nil
}

// parseFaults keeps the rows of the two reasons that show the form
// (D-887).
func parseFaults(rows []*mtgv1.UnresolvedRow) []*mtgv1.UnresolvedRow {
	var out []*mtgv1.UnresolvedRow
	for _, r := range rows {
		if IsParseFault(r.GetReason()) {
			out = append(out, r)
		}
	}
	return out
}

// IsParseFault reports whether a row of this reason shows the form
// (D-887). An unknown card, another language, and a card no format plays
// show none.
func IsParseFault(reason mtgv1.UnresolvedReason) bool {
	return reason == mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW ||
		reason == mtgv1.UnresolvedReason_UNRESOLVED_REASON_UNKNOWN_VALUE
}

// splitLines splits content into lines, with no line end. It answers at
// least one line.
func splitLines(content []byte) []string {
	text := strings.TrimPrefix(string(content), "\ufeff")
	lines := strings.Split(text, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSuffix(l, "\r")
	}
	return lines
}

// countRows counts the lines after the header that hold text.
func countRows(lines []string) int {
	n := 0
	for _, l := range lines[1:] {
		if strings.TrimSpace(l) != "" {
			n++
		}
	}
	return n
}

// everyRow answers each line after the header that holds text, as a row
// that does not parse.
func everyRow(lines []string) []*mtgv1.UnresolvedRow {
	var out []*mtgv1.UnresolvedRow
	for i, l := range lines[1:] {
		if strings.TrimSpace(l) == "" {
			continue
		}
		out = append(out, &mtgv1.UnresolvedRow{
			Line: int32(i + 2), //nolint:gosec // bounded by MaxContentBytes
			Raw:  cut(l), Reason: mtgv1.UnresolvedReason_UNRESOLVED_REASON_BAD_ROW,
		})
	}
	return out
}

// sourceRows gives each bad row the source line it names, in line order
// and once each. The parser report cuts a row at 200 bytes, and a CSV
// record that does not parse reads there as a note and not as the row.
// A line number off the file keeps the text of the parser.
func sourceRows(bad []*mtgv1.UnresolvedRow, lines []string) []*mtgv1.UnresolvedRow {
	seen := map[int32]bool{}
	var out []*mtgv1.UnresolvedRow
	for _, r := range bad {
		if seen[r.GetLine()] {
			continue
		}
		seen[r.GetLine()] = true
		raw := r.GetRaw()
		if n := int(r.GetLine()); n >= 1 && n <= len(lines) {
			raw = cut(lines[n-1])
		}
		out = append(out, &mtgv1.UnresolvedRow{Line: r.GetLine(), Raw: raw, Reason: r.GetReason()})
	}
	slices.SortStableFunc(out, func(a, b *mtgv1.UnresolvedRow) int { return int(a.GetLine() - b.GetLine()) })
	return out
}

// cut bounds a line at MaxRowBytes on a rune boundary.
func cut(s string) string {
	if len(s) > MaxRowBytes {
		n := MaxRowBytes
		for n > 0 && s[n]&0xC0 == 0x80 {
			n--
		}
		s = s[:n]
	}
	return strings.ToValidUTF8(s, "\uFFFD")
}
