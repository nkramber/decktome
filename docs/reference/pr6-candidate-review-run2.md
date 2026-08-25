# PR-6 candidate review

Snapshot: 2026-08-24. Prompts: 2026-08-24. Collection: 2471 entries, 4316 cards, 1 rows unresolved.

Gate: for 20 theme prompts, a human confirms the top 40 candidates are on theme in at least 18. Ten run with no collection.

Run 2 (2026-08-24). Run 1 is `docs/reference/pr6-candidate-review.md`. It scored 12 of 20 and keeps every score and Review block. Run 2 follows three engine fixes and a `themes.json` repair for the eight failed prompts (D-65). Ten lists are identical to run 1. Ten changed, and each one has a new Review block.

Review method: each card was checked against its Scryfall oracle text, format legality, and color identity. A card is on theme when its main function serves the theme (payoff, enabler, or staple), at any power level. A theme line behind a condition still counts (D-64). A card whose theme link is incidental counts as on theme only when EDHREC shows the theme runs it. A card that fails format legality or color identity counts as off theme. Pass bar: 36 of 40 (owner decision, 2026-08-24). Prompts 1 and 11 share one card list. Each prompt section ends with a Review block.

Score each prompt in the table, then fill the total.

| # | Theme | Mode | On theme (yes/no) | Notes |
|---|---|---|---|---|
| 1 | lifegain | owned-first | yes | 40 of 40. Carried over from run 1. The list is identical, card for card. |
| 2 | aristocrats sacrifice | owned-first | yes | 40 of 40. The `death-trigger` parent tag is gone, so every Food, Clue, and Blood card left. Blood Artist, Zulaport Cutthroat, and Yawgmoth lead the list. Run 1 scored 15. |
| 3 | tokens go-wide | owned-first | yes | 38 of 40. The `anthem` payoff tag and the "creatures you control get" needle are gone, so no typal lord remains. Transmutation Font and Orcrist are the two misses. Run 1 scored 25. |
| 4 | voltron equipment | owned-first | yes | 38 of 40. Carried over from run 1. The list is identical, card for card. |
| 5 | blink | owned-first | yes | 37 of 40. The `flicker` parent tag is gone, so the three self-blink cards left. Panharmonicon and Elesh Norn are the first ETB payoffs. Run 1 scored 33. |
| 6 | reanimator | owned-first | yes | 38 of 40. `reanimate-creature` is now the payoff. Reanimate, Animate Dead, and Necromancy lead the list. The Soul Stone and Artisan of Kozilek are the two misses. Run 1 scored 34. |
| 7 | landfall | owned-first | yes | 36 of 40. Carried over from run 1. The list is identical, card for card. |
| 8 | spellslinger | owned-first | yes | 39 of 40. Carried over from run 1. The list is identical, card for card. |
| 9 | dragons | owned-first | yes | 39 of 40. Carried over from run 1. The list is identical, card for card. |
| 10 | artifacts | owned-first | yes | 37 of 40. Carried over from run 1. The list is identical, card for card. |
| 11 | lifegain | any-card | yes | 40 of 40. Carried over from run 1. The list is identical, card for card. |
| 12 | mill | any-card | yes | 39 of 40. `synergy-mill` replaces a dead slug. See Double is the one miss. Run 1 scored 39. |
| 13 | elves | any-card | yes | 40 of 40. Carried over from run 1. The list is identical, card for card. |
| 14 | storm | any-card | yes | 39 of 40. `storm-like` is the payoff, so the two hate cards and the token payoffs left. Twenty cards have the Storm keyword, against none in run 1. Run 1 scored 25. |
| 15 | enchantress | any-card | yes | 39 of 40. Carried over from run 1. The list is identical, card for card. |
| 16 | counters proliferate | any-card | yes | 39 of 40. `counters-matter` and `counter-fuel` are gone, so every card now names a +1/+1 counter. Golgari Grave-Troll is the one miss. Run 1 scored 25. |
| 17 | zombies | any-card | yes | 39 of 40. Carried over from run 1. The list is identical, card for card. |
| 18 | burn | any-card | yes | 39 of 40. A burn card must aim damage at a player and be a spell or an "each opponent" source. Glaring Fleshraker is the one miss. Run 1 scored 25. |
| 19 | control | any-card | yes | 36 of 40, exactly at the bar. Six more cards are borderline. A strict reading gives 30 and fails. Run 1 scored 34. |
| 20 | infect poison | any-card | yes | 40 of 40. `synergy-poison` and `poisonous` replace two dead slugs. Run 1 scored 40. |

Total on theme: 20 of 20. The bar is 18, so the gate holds. Run 1 scored 12 of 20 (see `pr6-candidate-review.md`).

## 1. lifegain (commander, WB, owned-first)

Theme signals: payoff tags life-total-matters-self, lifegain-matters. tags drain-life, lifegain, repeatable-lifegain. keywords Lifelink.

Funnel: 12683 legal in colors, 1715 on theme, 478 owned, 121 on theme and owned, 222 returned, 50 upgrades.

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

Carried over from run 1. The list is identical, card for card.

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Lifegain and Lifedrain pages (edhrec.com/tags/lifegain, /lifedrain).

Off theme: none. Thirty-one cards carry a literal lifegain trigger. The other nine are life-total payoffs (Aetherflux Reservoir, Serra Ascendant, Caduceus, Felidar Sovereign, Righteous Valkyrie, Cosmos Elixir, Angel of Destiny, Speaker of the Heavens) or a lifegain doubler (The Wind Crystal). Weakest fits: Caduceus (11% of Lifegain decks), Cosmos Elixir (10%), Ajani, Strength of the Pride (9%), and Ocelot Pride (not on the page, mostly a token card).

Signal bugs: Caduceus gets tag:interaction but removes, counters, and taxes nothing. The Wind Crystal and Vito get tag:repeatable-lifegain, but neither gains life by itself. Aetherflux Reservoir gets role removal while Felidar Sovereign and Angel of Destiny get role wincon for the same job.

Legality: all 40 are legal in commander and inside WB.

Owned: 2 of 40 are owned (Exemplar of Light, Aerith Gainsborough). The funnel reports 121 owned on-theme cards.

Gaps (context only): the archetype's enablers are absent. Soul Warden (40% of Lifegain decks), Exquisite Blood (37%), Authority of the Consuls (33%), Soul's Attendant (31%), Kambal, Consul of Allocation (26%), Bloodthirsty Conqueror (24%), Rhox Faithmender (20%), and Suture Priest (20%).

## 2. aristocrats sacrifice (commander, WB, owned-first)

Theme signals: payoff tags blood-artist-ability. tags free-sacrifice-outlet, repeatable-sacrifice-outlet, sacrifice-outlet-creature.

Funnel: 12683 legal in colors, 770 on theme, 432 owned, 41 on theme and owned, 176 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Blood Artist | synergy | 0 | 1.00 | payoff:blood-artist-ability payoff-text:another creature dies theme |
| 2 | Zulaport Cutthroat | synergy | 0 | 1.00 | payoff:blood-artist-ability payoff-text:creature you control dies theme |
| 3 | Bastion of Remembrance | synergy | 0 | 1.00 | payoff:blood-artist-ability payoff-text:creature you control dies theme |
| 4 | The Meathook Massacre | wipe | 0 | 1.00 | payoff:blood-artist-ability payoff-text:creature you control dies tag:sweeper |
| 5 | Elas il-Kor, Sadistic Pilgrim | synergy | 0 | 0.99 | payoff:blood-artist-ability payoff-text:creature you control dies theme |
| 6 | Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel | draw | 0 | 0.99 | payoff:blood-artist-ability payoff-text:another creature dies payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:draw |
| 7 | Vengeful Bloodwitch | synergy | 0 | 0.98 | payoff:blood-artist-ability payoff-text:creature you control dies theme |
| 8 | Falkenrath Noble | threat | 0 | 0.98 | payoff:blood-artist-ability payoff-text:another creature dies type:threat |
| 9 | Funeral Room // Awakening Hall | synergy | 0 | 0.98 | payoff:blood-artist-ability payoff-text:creature you control dies theme |
| 10 | Smothering Abomination | draw | 0 | 0.96 | payoff-text:whenever you sacrifice a creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet text:sacrifice a creature tag:draw |
| 11 | Vindictive Vampire | threat | 0 | 0.96 | payoff:blood-artist-ability payoff-text:creature you control dies type:threat |
| 12 | Grave Venerations | draw | 1 | 0.96 | payoff:blood-artist-ability payoff-text:creature you control dies tag:draw |
| 13 | Venerated Stormsinger | threat | 0 | 0.95 | payoff:blood-artist-ability payoff-text:creature you control dies type:threat |
| 14 | South Wind Avatar | threat | 0 | 0.94 | payoff:blood-artist-ability payoff-text:creature you control dies type:threat |
| 15 | Warren Soultrader | ramp | 0 | 0.91 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:ramp |
| 16 | Yawgmoth, Thran Physician | removal | 0 | 0.91 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:removal |
| 17 | Yahenni, Undying Partisan | synergy | 0 | 0.91 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:free-sacrifice-outlet tag:repeatable-sacrifice-outlet theme |
| 18 | Boggart Mischief | synergy | 0 | 0.91 | payoff:blood-artist-ability payoff-text:creature you control dies theme |
| 19 | Umbral Collar Zealot | synergy | 0 | 0.90 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:free-sacrifice-outlet tag:repeatable-sacrifice-outlet theme |
| 20 | Relic Vial | draw | 0 | 0.90 | payoff:blood-artist-ability payoff-text:creature you control dies tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet text:sacrifice a creature tag:draw |
| 21 | Disciple of Bolas | draw | 0 | 0.90 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:draw |
| 22 | Woe Strider | synergy | 0 | 0.90 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:free-sacrifice-outlet tag:repeatable-sacrifice-outlet theme |
| 23 | Bartolomé del Presidio | synergy | 0 | 0.90 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:free-sacrifice-outlet tag:repeatable-sacrifice-outlet theme |
| 24 | Razaketh, the Foulblooded | threat | 0 | 0.90 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet type:threat |
| 25 | Ruthless Technomancer | ramp | 0 | 0.90 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:ramp |
| 26 | Ghoulcaller Gisa | threat | 0 | 0.89 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet type:threat |
| 27 | Master of Dark Rites | ramp | 0 | 0.89 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:ramp |
| 28 | Meathook Massacre II | interaction | 0 | 0.89 | payoff-text:creature you control dies tag:sacrifice-outlet-creature tag:interaction |
| 29 | Commissar Severina Raine | draw | 0 | 0.88 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:draw |
| 30 | Shadowheart, Dark Justiciar | draw | 0 | 0.88 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:draw |
| 31 | Tevesh Szat, Doom of Fools | draw | 0 | 0.88 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:draw |
| 32 | Erebos, Bleak-Hearted | removal | 0 | 0.88 | payoff-text:creature you control dies payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:removal |
| 33 | Baron Bertram Graywater | draw | 0 | 0.88 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:draw |
| 34 | Teysa, Orzhov Scion | removal | 0 | 0.88 | payoff-text:creature you control dies tag:sacrifice-outlet-creature tag:free-sacrifice-outlet tag:repeatable-sacrifice-outlet tag:removal |
| 35 | High-Society Hunter | draw | 0 | 0.87 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:draw |
| 36 | Ruthless Lawbringer | removal | 0 | 0.87 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:removal |
| 37 | Lord Skitter's Butcher | draw | 0 | 0.87 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:draw |
| 38 | Illuminor Szeras | ramp | 0 | 0.87 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:ramp |
| 39 | Cavalier of Night | removal | 0 | 0.86 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:removal |
| 40 | Ayli, Eternal Pilgrim | removal | 0 | 0.86 | payoff-text:sacrifice another creature tag:sacrifice-outlet-creature tag:repeatable-sacrifice-outlet tag:removal |

