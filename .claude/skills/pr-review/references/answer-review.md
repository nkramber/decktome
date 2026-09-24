# PR review: the author answer

Part of the `pr-review` skill. Load this file when you are the author, and you answer Gitar or a review, or you apply the `review-override` label.

## The Gitar pass

Gitar reviews each push. After each push, load the `gitar-review` skill and follow it. The Gitar pass comes before the Codex review, and it never replaces the Codex review (D-811).

- Answer each Gitar finding before you ask the owner for the Codex review.
- Do not answer a Gitar comment that names no specific item, such as a plan notice or a pause note. It is not a finding (D-842).
- A reply names no provider, harness, or model (hard rule 6).
- Record the pass in the hand-off: the count of findings, and the commit that answered each one.

## Start the Codex review

When a current Gitar review holds no open finding, start the review yourself (D-823). The owner approved each round of the loop (D-831).

1. Write the line `Author provider: Claude Code` in the hand-off record, and push it.
2. Wait until each check of the tip completes.
3. Run `make codex-review PR=<number>` in the background, and wait for the notice of its end.
4. Read the last line of the output: `outcome: <name> (exit <n>)`.
5. Run `git pull --ff-only`, because the reviewer pushed the record.

| Outcome | Next step |
|---|---|
| approve | Go to the auto-merge of the `one-pr-one-session` skill. The owner confirms the merge after the summary of four sections (D-834, D-836). |
| changes | Answer each finding with the procedure below, push, and do the Gitar pass again. |
| three-strike stop | Do the procedure of "The three-strike stop" below. |
| refusal | Correct the condition that the output names, then run the target again. |
| fault | Read the transcript that the output names. Ask the owner when the cause is not clear. |

A refusal spends nothing. The target refuses a dirty tree, a checkout that differs from origin, and an incomplete Gitar pass.

## The three-strike stop

The target exits 4 when a blocking finding is open at its third effective head (D-826). Do these steps:

1. Turn off the auto-merge with `gh pr merge <number> --disable-auto`.
2. Stop the fix loop. Change no file for that finding.
3. Ask the owner with `AskUserQuestion`.
4. Record the answer in `docs/reviews/pr-<number>-response.md`.
5. Record the answer as a new D- row too, when it sets a rule.

The question gives the finding, the evidence of the reviewer, each answer of the author so far, and the options with their pros and cons.

## Apply the label

A pull request of documents alone can merge with no Codex review (D-812). Apply the `review-override` label yourself when all of these conditions are true:

- Each changed path is in the documentation set of D-814.
- A current Gitar review holds no open finding.
- Each check of the pull request is green, except `review-gate`.
- The completion gate of the `one-pr-one-session` skill holds.

The documentation set holds `docs/`, `.claude/`, `CLAUDE.md`, `AGENTS.md`, `README.md`, and `.github/pull_request_template.md`. These paths are not in it: `docs/tools/`, `.claude/hooks/`, `.claude/settings.json`, the local settings file of the harness, and each workflow. A change of a decision row does not stop the label.

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
9. Do the Gitar pass, then start the repeat review with `make codex-review PR=<number>`.

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
