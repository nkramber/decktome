# Deck links: Archidekt and Moxfield, read 2026-10-03

This note holds each fact that PR-123 rests on (D-1100 to D-1108, D-1113 to D-1115). A session read each fact on 2026-10-03 from the owner's Mac, unless the row says otherwise. Each probe sent a named agent and no browser header.

## Archidekt

| Fact | Source | Read |
|---|---|---|
| `GET /api/decks/<id>/` answers 200 and the whole deck as JSON, with no login. | A probe of decks 7031486, 3045150, and 4406647 | 2026-10-03 |
| A deck id with no deck answers 404 and `{"error":"Deck not found."}`. | A probe of deck 999999999 | 2026-10-03 |
| `robots.txt` disallows `/login?`, `/partialCompare`, `/playtester-v2/`, and `/sandbox?`. It does not disallow `/api`. | `https://archidekt.com/robots.txt` | 2026-10-03 |
| The CORS header allows the origin `http://localhost:3000` alone. A page of `decktome.com` can not read the API, so the server reads it. | The headers of a probe with an `Origin` header | 2026-10-03 |
| The server header reads `nginx`. No Cloudflare header came back. | The same headers | 2026-10-03 |
| Staff call the API "open and public (as far as reading is concerned)". They ask for a link back when the data shows in public, and they can close the API under load. | `https://archidekt.com/forum/thread/40353`, a post of about 2019 | 2026-10-03 |
| The terms forbid "automated agents or scripts" that "generate automated searches, requests, or queries to the Site". | `https://archidekt.com/terms`, effective 2018-09-07 | 2026-10-03 |
| `card.uid` is the Scryfall id of the printing. The uid of Gaea's Gift answered `bro` 182 on Scryfall. | Deck 7031486 and the Scryfall API | 2026-10-03 |
| The commander category has `isPremier` true. | Decks 7031486 and 4406647 | 2026-10-03 |
| The Sideboard category has `includedInDeck` true. A card with the categories Sideboard and Instant is in the sideboard. | Deck 3045150 (60 main, 15 sideboard) | 2026-10-03 |
| The Maybeboard category has `includedInDeck` false. A user category can also have it false. | Decks 3045150 and 4406647 | 2026-10-03 |
| The modifier of a card reads `Normal`, `Foil`, or `Etched`. | Decks 3045150 and 4406647 | 2026-10-03 |
| Artifist Acumen (`fra` 73) is a card of Reality Fracture, released 2026-10-02. The local snapshot of the owner does not hold it yet. | The Scryfall API | 2026-10-03 |

UNVERIFIED: the answer to a private deck, and the answer to a read from Cloud Run. Every probe ran from the owner's Mac.

## Moxfield

| Fact | Source | Read |
|---|---|---|
| The deck page, `api2.moxfield.com/v3/decks/all/<id>`, the v2 API, and `api.moxfield.com/v2` answer 403 with the Cloudflare page "Attention Required". | A probe of deck `VaR9P-HceECgmgm55DC7ow` | 2026-10-03 |
| The terms page answers 403 with the same page. No session has read the terms. | `https://moxfield.com/help/terms-of-service` | 2026-10-03 |
| `robots.txt` disallows `/search/*`, `/account/*`, `/collection/*`, and `/binders/*`, and not `/decks`. | `https://moxfield.com/robots.txt` | 2026-10-03 |
| Moxfield limited its API because of bad actors, and a legitimate tool can ask support for access. | `https://github.com/chilli-axe/mpc-autofill/issues/265`, opened 2024-12-04 | 2026-10-03 |
| A developer with an agent that support allowed still got a Cloudflare error on the private API. The issue is open. | `https://github.com/moxfield/moxfield-public/issues/143`, opened 2025-11-23 | 2026-10-03 |
| The menu of a deck page reads More, then Export, then Copy for Arena. | The owner, in a browser | 2026-10-03 |
| That export holds an About section with the deck name, a Commander section, a Deck section, and a Sideboard section. | `go/internal/decklist/testdata/moxfield_arena.txt`, from the owner | 2026-10-03 |

## The probes that this note does not repeat

D-470, D-482, and D-493 hold the Moxfield probes of 2026-09-02 and 2026-09-03. D-845 holds the probes of 2026-09-23.
