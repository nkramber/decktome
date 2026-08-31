# mtg-deck-builder

An agentic Magic: The Gathering deck builder. You upload a ManaBox collection export (optional), describe the deck you want, answer a few questions, and get a legal, validated deck with card art. Go + Protobuf (Connect-RPC) + TypeScript (React). Design: `docs/design-roadmap.md`.

Status: `docs/SESSION-HANDOFF.md` says where the work stands. The local stack serves the card database, the collection import, the LLM role layer, and the question workflow (`AgentService.Chat`). It also serves the deck generator, the revision turn, and the web app.

## Run it locally

The whole stack runs on your machine with no cloud account and no credentials.

### 1. Check your tools

```bash
make doctor
```

Each `MISSING` line shows the fix command. Full install steps: `docs/setup.md`. You need: Go, Node 22.12 or newer on the 22 line (`.nvmrc` pins 22.23.2) + pnpm, the firebase CLI, and Java 17. An optional tool prints a `warn` line and not a failure: `brew`, the `docker` CLI, `gcloud`, `python3`, and `shellcheck`. A stopped Docker daemon is a `warn` line too.

### 2. Install the web dependencies

```bash
cd web && pnpm install && cd ..
```

Optional: copy `.env.example` to `.env` and add the LLM provider keys. Without keys, `make dev` sets `LLM_REQUIRE_KEYS=0`, and the fixture fake serves the health role only (D-267). A chat turn then fails with a clear message that names the role it can not serve.

### 3. Start the stack

```bash
make dev
```

This starts five processes under one supervisor. One Ctrl-C stops them all:

| Process | Port | What it is |
|---|---|---|
| Firestore + Auth emulators | 8281, 9199 (UI: 4100) | Local database, no cloud |
| fake-gcs-server | 4443 | Local file storage |
| API (`cmd/api`) | 8090 | Connect-RPC + `/healthz` (liveness) + `/readyz` (a snapshot is loaded) |
| Worker (`cmd/worker`) | - | Card-data refresh jobs |
| Web (Vite) | 5180 | http://localhost:5180 |

Note: the ports avoid the Wallabee dev stack (8080, 8181, 4000, 5173), which can run on the same machine.

The web app reads two optional variables. `VITE_API_BASE_URL` sets the API origin, and the default is empty, so the app calls the same origin (the Vite proxy in dev). `VITE_AUTH_EMULATOR_HOST` names the Auth emulator, and the dev default is `127.0.0.1:9199` (`web/apps/web/src/lib/api.ts` and `lib/firebase.ts`).

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

### 6. Test it in the browser

The web app is the front door of the product. Open http://localhost:5180 while `make dev` runs. The card-data line sits in the footer of every screen. This is what works today (Phase 3B, PR-16 to PR-17).

1. Sign in. The form runs over the Firebase Auth emulator, so any email and password work. Click "Create account" the first time. The account lives in the emulator and goes when it stops.
2. You land on Build. A signed-in reader starts at the chat, because that is what the app is for (D-334).
3. Add a collection, or skip it. Open "Collection" and drop a ManaBox CSV export on the zone. `go/internal/collections/testdata/manabox_collection.csv` is a real one. The head of the screen then shows the count, the unique cards, and the rarity spread.
4. Pick the pool. "Build" in the header opens a menu of your collections. Choose one, or choose "Any card" (D-332).
5. Write what you want, for example "a mono-green Commander deck around elves", and press Enter. The agent streams its questions, and the message box leaves while it works (D-325).
6. Read the deck. The moment a deck exists, the app moves to its own address, `/decks/<id>` (D-335). The deck fills the page, and the conversation docks at the bottom left. "History" opens the thread over it.
7. Use the deck. Rename it, star it, or delete it from the row above the deck. The export panel under the deck copies or downloads the Arena text, which ManaBox imports (D-15, D-307). The buy list opens on request.
8. Ask for a change. Write "no 6 or 7 mana cards" in the dock. The agent answers in words, and a new deck moves you to its address (PR-12B).
9. Open "Decks". The library reads `DeckService.ListDecks`. Search by name or commander, and filter by format, power, or favorites.

A chat turn calls the real providers and spends money, at the rates of `docs/SESSION-HANDOFF.md`. The card images come from the Scryfall CDN, so the deck view needs the internet.

Every request from the browser carries the Firebase ID token in the `Authorization` header (D-268, D-275). A request with no token falls back to the debug user `local-dev` under `make dev` only. The emulator forgets nothing while `.local/` stays, and `rm -rf .local` starts clean.

To test the sign-in path without the browser, create a user on the emulator and call the API with its token:

