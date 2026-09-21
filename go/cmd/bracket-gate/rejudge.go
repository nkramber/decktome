package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/rules"
)

// The re-judge mode reads the decks of a gate document and runs the
// judge lane over them. Run 1 of 2026-09-02 built 15 decks and lost
// every judge call to a schema fault, and a build costs more than a
// judge call, so the decks are read back and not built again.

var (
	deckHeader = regexp.MustCompile(`^### (\d+)\. Bracket (\d), (.+)$`)
	cardLine   = regexp.MustCompile(`^- (\d+) (.+)$`)
)

// readDecks parses the deck sections of a gate document: the header
// with the id, the bracket, the commander, and the theme, and the card
// list under "Cards:". A name the index does not know is an error, so
// the judge never reads a deck with a hole.
func readDecks(r io.Reader, idx *cards.Index) ([]result, error) {
	var out []result
	var cur *result
	inCards := false
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		line := sc.Text()
		if m := deckHeader.FindStringSubmatch(line); m != nil {
			id, _ := strconv.Atoi(m[1])
			br, _ := strconv.Atoi(m[2])
			c, theme, ok := splitHeader(m[3], idx)
			if !ok {
				return nil, fmt.Errorf("deck %d: no commander in %q", id, m[3])
			}
			out = append(out, result{
				prompt: prompt{ID: id, Bracket: int32(br), Commander: c.GetName(), Theme: theme},
				deck: &mtgv1.Deck{
					Format:             &mtgv1.Format{Id: mtgv1.FormatId_FORMAT_ID_COMMANDER},
					Power:              gatekit.PowerLevel(int32(br), ""),
					CommanderOracleIds: []string{c.GetOracleId()},
				},
			})
			cur = &out[len(out)-1]
			inCards = false
			continue
		}
		if cur == nil {
			continue
		}
		if line == "Cards:" {
			inCards = true
			continue
		}
		if !inCards {
			continue
		}
		m := cardLine.FindStringSubmatch(line)
		if m == nil {
			if strings.TrimSpace(line) != "" {
				inCards = false
			}
			continue
		}
		n, _ := strconv.Atoi(m[1])
		c, ok := idx.ByName(m[2])
		if !ok {
			return nil, fmt.Errorf("deck %d: no card named %q", cur.prompt.ID, m[2])
		}
		cur.deck.Cards = append(cur.deck.Cards, &mtgv1.DeckCard{OracleId: c.GetOracleId(), Name: c.GetName(), Count: int32(n)})
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("the document holds no deck section")
	}
	for _, r := range out {
		if len(r.deck.Cards) == 0 {
			return nil, fmt.Errorf("deck %d holds no cards", r.prompt.ID)
		}
	}
	return out, nil
}

// readContent profiles each stored deck, so the judge reads the combos
// of Commander Spellbook as the build path gives them (F-126, D-790). The
// endpoint is free, and a deck it fails keeps an unchecked profile. Each
// deck also reads its rules floor (F-162, D-793).
func readContent(ctx context.Context, rs []result, idx *cards.Index, log *slog.Logger) error {
	rcfg, err := rules.Load()
	if err != nil {
		return err
	}
	prof, err := gatekit.Profiler(idx, rcfg, log)
	if err != nil {
		return err
	}
	for i := range rs {
		rs[i].deck.Profile, _ = prof.Read(ctx, rs[i].deck, idx)
		rs[i].floor, rs[i].floorErr = prof.Floor(ctx, rs[i].deck, idx)
	}
	return nil
}

// splitHeader reads "commander, theme" where the commander can hold a
// comma, as "Gishath, Sun's Avatar" does. The commander is the longest
// prefix the index knows, and the rest is the theme.
func splitHeader(rest string, idx *cards.Index) (*mtgv1.Card, string, bool) {
	parts := strings.Split(rest, ", ")
	for i := len(parts) - 1; i >= 1; i-- {
		if c, ok := idx.ByName(strings.Join(parts[:i], ", ")); ok {
			return c, strings.Join(parts[i:], ", "), true
		}
	}
	return nil, "", false
}

