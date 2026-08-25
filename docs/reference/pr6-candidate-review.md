# PR-6 candidate review

Snapshot: 2026-08-24. Prompts: 2026-08-24. Collection: 2471 entries, 4316 cards, 1 rows unresolved.

Gate: for 20 theme prompts, a human confirms the top 40 candidates are on theme in at least 18. Ten run with no collection.

Score each prompt in the table, then fill the total.

| # | Theme | Mode | On theme (yes/no) | Notes |
|---|---|---|---|---|
| 1 | lifegain | owned-first | yes | 40 of 40. Clean list, payoff-heavy and enabler-light. Soul Warden, Exquisite Blood, and Authority of the Consuls are absent. Only 2 of 40 are owned. |
| 2 | aristocrats sacrifice | owned-first | no | 15 of 40. The death-trigger signal fires on any "when this dies" clause, so value creatures and artifacts fill the list. Blood Artist, Zulaport Cutthroat, Skullclamp, and Viscera Seer are absent. |
| 3 | tokens go-wide | owned-first | no | 25 of 40. The anthem matcher ignores the qualifier before "creatures you control get". Thirteen typal or conditional lords sit at ranks 9 to 37. No token doubler staple appears. |
| 4 | voltron equipment | owned-first | yes | 38 of 40. Greater Auramancy and Goro-Goro are the two misses, both close calls. The list is heavy on protection Equipment and has no tutor or draw engine. |
| 5 | blink | owned-first | no | 33 of 40. The theme keeps enabler tags only, because two payoff slugs do not exist in Tagger. Every ETB payoff is absent. Soulherder, Ghostly Flicker, and Panharmonicon miss the cut. |
| 6 | reanimator | owned-first | no | 34 of 40. Misses: Dredging Claw, Artisan of Kozilek, Drownyard Temple, Cauldron Familiar, The Soul Stone, Yawgmoth's Will. Reanimate, Animate Dead, Entomb, and Buried Alive are all absent. |
| 7 | landfall | owned-first | yes | 36 of 40, exactly at the bar. Two plain duals (Dreamroot Cascade at rank 1, Sodden Verdure) match "land enters" in their enters-tapped clause. Azusa and Ancient Greenwarden are absent. |
| 8 | spellslinger | owned-first | yes | 39 of 40. Will Kenrith is the one miss. 34 of 40 copy a spell. Magecraft counts as a keyword, so Storm-Kiln Artist and Archmage Emeritus fall past position 400. |
| 9 | dragons | owned-first | yes | 39 of 40. Clean Dragon typal list. Corroding Dragonstorm is the one miss (8% of Dragon decks). Crux of Fate, Haven of the Spirit Dragon, and Terror of the Peaks are absent. |
| 10 | artifacts | owned-first | yes | 37 of 40. Storm-Kiln Artist (rank 2) is a spellslinger card. Two generic rocks (The Mightstone and Weakstone, Solar Array) are off. No artifact land, cost reducer, or artifact cast trigger appears. |
| 11 | lifegain | any-card | yes | 40 of 40. The list is identical to prompt 1, card for card. Owned-first mode did not change the top 40. |
| 12 | mill | any-card | yes | 39 of 40. Jace, Wielder of Mysteries is a self-mill win condition. Eight-card-threshold text lifts weak Rogue bodies above Ruin Crab and Mesmeric Orb. Two text-match false positives. |
| 13 | elves | any-card | yes | 40 of 40. Clean Elf list. Llanowar Elves, Elvish Mystic, Beast Whisperer, and Reclamation Sage are absent because they carry no Elf-payoff text. |
| 14 | storm | any-card | no | 25 of 40. No Storm-keyword card appears (22 exist in UR). Second-spell token makers, colorless-matters cards, and a name match (Captain Lannery Storm) fill 15 slots. |
| 15 | enchantress | any-card | yes | 39 of 40. Clean enchantress list. Shambling Suit (artifact and enchantment count) is the one miss. Jukai Naturalist, Hall of Heliod's Generosity, and Enlightened Tutor are absent. |
| 16 | counters proliferate | any-card | no | 25 of 40. The counter-fuel tag fires on any counter type (wish, gold, energy, page, oil, quest, charge, mining). Fifteen such cards fill the list. Only one proliferate card appears. |
| 17 | zombies | any-card | yes | 39 of 40. Clean Zombie list. Liliana, the Last Hope is the one miss (Zombie text only in the emblem). Lord of the Undead is absent. |
| 18 | burn | any-card | no | 25 of 40. The damage tags fire on a sacrifice outlet, Equipment, Dragons, and a colorless planeswalker. EDHREC rank alone orders ranks 21 to 40, so a Modern prompt collects Commander staples. |
| 19 | control | any-card | no | 34 of 40. Six big creatures, Equipment, and Auras with "destroy target" or "exile target" text are off. Stock Up, Day of Judgment, Spell Snare, and Wan Shi Tong are absent. |
| 20 | infect poison | any-card | yes | 40 of 40. Every card has infect or toxic, or proliferates. The poison-counter signal fires on reminder text, so vanilla toxic bodies outrank Karn's Bastion and Evolution Sage. |

Total on theme: 12 of 20. All 20 prompts are scored. The gate needs 18, so this run fails on 2026-08-24.

Run 2 followed. It is `docs/reference/pr6-candidate-review-run2.md`, and it scores 20 of 20. This document stays as the record of run 1 (D-65).

Review method (2026-08-24): each card was checked against its Scryfall oracle text, format legality, and color identity. A card is on theme when its main function serves the theme (payoff, enabler, or staple), at any power level. A card whose theme link is incidental counts as on theme only when EDHREC shows the theme runs it for the theme: synergy of 0.10 or more on the theme's tag page, Top Commander status on that page, or presence in a format decklist (Modern, Standard). A card that fails format legality or color identity counts as off theme. Pass bar: 36 of 40 (owner decision, 2026-08-24). Prompts 1 and 11 share one card list. Each prompt section ends with a Review block that names the off-theme cards, the signal bugs, and the missing staples.

## 1. lifegain (commander, WB, owned-first)

Theme signals: payoff tags life-total-matters-self, lifegain-matters. tags drain-life, lifegain, repeatable-lifegain. keywords Lifelink.

Funnel: 12683 legal in colors, 1715 on theme, 478 owned, 121 on theme and owned, 224 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Aetherflux Reservoir | removal | 0 | 1.00 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain text:you gain tag:removal |
| 2 | Vito, Thorn of the Dusk Rose | synergy | 0 | 1.00 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 3 | Sanguine Bond | synergy | 0 | 1.00 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 4 | Enduring Tenacity | threat | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain type:threat |
| 5 | Serra Ascendant | synergy | 0 | 0.99 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain keyword:Lifelink theme |
| 6 | The Wind Crystal | synergy | 0 | 0.99 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 7 | Heliod, Sun-Crowned | synergy | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 8 | Marauding Blight-Priest | synergy | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 9 | Ocelot Pride | synergy | 0 | 0.99 | payoff:lifegain-matters payoff-text:if you gained life this turn tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain theme |
| 10 | Well of Lost Dreams | draw | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life payoff-text:life you gained text:you gain tag:draw |
| 11 | Exemplar of Light | draw | 3 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain tag:draw |
| 12 | Caduceus, Staff of Hermes | interaction | 0 | 0.99 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain tag:interaction |
| 13 | Archangel of Thune | threat | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain type:threat |
| 14 | Cleric Class | synergy | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain text:you gain theme |
| 15 | Haliya, Guided by Light | draw | 0 | 0.98 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain text:you gain tag:draw |
| 16 | Felidar Sovereign | wincon | 0 | 0.98 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain keyword:Lifelink tag:alternate-win-condition |
| 17 | Elenda's Hierophant | synergy | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain text:you gain theme |
| 18 | Starscape Cleric | synergy | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 19 | Righteous Valkyrie | synergy | 0 | 0.98 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 20 | Resplendent Angel | synergy | 0 | 0.98 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 21 | Aerith Gainsborough | synergy | 1 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain theme |
| 22 | Dawn of Hope | draw | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain text:you gain tag:draw |
| 23 | Witch of the Moors | removal | 0 | 0.98 | payoff:lifegain-matters payoff-text:if you gained life this turn text:you gain tag:removal |
| 24 | Ajani's Pridemate | synergy | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 25 | Cosmos Elixir | draw | 0 | 0.97 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain text:you gain tag:draw |
| 26 | Valkyrie Harbinger | threat | 0 | 0.97 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain type:threat |
| 27 | Celestine, the Living Saint | threat | 0 | 0.97 | payoff:lifegain-matters payoff-text:life you gained tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain type:threat |
| 28 | Angel of Destiny | wincon | 0 | 0.97 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain tag:alternate-win-condition |
| 29 | Nykthos Paragon | threat | 0 | 0.97 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain type:threat |
| 30 | Lunar Convocation | draw | 0 | 0.97 | payoff:lifegain-matters payoff-text:if you gained life this turn text:you gain tag:draw |
| 31 | Voice of the Blessed | synergy | 0 | 0.97 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 32 | Ajani, Strength of the Pride | wipe | 0 | 0.97 | payoff:lifegain-matters payoff:life-total-matters-self payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain text:you gain tag:sweeper |
| 33 | Astarion, the Decadent | threat | 0 | 0.97 | payoff:lifegain-matters payoff-text:life you gained tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain type:threat |
| 34 | Speaker of the Heavens | synergy | 0 | 0.96 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain keyword:Lifelink theme |
| 35 | Karlov of the Ghost Council | removal | 0 | 0.96 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain tag:removal |
| 36 | Sorin of House Markov // Sorin, Ravenous Neonate | removal | 0 | 0.96 | payoff:lifegain-matters payoff-text:life you gained tag:lifegain tag:repeatable-lifegain tag:drain-life keyword:Lifelink text:you gain tag:removal |
| 37 | Essence Channeler | synergy | 0 | 0.96 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 38 | Amalia Benavides Aguirre | wipe | 0 | 0.96 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain tag:sweeper |
| 39 | Veinwitch Coven | synergy | 0 | 0.96 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 40 | Indulging Patrician | synergy | 0 | 0.96 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain theme |

Top upgrades (unowned):

- Aetherflux Reservoir (removal, 1.00)
- Vito, Thorn of the Dusk Rose (synergy, 1.00)
- Sanguine Bond (synergy, 1.00)
- Enduring Tenacity (threat, 0.99)
- Serra Ascendant (synergy, 0.99)
- The Wind Crystal (synergy, 0.99)
- Heliod, Sun-Crowned (synergy, 0.99)
- Marauding Blight-Priest (synergy, 0.99)
- Ocelot Pride (synergy, 0.99)
- Well of Lost Dreams (draw, 0.99)

### Review

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Lifegain and Lifedrain pages (edhrec.com/tags/lifegain, /lifedrain).

Off theme: none. Thirty-one cards carry a literal lifegain trigger. The other nine are life-total payoffs (Aetherflux Reservoir, Serra Ascendant, Caduceus, Felidar Sovereign, Righteous Valkyrie, Cosmos Elixir, Angel of Destiny, Speaker of the Heavens) or a lifegain doubler (The Wind Crystal). Weakest fits: Caduceus (11% of Lifegain decks), Cosmos Elixir (10%), Ajani, Strength of the Pride (9%), and Ocelot Pride (not on the page, mostly a token card).

Signal bugs: Caduceus gets tag:interaction but removes, counters, and taxes nothing. The Wind Crystal and Vito get tag:repeatable-lifegain, but neither gains life by itself. Aetherflux Reservoir gets role removal while Felidar Sovereign and Angel of Destiny get role wincon for the same job.

Legality: all 40 are legal in commander and inside WB.

Owned: 2 of 40 are owned (Exemplar of Light, Aerith Gainsborough). The funnel reports 121 owned on-theme cards.

Gaps (context only): the archetype's enablers are absent. Soul Warden (40% of Lifegain decks), Exquisite Blood (37%), Authority of the Consuls (33%), Soul's Attendant (31%), Kambal, Consul of Allocation (26%), Bloodthirsty Conqueror (24%), Rhox Faithmender (20%), and Suture Priest (20%).

## 2. aristocrats sacrifice (commander, WB, owned-first)

Theme signals: payoff tags blood-artist-ability, death-trigger, sacrifice-matters, your-sacrifice-matters. tags death-trigger-self, repeatable-sacrifice-outlet, sacrifice-outlet.

Funnel: 12683 legal in colors, 2686 on theme, 471 owned, 172 on theme and owned, 217 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Solemn Simulacrum | ramp | 0 | 1.00 | payoff:death-trigger tag:death-trigger-self tag:ramp |
| 2 | Syr Konrad, the Grim | threat | 0 | 1.00 | payoff:death-trigger payoff-text:whenever another creature dies type:threat |
| 3 | Bastion of Remembrance | synergy | 0 | 1.00 | payoff:death-trigger payoff:blood-artist-ability payoff-text:whenever a creature you control dies theme |
| 4 | The Meathook Massacre | wipe | 0 | 1.00 | payoff:death-trigger payoff:blood-artist-ability payoff-text:whenever a creature you control dies tag:sweeper |
| 5 | Grave Pact | removal | 0 | 0.99 | payoff:death-trigger payoff-text:whenever a creature you control dies text:sacrifice tag:removal |
| 6 | Stitcher's Supplier | synergy | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self theme |
| 7 | Enduring Innocence | draw | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self tag:draw |
| 8 | Liliana, Dreadhorde General | wipe | 0 | 0.99 | payoff:death-trigger payoff-text:whenever a creature you control dies text:sacrifice tag:sweeper |
| 9 | Enduring Tenacity | threat | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self type:threat |
| 10 | Myr Retriever | synergy | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self theme |
| 11 | Ojer Taq, Deepest Foundation // Temple of Civilization | land | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self type:land |
| 12 | Ichor Wellspring | draw | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self tag:draw |
| 13 | Dictate of Erebos | removal | 0 | 0.99 | payoff:death-trigger payoff-text:whenever a creature you control dies text:sacrifice tag:removal |
| 14 | Yahenni, Undying Partisan | synergy | 0 | 0.99 | payoff:death-trigger tag:sacrifice-outlet tag:repeatable-sacrifice-outlet text:sacrifice theme |
| 15 | Wurmcoil Engine | threat | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self type:threat |
| 16 | Junji, the Midnight Sky | threat | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self type:threat |
| 17 | Elenda, the Dusk Rose | threat | 0 | 0.99 | payoff:death-trigger payoff-text:whenever another creature dies tag:death-trigger-self type:threat |
| 18 | Nuka-Cola Vending Machine | ramp | 0 | 0.99 | payoff:sacrifice-matters payoff:your-sacrifice-matters payoff-text:whenever you sacrifice text:sacrifice tag:ramp |
| 19 | Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | draw | 0 | 0.99 | payoff:death-trigger payoff:blood-artist-ability payoff-text:whenever another creature dies tag:sacrifice-outlet tag:repeatable-sacrifice-outlet text:sacrifice tag:draw |
| 20 | Hangarback Walker | synergy | 1 | 0.99 | payoff:death-trigger tag:death-trigger-self theme |
| 21 | Pawn of Ulamog | ramp | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self text:sacrifice tag:ramp |
| 22 | Unstoppable Slasher | synergy | 0 | 0.99 | payoff:death-trigger tag:death-trigger-self theme |
| 23 | Murderous Rider // Swift End | removal | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self tag:removal |
| 24 | Elenda's Hierophant | synergy | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self theme |
| 25 | Junk Diver | synergy | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self theme |
| 26 | Funeral Room // Awakening Hall | synergy | 0 | 0.98 | payoff:death-trigger payoff:blood-artist-ability payoff-text:whenever a creature you control dies theme |
| 27 | Triplicate Titan | threat | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self type:threat |
| 28 | Aerith Gainsborough | synergy | 1 | 0.98 | payoff:death-trigger tag:death-trigger-self theme |
| 29 | Kokusho, the Evening Star | threat | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self type:threat |
| 30 | Prized Statue | ramp | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self text:sacrifice tag:ramp |
| 31 | Greedy Freebooter | ramp | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self text:sacrifice tag:ramp |
| 32 | Mycosynth Wellspring | ramp | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self text:mana |
| 33 | Crowded Crypt | ramp | 0 | 0.98 | payoff:death-trigger payoff-text:whenever a creature you control dies text:sacrifice tag:ramp |
| 34 | God-Eternal Oketra | threat | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self type:threat |
| 35 | Nine-Lives Familiar | synergy | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self theme |
| 36 | Ancient Stone Idol | threat | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self type:threat |
| 37 | Aclazotz, Deepest Betrayal // Temple of the Dead | land | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self type:land |
| 38 | Spine of Ish Sah | removal | 0 | 0.98 | payoff:death-trigger tag:death-trigger-self tag:removal |
| 39 | Puppeteer Clique | threat | 1 | 0.98 | payoff:death-trigger tag:death-trigger-self type:threat |
| 40 | Persistent Constrictor | removal | 0 | 0.97 | payoff:death-trigger tag:death-trigger-self tag:removal |

Top upgrades (unowned):

- Solemn Simulacrum (ramp, 1.00)
- Syr Konrad, the Grim (threat, 1.00)
- Bastion of Remembrance (synergy, 1.00)
- The Meathook Massacre (wipe, 1.00)
- Grave Pact (removal, 0.99)
- Stitcher's Supplier (synergy, 0.99)
- Enduring Innocence (draw, 0.99)
- Liliana, Dreadhorde General (wipe, 0.99)
- Enduring Tenacity (threat, 0.99)
- Myr Retriever (synergy, 0.99)

### Review

Verdict: no. 15 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text, the EDHREC Aristocrats and Sacrifice pages, and their Orzhov filters (edhrec.com/tags/aristocrats/orzhov, /sacrifice/orzhov).

On theme (15): Syr Konrad, Bastion of Remembrance, The Meathook Massacre, Grave Pact, Stitcher's Supplier, Liliana, Dreadhorde General, Dictate of Erebos, Yahenni, Elenda, the Dusk Rose, Sephiroth, Pawn of Ulamog, Funeral Room, Greedy Freebooter, Crowded Crypt, and Nine-Lives Familiar.

Off theme (25):

