# The on-band plan, 2026-09-08

The owner read a build that took four minutes and gave the reader
nothing. The cause is not one slow call. Almost every deck leaves the
first model call off its bracket band. The app then spends two more
calls of 60 to 90 seconds. They move a number the code moves in 10
milliseconds.

This document holds the read, the four causes, and the plan. The goal is
one model call that lands in band.

## What the reader met

Session `833r7UccvAqFsyYJzHfz`, 2026-09-08. The reader asked for a
Commander deck of the Hobbit and Lord of the Rings sets. Aragorn leads
it, at bracket 4, from the cards they own.

| Time | What ran | Result |
|---|---|---|
| 14:43:53 | the build starts | the shortlist holds 657 cards of the named sets |
| 14:44:56 | the generate call ends, 63 s | 0 missed names, 2 findings |
| 14:46:24 | repair turn 1 ends, 88 s | 0 missed names, 2 findings |
| 14:47:52 | repair turn 2 is cut | the 4-minute build cap ends the build |

The reader waited 3 minutes 59 seconds and read an error. **A legal deck
existed at 14:46:24.** The second repair turn only runs when every
finding is a profile finding, and every profile finding is a warning. So
the deck of 14:46:24 held no block, and the build threw it away.

Session `CVN9U3QsXSbOLeCAg8JL` failed the same way on 2026-09-07, at
18:34:53 to 18:38:53. Its repair turn 1 read one finding and its repair
turn 2 read two, so that repair made the deck worse.

## What the measurement says

The two newest gate documents hold 40 built decks: deck gate run 16 (25
decks) and bracket gate run 1 (15 decks). They hold 26 off-band findings.

| Count | Feature | Does the prompt name it? |
|---|---|---|
| 8 | mana available on turn four | no |
| 5 | average mana value | yes |
| 4 | the color sources the deck needs | no |
| 2 | removal count | as a midpoint |
| 1 | land count | as a midpoint |
| 1 | ramp count | as a midpoint |
| 1 | draw count | as a midpoint |
| 1 | interaction count | as a midpoint |
| 3 | other | mixed |

**12 of the 26 findings are on two features the model never reads.**
Every one of the five average-mana-value findings sits above the
ceiling, at bracket 4 and bracket 5.

## The four causes

**Cause 1. Four bands never reach the model.** `profile.promptKeys`
names five features: the average mana value, the tapped lands, the
colorless lands, the tutors, and the fast mana. `bands.json` checks
fifteen. The prompt states none of `color_sources`, `mana_turn_four`,
`hands_two_to_four_lands`, and `commander_turn_over_mv`. The check reads
all four. So the model answers numbers it never sees.

**Cause 2. The model reads a point where the band is a range.**
`TargetsFor` sends the midpoint of each band, so the job block reads
"land: 33". The band is 31 to 36. A deck of 30 lands is one card from
the target. It is also one card outside the band, and the model can not
tell the two apart. Six findings are role counts.

**Cause 3. The shortlist ranking ignores the bracket's curve.**
Bracket 4 wants an average mana value of 1.6 to 3.0, and bracket 5 wants
1.2 to 2.4. The shortlist ranks on theme fit and popularity (D-94). So a
high bracket needs cheap cards that sit low in the list the model reads.
Five findings are that feature.

**Cause 4. The fix runs in the model and not in the code.** Every band
the profiler measures is a number this app computes itself. The goldfish
simulation of 10,000 hands takes about 10 milliseconds. The app spends a
60 to 90 second model call instead, twice, and the numbers do not move:
session `833r7` read two findings before the repair and two after.

## The plan

Four parts, in this order. Every part is free.

### Part 1. State every band the check reads

`promptKeys` grows to every feature that carries a band. The four silent
ones become instructions the model can act on:

- `color_sources` becomes a source count per color. `SourcesNeeded`
  computes it already, and the deck shape block states it.
- `mana_turn_four` becomes a floor on the lands and the cheap ramp
  together, derived from the band.
- `hands_two_to_four_lands` becomes a land-count floor.
- `commander_turn_over_mv` becomes a ramp instruction for a costly
  commander.

Gate: a test walks `bands.json` and fails when a feature with a band
reaches no prompt line.

### Part 2. Send the band, not the midpoint

The job target block reads "land: 31 to 36, aim for 33". The model then
knows which deviations cost nothing.

Gate: every job line names its band, and a test reads the two together.

### Part 3. The deterministic mana pass

After `assemble`, and before any repair call, the builder fixes the mana
base with no model call. While a mana feature sits off band it makes one
step, then measures again:

- swap a tapped land for an untapped land of the shortlist,
- move a basic land to the color that lacks sources,
- move one card between the land count and the spell count, inside the
  land band.

Each measurement costs about 10 milliseconds, so the whole pass costs
less than a second. It stops when every mana feature is in band, or when
no step helps.

Gate: the pass runs over the stored decks of deck gate run 16 and
bracket gate run 1. It puts every `mana_turn_four` and `color_sources`
finding in band, or it names the deck it can not fix and why. The lane
reads stored decks, so it is free.

### Part 4. The loop keeps what it has, and it watches the clock

- A repair call that fails leaves the deck before it standing. The
  builder does this for a repair that answers a worse deck already
  (`worseRepair`), and a failed call must follow the same rule.
- The loop starts a repair turn only when the build has more time left
  than the last call took.
- A deck whose findings are profile findings alone is a legal deck. It
  goes to the reader after Part 3, and no repair call runs for it.

Gate: a build that loses its deadline returns the last legal deck, and a
test proves it. A build with profile findings alone makes one model call.

## What the reader gets

| | Now | After |
|---|---|---|
| Model calls of a build | 1 to 3 | 1 |
| Time to the deck | 60 to 240 seconds | about 60 seconds |
| A build that ends with no deck | seen twice in two days | none |
| Decks off band | 26 findings over 40 decks | the gate of Part 3 |

## The effort measurement, run 2026-09-08

The generate call takes 48 to 91 seconds: `gpt-5.6-terra` at medium
effort, with a 16,384-token output cap. The owner asked for a
measurement of low effort against medium (D-610), and it ran first
(D-611). Deck gate run 17 changed one field of run 16,
`LLM_GENERATE_EFFORT=low`, on the pinned card snapshot of run 16.

| | Run 16, medium | Run 17, low |
|---|---|---|
| Verdict | PASS | **FAIL** |
| Wall clock | 3018 s, 120.7 s a prompt | 2823 s, 112.9 s a prompt |
| Cost | $3.79 | $3.33 |
| Provider calls | 93 | 93 |
| Decks through a repair turn | 12 of 25 | 12 of 25 |
| Warnings | 37 | 31 |
| Block findings | 0 | 1 |

Low effort saves 7.8 seconds of the 120.7 a prompt takes. It also
answers prompt 11 with a Commander deck of two copies of Skullport
Merchant, and the copy limit of the format is one. The repair turn of
that deck ran for a missed name and left the duplicate in place. The
block bar is zero tolerance, so the run fails. A shortlist can not make
a model list one card twice, so the fault is the call. **The generate
role stays at medium** (D-612).

The measurement also says the latency is not in the call. A build takes
four minutes because it makes three calls. This plan removes two of
them.

CAUTION: the epoch moved on two more axes, and the snapshot pin holds neither. The
plan_rubric prompt moved from version 1 to version 2, and the quality
model moved from `20260903T210658Z` to `20260907T081149Z`. So the
`grade` rows and the `plan_score` rows of the two runs do not compare.
The wall clock, the cost, the call count, the repair count, and the
block do compare, because no version of either one moves them.
