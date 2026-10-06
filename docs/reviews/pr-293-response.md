# The author response to the review of #293

Author provider: Claude Code.

## P2-1: The setup still directs the owner to copy the Codex login

Result: full merit.

Evidence: at `141fe63`, the error of `scripts/live-evals.sh` for a missing login said "Copy ~/.codex/auth.json". The comment above the token checks called the Codex login a copy of the login of the owner. `docs/reference/paid-targets.md` said the same. Only the setup header carried D-1184.

Correction:

- `scripts/live-evals.sh`: the error now says to run `codex login` with the `CODEX_HOME` of the sessions, and never to copy `~/.codex/auth.json` (D-1184). The comment names the own Codex login, and the gh token alone as a copy.
- `docs/reference/paid-targets.md`: a session uses its own Claude token, its own Codex login, and a copy of the gh login.

Regression check: a search of `scripts/live-evals.sh` and `docs/reference/paid-targets.md` for "Copy ~/.codex" and for the Codex login as a copy finds no line. `shellcheck scripts/*.sh` passes.
