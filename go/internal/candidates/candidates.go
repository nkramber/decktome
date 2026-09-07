// Package candidates builds the ranked, role-grouped card shortlist that
// the deck generator sees (roadmap PR-6). The model never sees the whole
// database. It picks from this list.
//
// Signals, in order of weight: Oracle tags (theme fit, F-5), keywords and
// type lines, Oracle text needles, EDHREC rank (popularity). Filters:
// legality, color identity, and the pool mode (D-37). Ownership partitions
// the list in the owned modes. It never gates a card in any-card mode.
package candidates

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/rules"
)

// Request describes one shortlist.
type Request struct {
	// Format selects legality and the deck shape. House rules skip legality.
	Format mtgv1.FormatId
	// Colors is the allowed color identity. Empty means every color.
	Colors []mtgv1.Color
	// Theme is the user's words, for example "lifegain aristocrats".
	Theme string
	// CommanderOracleIDs are excluded from the 99 candidates.
	CommanderOracleIDs []string
	// PoolRule is the ownership mode (D-37). UNSPECIFIED means any-card
	// without Owned, owned-first with it.
	PoolRule mtgv1.PoolRule
	// Owned maps Oracle id to owned count. Nil means no collection.
	Owned map[string]int32
	// Bracket is the Commander bracket, 0 when unknown. Brackets 1 and 2
	// drop Game Changers from the list, so the model never picks one.
	// Brackets 1 to 3 drop mass land denial, and bracket 1 drops extra
	// turns, by the tags (PR-14A).
	Bracket int32
	// SetCodes limits the list to cards printed in these paper sets. It
	// holds a whole set family (D-376). Empty means every set. Basic
	// lands are out of the shortlist in any case, and a set limit never
	// filters them (D-378).
	SetCodes []string
	// ExcludeOracleIDs are cards the deck must not hold: the cards of a
	// precon the reader excluded, with no copy to spare (D-408). Both
	// pools drop them, and the rules check blocks one that slips through.
	ExcludeOracleIDs []string
	// OutsideRoles names the roles that may be filled from outside
	// SetCodes, with the wanted count of each. A set family short of
	// mana cards gets them from the whole database, up to the count and
	// no further, and every such card is marked (D-382). Nil means no
	// card comes from outside.
	OutsideRoles map[mtgv1.CardRole]int
	// MetaBoost gives the quality model's card signal in [0,1] per
	// Oracle id: the inclusion rate in the top lists of the format
	// against the best rate (PR-14B). Nil with no model.
	MetaBoost func(oracleID string) float64
	// CommanderSignal gives the bracket 5 power signal of a commander or
	// a pair, in [0,1] (PR-14B, OQ-48). A bracket 4 or 5 request ranks
	// its commanders on it first, and theme breaks the tie. Nil with no
	// model, and the pool ranks on theme and popularity alone.
	CommanderSignal func(oracleIDs ...string) float64
	// WantBackground narrows a pair request to pairs that hold a
	// Background, because a Background rarely ranks on theme alone (D-154).
	WantBackground bool
	// WantPair asks CommanderPool to offer two-commander pairs beside
	// single commanders. The pool also offers them when too few single
	// commanders fit the colors, which is the four-color case (D-154).
	WantPair bool
	// Limits override the defaults. Zero fields keep the default.
	Limits Limits
}

// Limits bounds the list.
type Limits struct {
	// Total caps the main list (owned candidates in the owned modes, all
	// candidates in any-card mode).
	Total int
	// Upgrades caps the unowned upgrade list in owned-first mode.
	Upgrades int
	// PerRole caps each role in the main list.
	PerRole map[mtgv1.CardRole]int
}

// staplePenalty halves the score of a card with no theme signal.
const staplePenalty = 0.5

// DefaultLimits follows the roadmap numbers: about 300 candidates by role,
// about 50 upgrades. The role caps sum to more than Total on purpose: a
// role that comes up short leaves room for the others, and when every
// role is full the lowest-scored cards go, whatever their role.
var DefaultLimits = Limits{
	Total:    300,
	Upgrades: 50,
	PerRole: map[mtgv1.CardRole]int{
		mtgv1.CardRole_CARD_ROLE_LAND:        40,
		mtgv1.CardRole_CARD_ROLE_RAMP:        30,
		mtgv1.CardRole_CARD_ROLE_DRAW:        30,
		mtgv1.CardRole_CARD_ROLE_REMOVAL:     30,
		mtgv1.CardRole_CARD_ROLE_WIPE:        12,
		mtgv1.CardRole_CARD_ROLE_INTERACTION: 20,
		mtgv1.CardRole_CARD_ROLE_WINCON:      15,
		mtgv1.CardRole_CARD_ROLE_THREAT:      50,
		mtgv1.CardRole_CARD_ROLE_SYNERGY:     80,
		mtgv1.CardRole_CARD_ROLE_OTHER:       10,
	},
}

// The role keys of themes.json the bracket cut reads (PR-14A).
const (
	roleMassLandDenial = "mass_land_denial"
	roleExtraTurn      = "extra_turn"
)

// Candidate is one shortlisted card with its evidence.
type Candidate struct {
	Card  *mtgv1.Card
	Role  mtgv1.CardRole
	Score float64
	// Pop is the popularity share of the card, 1 for the most played
	// card of the snapshot and 0 for a card with no rank. The land cap
	// reads it apart from the score (D-450).
	Pop float64
	// Fix counts the deck colors a land produces, so a dual in the
	// colors outranks a fetch land that produces nothing. Zero for a
	// nonland (D-450).
	Fix int
	// Themed says the theme matched the card. The land cap fills its
	// theme half with these alone (D-450).
	Themed bool
	// Partner is the second commander of a pair, and nil for every other
	// candidate. The pair carries the union of the two color identities,
	// which is what lets a request reach four colors (D-154).
	Partner *mtgv1.Card
	// Owned is the owned count, 0 without a collection.
	Owned int32
	// Outside marks a card the request's sets do not hold. It reaches
	// the list only through OutsideRoles, and the deck marks it (D-382,
	// D-383).
	Outside bool
	// Signals lists why the card is here, for example "tag:lifegain".
	Signals []string
}

