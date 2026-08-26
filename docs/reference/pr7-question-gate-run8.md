# PR-7 question gate

Run date: 2026-08-25. Conversations: 2026-08-24.

Verdict: PASS. 26 of 30 conversations used catalog questions only. The bar is 25.

No conversation called itself complete with a slot unanswered.

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 135 |
| From the catalog | 131 |
| Invented by the model | 4 |
| Replacements refused as rewords | 6 |
| Catalog questions that closed a slot | 78 |
| Invented questions that closed a slot | 1 |
| Catalog-only conversations | 26 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `budget_scope` | 1 |
| `commander_pick` | 1 |
| `house_rules` | 1 |
| `pool` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 242. Time: 446.8 seconds.
- Tokens: 142857 input (0 cached), 20514 output.
- Cost: $0.0532.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 5. Invented: 1.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for your lifegain deck?

**Turn 2, the user:** Commander, and white and black is right.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander do you want: Lotho, Corrupt Shirriff, Giada, Font of Hope, or Stiltzkin, Moogle Merchant? Choose “None—name three more” if you’d like three additional suggestions.
  - Options: Lotho, Corrupt Shirriff / Giada, Font of Hope / Stiltzkin, Moogle Merchant / None—name three more

**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.

- [INVENTED slot=budget row=budget_scope fit=0.20 filled=false] Is the $50 cap for the cards you buy, or for the whole deck?
  - Options: The cards I buy / The whole deck
  - It replaced: Is that a cap on the cards you buy, or on the whole deck value?

### 2. blink with a thin library

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a preferred color or color combination?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your white-blue deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.20 filled=true] Which bracket does your white-blue Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
  - Refused as a reword (D-88): What power bracket does your table play: 2 (precon level), 3 (upgraded), or 4 (high power)?

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** No more than 40 dollars.

- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is the $40 limit for the cards you buy, or for the value of the whole deck?
  - Options: The cards I buy / The whole deck

### 3. a named card, role unknown

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Grist, the Hunger Tide deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your Commander table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.


### 4. fun and janky

Collection: true. Catalog: 5. Invented: 1.

Slots the deck needs and nobody answered: colors (asked, no answer), commander (asked, no answer).

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=power row=jank fit=0.90 filled=false] What does “janky” mean to you: a low-power deck, an odd card nobody expects, or a plan that almost never works?
  - Options: A low-power deck / An odd card nobody expects / A plan that almost never works
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.


**Turn 4, the user:** You pick the commander, and build from my library first.

- [INVENTED slot=commander row=commander_pick fit=0.05 filled=false] Which commander should I choose from your library: Glóin, Dwarf Emissary; Gorbag of Minas Morgul; or Lotho, Corrupt Shirriff? If none fit, name three more commanders.
  - Options: Glóin, Dwarf Emissary / Gorbag of Minas Morgul / Lotho, Corrupt Shirriff
  - It replaced: Which one do you want: Glóin, Dwarf Emissary, Gorbag of Minas Morgul, or Lotho, Corrupt Shirriff? Say "none" and I name three more.

### 5. 60-card anything goes

Collection: false. Catalog: 4. Invented: 1.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=format row=format fit=0.90 filled=true] What format should the deck use: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [INVENTED slot=house_rules row=house_rules fit=0.20 filled=false] Which house rules apply: any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules
  - It replaced: What does anything-goes mean at your table: any card with no ban list, or Vintage rules?

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should this dragon deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your dragon deck? Blue-red is generally the strongest option.

**Turn 3, the user:** The normal limits hold. Casual power, red and green.


### 6. the strongest deck, no collection

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for buying cards?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should your Modern deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=meta row=meta fit=0.90 filled=false] What decks do people play at your Modern event? I’ll tune the 15-card sideboard to that metagame.

**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.


### 7. FNM on Friday

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=true] Which format is your FNM: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** Pioneer. The best deck under budget.

- [catalog slot=meta row=meta fit=0.90 filled=false] What decks do people play at your Pioneer event? I’ll tune the 15-card sideboard to that metagame.

**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


### 8. mill for a playgroup

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the mill deck?
- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is your $100 budget a cap on the cards you buy, or on the whole deck's value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you already have a commander in mind for your blue-black deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which Commander bracket does your playgroup use? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.

- [catalog slot=power row=table_tolerance fit=0.90 filled=false] For your Bruvac the Grandiloquent mill deck at Bracket 3, does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.


**Turn 3, the user:** Add him to the buy list. Bracket 3, white and blue.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** Under 30 dollars, and build from my library first.

- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is your under-$30 limit for the cards you buy, or for the whole deck's value?
  - Options: Cards I buy / Whole deck

