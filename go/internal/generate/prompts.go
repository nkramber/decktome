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
// Version 2: the summary names no job target and no job count, because
// the summary is for the user and the counts are in the card list.
//
// Version 3: the repair turn writes the summary again from nothing, so a
// summary never describes the repair.
//
// Version 4: every shortlist line carries a price when a budget applies,
// the budget is a limit and not plan prose, and going over it buys the
// repair turn (D-244).
//
// Version 5: the precon share is a card count and not a percentage, and
// a shortfall buys the repair turn (D-248).
//
// Version 6: an upgrade gets no job targets, because a precon is a
// working deck (D-249).
//
// Version 7: an upgrade keeps the precon's land count, because the mana
// base is not a job (D-251).
//
// Version 8: the upgrade prompt names the precon's own land count and
// sends no generic land target, which was a second quota (D-251).
//
// Version 9: the share is 85 percent of the precon's nonbasic names,
// basic lands swap free, and the change count is a ceiling and not a
// target. The 60-card sideboard is "up to 15 cards" (D-218).
//
// Version 10: the shortlist omits the commander, which the prompt says
// is not one of the cards to list, and the house format gets the
// sideboard sentence (D-302).
//
// Version 11: a change that names a group of cards and a number touches
// that many cards of the group, and the repair turn reads a land swap
// the deck did not make (D-448).
//
// Version 12: the input carries a deck shape block with the bands of the
// bracket, and the repair turn reads a profile finding: a feature off
// its band, or a card or combo the bracket forbids (PR-14A).
const PromptVersion = 12

// generateInstructions is the stable prefix. It names no card, no format,
// and no session value, so every call of a session shares it.
//
// Two rules carry a fault number. F-13 says a model can name a card that
// exists and is not the card meant, so the deck list must copy names from
// the shortlist. F-26 says the model invents claims about the game, and
// the summary is prose with more room for that than a question ever had.
const generateInstructions = `You build a Magic: The Gathering deck from a shortlist.

You get the deck-building limits for one format, a shortlist of cards with the job each one does, a target count for each job, and the plan the user asked for. A deck shape block, when present, states the limits of the power level: the curve, the mana base, and the counts of tutors and fast mana. Build inside them.

Rules for the card list:
- Use only cards from the shortlist. Copy each name exactly as the shortlist writes it, character for character.
- Never name a card that is not on the shortlist. A card you remember is not on the shortlist.
- Give every card a count. Respect the copy limit in the limits block.
- Give every card one job from the job list, and one line that says why the card is in the deck.
- Meet the deck size in the limits block. Count the commander when the limits block says to.
- Come as close to each job target as the shortlist allows.

Rules for a revision, when the input holds the deck you are revising:
- The user read that deck and asked for a change. The input lists the change in short lines.
- Keep every card the change does not touch, at the same count and the same job. A revision is the smallest set of changes that does what the user asked, and not a new deck.
- Do what every line of the change says. A card the user wants out is not on the shortlist. A card the user wants kept must stay.
- When a change names a group of cards and a number, for example "replace at least 20 basic lands", it touches that many cards of the group. One card of the group is not the change.
- Write the summary for the deck as it is now. Never describe the change.

Rules for an upgrade, when the input names a precon:
- The precon is a working deck. Change as few cards as the upgrade needs, and keep its theme intact.
- The input states how many precon names you may change at most. That number is a ceiling and not a target. Never aim to replace that many.
- Basic lands are free to swap and are not counted.

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
- A finding that the deck holds too few new nonbasic lands means you kept basic lands the change told you to replace. Cut more basic lands and add nonbasic lands from the shortlist, of the kinds the change names, until the count is met. Keep the land total the same.
- A finding that the deck keeps too few precon names means you dropped too many. Put back the ones marked "precon" until the count is met, and drop cards that are not marked instead. Keep the theme of the precon, and change no more than the fix needs.
- A finding that a count is off its band names the count, the value, and the range the power level wants. Move the count into the range: add or cut cards of that job, or swap lands, and keep the deck size. A finding about the average mana value means swap dear cards for cheaper ones that do the same job, or the reverse.
- A finding that names a card or a combo the power level forbids means cut that card, or one card of the combo, and replace it with a shortlist card that does the same job.
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
