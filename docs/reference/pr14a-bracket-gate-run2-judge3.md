# PR-14A bracket gate, judge lane

Run date: 2026-09-21. Card snapshot: 2026-09-04. Decks read from `/Users/nate/repos/decktome/docs/reference/pr14a-bracket-gate-run2.md`.

Verdict: FAIL. The judge agreed with the bracket on 4 of 15 decks (26 percent, the bar is 80), with 0 judge errors. This document reads the judge bar alone: the block, band, and content bars are in the source document.

Calls 15. Cost $0.2446. Time 109 seconds.

## Run

- Suite `bracket-judge`, run `pr14a-bracket-gate-run2-judge3`, on 2026-09-21, commit `23b67ce`.
- Roles: judge on `claude-opus-5` (anthropic, effort medium).
- Versions: card snapshot 2026-09-04, bracket_judge prompt version 2, generate prompt version 15, source `pr14a-bracket-gate-run2.md`.
- Calls: 15. Cost: $0.2446. Time: 109 seconds.

| # | Built for | Judged | Agrees | Commander |
|---|---|---|---|---|
| 1 | 1 | 3 | no | Gishath, Sun's Avatar |
| 2 | 1 | 2 | no | Karlov of the Ghost Council |
| 3 | 1 | 3 | no | Adeline, Resplendent Cathar |
| 4 | 2 | 3 | no | Gishath, Sun's Avatar |
| 5 | 2 | 3 | no | Karlov of the Ghost Council |
| 6 | 2 | 3 | no | Adeline, Resplendent Cathar |
| 7 | 3 | 3 | yes | Karlov of the Ghost Council |
| 8 | 3 | 3 | yes | Denethor, Ruling Steward |
| 9 | 3 | 3 | yes | Zada, Hedron Grinder |
| 10 | 4 | 3 | no | Urza, Lord High Artificer |
| 11 | 4 | 3 | no | Korvold, Fae-Cursed King |
| 12 | 4 | 3 | no | Prosper, Tome-Bound |
| 13 | 5 | 5 | yes | Kinnan, Bonder Prodigy |
| 14 | 5 | 4 | no | Yuriko, the Tiger's Shadow |
| 15 | 5 | 3 | no | Najeela, the Blade-Blossom |

## Why

### 1. Bracket 1, Gishath, Sun's Avatar, dinosaurs

Judge: bracket 3. This is a straightforward Dinosaur tribal beatdown deck with no marked Game Changers, no mass land denial, no extra turns, and Spellbook finds no combos, so it can't be a 4 or 5. However, the mana base is well beyond precon quality (Plateau, Spire Garden, Bountiful Promenade, Temple Garden, plus efficient ramp like Nature's Lore/Three Visits/Farseek) and it packs strong card-advantage payoffs and haste enablers to land Gishath early, which is clearly upgraded past Core. That puts it comfortably in Bracket 3, where games tend to be decided around turns six and up.

Combos the judge read: none, Commander Spellbook finds no combo.

### 2. Bracket 1, Karlov of the Ghost Council, lifegain

Judge: bracket 2. This is a straightforward lifegain-payoff deck with no marked Game Changers, no tutors, no fast mana, no mass land denial or extra turns, and no combos found by the spellbook check. The mana base is 28 basics plus a handful of duals and pain lands, and the wincons are incremental (Vito, Defiant Bloodlord, big lifelinked angels) rather than explosive, so games will regularly run eight-plus turns. It is a tidy, focused build but sits right around precon power, placing it in Core.

Combos the judge read: none, Commander Spellbook finds no combo.

### 3. Bracket 1, Adeline, Resplendent Cathar, tokens

Judge: bracket 3. No marked Game Changers, no combos, and no mass land denial, but the card quality is clearly beyond a precon \u2014 Hero of Bladehold, Mondrak, Divine Visitation, Ocelot Pride, Felidar Retreat and a deep stack of token/draw engines with efficient removal. The basic-heavy, ramp-light mana base and lack of tutors or fast mana mean it needs a few turns to snowball, so games land in the six-to-eight-turn range rather than ending early. That puts it at the lower end of Upgraded.

