# M-12, the rules quantities of the Commander lists, 2026-09-11

This note plans M-12, the rules detector of D-666 and D-675. A free dump reads the quantities a rules detector checks. It reads every Commander list of gate run 19 and the 25 built decks of deck gate run 16. No run called a provider. Throwaway patches made the dump, and none of them merges.

**The result: the rules read the built decks right, and as the score they likely fail both bars** (F-116). Every candidate set flags decks 19 and 22 alone among the built decks. At the cuts that catch the broken copies, the rules flag about a tenth of the top lists. The land check flags 4 percent of them at a shortfall of 10, and the colors check flags 6 to 8 percent at 0.8. The owner chose to fit two designs over a grid of thresholds (D-676, D-677).

## The source

Frank Karsten, "How Many Lands Do You Need in Your Deck? An Updated Analysis", TCGplayer, 2022-07-29. A session read it on 2026-09-11 through `https://infinite-api.tcgplayer.com/content/article/cd1c1a24-d439-4a8e-b369-b936edb0b38a/`. The web page draws its text in the browser, so a plain fetch of the page reads no text.

- **The 99-card formula**: "31.42 + 3.13 * average mana value of your spells – 0.28 * number of cheap card draw or mana ramp spells". It ports the 60-card formula, counts the commander as a pseudo-companion, and removes 1.35 lands for the free mulligan and the draw on turn one.
- **The error**: Karsten calls the port "an imprecise back-of-the-envelope estimate". The 60-card model reads an R-squared of 0.395 and a root mean square error of 2.75 lands.
- **The average mana value**: the mana value of the nonland cards over their count. A land/spell MDFC counts as its front side.
- **A cheap card draw spell**: a nonland card of mana value two or less with a draw phrase in its Oracle text. The phrases are "draw a card", "draw two cards", "draw three cards", "draws cards", "draws two cards", and "draws three cards". The text must not hold "{4}", "Blood token", or "investigate", and a creature must also hold "when" and "enters". A noncreature spell with "look", "library", "put", and "your hand", and no "pay " or "pays", also counts. So does a spell that cycles for one mana.
- **A cheap mana ramp spell**: a nonland card of mana value two or less, not already cheap draw, whose text holds "add ". The text must not hold "add its ability" or "add a lore counter", and a creature must not hold "dies". Three more shapes count: a land search with no "sacrifice", "enchanted land is tapped" with "adds an additional", and "put a creature card with" with "from your hand onto the battlefield".
- **No cut**: the article gives no shortfall that makes a deck broken. The source table of 2022-08-02 gives no share of the need that does (`bracket-profile-2026-09-02.md`).

**The counter.** A throwaway Go function follows these rules word for word. A test ran it over the 33 cards the article names, in the card snapshot of 2026-09-04. Every card reads as the article says, Omen of the Sea included.

The profile reads the average mana value over the cards whose types hold no Land. So a spell with a land face counts as a land there, and Karsten counts its front side. The dump keeps the profile's value.

## The lists

| Tier | Source | Lists |
|---|---|---|
| baseline | MTGJSON precons | 189 |
| typical | EDHREC average decks | 1,040 |
| good | Topdeck.gg 3,950, MTGTop8 50 | 4,000 |
| great | Topdeck.gg 3,984, MTGTop8 16 | 4,000 |

The copies after the synergy check: 1,229 lands, 1,229 curve, 729 colors, and 2,732 synergy. So 7,934 of the 8,000 top lists are cEDH tournament lists.

## The quantities

Each cell reads the median, then the 5th and the 95th percentile. The shortfall is Karsten's need less the lands, with his cheap count.

