# PR-14A bracket gate, judge lane

Run date: 2026-09-13. Card snapshot: 2026-09-04. Decks read from `/Users/nate/Repos/decktome/docs/reference/pr14a-bracket-gate-run2.md`.

Verdict: FAIL. The judge agreed with the bracket on 3 of 15 decks (20 percent, the bar is 80), with 0 judge errors. This document reads the judge bar alone: the block, band, and content bars are in the source document.

Calls 15. Cost $0.2634. Time 128 seconds.

## Run

- Suite `bracket-judge`, run `pr14a-bracket-gate-run2-judge2`, on 2026-09-13, commit `8a66086`.
- Roles: judge on `claude-opus-5` (anthropic, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 12, source `pr14a-bracket-gate-run2.md`.
- Calls: 15. Cost: $0.2634. Time: 128 seconds.

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
| 13 | 5 | 4 | no | Kinnan, Bonder Prodigy |
| 14 | 5 | 4 | no | Yuriko, the Tiger's Shadow |
| 15 | 5 | 3 | no | Najeela, the Blade-Blossom |

## Why

### 1. Bracket 1, Gishath, Sun's Avatar, dinosaurs

Judge: bracket 3. This is a focused Dinosaur tribal beatdown deck with a strong dual/shock/pain land mana base, quality ramp (Nature's Lore, Three Visits, Farseek) and solid card draw, but no marked Game Changers, no fast mana like Mana Crypt, and no infinite combos. Interaction is limited to creature-based removal (Trumpeting Carnosaur, Thrashing Brontodon, Apex Altisaur) and Heroic Intervention protection. It's clearly upgraded beyond precon power and can close games around turns six to eight, which fits Upgraded rather than Core or Optimized.

### 2. Bracket 1, Karlov of the Ghost Council, lifegain

Judge: bracket 2. This is a straightforward white-black lifegain pile with no marked Game Changers, no mass land denial, no extra turns, and no genuine two-card infinite (Heliod and the Vito-style drainers lack their combo partners like Walking Ballista or Exquisite Blood). The mana base is basics-heavy with no fast mana and essentially no tutoring, and the game plan of stacking lifegain triggers until Vito, Defiant Bloodlord, or a big Voice of the Blessed closes takes many turns. Interaction is thin (a handful of spot removal and one wrath), putting it right at upgraded-precon power.

### 3. Bracket 1, Adeline, Resplendent Cathar, tokens

Judge: bracket 3. This is a mono-white go-wide tokens deck built well beyond a precon \u2014 Mondrak, Ocelot Pride, Divine Visitation, Felidar Retreat and a deep token-doubling/anthem package with real card advantage engines \u2014 but the mana base is nearly all basics with no fast mana and no tutors. No Game Changers are marked, there is no mass land denial or extra turns, and the only infinite-ish loops (Animation Module/Retrofitter style value) are mana-hungry rather than cheap two-card kills. Consistent but grindy pressure that usually closes around turns six to eight puts it squarely in Upgraded.

### 4. Bracket 2, Gishath, Sun's Avatar, dinosaurs

Judge: bracket 3. This is a dinosaur typal beatdown deck with no marked Game Changers, no mass land denial, no extra turns and no two-card infinite combos, so it isn't Bracket 4/5 material. However, the dual/shock/pain land base with Taiga, Plateau, Cavern of Souls, plus Sol Ring, mana dorks, and efficient ramp and protection (Heroic Intervention, Akroma's Will, Greaves/Boots) put it clearly beyond precon power. It wants to land Gishath around turn five or six and snowball, which is textbook Upgraded.

### 5. Bracket 2, Karlov of the Ghost Council, lifegain

Judge: bracket 3. This is a focused lifegain payoff deck whose mana base is basic-land-heavy with no fast mana or tutors, and no cards are flagged as Game Changers, so it sits above precon power but well short of optimized. It does pack a genuine two-card infinite in Heliod, Sun-Crowned plus Archangel of Thune (with Aetherflux Reservoir and Vito as alternate kill buttons), which pushes it out of Bracket 2, though at five- and six-mana pieces with no acceleration it won't assemble early. Removal and wraths are present but the clock is a grindy turn-seven-plus affair, squarely Upgraded.

### 6. Bracket 2, Adeline, Resplendent Cathar, tokens

Judge: bracket 3. No cards are marked as Game Changers, and there's no mass land denial, extra turns, or two-card infinite combo, so nothing pushes it past the Bracket 3 ceiling. Still, the card quality and land base (Cavern of Souls, Nykthos, Plaza of Heroes, Ocelot Pride, Sanctuary Warden, Myrel) are clearly souped up beyond a precon, with a dense token-and-anthem engine plus real removal that can close games around turn six or seven. That grind-out-value-with-upgrades profile is textbook Upgraded rather than Core.

### 7. Bracket 3, Karlov of the Ghost Council, lifegain

Judge: bracket 3. This is a well-tuned lifegain/drain build that goes beyond precon power "" Heliod, Sun-Crowned plus Archangel of Thune is a genuine two-card infinite, with Aetherflux Reservoir, Vito and Enduring Tenacity as backup kill conditions, plus solid interaction (Solitude, Kaya's Wrath, Fumigate, Meathook Massacre). No cards are flagged as Game Changers and there is no mass land denial or extra turns, so it doesn't need Bracket 4. The basic-heavy mana base and lack of fast mana or tutors mean the combo won't assemble in the first six turns, which fits Upgraded squarely.

### 8. Bracket 3, Denethor, Ruling Steward, aristocrats sacrifice

Judge: bracket 3. This is a well-tuned aristocrats deck with a strong dual/painland mana base, cheap sac outlets, free altars, and multiple drain engines, clearly beyond precon power, but no cards are marked as Game Changers and there's no fast mana or heavy tutor suite. Its kill patterns (Blasting Station/Chthonian Nightmare/Yawgmoth loops with Blood Artist effects) mostly need three pieces or expensive setup, so they're unlikely to assemble cheaply in the first six turns. Grindy, interactive, and capable of winning around turn seven or eight, it sits squarely in Upgraded.

### 9. Bracket 3, Zada, Hedron Grinder, goblins

Judge: bracket 3. This is an upgraded goblin tribal build with two marked Game Changers (Jeska's Will, The One Ring), a fair amount of fast mana and card advantage, but no mass land denial, no extra turns, and no clean cheap two-card infinite (Kiki-Jiki lacks its untapper partner here). Its power level is clearly above a precon — efficient rituals, Skullclamp, free counter-magic in Pyroblast/REB — yet it lacks the tutor density, staples, and consistency of an optimized or cEDH list. Games realistically end around turn six to eight, squarely Bracket 3.

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Judge: bracket 3. This is a clearly upgraded artifact deck — Mishra's Workshop, Mox Opal, Metalworker and Urza himself generate explosive mana, and there are combo-adjacent pieces (Power Artifact, Clock of Omens/Unwinding Clock, Aetherflux Reservoir, Blasting Station) — but only one marked Game Changer and no reliable cheap two-card kill assembled early. The mana base is 28 basic Islands with almost no tutors or free interaction, so it lacks the consistency and speed of an optimized or cEDH build. Its power sits comfortably above precon level, landing squarely in Upgraded.

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Judge: bracket 3. This is a well-tuned Jund aristocrats/sacrifice build with a strong dual-heavy mana base, multiple sac outlets and drain payoffs, and targeted tutors (Diabolic Intent, Savage Order, Dimir House Guard), but no marked Game Changers, no mass land denial, and no cheap guaranteed two-card infinite kill in the first six turns. It's clearly souped up well beyond precon power with engines like Yawgmoth, Grist, Blasting Station and Academy Manufactor, yet it lacks the fast mana and consistent fast kills of an optimized list. That places it squarely in Upgraded, aiming to close games around turn six to eight.

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Judge: bracket 3. This is a well-tuned Prosper treasure deck with a premium mana base (Badlands, shocks, Cavern of Souls), deep interaction, and card advantage engines, but only two Game Changers (Orcish Bowmasters, The One Ring) and no Sol Ring-tier fast mana or dedicated two-card infinite kill. Its wins come from grindy treasure value and incremental drain/beatdown rather than a turn-four combo, so games should still take six-plus turns. That profile is squarely Upgraded rather than Optimized.

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Judge: bracket 4. This is a dedicated Kinnan untap-combo deck with fast mana (Mana Vault, Grim Monolith, Basalt Monolith), fetch-dual mana, and multiple two-card kills available in the first few turns \u2014 Kinnan plus Basalt Monolith/Freed from the Real/Dramatic Reversal/Pili-Pala into Thassa's Oracle \u2014 backed by free-ish protection and counterspells. That speed and redundancy blows well past Bracket 3's 'no cheap two-card infinite before turn six' limit. It falls short of true cEDH, though, since it lacks the density of premium tutors, Force-style free interaction and fast rocks, and leans on a themed pile of marginal untappers (Twitch, Apathy, Pip-Boy 3000, Sonic Screwdriver), so Bracket 4 is the right home."}

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Judge: bracket 4. This is a highly optimized Yuriko turbo-ninja list with heavy fast mana (Sol Ring, Mana Vault, Mox Diamond/Opal/Amber, Chrome Mox, Lotus Petal, Dark Ritual), premium duals including Underground Sea, and a dense free-interaction suite (Force of Negation, Pact of Negation, Fierce Guardianship, Deadly Rollick, Snuff Out, Mental Misstep). Five marked Game Changers already push it past Bracket 3, and the deck can realistically kill with Yuriko triggers around turns four to five. It stops short of cEDH only in that it's a tribal ninja theme rather than a pure best-strategy build.

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Judge: bracket 3. This is a warrior-tribal battle-cruiser pile with no marked Game Changers, no fast mana, and essentially no tutors, so it sits well below optimized play. However, it does run shock lands and a genuine two-card engine in Najeela plus Combat Celebrant (and repeatable extra combats via Raiyuu/Samut), which pushes it past a precon-level Core deck. The clunky mana and low-power creature suite keep it firmly in Upgraded rather than Bracket 4.