// List is the shortlist.
type List struct {
	// Candidates is the main list, grouped by role in role order.
	Candidates []Candidate
	// Upgrades lists unowned cards worth a purchase (owned-first only).
	Upgrades []Candidate
	// Theme is the resolved theme: which words matched which signals.
	Theme ThemeMatch
	// Stats counts the funnel.
	Stats Stats
}

// Stats counts the funnel for the gate doc.
type Stats struct {
	Pool        int // cards after legality, color, and commander filters
	OnTheme     int // cards with at least one theme signal
	Owned       int // on-theme or role cards the collection covers
	Returned    int
	UpgradeSize int
	// OnThemeOwned counts owned cards with a theme signal (owned modes).
	OnThemeOwned int
	// ThinTheme marks an owned mode where the collection holds fewer than
	// ThinThemeFloor on-theme cards. PR-7 asks the pool-mode question
	// again on this flag (D-63).
	ThinTheme bool
	// InSet counts the pool cards the request's sets hold. It is zero when
	// no set limit applies, because no card is then inside a set the
	// reader named. The viability floor of D-380 reads the same number
	// through CountInSets.
	InSet int
	// Outside counts the cards the fill took from outside the sets
	// (D-382).
	Outside int
}

// ThinThemeFloor is the on-theme owned count under which a collection is
// too thin for owned-only play. The corpus guide asks for 20 to 35 theme
// pieces in a Commander deck.
const ThinThemeFloor = 30

// Builder holds the loaded theme table.
type Builder struct {
	themes *themeTable
}

// New loads the embedded theme table.
func New() (*Builder, error) {
	t, err := loadThemes()
	if err != nil {
		return nil, err
	}
	return &Builder{themes: t}, nil
}

// Roles answers a role reader over one index: the role a card fills
// with no theme in play. The quality fit of PR-14B reads published
// lists through it, so a list's role counts read the way a built
// deck's do. The reader resolves the role tags once.
func (b *Builder) Roles(idx *cards.Index) func(c *mtgv1.Card) mtgv1.CardRole {
	roleTags := b.themes.roleSets(idx.Tags())
	useText := idx.Tags().Len() == 0
	return func(c *mtgv1.Card) mtgv1.CardRole {
		role, _ := assignRole(c, roleTags, false, useText)
		return role
	}
}

// legalKeys maps a format to its Scryfall legality column. The app builds
// three formats (D-155). HOUSE has no key on purpose: the user defined
// the rules, so no ban list applies (D-3, D-306). An unknown format skips
// the legality filter as well.
var legalKeys = map[mtgv1.FormatId]string{
	mtgv1.FormatId_FORMAT_ID_COMMANDER: "commander",
	mtgv1.FormatId_FORMAT_ID_STANDARD:  "standard",
	mtgv1.FormatId_FORMAT_ID_MODERN:    "modern",
}

