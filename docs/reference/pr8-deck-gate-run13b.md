# PR-8 deck gate

Run date: 2026-09-03. Card snapshot: 2026-09-02.

Verdict: PASS. 24 of 24 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 24 |
| Decks returned | 24 |
| Decks with no block finding | 24 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 13 |
| Summaries judged (F-26) | 24 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 66 |
| Cost | $2.6201 |
| Time | 2150 seconds |

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `curve_summary` | 24 |
| `not_owned` | 23 |
| `profile_off_band` | 9 |
| `basics_added` | 3 |
| `outside_requested_set` | 2 |
| `two_card_combo` | 1 |
| `cards_trimmed` | 1 |

By severity: BLOCK 0. WARN 35. INFO 28. 

## The set filter (PR-17B)

A deck asked for a set holds cards of that set family alone. Three things earn an exception, and each one is marked: a card the reader named (D-381), a mana card the reader allowed from outside (D-382), and a basic land, which no set limit filters (D-378).

| # | Prompt | Sets | In set | Fill | Deck cards outside | Marked |
|---|---|---|---|---|---|---|
| 19 | the Hobbit family, two colours | `hob,hoc` | 170 | 0 | 0 | 0 |
| 20 | the Hobbit family, a delegated commander | `hob,hoc` | 75 | 0 | 0 | 0 |
| 21 | the Hobbit family, mana from outside | `hob,hoc` | 127 | 26 | 15 | 15 |
| 22 | a set family and a card from outside it | `hob,hoc` | 170 | 0 | 1 | 1 |
| 23 | two set families at once | `blb,blc,hob,hoc,pblb` | 201 | 0 | 0 | 0 |
| 24 | a 60-card deck from one set | `blb,blc,pblb` | 95 | 0 | 0 | 0 |

## Decks

### 1. lifegain Commander, any card

Format: Commander. Theme: lifegain. Pool: any_card. Shortlist: 303 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $446.28 to buy, $446.28 the whole deck.

**Summary:** A black-white lifegain deck that develops efficient mana, repeatedly gains small amounts of life, and turns those triggers into growing creatures, card advantage, and life-drain pressure. It protects its central threats, controls opposing boards with flexible answers, and can pivot from early combat to a resilient late-game battlefield.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 2.94 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | Basic white source for early lifegain spells and the commander.
- 9 Swamp | land | Basic black source for drain effects and removal.
- 1 Caves of Koilos | land | Untapped dual land for both deck colors.
- 1 Command Tower | land | Reliable untapped fixing.
- 1 Exotic Orchard | land | Usually provides either required color in Commander.
- 1 Godless Shrine | land | Fetchable dual source that can enter untapped.
- 1 Mana Confluence | land | Untapped access to either deck color.
- 1 Reflecting Pool | land | Efficient fixing alongside the dual-heavy mana base.
- 1 Shattered Sanctum | land | Untapped dual land through much of the game.
- 1 Vault of Champions | land | Multiplayer dual land that normally enters untapped.
- 1 Cavern of Souls | land | Protects key creature casts while remaining a mana source.
- 1 Spoils of Evil | ramp | Very cheap burst mana that can accelerate pivotal turns.
- 1 Bounty Board | ramp | Low-cost artifact acceleration.
- 1 Colossal Plow | ramp | Early mana production with creatures available to crew it.
- 1 Hot Dog Cart | ramp | Low-cost mana development that supports the deck's artifact package.
- 1 Oasis Gardener | ramp | Early creature-based mana fixing and acceleration.
- 1 Altar of the Pantheon | ramp | Fixes both colors while advancing mana.
- 1 Nuka-Cola Vending Machine | ramp | Produces mana and supplies incidental life-oriented resources.
- 1 Orazca Relic | ramp | Color fixing early with card value available later.
- 1 Phial of Galadriel | ramp | Mana rock that contributes to the deck's card-flow plan.
- 1 Pristine Talisman | ramp | Reliable mana rock whose life gain triggers lifegain payoffs.
- 1 Archivist of Oghma | draw | Efficient recurring card advantage against library searching.
- 1 Dawn of Hope | draw | Turns repeated life gain into cards and creature production.
- 1 Enduring Innocence | draw | Sustained card draw from small creatures entering the battlefield.
- 1 Frodo, Adventurous Hobbit | draw | Low-cost source of incremental cards in a creature-focused game.
- 1 Markov Purifier | draw | Draw engine rewarded by Vampire and lifegain synergies.
- 1 Pearl-Ear, Imperial Advisor | draw | Provides cards while supporting the deck's enchantment elements.
- 1 Scheming Silvertongue // Sign in Blood | draw | Flexible early card draw with a creature body afterward.
- 1 The Gaffer | draw | Steady draw when the deck is gaining life consistently.
- 1 Vampiric Rites | draw | Cheap sacrifice outlet that converts creatures into cards and life.
- 1 Well of Lost Dreams | draw | Converts larger lifegain events into substantial card advantage.
- 1 Exemplar of Light | draw | Creature-based card draw that benefits from a high life total.
- 1 Alseid of Life's Bounty | interaction | One-mana protection for the commander or a key payoff.
- 1 Courageous Resolve | interaction | Protects an important permanent from targeted disruption.
- 1 Faith's Shield | interaction | Flexible protection that can also force through combat damage.
- 1 Distinguished Conjurer | interaction | Creature-based protection that fits the deck's life theme.
- 1 Metropolis Reformer | interaction | Protects the player and key creatures from harmful targeting.
- 1 Restoration Magic | interaction | Versatile defensive spell for preserving the board.
- 1 Sword of Light and Shadow | interaction | Protection plus recurring creature value.
- 1 Werefox Bodyguard | interaction | Interactive creature that temporarily answers a problem permanent.
- 1 Ayli, Eternal Pilgrim | removal | Efficient sacrifice outlet and repeatable removal at high life totals.
- 1 Gumdrop Poisoner // Tempt with Treats | removal | Early creature removal with a lifegain-oriented alternative mode.
- 1 Murderous Rider // Swift End | removal | Flexible answer to creatures and planeswalkers.
- 1 Nightmare's Thirst | removal | Cheap creature removal improved by the deck's life gain.
- 1 Solitude | removal | Immediate answer to dangerous creatures, including at instant speed.
- 1 Starseer Mentor | removal | Creature-based removal that remains on-theme.
- 1 Syr Vondam, Sunstar Exemplar | removal | Repeatable removal attached to a lifelinking creature.
- 1 Umezawa's Jitte | removal | Efficient equipment that controls small creatures and rewards combat.
- 1 Vona, Butcher of Magan | removal | Repeatable permanent removal fueled by a healthy life total.
- 1 Ajani's Pridemate | synergy | Classic scalable threat that grows from every lifegain trigger.
- 1 Cleric Class | synergy | Enhances life gain and supplies a late-game payoff.
- 1 Essence Channeler | synergy | Efficient lifegain trigger source for the commander and payoffs.
- 1 Heliod, Sun-Crowned | synergy | Distributes counters from life gain and supports the creature plan.
- 1 Indulgent Aristocrat | synergy | Low-cost lifelink creature and sacrifice outlet with team growth.
- 1 Marauding Blight-Priest | synergy | Turns each lifegain event into pressure on opponents.
- 1 Ocelot Pride | synergy | Lifelink threat that snowballs into a token board.
- 1 Serra Ascendant | synergy | Early lifelink attacker that exploits Commander starting life totals.
- 1 Voice of the Blessed | synergy | Rapidly grows into a protected evasive lifegain payoff.
- 1 Vito, Thorn of the Dusk Rose | synergy | Converts life gain into direct life-loss pressure.
- 1 Attended Healer | threat | Creates a board presence from repeated lifegain triggers.
- 1 Celestine, the Living Saint | threat | Recurring creatures turns lifegain into lasting battlefield advantage.
- 1 Doctor Jane Foster | threat | Lifelink-oriented legend that pressures opponents while supporting the plan.
- 1 Enduring Tenacity | threat | Drains opponents whenever life is gained.
- 1 Exalted Sunborn | threat | Evasive lifelinking finisher that stabilizes combat.
- 1 Fiendish Panda | threat | Efficient creature threat that benefits from the deck's life-total plan.
- 1 Minwu, White Mage | threat | Low-cost legendary threat with useful lifegain synergy.
- 1 Qala, Ajani's Pridemate | threat | Lifegain-scaling legend that becomes a meaningful attacker quickly.
- 1 Star Charter | threat | Creature payoff that builds pressure from recurring lifegain.
- 1 Twinblade Paladin | threat | Double-striking lifelink attacker that scales well with buffs.
- 1 Valkyrie Harbinger | threat | Creates evasive threats once the deck gains life in larger chunks.
- 1 Wurmcoil Engine | threat | Resilient lifelink threat that leaves material behind after removal.
- 1 Amalia Benavides Aguirre | wipe | Low-cost lifegain payoff capable of clearing oversized boards.
- 1 Kaya's Wrath | wipe | Reliable full board reset that also restores life.
- 1 White Sun's Twilight | wipe | Flexible sweeper that can leave behind a stabilizing board.

</details>

### 2. aristocrats Commander, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 179 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $161.49 the whole deck.

**Summary:** A black-white aristocrats deck built around disposable creatures, recurring graveyard value, and incremental death-trigger pressure. It develops through efficient mana rocks and utility creatures, protects its central pieces, and uses flexible removal to keep opponents from stabilizing. The deck closes by grinding through creature exchanges, leveraging equipment and token-friendly value engines, and rebuilding effectively after a crowded board is cleared.

The quality model grades this deck good against the top lists of the format: the commander does not place in cEDH events, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves one color far better than another, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.81 over 63 nonland cards

<details><summary>The deck list</summary>

- 17 Plains | land | Reliable white mana for the deck’s creature and protection spells.
- 10 Swamp | land | Reliable black mana for sacrifice outlets and death-trigger spells.
- 1 Command Tower | land | Flexible commander-color mana.
- 1 Exotic Orchard | land | Flexible multiplayer color fixing.
- 1 City of Brass | land | Unrestricted color fixing.
- 1 Grand Coliseum | land | Broad color access.
- 1 Opal Palace | land | Commander-focused color fixing.
- 1 Plaza of Heroes | land | Supports legendary permanents while producing needed colors.
- 1 Spire of Industry | land | Artifact-supported color fixing.
- 1 Takenuma, Abandoned Mire | land | Black source with late-game recursion utility.
- 1 Minas Tirith | land | White source with card-advantage utility.
- 1 Arcane Signet | ramp | Efficient color fixing.
- 1 Commander's Sphere | ramp | Fixes mana and can later become a card.
- 1 Fellwar Stone | ramp | Efficient multiplayer mana rock.
- 1 Sol Ring | ramp | Fast colorless acceleration.
- 1 Thought Vessel | ramp | Mana acceleration with hand-size utility.
- 1 Wayfarer's Bauble | ramp | Early land-based acceleration.
- 1 Chromatic Lantern | ramp | Strong fixing for the two-color mana base.
- 1 Astral Cornucopia | ramp | Flexible scalable mana production.
- 1 Relic of Legends | ramp | Turns legendary creatures into mana sources.
- 1 Sword of the Animist | ramp | Equipment-based land ramp for attacking creatures.
- 1 Call of the Ring | draw | Steady card flow that complements the deck’s life-resource plan.
- 1 Grave Venerations | draw | Draw support tied to the graveyard-focused game plan.
- 1 Idol of Oblivion | draw | Efficient repeatable draw alongside token production.
- 1 Inspiring Overseer | draw | Creature-based card advantage with incidental life gain.
- 1 Lembas | draw | Low-cost card replacement that can be reused.
- 1 Mask of Memory | draw | Combat-based filtering and card advantage.
- 1 Night's Whisper | draw | Cheap, direct card advantage.
- 1 Painful Truths | draw | Efficient black card draw.
- 1 Skullclamp | draw | Converts expendable creatures into substantial card advantage.
- 1 Tome of Legends | draw | Reliable long-game card advantage.
- 1 Wall of Omens | draw | Early defense that replaces itself.
- 1 Bastion Protector | interaction | Keeps the commander protected while advancing the board.
- 1 Boromir, Warden of the Tower | interaction | Disrupts opposing explosive plays and protects the board.
- 1 Champion's Helm | interaction | Protects the commander and other legendary threats.
- 1 Clever Concealment | interaction | Preserves the board through opposing sweepers.
- 1 Darksteel Plate | interaction | Durable protection for a key creature.
- 1 Gift of Immortality | interaction | Returns an important sacrifice piece after it dies.
- 1 Lightning Greaves | interaction | Efficient haste and protection for key creatures.
- 1 Swiftfoot Boots | interaction | Additional haste and targeted protection.
- 1 Banishing Light | removal | Versatile answer to problematic nonland permanents.
- 1 Bitter Triumph | removal | Cheap, flexible creature or planeswalker removal.
- 1 Claim the Precious | removal | Clean answer to opposing creatures.
- 1 Crib Swap | removal | Exile-based creature removal.
- 1 Destroy Evil | removal | Efficient answer to large creatures and key enchantments.
- 1 Dispatch | removal | Low-cost creature interaction that improves with artifacts.
- 1 Fiend Hunter | removal | Creature-based removal that can be exploited with sacrifice effects.
- 1 Generous Gift | removal | Answers any troublesome permanent.
- 1 Infernal Grasp | removal | Straightforward unconditional creature removal.
- 1 Austere Command | wipe | Flexible reset that can spare the preferred board texture.
- 1 Fumigate | wipe | Creature reset with meaningful life recovery.
- 1 Vanquish the Horde | wipe | Efficient creature sweeper when the board is crowded.
- 1 Gollum the Abandoned | synergy | A recursive sacrifice-oriented creature for repeated value.
- 1 Gollum, Patient Plotter | synergy | Returns from the graveyard and rewards disposable creatures.
- 1 Gríma Wormtongue | synergy | Supports the deck’s attrition-focused sacrifice plan.
- 1 Phantom Train | synergy | Provides a resilient artifact-based value piece.
- 1 Arcade Cabinet | synergy | Artifact utility that supports the deck’s incremental value plan.
- 1 Blowfly Infestation | synergy | Turns creature deaths into cascading board control.
- 1 Contagion Clasp | synergy | Adds removal utility and supports counter-based attrition.
- 1 Deadly Dispute | synergy | Turns expendable permanents into mana and cards.
- 1 Stone of Erech | synergy | Graveyard pressure that complements death-focused play.
- 1 The Sackville-Bagginses | synergy | Creates value from food and sacrifice-oriented board states.
- 1 Angel of Serenity | threat | Large evasive finisher with impactful enter-the-battlefield value.
- 1 Bill the Pony | threat | Efficient creature that contributes to the board early.
- 1 Bronze Guardian | threat | A scalable artifact threat that also protects key permanents.
- 1 Exemplar of Light | threat | Evasive creature that provides ongoing value.
- 1 Frontline Medic | threat | Combat-focused creature that can protect an attacking board.
- 1 Giada, Font of Hope | threat | Builds a threatening flying board while supporting mana development.
- 1 Massacre Girl, Known Killer | threat | Converts creature deaths into a dangerous card-advantage engine.
- 1 Orcish Bowmasters | threat | Efficient creature that pressures opposing card draw and small creatures.
- 1 Palace Jailer | threat | Provides removal and a strong incentive to control combat.
- 1 Puresteel Paladin | threat | Develops equipment value while providing a capable body.
- 1 Rat King, Pale Piper | threat | A resilient black threat that fits the graveyard plan.
- 1 Vengeful Villagers | threat | Creature pressure that benefits from the deck’s attrition plan.

</details>

### 3. artifacts Commander, bracket 4

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 299 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $3906.04 to buy, $3906.04 the whole deck.

**Summary:** A high-power mono-blue artifact strategy built to develop mana quickly, establish a dense board of artifacts, and turn that infrastructure into relentless pressure. Efficient card advantage, protective countermagic, and flexible answers keep the engine intact while artifact creatures and scalable payoffs close games.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.08 over 66 nonland cards
- [WARN] `profile_off_band`: the average mana value of the nonland cards is 3.08, and bracket 4 wants 1.6 to 3
- [WARN] `profile_off_band`: the mana available on turn four is 4.47, and bracket 4 wants 4.8 or more (mean over 10000 hands)

<details><summary>The deck list</summary>

