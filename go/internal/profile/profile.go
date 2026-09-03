// Package profile is the bracket profile of a built deck (roadmap
// PR-14A, D-451 to D-453).
//
// A bracket reaches the generator as one line of prose, and no check
// read power after the build. The profile is the specification in data:
// a feature vector per deck, a band per feature per bracket, a goldfish
// simulation, and the content rules of the bracket read from Commander
// Spellbook. An off-band feature is a finding, and a finding buys the
// repair turn (D-244).
//
// The package is a library. It reads a deck and a card source, and it
// writes the profile and the findings. It never blocks a deck: the
// bands are guide numbers the gate tunes, and the content flags are
// community data (F-5). Every finding is a warning.
package profile

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
	"github.com/nkramber/mtg-deck-builder/go/internal/spellbook"
)

// The feature keys. bands.json is keyed by these.
const (
	KeyLand                = "land"
	KeyRamp                = "ramp"
	KeyDraw                = "draw"
	KeyRemoval             = "removal"
	KeyWipe                = "wipe"
	KeyInteraction         = "interaction"
	KeyAvgManaValue        = "avg_mana_value"
	KeyTutor               = "tutor"
	KeyFastMana            = "fast_mana"
	KeyTappedLand          = "tapped_land"
	KeyColorlessLand       = "colorless_land"
	KeyColorSources        = "color_sources"
	KeyGameChanger         = "game_changer"
	KeyManaTurnFour        = "mana_turn_four"
	KeyHandsTwoToFourLands = "hands_two_to_four_lands"
	KeyCommanderTurnOverMV = "commander_turn_over_mv"
)

// The finding codes. Every one is a warning that buys the repair turn.
const (
	// CodeOffBand reports a feature outside its band.
	CodeOffBand = "profile_off_band"
	// CodeMassLandDenial reports mass land denial in a bracket that
	// forbids it.
	CodeMassLandDenial = "mass_land_denial"
	// CodeExtraTurns reports more extra-turn cards than the bracket
	// allows, or an extra-turn combo.
	CodeExtraTurns = "extra_turn_limit"
	// CodeTwoCardCombo reports a two-card infinite combo faster than the
	// bracket allows.
	CodeTwoCardCombo = "two_card_combo"
	// CodeContentUnchecked reports that the content check did not run,
	// so the deck carries no content finding. It is information.
	CodeContentUnchecked = "content_unchecked"
)

// Classifier answers the content flags of a deck. The Commander
// Spellbook client is the one implementation, and a test fakes it.
type Classifier interface {
	EstimateBracket(ctx context.Context, commanders, main []string) (*spellbook.Result, error)
}

// TagSource answers the tag index of the current snapshot, or nil when
// the snapshot has none. The API's index swaps on refresh, so this is a
// function and not a value.
type TagSource func() *cards.TagIndex

// Profiler reads decks.
type Profiler struct {
	rules    *rules.Config
	bands    *Bands
	tags     TagSource
	classify Classifier
	hands    int
	seed     uint64
}

// New makes a profiler. tags and classify can be nil: without tags the
// tutor feature is not measured, and without a classifier the content
// check reports itself unchecked.
func New(cfg *rules.Config, tags TagSource, classify Classifier) (*Profiler, error) {
	bands, err := LoadBands()
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, fmt.Errorf("profile: the rules config is nil")
	}
	return &Profiler{rules: cfg, bands: bands, tags: tags, classify: classify, hands: DefaultHands, seed: 1}, nil
}

// Bands returns the band table, for the prompt and the reports.
func (p *Profiler) Bands() *Bands { return p.bands }

// SetHands sets how many games the simulation deals. Tests lower it.
func (p *Profiler) SetHands(n int) { p.hands = n }

// entry is one deck card the profiler resolved.
type entry struct {
	card  *mtgv1.Card
	count int
	role  mtgv1.CardRole
}

// Read builds the profile of a deck and the findings it earns. The deck
// is read as it is, the sideboard left out. A card the source does not
// know is skipped, and the rules engine has already reported it.
func (p *Profiler) Read(ctx context.Context, deck *mtgv1.Deck, src rules.CardSource) (*mtgv1.DeckProfile, []*mtgv1.Finding) {
	out, findings, entries, commanders := p.measure(deck, src)
	if deck.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		out.Content, findings = p.content(ctx, entries, commanders, out.GetBracket(), findings)
	}
	return out, findings
}

