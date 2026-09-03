package meta

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// MTGJSONBase is the deck product API (precon-data-2026-09-01).
const MTGJSONBase = "https://mtgjson.com/api/v5"

// MTGJSONDeckListURL lists every deck product.
const MTGJSONDeckListURL = MTGJSONBase + "/DeckList.json"

// MTGJSONDeckURL is the file of one product.
func MTGJSONDeckURL(fileName string) string { return MTGJSONBase + "/decks/" + fileName + ".json" }

// DeckEntry is one row of the deck list.
type DeckEntry struct {
	Code        string `json:"code"`
	FileName    string `json:"fileName"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
	Type        string `json:"type"`
}

// ParseDeckList reads the deck list and its version.
func ParseDeckList(data []byte) (version string, entries []DeckEntry, err error) {
	var doc struct {
		Meta struct {
			Version string `json:"version"`
		} `json:"meta"`
		Data []DeckEntry `json:"data"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		return "", nil, fmt.Errorf("meta: mtgjson deck list: %w", err)
	}
	if doc.Meta.Version == "" {
		return "", nil, fmt.Errorf("meta: mtgjson deck list carries no version")
	}
	return doc.Meta.Version, doc.Data, nil
}

// preconTypes are the product types the table holds (D-407): the
// Commander decks and the 60-card constructed decks. Jumpstart packs,
// Welcome decks, Secret Lair drops, and land packs stay out. The word
// is the format the quality fit reads the product as.
var preconTypes = map[string]string{
	"Commander Deck":          FormatCommander,
	"MTGO Commander Deck":     FormatCommander,
	"Brawl Deck":              "",
	"Starter Deck":            FormatSixty,
	"Intro Pack":              FormatSixty,
	"Planeswalker Deck":       FormatSixty,
	"Challenger Deck":         FormatSixty,
	"Pioneer Challenger Deck": FormatSixty,
	"Event Deck":              FormatSixty,
	"Modern Event Deck":       FormatSixty,
	"Theme Deck":              FormatSixty,
}

// KeepPrecon says whether the table holds a product type, and the
// format word its list carries. A Brawl deck sits in the table for the
// ownership check of PR-24 and serves no fit, because Brawl is not a
// format the app builds.
func KeepPrecon(deckType string) (format string, ok bool) {
	format, ok = preconTypes[deckType]
	return format, ok
}

// PreconCard is one printing of a product with its count.
type PreconCard struct {
	Name       string `json:"name"`
	Count      int    `json:"count"`
	SetCode    string `json:"set_code"`
	Number     string `json:"number"`
	ScryfallID string `json:"scryfall_id,omitempty"`
	OracleID   string `json:"oracle_id,omitempty"`
}

// Precon is one deck product, the row of the precon table (D-407).
type Precon struct {
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	Type        string       `json:"type"`
	ReleaseDate string       `json:"release_date"`
	Commanders  []PreconCard `json:"commanders,omitempty"`
	Cards       []PreconCard `json:"cards"`
	Sideboard   []PreconCard `json:"sideboard,omitempty"`
}

type mtgjsonCard struct {
	Name        string `json:"name"`
	Count       int    `json:"count"`
	SetCode     string `json:"setCode"`
	Number      string `json:"number"`
	Identifiers struct {
		ScryfallID string `json:"scryfallId"`
		OracleID   string `json:"scryfallOracleId"`
	} `json:"identifiers"`
}

// ParsePrecon reads one deck file.
func ParsePrecon(data []byte) (*Precon, error) {
	var doc struct {
		Data struct {
			Name        string        `json:"name"`
			Code        string        `json:"code"`
			Type        string        `json:"type"`
			ReleaseDate string        `json:"releaseDate"`
			Commander   []mtgjsonCard `json:"commander"`
			MainBoard   []mtgjsonCard `json:"mainBoard"`
			SideBoard   []mtgjsonCard `json:"sideBoard"`
		} `json:"data"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("meta: mtgjson deck: %w", err)
	}
	d := doc.Data
	if d.Name == "" || len(d.MainBoard) == 0 {
		return nil, fmt.Errorf("meta: mtgjson deck %q holds no main board", d.Name)
	}
	return &Precon{
		Name: d.Name, Code: d.Code, Type: d.Type, ReleaseDate: d.ReleaseDate,
		Commanders: preconCards(d.Commander), Cards: preconCards(d.MainBoard), Sideboard: preconCards(d.SideBoard),
	}, nil
}

func preconCards(rows []mtgjsonCard) []PreconCard {
	out := make([]PreconCard, 0, len(rows))
	for _, r := range rows {
		if r.Count <= 0 || r.Name == "" {
			continue
		}
		out = append(out, PreconCard{
			Name: r.Name, Count: r.Count, SetCode: r.SetCode, Number: r.Number,
			ScryfallID: r.Identifiers.ScryfallID, OracleID: r.Identifiers.OracleID,
		})
	}
	return out
}

// Key is the product's id in the table: the file name of MTGJSON, which
// is the name and the set code.
func (p *Precon) Key() string {
	return strings.ReplaceAll(p.Name, " ", "") + "_" + p.Code
}

// List reads the product as a baseline list (D-414). A product of a
// type that serves no fit answers nil.
func (p *Precon) List() *List {
	format, ok := KeepPrecon(p.Type)
	if !ok || format == "" {
		return nil
	}
	l := &List{
		Source: SourceMTGJSON, ID: p.Key(), Format: format, Event: p.Name,
		Date: p.ReleaseDate, Tier: TierBaseline,
	}
	for _, c := range p.Commanders {
		l.Commanders = append(l.Commanders, c.Name)
	}
	l.Cards = mergeCards(p.Cards)
	l.Sideboard = mergeCards(p.Sideboard)
	return l
}

// mergeCards sums the printings of one name into one row, because a
// product can hold two printings of a basic land.
func mergeCards(rows []PreconCard) []Card {
	byName := map[string]*Card{}
	var order []string
	for _, r := range rows {
		c, ok := byName[r.Name]
		if !ok {
			c = &Card{Name: r.Name, OracleID: r.OracleID, ScryfallID: r.ScryfallID}
			byName[r.Name] = c
			order = append(order, r.Name)
		}
		c.Count += r.Count
	}
	out := make([]Card, 0, len(order))
	for _, n := range order {
		out = append(out, *byName[n])
	}
	return out
}

// SortPrecons orders a table by release date, then name, so a stored
// table reads the same on every run.
func SortPrecons(list []Precon) {
	sort.Slice(list, func(i, j int) bool {
		if list[i].ReleaseDate != list[j].ReleaseDate {
			return list[i].ReleaseDate < list[j].ReleaseDate
		}
		return list[i].Name < list[j].Name
	})
}