- 1 Solemn Simulacrum: a ramp creature. The death trigger draws a card. Aristocrats page 20%, synergy 0.08: any deck runs it at that rate. The owner's own example of off theme.
- 7 Enduring Innocence: draws on small creatures. Its death clause returns it as a noncreature enchantment.
- 9 Enduring Tenacity: a lifegain payoff with the same return clause.
- 10 Myr Retriever and 25 Junk Diver: artifact-loop fodder. Not on any aristocrats page.
- 11 Ojer Taq and 37 Aclazotz: their death clauses return them as lands. A tokens card and a discard card.
- 12 Ichor Wellspring, 30 Prized Statue, 32 Mycosynth Wellspring, and 38 Spine of Ish Sah: noncreature artifacts. Creature sacrifice outlets cannot use them.
- 15 Wurmcoil Engine, 16 Junji, 27 Triplicate Titan, 29 Kokusho, and 36 Ancient Stone Idol: value creatures whose only link is a death trigger. Not on any aristocrats page. Kokusho is the closest call.
- 18 Nuka-Cola Vending Machine: a Food-to-Treasure loop with no creature death. Sacrifice page 7%, synergy 0.06.
- 20 Hangarback Walker: a counters card that leaves Thopters. Not on any aristocrats page. Owned.
- 22 Unstoppable Slasher: an aggro threat with a stun-counter return.
- 23 Murderous Rider: the death clause is a drawback (bottom of library).
- 24 Elenda's Hierophant: a lifegain card whose death makes tokens. Orzhov Aristocrats page 15%, synergy 0.05. This is a close call.
- 28 Aerith Gainsborough: a lifegain counters creature. Owned.
- 34 God-Eternal Oketra: a token maker whose death clause tucks it into the library.
- 39 Puppeteer Clique: an ETB value creature with persist. Not on any aristocrats page. Owned. This is a close call.
- 40 Persistent Constrictor: an upkeep drain with persist. From the Rakdos Duskmourn precon, not a sacrifice deck.

Root cause: payoff:death-trigger fires on every "when this dies" clause, including self-recursion, self-tuck, drawbacks, and "put into a graveyard from the battlefield" on noncreature artifacts. Twenty-four of the 25 misses carry it. text:sacrifice fires on Treasure and Eldrazi Spawn reminder text (Prized Statue, Greedy Freebooter, Pawn of Ulamog). Ojer Taq and Aclazotz get role land from their back faces.

Legality: all 40 are legal in commander and inside WB.

Owned: 3 of 40 are owned (Hangarback Walker, Aerith Gainsborough, Puppeteer Clique). The funnel reports 172 owned on-theme cards.

Gaps (context only): the archetype's core is absent. Blood Artist (64% of Orzhov Aristocrats decks), Cruel Celebrant (62%), Skullclamp (61%), Zulaport Cutthroat (59%), Viscera Seer (56%), Ashnod's Altar (49%), Pitiless Plunderer (48%), and Deadly Dispute (42%). Yahenni is the only free sacrifice outlet in the top 40.

## 3. tokens go-wide (commander, GW, owned-first)

Theme signals: payoff tags anthem, synergy-token, synergy-token-creature, token-doubler, token-increaser, tokenfall. tags repeatable-creature-tokens, repeatable-token-generator.

Funnel: 12612 legal in colors, 2070 on theme, 522 owned, 161 on theme and owned, 261 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Beastmaster Ascension | synergy | 0 | 0.99 | payoff:anthem payoff-text:creatures you control get theme |
| 2 | Banner of Kinship | synergy | 0 | 0.99 | payoff:anthem payoff-text:for each creature you control theme |
| 3 | Mirari's Wake | ramp | 0 | 0.99 | payoff:anthem payoff-text:creatures you control get tag:ramp |
| 4 | Elspeth, Storm Slayer | removal | 0 | 0.99 | payoff:token-doubler payoff:token-increaser tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token tag:removal |
| 5 | Peregrin Took | draw | 3 | 0.99 | payoff:synergy-token payoff:token-increaser tag:repeatable-token-generator tag:draw |
| 6 | Flowering of the White Tree | interaction | 0 | 0.99 | payoff:anthem payoff-text:creatures you control get tag:interaction |
| 7 | Elspeth, Sun's Champion | wipe | 0 | 0.99 | payoff:anthem payoff-text:creatures you control get tag:repeatable-token-generator tag:repeatable-creature-tokens text:creature token tag:sweeper |
| 8 | Elesh Norn, Grand Cenobite | wipe | 0 | 0.99 | payoff:anthem payoff-text:creatures you control get tag:sweeper |
| 9 | Forsaken Monument | ramp | 0 | 0.99 | payoff:anthem payoff-text:creatures you control get tag:ramp |
| 10 | Elvish Archdruid | ramp | 0 | 0.99 | payoff:anthem payoff-text:creatures you control get tag:ramp |
| 11 | Ocelot Pride | synergy | 0 | 0.99 | payoff:synergy-token payoff:tokenfall payoff:token-doubler payoff:token-increaser tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 12 | Tendershoot Dryad | threat | 0 | 0.99 | payoff:anthem tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 13 | Eldrazi Monument | interaction | 0 | 0.99 | payoff:anthem payoff-text:creatures you control get tag:interaction |
| 14 | Imperious Perfect | synergy | 0 | 0.99 | payoff:anthem tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 15 | Thunderfoot Baloth | threat | 0 | 0.98 | payoff:anthem payoff-text:creatures you control get type:threat |
| 16 | Righteous Valkyrie | synergy | 0 | 0.98 | payoff:anthem payoff-text:creatures you control get theme |
| 17 | Weaver of Harmony | synergy | 0 | 0.98 | payoff:anthem payoff-text:creatures you control get theme |
| 18 | Illustrious Wanderglyph | threat | 0 | 0.98 | payoff:anthem payoff-text:creatures you control get tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 19 | Ainok Strike Leader | interaction | 0 | 0.98 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token tag:interaction |
| 20 | Summon: Knights of Round | interaction | 0 | 0.98 | payoff-text:creatures you control get tag:repeatable-token-generator tag:repeatable-creature-tokens text:creature token tag:interaction |
| 21 | It That Heralds the End | synergy | 0 | 0.98 | payoff:anthem payoff-text:creatures you control get theme |
| 22 | Kozilek, the Broken Reality | draw | 0 | 0.98 | payoff:anthem payoff-text:creatures you control get tag:draw |
| 23 | The Immortal Sun | draw | 0 | 0.98 | payoff:anthem payoff-text:creatures you control get tag:draw |
| 24 | Dwynen, Gilt-Leaf Daen | threat | 0 | 0.98 | payoff:anthem payoff-text:creatures you control get type:threat |
| 25 | Fecund Greenshell | ramp | 0 | 0.97 | payoff:anthem payoff-text:creatures you control get tag:ramp |
| 26 | Sylvan Anthem | synergy | 0 | 0.97 | payoff:anthem payoff-text:creatures you control get theme |
| 27 | Chief of the Foundry | synergy | 0 | 0.97 | payoff:anthem payoff-text:creatures you control get theme |
| 28 | Esika's Chariot | synergy | 0 | 0.97 | payoff:synergy-token tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 29 | Chitterspitter | synergy | 0 | 0.97 | payoff:synergy-token payoff:anthem tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 30 | Nesting Dovehawk | threat | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 31 | Blossoming Tortoise | ramp | 0 | 0.97 | payoff:anthem payoff-text:creatures you control get tag:ramp |
| 32 | Trostani, Selesnya's Voice | threat | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 33 | Transmutation Font | draw | 0 | 0.97 | payoff:synergy-token tag:repeatable-token-generator tag:draw |
| 34 | Rhys the Redeemed | synergy | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature payoff:token-doubler payoff:token-increaser tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 35 | Maja, Bretagard Protector | threat | 0 | 0.97 | payoff:anthem payoff-text:creatures you control get tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 36 | Leyline of Hope | synergy | 0 | 0.97 | payoff:anthem payoff-text:creatures you control get theme |
| 37 | Knight Exemplar | interaction | 0 | 0.97 | payoff:anthem payoff-text:creatures you control get tag:interaction |
| 38 | Oltec Matterweaver | synergy | 0 | 0.97 | payoff:synergy-token tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 39 | On Wings of Gold | synergy | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature payoff:anthem tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 40 | Nissa, Ascended Animist | removal | 0 | 0.97 | payoff-text:creatures you control get tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token tag:removal |

Top upgrades (unowned):

- Beastmaster Ascension (synergy, 0.99)
- Banner of Kinship (synergy, 0.99)
- Mirari's Wake (ramp, 0.99)
- Elspeth, Storm Slayer (removal, 0.99)
- Flowering of the White Tree (interaction, 0.99)
- Elspeth, Sun's Champion (wipe, 0.99)
- Elesh Norn, Grand Cenobite (wipe, 0.99)
- Forsaken Monument (ramp, 0.99)
- Elvish Archdruid (ramp, 0.99)
- Ocelot Pride (synergy, 0.99)

### Review

Verdict: no. 25 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Tokens page (edhrec.com/tags/tokens).

Off theme (15):

- 2 Banner of Kinship: pumps one chosen creature type only. A typal lord, not a team anthem.
- 9 Forsaken Monument: pumps colorless creatures only. The other lines are colorless ramp and lifegain.
- 10 Elvish Archdruid: Elf lord and Elf mana. EDHREC lists it on the Elves page (86% of decks), not on the Tokens page.
- 16 Righteous Valkyrie: the team anthem needs 7 life above the start total. A lifegain card. This is a close call.
- 17 Weaver of Harmony: pumps enchantment creatures only. An enchantress card (31% of enchantress decks).
- 21 It That Heralds the End: pumps colorless creatures only. An Eldrazi card.
- 22 Kozilek, the Broken Reality: pumps colorless creatures only. A nine-mana Eldrazi.
- 24 Dwynen, Gilt-Leaf Daen: Elf lord. On the Elves page (39%), not on the Tokens page.
- 25 Fecund Greenshell: the anthem needs ten lands. A land ramp card.
- 26 Sylvan Anthem: pumps green creatures only. White tokens get nothing. This is a close call.
- 27 Chief of the Foundry: pumps artifact creatures only.
- 31 Blossoming Tortoise: "Land creatures you control get +1/+1". A lands card.
- 33 Transmutation Font: makes Blood, Clue, or Food tokens only. It does not widen the board.
- 36 Leyline of Hope: the team anthem needs 7 life above the start total. A lifegain card.
- 37 Knight Exemplar: pumps Knights only.

Root cause: the anthem matcher fires on "creatures you control get" and ignores the qualifier before it (creature type, color, or condition). Thirteen of the 15 misses come from this. Banner of Kinship drops "of the chosen type". Transmutation Font gets repeatable-token-generator for noncreature tokens.

Legality: all 40 are legal in commander and inside GW.

Owned: only Peregrin Took (rank 5) is owned. The other 39 are upgrades, in owned-first mode.

Gaps (context only): Anointed Procession, Parallel Lives, Doubling Season, Second Harvest, Intangible Virtue, Cathars' Crusade, Craterhoof Behemoth, and Idol of Oblivion are absent. Each is on the EDHREC Tokens page (12% to 28% of decks).

## 4. voltron equipment (commander, RW, owned-first)

Theme signals: payoff tags synergy-aura, synergy-equipment, synergy-modified. tags evasion, protection. keywords Equip, Double strike, Hexproof.

Funnel: 12737 legal in colors, 3362 on theme, 604 owned, 234 on theme and owned, 285 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Swiftfoot Boots | interaction | 1 | 1.00 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 2 | Lightning Greaves | interaction | 1 | 1.00 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 3 | Mithril Coat | interaction | 0 | 1.00 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 4 | Whispersilk Cloak | interaction | 0 | 1.00 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 5 | Commander's Plate | interaction | 0 | 1.00 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 6 | Sword of Feast and Famine | ramp | 0 | 0.99 | payoff-text:equipped creature tag:protection keyword:Equip tag:ramp |
| 7 | Brotherhood Regalia | interaction | 0 | 0.99 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 8 | Puresteel Paladin | draw | 1 | 0.99 | payoff:synergy-equipment payoff-text:whenever an equipment text:equipment tag:draw |
| 9 | Hammer of Nazahn | interaction | 0 | 0.99 | payoff:synergy-equipment payoff-text:equipped creature tag:protection keyword:Equip text:equipment tag:interaction |
| 10 | Champion's Helm | interaction | 1 | 0.99 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 11 | Sigarda's Aid | synergy | 0 | 0.99 | payoff:synergy-equipment payoff:synergy-aura payoff-text:whenever an equipment text:equipment theme |
| 12 | Sword of Hearth and Home | ramp | 0 | 0.99 | payoff-text:equipped creature tag:protection keyword:Equip tag:ramp |
| 13 | Darksteel Plate | interaction | 1 | 0.99 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 14 | Sword of Fire and Ice | interaction | 0 | 0.99 | payoff-text:equipped creature tag:protection keyword:Equip text:equipment tag:interaction |
| 15 | Mantle of the Ancients | synergy | 0 | 0.99 | payoff:synergy-equipment payoff:synergy-aura payoff-text:enchanted creature text:equipment theme |
| 16 | Caduceus, Staff of Hermes | interaction | 0 | 0.99 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 17 | Lavaspur Boots | interaction | 0 | 0.99 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 18 | Codsworth, Handy Helper | ramp | 0 | 0.99 | payoff:synergy-equipment payoff:synergy-aura tag:protection text:equipment tag:ramp |
| 19 | Thran Power Suit | interaction | 0 | 0.99 | payoff:synergy-equipment payoff:synergy-aura payoff-text:equipped creature tag:protection keyword:Equip text:equipment tag:interaction |
| 20 | Kaldra Compleat | interaction | 0 | 0.98 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 21 | Sword of Truth and Justice | interaction | 0 | 0.98 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 22 | Silver Shroud Costume | interaction | 0 | 0.98 | payoff-text:equipped creature tag:protection keyword:Equip text:equipment tag:interaction |
| 23 | Sword of Forge and Frontier | ramp | 0 | 0.98 | payoff-text:equipped creature tag:protection keyword:Equip tag:ramp |
| 24 | Akiri, Fearless Voyager | interaction | 0 | 0.98 | payoff:synergy-equipment payoff-text:equipped creature tag:protection text:equipment tag:interaction |
| 25 | Sword of Wealth and Power | ramp | 0 | 0.98 | payoff-text:equipped creature tag:protection keyword:Equip tag:ramp |
| 26 | Dragonfire Blade | interaction | 0 | 0.98 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 27 | Sword of Light and Shadow | interaction | 0 | 0.98 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 28 | Greater Auramancy | interaction | 0 | 0.98 | payoff:synergy-aura payoff-text:enchanted creature tag:protection tag:interaction |
| 29 | Halvar, God of Battle // Sword of the Realms | threat | 0 | 0.98 | payoff:synergy-equipment payoff:synergy-aura payoff-text:equipped creature keyword:Equip text:equipment type:threat |
| 30 | Armored Skyhunter | threat | 0 | 0.98 | payoff:synergy-equipment payoff:synergy-aura tag:evasion text:equipment type:threat |
| 31 | Sword of War and Peace | interaction | 0 | 0.97 | payoff-text:equipped creature tag:protection keyword:Equip text:equipment tag:interaction |
| 32 | Zack Fair | interaction | 2 | 0.97 | payoff:synergy-equipment tag:protection text:equipment tag:interaction |
| 33 | Cid, Freeflier Pilot | synergy | 1 | 0.97 | payoff:synergy-equipment tag:evasion text:equipment theme |
| 34 | Transcendent Envoy | synergy | 0 | 0.97 | payoff:synergy-aura tag:evasion theme |
| 35 | Robe of Stars | interaction | 0 | 0.97 | payoff-text:equipped creature tag:protection keyword:Equip tag:interaction |
| 36 | Goro-Goro, Disciple of Ryusei | synergy | 0 | 0.97 | payoff:synergy-equipment payoff:synergy-aura payoff:synergy-modified tag:evasion text:equipment theme |
| 37 | Celestial Armor | interaction | 0 | 0.97 | payoff-text:equipped creature tag:protection keyword:Equip text:equipment tag:interaction |
| 38 | Daybreak Coronet | synergy | 0 | 0.97 | payoff:synergy-aura payoff-text:enchanted creature theme |
| 39 | Songbirds' Blessing | synergy | 0 | 0.97 | payoff:synergy-aura payoff-text:enchanted creature theme |
| 40 | Eidolon of Countless Battles | synergy | 0 | 0.97 | payoff:synergy-aura payoff-text:enchanted creature theme |

Top upgrades (unowned):

- Mithril Coat (interaction, 1.00)
- Whispersilk Cloak (interaction, 1.00)
- Commander's Plate (interaction, 1.00)
- Sword of Feast and Famine (ramp, 0.99)
- Brotherhood Regalia (interaction, 0.99)
- Hammer of Nazahn (interaction, 0.99)
- Sigarda's Aid (synergy, 0.99)
- Sword of Hearth and Home (ramp, 0.99)
- Sword of Fire and Ice (interaction, 0.99)
- Mantle of the Ancients (synergy, 0.99)

### Review

Verdict: yes. 38 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Voltron, Equipment, and Auras pages (edhrec.com/tags/voltron, /equipment, /auras).

Off theme (2):

- 28 Greater Auramancy: gives your other enchantments and your enchanted creatures shroud. It has no Equipment text, and shroud blocks your own equip abilities. Not on the Voltron, Equipment, or Auras pages. An enchantress card (17% of enchantress decks). This is a close call.
- 36 Goro-Goro, Disciple of Ryusei: a haste anthem and a Dragon token engine that needs an attacking modified creature. The Equipment and Aura signals match only the reminder text. Not on any EDHREC page in the set. This is a close call.

Signal bugs: Goro-Goro gets payoff:synergy-equipment, payoff:synergy-aura, text:equipment, and tag:evasion from reminder text. The card has no Equipment, Aura, or evasion ability. Greater Auramancy gets payoff-text:enchanted creature from a static grant, not from an Aura.

Legality: all 40 are legal in commander and inside RW.

Owned: 7 of 40 are owned (ranks 1, 2, 8, 10, 13, 32, 33). The funnel reports 234 owned on-theme cards.

Gaps (context only): the archetype's tutor and draw core is absent. Sram, Senior Edificer (61% of Equipment decks), Open the Armory (48%), Steelshaper's Gift (46%), Forge Anew (46%), Stoneforge Mystic (42%), Danitha Capashen, Paragon (42%), Blackblade Reforged (40%), and Colossus Hammer (40%). The ranker favors "equipped creature" protection text over these engine cards.

## 5. blink (commander, WU, owned-first)

Theme signals: tags flicker, flicker-creature, flicker-self, flicker-slow.

Funnel: 12669 legal in colors, 207 on theme, 450 owned, 14 on theme and owned, 169 returned, 50 upgrades.

