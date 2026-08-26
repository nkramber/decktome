# PR-7 question gate

Run date: 2026-08-25. Conversations: 2026-08-24.

Verdict: FAIL. 21 of 30 conversations used catalog questions only. The bar is 25.

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 108 |
| From the catalog | 97 |
| Invented by the model | 11 |
| Catalog questions that closed a slot | 71 |
| Invented questions that closed a slot | 5 |
| Catalog-only conversations | 21 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `house_format_limits` | 2 |
| `pool` | 2 |
| `theme` | 2 |
| `acquisition` | 1 |
| `commander_pick` | 1 |
| `house_rules` | 1 |
| `locked` | 1 |
| `table_tolerance` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 212. Time: 364.9 seconds.
- Tokens: 103788 input (0 cached), 19321 output.
- Cost: $0.0439.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.95 filled=true] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Commander, and white and black is right.


**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.


**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.


### 2. blink with a thin library

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.95 filled=true] Do you have a color preference? Blink is strongest in white, blue, black, and green.

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.98 filled=true] Do you have a commander in mind, or would you like me to suggest one for your white-blue deck?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.72 filled=true] Which bracket does your white-blue Commander table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.

- [catalog slot=budget row=budget fit=0.88 filled=true] What budget would you like to set for cards to buy?

**Turn 4, the user:** No more than 40 dollars.

- [catalog slot=budget row=budget_scope fit=0.05 filled=false] Is the $40 limit for the cards you buy, or for the entire deck's value?
  - Options: Cards I buy / Entire deck

### 3. a named card, role unknown

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like for your Grist, the Hunger Tide deck: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=pool_rule row=pool fit=0.95 filled=true] For your Grist, the Hunger Tide Commander deck, should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.


### 4. fun and janky

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=power row=jank fit=1.00 filled=false] What does “janky” mean to you: a low-power deck, an odd card nobody expects, or a plan that almost never works?
  - Options: A low-power deck / An odd card nobody expects / A plan that almost never works
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Fun, janky decks are especially strong in blue, black, and green.

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.


**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.

- [catalog slot=pool_rule row=pool fit=0.05 filled=true] For your sacrifice deck, should I use your library first, use only your library, or ignore it for a fully optimized build?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 4, the user:** You pick the commander, and build from my library first.


### 5. 60-card anything goes

Collection: false. Catalog: 2. Invented: 1.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=theme row=theme fit=0.95 filled=true] What should the deck focus on: a creature type, mechanic, or play style?
- [INVENTED slot=house_rules row=house_format_limits fit=0.18 filled=true] Are there any other house rules or limits on card copies we should follow?
- [catalog slot=colors row=colors fit=0.75 filled=true] Do you have a preferred color or color combination?

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.


**Turn 3, the user:** The normal limits hold. Casual power, red and green.


### 6. the strongest deck, no collection

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme_competitive fit=0.90 filled=false] Since you own no cards, do you want a named tier-one deck or the best deck under your budget?
  - Options: A named deck / The best deck under budget
- [catalog slot=colors row=colors fit=0.25 filled=false] Do you have any color preference?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=theme row=theme fit=0.28 filled=false] What should the Modern deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=budget row=budget fit=1.00 filled=true] What’s your budget for buying cards?

**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.


### 7. FNM on Friday

Collection: false. Catalog: 3. Invented: 1.

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=theme row=theme_competitive fit=0.95 filled=false] For FNM this Friday, do you want a named tier-one deck, or the best deck under your budget?
  - Options: A named deck / The best deck under my budget
- [catalog slot=colors row=colors fit=0.70 filled=false] Do you have any color preference?

**Turn 2, the user:** Pioneer. The best deck under budget.

- [INVENTED slot=theme row=theme fit=0.05 filled=false] What is your budget in dollars?

**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


**Turn 4, the user:** I buy online, and I need the cards by Thursday.

- [catalog slot=budget row=acquisition fit=0.02 filled=false] You buy online—by what date do you need the cards?

### 8. mill for a playgroup

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.95 filled=true] Do you have a color preference? Mill is strongest in blue, black, and green.
- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is your $100 budget a cap on the cards you buy or on the whole deck’s value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=commander row=commander fit=1.00 filled=true] Do you have a blue-black commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.95 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.20 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.

- [catalog slot=power row=table_tolerance fit=0.95 filled=true] For this Bracket 3 Bruvac the Grandiloquent deck, does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=format row=format fit=0.92 filled=true] Which format should your Brago blink deck use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.

- [catalog slot=power row=power_commander fit=0.95 filled=true] Which bracket does your Brago, King Eternal Commander deck fit? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Add him to the buy list. Bracket 3, white and blue.

