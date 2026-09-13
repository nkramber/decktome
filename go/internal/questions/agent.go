package questions

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"sort"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
)

// DefaultFitThreshold is the gap-score threshold (D-27, D-69). A catalog
// question with a fit at or above it wins. Below it, the model may
// invent one. The value is conservative, so the agent almost never
// invents. The M-5 sheets read the value from the gate document (D-66).
const DefaultFitThreshold = 0.35

// Agent runs one conversation turn: it maps the user's words onto slots,
// picks the questions, and phrases them.
//
// A turn costs up to three model calls. The classify role fills the
// slots. A second classify call scores how well each planned catalog row
// fits the user's words, and offers a replacement when the fit is poor
// (D-25). The ask role then phrases what the agent decided to ask. A
// turn whose every row is fixed skips the last two calls. The score has
// its own call: a model that phrases a question must not also rate its
// own wording (D-69).
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
	// ChoseCommander is true when the pool offered no commander this turn
	// and the agent took the choice under D-127. The turn then asks
	// nothing about the commander, and the reader must be told why
	// (D-366).
	ChoseCommander bool
	// SetsApplied names the sets this turn read out of the reader's
	// words, when it read any. A reader who writes "the Hobbit set" gets
	// two sets, and a red mark on every card outside them. The turn must
	// say which sets it applied, or the marks explain nothing (D-390).
	SetsApplied []string
	// PreconsApplied names the precons the deck uses no card of, after a
	// turn that read one (D-496). PreconsPartial names the ones the
	// collection does not hold whole (D-497). PreconsNone says the reader
	// excluded their precons and the collection holds none whole, and
	// PreconsUnavailable says no precon table is loaded. The reader hears
	// each case: silence reads as a bug (D-390).
	PreconsApplied     []string
	PreconsPartial     []string
	PreconsNone        bool
	PreconsUnavailable bool
}

// Turn maps one user message onto the slots and returns the next
// questions.
func (a *Agent) Turn(ctx context.Context, st *State, message string, acc *llm.Accumulator) (Result, error) {
	if st == nil {
		return Result{}, fmt.Errorf("questions: Turn needs a state")
	}
	if err := a.classify(ctx, st, message, acc); err != nil {
		return Result{}, err
	}
	a.readFacts(st)
	// Both marks are of this turn alone (D-366, D-390).
	st.Ctx.ChoseCommander = false
	st.setsThisTurn = nil
	rows, resolved := a.plan(st, UserWords(message))
	// The planner has nothing to ask and the session is not ready. That
	// pair proves the turn is stuck: a question is out, the reader
	// answered it, and no value reached the slot. Ask it once more, in
	// this turn, rather than end with no question and no deck (D-599).
	if len(rows) == 0 && !st.Ready(a.cat) {
		if reasked := st.ReaskStalled(); len(reasked) > 0 {
			a.log.Info("the answer did not reach its slot, so the question goes out once more",
				"session", st.SessionID, "keys", reasked)
			rows, resolved = a.plan(st, UserWords(message))
		}
	}
	if len(rows) == 0 {
		// The classify call may have closed a key, so the M-4 report
		// changes even on a turn that asks nothing.
		res := Result{Slots: st.Slots, Ready: st.Ready(a.cat), Coverage: st.Metrics(),
			ChoseCommander: st.Ctx.ChoseCommander, SetsApplied: st.setsThisTurn}
		st.fillPrecons(&res)
		return res, nil
	}
	chosen, err := a.choose(ctx, st, message, rows, resolved.text, acc)
	if err != nil {
		return Result{}, err
	}
	return a.send(ctx, st, message, chosen, resolved, acc)
}

// classify runs the classify call, commits the message, and applies the
// result and the word rules. The message joins the state only after the
// call succeeds, so a failed call leaves no half turn behind.
func (a *Agent) classify(ctx context.Context, st *State, message string, acc *llm.Accumulator) error {
	open := openKeys(st)
	out, err := a.classifyCall(ctx, st, message, open, acc)
	if err != nil {
		return err
	}
	st.AddMessage(message)
	st.Turn++
	// The keys a deck slot fills before this turn. An out-of-scope
	// question closes when the user fills one afterwards (D-196).
	deckKeysBefore := deckKeysFilled(st)
	// The word rules read the user's own words. A quoted question is the
	// agent's text, and its format list is not a two-deck request.
	words := UserWords(message)
	a.apply(ctx, st, out, open, words, acc)
	a.applyWords(st, turnWords{Message: words, Declined: out.DeclinedKeys, Closed: out.ClosedKeys, Open: open})
	a.closeByOption(st, words)
	// The reader picked an option, and the engine sets the slot from it
	// with no model in the path (D-597). It runs after the classifier, so
	// the reader's own choice is the one that stands.
	a.applyOptionAnswers(st)
	a.applyManaPermission(st, words)
	// The scope question closes when the user answers it with a deck.
	// The row offers "Yes, a Magic deck", and a user who writes "a Modern
	// burn deck" instead has said the same thing. Nothing else closed
	// that key, so the session never reported ready (D-196).
	if st.Slots.GetSlotStates()["scope"] == mtgv1.SlotState_SLOT_STATE_ASKED &&
		!st.Ctx.OutOfScope && deckKeysFilled(st) > deckKeysBefore {
		a.log.Info("the user asked for a Magic deck, so the scope question closed",
			"session", st.SessionID)
		st.Skip("scope")
	}
	return nil
}

// applyManaPermission reads the answer to the mana row (D-382). The row
// asks one permission, so a "yes" fills the key and every other answer
// skips it.
//
// The key can not close by an option match: both options start with a
// word the match would take, and the difference between them is the
// whole answer. A skipped key means the deck stays inside the sets,
// which is the literal reading of what the reader asked for.
func (a *Agent) applyManaPermission(st *State, message string) {
	if st.Slots.GetSlotStates()[SlotSetOutsideMana] != mtgv1.SlotState_SLOT_STATE_ASKED {
		return
	}
	row, ok := a.askedRow(st, SlotSetOutsideMana)
	if !ok || row.ID != "set_outside_mana" {
		return
	}
	switch {
	case len(row.Options) > 0 && optionAnswered(message, row.Options[0]), acceptsOffer(message):
		a.log.Info("the reader allowed mana cards from outside the named sets",
			"session", st.SessionID)
		st.Close(SlotSetOutsideMana)
	case len(row.Options) > 1 && optionAnswered(message, row.Options[1]), BareNegative(message):
		a.log.Info("the reader kept the mana base inside the named sets",
			"session", st.SessionID)
		st.Skip(SlotSetOutsideMana)
	}
}

// readFacts hands the slots of this turn to the hint source and reads
// the planner facts back.
//
// The hint source reads the slots as they stand now. The caller built it
// before the turn, so the colors the classifier just filled would
// otherwise be invisible to it (D-124). The facts it answers move with
// the slots, so they are read again here and not only before the turn:
// a one-message answer filled the format and the theme, and the
// thin-theme count was never taken (D-198).
func (a *Agent) readFacts(st *State) {
	if pa, ok := a.hints.(PairAware); ok && st.Ctx.WantPair {
		pa.UseWantPair(true, st.Ctx.WantBackground)
	}
	if sa, ok := a.hints.(SlotAware); ok {
		sa.UseSlots(st.Slots.GetFormat().GetId(), st.Slots.GetColors(), st.Slots.GetPoolRule())
	}
	if ss, ok := a.hints.(SetAware); ok {
		ss.UseSets(st.Slots.GetSetCodes())
	}
	if ba, ok := a.hints.(BracketAware); ok {
		ba.UseBracket(st.Slots.GetPower().GetBracket())
	}
	if fs, ok := a.hints.(FactSource); ok {
		RefreshFacts(st, fs)
	}
	// The pick row asks again only when the names on the table changed.
	// The classify call is what changes them: a refusal empties the
	// table, and the color check drops a name the colors exclude (D-163).
	st.Ctx.OfferChanged = st.OfferChanged()
	// The decline rows ask again only when the user names another format
	// this app does not build (D-210).
	st.Ctx.BadFormatChanged = st.BadFormatChanged()
	// The set row follows the same rule (D-376).
	st.Ctx.SetChanged = st.BadSetChanged()
	// So does the precon row (D-496).
	st.Ctx.PreconChanged = st.BadPreconChanged()
	// So does the row that asks which card a commander name means (F-75).
	st.Ctx.CommanderChanged = st.BadCommanderChanged()
	// A named card that can lead a deck may fix the deck's color
	// identity, and nothing has settled its role yet. The color row
	// waits, or it asks for colors the commander already decides (D-388).
	st.Ctx.NamedLeader = false
	if cc, ok := a.hints.(CommanderChecker); ok && !st.Ctx.CommanderSet && !st.Ctx.Filled["named_card_role"] {
		for _, name := range st.NamedCards {
			if lead, known := cc.CanLead(name); known && lead {
				st.Ctx.NamedLeader = true
				break
			}
		}
	}
	// The mana row needs the count the sets offer against the count the
	// deck wants (D-382). It only matters while a set limit is on and
	// the key is open.
	st.Ctx.ThinSetMana = false
	if ms, ok := a.hints.(ManaSource); ok && st.Ctx.SetLimited && !st.Ctx.Filled[SlotSetOutsideMana] {
		thin, _, _ := ms.ThinSetMana(st.Slots.GetSetCodes(), st.Slots.GetFormat().GetId(),
			st.Slots.GetColors(), st.Slots.GetPower())
		st.Ctx.ThinSetMana = thin
	}
}

// resolvedRows holds the placeholder-free text, the options, and the
// commander names of every planned row.
type resolvedRows struct {
	text    map[string]string
	options map[string][]string
	offered map[string][]string
}

// plan picks the rows of this turn and resolves their placeholders. A
// brace that reaches a model comes back as a question aimed at the user.
func (a *Agent) plan(st *State, message string) ([]Row, resolvedRows) {
	rows := a.cat.Plan(st.Ctx)
	// A user can answer a question in the same message that raises it, and
	// the agent must not ask it (D-122).
	//
	// A closed key can free a row that waited on it, so the planner runs
	// once more.
	if a.closeAnsweredRows(st, rows, message) {
		rows = a.cat.Plan(st.Ctx)
	}
	res := resolvedRows{
		text:    make(map[string]string, len(rows)),
		options: make(map[string][]string, len(rows)),
		offered: make(map[string][]string, len(rows)),
	}
	live := rows[:0:0]
	for _, r := range rows {
		text, opts, names := resolve(r, st, a.hints)
		// A row that exists to name commanders can not do its job when
		// the pool gives none. The agent chooses instead, and the skipped
		// state records that nobody picked (D-127).
		if len(commanderKeysIn(r.Text)) > 0 && len(names) == 0 {
			a.log.Info("no commander to offer, so the agent chooses",
				"session", st.SessionID, "row", r.ID)
			// The reader hears about this. A refusal of the names on the
			// table that is answered with silence reads as a bug (D-366).
			st.Ctx.ChoseCommander = true
			st.Skip(r.StateKey())
			// The agent has taken the choice, so the slot is settled too, and
			// the session must not report ready with the commander never asked
			// (D-127). A slot the user filled stays filled: a row with its own
			// key must not overwrite the commander the user named (D-131).
			if r.Slot != r.StateKey() && !st.Ctx.Filled[r.Slot] {
				st.Skip(r.Slot)
			}
			continue
		}
		live = append(live, r)
		res.text[r.ID], res.options[r.ID], res.offered[r.ID] = text, opts, names
	}
	return live, res
}

