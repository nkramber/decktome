# decktome - CLAUDE.md

Read this file first. Then read `docs/SESSION-HANDOFF.md`. It tells you where the last session stopped. For all work on a pull request, load `.claude/skills/one-pr-one-session/SKILL.md` first (hard rule 12).

## Project

This repo is a Go + Protobuf + TypeScript monorepo for an agentic MtG deck builder. The app reads a user's collection export from ManaBox or Moxfield. The user gives a prompt. The agent asks questions, then builds a legal, useful deck.

Stage: **the app is live on `decktome.com`**, and Phase 3B, the product UI, is the current phase. A merge to `main` deploys itself on Cloud Build (D-584, D-586, D-588). `docs/SESSION-HANDOFF.md` holds the current state and the next step. Git holds each merge, and Cloud Build holds each deploy (D-747).

Two rules of the deployed app come from 2026-09-09. **A verdict keeps the object it names** (D-635), because a reader deletes the deck or the chat they complained about, and the verdict outlives it. **Every user has a record** at `users/<uid>` (D-638): the verified email, the dates, and six counters of what they made. No harvest reads that record, and a test refuses the import that joins them.

The look follows a reference design the owner gave on 2026-08-30 (D-328 to D-335). `docs/reference/autotune-readme.md` holds the loop commands.

**The repository is public** (D-639). Write no email and no personal address into a file, an issue, or a pull request. D-642 permits a reader's own words. Every pull request runs the verify workflow, and a change of documents alone skips its heavy jobs (D-818).

Run `make where` before you change anything. It prints the branch, the tree, and the state of the branch's pull request. The session commits on a branch, pushes it, and opens a pull request. The owner merges (hard rule 8, D-583, D-585). Never commit on `main`, and run `make hooks` once in a fresh checkout.

Read `docs/SESSION-HANDOFF.md` next. It is the resume point.

## Hard rules from the owner

1. **Write scope.** This repo permits writes (D-32, 2026-08-23). All other repos are read-only. Application code follows the roadmap order. Do not start a roadmap item before the gate of the item before it holds.
2. **Write in ASD-STE100.** Every doc, skill, and agent file must follow Simplified Technical English. Load the `ste-writing` skill before you write. Rules that apply most: max 20 words per procedural sentence, max 25 per descriptive sentence, and active voice. Also: one instruction per sentence, no semicolons, no "-ing" verb forms, one term per concept, paragraphs of max six sentences.
3. **Ask questions when you think of them.** Do not save questions for the end. Use `AskUserQuestion` in small batches. Record each answer in `docs/decisions.md`.
4. **Do the research.** Verify facts against sources (Scryfall API, Wizards announcements, the Comprehensive Rules). Record the date of each fact. MtG rules and ban lists change often.
5. **Make hand-off simple.** Update `docs/SESSION-HANDOFF.md` inside the pull request, before you call it ready (D-747). Record the completed work, the open work, and the next step.
6. **No AI-attribution text** in any PR, branch name, commit message, or comment. This house rule comes from connector-syncer. A review record and the `Author provider` line of the hand-off can name a provider (D-811).
7. **No mistakes.** Check card names, rules, and dates before you write them. When you are not sure, say so and mark the item as unverified.
8. **Every change starts on a branch.** Never commit to `main`, and never push to it (D-583). Make a branch, commit there, push it, and open a pull request. The owner merges. Run `make where` before every commit, push, and deploy. It prints the branch, the tree, and whether `main` is current. It also names the state of the branch's pull request. Run `make hooks` one time, and the pre-commit hook then refuses what these rules forbid (D-585).
9. **Deploy from `main` alone.** Never deploy any other branch to production, for any reason (D-579). Check the branch and the commit before every build, not only the tree. `docs/deploy-and-rollback.md` holds the procedure.
10. **Answer the review before you ask for a merge.** `gitar-bot` reviews every pull request (D-637). Load the `gitar-review` skill after each push, and follow its procedure (D-745). The skill proves that a review is current, and it answers each finding. These rules of this repo win over the skill:
   - **After Gitar, a Codex session reviews the pull request** with the `pr-review` skill (D-811). The `review-gate` check reads its record, and `main` requires the check (D-815).
   - A pull request of documents alone waits for Gitar too (D-679). When it is ready and each other check is green, apply the `review-override` label in place of the Codex review (D-812).
   - Tell the owner when the pull request is ready to merge. The owner merges.
   - Wait for a current Gitar review before you call the pull request ready. Fix each finding on the same pull request, in the same session (D-746).
   - A commit of `docs/SESSION-HANDOFF.md` or the hand-off archive alone does not make a Gitar pass stale (D-752).
