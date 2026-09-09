# PR-8 deck gate

Run date: 2026-09-09. Card snapshot: 2026-09-03.

Verdict: PASS. 2 of 2 decks passed every block check, 0 invented names reached the user, and 0 summaries stated a false rule of the game. All three bars are zero tolerance.

## Summary

| Measure | Value |
|---|---|
| Prompts | 2 |
| Decks returned | 2 |
| Decks with no block finding | 2 |
| Invented names that reached the user | 0 |
| Decks that needed the repair turn | 1 |
| Summaries judged (F-26) | 2 |
| Summaries that state a rule of the game | 0 |
| Summaries that state a FALSE rule | 0 |
| Judge errors | 0 |
| Decks the plan judge read (PR-15, information) | 2 |
| Mean plan score, 0 to 1 | 0.94 |
| Plan reasons the judge left empty | 0 |
| Errors | 0 |
| Prompt version | 12 |
| Calls | 7 |
| Cost | $0.3209 |
| Time | 267 seconds |

## Run

- Suite `decks`, run `pr8-deck-gate-run19-upgrades`, on 2026-09-09, commit `b361ce4`.
- Partial run over `17,18`. It reads the bars of its own items, and it never stands for the suite.
- Roles: generate on `gpt-5.6-terra` (openai, effort medium), judge on `claude-opus-5` (anthropic, effort medium), repair on `gpt-5.6-terra` (openai, effort medium).
- Versions: card snapshot 2026-09-03, generate prompt version 12, plan_rubric prompt version 2.
- Calls: 7. Cost: $0.3209. Time: 267 seconds.

## Findings by code

A block stops the deck. A warning and a note are reports.

| Code | Count |
|---|---|
| `profile_off_band` | 7 |
| `curve_summary` | 2 |
| `two_card_combo` | 1 |
| `basics_added` | 1 |

By severity: BLOCK 0. WARN 8. INFO 3. 

## Decks

### 17. upgrade a precon, owned first

Format: Commander. Theme: goblins. Pool: owned_first. Shortlist: 196 names.

Commander: Zada, Hedron Grinder.

Cards: 99 main, 0 sideboard. Repair turn: yes, for deck_size, precon_share, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band, profile_off_band. Block findings: 0.

Cost: $84.58 to buy, $288.25 the whole deck.

**Summary:** A mono-red Goblin swarm deck that develops a broad battlefield, converts spellcasting into explosive mana and card flow, and leverages mass-targeting tricks for dramatic combat turns. Token creation, sacrifice damage, haste effects, and tribal payoffs give it several overlapping ways to pressure the table while maintaining resilient late-game reach.

- PLAN plan_coherent=partly: The go-wide goblin core, Zada mass-targeting cantrips, and ritual/Grapeshot storm elements mostly interlock (tokens as Zada targets, Skirk Prospector and Storm-Kiln Artist fueling spells), but the storm payoff count is thin and Blasphemous Act actively sweeps the deck's own swarm, so the pieces pull slightly apart.
- PLAN theme_fit=yes: It is a mono-red goblin deck with a genuine storm subtheme (rituals, Mana Geyser, Past in Flames, Grapeshot, Empty the Warrens) at a reasonable bracket-3 power level, matching the precon upgrade request even though it leans on Zada as the helm.
- PLAN useful_as_built=yes: 37 lands plus heavy cheap ramp support a very low curve, and there are multiple real win routes (Impact Tremors/Bombardment drain, Grapeshot, pumped alpha strikes with Fists of Flame/Haze of Rage).
- PLAN summary_honest=yes: Every claim — token makers, sacrifice damage (Goblin Bombardment, Siege-Gang, Boggart Shenanigans), haste enablers, tribal lords, mass-targeting tricks, and spell-into-mana/cards effects — is backed by actual cards in the list.
- [INFO] `curve_summary`: average mana value 2.62 over 63 nonland cards
- [WARN] `profile_off_band`: the count of nonbasic lands that make only colorless mana is 5, and bracket 3 wants 4 at most (Fountainport, Goblin Burrows, Kher Keep, Reliquary Tower, War Room)
- [WARN] `two_card_combo`: bracket 3 allows no two-card infinite combo that needs four mana or less, and the deck holds Storm-Kiln Artist + Haze of Rage (speed 5, near two-card)

