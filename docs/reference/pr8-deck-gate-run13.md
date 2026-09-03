# PR-8 deck gate

Run date: 2026-09-02. Card snapshot: 2026-09-02.

Verdict: FAIL. 22 of 24 decks passed every block check, 6 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 24 |
| Decks returned | 24 |
| Decks with no block finding | 22 |
| Invented names that reached the user | 6 |
| Decks that needed the repair turn | 12 |
| Summaries judged (F-26) | 24 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 66 |
| Cost | $2.6424 |
| Time | 2295 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 24 |
| `not_owned` | 21 |
| `summary_rules_claim` | 15 |
| `profile_off_band` | 6 |
| `basics_added` | 2 |
| `outside_requested_set` | 2 |
| `precon_cards_restored` | 1 |
| `deck_size` | 1 |
| `copy_limit` | 1 |
| `two_card_combo` | 1 |

By severity: BLOCK 2. WARN 45. INFO 27. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 17 | 17 |
| 22 | a set family and a card from outside it | `hob,hoc` | 170 | 0 | 1 | 1 |
| 23 | two set families at once | `blb,blc,hob,hoc,pblb` | 201 | 0 | 0 | 0 |
| 24 | a 60-card deck from one set | `blb,blc,pblb` | 95 | 0 | 0 | 0 |

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $430.11 to buy, $430.11 the whole deck.

**Summary:** Karlov of the Ghost Council leads a resilient Orzhov lifegain strategy that develops its mana, repeatedly gains life through creatures and artifacts, and turns those triggers into growing attackers, draining effects, token armies, and powerful late-game pressure. Protective tools and broad removal keep the central engine intact while card-advantage pieces ensure the deck can continue applying pressure through longer multiplayer games.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. more opening hands hold two to four lands than the norm, and that raises the grade. the deck makes less mana on turn four than the norm, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 3.29 over 63 nonland cards
- [WARN] `two_card_combo`: bracket 3 allows no two-card infinite combo that needs four mana or less, and the deck holds Enduring Angel // Angelic Enforcer + Enduring Tenacity (speed 5, near two-card); Enduring Angel // Angelic Enforcer + Vito, Thorn of the Dusk Rose (speed 5, near two-card)

<details><summary>The deck list</summary>

- 10 Plains | land | Basic white sources for early lifegain spells and Karlov.
- 13 Swamp | land | Basic black sources for the deck's heavier black requirements.
- 1 Godless Shrine | land | Untapped-or-later dual source for both colors.
- 1 Caves of Koilos | land | Untapped dual source.
- 1 Vault of Champions | land | Reliable multiplayer dual source.
- 1 Shattered Sanctum | land | Dual source that is usually untapped as the game progresses.
- 1 Isolated Chapel | land | Dual source supported by the basic land base.
- 1 Scoured Barrens | land | Dual source that also triggers lifegain payoffs.
- 1 Restless Fortress | land | Dual source and resilient creature-land threat.
- 1 Command Tower | land | Reliable two-color fixing.
- 1 Exotic Orchard | land | Flexible multiplayer color fixing.
- 1 City of Brass | land | Untapped access to either deck color.
- 1 Mana Confluence | land | Untapped access to either deck color.
- 1 Reflecting Pool | land | Efficient fixing alongside the dense dual-land base.
- 1 Secluded Courtyard | land | Creature-heavy tribal fixing for the deck's Clerics and Vampires.
- 1 Altar of the Pantheon | ramp | Fixes mana while advancing toward larger lifegain threats.
- 1 Bounty Board | ramp | Low-cost mana development with later utility.
- 1 Colossal Plow | ramp | Provides burst mana for deploying larger creatures.
- 1 Cryptolith Fragment // Aurora of Emrakul | ramp | Fixes both colors and supplies repeatable mana.
- 1 Nuka-Cola Vending Machine | ramp | Treasure-based acceleration with useful lifegain texture.
- 1 Orazca Relic | ramp | Reliable mana rock with a card-draw fallback.
- 1 Phial of Galadriel | ramp | Mana acceleration that rewards drawing cards.
- 1 Potioner's Trove | ramp | Treasure production supported by lifegain.
- 1 Pristine Talisman | ramp | Mana rock that repeatedly triggers lifegain synergies.
- 1 The Celestus | ramp | Filtering, mana acceleration, and incidental life swings.
- 1 Archivist of Oghma | draw | Efficient card advantage that also gains life.
- 1 Dawn of Hope | draw | Turns surplus life into cards and creature production.
- 1 Enduring Innocence | draw | Sustained card advantage from small creature deployment.
- 1 Lunar Convocation | draw | Lifegain-oriented card advantage engine.
- 1 Mangara, the Diplomat | draw | Punishes opponents for overextending with cards.
- 1 Markov Purifier | draw | Draw engine supported by the Vampire core.
- 1 Sigarda's Splendor | draw | Life-total-based card advantage.
- 1 The Gaffer | draw | Repeatable draw while maintaining a healthy life total.
- 1 Tymna the Weaver | draw | Combat-based card advantage with lifegain upside.
- 1 Vampiric Rites | draw | Low-cost sacrifice outlet that converts creatures into cards and life.
- 1 Well of Lost Dreams | draw | Converts lifegain triggers into scalable card draw.
- 1 Alseid of Life's Bounty | interaction | Protects Karlov or a key payoff from targeted disruption.
- 1 Enduring Angel // Angelic Enforcer | interaction | Protective lifegain creature with a powerful transformed body.
- 1 Faith's Shield | interaction | Flexible protection for a critical permanent or combat step.
- 1 Metropolis Reformer | interaction | Defends the life total and grants protective coverage.
- 1 Restoration Magic | interaction | Protects the board and recovers from removal.
- 1 Sephara, Sky's Blade | interaction | Protects a developed flying board from destruction.
- 1 Sword of Light and Shadow | interaction | Protection plus recurring creature value.
- 1 Werefox Bodyguard | interaction | Creature-based disruption that preserves board presence.
- 1 Aetherflux Reservoir | removal | Lifegain payoff that can remove a player or decisive threat.
- 1 Ayli, Eternal Pilgrim | removal | Repeatable sacrifice outlet and permanent removal.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Flexible removal spell with a lifegain-adjacent Adventure.
- 1 Henrika Domnathi // Henrika, Infernal Seer | removal | Versatile creature interaction that can become an evasive threat.
- 1 Murderous Rider // Swift End | removal | Efficient answer that remains a creature afterward.
- 1 Nightmare's Thirst | removal | Cheap creature removal strengthened by the life total.
- 1 Solitude | removal | Immediate answer to problematic creatures.
- 1 Umezawa's Jitte | removal | Reusable combat control and utility.
- 1 Vona, Butcher of Magan | removal | Lifegain-supported repeatable permanent removal.
- 1 Fumigate | wipe | Creature reset that refills the life total.
- 1 Kaya's Wrath | wipe | Efficient creature sweep that rewards the deck's creature density.
- 1 The Battle of Bywater | wipe | Board reset that can preserve smaller utility creatures.
- 1 Ajani's Pridemate | synergy | Efficient scalable attacker from repeated lifegain.
- 1 Bloodthirsty Aerialist | synergy | Evasive creature that grows with every life-gain trigger.
- 1 Cleric Class | synergy | Enhances lifegain and rewards building the life total.
- 1 Cleric of Life's Bond | synergy | Creature deployment produces lifegain and scaling power.
- 1 Essence Channeler | synergy | Reliable lifegain trigger source for the creature plan.
- 1 Heliod, Sun-Crowned | synergy | Distributes counters from lifegain and supports combo-pressure lines.
- 1 Indulging Patrician | synergy | Converts substantial lifegain into opposing life loss.
- 1 Marauding Blight-Priest | synergy | Turns recurring lifegain into table-wide pressure.
- 1 Vito, Thorn of the Dusk Rose | synergy | Transforms life gained into direct opponent life loss.
- 1 Voice of the Blessed | synergy | Rapidly grows into a protected evasive threat.
- 1 Archangel of Thune | threat | Turns each lifegain event into a full-board growth trigger.
- 1 Attended Healer | threat | Builds a board of creatures while gaining life.
- 1 Celestine, the Living Saint | threat | Recurs valuable creatures through lifegain.
- 1 Crested Sunmare | threat | Creates a recurring army of difficult-to-answer attackers.
- 1 Divinity of Pride | threat | Large evasive attacker that benefits from a high life total.
- 1 Enduring Tenacity | threat | Weaponizes lifegain while remaining difficult to permanently answer.
- 1 Nykthos Paragon | threat | Converts large life gains into explosive board growth.
- 1 Regal Bloodlord | threat | Generates flying Vampire attackers from end-step lifegain.
- 1 Rhox Faithmender | threat | Doubles lifegain and enables overwhelming payoff turns.
- 1 Tivash, Gloom Summoner | threat | Creates powerful Demon tokens from a high life total.
- 1 Twinblade Paladin | threat | Becomes a double-striking finisher after substantial lifegain.
- 1 Valkyrie Harbinger | threat | Produces a steady stream of large Angel tokens.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $184.80 the whole deck.

**Summary:** This aristocrats build develops a creature-and-artifact board around its synergy pieces, keeps cards flowing with a deep draw package, and uses its creature threats to close games after opponents have been pressured by removal and board wipes. It gives up raw speed for a balanced, library-first mix of mana development, protection, and broad answers.

The quality model grades this deck good: the commander does not place in cEDH events, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade. the mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.05 over 63 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 14 Plains | land | Basic white mana source.
- 14 Swamp | land | Basic black mana source.
- 1 Command Tower | land | Flexible commander-color mana source.
- 1 City of Brass | land | Flexible colored mana source.
- 1 Exotic Orchard | land | Flexible colored mana source.
- 1 Grand Coliseum | land | Additional multicolor mana source.
- 1 Path of Ancestry | land | Multicolor land for the creature-heavy build.
- 1 Plaza of Heroes | land | Colored source that supports the legendary suite.
- 1 Secluded Courtyard | land | Creature-focused colored mana source.
- 1 Spire of Industry | land | Colored mana source alongside artifacts.
- 1 Call of the Ring | draw | Library draw option.
- 1 Grave Venerations | draw | Aristocrats-oriented draw support.
- 1 Idol of Oblivion | draw | Low-cost draw support.
- 1 Inspiring Overseer | draw | Creature-based draw support.
- 1 Lembas | draw | Artifact draw support.
- 1 Mask of Memory | draw | Equipment-based draw support.
- 1 Nasty End | draw | Black draw spell.
- 1 Night's Whisper | draw | Efficient black draw spell.
- 1 Skullclamp | draw | Core creature-focused draw engine.
- 1 Tome of Legends | draw | Repeatable artifact draw support.
- 1 Wall of Omens | draw | Early creature-based draw.
- 1 Bastion Protector | interaction | Creature protection for the commander.
- 1 Boromir, Warden of the Tower | interaction | Legendary creature interaction piece.
- 1 Clever Concealment | interaction | Protects the developed board.
- 1 Darksteel Plate | interaction | Durable equipment protection.
- 1 Gift of Immortality | interaction | Protective aura for a key creature.
- 1 Lightning Greaves | interaction | Efficient equipment protection.
- 1 Swiftfoot Boots | interaction | Additional creature protection.
- 1 Together Forever | interaction | Creature-focused protection piece.
- 1 Arcane Signet | ramp | Reliable two-color mana acceleration.
- 1 Bender's Waterskin | ramp | Artifact mana acceleration.
- 1 Deadly Dispute | ramp | Sacrifice-aligned mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost colored mana acceleration.
- 1 Lotho, Corrupt Shirriff | ramp | Creature-based mana acceleration.
- 1 Sol Ring | ramp | Efficient artifact mana acceleration.
- 1 Springleaf Drum | ramp | Low-cost creature-based mana acceleration.
- 1 Sword of the Animist | ramp | Equipment-based mana development.
- 1 Thought Vessel | ramp | Artifact mana acceleration.
- 1 Wayfarer's Bauble | ramp | Early mana development.
- 1 Bitter Triumph | removal | Flexible single-target answer.
- 1 Claim the Precious | removal | Black targeted removal.
- 1 Crib Swap | removal | Creature-focused instant-speed answer.
- 1 Fatal Push | removal | Efficient black removal.
- 1 Fiend Hunter | removal | Creature-based removal option.
- 1 Generous Gift | removal | Broad white permanent answer.
- 1 Infernal Grasp | removal | Reliable black creature answer.
- 1 Stroke of Midnight | removal | Broad white permanent answer.
- 1 Swords to Plowshares | removal | Efficient white creature answer.
- 1 Arcade Cabinet | synergy | Artifact synergy piece for the central plan.
- 1 Gollum the Abandoned | synergy | Black creature supporting the aristocrats plan.
- 1 Gollum, Patient Plotter | synergy | Black creature supporting the aristocrats plan.
- 1 Gríma Wormtongue | synergy | Legendary black synergy creature.
- 1 Heirloom Auntie | synergy | Creature synergy piece.
- 1 Nimble Hobbit | synergy | Low-cost creature synergy piece.
- 1 Palace Jailer | synergy | Creature-based support for the board plan.
- 1 Phantom Train | synergy | Artifact synergy support.
- 1 Rat King, Pale Piper | synergy | Black creature supporting the central plan.
- 1 The Sackville-Bagginses | synergy | Creature support for the aristocrats shell.
- 1 Angel of Serenity | threat | High-impact creature threat.
- 1 Archfiend of Ifnir | threat | Large black creature threat.
- 1 Bill the Pony | threat | Creature threat that adds board presence.
- 1 Bronze Guardian | threat | Artifact creature threat.
- 1 Cirith Ungol Patrol | threat | Black creature threat.
- 1 Exemplar of Light | threat | White creature threat.
- 1 Frontline Medic | threat | Creature threat that develops the board.
- 1 Kingpin's Enforcers | threat | Black creature threat.
- 1 Massacre Girl, Known Killer | threat | Legendary black creature threat.
- 1 Orcish Bowmasters | threat | Efficient black creature threat.
- 1 The Walls of Ba Sing Se | threat | Artifact creature threat.
- 1 Witch-king of Angmar | threat | Legendary black creature threat.
- 1 Austere Command | wipe | Flexible board-reset option.
- 1 Dusk // Dawn | wipe | Creature-focused board reset.
- 1 Vanquish the Horde | wipe | Efficient creature board reset.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $3926.59 to buy, $3926.59 the whole deck.

**Summary:** A high-power mono-blue artifact strategy that converts a dense field of cheap machines into mana, cards, and overwhelming board presence. The deck uses efficient disruption to protect its engine, then leverages cost reduction, recursion, tutoring, and artifact-based threats to seize control of the table and close quickly.

The quality model grades this deck below the precon baseline: the deck makes less mana on turn four than the norm, and that lowers the grade. many of the cards appear in no top list, and that lowers the grade. more opening hands hold two to four lands than the norm, and that raises the grade.

