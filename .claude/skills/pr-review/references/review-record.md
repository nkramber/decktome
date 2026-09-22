# PR review: the review record

Part of the `pr-review` skill. Load this file before you write a review record or correct the pull request body. Load it too before you read the head and the verdict for `review-gate`.

## Review record

Use one file for each pull request in `docs/reviews/` (D-803). Keep its name and its finding ids on each repeat review. For a new record, use `docs/reviews/pr-<number>.md` with the GitHub number, not the roadmap id. Name a provider only in the review record and in the `Author:` line of the hand-off (D-806).

The `review-gate` check reads this file (D-804). Three parts of it are machine-read. Keep their form exact:

| Part | Exact form | Rule |
|---|---|---|
| The file name | `docs/reviews/pr-<number>.md` | The number is the GitHub number of the pull request. |
| The head field | `- Head: ` and the hash in backticks, in the `## Identity` list | The hash is the effective head, of 7 or more letters. The rule reads that list alone. |
| The verdict | The verdict line of the `## Verdict` section, such as `**Ready for owner merge.**` | The section gives one bold span, and that span is the verdict name, with no other word. |

The effective head is the newest commit that changes a path outside the metadata set. The metadata set holds four paths of this pull request (D-752, D-803):

- `docs/reviews/pr-<number>.md`
- `docs/reviews/pr-<number>-response.md`
- `docs/SESSION-HANDOFF.md`
- `docs/reference/session-handoff-archive.md`

A commit that changes only those paths is a metadata commit, and it does not move the effective head. A commit that changes the record of another pull request moves it. The review commit holds the record and the hand-off entry alone, so it is a metadata commit. Record the effective head, and not the tip, when the review commit is the last commit.

A repeat review replaces the verdict of the `## Verdict` section. Put each earlier verdict in a section of its own, `## Earlier verdicts`. A second bold span in the `## Verdict` section is a fault, because the gate cannot know which span is current. Give the reason in prose after the name, with no bold.

Use this skeleton. Keep the heading text and the order.

```markdown
# Pull request <number> review

Date: <YYYY-MM-DD>

## Identity

- Pull request: <number>
- Target: `main`
- Base: `<sha>`
- Merge base: `<sha>`
- Head: `<effective head sha>`
- Branch: `<branch>`

## Provider gate

The author provider, the source of that fact, and the reviewer provider. The gate result against D-803.

## Intended behavior and scope

The intent, what the review inspected, and each affected contract. Each path of `git diff --stat`, as inspected or not inspected.

## Findings

One subsection for each finding, in severity order. Write "No finding." when the review found none.

## Out of scope

One line for each concern that a later item holds, with that item. No severity. Write "None." when there is none.

## PR comments

One line for each comment thread: the claim, the answer of the author, and what the review verified. Write "None." when there is none.

## Description edits

One line for each correction of the pull request body, with the old value and the new one. Write "None." when there is none.

## Verification

One line for each command or check, with its result. Name each check that did not run, and the reason.
- Push: <sha> is the head of origin/<branch>, verified with gh pr view.

## Open questions and accepted risks

Each open question with its row in `docs/owner-questions.md`, and each accepted risk with its D-id.

## Verdict

**<Blocked | Changes required | Ready for owner merge>.** This verdict applies to head `<sha>`. The reason in one or two sentences.
```

## Finding format

Give each finding a stable id: the letter `P`, the severity number, a hyphen, and an index. `P1-1` is the first P1 finding. Keep the id for the life of the pull request, and never renumber a finding.

```markdown
### P<severity>-<n>: <short title that states the defect>

Status: <open | fixed in `<sha>` | accepted risk, D-<id> | withdrawn>.

File: `<path>:<line range>`, or Commit: `<sha>`.

Trigger: the input or state that starts the defect.

Expected: the required behavior, with the contract or the D-id.

Actual: the observed behavior.

Consequence: the effect on the user, the data, the deploy, the cost, or the maintainer.

Correction: the smallest change that restores the contract.

Regression check: the command or test that proves the fix, and the result that must show.
```

A withdrawn finding stays in the file with the evidence that refuted it. Never delete a finding.

## Correct the pull request body

A pull request body that names a stale head, an old count, or a replaced correction misleads the owner at the merge. The reviewer corrects such a body directly. It needs no finding, and the author needs no extra pass for it.

The reviewer changes only a fact that the review verified:

- The effective head, the base, or the merge base.
- A count that the review ran, such as a test total or a finding total.
- A check result that the review read.
- A sentence that names a correction that a later commit replaced.

The reviewer never changes what the author says the pull request does, or why. It never changes a decision, a tradeoff, a recommendation, or a line that the owner ticks. It names no provider in the body (hard rule 6). Write one line for each edit under `## Description edits`, with the old value and the new one. A claim that is wrong in substance stays a finding, and the author corrects it.

## The review gate check

The `review-gate` check applies five rules (D-804). `docs/tools/review_gate.py` holds them:

1. RG 1: a pull request with the `review-override` label changes only documents and skills.
2. RG 2: a pull request with the label changes no decision row.
3. RG 3: `docs/reviews/pr-<number>.md` exists on the head.
4. RG 4: the verdict is `Ready for owner merge`.
5. RG 5: the head field of the Identity list names the effective head.

The label never covers `docs/tools/`, `.claude/hooks/`, the harness settings, or a workflow, because each one runs as code. The owner applies the label. The mode file `.github/review-gate-mode` on `main` reads `enforced` or `advisory`. In advisory mode the check reports each fault and passes.

RG 5 fails when the author pushes code after the approval. That result is correct. Assess the new diff, then update the head field and the verdict together. RG 5 does not fail when the last commits change only the metadata set.

The check runs from `main` on `pull_request_target`. So it cannot run on the pull request that creates or changes it. Such a pull request proves the command with its unit tests, and the next pull request gets the live check.
