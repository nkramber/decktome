# PR-8 deck gate, a rejudge of pr8-deck-gate-run31

Run date: 2026-09-21. Card snapshot: 2026-09-04. Decks read from `pr8-deck-gate-run31.md`.

Verdict: PASS. every deck of the source passed its build bars, 0 summaries stated a false rule of the game, and 0 judge calls failed. The build rows are the rows of the source run, and no deck was built again (D-789).

- Calls: 12. Cost: $0.3482. Time: 127 seconds.

## Run

- Suite `decks`, run `pr8-deck-gate-run33`, on 2026-09-21, commit `db2514e`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15, plan_rubric prompt version 4, summary_judge prompt version 2, kept_from `pr8-deck-gate-run32`, precons `5.3.0+20260903`, quality_model `20260914T154223Z`, rejudge_of `pr8-deck-gate-run31`.
- Calls: 12. Cost: $0.3482. Time: 127 seconds.
- Note: a rejudge of pr8-deck-gate-run31: the decks and the build rows are copied, and the judge rows are new (D-789). The judge rows of 19 decks are kept from pr8-deck-gate-run32, and its document holds their lines

## Decks

### 1. lifegain Commander, any card

**Summary:** This lifegain deck develops a resilient creature board, uses Karlov of the Ghost Council as a major source of pressure and control, and backs combat with several dedicated alternate ways to close the game. It has broad answers, card flow, and enough mana development to support its stronger threats, but it gives up some raw speed and can be vulnerable when its board is repeatedly cleared.

- PLAN plan_coherent=partly: The lifegain core (Karlov triggers, drain payoffs like Vito/Cliffhaven/Enduring Tenacity, life-per-draw engines) pulls one direction, but three symmetrical board wipes sit awkwardly in a wide creature deck and the ten-card "ramp" slot is largely marginal food/utility artifacts that barely produce mana.
- PLAN theme_fit=yes: It is a WB Commander deck led by Karlov of the Ghost Council built squarely around gaining life and converting it into damage, with a game-changer count and general tuning consistent with bracket 3.
- PLAN useful_as_built=yes: 36 lands with a strong WB mana base, a low curve, plenty of lifegain enablers and multiple closers (Aetherflux Reservoir, Vito, Approach, Test of Endurance, Angel of Destiny) make it playable out of the box.
- PLAN summary_honest=partly: The lifegain board, Karlov pressure, alt win conditions and card flow are all genuinely present, and it honestly admits the wipe vulnerability, but the claim of "enough mana development" overstates a rock package of Hot Dog Cart, Colossal Plow, Druidic Satchel and similar that mostly do not reliably ramp.

### 2. aristocrats Commander, owned first

**Summary:** This is a white-black aristocrats deck built to keep creatures flowing, turn their deaths into value, and use Denethor, Ruling Steward to make Soldier tokens and drain the table through sacrifice. Early mana and a large draw package keep the board supplied, while protection effects help Denethor and key creatures survive long enough to establish a sacrifice engine. The deck closes through steady life-loss pressure or its larger finishers, with broad removal and a few board wipes to recover from opposing boards. It gives up raw speed for a more incremental, board-dependent plan and can be vulnerable when its creatures or commander are repeatedly contained.

- PLAN plan_coherent=no: The list is really a white-black equipment/protection goodstuff pile — Puresteel Paladin, four hexproof/indestructible Equipment, Bastion Protector, Boromir, three board wipes — with almost no sacrifice outlets, no token generators for fodder, and only a couple of death-trigger payoffs, so the pieces pull against the stated sacrifice engine (the wipes actively fight the creature count).
- PLAN theme_fit=partly: It is a legal WB Commander deck under an aristocrats-capable commander, but the 80-card body delivers equipment, protection, and removal rather than the sacrifice-and-drain deck the player asked for, and bracket-3 goodstuff is the closest it gets.
- PLAN useful_as_built=yes: The mana base (24 basics plus fixing lands and nine rocks) supports the curve, removal and wipes are plentiful, and Angel of Serenity, Witch-king, Archfiend of Ifnir, and Vraska give real ways to close, so it can be picked up and played as a midrange deck.
- PLAN summary_honest=no: placeholder

### 3. artifacts Commander, bracket 4

**Summary:** Urza turns the early artifact mana into a fast, resilient board of artifact threats while holding up efficient protection and targeted answers. The deck wins by overwhelming the table with its artifact board, by deploying its dedicated artifact finishers, or through Thassa's Oracle. It gives up some flexibility for a dense artifact core, so its strongest games come from establishing mana quickly and protecting the key turn.

- PLAN plan_coherent=partly: The core is a consistent artifact plan — heavy zero-cost mana rocks, Urza's construct/mana engine, cost-reducers and artifact beaters — but it mixes cEDH-grade fast mana and free counters with filler bodies like Frogmite, Chrome Steed, and Foundry Assembler, and includes a Thassa's Oracle with no library-emptying enabler, so a few pieces pull away from the main line.
- PLAN theme_fit=yes: It is a mono-blue Urza, Lord High Artificer deck built around a dense artifact core with fast mana, free interaction and artifact payoffs, matching the requested high-power bracket-4 artifact brief.
- PLAN useful_as_built=yes: 33 lands plus a dozen cheap rocks support a low curve, and there are real win routes in the artifact creatures, Mechanized Production, Mirrodin Besieged and Urza tokens, so it plays fine off the shelf.
- PLAN summary_honest=partly: placeholder

