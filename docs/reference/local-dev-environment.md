# Reference: local test environment, including a local Firestore

Status: research note, frozen as history (D-42). Verified 2026-08-23. D-36 supersedes the port numbers in the recipe below (Firestore 8281, UI 4100, API 8090). Answers the owner's question: "Investigate how we can handle a local Firestore database. Is that possible?"

## Short answer

Yes. The Firebase Local Emulator Suite includes a Firestore emulator. connector-syncer uses it today (`make e2e-emulator`). The Go Firestore client honors the same environment variable as the Python client. A full local stack with the UI is possible with no cloud project and no cloud credentials.

## What exists on this machine (2026-08-23)

| Tool | Present | Version |
|---|---|---|
| `firebase` CLI | yes | 14.14.0 |
| Java (the emulator needs it) | yes | OpenJDK 17.0.17 |
| `gcloud` | yes | - |
| Docker | **no** | - |
| Go | yes (go-crash-course uses 1.26) | - |
| pnpm, Node | yes (wallabee-ui pins pnpm 9.2, Node 24.14) | - |

## How connector-syncer does it

- `firebase.json` declares the Firestore emulator on `localhost:8181` and the emulator UI on port 4000.
- `make e2e-emulator` runs `firebase emulators:start --only firestore --project=wallabee-koala`.
- Every local service sets `FIRESTORE_EMULATOR_HOST=localhost:8181`. The client library then sends all traffic to the emulator with no credentials.
- `make e2e-seed` writes fixtures into the emulator.
- The `koala-sandbox` skill (in `~/.claude/skills/`) adds three safety facts. A dead `PROJECT_ID` makes Cloud Tasks and BigQuery fail on purpose. A dead `OPENAI_BASE_URL` keeps content local. The UI runs with `VITE_FLASH_BASE_URL=` (empty) so the Vite proxy forwards to the local API.

## Recipe for the deck builder

### Firestore

1. Add `firebase.json` with a `firestore` emulator entry (port 8181) and the emulator UI (port 4000).
2. Run `firebase emulators:start --only firestore --project=mtg-local`.
3. Set `FIRESTORE_EMULATOR_HOST=localhost:8181` for every Go service. The Go client (`cloud.google.com/go/firestore`) reads this variable and skips authentication.
4. Commit `firestore.indexes.json` and `firestore.rules`. The emulator loads both.
5. Add `--export-on-exit=./.local/firestore` and `--import=./.local/firestore` so data survives a restart.

### Cloud Storage (card data snapshots, deck exports)

Use `fake-gcs-server` (github.com/fsouza/fake-gcs-server). It is a Go program. It runs without Docker: `go run github.com/fsouza/fake-gcs-server@latest -scheme http -port 4443 -data ./.local/gcs`. The Go storage client honors `STORAGE_EMULATOR_HOST=localhost:4443`.

### Cloud Tasks (background jobs: card-data refresh, meta refresh)

No official emulator exists. Two options:
- Define a `Dispatcher` interface in Go. The production implementation enqueues Cloud Tasks. The local implementation calls the handler in a goroutine. This is the recommended option. It removes a moving part from the local stack.
- Run the community emulator (`github.com/aertje/cloud-tasks-emulator`). It needs a gRPC endpoint override in the client. More parity, more setup.

### Authentication

Firebase Auth has an emulator (`firebase emulators:start --only auth`). The UI connects with `connectAuthEmulator`. Go verifies ID tokens with the Firebase Admin SDK, which also honors `FIREBASE_AUTH_EMULATOR_HOST`. If the owner picks a different auth provider, the local mode needs a debug bypass like Wallabee's `E2E__LOCAL_AUTH=true`.

### LLM

The role-to-model layer gets a `fake` provider. It returns canned structured responses from fixture files. Local mode defaults to it. A developer sets a real API key to test a real model.

### Card data

The card database comes from Scryfall bulk files (`oracle_cards`, 24.5 MB compressed, daily). Local mode reads one committed snapshot or one downloaded file from `.local/scryfall/`. Local mode needs no network after the first download. The browser loads card images from `*.scryfall.io` and needs network for them only.

### UI

The Vite dev server proxies `/api` to the Go API on `localhost:8080`. Connect-RPC works over HTTP/1.1, so the dev server needs no proxy for gRPC.

## One command

`make dev` starts, in order: the Firestore emulator (and Auth emulator if chosen), fake-gcs-server, the Go API, the Go worker, and the Vite dev server. A `make dev-seed` target loads a sample ManaBox export and the Scryfall snapshot. Use a process manager (`overmind` or `foreman` style, or a Go `cmd/dev` supervisor) so one Ctrl-C stops everything.

## Open points

- Docker is absent today. The owner decided 2026-08-23: install Docker, and make the install a roadmap item. Plan: native `make dev` first, then a Compose file for Cloud Run parity.
- Auth decided 2026-08-23: Firebase Auth, with the Auth emulator in local mode.
- Firestore emulator does not enforce quotas or index requirements. Commit `firestore.indexes.json` and test index-needing queries against a real project before release.
- The emulator is single-process and slow for large seeds. Keep the seed small.
