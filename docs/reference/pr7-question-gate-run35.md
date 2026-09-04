# PR-7 question gate

Run date: 2026-09-04. Conversations: 2026-08-31.

Verdict: PASS. 27 of 27 counted gate conversations used catalog questions only. The bar is 25. The set holds 30 gate conversations, and 3 of them start after a build and are not counted (A-9).

3 conversations after a build ran beside the count. They asked 0 questions, and the model offered 0 replacements.

77 probe conversations ran beside the gate. They asked 299 questions, and the model offered 4 replacements. A probe explores catalog coverage and does not move the verdict (D-96).

No conversation reached a dead end (D-357).

16 turns moved nothing while a question was out, and a later turn recovered each one. A turn that repeats here is a catalog or classifier candidate: extra turns commander turn 4 (budget); land destruction turn 4 (budget); poison in a tournament turn 2 (budget, colors); the user contradicts the budget turn 2 (colors); sideboard help only turn 3 (budget); copy a known deck turn 2 (budget, colors); the user answers a different question turn 2 (budget, colors, format); the user answers a different question turn 4 (budget, commander); terse: a competitive request turn 3 (colors); terse: sideboard help only turn 3 (budget); terse: a card the user does not own turn 4 (commander); terse: cEDH at bracket 5 turn 3 (theme); terse: all caps, no punctuation turn 3 (budget); terse: stax with no table named turn 3 (budget, power); terse: the user answers with a number turn 4 (budget, commander); a set that resolves to one family turn 3 (budget).

The net of D-351 held back on 67 turns, because every question that was out is still inside its grace period. The reader answers it next turn, or the net closes it on the turn after. StallGrace is 2 turns: lifegain with a collection turn 3 (commander_pick); fun and janky turn 3 (commander); 60-card anything goes turn 3 (budget); the strongest deck, no collection turn 2 (colors); FNM on Friday turn 3 (colors, format); a commander the library does not hold turn 3 (budget); a card to keep that is not the commander turn 2 (pool_rule); standard at the store, no library turn 3 (budget); the strongest modern deck turn 2 (colors); pauper burn on a budget turn 2 (budget, format); vintage with proxies turn 3 (format); extra turns commander turn 4 (budget); land destruction turn 4 (budget); poison in a tournament turn 2 (budget, colors); legacy for an event turn 2 (budget, format); a card for the 99 turn 4 (budget); reanimator with two plans turn 4 (budget); two hundred dollars turn 3 (budget_scope); FNM on Friday, Modern turn 2 (budget, colors); burn on a budget turn 2 (budget, power); the user changes the theme late turn 3 (power); the user contradicts the budget turn 2 (colors); shorthand and slang turn 2 (budget, commander_pick); a commander that can not lead turn 3 (budget); two partner commanders turn 2 (budget, power); a deck as a gift turn 3 (commander); sideboard help only turn 2 (budget); copy a known deck turn 2 (budget, colors); a format we do not support turn 4 (commander); a price cap per card turn 2 (power); the user swaps the commander turn 2 (budget, power); two decks at once turn 3 (commander_pick); the user answers a different question turn 2 (budget, colors, format); the user answers a different question turn 4 (commander); terse: a card for the 99 turn 4 (budget); terse: two decks at once turn 3 (commander_pick); terse: a format we do not build turn 3 (commander); terse: a proxy table turn 2 (colors, format, house_format_limits); terse: a competitive request turn 2 (colors); terse: a deck as a gift turn 3 (commander_pick); terse: a named commander and a locked card turn 2 (pool_rule); terse: sideboard help only turn 2 (budget); terse: a Pauper deck named once turn 2 (format); terse: a five-color deck turn 2 (budget, power); terse: a colorless deck turn 2 (budget, power); terse: a card the user does not own turn 3 (commander); terse: cEDH at bracket 5 turn 2 (theme); terse: a companion turn 2 (budget, power); terse: a Background pair turn 3 (commander_pick); terse: partners named once turn 2 (budget, power); terse: all caps, no punctuation turn 2 (budget); terse: voltron turn 2 (budget, power); terse: chaos turn 2 (budget); terse: teaching a new player turn 3 (commander_pick); terse: a deck for a spouse turn 3 (commander_pick); terse: a budget of zero turn 3 (commander, pool_rule); terse: two budget shapes at once turn 2 (power); terse: a collection of lands only turn 3 (commander_pick); terse: a library that does not exist turn 3 (budget, commander_pick); terse: Oathbreaker turn 3 (commander); terse: mill with no table named turn 2 (budget); terse: stax with no table named turn 2 (budget, power); terse: two themes at once turn 2 (budget); terse: the user answers with a number turn 4 (commander); a theme the library can not lead turn 2 (commander_pick, power); a set that resolves to one family turn 3 (budget); a set name that names two base sets turn 3 (budget, colors, commander).

The net of D-351 healed 36 turns. Each one asked nothing new and was not ready, so it closed the questions that were out and built with what it had. agentsvc.Chat runs the same net on every turn. A turn here is a catalog or classifier candidate, because the reader answered nothing the agent could read: lifegain with a collection turn 3 (budget); blink with a thin library turn 3 (budget); fun and janky turn 3 (colors); the strongest deck, no collection turn 3 (colors); mill for a playgroup turn 3 (budget_scope); the user says none, then picks turn 4 (budget); the strongest modern deck turn 3 (colors); pauper burn on a budget turn 3 (format); a big library, no theme turn 3 (budget); poison in a tournament turn 3 (colors); legacy for an event turn 3 (format); FNM on Friday, Modern turn 3 (colors); no ban list at our table turn 3 (colors); the user changes the theme late turn 3 (budget); the user contradicts the budget turn 3 (colors); shorthand and slang turn 3 (commander_pick); sideboard help only turn 3 (budget); copy a known deck turn 3 (colors); a format we do not support turn 4 (budget); the user swaps the commander turn 4 (budget); the user answers a different question turn 4 (budget); terse: a proxy table turn 3 (colors, format); terse: a competitive request turn 3 (colors); terse: sideboard help only turn 3 (budget); terse: one-word answers turn 3 (budget); terse: a card the user does not own turn 4 (commander); terse: cEDH at bracket 5 turn 3 (theme); terse: all caps, no punctuation turn 3 (budget); terse: the user corrects the colors turn 3 (budget); terse: enchantress turn 3 (budget); terse: landfall turn 3 (budget); terse: Pauper Commander turn 3 (format); terse: stax with no table named turn 3 (budget, power); terse: the user answers with a number turn 4 (budget); a set name that names two base sets turn 4 (budget, commander); a set beside a theme turn 3 (budget).

