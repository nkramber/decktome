# PR-7 question gate

Run date: 2026-08-28. Conversations: 2026-08-26.

Verdict: PASS. 28 of 30 gate conversations used catalog questions only. The bar is 25.

75 probe conversations ran beside the gate. They asked 304 questions, and the model offered 16 replacements. A probe explores catalog coverage and does not move the verdict (D-96).

No conversation called itself complete with a slot unanswered.

The linter found no defective question (D-115).

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 133 |
| From the catalog | 131 |
| Invented by the model | 2 |
| Replacements refused as rewords | 5 |
| Catalog questions that closed a slot | 101 |
| Invented questions that closed a slot | 2 |
| Catalog-only conversations | 28 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `budget` | 1 |
| `house_rules` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 742. Time: 1127.3 seconds.
- Tokens: 743878 input (487307 cached), 65810 output.
- Cost: $0.1400.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Commander, and white and black is right.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Cecil, Dark Knight // Cecil, Redeemed Paladin, Denethor, Ruling Steward, or Cloak and Dagger, Entwined? Say "none" and I name three more.
  - Options: Cecil, Dark Knight // Cecil, Redeemed Paladin / Denethor, Ruling Steward / Cloak and Dagger, Entwined / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the Commander deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.


**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.


### 2. blink with a thin library

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.


**Turn 4, the user:** No more than 40 dollars.

- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is the $40 limit for the cards you buy or for the deck’s total value?
  - Options: Cards I buy / Whole deck

### 3. a named card, role unknown

Collection: true. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: theme (asked, no answer).

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme_card_named fit=0.90 filled=false] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=commander row=named_card_role fit=0.90 filled=true] Should Grist, the Hunger Tide be your commander or one of the other cards in the deck?
  - Options: Use it as the commander / Include it in the 99
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target?
  - Options: 1 — Exhibition / 2 — Core / 3 — Upgraded / 4 — Optimized / 5 — cEDH

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

### 4. fun and janky

Collection: true. Catalog: 7. Invented: 0.

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=power row=jank fit=0.90 filled=true] When you say “janky,” do you mean a low-power deck, an unexpected oddball card, or a plan that almost never works?
  - Options: A low-power deck / An odd card nobody expects / A plan that almost never works
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.

- [catalog slot=pool_rule row=pool fit=0.20 filled=true] Should I build the sacrifice deck from your library first, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 4, the user:** You pick the commander, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

### 5. 60-card anything goes

Collection: false. Catalog: 6. Invented: 1.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the 60-card deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on, such as a creature type, mechanic, or play style?
- [INVENTED slot=house_rules row=house_rules fit=0.20 filled=true] Besides allowing any card with no ban list, are there any other house rules?
  - Options: Any card, no ban list / No, I meant something else
  - It replaced: When you say anything goes, do you mean any card with no ban list?

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [catalog slot=house_rules row=house_format_limits fit=0.90 filled=true] Inside your house format, do the normal 60-card deck limits hold?
  - Options: Yes, the normal limits / No, I will say the limits
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the dragon deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM level / Tournament meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the dragon deck?

**Turn 3, the user:** The normal limits hold. Casual power, red and green.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

### 6. the strongest deck, no collection

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep it general / I’ll name the decks

**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.


### 7. FNM on Friday (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (never asked).

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=false] Which format does your FNM run: Standard, Modern, or Commander?
  - Options: Standard / Modern / Commander
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Pioneer. The best deck under budget.

- [catalog slot=format row=format_unsupported_open fit=0.05 filled=false] I do not build Pioneer. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
  - Refused as a reword (D-88): Which format should I build instead: Commander, Standard, or Modern?

**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


### 8. mill for a playgroup

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the mill deck?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget_scope fit=0.90 filled=true] Is your $100 budget a cap on the cards you buy, or on the whole deck's value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 — Exhibition / 2 — Core / 3 — Upgraded / 4 — Optimized / 5 — cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.


**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target: 2 (precon level), 3 (upgraded), or 4 (high power)?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.


**Turn 3, the user:** Add him to the buy list. Bracket 3, white and blue.


**Turn 4, the user:** Under 30 dollars, and build from my library first.

- [catalog slot=budget row=budget_scope fit=0.05 filled=false] Should the $30 limit apply to the cards you buy, or to the entire deck's value?
  - Options: Cards I buy / Entire deck
  - Refused as a reword (D-88): Is your $30 budget a cap on the cards you buy, or on the whole deck value?
  - The guard compared against: Is that a cap on the cards you buy, or on the whole deck value?

### 10. no strong commander in the library

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Commander. White and black.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.


### 11. another version after a build

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=plan_variant row=variance fit=0.90 filled=true] Would you like the same plan with different cards, or a different plan using the same colors?
  - Options: Same plan, different cards / A different plan

**Turn 2, the user:** A different plan in the same colors.


### 12. a slot change after a build

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Switch to owned-only.

- [catalog slot=plan_variant row=variance fit=0.90 filled=false] Would you like to keep the same plan with different cards, or switch to a different plan in the same colors?
  - Options: Same plan, different cards / A different plan

### 13. a card to keep that is not the commander

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 2, the user:** Commander, white and black, bracket 3.


**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

### 14. the user says none, then picks

Collection: false. Catalog: 6. Invented: 0.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Vito, Thorn of the Dusk Rose, Heliod, Sun-Crowned, or Haliya, Guided by Light? Say "none" and I name three more.
  - Options: Vito, Thorn of the Dusk Rose / Heliod, Sun-Crowned / Haliya, Guided by Light / None, name three more
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 3, the user:** None of those.

- [catalog slot=commander row=commander_pick fit=0.05 filled=true] Which one do you want: Aerith Gainsborough, Dina, Soul Steeper, or Celestine, the Living Saint? Say "none" and I name three more.
  - Options: Aerith Gainsborough / Dina, Soul Steeper / Celestine, the Living Saint / None, name three more
  - Refused as a reword (D-88): Which commander do you want: Aerith Gainsborough, Dina, Soul Steeper, or Celestine, the Living Saint?