Thin theme: the collection holds under 30 on-theme cards. PR-7 asks the pool-mode question again here (D-63).

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Ephemerate | interaction | 0 | 0.80 | tag:flicker tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield tag:interaction |
| 2 | Y'shtola Rhul | removal | 0 | 0.78 | tag:flicker tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield text:destroy target |
| 3 | Momentary Blink | interaction | 0 | 0.78 | tag:flicker tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield tag:interaction |
| 4 | Splash Portal | draw | 0 | 0.78 | tag:flicker tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield tag:draw |
| 5 | Conjurer's Closet | removal | 0 | 0.69 | tag:flicker tag:flicker-creature text:exile target creature you control, then return text:destroy target |
| 6 | Cloudshift | interaction | 0 | 0.68 | tag:flicker tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 7 | Essence Flux | interaction | 0 | 0.68 | tag:flicker tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 8 | Nezahal, Primal Tide | draw | 0 | 0.68 | tag:flicker tag:flicker-slow tag:flicker-self text:return it to the battlefield tag:draw |
| 9 | Deadeye Navigator | interaction | 0 | 0.68 | tag:flicker tag:flicker-creature tag:flicker-self text:return it to the battlefield tag:interaction |
| 10 | Charming Prince | synergy | 0 | 0.68 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield theme |
| 11 | Touch the Spirit Realm | interaction | 0 | 0.67 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield tag:interaction |
| 12 | Flicker of Fate | interaction | 0 | 0.67 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:interaction |
| 13 | Planar Incision | interaction | 0 | 0.67 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:interaction |
| 14 | Blur | interaction | 0 | 0.67 | tag:flicker tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 15 | Airbender Ascension | removal | 0 | 0.67 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:removal |
| 16 | Estrid's Invocation | synergy | 0 | 0.67 | tag:flicker tag:flicker-self text:return it to the battlefield theme |
| 17 | Eldrazi Confluence | ramp | 0 | 0.67 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:ramp |
| 18 | Eldrazi Displacer | interaction | 0 | 0.66 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:interaction |
| 19 | Oath of Teferi | synergy | 0 | 0.66 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield theme |
| 20 | Cosmic Intervention | interaction | 0 | 0.65 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield tag:interaction |
| 21 | Settle Beyond Reality | removal | 0 | 0.65 | tag:flicker tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield tag:removal |
| 22 | Gossip's Talent | synergy | 0 | 0.65 | tag:flicker tag:flicker-creature text:return it to the battlefield theme |
| 23 | Distinguished Conjurer | interaction | 0 | 0.65 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:interaction |
| 24 | Far Traveler | synergy | 0 | 0.65 | tag:flicker tag:flicker-creature text:return it to the battlefield theme |
| 25 | Slip On the Ring | interaction | 1 | 0.65 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:interaction |
| 26 | Scrollshift | interaction | 0 | 0.64 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:interaction |
| 27 | Acrobatic Maneuver | interaction | 0 | 0.64 | tag:flicker tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 28 | Abuelo, Ancestral Echo | interaction | 0 | 0.64 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield tag:interaction |
| 29 | Siren's Ruse | interaction | 0 | 0.64 | tag:flicker tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 30 | Personify | removal | 3 | 0.64 | tag:flicker tag:flicker-creature text:exile target creature you control, then return text:destroy target |
| 31 | Long River Lurker | interaction | 0 | 0.63 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:interaction |
| 32 | Venser, the Sojourner | removal | 0 | 0.63 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield tag:removal |
| 33 | Getaway Glamer | interaction | 0 | 0.63 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield tag:interaction |
| 34 | Mystifying Maze | land | 0 | 0.63 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield type:land |
| 35 | Justiciar's Portal | interaction | 0 | 0.63 | tag:flicker tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 36 | Oji, the Exquisite Blade | interaction | 0 | 0.62 | tag:flicker tag:flicker-creature text:return it to the battlefield tag:interaction |
| 37 | Stenn, Paranoid Partisan | synergy | 0 | 0.62 | tag:flicker tag:flicker-slow tag:flicker-self text:return it to the battlefield theme |
| 38 | Meneldor, Swift Savior | threat | 1 | 0.62 | tag:flicker tag:flicker-creature text:return it to the battlefield type:threat |
| 39 | Intruder Alarm | synergy | 0 | 0.62 | payoff-text:whenever a creature enters theme |
| 40 | Guardian of Ghirapur | synergy | 0 | 0.62 | tag:flicker tag:flicker-creature tag:flicker-slow text:return it to the battlefield theme |

Top upgrades (unowned):

- Ephemerate (interaction, 0.80)
- Y'shtola Rhul (removal, 0.78)
- Momentary Blink (interaction, 0.78)
- Splash Portal (draw, 0.78)
- Conjurer's Closet (removal, 0.69)
- Cloudshift (interaction, 0.68)
- Essence Flux (interaction, 0.68)
- Nezahal, Primal Tide (draw, 0.68)
- Deadeye Navigator (interaction, 0.68)
- Charming Prince (synergy, 0.68)

### Review

Verdict: no. 33 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and legality, the EDHREC Blink page (edhrec.com/tags/blink), and a score probe of `internal/candidates` on the 2026-08-24 snapshot.

Off theme (7):

- 8 Nezahal, Primal Tide: a seven-mana draw engine. The self-exile costs three cards and only protects Nezahal.
- 16 Estrid's Invocation: an enchantment clone. The self-blink returns a copy of another enchantment.
- 17 Eldrazi Confluence: one mode of three flickers, and the permanent returns tapped. The card also needs {C}{C}, which a WU deck rarely makes.
- 31 Long River Lurker: a Frog lord with ward. The blink needs combat damage from one chosen creature.
- 34 Mystifying Maze: the target is an attacking creature an opponent controls. The land never repeats your own ETB triggers.
- 37 Stenn, Paranoid Partisan: a cost reducer. The self-blink is protection only.
- 39 Intruder Alarm: a combo untapper. Its only signal is payoff-text:whenever a creature enters.

Borderline (counted on theme): 20 Cosmic Intervention, 21 Settle Beyond Reality, 22 Gossip's Talent, and 38 Meneldor, Swift Savior. Each one holds a real blink line under a condition. Count these four as off theme, and the total falls to 29.

Root cause: the blink row in `themes.json` names two payoff slugs, `etb-matters` and `synergy-flicker`. Neither slug exists in Scryfall Tagger, so the loader drops both (`otag:etb-matters` returns no cards, 2026-08-24). Only the four flicker enabler tags remain. Therefore 39 of 40 cards are flicker effects, and one card fires a payoff signal. D-62 asks for payoffs before enablers, and this list gives the opposite. A probe of every slug in `themes.json` against the 2026-08-24 tag index found 16 slugs that do not exist, over 11 theme rows.

Second root cause: the two payoff needles almost never match. The needle "whenever a creature enters" matches 9 WU cards, but the common wording "whenever a (or another) creature you control enters" matches 45. The needle "enters the battlefield under your control" matches 1 WU card, because the 2024 templating writes "enters".

Cutoff: the two text needles are worth 0.4 each, and they select one-shot spells with a single target. Soulherder (58% of Blink decks) scores 0.56 at position 62. Guardian of Ghirapur (EDHREC rank 8127) takes rank 40 with 0.62.

Signal bugs: Y'shtola Rhul, Conjurer's Closet, and Personify get role removal from the text fallback for "exile target". The signal prints as text:destroy target, but none of the three destroys a permanent. Eldrazi Confluence gets role ramp from its Eldrazi Scion token.

Legality: all 40 are legal in commander and inside WU. Conjurer's Closet, Eldrazi Confluence, and Mystifying Maze are colorless.

Owned: 3 of 40 are owned (Slip On the Ring, Personify, Meneldor, Swift Savior). The funnel reports 14 owned on-theme cards, so the thin-theme flag is correct (D-63).

Gaps (context only): every ETB payoff is absent. Panharmonicon (40% of Blink decks) never enters the candidate pool, because it has no flicker tag and no staple role. Soulherder (58%), Ghostly Flicker (50%), Teleportation Circle (44%), Eerie Interlude (41%), Displacer Kitten (39%), Restoration Angel (35%), Thassa, Deep-Dwelling (31%), Felidar Guardian (30%), and Brago, King Eternal sit between position 54 and position 68. Mulldrifter (29%) and Aether Channeler (29%) fall to position 382 and position 633, because each keeps only a staple role.

## 6. reanimator (commander, B, owned-first)

Theme signals: payoff tags leaving-graveyard-matters, reanimate-matters. tags discard, mill-self, reanimate, reanimate-cast, reanimate-creature.

Funnel: 7465 legal in colors, 1042 on theme, 301 owned, 59 on theme and owned, 193 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Syr Konrad, the Grim | threat | 0 | 1.00 | payoff:leaving-graveyard-matters payoff:reanimate-matters payoff-text:leaves your graveyard tag:mill-self type:threat |
| 2 | Twilight Diviner | synergy | 0 | 0.97 | payoff:reanimate-matters tag:mill-self theme |
| 3 | Skeleton Crew | threat | 0 | 0.95 | payoff:leaving-graveyard-matters tag:reanimate text:from your graveyard to the battlefield type:threat |
| 4 | Canoptek Tomb Sentinel | removal | 0 | 0.93 | payoff:reanimate-matters tag:reanimate text:from your graveyard to the battlefield tag:removal |
| 5 | Triarch Praetorian | draw | 0 | 0.91 | payoff:reanimate-matters tag:reanimate text:from your graveyard to the battlefield tag:draw |
| 6 | Along the Crooked Way | synergy | 0 | 0.83 | payoff:leaving-graveyard-matters payoff-text:leaves your graveyard theme |
| 7 | Grim Reaper's Scythe | synergy | 0 | 0.81 | payoff:leaving-graveyard-matters tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 8 | Dredging Claw | synergy | 0 | 0.78 | payoff:reanimate-matters payoff-text:whenever a creature enters from your graveyard theme |
| 9 | Nether Traitor | synergy | 0 | 0.78 | tag:reanimate text:from your graveyard to the battlefield text:put into your graveyard theme |
| 10 | Polluted Cistern // Dim Oubliette | synergy | 0 | 0.75 | tag:reanimate tag:reanimate-creature tag:mill-self text:from your graveyard to the battlefield text:put into your graveyard theme |
| 11 | Teval's Judgment | ramp | 0 | 0.70 | payoff:leaving-graveyard-matters tag:ramp |
| 12 | Dread Return | synergy | 0 | 0.69 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 13 | Tormod, the Desecrator | threat | 0 | 0.69 | payoff:leaving-graveyard-matters payoff:reanimate-matters type:threat |
| 14 | Whip of Erebos | synergy | 0 | 0.69 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 15 | Sheoldred, Whispering One | removal | 0 | 0.69 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield tag:removal |
| 16 | Reassembling Skeleton | synergy | 0 | 0.68 | tag:reanimate text:from your graveyard to the battlefield theme |
| 17 | Unearth | draw | 0 | 0.68 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield tag:draw |
| 18 | Agadeem's Awakening // Agadeem, the Undercrypt | land | 0 | 0.68 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield type:land |
| 19 | The Soul Stone | ramp | 0 | 0.68 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield tag:ramp |
| 20 | Bloodghast | synergy | 0 | 0.68 | tag:reanimate text:from your graveyard to the battlefield theme |
| 21 | Artisan of Kozilek | removal | 0 | 0.68 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield tag:removal |
| 22 | Desecrated Tomb | synergy | 0 | 0.68 | payoff:leaving-graveyard-matters payoff:reanimate-matters theme |
| 23 | Persist | synergy | 1 | 0.68 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 24 | Stitch Together | synergy | 0 | 0.68 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 25 | Lively Dirge | synergy | 0 | 0.68 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 26 | Imotekh the Stormlord | threat | 0 | 0.67 | payoff:leaving-graveyard-matters type:threat |
| 27 | Drownyard Temple | land | 0 | 0.67 | tag:reanimate text:from your graveyard to the battlefield type:land |
| 28 | Liliana, Death's Majesty | wipe | 0 | 0.67 | tag:reanimate tag:reanimate-creature tag:mill-self text:from your graveyard to the battlefield tag:sweeper |
| 29 | Funeral Room // Awakening Hall | synergy | 0 | 0.67 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 30 | Ruthless Technomancer | ramp | 0 | 0.67 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield tag:ramp |
| 31 | Chthonian Nightmare | synergy | 0 | 0.67 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 32 | Forsaken Miner | synergy | 0 | 0.67 | tag:reanimate text:from your graveyard to the battlefield theme |
| 33 | Lich-Knights' Conquest | synergy | 0 | 0.67 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 34 | Yawgmoth's Will | synergy | 0 | 0.67 | tag:reanimate tag:reanimate-cast text:put into your graveyard theme |
| 35 | Vat of Rebirth | synergy | 0 | 0.67 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 36 | Will of the Abzan | removal | 0 | 0.67 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield tag:removal |
| 37 | Defiled Crypt // Cadaver Lab | synergy | 0 | 0.67 | payoff:leaving-graveyard-matters theme |
| 38 | Bloodline Bidding | synergy | 1 | 0.67 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield theme |
| 39 | Metamorphosis Fanatic | threat | 0 | 0.67 | tag:reanimate tag:reanimate-creature text:from your graveyard to the battlefield type:threat |
| 40 | Cauldron Familiar | synergy | 0 | 0.66 | tag:reanimate text:from your graveyard to the battlefield theme |

Top upgrades (unowned):

- Syr Konrad, the Grim (threat, 1.00)
- Twilight Diviner (synergy, 0.97)
- Skeleton Crew (threat, 0.95)
- Canoptek Tomb Sentinel (removal, 0.93)
- Triarch Praetorian (draw, 0.91)
- Along the Crooked Way (synergy, 0.83)
- Grim Reaper's Scythe (synergy, 0.81)
- Dredging Claw (synergy, 0.78)
- Nether Traitor (synergy, 0.78)
- Polluted Cistern // Dim Oubliette (synergy, 0.75)

### Review

Verdict: no. 34 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text, the EDHREC Reanimator and Self-Mill pages, and the mono-black Reanimator page (edhrec.com/tags/reanimator/mono-black).

Off theme (6):

- 8 Dredging Claw: a {2} Equipment for +1/+0 and menace. The theme link is a free attach when a creature enters from your graveyard. Not on any reanimator page. EDHREC card page: 1,106 decks.
- 19 The Soul Stone: a mana rock. Its upkeep reanimation needs {6}{B}, a tap, and an exiled creature first. Mono-black Reanimator page: 19% of decks, synergy 0.03, so generic mono-black decks run it at the same rate. This is a close call. The verdict stays no if it counts.
- 21 Artisan of Kozilek: reanimates only as a cast trigger on a nine-mana spell. When a reanimator deck returns it, the trigger does not fire. Not on any reanimator page. Its top commanders are Eldrazi ramp.
- 27 Drownyard Temple: a colorless land that returns itself for {3}. Land recursion, not creature reanimation. Self-Mill page 8%, synergy 0.07. Not on the Reanimator pages.
- 34 Yawgmoth's Will: casts from the graveyard rather than reanimates. A storm and combo card. Mono-black Reanimator page: 6%, synergy 0.02.
- 40 Cauldron Familiar: returns only when you sacrifice a Food. A Food and aristocrats card. Not on any reanimator page.

Borderline, counted on theme (1):

- 26 Imotekh the Stormlord: a leaves-graveyard payoff for artifact cards only. It is the top commander on the mono-black Reanimator page (628 decks), so the archetype builds around it.

Signal bugs: Yawgmoth's Will gets text:put into your graveyard from its exile-replacement clause. Nether Traitor gets the same signal from a death-trigger condition. Agadeem's Awakening gets role land from its back face. Teval's Judgment gets role ramp and Liliana, Death's Majesty gets role wipe, while the theme uses their other abilities.

Legality: all 40 are legal in commander and inside mono-black (four are colorless).

Owned: 2 of 40 are owned (Persist, Bloodline Bidding). The funnel reports 59 owned on-theme cards.

Gaps (context only): the archetype's defining spells are absent. Reanimate (59% of mono-black Reanimator decks), Victimize (46%), Buried Alive (36%), Animate Dead (35%), Entomb (30%), Living Death (25%), Rise of the Dark Realms (19%), and Stitcher's Supplier (18%). Likely cause: the text signal "from your graveyard to the battlefield" does not match "from a graveyard onto the battlefield", and tutor-to-graveyard cards have no tag.

## 7. landfall (commander, GU, owned-first)

Theme signals: payoff tags land-count-matters, landfall, landfall-other, lands-matter. tags land-ramp. keywords Landfall.

Funnel: 12456 legal in colors, 867 on theme, 465 owned, 106 on theme and owned, 176 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Dreamroot Cascade | land | 0 | 1.00 | payoff:lands-matter payoff-text:land enters type:land |
| 2 | Tireless Provisioner | ramp | 0 | 1.00 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall tag:ramp |
| 3 | Scute Swarm | synergy | 0 | 1.00 | payoff:landfall payoff:lands-matter payoff:land-count-matters payoff-text:whenever a land you control enters keyword:Landfall theme |
| 4 | Avenger of Zendikar | threat | 0 | 1.00 | payoff:landfall payoff:lands-matter payoff:land-count-matters payoff-text:whenever a land you control enters keyword:Landfall type:threat |
| 5 | Lotus Cobra | ramp | 0 | 1.00 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall tag:ramp |
| 6 | Rampaging Baloths | threat | 0 | 1.00 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall type:threat |
| 7 | Evolution Sage | synergy | 1 | 1.00 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 8 | Tatyova, Benthic Druid | draw | 0 | 1.00 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall tag:draw |
| 9 | Field of the Dead | land | 0 | 1.00 | payoff:landfall payoff:lands-matter payoff-text:land enters type:land |
| 10 | Tireless Tracker | draw | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall tag:draw |
| 11 | Icetill Explorer | ramp | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters tag:land-ramp keyword:Landfall tag:ramp |
| 12 | Springheart Nantuko | synergy | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 13 | Mossborn Hydra | synergy | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 14 | Aesi, Tyrant of Gyre Strait | ramp | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters tag:land-ramp keyword:Landfall tag:ramp |
| 15 | Burgeoning | ramp | 0 | 0.99 | payoff:landfall payoff:landfall-other payoff:lands-matter tag:land-ramp tag:ramp |
| 16 | Loot, Exuberant Explorer | ramp | 0 | 0.99 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 17 | Bristly Bill, Spine Sower | synergy | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 18 | Lumra, Bellow of the Woods | ramp | 0 | 0.99 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 19 | Topiary Stomper | ramp | 0 | 0.99 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 20 | Courser of Kruphix | synergy | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 21 | Ruin Crab | synergy | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 22 | Hedron Crab | synergy | 0 | 0.99 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 23 | Case of the Locked Hothouse | ramp | 0 | 0.99 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 24 | Sodden Verdure | land | 2 | 0.99 | payoff:lands-matter payoff-text:land enters type:land |
| 25 | Cultivator Colossus | ramp | 0 | 0.99 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 26 | Druid Class | ramp | 0 | 0.98 | payoff:landfall payoff:lands-matter payoff:land-count-matters payoff-text:whenever a land you control enters tag:land-ramp keyword:Landfall tag:ramp |
| 27 | Retreat to Coralhelm | synergy | 0 | 0.98 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 28 | Summon: Titan | ramp | 1 | 0.98 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 29 | Anticausal Vestige | ramp | 0 | 0.98 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 30 | Zendikar's Roil | synergy | 0 | 0.98 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 31 | Nissa, Resurgent Animist | ramp | 0 | 0.98 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall tag:ramp |
| 32 | Primeval Bounty | synergy | 0 | 0.98 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 33 | Earthbender Ascension | ramp | 1 | 0.98 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters tag:land-ramp keyword:Landfall tag:ramp |
| 34 | Beanstalk Giant // Fertile Footsteps | ramp | 0 | 0.98 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 35 | Wrenn and Seven | ramp | 0 | 0.98 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |
| 36 | Retreat to Kazandu | synergy | 0 | 0.98 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 37 | Khalni Heart Expedition | ramp | 0 | 0.98 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters tag:land-ramp keyword:Landfall tag:ramp |
| 38 | Greensleeves, Maro-Sorcerer | threat | 0 | 0.97 | payoff:landfall payoff:lands-matter payoff:land-count-matters payoff-text:whenever a land you control enters keyword:Landfall type:threat |
| 39 | Tifa Lockhart | synergy | 0 | 0.97 | payoff:landfall payoff:lands-matter payoff-text:whenever a land you control enters keyword:Landfall theme |
| 40 | Ulvenwald Hydra | ramp | 0 | 0.97 | payoff:lands-matter payoff:land-count-matters tag:land-ramp tag:ramp |