### 10. no strong commander in the library

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your lifegain deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a preferred color or color combination for your lifegain deck?

**Turn 2, the user:** Commander. White and black.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a white-and-black commander in mind, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.20 filled=true] Which bracket does your white-and-black lifegain deck's table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
  - Refused as a reword (D-88): What power bracket does your table play: 2 (precon level), 3 (upgraded), or 4 (high power)?

**Turn 3, the user:** Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Lotho, Corrupt Shirriff, Giada, Font of Hope, or Stiltzkin, Moogle Merchant? Choose “None—name three more” and I’ll suggest three more.
  - Options: Lotho, Corrupt Shirriff / Giada, Font of Hope / Stiltzkin, Moogle Merchant / None—name three more

### 11. another version after a build

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should this version use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on—a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** A different plan in the same colors.


### 12. a frozen run asks nothing

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), theme (asked, no answer), colors (asked, no answer), power (never asked).

**Turn 1, the user:** Switch to owned-only.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck do—focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

### 13. a card to keep that is not the commander

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library
- [catalog slot=locked row=locked fit=0.50 filled=false] Must the deck keep Sanguine Bond, or may I cut cards that do not fit the lifegain plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 2, the user:** Commander, white and black, bracket 3.


**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.


### 14. the user says none, then picks

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer), commander (asked, no answer).

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a lifegain commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference for your lifegain deck?

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Aclazotz, Deepest Betrayal // Temple of the Dead, Beledros Witherbloom, or Uro, Titan of Nature's Wrath? Choose “None—name three more” if you want three additional suggestions.
  - Options: Aclazotz, Deepest Betrayal // Temple of the Dead / Beledros Witherbloom / Uro, Titan of Nature's Wrath / None—name three more

**Turn 3, the user:** None of those.

- [catalog slot=commander row=commander_pick fit=0.20 filled=true] Which commander do you want: Aclazotz, Deepest Betrayal // Temple of the Dead, Beledros Witherbloom, or Uro, Titan of Nature's Wrath? Choose “None—name three more” if you want three additional options.
  - Options: Aclazotz, Deepest Betrayal // Temple of the Dead / Beledros Witherbloom / Uro, Titan of Nature's Wrath / None—name three more
  - Refused as a reword (D-88): Which one do you want: Aclazotz, Deepest Betrayal // Temple of the Dead, Beledros Witherbloom, or Uro, Titan of Nature's Wrath?

**Turn 4, the user:** The first of the new three is good.


### 15. standard at the store, no library

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Standard. An aggro deck, red and white.


**Turn 3, the user:** FNM level.


### 16. the strongest modern deck

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should your Modern deck be: casual, FNM-level, or tournament-meta? (You said you want the strongest.)
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference for your Modern deck?
- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your event so I can tune the 15-card sideboard to it?

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.


### 17. precon upgrade at bracket 2

Collection: true. Catalog: 0. Invented: 0.

The slots were full after turn 1. The last 1 messages never went out.

**This conversation asked nothing.** One message filled every slot, so it tests no question.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.


### 18. pauper burn on a budget

Collection: false. Catalog: 1. Invented: 0.

**Turn 1, the user:** A Pauper burn deck, as cheap as possible.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should your Pauper burn deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta

**Turn 2, the user:** Pauper. Burn, mono red.


**Turn 3, the user:** Casual power, and 25 dollars is the cap.


### 19. a big library, no theme

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I have a big collection. Build me something good.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. A token deck, green and white.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for your green-white token deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Peregrin Took, or Faramir, Field Commander? Choose “None—name three more” if you’d like three additional options.
  - Options: The Earth King / Peregrin Took / Faramir, Field Commander / None—name three more

### 20. dinosaur tribal for a child

Collection: true. Catalog: 3. Invented: 1.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Dinosaur decks are strongest in red and green.

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a Dinosaur commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one

**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [INVENTED slot=pool_rule row=pool fit=0.20 filled=true] Should I use your library first, only use your library, or ignore your library?
  - Options: Use my library first / Only use my library / Ignore my library
  - It replaced: Build from your library first, only your library, or ignore it for a fully optimized deck?

**Turn 4, the user:** Build from my library first.


### 21. vintage with proxies

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format are you playing: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=house_rules row=house_rules fit=0.90 filled=false] Since you proxy everything, what does “anything-goes” mean at your table: any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should this Vintage artifact prison deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.05 filled=false] Do you have any color preference? Artifact prison is strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 8. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format do you want: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Extra-turns decks are strongest in blue and green.

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a blue-red commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your blue-red Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For this blue-red extra-turns deck, should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Vivi Ornitier, Ioreth of the Healing House, or Pashalik Mons? Choose “None” and I’ll name three more.
  - Options: Vivi Ornitier / Ioreth of the Healing House / Pashalik Mons / None
