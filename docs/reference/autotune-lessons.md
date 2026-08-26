# What the tuning loop learned

The loop appends one block per iteration (D-182). The fixer reads this file before it edits. A dropped hypothesis is a dead end, and a kept one is a base to build on. Nobody edits this file by hand except the owner.

## Before the file existed (2026-08-26, written by hand from the evidence)

Two fixers tried the same idea on two different days, and both were rejected. Neither could know about the other, because nothing recorded the first attempt. This file exists so that stops.

### Dropped: the power confirm row goes out as written

Hypothesis: the ask role drops the first clause of "You asked for a competitive deck, so I will build to tournament level. Is that right?", so a `fixed` flag keeps the sentence whole.

Tried as D-162 (run 20260826-153915, orphaned) and again as D-179 (run 20260826-191225-001, rejected). In the second run the fixed text was judged bad in conversations 33, 49, and 68, all of which were fine before. Reason on conversation 33 ("A Modern deck for an event"): the question asserts a premise the user never gave. "Assumes an answer" faults rose from 9 to 15.

The first clause is the defect, not the rewrite. Do not fix the wording. Fix the trigger, so the row fires only when the user said competitive, strongest, or best, and never on "for an event".

### Kept in spirit, never landed: the house-limits row goes out as written

Hypothesis: the ask role splits "60 cards, four copies per name, and a 15-card sideboard" into three questions, so `fixed` keeps it as one.

In run 20260826-191225-001 both "two questions in one" faults on `house_format_limits` disappeared. The change was dropped with the rest of its iteration, because the loop was all-or-nothing then. It is a candidate to try again on its own.

### Kept in spirit, never landed: a partner phrase names the Commander format

In the same run, conversation 78 ("as partners") no longer got the format question. Same story: dropped with its iteration. A candidate to try again on its own.

### Noise, not a lesson: conversation 88 ("Surprise me")

Three questions of conversation 88 were judged bad in run -001 and good in run -000 and run 18, with identical text. That is the judge, not the fixer. The paired comparison now ignores such flips.

### The judge's own noise, measured

The same gate document scored twice by `gpt-5.6-luna` differs on 25 of 365 verdicts (7 percent). The holdout moved from 10 to 12 bad questions with no code change. Two full runs of identical code differed by three holdout questions. The checker's margin is three bad questions (D-183). A change that moves the holdout by less than that has proved nothing either way. That is why the paired comparison on changed text decides.

## 20260826-212512-000, baseline failed the gate (2026-08-26, written by hand)

The first baseline after the audit fixes failed its own gate: 24 catalog-only conversations against a bar of 25, and one premature session. Two lessons.

### Dropped: the not-owned commander row fires before PR-8

The row `commander_not_owned` was dead code until the key fix of D-197. It then fired in 14 conversations, and the score role rated it 0.05 every time. 13 of them went out as an invented replacement, and the eval refused 10 of the 14. That one row held 10 of the 42 bad questions of the run.

The owner ruled it dormant until PR-8 (D-207, OQ-36 stays open). Do not make this row fire, and do not reword it. It has no owned options to name before the generator exists.

### Confirmed: the budget row fires for every buy list

D-168 raised the budget row from 6 to 83 conversations, and the eval refused 0 of the 83. The rule holds (D-206). Do not narrow its trigger to make the count fall.

### Fixed by hand: conversation 39 was premature

"Make me a good deck", then "you pick", then "Whatever you think is best" ended the session with no commander. Run 18 did not do this on the same messages. The cause was a hole in the decline path and not a catalog row. A declined pick row closed its own key and not the commander slot (D-208). It is fixed in code. No catalog change can touch it.