No conversation called itself complete with a slot unanswered.

Every slot of the 27 counted conversations that name expectations ended on the expected value (PR-15).

The linter found no defective question (D-115).

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 27 |
| Questions asked | 118 |
| From the catalog | 118 |
| Invented by the model | 0 |
| Replacements refused as rewords | 1 |
| Catalog questions that closed a slot | 95 |
| Invented questions that closed a slot | 0 |
| Catalog-only conversations | 27 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

## Run

- Suite `questions`, run `pr7-question-gate-run35`, on 2026-09-04, commit `ec058b4`.
- Roles: ask on `gpt-5.6-luna` (openai, effort low), classify on `gpt-5.6-luna` (openai, effort none).
- Versions: card snapshot 2026-09-03, questions prompt version 17, conversations `2026-08-31`, slots_snapshot `4`.
- Calls: 666. Cost: $0.1860. Time: 1142 seconds.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Tokens: 1010761 input (522972 cached), 64978 output.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 5. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander, and white and black is right.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is the core level, near a precon; 3 is upgraded; 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.
  - Options: Astarion, the Decadent / Karlov of the Ghost Council / Sorin of House Markov // Sorin, Ravenous Neonate / None, name three more

**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.


**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.


### 2. blink with a thin library

Collection: true. Catalog: 5. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

Expected slots: every one met.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should I build the blink deck for: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the blink deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target: 1 exhibition, 2 core, 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.


### 3. a named card, role unknown

Collection: true. Catalog: 5. Invented: 0.

Expected slots: every one met.

Slots the deck needs and nobody answered: theme (asked, no answer).

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck focus on: a creature type, a mechanic, or a play style?

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=commander row=named_card_role fit=0.90 filled=true] Should Grist, the Hunger Tide lead the deck, or be included among the other cards?
  - Options: As my commander / In the 99
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

### 4. fun and janky

Collection: true. Catalog: 5. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.


**Turn 4, the user:** You pick the commander, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

### 5. 60-card anything goes

Collection: false. Catalog: 6. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=format row=format fit=0.05 filled=true] Which format should the 60-card deck use: Standard or Modern?
  - Options: Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=house_rules row=house_format_limits fit=0.35 filled=true] Inside your house format, do the normal 60-card deck limits hold?
  - Options: Yes, the normal limits / No, I will say the limits

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the dragon deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** The normal limits hold. Casual power, red and green.


### 6. the strongest deck, no collection

Collection: false. Catalog: 3. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.


### 7. FNM on Friday (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (asked, no answer).

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=false] Which format does Friday's FNM use: Standard, Modern, or Commander?
  - Options: Standard / Modern / Commander
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Pioneer. The best deck under budget.

- [catalog slot=format row=format_unsupported_open fit=0.35 filled=false] I do not build Pioneer. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


### 8. mill for a playgroup

Collection: true. Catalog: 6. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

Expected slots: every one met.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=0.90 filled=true] Which format do you want: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Does your $100 budget cover the cards you buy, or the whole deck's value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=power row=power_commander fit=0.90 filled=true] What power bracket should the deck target: 2 for core, near a preconstructed deck; 3 for upgraded; or 4 for high power? You can also choose 1 or 5.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.


### 9. a commander the library does not hold

Collection: true. Catalog: 4. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=commander row=named_card_role fit=0.90 filled=true] Do you want Brago to lead the deck, or should he be one card among the other 99?
  - Options: As my commander / In the 99
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck aim for?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.

- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Add him to the buy list. Bracket 3, white and blue.


**Turn 4, the user:** Under 30 dollars, and build from my library first.


### 10. no strong commander in the library

Collection: true. Catalog: 5. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the lifegain deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the lifegain deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander. White and black.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.


### 11. another version after a build (after a build)

Collection: true. Catalog: 0. Invented: 0.

The slots were full after turn 1. The last 1 messages never went out.

**This conversation asked nothing.** One message filled every slot, so it tests no question.

**Turn 1, the user:** Give me another version of that deck.


### 12. a slot change after a build (after a build)

Collection: true. Catalog: 0. Invented: 0.

**This conversation asked nothing.** One message filled every slot, so it tests no question.

**Turn 1, the user:** Switch to owned-only.


### 13. a card to keep that is not the commander

Collection: true. Catalog: 3. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck target? Choose from near-precon, upgraded, high-power, or cEDH levels.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should the deck use your library first, use only cards from your library, or ignore your library for a fully optimized build?
  - Options: Library first / Only my library / Ignore my library

**Turn 2, the user:** Commander, white and black, bracket 3.


**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

### 14. the user says none, then picks

Collection: false. Catalog: 5. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the lifegain deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the lifegain deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.

- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Vito, Thorn of the Dusk Rose, Heliod, Sun-Crowned, or Haliya, Guided by Light? Say "none" and I name three more.
  - Options: Vito, Thorn of the Dusk Rose / Heliod, Sun-Crowned / Haliya, Guided by Light / None, name three more

**Turn 3, the user:** None of those.

- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Aerith Gainsborough, Dina, Soul Steeper, or Celestine, the Living Saint? Say "none" and I name three more.
  - Options: Aerith Gainsborough / Dina, Soul Steeper / Celestine, the Living Saint / None, name three more

