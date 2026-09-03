package quality

import (
	"hash/fnv"
	"math/rand/v2"
	"slices"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
)

// The defect axes of a synthetic bad list (D-414). The engine breaks
// one axis of a real list, so each defect carries its own label.
const (
	DefectLands   = "lands"
	DefectCurve   = "curve"
	DefectColors  = "colors"
	DefectCopies  = "copies"
	DefectSynergy = "synergy"
)

// Defects lists the axes in the order the fit cycles them.
var Defects = []string{DefectLands, DefectCurve, DefectColors, DefectCopies, DefectSynergy}

// MinPlaysetCopies is the copies in playsets a deck needs before the
// copies axis breaks it, and MinNonbasicLands the nonbasic lands that
// fix two colors or more before the colors axis does (D-485).
const (
	MinPlaysetCopies = 12
	MinNonbasicLands = 4
)

// legalKeys maps a format to its legality column, as the candidate
// builder reads it.
var legalKeys = map[mtgv1.FormatId]string{
	mtgv1.FormatId_FORMAT_ID_COMMANDER: "commander",
	mtgv1.FormatId_FORMAT_ID_STANDARD:  "standard",
	mtgv1.FormatId_FORMAT_ID_MODERN:    "modern",
}

// pool is the legal nonland cards of a format inside one color
// identity, sorted by Oracle id, so a seeded pick is the same on every
// run.
type pool struct {
	idx    *cards.Index
	format mtgv1.FormatId
	cache  map[string][]*mtgv1.Card
	// seen holds the nonland cards the real lists of the format play.
	// A break draws from these, so it breaks the pairing and the shape
	// and not the card pool: random legal cards taught the model that a
	// card outside the top lists means a broken deck, and a themed
	// casual deck is full of those on purpose (D-488).
	seen map[string]bool
}

// MinSeenPool is the seen cards an identity needs before a break draws
// from them alone. A narrow identity falls back to every legal card.
const MinSeenPool = 100

func newPool(idx *cards.Index, format mtgv1.FormatId) *pool {
	return &pool{idx: idx, format: format, cache: map[string][]*mtgv1.Card{}}
}

// setSeen names the cards the real lists play, from the resolved lists.
func (p *pool) setSeen(reals []*Resolved) {
	p.seen = map[string]bool{}
	for _, r := range reals {
		for _, dc := range r.Deck.GetCards() {
			p.seen[dc.GetOracleId()] = true
		}
	}
	p.cache = map[string][]*mtgv1.Card{}
}

// cards answers the pool of an identity. An empty identity is every
// color, which is what a 60-card list reads as.
func (p *pool) cards(identity []mtgv1.Color) []*mtgv1.Card {
	key := colorKey(identity)
	if out, ok := p.cache[key]; ok {
		return out
	}
	allowed := candidates.ColorSet(identity)
	var all, seen []*mtgv1.Card
	for _, c := range p.idx.All() {
		if !legalIn(c, legalKeys[p.format]) || candidates.IsBasicLand(c) || slices.Contains(c.GetCardTypes(), "Land") {
			continue
		}
		if c.GetDefaultPrinting().GetDigital() {
			continue
		}
		if allowed != nil && !candidates.IdentityFits(c.GetColorIdentity(), allowed) {
			continue
		}
		all = append(all, c)
		if p.seen[c.GetOracleId()] {
			seen = append(seen, c)
		}
	}
	out := all
	if len(seen) >= MinSeenPool {
		out = seen
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GetOracleId() < out[j].GetOracleId() })
	p.cache[key] = out
	return out
}

func legalIn(c *mtgv1.Card, key string) bool {
	if key == "" {
		return true
	}
	s := c.GetLegalities()[key]
	return s == mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL || s == mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED
}

func colorKey(colors []mtgv1.Color) string {
	var b strings.Builder
	for _, c := range colors {
		b.WriteString(c.String())
	}
	return b.String()
}

