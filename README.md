# mtg-deck-builder

An agentic Magic: The Gathering deck builder. You upload a ManaBox collection export (optional), describe the deck you want, answer a few questions, and get a legal, validated deck with card art. Go + Protobuf (Connect-RPC) + TypeScript (React). Design: `docs/design-roadmap.md`.

Status (2026-08-26): Phase 1 complete. Phase 2 (agent) in progress: PR-7, PR-7B, and PR-10 merged (#1 to #13). The local stack serves the card database, the collection import, the LLM role layer, and the question workflow (`AgentService.Chat`). The tuning loop for the questions exists and ran once. The deck generator (PR-8) is not built yet.

## Run it locally

Everything runs on your machine with no cloud account and no credentials.

### 1. Check your tools

```bash
make doctor
```

Each `MISSING` line shows the fix command. Full install steps: `docs/setup.md`. You need: Go, Node 22 LTS (22.23.2) + pnpm, the firebase CLI, Java 17, and Docker (only for `make dev-docker`). A stopped Docker daemon is a `warn` line, not a failure.

### 2. Install the web dependencies

```bash
cd web && pnpm install && cd ..
```

Optional: copy `.env.example` to `.env` and add the LLM provider keys. Without keys, `make dev` uses the fixture fake for every LLM role.

### 3. Start the stack

```bash
make dev
```

This starts five processes under one supervisor. One Ctrl-C stops them all:

| Process | Port | What it is |
|---|---|---|
| Firestore + Auth emulators | 8281, 9199 (UI: 4100) | Local database, no cloud |
| fake-gcs-server | 4443 | Local file storage |
| API (`cmd/api`) | 8090 | Connect-RPC + `/healthz` |
| Worker (`cmd/worker`) | - | Card-data refresh jobs |
| Web (Vite) | 5180 | http://localhost:5180 |

Note: the ports avoid the Wallabee dev stack (8080, 8181, 4000, 5173), which can run on the same machine.

### 4. Load the card database (first run only)

The worker downloads the daily Scryfall bulk data (about 110 MB) into the local fake GCS on its first start. Watch `make dev` output, or force it in a second terminal:

```bash
make dev-seed
```

The API picks the snapshot up within 15 seconds. Check it:

```bash
curl -s localhost:8090/healthz
# {"status":"ok","version":"dev","card_snapshot":"2026-08-24T09:01:52Z",...}
```

Local state persists in `.local/` (gitignored). Delete `.local/` for a clean start.

### 5. Try it

Look up a card:

```bash
curl -s -X POST -H "Content-Type: application/json" \
  -d '{"name":"Sol Ring"}' localhost:8090/mtg.v1.CardService/Lookup
```

Search white lifegain cards that are legal in Commander:

```bash
curl -s -X POST -H "Content-Type: application/json" \
  -d '{"oracleTags":["lifegain"],"colorsWithin":["COLOR_W"],"legalIn":"commander","pageSize":5}' \
  localhost:8090/mtg.v1.CardService/Search
```

Import a ManaBox collection export (CSV):

```bash
python3 - <<'PY'
import base64, json, urllib.request
content = open("path/to/your_manabox_export.csv", "rb").read()
body = json.dumps({"name": "My binder", "source": "IMPORT_SOURCE_MANABOX_CSV",
                   "content": base64.b64encode(content).decode()}).encode()
req = urllib.request.Request("http://localhost:8090/mtg.v1.CollectionService/ImportCollection",
                             data=body, headers={"Content-Type": "application/json"})
print(json.load(urllib.request.urlopen(req))["collection"]["cardCount"], "cards imported")
PY
```

Open http://localhost:5180 for the web page (a health view for now).

### Container variant

```bash
make dev-docker
```

Compose runs the emulators, fake GCS, and the API in containers. The native `make dev` is the normal path. To load the card database in the container stack, run `docker compose --profile seed run --rm worker`. The API loads the snapshot within 15 seconds. Verified end to end on 2026-08-24.

## Development commands

```bash
make help          # every target, one line each
make proto         # regenerate Go + TS from proto/ (commit the output)
make lint          # go vet, golangci-lint, eslint, tsc
make test          # go test -race, vitest
make test-repeat TEST=TestName RUNS=25   # flake hunt
make test-smoke    # live LLM smoke test, reads the keys from .env
make llm-defaults-check   # warn when roles.json or prices.json changed
make proto-breaking       # buf breaking against main
make cover         # Go coverage summary
make build         # Go binaries + web bundle
make dev-seed      # one-shot card snapshot refresh into the local stack
make m5-sheet      # build the M-5 scoring sheet from a gate document
make m5-report     # read the scored sheet and compute the thresholds
make themes-check  # check the theme slugs and the commander ranking
make store-check   # session store against the local Firestore emulator
make candidates-review   # write the PR-6 gate document from the local snapshot
```

`cd go && go run ./cmd/tune-check` compares an eval summary with its baseline, and it costs nothing.

CAUTION: the four targets below call the real LLM providers and spend money. Ask the owner before each run, and write to a new output file (D-65).

```bash
make questions-gate    # 104 conversations, $0.15 to $0.17, about 20 minutes
make questions-eval    # score a gate run, $0.09 to $0.10, about 13 minutes
make eval-calibrate    # eval model against claude-sonnet-5, $0.25 to $0.30
make autotune          # print how to start the tuning loop (scripts/autotune.sh)
```

CI runs every job on every push and pull request. There are no path filters.

Rules for contributors and agents: `AGENTS.md`. Machine setup: `docs/setup.md`. Design and roadmap: `docs/design-roadmap.md`. Decisions: `docs/decisions.md`.

## Repo layout

- `proto/` - the one API contract (buf, `mtg.v1`).
- `go/` - Go module: `cmd/api`, `cmd/worker`, `internal/*`, `gen/` (generated, committed).
- `web/` - pnpm workspace: `apps/web` (React 19 + Vite), `packages/api-client` (generated, committed).
- `docs/` - roadmap, decisions, open questions, reference notes, `setup.md`.
- `.claude/` - agent skills, including the MtG rules corpus.

## Data and images

Card data and images come from Scryfall bulk data, updated daily. Prices are near-mint market estimates. This project follows the Wizards of the Coast Fan Content Policy through Scryfall's guidelines. Card data is free to access. Images keep artist and copyright visible. Wizards of the Coast and Scryfall endorse nothing here.
