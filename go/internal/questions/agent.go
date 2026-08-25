package questions

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/mtg-deck-builder/go/gen/mtg/v1"
	"github.com/nkramber/mtg-deck-builder/go/internal/llm"
)

// DefaultFitThreshold is the provisional gap-score threshold (D-27). A
// catalog question with a fit at or above it wins. Below it, the model may
// invent one. The value starts conservative on the owner's instruction of
// 2026-08-24, so the agent almost never invents. M-5 sets the real number
// from at least 50 scored questions (D-66).
const DefaultFitThreshold = 0.35

// Agent runs one conversation turn: it maps the user's words onto slots,
// picks the questions, and phrases them.
//
// A turn costs three model calls. The classify role fills the slots. A
// second classify call scores how well each planned catalog row fits the
// user's words, and offers a replacement when the fit is poor (D-25). The
// ask role then phrases what the agent decided to ask. The score has its
// own call on the owner's directive of 2026-08-24: a model that phrases a
// question must not also rate its own wording.
type Agent struct {
	cat       *Catalog
	llm       *llm.Client
	threshold float64
	log       *slog.Logger
	hints     Hints
}

// AgentOption changes one Agent field.
type AgentOption func(*Agent)

// WithFitThreshold overrides the provisional threshold.
func WithFitThreshold(f float64) AgentOption { return func(a *Agent) { a.threshold = f } }

// WithLogger sets the logger that carries the M-4 rows.
func WithLogger(l *slog.Logger) AgentOption { return func(a *Agent) { a.log = l } }

// WithHints supplies the values the catalog rows name in braces. Without
// it, every clause that needs a value is dropped.
func WithHints(h Hints) AgentOption { return func(a *Agent) { a.hints = h } }

// NewAgent builds an agent from a catalog and a model client.
func NewAgent(cat *Catalog, c *llm.Client, opts ...AgentOption) (*Agent, error) {
	if cat == nil || c == nil {
		return nil, fmt.Errorf("questions: NewAgent needs a catalog and a client")
	}
	a := &Agent{cat: cat, llm: c, threshold: DefaultFitThreshold, log: slog.Default()}
	for _, o := range opts {
		o(a)
	}
	return a, nil
}

// Result is one turn's output.
type Result struct {
	Questions []*mtgv1.Question
	Slots     *mtgv1.Slots
	// Ready is true when no slot is left to ask about.
	Ready bool
	// Invented counts the model-invented questions of this turn (M-4).
	Invented int
	// Coverage is the session's M-4 report after this turn.
	Coverage Coverage
}