func runRejudge(path, runOut string) error {
	if err := gatekit.SpendGuard("BRACKET_GATE"); err != nil {
		return err
	}
	if err := gatekit.RefuseExisting(runOut); err != nil {
		return err
	}
	run := evalrun.New("bracket-judge", evalrun.RunID(runOut))
	run.Header.Prompts["generate"] = generate.PromptVersion
	run.Header.Prompts["bracket_judge"] = generate.BracketJudgeVersion
	run.Header.Versions["source"] = filepath.Base(path)
	quiet := gatekit.Quiet()
	idx, err := gatekit.LoadSnapshot(context.Background(), quiet)
	if err != nil {
		return err
	}
	run.SetSnapshot(idx.AsOf)
	f, err := os.Open(path) // #nosec G304 -- the operator names the file.
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	results, err := readDecks(f, idx)
	if err != nil {
		return fmt.Errorf("rejudge %s: %w", path, err)
	}
	if err := readContent(context.Background(), results, idx, quiet); err != nil {
		return err
	}
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
	for i := range results {
		r := &results[i]
		r.judged, r.judgeErr = generate.JudgeBracket(context.Background(), client, r.deck, idx, acc)
		raiseToFloor(r)
		word := "error"
		if r.judged != nil {
			word = fmt.Sprintf("judged %d", r.judged.Bracket)
		}
		fmt.Fprintf(os.Stderr, "  %2d. bracket %d %-28s %s\n", r.prompt.ID, r.prompt.Bracket, r.prompt.Commander, word)
	}
	pass := reportJudge(os.Stdout, path, results, acc, idx, time.Since(start), run)
	if err := evalrun.WriteFile(runOut, run); err != nil {
		return err
	}
	if !pass {
		return errGateFailed
	}
	return nil
}

// raiseToFloor lifts the judge's bracket to the lowest bracket the rules
// allow the deck. The rules own the low end, and the judge reads above
// it (D-699). The judge named a listed combo and still read three precons
// below the bracket its rules allow (F-162, D-795).
func raiseToFloor(r *result) {
	if r.judged != nil && r.floorErr == nil && r.floor > r.judged.Bracket {
		r.raisedFrom = r.judged.Bracket
		r.judged.Bracket = r.floor
	}
}

// judgeWords writes the judge's bracket, and the bracket it named when
// the floor raised it.
func judgeWords(r result) string {
	if r.raisedFrom > 0 {
		return fmt.Sprintf("%d, raised from %d to the rules floor", r.judged.Bracket, r.raisedFrom)
	}
	return strconv.Itoa(int(r.judged.Bracket))
}

// isPrecon reports whether a deck is a precon of the judge calibration.
// A precon anchors no bracket, because the Commander Format Panel tied
// none to a bracket after 2025-10-21, so the lane scores it against its
// rules floor alone (F-162, D-793).
func isPrecon(p prompt) bool { return strings.HasPrefix(p.Theme, "precon ") }

// heldFloor reports whether the judge read a precon at or above the
// lowest bracket its rules allow.
func heldFloor(r result) bool {
	return r.floorErr == nil && r.judged != nil && r.judged.Bracket >= r.floor
}

