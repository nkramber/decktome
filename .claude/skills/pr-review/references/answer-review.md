# PR review: the author answer

Part of the `pr-review` skill. Load this file when the session is the author and answers Gitar or a review. The session of the author stays bound to its pull request, and it answers each review there (D-746).

## The Gitar pass

`scripts/gitar-state.sh` prints the state of Gitar (D-802).

- `paused`: no Gitar pass exists. Hand the pull request to the other provider when `make verify` and `make pr-check` pass.
- `on`: after each push, load `.claude/skills/gitar-review/SKILL.md` and follow it. Answer every Gitar finding before the hand-over to the other provider.

A reply to Gitar names no provider, harness, or model (hard rule 6).

## The hand-over

Tell the owner that the pull request is ready for the review of the other provider. Give the owner this prompt for that session:

```
Review PR #<number> of decktome. Role: reviewer.
Load `.claude/skills/one-pr-one-session/SKILL.md`, then `.claude/skills/pr-review/SKILL.md`.
Write `docs/reviews/pr-<number>.md`, add your hand-off entry, commit, and push to `<branch>`.
```

The owner starts that session. The author never starts a review of its own provider.

## Address review findings

**A finding is a claim, not a fact.** A review can be wrong. Assess each finding against the evidence before you change anything.

1. Run `git fetch` and `git status --short --branch`. Pull the review commit of the reviewer.
2. Read the finding, then read the file and the lines it names.
3. Reproduce the trigger. A finding that does not reproduce has no merit.
4. Read the contract that the finding cites, and each later decision that changes it.
5. Decide the disposition: full merit, partial merit, or no merit.
6. Correct each finding with merit. Use the smallest change that restores the contract.
7. Record each disposition in `docs/reviews/pr-<number>-response.md`.
8. Commit the response, the corrections, and the hand-off entry, and push them.
9. Tell the owner that the pull request is ready for a repeat review.

Push back when the evidence supports it. State the reason and show the proof:

| Reason to push back | What to show |
|---|---|
| The finding reads a rule too broadly. | Quote the rule. Name the other files that the broad reading also condemns. |
| The finding cites a superseded decision. | Quote the mark in the Question column, and name the current decision. |
| The trigger does not reproduce. | Give the command, the revision, and the result. |
| The correction breaks another contract. | Name the contract and the caller that it breaks. |
| The finding states a style preference. | Name the contract that the code does not break. |
| The finding repeats a risk that a decision accepted. | Quote the D-id and its accepted risk. |
| The finding asks for work outside the scope. | Quote the roadmap item and its gate. Name the item that holds the work. |
| The finding opens one id for the third time. | Name the three triggers, and ask the owner to settle the scope. |

Never accept a finding only to close the review faster. A wrong correction costs more than a written disagreement. Never make a correction wider than the contract that the finding names. When a finding and an owner decision conflict, quote both, and ask the owner (hard rule 11).

## The response file

The author answers a review in `docs/reviews/pr-<number>-response.md`. The `review-gate` check does not read it. Write one when the verdict is `Changes required` or `Blocked`. A clean first pass needs none.

For each finding, the response states:

- The disposition: full merit, partial merit, or no merit.
- The evidence, when the disposition is partial merit or no merit.
- The correction that landed, with the file and the decision id.
- The regression check that ran, and its result.

The response also lists each new D-id and F-id, and the final head. A disagreement belongs in the response file, with its evidence. Never delete a finding from the review record. The reviewer sets a refuted finding to `withdrawn`.
