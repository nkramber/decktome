# PR-7 question gate

Run date: 2026-08-25. Conversations: 2026-08-24.

Verdict: PASS. 28 of 30 conversations used catalog questions only. The bar is 25.

No conversation called itself complete with a slot unanswered.

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 138 |
| From the catalog | 136 |
| Invented by the model | 2 |
| Replacements refused as rewords | 7 |
| Catalog questions that closed a slot | 92 |
| Invented questions that closed a slot | 0 |
| Catalog-only conversations | 28 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `locked` | 1 |
| `table_tolerance` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 242. Time: 514.0 seconds.
- Tokens: 152364 input (1028 cached), 20974 output.
- Cost: $0.0555.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences for your lifegain deck?

**Turn 2, the user:** Commander, and white and black is right.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or would you like me to suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander do you want: Aerith Gainsborough, Minwu, White Mage, or Cecil, Dark Knight // Cecil, Redeemed Paladin? Choose “None—name three more” if you want three additional suggestions.
  - Options: Aerith Gainsborough / Minwu, White Mage / Cecil, Dark Knight // Cecil, Redeemed Paladin / None—name three more

**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.


### 2. blink with a thin library

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for your white-blue deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your white-blue Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** No more than 40 dollars.

- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is the $40 limit for the cards you buy, or for the entire deck value?
  - Options: Cards I buy / Entire deck

### 3. a named card, role unknown

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Grist, the Hunger Tide deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your Commander table use? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore your library for a fully optimized Grist, the Hunger Tide deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.


### 4. fun and janky

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=power row=jank fit=0.90 filled=true] What does “janky” mean to you: a low-power deck, an odd card nobody expects, or a plan that almost never works?
  - Options: A low-power deck / An odd card nobody expects / A plan that almost never works
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.

- [catalog slot=pool_rule row=pool fit=0.20 filled=true] For your sacrifice deck, should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library
  - Refused as a reword (D-88): Build from your library first, only your library, or use any cards for a fully optimized deck?

**Turn 4, the user:** You pick the commander, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Denethor, Ruling Steward; Gríma Wormtongue; or Old Man Willow? Choose “None—name three more” and I’ll suggest three additional options.
  - Options: Denethor, Ruling Steward / Gríma Wormtongue / Old Man Willow / None—name three more

### 5. 60-card anything goes

Collection: false. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the 60-card deck use: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=house_rules row=house_rules fit=0.20 filled=true] When you say anything goes, do you mean any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [catalog slot=house_rules row=house_format_limits fit=0.90 filled=true] Should we use the normal limits for this Modern-style deck: at least 60 cards, no more than four copies of any card by name, and a 15-card sideboard?
  - Options: Yes, the normal limits / No, I’ll specify different limits
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How powerful should your dragon deck be: casual, FNM-level, or tournament-meta?
  - Options: Casual / FNM-level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your dragon deck? Blue-red is generally the strongest pairing for it.

**Turn 3, the user:** The normal limits hold. Casual power, red and green.


### 6. the strongest deck, no collection

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a preferred color or color combination?
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

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the mill deck use: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the mill deck?
- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is your $100 budget a cap on the cards you buy, or on the whole deck’s value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you already have a blue-black commander in mind, or should I suggest one?
  - Options: I have a commander in mind / Suggest a commander
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which Commander bracket does your table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore it for a fully optimized blue-black mill deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.

- [catalog slot=power row=table_tolerance fit=0.90 filled=false] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the Brago blink deck use?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your Brago, King Eternal Commander table play at? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 — Exhibition / 2 — Core / 3 — Upgraded / 4 — Optimized / 5 — cEDH

**Turn 3, the user:** Add him to the buy list. Bracket 3, white and blue.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for the cards to buy?

**Turn 4, the user:** Under 30 dollars, and build from my library first.


### 10. no strong commander in the library

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a preferred color or color combination?

**Turn 2, the user:** Commander. White and black.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for your white-black deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Aerith Gainsborough, Minwu, White Mage, or Cecil, Dark Knight // Cecil, Redeemed Paladin? Choose “None—name three more” if you’d like three additional suggestions.
  - Options: Aerith Gainsborough / Minwu, White Mage / Cecil, Dark Knight // Cecil, Redeemed Paladin / None—name three more

### 11. another version after a build

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should this version use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on— a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** A different plan in the same colors.


### 12. a frozen run asks nothing

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), theme (asked, no answer), colors (asked, no answer), power (never asked).

**Turn 1, the user:** Switch to owned-only.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

### 13. a card to keep that is not the commander

