# Context budget of a session (2026-09-16)

This document holds the audit of the token use of this repo, the method, and the sizes of the start read before and after D-749. F-155 records the finding. D-750 holds what waits for a measurement.

## Scope and method

The audit read the local Claude Code transcripts of this repo. It read the usage fields of each API call, the tool names, the file paths, the command kinds, and the size of each tool result. It read no text of a prompt, and this document quotes none.

The sample is the ten newest sessions with 40 or more API calls and one or more prompts of the owner. They ran from 2026-09-11 23:30 UTC to 2026-09-16 19:55 UTC. The audit session itself, four empty sessions, and two sessions of 8 and 21 calls stay out.

Each number has a kind:

- **Reported:** the usage fields of the transcript.
- **Exact:** a count of bytes, words, or lines of a file.
- **Estimate:** a calculation from the numbers above. The method of each estimate follows.

No tokenizer is installed. The audit fitted the ratio from 161 steps with one tool result of 8,000 or more characters. Each step compares the context of two calls, less the output of the first call. The median is 2.43 characters for each token, and the middle half runs from 2.3 to 2.6. Each token estimate here divides bytes by 2.43.

## The baseline (reported)

| Measure | Value |
|---|---|
| API calls | 1,791 in the main thread, 67 in subagents |
| Input tokens, not cached | 36,906 |
| Cache read tokens | 772,099,913 |
| Cache write tokens, one-hour lifetime | 13,397,323 |
| Output tokens | 4,249,090 |
| Thinking tokens, part of the output | 2,647,911 |
| Input tokens for each session | median 85.4M, maximum 141.3M, minimum 17.3M |
| Context of one call | median 396,862, maximum 951,409 |
| Context of the first call | 43,953 to 50,187 |
| Compactions | 1 |

Cache reads made 98.3 percent of the input tokens. So the cost of a session follows the size of its context times the number of its calls.

## Where the context came from (estimate)

The audit gave each tool result and each output a carried cost: its tokens times the calls that follow it. Over nine sessions, the one with a compaction left out, the model explains 590.5M of the 644.2M reported input tokens.

| Source | Share of the carried context |
|---|---|
| Model output, thinking included | 53.0% |
| The first-call context of the harness and `CLAUDE.md` | 12.0% |
| Search commands | 6.3% |
| Read of code files | 7.2% |
| Shell reads such as `sed` and `cat` | 3.9% |
| The hand-off, all reads | 5.5% |
| The roadmap, all reads | 2.2% |
| `git diff` and `git log` | 1.6% |
| Tests, `make verify`, and GitHub commands | about 2% |

Nine of the ten largest tool results were full reads of the hand-off, from 38 to 53 KB each. The sample holds 13 full reads of it in 8 sessions.

## The findings

- **Session length drove the most cost.** Five sessions of more than 200 calls made about 72 percent of the estimated weighted cost. The weights are 0.1 for a cache read, 2 for a one-hour cache write, and 5 for an output token.
- **A wait of more than one hour wrote the whole context to the cache again.** 14 such writes came after gaps of 64 to 1,362 minutes. They held 6.1M tokens, 46 percent of all cache writes.
- **The hand-off grew fast.** It held 35,053 bytes on 2026-09-11 and 58,788 bytes on 2026-09-16. Its resume section held 31,711 bytes.
- **The start read conflicted with the skill.** The hand-off asked for `docs/decisions.md` from D-297, about 301 KB. Step 8 of the `one-pr-one-session` skill asks for the decisions that the change touches.
- **The two lists of paid targets differed.** `CLAUDE.md` named eleven targets and one script. The hand-off named twelve targets and two scripts, and the Makefile holds all twelve.
- **The audit refuted four other causes.** GitHub polls, test output, subagent work, and a change of `CLAUDE.md` between sessions each made 1 percent or less.

## The start read before and after D-749

Base `c2c8966` is the before column. The token column is an estimate.

| File | Before (bytes, words, lines) | After (bytes, words, lines) | Change in tokens |
|---|---|---|---|
| `CLAUDE.md`, loaded into every call | 15,971, 2,501, 123 | 9,366, 1,457, 82 | about -2,700 on each call |
| `docs/SESSION-HANDOFF.md` | 58,788, 10,002, 387 | 19,895, 3,206, 133 | about -16,000 on each read |
| Its resume section | 31,711 | 2,304 | about -12,100 |
| `docs/owner-questions.md` | 5,040, 921, 37 | 5,040, 921, 37 | none |
| `docs/open-questions.md`, out of the start read now | 3,085, 512, 41 | not read at start | about -1,300 |
| `docs/decisions.md` from D-297, out of the start read now | 301,172 | the decisions that the change touches | about -124,000 if a session read it all |

The start read that sessions made in practice held `CLAUDE.md`, the hand-off, and both question files: 82,884 bytes, about 34,100 tokens. After D-749 it holds 34,301 bytes, about 14,100 tokens. That is 59 percent less.

`docs/reference/paid-targets.md` holds 10,613 bytes. A session reads it only before it runs or changes a target.

## What the savings are likely to be (estimate)

| Change | Low | Expected | High |
|---|---|---|---|
| The hand-off budget and the start read | 2.2% | 3.3% | 3.9% |
| Hard rule 13, command output | 1.5% | 3.1% | 4.6% |
| The paid-target detail out of `CLAUDE.md` | 0.6% | 0.6% | 0.6% |

Each percentage is a share of the input tokens of the sample. The hand-off row assumes 40 to 70 percent fewer tokens for each read. The output row assumes 15 to 45 percent less search and shell-read output.

## What waits (D-750)

The checkpoint rule sends a session past about 300K tokens of context to a new clean session on the same pull request. A replay of the recorded context of each call cut the input tokens by 51 to 55 percent. That replay adds no reads that a new session makes again, so it is the high bound. The ten sessions ran before #184 and before D-749, so five new sessions measure first.

## How to measure again

1. Find the transcripts under `~/.claude/projects/`, in the folder of this repo.
2. Keep the ten newest sessions with 40 or more calls and one or more owner prompts.
3. Add the usage fields of each unique request id.
4. Divide the size of each tool result by 2.43 to get its tokens.
5. Multiply the tokens of each result and each output by the calls that follow.
6. Compare the shares with the table of this document.

Write the result to a new dated document. Quote no prompt text, and write no session id of the transcripts.