// Turn maps one user message onto the slots and returns the next
// questions. A frozen session asks nothing (D-68).
func (a *Agent) Turn(ctx context.Context, st *State, message string, acc *llm.Accumulator) (Result, error) {
	if st == nil {
		return Result{}, fmt.Errorf("questions: Turn needs a state")
	}
	st.AddWords(message)
	if st.Ctx.Frozen {
		return Result{Slots: st.Slots, Ready: true, Coverage: st.Metrics()}, nil
	}
	st.Turn++
	if err := a.classify(ctx, st, message, acc); err != nil {
		return Result{}, err
	}
	rows := a.cat.Plan(st.Ctx)
	if len(rows) == 0 {
		// The classify call may have closed a key, so the M-4 report
		// changes even on a turn that asks nothing.
		return Result{Slots: st.Slots, Ready: st.Ready(a.cat), Coverage: st.Metrics()}, nil
	}
	// Resolve every placeholder before a model sees the row. A brace that
	// reaches the model comes back as a question aimed at the user.
	resolved := make(map[string]string, len(rows))
	options := make(map[string][]string, len(rows))
	offered := make(map[string][]string, len(rows))
	for _, r := range rows {
		text, opts, names := resolve(r, st, a.hints)
		resolved[r.ID], options[r.ID], offered[r.ID] = text, opts, names
	}
	scores, err := a.score(ctx, st, message, rows, resolved, acc)
	if err != nil {
		return Result{}, err
	}
	// The agent decides, not the model. A replacement counts only when the
	// catalog fit is under the threshold and the model wrote one.
	chosen := make([]choice, 0, len(rows))
	for _, row := range rows {
		c := choice{Row: row, Text: resolved[row.ID], Options: options[row.ID], Fit: scores[row.ID].Fit, Reason: scores[row.ID].Reason}
		if custom := strings.TrimSpace(scores[row.ID].CustomText); custom != "" && c.Fit < a.threshold {
			switch {
			case nearCopy(custom, c.Text):
				// A reword is not an invention. It costs the user nothing
				// and it counts against the catalog (D-88).
				c.NearCopy, c.Refused = true, custom
				a.log.Info("replacement refused as a reword",
					"row", row.ID, "fit", c.Fit, "overlap", overlap(custom, c.Text),
					"replacement", custom, "row_text", c.Text)
			default:
				// The options stay. A row's option set is its decision
				// space, and a replacement must keep it (the score prompt
				// says so). Gate run 4 of 2026-08-25 replaced the
				// card-pool question twice, and each replacement offered
				// two of the three pool modes, so the user lost
				// owned-only or owned-first (D-37).
				c.Text, c.Invented = custom, true
			}
		}
		chosen = append(chosen, c)
	}
	phrased, err := a.ask(ctx, st, message, chosen, acc)
	if err != nil {
		return Result{}, err
	}
	res := Result{Slots: st.Slots}
	for _, c := range chosen {
		st.AskCount++
		q := &mtgv1.Question{
			Id:       fmt.Sprintf("q%d-%s", st.AskCount, c.Row.ID),
			Slot:     c.Row.Slot,
			Text:     c.Text,
			Options:  c.Options,
			Invented: c.Invented,
			GapScore: c.Fit,
		}
		if p, ok := phrased[c.Row.ID]; ok {
			// The guard keeps a bad phrasing off the wire. It falls back
			// to the resolved catalog text.
			if kept := guard(p.Text, c.Text); kept != c.Text {
				q.Text = kept
			}
			if len(p.Options) > 0 {
				q.Options = p.Options
			}
		}
		if c.Invented {
			// The UI shows both texts, so M-5 can score whether the
			// catalog was enough and whether the invention is better (D-66).
			q.CatalogText = resolved[c.Row.ID]
			res.Invented++
		} else {
			// The catalog text went out, so those names are on the table.
			// An invented question named none of them.
			st.SetOffer(offered[c.Row.ID])
		}
		st.MarkAsked(c.Row.ID, c.Row.StateKey())
		rec := Ask{
			QuestionID: q.Id, RowID: c.Row.ID, Slot: c.Row.Slot, Key: c.Row.StateKey(),
			Invented: c.Invented, Fit: c.Fit, Threshold: a.threshold, Turn: st.Turn,
		}
		if c.Invented {
			rec.CatalogText = resolved[c.Row.ID]
		}
		rec.NearCopy, rec.RefusedText = c.NearCopy, c.Refused
		st.Asks = append(st.Asks, rec)
		a.record(st, c.Row, q, c.Reason)
		res.Questions = append(res.Questions, q)
	}
	res.Coverage = st.Metrics()
	a.log.Info("turn coverage",
		"session", st.SessionID, "turn", st.Turn,
		"asked", res.Coverage.Asked, "catalog", res.Coverage.Catalog,
		"invented", res.Coverage.Invented, "invented_filled", res.Coverage.InventedFilled,
		"median_fit", res.Coverage.MedianFit())
	return res, nil
}

// record writes the M-4 row for one question: which slot, which source,
// and the fit score that allowed an invented question.
func (a *Agent) record(st *State, row Row, q *mtgv1.Question, reason string) {
	source := "catalog"
	if q.Invented {
		source = "invented"
	}
	a.log.Info("question asked",
		"session", st.SessionID, "turn", st.Turn, "question", q.Id,
		"slot", row.Slot, "key", row.StateKey(), "row", row.ID,
		"source", source, "fit", q.GapScore, "threshold", a.threshold,
		"reason", reason, "asked_so_far", len(st.Ctx.Asked))
}

