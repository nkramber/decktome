# PR-7 question gate

Run date: 2026-08-25. Conversations: 2026-08-24.

Verdict: PASS. 29 of 30 conversations used catalog questions only. The bar is 25.

No conversation called itself complete with a slot unanswered.

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 133 |
| From the catalog | 132 |
| Invented by the model | 1 |
| Replacements refused as rewords | 7 |
| Catalog questions that closed a slot | 85 |
| Invented questions that closed a slot | 1 |
| Catalog-only conversations | 29 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `pool` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 238. Time: 446.8 seconds.
- Tokens: 150680 input (0 cached), 20826 output.
- Cost: $0.0551.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for your lifegain deck?

**Turn 2, the user:** Commander, and white and black is right.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for your white-black lifegain deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander would you like: Lotho, Corrupt Shirriff; Giada, Font of Hope; or Stiltzkin, Moogle Merchant? Choose “None—name three more” if you want additional suggestions.
  - Options: Lotho, Corrupt Shirriff / Giada, Font of Hope / Stiltzkin, Moogle Merchant / None—name three more

**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.


### 2. blink with a thin library

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your blink deck?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your white-blue deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your white-blue table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** No more than 40 dollars.

- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Does the $40 limit apply to the cards you buy or to the entire deck's value?
  - Options: Cards I buy / Entire deck

### 3. a named card, role unknown

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Grist, the Hunger Tide deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your Commander table play at? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your Grist, the Hunger Tide deck, should I build from your library first, use only cards in your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.


### 4. fun and janky

Collection: true. Catalog: 8. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=power row=jank fit=0.90 filled=true] What does “janky” mean to you: a low-power deck, an odd card nobody expects, or a plan that almost never works?
  - Options: A low-power deck / An odd card nobody expects / A plan that almost never works
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for this fun, janky deck?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary; Gorbag of Minas Morgul; or Lotho, Corrupt Shirriff? You can also say “none” and I’ll name three more.
  - Options: Glóin, Dwarf Emissary / Gorbag of Minas Morgul / Lotho, Corrupt Shirriff / None—name three more
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 4, the user:** You pick the commander, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Sackville-Bagginses, June, Bounty Hunter, or Peregrin Took? Choose “None—name three more” if you want three additional options.
  - Options: The Sackville-Bagginses / June, Bounty Hunter / Peregrin Took / None—name three more

### 5. 60-card anything goes

Collection: false. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the 60-card deck use: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a particular play style?
- [catalog slot=house_rules row=house_rules fit=0.05 filled=true] When you say anything goes, do you mean any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [catalog slot=house_rules row=house_format_limits fit=0.05 filled=true] Should your Modern-named dragon deck follow the normal limits: at least 60 cards, no more than four copies of a card by name, and a 15-card sideboard?
  - Options: Yes, use the normal limits / No, I’ll specify the limits
  - Refused as a reword (D-88): Inside your house format, should the normal limits hold: 60 cards, up to four copies per name, and a 15-card sideboard?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should your dragon deck be: casual, FNM-level, or tournament-meta?
  - Options: Casual / FNM-level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference? A dragon deck is strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, red and green.


### 6. the strongest deck, no collection

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] What budget do you have for buying cards?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should your Modern deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your Modern event? I’ll tune the 15-card sideboard to it.

**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.


### 7. FNM on Friday

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=true] Which format is your FNM: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** Pioneer. The best deck under budget.

- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your Pioneer event? I’ll tune the 15 sideboard cards to it.

**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


### 8. mill for a playgroup

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the mill deck?
- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is your $100 budget a cap on the cards you buy or on the whole deck value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a blue-black commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your playgroup use? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your blue-black mill deck, should I use your library first, only cards from your library, or ignore it for a fully optimized build?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.

- [catalog slot=power row=table_tolerance fit=0.90 filled=false] Does your Bracket 3 table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your Brago blink deck need to fit? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 Exhibition / 2 Core / 3 Upgraded / 4 Optimized / 5 cEDH

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.


**Turn 3, the user:** Add him to the buy list. Bracket 3, white and blue.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for the cards to buy?

**Turn 4, the user:** Under 30 dollars, and build from my library first.


### 10. no strong commander in the library

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for your lifegain deck?