- [INFO] `curve_summary`: average mana value 3.18 over 66 nonland cards
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 3.18, and bracket 4 wants 1.6 to 3
- [WARN] `profile_off_band`: the mana available on turn four is 4.37, and bracket 4 wants 4.8 or more (mean over 10000 hands)
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 22 Island | land | Reliable blue sources for a mono-blue artifact mana base.
- 1 Academy Ruins | land | Recurs key artifacts through removal and trades.
- 1 Glimmervoid | land | Provides blue mana while the deck maintains a dense artifact board.
- 1 Inventors' Fair | land | Artifact-focused utility land with access to a needed artifact later.
- 1 Mishra's Workshop | land | Explosive acceleration for artifact spells.
- 1 Mystic Sanctuary | land | Returns an important instant or sorcery to the top of the library.
- 1 Otawara, Soaring City | land | Flexible uncounterable-style channel interaction from a land slot.
- 1 Power Depot | land | Artifact land that supports artifact casting.
- 1 Seat of the Synod | land | Blue-producing artifact land for affinity and artifact-count synergies.
- 1 Spire of Industry | land | Blue source that remains flexible with artifacts in play.
- 1 Urza's Saga | land | Produces constructs and finds a low-cost artifact utility piece.
- 1 Urza's Workshop | land | Artifact-focused utility mana source.
- 1 Mox Opal | ramp | Fast artifact-enabled mana production.
- 1 Moonsnare Prototype | ramp | Cheap acceleration that turns spare artifacts into mana.
- 1 Arc Reactor | ramp | Early artifact mana source.
- 1 Crystal Skull, Isu Spyglass | ramp | Mana rock with useful artifact utility.
- 1 Vedalken Engineer | ramp | Efficient creature-based mana for artifact spells.
- 1 Chief Engineer | ramp | Convoke converts artifact creatures into casting acceleration.
- 1 Grand Architect | ramp | Turns blue creatures into substantial artifact mana.
- 1 Metalworker | ramp | Provides explosive mana from an artifact-heavy hand.
- 1 Inspiring Statuary | ramp | Improvise lets inexpensive artifacts accelerate larger spells.
- 1 Slagstone Refinery | ramp | Artifact-based mana development.
- 1 Network Terminal | ramp | Fixes mana while looting through excess cards.
- 1 Powerstone Shard | ramp | Additional mana rock for powering large artifact turns.
- 1 Solar Array | ramp | Artifact mana source that advances the board.
- 1 Braided Net // Braided Quipu | draw | Artifact-based card advantage that supports the deck's central density.
- 1 Cerebral Download | draw | Efficient instant-speed card selection and advantage.
- 1 Era of Innovation | draw | Converts artifact development into sustained cards.
- 1 Esoteric Duplicator | draw | Clue-based card advantage with artifact synergy.
- 1 Forensic Gadgeteer | draw | Improves Clue efficiency and supplies ongoing value.
- 1 Reverse Engineer | draw | Refills the hand efficiently with improvise.
- 1 Riddlesmith | draw | Loots repeatedly as artifacts are deployed.
- 1 Sai, Master Thopterist | draw | Builds evasive artifact material and cashes it in for cards.
- 1 Thirst for Knowledge | draw | Instant-speed digging that rewards spare artifacts.
- 1 Thoughtcast | draw | Frequently becomes a very inexpensive draw spell.
- 1 Thought Monitor | draw | Affinity makes this a cheap artifact body that draws immediately.
- 1 Vedalken Archmage | draw | Provides a strong stream of cards from artifact casting.
- 1 Curator's Ward | interaction | Protects a key permanent while replacing itself.
- 1 Disruption Protocol | interaction | Artifact-assisted countermagic.
- 1 Escape Protocol | interaction | Protects artifacts and enables reusable enter-the-battlefield value.
- 1 Etched Champion | interaction | Resilient protection against creature-based pressure.
- 1 Ghostly Flicker | interaction | Saves important permanents and reuses artifact value.
- 1 Ice Out | interaction | Efficient counterspell in a snow-heavy basic mana base.
- 1 Metallic Rebuke | interaction | One of the deck's most efficient artifact-enabled counters.
- 1 Padeem, Consul of Innovation | interaction | Protects valuable artifacts while offering additional value.
- 1 Reality Ripple | interaction | Low-cost flexible disruption.
- 1 Stoic Rebuttal | interaction | Reliable artifact-enabled hard counterspell.
- 1 Welding Jar | interaction | Zero-cost protection for an essential artifact.
- 1 Aether Spellbomb | removal | Cheap, recurrable creature interaction.
- 1 Blasting Station | removal | Sacrifice outlet that converts artifact creatures into direct removal.
- 1 Cyber Conversion | removal | Efficient answer to a problematic creature or commander.
- 1 Into Thin Air | removal | Flexible tempo answer for troublesome permanents.
- 1 Resculpt | removal | Exiles an artifact or creature at instant speed.
- 1 Ravenform | removal | Exile-based answer that covers artifacts and creatures.
- 1 Unable to Scream | removal | Very cheap permanent-based creature shutdown.
- 1 Water Whip | removal | Low-cost answer for an opposing creature.
- 1 Lumengrid Drake | removal | Artifact-supported creature removal attached to a body.
- 1 Shape Anew | removal | Turns a small artifact into an answer or a major artifact deployment.
- 1 Engineered Explosives | wipe | Scalable reset that arrives early and can clear compact boards.
- 1 Hurkyl's Recall | wipe | Punishes opposing artifact boards and creates a strong tempo swing.
- 1 Emry, Lurker of the Loch | synergy | Recurs cheap artifacts and enables repeated artifact casting.
- 1 Etherium Sculptor | synergy | Reduces the cost of the deck's central spell type.
- 1 Foundry Inspector | synergy | Additional artifact cost reduction for explosive development.
- 1 Manifold Key | synergy | Untaps important mana artifacts or creates unblockable attacks.
- 1 Mystic Forge | synergy | Turns the top of the library into a continuing artifact resource.
- 1 Whir of Invention | synergy | Finds the exact artifact needed at instant speed.
- 1 Arcbound Crusher | threat | Grows rapidly as the artifact board develops.
- 1 Arcbound Reclaimer | threat | A modular threat that also preserves important artifacts.
- 1 Chrome Steed | threat | Efficient metalcraft attacker.
- 1 Darksteel Juggernaut | threat | Large indestructible attacker scaled by the artifact board.
- 1 Frogmite | threat | Affinity body that is often deployed with minimal mana.
- 1 Kappa Cannoneer | threat | A powerful evasive finisher that grows from artifact deployment.
- 1 Kuldotha Forgemaster | threat | Converts surplus artifacts into the deck's strongest permanents.
- 1 Lodestone Golem | threat | Pressures opponents while taxing nonartifact development.
- 1 Master Transmuter | threat | Protects artifacts and cheats expensive threats into play.
- 1 Myr Enforcer | threat | Large affinity attacker that is cheap in an established board.
- 1 Phyrexian Metamorph | threat | Copies the most powerful artifact or creature on the battlefield.
- 1 Traxos, Scourge of Kroog | threat | Large low-cost artifact attacker that naturally untaps during development.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $174.72 to buy, $174.72 the whole deck.

**Summary:** A welcoming Naya Dinosaur deck built around accelerating into large prehistoric creatures, rewarding creature deployment, and turning combat into the deck’s main source of pressure and advantage. It has ample mana fixing, several ways to rebuild its hand, flexible answers to troublesome permanents, and resilient protection for its most important board states.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. the commander is a popular one, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.42 over 62 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 1 Beast Whisperer | draw | Draws cards as Dinosaur creatures are cast.
- 1 Garruk's Uprising | draw | Provides creature-based card draw and supports large attackers.
- 1 Guardian Project | draw | Rewards the creature-heavy strategy with steady card draw.
- 1 Kutzil, Malamet Exemplar | draw | Draws from combat damage while protecting proactive turns.
- 1 Return of the Wildspeaker | draw | Offers either a large burst of cards or a finishing combat boost.
- 1 Ripjaw Raptor | draw | A Dinosaur that turns damage into cards.
- 1 Rishkar's Expertise | draw | Refills the hand based on a large Dinosaur's power.
- 1 Runic Armasaur | draw | Draws cards against common activated abilities.
- 1 Shamanic Revelation | draw | Refills the hand from a developed creature board.
- 1 Toski, Bearer of Secrets | draw | Rewards attacking with multiple creatures.
- 1 Vanquisher's Banner | draw | Supports the Dinosaur tribe while providing card draw.
- 1 Boros Charm | interaction | Protects the board or helps force through damage.
- 1 Heroic Intervention | interaction | Protects creatures and permanents from opposing disruption.
- 1 Lightning Greaves | interaction | Protects a key creature and gives it haste.
- 1 Snakeskin Veil | interaction | A low-cost protection spell for an important creature.
- 1 Swiftfoot Boots | interaction | Provides repeatable protection and haste.
- 1 Temple Altisaur | interaction | Reduces damage dealt to the Dinosaur board.
- 1 Arcane Signet | ramp | Reliable fixing for the three-color mana base.
- 1 Birds of Paradise | ramp | Early color fixing and acceleration.
- 1 Cultivate | ramp | Finds lands and ensures future land drops.
- 1 Drover of the Mighty | ramp | Fixes mana and becomes a sizable creature alongside Dinosaurs.
- 1 Farseek | ramp | Finds a dual land with a basic land type.
- 1 Nature's Lore | ramp | Efficiently finds a Forest dual land.
- 1 Rampant Growth | ramp | Simple early land acceleration and fixing.
- 1 Sol Ring | ramp | Efficient colorless acceleration.
- 1 Thunderherd Migration | ramp | Dinosaur-themed land fixing.
- 1 Topiary Stomper | ramp | A Dinosaur that ramps into larger threats.
- 1 Bronzebeak Foragers | removal | Exiles problematic opposing cards while advancing the board.
- 1 Burning Sun's Avatar | removal | Removes creatures and players' resources on entry.
- 1 Itzquinth, Firstborn of Gishath | removal | An efficient Dinosaur-based fight effect.
- 1 Needletooth Raptor | removal | Turns enrage into creature removal.
- 1 Ravenous Sailback | removal | A Dinosaur that removes an artifact or enchantment.
- 1 Savage Stomp | removal | Efficient creature removal that rewards Dinosaurs.
- 1 Thrashing Brontodon | removal | Flexible removal attached to a Dinosaur.
- 1 Belligerent Yearling | synergy | A low-cost Dinosaur that grows with the tribe.
- 1 Commune with Dinosaurs | synergy | Finds either a needed land or Dinosaur creature.
- 1 Dinosaur Egg | synergy | Supports the creature theme and replaces itself after dying.
- 1 Huatli's Raptor | synergy | Adds an early Dinosaur body and proliferates counters.
- 1 Kinjalli's Caller | synergy | Reduces the cost of large Dinosaur spells.
- 1 Marauding Raptor | synergy | Reduces Dinosaur costs and enables enrage triggers.
- 1 Otepec Huntmaster | synergy | Makes Dinosaur spells cheaper and grants haste.
- 1 Priest of the Wakening Sun | synergy | Finds Dinosaurs and gains life against aggressive decks.
- 1 Raptor Companion | synergy | An efficient early Dinosaur for tribal payoffs.
- 1 Raptor Hatchling | synergy | Creates more Dinosaur bodies through enrage.
- 1 Sky Terror | synergy | An evasive Dinosaur that pressures life totals early.
- 1 Sunfrill Imitator | synergy | Copies the combat strengths of larger Dinosaurs.
- 1 Territorial Hammerskull | synergy | Clears a blocker whenever it attacks.
- 1 Carnage Tyrant | threat | A resilient top-end Dinosaur threat.
- 1 Charging Tuskodon | threat | Makes Dinosaur combat damage especially punishing.
- 1 Ghalta and Mavren | threat | Creates a major combat payoff for attacking with large creatures.
- 1 Pantlaza, Sun-Favored | threat | Turns a stream of Dinosaurs into additional creature resources.
- 1 Quartzwood Crasher | threat | Builds trampling Dinosaur tokens from combat damage.
- 1 Regisaur Alpha | threat | Provides two Dinosaur bodies and gives the tribe haste.
- 1 Shifting Ceratops | threat | A difficult-to-answer hasty Dinosaur attacker.
- 1 Snapping Sailback | threat | A sturdy Dinosaur that can grow rapidly in combat.
- 1 Sun-Crested Pterodon | threat | An evasive Dinosaur that can become a strong attacker.
- 1 Thundering Spineback | threat | Creates Dinosaur tokens and boosts the team.
- 1 Zetalpa, Primal Dawn | threat | A powerful evasive finisher with multiple combat abilities.
- 1 Zilortha, Strength Incarnate | threat | Makes the deck's high-power creatures hit even harder.
- 1 Chain Reaction | wipe | Creature sweep that scales with crowded boards.
- 1 Raging Swordtooth | wipe | A Dinosaur-themed board sweeper that can trigger enrage.
- 1 Wrath of God | wipe | Straightforward answer to an overwhelming creature board.
- 8 Forest | land | Basic green sources for ramp spells and green-heavy Dinosaur cards.
- 5 Mountain | land | Basic red sources for Dinosaur threats and removal.
- 4 Plains | land | Basic white sources for protection and utility spells.
- 1 Battlefield Forge | land | Flexible red or white mana source.
- 1 Bountiful Promenade | land | Reliable green or white multiplayer land.
- 1 Brushland | land | Flexible green or white mana source.
- 1 Canopy Vista | land | Forest and Plains dual land for fixing.
- 1 Cinder Glade | land | Mountain and Forest dual land for fixing.
- 1 City of Brass | land | Produces any needed color.
- 1 Clifftop Retreat | land | Red and white fixing.
- 1 Command Tower | land | Reliable three-color commander mana.
- 1 Exotic Orchard | land | Usually supplies the needed colors in multiplayer.
- 1 Fortified Village | land | Green and white fixing.
- 1 Game Trail | land | Red and green fixing.
- 1 Jetmir's Garden | land | Three-color fixing with basic land types.
- 1 Jungle Shrine | land | Three-color fixing for the deck's identity.
- 1 Karplusan Forest | land | Flexible red or green mana source.
- 1 Path of Ancestry | land | Tribal mana fixing with occasional card selection.
- 1 Rockfall Vale | land | Red and green fixing.
- 1 Rootbound Crag | land | Red and green fixing.
- 1 Sacred Foundry | land | Mountain and Plains dual land for fixing.
- 1 Stomping Ground | land | Mountain and Forest dual land for fixing.
- 1 Sunpetal Grove | land | Green and white fixing.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $225.66 the whole deck.

**Summary:** A mono-white blink deck built around protecting a legendary creature, repeatedly reusing enter-the-battlefield effects, and accruing value through creatures and equipment. It establishes mana efficiently, controls troublesome permanents with broad answers, and wins through an evasive, resilient creature board backed by protective spells and sweepers.

The quality model grades this deck good: the commander does not place in cEDH events, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade. the mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.97 over 63 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 25 Plains | land | Reliable white mana base for a mono-white creature deck.
- 1 Command Tower | land | Efficient colored mana source.
- 1 Exotic Orchard | land | Flexible colored mana source in multiplayer games.
- 1 City of Brass | land | Provides white mana without entering tapped.
- 1 Plaza of Heroes | land | Supports the legendary commander while producing colored mana.
- 1 Secluded Courtyard | land | Names Human to cast much of the creature suite.
- 1 Spire of Industry | land | Colored mana source supported by the artifact package.
- 1 Marsh Flats | land | Finds a Plains while improving land access.
- 1 Fabled Passage | land | Finds a Plains and fixes early land drops.
- 1 Ash Barrens | land | Basic landcycling provides dependable access to Plains.
- 1 Minas Tirith | land | White-producing utility land with later card advantage.
- 1 War Room | land | Repeatable card advantage from a land slot.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Arcane Signet | ramp | Reliable colored mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost multiplayer mana fixing.
- 1 Wayfarer's Bauble | ramp | Finds a Plains and permanently advances mana.
- 1 Sword of the Animist | ramp | Turns creature attacks into repeatable Plains ramp.
- 1 Thought Vessel | ramp | Mana acceleration with a useful hand-size benefit.
- 1 White Lotus Tile | ramp | Low-cost artifact acceleration.
- 1 Springleaf Drum | ramp | Converts inexpensive creatures into mana.
- 1 Bender's Waterskin | ramp | Artifact-based mana acceleration.
- 1 Commander's Sphere | ramp | Fixes mana early and can later become a card.
- 1 Banishing Light | removal | Versatile answer to problematic nonland permanents.
- 1 Crib Swap | removal | Exiles creatures at instant speed.
- 1 Destroy Evil | removal | Efficient answer to large creatures or key enchantments.
- 1 Dispatch | removal | Cheap creature exile supported by the artifact count.
- 1 Generous Gift | removal | Answers any permanent type.
- 1 Get Lost | removal | Efficient answer to creatures, enchantments, and planeswalkers.
- 1 Journey to Nowhere | removal | Low-cost creature exile effect.
- 1 March of Otherworldly Light | removal | Flexible instant-speed exile removal.
- 1 Swords to Plowshares | removal | Premium one-mana creature exile.
- 1 Austere Command | wipe | Flexible sweeper that can preserve the preferred board.
- 1 Dusk // Dawn | wipe | Sweeps larger creatures and can recover the creature suite.
- 1 Vanquish the Horde | wipe | Efficient creature reset when opposing boards grow.
- 1 Adventurer's Airship | draw | Creature-based card advantage that benefits from a board presence.
- 1 Buster Sword | draw | Equipment that provides ongoing card advantage.
- 1 Champions of Minas Tirith | draw | Creature-based card advantage for the Human shell.
- 1 Crown of Gondor | draw | Equipment that rewards combat with cards.
- 1 Diary of Dreams | draw | Artifact source of repeatable card advantage.
- 1 Idol of Oblivion | draw | Low-cost repeatable card draw.
- 1 Lembas | draw | Early card selection that can later replace itself.
- 1 Mask of Memory | draw | Combat-based filtering and card advantage.
- 1 Skullclamp | draw | Efficient card advantage with disposable creatures.
- 1 Tome of Legends | draw | Steady card draw alongside the commander.
- 1 Puresteel Paladin | draw | Draw engine for the equipment package.
- 1 Clever Concealment | interaction | Protects the board from removal and sweepers.
- 1 Darksteel Plate | interaction | Makes an important creature resilient to destruction.
- 1 Lightning Greaves | interaction | Provides haste and protective shroud.
- 1 Reprieve | interaction | Temporarily answers a spell while replacing itself.
- 1 Sheltered by Ghosts | interaction | Protects a creature while pressuring opposing permanents.
- 1 Swiftfoot Boots | interaction | Protects the commander while allowing immediate activation.
- 1 Ultimate Magic: Holy | interaction | Flexible protection for the creature board.
- 1 Unbreakable Formation | interaction | Protects the board and can turn creatures into a finishing force.
- 1 Flickerwisp | synergy | Reusable blink payoff that resets creatures and permanents.
- 1 Angel of Condemnation | synergy | Provides repeatable blinking and can temporarily exile opposing creatures.
- 1 Personify | synergy | Blink effect that reuses enter-the-battlefield abilities.
- 1 Slip On the Ring | synergy | Instant-speed blink protects a creature and retriggers it.
- 1 Angel of Sanctions | synergy | Blinking renews its versatile exile trigger.
- 1 Fiend Hunter | synergy | Blinking can repeatedly remove opposing creatures.
- 1 Palace Jailer | synergy | Blinking refreshes the monarch and exile effects.
- 1 Wall of Omens | synergy | A cheap defensive body with a reusable entry trigger.
- 1 Inspiring Overseer | synergy | Blinking repeatedly provides a card and life.
- 1 Gift of Immortality | synergy | Keeps a key creature available through removal and sweeps.
- 1 Angel of Serenity | threat | Large evasive finisher with a powerful entry trigger.
- 1 Giada, Font of Hope | threat | Develops the Angel board while growing later Angels.
- 1 Exemplar of Light | threat | Evasive creature that contributes to sustained board pressure.
- 1 Bronze Guardian | threat | Protective artifact threat that scales with the board.
- 1 Boromir, Warden of the Tower | threat | Disruptive legendary creature that protects the board from free spells.
- 1 Faramir, Field Commander | threat | Builds a Human army and rewards attacking.
- 1 Frontline Medic | threat | Combat-focused creature that can protect an attacking force.
- 1 Zack Fair | threat | Legendary attacker that advances the creature plan.
- 1 Westfold Rider | threat | Creature pressure with useful battlefield utility.
- 1 Troop of Ponies | threat | Develops a broad creature board for combat and support effects.
- 1 The Walls of Ba Sing Se | threat | Durable artifact creature that contributes to the board plan.
- 1 Summon: Primal Garuda | threat | Evasive saga creature that supplies a meaningful finishing body.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $416.38 to buy, $416.38 the whole deck.