**Turn 4, the user:** The first of the new three is good.


### 15. standard at the store, no library

Collection: false. Catalog: 4. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Standard. An aggro deck, red and white.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** FNM level.


### 16. the strongest modern deck

Collection: false. Catalog: 1. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference for the deck?

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.


### 17. precon upgrade at bracket 2

Collection: true. Catalog: 2. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?

**Turn 2, the user:** A superfriends deck. Build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

### 18. pauper burn on a budget (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A Pauper burn deck, as cheap as possible.

- [catalog slot=format row=format_unsupported_open fit=0.35 filled=false] I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=budget row=budget fit=0.90 filled=true] How low should I keep the cost for cards you do not already own?

**Turn 2, the user:** Pauper. Burn, mono red.


**Turn 3, the user:** Casual power, and 25 dollars is the cap.


### 19. a big library, no theme

Collection: true. Catalog: 5. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** I have a big collection. Build me something good.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like to build for: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander. A token deck, green and white.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Trostani, Selesnya's Voice, Rhys the Redeemed, or Cadira, Caller of the Small? Say "none" and I name three more.
  - Options: Trostani, Selesnya's Voice / Rhys the Redeemed / Cadira, Caller of the Small / None, name three more

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


### 20. dinosaur tribal for a child

Collection: true. Catalog: 5. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the dinosaur deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.

- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Pantlaza, Sun-Favored, Itzquinth, Firstborn of Gishath, or Gishath, Sun's Avatar? Say "none" and I name three more.
  - Options: Pantlaza, Sun-Favored / Itzquinth, Firstborn of Gishath / Gishath, Sun's Avatar / None, name three more

**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [catalog slot=pool_rule row=pool_thin fit=0.90 filled=true] You have four Dinosaur cards, and the deck needs at least 30. Should I start with the cards you own and add a buy list, or use the entire pool?
  - Options: Owned-first with a buy list / Use the whole pool

**Turn 4, the user:** Build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you don’t own?

### 21. vintage with proxies (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer).

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=format row=format_unsupported_open fit=0.35 filled=false] I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=house_rules row=house_format_limits fit=0.35 filled=true] Inside your house format, do the normal 60-card deck limits hold?
  - Options: Yes, the normal limits / No, I will say the limits
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the artifact prison deck?

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 6. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the extra-turns deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? Choose from 1 (exhibition), 2 (core), 3 (upgraded), 4 (optimized), or 5 (cEDH).
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool_thin fit=0.90 filled=true] Since the card pool is thin for this theme, should I build around the cards you own and provide a buy list, or use the whole pool?
  - Options: Owned-first with a buy list / Use the whole pool
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 4, the user:** My table is fine with extra turns.


### 23. land destruction

Collection: true. Catalog: 6. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the land destruction deck target? Choose from exhibition, core, upgraded, optimized, or cEDH.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Syr Konrad, the Grim, Etali, Primal Storm, or Ragavan, Nimble Pilferer? Say "none" and I name three more.
  - Options: Syr Konrad, the Grim / Etali, Primal Storm / Ragavan, Nimble Pilferer / None, name three more

**Turn 2, the user:** Commander. Red and green.

- [catalog slot=pool_rule row=pool_thin fit=0.90 filled=true] Your library is small for land destruction. Should I use owned cards first and make a buy list, or use the entire pool?
  - Options: Owned-first with a buy list / Use the whole pool
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Etali, Primal Conqueror // Etali, Primal Sickness, Omnath, Locus of Rage, or Xenagos, God of Revels? Say "none" and I name three more.
  - Options: Etali, Primal Conqueror // Etali, Primal Sickness / Omnath, Locus of Rage / Xenagos, God of Revels / None, name three more

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 4, the user:** My table does not mind.


### 24. stax, owned only

Collection: true. Catalog: 4. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

Expected slots: every one met.

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.


### 25. poison in a tournament

Collection: false. Catalog: 2. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.


### 26. legacy for an event (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Legacy deck for an event.

- [catalog slot=format row=format_unsupported_open fit=0.35 filled=false] I do not build Legacy. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.


**Turn 3, the user:** FNM level, 600 dollars.


### 27. a card for the 99

Collection: true. Catalog: 6. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: feature a creature type, use a mechanic, or follow a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [catalog slot=pool_rule row=pool fit=0.20 filled=true] Should I build from your library first, use only cards from your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library
  - Refused as a reword (D-88): Do you want to build from your library first, only your library, or ignore your library for a fully optimized deck?
  - The guard compared against: Build from your library first, only your library, or ignore it for a fully optimized deck?
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 4, the user:** You pick the commander, and keep Grist in the deck.


### 28. reanimator with two plans

Collection: true. Catalog: 6. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is the core level, near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Syr Konrad, the Grim, Sheoldred, Whispering One, or Junji, the Midnight Sky? Say "none" and I name three more.
  - Options: Syr Konrad, the Grim / Sheoldred, Whispering One / Junji, the Midnight Sky / None, name three more

**Turn 2, the user:** Commander. Black and green.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards from your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Meren of Clan Nel Toth, Grist, Voracious Larva // Grist, the Plague Swarm, or Zask, Skittering Swarmlord? Say "none" and I name three more.
  - Options: Meren of Clan Nel Toth / Grist, Voracious Larva // Grist, the Plague Swarm / Zask, Skittering Swarmlord / None, name three more

**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 4, the user:** Reanimate one big creature, not a value loop.


### 29. two hundred dollars

Collection: true. Catalog: 5. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. An artifact deck, blue and red.

- [catalog slot=budget row=budget_scope fit=0.90 filled=true] Does your budget apply to the cards you buy, or to the deck's total value?
  - Options: The cards I buy / The whole deck
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** The 200 caps the cards I buy.


### 30. another version, same plan (after a build)

Collection: true. Catalog: 0. Invented: 0.

The slots were full after turn 1. The last 1 messages never went out.