// Measure builds the profile with no content check and no finding. The
// quality fit of PR-14B reads thousands of published lists through it,
// and the endpoint's rate would make that a day's work (D-459).
func (p *Profiler) Measure(deck *mtgv1.Deck, src rules.CardSource) *mtgv1.DeckProfile {
	out, _, _, _ := p.measure(deck, src)
	return out
}

// measure resolves the deck and measures every feature.
func (p *Profiler) measure(deck *mtgv1.Deck, src rules.CardSource) (*mtgv1.DeckProfile, []*mtgv1.Finding, []entry, []*mtgv1.Card) {
	format := deck.GetFormat().GetId()
	commander := format == mtgv1.FormatId_FORMAT_ID_COMMANDER
	table, bracket := p.bands.For(format, deck.GetPower())
	out := &mtgv1.DeckProfile{Bracket: bracket, BandsVerifiedAt: p.bands.VerifiedAt}

	var entries []entry
	var commanders []*mtgv1.Card
	for _, dc := range deck.GetCards() {
		c, ok := src.ByOracleID(dc.GetOracleId())
		if !ok || dc.GetCount() <= 0 {
			continue
		}
		entries = append(entries, entry{card: c, count: int(dc.GetCount()), role: dc.GetRole()})
	}
	for _, id := range deck.GetCommanderOracleIds() {
		if c, ok := src.ByOracleID(id); ok {
			commanders = append(commanders, c)
		}
	}
	colors := deckColors(entries, commanders, commander)

	var tags *cards.TagIndex
	if p.tags != nil {
		tags = p.tags()
	}
	f := &features{}
	f.counts(entries, commanders)
	f.tutors(entries, tags)
	f.sources(entries, colors, deckSize(format))
	f.goldfish(entries, commanders, commander, p.hands, p.seed)

	var findings []*mtgv1.Finding
	for _, key := range featureOrder {
		m, ok := f.values[key]
		if !ok {
			continue
		}
		if !commander && key == KeyCommanderTurnOverMV {
			continue
		}
		row := &mtgv1.ProfileFeature{Key: key, Value: m.value, Note: m.note}
		if band, ok := table[key]; ok {
			row.Low = band.Low
			if band.High != nil {
				row.High = *band.High
				row.HasHigh = true
			}
			row.OffBand = !band.Holds(m.value)
		}
		if key == KeyGameChanger && commander {
			// The rules engine reports the Game Changer limit as a block,
			// so the profile marks the band and adds no second finding.
			if br, ok := p.rules.Brackets[bracket]; ok && br.MaxGameChangers >= 0 {
				row.Low, row.High, row.HasHigh = 0, float64(br.MaxGameChangers), true
				row.OffBand = m.value > float64(br.MaxGameChangers)
			}
		} else if row.OffBand {
			findings = append(findings, &mtgv1.Finding{
				Code: CodeOffBand, Severity: mtgv1.Severity_SEVERITY_WARN,
				Message: offBandMessage(row, bracket, format, deck.GetPower()),
			})
		}
		out.Features = append(out.Features, row)
	}
	out.Goldfish = f.sim
	return out, findings, entries, commanders
}

// featureOrder is the row order of a profile, the way a reader scans
// a deck: the mana base first, then the jobs, then the power signals,
// then the simulation.
var featureOrder = []string{
	KeyLand, KeyTappedLand, KeyColorlessLand, KeyColorSources, KeyAvgManaValue,
	KeyRamp, KeyDraw, KeyRemoval, KeyWipe, KeyInteraction,
	KeyTutor, KeyFastMana, KeyGameChanger,
	KeyManaTurnFour, KeyHandsTwoToFourLands, KeyCommanderTurnOverMV,
}

// measure is one feature value with the note the reader sees.
type measure struct {
	value float64
	note  string
}

type features struct {
	values map[string]measure
	sim    *mtgv1.Goldfish
}

func (f *features) set(key string, v float64, note string) {
	if f.values == nil {
		f.values = map[string]measure{}
	}
	f.values[key] = measure{value: v, note: note}
}

