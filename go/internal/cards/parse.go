// Package cards holds the in-memory card database (roadmap PR-2).
//
// The data source is the daily Scryfall bulk snapshot. Guardrail 2: the
// legalities in this package are the only legality source in the app.
package cards

import (
	"encoding/json"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// rawCard mirrors the Scryfall card JSON fields this app reads.
type rawCard struct {
	ID            string            `json:"id"`
	OracleID      string            `json:"oracle_id"`
	Name          string            `json:"name"`
	ManaCost      string            `json:"mana_cost"`
	CMC           float64           `json:"cmc"`
	Colors        []string          `json:"colors"`
	ColorIdentity []string          `json:"color_identity"`
	TypeLine      string            `json:"type_line"`
	OracleText    string            `json:"oracle_text"`
	Keywords      []string          `json:"keywords"`
	Layout        string            `json:"layout"`
	CardFaces     []rawFace         `json:"card_faces"`
	Power         string            `json:"power"`
	Toughness     string            `json:"toughness"`
	Loyalty       string            `json:"loyalty"`
	ProducedMana  []string          `json:"produced_mana"`
	Legalities    map[string]string `json:"legalities"`
	GameChanger   bool              `json:"game_changer"`
	EdhrecRank    int32             `json:"edhrec_rank"`
	Set           string            `json:"set"`
	SetName       string            `json:"set_name"`
	CollectorNo   string            `json:"collector_number"`
	Rarity        string            `json:"rarity"`
	Artist        string            `json:"artist"`
	ImageUris     *rawImages        `json:"image_uris"`
	Digital       bool              `json:"digital"`
	Prices        map[string]string `json:"prices"`
}

type rawFace struct {
	OracleID   string     `json:"oracle_id"`
	Name       string     `json:"name"`
	Artist     string     `json:"artist"`
	ManaCost   string     `json:"mana_cost"`
	TypeLine   string     `json:"type_line"`
	OracleText string     `json:"oracle_text"`
	Power      string     `json:"power"`
	Toughness  string     `json:"toughness"`
	Loyalty    string     `json:"loyalty"`
	ImageUris  *rawImages `json:"image_uris"`
}

type rawImages struct {
	Small   string `json:"small"`
	Normal  string `json:"normal"`
	Large   string `json:"large"`
	ArtCrop string `json:"art_crop"`
}

func (r *rawImages) proto() *mtgv1.ImageUris {
	if r == nil {
		return nil
	}
	return &mtgv1.ImageUris{Small: r.Small, Normal: r.Normal, Large: r.Large, ArtCrop: r.ArtCrop}
}

var colorByLetter = map[string]mtgv1.Color{
	"W": mtgv1.Color_COLOR_W, "U": mtgv1.Color_COLOR_U, "B": mtgv1.Color_COLOR_B,
	"R": mtgv1.Color_COLOR_R, "G": mtgv1.Color_COLOR_G, "C": mtgv1.Color_COLOR_C,
}

func colors(letters []string) []mtgv1.Color {
	out := make([]mtgv1.Color, 0, len(letters))
	for _, l := range letters {
		if c, ok := colorByLetter[l]; ok {
			out = append(out, c)
		}
	}
	return out
}

var legalityByString = map[string]mtgv1.LegalityStatus{
	"legal":      mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL,
	"not_legal":  mtgv1.LegalityStatus_LEGALITY_STATUS_NOT_LEGAL,
	"banned":     mtgv1.LegalityStatus_LEGALITY_STATUS_BANNED,
	"restricted": mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED,
}

// parseCard converts one Scryfall Oracle-card JSON line into a proto Card.
// Face normalization (F-9): a single-faced card gets one face that carries
// the card-level image. A multi-faced card gets one face per printed face.
func parseCard(line []byte) (*mtgv1.Card, error) {
	var r rawCard
	if err := json.Unmarshal(line, &r); err != nil {
		return nil, err
	}
	// Reversible cards carry oracle_id on the faces, not the top level (C-8).
	if r.OracleID == "" && len(r.CardFaces) > 0 {
		r.OracleID = r.CardFaces[0].OracleID
	}
	c := &mtgv1.Card{
		OracleId:      r.OracleID,
		Name:          r.Name,
		ManaCost:      r.ManaCost,
		ManaValue:     r.CMC,
		Colors:        colors(r.Colors),
		ColorIdentity: colors(r.ColorIdentity),
		TypeLine:      r.TypeLine,
		OracleText:    r.OracleText,
		Keywords:      r.Keywords,
		Layout:        r.Layout,
		Power:         r.Power,
		Toughness:     r.Toughness,
		Loyalty:       r.Loyalty,
		ProducedMana:  colors(r.ProducedMana),
		GameChanger:   r.GameChanger,
		EdhrecRank:    r.EdhrecRank,
		Legalities:    buildLegalities(r.Legalities),
		DefaultPrinting: &mtgv1.Printing{
			ScryfallId:      r.ID,
			SetCode:         r.Set,
			SetName:         r.SetName,
			CollectorNumber: r.CollectorNo,
			Rarity:          r.Rarity,
			Artist:          r.Artist,
			ImageUris:       r.ImageUris.proto(),
			Digital:         r.Digital,
		},
	}
	if len(r.CardFaces) > 0 {
		for _, f := range r.CardFaces {
			face := &mtgv1.CardFace{
				Name: f.Name, ManaCost: f.ManaCost, TypeLine: f.TypeLine,
				OracleText: f.OracleText, Power: f.Power, Toughness: f.Toughness,
				Loyalty: f.Loyalty, ImageUris: f.ImageUris.proto(), Artist: f.Artist,
			}
			// Split and adventure faces share one printed image at card level.
			if face.ImageUris == nil {
				face.ImageUris = r.ImageUris.proto()
			}
			// Faces of one card can have different artists (D-6). A face
			// with no artist of its own takes the card-level one.
			if face.Artist == "" {
				face.Artist = r.Artist
			}
			c.Faces = append(c.Faces, face)
		}
		// A multi-face card's text often lives only on the faces.
		if c.OracleText == "" {
			texts := make([]string, 0, len(c.Faces))
			for _, f := range c.Faces {
				texts = append(texts, f.OracleText)
			}
			c.OracleText = strings.Join(texts, "\n//\n")
		}
	} else {
		c.Faces = []*mtgv1.CardFace{{
			Name: r.Name, ManaCost: r.ManaCost, TypeLine: r.TypeLine,
			OracleText: r.OracleText, Power: r.Power, Toughness: r.Toughness,
			Loyalty: r.Loyalty, ImageUris: r.ImageUris.proto(), Artist: r.Artist,
		}}
	}
	if usd := r.Prices["usd"]; usd != "" {
		if v, err := strconv.ParseFloat(usd, 64); err == nil {
			c.PriceUsd = v
		}
	}
	derive(c)
	return c, nil
}

func buildLegalities(in map[string]string) map[string]mtgv1.LegalityStatus {
	out := make(map[string]mtgv1.LegalityStatus, len(in))
	for k, v := range in {
		out[k] = legalityByString[v]
	}
	return out
}

// Printing is one physical printing row from default_cards.
type Printing struct {
	ScryfallID      string
	OracleID        string
	Name            string
	SetCode         string
	CollectorNumber string
	// Layout is the Scryfall layout. Non-playable layouts (tokens, emblems,
	// art cards) stay out of the index but are remembered by id, so an
	// import can report them instead of a silent name fallback.
	Layout string
	// Digital marks an online-only printing. The oracle_cards file names
	// one printing per card, and for some cards that printing is digital.
	// Diamond Valley reads as Masters Edition, and its paper printing is
	// Arabian Nights (D-221).
	Digital bool
	// ReleasedAt is the printing date, ISO 8601. The index prefers the
	// newest paper printing, because that is the one a player can buy.
	ReleasedAt string
	SetName    string
	Rarity     string
	Artist     string
	ImageUris  *mtgv1.ImageUris
	// PriceUSD is the printing's USD price, 0 when it has none. A digital
	// printing carries no USD price at all, only MTGO tickets, so a card
	// whose default printing is digital reads as free. The price must
	// follow the printing the index shows (D-231, extends D-221 and D-17).
	PriceUSD float64
}

// parsePrinting reads the minimal printing row for collection resolution.
func parsePrinting(line []byte) (Printing, error) {
	var r struct {
		ID              string            `json:"id"`
		OracleID        string            `json:"oracle_id"`
		Name            string            `json:"name"`
		Set             string            `json:"set"`
		SetName         string            `json:"set_name"`
		CollectorNumber string            `json:"collector_number"`
		Layout          string            `json:"layout"`
		Rarity          string            `json:"rarity"`
		Artist          string            `json:"artist"`
		Digital         bool              `json:"digital"`
		ReleasedAt      string            `json:"released_at"`
		Prices          map[string]string `json:"prices"`
		ImageUris       *rawImages        `json:"image_uris"`
		CardFaces       []struct {
			OracleID string `json:"oracle_id"`
		} `json:"card_faces"`
	}
	if err := json.Unmarshal(line, &r); err != nil {
		return Printing{}, err
	}
	// Reversible cards carry oracle_id on the faces, not the top level.
	if r.OracleID == "" && len(r.CardFaces) > 0 {
		r.OracleID = r.CardFaces[0].OracleID
	}
	return Printing{ScryfallID: r.ID, OracleID: r.OracleID, Name: r.Name,
		SetCode: r.Set, CollectorNumber: r.CollectorNumber, Layout: r.Layout,
		Digital: r.Digital, ReleasedAt: r.ReleasedAt, SetName: r.SetName,
		Rarity: r.Rarity, Artist: r.Artist, ImageUris: r.ImageUris.proto(),
		PriceUSD: usdPrice(r.Prices)}, nil
}

// usdPrice reads the USD price of a printing. A digital printing has
// none, and the field may be absent or null (D-231).
func usdPrice(prices map[string]string) float64 {
	v, err := strconv.ParseFloat(prices["usd"], 64)
	if err != nil {
		return 0
	}
	return v
}
