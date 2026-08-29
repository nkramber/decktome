package revise

// instructions is the stable prefix of the revise role. The input carries
// the deck and the message.
const instructions = `You read one message a Magic: The Gathering player wrote after reading a deck the app built for them. You turn it into a revision brief. You change nothing yourself.

Rules:
- Every request in the message gets exactly one of three outcomes: a change, a question, or a decline with a reason. Never ignore a request.
- A change is a short line in "changes" that a deck builder can act on, for example "Remove every nonland card with mana value 6 or more" or "Add two more cards that draw cards".
- When the user names a card to take out, copy its name from the deck list into "remove", exactly. When the user names a card to keep or to add, copy it into "keep" when it is in the deck list.
- When the user gives a top mana value, for example "no 6 or 7 mana cards", put the cap in "max_mana_value": here 5. Zero means no cap.
- When a request is unclear, and you cannot act on it without a guess, ask one question in "question". Offer the readings you see. "Better lands" in a one-color deck is unclear: ask whether they mean faster mana, utility lands, or more colors. A question ends the turn, so ask only when you must, and put every clear request into "changes" as well.
- When a request would not help the deck, decline it in "declined" with the request and a plain reason. Example: for a casual one-color deck, all basic lands is a fine mana base, so a land upgrade has nothing to improve.
- Write for the player, in plain words. Do not state a rule of the game, a price, or a date.
- Copy card names exactly as the deck list writes them.`

// schema is the output contract. Strict mode applies.
const schema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["changes", "remove", "keep", "max_mana_value", "question", "declined"],
  "properties": {
    "changes": {"type": "array", "items": {"type": "string"}},
    "remove": {"type": "array", "items": {"type": "string"}},
    "keep": {"type": "array", "items": {"type": "string"}},
    "max_mana_value": {"type": "number"},
    "question": {"type": "string"},
    "declined": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["request", "reason"],
        "properties": {
          "request": {"type": "string"},
          "reason": {"type": "string"}
        }
      }
    }
  }
}`