**Turn 2, the user:** Commander. White and black.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a white-black commander in mind, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Lotho, Corrupt Shirriff, Giada, Font of Hope, or Stiltzkin, Moogle Merchant? Choose “None—show three more” and I’ll name three more.
  - Options: Lotho, Corrupt Shirriff / Giada, Font of Hope / Stiltzkin, Moogle Merchant / None—show three more

### 11. another version after a build

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the new version use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on, such as a creature type, mechanic, or play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** A different plan in the same colors.


### 12. a frozen run asks nothing

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), theme (asked, no answer), colors (asked, no answer), power (never asked).

**Turn 1, the user:** Switch to owned-only.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

### 13. a card to keep that is not the commander

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the Karlov of the Ghost Council lifegain deck use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, white and black, bracket 3.

- [catalog slot=pool_rule row=pool fit=0.20 filled=true] Should I build from your library first, use only cards in your library, or ignore it for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library
  - Refused as a reword (D-88): Should I build from your existing library first, only your existing library, or ignore it for a fully optimized deck?
- [catalog slot=locked row=locked fit=0.90 filled=false] Must the deck keep Sanguine Bond, or may I cut cards that do not fit the lifegain plan?
  - Options: Keep every card I named / Cut cards that do not fit

**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.


### 14. the user says none, then picks

Collection: false. Catalog: 4. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a lifegain commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for your lifegain deck?

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander do you want for your bracket 3 lifegain deck?
  - Options: Aclazotz, Deepest Betrayal // Temple of the Dead / Beledros Witherbloom / Uro, Titan of Nature's Wrath / None—name three more

**Turn 3, the user:** None of those.


### 15. standard at the store, no library

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Standard. An aggro deck, red and white.


**Turn 3, the user:** FNM level.


### 16. the strongest modern deck

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should your Modern deck be: casual, FNM-level, or tournament-meta? Provide one of the following: Casual, FNM, or Tournament-meta.
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your Modern deck?
- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your Modern event? I’ll tune the 15-card sideboard to it.

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.


### 17. precon upgrade at bracket 2

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Use my library first / Use only cards from my library / Ignore my library

**Turn 2, the user:** A superfriends deck. Build from my library first.


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
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a preferred color or color combination?

**Turn 2, the user:** Commander. A token deck, green and white.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a green-white commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.20 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
  - Refused as a reword (D-88): What power bracket does your table play: 2 (precon level), 3 (upgraded), or 4 (high power)?

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Peregrin Took, or Faramir, Field Commander? Choose “None—name three more” and I’ll name three additional options.
  - Options: The Earth King / Peregrin Took / Faramir, Field Commander / None—name three more

### 20. dinosaur tribal for a child

Collection: true. Catalog: 3. Invented: 1.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Dinosaur decks are strongest in red and green.

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a dinosaur commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one

**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [INVENTED slot=pool_rule row=pool fit=0.20 filled=true] Would you like the deck built using your existing library first, only from your existing library, or without a library restriction for a fully optimized deck?
  - Options: Use my library first / Use only my library / No library restriction
  - It replaced: Build from your library first, only your library, or ignore it for a fully optimized deck?

**Turn 4, the user:** Build from my library first.


### 21. vintage with proxies

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format are you building for: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a particular play style?
- [catalog slot=house_rules row=house_rules fit=0.50 filled=false] When you say you proxy everything, does your table allow any card with no ban list, or do you use Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should this artifact prison deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.05 filled=false] Do you have any color preference? Artifact prison decks are strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 8. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Extra-turns decks are strongest in blue and green.

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for your blue-red deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized extra-turns deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Vivi Ornitier, Ioreth of the Healing House, or Pashalik Mons? Choose “None” and I’ll name three more.
  - Options: Vivi Ornitier / Ioreth of the Healing House / Pashalik Mons / None
- [catalog slot=power row=table_tolerance fit=0.20 filled=false] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
  - Refused as a reword (D-88): Does your table accept mill, land destruction, or stax?

**Turn 4, the user:** My table is fine with extra turns.

