package triage

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nkramber/decktome/go/internal/harvest"
)

// Need says what a verdict still wants before it becomes a case.
type Need string

const (
	// NeedNothing means the reason keys answered the class, so the case
	// costs no model call.
	NeedNothing Need = "none"
	// NeedJudge means the judge role has to read the item. Three things
	// ask for it: a verdict with free text and no reason, a verdict whose
	// reasons cross two classes, and a verdict of a disputing class that
	// carries free text. A dry run stops here and reports it.
	NeedJudge Need = "judge"
)

// Route is what the triage decided for one verdict.
type Route struct {
	Record harvest.Record
	// Class is the fault, when one is known. A verdict that needs the
	// judge carries the zero class until the judge answers.
	Class Class
	Need  Need
	// Why says how the route was reached, in one line.
	Why string
	// Keep marks a thumbs up. It is a case a change must not flip, and it
	// carries no class: nothing was wrong.
	Keep bool
	// Owner marks a verdict the judge sent to the owner. It writes a row
	// in docs/owner-questions.md and no case (D-558).
	Owner bool
	// Decision names the decision the reader argues with, when Owner is
	// set or when the class disputes one.
	Decision string
}

// RouteOf reads one verdict and answers where it goes. It calls no
// model: a route that needs one says so, and the caller decides whether
// to spend (D-643).
func RouteOf(rec harvest.Record) Route {
	r := Route{Record: rec}
	if rec.Verdict == "up" {
		r.Keep, r.Need, r.Why = true, NeedNothing, "a thumbs up is a keep case, and a change must not flip it"
		return r
	}
	found := classesOf(rec)
	switch len(found) {
	case 0:
		if strings.TrimSpace(rec.Text) == "" {
			r.Need, r.Why = NeedNothing, "the verdict names no reason this build knows and no text, so it holds nothing to triage"
			return r
		}
		r.Need, r.Why = NeedJudge, "the verdict names no reason this build knows, so the judge reads the words"
		return r
	case 1:
		r.Class = found[0]
		r.Decision = found[0].Disputes
		if found[0].Disputes != "" && strings.TrimSpace(rec.Text) != "" {
			r.Need = NeedJudge
			r.Why = fmt.Sprintf("%s meets %s, and the reader wrote words, so the judge says case or owner question", found[0].ID, found[0].Disputes)
			return r
		}
		r.Need, r.Why = NeedNothing, "the reason key names the class"
		return r
	default:
		ids := make([]string, 0, len(found))
		for _, c := range found {
			ids = append(ids, c.ID)
		}
		r.Need = NeedJudge
		r.Why = fmt.Sprintf("the reasons cross %d classes, %s, so the judge picks one", len(found), strings.Join(ids, " and "))
		return r
	}
}

// classesOf reads the classes the checked reasons name, in the order of
// the class table. A reason key this build does not know is dropped: a
// dialog may have offered a key a later build renamed, and the judge
// then reads the words.
func classesOf(rec harvest.Record) []Class {
	seen := map[string]Class{}
	for _, key := range rec.Reasons {
		if c, ok := ClassOf(rec.Kind, strings.TrimSpace(key)); ok {
			seen[c.ID] = c
		}
	}
	out := make([]Class, 0, len(seen))
	for _, c := range seen {
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Apply writes a judge answer onto a route. An answer that names no
// class this build knows leaves the route where it was, so a judge that
// invents a class never writes a case.
func (r Route) Apply(v Verdict) Route {
	r.Why = strings.TrimSpace(v.Why)
	if v.OwnerQuestion {
		r.Owner, r.Need, r.Decision = true, NeedNothing, strings.TrimSpace(v.Decision)
		if r.Why == "" {
			r.Why = "the judge sent it to the owner"
		}
		return r
	}
	c, ok := ClassByID(strings.TrimSpace(v.Class))
	if !ok {
		r.Why = fmt.Sprintf("the judge named class %q, which this build does not hold", v.Class)
		return r
	}
	r.Class, r.Need = c, NeedNothing
	if r.Decision == "" {
		r.Decision = c.Disputes
	}
	return r
}

// RowOf reads the catalog row a question verdict is about. The session
// names a question "q<n>-<row>" (agent.go), so the row costs no model
// call and no text match. It answers an empty string when the id does
// not read that way.
func RowOf(questionID string) string {
	_, row, ok := strings.Cut(strings.TrimSpace(questionID), "-")
	if !ok {
		return ""
	}
	return row
}
