package cards

import (
	"slices"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// Index is the immutable in-memory card database. Build it once per
// snapshot and swap the pointer. Readers never see a partial state.
type Index struct {
	cards      []*mtgv1.Card
	byOracleID map[string]*mtgv1.Card
	byName     map[string]*mtgv1.Card // key: normalized full or face name
	byPrinting map[string]*mtgv1.Card // key: scryfall printing id
	// printings keeps the display fields of every playable printing, so
	// the deck view can show the printing the user owns (D-299).
	printings map[string]*mtgv1.Printing
	bySetNo   map[string]*mtgv1.Card // key: "setcode/collectornumber", lowercase
	// nonPlayable maps a dropped printing (Scryfall id and set/collector
	// key) to its layout, so an import can name the reason.
	nonPlayable map[string]string
	// paperSwaps counts the cards whose digital default printing was
	// replaced by a paper one (D-221).
	paperSwaps int
	tags       *TagIndex
	collisions Collisions
	// AsOf is the Scryfall updated_at of the snapshot (roadmap PR-3).
	AsOf time.Time
}

// Collisions counts name keys that more than one card claimed at build
// time (C-16). On a full-name tie the card that is legal in at least one
// format wins. A name that two equally playable cards share resolves to
// nothing, full name or face name: the lookup is exact, and a guess is
// a wrong card. A face name never overrides a full name. The log line
// shows the counts.
type Collisions struct {
	// FullNames counts a full name that a later card also carried.
	FullNames int
	// FaceNames counts a face name that another card's full or face
	// name already held.
	FaceNames int
}

// legalSomewhere reports whether a card is legal or restricted in at
// least one format. A playtest card or a front card is legal nowhere.
func legalSomewhere(c *mtgv1.Card) bool {
	for _, s := range c.Legalities {
		if s == mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL || s == mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED {
			return true
		}
	}
	return false
}

// Collisions returns the name collision counts of the build.
func (x *Index) Collisions() Collisions { return x.collisions }

// normName is the lookup key: lowercase, trimmed. Exact otherwise
// (guardrail 4: no fuzzy match).
func normName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

// NewIndex builds the index. printings and tags are optional. The card
// list is copied and sorted by EDHREC rank once, so Search never sorts.
// A card with a price gets price_as_of set to the snapshot date.
func NewIndex(cardList []*mtgv1.Card, printings []Printing, tags *TagIndex, asOf time.Time) *Index {
	sorted := slices.Clone(cardList)
	slices.SortStableFunc(sorted, func(a, b *mtgv1.Card) int {
		return rankOf(a) - rankOf(b)
	})
	priceDate := asOf.UTC().Format("2006-01-02")
	idx := &Index{
		cards:       sorted,
		byOracleID:  make(map[string]*mtgv1.Card, len(cardList)),
		byName:      make(map[string]*mtgv1.Card, len(cardList)*2),
		byPrinting:  make(map[string]*mtgv1.Card, len(printings)),
		printings:   make(map[string]*mtgv1.Printing, len(printings)),
		bySetNo:     make(map[string]*mtgv1.Card, len(printings)),
		nonPlayable: make(map[string]string),
		tags:        tags,
		AsOf:        asOf,
	}
	// fullAmbiguous holds a full name that two equally playable cards
	// share. It stays in byName through the face walk, so no face name
	// takes it, and leaves before the index is returned.
	fullAmbiguous := map[string]bool{}
	for _, c := range cardList {
		if c.PriceUsd > 0 && c.PriceAsOf == "" {
			c.PriceAsOf = priceDate
		}
		idx.byOracleID[c.OracleId] = c
		k := normName(c.Name)
		if taken, ok := idx.byName[k]; ok {
			idx.collisions.FullNames++
			// A playable card beats one that is legal nowhere. Two cards
			// of equal standing make the name ambiguous.
			switch {
			case !legalSomewhere(taken) && legalSomewhere(c):
				idx.byName[k] = c
				delete(fullAmbiguous, k)
			case legalSomewhere(taken) == legalSomewhere(c):
				fullAmbiguous[k] = true
			}
		} else {
			idx.byName[k] = c
		}
		if c.DefaultPrinting != nil {
			idx.byPrinting[c.DefaultPrinting.ScryfallId] = c
		}
	}
	// Face names resolve to the whole card ("Stomp" finds "Bonecrusher
	// Giant // Stomp"). A face name never overrides a real full name. A
	// face name that two cards share is ambiguous: "Fire" is a face of
	// both "Fire // Ice" and "Start // Fire", so it resolves to nothing.
	faceOwner := map[string]*mtgv1.Card{}
	ambiguous := map[string]bool{}
	for _, c := range cardList {
		for _, f := range c.Faces {
			k := normName(f.Name)
			if taken, ok := idx.byName[k]; ok {
				if taken != c {
					idx.collisions.FaceNames++
				}
				continue
			}
			if owner, ok := faceOwner[k]; ok {
				if owner != c {
					idx.collisions.FaceNames++
					ambiguous[k] = true
				}
				continue
			}
			faceOwner[k] = c
		}
	}
	for k, c := range faceOwner {
		if !ambiguous[k] {
			idx.byName[k] = c
		}
	}
	for k := range fullAmbiguous {
		delete(idx.byName, k)
	}
	// paper collects a replacement for every card whose default printing
	// is digital and whose paper printing the file also holds (D-221).
	paper := map[string]Printing{}
	for _, p := range printings {
		c, ok := idx.byOracleID[p.OracleID]
		if !ok {
			if SkipLayouts[p.Layout] {
				idx.nonPlayable[p.ScryfallID] = p.Layout
				if p.SetCode != "" && p.CollectorNumber != "" {
					idx.nonPlayable[setNoKey(p.SetCode, p.CollectorNumber)] = p.Layout
				}
			}
			continue
		}
		idx.byPrinting[p.ScryfallID] = c
		idx.printings[p.ScryfallID] = &mtgv1.Printing{
			ScryfallId:      p.ScryfallID,
			SetCode:         p.SetCode,
			SetName:         p.SetName,
			CollectorNumber: p.CollectorNumber,
			Rarity:          p.Rarity,
			Artist:          p.Artist,
			ImageUris:       p.ImageUris,
			Digital:         p.Digital,
			PriceUsd:        p.PriceUSD,
		}
		if p.SetCode != "" && p.CollectorNumber != "" {
			idx.bySetNo[setNoKey(p.SetCode, p.CollectorNumber)] = c
		}
		// A digital default printing is replaced by a paper one: the
		// price (D-17), the image, and the artist credit (D-6) follow the
		// printing (D-221).
		if !p.Digital && c.DefaultPrinting.GetDigital() {
			paper[c.OracleId] = newerPaper(paper[c.OracleId], p)
		}
	}
	// The swap runs after the walk, so the chosen paper printing wins
	// whatever order the file holds.
	for oid, p := range paper {
		c, ok := idx.byOracleID[oid]
		if !ok {
			continue
		}
		c.DefaultPrinting = &mtgv1.Printing{
			ScryfallId:      p.ScryfallID,
			SetCode:         p.SetCode,
			SetName:         p.SetName,
			CollectorNumber: p.CollectorNumber,
			Rarity:          p.Rarity,
			Artist:          p.Artist,
			ImageUris:       p.ImageUris,
		}
		// The price follows the printing (D-231).
		if p.PriceUSD > 0 {
			c.PriceUsd = p.PriceUSD
			c.PriceAsOf = priceDate
		}
		idx.paperSwaps++
	}
	return idx
}

// newerPaper keeps the later of two paper printings, and a printing with
// a USD price beats one without. A player buys the newest one, and a
// printing with no price would make the card free to the budget (D-231).
func newerPaper(have, next Printing) Printing {
	if have.ScryfallID == "" {
		return next
	}
	if (next.PriceUSD > 0) != (have.PriceUSD > 0) {
		if next.PriceUSD > 0 {
			return next
		}
		return have
	}
	if next.ReleasedAt > have.ReleasedAt {
		return next
	}
	return have
}

// NonPlayablePrinting reports a printing the index dropped on purpose
// (token, emblem, art card). key is a Scryfall id or a set/collector key.
func (x *Index) NonPlayablePrinting(scryfallID, setCode, collector string) (layout string, ok bool) {
	if scryfallID != "" {
		if l, ok := x.nonPlayable[scryfallID]; ok {
			return l, true
		}
	}
	if setCode != "" && collector != "" {
		if l, ok := x.nonPlayable[setNoKey(setCode, collector)]; ok {
			return l, true
		}
	}
	return "", false
}

// ByName finds a card by exact full name or exact face name. A face
// name that two cards share finds nothing.
func (x *Index) ByName(name string) (*mtgv1.Card, bool) {
	c, ok := x.byName[normName(name)]
	return c, ok
}

// ByOracleID finds a card by Oracle id.
func (x *Index) ByOracleID(id string) (*mtgv1.Card, bool) {
	c, ok := x.byOracleID[id]
	return c, ok
}

// Printing returns the display fields of one playable printing, or false
// when the snapshot does not hold it (D-299).
func (x *Index) Printing(id string) (*mtgv1.Printing, bool) {
	p, ok := x.printings[id]
	return p, ok
}

// ByPrintingID finds a card by a Scryfall printing id (ManaBox join key).
func (x *Index) ByPrintingID(id string) (*mtgv1.Card, bool) {
	c, ok := x.byPrinting[id]
	return c, ok
}

func setNoKey(set, num string) string {
	return strings.ToLower(strings.TrimSpace(set)) + "/" + strings.ToLower(strings.TrimSpace(num))
}

// BySetCollector finds a card by set code and collector number, the
// ManaBox fallback join when the Scryfall id column is absent (PR-4).
func (x *Index) BySetCollector(set, num string) (*mtgv1.Card, bool) {
	c, ok := x.bySetNo[setNoKey(set, num)]
	return c, ok
}

// Len returns the card count.
func (x *Index) Len() int { return len(x.cards) }

// All returns every card in rank order. The slice is shared: do not
// modify it.
func (x *Index) All() []*mtgv1.Card { return x.cards }

// Tags returns the tag index, or nil when the snapshot had no tags file.
func (x *Index) Tags() *TagIndex { return x.tags }

// SearchQuery is the structured filter set for Search.
type SearchQuery struct {
	// ColorsWithin keeps cards whose color identity fits these colors.
	// Empty means no color filter.
	ColorsWithin []mtgv1.Color
	TypeContains string
	Keywords     []string
	// OracleTags are Tagger slugs. A card must match one of them.
	OracleTags []string
	// LegalIn is a Scryfall format key, for example "commander".
	LegalIn string
	Limit   int
	Offset  int
}

// Search filters the database. Order: by EDHREC rank, unranked last.
// The order comes from the build-time sort, so a search never sorts.
func (x *Index) Search(q SearchQuery) (result []*mtgv1.Card, total int) {
	var tagSet map[string]bool
	if len(q.OracleTags) > 0 {
		tagSet = map[string]bool{}
		if x.tags != nil {
			for _, slug := range q.OracleTags {
				for _, oid := range x.tags.Resolve(slug) {
					tagSet[oid] = true
				}
			}
		}
	}
	var colorSet map[mtgv1.Color]bool
	if len(q.ColorsWithin) > 0 {
		colorSet = map[mtgv1.Color]bool{}
		for _, c := range q.ColorsWithin {
			colorSet[c] = true
		}
	}
	var matched []*mtgv1.Card
	for _, c := range x.cards {
		if q.LegalIn != "" {
			s := c.Legalities[q.LegalIn]
			if s != mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL && s != mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED {
				continue
			}
		}
		if colorSet != nil && !identityFits(c.ColorIdentity, colorSet) {
			continue
		}
		if q.TypeContains != "" && !strings.Contains(strings.ToLower(c.TypeLine), strings.ToLower(q.TypeContains)) {
			continue
		}
		if !hasAllKeywords(c.Keywords, q.Keywords) {
			continue
		}
		if tagSet != nil && !tagSet[c.OracleId] {
			continue
		}
		matched = append(matched, c)
	}
	// x.cards is rank-sorted at build time, so matched is too.
	total = len(matched)
	if q.Offset >= len(matched) {
		return nil, total
	}
	matched = matched[q.Offset:]
	if q.Limit > 0 && len(matched) > q.Limit {
		matched = matched[:q.Limit]
	}
	return matched, total
}

func rankOf(c *mtgv1.Card) int {
	if c.EdhrecRank == 0 {
		return 1 << 30
	}
	return int(c.EdhrecRank)
}

func identityFits(identity []mtgv1.Color, allowed map[mtgv1.Color]bool) bool {
	for _, c := range identity {
		if !allowed[c] {
			return false
		}
	}
	return true
}

func hasAllKeywords(have, want []string) bool {
	for _, w := range want {
		if !slices.Contains(have, w) {
			return false
		}
	}
	return true
}

// PaperSwaps counts the cards whose default printing was digital and
// whose paper printing the snapshot also held. LoadIndex logs it, so a
// snapshot that suddenly swaps thousands of cards is visible (D-221).
func (x *Index) PaperSwaps() int { return x.paperSwaps }