// classifyOut is the classify role's structured output.
type classifyOut struct {
	Format         string   `json:"format"`
	Theme          string   `json:"theme"`
	Colors         []string `json:"colors"`
	CommanderNames []string `json:"commander_names"`
	LockedNames    []string `json:"locked_names"`
	Power          string   `json:"power"`
	PoolRule       string   `json:"pool_rule"`
	BudgetUSD      float64  `json:"budget_usd"`
	ClosedKeys     []string `json:"closed_keys"`
	// DeclinedKeys are the keys the user handed back to the agent. A
	// decline is not an answer: it holds no value, and a default applies
	// (D-93).
	DeclinedKeys []string `json:"declined_keys"`
	Facts        struct {
		NamedCard        bool `json:"named_card"`
		BuyList          bool `json:"buy_list"`
		HouseFormat      bool `json:"house_format"`
		TwoPlans         bool `json:"two_plans"`
		BudgetAmbiguous  bool `json:"budget_ambiguous"`
		PowerCompetitive bool `json:"power_competitive"`
		// WantsSuggestion says the user asked the agent to name a
		// commander. It is what fires the commander_pick row (D-71).
		WantsSuggestion bool `json:"wants_suggestion"`
		// OutOfScope says the user asked for something other than a
		// Magic: The Gathering deck (D-99).
		OutOfScope bool `json:"out_of_scope"`
	} `json:"facts"`
}

func (a *Agent) classify(ctx context.Context, st *State, message string, acc *llm.Accumulator) error {
	open := openKeys(st)
	input, err := json.Marshal(map[string]any{
		"message":     message,
		"slots_known": st.Slots.SlotStates,
		"theme":       st.Slots.Theme,
		"open_keys":   open,
	})
	if err != nil {
		return fmt.Errorf("questions: classify input: %w", err)
	}
	res, err := a.llm.Complete(ctx, llm.RoleClassify, llm.Request{
		Instructions: classifyInstructions,
		Input:        string(input),
		SchemaName:   "slot_fill",
		Schema:       json.RawMessage(classifySchema),
		CacheKey:     st.SessionID,
	}, acc)
	if err != nil {
		return fmt.Errorf("questions: classify: %w", err)
	}
	var out classifyOut
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return fmt.Errorf("questions: classify output: %w", err)
	}
	a.apply(st, out, open)
	return nil
}