Collection: true. Catalog: 2. Invented: 1.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table use? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 — Exhibition / 2 — Core / 3 — Upgraded / 4 — Optimized / 5 — cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore your library for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library
- [INVENTED slot=locked row=locked fit=0.20 filled=false] Besides keeping Sanguine Bond, should the deck keep every other card you named, or may I cut cards that do not fit?
  - Options: Keep every card I named / Cut anything that does not fit
  - It replaced: Must the deck keep Sanguine Bond, or may I cut a card that does not fit the plan?

**Turn 2, the user:** Commander, white and black, bracket 3.


**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.


### 14. the user says none, then picks

Collection: false. Catalog: 4. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a lifegain commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your lifegain deck's table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your lifegain commander deck?

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander do you want: Vito, Thorn of the Dusk Rose, Heliod, Sun-Crowned, or Haliya, Guided by Light? Choose “None—name three more” and I’ll suggest three additional commanders.
  - Options: Vito, Thorn of the Dusk Rose / Heliod, Sun-Crowned / Haliya, Guided by Light / None—name three more

**Turn 3, the user:** None of those.


### 15. standard at the store, no library

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Standard. An aggro deck, red and white.


**Turn 3, the user:** FNM level.


### 16. the strongest modern deck

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the Modern deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the Modern deck?
- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your event so I can tune the 15-card sideboard to it?

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.


### 17. precon upgrade at bracket 2

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=pool_rule row=pool fit=0.20 filled=true] Should I build from your precon cards first, use only cards from the precon, or ignore the precon for a fully optimized deck?
  - Options: Use the precon cards first / Use only cards from the precon / Ignore the precon and fully optimize
  - Refused as a reword (D-88): Build from your library first, use only cards from your library, or ignore it for a fully optimized deck?

**Turn 2, the user:** A superfriends deck. Build from my library first.


### 18. pauper burn on a budget

Collection: false. Catalog: 1. Invented: 0.

**Turn 1, the user:** A Pauper burn deck, as cheap as possible.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the Pauper burn deck be: casual, FNM level, or tournament-meta?
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

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Peregrin Took, Rosie Cotton of South Lane, or Quina, Qu Gourmet? Choose “None—name three more” if you want three other options.
  - Options: Peregrin Took / Rosie Cotton of South Lane / Quina, Qu Gourmet / None—name three more

### 20. dinosaur tribal for a child

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Dinosaur decks are strongest in red and green.

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for your casual Dinosaur deck?
  - Options: I have one in mind / Suggest one

**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore your library for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 4, the user:** Build from my library first.


### 21. vintage with proxies

Collection: false. Catalog: 7. Invented: 1.

Slots the deck needs and nobody answered: colors (asked, no answer), commander (asked, no answer).

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format are you playing: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—focus on a creature type, a mechanic, or a play style?
- [catalog slot=house_rules row=house_rules fit=0.50 filled=true] When you say you proxy everything, does your table allow any card with no ban list, or do you use Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=house_rules row=house_format_limits fit=0.90 filled=true] For your Vintage artifact prison deck, should the usual limits apply: a 60-card minimum, four copies of each card name, and a 15-card sideboard?
  - Options: Yes, the normal limits / No, I will specify the limits
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How powerful should your Vintage artifact prison deck be: casual, FNM-level, or tournament-meta?
  - Options: Casual / FNM-level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference for the artifact prison deck? It is strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [INVENTED slot=power row=table_tolerance fit=0.20 filled=false] What power level should the deck target: casual or high power?
  - Options: Casual / High power
  - It replaced: Does your table accept mill, land destruction, extra turns, or stax?

### 22. extra turns commander

Collection: true. Catalog: 8. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Extra-turns decks are strongest in blue and green.

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a blue-red commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your blue-red Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your blue-red extra-turns deck, should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Birgi, God of Storytelling // Harnfel, Horn of Bounty, Emrakul, the Promised End, or Vnwxt, Verbose Host? Choose “None—name three more” if you want three additional options.
  - Options: Birgi, God of Storytelling // Harnfel, Horn of Bounty / Emrakul, the Promised End / Vnwxt, Verbose Host / None—name three more
- [catalog slot=power row=table_tolerance fit=0.20 filled=false] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
  - Refused as a reword (D-88): What power level does your table accept: mill, land destruction, extra turns, or stax?

**Turn 4, the user:** My table is fine with extra turns.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Birgi, God of Storytelling // Harnfel, Horn of Bounty; Emrakul, the Promised End; or Vnwxt, Verbose Host? Choose “None—name three more” and I’ll name three more.
  - Options: Birgi, God of Storytelling // Harnfel, Horn of Bounty / Emrakul, the Promised End / Vnwxt, Verbose Host / None—name three more

