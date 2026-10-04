# The author response to the review of #283

Author provider: Claude Code.

## P2-1: The classifier drops hard avoidance

Result: full merit.

Evidence: questions prompt version 21 told the classifier to write the thing to avoid "without less, fewer, or no". So "no artifacts" reached `Slots.avoid` as `artifacts`, and `avoidMatch` read it as a soft request. The test `TestF212AvoidRanksLower` gave "no artifacts" to the builder directly, so no test read the path from the classifier.

The review also exposed a second gap in the same function. The slot joins the things of each message with ";". `avoidMatch` read one hard word for the whole slot, so "artifacts; no creatures" made "artifacts" hard too.

Correction:

- `go/internal/questions/prompts.go`: avoid keeps "no" when the user wants none of the thing, and `PromptVersion` reads 22 (D-1122). The owner chose a change of the prompt over a parser of the message.
- `go/internal/candidates/theme.go`: `avoidMatch` reads each part of the slot alone, through the new function `avoidPart`.
- `docs/decisions.md`: the row of D-1122 records the owner choice. `docs/design-roadmap.md` and `docs/SESSION-HANDOFF.md` name prompt version 22.

Regression checks:

- `TestAvoidKeepsAHardRequest` in `go/internal/questions/avoid_test.go`: the classifier output "no artifacts" reaches the slot beside an earlier soft "artifacts". It passes.
- `TestF212AvoidRanksLower` in `go/internal/candidates/shortlist_f212_test.go` reads the slot strings that the classifier writes: "artifacts", "artifacts; no creatures", and "creatures with flying; no artifacts". It passes. On the matcher of `63bf6c7`, it fails: "no creatures" made "artifacts" hard, and it lowered Swiftfoot Boots.
- `go/internal/agentsvc/build.go` copies the slot to `Request.Avoid` with no change, so these two tests cover the path from the classifier to the ranking.
- A question gate run on prompt version 22 reads the classifier on the real model (D-66).

## Scope added after round 1

The owner put two more findings into this pull request after round 1. Neither one answers a finding of the review.

- F-221: a deck of bracket 1 to 4 reads the mean of the model for the cEDH signal. So that signal moves no grade and names no reason (D-1123). The roadmap gives that signal to bracket 5 alone. `TestCEDHSignalReadsBracketFiveAlone` fails on the old code. Quality gate run 26 reads PASS on `6503306`. Its decks name no bracket, so it shows that a deck with no bracket keeps its grade.
- F-222: `chat-probe -decks-out` writes each deck that reached the user, and `user-case -decks` reads the deck bars of the case (D-1124). The paid replay of `6503306` passes each shortlist bar and each deck bar, and its grade names no cEDH reason.
- Question gate run 60 reads PASS on prompt version 22 for $0.1043, with 75 of 75 conversations on catalog questions alone.
- F-223: a deck of bracket 1 to 4 also reads the mean of the model for the deck count of its commander (D-1123). The typical rung holds the average decks of the 100 most built commanders alone, so the fit read a rare commander as a better deck. `TestCEDHSignalReadsBracketFiveAlone` covers both signals. A free rescore of the replay decks gives 0.565 and 0.510, above the bars of 0.354 and 0.378.
- The case adds four deck bars to build 2: cantrips, hand size, Equipment, and artifact creatures. The baselines come from the two stored decks of the user, read on 2026-10-04. Those decks fail 9 deck bars, and the replay decks pass each one.