**Turn 4, the user:** The first of the new three is good.


### 15. standard at the store, no library

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be?
  - Options: Casual / FNM-level / Tournament meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Standard. An aggro deck, red and white.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 3, the user:** FNM level.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks I expect to face

### 16. the strongest modern deck

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks I expect to face

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.


### 17. precon upgrade at bracket 2

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=pool_rule row=pool_precon fit=0.05 filled=true] Should I start with your precon, restrict the deck to cards from it, or build without using it?
  - Options: Start with the precon / Use only cards from the precon / Ignore the precon
  - Refused as a reword (D-88): Should I build from your Atraxa, Praetors' Voice precon first, or ignore it for a fully optimized deck?
  - The guard compared against: Should I build from your Atraxa, Praetors' Voice precon first, use only cards from it, or ignore it for a fully optimized deck?

**Turn 2, the user:** A superfriends deck. Build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

### 18. pauper burn on a budget (probe)

Collection: false. Catalog: 2. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer).

**Turn 1, the user:** A Pauper burn deck, as cheap as possible.

- [catalog slot=format row=format_unsupported_open fit=0.90 filled=false] I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=budget row=budget fit=0.90 filled=true] How much would you like to spend on cards?

**Turn 2, the user:** Pauper. Burn, mono red.


**Turn 3, the user:** Casual power, and 25 dollars is the cap.


### 19. a big library, no theme

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** I have a big collection. Build me something good.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Commander. A token deck, green and white.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Samwise Gamgee, Sam, Loyal Attendant, or Suki, Kyoshi Warrior? Say "none" and I name three more.
  - Options: Samwise Gamgee / Sam, Loyal Attendant / Suki, Kyoshi Warrior / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1: Exhibition / 2: Core / 3: Upgraded / 4: Optimized / 5: cEDH

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


### 20. dinosaur tribal for a child

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.


**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.

- [catalog slot=pool_rule row=pool_thin fit=0.90 filled=true] You have four Dinosaur cards in your library, while the deck needs at least 30. Should I prioritize cards you own and add a buy list, or use the entire card pool?
  - Options: Prioritize owned cards and add a buy list / Use the entire card pool

**Turn 4, the user:** Build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for buying additional cards?

### 21. vintage with proxies (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (never asked).

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=format row=format_unsupported_open fit=0.20 filled=false] I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
  - Refused as a reword (D-88): Which format should I build instead: Commander, Standard, or Modern?
- [catalog slot=house_rules row=house_format_limits fit=0.90 filled=true] Inside your house format, do the normal 60-card deck limits hold?
  - Options: Yes, the normal limits / No, I will say the limits

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. Blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 — Exhibition / 2 — Core / 3 — Upgraded / 4 — Optimized / 5 — cEDH
- [catalog slot=pool_rule row=pool_thin fit=0.90 filled=true] The card pool is thin for this theme. Should I build owned-first with a buy list, or use the whole pool?
  - Options: Owned-first with a buy list / Use the whole pool

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 4, the user:** My table is fine with extra turns.


### 23. land destruction

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. Red and green.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards from it, or ignore it for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Is there a budget for cards to buy?

**Turn 4, the user:** My table does not mind.


### 24. stax, owned only

Collection: true. Catalog: 4. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. White and blue.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.


### 25. poison in a tournament

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** An infect deck for a Modern tournament.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?
- [catalog slot=meta row=meta fit=0.90 filled=true] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks I expect to face

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.


### 26. legacy for an event (probe)

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer).

**Turn 1, the user:** A Legacy deck for an event.

- [catalog slot=format row=format_unsupported_open fit=0.90 filled=false] I do not build Legacy. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.


**Turn 3, the user:** FNM level, 600 dollars.


### 27. a card for the 99

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 Exhibition / 2 Core / 3 Upgraded / 4 Optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. A sacrifice deck, black and green.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for the deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=pool_rule row=pool fit=0.50 filled=true] Should I build from your library first, use only cards in your library, or ignore it for the deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 4, the user:** You pick the commander, and keep Grist in the deck.


### 28. reanimator with two plans

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. Black and green.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, restrict the deck to cards in your library, or ignore your library for an optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 4, the user:** Reanimate one big creature, not a value loop.


### 29. two hundred dollars

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: Exhibition / Core / Upgraded / Optimized / cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander. An artifact deck, blue and red.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a blue-red commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=budget row=budget_scope fit=0.20 filled=true] Is your budget a cap on the cards you buy, or on the value of the whole deck?
  - Options: The cards I buy / The whole deck
  - Refused as a reword (D-88): Is your $200 budget a cap on the cards you buy, or on the whole deck value?
  - The guard compared against: Is that a cap on the cards you buy, or on the whole deck value?

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** The 200 caps the cards I buy.


### 30. another version, same plan

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), theme (asked, no answer), colors (asked, no answer), power (never asked), pool_rule (never asked).

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** The same plan with different cards.


### 31. FNM on Friday, Modern

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=true] Which format does Friday’s event use: Standard, Modern, or Commander?
  - Options: Standard / Modern / Commander
- [catalog slot=theme row=theme_competitive fit=0.90 filled=true] Would you prefer a specific competitive deck, or the best option within your budget?
  - Options: A named deck / The best deck under budget
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?
- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks I expect to face

**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


### 32. burn on a budget

Collection: false. Catalog: 1. Invented: 1.

**Turn 1, the user:** A Standard burn deck, as cheap as possible.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How competitive should the Standard burn deck be: casual, FNM-level, or tournament meta?
  - Options: Casual / FNM-level / Tournament meta
- [INVENTED slot=budget row=budget fit=0.20 filled=true] What is your budget for buying cards?
  - It replaced: Is there a budget for cards to buy?

**Turn 2, the user:** Standard. Burn, mono red.


**Turn 3, the user:** Casual power, and 25 dollars is the cap.


### 33. a tempo deck for an event