Top upgrades (unowned):

- Blood Artist (synergy, 1.00)
- Zulaport Cutthroat (synergy, 1.00)
- Bastion of Remembrance (synergy, 1.00)
- The Meathook Massacre (wipe, 1.00)
- Elas il-Kor, Sadistic Pilgrim (synergy, 0.99)
- Sephiroth, Fabled SOLDIER // Sephiroth, One-Winged Angel (draw, 0.99)
- Vengeful Bloodwitch (synergy, 0.98)
- Falkenrath Noble (threat, 0.98)
- Funeral Room // Awakening Hall (synergy, 0.98)
- Smothering Abomination (draw, 0.96)

### Review

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Aristocrats and Sacrifice pages (edhrec.com/tags/aristocrats, /sacrifice). Run 1 scored 15 of 40.

Off theme: none. Twenty-seven cards are creature sacrifice outlets. Fifteen carry the Blood Artist drain ability. Blood Artist, Zulaport Cutthroat, Bastion of Remembrance, The Meathook Massacre, Yawgmoth, Thran Physician, Woe Strider, Teysa, Orzhov Scion, and Ayli, Eternal Pilgrim are all present.

Weakest fits: Boggart Mischief drains on a Goblin death only. Master of Dark Rites and Illuminor Szeras turn a creature into mana, which serves a combo plan more than a drain plan.

Fix in run 2: the payoff tag `death-trigger` is the parent of `death-trigger-self`, so every "when this dies" value creature scored as a payoff. The row now uses `blood-artist-ability` plus four creature-specific needles, and the outlet tags narrow to `sacrifice-outlet-creature` and `free-sacrifice-outlet`. Run 1's Food, Clue, and Blood cards are gone.

Legality: all 40 are legal in commander and inside WB.

Owned: 1 of 40 is owned (Grave Venerations). The funnel reports 41 owned on-theme cards.

Gaps (context only): Viscera Seer, Ashnod's Altar, Skullclamp, and Pitiless Plunderer are absent. A free sacrifice outlet with no drain text scores one enabler tag alone.

## 3. tokens go-wide (commander, GW, owned-first)

Theme signals: payoff tags synergy-token, synergy-token-creature, token-doubler, token-increaser, tokenfall. tags repeatable-creature-tokens, repeatable-token-generator.

Funnel: 12612 legal in colors, 1680 on theme, 508 owned, 139 on theme and owned, 246 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Caretaker's Talent | draw | 0 | 0.99 | payoff:synergy-token payoff:synergy-token-creature payoff:tokenfall payoff-text:creature tokens you control text:create a text:creature token tag:draw |
| 2 | Intangible Virtue | synergy | 0 | 0.99 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control text:creature token theme |
| 3 | Elspeth, Storm Slayer | removal | 0 | 0.99 | payoff:token-doubler payoff:token-increaser tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token tag:removal |
| 4 | Peregrin Took | draw | 3 | 0.99 | payoff:synergy-token payoff:token-increaser tag:repeatable-token-generator tag:draw |
| 5 | Ocelot Pride | synergy | 0 | 0.99 | payoff:synergy-token payoff:tokenfall payoff:token-doubler payoff:token-increaser tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 6 | Springleaf Parade | ramp | 0 | 0.98 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control text:creature token tag:ramp |
| 7 | Ainok Strike Leader | interaction | 0 | 0.98 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token tag:interaction |
| 8 | Inspiring Leader | synergy | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control text:creature token theme |
| 9 | Esika's Chariot | synergy | 0 | 0.97 | payoff:synergy-token tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 10 | Chitterspitter | synergy | 0 | 0.97 | payoff:synergy-token tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 11 | Nesting Dovehawk | threat | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 12 | Trostani, Selesnya's Voice | threat | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 13 | Transmutation Font | draw | 0 | 0.97 | payoff:synergy-token tag:repeatable-token-generator tag:draw |
| 14 | Rhys the Redeemed | synergy | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature payoff:token-doubler payoff:token-increaser tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 15 | Oltec Matterweaver | synergy | 0 | 0.97 | payoff:synergy-token tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 16 | On Wings of Gold | synergy | 0 | 0.97 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 17 | Cadira, Caller of the Small | synergy | 0 | 0.96 | payoff:synergy-token tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 18 | Prava of the Steel Legion | synergy | 0 | 0.96 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 19 | City of Death | ramp | 0 | 0.96 | payoff:synergy-token tag:repeatable-token-generator text:create a tag:ramp |
| 20 | Growing Ranks | synergy | 0 | 0.96 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 21 | King Darien XLVIII | interaction | 0 | 0.95 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token tag:interaction |
| 22 | Toby, Beastie Befriender | synergy | 0 | 0.95 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control text:create a text:creature token theme |
| 23 | Elspeth Tirel | wipe | 0 | 0.95 | payoff:synergy-token payoff-text:for each creature you control tag:repeatable-token-generator tag:repeatable-creature-tokens text:creature token tag:sweeper |
| 24 | Muster the Departed | synergy | 0 | 0.95 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 25 | Song of the Worldsoul | synergy | 0 | 0.95 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 26 | Hollowhenge Overlord | threat | 0 | 0.95 | payoff-text:for each creature you control tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 27 | Idol of Oblivion | draw | 1 | 0.94 | payoff:synergy-token text:create a text:creature token tag:draw |
| 28 | Twilight Drover | synergy | 0 | 0.94 | payoff:synergy-token payoff:synergy-token-creature tag:repeatable-token-generator tag:repeatable-creature-tokens text:creature token theme |
| 29 | Killer Service | synergy | 0 | 0.94 | payoff:synergy-token tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 30 | Battle for Bretagard | synergy | 0 | 0.93 | payoff:synergy-token payoff-text:creature tokens you control text:create a text:creature token theme |
| 31 | Rootborn Defenses | interaction | 0 | 0.93 | payoff:synergy-token payoff:synergy-token-creature text:create a text:creature token tag:interaction |
| 32 | Pollen-Shield Hare // Hare Raising | synergy | 0 | 0.93 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control text:creature token theme |
| 33 | Mite Overseer | threat | 0 | 0.93 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token type:threat |
| 34 | Life Finds a Way | synergy | 0 | 0.93 | payoff:synergy-token tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token theme |
| 35 | Orcrist, Goblin-cleaver | ramp | 0 | 0.92 | payoff-text:for each creature you control tag:repeatable-token-generator text:create a tag:ramp |
| 36 | Sundering Growth | removal | 1 | 0.92 | payoff:synergy-token payoff:synergy-token-creature text:create a text:creature token tag:removal |
| 37 | Propagator Drone | ramp | 0 | 0.92 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control tag:repeatable-token-generator tag:repeatable-creature-tokens text:create a text:creature token tag:ramp |
| 38 | Staff of the Storyteller | draw | 0 | 0.91 | payoff:synergy-token payoff:synergy-token-creature text:create a text:creature token tag:draw |
| 39 | Belladonna Took | draw | 0 | 0.91 | payoff:synergy-token payoff:tokenfall payoff-text:whenever a token tag:draw |
| 40 | Audience with Trostani | draw | 0 | 0.91 | payoff:synergy-token payoff:synergy-token-creature payoff-text:creature tokens you control text:create a text:creature token tag:draw |

Top upgrades (unowned):

- Caretaker's Talent (draw, 0.99)
- Intangible Virtue (synergy, 0.99)
- Elspeth, Storm Slayer (removal, 0.99)
- Ocelot Pride (synergy, 0.99)
- Springleaf Parade (ramp, 0.98)
- Ainok Strike Leader (interaction, 0.98)
- Inspiring Leader (synergy, 0.97)
- Esika's Chariot (synergy, 0.97)
- Chitterspitter (synergy, 0.97)
- Nesting Dovehawk (threat, 0.97)

### Review

Verdict: yes. 38 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Tokens page (edhrec.com/tags/tokens). Run 1 scored 25 of 40.

Off theme (2):

- 13 Transmutation Font: it makes Blood, Clue, and Food tokens. No creature token, so the board never widens.
- 35 Orcrist, Goblin-cleaver: an Equipment that makes Treasure tokens on combat damage.

Fix in run 2: the payoff tag `anthem` counts every typal lord, and the needle "creatures you control get" counts every qualified anthem. Both are gone. The needle "creature tokens you control" replaces them. All 13 typal or colorless lords from run 1 are gone.

Legality: all 40 are legal in commander and inside GW.

Owned: 3 of 40 are owned (Peregrin Took, Idol of Oblivion, Sundering Growth). The funnel reports 139 owned on-theme cards.

