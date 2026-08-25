# PR-7 question gate

Run date: 2026-08-25. Conversations: 2026-08-24.

Verdict: PASS. 27 of 30 gate conversations used catalog questions only. The bar is 25.

22 probe conversations ran beside the gate. They asked 84 questions, and the model offered 6 replacements. A probe explores catalog coverage and does not move the verdict (D-96).

No conversation called itself complete with a slot unanswered.

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 132 |
| From the catalog | 128 |
| Invented by the model | 4 |
| Replacements refused as rewords | 7 |
| Catalog questions that closed a slot | 85 |
| Invented questions that closed a slot | 2 |
| Catalog-only conversations | 27 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `colors` | 1 |
| `house_rules` | 1 |
| `named_card_role` | 1 |
| `pool` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 391. Time: 688.0 seconds.
- Tokens: 260280 input (82017 cached), 35285 output.
- Cost: $0.0796.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your lifegain deck?

**Turn 2, the user:** Commander, and white and black is right.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for your white-black lifegain deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table play: 1 exhibition, 2 core (precon level), 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander do you want: Aerith Gainsborough, Minwu, White Mage, or Cecil, Dark Knight // Cecil, Redeemed Paladin? Choose “None—name three more” if you’d like three additional suggestions.
  - Options: Aerith Gainsborough / Minwu, White Mage / Cecil, Dark Knight // Cecil, Redeemed Paladin / None—name three more

**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.


### 2. blink with a thin library

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a preferred color identity?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a white-blue commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** No more than 40 dollars.


### 3. a named card, role unknown

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Grist, the Hunger Tide deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your Commander table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards from your library, or ignore your library for a fully optimized Grist, the Hunger Tide deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.


### 4. fun and janky

Collection: true. Catalog: 7. Invented: 0.

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
- [catalog slot=power row=power_commander fit=0.20 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
  - Refused as a reword (D-88): What power level does your table play: bracket 2 (precon level), bracket 3 (upgraded), or bracket 4 (high power)?

**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Denethor, Ruling Steward, Gríma Wormtongue, or Old Man Willow? Choose “None—name three more” if you want three additional suggestions.
  - Options: Denethor, Ruling Steward / Gríma Wormtongue / Old Man Willow / None—name three more

**Turn 4, the user:** You pick the commander, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.20 filled=false] Which commander do you want: The Sackville-Bagginses, June, Bounty Hunter, or Hei Bai, Spirit of Balance? Choose “None—name three more” and I’ll name three more.
  - Options: The Sackville-Bagginses / June, Bounty Hunter / Hei Bai, Spirit of Balance / None—name three more
  - Refused as a reword (D-88): Which commander do you want: The Sackville-Bagginses, June, Bounty Hunter, or Hei Bai, Spirit of Balance? Say "none" and I’ll name three more.

### 5. 60-card anything goes

Collection: false. Catalog: 3. Invented: 2.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—focus on a creature type, mechanic, or play style?
- [INVENTED slot=house_rules row=house_rules fit=0.20 filled=false] Which house rules should apply: any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules
  - It replaced: What does anything-goes mean at your table: any card with no ban list, or Vintage rules?

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should your dragon deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [INVENTED slot=colors row=colors fit=0.20 filled=true] Which colors would you like for your dragon deck?
  - It replaced: Any color preference? A a dragon deck deck is strongest in blue and red.

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

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the Modern deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=meta row=meta fit=0.20 filled=false] What do people play at your Modern event? I’ll tune the 15 sideboard cards to it.

**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.


### 7. FNM on Friday

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=true] Which format is your FNM: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** Pioneer. The best deck under budget.

- [catalog slot=meta row=meta fit=0.90 filled=false] What decks do people play at your Pioneer event? I’ll tune the 15-card sideboard to that metagame.

**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


### 8. mill for a playgroup

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your mill deck?
- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is your $100 budget a cap on the cards you buy, or on the whole deck's value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a blue-black commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your blue-black Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your blue-black mill deck, should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.

