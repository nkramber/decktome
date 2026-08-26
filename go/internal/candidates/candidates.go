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

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
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
	Bracket int32
	// MetaBoost gives a meta score in [0,1] per Oracle id (PR-14). Nil today.
	MetaBoost func(oracleID string) float64
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
// about 50 upgrades.
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

// Candidate is one shortlisted card with its evidence.
type Candidate struct {
	Card  *mtgv1.Card
	Role  mtgv1.CardRole
	Score float64
	// Owned is the owned count, 0 without a collection.
	Owned int32
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

// Stats counts the funnel for M-3 style reporting and for the gate doc.
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

// legalKeys maps a format to its Scryfall legality key. House rules and
// an unknown format skip the legality filter.
var legalKeys = map[mtgv1.FormatId]string{
	mtgv1.FormatId_FORMAT_ID_COMMANDER: "commander",
	mtgv1.FormatId_FORMAT_ID_STANDARD:  "standard",
	mtgv1.FormatId_FORMAT_ID_PIONEER:   "pioneer",
	mtgv1.FormatId_FORMAT_ID_MODERN:    "modern",
	mtgv1.FormatId_FORMAT_ID_LEGACY:    "legacy",
	mtgv1.FormatId_FORMAT_ID_VINTAGE:   "vintage",
	mtgv1.FormatId_FORMAT_ID_PAUPER:    "pauper",
}

// Build makes the shortlist. It is deterministic: the same index and
// request give the same list. PR-9 adds the seeded shuffle on top.
func (b *Builder) Build(idx *cards.Index, req Request) (*List, error) {
	if idx == nil {
		return nil, fmt.Errorf("candidates: no card index")
	}
	lim := req.Limits.withDefaults()
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
	theme := b.themes.match(req.Theme, idx.Tags())
	roleTags := b.themes.roleSets(idx.Tags())
	legalKey := legalKeys[req.Format]
	colorSet := colorSetOf(req.Colors)
	excluded := map[string]bool{}
	for _, id := range req.CommanderOracleIDs {
		excluded[id] = true
	}
	maxRank := maxRankOf(idx)

	var stats Stats
	var scored []Candidate
	for _, c := range idx.All() {
		if excluded[c.OracleId] || isBasicLand(c) {
			continue
		}
		if legalKey != "" && !legalIn(c, legalKey) {
			continue
		}
		if colorSet != nil && !identityFits(c.ColorIdentity, colorSet) {
			continue
		}
		if c.GameChanger && req.Bracket > 0 && req.Bracket <= 2 {
			continue
		}
		stats.Pool++
		themeScore, signals := theme.score(c)
		role, roleSignal := assignRole(c, roleTags, themeScore > 0)
		if roleSignal != "" {
			signals = append(signals, roleSignal)
		}
		if themeScore > 0 {
			stats.OnTheme++
		}
		// A card with no theme signal stays only when it fills a staple
		// role (lands, ramp, draw, removal, wipes, interaction). Threats
		// and synergy pieces need a theme signal.
		if themeScore == 0 && !stapleRole(role) {
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
		scored = append(scored, Candidate{Card: c, Role: role, Score: score, Owned: owned, Signals: signals})
	}
	sortCandidates(scored)

	var main, upgrades []Candidate
	switch mode {
	case mtgv1.PoolRule_POOL_RULE_ANY_CARD:
		main = capByRole(scored, lim)
	case mtgv1.PoolRule_POOL_RULE_OWNED_ONLY:
		main = capByRole(filterOwned(scored, true), lim)
	case mtgv1.PoolRule_POOL_RULE_OWNED_FIRST:
		main = capByRole(filterOwned(scored, true), lim)
		upgrades = topUpgrades(filterOwned(scored, false), main, lim.Upgrades)
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
func capByRole(in []Candidate, lim Limits) []Candidate {
	byRole := map[mtgv1.CardRole][]Candidate{}
	for _, c := range in {
		byRole[c.Role] = append(byRole[c.Role], c)
	}
	var out []Candidate
	for _, r := range roleOrder {
		cs := byRole[r]
		if n := lim.PerRole[r]; n > 0 && len(cs) > n {
			cs = cs[:n]
		}
		out = append(out, cs...)
	}
	if len(out) > lim.Total {
		out = out[:lim.Total]
	}
	return out
}

// topUpgrades picks the best unowned cards that beat the weakest owned
// card of the same role, so an upgrade is a real improvement.
func topUpgrades(unowned, owned []Candidate, limit int) []Candidate {
	floor := map[mtgv1.CardRole]float64{}
	for _, c := range owned {
		if v, ok := floor[c.Role]; !ok || c.Score < v {
			floor[c.Role] = c.Score
		}
	}
	var out []Candidate
	for _, c := range unowned {
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

func identityFits(identity []mtgv1.Color, allowed map[mtgv1.Color]bool) bool {
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

func isBasicLand(c *mtgv1.Card) bool {
	return slices.Contains(c.Supertypes, "Basic") && slices.Contains(c.CardTypes, "Land")
}

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

// ThemeColors returns the colors a theme is strongest in, in WUBRG order.
// It tallies the color identity of the theme's best cards and keeps a
// color that carries at least share of them. PR-7 uses it to fill the
// {colors} clause of the color question, so the agent states a fact
// instead of asking the user for it.
func (b *Builder) ThemeColors(idx *cards.Index, format mtgv1.FormatId, theme string, top int, share float64) ([]mtgv1.Color, error) {
	if top <= 0 {
		top = 100
	}
	if share <= 0 {
		share = 0.25
	}
	list, err := b.Build(idx, Request{Format: format, Theme: theme})
	if err != nil {
		return nil, err
	}
	count := map[mtgv1.Color]int{}
	seen := 0
	for _, c := range list.Candidates {
		if seen >= top {
			break
		}
		if len(c.Card.ColorIdentity) == 0 {
			continue // colorless cards say nothing about a theme's colors
		}
		seen++
		for _, col := range c.Card.ColorIdentity {
			count[col]++
		}
	}
	if seen == 0 {
		return nil, nil
	}
	var out []mtgv1.Color
	for _, col := range []mtgv1.Color{mtgv1.Color_COLOR_W, mtgv1.Color_COLOR_U, mtgv1.Color_COLOR_B, mtgv1.Color_COLOR_R, mtgv1.Color_COLOR_G} {
		if float64(count[col])/float64(seen) >= share {
			out = append(out, col)
		}
	}
	return out, nil
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
// with lands first, so it is not ordered by score at all. Reading the
// first legends out of it returned whichever ones landed in the land
// bucket: a blink request was answered with three Ojer modal double-faced
// cards, scoring 0.16 on the theme, while 85 on-theme blink commanders
// existed and Emiel the Blessed scored 0.56 (measured 2026-08-25, D-94).
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
	theme := b.themes.match(req.Theme, idx.Tags())
	colorSet := colorSetOf(req.Colors)
	maxRank := maxRankOf(idx)

	var out []Candidate
	for _, c := range idx.All() {
		if !c.GetCanBeCommander() || !legalIn(c, legalKeys[mtgv1.FormatId_FORMAT_ID_COMMANDER]) {
			continue
		}
		if colorSet != nil && !identityFits(c.ColorIdentity, colorSet) {
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
	// Owned-first offers what the user already has, before a card they
	// would need to buy.
	if mode == mtgv1.PoolRule_POOL_RULE_OWNED_FIRST {
		out = append(filterOwned(out, true), filterOwned(out, false)...)
	}
	return out, nil
}