// Build makes the shortlist. It is deterministic: the same index and
// request give the same list. PR-9 adds the seeded shuffle on top.
func (b *Builder) Build(idx *cards.Index, req Request) (*List, error) {
	if idx == nil {
		return nil, fmt.Errorf("candidates: no card index")
	}
	lim := req.Limits.withDefaults()
	// A set limit drops the per-role caps as well as the theme cut
	// (D-379). The caps shape a shortlist drawn from 12,718 cards. Drawn
	// from 128, they only lose: the "other" cap of 10 cut 40 of the 128
	// cards the Hobbit family offers in black-red, and a Commander deck
	// needs 99. The total cap still bounds the list, and a caller that
	// names its own caps keeps them.
	if len(req.SetCodes) > 0 && req.Limits.PerRole == nil {
		lim.PerRole = map[mtgv1.CardRole]int{}
	}
	mode := req.PoolRule
	if mode == mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		mode = mtgv1.PoolRule_POOL_RULE_ANY_CARD
		if req.Owned != nil {
			mode = mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
		}
	}
	if mode != mtgv1.PoolRule_POOL_RULE_ANY_CARD && req.Owned == nil {
		return nil, fmt.Errorf("candidates: pool rule %s needs a collection", mode)
	}
	theme := b.themes.matchIn(req.Theme, idx)
	roleTags := b.themes.roleSets(idx.Tags())
	// Text fallbacks stand in for the tags only when the snapshot has
	// none. With tags loaded, an untagged card is not a staple.
	useText := idx.Tags().Len() == 0
	legalKey := legalKeys[req.Format]
	colorSet := colorSetOf(req.Colors)
	setCodes := cards.CodeSet(req.SetCodes)
	excluded := map[string]bool{}
	for _, id := range req.CommanderOracleIDs {
		excluded[id] = true
	}
	for _, id := range req.ExcludeOracleIDs {
		excluded[id] = true
	}
	maxRank := maxRankOf(idx)

	var stats Stats
	var scored []Candidate
	fired := firedSignals{}
	for _, c := range idx.All() {
		if excluded[c.OracleId] || IsBasicLand(c) {
			continue
		}
		if legalKey != "" && !legalIn(c, legalKey) {
			continue
		}
		// Every format offers paper cards only. A card with no paper
		// printing never reaches the shortlist, and a banned card stays
		// allowed where no legality key applies (D-306).
		if !hasPaperPrinting(c) {
			continue
		}
		if colorSet != nil && !IdentityFits(c.ColorIdentity, colorSet) {
			continue
		}
		if c.GameChanger && req.Bracket > 0 && req.Bracket <= 2 {
			continue
		}
		// The content rules of a bracket, as the tags read them: no mass
		// land denial through bracket 3, and no extra turn at bracket 1
		// (rules/brackets.json). The tags seed the list, and Commander
		// Spellbook checks the deck after the build (PR-14A, F-5).
		if req.Bracket > 0 && req.Bracket <= 3 && roleTags[roleMassLandDenial][c.OracleId] {
			continue
		}
		if req.Bracket == 1 && roleTags[roleExtraTurn][c.OracleId] {
			continue
		}
		stats.Pool++
		themeScore, signals := theme.score(c)
		fired.mark(signals)
		role, roleSignal := assignRole(c, roleTags, themeScore > 0, useText)
		if roleSignal != "" {
			signals = append(signals, roleSignal)
		}
		if themeScore > 0 {
			stats.OnTheme++
		}
		// The set limit (D-373). A card the sets do not hold reaches the
		// list only through the mana fill, which OutsideRoles names.
		outside := !cards.InSets(c, setCodes)
		if outside {
			if len(req.OutsideRoles) == 0 || req.OutsideRoles[role] == 0 {
				continue
			}
			signals = append(signals, "outside the sets")
		} else if setCodes != nil {
			stats.InSet++
		}
		// A card with no theme signal stays only when it fills a staple
		// role (lands, ramp, draw, removal, wipes, interaction). Threats
		// and synergy pieces need a theme signal.
		//
		// A set limit lifts the cut (D-379). The Hobbit family offers 128
		// cards in black-red, and the cut left 75 to 89 of them, which
		// can not fill 99. Inside a set the theme ranks the list and does
		// not cut it: the staple penalty below still puts every on-theme
		// card first.
		if themeScore == 0 && !stapleRole(role) && setCodes == nil {
			continue
		}
		pop := popularity(c, maxRank)
		score := themeScore*0.7 + pop*0.3
		if req.MetaBoost != nil {
			score += req.MetaBoost(c.OracleId) * 0.1
		}
		// A staple with no theme signal ranks below every on-theme card of
		// its score band. Theme leads, popularity breaks ties (F-5).
		if themeScore == 0 {
			score *= staplePenalty
		}
		owned := req.Owned[c.OracleId]
		if owned > 0 {
			stats.Owned++
			if themeScore > 0 {
				stats.OnThemeOwned++
			}
		}
		scored = append(scored, Candidate{Card: c, Role: role, Score: score, Pop: pop, Fix: fixCount(c, colorSet), Themed: themeScore > 0, Owned: owned, Outside: outside, Signals: signals})
	}
	sortCandidates(scored)
	theme.Unmatched = theme.unmatchedWords(fired)

	// The cards the sets hold rank on their own. The outside cards are a
	// fill and never a competitor, so they are held back and added after
	// the caps run (D-382).
	inSet, outsideCards := scored, []Candidate(nil)
	if setCodes != nil && len(req.OutsideRoles) > 0 {
		inSet, outsideCards = splitOutside(scored)
	}
	var main, upgrades []Candidate
	switch mode {
	case mtgv1.PoolRule_POOL_RULE_ANY_CARD:
		main = capByRole(inSet, lim)
	case mtgv1.PoolRule_POOL_RULE_OWNED_ONLY:
		main = capByRole(filterOwned(inSet, true), lim)
	case mtgv1.PoolRule_POOL_RULE_OWNED_FIRST:
		main = ownedFirst(inSet, lim)
		upgrades = topUpgrades(filterOwned(inSet, false), main, lim.Upgrades)
	}
	if fill := outsideFill(main, outsideCards, req.OutsideRoles); len(fill) > 0 {
		main = mergeByRole(main, fill)
		stats.Outside = len(fill)
	}
	stats.Returned = len(main)
	stats.UpgradeSize = len(upgrades)
	stats.ThinTheme = mode != mtgv1.PoolRule_POOL_RULE_ANY_CARD && stats.OnThemeOwned < ThinThemeFloor
	return &List{Candidates: main, Upgrades: upgrades, Theme: theme, Stats: stats}, nil
}

func (l Limits) withDefaults() Limits {
	out := l
	if out.Total <= 0 {
		out.Total = DefaultLimits.Total
	}
	if out.Upgrades <= 0 {
		out.Upgrades = DefaultLimits.Upgrades
	}
	if out.PerRole == nil {
		out.PerRole = DefaultLimits.PerRole
	}
	return out
}

// roleOrder is the display and cap order.
var roleOrder = []mtgv1.CardRole{
	mtgv1.CardRole_CARD_ROLE_LAND,
	mtgv1.CardRole_CARD_ROLE_RAMP,
	mtgv1.CardRole_CARD_ROLE_DRAW,
	mtgv1.CardRole_CARD_ROLE_REMOVAL,
	mtgv1.CardRole_CARD_ROLE_WIPE,
	mtgv1.CardRole_CARD_ROLE_INTERACTION,
	mtgv1.CardRole_CARD_ROLE_WINCON,
	mtgv1.CardRole_CARD_ROLE_THREAT,
	mtgv1.CardRole_CARD_ROLE_SYNERGY,
	mtgv1.CardRole_CARD_ROLE_OTHER,
}

// capByRole keeps the best cards per role, in role order, under the total.
// When the role caps together exceed the total, the lowest-scored cards
// go across every role. A plain cut at Total dropped the whole tail of
// the last roles, which is where the synergy pieces sit.
func capByRole(in []Candidate, lim Limits) []Candidate {
	byRole := map[mtgv1.CardRole][]Candidate{}
	for _, c := range in {
		byRole[c.Role] = append(byRole[c.Role], c)
	}
	var out []Candidate
	for _, r := range roleOrder {
		cs := byRole[r]
		if n := lim.PerRole[r]; n > 0 && len(cs) > n {
			if r == mtgv1.CardRole_CARD_ROLE_LAND {
				cs = capLands(cs, n)
			} else {
				cs = cs[:n]
			}
		}
		out = append(out, cs...)
	}
	if len(out) <= lim.Total {
		return out
	}
	ranked := slices.Clone(out)
	sortCandidates(ranked)
	keep := make(map[*mtgv1.Card]bool, lim.Total)
	for _, c := range ranked[:lim.Total] {
		keep[c.Card] = true
	}
	kept := make([]Candidate, 0, lim.Total)
	for _, c := range out {
		if keep[c.Card] {
			kept = append(kept, c)
		}
	}
	return kept
}

