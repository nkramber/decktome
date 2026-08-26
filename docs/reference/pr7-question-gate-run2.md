# PR-7 question gate

Run date: 2026-08-25. Conversations: 2026-08-24.

Verdict: FAIL. 23 of 30 conversations used catalog questions only. The bar is 25.

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 91 |
| From the catalog | 84 |
| Invented by the model | 7 |
| Catalog questions that closed a slot | 66 |
| Invented questions that closed a slot | 3 |
| Catalog-only conversations | 23 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `colors` | 2 |
| `acquisition` | 1 |
| `commander_pick` | 1 |
| `house_format_limits` | 1 |
| `pool` | 1 |
| `power_commander` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 192. Time: 320.0 seconds.
- Tokens: 100936 input (0 cached), 17426 output.
- Cost: $0.0411.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 2. Invented: 0.

The slots were full after turn 2. The last 2 messages never went out.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.95 filled=true] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Commander, and white and black is right.


### 2. blink with a thin library

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.98 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.95 filled=true] Do you have a color preference? Blink is strongest in white, blue, black, and green.

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=1.00 filled=true] Do you have a white-blue commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.30 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.

- [catalog slot=budget row=budget fit=0.95 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** No more than 40 dollars.

- [catalog slot=budget row=budget_scope fit=0.05 filled=false] Is the $40 limit for the cards you buy, or for the whole deck’s value?
  - Options: The cards I buy / The whole deck

### 3. a named card, role unknown

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like for your Grist, the Hunger Tide deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=pool_rule row=pool fit=0.95 filled=true] For your Grist, the Hunger Tide Commander deck, should I build from your library first, use only cards in your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.


### 4. fun and janky

Collection: true. Catalog: 3. Invented: 1.

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=power row=jank fit=1.00 filled=false] What does “janky” mean to you: a low-power deck, an odd card nobody expects, or a plan that almost never works?
  - Options: A low-power deck / An odd card nobody expects / A plan that almost never works
- [catalog slot=colors row=colors fit=0.55 filled=true] Do you have a color preference? Fun, janky decks are especially strong in blue, black, and green.

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.


**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.

- [INVENTED slot=pool_rule row=pool fit=0.12 filled=true] Should I use cards from your existing collection, or build from the full card pool?
  - Options: Use my existing collection / Build from the full card pool

**Turn 4, the user:** You pick the commander, and build from my library first.


### 5. 60-card anything goes

Collection: false. Catalog: 3. Invented: 1.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=theme row=theme fit=1.00 filled=true] What should the deck do—focus on a creature type, mechanic, or play style?
- [catalog slot=house_rules row=house_rules fit=1.00 filled=true] When you say anything goes, do you mean any card with no ban list or Vintage rules?
  - Options: Any card, no ban list / Vintage rules
- [catalog slot=colors row=colors fit=1.00 filled=true] Do you have a color preference?

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [INVENTED slot=house_rules row=house_format_limits fit=0.25 filled=true] Should the Modern dragon deck use up to four copies of each card name and a 15-card sideboard?
  - Options: Yes / No

**Turn 3, the user:** The normal limits hold. Casual power, red and green.


### 6. the strongest deck, no collection

Collection: false. Catalog: 2. Invented: 1.

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [INVENTED slot=colors row=colors fit=0.25 filled=false] Which colors, if any, would you like the deck to use?
- [catalog slot=budget row=budget fit=1.00 filled=true] Since you own no cards, what budget do you have for buying the deck?

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.


### 7. FNM on Friday

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=1.00 filled=true] Which format is your Friday FNM: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
- [catalog slot=theme row=theme_competitive fit=0.80 filled=true] Would you like a named tier-one deck or the best competitive deck within your budget?
  - Options: A named deck / The best deck under budget
- [catalog slot=colors row=colors fit=0.55 filled=false] Do you have any color preference?

**Turn 2, the user:** Pioneer. The best deck under budget.


**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


**Turn 4, the user:** I buy online, and I need the cards by Thursday.

- [catalog slot=budget row=acquisition fit=1.00 filled=false] Since you buy online, do you need the cards by Thursday?
  - Options: Online by Thursday

