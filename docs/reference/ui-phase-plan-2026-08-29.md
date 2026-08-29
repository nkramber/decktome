# The product UI, Phase 3B: the plan behind the roadmap

Status: plan, 2026-08-29. The roadmap holds the entries PR-16 to PR-23 (`docs/design-roadmap.md`, Phase 3B). This document holds the detail: the screens, the components, the contract changes, the deploy shape, and the gates. The owner asked for the roadmap before any code (D-319). Decisions D-310 to D-319 scope the phase.

## 1. What exists, and why it is not a product

PR-11 to PR-13 built a live-test UI (D-273). It has four screens on plain Tailwind: sign-in, the collection upload, the chat beside the deck, and the deck list. It has no design system, one theme, no navigation on a phone, no way to rename or delete anything, no card detail, and no share. The bundle loads `firebase/auth` on every route. It serves one purpose: the owner can test the agent in a browser. F-28 records this.

## 2. The decisions that shape the phase

| Decision | What it sets |
|---|---|
| D-310 | Local first, then GCP for invited users. No public sign-up. |
| D-311 | shadcn/Radix, light and dark, card-art forward, responsive. |
| D-312 | Four flows: deck library, collections, chat and build, deck view and card detail. |
| D-313 | Vitest and axe per pull request. One Playwright smoke flow on a manual trigger. |
| D-314 | An email allowlist gates the API on GCP. |
| D-315 | A public read-only share page per deck, with a revocable token. |
| D-316 | Phase 3B before Phase 4. The paid re-baseline runs in parallel. |
| D-317 | The design system and the shell come first. |
| D-318 | The sample hand is in. The goldfish simulator stays parked. |

## 3. The design system (PR-16)

### 3.1 Primitives

shadcn/ui components on Radix, copied into `web/apps/web/src/components/ui/` as the shadcn convention. The first set: Button, Input, Textarea, Select, Checkbox, Switch, Dialog, AlertDialog, Sheet, DropdownMenu, Tabs, Tooltip, Toast, Skeleton, Badge, Card, Separator, ScrollArea, and Command (the search palette). Each one is a file the repo owns, so a fix never waits on a release.

### 3.2 Tokens and themes

Tailwind 4 tokens in `src/styles/tokens.css`: a neutral scale, one accent, the five mana colors as fixed tokens (white, blue, black, red, green) plus colorless, and the semantic roles (background, surface, border, text, muted, danger, success). The dark theme overrides the neutral scale and the surfaces only. The mana colors do not change between themes. The theme follows `prefers-color-scheme` by default, and a toggle in the header stores a choice in `localStorage`.

Card art carries the color. Surfaces are neutral, and a deck page takes one accent from the commander's colors, through the mana tokens.

### 3.3 The shell

- Desktop: a left sidebar with the navigation (Build, Decks, Collection), the theme toggle, and the account menu. The content fills the rest.
- Phone: a bottom tab bar with the same three entries, and the account menu behind the avatar.
- One `PageHeader` component: title, one-line description, and the actions on the right.
- One `EmptyState` component: an icon, a sentence, and one action.
- One `ErrorState` component for a failed query, with a retry.
- Toasts for every mutation result. A destructive action opens an AlertDialog first.
- The route error page and the loading page use the same components.

### 3.4 What moves

Every existing screen moves onto the primitives with no new feature. The screens are sign-in, the collection upload, the chat, the question cards, the deck view, the export panel, and the deck list. The bundle splits by route, so `firebase/auth` loads on the sign-in route and the app shell loads first.

### 3.5 Gate

- axe passes on every route in both themes.
- The 118 web tests hold, with the assertions updated to the new markup only.
- The first paint of the app shell loads under 200 kB of JavaScript, measured by the Vite build report and recorded in the roadmap.
- The owner walks the whole path in the browser on a desktop and on a phone.

## 4. The deck library (PR-17)

### 4.1 Screens