**Summary:** This blue-red tempo deck establishes early creature pressure with Ledger Shredder, Faerie Mastermind, and Ragavan, Nimble Pilferer, then uses efficient draw, removal, and interaction to keep opponents from stabilizing. Temporal Mastery and Temporal Trespass provide a powerful payoff when the game develops beyond the opening exchanges. The deck gives up larger standalone threats for a lean spell package and a proactive, disruptive game plan.

The quality model grades this deck good: the cards pair in ways the top lists do not, and that lowers the grade. the deck runs few spells as one copy, and that raises the grade. the deck holds many cheap spells, and that raises the grade.

- [INFO] `curve_summary`: average mana value 2.33 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Island | land | Provides a reliable blue source for the deck's core spells.
- 6 Mountain | land | Provides red mana for removal and the red creature package.
- 4 Steam Vents | land | Supplies both deck colors without entering tapped.
- 4 Spirebluff Canal | land | Provides an untapped blue-red source for early tempo turns.
- 4 Consider | draw | Cheap card selection supports a spell-dense tempo plan.
- 2 Preordain | draw | Adds efficient selection to find lands, pressure, or answers.
- 2 Counterspell | interaction | Provides broad stack interaction for protecting a lead.
- 2 Force of Negation | interaction | Offers additional interaction while maintaining tempo.
- 2 Spell Pierce | interaction | Efficient early interaction against opposing noncreature plays.
- 4 Lightning Bolt | removal | Low-cost removal supports an aggressive tempo posture.
- 2 Into the Flood Maw | removal | Flexible cheap removal helps clear the way for attackers.
- 2 Abrade | removal | Adds removal coverage for problematic opposing permanents.
- 2 Temporal Mastery | synergy | Supports the deck's temporal-spell synergy package.
- 2 Temporal Trespass | synergy | Provides a second temporal payoff alongside Temporal Mastery.
- 4 Ledger Shredder | draw | A low-cost creature that supplies ongoing card selection.
- 4 Faerie Mastermind | draw | Adds evasive creature pressure while contributing card draw.
- 4 Ragavan, Nimble Pilferer | ramp | Provides an early creature that supports the deck's resource development.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $38.08 to buy, $38.08 the whole deck.

**Summary:** This mono-red burn deck keeps the mana simple, applies pressure with a steady stream of creatures, and backs that pressure with direct-damage removal. Its draw cards help sustain action, while the synergy creatures reward the deck for staying focused on cheap red spells. It aims to finish games through repeated burn pressure and combat damage, giving up flexibility and broader answers in exchange for a direct, streamlined plan.

The quality model grades this deck below the precon baseline: the deck runs its spells as playsets, and that lowers the grade. the cards pair in ways the top lists do not, and that lowers the grade. many of the cards appear in no top list, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 2.78 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides a consistent mono-red mana base.
- 4 Lightning Bolt | removal | Efficient burn removal supports the deck's direct-damage plan.
- 2 Lightning Strike | removal | Adds reliable direct-damage removal.
- 4 Risk Factor | draw | Provides the deck with a dedicated draw package.
- 2 Ancestral Anger | draw | Rounds out the draw suite while fitting the burn shell.
- 2 Chandra, Dressed to Kill | ramp | Supplies the requested ramp component.
- 4 Eidolon of the Great Revel | synergy | A burn-focused synergy piece that pressures the opponent's plan.
- 4 Thermo-Alchemist | synergy | Works with the deck's large instant and sorcery package.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | Provides an early red threat for the pressure plan.
- 4 Keral Keep Disciples | threat | Adds a full set of red threats to maintain board pressure.
- 4 Hazoret the Fervent | threat | Serves as a resilient top-end threat.
- 2 Torbran, Thane of Red Fell | threat | Completes the threat package with a powerful red finisher.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $739.34 to buy, $739.34 the whole deck.

**Summary:** This is a proactive white-black lifegain deck that starts with Soul Warden and Voice of the Blessed, then uses its steady life total gains to develop an increasingly dangerous creature board. Enduring Innocence and Inspiring Overseer keep the pressure supplied, while Path to Exile and Solitude clear away opposing creatures. Attended Healer, Sheoldred, the Apocalypse, Archangel of Thune, Enduring Tenacity, and Bloodthirsty Conqueror provide the closing power; the tradeoff is that the deck leans heavily on creatures and its more expensive finishers can be slower against very fast opponents.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. the cards pair in ways the top lists do not, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.00 over 36 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 10 Plains | land | Provides reliable white mana without entering tapped.
- 4 Swamp | land | Provides reliable black mana without entering tapped.
- 4 Godless Shrine | land | Supplies both colors for the white-black mana base.
- 4 Caves of Koilos | land | Supplies both colors while supporting early plays.
- 2 Shattered Sanctum | land | Rounds out the two-color land base.
- 2 Arcane Signet | ramp | Provides the requested mana acceleration for the deck's higher-cost threats.
- 4 Enduring Innocence | draw | Keeps cards flowing alongside the deck's creature-heavy plan.
- 2 Inspiring Overseer | draw | Adds card flow on a lifegain-themed creature body.
- 4 Path to Exile | removal | Efficiently answers opposing creatures.
- 2 Solitude | removal | Provides additional creature removal while remaining a threat in longer games.
- 4 Soul Warden | synergy | A low-cost lifegain engine that turns on the deck's payoffs.
- 4 Voice of the Blessed | synergy | Rewards repeated lifegain and gives the deck an early scaling payoff.
- 4 Attended Healer | threat | Builds a board presence from the deck's lifegain plan.
- 4 Sheoldred, the Apocalypse | threat | A resilient top-end threat that pressures opponents while supporting lifegain.
- 2 Archangel of Thune | threat | A powerful lifegain payoff that helps convert small advantages into a winning board.
- 2 Enduring Tenacity | threat | Adds another threatening payoff for the deck's lifegain plan.
- 2 Bloodthirsty Conqueror | threat | Provides a high-impact finisher for games that go long.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $186.82 to buy, $186.82 the whole deck.

**Summary:** This black-green midrange deck develops with early ramp, then presses the game with a layered creature suite while using its removal to clear opposing obstacles. Its draw cards and creature-focused synergy help it sustain pressure into longer games, and its larger threats provide the primary way to finish. The deck gives up sideboard flexibility and relies on its main-deck mix of threats, removal, and card flow to cover a broad FNM field.

The quality model grades this deck below the precon baseline: few of the cards are ones the top lists of the format play, and that lowers the grade. the mana base serves the colors evenly, and that lowers the grade. many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.19 over 36 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 8 Forest | land | Provides the green mana base for the deck.
- 8 Swamp | land | Provides the black mana base for the deck.
- 4 Overgrown Tomb | land | Provides both deck colors in the land base.
- 4 Blooming Marsh | land | Provides black-green color support.
- 2 Llanowar Elves | ramp | Fills the early ramp slots.
- 4 Phyrexian Arena | draw | Provides a dedicated draw package.
- 2 Unholy Annex // Ritual Chamber | draw | Adds to the deck's draw package.
- 4 Bitter Triumph | removal | Forms the efficient core of the removal suite.
- 2 Maelstrom Pulse | removal | Complements the removal suite.
- 4 Midnight Reaper | synergy | Supports the deck's creature-focused synergy plan.
- 4 Undying Malice | synergy | Supports the deck's creature-focused synergy plan.
- 4 Darkstar Augur | threat | Serves as an early threat for the midrange plan.
- 4 Goldvein Hydra | threat | Serves as a substantial green threat.
- 3 Rottenmouth Viper | threat | Serves as a black threat in the creature suite.
- 3 Vaultborn Tyrant | threat | Provides a top-end threat for closing games.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $117.74 to buy, $117.74 the whole deck.

**Summary:** A red-white aggressive strategy built to establish early creature pressure, amplify combat with coordinated battlefield payoffs, and maintain momentum through flexible answers and resource-generating attackers. Its mana base emphasizes dependable access to both colors while preserving a fast, attack-oriented game plan.

The quality model grades this deck below the precon baseline: few of the cards are ones the top lists of the format play, and that lowers the grade. many of the cards appear in no top list, and that lowers the grade. the color sources cover the pips, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.06 over 36 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 4 Fugitive Codebreaker | draw | Early red pressure that also replenishes resources.
- 2 Reckless Lackey | draw | Low-cost aggressive creature that provides card flow.
- 4 Boros Charm | interaction | Flexible protection and reach for an aggressive Boros plan.
- 2 Swiftfoot Boots | interaction | Protects key attackers and lets threats contribute quickly.
- 4 Emeritus of Truce // Swords to Plowshares | removal | Efficient white answer that remains useful alongside creature pressure.
- 4 Aetherjacket | removal | Removal attached to a board-developing artifact creature.
- 4 Warleader's Call | synergy | Rewards a wide battlefield and turns repeated attacks into extra damage.
- 4 Dragonback Lancer | threat | Red creature threat that keeps the deck's curve focused on attacking.
- 4 Vanguard Seraph | threat | Evasive white threat for pressuring stalled boards.
- 4 Serra Redeemer | threat | Resilient white finisher that supports the creature-heavy attack plan.
- 4 Sacred Foundry | land | Reliable untapped access to both primary colors.
- 4 Inspiring Vantage | land | Fast dual land for early aggressive development.
- 4 Sundown Pass | land | Dual-color source that supports both halves of the deck.
- 4 Sunbillow Verge | land | Additional red-white fixing without slowing the deck's early development.
- 4 Plains | land | Stable white source for the deck's white spells.
- 4 Mountain | land | Stable red source for the deck's red spells.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $726.80 to buy, $726.80 the whole deck.

**Summary:** Karlov leads a white-black sacrifice deck that develops creatures, converts them into mana, cards, removal, and pressure, then uses durable synergy pieces to keep value flowing. The deck wins by building a threatening board and turning repeated creature sacrifices into steady advantages while its larger creatures finish the game. It gives up some speed to maintain a broad creature-based engine and relies on its board to make the sacrifice plan work.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade. few decks lead with the commander, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.17 over 63 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 13 Plains | land | Basic white land for the mana base.
- 15 Swamp | land | Basic black land for the mana base.
- 1 Command Tower | land | Two-color land for the mana base.
- 1 Godless Shrine | land | White-black land for the mana base.
- 1 Scrubland | land | White-black land for the mana base.
- 1 Caves of Koilos | land | White-black land for the mana base.
- 1 Vault of Champions | land | White-black land for the mana base.
- 1 City of Brass | land | Flexible colored land for the mana base.
- 1 Mana Confluence | land | Flexible colored land for the mana base.
- 1 Exotic Orchard | land | Flexible colored land for the mana base.
- 1 Sol Ring | ramp | Required efficient ramp for deploying the deck's engine and threats.
- 1 Ashnod's Altar | ramp | Sacrifice-focused ramp that supports the central plan.
- 1 Phyrexian Altar | ramp | Sacrifice-focused ramp that supports the central plan.
- 1 Pitiless Plunderer | ramp | Creature-based ramp that fits the sacrifice shell.
- 1 Priest of Forgotten Gods | ramp | Sacrifice-focused ramp for the creature-heavy plan.
- 1 Pawn of Ulamog | ramp | Creature-based ramp that rewards the sacrifice plan.
- 1 Sifter of Skulls | ramp | Creature-based ramp that rewards creature sacrifices.
- 1 Skullport Merchant | ramp | Creature-based ramp that supports the sacrifice shell.
- 1 Crowded Crypt | ramp | Ramp piece for the deck's creature-focused plan.
- 1 Culling the Weak | ramp | Efficient sacrifice-focused ramp.
- 1 Corrupted Conviction | draw | Sacrifice-focused card draw.
- 1 Village Rites | draw | Efficient sacrifice-focused card draw.
- 1 Vampiric Rites | draw | Repeatable sacrifice-focused card draw.
- 1 Disciple of Bolas | draw | Creature-based draw that fits the sacrifice plan.
- 1 Smothering Abomination | draw | Creature-based draw for a sacrifice-heavy board.
- 1 Shadowheart, Dark Justiciar | draw | Creature-based draw that supports sacrificing creatures.
- 1 Relic Vial | draw | Artifact draw for the creature-focused shell.
- 1 Baron Bertram Graywater | draw | Creature-based draw for the deck's white-black core.
- 1 Bushmeat Poacher | draw | Creature-based draw that supports sacrificing creatures.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | Draw threat that fits the sacrifice-oriented creature plan.
- 1 Tevesh Szat, Doom of Fools | draw | Draw engine for the creature-focused strategy.
- 1 Cartel Aristocrat | interaction | Sacrifice-oriented interaction that protects the board plan.
- 1 Fanatical Devotion | interaction | Sacrifice-oriented interaction for protecting key creatures.
- 1 Flare of Fortitude | interaction | Protective interaction for the board-based strategy.
- 1 Gift of Doom | interaction | Protective interaction that fits the sacrifice shell.
- 1 Promise of Tomorrow | interaction | Interaction that supports the creature-focused plan.
- 1 Spirit Bonds | interaction | Creature-focused interaction for the deck's board plan.
- 1 Sunstone | interaction | Defensive interaction for buying time to assemble the engine.
- 1 Nightmare Shepherd | interaction | Creature-based interaction that supports the sacrifice strategy.
- 1 Bone Shards | removal | Low-cost removal that fits a sacrifice-heavy deck.
- 1 Bone Splinters | removal | Low-cost removal that fits a sacrifice-heavy deck.
- 1 Attrition | removal | Sacrifice-focused repeatable removal.
- 1 Ayli, Eternal Pilgrim | removal | Creature-based removal that fits the white-black sacrifice plan.
- 1 Eaten Alive | removal | Sacrifice-focused removal for problematic permanents.
- 1 Dictate of Erebos | removal | Sacrifice-focused removal payoff.
- 1 Grave Pact | removal | Sacrifice-focused removal payoff.
- 1 Teysa, Orzhov Scion | removal | Creature-based removal that fits the sacrifice shell.
- 1 Yawgmoth, Thran Physician | removal | Creature-based removal for the sacrifice-oriented plan.
- 1 Carrion Feeder | synergy | Low-cost sacrifice synergy for the deck's core plan.
- 1 Viscera Seer | synergy | Low-cost sacrifice synergy for the deck's core plan.
- 1 Zulaport Cutthroat | synergy | Creature-sacrifice synergy that helps convert losses into pressure.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | White-black creature-sacrifice synergy.
- 1 Bastion of Remembrance | synergy | Persistent sacrifice synergy for the central plan.
- 1 Fleshtaker | synergy | Creature-based sacrifice synergy.
- 1 Woe Strider | synergy | Creature-based sacrifice synergy with lasting board presence.
- 1 Hidden Stockpile | synergy | Enchantment-based sacrifice synergy.
- 1 Victimize | synergy | Sacrifice-focused synergy that sustains the creature plan.
- 1 Chthonian Nightmare | synergy | Recurring sacrifice synergy for the creature-heavy shell.
- 1 Liesa, Forgotten Archangel | threat | Powerful white-black creature threat for closing games.
- 1 Felisa, Fang of Silverquill | threat | White-black creature threat that fits Karlov's board plan.
- 1 Mondrak, Glory Dominus | threat | High-impact creature threat for a token-oriented sacrifice board.
- 1 Ratadrabik of Urborg | threat | White-black legendary creature threat for the board plan.
- 1 Requiem Angel | threat | Creature threat that supports a creature-focused strategy.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | Creature-based threat for the sacrifice shell.
- 1 Razaketh, the Foulblooded | threat | Major creature threat at the top of the curve.
- 1 Ghoulcaller Gisa | threat | Creature threat for building a sacrifice-oriented board.
- 1 Sadistic Hypnotist | threat | Creature threat that leverages the sacrifice plan.
- 1 Sidisi, Undead Vizier | threat | Creature threat that fits the sacrifice-heavy shell.
- 1 Vindictive Vampire | threat | Creature threat for a vampire-leaning sacrifice board.
- 1 Vito's Inquisitor | threat | Low-cost vampire threat for the white-black creature plan.
- 1 Toxic Deluge | wipe | Efficient board wipe when the table gets ahead.
- 1 Austere Command | wipe | Flexible board wipe for resetting difficult boards.
- 1 The Meathook Massacre | wipe | Board wipe that fits the deck's creature-loss theme.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $62.51 to buy, $62.51 the whole deck.

