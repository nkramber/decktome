# The review of the Gríma deck, as the owner wrote it, 2026-09-14

The owner wrote this review of deck `sFLbEUuKzI0QyLPft0zI` of session `z1hshyY6Npig1FN2NuV7` on 2026-09-14. The text below is word for word, and D-642 permits the words of a reader. `docs/reference/deck-review-grima-2026-09-14.md` checks each claim. Some claims are wrong, so read the checked document before you act on a claim.

## The review

```text
## Verdict

By the Commander Brackets rules it's a legal **Bracket 3** deck — one Game Changer (Orcish Bowmasters), no tutors, no two-card infinites, no extra turns, no land denial. But it *plays* like a high Bracket 2. It won't embarrass itself at a Bracket 3 table and it will get lapped at a Bracket 4 one, mostly for two reasons: the mana base is a budget precon mana base, and **the deck has no win condition**.

Structurally it's sound: 99 + commander, legal color identity, 33 lands, ~9 mana rocks, ~13 draw pieces, ~16 pieces of interaction. That's a real curve built by something that understands deck roles. The problems are all in the layer above that — synergy reasoning.

## The central misread

Gríma's trigger is **binary**. It fires once when he connects, regardless of whether he's a 1/4 or a 15/18. He's also already unblockable, so he doesn't need help connecting.

The generator appears to have read "combat damage trigger + can't be blocked" and reached for the voltron shelf. Count the cards whose primary function is *making a creature bigger*: Buster Sword, Orcrist, Explorer's Scope, Blitzball, Iron Spider, Bender's Waterskin, Hot Dog Cart, Interdimensional Web Watch. On a binary trigger, the power boost contributes almost nothing.

What actually multiplies Gríma's output is **more damage events**, not more damage:

- **Strionic Resonator** — copy the trigger outright. This is the single best card not in the deck.
- **Fireshrieker** / **Grappling Hook** — double strike is literally two triggers per swing.
- **Sword of Feast and Famine**, **Sword of Fire and Ice** (already in, correctly) — protection plus a real rider, unlike the vanilla pump equipment.

Dimir is thin on extra combats, which is a fair constraint, but the copy/double-strike axis is fully available and entirely absent.

The protection suite, by contrast, is genuinely well chosen. Lightning Greaves, Swiftfoot Boots, Champion's Helm, Darksteel Plate, Octopus Form, Hide on the Ceiling — a 1/4 commander that must connect repeatedly needs exactly this. That part of the build is correct.

## Cards with unmet preconditions

The deck has **two nonland creatures** (Orcish Bowmasters and Ingenious Prodigy). Several inclusions quietly assume otherwise:

- **Kindred Discovery** — five mana to draw roughly one card per turn off Gríma attacking. Not dead, but a terrible rate.
- **Skullclamp** — needs bodies or tokens. Bowmasters' Orc Army is the only real target.
- **Springleaf Drum** — needs an untapped creature you aren't attacking with. Mostly a blank.
- **Idol of Oblivion** — needs a token engine. Its realistic mode here is the {8} activation.
- **Inspiring Statuary** — improvise wants expensive nonartifact spells to cheat on. The curve tops out around four.

This is the most fixable class of bug: these are all threshold-gated cards, and the thresholds aren't being checked.

## No wincon

Serious question the deck can't answer: how does it win? Gríma's free-spell trigger is value, not pressure — it casts a random instant or sorcery off an opponent's deck, which is often a cantrip or a removal spell pointed at a board you don't have. Commander damage from a 1/4 with a couple of swords is a seven-plus turn clock with no redundancy if he eats an exile effect.

The Ring package (Call of the Ring, Birthday Escape, Lembas, Inherited Envelope, My Precious) is the closest thing to a kill — Ring level 4 drains three per hit — and it's still glacial.

Add finishers that use the ramp already present: **Torment of Hailfire**, **Exsanguinate**, **Blue Sun's Zenith** as a deck-out, or a couple of large evasive threats so Gríma isn't the entire offense.

## Mana base

This is the clearest tell that it was machine-built.

**Zero true duals.** 13 Island, 13 Swamp, Command Tower, City of Brass, and then three lands that don't belong: **Exotic Orchard** (unreliable in two colors), **Plaza of Heroes** (you have maybe three legendary permanents), **Spire of Industry** (pays life to do what a Watery Grave does better). Two taplands round it out.

Then **Cryptic Command** at {1}{U}{U}{U} on roughly 18 blue sources. Karsten's math wants ~23–24 for triple pips at four mana. That's a straight-up castability failure, and there's no pip-weighting pass catching it.

The replacements are obvious and mostly cheap: Watery Grave, Darkslick Shores, Drowned Catacomb, Choked Estuary, Undercity Sewers, Morphic Pool, River of Tears, Underground River, Bojuka Bog, Otawara, Takenuma. Also add Dimir Signet and Talisman of Dominance over Astral Cornucopia.

## Missing staples

For a Dimir deck at this bracket: no Counterspell, no Swan Song, no An Offer You Can't Refuse, no Rhystic Study, no Mystic Remora, no Toxic Deluge, no Reliquary Tower. Meanwhile it runs Sword of Fire and Ice and Orcish Bowmasters, so it isn't budget-capped. The pool skews hard toward 2025–26 sets with almost no pre-2020 cards outside a few reprints — if your selection is weighted by recency, set popularity, or EDHREC synergy scores computed on small recent samples, that's likely the cause.

## Heuristics worth adding to Deck Tome

1. **Classify the commander's payoff shape** before selecting. Binary-trigger commanders (connect once) vs. scaling ones (damage matters) vs. count-based ones. Gríma should down-weight +X/+X and up-weight trigger copiers, double strike, and extra combats.
2. **Evasion redundancy check.** If the commander already has can't-be-blocked, flying, menace, or shadow, discount cards whose only function is granting that keyword.
3. **Precondition solver.** Tag cards with numeric requirements (≥10 creatures, token generator present, ≥18 artifacts, ≥8 of a creature type, ≥3 colors) and run a validation pass after selection, cutting anything unmet. Applies to lands too — Exotic Orchard, Plaza of Heroes, Spire of Industry all failed here.
4. **Pip-weighted mana check.** Compute colored sources against Karsten thresholds per spell; reject or flag spells whose requirements can't be met, and fill the nonbasic slots from a color-identity-derived dual list before reaching for generic filler.
5. **Require a win condition and estimate a clock.** A deck should fail validation if it can't produce a kill — count damage output, drain, X-spells, alt-wins, and reject builds whose fastest kill is 8+ turns.
6. **Effect-class saturation caps.** Seven counterspells is fine; eight pump equipment on a binary trigger is not. Diminishing-returns curves per effect class, weighted by how much the commander cares.
7. **Bracket reporter.** You're already close to compliant — surface it explicitly: Game Changer count, tutor count, combo detection, extra turns, MLD, plus a power estimate that's separate from the rules-legal bracket. Those two numbers disagreeing here (rules: 3, power: 2.5) is exactly the insight a user wants.

The bones are good. It built a coherent Dimir control shell with real interaction and real protection. What it's missing is the layer that asks *what does this specific commander actually want* and *can this pile of good cards close a game*.
```

## The request

```text
Ingest this feedback as it relates to session ID z1hshyY6Npig1FN2NuV7

We need to consider win condition more carefully when constructing decks. Evaluate how this may be done. We also need to prioritize better mana cards for tier 4/5 decks.
```
