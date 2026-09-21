# PR-8 deck gate, a rejudge of pr8-deck-gate-run31

Run date: 2026-09-21. Card snapshot: 2026-09-04. Decks read from `pr8-deck-gate-run31.md`.

Verdict: PASS. every deck of the source passed its build bars, 0 summaries stated a false rule of the game, and 0 judge calls failed. The build rows are the rows of the source run, and no deck was built again (D-789).

- Calls: 25. Cost: $0.4653. Time: 141 seconds.

## Run

- Suite `decks`, run `pr8-deck-gate-run34`, on 2026-09-21, commit `93a07a8`.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 15, plan_rubric prompt version 4, summary_judge prompt version 2, kept_from `pr8-deck-gate-run33`, precons `5.3.0+20260903`, quality_model `20260914T154223Z`, rejudge_of `pr8-deck-gate-run31`.
- Calls: 25. Cost: $0.4653. Time: 141 seconds.
- Note: a rejudge of pr8-deck-gate-run31: the decks and the build rows are copied, and the judge rows are new (D-789). The summary judge read every deck again, and the plan rows are kept from pr8-deck-gate-run33

## Decks

### 1. lifegain Commander, any card

**Summary:** This lifegain deck develops a resilient creature board, uses Karlov of the Ghost Council as a major source of pressure and control, and backs combat with several dedicated alternate ways to close the game. It has broad answers, card flow, and enough mana development to support its stronger threats, but it gives up some raw speed and can be vulnerable when its board is repeatedly cleared.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 2. aristocrats Commander, owned first

**Summary:** This is a white-black aristocrats deck built to keep creatures flowing, turn their deaths into value, and use Denethor, Ruling Steward to make Soldier tokens and drain the table through sacrifice. Early mana and a large draw package keep the board supplied, while protection effects help Denethor and key creatures survive long enough to establish a sacrifice engine. The deck closes through steady life-loss pressure or its larger finishers, with broad removal and a few board wipes to recover from opposing boards. It gives up raw speed for a more incremental, board-dependent plan and can be vulnerable when its creatures or commander are repeatedly contained.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 3. artifacts Commander, bracket 4

**Summary:** Urza turns the early artifact mana into a fast, resilient board of artifact threats while holding up efficient protection and targeted answers. The deck wins by overwhelming the table with its artifact board, by deploying its dedicated artifact finishers, or through Thassa's Oracle. It gives up some flexibility for a dense artifact core, so its strongest games come from establishing mana quickly and protecting the key turn.

The quality model grades this deck typical against the top lists of the format: the deck makes more mana on turn four than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 4. dinosaur tribal, bracket 2

**Summary:** This deck ramps into a broad Dinosaur board, uses tribal support to make its creatures more effective, and lets Gishath, Sun's Avatar turn successful attacks into even more Dinosaurs. It wins primarily by overwhelming opponents in combat with large threats and its finisher cards, while retaining enough removal, protection, and wipes to keep the board manageable. It gives up fast, highly precise play for a straightforward creature plan with several expensive cards at the top of the curve.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

- JUDGE [true]: "lets Gishath, Sun's Avatar turn successful attacks into even more Dinosaurs". This asserts a card ability: Gishath, Sun's Avatar reads that whenever it deals combat damage to a player, you reveal that many cards from the top of your library and put any Dinosaur cards revealed this way onto the battlefield tapped and attacking. So connecting in combat does yield more Dinosaurs.
- PLAN rows kept from pr8-deck-gate-run33.

### 5. blink Commander, owned first

**Summary:** Gilraen, Dúnedain Protector keeps a creature-heavy board moving through blink play, repeatedly leaning on creatures that supply draw, removal, and interaction. The deck develops through a substantial artifact ramp package, protects Gilraen and its key permanents, and uses broad reset buttons when the table gets out of hand. Angel of Serenity is chief among its finishers. The tradeoff is that the list is deliberately reactive and board-focused, with its closing power concentrated in a small finishing package rather than a fast combo.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

- PLAN rows kept from pr8-deck-gate-run33.

### 6. Modern tempo, tournament

**Summary:** This blue-red tempo deck establishes pressure with a compact creature suite, then uses cheap removal and interaction to preserve that lead while cantrips keep the hand moving. It wins by maintaining steady pressure while denying the opponent room to stabilize. The tradeoff is a lean, low-end curve that gives up broader late-game power and more specialized answers.

The quality model grades this deck typical against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The curve sits low for the format, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 7. Modern burn, casual

**Summary:** This mono-red burn deck applies pressure with efficient removal, a dense package of spell-focused creatures, and a steady stream of attackers. It wins by combining burn pressure with combat damage, using draw to keep action flowing into the middle turns. It gives up broad answers and matchup-specific tools for a focused, consistent red game plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 8. Modern lifegain, FNM

