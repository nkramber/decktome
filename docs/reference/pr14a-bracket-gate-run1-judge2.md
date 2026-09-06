# PR-14A bracket gate, judge lane

Run date: 2026-09-06. Card snapshot: 2026-09-04. Decks read from `/Users/nate/Repos/mtg-deck-builder/docs/reference/pr14a-bracket-gate-run1.md`.

Verdict: FAIL. The judge agreed with the bracket on 8 of 15 decks (53 percent, the bar is 80), with 0 judge errors. This document reads the judge bar alone: the block, band, and content bars are in the source document.

Calls 15. Cost $0.2787. Time 141 seconds.

## Run

- Suite `bracket-judge`, run `pr14a-bracket-gate-run1-judge2`, on 2026-09-06, commit `efa629a`.
- Roles: judge on `claude-opus-5` (anthropic, effort medium).
- Versions: card snapshot 2026-09-04, generate prompt version 12, source `pr14a-bracket-gate-run1.md`.
- Calls: 15. Cost: $0.2787. Time: 141 seconds.

| # | Built for | Judged | Agrees | Commander |
|---|---|---|---|---|
| 1 | 1 | 2 | no | Gishath, Sun's Avatar |
| 2 | 1 | 3 | no | Karlov of the Ghost Council |
| 3 | 1 | 3 | no | Adeline, Resplendent Cathar |
| 4 | 2 | 2 | yes | Gishath, Sun's Avatar |
| 5 | 2 | 3 | no | Karlov of the Ghost Council |
| 6 | 2 | 2 | yes | Adeline, Resplendent Cathar |
| 7 | 3 | 3 | yes | Karlov of the Ghost Council |
| 8 | 3 | 3 | yes | Denethor, Ruling Steward |
| 9 | 3 | 3 | yes | Zada, Hedron Grinder |
| 10 | 4 | 4 | yes | Urza, Lord High Artificer |
| 11 | 4 | 3 | no | Korvold, Fae-Cursed King |
| 12 | 4 | 4 | yes | Prosper, Tome-Bound |
| 13 | 5 | 5 | yes | Kinnan, Bonder Prodigy |
| 14 | 5 | 4 | no | Yuriko, the Tiger's Shadow |
| 15 | 5 | 3 | no | Najeela, the Blade-Blossom |

## Why

### 1. Bracket 1, Gishath, Sun's Avatar, dinosaurs

Judge: bracket 2. This is a straightforward dinosaur tribal battlecruiser deck with no Game Changers, no tutors, no fast mana beyond basic ramp spells, and no infinite combos. Its wins come from stacking big creatures and connecting with Gishath, which is a slow, fair plan around turn eight or later. The mana base is decent but the card quality and interaction suite sit right around a strong precon, so Bracket 2 with only a slight lean toward 3.

### 2. Bracket 1, Karlov of the Ghost Council, lifegain

Judge: bracket 3. This is a lifegain midrange pile with a mostly basic-land mana base and no fast mana, but it packs one Game Changer (Solitude) and a genuine two-card infinite (Heliod, Sun-Crowned + Archangel of Thune, with Vito/Cliffhaven Vampire as win outlets). At nine total mana that combo is not a realistic turn-six kill, and the deck's interaction is fair removal rather than stax or fast wins, so it sits above precon power without reaching Optimized. Bracket 3, Upgraded, is the right home.

### 3. Bracket 1, Adeline, Resplendent Cathar, tokens

Judge: bracket 3. This is a tightly-tuned mono-white go-wide token build \u2014 Adeline plus Hero of Bladehold, Mondrak, Ojer Taq, Divine Visitation and a deep suite of anthem/doubler payoffs is well beyond precon consistency and can close a game in the six-to-eight turn range. However, it runs no Game Changers, no fast mana or tutor package, no mass land denial or extra turns, and no two-card infinite combos (Halo Fountain needs a wide board rather than a compact loop). Its plain-heavy, ramp-light mana base and reliance on creature-based value rather than efficient interaction keep it clearly short of Bracket 4 optimization.

