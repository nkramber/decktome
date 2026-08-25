package questions

// The two prompts of the question workflow. Keep them stable: both
// providers cache a stable instruction prefix, and a changed prompt
// invalidates the M-5 scores that carry a prompt version (D-66).

const classifyInstructions = `You map one message from a Magic: The Gathering deck-building conversation onto slots.

Rules:
- Fill a slot only from what the user wrote or clearly implied. Never guess.
- When the message answers a question about the format, the theme, the colors, the power, the pool rule, the budget, or the commander, put the answer in that field. A field is the only place an answer counts. Naming the slot in closed_keys does not record the answer, and the agent asks again.
- Repeat a value the user gave in an earlier message when the field is still empty. "Pioneer" said two messages ago is still the format.
- Leave a field empty, zero, or "unknown" when the message does not answer it.
- format: the format the user named, even as one word on its own. "Commander" means commander, "Modern" means modern, "Standard" means standard. Copy it into the format field every time the user names one. Use "unknown" only when the message names no format at all. "anything goes" is not a format.
- theme: the plan in the user's own words, for example "lifegain" or "mill". An answer such as "the best deck under budget" or "a named tier-one deck" is also a theme.
- power: a Commander bracket as "bracket 3", or a 60-card step as "casual", "fnm", or "tournament". Vague words such as "strongest", "competitive", or "best" are not a step. Leave power empty for those and set facts.power_competitive.
- pool_rule: "owned_first" when the user builds from their library first, "owned_only" when only owned cards may be used, "any_card" when the library does not constrain the deck.
- declined_keys: the keys in open_keys that the user handed back to you. A decline is not an answer, and it names no value. Name a key only when the user's words are about that key. "Any colors are fine" declines the colors and nothing else. "You decide" with no subject declines every key in open_keys. Never put a key in both lists.
- closed_keys: for the advisory keys only, and only those listed in open_keys. It never carries a format, theme, colors, power, pool rule, budget, or commander answer. A refusal, "none", or "name three more" closes no key. A phrase that only raises a topic closes no key: "we proxy everything" raises house rules, and it does not say which cards are legal.
- facts.out_of_scope: the user asked for something this app does not build, such as a deck for another card game. A Magic: The Gathering request is always in scope.
- facts.named_card: the user named a specific card.
- facts.buy_list: the deck will need cards the user does not own.
- facts.house_format: the user described their own rule set instead of a real format.
- facts.two_plans: the theme has two common plans and the user has not chosen one.
- facts.budget_ambiguous: the user named one money number without saying whether it caps purchases or the whole deck.
- facts.power_competitive: the user asked for a strong, competitive, or winning deck.
- facts.wants_suggestion: the user asked you to name a commander, or said they have none in mind.
- commander_names: a card the user wants as the commander. locked_names: a card the user wants in the deck but not as the commander. Never put one name in both lists.

Answer with the schema only.`

const classifySchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["format","theme","colors","commander_names","locked_names","power","pool_rule","budget_usd","closed_keys","declined_keys","facts"],
  "properties": {
    "format": {"type": "string"},
    "theme": {"type": "string"},
    "colors": {"type": "array", "items": {"type": "string", "enum": ["W","U","B","R","G"]}},
    "commander_names": {"type": "array", "items": {"type": "string"}},
    "locked_names": {"type": "array", "items": {"type": "string"}},
    "power": {"type": "string"},
    "pool_rule": {"type": "string"},
    "budget_usd": {"type": "number"},
    "closed_keys": {"type": "array", "items": {"type": "string"}},
    "declined_keys": {"type": "array", "items": {"type": "string"}},
    "facts": {
      "type": "object",
      "additionalProperties": false,
      "required": ["named_card","buy_list","house_format","two_plans","budget_ambiguous","power_competitive","wants_suggestion","out_of_scope"],
      "properties": {
        "named_card": {"type": "boolean"},
        "buy_list": {"type": "boolean"},
        "house_format": {"type": "boolean"},
        "two_plans": {"type": "boolean"},
        "budget_ambiguous": {"type": "boolean"},
        "power_competitive": {"type": "boolean"},
        "wants_suggestion": {"type": "boolean"},
        "out_of_scope": {"type": "boolean"}
      }
    }
  }
}`

const askInstructions = `You phrase clarifying questions for a Magic: The Gathering deck builder.

You get the user's last message and the questions the agent decided to ask. For each one:
- text: the same question in natural words, fitted to what the user wrote. Keep the meaning. Replace any {placeholder} with a real value, or drop that clause when you have no value.
- options: the given options, reworded to match. Keep them short. An empty list is fine.

Never change what a question asks. Never merge two questions. Never add a question. Answer with the schema only.`

const askSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["questions"],
  "properties": {
    "questions": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["row_id","text","options"],
        "properties": {
          "row_id": {"type": "string"},
          "text": {"type": "string"},
          "options": {"type": "array", "items": {"type": "string"}}
        }
      }
    }
  }
}`

// The gap score has its own classifier call (owner directive, 2026-08-24).
// The roadmap asks for a small classifier, and a scorer that also phrases
// the question rates its own work. The cost is one more call per turn.

const scoreInstructions = `You check a fixed catalog question against one user of a Magic: The Gathering deck builder.

You get the user's last message, the conversation so far, and the catalog rows the agent plans to ask. Judge each row against four faults, and nothing else. Do not reward wording you merely like better.

1. Contradiction: the row assumes something the user has denied.
2. Already answered: the user has given this value.
3. Wrong slot: the row asks about a subject other than its slot.
4. Unanswerable: the row names a thing this user can not know, or it asks two things at once.

Score by the count of faults you find:
- fit 0.9: no fault. The row asks a thing this user has not answered.
- fit 0.5: one fault, and it is small. The row still works.
- fit 0.2: one clear fault.
- fit 0.05: two or more faults.

A row that repeats the catalog wording is not a fault. A row that is short is not a fault. A row that does not name the user's colors, format, or cards is not a fault, because another step adds those words.

custom_text: write one only when you scored under the threshold. It must ask for the same slot, keep every option the row offers, and ask one thing. Leave it empty otherwise.
reason: name the fault you found, or say "no fault".

The catalog is the default. A replacement is the exception. Answer with the schema only.`

const scoreSchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["scores"],
  "properties": {
    "scores": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["row_id","fit","custom_text","reason"],
        "properties": {
          "row_id": {"type": "string"},
          "fit": {"type": "number"},
          "custom_text": {"type": "string"},
          "reason": {"type": "string"}
        }
      }
    }
  }
}`
