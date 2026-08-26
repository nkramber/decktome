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
