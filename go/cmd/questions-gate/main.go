// Command questions-gate writes the PR-7 gate document. It runs the gate
// conversations against the real providers and reports M-4: how many
// questions came from the catalog, how many the model invented, the gap
// score of each, and whether an invented question filled its slot.
//
// The file also holds probe conversations. A probe explores catalog
// coverage rather than the gate, so it holds a shape the gate set never
// takes: a user who changes their mind, contradicts themselves, asks a
// question back, or wants something the app can not serve. Counting a
// probe would move the bar, so the verdict reads the gate set alone. The
// questions a probe raises still reach the M-5 sheet (D-96).
//
// TestConversations proves the other half of the gate offline: a complete
// slot set in at most four turns, with no repeated question. Only the
// model invents a question, so this half needs a live run.
//
// CAUTION: this command spends money. It makes about three model calls
// per turn, on the cost tier. QUESTIONS_GATE=1 is required, so it can
// not run by accident.
//
// Usage:
//
//	set -a && . ./.env && set +a
//	QUESTIONS_GATE=1 CARDS_SNAPSHOT_DIR=.local/gcs/mtg-local-cards/scryfall \
//	  go run ./cmd/questions-gate -collection internal/collections/testdata/manabox_collection.csv \
//	  > ../docs/reference/pr7-question-gate.md
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/candidates"
	"github.com/nkramber/mtg-deck-builder/go/internal/cards"
	"github.com/nkramber/mtg-deck-builder/go/internal/gatekit"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
	"github.com/nkramber/mtg-deck-builder/go/internal/questions"
)

//go:embed conversations.json
var conversationsJSON []byte

// CatalogOnlyBar is the roadmap gate: at least 25 of the counted gate
// conversations use catalog questions only. A gate conversation that
// starts after a build is not counted: its build slots are closed before
// the first message, so it can not invent one (A-9).
const CatalogOnlyBar = 25

type gateFile struct {
	VerifiedAt    string         `json:"verified_at"`
	Note          string         `json:"note"`
	Conversations []conversation `json:"conversations"`
}

type conversation struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Collection bool     `json:"collection"`
	Messages   []string `json:"messages"`
	// Note says why the conversation exists. The struct names it so the
	// file can not carry a field nothing reads (T-19).
	Note string `json:"note,omitempty"`
	// Probe marks a conversation that explores catalog coverage rather
	// than one the gate scores. A probe holds a shape the gate set never
	// takes: a user who changes their mind, contradicts themselves, or
	// wants something the app can not serve. Counting probes would move
	// the bar, so the verdict reads the gate set alone. Their questions
	// still reach the M-5 sheet (D-96).
	Probe bool `json:"probe"`
	// HasDeck starts the conversation after a finished build. agentsvc
	// sets the fact from the session's stored decks, and this harness runs
	// the questions agent directly, so without the flag the fact can never
	// be true. A deck exists only when the build slots were filled, so the
	// flag fills them too (D-239, OQ-38).
	HasDeck bool `json:"has_deck"`
}

// builtSlots are the slots a finished build must have settled. A stored
// deck is proof that each one was answered (D-239). The document reports
// the same slots for every conversation.
var builtSlots = []string{"format", "theme", "colors", "commander", "power", "pool_rule", "budget"}

// gateCount counts the conversations the verdict reads: every one that
// is not a probe. The size test counts the same way (T-20).
func gateCount(convs []conversation) int {
	n := 0
	for _, c := range convs {
		if !c.Probe {
			n++
		}
	}
	return n
}

// asked is one question as the run produced it.
type asked struct {
	Turn     int
	Slot     string
	Row      string
	Source   string
	Fit      float64
	Text     string
	Options  []string
	Filled   bool
	Invented bool
	// Catalog is the row text the model replaced. The M-5 rubric asks
	// whether the catalog was enough and whether the invention is better,
	// and neither can be scored without both texts side by side (D-66).
	Catalog string
	// NearCopy marks a turn where the model offered a replacement and the
	// agent refused it as a reword (D-88).
	NearCopy bool
	// Refused is the text of that replacement, for the M-5 sheet to judge.
	Refused string
	// Resolved is the catalog row before the ask role phrased it. The
	// reword guard compared the replacement against this text (D-116).
	Resolved string
}