- 20 Island | land | Reliable blue mana base for early development and holding up interaction.
- 1 Academy Ruins | land | Recurs key artifacts after removal.
- 1 Buried Ruin | land | Artifact recursion from the mana base.
- 1 Inventors' Fair | land | Finds an important artifact while supporting the artifact count.
- 1 Mishra's Workshop | land | Explosive artifact mana for deploying the deck’s core permanents.
- 1 Urza's Workshop | land | Artifact-focused mana production.
- 1 Seat of the Synod | land | Blue source that also contributes to artifact synergies.
- 1 Mystic Sanctuary | land | Blue source with late-game spell recursion.
- 1 Otawara, Soaring City | land | Flexible blue source and uncounterable-style utility interaction.
- 1 Command Tower | land | Reliable untapped blue source.
- 1 Glimmervoid | land | Efficient colored mana in an artifact-heavy deck.
- 1 Spire of Industry | land | Colored mana source that is consistently active with artifacts in play.
- 1 City of Brass | land | Untapped blue fixing.
- 1 Mana Confluence | land | Untapped blue fixing.
- 1 Mox Opal | ramp | Fast artifact mana that accelerates early development.
- 1 Moonsnare Prototype | ramp | Low-cost mana production that turns artifacts into acceleration.
- 1 Arc Reactor | ramp | Early artifact mana for faster commander and spell deployment.
- 1 Crystal Skull, Isu Spyglass | ramp | Mana acceleration with useful artifact utility.
- 1 Chief Engineer | ramp | Convoke-style artifact acceleration for artifact spells.
- 1 Vedalken Engineer | ramp | Efficient mana creature dedicated to casting artifacts.
- 1 Grand Architect | ramp | Turns blue creatures into substantial artifact mana.
- 1 Metalworker | ramp | High-output artifact mana engine.
- 1 Coin of Mastery | ramp | Artifact mana source that advances the board early.
- 1 Cargo Ship | ramp | Mana-producing artifact that supports the artifact density.
- 1 Network Terminal | ramp | Reliable artifact mana production.
- 1 Powerstone Shard | ramp | Artifact ramp that scales with copied or recurring artifacts.
- 1 Solar Array | ramp | Additional artifact acceleration for explosive midgames.
- 1 Cerebral Download | draw | Efficient card flow for an artifact-heavy hand.
- 1 Era of Innovation | draw | Artifact and creature development converts into cards.
- 1 Esoteric Duplicator | draw | Clue-based card advantage that fits the artifact shell.
- 1 Forensic Gadgeteer | draw | Artifact-focused value and card advantage.
- 1 Riddlesmith | draw | Filters draws whenever artifacts are deployed.
- 1 Sai, Master Thopterist | draw | Creates material and converts spare artifacts into cards.
- 1 Thirst for Knowledge | draw | Efficient instant-speed card selection using artifacts.
- 1 Vedalken Archmage | draw | Turns artifact spells into sustained card advantage.
- 1 Waterbending Lesson | draw | Low-cost card advantage.
- 1 Katara, Bending Prodigy | draw | Creature-based source of ongoing cards.
- 1 Edgar, King of Figaro | draw | Artifact-centric card advantage engine.
- 1 Braided Net // Braided Quipu | draw | Flexible artifact-based draw resource.
- 1 Disruption Protocol | interaction | Artifact-enabled protection against opposing spells.
- 1 Metallic Rebuke | interaction | Efficient counterspell in an artifact-dense list.
- 1 Stoic Rebuttal | interaction | Reliable hard counterspell with metalcraft.
- 1 Ice Out | interaction | Flexible countermagic that becomes efficient with artifacts.
- 1 Curator's Ward | interaction | Protects a critical permanent while replacing itself.
- 1 Welding Jar | interaction | Zero-cost protection for essential artifacts.
- 1 Fugitive Droid | interaction | Artifact creature that provides disruptive utility.
- 1 Escape Protocol | interaction | Protects key artifacts and enables value loops.
- 1 Ghostly Flicker | interaction | Protects permanents while retriggering useful abilities.
- 1 Reality Ripple | interaction | Low-cost flexible disruption.
- 1 Waterbender's Restoration | interaction | Instant-speed defensive interaction.
- 1 Aether Spellbomb | removal | Cheap, recurrable answer that remains an artifact.
- 1 Contagion Clasp | removal | Early permanent answer with useful proliferate utility.
- 1 Cyber Conversion | removal | Efficient answer to problematic creatures.
- 1 Into Thin Air | removal | Tempo-positive instant-speed removal.
- 1 Resculpt | removal | Versatile answer to creatures or artifacts.
- 1 Ravenform | removal | Exiles troublesome creatures or artifacts.
- 1 Unable to Scream | removal | Low-cost creature neutralization.
- 1 Water Whip | removal | Efficient targeted removal.
- 1 Watery Grasp | removal | Permanent-based answer to a major threat.
- 1 Zuko's Exile | removal | Low-cost removal option.
- 1 Engineered Explosives | wipe | Flexible scalable sweeper that works well with artifact recursion.
- 1 Hurkyl's Recall | wipe | Fast mass reset against opposing artifact boards.
- 1 Emry, Lurker of the Loch | synergy | Recurs cheap artifacts and rewards a dense artifact graveyard.
- 1 Etherium Sculptor | synergy | Reduces artifact costs to improve explosive turns.
- 1 Foundry Inspector | synergy | Additional artifact cost reduction.
- 1 Mystic Forge | synergy | Lets the deck deploy artifacts from the top of the library.
- 1 Clock of Omens | synergy | Converts excess artifacts into powerful untap utility.
- 1 Voltaic Key | synergy | Cheap untap engine for mana artifacts and utility pieces.
- 1 Filigree Attendant | threat | Artifact payoff creature that applies pressure efficiently.
- 1 Frogmite | threat | Deploys cheaply in a high-artifact board and adds pressure.
- 1 Ironheart, Clever Champion | threat | Artifact-based attacker with meaningful board presence.
- 1 Jhoira's Familiar | threat | Evasive artifact threat that supports historic spells.
- 1 Karn, Scion of Urza | threat | Produces artifact-based threats and supplies continuing value.
- 1 Panther Robot | threat | Efficient artifact creature for board pressure.
- 1 Rust Elemental | threat | Large artifact threat that rewards a stocked board.
- 1 Somber Hoverguard | threat | Evasive artifact payoff that can be deployed ahead of curve.
- 1 Synth Infiltrator | threat | Artifact creature that contributes resilient pressure.
- 1 Traxos, Scourge of Kroog | threat | Large low-cost artifact attacker that untaps through normal development.
- 1 Whirler Rogue | threat | Builds artifact material and makes a key attacker unblockable.
- 1 Darksteel Juggernaut | threat | Scales into a major attacker as the artifact board grows.

</details>

### 4. dinosaur tribal, bracket 2

Format: Commander. Theme: dinosaurs. Pool: any_card. Shortlist: 269 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $328.36 to buy, $328.36 the whole deck.

**Summary:** A Naya Dinosaur deck built around developing reliable mana, filling the battlefield with prehistoric creatures, and turning combat into an overwhelming finishing plan. It balances straightforward creature pressure with card advantage, creature-based answers, and protective spells, giving a new player clear decisions throughout the game.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. The commander is a popular one, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.60 over 62 nonland cards

<details><summary>The deck list</summary>

- 1 Command Tower | land | Reliable three-color fixing.
- 1 Exotic Orchard | land | Flexible multiplayer color fixing.
- 1 Bountiful Promenade | land | Green-white dual land.
- 1 Brushland | land | Green-white source that enters ready.
- 1 Battlefield Forge | land | Red-white source that enters ready.
- 1 Karplusan Forest | land | Red-green source that enters ready.
- 1 Clifftop Retreat | land | Red-white fixing.
- 1 Rootbound Crag | land | Red-green fixing.
- 1 Sunpetal Grove | land | Green-white fixing.
- 1 Rockfall Vale | land | Red-green fixing.
- 1 Sundown Pass | land | Red-white fixing.
- 1 Spire Garden | land | Multiplayer red-green dual land.
- 1 Spectator Seating | land | Multiplayer red-white dual land.
- 1 Sacred Foundry | land | Red-white land with basic land types.
- 1 Stomping Ground | land | Red-green land with basic land types.
- 1 Temple Garden | land | Green-white land with basic land types.
- 1 Canopy Vista | land | Green-white dual land.
- 1 Cinder Glade | land | Red-green dual land.
- 1 Jetmir's Garden | land | Three-color fixing with useful land types.
- 1 Jungle Shrine | land | Three-color fixing.
- 1 Path of Ancestry | land | Tribal mana fixing and incidental selection.
- 1 Unclaimed Territory | land | Reliable Dinosaur tribal fixing.
- 1 Secluded Courtyard | land | Reliable Dinosaur tribal fixing.
- 1 Cavern of Souls | land | Protects key Dinosaur spells from counters.
- 1 Three Tree City | land | Tribal land that can produce substantial mana.
- 1 Reflecting Pool | land | Flexible fixing alongside the multicolor lands.
- 1 Fortified Village | land | Green-white fixing.
- 1 Game Trail | land | Red-green fixing.
- 4 Forest | land | Core green sources for ramp and Dinosaur spells.
- 3 Mountain | land | Core red sources for Dinosaur spells.
- 2 Plains | land | Core white sources for the commander and support spells.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Arcane Signet | ramp | Reliable three-color mana acceleration.
- 1 Birds of Paradise | ramp | Early mana creature that fixes every color.
- 1 Llanowar Elves | ramp | Low-cost green acceleration.
- 1 Farseek | ramp | Finds dual lands with basic land types.
- 1 Nature's Lore | ramp | Finds a Forest dual land untapped.
- 1 Rampant Growth | ramp | Stable basic-land acceleration.
- 1 Cultivate | ramp | Develops mana while ensuring the next land drop.
- 1 Kodama's Reach | ramp | Additional dependable land ramp.
- 1 Thunderherd Migration | ramp | Dinosaur-themed land ramp.
- 1 Esper Sentinel | draw | Early source of steady card advantage.
- 1 Faithless Looting | draw | Filters cards early and can be reused later.
- 1 Demand Answers | draw | Turns spare material into fresh cards.
- 1 Thrill of Possibility | draw | Instant-speed card filtering.
- 1 Beast Whisperer | draw | Rewards the creature-heavy Dinosaur plan.
- 1 Garruk's Uprising | draw | Draws for large creatures and grants trample.
- 1 Harmonize | draw | Simple, dependable card advantage.
- 1 Return of the Wildspeaker | draw | Draws from a large creature or boosts the team.
- 1 Ripjaw Raptor | draw | Dinosaur threat that rewards taking damage with cards.
- 1 Earthshaker Dreadmaw | draw | Large Dinosaur that refills the hand.
- 1 Shamanic Revelation | draw | Rewards building a wide creature board.
- 1 Boros Charm | interaction | Protects the board or finishes combat.
- 1 Ephemerate | interaction | Protects a creature and reuses its enter-the-battlefield ability.
- 1 Heroic Intervention | interaction | Broad protection for the Dinosaur board.
- 1 Snakeskin Veil | interaction | Low-cost creature protection.
- 1 Swiftfoot Boots | interaction | Protects an important creature and grants haste.
- 1 Temple Altisaur | interaction | Reduces damage dealt to the Dinosaur team.
- 1 Itzquinth, Firstborn of Gishath | removal | Efficient Dinosaur removal on a creature.
- 1 Savage Stomp | removal | Efficient fight spell for a Dinosaur deck.
- 1 Thrashing Brontodon | removal | Creature-based artifact and enchantment removal.
- 1 Needletooth Raptor | removal | Dinosaur that can fight opposing creatures.
- 1 Raging Regisaur | removal | Attack trigger picks off smaller creatures.
- 1 Trumpeting Carnosaur | removal | Large Dinosaur with immediate creature removal.
- 1 Ravenous Sailback | removal | Dinosaur that answers artifacts or enchantments.
- 1 Amped Raptor | synergy | Early Dinosaur that supports the tribal plan.
- 1 Belligerent Yearling | synergy | Low-cost Dinosaur that grows with the herd.
- 1 Commune with Dinosaurs | synergy | Finds either a needed land or Dinosaur.
- 1 Deathgorge Scavenger | synergy | Useful Dinosaur with graveyard disruption.
- 1 Dinosaur Egg | synergy | Tribal body that leaves value behind.
- 1 Dromosaur | synergy | Early Dinosaur for tribal development.
- 1 Drowsing Tyrannodon | synergy | Efficient Dinosaur body for the curve.
- 1 Huatli's Raptor | synergy | Early Dinosaur that supports counter themes.
- 1 Kinjalli's Caller | synergy | Reduces the cost of large Dinosaur spells.
- 1 Marauding Raptor | synergy | Cost reduction and growth for Dinosaur casting.
- 1 Nest Robber | synergy | Early Dinosaur that smooths draws.
- 1 Otepec Huntmaster | synergy | Reduces Dinosaur costs and can grant haste.
- 1 Raptor Companion | synergy | Straightforward early Dinosaur presence.
- 1 Carnage Tyrant | threat | Resilient top-end Dinosaur attacker.
- 1 Ghalta, Primal Hunger | threat | Massive Dinosaur that can be cast efficiently.
- 1 Regisaur Alpha | threat | Creates multiple Dinosaur bodies and grants haste.
- 1 Quartzwood Crasher | threat | Creates huge trampling Dinosaur tokens after combat damage.
- 1 Pantlaza, Sun-Favored | threat | Provides repeated value while deploying Dinosaurs.
- 1 Ghalta and Mavren | threat | Builds a wide board of powerful tokens.
- 1 Goring Ceratops | threat | Combat-focused Dinosaur that doubles strike damage.
- 1 Charging Tuskodon | threat | Turns blocked combat into major damage.
- 1 Colossal Dreadmaw | threat | Simple, iconic trampling Dinosaur finisher.
- 1 Etali, Primal Storm | threat | Generates explosive value when attacking.
- 1 Zetalpa, Primal Dawn | threat | Durable evasive Dinosaur finisher.
- 1 Ghalta, Stampede Tyrant | threat | Deploys a dramatic board of large creatures.
- 1 Blasphemous Act | wipe | Efficient reset button when the board gets crowded.
- 1 Chain Reaction | wipe | Creature-board reset that scales with the battlefield.
- 1 Vanquish the Horde | wipe | Often inexpensive creature-board reset.

</details>

### 5. blink Commander, owned first

Format: Commander. Theme: blink. Pool: owned_first. Shortlist: 168 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $127.09 the whole deck.

**Summary:** A mono-white blink deck centered on repeatedly harvesting creature entry abilities for cards, removal, and resilient board development. It combines protective equipment, artifact mana, and efficient answers with a creature-forward endgame that can rebuild well after opposing disruption.

The quality model grades this deck good against the top lists of the format: the commander does not place in cEDH events, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.87 over 63 nonland cards

<details><summary>The deck list</summary>

- 27 Plains | land | Reliable white mana base for a mono-white creature deck.
- 1 Command Tower | land | Untapped white source.
- 1 City of Brass | land | Flexible colored source.
- 1 Exotic Orchard | land | Flexible colored source.
- 1 Grand Coliseum | land | Additional white-producing land.
- 1 Minas Tirith | land | White source with late-game card advantage.
- 1 Plaza of Heroes | land | Supports the legendary commander while producing white mana.
- 1 Secluded Courtyard | land | Names Human for the deck's creature base.
- 1 Evolving Wilds | land | Fetches a Plains and improves land access.
- 1 Fabled Passage | land | Fetches a Plains and improves land access.
- 1 Sol Ring | ramp | Efficient artifact acceleration.
- 1 Arcane Signet | ramp | Reliable white mana acceleration.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Wayfarer's Bauble | ramp | Finds a basic Plains permanently.
- 1 Thought Vessel | ramp | Mana acceleration with hand-size utility.
- 1 Commander's Sphere | ramp | Fixes mana early and converts to a card later.
- 1 Sword of the Animist | ramp | Turns creature attacks into Plains ramp.
- 1 Relic of Legends | ramp | Produces mana from the commander and other legends.
- 1 Springleaf Drum | ramp | Efficient acceleration alongside cheap creatures.
- 1 Bender's Waterskin | ramp | Artifact-based mana development.
- 1 Wall of Omens | draw | Early card replacement that rewards repeated blinking.
- 1 Inspiring Overseer | draw | Creature-based card advantage with a useful blink trigger.
- 1 Lembas | draw | Reusable artifact cantrip through blinking and recursion.
- 1 Idol of Oblivion | draw | Steady card advantage with the deck's token and artifact support.
- 1 Mask of Memory | draw | Equipment-based card filtering on evasive attackers.
- 1 Tome of Legends | draw | Efficient long-game card advantage.
- 1 Skullclamp | draw | Converts smaller creatures into cards.
- 1 Puresteel Paladin | draw | Rewards the equipment package with cards.
- 1 Mirror of Galadriel | draw | Artifact-based card selection and advantage.
- 1 Crown of Gondor | draw | Equipment that supplies repeatable card advantage.
- 1 Stone of Erech | draw | Low-cost artifact card advantage.
- 1 Swords to Plowshares | removal | Premium answer to a creature.
- 1 Generous Gift | removal | Answers any problematic permanent.
- 1 Get Lost | removal | Efficient answer to creatures, enchantments, and planeswalkers.
- 1 Destroy Evil | removal | Flexible answer to large creatures or key enchantments.
- 1 Dispatch | removal | Efficient creature exile supported by artifacts.
- 1 Crib Swap | removal | Tribal-compatible instant-speed creature exile.
- 1 Fiend Hunter | removal | Blinkable creature removal.
- 1 Journey to Nowhere | removal | Low-cost creature exile.
- 1 Angel of Sanctions | removal | A flying body with a blinkable exile trigger.
- 1 Flickerwisp | synergy | Core blink effect that reuses enter-the-battlefield abilities.
- 1 Angel of Condemnation | synergy | Repeatable creature blinking and flexible board control.
- 1 Angel of Serenity | synergy | High-impact enter-the-battlefield effect to reuse.
- 1 Slip On the Ring | synergy | Protects a creature while retriggering its entry ability.
- 1 Gift of Immortality | synergy | Keeps key creatures available through removal.
- 1 Together Forever | synergy | Preserves important creatures for repeated deployment.
- 1 Personify | synergy | Supports the deck's creature-focused value plan.
- 1 Sheltered by Ghosts | synergy | Protects a key creature while advancing the board.
- 1 Duty Beyond Death | synergy | Protects creatures and supports recurring value bodies.
- 1 Crystal Fragments // Summon: Alexander | synergy | Provides a durable value permanent for the blink plan.
- 1 Bronze Guardian | threat | Large artifact threat that protects the deck's artifacts.
- 1 Grim Poppet | threat | Artifact creature that grows into a meaningful board presence.
- 1 PuPu UFO | threat | Artifact creature threat that contributes to the board.
- 1 The Walls of Ba Sing Se | threat | Resilient artifact creature that pressures opposing boards.
- 1 The Vision | threat | Legendary artifact creature that provides a substantial body.
- 1 Bastion Protector | threat | Protective creature that strengthens the commander.
- 1 Frontline Medic | threat | Creature threat with valuable combat protection.
- 1 Giada, Font of Hope | threat | Builds an increasingly powerful Angel board.
- 1 Kataki, War's Wage | threat | Disruptive creature that pressures artifact-heavy opponents.
- 1 Westfold Rider | threat | Creature body that contributes to combat pressure.
- 1 Champions of Minas Tirith | threat | Creature threat that rewards a developed board.
- 1 Faramir, Field Commander | threat | Legendary creature that builds a commanding board position.
- 1 Austere Command | wipe | Flexible reset that can preserve the preferred board.
- 1 Fumigate | wipe | Creature reset with life recovery.
- 1 Vanquish the Horde | wipe | Efficient creature board wipe.
- 1 Lightning Greaves | interaction | Protects key creatures and grants haste.
- 1 Swiftfoot Boots | interaction | Repeatable commander protection.
- 1 Darksteel Plate | interaction | Durable protection for an important creature.
- 1 Champion's Helm | interaction | Protects the commander and legendary threats.
- 1 Clever Concealment | interaction | Protects the board from a sweeping answer.
- 1 Reprieve | interaction | Temporarily answers a spell while replacing itself.
- 1 Unbreakable Formation | interaction | Protects the creature board and supports combat.
- 1 Ultimate Magic: Holy | interaction | Flexible instant-speed protection and disruption.

