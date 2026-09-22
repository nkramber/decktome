---
name: pr-review
description: Review a pull request as the provider that did not write it, or answer such a review as the author. Require the opposite provider, precise evidence, regression checks, and a verdict for one effective head. A finding is a claim, not a fact, and the author can refute one with evidence. Use for each review, each repeat review after a fix, and each answer to review findings.
---

# PR review skill

Review the change as the engineer who answers for its effect on the whole system. Judge correctness, contracts, failure recovery, test quality, and future maintenance. Apply this standard to code, tools, CI, skills, and document pull requests. A green test suite or a persuasive pull request body does not prove correctness.

The owner decisions are D-803 to D-806. `CLAUDE.md` holds the other rules, and this skill does not repeat them.

## Reference files

This file holds the rules of every review. Each reference file holds the rules of one case. Load a reference file before the work of its case.

| Reference file | Load it when |
|---|---|
| `references/project-contracts.md` | The pull request changes code, tools, CI, or the text of a contract. |
| `references/repeat-review.md` | The session reviews a pull request again after a correction. |
| `references/review-record.md` | The session writes or edits a review record, or reads the head and the verdict for `review-gate`. |
| `references/answer-review.md` | The session is the author, and it answers a review or Gitar. |
| `references/commit-and-end.md` | The session commits a review record or a response file, and before the session ends. |

## Mandatory provider gate

**The reviewer MUST NOT come from the provider that wrote the pull request.** This rule applies before the review starts and before any approval (D-803).

| Provider that wrote the pull request | Required reviewer |
|---|---|
| Claude Code, Anthropic | Codex, OpenAI |
| Codex, OpenAI | Claude Code, Anthropic |

A different model, account, session, or subagent of the same provider does not qualify. A prompt that gives the other provider's name does not change the actual provider. A self-check of the author and the automated tests do not satisfy this gate.

A substantive change alters code, data, configuration, requirements, or instructions that a tool or an agent runs. Review findings and test reports alone do not make the reviewer an author.

1. Identify the actual provider of this session from the active environment.
2. Identify each provider that wrote a substantive change or fix in this pull request.
3. Read the `Author:` line of each hand-off session entry of this pull request (D-806).
4. Read the owner's statement and the earlier review records of this pull request too.
5. Record the providers, the source of each fact, and the gate result in the review record.

The newest hand-off entry can describe a review, and not the authorship. A Git account alone does not identify the provider. Do not infer authorship from prose style, commit email, or a branch name.

**Stop with `Blocked` if the providers match, the authorship is unknown, or the evidence conflicts.** State the fact or the eligible reviewer that the review needs. Ask the owner to supply that fact, or to start a session of the opposite provider. Do not do a substitute review with another model of the same provider.

When both providers wrote substantive changes, neither one qualifies for the whole pull request. Record the conflict, and ask the owner how to divide the changes. Do not approve through a reciprocal review of selected hunks.

## Establish the review scope

- Read `CLAUDE.md` and the resume section of `docs/SESSION-HANDOFF.md`. Read each other document with a targeted read by id or heading (D-749).
- Load `.claude/skills/one-pr-one-session/SKILL.md` first. A review session works on one pull request alone (D-746).
- Load `.claude/skills/ste-writing/SKILL.md` before any review text (hard rule 2).
- Load `.claude/skills/mtg-corpus/SKILL.md` before you judge a claim about a format, a legality, or a card term.
- Read the pull request body, its roadmap item and gate, and each earlier review record.
- Read `docs/owner-questions.md` for an open choice that affects this change.
- Save every comment of the pull request to one file with one command, then read that file. See "Gitar and other comments".
- Record the pull request number, the target branch, the base commit, the merge base, and the head commit.
- Make sure that the local checkout and the diff match those commits.
- Keep unrelated local edits. Use a separate worktree when necessary.

Read the diff in stages (D-749). Run `git diff --stat` of the merge base and the head first. Then read the diff of each listed path: deleted files, renamed files, configuration, schemas, and tests. List each path in the review record as inspected, or name it as not inspected. Read each changed file in context, and follow the callers, the consumers, and the stored data beyond the diff.

