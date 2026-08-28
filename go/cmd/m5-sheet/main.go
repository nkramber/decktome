// Command m5-sheet builds the M-5 scoring sheet from one or more gate
// documents (D-66, D-104).
//
// Version 1 held only a question the model offered to replace. That was
// 60 of the 793 questions of gate runs 10 to 13, or 7.6 percent, and a
// question that was wrong but that the model never challenged could not
// become an item. The owner found such questions anyway, by reading
// around the items. The sheet now holds both kinds (D-104):
//
//   - A replacement item pairs the question the model wrote with the
//     catalog row it replaced. It carries the six fields of D-66, and it
//     decides the D-27 threshold.
//   - A plain item holds one question the model left alone. It carries a
//     `warranted` field, and it decides whether the row deserved to fire.
//
// Every replacement goes in. The rest of the sheet is a sample spread
// over the rows, so one noisy row does not fill it and every row appears.
//
// Every tenth item repeats an earlier one, unlabeled, so the sheet
// measures the owner's self-consistency (D-66).
//
// Usage:
//
//	go run ./cmd/m5-sheet -n 60 ../docs/reference/pr7-question-gate-run14.md > ../docs/reference/pr7-m5-scoring-run14.md
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

// RubricVersion versions the fields. A change to them invalidates every
// score already given (D-66). Version 2 adds the `warranted` field, the
// three faults the owner used while scoring version 1, and the `n/a`
// value for a row that must not exist (D-104, D-114).
const RubricVersion = 2

// DefaultItems is the sheet size the owner chose on 2026-08-25. It keeps
// the scoring workload the same as version 1.
const DefaultItems = 60

// item is one question the agent asked.
type item struct {
	Run    string
	Conv   string
	Turn   int
	Row    string
	Slot   string
	Fit    float64
	Filled bool
	// Asked is the text that reached the user.
	Asked string
	// Catalog is the resolved catalog row behind the question.
	Catalog string
	// Invented is the replacement the model offered, or empty when it
	// offered none.
	Invented string
	// Refused marks a replacement the agent refused as a reword (D-88).
	// The catalog question went out in its place.
	Refused bool
	// Resolved is the catalog row before the ask role phrased it. The
	// reword guard compared the replacement against this text, so the
	// owner needs it to judge whether the refusal was right (D-116).
	Resolved string
}

// replacement reports whether the model offered another question here.
func (it item) replacement() bool { return strings.TrimSpace(it.Invented) != "" }

var (
	convRe = regexp.MustCompile(`(?m)^### (.+)$`)
	turnRe = regexp.MustCompile(`(?m)^\*\*Turn (\d+), the user:\*\*`)
	qRe    = regexp.MustCompile(`(?m)^- \[(INVENTED|catalog) slot=(\S+) row=(\S+) fit=([0-9.]+) filled=(\S+)\] (.+)$`)
	repRe  = regexp.MustCompile(`(?m)^  - It replaced: (.+)$`)
	refRe  = regexp.MustCompile(`(?m)^  - Refused as a reword \(D-88\): (.+)$`)
	resRe  = regexp.MustCompile(`(?m)^  - The guard compared against: (.+)$`)
)

func main() {
	n := flag.Int("n", DefaultItems, "how many items the sheet holds (0 holds every question)")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "give one or more gate documents")
		os.Exit(1)
	}
	found, err := collect(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	items, dropped, err := scorable(found)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(items) == 0 {
		fmt.Fprintln(os.Stderr, "no scorable question found in those documents")
		os.Exit(1)
	}
	write(os.Stdout, pick(items, *n), len(found), dropped)
}

