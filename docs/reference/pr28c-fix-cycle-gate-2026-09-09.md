# PR-28c gate, 2026-09-09

The roadmap gate of PR-28c has one item.

> One live cycle on the owner's word ends with a pull request and a passing gate on its new cases. `make eval-check` shows no flip on the baselines.

**Verdict: the free half holds. The live cycle waits for the owner's word.**

A live cycle spends money and it edits code with nobody watching. It also needs a harvest that carries a snapshot, and the three verdicts on record predate D-635. So this document records what runs for nothing, what a live cycle would do, and what is left to measure.

## What holds, for nothing

### The plan

`make feedback-loop-dry LOOP_ARGS="--in go/internal/triage/testdata/fixture.jsonl"` reads the ten fixture verdicts and names the whole cycle. It calls no model, it makes no branch, it commits nothing, and it writes no case into a gate file.

```
dry run: no branch, no model call, no commit. The checkout stays on pr-28c.
cap $2.00. State under .local/tune/feedback-20260909-212326.
step 1: the triage, dry. The reason keys place a case, no judge runs, and no gate file changes.
the triage wrote 6 case(s) a gate measures
step 2: confirm. Every case must fail before the fixer touches anything.
dry run: no gate ran. A live cycle would run: questions decks brackets
  questions gate, -only 110,111
  decks gate, -only 26,27,28
  brackets gate, -only 16
```

`git status` after the run reports no change to `conversations.json`, `prompts.json`, or the bracket prompts.

### The manifest

The manifest is the hand-off from the triage of PR-28b to the cycle. It names the gate that owns each case, the id in that gate's file, where the fix probably lives, and the reader's own words.

```json
{
  "class": "Q1",
  "from": "f01",
  "gate": "questions",
  "target": "go/cmd/questions-gate/conversations.json",
  "id": 110,
  "where": "the catalog trigger, the word rules, the classify prompt",
  "said": "Reasons: already_answered. The reader wrote: I already told you the bracket in my first message.",
  "gaps": ["the commander is an Oracle id and no card index named it"]
}
```

`TestTheManifestHandsTheCasesToTheCycle` reads the same fixture and pins the counts: six cases a gate measures, two on the question gate, three on the deck gate, and one on the bracket gate. Four of the ten verdicts reach no gate, because two want the judge and two write a defect row.

### The verdict reader

`cmd/case-check` reads the run file the gate wrote and says whether each case passed. It reads the same bars the gate's own verdict reads: a metric the run marks lower-is-better is good at zero, and every other gate bar is good at one.

Six tests fire both sides of it.

| Test | What it holds |
|---|---|
| `TestAFailingCaseReadsFail` | The confirm run. The fault is real, and the same run after a fix is a failure of the cycle. |
| `TestAPassingCaseReadsPass` | The measure run, and a case that already passed on the confirm run. |
| `TestALowerIsBetterBarReadsTheOtherWay` | A lint finding is good at zero, and a slot bar is good at one. |
| `TestARunThatMeasuredAnotherItemIsAFault` | A run of the wrong ids is a fault and never a failing case (T-12). |
| `TestAGateWithNoCaseIsAFault` | A pass over an empty list is a fault. |
| `TestTheGateFilterReadsOneSuite` | A deck case never counts against a question run. |

### The guards

| Test | What it holds |
|---|---|
| `TestEveryFrozenPathExists` | A frozen path that moved guards a file nobody has. |
| `TestTheCaseFilesAreFrozen` | The three gate files, the triage, and `case-check` are frozen. A fixer that edits a case makes the gate agree with the code instead of with the reader. |
| `TestTheFixerPromptAndTheCycleAgreeOnTheFrozenList` | The prompt warns about every path the cycle enforces. It caught nine paths the prompt had missed. |
| `TestTheCycleNeedsItsOwnPermission` | `FEEDBACK_LOOP_ALLOW`, `AUTOTUNE_FIXER_CMD`, and the $2 cap. A dry run reaches none of them. |

`shellcheck scripts/*.sh` reads clean, and `make verify` passes in full.

## What a live cycle does

1. **The cases.** The triage runs with `-apply`, so each case joins the gate file that owns it. One commit.
2. **Confirm.** Each gate runs over its case ids alone, and every case must FAIL. A case that already passes stays as a case a change must not flip, and the fixer never sees it.
3. **The fixer.** `scripts/autotune-fix.sh` runs the agent of `AUTOTUNE_FIXER_CMD` with `docs/reference/feedback-fixer-prompt.md`, the failing cases, and the reader's words.
4. **The free guards.** A frozen path, a removed line of `docs/decisions.md`, or a red tree reverts the fixer and keeps the cases.
5. **Measure.** The same gates run over the same ids, and every case must PASS.
6. **The baselines.** `make eval-check` must show no flip.
7. **The pull request.** The cycle commits the evidence, pushes, and opens the pull request with the measure run in its body.
8. **The review.** `scripts/feedback-review.sh` waits for `gitar-bot`, hands every open finding to the fixer, runs the free checks, pushes, and replies on each thread. Three rounds at most.

Every failure between step 3 and step 6 reverts the fixer and keeps the cases, so the branch never holds a case with no fix and no evidence.

## The cost

The ledger reads `cost_usd` from the header of each run file, and the cycle checks the cap before every paid run. A run with no priced cost is a fault, because an uncounted call would let the cycle pass its cap without knowing.

The cap is $2.00 (D-559). On the fixture the cycle would run six cases over three gates, twice. A question-gate case costs about a cent, a deck-gate case about $0.13, and a bracket-gate case about $0.14 with its judge call. So the fixture plan reads about $1.10 for both passes, inside the cap.

The fixer's own tokens are outside this ledger. They bill against the owner's monthly plan (D-159).

## What this gate does not measure

**The live cycle.** No run of it exists. It needs three things: the owner's word, `AUTOTUNE_FIXER_CMD` set to an agent, and a harvest whose verdicts carry a snapshot. The three verdicts on record predate D-635, so the first live cycle belongs with the first real harvest after it.

**The review round.** `scripts/feedback-review.sh` reads clean and its shape is pinned, and no run of it exists either. It writes to a public pull request, so the first run belongs with the first live cycle, on the owner's word.

**The judge lane of the triage.** Still unmeasured, as PR-28b recorded. The cycle passes `-dry` to the triage on a dry run, so the free lane never reaches it.
