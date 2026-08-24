# mtg-deck-builder

An agentic Magic: The Gathering deck builder. You upload a ManaBox collection export (optional), describe the deck you want, answer a few questions, and get a legal, validated deck with card art. Go + Protobuf (Connect-RPC) + TypeScript (React). Design: `docs/design-roadmap.md`.

Status (2026-08-24): Phase 1 in progress. The local stack serves the card database and collection import. The deck-building agent is not built yet.

## Run it locally

Everything runs on your machine with no cloud account and no credentials.

### 1. Check your tools

```bash
make doctor
```

Each `MISSING` line shows the fix command. Full install steps: `docs/setup.md`. You need: Go, Node + pnpm, the firebase CLI, Java 17, and Docker (only for `make dev-docker`).

### 2. Install the web dependencies

```bash
cd web && pnpm install && cd ..
```

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

Compose runs the emulators, fake GCS, and the API in containers. The native `make dev` is the normal path.

## Development commands

```bash
make help          # every target, one line each
make proto         # regenerate Go + TS from proto/ (commit the output)
make lint          # go vet, golangci-lint, eslint, tsc
make test          # go test -race, vitest
make test-repeat TEST=TestName RUNS=25   # flake hunt
```

Rules for contributors and agents: `AGENTS.md`. Machine setup: `docs/setup.md`. Design and roadmap: `docs/design-roadmap.md`. Decisions: `docs/decisions.md`.

## Repo layout

- `proto/` - the one API contract (buf, `mtg.v1`).
- `go/` - Go module: `cmd/api`, `cmd/worker`, `internal/*`, `gen/` (generated, committed).
- `web/` - pnpm workspace: `apps/web` (React 19 + Vite), `packages/api-client` (generated, committed).
- `docs/` - roadmap, decisions, open questions, reference notes, `setup.md`.
- `.claude/` - agent skills, including the MtG rules corpus.

## Data and images

Card data and images come from Scryfall bulk data, updated daily. Prices are near-mint market estimates. This project follows the Wizards of the Coast Fan Content Policy through Scryfall's guidelines. Card data is free to access. Images keep artist and copyright visible. Wizards of the Coast and Scryfall endorse nothing here.
