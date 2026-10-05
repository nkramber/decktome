# The author response to the review of #289

Author provider: Claude Code.

## P1-1: The setup does not protect the copied credentials

Result: full merit.

Evidence: under umask 022, `echo x > open-file` made a file of mode 644. The header said "Mode 600 for each file", but no step set the mode, and no check read it. The three real login files have mode 600, because the session wrote them under `umask 077`.

Correction:

- `scripts/live-evals.sh`: the setup in the header starts with `umask 077`. The new function `private_file` refuses a file with a group bit or an other bit. The script calls it for `claude-token`, `gh-token`, and the `auth.json` of the session Codex home, before it reads them (D-1164).
- `scripts/live-evals-launchd.sh`: the preflight of the install refuses the same three files with the same rule.
- `docs/decisions.md`, `docs/design-roadmap.md`, and `docs/SESSION-HANDOFF.md` record the check.

Regression checks:

- Under umask 022, `private_file` refused a file of mode 644: "has mode 644, and another account can read it".
- The same function passed a file that `umask 077` made, of mode 600.
- It passed the three real login files: `claude-token`, `gh-token`, and the session `auth.json`.
- The preflight of the install did not run in a test, because a pass loads the agent. It uses the same mode rule as `private_file`.

## The Gitar finding on the Push line of the record

Result: full merit, and the review tool fixed it. The record commit `6de6c36` names `b58f50c` in the Push line, and `b58f50c` is on the branch.