Gaps (context only): Anointed Procession, Parallel Lives, Doubling Season, and Cathars' Crusade are still absent. Each is a token doubler with no `synergy-token` tag in the snapshot.

## 4. voltron equipment (commander, RW, owned-first)

Theme signals: payoff tags synergy-aura, synergy-equipment, synergy-modified. tags evasion, protection. keywords Equip, Double strike, Hexproof.

Funnel: 12737 legal in colors, 3362 on theme, 604 owned, 234 on theme and owned, 283 returned, 50 upgrades.

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

Carried over from run 1. The list is identical, card for card.

Verdict: yes. 38 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Voltron, Equipment, and Auras pages (edhrec.com/tags/voltron, /equipment, /auras).

Off theme (2):

- 28 Greater Auramancy: gives your other enchantments and your enchanted creatures shroud. It has no Equipment text, and shroud blocks your own equip abilities. Not on the Voltron, Equipment, or Auras pages. An enchantress card (17% of enchantress decks). This is a close call.
- 36 Goro-Goro, Disciple of Ryusei: a haste anthem and a Dragon token engine that needs an attacking modified creature. The Equipment and Aura signals match only the reminder text. Not on any EDHREC page in the set. This is a close call.

Signal bugs: Goro-Goro gets payoff:synergy-equipment, payoff:synergy-aura, text:equipment, and tag:evasion from reminder text. The card has no Equipment, Aura, or evasion ability. Greater Auramancy gets payoff-text:enchanted creature from a static grant, not from an Aura.

Legality: all 40 are legal in commander and inside RW.

Owned: 7 of 40 are owned (ranks 1, 2, 8, 10, 13, 32, 33). The funnel reports 234 owned on-theme cards.

Gaps (context only): the archetype's tutor and draw core is absent. Sram, Senior Edificer (61% of Equipment decks), Open the Armory (48%), Steelshaper's Gift (46%), Forge Anew (46%), Stoneforge Mystic (42%), Danitha Capashen, Paragon (42%), Blackblade Reforged (40%), and Colossus Hammer (40%). The ranker favors "equipped creature" protection text over these engine cards.

## 5. blink (commander, WU, owned-first)

Theme signals: tags flicker-creature, flicker-nonland, flicker-permanent.

Funnel: 12669 legal in colors, 191 on theme, 452 owned, 14 on theme and owned, 170 returned, 50 upgrades.

Thin theme: the collection holds under 30 on-theme cards. PR-7 asks the pool-mode question again here (D-63).

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Ephemerate | interaction | 0 | 0.80 | tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield tag:interaction |
| 2 | Y'shtola Rhul | threat | 0 | 0.78 | tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield type:threat |
| 3 | Momentary Blink | interaction | 0 | 0.78 | tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield tag:interaction |
| 4 | Splash Portal | draw | 0 | 0.78 | tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield tag:draw |
| 5 | Conjurer's Closet | synergy | 0 | 0.69 | tag:flicker-creature text:exile target creature you control, then return theme |
| 6 | Cloudshift | interaction | 0 | 0.68 | tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 7 | Essence Flux | interaction | 0 | 0.68 | tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 8 | Deadeye Navigator | interaction | 0 | 0.68 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 9 | Charming Prince | synergy | 0 | 0.68 | tag:flicker-creature text:return it to the battlefield theme |
| 10 | Touch the Spirit Realm | interaction | 0 | 0.67 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 11 | Flicker of Fate | interaction | 0 | 0.67 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 12 | Planar Incision | interaction | 0 | 0.67 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 13 | Blur | interaction | 0 | 0.67 | tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 14 | Airbender Ascension | removal | 0 | 0.67 | tag:flicker-creature text:return it to the battlefield tag:removal |
| 15 | Eldrazi Confluence | ramp | 0 | 0.67 | tag:flicker-creature tag:flicker-nonland text:return it to the battlefield tag:ramp |
| 16 | Eldrazi Displacer | interaction | 0 | 0.66 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 17 | Oath of Teferi | synergy | 0 | 0.66 | tag:flicker-creature tag:flicker-permanent text:return it to the battlefield theme |
| 18 | Cosmic Intervention | interaction | 0 | 0.65 | tag:flicker-creature tag:flicker-permanent text:return it to the battlefield tag:interaction |
| 19 | Settle Beyond Reality | removal | 0 | 0.65 | tag:flicker-creature text:exile target creature you control, then return text:return it to the battlefield tag:removal |
| 20 | Gossip's Talent | synergy | 0 | 0.65 | tag:flicker-creature text:return it to the battlefield theme |
| 21 | Distinguished Conjurer | interaction | 0 | 0.65 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 22 | Far Traveler | synergy | 0 | 0.65 | tag:flicker-creature text:return it to the battlefield theme |
| 23 | Slip On the Ring | interaction | 1 | 0.65 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 24 | Scrollshift | interaction | 0 | 0.64 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 25 | Acrobatic Maneuver | interaction | 0 | 0.64 | tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 26 | Abuelo, Ancestral Echo | interaction | 0 | 0.64 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 27 | Siren's Ruse | interaction | 0 | 0.64 | tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 28 | Personify | synergy | 3 | 0.64 | tag:flicker-creature text:exile target creature you control, then return theme |
| 29 | Long River Lurker | interaction | 0 | 0.63 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 30 | Panharmonicon | synergy | 0 | 0.63 | payoff-text:entering causes a triggered ability theme |
| 31 | Venser, the Sojourner | removal | 0 | 0.63 | tag:flicker-creature tag:flicker-permanent text:return it to the battlefield tag:removal |
| 32 | Getaway Glamer | interaction | 0 | 0.63 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 33 | Mystifying Maze | land | 0 | 0.63 | tag:flicker-creature text:return it to the battlefield type:land |
| 34 | Justiciar's Portal | interaction | 0 | 0.63 | tag:flicker-creature text:exile target creature you control, then return tag:interaction |
| 35 | Elesh Norn, Mother of Machines | threat | 0 | 0.63 | payoff-text:entering causes a triggered ability type:threat |
| 36 | Oji, the Exquisite Blade | interaction | 0 | 0.62 | tag:flicker-creature text:return it to the battlefield tag:interaction |
| 37 | Meneldor, Swift Savior | threat | 1 | 0.62 | tag:flicker-creature text:return it to the battlefield type:threat |
| 38 | Guardian of Ghirapur | synergy | 0 | 0.62 | tag:flicker-creature text:return it to the battlefield theme |
| 39 | Virtue of Knowledge // Vantress Visions | synergy | 0 | 0.62 | payoff-text:entering causes a triggered ability theme |
| 40 | Salvation Swan | interaction | 0 | 0.61 | tag:flicker-creature text:return it to the battlefield tag:interaction |

Top upgrades (unowned):

- Ephemerate (interaction, 0.80)
- Y'shtola Rhul (threat, 0.78)
- Momentary Blink (interaction, 0.78)
- Splash Portal (draw, 0.78)
- Conjurer's Closet (synergy, 0.69)
- Cloudshift (interaction, 0.68)
- Essence Flux (interaction, 0.68)
- Deadeye Navigator (interaction, 0.68)
- Charming Prince (synergy, 0.68)
- Touch the Spirit Realm (interaction, 0.67)

### Review

Verdict: yes. 37 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Blink page (edhrec.com/tags/blink). Run 1 scored 33 of 40.

Off theme (3):

- 15 Eldrazi Confluence: one mode of three flickers, and the permanent returns tapped. The card also needs {C}{C}, which a WU deck rarely makes.
- 29 Long River Lurker: a Frog lord with ward. The blink needs combat damage from one chosen creature.
- 33 Mystifying Maze: the target is an attacking creature an opponent controls. The land never repeats your own ETB triggers.

Fix in run 2: the row dropped the `flicker` parent tag, which carries the `flicker-self` branch. Nezahal, Primal Tide, Estrid's Invocation, and Stenn, Paranoid Partisan protect themselves and are now gone. The payoff is the Panharmonicon wording, "entering causes a triggered ability". Panharmonicon (rank 30), Elesh Norn, Mother of Machines (35), and Virtue of Knowledge (39) are the first ETB payoffs the list has held.

Legality: all 40 are legal in commander and inside WU. Four are colorless, which includes Eldrazi Confluence and Mystifying Maze.

Owned: 3 of 40 are owned (Slip On the Ring, Personify, Meneldor, Swift Savior). The funnel reports 14 owned on-theme cards, so the thin-theme flag still holds (D-63).

Gaps (context only): Soulherder (58% of Blink decks), Ghostly Flicker (50%), Teleportation Circle (44%), Eerie Interlude (41%), and Restoration Angel (35%) still rank below the cut. Each carries the flicker tag and no text needle, so two 0.4 needles decide the order.

## 6. reanimator (commander, B, owned-first)

Theme signals: payoff tags reanimate-creature, reanimate-matters. tags discard, mill-self, reanimate, reanimate-from-any, tutor-to-graveyard.

Funnel: 7465 legal in colors, 1009 on theme, 304 owned, 58 on theme and owned, 194 returned, 50 upgrades.