**Summary:** This deck develops a wide white creature board, using token-focused permanents and support cards to turn Adeline-led attacks into mounting pressure. It wins by maintaining a stream of threats, improving the value of its token army, and attacking through a clear board after targeted removal or a reset. The tradeoff is a deliberately budget-conscious mana base and a slower, board-centric plan that depends on establishing creatures and keeping them in play.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. the deck makes less mana on turn four than the norm, and that lowers the grade. more opening hands hold two to four lands than the norm, and that raises the grade.

- [INFO] `curve_summary`: average mana value 3.27 over 62 nonland cards

<details><summary>The deck list</summary>

- 25 Plains | land | Provides dependable white mana across the mana base.
- 1 Ancient Den | land | A white-producing land slot for the mana base.
- 1 Castle Ardenvale | land | A white land that supports the deck's token-focused plan.
- 1 Command Tower | land | Provides commander-aligned mana.
- 1 Eiganjo, Seat of the Empire | land | A white land slot with added utility.
- 1 Exotic Orchard | land | A flexible land for producing needed mana.
- 1 Fountainport | land | A utility land that fits the token strategy.
- 1 Heap Gate | land | A budget utility land for the mana base.
- 1 Jasmine Dragon Tea Shop | land | A land slot that supports the deck's white mana base.
- 1 Kjeldoran Outpost | land | A utility land suited to a token-focused deck.
- 1 Legion's Landing // Adanto, the First Fort | land | A token-oriented land slot for the deck.
- 1 Minas Tirith | land | A white legendary land that adds utility.
- 1 Mirrex | land | A utility land that contributes to the token plan.
- 1 Angelic Sell-Sword | draw | A low-cost draw option.
- 1 Bygone Bishop | draw | Adds draw support while fitting the creature plan.
- 1 Dawn of Hope | draw | Provides repeatable draw support for the deck.
- 1 Faramir, Field Commander | draw | A creature-based source of card advantage.
- 1 Glimmer Seeker | draw | A budget draw creature.
- 1 Idol of Oblivion | draw | A token-friendly artifact draw engine.
- 1 Nexus of Becoming | draw | Provides additional card advantage.
- 1 Search the Premises | draw | An enchantment source of draw support.
- 1 Staff of the Storyteller | draw | A budget artifact that supports drawing cards.
- 1 Valiant Rescuer | draw | A creature that contributes card advantage.
- 1 Wedding Announcement // Wedding Festivity | draw | Supports both card advantage and the token-focused game plan.
- 1 Ainok Strike Leader | interaction | Provides creature-based interaction.
- 1 Basri Ket | interaction | A flexible planeswalker interaction piece.
- 1 Basri, Tomorrow's Champion | interaction | Supports the board while providing interaction.
- 1 Lena, Selfless Champion | interaction | Helps protect the token board.
- 1 Rootborn Defenses | interaction | A budget defensive interaction spell for a wide board.
- 1 Spirit Bonds | interaction | An enchantment-based protection option.
- 1 Currency Converter | ramp | A cheap artifact ramp piece.
- 1 Druidic Satchel | ramp | Provides incremental mana development.
- 1 Goldvein Pick | ramp | An equipment-based ramp option.
- 1 Idol of False Gods | ramp | A low-cost ramp artifact.
- 1 Karn, Living Legacy | ramp | A planeswalker that supports mana development.
- 1 Keeper of the Accord | ramp | A creature-based ramp option.
- 1 Monologue Tax | ramp | An enchantment ramp piece for longer games.
- 1 Noble's Purse | ramp | A budget artifact mana source.
- 1 Prying Blade | ramp | An inexpensive equipment-based ramp option.
- 1 The Restoration of Eiganjo // Architect of Restoration | ramp | A ramp card that also fits the creature-focused plan.
- 1 Aerial Assault | removal | A low-cost removal spell.
- 1 Banishing Slash | removal | A budget sorcery removal option.
- 1 Citizen's Crowbar | removal | An equipment-based removal piece.
- 1 Generous Gift | removal | Flexible instant-speed removal.
- 1 Kellan's Lightblades | removal | A cheap removal spell.
- 1 Skyclave Apparition | removal | Creature-based removal that advances the board.
- 1 Stroke of Midnight | removal | Flexible instant-speed removal.
- 1 Ceaseless Conflict | wipe | A budget board-reset option.
- 1 Crisis of Conscience | wipe | Provides another efficient reset button.
- 1 Hour of Reckoning | wipe | A token-oriented board wipe.
- 1 Aligned Heart | synergy | A low-cost synergy piece for the token plan.
- 1 Anafenza, Unyielding Lineage | synergy | Supports the deck's creature and token theme.
- 1 Automated Assembly Line | synergy | An artifact synergy piece for building a board.
- 1 Cat Collector | synergy | A cheap creature that supports token synergies.
- 1 Cathar's Call | synergy | An enchantment that reinforces the token plan.
- 1 Clarion Spirit | synergy | A low-cost creature for token production synergies.
- 1 Divine Visitation | synergy | A powerful enchantment payoff for making tokens.
- 1 Eyes in the Skies | synergy | A token-focused spell that fits the main plan.
- 1 Hero of Precinct One | synergy | A creature that contributes to building a token board.
- 1 Intangible Virtue | synergy | A dedicated anthem-style token payoff.
- 1 Mavren Fein, Dusk Apostle | synergy | A creature that supports attacking with tokens.
- 1 Retrofitter Foundry | synergy | A flexible artifact token engine.
- 1 Rosie Cotton of South Lane | synergy | A creature payoff for the deck's token production.
- 1 Ajani's Chosen | threat | A token-oriented creature threat.
- 1 Archon of Sun's Grace | threat | An evasive token-focused threat.
- 1 Attended Healer | threat | A low-cost creature threat that fits the theme.
- 1 Basri's Lieutenant | threat | A creature threat that rewards building the board.
- 1 Cemetery Protector | threat | A resilient creature threat with useful board presence.
- 1 Defiler of Faith | threat | A creature threat for the white permanent plan.
- 1 Dragonback Lancer | threat | A budget creature threat.
- 1 Drogskol Cavalry | threat | A token-oriented flying threat.
- 1 Emeria Angel | threat | A creature threat that supports making tokens.
- 1 Gideon, Ally of Zendikar | threat | A planeswalker threat suited to a wide board.
- 1 God-Eternal Oketra | threat | A durable creature threat for a creature-heavy deck.
- 1 Hero of Bladehold | threat | An attack-focused creature threat that fits the token plan.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $174.32 the whole deck.

**Summary:** This Karlov deck develops a lifegain-focused board, uses protective interaction to keep its important creatures in play, and steadily builds toward attacks from Angels and other creature threats. Broad removal and a small suite of reset cards keep opposing boards manageable while Equipment and artifact support provide staying power. It gives up the explosive speed of the strongest artifact-heavy commander shells for a more creature-driven game that can pivot between pressure, protection, and battlefield control.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade. few decks lead with the commander, and that raises the grade.

Summary rules claims (F-26): the format play

- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.03 over 63 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 1 Command Tower | land | Reliable access to the deck's two colors.
- 1 City of Brass | land | Flexible color fixing for early plays.
- 1 Exotic Orchard | land | Adds another flexible source of colored mana.
- 1 Path of Ancestry | land | Provides fixing in the mana base.
- 1 Plaza of Heroes | land | Supports the commander-centered plan while fixing mana.
- 1 Secluded Courtyard | land | Provides tribal-oriented color fixing for the creature base.
- 1 Spire of Industry | land | Provides flexible mana alongside the deck's artifacts.
- 1 Takenuma, Abandoned Mire | land | A black source with additional utility.
- 15 Plains | land | Core white mana for the deck's lifegain and Angel cards.
- 13 Swamp | land | Core black mana for the deck's support and removal cards.
- 1 Sol Ring | ramp | Efficiently accelerates the deck's setup.
- 1 Arcane Signet | ramp | Reliable two-color mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost fixing and acceleration.
- 1 Commander's Sphere | ramp | Fixes mana while remaining useful later.
- 1 Relic of Legends | ramp | Adds mana acceleration for the legendary-focused deck.
- 1 Thought Vessel | ramp | Provides durable mana acceleration.
- 1 Sword of the Animist | ramp | An Equipment-based source of continued mana development.
- 1 Wayfarer's Bauble | ramp | Early mana development from a compact artifact.
- 1 Chromatic Lantern | ramp | Smooths the mana base while accelerating.
- 1 Inspiring Statuary | ramp | Turns the artifact package into additional mana support.
- 1 Buster Sword | draw | Equipment-based card advantage for the creature plan.
- 1 Call of the Ring | draw | A persistent source of card advantage.
- 1 Idol of Oblivion | draw | Low-cost artifact card advantage.
- 1 Inspiring Overseer | draw | A creature that contributes card advantage.
- 1 Lembas | draw | A compact artifact source of card advantage.
- 1 Mask of Memory | draw | Rewards the deck for attacking with creatures.
- 1 Night's Whisper | draw | Efficient black card advantage.
- 1 Puresteel Paladin | draw | Supports the Equipment package with card advantage.
- 1 Skullclamp | draw | Efficient Equipment-based card advantage.
- 1 Tome of Legends | draw | Steady card advantage for a commander-focused deck.
- 1 Wall of Omens | draw | Early board presence that replaces itself.
- 1 Banishing Light | removal | Flexible answer to a troublesome permanent.
- 1 Bitter Triumph | removal | Efficient answer to a problematic creature or planeswalker.
- 1 Crib Swap | removal | Creature removal that fits the deck's colors.
- 1 Dispatch | removal | Low-cost targeted removal.
- 1 Fatal Push | removal | Efficient black creature removal.
- 1 Generous Gift | removal | Broad permanent removal.
- 1 Get Lost | removal | Versatile answer to opposing permanents.
- 1 Infernal Grasp | removal | Straightforward creature removal.
- 1 Stroke of Midnight | removal | Flexible instant-speed permanent removal.
- 1 Bastion Protector | interaction | Helps preserve the commander on the battlefield.
- 1 Champion's Helm | interaction | Protects a key legendary creature.
- 1 Clever Concealment | interaction | Protects the developed board from opposing disruption.
- 1 Darksteel Plate | interaction | A durable protection tool for a key creature.
- 1 Lightning Greaves | interaction | Efficient protection for Karlov or a major threat.
- 1 Reprieve | interaction | Provides flexible stack interaction.
- 1 Swiftfoot Boots | interaction | Protects important creatures while supporting attacks.
- 1 Unbreakable Formation | interaction | Protects the creature board at a pivotal moment.
- 1 Aerith Gainsborough | synergy | A lifegain-focused synergy piece for the commander plan.
- 1 Angel of Vitality | synergy | A creature that supports the lifegain theme.
- 1 Aettir and Priwen | synergy | Equipment support for the deck's creature-focused plan.
- 1 Compassionate Healer | synergy | A dedicated lifegain synergy creature.
- 1 Elixir | synergy | A compact artifact supporting the deck's lifegain theme.
- 1 Kor Firewalker | synergy | A resilient creature aligned with the lifegain strategy.
- 1 Light of Promise | synergy | Rewards the deck's lifegain-focused development.
- 1 Night Nurse, Healer of Heroes | synergy | A lifegain-themed legendary support creature.
- 1 Rosie Cotton of South Lane | synergy | Supports the deck's token and lifegain-oriented pieces.
- 1 Second Breakfast | synergy | A flexible lifegain synergy card.
- 1 Angel of Invention | threat | A high-impact white creature threat.
- 1 Dawnhand Eulogist | threat | A creature threat that advances the board.
- 1 Foggy Swamp Hunters | threat | A creature threat for applying battlefield pressure.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A black creature threat with added flexibility.
- 1 Lyra Dawnbringer | threat | A powerful Angel threat for closing games.
- 1 Minwu, White Mage | threat | A legendary white creature threat.
- 1 Rabaroo Troop | threat | A creature threat that helps build a meaningful board.
- 1 Reaping Willow | threat | A substantial creature threat for the midgame.
- 1 Rooftop Percher | threat | A creature threat that adds pressure to the board.
- 1 Shattered Angel | threat | An Angel threat aligned with the lifegain theme.
- 1 Sneering Shadewriter | threat | A black creature threat for the deck's curve.
- 1 Victory's Herald | threat | A top-end Angel threat for finishing games.
- 1 Austere Command | wipe | A flexible reset that can spare the most useful parts of the board.
- 1 Fumigate | wipe | A full creature-board reset that fits the lifegain plan.
- 1 Vanquish the Horde | wipe | An efficient creature-board reset.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $599.99 to buy, $599.99 the whole deck.

