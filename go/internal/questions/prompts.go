package questions

// The two prompts of the question workflow. Keep them stable: both
// providers cache a stable instruction prefix, and a changed prompt
// invalidates the M-5 scores that carry a prompt version (D-66).

const classifyInstructions = `You map one message from a Magic: The Gathering deck-building conversation onto slots.

Rules:
- Fill a slot only from what the user wrote or clearly implied. Never guess.
- Leave a field empty, zero, or "unknown" when the message does not answer it.
- format: the format the user named. "unknown" when none is clear. "anything goes" is not a format.
- theme: the plan in the user's own words, for example "lifegain" or "mill".
- power: a Commander bracket as "bracket 3", or a 60-card step as "casual", "fnm", or "tournament". Vague words such as "strongest", "competitive", or "best" are not a step. Leave power empty for those and set facts.power_competitive.
- pool_rule: "owned_first" when the user builds from their library first, "owned_only" when only owned cards may be used, "any_card" when the library does not constrain the deck.
- closed_keys: the open keys this message answered. Use only keys from open_keys.
- facts.named_card: the user named a specific card.
- facts.buy_list: the deck will need cards the user does not own.
- facts.deadline: the user named a date or an event day.
- facts.house_format: the user described their own rule set instead of a real format.
- facts.two_plans: the theme has two common plans and the user has not chosen one.
- facts.budget_ambiguous: the user named one money number without saying whether it caps purchases or the whole deck.
- facts.power_competitive: the user asked for a strong, competitive, or winning deck.

Answer with the schema only.`

const classifySchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["format","theme","colors","commander_names","locked_names","power","pool_rule","budget_usd","closed_keys","facts"],
  "properties": {
    "format": {"type": "string", "enum": ["commander","standard","pioneer","modern","legacy","vintage","pauper","house","unknown"]},
    "theme": {"type": "string"},
    "colors": {"type": "array", "items": {"type": "string", "enum": ["W","U","B","R","G"]}},
    "commander_names": {"type": "array", "items": {"type": "string"}},
    "locked_names": {"type": "array", "items": {"type": "string"}},
    "power": {"type": "string"},
    "pool_rule": {"type": "string", "enum": ["owned_first","owned_only","any_card","unknown"]},
    "budget_usd": {"type": "number"},
    "closed_keys": {"type": "array", "items": {"type": "string"}},
    "facts": {
      "type": "object",
      "additionalProperties": false,
      "required": ["named_card","buy_list","deadline","house_format","two_plans","budget_ambiguous","power_competitive"],
      "properties": {
        "named_card": {"type": "boolean"},
        "buy_list": {"type": "boolean"},
        "deadline": {"type": "boolean"},
        "house_format": {"type": "boolean"},
        "two_plans": {"type": "boolean"},
        "budget_ambiguous": {"type": "boolean"},
        "power_competitive": {"type": "boolean"}
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

const scoreInstructions = `You score how well a fixed catalog question fits one user of a Magic: The Gathering deck builder.

You get the user's last message, the conversation so far, and the catalog rows the agent plans to ask. For each row:
- fit: 0 to 1. How well this row asks what this user needs to be asked next. 1 means it fits exactly. Score the wording and the assumption behind it, not the slot.
- custom_text: a replacement question that fills the same slot, and only when the fit is below the threshold. Leave it empty otherwise. One question, no compound clauses.
- reason: one short sentence for the score.

Score low when the row assumes something the user contradicted, when it reads as if the agent ignored the message, or when it asks for a thing the user already gave. Score high when the row is the obvious next question. The catalog is the default: a replacement is the exception. Answer with the schema only.`

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
