package questions

// The three prompts of the question workflow. Keep them stable: both
// providers cache a stable instruction prefix, and a changed prompt
// invalidates the M-5 scores that carry a prompt version (D-66).

// PromptVersion counts the changes to the prompts below. The gate
// document records it, so a scored M-5 sheet names the prompts that
// produced it (D-66).
//
// Version 2: the classify role reads a format from an adjective and from
// a commander phrase, and the ask role adds no clause that repeats a
// value the user gave (D-116).
//
// Version 3 added the named_cards list. "Build around X" names a card
// and no role, and the classifier must not report X as the commander
// (D-118).
//
// Version 6: the ask role must not fit a question to the format the
// agent has just declined (D-150).
//
// Version 7: the ask role must not name one card's color identity, in
// any preposition (D-151).
//
// Version 8 followed D-155, which narrowed the app to Commander,
// Standard, and Modern. The classify role must not resolve a format the
// app no longer builds, because the unsupported-format row declines it
// by name instead.
//
// Version 9: the classify role reads cEDH as bracket 5 (D-164).
//
// Version 10: the classify role reports the budget scope, so the answer
// to the budget-scope row is stored (D-238).
//
// Version 11: the input carries prior_messages, the user's last five
// earlier messages, in place of a rule that told the model to repeat a
// value it never saw (D-90). The input also carries nearest_format, so
// a "yes" to the decline row fills the format (D-112). The instruction
// prefix is unchanged in shape, so the provider cache still serves it.
//
// Version 12 stores the house-rules answer (D-265). The schema gains a
// house_rules string, and five rows left the catalog (D-260), so the
// instruction text names fewer rows.
//
// Version 13 rewrites the example of the ask prompt to the commander row
// of D-290. The instruction text changed, so the provider cache prefix
// changed with it, and the question gate re-baselines (D-66).
// Version 14 drops the two_plans fact from the classify role. No catalog
// row reads it, and the owned_mode fact left the planner with it (D-302).
//
// Version 15 adds set_names to the classify role (D-373). A set is a
// constraint the app applies now, so the words that name one no longer
// go into the theme alone. The D-371 sentence that read a set limit as a
// theme is rewritten, and the ownership half of it stays: "only" is
// about a library and never about a set.
const PromptVersion = 15

const classifyInstructions = `You map one message from a Magic: The Gathering deck-building conversation onto slots.

Rules:
- Fill a slot only from what the user wrote or clearly implied. Never guess.
- When the message answers a question about the format, the theme, the colors, the power, the pool rule, the budget, or the commander, put the answer in that field. A field is the only place an answer counts. Naming the slot in closed_keys does not record the answer, and the agent asks again.
- prior_messages in the input are the user's earlier messages, oldest first. Read them for a value the user gave before and has not replaced. "Modern" in an earlier message is still the format when no later message names another.
- nearest_format in the input is the format the agent offered in place of one it does not build. When the message accepts it ("yes", "use that", "fine, treat it as Commander"), copy nearest_format into the format field.
- Leave a field empty, zero, or "unknown" when the message does not answer it.
- format: the format the user named, even as one word on its own. This app builds three: "Commander", "Standard", and "Modern". Copy one of those three into the format field every time the user names one. Use "unknown" for every other format, including Pioneer, Legacy, Vintage, Pauper, Brawl, and Historic: another step declines those by name, and naming one here would build the wrong deck. Use "unknown" also when the message names no format at all. "anything goes" is not a format.
- format from an adjective: "a Commander deck", "a Modern burn deck", and "a Standard burn deck" all name the format. Read it. "EDH" means commander.
- format from a commander phrase: a message that says "my commander", "not as my commander", "in the 99", "bracket 3", or "my precon" means the commander format, even when the word Commander is absent. Fill the format field from it.
- theme: the plan in the user's own words, for example "lifegain" or "mill". An answer such as "the best deck under budget" or "a named tier-one deck" is also a theme.
- power: a Commander bracket as "bracket 3", or a 60-card step as "casual", "fnm", or "tournament". Vague words such as "strongest", "competitive", or "best" are not a step. Leave power empty for those and set facts.power_competitive.
- "cEDH" is a power level and a format. It means bracket 5, and the deck is Commander. "Competitive Commander" is not the same thing: it names no bracket.
- pool_rule: "owned_first" when the user builds from their library first, "owned_only" when only owned cards may be used, "any_card" when the library does not constrain the deck.
- pool_rule is about ownership alone. A limit to a set, a block, a color, or a card type names no pool rule. "Build only from the Hobbit set", "only cards from Bloomburrow", and "only artifacts" all leave pool_rule empty. The word "only" means owned_only when it is about the user's own cards, as in "only cards I own" or "only what is in my collection".
- set_names: the Magic sets or products the user wants the deck built from, in the user's own words. "Build only from the Hobbit set" gives ["the Hobbit set"]. "Cards from Bloomburrow and Duskmourn" gives ["Bloomburrow", "Duskmourn"]. Write the name the user wrote, and add no set the user did not name. Leave the list empty when the user named no set.
- A set is not a theme. "Build only from the Hobbit set" names a set and no theme, so set_names holds it and theme stays empty. "A Hobbit-set dragons deck" names both: set_names holds "the Hobbit set" and theme holds "dragons".
- A creature type, a mechanic, a play style, or a card type is a theme and never a set. "only artifacts" is a theme. A set is a product name, such as Bloomburrow, Duskmourn, Final Fantasy, or Modern Horizons 3.
- A refusal of the names on the table is neither an answer nor a decline. Under commander_pick alone, "None of those", "none of these", and "name three more" leave that key open, and they name no key in either list. This rule is about the offered names only. It never applies to another key.
- declined_keys: the keys in open_keys that the user handed back to you. A decline is not an answer, and it names no value. Name a key only when the user's words are about that key. "Any colors are fine" declines the colors and nothing else. "You decide" with no subject declines every key in open_keys. Never put a key in both lists.
- A negative answer to a question that invites a yes or a no is a decline. "Do you have a color preference?" answered "No" declines colors. "Do you have a budget for cards to buy?" answered "No" declines budget. Read "no", "none", "no preference", "not really", "any", and "it does not matter" the same way. The user has said there is no such constraint, so the key must close. Leaving it open stops the deck for good.
- A quoted question tells you which key an answer belongs to. A line that starts "Q: " is the question this app asked, and the line under it that starts "A: " is the user's answer to that question and to nothing else.
- closed_keys: for the advisory keys only, and only those listed in open_keys. It never carries a format, theme, colors, power, pool rule, budget, or commander answer. A refusal, "none", or "name three more" closes no key. A phrase that only raises a topic closes no key: "we proxy everything" raises house rules, and it does not say which cards are legal.
- facts.out_of_scope: the user asked for something this app does not build, such as a deck for another card game. A Magic: The Gathering request is always in scope.
- facts.named_card: the user named a specific card.
- facts.buy_list: the deck will need cards the user does not own.
- facts.house_format: the user described their own rule set instead of a real format.
- facts.budget_ambiguous: the user named one money number without saying whether it caps purchases or the whole deck.
- house_rules: what the user means by "anything goes", "kitchen table", or "no ban list", in the user's own words, for example "any card, no ban list" or "Modern with proxies". Fill it when the user answers the house-rules question, or states the rules unprompted. Leave it empty otherwise.
- budget_scope: what the cap covers, when the user says. "buy" means the cards they must acquire, and "deck" means the whole deck value, owned copies included. Leave it "unknown" when the user did not say.
- facts.power_competitive: the user asked for a strong, competitive, or winning deck.
- facts.wants_suggestion: the user asked you to name a commander, or said they have none in mind.
- offered_commanders in the input are the commanders the agent just named. When the user picks one of them, by name or by place ("the first", "the second one"), put that commander in commander_names.
- Three card lists, and a name goes in exactly one of them. commander_names: a card the user wants as the commander. locked_names: a card the user wants in the deck but not as the commander. named_cards: a card the user named without saying what role it plays.
- "Build around X" names a card and no role. Put X in named_cards. Do not guess that X is the commander, because the agent asks which role the user wants.
- "Build around X, but not as my commander" puts X in locked_names, and never in commander_names. The user answered the role question before you asked it.
- Write a card name once, and write it in full. Do not report both "Grist" and "Grist, the Hunger Tide".

Answer with the schema only.`

