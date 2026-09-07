package questions

import (
	"context"
	"log/slog"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/precons"
)

// CandidateHints answers the placeholder values from the card index, so
// the color question states a fact and the commander question names real
// cards (PR-6 feeds PR-7). A zero value answers nothing, which drops the
// clause that needs a value.
type CandidateHints struct {
	Index   *cards.Index
	Builder *candidates.Builder
	Format  mtgv1.FormatId
	Colors  []mtgv1.Color
	Owned   map[string]int32
	Pool    mtgv1.PoolRule
	// SetCodes are the paper sets the reader limited the deck to, a whole
	// set family (D-376). The commander offer reads them (D-437).
	SetCodes []string
	// WantPair asks the candidate builder for two-commander pairs. The
	// user asked for partners or a Background (D-154).
	WantPair bool
	// WantBackground narrows that to pairs that hold a Background.
	WantBackground bool
	// Bracket is the Commander bracket the reader named, 0 before the
	// power row. CommanderSignal is the quality model's cEDH signal, nil
	// with no model. A bracket 4 or 5 offer ranks on it (PR-14B, OQ-48).
	Bracket         int32
	CommanderSignal func(oracleIDs ...string) float64
	// Precons is the precon table, nil before the meta job stored one.
	// The precon exclusion resolves the reader's words against it
	// (D-496).
	Precons *precons.Table
	// PrintingCounts reads the collection's copies per Scryfall id, for
	// the ownership check of D-408. It runs once per turn at most, and
	// only on a turn that needs it. Nil means no collection.
	PrintingCounts func() map[string]int32
	// OnThemeOwned is PR-6's count for the {n} clause of the thin-theme
	// question. The caller reads it from candidates.Stats.
	OnThemeOwned int
	Log          *slog.Logger

	printingsDone bool
	printings     map[string]int32

	commanders map[string][]string
	// The thin-theme count runs the whole PR-6 build, so it runs once per
	// key and the answer is kept. The key carries the format, the
	// colors, and the pool rule, as every other hint's does (D-82).
	thinDone  map[string]bool
	thin      map[string]bool
	thinCount map[string]int
	// The mana count of D-382 walks the index, so it runs once per key
	// and the answer is kept. The key carries the sets, the format, and
	// the colors.
	manaDone map[string]bool
	manaThin map[string]bool
	manaHave map[string]int
	manaWant map[string]int
}

// CanLead reports whether a named card can lead a deck. "Lightning Bolt
// as my commander" must not pass in silence (D-129).
//
// It answers "not known" for any legendary card it can not confirm. The
// engine reads a type line, and a type line is wrong about a whole class
// of commander. Grist, the Hunger Tide is "Legendary Planeswalker" and it
// is a legal commander, because a characteristic-defining ability makes
// it a creature card everywhere except the battlefield (D-269). A false
// "can not lead" passes the gate and the linter, so the engine says
// nothing it can not prove.
//
// A confident "no" therefore needs a card that is not legendary at all.
// Lightning Bolt is an instant and Sol Ring is not legendary, so both
// still answer (D-140).
func (h *CandidateHints) CanLead(name string) (canLead, known bool) {
	if h == nil || h.Index == nil {
		return false, false
	}
	card, ok := h.Index.ByName(strings.TrimSpace(name))
	if !ok {
		return false, false
	}
	if card.GetCanBeCommander() {
		return true, true
	}
	// A Background alone can not lead a deck. It joins a creature that
	// chooses a Background, and the pair is the commander (D-154).
	if card.GetIsBackground() {
		return false, true
	}
	if strings.Contains(strings.ToLower(card.GetTypeLine()), "legendary") {
		// A legendary card the engine can not confirm. Say nothing.
		return false, false
	}
	return false, true
}

// FitsColors reports whether the named card holds every color named and
// no other. It is the same test CommanderPool applies when it builds the
// pool (D-148), so a name that survives here would be offered again.
//
// An empty color list fits everything: the user has named no colors, so
// nothing is out of them (D-153).
func (h *CandidateHints) FitsColors(name string, colors []mtgv1.Color) (fits, known bool) {
	if h == nil || h.Index == nil {
		return false, false
	}
	card, ok := h.Index.ByName(strings.TrimSpace(name))
	if !ok {
		return false, false
	}
	if len(colors) == 0 {
		return true, true
	}
	return candidates.IdentityMatches(card.GetColorIdentity(), colors), true
}

