// Package archidekt reads one public Archidekt deck by the URL that a user
// pasted, and writes its list as text that the deck import reads (PR-123,
// D-1100).
//
// The read API needs no login. Archidekt staff call it open for reads, and
// they ask for a link back when the data shows in public (D-1101). The
// client sends the named agent of every other source, and it never fetches
// the URL of the user: it makes the API URL again from the deck id alone.
// Package decklink sorts a link by its site before this package reads it.
package archidekt

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

// UserAgent identifies this app to Archidekt. It is the agent of every
// other source (meta.UserAgent).
const UserAgent = "mtg-deck-builder/0.1 (github.com/nkramber/decktome)"

// DefaultBaseURL is the site and its API.
const DefaultBaseURL = "https://archidekt.com"

// Timeout bounds one read. A deck of 100 cards answers in under a second.
const Timeout = 10 * time.Second

// maxBodyBytes caps one answer. A deck of 100 cards is about 250 KB.
const maxBodyBytes = 4 << 20

// maxURLBytes caps the URL that a user pastes.
const maxURLBytes = 2048

var (
	// ErrNotDeckURL refuses a URL that names no Archidekt deck.
	ErrNotDeckURL = errors.New("url: paste the link of an Archidekt deck, such as https://archidekt.com/decks/123456")
	// ErrNotFound says Archidekt has no public deck at the id. A private
	// deck reads the same.
	ErrNotFound = errors.New("url: Archidekt has no public deck at that link. Make the deck public or unlisted on Archidekt, or export it as text")
	// ErrUnavailable says Archidekt did not answer with a deck.
	ErrUnavailable = errors.New("url: Archidekt did not answer. Try again later, or export the deck as text")
)

// ParseURL reads the deck id out of a deck URL. It takes the forms that
// a user copies: with or without the scheme, with or without "www.", and
// with the name of the deck after the id.
func ParseURL(raw string) (int64, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" || len(raw) > maxURLBytes {
		return 0, ErrNotDeckURL
	}
	if !strings.Contains(raw, "://") {
		raw = "https://" + raw
	}
	u, err := url.Parse(raw)
	if err != nil || (u.Scheme != "https" && u.Scheme != "http") || u.Port() != "" || u.User != nil {
		return 0, ErrNotDeckURL
	}
	host := strings.TrimPrefix(strings.ToLower(u.Hostname()), "www.")
	if host != "archidekt.com" {
		return 0, ErrNotDeckURL
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 || parts[0] != "decks" {
		return 0, ErrNotDeckURL
	}
	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil || id <= 0 || strconv.FormatInt(id, 10) != parts[1] {
		return 0, ErrNotDeckURL
	}
	return id, nil
}

// DeckURL is the page of a deck. The import keeps this link and never the
// URL that a user pasted (D-1101).
func DeckURL(id int64) string {
	return DefaultBaseURL + "/decks/" + strconv.FormatInt(id, 10)
}

// Deck is one deck that the client read.
type Deck struct {
	Name string
	// Text is the list in Arena sections, which decklist.Parse reads.
	Text string
	// LeftOut counts the copies of a category that Archidekt keeps out
	// of the deck, such as the maybeboard.
	LeftOut int
}

// Client reads decks from the API.
type Client struct {
	http    *http.Client
	baseURL string
}

// New returns a client with the timeout of one read. It follows no
// redirect, so an answer can not send the read to another host.
func New() *Client {
	return &Client{
		http: &http.Client{
			Timeout: Timeout,
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		baseURL: DefaultBaseURL,
	}
}

// WithBaseURL points the client at a test server.
func (c *Client) WithBaseURL(u string) *Client {
	c.baseURL = strings.TrimRight(u, "/")
	return c
}

// Fetch reads one deck by its id.
func (c *Client) Fetch(ctx context.Context, id int64) (*Deck, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/api/decks/"+strconv.FormatInt(id, 10)+"/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnavailable, err)
	}
	defer res.Body.Close() //nolint:errcheck // a read body has nothing to report
	switch {
	case res.StatusCode == http.StatusNotFound, res.StatusCode == http.StatusForbidden, res.StatusCode == http.StatusUnauthorized:
		return nil, ErrNotFound
	case res.StatusCode != http.StatusOK:
		return nil, fmt.Errorf("%w: status %d", ErrUnavailable, res.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(res.Body, maxBodyBytes+1))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnavailable, err)
	}
	if len(body) > maxBodyBytes {
		return nil, fmt.Errorf("%w: the answer is over %d bytes", ErrUnavailable, maxBodyBytes)
	}
	var d apiDeck
	if err := json.Unmarshal(body, &d); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnavailable, err)
	}
	return d.list(), nil
}