**Summary:** This deck builds its mana with inexpensive Dragon-focused support, then fills the board with efficient Dragons and uses Atarka, World Render to turn combat into the main finishing route. Its removal suite keeps opposing creatures from stabilizing, while several board resets and protection spells help preserve momentum through contested games. It gives up some raw speed for a durable, creature-forward plan that needs time and mana to establish its strongest turns.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade. the mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.13 over 63 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 1 Faithless Looting | draw | Early card selection keeps the hand moving.
- 1 Demand Answers | draw | Flexible card flow supports the deck’s setup turns.
- 1 Thrill of Possibility | draw | Efficient card selection helps find threats or mana.
- 1 Skullclamp | draw | A low-cost source of sustained card flow.
- 1 Sylvan Library | draw | Provides repeatable card selection.
- 1 Elemental Bond | draw | Turns the creature-heavy plan into card flow.
- 1 Garruk's Uprising | draw | Rewards the deck for deploying large creatures.
- 1 Guardian Project | draw | Provides steady cards alongside the creature plan.
- 1 Beast Whisperer | draw | Converts creature deployment into additional cards.
- 1 Dragonborn Champion | draw | A Dragon-aligned draw engine for combat-heavy games.
- 1 Return of the Wildspeaker | draw | Flexible late-game card refill.
- 1 Heroic Intervention | interaction | Protects the developed board from opposing disruption.
- 1 Fog | interaction | Buys time against a decisive combat step.
- 1 Lightning Greaves | interaction | Protects an important creature efficiently.
- 1 Snakeskin Veil | interaction | A compact protection spell for a key threat.
- 1 Tamiyo's Safekeeping | interaction | Protects a valuable permanent when needed.
- 1 Veil of Summer | interaction | Efficient defensive interaction.
- 1 Swiftfoot Boots | interaction | Adds repeatable protection for major threats.
- 1 Tibalt's Trickery | interaction | Broad reactive interaction for a pivotal spell.
- 1 Carnelian Orb of Dragonkind | ramp | Early mana acceleration that supports Dragons.
- 1 Dragon's Hoard | ramp | Dragon-focused mana development for the midgame.
- 1 Jade Orb of Dragonkind | ramp | Low-cost acceleration for the deck’s creature plan.
- 1 Orb of Dragonkind | ramp | Reliable early mana acceleration.
- 1 Scaled Nurturer | ramp | Early creature-based mana development.
- 1 Reckless Barbarian | ramp | Cheap acceleration toward larger threats.
- 1 Embermouth Sentinel | ramp | Adds low-cost mana development.
- 1 Mox Jasper | ramp | Provides a fast mana option within the Dragon shell.
- 1 Sarkhan, Fireblood | ramp | Supports mana development while fitting the Dragon theme.
- 1 Ganax, Astral Hunter | ramp | Dragon-themed acceleration for building a large board.
- 1 Draconic Roar | removal | Efficient Dragon-themed spot removal.
- 1 Dragon's Fire | removal | Low-cost removal that fits the creature base.
- 1 Piercing Exhale | removal | A compact removal spell for problematic creatures.
- 1 Spit Flame | removal | Dragon-aligned removal for creature matchups.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Provides removal while remaining part of the Dragon plan.
- 1 Glorybringer | removal | A Dragon threat that also handles opposing creatures.
- 1 Terror of the Peaks | removal | Turns Dragon deployment into removal pressure.
- 1 Scourge of Valkas | removal | A Dragon payoff that supplies removal pressure.
- 1 Wrathful Red Dragon | removal | A Dragon-based removal payoff.
- 1 Crucible of Fire | synergy | Strengthens the deck’s central Dragon theme.
- 1 Dragonlord's Servant | synergy | Reduces the cost pressure of the Dragon curve.
- 1 Dragonspeaker Shaman | synergy | Helps deploy the deck’s Dragon threats efficiently.
- 1 Dragon Egg | synergy | A low-cost Dragon-themed setup piece.
- 1 Dragon Hatchling | synergy | A cheap Dragon body that supports the theme.
- 1 Dragonkin Berserker | synergy | A low-cost creature that rewards the Dragon plan.
- 1 Firespitter Whelp | synergy | A Dragon-themed early-game synergy piece.
- 1 Kargan Dragonrider | synergy | An inexpensive creature supporting Dragon deployment.
- 1 Last Light of Durin's Day | synergy | A thematic support piece for the Dragon strategy.
- 1 Dragon Tempest | removal | A Dragon payoff that adds removal pressure.
- 1 Breath Weapon | wipe | A low-cost board reset for creature-heavy tables.
- 1 Draconic Intervention | wipe | A scalable reset that fits the Dragon shell.
- 1 Ryusei, the Falling Star | wipe | A Dragon threat that can clear crowded boards.
- 1 Dragonloft Idol | threat | A low-cost Dragon threat for applying pressure.
- 1 Thunderbreak Regent | threat | An efficient Dragon threat for the combat plan.
- 1 Backdraft Hellkite | threat | A Dragon attacker that advances the deck’s pressure plan.
- 1 Blast-Furnace Hellkite | threat | A sizeable Dragon threat for closing games.
- 1 Hoarding Dragon | threat | A Dragon threat that supports the artifact package.
- 1 Manaform Hellkite | threat | A Dragon threat that keeps pressure on opponents.
- 1 Shivan Dragon | threat | A classic Dragon body for combat finishes.
- 1 Stirring Bard | threat | A Dragon-aligned threat for the board.
- 1 Tiamat's Fanatics | threat | A Dragon threat that adds to combat pressure.
- 1 Verix Bladewing | threat | A Dragon threat suited to the deck’s main plan.
- 1 Mirrorwing Dragon | threat | A resilient Dragon threat for combat-focused games.
- 1 Stormbreath Dragon | threat | An efficient Dragon attacker for finishing pressure.
- 12 Forest | land | Basic green mana for the deck’s Gruul base.
- 12 Mountain | land | Basic red mana for the deck’s Gruul base.
- 1 Cavern of Souls | land | A tribal land for the Dragon creature base.
- 1 Command Tower | land | Reliable color fixing for the commander’s colors.
- 1 Karplusan Forest | land | An untapped dual source for early development.
- 1 Stomping Ground | land | An efficient red-green dual land.
- 1 Taiga | land | An untapped red-green dual land.
- 1 Spire Garden | land | A multiplayer-focused red-green dual land.
- 1 Rockfall Vale | land | A red-green source for the mana base.
- 1 Rootbound Crag | land | A red-green source for the mana base.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible access to either deck color.
- 1 Secluded Courtyard | land | Tribal color fixing for Dragons.
- 1 Unclaimed Territory | land | Additional tribal color fixing.
- 1 Three Tree City | land | A tribal land supporting the Dragon-heavy board.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 93 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 1.

Cost: $0.00 to buy, $135.95 the whole deck.

**Summary:** An Orzhov lifegain deck built to establish a durable board, turn steady life gain into growing creature pressure, and finish through evasive Angels and equipped attackers. It balances resilient threats with broad removal, protective spells, and card advantage so it can stabilize early and steadily overwhelm the table.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. the commander does not place in cEDH events, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- NOTE: I could not place "Commander’s Sphere", so it is not in the deck. The shortlist holds Command Tower, Commander's Sphere.
- NOTE: I could not place "Wayfarer’s Bauble", so it is not in the deck. The shortlist holds Wayfarer's Bauble.
- NOTE: I could not place "Night’s Whisper", so it is not in the deck. The shortlist holds Night Nurse, Healer of Heroes, Night's Whisper.
- NOTE: I could not place "Champion’s Helm", so it is not in the deck. The shortlist holds Champion's Helm.
- NOTE: I could not place "White Mage’s Staff", so it is not in the deck. The shortlist holds White Auracite, White Lotus Hideout, White Lotus Tile.
- NOTE: I could not place "Victory’s Herald", so it is not in the deck. The shortlist holds Victory's Herald.
- [BLOCK] `deck_size`: deck has 94 cards, the format needs exactly 100 (commander included)
- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.95 over 56 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 22 Plains | land | Reliable white mana for the lifegain core and white-intensive spells.
- 15 Swamp | land | Reliable black mana for drain effects, card advantage, and removal.
- 1 Arcane Signet | ramp | Efficient fixing and acceleration.
- 1 Astral Cornucopia | ramp | Flexible mana production that scales into the late game.
- 1 Chromatic Lantern | ramp | Fixes both colors while accelerating mana.
- 1 Fellwar Stone | ramp | Efficient multiplayer mana acceleration.
- 1 Giada, Font of Hope | ramp | Accelerates the Angel portion of the deck.
- 1 Lotho, Corrupt Shirriff | ramp | Generates Treasures while advancing the deck’s game plan.
- 1 Sol Ring | ramp | Fast, reliable colorless acceleration.
- 1 Sword of the Animist | ramp | Turns attacks into lasting land-based ramp.
- 1 Buster Sword | draw | Equipment-based card advantage for attacking creatures.
- 1 Call of the Ring | draw | Steady card flow that pairs well with lifegain.
- 1 Exemplar of Light | draw | Lifegain-themed creature that converts the plan into cards.
- 1 Idol of Oblivion | draw | Low-cost repeatable card advantage.
- 1 Inspiring Overseer | draw | Provides an evasive body, life, and a card.
- 1 Lembas | draw | A life-oriented artifact that replaces itself.
- 1 Puresteel Paladin | draw | Draw engine for the deck’s Equipment package.
- 1 Skullclamp | draw | Efficient card advantage with smaller creatures and tokens.
- 1 Tome of Legends | draw | Repeatable card advantage alongside the commander.
- 1 Wall of Omens | draw | Early defense that immediately replaces itself.
- 1 Banishing Light | removal | Versatile answer to a troublesome permanent.
- 1 Bitter Triumph | removal | Flexible instant-speed creature or planeswalker removal.
- 1 Crib Swap | removal | Exiles a creature cleanly.
- 1 Dispatch | removal | Efficient creature interaction supported by artifacts.
- 1 Fatal Push | removal | Low-cost answer to early opposing threats.
- 1 Generous Gift | removal | Answers nearly any permanent type.
- 1 Infernal Grasp | removal | Reliable unconditional creature removal.
- 1 Bastion Protector | interaction | Keeps the commander protected while adding a useful body.
- 1 Clever Concealment | interaction | Guards the board from sweeping removal.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Lightning Greaves | interaction | Immediate protection and haste for important creatures.
- 1 Reprieve | interaction | Temporarily answers a pivotal spell while replacing itself.
- 1 Austere Command | wipe | Flexible reset button that can spare the preferred board.
- 1 Dusk // Dawn | wipe | Creature sweeper with later creature recovery.
- 1 Fumigate | wipe | Board reset that turns opposing creatures into life.
- 1 Aerith Gainsborough | synergy | Central lifegain payoff creature.
- 1 Angel of Vitality | synergy | Rewards repeated life gain with a larger evasive threat.
- 1 Compassionate Healer | synergy | Consistent life gain that supports the deck’s payoffs.
- 1 Crowd of True Believers | synergy | Adds a board presence aligned with the life-focused plan.
- 1 Eastfarthing Farmer | synergy | Supports the deck’s Food and lifegain elements.
- 1 Elixir | synergy | Provides a compact life-oriented utility piece.
- 1 Kor Firewalker | synergy | Efficient creature that reinforces the lifegain strategy.
- 1 Light of Promise | synergy | Converts life gain into a rapidly growing creature.
- 1 Night Nurse, Healer of Heroes | synergy | Life-focused legend that supports the creature plan.
- 1 Rosie Cotton of South Lane | synergy | Transforms life gain into permanent creature growth.
- 1 Second Breakfast | synergy | Food-focused lifegain support and value.
- 1 Well-Worn Spatula | synergy | Equipment support for the deck’s creature-based lifegain plan.
- 1 Angel of Invention | threat | Produces a meaningful airborne board and benefits from creature support.
- 1 Bill the Pony | threat | Efficient legendary threat that contributes to board pressure.
- 1 Dawnhand Eulogist | threat | Threat that benefits from the deck’s life-focused game plan.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A resilient black threat with useful incidental value.
- 1 Lyra Dawnbringer | threat | Powerful lifelinking finisher that stabilizes races.
- 1 Minwu, White Mage | threat | Life-oriented legendary threat for the midgame.
- 1 Rabaroo Troop | threat | Creature threat that adds pressure to the board.
- 1 Rooftop Percher | threat | Flexible creature threat for combat-focused turns.
- 1 Shattered Angel | threat | Evasive lifelink threat that punishes opposing land development.
- 1 Sneering Shadewriter | threat | Black evasive threat that pressures opponents.
- 1 Witch-king of Angmar | threat | High-impact black threat that makes combat difficult for opponents.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $12.75 to buy, $36.37 the whole deck.

**Summary:** This aristocrats deck develops a board of creatures, uses sacrifice-focused synergies to turn creature losses into incremental pressure, and backs that plan with broad answers and a few larger threats. It aims to win by maintaining a steady flow of creatures and repeatedly leveraging the deck’s death-focused pieces rather than relying on a single expensive finisher. It gives up premium acceleration and high-cost staples for a library-first, budget-conscious build.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. the commander does not place in cEDH events, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Sivriss, Nightmare Speaker: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.18 over 62 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 1 Ahriman | draw | Provides a library-owned draw option.
- 1 Beetle-Headed Merchants | draw | Adds inexpensive, library-owned card flow.
- 1 Buzzard-Wasp Colony | draw | Supplies another owned draw piece.
- 1 Circle of Power | draw | Adds card flow from the existing library.
- 1 Cirith Ungol Patrol | draw | Provides a low-cost owned draw slot.
- 1 Exemplar of Light | draw | Adds creature-based card flow.
- 1 Grave Venerations | draw | Supports continued access to cards.
- 1 Inspiring Overseer | draw | Provides an owned creature-based draw option.
- 1 June, Bounty Hunter | draw | Adds another inexpensive draw card.
- 1 Kingpin's Enforcers | draw | Fills out the deck's creature-based card flow.
- 1 Nasty End | draw | Provides efficient card flow from the library.
- 1 Bastion Protector | interaction | Helps keep an important creature in play.
- 1 Boromir, Warden of the Tower | interaction | Provides an owned defensive interaction piece.
- 1 Bronze Guardian | interaction | Adds resilient artifact-based protection.
- 1 Gift of Immortality | interaction | Helps preserve a key creature through removal.
- 1 Sheltered by Ghosts | interaction | Offers a flexible protective interaction slot.
- 1 Swiftfoot Boots | interaction | Protects an important creature immediately.
- 1 Arcane Signet | ramp | Provides dependable color fixing from the library.
- 1 Astral Cornucopia | ramp | Adds an owned mana source for longer games.
- 1 Bender's Waterskin | ramp | Provides another library-owned mana piece.
- 1 Blitzball | ramp | Adds efficient artifact-based mana support.
- 1 Chromatic Lantern | ramp | Provides mana support and color fixing.
- 1 Commander's Sphere | ramp | Adds flexible mana support.
- 1 Deadly Dispute | ramp | Supports the mana plan while fitting the sacrifice theme.
- 1 Elementalist's Palette | ramp | Adds a library-owned mana option.
- 1 Explorer's Scope | ramp | Provides low-cost mana development.
- 1 Fellwar Stone | ramp | Adds inexpensive color fixing.
- 1 Angel of Serenity | removal | Provides a high-impact creature answer.
- 1 Bitter Triumph | removal | Adds efficient instant-speed removal.
- 1 Claim the Precious | removal | Provides a straightforward answer to a threat.
- 1 Crib Swap | removal | Adds flexible creature removal.
- 1 Deadly Precision | removal | Provides inexpensive removal from the library.
- 1 Destroy Evil | removal | Adds a versatile answer to troublesome permanents.
- 1 Fiend Hunter | removal | Provides creature-based removal that fits the deck's board plan.
- 1 Archfiend of Ifnir | wipe | Supplies a creature-based board-clearing option.
- 1 Austere Command | wipe | Provides a flexible reset button.
- 1 Dusk // Dawn | wipe | Adds a board reset with useful follow-up value.
- 1 Aron, Benalia's Ruin | synergy | Supports the deck's aristocrats synergy.
- 1 Ayli, Eternal Pilgrim | synergy | Adds a low-cost aristocrats-focused synergy piece.
- 1 Bartolomé del Presidio | synergy | Supports the creature-sacrifice game plan.
- 1 Bastion of Remembrance | synergy | Adds a dedicated aristocrats synergy permanent.
- 1 Blood Artist | synergy | Provides a central death-focused payoff.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Adds a compact aristocrats payoff.
- 1 Falkenrath Noble | synergy | Provides another death-focused payoff.
- 1 Gollum the Abandoned | synergy | Supports the deck's existing sacrifice synergies.
- 1 Gollum, Patient Plotter | synergy | Adds another library-owned synergy creature.
- 1 Vindictive Vampire | synergy | Provides an additional aristocrats payoff.
- 1 Woe Strider | synergy | Supports the creature-sacrifice plan.
- 1 Yahenni, Undying Partisan | synergy | Adds a low-cost sacrifice-focused creature.
- 1 Zulaport Cutthroat | synergy | Provides another efficient death-focused payoff.
- 1 Bill the Pony | threat | Adds an owned creature threat.
- 1 Hei Bai, Spirit of Balance | threat | Provides a library-owned creature threat.
- 1 Namazu Trader | threat | Adds a low-cost creature threat.
- 1 Vengeful Villagers | threat | Supplies another owned threat for the board.
- 1 Baron Bertram Graywater | threat | Adds an inexpensive creature threat.
- 1 Lord Skitter's Butcher | threat | Provides a budget creature threat.
- 1 Old Flitterfang | threat | Adds another low-cost threat.
- 1 Sivriss, Nightmare Speaker | threat | Provides an inexpensive threat that supports the creature plan.
- 1 Skullport Merchant | threat | Adds a budget creature threat.
- 1 Smothering Abomination | threat | Provides a larger threat for the sacrifice-focused plan.
- 1 Vampire Gourmand | threat | Adds an inexpensive creature threat.
- 1 Vengeful Bloodwitch | threat | Supplies another budget creature threat.
- 1 Command Tower | land | Provides reliable access to both deck colors.
- 1 Exotic Orchard | land | Adds flexible color fixing.
- 1 Grand Coliseum | land | Provides flexible colored mana.
- 1 Opal Palace | land | Adds a colored land from the library.
- 1 Path of Ancestry | land | Provides a budget dual-color land.
- 1 Scavenger Grounds | land | Adds a utility land slot.
- 1 Secluded Courtyard | land | Provides creature-oriented color fixing.
- 1 Spire of Industry | land | Adds flexible colored mana alongside artifacts.
- 1 Thriving Moor | land | Provides additional black mana fixing.
- 1 Unclaimed Territory | land | Adds another creature-oriented fixing land.
- 14 Plains | land | Provides a stable base of white mana.
- 13 Swamp | land | Provides a stable base of black mana.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band, two_card_combo. Block findings: 0.

Cost: $79.08 to buy, $231.84 the whole deck.