The pull request body states the intent. The diff and the verified behavior show what the pull request does. A review of an uncommitted patch is provisional, and it cannot satisfy the gate. When the base or the head changes, assess the new diff and its evidence before a final verdict.

## Stay inside the pull request

A review judges the change in front of it. It does not design the next one. Read the roadmap item of this pull request and its gate before the first finding. Those two texts set the boundary. This section limits the reach of a review. It never lowers the standard for the code that the pull request changes.

A concern is in scope when one of these holds:

- The changed code gives a wrong result under a supported condition.
- The change breaks a caller, stored data, a deploy, or a build that exists today.
- The gate of this pull request does not hold.
- A contract that the pull request names does not hold for the code that it adds.

A concern belongs to a later pull request when one of these holds:

- It asks a tool of this pull request to cover a surface that no gate names.
- It asks for behavior that the roadmap gives to a later item.
- It repeats a class of defect in a surface that this pull request does not touch.
- It needs an owner decision about scope, and not a correction.

Write the second kind under `## Out of scope` in the review record, and name the item that holds it. Give it no severity. A line in that section never blocks the merge. A pull request that creates a check must pass that check.

## The review standard

Build your own account of the behavior before you compare it with the explanation of the author. For each changed behavior, trace the input, the state change, the output, the side effects, and the recovery path. State the invariant that each boundary must keep.

### Correctness and system effects

- Check normal use, boundary values, absent data, invalid data, repeated actions, and interrupted actions.
- Trace the owner and the lifetime of each state across the API, the worker, Firestore, GCS, and the web app.
- Inspect a retry, a timeout, a cancel, a cold start, and a restart when the change affects one.
- Check compatibility with current callers, stored documents, the proto contract, and deployed clients.
- Check whether a local fix breaks another consumer of the same contract.
- Verify each gate line against the implementation and the evidence.

Do not turn the review into an unrelated rewrite. Separate a defect of the pull request, a defect that it exposes, and an older defect. An older defect blocks this pull request only when it stops the changed behavior or a required check.

### Project contracts

Load `references/project-contracts.md` when the pull request changes code, tools, or CI, or the text of a contract.

### Design, maintainability, and documents

- Confirm one concern for each pull request, and a clear reason for each changed subsystem.
- Explain the concrete maintenance cost of a design objection.
- Do not report a personal style preference as a correctness defect.
- Check that the roadmap, the decisions, the questions, the code, and the gate agree.
- Separate proposed work, implemented work, measured behavior, and owner approval.
- Verify each MtG fact and each external claim against a dated primary source (hard rules 4 and 7).
- Check the documentation matrix against section 2 of the `one-pr-one-session` skill. A deferral to a later pull request is a finding (D-747).
- Check the attribution rule in commits, the pull request body, comments, and files (hard rule 6, D-806).
- Check that no file, issue, or pull request holds an email or a personal address (D-639).

A document or skill pull request needs the same provider independence and evidence as a code pull request. For a skill change, examine its trigger, scope, instructions, references, and behavior on a realistic request. A contradictory instruction and a gate that cannot pass are defects.

## Verification

Run the focused checks that can prove the changed behavior wrong, and the checks of the project. Use the commands of `CLAUDE.md` and `AGENTS.md`. Never run a paid target without the owner's word (D-65).

- Read the tests as critically as the implementation.
- Verify that each bug fix has a regression test that fails on the old code.
- When practical, run that test against the base in a separate worktree. Else, give the causal reason that the old code fails it.
- Check each assertion against the contract, and not against a copy of the implementation.
- Check test discovery, skipped tests, fakes, fixtures, and assertions that pass without the intended behavior.
- Separate a passed check from a skipped, unavailable, failed, or author-reported check.
- Record the command, the revision, the result, and the artifact of each required check.
- Verify the CI results against the reviewed commit. A docs-only skip of the code jobs is a pass only when `verify:gate` passed (D-805).

Do not repeat a broad suite with no new change, failure, or risk. Do not weaken a test or a threshold to get a pass. Absent required evidence blocks the approval. An optional gap goes under the limits of the record.

