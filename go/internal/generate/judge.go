package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/llm"
	"github.com/nkramber/decktome/go/internal/rules"
)

// The judge lane is the real check for F-26. The deterministic net of
// LintSummary reads the shape of a rules claim and can not read its
// truth.
//
// The judge runs on another provider than the generator, so it never
// rates its own work (D-22). It is cheap enough to run beside every gate
// (D-229).

const judgeInstructions = `You check one paragraph from a Magic: The Gathering deck builder.

The paragraph is a deck summary written for the person who will play the deck. It must describe the deck and state no rule of the game. The rules engine reports the rules, and the summary must not.

Report two things.

First, every statement in the paragraph that asserts a rule of the game. A rule of the game is anything about what a card may do, what a format allows, what is banned or legal, what counts toward a limit, how many copies a deck may hold, or whether a card can lead a deck. A description of what the deck does on the table is not a rule.

Second, for each such statement, whether it is true. Judge it against the real rules of Magic: The Gathering. Say "unknown" when you can not tell.

Be strict about what counts as a rules claim and honest about truth. A summary with no rules claim is the expected result.

When the input holds the card list with the mana cost and the type line of each card from the card data, judge a claim about the cost, the color, or the type of a card against that list, and not against your memory of the card.`

// BracketJudgeVersion changes when the instructions or the input of the
// bracket judge change. Version 2 reads the combos of Commander
// Spellbook (F-126, D-790). Version 3 states the combo rules of brackets
// 2 and 3 as brackets.json does, and it compares no bracket to a precon
// (F-162, D-794).
const BracketJudgeVersion = 3

// SixtyJudgeVersion names the prompt of the step judge of a 60-card
// import. A sixty-gate document reads it, so a run names the prompt it
// measured. Version 1 defines each step by its anchor (D-866).
const SixtyJudgeVersion = 1

// SummaryJudgeVersion changes when the instructions or the input of the
// summary judge change. Version 2 reads the card facts of the deck
// (D-789). A change starts a new epoch of the false-rule rows.
const SummaryJudgeVersion = 2

const judgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["claims", "verdict"],
  "properties": {
    "claims": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["text", "truth", "why"],
        "properties": {
          "text": {"type": "string"},
          "truth": {"type": "string", "enum": ["true", "false", "unknown"]},
          "why": {"type": "string"}
        }
      }
    },
    "verdict": {"type": "string", "enum": ["clean", "states_a_rule", "states_a_false_rule"]}
  }
}`

// Claim is one rules statement the judge found in a summary.
type Claim struct {
	Text  string `json:"text"`
	Truth string `json:"truth"`
	Why   string `json:"why"`
}

// Judgement is the judge's answer for one summary.
type Judgement struct {
	Claims  []Claim `json:"claims"`
	Verdict string  `json:"verdict"`
}

// StatesAFalseRule reports the F-26 failure: the summary asserts a rule
// of the game and the rule is wrong.
func (j Judgement) StatesAFalseRule() bool {
	if j.Verdict == "states_a_false_rule" {
		return true
	}
	for _, c := range j.Claims {
		if c.Truth == "false" {
			return true
		}
	}
	return false
}

// JudgeSummary asks the judge role whether one summary states a rule of
// the game. name names the deck, for the judge's context. A deck with its
// card source adds the cost and the type of each card, because the judge
// read a true cost as false from its memory of the card (D-789). A nil
// deck sends the summary alone.
func JudgeSummary(ctx context.Context, c *llm.Client, name, summary string, deck *mtgv1.Deck, cards rules.CardSource, acc *llm.Accumulator) (*Judgement, error) {
	input := fmt.Sprintf("Deck: %s\n\nSummary:\n%s", name, summary)
	if deck != nil && cards != nil {
		input += "\n\n" + factsDeckText(deck, cards, nil)
	}
	res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
		Instructions: judgeInstructions,
		Input:        input,
		SchemaName:   "summary_check",
		Schema:       json.RawMessage(judgeSchema),
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("judge summary: %w", err)
	}
	var out Judgement
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("judge summary output: %w", err)
	}
	return &out, nil
}

// The bracket judge is the second bar of the bracket gate (PR-14A). It
// reads a deck with the bracket definitions in hand and names the
// bracket it would play at. The gate asks it to agree with the bracket
// the deck was built for in eight of ten.

const bracketJudgeInstructions = `You read one Commander deck and name the bracket it plays at.