// UseSlots takes the slot values as they stand inside the turn. The cache
// key carries the format, the colors, and the pool rule, so an answer
// computed under other values stays keyed to those values (D-82).
func (h *CandidateHints) UseSlots(format mtgv1.FormatId, colors []mtgv1.Color, pool mtgv1.PoolRule) {
	if h == nil {
		return
	}
	if format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		h.Format = format
	}
	if len(colors) > 0 {
		h.Colors = colors
	}
	if pool != mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		h.Pool = pool
	}
}

// UseSets takes the set limit as it stands inside the turn (D-437). A
// set the classifier resolved this turn reaches the offer of this turn.
func (h *CandidateHints) UseSets(codes []string) {
	if h == nil || len(codes) == 0 {
		return
	}
	h.SetCodes = codes
}

// UseWantPair records that the user asked for a two-commander pair. The
// pool offers pairs on its own when too few singles fit the colors, so
// this only adds the case the words ask for (D-154).
func (h *CandidateHints) UseWantPair(want, background bool) {
	if h == nil || !want {
		return
	}
	h.WantPair = true
	if background {
		h.WantBackground = true
	}
}

// Commanders names up to three commanders for the theme. It never names
// one the agent already offered, so a user who answers "none" sees three
// others (D-73).
func (h *CandidateHints) Commanders(theme string, skip []string) []string {
	// An empty theme is a pool of its own: the reader declined the theme
	// or named none, and the pool ranks the commanders that fit the
	// format, the colors, and the sets on popularity (D-367, D-437).
	// Before this, an empty theme answered no name, the pick row went
	// out bare, and the build chose a commander with no word to the
	// reader.
	if h == nil || h.Index == nil || h.Builder == nil {
		return nil
	}
	cacheKey := h.key(theme) + "\x00" + strings.Join(skip, "\x00") + "\x00" + strings.Join(h.SetCodes, ",")
	if h.Bracket >= 4 && h.CommanderSignal != nil {
		cacheKey += "\x00power"
	}
	if h.WantPair {
		cacheKey += "\x00pair"
	}
	if h.WantBackground {
		cacheKey += "\x00background"
	}
	if v, ok := h.commanders[cacheKey]; ok {
		return v
	}
	req := candidates.Request{
		Format:         mtgv1.FormatId_FORMAT_ID_COMMANDER,
		Theme:          theme,
		Colors:         h.Colors,
		PoolRule:       h.Pool,
		Owned:          h.Owned,
		WantPair:       h.WantPair,
		WantBackground: h.WantBackground,
		Bracket:        h.Bracket,
		// A bracket 4 or 5 request offers the strongest commanders, not
		// the most popular (PR-14B, OQ-48).
		CommanderSignal: h.CommanderSignal,
		// A commander comes from the sets the reader named (D-382). The
		// offer read no set before D-437, so a Hobbit-only request could
		// offer a commander of any set.
		SetCodes: h.SetCodes,
	}
	list, err := h.Builder.Commanders(h.Index, req, 3+len(skip))
	if err != nil {
		h.warn("commanders", err)
		return nil
	}
	var names []string
	for _, c := range list {
		// A pair reads "A + B". The pick row offers it as one choice,
		// because the user chooses a pair and not half of one (D-154).
		// hasName reads a short name and a full name as one card, which
		// is the sameCard rule of D-70.
		name := c.DisplayName()
		if hasName(skip, name) {
			continue
		}
		names = append(names, name)
		if len(names) == 3 {
			break
		}
	}
	if h.commanders == nil {
		h.commanders = map[string][]string{}
	}
	h.commanders[cacheKey] = names
	return names
}

// OwnedThemeCount is the on-theme owned count PR-6 reported. A count
// ThinTheme measured under the same key wins over the one the caller
// set, so the {n} clause of the thin-theme question reads the count that
// belongs to these colors (D-198, D-82).
func (h *CandidateHints) OwnedThemeCount(theme string) int {
	if h == nil {
		return 0
	}
	if n, ok := h.thinCount[h.key(theme)]; ok {
		return n
	}
	return h.OnThemeOwned
}