- [catalog slot=budget row=budget fit=0.95 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** Under 30 dollars, and build from my library first.

- [catalog slot=budget row=budget_scope fit=0.95 filled=false] Is the under-$30 limit for the cards you buy, or for the whole deck's value?
  - Options: Cards I buy / Whole deck value

### 10. no strong commander in the library

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like for your lifegain deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Commander. White and black.

- [catalog slot=commander row=commander fit=1.00 filled=true] Do you have a white-black commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.95 filled=true] Which bracket does your white-black Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.

- [catalog slot=commander row=commander_pick fit=0.96 filled=false] Which commander do you want: Zidane, Tantalus Thief, Ignis Scientia, or Peregrin Took? Choose “None—name three more” if you want three additional options.
  - Options: Zidane, Tantalus Thief / Ignis Scientia / Peregrin Took / None—name three more

### 11. another version after a build

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=format row=format fit=0.95 filled=false] Which format should the new version use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on—a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.75 filled=true] Do you have a preferred color or color combination?

**Turn 2, the user:** A different plan in the same colors.

- [catalog slot=plan_variant row=plan_choice fit=0.95 filled=false] Which of the two common plans in these colors do you want?

### 12. a frozen run asks nothing

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Switch to owned-only.

- [catalog slot=format row=format fit=0.05 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.05 filled=false] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.05 filled=false] Do you have any color preference?

### 13. a card to keep that is not the commander

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=format row=format fit=0.95 filled=true] Which format should the Karlov of the Ghost Council lifegain deck use: Commander, Standard, Modern, Pioneer, or another format?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=locked row=locked fit=0.02 filled=true] Must the deck keep Sanguine Bond, or may cards that do not fit the lifegain plan be cut?
  - Options: Keep every card I named / Cut what does not fit

**Turn 2, the user:** Commander, white and black, bracket 3.

- [catalog slot=pool_rule row=pool fit=0.95 filled=true] Should I build your white-black Commander lifegain deck from your library first, only from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.


### 14. the user says none, then picks

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=format row=format fit=0.98 filled=false] Which format would you like: Commander, Standard, Modern, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.72 filled=true] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.


**Turn 3, the user:** None of those.


**Turn 4, the user:** The first of the new three is good.


### 15. standard at the store, no library

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=theme row=theme fit=0.98 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Standard. An aggro deck, red and white.


**Turn 3, the user:** FNM level.


### 16. the strongest modern deck

Collection: false. Catalog: 2. Invented: 1.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=theme row=theme_competitive fit=0.96 filled=false] Do you want a named tier-one deck, or the strongest deck regardless of cost?
  - Options: A named deck / The best deck under budget
- [catalog slot=colors row=colors fit=0.18 filled=false] Do you have any color preference?

**Turn 2, the user:** The best deck under budget. Any colors.

- [INVENTED slot=theme row=theme fit=0.05 filled=false] What budget should the deck stay within?

**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.


### 17. precon upgrade at bracket 2

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=theme row=theme_card_named fit=0.95 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?

**Turn 2, the user:** A superfriends deck. Build from my library first.


### 18. pauper burn on a budget

Collection: false. Catalog: 0. Invented: 0.

**Turn 1, the user:** A Pauper burn deck, as cheap as possible.


**Turn 2, the user:** Pauper. Burn, mono red.


**Turn 3, the user:** Casual power, and 25 dollars is the cap.


### 19. a big library, no theme

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** I have a big collection. Build me something good.

- [catalog slot=format row=format fit=0.98 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme_competitive fit=0.72 filled=true] Do you want a named tier-one deck, or the best deck possible within your budget?
  - Options: A named deck / The best deck under budget
- [catalog slot=colors row=colors fit=0.58 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. A token deck, green and white.

- [catalog slot=commander row=commander fit=1.00 filled=true] Do you have a green-white commander in mind for your token deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


### 20. dinosaur tribal for a child

Collection: true. Catalog: 2. Invented: 2.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [INVENTED slot=house_rules row=house_rules fit=0.10 filled=false] What power level and house rules should the casual dinosaur deck follow?
- [catalog slot=colors row=colors fit=0.75 filled=true] Do you have a color preference? Dinosaurs are strongest in red and green.

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.


**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [INVENTED slot=pool_rule row=pool fit=0.05 filled=true] Should I prioritize your existing card collection when building the Gishath, Sun's Avatar Dinosaur deck?

**Turn 4, the user:** Build from my library first.


### 21. vintage with proxies

Collection: false. Catalog: 3. Invented: 1.

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.50 filled=true] Which format are you playing: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.45 filled=true] What should the deck focus on—a creature type, a mechanic, or a play style?
- [INVENTED slot=house_rules row=house_format_limits fit=0.20 filled=false] What special deckbuilding rules should I follow at your table, given that you proxy everything?

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference? Artifact prison is strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 7. Invented: 1.

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=1.00 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.95 filled=true] Do you have a color preference? Extra-turns decks are strongest in blue and green.

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=commander row=commander fit=0.98 filled=true] Do you have a blue-red commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.86 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.02 filled=false] Which commander do you want: Lotho, Corrupt Shirriff; Peregrin Took; or Massacre Girl, Known Killer? Choose “None—name three more” if you want three additional options.
  - Options: Lotho, Corrupt Shirriff / Peregrin Took / Massacre Girl, Known Killer / None—name three more