</details>

### 6. Modern tempo, tournament

Format: Modern. Theme: tempo. Pool: any_card. Shortlist: 167 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $205.06 to buy, $205.06 the whole deck.

**Summary:** This blue-red tempo build uses inexpensive draw, interaction, and removal to establish and protect a tempo advantage, while its creature-based draw cards provide pressure alongside that plan. It gives up a clearly supported dedicated finisher: the supplied shortlist contains no cards tagged as threats, so the deck must convert its tempo advantage without a listed primary win-condition package.

The quality model grades this deck good against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade. The deck runs its spells as playsets, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 2.39 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Island | land | Basic blue land for the mana base.
- 4 Mountain | land | Basic red land for the mana base.
- 4 Spirebluff Canal | land | Blue-red land for the tempo mana base.
- 4 Steam Vents | land | Blue-red land for the tempo mana base.
- 4 Shivan Reef | land | Blue-red land for the tempo mana base.
- 4 Riverglide Pathway // Lavaglide Pathway | land | Flexible blue-red land slot.
- 4 Consider | draw | Low-cost draw support for a spell-heavy plan.
- 2 Preordain | draw | Additional early draw and selection.
- 4 Ledger Shredder | draw | Creature-based draw support.
- 4 Faerie Mastermind | draw | Creature-based draw support.
- 4 Enduring Curiosity | draw | Additional draw support that stays on the table.
- 4 Counterspell | interaction | Core broad interaction for protecting a tempo advantage.
- 2 Spell Pierce | interaction | Efficient interaction for early exchanges.
- 4 Lightning Bolt | removal | Efficient removal for opposing threats.
- 2 Abrade | removal | Flexible removal coverage.
- 2 Into the Flood Maw | removal | Low-cost removal option.
- 4 Temporal Mastery | synergy | Supplies the available synergy package.

</details>

### 7. Modern burn, casual

Format: Modern. Theme: burn. Pool: any_card. Shortlist: 290 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $33.50 to buy, $33.50 the whole deck.

**Summary:** This mono-red burn deck applies pressure with a steady threat package, then uses direct removal and damage-focused synergy cards to keep the opponent under pressure. Its draw spells help it continue deploying action, while a small ramp package supports its broader game plan. It gives up flexibility against unusual permanent types in exchange for a focused, straightforward damage strategy.

The quality model grades this deck below the precon baseline against the top lists of the format: the deck runs its spells as playsets, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 2.83 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Supplies the mono-red mana base.
- 2 Chandra, Dressed to Kill | ramp | Fills the ramp slots for the deck.
- 4 Risk Factor | draw | Provides burn-oriented card draw.
- 2 Browbeat | draw | Adds further card draw for the burn plan.
- 4 Lightning Bolt | removal | Primary efficient burn removal.
- 2 Burst Lightning | removal | Adds flexible burn removal.
- 4 Eidolon of the Great Revel | synergy | A synergy piece for a low-cost burn shell.
- 4 Roiling Vortex | synergy | Supports the deck's damage-focused synergy plan.
- 4 Akki Lavarunner // Tok-Tok, Volcano Born | threat | Provides an aggressive threat.
- 4 Fuming Effigy | threat | Adds pressure as a threat.
- 4 Keral Keep Disciples | threat | Supplies another set of threats.
- 2 Hazoret the Fervent | threat | Rounds out the threat package.

</details>

### 8. Modern lifegain, FNM

Format: Modern. Theme: lifegain. Pool: any_card. Shortlist: 301 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $463.44 to buy, $463.44 the whole deck.

**Summary:** This white-black lifegain deck builds an early board around Soul Warden and Ajani's Pridemate, then converts that foundation into sustained creature pressure with its Vampire threats. Card-draw creatures help it keep developing, while Solitude and Murderous Rider protect the board from opposing threats. It wins primarily through combat after its creatures take over the table, giving up some speed and broad utility for a focused, creature-based plan.

The quality model grades this deck below the precon baseline against the top lists of the format: the cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The deck runs few spells as one copy, and that raises the grade.

- [INFO] `curve_summary`: average mana value 3.28 over 36 nonland cards

<details><summary>The deck list</summary>

- 10 Plains | land | Provides reliable white mana for the deck's white spells.
- 10 Swamp | land | Provides reliable black mana for the deck's black spells.
- 4 Godless Shrine | land | Supplies access to both deck colors without adding a tapped land.
- 2 Arcane Signet | ramp | Helps accelerate the deck's mana development.
- 4 Preacher of the Schism | draw | Provides card advantage while contributing a creature body.
- 2 Inspiring Overseer | draw | Adds card draw on a creature that supports the white side of the deck.
- 4 Solitude | removal | Provides efficient answers to opposing threats.
- 2 Murderous Rider // Swift End | removal | Serves as removal while remaining a creature afterward.
- 4 Soul Warden | synergy | Establishes the deck's early lifegain synergy.
- 4 Ajani's Pridemate | synergy | Turns the lifegain plan into growing creature pressure.
- 4 Attended Healer | threat | Provides a resilient midgame creature threat for the lifegain plan.
- 4 Bloodbond Vampire | threat | Offers another creature threat that fits the deck's central plan.
- 2 Enduring Tenacity | threat | Adds a black creature threat at the middle of the curve.
- 2 Sheoldred, the Apocalypse | threat | Supplies a high-impact top-end threat.
- 2 Bloodthirsty Conqueror | threat | Provides powerful finishing pressure for longer games.

</details>

### 9. Standard midrange, FNM

Format: Standard. Theme: midrange. Pool: any_card. Shortlist: 164 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $250.82 to buy, $250.82 the whole deck.

**Summary:** This black-green midrange deck develops with Llanowar Elves, builds card advantage through Phyrexian Arena and Unholy Annex // Ritual Chamber, and clears resistance with a focused removal package. It wins by establishing creature pressure from Darkstar Augur and Goldvein Hydra before closing with Vein Ripper or Vaultborn Tyrant, while Undying Malice and Snakeskin Veil support the board-centered plan. It gives up some speed for a sturdier long-game approach and depends on keeping its creature threats in play.

The quality model grades this deck good against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.92 over 36 nonland cards

<details><summary>The deck list</summary>

- 4 Overgrown Tomb | land | Primary black-green source for casting the deck's core spells.
- 4 Blooming Marsh | land | Provides efficient early black-green mana.
- 4 Underground Mortuary | land | Adds another black-green source to support a consistent mana base.
- 7 Swamp | land | Reliable black mana for removal, draw, and threats.
- 5 Forest | land | Reliable green mana for early development and threats.
- 4 Phyrexian Arena | draw | A dedicated source of ongoing card advantage.
- 2 Unholy Annex // Ritual Chamber | draw | Adds card advantage while remaining within the black core.
- 2 Llanowar Elves | ramp | Accelerates the deck into its stronger midgame plays.
- 4 Bitter Triumph | removal | Efficient black removal for opposing problems.
- 2 Assassin's Trophy | removal | Flexible black-green removal for difficult permanents.
- 4 Undying Malice | synergy | Supports the creature-heavy plan and helps maintain pressure.
- 4 Snakeskin Veil | synergy | Supports key creatures and the deck's board-focused plan.
- 4 Darkstar Augur | threat | An early creature threat that starts the deck's pressure.
- 4 Goldvein Hydra | threat | A scalable green creature threat for the midgame.
- 3 Vein Ripper | threat | A powerful black finisher for closing longer games.
- 3 Vaultborn Tyrant | threat | A large top-end creature threat for decisive late-game pressure.

</details>

### 10. Standard aggro, tournament

Format: Standard. Theme: aggro. Pool: any_card. Shortlist: 294 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $57.56 to buy, $57.56 the whole deck.

**Summary:** This red-white aggro deck establishes pressure with a dense creature suite, then uses Warleader's Call to reinforce its combat plan. Removal and interaction keep opposing defenses from slowing the attack, while its draw cards help maintain momentum after the first wave. It wins by continuing to present threats and push combat, giving up the broader answers and late-game power of a slower strategy.

The quality model grades this deck below the precon baseline against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The mana base serves the colors evenly, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.00 over 36 nonland cards

<details><summary>The deck list</summary>

- 12 Mountain | land | Basic red mana for the aggressive red-white base.
- 12 Plains | land | Basic white mana for the aggressive red-white base.
- 2 Fugitive Codebreaker | draw | Early red creature that fills the deck's draw slot.
- 2 Reckless Lackey | draw | Low-cost red creature that contributes card access.
- 2 Irreverent Gremlin | draw | An aggressive red creature assigned to card draw.
- 4 Boros Charm | interaction | Efficient red-white interaction for protecting the attack or disrupting opponents.
- 2 Aven Interrupter | interaction | A creature-based interaction piece that keeps pressure on the board.
- 4 Harsh Annotation | removal | Cheap red removal to clear opposing resistance.
- 4 Case of the Gateway Express | removal | Additional removal that supports the creature-forward plan.
- 4 Warleader's Call | synergy | A red-white synergy card that reinforces the deck's attacking plan.
- 4 Redcap Gutter-Dweller | threat | A red threat that provides a concentrated aggressive creature package.
- 4 Bedhead Beastie | threat | A creature threat for maintaining pressure through combat.
- 4 Frilled Sparkshooter | threat | A red creature threat that rounds out the low-to-the-ground attack force.

</details>

### 11. Commander with a locked card

Format: Commander. Theme: sacrifice. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $706.45 to buy, $706.45 the whole deck.

**Summary:** Karlov leads a black-white sacrifice deck that develops expendable creatures, sacrifice outlets, and value pieces before converting that board into steady pressure. It closes through its creature threats and sacrifice payoffs while using removal and sweepers to keep opposing boards manageable. The deck gives up some speed to maintain a board-centric engine, so it is strongest when it can establish several complementary permanents rather than relying on a single threat.

The quality model grades this deck below the precon baseline against the top lists of the format: the mana base serves one color far better than another, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. Many of the cards appear in no top list, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 3.24 over 63 nonland cards

<details><summary>The deck list</summary>

- 1 Corrupted Conviction | draw | Low-cost draw that fits the sacrifice plan.
- 1 Disciple of Bolas | draw | A sacrifice-oriented draw piece for turning a developed board into cards.
- 1 Vampiric Rites | draw | Repeatable draw support within the sacrifice shell.
- 1 Village Rites | draw | Efficient draw that uses the deck's sacrifice theme.
- 1 Smothering Abomination | draw | A draw threat that rewards the deck's core plan.
- 1 Tevesh Szat, Doom of Fools | draw | A durable draw engine for longer games.
- 1 Relic Vial | draw | A draw artifact that supports the creature-heavy build.
- 1 Shadowheart, Dark Justiciar | draw | A draw option that works with expendable creatures.
- 1 Baron Bertram Graywater | draw | A draw creature that contributes to the sacrifice strategy.
- 1 Jerren, Corrupted Bishop // Ormendahl, the Corrupter | draw | A draw-oriented threat for the creature plan.
- 1 Thraxodemon | draw | Additional draw attached to a sacrifice-focused creature.
- 1 Cartel Aristocrat | interaction | A flexible interaction creature for protecting the board plan.
- 1 Dark Privilege | interaction | An interaction piece that supports sacrificing creatures.
- 1 Fanatical Devotion | interaction | Board-protection interaction that fits the deck's theme.
- 1 Flare of Fortitude | interaction | A protection spell for preserving a developed board.
- 1 Gift of Doom | interaction | Protective interaction for an important permanent.
- 1 Kinzu of the Bleak Coven | interaction | An interaction creature that complements the sacrifice shell.
- 1 Nightmare Shepherd | interaction | Creature-based interaction for the deck's board-centric plan.
- 1 Rescue from the Underworld | interaction | Flexible interaction that supports the sacrifice strategy.
- 9 Plains | land | Basic white mana for the two-color mana base.
- 20 Swamp | land | Basic black mana for the sacrifice-heavy deck.
- 1 Caves of Koilos | land | Dual-color land for consistent mana.
- 1 Command Tower | land | Reliable color fixing for the commander deck.
- 1 Godless Shrine | land | Dual-color land for the mana base.
- 1 Scrubland | land | Dual-color land for dependable early mana.
- 1 Vault of Champions | land | Dual-color land that supports the color requirements.
- 1 Shattered Sanctum | land | Dual-color land for the black-white core.
- 1 Isolated Chapel | land | Additional dual-color support for the mana base.
- 1 Sol Ring | ramp | Required card that fills a ramp slot.
- 1 Ashnod's Altar | ramp | Sacrifice-based mana acceleration for explosive turns.
- 1 Phyrexian Altar | ramp | A central sacrifice outlet that also accelerates mana.
- 1 Pitiless Plunderer | ramp | Ramp creature that rewards the sacrifice plan.
- 1 Pawn of Ulamog | ramp | Creature-based ramp that supports expendable bodies.
- 1 Sifter of Skulls | ramp | Ramp payoff for a creature-sacrifice strategy.
- 1 Priest of Forgotten Gods | ramp | A sacrifice-focused ramp creature.
- 1 Culling the Weak | ramp | Fast sacrifice-based mana burst.
- 1 Deadly Dispute | ramp | Sacrifice-compatible resource acceleration.
- 1 Skullport Merchant | ramp | A ramp creature that fits the value-oriented plan.
- 1 Attrition | removal | Sacrifice-based removal for problematic permanents.
- 1 Ayli, Eternal Pilgrim | removal | Creature-based removal that belongs in the sacrifice shell.
- 1 Blasting Station | removal | A sacrifice outlet that doubles as removal.
- 1 Bone Shards | removal | Efficient removal that uses expendable creatures.
- 1 Dictate of Erebos | removal | A powerful removal payoff for sacrificing creatures.
- 1 Grave Pact | removal | A major removal engine tied to the deck's main theme.
- 1 Teysa, Orzhov Scion | removal | A sacrifice-focused removal creature.
- 1 Yawgmoth, Thran Physician | removal | Versatile sacrifice-based removal and utility.
- 1 Erebos, Bleak-Hearted | removal | Removal support that complements creature sacrifices.
- 1 Altar of Dementia | synergy | A sacrifice outlet and dedicated synergy piece.
- 1 Bartolomé del Presidio | synergy | A low-cost creature that supports sacrificing permanents.
- 1 Bastion of Remembrance | synergy | A core sacrifice payoff for the deck's theme.
- 1 Carrion Feeder | synergy | An efficient sacrifice outlet.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | A sacrifice payoff creature for the black-white shell.
- 1 Fleshtaker | synergy | A synergy creature that rewards sacrificing creatures.
- 1 Hidden Stockpile | synergy | A sacrifice-oriented value engine.
- 1 Viscera Seer | synergy | A cheap sacrifice outlet for setting up stronger turns.
- 1 Woe Strider | synergy | A sacrifice outlet that reinforces the creature plan.
- 1 Zulaport Cutthroat | synergy | A key payoff for sacrificing creatures.
- 1 Felisa, Fang of Silverquill | threat | A sacrifice-compatible threat for the commander strategy.
- 1 Mondrak, Glory Dominus | threat | A high-impact creature threat for building a board.
- 1 Ratadrabik of Urborg | threat | A resilient legendary threat in the creature-heavy deck.
- 1 Liesa, Forgotten Archangel | threat | A powerful top-end threat for the black-white shell.
- 1 Ghoulcaller Gisa | threat | A creature-based threat that supports the sacrifice plan.
- 1 Liliana, Heretical Healer // Liliana, Defiant Necromancer | threat | A sacrifice-compatible threat with lasting board presence.
- 1 Vindictive Vampire | threat | A creature threat that fits the sacrifice theme.
- 1 Sidisi, Undead Vizier | threat | A substantial black threat for the creature strategy.
- 1 Requiem Angel | threat | An evasive top-end threat for the board-focused plan.
- 1 Threefold Thunderhulk | threat | A large artifact threat that adds battlefield pressure.
- 1 Kuldotha Forgemaster | threat | An artifact creature threat for the deck's upper curve.
- 1 Thallid Omnivore | threat | A sacrifice-oriented threat that benefits from expendable creatures.
- 1 Toxic Deluge | wipe | A flexible reset button when the board gets away from the deck.
- 1 The Meathook Massacre | wipe | A sweeping board-control option that matches the creature plan.
- 1 Austere Command | wipe | A versatile wipe for handling difficult boards.

