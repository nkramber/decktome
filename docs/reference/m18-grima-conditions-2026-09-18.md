# M-18, the Gríma replay after PR-54, 2026-09-18

The roadmap gate of M-18 reads:

> Gate: a dated document holds the conditions of the replay, and the owner reads them before F-140 closes.

**Verdict: the gate holds. F-140 is fixed.** D-727 held F-140 as a record until a replay after PR-54. This document holds that replay. Three real builds of the same request each clear every condition line of M-17. The three builds cost $0.2350 in total (D-757). Every other number here is free.

**The cause is gone.** M-17 found that the theme "opponent milling cards" matched no card, because `themes.json` held no row for "milling" (F-142). The word-form rule of PR-54 now reads that word as the mill row (D-724). The theme matches 823 cards, and the shortlist carries threats, synergy cards, and wincons again.

## The method

A scratch Go test inside `go/cmd/deck-gate` made every count, and a copy sits in `.local/m18/` out of git. It reads the card snapshot of 2026-09-04, the newest on this machine. The condition rule is the rule of M-17, word for word, so the two documents compare.

The request is the request of M-17: Commander, the theme "opponent milling cards", Gríma, Saruman's Footman, owned only, and bracket 5. The replay builds the shortlist of the app, because `deck-gate` builds it (D-706).

The collection comes from a ManaBox export of the owner, dated 2026-09-02 (D-756). The session `z1hshyY6Npig1FN2NuV7` no longer exists in `decktome-prod`, so no session read was possible. The import path is the path of the app: `Detect`, `Parse`, `Resolve`, and `OracleCounts`.

| Input | Value |
|---|---|
| Export rows | 2,956 |
| Rows that no card resolved | 1 |
| Collection entries | 2,848 |
| Cards, copies counted | 5,335 |
| Distinct oracle ids | 2,178 |
| Legal cards in the color identity | 475 |

## The theme now matches

| Measure | M-17, 2026-09-14 | M-18, 2026-09-18 |
|---|---|---|
| The word "milling" | no row, and the generic rule fired no signal | reads the mill row through the word-form rule |
| Cards on theme, whole pool | 0 | 823 of 12,599 |
| Cards on theme that the owner owns | 0 | 52 |

The resolved theme reads the payoff tag `synergy-mill` and the tags `mill`, `mill-any`, and `mill-opponent`. The word "opponent" still matches no card, and the match reports it as unmatched. The mill row carries the signal, so the theme is not thin.

## The shortlist holds threats again

Each count reads the owned-only shortlist of the replay.

| Role | M-17 | M-18 |
|---|---|---|
| Land | 40 | 40 |
| Ramp | 32 | 32 |
| Removal | 31 | 31 |
| Draw | 30 | 30 |
| Synergy | 0 | 25 |
| Interaction | 20 | 20 |
| Wipe | 12 | 12 |
| Threat | 0 | 10 |
| Wincon | 0 | 1 |
| Total | 165 | 201 |

M-17 read no threat, no synergy card, and no wincon, because no card of the pool held a theme signal. The shortlist of M-18 holds 52 cards that the theme matched.

The uncapped build reads all 475 owned cards. It holds the same 10 threats, 25 synergy cards, and 1 wincon. So no cap hides a role, and the owned pool itself is thin at those three roles (F-157).

## The conditions of three built decks

The line of each class is the line of M-17: the lowest low decile over the four large groups of real lists. A condition fails when the deck holds fewer enablers than its line.

| Deck | Cards | Lands | Creature enablers | Token enablers | Artifact enablers |
|---|---|---|---|---|---|
| Line | | | 11 | 6 | 11 |
| The deck of the review | 99 | | **8, and it fails** | | |
| Build 1 | 99 | 32 | 20 | 8 | 33 |
| Build 2 | 99 | 32 | 18 | 8 | 26 |
| Build 3 | 99 | 31 | 24 | 8 | 26 |

Every condition of every build holds. The lowest creature count is 18, and the line is 11. The cards that need creatures read 13, 8, and 12 over the three builds, against 19 in the deck of the review.

The three builds agree, so one sample does not carry this result. The judge noise of D-230 does not apply, because no judge read these decks. The condition rule is deterministic.

## Every build ran a repair turn

Each of the three builds ran its repair turn. The reason of each build names one `not_owned` finding and three `profile_off_band` findings.

- `not_owned` means the model named a card that the collection does not cover. An owned-only request must never do this.
- `profile_off_band` means a feature of the deck sits outside the band of its bracket (PR-14A).

F-157 records this. No number holds the state after the repair turn, because the test read the reason and not the validation of the final deck. The thin owned pool is the likely cause of the off-band findings at bracket 5, and no measurement proves that yet.

## Limits

- The export is dated 2026-09-02, and the app built the deck of the review on 2026-09-14. The export holds no Gríma, so it predates the acquisition of the commander. The replay is conservative, because a later collection holds more cards.
- The snapshot of 2026-09-04 is the newest on this machine. The deployed app read a snapshot of 2026-09-14 for the deck of the review.
- The local quality model differs from the deployed one, as it did for M-17.
- The condition rule reads rules text. A card with rare text can take the wrong class.
- No measurement reads the state of each deck after its repair turn.
- Three builds are three samples. A fourth build can read a lower count, and the margin over the line is wide.

## What the plan takes from M-18

- F-140 closes as fixed. PR-54 removed its cause, and three builds measure the result.
- F-157 opens on the thin owned pool and the repair turn of every build.
- No new code change comes from this measurement.
