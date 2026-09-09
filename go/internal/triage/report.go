package triage

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

// Result is one verdict after the triage: where it went and what it
// wrote.
type Result struct {
	Route Route
	// Case is the artifact, when the route wrote one.
	Case Case
	// Applied is the file the case joined, when a run wrote it.
	Applied string
	// Err is the failure of one item. One bad verdict never stops the
	// run: the rest still reach the document.
	Err error
}

// Report writes the triage document and answers the verdict. The bar is
// that every verdict reached a class or an owner question, and that no
// item failed. A dry run reads no bar: it names what the judge would
// read and stops.
func Report(w io.Writer, rs []Result, day time.Time, dry bool, cost string) bool {
	byClass := map[string]int{}
	needJudge, keeps, owners, errs := 0, 0, 0, 0
	for _, r := range rs {
		switch {
		case r.Err != nil:
			errs++
		case r.Route.Keep:
			keeps++
		case r.Route.Owner:
			owners++
		case r.Route.Need == NeedJudge:
			needJudge++
		}
		if r.Route.Class.ID != "" {
			byClass[r.Route.Class.ID]++
		}
	}
	pass := errs == 0 && (dry || needJudge == 0) && len(rs) > 0

	_, _ = fmt.Fprintf(w, "# PR-28b feedback triage\n\n")
	_, _ = fmt.Fprintf(w, "Run date: %s.\n\n", day.UTC().Format("2006-01-02"))
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	lane := "live"
	if dry {
		lane = "dry, so it called no model and wrote no case"
	}
	_, _ = fmt.Fprintf(w, "Verdict: %s. %d verdicts, %s. %d reached a class, %d are keep cases, %d go to the owner, %d still want the judge, and %d failed.\n\n",
		verdict, len(rs), lane, len(rs)-keeps-owners-needJudge-errs, keeps, owners, needJudge, errs)
	if len(rs) == 0 {
		_, _ = fmt.Fprintf(w, "The harvest held no verdict. A triage can not pass with nothing to read.\n\n")
	}
	_, _ = fmt.Fprintf(w, "CAUTION: this file holds what a reader wrote. It holds no email (D-559).\n\n")

	_, _ = fmt.Fprintf(w, "## Summary\n\n| Measure | Value |\n|---|---|\n")
	_, _ = fmt.Fprintf(w, "| Verdicts read | %d |\n", len(rs))
	_, _ = fmt.Fprintf(w, "| Keep cases, from a thumbs up | %d |\n", keeps)
	_, _ = fmt.Fprintf(w, "| Owner questions | %d |\n", owners)
	_, _ = fmt.Fprintf(w, "| Still want the judge | %d |\n", needJudge)
	_, _ = fmt.Fprintf(w, "| Failed | %d |\n", errs)
	_, _ = fmt.Fprintf(w, "| Cost | %s |\n\n", cost)

	if len(byClass) > 0 {
		ids := make([]string, 0, len(byClass))
		for id := range byClass {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		_, _ = fmt.Fprintf(w, "## By class\n\n| Class | Fault | Artifact | Verdicts |\n|---|---|---|---|\n")
		for _, id := range ids {
			c, _ := ClassByID(id)
			_, _ = fmt.Fprintf(w, "| %s | %s | %s | %d |\n", id, c.Name, c.Artifact, byClass[id])
		}
		_, _ = fmt.Fprintf(w, "\n")
	}

	_, _ = fmt.Fprintf(w, "## The verdicts\n\n")
	for _, r := range rs {
		rec := r.Route.Record
		_, _ = fmt.Fprintf(w, "### %s %s/%s\n\n", rec.CreatedAt.Format("2006-01-02 15:04"), rec.Kind, rec.Verdict)
		_, _ = fmt.Fprintf(w, "- id `%s`, user `%s`\n", rec.ID, rec.UID)
		_, _ = fmt.Fprintf(w, "- route: %s\n", routeWord(r.Route))
		_, _ = fmt.Fprintf(w, "- why: %s\n", r.Route.Why)
		if said := readerSaid(rec); said != "" {
			_, _ = fmt.Fprintf(w, "- the reader: %s\n", said)
		}
		if r.Err != nil {
			_, _ = fmt.Fprintf(w, "- **ERROR: %v**\n", r.Err)
		}
		if r.Case.Artifact != "" {
			_, _ = fmt.Fprintf(w, "- artifact: %s", r.Case.Artifact)
			if r.Case.Target != "" {
				_, _ = fmt.Fprintf(w, ", for `%s`", r.Case.Target)
			}
			_, _ = fmt.Fprintf(w, "\n")
		}
		if r.Case.Detail != "" {
			_, _ = fmt.Fprintf(w, "- detail: %s\n", r.Case.Detail)
		}
		for _, g := range r.Case.Gaps {
			_, _ = fmt.Fprintf(w, "- **gap**: %s\n", g)
		}
		if r.Applied != "" {
			_, _ = fmt.Fprintf(w, "- written to `%s`\n", r.Applied)
		}
		if len(r.Case.Body) > 0 {
			_, _ = fmt.Fprintf(w, "\n```json\n%s\n```\n", string(r.Case.Body))
		}
		_, _ = fmt.Fprintf(w, "\n")
	}
	return pass
}

// routeWord names where a verdict went, in one word or two.
func routeWord(r Route) string {
	switch {
	case r.Keep:
		return "keep case"
	case r.Owner:
		return "the owner, on " + orUnknown(r.Decision)
	case r.Need == NeedJudge:
		return "the judge"
	case r.Class.ID != "":
		return r.Class.ID + ", " + r.Class.Name
	}
	return "nothing to triage"
}

// Cases lists the cases of a run, newest last, so a caller can write
// them in one pass.
func Cases(rs []Result) []Case {
	var out []Case
	for _, r := range rs {
		if r.Err == nil && len(r.Case.Body) > 0 {
			out = append(out, r.Case)
		}
	}
	return out
}

// Owners lists the verdicts that go to the owner and no further.
func Owners(rs []Result) []Route {
	var out []Route
	for _, r := range rs {
		if r.Err == nil && r.Route.Owner {
			out = append(out, r.Route)
		}
	}
	return out
}

// OwnerRow writes the three cells of an owner-question row from one
// route. The words are the triage's own, and the reader's words follow
// them, because the owner answers better with both (D-642).
func OwnerRow(r Route) (question, whyYou, blocks string) {
	question = fmt.Sprintf("A reader disagrees with %s. %s The reader gave a thumbs down on a %s. %s",
		orUnknown(r.Decision), decisionSays(r.Decision), r.Record.Kind, readerSaid(r.Record))
	whyYou = "It changes a rule you set, and the loop refuses to decide one."
	blocks = "The fix for this verdict. No case is written until you answer."
	return question, whyYou, blocks
}

// decisionSays reads the words of a disputed decision, for the row.
func decisionSays(id string) string {
	for _, d := range disputedDecisions {
		if d.ID == id {
			return d.Says
		}
	}
	return ""
}

// Kinds counts the verdicts by kind, for the run line of a caller that
// prints one.
func Kinds(rs []Result) string {
	byKind := map[string]int{}
	for _, r := range rs {
		byKind[r.Route.Record.Kind]++
	}
	kinds := make([]string, 0, len(byKind))
	for k := range byKind {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)
	var parts []string
	for _, k := range kinds {
		parts = append(parts, fmt.Sprintf("%s %d", k, byKind[k]))
	}
	return strings.Join(parts, ", ")
}
