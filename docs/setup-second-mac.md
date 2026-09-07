# Continue the work on a second Mac

This page tells you how to move the work of this repo to another Mac and continue it there. It names what to carry from this Mac, what to install, how to prove the machine, and where the work stands. `docs/setup.md` holds the base install of the tools, and this page repeats its commands with the versions of 2026-09-04. Read `CLAUDE.md` and `docs/SESSION-HANDOFF.md` on the new Mac before you change anything.

Verified on 2026-09-04 against this Mac: macOS 26.5.2, Apple Silicon, Go 1.27.0, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, OpenJDK 17, Claude Code 2.1.233.

## 1. What to carry from this Mac

The repo comes from GitHub. Five things do not, and the paid gates and the dev stack need them.

| Item | Path on this Mac | Size | Why |
|---|---|---|---|
| The secrets | `.env` in the repo root | 4 KB | Four keys: `OPENAI_API_KEY`, `ANTHROPIC_API_KEY`, `TOPDECK_API_KEY`, `AUTOTUNE_FIXER_CMD`. Every paid target and `make meta-refresh` read it. |
| The card snapshots and the meta store | `.local/gcs/mtg-local-cards/` and `.local/gcs/mtg-local-cards.bucketMetadata` | 574 MB | Ten Scryfall snapshot versions (310 MB) and the meta store: the deck lists, the commanders, the quality model, the precon table (263 MB). Every gate reads the newest complete snapshot, and the quality model version is part of every fingerprint. |
| The eval run files | `.local/tune/` | 4.6 MB | The JSON of every question eval run. `make questions-eval` writes there, and `tune-check` reads it. |
| The emulator data | `.local/firestore/` | 744 KB | The sessions, decks, and collections of the local dev stack. Optional. |
| The Claude Code memory | `~/.claude/projects/-Users-nate-Repos-decktome/memory/` | 60 KB | The facts Claude Code keeps about you and this repo. Not in git. |

CAUTION: the fake GCS server keeps the metadata of each object in an extended attribute, `user.metadata`. A copy that drops the attribute breaks every listing of the bucket, and the API then loads no card index. Use `tar` or `ditto`, which keep the attributes on macOS. Do not use a cloud drive that strips them, and do not write a file into `.local/gcs` by hand.

Pack the data on this Mac:

```
cd ~/Repos/decktome
tar -czf ~/Desktop/mtg-carry.tgz .env .local/gcs/mtg-local-cards .local/gcs/mtg-local-cards.bucketMetadata .local/tune .local/firestore
tar -czf ~/Desktop/mtg-claude-memory.tgz -C ~/.claude/projects ./-Users-nate-Repos-decktome/memory
```

CAUTION: `mtg-carry.tgz` holds your API keys. Move it with AirDrop or a USB drive, and delete both copies after the unpack.

## 2. Get the code

1. Install Homebrew from https://brew.sh, if `brew --version` prints nothing.
2. Run `xcode-select --install` and accept the dialog.
3. Run `brew install gh`.
4. Run `gh auth login`. Choose GitHub.com, then SSH, and let it upload a new key.
5. Run `ssh -T git@github.com`. The answer names your account.
6. Run `git clone git@github.com:nkramber/decktome.git ~/Repos/decktome`.
7. Run `cd ~/Repos/decktome`. The work continues on `main`, and every new branch starts from it (D-494).

Keep the path `~/Repos/decktome`. Claude Code names its memory directory after the repo path, and section 6 explains the rule.

## 3. Install the tools

The pinned versions live in `go/go.mod`, `.nvmrc`, `web/package.json`, and `scripts/doctor.sh`. `make doctor` reads them and prints one line per tool, with the fix command on each `MISSING` line.

1. Run `brew install go`. `go version` must print 1.27.0 or newer.
2. Install nvm from https://github.com/nvm-sh/nvm, then open a new terminal.
3. In the repo root, run `nvm install`. It reads `.nvmrc` and installs Node 22.23.2.
4. Run `nvm alias default 22.23.2`. Node 20 on the PATH fails every web test with `ERR_REQUIRE_ESM`.
5. Run `corepack enable`, then `corepack prepare pnpm@9.2.0 --activate`.
6. Run `npm install -g firebase-tools@14.14.0`. The Firestore and Auth emulators need it.
7. Run `brew install openjdk@17`. The emulators run on Java.
8. Run `sudo ln -sfn /opt/homebrew/opt/openjdk@17/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk-17.jdk`.
9. Run `brew install shellcheck`. CI lints the scripts with it.
10. Run `cd web && pnpm install && cd ..`.
11. Run `make buf`. It builds `.bin/buf` from `go/go.mod`, and `make proto` uses it.
12. Run `make doctor`. Every line must read `ok` or `warn`.

Two tools are optional. `brew install --cask docker-desktop` serves `make dev-docker` only, and `brew install --cask gcloud-cli` serves the cloud deploy of PR-22 only. `make dev` needs neither. Python 3 ships with macOS, and the STE checker uses it.

## 4. Put the data in place

1. Move `mtg-carry.tgz` and `mtg-claude-memory.tgz` to the new Mac.
2. Run `cd ~/Repos/decktome`.
3. Run `tar -xzf ~/Desktop/mtg-carry.tgz`. It writes `.env` and `.local/`.
4. Run `chmod 600 .env`.
5. Check the attributes on one snapshot object with the two lines below.

The output must hold `user.metadata`:

```
V=$(ls .local/gcs/mtg-local-cards/scryfall | tail -1)
xattr -l ".local/gcs/mtg-local-cards/scryfall/$V/complete"
```