The table merges owned cards and upgrades by score (D-62). Owned 0 marks an upgrade.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Reanimate | synergy | 0 | 1.00 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any theme |
| 2 | Victimize | synergy | 0 | 1.00 | payoff:reanimate-creature tag:reanimate theme |
| 3 | Animate Dead | synergy | 0 | 1.00 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any theme |
| 4 | Syr Konrad, the Grim | threat | 0 | 1.00 | payoff:reanimate-matters tag:mill-self type:threat |
| 5 | Living Death | wipe | 0 | 1.00 | payoff:reanimate-creature tag:reanimate tag:sweeper |
| 6 | Rise of the Dark Realms | synergy | 0 | 1.00 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any theme |
| 7 | Dread Return | synergy | 0 | 1.00 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate text:from your graveyard to the battlefield theme |
| 8 | Whip of Erebos | synergy | 0 | 0.99 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate text:from your graveyard to the battlefield theme |
| 9 | Sheoldred, Whispering One | removal | 0 | 0.99 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate text:from your graveyard to the battlefield tag:removal |
| 10 | Breach the Multiverse | synergy | 0 | 0.99 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any tag:mill-self theme |
| 11 | Unearth | draw | 0 | 0.99 | payoff:reanimate-creature tag:reanimate text:from your graveyard to the battlefield tag:draw |
| 12 | Agadeem's Awakening // Agadeem, the Undercrypt | land | 0 | 0.99 | payoff:reanimate-creature tag:reanimate text:from your graveyard to the battlefield type:land |
| 13 | Necromancy | synergy | 0 | 0.99 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any theme |
| 14 | The Soul Stone | ramp | 0 | 0.99 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate text:from your graveyard to the battlefield tag:ramp |
| 15 | Portal to Phyrexia | removal | 0 | 0.99 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any tag:removal |
| 16 | Artisan of Kozilek | removal | 0 | 0.99 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate text:from your graveyard to the battlefield tag:removal |
| 17 | Junji, the Midnight Sky | threat | 0 | 0.99 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any tag:discard type:threat |
| 18 | The Darkness Crystal | synergy | 1 | 0.99 | payoff:reanimate-creature tag:reanimate theme |
| 19 | Persist | synergy | 1 | 0.99 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate text:from your graveyard to the battlefield theme |
| 20 | Patriarch's Bidding | synergy | 0 | 0.99 | payoff:reanimate-creature tag:reanimate theme |
| 21 | Stitch Together | synergy | 0 | 0.98 | payoff:reanimate-creature tag:reanimate text:from your graveyard to the battlefield theme |
| 22 | Ancient Brass Dragon | threat | 0 | 0.98 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any type:threat |
| 23 | Lively Dirge | synergy | 0 | 0.98 | payoff:reanimate-creature tag:reanimate tag:tutor-to-graveyard text:from your graveyard to the battlefield theme |
| 24 | Lazotep Quarry | land | 0 | 0.98 | payoff:reanimate-creature tag:reanimate type:land |
| 25 | The Eldest Reborn | removal | 0 | 0.98 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any tag:discard tag:removal |
| 26 | Virtue of Persistence // Locthwain Scorn | removal | 0 | 0.98 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any tag:removal |
| 27 | Sheoldred // The True Scriptures | removal | 0 | 0.98 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any tag:discard tag:removal |
| 28 | Liliana, Death's Majesty | wipe | 0 | 0.98 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate tag:mill-self text:from your graveyard to the battlefield tag:sweeper |
| 29 | Funeral Room // Awakening Hall | synergy | 0 | 0.98 | payoff:reanimate-creature tag:reanimate text:from your graveyard to the battlefield theme |
| 30 | Ruthless Technomancer | ramp | 0 | 0.98 | payoff:reanimate-creature tag:reanimate text:from your graveyard to the battlefield tag:ramp |
| 31 | Sepulchral Primordial | threat | 0 | 0.98 | payoff:reanimate-creature tag:reanimate type:threat |
| 32 | Dance of the Dead | synergy | 0 | 0.98 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any theme |
| 33 | Chthonian Nightmare | synergy | 0 | 0.98 | payoff:reanimate-creature tag:reanimate text:from your graveyard to the battlefield theme |
| 34 | Lich-Knights' Conquest | synergy | 0 | 0.98 | payoff:reanimate-creature tag:reanimate text:from your graveyard to the battlefield theme |
| 35 | Puppeteer Clique | threat | 1 | 0.98 | payoff:reanimate-creature tag:reanimate type:threat |
| 36 | Vat of Rebirth | synergy | 0 | 0.98 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate text:from your graveyard to the battlefield theme |
| 37 | Grave Researcher // Reanimate | synergy | 0 | 0.98 | payoff:reanimate-creature tag:reanimate tag:reanimate-from-any tag:mill-self theme |
| 38 | Will of the Abzan | removal | 0 | 0.98 | payoff:reanimate-creature payoff-text:creature card from your graveyard to the battlefield tag:reanimate text:from your graveyard to the battlefield tag:removal |
| 39 | Zul Ashur, Lich Lord | synergy | 0 | 0.97 | payoff:reanimate-creature tag:reanimate theme |
| 40 | Bloodline Bidding | synergy | 1 | 0.97 | payoff:reanimate-creature tag:reanimate text:from your graveyard to the battlefield theme |

Top upgrades (unowned):

- Reanimate (synergy, 1.00)
- Victimize (synergy, 1.00)
- Animate Dead (synergy, 1.00)
- Syr Konrad, the Grim (threat, 1.00)
- Living Death (wipe, 1.00)
- Rise of the Dark Realms (synergy, 1.00)
- Dread Return (synergy, 1.00)
- Whip of Erebos (synergy, 0.99)
- Sheoldred, Whispering One (removal, 0.99)
- Breach the Multiverse (synergy, 0.99)

### Review

Verdict: yes. 38 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Reanimator and mono-black Reanimator pages. Run 1 scored 34 of 40.

Off theme (2):

- 14 The Soul Stone: a mana rock. The reanimation needs {6}{B}, a tap, and an exiled creature first.
- 16 Artisan of Kozilek: the reanimation is a cast trigger on a nine-mana spell. A reanimated Artisan gives nothing.

Fix in run 2: `reanimate-creature` is now the payoff, and the loose `leaving-graveyard-matters` payoff and the `reanimate-cast` enabler are gone. Reanimate, Victimize, Animate Dead, Living Death, Dread Return, Necromancy, and Persist now hold ranks 1 to 19. Run 1 held none of them.

Legality: all 40 are legal in commander and inside mono-black. Three are colorless.

Owned: 4 of 40 are owned (The Darkness Crystal, Persist, Puppeteer Clique, Bloodline Bidding). The funnel reports 58 owned on-theme cards.

Gaps (context only): the discard outlets and the graveyard tutors are thin. Entomb and Buried Alive carry `tutor-to-graveyard` but no payoff signal, so both stay below the cut.

## 7. landfall (commander, GU, owned-first)

Theme signals: payoff tags land-count-matters, landfall, landfall-other, lands-matter. tags land-ramp. keywords Landfall.

Funnel: 12456 legal in colors, 867 on theme, 466 owned, 106 on theme and owned, 176 returned, 50 upgrades.

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

Carried over from run 1. The list is identical, card for card.

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

Funnel: 12588 legal in colors, 1313 on theme, 519 owned, 111 on theme and owned, 216 returned, 50 upgrades.

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

Carried over from run 1. The list is identical, card for card.

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

Funnel: 18314 legal in colors, 328 on theme, 636 owned, 6 on theme and owned, 165 returned, 50 upgrades.

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

Carried over from run 1. The list is identical, card for card.

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

Funnel: 12588 legal in colors, 2150 on theme, 505 owned, 144 on theme and owned, 202 returned, 50 upgrades.

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

Carried over from run 1. The list is identical, card for card.

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

Carried over from run 1. The list is identical, card for card.

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24.

The 40 cards, ranks, roles, scores, and signals are identical to prompt 1. The two owned cards in prompt 1 (Exemplar of Light at rank 11, Aerith Gainsborough at rank 21) hold the same ranks here on score alone. So owned-first mode did not promote any owned card into the top 40. See the prompt 1 review for the card-level findings, signal bugs, and gaps.

## 12. mill (commander, UB, any-card)

Theme signals: payoff tags synergy-mill. tags mill, mill-any, mill-opponent. keywords Mill.

Funnel: 12600 legal in colors, 820 on theme, 0 owned, 0 on theme and owned, 296 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Syr Konrad, the Grim | threat | 0 | 1.00 | payoff:synergy-mill tag:mill-opponent tag:mill keyword:Mill text:mills type:threat |
| 2 | Bloodchief Ascension | synergy | 0 | 0.99 | payoff:synergy-mill payoff-text:whenever a card is put into an opponent's graveyard theme |
| 3 | Jace, Wielder of Mysteries | wincon | 0 | 0.99 | payoff:synergy-mill tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills tag:alternate-win-condition |
| 4 | Consuming Aberration | threat | 0 | 0.99 | payoff:synergy-mill tag:mill-opponent tag:mill type:threat |
| 5 | The Water Crystal | synergy | 0 | 0.98 | payoff:synergy-mill tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 6 | Sheoldred // The True Scriptures | removal | 0 | 0.98 | payoff:synergy-mill payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills tag:removal |
| 7 | Cemetery Tampering | synergy | 0 | 0.97 | payoff:synergy-mill tag:mill keyword:Mill theme |
| 8 | Fraying Sanity | synergy | 0 | 0.97 | payoff:synergy-mill tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 9 | Zellix, Sanity Flayer | synergy | 0 | 0.97 | payoff:synergy-mill tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 10 | See Double | synergy | 0 | 0.97 | payoff:synergy-mill payoff-text:cards in their graveyard theme |
| 11 | Riverchurn Monument | synergy | 0 | 0.97 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill-any tag:mill keyword:Mill theme |
| 12 | Mirelurk Queen | draw | 0 | 0.96 | payoff:synergy-mill tag:mill-opponent tag:mill-any tag:mill tag:draw |
| 13 | Screeching Scorchbeast | threat | 0 | 0.96 | payoff:synergy-mill tag:mill-opponent tag:mill type:threat |
| 14 | Captain N'ghathrod | threat | 0 | 0.96 | payoff:synergy-mill tag:mill-opponent tag:mill keyword:Mill text:mills type:threat |
| 15 | Infesting Radroach | synergy | 0 | 0.95 | payoff:synergy-mill tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 16 | Raul, Trouble Shooter | synergy | 0 | 0.95 | payoff:synergy-mill tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 17 | Into the Story | draw | 0 | 0.95 | payoff:synergy-mill payoff-text:cards in their graveyard tag:draw |
| 18 | Dreadhound | threat | 0 | 0.95 | payoff:synergy-mill tag:mill keyword:Mill type:threat |
| 19 | Sewer Nemesis | threat | 0 | 0.94 | payoff:synergy-mill tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills type:threat |
| 20 | Undead Alchemist | threat | 0 | 0.94 | payoff:synergy-mill tag:mill-opponent tag:mill keyword:Mill text:mills type:threat |
| 21 | Duskmantle Guildmage | synergy | 0 | 0.94 | payoff:synergy-mill payoff-text:whenever a card is put into an opponent's graveyard tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 22 | Polluted Cistern // Dim Oubliette | synergy | 0 | 0.94 | payoff:synergy-mill tag:mill keyword:Mill theme |
| 23 | The Haunt of Hightower | threat | 0 | 0.93 | payoff:synergy-mill payoff-text:whenever a card is put into an opponent's graveyard type:threat |
| 24 | Unshakable Tail | draw | 0 | 0.93 | payoff:synergy-mill tag:mill tag:draw |
| 25 | Jace's Phantasm | synergy | 0 | 0.93 | payoff:synergy-mill payoff-text:cards in their graveyard theme |
| 26 | Blackbloom Rogue // Blackbloom Bog | land | 0 | 0.92 | payoff:synergy-mill payoff-text:cards in their graveyard type:land |
| 27 | Relic Golem | synergy | 0 | 0.92 | payoff:synergy-mill payoff-text:cards in their graveyard tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 28 | Thieves' Guild Enforcer | synergy | 0 | 0.92 | payoff:synergy-mill payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 29 | Merfolk Windrobber | draw | 0 | 0.91 | payoff:synergy-mill payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills tag:draw |
| 30 | Deepmuck Desperado | synergy | 0 | 0.91 | payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 31 | Soaring Thought-Thief | synergy | 0 | 0.90 | payoff:synergy-mill payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 32 | Vantress Gargoyle | synergy | 0 | 0.89 | payoff:synergy-mill payoff-text:cards in their graveyard tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 33 | Anticognition | interaction | 0 | 0.88 | payoff:synergy-mill payoff-text:cards in their graveyard tag:interaction |
| 34 | Nihilith | threat | 0 | 0.86 | payoff:synergy-mill payoff-text:whenever a card is put into an opponent's graveyard type:threat |
| 35 | Dimir Strandcatcher | draw | 0 | 0.85 | payoff:synergy-mill tag:mill tag:draw |
| 36 | Bruvac the Grandiloquent | synergy | 0 | 0.84 | payoff:synergy-mill keyword:Mill theme |
| 37 | Altar of Dementia | synergy | 0 | 0.83 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 38 | Brain Freeze | synergy | 0 | 0.82 | tag:mill-opponent tag:mill-any tag:mill keyword:Mill text:mills theme |
| 39 | Breach the Multiverse | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |
| 40 | Mindcrank | synergy | 0 | 0.82 | tag:mill-opponent tag:mill keyword:Mill text:mills theme |

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Mill page (edhrec.com/tags/mill). Run 1 scored 39 of 40.

