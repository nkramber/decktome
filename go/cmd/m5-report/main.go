// Command m5-report reads a scored M-5 sheet and computes the two
// thresholds the owner's answers decide (D-27, D-66, D-88).
//
// It counts, it does not judge. An item is warranted when
// catalog_enough is "no" and invented_better is not "worse" (D-66). For
// each candidate gap score it reports how many inventions that threshold
// would allow and what share of them were warranted. The D-27 threshold
// is the highest one that keeps the share at or above the precision
// floor.
//
// It reads only the first word of each field, so free text after a dash
// passes through untouched (D-100).
//
// Usage:
//
//	go run ./cmd/m5-report ../docs/reference/pr7-m5-scoring.md
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

// PrecisionFloor is the D-66 rule: at least this share of the inventions
// a threshold allows must be warranted.
const PrecisionFloor = 0.80

// scored is one item after the owner filled it in.
type scored struct {
	Item     int
	Row      string
	Fit      float64
	Refused  bool
	Enough   string
	Better   string
	Slot     string
	Faults   string
	Action   string
	Complete bool
}

// warranted is the D-66 rule. An unsure item is neither warranted nor
// unwarranted, and the caller leaves it out.
func (s scored) warranted() bool { return s.Enough == "no" && s.Better != "worse" }
func (s scored) unsure() bool    { return s.Enough == "unsure" || s.Enough == "" }

var (
	itemRe   = regexp.MustCompile(`(?m)^## Item (\d+)$`)
	sourceRe = regexp.MustCompile(`Row ` + "`" + `(\w+)` + "`" + `.*gap score ([0-9.]+)`)
	fieldRe  = regexp.MustCompile(`(?m)^\| (\w+) \| [^|]* \| ([^|]*)\|`)
)

func main() {
	floor := flag.Float64("floor", PrecisionFloor, "precision floor for a candidate threshold")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "give one scored sheet")
		os.Exit(1)
	}
	raw, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	items := parse(string(raw))
	if len(items) == 0 {
		fmt.Fprintln(os.Stderr, "no items in that sheet")
		os.Exit(1)
	}
	report(os.Stdout, items, *floor)
}

// parse reads every item and the fields the owner filled in. It takes the
// first word of a field, so free text after it is ignored here and kept
// in the sheet.
func parse(doc string) []scored {
	var out []scored
	locs := itemRe.FindAllStringSubmatchIndex(doc, -1)
	for i, loc := range locs {
		end := len(doc)
		if i+1 < len(locs) {
			end = locs[i+1][0]
		}
		block := doc[loc[0]:end]
		n, _ := strconv.Atoi(doc[loc[2]:loc[3]])
		it := scored{Item: n, Refused: strings.Contains(block, "refused this one as a reword")}
		if m := sourceRe.FindStringSubmatch(block); m != nil {
			it.Row = m[1]
			it.Fit, _ = strconv.ParseFloat(strings.TrimSuffix(m[2], "."), 64)
		}
		for _, f := range fieldRe.FindAllStringSubmatch(block, -1) {
			value := firstWord(f[2])
			switch f[1] {
			case "catalog_enough":
				it.Enough = value
			case "invented_better":
				it.Better = value
			case "right_slot":
				it.Slot = value
			case "faults":
				it.Faults = strings.TrimSpace(f[2])
			case "catalog_action":
				it.Action = value
			}
		}
		it.Complete = it.Enough != "" && it.Better != ""
		out = append(out, it)
	}
	return out
}

// firstWord lowercases a field and takes the word before any separator,
// so "reword - presumes a list" reads as "reword".
func firstWord(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	for _, cut := range []string{" - ", " — ", ", ", ": "} {
		if i := strings.Index(s, cut); i >= 0 {
			s = s[:i]
		}
	}
	return strings.TrimSpace(strings.Trim(s, ".,"))
}

