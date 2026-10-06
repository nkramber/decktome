# The ledger evidence of the live-evals session v-lKqWOkMKOmppB0RJBmRb (2026-10-06)

This record holds the evidence of the first gate of PR-135. The ledger and the session log stay outside the repository, under the folder of the live evals (D-1171). So a review of the checkout can not read them. The author copied the lines here on 2026-10-06.

## The ledger

The file is `state/runs/v-lKqWOkMKOmppB0RJBmRb/ledger/spend.jsonl` under the folder of the live evals. It holds four lines. `ls -lO` shows the flag `uappnd`, so the file takes appends alone (D-1171).

```
{"id":"fe89cf50f975dbb3","event":"start","target":"chat-probe","usd":0,"measured":false}
{"id":"fe89cf50f975dbb3","event":"end","target":"chat-probe","calls":4,"tokens":{"input_tokens":6875,"cached_input_tokens":2725,"cache_write_tokens":2878,"output_tokens":443,"reasoning_tokens":0},"usd":0.0007356999999999999,"measured":true}
{"id":"16e39f562667ef2f","event":"start","target":"chat-probe","usd":0,"measured":false}
{"id":"16e39f562667ef2f","event":"end","target":"chat-probe","calls":6,"tokens":{"input_tokens":8097,"cached_input_tokens":5603,"cache_write_tokens":0,"output_tokens":609,"reasoning_tokens":79},"usd":0.0006099300000000001,"measured":true}
```

The ledger holds two runs. Each run has a start line and a measured end line.

## The order in the session log

The spend tool reads the session logs of the run folder, `session-<n>.log` (D-1173). This session has one log, `state/runs/v-lKqWOkMKOmppB0RJBmRb/session-1.log`. It holds 168 lines.

Each line of this log is one JSON event of the session. A tool-result event holds the full output of one command. The logger writes one spend event per output line (`go/internal/livespend/livespend.go`). So one log line can hold several spend events, each on its own output line.

| Log line | Time (UTC) | Output line | Run | Event |
|---|---|---|---|---|
| 67 | 2026-10-06T03:28:46Z | 12 of 17 | `fe89cf50f975dbb3` | end |
| 72 | 2026-10-06T03:29:09Z | 1 of 20 | `16e39f562667ef2f` | start |
| 72 | 2026-10-06T03:29:09Z | 16 of 20 | `16e39f562667ef2f` | end |

The spend lines of the output, word for word:

```
67, line 12: LIVE-EVAL-SPEND {"id":"fe89cf50f975dbb3","event":"end","target":"chat-probe","calls":4,...,"usd":0.0007356999999999999,"measured":true}
72, line 1:  LIVE-EVAL-SPEND {"id":"16e39f562667ef2f","event":"start","target":"chat-probe","usd":0,"measured":false}
72, line 16: LIVE-EVAL-SPEND {"id":"16e39f562667ef2f","event":"end","target":"chat-probe","calls":6,...,"usd":0.0006099300000000001,"measured":true}
```

The end of the first run was in the log 23 seconds before the second start. The start of a paid run reads the session logs. A run with no end line in a log counts as the full budget (D-1173).

The log holds no start line of the first run. The session sent the output of that command through `tail -15`, and the start line was the first output line. The ledger holds that start line.

An earlier copy of this record cited the transcript of the session, not `session-1.log` (D-1186).

## The recompute

The spend tool of the script reads the logs and the ledger, and it calls no provider. The author ran it on 2026-10-06 with the default budget of the script, $3.00:

```
live-evals spend -logs state/runs/v-lKqWOkMKOmppB0RJBmRb -ledger state/runs/v-lKqWOkMKOmppB0RJBmRb/ledger/spend.jsonl -budget 3.00
{"budget":3,"usd":0.00134563,"runs":2,"charged":0.00134563,"text":"measured $0.0013 in 2 runs"}
```

The charge is the measured sum, not the full budget. So the tool found the end line of each run in the logs. The text is the text of the notice in the hand-off.