// counts measures the lands, the roles, the curve, the fast mana, and
// the Game Changers.
func (f *features) counts(entries []entry, commanders []*mtgv1.Card) {
	var lands, tapped, colorless, nonlands, fast, changers int
	var mvSum float64
	roles := map[mtgv1.CardRole]int{}
	var tappedNames, colorlessNames, fastNames, changerNames []string
	for _, e := range entries {
		c := e.card
		if isLand(c) {
			lands += e.count
			if entersTapped(c) {
				tapped += e.count
				tappedNames = append(tappedNames, c.GetName())
			}
			if isColorlessLand(c) {
				colorless += e.count
				colorlessNames = append(colorlessNames, c.GetName())
			}
		} else {
			nonlands += e.count
			mvSum += c.GetManaValue() * float64(e.count)
			if isFastMana(c) {
				fast += e.count
				fastNames = append(fastNames, c.GetName())
			}
		}
		roles[e.role] += e.count
		if c.GetGameChanger() {
			changers += e.count
			changerNames = append(changerNames, c.GetName())
		}
	}
	for _, c := range commanders {
		if c.GetGameChanger() {
			changers++
			changerNames = append(changerNames, c.GetName())
		}
	}
	f.set(KeyLand, float64(lands), "")
	f.set(KeyTappedLand, float64(tapped), names(tappedNames))
	f.set(KeyColorlessLand, float64(colorless), names(colorlessNames))
	avg := 0.0
	if nonlands > 0 {
		avg = math.Round(mvSum/float64(nonlands)*100) / 100
	}
	f.set(KeyAvgManaValue, avg, fmt.Sprintf("over %d nonland cards", nonlands))
	f.set(KeyRamp, float64(roles[mtgv1.CardRole_CARD_ROLE_RAMP]), "")
	f.set(KeyDraw, float64(roles[mtgv1.CardRole_CARD_ROLE_DRAW]), "")
	f.set(KeyRemoval, float64(roles[mtgv1.CardRole_CARD_ROLE_REMOVAL]), "")
	f.set(KeyWipe, float64(roles[mtgv1.CardRole_CARD_ROLE_WIPE]), "")
	f.set(KeyInteraction, float64(roles[mtgv1.CardRole_CARD_ROLE_INTERACTION]), "")
	f.set(KeyFastMana, float64(fast), names(fastNames))
	f.set(KeyGameChanger, float64(changers), names(changerNames))
}

// tutorSlug and tutorLandSlug are the Scryfall Tagger slugs. A tutor is
// any card under tutor, less the land searches, which are ramp and not
// the tutors the bracket text means. Slugs checked against the
// oracle_tags file of 2026-09-02.
const (
	tutorSlug     = "tutor"
	tutorLandSlug = "tutor-land"
)

// tutors measures the tutor count when the snapshot has tags. Without
// them the feature is absent, and the reader sees no row.
func (f *features) tutors(entries []entry, tags *cards.TagIndex) {
	if tags == nil {
		return
	}
	set := map[string]bool{}
	for _, id := range tags.Resolve(tutorSlug) {
		set[id] = true
	}
	for _, id := range tags.Resolve(tutorLandSlug) {
		delete(set, id)
	}
	n := 0
	var found []string
	for _, e := range entries {
		if set[e.card.GetOracleId()] && !isLand(e.card) {
			n += e.count
			found = append(found, e.card.GetName())
		}
	}
	f.set(KeyTutor, float64(n), names(found))
}

