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
		// D-218 sets the share, and the prompt states it as a limit the
		// model must meet, not as a preference.
		fmt.Fprintf(&s, "\nThis deck upgrades the %s precon. Keep at least %d percent of its cards. The shortlist marks them \"precon\".\n",
			req.Precon, PreconSharePercent)
	}
	if len(req.Targets) > 0 {
		s.WriteString("\n## Job targets\n\n")
		keys := make([]string, 0, len(req.Targets))
		for k := range req.Targets {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(&s, "- %s: %d\n", k, req.Targets[k])
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

// CAUTION: nothing reaches this function today, and nothing can. The
// share needs the precon's card list, and no precon decklist source
// exists in this repo. The question workflow holds only PreconName, which
// is the first card the user named and not a list. The Scryfall snapshot
// carries set codes and no per-product decklist.
//
// D-218 stands as the owner's answer. It waits on a precon ingester,
// which OQ-40 asks for (D-240).
//
// checkPreconShare adds a finding when the deck keeps less of the precon
// than D-218 requires. It is a build rule and not a rule of the game, so
// it is a warning and never a block: the user asked for an upgrade, and a
// refusal to return a deck serves nobody.
func checkPreconShare(deck *mtgv1.Deck, req Request) {
	want := len(req.PreconOracleIDs)
	if want == 0 {
		return
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
	if kept*100 >= want*PreconSharePercent {
		return
	}
	deck.Validation.Findings = append(deck.GetValidation().GetFindings(), &mtgv1.Finding{
		Code:     CodePreconShare,
		Severity: mtgv1.Severity_SEVERITY_WARN,
		Message: fmt.Sprintf("the deck keeps %d of the %d %s precon cards, and the rule asks for %d percent",
			kept, want, req.Precon, PreconSharePercent),
	})
}