<details><summary>The deck list</summary>

- 1 Ancestors' Aid | synergy | A low-cost targeted spell that supports Zada-style mass targeting turns.
- 1 Battle Hymn | ramp | Converts a wide Goblin board into a burst of mana.
- 1 Blasphemous Act | wipe | Efficient emergency creature-board reset.
- 1 Boggart Shenanigans | removal | Turns Goblin deaths into reliable damage.
- 1 Brightstone Ritual | ramp | Explosive ritual mana for Goblin-heavy turns.
- 1 Chaos Warp | removal | Versatile answer to problematic permanents.
- 1 Conspicuous Snoop | synergy | Goblin-focused value and board development.
- 1 Crimson Wisps | draw | Cantrip that grants haste and scales with Zada.
- 1 Daring Discovery | synergy | Supports the deck's proactive spell-based game plan.
- 1 Dragon Fodder | synergy | Produces multiple Goblins for go-wide pressure and Zada turns.
- 1 Empty the Warrens | synergy | Creates a large Goblin board after a spell-heavy turn.
- 1 Expedite | draw | Cheap haste cantrip that becomes powerful with Zada.
- 1 Fable of the Mirror-Breaker // Reflection of Kiki-Jiki | ramp | Provides Treasure mana, filtering, and a value-copy threat.
- 1 Faithless Looting | draw | Efficient card selection and graveyard setup.
- 1 Fists of Flame | draw | A Zada-scalable cantrip and combat finisher.
- 1 Frontline Heroism | synergy | Supports the deck's creature-forward combat plan.
- 1 Gempalm Incinerator | removal | Goblin-scaled creature removal with cycling utility.
- 1 Grapeshot | removal | Spell-chain payoff that can pick off creatures or finish opponents.
- 1 Haze of Rage | synergy | Repeatable storm-style pump effect for a wide board.
- 1 Idol of Oblivion | draw | Steady card advantage alongside frequent token creation.
- 1 Impact Tremors | synergy | Converts every Goblin and token burst into direct damage.
- 1 Impulsive Pilferer | ramp | Early Goblin that leaves behind Treasure mana.
- 1 Krenko's Command | synergy | Efficiently adds multiple Goblins to the battlefield.
- 1 Krenko, Mob Boss | threat | Major repeatable Goblin-token engine and battlefield threat.
- 1 Mana Geyser | ramp | Large ritual effect that enables explosive turns.
- 1 Mogg War Marshal | synergy | Multiple Goblin bodies from one card.
- 1 Past in Flames | synergy | Reuses a graveyard of cantrips, rituals, and finishing spells.
- 1 Quest for the Goblin Lord | synergy | Rewards building and attacking with a large Goblin force.
- 1 Roaming Throne | synergy | Doubles important Goblin triggered abilities.
- 1 Ruby Medallion | ramp | Red spell cost reduction supports multi-spell turns.
- 1 Sazacap's Brew | draw | Cheap card flow for a spell-dense deck.
- 1 Seething Song | ramp | Burst mana for deploying Zada or chaining spells.
- 1 Siege-Gang Commander | removal | Creates a board and turns spare Goblins into damage.
- 1 Siege-Gang Lieutenant | removal | Goblin token production with direct-damage utility.
- 1 Skirk Prospector | ramp | Goblin sacrifice outlet and efficient mana acceleration.
- 1 Skullclamp | draw | Powerful card-draw engine with disposable Goblin tokens.
- 1 Sol Ring | ramp | Efficient early acceleration.
- 1 Spreading Insurrection | threat | High-impact creature swing that can break stalled boards.
- 1 Storm-Kiln Artist | ramp | Turns instant and sorcery casting into Treasure mana.
- 1 Swiftfoot Boots | interaction | Protects Zada while granting haste.
- 1 Throne of Eldraine | ramp | Flexible mana piece with late-game utility.
- 1 Vandalblast | wipe | Artifact sweep that can clear opposing mana and engines.
- 1 Warren Torchmaster | synergy | Goblin combat payoff that maintains tribal pressure.
- 1 Wild Ride | synergy | Supports aggressive combat and the spell-based plan.
- 1 Witch's Mark | draw | Card filtering that can also push combat damage.
- 1 Goblin Chieftain | synergy | Anthem and haste support for the Goblin army.
- 1 Goblin Lackey | synergy | Efficient early Goblin deployment engine.
- 1 Goblin Matron | synergy | Finds the most useful Goblin for the current board state.
- 1 Goblin Negotiation | removal | Goblin-themed removal option.
- 1 Goblin Trashmaster | removal | Goblin anthem with artifact removal utility.
- 1 Goblin Warchief | synergy | Reduces Goblin costs and gives the team haste.
- 1 Rundvelt Hordemaster | synergy | Anthem effect that supplies cards when Goblins die.
- 1 Searslicer Goblin | synergy | Adds another Goblin body and tribal payoff.
- 1 Arcane Signet | ramp | Reliable early mana acceleration.
- 1 Fellwar Stone | ramp | Additional inexpensive mana acceleration.
- 1 Abrade | removal | Flexible answer to creatures or artifacts.
- 1 Lightning Bolt | removal | Efficient instant-speed creature removal or reach.
- 1 Goblin Bombardment | removal | Sacrifice outlet that turns tokens into repeatable damage.
- 1 Ancestral Anger | draw | One-mana cantrip that becomes a wide-board pump spell with Zada.
- 1 Lightning Greaves | interaction | Fast protection and haste for Zada or key Goblins.
- 1 Champion's Helm | interaction | Protects the commander while improving combat.
- 1 Darksteel Plate | interaction | Durable protection for the commander or a key threat.
- 1 The Walls of Ba Sing Se | interaction | Flexible artifact-based defensive interaction.
- 1 Arena of Glory | land | Red-producing utility land that helps creatures act immediately.
- 1 Castle Embereth | land | Red source with a useful go-wide combat boost.
- 1 Den of the Bugbear | land | Red source that becomes a creature threat when needed.
- 1 Dwarven Mine | land | Mountain land that can provide an extra body.
- 1 Forgotten Cave | land | Red source with cycling utility.
- 1 Fountainport | land | Utility land for converting resources into value.
- 1 Goblin Burrows | land | Red source with a Goblin-focused combat ability.
- 1 Hidden Volcano | land | Red-producing land with utility.
- 1 Kher Keep | land | Utility land that produces small creatures for sacrifice and targeting synergies.
- 1 Reliquary Tower | land | Utility land that preserves a large refilled hand.
- 1 Shinka, the Bloodsoaked Keep | land | Untapped red source with combat utility.
- 1 Smoldering Crater | land | Red source with cycling utility.
- 1 Spinerock Knoll | land | Red utility land that can deploy a spell after a large attack.
- 1 War Room | land | Utility land that supplies card draw in longer games.
- 22 Mountain | land | Primary untapped red mana base for consistently casting red spells and the commander.