func (h *CandidateHints) warn(what string, err error) {
	log := h.Log
	if log == nil {
		log = slog.Default()
	}
	log.WarnContext(context.Background(), "questions: hint failed", "hint", what, "error", err)
}

// key is the cache key of one hint. It carries every value the answer
// depends on. A key of the theme alone keeps a colorless commander list
// after the user names their colors (D-82).
func (h *CandidateHints) key(theme string) string {
	var b strings.Builder
	b.WriteString(theme)
	b.WriteString("|")
	b.WriteString(h.Format.String())
	b.WriteString("|")
	for _, c := range h.Colors {
		b.WriteString(c.String())
	}
	b.WriteString("|")
	b.WriteString(h.Pool.String())
	return b.String()
}

// ThinTheme runs the PR-6 count and reports whether the collection holds
// fewer than 30 on-theme cards (D-63). The second value is the count, for
// the {n} clause of the thin-theme question.
//
// The mode is owned-first when the user has not answered the pool
// question, because that is the default with a collection (D-37).
//
// The answer is cached by key, so a count taken before the user named
// the colors is not served after. It leaves OnThemeOwned alone: the
// count reaches OwnedThemeCount through the same cache (D-82).
func (h *CandidateHints) ThinTheme(theme string) (bool, int) {
	if h == nil || strings.TrimSpace(theme) == "" {
		return false, 0
	}
	key := h.key(theme)
	if h.thinDone[key] {
		return h.thin[key], h.thinCount[key]
	}
	if h.Index == nil || h.Builder == nil || len(h.Owned) == 0 {
		return false, 0
	}
	pool := h.Pool
	if pool == mtgv1.PoolRule_POOL_RULE_UNSPECIFIED {
		pool = mtgv1.PoolRule_POOL_RULE_OWNED_FIRST
	}
	list, err := h.Builder.Build(h.Index, candidates.Request{
		Format:   h.Format,
		Theme:    theme,
		Colors:   h.Colors,
		PoolRule: pool,
		Owned:    h.Owned,
	})
	if err != nil {
		h.warn("thin theme", err)
		return false, 0
	}
	h.cacheThin(key, list.Stats.ThinTheme, list.Stats.OnThemeOwned)
	return h.thin[key], h.thinCount[key]
}

// cacheThin keeps one thin-theme answer under its key.
func (h *CandidateHints) cacheThin(key string, thin bool, count int) {
	if h.thinDone == nil {
		h.thinDone, h.thin, h.thinCount = map[string]bool{}, map[string]bool{}, map[string]int{}
	}
	h.thinDone[key], h.thin[key], h.thinCount[key] = true, thin, count
}

// ResolveSet maps the words a reader wrote onto a set family (D-376). It
// answers the SetResolver contract of the question workflow.
//
// ok is false when the phrase names no set this snapshot holds, or when
// it names two or more base sets. options then holds the names the
// question offers, and it is empty for an unknown name.
func (h *CandidateHints) ResolveSet(phrase string) (codes, names, options []string, ok bool) {
	if h == nil || h.Index == nil {
		return nil, nil, nil, false
	}
	tbl := h.Index.Sets()
	res := tbl.Resolve(phrase)
	switch res.Kind {
	case cards.ResolveOne:
		return res.Codes, tbl.Names(res.Codes), nil, true
	case cards.ResolveMany:
		for _, s := range res.Candidates {
			options = append(options, s.Name)
		}
		return nil, nil, options, false
	}
	return nil, nil, nil, false
}

// ResolveSetGroup maps a franchise word onto every family it names
// (D-525). It answers the SetResolver contract for a group request.
func (h *CandidateHints) ResolveSetGroup(phrase string) (codes, names []string, ok bool) {
	if h == nil || h.Index == nil {
		return nil, nil, false
	}
	tbl := h.Index.Sets()
	codes = tbl.ResolveGroup(phrase)
	if len(codes) == 0 {
		return nil, nil, false
	}
	return codes, tbl.Names(codes), true
}

// printingCounts reads the collection's copies per printing once.
func (h *CandidateHints) printingCounts() map[string]int32 {
	if h.printingsDone {
		return h.printings
	}
	h.printingsDone = true
	if h.PrintingCounts != nil {
		h.printings = h.PrintingCounts()
	}
	return h.printings
}