Top upgrades (unowned):

- Dreamroot Cascade (land, 1.00)
- Tireless Provisioner (ramp, 1.00)
- Scute Swarm (synergy, 1.00)
- Avenger of Zendikar (threat, 1.00)
- Lotus Cobra (ramp, 1.00)
- Rampaging Baloths (threat, 1.00)
- Tatyova, Benthic Druid (draw, 1.00)
- Field of the Dead (land, 1.00)
- Tireless Tracker (draw, 0.99)
- Icetill Explorer (ramp, 0.99)

### Review

Verdict: yes. 36 of 40 are on theme. The bar is 36, so this prompt sits exactly at the bar. Review date: 2026-08-24. Sources: Scryfall oracle text, the EDHREC Landfall and Lands Matter pages (edhrec.com/tags/landfall, /lands-matter), and the Tatyova, Aesi, and Bristly Bill commander pages.

Off theme (4):

- 1 Dreamroot Cascade: a plain Simic slow land. "This land enters tapped unless you control two or more other lands." No landfall trigger. It scored 1.00, above every real landfall card.
- 24 Sodden Verdure: a plain Forest Island dual with the same enters-tapped clause. No landfall trigger. Owned (2).
- 29 Anticausal Vestige: a colorless Eldrazi. Its only lands link is a cap on a cheat effect when it leaves the battlefield. No landfall trigger, land drop, or ramp. Not on the Landfall or Lands Matter pages.
- 32 Primeval Bounty: two cast triggers and a landfall clause that gains 3 life. Landfall page: 8% of decks, synergy 0.06. This is a close call.

Signal bugs: payoff-text:land enters fires on the enters-tapped clause of the two duals. The pattern should require "whenever a land" or the Landfall keyword, and exclude "enters tapped". Anticausal Vestige gets tag:land-ramp and payoff:land-count-matters from a cap, not a payoff. Topiary Stomper gets payoff:land-count-matters from an attack restriction.

Legality: all 40 are legal in commander and inside GU.

Owned: 4 of 40 are owned (Evolution Sage, Sodden Verdure, Summon: Titan, Earthbender Ascension). The funnel reports 106 owned on-theme cards.

Gaps (context only): the extra-land-drop and double-landfall staples are absent. Evolving Wilds (54% of Landfall decks), Azusa, Lost but Seeking (51%), Ancient Greenwarden (47%), Harrow (45%), Ramunap Excavator (42%), Oracle of Mul Daya (38%), Dryad of the Ilysian Grove (37%), and Exploration (36%).

## 8. spellslinger (commander, UR, owned-first)

Theme signals: payoff tags second-spell-matters, storm-count-matters. tags cantrip, copy. keywords Prowess, Magecraft.

Funnel: 12588 legal in colors, 1313 on theme, 518 owned, 111 on theme and owned, 216 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Thousand-Year Storm | synergy | 0 | 0.99 | payoff:storm-count-matters payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 2 | Primal Amulet // Primal Wellspring | land | 0 | 0.98 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery type:land |
| 3 | Kitsa, Otterball Elite | draw | 0 | 0.98 | payoff-text:whenever you cast a noncreature spell tag:copy keyword:Prowess text:instant or sorcery tag:draw |
| 4 | Ral, Crackling Wit | draw | 0 | 0.98 | payoff-text:whenever you cast an instant or sorcery payoff-text:whenever you cast a noncreature spell tag:copy text:instant or sorcery tag:draw |
| 5 | Rionya, Fire Dancer | threat | 0 | 0.97 | payoff:storm-count-matters tag:copy type:threat |
| 6 | Zada, Hedron Grinder | threat | 2 | 0.97 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery type:threat |
| 7 | Krark, the Thumbless | synergy | 0 | 0.97 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 8 | Storm of Saruman | synergy | 0 | 0.97 | payoff:second-spell-matters tag:copy theme |
| 9 | Ral, Monsoon Mage // Ral, Leyline Prodigy | removal | 0 | 0.97 | payoff:storm-count-matters payoff-text:whenever you cast an instant or sorcery text:instant or sorcery tag:removal |
| 10 | Sorcerer Class | ramp | 0 | 0.97 | payoff:storm-count-matters payoff-text:whenever you cast an instant or sorcery text:instant or sorcery tag:ramp |
| 11 | Stormsplitter | threat | 0 | 0.96 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery type:threat |
| 12 | Chandra, Hope's Beacon | ramp | 0 | 0.96 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery tag:ramp |
| 13 | Cori-Steel Cutter | synergy | 0 | 0.96 | payoff:second-spell-matters payoff-text:whenever you cast a noncreature spell theme |
| 14 | Swarm Intelligence | synergy | 0 | 0.96 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 15 | Prismari, the Inspiration | threat | 0 | 0.96 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery type:threat |
| 16 | Stella Lee, Wild Card | synergy | 0 | 0.95 | payoff:second-spell-matters tag:copy text:instant or sorcery theme |
| 17 | Taigam, Master Opportunist | synergy | 0 | 0.95 | payoff:second-spell-matters tag:copy theme |
| 18 | Leyline of Resonance | synergy | 0 | 0.95 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 19 | Orvar, the All-Form | ramp | 0 | 0.95 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery tag:ramp |
| 20 | Mercurial Spelldancer | synergy | 0 | 0.95 | payoff-text:whenever you cast a noncreature spell tag:copy text:instant or sorcery theme |
| 21 | Tomb of Horrors Adventurer | threat | 0 | 0.94 | payoff:second-spell-matters tag:copy type:threat |
| 22 | Melek, Izzet Paragon | threat | 0 | 0.93 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery type:threat |
| 23 | Gandalf the Grey | threat | 1 | 0.93 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery type:threat |
| 24 | Slick Sequence | removal | 0 | 0.92 | payoff:second-spell-matters tag:cantrip tag:removal |
| 25 | Aria of Flame | removal | 0 | 0.92 | payoff:storm-count-matters payoff-text:whenever you cast an instant or sorcery text:instant or sorcery tag:removal |
| 26 | Breeches, the Blastmaker | removal | 0 | 0.92 | payoff:second-spell-matters tag:copy tag:removal |
| 27 | Adaptive Training Post | synergy | 0 | 0.92 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 28 | Coruscation Mage | synergy | 0 | 0.91 | payoff-text:whenever you cast a noncreature spell tag:copy theme |
| 29 | Spinerock Tyrant | threat | 0 | 0.91 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery type:threat |
| 30 | Invasion of Arcavios // Invocation of the Founders | synergy | 0 | 0.90 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 31 | Cursed Recording | synergy | 0 | 0.90 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 32 | Muddle, the Ever-Changing | threat | 0 | 0.90 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery type:threat |
| 33 | Rowan, Scholar of Sparks // Will, Scholar of Frost | removal | 0 | 0.89 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery tag:removal |
| 34 | Saheeli, Sublime Artificer | threat | 0 | 0.89 | payoff-text:whenever you cast a noncreature spell tag:copy type:threat |
| 35 | Pyromancer Ascension | synergy | 0 | 0.89 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 36 | Will Kenrith | draw | 0 | 0.88 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery tag:draw |
| 37 | Pinnacle Monk // Mystic Peak | land | 0 | 0.88 | payoff-text:whenever you cast a noncreature spell keyword:Prowess text:instant or sorcery type:land |
| 38 | The Mirari Conjecture | synergy | 0 | 0.88 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery theme |
| 39 | Mathise, Surge Channeler | draw | 0 | 0.88 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery tag:draw |
| 40 | Leitmotif Composer | draw | 0 | 0.87 | payoff-text:whenever you cast an instant or sorcery tag:copy text:instant or sorcery tag:draw |

Top upgrades (unowned):

- Thousand-Year Storm (synergy, 0.99)
- Primal Amulet // Primal Wellspring (land, 0.98)
- Kitsa, Otterball Elite (draw, 0.98)
- Ral, Crackling Wit (draw, 0.98)
- Rionya, Fire Dancer (threat, 0.97)
- Krark, the Thumbless (synergy, 0.97)
- Storm of Saruman (synergy, 0.97)
- Ral, Monsoon Mage // Ral, Leyline Prodigy (removal, 0.97)
- Sorcerer Class (ramp, 0.97)
- Stormsplitter (threat, 0.96)

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and legality, the EDHREC Spellslinger and Spell Copy pages (edhrec.com/tags/spellslinger, /spell-copy), and a score probe of `internal/candidates`.

Off theme (1):

- 36 Will Kenrith: a six-mana planeswalker. The +2 shrinks two creatures, and the -2 draws for a player. Only the -8 emblem copies spells. Prompt 14 counts Chandra, Torch of Defiance off theme for the same shape.

Weakest fits: Rionya, Fire Dancer and Muddle, the Ever-Changing turn a spell count into combat damage. Mathise, Surge Channeler needs mana value 3 or more. Leitmotif Composer needs mana value 5 or more. Each of the four still counts spells, so each stays on theme.

Shape: 34 of the 40 copy a spell. Twenty-seven trigger on "whenever you cast an instant or sorcery". Seven trigger on the second spell each turn. One card carries tag:cantrip.

Signal bugs: Magecraft counts as a keyword (weight 0.5), not as a payoff. Storm-Kiln Artist (58% of Spellslinger decks, and owned) falls to position 404, and Archmage Emeritus (52%) to position 409. Primal Amulet // Primal Wellspring and Pinnacle Monk // Mystic Peak get role land from the back face, the bug that prompt 14 records. Sorcerer Class, Chandra, Hope's Beacon, and Orvar, the All-Form get role ramp.

Legality: all 40 are legal in commander and inside UR. Primal Amulet // Primal Wellspring is colorless.

Owned: 2 of 40 are owned (Zada, Hedron Grinder and Gandalf the Grey). The funnel reports 111 owned on-theme cards.

Gaps (context only): the deck's own spells are absent. A cantrip carries the enabler tag alone, so its theme score caps at 0.40. Brainstorm (47% of Spellslinger decks), Ponder (40%), Opt (39%), and Preordain (36%) sit between position 247 and position 253. Frantic Search (51%) keeps only tag:draw and falls to position 1319.

The cheap payoff creatures also miss the cut. Guttersnipe (38%), Talrand, Sky Summoner, and Young Pyromancer score 0.74 at position 60 to position 62. Goblin Electromancer (37%) never enters the pool, because a cost reducer matches no signal and fills no staple role.

## 9. dragons (commander, RGB, owned-first)

Theme signals: payoff tags typal-dragon. subtypes Dragon.

Funnel: 18314 legal in colors, 328 on theme, 634 owned, 6 on theme and owned, 165 returned, 50 upgrades.

Thin theme: the collection holds under 30 on-theme cards. PR-7 asks the pool-mode question again here (D-63).

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Dragon Tempest | removal | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon you control payoff-text:dragons you control tag:removal |
| 2 | Lathliss, Dragon Queen | threat | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon you control payoff-text:dragons you control subtype:Dragon type:threat |
| 3 | Dragon's Hoard | ramp | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon you control tag:ramp |
| 4 | Dragonspeaker Shaman | synergy | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon spell theme |
| 5 | Atarka, World Render | threat | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon type:threat |
| 6 | Dragonlord's Servant | synergy | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon spell theme |
| 7 | Scourge of Valkas | removal | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon you control payoff-text:dragons you control subtype:Dragon tag:removal |
| 8 | Utvara Hellkite | threat | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon type:threat |
| 9 | Ganax, Astral Hunter | ramp | 0 | 0.99 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon tag:ramp |
| 10 | Maelstrom of the Spirit Dragon | land | 0 | 0.98 | payoff:typal-dragon payoff-text:dragon spell type:land |
| 11 | Wrathful Red Dragon | removal | 0 | 0.98 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon tag:removal |
| 12 | Dracogenesis | synergy | 0 | 0.98 | payoff:typal-dragon payoff-text:dragon spell theme |
| 13 | Encroaching Dragonstorm | ramp | 0 | 0.98 | payoff:typal-dragon payoff-text:dragon you control tag:ramp |
| 14 | Thrakkus the Butcher | threat | 0 | 0.98 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon type:threat |
| 15 | Dragonstorm Globe | ramp | 0 | 0.97 | payoff:typal-dragon payoff-text:dragon you control tag:ramp |
| 16 | Thunderbreak Regent | threat | 0 | 0.97 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon type:threat |
| 17 | Earthquake Dragon | threat | 0 | 0.97 | payoff:typal-dragon payoff-text:dragons you control subtype:Dragon type:threat |
| 18 | Bloomvine Regent // Claim Territory | ramp | 0 | 0.97 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon tag:ramp |
| 19 | Parapet Thrasher | removal | 0 | 0.97 | payoff:typal-dragon payoff-text:dragons you control subtype:Dragon tag:removal |
| 20 | Orb of Dragonkind | ramp | 0 | 0.96 | payoff:typal-dragon payoff-text:dragon spell tag:ramp |
| 21 | Broodcaller Scourge | ramp | 0 | 0.96 | payoff:typal-dragon payoff-text:dragons you control subtype:Dragon tag:ramp |
| 22 | Breaching Dragonstorm | synergy | 0 | 0.96 | payoff:typal-dragon payoff-text:dragon you control theme |
| 23 | Firespitter Whelp | synergy | 0 | 0.96 | payoff:typal-dragon payoff-text:dragon spell subtype:Dragon theme |
| 24 | Crucible of the Spirit Dragon | land | 0 | 0.96 | payoff:typal-dragon payoff-text:dragon spell type:land |
| 25 | Spit Flame | removal | 0 | 0.96 | payoff:typal-dragon payoff-text:dragon you control tag:removal |
| 26 | Nogi, Draco-Zealot | synergy | 0 | 0.96 | payoff:typal-dragon payoff-text:dragon spell theme |
| 27 | Stormscale Scion | threat | 0 | 0.95 | payoff:typal-dragon payoff-text:dragons you control subtype:Dragon type:threat |
| 28 | Sarkhan the Masterless | removal | 0 | 0.95 | payoff:typal-dragon payoff-text:dragon you control tag:removal |
| 29 | Sarkhan, Fireblood | ramp | 0 | 0.95 | payoff:typal-dragon payoff-text:dragon spell tag:ramp |
| 30 | Kolaghan, the Storm's Fury | threat | 0 | 0.94 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon type:threat |
| 31 | Dragon's Fire | removal | 0 | 0.93 | payoff:typal-dragon payoff-text:dragon you control tag:removal |
| 32 | Sarkhan, Dragon Ascendant | ramp | 0 | 0.93 | payoff:typal-dragon payoff-text:dragon you control tag:ramp |
| 33 | Invasion of Tarkir // Defiant Thundermaw | removal | 0 | 0.92 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon tag:removal |
| 34 | Ran and Shaw | ramp | 2 | 0.92 | payoff:typal-dragon payoff-text:dragons you control subtype:Dragon tag:ramp |
| 35 | Acolyte of Bahamut | synergy | 0 | 0.91 | payoff:typal-dragon payoff-text:dragon spell theme |
| 36 | Scaled Nurturer | ramp | 0 | 0.91 | payoff:typal-dragon subtype:Dragon tag:ramp |
| 37 | Bladewing the Risen | threat | 0 | 0.91 | payoff:typal-dragon subtype:Dragon type:threat |
| 38 | Boneyard Scourge | threat | 0 | 0.90 | payoff:typal-dragon payoff-text:dragon you control subtype:Dragon type:threat |
| 39 | Corroding Dragonstorm | synergy | 0 | 0.90 | payoff:typal-dragon payoff-text:dragon you control theme |
| 40 | Skanos Dragonheart | threat | 0 | 0.89 | payoff:typal-dragon payoff-text:dragons you control subtype:Dragon type:threat |

Top upgrades (unowned):

- Dragon Tempest (removal, 0.99)
- Lathliss, Dragon Queen (threat, 0.99)
- Dragon's Hoard (ramp, 0.99)
- Dragonspeaker Shaman (synergy, 0.99)
- Atarka, World Render (threat, 0.99)
- Dragonlord's Servant (synergy, 0.99)
- Scourge of Valkas (removal, 0.99)
- Utvara Hellkite (threat, 0.99)
- Ganax, Astral Hunter (ramp, 0.99)
- Maelstrom of the Spirit Dragon (land, 0.98)

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Dragons page (edhrec.com/tags/dragons).

Off theme (1):

- 39 Corroding Dragonstorm: a drain-2 and surveil-2 enchantment that returns to hand when a Dragon enters. Dragons page: 8% of decks, synergy 0.08, under the 0.10 line. This is a close call.

Ranks 1 to 11 are the Dragons page's own high-synergy cards (35% to 70% of decks). Seven cards are absent from the page but are Dragons or Dragon payoffs by oracle text: Firespitter Whelp, Nogi, Draco-Zealot, Stormscale Scion, Sarkhan, Dragon Ascendant, Ran and Shaw, Boneyard Scourge, and Skanos Dragonheart.

Signal bugs: Ran and Shaw gets role ramp from its firebending mana. It is a five-mana Dragon lord, so threat fits better. No wrong text match found.

Legality: all 40 are legal in commander and inside RGB (some are colorless).

Owned: 1 of 40 is owned (Ran and Shaw, rank 34). The funnel reports 6 owned on-theme cards.

Gaps (context only): Crux of Fate (65% of Dragon decks), Haven of the Spirit Dragon (56%), Terror of the Peaks (53%), Old Gnawbone (51%), Goldspan Dragon (49%), Ancient Copper Dragon (40%), Carnelian Orb of Dragonkind (37%), and Urza's Incubator (34%) are absent.

## 10. artifacts (commander, UR, owned-first)