- [catalog slot=power row=table_tolerance fit=0.90 filled=false] For your Bruvac the Grandiloquent deck at Bracket 3, does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your Brago blink deck table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.


**Turn 3, the user:** Add him to the buy list. Bracket 3, white and blue.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for the cards to buy?

**Turn 4, the user:** Under 30 dollars, and build from my library first.


### 10. no strong commander in the library

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for your lifegain deck?

**Turn 2, the user:** Commander. White and black.

- [catalog slot=commander row=commander fit=0.05 filled=false] Do you have a white-black commander in mind, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Aerith Gainsborough, Minwu, White Mage, or Cecil, Dark Knight // Cecil, Redeemed Paladin? Choose “None—name three more” if you want three additional suggestions.
  - Options: Aerith Gainsborough / Minwu, White Mage / Cecil, Dark Knight // Cecil, Redeemed Paladin / None—name three more

### 11. another version after a build

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the new version use: Commander, Standard, Modern, Pioneer, or Pauper?
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
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a preferred color or color combination?

### 13. a card to keep that is not the commander

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Karlov of the Ghost Council lifegain deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, white and black, bracket 3.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your white-black Commander deck, should I use your library first, use only your library, or ignore it for a fully optimized build?
  - Options: Use my library first / Use only my library / Ignore my library
- [catalog slot=locked row=locked fit=0.20 filled=false] Should the deck keep Sanguine Bond, or may I cut it if it does not fit the lifegain plan?
  - Options: Keep every card I named / Cut what does not fit
  - Refused as a reword (D-88): Must the deck keep Sanguine Bond, or may I cut a card that does not fit the plan?

**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.


### 14. the user says none, then picks

Collection: false. Catalog: 4. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a lifegain commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your lifegain deck?

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander do you want: Vito, Thorn of the Dusk Rose, Heliod, Sun-Crowned, or Haliya, Guided by Light? Choose “None—name three more” and I’ll suggest three additional commanders.
  - Options: Vito, Thorn of the Dusk Rose / Heliod, Sun-Crowned / Haliya, Guided by Light / None—name three more

**Turn 3, the user:** None of those.


### 15. standard at the store, no library

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM level / Tournament meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences?

**Turn 2, the user:** Standard. An aggro deck, red and white.


**Turn 3, the user:** FNM level.


### 16. the strongest modern deck

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=power row=power_sixty fit=0.05 filled=true] How strong should your Modern deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
  - Refused as a reword (D-88): How strong should this be: casual, FNM level, or tournament-meta?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your Modern deck?
- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your Modern event? I’ll tune the 15-card sideboard to it.

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.


### 17. precon upgrade at bracket 2

Collection: true. Catalog: 0. Invented: 1.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [INVENTED slot=pool_rule row=pool fit=0.20 filled=true] Do you want to start with your current Atraxa precon, use only cards already in it, or ignore it for a fully optimized deck?
  - Options: Start with the current precon / Use only cards in the precon / Ignore the precon
  - It replaced: Build from your library first, only your library, or ignore it for a fully optimized deck?

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

**Turn 1, the user:** I have a big collection. Build me something good.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. A token deck, green and white.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your green-white token deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Peregrin Took, Rosie Cotton of South Lane, or Quina, Qu Gourmet? Choose “None—name three more” and I’ll name three more.
  - Options: Peregrin Took / Rosie Cotton of South Lane / Quina, Qu Gourmet / None—name three more

### 20. dinosaur tribal for a child

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference? Dinosaur decks are strongest in red and green.

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for your casual Dinosaur deck?
  - Options: I have one in mind / Suggest one

**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [catalog slot=pool_rule row=pool fit=0.05 filled=true] Should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library
  - Refused as a reword (D-88): Build from your library first, or ignore it for a fully optimized deck?

**Turn 4, the user:** Build from my library first.