// apply writes one classify result onto the state. It never clears a slot
// the session already filled: a value stays until the user replaces it.
func (a *Agent) apply(st *State, out classifyOut, open []string) {
	if id, ok := formatIDs[slotWord(out.Format)]; ok && id != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		st.Slots.Format = &mtgv1.Format{Id: id}
		st.Ctx.Format = id
		st.Close("format")
	}
	if s := strings.TrimSpace(out.Theme); s != "" {
		st.Slots.Theme, st.Ctx.Theme = s, strings.ToLower(s)
		st.Close("theme")
	}
	if len(out.Colors) > 0 {
		st.Slots.Colors = nil
		for _, c := range out.Colors {
			if id, ok := colorIDs[strings.ToUpper(c)]; ok {
				st.Slots.Colors = append(st.Slots.Colors, id)
			}
		}
		if len(st.Slots.Colors) > 0 {
			st.Close("colors")
		}
	}
	// The two name lists are kept apart. Merged, one name that becomes the
	// commander also reads as a card to keep, which is the defect the live
	// run of 2026-08-24 found (D-70).
	for _, name := range append(append([]string(nil), out.CommanderNames...), out.LockedNames...) {
		if name = strings.TrimSpace(name); name != "" {
			st.NamedCards = append([]string{name}, st.NamedCards...)
			st.Ctx.NamedCard = true
		}
	}
	for _, name := range out.LockedNames {
		st.AddLocked(name)
	}
	for _, name := range out.CommanderNames {
		st.SetCommander(name)
	}
	if rule, ok := poolRules[slotWord(out.PoolRule)]; ok {
		st.Slots.PoolRule = rule
		st.Ctx.OwnedMode = rule != mtgv1.PoolRule_POOL_RULE_ANY_CARD
		st.Close("pool_rule")
	}
	if p := power(out.Power); p != nil {
		st.Slots.Power = p
		st.Close("power")
	}
	if out.BudgetUSD > 0 {
		st.Slots.BudgetUsd = out.BudgetUSD
		st.Close("budget")
	}
	// A key closes by name only when two things hold: its question is
	// out, and it carries no typed value. Three gate runs paid for that
	// pair of conditions.
	//
	// The question must be out, because on 2026-08-25 the classifier
	// retired power and the pool rule from "Brago blink deck from my
	// library". The session then called itself complete after one
	// question, and neither key had ever been asked.
	//
	// A typed slot must close on its value, because a name says only
	// "answered" and never says what the answer was. Run 6 of 2026-08-25
	// closed the format by name, so the format value stayed empty. Every
	// row that triggers on the format then stopped firing, five sessions
	// ended with no power level, and PR-8 would have had no format to
	// build from. Asking the format twice is the safe failure. Building
	// a deck with no format is not (D-83).
	for _, k := range out.ClosedKeys {
		k = strings.TrimSpace(k)
		if k == "" {
			continue
		}
		if typedSlots[k] {
			a.log.Warn("classifier closed a typed slot by name and gave no value", "key", k)
			continue
		}
		if st.Slots.GetSlotStates()[k] != mtgv1.SlotState_SLOT_STATE_ASKED {
			a.log.Warn("classifier closed a key whose question is not out", "key", k, "offered", open)
			continue
		}
		st.Close(k)
	}
	// A decline closes any key, typed or not. It carries no value on
	// purpose: the user handed the choice back, and the generator applies
	// the default the corpus names. The question must still be out,
	// because nobody can decline a question they never saw (D-93).
	for _, k := range out.DeclinedKeys {
		k = strings.TrimSpace(k)
		if k == "" {
			continue
		}
		if st.Slots.GetSlotStates()[k] != mtgv1.SlotState_SLOT_STATE_ASKED {
			a.log.Warn("classifier declined a key whose question is not out", "key", k, "offered", open)
			continue
		}
		a.log.Info("the user declined a slot", "key", k)
		st.Skip(k)
		// The format is the one slot the planner routes on. Every power
		// row, and every Commander row, triggers on it. A declined format
		// left it empty, so no power row could ever fire and the session
		// finished with no power level (probe 35 of gate run 11).
		//
		// The corpus names the default under "default answers", so
		// nothing is invented here. The slot keeps the SKIPPED state,
		// which records that the user did not choose it. PR-8 reads the
		// other declined slots the same way (D-98).
		if k == "format" && st.Ctx.Format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
			st.Slots.Format = &mtgv1.Format{Id: DefaultFormat}
			st.Ctx.Format = DefaultFormat
			a.log.Info("a declined format took the corpus default", "format", DefaultFormat.String())
		}
	}
	f := out.Facts
	st.Ctx.NamedCard = st.Ctx.NamedCard || f.NamedCard
	st.Ctx.BuyList = st.Ctx.BuyList || f.BuyList
	st.Ctx.HouseFormat = st.Ctx.HouseFormat || f.HouseFormat
	st.Ctx.TwoPlans = st.Ctx.TwoPlans || f.TwoPlans
	st.Ctx.BudgetAmbiguous = st.Ctx.BudgetAmbiguous || f.BudgetAmbiguous
	st.Ctx.PowerCompetitive = st.Ctx.PowerCompetitive || f.PowerCompetitive
	// The fact is not sticky. A user who asks for a Magic deck after the
	// agent declines is back in scope.
	st.Ctx.OutOfScope = f.OutOfScope
	// A user who asks for a suggestion gets the pick row next turn. A user
	// who then names one closes every commander row, so the fact does not
	// need to be cleared.
	if f.WantsSuggestion {
		// A second request retires the names on the table, so the next
		// round names three others (D-73).
		st.RetireOffer()
	}
	st.Ctx.Suggested = st.Ctx.Suggested || f.WantsSuggestion
	// A word rule can fire when the classifier misses one (corpus 11).
	if slot := Route(st.Ctx.Words); slot == "house_rules" {
		st.Ctx.HouseFormat = st.Ctx.HouseFormat || strings.Contains(st.Ctx.Words, "no ban list")
	}
}

// power maps the classifier's power word onto the proto message.
func power(s string) *mtgv1.PowerLevel {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" || s == "unknown" {
		return nil
	}
	if step, ok := sixtySteps[s]; ok {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: step}}
	}
	if n, err := strconv.Atoi(strings.TrimPrefix(s, "bracket ")); err == nil && n >= 1 && n <= 5 {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: int32(n)}}
	}
	return nil
}

// typedSlots are the slots that carry a value the deck generator reads.
// The classifier fills each one through its own field, so a name in the
// closed_keys list can never close one: the name says "answered" and
// leaves the value empty (D-83).
var typedSlots = map[string]bool{
	"format": true, "theme": true, "colors": true,
	"power": true, "pool_rule": true, "budget": true,
	// The commander closes on a name, and that name closes the color
	// slot and the two other commander rows with it.
	"commander": true,
}