// choose decides the source of every planned row: the catalog row as
// written, or the model's replacement. The agent decides, not the model.
// A replacement counts only when the catalog fit is under the threshold
// and the model wrote one.
//
// A turn whose every row is fixed makes no score call. A fixed row goes
// out as written whatever the score says, so the call could change
// nothing (D-117).
func (a *Agent) choose(ctx context.Context, st *State, message string, rows []Row, resolved map[string]string, acc *llm.Accumulator) ([]choice, error) {
	scores := map[string]scored{}
	if open := notFixed(rows); len(open) > 0 {
		var err error
		scores, err = a.score(ctx, st, message, open, resolved, acc)
		if err != nil {
			return nil, err
		}
	}
	chosen := make([]choice, 0, len(rows))
	for _, row := range rows {
		sc, ok := scores[row.ID]
		switch {
		case ok:
		case row.Fixed:
			sc = scored{RowID: row.ID, Fit: a.threshold, Reason: "fixed row, not scored"}
		default:
			// The model omitted the row. Its zero value would enter the
			// M-4 record as a fit of 0, which reads as the worst catalog
			// fit ever measured. The catalog is the default source (D-25),
			// so the row goes out as written with the threshold as its fit.
			a.log.Warn("the score call omitted a row, so the catalog row goes out as written",
				"session", st.SessionID, "row", row.ID)
			sc = scored{RowID: row.ID, Fit: a.threshold, Reason: "no score returned"}
		}
		sc.Fit = clampFit(sc.Fit)
		c := choice{Row: row, Text: resolved[row.ID], Fit: sc.Fit, Reason: sc.Reason}
		if custom := strings.TrimSpace(sc.CustomText); custom != "" && c.Fit < a.threshold {
			switch {
			case row.Fixed:
				// A row that states what the app does or does not do goes
				// out as written. A replacement drops the sentence that
				// names the limit, and that sentence is the point.
				c.NearCopy, c.Refused = true, custom
				a.log.Info("replacement refused: the row states an app limit",
					"row", row.ID, "fit", c.Fit, "replacement", custom)
			case nearCopy(custom, c.Text):
				// A reword is not an invention. It costs the user nothing
				// and it counts against the catalog (D-88).
				c.NearCopy, c.Refused = true, custom
				a.log.Info("replacement refused as a reword",
					"row", row.ID, "fit", c.Fit, "overlap", overlap(custom, c.Text),
					"replacement", custom, "row_text", c.Text)
			default:
				// The options stay. A row's option set is its decision
				// space, and a replacement must keep it, or the user loses a
				// pool mode (D-37).
				c.Text, c.Invented = custom, true
			}
		}
		chosen = append(chosen, c)
	}
	return chosen, nil
}

// notFixed returns the rows a score can change. A fixed row goes out as
// written whatever its score, so it never reaches the score call
// (D-117).
func notFixed(rows []Row) []Row {
	var out []Row
	for _, r := range rows {
		if !r.Fixed {
			out = append(out, r)
		}
	}
	return out
}

// send phrases the chosen questions, records them, and builds the
// result.
//
// A fixed row never reaches the ask role. Making it unreplaceable is
// not enough: a rephrasing drops the sentence that names the limit
// (D-117, D-131).
func (a *Agent) send(ctx context.Context, st *State, message string, chosen []choice, resolved resolvedRows, acc *llm.Accumulator) (Result, error) {
	var phrase []choice
	for _, c := range chosen {
		if !c.Row.Fixed {
			c.Options = resolved.options[c.Row.ID]
			phrase = append(phrase, c)
		}
	}
	phrased := map[string]phrasing{}
	if len(phrase) > 0 {
		var err error
		phrased, err = a.ask(ctx, st, message, phrase, acc)
		if err != nil {
			return Result{}, err
		}
	}
	res := Result{Slots: st.Slots, ChoseCommander: st.Ctx.ChoseCommander, SetsApplied: st.setsThisTurn}
	st.fillPrecons(&res)
	for _, c := range chosen {
		st.AskCount++
		q := &mtgv1.Question{
			Id:        fmt.Sprintf("q%d-%s", st.AskCount, c.Row.ID),
			Slot:      c.Row.Slot,
			Text:      c.Text,
			Options:   resolved.options[c.Row.ID],
			Invented:  c.Invented,
			GapScore:  c.Fit,
			Closed:    c.Row.Closed && !c.Invented,
			NoDecline: c.Row.NoDecline && !c.Invented,
		}
		if p, ok := phrased[c.Row.ID]; ok && !c.Row.Fixed {
			// The guard keeps a bad phrasing off the wire. It falls back
			// to the resolved catalog text.
			if kept := guard(c.Row.ID, p.Text, c.Text); kept != c.Text {
				q.Text = kept
			}
			// A closed row keeps the catalog options. The UI offers them
			// as the whole answer space, and the option net compares the
			// answer with the catalog words (D-295, D-119).
			if len(p.Options) > 0 && !c.Row.Closed {
				q.Options = p.Options
			}
		}
		if c.Invented {
			// The UI shows both texts, so M-5 can score whether the
			// catalog was enough and whether the invention is better (D-66).
			q.CatalogText = resolved.text[c.Row.ID]
			res.Invented++
		}
		// The names the resolver produced are on the table, whatever the
		// wording. An invented pick question still carries them in its
		// options, and the user answers "the first" against that list.
		// Before D-121 an invented question recorded no offer, so the
		// next turn could not read an ordinal answer.
		st.SetOffer(resolved.offered[c.Row.ID])
		// A row that repeats only on a change records what it just sent,
		// so the next turn can tell a new list from the same list.
		if c.Row.RepeatOnChange {
			st.RecordAskedOffer(resolved.offered[c.Row.ID])
			if declinesFormat(c.Row.ID) {
				st.RecordAskedBadFormat()
			}
			// The set row follows the same rule (D-376). Without this it
			// asked the same question every turn until the reader
			// answered it, which gate run 29 showed twice in one
			// conversation.
			if c.Row.StateKey() == SlotSetUnresolved {
				st.RecordAskedSet()
			}
			// The row of F-75 follows it too, so a reader who repeats a
			// name this app can not settle reads the sentence once.
			if c.Row.StateKey() == SlotCommanderUnresolved {
				st.RecordAskedCommander()
			}
		}
		st.MarkAsked(c.Row.ID, c.Row.StateKey(), c.Row.Slot)
		rec := Ask{
			QuestionID: q.Id, RowID: c.Row.ID, Slot: c.Row.Slot, Key: c.Row.StateKey(),
			Invented: c.Invented, Fit: c.Fit, Threshold: a.threshold, Turn: st.Turn,
		}
		if c.Invented {
			rec.CatalogText = resolved.text[c.Row.ID]
		}
		rec.NearCopy, rec.RefusedText = c.NearCopy, c.Refused
		// Keep the resolved row whenever the phrasing changed it. The
		// guard read the resolved text, so the M-5 sheet must show it (D-116).
		if q.GetText() != resolved.text[c.Row.ID] {
			rec.ResolvedText = resolved.text[c.Row.ID]
		}
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
	// NamedCards are the cards the user named without a role. "Build
	// around Grist, the Hunger Tide" is such a message. The role question
	// exists for exactly this case, and the classifier used to report the
	// card as the commander, so the question never fired (D-118).
	NamedCards []string `json:"named_cards"`
	// SetNames are the sets the reader wants the deck built from, in the
	// reader's own words. A set is a constraint the app applies now, so
	// the words no longer reach the theme alone (D-373).
	SetNames []string `json:"set_names"`
	// SetGroups are the franchises the reader names as a group of sets,
	// "sets with Marvel characters", in the reader's own words. A group
	// reaches every family of the franchise, and one product name is a
	// set and never a group (D-525).
	SetGroups []string `json:"set_groups"`
	// PreconNames are the precons the reader wants the deck to use no
	// card of, in the reader's own words (D-496). An upgrade names no
	// precon here: that is the word rule of D-113.
	PreconNames []string `json:"precon_names"`
	Power       string   `json:"power"`
	PoolRule    string   `json:"pool_rule"`
	BudgetUSD   float64  `json:"budget_usd"`
	// BudgetScope is "buy", "deck", or "unknown". The budget-scope row
	// asks it, and nothing stored the answer before D-238.
	BudgetScope string `json:"budget_scope"`
	// HouseRules is what the user means by "anything goes", in the
	// user's own words. The house-rules row asks it, and the build reads
	// the answer (D-3, D-265).
	HouseRules string   `json:"house_rules"`
	ClosedKeys []string `json:"closed_keys"`
	// DeclinedKeys are the keys the user handed back to the agent. A
	// decline is not an answer: it holds no value, and a default applies
	// (D-93).
	DeclinedKeys []string `json:"declined_keys"`
	Facts        struct {
		NamedCard        bool `json:"named_card"`
		BuyList          bool `json:"buy_list"`
		HouseFormat      bool `json:"house_format"`
		BudgetAmbiguous  bool `json:"budget_ambiguous"`
		PowerCompetitive bool `json:"power_competitive"`
		// WantsSuggestion says the user asked the agent to name a
		// commander. It is what fires the commander_pick row (D-71).
		WantsSuggestion bool `json:"wants_suggestion"`
		// OutOfScope says the user asked for something other than a
		// Magic: The Gathering deck (D-99).
		OutOfScope bool `json:"out_of_scope"`
		// ExcludePrecons says the user wants no card from the precons
		// they own, and named no product: "not from my precons" (D-496).
		ExcludePrecons bool `json:"exclude_precons"`
	} `json:"facts"`
}

// classifyCall runs the classify role and reads its answer. It changes
// no state.
//
// prior_messages carries the user's earlier messages, so the model reads
// them instead of a rule that told it to repeat values it never saw
// (D-90). The provider cache keys on the instruction prefix and the
// session, and the input differs on every turn in any case, so the list
// costs no cache hit.
func (a *Agent) classifyCall(ctx context.Context, st *State, message string, open []string, acc *llm.Accumulator) (classifyOut, error) {
	prior := st.Prior()
	if prior == nil {
		prior = []string{}
	}
	input, err := json.Marshal(map[string]any{
		"message":            message,
		"prior_messages":     prior,
		"slots_known":        st.Slots.SlotStates,
		"theme":              st.Slots.Theme,
		"open_keys":          open,
		"offered_commanders": st.CurrentOffer,
		// The format the agent offered in place of one it does not build.
		// A "yes" fills the format from it (D-112).
		"nearest_format": st.NearestFormat,
	})
	if err != nil {
		return classifyOut{}, fmt.Errorf("questions: classify input: %w", err)
	}
	res, err := a.llm.Complete(ctx, llm.RoleClassify, llm.Request{
		Instructions: classifyInstructions,
		Input:        string(input),
		SchemaName:   "slot_fill",
		Schema:       json.RawMessage(classifySchema),
		CacheKey:     st.SessionID,
	}, acc)
	if err != nil {
		return classifyOut{}, fmt.Errorf("questions: classify: %w", err)
	}
	var out classifyOut
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return classifyOut{}, fmt.Errorf("questions: classify output: %w", err)
	}
	return out, nil
}

