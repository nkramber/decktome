package meta

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GoldfishBase is the deck site. The user-deck listing of a format is
// server-rendered, thirty decks a page, and a deck page embeds the whole
// list in a form field (read 2026-09-03, pr14c-sources-2026-09-03). The
// robots file allows the deck page and disallows the download endpoint,
// so the reader takes the list from the page (D-503).
const GoldfishBase = "https://www.mtggoldfish.com"

// goldfishFormats maps the format word of a deck page to the model's
// word. The user decks of the 60-card formats are the typical rung
// (D-490).
var goldfishFormats = map[string]string{"Modern": FormatModern, "Standard": FormatStandard}

// GoldfishFormatWords lists the listing words the reader walks.
var GoldfishFormatWords = []string{"modern", "standard"}

// goldfishUserDeck marks a deck page of a user deck. A tournament deck
// page reads otherwise, and the reader leaves it out.
const goldfishUserDeck = "User Submitted Deck"

// GoldfishListingURL is one page of the user-deck listing of a format.
func GoldfishListingURL(base, word string, page int) string {
	if page <= 1 {
		return base + "/deck/custom/" + word
	}
	return base + "/deck/custom/" + word + "?page=" + strconv.Itoa(page)
}

// GoldfishDeckURL is the page of one deck.
func GoldfishDeckURL(base, id string) string { return base + "/deck/" + id }

var (
	goldfishTile   = regexp.MustCompile(`href="/deck/(\d+)#paper"`)
	goldfishTotal  = regexp.MustCompile(`Displaying <b>\d+ - (\d+)</b> of <b>(\d+)</b>`)
	goldfishInfo   = regexp.MustCompile(`(?s)<p class='deck-container-information'>(.*?)</p>`)
	goldfishFormat = regexp.MustCompile(`Format:\s*([A-Za-z ]+?)\s*<br>`)
	goldfishDate   = regexp.MustCompile(`Deck Date:\s*([A-Z][a-z]{2} \d{1,2}, \d{4})`)
	goldfishDeck   = regexp.MustCompile(`name="deck_input\[deck\]"[^>]*value="([^"]*)"`)
	goldfishName   = regexp.MustCompile(`name="deck_input\[name\]"[^>]*value="([^"]*)"`)
)

// ParseGoldfishListing reads the deck ids of a listing page, each once
// and in page order, and whether a later page exists.
func ParseGoldfishListing(page []byte) (ids []string, more bool) {
	seen := map[string]bool{}
	for _, m := range goldfishTile.FindAllSubmatch(page, -1) {
		id := string(m[1])
		if seen[id] {
			continue
		}
		seen[id] = true
		ids = append(ids, id)
	}
	if m := goldfishTotal.FindSubmatch(page); m != nil {
		last, _ := strconv.Atoi(string(m[1]))
		total, _ := strconv.Atoi(string(m[2]))
		more = last < total
	}
	return ids, more
}

// ParseGoldfishDeck reads a deck page as a typical list (D-490). A page
// of a tournament deck, or of a format the model does not cover, answers
// nil and no error. A page with no list is a parse failure (M-6).
func ParseGoldfishDeck(id string, page []byte) (*List, error) {
	info := goldfishInfo.FindSubmatch(page)
	if info == nil {
		return nil, fmt.Errorf("meta: mtggoldfish deck %s holds no information block", id)
	}
	if !strings.Contains(string(info[1]), goldfishUserDeck) {
		return nil, nil
	}
	fm := goldfishFormat.FindSubmatch(info[1])
	if fm == nil {
		return nil, fmt.Errorf("meta: mtggoldfish deck %s names no format", id)
	}
	format, ok := goldfishFormats[strings.TrimSpace(string(fm[1]))]
	if !ok {
		return nil, nil
	}
	dm := goldfishDate.FindSubmatch(info[1])
	if dm == nil {
		return nil, fmt.Errorf("meta: mtggoldfish deck %s names no date", id)
	}
	day, err := time.Parse("Jan 2, 2006", string(dm[1]))
	if err != nil {
		return nil, fmt.Errorf("meta: mtggoldfish deck %s: %w", id, err)
	}
	raw := goldfishDeck.FindSubmatch(page)
	if raw == nil {
		return nil, fmt.Errorf("meta: mtggoldfish deck %s holds no list", id)
	}
	main, side := ParseMTGTop8Deck([]byte(plainLines(string(raw[1]))))
	if len(main) == 0 {
		return nil, fmt.Errorf("meta: mtggoldfish deck %s holds an empty list", id)
	}
	name := "User deck " + id
	if nm := goldfishName.FindSubmatch(page); nm != nil {
		if s := plainText(string(nm[1])); s != "" {
			name = s
		}
	}
	return &List{
		Source: SourceGoldfish, ID: id, Format: format, Event: name,
		Date: day.Format("2006-01-02"), Tier: TierTypical, Cards: main, Sideboard: side,
	}, nil
}

// plainLines unescapes a form value and keeps its line breaks.
func plainLines(s string) string {
	var out []string
	for _, line := range strings.Split(s, "\n") {
		out = append(out, plainText(line))
	}
	return strings.Join(out, "\n")
}
