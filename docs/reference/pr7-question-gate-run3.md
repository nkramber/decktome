# PR-7 question gate

Run date: 2026-08-25. Conversations: 2026-08-24.

Verdict: FAIL. 23 of 30 conversations used catalog questions only. The bar is 25.

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 30 |
| Questions asked | 86 |
| From the catalog | 76 |
| Invented by the model | 10 |
| Catalog questions that closed a slot | 27 |
| Invented questions that closed a slot | 3 |
| Catalog-only conversations | 23 |
| Median gap score | 0.90 |
| Fit threshold | 0.35 |

### Invented questions by the row they replaced

A row that repeats here is a catalog change candidate (D-25, PR-15).

| Row | Count |
|---|---|
| `colors` | 3 |
| `acquisition` | 2 |
| `format` | 2 |
| `format_store` | 2 |
| `house_format_limits` | 1 |

## Run

- Roles: classify on `gpt-5.6-luna`, ask on `gpt-5.6-luna`.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Calls: 192. Time: 322.6 seconds.
- Tokens: 111542 input (0 cached), 16368 output.
- Cost: $0.0420.

## Conversations

### 1. lifegain with a collection

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** Build me a lifegain deck. I have a collection.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Commander, and white and black is right.


**Turn 3, the user:** Bracket 3. Suggest a commander from my library, and build from my library first.


**Turn 4, the user:** Karlov of the Ghost Council. Keep the buy list under 50 dollars.


### 2. blink with a thin library

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** I want a blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Blink is strongest in white, blue, black, and green.

**Turn 2, the user:** Commander. White and blue.


**Turn 3, the user:** Bracket 3, and Brago, King Eternal is my commander. Build owned-first with a buy list.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** No more than 40 dollars.

- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is the $40 cap for the cards you buy, or for the whole deck value?
  - Options: Cards I buy / Whole deck

### 3. a named card, role unknown

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Build around Grist, the Hunger Tide.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, or another format?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme_card_named fit=0.90 filled=false] What should the deck focus on: a creature type, a mechanic, or a particular play style?

**Turn 2, the user:** Commander. Black and green is fine.

- [catalog slot=theme row=theme fit=0.90 filled=false] What should the black-green Commander deck focus on: a creature type, a mechanic, or a play style?

**Turn 3, the user:** Grist is my commander. Bracket 3, and use my library first.


### 4. fun and janky

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Make me something fun and janky.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=power row=jank fit=0.90 filled=false] What does “janky” mean to you: a low-power deck, an unexpected oddball card, or a plan that almost never works?
  - Options: A low-power deck / An odd card nobody expects / A plan that almost never works
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference? Fun, janky decks are especially strong in blue, black, and green.

**Turn 2, the user:** Commander. Janky means an odd card nobody expects.


**Turn 3, the user:** A sacrifice deck. Bracket 2, and any colors are fine.


**Turn 4, the user:** You pick the commander, and build from my library first.


### 5. 60-card anything goes

Collection: false. Catalog: 4. Invented: 1.

**Turn 1, the user:** I want a 60-card deck, and anything goes at our table.

- [catalog slot=format row=format fit=0.90 filled=false] Since you want a 60-card deck, which format should it follow: Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—focus on a creature type, a mechanic, or a particular play style?
- [catalog slot=house_rules row=house_rules fit=0.20 filled=true] When you say anything goes, do you mean any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Modern. A dragon deck.

- [INVENTED slot=house_rules row=house_format_limits fit=0.05 filled=true] For your house format, should we use a 60-card minimum, up to four copies of each card name, and a 15-card sideboard?
  - Options: Yes, use those limits / No, different limits
  - It replaced: Inside your house format, do the normal limits hold: 60-card minimum, four copies per name, and a 15-card sideboard?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Dragon decks are strongest in blue and red.
  - Options: Blue-red / Another color combination / Any colors

**Turn 3, the user:** The normal limits hold. Casual power, red and green.


### 6. the strongest deck, no collection

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** Build the strongest deck possible. I own nothing.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference? For the strongest deck possible, blue, black, and green are recommended.
- [catalog slot=budget row=budget fit=0.90 filled=true] What budget do you have for buying cards?