// minOptionMatch is the shortest option that may close a key by itself.
// "Yes" and "No" are too common to read as an answer to one question.
const minOptionMatch = 8

// OptionAnswer is one option the reader picked, by the id of the
// question and the index of the option. The caller reads them off the
// structured answers of the request, and they live for one turn.
type OptionAnswer struct {
	QuestionID string
	Index      int
}

// applyOptionAnswers sets a slot from the option the reader picked
// (D-597).
//
// The client sends the question id and the option index, which name one
// value exactly. Before this the index became the option text, the text
// joined the message, and the slot filled only when the classifier wrote
// a string the slot could read. A model that echoed the option left the
// slot asked, and the turn then asked nothing and built nothing (F-70).
// Seventeen of the eighteen rows with options carry a typed slot, so
// every one of them had that failure in it.
//
// A row with no option values is unchanged: the classifier still reads
// it, and CloseStalled is still the net under it (D-351).
func (a *Agent) applyOptionAnswers(st *State) {
	for _, ans := range st.OptionAnswers {
		ask, ok := st.askOf(ans.QuestionID)
		if !ok {
			continue
		}
		row, ok := a.cat.Row(ask.RowID)
		if !ok {
			continue
		}
		// The row of F-75 offers card names, and the catalog can not hold
		// them: they change with the name the reader wrote (D-606).
		if row.StateKey() == SlotCommanderUnresolved {
			a.applyCommanderOption(st, ans.Index)
			continue
		}
		if ans.Index < 0 || ans.Index >= len(row.OptionValues) {
			continue
		}
		value := row.OptionValues[ans.Index]
		if value == "" || !a.setSlotValue(st, row, ask.Key, value) {
			continue
		}
		a.log.Info("the reader picked an option, so the slot took its value with no model in the path",
			"session", st.SessionID, "row", row.ID, "key", ask.Key, "value", value)
	}
}

// setSlotValue writes one typed value onto its slot and closes the key.
// It answers false for a slot it does not type, so the caller leaves
// that row to the classifier.
func (a *Agent) setSlotValue(st *State, row Row, key, value string) bool {
	switch row.Slot {
	case "power":
		p := power(value)
		if p == nil {
			return false
		}
		st.Slots.Power = p
	case "format":
		id, ok := formatIDs[slotWord(value)]
		if !ok {
			return false
		}
		st.Slots.Format = &mtgv1.Format{Id: id}
	case "pool_rule":
		rule, ok := poolRules[slotWord(value)]
		if !ok {
			return false
		}
		// An owned rule needs a collection behind it (D-371).
		if rule != mtgv1.PoolRule_POOL_RULE_ANY_CARD && !st.Ctx.HasCollection {
			rule = mtgv1.PoolRule_POOL_RULE_ANY_CARD
		}
		st.Slots.PoolRule = rule
	default:
		return false
	}
	st.Close(key)
	return true
}

// applyCommanderOption reads the answer to the row that asks which card
// a commander name means (F-75). The index names one card of the list
// that went out, and the index after the last one is "None of these"
// (D-607).
func (a *Agent) applyCommanderOption(st *State, index int) {
	options := st.CommanderOptions
	switch {
	case index < 0 || index > len(options) || len(options) == 0:
		return
	case index == len(options):
		a.log.Info("the reader wanted none of the cards of that name, so the app offers its own",
			"session", st.SessionID, "name", st.UnresolvedCommander)
		st.DropUnresolvedCommander()
	default:
		a.log.Info("the reader picked the card the name meant, with no model in the path",
			"session", st.SessionID, "name", st.UnresolvedCommander, "card", options[index])
		st.CommanderResolvedName(options[index])
	}
}

// closeByOption closes an advisory key when the user repeats one of the
// options that question offered. It is a net under the classifier, and it
// closes nothing the classifier already closed.
//
// The classifier can leave the key open when the user repeats an option
// word for word. The house-format row waits on that key, so it would
// never fire (D-119).
//
// A typed slot is left alone. Such a slot closes on its value, and a
// name says only "answered" (D-83).
func (a *Agent) closeByOption(st *State, message string) {
	msg := strings.ToLower(message)
	for key, state := range st.Slots.GetSlotStates() {
		if state != mtgv1.SlotState_SLOT_STATE_ASKED || typedSlots[key] {
			continue
		}
		row, ok := a.askedRow(st, key)
		if !ok {
			continue
		}
		for _, opt := range row.Options {
			if !optionAnswered(msg, opt) {
				continue
			}
			a.log.Info("the answer repeated an option, so the key closed",
				"session", st.SessionID, "key", key, "option", opt)
			st.Close(key)
			break
		}
	}
}

// optionPrefixes are the words an option puts before its real content.
// The user repeats the content and drops the prefix: the option reads
// "Yes, the normal limits" and the answer reads "The normal limits hold".
var optionPrefixes = []string{"yes, ", "no, ", "yes ", "no "}

// optionAnswered reports whether a message repeats one option.
func optionAnswered(msg, opt string) bool {
	o := strings.ToLower(strings.TrimSpace(opt))
	if len(o) >= minOptionMatch && strings.Contains(msg, o) {
		return true
	}
	for _, p := range optionPrefixes {
		if !strings.HasPrefix(o, p) {
			continue
		}
		rest := strings.TrimSpace(strings.TrimPrefix(o, p))
		if len(rest) >= minOptionMatch && strings.Contains(msg, rest) {
			return true
		}
	}
	return false
}

// closeAnsweredRows closes the key of a planned row that the user has
// already answered, and reports whether it closed anything.
//
// It reads the options alone, so it needs the same evidence closeByOption
// needs: the message repeats one of the answers the row offers. A typed
// slot is left alone, because it closes on its value (D-83).
func (a *Agent) closeAnsweredRows(st *State, rows []Row, message string) bool {
	msg := strings.ToLower(message)
	closed := false
	for _, r := range rows {
		key := r.StateKey()
		if typedSlots[key] || st.Ctx.Filled[key] {
			continue
		}
		for _, opt := range r.Options {
			if !optionAnswered(msg, opt) {
				continue
			}
			a.log.Info("the message answered a question before it went out",
				"session", st.SessionID, "row", r.ID, "key", key, "option", opt)
			st.Close(key)
			closed = true
			break
		}
	}
	return closed
}

// askedRow finds the row behind the newest open question for one key.
func (a *Agent) askedRow(st *State, key string) (Row, bool) {
	for i := len(st.Asks) - 1; i >= 0; i-- {
		if st.Asks[i].Key == key && !st.Asks[i].Filled {
			return a.cat.Row(st.Asks[i].RowID)
		}
	}
	return Row{}, false
}

// questionOut reports whether the newest open question for a key came
// from the named row.
func (a *Agent) questionOut(st *State, key, rowID string) bool {
	if st.Slots.GetSlotStates()[key] != mtgv1.SlotState_SLOT_STATE_ASKED {
		return false
	}
	row, ok := a.askedRow(st, key)
	return ok && row.ID == rowID
}

// turnWords is what a word rule reads: the message of this turn and the
// keys the classifier assigned it to.
type turnWords struct {
	Message string
	// Declined and Closed are the keys the classifier put the message
	// under. A delegation reads them to learn which question the user
	// answered (D-147).
	Declined, Closed []string
	// Open are the keys that were out when the turn began. A rule reads
	// them to learn which question a value answers (D-288).
	Open []string
}

// hasKey reports whether the classifier assigned the message to a key.
func (w turnWords) hasKey(key string) bool {
	for _, list := range [][]string{w.Declined, w.Closed} {
		for _, k := range list {
			if strings.TrimSpace(k) == key {
				return true
			}
		}
	}
	return false
}

// assigned reports whether the classifier assigned the message to any
// key at all.
func (w turnWords) assigned() bool { return len(w.Declined)+len(w.Closed) > 0 }

// wordRule is one deterministic rule. The rules run in table order, and
// TestWordRulesRunInOrder holds that order.
type wordRule struct {
	name  string
	apply func(a *Agent, st *State, in turnWords)
}

// wordRules are the deterministic word rules, in the order they run.
// Every rule here needs no model call, so it costs nothing and it can not
// drift between runs. They fill a gap the classify role left, and they
// never replace a value it gave.
//
// The order matters in four places. The format rules run first, because
// every later row routes on the format. The nearest-format acceptance
// runs before the decline rule, so a "yes" fills the format before the
// decline rule reads the same message. The card-in-the-99 rule runs
// before the commander rules, because a card in the 99 settles the role
// row that the commander row waits on. The power inference runs last, since
// it reads the format and the competitive fact the other rules set.
var wordRules = []wordRule{
	{"format_from_words", ruleFormatFromWords},
	{"accept_nearest_format", ruleAcceptNearestFormat},
	{"unsupported_format", ruleUnsupportedFormat},
	{"one_deck", ruleOneDeck},
	{"precon", rulePrecon},
	{"proxy_user", ruleProxyUser},
	{"budget_scope", ruleBudgetScope},
	{"no_spending_limit", ruleNoSpendingLimit},
	{"buy_list", ruleBuyList},
	{"colorless", ruleColorless},
	{"cedh", ruleCEDH},
	{"house_rules", ruleHouseRules},
	{"card_in_the_99", ruleCardInThe99},
	{"named_card_as_commander", ruleNamedCardAsCommander},
	{"swap_commander", ruleSwapCommander},
	{"commander_pair", ruleCommanderPair},
	{"delegate_commander", ruleDelegateCommander},
	{"delegate_colors", ruleDelegateColors},
	{"format_from_named_leader", ruleFormatFromNamedLeader},
	{"pick_by_place", rulePickByPlace},
	{"refuse_offer", ruleRefuseOffer},
	{"infer_power", ruleInferPower},
}

