# The author response to the review of #293

Author provider: Claude Code.

## P2-1: The setup still directs the owner to copy the Codex login

Result: full merit.

Evidence: at `141fe63`, the error of `scripts/live-evals.sh` for a missing login said "Copy ~/.codex/auth.json". The comment above the token checks called the Codex login a copy of the login of the owner. `docs/reference/paid-targets.md` said the same. Only the setup header carried D-1184.

Correction:

- `scripts/live-evals.sh`: the error now says to run `codex login` with the `CODEX_HOME` of the sessions, and never to copy `~/.codex/auth.json` (D-1184). The comment names the own Codex login, and the gh token alone as a copy.
- `docs/reference/paid-targets.md`: a session uses its own Claude token, its own Codex login, and a copy of the gh login.

Regression check: a search of `scripts/live-evals.sh` and `docs/reference/paid-targets.md` for "Copy ~/.codex" and for the Codex login as a copy finds no line. `shellcheck scripts/*.sh` passes.

## Round 2: The ledger gate has no source evidence in the checkout

Result: full merit. The gate holds, but the checkout did not hold its evidence.

Evidence: the ledger and the session log stay under the folder of the live evals, outside the repository (D-1171). The ledger holds a start line and a measured end line for each of two runs. In the session log, the first end line is at line 106, and the second start is at line 113. The spend tool reads the logs and the ledger, and it calls no provider. It gave `measured $0.0013 in 2 runs`, and the charge was the measured sum, not the full budget (D-1173).

Correction: `docs/reference/live-evals-ledger-2026-10-06.md` holds the four ledger lines, the order of the log lines, and the recompute with its output. The hand-off cites it.