### 21. vintage with proxies

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format are you playing: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—focus on a creature type, a mechanic, or a play style?
- [catalog slot=house_rules row=house_rules fit=0.50 filled=false] When you say you proxy everything, does your table allow any card with no ban list, or do you use Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should this Vintage artifact prison deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference? Artifact prison decks are strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 8. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Extra-turns decks are strongest in blue and green.

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for your blue-red deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your blue-red Commander table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build your blue-red extra-turns deck from your library first, only from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Birgi, God of Storytelling // Harnfel, Horn of Bounty; Emrakul, the Promised End; or Vnwxt, Verbose Host? Choose “None—name three more” if you want three additional options.
  - Options: Birgi, God of Storytelling // Harnfel, Horn of Bounty / Emrakul, the Promised End / Vnwxt, Verbose Host / None—name three more
- [catalog slot=power row=table_tolerance fit=0.90 filled=false] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table is fine with extra turns.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Birgi, God of Storytelling // Harnfel, Horn of Bounty, Emrakul, the Promised End, or Vnwxt, Verbose Host?
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

- [catalog slot=power row=table_tolerance fit=0.90 filled=false] Does your Commander table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Choose “None—name three more” and I’ll name three more.
  - Options: The Earth King / Toph, Hardheaded Teacher / Toph, the Blind Bandit / None—name three more

**Turn 4, the user:** My table does not mind.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: The Earth King, Toph, Hardheaded Teacher, or Toph, the Blind Bandit? Choose “None” and I’ll name three more.
  - Options: The Earth King / Toph, Hardheaded Teacher / Toph, the Blind Bandit / None

### 24. stax, owned only

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the stax deck?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your white-blue deck, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.05 filled=true] Which bracket does your white-blue table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.

- [catalog slot=commander row=commander_pick fit=0.05 filled=false] Which commander do you want: Thalia, Guardian of Thraben, Pippin, Guard of the Citadel, or Kataki, War's Wage? Choose “None—name three more” if you want three additional options.
  - Options: Thalia, Guardian of Thraben / Pippin, Guard of the Citadel / Kataki, War's Wage / None—name three more
- [catalog slot=power row=table_tolerance fit=0.90 filled=true] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only

**Turn 4, the user:** My table accepts stax.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Thalia, Guardian of Thraben, Pippin, Guard of the Citadel, or Kataki, War's Wage? Choose “None—name three more” if you want three additional suggestions.
  - Options: Thalia, Guardian of Thraben / Pippin, Guard of the Citadel / Kataki, War's Wage / None—name three more

### 25. poison in a tournament

Collection: false. Catalog: 2. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preferences for your infect deck?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your Modern event? I’ll tune the 15-card sideboard to it.

**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.


### 26. legacy for an event

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Legacy deck for an event.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the Legacy deck do: focus on a creature type, a mechanic, or a particular play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the Legacy deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM level / Tournament meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for the Legacy deck?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.

- [catalog slot=locked row=locked fit=0.90 filled=false] Must the deck keep Delver of Secrets, or may I cut cards that don’t fit the blue-red Legacy tempo plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 3, the user:** FNM level, 600 dollars.


### 27. a card for the 99

Collection: true. Catalog: 7. Invented: 1.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences?

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [INVENTED slot=commander row=named_card_role fit=0.20 filled=false] Do you want Grist, the Hunger Tide as your commander, or in the 99 with a different commander?
  - Options: Use Grist as commander / Put Grist in the 99
  - It replaced: Do you want Grist, the Hunger Tide as your commander, or as one card in the 99?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander fit=0.20 filled=true] Do you have a commander in mind, or should I suggest one for this Grist, the Hunger Tide deck?
  - Options: I have one in mind / Suggest one
  - Refused as a reword (D-88): Do you have a different commander in mind, or should I suggest one?
- [catalog slot=locked row=locked fit=0.90 filled=false] Must the deck keep Grist, the Hunger Tide, or may I cut a card that does not fit the plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 4, the user:** You pick the commander, and keep Grist in the deck.