- [catalog slot=power row=table_tolerance fit=0.20 filled=false] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
  - Refused as a reword (D-88): Which of these does your table accept: mill, land destruction, or stax?

**Turn 4, the user:** My table is fine with extra turns.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Vivi Ornitier, Ioreth of the Healing House, or Pashalik Mons? Say “none” and I’ll name three more.
  - Options: Vivi Ornitier / Ioreth of the Healing House / Pashalik Mons / None—name three more

### 23. land destruction

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for your land destruction deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your Commander table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the land destruction deck?

**Turn 2, the user:** Commander. Red and green.

- [catalog slot=power row=table_tolerance fit=0.90 filled=true] Does your Commander table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized Commander deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Rydia, Summoner of Mist, or Peregrin Took? Choose “None—name three more” if you want three additional options.
  - Options: The Earth King / Rydia, Summoner of Mist / Peregrin Took / None—name three more

**Turn 4, the user:** My table does not mind.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Rydia, Summoner of Mist, or Peregrin Took? You can choose “None—name three more.”
  - Options: The Earth King / Rydia, Summoner of Mist / Peregrin Took / None—name three more

### 24. stax, owned only

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the stax deck?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a white-blue commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.50 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.

- [catalog slot=commander row=commander_pick fit=0.05 filled=false] Which commander do you want: Giada, Font of Hope; Ioreth of the Healing House; or Gandalf, Wandering Wizard? Choose “None—name three more” if you want three other options.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more
- [catalog slot=power row=table_tolerance fit=0.20 filled=false] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
  - Refused as a reword (D-88): Does your table accept mill, land destruction, or extra turns?

**Turn 4, the user:** My table accepts stax.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Giada, Font of Hope, Ioreth of the Healing House, or Gandalf, Wandering Wizard? Say “none” and I’ll name three more.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more

### 25. poison in a tournament

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference for the infect deck?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=meta row=meta fit=0.90 filled=false] What decks do people play at your Modern event? I’ll tune the 15-card sideboard to them.

**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.

- [catalog slot=plan_variant row=plan_choice fit=0.90 filled=false] For the best tournament-level deck you can build under $400 for a shop with aggro and control decks, which of the two common plans do you want?

### 26. legacy for an event

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Legacy deck for an event.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the Legacy deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the Legacy deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the Legacy deck?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.

- [catalog slot=locked row=locked fit=0.90 filled=false] Must the deck keep Delver of Secrets, or may I cut a card that does not fit the plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 3, the user:** FNM level, 600 dollars.


### 27. a card for the 99

Collection: true. Catalog: 8. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Grist, the Hunger Tide deck: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences for the deck?

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [catalog slot=commander row=named_card_role fit=0.20 filled=true] Do you want Grist, the Hunger Tide as your commander, or as one card in the 99?
  - Options: As my commander / In the 99
  - Refused as a reword (D-88): Do you want Grist, the Hunger Tide in the 99, or a different commander?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=locked row=locked fit=0.90 filled=true] Must the deck keep Grist, the Hunger Tide and Grist, or may I cut a card that does not fit the plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 4, the user:** You pick the commander, and keep Grist in the deck.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Gorbag of Minas Morgul, The Sackville-Bagginses, or June, Bounty Hunter? Choose “None—name three more” if you want three other options.
  - Options: Gorbag of Minas Morgul / The Sackville-Bagginses / June, Bounty Hunter / None—name three more

### 28. reanimator with two plans

Collection: true. Catalog: 4. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your reanimator deck, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play for the reanimator deck? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your reanimator deck?

**Turn 2, the user:** Commander. Black and green.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your black-green Commander reanimator deck, should I build from your library first, use only cards in your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.


### 29. two hundred dollars

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander would you like, or should I suggest three options?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your playgroup use? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** Commander. An artifact deck, blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for your blue-red artifact deck, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=budget row=budget_scope fit=0.50 filled=false] Is your budget a cap on the cards you buy, or on the whole blue-red artifact deck's total value?
  - Options: The cards I buy / The whole deck

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander should I choose: Glóin, Dwarf Emissary; Katara, Bending Prodigy; or Sokka, Bold Boomeranger? Choose “None—name three more” if you want three additional options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Sokka, Bold Boomeranger / None—name three more

**Turn 4, the user:** The 200 caps the cards I buy.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary; Katara, Bending Prodigy; or Sokka, Bold Boomeranger? Choose “None” and I’ll name three more.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Sokka, Bold Boomeranger / None

### 30. another version, same plan

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), theme (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the same deck use with different cards: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** The same plan with different cards.


