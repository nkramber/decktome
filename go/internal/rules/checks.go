package rules

import (
	"fmt"
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// mainDeckCount sums main-deck copies, the commanders included:
// a Commander deck is exactly 100 cards with the command zone.
func mainDeckCount(deck *mtgv1.Deck) int {
	n := len(deck.CommanderOracleIds)
	for _, c := range deck.Cards {
		n += int(c.Count)
	}
	return n
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
	for _, c := range deck.Sideboard {
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
// cards are exempt. Vintage-restricted cards cap at one copy.
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
		if fr.ScryfallKey != "" && card.Legalities[fr.ScryfallKey] == mtgv1.LegalityStatus_LEGALITY_STATUS_RESTRICTED {
			limit = 1
			if n > limit {
				add(res, CodeRestrictedCard, mtgv1.Severity_SEVERITY_BLOCK,
					fmt.Sprintf("%s is restricted: max 1 copy, the deck has %d", names[oid], n), oid)
				continue
			}
		}
		if n > limit {
			add(res, CodeCopyLimit, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("%s: %d copies, the limit is %d", names[oid], n, limit), oid)
		}
	}
}

// allCards walks the main deck, the sideboard, and the commanders.
func allCards(deck *mtgv1.Deck) []*mtgv1.DeckCard {
	out := make([]*mtgv1.DeckCard, 0, len(deck.Cards)+len(deck.Sideboard)+len(deck.CommanderOracleIds))
	out = append(out, deck.Cards...)
	out = append(out, deck.Sideboard...)
	for _, oid := range deck.CommanderOracleIds {
		out = append(out, &mtgv1.DeckCard{OracleId: oid, Count: 1})
	}
	return out
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
	if len(cmdrs) == 2 {
		if !validPair(cmdrs[0], cmdrs[1]) {
			add(res, CodeBadPartner, mtgv1.Severity_SEVERITY_BLOCK,
				fmt.Sprintf("%s and %s are not a legal commander pair", cmdrs[0].Name, cmdrs[1].Name), "")
		}
	}
}

// validPair checks the two-commander mechanics (corpus section 2.2).
func validPair(a, b *mtgv1.Card) bool {
	pk := func(c *mtgv1.Card) mtgv1.PartnerKind { return c.Partner }
	switch {
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_PARTNER && pk(b) == mtgv1.PartnerKind_PARTNER_KIND_PARTNER:
		return true
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_WITH && pk(b) == mtgv1.PartnerKind_PARTNER_KIND_WITH:
		return a.PartnerWithName == b.Name && b.PartnerWithName == a.Name
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_FRIENDS_FOREVER && pk(b) == mtgv1.PartnerKind_PARTNER_KIND_FRIENDS_FOREVER:
		return true
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND && b.IsBackground:
		return true
	case pk(b) == mtgv1.PartnerKind_PARTNER_KIND_CHOOSE_BACKGROUND && a.IsBackground:
		return true
	case pk(a) == mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION && isDoctor(b):
		return true
	case pk(b) == mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION && isDoctor(a):
		return true
	default:
		return false
	}
}

func isDoctor(c *mtgv1.Card) bool {
	return slices.Contains(c.Subtypes, "Doctor") && slices.Contains(c.Subtypes, "Time") &&
		c.Partner != mtgv1.PartnerKind_PARTNER_KIND_DOCTORS_COMPANION
}

func checkColorIdentity(res *mtgv1.ValidationResult, in Input) {
	allowed := map[mtgv1.Color]bool{}
	for _, oid := range in.Deck.CommanderOracleIds {
		if c, ok := in.Cards.ByOracleID(oid); ok {
			for _, col := range c.ColorIdentity {
				allowed[col] = true
			}
		}
	}
	for _, dc := range in.Deck.Cards {
		card, ok := in.Cards.ByOracleID(dc.OracleId)
		if !ok {
			continue
		}
		for _, col := range card.ColorIdentity {
			if !allowed[col] {
				add(res, CodeOffColor, mtgv1.Severity_SEVERITY_BLOCK,
					fmt.Sprintf("%s is outside the commander color identity", card.Name), dc.OracleId)
				break
			}
		}
	}
}

func checkBracket(cfg *Config, res *mtgv1.ValidationResult, in Input) {
	bracket := in.Deck.GetPower().GetBracket()
	if bracket == 0 {
		return
	}
	limit, ok := cfg.MaxGameChangers[bracket]
	if !ok {
		add(res, CodeGameChangers, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("unknown bracket %d", bracket), "")
		return
	}
	var changers []string
	for _, dc := range allCards(in.Deck) {
		if c, ok := in.Cards.ByOracleID(dc.OracleId); ok && c.GameChanger {
			changers = append(changers, c.Name)
		}
	}
	if limit >= 0 && len(changers) > limit {
		add(res, CodeGameChangers, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("bracket %d allows %d Game Changers, the deck has %d: %s",
				bracket, limit, len(changers), strings.Join(changers, ", ")), "")
	}
	if bracket <= 3 {
		add(res, CodeBracketProse, mtgv1.Severity_SEVERITY_INFO,
			"bracket prose rules (mass land denial, extra-turn chains, combo timing) are not machine-checked yet (F-11)", "")
	}
}

// checkCompanion enforces F-18: the companion flag, the sideboard seat,
// and the banned-as-companion list that Scryfall can not express.
func checkCompanion(cfg *Config, res *mtgv1.ValidationResult, in Input) {
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
	if cfg.BannedAsCompanion[card.Name] {
		add(res, CodeCompanionBanned, mtgv1.Severity_SEVERITY_BLOCK,
			fmt.Sprintf("%s is banned as a companion (it can still sit in the deck or command it)", card.Name), oid)
		return
	}
	add(res, CodeCompanionUncheck, mtgv1.Severity_SEVERITY_INFO,
		fmt.Sprintf("the companion condition of %s is not machine-checked yet", card.Name), oid)
}

// checkOwnership applies only in the owned pool modes (D-37).
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
	for _, dc := range allCards(in.Deck) {
		card, ok := in.Cards.ByOracleID(dc.OracleId)
		if !ok {
			continue
		}
		if isBasic(card) {
			continue // basic lands are treated as always available
		}
		if owned := in.OracleCounts[dc.OracleId]; owned < dc.Count {
			add(res, CodeNotOwned, sev,
				fmt.Sprintf("%s: the deck needs %d, the collection has %d", card.Name, dc.Count, owned), dc.OracleId)
		}
	}
}

// checkManaBase gives guide-range warnings, never blocks (corpus section 6).
func checkManaBase(res *mtgv1.ValidationResult, in Input, fr FormatRules) {
	var lands, nonlands int32
	var mvSum float64
	for _, dc := range in.Deck.Cards {
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
		lo, hi = 34, 40
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