// manaShare is the share of the land cap that goes to the mana order,
// theme or not: the duals of the colors and Command Tower. The theme
// ranks the rest of the cap. A theme whose word sits in land text can
// fill the whole cap with no fixing otherwise (F-32, D-450).
const manaShare = 0.5

// capLands takes n lands from a bucket in score order: the mana half
// first, then the theme lands in score order, then the mana order again
// when the theme matched too few lands. The mana order ranks by the deck
// colors a land produces, then by play, so a dual in the colors comes
// before Command Tower's cousins and an off-color fetch land comes last.
// The result keeps the bucket's order, so the theme lands still read
// first.
func capLands(cs []Candidate, n int) []Candidate {
	if n <= 0 || len(cs) <= n {
		return cs
	}
	mana := int(math.Round(float64(n) * manaShare))
	// A land that makes two or more of the deck colors is fixing, and
	// play ranks the fixing. A count above two would put a Thriving land
	// or Cavern of Souls, which Scryfall lists as every color, over a
	// shock land in a three-color deck.
	byMana := slices.Clone(cs)
	sort.SliceStable(byMana, func(i, j int) bool {
		if fi, fj := min(byMana[i].Fix, 2), min(byMana[j].Fix, 2); fi != fj {
			return fi > fj
		}
		return byMana[i].Pop > byMana[j].Pop
	})
	keep := make(map[*mtgv1.Card]bool, n)
	for _, c := range byMana[:mana] {
		keep[c.Card] = true
	}
	for _, c := range cs {
		if len(keep) >= n {
			break
		}
		if c.Themed {
			keep[c.Card] = true
		}
	}
	for _, c := range byMana {
		if len(keep) >= n {
			break
		}
		keep[c.Card] = true
	}
	out := make([]Candidate, 0, n)
	for _, c := range cs {
		if keep[c.Card] {
			out = append(out, c)
		}
	}
	return out
}

// fixCount is the number of deck colors a land produces. With no color
// limit every color counts. A nonland is 0, whatever it produces: the
// mana half of the land cap is for lands.
func fixCount(c *mtgv1.Card, colorSet map[mtgv1.Color]bool) int {
	if !slices.Contains(c.GetCardTypes(), "Land") {
		return 0
	}
	n := 0
	for _, col := range c.GetProducedMana() {
		if col == mtgv1.Color_COLOR_C {
			continue
		}
		if colorSet == nil || colorSet[col] {
			n++
		}
	}
	return n
}

// topUpgrades picks the best unowned cards that beat the weakest owned
// card of the same role, so an upgrade is a real improvement.
func topUpgrades(unowned, main []Candidate, limit int) []Candidate {
	floor := map[mtgv1.CardRole]float64{}
	// A card the main list already carries is not an upgrade. It is in
	// the deck, and the buy list already names it when the user does not
	// own it (D-359).
	inMain := make(map[*mtgv1.Card]bool, len(main))
	for _, c := range main {
		inMain[c.Card] = true
		if v, ok := floor[c.Role]; !ok || c.Score < v {
			floor[c.Role] = c.Score
		}
	}
	var out []Candidate
	for _, c := range unowned {
		if inMain[c.Card] {
			continue
		}
		if v, ok := floor[c.Role]; ok && c.Score <= v {
			continue
		}
		out = append(out, c)
		if len(out) >= limit {
			break
		}
	}
	return out
}

// OwnedFillFloor is the main-list size under which a collection can not
// build a deck on its own (D-362). A Commander deck is 99 cards, and
// deck gate run 8 built one with nothing to buy from a shortlist of 168
// owned names. Above this floor the collection is enough, and the fill
// stays away.
//
// The floor is not the shortlist cap. Filling to the cap offered the
// model 120 unowned cards it did not need, and gate run 9 turned four
// decks that cost nothing into decks that cost $40 to $168. A reader
// who says "my collection first" is not asking for that.
const OwnedFillFloor = 150

// ownedFirst builds the main list from the collection, then fills what
// the collection can not (D-359).
//
// The collection leads: every owned candidate that fits the caps is in
// the list before one unowned card is read. A thin collection then
// leaves a hole, and a deck with a hole is not a deck, so the database
// fills the rest in score order. A reader who says "only cards I own"
// takes POOL_RULE_OWNED_ONLY and no fill at all.
//
// The color identity is applied before this, so a card outside the
// deck's colors is in neither half.
func ownedFirst(in []Candidate, lim Limits) []Candidate {
	owned := capByRole(filterOwned(in, true), lim)
	// The fill reaches the floor, never the cap (D-362). A collection
	// that already builds a deck is left exactly as it is.
	target := OwnedFillFloor
	if lim.Total < target {
		target = lim.Total
	}
	room := target - len(owned)
	if room <= 0 {
		return owned
	}
	// The caps the owned half did not use. A role the collection filled
	// takes no fill, and the total never grows.
	used := map[mtgv1.CardRole]int{}
	for _, c := range owned {
		used[c.Role]++
	}
	left := Limits{Total: room, PerRole: map[mtgv1.CardRole]int{}}
	for role, n := range lim.PerRole {
		if free := n - used[role]; free > 0 {
			left.PerRole[role] = free
		}
	}
	fill := capByRole(filterOwned(in, false), left)
	if len(fill) == 0 {
		return owned
	}
	// The list reads by role, the way capByRole returns one, so the
	// owned half and the fill do not sit in two blocks.
	merged := make([]Candidate, 0, len(owned)+len(fill))
	byRole := map[mtgv1.CardRole][]Candidate{}
	for _, c := range append(append([]Candidate{}, owned...), fill...) {
		byRole[c.Role] = append(byRole[c.Role], c)
	}
	for _, r := range roleOrder {
		merged = append(merged, byRole[r]...)
	}
	return merged
}