- [catalog slot=commander row=commander_pick fit=0.05 filled=false] Which commander do you want: Grist, Voracious Larva // Grist, the Plague Swarm; Baron Sengir; or Galactus, Devourer of Worlds? Choose “None—name three more” and I’ll name three additional options.
  - Options: Grist, Voracious Larva // Grist, the Plague Swarm / Baron Sengir / Galactus, Devourer of Worlds / None—name three more
  - Refused as a reword (D-88): Which commander do you want: Baron Sengir or Galactus, Devourer of Worlds? Say "none" and I’ll name three more.

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

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] For your black-green Commander reanimator deck, should I build from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.


### 29. two hundred dollars

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. An artifact deck, blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind for your blue-red artifact deck, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=budget row=budget_scope fit=0.05 filled=false] Is your budget cap for the cards you buy, or for the whole blue-red artifact deck's total value?
  - Options: The cards I buy / The whole deck

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary; Katara, Bending Prodigy; or Yue, the Moon Spirit? Choose “None—name three more” if you want three other options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Yue, the Moon Spirit / None—name three more

**Turn 4, the user:** The 200 caps the cards I buy.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Glóin, Dwarf Emissary; Katara, Bending Prodigy; or Yue, the Moon Spirit? Choose “None—name three more” if you want three additional options.
  - Options: Glóin, Dwarf Emissary / Katara, Bending Prodigy / Yue, the Moon Spirit / None—name three more

### 30. another version, same plan

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** The same plan with different cards.


### 31. the user changes the format (probe)

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander, white and black.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a white-and-black commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=false] Which bracket does your white-and-black lifegain table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Actually, make it Modern instead. Sixty cards.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should this 60-card Modern lifegain deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta

**Turn 4, the user:** FNM level, and 100 dollars.


### 32. the user changes the theme late (probe)

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** A sacrifice deck from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Sacrifice decks are often strongest in black.

**Turn 2, the user:** Commander, black and green. Meren of Clan Nel Toth is the commander.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your black-green Meren of Clan Nel Toth Commander deck play at? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** I changed my mind, I want a token deck instead.


**Turn 4, the user:** Bracket 3, and build from my library first.


### 33. the user contradicts the budget (probe)

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** A competitive Modern deck, my budget is 50 dollars.

- [catalog slot=power row=power_sixty fit=0.20 filled=true] How strong should your competitive Modern deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference for your competitive Modern deck?
- [catalog slot=meta row=meta fit=0.90 filled=false] What do people play at your event? I’ll tune the 15-card sideboard to that metagame.

**Turn 2, the user:** Modern. Actually money is no object, spend what you need.


**Turn 3, the user:** Tournament level. My shop plays aggro.


### 34. everything in one message (probe)

Collection: true. Catalog: 0. Invented: 0.

**This conversation asked nothing.** One message filled every slot, so it tests no question.

**Turn 1, the user:** Commander, bracket 3, Karlov of the Ghost Council, white and black lifegain, build from my library first, 50 dollar buy list.


### 35. the user stays vague (probe)

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** Make me a good deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** I dunno, you pick.

- [catalog slot=commander row=commander fit=0.90 filled=true] Should I suggest a commander, or do you already have one in mind?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should I build for: 1 exhibition, 2 core (precon level), 3 upgraded, 4 optimized (high power), or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I use your library first, restrict the deck to only your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Whatever you think is best.

- [catalog slot=house_rules row=house_rules fit=0.90 filled=true] For your anything-goes table, do you mean any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules
- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander would you like, or should I suggest three more?

**Turn 4, the user:** Sure, that works.


### 36. the user asks a question back (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want a Commander deck.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—for example, focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 Exhibition / 2 Core / 3 Upgraded / 4 Optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences?

**Turn 2, the user:** What is a bracket? I have never heard of that.


**Turn 3, the user:** Alright, bracket 2 then. A dragon deck, red and green.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for your red-green dragon deck, or should I suggest one?
  - Options: I have a commander in mind / Suggest one

