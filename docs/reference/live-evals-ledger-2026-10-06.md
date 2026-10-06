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

The session log is `claude-config/projects/<project>/6aaa12ad-9f3d-47e6-8aa2-dd15475458f1.jsonl` under the folder of the session. It holds 213 lines. The `LIVE-EVAL-SPEND` lines come in this order:

| Log line | Time (UTC) | Run | Event |
|---|---|---|---|
| 106 | 2026-10-06T03:28:46Z | `fe89cf50f975dbb3` | end |
| 113 | 2026-10-06T03:29:09Z | `16e39f562667ef2f` | start |
| 113 | 2026-10-06T03:29:09Z | `16e39f562667ef2f` | end |

The end line of the first run was in the log 23 seconds before the second start. The start of a paid run reads the session logs. A run with no end line in a log counts as the full budget (D-1173).

## The recompute

The spend tool of the script reads the logs and the ledger, and it calls no provider. The author ran it on 2026-10-06 with the default budget of the script, $3.00:

```
live-evals spend -logs state/runs/v-lKqWOkMKOmppB0RJBmRb -ledger state/runs/v-lKqWOkMKOmppB0RJBmRb/ledger/spend.jsonl -budget 3.00
{"budget":3,"usd":0.00134563,"runs":2,"charged":0.00134563,"text":"measured $0.0013 in 2 runs"}
```

The charge is the measured sum, not the full budget. So the tool found the end line of each run in the logs. The text is the text of the notice in the hand-off.