// splitOutside separates the cards the sets hold from the ones the fill
// may take.
func splitOutside(in []Candidate) (inSet, outside []Candidate) {
	for _, c := range in {
		if c.Outside {
			outside = append(outside, c)
			continue
		}
		inSet = append(inSet, c)
	}
	return inSet, outside
}

// outsideFill takes the best outside cards of each named role, up to the
// shortfall against the role's wanted count (D-382). A role the sets
// already fill takes nothing.
//
// The fill is additive: it runs after the caps, because it exists to
// close a hole the caps can not close. A shortfall of ten ramp cards is
// ten names, so the total grows by a little and never by a lot.
func outsideFill(main, outside []Candidate, want map[mtgv1.CardRole]int) []Candidate {
	if len(outside) == 0 || len(want) == 0 {
		return nil
	}
	have := map[mtgv1.CardRole]int{}
	for _, c := range main {
		have[c.Role]++
	}
	// outside is already in score order, because the whole list was
	// sorted before the split.
	var out []Candidate
	taken := map[mtgv1.CardRole]int{}
	for _, c := range outside {
		room := want[c.Role] - have[c.Role] - taken[c.Role]
		if room <= 0 {
			continue
		}
		taken[c.Role]++
		out = append(out, c)
	}
	return out
}

// mergeByRole joins two lists and returns them in role order, which is
// the order capByRole emits. Without it the fill would sit in one block
// at the end of the shortlist.
func mergeByRole(a, b []Candidate) []Candidate {
	byRole := map[mtgv1.CardRole][]Candidate{}
	for _, c := range a {
		byRole[c.Role] = append(byRole[c.Role], c)
	}
	for _, c := range b {
		byRole[c.Role] = append(byRole[c.Role], c)
	}
	out := make([]Candidate, 0, len(a)+len(b))
	for _, r := range roleOrder {
		out = append(out, byRole[r]...)
	}
	return out
}

func filterOwned(in []Candidate, want bool) []Candidate {
	var out []Candidate
	for _, c := range in {
		if (c.Owned > 0) == want {
			out = append(out, c)
		}
	}
	return out
}

// sortCandidates orders by score, then rank, then name. The name breaks
// every tie so the list is stable across runs.
func sortCandidates(cs []Candidate) {
	sort.SliceStable(cs, func(i, j int) bool {
		a, b := cs[i], cs[j]
		if a.Score != b.Score {
			return a.Score > b.Score
		}
		if ra, rb := rankOf(a.Card), rankOf(b.Card); ra != rb {
			return ra < rb
		}
		return a.Card.Name < b.Card.Name
	})
}

func rankOf(c *mtgv1.Card) int32 {
	if c.EdhrecRank == 0 {
		return math.MaxInt32
	}
	return c.EdhrecRank
}

func maxRankOf(idx *cards.Index) float64 {
	var m int32
	for _, c := range idx.All() {
		if c.EdhrecRank > m {
			m = c.EdhrecRank
		}
	}
	if m == 0 {
		return 1
	}
	return float64(m)
}

// popularity maps the EDHREC rank to [0,1]. Unranked cards get 0.
func popularity(c *mtgv1.Card, maxRank float64) float64 {
	if c.EdhrecRank == 0 {
		return 0
	}
	return 1 - float64(c.EdhrecRank-1)/maxRank
}

func legalIn(c *mtgv1.Card, key string) bool {
	s := c.Legalities[key]
	return s == mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL || s == mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED
}

func colorSetOf(colors []mtgv1.Color) map[mtgv1.Color]bool {
	if len(colors) == 0 {
		return nil
	}
	m := map[mtgv1.Color]bool{}
	for _, c := range colors {
		m[c] = true
	}
	return m
}

// IdentityMatches reports whether a commander's color identity holds
// every color named and no other. It is the exported form of the test
// CommanderPool applies (D-148), so PR-7 can check a name it already
// offered against colors that arrived later (D-153).
func IdentityMatches(identity []mtgv1.Color, colors []mtgv1.Color) bool {
	set := colorSetOf(colors)
	if set == nil {
		return true
	}
	return IdentityFits(identity, set) && identityCovers(identity, set)
}

// identityCovers reports whether the identity holds every color the user
// named. IdentityFits is a subset test, which is right for the 99: a
// mono-red card belongs in a blue-red deck. It is wrong for the commander,
// because the commander's identity is the deck's identity (CR 903.4), so
// a commander must hold every named color (D-148).
//
// Colorless is skipped on both sides, as it is in IdentityFits. It is not
// a color, so it can neither fail the test nor satisfy it.
func identityCovers(identity []mtgv1.Color, allowed map[mtgv1.Color]bool) bool {
	have := map[mtgv1.Color]bool{}
	for _, c := range identity {
		have[c] = true
	}
	for c := range allowed {
		if c == mtgv1.Color_COLOR_C {
			continue
		}
		if !have[c] {
			return false
		}
	}
	return true
}

// IdentityFits reports whether a color identity is a subset of the allowed
// colors. Colorless is not a color, so it never fails the test. The
// generate package shares it, so the 99 and the precon swap apply one
// test (D-154).
func IdentityFits(identity []mtgv1.Color, allowed map[mtgv1.Color]bool) bool {
	for _, c := range identity {
		if c == mtgv1.Color_COLOR_C {
			continue
		}
		if !allowed[c] {
			return false
		}
	}
	return true
}

// ColorSet is the allowed-color set IdentityFits reads. Empty colors give
// nil, which every caller reads as "any color".
func ColorSet(colors []mtgv1.Color) map[mtgv1.Color]bool { return colorSetOf(colors) }

// IsBasicLand reports whether a card is a basic land: the Basic supertype
// on a Land. The generate package shares it, so one test decides what a
// shortlist omits and what a deck pads with (D-225).
func IsBasicLand(c *mtgv1.Card) bool {
	return slices.Contains(c.GetSupertypes(), "Basic") && slices.Contains(c.GetCardTypes(), "Land")
}