// apiDeck holds the fields of the answer that the list reads.
type apiDeck struct {
	Name       string        `json:"name"`
	Categories []apiCategory `json:"categories"`
	Cards      []apiCard     `json:"cards"`
}

type apiCategory struct {
	Name           string `json:"name"`
	IsPremier      bool   `json:"isPremier"`
	IncludedInDeck bool   `json:"includedInDeck"`
}

type apiCard struct {
	Quantity   int      `json:"quantity"`
	Categories []string `json:"categories"`
	Companion  bool     `json:"companion"`
	Modifier   string   `json:"modifier"`
	DeletedAt  *string  `json:"deletedAt"`
	Card       struct {
		CollectorNumber string `json:"collectorNumber"`
		Edition         struct {
			Code string `json:"editioncode"`
		} `json:"edition"`
		OracleCard struct {
			Name string `json:"name"`
		} `json:"oracleCard"`
	} `json:"card"`
}

type section int

const (
	commander section = iota
	companion
	mainDeck
	sideboard
	leftOut
)

// headers are the Arena headers of the sections, in the order of the text.
var headers = []string{commander: "Commander", companion: "Companion", mainDeck: "Deck", sideboard: "Sideboard"}

// sectionOf places one card. The premier category is the commander
// (D-847). The first category of a card decides the rest, as Archidekt
// counts it: "Sideboard" is the sideboard, and a category kept out of
// the deck, such as "Maybeboard", is left out (D-1100).
func (d *apiDeck) sectionOf(c apiCard) section {
	premier, included := map[string]bool{}, map[string]bool{}
	for _, cat := range d.Categories {
		premier[cat.Name], included[cat.Name] = cat.IsPremier, cat.IncludedInDeck
	}
	for _, name := range c.Categories {
		if premier[name] {
			return commander
		}
	}
	if c.Companion {
		return companion
	}
	if len(c.Categories) == 0 {
		return mainDeck
	}
	first := c.Categories[0]
	if strings.EqualFold(first, "sideboard") {
		return sideboard
	}
	if in, known := included[first]; known && !in {
		return leftOut
	}
	return mainDeck
}

// list writes the deck as Arena sections, one line a card, sorted by name
// within a section. A line names its set and number, so the import reads
// the same printing.
func (d *apiDeck) list() *Deck {
	out := &Deck{Name: oneLine(d.Name)}
	lines := make([][]string, len(headers))
	for _, c := range d.Cards {
		if c.DeletedAt != nil || c.Quantity <= 0 {
			continue
		}
		sec := d.sectionOf(c)
		if sec == leftOut {
			out.LeftOut += c.Quantity
			continue
		}
		lines[sec] = append(lines[sec], line(c))
	}
	var b strings.Builder
	for sec, ls := range lines {
		if len(ls) == 0 {
			continue
		}
		slices.SortFunc(ls, func(x, y string) int { return cmp.Compare(nameOf(x), nameOf(y)) })
		if b.Len() > 0 {
			b.WriteString("\n")
		}
		b.WriteString(headers[sec] + "\n")
		for _, l := range ls {
			b.WriteString(l + "\n")
		}
	}
	out.Text = b.String()
	return out
}

// line writes one card: "2 Name (set) 123", with "*F*" for a foil and
// "*E*" for an etched card.
func line(c apiCard) string {
	s := strconv.Itoa(c.Quantity) + " " + oneLine(c.Card.OracleCard.Name)
	set, num := oneLine(c.Card.Edition.Code), oneLine(c.Card.CollectorNumber)
	if set != "" && num != "" && !strings.ContainsAny(set+num, " ()") {
		s += " (" + set + ") " + num
	}
	switch c.Modifier {
	case "Foil":
		s += " *F*"
	case "Etched":
		s += " *E*"
	}
	return s
}

// nameOf is the part of a line after its count, for the sort.
func nameOf(l string) string {
	_, rest, _ := strings.Cut(l, " ")
	return rest
}

// oneLine keeps a field on one line of the list.
func oneLine(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