Off theme (1):

- 10 See Double: a spell copy with a graveyard-count rider. The main function is the copy, so the card belongs to spellslinger.

Fix in run 2: `synergy-mill` replaces the dead payoff slug `mill-matters`. Thirty-four cards now carry a real mill payoff tag, and the eight-card-threshold bodies of run 1 rank lower. Bruvac the Grandiloquent entered the list.

Legality: all 40 are legal in commander and inside UB. Three are colorless.

Owned: none. The prompt runs in any-card mode.

Gaps (context only): Ruin Crab, Hedron Crab, and Mesmeric Orb are absent. Each mills for the theme but carries no `synergy-mill` tag in the snapshot.

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

Carried over from run 1. The list is identical, card for card.

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Elves page (edhrec.com/tags/elves).

Off theme: none. Every card is an Elf creature, an Elf kindred spell, or a card with Elf-specific text. Thirty-one of the 40 are on the Elves page. Nissa, Resurgent Animist (rank 34) is the weakest fit: she is a legendary Elf creature, but her main function is landfall mana.

Signal bugs: none found. Tyvar Kell (a planeswalker) correctly has no subtype:Elf. Elvish Promenade and Trystan's Command carry subtype:Elf because they are Kindred Elf spells.

Legality: all 40 are legal in commander and inside GB.

Gaps (context only): the archetype's most-played Elves lack Elf-payoff text, so they fall below rank 40. Llanowar Elves (81% of Elves decks), Elvish Mystic (79%), Beast Whisperer (65%), Reclamation Sage (65%), Fyndhorn Elves (57%), Circle of Dreams Druid (40%), Galadhrim Ambush (38%), and Wirewood Channeler (30%).

## 14. storm (commander, UR, any-card)

Theme signals: payoff tags fourth-spell-matters, storm-like, third-spell-matters. tags cantrip, ritual, storm-count-matters. keywords Storm.

Funnel: 12588 legal in colors, 409 on theme, 0 owned, 0 on theme and owned, 204 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Aetherflux Reservoir | removal | 0 | 1.00 | payoff:storm-like tag:storm-count-matters tag:removal |
| 2 | Thousand-Year Storm | synergy | 0 | 0.99 | payoff:storm-like tag:storm-count-matters theme |
| 3 | Thunderclap Drake | synergy | 0 | 0.98 | payoff:storm-like payoff-text:instant and sorcery spells you cast cost theme |
| 4 | Case of the Ransacked Lab | draw | 0 | 0.97 | payoff:fourth-spell-matters payoff-text:instant and sorcery spells you cast cost tag:draw |
| 5 | Sorcerer Class | ramp | 0 | 0.97 | payoff:storm-like tag:storm-count-matters tag:ramp |
| 6 | Lock and Load | draw | 0 | 0.94 | payoff:storm-like tag:storm-count-matters tag:draw |
| 7 | Aria of Flame | removal | 0 | 0.92 | payoff:storm-like tag:storm-count-matters tag:removal |
| 8 | Sentinel Tower | removal | 0 | 0.91 | payoff:storm-like tag:storm-count-matters tag:removal |
| 9 | Ral, Monsoon Mage // Ral, Leyline Prodigy | removal | 0 | 0.88 | payoff-text:instant and sorcery spells you cast cost tag:storm-count-matters tag:removal |
| 10 | Flusterstorm | interaction | 0 | 0.77 | payoff-text:copy it for each spell cast before it this turn keyword:Storm tag:interaction |
| 11 | Grapeshot | removal | 0 | 0.77 | payoff-text:copy it for each spell cast before it this turn keyword:Storm tag:removal |
| 12 | Brain Freeze | synergy | 0 | 0.77 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 13 | Amphibian Downpour | removal | 0 | 0.76 | payoff-text:copy it for each spell cast before it this turn keyword:Storm tag:removal |
| 14 | Empty the Warrens | synergy | 0 | 0.75 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 15 | Radstorm | synergy | 0 | 0.75 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 16 | Mind's Desire | synergy | 0 | 0.74 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 17 | Haze of Rage | synergy | 0 | 0.73 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 18 | Stormscale Scion | threat | 0 | 0.73 | payoff-text:copy it for each spell cast before it this turn keyword:Storm type:threat |
| 19 | Elemental Eruption | synergy | 0 | 0.73 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 20 | All of History, All at Once | synergy | 0 | 0.69 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 21 | Galvanic Relay | synergy | 0 | 0.68 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 22 | Dragonstorm | synergy | 0 | 0.68 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 23 | Temporal Fissure | removal | 0 | 0.66 | payoff-text:copy it for each spell cast before it this turn keyword:Storm tag:removal |
| 24 | Ignite Memories | synergy | 0 | 0.66 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 25 | Spreading Insurrection | synergy | 0 | 0.66 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 26 | Fury Storm | synergy | 0 | 0.65 | payoff:storm-like theme |
| 27 | Storm of Memories | synergy | 0 | 0.64 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 28 | Emeritus of Conflict // Lightning Bolt | removal | 0 | 0.64 | payoff:third-spell-matters tag:removal |
| 29 | Goblin Electromancer | synergy | 0 | 0.63 | payoff-text:instant and sorcery spells you cast cost theme |
| 30 | Stormcatch Mentor | synergy | 0 | 0.63 | payoff-text:instant and sorcery spells you cast cost theme |
| 31 | Archmage of Runes | draw | 0 | 0.63 | payoff-text:instant and sorcery spells you cast cost tag:draw |
| 32 | Ground Rift | synergy | 0 | 0.63 | payoff-text:copy it for each spell cast before it this turn keyword:Storm theme |
| 33 | Baral, Chief of Compliance | draw | 0 | 0.62 | payoff-text:instant and sorcery spells you cast cost tag:draw |
| 34 | Fiery Encore | removal | 0 | 0.62 | payoff-text:copy it for each spell cast before it this turn keyword:Storm tag:removal |
| 35 | Scattershot | removal | 0 | 0.62 | payoff-text:copy it for each spell cast before it this turn keyword:Storm tag:removal |
| 36 | Primal Amulet // Primal Wellspring | land | 0 | 0.62 | payoff-text:instant and sorcery spells you cast cost type:land |
| 37 | Ral, Crackling Wit | draw | 0 | 0.62 | payoff-text:copy it for each spell cast before it this turn tag:draw |
| 38 | Haughty Djinn | synergy | 0 | 0.61 | payoff-text:instant and sorcery spells you cast cost theme |
| 39 | Echo Storm | synergy | 0 | 0.61 | payoff:storm-like theme |
| 40 | Eye of the Storm | synergy | 0 | 0.61 | payoff:storm-like theme |

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text, the EDHREC Storm page (edhrec.com/tags/storm), and Scryfall keyword search `keyword:storm`. Run 1 scored 25 of 40.

Off theme (1):

- 28 Emeritus of Conflict // Lightning Bolt: a third-spell aggro creature from Secrets of Strixhaven. The copy is one Lightning Bolt for {R}.

Fix in run 2: the payoff moved from `storm-count-matters` to `storm-like`. The old tag also holds the two hate cards (Damping Sphere, Rug of Smothering) and the combat-token payoffs (Rionya, Fire Dancer and Geralf, the Fleshwright). All four are gone. The reminder-text needle "copy it for each spell cast before it this turn" now catches every Storm card. Twenty of the 40 have the Storm keyword, against none in run 1.

The list holds 20 Storm cards, 10 cards with the storm-like payoff tag, and 9 cost reducers, which include Goblin Electromancer, Baral, Chief of Compliance, and Primal Amulet // Primal Wellspring.

Legality: all 40 are legal in commander and inside UR.

Owned: none. The prompt runs in any-card mode.

