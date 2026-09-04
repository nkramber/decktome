package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
)

// The lines of a deck gate document the import reads. They are the
// lines cmd/deck-gate writes, and a document from before a line existed
// gives no row for it.
var (
	importDate    = regexp.MustCompile(`^Run date: (\S+)\. Card snapshot: (\S+)\.`)
	importVerdict = regexp.MustCompile(`^Verdict: (PASS|FAIL)\.`)
	importTable   = regexp.MustCompile(`^\| (Prompt version|Calls|Cost|Time) \| ([^|]+?) \|$`)
	importDeck    = regexp.MustCompile(`^### (\d+)\. (.+)$`)
	importError   = regexp.MustCompile(`^ERROR: (.+)$`)
	importGrade   = regexp.MustCompile(`^Grade: (\w+), score ([0-9.]+), model (\S+)\.`)
	importCards   = regexp.MustCompile(`^Cards: (\d+) main, (\d+) sideboard\. Repair turn: (.+?)\. Block findings: (\d+)\.`)
	importCost    = regexp.MustCompile(`^Cost: \$([0-9.]+) to buy, \$([0-9.]+) the whole deck\.`)
	importNote    = regexp.MustCompile(`^- NOTE: (.+)$`)
	importJudge   = regexp.MustCompile(`^- JUDGE \[(true|false|unknown)\]:`)
	importJudgeEr = regexp.MustCompile(`^- JUDGE ERROR: (.+)$`)
	importFinding = regexp.MustCompile("^- \\[(BLOCK|WARN|INFO)\\] `([a-z_]+)`:")
	importPool    = regexp.MustCompile(`Shortlist: (\d+) names\.`)
)

// importDeckGate reads a deck gate document into a run of the decks
// suite. The roles are unknown: a document names no model, which is the
// gap the run files close (PR-15).
func importDeckGate(r io.Reader, runID, source string) (*evalrun.Run, error) {
	run := evalrun.New("decks", runID)
	run.Header.Commit = ""
	run.Header.Note = "imported from " + source + ", so the roles are unknown"
	run.LowerIsBetter("blocks", "invented_names", "false_rules", "judge_error", "warnings", "repaired", "buy_cost", "deck_cost")

	type deck struct {
		item                          string
		err                           string
		cards, blocks, warnings, pool int
		repaired                      bool
		repairReason                  string
		notes, codes                  []string
		falseRules, judged            int
		judgeErr                      string
		buy, cost, score              float64
		tier, model                   string
		seen                          bool
	}
	var decks []*deck
	var cur *deck
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64<<10), 8<<20)
	for sc.Scan() {
		line := sc.Text()
		if cur == nil {
			if m := importDate.FindStringSubmatch(line); m != nil {
				run.Header.Date, run.Header.Snapshot = m[1], m[2]
				continue
			}
			if m := importVerdict.FindStringSubmatch(line); m != nil {
				run.Header.Verdict = m[1]
				continue
			}
			if m := importTable.FindStringSubmatch(line); m != nil {
				readHeaderCell(run, m[1], strings.TrimSpace(m[2]))
				continue
			}
		}
		if m := importDeck.FindStringSubmatch(line); m != nil {
			cur = &deck{item: m[1], seen: true}
			decks = append(decks, cur)
			continue
		}
		if cur == nil {
			continue
		}
		switch {
		case importError.MatchString(line):
			cur.err = importError.FindStringSubmatch(line)[1]
		case importPool.MatchString(line):
			cur.pool, _ = strconv.Atoi(importPool.FindStringSubmatch(line)[1])
		case importGrade.MatchString(line):
			m := importGrade.FindStringSubmatch(line)
			cur.tier, cur.model = m[1], m[3]
			cur.score, _ = strconv.ParseFloat(m[2], 64)
		case importCards.MatchString(line):
			m := importCards.FindStringSubmatch(line)
			cur.cards, _ = strconv.Atoi(m[1])
			cur.blocks, _ = strconv.Atoi(m[4])
			repair := m[3]
			cur.repaired = repair != "no" && repair != "false"
			cur.repairReason = strings.TrimPrefix(repair, "yes, for ")
			if !cur.repaired {
				cur.repairReason = ""
			}
		case importCost.MatchString(line):
			m := importCost.FindStringSubmatch(line)
			cur.buy, _ = strconv.ParseFloat(m[1], 64)
			cur.cost, _ = strconv.ParseFloat(m[2], 64)
		case importNote.MatchString(line):
			cur.notes = append(cur.notes, importNote.FindStringSubmatch(line)[1])
		case importJudgeEr.MatchString(line):
			cur.judgeErr = importJudgeEr.FindStringSubmatch(line)[1]
		case importJudge.MatchString(line):
			cur.judged++
			if importJudge.FindStringSubmatch(line)[1] == "false" {
				cur.falseRules++
			}
		case importFinding.MatchString(line):
			m := importFinding.FindStringSubmatch(line)
			switch m[1] {
			case "BLOCK":
				cur.codes = append(cur.codes, m[2])
			case "WARN":
				cur.warnings++
			}
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if len(decks) == 0 {
		return nil, fmt.Errorf("no deck section in the document")
	}
	if run.Header.Verdict == "" {
		return nil, fmt.Errorf("no verdict line in the document")
	}
	for _, d := range decks {
		if d.err != "" {
			run.Gate(d.item, "built", 0, d.err)
			continue
		}
		run.Gate(d.item, "built", 1, "")
		run.Gate(d.item, "blocks", float64(d.blocks), strings.Join(d.codes, ", "))
		run.Gate(d.item, "invented_names", float64(len(d.notes)), strings.Join(d.notes, " | "))
		switch {
		case d.judgeErr != "":
			run.Gate(d.item, "judge_error", 1, d.judgeErr)
		default:
			run.Gate(d.item, "false_rules", float64(d.falseRules), "")
			run.Info(d.item, "rules_claims", float64(d.judged), "")
		}
		repaired := 0.0
		if d.repaired {
			repaired = 1
		}
		run.Info(d.item, "repaired", repaired, d.repairReason)
		run.Info(d.item, "cards", float64(d.cards), "")
		run.Info(d.item, "pool", float64(d.pool), "")
		run.Info(d.item, "warnings", float64(d.warnings), "")
		run.Info(d.item, "buy_cost", d.buy, "")
		run.Info(d.item, "deck_cost", d.cost, "")
		if d.tier != "" {
			run.Info(d.item, "grade", d.score, d.tier)
			if run.Header.Versions["quality_model"] == "" {
				run.Header.Versions["quality_model"] = d.model
			}
		}
	}
	return run, nil
}

// readHeaderCell reads one row of the summary table into the header.
func readHeaderCell(run *evalrun.Run, name, value string) {
	switch name {
	case "Prompt version":
		if n, err := strconv.Atoi(value); err == nil {
			run.Header.Prompts["generate"] = n
		}
	case "Calls":
		run.Header.Calls, _ = strconv.Atoi(value)
	case "Cost":
		if v, err := strconv.ParseFloat(strings.TrimPrefix(value, "$"), 64); err == nil {
			run.Header.CostUSD = &v
		}
	case "Time":
		if v, err := strconv.ParseFloat(strings.TrimSuffix(value, " seconds"), 64); err == nil {
			run.Header.Seconds = v
		}
	}
}