Combos the judge read: none, Commander Spellbook finds no combo.

### 4. Bracket 2, Gishath, Sun's Avatar, dinosaurs

Judge: bracket 3. This is a straightforward Dinosaur typal beatdown deck with no Game Changers, no mass land denial, no extra turns and no combos found, so nothing pushes it into Bracket 4. However, the mana base (Taiga, Plateau, shocks, Cavern of Souls), Sol Ring plus multiple mana dorks and ramp spells, and a suite of efficient payoffs and protection (Heroic Intervention, Akroma's Will, greaves/boots) make it clearly stronger and more consistent than a precon. That places it comfortably in Bracket 3, with games typically resolving via combat around turns six to eight.

Combos the judge read: none, Commander Spellbook finds no combo.

### 5. Bracket 2, Karlov of the Ghost Council, lifegain

Judge: bracket 3. This is a focused lifegain/Karlov deck clearly built past precon level, with premium cards like Solitude, Archangel of Thune, Heliod, Umezawa's Jitte and Aetherflux Reservoir, plus real interaction and three board wipes, but its mana base is heavily basic-land and it lacks fast mana or tutors. No cards are flagged as Game Changers, there is no mass land denial or extra turns, and the combo check found nothing, so it isn't Bracket 4. Its power and speed sit above a Core-level deck, making Upgraded the right fit.

Combos the judge read: none, Commander Spellbook finds no combo.

### 6. Bracket 2, Adeline, Resplendent Cathar, tokens

Judge: bracket 3. A focused mono-white token/go-wide deck with no marked Game Changers, no mass land denial, no extra turns, and no combos found — so it sits in the no-restriction-violations zone. However, the card quality and mana base (Ocelot Pride, Retrofitter Foundry, Sanctuary Warden, Hero of Bladehold, Cavern of Souls, Nykthos) plus a deep token-payoff and removal suite clearly exceed precon strength, with Adeline pressuring life totals fast. That puts it comfortably at Upgraded rather than Core, though it lacks the fast mana and tutors of an Optimized build.

Combos the judge read: none, Commander Spellbook finds no combo.

### 7. Bracket 3, Karlov of the Ghost Council, lifegain

Judge: bracket 3. This is a lifegain/aristocrat value deck with no marked Game Changers, no mass land denial, no extra turns, and Spellbook finds no combos, but it is clearly upgraded past precon level with premium cards like Solitude, Vein Ripper, Archangel of Thune, Aetherflux Reservoir and strong lifegain payoff engines. The mana base is basic-heavy with no fast mana or tutors, so it isn't fast enough for Bracket 4, yet its density of high-end payoffs and engine pieces (Heliod, Vito, Well of Lost Dreams) puts it above Core. Bracket 3 fits: grindy but capable of closing games around turn six to eight.

Combos the judge read: none, Commander Spellbook finds no combo.

### 8. Bracket 3, Denethor, Ruling Steward, aristocrats sacrifice

Judge: bracket 3. This is a well-tuned aristocrats deck with strong sac outlets, free altars (Ashnod's/Phyrexian), and payoffs like Blood Artist and Yawgmoth, but the only listed combo is a three-card loop, so it dodges the two-card infinite restriction. No cards are flagged as Game Changers and there's no mass land denial or extra turns, while the mana base and tutoring are only moderately optimized. That places it as an upgraded, clearly above-precon deck aiming to win around turns six to eight.

Combos the judge read: Chthonian Nightmare + Pitiless Plunderer + Mondrak, Glory Dominus.

### 9. Bracket 3, Zada, Hedron Grinder, goblins

Judge: bracket 3. This is an upgraded goblin-tribal Zada deck with two Game Changers (Jeska's Will, The One Ring), light fast mana (Sol Ring, rituals) and solid card draw, but it wins through go-wide combat damage rather than a tight combo line. The only two-card infinite (Kiki-Jiki + Akki Battle Squad) costs 5 and 6 mana, so it isn't a cheap early-turn kill, and the three-card Krenko loops are likewise slow. That puts it squarely within Bracket 3's limits of three or fewer Game Changers, no mass land denial, and no cheap early two-card combo.

Combos the judge read: Krenko, Mob Boss + Skirk Prospector + Goblin Warchief; Krenko, Mob Boss + Skirk Prospector + Goblin Chieftain; Kiki-Jiki, Mirror Breaker + Akki Battle Squad.

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Judge: bracket 3. Only one marked Game Changer (Mishra's Workshop) and Spellbook finds no combos, so it fits inside Bracket 3's limits despite Urza's inherent power. The fast mana (Mox Opal, Metalworker, Workshop, Urza's Saga) and Power Artifact/Clock of Omens engines push it well past precon speed, but there are no tutors, no mass land denial, and plenty of low-impact filler (Cargo Ship, Luck Bobblehead, Strixhaven Stadium) that keeps it from optimized. Interaction is decent counterspells and removal, giving a typical upgraded midrange clock around turn six or later.

Combos the judge read: none, Commander Spellbook finds no combo.

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Judge: bracket 3. This is a well-tuned Korvold aristocrats/sacrifice build with a strong dual-land mana base, treasure ramp, and tutors (Diabolic Intent, Savage Order, Dimir House Guard, Wight of the Reliquary), clearly beyond precon power. No cards are marked as Game Changers and Commander Spellbook finds no combos, and there's no mass land denial or extra turns, so it stays under the Bracket 4 threshold. Its grindy value engines (Yawgmoth, Blasting Station, Meathook Massacre) win in the mid-game rather than turn four, fitting Upgraded.

Combos the judge read: none, Commander Spellbook finds no combo.

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Judge: bracket 3. This is a well-tuned treasure/Prosper deck with strong dual lands, efficient interaction and card draw, but only two Game Changers (Orcish Bowmasters, The One Ring) and no mass land denial or extra turns. The only combo found is a clunky four-card engine, not a cheap two-card infinite. That profile — clearly beyond precon power yet within the three Game Changer cap — lands squarely in Upgraded.

Combos the judge read: Prosper, Tome-Bound + Professional Face-Breaker + Sensei's Divining Top + Xorn.

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Judge: bracket 5. This is a classic cEDH Kinnan combo deck: fast mana (Mana Vault, Grim Monolith, Basalt Monolith), fetch/dual mana base, a dense free/cheap interaction suite (Flusterstorm, Muddle the Mixture, Narset's Reversal, Cyclonic Rift), and multiple two-card infinite mana lines with Thassa's Oracle as the wincon. Kinnan + Basalt Monolith and Kelpie Guide + Freed from the Real can win as early as turn 2-3 with untapper/clone redundancy. The build is metagame-optimized with no theme concessions, placing it firmly at cEDH rather than Bracket 4.

Combos the judge read: Kelpie Guide + Freed from the Real; Kinnan, Bonder Prodigy + Grim Monolith + Spark Double; Kinnan, Bonder Prodigy + Basalt Monolith.

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Judge: bracket 4. This is a highly optimized Yuriko list with heavy fast mana (Sol Ring, Mana Vault, Chrome Mox, Mox Diamond/Opal/Amber, Lotus Petal, Dark Ritual), duals including Underground Sea, and a dense free-interaction suite (Force of Negation, Pact of Negation, Fierce Guardianship, Deadly Rollick, Snuff Out, Mental Misstep). It carries five Game Changers, well past Bracket 3's limit, and Yuriko's ninjutsu clock can close games by turn four or five. It lacks a dedicated cEDH combo win line and plays a tribal ninja theme, so it sits at Optimized rather than 5.

Combos the judge read: none, Commander Spellbook finds no combo.

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Judge: bracket 3. No marked Game Changers, no mass land denial, and no fast mana or tutor package f02 the deck is mostly a wide warrior-tribal battlecruiser with shocklands and a lot of small creatures. However, it packs the Najeela + Professional Face-Breaker infinite-combat combo, which at seven total mana with the commander is realistically assembled around turn six or later rather than as an early kill. That mix of a functional two-card win alongside otherwise casual, unoptimized cards lands it squarely in Upgraded.

Combos the judge read: Najeela, the Blade-Blossom + Professional Face-Breaker.