**Summary:** A mono-red Goblin swarm deck that develops a broad battlefield, turns targeted cantrips into overwhelming value, and converts token production, combat pressure, and sacrifice effects into explosive finishing turns.

The quality model grades this deck good: few of the cards are ones the top lists of the format play, and that lowers the grade. the mana base serves the colors evenly, and that raises the grade. nearly every card appears in a top list, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.61 over 61 nonland cards
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | draw | Targeted cantrip that scales strongly with the commander and a wide board.
- 1 Battle Hymn | ramp | Converts a Goblin board into a burst of red mana for explosive turns.
- 1 Blasphemous Act | wipe | Efficient reset button when opposing creature boards get ahead.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into reach and supports sacrifice lines.
- 1 Chaos Warp | removal | Versatile answer to troublesome permanents.
- 1 Conspicuous Snoop | synergy | Tribal card advantage and a Goblin body for the board.
- 1 Crimson Wisps | draw | Cheap targeted cantrip that can give the team haste through the commander.
- 1 Daring Discovery | other | Flexible spell that supports the deck's proactive game plan.
- 1 Dragon Fodder | synergy | Efficiently supplies multiple Goblins for go-wide and sacrifice synergies.
- 1 Empty the Warrens | wincon | Creates a large Goblin board after a spell-heavy turn.
- 1 Expedite | draw | Low-cost targeted cantrip that enables immediate attacks.
- 1 Faithless Looting | draw | Efficient hand filtering and graveyard setup.
- 1 Fists of Flame | draw | A targeted card-draw spell that can make a wide attack lethal.
- 1 Frontline Heroism | synergy | Supports aggressive combat and rewards building a creature board.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal with cycling utility.
- 1 General Kreat, the Boltbringer | synergy | Goblin-focused payoff that advances the token plan.
- 1 Glimpse the Impossible | ramp | Provides temporary access to extra resources for explosive turns.
- 1 Goblin Bombardment | removal | Reliable sacrifice outlet and repeatable creature-based damage.
- 1 Goblin Bushwhacker | synergy | Turns a token deployment into a fast, high-damage attack.
- 1 Goblin Chieftain | synergy | Anthem and haste effect for the Goblin army.
- 1 Goblin Dark-Dwellers | threat | Recasts a useful instant or sorcery while adding a Goblin body.
- 1 Goblin Lackey | synergy | Accelerates Goblin deployment after connecting in combat.
- 1 Goblin Matron | synergy | Finds the right Goblin payoff or utility creature.
- 1 Goblin Negotiation | removal | Tribal removal that keeps opponents' boards manageable.
- 1 Goblin Trashmaster | removal | Goblin lord that also answers artifacts.
- 1 Goblin Warchief | synergy | Reduces Goblin costs and grants haste to pressure quickly.
- 1 Grapeshot | removal | Storm payoff that can clear small creatures or finish weakened opponents.
- 1 Idol of Oblivion | draw | Steady card advantage alongside regular token production.
- 1 Impact Tremors | wincon | Converts every Goblin and token burst into direct damage.
- 1 Impulsive Pilferer | ramp | Early Goblin that leaves behind Treasure when it dies.
- 1 Krenko's Command | synergy | Adds multiple Goblins for token and combat synergies.
- 1 Krenko, Mob Boss | threat | Major repeatable token engine and primary board-building threat.
- 1 Mana Geyser | ramp | Produces a large ritual burst for a decisive turn.
- 1 Mogg War Marshal | synergy | Multiple Goblin bodies from one card support sacrifice and go-wide plans.
- 1 Past in Flames | other | Lets the deck reuse a graveyard full of cheap spells.
- 1 Quest for the Goblin Lord | synergy | Rewards repeated Goblin development with a powerful team boost.
- 1 Roaming Throne | synergy | Doubles key Goblin triggered abilities after choosing the creature type.
- 1 Ruby Medallion | ramp | Reduces the cost of the deck's dense red spell suite.
- 1 Rundvelt Hordemaster | synergy | Goblin anthem and death-trigger card advantage.
- 1 Sazacap's Brew | draw | Targeted card draw that can be copied across the creature board.
- 1 Searslicer Goblin | synergy | Contributes another Goblin body and tribal payoff.
- 1 Seething Song | ramp | Efficient ritual mana for high-tempo turns.
- 1 Siege-Gang Commander | removal | Provides a Goblin army and repeatable sacrifice-based damage.
- 1 Siege-Gang Lieutenant | removal | Adds token production and Goblin-based board presence.
- 1 Skirk Prospector | ramp | Sacrifices Goblins for explosive mana and enables death synergies.
- 1 Skullclamp | draw | Exceptional card draw with expendable Goblin tokens.
- 1 Sol Ring | ramp | Fast, reliable mana acceleration.
- 1 Storm-Kiln Artist | ramp | Builds Treasure from the deck's frequent instant and sorcery spells.
- 1 Swiftfoot Boots | interaction | Protects the commander or a key Goblin while enabling haste.
- 1 Throne of Eldraine | ramp | Mana rock that supports the deck's tribal creature plan.
- 1 Vandalblast | wipe | Efficient artifact interaction with a powerful overloaded mode.
- 1 Warren Torchmaster | synergy | Goblin-focused combat payoff that supports attacking with a wide board.
- 1 Wild Ride | other | Supports proactive combat and creature pressure.
- 1 Witch's Mark | draw | Cheap hand filtering that can grant haste to a key creature.
- 1 Ancestral Anger | draw | A cheap targeted cantrip that rewards a wide board and fuels spell chains.
- 1 Brightstone Ritual | ramp | Converts a developed Goblin board into a large burst of mana.
- 1 Broadside Bombardiers | removal | Turns expendable tokens into substantial removal or direct damage.
- 1 Goblin Fireleaper | removal | Rewards targeting a Goblin and provides scalable direct damage.
- 1 Moria Marauder | synergy | Adds a Goblin attacker that rewards the deck's tribal development.
- 1 Pashalik Mons | removal | Creates Goblins and turns Goblin deaths into damage.
- 1 Raging Goblinoids | threat | Adds a Goblin threat that helps maintain combat pressure.
- 1 Arena of Glory | land | Red-producing utility land that can give a key creature immediate impact.
- 1 Castle Embereth | land | Red source with a late-game team-pump option.
- 1 Den of the Bugbear | land | Red source that becomes an attacking threat when needed.
- 1 Dwarven Mine | land | Red source that can create an extra body after sufficient Mountain development.
- 1 Forgotten Cave | land | Red source with cycling utility in low-resource situations.
- 1 Goblin Burrows | land | Utility land that can enhance a Goblin in combat.
- 1 Hidden Volcano | land | Red-producing utility land for the mana base.
- 1 Kher Keep | land | Utility land that supplies small tokens for sacrifice synergies.
- 1 Reliquary Tower | land | Utility land that supports retaining a full hand after large draw turns.
- 1 War Room | land | Repeatable card-draw utility from the mana base.
- 28 Mountain | land | Reliable red mana for a mono-red Goblin and spell-heavy strategy.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for game_changer_limit, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $216.33 to buy, $216.33 the whole deck.

**Summary:** A five-color Turtle Power creature deck that develops its mana early, builds a resilient board of themed legends and allies, and supports its combat-focused plan with steady card advantage. It combines targeted answers, protective spells, and broad reset buttons to keep the table manageable before closing with a reinforced Turtle assault.

The quality model grades this deck good: the cards pair the way the top lists pair them, and that raises the grade. few of the cards are ones the top lists of the format play, and that lowers the grade. the mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.29 over 63 nonland cards
- [INFO] `precon_cards_restored`: the deck was 3 cards short of the Turtle Power precon share, so the builder put 3 cards back
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 1 Ash Barrens | land | Fixes colors while preserving a precon land.
- 1 Big Apple, 3 a.m. | land | Retained themed fixing land.
- 1 Cinder Glade | land | Retained dual land for the multicolor base.
- 1 Command Tower | land | Reliable five-color fixing.
- 1 Escape Tunnel | land | Finds a needed basic color.
- 1 Evolving Wilds | land | Flexible basic-land fixing.
- 1 Fabled Passage | land | Flexible basic-land fixing.
- 1 Grand Coliseum | land | Retained broad color fixing.
- 1 Hidden Hideout | land | Retained themed mana source.
- 1 Path of Ancestry | land | Creature-focused fixing.
- 1 Rain-Slicked Copse | land | Retained dual land.
- 1 Rootbound Crag | land | Retained dual land.
- 1 Sodden Verdure | land | Retained dual land.
- 1 Thriving Grove | land | Flexible color fixing.
- 1 Thriving Isle | land | Flexible color fixing.
- 1 Thriving Moor | land | Flexible color fixing.
- 1 Turtle Lair | land | Retained on-theme land.
- 1 Undergrowth Stadium | land | Retained dual land.
- 1 Vernal Fen | land | Retained dual land.
- 1 Vibrant Cityscape | land | Retained broad fixing.
- 1 Breeding Pool | land | Untapped green-blue fixing for early development.
- 1 Hallowed Fountain | land | Untapped white-blue fixing.
- 1 Watery Grave | land | Untapped blue-black fixing.
- 1 Overgrown Tomb | land | Untapped black-green fixing.
- 1 Temple Garden | land | Untapped green-white fixing.
- 1 Blood Crypt | land | Untapped black-red fixing.
- 1 Steam Vents | land | Untapped blue-red fixing.
- 1 Sacred Foundry | land | Untapped red-white fixing.
- 3 Forest | land | Basic source for green early plays.
- 2 Island | land | Basic source for blue interaction.
- 1 Plains | land | Basic source for white spells.
- 1 Swamp | land | Basic source for black spells.
- 1 Mountain | land | Basic source for red spells.
- 1 Acidic Slime | removal | Reusable creature-based answer to troublesome permanents.
- 1 April O'Neil, Live on the Scene | draw | Themed card advantage engine.
- 1 Arcade Cabinet | synergy | Retained Turtle Power support piece.
- 1 Arcane Signet | ramp | Efficient multicolor acceleration.
- 1 Assassin's Trophy | removal | Answers any problematic permanent.
- 1 Baxter, Fly in the Ointment | other | Retained precon character and theme support.
- 1 Bebop, Skull & Crossbones | threat | Retained precon character threat.
- 1 Big Mother Mouser | other | Retained themed creature.
- 1 Biogenic Ooze | threat | Provides a growing board presence.
- 1 Chromatic Lantern | ramp | Acceleration and comprehensive color fixing.
- 1 Coin of Mastery | other | Retained themed value artifact.
- 1 Continue? | draw | Retained themed source of value.
- 1 Corpsejack Menace | synergy | Amplifies counter-based Turtle Power play patterns.
- 1 Dimension X Pizzasaur | other | Retained themed creature.
- 1 Double Jump // Flying Kick | other | Retained flexible themed spell.
- 1 Electric Seaweed | other | Retained precon utility creature.
- 1 Endless Foot Assault | synergy | Retained thematic board-building piece.
- 1 Everything Pizza | draw | Retained themed value artifact.
- 1 Exploding Barrel | other | Retained themed utility artifact.
- 1 Fast Forward | ramp | Retained themed mana-development spell.
- 1 Game Over | wincon | Retained thematic finisher.
- 1 Harmonize | draw | Straightforward card advantage.
- 1 Here Comes a New Hero! | draw | Themed card advantage and creature support.
- 1 High Score | draw | Retained themed value engine.
- 1 Krang, the All-Powerful | threat | Retained precon centerpiece.
- 1 Leatherhead, Iron Gator | threat | Retained themed creature.
- 1 Lessons from Life | draw | Themed card-advantage spell.
- 1 Level Up | synergy | Retained creature-enhancement support.
- 1 Lita, Little Orphan Amphibian | synergy | Supports the Turtle-focused creature plan.
- 1 Michelangelo, the Heart | synergy | Core Turtle Power synergy piece.
- 1 Mole Module | other | Retained themed utility artifact.
- 1 Mona Lisa, Science Geek | other | Retained precon character support.
- 1 Ninja Pizza | synergy | Retained thematic value piece.
- 1 Rat King, Pale Piper | threat | Retained precon character.
- 1 Ray Fillet, Wave Warrior | threat | Retained precon character.
- 1 Roadkill Rodney | other | Retained themed utility creature.
- 1 Rocksteady, Mutant Marauder | threat | Retained precon character.
- 1 Saved by the Shell | interaction | Protects the team from opposing disruption.
- 1 Shellshock | removal | Retained flexible themed disruption.
- 1 Shredder, Shadow Master | threat | Retained precon antagonist and threat.
- 1 Sol Ring | ramp | Fast early acceleration.
- 1 Special Move | other | Retained themed utility spell.
- 1 Spikeshell Harrier | removal | Creature-based removal that fits the theme.
- 1 Splinter & Leo, Father & Son | synergy | Supports the Turtle legend theme.
- 1 Splinter, the Mentor | synergy | Retained precon character support.
- 1 Steelbane Hydra | removal | Repeatable artifact and enchantment removal.
- 1 Super Combo | wincon | Retained thematic payoff.
- 1 Swift Demise | removal | Efficient themed creature answer.
- 1 Tempestra, Dame of Games | other | Retained precon character.
- 1 Together Forever | synergy | Helps preserve important creatures.
- 1 Tokka & Rahzar, Unsupervised | ramp | On-theme creature acceleration.
- 1 Turtles in Time | wipe | Themed reset for crowded boards.
- 1 Vanquish the Horde | wipe | Reliable creature-board reset.
- 1 Voracious Hydra | threat | Scalable threat with creature interaction.
- 1 Wave Goodbye | wipe | Retained precon board-control spell.
- 1 Smothering Tithe | ramp | Powerful long-game mana production.
- 1 Cyclonic Rift | wipe | Flexible emergency answer and asymmetrical reset.
- 1 Teferi's Protection | interaction | Protects the board and life total through critical turns.
- 1 Farseek | ramp | Efficiently finds a needed dual land.
- 1 Nature's Lore | ramp | Efficiently develops mana and color access.
- 1 Raphael, the Muscle | draw | the deck upgrades a precon, and this card is one the precon holds
- 1 Foot Chopper | interaction | the deck upgrades a precon, and this card is one the precon holds
- 1 Donatello, the Brains | interaction | the deck upgrades a precon, and this card is one the precon holds

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $368.87 to buy, $368.87 the whole deck.

**Summary:** A Sultai Elf-focused Commander deck that develops a deep creature board, converts that board into mana and card advantage, and uses flexible answers to keep opponents from stabilizing. Its main path to victory is resilient combat pressure backed by legendary Elf support, protective interaction, and tribal-friendly sweepers.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. more opening hands hold two to four lands than the norm, and that raises the grade. the deck makes less mana on turn four than the norm, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 3.37 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.84, and bracket 3 wants 0.85 or more (sources of requirement: U 16.75 of 20, B 16.75 of 19, G 18.25 of 20)

<details><summary>The deck list</summary>