### 8. mill for a playgroup

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.80 filled=true] Do you have a color preference? Mill is strongest in blue, black, and green.
- [catalog slot=budget row=budget_scope fit=0.90 filled=true] Is your $100 budget a cap on the cards you buy, or on the whole deck's value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=pool_rule row=pool fit=0.95 filled=true] For your blue-black Commander mill deck, should I use your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.


**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 1. Invented: 0.

The slots were full after turn 2. The last 2 messages never went out.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like for your Brago blink deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.


### 10. no strong commander in the library

Collection: true. Catalog: 2. Invented: 0.

The slots were full after turn 2. The last 1 messages never went out.

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.98 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Commander. White and black.


### 11. another version after a build

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=format row=format fit=0.85 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.95 filled=true] What should the deck focus on—a creature type, mechanic, or play style?
- [catalog slot=colors row=colors fit=0.75 filled=false] Do you have any color preference?

**Turn 2, the user:** A different plan in the same colors.


### 12. a frozen run asks nothing

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Switch to owned-only.

- [catalog slot=format row=format fit=0.05 filled=false] Which format should the owned-only deck be for: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.05 filled=false] What should the owned-only deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.05 filled=false] Do you have any color preference for the owned-only deck?

### 13. a card to keep that is not the commander

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=format row=format fit=0.98 filled=true] Which format should the Karlov of the Ghost Council lifegain deck use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, white and black, bracket 3.

- [catalog slot=pool_rule row=pool fit=0.92 filled=true] For your white-black Commander lifegain deck at bracket 3, should I use your library first, only your library, or ignore it for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library
- [catalog slot=locked row=locked fit=0.05 filled=true] Must the deck keep Sanguine Bond, or may I cut cards that do not fit the lifegain plan?
  - Options: Keep every card I named / Cut cards that do not fit

**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.


### 14. the user says none, then picks

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=format row=format fit=0.02 filled=false] Which format would you like for your lifegain deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.


**Turn 3, the user:** None of those.


**Turn 4, the user:** The first of the new three is good.


### 15. standard at the store, no library

Collection: false. Catalog: 2. Invented: 0.

The slots were full after turn 2. The last 1 messages never went out.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=theme row=theme fit=1.00 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any preferred colors?

**Turn 2, the user:** Standard. An aggro deck, red and white.


### 16. the strongest modern deck

Collection: false. Catalog: 1. Invented: 0.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=colors row=colors fit=0.05 filled=false] Do you have a color preference? The strongest Modern deck is blue, black, and green.

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.


### 17. precon upgrade at bracket 2

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=format row=format fit=0.05 filled=true] Which format are you playing: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** A superfriends deck. Build from my library first.


### 18. pauper burn on a budget

Collection: false. Catalog: 0. Invented: 0.

The slots were full after turn 1. The last 2 messages never went out.

**This conversation asked nothing.** One message filled every slot, so it tests no question.

**Turn 1, the user:** A Pauper burn deck, as cheap as possible.


### 19. a big library, no theme

Collection: true. Catalog: 5. Invented: 1.

**Turn 1, the user:** I have a big collection. Build me something good.

- [catalog slot=format row=format fit=0.98 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.93 filled=true] What should the deck focus on— a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. A token deck, green and white.

- [catalog slot=commander row=commander fit=0.98 filled=true] Do you have a commander in mind for your green-white token deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.55 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [INVENTED slot=commander row=commander_pick fit=0.12 filled=false] Which commander from your library would you like me to choose for the green-white token deck?

### 20. dinosaur tribal for a child

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.85 filled=true] Do you have any color preference? Dinosaurs are strongest in red and green.

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.


**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [catalog slot=pool_rule row=pool fit=0.97 filled=true] Should I build with your library as the first source, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 4, the user:** Build from my library first.


### 21. vintage with proxies

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.10 filled=true] Which format are you building for: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.15 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=house_rules row=house_rules fit=0.90 filled=true] When you say anything-goes, do you mean any card with no ban list or Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=house_rules row=house_format_limits fit=0.98 filled=true] For your house Vintage format, should the normal limits apply: a 60-card minimum, up to four copies of each card name, and a 15-card sideboard?
  - Options: Yes, the normal limits / No, I’ll specify the limits