Collection: false. Catalog: 6. Invented: 0.

**Turn 1, the user:** A Modern deck for an event.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be?
  - Options: Casual / FNM level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Modern. A Delver of Secrets tempo deck, blue and red.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?
- [catalog slot=locked row=locked fit=0.20 filled=false] Does the deck need to keep Delver of Secrets, or may I cut cards that do not fit the plan?
  - Options: Keep every card I named / Cut what does not fit
  - Refused as a reword (D-88): Should the deck keep Delver of Secrets, or may you cut a card that does not fit the plan?
  - The guard compared against: Must the deck keep Delver of Secrets, or may I cut a card that does not fit the plan?

**Turn 3, the user:** FNM level, 600 dollars.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep it general / I’ll name the decks

### 34. no ban list at our table

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** Any card, no ban list. Call it Modern. An artifact prison deck.

- [catalog slot=house_rules row=house_format_limits fit=0.90 filled=true] Inside your house format, do the normal 60-card deck limits hold?
  - Options: Yes, the normal limits / No, I will say the limits
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the artifact prison deck be?
  - Options: Casual / FNM-level / Tournament-meta

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 35. the user changes the format (probe)

Collection: false. Catalog: 7. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Commander, white and black.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=false] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 3, the user:** Actually, make it Modern instead. Sixty cards.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should this 60-card Modern deck be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta

**Turn 4, the user:** FNM level, and 100 dollars.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks

### 36. the user changes the theme late (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A sacrifice deck from my collection.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Denethor, Ruling Steward, Gríma Wormtongue, or Old Man Willow? Say "none" and I name three more.
  - Options: Denethor, Ruling Steward / Gríma Wormtongue / Old Man Willow / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 — exhibition / 2 — core / 3 — upgraded / 4 — optimized / 5 — cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander, black and green. Meren of Clan Nel Toth is the commander.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 3, the user:** I changed my mind, I want a token deck instead.


**Turn 4, the user:** Bracket 3, and build from my library first.


### 37. the user contradicts the budget (probe)

Collection: false. Catalog: 2. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** A competitive Modern deck, my budget is 50 dollars.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=meta row=meta fit=0.20 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks I expect to face
  - Refused as a reword (D-88): Which decks do you expect to face, or should I keep the sideboard general?

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

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** I dunno, you pick.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Peregrin Took, Joshua, Phoenix's Dominant // Phoenix, Warden of Fire, or Éowyn, Shieldmaiden? Say "none" and I name three more.
  - Options: Peregrin Took / Joshua, Phoenix's Dominant // Phoenix, Warden of Fire / Éowyn, Shieldmaiden / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck aim for? Choose 1 for exhibition, 2 for core, 3 for upgraded, 4 for optimized, or 5 for cEDH.
  - Options: Exhibition / Core / Upgraded / Optimized / cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] How should your library be used: as the starting point, as the only card pool, or not at all for a fully optimized deck?
  - Options: Use my library first / Use only my library / Ignore my library

**Turn 3, the user:** Whatever you think is best.


### 40. the user asks a question back (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want a Commander deck.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** What is a bracket? I have never heard of that.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for buying cards?

**Turn 3, the user:** Alright, bracket 2 then. A dragon deck, red and green.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for the dragon deck?
  - Options: I have a commander in mind / Suggest one

### 41. a request we can not serve (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** Can you build me a Yu-Gi-Oh deck?

- [catalog slot=scope row=out_of_scope fit=0.90 filled=true] I build Magic: The Gathering decks only. Would you like one instead?
  - Options: Yes, a Magic deck / No thanks

**Turn 2, the user:** Fine, Magic then. Commander, a dragon deck.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Lathliss, Dragon Queen, Atarka, World Render, or Miirym, Sentinel Wyrm? Say "none" and I name three more.
  - Options: Lathliss, Dragon Queen / Atarka, World Render / Miirym, Sentinel Wyrm / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck aim for? Choose from 1 (exhibition), 2 (core), 3 (upgraded), 4 (optimized), or 5 (cEDH).
  - Options: 1: Exhibition / 2: Core / 3: Upgraded / 4: Optimized / 5: cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 3, the user:** Bracket 3, red and green.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Atarka, World Render, Thrakkus the Butcher, or Dragonlord Atarka? Say "none" and I name three more.
  - Options: Atarka, World Render / Thrakkus the Butcher / Dragonlord Atarka / None, name three more
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

### 42. shorthand and slang (probe)

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** edh gruul dino stompy pls, no proxies

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Itzquinth, Firstborn of Gishath, Owen Grady, Raptor Trainer, or Etali, Primal Conqueror // Etali, Primal Sickness? Say "none" and I name three more.
  - Options: Itzquinth, Firstborn of Gishath / Owen Grady, Raptor Trainer / Etali, Primal Conqueror // Etali, Primal Sickness / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool_thin fit=0.90 filled=true] You have five Gruul dinosaur stompy cards, but the deck needs at least 30. Should I build around what you own and provide a buy list, or use the entire pool?
  - Options: Owned-first with a buy list / Use the whole pool

**Turn 2, the user:** bracket 3, from my binder

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** yeah 50 bucks max

- [catalog slot=budget row=budget_scope fit=0.20 filled=false] Does the $50 limit apply to the cards you purchase, or to the deck’s total value?
  - Options: The cards I buy / The whole deck

### 43. a card name with a typo (probe)

Collection: true. Catalog: 7. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** Build around Karlov of the Ghost Counsel.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander, white and black lifegain.

- [catalog slot=commander row=named_card_role fit=0.90 filled=false] Should Karlov of the Ghost Council lead the deck, or should it be included among the other cards?
  - Options: As my commander / In the 99
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power bracket should the deck target? Level 2 is preconstructed-deck strength, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] How should your existing card collection affect the build: prioritize it, use only it, or set it aside for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Yes that is the card I meant. Bracket 3, from my library.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

