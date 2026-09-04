# PR-7 question gate

Run date: 2026-09-04. Conversations: 2026-08-31.

Verdict: PASS. 0 of 0 counted gate conversations used catalog questions only. The bar is 25. The set holds 0 gate conversations, and 0 of them start after a build and are not counted (A-9).

This is a partial run over `109`: 1 of 108 conversations ran. The catalog-only bar and the gate size bar need the whole set, so the verdict reads neither, and this run never stands as the gate. The item bars hold: no premature build, no dead end, no lint finding, and no expectation miss.

1 probe conversations ran beside the gate. They asked 3 questions, and the model offered 0 replacements. A probe explores catalog coverage and does not move the verdict (D-96).

No conversation reached a dead end (D-357).

The net of D-351 held back on 1 turns, because every question that was out is still inside its grace period. The reader answers it next turn, or the net closes it on the turn after. StallGrace is 2 turns: a group of sets by franchise turn 2 (budget).

The net of D-351 healed 1 turns. Each one asked nothing new and was not ready, so it closed the questions that were out and built with what it had. agentsvc.Chat runs the same net on every turn. A turn here is a catalog or classifier candidate, because the reader answered nothing the agent could read: a group of sets by franchise turn 3 (budget).

No conversation called itself complete with a slot unanswered.

The linter found no defective question (D-115).

## M-4 report

| Measure | Value |
|---|---|
| Conversations | 0 |
| Questions asked | 0 |
| From the catalog | 0 |
| Invented by the model | 0 |
| Replacements refused as rewords | 0 |
| Catalog questions that closed a slot | 0 |
| Invented questions that closed a slot | 0 |
| Catalog-only conversations | 0 |
| Median gap score | 0.00 |
| Fit threshold | 0.35 |

## Run

- Suite `questions`, run `pr7-question-gate-run37`, on 2026-09-04, commit `07856c2`.
- Partial run over `109`. It reads the bars of its own items, and it never stands for the suite.
- Roles: ask on `gpt-5.6-luna` (openai, effort low), classify on `gpt-5.6-luna` (openai, effort none).
- Versions: card snapshot 2026-09-04, questions prompt version 18, conversations `2026-08-31`, slots_snapshot `4`.
- Calls: 5. Cost: $0.0011. Time: 10 seconds.
- Cards: snapshot loaded, 2471 entries, 4316 cards, 1 rows unresolved.
- Tokens: 8961 input (7721 cached), 562 output.

## Conversations

### 109. a group of sets by franchise (probe)

Collection: false. Catalog: 3. Invented: 0.

Expected slots: every one met.

**Turn 1, the user:** Build the best possible deck from the sets with Marvel characters.

- [catalog slot=format row=format fit=0.90 filled=true] Which format should the deck use: Commander, Standard, or Modern?
  - Options: Commander / Standard / Modern
- [catalog slot=colors row=colors fit=0.90 filled=true] Do you have a color preference?
- [catalog slot=budget row=budget fit=0.90 filled=false] What spending limit, if any, should I use for cards you do not own?

**Turn 2, the user:** Commander, bracket 3, and you pick the commander.


**Turn 3, the user:** Any colors are fine.


