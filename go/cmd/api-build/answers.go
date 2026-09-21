package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// plan holds one reply for each slot the run expects to be asked about.
// A question id is new in every turn, so the plan keys on Question.slot,
// which is stable.
type plan map[string]reply

// reply is one prepared answer. Exactly one of its three forms applies.
type reply struct {
	// option is the 1-based position in Question.options. Zero means the
	// reply carries no option.
	option int
	text   string
	// decline hands the choice back, and the generator then applies the
	// default the corpus names (D-353).
	decline bool
}

// parsePlan reads the -answers flag: "slot=text;slot=#2;slot=decline".
// "#N" names the Nth option, counted from one. An empty string gives an
// empty plan, and the defaults of answerFor then carry the run.
func parsePlan(s string) (plan, error) {
	out := plan{}
	for _, part := range strings.Split(s, ";") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		slot, value, found := strings.Cut(part, "=")
		slot = strings.TrimSpace(slot)
		value = strings.TrimSpace(value)
		if !found || slot == "" {
			return nil, fmt.Errorf("answer %q is not slot=value", part)
		}
		if _, dup := out[slot]; dup {
			return nil, fmt.Errorf("slot %q has more than one answer", slot)
		}
		switch {
		case value == "decline":
			out[slot] = reply{decline: true}
		case strings.HasPrefix(value, "#"):
			n, err := strconv.Atoi(value[1:])
			if err != nil || n < 1 {
				return nil, fmt.Errorf("slot %q: %q is not an option number of one or more", slot, value)
			}
			out[slot] = reply{option: n}
		case value == "":
			return nil, fmt.Errorf("slot %q has an empty answer", slot)
		default:
			out[slot] = reply{text: value}
		}
	}
	return out, nil
}

// answerFor builds the reply to one question. The plan wins. With no
// plan entry the run still answers, so it never stalls with nobody at
// the keyboard:
//
//   - A closed question takes its first option, because the options are
//     the whole answer space (D-295).
//   - A no_decline question takes its first option, because the UI shows
//     no decline control for it (D-690).
//   - Every other question is declined (D-353).
func answerFor(q *mtgv1.Question, p plan) (*mtgv1.Answer, error) {
	a := &mtgv1.Answer{QuestionId: q.GetId()}
	r, planned := p[q.GetSlot()]
	if !planned {
		if q.GetClosed() || q.GetNoDecline() {
			if len(q.GetOptions()) == 0 {
				return nil, fmt.Errorf("question %q of slot %q offers no option and allows no decline", q.GetText(), q.GetSlot())
			}
			i := int32(0)
			a.OptionIndex = &i
			return a, nil
		}
		a.Declined = true
		return a, nil
	}
	switch {
	case r.decline:
		if q.GetNoDecline() {
			return nil, fmt.Errorf("slot %q refuses a decline (D-690): name an option or a text answer", q.GetSlot())
		}
		a.Declined = true
	case r.option > 0:
		if r.option > len(q.GetOptions()) {
			return nil, fmt.Errorf("slot %q: option #%d of %d offered", q.GetSlot(), r.option, len(q.GetOptions()))
		}
		i := int32(r.option - 1)
		a.OptionIndex = &i
	default:
		if q.GetClosed() {
			return nil, fmt.Errorf("slot %q is closed and takes no free text (D-295): name an option", q.GetSlot())
		}
		a.Text = r.text
	}
	return a, nil
}

// answerAll builds one reply for each question of a turn, in the order
// the agent asked them.
func answerAll(qs []*mtgv1.Question, p plan) ([]*mtgv1.Answer, error) {
	out := make([]*mtgv1.Answer, 0, len(qs))
	for _, q := range qs {
		a, err := answerFor(q, p)
		if err != nil {
			return nil, err
		}
		out = append(out, a)
	}
	return out, nil
}

// describe prints one answer beside the question it replies to, so the
// log of a run shows what the command chose on the reader's behalf.
func describe(q *mtgv1.Question, a *mtgv1.Answer) string {
	switch {
	case a.GetDeclined():
		return "declined"
	case a.OptionIndex != nil:
		i := int(a.GetOptionIndex())
		if i < len(q.GetOptions()) {
			return fmt.Sprintf("option #%d %q", i+1, q.GetOptions()[i])
		}
		return fmt.Sprintf("option #%d", i+1)
	default:
		return fmt.Sprintf("text %q", a.GetText())
	}
}

// slots lists the plan's slots in order, for the report.
func (p plan) slots() []string {
	out := make([]string, 0, len(p))
	for k := range p {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