### 44. a card that is not legal (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** I want a Modern deck built around Black Lotus.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be?
  - Options: Casual / FNM level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Then pick something else. A burn deck, mono red.


**Turn 3, the user:** FNM level, 200 dollars.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks I expect to face

### 45. a commander that can not lead (probe)

Collection: false. Catalog: 6. Invented: 0.

**Turn 1, the user:** Commander deck with Lightning Bolt as my commander.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=commander row=commander_illegal fit=0.20 filled=true] Lightning Bolt can not lead a deck. Shall I suggest a commander instead?
  - Options: Yes, suggest one / No, I will name another
  - Refused as a reword (D-88): What should the commander be instead of Lightning Bolt?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** Alright, suggest a red commander for a burn deck.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Purphoros, God of the Forge, Ashling, Flame Dancer, or Dragonhawk, Fate's Tempest? Say "none" and I name three more.
  - Options: Purphoros, God of the Forge / Ashling, Flame Dancer / Dragonhawk, Fate's Tempest / None, name three more
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?
- [catalog slot=locked row=locked fit=0.05 filled=false] Must the deck keep Lightning Bolt, or may I cut cards that do not fit the plan?
  - Options: Keep every card you named / Cut cards that do not fit

**Turn 3, the user:** The first one. Bracket 3.


### 46. two partner commanders (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Commander deck with Thrasios, Triton Hero and Tymna the Weaver.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 Exhibition / 2 Core / 3 Upgraded / 4 Optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Yes, both as partners. A value deck.


**Turn 3, the user:** Bracket 5, cEDH. Money is no object.


### 47. a deck as a gift (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** I want to build a deck as a gift for my brother. He likes zombies.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Commander, black. He is new to the game.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Gisa, the Hellraiser, Zul Ashur, Lich Lord, or Hancock, Ghoulish Mayor? Say "none" and I name three more.
  - Options: Gisa, the Hellraiser / Zul Ashur, Lich Lord / Hancock, Ghoulish Mayor / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? Level 2 is precon-level, 3 is upgraded, and 4 is high power.
  - Options: Exhibition / Core / Upgraded / Optimized / cEDH

**Turn 3, the user:** Bracket 2, and under 80 dollars.


### 48. sideboard help only (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** I already have a Modern burn deck. I only need sideboard help.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the sideboard be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the sideboard?
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 2, the user:** Modern, burn, mono red. Tournament level.

- [catalog slot=meta row=meta fit=0.90 filled=true] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I'll name the decks

**Turn 3, the user:** My shop plays a lot of control.


### 49. copy a known deck (probe)

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: colors (asked, no answer).

**Turn 1, the user:** Build me the current best Standard deck, whatever is winning.

- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?
- [catalog slot=meta row=meta fit=0.20 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep it general / I’ll name the decks
  - Refused as a reword (D-88): Should I keep the sideboard general?

**Turn 2, the user:** Standard. Tournament level.


**Turn 3, the user:** 300 dollars, and my shop plays aggro.


### 50. a format we do not support (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want a Brawl deck for Arena.

- [catalog slot=format row=format_unsupported fit=0.90 filled=true] I do not build Brawl. The nearest format I build is Commander. Shall I use that?
  - Options: Yes, use the nearest format / No, I will name another format
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** It is like Commander but 60 cards on Arena.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 3, the user:** Fine, treat it as Commander. A dragon deck, red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a red commander in mind for the dragon deck, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the red dragon deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 4, the user:** Bracket 2.


### 51. the user declines everything (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** Build me a Commander deck.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target: 1 exhibition, 2 core, 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** You decide.

- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 3, the user:** You decide.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I start with cards from your library, use only cards from it, or ignore it for a fully optimized deck?
  - Options: Start with my library / Use only my library / Ignore my library

**Turn 4, the user:** You decide.


### 52. a price cap per card (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Modern deck where nothing costs more than 5 dollars a card.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be?
  - Options: Casual / FNM-level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Modern, a burn deck, mono red.


**Turn 3, the user:** Casual power, and 100 dollars total.


### 53. the user swaps the commander (probe)

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** A lifegain Commander deck from my library.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Aerith Gainsborough, Aerith, Last Ancient, or Hope Estheim? Say "none" and I name three more.
  - Options: Aerith Gainsborough / Aerith, Last Ancient / Hope Estheim / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target: 1 exhibition, 2 core, 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** White and black. Karlov of the Ghost Council is the commander.

- [catalog slot=budget row=budget fit=0.90 filled=false] Is there a budget for cards to buy?

**Turn 3, the user:** Actually use a different commander, suggest one.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Cecil, Dark Knight // Cecil, Redeemed Paladin, Denethor, Ruling Steward, or Cloak and Dagger, Entwined? Say "none" and I name three more.
  - Options: Cecil, Dark Knight // Cecil, Redeemed Paladin / Denethor, Ruling Steward / Cloak and Dagger, Entwined / None, name three more

**Turn 4, the user:** The second one. Bracket 3, library first.


### 54. two decks at once (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want two decks, one Commander and one Modern.

- [catalog slot=deck_count row=one_deck fit=0.05 filled=true] I build one deck at a time. Which deck do you want first?
  - Refused as a reword (D-88): Which deck do you want to build first: Commander or Modern?

**Turn 2, the user:** Start with the Commander one. A dragon deck, red and green.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for the dragon deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 55. a returning player (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I have not played since 2015 and I want to get back in.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like to build for: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What would you like the deck to focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander, I think. I like big green creatures.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.50 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Is there a budget for cards to buy?

**Turn 3, the user:** Green. Bracket 2, and about 100 dollars.


### 56. the user answers a different question (probe)

Collection: true. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** A mill deck from my collection.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like for the mill deck: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference for the mill deck?
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 2, the user:** My friends hate mill.


**Turn 3, the user:** Commander, blue and black. Bracket 3.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: The Mouth of Sauron, Hama, the Bloodbender, or Ultimecia, Time Sorceress // Ultimecia, Omnipotent? Say "none" and I name three more.
  - Options: The Mouth of Sauron / Hama, the Bloodbender / Ultimecia, Time Sorceress // Ultimecia, Omnipotent / None, name three more

**Turn 4, the user:** Build from my library first.


### 57. terse: the format is named once (probe)

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Red and green.

- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I build from your library first, use only cards from your library, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Is there a budget for cards to buy?

### 58. terse: a card for the 99 (probe)

Collection: true. Catalog: 6. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** A sacrifice deck, black and green.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a black-green commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should I start with cards from your library, use only cards from your library, or ignore your library for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 3, and build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 4, the user:** You pick the commander.


### 59. terse: two decks at once (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** I want two decks, one Commander and one Modern.

- [catalog slot=deck_count row=one_deck fit=0.05 filled=true] I build one deck at a time. Which deck do you want first?
  - Refused as a reword (D-88): Which deck do you want to build first: Commander or Modern?

**Turn 2, the user:** The Commander one. A dragon deck, red and green.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Atarka, World Render, Thrakkus the Butcher, or Dragonlord Atarka? Say "none" and I name three more.
  - Options: Atarka, World Render / Thrakkus the Butcher / Dragonlord Atarka / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the Commander deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 60. terse: a format we do not build (probe)

Collection: false. Catalog: 6. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** I want a Brawl deck for Arena.

- [catalog slot=format row=format_unsupported fit=0.90 filled=true] I do not build Brawl. The nearest format I build is Commander. Shall I use that?
  - Options: Yes, use the nearest format / No, I will name another format
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander is fine. A dragon deck, red.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=false] Is there a budget for cards to buy?

**Turn 3, the user:** Bracket 2.


### 61. terse: a proxy table (probe)

Collection: false. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer), colors (asked, no answer).