// BasicLandByOracle is IsBasicLand over an Oracle id, for the precon
// exclusion and the ownership check (D-37, D-523). An id the index lacks
// is not a basic land, and a nil index knows none.
func BasicLandByOracle(idx *cards.Index) func(oracleID string) bool {
	return func(oracleID string) bool {
		if idx == nil {
			return false
		}
		c, ok := idx.ByOracleID(oracleID)
		return ok && IsBasicLand(c)
	}
}

// FoldName is the name match key: lower case, with the outer spaces
// removed. Nothing else is folded, because a punctuation change makes a
// different card name (F-13).
func FoldName(s string) string { return strings.ToLower(strings.TrimSpace(foldQuotes.Replace(s))) }

// foldQuotes reads a curly apostrophe as the straight one a card name
// holds. Deck gate run 13 lost six cards to "Commander’s Sphere".
var foldQuotes = strings.NewReplacer("\u2019", "'", "\u2018", "'", "\u201c", "\"", "\u201d", "\"")

// hasPaperPrinting reports whether the card's shown printing is a paper
// one. The index swaps a digital default for a paper printing when one
// exists (D-221), so a digital default means the card has none (D-306).
func hasPaperPrinting(c *mtgv1.Card) bool { return !c.GetDefaultPrinting().GetDigital() }

func stapleRole(r mtgv1.CardRole) bool {
	switch r {
	case mtgv1.CardRole_CARD_ROLE_LAND, mtgv1.CardRole_CARD_ROLE_RAMP, mtgv1.CardRole_CARD_ROLE_DRAW,
		mtgv1.CardRole_CARD_ROLE_REMOVAL, mtgv1.CardRole_CARD_ROLE_WIPE, mtgv1.CardRole_CARD_ROLE_INTERACTION:
		return true
	}
	return false
}

// RoleName is the short lowercase role name for logs and the gate doc.
func RoleName(r mtgv1.CardRole) string {
	return strings.ToLower(strings.TrimPrefix(r.String(), "CARD_ROLE_"))
}

// Commanders returns up to n commander-eligible candidates for a request,
// best first. PR-7 uses it to name real commanders in the commander
// question instead of a placeholder.
func (b *Builder) Commanders(idx *cards.Index, req Request, n int) ([]Candidate, error) {
	if n <= 0 {
		n = 3
	}
	pool, err := b.CommanderPool(idx, req)
	if err != nil {
		return nil, err
	}
	if len(pool) > n {
		pool = pool[:n]
	}
	return pool, nil
}

// CommanderPool ranks every commander that fits the request, best first.
//
// It walks the card index itself rather than the 99-card shortlist. The
// shortlist ends in capByRole, which emits one role bucket after another
// with lands first, so it is not ordered by score at all (D-94).
//
// Two rules differ from the 99. A commander must carry a theme signal,
// because the staple-role fallback that keeps a useful land in the deck
// says nothing about leading it. Role caps do not apply, because one card
// fills no role quota.
//
// The length of the result is the weak-pool signal PR-7 needs (D-63):
// an owned mode with no on-theme commander leaves it empty.
func (b *Builder) CommanderPool(idx *cards.Index, req Request) ([]Candidate, error) {
	if idx == nil {
		return nil, fmt.Errorf("candidates: no card index")
	}
	// An unasked pool rule ranks on quality alone. The commander offer
	// goes out before the pool question, and a collection must not turn
	// it into a list of the legends the user happens to own (D-293).
	mode := req.PoolRule
	if mode == mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		mode = mtgv1.PoolRule_POOL_RULE_ANY_CARD
	}
	if mode != mtgv1.PoolRule_POOL_RULE_ANY_CARD && req.Owned == nil {
		return nil, fmt.Errorf("candidates: pool rule %s needs a collection", mode)
	}
	theme := b.themes.matchIn(req.Theme, idx)
	colorSet := colorSetOf(req.Colors)
	setCodes := cards.CodeSet(req.SetCodes)
	maxRank := maxRankOf(idx)

	excluded := map[string]bool{}
	for _, id := range req.ExcludeOracleIDs {
		excluded[id] = true
	}

	var out []Candidate
	for _, c := range idx.All() {
		if !c.GetCanBeCommander() || !legalIn(c, legalKeys[mtgv1.FormatId_FORMAT_ID_COMMANDER]) {
			continue
		}
		// A commander of an excluded precon can not lead the deck (D-408).
		if excluded[c.OracleId] {
			continue
		}
		// Every format offers paper cards only (D-306). The unthemed fill
		// of D-367 always tested this, and the themed half above it never
		// did, so a digital-only legend could lead an offer.
		if !hasPaperPrinting(c) {
			continue
		}
		// A commander is the identity of the deck, so it comes from the
		// sets the reader named and never from the mana fill (D-382).
		if !cards.InSets(c, setCodes) {
			continue
		}
		// A commander must hold every color the user named, and no other
		// (D-148). The 99 keeps the subset test, because a mono-red card
		// belongs in a blue-red deck.
		if colorSet != nil &&
			(!IdentityFits(c.ColorIdentity, colorSet) || !identityCovers(c.ColorIdentity, colorSet)) {
			continue
		}
		if c.GameChanger && req.Bracket > 0 && req.Bracket <= 2 {
			continue
		}
		themeScore, signals := theme.score(c)
		if themeScore <= 0 {
			continue
		}
		owned := req.Owned[c.OracleId]
		if mode == mtgv1.PoolRule_POOL_RULE_OWNED_ONLY && owned == 0 {
			continue
		}
		out = append(out, Candidate{
			Card: c, Role: mtgv1.CardRole_CARD_ROLE_THREAT,
			Score: themeScore*0.7 + popularity(c, maxRank)*0.3,
			Owned: owned, Signals: signals,
		})
	}
	sortCandidates(out)
	// A pair carries the union of two color identities. It is the only
	// way to reach four colors, and it is what a user means when they ask
	// for partners or a Background (D-154). The pool offers pairs when
	// the user asked for one, and when too few single commanders fit the
	// colors to fill the three names the pick row holds.
	if req.WantPair || len(out) < commanderNames {
		out = append(out, b.commanderPairs(idx, req, theme, colorSet, setCodes, mode, maxRank)...)
		sortCandidates(out)
	}
	// A theme the tag table does not know leaves the pool nearly empty.
	// "Hobbit" names a handful of legends, and a reader who refuses those
	// has nothing left to be offered: the row falls silent and D-127
	// hands the choice to the agent with no word to the reader (D-367).
	//
	// The theme still leads. Below the floor the pool takes the
	// commanders that fit the format and the colors on popularity alone,
	// so "name three more" always has three more.
	if len(out) < CommanderPoolFloor {
		out = append(out, b.unthemed(idx, req, colorSet, setCodes, mode, maxRank, out)...)
	}
	// A bracket 4 or 5 request wants the strongest commander, not the
	// most popular one (OQ-48). The cEDH signal of the quality model
	// leads, and the theme order above breaks the tie (PR-14B).
	if req.Bracket >= 4 && req.CommanderSignal != nil {
		signal := func(c Candidate) float64 {
			if c.Partner != nil {
				return req.CommanderSignal(c.Card.GetOracleId(), c.Partner.GetOracleId())
			}
			return req.CommanderSignal(c.Card.GetOracleId())
		}
		sort.SliceStable(out, func(i, j int) bool { return signal(out[i]) > signal(out[j]) })
	}
	// Owned-first ranks on quality like any card. The commander is one
	// card, the buy list carries it, and a deck led by the best fit beats
	// a deck led by a legend the user happens to own (D-297). Owned-only
	// filtered above, because there the commander must be owned.
	return out, nil
}