11. **Never hesitate to ask or to push back.** Ask a question the moment you have one. When the owner's two statements conflict, say so and quote both. When a request rests on a wrong premise, say so with the evidence. The owner sees this as the key to good LLM-user interaction. Silence is the mistake, not the question.
12. **One pull request, one clean session.** A session works on one pull request, and the pull request carries all its documents and its hand-off. No pull request exists to record an earlier merge. Load `.claude/skills/one-pr-one-session/SKILL.md` for all work on a pull request (D-746 to D-748). The owner says that the pull request merged. Then write the transitional prompt of section 5, and do no other work (D-754, D-764).
13. **Keep command output small** (D-749). Every line of output stays in the context of every later call. Count or list the matches first, with `grep -c` or `grep -l`. Then read a bounded range, with `sed -n`, `head`, or the offset and limit of the Read tool. Read only the section that you need. Do not print a whole document. Show the output of a failed test, build, or gate in full, because the error is the evidence.

## Reference material

- Style model for the design doc: `/Users/nate/Repos/connector-syncer-docs/docs/document-summary-roadmap.md`. Load the `design-doc-style` skill for the section template.
- AI patterns to borrow: `/Users/nate/Repos/connector-syncer` (Python RAG app). See `docs/reference/connector-syncer-ai-patterns.md`.
- Frontend patterns to borrow: `/Users/nate/Repos/wallabee-ui`. See `docs/reference/wallabee-toolchain-conventions.md`.
- MtG terminology and rules corpus: load the `mtg-corpus` skill. It is the app's future knowledge base.

## Skills in this repo

| Skill | Use when |
|---|---|
| `ste-writing` | Before you write or edit any `.md` file here. |
| `one-pr-one-session` | Before any work on a pull request: a start, a revision, a review, a merge message, or the hand-off. |
| `gitar-review` | After each push to a pull request, documents alone included. |
| `pr-review` | For a Codex review, for each answer to a review, and before you apply the `review-override` label. |
| `design-doc-style` | Before you edit `docs/design-roadmap.md`. |
| `mtg-corpus` | Before you reason about formats, legality, archetypes, or card terms. |

## File map

- `docs/design-roadmap.md` - the design document (the deliverable).
- `docs/decisions.md` - every owner decision, with date.
- `docs/SESSION-HANDOFF.md` - resume point for a fresh session.
- `docs/setup-second-mac.md` - what to carry and what to install to continue the work on another Mac.
- `docs/setup-gcp.md` - the deploy on Google Cloud from the domain to the running app, with the cost estimate.
- `docs/deploy-and-rollback.md` - the redeploy after a merge to `main`, and the rollback of each part.
- `docs/reference/mobile-and-engagement-2026-09-05.md` - the phone and engagement proposal in three stages.
- `docs/open-questions.md` - questions not yet asked or not yet answered.
- `docs/owner-questions.md` - the decision queue. Every question here waits for the owner, and the tuning loop refuses to decide one.
- `docs/reference/` - research notes with sources and dates, and every dated gate document.
- `docs/audit-2026-08-28.md` - the full audit of 2026-08-28, its owner answers, and the change plan.
- `docs/audit-2026-08-29.md` - the quality audit of 2026-08-29 and its fixes (D-302 to D-306).
- `docs/reference/set-data-2026-08-31.md` - every set number PR-17B rests on, with its source and date.
- `docs/reference/bracket-profile-2026-09-02.md` - every bracket rule, Karsten table, and Spellbook threshold PR-14A rests on.
- `docs/reference/deck-quality-model-2026-09-02.md` - every source fact, feature, and fit rule PR-14B rests on.

## Commands that cost money

Thirteen targets and two loop scripts spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `make api-build`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before every run. A rerun writes to a new `*_OUT` file, and a guard refuses a document that holds a result (D-65).

`docs/reference/paid-targets.md` holds the cost, the flags, and the guards of each paid target, and every free target (D-749). Read it before you run or change a target.

`make verify` runs every check the verify workflow runs, on this machine, for nothing (D-578). Run it before every pull request. `make lint` also runs `make ste-check`, `make ref-check`, `make lifecycle-check`, `make context-budget`, and `make pipefail-check` (F-160). `make pr-check` reads the pull request body and diff against the contract of D-747. `make where` prints the branch, the tree, and the state of the branch's pull request. `make hooks` installs the pre-commit hook that refuses a commit on `main` (D-585).
