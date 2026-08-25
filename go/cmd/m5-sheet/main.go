// Command m5-sheet builds the M-5 scoring sheet from one or more gate
// documents (D-66). Each row pairs an invented question with the catalog
// question it replaced, which is what the rubric needs: the owner can not
// judge catalog_enough or invented_better from one text alone.
//
// Every tenth item repeats an earlier one, unlabeled, so the sheet
// measures the owner's self-consistency (D-66).
//
// Usage:
//
//	go run ./cmd/m5-sheet -n 20 ../docs/reference/pr7-question-gate*.md > ../docs/reference/pr7-m5-scoring.md
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

// RubricVersion versions the six fields. A change to them invalidates
// every score already given (D-66).
const RubricVersion = 1

// item is one replacement the model offered, with the row it replaced.
// The agent either sent it (Refused false) or kept the catalog row and
// refused it as a reword (Refused true, D-88).
//
// A refused item is scored the same way. The rubric asks whether the
// catalog was enough and whether the replacement is better, and those
// two questions decide whether the refusal was right.
type item struct {
	Run      string
	Conv     string
	Turn     int
	Row      string
	Slot     string
	Fit      float64
	Filled   bool
	Invented string
	Catalog  string
	Refused  bool
}

var (
	convRe = regexp.MustCompile(`(?m)^### (.+)$`)
	turnRe = regexp.MustCompile(`(?m)^\*\*Turn (\d+), the user:\*\*`)
	qRe    = regexp.MustCompile(`(?m)^- \[(INVENTED|catalog) slot=(\S+) row=(\S+) fit=([0-9.]+) filled=(\S+)\] (.+)$`)
	repRe  = regexp.MustCompile(`(?m)^  - It replaced: (.+)$`)
	refRe  = regexp.MustCompile(`(?m)^  - Refused as a reword \(D-88\): (.+)$`)
)

func main() {
	n := flag.Int("n", 0, "cap the sample (0 scores every replacement)")
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
		fmt.Fprintln(os.Stderr, "no scorable invented question found in those documents")
		os.Exit(1)
	}
	write(os.Stdout, pick(items, *n), len(found), dropped)
}

// collect reads every invented question out of the given documents. A
// pattern is expanded here, not by the shell: the command runs with the
// Go module as its directory, and the shell expands a glob somewhere
// else.
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
		text := string(raw)
		// Walk the file once, tracking the conversation and turn each
		// question sits under.
		conv, turn := "", 0
		for _, line := range strings.Split(text, "\n") {
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
					Fit: fit, Filled: m[5] == "true",
				}
				if m[1] == "INVENTED" {
					it.Invented = strings.TrimSpace(m[6])
				} else {
					// A catalog question. It becomes an item only when a
					// refusal line follows, and then its own text is the
					// catalog side of the pair.
					it.Catalog, it.Refused = strings.TrimSpace(m[6]), true
				}
				out = append(out, it)
			case repRe.MatchString(line) && len(out) > 0:
				out[len(out)-1].Catalog = strings.TrimSpace(repRe.FindStringSubmatch(line)[1])
			case refRe.MatchString(line) && len(out) > 0:
				out[len(out)-1].Invented = strings.TrimSpace(refRe.FindStringSubmatch(line)[1])
			}
		}
	}
	return out, nil
}

// scorable keeps the items the owner can actually score. Two kinds go:
// an item with no catalog text, because the rubric compares the two
// questions, and an item whose row the catalog no longer holds.
func scorable(items []item) (kept []item, dropped map[string]int, err error) {
	cat, err := questions.Load()
	if err != nil {
		return nil, nil, err
	}
	dropped = map[string]int{}
	for _, it := range items {
		switch {
		case strings.TrimSpace(it.Invented) == "":
			// A catalog question the model left alone. That is the normal
			// case, and it is not a dropped item.
			continue
		case strings.TrimSpace(it.Catalog) == "":
			dropped["no catalog text recorded"]++
		default:
			if _, ok := cat.Row(it.Row); !ok {
				dropped["the catalog no longer holds the row"]++
				continue
			}
			kept = append(kept, it)
		}
	}
	return kept, dropped, nil
}

// pick spreads the sample over the rows, so one noisy row does not fill
// the sheet. Every tenth item repeats an earlier one (D-66).
func pick(items []item, n int) []item {
	byRow := map[string][]item{}
	for _, it := range items {
		byRow[it.Row] = append(byRow[it.Row], it)
	}
	rows := make([]string, 0, len(byRow))
	for row := range byRow {
		rows = append(rows, row)
	}
	sort.Strings(rows)
	// Round-robin over the rows, so one noisy row does not fill the sheet.
	var out []item
	for i := 0; n <= 0 || len(out) < n; i++ {
		added := false
		for _, row := range rows {
			if i < len(byRow[row]) {
				out = append(out, byRow[row][i])
				added = true
				if n > 0 && len(out) == n {
					break
				}
			}
		}
		if !added {
			break
		}
	}
	return withRepeats(out)
}