Theme signals: payoff tags artifact-matters, synergy-artifact. tags mana-rock. keywords Affinity, Improvise, Metalcraft.

Funnel: 12588 legal in colors, 2150 on theme, 504 owned, 144 on theme and owned, 202 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Urza's Saga | land | 0 | 1.00 | payoff:synergy-artifact payoff-text:for each artifact you control text:artifact type:land |
| 2 | Storm-Kiln Artist | ramp | 1 | 1.00 | payoff:artifact-matters payoff:synergy-artifact payoff-text:for each artifact you control text:artifact tag:ramp |
| 3 | Mox Opal | ramp | 0 | 1.00 | payoff:synergy-artifact tag:mana-rock keyword:Metalcraft text:artifact tag:ramp |
| 4 | Unwinding Clock | synergy | 0 | 0.99 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact theme |
| 5 | Emry, Lurker of the Loch | synergy | 0 | 0.99 | payoff:synergy-artifact payoff-text:for each artifact you control keyword:Affinity text:artifact theme |
| 6 | Thought Monitor | draw | 0 | 0.99 | payoff:synergy-artifact payoff-text:for each artifact you control keyword:Affinity text:artifact tag:draw |
| 7 | Urza, Lord High Artificer | ramp | 0 | 0.99 | payoff:synergy-artifact payoff-text:for each artifact you control text:artifact tag:ramp |
| 8 | Padeem, Consul of Innovation | interaction | 0 | 0.99 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact tag:interaction |
| 9 | Thoughtcast | draw | 0 | 0.99 | payoff:synergy-artifact payoff-text:for each artifact you control keyword:Affinity text:artifact tag:draw |
| 10 | Inspiring Statuary | ramp | 1 | 0.99 | payoff:synergy-artifact tag:mana-rock text:artifact tag:ramp |
| 11 | Krark-Clan Ironworks | ramp | 0 | 0.99 | payoff:synergy-artifact tag:mana-rock text:artifact tag:ramp |
| 12 | Forensic Gadgeteer | draw | 0 | 0.99 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact tag:draw |
| 13 | Uthros, Titanic Godcore | land | 0 | 0.99 | payoff:synergy-artifact payoff-text:for each artifact you control text:artifact type:land |
| 14 | Darksteel Forge | interaction | 0 | 0.99 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact tag:interaction |
| 15 | Fomori Vault | land | 0 | 0.98 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact type:land |
| 16 | Metalwork Colossus | threat | 0 | 0.98 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact type:threat |
| 17 | Tezzeret the Seeker | ramp | 0 | 0.98 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact tag:ramp |
| 18 | Clock of Omens | synergy | 0 | 0.98 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact theme |
| 19 | Simulacrum Synthesizer | synergy | 0 | 0.98 | payoff:synergy-artifact payoff-text:for each artifact you control text:artifact theme |
| 20 | Ingenious Artillerist | synergy | 0 | 0.98 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact theme |
| 21 | Uthros Research Craft | draw | 0 | 0.98 | payoff:synergy-artifact payoff-text:for each artifact you control text:artifact tag:draw |
| 22 | Crystal Skull, Isu Spyglass | ramp | 0 | 0.98 | payoff:synergy-artifact tag:mana-rock text:artifact tag:ramp |
| 23 | Adaptive Omnitool | synergy | 0 | 0.98 | payoff:synergy-artifact payoff-text:for each artifact you control text:artifact theme |
| 24 | Master of Etherium | synergy | 0 | 0.98 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact theme |
| 25 | Grinding Station | synergy | 0 | 0.98 | payoff:synergy-artifact payoff-text:whenever an artifact enters text:artifact theme |
| 26 | Whirler Rogue | threat | 0 | 0.98 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact type:threat |
| 27 | Karn, Legacy Reforged | ramp | 0 | 0.98 | payoff:synergy-artifact payoff-text:for each artifact you control payoff-text:artifacts you control text:artifact tag:ramp |
| 28 | Shimmer Dragon | interaction | 0 | 0.97 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact tag:interaction |
| 29 | One with the Machine | draw | 0 | 0.97 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact tag:draw |
| 30 | The Mightstone and Weakstone | ramp | 0 | 0.97 | payoff:synergy-artifact tag:mana-rock text:artifact tag:ramp |
| 31 | Mycosynth Golem | threat | 0 | 0.97 | payoff:synergy-artifact payoff-text:for each artifact you control keyword:Affinity text:artifact type:threat |
| 32 | Storm the Vault // Vault of Catlacan | land | 0 | 0.97 | payoff:synergy-artifact payoff-text:for each artifact you control text:artifact type:land |
| 33 | Kappa Cannoneer | threat | 0 | 0.96 | payoff:synergy-artifact keyword:Improvise text:artifact type:threat |
| 34 | Secluded Starforge | land | 0 | 0.96 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact type:land |
| 35 | Whir of Invention | synergy | 0 | 0.96 | payoff:synergy-artifact keyword:Improvise text:artifact theme |
| 36 | Darksteel Juggernaut | threat | 0 | 0.96 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact type:threat |
| 37 | Depthshaker Titan | threat | 0 | 0.96 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact type:threat |
| 38 | Solar Array | ramp | 0 | 0.96 | payoff:synergy-artifact tag:mana-rock text:artifact tag:ramp |
| 39 | Brotherhood Vertibird | synergy | 0 | 0.96 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact theme |
| 40 | Mm'menon, the Right Hand | ramp | 0 | 0.95 | payoff:synergy-artifact payoff-text:artifacts you control text:artifact text:mana |

Top upgrades (unowned):

- Urza's Saga (land, 1.00)
- Mox Opal (ramp, 1.00)
- Unwinding Clock (synergy, 0.99)
- Emry, Lurker of the Loch (synergy, 0.99)
- Thought Monitor (draw, 0.99)
- Urza, Lord High Artificer (ramp, 0.99)
- Padeem, Consul of Innovation (interaction, 0.99)
- Thoughtcast (draw, 0.99)
- Krark-Clan Ironworks (ramp, 0.99)
- Forensic Gadgeteer (draw, 0.99)

### Review

Verdict: yes. 37 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text, the EDHREC Artifacts page (edhrec.com/tags/artifacts), and EDHREC card pages for the three close calls.

Off theme (3):

- 2 Storm-Kiln Artist: gets +1/+0 for each artifact, but the Treasure trigger needs instant or sorcery casts. Its EDHREC card page shows only spellslinger commanders as top homes, and the average deck that runs it has 9 artifacts. Not on the Artifacts page. Score 1.00 overstates it. This is a close call.
- 30 The Mightstone and Weakstone: a five-mana rock whose mana pays only for artifact spells. Artifacts page: 6.5% of decks, synergy 0.06. This is a close call.
- 38 Solar Array: a three-mana rock for any color. The only artifact text is a sunburst rider for the next artifact spell, worth at most two counters in UR. Not on the Artifacts page. This is a close call.

Signal bugs: Storm the Vault // Vault of Catlacan gets role land and type:land from its back face. The front face is an enchantment. Inspiring Statuary gets tag:mana-rock but has no mana ability. Shimmer Dragon, Padeem, and Darksteel Forge get tag:interaction for self-protection effects.

Legality: all 40 are legal in commander and inside UR.

Owned: 2 of 40 are owned (ranks 2 and 10). The funnel reports 144 owned on-theme cards.

Gaps (context only): the highest-synergy artifact cards are absent. Seat of the Synod (59% of Artifacts decks), Etherium Sculptor (53%), Great Furnace (48%), Foundry Inspector (42%), Jhoira, Weatherlight Captain (39%), Sai, Master Thopterist (34%), Inventors' Fair (32%), and Mystic Forge (29%). The signal set seems to lack "Artifact Land", "artifact spells cost {1} less", and "whenever you cast an artifact spell".

## 11. lifegain (commander, WB, any-card)

Theme signals: payoff tags life-total-matters-self, lifegain-matters. tags drain-life, lifegain, repeatable-lifegain. keywords Lifelink.

Funnel: 12683 legal in colors, 1715 on theme, 0 owned, 0 on theme and owned, 300 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Aetherflux Reservoir | removal | 0 | 1.00 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain text:you gain tag:removal |
| 2 | Vito, Thorn of the Dusk Rose | synergy | 0 | 1.00 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 3 | Sanguine Bond | synergy | 0 | 1.00 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 4 | Enduring Tenacity | threat | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain type:threat |
| 5 | Serra Ascendant | synergy | 0 | 0.99 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain keyword:Lifelink theme |
| 6 | The Wind Crystal | synergy | 0 | 0.99 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 7 | Heliod, Sun-Crowned | synergy | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 8 | Marauding Blight-Priest | synergy | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 9 | Ocelot Pride | synergy | 0 | 0.99 | payoff:lifegain-matters payoff-text:if you gained life this turn tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain theme |
| 10 | Well of Lost Dreams | draw | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life payoff-text:life you gained text:you gain tag:draw |
| 11 | Exemplar of Light | draw | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain tag:draw |
| 12 | Caduceus, Staff of Hermes | interaction | 0 | 0.99 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain tag:interaction |
| 13 | Archangel of Thune | threat | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain type:threat |
| 14 | Cleric Class | synergy | 0 | 0.99 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain text:you gain theme |
| 15 | Haliya, Guided by Light | draw | 0 | 0.98 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain text:you gain tag:draw |
| 16 | Felidar Sovereign | wincon | 0 | 0.98 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain keyword:Lifelink tag:alternate-win-condition |
| 17 | Elenda's Hierophant | synergy | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain text:you gain theme |
| 18 | Starscape Cleric | synergy | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 19 | Righteous Valkyrie | synergy | 0 | 0.98 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 20 | Resplendent Angel | synergy | 0 | 0.98 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain text:you gain theme |
| 21 | Aerith Gainsborough | synergy | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain theme |
| 22 | Dawn of Hope | draw | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain text:you gain tag:draw |
| 23 | Witch of the Moors | removal | 0 | 0.98 | payoff:lifegain-matters payoff-text:if you gained life this turn text:you gain tag:removal |
| 24 | Ajani's Pridemate | synergy | 0 | 0.98 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 25 | Cosmos Elixir | draw | 0 | 0.97 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain text:you gain tag:draw |
| 26 | Valkyrie Harbinger | threat | 0 | 0.97 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain type:threat |
| 27 | Celestine, the Living Saint | threat | 0 | 0.97 | payoff:lifegain-matters payoff-text:life you gained tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain type:threat |
| 28 | Angel of Destiny | wincon | 0 | 0.97 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain tag:alternate-win-condition |
| 29 | Nykthos Paragon | threat | 0 | 0.97 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain type:threat |
| 30 | Lunar Convocation | draw | 0 | 0.97 | payoff:lifegain-matters payoff-text:if you gained life this turn text:you gain tag:draw |
| 31 | Voice of the Blessed | synergy | 0 | 0.97 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 32 | Ajani, Strength of the Pride | wipe | 0 | 0.97 | payoff:lifegain-matters payoff:life-total-matters-self payoff-text:whenever you gain life tag:lifegain tag:repeatable-lifegain text:you gain tag:sweeper |
| 33 | Astarion, the Decadent | threat | 0 | 0.97 | payoff:lifegain-matters payoff-text:life you gained tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain type:threat |
| 34 | Speaker of the Heavens | synergy | 0 | 0.96 | payoff:lifegain-matters payoff:life-total-matters-self tag:lifegain tag:repeatable-lifegain keyword:Lifelink theme |
| 35 | Karlov of the Ghost Council | removal | 0 | 0.96 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain tag:removal |
| 36 | Sorin of House Markov // Sorin, Ravenous Neonate | removal | 0 | 0.96 | payoff:lifegain-matters payoff-text:life you gained tag:lifegain tag:repeatable-lifegain tag:drain-life keyword:Lifelink text:you gain tag:removal |
| 37 | Essence Channeler | synergy | 0 | 0.96 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 38 | Amalia Benavides Aguirre | wipe | 0 | 0.96 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain tag:sweeper |
| 39 | Veinwitch Coven | synergy | 0 | 0.96 | payoff:lifegain-matters payoff-text:whenever you gain life text:you gain theme |
| 40 | Indulging Patrician | synergy | 0 | 0.96 | payoff:lifegain-matters tag:lifegain tag:repeatable-lifegain keyword:Lifelink text:you gain theme |

### Review

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24.

The 40 cards, ranks, roles, scores, and signals are identical to prompt 1. The two owned cards in prompt 1 (Exemplar of Light at rank 11, Aerith Gainsborough at rank 21) hold the same ranks here on score alone. So owned-first mode did not promote any owned card into the top 40. See the prompt 1 review for the card-level findings, signal bugs, and gaps.

## 12. mill (commander, UB, any-card)

Theme signals: tags mill, mill-any, mill-opponent. keywords Mill.

Funnel: 12600 legal in colors, 801 on theme, 0 owned, 0 on theme and owned, 294 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Sheoldred // The True Scriptures | removal | 0 | 0.98 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills tag:removal |
| 2 | Riverchurn Monument | synergy | 0 | 0.97 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill-any tag:mill keyword:Mill theme |
| 3 | Duskmantle Guildmage | synergy | 0 | 0.94 | payoff-text:whenever a card is put into an opponent's graveyard tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 4 | Relic Golem | synergy | 0 | 0.92 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 5 | Thieves' Guild Enforcer | synergy | 0 | 0.92 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 6 | Merfolk Windrobber | draw | 0 | 0.91 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills tag:draw |
| 7 | Deepmuck Desperado | synergy | 0 | 0.91 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 8 | Soaring Thought-Thief | synergy | 0 | 0.90 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 9 | Vantress Gargoyle | synergy | 0 | 0.89 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 10 | Syr Konrad, the Grim | threat | 0 | 0.83 | tag:mill-opponent tag:mill keyword:Mill text:mills type:threat |
| 11 | Altar of Dementia | synergy | 0 | 0.83 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 12 | Brain Freeze | synergy | 0 | 0.82 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 13 | Breach the Multiverse | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 14 | Mindcrank | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 15 | Mesmeric Orb | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 16 | Jace, Wielder of Mysteries | wincon | 0 | 0.82 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills tag:alternate-win-condition |
| 17 | Ruin Crab | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 18 | Altar of the Brood | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 19 | Hedron Crab | synergy | 0 | 0.82 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 20 | The Water Crystal | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 21 | Maddening Cacophony | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 22 | Psychic Corrosion | draw | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills text:draw |
| 23 | Thought Scour | draw | 0 | 0.82 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills tag:draw |
| 24 | Fractured Sanity | draw | 0 | 0.81 | tag:mill-opponent tag:mill keyword:Mill text:mills tag:draw |
| 25 | Drown in Dreams | draw | 0 | 0.81 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills tag:draw |
| 26 | The Mindskinner | synergy | 0 | 0.81 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 27 | Grinding Station | synergy | 0 | 0.81 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 28 | Jidoor, Aristocratic Capital // Overture | land | 0 | 0.81 | tag:mill-opponent tag:mill keyword:Mill text:mills type:land |
| 29 | Codex Shredder | synergy | 0 | 0.81 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 30 | Ashiok, Dream Render | threat | 0 | 0.81 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills type:threat |
| 31 | Court of Cunning | draw | 0 | 0.81 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills tag:draw |
| 32 | Traumatize | synergy | 0 | 0.81 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 33 | Folio of Fancies | draw | 0 | 0.81 | tag:mill-opponent tag:mill keyword:Mill text:mills tag:draw |
| 34 | Fraying Sanity | synergy | 0 | 0.81 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 35 | Zellix, Sanity Flayer | synergy | 0 | 0.80 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 36 | One Ring to Rule Them All | wipe | 0 | 0.80 | tag:mill-opponent tag:mill keyword:Mill text:mills tag:sweeper |
| 37 | Extract from Darkness | synergy | 0 | 0.80 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 38 | Nephalia Drownyard | land | 0 | 0.80 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills type:land |
| 39 | Cut Your Losses | synergy | 0 | 0.80 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 40 | Didn't Say Please | interaction | 0 | 0.80 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills tag:interaction |

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Mill page (edhrec.com/tags/mill).

Off theme (1):

- 16 Jace, Wielder of Mysteries: the +1 mills two, but the static ability and the -8 are a self-mill win condition. Mill page: 9% of decks, synergy 0.07, the lowest of any listed card here. This is a close call.

Signal bugs: Deepmuck Desperado gets payoff-text:cards in their graveyard from the crime reminder text. Psychic Corrosion gets role draw and text:draw from "whenever you draw a card". It does not draw. The payoff-text signal lifts eight-card-threshold bodies (ranks 4 to 9) above the archetype's engines (Ruin Crab, Mesmeric Orb, and Maddening Cacophony at ranks 15 to 21).

Legality: all 40 are legal in commander and inside UB.

Gaps (context only): Consuming Aberration (41% of Mill decks), Drown in the Loch (33%), Bruvac the Grandiloquent (28%), Memory Erosion (21%), Nemesis of Reason (18%), Bloodchief Ascension (16%), Mind Grind (16%), and Glimpse the Unthinkable (15%) are absent.

## 13. elves (commander, GB, any-card)

Theme signals: payoff tags typal-elf. subtypes Elf.

