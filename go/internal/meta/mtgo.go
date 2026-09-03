package meta

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// MTGOBase is the decklist site. A month page lists the events, and an
// event page embeds the lists in window.MTGO.decklists.data (read
// 2026-09-02, deck-quality-sources-2026-09-01).
const MTGOBase = "https://www.mtgo.com"

// MTGOMonthURL is the page of one month.
func MTGOMonthURL(year int, month time.Month) string {
	return fmt.Sprintf("%s/decklists/%04d/%02d", MTGOBase, year, int(month))
}

// MTGOEventURL is the page of one event slug.
func MTGOEventURL(slug string) string { return MTGOBase + "/decklist/" + slug }

var mtgoEventHref = regexp.MustCompile(`href="/decklist/([a-z0-9-]+)"`)

// ParseMTGOMonth reads the event slugs of a month page, in page order,
// each once.
func ParseMTGOMonth(page []byte) []string {
	seen := map[string]bool{}
	var out []string
	for _, m := range mtgoEventHref.FindAllSubmatch(page, -1) {
		slug := string(m[1])
		if seen[slug] {
			continue
		}
		seen[slug] = true
		out = append(out, slug)
	}
	return out
}

// mtgoSlug reads the format, the date, and the event id of a slug such
// as "modern-challenge-32-2026-09-0212853228". The date and the id run
// together at the end.
var mtgoSlug = regexp.MustCompile(`^([a-z]+(?:-[a-z]+)*?)-((?:challenge|league|showcase|qualifier|super-qualifier|preliminary|premier|last-chance)[a-z0-9-]*?)-?(\d{4}-\d{2}-\d{2})(\d+)$`)

// MTGOSlugFormat answers the format word of a slug, or "" for a format
// the model does not cover. Duel Commander is not Commander (D-112).
func MTGOSlugFormat(slug string) string {
	switch {
	case strings.HasPrefix(slug, "modern-"):
		return FormatModern
	case strings.HasPrefix(slug, "standard-"):
		return FormatStandard
	default:
		return ""
	}
}

// mtgoEvent is the embedded object of an event page. A challenge carries
// the standings, and a league carries the wins of each list.
type mtgoEvent struct {
	Description string `json:"description"`
	Format      string `json:"format"`
	Type        string `json:"type"`
	PublishDate string `json:"publish_date"`
	PlayerCount struct {
		Players string `json:"players"`
	} `json:"player_count"`
	Decklists []struct {
		LoginID   string     `json:"loginid"`
		Player    string     `json:"player"`
		MainDeck  []mtgoCard `json:"main_deck"`
		Sideboard []mtgoCard `json:"sideboard_deck"`
		Wins      *struct {
			Wins   string `json:"wins"`
			Losses string `json:"losses"`
		} `json:"wins"`
	} `json:"decklists"`
	Standings []struct {
		LoginID string `json:"loginid"`
		Rank    string `json:"rank"`
	} `json:"standings"`
	FinalRank []struct {
		LoginID string `json:"loginid"`
		Rank    string `json:"rank"`
	} `json:"final_rank"`
	WinLoss []struct {
		LoginID string `json:"loginid"`
		Wins    string `json:"wins"`
		Losses  string `json:"losses"`
	} `json:"winloss"`
}

type mtgoCard struct {
	Qty        string `json:"qty"`
	Attributes struct {
		Name string `json:"card_name"`
	} `json:"card_attributes"`
}

var mtgoData = regexp.MustCompile(`window\.MTGO\.decklists\.data\s*=\s*`)

// ParseMTGOEvent reads every list of an event page. A page with no
// embedded object is an error, because that is the markup change M-6
// counts. A slug of a format the model does not cover answers no list
// and no error.
func ParseMTGOEvent(slug string, page []byte) ([]List, error) {
	format := MTGOSlugFormat(slug)
	if format == "" {
		return nil, nil
	}
	m := mtgoSlug.FindStringSubmatch(slug)
	if m == nil {
		return nil, fmt.Errorf("meta: mtgo slug %q has no date", slug)
	}
	date := m[3]
	loc := mtgoData.FindIndex(page)
	if loc == nil {
		return nil, fmt.Errorf("meta: mtgo page %s holds no decklist object", slug)
	}
	dec := json.NewDecoder(strings.NewReader(string(page[loc[1]:])))
	var ev mtgoEvent
	if err := dec.Decode(&ev); err != nil {
		return nil, fmt.Errorf("meta: mtgo page %s: %w", slug, err)
	}
	league := strings.Contains(slug, "-league-")
	players, _ := strconv.Atoi(ev.PlayerCount.Players)
	rank := map[string]int{}
	for _, s := range ev.FinalRank {
		if r, err := strconv.Atoi(s.Rank); err == nil {
			rank[s.LoginID] = r
		}
	}
	for _, s := range ev.Standings {
		if _, ok := rank[s.LoginID]; ok {
			continue
		}
		if r, err := strconv.Atoi(s.Rank); err == nil {
			rank[s.LoginID] = r
		}
	}
	record := map[string][2]int{}
	for _, w := range ev.WinLoss {
		wins, _ := strconv.Atoi(w.Wins)
		losses, _ := strconv.Atoi(w.Losses)
		record[w.LoginID] = [2]int{wins, losses}
	}
	event := ev.Description
	if event == "" {
		event = slug
	}
	out := make([]List, 0, len(ev.Decklists))
	for _, d := range ev.Decklists {
		l := List{
			Source: SourceMTGO, ID: slug + "/" + d.LoginID, Format: format,
			Event: event, Date: date, Players: players,
			Cards: mtgoCards(d.MainDeck), Sideboard: mtgoCards(d.Sideboard),
		}
		if d.Wins != nil {
			l.Wins, _ = strconv.Atoi(d.Wins.Wins)
			l.Losses, _ = strconv.Atoi(d.Wins.Losses)
		} else if r, ok := record[d.LoginID]; ok {
			l.Wins, l.Losses = r[0], r[1]
		}
		l.Placement = rank[d.LoginID]
		l.Tier = mtgoTier(league, l.Placement)
		if len(l.Cards) == 0 {
			continue
		}
		out = append(out, l)
	}
	return out, nil
}

// mtgoTier labels a list (D-414): a top-8 finish is great, and a league
// finish or the rest of a challenge is good. A league page shows the
// 5-0 lists alone, so every list there is a finish.
func mtgoTier(league bool, placement int) string {
	if !league && placement > 0 && placement <= 8 {
		return TierGreat
	}
	return TierGood
}

func mtgoCards(rows []mtgoCard) []Card {
	out := make([]Card, 0, len(rows))
	for _, r := range rows {
		n, _ := strconv.Atoi(r.Qty)
		name := cleanName(r.Attributes.Name)
		if n <= 0 || name == "" {
			continue
		}
		out = append(out, Card{Name: name, Count: n})
	}
	return out
}