</details>

### 12. Commander on a budget

Format: Commander. Theme: tokens. Pool: any_card. Shortlist: 296 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $59.89 to buy, $59.89 the whole deck.

**Summary:** Adeline leads a deliberate mono-white token strategy that builds a wide creature board, reinforces it with token-focused permanents, and turns repeated combat pressure into the primary path to victory. The deck has steady mana development, substantial card flow, and several ways to answer troublesome permanents or reset an opposing board. It gives up explosive speed and premium individual-card power in exchange for a cohesive, resilient battlefield plan that rewards committing creatures and tokens over multiple turns.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- [INFO] `curve_summary`: average mana value 3.27 over 62 nonland cards

<details><summary>The deck list</summary>

- 30 Plains | land | Primary basic land for the mono-white mana base.
- 1 Ancient Den | land | White-producing land slot for the mana base.
- 1 Castle Ardenvale | land | White land slot for the mana base.
- 1 Command Tower | land | Reliable color-producing land slot.
- 1 Eiganjo, Seat of the Empire | land | White land slot for the mana base.
- 1 Exotic Orchard | land | Color-producing land slot for the mana base.
- 1 Secluded Courtyard | land | Creature-focused color-producing land slot.
- 1 Windbrisk Heights | land | White land slot for the mana base.
- 1 Bucknard's Everfull Purse | ramp | Ramp piece from the shortlist.
- 1 Charisma Bobblehead | ramp | Ramp piece from the shortlist.
- 1 Coin of Mastery | ramp | Ramp piece from the shortlist.
- 1 Collector's Vault | ramp | Ramp piece from the shortlist.
- 1 Currency Converter | ramp | Low-cost ramp piece from the shortlist.
- 1 Discerning Financier | ramp | Creature-based ramp for the deck's board plan.
- 1 Druidic Satchel | ramp | Ramp piece from the shortlist.
- 1 Goldvein Pick | ramp | Equipment-based ramp piece.
- 1 Idol of False Gods | ramp | Ramp piece from the shortlist.
- 1 Karn, Living Legacy | ramp | Ramp support that also adds a durable permanent.
- 1 Angelic Sell-Sword | draw | Draw option from the shortlist.
- 1 Bygone Bishop | draw | Creature-based draw for the token plan.
- 1 Dawn of Hope | draw | Draw support for a long board-focused game.
- 1 Faramir, Field Commander | draw | Draw creature that fits the deck's creature plan.
- 1 Glimmer Seeker | draw | Low-cost draw option.
- 1 Idol of Oblivion | draw | Token-oriented draw support.
- 1 Platoon Dispenser | draw | Draw permanent for the board-focused strategy.
- 1 Sarah Jane Smith | draw | Creature-based draw support.
- 1 Search the Premises | draw | Draw engine for a longer game.
- 1 Staff of the Storyteller | draw | Affordable draw support.
- 1 Wedding Announcement // Wedding Festivity | draw | Draw piece that fits the token theme.
- 1 Ainok Strike Leader | interaction | Interaction creature from the shortlist.
- 1 Basri Ket | interaction | Interaction planeswalker that supports a board plan.
- 1 Basri, Tomorrow's Champion | interaction | Interaction creature for the deck's creature-heavy plan.
- 1 Lena, Selfless Champion | interaction | Interaction creature that fits the token strategy.
- 1 Rootborn Defenses | interaction | Low-cost interaction for protecting the board plan.
- 1 Spirit Bonds | interaction | Interaction enchantment that fits a token-focused shell.
- 1 Generous Gift | removal | Flexible removal from the shortlist.
- 1 Hanged Executioner | removal | Creature-based removal that fits the board plan.
- 1 Kellan's Lightblades | removal | Low-cost removal option.
- 1 Righteous Confluence | removal | Flexible removal option.
- 1 Shire Shirriff | removal | Affordable creature-based removal.
- 1 Skyclave Apparition | removal | Efficient creature-based removal.
- 1 Stroke of Midnight | removal | Flexible instant-speed removal.
- 1 Hour of Reckoning | wipe | Board reset selected for a token strategy.
- 1 Martial Coup | wipe | Board reset that remains aligned with the token plan.
- 1 Phyrexian Rebirth | wipe | Additional board reset for established opposing boards.
- 1 Aligned Heart | synergy | Token-plan synergy piece from the shortlist.
- 1 Anafenza, Unyielding Lineage | synergy | Creature-based synergy for the token strategy.
- 1 Animation Module | synergy | Artifact synergy piece for building a board.
- 1 Automated Assembly Line | synergy | Affordable token-plan synergy permanent.
- 1 Cat Collector | synergy | Creature synergy for the token theme.
- 1 Clarion Spirit | synergy | Low-cost creature synergy for the board plan.
- 1 Divine Visitation | synergy | Token-focused synergy enchantment.
- 1 Felidar Retreat | synergy | Token-plan synergy permanent.
- 1 Horn of Gondor | synergy | Token-focused artifact synergy.
- 1 Intangible Virtue | synergy | Efficient synergy for a wide board.
- 1 Rosie Cotton of South Lane | synergy | Creature synergy for the token strategy.
- 1 Siege Veteran | synergy | Creature synergy that supports a wide board.
- 1 Skrelv's Hive | synergy | Token-focused enchantment synergy.
- 1 Ajani's Chosen | threat | Token-oriented threat for the creature suite.
- 1 Archangel Elspeth | threat | High-impact token-plan threat.
- 1 Archon of Sun's Grace | threat | Creature threat that fits the enchantment-heavy support package.
- 1 Attended Healer | threat | Affordable creature threat for the board plan.
- 1 Basri's Lieutenant | threat | Creature threat that complements a wide board.
- 1 Cemetery Protector | threat | Standalone creature threat.
- 1 Defiler of Faith | threat | Creature threat for the mono-white shell.
- 1 Emeria Angel | threat | Token-focused creature threat.
- 1 God-Eternal Oketra | threat | Resilient creature threat for the deck's board plan.
- 1 Hero of Bladehold | threat | Combat-focused creature threat.
- 1 Threefold Thunderhulk | threat | Artifact creature threat that supports going wide.
- 1 Warren Warleader | threat | Creature threat that fits the token strategy.

</details>

### 13. owned first, and the commander is not owned

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $0.00 to buy, $152.40 the whole deck.

**Summary:** This is a black-white lifegain deck that establishes Karlov alongside a creature-heavy board, then uses its lifegain-focused synergies and Angels to apply steady combat pressure. It wins by protecting key creatures, clearing away resistance with focused removal or board wipes, and maintaining enough card flow to rebuild. The deck gives up multicolor flexibility for a consistent, board-centered plan with strong access to protection and answers.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade.

Summary rules claims (F-26): the format play

- [WARN] `not_owned`: Karlov of the Ghost Council: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.98 over 63 nonland cards

<details><summary>The deck list</summary>

- 18 Plains | land | White basic land for a stable mana base.
- 18 Swamp | land | Black basic land for a stable mana base.
- 1 Buster Sword | draw | Equipment-based draw option that supports the deck's creature plan.
- 1 Call of the Ring | draw | Dedicated card-advantage piece.
- 1 Exemplar of Light | draw | Creature-based draw that fits the white-focused shell.
- 1 Idol of Oblivion | draw | Low-cost artifact draw source.
- 1 Inspiring Overseer | draw | Creature-based card advantage for the board-focused plan.
- 1 Lembas | draw | Flexible artifact draw option.
- 1 Night's Whisper | draw | Efficient black card-advantage spell.
- 1 Puresteel Paladin | draw | Draw creature that also fits the equipment package.
- 1 Skullclamp | draw | Efficient equipment-based card-advantage tool.
- 1 Tome of Legends | draw | Repeatable artifact draw source.
- 1 Wall of Omens | draw | Early defensive body with card advantage.
- 1 Bastion Protector | interaction | Protective creature that helps preserve the deck's centerpiece.
- 1 Boromir, Warden of the Tower | interaction | Creature-based disruptive interaction.
- 1 Champion's Helm | interaction | Equipment protection for an important creature.
- 1 Clever Concealment | interaction | Broad protection against opposing interaction.
- 1 Darksteel Plate | interaction | Durable equipment protection for a key threat.
- 1 Lightning Greaves | interaction | Efficient protection equipment.
- 1 Swiftfoot Boots | interaction | Additional protection equipment for key creatures.
- 1 Unbreakable Formation | interaction | Board-protection spell for committed creature positions.
- 1 Arcane Signet | ramp | Efficient color-fixing mana rock.
- 1 Commander's Sphere | ramp | Reliable fixing that remains useful later.
- 1 Fellwar Stone | ramp | Efficient two-mana mana rock.
- 1 Giada, Font of Hope | ramp | Creature-based acceleration that fits the Angel portion of the deck.
- 1 Lotho, Corrupt Shirriff | ramp | Theme-compatible creature ramp.
- 1 Relic of Legends | ramp | Mana rock that supports a legendary commander shell.
- 1 Sol Ring | ramp | Efficient early mana acceleration.
- 1 Sword of the Animist | ramp | Equipment-based mana development for attacking creatures.
- 1 Thought Vessel | ramp | Colorless mana acceleration.
- 1 Wayfarer's Bauble | ramp | Basic-land ramp for steady development.
- 1 Banishing Light | removal | Flexible permanent removal.
- 1 Bitter Triumph | removal | Efficient black spot removal.
- 1 Crib Swap | removal | Creature removal with useful flexibility.
- 1 Dispatch | removal | Low-cost targeted removal.
- 1 Fatal Push | removal | Efficient early targeted removal.
- 1 Generous Gift | removal | Broad answer to troublesome permanents.
- 1 Get Lost | removal | Versatile white permanent removal.
- 1 Infernal Grasp | removal | Reliable black creature removal.
- 1 Swords to Plowshares | removal | Premium efficient creature answer.
- 1 Aerith Gainsborough | synergy | Lifegain-focused synergy creature.
- 1 Angel of Vitality | synergy | Angel synergy that supports the lifegain plan.
- 1 Compassionate Healer | synergy | Dedicated lifegain synergy piece.
- 1 Kor Firewalker | synergy | Creature synergy for the deck's lifegain focus.
- 1 Light of Promise | synergy | Payoff aura for the deck's central theme.
- 1 Night Nurse, Healer of Heroes | synergy | Theme-focused legendary synergy creature.
- 1 Prideful Feastling | synergy | Creature synergy for the lifegain shell.
- 1 Rosie Cotton of South Lane | synergy | Creature payoff that supports the board plan.
- 1 Second Breakfast | synergy | Theme-supporting instant synergy piece.
- 1 Wanderbrine Preacher | synergy | Creature synergy that complements the deck's core plan.
- 1 Angel of Invention | threat | Evasive creature threat for closing games.
- 1 Bill the Pony | threat | Efficient creature threat that fits the white creature base.
- 1 Dawnhand Eulogist | threat | Creature threat that adds battlefield pressure.
- 1 Foggy Swamp Hunters | threat | Creature threat for sustaining pressure.
- 1 Gollum, Silent Slinker // Meager Meal | threat | Black creature threat with an attached adventure.
- 1 Lyra Dawnbringer | threat | High-impact Angel threat.
- 1 Minwu, White Mage | threat | Legendary creature threat for the white core.
- 1 Rabaroo Troop | threat | Creature threat that helps fill out the attacking board.
- 1 Rooftop Percher | threat | Creature threat for additional board presence.
- 1 Shattered Angel | threat | Angel threat that fits the lifegain theme.
- 1 Sneering Shadewriter | threat | Black creature threat for pressure and curve depth.
- 1 Victory's Herald | threat | Top-end Angel threat for decisive attacks.
- 1 Austere Command | wipe | Flexible reset option for unfavorable boards.
- 1 Fumigate | wipe | Board wipe that matches the deck's lifegain focus.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset.

</details>

### 14. the user delegates the commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 244 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $614.09 to buy, $614.09 the whole deck.

**Summary:** A Gruul Dragon deck that ramps into a steady parade of evasive threats, using tribal cost reduction, token production, and combat-focused payoffs to turn the skies into a decisive advantage. Flexible creature control and protection preserve the board, while repeated Dragon entries provide damage, cards, mana, and overwhelming finishing attacks.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade. Few decks lead with the commander, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.29 over 63 nonland cards

<details><summary>The deck list</summary>

- 11 Mountain | land | Basic red sources support the deck's dragon-heavy red requirements.
- 9 Forest | land | Basic green sources support early ramp and green utility spells.
- 1 Cavern of Souls | land | Helps key Dragon creatures resolve through counterspells.
- 1 Cinder Glade | land | Dual land that provides both commander colors.
- 1 City of Brass | land | Flexible fixing for either color.
- 1 Command Tower | land | Reliable two-color commander mana.
- 1 Commercial District | land | Provides dependable red-green fixing.
- 1 Cragcrown Pathway // Timbercrown Pathway | land | Flexible early colored source.
- 1 Exotic Orchard | land | Usually supplies either needed color in multiplayer.
- 1 Game Trail | land | Red-green dual source.
- 1 Haven of the Spirit Dragon | land | Dragon-focused mana source with late-game recursion utility.
- 1 Karplusan Forest | land | Untapped red-green fixing.
- 1 Mana Confluence | land | Untapped fixing for both colors.
- 1 Rockfall Vale | land | Red-green dual source.
- 1 Rootbound Crag | land | Red-green dual source.
- 1 Spire Garden | land | Usually enters ready to produce either color.
- 1 Stomping Ground | land | Fetchable red-green dual source.
- 1 Taiga | land | Untapped fetchable red-green dual source.
- 1 Faithless Looting | draw | Efficient card selection and graveyard setup.
- 1 Demand Answers | draw | Turns spare artifacts or cards into fresh resources.
- 1 Thrill of Possibility | draw | Instant-speed filtering helps find lands or threats.
- 1 Idol of Oblivion | draw | Repeatable card draw alongside Dragon token makers.
- 1 Skullclamp | draw | Converts small creature tokens into cards.
- 1 Elemental Bond | draw | Most Dragons entering the battlefield trigger steady card draw.
- 1 Garruk's Uprising | draw | Draws from large creatures and grants trample.
- 1 Beast Whisperer | draw | Creature-heavy sequencing turns Dragons into cards.
- 1 Dragonborn Champion | draw | Combat damage from large fliers keeps cards flowing.
- 1 Harmonize | draw | Straightforward refill after committing to the board.
- 1 Return of the Wildspeaker | draw | Can refill from a large Dragon or push lethal combat damage.
- 1 Lightning Greaves | interaction | Protects a major threat while granting haste.
- 1 Swiftfoot Boots | interaction | Provides repeatable haste and targeted protection.
- 1 Snakeskin Veil | interaction | Cheap protection against removal and sweepers that deal damage.
- 1 Tamiyo's Safekeeping | interaction | Protects a key permanent while buying a little life.
- 1 Heroic Intervention | interaction | Protects the developed Dragon board from removal.
- 1 Veil of Summer | interaction | Efficient defense against blue or black disruption.
- 1 Tibalt's Trickery | interaction | Flexible answer to a pivotal spell.
- 1 Pyroblast | interaction | Low-cost answer to dangerous blue spells or permanents.
- 1 Carnelian Orb of Dragonkind | ramp | Early fixing that gives Dragons haste when spent on them.
- 1 Dragon's Hoard | ramp | Dragon casting produces mana and later converts to cards.
- 1 Dragonstorm Globe | ramp | Provides fixing and rewards casting Dragons.
- 1 Jade Orb of Dragonkind | ramp | Early mana that improves Dragon bodies.
- 1 Mox Jasper | ramp | Low-cost Dragon-dependent color fixing.
- 1 Orb of Dragonkind | ramp | Fixes mana and can find a Dragon later.
- 1 Scaled Nurturer | ramp | Early creature ramp tailored to Dragons.
- 1 Ganax, Astral Hunter | ramp | Dragon entries create Treasure for explosive follow-up turns.
- 1 Gadrak, the Crown-Scourge | ramp | Creates Treasure while providing an evasive body.
- 1 Sarkhan, Fireblood | ramp | Accelerates Dragons and filters excess cards.
- 1 Draconic Roar | removal | Efficient creature removal with bonus damage from a Dragon.
- 1 Dragon's Fire | removal | Scales into strong removal with a Dragon in hand or play.
- 1 Molten Exhale | removal | Efficient Dragon-scaled creature removal.
- 1 Piercing Exhale | removal | Cheap interaction that rewards controlling Dragons.
- 1 Spit Flame | removal | Reusable removal whenever Dragons enter the battlefield.
- 1 Scourge of Valkas | removal | Turns every Dragon arrival into direct damage.
- 1 Glorybringer | removal | Hasty evasive threat that removes problematic creatures.
- 1 Invasion of Tarkir // Defiant Thundermaw | removal | Dragon-revealing removal that becomes a threatening flier.
- 1 Dragon Tempest | removal | Dragon entries can immediately pick off creatures or players.
- 1 Crucible of Fire | synergy | Makes every Dragon substantially more threatening in combat.
- 1 Dragonkin Berserker | synergy | Uses spare mana to develop Dragon tokens and stronger attacks.
- 1 Firespitter Whelp | synergy | Early Dragon that rewards continued Dragon casting.
- 1 Kargan Dragonrider | synergy | Gets stronger and evasive alongside Dragons.
- 1 Minion of the Mighty | synergy | Can deploy a large Dragon ahead of schedule after attacking.
- 1 Nogi, Draco-Zealot | synergy | Supports the tribal plan while leveraging Dragon attacks.
- 1 Slumbering Dragon | synergy | Early Dragon body that becomes a serious deterrent.
- 1 Dragonspeaker Shaman | synergy | Reduces the cost of the deck's central creature type.
- 1 Sarkhan's Triumph | synergy | Finds the most suitable Dragon for the current board state.
- 1 The Dragon-Kami Reborn // Dragon-Kami's Egg | synergy | Builds toward a powerful Dragon payoff with card selection.
- 1 Ancient Bronze Dragon | threat | Combat trigger can generate an overwhelming board of permanents.
- 1 Backdraft Hellkite | threat | Evasive attacker that reuses impactful graveyard spells.
- 1 Blast-Furnace Hellkite | threat | A punishing combat threat that amplifies attacking Dragons.
- 1 Caldera Pyremaw | threat | Large flier that pressures life totals and boards.
- 1 Canopy Gargantuan | threat | Massive Dragon finisher for aerial combat.
- 1 Dragon Broodmother | threat | Builds a growing Dragon token army over successive turns.
- 1 Hellkite Charger | threat | Threatens extra combat steps with available mana.
- 1 Lathliss, Dragon Queen | threat | Every nontoken Dragon expands the aerial army.
- 1 Scourge of the Throne | threat | Extra combats turn a successful attack into a finishing sequence.
- 1 Thrakkus the Butcher | threat | Doubles Dragon power for explosive combat turns.
- 1 Thunderbreak Regent | threat | Efficient flying pressure that punishes targeted Dragon removal.
- 1 Utvara Hellkite | threat | Repeated attacks rapidly create a game-ending Dragon swarm.
- 1 Breath Weapon | wipe | Low-cost sweeper that largely spares the deck's Dragons.
- 1 Draconic Intervention | wipe | Flexible scalable sweeper that leverages a large Dragon card.
- 1 Ryusei, the Falling Star | wipe | Dragon threat that clears smaller opposing creatures on death.