```bash
TOKEN=$(curl -s -X POST 'http://127.0.0.1:9199/identitytoolkit.googleapis.com/v1/accounts:signUp?key=demo-key' \
  -H 'Content-Type: application/json' \
  -d '{"email":"me@example.com","password":"password123","returnSecureToken":true}' | python3 -c 'import sys,json;print(json.load(sys.stdin)["idToken"])')
curl -s -X POST -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' \
  -d '{}' localhost:8090/mtg.v1.CollectionService/ListCollections
```

The web checks run with `pnpm --dir web lint`, `typecheck`, `test`, and `build`. `make lint` and `make test` run them too.

### Container variant

```bash
make dev-docker
```

Compose runs the emulators, fake GCS, and the API in containers. Compose reads `.env`, and the API fails fast with a key error when `.env` holds no provider key (D-267). The native `make dev` is the normal path. To load the card database in the container stack, run `docker compose --profile seed run --rm worker`. The API loads the snapshot within 15 seconds. Verified end to end on 2026-08-24.

## Development commands

```bash
make help          # every target, one line each
make doctor        # check the tools
make buf           # build buf into .bin/
make proto         # regenerate Go + TS from proto/ (commit the output)
make proto-check   # fail when the generated code is stale
make proto-breaking       # buf breaking against main
make lint          # go vet, golangci-lint, eslint, tsc, and the STE check
make lint-go / make lint-web / make ste-check   # one lane each
make test          # go test -race, vitest
make test-repeat TEST=TestName RUNS=25   # flake hunt
make test-smoke    # live LLM smoke test, reads the keys from .env
make llm-defaults-check   # warn when roles.json or prices.json changed
make cover         # Go coverage summary
make build         # Go binaries + web bundle
make run-api / make run-worker / make run-web   # one process each
make dev           # the whole local stack
make dev-docker    # the same stack in containers
make dev-seed      # one-shot card snapshot refresh into the local stack
make clean         # remove the build output
make m5-sheet      # build the M-5 scoring sheet from a gate document
make m5-report     # read the scored sheet and compute the thresholds
make themes-check  # check the theme slugs and the commander ranking
make store-check   # session store against the local Firestore emulator
make candidates-review   # write the PR-6 gate document from the local snapshot
```

`make help` shows every target, the money targets included.

`cd go && go run ./cmd/tune-check` compares an eval summary with its baseline, and it costs nothing.

CAUTION: the nine targets below and `scripts/autotune.sh` call the real LLM providers and spend money. `make autotune` is free. Ask the owner before each run, and write to a new output file (D-65). Each gate target refuses to overwrite a scored output, and each one needs `.env`.

```bash
make questions-gate    # 104 conversations, $0.15 to $0.17, about 20 minutes
make questions-eval    # score a gate run, $0.092 to $0.104 (runs 14 to 25), about 13 minutes
make eval-calibrate    # eval model against claude-sonnet-5, $0.25 to $0.30
make autotune          # free: print how to start the paid loop, scripts/autotune.sh ($0.25 an iteration)
make deck-gate         # the PR-8 gate document, $1.09 for 18 prompts (run 8)
make revise-gate       # two base decks and six revisions (PR-12B), about $0.30 (run 2, $0.29)
make chat-probe        # drive the real Chat RPC to a deck, a few cents
make generate-probe    # build one deck with the real generate role, a few cents
make summary-judge     # judge every deck summary of a gate document (F-26), a few cents
make test-smoke        # live LLM smoke test, reads .env, a few cents
```

The question gate cost is from 2026-08-26, and the deck gate cost is from run 8 (2026-08-29).

CI runs on pull requests only, and a new push to a branch cancels the run in progress. A first job reads the diff against the base branch. Each job runs only when its inputs changed, so a docs change runs the STE check and nothing else. A merge to `main` runs nothing, because the pull request verified the same tree. A weekly schedule runs govulncheck alone, at about 2 minutes a week (D-305). The owner hit 90 percent of the monthly minutes in six days on 2026-08-28, and each run cost 25 billed minutes before this rule (D-286).

Rules for contributors and agents: `AGENTS.md`. Machine setup: `docs/setup.md`. Design and roadmap: `docs/design-roadmap.md`. Decisions: `docs/decisions.md`.

## Repo layout

- `proto/` - the one API contract (buf, `mtg.v1`).
- `go/` - Go module: `cmd/api`, `cmd/worker`, `internal/*`, `gen/` (generated, committed).
- `web/` - pnpm workspace: `apps/web` (React 19 + Vite), `packages/api-client` (generated, committed).
- `docs/` - roadmap, decisions, open questions, reference notes, `setup.md`.
- `.claude/` - agent skills, including the MtG rules corpus.

## Data and images

Card data and images come from Scryfall bulk data, updated daily. Prices are near-mint market estimates. This project follows the Wizards of the Coast Fan Content Policy through Scryfall's guidelines. Card data is free to access. Images keep artist and copyright visible. Wizards of the Coast and Scryfall endorse nothing here.
