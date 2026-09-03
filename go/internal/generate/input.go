package generate

import (
	"fmt"
	"sort"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// input writes the session text the model reads. The stable instructions
// carry the rules, and this block carries the session, so the cached
// prefix holds across a repair turn (roadmap PR-8).
//
// The shortlist goes last and it is the longest part. Findings and misses
// are empty on the first turn. On the repair turn, findings holds every
// finding that bought the turn, the two warnings included (D-244, D-248).
func (b *Builder) input(req Request, misses []Miss, findings []*mtgv1.Finding) string {
	var s strings.Builder
	fmt.Fprintf(&s, "## Limits\n\n%s\n", strings.TrimSpace(req.Limits))
	fmt.Fprintf(&s, "\n## The deck the user asked for\n\n%s\n", strings.TrimSpace(req.Plan))
	// Only Commander has a command zone, so a 60-card session never hears
	// that the deck has one (D-233).
	if req.Format == mtgv1.FormatId_FORMAT_ID_COMMANDER && len(req.Commanders) > 0 {
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
		for _, id := range req.Locked {
			if c, ok := req.Pool.ByOracleID(id); ok {
				names = append(names, c.GetName())
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
		// arithmetic against a list it is still writing, so the prompt
		// states the count instead, and what it may change (D-248).
		//
		// The count is the precon's nonbasic names. Basic lands swap free,
		// and the change count is a ceiling and not a target: an upgrade
		// makes the smallest set of changes that keeps the theme (D-218).
		names := len(preconNonbasics(req, b.cards))
		keep := PreconKeepCount(names)
		change := names - keep
		fmt.Fprintf(&s, "\n## The precon\n\nThis deck upgrades the %s precon. Its nonbasic cards are %d names, and the shortlist marks each one \"precon\". Basic lands are not counted, and you may swap them freely.\n",
			req.Precon, names)
		fmt.Fprintf(&s, "Keep at least %d of those names. You may change at most %d, and that number is a ceiling and not a target.\n",
			keep, change)
		s.WriteString("Change as few cards as the upgrade needs, and keep the theme of the precon intact. Never aim to replace the full number you may change.\n")
		s.WriteString("An upgrade is a small number of better cards, and not a new deck.\n")
		s.WriteString("Keep the precon's own shape. It is a working deck, so do not rebuild it to a role template, and no job target is given.\n")
		// The mana base is the one thing a role template got right, and
		// the precon already holds it. Naming its own count keeps it
		// without setting a second quota against the share (D-251).
		if lands := req.PreconLands; lands > 0 {
			fmt.Fprintf(&s, "The precon holds %s. Keep about that many, and count the precon's own lands toward it.\n",
				plural(lands, "land"))
		}
		s.WriteString("Count the names you keep before you answer. The count above is a limit and not a goal.\n")
	}
	if r := req.Revision; r != nil {
		s.WriteString("\n## The deck you are revising\n\n")
		s.WriteString("The user read this deck and asked for a change. Keep every card the change does not touch.\n\n")
		for _, c := range r.Base {
			fmt.Fprintf(&s, "- %d %s (%s)\n", c.GetCount(), c.GetName(), roleWord(c.GetRole()))
		}
		s.WriteString("\n## The change the user asked for\n\n")
		for _, line := range r.Instructions {
			fmt.Fprintf(&s, "- %s\n", line)
		}
		if len(r.Remove) > 0 {
			fmt.Fprintf(&s, "- Out: %s. These are not on the shortlist, so never name them.\n", strings.Join(r.Remove, ", "))
		}
		if len(r.Keep) > 0 {
			fmt.Fprintf(&s, "- Keep: %s.\n", strings.Join(r.Keep, ", "))
		}
		if r.MaxManaValue > 0 {
			fmt.Fprintf(&s, "- No nonland card above mana value %g. The shortlist holds none.\n", r.MaxManaValue)
		}
		if r.SwapBasics > 0 {
			fmt.Fprintf(&s, "- Replace at least %d basic lands with nonbasic lands from the shortlist", r.SwapBasics)
			if k := strings.TrimSpace(r.LandKinds); k != "" {
				fmt.Fprintf(&s, ": %s", strings.TrimRight(k, "."))
			}
			s.WriteString(". Keep the land total the same. Count the nonbasic lands you add before you answer: the number is a floor and not a ceiling.\n")
		}
	}
	// An upgrade keeps the precon's own composition. The generic job
	// targets prescribe the whole deck, and the share demands most of
	// those slots come from the precon, so the two instructions fight. A
	// precon is a working deck already (D-249).
	//
	// The land count is not one of those jobs. It is the mana base, and
	// a generic land target is a second quota against the share. The
	// upgrade prompt names the precon's land count in its own block and
	// sends no target here (D-251).
	targets := req.Targets
	if req.Precon != "" || req.Revision != nil {
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
		// The deck shape is the rest of the bracket's band: the curve,
		// the mana base, and the power signals the profile checks after
		// the build (PR-14A). It goes out with the targets, and not to
		// an upgrade or a revision.
		if b.profiler != nil {
			if lines := b.profiler.Bands().Lines(req.Format, req.Power); len(lines) > 0 {
				s.WriteString("\n## Deck shape\n\nBuild inside these limits. A check reads them after the build.\n\n")
				for _, line := range lines {
					s.WriteString(line + "\n")
				}
			}
		}
		// The format shape is what the top lists of the format look
		// like: their land count, their curve, and the cards they hold
		// most (PR-14B). It is a description, and the bands above are
		// the limits.
		if b.scorer != nil {
			if lines := b.scorer.ShapeLines(req.Format); len(lines) > 0 {
				s.WriteString("\n## Format shape\n\nThe published top lists of the format look like this.\n\n")
				for _, line := range lines {
					s.WriteString(line + "\n")
				}
			}
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
	if len(findings) > 0 {
		s.WriteString("\n## Findings against your deck\n\n")
		for _, f := range findings {
			fmt.Fprintf(&s, "- %s: %s\n", f.GetCode(), f.GetMessage())
		}
	}
	s.WriteString("\n## The shortlist\n\n")
	s.WriteString(b.shortlist(req))
	return s.String()
}

// shortlist writes one line per card: the name as the model must copy it,
// the job, and the owned count when a collection is attached. The
// commander is in the command zone, so the shortlist omits it and the
// pool keeps it (D-302).
func (b *Builder) shortlist(req Request) string {
	precon := make(map[string]bool, len(req.PreconOracleIDs))
	for _, id := range req.PreconOracleIDs {
		precon[id] = true
	}
	omit := map[string]bool{}
	if req.Format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		for _, id := range req.Commanders {
			omit[id] = true
		}
	}
	var s strings.Builder
	for _, name := range req.Pool.Names() {
		c, ok := req.Pool.Card(name)
		if !ok || omit[c.GetOracleId()] {
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
		// The model cannot budget what it cannot see, so every line
		// carries a price when a budget applies (D-244).
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

// checkPreconShare adds a finding when the deck keeps fewer of the
// precon's nonbasic names than D-218 requires. It is a build rule and
// not a rule of the game, so it is a warning and never a block: the user
// asked for an upgrade, and a refusal to return a deck serves nobody. It
// reports whether the finding was added.
func checkPreconShare(deck *mtgv1.Deck, req Request, cards rules.CardSource) bool {
	in := preconNonbasics(req, cards)
	want := len(in)
	if want == 0 {
		return false
	}
	kept := len(heldPrecon(deck, in))
	if kept >= PreconKeepCount(want) {
		return false
	}
	addFinding(deck, CodePreconShare, mtgv1.Severity_SEVERITY_WARN,
		fmt.Sprintf("the deck keeps %d of the %d nonbasic %s precon names, and the rule asks for %d",
			kept, want, req.Precon, PreconKeepCount(want)))
	return true
}