- 10 Island | land | Provides reliable blue mana for card advantage and reactive spells.
- 9 Swamp | land | Provides black mana for removal and recursion-oriented effects.
- 9 Forest | land | Provides green mana for Elf development and creature-based acceleration.
- 1 Elven Passage | land | Thematic colored land supporting the Elf-focused mana base.
- 1 Elvenking's Halls | land | Thematic fixing land for the deck's three-color requirements.
- 1 Minas Morgul, Dark Fortress | land | Black-producing thematic utility land.
- 1 Mirkwood | land | Thematic colored land supporting creature development.
- 1 Rivendell | land | Blue-producing thematic land for the reactive game plan.
- 1 The Black Gate | land | Thematic black mana source.
- 1 The Shire | land | Thematic green mana source.
- 1 Dragon-Cursed Halls | land | Additional thematic fixing land for the mana base.
- 1 Arcane Signet | ramp | Efficient fixing for all of the deck's colors.
- 1 Mox Amber | ramp | Low-cost acceleration enabled by the legendary commander.
- 1 Elvish Mystic | ramp | Early green acceleration that contributes to the Elf board.
- 1 Delighted Halfling | ramp | Colored acceleration that helps deploy important legendary spells.
- 1 Elvish Archdruid | ramp | Elf-based mana production that scales with the board.
- 1 Wood Elves | ramp | Develops the board while finding a Forest.
- 1 Woodland Weavemaster | ramp | Creature-based acceleration that supports the deck's tribal plan.
- 1 Elven Chorus | ramp | Turns the creature-heavy board into additional mana.
- 1 Thranduil's Company | ramp | Thematic Elf acceleration for building a broad board.
- 1 Through the Forest Gate | ramp | Land-based ramp and color access.
- 1 Elvish Visionary | draw | A cheap Elf that replaces itself and supports creature synergies.
- 1 Lórien Revealed | draw | Flexible card selection with early landcycling utility.
- 1 Hithlain Knots | draw | Instant-speed card advantage.
- 1 Night's Whisper | draw | Efficient, low-cost card draw.
- 1 Fateful Discovery | draw | Persistent card-advantage engine.
- 1 Gandalf, Shadow's Foe | draw | A substantial legend that generates card advantage.
- 1 Gandalf, Wandering Wizard | draw | Legendary card advantage suited to a longer game.
- 1 Last March of the Ents | draw | Refills the hand from a developed creature board.
- 1 Plunder the Trollshaws | draw | Instant-speed draw that rewards combat development.
- 1 Reverent Howl | draw | Flexible card draw for maintaining resources.
- 1 Uncover the Moon-Letters | draw | Ongoing card selection and hand replenishment.
- 1 Confusticate and Bebother | interaction | Flexible stack interaction.
- 1 Elrond, Moon-Reader | interaction | A thematic legendary defensive tool.
- 1 Mithril Coat | interaction | Protects a key creature from removal.
- 1 My Precious // Allure of Power | interaction | Versatile protection and disruption on a single card.
- 1 Old Fat Spider Can't See Me | interaction | Defensive interaction that disrupts opposing plans.
- 1 Sound the Trumpets | interaction | Reactive spell for protecting the board or disrupting combat.
- 1 Stern Scolding | interaction | Efficient early interaction against problematic creatures.
- 1 Thranduil's Decree | interaction | Thematic instant-speed disruption.
- 1 Bilbo's Deadly Slice | removal | Efficient targeted creature removal.
- 1 Bitter Downfall | removal | Broad answer to troublesome permanents.
- 1 Colossal Whale | removal | Creature threat that repeatedly removes opposing creatures.
- 1 Enchanted River's Grasp | removal | Thematic enchantment-based answer to a permanent.
- 1 Merciless Executioner | removal | Edict removal that can exploit disposable creature bodies.
- 1 Orcish Bowmasters | removal | Efficient creature control that punishes opposing card draw.
- 1 Quarrel | removal | Low-cost spot removal.
- 1 Stir Up Trouble | removal | Versatile removal spell.
- 1 Uneasy Partings | removal | Instant-speed answer to a problematic permanent.
- 1 Gnashing of Teeth | wipe | Board reset for creature-heavy opposing positions.
- 1 Languish | wipe | Efficient mass creature removal.
- 1 Raise the Palisade | wipe | Tribal-friendly sweep that can preserve the Elf board.
- 1 Arwen, Weaver of Hope | synergy | Builds resilient creatures and rewards continued board development.
- 1 Boughside Wanderers | synergy | Elf body that advances the tribe's battlefield presence.
- 1 Cantankerous Keepers | synergy | Elf creature supporting the deck's creature density.
- 1 Celeborn the Wise | synergy | Legendary Elf payoff for the deck's thematic plan.
- 1 Elven Raft-Steerer | synergy | Elf utility creature that supports the board-focused strategy.
- 1 Elvenking's Harper | synergy | Thematic Elf that contributes to tribal development.
- 1 Galadhrim Guide | synergy | Elf support creature for advancing attacks and board presence.
- 1 Galion, Elvenking's Butler | synergy | Thematic legendary Elf support piece.
- 1 Guardian of the Halls | synergy | Elf defender that helps protect the developing board.
- 1 Lothlórien Lookout | synergy | Elf utility creature that reinforces the tribal core.
- 1 Attercop | threat | Large thematic creature that pressures opponents in combat.
- 1 Haunt of the Dead Marshes | threat | A dangerous evasive threat with strong thematic flavor.
- 1 Mirkwood Elk | threat | Efficient green creature for board presence and combat pressure.
- 1 Mirkwood Meditator | threat | Elf threat that benefits from the deck's creature infrastructure.
- 1 Mirkwood Nurturer | threat | Creature threat that develops alongside the Elf plan.
- 1 Mirkwood Pathmaker | threat | Elf body that helps turn a wide board into pressure.
- 1 Nasty Little Rabbit | threat | Low-cost creature that contributes to a steady combat board.
- 1 Nimrodel Watcher | threat | Elf threat that supports sustained battlefield development.
- 1 Silvan Reveler | threat | Elf creature that remains useful while applying pressure.
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | Legendary Elf threat with a useful rally effect.
- 1 Willow-Wind | threat | Substantial green creature that closes games through combat.
- 1 Great Fierce Bee | threat | Thematic creature that adds another meaningful combat threat.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $506.23 to buy, $506.23 the whole deck.

**Summary:** A Smaug-led dragon deck that develops its mana through Treasure and artifacts, then turns equipped creatures and imposing threats into sustained combat pressure. It combines flavorful legendary gear, resilient value engines, targeted answers, and dramatic battlefield resets to dominate the middle and late game.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. the commander does not place in cEDH events, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.29 over 63 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 1 Dragon-Cursed Halls | land | Red-producing themed land.
- 1 Hobbit Hole | land | Themed nonbasic mana source.
- 31 Mountain | land | Reliable red mana base.
- 1 Rogue's Passage | land | Makes a large attacker difficult to block.
- 1 The Lonely Mountain | land | Themed red mana source.
- 1 Treasure Vault | land | Late-game Treasure outlet and mana source.
- 1 Arcane Signet | ramp | Efficient color fixing.
- 1 Bag End Banquet | ramp | Treasure-oriented mana acceleration.
- 1 Burn, Burn, Tree and Fern | ramp | Themed mana development.
- 1 Cavern-Hoard Dragon | ramp | Dragon threat that converts combat into mana.
- 1 Dragon's Desire | ramp | Mana acceleration for expensive threats.
- 1 Fíli and Kíli, Joyous | ramp | Creature-based mana development.
- 1 Mox Amber | ramp | Low-cost legendary mana acceleration.
- 1 The Misty Mountains Cold | ramp | Saga-based mana development.
- 1 The Reaver Cleaver | ramp | Turns combat damage into Treasure.
- 1 Wayfarer's Bauble | ramp | Early land-based ramp.
- 1 Balin, Loremaster | draw | Repeatable card advantage on a themed body.
- 1 Bothersome Noisemaker | draw | Supports the deck's ongoing card flow.
- 1 Gundabad Opportunist | draw | Provides incremental value while developing the board.
- 1 Key to the Side-Door | draw | Artifact card-advantage piece.
- 1 Old Thrush | draw | Low-cost value creature.
- 1 Óin the Brave | draw | Themed creature-based card advantage.
- 1 Palantír of Orthanc | draw | Powerful persistent card selection.
- 1 Ragged Short Spear | draw | Equipment that supplies card advantage.
- 1 Smaug's Fury | draw | Flexible spell that keeps resources flowing.
- 1 Thrór's Map | draw | Reliable artifact card advantage.
- 1 Tidings of War | draw | Refills resources while advancing the game plan.
- 1 Battle-Scarred Goblin | removal | Creature-based spot removal.
- 1 Fire of Orthanc | removal | Efficient targeted answer.
- 1 Gandalf, Spark Starter | removal | Themed removal on a creature.
- 1 Giant's Boulder | removal | Artifact-based targeted removal.
- 1 Goblin Cratermaker | removal | Flexible creature or artifact answer.
- 1 Improvised Club | removal | Low-cost instant-speed removal.
- 1 Inferno Titan | removal | Threat that removes small creatures.
- 1 Smaug, the Great Calamity // Spew Flame | removal | Dragon threat with removal attached.
- 1 Smite the Deathless | removal | Clean targeted removal.
- 1 Call Forth the Tempest | wipe | Broad reset for crowded boards.
- 1 Desolation of Smaug | wipe | Flavorful board-clearing effect.
- 1 Glóin the Mighty // Easy Pickings | wipe | Creature that offers a sweeping Adventure.
- 1 Bilbo's Ring | interaction | Protective and disruptive equipment.
- 1 Dwarven Mattock | interaction | Flexible equipment-based interaction.
- 1 Getaway Barrel | interaction | Utility artifact for protecting key permanents.
- 1 Glamdring | interaction | Legendary equipment utility.
- 1 Long-Lost Lances | interaction | Combat utility for important attackers.
- 1 Mithril Coat | interaction | Protects the commander or a major threat.
- 1 Sting, Bilbo's Sword | interaction | Low-cost equipment utility.
- 1 The One Ring | interaction | Protective legendary artifact with substantial utility.
- 1 Andúril, Flame of the West | synergy | Legendary equipment supporting the combat plan.
- 1 Andúril, Narsil Reforged | synergy | Legendary equipment reinforcing the theme.
- 1 Dwarven Warriors | synergy | Themed supporting creature.
- 1 Goblin-town Flunkies | synergy | Low-cost creature that supports the aggressive board plan.
- 1 Guttersnipe | synergy | Rewards the deck's suite of impactful spells.
- 1 Last Light of Durin's Day | synergy | Central enchantment for the deck's value plan.
- 1 Misty Mountains Raider | synergy | Themed attacker supporting combat pressure.
- 1 Snowslope Hunter | synergy | Themed creature supporting the board plan.
- 1 Thorin, Company's Leader | synergy | Legendary creature that supports the broader theme.
- 1 Well-Worn Spatula | synergy | Equipment that complements the artifact package.
- 1 Dáin Ironfoot | threat | Legendary creature that applies meaningful combat pressure.
- 1 Desert Were-Worm | threat | Large Dragon creature and primary finisher.
- 1 Dori, Bearer of Friends | threat | Legendary body that advances the board.
- 1 Dwarven Mauler | threat | Aggressive creature threat.
- 1 Iron Hills Stalwart | threat | Efficient themed combat body.
- 1 Oliphaunt | threat | Large creature capable of ending games.
- 1 Olog-hai Crusher | threat | Heavy combat finisher.
- 1 Orcish Siegemaster | threat | Board-building creature threat.
- 1 Stone-Giant of High Pass | threat | Large attacker with battlefield impact.
- 1 The Black Arrow | threat | High-impact legendary artifact threat.
- 1 Thorin, Mountain-king | threat | Legendary threat with strong battlefield presence.
- 1 Troop of Ponies | threat | Creature-based pressure that contributes to the board.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $1022.95 to buy, $1022.95 the whole deck.

**Summary:** A red-black Middle-earth deck that builds a treasure-backed hoard, develops a board of Goblins, legends, and monstrous creatures, then leverages equipment and combat pressure to dominate the table. It combines efficient answers with durable card advantage and sweeping effects to keep opponents from stabilizing before its great threats take over.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. more opening hands hold two to four lands than the norm, and that raises the grade. the deck makes less mana on turn four than the norm, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 3.03 over 63 nonland cards
- [WARN] `profile_off_band`: the turns the commander comes down after its mana value is 0.87, and bracket 3 wants 0.8 at most (the commander costs 7 and comes down on turn 7.87 on average)
- [WARN] `outside_requested_set`: the sets you named do not hold 17 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Evolving Wilds, Exotic Orchard, Fabled Passage, Marsh Flats, Misty Rainforest, Polluted Delta, Scalding Tarn, Terramorphic Expanse, Urborg, Tomb of Yawgmoth, Verdant Catacombs, Wooded Foothills

<details><summary>The deck list</summary>

- 1 Ancient Tomb | land | Accelerates high-impact plays while remaining within the colorless-land allowance.
- 1 Arid Mesa | land | Fetches a Mountain or the deck’s red-black dual land.
- 1 Blood Crypt | land | Reliable untapped access to both commander colors.
- 1 Bloodstained Mire | land | Fetches either basic color or the red-black dual land.
- 1 Command Tower | land | Provides either required color without conditions.
- 1 Dragonskull Summit | land | Dual-color source suited to the deck’s basic-land base.
- 1 Evolving Wilds | land | Fixes colors while increasing the basic-land count in play.
- 1 Exotic Orchard | land | Flexible color fixing in multiplayer games.
- 1 Fabled Passage | land | Fetches the needed basic color and improves later mana.
- 1 Marsh Flats | land | Fetches Swamp or the red-black dual land.
- 1 Misty Rainforest | land | Fetches the Mountain-containing dual land when needed.
- 1 Mount Doom | land | Thematic mana source for the dragon’s domain.
- 1 Polluted Delta | land | Fetches Swamp or the red-black dual land.
- 1 Scalding Tarn | land | Fetches Mountain or the red-black dual land.
- 1 Terramorphic Expanse | land | Additional basic-land fixing.
- 1 Urborg, Tomb of Yawgmoth | land | Improves black access and lets utility lands produce black mana.
- 1 Verdant Catacombs | land | Fetches Swamp or the red-black dual land.
- 1 Wooded Foothills | land | Fetches Mountain or the red-black dual land.
- 9 Mountain | land | Stable red mana for dragons, burn, and Goblins.
- 9 Swamp | land | Stable black mana for removal and darker threats.
- 1 Arcane Signet | ramp | Efficient, dependable color fixing.
- 1 Bag End Banquet | ramp | Treasure-oriented mana development that fits the hoard theme.
- 1 Bolg's Company | ramp | Creature-based acceleration that advances the board.
- 1 Burn, Burn, Tree and Fern | ramp | Thematic Saga-based mana acceleration.
- 1 Dragon's Desire | ramp | Helps produce the mana needed for expensive dragon plays.
- 1 Mox Amber | ramp | Fast legendary-focused acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment that contributes to mana development.
- 1 Smaug the Magnificent | ramp | A thematic Dragon that helps build toward larger plays.
- 1 The Reaver Cleaver | ramp | Turns combat pressure into a substantial Treasure hoard.
- 1 Wayfarer's Bauble | ramp | Early color fixing through a basic land.
- 1 Balin, Loremaster | draw | Legendary creature providing repeatable thematic card advantage.
- 1 Gollum, Riddle Master | draw | Provides card selection and fits the Middle-earth cast.
- 1 Key to the Side-Door | draw | Artifact-based card advantage with thematic utility.
- 1 Night's Whisper | draw | Low-cost, efficient card flow.
- 1 Palantír of Orthanc | draw | Persistent card advantage that pressures opponents.
- 1 Ragged Short Spear | draw | Equipment that supports combat while replacing resources.
- 1 Reverent Howl | draw | Instant-speed card advantage.
- 1 The Master of Lake-town | draw | A thematic source of continued cards.
- 1 The Sackville-Bagginses | draw | Creature-based card advantage for longer games.
- 1 Thrór's Map | draw | Provides artifact-based card selection and value.
- 1 Óin the Brave | draw | Legendary Dwarf that contributes cards while developing the board.
- 1 Andúril, Flame of the West | interaction | A powerful legendary equipment for protecting and empowering attackers.
- 1 Andúril, Narsil Reforged | interaction | Versatile legendary equipment supporting combat-based play.
- 1 Bilbo's Ring | interaction | Useful protective equipment for key creatures.
- 1 Dwarven Mattock | interaction | Equipment-based utility that supports creature combat.
- 1 Mithril Coat | interaction | Protects an important legendary creature from removal.
- 1 My Precious // Allure of Power | interaction | Flexible legendary equipment with an interactive adventure.
- 1 Sting, Bilbo's Sword | interaction | Low-cost equipment utility for combat and creature support.
- 1 The One Ring | interaction | Provides protection and resilient long-game value.
- 1 Azog, Moria's Ruin | removal | Thematic creature removal attached to a threat.
- 1 Battle-Scarred Goblin | removal | Creature-based removal that remains useful on the battlefield.
- 1 Bilbo's Deadly Slice | removal | Efficient targeted removal.
- 1 Bitter Downfall | removal | Flexible instant-speed answer to troublesome permanents.
- 1 Fire of Orthanc | removal | Red removal for creatures or artifacts.
- 1 Goblin Cratermaker | removal | Versatile answer to small creatures and artifacts.
- 1 Improvised Club | removal | Cheap instant-speed creature interaction.
- 1 Orcish Bowmasters | removal | Punishes opposing card draw while picking off small creatures.
- 1 Smite the Deathless | removal | Direct answer to durable opposing threats.
- 1 Bolg, Erebor's Reckoning | wipe | A thematic creature-based battlefield reset.
- 1 Call Forth the Tempest | wipe | Sweeps away developed opposing boards.
- 1 Languish | wipe | Efficient mass creature removal.
- 1 Desert Were-Worm | threat | Large thematic body that presents a meaningful combat clock.
- 1 Dreaded Bat-Cloud | threat | Evasive creature pressure fitting the darker mountain theme.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Flexible legendary threat with a useful adventure.
- 1 Gollum the Abandoned | threat | A dangerous legendary creature that pressures opponents.
- 1 Great Goblin, Foul-Hearted | threat | Thematic Goblin leader that strengthens an aggressive board.
- 1 Guttersnipe | threat | Converts the deck’s spells into repeated damage.
- 1 Haunt of the Dead Marshes | threat | Menacing black creature for board presence.
- 1 Inferno Titan | threat | Immediate damage plus a powerful finishing body.
- 1 Olog-hai Crusher | threat | Large combat threat for closing games.
- 1 Sauron, the Lidless Eye | threat | A formidable legendary threat for the late game.
- 1 The Great Goblin | threat | A Goblin-focused legendary threat.
- 1 Troll of Khazad-dûm | threat | A large thematic creature that can pressure life totals.
- 1 Along the Crooked Way | synergy | Supports the deck’s sinister Middle-earth theme and value plan.
- 1 Down, Down to Goblin-town | synergy | Builds around the Goblin and mountain-dwelling side of the deck.
- 1 Getaway Barrel | synergy | Artifact utility that works with the deck’s Treasure and equipment elements.
- 1 Glamdring | synergy | Legendary equipment supporting the deck’s combat plan.
- 1 Last Light of Durin's Day | synergy | Thematic enchantment value for the deck’s legendary cast.
- 1 Long-Lost Lances | synergy | Equipment that rewards committing to the battlefield.
- 1 Smaug's Fury | synergy | A thematic spell that reinforces the dragon-led game plan.
- 1 Supper for Spiders | synergy | Directly supports the deck’s themed creature strategy.
- 1 Tidings of War | synergy | Develops the board and supports an aggressive finish.
- 1 Well-Worn Spatula | synergy | Additional equipment synergy for creatures and artifact payoffs.

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 1.