// CommanderPoolFloor is the pool size under which the theme filter has
// left too little to choose from. The pick row names three at a time, so
// a reader who refuses twice needs nine, and a little room over that.
const CommanderPoolFloor = 12

// unthemed ranks the commanders that fit the format and the colors but
// carry no theme signal, best first on popularity (D-367). They go after
// every themed commander, so the theme still leads.
func (b *Builder) unthemed(idx *cards.Index, req Request, colorSet map[mtgv1.Color]bool,
	setCodes map[string]bool, mode mtgv1.PoolRule, maxRank float64, have []Candidate,
) []Candidate {
	seen := make(map[string]bool, len(have))
	for _, c := range have {
		seen[c.Card.GetOracleId()] = true
	}
	// The fill drops an excluded commander as the themed half does (D-408).
	for _, id := range req.ExcludeOracleIDs {
		seen[id] = true
	}
	var out []Candidate
	for _, c := range idx.All() {
		if seen[c.GetOracleId()] {
			continue
		}
		if !c.GetCanBeCommander() || !legalIn(c, legalKeys[mtgv1.FormatId_FORMAT_ID_COMMANDER]) {
			continue
		}
		if !hasPaperPrinting(c) || !cards.InSets(c, setCodes) {
			continue
		}
		// The same color rule as the themed half: a commander holds every
		// color the user named, and no other (D-148).
		if colorSet != nil &&
			(!IdentityFits(c.GetColorIdentity(), colorSet) || !identityCovers(c.GetColorIdentity(), colorSet)) {
			continue
		}
		if c.GetGameChanger() && req.Bracket > 0 && req.Bracket <= 2 {
			continue
		}
		owned := req.Owned[c.GetOracleId()]
		if mode == mtgv1.PoolRule_POOL_RULE_OWNED_ONLY && owned == 0 {
			continue
		}
		out = append(out, Candidate{
			Card: c, Role: mtgv1.CardRole_CARD_ROLE_THREAT,
			Score: popularity(c, maxRank),
			Owned: owned, Signals: []string{"no theme signal"},
		})
	}
	sortCandidates(out)
	return out
}

// commanderNames is how many names the pick row holds (catalog row
// commander_pick). A pool shorter than this can not fill the question.
const commanderNames = 3

// commanderPairs builds the two-commander candidates whose combined color
// identity matches the request.
//
// Few leaders can pair at all, so the search space is small enough to
// walk whole (D-154).
//
// The pairing rules live in internal/rules, which the deck validator
// already uses. One term per concept: a pair this offers is a pair that
// passes validation.
func (b *Builder) commanderPairs(idx *cards.Index, req Request, theme ThemeMatch,
	colorSet map[mtgv1.Color]bool, setCodes map[string]bool, mode mtgv1.PoolRule, maxRank float64) []Candidate {
	excluded := map[string]bool{}
	for _, id := range req.CommanderOracleIDs {
		excluded[id] = true
	}
	var pairable []*mtgv1.Card
	for _, c := range idx.All() {
		// An excluded commander is out of the pairs, as it is out of the 99.
		if excluded[c.OracleId] || !legalIn(c, legalKeys[mtgv1.FormatId_FORMAT_ID_COMMANDER]) {
			continue
		}
		// A card that can not pair never reaches the walk, which keeps the
		// pair loop small (D-154).
		if !canPair(c) || !hasPaperPrinting(c) || !cards.InSets(c, setCodes) {
			continue
		}
		if !c.GetCanBeCommander() && !c.GetIsBackground() {
			continue
		}
		if c.GameChanger && req.Bracket > 0 && req.Bracket <= 2 {
			continue
		}
		if mode == mtgv1.PoolRule_POOL_RULE_OWNED_ONLY && req.Owned[c.OracleId] == 0 {
			continue
		}
		pairable = append(pairable, c)
	}
	var out []Candidate
	for i, a := range pairable {
		for _, c := range pairable[i+1:] {
			if !rules.ValidPair(a, c) {
				continue
			}
			// The union of the two identities is the deck's identity.
			union := append(append([]mtgv1.Color(nil), a.ColorIdentity...), c.ColorIdentity...)
			if colorSet != nil && (!IdentityFits(union, colorSet) || !identityCovers(union, colorSet)) {
				continue
			}
			// The user asked for a Background by name, so a pair without
			// one answers a different question.
			if req.WantBackground && !a.GetIsBackground() && !c.GetIsBackground() {
				continue
			}
			// One of the two must carry the theme. A Background rarely
			// does, and the leader is what the deck is built around.
			aScore, aSignals := theme.score(a)
			cScore, cSignals := theme.score(c)
			if aScore <= 0 && cScore <= 0 {
				continue
			}
			lead, second, signals := a, c, aSignals
			if cScore > aScore {
				lead, second, signals = c, a, cSignals
			}
			// A Background never leads: it is not a commander on its own
			// (corpus section 2.2).
			if lead.GetIsBackground() {
				lead, second = second, lead
			}
			best := aScore
			if cScore > best {
				best = cScore
			}
			owned := req.Owned[lead.OracleId]
			if o := req.Owned[second.OracleId]; o < owned {
				owned = o
			}
			out = append(out, Candidate{
				Card: lead, Partner: second, Role: mtgv1.CardRole_CARD_ROLE_THREAT,
				Score: best*0.7 + popularity(lead, maxRank)*0.3,
				Owned: owned, Signals: signals,
			})
		}
	}
	sortCandidates(out)
	// A long tail of pairs would bury the singles. The pick row holds
	// three names, so a few good pairs are enough to fill it.
	if len(out) > 2*commanderNames {
		out = out[:2*commanderNames]
	}
	return out
}

