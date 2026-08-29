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

The web app is the front door of the product. Open http://localhost:5180 while `make dev` runs. It shows the sign-in form, and the API health line sits in the footer. This is what works today, and what does not. The roadmap slices are PR-11 and PR-12 (`docs/reference/ui-plan-2026-08-28.md`).

1. Sign in. The form runs over the Firebase Auth emulator, so any email and password work. Click "Create account" the first time. The account survives a restart of `make dev`, because the emulator exports its users to `.local/firestore` at exit.
2. Upload a collection, or skip. Choose a ManaBox CSV export. `go/internal/collections/testdata/manabox_collection.csv` is a real one with 2,548 rows. The screen shows the card count, the rows that did not resolve, and the reason for each. Earlier uploads appear as a list, and one of them is the active collection. "Skip, build from any card" clears the active collection (D-37).
3. Continue to chat. Write what you want, for example "a mono-green Commander deck around elves", and press Enter. The agent streams its questions. Click an option, or type an answer in the field under each question, then click "Submit answers". A commander offer shows each card image with a pick button (D-291). The button waits until every open question has an answer. The message box hides while the agent works and while a question waits. "Stop" ends a turn that hangs. The status line says when the build runs, and the spend line shows the session cost (M-1).
4. Read the deck. It opens beside the thread after the agent builds one. The cards group by role, and each card tile shows the full image and no caption (D-290). A double-faced card shows both faces. An owned card carries a mark, and a card to buy shows its price. The findings, the legality date, the mana curve, and the color sources sit above the cards.
5. Open "Decks". The list reads `DeckService.ListDecks`. "View" opens a deck in place. The session link returns to its chat, and a reload of the chat rebuilds the thread from the stored session.
6. Read the footer. It shows the API status and the date of the card snapshot. "Card data: not loaded yet" means step 4 above did not run.

After the deck, write what you want changed, for example "no 6 or 7 mana cards". The agent answers in words. Then it asks a question when the request is unclear, revises the deck when it is clear, or says why it made no change (D-283, D-284). The deck view shows what changed. A change to a setting, for example the format or the bracket, builds the deck again from the start.

What the browser cannot do yet: export a deck (PR-13). A chat turn calls the real providers and spends money, at the rates of `docs/SESSION-HANDOFF.md`. The card images come from the Scryfall CDN, so the deck view needs the internet.

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