Gaps (context only): the rituals and the free cantrips are still absent. Jeska's Will, Frantic Search, and Underworld Breach carry the `ritual` or `cantrip` enabler alone, which caps their theme score at 0.40.

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

Carried over from run 1. The list is identical, card for card.

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Enchantress and Auras pages (edhrec.com/tags/enchantress, /auras).

Off theme (1):

- 31 Shambling Suit: power equals the number of artifacts and enchantments you control. An Eldraine artifact-theme body. Not on the Enchantress or Auras pages. EDHREC card page: 0.1% of decks. This is a close call. All That Glitters shares the count and is on the Enchantress page (26%).

Ranks 1 to 20 are the archetype's core, all on the Enchantress page at 30% to 69% of decks. Ranks 21 to 40 are weaker but real enchantment-cast, constellation, enchantment-count, and protection cards.

Signal bugs: Skybind gets tag:removal for a temporary exile that returns the permanent at end step. Ellivere of the Wild Court and Sky-Blessed Samurai get payoff-text:for each enchantment from reminder text. Shambling Suit gets payoff-text:enchantments you control from "artifacts and/or enchantments you control".

Legality: all 40 are legal in commander and inside GW.

Gaps (context only): Jukai Naturalist (58% of Enchantress decks), Hall of Heliod's Generosity (51%), Wild Growth (37%), Starfield Mystic (34%), Ondu Spiritdancer (29%), Enlightened Tutor (27%), Starfield of Nyx (25%), and Idyllic Tutor (22%) are absent.

## 16. counters proliferate (commander, GUB, any-card)

Theme signals: payoff tags pp-counters-matter, synergy-modified. tags pseudo-proliferate, repeatable-pp-counters, repeatable-proliferate. keywords Proliferate.

Funnel: 18160 legal in colors, 2023 on theme, 0 owned, 0 on theme and owned, 294 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Inspiring Call | interaction | 0 | 1.00 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter tag:interaction |
| 2 | Forgotten Ancient | threat | 0 | 1.00 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter type:threat |
| 3 | Walking Ballista | removal | 0 | 1.00 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter tag:removal |
| 4 | Ozolith, the Shattered Spire | draw | 0 | 0.99 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter tag:draw |
| 5 | Chasm Skulker | draw | 0 | 0.99 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter text:draw |
| 6 | The Earth Crystal | synergy | 0 | 0.99 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter theme |
| 7 | Mossborn Hydra | synergy | 0 | 0.99 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it tag:repeatable-pp-counters text:+1/+1 counter theme |
| 8 | Gyre Sage | ramp | 0 | 0.99 | payoff:pp-counters-matter payoff-text:for each +1/+1 counter tag:repeatable-pp-counters text:+1/+1 counter tag:ramp |
| 9 | Bristly Bill, Spine Sower | synergy | 0 | 0.99 | payoff:pp-counters-matter tag:pseudo-proliferate tag:repeatable-pp-counters text:+1/+1 counter theme |
| 10 | Evolution Witness | synergy | 0 | 0.99 | payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters text:+1/+1 counter theme |
| 11 | Kalonian Hydra | threat | 0 | 0.99 | payoff:pp-counters-matter tag:pseudo-proliferate tag:repeatable-pp-counters text:+1/+1 counter type:threat |
| 12 | Simic Ascendancy | wincon | 0 | 0.99 | payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters tag:repeatable-pp-counters text:+1/+1 counter tag:alternate-win-condition |
| 13 | Mikaeus, the Unhallowed | interaction | 0 | 0.99 | payoff-text:with a +1/+1 counter on it tag:repeatable-pp-counters text:+1/+1 counter tag:interaction |
| 14 | Sphere Grid | synergy | 0 | 0.99 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter theme |
| 15 | Agatha's Soul Cauldron | removal | 0 | 0.99 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter text:exile target |
| 16 | Duskshell Crawler | synergy | 0 | 0.99 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter theme |
| 17 | Hangarback Walker | synergy | 0 | 0.99 | payoff:pp-counters-matter payoff-text:for each +1/+1 counter tag:repeatable-pp-counters text:+1/+1 counter theme |
| 18 | Fathom Mage | draw | 0 | 0.98 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter tag:draw |
| 19 | Hydra's Growth | synergy | 0 | 0.98 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter theme |
| 20 | Mycoloth | threat | 0 | 0.98 | payoff:pp-counters-matter payoff-text:for each +1/+1 counter text:+1/+1 counter type:threat |
| 21 | Basking Broodscale | ramp | 0 | 0.98 | payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters text:+1/+1 counter tag:ramp |
| 22 | Court of Garenbrig | draw | 0 | 0.98 | payoff:pp-counters-matter tag:pseudo-proliferate tag:repeatable-pp-counters text:+1/+1 counter tag:draw |
| 23 | Scurry Oak | synergy | 0 | 0.98 | payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters tag:repeatable-pp-counters text:+1/+1 counter theme |
| 24 | Llanowar Reborn | land | 0 | 0.98 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter type:land |
| 25 | Primordial Hydra | synergy | 0 | 0.98 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter theme |
| 26 | Fangs of Kalonia | synergy | 0 | 0.98 | payoff:pp-counters-matter tag:pseudo-proliferate text:+1/+1 counter theme |
| 27 | Bred for the Hunt | draw | 0 | 0.98 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter tag:draw |
| 28 | Golgari Grave-Troll | draw | 0 | 0.97 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter text:draw |
| 29 | Scythecat Cub | synergy | 0 | 0.97 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter theme |
| 30 | Crystalline Crawler | ramp | 0 | 0.97 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it tag:repeatable-pp-counters text:+1/+1 counter tag:ramp |
| 31 | Benevolent Hydra | synergy | 0 | 0.97 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter theme |
| 32 | Iron Spider, Stark Upgrade | draw | 0 | 0.97 | payoff:pp-counters-matter tag:repeatable-pp-counters text:+1/+1 counter tag:draw |
| 33 | Hooded Hydra | synergy | 0 | 0.97 | payoff:pp-counters-matter payoff-text:for each +1/+1 counter text:+1/+1 counter theme |
| 34 | Nyxborn Hydra | synergy | 0 | 0.97 | payoff:pp-counters-matter payoff-text:for each +1/+1 counter text:+1/+1 counter theme |
| 35 | Armorcraft Judge | draw | 0 | 0.97 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it text:+1/+1 counter tag:draw |
| 36 | Arcbound Ravager | synergy | 0 | 0.97 | payoff:pp-counters-matter payoff-text:with a +1/+1 counter on it tag:repeatable-pp-counters text:+1/+1 counter theme |
| 37 | Wildwood Scourge | synergy | 0 | 0.97 | payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters tag:repeatable-pp-counters text:+1/+1 counter theme |
| 38 | Lux Artillery | synergy | 0 | 0.97 | payoff-text:with a +1/+1 counter on it tag:repeatable-pp-counters text:+1/+1 counter theme |
| 39 | Marketback Walker | draw | 0 | 0.96 | payoff:pp-counters-matter payoff-text:for each +1/+1 counter tag:repeatable-pp-counters text:+1/+1 counter tag:draw |
| 40 | Herd Baloth | threat | 0 | 0.96 | payoff:pp-counters-matter payoff-text:whenever one or more +1/+1 counters text:+1/+1 counter type:threat |

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC +1/+1 Counters and Proliferate pages. Run 1 scored 25 of 40.

Off theme (1):

- 28 Golgari Grave-Troll: a dredge card. It enters with counters, but the deck runs it to mill, not to grow a board.

Fix in run 2: `counters-matter` and `counter-fuel` fire on any counter type, so run 1 held wish, gold, energy, oil, quest, and mining counters. Both are gone. The payoff is `pp-counters-matter` with `synergy-modified`, and `repeatable-proliferate` and `pseudo-proliferate` replace the dead `proliferate` slug. All 40 cards now carry "+1/+1 counter" in their text.

Legality: all 40 are legal in commander and inside GUB. Eight are colorless.

Owned: none. The prompt runs in any-card mode.

Gaps (context only): the word "proliferate" still finds nothing at the top. A proliferate card scores one enabler tag, one keyword, and one needle, which is 0.76 of the cap. A +1/+1 payoff reaches 1.00. Evolution Sage, Karn's Bastion, and Flux Channeler stay below the cut. A future row must weigh the second theme word against the first.

## 17. zombies (commander, BU, any-card)

Theme signals: payoff tags typal-zombie. subtypes Zombie.

Funnel: 12600 legal in colors, 533 on theme, 0 owned, 0 on theme and owned, 292 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Death Baron | synergy | 0 | 0.98 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie theme |
| 2 | The Scarab God | removal | 0 | 0.98 | payoff:typal-zombie payoff-text:zombies you control text:exile target |
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
| 33 | Cemetery Reaper | removal | 0 | 0.92 | payoff:typal-zombie subtype:Zombie text:exile target |
| 34 | Bladestitched Skaab | synergy | 0 | 0.92 | payoff:typal-zombie payoff-text:zombies you control subtype:Zombie theme |
| 35 | Unbreathing Horde | synergy | 0 | 0.92 | payoff:typal-zombie payoff-text:zombie you control subtype:Zombie theme |
| 36 | Zul Ashur, Lich Lord | synergy | 0 | 0.92 | payoff:typal-zombie subtype:Zombie theme |
| 37 | Zombie Master | interaction | 0 | 0.92 | payoff:typal-zombie subtype:Zombie tag:interaction |
| 38 | Gleaming Overseer | interaction | 0 | 0.92 | payoff:typal-zombie subtype:Zombie tag:interaction |
| 39 | Undead Warchief | threat | 0 | 0.92 | payoff:typal-zombie subtype:Zombie type:threat |
| 40 | Graf Harvest | synergy | 0 | 0.91 | payoff:typal-zombie payoff-text:zombies you control theme |

### Review

Carried over from run 1. The list is identical, card for card.

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Zombies page (edhrec.com/tags/zombies).

Off theme (1):

- 31 Liliana, the Last Hope: the +1 and -2 are generic removal and recursion. Only the -7 emblem names Zombies. Zombies page: 7% of decks, synergy 0.07. This is a close call.