The Commander brackets, from the Commander Format Panel (2025-02-11, revised 2025-10-21):
- Bracket 1, Exhibition: an ultra-casual deck, games of nine turns or more. No Game Changers, no mass land denial, no extra turns, no two-card infinite combos.
- Bracket 2, Core: unoptimized and straightforward, with wins that are incremental, telegraphed on the board, and disruptable, and games of eight turns or more. No Game Changers, no mass land denial, few extra turns and never chained, no two-card infinite combo that needs six mana or less.
- Bracket 3, Upgraded: powered up, with strong synergy and high card quality, and games of six turns or more. Up to three Game Changers, no mass land denial, few extra turns and never chained, no two-card infinite combo that needs four mana or less.
- Bracket 4, Optimized: the strongest cards and decks, games can end from turn four. Only the ban list applies.
- Bracket 5, cEDH: competitive and metagame-focused, a game can end on any turn. Only the ban list applies. The deck plays the best strategy and not a theme.

The card list marks each Game Changer from the card data, the commander included. Count only the marked cards as Game Changers, and never a card you recall from another list.

The list under "Combos" names each combo that Commander Spellbook finds in the deck, with the mana each needs. Count only a listed combo, and never a combo you recall. When the list says the combo check did not run, read the combos from the cards.

Read the card list for its speed, its mana base, its fast mana and tutors, its interaction, its combos, and its Game Changers. Name one bracket, 1 to 5, and say why in two or three sentences. Judge the deck as it is, and not the bracket the builder may have aimed at.`

// The bracket is a string enum and not a bounded integer: the Anthropic
// structured-output schema refuses minimum and maximum on an integer
// (bracket gate run 1, 2026-09-02).
const bracketJudgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["bracket", "why"],
  "properties": {
    "bracket": {"type": "string", "enum": ["1", "2", "3", "4", "5"]},
    "why": {"type": "string"}
  }
}`

// BracketJudgement is the judge's bracket for one deck.
type BracketJudgement struct {
	Bracket int32  `json:"-"`
	Why     string `json:"why"`
}

// bracketOut is the judge's answer as the schema shapes it.
type bracketOut struct {
	Bracket string `json:"bracket"`
	Why     string `json:"why"`
}

// JudgeBracket asks the judge role which bracket a deck plays at. The
// deck goes out as a card list with the commander first, and the judge
// never sees the bracket the deck was built for.
func JudgeBracket(ctx context.Context, c *llm.Client, deck *mtgv1.Deck, cards rules.CardSource, acc *llm.Accumulator) (*BracketJudgement, error) {
	res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
		Instructions: bracketJudgeInstructions,
		Input:        bracketDeckText(deck, cards),
		SchemaName:   "bracket_check",
		Schema:       json.RawMessage(bracketJudgeSchema),
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("judge bracket: %w", err)
	}
	var out bracketOut
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("judge bracket output: %w", err)
	}
	n, err := strconv.Atoi(out.Bracket)
	if err != nil || n < 1 || n > 5 {
		return nil, fmt.Errorf("judge bracket output: bracket %q", out.Bracket)
	}
	return &BracketJudgement{Bracket: int32(n), Why: out.Why}, nil
}

// The step judge reads the power step of a 60-card import (PR-70,
// D-858). Each step is the level of a kind of deck, and the calibration
// set labels its lists by the same kinds (D-863, D-866).
const sixtyJudgeInstructions = `You read one 60-card constructed deck and name the power step it plays at.

The three steps, each the level of a kind of deck:
- casual: the level of a preconstructed deck that Wizards of the Coast sells for play at home: a Theme deck, an Intro pack, or a Planeswalker deck.
- fnm: the level of a Challenger deck, a deck that Wizards of the Coast sells ready to play at Friday Night Magic in a local store.
- tournament: a deck that can finish in the top 8 of a Magic Online Challenge or of a Regional Championship Qualifier.

The first line names the format of the deck. Each card line gives the count, the name, the mana cost, and the type line from the card data.

Read the card list for its card quality, its mana base, its curve, its consistency, its interaction, and its sideboard. Judge the deck against the other decks of its format when its cards were new. Do not lower the step because the cards are old, and do not raise it because they are new. Name one step, and say why in two or three sentences. Judge the deck as it is, and not the step its builder may have aimed at.`

const sixtyJudgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["step", "why"],
  "properties": {
    "step": {"type": "string", "enum": ["casual", "fnm", "tournament"]},
    "why": {"type": "string"}
  }
}`

// SixtyJudgement is the judge's power step for one 60-card deck.
type SixtyJudgement struct {
	Step mtgv1.SixtyStep `json:"-"`
	Why  string          `json:"why"`
}

type sixtyOut struct {
	Step string `json:"step"`
	Why  string `json:"why"`
}

var sixtySteps = map[string]mtgv1.SixtyStep{
	"casual":     mtgv1.SixtyStep_SIXTY_STEP_CASUAL,
	"fnm":        mtgv1.SixtyStep_SIXTY_STEP_FNM,
	"tournament": mtgv1.SixtyStep_SIXTY_STEP_TOURNAMENT,
}

// JudgeSixtyStep asks the judge role which power step a 60-card deck
// plays at. The format word heads the list, because a calibration list
// can name a format the app does not build, such as Pioneer (D-869).
func JudgeSixtyStep(ctx context.Context, c *llm.Client, deck *mtgv1.Deck, format string, cards rules.CardSource, acc *llm.Accumulator) (*SixtyJudgement, error) {
	res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
		Instructions: sixtyJudgeInstructions,
		Input:        sixtyDeckText(deck, format, cards),
		SchemaName:   "sixty_step",
		Schema:       json.RawMessage(sixtyJudgeSchema),
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("judge step: %w", err)
	}
	var out sixtyOut
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("judge step output: %w", err)
	}
	step, ok := sixtySteps[out.Step]
	if !ok {
		return nil, fmt.Errorf("judge step output: step %q", out.Step)
	}
	return &SixtyJudgement{Step: step, Why: out.Why}, nil
}

// SixtyFormatWord names the format of a 60-card import for the step
// judge. The house format checks no legality (D-3, D-857).
func SixtyFormatWord(f mtgv1.FormatId) string {
	if f == mtgv1.FormatId_FORMAT_ID_HOUSE {
		return "house rules, any card and no ban list"
	}
	return FormatWord(f)
}

// sixtyDeckText is the deck with the facts of each card, under a format
// line, and the sideboard after the main deck.
func sixtyDeckText(deck *mtgv1.Deck, format string, cards rules.CardSource) string {
	var s strings.Builder
	fmt.Fprintf(&s, "Format: %s\n", format)
	s.WriteString(factsDeckText(deck, cards, nil))
	if len(deck.GetSideboard()) == 0 {
		return s.String()
	}
	s.WriteString("\nSideboard:\n")
	for _, dc := range deck.GetSideboard() {
		name := dc.GetName()
		if c, ok := cards.ByOracleID(dc.GetOracleId()); ok {
			name += cardFacts(c)
		}
		fmt.Fprintf(&s, "%d %s\n", dc.GetCount(), name)
	}
	return s.String()
}

// DeckText writes a deck as the judge reads it: the commander, then one
// line per card with its count and its job, when the deck names one. No
// bracket, no summary, and no finding.
func DeckText(deck *mtgv1.Deck, cards rules.CardSource) string {
	return deckText(deck, cards, deckTextOpts{})
}

// bracketDeckText is DeckText with a mark on each Game Changer that the
// card data flags, the commander included. The bracket judge counts those
// marks and not a list it recalls (F-123, D-696). The tier judge reads
// DeckText unmarked. The combos of the stored profile follow the cards,
// so the judge counts no combo it recalls (F-126, D-790).
func bracketDeckText(deck *mtgv1.Deck, cards rules.CardSource) string {
	return deckText(deck, cards, deckTextOpts{gameChangers: true}) + comboText(deck.GetProfile().GetContent())
}

// comboText writes the combos of a content check as the bracket judge
// reads them. A check that did not run says so, and the judge then reads
// the combos from the cards.
func comboText(c *mtgv1.ContentCheck) string {
	var s strings.Builder
	s.WriteString("\nCombos:\n")
	switch {
	case c == nil:
		s.WriteString("The combo check did not run for this deck.\n")
	case !c.GetChecked():
		fmt.Fprintf(&s, "The combo check did not run for this deck: %s.\n", c.GetError())
	case len(c.GetCombos()) == 0:
		s.WriteString("Commander Spellbook finds no combo in this deck.\n")
	}
	for _, combo := range c.GetCombos() {
		size := fmt.Sprintf("%d cards, not a two-card combo", len(combo.GetCards()))
		if combo.GetTwoCard() {
			size = "two cards"
		}
		fmt.Fprintf(&s, "- %s (%s, %s)\n", strings.Join(combo.GetCards(), " + "), size, speedWord(combo.GetSpeed()))
	}
	return s.String()
}

// speedWord writes the speed of a Spellbook combo as the mana it needs:
// 5 needs no mana, 4 needs four or less, 3 six or less, 2 eight or less,
// and 1 more. An uncertain minimum adds one, so 6 reads as 5 (variant.py,
// read 2026-09-02).
func speedWord(speed int32) string {
	switch {
	case speed >= 5:
		return "no mana past the cards"
	case speed == 4:
		return "four mana or less"
	case speed == 3:
		return "six mana or less"
	case speed == 2:
		return "eight mana or less"
	}
	return "more than eight mana"
}

// factsDeckText is DeckText with the mana cost and the type line of each
// card, and the color identity of the commander, from the card data. The
// plan judge reads the colors, the curve, and the mana base from these
// facts and not from its memory of the cards (F-39, D-787), and the
// summary judge reads a claim about a card against them (D-789). A request that
// names sets also marks each card in or outside them, because the judge
// read reprints of the Hobbit Commander set as outside it (D-788).
func factsDeckText(deck *mtgv1.Deck, cards rules.CardSource, setCodes []string) string {
	return deckText(deck, cards, deckTextOpts{facts: true, sets: setCodes})
}

type deckTextOpts struct {
	gameChangers bool
	facts        bool
	sets         []string
}

// setNote reads whether a card has a printing in the sets of a request.
func (o deckTextOpts) setNote(c *mtgv1.Card) string {
	for _, have := range c.GetSetCodes() {
		for _, want := range o.sets {
			if strings.EqualFold(have, want) {
				return "in the sets"
			}
		}
	}
	return "outside the sets"
}

func deckText(deck *mtgv1.Deck, cards rules.CardSource, o deckTextOpts) string {
	flagged := func(id string) bool {
		if !o.gameChangers {
			return false
		}
		c, ok := cards.ByOracleID(id)
		return ok && c.GetGameChanger()
	}
	var s strings.Builder
	var identity []mtgv1.Color
	for _, id := range deck.GetCommanderOracleIds() {
		c, ok := cards.ByOracleID(id)
		if !ok {
			continue
		}
		identity = append(identity, c.GetColorIdentity()...)
		name := c.GetName()
		if o.facts {
			name += cardFacts(c)
		}
		var notes []string
		if flagged(id) {
			notes = append(notes, "Game Changer")
		}
		if len(o.sets) > 0 {
			notes = append(notes, o.setNote(c))
		}
		if len(notes) > 0 {
			name += " (" + strings.Join(notes, ", ") + ")"
		}
		fmt.Fprintf(&s, "Commander: %s\n", name)
	}
	if o.facts && len(deck.GetCommanderOracleIds()) > 0 {
		fmt.Fprintf(&s, "Color identity: %s\n", identityWord(identity))
	}
	if len(o.sets) > 0 {
		fmt.Fprintf(&s, "Sets of the request: %s\n", strings.Join(o.sets, ", "))
	}
	s.WriteString("\nCards:\n")
	for _, dc := range deck.GetCards() {
		var notes []string
		if dc.GetRole() != mtgv1.CardRole_CARD_ROLE_UNSPECIFIED {
			notes = append(notes, roleWord(dc.GetRole()))
		}
		if flagged(dc.GetOracleId()) {
			notes = append(notes, "Game Changer")
		}
		name := dc.GetName()
		if c, ok := cards.ByOracleID(dc.GetOracleId()); ok {
			if o.facts {
				name += cardFacts(c)
			}
			if len(o.sets) > 0 {
				notes = append(notes, o.setNote(c))
			}
		}
		if len(notes) == 0 {
			fmt.Fprintf(&s, "%d %s\n", dc.GetCount(), name)
			continue
		}
		fmt.Fprintf(&s, "%d %s (%s)\n", dc.GetCount(), name, strings.Join(notes, ", "))
	}
	return s.String()
}

// cardFacts writes the mana cost and the type line of a card, with a
// leading bar: " | {1}{W}{B} | Legendary Creature". A card with no cost,
// such as a land, writes the type line alone. A card with two faces and
// no cost of its own writes the cost of each face.
func cardFacts(c *mtgv1.Card) string {
	cost := c.GetManaCost()
	if cost == "" {
		var faces []string
		for _, f := range c.GetFaces() {
			if f.GetManaCost() != "" {
				faces = append(faces, f.GetManaCost())
			}
		}
		cost = strings.Join(faces, " // ")
	}
	var parts []string
	if cost != "" {
		parts = append(parts, cost)
	}
	if t := c.GetTypeLine(); t != "" {
		parts = append(parts, t)
	}
	if len(parts) == 0 {
		return ""
	}
	return " | " + strings.Join(parts, " | ")
}

// identityWord writes a color identity in WUBRG letters, and the word
// colorless for none.
func identityWord(colors []mtgv1.Color) string {
	letters := map[mtgv1.Color]string{
		mtgv1.Color_COLOR_W: "W", mtgv1.Color_COLOR_U: "U", mtgv1.Color_COLOR_B: "B",
		mtgv1.Color_COLOR_R: "R", mtgv1.Color_COLOR_G: "G",
	}
	var s strings.Builder
	for _, col := range AllColors {
		if hasColor(colors, col) {
			s.WriteString(letters[col])
		}
	}
	if s.Len() == 0 {
		return "colorless"
	}
	return s.String()
}

// The tier judge of PR-14B: the judge names the rung of the ladder of
// D-414 a deck sits on, and the gate reads whether it agrees with the
// quality model's grade in eight of ten.
const tierJudgeInstructions = `You read one deck list and name its quality tier.

