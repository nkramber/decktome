# PR-7 question gate

Run date: 2026-08-25. Conversations: 2026-08-24.

Verdict: PASS. 29 of 30 conversations used catalog questions only. The bar is 25.

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 126 |
| From the catalog | 125 |
| Invented by the model | 1 |
| Replacements refused as rewords | 15 |
| Catalog questions that closed a slot | 98 |
| Invented questions that closed a slot | 0 |
| Catalog-only conversations | 29 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `locked` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 220. Time: 329.2 seconds.
- Tokens: 119040 input (0 cached), 19604 output.
- Cost: $0.0473.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your lifegain deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your lifegain deck?

**Turn 2, the user:** Commander, and white and black is right.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for your white-black lifegain deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander do you want: Lotho, Corrupt Shirriff; Giada, Font of Hope; or Stiltzkin, Moogle Merchant? Choose “None—name three more” if you’d like three additional suggestions.
  - Options: Lotho, Corrupt Shirriff / Giada, Font of Hope / Stiltzkin, Moogle Merchant / None—name three more

**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.


### 2. blink with a thin library

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a white-blue commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.

- [catalog slot=budget row=budget fit=0.90 filled=true] What budget do you have for cards to buy?

**Turn 4, the user:** No more than 40 dollars.

- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is the $40 limit for the cards you buy or for the entire deck's value?
  - Options: Cards I buy / Entire deck

### 3. a named card, role unknown

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the Grist, the Hunger Tide deck be built for: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your Commander table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore it for a fully optimized Grist, the Hunger Tide deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.


### 4. fun and janky

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=power row=jank fit=0.90 filled=true] What does “janky” mean to you: a low-power deck, an odd card nobody expects, or a plan that almost never works?
  - Options: A low-power deck / An odd card nobody expects / A plan that almost never works
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a preferred color or color combination?

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.


**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your sacrifice deck, should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Use your library first / Use only your library / Ignore your library

**Turn 4, the user:** You pick the commander, and build from my library first.


### 5. 60-card anything goes

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=format row=format fit=0.90 filled=true] For your 60-card deck, which format should we use: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—focus on a creature type, a mechanic, or a particular play style?
- [catalog slot=house_rules row=house_rules fit=0.20 filled=true] When you say anything goes, do you mean any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules
  - A replacement was refused as a reword (D-88).

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [catalog slot=house_rules row=house_format_limits fit=0.90 filled=false] For your house Modern format, should the normal limits apply: a 60-card minimum, up to four copies of each card name, and a 15-card sideboard?
  - Options: Yes, the normal limits / No, I’ll specify the limits
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? A dragon deck is often strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, red and green.


### 6. the strongest deck, no collection

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a preferred color or color combination?
- [catalog slot=budget row=budget fit=0.90 filled=true] What budget do you have for buying the cards?

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.


### 7. FNM on Friday

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=true] Which format is your FNM event: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** Pioneer. The best deck under budget.


**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


### 8. mill for a playgroup

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your mill deck?
- [catalog slot=budget row=budget_scope fit=0.90 filled=true] Is your $100 budget a cap on the cards you buy or on the whole deck’s value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you already have a blue-black commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your blue-black Commander table use? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your blue-black mill deck, should I use your library first, only cards from your library, or ignore it for a fully optimized build?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.

- [catalog slot=power row=table_tolerance fit=0.20 filled=true] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
  - A replacement was refused as a reword (D-88).

**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 1. Invented: 0.

The slots were full after turn 2. The last 2 messages never went out.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Brago blink deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.


### 10. no strong commander in the library

Collection: true. Catalog: 2. Invented: 0.

The slots were full after turn 2. The last 1 messages never went out.

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your lifegain deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for your lifegain deck?

**Turn 2, the user:** Commander. White and black.


### 11. another version after a build

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the new version use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** A different plan in the same colors.

- [catalog slot=plan_variant row=plan_choice fit=0.90 filled=false] Which of the two common plans in these colors do you want?

### 12. a frozen run asks nothing

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Switch to owned-only.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

### 13. a card to keep that is not the commander

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the Karlov of the Ghost Council lifegain deck use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, white and black, bracket 3.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your white-black Commander deck in bracket 3, should I start with your library, use only your library, or ignore it for a fully optimized build?
  - Options: Start with my library / Use only my library / Ignore my library
