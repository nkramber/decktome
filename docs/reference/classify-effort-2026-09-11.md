# Classify effort, none against low, 2026-09-11

This note records the measurement of D-672. It asks whether the classify role misreads less at effort `low` than at effort `none`. Every number here comes from partial runs of the question gate.

## The method

- The code is `193618a`: the offer rule of D-669, and not the guards of D-670. So a misread still shows as an expectation miss.
- The runs read conversation 100, "terse: a deck for a team event", and conversation 103, "terse: two themes at once". Gate run 46 missed the first, and gate run 47 missed the second.
- Each round runs one partial run at effort `none` and one at effort `low`, through `LLM_CLASSIFY_EFFORT`. Eight rounds ran, from 03:00 to 03:08 UTC.
- The classify role runs on `gpt-5.6-luna`. The client sets no temperature and no seed.

## The result

| Effort | Runs | Conversations met | Missed | Cost per run | Time per run | Calls per run | Output tokens per run |
|---|---|---|---|---|---|---|---|
| `none` | 8 | 16 of 16 | 0 | $0.0022 | 20 s | 10.8 | 1,117 |
| `low` | 8 | 16 of 16 | 0 | $0.0033 | 27 s | 11.5 | 1,802 |

The measurement cost $0.0439 in total.

## What the result says

- **No run missed at either effort.** Each conversation ran 8 times at each effort, and every expectation held. A recent whole run misses one of its 74 conversations, so a single conversation misses rarely. Sixteen runs can not tell the two efforts apart.
- **`low` costs more on every turn.** It costs 1.5 times as much per run, takes 35 percent longer, and writes 61 percent more output tokens.
- **The guards of D-670 close both gaps whatever the classifier reads.** The measurement gives no reason to raise the effort, so the classify role keeps effort `none`.