Funnel: 12525 legal in colors, 529 on theme, 0 owned, 0 on theme and owned, 292 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Priest of Titania | ramp | 0 | 0.99 | payoff:typal-elf payoff-text:for each elf subtype:Elf tag:ramp |
| 2 | Elvish Archdruid | ramp | 0 | 0.99 | payoff:typal-elf payoff-text:elf you control payoff-text:for each elf subtype:Elf tag:ramp |
| 3 | Marwyn, the Nurturer | ramp | 0 | 0.99 | payoff:typal-elf payoff-text:elf you control subtype:Elf tag:ramp |
| 4 | Imperious Perfect | synergy | 0 | 0.99 | payoff:typal-elf payoff-text:elves you control subtype:Elf theme |
| 5 | Elvish Warmaster | synergy | 0 | 0.98 | payoff:typal-elf payoff-text:elves you control subtype:Elf theme |
| 6 | Leaf-Crowned Visionary | draw | 0 | 0.98 | payoff:typal-elf payoff-text:elves you control subtype:Elf tag:draw |
| 7 | Lathril, Blade of the Elves | threat | 0 | 0.98 | payoff:typal-elf payoff-text:elves you control subtype:Elf type:threat |
| 8 | Dionus, Elvish Archdruid | threat | 0 | 0.98 | payoff:typal-elf payoff-text:elves you control subtype:Elf type:threat |
| 9 | Dwynen, Gilt-Leaf Daen | threat | 0 | 0.98 | payoff:typal-elf payoff-text:elf you control subtype:Elf type:threat |
| 10 | Elven Ambush | synergy | 0 | 0.97 | payoff:typal-elf payoff-text:elf you control payoff-text:for each elf theme |
| 11 | Wolverine Riders | threat | 0 | 0.97 | payoff:typal-elf payoff-text:elf you control subtype:Elf type:threat |
| 12 | Elvish Promenade | synergy | 0 | 0.97 | payoff:typal-elf payoff-text:elf you control payoff-text:for each elf subtype:Elf theme |
| 13 | Canopy Tactician | ramp | 0 | 0.97 | payoff:typal-elf payoff-text:elves you control subtype:Elf tag:ramp |
| 14 | Wirewood Symbiote | synergy | 0 | 0.97 | payoff:typal-elf payoff-text:elf you control theme |
| 15 | Immaculate Magistrate | threat | 0 | 0.97 | payoff:typal-elf payoff-text:elf you control payoff-text:for each elf subtype:Elf type:threat |
| 16 | Galadhrim Brigade | synergy | 0 | 0.96 | payoff:typal-elf payoff-text:elves you control subtype:Elf theme |
| 17 | High Perfect Morcant | removal | 0 | 0.96 | payoff:typal-elf payoff-text:elves you control payoff-text:elf you control subtype:Elf tag:removal |
| 18 | Joraga Treespeaker | ramp | 0 | 0.96 | payoff:typal-elf payoff-text:elves you control subtype:Elf tag:ramp |
| 19 | Tyvar Kell | ramp | 0 | 0.96 | payoff:typal-elf payoff-text:elves you control tag:ramp |
| 20 | Champions of the Perfect | draw | 0 | 0.96 | payoff:typal-elf payoff-text:elf you control subtype:Elf tag:draw |
| 21 | Shaman of the Pack | synergy | 0 | 0.96 | payoff:typal-elf payoff-text:elves you control subtype:Elf theme |
| 22 | Abomination of Llanowar | synergy | 0 | 0.95 | payoff:typal-elf payoff-text:elves you control subtype:Elf theme |
| 23 | Wellwisher | synergy | 0 | 0.95 | payoff:typal-elf payoff-text:for each elf subtype:Elf theme |
| 24 | Elderfang Venom | synergy | 0 | 0.95 | payoff:typal-elf payoff-text:elves you control payoff-text:elf you control theme |
| 25 | Morcant's Loyalist | synergy | 0 | 0.95 | payoff:typal-elf payoff-text:elves you control subtype:Elf theme |
| 26 | Heritage Druid | ramp | 0 | 0.94 | payoff:typal-elf payoff-text:elves you control subtype:Elf tag:ramp |
| 27 | Tyvar the Bellicose | threat | 0 | 0.94 | payoff:typal-elf payoff-text:elves you control subtype:Elf type:threat |
| 28 | Elvish Guidance | ramp | 0 | 0.94 | payoff:typal-elf payoff-text:for each elf tag:ramp |
| 29 | Haldir, Lórien Lieutenant | synergy | 0 | 0.94 | payoff:typal-elf payoff-text:elves you control subtype:Elf theme |
| 30 | Allosaurus Shepherd | synergy | 0 | 0.93 | payoff:typal-elf subtype:Elf theme |
| 31 | Miara, Thorn of the Glade | draw | 0 | 0.93 | payoff:typal-elf payoff-text:elf you control subtype:Elf tag:draw |
| 32 | Trystan's Command | removal | 0 | 0.93 | payoff:typal-elf payoff-text:elf you control subtype:Elf tag:removal |
| 33 | Jagged-Scar Archers | removal | 0 | 0.93 | payoff:typal-elf payoff-text:elves you control subtype:Elf tag:removal |
| 34 | Nissa, Resurgent Animist | ramp | 0 | 0.92 | payoff:typal-elf subtype:Elf tag:ramp |
| 35 | Crown of Skemfar | synergy | 0 | 0.92 | payoff:typal-elf payoff-text:elf you control payoff-text:for each elf theme |
| 36 | Lys Alana Huntmaster | threat | 0 | 0.92 | payoff:typal-elf subtype:Elf type:threat |
| 37 | Ezuri, Renegade Leader | interaction | 0 | 0.92 | payoff:typal-elf subtype:Elf tag:interaction |
| 38 | Rhys the Exiled | synergy | 0 | 0.92 | payoff:typal-elf payoff-text:elf you control payoff-text:for each elf subtype:Elf theme |
| 39 | Elvish Harbinger | ramp | 0 | 0.92 | payoff:typal-elf subtype:Elf tag:ramp |
| 40 | Elvish Champion | synergy | 0 | 0.92 | payoff:typal-elf subtype:Elf theme |

### Review

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Elves page (edhrec.com/tags/elves).

Off theme: none. Every card is an Elf creature, an Elf kindred spell, or a card with Elf-specific text. Thirty-one of the 40 are on the Elves page. Nissa, Resurgent Animist (rank 34) is the weakest fit: she is a legendary Elf creature, but her main function is landfall mana.

Signal bugs: none found. Tyvar Kell (a planeswalker) correctly has no subtype:Elf. Elvish Promenade and Trystan's Command carry subtype:Elf because they are Kindred Elf spells.

Legality: all 40 are legal in commander and inside GB.

Gaps (context only): the archetype's most-played Elves lack Elf-payoff text, so they fall below rank 40. Llanowar Elves (81% of Elves decks), Elvish Mystic (79%), Beast Whisperer (65%), Reclamation Sage (65%), Fyndhorn Elves (57%), Circle of Dreams Druid (40%), Galadhrim Ambush (38%), and Wirewood Channeler (30%).

## 14. storm (commander, UR, any-card)

Theme signals: payoff tags fourth-spell-matters, second-spell-matters, storm-count-matters, third-spell-matters. tags cantrip, ramp, storm-like. keywords Storm.

Funnel: 12588 legal in colors, 1761 on theme, 0 owned, 0 on theme and owned, 296 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Aetherflux Reservoir | removal | 0 | 1.00 | payoff:storm-count-matters payoff-text:whenever you cast tag:storm-like tag:removal |
| 2 | Thousand-Year Storm | synergy | 0 | 0.99 | payoff:storm-count-matters payoff-text:whenever you cast tag:storm-like theme |
| 3 | Case of the Ransacked Lab | draw | 0 | 0.97 | payoff:fourth-spell-matters payoff-text:whenever you cast tag:draw |
| 4 | Storm of Saruman | synergy | 0 | 0.97 | payoff:second-spell-matters payoff-text:whenever you cast theme |
| 5 | Ral, Monsoon Mage // Ral, Leyline Prodigy | removal | 0 | 0.97 | payoff:storm-count-matters payoff-text:whenever you cast tag:removal |
| 6 | Sorcerer Class | ramp | 0 | 0.97 | payoff:storm-count-matters payoff-text:whenever you cast tag:storm-like tag:ramp tag:ramp |
| 7 | Cori-Steel Cutter | synergy | 0 | 0.96 | payoff:second-spell-matters payoff-text:whenever you cast theme |
| 8 | Alphinaud Leveilleur | draw | 0 | 0.96 | payoff:second-spell-matters payoff-text:whenever you cast tag:draw |
| 9 | Jori En, Ruin Diver | draw | 0 | 0.95 | payoff:second-spell-matters payoff-text:whenever you cast tag:draw |
| 10 | Stella Lee, Wild Card | synergy | 0 | 0.95 | payoff:second-spell-matters payoff-text:whenever you cast theme |
| 11 | Taigam, Master Opportunist | synergy | 0 | 0.95 | payoff:second-spell-matters payoff-text:whenever you cast theme |
| 12 | Lock and Load | draw | 0 | 0.94 | payoff:storm-count-matters tag:storm-like tag:draw |
| 13 | Tomb of Horrors Adventurer | threat | 0 | 0.94 | payoff:second-spell-matters payoff-text:whenever you cast payoff-text:copy that spell type:threat |
| 14 | Eris, Roar of the Storm | threat | 0 | 0.94 | payoff:second-spell-matters payoff-text:whenever you cast type:threat |
| 15 | Geralf, the Fleshwright | synergy | 0 | 0.94 | payoff:storm-count-matters payoff-text:whenever you cast theme |
| 16 | Malcolm, the Eyes | draw | 0 | 0.93 | payoff:second-spell-matters payoff-text:whenever you cast tag:draw |
| 17 | Saruman the White | threat | 0 | 0.93 | payoff:second-spell-matters payoff-text:whenever you cast type:threat |
| 18 | Slick Sequence | removal | 0 | 0.92 | payoff:second-spell-matters tag:cantrip tag:removal |
| 19 | Aria of Flame | removal | 0 | 0.92 | payoff:storm-count-matters payoff-text:whenever you cast tag:storm-like tag:removal |
| 20 | Breeches, the Blastmaker | removal | 0 | 0.92 | payoff:second-spell-matters payoff-text:whenever you cast payoff-text:copy that spell tag:removal |
| 21 | Emeritus of Conflict // Lightning Bolt | removal | 0 | 0.92 | payoff:third-spell-matters payoff-text:whenever you cast tag:removal |
| 22 | Kraum, Violent Cacophony | draw | 0 | 0.92 | payoff:second-spell-matters payoff-text:whenever you cast tag:draw |
| 23 | Storm-Kiln Artist | ramp | 0 | 0.91 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 24 | Birgi, God of Storytelling // Harnfel, Horn of Bounty | ramp | 0 | 0.91 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 25 | Sentinel Tower | removal | 0 | 0.91 | payoff:storm-count-matters tag:storm-like tag:removal |
| 26 | Forsaken Monument | ramp | 0 | 0.91 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 27 | Vivi Ornitier | ramp | 0 | 0.91 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 28 | Captain Lannery Storm | ramp | 0 | 0.91 | payoff-text:storm tag:ramp tag:ramp |
| 29 | Glaring Fleshraker | ramp | 0 | 0.90 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 30 | Electro, Assaulting Battery | ramp | 0 | 0.90 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 31 | Chandra, Torch of Defiance | ramp | 0 | 0.90 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 32 | Ashling, Flame Dancer | wipe | 0 | 0.90 | payoff-text:whenever you cast tag:ramp tag:sweeper |
| 33 | Brass's Tunnel-Grinder // Tecutlan, the Searing Rift | land | 0 | 0.90 | payoff-text:whenever you cast tag:cantrip tag:ramp type:land |
| 34 | Runaway Steam-Kin | ramp | 0 | 0.90 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 35 | Ugin, Eye of the Storms | ramp | 0 | 0.90 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 36 | Urabrask // The Great Work | wipe | 0 | 0.90 | payoff-text:whenever you cast tag:ramp tag:sweeper |
| 37 | Brimstone Roundup | synergy | 0 | 0.90 | payoff:second-spell-matters payoff-text:whenever you cast theme |
| 38 | Alchemist's Talent | ramp | 0 | 0.90 | payoff-text:whenever you cast tag:ramp tag:ramp |
| 39 | Primal Amulet // Primal Wellspring | land | 0 | 0.90 | payoff-text:whenever you cast payoff-text:copy that spell tag:ramp type:land |
| 40 | Sword of Wealth and Power | ramp | 0 | 0.90 | payoff-text:copy that spell tag:ramp tag:ramp |

### Review

Verdict: no. 25 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text, the EDHREC Storm, Spellslinger, and Spell Copy pages (edhrec.com/tags/storm, /spellslinger, /spell-copy), and the Wizards mechanics article for Secrets of Strixhaven.

Off theme (15):

- 7 Cori-Steel Cutter: second-spell trigger, but the output is a 1/1 Monk token and Equipment stats. A prowess aggro card.
- 14 Eris, Roar of the Storm: second-spell trigger makes a 4/4 Dragon token. A ten-mana combat finisher.
- 15 Geralf, the Fleshwright: each spell after the first makes a 2/2 Zombie. A go-wide win path storm decks do not use.
- 16 Malcolm, the Eyes: second-spell trigger investigates. A Clue costs {2} to crack, so it does not fuel the turn.
- 17 Saruman the White: second-spell trigger amasses Orcs 2. A combat threat.
- 21 Emeritus of Conflict // Lightning Bolt: third-spell trigger lets you cast a copy of Lightning Bolt for {R}. A spells-matter aggro creature (Secrets of Strixhaven, 2026-04-24).
- 26 Forsaken Monument: triggers on colorless spells only. Colorless anthem and ramp.
- 28 Captain Lannery Storm: a name match. The text is haste, a Treasure on attack, and +1/+0 on Treasure sacrifice.
- 29 Glaring Fleshraker: triggers on colorless spells only. An Eldrazi card.
- 31 Chandra, Torch of Defiance: a generic red planeswalker. The only cast link is the -7 emblem.
- 33 Brass's Tunnel-Grinder // Tecutlan, the Searing Rift: a one-shot loot that transforms after three turns of descend. The land triggers on permanent spells.
- 35 Ugin, Eye of the Storms: a name match plus a colorless-spell trigger. A colorless-matters card.
- 37 Brimstone Roundup: second-spell trigger makes a 1/1 Mercenary. A go-wide combat payoff.
- 38 Alchemist's Talent: Treasure ramp. Only level 3 ({4}{R} more) adds a spell trigger.
- 40 Sword of Wealth and Power: combat Equipment. The spell copy needs combat damage to a player.

None of the 15 is on the Storm, Spellslinger, or Spell Copy pages.

Root cause: payoff-text:whenever you cast matches "whenever you cast a colorless spell" (three cards) and the -7 emblem of Chandra. payoff-text:storm matches card names (Captain Lannery Storm, Ugin, Eye of the Storms). payoff:second-spell-matters treats every "second spell each turn" trigger as storm, but six of them make combat tokens. tag:ramp appears twice on thirteen cards. Brass's Tunnel-Grinder and Primal Amulet get role land from their back faces.

Legality: all 40 are legal in commander and inside UR.

Gaps (context only): Scryfall lists 22 Storm-keyword cards legal in commander inside UR. None is in the top 40. Jeska's Will (54% of Storm decks), Frantic Search (51%), Underworld Breach (42%), Grapeshot (40%), Archmage Emeritus (38%), Seething Song (37%), Brain Freeze (36%), and Goblin Electromancer (24%) are absent.

## 15. enchantress (commander, GW, any-card)

Theme signals: payoff tags synergy-enchantment. keywords Constellation.

Funnel: 12612 legal in colors, 851 on theme, 0 owned, 0 on theme and owned, 291 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Sanctum Weaver | ramp | 0 | 0.99 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment tag:ramp |
| 2 | Sphere of Safety | interaction | 0 | 0.99 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment tag:interaction |
| 3 | Mesa Enchantress | draw | 0 | 0.99 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment tag:draw |
| 4 | Destiny Spinner | synergy | 0 | 0.99 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment theme |
| 5 | Enchantress's Presence | draw | 0 | 0.99 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment tag:draw |
| 6 | Sythis, Harvest's Hand | draw | 0 | 0.99 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment tag:draw |
| 7 | Ethereal Armor | synergy | 0 | 0.99 | payoff:synergy-enchantment payoff-text:for each enchantment text:enchantment theme |
| 8 | Sterling Grove | interaction | 0 | 0.99 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment tag:interaction |
| 9 | Sigil of the Empty Throne | synergy | 0 | 0.99 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment theme |
| 10 | Herald of the Pantheon | synergy | 0 | 0.98 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment theme |
| 11 | Satyr Enchanter | draw | 0 | 0.98 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment tag:draw |
| 12 | Verduran Enchantress | draw | 0 | 0.98 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment tag:draw |
| 13 | Greater Auramancy | interaction | 0 | 0.98 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment tag:interaction |
| 14 | Serra's Sanctum | land | 0 | 0.97 | payoff:synergy-enchantment payoff-text:for each enchantment text:enchantment type:land |
| 15 | Hallowed Haunting | synergy | 0 | 0.97 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment theme |
| 16 | Argothian Enchantress | draw | 0 | 0.97 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment tag:draw |
| 17 | Eidolon of Blossoms | draw | 0 | 0.96 | payoff:synergy-enchantment keyword:Constellation text:enchantment tag:draw |
| 18 | Setessan Champion | draw | 0 | 0.96 | payoff:synergy-enchantment keyword:Constellation text:enchantment tag:draw |
| 19 | Secret Arcade // Dusty Parlor | synergy | 0 | 0.96 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment theme |
| 20 | Archon of Sun's Grace | threat | 0 | 0.96 | payoff:synergy-enchantment keyword:Constellation text:enchantment type:threat |
| 21 | Helm of the Gods | synergy | 0 | 0.96 | payoff:synergy-enchantment payoff-text:for each enchantment text:enchantment theme |
| 22 | Tempest Technique | synergy | 0 | 0.95 | payoff:synergy-enchantment payoff-text:for each enchantment text:enchantment theme |
| 23 | Ellivere of the Wild Court | draw | 0 | 0.94 | payoff:synergy-enchantment payoff-text:for each enchantment text:enchantment tag:draw |
| 24 | Calix, Guided by Fate | synergy | 0 | 0.94 | payoff:synergy-enchantment keyword:Constellation text:enchantment theme |
| 25 | Composer of Spring | ramp | 0 | 0.94 | payoff:synergy-enchantment keyword:Constellation text:enchantment tag:ramp |
| 26 | Nyxborn Behemoth | threat | 0 | 0.93 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment type:threat |
| 27 | Boon of the Spirit Realm | synergy | 0 | 0.93 | payoff:synergy-enchantment keyword:Constellation text:enchantment theme |
| 28 | Elvish Archivist | draw | 0 | 0.92 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment tag:draw |
| 29 | Nylea's Colossus | threat | 0 | 0.90 | payoff:synergy-enchantment keyword:Constellation text:enchantment type:threat |
| 30 | Kami of Transience | synergy | 0 | 0.90 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment theme |
| 31 | Shambling Suit | synergy | 0 | 0.90 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment theme |
| 32 | Generous Visitor | synergy | 0 | 0.90 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment theme |
| 33 | Nessian Wanderer | synergy | 0 | 0.90 | payoff:synergy-enchantment keyword:Constellation text:enchantment theme |
| 34 | Slumbering Keepguard | synergy | 0 | 0.89 | payoff:synergy-enchantment payoff-text:for each enchantment text:enchantment theme |
| 35 | Celestial Ancient | threat | 0 | 0.88 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment type:threat |
| 36 | Yavimaya Enchantress | synergy | 0 | 0.87 | payoff:synergy-enchantment payoff-text:for each enchantment text:enchantment theme |
| 37 | Skybind | removal | 0 | 0.87 | payoff:synergy-enchantment keyword:Constellation text:enchantment tag:removal |
| 38 | Fountain Watch | interaction | 0 | 0.86 | payoff:synergy-enchantment payoff-text:enchantments you control text:enchantment tag:interaction |
| 39 | Sky-Blessed Samurai | threat | 0 | 0.84 | payoff:synergy-enchantment payoff-text:for each enchantment text:enchantment type:threat |
| 40 | Dawnhart Geist | synergy | 0 | 0.83 | payoff:synergy-enchantment payoff-text:whenever you cast an enchantment text:enchantment theme |

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Enchantress and Auras pages (edhrec.com/tags/enchantress, /auras).