- [catalog slot=locked row=locked fit=0.20 filled=true] Must the deck keep Sanguine Bond, or may I cut cards that do not fit the lifegain plan?
  - Options: Keep every card I named / Cut cards that do not fit
  - A replacement was refused as a reword (D-88).

**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.


### 14. the user says none, then picks

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like for your lifegain deck?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for your lifegain commander deck?

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.


**Turn 3, the user:** None of those.


**Turn 4, the user:** The first of the new three is good.


### 15. standard at the store, no library

Collection: false. Catalog: 3. Invented: 0.

The slots were full after turn 2. The last 1 messages never went out.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=format row=format_store fit=0.90 filled=true] Which format does your local store's event run: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Standard. An aggro deck, red and white.


### 16. the strongest modern deck

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=format row=format fit=0.05 filled=false] You mentioned Modern—should the deck be built for Modern?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.

- [catalog slot=format row=format_store fit=0.05 filled=false] Which format does your tournament run: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
  - A replacement was refused as a reword (D-88).

### 17. precon upgrade at bracket 2

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=format row=format fit=0.05 filled=false] Which format are you playing: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
  - A replacement was refused as a reword (D-88).

**Turn 2, the user:** A superfriends deck. Build from my library first.


### 18. pauper burn on a budget

Collection: false. Catalog: 1. Invented: 0.

The slots were full after turn 2. The last 1 messages never went out.

**Turn 1, the user:** A Pauper burn deck, as cheap as possible.

- [catalog slot=format row=format fit=0.05 filled=true] You mentioned Pauper—should the deck be built for Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Pauper. Burn, mono red.


### 19. a big library, no theme

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I have a big collection. Build me something good.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on— a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. A token deck, green and white.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for your green-white token deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Peregrin Took, or Faramir, Field Commander? Choose “None—name three more” and I’ll name three more.
  - Options: The Earth King / Peregrin Took / Faramir, Field Commander / None—name three more

### 20. dinosaur tribal for a child

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Dinosaur decks are strongest in red and green.

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.


**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build the Gishath Dinosaur deck from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 4, the user:** Build from my library first.


### 21. vintage with proxies

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format are you building for: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=house_rules row=house_rules fit=0.50 filled=true] When you say you proxy everything, do you mean any card with no ban list, or using Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=house_rules row=house_format_limits fit=0.20 filled=true] For your house Vintage format, should the normal limits apply: a 60-card minimum, four copies per card name, and a 15-card sideboard?
  - Options: Yes, use the normal limits / No, I’ll specify the limits
  - A replacement was refused as a reword (D-88).
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? An artifact prison deck is often strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 8. Invented: 0.

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Extra-turns decks are strongest in blue and green.

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a blue-red commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your blue-red extra-turns deck, should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Vivi Ornitier, Ioreth of the Healing House, or Pashalik Mons? You can say “None—name three more.”
  - Options: Vivi Ornitier / Ioreth of the Healing House / Pashalik Mons / None—name three more
- [catalog slot=power row=table_tolerance fit=0.20 filled=true] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
  - A replacement was refused as a reword (D-88).

**Turn 4, the user:** My table is fine with extra turns.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Vivi Ornitier, Ioreth of the Healing House, or Pashalik Mons? Choose “None—name three more” and I’ll name three more.
  - Options: Vivi Ornitier / Ioreth of the Healing House / Pashalik Mons / None—name three more

### 23. land destruction

Collection: true. Catalog: 8. Invented: 0.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your land destruction deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your land destruction Commander deck?

**Turn 2, the user:** Commander. Red and green.

- [catalog slot=commander row=commander fit=0.20 filled=true] Do you have a red-green commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
  - A replacement was refused as a reword (D-88).
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Rydia, Summoner of Mist, or Peregrin Took? Or choose “None—name three more.”
  - Options: The Earth King / Rydia, Summoner of Mist / Peregrin Took / None—name three more
- [catalog slot=power row=table_tolerance fit=0.20 filled=true] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
  - A replacement was refused as a reword (D-88).

