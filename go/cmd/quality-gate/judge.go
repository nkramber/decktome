package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/evalrun"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/generate"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/meta"
	"github.com/nkramber/mtg-deck-builder/go/internal/profile"
	"github.com/nkramber/mtg-deck-builder/go/internal/quality"
	"github.com/nkramber/mtg-deck-builder/go/internal/rules"
)

// The judge lane reads the decks of a deck gate document and asks the
// judge role for the tier of each. The bar: the judge agrees with the
// grade the summary names in eight of ten (roadmap PR-14B).
//
// CAUTION: this calls a real provider and it costs money. QUALITY_JUDGE=1
// is required, so it can not run by accident. Each deck is one judge
// call, a few cents on the judge role (D-430).

// judgeBar is the share of decks the judge must agree on.
const judgeBar = 0.8

var (
	judgeDeckHeader = regexp.MustCompile(`^### (\d+)\. (.+)$`)
	judgeFormat     = regexp.MustCompile(`^Format: ([A-Za-z]+)\.`)
	judgeCommander  = regexp.MustCompile(`^Commander: (.+)\.$`)
	judgeGrade      = regexp.MustCompile(`The quality model grades this deck (.+?)( against the top lists of the format)?[:.]`)
	judgeCard       = regexp.MustCompile(`^- (\d+) (.+?) \| `)
)

// judged is one deck of the document with the judge's answer.
type judged struct {
	id     int
	title  string
	format mtgv1.FormatId
	deck   *mtgv1.Deck
	// grade is the tier the document's summary names, and modelGrade
	// the tier the stored model gives the same deck now. The bar reads
	// the model's, so a refit needs no paid deck gate run.
	grade      string
	modelGrade string
	judgement  *generate.TierJudgement
	err        error
}