type result struct {
	conversation
	Questions []asked
	Coverage  questions.Coverage
	// Turns counts the turns that ran. A session that reaches a full slot
	// set stops early, and the document must not print the messages that
	// never went out.
	Turns int
	// Ready and Slots record where the conversation ended. A session that
	// stops after one question is either complete or premature, and only
	// the slot state tells the two apart.
	Ready bool
	Slots map[string]string
	// Unanswered are the slots the deck needs that no answer filled.
	Unanswered []string
	// Premature marks a session that called itself complete with a slot
	// still unanswered. That is a gate failure: the deck would carry a
	// value nobody chose.
	Premature bool
	// Findings are the deterministic defects the linter found in the
	// questions this conversation sent (D-115).
	Findings []questions.Finding
	Err      error
}

func main() {
	collectionPath := flag.String("collection", "", "ManaBox CSV for the owned-mode conversations")
	limit := flag.Int("n", 0, "run only the first n conversations (0 runs all)")
	// only names the conversation ids to run, for a cheap check of one
	// row. A full run costs money, and a wasted one costs it twice.
	only := flag.String("only", "", "run only these conversation ids, comma separated")
	flag.Parse()
	if err := run(*collectionPath, *limit, *only, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(collectionPath string, limit int, only string, w io.Writer) error {
	if err := gatekit.SpendGuard("QUESTIONS_GATE"); err != nil {
		return err
	}
	var file gateFile
	if err := json.Unmarshal(conversationsJSON, &file); err != nil {
		return fmt.Errorf("conversations.json: %w", err)
	}
	if n := gateCount(file.Conversations); n < questions.MinGateSize {
		return fmt.Errorf("conversations.json holds %d gate conversations, the gate needs %d", n, questions.MinGateSize)
	}
	cat, err := questions.Load()
	if err != nil {
		return err
	}
	quiet := gatekit.Quiet()
	client, err := llm.NewFromEnv(gatekit.Env, quiet)
	if err != nil {
		return err
	}
	prices, err := llm.LoadPrices()
	if err != nil {
		return err
	}
	idx, owned, ownedNote, err := loadIndex(collectionPath)
	if err != nil {
		return err
	}
	builder, err := candidates.New()
	if err != nil {
		return err
	}

	list := file.Conversations
	ids, err := gatekit.ParseIDs(only)
	if err != nil {
		return fmt.Errorf("questions-gate: %w", err)
	}
	if ids != nil {
		want := map[int]bool{}
		for _, id := range ids {
			want[id] = true
		}
		var kept []conversation
		for _, c := range list {
			if want[c.ID] {
				kept = append(kept, c)
			}
		}
		if len(kept) == 0 {
			return fmt.Errorf("-only %q matches no conversation", only)
		}
		list = kept
	}
	if limit > 0 && limit < len(list) {
		list = list[:limit]
	}
	acc := llm.NewAccumulator(prices)
	started := time.Now()
	results := make([]result, 0, len(list))
	var cov coverages
	for _, conv := range list {
		res := runOne(cat, client, idx, builder, owned, conv, acc)
		results = append(results, res)
		cov.add(res)
		fmt.Fprintf(os.Stderr, "%2d/%d %-40s catalog=%d invented=%d%s\n",
			conv.ID, len(list), conv.Name, res.Coverage.Catalog, res.Coverage.Invented, res.kind())
	}
	return write(w, file, results, cov, acc.Report(), client.Config(), ownedNote, time.Since(started))
}

// coverages splits the M-4 counts three ways: the counted gate
// conversations, the ones that start after a build, and the probes.
type coverages struct {
	total, afterBuild, probes questions.Coverage
}

func (c *coverages) add(res result) {
	switch {
	case res.Probe:
		c.probes.Add(asksOf(res))
	case res.HasDeck:
		c.afterBuild.Add(asksOf(res))
	default:
		c.total.Add(asksOf(res))
	}
}

// kind is the heading suffix that tells the reader, and tune.ReadRun,
// which set a conversation belongs to.
func (r result) kind() string {
	switch {
	case r.Probe:
		return " (probe)"
	case r.HasDeck:
		return " (after a build)"
	}
	return ""
}

// runOne plays one conversation. A failed turn ends that conversation and
// the document names the error.
func runOne(cat *questions.Catalog, client *llm.Client, idx *cards.Index, builder *candidates.Builder,
	owned map[string]int32, conv conversation, acc *llm.Accumulator) result {
	res := result{conversation: conv}
	quiet := gatekit.Quiet()
	st := questions.NewState(conv.Collection)
	st.SessionID = fmt.Sprintf("gate-%02d", conv.ID)
	// agentsvc reads these from the session (D-239). A stored deck means a
	// build finished, so the slots it needed are settled.
	st.Ctx.AfterBuild = conv.HasDeck
	if conv.HasDeck {
		for _, k := range builtSlots {
			// Close, not Ctx.Filled: a slot the planner calls filled and
			// the record calls never asked reads as a premature session
			// (D-252).
			st.Close(k)
		}
		st.Ctx.CommanderSet = true
	}
	for i, msg := range conv.Messages {
		// One agent per turn, with hints built from the slots as they
		// stand. agentsvc does the same, and a hint source that outlives
		// a turn keeps an answer from before the user named the colors.
		opts := []questions.AgentOption{questions.WithLogger(quiet)}
		if idx != nil {
			hints := &questions.CandidateHints{
				Index: idx, Builder: builder, Log: quiet,
				Format: st.Slots.GetFormat().GetId(),
				Colors: st.Slots.GetColors(),
				Pool:   st.Slots.GetPoolRule(),
			}
			if hints.Format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
				hints.Format = mtgv1.FormatId_FORMAT_ID_COMMANDER
			}
			if conv.Collection {
				hints.Owned = owned
			}
			opts = append(opts, questions.WithHints(hints))
		}
		agent, err := questions.NewAgent(cat, client, opts...)
		if err != nil {
			res.Err = err
			break
		}
		res.Turns = i + 1
		turn, err := agent.Turn(context.Background(), st, msg, acc)
		if err != nil {
			res.Err = fmt.Errorf("turn %d: %w", i+1, err)
			break
		}
		for _, q := range turn.Questions {
			res.Questions = append(res.Questions, asked{
				Turn: i + 1, Slot: q.GetSlot(), Source: source(q.GetInvented()),
				Fit: q.GetGapScore(), Text: q.GetText(), Options: q.GetOptions(),
				Invented: q.GetInvented(), Catalog: q.GetCatalogText(),
			})
		}
		if turn.Ready {
			break
		}
	}
	// The row id and the filled flag come from the M-4 records. The two
	// lists are written in step, one record per question sent, and a
	// length that differs would map a row onto the wrong question.
	if res.Err == nil && len(st.Asks) != len(res.Questions) {
		res.Err = fmt.Errorf("%d M-4 records for %d questions sent", len(st.Asks), len(res.Questions))
	}
	for i, rec := range st.Asks {
		if i < len(res.Questions) {
			res.Questions[i].Row = rec.RowID
			res.Questions[i].Filled = rec.Filled
			res.Questions[i].Catalog = rec.CatalogText
			res.Questions[i].NearCopy = rec.NearCopy
			res.Questions[i].Refused = rec.RefusedText
			res.Questions[i].Resolved = rec.ResolvedText
		}
	}
	// The linter reads the questions that went out against the messages
	// that came before them. It costs no model call (D-115).
	lint := make([]questions.LintQuestion, 0, len(res.Questions))
	for _, q := range res.Questions {
		lint = append(lint, questions.LintQuestion{Turn: q.Turn, RowID: q.Row, Slot: q.Slot, Text: q.Text})
	}
	res.Findings = questions.LintConversation(conv.Messages[:res.Turns], lint)
	res.Coverage = st.Metrics()
	res.Ready = st.Ready(cat)
	res.Slots = map[string]string{}
	for _, key := range builtSlots {
		res.Slots[key] = slotState(st, key)
	}
	for _, key := range required(st, conv.Collection) {
		if state := res.Slots[key]; state != "filled" && state != "skipped" {
			res.Unanswered = append(res.Unanswered, key+" ("+state+")")
		}
	}
	// A session that ran out of messages is not a defect: the script
	// ended. A session that called itself complete with a slot open is.
	res.Premature = res.Ready && len(res.Unanswered) > 0
	return res
}

// required lists the slots the deck can not be built without. The
// commander applies to Commander alone, and the pool rule applies only
// when the user has a library (D-37).
func required(st *questions.State, hasCollection bool) []string {
	req := []string{"format", "theme", "colors", "power"}
	if st.Slots.GetFormat().GetId() == mtgv1.FormatId_FORMAT_ID_COMMANDER {
		req = append(req, "commander")
	}
	if hasCollection {
		req = append(req, "pool_rule")
	}
	return req
}

// slotState reads one slot as a word a reader can scan.
func slotState(st *questions.State, key string) string {
	switch st.Slots.GetSlotStates()[key] {
	case mtgv1.SlotState_SLOT_STATE_FILLED:
		return "filled"
	case mtgv1.SlotState_SLOT_STATE_SKIPPED:
		return "skipped"
	case mtgv1.SlotState_SLOT_STATE_ASKED:
		return "asked, no answer"
	default:
		return "never asked"
	}
}

func asksOf(r result) []questions.Ask {
	out := make([]questions.Ask, 0, len(r.Questions))
	for _, q := range r.Questions {
		out = append(out, questions.Ask{RowID: q.Row, Slot: q.Slot, Fit: q.Fit,
			Invented: q.Invented, Filled: q.Filled, NearCopy: q.NearCopy})
	}
	return out
}

func source(invented bool) string {
	if invented {
		return "INVENTED"
	}
	return "catalog"
}

// loadIndex reads the local snapshot and the collection export. Both are
// optional: without them every clause that needs a value is dropped.
func loadIndex(collectionPath string) (*cards.Index, map[string]int32, string, error) {
	dir := os.Getenv("CARDS_SNAPSHOT_DIR")
	if dir == "" {
		return nil, nil, "no snapshot: the hint clauses were dropped", nil
	}
	idx, err := gatekit.LoadSnapshot(context.Background(), gatekit.Quiet())
	if err != nil {
		return nil, nil, "", err
	}
	if collectionPath == "" {
		return idx, nil, "snapshot loaded, no collection", nil
	}
	owned, note, err := gatekit.LoadOwned(collectionPath, idx)
	if err != nil {
		return nil, nil, "", err
	}
	return idx, owned, "snapshot loaded, " + note, nil
}

func write(w io.Writer, file gateFile, results []result, cov coverages,
	report llm.Report, cfg *llm.Config, ownedNote string, elapsed time.Duration) error {
	total, probes := cov.total, cov.probes
	var premature []string
	gate, counted, afterBuild := 0, 0, 0
	for _, r := range results {
		if r.Premature {
			premature = append(premature, r.Name)
		}
		switch {
		case r.Probe:
		case r.HasDeck:
			gate++
			afterBuild++
		default:
			gate++
			counted++
		}
	}
	// The linter is the third bar. A question that names a format the user
	// gave, presumes a table, or states a fact about the game fails the
	// gate, whatever the catalog-only count says (D-115).
	findings := 0
	byRule := map[string]int{}
	for _, r := range results {
		findings += len(r.Findings)
		for _, f := range r.Findings {
			byRule[f.Rule]++
		}
	}
	pass := total.CatalogOnly >= CatalogOnlyBar && gate >= questions.MinGateSize &&
		len(premature) == 0 && findings == 0
	verdict := "FAIL"
	if pass {
		verdict = "PASS"
	}
	for _, r := range results {
		if r.Err != nil {
			verdict, pass = "FAIL", false
		}
	}

	_, _ = fmt.Fprintf(w, "# PR-7 question gate\n\n")
	_, _ = fmt.Fprintf(w, "Run date: %s. Conversations: %s.\n\n", time.Now().UTC().Format("2006-01-02"), file.VerifiedAt)
	_, _ = fmt.Fprintf(w, "Verdict: %s. %d of %d counted gate conversations used catalog questions only. The bar is %d. The set holds %d gate conversations, and %d of them start after a build and are not counted (A-9).\n\n",
		verdict, total.CatalogOnly, counted, CatalogOnlyBar, gate, afterBuild)
	if cov.afterBuild.Sessions > 0 {
		_, _ = fmt.Fprintf(w, "%d conversations after a build ran beside the count. They asked %d questions, and the model offered %d replacements.\n\n",
			cov.afterBuild.Sessions, cov.afterBuild.Asked, cov.afterBuild.Invented+cov.afterBuild.NearCopies)
	}
	if probes.Sessions > 0 {
		_, _ = fmt.Fprintf(w, "%d probe conversations ran beside the gate. They asked %d questions, and the model offered %d replacements. A probe explores catalog coverage and does not move the verdict (D-96).\n\n",
			probes.Sessions, probes.Asked, probes.Invented+probes.NearCopies)
	}
	if len(premature) > 0 {
		_, _ = fmt.Fprintf(w, "%d conversations called themselves complete with a slot still unanswered, which fails the gate: %s.\n\n",
			len(premature), strings.Join(premature, "; "))
	} else {
		_, _ = fmt.Fprintf(w, "No conversation called itself complete with a slot unanswered.\n\n")
	}
	if findings > 0 {
		_, _ = fmt.Fprintf(w, "The linter found %d defective questions, which fails the gate (D-115).\n\n", findings)
		_, _ = fmt.Fprintf(w, "| Rule | Count |\n|---|---|\n")
		rules := make([]string, 0, len(byRule))
		for rule := range byRule {
			rules = append(rules, rule)
		}
		sort.Strings(rules)
		for _, rule := range rules {
			_, _ = fmt.Fprintf(w, "| `%s` | %d |\n", rule, byRule[rule])
		}
		_, _ = fmt.Fprintf(w, "\n")
		for _, r := range results {
			for _, f := range r.Findings {
				_, _ = fmt.Fprintf(w, "- %s: %s\n", r.Name, f)
			}
		}
		_, _ = fmt.Fprintf(w, "\n")
	} else {
		_, _ = fmt.Fprintf(w, "The linter found no defective question (D-115).\n\n")
	}

	_, _ = fmt.Fprintf(w, "## M-4 report\n\n")
	_, _ = fmt.Fprintf(w, "| Measure | Value |\n|---|---|\n")
	_, _ = fmt.Fprintf(w, "| Conversations | %d |\n", total.Sessions)
	_, _ = fmt.Fprintf(w, "| Questions asked | %d |\n", total.Asked)
	_, _ = fmt.Fprintf(w, "| From the catalog | %d |\n", total.Catalog)
	_, _ = fmt.Fprintf(w, "| Invented by the model | %d |\n", total.Invented)
	_, _ = fmt.Fprintf(w, "| Replacements refused as rewords | %d |\n", total.NearCopies)
	_, _ = fmt.Fprintf(w, "| Catalog questions that closed a slot | %d |\n", total.CatalogFilled)
	_, _ = fmt.Fprintf(w, "| Invented questions that closed a slot | %d |\n", total.InventedFilled)
	_, _ = fmt.Fprintf(w, "| Catalog-only conversations | %d |\n", total.CatalogOnly)
	_, _ = fmt.Fprintf(w, "| Median gap score | %.2f |\n", total.MedianFit())
	_, _ = fmt.Fprintf(w, "| Fit threshold | %.2f |\n\n", questions.DefaultFitThreshold)

	if len(total.InventedByRow) > 0 {
		_, _ = fmt.Fprintf(w, "### Invented questions by the row they replaced\n\n")
		_, _ = fmt.Fprintf(w, "A row that repeats here is a catalog change candidate (D-25, PR-15).\n\n")
		_, _ = fmt.Fprintf(w, "| Row | Count |\n|---|---|\n")
		rows := make([]string, 0, len(total.InventedByRow))
		for row := range total.InventedByRow {
			rows = append(rows, row)
		}
		sort.Slice(rows, func(i, j int) bool {
			if total.InventedByRow[rows[i]] != total.InventedByRow[rows[j]] {
				return total.InventedByRow[rows[i]] > total.InventedByRow[rows[j]]
			}
			return rows[i] < rows[j]
		})
		for _, row := range rows {
			_, _ = fmt.Fprintf(w, "| `%s` | %d |\n", row, total.InventedByRow[row])
		}
		_, _ = fmt.Fprintln(w)
	}

	_, _ = fmt.Fprintf(w, "## Run\n\n")
	_, _ = fmt.Fprintf(w, "- Roles: classify on `%s`, ask on `%s`.\n",
		cfg.Roles[llm.RoleClassify].Model, cfg.Roles[llm.RoleAsk].Model)
	_, _ = fmt.Fprintf(w, "- Cards: %s.\n", ownedNote)
	_, _ = fmt.Fprintf(w, "- Calls: %d. Time: %.1f seconds.\n", report.Calls, elapsed.Seconds())
	if report.Tokens != nil {
		_, _ = fmt.Fprintf(w, "- Tokens: %d input (%d cached), %d output.\n",
			report.Tokens.InputTokens, report.Tokens.CachedInputTokens, report.Tokens.OutputTokens)
	}
	if report.CostUSD != nil {
		_, _ = fmt.Fprintf(w, "- Cost: $%.4f.\n", *report.CostUSD)
	} else {
		_, _ = fmt.Fprintf(w, "- Cost: unpriced. A model in this run has no price row.\n")
	}
	_, _ = fmt.Fprintln(w)

	_, _ = fmt.Fprintf(w, "## Conversations\n\n")
	for _, r := range results {
		_, _ = fmt.Fprintf(w, "### %d. %s%s\n\n", r.ID, r.Name, r.kind())
		if r.Err != nil {
			_, _ = fmt.Fprintf(w, "**Failed: %v**\n\n", r.Err)
		}
		_, _ = fmt.Fprintf(w, "Collection: %v. Catalog: %d. Invented: %d.\n\n",
			r.Collection, r.Coverage.Catalog, r.Coverage.Invented)
		if r.Turns < len(r.Messages) {
			_, _ = fmt.Fprintf(w, "The slots were full after turn %d. The last %d messages never went out.\n\n",
				r.Turns, len(r.Messages)-r.Turns)
		}
		if len(r.Unanswered) > 0 {
			label := "Slots the deck needs and nobody answered"
			if r.Premature {
				label = "**PREMATURE.** It called itself complete with these slots unanswered"
			}
			_, _ = fmt.Fprintf(w, "%s: %s.\n\n", label, strings.Join(r.Unanswered, ", "))
		}
		if r.Turns > 0 && len(r.Questions) == 0 {
			_, _ = fmt.Fprintf(w, "**This conversation asked nothing.** One message filled every slot, so it tests no question.\n\n")
		}
		for i, msg := range r.Messages[:r.Turns] {
			_, _ = fmt.Fprintf(w, "**Turn %d, the user:** %s\n\n", i+1, msg)
			for _, q := range r.Questions {
				if q.Turn != i+1 {
					continue
				}
				_, _ = fmt.Fprintf(w, "- [%s slot=%s row=%s fit=%.2f filled=%v] %s\n",
					q.Source, q.Slot, q.Row, q.Fit, q.Filled, q.Text)
				if len(q.Options) > 0 {
					_, _ = fmt.Fprintf(w, "  - Options: %s\n", strings.Join(q.Options, " / "))
				}
				if q.Invented && q.Catalog != "" {
					_, _ = fmt.Fprintf(w, "  - It replaced: %s\n", q.Catalog)
				}
				if q.NearCopy {
					_, _ = fmt.Fprintf(w, "  - Refused as a reword (D-88): %s\n", q.Refused)
					if q.Resolved != "" {
						_, _ = fmt.Fprintf(w, "  - The guard compared against: %s\n", q.Resolved)
					}
				}
			}
			_, _ = fmt.Fprintln(w)
		}
	}
	if !pass {
		return fmt.Errorf("gate failed: %d of %d counted catalog-only (bar %d), %d premature, %d lint findings",
			total.CatalogOnly, counted, CatalogOnlyBar, len(premature), findings)
	}
	return nil
}