// openKeys lists every key whose question is out. The classifier reads
// them for two lists. It may close an advisory key by name (D-83), and it
// may decline any of them, typed or not (D-93).
func openKeys(st *State) []string {
	var out []string
	for key, state := range st.Slots.GetSlotStates() {
		if state == mtgv1.SlotState_SLOT_STATE_ASKED {
			out = append(out, key)
		}
	}
	// Sorted, because a map gives a new order on every turn, and the
	// input of a model call must not change without a reason.
	sort.Strings(out)
	return out
}

// choice is one question after the agent decided its source.
type choice struct {
	Row      Row
	Text     string
	Options  []string
	Fit      float64
	Reason   string
	Invented bool
	// NearCopy marks a replacement the agent refused as a reword (D-88).
	NearCopy bool
	// Refused is the text of that replacement.
	Refused string
}

// scored is one row as the scoring classifier returned it.
type scored struct {
	RowID string `json:"row_id"`
	// Fit is how well the catalog row fits the user's words, 0 to 1.
	Fit        float64 `json:"fit"`
	CustomText string  `json:"custom_text"`
	Reason     string  `json:"reason"`
}

// score rates the planned catalog rows against the user's words. It is a
// classify-role call of its own, so the model that phrases a question
// never rates its own wording.
func (a *Agent) score(ctx context.Context, st *State, message string, rows []Row, resolved map[string]string, acc *llm.Accumulator) (map[string]scored, error) {
	type rowIn struct {
		ID   string `json:"id"`
		Slot string `json:"slot"`
		Text string `json:"text"`
	}
	in := make([]rowIn, 0, len(rows))
	for _, r := range rows {
		in = append(in, rowIn{ID: r.ID, Slot: r.Slot, Text: resolved[r.ID]})
	}
	input, err := json.Marshal(map[string]any{
		"message":      message,
		"conversation": st.Ctx.Words,
		"theme":        st.Slots.Theme,
		"rows":         in,
		"threshold":    a.threshold,
	})
	if err != nil {
		return nil, fmt.Errorf("questions: score input: %w", err)
	}
	res, err := a.llm.Complete(ctx, llm.RoleClassify, llm.Request{
		Instructions: scoreInstructions,
		Input:        string(input),
		SchemaName:   "score_questions",
		Schema:       json.RawMessage(scoreSchema),
		CacheKey:     st.SessionID,
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("questions: score: %w", err)
	}
	var out struct {
		Scores []scored `json:"scores"`
	}
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("questions: score output: %w", err)
	}
	byRow := make(map[string]scored, len(out.Scores))
	for _, s := range out.Scores {
		byRow[s.RowID] = s
	}
	return byRow, nil
}

// phrasing is one question as the ask role returned it.
type phrasing struct {
	RowID   string   `json:"row_id"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
}

// ask puts the chosen questions into natural words. It never changes what
// a question asks.
func (a *Agent) ask(ctx context.Context, st *State, message string, chosen []choice, acc *llm.Accumulator) (map[string]phrasing, error) {
	type rowIn struct {
		ID      string   `json:"id"`
		Slot    string   `json:"slot"`
		Text    string   `json:"text"`
		Options []string `json:"options"`
	}
	in := make([]rowIn, 0, len(chosen))
	for _, c := range chosen {
		in = append(in, rowIn{ID: c.Row.ID, Slot: c.Row.Slot, Text: c.Text, Options: c.Options})
	}
	input, err := json.Marshal(map[string]any{
		"message":   message,
		"theme":     st.Slots.Theme,
		"questions": in,
	})
	if err != nil {
		return nil, fmt.Errorf("questions: ask input: %w", err)
	}
	res, err := a.llm.Complete(ctx, llm.RoleAsk, llm.Request{
		Instructions: askInstructions,
		Input:        string(input),
		SchemaName:   "phrase_questions",
		Schema:       json.RawMessage(askSchema),
		CacheKey:     st.SessionID,
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("questions: ask: %w", err)
	}
	var out struct {
		Questions []phrasing `json:"questions"`
	}
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("questions: ask output: %w", err)
	}
	byRow := make(map[string]phrasing, len(out.Questions))
	for _, p := range out.Questions {
		byRow[p.RowID] = p
	}
	return byRow, nil
}