### 23. land destruction

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for your land destruction deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your land destruction deck?

**Turn 2, the user:** Commander. Red and green.

- [catalog slot=power row=table_tolerance fit=0.90 filled=false] Does your Commander table allow mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Choose “None—name three more” if you’d like me to name three additional options.
  - Options: The Earth King / Toph, Hardheaded Teacher / Toph, the Blind Bandit / None—name three more

**Turn 4, the user:** My table does not mind.

- [catalog slot=commander row=commander_pick fit=0.20 filled=false] Which commander do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Say “none” and I’ll name three more.
  - Options: The Earth King / Toph, Hardheaded Teacher / Toph, the Blind Bandit / None; name three more
  - Refused as a reword (D-88): Which one do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Say "none" and I name three more.

### 24. stax, owned only

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the stax deck?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a white-blue commander in mind, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.05 filled=true] Which power bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
  - Refused as a reword (D-88): Which bracket does your table play: 2 (precon level), 3 (upgraded), or 4 (high power)?

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.

- [catalog slot=commander row=commander_pick fit=0.05 filled=false] Which commander do you want: Thalia, Guardian of Thraben, Pippin, Guard of the Citadel, or Kataki, War's Wage? Choose “None” and I’ll name three more.
  - Options: Thalia, Guardian of Thraben / Pippin, Guard of the Citadel / Kataki, War's Wage / None
- [catalog slot=power row=table_tolerance fit=0.05 filled=true] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts stax.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Thalia, Guardian of Thraben, Pippin, Guard of the Citadel, or Kataki, War's Wage? Say “none” and I’ll name three more.
  - Options: Thalia, Guardian of Thraben / Pippin, Guard of the Citadel / Kataki, War's Wage / None—name three more

### 25. poison in a tournament

Collection: false. Catalog: 2. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference for your infect deck?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your Modern event? I’ll tune the 15-card sideboard to it.

**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.


### 26. legacy for an event

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Legacy deck for an event.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the Legacy deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the Legacy deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the Legacy deck?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.

- [catalog slot=locked row=locked fit=0.90 filled=false] Must the deck keep Delver of Secrets, or may I cut cards that do not fit the blue-red tempo plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 3, the user:** FNM level, 600 dollars.


### 27. a card for the 99

Collection: true. Catalog: 8. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Grist, the Hunger Tide deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences for the deck?

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [catalog slot=commander row=named_card_role fit=0.90 filled=false] Would you like Grist, the Hunger Tide as your commander, or included in the 99?
  - Options: As my commander / In the 99
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power bracket does your Commander table use? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for the sacrifice deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=locked row=locked fit=0.20 filled=false] Must the deck keep Grist, the Hunger Tide and Grist, or may I cut a card that does not fit the sacrifice plan?
  - Options: Keep every card I named / Cut what does not fit
  - Refused as a reword (D-88): Must the deck keep Grist, the Hunger Tide, or may I cut a card that does not fit the plan?

**Turn 4, the user:** You pick the commander, and keep Grist in the deck.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Gríma Wormtongue, Old Man Willow, or The Sackville-Bagginses? Choose “None—name three more” and I’ll name three more.
  - Options: Gríma Wormtongue / Old Man Willow / The Sackville-Bagginses / None—name three more

### 28. reanimator with two plans

Collection: true. Catalog: 4. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your reanimator deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play for the reanimator deck? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your reanimator deck?

**Turn 2, the user:** Commander. Black and green.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your black-green Commander reanimator deck, should I use your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.


### 29. two hundred dollars

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table use: 2 (precon level), 3 (upgraded), 4 (high power), or another bracket?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. An artifact deck, blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=budget row=budget_scope fit=0.05 filled=false] Is your budget a cap on the cards you buy, or on the whole deck's total value?
  - Options: The cards I buy / The whole deck
  - Refused as a reword (D-88): What is the $200 budget a cap on: the cards you buy or the whole deck value?

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary, Katara, Bending Prodigy, or Yue, the Moon Spirit? Choose “None—name three more” if you want three additional options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Yue, the Moon Spirit / None—name three more

**Turn 4, the user:** The 200 caps the cards I buy.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary; Katara, Bending Prodigy; or Yue, the Moon Spirit? Choose “None—name three more” for three additional options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Yue, the Moon Spirit / None—name three more

### 30. another version, same plan

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the deck use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** The same plan with different cards.