</details>

### 15. delegated commander, owned first

Format: Commander. Theme: lifegain. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, profile_off_band. Block findings: 0.

Cost: $0.00 to buy, $169.57 the whole deck.

**Summary:** A black-white lifegain deck built to turn steady life-total growth into resilient creature pressure. Efficient mana and card advantage establish a board of lifegain payoffs, while protective equipment and spells preserve key threats. Flexible spot removal and sweeping answers keep opposing boards contained until evasive Angels and enhanced creatures can finish the game.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [WARN] `not_owned`: Astarion, the Decadent: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 3.21 over 62 nonland cards
- [INFO] `basics_added`: the list was 2 cards short, so the builder added 2 basic lands

<details><summary>The deck list</summary>

- 12 Plains | land | Reliable white source for the deck’s lifegain creatures and protective spells.
- 12 Swamp | land | Reliable black source for the commander and efficient removal.
- 1 Command Tower | land | Reliable dual-color commander land.
- 1 Exotic Orchard | land | Flexible dual-color fixing in multiplayer games.
- 1 City of Brass | land | Unrestricted color fixing.
- 1 Grand Coliseum | land | Provides either deck color when needed.
- 1 Castle Locthwain | land | Black source with late-game card advantage.
- 1 Minas Tirith | land | White source with a useful legendary-land payoff.
- 1 Takenuma, Abandoned Mire | land | Black source that can recover an important creature.
- 1 Plaza of Heroes | land | Supports the legendary commander and other legends.
- 1 Opal Palace | land | Fixes mana while improving commander recasts.
- 1 Spire of Industry | land | Flexible fixing supported by the artifact package.
- 1 Secluded Courtyard | land | Fixes mana for the deck’s substantial creature suite.
- 1 Unclaimed Territory | land | Additional creature-focused color fixing.
- 1 Path of Ancestry | land | Fixes colors and rewards casting creatures that share a type with the commander.
- 1 Buster Sword | draw | Equipment-based card advantage that rewards connecting in combat.
- 1 Call of the Ring | draw | Steady card flow with life-loss that lifegain can offset.
- 1 Circle of Power | draw | Card advantage that helps maintain resources through the midgame.
- 1 Exemplar of Light | draw | A lifegain-friendly creature that converts the theme into cards.
- 1 Grave Venerations | draw | Ongoing card advantage for a creature-focused strategy.
- 1 Idol of Oblivion | draw | Efficient repeatable draw alongside token-making effects.
- 1 Inspiring Overseer | draw | Flying lifegain body that replaces itself.
- 1 Lembas | draw | Early card selection with a Food-themed lifegain resource.
- 1 Mask of Memory | draw | Combat-based draw that fills the graveyard selectively.
- 1 Night's Whisper | draw | Low-cost unconditional card draw.
- 1 Painful Truths | draw | Efficient card refill whose life payment fits the theme.
- 1 Bastion Protector | interaction | Keeps the commander safer while contributing a relevant body.
- 1 Clever Concealment | interaction | Protects the board from opposing sweepers or targeted disruption.
- 1 Lightning Greaves | interaction | Gives the commander immediate protection and haste.
- 1 Reprieve | interaction | Temporarily answers a problematic spell while replacing itself.
- 1 Swiftfoot Boots | interaction | Reliable commander protection that preserves targeting flexibility.
- 1 Unbreakable Formation | interaction | Protects the creature board and can turn a developed board into pressure.
- 1 Arcane Signet | ramp | Efficient two-color mana acceleration.
- 1 Chromatic Lantern | ramp | Fixes colors while accelerating into the upper end of the deck.
- 1 Commander's Sphere | ramp | Reliable fixing that can become a card later.
- 1 Fellwar Stone | ramp | Efficient multiplayer color fixing.
- 1 Inherited Envelope | ramp | Mana development that remains useful to the deck’s artifact package.
- 1 Oath of the Grey Host | ramp | Theme-friendly mana development with value beyond the initial investment.
- 1 Sol Ring | ramp | Strong colorless acceleration for equipment and larger threats.
- 1 Sword of the Animist | ramp | Equipment that converts attacks into lasting mana development.
- 1 Thought Vessel | ramp | Colorless acceleration that supports a full hand.
- 1 Wayfarer's Bauble | ramp | Early land-based acceleration and color fixing.
- 1 Banishing Light | removal | Versatile answer to troublesome permanents.
- 1 Bitter Triumph | removal | Cheap instant-speed answer to creatures and planeswalkers.
- 1 Generous Gift | removal | Answers any permanent at instant speed.
- 1 Get Lost | removal | Efficient answer to creatures, artifacts, and enchantments.
- 1 Infernal Grasp | removal | Clean, inexpensive creature removal.
- 1 Stroke of Midnight | removal | Flexible instant-speed permanent removal.
- 1 Swords to Plowshares | removal | Premium creature answer that also reinforces the life-total theme.
- 1 Aerith Gainsborough | synergy | A central lifegain payoff that helps turn gained life into board value.
- 1 Aettir and Priwen | synergy | Equipment that rewards attacking with the deck’s lifegain creatures.
- 1 Angel of Vitality | synergy | Improves every life-gain event and grows into a meaningful attacker.
- 1 Compassionate Healer | synergy | Repeatable life gain that supports the commander’s game plan.
- 1 Kor Firewalker | synergy | Resilient lifegain creature that advances the central theme.
- 1 Light of Promise | synergy | Converts repeated life gain into a rapidly growing threat.
- 1 Night Nurse, Healer of Heroes | synergy | Life-total support on a useful legendary creature.
- 1 Prideful Feastling | synergy | Food and life-gain synergy with relevant creature presence.
- 1 Rosie Cotton of South Lane | synergy | Builds a creature into a major threat as life gain and tokens accumulate.
- 1 Second Breakfast | synergy | Efficient Food-based life-gain support and resource generation.
- 1 Wanderbrine Preacher | synergy | Creature-based payoff for repeatedly gaining life.
- 1 Well-Worn Spatula | synergy | Food-adjacent equipment that complements the deck’s life-gain plan.
- 1 White Mage's Staff | synergy | Equipment support that complements white lifegain creatures.
- 1 Angel of Invention | threat | Creates immediate board presence and threatens the air.
- 1 Dawnhand Eulogist | threat | A substantial creature that pressures opponents while supporting the deck’s creature plan.
- 1 Gollum, Silent Slinker // Meager Meal | threat | A difficult-to-block threat with an attached Food-oriented value spell.
- 1 Invisible Woman, Sue Storm | threat | Legendary threat that provides durable board presence.
- 1 Lo and Li, Twin Tutors | threat | A high-impact legendary creature that contributes meaningful pressure.
- 1 Lyra Dawnbringer | threat | Powerful evasive lifelink threat that stabilizes races.
- 1 Minwu, White Mage | threat | A white creature threat that naturally complements lifegain.
- 1 Rabaroo Troop | threat | Creature pressure that helps establish a broad battlefield.
- 1 Rooftop Percher | threat | Evasive creature pressure for closing games through stalled boards.
- 1 Shattered Angel | threat | Flying threat that gains life while punishing opposing land development.
- 1 Sneering Shadewriter | threat | Black creature pressure that benefits from a longer game.
- 1 Victory's Herald | threat | A top-end flying finisher that makes attacks decisive.
- 1 Austere Command | wipe | Flexible reset that can spare the deck’s preferred permanents.
- 1 Fumigate | wipe | Creature reset that replenishes life after a crowded board.
- 1 Vanquish the Horde | wipe | Efficient creature-board reset when opponents commit heavily.

</details>

### 16. a tight budget, owned first

Format: Commander. Theme: aristocrats sacrifice. Pool: owned_first. Shortlist: 229 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $15.99 to buy, $40.96 the whole deck.

**Summary:** This black-white aristocrats deck develops a board through sacrifice-oriented synergy creatures and enchantments, protects its key pieces, and keeps opposing boards in check with efficient removal and wipes. It closes through steady pressure from its creature threats while its draw package helps sustain the board. The list leans on cards already in the library and a straightforward two-color mana base, giving up pricier high-impact options for a lower-cost, lower-powered game plan.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [WARN] `not_owned`: Aron, Benalia's Ruin: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Ayli, Eternal Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bartolomé del Presidio: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Bastion of Remembrance: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Blood Artist: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Elas il-Kor, Sadistic Pilgrim: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Falkenrath Noble: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Meathook Massacre II: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Skullport Merchant: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Smothering Abomination: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Woe Strider: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Yahenni, Undying Partisan: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Zulaport Cutthroat: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Baron Bertram Graywater: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Lord Skitter's Butcher: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Old Flitterfang: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Shilgengar, Sire of Famine: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vengeful Bloodwitch: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vindictive Vampire: the deck needs 1, the collection has 0
- [WARN] `not_owned`: Vampire Gourmand: the deck needs 1, the collection has 0
- [WARN] `not_owned`: High-Society Hunter: the deck needs 1, the collection has 0
- [INFO] `curve_summary`: average mana value 2.85 over 62 nonland cards

<details><summary>The deck list</summary>

- 12 Plains | land | Basic white mana source.
- 17 Swamp | land | Basic black mana source.
- 1 Command Tower | land | Reliable two-color commander mana source.
- 1 Exotic Orchard | land | Flexible color-producing land.
- 1 Grand Coliseum | land | Additional color-fixing land.
- 1 Opal Palace | land | Commander-focused mana source.
- 1 Scavenger Grounds | land | Utility land slot.
- 1 Secluded Courtyard | land | Creature-focused color fixing.
- 1 Unclaimed Territory | land | Creature-focused color fixing.
- 1 Study Hall | land | Utility land slot.
- 1 Sol Ring | ramp | Efficient artifact mana.
- 1 Arcane Signet | ramp | Reliable color-fixing mana rock.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Thought Vessel | ramp | Mana rock that supports longer games.
- 1 Chromatic Lantern | ramp | Artifact ramp and mana fixing.
- 1 Commander's Sphere | ramp | Flexible mana rock.
- 1 Deadly Dispute | ramp | Sacrifice-compatible mana acceleration.
- 1 Explorer's Scope | ramp | Low-cost land-based acceleration.
- 1 Springleaf Drum | ramp | Cheap creature-based mana source.
- 1 White Auracite | ramp | Low-cost artifact mana source.
- 1 Skullclamp | draw | Efficient repeatable card advantage.
- 1 Idol of Oblivion | draw | Artifact card-advantage source.
- 1 Mask of Memory | draw | Combat-based card advantage.
- 1 Tome of Legends | draw | Commander-oriented card advantage.
- 1 Wall of Omens | draw | Early body that replaces itself.
- 1 Inspiring Overseer | draw | Creature-based card advantage.
- 1 Nasty End | draw | Sacrifice-compatible card draw.
- 1 Painful Truths | draw | Efficient spell-based card advantage.
- 1 Stone of Erech | draw | Utility artifact that provides card advantage.
- 1 Puresteel Paladin | draw | Equipment-based card advantage.
- 1 Folk Hero | draw | Creature-focused card advantage.
- 1 Bastion Protector | interaction | Protective support for the commander.
- 1 Frontline Medic | interaction | Creature-based protective interaction.
- 1 Gift of Immortality | interaction | Protective recursion support.
- 1 Swiftfoot Boots | interaction | Low-cost creature protection.
- 1 Together Forever | interaction | Protective recursion support.
- 1 Take Up the Shield | interaction | Efficient protective combat trick.
- 1 Bitter Triumph | removal | Flexible single-target answer.
- 1 Claim the Precious | removal | Single-target removal spell.
- 1 Crib Swap | removal | Creature answer with flexible timing.
- 1 Heartless Act | removal | Efficient creature removal.
- 1 Infernal Grasp | removal | Reliable creature answer.
- 1 Swords to Plowshares | removal | Efficient single-target answer.
- 1 Stroke of Midnight | removal | Flexible permanent removal.
- 1 Austere Command | wipe | Flexible board-clearing option.
- 1 Dusk // Dawn | wipe | Board reset with later value.
- 1 Fumigate | wipe | Creature-focused board reset.
- 1 Aron, Benalia's Ruin | synergy | Low-cost aristocrats support creature.
- 1 Ayli, Eternal Pilgrim | synergy | Sacrifice-oriented support creature.
- 1 Bartolomé del Presidio | synergy | Low-cost sacrifice-oriented creature.
- 1 Bastion of Remembrance | synergy | Aristocrats-focused enchantment support.
- 1 Blood Artist | synergy | Core aristocrats payoff creature.
- 1 Elas il-Kor, Sadistic Pilgrim | synergy | Aristocrats-focused payoff creature.
- 1 Falkenrath Noble | synergy | Creature-based aristocrats payoff.
- 1 Meathook Massacre II | synergy | Aristocrats-focused enchantment support.
- 1 Skullport Merchant | synergy | Sacrifice-oriented value creature.
- 1 Smothering Abomination | synergy | Sacrifice-oriented card-advantage creature.
- 1 Woe Strider | synergy | Sacrifice-oriented support creature.
- 1 Yahenni, Undying Partisan | synergy | Resilient sacrifice-oriented creature.
- 1 Zulaport Cutthroat | synergy | Core aristocrats payoff creature.
- 1 Bill the Pony | threat | Low-cost creature threat.
- 1 Hei Bai, Spirit of Balance | threat | Creature threat for the board plan.
- 1 Namazu Trader | threat | Creature threat that develops the board.
- 1 Vengeful Villagers | threat | Creature threat for sustained pressure.
- 1 Baron Bertram Graywater | threat | Creature threat that fits the black-white plan.
- 1 Lord Skitter's Butcher | threat | Creature threat that supports the board plan.
- 1 Old Flitterfang | threat | Creature threat for the aristocrats shell.
- 1 Shilgengar, Sire of Famine | threat | Creature threat for closing games.
- 1 Vengeful Bloodwitch | threat | Creature threat that adds board pressure.
- 1 Vindictive Vampire | threat | Creature threat for the aristocrats shell.
- 1 Vampire Gourmand | threat | Low-cost creature threat.
- 1 High-Society Hunter | threat | Creature threat that adds board pressure.