**Summary:** This white-black lifegain deck establishes early creatures, turns repeated life gain into increasingly serious combat pressure, and uses Vito, Thorn of the Dusk Rose to make that plan harder to race. Solitude and The Wandering Emperor clear away opposing threats while the draw package helps keep the board stocked. It gives up some speed and relies on its creatures surviving long enough for the lifegain payoffs to take over.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 9. Standard midrange, FNM

**Summary:** This black-green midrange deck develops its mana, establishes a steady creature board, and uses efficient removal to keep opposing threats from taking over. Insidious Roots and Corrupted Conviction support the creature-heavy core, while Darkstar Augur and Phyrexian Arena help maintain resources through longer games. It wins by applying sustained pressure with its creatures, chief among them Goldvein Hydra and Vein Ripper, but gives up some speed against decks built to end the game before its larger threats matter.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 10. Standard aggro, tournament

**Summary:** This red-white aggro deck aims to establish pressure early with a focused creature suite, then clear obstacles with removal and protect its tempo with interaction. Its draw cards help sustain that pressure when the first wave is answered, while Warleader's Call supports the attacking plan. It wins by keeping opponents under consistent combat pressure, giving up broader late-game flexibility for a streamlined, proactive approach.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 11. Commander with a locked card

**Summary:** Karlov of the Ghost Council leads a low-curve Orzhov sacrifice deck that builds resources from expendable creatures, protects its important pieces, and turns repeated creature losses into pressure. Karlov grows from the deck’s life-focused synergies and provides a dependable creature-control outlet, while sacrifice payoffs such as Bastion of Remembrance and Relic Vial help close games. The deck gives up some raw speed to keep a broad creature package, but it has strong staying power through recursion, card flow, and reset buttons.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 12. Commander on a budget

**Summary:** Adeline, Resplendent Cathar leads a white go-wide deck that attacks early, turns those attacks into a growing Human token force, and reinforces that force with token payoffs, creatures, and enchantments. The deck usually wins by overwhelming the table through combat, while Halo Fountain, Luck Bobblehead, and Sword of Body and Mind provide additional closing pressure. It gives up premium mana acceleration and expensive staple cards in favor of a broad, budget-conscious board-building plan.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 13. owned first, and the commander is not owned

**Summary:** Karlov of the Ghost Council is the center of a steady Orzhov lifegain plan: gain life repeatedly to grow Karlov, protect it, and use its creature-control pressure to keep attacks clear. The deck backs that plan with efficient removal, board resets, and a broad creature suite, then closes through Karlov's accumulated power or chief among its finishers, Frodo, Sauron's Bane, Grave Venerations, and Lyra Dawnbringer. It gives up explosive multicolor mana and relies on its basics, artifact acceleration, and steady card flow to win longer games.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few lands enter tapped, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 14. the user delegates the commander

**Summary:** This deck ramps through its early turns, uses tribal cost support to deploy Dragons efficiently, and then turns a wide aerial board into devastating attacks with Atarka, World Render. Lathliss, Dragon Queen, Utvara Hellkite, and Wrathful Red Dragon provide especially potent closing pressure, while the removal suite keeps opposing boards from stabilizing. It gives up some resilience to repeated sweepers and relies on its mana development and creature board to convert its powerful top end into a win.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The curve sits high for the format, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 15. delegated commander, owned first

**Summary:** Astarion, the Decadent leads a white-black lifegain deck that develops mana, establishes creatures and Equipment, and uses life gained during the turn to make the Friends choice at end step especially rewarding. The board is supported by efficient answers and reset buttons, then closes through its dedicated victory cards and finisher-class threats, with Astarion's Feed mode providing a way to press an opponent who has already lost life. The deck gives up some speed for a steady creature-based battlefield plan and a commander that costs six mana.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The land count sits above the norm of the format, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 16. a tight budget, owned first

**Summary:** Denethor, Ruling Steward anchors a creature-resource game: establish mana and card flow, develop a board of expendable creatures and sacrifice-focused pieces, then turn repeated creature losses into pressure. Bastion of Remembrance, Falkenrath Noble, and Zulaport Cutthroat provide the clearest finishing pressure alongside Denethor’s own sacrifice outlet. The deck trades explosive starts and premium individual card power for a patient, board-based plan that can rebuild through creature synergies and use wipes to reset opponents who get ahead.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- JUDGE [unknown]: "…alongside Denethor's own sacrifice outlet.". This asserts that the commander card Denethor, Ruling Steward has an ability that lets you sacrifice creatures — a claim about what a card may do. The provided card data lists only its mana cost and type line, and I cannot confirm from memory whether its printed text includes a sacrifice outlet.
- PLAN rows kept from pr8-deck-gate-run33.

### 17. upgrade a precon, owned first