- [catalog slot=commander row=commander_pick fit=0.20 filled=false] Which commander do you want: Vivi Ornitier, Ioreth of the Healing House, or Pashalik Mons? Say “none” and I’ll name three more.
  - Options: Vivi Ornitier / Ioreth of the Healing House / Pashalik Mons / None—name three more
  - Refused as a reword (D-88): Which commander should I use: Vivi Ornitier, Ioreth of the Healing House, or Pashalik Mons? Say "none" and I’ll name three more.

### 23. land destruction

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your land destruction deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your land destruction deck?

**Turn 2, the user:** Commander. Red and green.

- [catalog slot=power row=table_tolerance fit=0.90 filled=false] Does your Commander table allow mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Rydia, Summoner of Mist, or Peregrin Took? Choose “None—name three more” if you want three additional options.
  - Options: The Earth King / Rydia, Summoner of Mist / Peregrin Took / None—name three more

**Turn 4, the user:** My table does not mind.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Rydia, Summoner of Mist, or Peregrin Took? Choose “None” and I’ll name three more.
  - Options: The Earth King / Rydia, Summoner of Mist / Peregrin Took / None

### 24. stax, owned only

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a preferred color or color combination?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a white-blue commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.50 filled=true] Which power bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Giada, Font of Hope, Ioreth of the Healing House, or Gandalf, Wandering Wizard? You can say “none” and I’ll name three more.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more
- [catalog slot=power row=table_tolerance fit=0.05 filled=false] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts stax.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Giada, Font of Hope; Ioreth of the Healing House; or Gandalf, Wandering Wizard? You can also say “none” and I’ll name three more.
  - Options: Giada, Font of Hope / Ioreth of the Healing House / Gandalf, Wandering Wizard / None—name three more

### 25. poison in a tournament

Collection: false. Catalog: 2. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a preferred color combination for the infect deck?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=meta row=meta fit=0.90 filled=false] What decks do people play at your Modern event? I’ll tune the 15-card sideboard to that metagame.

**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.


### 26. legacy for an event

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A Legacy deck for an event.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should this Legacy deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference? Legacy event decks are often strongest in white and red.

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.


**Turn 3, the user:** FNM level, 600 dollars.


### 27. a card for the 99

Collection: true. Catalog: 8. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences for the Grist, the Hunger Tide deck?

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [catalog slot=commander row=named_card_role fit=0.05 filled=false] Do you want Grist, the Hunger Tide as your commander, or in the 99?
  - Options: As my commander / In the 99
  - Refused as a reword (D-88): Do you want Grist, the Hunger Tide included as one card in the 99?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table play: 1 exhibition, 2 core, 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for the bracket 3 sacrifice deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=locked row=locked fit=0.20 filled=false] Must the deck keep Grist, the Hunger Tide and Grist in the 99, or may I cut a card that does not fit the sacrifice plan?
  - Options: Keep every card I named / Cut what does not fit
  - Refused as a reword (D-88): Must the deck keep Grist, the Hunger Tide, or may I cut a card that does not fit the plan?

**Turn 4, the user:** You pick the commander, and keep Grist in the deck.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Gorbag of Minas Morgul, The Sackville-Bagginses, or June, Bounty Hunter? You can also choose “None—name three more.”
  - Options: Gorbag of Minas Morgul / The Sackville-Bagginses / June, Bounty Hunter / None—name three more

### 28. reanimator with two plans

Collection: true. Catalog: 4. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your reanimator deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play for this reanimator deck? 2 is precon level, 3 is upgraded, and 4 is high power.
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

- [catalog slot=theme row=theme fit=0.90 filled=true] What should your Commander deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your playgroup use? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the deck?

**Turn 2, the user:** Commander. An artifact deck, blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for your blue-red artifact deck, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=budget row=budget_scope fit=0.05 filled=false] Is your budget a cap on the cards you buy, or on the whole blue-red artifact deck's total value?
  - Options: The cards I buy / The whole deck

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary; Katara, Bending Prodigy; or Sokka, Bold Boomeranger? Choose “None—name three more” if you want three additional options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Sokka, Bold Boomeranger / None—name three more

**Turn 4, the user:** The 200 caps the cards I buy.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary, Katara, Bending Prodigy, or Sokka, Bold Boomeranger? Choose “None—name three more” if you’d like three additional options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Sokka, Bold Boomeranger / None—name three more

### 30. another version, same plan

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), theme (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the same deck use with different cards: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** The same plan with different cards.


