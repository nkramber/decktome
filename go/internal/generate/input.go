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