// canPair reports whether a card can be half of a two-commander pair.
// Few leaders can, so this cut is what keeps the pair walk small (D-154).
//
// A Doctor carries no partner kind: the Doctor's companion card carries
// it (derive.go). The Doctor passes on its creature types instead, which
// is the test rules.ValidPair applies (CR 702.124m).
func canPair(c *mtgv1.Card) bool {
	if c.GetIsBackground() || rules.IsDoctor(c) {
		return true
	}
	switch c.Partner {
	case mtgv1.PartnerKind_PARTNER_KIND_PARTNER,
		mtgv1.PartnerKind_PARTNER_KIND_WITH,
		mtgv1.PartnerKind_PARTNER_KIND_FRIENDS_FOREVER,
		mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND,
		mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION:
		return true
	default:
		return false
	}
}

// DisplayName names a candidate the way a question does. A pair reads
// "A + B", which is how the pick row offers it (D-154).
func (c Candidate) DisplayName() string {
	if c.Partner == nil {
		return c.Card.GetName()
	}
	return c.Card.GetName() + " + " + c.Partner.GetName()
}

// CountInSets counts the distinct nonbasic cards a set-limited request
// can draw on: format-legal, on paper, inside the color identity, and
// printed in one of the named sets. It is the number the viability floor
// of D-380 reads.
//
// It walks the index and scores nothing, so it costs a fraction of a
// Build. The floor runs twice per build, once before the commander is
// chosen and once after, because the commander narrows the colors.
//
// Basic lands are out of the count. They repeat without limit and they
// are never filtered by a set (D-378), so counting them would say a
// five-card set can build a deck.
func CountInSets(idx *cards.Index, req Request) int {
	if idx == nil || len(req.SetCodes) == 0 {
		return 0
	}
	setCodes := cards.CodeSet(req.SetCodes)
	colorSet := colorSetOf(req.Colors)
	legalKey := legalKeys[req.Format]
	n := 0
	for _, c := range idx.All() {
		if IsBasicLand(c) || !hasPaperPrinting(c) {
			continue
		}
		if legalKey != "" && !legalIn(c, legalKey) {
			continue
		}
		if colorSet != nil && !IdentityFits(c.GetColorIdentity(), colorSet) {
			continue
		}
		if !cards.InSets(c, setCodes) {
			continue
		}
		n++
	}
	return n
}

// SetFloor is the distinct nonbasic count under which a set family can
// not build a deck of this format (D-380).
//
// A Commander deck holds 99 cards beside the commander. About 36 of them
// are lands, and basic lands fill most of that, so about 70 distinct
// nonbasic cards is the real need. Measured on the 2026-08-31 snapshot,
// the Hobbit family clears it in every two-color identity (114 to 132)
// and in mono white, red, and green, and falls under it in mono black
// (68). Of 287 playable-product families, 163 clear it.
//
// A 60-card deck runs four copies of a name, so half the count builds it.
func SetFloor(format mtgv1.FormatId) int {
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		return 70
	}
	return 35
}

// CountManaInSets counts the ramp cards and the nonbasic lands a set
// family offers in the deck's colors (D-382). The mana row compares it
// with the role targets of the format.
//
// It assigns a role, so it costs more than CountInSets and less than a
// whole Build: nothing is scored, sorted, or capped.
func (b *Builder) CountManaInSets(idx *cards.Index, req Request) int {
	if idx == nil || len(req.SetCodes) == 0 {
		return 0
	}
	setCodes := cards.CodeSet(req.SetCodes)
	colorSet := colorSetOf(req.Colors)
	legalKey := legalKeys[req.Format]
	roleTags := b.themes.roleSets(idx.Tags())
	useText := idx.Tags().Len() == 0
	n := 0
	for _, c := range idx.All() {
		if IsBasicLand(c) || !hasPaperPrinting(c) {
			continue
		}
		if legalKey != "" && !legalIn(c, legalKey) {
			continue
		}
		if colorSet != nil && !IdentityFits(c.GetColorIdentity(), colorSet) {
			continue
		}
		if !cards.InSets(c, setCodes) {
			continue
		}
		switch role, _ := assignRole(c, roleTags, false, useText); role {
		case mtgv1.CardRole_CARD_ROLE_RAMP, mtgv1.CardRole_CARD_ROLE_LAND:
			n++
		}
	}
	return n
}