// promptCommanders reads the commander of each prompt of the deck gate,
// for a document written before the report carried a commander line.
func promptCommanders(path string) (map[int][]string, error) {
	if path == "" {
		return nil, nil
	}
	data, err := os.ReadFile(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, err
	}
	var doc struct {
		Prompts []struct {
			ID        int    `json:"id"`
			Commander string `json:"commander"`
		} `json:"prompts"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("prompts %s: %w", path, err)
	}
	out := map[int][]string{}
	for _, p := range doc.Prompts {
		if p.Commander != "" {
			out[p.ID] = strings.Split(p.Commander, " + ")
		}
	}
	return out, nil
}

// readDeckGate parses the deck sections of a deck gate document: the
// header with the id and the title, the format line, the commander line
// when the report wrote one, the grade the summary names, and the card
// list. A deck with no grade is skipped, because the bar reads the
// agreement with a grade.
func readDeckGate(r io.Reader, idx *cards.Index, commanders map[int][]string) ([]judged, error) {
	var out []judged
	var cur *judged
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		line := sc.Text()
		if m := judgeDeckHeader.FindStringSubmatch(line); m != nil {
			id, _ := strconv.Atoi(m[1])
			out = append(out, judged{id: id, title: m[2], deck: &mtgv1.Deck{}})
			cur = &out[len(out)-1]
			continue
		}
		if cur == nil {
			continue
		}
		if m := judgeFormat.FindStringSubmatch(line); m != nil {
			cur.format = gatekit.FormatID(strings.ToLower(m[1]))
			cur.deck.Format = &mtgv1.Format{Id: cur.format}
			continue
		}
		if m := judgeCommander.FindStringSubmatch(line); m != nil {
			for _, name := range strings.Split(m[1], ", ") {
				if c, ok := idx.ByName(name); ok {
					cur.deck.CommanderOracleIds = append(cur.deck.CommanderOracleIds, c.GetOracleId())
				}
			}
			continue
		}
		if m := judgeGrade.FindStringSubmatch(line); m != nil && cur.grade == "" {
			cur.grade = tierOfWords(m[1])
			continue
		}
		if m := judgeCard.FindStringSubmatch(line); m != nil {
			n, _ := strconv.Atoi(m[1])
			c, ok := idx.ByName(m[2])
			if !ok {
				return nil, fmt.Errorf("deck %d: the index does not know %q", cur.id, m[2])
			}
			cur.deck.Cards = append(cur.deck.Cards, &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: int32(n)})
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	var kept []judged
	for _, j := range out {
		if len(j.deck.CommanderOracleIds) == 0 && j.format == mtgv1.FormatId_FORMAT_ID_COMMANDER {
			for _, name := range commanders[j.id] {
				if c, ok := idx.ByName(name); ok {
					j.deck.CommanderOracleIds = append(j.deck.CommanderOracleIds, c.GetOracleId())
				}
			}
		}
		if j.grade == "" || len(j.deck.Cards) == 0 {
			continue
		}
		kept = append(kept, j)
	}
	if len(kept) == 0 {
		return nil, errors.New("the document holds no graded deck")
	}
	return kept, nil
}

// tierOfWords reads the grade word of a summary back to its tier.
func tierOfWords(words string) string {
	switch strings.TrimSpace(words) {
	case "at the precon baseline":
		return meta.TierBaseline
	case "below the precon baseline":
		return meta.TierBad
	default:
		w := strings.TrimSpace(words)
		if meta.TierLevel(w) >= 0 {
			return w
		}
		return ""
	}
}

// runJudge runs the judge lane over a deck gate document and writes
// the judge document to stdout.
func runJudge(path, promptsPath, runOut string) error {
	if err := gatekit.SpendGuard("QUALITY_JUDGE"); err != nil {
		return err
	}
	// The run file is never overwritten (D-65), and the check runs before
	// the first provider call.
	if err := gatekit.RefuseExisting(runOut); err != nil {
		return err
	}
	run := evalrun.New("tier-judge", evalrun.RunID(runOut))
	run.Header.Versions["source"] = filepath.Base(path)
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(context.Background(), quiet)
	if err != nil {
		return err
	}
	run.SetSnapshot(idx.AsOf)
	commanders, err := promptCommanders(promptsPath)
	if err != nil {
		return err
	}
	f, err := os.Open(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	decks, err := readDeckGate(f, idx, commanders)
	if err != nil {
		return fmt.Errorf("judge %s: %w", path, err)
	}
	modelVersion, err := regrade(decks, idx)
	if err != nil {
		return err
	}
	run.Header.Versions["quality_model"] = modelVersion
	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return err
	}
	run.SetRoles(client.Config(), llm.RoleJudge)
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	acc := llm.NewAccumulator(prices)
	start := time.Now()
	for i := range decks {
		d := &decks[i]
		d.judgement, d.err = generate.JudgeTier(context.Background(), client, d.deck, idx, acc)
		word := "error"
		if d.judgement != nil {
			word = "judged " + d.judgement.Tier
		}
		fmt.Fprintf(os.Stderr, "  %2d. document %-9s model %-9s %s\n", d.id, d.grade, d.modelGrade, word)
	}
	pass := reportJudge(os.Stdout, path, decks, acc, idx, time.Since(start), run)
	if err := evalrun.WriteFile(runOut, run); err != nil {
		return err
	}
	if !pass {
		return errors.New("quality judge: FAIL")
	}
	return nil
}

// regrade grades every deck with the stored model: the profile with no
// content check, then the scorer. The roles come from the candidate
// builder, as the fit reads them.
func regrade(decks []judged, idx *cards.Index) (modelVersion string, err error) {
	scorer, err := gatekit.Scorer(context.Background())
	if err != nil {
		return "", err
	}
	if scorer.Version() == "" {
		return "", errors.New("no stored quality model, so the judge lane has no grade to check")
	}
	cfg, err := rules.Load()
	if err != nil {
		return "", err
	}
	prof, err := profile.New(cfg, func() *cards.TagIndex { return idx.Tags() }, nil)
	if err != nil {
		return "", err
	}
	builder, err := candidates.New()
	if err != nil {
		return "", err
	}
	roles := builder.Roles(idx)
	for i := range decks {
		d := &decks[i]
		for _, dc := range d.deck.GetCards() {
			if c, ok := idx.ByOracleID(dc.GetOracleId()); ok {
				dc.Role = roles(c)
			}
		}
		q := scorer.Score(quality.Input{Deck: d.deck, Profile: prof.Measure(d.deck, idx), Cards: idx})
		if q == nil {
			return "", fmt.Errorf("deck %d: the model covers no %s", d.id, generate.FormatWord(d.format))
		}
		d.modelGrade = q.GetTier()
	}
	return scorer.Version(), nil
}

// reportJudge writes the judge document and answers the verdict.
func reportJudge(w io.Writer, source string, decks []judged, acc *llm.Accumulator, idx *cards.Index, took time.Duration, run *evalrun.Run) bool {
	p := func(format string, a ...any) { _, _ = fmt.Fprintf(w, format, a...) }
	agreed, judgedCount, errs := 0, 0, 0
	// Off by one rung counts apart, so a reader sees a near miss.
	near := 0
	for _, d := range decks {
		if d.err != nil {
			errs++
			continue
		}
		judgedCount++
		if d.judgement.Tier == d.modelGrade {
			agreed++
		} else if diff := meta.TierLevel(d.judgement.Tier) - meta.TierLevel(d.modelGrade); diff == 1 || diff == -1 {
			near++
		}
	}
	share := 0.0
	if judgedCount > 0 {
		share = float64(agreed) / float64(judgedCount)
	}
	pass := errs == 0 && judgedCount > 0 && share >= judgeBar
	verdict := "PASS"
	if !pass {
		verdict = "FAIL"
	}
	for _, d := range decks {
		item := fmt.Sprintf("%d", d.id)
		if d.err != nil {
			run.Gate(item, "judged", 0, d.err.Error())
			continue
		}
		run.Gate(item, "judged", 1, "")
		agrees := 0.0
		if d.judgement != nil && d.judgement.Tier == d.modelGrade {
			agrees = 1
		}
		tier := ""
		if d.judgement != nil {
			tier = d.judgement.Tier
		}
		run.Info(item, "judge_agrees", agrees, fmt.Sprintf("graded %s, judged %s", d.modelGrade, tier))
	}
	run.Gate("suite", "judge_agreement", share*100, fmt.Sprintf("%d of %d, the bar is %.0f", agreed, judgedCount, judgeBar*100))
	run.Finish(acc.Report(), took, verdict)
	p("# PR-14B quality judge lane\n\n")
	p("Run date: %s. Card snapshot: %s. Source: %s.\n\n", time.Now().UTC().Format("2006-01-02"), idx.AsOf.Format("2006-01-02"), source)
	p("Verdict: %s. The judge agreed with the model's grade on %d of %d decks (%.0f percent, the bar is %.0f), %d off by one rung, %d judge errors.\n\n",
		verdict, agreed, judgedCount, share*100, judgeBar*100, near, errs)
	p("The grade column is the stored model's read of the deck now, and the document column the grade the deck gate printed. The bar reads the model.\n\n")
	p("## Summary\n\n| Measure | Value |\n|---|---|\n")
	p("| Decks | %d |\n| Judged | %d |\n| Agreed | %d |\n| Off by one rung | %d |\n| Judge errors | %d |\n", len(decks), judgedCount, agreed, near, errs)
	rep := acc.Report()
	p("| Calls | %d |\n| Cost | %s |\n| Time | %.0f seconds |\n\n", rep.Calls, gatekit.CostWord(rep), took.Seconds())
	run.Markdown(w)
	p("\n")
	p("## Decks\n\n| # | Deck | Format | Grade | Document | Judge | Agree |\n|---|---|---|---|---|---|---|\n")
	for _, d := range decks {
		judge, agree := "error", ""
		if d.judgement != nil {
			judge = d.judgement.Tier
			agree = "no"
			if judge == d.modelGrade {
				agree = "yes"
			}
		}
		p("| %d | %s | %s | %s | %s | %s | %s |\n", d.id, d.title, generate.FormatWord(d.format), d.modelGrade, d.grade, judge, agree)
	}
	p("\n## The judge's reasons\n\n")
	for _, d := range decks {
		if d.err != nil {
			p("- %d. %s: JUDGE ERROR: %v\n", d.id, d.title, d.err)
			continue
		}
		p("- %d. %s: graded %s, judged %s. %s\n", d.id, d.title, d.modelGrade, d.judgement.Tier, d.judgement.Why)
	}
	p("\nVerdict: %s.\n", verdict)
	return pass
}