// applyWords runs the word rules in table order.
func (a *Agent) applyWords(st *State, in turnWords) {
	for _, r := range wordRules {
		r.apply(a, st, in)
	}
}

// ruleFormatFromWords reads a format the classifier missed. "A land
// destruction Commander deck" names the format, and "not as my
// commander" names no format and can only mean Commander (D-116).
//
// It reads this message alone (D-125).
func ruleFormatFromWords(a *Agent, st *State, in turnWords) {
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return
	}
	if id, ok := FormatFromWords(in.Message); ok {
		a.setFormat(st, id)
		a.log.Info("a word rule read the format the classifier left empty",
			"session", st.SessionID, "format", id.String())
	}
}

// ruleAcceptNearestFormat fills the format when the user accepts the
// nearest one. The decline row offers "Yes, use the nearest format", and
// the slot must fill on that answer (D-112).
//
// The rule fires only while the decline row is the open format question,
// so a bare "yes" reaches nothing else.
func ruleAcceptNearestFormat(a *Agent, st *State, in turnWords) {
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED || !a.questionOut(st, "format", "format_unsupported") {
		return
	}
	id, ok := formatIDs[slotWord(st.NearestFormat)]
	if !ok || id == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return
	}
	msg := strings.ToLower(in.Message)
	if !optionAnswered(msg, "Yes, use the nearest format") && !acceptsOffer(in.Message) {
		return
	}
	a.log.Info("the user accepted the nearest format",
		"session", st.SessionID, "format", id.String())
	a.setFormat(st, id)
}

// setFormat writes a format the words gave. A filled format ends the
// decline: the row that declined the other format has nothing to say.
func (a *Agent) setFormat(st *State, id mtgv1.FormatId) {
	st.Slots.Format = &mtgv1.Format{Id: id}
	st.Ctx.Format = id
	st.Ctx.UnsupportedFormat, st.Ctx.NoNearFormat = false, false
	st.Close("format")
}

// ruleUnsupportedFormat reads a format this app does not build (D-112).
//
// It reads this message alone. A user who names one after a supported
// format is filled gets the decline row, through the path that a format
// change takes: the open questions retire and the format reopens
// (D-125).
func ruleUnsupportedFormat(a *Agent, st *State, in turnWords) {
	name, near, ok := unsupportedFormat(in.Message)
	declared := st.Ctx.Asked["format_unsupported"] || st.Ctx.Asked["format_unsupported_open"]
	if !ok {
		// The row states the limit once, and it asks again only while
		// the user keeps naming the format (D-158, narrows D-157).
		if declared {
			st.Ctx.UnsupportedFormat = false
		}
		return
	}
	if st.Ctx.Filled["format"] {
		a.log.Info("the user named a format this app does not build after a format was filled, so the format reopens",
			"session", st.SessionID, "from", st.Ctx.Format.String(), "format", name)
		st.RetireOutstanding(a.cat)
		st.ReopenRows(a.cat, "format")
		st.Slots.Format, st.Ctx.Format = nil, mtgv1.FormatId_FORMAT_ID_UNSPECIFIED
	} else if !declared && !st.Ctx.UnsupportedFormat {
		// A newly named unsupported format retires the questions that
		// are out. The plain format row is one of them, and D-126
		// blocks every other row on that key while it waits, so the
		// decline could never go out (D-157).
		// "Declared" is sticky, and the firing fact is not. The fact
		// alone oscillates from turn to turn, and the row would ask a
		// third time (D-158).
		a.log.Info("the user named a format this app does not build",
			"session", st.SessionID, "format", name)
		st.RetireOutstanding(a.cat)
	}
	st.UnsupportedFormatName, st.NearestFormat = name, near
	st.Ctx.UnsupportedFormat = true
	// Historic and Timeless name no substitute, so a second row
	// asks which format to build instead (D-146).
	st.Ctx.NoNearFormat = near == ""
}

// ruleOneDeck reads a request for two decks from one message, and never
// from the whole conversation. A user who changes the format across two
// turns has not asked for two decks (D-112).
//
// The fact is not sticky. A second request on a later turn raises it
// again, and the one-deck row says its sentence again (D-112).
func ruleOneDeck(a *Agent, st *State, in turnWords) {
	st.Ctx.TwoDecks = oneDeckRequest(in.Message)
	switch {
	case st.Ctx.TwoDecks && st.Ctx.Filled["deck_count"]:
		a.log.Info("the user asked for two decks again, so the one-deck question reopens",
			"session", st.SessionID)
		st.ReopenRows(a.cat, "deck_count")
	case !st.Ctx.TwoDecks && st.Slots.GetSlotStates()["deck_count"] == mtgv1.SlotState_SLOT_STATE_ASKED:
		// The one-deck question is out, and the user wrote about one
		// deck. That is the answer. Nothing else closed the key, so the
		// session never reported ready (D-196).
		a.log.Info("the user named one deck, so the one-deck question closed",
			"session", st.SessionID)
		st.Skip("deck_count")
	}
}

// rulePrecon reads a precon, which names the card pool (D-113).
func rulePrecon(_ *Agent, st *State, _ turnWords) {
	if !preconRequest(st.Ctx.Words) {
		return
	}
	st.Ctx.Precon = true
	if st.PreconName == "" && len(st.NamedCards) > 0 {
		st.PreconName = st.NamedCards[0]
	}
}

// ruleProxyUser closes the budget for a user who proxies every card.
// Such a user has no budget, so no budget question goes out. "No
// proxies" is not such a user, and anyPhrase reads the negation (D-111).
func ruleProxyUser(a *Agent, st *State, _ turnWords) {
	if !proxyUser(st.Ctx.Words) || st.Ctx.Filled["budget"] {
		return
	}
	st.Skip("budget")
	// No budget means no scope to ask about (D-253).
	st.Skip("budget_scope")
	a.log.Info("the user proxies their cards, so the budget slot is closed",
		"session", st.SessionID)
}

// ruleBudgetScope reads the scope of a cap from the words. A message
// that names the buy list names the scope with it, so the scope row has
// its answer (D-253). "The whole deck" is the other option of the row,
// and the scope is typed, so the words must carry the value (D-238).
func ruleBudgetScope(a *Agent, st *State, in turnWords) {
	// The user's own words for the scope win over a scope the classifier
	// inferred on an earlier turn: "the 100 caps the cards I buy" settles
	// it whatever came before (D-535).
	switch {
	case namesTheBuyList(in.Message):
		st.Slots.BudgetScope = mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY
		a.log.Info("the user named the buy list, so the budget scope is the cards to buy",
			"session", st.SessionID)
	case namesTheWholeDeck(in.Message) && st.Slots.GetSlotStates()["budget_scope"] == mtgv1.SlotState_SLOT_STATE_ASKED:
		st.Slots.BudgetScope = mtgv1.BudgetScope_BUDGET_SCOPE_WHOLE_DECK
		a.log.Info("the user named the whole deck, so the budget scope is the deck value",
			"session", st.SessionID)
	case st.Slots.GetBudgetScope() != mtgv1.BudgetScope_BUDGET_SCOPE_UNSPECIFIED:
		return
	case slices.Contains(in.Open, "budget") && st.Slots.GetBudgetUsd() > 0:
		// The budget row asks "Is there a budget for cards to buy?", so a
		// number that answers it is a cap on the cards to buy, and the
		// scope row has its answer (D-288).
		st.Slots.BudgetScope = mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY
		a.log.Info("the number answers the cards-to-buy question, so the budget scope is the cards to buy",
			"session", st.SessionID)
	default:
		return
	}
	st.Close("budget_scope")
}

// ruleNoSpendingLimit closes the budget for a user who named no cap. It
// reads the message alone, so a cap the user names later still closes
// the slot on its value (D-168).
func ruleNoSpendingLimit(a *Agent, st *State, in turnWords) {
	if !noSpendingLimit(in.Message) || st.Ctx.Filled["budget"] {
		return
	}
	st.Skip("budget")
	a.log.Info("the user answered the budget row without a number, so the slot is closed",
		"session", st.SessionID)
}

// ruleBuyList sets the buy-list fact. A buy list exists when the user
// owns nothing to build from, or when the pool rule lets the deck hold a
// card they do not own. The corpus fires the budget row on a buy list or
// an any-card pool, and a model fact alone misses most of them (D-168).
func ruleBuyList(_ *Agent, st *State, _ turnWords) {
	if !st.Ctx.HasCollection || buysCards(st.Slots.GetPoolRule()) {
		st.Ctx.BuyList = true
	}
}

// ruleColorless closes the color slot on "colorless". It is an answer to
// the color question, and no model call can report it: the classify
// schema holds the five colors alone (D-165).
func ruleColorless(a *Agent, st *State, in turnWords) {
	if !colorlessRequest(in.Message) || st.Ctx.Filled["colors"] {
		return
	}
	st.Skip("colors")
	a.log.Info("the user asked for a colorless deck, so the color slot is closed",
		"session", st.SessionID)
}

// ruleCEDH fills bracket 5 for cEDH. It is bracket 5 by definition, so
// the bracket question has its answer (D-164).
func ruleCEDH(a *Agent, st *State, _ turnWords) {
	if !cedhRequest(st.Ctx.Words) || st.Slots.GetPower() != nil {
		return
	}
	st.Slots.Power = &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: cedhBracket}}
	st.Close("power")
	a.log.Info("cEDH names bracket 5, so the power slot is closed",
		"session", st.SessionID)
}

// ruleHouseRules stores the house rules when the user repeats the option
// the house-rules row offers. The slot is typed, so the D-119 net can not
// close it by name, and the classifier can miss the answer. The option
// the user repeated is the user's own words for what "anything goes"
// means (D-265).
//
// The rule fires when the row's question is out, or when the same
// message raises the row and answers it (D-122).
func ruleHouseRules(a *Agent, st *State, in turnWords) {
	if st.Slots.GetHouseRules() != "" {
		return
	}
	out := st.Slots.GetSlotStates()["house_rules"] == mtgv1.SlotState_SLOT_STATE_ASKED
	if !out && Route(in.Message) != "house_rules" {
		return
	}
	row, ok := a.cat.Row("house_rules")
	if !ok {
		return
	}
	msg := strings.ToLower(in.Message)
	for _, opt := range row.Options {
		if !optionAnswered(msg, opt) {
			continue
		}
		st.Slots.HouseRules = opt
		st.Close("house_rules")
		a.log.Info("the answer repeated a house-rules option, so the slot holds it",
			"session", st.SessionID, "house_rules", opt)
		return
	}
}

// ruleCardInThe99 settles the role of a named card the user put in the
// 99. It is not a commander, so the role question is already answered
// (D-70).
func ruleCardInThe99(_ *Agent, st *State, in turnWords) {
	if name := lockedByWords(st, in.Message); name != "" {
		st.AddLocked(name)
	}
}