- `/decks`: a grid of deck cards. Each card shows the commander art or the first threat as its image. It shows the name, the format, the power, the count, the buy cost, and the date. Search by name and commander. Filter by format and by power. Sort by date, name, and cost. A favorite star.
- `/decks/<id>`: the deck page. The deck view of PR-12 with the export panel, plus the actions: rename, favorite, delete, share (PR-21), and "Open the chat".
- Version history: the chain of `revised_from_deck_id`, as a list of versions on the deck page. Each version opens, and any two open side by side with the diff of `deck-stats.diffDecks`.
- Compare: pick two decks from the grid and read the diff.

### 4.2 Contract

- `DeckService.UpdateDeck(deck_id, name, favorite)`: the two mutable fields. `Deck.favorite` is a new additive field.
- `DeckService.DeleteDeck(deck_id)`: a hard delete of the deck document. The session keeps the id in `deck_ids`, and the chat shows "deck deleted" for it.
- `ListDecksRequest` gains `page_size`, `page_token`, `format`, `favorite`, and `query`. The list stays flat fields only (D-245).

### 4.3 Gate

Each action round-trips through the API and shows in the grid without a reload. A deleted deck answers `NotFound`. The grid of 100 decks renders under one second.

## 5. Collection management (PR-18)

### 5.1 Screens

- `/collection`: the list of collections with the name, the card count, the upload date, and the active mark. Rename and delete. "Upload" opens a dialog with the file picker, the progress, and the import report.
- Re-upload with a diff: an upload whose hash differs from the active collection shows the diff first: cards added, removed, and changed in count. "Replace" makes it active. The old one stays until the user deletes it.
- `/collection/<id>`: the binder. A grid of the cards with art, the count, the finish, and the condition. Search by name, filter by color, type, set, and owned count, sort by name, price, and count. Virtualized, because the owner's export holds 2,657 rows.

### 5.2 Contract

- `CollectionService.UpdateCollection(collection_id, name)` and `DeleteCollection(collection_id)`.
- `CollectionService.DiffCollections(a, b)`: the added, removed, and changed entries by Oracle id.
- `GetCollectionRequest` gains `page_size` and `page_token`. A binder page carries 200 entries.

### 5.3 Gate

The diff of the same file uploaded twice is empty. The binder of the owner's export scrolls at 60 frames per second on the owner's laptop, measured in the browser's performance panel. A deleted collection that a session names makes the session fall back to any-card mode with a notice (D-37).

## 6. The chat and build experience (PR-19)

### 6.1 Screens

- `/build`: the start screen. A short form seeds the session: the format, the power, the pool rule, the budget, and a free-text line for the theme. Each form answer goes out as a structured answer to the catalog row it fills, so the agent asks nothing the form answered. "Just chat" skips the form.
- The chat: the thread of PR-12 on the new primitives. A question card shows the options as choice chips, a card option as an art tile, and a closed question with no field (D-295).
- Build progress: the status events of the turn as a stepper (understand, shortlist, build, check, repair). The stepper reads the `status` events the API streams today, and PR-19 adds a `phase` field to them, additive.
- Error recovery: a failed turn shows the reason and a retry. A build in progress shows the D-303 notice. A session that lost its stream resumes from `GetSession`.
- `/sessions`: the list of sessions with the first message, the date, the deck count, and the cost. Rename and delete. Resume opens the chat.

### 6.2 Contract

- `AgentService.ListSessions(page_size, page_token)`, `UpdateSession(session_id, name)`, `DeleteSession(session_id)`. `Session.name` is a new additive field.
- `ChatResponse.status` gains `phase`, additive.
- The form sends the answers through `ChatRequest.answers` with the catalog question ids the API publishes through a new `AgentService.GetCatalog` read. The catalog is static per build, and the read is cached.

### 6.3 Gate

A session started from a full form asks no catalog question the form answered. The check replays the 30 gate conversations with their answers as form input. The stepper shows every phase of a real build. The owner builds one deck from the form and one from the chat in the browser.

CAUTION: the form path sends the `Q:`/`A:` shape to the classify call (D-280), which the question gate measures only from the browser. PR-19 adds the form shape to the gate's conversation set, so the next paid run measures it.

## 7. The deck view and card detail (PR-20)

### 7.1 Screens