**Turn 1, the user:** We proxy everything at our table. A Vintage artifact prison deck.

- [catalog slot=format row=format_unsupported_open fit=0.20 filled=false] I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
  - Refused as a reword (D-88): I do not build Vintage. Which format should I build instead: Commander, Standard, or Modern?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference for the deck?

**Turn 2, the user:** Any card, no ban list.

- [catalog slot=house_rules row=house_format_limits fit=0.90 filled=true] Inside your house format, do the normal 60-card deck limits hold?
  - Options: Yes, the normal limits / No, I will say the limits

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 62. terse: a precon upgrade (probe)

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 2, the user:** A superfriends deck.


**Turn 3, the user:** Build from the precon first.


### 63. terse: a competitive request (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=meta row=meta fit=0.90 filled=true] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks I expect to face

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Yes, tournament level. My shop plays aggro and combo.


### 64. terse: a deck as a gift (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** I want to build a Commander deck as a gift for my brother. He likes zombies.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: The Scarab God, Wilhelt, the Rotcleaver, or Ratadrabik of Urborg? Say "none" and I name three more.
  - Options: The Scarab God / Wilhelt, the Rotcleaver / Ratadrabik of Urborg / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck aim for?
  - Options: 1 — exhibition / 2 — core / 3 — upgraded / 4 — optimized / 5 — cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Black. He is new to the game.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Gisa, the Hellraiser, Zul Ashur, Lich Lord, or Hancock, Ghoulish Mayor? Say "none" and I name three more.
  - Options: Gisa, the Hellraiser / Zul Ashur, Lich Lord / Hancock, Ghoulish Mayor / None, name three more
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 2, and under 80 dollars.


### 65. terse: a named commander and a locked card (probe)

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target?
  - Options: Exhibition / Core / Upgraded / Optimized / cEDH
- [catalog slot=pool_rule row=pool fit=0.90 filled=true] Should the deck use your library first, only cards from your library, or ignore your library?
  - Options: Library first / Only my library / Ignore my library

**Turn 2, the user:** Bracket 3.


**Turn 3, the user:** Build from my library first.

- [catalog slot=budget row=budget fit=0.90 filled=false] Is there a budget for cards to buy?

### 66. terse: sideboard help only (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** I already have a Modern burn deck, mono red. I only need sideboard help.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the sideboard be: casual, FNM level, or tournament-meta?
  - Options: Casual / FNM / Tournament-meta
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 2, the user:** Tournament level.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I'll name the decks

**Turn 3, the user:** My shop plays a lot of control.


### 67. terse: everything in one message (probe)

Collection: true. Catalog: 0. Invented: 0.

**This conversation asked nothing.** One message filled every slot, so it tests no question.

**Turn 1, the user:** Commander, bracket 3, Karlov of the Ghost Council, white and black lifegain, build from my library first, 50 dollar buy list.


### 68. terse: the strongest deck, no collection (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build the strongest Modern deck possible. I own nothing.

- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?
- [catalog slot=meta row=meta fit=0.90 filled=true] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I'll name the decks I expect to face

**Turn 2, the user:** A burn deck, mono red.


**Turn 3, the user:** Yes, tournament level. 300 dollars.


### 69. terse: a Pauper deck named once (probe)

Collection: false. Catalog: 2. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer).

**Turn 1, the user:** A Pauper burn deck, as cheap as possible. Mono red.

- [catalog slot=format row=format_unsupported_open fit=0.90 filled=false] I do not build Pauper. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=budget row=budget fit=0.90 filled=true] How much can you spend on cards for the deck?

**Turn 2, the user:** Casual power, and 25 dollars is the cap.


### 70. terse: a mill deck for a playgroup (probe)

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** A Commander mill deck for my playgroup, budget 100.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Hope Estheim, The Mouth of Sauron, or Saruman of Many Colors? Say "none" and I name three more.
  - Options: Hope Estheim / The Mouth of Sauron / Saruman of Many Colors / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Blue and black. Bruvac the Grandiloquent is the commander.

- [catalog slot=pool_rule row=pool fit=0.20 filled=true] Should I build from your library first, use only cards from it, or ignore it for a fully optimized deck?
  - Options: Library first / Only my library / Ignore my library
  - Refused as a reword (D-88): Do you want to build from your library first, use only your library, or ignore your library for a fully optimized deck?
  - The guard compared against: Build from your library first, only your library, or ignore it for a fully optimized deck?
