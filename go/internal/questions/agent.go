package questions

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
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
}

// Turn maps one user message onto the slots and returns the next
// questions. A frozen session asks nothing (D-68).
func (a *Agent) Turn(ctx context.Context, st *State, message string, acc *llm.Accumulator) (Result, error) {
	if st == nil {
		return Result{}, fmt.Errorf("questions: Turn needs a state")
	}
	st.AddWords(message)
	if st.Ctx.Frozen {
		return Result{Slots: st.Slots, Ready: true}, nil
	}
	if err := a.classify(ctx, st, message, acc); err != nil {
		return Result{}, err
	}
	rows := a.cat.Plan(st.Ctx)
	if len(rows) == 0 {
		return Result{Slots: st.Slots, Ready: st.Ready(a.cat)}, nil
	}
	// Resolve every placeholder before a model sees the row. A brace that
	// reaches the model comes back as a question aimed at the user.
	resolved := make(map[string]string, len(rows))
	options := make(map[string][]string, len(rows))
	for _, r := range rows {
		text, opts := resolve(r, st, a.hints)
		resolved[r.ID], options[r.ID] = text, opts
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
			c.Text, c.Options, c.Invented = custom, nil, true
		}
		chosen = append(chosen, c)
	}
	phrased, err := a.ask(ctx, st, message, chosen, acc)
	if err != nil {
		return Result{}, err
	}
	res := Result{Slots: st.Slots}
	for i, c := range chosen {
		q := &mtgv1.Question{
			Id:       fmt.Sprintf("q%d-%s", len(st.Ctx.Asked)+i+1, c.Row.ID),
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
			res.Invented++
		}
		st.MarkAsked(c.Row.ID, c.Row.StateKey())
		a.record(st, c.Row, q, c.Reason)
		res.Questions = append(res.Questions, q)
	}
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
	Facts          struct {
		NamedCard        bool `json:"named_card"`
		BuyList          bool `json:"buy_list"`
		Deadline         bool `json:"deadline"`
		HouseFormat      bool `json:"house_format"`
		TwoPlans         bool `json:"two_plans"`
		BudgetAmbiguous  bool `json:"budget_ambiguous"`
		PowerCompetitive bool `json:"power_competitive"`
	} `json:"facts"`
}

func (a *Agent) classify(ctx context.Context, st *State, message string, acc *llm.Accumulator) error {
	open := openKeys(a.cat, st)
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
	if id, ok := formatIDs[out.Format]; ok && id != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
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
	for _, name := range append(append([]string(nil), out.CommanderNames...), out.LockedNames...) {
		if name = strings.TrimSpace(name); name != "" {
			st.NamedCards = append([]string{name}, st.NamedCards...)
			st.Ctx.NamedCard = true
		}
	}
	if rule, ok := poolRules[out.PoolRule]; ok {
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
	// A key closes only when the agent offered it this turn. The live run
	// of 2026-08-24 showed the classifier naming keys it was never asked
	// about, which ended a session with three slots still empty.
	offered := make(map[string]bool, len(open))
	for _, k := range open {
		offered[k] = true
	}
	for _, k := range out.ClosedKeys {
		k = strings.TrimSpace(k)
		if k == "" {
			continue
		}
		if !offered[k] {
			a.log.Warn("classifier named a key it was not offered", "key", k, "offered", open)
			continue
		}
		st.Close(k)
	}
	f := out.Facts
	st.Ctx.NamedCard = st.Ctx.NamedCard || f.NamedCard
	st.Ctx.BuyList = st.Ctx.BuyList || f.BuyList
	st.Ctx.Deadline = st.Ctx.Deadline || f.Deadline
	st.Ctx.HouseFormat = st.Ctx.HouseFormat || f.HouseFormat
	st.Ctx.TwoPlans = st.Ctx.TwoPlans || f.TwoPlans
	st.Ctx.BudgetAmbiguous = st.Ctx.BudgetAmbiguous || f.BudgetAmbiguous
	st.Ctx.PowerCompetitive = st.Ctx.PowerCompetitive || f.PowerCompetitive
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

// openKeys lists the keys the planner would ask about next. The classifier
// reads them so it can close one from free text before a question goes out.
func openKeys(c *Catalog, st *State) []string {
	var out []string
	for _, r := range c.Plan(st.Ctx) {
		out = append(out, r.StateKey())
	}
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
