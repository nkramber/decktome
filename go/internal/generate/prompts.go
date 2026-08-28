package generate

// The generate and repair prompts. Keep them stable: both providers cache
// a stable instruction prefix, and the session text goes last so the
// prefix stays the same across a repair turn and across PR-9's re-rolls
// (roadmap PR-8).

// PromptVersion counts the changes to the prompts below. A deck scored at
// an earlier version does not carry over, which is the D-66 rule for the
// question prompts.
//
// Version 1 is the first generator.
//
// Version 4 followed deck gate run 5. The deck cost $268.37 against a
// $100.00 cap, on a shortlist whose cheapest 99 cards cost $25.66. No
// shortlist line carried a price, so the model could not budget at all.
// The lines carry prices now, the budget is a limit rather than a line of
// the plan prose, and going over it buys the repair turn (D-244).
//
// Version 3 followed deck gate run 3. The repair turn wrote a changelog
// into the summary: "Sol Ring remains included. Skullport Merchant now
// appears once; Warren Soultrader fills the replaced ramp slot." The
// reader of a summary does not know a repair happened, so the repair
// prompt now asks for the summary again from nothing.
//
// Version 2 followed the first real call, on 2026-08-27. The summary read
// "uses its 20 synergy cards ... while maintaining 10 draw cards, 10 ramp
// cards, eight removal cards, and three wipes. There are no target
// shortfalls." The model recited the job targets back, because version 1
// asked it to state a shortfall there. The summary is for the user, and
// the counts are in the card list.
const PromptVersion = 4

// generateInstructions is the stable prefix. It names no card, no format,
// and no session value, so every call of a session shares it.
//
// Two rules carry a fault number. F-13 says a model can name a card that
// exists and is not the card meant, so the deck list must copy names from
// the shortlist. F-26 says the model invents claims about the game, and
// the summary is prose with more room for that than a question ever had.
const generateInstructions = `You build a Magic: The Gathering deck from a shortlist.

You get the deck-building limits for one format, a shortlist of cards with the job each one does, a target count for each job, and the plan the user asked for.

Rules for the card list:
- Use only cards from the shortlist. Copy each name exactly as the shortlist writes it, character for character.
- Never name a card that is not on the shortlist. A card you remember is not on the shortlist.
- Give every card a count. Respect the copy limit in the limits block.
- Give every card one job from the job list, and one line that says why the card is in the deck.
- Meet the deck size in the limits block. Count the commander when the limits block says to.
- Come as close to each job target as the shortlist allows.

Rules for the summary:
- Write one paragraph for the person who will play the deck. Say what the deck does on the table, how it wins, and what it gives up.
- Name no job target and no job count. The card list carries them, and a count means nothing to the reader.
- Write plainly. Do not report on your own work, and do not say whether you met the targets.
- State no rule of the game. Do not say what a card may do, what a format allows, what is banned, or what is legal. The rules engine checks the deck and reports that.
- Claim nothing about a card that the shortlist does not say.
- Name no price and no date.
- Do not address the user by name, and do not ask a question.`

// repairInstructions runs after the engine or the normalizer refuses the
// deck. It is a second stable prefix, and the findings arrive in the
// input with the session text.
const repairInstructions = `You repair a Magic: The Gathering deck you built.

You get the same limits, the same shortlist, and the deck you returned. You also get the findings against it.

Rules:
- Fix every finding. Change as few cards as the fix needs.
- Use only cards from the shortlist, and copy each name exactly as the shortlist writes it.
- A finding that names a card you invented means the card is not on the shortlist. Replace it with a shortlist card that does the same job. Never write the name again.
- A finding that the deck costs too much means you must swap dear cards for cheaper ones that do the same job. Each shortlist line ends with the price of one copy. Come under the cap.
- Return the whole deck, and not the change alone.
- Write the summary again from nothing. It describes the deck, and never the repair. Name no card you changed, no count, and no slot you filled. A reader of the summary does not know a first turn happened.
- The summary rules of the first turn still hold. State no rule of the game.`

// deckSchema is the output contract. Strict mode applies, so every object
// carries additionalProperties false and a full required list.
const deckSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["summary", "cards", "sideboard"],
  "properties": {
    "summary": {"type": "string"},
    "cards": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["name", "count", "role", "reason"],
        "properties": {
          "name": {"type": "string"},
          "count": {"type": "integer"},
          "role": {
            "type": "string",
            "enum": ["land", "ramp", "draw", "removal", "wipe", "threat", "interaction", "synergy", "wincon", "other"]
          },
          "reason": {"type": "string"}
        }
      }
    },
    "sideboard": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["name", "count", "role", "reason"],
        "properties": {
          "name": {"type": "string"},
          "count": {"type": "integer"},
          "role": {
            "type": "string",
            "enum": ["land", "ramp", "draw", "removal", "wipe", "threat", "interaction", "synergy", "wincon", "other"]
          },
          "reason": {"type": "string"}
        }
      }
    }
  }
}`

// deckOut is the model's answer.
type deckOut struct {
	Summary   string  `json:"summary"`
	Cards     []Entry `json:"cards"`
	Sideboard []Entry `json:"sideboard"`
}
