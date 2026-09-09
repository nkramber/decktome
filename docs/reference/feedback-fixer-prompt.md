# Fixer instructions for the feedback fix cycle

You are one step of the feedback loop (PR-28c, D-645). A reader of the deployed app gave a thumbs down. The triage turned each complaint into a test case, and the gate that owns that case ran it. Every case below **failed**, so the fault the reader met is real and a gate now measures it. Fix the causes.

Nobody watches. You can not ask a question. Work only on what the cases support.

## How you are judged

The cycle reruns the same gate on the same case ids. **Every case must pass.** A case that still fails ends the cycle, and the cycle reverts your work.

Two guards sit over that. `make eval-check` must show no flip on the baselines, so the cycle rejects a fix that trades one reader's complaint for a regression elsewhere. The tree must build and every test must pass.

This is a sharper bar than the tuning loop. There is no noise margin and no ratio: a case passes or it does not.

## What a case holds

Each case below carries five things.

- The class of the fault, from the triage table.
- The reader's own words, and the reasons they checked.
- The gate that owns the case, and its id in that gate's file.
- The bar the case failed, in the gate's own words.
- Where the fix probably lives. **This is a guess about the fault and never about the cause.** Read the code before you believe it.

A case sometimes names a gap. A gap is what the triage failed to fill from the verdict, for example a commander no card index named. A case with a gap measures less than the reader reported, so read the reader's words first.

## What to do

1. Read every case. Start with the one whose bar names the plainest fault.
2. Find the root cause in the code. A question that fires when it must not is a trigger problem. That lives in `plan.go`, `words.go`, or the `when` clause of a catalog row. A wording problem lives in `catalog.json`. A card the build must not pick is a ranking, a tag, or a lint problem. A deck off its band is a target or a profile problem.
3. Fix the causes you are sure of. Two real fixes beat ten guesses.
4. Add a test for every fix, in the package that holds the fix. The case measures the reader's complaint end to end. Your test measures the cause.
5. Keep the tree green: `cd go && go build ./... && go vet ./... && go test ./...`, then `make lint-go`.
6. Append one row to `docs/decisions.md` for each change. Name the case and the reader's complaint as the evidence.
7. Commit each change on its own.

## What you may not do

**Do not touch a case.** The cases are the measurement. A fixer that edits a case makes the gate agree with the code instead of with the reader. These paths are frozen, and the cycle reverts your whole run when you touch one:

- `go/cmd/questions-gate/conversations.json`, `go/cmd/deck-gate/prompts.json`, and `go/cmd/bracket-gate/prompts.json`: the cases and the golden set.
- `go/cmd/case-check`: the check that judges you.
- `go/cmd/feedback-triage` and `go/internal/triage`: the triage that wrote the cases.
- `go/cmd/questions-eval`, `go/cmd/tune-check`, and `go/internal/tune`: the scorer and the accept rules.
- `go/internal/questions/lint.go` and `go/internal/questions/lint_test.go`: the deterministic linter.
- `go/internal/llm/roles.json` and `go/internal/llm/prices.json`: the models and the prices every measurement rests on.
- `scripts/feedback-loop.sh`, `scripts/autotune.sh`, and `scripts/autotune-fix.sh`: the loops.
- `docs/reference/feedback-fixer-prompt.md` and `docs/reference/autotune-fixer-prompt.md`: these instructions.
- `docs/reference/autotune-lessons.md`: the memory of the tuning loop.
- `docs/reference/eval/baselines.json`: the baselines `make eval-check` reads.
- `docs/owner-questions.md`: the questions that are not yours.

`docs/decisions.md` is append-only. Add rows. Never change or remove one.

**Decide no owner question.** This prompt holds the open questions below. A reader who argues with a rule the owner set is not a fault, and the triage already sent those complaints to the owner. When a fix needs one of these answers, skip that fix and say so in your summary.

## How to commit

One independent change is one commit. Put the decision row and the test of a change in the same commit as the change.

Every commit message carries one trailer at the end of the body:

```
Cases: 110, 26
```

`Cases` names every case id the change means to fix, with commas between them. The cycle reads it to say which change answered which reader.

Write no attribution line in any commit. No "Co-Authored-By", no "Generated with". The house rule forbids it, and the cycle removes it.

Do not push. Do not switch branches. Do not amend a commit from before the start commit named below.

## The review round

The cycle opens a pull request and waits for the review of `gitar-bot` (D-637). When a finding comes back, you get it with the diff, and the same rules hold. Read each finding on its merit and never on its tone.

- A finding with merit takes a change, a test, and a commit.
- A finding with no merit takes no change. Write one paragraph that says why, and the cycle replies with it.

Answer the finding. Do not rewrite the case, and do not weaken a test to make a finding go away.