The ladder, from the top:
- great: a list that would place in the top 8 of a Challenge or the top cut of a cEDH tournament. Tuned, efficient, the cards and the pairs the best lists of the format play.
- good: a list that would post a winning record in a league or finish mid-field at a Challenge. Sound, on plan, a step below the top.
- typical: the average deck of its commander or archetype as the community builds it. Reasonable choices, no sharp edge.
- baseline: the strength of a preconstructed deck as sold. Playable, slow, generic cards, few of the format's best.
- bad: below a precon. A broken mana base, a curve that never lands its spells, playsets turned to singletons, or cards that do not work together.

Read the card list for its card quality against the format's best lists, its synergy, its mana base, its curve, and its interaction. Name one tier word and say why in two or three sentences. Judge the deck as it is.`

const tierJudgeSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["tier", "why"],
  "properties": {
    "tier": {"type": "string", "enum": ["great", "good", "typical", "baseline", "bad"]},
    "why": {"type": "string"}
  }
}`

// TierJudgement is the judge's tier for one deck.
type TierJudgement struct {
	Tier string `json:"tier"`
	Why  string `json:"why"`
}

// JudgeTier asks the judge role for the tier of a deck. The judge sees
// the format and the card list, and never the grade.
func JudgeTier(ctx context.Context, c *llm.Client, deck *mtgv1.Deck, cards rules.CardSource, acc *llm.Accumulator) (*TierJudgement, error) {
	res, err := c.Complete(ctx, llm.RoleJudge, llm.Request{
		Instructions: tierJudgeInstructions,
		Input:        "Format: " + FormatWord(deck.GetFormat().GetId()) + "\n" + DeckText(deck, cards),
		SchemaName:   "tier_check",
		Schema:       json.RawMessage(tierJudgeSchema),
	}, acc)
	if err != nil {
		return nil, fmt.Errorf("judge tier: %w", err)
	}
	var out TierJudgement
	if err := json.Unmarshal(res.Output, &out); err != nil {
		return nil, fmt.Errorf("judge tier output: %w", err)
	}
	if out.Tier == "" {
		return nil, fmt.Errorf("judge tier output: no tier")
	}
	return &out, nil
}