- [catalog slot=power row=table_tolerance fit=0.05 filled=true] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table is fine with extra turns.

- [INVENTED slot=commander row=commander_pick fit=0.05 filled=false] Which blue-red commander would you like to use for your extra-turns deck?

### 23. land destruction

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=format row=format fit=0.05 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Land destruction is strongest in blue, black, and green.

**Turn 2, the user:** Commander. Red and green.

- [catalog slot=pool_rule row=pool fit=0.95 filled=true] For your Commander deck in red and green, should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** My table does not mind.


### 24. stax, owned only

Collection: true. Catalog: 6. Invented: 1.

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.95 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.85 filled=true] Do you have a color preference? Stax is strongest in white, blue, and green.

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.95 filled=true] Do you have a commander in mind for your white-blue deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.45 filled=true] Which bracket does your white-blue table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.

- [catalog slot=commander row=commander_pick fit=0.98 filled=false] Which commander do you want: Lotho, Corrupt Shirriff, Zhao, the Moon Slayer, or Troyan, Gutsy Explorer? You can say “none” and I’ll name three more.
  - Options: Lotho, Corrupt Shirriff / Zhao, the Moon Slayer / Troyan, Gutsy Explorer / None—name three more
- [INVENTED slot=power row=table_tolerance fit=0.30 filled=true] Are mass land destruction and extra turns acceptable at your table?
  - Options: Yes / No

**Turn 4, the user:** My table accepts stax.

- [catalog slot=commander row=commander_pick fit=0.95 filled=false] Which commander do you want: Giada, Font of Hope, Ioreth of the Healing House, or Gandalf, Wandering Wizard? Choose “None—name three more” if you want three additional options.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more

### 25. poison in a tournament

Collection: false. Catalog: 1. Invented: 0.

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=colors row=colors fit=0.95 filled=false] Do you have a color preference? Infect is strongest in blue, black, and green.
  - Options: Blue / Black / Green / Blue-Black / Blue-Green / Black-Green / Blue-Black-Green / No preference

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.


### 26. legacy in three weeks

Collection: false. Catalog: 3. Invented: 1.

**Turn 1, the user:** A Legacy deck. I need it in three weeks for an event.

- [catalog slot=theme row=theme fit=0.72 filled=true] What should the Legacy deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.62 filled=true] Do you have any color preferences for the Legacy deck?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.

- [catalog slot=locked row=locked fit=0.93 filled=false] In this Legacy blue-red Delver of Secrets tempo deck, must I keep every card you named, or may I cut cards that do not fit the plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 3, the user:** FNM level, 600 dollars.


**Turn 4, the user:** I buy online, and I need the cards two days before the event.

- [INVENTED slot=budget row=acquisition fit=0.02 filled=false] What is your total budget for the Delver of Secrets tempo deck?

### 27. a card for the 99

Collection: true. Catalog: 1. Invented: 2.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=format row=format fit=0.95 filled=true] Which format should the deck be built for: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [INVENTED slot=locked row=locked fit=0.20 filled=true] Which Grist-centered strategy would you like to use in a format where Grist isn’t the commander?

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [INVENTED slot=pool_rule row=pool fit=0.22 filled=true] Would you like to build the black-green sacrifice Commander deck from your existing card pool, or should I use any cards available?
  - Options: Use my existing card pool / Use any available cards

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.


**Turn 4, the user:** You pick the commander, and keep Grist in the deck.


### 28. reanimator with two plans

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=commander row=commander fit=1.00 filled=true] Do you have a commander in mind for your reanimator deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.80 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference? Reanimator is strongest in white, black, and green.

**Turn 2, the user:** Commander. Black and green.

- [catalog slot=pool_rule row=pool fit=0.95 filled=true] For your black-green Commander reanimator deck, should I use your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** Reanimate one big creature, not a value loop.


### 29. two hundred dollars

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=theme row=theme fit=0.98 filled=true] What should the Commander deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.62 filled=true] Do you have any color preferences for the deck?

**Turn 2, the user:** Commander. An artifact deck, blue and red.


**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** The 200 caps the cards I buy.


### 30. another version, same plan

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.92 filled=false] Which format should the same deck use with different cards: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.18 filled=true] What should the deck do—focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.25 filled=false] Do you have any color preference?

**Turn 2, the user:** The same plan with different cards.