// Synthesize breaks one axis of a resolved list and answers the bad
// list. The seed comes from the list's key and the axis, so a fit
// makes the same bad list every run. A list the axis can not break,
// for example a copies defect on a singleton deck, breaks on the
// synergy axis instead.
func Synthesize(r *Resolved, axis string, pl *pool, roles Roles) *Resolved {
	deck := r.Deck
	format := deck.GetFormat().GetId()
	commander := format == mtgv1.FormatId_FORMAT_ID_COMMANDER
	h := fnv.New64a()
	_, _ = h.Write([]byte(r.List.Key() + "|" + axis))
	rng := rand.New(rand.NewPCG(h.Sum64(), 7))
	identity := deckIdentity(deck, pl.idx, commander)
	candidatesPool := pl.cards(identity)
	inDeck := map[string]bool{}
	for _, dc := range deck.GetCards() {
		inDeck[dc.GetOracleId()] = true
	}
	for _, id := range deck.GetCommanderOracleIds() {
		inDeck[id] = true
	}
	// pick draws a random pool card the deck does not hold.
	pick := func(filter func(*mtgv1.Card) bool) *mtgv1.Card {
		for tries := 0; tries < 200 && len(candidatesPool) > 0; tries++ {
			c := candidatesPool[rng.IntN(len(candidatesPool))]
			if inDeck[c.GetOracleId()] || (filter != nil && !filter(c)) {
				continue
			}
			inDeck[c.GetOracleId()] = true
			return c
		}
		return nil
	}
	var out []*mtgv1.DeckCard
	clone := func() {
		out = make([]*mtgv1.DeckCard, 0, len(deck.GetCards()))
		for _, dc := range deck.GetCards() {
			out = append(out, &mtgv1.DeckCard{OracleId: dc.GetOracleId(), Name: dc.GetName(), Count: dc.GetCount(), Role: dc.GetRole()})
		}
	}
	add := func(c *mtgv1.Card, count int32) {
		if c == nil {
			return
		}
		var role mtgv1.CardRole
		if roles != nil {
			role = roles(c)
		}
		out = append(out, &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: count, Role: role})
	}
	isLand := func(dc *mtgv1.DeckCard) bool {
		c, ok := pl.idx.ByOracleID(dc.GetOracleId())
		return ok && slices.Contains(c.GetCardTypes(), "Land")
	}
	mv := func(dc *mtgv1.DeckCard) float64 {
		c, _ := pl.idx.ByOracleID(dc.GetOracleId())
		return c.GetManaValue()
	}
	// removeCopies takes n copies off the rows the order names first.
	removeCopies := func(n int, order func(a, b *mtgv1.DeckCard) bool, keep func(*mtgv1.DeckCard) bool) int {
		rows := make([]*mtgv1.DeckCard, 0, len(out))
		for _, dc := range out {
			if keep == nil || keep(dc) {
				rows = append(rows, dc)
			}
		}
		sort.SliceStable(rows, func(i, j int) bool { return order(rows[i], rows[j]) })
		removed := 0
		for _, dc := range rows {
			if removed >= n {
				break
			}
			take := min(int(dc.GetCount()), n-removed)
			dc.Count -= int32(take)
			removed += take
		}
		out = slices.DeleteFunc(out, func(dc *mtgv1.DeckCard) bool { return dc.GetCount() <= 0 })
		return removed
	}
	fill := func(n int, filter func(*mtgv1.Card) bool) {
		for i := 0; i < n; i++ {
			c := pick(filter)
			if c == nil {
				return
			}
			count := int32(1)
			if !commander && n-i >= 4 && rng.IntN(2) == 0 {
				count = 4
				i += 3
			}
			add(c, count)
		}
	}
	byRandom := func(_, _ *mtgv1.DeckCard) bool { return rng.IntN(2) == 0 }

	// A break must be material, or it is no defect and the label lies.
	// A deck with under three playsets loses too little to the copies
	// axis, and a deck with under four nonbasic lands too little to the
	// colors axis, so each breaks on synergy instead (D-485).
	if axis == DefectCopies {
		playsetCopies := 0
		for _, dc := range deck.GetCards() {
			if c, ok := pl.idx.ByOracleID(dc.GetOracleId()); ok && dc.GetCount() >= 4 && !slices.Contains(c.GetCardTypes(), "Land") {
				playsetCopies += int(dc.GetCount())
			}
		}
		if commander || playsetCopies < MinPlaysetCopies {
			axis = DefectSynergy
		}
	}
	if axis == DefectColors {
		// The lands that fix: nonbasic lands that make two colors or
		// more. A utility land of one color starves nothing when it
		// becomes a basic.
		fixing := 0
		for _, dc := range deck.GetCards() {
			if c, ok := pl.idx.ByOracleID(dc.GetOracleId()); ok && slices.Contains(c.GetCardTypes(), "Land") && !candidates.IsBasicLand(c) && coloredSources(c) >= 2 {
				fixing += int(dc.GetCount())
			}
		}
		if fixing < MinNonbasicLands || len(identity) < 2 {
			axis = DefectSynergy
		}
	}
	clone()
	switch axis {
	case DefectLands:
		// A third of the lands become spells.
		lands := 0
		for _, dc := range out {
			if isLand(dc) {
				lands += int(dc.GetCount())
			}
		}
		n := removeCopies(lands/3, byRandom, isLand)
		fill(n, nil)
	case DefectCurve:
		// The cheapest spells become six-drops and up.
		n := removeCopies(15, func(a, b *mtgv1.DeckCard) bool { return mv(a) < mv(b) }, func(dc *mtgv1.DeckCard) bool { return !isLand(dc) })
		fill(n, func(c *mtgv1.Card) bool { return c.GetManaValue() >= 6 })
	case DefectColors:
		// Every nonbasic land becomes a basic of one color, so the
		// sources of the other colors collapse.
		basic, ok := pl.idx.ByName(basicOf(identity[0]))
		n := removeCopies(1<<20, byRandom, func(dc *mtgv1.DeckCard) bool {
			c, found := pl.idx.ByOracleID(dc.GetOracleId())
			return found && slices.Contains(c.GetCardTypes(), "Land") && !candidates.IsBasicLand(c)
		})
		if ok {
			add(basic, int32(n))
		}
	case DefectCopies:
		// Every playset becomes one copy, and singletons fill the room.
		n := 0
		for _, dc := range out {
			if dc.GetCount() >= 4 && !isLand(dc) {
				n += int(dc.GetCount()) - 1
				dc.Count = 1
			}
		}
		for i := 0; i < n; i++ {
			add(pick(nil), 1)
		}
	default:
		// Half the spells become random legal cards.
		spells := 0
		for _, dc := range out {
			if !isLand(dc) {
				spells += int(dc.GetCount())
			}
		}
		n := removeCopies(spells/2, byRandom, func(dc *mtgv1.DeckCard) bool { return !isLand(dc) })
		fill(n, nil)
	}
	list := *r.List
	list.Source = meta.SourceSynthetic
	list.ID = r.List.Key() + "/" + axis
	list.Tier = meta.TierBad
	list.Defect = axis
	return &Resolved{
		List: &list,
		Deck: &mtgv1.Deck{Format: deck.GetFormat(), CommanderOracleIds: deck.GetCommanderOracleIds(), Cards: out},
	}
}