// sources measures the color sources against the pips by the Karsten
// tables. The value is the worst color's ratio of sources to
// requirement, and the note lists every color.
func (f *features) sources(entries []entry, colors []mtgv1.Color, size int) {
	if len(colors) == 0 {
		return
	}
	have := map[mtgv1.Color]float64{}
	needs := map[mtgv1.Color][]int{}
	for _, e := range entries {
		c := e.card
		if isLand(c) {
			if isFetch(c) {
				for _, col := range colors {
					have[col] += float64(e.count)
				}
				continue
			}
			for _, col := range c.GetProducedMana() {
				have[col] += float64(e.count)
			}
			continue
		}
		if producesMana(c) {
			for _, col := range c.GetProducedMana() {
				have[col] += RockSource * float64(e.count)
			}
		}
		pips := colorPips(c.GetManaCost())
		total := 0
		for _, n := range pips {
			total += n
		}
		generic := int(c.GetManaValue()) - total
		for col, n := range pips {
			need := sourcesNeeded(size, generic, n)
			for range e.count {
				needs[col] = append(needs[col], need)
			}
		}
	}
	worst := math.Inf(1)
	var parts []string
	for _, col := range colorOrder {
		if !hasColor(colors, col) {
			continue
		}
		req := requirement(needs[col])
		ratio := 1.0
		if req > 0 {
			ratio = have[col] / float64(req)
		}
		if ratio < worst {
			worst = ratio
		}
		parts = append(parts, fmt.Sprintf("%s %s of %d", colorLetters[col], num(have[col]), req))
	}
	if math.IsInf(worst, 1) {
		return
	}
	f.set(KeyColorSources, math.Round(worst*100)/100, "sources of requirement: "+strings.Join(parts, ", "))
}

// goldfish runs the simulation and records its three numbers.
func (f *features) goldfish(entries []entry, commanders []*mtgv1.Card, commander bool, hands int, seed uint64) {
	var list []simCard
	for _, e := range entries {
		sc := simCardOf(e.card)
		for range e.count {
			list = append(list, sc)
		}
	}
	mv := 0
	for i, c := range commanders {
		if v := int(c.GetManaValue()); i == 0 || v < mv {
			mv = v
		}
	}
	res := simulate(simInput{cards: list, commander: commander, commanderMV: mv, hands: hands, seed: seed})
	if res.hands == 0 {
		return
	}
	f.sim = &mtgv1.Goldfish{
		Hands:               int32(res.hands),
		CommanderTurn:       round2(res.commanderTurn),
		ManaTurnFour:        round2(res.manaTurnFour),
		ShareTwoToFourLands: round2(res.shareTwoToFour),
	}
	f.set(KeyManaTurnFour, f.sim.ManaTurnFour, fmt.Sprintf("mean over %d hands", res.hands))
	f.set(KeyHandsTwoToFourLands, f.sim.ShareTwoToFourLands, "share of first seven-card hands")
	if commander && len(commanders) > 0 {
		f.set(KeyCommanderTurnOverMV, round2(res.commanderTurn-float64(mv)),
			fmt.Sprintf("the commander costs %d and comes down on turn %s on average", mv, num(f.sim.CommanderTurn)))
	}
}

func round2(v float64) float64 { return math.Round(v*100) / 100 }