// ruleNamedCardAsCommander sets the commander from the named card when
// the user gives it that role. The role row offers "As my commander",
// and the key is typed: its value is the card (D-118, D-83).
func ruleNamedCardAsCommander(a *Agent, st *State, in turnWords) {
	if st.Ctx.CommanderSet || len(st.NamedCards) == 0 {
		return
	}
	if st.Slots.GetSlotStates()["named_card_role"] != mtgv1.SlotState_SLOT_STATE_ASKED {
		return
	}
	msg := strings.ToLower(in.Message)
	if !optionAnswered(msg, "As my commander") && !namedCardAsCommander(in.Message) {
		return
	}
	a.log.Info("the user gave the named card the commander role",
		"session", st.SessionID, "card", st.NamedCards[0])
	st.SetCommander(st.NamedCards[0])
}

// ruleSwapCommander reopens the choice when the user asks for a commander
// other than the one they chose. Every commander row is closed by then,
// so nothing could ask (D-130).
func ruleSwapCommander(a *Agent, st *State, in turnWords) {
	if !st.Ctx.CommanderSet || !swapsCommander(in.Message) {
		return
	}
	a.log.Info("the user asked for another commander, so the choice reopens",
		"session", st.SessionID)
	st.ClearCommander()
}

// ruleCommanderPair reads a request for two commanders. A pair carries
// the union of two color identities, and it is the only practical way to
// reach four colors (D-154).
func ruleCommanderPair(a *Agent, st *State, in turnWords) {
	if st.Ctx.WantPair || !wantsCommanderPair(in.Message) {
		return
	}
	a.log.Info("the user asked for a two-commander pair", "session", st.SessionID)
	st.Ctx.WantPair = true
	st.Ctx.WantBackground = wantsBackgroundPair(in.Message)
	// The names on the table were single commanders, so they answer a
	// different question now.
	st.RetireOffer()
	st.Ctx.Suggested = true
}

// ruleDelegateCommander reads a commander choice the user handed to the
// agent. "You pick the commander" is neither a refusal nor a pick, and
// the pick row would otherwise ask again every turn (D-147).
//
// A delegation is a decline (D-93). It closes the key, it names no
// value, and the generator takes the best commander of the pool. The
// names on the table leave with it, so a later color change has nothing
// to reopen (D-153).
//
// The scope guard keeps "up to you" from closing the commander choice
// when the user answered some other question with it. The classifier
// says which question the message answered: a delegation it filed under
// the colors is not about the commander, even while a commander question
// is out. Without a classifier verdict, the commander question must be
// the only one out, or the message must name a commander (D-147).
//
// A superlative delegates too. "Buy the best lifegain commander" names
// no card and asks the agent to select one. That form carries its own
// guard: the message must name a commander (D-167).
func ruleDelegateCommander(a *Agent, st *State, in turnWords) {
	if st.Ctx.CommanderSet || (!delegatesChoice(in.Message) && !delegatesCommander(in.Message)) {
		return
	}
	if !delegationIsAboutTheCommander(st, in) {
		return
	}
	a.log.Info("the user handed the commander choice to the agent",
		"session", st.SessionID)
	st.Skip("commander_pick")
	st.Skip("commander")
	st.CurrentOffer = nil
}

// ruleDelegateColors closes the color slot on a delegation that is not
// about the commander (D-388). "Surprise me" hands back every open key,
// and the classifier reads the commander half of that and misses the
// colors.
//
// A delegation about the commander closes the colors on its own, because
// the commander's identity is the deck's identity (D-70).
//
// "Whatever is winning" is not a delegation here. The corpus routes
// "whatever" nowhere, after gate runs 11 to 13 read it as house rules
// six times (D-111).
func ruleDelegateColors(a *Agent, st *State, in turnWords) {
	if st.Ctx.Filled["colors"] || len(st.Slots.GetColors()) > 0 {
		return
	}
	if !delegatesChoice(in.Message) {
		return
	}
	// A delegation the classifier assigned to a commander key answers
	// that key alone. "You pick" against the pick row chooses a
	// commander, and it says nothing about the colors.
	//
	// The word test of delegationIsAboutTheCommander is too wide here.
	// It reads "a Commander deck, surprise me" as a commander answer,
	// and that message delegates every open key (D-93).
	if in.hasKey("commander") || in.hasKey("commander_pick") {
		return
	}
	a.log.Info("the user handed the color choice to the agent",
		"session", st.SessionID)
	st.Skip("colors")
}

// ruleFormatFromNamedLeader reads the format from a card that can only
// lead a Commander deck (D-388).
//
// A legendary creature names no format on its own: many of them play in
// Standard and Modern too. A card that can lead and is legal in neither
// leaves one format this app builds, so the format row must not offer
// three. Eval run 31 flagged that question on Atraxa, Praetor's Voice.
func ruleFormatFromNamedLeader(a *Agent, st *State, in turnWords) {
	if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return
	}
	// A format this app does not build is declined by its own row, and a
	// card must never override that (D-112).
	if _, _, bad := unsupportedFormat(in.Message); bad {
		return
	}
	fc, ok := a.hints.(FormatChecker)
	if !ok {
		return
	}
	for _, name := range st.NamedCards {
		only, known := fc.OnlyCommander(name)
		if !known || !only {
			continue
		}
		a.log.Info("the user named a card that can only lead a Commander deck",
			"session", st.SessionID, "card", name)
		a.setFormat(st, mtgv1.FormatId_FORMAT_ID_COMMANDER)
		return
	}
}

// delegationIsAboutTheCommander reports whether a delegation in the
// message answers a commander question.
func delegationIsAboutTheCommander(st *State, in turnWords) bool {
	if namesCommander(in.Message) {
		return true
	}
	if in.hasKey("commander") || in.hasKey("commander_pick") {
		return true
	}
	if in.assigned() {
		// The classifier filed the answer under other keys.
		return false
	}
	_, pickOut := st.Ctx.Outstanding["commander_pick"]
	_, baseOut := st.Ctx.Outstanding["commander"]
	if !pickOut && !baseOut {
		return false
	}
	for key := range st.Ctx.Outstanding {
		if key != "commander" && key != "commander_pick" {
			return false
		}
	}
	return true
}

// rulePickByPlace reads a commander chosen by its place, such as "the
// first of the new three". The classifier can not map that onto a name,
// because it never sees the names (D-121).
func rulePickByPlace(a *Agent, st *State, in turnWords) {
	if !st.Ctx.Asked["commander_pick"] || st.Ctx.Filled["commander_pick"] || refusedOffer(in.Message) {
		return
	}
	if i, ok := offeredPick(in.Message); ok && i < len(st.CurrentOffer) {
		a.log.Info("the user chose a commander by its place",
			"session", st.SessionID, "place", i+1, "commander", st.CurrentOffer[i])
		st.SetCommander(st.CurrentOffer[i])
	}
}

// ruleRefuseOffer retires the names on the table when the user refuses
// them, and the pick row asks again with three others. The classifier
// can report the refusal as an answer, so a word rule reads it (D-73,
// D-120).
func ruleRefuseOffer(a *Agent, st *State, in turnWords) {
	if !st.Ctx.Asked["commander_pick"] || st.Ctx.Filled["commander_pick"] || !refusedOffer(in.Message) {
		return
	}
	st.RetireOffer()
	st.Ctx.Suggested = true
	a.log.Info("the user refused the commanders on the table, so three others follow",
		"session", st.SessionID)
}

// ruleInferPower infers the tournament step from a competitive request.
// The slot ends with a value whatever the user answers, which D-90
// requires. Commander keeps its bracket question, because bracket 4 and
// bracket 5 are too far apart to infer (D-107).
//
// A step the user named reaches the slot through the classify call, so
// no inference runs (D-209). The agent states the step and asks nothing:
// a user who asked for the strongest deck has given the answer (D-216).
func ruleInferPower(a *Agent, st *State, _ turnWords) {
	if !st.Ctx.PowerCompetitive || !sixtyCard(st.Ctx.Format) || st.Slots.GetPower() != nil {
		return
	}
	st.Slots.Power = &mtgv1.PowerLevel{
		Level: &mtgv1.PowerLevel_SixtyStep{SixtyStep: mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT},
	}
	st.Close("power")
	a.log.Info("the agent inferred the tournament step from a competitive request",
		"session", st.SessionID)
}

// lockedByWords names the card the user placed in the 99, or an empty
// string. It reads the current message first. It falls back to the whole
// conversation only when one card was ever named, so a later card can not
// take the place of the one the user meant.
func lockedByWords(st *State, message string) string {
	if len(st.NamedCards) == 0 {
		return ""
	}
	if namedCardNotCommander(message) {
		return st.NamedCards[0]
	}
	if len(st.NamedCards) == 1 && namedCardNotCommander(st.Ctx.Words) {
		return st.NamedCards[0]
	}
	return ""
}

// apply writes one classify result onto the state. It never clears a slot
// the session already filled: a value stays until the user replaces it.
func (a *Agent) apply(ctx context.Context, st *State, out classifyOut, open []string, message string, acc *llm.Accumulator) {
	a.applyFormat(st, out, message)
	a.applyTheme(st, out)
	// A colorless request names no color, and the classifier can answer
	// the open color question with all five. The colorless rule closes
	// the slot instead (D-165, D-535).
	if !colorlessRequest(message) {
		a.applyColors(st, out, message)
	}
	a.applyNames(st, out)
	a.applySets(ctx, st, out, acc)
	a.applyPrecons(st, out)
	if rule, ok := poolRules[slotWord(out.PoolRule)]; ok {
		// The reader's choice on the chat screen wins, always (D-591).
		// A reader who picked a collection and "Only cards I own" and
		// then named a set read "Pool: any card", because the classifier
		// took the set phrase for a pool answer and wrote over the
		// choice. The flip also turned the buy list on, so the deck
		// named cards the reader does not own. It is the D-371 collision
		// from the other side.
		if st.Ctx.PoolFromReader {
			a.log.Info("the reader chose the card pool, so the classifier does not write over it",
				"session", st.SessionID, "reader", st.Slots.GetPoolRule(), "classifier", rule)
		} else {
			// An owned rule needs a collection. A reader with none who
			// says "build only from the Hobbit set" names a set, not
			// their library, and the classifier reads the word "only" as
			// ownership. The rule then empties the card pool and the
			// commander pool, and the deck can not be built at all
			// (D-371).
			if rule != mtgv1.PoolRule_POOL_RULE_ANY_CARD && !st.Ctx.HasCollection {
				a.log.Info("an owned pool rule needs a collection, and this session has none",
					"session", st.SessionID, "rule", rule)
				rule = mtgv1.PoolRule_POOL_RULE_ANY_CARD
			}
			st.Slots.PoolRule = rule
			st.Close("pool_rule")
		}
	}
	a.applyPower(st, out, message)
	// A budget applies only when the message names it (D-537, the D-125
	// rule for the budget). The classifier answered the open budget
	// question with 2 on "the best deck under budget", and the session
	// called itself complete before the user named 400.
	switch {
	case out.BudgetUSD <= 0:
	case !budgetNamed(message, out.BudgetUSD):
		a.log.Info("the classifier reported a budget the message does not name, so the budget slot keeps its state",
			"session", st.SessionID, "reported", out.BudgetUSD)
	default:
		st.Slots.BudgetUsd = out.BudgetUSD
		st.Close("budget")
	}
	// The scope answers its own row, so a user who says "on the whole
	// deck" closes it without naming a number again (D-238).
	if sc := budgetScope(out.BudgetScope); sc != mtgv1.BudgetScope_BUDGET_SCOPE_UNSPECIFIED {
		st.Slots.BudgetScope = sc
		st.Close("budget_scope")
	}
	// The house rules close on the user's words, so the build can copy
	// them to Format.house_rules (D-265).
	if s := strings.TrimSpace(out.HouseRules); s != "" {
		st.Slots.HouseRules = s
		st.Close("house_rules")
	}
	a.applyKeys(st, out, open, message)
	a.applyFacts(st, out, message)
}