</details>

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, two_card_combo. Block findings: 0.

Cost: $81.50 to buy, $265.52 the whole deck.

**Summary:** A mono-red Goblin storm deck that develops a wide battlefield, converts that board into mana and cards, and uses rapid spell chains to fuel overwhelming combat steps or direct damage. It balances token production, tribal payoffs, graveyard recursion, and flexible interaction while preserving the aggressive preconstructed-deck identity.

The quality model grades this deck good against the top lists of the format: few of the cards are ones the top lists of the format play, and that lowers the grade. The mana base serves the colors evenly, and that raises the grade. Nearly every card appears in a top list, and that raises the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 2.51 over 63 nonland cards
- [WARN] `profile_off_band`: the wipe count is 1, and bracket 3 wants 2 to 5
- [WARN] `profile_off_band`: the interaction count is 1, and bracket 3 wants 4 to 12
- [WARN] `two_card_combo`: bracket 3 allows no two-card infinite combo that needs four mana or less, and the deck holds Storm-Kiln Artist + Haze of Rage (speed 5, near two-card)
- [INFO] `cards_trimmed`: the list was 1 card over, so the builder cut this card: Blasphemous Act

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | draw | Cantrip-style targeting spell that scales across the Goblin board with the commander.
- 1 Ancestral Anger | draw | Efficient targeting cantrip for storm turns and combat pressure.
- 1 Arena of Glory | land | Red source that can give a key creature immediate impact.
- 1 Battle Hymn | ramp | Converts a wide Goblin board into explosive red mana.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into direct damage.
- 1 Brightstone Ritual | ramp | Burst mana for deploying Goblins and chaining spells.
- 1 Broadside Bombardiers | removal | Turns expendable tokens into powerful removal or direct damage.
- 1 Castle Embereth | land | Red source with a late-game team pump.
- 1 Chaos Warp | removal | Flexible answer to problematic permanents.
- 1 Conspicuous Snoop | synergy | Goblin-focused card access and tribal synergy.
- 1 Crimson Wisps | draw | Cheap cantrip that grants haste to the whole copied targeting sequence.
- 1 Daring Discovery | synergy | Supports the deck's token-making storm plan.
- 1 Den of the Bugbear | land | Red mana source that becomes a Goblin-producing threat.
- 1 Dragon Fodder | synergy | Early bodies for pressure, mana rituals, and sacrifice effects.
- 1 Dwarven Mine | land | Red land that can add a Goblin body when enabled.
- 1 Empty the Warrens | wincon | Major token payoff for turns with several spells cast.
- 1 Expedite | draw | Low-cost haste cantrip that works especially well with a wide board.
- 1 Faithless Looting | draw | Efficient filtering that stocks the graveyard for flashback turns.
- 1 Fists of Flame | draw | Card-draw combat finisher that rewards storming off.
- 1 Forgotten Cave | land | Red source with cycling when a land is not needed.
- 1 Frontline Heroism | synergy | Supports aggressive combat with the Goblin token force.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal with cycling card flow.
- 1 General Kreat, the Boltbringer | synergy | Tribal payoff that improves the Goblin swarm.
- 1 Goblin Bombardment | removal | Reliable sacrifice outlet that converts tokens into damage.
- 1 Goblin Burrows | land | Colorless utility land that can turn a Goblin into a meaningful attacker.
- 1 Goblin Bushwhacker | synergy | Haste and power boost for explosive token turns.
- 1 Goblin Chieftain | synergy | Anthem and haste lord for the Goblin army.
- 1 Goblin Dark-Dwellers | threat | Recasts an impactful instant or sorcery from the graveyard.
- 1 Goblin Fireleaper | removal | Goblin body that provides repeatable reach in combat.
- 1 Goblin Glasswright // Craft with Pride | ramp | Goblin-based mana support that remains useful as a spell payoff.
- 1 Goblin Lackey | synergy | Accelerates powerful Goblins onto the battlefield through combat.
- 1 Goblin Matron | synergy | Finds the most useful Goblin for the current board state.
- 1 Goblin Negotiation | removal | Tribal removal that leverages a broad battlefield.
- 1 Goblin Trashmaster | removal | Goblin lord that also answers artifacts.
- 1 Goblin Warchief | synergy | Cost reduction and haste make storm turns much more explosive.
- 1 Grapeshot | removal | Storm payoff that can clear small creatures or finish damaged opponents.
- 1 Haze of Rage | wincon | Reusable storm payoff that turns a developed board into lethal combat damage.
- 1 Hidden Volcano | land | Red land with useful late-game flexibility.
- 1 Idol of Oblivion | draw | Consistent card draw alongside frequent token production.
- 1 Impact Tremors | wincon | Converts every Goblin token entering into table-wide pressure.
- 1 Impulsive Pilferer | ramp | Early Goblin that leaves behind a Treasure for storm turns.
- 1 Kher Keep | land | Colorless utility land that supplies expendable creature tokens.
- 1 Krenko's Command | synergy | Efficient token production for a wide-board strategy.
- 1 Krenko, Mob Boss | threat | Premier repeatable Goblin token engine and primary threat.
- 1 Mana Geyser | ramp | High-output ritual that enables large multi-spell turns.
- 25 Mountain | land | Reliable untapped red mana for a mono-red spell-heavy deck.
- 1 Mogg War Marshal | synergy | Multiple Goblin bodies from one card for sacrifice and go-wide payoffs.
- 1 Moggcatcher | threat | Repeatedly tutors key Goblin threats directly to the battlefield.
- 1 Moria Marauder | synergy | Goblin attacker that helps keep pressure and cards flowing.
- 1 Pashalik Mons | removal | Goblin payoff that turns sacrifices and deaths into damage.
- 1 Past in Flames | wincon | Lets the deck replay a stocked graveyard for a major storm turn.
- 1 Quest for the Goblin Lord | synergy | Tribal payoff that rewards building a large Goblin presence.
- 1 Reliquary Tower | land | Colorless utility land that preserves cards accumulated during explosive turns.
- 1 Roaming Throne | synergy | Doubles valuable Goblin triggered abilities after choosing the tribe.
- 1 Ruby Medallion | ramp | Red spell cost reduction supports chaining spells together.
- 1 Rundvelt Hordemaster | synergy | Goblin anthem and card-advantage engine when the tribe dies.
- 1 Sazacap's Brew | draw | Cheap filtering spell that also supports storm count.
- 1 Searslicer Goblin | synergy | Tribal attacker that contributes to the go-wide plan.
- 1 Seething Song | ramp | Efficient ritual for explosive turns.
- 1 Shinka, the Bloodsoaked Keep | land | Red legendary land that gives a relevant combat keyword when needed.
- 1 Siege-Gang Commander | removal | Produces a Goblin squad and converts bodies into direct damage.
- 1 Siege-Gang Lieutenant | removal | Adds tokens and a Goblin-based damage outlet.
- 1 Skirk Prospector | ramp | Sacrifices Goblins for the mana needed to extend storm turns.
- 1 Skullclamp | draw | Exceptional card-draw engine with disposable Goblin tokens.
- 1 Sol Ring | ramp | Fast mana that accelerates the deck's development.
- 1 Spreading Insurrection | wincon | High-impact spell that exploits crowded opposing boards.
- 1 Sting-Slinger | synergy | Goblin tribal piece that supports aggressive board development.
- 1 Storm-Kiln Artist | ramp | Creates Treasure from the deck's many instants and sorceries.
- 1 Swiftfoot Boots | interaction | Protects an important engine creature while granting haste.
- 1 Throne of Eldraine | ramp | Mana rock that offers useful late-game card selection.
- 1 Vandalblast | wipe | Efficient artifact interaction with a strong overloaded mode.
- 1 War Room | land | Colorless utility land that provides repeatable card draw.
- 1 Warren Torchmaster | synergy | Goblin payoff that helps convert the swarm into decisive combat.
- 1 Wild Ride | synergy | Combat-oriented spell that supports explosive attacks.
- 1 Witch's Mark | draw | Cheap filtering and haste support for a key creature.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $170.18 to buy, $170.18 the whole deck.

**Summary:** A five-color Turtle-themed creature deck that develops its mana, builds a resilient board of iconic allies and shelled threats, and converts that board into sustained card advantage and decisive combat. Flexible answers, protective plays, and sweepers keep opponents from disrupting the team’s momentum.

The quality model grades this deck good against the top lists of the format: the cards pair the way the top lists pair them, and that raises the grade. Nearly every card appears in a top list, and that raises the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.28 over 65 nonland cards

<details><summary>The deck list</summary>

- 1 Ash Barrens | land | Flexible fixing that can find a needed basic.
- 1 Big Apple, 3 a.m. | land | Retained thematic multicolor land.
- 1 Cinder Glade | land | Retained dual-color source.
- 1 Command Tower | land | Reliable five-color fixing.
- 1 Dragonskull Summit | land | Retained dual-color source.
- 1 Escape Tunnel | land | Fixes mana while retaining a useful utility land.
- 1 Evolving Wilds | land | Budget fixing for the broad color requirements.
- 1 Fabled Passage | land | Fetches the basic color currently needed.
- 1 Grand Coliseum | land | Retained all-color fixing.
- 1 Hidden Hideout | land | Retained thematic utility land.
- 1 Path of Ancestry | land | Color fixing with tribal upside.
- 1 Rain-Slicked Copse | land | Retained dual-color source.
- 1 Rootbound Crag | land | Retained dual-color source.
- 1 Sodden Verdure | land | Retained dual-color source.
- 1 Thriving Grove | land | Flexible color fixing.
- 1 Thriving Isle | land | Flexible color fixing.
- 1 Thriving Moor | land | Flexible color fixing.
- 1 Turtle Lair | land | Thematic land supporting the Turtle plan.
- 1 Undergrowth Stadium | land | Retained dual-color source.
- 1 Vernal Fen | land | Retained dual-color source.
- 1 Vibrant Cityscape | land | Retained flexible fixing.
- 3 Forest | land | Basic green sources for early ramp and Turtle spells.
- 3 Island | land | Basic blue sources for draw and interaction.
- 3 Mountain | land | Basic red sources for the deck's Turtle suite.
- 2 Plains | land | Basic white sources for protection and answers.
- 2 Swamp | land | Basic black sources for removal and support.
- 1 Acidic Slime | removal | Theme-friendly creature that answers problematic permanents.
- 1 April O'Neil, Live on the Scene | draw | Retained Turtle-supporting card advantage.
- 1 Arcade Cabinet | synergy | Retained thematic support piece.
- 1 Arcane Signet | ramp | Efficient color fixing.
- 1 Assassin's Trophy | removal | Versatile answer to any troublesome permanent.
- 1 Baxter, Fly in the Ointment | synergy | Retained thematic character support.
- 1 Bebop, Skull & Crossbones | threat | Retained thematic character threat.
- 1 Big Mother Mouser | synergy | Retained themed board presence.
- 1 Biogenic Ooze | threat | Retained token-making threat.
- 1 Blasphemous Act | wipe | Efficient reset when creature boards get crowded.
- 1 Casey Jones, Back Alley Brute | threat | Retained thematic combat threat.
- 1 Chromatic Lantern | ramp | Fixes all colors while advancing mana.
- 1 Coin of Mastery | synergy | Retained thematic utility artifact.
- 1 Continue? | interaction | Retained reactive thematic spell.
- 1 Corpsejack Menace | synergy | Retained counter-focused payoff.
- 1 Cultivate | ramp | Reliable land-based fixing and acceleration.
- 1 Dimension X Pizzasaur | synergy | Retained thematic artifact creature.
- 1 Double Jump // Flying Kick | interaction | Retained flexible combat trick and response.
- 1 Electric Seaweed | synergy | Retained defensive thematic permanent.
- 1 Endless Foot Assault | synergy | Retained thematic combat support.
- 1 Everything Pizza | synergy | Retained thematic value artifact.
- 1 Exploding Barrel | synergy | Retained thematic utility artifact.
- 1 Fast Forward | interaction | Retained reactive spell for protecting momentum.
- 1 Foot Chopper | synergy | Retained equipment support.
- 1 Game Over | wincon | Retained thematic finisher.
- 1 Harmonize | draw | Straightforward refill after committing creatures.
- 1 Here Comes a New Hero! | synergy | Retained thematic creature support.
- 1 High Score | synergy | Retained long-game value engine.
- 1 Krang, the All-Powerful | threat | Retained thematic artifact threat.
- 1 Leatherhead, Iron Gator | threat | Retained thematic creature threat.
- 1 Lessons from Life | synergy | Retained thematic value spell.
- 1 Level Up | synergy | Retained enhancement for the creature plan.
- 1 Michelangelo, the Heart | synergy | Retained Turtle-themed support legend.
- 1 Mole Module | synergy | Retained thematic vehicle support.
- 1 Mona Lisa, Science Geek | synergy | Retained thematic character support.
- 1 Ninja Pizza | synergy | Retained thematic value enchantment.
- 1 Rat King, Pale Piper | threat | Retained thematic creature threat.
- 1 Ray Fillet, Wave Warrior | synergy | Retained thematic creature support.
- 1 Roadkill Rodney | synergy | Retained thematic artifact creature.
- 1 Rocksteady, Mutant Marauder | threat | Retained thematic creature threat.
- 1 Shellshock | interaction | Retained thematic response spell.
- 1 Sol Ring | ramp | Fast, reliable artifact acceleration.
- 1 Special Move | interaction | Retained flexible reactive support.
- 1 Splinter, the Mentor | synergy | Retained thematic support legend.
- 1 Steelbane Hydra | removal | Creature-based answer that fits the board-focused plan.
- 1 Super Combo | wincon | Retained thematic closing tool.
- 1 Swift Demise | removal | Retained targeted answer.
- 1 Tempestra, Dame of Games | synergy | Retained thematic value creature.
- 1 Together Forever | synergy | Retained resilience for the creature plan.
- 1 Tokka & Rahzar, Unsupervised | ramp | Turtle-themed mana development.
- 1 Vanquish the Horde | wipe | Additional board reset for creature-heavy tables.
- 1 Vigor | synergy | Protects the board while rewarding creature combat.
- 1 Voracious Hydra | threat | Retained scalable creature threat.
- 1 Wave Goodbye | synergy | Retained thematic disruption spell.
- 1 Beast Whisperer | draw | Turns the creature-heavy Turtle plan into sustained cards.
- 1 Beast Within | removal | Broad permanent removal in green.
- 1 Esper Sentinel | draw | Early card advantage that discourages opponents from racing ahead.
- 1 Farseek | ramp | Efficient fixing for the multicolor mana base.
- 1 Fellwar Stone | ramp | Low-cost acceleration that usually fixes several colors.
- 1 Garruk's Uprising | draw | Creature-focused draw engine with combat upside.
- 1 Generous Gift | removal | Flexible instant-speed answer to any permanent.
- 1 Mystic Remora | draw | Efficient early engine against spell-heavy opponents.
- 1 Nature's Lore | ramp | Efficient land ramp and fixing.
- 1 Phyrexian Arena | draw | Steady card flow for the midgame.
- 1 Return of the Wildspeaker | draw | Refills the hand from a developed Turtle board.

</details>

### 19. the Hobbit family, two colours

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 174 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $485.38 to buy, $485.38 the whole deck.

**Summary:** A Sultai Elven midrange deck built around assembling a resilient woodland court, accelerating through mana creatures and artifacts, and converting a broad creature board into sustained value. It combines steady card advantage with efficient disruption, selective sweeping effects, and enough threats to close games through combat after controlling the table.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves one color far better than another, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 3.16 over 63 nonland cards
- [WARN] `profile_off_band`: the worst color's share of the sources it needs is 0.49, and bracket 3 wants 0.85 or more (sources of requirement: U 13.75 of 19, B 12.75 of 26, G 23 of 19)

<details><summary>The deck list</summary>