### 4. Bracket 2, Gishath, Sun's Avatar, dinosaurs

Judge: bracket 2. This is a straightforward Dinosaur tribal deck built around Gishath's combat trigger, with no Game Changers, no tutors, no fast mana beyond basic ramp spells, and no infinite combos. The mana base is decent but budget-ish (few true duals, no fetches), and the game plan of casting a seven-mana commander and connecting is slow. Interaction is limited to a couple of board wipes and Savage Stomp, putting it right around precon-plus power.

### 5. Bracket 2, Karlov of the Ghost Council, lifegain

Judge: bracket 3. This is a battlecruiser lifegain/Karlov value deck with a slow mana base (no fast mana, no real tutors, 27 basics), so its clock is more like turn 8-10, but it includes Game Changers (Vein Ripper, plus premium cards like Solitude and Aetherflux Reservoir) that push it out of Bracket 2. There's no compact two-card infinite here — Heliod and Archangel of Thune lack their usual partners — and interaction is limited to a few wraths and spot removal. That combination of one Game Changer and no combo lands it squarely as an upgraded casual deck.

### 6. Bracket 2, Adeline, Resplendent Cathar, tokens

Judge: bracket 2. This is a mono-white token deck with 37 basic Plains, no fast mana, no tutors, and no combos \u2014 the mana base and card quality sit right around precon level. Its power comes from slow value engines and token payoffs that need several turns to snowball, with only a handful of one-for-one removal spells. It has essentially no Game Changers, so it lands comfortably in Core, maybe brushing the low end of Upgraded.

### 7. Bracket 3, Karlov of the Ghost Council, lifegain

Judge: bracket 3. This is a well-tuned lifegain/Karlov deck with one Game Changer (Tymna the Weaver), solid interaction (Solitude, Murderous Rider, three wraths) and strong payoffs, but its mana base is casual (22 basics, few fast mana, no real tutor package) and it wins through incremental drain and Aetherflux Reservoir rather than a cheap two-card infinite. Nothing here reliably kills before turn six, yet it clearly outpaces a precon.

### 8. Bracket 3, Denethor, Ruling Steward, aristocrats sacrifice

