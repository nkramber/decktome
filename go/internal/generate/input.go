package generate

import (
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
)

// input writes the session text the model reads. The stable instructions
// carry the rules, and this block carries the session, so the cached
// prefix holds across a repair turn (roadmap PR-8).
//
// The shortlist goes last and it is the longest part. Findings and misses
// are empty on the first turn.
func (b *Builder) input(req Request, misses []Miss, blocks []*mtgv1.Finding) string {
	var s strings.Builder
	fmt.Fprintf(&s, "## Limits\n\n%s\n", strings.TrimSpace(req.Limits))
	fmt.Fprintf(&s, "\n## The deck the user asked for\n\n%s\n", strings.TrimSpace(req.Plan))
	if len(req.Commanders) > 0 {
		var names []string
		for _, id := range req.Commanders {
			if c, ok := b.cards.ByOracleID(id); ok {
				names = append(names, c.GetName())
			}
		}
		if len(names) > 0 {
			fmt.Fprintf(&s, "\nThe commander is %s. It is chosen, and it is not one of the cards you list.\n",
				strings.Join(names, " and "))
		}
	}
	if req.BudgetUSD > 0 {
		what := "the cards you must buy"
		if req.BudgetWholeDeck {
			what = "every card in the deck, the copies the user owns included"
		}
		fmt.Fprintf(&s, "\n## Budget\n\nThe deck must cost $%.2f or less, counting %s. Each shortlist line ends with the price of one copy.\n",
			req.BudgetUSD, what)
		if req.OracleCounts != nil && !req.BudgetWholeDeck {
			s.WriteString("A copy the user already owns costs nothing, so prefer the cards marked owned.\n")
		}
		s.WriteString("Stay under the cap. Choose a cheaper card that does the same job when one is on the list.\n")
	}
	if len(req.Locked) > 0 {
		var names []string
		for _, n := range req.Pool.Names() {
			c, ok := req.Pool.Card(n)
			if !ok {
				continue
			}
			for _, id := range req.Locked {
				if c.GetOracleId() == id {
					names = append(names, c.GetName())
				}
			}
		}
		if len(names) > 0 {
			// The user named these, so they are not a preference (D-242).
			fmt.Fprintf(&s, "\nThe deck must hold %s. The user asked to keep %s.\n",
				strings.Join(names, ", "), these(len(names)))
		}
	}
	if req.Precon != "" {
		// D-218 sets the share. A percentage asks the model to do
		// arithmetic against a list it is still writing, and deck gate
		// prompts 17 and 18 of 2026-08-28 kept 68 and 29 percent. The
		// prompt states the count instead, and what it may change (D-248).
		keep := PreconKeepCount(len(req.PreconOracleIDs))
		change := len(req.PreconOracleIDs) - keep
		fmt.Fprintf(&s, "\n## The precon\n\nThis deck upgrades the %s precon, which holds %d cards. The shortlist marks each one \"precon\".\n",
			req.Precon, len(req.PreconOracleIDs))
		fmt.Fprintf(&s, "Keep at least %d of them. You may drop at most %d, and replace those with anything else on the shortlist.\n",
			keep, change)
		s.WriteString("An upgrade is a small number of better cards, and not a new deck.\n")
		s.WriteString("Keep the precon's own shape. It is a working deck, so do not rebuild it to a role template, and no job target is given.\n")
		// The mana base is the one thing a role template got right, and
		// the precon already holds it. Naming its own count keeps it
		// without setting a second quota against the share (D-251).
		if lands := req.PreconLands; lands > 0 {
			fmt.Fprintf(&s, "The precon holds %s. Keep about that many, and count the precon's own lands toward it.\n",
				plural(lands, "land"))
		}
		s.WriteString("Count the cards you keep before you answer. The count above is a limit and not a goal.\n")
	}
	// An upgrade keeps the precon's own composition. The generic job
	// targets prescribe the whole deck, and the share demands most of
	// those slots come from the precon, so the two instructions fight and
	// the model splits the difference: prompt 17 kept 54 of the 68 it
	// needed. A precon is a working deck already (D-249).
	//
	// The land count is not one of those jobs. It is the mana base, and
	// dropping it with the rest gave both precon decks 25 lands against a
	// guide of 34 to 38, which the land-count advisory caught (D-251).
	// A generic land target is a second quota. With "keep 68 precon
	// cards" it reads as 36 plus 68 of 99 slots, which cannot be met, and
	// the share fell to 61. The precon's own lands are precon cards, so
	// the upgrade prompt names the precon's land count in its own block
	// and sends no target here (D-251).
	targets := req.Targets
	if req.Precon != "" {
		targets = nil
	}
	if len(targets) > 0 {
		s.WriteString("\n## Job targets\n\n")
		keys := make([]string, 0, len(targets))
		for k := range targets {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(&s, "- %s: %d\n", k, targets[k])
		}
	}
	if len(misses) > 0 {
		s.WriteString("\n## Names that are not on the shortlist\n\n")
		for _, m := range misses {
			if len(m.Near) == 0 {
				fmt.Fprintf(&s, "- %q is not on the shortlist. Replace it with a shortlist card that does the same job.\n", m.Name)
				continue
			}
			fmt.Fprintf(&s, "- %q is not on the shortlist. The shortlist holds %s.\n", m.Name, strings.Join(m.Near, ", "))
		}
	}
	if len(blocks) > 0 {
		s.WriteString("\n## Findings against your deck\n\n")
		for _, f := range blocks {
			fmt.Fprintf(&s, "- %s: %s\n", f.GetCode(), f.GetMessage())
		}
	}
	s.WriteString("\n## The shortlist\n\n")
	s.WriteString(b.shortlist(req))
	return s.String()
}

