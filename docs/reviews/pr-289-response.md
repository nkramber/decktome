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
- The install itself did not run in a test, because a pass loads the agent. The answer to P1-2 tests the loop of its preflight.

## The Gitar finding on the Push line of the record

Result: full merit, and the review tool fixed it. The record commit `6de6c36` names `b58f50c` in the Push line, and `b58f50c` is on the branch.

## P1-2: The credential check ignores read access control lists

Result: full merit.

Evidence: `chmod +a 'everyone allow read'` on a file of mode 600 left `stat -f %Lp` at 600. `ls -lde` printed the entry `0: group:everyone allow read` on a second line.

Correction:

- `scripts/live-evals.sh`: `private_file` also refuses a file with any entry of an access control list. `ls -lde` prints one line for each entry after the line of the file (D-1164).
- `scripts/live-evals-launchd.sh`: the preflight of the install refuses the same files with the same rule.
- A `shellcheck disable=SC2012` comment states the reason: `find` prints no access control list, and each path is fixed.

Regression checks:

- `private_file` refused a file of mode 600 with `everyone allow read`: "has an access control list".
- It refused a file of mode 644, and it passed a clean file of mode 600 and the three real login files.
- The loop of the install preflight, with its `SECRETS` and `HOME_DIR` set to a test folder, refused a `claude-token` with the same entry. It passed the same folder with no entry.
- `shellcheck scripts/*.sh` passed.

## P2-1: The version gate does not pin jq

Result: full merit.

Evidence: `PINS` held no line for jq. Under the PATH of the agent, jq is `/usr/bin/jq`, so the macOS build pinned it. A Homebrew jq that comes first on the PATH at the time of the install has no pin.

Correction: `PINS` in `scripts/live-evals.sh` holds `jq --version` with the line `jq-1.7.1-apple` (D-1165).

Regression check: under `/bin/bash` and the PATH of the agent, the pin loop read 10 pins and no miss. The test of a wrong go pin proved the stop and the single notice of a miss for each program of the loop.

## The Blocked verdict of round 3

Result: the review found no open defect. It asked for independent evidence of two gate items, and the owner answered both (D-1167).

- The notice of a missed pin: the owner confirmed that one notice reached the phone, and that the second run sent none.
- The launchd probe of the branch: the owner chose the first real tick after the merge as the one proof. The gate of PR-131 in `docs/design-roadmap.md` lost its pre-merge launchd item, and its last item names the first tick.