### 37. a request we can not serve (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Can you build me a Yu-Gi-Oh deck?

- [catalog slot=scope row=out_of_scope fit=0.90 filled=false] I build Magic: The Gathering decks only. Would you like a Magic deck instead?
  - Options: Yes, a Magic deck / No thanks

**Turn 2, the user:** Fine, Magic then. Commander, a dragon deck.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for your dragon deck, or should I suggest one?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your dragon deck? It's strongest in blue and red.

**Turn 3, the user:** Bracket 3, red and green.


### 38. shorthand and slang (probe)

Collection: true. Catalog: 4. Invented: 1.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** edh gruul dino stompy pls, no proxies

- [INVENTED slot=house_rules row=house_rules fit=0.05 filled=false] Do you have any house rules for your Gruul dinosaur stompy deck?
  - Options: Any card, no ban list / Vintage rules
  - It replaced: What does anything-goes mean at your table: any card with no ban list, or Vintage rules?
- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for your Gruul dinosaur stompy deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your table play: 1 exhibition, 2 core (precon level), 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** bracket 3, from my binder

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander would you like for your bracket 3 dinosaur stompy deck, or should I suggest three more?

**Turn 3, the user:** yeah 50 bucks max

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] With a $50 maximum, which commander do you want: Ghalta, Primal Hunger; Reptil, Dinomorpher; or Itzquinth, Firstborn of Gishath? Choose “None—name three more” if you want three additional options.
  - Options: Ghalta, Primal Hunger / Reptil, Dinomorpher / Itzquinth, Firstborn of Gishath / None—name three more

### 39. a card name with a typo (probe)

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build around Karlov of the Ghost Counsel.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for your Karlov of the Ghost Council deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, white and black lifegain.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket does your Commander table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Yes that is the card I meant. Bracket 3, from my library.


### 40. a card that is not legal (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I want a Modern deck built around Black Lotus.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the Black Lotus Modern deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Black Lotus decks are strongest in black.
- [catalog slot=locked row=locked fit=0.90 filled=false] Does the deck have to keep Black Lotus, or may I cut it if it does not fit the plan?
  - Options: Keep every card I named / Cut what does not fit

**Turn 2, the user:** Then pick something else. A burn deck, mono red.


**Turn 3, the user:** FNM level, 200 dollars.


### 41. a commander that can not lead (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** Commander deck with Lightning Bolt as my commander.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the Lightning Bolt deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play: 2 is precon level, 3 is upgraded, and 4 is high power?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** Alright, suggest a red commander for a burn deck.


**Turn 3, the user:** The first one. Bracket 3.


### 42. two partner commanders (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A Commander deck with Thrasios, Triton Hero and Tymna the Weaver.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** Yes, both as partners. A value deck.


**Turn 3, the user:** Bracket 5, cEDH. Money is no object.


### 43. a deck as a gift (probe)

Collection: false. Catalog: 4. Invented: 1.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want to build a deck as a gift for my brother. He likes zombies.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Does your brother have any color preference? Zombie decks are often strongest in blue and black.

**Turn 2, the user:** Commander, black. He is new to the game.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [INVENTED slot=power row=power_commander fit=0.20 filled=true] What power level should the black Zombie deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
  - It replaced: Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power.

**Turn 3, the user:** Bracket 2, and under 80 dollars.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which commander do you want: Gisa, the Hellraiser; Zul Ashur, Lich Lord; or Hancock, Ghoulish Mayor? Choose “None—name three more” if you’d like three other options.
  - Options: Gisa, the Hellraiser / Zul Ashur, Lich Lord / Hancock, Ghoulish Mayor / None—name three more

### 44. sideboard help only (probe)

Collection: false. Catalog: 2. Invented: 1.

**Turn 1, the user:** I already have a Modern burn deck. I only need sideboard help.

- [catalog slot=power row=power_sixty fit=0.20 filled=true] How strong should your Modern Burn sideboard be: casual, FNM-level, or tournament meta?
  - Options: Casual / FNM-level / Tournament meta
  - Refused as a reword (D-88): How strong should your sideboard be: casual, FNM level, or tournament-meta?