Off theme (1):

- 31 Shambling Suit: power equals the number of artifacts and enchantments you control. An Eldraine artifact-theme body. Not on the Enchantress or Auras pages. EDHREC card page: 0.1% of decks. This is a close call. All That Glitters shares the count and is on the Enchantress page (26%).

Ranks 1 to 20 are the archetype's core, all on the Enchantress page at 30% to 69% of decks. Ranks 21 to 40 are weaker but real enchantment-cast, constellation, enchantment-count, and protection cards.

Signal bugs: Skybind gets tag:removal for a temporary exile that returns the permanent at end step. Ellivere of the Wild Court and Sky-Blessed Samurai get payoff-text:for each enchantment from reminder text. Shambling Suit gets payoff-text:enchantments you control from "artifacts and/or enchantments you control".

Legality: all 40 are legal in commander and inside GW.

Gaps (context only): Jukai Naturalist (58% of Enchantress decks), Hall of Heliod's Generosity (51%), Wild Growth (37%), Starfield Mystic (34%), Ondu Spiritdancer (29%), Enlightened Tutor (27%), Starfield of Nyx (25%), and Idyllic Tutor (22%) are absent.

## 16. counters proliferate (commander, GUB, any-card)

Theme signals: payoff tags counters-matter, pp-counters-matter, synergy-modified. tags counter-fuel. keywords Proliferate.

Funnel: 18160 legal in colors, 2386 on theme, 0 owned, 0 on theme and owned, 296 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Inspiring Call | interaction | 0 | 1.00 | payoff:counters-matter payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter tag:interaction |
| 2 | Walking Ballista | removal | 0 | 1.00 | payoff:counters-matter payoff:pp-counters-matter tag:counter-fuel text:+1/+1 counter tag:removal |
| 3 | Wishclaw Talisman | synergy | 0 | 1.00 | payoff:counters-matter tag:counter-fuel theme |
| 4 | Mossborn Hydra | synergy | 0 | 0.99 | payoff:counters-matter payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter theme |
| 5 | Gyre Sage | ramp | 0 | 0.99 | payoff:counters-matter payoff:pp-counters-matter payoff-text:for each +1/+1 counter text:+1/+1 counter tag:ramp |
| 6 | Evolution Witness | synergy | 0 | 0.99 | payoff:counters-matter payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters text:+1/+1 counter theme |
| 7 | Dragon's Hoard | ramp | 0 | 0.99 | payoff:counters-matter tag:counter-fuel tag:ramp |
| 8 | Tekuthal, Inquiry Dominus | threat | 0 | 0.99 | payoff:counters-matter tag:counter-fuel keyword:Proliferate type:threat |
| 9 | Simic Ascendancy | wincon | 0 | 0.99 | payoff:counters-matter payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters text:+1/+1 counter tag:alternate-win-condition |
| 10 | Ominous Seas | draw | 0 | 0.99 | payoff:counters-matter tag:counter-fuel tag:draw |
| 11 | Duskshell Crawler | synergy | 0 | 0.99 | payoff:counters-matter payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter theme |
| 12 | Hangarback Walker | synergy | 0 | 0.99 | payoff:counters-matter payoff:pp-counters-matter payoff-text:for each +1/+1 counter text:+1/+1 counter theme |
| 13 | Volatile Stormdrake | removal | 0 | 0.98 | payoff:counters-matter tag:counter-fuel tag:removal |
| 14 | Mycoloth | threat | 0 | 0.98 | payoff:counters-matter payoff:pp-counters-matter payoff-text:for each +1/+1 counter text:+1/+1 counter type:threat |
| 15 | Basking Broodscale | ramp | 0 | 0.98 | payoff:counters-matter payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters text:+1/+1 counter tag:ramp |
| 16 | Scurry Oak | synergy | 0 | 0.98 | payoff:counters-matter payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters text:+1/+1 counter theme |
| 17 | Quilled Greatwurm | threat | 0 | 0.98 | payoff:counters-matter tag:counter-fuel text:+1/+1 counter type:threat |
| 18 | Tome of Legends | draw | 0 | 0.98 | payoff:counters-matter tag:counter-fuel tag:draw |
| 19 | Chthonian Nightmare | synergy | 0 | 0.98 | payoff:counters-matter tag:counter-fuel theme |
| 20 | Llanowar Reborn | land | 0 | 0.98 | payoff:counters-matter payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter type:land |
| 21 | Steelbane Hydra | removal | 0 | 0.98 | payoff:counters-matter payoff:pp-counters-matter tag:counter-fuel text:+1/+1 counter tag:removal |
| 22 | Vat of Rebirth | synergy | 0 | 0.98 | payoff:counters-matter tag:counter-fuel theme |
| 23 | Khalni Heart Expedition | ramp | 0 | 0.98 | payoff:counters-matter tag:counter-fuel tag:ramp |
| 24 | Bred for the Hunt | draw | 0 | 0.98 | payoff:counters-matter payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter tag:draw |
| 25 | Golgari Grave-Troll | draw | 0 | 0.97 | payoff:counters-matter payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it tag:counter-fuel text:+1/+1 counter text:draw |
| 26 | Umezawa's Jitte | removal | 0 | 0.97 | payoff:counters-matter tag:counter-fuel tag:removal |
| 27 | Fertilid | ramp | 0 | 0.97 | payoff:counters-matter payoff:pp-counters-matter tag:counter-fuel text:+1/+1 counter tag:ramp |
| 28 | Crystalline Crawler | ramp | 0 | 0.97 | payoff:counters-matter payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it tag:counter-fuel text:+1/+1 counter tag:ramp |
| 29 | Benevolent Hydra | synergy | 0 | 0.97 | payoff:counters-matter payoff:pp-counters-matter tag:counter-fuel text:+1/+1 counter theme |
| 30 | Power Conduit | synergy | 0 | 0.97 | payoff:counters-matter tag:counter-fuel text:+1/+1 counter theme |
| 31 | Ingenious Prodigy | draw | 0 | 0.97 | payoff:counters-matter payoff:pp-counters-matter tag:counter-fuel text:+1/+1 counter tag:draw |
| 32 | Solar Transformer | ramp | 0 | 0.97 | payoff:counters-matter tag:counter-fuel tag:ramp |
| 33 | Fain, the Broker | ramp | 0 | 0.97 | payoff:counters-matter tag:counter-fuel text:+1/+1 counter tag:ramp |
| 34 | Iron Spider, Stark Upgrade | draw | 0 | 0.97 | payoff:counters-matter payoff:pp-counters-matter tag:counter-fuel text:+1/+1 counter tag:draw |
| 35 | Gemstone Mine | land | 0 | 0.97 | payoff:counters-matter tag:counter-fuel type:land |
| 36 | Hooded Hydra | synergy | 0 | 0.97 | payoff:counters-matter payoff:pp-counters-matter payoff-text:for each +1/+1 counter text:+1/+1 counter theme |
| 37 | Nyxborn Hydra | synergy | 0 | 0.97 | payoff:counters-matter payoff:pp-counters-matter payoff-text:for each +1/+1 counter text:+1/+1 counter theme |
| 38 | Aether Hub | land | 0 | 0.97 | payoff:counters-matter tag:counter-fuel type:land |
| 39 | O'aka, Traveling Merchant | draw | 0 | 0.97 | payoff:counters-matter tag:counter-fuel tag:draw |
| 40 | Aetherworks Marvel | synergy | 0 | 0.97 | payoff:counters-matter tag:counter-fuel theme |

### Review

Verdict: no. 25 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text, the EDHREC +1/+1 Counters, Proliferate, and Counters Matter pages (edhrec.com/tags/plus-1-plus-1-counters, /proliferate, /counters-matter), and EDHREC card pages for the borderline cards.

Off theme (15):

- 3 Wishclaw Talisman: a tutor with wish counters. Score 1.00. Its top commanders are cEDH tutor decks. Not on any counters page.
- 7 Dragon's Hoard: a Dragons mana rock with gold counters (58% of Dragon decks). Not on any counters page.
- 10 Ominous Seas: a draw-matters payoff with foreshadow counters. Not on any counters page.
- 13 Volatile Stormdrake: a theft creature whose own trigger spends the energy it grants. Not on any counters page.
- 18 Tome of Legends: a commander draw rock with page counters. Not on any counters page.
- 19 Chthonian Nightmare: sacrifice reanimation paid with energy. Not on any counters page.
- 22 Vat of Rebirth: oil-counter reanimation for aristocrats decks. Not on any counters page.
- 23 Khalni Heart Expedition: landfall ramp with quest counters. Not on any counters page.
- 25 Golgari Grave-Troll: a dredge card that carries +1/+1 counters. Not on any counters page.
- 26 Umezawa's Jitte: an Equipment staple with charge counters. Not on any counters page.
- 32 Solar Transformer: an energy mana rock. Not on any counters page.
- 34 Iron Spider, Stark Upgrade: puts +1/+1 counters on artifact creatures and Vehicles only. Not on any counters page. This is a close call.
- 35 Gemstone Mine: a color-fixing land with mining counters. Not on any counters page.
- 38 Aether Hub: an energy fixing land. Not on any counters page.
- 40 Aetherworks Marvel: an energy engine. Not on any counters page.

Borderline, counted on theme (1):

- 30 Power Conduit: turns a counter you remove into a +1/+1 counter on a creature every turn. Repeatable +1/+1 counter placement is theme text, so it counts on text alone.

Root cause: payoff:counters-matter and tag:counter-fuel fire on any card whose text names a counter type. Thirteen of the 15 misses come from this pair. They rank Wishclaw Talisman (1.00) and Dragon's Hoard (0.99) above every real counters card except Walking Ballista. Golgari Grave-Troll gets role draw from the dredge reminder text. Quilled Greatwurm places +1/+1 counters but gets no pp-counters-matter signal.

Legality: all 40 are legal in commander and inside GUB.

Gaps (context only): Karn's Bastion (63% of Proliferate decks), Evolution Sage (59%), Inexorable Tide (51%), Hardened Scales (47% of +1/+1 Counters decks), Branching Evolution (32%), The Ozolith (21%), Winding Constrictor (21%), and Doubling Season (20%) are absent. Tekuthal is the only proliferate card in the top 40.

## 17. zombies (commander, BU, any-card)

Theme signals: payoff tags typal-zombie. subtypes Zombie.

Funnel: 12600 legal in colors, 533 on theme, 0 owned, 0 on theme and owned, 292 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Death Baron | synergy | 0 | 0.98 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie theme |
| 2 | The Scarab God | removal | 0 | 0.98 | payoff:typal-zombie payoff-text:zombies you control text:destroy target |
| 3 | Undead Augur | draw | 0 | 0.98 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie tag:draw |
| 4 | Lord of the Accursed | synergy | 0 | 0.98 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie theme |
| 5 | Cryptbreaker | draw | 0 | 0.98 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie tag:draw |
| 6 | Diregraf Colossus | synergy | 0 | 0.98 | payoff:typal-zombie payoff-text:for each zombie subtype:Zombie theme |
| 7 | Diregraf Captain | synergy | 0 | 0.98 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie theme |
| 8 | Endless Ranks of the Dead | synergy | 0 | 0.98 | payoff:typal-zombie payoff-text:zombies you control theme |
| 9 | Wilhelt, the Rotcleaver | draw | 0 | 0.97 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie tag:draw |
| 10 | Champion of the Perished | synergy | 0 | 0.97 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie theme |
| 11 | Headless Rider | synergy | 0 | 0.97 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie theme |
| 12 | Necroduality | synergy | 0 | 0.97 | payoff:typal-zombie payoff-text:zombie you control theme |
| 13 | Lost Monarch of Ifnir | threat | 0 | 0.97 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie type:threat |
| 14 | Gisa, the Hellraiser | threat | 0 | 0.97 | payoff:typal-zombie payoff-text:zombies you control type:threat |
| 15 | Plague Belcher | synergy | 0 | 0.97 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie theme |
| 16 | Hordewing Skaab | draw | 0 | 0.96 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie tag:draw |
| 17 | Wand of Orcus | synergy | 0 | 0.96 | payoff:typal-zombie payoff-text:zombies you control theme |
| 18 | Liliana's Mastery | synergy | 0 | 0.96 | payoff:typal-zombie payoff-text:zombies you control theme |
| 19 | Geralf, Visionary Stitcher | synergy | 0 | 0.95 | payoff:typal-zombie payoff-text:zombies you control theme |
| 20 | Tomb Tyrant | threat | 0 | 0.95 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie type:threat |
| 21 | Prophet of the Scarab | draw | 0 | 0.95 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie tag:draw |
| 22 | Shepherd of Rot | synergy | 0 | 0.95 | payoff:typal-zombie payoff-text:for each zombie subtype:Zombie theme |
| 23 | Undead Alchemist | threat | 0 | 0.94 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie type:threat |
| 24 | Archghoul of Thraben | synergy | 0 | 0.94 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie theme |
| 25 | Gravecrawler | synergy | 0 | 0.94 | payoff:typal-zombie subtype:Zombie theme |
| 26 | Geralf, the Fleshwright | synergy | 0 | 0.94 | payoff:typal-zombie payoff-text:zombie you control theme |
| 27 | Liliana, Untouched by Death | removal | 0 | 0.93 | payoff:typal-zombie payoff-text:zombies you control tag:removal |
| 28 | Graveborn Muse | draw | 0 | 0.93 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie tag:draw |
| 29 | Dark Salvation | removal | 0 | 0.93 | payoff:typal-zombie payoff-text:for each zombie tag:removal |
| 30 | Gravespawn Sovereign | threat | 0 | 0.93 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie type:threat |
| 31 | Liliana, the Last Hope | removal | 0 | 0.93 | payoff:typal-zombie payoff-text:zombies you control tag:removal |
| 32 | Cursecloth Wrappings | synergy | 0 | 0.92 | payoff:typal-zombie payoff-text:zombies you control theme |
| 33 | Cemetery Reaper | removal | 0 | 0.92 | payoff:typal-zombie subtype:Zombie text:destroy target |
| 34 | Bladestitched Skaab | synergy | 0 | 0.92 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie theme |
| 35 | Unbreathing Horde | synergy | 0 | 0.92 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie theme |
| 36 | Zul Ashur, Lich Lord | synergy | 0 | 0.92 | payoff:typal-zombie subtype:Zombie theme |
| 37 | Zombie Master | interaction | 0 | 0.92 | payoff:typal-zombie subtype:Zombie tag:interaction |
| 38 | Gleaming Overseer | interaction | 0 | 0.92 | payoff:typal-zombie subtype:Zombie tag:interaction |
| 39 | Undead Warchief | threat | 0 | 0.92 | payoff:typal-zombie subtype:Zombie type:threat |
| 40 | Graf Harvest | synergy | 0 | 0.91 | payoff:typal-zombie payoff-text:zombies you control theme |

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Zombies page (edhrec.com/tags/zombies).

Off theme (1):

- 31 Liliana, the Last Hope: the +1 and -2 are generic removal and recursion. Only the -7 emblem names Zombies. Zombies page: 7% of decks, synergy 0.07. This is a close call.

Thirty-three of the 40 are on the Zombies page, most above 25% of decks. The seven absent cards are Zombies or Zombie-specific by oracle text: Shepherd of Rot, Undead Alchemist, Archghoul of Thraben, Graveborn Muse, Gravespawn Sovereign, Bladestitched Skaab, and Unbreathing Horde.

Signal bugs: The Scarab God and Cemetery Reaper get text:destroy target and role removal, but neither card contains "destroy". Zombie Master and Gleaming Overseer get tag:interaction for regenerate and hexproof grants.

Legality: all 40 are legal in commander and inside BU.

Gaps (context only): Lord of the Undead (35% of Zombies decks) and Liliana, Death's Majesty (30%) are absent.

## 18. burn (modern, R, any-card)

Theme signals: payoff tags synergy-burn. tags burn, burn-any, burn-player.

Funnel: 5166 legal in colors, 1292 on theme, 0 owned, 0 on theme and owned, 293 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Artist's Talent | draw | 0 | 0.99 | payoff:synergy-burn payoff-text:deals that much damage plus tag:draw |
| 2 | Chandra's Incinerator | removal | 0 | 0.98 | payoff:synergy-burn tag:burn tag:removal |
| 3 | Virtue of Courage // Embereth Blaze | removal | 0 | 0.96 | payoff:synergy-burn tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 4 | Toralf, God of Fury // Toralf's Hammer | removal | 0 | 0.96 | payoff:synergy-burn tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 5 | The Rollercrusher Ride | removal | 0 | 0.95 | payoff:synergy-burn tag:burn tag:removal |
| 6 | Fear of Burning Alive | removal | 0 | 0.94 | payoff:synergy-burn tag:burn-player tag:burn tag:removal |
| 7 | Aether Revolt | removal | 0 | 0.94 | payoff:synergy-burn payoff-text:deals that much damage plus tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 8 | Fall of Cair Andros | removal | 0 | 0.93 | payoff:synergy-burn tag:burn tag:removal |
| 9 | Magmatic Galleon | ramp | 0 | 0.93 | payoff:synergy-burn tag:burn tag:ramp |
| 10 | Imodane, the Pyrohammer | threat | 0 | 0.93 | payoff:synergy-burn tag:burn-player tag:burn type:threat |
| 11 | Hawkeye, Young Avenger | threat | 0 | 0.89 | payoff:synergy-burn payoff-text:deals that much damage plus type:threat |
| 12 | Jaya, Venerated Firemage | removal | 0 | 0.87 | payoff-text:deals that much damage plus tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 13 | Pyromancer's Swath | synergy | 0 | 0.87 | payoff:synergy-burn payoff-text:deals that much damage plus theme |
| 14 | Syr Carah, the Bold | removal | 0 | 0.87 | payoff:synergy-burn tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 15 | Invasion of Regatha // Disciples of the Inferno | removal | 0 | 0.86 | payoff:synergy-burn payoff-text:deals that much damage plus tag:burn-player tag:burn tag:removal |
| 16 | Pyromancer's Gauntlet | synergy | 0 | 0.85 | payoff:synergy-burn payoff-text:deals that much damage plus theme |
| 17 | Satyr Firedancer | removal | 0 | 0.85 | payoff:synergy-burn tag:burn tag:removal |
| 18 | Talon of Pain | removal | 0 | 0.80 | payoff:synergy-burn tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 19 | Solphim, Mayhem Dominus | threat | 0 | 0.71 | payoff:synergy-burn type:threat |
| 20 | Ojer Axonil, Deepest Might // Temple of Power | land | 0 | 0.71 | payoff:synergy-burn type:land |
| 21 | Lightning Bolt | removal | 0 | 0.69 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 22 | Aetherflux Reservoir | removal | 0 | 0.69 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 23 | Walking Ballista | removal | 0 | 0.69 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 24 | Goblin Bombardment | removal | 0 | 0.69 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 25 | Grapeshot | removal | 0 | 0.69 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 26 | Dragon Tempest | removal | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 27 | Drakuseth, Maw of Flames | removal | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 28 | Valakut, the Molten Pinnacle | land | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target type:land |
| 29 | Sword of Fire and Ice | interaction | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:interaction |
| 30 | Electro, Assaulting Battery | ramp | 0 | 0.68 | tag:burn-player tag:burn text:damage to target player tag:ramp |
| 31 | Electrodominance | removal | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 32 | Scourge of Valkas | removal | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 33 | Chandra, Torch of Defiance | ramp | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:ramp |
| 34 | Siege-Gang Commander | removal | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 35 | Ugin, the Spirit Dragon | wipe | 0 | 0.68 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:sweeper |
| 36 | Shock | removal | 0 | 0.67 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 37 | Gut Shot | removal | 0 | 0.67 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 38 | Outpost Siege | removal | 0 | 0.67 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 39 | Lightning Strike | removal | 0 | 0.67 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:removal |
| 40 | Leyline Tyrant | ramp | 0 | 0.67 | tag:burn-player tag:burn-any tag:burn text:damage to any target tag:ramp |