// applyTheme writes the theme. A superlative phrase such as "the best
// deck under budget" is a theme for a deck with none, as the prompt says,
// and it must not replace a theme the user named: "Modern. The best deck
// under budget." answers the budget question of an infect deck, and
// infect stays (D-535, the D-125 rule for the theme).
func (a *Agent) applyTheme(st *State, out classifyOut) {
	s := strings.TrimSpace(out.Theme)
	if s == "" {
		return
	}
	// An occasion is not a theme. "A Modern deck for a team event" names a
	// happening, and a theme written from it stops the theme row, so the
	// reader never names the plan (D-670, extends D-219).
	if occasionTheme(s) {
		a.log.Info("an occasion is not a theme, so the classifier theme is dropped",
			"session", st.SessionID, "theme", s)
		return
	}
	if st.Slots.GetTheme() != "" && anyPhrase(strings.ToLower(s), bestSigns) {
		a.log.Info("a superlative phrase does not replace the theme the user named",
			"session", st.SessionID, "theme", st.Slots.GetTheme(), "phrase", s)
		return
	}
	st.Slots.Theme, st.Ctx.Theme = s, strings.ToLower(s)
	st.Close("theme")
}

// applyFormat writes the format the classifier reported. A filled format
// changes only when the message names the new one, because the model can
// repeat a value the user has just replaced (D-125).
//
// A message that names a format this app does not build gives no format
// at all. "I play Duel Commander" holds the word "commander", and the
// classifier can report Commander for it. The unsupported-format row
// must decline it instead (D-199, D-112).
func (a *Agent) applyFormat(st *State, out classifyOut, message string) {
	id, ok := formatIDs[slotWord(out.Format)]
	if !ok || id == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return
	}
	_, _, unsupportedNow := unsupportedFormat(message)
	switch {
	case unsupportedNow:
		a.log.Warn("the classifier reported a format on a message that names one this app does not build",
			"session", st.SessionID, "reported", id.String())
	case st.Ctx.Format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED, namesFormat(message, id):
		// A changed format retires every question that is still out.
		// The bracket question means nothing in Modern, and D-126
		// would otherwise block the 60-card power question behind it.
		// A retired row may ask again (D-195).
		if st.Ctx.Format != mtgv1.FormatId_FORMAT_ID_UNSPECIFIED && st.Ctx.Format != id {
			a.log.Info("the user changed the format, so the open questions retire",
				"session", st.SessionID, "from", st.Ctx.Format.String(), "to", id.String())
			st.RetireOutstanding(a.cat)
		}
		a.setFormat(st, id)
	case id != st.Ctx.Format:
		a.log.Warn("the classifier reported a format this message does not name",
			"session", st.SessionID, "reported", id.String(), "kept", st.Ctx.Format.String())
	}
}

// applyColors writes the colors. The old value stays until at least one
// new color is valid. A list of unknown words wiped the colors the user
// gave before.
//
// A closed color slot changes only when the message names a color
// (D-535, the D-125 rule for colors). The classifier repeats values, and
// it answered "Bracket 3." with all five colors for a deck the user had
// called colorless, over the slot the colorless rule had closed.
func (a *Agent) applyColors(st *State, out classifyOut, message string) {
	if len(out.Colors) == 0 {
		return
	}
	if st.Ctx.Filled["colors"] && !namesAColor(message) {
		a.log.Info("the classifier reported colors on a message that names none, so the closed color slot keeps its value",
			"session", st.SessionID, "reported", out.Colors)
		return
	}
	var colors []mtgv1.Color
	for _, c := range out.Colors {
		if id, ok := colorIDs[strings.ToUpper(c)]; ok {
			colors = append(colors, id)
		}
	}
	if len(colors) == 0 {
		return
	}
	st.Slots.Colors = colors
	st.Close("colors")
	// The names on the table were chosen before these colors arrived,
	// so they are checked against the colors again (D-153).
	a.dropOffColorOffers(st)
}

// applySets resolves the set names the reader wrote onto a set family
// and writes it to the slots (D-376).
//
// A phrase that names one base set closes the set row. A phrase this app
// can not settle opens it, and the row asks. A message that names no set
// changes nothing: a set the reader gave before stays until they replace
// it, which is the rule every other slot follows.
// setMatchOut is what the set matcher answers (D-581, F-64).
type setMatchOut struct {
	Codes      []string `json:"codes"`
	Candidates []string `json:"candidates"`
}

// maxSetOptions bounds the sets the question offers after a match. The
// owner asked for the three most likely (D-581).
const maxSetOptions = 3

// matchSets asks the model what set a phrase names, when the set table
// settles nothing (D-581, F-64). It answers the family codes and their
// names for a definitive match, or the options the question offers.
//
// Every code the model answers must be a base set of this snapshot. A
// code the snapshot does not hold is dropped, so an invented set reaches
// no deck.
func (a *Agent) matchSets(ctx context.Context, phrase string, acc *llm.Accumulator) (codes, names, options []string) {
	src, ok := a.hints.(SetMatchSource)
	if !ok || a.llm == nil {
		return nil, nil, nil
	}
	rows := src.SetRows()
	if len(rows) == 0 {
		return nil, nil, nil
	}
	sets := make([]map[string]string, 0, len(rows))
	for _, r := range rows {
		sets = append(sets, map[string]string{"code": r.Code, "name": r.Name, "released": r.Released})
	}
	input, err := json.Marshal(map[string]any{"phrase": phrase, "sets": sets})
	if err != nil {
		return nil, nil, nil
	}
	res, err := a.llm.Complete(ctx, llm.RoleSetMatch, llm.Request{
		Instructions: setMatchInstructions,
		Input:        string(input),
		SchemaName:   "set_match",
		Schema:       json.RawMessage(setMatchSchema),
		// The set list is the same for every caller of a snapshot, so one
		// cache key serves them all.
		CacheKey: "setmatch",
	}, acc)
	if err != nil {
		a.log.Warn("the set matcher failed, so the set row asks", "phrase", phrase, "err", err)
		return nil, nil, nil
	}
	var got setMatchOut
	if err := json.Unmarshal(res.Output, &got); err != nil {
		a.log.Warn("the set matcher answered no object, so the set row asks", "phrase", phrase, "err", err)
		return nil, nil, nil
	}
	seen := map[string]bool{}
	for _, code := range got.Codes {
		famCodes, famNames, valid := src.SetFamily(strings.ToLower(strings.TrimSpace(code)))
		if !valid {
			a.log.Warn("the set matcher named a set the snapshot does not hold", "phrase", phrase, "code", code)
			continue
		}
		for i, c := range famCodes {
			if seen[c] {
				continue
			}
			seen[c] = true
			codes = append(codes, c)
			names = append(names, famNames[i])
		}
	}
	if len(codes) > 0 {
		return codes, names, nil
	}
	for _, code := range got.Candidates {
		_, famNames, valid := src.SetFamily(strings.ToLower(strings.TrimSpace(code)))
		if !valid || len(famNames) == 0 {
			continue
		}
		options = append(options, famNames[0])
		if len(options) == maxSetOptions {
			break
		}
	}
	return nil, nil, options
}

func (a *Agent) applySets(ctx context.Context, st *State, out classifyOut, acc *llm.Accumulator) {
	if len(out.SetNames) == 0 && len(out.SetGroups) == 0 {
		return
	}
	r, ok := a.hints.(SetResolver)
	if !ok {
		return
	}
	phrases := append(append([]string(nil), out.SetNames...), out.SetGroups...)
	var codes, names []string
	var unresolved string
	var options []string
	for _, phrase := range out.SetNames {
		if strings.TrimSpace(phrase) == "" {
			continue
		}
		gotCodes, gotNames, gotOptions, done := r.ResolveSet(phrase)
		if !done {
			// The table settles a name and a code, and it settles no
			// abbreviation. The model reads the phrase against the set
			// list before the row asks (D-581, F-64).
			matchCodes, matchNames, matchOptions := a.matchSets(ctx, phrase, acc)
			if len(matchCodes) > 0 {
				codes = append(codes, matchCodes...)
				names = append(names, matchNames...)
				continue
			}
			// The first phrase this app can not settle is the one the row
			// asks about. A second one waits for its turn.
			if unresolved == "" {
				unresolved = phrase
				options = gotOptions
				if len(options) == 0 {
					options = matchOptions
				}
			}
			continue
		}
		codes = append(codes, gotCodes...)
		names = append(names, gotNames...)
	}
	// A group reaches every family of a franchise (D-525). A group the
	// snapshot does not know asks the set row, as an unknown name does,
	// and the row offers no option for it.
	for _, phrase := range out.SetGroups {
		if strings.TrimSpace(phrase) == "" {
			continue
		}
		gotCodes, gotNames, done := r.ResolveSetGroup(phrase)
		if !done {
			// The group table matches a set name word by word, and it
			// settles no abbreviation. The model reads the phrase as it
			// does for a name (D-581, F-74). A reader who wrote "LOTR"
			// answered the set question already, and the row asked it
			// again with no option to pick.
			matchCodes, matchNames, matchOptions := a.matchSets(ctx, phrase, acc)
			if len(matchCodes) > 0 {
				codes = append(codes, matchCodes...)
				names = append(names, matchNames...)
				continue
			}
			if unresolved == "" {
				unresolved, options = phrase, matchOptions
			}
			continue
		}
		codes = append(codes, gotCodes...)
		names = append(names, gotNames...)
	}
	// A message can name two sets and resolve one of them. The resolved
	// set fills the slot, and the row still asks about the other. Both
	// halves run, so neither answer is dropped in silence (D-376).
	if len(codes) > 0 {
		codes = append(codes, st.Slots.GetSetCodes()...)
		names = append(names, st.SetNames...)
		codes, names = dedupeSets(codes, names)
		a.log.Info("the deck is limited to the sets the reader named",
			"session", st.SessionID, "phrase", strings.Join(phrases, ", "),
			"sets", strings.Join(codes, ","))
		st.SetLimit(strings.Join(phrases, ", "), codes, names)
		// The reader hears which sets the words became. "The Hobbit"
		// is two sets, and a red mark on a card explains nothing until
		// the reader knows what the limit is (D-390).
		st.setsThisTurn = names
	}
	if unresolved != "" {
		a.log.Info("the reader named a set this app can not settle",
			"session", st.SessionID, "phrase", unresolved, "options", len(options))
		st.SetUnresolved(unresolved, options)
		return
	}
	// Every phrase of this message named a set, so the row that asks
	// which set a name means has its answer.
	if st.Ctx.SetUnresolved || st.UnresolvedSet != "" {
		st.SetResolved()
	}
}