**This conversation asked nothing.** One message filled every slot, so it tests no question.

**Turn 1, the user:** Give me the same deck with other cards.


### 31. FNM on Friday, Modern

Collection: false. Catalog: 3. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=true] Which format does Friday’s event run: Standard, Modern, or Commander?
  - Options: Standard / Modern / Commander
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


### 32. burn on a budget

Collection: false. Catalog: 2. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** A Standard burn deck, as cheap as possible.

- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Standard. Burn, mono red.


**Turn 3, the user:** Casual power, and 25 dollars is the cap.


### 33. a tempo deck for an event

Collection: false. Catalog: 4. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** A Modern deck for an event.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the Modern deck do: use a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Modern. A Delver of Secrets tempo deck, blue and red.

- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** FNM level, 600 dollars.


### 34. no ban list at our table

Collection: false. Catalog: 5. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** Any card, no ban list. Call it Modern. An artifact prison deck.

- [catalog slot=house_rules row=house_format_limits fit=0.35 filled=true] Inside your house format, do the normal 60-card deck limits hold?
  - Options: Yes, the normal limits / No, I will say the limits
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 35. the user changes the format (probe)

Collection: false. Catalog: 7. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander, white and black.

- [catalog slot=power row=power_commander fit=0.90 filled=false] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Actually, make it Modern instead. Sixty cards.

- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 4, the user:** FNM level, and 100 dollars.


### 36. the user changes the theme late (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A sacrifice deck from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the sacrifice deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the sacrifice deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander, black and green. Meren of Clan Nel Toth is the commander.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** I changed my mind, I want a token deck instead.


**Turn 4, the user:** Bracket 3, and build from my library first.


### 37. the user contradicts the budget (probe)

Collection: false. Catalog: 1. Invented: 0.

**Turn 1, the user:** A competitive Modern deck, my budget is 50 dollars.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference for the deck?

**Turn 2, the user:** Modern. Actually money is no object, spend what you need.


**Turn 3, the user:** Tournament level. My shop plays aggro.


### 38. everything in one message (probe)

Collection: true. Catalog: 0. Invented: 0.

**This conversation asked nothing.** One message filled every slot, so it tests no question.

**Turn 1, the user:** Commander, bracket 3, Karlov of the Ghost Council, white and black lifegain, build from my library first, 50 dollar buy list.


### 39. the user stays vague (probe)

Collection: true. Catalog: 5. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** Make me a good deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format do you want: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** I dunno, you pick.

- [catalog slot=power row=power_commander fit=0.90 filled=true] What power bracket should I aim for: 1 exhibition, 2 core, 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool_thin fit=0.90 filled=true] Should I prioritize the cards you already own and add a buy list, or use the entire pool?
  - Options: Owned-first with a buy list / Use the whole pool
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Syr Konrad, the Grim, Etali, Primal Storm, or Ragavan, Nimble Pilferer? Say "none" and I name three more.
  - Options: Syr Konrad, the Grim / Etali, Primal Storm / Ragavan, Nimble Pilferer / None, name three more

**Turn 3, the user:** Whatever you think is best.


### 40. the user asks a question back (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want a Commander deck.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: use a creature type, mechanic, or play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** What is a bracket? I have never heard of that.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Alright, bracket 2 then. A dragon deck, red and green.

- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

### 41. a request we can not serve (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** Can you build me a Yu-Gi-Oh deck?

- [catalog slot=scope row=out_of_scope fit=0.35 filled=true] I build Magic: The Gathering decks only. Would you like one instead?
  - Options: Yes, a Magic deck / No thanks

**Turn 2, the user:** Fine, Magic then. Commander, a dragon deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck aim for? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the dragon deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Bracket 3, red and green.

- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Atarka, World Render, Thrakkus the Butcher, or Dragonlord Atarka? Say "none" and I name three more.
  - Options: Atarka, World Render / Thrakkus the Butcher / Dragonlord Atarka / None, name three more

### 42. shorthand and slang (probe)

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** edh gruul dino stompy pls, no proxies

- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the Dinosaur deck target? Choose from near-precon to cEDH.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Itzquinth, Firstborn of Gishath, Owen Grady, Raptor Trainer, or Etali, Primal Conqueror // Etali, Primal Sickness? Say "none" and I name three more.
  - Options: Itzquinth, Firstborn of Gishath / Owen Grady, Raptor Trainer / Etali, Primal Conqueror // Etali, Primal Sickness / None, name three more

**Turn 2, the user:** bracket 3, from my binder


**Turn 3, the user:** yeah 50 bucks max


### 43. a card name with a typo (probe)

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** Build around Karlov of the Ghost Counsel.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck? Choose Commander, Standard, or Modern.
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander, white and black lifegain.

- [catalog slot=commander row=named_card_role fit=0.90 filled=false] Should Karlov of the Ghost Counsel lead the deck, or should it be included among the other cards?
  - Options: As my commander / In the 99
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck target: 1 exhibition, 2 core and near a precon, 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I use cards from your library first, restrict the deck to your library, or ignore your library for a fully optimized build?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Yes that is the card I meant. Bracket 3, from my library.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

### 44. a card that is not legal (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** I want a Modern deck built around Black Lotus.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Then pick something else. A burn deck, mono red.

- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** FNM level, 200 dollars.


### 45. a commander that can not lead (probe)

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** Commander deck with Lightning Bolt as my commander.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=commander row=commander_illegal fit=0.35 filled=true] Lightning Bolt can not lead a deck. Shall I suggest a commander instead?
  - Options: Yes, suggest one / No, I will name another
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** Alright, suggest a red commander for a burn deck.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Purphoros, God of the Forge, Ashling, Flame Dancer, or Dragonhawk, Fate's Tempest? Say "none" and I name three more.
  - Options: Purphoros, God of the Forge / Ashling, Flame Dancer / Dragonhawk, Fate's Tempest / None, name three more

**Turn 3, the user:** The first one. Bracket 3.


