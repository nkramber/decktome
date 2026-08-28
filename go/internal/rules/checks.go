package rules

import (
	"fmt"
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// mainDeckCount sums main-deck copies, the commanders included:
// a Commander deck is exactly 100 cards with the command zone.
// The companion is not part of the count: it sits in the sideboard
// (60-card formats) or outside the 100 (Commander).
func mainDeckCount(deck *mtgv1.Deck) int {
	n := len(deck.CommanderOracleIds)
	for _, c := range counted(deck.Cards) {
		n += int(c.Count)
	}
	return n
}

// counted drops every entry with a count under one. Such an entry is a
// bad_count finding, and no sum may read it: a count of -1 beside a
// 101-card list passed the size check before this filter existed.
func counted(list []*mtgv1.DeckCard) []*mtgv1.DeckCard {
	out := make([]*mtgv1.DeckCard, 0, len(list))
	for _, c := range list {
		if c.Count >= 1 {
			out = append(out, c)
		}
	}
	return out
}

// checkCounts blocks every main-deck or sideboard entry whose count is
// under one. A zero says nothing, and a negative count would shrink
// every sum the other checks make.
func checkCounts(res *mtgv1.ValidationResult, deck *mtgv1.Deck) {
	report := func(zone string, list []*mtgv1.DeckCard) {
		for _, c := range list {
			if c.Count < 1 {
				add(res, CodeBadCount, mtgv1.Severity_SEVERITY_BLOCK,
					fmt.Sprintf("%s entry %s has count %d, the count must be 1 or more", zone, c.Name, c.Count), c.OracleId)
			}
		}
	}
	report("main deck", deck.Cards)
	report("sideboard", deck.Sideboard)
}

func checkSize(res *mtgv1.ValidationResult, deck *mtgv1.Deck, fr FormatRules) {
	n := mainDeckCount(deck)
	switch {
	case fr.ExactSize > 0 && n != fr.ExactSize:
		add(res, CodeDeckSize, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("deck has %d cards, the format needs exactly %d (commander included)", n, fr.ExactSize), "")
	case fr.ExactSize == 0 && n < fr.MinSize:
		add(res, CodeDeckSize, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("deck has %d cards, the format needs at least %d", n, fr.MinSize), "")
	}
	var sb int
	for _, c := range counted(deck.Sideboard) {
		sb += int(c.Count)
	}
	if sb > fr.SideboardMax {
		add(res, CodeSideboardSize, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("sideboard has %d cards, the format allows %d", sb, fr.SideboardMax), "")
	}
}

func isBasic(c *mtgv1.Card) bool {
	return slices.Contains(c.Supertypes, "Basic")
}

// checkCopies enforces the per-name limit. Basic lands and any-number
// cards are exempt. A card-text limit (Seven Dwarves: 7, Nazgûl: 9)
// replaces the format limit.
func checkCopies(res *mtgv1.ValidationResult, in Input, fr FormatRules) {
	counts := map[string]int32{}
	names := map[string]string{}
	for _, dc := range allCards(in.Deck) {
		counts[dc.OracleId] += dc.Count
		names[dc.OracleId] = dc.Name
	}
	for oid, n := range counts {
		card, ok := in.Cards.ByOracleID(oid)
		if !ok {
			continue // unknown_card comes from checkLegality
		}
		if isBasic(card) || card.AnyCountInDeck {
			continue
		}
		limit := int32(fr.MaxCopies)
		if card.MaxCopiesOverride > 0 {
			limit = card.MaxCopiesOverride
		}
		if n > limit {
			add(res, CodeCopyLimit, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("%s: %d copies, the limit is %d", names[oid], n, limit), oid)
		}
	}
}

// allCards walks the main deck, the sideboard, the commanders, and the
// companion. A companion that already sits in the sideboard is not
// counted twice.
func allCards(deck *mtgv1.Deck) []*mtgv1.DeckCard {
	out := make([]*mtgv1.DeckCard, 0, len(deck.Cards)+len(deck.Sideboard)+len(deck.CommanderOracleIds)+1)
	out = append(out, counted(deck.Cards)...)
	out = append(out, counted(deck.Sideboard)...)
	for _, oid := range deck.CommanderOracleIds {
		out = append(out, &mtgv1.DeckCard{OracleId: oid, Count: 1})
	}
	if deck.CompanionOracleId != "" && !inSideboard(deck, deck.CompanionOracleId) {
		out = append(out, &mtgv1.DeckCard{OracleId: deck.CompanionOracleId, Count: 1})
	}
	return out
}

func inSideboard(deck *mtgv1.Deck, oid string) bool {
	for _, dc := range deck.Sideboard {
		if dc.OracleId == oid && dc.Count > 0 {
			return true
		}
	}
	return false
}

func checkLegality(res *mtgv1.ValidationResult, in Input, fr FormatRules) {
	seen := map[string]bool{}
	for _, dc := range allCards(in.Deck) {
		if seen[dc.OracleId] {
			continue
		}
		seen[dc.OracleId] = true
		card, ok := in.Cards.ByOracleID(dc.OracleId)
		if !ok {
			add(res, CodeUnknownCard, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("unknown card id %q (%s): not in the card database", dc.OracleId, dc.Name), dc.OracleId)
			continue
		}
		if fr.ScryfallKey == "" {
			continue
		}
		switch card.Legalities[fr.ScryfallKey] {
		// RESTRICTED still reads as legal. No format the app builds has a
		// restricted list since D-155, so this arm is about the card data
		// and not about a deck the app can produce.
		case mtgv1.LegalityStatus_LEGALITY_STATUS_LEGAL, mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED:
		case mtgv1.LegalityStatus_LEGALITY_STATUS_BANNED:
			add(res, CodeBannedCard, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("%s is banned in this format", card.Name), dc.OracleId)
		default:
			add(res, CodeNotLegal, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("%s is not legal in this format", card.Name), dc.OracleId)
		}
	}
}

func checkCommander(res *mtgv1.ValidationResult, in Input) {
	ids := in.Deck.CommanderOracleIds
	if len(ids) == 0 {
		add(res, CodeNoCommander, mtgv1.Severity_SEVERITY_BLOCK, "the deck has no commander", "")
		return
	}
	if len(ids) > 2 {
		add(res, CodeBadCommander, mtgv1.Severity_SEVERITY_BLOCK, "more than two commanders", "")
		return
	}
	cmdrs := make([]*mtgv1.Card, 0, 2)
	for _, oid := range ids {
		c, ok := in.Cards.ByOracleID(oid)
		if !ok {
			return // unknown_card already reported
		}
		cmdrs = append(cmdrs, c)
		if !c.CanBeCommander && !c.IsBackground {
			add(res, CodeBadCommander, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("%s can not be a commander", c.Name), c.OracleId)
		}
	}
	// A Background is a commander only beside a "choose a Background"
	// commander (CR 702.124k). Alone, or beside any other card, it is not.
	for i, c := range cmdrs {
		if !c.IsBackground {
			continue
		}
		if len(cmdrs) < 2 || cmdrs[1-i].Partner != mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND {
			add(res, CodeBadCommander, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("%s is a Background: it needs a commander with \"choose a Background\"", c.Name), c.OracleId)
			return
		}
	}
	if len(cmdrs) == 2 {
		if !ValidPair(cmdrs[0], cmdrs[1]) {
			add(res, CodeBadPartner, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("%s and %s are not a legal commander pair", cmdrs[0].Name, cmdrs[1].Name), "")
		}
	}
}

// ValidPair checks the two-commander mechanics (corpus section 2.2).
// Two Partner cards pair only when their variant text is equal
// (CR 702.124i, effective 2026-08-07): plain Partner with plain Partner,
// Survivors with Survivors, and so on. CR 702.124f is the rule that
// forbids a cross of two different partner abilities.
//
// It is exported so the candidate builder can offer a pair as one choice.
// A pair carries the union of two color identities, which is the only way
// to reach four colors: WUBR, WBRG, and UBRG hold one legal single
// commander each (D-154).
func ValidPair(a, b *mtgv1.Card) bool {
	pk := func(c *mtgv1.Card) mtgv1.PartnerKind { return c.Partner }
	switch {
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_PARTNER && pk(b) == mtgv1.PartnerKind_PARTNER_KIND_PARTNER:
		return a.PartnerText == b.PartnerText
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_WITH && pk(b) == mtgv1.PartnerKind_PARTNER_KIND_WITH:
		return a.PartnerWithName == b.Name && b.PartnerWithName == a.Name
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_FRIENDS_FOREVER && pk(b) == mtgv1.PartnerKind_PARTNER_KIND_FRIENDS_FOREVER:
		return true
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND && b.IsBackground:
		return true
	case pk(b) == mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND && a.IsBackground:
		return true
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION && IsDoctor(b):
		return true
	case pk(b) == mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION && IsDoctor(a):
		return true
	default:
		return false
	}
}

// IsDoctor applies CR 702.124m: the creature types are exactly Time Lord
// Doctor, with no other creature type. It is exported because a Doctor
// carries no partner kind of its own, so the candidate builder needs
// this test to let a Doctor into a pair.
func IsDoctor(c *mtgv1.Card) bool {
	if len(c.Subtypes) != 3 || !slices.Contains(c.CardTypes, "Creature") {
		return false
	}
	return slices.Contains(c.Subtypes, "Time") && slices.Contains(c.Subtypes, "Lord") &&
		slices.Contains(c.Subtypes, "Doctor")
}

// checkColorIdentity covers the 99 and the companion (the 101st card).
func checkColorIdentity(res *mtgv1.ValidationResult, in Input) {
	allowed := map[mtgv1.Color]bool{}
	for _, oid := range in.Deck.CommanderOracleIds {
		if c, ok := in.Cards.ByOracleID(oid); ok {
			for _, col := range c.ColorIdentity {
				allowed[col] = true
			}
		}
	}
	check := func(oid string) {
		card, ok := in.Cards.ByOracleID(oid)
		if !ok {
			return
		}
		for _, col := range card.ColorIdentity {
			if !allowed[col] {
				add(res, CodeOffColor, mtgv1.Severity_SEVERITY_BLOCK,
					fmt.Sprintf("%s is outside the commander color identity", card.Name), oid)
				return
			}
		}
	}
	for _, dc := range in.Deck.Cards {
		check(dc.OracleId)
	}
	if in.Deck.CompanionOracleId != "" {
		check(in.Deck.CompanionOracleId)
	}
}

// checkBracket counts Game Changers in the 99, the command zone, and the
// companion. A Game Changer commander counts as one of the three at
// Bracket 3, and can not play in Brackets 1 and 2. Verified 2026-08-26
// against the Wizards announcement (see brackets.json).
func checkBracket(cfg *Config, res *mtgv1.ValidationResult, in Input) {
	bracket := in.Deck.GetPower().GetBracket()
	if bracket == 0 {
		return
	}
	br, ok := cfg.Brackets[bracket]
	if !ok {
		// A bracket the data does not know is its own finding, and not a
		// Game Changer count.
		add(res, CodeUnknownBracket, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("unknown bracket %d", bracket), "")
		return
	}
	var changers []string
	for _, dc := range allCards(in.Deck) {
		if c, ok := in.Cards.ByOracleID(dc.OracleId); ok && c.GameChanger {
			changers = append(changers, c.Name)
		}
	}
	if br.MaxGameChangers >= 0 && len(changers) > br.MaxGameChangers {
		add(res, CodeGameChangers, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("bracket %d allows %d Game Changers, the deck has %d: %s",
				bracket, br.MaxGameChangers, len(changers), strings.Join(changers, ", ")), "")
	}
	if bracket <= 3 {
		add(res, CodeBracketProse, mtgv1.Severity_SEVERITY_INFO,
			fmt.Sprintf("bracket %d (%s, %s turns): the prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11). Bracket data verified %s",
				bracket, br.Name, br.ExpectedTurns, cfg.VerifiedAt["brackets.json"]), "")
	}
}

// checkCompanion enforces F-18: the companion flag, the sideboard seat,
// and the banned-as-companion list that Scryfall can not express.
// Legality, color identity, and copy limits of the companion run in the
// shared checks through allCards.
func checkCompanion(cfg *Config, res *mtgv1.ValidationResult, in Input, fr FormatRules) {
	oid := in.Deck.CompanionOracleId
	if oid == "" {
		return
	}
	card, ok := in.Cards.ByOracleID(oid)
	if !ok {
		add(res, CodeBadCompanion, mtgv1.Severity_SEVERITY_BLOCK, "companion id is not in the card database", oid)
		return
	}
	if !card.IsCompanion {
		add(res, CodeBadCompanion, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("%s is not a companion", card.Name), oid)
		return
	}
	if fr.ScryfallKey != "" && slices.Contains(cfg.BannedAsCompanion[card.Name], fr.ScryfallKey) {
		add(res, CodeCompanionBanned, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("%s is banned as a companion in this format (it can still sit in the deck or command it)", card.Name), oid)
		return
	}
	// In a 60-card format the companion is one of the sideboard cards
	// (CR 702.139a). In Commander it is the 101st card.
	if !fr.Commander && !inSideboard(in.Deck, oid) {
		add(res, CodeCompanionNotSide, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("%s is the companion but is not in the sideboard", card.Name), oid)
	}
	add(res, CodeCompanionUncheck, mtgv1.Severity_SEVERITY_INFO,
		fmt.Sprintf("the companion condition of %s is not machine-checked yet", card.Name), oid)
}

// checkOwnership applies only in the owned pool modes (D-37). Copies are
// summed per Oracle id over the main deck, the sideboard, the commanders,
// and the companion.
func checkOwnership(res *mtgv1.ValidationResult, in Input) {
	if in.OracleCounts == nil {
		return
	}
	var sev mtgv1.Severity
	switch in.PoolRule {
	case mtgv1.PoolRule_POOL_RULE_OWNED_ONLY:
		sev = mtgv1.Severity_SEVERITY_BLOCK
	case mtgv1.PoolRule_POOL_RULE_OWNED_FIRST:
		sev = mtgv1.Severity_SEVERITY_WARN
	default:
		return // any-card: ownership is information, never a finding
	}
	need := map[string]int32{}
	order := []string{}
	for _, dc := range allCards(in.Deck) {
		if _, seen := need[dc.OracleId]; !seen {
			order = append(order, dc.OracleId)
		}
		need[dc.OracleId] += dc.Count
	}
	for _, oid := range order {
		card, ok := in.Cards.ByOracleID(oid)
		if !ok {
			continue
		}
		if isBasic(card) {
			// D-37 exception (owner decision 2026-08-24): basic lands are
			// always available, also in owned-only mode.
			continue
		}
		if owned := in.OracleCounts[oid]; owned < need[oid] {
			add(res, CodeNotOwned, sev,
				fmt.Sprintf("%s: the deck needs %d, the collection has %d", card.Name, need[oid], owned), oid)
		}
	}
}

// checkManaBase gives guide-range warnings, never blocks (corpus section 6).
func checkManaBase(res *mtgv1.ValidationResult, in Input, fr FormatRules) {
	var lands, nonlands int32
	var mvSum float64
	for _, dc := range counted(in.Deck.Cards) {
		card, ok := in.Cards.ByOracleID(dc.OracleId)
		if !ok {
			continue
		}
		if slices.Contains(card.CardTypes, "Land") {
			lands += dc.Count
		} else {
			nonlands += dc.Count
			mvSum += card.ManaValue * float64(dc.Count)
		}
	}
	lo, hi := int32(20), int32(27)
	if fr.Commander {
		lo, hi = 34, 38 // corpus section 6, the owner's guide
	}
	if lands < lo || lands > hi {
		add(res, CodeLandCount, mtgv1.Severity_SEVERITY_WARN,
			fmt.Sprintf("%d lands: the guide range for this format is %d to %d", lands, lo, hi), "")
	}
	if nonlands > 0 {
		add(res, CodeCurve, mtgv1.Severity_SEVERITY_INFO,
			fmt.Sprintf("average mana value %.2f over %d nonland cards", mvSum/float64(nonlands), nonlands), "")
	}
}

// checkPrintings flags the printing exception (roadmap F-12). The index
// carries one default printing per card with its digital flag. A digital
// default printing is information: a paper printing can exist.
// TODO(F-12): the index has no per-printing legality and no list of
// paper printings per Oracle id. A real "digital only" verdict needs that
// data from default_cards.
func checkPrintings(res *mtgv1.ValidationResult, in Input) {
	seen := map[string]bool{}
	for _, dc := range allCards(in.Deck) {
		if seen[dc.OracleId] {
			continue
		}
		seen[dc.OracleId] = true
		card, ok := in.Cards.ByOracleID(dc.OracleId)
		if !ok || !card.GetDefaultPrinting().GetDigital() {
			continue
		}
		add(res, CodeDigitalPrinting, mtgv1.Severity_SEVERITY_INFO,
			fmt.Sprintf("%s: the default printing is digital only, check for a paper printing", card.Name), dc.OracleId)
	}
}