// ResolvePrecon maps a phrase onto the products of the precon table, and
// says whether the collection holds them whole (D-496, D-497). It answers
// the PreconResolver contract.
func (h *CandidateHints) ResolvePrecon(phrase string) PreconMatch {
	if h == nil || h.Precons == nil {
		return PreconMatch{}
	}
	m := h.Precons.Resolve(phrase)
	if !m.OK() {
		return PreconMatch{Options: m.Options}
	}
	res := PreconMatch{OK: true}
	for _, p := range m.Products {
		res.Products = append(res.Products, PreconRef{Key: p.Key, Name: p.Name})
	}
	if counts := h.printingCounts(); counts != nil {
		whole := false
		for _, p := range m.Products {
			if p.OwnedWhole(counts, candidates.BasicLandByOracle(h.Index)) {
				whole = true
				break
			}
		}
		if !whole {
			res.Partial = m.Names()
		}
	}
	return res
}

// OwnedPrecons lists the products the collection holds whole (D-408),
// basic lands aside (D-523). It answers the OwnedPreconSource contract.
func (h *CandidateHints) OwnedPrecons() ([]PreconRef, bool) {
	if h == nil || h.Precons == nil {
		return nil, false
	}
	counts := h.printingCounts()
	if counts == nil {
		return nil, false
	}
	var out []PreconRef
	for _, p := range h.Precons.Owned(counts, candidates.BasicLandByOracle(h.Index)) {
		out = append(out, PreconRef{Key: p.Key, Name: p.Name})
	}
	return out, true
}

// ThinSetMana counts the mana cards a set family offers against the count
// the deck wants (D-382). It answers the ManaSource contract.
//
// The mana roles are ramp and land. Basic lands are out of both counts:
// no set limit filters them, and they repeat without limit, so counting
// them would say every set has a mana base.
//
// want is zero when this source can not answer, and the row then drops
// the clause that names the two counts.
func (h *CandidateHints) ThinSetMana(codes []string, format mtgv1.FormatId,
	colors []mtgv1.Color, power *mtgv1.PowerLevel,
) (thin bool, have, want int) {
	if h == nil || h.Index == nil || h.Builder == nil || len(codes) == 0 {
		return false, 0, 0
	}
	if format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		format = h.Format
	}
	if format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return false, 0, 0
	}
	key := strings.Join(codes, ",") + "|" + format.String() + "|" + colorKey(colors)
	if h.manaDone[key] {
		return h.manaThin[key], h.manaHave[key], h.manaWant[key]
	}
	targets := generate.TargetsFor(format, power)
	want = targets["ramp"] + targets["land"]
	have = h.Builder.CountManaInSets(h.Index, candidates.Request{
		Format: format, Colors: colors, SetCodes: codes,
	})
	thin = have < want
	if h.manaDone == nil {
		h.manaDone, h.manaThin = map[string]bool{}, map[string]bool{}
		h.manaHave, h.manaWant = map[string]int{}, map[string]int{}
	}
	h.manaDone[key], h.manaThin[key] = true, thin
	h.manaHave[key], h.manaWant[key] = have, want
	return thin, have, want
}

// colorKey reads a color list as one cache key.
func colorKey(colors []mtgv1.Color) string {
	var b strings.Builder
	for _, c := range colors {
		b.WriteString(c.String())
	}
	return b.String()
}

// OnlyCommander reports whether a named card can only be played in
// Commander here (D-388). It answers the FormatChecker contract.
//
// A legendary creature is not proof of the format. Sheoldred, the
// Apocalypse leads a Commander deck and plays in Standard, so it names
// no format. Atraxa, Praetor's Voice leads a Commander deck and is legal
// in neither of the other two formats this app builds, so it does.
//
// known is false for a name the index does not hold, and the caller
// claims nothing about it.
func (h *CandidateHints) OnlyCommander(name string) (only, known bool) {
	if h == nil || h.Index == nil {
		return false, false
	}
	card, ok := h.Index.ByName(strings.TrimSpace(name))
	if !ok {
		return false, false
	}
	if !card.GetCanBeCommander() {
		return false, true
	}
	for _, key := range []string{"standard", "modern"} {
		switch card.GetLegalities()[key] {
		case mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL,
			mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED:
			return false, true
		}
	}
	return true, true
}