- 1 Arcane Signet | ramp | Reliable fixing for all three colors.
- 1 Delighted Halfling | ramp | Early acceleration that supports legendary spells.
- 1 Elvish Archdruid | ramp | Elf-based mana production with a substantial ceiling.
- 1 Elvish Mystic | ramp | Efficient early green acceleration.
- 1 Elven Chorus | ramp | Turns the Elf-heavy creature suite into mana.
- 1 Mox Amber | ramp | Fast legendary-focused acceleration.
- 1 Necklace of Girion | ramp | Artifact mana fixing.
- 1 Wayfarer's Bauble | ramp | Basic-land ramp and color support.
- 1 Wood Elves | ramp | Creature-based land ramp that develops the board.
- 1 Woodland Weavemaster | ramp | Elf synergy and mana development.
- 1 Elvish Visionary | draw | A cheap Elf that replaces itself.
- 1 Beorn the Fierce | draw | Creature-based card advantage.
- 1 Bilbo Baggins, Burglar // Take a Glance | draw | Flexible adventure card advantage.
- 1 Fateful Discovery | draw | Persistent source of cards.
- 1 Gandalf, Shadow's Foe | draw | A substantial value creature.
- 1 Gandalf, Wandering Wizard | draw | Wizard-themed card advantage.
- 1 Hithlain Knots | draw | Instant-speed card flow.
- 1 Lórien Revealed | draw | Early land access or later card advantage.
- 1 Night's Whisper | draw | Efficient low-cost card draw.
- 1 Palantír of Orthanc | draw | Repeatable artifact card advantage.
- 1 Plunder the Trollshaws | draw | Instant-speed card advantage.
- 1 Confusticate and Bebother | interaction | Protects the board and disrupts opposing plays.
- 1 Elrond, Moon-Reader | interaction | Legendary Elf utility and reactive play.
- 1 Bilbo's Ring | interaction | Equipment-based protection and utility.
- 1 Mithril Coat | interaction | Protects a crucial creature at instant speed.
- 1 My Precious // Allure of Power | interaction | Flexible protection and disruptive interaction.
- 1 Old Fat Spider Can't See Me | interaction | Defensive disruption that supports the theme.
- 1 Stern Scolding | interaction | Efficient answer to early creatures and commanders.
- 1 The One Ring | interaction | A powerful defensive utility artifact.
- 1 Bilbo's Deadly Slice | removal | Efficient creature removal.
- 1 Bitter Downfall | removal | Broad answer to problematic permanents.
- 1 Colossal Whale | removal | Creature threat that also removes opposing creatures.
- 1 Enchanted River's Grasp | removal | Blue removal that answers troublesome permanents.
- 1 Merciless Executioner | removal | Edict removal attached to a creature.
- 1 Orcish Bowmasters | removal | Punishes opposing card draw while controlling small creatures.
- 1 Quarrel | removal | Flexible spot removal.
- 1 The Black Arrow | removal | Equipment-based removal utility.
- 1 Uneasy Partings | removal | Instant-speed removal option.
- 1 Gnashing of Teeth | wipe | Creature sweep that clears a developed board.
- 1 Languish | wipe | Efficient black creature sweeper.
- 1 Raise the Palisade | wipe | Tribal sweeper that can spare the Elf board.
- 1 Attercop | threat | Large themed creature that pressures opponents.
- 1 Boughside Wanderers | threat | Elf body that advances the tribal battlefield.
- 1 Cantankerous Keepers | threat | Elf creature for board presence.
- 1 Celeborn the Wise | threat | Legendary Elf payoff and resilient threat.
- 1 Elven Raft-Steerer | threat | Evasive Elf board presence.
- 1 Galadhrim Guide | threat | Low-cost Elf attacker and tribal body.
- 1 Galion, Elvenking's Butler | threat | Legendary Elf that supports the courtly theme.
- 1 Guardian of the Halls | threat | Elf Soldier that helps establish the board.
- 1 Haunt of the Dead Marshes | threat | A black-green creature threat with thematic value.
- 1 Mirkwood Elk | threat | Efficient green creature for combat pressure.
- 1 Mirkwood Pathmaker | threat | Elf body that complements the tribal plan.
- 1 Nimrodel Watcher | threat | Elf threat that supports a wide battlefield.
- 1 Along the Crooked Way | synergy | Thematic permanent supporting the deck's long game.
- 1 Down in the Valley | synergy | Saga value that complements the Middle-earth theme.
- 1 Elvenking's Harper | synergy | Elf support piece for the creature-focused plan.
- 1 Grey Havens Navigator | synergy | Elf utility creature with thematic cohesion.
- 1 Lothlórien Lookout | synergy | Additional Elf presence for tribal payoffs.
- 1 Mirkwood Meditator | synergy | Elf utility that strengthens the tribal core.
- 1 Mirkwood Nurturer | synergy | Elf-focused support for the developing board.
- 1 Supper for Spiders | synergy | Theme-forward support spell.
- 1 Thranduil, Sindarin Liege // Silvan Rally | synergy | A powerful Elf-themed legendary payoff.
- 1 Uncover the Moon-Letters | synergy | Thematic value engine for the longer game.
- 17 Forest | land | Primary green source base for the Elf-heavy spell suite.
- 10 Island | land | Blue sources for reactive spells and card advantage.
- 9 Swamp | land | Black sources for removal and value spells.

</details>

### 20. the Hobbit family, a delegated commander

Format: Commander. Theme: dragons. Pool: any_card. Shortlist: 77 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $521.10 to buy, $521.10 the whole deck.

**Summary:** A mono-red Middle-earth dragon deck that accelerates through treasures and artifacts, develops a supporting host of Dwarves and raiders, and uses equipment to make its largest attackers especially dangerous. Its game plan combines resilient resource development, steady card advantage, focused burn, and dramatic battlefield-clearing effects before overwhelming opponents with legendary creatures and immense monsters.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. The commander does not place in cEDH events, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.32 over 63 nonland cards
- [INFO] `basics_added`: the list was 1 card short, so the builder added 1 basic land

<details><summary>The deck list</summary>

- 31 Mountain | land | Basic red mana base for a mono-red dragon strategy.
- 1 Dragon-Cursed Halls | land | Thematic utility land.
- 1 Hobbit Hole | land | Utility land supporting the deck’s resource plan.
- 1 Rogue's Passage | land | Lets a large dragon connect through opposing boards.
- 1 The Lonely Mountain | land | Thematic red-producing land.
- 1 Treasure Vault | land | Flexible artifact land and late-game mana outlet.
- 1 Balin, Loremaster | draw | Repeatable card advantage.
- 1 Key to the Side-Door | draw | Provides card access while supporting artifact synergies.
- 1 Palantír of Orthanc | draw | Steady card advantage.
- 1 Ragged Short Spear | draw | Equipment-based card advantage.
- 1 Thrór's Map | draw | Provides card selection and value.
- 1 Óin the Brave | draw | Creature-based card advantage.
- 1 The One Ring | draw | Powerful ongoing card advantage with protection utility.
- 1 Last Light of Durin's Day | draw | Sustains the hand while rewarding the deck’s permanent-heavy plan.
- 1 Arcane Signet | ramp | Reliable early mana acceleration.
- 1 Bag End Banquet | ramp | Artifact mana acceleration.
- 1 Burn, Burn, Tree and Fern | ramp | Advances mana while fitting the setting’s theme.
- 1 Cavern-Hoard Dragon | ramp | Dragon threat that creates mana resources.
- 1 Dragon's Desire | ramp | Accelerates into larger threats.
- 1 Mox Amber | ramp | Efficient legendary-focused acceleration.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment that helps generate mana advantage.
- 1 The Reaver Cleaver | ramp | Turns combat damage into a large mana burst.
- 1 Wayfarer's Bauble | ramp | Reliable land-based ramp.
- 1 The Misty Mountains Cold | ramp | Thematic mana development.
- 1 Battle-Scarred Goblin | removal | Creature-based targeted removal.
- 1 Fire of Orthanc | removal | Efficient targeted removal.
- 1 Gandalf, Spark Starter | removal | Flexible creature-based removal.
- 1 Giant's Boulder | removal | Reusable thematic removal tool.
- 1 Goblin Cratermaker | removal | Versatile answer to creatures and artifacts.
- 1 Improvised Club | removal | Low-cost targeted removal.
- 1 Inferno Titan | removal | Large threat with immediate damage-based removal.
- 1 Smite the Deathless | removal | Clean answer to problematic creatures.
- 1 The Black Arrow | removal | Thematic removal attached to an artifact.
- 1 Call Forth the Tempest | wipe | Resets crowded opposing boards.
- 1 Desolation of Smaug | wipe | Thematic board reset.
- 1 Glóin the Mighty // Easy Pickings | wipe | Creature with a flexible sweeping Adventure.
- 1 Bilbo's Ring | interaction | Protective and evasive utility for key attackers.
- 1 Dwarven Mattock | interaction | Equipment utility for combat and problem permanents.
- 1 Mithril Coat | interaction | Protects the commander or a pivotal threat.
- 1 Andúril, Flame of the West | interaction | Combat-focused legendary equipment.
- 1 Andúril, Narsil Reforged | interaction | Legendary equipment that improves combat positioning.
- 1 Glamdring | interaction | Flexible equipment utility.
- 1 Long-Lost Lances | interaction | Combat utility and creature enhancement.
- 1 Sting, Bilbo's Sword | interaction | Efficient equipment for protecting and empowering attackers.
- 1 Dáin Ironfoot | synergy | Supports the deck’s Dwarf and artifact subtheme.
- 1 Dwarven Mauler | synergy | Thematic creature that benefits from the equipment package.
- 1 Dwarven Warriors | synergy | Adds bodies that work with the deck’s combat plan.
- 1 Gundabad Opportunist | synergy | Supports the aggressive creature package.
- 1 Guttersnipe | synergy | Converts removal and utility spells into extra damage.
- 1 Iron Hills Stalwart | synergy | Thematic body for the creature-based combat plan.
- 1 Orcish Siegemaster | synergy | Adds pressure alongside the deck’s aggressive threats.
- 1 Tidings of War | synergy | Reinforces the deck’s battlefield-focused strategy.
- 1 Thorin, Company's Leader | synergy | Supports the legendary and creature themes.
- 1 Goblin-town Flunkies | synergy | Low-cost thematic attacker and support creature.
- 1 Desert Were-Worm | threat | Large dragon-adjacent finisher.
- 1 Smaug, the Great Calamity // Spew Flame | threat | Powerful dragon threat with flexible damage.
- 1 Stone-Giant of High Pass | threat | Large battlefield threat.
- 1 Olog-hai Crusher | threat | High-impact combat finisher.
- 1 Oliphaunt | threat | Large thematic creature that pressures life totals.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Versatile legendary threat.
- 1 Goblin Fireleaper | threat | Aggressive creature that adds reach.
- 1 Long-Bodied Grey Dog | threat | Creature that contributes to pressure and development.
- 1 Misty Mountains Raider | threat | Thematic attacker for the combat plan.
- 1 Snowslope Hunter | threat | Creature threat that helps maintain pressure.
- 1 Thorin, Mountain-king | threat | Legendary threat with meaningful battlefield impact.
- 1 Troop of Ponies | threat | Broadens the creature assault.
- 1 Bothersome Noisemaker | other | Thematic utility creature.
- 1 Old Thrush | other | Low-cost thematic utility body.
- 1 Getaway Barrel | other | Flexible artifact utility.

</details>

### 21. the Hobbit family, mana from outside

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 156 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band, profile_off_band. Block findings: 0.

Cost: $894.01 to buy, $894.01 the whole deck.

**Summary:** Smaug leads a Rakdos Middle-earth arsenal built around legendary relics, ravenous Goblins, and towering monsters. Early resource development supports a steady stream of equipment-enhanced attackers, while plentiful removal keeps rival boards contained. The deck wins by converting artifact value and combat pressure into overwhelming Dragon-led finishes.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. The deck makes less mana on turn four than the norm, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade.

- [INFO] `curve_summary`: average mana value 3.22 over 63 nonland cards
- [WARN] `profile_off_band`: the mana available on turn four is 4.17, and bracket 3 wants 4.2 or more (mean over 10000 hands)
- [WARN] `profile_off_band`: the turns the commander comes down after its mana value is 0.91, and bracket 3 wants 0.8 at most (the commander costs 7 and comes down on turn 7.91 on average)
- [WARN] `outside_requested_set`: the sets you named do not hold 15 cards: Ancient Tomb, Arid Mesa, Blood Crypt, Bloodstained Mire, Command Tower, Dragonskull Summit, Exotic Orchard, Fabled Passage, Marsh Flats, Myriad Landscape, Polluted Delta, Scalding Tarn, Stardew Valley, Verdant Catacombs, Wooded Foothills

<details><summary>The deck list</summary>

- 12 Mountain | land | Reliable red source for the deck’s core dragon, Goblin, and burn effects.
- 8 Swamp | land | Reliable black source for removal and graveyard-focused spells.
- 1 Ancient Tomb | land | Accelerates expensive artifacts and major threats.
- 1 Myriad Landscape | land | Provides a land-based source of later mana development.
- 1 Stardew Valley | land | Utility mana source within the land package.
- 1 Treasure Vault | land | Artifact land that can convert into a burst of Treasure.
- 1 Blood Crypt | land | Untapped dual source when needed.
- 1 Command Tower | land | Fixes both commander colors cleanly.
- 1 Dragonskull Summit | land | Dual land for the deck’s main colors.
- 1 Exotic Orchard | land | Flexible multiplayer color fixing.
- 1 Fabled Passage | land | Finds the needed basic color while thinning the deck.
- 1 Bloodstained Mire | land | Fetches either primary basic land type or the dual land.
- 1 Polluted Delta | land | Finds black sources and the dual land.
- 1 Marsh Flats | land | Finds black sources and the dual land.
- 1 Arid Mesa | land | Finds red sources and the dual land.
- 1 Scalding Tarn | land | Finds red sources and the dual land.
- 1 Wooded Foothills | land | Finds red sources and the dual land.
- 1 Verdant Catacombs | land | Finds black sources and the dual land.
- 1 Arcane Signet | ramp | Efficient, dependable color fixing.
- 1 Mox Amber | ramp | Cheap acceleration alongside the deck’s legendary creatures.
- 1 Wayfarer's Bauble | ramp | Early land ramp that fixes either main color.
- 1 Bag End Banquet | ramp | Artifact-based mana development that supports the deck’s resource plan.
- 1 Burn, Burn, Tree and Fern | ramp | Saga-based mana acceleration with thematic flavor.
- 1 Dragon's Desire | ramp | Advances mana toward the deck’s larger creatures.
- 1 The Misty Mountains Cold | ramp | Provides additional thematic mana development.
- 1 The Reaver Cleaver | ramp | Turns combat damage into a substantial Treasure advantage.
- 1 Orcrist, Goblin-cleaver | ramp | Equipment that contributes to the artifact-based mana plan.
- 1 Smaug the Magnificent | ramp | A dragon threat that also expands available resources.
- 1 Gollum, Riddle Master | draw | Thematic creature-based card advantage.
- 1 Key to the Side-Door | draw | Reusable artifact card advantage.
- 1 Night's Whisper | draw | Low-cost, efficient card selection.
- 1 Palantír of Orthanc | draw | Sustained card advantage with a threatening payoff.
- 1 Rage into the Valley | draw | Refills resources while supporting the aggressive plan.
- 1 Ragged Short Spear | draw | Equipment that supplies incremental card advantage.
- 1 Reverent Howl | draw | Instant-speed card draw for maintaining momentum.
- 1 The Master of Lake-town | draw | Creature-based value engine.
- 1 The Sackville-Bagginses | draw | Provides recurring thematic card advantage.
- 1 Thrór's Map | draw | Artifact card advantage that fits the treasure-and-relic package.
- 1 Óin the Brave | draw | Dwarf value creature that keeps cards flowing.
- 1 Bilbo's Ring | interaction | Protective equipment with flexible utility.
- 1 Dwarven Mattock | interaction | Equipment-based answer to troublesome permanents.
- 1 Mithril Coat | interaction | Protects a key creature from removal.
- 1 My Precious // Allure of Power | interaction | Flexible equipment and instant-speed disruption.
- 1 The One Ring | interaction | Protects its controller while providing long-term value.
- 1 Along the Crooked Way | interaction | Versatile enchantment utility for contested boards.
- 1 Getaway Barrel | interaction | Artifact utility that can disrupt attacks and preserve creatures.
- 1 Smaug's Fury | interaction | Instant-speed tactical effect that supports combat pressure.
- 1 Azog, Moria's Ruin | removal | Legendary creature that answers opposing threats.
- 1 Battle-Scarred Goblin | removal | Creature-based removal that contributes to board presence.
- 1 Bilbo's Deadly Slice | removal | Efficient spot removal.
- 1 Bitter Downfall | removal | Flexible answer to problematic creatures or planeswalkers.
- 1 Fire of Orthanc | removal | Direct removal for creatures and other key targets.
- 1 Goblin Cratermaker | removal | Reusable body with flexible removal options.
- 1 Improvised Club | removal | Low-cost instant-speed creature interaction.
- 1 Orcish Bowmasters | removal | Punishes opposing card draw while removing small creatures.
- 1 Smite the Deathless | removal | Clean instant-speed answer to a major threat.
- 1 Bolg, Erebor's Reckoning | wipe | Creature-based mass removal option.
- 1 Desolation of Smaug | wipe | Thematic board reset that clears opposing development.
- 1 Languish | wipe | Efficient sweeper against creature-heavy boards.
- 1 Supper for Spiders | synergy | Supports the deck’s artifact, food, and creature-resource themes.
- 1 Andúril, Flame of the West | synergy | Equipment that rewards committing powerful creatures to combat.
- 1 Andúril, Narsil Reforged | synergy | Legendary equipment that strengthens the creature suite.
- 1 Glamdring | synergy | Equipment that improves combat-focused value creatures.
- 1 Goblin Plate Mail | synergy | Supports the Goblin contingent while enhancing combat.
- 1 Long-Lost Lances | synergy | Equipment that helps large attackers break through.
- 1 Sting, Bilbo's Sword | synergy | Low-cost legendary equipment for the artifact package.
- 1 Well-Worn Spatula | synergy | Flavorful equipment that strengthens creature combat lines.
- 1 Down, Down to Goblin-town | synergy | Builds around the deck’s Goblin and under-mountain themes.
- 1 Tidings of War | synergy | Reinforces the deck’s creature-forward aggressive plan.
- 1 Cavern-Hoard Dragon | threat | Large Dragon finisher that rewards the deck’s artifact resources.
- 1 Smaug, Wicked Worm | threat | A marquee Dragon threat with a punishing battlefield presence.
- 1 Smaug, the Great Calamity // Spew Flame | threat | Dragon threat with a flexible removal-backed adventure.
- 1 Inferno Titan | threat | Immediate damage and a powerful combat-closing body.
- 1 Sauron, the Lidless Eye | threat | Legendary late-game threat that pressures opposing boards.
- 1 Gandalf, Goblins' Bane // Flameshape | threat | Versatile legend that supplies pressure and spell utility.
- 1 The Great Goblin | threat | Goblin legend that anchors the deck’s under-mountain forces.
- 1 Olog-hai Crusher | threat | Large creature that applies sustained combat pressure.
- 1 Troll of Khazad-dûm | threat | Substantial body that complements the dark mountain theme.
- 1 Desert Were-Worm | threat | Large evasive-style threat for closing stalled games.
- 1 Oliphaunt | threat | High-impact creature that adds a formidable combat presence.
- 1 Great Ugly-Looking Goblin // Clap! Snap! | threat | Flexible Goblin threat with an attached spell option.