- [catalog slot=budget row=budget_scope fit=0.20 filled=true] Is the budget a cap on the cards you buy or on the whole deck's value?
  - Options: The cards I buy / The whole deck
  - Refused as a reword (D-88): Is the $100 budget a cap on the cards you buy, or on the whole deck value?
  - The guard compared against: Is that a cap on the cards you buy, or on the whole deck value?

**Turn 3, the user:** Bracket 3, build from my library first. The 100 caps the cards I buy.


### 71. terse: one-word answers (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** Mono-white weenie deck.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 2, the user:** Modern.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be?
  - Options: Casual / FNM-level / Tournament-meta

**Turn 3, the user:** FNM.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face at FNM?
  - Options: Keep the sideboard general / I'll name the decks

**Turn 4, the user:** Control.


### 72. terse: a five-color deck (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A five-color Commander deck. Kenrith, the Returned King is the commander.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 Exhibition / 2 Core / 3 Upgraded / 4 Optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Good stuff, whatever is strongest.


**Turn 3, the user:** Bracket 4, and 400 dollars.


### 73. terse: a colorless deck (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A colorless Commander deck built around big artifacts.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Is there a budget for cards to buy?

**Turn 2, the user:** Kozilek, Butcher of Truth is the commander.


**Turn 3, the user:** Bracket 3, and 250 dollars.


### 74. terse: a card the user does not own (probe)

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A Commander deck from only the cards I own. Sol Ring goes in it.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What would you like the deck to focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** An artifact deck, colorless.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a colorless commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one

**Turn 3, the user:** Bracket 3.


**Turn 4, the user:** Buy it then.

- [catalog slot=budget row=budget fit=0.05 filled=false] Do you have a budget for cards to buy?
  - Refused as a reword (D-88): Is there a budget for cards to buy?
  - The guard compared against: Is there a budget for cards to buy?

### 75. terse: cEDH at bracket 5 (probe)

Collection: false. Catalog: 2. Invented: 0.

Slots the deck needs and nobody answered: theme (asked, no answer).

**Turn 1, the user:** A cEDH deck. Najeela, the Blade-Blossom.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=false] What should the deck focus on: a creature type, a mechanic, or a play style?
  - Options: Creature type / Mechanic / Play style
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Money is no object.


**Turn 3, the user:** Yes, bracket 5.


### 76. terse: a companion (probe)

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** A Modern deck with Lurrus of the Dream-Den as my companion.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM level / Tournament meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Mono black, aggro.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Tournament level, 300 dollars.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep it general / I’ll name the decks

### 77. terse: a Background pair (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** A Commander deck with a Background commander pair.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Jaheira, Friend of the Forest + Faceless One, Karlach, Fury of Avernus + Faceless One, or Lae'zel, Vlaakith's Champion + Faceless One? Say "none" and I name three more.
  - Options: Jaheira, Friend of the Forest + Faceless One / Karlach, Fury of Avernus + Faceless One / Lae'zel, Vlaakith's Champion + Faceless One / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: Exhibition / Core / Upgraded / Optimized / cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Red and white, aggressive.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 78. terse: partners named once (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** Krark, the Thumbless and Sakashima of a Thousand Faces as partners.

- [catalog slot=format row=format fit=0.90 filled=true] Which format would you like: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** A spellslinger deck, blue and red.


**Turn 3, the user:** Bracket 4, and 500 dollars.


### 79. terse: everything buried in one long message (probe)

Collection: true. Catalog: 3. Invented: 0.

Slots the deck needs and nobody answered: power (asked, no answer), commander (never asked).

**Turn 1, the user:** So I have been playing with my friends on Thursdays for about a year and we mostly play Commander, and I picked up a lot of black and green cards over that time, and what I really want is something that sacrifices creatures for value, nothing too strong because we are all pretty casual, maybe bracket 2 or 3, and I would rather use what I already have than buy much.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Old Man Willow, Killmonger, Scourge of Wakanda, or Jenova, Ancient Calamity? Say "none" and I name three more.
  - Options: Old Man Willow / Killmonger, Scourge of Wakanda / Jenova, Ancient Calamity / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=false] Which power bracket should the deck target?
  - Options: Bracket 1 / Bracket 2 / Bracket 3 / Bracket 4 / Bracket 5
- [catalog slot=budget row=budget fit=0.90 filled=false] How much would you like to spend on cards you do not already have?

### 80. terse: all caps, no punctuation (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** BUILD ME A MODERN BURN DECK MONO RED

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM level / Tournament meta
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 2, the user:** TOURNAMENT

- [catalog slot=meta row=meta fit=0.90 filled=true] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks

**Turn 3, the user:** MY SHOP PLAYS CONTROL


### 81. terse: the user corrects the colors (probe)

Collection: false. Catalog: 3. Invented: 0.

**Turn 1, the user:** A Commander lifegain deck, white and black.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 Exhibition / 2 Core / 3 Upgraded / 4 Optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=false] Is there a budget for cards to buy?

**Turn 2, the user:** Actually make it white and green instead.


**Turn 3, the user:** You pick the commander. Bracket 3.


### 82. terse: the user corrects a card name (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** Build around Atraxa, Praetor's Voice.

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=commander row=named_card_role fit=0.90 filled=true] Do you want Atraxa, Praetors Voice as your commander, or as one card in the 99?
  - Options: As my commander / In the 99
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH

**Turn 2, the user:** I meant Atraxa, Praetors' Voice. Commander, superfriends.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 3, and 300 dollars.


### 83. terse: group hug (probe)

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A group hug Commander deck. Phelddagrif is the commander.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** Bracket 2, and 100 dollars.


### 84. terse: voltron (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A voltron Commander deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** White and green. You pick the commander.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** The first one. Bracket 3, and 200 dollars.