// reportJudge writes the judge lane document and returns its verdict:
// every deck judged, the agreement at the bar on each anchored deck, and
// each precon at or above its rules floor.
func reportJudge(w io.Writer, source string, rs []result, acc *llm.Accumulator, idx *cards.Index, took time.Duration, run *evalrun.Run) bool {
	judged, anchored, agreed, errs, precons, held := 0, 0, 0, 0, 0, 0
	for _, r := range rs {
		switch {
		case r.judgeErr != nil:
			errs++
		case r.judged != nil:
			judged++
			if isPrecon(r.prompt) {
				precons++
				if heldFloor(r) {
					held++
				}
				continue
			}
			anchored++
			if r.judged.Bracket == r.prompt.Bracket {
				agreed++
			}
		}
	}
	agreement := 0
	if anchored > 0 {
		agreement = agreed * 100 / anchored
	}
	pass := len(rs) > 0 && errs == 0 && judged == len(rs) && held == precons &&
		(anchored == 0 || agreement >= JudgeAgreementPercent)
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	for _, r := range rs {
		item := fmt.Sprintf("%d", r.prompt.ID)
		switch {
		case r.judgeErr != nil:
			run.Gate(item, "judged", 0, r.judgeErr.Error())
		case r.judged != nil && isPrecon(r.prompt):
			run.Gate(item, "judged", 1, "")
			if r.floorErr != nil {
				run.Gate(item, "precon_floor", 0, "no floor: "+r.floorErr.Error())
				break
			}
			ok := 0.0
			if heldFloor(r) {
				ok = 1
			}
			run.Gate(item, "precon_floor", ok, fmt.Sprintf("floor %d, judged %d", r.floor, r.judged.Bracket))
		case r.judged != nil:
			run.Gate(item, "judged", 1, "")
			agrees := 0.0
			if r.judged.Bracket == r.prompt.Bracket {
				agrees = 1
			}
			run.Info(item, "judge_agrees", agrees, fmt.Sprintf("built for %d, judged %d", r.prompt.Bracket, r.judged.Bracket))
		default:
			run.Gate(item, "judged", 0, "no verdict")
		}
	}
	run.Gate("suite", "judge_agreement", float64(agreement), fmt.Sprintf("%d of %d, the bar is %d", agreed, anchored, JudgeAgreementPercent))
	run.Finish(acc.Report(), took, verdict)
	_, _ = fmt.Fprintf(w, "# PR-14A bracket gate, judge lane\n\n")
	_, _ = fmt.Fprintf(w, "Run date: %s. Card snapshot: %s. Decks read from `%s`.\n\n", time.Now().UTC().Format("2006-01-02"), idx.AsOf.Format("2006-01-02"), source)
	_, _ = fmt.Fprintf(w, "Verdict: %s. The judge agreed with the bracket on %d of %d decks (%d percent, the bar is %d), with %d judge errors. This document reads the judge bar alone: the block, band, and content bars are in the source document.\n\n",
		verdict, agreed, anchored, agreement, JudgeAgreementPercent, errs)
	if precons > 0 {
		_, _ = fmt.Fprintf(w, "The judge read %d of %d precons at or above the lowest bracket their rules allow. A precon anchors no bracket, so the agreement leaves it out (F-162, D-793).\n\n", held, precons)
	}
	rep := acc.Report()
	_, _ = fmt.Fprintf(w, "Calls %d. Cost %s. Time %.0f seconds.\n\n", rep.Calls, gatekit.CostWord(rep), took.Seconds())
	run.Markdown(w)
	_, _ = fmt.Fprintf(w, "\n")
	_, _ = fmt.Fprintf(w, "| # | Built for | Judged | Agrees | Commander |\n|---|---|---|---|---|\n")
	for _, r := range rs {
		judgedWord, agrees := "error", ""
		if r.judged != nil {
			judgedWord = judgeWords(r)
			agrees = "no"
			if r.judged.Bracket == r.prompt.Bracket {
				agrees = "yes"
			}
		}
		built := strconv.Itoa(int(r.prompt.Bracket))
		if isPrecon(r.prompt) {
			built, agrees = "no floor", "no"
			if r.floorErr == nil {
				built = fmt.Sprintf("floor %d", r.floor)
			}
			if heldFloor(r) {
				agrees = "at or above"
			}
		}
		_, _ = fmt.Fprintf(w, "| %d | %s | %s | %s | %s |\n", r.prompt.ID, built, judgedWord, agrees, r.prompt.Commander)
	}
	_, _ = fmt.Fprintf(w, "\n## Why\n\n")
	for _, r := range rs {
		head := fmt.Sprintf("Bracket %d", r.prompt.Bracket)
		if isPrecon(r.prompt) {
			head = "No floor"
			if r.floorErr == nil {
				head = fmt.Sprintf("Floor %d", r.floor)
			}
		}
		_, _ = fmt.Fprintf(w, "### %d. %s, %s, %s\n\n", r.prompt.ID, head, r.prompt.Commander, r.prompt.Theme)
		if isPrecon(r.prompt) && r.floorErr != nil {
			_, _ = fmt.Fprintf(w, "Floor: %s\n\n", r.floorErr.Error())
		}
		switch {
		case r.judgeErr != nil:
			_, _ = fmt.Fprintf(w, "Judge: error, %s\n\n", r.judgeErr.Error())
		case r.judged != nil:
			_, _ = fmt.Fprintf(w, "Judge: bracket %s. %s\n\n", judgeWords(r), strings.TrimSpace(r.judged.Why))
		}
		if c := r.deck.GetProfile().GetContent(); c != nil {
			_, _ = fmt.Fprintf(w, "Combos the judge read: %s.\n\n", comboNames(c))
		}
	}
	return pass
}

// comboNames lists the combos of a content check on one line, or says
// that the check found none or did not run.
func comboNames(c *mtgv1.ContentCheck) string {
	if !c.GetChecked() {
		return "none, the check did not run: " + c.GetError()
	}
	var out []string
	for _, combo := range c.GetCombos() {
		out = append(out, strings.Join(combo.GetCards(), " + "))
	}
	if len(out) == 0 {
		return "none, Commander Spellbook finds no combo"
	}
	return strings.Join(out, "; ")
}