const classifySchema = `{
  "type": "object",
  "additionalProperties": false,
  "required": ["format","theme","colors","commander_names","locked_names","named_cards","set_names","power","pool_rule","budget_usd","budget_scope","house_rules","closed_keys","declined_keys","facts"],
  "properties": {
    "format": {"type": "string"},
    "theme": {"type": "string"},
    "colors": {"type": "array", "items": {"type": "string", "enum": ["W","U","B","R","G"]}},
    "commander_names": {"type": "array", "items": {"type": "string"}},
    "locked_names": {"type": "array", "items": {"type": "string"}},
    "named_cards": {"type": "array", "items": {"type": "string"}},
    "set_names": {"type": "array", "items": {"type": "string"}},
    "power": {"type": "string"},
    "pool_rule": {"type": "string"},
    "budget_usd": {"type": "number"},
    "budget_scope": {"type": "string", "enum": ["buy", "deck", "unknown"]},
    "house_rules": {"type": "string"},
    "closed_keys": {"type": "array", "items": {"type": "string"}},
    "declined_keys": {"type": "array", "items": {"type": "string"}},
    "facts": {
      "type": "object",
      "additionalProperties": false,
      "required": ["named_card","buy_list","house_format","budget_ambiguous","power_competitive","wants_suggestion","out_of_scope"],
      "properties": {
        "named_card": {"type": "boolean"},
        "buy_list": {"type": "boolean"},
        "house_format": {"type": "boolean"},
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

Add no clause that only repeats a value the user already gave. "Which commander do you want for your white-black lifegain deck? Name one, or I suggest three." tells the user nothing they did not write themselves. Ask "Which commander do you want? Name one, or I suggest three." instead.

Keep a clause that narrows the question. "Do you have a red-green commander in mind?" tells the user which commanders you will accept, so it earns its words. The test is whether the clause changes what a useful answer looks like.

State no fact about the game. Do not say which colors, cards, or archetypes are strongest. Another step owns that, and a wrong claim costs the user's trust.

In Commander, the color identity of the commander is the color identity of the deck. Never offer to go beyond it, outside it, or to add a color to it. Gate run 14 asked "Do you want to use any colors beyond Grist's color identity?", and the rules do not allow that answer.

Presume nothing the user did not write. Do not say "your table", "your playgroup", or "your event" unless the user named one. A deck can be a gift.

Never name one card's color identity. Do not write "within Grist's color identity", "beyond Grist's color identity", or any clause of that shape. You are asked about the colors only when no commander is settled, so naming a card presumes that the card leads the deck, and the user may not have said so. Ask "Do you have a color preference?" instead.

Never name a format this app does not build: Brawl, Oathbreaker, Duel Commander, Canadian Highlander, Alchemy, Historic, or Timeless. The agent declines such a format in its own row, and every other question must leave it out. Gate run 15 asked "What should the Oathbreaker deck focus on?" one line under "I do not build Oathbreaker", so the agent contradicted itself in one message. Write "the deck", and never the declined format.

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

// The gap score has its own classifier call (D-69). A scorer that also
// phrases the question rates its own work. The cost is one more call
// per turn.

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