Cost: $469.89 to buy, $469.89 the whole deck.

**Summary:** This is a Sultai Elf-tribal Commander deck built around developing a broad woodland court, protecting its key legends, and leveraging creature-based mana and card advantage. It uses efficient answers to keep opposing threats contained, then turns a growing force of Elves into sustained combat pressure backed by powerful tribal sweepers and resilient value pieces.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. more opening hands hold two to four lands than the norm, and that raises the grade. the mana base serves one color far better than another, and that lowers the grade.

- [BLOCK] `copy_limit`: Mirkwood Nurturer: 2 copies, the limit is 1
- [INFO] `curve_summary`: average mana value 3.29 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.46, and bracket 3 wants 0.85 or more (sources of requirement: U 13.75 of 19, B 12 of 26, G 22.25 of 22)
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 17 Forest | land | Heavy green base for Elf spells and mana development.
- 10 Island | land | Reliable blue access for card advantage and protective interaction.
- 9 Swamp | land | Supports the deck's efficient black answers.
- 1 Sol Ring | ramp | Efficient artifact acceleration retained as requested.
- 1 Arcane Signet | ramp | Fixes all of the commander's colors while accelerating the deck.
- 1 Mox Amber | ramp | Early acceleration supported by the deck's legendary creatures.
- 1 Wayfarer's Bauble | ramp | Finds a basic land and improves color consistency.
- 1 Elvish Mystic | ramp | Cheap Elf acceleration that contributes to the creature theme.
- 1 Elvish Archdruid | ramp | Tribal mana engine for the large Elf contingent.
- 1 Silvan Reveler | ramp | Creature-based acceleration that works with Elf synergies.
- 1 Wood Elves | ramp | Develops mana while adding another Elf body.
- 1 Thranduil the Strategist | ramp | On-theme mana development from an Elf legend.
- 1 Through the Forest Gate | ramp | Land acceleration that strengthens green access.
- 1 Elvish Visionary | draw | An Elf body that replaces itself.
- 1 Hithlain Knots | draw | Flexible instant-speed card advantage.
- 1 Ithilien Kingfisher | draw | Creature-based card advantage for the board plan.
- 1 Key to the Side-Door | draw | Provides repeatable access to additional cards.
- 1 Last March of the Ents | draw | A powerful refill for a creature-focused strategy.
- 1 Lórien Revealed | draw | Early landcycling or later card selection.
- 1 Long Lake Nuisance | draw | Adds card advantage on an evasive creature.
- 1 Palantír of Orthanc | draw | Steady artifact-based card advantage.
- 1 Plunder the Trollshaws | draw | Instant-speed card advantage.
- 1 Rage into the Valley | draw | Refuels the hand while remaining on theme.
- 1 Reverent Howl | draw | Efficient draw tied to the creature plan.
- 1 Confusticate and Bebother | interaction | Protects the board and disrupts opposing plays.
- 1 Elrond, Moon-Reader | interaction | Legendary Elf interaction with strong thematic overlap.
- 1 Dwarven Mattock | interaction | Equipment-based utility against problematic permanents.
- 1 Mithril Coat | interaction | Protects the commander or a key threat.
- 1 My Precious // Allure of Power | interaction | Versatile protection and disruption from one card.
- 1 Stern Scolding | interaction | Low-cost answer to efficient opposing creatures.
- 1 The One Ring | interaction | Protective legendary utility with sustained value.
- 1 Thranduil's Decree | interaction | On-theme instant-speed disruption.
- 1 Bilbo's Deadly Slice | removal | Efficient black spot removal.
- 1 Bitter Downfall | removal | Broad instant-speed answer to a dangerous permanent.
- 1 Colossal Whale | removal | A large threat that removes opposing creatures.
- 1 Crude Bent Blade | removal | Equipment that turns creatures into removal tools.
- 1 Enchanted River's Grasp | removal | Blue permanent-based removal.
- 1 Giant's Boulder | removal | Colorless removal utility available to the whole deck.
- 1 Merciless Executioner | removal | Creature-based sacrifice removal.
- 1 Orcish Bowmasters | removal | Efficient creature removal with incidental pressure.
- 1 Troll Negotiations | removal | Additional black answer for troublesome threats.
- 1 Gnashing of Teeth | wipe | Creature sweep that can reset a crowded board.
- 1 Languish | wipe | Reliable black sweep against creature boards.
- 1 Raise the Palisade | wipe | Tribal asymmetrical board reset that favors Elves.
- 1 Boughside Wanderers | synergy | Elf body that advances the tribe's battlefield presence.
- 1 Cantankerous Keepers | synergy | Elf creature that benefits from the tribal core.
- 1 Elven Raft-Steerer | synergy | Adds an Elf body and supports the creature plan.
- 1 Galadhrim Guide | synergy | Elf support creature for building a coordinated board.
- 1 Galion, Elvenking's Butler | synergy | Legendary Elf support for the commander's court.
- 1 Grey Havens Navigator | synergy | Elf utility creature that reinforces the theme.
- 1 Lothlórien Lookout | synergy | Early Elf presence for tribal payoffs.
- 1 Mirkwood Meditator | synergy | Elf support body for the wider board strategy.
- 1 Mirkwood Nurturer | synergy | Thematic Elf creature that supports development.
- 1 Supper for Spiders | synergy | Flexible tribal payoff for the creature-heavy plan.
- 1 Arwen, Weaver of Hope | threat | Resilient legendary Elf that strengthens the team.
- 1 Celeborn the Wise | threat | High-impact Elf legend and a natural centerpiece threat.
- 1 Elvenking's Harper | threat | Elf creature that adds to the deck's attacking board.
- 1 Guardian of the Halls | threat | Sturdy Elf threat for combat-focused games.
- 1 Mirkwood Elk | threat | A substantial green creature for applying pressure.
- 1 Mirkwood Pathmaker | threat | Elf threat that advances the tribal battlefield.
- 1 Nimrodel Watcher | threat | Elf body that helps establish a threatening board.
- 1 Thranduil's Company | threat | A thematic Elf force that expands battlefield pressure.
- 1 Thranduil, Sindarin Liege // Silvan Rally | threat | Legendary Elf threat with a valuable rally option.
- 1 Willow-Wind | threat | Large green threat for closing games through combat.
- 1 Woodland Weavemaster | threat | Elf creature that contributes meaningful board presence.
- 1 Mirkwood Nurturer | threat | Thematic Elf creature that supports development.

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $312.00 to buy, $312.00 the whole deck.

**Summary:** This is a white creature-and-artifact deck centered on Kíli, with Dwarves, Halflings, Rabbits, and equipment creating a broad battlefield. It develops its resources through efficient artifacts and creature-based support, keeps opposing boards in check with targeted answers and resets, then wins by applying combat pressure with large creatures, legendary cards, and equipped attackers. The tradeoff is a deliberately board-centric plan that depends on keeping permanents in play rather than relying on explosive mana or extensive library searching.

The quality model grades this deck good: many of the cards appear in no top list, and that lowers the grade. the commander does not place in cEDH events, and that lowers the grade. few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.90 over 63 nonland cards
- [WARN] `summary_rules_claim`: the summary states a rule of the game, "the format play". The engine reports the rules, and the summary must not.

<details><summary>The deck list</summary>

- 1 Belladonna Took | draw | Draw piece that supports the deck's steady resource plan.
- 1 Caretaker's Talent | draw | Draw engine for a board-focused creature deck.
- 1 Circuit Mender | draw | Artifact-bodied draw option that fits the deck's permanent-heavy build.
- 1 Errand-Rider of Gondor | draw | Creature-based draw that adds to the board.
- 1 Fountainport Bell | draw | Artifact draw piece for the deck's resource package.
- 1 Idol of Oblivion | draw | Efficient artifact draw option.
- 1 Inspiring Overseer | draw | Creature-based draw that helps maintain board presence.
- 1 Mentor of the Meek | draw | Draw engine for the deck's many smaller creatures.
- 1 Skullclamp | draw | Low-cost draw equipment for a creature-heavy board.
- 1 Spirited Companion | draw | Early creature that contributes to the draw package.
- 1 The Gaffer | draw | Legendary draw threat that supports longer games.
- 1 Baird, Steward of Argive | interaction | Interaction piece that helps discourage opposing pressure.
- 1 Bilbo's Gambit | interaction | Flexible instant-speed interaction.
- 1 Dawn's Truce | interaction | Protective interaction for the deck's established board.
- 1 Galadriel's Dismissal | interaction | Instant interaction that can answer a difficult board state.
- 1 Reprieve | interaction | Low-cost interaction that preserves tempo.
- 1 Selfless Spirit | interaction | Creature-based protection for the deck's board.
- 1 Swiftfoot Boots | interaction | Equipment interaction that helps protect a key creature.
- 1 Stone by Sunlight | interaction | Versatile instant-speed interaction.
- 1 Arcane Signet | ramp | Reliable early mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Mind Stone | ramp | Early acceleration that remains useful later.
- 1 Mox Amber | ramp | Fast legendary-focused mana acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment-based acceleration that fits the artifact plan.
- 1 Ornithopter of Paradise | ramp | Creature-based mana acceleration.
- 1 Patchwork Banner | ramp | Artifact acceleration aligned with the creature theme.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Thought Vessel | ramp | Artifact mana acceleration for the midgame.
- 1 Wayfarer's Bauble | ramp | Low-cost mana development.
- 1 Banishing Light | removal | Broad permanent removal.
- 1 Fiend Hunter | removal | Creature-based removal that contributes to board presence.
- 1 Generous Gift | removal | Flexible answer to a problematic permanent.
- 1 Loran of the Third Path | removal | Legendary artifact-focused removal option.
- 1 Nettle Guard | removal | Creature removal attached to a Mouse body.
- 1 Skyclave Apparition | removal | Efficient creature-based removal.
- 1 Swords to Plowshares | removal | Low-cost targeted removal.
- 1 The Black Arrow | removal | Equipment-based removal that supports the artifact package.
- 1 Westfold Rider | removal | Creature-based removal for the board plan.
- 1 Academy Manufactor | synergy | Artifact creature that strengthens the deck's artifact-centered permanent package.
- 1 Andúril, Narsil Reforged | synergy | Legendary equipment that supports the deck's equipment theme.
- 1 Blade Splicer | synergy | Artifact-adjacent creature that contributes to the board plan.
- 1 Carrot Cake | synergy | Food artifact that supports the deck's artifact package.
- 1 Dáin, Lord of the Iron Hills | synergy | Legendary Dwarf that anchors the Dwarf portion of the deck.
- 1 Iron Hills Blacksmith | synergy | Dwarf Artificer that joins the artifact and Dwarf themes.
- 1 Maskwood Nexus | synergy | Artifact that ties together the deck's varied creature types.
- 1 Ori, Keeper of Songs | synergy | Legendary Dwarf that supports the deck's tribal core.
- 1 Tangle Tumbler | synergy | Artifact Vehicle that reinforces the artifact package.
- 1 Three Tree Mascot | synergy | Artifact creature that supports the deck's creature-type cohesion.
- 1 Dusk // Dawn | wipe | Board reset for games where opposing creatures get ahead.
- 1 Martial Coup | wipe | Sweep effect that also supports rebuilding a creature board.
- 1 Promise of Loyalty | wipe | Board-control reset for difficult creature-heavy games.
- 1 Angel of the Ruins | threat | Large artifact creature threat for closing games.
- 1 Andúril, Flame of the West | threat | Legendary equipment that turns a creature into a meaningful threat.
- 1 Boss's Chauffeur | threat | Creature threat that adds weight to the battlefield.
- 1 Fíli the Pathfinder | threat | Legendary Dwarf threat that fits the deck's tribal focus.
- 1 Helm of the Host | threat | Equipment threat that rewards maintaining a creature board.
- 1 Jazal Goldmane | threat | Legendary creature threat for the deck's combat plan.
- 1 Karn, the Great Creator | threat | Planeswalker threat that adds a durable noncreature angle.
- 1 Psychosis Crawler | threat | Artifact creature threat aligned with the deck's draw package.
- 1 Realm-Cloaked Giant // Cast Off | threat | Large creature threat with a flexible second role.
- 1 Serra Redeemer | threat | Flying creature threat that supports the permanent-heavy plan.
- 1 Sunscorch Regent | threat | Large Dragon threat for finishing games.
- 1 Warren Warleader | threat | Rabbit creature threat that broadens the combat plan.
- 1 Castle Ardenvale | land | Nonbasic land slot in the white mana base.
- 1 Command Tower | land | Reliable color-fixing land.
- 1 Elven Passage | land | Nonbasic land slot supporting the deck's creature base.
- 1 Evolving Wilds | land | Land slot that improves mana development.
- 1 Exotic Orchard | land | Color-fixing nonbasic land.
- 1 Fabled Passage | land | Land slot that improves mana development.
- 1 Hidden Grotto | land | Nonbasic land slot for the deck's mana base.
- 1 Hobbit Hole | land | Thematic nonbasic land slot.
- 1 Lupinflower Village | land | Bloomburrow-themed nonbasic land slot.
- 1 Minas Tirith | land | Thematic legendary land slot.
- 1 Path of Ancestry | land | Land slot that supports the deck's creature-type focus.
- 1 Reliquary Tower | land | Utility land slot for resource-heavy games.
- 1 Rogue's Passage | land | Utility land that supports the combat plan.
- 1 Swarmyard | land | Utility land that fits several of the deck's creature types.
- 1 Three Tree City | land | Legendary land aligned with the deck's creature-type theme.
- 1 Thriving Heath | land | Color-fixing nonbasic land.
- 1 Treasure Vault | land | Artifact land that reinforces the artifact package.
- 1 Terramorphic Expanse | land | Land slot that improves mana development.
- 1 Uncharted Haven | land | Color-fixing nonbasic land.
- 17 Plains | land | Basic white mana base for the deck's white cards.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $32.64 to buy, $32.64 the whole deck.

**Summary:** This mono-red aggro deck establishes pressure with Mouse, Lizard, Raccoon, and Otter creatures, then keeps attacking while its removal clears the way. Emberheart Challenger and Hearthborn Battler supply the central synergy pieces, while the draw cards help maintain momentum through a longer game. It gives up broad answers and large late-game threats in favor of a focused, low-land-pressure attacking plan.

The quality model grades this deck below the precon baseline: many of the cards appear in no top list, and that lowers the grade. the cards pair in ways the top lists do not, and that lowers the grade. the deck runs few spells as one copy, and that raises the grade.

- [INFO] `curve_summary`: average mana value 3.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides a consistent untapped mono-red mana base.
- 2 Brazen Collector | ramp | Adds the requested ramp while remaining part of the creature plan.
- 2 Artist's Talent | draw | Provides draw support for the aggressive game plan.
- 2 Might of the Meek | draw | Supplies efficient draw alongside the deck's small creatures.
- 2 Sazacap's Brew | draw | Rounds out the deck's draw package.
- 4 Abrade | removal | Provides flexible early removal.
- 2 Agate Assault | removal | Adds more low-cost removal for opposing threats.
- 4 Emberheart Challenger | synergy | Forms a core Mouse synergy package.
- 4 Hearthborn Battler | synergy | Adds a second aggressive synergy-focused creature package.
- 4 Frilled Sparkshooter | threat | Provides a full set of Lizard threats for the attack plan.
- 4 Reptilian Recruiter | threat | Supplies efficient creature pressure.
- 4 Teapot Slinger | threat | Adds another set of attacking threats.
- 2 Stormsplitter | threat | Completes the threat suite with Otter Wizard pressure.

</details>