// dedupeSets drops a repeated code and keeps the names beside it. Two
// phrases can name one family, and one set must not be listed twice.
func dedupeSets(codes, names []string) ([]string, []string) {
	byCode := map[string]string{}
	for i, c := range codes {
		if _, seen := byCode[c]; seen {
			continue
		}
		name := c
		if i < len(names) && strings.TrimSpace(names[i]) != "" {
			name = names[i]
		}
		byCode[c] = name
	}
	outCodes := make([]string, 0, len(byCode))
	for c := range byCode {
		outCodes = append(outCodes, c)
	}
	// Sorted, so two runs of one session read the same and a log line
	// compares with the next one.
	sort.Strings(outCodes)
	outNames := make([]string, 0, len(outCodes))
	for _, c := range outCodes {
		outNames = append(outNames, byCode[c])
	}
	return outCodes, outNames
}

// applyNames writes the three card lists. The lists are kept apart.
// Merged, one name that becomes the commander also reads as a card to
// keep (D-70).
func (a *Agent) applyNames(st *State, out classifyOut) {
	named := append([]string(nil), out.CommanderNames...)
	named = append(named, out.LockedNames...)
	named = append(named, out.NamedCards...)
	for _, name := range named {
		if name = strings.TrimSpace(name); name != "" {
			// mergeFront keeps one entry per card. The classifier reports
			// the full name in one turn and the short name in the next
			// (D-70).
			st.NamedCards = mergeFront(st.NamedCards, name)
			st.Ctx.NamedCard = true
		}
	}
	// A card named with no role after the commander is set gets the role
	// question. SetCommander closed the role key, and a card the user
	// names later must not sit in the 99 in silence (D-118).
	if st.Ctx.CommanderSet {
		for _, name := range out.NamedCards {
			if name = strings.TrimSpace(name); name != "" && !hasName(st.CommanderNames, name) && !hasName(st.LockedNames, name) {
				a.log.Info("a card was named after the commander was set, so the role question reopens",
					"session", st.SessionID, "card", name)
				st.ReopenRows(a.cat, "named_card_role")
				break
			}
		}
	}
	for _, name := range out.LockedNames {
		st.AddLocked(name)
	}
	ck, checks := a.hints.(CommanderChecker)
	// A named card that can not lead a deck has a settled role: it can
	// only sit in the 99. The role question would ask what the card
	// itself answers (D-220, extends D-129 and D-70).
	if checks {
		for _, name := range out.NamedCards {
			if name = strings.TrimSpace(name); name == "" {
				continue
			}
			if canLead, known := ck.CanLead(name); known && !canLead {
				st.AddLocked(name)
				a.log.Info("a named card can not lead a deck, so its role is settled",
					"session", st.SessionID, "card", name)
			}
		}
	}
	for _, name := range out.CommanderNames {
		// A card that can not lead a deck is not a commander (D-129).
		if checks {
			if canLead, known := ck.CanLead(name); known && !canLead {
				st.IllegalCommander = strings.TrimSpace(name)
				st.Ctx.CommanderIllegal = true
				// A legal commander closed the illegal row (D-196). A new
				// illegal name raises it again.
				st.ReopenRows(a.cat, "commander_illegal")
				// The user still wants the card in the deck, and its role
				// is settled: it can only sit in the 99. AddLocked closes
				// the role question with it (D-70).
				st.AddLocked(name)
				a.log.Info("the named commander can not lead a deck",
					"session", st.SessionID, "card", name)
				continue
			}
			_, known := ck.CanLead(name)
			switch {
			case known:
				a.dropUnknownCommanders(st, ck, name)
			case knownCommander(st, ck):
				// An unknown name never joins a known commander (D-538).
				// The classifier repeats values, so "I meant Atraxa,
				// Praetors' Voice" can come back with both spellings in
				// one turn, and the misspelled one must not lead the deck
				// beside the right one.
				a.log.Info("a commander name the index does not know does not join a known one",
					"session", st.SessionID, "dropped", name)
				st.NamedCards = withoutName(st.NamedCards, name)
				continue
			case a.askWhichCommander(st, name):
				// The name is not the commander until the reader picks a
				// card (F-75, D-606).
				continue
			}
		}
		st.SetCommander(name)
	}
	// A commander the index knows answers the row that asks which card a
	// name means, whatever words brought it (F-75).
	if st.Ctx.CommanderSet && st.Ctx.CommanderUnresolved {
		st.DropUnresolvedCommander()
	}
}

// askWhichCommander records a commander name the card index does not
// hold, so the row asks which card the reader means (F-75, D-606). It
// reports whether the row took the name.
//
// The reader wrote "Aragorn", and no card carries that name alone.
// applyNames set the name, build.go missed it in the index, and the
// build took the delegation path of D-232 with no word to the reader.
//
// The name leaves the named cards with it. It names no card, so the role
// row would ask about a card that does not exist.
func (a *Agent) askWhichCommander(st *State, name string) bool {
	r, ok := a.hints.(CommanderNameResolver)
	if !ok {
		return false
	}
	options := r.ResolveCommander(name)
	a.log.Info("the card index holds no commander of this name, so the row asks which card",
		"session", st.SessionID, "name", name, "options", len(options))
	st.NamedCards = withoutName(st.NamedCards, name)
	st.Ctx.NamedCard = len(st.NamedCards) > 0
	st.CommanderUnresolved(name, options)
	return true
}

// knownCommander reports whether a commander name the index knows is set.
func knownCommander(st *State, ck CommanderChecker) bool {
	for _, n := range st.CommanderNames {
		if _, known := ck.CanLead(n); known {
			return true
		}
	}
	return false
}

// dropUnknownCommanders removes the commander names the index does not
// know when a known name arrives, so "I meant Atraxa, Praetors' Voice"
// replaces the misspelled name instead of joining it in the command zone
// (D-535). The misspelled name leaves the named cards too. keep is the
// name that arrived.
func (a *Agent) dropUnknownCommanders(st *State, ck CommanderChecker, keep string) {
	var kept []string
	for _, old := range st.CommanderNames {
		if _, known := ck.CanLead(old); known || sameCard(old, keep) {
			kept = append(kept, old)
			continue
		}
		a.log.Info("a commander name the index does not know gives way to a known one",
			"session", st.SessionID, "dropped", old, "kept", keep)
		st.NamedCards = withoutName(st.NamedCards, old)
	}
	st.CommanderNames = kept
}

// withoutName is the list without the named card.
func withoutName(list []string, name string) []string {
	var out []string
	for _, n := range list {
		if !sameCard(n, name) {
			out = append(out, n)
		}
	}
	return out
}

// applyPower writes the power level. An occasion is not a power level:
// "a Modern deck for an event" names no step, and the classifier can
// answer the tournament step for it. The open power row asks instead
// (D-219, extends D-215).
func (a *Agent) applyPower(st *State, out classifyOut, message string) {
	p := power(out.Power)
	if p == nil {
		return
	}
	if p.GetSixtyStep() != mtgv1.SixtyStep_SIXTY_STEP_UNSPECIFIED && occasionOnly(message) {
		a.log.Info("an occasion names no power step, so the classifier step is dropped",
			"session", st.SessionID)
		return
	}
	st.Slots.Power = p
	st.Close("power")
	// The corpus triggers the competitive theme row on the power
	// value: FNM or tournament-meta. A user who names the step
	// outright must reach it too (D-132).
	switch p.GetSixtyStep() {
	case mtgv1.SixtyStep_SIXTY_STEP_FNM, mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT:
		st.Ctx.PowerCompetitive = true
	}
}

// applyKeys closes and declines keys by name.
//
// A key closes by name only when two things hold: its question is out,
// and it carries no typed value.
//
// The question must be out, because the classifier can retire a key
// nobody asked, and the session then calls itself complete (D-76).
//
// A typed slot must close on its value, because a name says only
// "answered" and never says what the answer was. A format closed by
// name stays empty, and every row that triggers on the format stops.
// Asking the format twice is the safe failure. Building a deck with no
// format is not (D-83).
func (a *Agent) applyKeys(st *State, out classifyOut, open []string, message string) {
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
	// A delegation that names its slot hands over that slot alone. "You pick
	// the commander" declines the commander, and the classifier can decline
	// every open key for it, so the bracket and the budget take defaults the
	// reader never chose (D-670).
	delegated, bare, delegates := delegationObjects(message)
	for _, k := range out.DeclinedKeys {
		k = strings.TrimSpace(k)
		if k == "" {
			continue
		}
		if delegates && !bare && len(delegated) > 0 && !delegated[k] && slotHasNouns(k) {
			a.log.Warn("a delegation that names its slot declines no other key", "key", k)
			continue
		}
		// One shape must not reach it: "None of those" refuses the names
		// on the table and asks for others. It is not a decline, and the
		// classifier can report it as one (D-120).
		if k == "commander_pick" && refusedOffer(message) {
			a.log.Warn("a refusal of the offered names is not a decline", "key", k)
			continue
		}
		if st.Slots.GetSlotStates()[k] != mtgv1.SlotState_SLOT_STATE_ASKED {
			a.log.Warn("classifier declined a key whose question is not out", "key", k, "offered", open)
			continue
		}
		a.log.Info("the user declined a slot", "key", k)
		formatWasUnset := st.Ctx.Format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED
		st.DeclineKey(k)
		if k == "format" && formatWasUnset {
			a.log.Info("a declined format took the corpus default", "format", DefaultFormat.String())
		}
	}
}

