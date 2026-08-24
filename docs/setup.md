# Developer machine setup (PR-0b)

This procedure prepares a macOS machine (Apple Silicon) for this repo. Run `make doctor` after each step to see what is still absent. The pinned versions live in `go/go.mod`, `.nvmrc`, `web/package.json`, and `Makefile`. This page does not repeat them.

Verified 2026-08-23 on macOS 26.5.2, arm64, Homebrew 6.0.12. Node moved to the 22 LTS line (22.23.2) on 2026-08-24 (B-1, D-52).

## 1. Check the machine

1. Open a terminal.
2. Run `make doctor`.
3. Read the list. Each `MISSING` line shows the fix command.

## 2. Install Homebrew

Skip this step if `brew --version` prints a version.

1. Follow the one-line install at https://brew.sh.
2. Run `brew --version` to confirm.

## 3. Install Git and the Xcode command line tools

1. Run `xcode-select --install`.
2. Accept the dialog.
3. Run `git --version` to confirm.

## 4. Install Go

1. Read the `go` line in `go/go.mod`. That is the required version.
2. Download that version from https://go.dev/dl/ and run the installer. Or run `brew install go`.
3. Run `go version` to confirm. The version must match `go/go.mod`.

Note: `buf`, `protoc-gen-go`, and `protoc-gen-connect-go` are not separate installs. `make proto` builds them from the `tool` directives in `go/go.mod`.

## 5. Install Node and pnpm

1. Install nvm from https://github.com/nvm-sh/nvm.
2. In the repo root, run `nvm install`. It reads `.nvmrc` (Node 22.23.2, the 22 LTS line).
3. Run `nvm use`. Vite 7 needs Node 20.19 or 22.12 and newer. An older Node prints a warning on each build and fails the web tests.
4. Run `corepack enable`.
5. Run `corepack prepare pnpm@9.2.0 --activate`. The version comes from `packageManager` in `web/package.json`.
6. Run `pnpm --version` to confirm.

## 6. Install the Firebase CLI and Java

The Firestore and Auth emulators (PR-0c) need both.

1. Run `npm install -g firebase-tools@14.14.0`. This is the version in `docker/emulators.Dockerfile` and `scripts/doctor.sh`.
2. Run `brew install openjdk@17`.
3. Follow the `brew` note to put Java on your PATH, or run `sudo ln -sfn /opt/homebrew/opt/openjdk@17/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk-17.jdk`.
4. Run `firebase --version` and `java -version` to confirm.

## 7. Install Docker (D-10)

Docker is not needed for `make dev`. It is needed for the Compose file (PR-0c) and for Cloud Run parity. When the daemon is stopped, `make doctor` prints a `warn` line and still passes.

1. Run `brew install --cask docker`. Or download Docker Desktop from https://docs.docker.com/desktop/setup/install/mac-install/.
2. Open Docker Desktop once. Accept the license. Wait for the whale icon to show "running".
3. Run `docker --version` to confirm.
4. Run `docker run --rm hello-world` to confirm the daemon works.

CAUTION: Docker Desktop needs a paid subscription for large companies. Personal use is free. Read the license before you accept it.

## 8. Install the Google Cloud CLI

Needed for `mtg-dev` and `mtg-prod` (D-24). Not needed for local work.

1. Run `brew install --cask google-cloud-sdk`.
2. Run `gcloud init`.
3. Run `gcloud --version` to confirm.

## 9. Install the project

1. Run `cd web && pnpm install && cd ..`.
2. Optional: run `cp .env.example .env` and add the LLM provider keys. `make dev` and `make test-smoke` read this file.
3. Run `make proto`. It builds buf and regenerates code. `git status` must stay clean.
4. Run `make lint`.
5. Run `make test`.
6. Run `make dev`. Open http://localhost:5180. The page shows the API health check.
7. Press Ctrl-C to stop the stack.

Other targets: `make test-smoke` (live LLM test, needs `.env`), `make llm-defaults-check`, `make dev-seed`, `make cover`, `make build`, `make proto-breaking`. Run `make help` for the full list.

## 10. Final check

1. Run `make doctor`.
2. Every line must read `ok`. This is the PR-0b gate.

## Known problems

| Symptom | Cause | Fix |
|---|---|---|
| `make proto` fails with "no required module provides package" | `go/go.sum` is stale | Run `go -C go mod download all`, then `make proto`. |
| Port 5180 or 8090 is in use | Another dev server runs | Run `lsof -ti :5180 -ti :8090 \| xargs kill`. |
| Vite reports "Cannot find native binding" | Vite 8 with pnpm 9 | The repo pins Vite 7 (D-35). Run `pnpm install` again. |
| `firebase emulators:start` fails with a Java error | Java is not on PATH | Repeat step 6.3. |
| `make test` fails in vitest with `ERR_REQUIRE_ESM` | Node is older than 22.12 | Run `nvm install` and `nvm use`. |
| `make test-smoke` says `.env is absent` | No `.env` file | Run `cp .env.example .env` and add the keys. |