// content runs the content check and adds its findings.
func (p *Profiler) content(ctx context.Context, entries []entry, commanders []*mtgv1.Card,
	bracket int32, findings []*mtgv1.Finding) (*mtgv1.ContentCheck, []*mtgv1.Finding) {
	check := &mtgv1.ContentCheck{}
	unchecked := func(why string) (*mtgv1.ContentCheck, []*mtgv1.Finding) {
		check.Checked = false
		check.Error = why
		return check, append(findings, &mtgv1.Finding{
			Code: CodeContentUnchecked, Severity: mtgv1.Severity_SEVERITY_INFO,
			Message: "the content rules of the bracket were not checked: " + why,
		})
	}
	if p.classify == nil {
		return unchecked("no classifier is wired")
	}
	var names, cmdNames []string
	for _, e := range entries {
		if !isBasic(e.card) {
			names = append(names, e.card.GetName())
		}
	}
	for _, c := range commanders {
		cmdNames = append(cmdNames, c.GetName())
	}
	res, err := p.classify.EstimateBracket(ctx, cmdNames, names)
	if err != nil {
		return unchecked(err.Error())
	}
	check.Checked = true
	check.SourceTag = res.BracketTag
	var extraCards int
	for _, c := range res.Cards {
		if c.GameChanger {
			check.GameChangers = append(check.GameChangers, c.Card.Name)
		}
		if c.MassLandDenial {
			check.MassLandDenial = append(check.MassLandDenial, c.Card.Name)
		}
		if c.ExtraTurn {
			check.ExtraTurns = append(check.ExtraTurns, c.Card.Name)
			extraCards += max(c.Quantity, 1)
		}
	}
	var extraCombos, mldCombos []string
	var fastCombos []string
	br, known := p.rules.Brackets[bracket]
	for _, v := range res.Combos {
		// A near two-card combo, one that needs a common third piece
		// such as a lifegain trigger, reads one speed step slower than a
		// sure one. That is the endpoint's own rule: a sure combo at
		// speed 4 is Ruthless, and a near one at speed 4 is Spicy
		// (variant.py, read 2026-09-02).
		two := v.Relevant && v.DefinitelyTwoCard
		speed := v.Speed
		near := ""
		if !two && v.Relevant && v.ArguablyTwoCard {
			two = true
			speed--
			near = ", near two-card"
		}
		hit := &mtgv1.ComboHit{
			Id: v.Combo.ID, Cards: v.Combo.Names(), TwoCard: two,
			Speed: int32(v.Speed), ExtraTurn: v.ExtraTurn, MassLandDenial: v.MassLandDenial,
		}
		check.Combos = append(check.Combos, hit)
		line := strings.Join(hit.Cards, " + ")
		if v.ExtraTurn {
			extraCombos = append(extraCombos, line)
		}
		if v.MassLandDenial {
			mldCombos = append(mldCombos, line)
		}
		if known && two && br.MaxComboSpeed >= 0 && speed > br.MaxComboSpeed {
			fastCombos = append(fastCombos, fmt.Sprintf("%s (speed %d%s)", line, v.Speed, near))
		}
	}
	if !known {
		return check, findings
	}
	warn := func(code, msg string) {
		findings = append(findings, &mtgv1.Finding{Code: code, Severity: mtgv1.Severity_SEVERITY_WARN, Message: msg})
	}
	if !br.MassLandDenial {
		if all := append(append([]string(nil), check.MassLandDenial...), mldCombos...); len(all) > 0 {
			warn(CodeMassLandDenial, fmt.Sprintf("bracket %d allows no mass land denial, and the deck holds %s",
				bracket, strings.Join(all, ", ")))
		}
	}
	if br.MaxExtraTurnCards >= 0 {
		if extraCards > br.MaxExtraTurnCards {
			warn(CodeExtraTurns, fmt.Sprintf("bracket %d allows %s, and the deck holds %d: %s",
				bracket, plural(br.MaxExtraTurnCards, "extra-turn card"), extraCards, strings.Join(check.ExtraTurns, ", ")))
		}
		if len(extraCombos) > 0 {
			warn(CodeExtraTurns, fmt.Sprintf("bracket %d allows no extra-turn combo, and the deck holds %s",
				bracket, strings.Join(extraCombos, "; ")))
		}
	}
	if br.MaxComboSpeed >= 0 && len(fastCombos) > 0 {
		what := "no two-card infinite combo"
		if br.MaxComboSpeed > 0 {
			what = fmt.Sprintf("no two-card infinite combo that needs %s or less", speedWords[br.MaxComboSpeed+1])
		}
		warn(CodeTwoCardCombo, fmt.Sprintf("bracket %d allows %s, and the deck holds %s",
			bracket, what, strings.Join(fastCombos, "; ")))
	}
	return check, findings
}

// speedWords names the mana a combo of each speed needs, from the
// Spellbook scale: speed 5 no mana, 4 four or less, 3 six or less, 2
// eight or less.
var speedWords = map[int]string{5: "no mana", 4: "four mana", 3: "six mana", 2: "eight mana", 6: "no mana"}

// offBandMessage writes the finding the repair turn reads. It names
// the feature, the value, and the band, in the words of the bracket.
func offBandMessage(row *mtgv1.ProfileFeature, bracket int32, format mtgv1.FormatId, power *mtgv1.PowerLevel) string {
	who := fmt.Sprintf("bracket %d", bracket)
	if format != mtgv1.FormatId_FORMAT_ID_COMMANDER {
		who = "this power level"
		if step := power.GetSixtyStep(); step != mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED {
			who = "the " + strings.ToLower(strings.TrimPrefix(step.String(), "SIXTY_STEP_")) + " level"
		}
	}
	want := ""
	switch {
	case row.HasHigh && row.Low > 0:
		want = fmt.Sprintf("%s to %s", num(row.Low), num(row.High))
	case row.HasHigh:
		want = fmt.Sprintf("%s at most", num(row.High))
	default:
		want = fmt.Sprintf("%s or more", num(row.Low))
	}
	msg := fmt.Sprintf("%s is %s, and %s wants %s", featureWords[row.Key], num(row.Value), who, want)
	if row.Note != "" && row.Key != KeyAvgManaValue {
		msg += " (" + row.Note + ")"
	}
	return msg
}