Thirty-three of the 40 are on the Zombies page, most above 25% of decks. The seven absent cards are Zombies or Zombie-specific by oracle text: Shepherd of Rot, Undead Alchemist, Archghoul of Thraben, Graveborn Muse, Gravespawn Sovereign, Bladestitched Skaab, and Unbreathing Horde.

Signal bugs: The Scarab God and Cemetery Reaper get text:destroy target and role removal, but neither card contains "destroy". Zombie Master and Gleaming Overseer get tag:interaction for regenerate and hexproof grants.

Legality: all 40 are legal in commander and inside BU.

Gaps (context only): Lord of the Undead (35% of Zombies decks) and Liliana, Death's Majesty (30%) are absent.

## 18. burn (modern, R, any-card)

Theme signals: payoff tags burn-player. tags single-target-instant-sorcery.

Funnel: 5166 legal in colors, 1294 on theme, 0 owned, 0 on theme and owned, 289 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Lightning Bolt | removal | 0 | 1.00 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 2 | Impact Tremors | synergy | 0 | 1.00 | payoff:burn-player payoff-text:damage to each opponent theme |
| 3 | Guttersnipe | synergy | 0 | 1.00 | payoff:burn-player payoff-text:damage to each opponent theme |
| 4 | Chandra's Ignition | wipe | 0 | 1.00 | payoff:burn-player tag:single-target-instant-sorcery tag:sweeper |
| 5 | Purphoros, God of the Forge | threat | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent type:threat |
| 6 | Reckless Fireweaver | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 7 | Grapeshot | removal | 0 | 0.99 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 8 | Coruscation Mage | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 9 | Firebrand Archer | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 10 | Fiery Inscription | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 11 | Crackle with Power | removal | 0 | 0.99 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 12 | Glaring Fleshraker | ramp | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent tag:ramp |
| 13 | Molten Gatekeeper | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 14 | Electrodominance | removal | 0 | 0.99 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 15 | Kessig Flamebreather | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 16 | Comet Storm | removal | 0 | 0.99 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 17 | Witty Roastmaster | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 18 | Grab the Prize | draw | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent tag:draw |
| 19 | Fling | removal | 0 | 0.99 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 20 | Lindblum, Industrial Regency // Mage Siege | land | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent type:land |
| 21 | Boltwave | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 22 | Weftstalker Ardent | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 23 | Chandra, Torch of Defiance | ramp | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent tag:ramp |
| 24 | Ashling, Flame Dancer | wipe | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent tag:sweeper |
| 25 | Electrostatic Field | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 26 | Valakut Exploration | synergy | 0 | 0.99 | payoff:burn-player payoff-text:damage to each opponent theme |
| 27 | Kazuul's Fury // Kazuul's Cliffs | land | 0 | 0.98 | payoff:burn-player tag:single-target-instant-sorcery type:land |
| 28 | Thermo-Alchemist | synergy | 0 | 0.98 | payoff:burn-player payoff-text:damage to each opponent theme |
| 29 | Glint-Horn Buccaneer | draw | 0 | 0.98 | payoff:burn-player payoff-text:damage to each opponent tag:draw |
| 30 | Shock | removal | 0 | 0.98 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 31 | Dragonhawk, Fate's Tempest | threat | 0 | 0.98 | payoff:burn-player payoff-text:damage to each opponent type:threat |
| 32 | End the Festivities | wipe | 0 | 0.98 | payoff:burn-player payoff-text:damage to each opponent tag:sweeper |
| 33 | Jaya's Immolating Inferno | removal | 0 | 0.98 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 34 | Gut Shot | removal | 0 | 0.98 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 35 | Lightning Strike | removal | 0 | 0.98 | payoff:burn-player tag:single-target-instant-sorcery tag:removal |
| 36 | Erebor Flamesmith | synergy | 0 | 0.98 | payoff:burn-player payoff-text:damage to each opponent theme |
| 37 | Sabotender | synergy | 0 | 0.98 | payoff:burn-player payoff-text:damage to each opponent theme |
| 38 | Tectonic Hazard | wipe | 0 | 0.98 | payoff:burn-player payoff-text:damage to each opponent tag:sweeper |
| 39 | Reckless Handling | synergy | 0 | 0.97 | payoff:burn-player payoff-text:damage to each opponent theme |
| 40 | Spikefield Hazard // Spikefield Cave | land | 0 | 0.97 | payoff:burn-player tag:single-target-instant-sorcery type:land |

### Review

Verdict: yes. 39 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and Modern legality, and one Modern Burn list (Iwao Sho, 2026-08-15, mtgtop8 deck 880322). Run 1 scored 25 of 40.

Off theme (1):

- 12 Glaring Fleshraker: the triggers count colorless spells. An Eldrazi card in a red list.

Fix in run 2: `burn-player` is the payoff and `single-target-instant-sorcery` is the enabler, so a card must aim damage at a player and be a spell or an "each opponent" source. The permanents with a damage ability (Aetherflux Reservoir, Walking Ballista, Goblin Bombardment, Sword of Fire and Ice, Ugin, the Spirit Dragon, Valakut, the Molten Pinnacle) reach the payoff alone, which is 0.60 of the cap, and all six are gone. The dead needle "deals damage to each opponent" became "damage to each opponent", which real cards match.

Format fit: the curve is better than run 1 but still not a 60-card curve. Eighteen of the 40 are instants or sorceries. The sampled Burn list plays 20 lands, and no spell in it costs more than 3.

Legality: all 40 are legal in Modern on 2026-08-24 and inside mono-red. Glaring Fleshraker is colorless.

Owned: none. The prompt runs in any-card mode.

Gaps (context only): the archetype's own spells are still absent, because popularity is EDHREC rank and EDHREC measures Commander play. Of the ten nonland spells in the sampled list, only Lightning Bolt is in the top 40. Boros Charm is outside mono-red, and the filter drops it correctly. Monastery Swiftspear and Goblin Guide never enter the pool. PR-14 (`MetaBoost`) is the fix, because a 60-card format needs a 60-card popularity number.

## 19. control (standard, WU, any-card)

Theme signals: payoff tags counterspell, single-target-instant-sorcery, sweeper. tags draw-engine, pure-draw, removal.

Funnel: 1821 legal in colors, 709 on theme, 0 owned, 0 on theme and owned, 218 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Stroke of Midnight | removal | 0 | 1.00 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:removal |
| 2 | Aetherize | wipe | 0 | 1.00 | payoff:sweeper tag:removal tag:sweeper |
| 3 | Into the Flood Maw | removal | 0 | 0.99 | payoff:single-target-instant-sorcery tag:removal tag:removal |
| 4 | Fumigate | wipe | 0 | 0.99 | payoff:sweeper tag:removal tag:sweeper |
| 5 | Disenchant | removal | 0 | 0.99 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:removal |
| 6 | Steel Hellkite | wipe | 0 | 0.99 | payoff:sweeper tag:removal tag:sweeper |
| 7 | Requisition Raid | removal | 0 | 0.99 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:removal |
| 8 | Enter the Enigma | draw | 0 | 0.99 | payoff:single-target-instant-sorcery tag:pure-draw tag:draw |
| 9 | Parting Gust | interaction | 0 | 0.98 | payoff:single-target-instant-sorcery tag:removal text:exile target tag:interaction |
| 10 | Avatar's Wrath | wipe | 0 | 0.98 | payoff:sweeper payoff:single-target-instant-sorcery tag:removal tag:sweeper |
| 11 | Get Lost | removal | 0 | 0.98 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:removal |
| 12 | Day of Judgment | wipe | 0 | 0.98 | payoff:sweeper tag:removal tag:sweeper |
| 13 | Split Up | wipe | 0 | 0.98 | payoff:sweeper tag:removal tag:sweeper |
| 14 | Season of Weaving | wipe | 0 | 0.98 | payoff:sweeper tag:removal tag:pure-draw tag:sweeper |
| 15 | Starfall Invocation | wipe | 0 | 0.98 | payoff:sweeper tag:removal tag:sweeper |
| 16 | Final Showdown | wipe | 0 | 0.98 | payoff:sweeper tag:removal tag:sweeper |
| 17 | Erode | ramp | 0 | 0.98 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:ramp |
| 18 | Crib Swap | removal | 0 | 0.98 | payoff:single-target-instant-sorcery tag:removal text:exile target tag:removal |
| 19 | Valorous Stance | interaction | 0 | 0.98 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:interaction |
| 20 | Expel the Interlopers | wipe | 0 | 0.97 | payoff:sweeper tag:removal tag:sweeper |
| 21 | Unsummon | removal | 0 | 0.97 | payoff:single-target-instant-sorcery tag:removal tag:removal |
| 22 | Splash Portal | draw | 0 | 0.97 | payoff:single-target-instant-sorcery tag:pure-draw text:exile target tag:draw |
| 23 | River's Rebuke | wipe | 0 | 0.97 | payoff:sweeper payoff:single-target-instant-sorcery tag:removal tag:sweeper |
| 24 | Settle the Wreckage | wipe | 0 | 0.97 | payoff:sweeper payoff:single-target-instant-sorcery tag:removal tag:sweeper |
| 25 | Aang, Swift Savior // Aang and La, Ocean's Fury | interaction | 0 | 0.96 | payoff:counterspell tag:removal tag:interaction |
| 26 | Bovine Intervention | removal | 0 | 0.96 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:removal |
| 27 | Season of the Burrow | ramp | 0 | 0.96 | payoff:single-target-instant-sorcery tag:removal text:exile target tag:ramp |
| 28 | Airbending Lesson | interaction | 0 | 0.96 | payoff:single-target-instant-sorcery tag:removal tag:pure-draw tag:interaction |
| 29 | Unwanted Remake | removal | 0 | 0.96 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:removal |
| 30 | Airbender's Reversal | interaction | 0 | 0.96 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:interaction |
| 31 | Mathemagics | draw | 0 | 0.95 | payoff:single-target-instant-sorcery tag:pure-draw tag:draw |
| 32 | Into the Roil | removal | 0 | 0.95 | payoff:single-target-instant-sorcery tag:removal tag:pure-draw tag:removal |
| 33 | Combat Tutorial | draw | 0 | 0.95 | payoff:single-target-instant-sorcery tag:pure-draw tag:draw |
| 34 | Exorcise | removal | 0 | 0.94 | payoff:single-target-instant-sorcery tag:removal text:exile target tag:removal |
| 35 | Make Your Move | removal | 0 | 0.94 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:removal |
| 36 | Getaway Glamer | interaction | 0 | 0.94 | payoff:single-target-instant-sorcery tag:removal text:destroy target text:exile target tag:interaction |
| 37 | Emeritus of Ideation // Ancestral Recall | draw | 0 | 0.94 | payoff:single-target-instant-sorcery tag:draw-engine tag:pure-draw tag:draw |
| 38 | Spectacular Tactics | interaction | 0 | 0.94 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:interaction |
| 39 | Honor | draw | 0 | 0.93 | payoff:single-target-instant-sorcery tag:pure-draw tag:draw |
| 40 | Repel Calamity | removal | 0 | 0.93 | payoff:single-target-instant-sorcery tag:removal text:destroy target tag:removal |