| Group | Lands | Average mana value | Cheap | Shortfall | Color sources | Mana on turn four | Fast mana |
|---|---|---|---|---|---|---|---|
| precons | 38 (36, 41) | 3.68 (3.05, 4.27) | 7 (3, 12) | 2.7 (-0.6, 5.3) | 1.08 (0.79, 1.43) | 4.30 (3.96, 4.63) | 1 (1, 2) |
| average decks | 36 (34, 39) | 3.02 (2.33, 4.08) | 10 (6, 17) | 2.6 (-1.5, 6.3) | 1.30 (0.88, 1.78) | 4.57 (4.14, 5.29) | 2 (1, 4) |
| good lists | 27 (23, 37) | 2.06 (1.75, 3.25) | 19 (8, 25) | 6.0 (0.9, 9.4) | 1.07 (0.78, 1.38) | 5.13 (4.26, 5.82) | 9 (1, 16) |
| great lists | 27 (23, 37) | 2.05 (1.75, 3.24) | 19 (8, 25) | 6.0 (1.2, 9.6) | 1.08 (0.79, 1.37) | 5.16 (4.32, 5.81) | 9 (1, 15) |
| lands copies | 24 (23, 26) | 3.15 (2.51, 4.01) | 10 (6, 18) | 14.0 (10.6, 17.1) | 0.82 (0.50, 1.24) | 3.94 (3.50, 4.62) | 2 (1, 4) |
| curve copies | 36 (34, 39) | 4.42 (3.75, 5.42) | 3 (0, 9) | 8.4 (4.9, 12.0) | 1.12 (0.76, 1.58) | 4.00 (3.66, 4.50) | 0 (0, 0) |
| colors copies | 36 (34, 40) | 3.17 (2.39, 4.14) | 9 (5, 17) | 2.5 (-1.1, 5.5) | 0.56 (0.24, 0.79) | 4.65 (4.29, 5.32) | 1 (1, 4) |
| synergy copies | 36 (34, 39) | 3.22 (2.75, 3.95) | 6 (2, 12) | 4.0 (0.2, 6.9) | 1.27 (0.78, 1.76) | 4.21 (3.77, 4.81) | 1 (0, 3) |

- **The top lists sit short of Karsten's need.** They run 27 lands over a curve of 2.05 with 9 fast mana. The formula counts a cheap spell at 0.28 of a land and reads no fast mana past that.
- **The curve break raises the need and removes the cheap spells.** So a curve copy sits short too, but it overlaps the top lists.
- **The goldfish does not split them.** The lands copies make 3.94 mana on turn four and the precons 4.30. The share of first hands with two to four lands reads 0.54 for the lands copies and 0.60 for the top lists.

## The single checks

A pair wins when the check flags the copy and passes its precon, or flags both and reads the copy as worse. The flag shares read each group.

| Check | Lands pairs | Curve pairs | Colors pairs | Precons | Average decks | Good lists | Great lists |
|---|---|---|---|---|---|---|---|
| shortfall 4 or more | 189 of 189 | 185 of 189 | 0 of 178 | 0.22 | 0.23 | 0.81 | 0.82 |
| shortfall 6 or more | 189 of 189 | 173 of 189 | 0 of 178 | 0.01 | 0.06 | 0.49 | 0.49 |
| shortfall 8 or more | 189 of 189 | 96 of 189 | 0 of 178 | 0.00 | 0.02 | 0.15 | 0.16 |
| shortfall 10 or more | 188 of 189 | 26 of 189 | 0 of 178 | 0.00 | 0.00 | 0.04 | 0.04 |
| shortfall 12 or more | 174 of 189 | 2 of 189 | 0 of 178 | 0.00 | 0.00 | 0.01 | 0.01 |
| average mana value over 4.2 | 1 of 189 | 184 of 189 | 0 of 178 | 0.09 | 0.04 | 0.01 | 0.01 |
| average mana value over 4.4 | 0 of 189 | 174 of 189 | 0 of 178 | 0.02 | 0.03 | 0.01 | 0.01 |
| color sources under 0.75 | 125 of 189 | 10 of 189 | 163 of 178 | 0.04 | 0.03 | 0.03 | 0.02 |
| color sources under 0.8 | 148 of 189 | 24 of 189 | 170 of 178 | 0.06 | 0.04 | 0.08 | 0.06 |
| mana on turn four under 3.75 | 69 of 189 | 56 of 189 | 0 of 178 | 0.02 | 0.00 | 0.00 | 0.00 |

The role count of the app reads more cheap spells in the top lists, a median of 22, and it tells the same story. A shortfall of 4 or more flags 72 percent of them there.

## The combined sets

