package meta

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// EDHRECBase serves the JSON behind each EDHREC page. A commander page
// carries the deck count, the rank, and the bracket counts, and the
// average deck page carries one list (read 2026-09-02).
const EDHRECBase = "https://json.edhrec.com/pages"

// EDHRECCommanderURL is the commander page of a slug.
func EDHRECCommanderURL(slug string) string { return EDHRECBase + "/commanders/" + slug + ".json" }

// EDHRECAverageDeckURL is the average deck page of a slug.
func EDHRECAverageDeckURL(slug string) string { return EDHRECBase + "/average-decks/" + slug + ".json" }

// EDHRECTopURL lists the 100 most built commanders of a window: "year",
// "month", or "week".
func EDHRECTopURL(window string) string { return EDHRECBase + "/commanders/" + window + ".json" }

// EDHRECSlug turns a card name into the page slug: lower case, ASCII
// letters and digits kept, an accent dropped from its letter, the rest
// cut, and the words joined by hyphens. "Y'shtola, Night's Blessed" is
// "yshtola-nights-blessed", and "Bartolomé del Presidio" is
// "bartolome-del-presidio" (read 2026-09-02, the accented slug answered
// 403). A partner pair joins the two slugs with a hyphen.
func EDHRECSlug(names ...string) string {
	parts := make([]string, 0, len(names))
	for _, name := range names {
		if i := strings.Index(name, " // "); i > 0 {
			name = name[:i]
		}
		var b strings.Builder
		space := false
		for _, r := range strings.ToLower(foldAccents(name)) {
			switch {
			case r < 128 && (unicode.IsLetter(r) || unicode.IsDigit(r)):
				if space && b.Len() > 0 {
					b.WriteByte('-')
				}
				space = false
				b.WriteRune(r)
			case r == ' ' || r == '-':
				space = true
			}
		}
		if b.Len() > 0 {
			parts = append(parts, b.String())
		}
	}
	return strings.Join(parts, "-")
}

// Commander is the EDHREC read of one commander.
type Commander struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	// NumDecks counts the decks EDHREC holds for the commander.
	NumDecks int `json:"num_decks"`
	// Rank is the popularity rank, 1 for the most built.
	Rank int `json:"rank,omitempty"`
	// BracketCounts counts the decks per bracket their owners set, keyed
	// 1 to 5. Not every deck carries a bracket, so the sum is below
	// NumDecks.
	BracketCounts map[int]int `json:"bracket_counts,omitempty"`
	// ReadAt is the date of the read, YYYY-MM-DD.
	ReadAt string `json:"read_at,omitempty"`
	// Competitive marks a commander of the competitive section of the
	// cEDH Decklist Database.
	Competitive bool `json:"competitive,omitempty"`
	// TopCuts counts the top-cut finishes of the commander in the
	// Topdeck.gg cEDH tournaments of the window, and Entries the lists.
	TopCuts int `json:"top_cuts,omitempty"`
	Entries int `json:"entries,omitempty"`
}

// The shrink of the cEDH signal (D-485). A commander with one entry
// and one top cut is not a 1.00: the share starts at the pod prior, a
// quarter, with the weight of SignalPrior entries, so ten top cuts in
// ten entries read 0.75 and one in one reads 0.38.
const (
	SignalPrior      = 5.0
	SignalPriorShare = 0.25
)

// CEDHSignal is the bracket 5 power signal of the roadmap, in [0, 1]:
// the shrunk top-cut share of the commander's tournament entries, or
// the competitive tier of the database when no tournament list exists.
func (c *Commander) CEDHSignal() float64 {
	if c.Entries > 0 {
		share := (float64(c.TopCuts) + SignalPrior*SignalPriorShare) / (float64(c.Entries) + SignalPrior)
		if c.Competitive && share < 0.5 {
			share = 0.5
		}
		return share
	}
	if c.Competitive {
		return 0.5
	}
	return 0
}

// HighBracketShare is the share of bracketed decks at bracket 4 or 5.
// It is the lower-bracket power signal of the roadmap, and 0 for a
// commander with no bracketed deck.
func (c *Commander) HighBracketShare() float64 {
	total := 0
	for _, n := range c.BracketCounts {
		total += n
	}
	if total == 0 {
		return 0
	}
	return float64(c.BracketCounts[4]+c.BracketCounts[5]) / float64(total)
}