- [catalog slot=colors row=colors fit=0.58 filled=true] Do you have a color preference? Artifact prison is often strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 3. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.95 filled=true] Do you have a color preference? Extra-turns decks are strongest in blue and green.

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=pool_rule row=pool fit=0.55 filled=true] For your blue-red Commander deck, should I build from your library first, use only cards in your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


### 23. land destruction

Collection: true. Catalog: 3. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=format row=format fit=0.05 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.85 filled=true] Do you have a color preference? Land destruction is strongest in blue, black, and green.

**Turn 2, the user:** Commander. Red and green.

- [catalog slot=pool_rule row=pool fit=0.95 filled=true] For your red-green Commander land-destruction deck, should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


### 24. stax, owned only

Collection: true. Catalog: 6. Invented: 1.

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.98 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.86 filled=true] Do you have a color preference? Stax is strongest in white, blue, and green.

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.95 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [INVENTED slot=power row=power_commander fit=0.25 filled=true] What power level should the white-blue stax deck target?

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.

- [catalog slot=commander row=commander_pick fit=0.02 filled=false] Which commander do you want: Giada, Font of Hope, Ioreth of the Healing House, or Gandalf, Wandering Wizard? You can say “none” and I’ll name three more.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more
- [catalog slot=power row=table_tolerance fit=0.18 filled=true] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts stax.

- [catalog slot=commander row=commander_pick fit=0.98 filled=false] Which commander do you want: Giada, Font of Hope, Ioreth of the Healing House, or Gandalf, Wandering Wizard? Say “none” and I’ll name three more.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more

### 25. poison in a tournament

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=colors row=colors fit=0.95 filled=false] Do you have a color preference? Infect is strongest in blue, black, and green.
  - Options: Blue / Black / Green / A combination of these / No preference

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.

- [catalog slot=plan_variant row=plan_choice fit=0.92 filled=false] For your Infect deck, which of the two common plans would you like to use?

### 26. legacy in three weeks

Collection: false. Catalog: 3. Invented: 1.

**Turn 1, the user:** A Legacy deck. I need it in three weeks for an event.

- [catalog slot=theme row=theme fit=0.82 filled=true] What should the Legacy deck do: focus on a creature type, a mechanic, or a particular play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences for the Legacy deck?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.

- [catalog slot=locked row=locked fit=0.95 filled=false] In this Legacy blue-red Delver of Secrets tempo deck, must you keep Delver of Secrets, or may I cut a card that does not fit the plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 3, the user:** FNM level, 600 dollars.


**Turn 4, the user:** I buy online, and I need the cards two days before the event.

- [INVENTED slot=budget row=acquisition fit=0.05 filled=false] What is your total budget for the deck?

### 27. a card for the 99

Collection: true. Catalog: 3. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=format row=format fit=0.95 filled=true] Which format should the Grist, the Hunger Tide deck be built for: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [catalog slot=pool_rule row=pool fit=0.05 filled=true] For your black-green Commander sacrifice deck, should I build from your library first, use only cards from your library, or ignore it for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library
- [catalog slot=locked row=locked fit=0.90 filled=true] Should the deck keep Grist, the Hunger Tide, or may I cut it if it does not fit the sacrifice plan?
  - Options: Keep every named card / Cut cards that do not fit

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.


### 28. reanimator with two plans

Collection: true. Catalog: 2. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=colors row=colors fit=0.96 filled=true] Do you have any color preferences? Reanimator is strongest in white, black, and green.

**Turn 2, the user:** Commander. Black and green.

- [catalog slot=pool_rule row=pool fit=0.95 filled=true] For your black-green Commander reanimator deck, should I build from your library first, use only cards in your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.


### 29. two hundred dollars

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=theme row=theme fit=1.00 filled=true] What should the Commander deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.85 filled=true] Do you have any color preference for the deck?
- [catalog slot=budget row=budget_scope fit=0.15 filled=true] Is your $200 budget a cap on the cards you buy or on the whole deck's value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. An artifact deck, blue and red.


**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** The 200 caps the cards I buy.


### 30. another version, same plan

Collection: true. Catalog: 2. Invented: 1.

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the deck be for: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.95 filled=false] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [INVENTED slot=colors row=colors fit=0.30 filled=false] Which colors should the replacement cards use?

**Turn 2, the user:** The same plan with different cards.