### Review

Verdict: no. 25 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and Modern legality, the mtgtop8 Modern metagame page, and one Modern Burn list (Iwao Sho, 2026-08-15, mtgtop8 deck 880322).

Off theme (15):

- 8 Fall of Cair Andros: excess damage amasses Orcs. An Army payoff.
- 9 Magmatic Galleon: a five-mana Vehicle. The damage hits a creature and makes a Treasure.
- 22 Aetherflux Reservoir: a lifegain combo, and rank 1 of prompt 1.
- 23 Walking Ballista: a colorless counter creature. Each counter after the first costs four mana.
- 24 Goblin Bombardment: a sacrifice outlet. It needs creatures to sacrifice.
- 26 Dragon Tempest: Dragon typal. The damage needs Dragons.
- 27 Drakuseth, Maw of Flames: a seven-mana attacker.
- 28 Valakut, the Molten Pinnacle: a ramp payoff land. It needs five other Mountains.
- 29 Sword of Fire and Ice: combat Equipment.
- 30 Electro, Assaulting Battery: a mana engine. The damage needs Electro to leave the battlefield.
- 32 Scourge of Valkas: Dragon typal.
- 34 Siege-Gang Commander: a Goblin token maker.
- 35 Ugin, the Spirit Dragon: an eight-mana colorless planeswalker.
- 38 Outpost Siege: a card-advantage enchantment. The Dragons mode needs a creature to leave the battlefield.
- 40 Leyline Tyrant: a mana-storage Dragon. The damage needs the Dragon to die.

Borderline (counted on theme): 4 Toralf, God of Fury, 17 Satyr Firedancer, 18 Talon of Pain, and 25 Grapeshot. Toralf and Satyr Firedancer send the damage to a creature. Talon of Pain needs charge counters first. Grapeshot needs a storm count. Count these four as off theme, and the total falls to 21.

Format fit: the list does not fit a 60-card Modern deck. Only 6 of the 40 are instants or sorceries. The average mana value is 3.5, and 11 cards cost 5 or more. The sampled Burn list plays 20 lands, and no spell in it costs more than 3.

Root cause: the three enabler tags fire on any "damage to any target" line. Ranks 21 to 40 share one theme score, and their EDHREC ranks rise in strict order from 160 to 2374. Popularity alone sorts the second half of the list. EDHREC rank measures Commander play, so a Modern prompt collects Commander staples. The hand-off records this limit. The numbers above measure it.

Signal bugs: the payoff needle "deals damage to each opponent" matches 11 red cards. Real cards write a number, and "deals N damage to each opponent" matches 176 red cards. Boltwave is a one-mana sorcery that deals 3 damage to each opponent. It gets no payoff signal and sits at position 193. The payoff tag synergy-burn adds 19 damage amplifiers, which is a Commander pattern.

Legality: all 40 are legal in Modern on 2026-08-24 and inside mono-red. Aetherflux Reservoir, Pyromancer's Gauntlet, Sword of Fire and Ice, Ugin, the Spirit Dragon, and Walking Ballista are colorless.

Owned: none. The prompt runs in any-card mode with no collection.

Gaps (context only): the sampled Burn list holds 10 nonland spells, and one of them reaches the top 40. That card is Lightning Bolt. The filter drops Boros Charm correctly, because it is outside mono-red. Monastery Swiftspear and Goblin Guide never enter the pool, because they have no damage text and no staple role. Skewer the Critics sits at position 101, Skullcrack at 114, Lava Spike at 153, Boltwave at 193, Rift Bolt at 225, and Searing Blaze at 337. Eidolon of the Great Revel sits at position 469.

## 19. control (standard, WU, any-card)

Theme signals: tags counterspell, draw-engine, removal, sweeper.

Funnel: 1821 legal in colors, 479 on theme, 0 owned, 0 on theme and owned, 162 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Getaway Glamer | interaction | 0 | 0.74 | tag:removal text:destroy target text:exile target tag:interaction |
| 2 | Stroke of Midnight | removal | 0 | 0.69 | tag:removal text:destroy target tag:removal |
| 3 | Demolition Field | land | 0 | 0.69 | tag:removal text:destroy target type:land |
| 4 | Elspeth, Storm Slayer | removal | 0 | 0.69 | tag:removal text:destroy target tag:removal |
| 5 | Meteor Golem | removal | 0 | 0.68 | tag:removal text:destroy target tag:removal |
| 6 | Three Steps Ahead | interaction | 0 | 0.68 | tag:counterspell text:counter target spell tag:interaction |
| 7 | Disenchant | removal | 0 | 0.68 | tag:removal text:destroy target tag:removal |
| 8 | Requisition Raid | removal | 0 | 0.68 | tag:removal text:destroy target tag:removal |
| 9 | Sheltered by Ghosts | interaction | 0 | 0.68 | tag:removal text:exile target tag:interaction |
| 10 | Cathar Commando | removal | 0 | 0.68 | tag:removal text:destroy target tag:removal |
| 11 | Long River's Pull | interaction | 0 | 0.68 | tag:counterspell text:counter target spell tag:interaction |
| 12 | Parting Gust | interaction | 0 | 0.68 | tag:removal text:exile target tag:interaction |
| 13 | Get Lost | removal | 0 | 0.68 | tag:removal text:destroy target tag:removal |
| 14 | Aven Interrupter | interaction | 0 | 0.67 | tag:counterspell text:exile target tag:interaction |
| 15 | Refute | interaction | 0 | 0.67 | tag:counterspell text:counter target spell tag:interaction |
| 16 | Cancel | interaction | 0 | 0.67 | tag:counterspell text:counter target spell tag:interaction |
| 17 | Erode | ramp | 0 | 0.67 | tag:removal text:destroy target tag:ramp |
| 18 | Banishing Light | removal | 0 | 0.67 | tag:removal text:exile target tag:removal |
| 19 | Disdainful Stroke | interaction | 0 | 0.67 | tag:counterspell text:counter target spell tag:interaction |
| 20 | Crib Swap | removal | 0 | 0.67 | tag:removal text:exile target tag:removal |
| 21 | Volatile Fault | land | 0 | 0.67 | tag:removal text:destroy target type:land |
| 22 | Valorous Stance | interaction | 0 | 0.67 | tag:removal text:destroy target tag:interaction |
| 23 | Venat, Heart of Hydaelyn // Hydaelyn, the Mothercrystal | interaction | 0 | 0.66 | tag:removal tag:draw-engine text:exile target tag:interaction |
| 24 | Syncopate | interaction | 0 | 0.66 | tag:counterspell text:counter target spell tag:interaction |
| 25 | It'll Quench Ya! | interaction | 0 | 0.66 | tag:counterspell text:counter target spell tag:interaction |
| 26 | Bovine Intervention | removal | 0 | 0.66 | tag:removal text:destroy target tag:removal |
| 27 | White Auracite | ramp | 0 | 0.66 | tag:removal text:exile target tag:ramp |
| 28 | Season of the Burrow | ramp | 0 | 0.65 | tag:removal text:exile target tag:ramp |
| 29 | Spell Stutter | interaction | 0 | 0.65 | tag:counterspell text:counter target spell tag:interaction |
| 30 | Dazzling Denial | interaction | 0 | 0.65 | tag:counterspell text:counter target spell tag:interaction |
| 31 | Dion, Bahamut's Dominant // Bahamut, Warden of Light | removal | 0 | 0.65 | tag:removal text:destroy target tag:removal |
| 32 | Unwanted Remake | removal | 0 | 0.65 | tag:removal text:destroy target tag:removal |
| 33 | Extinguisher Battleship | wipe | 0 | 0.65 | tag:removal tag:sweeper text:destroy target tag:sweeper |
| 34 | Mana Sculpt | ramp | 0 | 0.65 | tag:counterspell text:counter target spell tag:ramp |
| 35 | Airbender's Reversal | interaction | 0 | 0.65 | tag:removal text:destroy target tag:interaction |
| 36 | Amazing Acrobatics | interaction | 0 | 0.65 | tag:counterspell text:counter target spell tag:interaction |
| 37 | Zuko's Exile | removal | 0 | 0.65 | tag:removal text:exile target tag:removal |
| 38 | Ultima Weapon | removal | 0 | 0.64 | tag:removal text:destroy target tag:removal |
| 39 | Fear of Impostors | interaction | 0 | 0.64 | tag:counterspell text:counter target spell tag:interaction |
| 40 | No More Lies | interaction | 0 | 0.64 | tag:counterspell text:counter target spell tag:interaction |

### Review

Verdict: no. 34 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and legality, MTGGoldfish Standard Azorius, Esper, and Dimir Control and Azorius Tempo pages, mtgtop8 UW Control lists (RCQ 2026-08-15 and 2026-08-16, MTGO Challenge 2026-08-21), and magic.gg Traditional Standard decklists (2026-06-29). Standard has no rotation in 2026. Wilds of Eldraine through Duskmourn leave in February 2027 (Star City Games rotation guide, 2026-08-11).

Off theme (6):

- 5 Meteor Golem: a seven-mana 3/3 whose ETB destroys a nonland permanent. It scores 0.68, above No More Lies (0.64). Not in any sampled control list.
- 9 Sheltered by Ghosts: an Aura for your own creature. Its exile needs a creature to enchant. An aggro card. Not in any sampled control list.
- 23 Venat, Heart of Hydaelyn: a legendary-matters value creature. The exile costs {7} at sorcery speed. Not in any sampled control list.
- 31 Dion, Bahamut's Dominant: a go-wide Knight. The "destroy" is chapter III of its back-face Saga, ten mana and three turns deep. Not in any sampled control list.
- 33 Extinguisher Battleship: an eight-mana Spacecraft. Its ETB is a sweeper, but the cost and the station clause fit ramp decks. Seen only in Simic and Grixis ramp lists.
- 38 Ultima Weapon: a seven-mana Equipment that destroys a creature when the equipped creature attacks. A Voltron payoff. Not the sweeper Ultima that UW Control runs.

The 34 on-theme cards are 16 removal spells or removal permanents, 14 counterspells, 2 land-destruction utility lands, Elspeth, Storm Slayer, and Cathar Commando. Current UW Control staples present: No More Lies, Get Lost, Three Steps Ahead, Erode, Requisition Raid, Demolition Field, and Elspeth, Storm Slayer.

Signal bugs: Erode, Season of the Burrow, White Auracite, and Mana Sculpt get role ramp. Erode's basic land goes to the opponent. Season of the Burrow has no mana text. Dion's tag:removal comes from a back-face Saga chapter. Ultima Weapon's tag:removal is an attack trigger on an Equipment. Venat's tag:draw-engine is a once-per-turn legendary trigger. tag:removal fires on land destruction, which lifts Demolition Field to rank 3 above every counterspell.

Legality: all 40 are legal in Standard on 2026-08-24 and inside WU.

Gaps (context only): the card-advantage engines and sweepers are absent. Stock Up, Consult the Star Charts, Day of Judgment, Ultima, Spell Snare, Negate, Wan Shi Tong, Librarian, and Seam Rip appear in the sampled UW, Esper, and Dimir Control lists. The draw-engine tag found only Venat.

## 20. infect poison (commander, GB, any-card)

Theme signals: keywords Infect, Toxic, Proliferate.

Funnel: 12525 legal in colors, 149 on theme, 0 owned, 0 on theme and owned, 262 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Bloated Contaminator | synergy | 0 | 0.89 | payoff-text:poison counter keyword:Toxic keyword:Proliferate theme |
| 2 | Blightbelly Rat | synergy | 0 | 0.89 | payoff-text:poison counter keyword:Toxic keyword:Proliferate theme |
| 3 | Contaminant Grafter | ramp | 0 | 0.88 | payoff-text:poison counter payoff-text:corrupted keyword:Toxic keyword:Proliferate tag:ramp |
| 4 | Venomous Brutalizer | threat | 0 | 0.84 | payoff-text:poison counter keyword:Toxic keyword:Proliferate type:threat |
| 5 | Core Prowler | threat | 0 | 0.84 | payoff-text:poison counter keyword:Infect keyword:Proliferate type:threat |
| 6 | Vraska, Betrayal's Sting | ramp | 0 | 0.76 | payoff-text:poison counter keyword:Proliferate tag:ramp |
| 7 | Plague Myr | ramp | 0 | 0.76 | payoff-text:poison counter keyword:Infect tag:ramp |
| 8 | Myr Convert | ramp | 0 | 0.75 | payoff-text:poison counter keyword:Toxic tag:ramp |
| 9 | Skithiryx, the Blight Dragon | threat | 0 | 0.75 | payoff-text:poison counter keyword:Infect type:threat |
| 10 | Venerated Rotpriest | synergy | 0 | 0.75 | payoff-text:poison counter keyword:Toxic theme |
| 11 | Glistening Sphere | ramp | 0 | 0.75 | payoff-text:poison counter payoff-text:corrupted keyword:Proliferate tag:ramp |
| 12 | Tyrranax Rex | threat | 0 | 0.75 | payoff-text:poison counter keyword:Toxic type:threat |
| 13 | Ichor Rats | synergy | 0 | 0.74 | payoff-text:poison counter keyword:Infect theme |
| 14 | Karumonix, the Rat King | synergy | 0 | 0.74 | payoff-text:poison counter keyword:Toxic theme |
| 15 | Bloodroot Apothecary | synergy | 0 | 0.74 | payoff-text:poison counter keyword:Toxic theme |
| 16 | Bilious Skulldweller | synergy | 0 | 0.73 | payoff-text:poison counter keyword:Toxic theme |
| 17 | Ichorclaw Myr | synergy | 0 | 0.73 | payoff-text:poison counter keyword:Infect theme |
| 18 | Plague Stinger | synergy | 0 | 0.73 | payoff-text:poison counter keyword:Infect theme |
| 19 | Phyrexian Swarmlord | threat | 0 | 0.73 | payoff-text:poison counter keyword:Infect type:threat |
| 20 | Viridian Corrupter | removal | 0 | 0.72 | payoff-text:poison counter keyword:Infect tag:removal |
| 21 | Pestilent Syphoner | synergy | 0 | 0.72 | payoff-text:poison counter keyword:Toxic theme |
| 22 | Glistener Elf | synergy | 0 | 0.71 | payoff-text:poison counter keyword:Infect theme |
| 23 | Necrogen Rotpriest | threat | 0 | 0.71 | payoff-text:poison counter keyword:Toxic type:threat |
| 24 | Blight Mamba | synergy | 0 | 0.70 | payoff-text:poison counter keyword:Infect theme |
| 25 | Ichorspit Basilisk | synergy | 0 | 0.70 | payoff-text:poison counter keyword:Toxic theme |
| 26 | Phyrexian Crusader | synergy | 0 | 0.70 | payoff-text:poison counter keyword:Infect theme |
| 27 | Hand of the Praetors | threat | 0 | 0.69 | payoff-text:poison counter keyword:Infect type:threat |
| 28 | Necropede | removal | 0 | 0.69 | payoff-text:poison counter keyword:Infect tag:removal |
| 29 | Glissa's Retriever | threat | 0 | 0.69 | payoff-text:poison counter payoff-text:corrupted keyword:Toxic type:threat |
| 30 | Dune Mover | ramp | 0 | 0.68 | payoff-text:poison counter keyword:Toxic text:mana |
| 31 | Septic Rats | synergy | 0 | 0.68 | payoff-text:poison counter keyword:Infect theme |
| 32 | Paladin of Predation | threat | 0 | 0.67 | payoff-text:poison counter keyword:Toxic type:threat |
| 33 | Flesh-Eater Imp | threat | 0 | 0.67 | payoff-text:poison counter keyword:Infect type:threat |
| 34 | Tyrranax Atrocity | threat | 0 | 0.67 | payoff-text:poison counter keyword:Toxic type:threat |
| 35 | Whispering Specter | synergy | 0 | 0.67 | payoff-text:poison counter keyword:Infect theme |
| 36 | Corpse Cur | threat | 0 | 0.67 | payoff-text:poison counter keyword:Infect type:threat |
| 37 | Reaper of Sheoldred | threat | 0 | 0.66 | payoff-text:poison counter keyword:Infect type:threat |
| 38 | Phyrexian Hydra | threat | 0 | 0.66 | payoff-text:poison counter keyword:Infect type:threat |
| 39 | Branchblight Stalker | synergy | 0 | 0.65 | payoff-text:poison counter keyword:Toxic theme |
| 40 | Rot Wolf | draw | 0 | 0.65 | payoff-text:poison counter keyword:Infect tag:draw |

### Review

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Infect and Proliferate pages (edhrec.com/tags/infect, /proliferate).

Off theme: none. Every card has infect or toxic in its oracle text, or proliferates, or gives poison counters. Twenty-six of the 40 are on the Infect or Proliferate pages. The other 14 are infect or toxic creatures that EDHREC does not list.

Signal bugs: all 40 carry payoff-text:poison counter. On 28 of them the phrase appears only in the reminder text of Infect or Toxic. The signal does not separate real poison payoffs (Venerated Rotpriest, Phyrexian Swarmlord) from vanilla keyword creatures. Dune Mover gets text:mana and role ramp, but its text has no "mana". Vraska, Betrayal's Sting gets role ramp for a proliferate planeswalker.

Legality: all 40 are legal in commander and inside GB (some are colorless).

Gaps (context only): the proliferate support is absent. Infectious Inquiry (63% of Infect decks), Vraska's Fall (60%), Evolution Sage (57%), Karn's Bastion (56%), Infectious Bite (55%), Phyresis Outbreak (52%), Unnatural Restoration (48%), and Cankerbloom (46%). Their reminder text says "counter of each kind", not "poison counter".