type edhrecPage struct {
	BracketCounts map[string]int `json:"bracket_counts"`
	Container     struct {
		JSONDict struct {
			Card struct {
				Name     string `json:"name"`
				NumDecks int    `json:"num_decks"`
				Rank     int    `json:"rank"`
			} `json:"card"`
			CardLists []struct {
				Header    string `json:"header"`
				CardViews []struct {
					Name     string `json:"name"`
					Slug     string `json:"slug"`
					NumDecks int    `json:"num_decks"`
					Rank     int    `json:"rank"`
				} `json:"cardviews"`
			} `json:"cardlists"`
		} `json:"json_dict"`
	} `json:"container"`
	Deck struct {
		Commander []string                   `json:"commander"`
		RawCards  map[string]json.RawMessage `json:"cards"`
	} `json:"deck"`
}

// ParseEDHRECCommander reads a commander page.
func ParseEDHRECCommander(slug string, data []byte) (*Commander, error) {
	var page edhrecPage
	if err := json.Unmarshal(data, &page); err != nil {
		return nil, fmt.Errorf("meta: edhrec page %s: %w", slug, err)
	}
	card := page.Container.JSONDict.Card
	if card.Name == "" {
		return nil, fmt.Errorf("meta: edhrec page %s names no card", slug)
	}
	c := &Commander{Name: card.Name, Slug: slug, NumDecks: card.NumDecks, Rank: card.Rank}
	for k, n := range page.BracketCounts {
		b, err := strconv.Atoi(k)
		if err != nil || b < 1 || b > 5 {
			continue
		}
		if c.BracketCounts == nil {
			c.BracketCounts = map[int]int{}
		}
		c.BracketCounts[b] = n
	}
	return c, nil
}

// ParseEDHRECAverageDeck reads the average deck of a commander as a
// typical list (D-414). The deck block maps a type word to rows of name
// and count.
func ParseEDHRECAverageDeck(slug string, data []byte) (*List, error) {
	var page edhrecPage
	if err := json.Unmarshal(data, &page); err != nil {
		return nil, fmt.Errorf("meta: edhrec average deck %s: %w", slug, err)
	}
	if len(page.Deck.Commander) == 0 || len(page.Deck.RawCards) == 0 {
		return nil, fmt.Errorf("meta: edhrec average deck %s holds no deck", slug)
	}
	l := &List{
		Source: SourceEDHREC, ID: slug, Format: FormatCommander, Event: "EDHREC average deck",
		Tier: TierTypical, Commanders: page.Deck.Commander,
	}
	types := make([]string, 0, len(page.Deck.RawCards))
	for t := range page.Deck.RawCards {
		types = append(types, t)
	}
	sort.Strings(types)
	for _, t := range types {
		var rows [][2]json.RawMessage
		if err := json.Unmarshal(page.Deck.RawCards[t], &rows); err != nil {
			return nil, fmt.Errorf("meta: edhrec average deck %s, %s rows: %w", slug, t, err)
		}
		for _, r := range rows {
			var name string
			var count int
			if err := json.Unmarshal(r[0], &name); err != nil {
				return nil, fmt.Errorf("meta: edhrec average deck %s, %s name: %w", slug, t, err)
			}
			if err := json.Unmarshal(r[1], &count); err != nil {
				return nil, fmt.Errorf("meta: edhrec average deck %s, %s count: %w", slug, t, err)
			}
			if count <= 0 || name == "" {
				continue
			}
			l.Cards = append(l.Cards, Card{Name: name, Count: count})
		}
	}
	return l, nil
}

// ParseEDHRECTop reads a top-commanders page: the name, the slug, the
// deck count, and the rank of each.
func ParseEDHRECTop(data []byte) ([]Commander, error) {
	var page edhrecPage
	if err := json.Unmarshal(data, &page); err != nil {
		return nil, fmt.Errorf("meta: edhrec top page: %w", err)
	}
	var out []Commander
	for _, cl := range page.Container.JSONDict.CardLists {
		for _, v := range cl.CardViews {
			if v.Name == "" || v.Slug == "" {
				continue
			}
			out = append(out, Commander{Name: v.Name, Slug: v.Slug, NumDecks: v.NumDecks, Rank: v.Rank})
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("meta: edhrec top page names no commander")
	}
	return out, nil
}

// foldAccents drops the combining marks of a name: "é" becomes "e".
func foldAccents(name string) string {
	var b strings.Builder
	for _, r := range norm.NFD.String(name) {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
