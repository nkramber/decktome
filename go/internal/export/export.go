// Package export renders a deck as text (D-15). The Arena text is the
// shape ManaBox imports, and the buy list is the shape a shop's
// mass-entry form reads (D-307 to D-309).
package export

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// Lookup answers a card by Oracle id. *cards.Index satisfies it.
type Lookup interface {
	ByOracleID(id string) (*mtgv1.Card, bool)
}

// ErrFormat says the format is not one this package renders.
var ErrFormat = errors.New("unknown export format")

// Render returns the text and the file name for one format. An
// unspecified format is the Arena text.
func Render(d *mtgv1.Deck, cards Lookup, f mtgv1.ExportFormat) (text, fileName string, err error) {
	switch f {
	case mtgv1.ExportFormat_EXPORT_FORMAT_UNSPECIFIED, mtgv1.ExportFormat_EXPORT_FORMAT_ARENA_TEXT:
		return ArenaText(d, cards), FileName(d, mtgv1.ExportFormat_EXPORT_FORMAT_ARENA_TEXT), nil
	case mtgv1.ExportFormat_EXPORT_FORMAT_BUY_LIST_TEXT:
		return BuyListText(d, cards), FileName(d, f), nil
	default:
		return "", "", fmt.Errorf("%w: %s", ErrFormat, f)
	}
}

// ArenaText renders the deck one card per line: the count, the full
// name, then "(SET) number" when a printing is known. The commander
// comes first under "Commander", the main deck under "Deck", and the
// sideboard under "Sideboard". An owned card names its owned printing,
// else the default paper printing (D-307).
func ArenaText(d *mtgv1.Deck, cards Lookup) string {
	var b strings.Builder
	commanders := commanderSet(d)
	if entries := commanderEntries(d, cards); len(entries) > 0 {
		b.WriteString("Commander\n")
		for _, e := range entries {
			b.WriteString(arenaLine(e, cards))
		}
		b.WriteString("\n")
	}
	b.WriteString("Deck\n")
	for _, e := range d.GetCards() {
		if commanders[e.GetOracleId()] {
			continue
		}
		b.WriteString(arenaLine(e, cards))
	}
	if len(d.GetSideboard()) > 0 {
		b.WriteString("\nSideboard\n")
		for _, e := range d.GetSideboard() {
			b.WriteString(arenaLine(e, cards))
		}
	}
	return b.String()
}

func arenaLine(e *mtgv1.DeckCard, cards Lookup) string {
	name := e.GetName()
	var p *mtgv1.Printing
	if c, ok := cards.ByOracleID(e.GetOracleId()); ok {
		if c.GetName() != "" {
			name = c.GetName()
		}
		p = c.GetDefaultPrinting()
	}
	if op := e.GetOwnedPrinting(); e.GetOwned() && op.GetSetCode() != "" && op.GetCollectorNumber() != "" {
		p = op
	}
	if p.GetSetCode() != "" && p.GetCollectorNumber() != "" {
		return fmt.Sprintf("%d %s (%s) %s\n", e.GetCount(), name, strings.ToUpper(p.GetSetCode()), p.GetCollectorNumber())
	}
	return fmt.Sprintf("%d %s\n", e.GetCount(), name)
}

// Row is one card the user still needs (D-308).
type Row struct {
	OracleID        string
	Name            string
	Count           int32
	SetCode         string
	CollectorNumber string
}

// BuyList lists the shortfall of the commander, the main deck, and the
// sideboard: the count minus the owned count, summed per Oracle id. The
// upgrades come back in the second list, with their full count, because
// an upgrade is unowned by definition (D-308).
func BuyList(d *mtgv1.Deck, cards Lookup) (needed, upgrades []Row) {
	type tally struct {
		need  int32
		entry *mtgv1.DeckCard
	}
	order := []string{}
	byID := map[string]*tally{}
	add := func(e *mtgv1.DeckCard, need int32) {
		if need <= 0 {
			return
		}
		t, ok := byID[e.GetOracleId()]
		if !ok {
			t = &tally{entry: e}
			byID[e.GetOracleId()] = t
			order = append(order, e.GetOracleId())
		}
		t.need += need
	}
	for _, e := range commanderEntries(d, cards) {
		add(e, shortfall(e))
	}
	commanders := commanderSet(d)
	for _, e := range d.GetCards() {
		if commanders[e.GetOracleId()] {
			continue
		}
		add(e, shortfall(e))
	}
	for _, e := range d.GetSideboard() {
		add(e, shortfall(e))
	}
	for _, id := range order {
		t := byID[id]
		needed = append(needed, row(t.entry, t.need, cards))
	}
	for _, e := range d.GetUpgrades() {
		if e.GetCount() > 0 {
			upgrades = append(upgrades, row(e, e.GetCount(), cards))
		}
	}
	return needed, upgrades
}