**Turn 2, the user:** Modern. The best deck under budget.


**Turn 3, the user:** 300 dollars, tournament level. People at my shop play aggro and control.

- [catalog slot=format row=format_store fit=0.05 filled=false] Which format does your tournament run: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper

### 7. FNM on Friday

Collection: false. Catalog: 3. Invented: 2.

**Turn 1, the user:** I need a deck for FNM on Friday, something competitive.

- [catalog slot=format row=format_store fit=0.90 filled=false] Which format is your Friday FNM: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
- [catalog slot=theme row=theme_competitive fit=0.90 filled=true] Do you want a named tier-one deck or the best competitive deck within your budget?
  - Options: A named deck / The best deck under budget
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference?

**Turn 2, the user:** Pioneer. The best deck under budget.

- [INVENTED slot=format row=format fit=0.20 filled=false] You mentioned Pioneer—would you like to build for the Pioneer format?
  - It replaced: Which format: Commander, Standard, Modern, or something else?

**Turn 3, the user:** 200 dollars, tournament level. My shop is mostly aggro decks.


**Turn 4, the user:** I buy online, and I need the cards by Thursday.

- [INVENTED slot=budget row=acquisition fit=0.05 filled=false] What is your budget for the deck?
  - It replaced: Do you buy in person or online, and by what date do you need the cards?

### 8. mill for a playgroup

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Make a mill deck for my playgroup, budget 100.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Mill is strongest in blue, black, and green.
- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is your $100 budget a cap on the cards you buy or on the whole deck's value?
  - Options: The cards I buy / The whole deck

**Turn 2, the user:** Commander. Blue and black.


**Turn 3, the user:** Bracket 3. Bruvac the Grandiloquent is the commander, and build from my library first.


**Turn 4, the user:** My table accepts mill. The 100 caps the cards I buy. Mill as the win condition.


### 9. a commander the library does not hold

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Brago blink deck from my library.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like for your Brago blink deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, and Brago, King Eternal is the commander.


**Turn 3, the user:** Add him to the buy list. Bracket 3, white and blue.

- [catalog slot=budget row=budget fit=0.90 filled=true] Do you have a budget for cards to buy?

**Turn 4, the user:** Under 30 dollars, and build from my library first.

- [catalog slot=budget row=budget_scope fit=0.90 filled=false] Is the $30 limit for the cards you buy, or for the whole deck's value?
  - Options: Cards I buy / Whole deck

### 10. no strong commander in the library

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** Lifegain from my collection.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or another format?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Commander. White and black.


**Turn 3, the user:** Buy the best lifegain commander. Bracket 3, owned-first, and 60 dollars is the cap.


### 11. another version after a build

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Give me another version of that deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the new version use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck focus on—a creature type, mechanic, or play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** A different plan in the same colors.


### 12. a frozen run asks nothing

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Switch to owned-only.

- [catalog slot=format row=format fit=0.90 filled=false] Which format should the owned-only deck use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

### 13. a card to keep that is not the commander

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Karlov of the Ghost Council lifegain deck, and keep Sanguine Bond in it.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like for your Karlov of the Ghost Council lifegain deck: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Commander, white and black, bracket 3.


**Turn 3, the user:** Build from my library first, and lock Sanguine Bond in.


### 14. the user says none, then picks

Collection: false. Catalog: 2. Invented: 0.

**Turn 1, the user:** A lifegain commander deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference? Lifegain is strongest in white, black, and green.

**Turn 2, the user:** Bracket 3. I have no commander in mind, so suggest one.


**Turn 3, the user:** None of those.


**Turn 4, the user:** The first of the new three is good.


### 15. standard at the store, no library

Collection: false. Catalog: 4. Invented: 0.

**Turn 1, the user:** I need a Standard deck for my local store.

- [catalog slot=format row=format_store fit=0.90 filled=false] Which format does your local store’s event run: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Standard. An aggro deck, red and white.

- [catalog slot=format row=format fit=0.05 filled=false] Which format would you like to use: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 3, the user:** FNM level.


### 16. the strongest modern deck