Judge: bracket 3. This is a well-tuned Orzhov aristocrats deck that goes beyond precon power with free sac outlets (Carrion Feeder, Viscera Seer, Phyrexian/Ashnod's Altar), a thick drain package, and one Game Changer in Yawgmoth, Thran Physician, but it lacks fast mana, tutors, and a compact two-card kill in the early turns. Its wins come from grinding incremental drains and board control (Grave Pact, Meathook Massacre, Toxic Deluge), which typically lands around turn six or later. That profile fits Upgraded rather than Optimized.

### 9. Bracket 3, Zada, Hedron Grinder, goblins

Judge: bracket 3. The deck runs exactly three Game Changers (Wheel of Fortune, Jeska's Will, The One Ring) alongside Sol Ring and a fairly optimized goblin tribal shell, which sits right at the Upgraded ceiling. Its Conspicuous Snoop + Kiki-Jiki infinite is an eight-mana, fragile two-card line that rarely assembles before turn six, and there's no mass land denial or extra turns. Speed and interaction (Pyroblast, Cratermaker, Chain Reaction) are a clear step above precon but well short of Optimized.

### 10. Bracket 4, Urza, Lord High Artificer, artifacts

Judge: bracket 4. Between Urza himself, Mishra's Workshop, Gemstone Caverns and Mox Opal-style fast mana, the deck exceeds the three Game Changer cap and can produce huge explosive turns very early. Krark-Clan Ironworks plus Urza's construct/tap ability and the Ghostly Flicker + Mystic Sanctuary loop give it cheap engine combos well inside six turns. It's still a themed artifact build with clunky filler and only modest interaction rather than a lean cEDH list, so Bracket 4 rather than 5.

### 11. Bracket 4, Korvold, Fae-Cursed King, treasure sacrifice

Judge: bracket 3. This is a well-tuned Korvold aristocrats pile but it lacks the fast mana (no Crypt/Vault/Sol Ring-tier acceleration), tutor density, and free interaction of an Optimized list 
instead relying on grindy sac value engines and a basic-heavy, fetchless mana base. It runs only about two Game Changers (Yawgmoth, Thran Physician and Grist, the Hunger Tide), within Bracket 3's allowance, and while pieces like Yawgmoth, Nightmare Shepherd, Pitiless Plunderer, Warren Soultrader and Blasting Station can assemble loops, they're multi-card, mana-hungry setups unlikely to land in the first six turns. Games plausibly end around turn six to eight, which is squarely Upgraded.

### 12. Bracket 4, Prosper, Tome-Bound, treasure

Judge: bracket 4. The deck stacks up multiple Game Changers 	a Ragavan, Nimble Pilferer, Wheel of Fortune, The One Ring, and Ancient Copper Dragon 	a which is beyond the three allowed in Bracket 3, and it backs them with an expensive optimized mana base (Badlands, Blood Crypt, Cavern of Souls, filters and fastlands), premium cheap interaction (Deadly Rollick, Snuff Out, Terminate, Chaos Warp, Pyroblast) and card advantage engines like Black Market Connections and Skullclamp. Treasure-fueled Prosper drains with Torment of Hailfire or Revel in Riches can close games well before turn six even without a dedicated two-card combo. That speed, redundancy, and Game Changer density places it squarely in Optimized rather than Upgraded, though it lacks the tutor-and-combo focus of cEDH.

### 13. Bracket 5, Kinnan, Bonder Prodigy, combo

Judge: bracket 5. This is a Kinnan cEDH combo deck: fast mana (Mana Vault, Grim Monolith, Basalt Monolith, Gemstone Caverns), an optimized fetch/dual mana base, and multiple two-card infinite mana lines (Kinnan + Basalt Monolith, Devoted Druid, Voltaic Construct/Filigree Sages/Staff of Domination) that win via Thassa's Oracle. It is packed with free/cheap interaction and untap-enablers plus Game Changers like Cyclonic Rift and Thassa's Oracle, aiming at turn 2-4 wins with protection rather than any theme.

### 14. Bracket 5, Yuriko, the Tiger's Shadow, ninjas

Judge: bracket 4. The mana base and acceleration are cEDH-grade erry\u2014 Underground Sea, Chrome Mox, Mox Diamond, Mox Opal, Mana Vault, Lotus Petal, Dark Ritual erry\u2014 backed by a full free-interaction suite (Force of Will, Force of Negation, Fierce Guardianship, Pact of Negation, Deadly Rollick) and well over three Game Changers, so it's far past Bracket 3. It can realistically start flipping Yuriko triggers turn two and kill from the air by turn four or five. It stops short of Bracket 5 because it's still built as a ninja-tribal theme erry\u2014 clunky payoffs like Okiba-Gang Shinobi, Ink-Eyes, and Fallen Shinobi over the dedicated one-mana unblockable creatures and tutor/combo package a true cEDH Yuriko list runs.

### 15. Bracket 5, Najeela, the Blade-Blossom, warriors combat

Judge: bracket 3. The deck plays two Game Changers (Najeela in the command zone plus Winota, Joiner of Forces) and packs the classic Najeela + Combat Celebrant / Lightning Runner infinite-combat lines, which pushes it past Core. However, there's no fast mana, no tutor package, and the creature suite is a pile of midrange warriors and battle-cry filler, so the combo is expensive and unlikely before turn six. That profile — precon-plus power, a few Game Changers, a slow two-card kill — sits squarely in Upgraded.

