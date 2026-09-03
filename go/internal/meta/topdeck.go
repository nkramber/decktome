package meta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TopdeckBase is the tournaments API (D-417). Every request carries the
// key in the Authorization header, the limit is 100 requests a minute,
// and the terms ask for a visible credit (docs read 2026-09-02).
const TopdeckBase = "https://topdeck.gg/api"

// TopdeckCredit is the credit line the app shows, with TopdeckURL as
// its link (D-417, guardrail 7).
const (
	TopdeckCredit = "Tournament data by TopDeck.gg"
	TopdeckURL    = "https://topdeck.gg"
)

// TopdeckKeyEnv names the key in .env and in Secret Manager. The worker
// refuses its Topdeck job without one.
const TopdeckKeyEnv = "TOPDECK_API_KEY"

// TopdeckFormatEDH is the format word of the API for Commander. The
// words are case-sensitive.
const TopdeckFormatEDH = "EDH"

// TopdeckGame is the game word of the API.
const TopdeckGame = "Magic: The Gathering"

// TopdeckRequestsPerMinute is the documented limit of most endpoints.
// The bulk search has a lower one the docs do not number, and it
// answered 429 to weekly slices 600 ms apart on 2026-09-02, so the
// client spaces its calls by TopdeckBulkGap.
const TopdeckRequestsPerMinute = 100

// TopdeckBulkGap is the wait between two bulk searches.
const TopdeckBulkGap = 10 * time.Second

// TopdeckMaxBody bounds one answer of the bulk search. A two-week
// window of cEDH answered 43 MB on 2026-09-02, and the client asks for
// one week at a time.
const TopdeckMaxBody = 256 << 20

// TopdeckSliceDays is the window of one bulk call.
const TopdeckSliceDays = 7

// DefaultMinPlayers is the field a tournament needs before its lists
// count. On 2026-09-02 a two-week window held 617 cEDH events, 353 of
// them under 16 players, and those held a tenth of the structured
// decks. A pod of eight is a game night, not a tournament (D-479).
const DefaultMinPlayers = 16

// Topdeck calls the API.
type Topdeck struct {
	fetch *Fetcher
	key   string
	base  string
	// MinPlayers is the field a tournament needs. Zero means the
	// default.
	MinPlayers int
	// Now is the clock, for the tests.
	Now func() time.Time
}

// NewTopdeck makes a client. An empty key is an error, so no job starts
// without one.
func NewTopdeck(fetch *Fetcher, key, base string, logger *slog.Logger) (*Topdeck, error) {
	if key == "" {
		return nil, fmt.Errorf("meta: %s is empty, and the Topdeck job needs a key", TopdeckKeyEnv)
	}
	if base == "" {
		base = TopdeckBase
	}
	if fetch == nil {
		fetch = NewFetcher(nil, TopdeckBulkGap, logger)
	}
	fetch.MaxBody = TopdeckMaxBody
	return &Topdeck{fetch: fetch, key: key, base: base, Now: time.Now}, nil
}

// Tournament is one completed tournament with its standings.
type Tournament struct {
	TID  string `json:"TID"`
	Name string `json:"tournamentName"`
	// Start is unix seconds.
	Start     int64      `json:"startDate"`
	Players   int        `json:"-"`
	TopCut    int        `json:"topCut"`
	Standings []Standing `json:"standings"`
}

// Standing is one player's row. The deck comes as text, and as deckObj
// when the organizer collected a structured list.
type Standing struct {
	Standing int             `json:"standing"`
	Name     string          `json:"name"`
	ID       string          `json:"id"`
	Decklist string          `json:"decklist"`
	DeckObj  json.RawMessage `json:"deckObj"`
	Wins     int             `json:"wins"`
	Losses   int             `json:"losses"`
	Draws    int             `json:"draws"`
	WinRate  float64         `json:"winRate"`
}

// Tournaments searches the completed tournaments of a format over the
// last days, with the standings and the decklists. The window goes to
// the API one week at a time, oldest first, with the field floor, and
// a tournament dated after now is a test event and stays out.
func (t *Topdeck) Tournaments(ctx context.Context, format string, lastDays int) ([]Tournament, error) {
	now := t.Now().UTC()
	minPlayers := t.MinPlayers
	if minPlayers <= 0 {
		minPlayers = DefaultMinPlayers
	}
	var out []Tournament
	seen := map[string]bool{}
	start := now.AddDate(0, 0, -lastDays)
	for start.Before(now) {
		end := start.AddDate(0, 0, TopdeckSliceDays)
		if end.After(now) {
			end = now
		}
		body, err := json.Marshal(map[string]any{
			"game":           TopdeckGame,
			"format":         format,
			"start":          start.Unix(),
			"end":            end.Unix(),
			"participantMin": minPlayers,
			"columns":        []string{"name", "id", "decklist", "wins", "losses", "draws", "winRate"},
		})
		if err != nil {
			return nil, err
		}
		data, err := t.fetch.do(ctx, http.MethodPost, t.base+"/v2/tournaments", bytes.NewReader(body), map[string]string{
			"Authorization": t.key,
			"Content-Type":  "application/json",
		})
		if err != nil {
			return nil, err
		}
		slice, err := ParseTopdeckTournaments(data)
		if err != nil {
			return nil, err
		}
		for _, tr := range slice {
			if seen[tr.TID] || tr.Start > now.Unix() || tr.Players < minPlayers {
				continue
			}
			seen[tr.TID] = true
			out = append(out, tr)
		}
		start = end
	}
	return out, nil
}

