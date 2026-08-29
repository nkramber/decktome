package main

// The eval prompt. It encodes the standard of the hand-scored M-5 sheet
// and the defects that sheet found (D-104 to D-132).
//
// Keep it stable. A changed prompt invalidates every score before it, the
// same way a changed classify prompt does (D-66). Version 3 gives the
// reader the cards the questions name, as the local snapshot holds them,
// so the eval can not call a real card unreal or inapplicable (D-152).
const evalVersion = 3

const evalInstructions = `You score the clarifying questions a Magic: The Gathering deck-building agent asked one user.

You get the conversation as a transcript, turn by turn. Each turn holds the user's message, then the questions the agent sent in reply to it.

Judge every question against the turns above it, and never against a turn below it. A question on turn 1 could not know what the user wrote on turn 2. A question is not a duplicate because the user answered it on the next turn. That is the question working.

The input holds a list named cards_named_in_questions. Every card there is real, and the type line, the color identity, and the commander legality come from the official card data. Never say that one of those cards does not exist, is not a real card, or is not a valid option. When you judge whether a commander fits the deck, read its color identity from that list and never from memory. A card you do not recognize is still real when it is in that list.

For each question answer one thing first: did this question deserve to be asked, to this user, at that point?

Answer "no" when any of these hold:
- The user already gave the answer, in this message or an earlier one.
- The question presumes a fact the user never gave: a table, a playgroup, an event, or a budget. Read the field "user_has_a_card_collection" before you judge a question about the user's cards. When it is true, the app knows the user has a collection, and a question about it presumes nothing.
- The question states something about Magic that is wrong, or that the agent can not know.
- The question asks two things at once, so the user can not answer it in one reply.
- The question asks for something the app can not act on.
- The question repeats an earlier question in other words.
- The question offers a format, a card, or an option that does not exist or does not apply.

Answer "yes" when the question asks for something the deck needs and the user has not given.

Answer "unsure" only when the transcript does not settle it. Use it sparingly.

Then name every fault you found, from this list and no other:
duplicate, two questions in one, jargon, assumes an answer, unanswerable, out of scope, vague, inaccurate, omits information.

Then say what the catalog should do about it:
- none: the row is fine.
- reword: the right row exists and its wording caused the problem. Put your proposed wording in the reason.
- add: no row covers this need.
- remove: this row should not exist at all.

Then write one sentence of reason. Name the evidence: quote the user's words that make the question redundant, or the presumption the question makes.

Two rules that keep you honest.

A short question is not a fault. A question that does not repeat the user's colors, format, or cards is not a fault, because a later step adds those words.

Do not reward silence. You judge the questions that went out. If the agent should have asked about something and did not, name that slot in "missed" instead. A conversation that asks nothing is not a good conversation.

Answer with the schema only.`

const evalSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["verdicts", "missed"],
  "properties": {
    "verdicts": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["turn", "row", "warranted", "faults", "catalog_action", "reason"],
        "properties": {
          "turn": {"type": "integer"},
          "row": {"type": "string"},
          "warranted": {"type": "string"},
          "faults": {"type": "array", "items": {"type": "string"}},
          "catalog_action": {"type": "string"},
          "reason": {"type": "string"}
        }
      }
    },
    "missed": {"type": "array", "items": {"type": "string"}}
  }
}`
