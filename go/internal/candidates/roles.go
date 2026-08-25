package candidates

import (
	"slices"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// assignRole gives a card one role (corpus section 6). Tags decide first.
// Type line and text decide the rest. onTheme marks a theme signal.
//
// Order matters: a land is a land even when it also ramps. A sweeper is a
// wipe, not removal. A card with a theme signal and no staple role is a
// synergy piece or, for a big creature or planeswalker, a threat.
func assignRole(c *mtgv1.Card, roleTags map[string]map[string]bool, onTheme bool) (mtgv1.CardRole, string) {
	in := func(role string) bool { return roleTags[role][c.OracleId] }
	text := strings.ToLower(c.OracleText)
	isLand := slices.Contains(c.CardTypes, "Land")
	switch {
	case isLand:
		return mtgv1.CardRole_CARD_ROLE_LAND, "type:land"
	case in("wipe"):
		return mtgv1.CardRole_CARD_ROLE_WIPE, "tag:sweeper"
	case in("wincon"):
		return mtgv1.CardRole_CARD_ROLE_WINCON, "tag:alternate-win-condition"
	case strings.Contains(text, "you win the game"):
		return mtgv1.CardRole_CARD_ROLE_WINCON, "text:you win the game"
	case in("ramp"):
		return mtgv1.CardRole_CARD_ROLE_RAMP, "tag:ramp"
	case in("interaction"):
		return mtgv1.CardRole_CARD_ROLE_INTERACTION, "tag:interaction"
	case in("removal"):
		return mtgv1.CardRole_CARD_ROLE_REMOVAL, "tag:removal"
	case in("draw"):
		return mtgv1.CardRole_CARD_ROLE_DRAW, "tag:draw"
	}
	// Text fallbacks for snapshots without tags.
	switch {
	case strings.Contains(text, "counter target spell"):
		return mtgv1.CardRole_CARD_ROLE_INTERACTION, "text:counter"
	case strings.Contains(text, "destroy all") || strings.Contains(text, "exile all"):
		return mtgv1.CardRole_CARD_ROLE_WIPE, "text:destroy all"
	case strings.Contains(text, "destroy target"):
		return mtgv1.CardRole_CARD_ROLE_REMOVAL, "text:destroy target"
	case exileIsRemoval(text):
		return mtgv1.CardRole_CARD_ROLE_REMOVAL, "text:exile target"
	case strings.Contains(text, "draw a card") || strings.Contains(text, "draw two cards"):
		return mtgv1.CardRole_CARD_ROLE_DRAW, "text:draw"
	case strings.Contains(text, "search your library for a basic land") || strings.Contains(text, "add {"):
		return mtgv1.CardRole_CARD_ROLE_RAMP, "text:mana"
	}
	if !onTheme {
		return mtgv1.CardRole_CARD_ROLE_OTHER, ""
	}
	isCreature := slices.Contains(c.CardTypes, "Creature")
	isWalker := slices.Contains(c.CardTypes, "Planeswalker")
	if isWalker || (isCreature && c.ManaValue >= 4) {
		return mtgv1.CardRole_CARD_ROLE_THREAT, "type:threat"
	}
	return mtgv1.CardRole_CARD_ROLE_SYNERGY, "theme"
}

// exileIsRemoval reports whether an "exile target" clause removes a
// permanent for good. A blink spell also says "exile target", but it
// exiles a permanent you own and returns it, so it is not removal.
func exileIsRemoval(text string) bool {
	found := false
	for i := 0; i < len(text); {
		j := strings.Index(text[i:], "exile ")
		if j < 0 {
			break
		}
		i += j + len("exile ")
		clause := text[i:min(i+120, len(text))]
		if !strings.HasPrefix(clause, "target") && !strings.HasPrefix(clause, "another target") &&
			!strings.HasPrefix(clause, "up to one target") && !strings.HasPrefix(clause, "up to two target") {
			continue
		}
		found = true
		if strings.Contains(clause, "you control") || strings.Contains(clause, "you own") ||
			strings.Contains(clause, "then return") || strings.Contains(clause, "return it") ||
			strings.Contains(clause, "return that card") || strings.Contains(clause, "return them") {
			return false
		}
	}
	return found
}