// applyPrecons reads the precons the deck must use no card of (D-496):
// the products the message named, and every precon the collection holds
// whole when the reader says "not from my precons" (D-408). A named
// product the collection does not hold whole is excluded anyway, and the
// turn says so (D-497). A phrase the table can not settle opens the
// precon row, as an unknown set name opens the set row.
func (a *Agent) applyPrecons(st *State, out classifyOut) {
	if len(out.PreconNames) == 0 && !out.Facts.ExcludePrecons {
		return
	}
	var refs []PreconRef
	var partial []string
	var unresolved string
	var options []string
	r, hasResolver := a.hints.(PreconResolver)
	for _, phrase := range out.PreconNames {
		if strings.TrimSpace(phrase) == "" {
			continue
		}
		if !hasResolver || !hasTable(r) {
			st.preconsUnavailableThisTurn = true
			continue
		}
		m := r.ResolvePrecon(phrase)
		if !m.OK {
			// The first phrase this app can not settle is the one the row
			// asks about. A second one waits for its turn.
			if unresolved == "" {
				unresolved, options = phrase, m.Options
			}
			continue
		}
		refs = append(refs, m.Products...)
		partial = append(partial, m.Partial...)
	}
	none := false
	if out.Facts.ExcludePrecons {
		if o, ok := a.hints.(OwnedPreconSource); ok {
			owned, ok := o.OwnedPrecons()
			switch {
			case !ok:
				st.preconsUnavailableThisTurn = true
			case len(owned) == 0:
				none = true
			default:
				refs = append(refs, owned...)
			}
		} else {
			st.preconsUnavailableThisTurn = true
		}
	}
	if len(refs) > 0 || none {
		for i, key := range st.Slots.GetExcludePreconKeys() {
			name := key
			if i < len(st.ExcludedPreconNames) {
				name = st.ExcludedPreconNames[i]
			}
			refs = append(refs, PreconRef{Key: key, Name: name})
		}
		refs = dedupePrecons(refs)
		phrase := strings.Join(out.PreconNames, ", ")
		if out.Facts.ExcludePrecons {
			phrase = strings.TrimPrefix(phrase+", my precons", ", ")
		}
		a.log.Info("the deck uses no card of the reader's precons",
			"session", st.SessionID, "phrase", phrase, "products", len(refs))
		st.ExcludePrecons(phrase, refs)
		st.preconsThisTurn = st.ExcludedPreconNames
		st.preconsPartialThisTurn = partial
		st.preconsNoneThisTurn = len(refs) == 0
	}
	if unresolved != "" {
		a.log.Info("the reader named a precon this app can not settle",
			"session", st.SessionID, "phrase", unresolved, "options", len(options))
		st.PreconUnresolved(unresolved, options)
		return
	}
	if st.Ctx.PreconUnresolved || st.UnresolvedPrecon != "" {
		st.PreconResolved()
	}
}

// hasTable reports whether a resolver holds a precon table at all. A
// resolver with none answers every phrase with no product and no option.
func hasTable(r PreconResolver) bool {
	if h, ok := r.(*CandidateHints); ok {
		return h != nil && h.Precons != nil
	}
	return true
}

// dedupePrecons keeps the first of each key, in order.
func dedupePrecons(refs []PreconRef) []PreconRef {
	seen := map[string]bool{}
	out := make([]PreconRef, 0, len(refs))
	for _, r := range refs {
		if r.Key == "" || seen[r.Key] {
			continue
		}
		seen[r.Key] = true
		out = append(out, r)
	}
	return out
}

// fillPrecons copies the precon marks of this turn onto the result, and
// clears them: they are turn state, as setsThisTurn is.
func (s *State) fillPrecons(res *Result) {
	res.PreconsApplied = s.preconsThisTurn
	res.PreconsPartial = s.preconsPartialThisTurn
	res.PreconsNone = s.preconsNoneThisTurn
	res.PreconsUnavailable = s.preconsUnavailableThisTurn
	s.preconsThisTurn, s.preconsPartialThisTurn = nil, nil
	s.preconsNoneThisTurn, s.preconsUnavailableThisTurn = false, false
}

// applyFacts writes the classifier facts onto the context.
func (a *Agent) applyFacts(st *State, out classifyOut, message string) {
	f := out.Facts
	st.Ctx.NamedCard = st.Ctx.NamedCard || f.NamedCard
	st.Ctx.BuyList = st.Ctx.BuyList || f.BuyList
	st.Ctx.HouseFormat = st.Ctx.HouseFormat || f.HouseFormat
	st.Ctx.BudgetAmbiguous = st.Ctx.BudgetAmbiguous || f.BudgetAmbiguous
	// The classifier can read an occasion as a power level. The
	// user's own words carry the fact, as they carry the format, and
	// the fact is sticky, so one read per message is enough (D-215,
	// extends D-199).
	st.Ctx.PowerCompetitive = st.Ctx.PowerCompetitive || competitiveRequest(message)
	// The fact is not sticky. A user who asks for a Magic deck after the
	// agent declines is back in scope. A user who asks for another game
	// again hears the sentence again, so the key reopens (D-99).
	st.Ctx.OutOfScope = f.OutOfScope
	if f.OutOfScope && st.Ctx.Filled["scope"] {
		a.log.Info("the user asked for another game again, so the scope question reopens",
			"session", st.SessionID)
		st.ReopenRows(a.cat, "scope")
	}
	// A user who asks for a suggestion gets the pick row next turn. A user
	// who then names one closes every commander row, so the fact does not
	// need to be cleared.
	// A request for a suggestion does not retire the names on the
	// table. The same three stay until the user asks for others, and
	// a refusal is what asks (D-80, D-120, D-123).
	st.Ctx.Suggested = st.Ctx.Suggested || f.WantsSuggestion
	// A word rule can fire when the classifier misses one (corpus 11).
	// anyPhrase reads the negation, so "we have a ban list" fires nothing.
	if slot := Route(st.Ctx.Words); slot == "house_rules" {
		st.Ctx.HouseFormat = st.Ctx.HouseFormat || anyPhrase(st.Ctx.Words, []string{"no ban list"})
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
	// The classifier may answer with the word instead of the number.
	if s == "cedh" {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: cedhBracket}}
	}
	if n, err := strconv.Atoi(strings.TrimPrefix(s, "bracket ")); err == nil && n >= 1 && n <= 5 {
		return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: int32(n)}}
	}
	// The option of the power row leads with its number: "4 optimized"
	// (D-593). The prompt asks for "bracket 4", and the field is free
	// text, so a model that echoes the option the reader picked wrote a
	// string this function could not read. The slot then stayed asked,
	// the readiness gate waits on it, and the agent never repeats a
	// question, so session oUZMC0F2vHe7GGl24LIP could not go on.
	if head, _, ok := strings.Cut(s, " "); ok {
		if n, err := strconv.Atoi(head); err == nil && n >= 1 && n <= 5 {
			return &mtgv1.PowerLevel{Level: &mtgv1.PowerLevel_Bracket{Bracket: int32(n)}}
		}
	}
	return nil
}

// typedSlots are the keys that carry a value the deck generator reads.
// The classifier fills each one through its own field, so a name in the
// closed_keys list can never close one: the name says "answered" and
// leaves the value empty (D-83).
var typedSlots = map[string]bool{
	"format": true, "theme": true, "colors": true,
	"power": true, "pool_rule": true, "budget": true,
	// The scope of a cap is a typed value. An option match would close
	// it with the value UNSPECIFIED (D-238).
	"budget_scope": true,
	// The house rules close on the user's words, which the build copies
	// to Format.house_rules (D-265).
	"house_rules": true,
	// The commander closes on a name, and that name closes the color
	// slot and the two other commander rows with it.
	"commander": true,
	// The set closes on a resolved set family, and the mana row closes
	// on a permission. An option match would close the mana row on a
	// "no" as if it were a "yes" (D-376, D-382).
	SlotSet:            true,
	SlotSetOutsideMana: true,
	// The pick row closes on a name as well. Its answer is a commander,
	// and "none of those" is a refusal and not an answer. A pick closed
	// by name would carry a commander nobody chose (D-120).
	"commander_pick": true,
	// The role row closes on the card's role: "As my commander" sets the
	// commander, and "In the 99" locks the card. A name closes it with
	// neither (D-118).
	"named_card_role": true,
}

// deckKeys are the slots a deck request fills. A user who fills one has
// asked for a Magic deck, whatever else they asked for before (D-196).
var deckKeys = []string{"format", "theme", "colors", "commander"}

// deckKeysFilled counts the deck keys the session holds.
func deckKeysFilled(st *State) int {
	n := 0
	for _, k := range deckKeys {
		if st.Ctx.Filled[k] {
			n++
		}
	}
	return n
}

// clampFit keeps a fit score inside 0 to 1. The schema asks for that
// range, and the M-4 record must not carry a value outside it.
func clampFit(f float64) float64 {
	switch {
	case f < 0:
		return 0
	case f > 1:
		return 1
	}
	return f
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

// dropOffColorOffers removes an offered commander that the colors just
// named exclude. It keeps every name that still fits, so D-80 and D-123
// still hold: the agent swaps no name the user could still choose.
//
// A dropped name is retired, so the pool never offers it again. When
// nothing is left, the pick row asks with three new names, and D-127
// settles the case where the pool has none.
//
// A settled choice stays settled. A user who delegated the commander, or
// who named one, gets no pick question from a color change: the offer
// leaves, and the delegation stands (D-153).
func (a *Agent) dropOffColorOffers(st *State) {
	if st.Ctx.CommanderSet || st.Ctx.Filled["commander_pick"] {
		st.CurrentOffer = nil
		return
	}
	checker, ok := a.hints.(IdentityChecker)
	if !ok || len(st.CurrentOffer) == 0 {
		return
	}
	kept := make([]string, 0, len(st.CurrentOffer))
	var dropped []string
	for _, name := range st.CurrentOffer {
		fits, known := checker.FitsColors(name, st.Slots.Colors)
		// An unknown name proves nothing, so it stays on the table.
		if !known || fits {
			kept = append(kept, name)
			continue
		}
		dropped = append(dropped, name)
	}
	if len(dropped) == 0 {
		return
	}
	a.log.Info("an offered commander no longer fits the colors, so it leaves the table",
		"session", st.SessionID, "dropped", dropped, "kept", kept)
	for _, name := range dropped {
		if name = strings.TrimSpace(name); name != "" && !hasName(st.OfferedCommanders, name) {
			st.OfferedCommanders = append(st.OfferedCommanders, name)
		}
	}
	st.CurrentOffer = kept
	// The pick row must ask again with a full set of names.
	st.ReopenRows(a.cat, "commander_pick")
	st.Ctx.Suggested = true
}

// budgetScope reads the scope word the classifier reported (D-238).
func budgetScope(s string) mtgv1.BudgetScope {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "buy", "cards", "purchases":
		return mtgv1.BudgetScope_BUDGET_SCOPE_CARDS_TO_BUY
	case "deck", "whole", "whole_deck", "whole deck":
		return mtgv1.BudgetScope_BUDGET_SCOPE_WHOLE_DECK
	}
	return mtgv1.BudgetScope_BUDGET_SCOPE_UNSPECIFIED
}