// collect reads every question out of the given documents. A pattern is
// expanded here, not by the shell: the command runs with the Go module as
// its directory, and the shell expands a glob somewhere else.
func collect(patterns []string) ([]item, error) {
	var paths []string
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", pattern, err)
		}
		if len(matches) == 0 {
			return nil, fmt.Errorf("%s matches no file", pattern)
		}
		paths = append(paths, matches...)
	}
	sort.Strings(paths)
	var out []item
	for _, path := range paths {
		raw, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		run := strings.TrimSuffix(filepath.Base(path), ".md")
		conv, turn := "", 0
		for _, line := range strings.Split(string(raw), "\n") {
			switch {
			case convRe.MatchString(line):
				conv, turn = convRe.FindStringSubmatch(line)[1], 0
			case turnRe.MatchString(line):
				turn, _ = strconv.Atoi(turnRe.FindStringSubmatch(line)[1])
			case qRe.MatchString(line):
				m := qRe.FindStringSubmatch(line)
				fit, _ := strconv.ParseFloat(m[4], 64)
				it := item{
					Run: run, Conv: conv, Turn: turn, Slot: m[2], Row: m[3],
					Fit: fit, Filled: m[5] == "true", Asked: strings.TrimSpace(m[6]),
				}
				if m[1] == "INVENTED" {
					// The catalog side arrives on the "It replaced" line.
					it.Invented = it.Asked
				} else {
					it.Catalog = it.Asked
				}
				out = append(out, it)
			case repRe.MatchString(line) && len(out) > 0:
				out[len(out)-1].Catalog = strings.TrimSpace(repRe.FindStringSubmatch(line)[1])
			case refRe.MatchString(line) && len(out) > 0:
				// The catalog question went out, and this is what the model
				// offered in its place (D-88).
				out[len(out)-1].Invented = strings.TrimSpace(refRe.FindStringSubmatch(line)[1])
				out[len(out)-1].Refused = true
			case resRe.MatchString(line) && len(out) > 0:
				out[len(out)-1].Resolved = strings.TrimSpace(resRe.FindStringSubmatch(line)[1])
			}
		}
	}
	return out, nil
}

// scorable keeps the items the owner can actually score. A replacement
// with no catalog text goes, because the rubric compares the two. Any
// item whose row the catalog no longer holds goes too.
func scorable(items []item) (kept []item, dropped map[string]int, err error) {
	cat, err := questions.Load()
	if err != nil {
		return nil, nil, err
	}
	dropped = map[string]int{}
	for _, it := range items {
		if _, ok := cat.Row(it.Row); !ok {
			dropped["the catalog no longer holds the row"]++
			continue
		}
		if it.replacement() && strings.TrimSpace(it.Catalog) == "" {
			dropped["no catalog text recorded"]++
			continue
		}
		kept = append(kept, it)
	}
	return kept, dropped, nil
}

// pick builds the sheet. Every replacement goes in, because those decide
// the D-27 threshold and there are never many. The rest is a sample of
// the plain questions, spread so that every row appears at least once and
// no row fills the sheet (D-104).
func pick(items []item, n int) []item {
	var reps, plain []item
	for _, it := range items {
		if it.replacement() {
			reps = append(reps, it)
		} else {
			plain = append(plain, it)
		}
	}
	base := append([]item(nil), reps...)
	byRow := map[string][]item{}
	for _, it := range plain {
		byRow[it.Row] = append(byRow[it.Row], it)
	}
	rows := make([]string, 0, len(byRow))
	for row := range byRow {
		rows = append(rows, row)
	}
	sort.Strings(rows)
	// Round-robin over the rows. The first pass gives every row one item,
	// which is what makes a rare row reach the sheet.
	for i := 0; n <= 0 || len(base) < n; i++ {
		added := false
		for _, row := range rows {
			if i >= len(byRow[row]) {
				continue
			}
			base = append(base, byRow[row][i])
			added = true
			if n > 0 && len(base) >= n {
				break
			}
		}
		if !added {
			break
		}
	}
	// The repeats sit inside the count. The sheet used to be cut after
	// the repeats went in, which dropped real items off the end while
	// every repeat stayed (audit 2026-08-28).
	for n > 0 && len(withRepeats(base)) > n {
		base = base[:len(base)-1]
	}
	out := withRepeats(base)
	return out
}