| Set | Lands pairs | Curve pairs | Colors pairs | Lost | Precons | Good lists | Great lists | Built decks flagged |
|---|---|---|---|---|---|---|---|---|
| shortfall 10, curve 4.2, colors 0.75 | 188 | 184 | 163 | 21 | 0.12 | 0.06 | 0.06 | 19, 22 |
| shortfall 10, curve 4.2, colors 0.8 | 188 | 184 | 170 | 14 | 0.14 | 0.11 | 0.10 | 19, 22 |
| shortfall 10, curve 4.4, colors 0.75 | 188 | 176 | 163 | 29 | 0.05 | 0.06 | 0.06 | 19, 22 |
| shortfall 10, curve 4.4, colors 0.8 | 188 | 177 | 170 | 21 | 0.08 | 0.11 | 0.10 | 19, 22 |

**The room of the precon bar.** The bar asks 749 of 788, so it allows 39 lost pairs. Gate run 19 loses 36 on the synergy axis. So the lands, curve, and colors axes together keep a room of 3 lost pairs. The dump counts a pair that neither check flags as lost, and the ladder can win some of those pairs. Of the 14 losses of the second set, 2 are certain, and the ladder decides 12.

**The great-over-precon bar, by estimate.** A flagged top list scores under the floor and loses to every unflagged precon. The second set flags about 10 percent of the great lists and 14 percent of the precons. So the bar falls near 0.88 against its floor of 0.90. Only a fit measures it.

## The built decks

| Deck | Tier today | Owner | Lands | Average mana value | Cheap | Need | Shortfall | Color sources |
|---|---|---|---|---|---|---|---|---|
| 1 | bad | typical | 36 | 3.37 | 2 | 41.4 | 5.4 | 0.86 |
| 3 | bad | baseline | 33 | 3.23 | 6 | 39.8 | 6.8 | 1.21 |
| 4 | bad | typical | 37 | 3.08 | 12 | 37.7 | 0.7 | 1.01 |
| 11 | bad | typical | 36 | 2.87 | 7 | 38.4 | 2.4 | 0.95 |
| 13 | bad | baseline | 36 | 3.06 | 13 | 37.4 | 1.4 | 0.88 |
| 14 | bad | typical | 36 | 3.21 | 9 | 38.9 | 2.9 | 1.11 |
| 15 | bad | baseline | 37 | 3.10 | 14 | 37.2 | 0.2 | 1.18 |
| 16 | bad | typical | 37 | 3.15 | 8 | 39.0 | 2.0 | 1.20 |
| 19 | bad | | 36 | 3.24 | 12 | 38.2 | 2.2 | 0.71 |
| 20 | bad | | 36 | 3.30 | 6 | 40.1 | 4.1 | 1.76 |
| 21 | bad | baseline | 36 | 3.24 | 7 | 39.6 | 3.6 | 1.18 |
| 22 | bad | bad | 36 | 3.14 | 14 | 37.3 | 1.3 | 0.70 |

No built deck runs an average mana value over 3.40, and no built deck falls 8 lands short. So every set of the grid flags decks 19 and 22 alone. Deck 22 is the one real defect of D-659, and it reads 0.70 on its color sources.

## The two designs

A rules check says "flagged" or "not flagged", so it never tips an unflagged deck to bad (F-115). A flagged deck grades bad, and every other deck takes the tier of the ladder.

- **Design A**: the rules set the tier, and the fitted detector keeps the score. No bar can move, because the score stays. The bars then guard a score no reader sees, and the reasons must name the rule.
- **Design B**: the rules and a fitted synergy check replace the detector in the score and the tier. The bars then read what a reader sees. A flagged deck scores under the floor, lower the further past its cut. The dump says both bars likely fail.

Under the rules alone, both designs grade the built decks the same. Read from the ladders of `notier` in M-11, that is 9 graded bad, a judge agreement of 10, and 5 owner matches. No fit measured it. **B also carries the fitted synergy check in its tier.** So B grades the built decks as A does only where that check flags no built deck.

The owner chose to fit both designs (D-676) over a grid of eight combinations each (D-677). The owner decides which one PR-40 builds, if any.