- Card detail: a click on a card tile opens a Sheet. It shows the full image, both faces, the Oracle text, the type line, and the mana cost. It also shows the rulings with dates, the legalities, the printings with prices, the reason line of the deck, and "Open on Scryfall".
- Filters and sort on the deck view: by role, color, mana value, type, and owned. Sort by mana value, name, and price.
- Stats: the curve, the color sources, the type counts, the average mana value, and the buy cost. Each one is a small chart with a text table under it.
- Sample hand (D-318): draw seven from the exact main deck, mulligan to six and five, draw one. No turn simulation.

### 7.2 Contract

- `CardService.GetRulings(oracle_id)`: the rulings with dates. The worker downloads the Scryfall rulings bulk file with the daily snapshot, and the index holds it per Oracle id.
- `Card.legalities` is in the proto already, and the detail reads it.

### 7.3 Gate

axe passes on the Sheet. The sample hand draws from the exact list, checked by a test that draws every card of a 60-card deck. The rulings carry their dates, and the snapshot date shows.

## 8. The share link and the print view (PR-21)

### 8.1 Screens

- "Share" on the deck page makes a token and shows the link. "Revoke" ends it.
- `/d/<token>`: a public read-only deck page: the name, the format, the power, the summary, the cards by role with art, and the export button. No owner name, no collection, no session, no prices of the owner's copies.
- Print: a print stylesheet on the deck page: the list by role, the counts, and the commander, in black on white, with no images.

### 8.2 Contract

- `DeckService.ShareDeck(deck_id)` returns the token. `RevokeShare(deck_id)`. The deck stores `share_token_hash`, never the token.
- `DeckService.GetSharedDeck(token)`: unauthenticated, rate-limited per IP, and it strips every user field before it answers.

### 8.3 Gate

A revoked link answers `NotFound`. A test reads the shared deck message and finds no uid, session id, collection id, owned flag, or owned printing. The rate limit refuses the 61st call in a minute.

## 9. Deploy to GCP for invited users (PR-22)

### 9.1 Shape

- The API and the worker on Cloud Run, from the Dockerfiles of PR-0c, with the min instances at zero.
- The web app on Firebase Hosting, with a rewrite of `/mtg.v1.*` to the API.
- Real Firebase Auth with email and password. The allowlist of D-314 in one env var, `ALLOWED_EMAILS`, read by the interceptor. A uid off the list gets `CodePermissionDenied` with one sentence.
- Firestore in Native mode with the rules of the repo, which deny every client read (the Admin SDK reads).
- The card snapshot in a GCS bucket, refreshed by the worker on Cloud Scheduler.
- Secrets in Secret Manager: the provider keys.
- A per-user monthly spend cap from `Usage`, refused with one sentence when reached. A budget alert on the project.

### 9.2 Gate

An allowlisted user signs in on the deployed URL, uploads a collection, builds a deck, revises it, and exports it. The first RPC refuses a user off the list. The roadmap records the measured monthly cost at idle. Guardrail 9 holds: `make dev` still runs with no cloud dependency.

## 10. The Playwright smoke flow (PR-23)

One flow, `workflow_dispatch` only (D-313). It signs in over the emulator and uploads the fixture export. Then it starts a session from the form with the fake provider, opens the deck, and exports it. The fake provider serves canned answers for the classify, ask, and generate roles, so the flow costs nothing. About 5 minutes of Actions time per run. The owner triggers it before a merge that touches the user path.

## 11. Order and size

| Slice | Size | Gate in one line |
|---|---|---|
| PR-16 design system and shell | Three days | Every route, both themes, axe clean, shell under 200 kB. |
| PR-17 deck library | Two days | Every action round-trips, 100 decks under one second. |
| PR-18 collections | Two days | The same file twice diffs empty, the binder scrolls smooth. |
| PR-19 chat and build | Three days | The form asks nothing it answered, the stepper shows every phase. |
| PR-20 deck view and card detail | Two days | axe on the Sheet, the sample hand draws the exact list. |
| PR-21 share and print | One day | A revoked link is NotFound, the public message carries no user data. |
| PR-22 deploy | Two days | An invited user builds a deck on GCP, an outsider is refused. |
| PR-23 Playwright smoke | One day | The flow passes on the emulators for nothing. |

The paid re-baseline of the question gate and the deck gate (D-302) runs in parallel, on the owner's word (D-316).
