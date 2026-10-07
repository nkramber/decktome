package generate

import (
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// CodeThemeLeftOut is the warning of a deck that plays a card of the set
// fill and leaves out a nonland card on theme (D-1198). It buys the
// repair turn, because the swap is the fix, and the model can make it.
const CodeThemeLeftOut = "theme_left_out"

// themeCheckNames caps the card names of the finding, so a long list
// stays readable.
const themeCheckNames = 12

// checkThemeLeftOut adds CodeThemeLeftOut when the deck plays a card that
// only the set limit put on the list, and the shortlist holds a nonland
// card on theme that the deck leaves out. A set fill card has no theme
// signal and no staple job, so the card on theme is the better pick.
// Replays of D-1197 left out Chandra, Torch of Defiance and Garruk,
// Veiled Butcher, and played Yargle, Glutton of Urborg (D-1198).
//
// A revision follows the brief of the reader, and an upgrade keeps the
// precon, so neither one reads the check. A land on theme stays out of
// it, because a land swap moves the mana base.
func checkThemeLeftOut(deck *mtgv1.Deck, req Request) {
	if len(req.Themed) == 0 || len(req.SetFill) == 0 || req.Revision != nil || req.Precon != "" || req.Pool == nil {
		return
	}
	held := heldIDs(deck)
	var fill []string
	for _, dc := range deck.GetCards() {
		if req.SetFill[dc.GetOracleId()] {
			fill = append(fill, dc.GetName())
		}
	}
	if len(fill) == 0 {
		return
	}
	commander := map[string]bool{}
	for _, id := range deck.GetCommanderOracleIds() {
		commander[id] = true
	}
	var out []string
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok || !req.Themed[c.GetOracleId()] || held[c.GetOracleId()] || commander[c.GetOracleId()] {
			continue
		}
		if strings.Contains(c.GetTypeLine(), "Land") {
			continue
		}
		out = append(out, c.GetName())
	}
	if len(out) == 0 {
		return
	}
	sort.Strings(fill)
	sort.Strings(out)
	addFinding(deck, CodeThemeLeftOut, mtgv1.Severity_SEVERITY_WARN,
		fmt.Sprintf("the deck leaves out %s on theme, %s, and plays %s of the set with no theme and no job, %s",
			plural(len(out), "card"), capNames(out), plural(len(fill), "card"), capNames(fill)))
}

// capNames joins the names, and it counts the names past the cap.
func capNames(names []string) string {
	if len(names) <= themeCheckNames {
		return strings.Join(names, ", ")
	}
	return fmt.Sprintf("%s, and %d more", strings.Join(names[:themeCheckNames], ", "), len(names)-themeCheckNames)
}