// withRepeats inserts a repeat of an earlier item at every tenth
// position (D-66). It inserts rather than overwrites, so no real item
// is lost to the self-consistency check.
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
	p("# PR-7 M-5 scoring sheet\n\n")
	sent, refused := 0, 0
	for _, it := range items {
		if it.Refused {
			refused++
		} else {
			sent++
		}
	}
	p("Rubric version %d (D-66). Items: %d, from %d questions across the gate runs.\n\n", RubricVersion, len(items), total)
	p("A question the model left alone is not scored. Only a question it offered to replace is.\n\n")
	p("%d went out in place of a catalog question. %d were refused as rewords and never reached a user (D-88). ", sent, refused)
	p("Both kinds are scored the same way. The refused ones decide whether the reword guard is set right.\n\n")
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
	p("The agent scores every catalog question from 0 to 1 for how well it fits the user. ")
	p("Under a threshold, the model may throw the row away and write its own question. ")
	p("That threshold is a guess until you score these items. Your answers replace the guess (D-27).\n\n")

	p("## How to read an item\n\n")
	p("Each item holds two questions. **The catalog question** is the fixed row the agent planned to ask. ")
	p("**What the model asked instead** is the replacement it wrote. Every field below reads one of the two, and the list says which.\n\n")

	p("## The six fields\n\n")
	p("### `catalog_enough` (reads the catalog question)\n\n")
	p("Values: yes / no / unsure. Ignore the replacement. Read the catalog question alone. ")
	p("Had the agent simply asked that, would this conversation have been fine?\n\n")
	p("Use `unsure` freely. An unsure item is reported and left out of the calculation (D-66).\n\n")

	p("### `invented_better` (compares the two)\n\n")
	p("Values: worse / same / better. The only field that compares.\n\n")
	p("- `worse`: the replacement loses something the row had.\n")
	p("- `same`: no real difference. A replacement that only adds the user's colors or format is `same`.\n")
	p("- `better`: it serves this user better.\n\n")

	p("### `right_slot` (reads the replacement)\n\n")
	p("Values: yes / no. Does the replacement ask about the slot named in the item header?\n\n")
	p("This is a break, not a matter of taste. A question filed under `theme` that asks about money sends the answer to the wrong field.\n\n")

	p("### `filled_slot` (already filled in)\n\n")
	p("Values: yes / partly / no. The run recorded it. Change it only if you disagree.\n\n")
	p("It means one of two things. On an item the agent sent, it says whether the replacement closed the slot. ")
	p("On a refused item the replacement never went out, so it says whether the catalog question closed the slot.\n\n")

	p("### `faults` (reads the replacement)\n\n")
	p("A list, not one value. Leave it blank only when nothing is wrong. Mandatory otherwise.\n\n")
	p("| Fault | Means |\n|---|---|\n")
	p("| duplicate | Asks what the user already answered |\n")
	p("| two questions in one | The user can not answer it in one reply |\n")
	p("| jargon | Uses words a player may not know |\n")
	p("| assumes an answer | Presumes something the user never said |\n")
	p("| unanswerable | Asks for a thing the user can not know |\n")
	p("| out of scope | Asks about a thing the app can not act on |\n")
	p("| vague | Does not say what kind of answer is useful (D-103) |\n\n")

	p("### `catalog_action` (reads our catalog)\n\n")
	p("Values: none / add / reword. Not about the replacement. This field is how a scored item becomes a catalog change (D-25).\n\n")
	p("- `none`: leave the row alone.\n")
	p("- `reword`: the right row exists, and its wording caused the problem.\n")
	p("- `add`: no row covers this need at all.\n\n")

	p("## Free text\n\n")
	p("`catalog_action` and `faults` accept free text after the keyword, and a proposed wording is the useful part of `reword`. ")
	p("Put the keyword first, then a dash, then whatever you want:\n\n")
	p("    | catalog_action | reword - presumes a list, try \"Is there a card the deck must keep?\" |\n\n")
	p("Keep `catalog_enough` and `invented_better` to their listed values. ")
	p("Those two are counted, and free text there would measure the reader rather than you (D-66).\n\n")

	p("## How your answers are used\n\n")
	p("An invention is warranted when `catalog_enough` is no and `invented_better` is not worse. ")
	p("The D-27 threshold is the gap score that best separates warranted from unwarranted. ")
	p("At least 80%% of the inventions a threshold allows must be warranted.\n\n")
	p("Every item here scored 0.05 or 0.20, because the scorer counts faults rather than shades of fit. ")
	p("A 0.05 means it found two faults or more, and a 0.20 means one clear fault. ")
	p("So the threshold question is nearly a yes or no: does the 0.20 group deserve to invent?\n\n")
	p("Some items repeat. That is on purpose, and it measures self-consistency. Do not look for them.\n\n")
	p("---\n\n")
	for i, it := range items {
		filled := "no"
		if it.Filled {
			filled = "yes"
		}
		p("## Item %d\n\n", i+1)
		p("Source: %s, conversation \"%s\", turn %d. Row `%s`, slot `%s`, gap score %.2f.\n\n", it.Run, it.Conv, it.Turn, it.Row, it.Slot, it.Fit)
		p("**The catalog question:** %s\n\n", orNone(it.Catalog))
		if it.Refused {
			p("**What the model offered instead:** %s\n\n", it.Invented)
			p("The agent refused this one as a reword and sent the catalog question (D-88). Score it as if it had gone out. A `better` here says the refusal was wrong.\n\n")
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