// shortlist writes one line per card: the name as the model must copy it,
// the job, and the owned count when a collection is attached.
func (b *Builder) shortlist(req Request) string {
	precon := make(map[string]bool, len(req.PreconOracleIDs))
	for _, id := range req.PreconOracleIDs {
		precon[id] = true
	}
	var s strings.Builder
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok {
			continue
		}
		fmt.Fprintf(&s, "- %s", c.GetName())
		if t := strings.TrimSpace(c.GetTypeLine()); t != "" {
			fmt.Fprintf(&s, " | %s", t)
		}
		if job := req.Roles[c.GetOracleId()]; job != "" {
			fmt.Fprintf(&s, " | %s", job)
		}
		if req.OracleCounts != nil {
			fmt.Fprintf(&s, " | owned %d", req.OracleCounts[c.GetOracleId()])
		}
		// The model cannot budget what it cannot see. Deck gate run 5
		// spent $268.37 against a $100.00 cap on a shortlist whose
		// cheapest 99 cards cost $25.66, because no line carried a price
		// (D-244).
		if req.BudgetUSD > 0 {
			fmt.Fprintf(&s, " | $%.2f", c.GetPriceUsd())
		}
		if precon[c.GetOracleId()] {
			s.WriteString(" | precon")
		}
		s.WriteString("\n")
	}
	return s.String()
}

// PreconSharePercent is how much of a named precon a built deck keeps.
// The owner set it on 2026-08-26, and called it a start and not a settled
// figure (D-218, answers OQ-21).
const PreconSharePercent = 85

// CodePreconShare is the finding an upgrade gets when it drops too much
// of the precon it was asked to upgrade.
const CodePreconShare = "precon_share"

// internal/precons holds the decklists, and agentsvc reads the product
// name from the user's own words, because the classifier reports a card
// name for an upgrade request and not a product (D-247).
//
// checkPreconShare adds a finding when the deck keeps less of the precon
// than D-218 requires. It is a build rule and not a rule of the game, so
// it is a warning and never a block: the user asked for an upgrade, and a
// refusal to return a deck serves nobody.
func checkPreconShare(deck *mtgv1.Deck, req Request) bool {
	want := len(req.PreconOracleIDs)
	if want == 0 {
		return false
	}
	in := make(map[string]bool, want)
	for _, id := range req.PreconOracleIDs {
		in[id] = true
	}
	kept := 0
	for _, c := range deck.GetCards() {
		if in[c.GetOracleId()] {
			kept++
		}
	}
	// The commander counts: it is in the deck, in the command zone.
	for _, id := range deck.GetCommanderOracleIds() {
		if in[id] {
			kept++
		}
	}
	if kept >= PreconKeepCount(want) {
		return false
	}
	deck.Validation.Findings = append(deck.GetValidation().GetFindings(), &mtgv1.Finding{
		Code:     CodePreconShare,
		Severity: mtgv1.Severity_SEVERITY_WARN,
		Message: fmt.Sprintf("the deck keeps %d of the %d %s precon cards, and the rule asks for %d",
			kept, want, req.Precon, PreconKeepCount(want)),
	})
	return true
}