### 46. two partner commanders (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Commander deck with Thrasios, Triton Hero and Tymna the Weaver.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Yes, both as partners. A value deck.


**Turn 3, the user:** Bracket 5, cEDH. Money is no object.


### 47. a deck as a gift (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want to build a deck as a gift for my brother. He likes zombies.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the zombie deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not already own?

**Turn 2, the user:** Commander, black. He is new to the game.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a preconstructed deck, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Bracket 2, and under 80 dollars.


### 48. sideboard help only (probe)

Collection: false. Catalog: 1. Invented: 1.

**Turn 1, the user:** I already have a Modern burn deck. I only need sideboard help.

- [INVENTED slot=colors row=colors fit=0.20 filled=true] What colors should the sideboard use?
  - It replaced: Any color preference?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Modern, burn, mono red. Tournament level.


**Turn 3, the user:** My shop plays a lot of control.


### 49. copy a known deck (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** Build me the current best Standard deck, whatever is winning.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Standard. Tournament level.


**Turn 3, the user:** 300 dollars, and my shop plays aggro.


### 50. a format we do not support (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want a Brawl deck for Arena.

- [catalog slot=format row=format_unsupported fit=0.35 filled=true] I do not build Brawl. The nearest format I build is Commander. Shall I use that?
  - Options: Yes, use the nearest format / No, I will name another format
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** It is like Commander but 60 cards on Arena.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Fine, treat it as Commander. A dragon deck, red.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the dragon deck target? 2 is near a preconstructed deck, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 4, the user:** Bracket 2.


### 51. the user declines everything (probe)

Collection: true. Catalog: 5. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** Build me a Commander deck.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: use a creature type, mechanic, or play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** You decide.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should the deck use your library first, use only your library, or ignore it for a fully optimized build?
  - Options: Library first / Only my library / Ignore my library
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Syr Konrad, the Grim, Etali, Primal Storm, or Ragavan, Nimble Pilferer? Say "none" and I name three more.
  - Options: Syr Konrad, the Grim / Etali, Primal Storm / Ragavan, Nimble Pilferer / None, name three more

**Turn 3, the user:** You decide.


### 52. a price cap per card (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Modern deck where nothing costs more than 5 dollars a card.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Modern, a burn deck, mono red.


**Turn 3, the user:** Casual power, and 100 dollars total.


### 53. the user swaps the commander (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A lifegain Commander deck from my library.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? Level 2 is close to a preconstructed deck, level 3 is upgraded, and level 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] Should I follow a spending limit for cards you do not already own?

**Turn 2, the user:** White and black. Karlov of the Ghost Council is the commander.


**Turn 3, the user:** Actually use a different commander, suggest one.

- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.
  - Options: Astarion, the Decadent / Karlov of the Ghost Council / Sorin of House Markov // Sorin, Ravenous Neonate / None, name three more

**Turn 4, the user:** The second one. Bracket 3, library first.


### 54. two decks at once (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** I want two decks, one Commander and one Modern.

- [catalog slot=deck_count row=one_deck fit=0.35 filled=false] I build one deck at a time. Which deck do you want first?

**Turn 2, the user:** Start with the Commander one. A dragon deck, red and green.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the red-green dragon deck target? 2 is the core level, near a precon; 3 is upgraded; 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Atarka, World Render, Thrakkus the Butcher, or Dragonlord Atarka? Say "none" and I name three more.
  - Options: Atarka, World Render / Thrakkus the Butcher / Dragonlord Atarka / None, name three more

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 55. a returning player (probe)

Collection: false. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I have not played since 2015 and I want to get back in.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander, I think. I like big green creatures.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? Level 2 is near a precon, level 3 is upgraded, and level 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Green. Bracket 2, and about 100 dollars.

- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Freyalise, Llanowar's Fury, Greensleeves, Maro-Sorcerer, or Renata, Called to the Hunt? Say "none" and I name three more.
  - Options: Freyalise, Llanowar's Fury / Greensleeves, Maro-Sorcerer / Renata, Called to the Hunt / None, name three more

### 56. the user answers a different question (probe)

Collection: true. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A mill deck from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the mill deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the mill deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** My friends hate mill.


**Turn 3, the user:** Commander, blue and black. Bracket 3.

- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 4, the user:** Build from my library first.


### 57. terse: the format is named once (probe)

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target: 2 for core, near a precon; 3 for upgraded; 4 for high power; or another level?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Syr Konrad, the Grim, Etali, Primal Storm, or Ragavan, Nimble Pilferer? Say "none" and I name three more.
  - Options: Syr Konrad, the Grim / Etali, Primal Storm / Ragavan, Nimble Pilferer / None, name three more

**Turn 2, the user:** Red and green.

- [catalog slot=pool_rule row=pool_thin fit=0.90 filled=true] Your pool is thin for this theme. Should I start with cards you own and add a buy list, or use the full pool?
  - Options: Owned-first with a buy list / Use the whole pool
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Etali, Primal Conqueror // Etali, Primal Sickness, Omnath, Locus of Rage, or Xenagos, God of Revels? Say "none" and I name three more.
  - Options: Etali, Primal Conqueror // Etali, Primal Sickness / Omnath, Locus of Rage / Xenagos, God of Revels / None, name three more

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

### 58. terse: a card for the 99 (probe)

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** A sacrifice deck, black and green.

- [catalog slot=pool_rule row=pool fit=0.20 filled=true] Should I build the sacrifice deck from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library
  - Refused as a reword (D-88): Build from your library first, only your library, or use any cards for a fully optimized deck?
  - The guard compared against: Build from your library first, only your library, or ignore it for a fully optimized deck?
- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you don’t own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Dina, Essence Brewer, Jarad, Golgari Lich Lord, or Dina, Soul Steeper? Say "none" and I name three more.
  - Options: Dina, Essence Brewer / Jarad, Golgari Lich Lord / Dina, Soul Steeper / None, name three more