Without the carried data, the machine can rebuild it, at a cost in time and in fingerprints:

- `make dev`, then `make dev-seed` in a second terminal, downloads the newest Scryfall day, about 110 MB. The gates then read a newer snapshot date than the runs of 2026-09-04.
- `make meta-refresh` reads the deck list sources over the network, about 40 minutes on the first run, and fits the quality model. It needs `TOPDECK_API_KEY` in `.env`. The model version then differs from every fingerprint on record, and the compare reports an epoch move.

## 5. Prove the machine, for nothing

1. Run `make doctor`.
2. Run `make lint`. It runs vet, golangci-lint, the web lint and typecheck, and the STE check.
3. Run `make test`. It runs the Go race suite and the web tests.
4. Run `make eval-check`. A baseline with no newer run reads NOT EVALUATED, and that is right.
5. Run `make themes-check`. It reads the snapshot.
6. Run `QUALITY_GATE_OUT=/tmp/qg.md QUALITY_GATE_RUN=/tmp/qg.jsonl make quality-gate`. It reads the meta store.
7. Run `make dev`, then open http://localhost:5180. The footer shows the card snapshot date.
8. Press Ctrl-C to stop the stack.
9. Run `make test-smoke`. It spends a few cents and proves the keys.

The quality gate of step 6 reads FAIL on the two precon bars, as run 12 did on 2026-09-03. That is the state of the corpus and not the machine.

## 6. Claude Code

1. Install Claude Code and its VS Code extension from https://claude.com/claude-code.
2. Sign in with the same account as this Mac.
3. Run `claude --version`. This Mac runs 2.1.233.
4. Run `tar -xzf ~/Desktop/mtg-claude-memory.tgz -C ~/.claude/projects`.
5. Run `claude mcp add playwright -- npx @playwright/mcp@latest`.
6. Open the repo in VS Code, or run `claude` in the repo root.

The memory directory is `~/.claude/projects/<key>/memory/`, and the key is the repo path with each slash as a dash. On this Mac the key is `-Users-nate-Repos-decktome`. A different user name or a different path gives a different key, and Claude Code then loads no memory. Rename the unpacked directory to the key of the new Mac when they differ.

The skills and the agent of this repo live in `.claude/` and come with the clone. `~/.claude/settings.json` on this Mac holds no secret: the model, the effort level, and the thinking flag. Set them again on the new Mac, or copy the file. `AUTOTUNE_FIXER_CMD` in `.env` names the `claude` command, so the tuning loop needs it on the PATH.

## 7. The paid targets

Ask before every paid run, and write every run to a new document (D-65). Each target keeps a spend guard, and the Makefile sets it. The sweep of PR-15 runs the paid suites under a cap. Its estimate of a step is the cost of the last run file of that suite. The costs measured on 2026-09-04:

| Run | Command | Cost |
|---|---|---|
| Question gate with its eval | `EVAL_SWEEP=1 go -C go run ./cmd/eval sweep -cap 0.50 -suites questions,question-eval -continue` | $0.28 |
| Deck gate, 25 prompts with the plan judge | `EVAL_SWEEP=1 go -C go run ./cmd/eval sweep -cap 4.50 -suites decks` | $3.79 |
| Tier judge of a deck gate document | `EVAL_SWEEP=1 go -C go run ./cmd/eval sweep -cap 0.50 -suites tier-judge` | $0.36 |
| Revise gate | `EVAL_SWEEP=1 go -C go run ./cmd/eval sweep -cap 1.50 -suites revise` | $1.07 |
| The whole list | `EVAL_SWEEP=1 go -C go run ./cmd/eval sweep -cap 6.00 -continue` | about $5.50 |

`-dry` prints the plan and spends nothing. The tier judge reads FAIL on the open bar of D-488, and the sweep stops there without `-continue`. After a run that passes, `go -C go run ./cmd/eval baseline -suite <suite> -run <file.jsonl>` records it, and `make eval-check` compares the next run with it. `CLAUDE.md` holds every other paid target and its cost.

## 8. Where the work stands on 2026-09-04

- PR-15, the eval harness, merged as #65 on 2026-09-04 after its paid gate passed on every lane (D-514 to D-517). The three suites have baselines: question gate run 35, deck gate run 16, and revise gate run 9.
- `main` is the merge commit, and the tree is clean. PR-22, the deploy for invited users, is next on a new branch from `main`, then PR-23.
- `docs/SESSION-HANDOFF.md` is the resume point, and `docs/owner-questions.md` holds no open row on 2026-09-04.

## Known problems

| Symptom | Cause | Fix |
|---|---|---|
| Every web test fails with `ERR_REQUIRE_ESM` | Node 20 is on the PATH | Run `nvm use`, or set `nvm alias default 22.23.2`. |
| `make dev` logs "bucket doesn't exist" every 15 seconds | An object under `.local/gcs` lost its `user.metadata` attribute | Copy the data again with `tar` or `ditto`. Never write into `.local/gcs` by hand. |
| `go run ./cmd/eval` says "no required module provides package" | The shell is not in the repo root, or the branch predates #65 | Run `cd ~/Repos/decktome && git checkout main && git pull`. |
| `make lint` says "No rule to make target" | The shell is in `go/` | Run `make` from the repo root. |
| A paid target says `.env is absent` | No `.env` in the repo root | Unpack `mtg-carry.tgz` again, or copy `.env.example` and add the keys. |
| `firebase emulators:start` fails with a Java error | Java is not on the PATH | Repeat step 8 of section 3. |
| A gate refuses to write | The document or its run file holds a result | Name a new file (D-65). |