func report(w io.Writer, items []scored, floor float64) {
	p := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	var sent, refused, done, unsure int
	for _, it := range items {
		if it.Refused {
			refused++
		} else {
			sent++
		}
		switch {
		case !it.Complete:
		case it.unsure():
			unsure++
			done++
		default:
			done++
		}
	}
	p("# M-5 report\n\n")
	p("Items: %d. Scored: %d. Not scored yet: %d.\n\n", len(items), done, len(items)-done)
	p("Sent to a user: %d. Refused as rewords: %d (D-88).\n\n", sent, refused)
	if done < len(items) {
		p("**The sheet is not finished. Every number below reads the scored items only.**\n\n")
	}
	if unsure > 0 {
		p("%s unsure, and D-66 leaves %s out of the fit.\n\n", plural(unsure, "item is", "items are"), them(unsure))
	}

	// The D-27 threshold reads the questions that went out. A refused
	// reword never reached a user, so it can not tell us what a threshold
	// would have allowed.
	p("## The fit threshold (D-27)\n\n")
	p("A threshold allows every invention whose gap score is under it. ")
	p("An invention is warranted when the catalog was not enough and the replacement is not worse.\n\n")
	// The fit threshold reads the questions that went out. Count them
	// first, because "no threshold holds" and "nothing is scored yet" are
	// different answers and must not print the same line.
	usable := 0
	for _, it := range items {
		if !it.Refused && it.Complete && !it.unsure() {
			usable++
		}
	}
	if usable == 0 {
		p("No sent invention is scored yet, so no threshold follows. ")
		p("A refused reword never reached a user, so it can not say what a threshold would have allowed.\n\n")
	} else {
		p("| Threshold | Allowed | Warranted | Precision | Meets the %.0f%% floor |\n|---|---|---|---|---|\n", floor*100)
	}
	best, bestOK := 0.0, false
	for _, t := range candidates(items) {
		if usable == 0 {
			break
		}
		allowed, ok := 0, 0
		for _, it := range items {
			if it.Refused || !it.Complete || it.unsure() || it.Fit >= t {
				continue
			}
			allowed++
			if it.warranted() {
				ok++
			}
		}
		precision, meets := 0.0, false
		if allowed > 0 {
			precision = float64(ok) / float64(allowed)
			meets = precision >= floor
		}
		mark := "no"
		if allowed == 0 {
			mark = "nothing allowed"
		} else if meets {
			mark = "yes"
			if t > best {
				best, bestOK = t, true
			}
		}
		p("| %.2f | %d | %d | %.0f%% | %s |\n", t, allowed, ok, precision*100, mark)
	}
	p("\n")
	switch {
	case usable == 0:
		p("Score an item the agent sent, and this table fills in.\n\n")
	case bestOK:
		p("**Set the fit threshold to %.2f.** It is the highest candidate that keeps the precision floor.\n\n", best)
	default:
		p("**No candidate meets the floor.** Every threshold lets through more unwarranted inventions than D-66 allows. ")
		p("The catalog rows the model replaced need work before a threshold is worth setting.\n\n")
	}

	// The reword guard reads the refused items alone.
	p("## The reword guard (D-88)\n\n")
	better, total := 0, 0
	for _, it := range items {
		if !it.Refused || !it.Complete || it.unsure() {
			continue
		}
		total++
		if it.Better == "better" {
			better++
		}
	}
	switch {
	case total == 0:
		p("No refused reword is scored yet.\n\n")
	case better == 0:
		p("None of the %s better than the row it replaced. The guard at %.2f overlap is not too tight.\n\n",
			plural(total, "scored refusal is", "scored refusals are"), questions.MaxRewordOverlap)
	default:
		p("**%d of %d scored refusals are better than the row they replaced.** The guard at %.2f overlap suppressed a real improvement, so it is too tight.\n\n",
			better, total, questions.MaxRewordOverlap)
	}

	// What the catalog needs.
	p("## What the catalog needs\n\n")
	byAction := map[string][]string{}
	for _, it := range items {
		if it.Action == "" || it.Action == "none" {
			continue
		}
		byAction[it.Action] = append(byAction[it.Action], it.Row)
	}
	if len(byAction) == 0 {
		p("No item asks for a catalog change yet.\n\n")
		return
	}
	actions := make([]string, 0, len(byAction))
	for a := range byAction {
		actions = append(actions, a)
	}
	sort.Strings(actions)
	for _, a := range actions {
		counts := map[string]int{}
		for _, row := range byAction[a] {
			counts[row]++
		}
		rows := make([]string, 0, len(counts))
		for row := range counts {
			rows = append(rows, row)
		}
		sort.Slice(rows, func(i, j int) bool {
			if counts[rows[i]] != counts[rows[j]] {
				return counts[rows[i]] > counts[rows[j]]
			}
			return rows[i] < rows[j]
		})
		p("**%s** (%s):", a, plural(len(byAction[a]), "item", "items"))
		for _, row := range rows {
			p(" `%s` x%d", row, counts[row])
		}
		p("\n\n")
	}
	p("The owner approves every catalog change (D-28).\n")
}

// plural writes a count with the word that fits it.
func plural(n int, one, many string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, one)
	}
	return fmt.Sprintf("%d %s", n, many)
}

// them writes the pronoun that fits a count.
func them(n int) string {
	if n == 1 {
		return "it"
	}
	return "them"
}

// candidates are the thresholds worth testing: just above each gap score
// the sheet holds, so every distinct group is tried.
func candidates(items []scored) []float64 {
	seen := map[float64]bool{}
	for _, it := range items {
		if !it.Refused {
			seen[it.Fit] = true
		}
	}
	var out []float64
	for fit := range seen {
		out = append(out, fit+0.01)
	}
	out = append(out, 0.35)
	sort.Float64s(out)
	return out
}
