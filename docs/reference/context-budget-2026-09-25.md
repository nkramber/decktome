# Context budget of five sessions on the new files (2026-09-25)

This document holds the measurement that D-750 asked for. It repeats the method of `docs/reference/context-budget-2026-09-16.md` on five sessions that ran after #184 and after D-749. It then replays the checkpoint rule on the same sessions.

## Scope and method

The measurement read the local Claude Code transcripts of this repo. It read the usage fields of each API call, the tool names, the file paths, the command kinds, and the size of each tool result. It read no text of a prompt, and this document quotes none. It names each session by its pull request, and it writes no session id.

The sample is the five newest sessions with 40 or more API calls and one or more prompts of the owner. They are the author sessions of #227 to #231, and they ran from 2026-09-25 03:30 UTC to 2026-09-25 19:50 UTC. Each one worked on one pull request, and each one ended at the merge. The session of this measurement stays out.

CAUTION: The five sessions are not typical. The owner asked for as many corrections as one session can finish (D-944), so each session ran long. A checkpoint rule acts on long sessions alone, so the sample shows its effect where it acts. The cross-check below reads all 47 sessions of the same kind since D-749.

Each number has a kind, as in the first audit:

- **Reported:** the usage fields of the transcript.
- **Exact:** a count of bytes, calls, or results.
- **Estimate:** a calculation from the numbers above.

Each token estimate divides characters by 2.43, the ratio of the first audit.

## The five sessions (reported)

| Measure | Before D-749: ten sessions | After D-749: five sessions |
|---|---|---|
| API calls | 1,791 main, 67 subagents | 1,624 main, 175 subagents |
| Input tokens, not cached | 36,906 | 3,330 |
| Cache read tokens | 772,099,913 | 450,690,161 |
| Cache write tokens, one-hour lifetime | 13,397,323 | 2,016,776 |
| Output tokens | 4,249,090 | 947,587 |
| Thinking tokens, part of the output | 2,647,911 | 271,349 |
| Input tokens for each session | median 85.4M, maximum 141.3M, minimum 17.3M | median 70.5M, maximum 187.8M, minimum 40.4M |
| Context of one call | median 396,862, maximum 951,409 | median 258,598, maximum 637,953 |
| Context of the first call | 43,953 to 50,187 | 32,719 to 44,339 |
| Compactions | 1 | 0 |
| Sessions of more than 200 calls | 5 | 5 |
| Sessions past 300K of context | not recorded | 4 |
| Reasoning effort of each call | not recorded | high |

Cache reads made 99.6 percent of the input tokens. So the cost of a session still follows the size of its context times the number of its calls.

## Where the context came from (estimate)

The model explains 408.2M of the 452.7M reported input tokens. The rest holds the prompts, the system reminders, and the text of each skill load.

| Source | Before D-749 | After D-749 |
|---|---|---|
| Model output, thinking included | 53.0% | 42.2% |
| The first-call context of the harness and `CLAUDE.md` | 12.0% | 16.1% |
| Search commands | 6.3% | 14.6% |
| Read of code files | 7.2% | 0.1% |
| Shell reads such as `sed` and `cat` | 3.9% | 13.9% |
| The hand-off, all reads | 5.5% | 3.2% |
| The roadmap, all reads | 2.2% | 1.1% |
| `git diff` and `git log` | 1.6% | 1.7% |
| Tests, `make verify`, and GitHub commands | about 2% | 2.5% |

The Bash tool made 1,515 of the 1,615 tool calls. So a read of a code file went through `sed` or `grep`, and the Read share moved to the shell rows. The three read rows together grew from 17.4 to 28.6 percent. A long session carries each result over more calls, so the share grows with the session length.

The hand-off target held. The five sessions made one read of the hand-off of more than 15,000 characters. The ten sessions of the first audit made 13 full reads of 38 to 53 KB. The largest tool result held 21,303 characters, a shell read. The hand-off holds 23,230 bytes on the base of this pull request.

## The checkpoint replay (estimate)

The replay walks the calls of each session. When the context of a call passes 300K tokens, a new session starts on the same pull request. That session holds the first-call context plus a re-read cost R. Each later call keeps its reported growth over that start. The replay counts the start of each new session as a cache write.

R is the cost of the reads that a new session makes again. It holds the hand-off, the skills, the diff, and the state of the pull request. The five sessions grew by these amounts after their first call:

| Calls after the first call | Growth, lowest to highest |
|---|---|
| 10 | 22,494 to 47,168 |
| 20 | 35,844 to 62,985 |
| 30 | 49,226 to 76,974 |

| R | Input tokens saved | Weighted cost saved |
|---|---|---|
| 0, the high bound | 39.9% | 32.8% |
| 25,000 | 37.1% | 30.0% |
| 44,000, the expected value | 34.8% | 27.8% |
| 55,000 | 33.5% | 26.4% |

The limit of 300K comes from D-750. The replay also read two other limits, with R at 44,000:

| Limit | Five sessions: new sessions, input saved, weighted saved | All 47: new sessions, input saved, weighted saved |
|---|---|---|
| 200K | 12, 51.7%, 39.9% | 42, 34.7%, 22.6% |
| 300K | 5, 34.8%, 27.8% | 13, 18.0%, 12.3% |
| 400K | 2, 22.7%, 18.5% | 5, 9.2%, 6.5% |

The weights are 0.1 for a cache read, 2 for a one-hour cache write, and 5 for an output token, as in the first audit. The replay keeps the output tokens of each call. A new session also writes an output to orient itself, so each saving is a high value for its R.

## The cross-check: all 47 sessions since D-749 (estimate)

The same filter over the whole period since the merge of #186 finds 47 sessions. They ran from 2026-09-18 to 2026-09-25.

| Measure | Value |
|---|---|
| API calls | 8,028 main, 1,918 subagents |
| Context of one call | median 164,681 |
| Sessions of more than 200 calls | 13, with 54.0% of the weighted cost |
| Sessions past 300K of context | 11 |
| Reads of the hand-off of more than 15,000 characters | 13 |
| Cache writes after a gap of more than one hour | 15, with 18.9% of all cache writes |
| Input tokens saved by the replay, R from 0 to 55,000 | 20.6% to 17.2% |
| Weighted cost saved by the replay, R from 0 to 55,000 | 14.8% to 11.6% |

The first audit replayed ten sessions before D-749 and found 51 to 55 percent, with R at 0. D-749 made each call smaller, so less context passes 300K, and the rule saves less.

## The findings

- **D-749 cut the context of each call.** The median fell from 397K to 259K tokens in the five sessions, and to 165K in all 47.
- **The checkpoint rule still saves a third of the input of a long session.** At the expected R, the five sessions save 34.8 percent of their input tokens and 27.8 percent of the weighted cost.
- **Over all sessions the saving is smaller.** It is 17 to 21 percent of the input tokens, because 36 of the 47 sessions never passed 300K.
- **Reads through the shell grew.** Search and shell reads made 28.5 percent of the carried context in the five sessions, against 10.2 percent before D-749.
- **The hand-off budget holds.** Full reads of the hand-off fell from 13 in ten sessions to 1 in five sessions.
- **The line of D-750 in the skill never fired.** Section 4 of the `one-pr-one-session` skill tells a session to offer a context compaction past 300K. Four of the five sessions passed 300K, and no reply of the five named a compaction. None of the 47 sessions holds a compaction. A session can not see the size of its context.
- **No wait of more than one hour occurred in the five sessions.** Over all 47 sessions, 15 such waits wrote 2.6M tokens to the cache again.

## The owner decision (D-946)

The owner adopted the checkpoint rule at 300K on 2026-09-25. The hook `.claude/hooks/context_checkpoint.py` reads the usage of the last call, and it tells the session at 300K and at each further 100K. Section 4 of the `one-pr-one-session` skill holds the steps. D-750 also parked the split of the `mtg-corpus` skill and the reasoning effort of the harness. Both stay as they are. The five sessions ran each call at high effort.

## How to measure again

1. Find the transcripts under `~/.claude/projects/`, in the folder of this repo.
2. Keep the newest sessions with 40 or more calls and one or more owner prompts.
3. Add the usage fields of each unique request id, main thread and subagents apart.
4. Divide the size of each tool result by 2.43 to get its tokens.
5. Multiply the tokens of each result and each output by the calls that follow.
6. Replay each session with a reset at 300K tokens, and add R at each reset.
7. Compare the shares and the savings with the tables of this document.

Write the result to a new dated document. Quote no prompt text, and write no session id of the transcripts.