// deckIdentity is the commander's identity, or the union of the card
// colors of a 60-card deck.
func deckIdentity(deck *mtgv1.Deck, idx *cards.Index, commander bool) []mtgv1.Color {
	seen := map[mtgv1.Color]bool{}
	var out []mtgv1.Color
	addAll := func(colors []mtgv1.Color) {
		for _, c := range colors {
			if !seen[c] {
				seen[c] = true
				out = append(out, c)
			}
		}
	}
	if commander {
		for _, id := range deck.GetCommanderOracleIds() {
			if c, ok := idx.ByOracleID(id); ok {
				addAll(c.GetColorIdentity())
			}
		}
		if len(out) > 0 {
			return out
		}
	}
	for _, dc := range deck.GetCards() {
		if c, ok := idx.ByOracleID(dc.GetOracleId()); ok {
			addAll(c.GetColors())
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// basicOf names the basic land of a color.
func basicOf(c mtgv1.Color) string {
	switch c {
	case mtgv1.Color_COLOR_W:
		return "Plains"
	case mtgv1.Color_COLOR_U:
		return "Island"
	case mtgv1.Color_COLOR_B:
		return "Swamp"
	case mtgv1.Color_COLOR_R:
		return "Mountain"
	case mtgv1.Color_COLOR_G:
		return "Forest"
	default:
		return "Wastes"
	}
}

// coloredSources counts the colors a land makes.
func coloredSources(c *mtgv1.Card) int {
	n := 0
	for _, col := range c.GetProducedMana() {
		if col != mtgv1.Color_COLOR_C {
			n++
		}
	}
	return n
}