**Turn 4, the user:** You pick the commander.


### 59. terse: two decks at once (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** I want two decks, one Commander and one Modern.

- [catalog slot=deck_count row=one_deck fit=0.35 filled=true] I build one deck at a time. Which deck do you want first?

**Turn 2, the user:** The Commander one. A dragon deck, red and green.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? 2 is near a preconstructed deck, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Atarka, World Render, Thrakkus the Butcher, or Dragonlord Atarka? Say "none" and I name three more.
  - Options: Atarka, World Render / Thrakkus the Butcher / Dragonlord Atarka / None, name three more

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 60. terse: a format we do not build (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want a Brawl deck for Arena.

- [catalog slot=format row=format_unsupported fit=0.35 filled=true] I do not build Brawl. The nearest format I build is Commander. Shall I use that?
  - Options: Yes, use the nearest format / No, I will name another format
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander is fine. A dragon deck, red.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the dragon deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Bracket 2.


### 61. terse: a proxy table (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** We proxy everything at our table. A Vintage artifact prison deck.

- [catalog slot=format row=format_unsupported_open fit=0.35 filled=false] I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=house_rules row=house_format_limits fit=0.35 filled=true] Inside your house format, do the normal 60-card deck limits hold?
  - Options: Yes, the normal limits / No, I will say the limits
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference for the deck?

**Turn 2, the user:** Any card, no ban list.


**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 62. terse: a precon upgrade (probe)

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
  - Options: Creature type / Mechanic / Play style

**Turn 2, the user:** A superfriends deck.

- [catalog slot=pool_rule row=pool_precon fit=0.90 filled=true] Should I start with the precon, use only cards from it, or ignore it for a fully optimized deck?
  - Options: The precon first / Only the precon / Ignore the precon

**Turn 3, the user:** Build from the precon first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

### 63. terse: a competitive request (probe)

Collection: false. Catalog: 1. Invented: 0.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference for the deck?

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Yes, tournament level. My shop plays aggro and combo.


### 64. terse: a deck as a gift (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** I want to build a Commander deck as a gift for my brother. He likes zombies.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a preconstructed deck, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Black. He is new to the game.

- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Gisa, the Hellraiser, Zul Ashur, Lich Lord, or Hancock, Ghoulish Mayor? Say "none" and I name three more.
  - Options: Gisa, the Hellraiser / Zul Ashur, Lich Lord / Hancock, Ghoulish Mayor / None, name three more

**Turn 3, the user:** Bracket 2, and under 80 dollars.


### 65. terse: a named commander and a locked card (probe)

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? Choose core, near a precon; upgraded; high power; optimized; or cEDH.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should the deck use your library first, only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 2, the user:** Bracket 3.


**Turn 3, the user:** Build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you don’t own?

### 66. terse: sideboard help only (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** I already have a Modern burn deck, mono red. I only need sideboard help.

- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Tournament level.


**Turn 3, the user:** My shop plays a lot of control.


### 68. terse: the strongest deck, no collection (probe)

Collection: false. Catalog: 2. Invented: 0.

The slots were full after turn 2. The last 1 messages never went out.

**Turn 1, the user:** Build the strongest Modern deck possible. I own nothing.

- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** A burn deck, mono red.


### 69. terse: a Pauper deck named once (probe)

Collection: false. Catalog: 2. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer).

**Turn 1, the user:** A Pauper burn deck, as cheap as possible. Mono red.

- [catalog slot=format row=format_unsupported_open fit=0.35 filled=false] I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Casual power, and 25 dollars is the cap.


### 70. terse: a mill deck for a playgroup (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Commander mill deck for my playgroup, budget 100.

- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck aim for? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget_scope fit=0.90 filled=true] Does your 100 budget cover only the cards you buy, or the deck's total value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Blue and black. Bruvac the Grandiloquent is the commander.

- [catalog slot=pool_rule row=pool fit=0.05 filled=true] Should I build from your library first, use only cards from it, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library
  - Refused as a reword (D-88): Build from your library first, use only your library, or ignore it for a fully optimized deck?
  - The guard compared against: Build from your library first, only your library, or ignore it for a fully optimized deck?

**Turn 3, the user:** Bracket 3, build from my library first. The 100 caps the cards I buy.


### 71. terse: one-word answers (probe)

Collection: false. Catalog: 3. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** Mono-white weenie deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format do you want for the mono-white weenie deck?
  - Options: Commander / Standard / Modern
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Modern.

- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta

**Turn 3, the user:** FNM.


### 72. terse: a five-color deck (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A five-color Commander deck. Kenrith, the Returned King is the commander.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Good stuff, whatever is strongest.


**Turn 3, the user:** Bracket 4, and 400 dollars.


### 73. terse: a colorless deck (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A colorless Commander deck built around big artifacts.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Emry, Lurker of the Loch, Urza, Lord High Artificer, or Padeem, Consul of Innovation? Say "none" and I name three more.
  - Options: Emry, Lurker of the Loch / Urza, Lord High Artificer / Padeem, Consul of Innovation / None, name three more

**Turn 2, the user:** Kozilek, Butcher of Truth is the commander.


**Turn 3, the user:** Bracket 3, and 250 dollars.


### 74. terse: a card the user does not own (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Commander deck from only the cards I own. Sol Ring goes in it.

- [catalog slot=theme row=theme fit=0.90 filled=true] What would you like the deck to do—focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** An artifact deck, colorless.

- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Bracket 3.


**Turn 4, the user:** Buy it then.


### 75. terse: cEDH at bracket 5 (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A cEDH deck. Najeela, the Blade-Blossom.

- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Money is no object.


**Turn 3, the user:** Yes, bracket 5.


### 76. terse: a companion (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Modern deck with Lurrus of the Dream-Den as my companion.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Mono black, aggro.


**Turn 3, the user:** Tournament level, 300 dollars.


### 77. terse: a Background pair (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** A Commander deck with a Background commander pair.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the Background commander pair?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Red and white, aggressive.

- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Karlach, Fury of Avernus + Inspiring Leader, Karlach, Fury of Avernus + Folk Hero, or Karlach, Fury of Avernus + Noble Heritage? Say "none" and I name three more.
  - Options: Karlach, Fury of Avernus + Inspiring Leader / Karlach, Fury of Avernus + Folk Hero / Karlach, Fury of Avernus + Noble Heritage / None, name three more

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 78. terse: partners named once (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** Krark, the Thumbless and Sakashima of a Thousand Faces as partners.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** A spellslinger deck, blue and red.


**Turn 3, the user:** Bracket 4, and 500 dollars.


### 79. terse: everything buried in one long message (probe)

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: power (asked, no answer), commander (never asked).

**Turn 1, the user:** So I have been playing with my friends on Thursdays for about a year and we mostly play Commander, and I picked up a lot of black and green cards over that time, and what I really want is something that sacrifices creatures for value, nothing too strong because we are all pretty casual, maybe bracket 2 or 3, and I would rather use what I already have than buy much.

- [catalog slot=power row=power_commander fit=0.20 filled=false] Which power bracket should the deck target? You mentioned 2 or 3; choose one, or pick a different level.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=false] How much, if anything, would you like to spend on cards you do not already own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Amzu, Swarm's Hunger, Chatterfang, Squirrel General, or Grist, the Hunger Tide? Say "none" and I name three more.
  - Options: Amzu, Swarm's Hunger / Chatterfang, Squirrel General / Grist, the Hunger Tide / None, name three more

### 80. terse: all caps, no punctuation (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** BUILD ME A MODERN BURN DECK MONO RED

- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** TOURNAMENT


**Turn 3, the user:** MY SHOP PLAYS CONTROL


### 81. terse: the user corrects the colors (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Commander lifegain deck, white and black.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the lifegain deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Astarion, the Decadent, Karlov of the Ghost Council, or Sorin of House Markov // Sorin, Ravenous Neonate? Say "none" and I name three more.
  - Options: Astarion, the Decadent / Karlov of the Ghost Council / Sorin of House Markov // Sorin, Ravenous Neonate / None, name three more

**Turn 2, the user:** Actually make it white and green instead.

- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Lathiel, the Bounteous Dawn, Treebeard, Gracious Host, or Aerith, Last Ancient? Say "none" and I name three more.
  - Options: Lathiel, the Bounteous Dawn / Treebeard, Gracious Host / Aerith, Last Ancient / None, name three more

**Turn 3, the user:** You pick the commander. Bracket 3.


### 82. terse: the user corrects a card name (probe)

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** Build around Atraxa, Praetor's Voice.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the Atraxa, Praetors' Voice deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** I meant Atraxa, Praetors' Voice. Commander, superfriends.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? Choose 2 for near-precon, 3 for upgraded, 4 for high power, or another level if preferred.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Bracket 3, and 300 dollars.


### 83. terse: group hug (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A group hug Commander deck. Phelddagrif is the commander.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target: 2, near a precon; 3, upgraded; or 4, high power? You can also choose 1, exhibition, or 5, cEDH.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Bracket 2, and 100 dollars.


### 84. terse: voltron (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A voltron Commander deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** White and green. You pick the commander.


**Turn 3, the user:** The first one. Bracket 3, and 200 dollars.


### 85. terse: chaos (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A chaos Commander deck, as random as possible.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? 2 is near a preconstructed deck, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the chaos deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Yidris, Maelstrom Wielder. Bracket 3.


**Turn 3, the user:** 150 dollars.


### 86. terse: enchantress (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** An enchantress deck from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the enchantress deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the enchantress deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander, white and green. Sythis, Harvest's Hand.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, library first.


### 87. terse: landfall (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A landfall Commander deck from my binder.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Green, mono.

- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Loot, Exuberant Explorer, Bristly Bill, Spine Sower, or Lumra, Bellow of the Woods? Say "none" and I name three more.
  - Options: Loot, Exuberant Explorer / Bristly Bill, Spine Sower / Lumra, Bellow of the Woods / None, name three more

**Turn 3, the user:** You pick the commander. Bracket 3, library first.


### 88. terse: no theme at all (probe)

Collection: false. Catalog: 3. Invented: 0.

The slots were full after turn 2. The last 1 messages never went out.

**Turn 1, the user:** A Commander deck. Surprise me.

- [catalog slot=theme row=theme fit=0.90 filled=true] Should the deck center on a creature type, a mechanic, or a play style? You can also leave this open for a surprise.
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** You decide.


### 89. terse: teaching a new player (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** A Commander deck for someone who has never played before.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: use a creature type, a mechanic, or a particular play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target, from exhibition to cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Green, big creatures.

- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Marwyn, the Nurturer, Kogla, the Titan Ape, or Nissa, Resurgent Animist? Say "none" and I name three more.
  - Options: Marwyn, the Nurturer / Kogla, the Titan Ape / Nissa, Resurgent Animist / None, name three more

**Turn 3, the user:** Bracket 1, and under 50 dollars.


### 90. terse: a deck for a spouse (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** My wife likes cats. Build her a Commander deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** White and green. Bracket 2.

- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Arahbo, Roar of the World, Kaheera, the Orphanguard, or Kutzil, Malamet Exemplar? Say "none" and I name three more.
  - Options: Arahbo, Roar of the World / Kaheera, the Orphanguard / Kutzil, Malamet Exemplar / None, name three more

**Turn 3, the user:** Under 60 dollars.


### 91. terse: a budget of zero (probe)

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer), pool_rule (asked, no answer).

**Turn 1, the user:** A Commander deck and I can not spend anything.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a preconstructed deck, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Black and green, a sacrifice deck.

- [catalog slot=pool_rule row=pool fit=0.05 filled=false] Should the deck use your library first, use only cards from your library, or ignore your library for a fully optimized build?
  - Options: Library first / Only my library / Ignore my library
  - Refused as a reword (D-88): Build from your library first, use only your library, or ignore it for a fully optimized deck?
  - The guard compared against: Build from your library first, only your library, or ignore it for a fully optimized deck?
- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Bracket 2.


### 92. terse: proxy the expensive cards only (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Commander deck. I proxy anything over 20 dollars.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Blue and black, a mill deck. Bruvac the Grandiloquent.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I prioritize cards from your library, use only cards from it, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3, library first.


### 93. terse: two budget shapes at once (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Modern deck. Nothing over 5 dollars a card, and 120 dollars in total.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** A burn deck, mono red.


**Turn 3, the user:** FNM level.


### 94. terse: a collection of lands only (probe)

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** I only own lands. Build me a Commander deck anyway.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Red and green, a dinosaur deck.

- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Itzquinth, Firstborn of Gishath, Owen Grady, Raptor Trainer, or Etali, Primal Conqueror // Etali, Primal Sickness? Say "none" and I name three more.
  - Options: Itzquinth, Firstborn of Gishath / Owen Grady, Raptor Trainer / Etali, Primal Conqueror // Etali, Primal Sickness / None, name three more

**Turn 3, the user:** Bracket 3, and 200 dollars.


### 95. terse: a library that does not exist (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** Build a Commander deck from my library.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** White and blue, a blink deck.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander_pick fit=0.35 filled=false] Which commander do you want: Abuelo, Ancestral Echo, Oji, the Exquisite Blade, or Niko, Light of Hope? Say "none" and I name three more.
  - Options: Abuelo, Ancestral Echo / Oji, the Exquisite Blade / Niko, Light of Hope / None, name three more

**Turn 3, the user:** Bracket 3.


### 96. terse: the user asks about a card (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** Is Sol Ring legal in Modern?

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: follow a creature type, use a mechanic, or have a particular play style?
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Alright, a Modern artifact deck then, blue.

- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** FNM level, and 150 dollars.


### 97. terse: Pauper Commander (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Pauper Commander deck.

- [catalog slot=format row=format_unsupported_open fit=0.35 filled=false] I do not build Pauper Commander. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Black, a sacrifice deck.

- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Casual power, and 40 dollars.


### 98. terse: Oathbreaker (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** An Oathbreaker deck.

- [catalog slot=format row=format_unsupported fit=0.35 filled=true] I do not build Oathbreaker. The nearest format I build is Commander. Shall I use that?
  - Options: Yes, use the nearest format / No, I will name another format
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander is fine. Blue and red, a spellslinger deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 99. terse: Historic on Arena (probe)

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** A Historic deck for Arena.

- [catalog slot=format row=format_unsupported_open fit=0.35 filled=true] I do not build Historic. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Modern works. A mono-black aggro deck.

- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Tournament level, and 200 dollars.


### 100. terse: a deck for a team event (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** I need a Modern deck for a team event on Saturday.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: use a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.35 filled=true] How strong should this be: casual, Friday Night Magic (FNM) level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** My teammates play aggro and combo, so I take control.

- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Tournament level, 600 dollars, blue and white.


### 101. terse: mill with no table named (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Commander mill deck.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is the core level, near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Blue and black. You pick the commander.


**Turn 3, the user:** The second one. Bracket 3, and 100 dollars.


### 102. terse: stax with no table named (probe)

Collection: false. Catalog: 3. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A stax Commander deck.

- [catalog slot=power row=power_commander fit=0.90 filled=false] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** White and blue. You pick the commander.


**Turn 3, the user:** None of those.


### 103. terse: two themes at once (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Commander deck that does artifacts and lifegain.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=budget row=budget fit=0.90 filled=true] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** White and blue. You pick the commander.


**Turn 3, the user:** The first one. Bracket 3, and 200 dollars.


### 104. terse: the user answers with a number (probe)

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A Commander deck from my collection.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is core, near a precon; 3 is upgraded; 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** 3

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Black and green, a sacrifice deck.

- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 4, the user:** 2


### 105. a theme the library can not lead (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** Build me an Eldrazi Commander deck using only the cards in my library.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?
- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Lotho, Corrupt Shirriff, Ghalta, Primal Hunger, or Peregrin Took? Say "none" and I name three more.
  - Options: Lotho, Corrupt Shirriff / Ghalta, Primal Hunger / Peregrin Took / None, name three more

**Turn 2, the user:** Colorless.


**Turn 3, the user:** Bracket 3.

- [catalog slot=commander row=commander_pick fit=0.35 filled=true] Which commander do you want: Jodah, the Unifier, Leonardo, the Balance, or Cosmic Spider-Man? Say "none" and I name three more.
  - Options: Jodah, the Unifier / Leonardo, the Balance / Cosmic Spider-Man / None, name three more

**Turn 4, the user:** You pick the commander.


### 106. a set that resolves to one family (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** Build me a Commander deck only from the Hobbit set.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is near a precon, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Bracket 3, and you pick the commander.

- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 3, the user:** Any colors are fine.


### 107. a set name that names two base sets (probe)

Collection: false. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a Tarkir deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format do you want: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the Tarkir deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=set row=set_unresolved fit=0.35 filled=true] I do not know which set you mean by Tarkir. Which one: Tarkir: Dragonstorm, Dragons of Tarkir, and Khans of Tarkir?
  - Options: Tarkir: Dragonstorm / Dragons of Tarkir / Khans of Tarkir / Use every set

**Turn 2, the user:** Commander, dragons, bracket 3.

- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the dragon deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?
- [catalog slot=commander row=commander fit=0.35 filled=false] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Tarkir: Dragonstorm.


**Turn 4, the user:** No color preference.


### 108. a set beside a theme (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Bloomburrow deck about animals.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander, bracket 2.

- [catalog slot=commander row=commander fit=0.35 filled=true] Which commander do you want? Name one, or I suggest three.
  - Options: Suggest one

**Turn 3, the user:** Green and white, and you pick the commander.


