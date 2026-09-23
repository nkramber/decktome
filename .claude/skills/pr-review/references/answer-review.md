# PR review: the author answer

Part of the `pr-review` skill. Load this file when you are the author, and you answer Gitar or a review, or you apply the `review-override` label.

## The Gitar pass

Gitar reviews each push. After each push, load the `gitar-review` skill and follow it. The Gitar pass comes before the Codex review, and it never replaces the Codex review (D-811).

- Answer each Gitar finding before you ask the owner for the Codex review.
- A reply names no provider, harness, or model (hard rule 6).
- Record the pass in the hand-off: the count of findings, and the commit that answered each one.

## Ask for the Codex review

When a current Gitar review holds no open finding, tell the owner that the pull request is ready for the Codex review. The owner starts the Codex session. Give the owner these facts:

- The pull request number and the effective head.
- The line `Author provider: Claude Code` in the hand-off record of the pull request.
- The result of `make verify` on the effective head.

## Apply the label

A pull request of documents alone can merge with no Codex review (D-812). Apply the `review-override` label yourself when all of these conditions are true:

- Each changed path is in the documentation set of D-814.
- A current Gitar review holds no open finding.
- Each check of the pull request is green, except `review-gate`.
- The completion gate of the `one-pr-one-session` skill holds.

The documentation set holds `docs/`, `.claude/`, `CLAUDE.md`, `AGENTS.md`, `README.md`, and `.github/pull_request_template.md`. These paths are not in it: `docs/tools/`, `.claude/hooks/`, `.claude/settings.json`, `.claude/settings.local.json`, and each workflow. A change of a decision row does not stop the label.

Apply the label with this command:

```bash
gh pr edit <number> --add-label review-override
```

The label event runs `review-gate` again. Read its result.

When you push again after the label, remove the label first. Apply it again when each condition above holds again.

## Answer the findings of a review

**A finding is a claim, not a fact.** A review can be wrong. Examine each finding against the evidence before you change anything.

1. Run `git fetch` and `git status --short --branch`.
2. Read the finding, and read the file and the lines that it names.
3. Reproduce the trigger. A finding that does not reproduce has no merit.
4. Read the contract that the finding cites. Look for a later decision that revises it.
5. Decide the result: full merit, partial merit, or no merit.
6. Correct each part with merit. Make the smallest change that restores the contract.
7. Record each result in `docs/reviews/pr-<number>-response.md`.
8. Commit the response, the corrections, and the hand-off, then push.
9. Ask the owner for a repeat review by the Codex session.

Refute a finding when the evidence supports it:

| Reason | What to show |
|---|---|
| The finding reads a rule too broadly. | Quote the rule. Name the other files that the broad reading also condemns. |
| The finding cites a revised decision. | Quote the later decision. |
| The trigger does not reproduce. | Give the command, the commit, and the result. |
| The correction breaks another contract. | Name the contract and the caller. |
| The finding asks for work of a later item. | Quote the roadmap entry and its gate. |

Never accept a finding only to close the review faster. Ask the owner when a finding and an owner decision conflict, and quote both.

## The response file

Write `docs/reviews/pr-<number>-response.md` when the verdict is `Changes required` or `Blocked`. The `review-gate` check does not read it. For each finding, the file states:

- The result: full merit, partial merit, or no merit.
- The evidence, for partial merit or no merit.
- The correction, with the file and the decision id.
- The regression check that ran, and its result.