Collection: false. Catalog: 2. Invented: 1.

**Turn 1, the user:** I want the strongest Modern deck, money is no object.

- [catalog slot=format row=format fit=0.05 filled=false] Which format would you like to build for: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [INVENTED slot=colors row=colors fit=0.05 filled=false] Do you have any color preference?
  - It replaced: Any color preference? the strongest Modern deck is strongest in blue, black, and green.

**Turn 2, the user:** The best deck under budget. Any colors.


**Turn 3, the user:** Tournament level, 800 dollars. My shop plays aggro and combo decks.

- [catalog slot=format row=format_store fit=0.05 filled=false] Which format does your tournament run: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper

### 17. precon upgrade at bracket 2

Collection: true. Catalog: 1. Invented: 0.

**Turn 1, the user:** Upgrade my Atraxa, Praetors' Voice precon. We play bracket 2.

- [catalog slot=format row=format fit=0.05 filled=false] Which format are you upgrading for: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** A superfriends deck. Build from my library first.


### 18. pauper burn on a budget

Collection: false. Catalog: 1. Invented: 0.

**Turn 1, the user:** A Pauper burn deck, as cheap as possible.

- [catalog slot=format row=format fit=0.05 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 2, the user:** Pauper. Burn, mono red.


**Turn 3, the user:** Casual power, and 25 dollars is the cap.


### 19. a big library, no theme

Collection: true. Catalog: 1. Invented: 1.

**Turn 1, the user:** I have a big collection. Build me something good.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [INVENTED slot=colors row=colors fit=0.20 filled=true] Are there any colors you’d like to play?
  - It replaced: Any color preference? the best deck under budget is strongest in white, blue, black, and green.

**Turn 2, the user:** Commander. A token deck, green and white.


**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


### 20. dinosaur tribal for a child

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** A dinosaur deck for my kid. Keep it casual.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Dinosaurs are strongest in red and green.

**Turn 2, the user:** Commander. Dinosaurs. Casual means low power, not a house format.


**Turn 3, the user:** Gishath, Sun's Avatar is the commander. Bracket 2, red green and white.


**Turn 4, the user:** Build from my library first.


### 21. vintage with proxies

Collection: false. Catalog: 5. Invented: 0.

**Turn 1, the user:** We proxy everything at our table.

- [catalog slot=format row=format fit=0.90 filled=false] Which format are you playing: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do—focus on a creature type, a mechanic, or a play style?
- [catalog slot=house_rules row=house_rules fit=0.90 filled=true] Since you proxy everything, what does “anything-goes” mean at your table: any card with no ban list, or Vintage rules?
  - Options: Any card, no ban list / Vintage rules

**Turn 2, the user:** Any card, no ban list. Call it Vintage. An artifact prison deck.

- [catalog slot=house_rules row=house_format_limits fit=0.90 filled=true] For your house Vintage format, should the normal limits apply: a 60-card minimum, four copies per card name, and a 15-card sideboard?
  - Options: Yes, the normal limits / No, I’ll specify the limits
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference, or should I choose the colors? Artifact prison is often strongest in blue and red.

**Turn 3, the user:** The normal limits hold. Casual power, mono brown.


### 22. extra turns commander

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** I want an extra turns deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, Pauper, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Extra-turns decks are strongest in blue and green.

**Turn 2, the user:** Commander. Blue and red.


**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** My table is fine with extra turns.


### 23. land destruction

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** A land destruction Commander deck.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Land destruction is strongest in blue, black, and green.

**Turn 2, the user:** Commander. Red and green.


**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** My table does not mind.


### 24. stax, owned only

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** A stax deck from only the cards I own.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Stax is strongest in white, blue, and green.

**Turn 2, the user:** Commander. White and blue.


**Turn 3, the user:** You pick the commander. Bracket 4, owned-only.


**Turn 4, the user:** My table accepts stax.


### 25. poison in a tournament

Collection: false. Catalog: 2. Invented: 1.

**Turn 1, the user:** An infect deck for a Modern tournament.

- [INVENTED slot=format row=format_store fit=0.20 filled=false] Which format is your event: Standard, Pioneer, Modern, or Pauper?
  - Options: Standard / Pioneer / Modern / Pauper
  - It replaced: Which format does your event run: Standard, Pioneer, Modern, or Pauper?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have a color preference? Infect is strongest in blue, black, and green.

**Turn 2, the user:** Modern. The best deck under budget.

- [catalog slot=format row=format fit=0.05 filled=false] You mentioned Modern—should the deck be for Modern, or another format?
  - Options: Commander / Standard / Modern / Pioneer / Pauper

**Turn 3, the user:** Tournament level, 400 dollars. My shop plays aggro and control decks.


### 26. legacy in three weeks

Collection: false. Catalog: 2. Invented: 3.

**Turn 1, the user:** A Legacy deck. I need it in three weeks for an event.

- [INVENTED slot=format row=format_store fit=0.20 filled=false] Legacy is noted—can you confirm that’s the format your event runs?
  - It replaced: Which format does your event run: Standard, Pioneer, Modern, or Pauper?
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Legacy. A Delver of Secrets tempo deck, blue and red.

- [INVENTED slot=format row=format fit=0.05 filled=false] Which format is your deck for: Commander, Standard, Modern, Legacy, or something else?
  - Options: Commander / Standard / Modern / Legacy / Something else
  - It replaced: Which format: Commander, Standard, Modern, or something else?

**Turn 3, the user:** FNM level, 600 dollars.


**Turn 4, the user:** I buy online, and I need the cards two days before the event.

- [INVENTED slot=budget row=acquisition fit=0.05 filled=false] What is your budget for the Delver of Secrets tempo deck?
  - It replaced: Do you buy in person or online, and by what date do you need the cards?

### 27. a card for the 99

Collection: true. Catalog: 1. Invented: 1.

**Turn 1, the user:** Build around Grist, the Hunger Tide, but not as my commander.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [INVENTED slot=colors row=colors fit=0.20 filled=true] Do you have any color preferences?
  - It replaced: Any color preference? Grist, the Hunger Tide is strongest in blue, black, red, and green.

**Turn 2, the user:** Commander. A sacrifice deck, black and green.


**Turn 3, the user:** Grist goes in the 99. Bracket 3, and build from my library first.


**Turn 4, the user:** You pick the commander, and keep Grist in the deck.


### 28. reanimator with two plans

Collection: true. Catalog: 2. Invented: 0.

**Turn 1, the user:** A reanimator Commander deck.

- [catalog slot=format row=format fit=0.20 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference? Reanimator is especially strong in white, black, and green.

**Turn 2, the user:** Commander. Black and green.


**Turn 3, the user:** Meren of Clan Nel Toth is the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** Reanimate one big creature, not a value loop.


### 29. two hundred dollars

Collection: true. Catalog: 4. Invented: 0.

**Turn 1, the user:** A Commander deck. I can spend 200 dollars, and I have a library.

- [catalog slot=format row=format fit=0.90 filled=false] You mentioned a Commander deck—should the format be Commander, Standard, Modern, Pioneer, or Pauper?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=true] What should the deck do: focus on a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have any color preference?

**Turn 2, the user:** Commander. An artifact deck, blue and red.

- [catalog slot=budget row=budget_scope fit=0.90 filled=true] Is your budget a cap on the cards you buy, or on the value of the whole deck?
  - Options: Cards I buy / Entire deck

**Turn 3, the user:** You pick the commander. Bracket 3, and build from my library first.


**Turn 4, the user:** The 200 caps the cards I buy.


### 30. another version, same plan

Collection: true. Catalog: 3. Invented: 0.

**Turn 1, the user:** Give me the same deck with other cards.

- [catalog slot=format row=format fit=0.90 filled=false] Which format would you like: Commander, Standard, Modern, Pioneer, or something else?
  - Options: Commander / Standard / Modern / Pioneer / Pauper
- [catalog slot=theme row=theme fit=0.90 filled=false] What should the deck focus on: a creature type, a mechanic, or a play style?
- [catalog slot=colors row=colors fit=0.90 filled=false] Do you have any color preference?

**Turn 2, the user:** The same plan with different cards.