// ParseTopdeckTournaments reads the answer of the bulk search.
func ParseTopdeckTournaments(data []byte) ([]Tournament, error) {
	var out []Tournament
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("meta: topdeck tournaments: %w", err)
	}
	for i := range out {
		out[i].Players = len(out[i].Standings)
		// The bulk answer carries no standing field, and its rows come
		// in rank order (read 2026-09-02), so the row order stands in.
		for j := range out[i].Standings {
			if out[i].Standings[j].Standing == 0 {
				out[i].Standings[j].Standing = j + 1
			}
		}
	}
	return out, nil
}

// Date is the start day, YYYY-MM-DD, in UTC.
func (t *Tournament) Date() string {
	return time.Unix(t.Start, 0).UTC().Format("2006-01-02")
}

// TopCutSize is the size of the top cut: the API's own when it names
// one, else the pod cut of cEDH, which is 4 up to 16 players and 16
// above 64.
func (t *Tournament) TopCutSize() int {
	if t.TopCut > 0 {
		return t.TopCut
	}
	switch {
	case t.Players > 64:
		return 16
	case t.Players > 16:
		return 8
	default:
		return 4
	}
}

// Lists reads the standings as lists. A row with no readable deck is
// left out. A top-cut finish is great, and the rest is good (D-414).
func (t *Tournament) Lists() []List {
	cut := t.TopCutSize()
	var out []List
	for _, s := range t.Standings {
		commanders, cards := ParseTopdeckDeck(s.DeckObj, s.Decklist)
		if len(cards) == 0 {
			continue
		}
		tier := TierGood
		if s.Standing > 0 && s.Standing <= cut {
			tier = TierGreat
		}
		out = append(out, List{
			Source: SourceTopdeck, ID: t.TID + "/" + s.ID, Format: FormatCommander,
			Event: t.Name, Date: t.Date(), Players: t.Players, Placement: s.Standing,
			Wins: s.Wins, Losses: s.Losses, Tier: tier, Commanders: commanders, Cards: cards,
		})
	}
	return out
}

// topdeckSection is a heading of the text deck, "~~Commanders~~".
var topdeckSection = regexp.MustCompile(`^~~\s*(.+?)\s*~~$`)

// ParseTopdeckDeck reads a deck from deckObj when it holds one, else
// from the text. deckObj maps a section word to a map of card name to a
// row, and the row is an object with a count, or the count itself. The
// text holds "~~Section~~" lines and "N Name" lines. A live answer of
// 2026-09-02 held the object as Commanders, Mainboard, and metadata,
// each card a row with an id of the site's own and a count.
func ParseTopdeckDeck(obj json.RawMessage, text string) (commanders []string, cards []Card) {
	if len(obj) > 0 && !bytes.Equal(bytes.TrimSpace(obj), []byte("null")) {
		commanders, cards = parseTopdeckObj(obj)
		if len(cards) > 0 {
			return commanders, cards
		}
	}
	// A row with no deck reads the word "None" in both fields, and a
	// list the organizer took as a link reads a URL (read 2026-09-02).
	text = strings.TrimSpace(text)
	if text == "" || text == "None" || strings.HasPrefix(text, "http") {
		return nil, nil
	}
	return parseTopdeckText(text)
}

func parseTopdeckObj(obj json.RawMessage) (commanders []string, cards []Card) {
	var sections map[string]map[string]json.RawMessage
	if err := json.Unmarshal(obj, &sections); err != nil {
		return nil, nil
	}
	names := make([]string, 0, len(sections))
	for s := range sections {
		names = append(names, s)
	}
	sort.Strings(names)
	for _, section := range names {
		rows := sections[section]
		cardNames := make([]string, 0, len(rows))
		for n := range rows {
			cardNames = append(cardNames, n)
		}
		sort.Strings(cardNames)
		for _, name := range cardNames {
			count := topdeckCount(rows[name])
			if count <= 0 {
				continue
			}
			if strings.EqualFold(section, "commanders") || strings.EqualFold(section, "commander") {
				commanders = append(commanders, name)
				continue
			}
			if strings.EqualFold(section, "sideboard") {
				continue
			}
			cards = append(cards, Card{Name: cleanName(name), Count: count})
		}
	}
	return commanders, cards
}

// topdeckCount reads a row as an object with a count, or as a number.
func topdeckCount(raw json.RawMessage) int {
	var n int
	if err := json.Unmarshal(raw, &n); err == nil {
		return n
	}
	var row struct {
		Count int `json:"count"`
		Qty   int `json:"quantity"`
	}
	if err := json.Unmarshal(raw, &row); err != nil {
		return 0
	}
	if row.Count > 0 {
		return row.Count
	}
	return row.Qty
}

func parseTopdeckText(text string) (commanders []string, cards []Card) {
	section := "mainboard"
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if m := topdeckSection.FindStringSubmatch(line); m != nil {
			section = strings.ToLower(m[1])
			continue
		}
		count, name := splitCountName(line)
		if name == "" {
			continue
		}
		switch {
		case strings.HasPrefix(section, "commander"):
			commanders = append(commanders, name)
		case strings.HasPrefix(section, "sideboard"):
		default:
			cards = append(cards, Card{Name: name, Count: count})
		}
	}
	return commanders, cards
}

// splitCountName reads "1 Sol Ring" and "1x Sol Ring", and a bare name
// as one copy.
func splitCountName(line string) (int, string) {
	fields := strings.SplitN(line, " ", 2)
	if len(fields) == 2 {
		n, err := strconv.Atoi(strings.TrimSuffix(fields[0], "x"))
		if err == nil && n > 0 {
			return n, cleanName(fields[1])
		}
	}
	return 1, cleanName(line)
}