**Turn 4, the user:** My table does not mind.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: The Earth King, Rydia, Summoner of Mist, or Peregrin Took? Say “none” and I’ll name three more.
  - Options: The Earth King / Rydia, Summoner of Mist / Peregrin Took / None—name three more

### 24. stax, owned only

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the stax deck use: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the stax deck?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your white-blue deck, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.20 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
  - A replacement was refused as a reword (D-88).

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.

- [catalog slot=commander row=commander_pick fit=0.05 filled=false] Which commander do you want: Giada, Font of Hope, Ioreth of the Healing House, or Gandalf, Wandering Wizard? Choose “None—name three more” if you want three additional options.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more
- [catalog slot=power row=table_tolerance fit=0.90 filled=true] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts stax.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Giada, Font of Hope, Ioreth of the Healing House, or Gandalf, Wandering Wizard? Say “none” and I’ll name three more.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more

### 25. poison in a tournament

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=format row=format_store fit=0.20 filled=true] Which format does your event run: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
  - A replacement was refused as a reword (D-88).
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference for your infect deck?

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.


### 26. legacy for an event

Collection: false. Catalog: 3. Invented: 1.

**Turn 1, the user:** A Legacy deck for an event.

- [catalog slot=format row=format_store fit=0.20 filled=true] Which format does your event run: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
  - A replacement was refused as a reword (D-88).
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Any color preference?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.

- [INVENTED slot=locked row=locked fit=0.05 filled=false] In this Legacy blue-red Delver of Secrets tempo deck, must Delver of Secrets stay?
  - Options: Keep every card you named / Cut cards that don't fit
  - It replaced: Must the deck keep Delver of Secrets, or may I cut a card that does not fit the plan?

**Turn 3, the user:** FNM level, 600 dollars.


### 27. a card for the 99

Collection: true. Catalog: 9. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck be built for: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] Besides building around Grist, the Hunger Tide, what creature type, mechanic, or play style should the deck focus on?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences for the deck?

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [catalog slot=commander row=named_card_role fit=0.05 filled=true] Do you want Grist, the Hunger Tide as your commander, or as one card in the 99?
  - Options: As your commander / In the 99
  - A replacement was refused as a reword (D-88).
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only your library / Ignore my library

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=locked row=locked fit=0.20 filled=true] Must the deck keep Grist, the Hunger Tide and Grist, or may I cut a card that does not fit the plan?
  - Options: Keep every card I named / Cut what does not fit
  - A replacement was refused as a reword (D-88).

**Turn 4, the user:** You pick the commander, and keep Grist in the deck.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Gorbag of Minas Morgul, The Sackville-Bagginses, or June, Bounty Hunter? Choose “None—name three more” if you’d like three other options.
  - Options: Gorbag of Minas Morgul / The Sackville-Bagginses / June, Bounty Hunter / None—name three more

### 28. reanimator with two plans

Collection: true. Catalog: 5. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=format row=format fit=0.20 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the reanimator Commander deck?

**Turn 2, the user:** Commander. Black and green.

- [catalog slot=commander row=commander fit=0.05 filled=true] Do you have a black-green commander in mind, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For this black-green reanimator deck, should I build from your library first, use only cards in your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.


### 29. two hundred dollars

Collection: true. Catalog: 8. Invented: 0.

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=format row=format fit=0.90 filled=true] You mentioned a Commander deck—should the format be Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a particular play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the deck?

**Turn 2, the user:** Commander. An artifact deck, blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest a blue-red artifact commander?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget_scope fit=0.90 filled=true] Is your budget cap for the cards you buy, or for the whole blue-red artifact deck's value?
  - Options: The cards I buy / The whole deck

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.05 filled=false] Which commander do you want: Glóin, Dwarf Emissary, Katara, Bending Prodigy, or Sokka, Bold Boomeranger? Choose “None—name three more” if you want three additional options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Sokka, Bold Boomeranger / None—name three more
  - A replacement was refused as a reword (D-88).

**Turn 4, the user:** The 200 caps the cards I buy.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary, Katara, Bending Prodigy, or Sokka, Bold Boomeranger? Choose “None—name three more” if you want three other options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Sokka, Bold Boomeranger / None—name three more

### 30. another version, same plan

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the deck use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** The same plan with different cards.