### 85. terse: chaos (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A chaos Commander deck, as random as possible.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Vincent, Vengeful Atoner, Shadow the Hedgehog, or Be'lakor, the Dark Master? Say "none" and I name three more.
  - Options: Vincent, Vengeful Atoner / Shadow the Hedgehog / Be'lakor, the Dark Master / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] What power level should the deck target?
  - Options: Exhibition / Core / Upgraded / Optimized / cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Yidris, Maelstrom Wielder. Bracket 3.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** 150 dollars.


### 86. terse: enchantress (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** An enchantress deck from my collection.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 — exhibition / 2 — core / 3 — upgraded / 4 — optimized / 5 — cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander, white and green. Sythis, Harvest's Hand.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 3, library first.


### 87. terse: landfall (probe)

Collection: true. Catalog: 5. Invented: 0.

**Turn 1, the user:** A landfall Commander deck from my binder.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Rydia, Summoner of Mist, Ambrosia Whiteheart, or Choco, Seeker of Paradise? Say "none" and I name three more.
  - Options: Rydia, Summoner of Mist / Ambrosia Whiteheart / Choco, Seeker of Paradise / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target? Level 2 is precon-level, 3 is upgraded, and 4 is high-power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Green, mono.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Mole Man, Moloid Master, Ka-Zar of the Savage Land, or Toph, the Blind Bandit? Say "none" and I name three more.
  - Options: Mole Man, Moloid Master / Ka-Zar of the Savage Land / Toph, the Blind Bandit / None, name three more
- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 3, the user:** You pick the commander. Bracket 3, library first.


### 88. terse: no theme at all (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Commander deck. Surprise me.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** You decide.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** You decide.


### 89. terse: teaching a new player (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** A Commander deck for someone who has never played before.

- [catalog slot=theme row=theme fit=0.90 filled=true] What would you like the deck to do: feature a creature type, use a mechanic, or follow a particular play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Green, big creatures.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a green commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 1, and under 50 dollars.


### 90. terse: a deck for a spouse (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** My wife likes cats. Build her a Commander deck.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Arahbo, the First Fang, Ajani, Nacatl Pariah // Ajani, Nacatl Avenger, or Rin and Seri, Inseparable? Say "none" and I name three more.
  - Options: Arahbo, the First Fang / Ajani, Nacatl Pariah // Ajani, Nacatl Avenger / Rin and Seri, Inseparable / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 — exhibition / 2 — core / 3 — upgraded / 4 — optimized / 5 — cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** White and green. Bracket 2.

- [catalog slot=commander row=commander_pick fit=0.90 filled=false] Which one do you want: Arahbo, Roar of the World, Kaheera, the Orphanguard, or Kutzil, Malamet Exemplar? Say "none" and I name three more.
  - Options: Arahbo, Roar of the World / Kaheera, the Orphanguard / Kutzil, Malamet Exemplar / None, name three more
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Under 60 dollars.


### 91. terse: a budget of zero (probe)

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer), pool_rule (asked, no answer).

**Turn 1, the user:** A Commander deck and I can not spend anything.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 Exhibition / 2 Core / 3 Upgraded / 4 Optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Black and green, a sacrifice deck.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
- [catalog slot=pool_rule row=pool fit=0.90 filled=false] Should I use your library first, restrict the deck to cards you own, or build without considering your library?
  - Options: Library first / Only my library / Ignore my library

**Turn 3, the user:** Bracket 2.


### 92. terse: proxy the expensive cards only (probe)

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (never asked).

**Turn 1, the user:** A Commander deck. I proxy anything over 20 dollars.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
  - Options: Creature type / Mechanic / Play style
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Blue and black, a mill deck. Bruvac the Grandiloquent.

- [catalog slot=commander row=named_card_role fit=0.90 filled=false] Should Bruvac the Grandiloquent be your commander or one card in the 99?
  - Options: As the commander / In the 99
- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is your budget a cap on the cards you buy or on the deck's total value?
  - Options: The cards I buy / The whole deck

**Turn 3, the user:** Bracket 3, library first.


### 93. terse: two budget shapes at once (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Modern deck. Nothing over 5 dollars a card, and 120 dollars in total.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM level / Tournament meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** A burn deck, mono red.


**Turn 3, the user:** FNM level.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / I’ll name the decks

### 94. terse: a collection of lands only (probe)

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** I only own lands. Build me a Commander deck anyway.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
  - Options: A creature type / A mechanic / A play style
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target?
  - Options: 1 — Exhibition / 2 — Core / 3 — Upgraded / 4 — Optimized / 5 — cEDH

**Turn 2, the user:** Red and green, a dinosaur deck.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 3, and 200 dollars.

- [catalog slot=budget row=budget_scope fit=0.05 filled=false] Is the $200 a cap on the cards you buy, or on the whole deck's value?
  - Options: The cards I buy / The whole deck
  - Refused as a reword (D-88): Is the $200 budget a cap on the cards you buy, or on the whole deck value?
  - The guard compared against: Is that a cap on the cards you buy, or on the whole deck value?

### 95. terse: a library that does not exist (probe)

Collection: false. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** Build a Commander deck from my library.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** White and blue, a blink deck.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a commander in mind, or should I suggest one for the deck?
  - Options: I have a commander in mind / Suggest one
- [catalog slot=budget row=budget fit=0.90 filled=true] Is there a budget for cards to buy?

**Turn 3, the user:** Bracket 3.


### 96. terse: the user asks about a card (probe)

Collection: false. Catalog: 5. Invented: 1.

**Turn 1, the user:** Is Sol Ring legal in Modern?