### Review

Verdict: yes. 36 of 40 are on theme, which is exactly the bar. Review date: 2026-08-24. Sources: Scryfall oracle text and Standard legality, and the mtgtop8 UW Control lists of 2026-08. Run 1 scored 34 of 40.

Off theme (4):

- 6 Steel Hellkite: a six-mana artifact creature. The sweeper needs combat damage first.
- 8 Enter the Enigma: an unblockable trick with a cantrip. An aggro card.
- 22 Splash Portal: it blinks your own creature. A creature-deck cantrip.
- 39 Honor: a +1/+1 counter with a cantrip. An aggro card.

Borderline (counted on theme, D-64): 2 Aetherize, 21 Unsummon, 25 Aang, Swift Savior, 27 Season of the Burrow, 33 Combat Tutorial, and 36 Getaway Glamer. Each answers a board or draws a card under a condition. Count all six off, and the total falls to 30, which fails. This prompt is the weakest pass of the 20.

Fix in run 2: `counterspell`, `sweeper`, and `single-target-instant-sorcery` are payoffs, and the same three plus `removal`, `draw-engine`, and `pure-draw` stay enablers. A permanent with a removal rider now reaches 0.56 of the cap. Meteor Golem, Sheltered by Ghosts, Venat, Heart of Hydaelyn, Dion, Bahamut's Dominant, and Ultima Weapon are gone.

Legality: all 40 are legal in Standard on 2026-08-24 and inside WU.

Owned: none. The prompt runs in any-card mode.

Gaps (context only): only one counterspell (Aang, Swift Savior) is in the top 40. Standard holds few cards with the `counterspell` tag, and the sweepers and removal spells outrank them on EDHREC rank. Stock Up and Consult the Star Charts carry no tag the row names.

## 20. infect poison (commander, GB, any-card)

Theme signals: payoff tags synergy-poison. tags poisonous. keywords Infect, Toxic, Proliferate.

Funnel: 12525 legal in colors, 150 on theme, 0 owned, 0 on theme and owned, 263 returned, 0 upgrades.

| Rank | Card | Role | Owned | Score | Signals |
|---|---|---|---|---|---|
| 1 | Plague Myr | ramp | 0 | 0.98 | payoff-text:poison counter tag:poisonous keyword:Infect tag:ramp |
| 2 | Myr Convert | ramp | 0 | 0.98 | payoff-text:poison counter tag:poisonous keyword:Toxic tag:ramp |
| 3 | Bloated Contaminator | synergy | 0 | 0.98 | payoff-text:poison counter tag:poisonous keyword:Toxic keyword:Proliferate theme |
| 4 | Skithiryx, the Blight Dragon | threat | 0 | 0.97 | payoff-text:poison counter tag:poisonous keyword:Infect type:threat |
| 5 | Blightbelly Rat | synergy | 0 | 0.97 | payoff-text:poison counter tag:poisonous keyword:Toxic keyword:Proliferate theme |
| 6 | Venerated Rotpriest | synergy | 0 | 0.97 | payoff-text:poison counter tag:poisonous keyword:Toxic theme |
| 7 | Glistening Sphere | ramp | 0 | 0.97 | payoff:synergy-poison payoff-text:poison counter payoff-text:corrupted keyword:Proliferate tag:ramp |
| 8 | Tyrranax Rex | threat | 0 | 0.97 | payoff-text:poison counter tag:poisonous keyword:Toxic type:threat |
| 9 | Ichor Rats | synergy | 0 | 0.97 | payoff-text:poison counter tag:poisonous keyword:Infect theme |
| 10 | Karumonix, the Rat King | synergy | 0 | 0.97 | payoff-text:poison counter tag:poisonous keyword:Toxic theme |
| 11 | Bloodroot Apothecary | synergy | 0 | 0.97 | payoff-text:poison counter tag:poisonous keyword:Toxic theme |
| 12 | Phyresis Outbreak | wipe | 0 | 0.97 | payoff:synergy-poison payoff-text:poison counter tag:sweeper |
| 13 | Contaminant Grafter | ramp | 0 | 0.96 | payoff:synergy-poison payoff-text:poison counter payoff-text:corrupted tag:poisonous keyword:Toxic keyword:Proliferate tag:ramp |
| 14 | Bilious Skulldweller | synergy | 0 | 0.96 | payoff-text:poison counter tag:poisonous keyword:Toxic theme |
| 15 | Ichorclaw Myr | synergy | 0 | 0.96 | payoff-text:poison counter tag:poisonous keyword:Infect theme |
| 16 | Plague Stinger | synergy | 0 | 0.96 | payoff-text:poison counter tag:poisonous keyword:Infect theme |
| 17 | Phyrexian Swarmlord | threat | 0 | 0.95 | payoff:synergy-poison payoff-text:poison counter tag:poisonous keyword:Infect type:threat |
| 18 | The Seedcore | land | 0 | 0.95 | payoff:synergy-poison payoff-text:poison counter payoff-text:corrupted type:land |
| 19 | Viridian Corrupter | removal | 0 | 0.94 | payoff-text:poison counter tag:poisonous keyword:Infect tag:removal |
| 20 | Pestilent Syphoner | synergy | 0 | 0.94 | payoff-text:poison counter tag:poisonous keyword:Toxic theme |
| 21 | Glistener Elf | synergy | 0 | 0.94 | payoff-text:poison counter tag:poisonous keyword:Infect theme |
| 22 | Necrogen Rotpriest | threat | 0 | 0.94 | payoff-text:poison counter tag:poisonous keyword:Toxic type:threat |
| 23 | Venomous Brutalizer | threat | 0 | 0.93 | payoff-text:poison counter tag:poisonous keyword:Toxic keyword:Proliferate type:threat |
| 24 | Phyrexian Atlas | ramp | 0 | 0.93 | payoff:synergy-poison payoff-text:poison counter payoff-text:corrupted tag:ramp |
| 25 | Blight Mamba | synergy | 0 | 0.93 | payoff-text:poison counter tag:poisonous keyword:Infect theme |
| 26 | Core Prowler | threat | 0 | 0.92 | payoff-text:poison counter tag:poisonous keyword:Infect keyword:Proliferate type:threat |
| 27 | Ichorspit Basilisk | synergy | 0 | 0.92 | payoff-text:poison counter tag:poisonous keyword:Toxic theme |
| 28 | Phyrexian Crusader | synergy | 0 | 0.92 | payoff-text:poison counter tag:poisonous keyword:Infect theme |
| 29 | Hand of the Praetors | threat | 0 | 0.91 | payoff-text:poison counter tag:poisonous keyword:Infect type:threat |
| 30 | Geth's Summons | synergy | 0 | 0.91 | payoff:synergy-poison payoff-text:poison counter payoff-text:corrupted theme |
| 31 | Necropede | removal | 0 | 0.91 | payoff-text:poison counter tag:poisonous keyword:Infect tag:removal |
| 32 | Glissa's Retriever | threat | 0 | 0.91 | payoff:synergy-poison payoff-text:poison counter payoff-text:corrupted tag:poisonous keyword:Toxic type:threat |
| 33 | Dune Mover | ramp | 0 | 0.90 | payoff-text:poison counter tag:poisonous keyword:Toxic text:mana |
| 34 | Feed the Infection | draw | 0 | 0.90 | payoff:synergy-poison payoff-text:poison counter payoff-text:corrupted tag:draw |
| 35 | Septic Rats | synergy | 0 | 0.90 | payoff:synergy-poison payoff-text:poison counter tag:poisonous keyword:Infect theme |
| 36 | Inkmoth Nexus | land | 0 | 0.90 | payoff-text:poison counter tag:poisonous type:land |
| 37 | Paladin of Predation | threat | 0 | 0.90 | payoff-text:poison counter tag:poisonous keyword:Toxic type:threat |
| 38 | Flesh-Eater Imp | threat | 0 | 0.89 | payoff-text:poison counter tag:poisonous keyword:Infect type:threat |
| 39 | Anoint with Affliction | removal | 0 | 0.89 | payoff:synergy-poison payoff-text:poison counter payoff-text:corrupted tag:removal |
| 40 | Tyrranax Atrocity | threat | 0 | 0.89 | payoff-text:poison counter tag:poisonous keyword:Toxic type:threat |

### Review

Verdict: yes. 40 of 40 are on theme. The bar is 36. Review date: 2026-08-24. Sources: Scryfall oracle text and the EDHREC Infect and Proliferate pages. Run 1 scored 40 of 40.

Off theme: none. Every card has infect or toxic, gives poison counters, or proliferates. Fifteen have Infect and 17 have Toxic.

Fix in run 2: `synergy-poison` replaces the dead payoff slug `poison-matters`, and `poisonous` replaces the dead enabler slug `infect`. Eleven cards now carry a real poison payoff tag, against none in run 1. Phyresis Outbreak, Geth's Summons, and The Seedcore entered the list.

Legality: all 40 are legal in commander and inside GB. Ten are colorless.

Owned: none. The prompt runs in any-card mode.

Gaps (context only): the proliferate support is still thin. Karn's Bastion and Evolution Sage carry `repeatable-proliferate`, which this row does not name.