### 4. dinosaur tribal, bracket 2

**Summary:** This deck ramps into a broad Dinosaur board, uses tribal support to make its creatures more effective, and lets Gishath, Sun's Avatar turn successful attacks into even more Dinosaurs. It wins primarily by overwhelming opponents in combat with large threats and its finisher cards, while retaining enough removal, protection, and wipes to keep the board manageable. It gives up fast, highly precise play for a straightforward creature plan with several expensive cards at the top of the curve.

- PLAN plan_coherent=yes: Every slot supports one line — ramp into big Dinosaurs, tribal cost reducers and anthems, and Gishath cheating more Dinosaurs off combat damage — with no competing subtheme pulling elsewhere.
- PLAN theme_fit=yes: It is a Gishath, Sun's Avatar Naya Dinosaur tribal Commander deck with straightforward creature-combat gameplay and no Game Changer-tier or combo pieces, matching the requested theme, commander, and bracket-2 casual level for a new player.
- PLAN useful_as_built=yes: 37 lands with a solid Naya fixing base plus about ten ramp pieces support the heavy top end, and the deck has ample redundant win paths in trample threats, Shared Animosity, and Gishath triggers, so it plays fine out of the box.
- PLAN summary_honest=yes: Each claim is backed by actual cards: ramp (Drover, Thunderherd Migration, signets), tribal support (Herald's Horn, Icon of Ancestry, Shared Animosity), removal (Trumpeting Carnosaur, Triumphant Chomp), protection (Heroic Intervention, boots/greaves), and wipes (Blasphemous Act, Harsh Mercy), and it openly admits the top-heavy curve.

### 5. blink Commander, owned first

**Summary:** Gilraen, Dúnedain Protector keeps a creature-heavy board moving through blink play, repeatedly leaning on creatures that supply draw, removal, and interaction. The deck develops through a substantial artifact ramp package, protects Gilraen and its key permanents, and uses broad reset buttons when the table gets out of hand. Angel of Serenity is chief among its finishers. The tradeoff is that the list is deliberately reactive and board-focused, with its closing power concentrated in a small finishing package rather than a fast combo.

- PLAN plan_coherent=partly: There is a recognizable white ETB-value-and-protect core, but the equipment/aura protection package (Buster Sword, Champion's Helm, Gift of Immortality, Darksteel Plate) and four sweepers pull against the small blink/ETB-creature subtheme rather than reinforcing it, and true blink enablers are few.
- PLAN theme_fit=partly: The commander is a blink-oriented white legend and the power level is plausibly bracket 3, but the actual list is mostly artifact ramp, equipment, protection and wipes with only a modest handful of blink payoffs (Angel of Sanctions, Fiend Hunter, Wall of Omens, Inspiring Overseer, Angel of Condemnation), so it is not really the blink deck asked for.
- PLAN useful_as_built=yes: Mono-white with 38 lands plus heavy cheap ramp, a smooth curve, ample removal and sweepers, and workable win routes in Angel of Serenity, Martial Coup and an equipped creature board — it can be picked up and played as is.
- PLAN summary_honest=partly: The ramp, protection, sweeper and Angel of Serenity claims all check out and the limited finishing power is honestly disclosed, but calling the list "creature-heavy" and built on "repeatedly leaning on" blink overstates a deck with roughly twenty creatures and only a handful of flicker effects.

### 6. Modern tempo, tournament

**Summary:** This blue-red tempo deck establishes pressure with a compact creature suite, then uses cheap removal and interaction to preserve that lead while cantrips keep the hand moving. It wins by maintaining steady pressure while denying the opponent room to stabilize. The tradeoff is a lean, low-end curve that gives up broader late-game power and more specialized answers.

- PLAN plan_coherent=yes: Every card serves one low-curve tempo plan: three cheap evasive/value threats, eight one-mana removal spells, six counterspells, and cantrips plus Bauble that both dig and connive Ledger Shredder.
- PLAN theme_fit=partly: It is a Modern blue-red tempo deck as asked, but the requested Delver archetype's namesake card (and any comparable one-mana flipping threat) is absent, so the threat base is Ragavan/Shredder/Mastermind rather than a true Delver shell.
- PLAN useful_as_built=yes: 24 lands with four duals, four fastlands, and two painlands comfortably cast everything on a curve topping at three, and twelve evasive creatures backed by burn and counters give a real clock and a coherent way to close.
- PLAN summary_honest=yes: The summary's claims — compact creature suite, cheap removal, counters, cantrips, and a lean curve short on late-game power — all match the 60 cards as listed.

### 7. Modern burn, casual

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 8. Modern lifegain, FNM

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 9. Standard midrange, FNM

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 10. Standard aggro, tournament

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 11. Commander with a locked card

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 12. Commander on a budget

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 13. owned first, and the commander is not owned

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 14. the user delegates the commander

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 15. delegated commander, owned first

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 16. a tight budget, owned first

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 17. upgrade a precon, owned first

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 18. upgrade a precon, any card

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 19. the Hobbit family, two colours

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 20. the Hobbit family, a delegated commander

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 21. the Hobbit family, mana from outside

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 22. a set family and a card from outside it

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 23. two set families at once

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 24. a 60-card deck from one set

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

### 25. use no card of an owned precon

Kept from pr8-deck-gate-run32, and its document holds the judge lines.