</details>

### 22. a set family and a card from outside it

Format: Commander. Theme: hobbits and food. Pool: any_card. Shortlist: 175 names.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $471.85 to buy, $471.85 the whole deck.

**Summary:** This deck develops an Elf-centered board while using inexpensive ramp and steady draw to keep deploying creatures and resources. It wins by establishing battlefield pressure with its creature suite, backed by targeted answers and reset buttons when opposing boards get ahead. The mana base is deliberately stable and untapped, giving up utility-land flexibility and the explosiveness of a more aggressive mana package for dependable three-color development.

The quality model grades this deck below the precon baseline against the top lists of the format: many of the cards appear in no top list, and that lowers the grade. More opening hands hold two to four lands than the norm, and that raises the grade. The mana base serves the colors evenly, and that raises the grade.

- [INFO] `curve_summary`: average mana value 3.37 over 63 nonland cards
- [WARN] `outside_requested_set`: the sets you named do not hold 1 card: Sol Ring

<details><summary>The deck list</summary>

- 1 Elvish Visionary | draw | Provides a compact draw piece.
- 1 Fateful Discovery | draw | Adds another dedicated draw card.
- 1 Hithlain Knots | draw | Supplies draw from an instant-speed slot.
- 1 Last March of the Ents | draw | Adds a draw spell to keep resources flowing.
- 1 Lórien Revealed | draw | Provides a draw-focused sorcery.
- 1 Night's Whisper | draw | Adds efficient dedicated draw.
- 1 Palantír of Orthanc | draw | Provides a lasting draw-focused artifact.
- 1 Plunder the Trollshaws | draw | Adds an instant draw option.
- 1 Thrór's Map | draw | Provides another draw artifact.
- 1 Uncover the Moon-Letters | draw | Adds a draw-focused enchantment.
- 1 Beorn the Fierce | draw | Adds a creature-based draw card.
- 1 Confusticate and Bebother | interaction | Provides a dedicated interaction spell.
- 1 Bilbo's Ring | interaction | Adds interaction from an artifact slot.
- 1 Elrond, Moon-Reader | interaction | Provides interaction on an Elf creature.
- 1 Mithril Coat | interaction | Adds a protective interaction artifact.
- 1 My Precious // Allure of Power | interaction | Provides interaction with flexible card form.
- 1 Old Fat Spider Can't See Me | interaction | Adds interaction through an enchantment.
- 1 The One Ring | interaction | Provides an additional interaction artifact.
- 1 Stern Scolding | interaction | Adds a compact interaction spell.
- 1 Sol Ring | ramp | Kept as requested and fills a ramp slot.
- 1 Arcane Signet | ramp | Provides reliable artifact ramp.
- 1 Delighted Halfling | ramp | Adds creature-based ramp.
- 1 Elven Chorus | ramp | Provides enchantment-based ramp.
- 1 Elvish Archdruid | ramp | Adds Elf-centered creature ramp.
- 1 Mox Amber | ramp | Provides a low-cost ramp artifact.
- 1 Silvan Reveler | ramp | Adds another Elf ramp creature.
- 1 Through the Forest Gate | ramp | Provides sorcery ramp.
- 1 Wood Elves | ramp | Adds a creature-based ramp option.
- 1 Woodland Weavemaster | ramp | Provides another Elf ramp body.
- 1 Bilbo's Deadly Slice | removal | Provides dedicated removal.
- 1 Bitter Downfall | removal | Adds an instant-speed removal spell.
- 1 Crude Bent Blade | removal | Provides removal from an artifact slot.
- 1 Enchanted River's Grasp | removal | Adds enchantment-based removal.
- 1 Giant's Boulder | removal | Provides artifact-based removal.
- 1 Merciless Executioner | removal | Adds creature-based removal.
- 1 Orcish Bowmasters | removal | Provides a compact removal creature.
- 1 Quarrel | removal | Adds a dedicated removal instant.
- 1 Uneasy Partings | removal | Provides another flexible removal spell.
- 1 Gnashing of Teeth | wipe | Provides a dedicated board wipe.
- 1 Languish | wipe | Adds a second board wipe.
- 1 Raise the Palisade | wipe | Provides a third board-clearing option.
- 1 Supper for Spiders | synergy | Supplies the shortlist's dedicated synergy card.
- 1 Arwen, Weaver of Hope | synergy | An Elf creature that supports the deck's Elf-centered plan.
- 1 Boughside Wanderers | synergy | Adds an Elf body to the tribal core.
- 1 Cantankerous Keepers | synergy | Adds another Elf creature for the shared plan.
- 1 Celeborn the Wise | synergy | Provides a legendary Elf for the deck's core theme.
- 1 Galion, Elvenking's Butler | synergy | Reinforces the Elvenking-themed Elf package.
- 1 Guardian of the Halls | synergy | Adds an Elf Soldier to the shared creature base.
- 1 Lothlórien Lookout | synergy | Provides another Elf for the tribal package.
- 1 Mirkwood Meditator | synergy | Adds an Elf Druid to support the central theme.
- 1 Mirkwood Nurturer | synergy | Provides another Elf creature in the synergy suite.
- 1 Attercop | threat | A large creature threat for battlefield pressure.
- 1 Colossal Whale | threat | Provides a substantial creature threat.
- 1 Dreaded Bat-Cloud | threat | Adds an evasive creature threat.
- 1 Gigantic Big Bear | threat | Provides a sizeable creature threat.
- 1 Great Fierce Bee | threat | Adds another creature that can pressure the table.
- 1 Haunt of the Dead Marshes | threat | Provides a Nightmare Elf threat.
- 1 Mirkwood Elk | threat | Adds a creature threat that fits the woodland theme.
- 1 Mirkwood Pathmaker | threat | Provides an Elf creature threat.
- 1 Nasty Little Rabbit | threat | Adds another creature for board pressure.
- 1 Rhovanion Rampager | threat | Provides a Wolf creature threat.
- 1 Troll of Khazad-dûm | threat | Adds a Troll threat to diversify the creature suite.
- 1 Willow-Wind | threat | Provides an Elemental creature threat.
- 12 Forest | land | Provides a stable basic green mana base.
- 12 Island | land | Provides a stable basic blue mana base.
- 12 Swamp | land | Provides a stable basic black mana base.

</details>

### 23. two set families at once

Format: Commander. Theme: artifacts. Pool: any_card. Shortlist: 203 names.

Cards: 99 main, 0 sideboard. Repair turn: yes, for profile_off_band. Block findings: 0.

Cost: $327.36 to buy, $327.36 the whole deck.

**Summary:** Kíli the Resourceful leads a mono-white Dwarf and artifact deck built around legendary gear, creature tokens, and a resilient battlefield presence. The deck develops mana through artifacts, turns small creatures into cards and resources, controls opposing permanents with flexible answers, and closes through equipped attackers, evasive finishers, and overwhelming combat boards.

The quality model grades this deck good against the top lists of the format: the commander does not place in cEDH events, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. Few of the cards are ones the top lists of the format play, and that lowers the grade.

Summary rules claims (F-26): the format play

- [INFO] `curve_summary`: average mana value 3.07 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the count of nonbasic lands that make only colorless mana is 5, and bracket 3 wants 4 at most (Fountainport, Reliquary Tower, Rogue's Passage, Swarmyard, Treasure Vault)
- [INFO] `basics_added`: the list was 3 cards short, so the builder added 3 basic lands

<details><summary>The deck list</summary>

- 1 Academy Manufactor | synergy | Converts artifact-token production into broader resource generation.
- 1 Andúril, Flame of the West | threat | A powerful Equipment that turns equipped attackers into meaningful pressure.
- 1 Andúril, Narsil Reforged | threat | A resilient legendary Equipment that rewards combat-focused play.
- 1 Angel of the Ruins | removal | Artifact and enchantment removal attached to a sizeable evasive body.
- 1 Arcane Signet | ramp | Reliable early color fixing and acceleration.
- 1 Bag End Banquet | ramp | Artifact-based mana development that supports the deck's resource theme.
- 1 Baird, Steward of Argive | interaction | Taxes opposing attacks and protects a developing board.
- 1 Banishing Light | removal | Flexible answer to troublesome nonland permanents.
- 1 Blade Splicer | synergy | Produces artifact material while contributing to a go-wide board.
- 1 Bumbleflower's Sharepot | removal | Versatile artifact removal that fits the artifact shell.
- 1 Burnished Hart | ramp | Permanent land-based acceleration from an artifact creature.
- 1 Carrot Cake | synergy | Creates multiple useful artifact tokens and a body for the board.
- 1 Caretaker's Talent | draw | Sustained card advantage for a deck that makes creatures and tokens.
- 1 Circuit Mender | draw | Replaces itself while providing a useful artifact creature.
- 1 Cut a Deal | draw | Efficient multiplayer card refill.
- 1 Dáin, Lord of the Iron Hills | synergy | A Dwarf payoff that improves the deck's primary creature type.
- 1 Dawn of a New Age | draw | Ongoing card advantage from a board of creatures.
- 1 Dawn's Truce | interaction | Protects the board from removal and unfavorable combat.
- 1 Dusk // Dawn | wipe | A creature reset with a valuable graveyard-to-hand follow-up.
- 1 Eagles of the North | threat | Evasive creature pressure with an additional combat trick mode.
- 1 Fellwar Stone | ramp | Low-cost mana acceleration.
- 1 Fíli the Pathfinder | threat | A thematic legendary Dwarf that provides an impactful threat.
- 1 Fiend Hunter | removal | Creature-based exile removal that benefits from creature synergies.
- 1 Fountainport Bell | draw | Repeatable artifact-based card selection and draw.
- 1 Galadriel's Dismissal | interaction | Flexible protection or tempo interaction at instant speed.
- 1 Generous Gift | removal | Answers nearly any problematic permanent.
- 1 Heirloom Epic | draw | Artifact-based card advantage that rewards the deck's creature plan.
- 1 Idol of Oblivion | draw | Efficient repeatable draw alongside token production.
- 1 Inspiring Overseer | draw | A creature that immediately replaces itself while adding an evasive body.
- 1 Iron Hills Blacksmith | synergy | Supports Equipment and artifact creatures in the Dwarf shell.
- 1 Jazal Goldmane | threat | Turns a developed creature board into a lethal combat force.
- 1 Karn, the Great Creator | threat | A durable artifact-focused threat that disrupts opposing artifacts.
- 1 Loran of the Third Path | removal | Efficient artifact or enchantment removal with optional card advantage.
- 1 Martial Coup | wipe | Resets crowded boards while leaving behind an army.
- 1 Maskwood Nexus | synergy | Unifies creature types and supplies additional artifact creature material.
- 1 Mentor of the Meek | draw | Converts the deck's smaller creatures and tokens into cards.
- 1 Mind Stone | ramp | Early acceleration that later cycles for a card.
- 1 Mithril Coat | interaction | Protects the commander or a key creature from destruction.
- 1 Mox Amber | ramp | Efficient legendary-supported mana acceleration.
- 1 Ornithopter of Paradise | ramp | Creature-based fixing that also contributes to board development.
- 1 Ori, Keeper of Songs | synergy | A thematic Dwarf payoff for the artifact-centered plan.
- 1 Patchwork Banner | ramp | Mana acceleration that also strengthens the chosen creature type.
- 1 Promise of Loyalty | wipe | Selective board control that preserves each player's favored creature.
- 1 Psychosis Crawler | threat | A scalable artifact creature that pressures opponents through card draw.
- 1 Realm-Cloaked Giant // Cast Off | threat | A large follow-up threat paired with a creature sweep.
- 1 Reprieve | interaction | Efficiently delays a spell while replacing itself.
- 1 Selfless Spirit | interaction | Protects the creature board from destruction-based sweepers.
- 1 Serra Redeemer | threat | Provides flying pressure and rewards deploying smaller creatures.
- 1 Skullclamp | draw | Exceptional card flow with token creatures and expendable bodies.
- 1 Skyclave Apparition | removal | Broad permanent removal on a creature body.
- 1 Sol Ring | ramp | Highly efficient artifact mana acceleration.
- 1 Spirited Companion | draw | A low-cost body that immediately draws a card.
- 1 Sunscorch Regent | threat | A growing evasive finisher that gains life while attacking.
- 1 Sun Titan | threat | A substantial threat that repeatedly returns useful permanents.
- 1 Swiftfoot Boots | interaction | Gives important creatures haste and reliable protection.
- 1 Swords to Plowshares | removal | Premium low-cost creature exile.
- 1 Tangle Tumbler | synergy | Artifact-focused support that advances the deck's resource engine.
- 1 The Black Arrow | removal | Equipment-based removal that remains useful on the battlefield.
- 1 Three Tree Mascot | synergy | An artifact creature that supports the deck's typal and artifact themes.
- 1 Warren Warleader | threat | Creates a substantial attacking board and serves as a combat finisher.
- 1 Castle Ardenvale | land | Mana source with a late-game token-making option.
- 1 Command Tower | land | Reliable color fixing for the commander's identity.
- 1 Evolving Wilds | land | Fetches a basic land while improving land access.
- 1 Exotic Orchard | land | Usually provides the colors needed in multiplayer games.
- 1 Fabled Passage | land | Fetches a basic land and improves mana consistency.
- 1 Fountainport | land | Utility land that supports artifact-token resource production.
- 1 Hobbit Hole | land | Thematic mana source with useful Food production.
- 1 Minas Tirith | land | Thematic white source with a useful card-advantage mode.
- 1 Path of Ancestry | land | Fixes mana and rewards the deck's shared creature types.
- 23 Plains | land | Stable basic mana base for a mono-white deck.
- 1 Reliquary Tower | land | Utility land that preserves a full grip of cards.
- 1 Rogue's Passage | land | Makes a key equipped attacker difficult to block.
- 1 Swarmyard | land | Regenerates several relevant small creature types.
- 1 Three Tree City | land | Typal land that can generate a large burst of mana.
- 1 Thriving Heath | land | Fixes mana while retaining a basic-land style role.
- 1 Treasure Vault | land | Artifact land that converts spare mana into Treasure resources.
- 1 Uncharted Haven | land | Color-fixing land for consistent white access.

</details>

### 24. a 60-card deck from one set

Format: Modern. Theme: aggro. Pool: any_card. Shortlist: 96 names.

Cards: 60 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $47.58 to buy, $47.58 the whole deck.

**Summary:** This deck aims to establish pressure early with a dense red creature suite, using Emberheart Challenger and Hearthborn Battler to reinforce that attack plan. Removal clears opposing obstacles while the draw cards help keep threats coming, and Dragonhawk, Fate's Tempest gives the deck a heavier way to close games. It gives up broad flexibility and a deep late-game plan in favor of direct, creature-led aggression.

The quality model grades this deck below the precon baseline against the top lists of the format: the deck runs its spells as playsets, and that lowers the grade. Many of the cards appear in no top list, and that lowers the grade. The cards pair in ways the top lists do not, and that lowers the grade.

- [INFO] `curve_summary`: average mana value 3.11 over 36 nonland cards

<details><summary>The deck list</summary>

- 24 Mountain | land | Provides a consistent mono-red mana base.
- 2 Arcane Signet | ramp | Supplies the deck's available mana acceleration.
- 4 Artist's Talent | draw | Fills a draw slot while staying in the deck's red plan.
- 2 Sazacap's Brew | draw | Adds more card flow for an aggressive hand.
- 4 Abrade | removal | Provides efficient removal coverage.
- 2 Agate Assault | removal | Completes the removal package with a red spell.
- 4 Emberheart Challenger | synergy | Forms a core synergy creature for the aggressive plan.
- 4 Hearthborn Battler | synergy | Adds another full playset of synergy-focused pressure.
- 4 Frilled Sparkshooter | threat | Provides a full playset of aggressive threats.
- 4 Reptilian Recruiter | threat | Adds reliable creature pressure.
- 4 Teapot Slinger | threat | Supplies another set of threats for combat-focused games.
- 2 Dragonhawk, Fate's Tempest | threat | Rounds out the threat suite with impactful top-end creatures.

</details>