## Precise findings

Investigate each suspected defect before it becomes a finding. Search for a caller guarantee, a validation layer, a test, or a later decision that can refute the concern. Use a reproduction, a failed assertion, or a complete causal trace as evidence. Separate a verified defect from an open question or an optional suggestion.

Each finding holds:

- A stable id, a severity, and a short title that states the defect.
- The reviewed commit, and the smallest useful file and line range.
- The input or state that starts the defect.
- The expected behavior, with its contract or D-id.
- The actual behavior, and its effect on the user, the data, the deploy, the cost, or the maintainer.
- The evidence: the command, the trace, or the artifact.
- The direction of a correction, and the regression check that proves it.

Group repeated symptoms under one cause. Do not prescribe a broad rewrite when a smaller correction restores the contract. Do not invent findings to meet a quota. A thorough review can find nothing.

Answer two questions before a finding enters the record:

1. Does the changed code break a contract that this pull request names?
2. Does the gate of this pull request fail?

A finding needs one yes. A concern with two answers of no goes under `## Out of scope`.

| Severity | Meaning |
|---|---|
| P0 | An immediate critical failure, such as a loss of stored user data or a deploy that cannot start. State the shown scope. |
| P1 | A major correctness, recovery, security, cost, or required-check failure. Resolve it before the merge. |
| P2 | A concrete defect or contract gap under a supported condition. Resolve it before the merge, or get an owner disposition. |
| P3 | An optional improvement with no broken contract. It does not block the merge. |

Scope decides whether a concern enters the table. Severity decides how much it blocks. Do not lower a severity because the patch is small or the author calls the change safe. When two owner decisions conflict, quote both, file the question in `docs/owner-questions.md`, and stop the dependent work.

## Verdicts

| Verdict | Condition |
|---|---|
| Blocked | The provider gate, the review target, an owner decision, or required evidence is not settled. Record each verified defect too. |
| Changes required | The eligible review found defects or contract breaks in scope. List the required changes. |
| Ready for owner merge | The provider gate passes, the whole scope has review coverage, each required check passes, and no blocking finding stays open. |

A line under `## Out of scope` never gives the verdict `Changes required`. No finding does not mean no risk. State each material limit, and make no claim of zero regressions. The approval applies to the recorded effective head alone. The owner alone merges the pull request (D-583).

## Do not raise a tool name as attribution

Hard rule 6 and D-806 forbid text that names an agent, a harness, or a model as the source of the work. A tool name that identifies a file, a schema, or a verified version is not attribution.

| Raise it | Do not raise it |
|---|---|
| A commit body that says an agent wrote the change. | The path `.claude/settings.json`. |
| A co-author trailer or a generation line. | A decision that names the model a role runs on. |
| A pull request body that credits a model. | The `Author:` line of a hand-off entry, or a review record (D-806). |

## Gitar and other comments

The reviewer reads each comment of the pull request, and takes it into its own review. It never replies to Gitar, never resolves a thread, and never writes a comment on the pull request.

- A comment of Gitar is a claim about the code. Verify it against the head, and record the result under `## PR comments`.
- A reply of the author is evidence. Check its trigger, its contract, and the commit it names.
- While `.github/gitar-review` reads `paused`, no Gitar pass needs to exist (D-802).
- While it reads `on`, prove that the Gitar pass is current with the `gitar-review` skill. A Gitar comment with no answer blocks the verdict.
- A Gitar pass does not make Gitar an author. The provider gate reads the providers of the substantive commits alone.

## Scope limits

A review request permits these actions and no other:

- Inspection, verification, the review record, and a session entry of the hand-off.
- A commit of those files, and a push of that commit to the branch of the pull request.
- A correction of a stale fact in the pull request body, under "Correct the pull request body" in `references/review-record.md`.

It does not permit a code fix, a merge, or another external message. A reviewer never pushes to `main` (D-583). When the reviewer writes a substantive fix, check the provider gate again. The reviewer cannot approve its own change. Do not hide a fix as review metadata to pass the provider gate.
