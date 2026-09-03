package meta

import (
	"bytes"
	"fmt"
	"html"
	"regexp"
	"strconv"
	"strings"
)

// MTGTop8Base is the paper tournament site. A format page lists the last
// events with a paper or an online mark, an event page lists the decks
// with their placements, and the text export of a deck is one plain
// list (read 2026-09-03, pr14c-sources-2026-09-03). The site publishes
// no robots file.
const MTGTop8Base = "https://www.mtgtop8.com"

// mtgtop8Formats maps the site's format code to the model's word.
// Duel Commander, code EDH, is not Commander (D-112).
var mtgtop8Formats = map[string]string{"MO": FormatModern, "ST": FormatStandard, "cEDH": FormatCommander}

// MTGTop8FormatCodes lists the codes the reader walks, in order (D-504).
var MTGTop8FormatCodes = []string{"MO", "ST", "cEDH"}

// MTGTop8FormatURL is the first page of a format's event list.
func MTGTop8FormatURL(base, code string) string { return base + "/format?f=" + code }

// MTGTop8PageURL resolves a page href of the event list, which the page
// writes as "?f=MO&meta=54&cp=2".
func MTGTop8PageURL(base, href string) string { return base + "/format" + href }

// MTGTop8EventURL is the page of one event.
func MTGTop8EventURL(base, id, code string) string {
	return base + "/event?e=" + id + "&f=" + code
}

// MTGTop8DeckURL is the text export of one deck.
func MTGTop8DeckURL(base, id string) string { return base + "/mtgo?d=" + id }

// MTGTop8Event is one row of a format page.
type MTGTop8Event struct {
	ID   string
	Name string
	// Date is YYYY-MM-DD.
	Date string
	// Stars is the site's size mark, one to three.
	Stars int
	// Paper says the row carries the paper mark and not the online one.
	Paper bool
}

var (
	mtgtop8Row   = regexp.MustCompile(`(?s)<tr class=hover_tr>(.*?)</tr>`)
	mtgtop8Href  = regexp.MustCompile(`href=event\?e=(\d+)&f=[A-Za-z]+>([^<]*)</a>`)
	mtgtop8Date  = regexp.MustCompile(`(\d{2})/(\d{2})/(\d{2})`)
	mtgtop8Pages = regexp.MustCompile(`href="?(\?f=[A-Za-z]+&meta=\d+&cp=\d+)`)
)

// ParseMTGTop8Format reads the event rows of a format page, each event
// once and in page order, and the hrefs of the other pages of the list.
// The page holds two event tables of one row shape, the major events
// and the last twenty.
func ParseMTGTop8Format(page []byte) (events []MTGTop8Event, pages []string) {
	seen := map[string]bool{}
	for _, m := range mtgtop8Row.FindAllSubmatch(page, -1) {
		row := m[1]
		h := mtgtop8Href.FindSubmatch(row)
		if h == nil {
			continue
		}
		id := string(h[1])
		if seen[id] {
			continue
		}
		seen[id] = true
		ev := MTGTop8Event{ID: id, Name: plainText(string(h[2]))}
		ev.Stars = bytes.Count(row, []byte("/graph/star.png"))
		ev.Paper = bytes.Contains(row, []byte("/graph/online/paper.png"))
		if d := mtgtop8Date.FindSubmatch(row); d != nil {
			ev.Date = "20" + string(d[3]) + "-" + string(d[2]) + "-" + string(d[1])
		}
		events = append(events, ev)
	}
	seenPage := map[string]bool{}
	for _, m := range mtgtop8Pages.FindAllSubmatch(page, -1) {
		href := string(m[1])
		if seenPage[href] {
			continue
		}
		seenPage[href] = true
		pages = append(pages, href)
	}
	return events, pages
}

// MTGTop8EventInfo is the head of an event page.
type MTGTop8EventInfo struct {
	Name string
	// Date is YYYY-MM-DD.
	Date string
	// Format is the site's code, from the deck links.
	Format string
	// Players is the field, 0 when the page names none.
	Players int
	// Online says the page names mtgo.com as its source: the MTGO lane
	// holds the same lists, so the reader skips it (D-504).
	Online bool
}

// MTGTop8Row is one deck of an event page.
type MTGTop8Row struct {
	// Placement is the first number of the place, so "3-4" reads 3.
	Placement int
	DeckID    string
	Name      string
	Player    string
}

var (
	mtgtop8Title   = regexp.MustCompile(`<div class=event_title>([^<]*)</div>`)
	mtgtop8Players = regexp.MustCompile(`(\d+) players`)
	mtgtop8Block   = regexp.MustCompile(`(?s)<div class=(?:chosen_tr|hover_tr)[^>]*>\s*<div style="display:flex;align-items:center;">(.*?)class=player href=search\?player=[^>]*>([^<]*)</a>`)
	mtgtop8Place   = regexp.MustCompile(`class=S14>\s*(\d+)(?:-\d+)?\s*</div>`)
	mtgtop8Deck    = regexp.MustCompile(`href=\?e=\d+&d=(\d+)&f=([A-Za-z]+)>([^<]+)</a>`)
	mtgtop8Source  = regexp.MustCompile(`Source: <a[^>]*href=https?://www\.mtgo\.com`)
)

