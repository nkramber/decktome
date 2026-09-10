# PR-29 gate, 2026-09-10

The roadmap gate of PR-29 reads:

> The Commander synergy axis reads 0.95 or better, and no axis falls under its run 14 read. The tier judge lane over deck gate run 16 reads the built decks, $0.36. The count graded bad in the explain mode falls under 19 of 24.

**Verdict: UNMET.** The Commander synergy axis reads **0.81** against a bar of 0.95. No axis fell. The two reader-facing numbers need the paid judge lane, and the owner held it (D-648).

The owner merged the item at 0.81 on this evidence, and M-8 audits the bar before anybody fits again (D-649).

## The measurement needs a control

Run 14 is the reference the plan names, and it is no fair one. The meta store grew after it: Standard reads 7356 lists in run 14 and 8228 today, and Modern 17666 and 20292. A comparison against run 14 would charge that growth to PR-29.

So the numbers below come from two runs on the same day, over the same meta store and the same card snapshot of 2026-09-04. The control is the same code with the change reverted.

| Read | Run 14 | Control | PR-29 | Bar |
|---|---|---|---|---|
| Commander synergy axis, own copy | 0.70 | 0.70 | **0.81** | 0.95 |
| Commander precon over own copy | 0.88 | 0.88 | **0.92** | 0.95 |
| Commander precon over bad, cross | 0.86 | 0.86 | 0.91 | - |
| Commander great over precon | 0.96 | 0.96 | 0.95 | 0.90 |
| Commander accuracy | 0.65 | 0.65 | 0.66 | - |
| Standard precon over own copy | 0.97 | 1.00 | 1.00 | 0.95 |
| Standard accuracy | 0.47 | 0.46 | 0.44 | - |
| Modern precon over own copy | 0.99 | 0.99 | 0.99 | 0.95 |

The other three own-copy axes read 1.00 in both runs: lands, curve, and colors.

A second run of the same code reproduced the PR-29 column exactly. This gate has no run-to-run noise, unlike the question gate (D-230).

## The mechanism holds

`casual_rate` is the feature the synergy break moves most, at 1.03 standard deviations, ahead of `synergy` at 0.91 and `unseen_share` at 0.93. Before PR-29 the most-moved feature was `unseen_share`.

The own-copy misses fell from 118 to 74, and the fall is in every kind:

| Miss kind | Control | PR-29 |
|---|---|---|
| Both passed, the ladder put the copy at or above the precon | 54 | 40 |
| Precon flagged, the detector flagged the precon and passed the copy | 24 | 6 |
| Both flagged, and it read the precon as the more broken | 40 | 28 |

The detector reads the same standardized feature vector as the ladder, so a corpus change reaches both. The detector misses fell from 64 to 34.

The fitted weights are the sense the plan asked for. Commander reads `casual_rate` +0.224 and `casual_synergy` +0.074. Standard and Modern read both negative, which is right: in a 60-card format, a deck that reads like a precon is the worse deck.

## The corpus sizes

Read from the stored model of 2026-09-10.

| Format | Top rates | Top pairs | Casual rates | Casual pairs |
|---|---|---|---|---|
| commander | 9942 | 200000 | 10278 | 103183 |
| modern | 891 | 6786 | 4747 | 806 |
| standard | 815 | 6032 | 474 | 161 |

Commander holds a healthy casual corpus, and it sits under the cap of 200000, so no pair is dropped. The two 60-card formats are thin: Modern holds 806 casual pairs and Standard 161. PR-35 answers that (D-650).

The model grew from 4.85 MB to 7.23 MB, gzipped. That is 49 percent and not the doubling the pair cap allows, because the casual pairs stop well under the cap.

## The measurement refutes two predictions of the plan

### The four weights against sense moved further against sense

The plan says: "The four weights against sense have a reason to move." They moved, and every one moved the wrong way.

| Weight | Control | PR-29 |
|---|---|---|
| `draw` | -0.227 | -0.267 |
| `wipe` | -0.146 | -0.184 |
| `commander_decks` | -0.129 | -0.200 |
| `land` | -0.256 | -0.258 |

One reading, untested: those shape features were partly standing in for "this is a casual deck". The casual features carry that signal directly now, so the shape features measure a purer distance to a tournament list, which is more negative. That reading makes the new numbers the more honest ones. Nothing here proves it.

### The gate of 0.95 is out of reach by corpus work

Two further corpus features were built, measured, and reverted on 2026-09-10. Neither is in the shipped code, and both are recorded here so nobody builds them again.

**`casual_unseen_share`**, the share of copies no casual list holds. The reasoning: the break draws its filler from the cards real lists play, and that pool is dominated by the great and the good lists, so a filler card should be absent from the casual lists. The result: the target axis moved from 0.81 to 0.81, and three other bars fell hard. Commander cross pairs read 0.76 against 0.91, Modern own copy 0.90 against 0.99, and every cross axis fell. The feature taught the model that a deck holding cards casual lists do not hold is a broken deck, and that describes every good cEDH deck.

**`casual_pair_share`**, the share of card pairs that lift at all, rather than the mean lift. The result: 316 wins of 389 against 315, one pair. The weights say why: `casual_synergy` read +0.284 and `casual_pair_share` -0.262. The two are collinear, the fit split them with opposite signs, and they cancel.

## Why the gap stays

The ladder orders the five tiers, and the same weights decide this bar: the bar asks whether the ladder scores a precon above its copy. `casual_synergy` is the feature the break moves most cleanly, and it carries +0.074, because it is mediocre at ordering the whole ladder. One fit serves two jobs (F-94).

The break replaces half a precon's spells with cards drawn from lists people play (D-488). For a weak precon that is arguably the better pile of cards, and only the pairing tells the two apart. M-8 asks whether the bar is true before anybody fits again (D-649).

## What this gate does not measure

**The two reader-facing numbers of D-648.** The explain mode count of built decks graded bad, 19 of 24 at run 13, and the tier judge agreement, 7 of 25 at run 4. Both need the judge lane at $0.36, and the owner held the spend until the code settles. Every quality item reports all three numbers from now on, and an item that moves none is dropped.

**Whether the bar is true.** M-8.