// withRepeats inserts a repeat of an earlier item at every tenth
// position (D-66). It inserts rather than overwrites, so no real item is
// lost to the self-consistency check.
func withRepeats(items []item) []item {
	out := make([]item, 0, len(items)+len(items)/9+1)
	repeats := 0
	for _, it := range items {
		out = append(out, it)
		if len(out)%10 == 9 && repeats < len(items) {
			out = append(out, items[repeats])
			repeats++
		}
	}
	return out
}

func write(w io.Writer, items []item, total int, dropped map[string]int) {
	p := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	var sent, refused, plain int
	for _, it := range items {
		switch {
		case !it.replacement():
			plain++
		case it.Refused:
			refused++
		default:
			sent++
		}
	}
	p("# PR-7 M-5 scoring sheet\n\n")
	p("Rubric version %d (D-66, D-104). Prompt version %d. Items: %d, from %d questions across the gate runs.\n\n",
		RubricVersion, questions.PromptVersion, len(items), total)
	p("This sheet holds two kinds of item.\n\n")
	p("- **A replacement item** (%d of %d) pairs a question the model wrote with the catalog row it replaced. ", sent+refused, len(items))
	p("Of those, %d went out and %d were refused as rewords and never reached a user (D-88).\n", sent, refused)
	p("- **A plain item** (%d of %d) holds one question the model left alone. ", plain, len(items))
	p("Version 1 of this sheet could not hold such a question, and that is where most defects hid (D-104).\n\n")
	if len(dropped) > 0 {
		reasons := make([]string, 0, len(dropped))
		for reason := range dropped {
			reasons = append(reasons, reason)
		}
		sort.Strings(reasons)
		p("Not scorable, and left out:\n\n")
		for _, reason := range reasons {
			p("- %d, because %s.\n", dropped[reason], reason)
		}
		p("\n")
	}

	p("## What this sheet decides\n\n")
	p("Two things. The replacement items set the gap-score threshold: under it, the model may throw a catalog row away and write its own question (D-27). ")
	p("The plain items say whether our rows deserve to fire at all, which is the question version 1 could not ask.\n\n")

	p("## Fields on every item\n\n")
	p("### `right_slot`\n\n")
	p("Values: yes / no / n/a. Does the question ask about the slot named in the item header?\n\n")
	p("This is a break, not a matter of taste. A question filed under `theme` that asks about money sends the answer to the wrong field.\n\n")

	p("### `filled_slot` (already filled in)\n\n")
	p("Values: yes / partly / no. The run recorded it. Change it only if you disagree.\n\n")

	p("### `faults`\n\n")
	p("A list, not one value. Leave it blank only when nothing is wrong. Mandatory otherwise.\n\n")
	p("| Fault | Means |\n|---|---|\n")
	p("| duplicate | Asks what the user already answered |\n")
	p("| two questions in one | The user can not answer it in one reply |\n")
	p("| jargon | Uses words a player may not know |\n")
	p("| assumes an answer | Presumes something the user never said |\n")
	p("| unanswerable | Asks for a thing the user can not know |\n")
	p("| out of scope | Asks about a thing the app can not act on |\n")
	p("| vague | Does not say what kind of answer is useful (D-103) |\n")
	p("| inaccurate | States something that is not true of this user or this game (D-114) |\n")
	p("| omits information | Drops something the catalog row said, and says less for it (D-114) |\n\n")

	p("### `catalog_action`\n\n")
	p("Values: none / add / reword / remove. This field is how a scored item becomes a catalog change (D-25).\n\n")
	p("- `none`: leave the row alone.\n")
	p("- `reword`: the right row exists, and its wording caused the problem.\n")
	p("- `add`: no row covers this need at all.\n")
	p("- `remove`: this row must not exist. Use `n/a` in the other fields when you say this (D-114).\n\n")

	p("## Fields on a replacement item only\n\n")
	p("### `catalog_enough` (reads the catalog question)\n\n")
	p("Values: yes / no / unsure / n/a. Ignore the replacement. Read the catalog question alone. ")
	p("Had the agent simply asked that, would this conversation have been fine?\n\n")
	p("Use `unsure` freely. An unsure item is reported and left out of the calculation (D-66).\n\n")

	p("### `invented_better` (compares the two)\n\n")
	p("Values: worse / same / better / n/a. The only field that compares.\n\n")
	p("- `worse`: the replacement loses something the row had.\n")
	p("- `same`: no real difference. A replacement that only adds the user's colors or format is `same`.\n")
	p("- `better`: it serves this user better.\n\n")

	p("## The field on a plain item only\n\n")
	p("### `warranted` (reads the question)\n\n")
	p("Values: yes / no / unsure. Did this question deserve to be asked, to this user, at this point?\n\n")
	p("A `no` says the row fired when it should not have. That is a catalog or trigger defect, ")
	p("and it is what version 1 of this sheet could not record (D-104).\n\n")

	p("## Free text\n\n")
	p("`catalog_action` and `faults` accept free text after the keyword. A proposed wording is the useful part of `reword`. ")
	p("Put the keyword first, then a dash, then whatever you want:\n\n")
	p("    | catalog_action | reword - presumes a list, try \"Is there a card the deck must keep?\" |\n\n")
	p("Keep `catalog_enough`, `invented_better`, and `warranted` to their listed values. ")
	p("Those three are counted, and free text there would measure the reader rather than you (D-66).\n\n")

	p("## How your answers are used\n\n")
	p("An invention is warranted when `catalog_enough` is no and `invented_better` is not worse. ")
	p("The D-27 threshold is the gap score that best separates warranted from unwarranted. ")
	p("At least 80%% of the inventions a threshold allows must be warranted.\n\n")
	p("An `n/a` item counts toward no threshold. It says the row must not exist, which is a catalog change and not a scoring signal (D-114).\n\n")
	p("Some items repeat. That is on purpose, and it measures self-consistency. Do not look for them.\n\n")
	p("---\n\n")

	for i, it := range items {
		filled := "no"
		if it.Filled {
			filled = "yes"
		}
		p("## Item %d\n\n", i+1)
		p("Source: %s, conversation \"%s\", turn %d. Row `%s`, slot `%s`, gap score %.2f.\n\n",
			it.Run, it.Conv, it.Turn, it.Row, it.Slot, it.Fit)
		if !it.replacement() {
			p("**The question that went out:** %s\n\n", it.Asked)
			p("The model offered no replacement here. Score whether the question deserved to be asked.\n\n")
			p("| Field | Reads | Value |\n|---|---|---|\n")
			p("| warranted | the question | |\n")
			p("| right_slot | the question | |\n")
			p("| filled_slot | the run | %s |\n", filled)
			p("| faults | the question | |\n")
			p("| catalog_action | our catalog | |\n\n")
			continue
		}
		p("**The catalog question:** %s\n\n", orNone(it.Catalog))
		if it.Refused {
			p("**What the model offered instead:** %s\n\n", it.Invented)
			if it.Resolved != "" {
				p("**The guard compared against:** %s\n\n", it.Resolved)
				p("The line above is the catalog row before it was fitted to this user. ")
				p("The reword guard read that text, not the wording that went out (D-116).\n\n")
			}
			p("The agent refused this one as a reword and sent the catalog question (D-88). ")
			p("Score it as if it had gone out. A `better` here says the refusal was wrong.\n\n")
		} else {
			p("**What the model asked instead:** %s\n\n", it.Invented)
			p("The agent sent this one in place of the catalog question.\n\n")
		}
		p("| Field | Reads | Value |\n|---|---|---|\n")
		p("| catalog_enough | the catalog question | |\n")
		p("| invented_better | both | |\n")
		p("| right_slot | the replacement | |\n")
		p("| filled_slot | the run | %s |\n", filled)
		p("| faults | the replacement | |\n")
		p("| catalog_action | our catalog | |\n\n")
	}
}

func orNone(s string) string {
	if strings.TrimSpace(s) == "" {
		return "(the document did not record it: the run predates the catalog_text field)"
	}
	return s
}