**Summary:** This deck builds a wide Goblin board, then uses Zada to turn targeted spells into explosive team-wide turns. Goblin synergies create pressure while draw and burst mana help assemble a decisive combat, with Great Train Heist and Earthquake as major closing tools. It gives up broad flexibility for a focused Goblin-and-spells game plan that is strongest when its creature board remains established.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The cards pair the way the top lists pair them, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 18. upgrade a precon, any card

**Summary:** This keeps Turtle Power focused on a broad team of Turtles, Mutants, and Ninjas that wants to attack together, with Heroes in a Half Shell turning successful combat into larger attackers and more cards. The upgraded mana base and compact ramp package make the five-color plan steadier, while Rhystic Study and Kindred Discovery support a longer game. Dimension X Pizzasaur and Everything Pizza remain chief among the finishing threats, backed by efficient removal and the precon’s sweepers. The tradeoff is that the deck remains a combat-focused tribal build rather than a tightly streamlined combo deck.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 19. the Hobbit family, two colours

**Summary:** Thranduil leads a Sultai Elf creature deck that develops its mana early, builds pressure with a broad Elf board, and uses legendary Elves to fuel the commander’s draw-and-discard trigger. The deck keeps opponents in check with a substantial removal and interaction suite, then closes through Troll of Khazad-dûm or either Witch-king. Its main concession is that much of its pressure and synergy is creature-based, so repeated board clears can slow its ability to establish a decisive battlefield.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 20. the Hobbit family, a delegated commander

**Summary:** Smaug the Magnificent is the center of a red dragon-and-Treasure strategy that develops mana, attacks early, and turns its Treasure stockpile into mounting pressure. Cavern-Hoard Dragon, Desert Were-Worm, Smaug, the Great Calamity // Spew Flame, and chief among the supporting threats Dáin Ironfoot provide the closing force, while a broad spread of removal, interaction, and board resets keeps opponents from assembling an easy defense. The deck gives up some card selection and specialized answers in exchange for a direct, creature-driven plan built around Smaug’s attacks and a strongly themed supporting cast.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 21. the Hobbit family, mana from outside

**Summary:** This is a black-red Smaug deck that develops its mana with artifacts, Treasures, and Dragon-oriented acceleration before bringing Smaug the Impenetrable to the table as its central threat. It maintains pressure through Goblins, Orcs, Wargs, and larger monsters, uses a broad removal package to force attacks through, and closes with Smaug, the Great Calamity, Troll of Khazad-dûm, or Witch-king of Angmar. The deck gives up some early speed for a mana-heavy, thematic midrange plan built to support expensive Dragons and legendary equipment.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 22. a set family and a card from outside it

**Summary:** Thranduil leads a creature-centered Sultai deck built around Elves, with legendary Elves chief among them for recurring draw-and-discard triggers. Early mana creatures and artifacts help establish the board, while removal, interaction, and sweepers keep opponents from pulling too far ahead. The deck closes with Troll of Khazad-dûm and the Witch-kings after its Elf force has developed; it gives up some speed and consistency for a broad thematic creature package and a mana base that supports all three colors.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 23. two set families at once

**Summary:** Kíli leads a white Dwarf-and-Equipment deck that builds a durable artifact board, keeps cards flowing through its many value pieces, and turns even modest creatures into meaningful attackers. The deck wins through equipped combat, supported by large evasive closers and artifact-based pressure. It gives up multicolor flexibility and depends on keeping creatures and Equipment together, so its protective spells and selective removal are important for maintaining momentum.

The quality model grades this deck typical against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

This deck plan aims for 2 or more finishers, the cards that win the game, and the deck holds 1.

- PLAN rows kept from pr8-deck-gate-run33.

### 24. a 60-card deck from one set

**Summary:** This mono-red aggro deck aims to establish creature pressure early, reinforce it with Emberheart Challenger and Hearthborn Battler, and clear resistance with direct removal. Its threats carry the game through combat, with Dragonhawk, Fate's Tempest chief among them, while its draw package helps keep pressure coming. The deck gives up broad answers and defensive options in favor of a direct, consistent red attacking plan.

The quality model grades this deck at the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The color sources cover the pips, and that raises the grade.

- PLAN rows kept from pr8-deck-gate-run33.

### 25. use no card of an owned precon

**Summary:** Captain America, Super-Soldier anchors a white Avengers and S.H.I.E.L.D. Hero shell that develops a protected combat board, uses Equipment to keep the cards and mana flowing, and presses its advantage with superhero threats. The deck closes through Angel of Serenity, Origin of Spider-Man, and The Sentry, Golden Guardian, while its removal and wipes keep opposing boards from taking over. It gives up multicolor flexibility and depends heavily on maintaining creatures and artifacts on the table.

The quality model grades this deck at the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The deck makes less mana on turn four than the norm, and that lowers the grade.

- PLAN rows kept from pr8-deck-gate-run33.