// ParseMTGTop8Event reads the head and the deck rows of an event page.
// A page with no deck row is a parse failure (M-6).
func ParseMTGTop8Event(page []byte) (MTGTop8EventInfo, []MTGTop8Row, error) {
	var info MTGTop8EventInfo
	if m := mtgtop8Title.FindSubmatch(page); m != nil {
		info.Name = plainText(string(m[1]))
	}
	if d := mtgtop8Date.FindSubmatch(page); d != nil {
		info.Date = "20" + string(d[3]) + "-" + string(d[2]) + "-" + string(d[1])
	}
	if m := mtgtop8Players.FindSubmatch(page); m != nil {
		info.Players, _ = strconv.Atoi(string(m[1]))
	}
	info.Online = mtgtop8Source.Match(page)
	var rows []MTGTop8Row
	seen := map[string]bool{}
	for _, m := range mtgtop8Block.FindAllSubmatch(page, -1) {
		block := m[1]
		d := mtgtop8Deck.FindSubmatch(block)
		if d == nil {
			continue
		}
		id := string(d[1])
		if seen[id] {
			continue
		}
		seen[id] = true
		if info.Format == "" {
			info.Format = string(d[2])
		}
		row := MTGTop8Row{DeckID: id, Name: plainText(string(d[3])), Player: plainText(string(m[2]))}
		if p := mtgtop8Place.FindSubmatch(block); p != nil {
			row.Placement, _ = strconv.Atoi(string(p[1]))
		}
		rows = append(rows, row)
	}
	if len(rows) == 0 {
		return info, nil, fmt.Errorf("meta: mtgtop8 event page holds no deck row")
	}
	return info, rows, nil
}

// ParseMTGTop8Deck reads the text export: "N Name" lines, then a
// "Sideboard" line and the sideboard. A Commander export lists the
// commander under the sideboard, as MTGO does.
func ParseMTGTop8Deck(text []byte) (main, side []Card) {
	inSide := false
	for _, line := range strings.Split(string(text), "\n") {
		line = strings.TrimSpace(strings.TrimSuffix(line, "\r"))
		if line == "" {
			continue
		}
		if strings.EqualFold(line, "sideboard") {
			inSide = true
			continue
		}
		count, name, ok := strings.Cut(line, " ")
		n, err := strconv.Atoi(count)
		if !ok || err != nil || n <= 0 {
			continue
		}
		name = plainText(name)
		if name == "" {
			continue
		}
		if inSide {
			side = append(side, Card{Name: name, Count: n})
		} else {
			main = append(main, Card{Name: name, Count: n})
		}
	}
	return main, side
}

// mtgtop8Tier labels a paper list (D-504, D-414): a top-8 finish in a
// field of 32 or more is great, and the rest is good. A page that names
// no field takes the size mark of the listing: two or three stars mark a
// field of that size.
func mtgtop8Tier(placement, players, stars int) string {
	if placement <= 0 || placement > 8 {
		return TierGood
	}
	if players >= 32 || (players == 0 && stars >= 2) {
		return TierGreat
	}
	return TierGood
}

// MTGTop8List builds the list of one deck of an event. ev is the listing
// row, empty on a re-parse, and info the event page's head. A Commander
// export carries the commanders as its sideboard, and the list moves
// them to their own field.
func MTGTop8List(code string, ev MTGTop8Event, info MTGTop8EventInfo, row MTGTop8Row, text []byte) (List, bool) {
	format, ok := mtgtop8Formats[code]
	if !ok {
		return List{}, false
	}
	main, side := ParseMTGTop8Deck(text)
	if len(main) == 0 {
		return List{}, false
	}
	name, date := info.Name, info.Date
	if name == "" {
		name = ev.Name
	}
	if date == "" {
		date = ev.Date
	}
	eventID := ev.ID
	if eventID == "" {
		eventID = "event"
	}
	l := List{
		Source: SourceMTGTop8, ID: eventID + "/" + row.DeckID, Format: format,
		Event: name, Date: date, Players: info.Players, Placement: row.Placement,
		Tier: mtgtop8Tier(row.Placement, info.Players, ev.Stars), Cards: main, Sideboard: side,
	}
	if format == FormatCommander {
		for _, c := range side {
			l.Commanders = append(l.Commanders, c.Name)
		}
		l.Sideboard = nil
	}
	return l, true
}

// plainText unescapes a page fragment, drops the bytes the site writes
// outside UTF-8, and trims it.
func plainText(s string) string {
	s = html.UnescapeString(s)
	s = strings.ToValidUTF8(s, "")
	return strings.TrimSpace(strings.Join(strings.Fields(s), " "))
}