- [INVENTED slot=colors row=colors fit=0.20 filled=true] Do you have any color preference for your Modern Burn sideboard?
  - It replaced: Any color preference? A burn deck is strongest in red.

**Turn 2, the user:** Modern, burn, mono red. Tournament level.

- [catalog slot=meta row=meta fit=0.90 filled=false] What decks do people play at your tournament? I’ll tune the 15 sideboard cards to that metagame.

**Turn 3, the user:** My shop plays a lot of control.


### 45. copy a known deck (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** Build me the current best Standard deck, whatever is winning.

- [catalog slot=house_rules row=house_rules fit=0.05 filled=false] What does “anything-goes” mean at your table: any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules
- [catalog slot=power row=power_sixty fit=0.50 filled=true] How strong should this Standard deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM / Tournament meta
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference for the Standard deck?

**Turn 2, the user:** Standard. Tournament level.

- [catalog slot=meta row=meta fit=0.90 filled=false] What decks do people play at your Standard tournament? I’ll tune the 15-card sideboard to that metagame.

**Turn 3, the user:** 300 dollars, and my shop plays aggro.


### 46. a format we do not support (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want a Brawl deck for Arena.

- [catalog slot=format row=format fit=0.05 filled=true] Which format would you like for Arena: Brawl or something else?
  - Options: Brawl
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** It is like Commander but 60 cards on Arena.


**Turn 3, the user:** Fine, treat it as Commander. A dragon deck, red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for your red Dragon deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 4, the user:** Bracket 2.


### 47. the user declines everything (probe)

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** Build me a Commander deck.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—focus on a creature type, mechanic, or play style?
- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which commander would you like, or should I suggest three options?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** You decide.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 3, the user:** You decide.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards in your library, or ignore your library for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 4, the user:** You decide.


### 48. a price cap per card (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Modern deck where nothing costs more than 5 dollars a card.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a preferred color or color combination?

**Turn 2, the user:** Modern, a burn deck, mono red.


**Turn 3, the user:** Casual power, and 100 dollars total.


### 49. the user swaps the commander (probe)

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** A lifegain Commander deck from my library.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for your lifegain deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference for your lifegain deck?

**Turn 2, the user:** White and black. Karlov of the Ghost Council is the commander.


**Turn 3, the user:** Actually use a different commander, suggest one.


**Turn 4, the user:** The second one. Bracket 3, library first.


### 50. two decks at once (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want two decks, one Commander and one Modern.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should we focus on first: Commander, Modern, Standard, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] For the decks, what creature type, mechanic, or play style should they use?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences for either deck?

**Turn 2, the user:** Start with the Commander one. A dragon deck, red and green.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind for your red-green dragon deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which bracket does your table play for the Commander deck? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 51. a returning player (probe)

Collection: false. Catalog: 4. Invented: 1.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I have not played since 2015 and I want to get back in.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like to play: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What would you like the deck to focus on—a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preferences?

**Turn 2, the user:** Commander, I think. I like big green creatures.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or would you like me to suggest one for your big green creatures deck?
  - Options: I have one in mind / Suggest one
- [INVENTED slot=power row=power_commander fit=0.05 filled=true] What power level or bracket is your Commander table aiming for?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
  - It replaced: Which bracket does your table play? 2 is precon level, 3 is upgraded, 4 is high power.

**Turn 3, the user:** Green. Bracket 2, and about 100 dollars.


### 52. the user answers a different question (probe)

Collection: true. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A mill deck from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** My friends hate mill.


**Turn 3, the user:** Commander, blue and black. Bracket 3.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=table_tolerance fit=0.20 filled=false] Does your table accept mill, land destruction, extra turns, or stax?
  - Options: Yes / No / Mill only
  - Refused as a reword (D-88): What power level does your table accept for mill, land destruction, extra turns, or stax?

**Turn 4, the user:** Build from my library first.