</details>

### 18. upgrade a precon, any card

Format: Commander. Theme: turtles. Pool: any_card. Shortlist: 300 names.

Commander: Heroes in a Half Shell.

Cards: 99 main, 0 sideboard. Repair turn: no. Block findings: 0.

Cost: $250.37 to buy, $250.37 the whole deck.

**Summary:** This Turtle Power upgrade keeps the precon’s character-heavy Turtle plan intact while smoothing its early mana and adding dependable ramp and card flow. Develop the board with the deck’s themed creatures and support pieces, then press the advantage with its larger Turtle threats while holding removal and sweepers for opposing boards. The deck gives up some of the precon’s slower color-fixing lands in exchange for faster access to its colors, but remains focused on the original theme rather than pursuing a compact combo finish.

- PLAN plan_coherent=yes: The list keeps a single coherent precon plan — deploy the TMNT character creatures and support pieces, back them with ramp, interaction and sweepers — and the additions (Corpsejack Menace, Vigor, hydras, Biogenic Ooze) all point the same direction rather than splitting into a second strategy.
- PLAN theme_fit=yes: It is a Heroes in a Half Shell Turtle Power precon upgrade that retains the themed cards while adding modest staples, matching the bracket 3 ask with no tutors, fast mana beyond Sol Ring, or two-card combos.
- PLAN useful_as_built=yes: 39 lands with Command Tower, Cavern of Souls, City of Brass, Mana Confluence, two shocklands, fetches and Chromatic Lantern give a workable if pain-heavy five-color base, and the deck has plenty of creature-based ways to close a game.
- PLAN summary_honest=yes: Everything claimed is present: added ramp (Sol Ring, Arcane Signet, Chromatic Lantern, Cultivate, Nature's Lore), card flow (Rhystic Study, Harmonize, April O'Neil), removal/sweepers (Assassin's Trophy, Blasphemous Act, Vanquish the Horde, Swift Demise), swapped tap-lands for painlands/shocklands, and no combo finish is claimed.
- [INFO] `curve_summary`: average mana value 3.42 over 60 nonland cards
- [WARN] `profile_off_band`: the land count is 39, and bracket 3 wants 34 to 38
- [WARN] `profile_off_band`: the ramp count is 6, and bracket 3 wants 8 to 13
- [WARN] `profile_off_band`: the draw count is 2, and bracket 3 wants 8 to 14
- [WARN] `profile_off_band`: the removal count is 2, and bracket 3 wants 6 to 12
- [WARN] `profile_off_band`: the wipe count is 1, and bracket 3 wants 2 to 5
- [WARN] `profile_off_band`: the interaction count is 0, and bracket 3 wants 4 to 12
- [INFO] `basics_added`: the list was 5 cards short, so the builder added 5 basic lands

<details><summary>The deck list</summary>

- 1 Acidic Slime | other | Flexible precon removal on a creature body.
- 1 April O'Neil, Live on the Scene | draw | Retained precon card for the deck's established theme.
- 1 Arcade Cabinet | other | Retained precon card for the deck's established theme.
- 1 Arcane Signet | ramp | Reliable multicolor acceleration from the precon.
- 1 Assassin's Trophy | removal | Versatile precon answer to troublesome permanents.
- 1 Baxter, Fly in the Ointment | other | Retained precon character supporting the established theme.
- 1 Bebop, Skull & Crossbones | other | Retained precon character supporting the established theme.
- 1 Big Mother Mouser | other | Retained precon card for the deck's established theme.
- 1 Biogenic Ooze | other | Retained precon creature for its existing board presence.
- 1 Blasphemous Act | wipe | Efficient precon reset when the board gets crowded.
- 1 Chromatic Lantern | ramp | Precon mana fixing and acceleration for the multicolor shell.
- 1 Coin of Mastery | other | Retained precon card for the deck's established theme.
- 1 Continue? | other | Retained precon card for the deck's established theme.
- 1 Corpsejack Menace | other | Retained precon creature supporting its existing counter theme.
- 1 Cultivate | ramp | Stable precon land ramp for a five-color mana base.
- 1 Dimension X Pizzasaur | other | Retained precon card for the deck's established theme.
- 1 Donatello, the Brains | synergy | Core Turtle-themed precon synergy piece.
- 1 Double Jump // Flying Kick | other | Retained precon card for the deck's established theme.
- 1 Electric Seaweed | other | Retained precon card for the deck's established theme.
- 1 Endless Foot Assault | other | Retained precon card for the deck's established theme.
- 1 Everything Pizza | other | Retained precon card for the deck's established theme.
- 1 Exploding Barrel | other | Retained precon card for the deck's established theme.
- 1 Fast Forward | other | Retained precon card for the deck's established theme.
- 1 Foot Chopper | other | Retained precon Equipment for the deck's theme.
- 1 Game Over | other | Retained precon card for the deck's established theme.
- 1 Harmonize | other | Precon card advantage that helps keep threats flowing.
- 1 Here Comes a New Hero! | other | Retained precon card for the deck's established theme.
- 1 High Score | other | Retained precon card for the deck's established theme.
- 1 Irma, Part-Time Mutant | other | Retained precon character supporting the established theme.
- 1 Krang, the All-Powerful | other | Retained precon character supporting the established theme.
- 1 Leatherhead, Iron Gator | other | Retained precon character supporting the established theme.
- 1 Leonardo, the Balance | threat | Retained precon Turtle threat.
- 1 Lessons from Life | other | Retained precon card for the deck's established theme.
- 1 Level Up | other | Retained precon card for the deck's established theme.
- 1 Michelangelo, the Heart | synergy | Core Turtle-themed precon synergy piece.
- 1 Mole Module | other | Retained precon card for the deck's established theme.
- 1 Mona Lisa, Science Geek | other | Retained precon character supporting the established theme.
- 1 Ninja Pizza | other | Retained precon card for the deck's established theme.
- 1 Raphael, the Muscle | threat | Retained precon Turtle threat.
- 1 Rat King, Pale Piper | other | Retained precon character supporting the established theme.
- 1 Ray Fillet, Wave Warrior | other | Retained precon character supporting the established theme.
- 1 Roadkill Rodney | other | Retained precon card for the deck's established theme.
- 1 Rocksteady, Mutant Marauder | other | Retained precon character supporting the established theme.
- 1 Shellshock | other | Retained precon card for the deck's established theme.
- 1 Shredder, Shadow Master | other | Retained precon character supporting the established theme.
- 1 Sol Ring | ramp | Efficient precon acceleration.
- 1 Special Move | other | Retained precon card for the deck's established theme.
- 1 Splinter, the Mentor | other | Retained precon character supporting the established theme.
- 1 Steelbane Hydra | removal | Retained precon creature that provides removal.
- 1 Super Combo | other | Retained precon card for the deck's established theme.
- 1 Swift Demise | other | Retained precon card for the deck's established theme.
- 1 Tempestra, Dame of Games | other | Retained precon character supporting the established theme.
- 1 Together Forever | other | Retained precon card for the deck's established theme.
- 1 Tokka & Rahzar, Unsupervised | ramp | Retained precon Turtle-themed ramp piece.
- 1 Vanquish the Horde | other | Retained precon board-reset option.
- 1 Vigor | other | Retained precon creature supporting the established theme.
- 1 Voracious Hydra | other | Retained precon creature for its existing board impact.
- 1 Wave Goodbye | other | Retained precon card for the deck's established theme.
- 1 Nature's Lore | ramp | Low-cost land ramp improves early development.
- 1 Rhystic Study | draw | Adds a strong ongoing source of cards.
- 1 Ash Barrens | land | Retained precon mana-fixing land.
- 1 Cinder Glade | land | Retained precon red-green source.
- 1 Command Tower | land | Reliable multicolor precon land.
- 1 Escape Tunnel | land | Retained precon utility land.
- 1 Evolving Wilds | land | Retained precon land that fixes basic colors.
- 1 Fabled Passage | land | Retained precon land that fixes basic colors.
- 4 Forest | land | Green basic sources support the deck's ramp and Turtle cards.
- 4 Island | land | Blue basic sources support the deck's Turtle cards.
- 3 Mountain | land | Red basic sources support the deck's Turtle cards.
- 1 Path of Ancestry | land | Retained precon tribal mana source.
- 4 Plains | land | White basic sources support the deck's Turtle cards.
- 1 Rain-Slicked Copse | land | Retained precon blue-green source.
- 1 Rootbound Crag | land | Retained precon red-green source.
- 1 Sodden Verdure | land | Retained precon blue-green source.
- 4 Swamp | land | Black basic sources support the deck's Turtle cards.
- 1 Thriving Moor | land | Retained precon flexible color source.
- 1 Turtle Lair | land | Retained precon Turtle-themed land.
- 1 Undergrowth Stadium | land | Retained precon black-green source.
- 1 Vernal Fen | land | Retained precon black-green source.
- 1 Vibrant Cityscape | land | Retained precon multicolor land.
- 1 City of Brass | land | Untapped multicolor fixing improves early color access.
- 1 Mana Confluence | land | Untapped multicolor fixing improves early color access.
- 1 Cavern of Souls | land | Tribal-focused color fixing for the Turtle plan.
- 1 Hallowed Fountain | land | Untapped-access white-blue source improves the mana base.
- 1 Watery Grave | land | Untapped-access blue-black source improves the mana base.

</details>