- [catalog slot=theme row=theme_card_named fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be?
  - Options: Casual / FNM level / Tournament-meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Alright, a Modern artifact deck then, blue.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?
- [INVENTED slot=locked row=locked fit=0.20 filled=false] Should the deck keep every card you named, or may cards be cut if they do not fit the plan?
  - Options: Keep every card I named / Cut what does not fit
  - It replaced: Must the deck keep Sol Ring, or may I cut a card that does not fit the plan?

**Turn 3, the user:** FNM level, and 150 dollars.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep it general / I'll name the decks

### 97. terse: Pauper Commander (probe)

Collection: false. Catalog: 4. Invented: 0.

Slots the deck needs and nobody answered: format (asked, no answer).

**Turn 1, the user:** A Pauper Commander deck.

- [catalog slot=format row=format_unsupported_open fit=0.90 filled=false] I do not build Pauper Commander. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Black, a sacrifice deck.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards you’d need to buy?

**Turn 3, the user:** Casual power, and 40 dollars.


### 98. terse: Oathbreaker (probe)

Collection: false. Catalog: 5. Invented: 1.

Slots the deck needs and nobody answered: commander (asked, no answer).

**Turn 1, the user:** An Oathbreaker deck.

- [catalog slot=format row=format_unsupported fit=0.90 filled=true] I do not build Oathbreaker. The nearest format I build is Commander. Shall I use that?
  - Options: Yes, use the nearest format / No, I will name another format
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Commander is fine. Blue and red, a spellslinger deck.

- [INVENTED slot=commander row=commander fit=0.20 filled=false] Do you have a commander in mind, or should I suggest one?
  - Options: I have one in mind / Suggest one
  - It replaced: Do you have a commander in mind, or should I suggest one?
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power level should the deck target: 1 exhibition, 2 core, 3 upgraded, 4 optimized, or 5 cEDH?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Bracket 3, and 150 dollars.


### 99. terse: Historic on Arena (probe)

Collection: false. Catalog: 6. Invented: 0.

**Turn 1, the user:** A Historic deck for Arena.

- [catalog slot=format row=format_unsupported_open fit=0.20 filled=true] I do not build Historic. Which format should I build instead: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
  - Refused as a reword (D-88): Which format should I build instead: Commander, Standard, or Modern?
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Modern works. A mono-black aggro deck.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM level / Tournament meta
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** Tournament level, and 200 dollars.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep the sideboard general / Name the decks you expect to face

### 100. terse: a deck for a team event (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** I need a Modern deck for a team event on Saturday.

- [catalog slot=power row=power_sixty fit=0.90 filled=true] How strong should the deck be: casual, FNM level, or tournament meta?
  - Options: Casual / FNM level / Tournament meta
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 2, the user:** My teammates play aggro and combo, so I take control.


**Turn 3, the user:** Tournament level, 600 dollars, blue and white.

- [catalog slot=meta row=meta fit=0.90 filled=false] Should I keep the sideboard general, or will you name the decks you expect to face?
  - Options: Keep it general / I’ll name the decks

### 101. terse: mill with no table named (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Commander mill deck.

- [catalog slot=commander row=commander fit=0.90 filled=true] Do you have a commander in mind, or should I suggest one for the deck?
  - Options: I have one in mind / Suggest one
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Blue and black. You pick the commander.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 3, the user:** The second one. Bracket 3, and 100 dollars.


### 102. terse: stax with no table named (probe)

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** A stax Commander deck.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Wilson, Refined Grizzly, Gaddock Teeg, or Niv-Mizzet, Parun? Say "none" and I name three more.
  - Options: Wilson, Refined Grizzly / Gaddock Teeg / Niv-Mizzet, Parun / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: 1 — exhibition / 2 — core / 3 — upgraded / 4 — optimized / 5 — cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** White and blue. You pick the commander.

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 3, the user:** None of those.


**Turn 4, the user:** The first one. Bracket 4.


### 103. terse: two themes at once (probe)

Collection: false. Catalog: 4. Invented: 1.

**Turn 1, the user:** A Commander deck that does artifacts and lifegain.

- [catalog slot=commander row=commander_pick fit=0.90 filled=true] Which one do you want: Vito, Thorn of the Dusk Rose, Emry, Lurker of the Loch, or Danitha Capashen, Paragon? Say "none" and I name three more.
  - Options: Vito, Thorn of the Dusk Rose / Emry, Lurker of the Loch / Danitha Capashen, Paragon / None, name three more
- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target? Bracket 2 is precon level, 3 is upgraded, and 4 is high power.
  - Options: Exhibition / Core / Upgraded / Optimized / cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** White and blue. You pick the commander.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?
- [INVENTED slot=plan_variant row=plan_choice fit=0.20 filled=false] Would you like to choose between the two common plans for an artifacts-and-lifegain deck?
  - It replaced: This theme has two common plans. Which one do you want?

**Turn 3, the user:** The first one. Bracket 3, and 200 dollars.


### 104. terse: the user answers with a number (probe)

Collection: true. Catalog: 5. Invented: 0.

Slots the deck needs and nobody answered: power (asked, no answer), commander (asked, no answer).

**Turn 1, the user:** A Commander deck from my collection.

- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, mechanic, or play style?
- [catalog slot=power row=power_commander fit=0.90 filled=false] Which power bracket should the deck target?
  - Options: 1 exhibition / 2 core / 3 upgraded / 4 optimized / 5 cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** 3

- [catalog slot=budget row=budget fit=0.90 filled=false] Do you have a budget for cards to buy?

**Turn 3, the user:** Black and green, a sacrifice deck.

- [catalog slot=commander row=commander fit=0.90 filled=false] Do you have a black-green commander in mind, or should I suggest one?
  - Options: I have a commander in mind / Suggest one

**Turn 4, the user:** 2


### 105. probe: a theme the library can not lead (probe)

Collection: true. Catalog: 2. Invented: 0.

The slots were full after turn 3. The last 1 messages never went out.

**Turn 1, the user:** Build me an Eldrazi Commander deck using only the cards in my library.

- [catalog slot=power row=power_commander fit=0.90 filled=true] Which power bracket should the deck target: exhibition, core, upgraded, optimized, or cEDH?
  - Options: Exhibition / Core / Upgraded / Optimized / cEDH
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?

**Turn 2, the user:** Colorless.


**Turn 3, the user:** Bracket 3.


