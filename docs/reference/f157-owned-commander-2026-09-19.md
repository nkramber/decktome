# F-157, the unowned commander and the repair turn, 2026-09-19

**Verdict: the cause of F-157 is one unowned commander, and not a card of the model.** M-18 read one `not_owned` finding and three `profile_off_band` findings in the repair reason of each of its three builds. A free replay names the card of that finding. It is Gríma, Saruman's Footman, the commander the reader named. The repair turn can never fix it, so each build spent a model call of 60 to 90 seconds for nothing.

Every number here comes from a free run. No paid target ran for this document.

## The method

A scratch Go test in `go/internal/generate` makes every count, and a copy sits in `.local/f157/` out of git. It reads the card snapshot of 2026-09-04, the newest on this machine. The request is the request of M-17 and M-18: Commander, the theme "opponent milling cards", Gríma, Saruman's Footman, owned only, and bracket 5.

The collection comes from the same ManaBox export of the owner, dated 2026-09-02 (D-756). The import path is the path of the app: `Detect`, `Parse`, `Resolve`, and `OracleCounts`.

The test does three things. It builds the owned-only shortlist of the app. It assembles a 99-card deck from that shortlist by a fixed role rule, with no model call. It then reads the findings and the repair decision that a real build reads.

| Input | Value |
|---|---|
| Export rows | 2,956 |
| Collection entries | 2,848 |
| Cards, copies counted | 5,335 |
| Distinct oracle ids | 2,178 |
| Rows that no card resolved | 1 |
| Shortlist | 201 cards |
| The assembled deck | 99 cards, 0 misses |

## The finding names the commander

| Severity | Code | Message |
|---|---|---|
| BLOCK | `not_owned` | Gríma, Saruman's Footman: the deck needs 1, the collection has 0 |

The export holds no copy of that card. It holds 1 copy of Gríma Wormtongue, a different card. M-18 records the same limit: the export predates the day the owner got the commander.

## The model named no card of the collection

Three code paths make that impossible in an owned-only build.

- The shortlist drops every card with no copy. `go/internal/candidates/candidates.go` reads `POOL_RULE_OWNED_ONLY` in three places, and each one refuses a card at 0 copies.
- `Normalize` matches a name of the model against the pool alone. A name outside the pool becomes a miss, and no miss reached this deck.
- `FromListOwned` adds the commander and the basic lands to the pool, beside the shortlist (`go/internal/generate/pool.go`). A basic land is exempt from the ownership check (D-37). So the commander is the one card of the deck that the collection can fail to cover.

## No repair turn can fix such a finding

`assemble` takes the commander from the request on every pass (`go/internal/generate/generate.go`). The answer of the model holds a card list, and that list never changes the command zone. So the block survives each repair turn, and it reaches the reader with the deck.

D-226 asks for exactly that report. The owner retired the not-owned commander question row on 2026-08-27, because the check after the build "costs no turn, names the card, and arrives with the deck". D-300 keeps the block in the deck view, and that view hides a `not_owned` warning.

## The control proves the cause

The test runs the same deck and the same pool a second time, with one change. The collection covers the commander.

| Measure | The replay | The control |
|---|---|---|
| `not_owned` findings | 1 | 0 |
| Findings the repair turn reads | 7 | 6 |
| Profile findings alone | no | yes |
| A repair turn runs | yes | **no** |

So the unowned commander alone bought the repair turn. A profile finding buys no repair turn since PR-33 (D-609).

## One count differs from M-18

The repair reason of this deck names one `not_owned` finding and **six** `profile_off_band` findings. M-18 read three. The decks differ: M-18 built three decks with the model, and this test assembles one deck by a fixed rule. The kind of each finding matches, and the count of the profile findings does not.

## The bracket 5 floors the collection can not reach

The uncapped shortlist holds all 475 owned cards that the pool rule lets through. A floor the whole pool misses is a floor no deck of this collection reaches.

| Power floor | The whole owned pool | Bracket 5 wants |
|---|---|---|
| `fast_mana` | 7 | 6 or more |
| `finisher` | 18 | 1 or more |
| `game_changer` | 1 | 8 or more |
| `tutor` | 1 | 4 or more |

The gap note of D-709 already answers this for the reader. It names Vampiric Tutor, Demonic Tutor, Imperial Seal, and Mystical Tutor for the tutor floor, and it marks each one "to buy".

## The change

`fixable` drops every finding that names a commander of the deck, before the build decides on a repair turn (D-762). The deck keeps the finding. `go/internal/generate/ownedcommander_test.go` holds two tests, and both fail on the old code:

- The M-18 shape, one ownership block on the commander and three profile findings, buys no repair turn.
- A whole build with an unowned commander spends one model call, and not two.

## Limits

- The export is dated 2026-09-02, and the app built the deck of the review on 2026-09-14. A later collection holds more cards.
- The snapshot of 2026-09-04 is the newest on this machine. The deployed app read a snapshot of 2026-09-14.
- The local quality model differs from the deployed one, as it did for M-17 and M-18.
- No number here reads a deck of a model. The deck of this test comes from a fixed role rule.
- The three builds of M-18 read no card name for their `not_owned` findings. The test of M-18 read the reason alone, so the identity rests on the code paths above.

## What the plan takes from F-157

- F-157 closes as fixed (D-762, D-763).
- The premise "an owned-only request must never name a card the collection does not cover" is wrong for the commander. The reader names the commander, and D-226 permits a commander outside the collection.
- The thin pool half closes with the ceiling table above (D-763). No build tells the reader before the build that a bracket sits outside the reach of their collection.