// shortfall is the count the collection does not cover. An entry marked
// owned needs nothing, whatever its owned count says.
func shortfall(e *mtgv1.DeckCard) int32 {
	if e.GetOwned() {
		return 0
	}
	n := e.GetCount() - e.GetOwnedCount()
	if n < 0 {
		return 0
	}
	return n
}

func row(e *mtgv1.DeckCard, count int32, cards Lookup) Row {
	r := Row{OracleID: e.GetOracleId(), Name: e.GetName(), Count: count}
	if c, ok := cards.ByOracleID(e.GetOracleId()); ok {
		if c.GetName() != "" {
			r.Name = c.GetName()
		}
		r.SetCode = strings.ToUpper(c.GetDefaultPrinting().GetSetCode())
		r.CollectorNumber = c.GetDefaultPrinting().GetCollectorNumber()
	}
	return r
}

// BuyListText renders the buy list one "count name" line per row, then
// the upgrades under an "Upgrades" header (D-309). An empty list renders
// as one comment line, so a file is never blank.
func BuyListText(d *mtgv1.Deck, cards Lookup) string {
	needed, upgrades := BuyList(d, cards)
	var b strings.Builder
	if len(needed) == 0 {
		b.WriteString("// Nothing to buy: every card is owned.\n")
	}
	for _, r := range needed {
		fmt.Fprintf(&b, "%d %s\n", r.Count, r.Name)
	}
	if len(upgrades) > 0 {
		b.WriteString("\nUpgrades\n")
		for _, r := range upgrades {
			fmt.Fprintf(&b, "%d %s\n", r.Count, r.Name)
		}
	}
	return b.String()
}

// FileName is the download name: the deck name as a slug, or "deck",
// with "-buy-list" for the buy list, and ".txt".
func FileName(d *mtgv1.Deck, f mtgv1.ExportFormat) string {
	name := slug(d.GetName())
	if name == "" {
		name = "deck"
	}
	if f == mtgv1.ExportFormat_EXPORT_FORMAT_BUY_LIST_TEXT {
		name += "-buy-list"
	}
	return name + ".txt"
}

const maxSlug = 60

func slug(s string) string {
	var b strings.Builder
	dash := false
	for _, r := range strings.ToLower(s) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			dash = false
		case !dash && b.Len() > 0:
			b.WriteByte('-')
			dash = true
		}
		if b.Len() >= maxSlug {
			break
		}
	}
	return strings.Trim(b.String(), "-")
}

func commanderSet(d *mtgv1.Deck) map[string]bool {
	set := map[string]bool{}
	for _, id := range d.GetCommanderOracleIds() {
		set[id] = true
	}
	return set
}

// commanderEntries gives one entry per commander id: the deck's own
// entry when the card list holds it, else an entry of one from the card
// data (D-289). A commander with no entry counts as not owned, because
// nothing recorded the owned count.
func commanderEntries(d *mtgv1.Deck, cards Lookup) []*mtgv1.DeckCard {
	var out []*mtgv1.DeckCard
	for _, id := range d.GetCommanderOracleIds() {
		var found *mtgv1.DeckCard
		for _, e := range d.GetCards() {
			if e.GetOracleId() == id {
				found = e
				break
			}
		}
		if found == nil {
			found = &mtgv1.DeckCard{OracleId: id, Count: 1}
			if c, ok := cards.ByOracleID(id); ok {
				found.Name = c.GetName()
			}
		}
		out = append(out, found)
	}
	return out
}