// featureWords names each feature for a finding.
var featureWords = map[string]string{
	KeyLand:                "the land count",
	KeyRamp:                "the ramp count",
	KeyDraw:                "the draw count",
	KeyRemoval:             "the removal count",
	KeyWipe:                "the wipe count",
	KeyInteraction:         "the interaction count",
	KeyAvgManaValue:        "the average mana value of the nonland cards",
	KeyTutor:               "the tutor count",
	KeyFastMana:            "the fast mana count",
	KeyTappedLand:          "the count of lands that enter tapped",
	KeyColorlessLand:       "the count of nonbasic lands that make only colorless mana",
	KeyColorSources:        "the worst color's share of the sources it needs",
	KeyGameChanger:         "the Game Changer count",
	KeyManaTurnFour:        "the mana available on turn four",
	KeyHandsTwoToFourLands: "the share of opening hands with two to four lands",
	KeyCommanderTurnOverMV: "the turns the commander comes down after its mana value",
}

// Word names a feature for a report.
func Word(key string) string {
	if w, ok := featureWords[key]; ok {
		return w
	}
	return key
}

// deckColors is the color identity the sources are measured against:
// the commanders' identity in Commander, and the union of the cards'
// colors otherwise.
func deckColors(entries []entry, commanders []*mtgv1.Card, commander bool) []mtgv1.Color {
	set := map[mtgv1.Color]bool{}
	if commander && len(commanders) > 0 {
		for _, c := range commanders {
			for _, col := range c.GetColorIdentity() {
				set[col] = true
			}
		}
	} else {
		for _, e := range entries {
			for _, col := range e.card.GetColors() {
				set[col] = true
			}
		}
	}
	var out []mtgv1.Color
	for _, col := range colorOrder {
		if set[col] {
			out = append(out, col)
		}
	}
	return out
}

func hasColor(list []mtgv1.Color, want mtgv1.Color) bool {
	for _, c := range list {
		if c == want {
			return true
		}
	}
	return false
}

func deckSize(format mtgv1.FormatId) int {
	if format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		return 99
	}
	return 60
}

// names writes a sorted card list for a note, or nothing.
func names(list []string) string {
	if len(list) == 0 {
		return ""
	}
	sorted := append([]string(nil), list...)
	sort.Strings(sorted)
	return strings.Join(sorted, ", ")
}

func plural(n int, noun string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, noun)
	}
	return fmt.Sprintf("%d %ss", n, noun)
}

// CutShortlist asks the endpoint which cards of a shortlist the bracket
// forbids, before the model sees the list (D-468). It returns the names
// to drop: mass land denial through bracket 3, and an extra-turn card
// at bracket 1, by the rules of brackets.json. A call that fails drops
// nothing and returns the error, so the build goes on and the check
// after the build still runs. A 60-card format and a bracket the rules
// do not know drop nothing.
func (p *Profiler) CutShortlist(ctx context.Context, format mtgv1.FormatId, bracket int32, commanders, names []string) ([]string, error) {
	if format != mtgv1.FormatId_FORMAT_ID_COMMANDER || p.classify == nil || len(names) == 0 {
		return nil, nil
	}
	br, ok := p.rules.Brackets[bracket]
	if !ok || (br.MassLandDenial && br.MaxExtraTurnCards != 0) {
		return nil, nil
	}
	res, err := p.classify.EstimateBracket(ctx, commanders, names)
	if err != nil {
		return nil, err
	}
	var drop []string
	for _, c := range res.Cards {
		if (!br.MassLandDenial && c.MassLandDenial) || (br.MaxExtraTurnCards == 0 && c.ExtraTurn) {
			drop = append(drop, c.Card.Name)
		}
	}
	sort.Strings(drop)
	return drop, nil
}
