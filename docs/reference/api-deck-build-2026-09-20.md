# The API-only deck build, 2026-09-20

This is a dated record. It holds the method of `make api-build`, the result of the first live run, and the limits of the check. D-778 asked for it, and D-779 to D-781 set its account, its cleanup, and its cost. The roadmap entry is PR-61.

## Why this exists

The live check of #196 proved the deploy and the released chunks. It proved that the deck page holds the power counts of PR-60. It never read a built deck.

Two facts caused that hole:

- A session cannot open a browser as the owner. So no session signs in, uploads a collection, and reads a deck page.
- The sandbox of a session refuses every read under `users/<uid>` but one session by id (D-638). `scripts/read-session.sh` reads one session, and no session lists the decks of a user.

So a merge reached production, and no session had a way to say whether the app still builds a deck. `make api-build` closes that hole.

## The method

The command is `go/cmd/api-build`, and `make api-build` runs it.

1. **The token.** The command posts the email and the password to `accounts:signInWithPassword` of Identity Toolkit. The deployed project enables the email and password provider alone, so this needs no browser. The answer holds a Firebase id token, and the token lasts one hour.
2. **The header.** A `http.RoundTripper` adds `Authorization: Bearer <token>` to every call. The streamed `Chat` call carries it too, which a test proves.
3. **The import.** `CollectionService.ImportCollection` takes the raw ManaBox CSV as bytes. The run uses `go/internal/collections/testdata/manabox_collection.csv`, which holds 2,547 rows.
4. **The chat.** `AgentService.Chat` streams the turn. The first event returns the session id.
5. **The answers.** The run reads each `Question` of the turn and builds one `Answer` for it. The next request carries the answers and an empty message, which the service takes.
6. **The deck.** The stream ends with a `Deck` event when the build finishes.
7. **The read-back.** `DeckService.GetDeck`, `DeckService.ListDecks` with the session id, and `AgentService.GetSession` read what the deployed project stored.

## The answer rules

`-answers "slot=text;slot=#2;slot=decline"` names one reply for each slot. The key is the slot, and never the question id, because a question id is new in every turn.

A slot with no entry still gets an answer, so the run never waits for a person:

- A closed question takes its first option. The options are the whole answer space (D-295).
- A `no_decline` question takes its first option. The commander row shows no decline control (D-690).
- The run declines every other question. The slot closes, and the generator applies the default of the corpus (D-353).

CAUTION: the first option of the power question is bracket 1. So a run with no plan builds a bracket 1 deck, and it reads no floor of a high bracket. `API_BUILD_ANSWERS="power=#3"` asks for a bracket 3 deck. The run prints each option that it chose, so a change of the offered options is visible in the log.

## The endpoints and the keys

| Item | Value | Source |
|---|---|---|
| The API origin | `https://mtg-api-qk2ackpb3q-uc.a.run.app` | `cloudbuild/web.yaml`, substitution `_API_BASE_URL` |
| The web API key | in `cloudbuild/web.yaml`, substitution `_FIREBASE_API_KEY` | the built web app serves it to every reader |
| The project | `decktome-prod` | `.firebaserc` |
| The sign-in | `https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword` | Firebase Auth REST, read 2026-09-20 |

CAUTION: `decktome.com` serves the web app alone. `firebase.json` holds one rewrite, and it sends every path to `index.html`. No path of `decktome.com` reaches the API (D-544). So the command must call the run.app origin.

The email and the password are not in this repository, and they never join a file of it (D-639). `.env` holds both, and `.gitignore` covers `.env`.

## The result of the first live run

The run passed. It is the first proof that a session builds a deck on the deployed app with no GUI.

| Item | Value |
|---|---|
| The account | the check account of D-779, uid `n60IL9CX0NOCj3JbuI7ZrtlrgXf2` |
| The card snapshot | `2026-09-20T21:01:54Z` |
| The import | 2,471 rows, 4,316 cards, 2,547 rows resolved, 1 unresolved |
| The session | `4mmIdVLGjGXIbbPGqtE7`, 3 turns, 8 slot states |
| The deck | `4aiklNKrKHGYqyaMhN4H`, "lifegain", Commander, 1 commander and 99 main cards, bracket 1 |
| The time | 100 seconds over 3 turns |
| The read-back | `GetDeck` returned the deck, `ListDecks` found 1 deck under the session |

The agent asked four questions over two turns. Turn 1 asked the power bracket, the colors, and the budget. Turn 2 asked the commander. The command answered each one with no person at the keyboard.

The validation of the deck holds two findings. `not_owned` says the deck needs one Vito, Thorn of the Dusk Rose, and the collection holds none. `curve_summary` reads an average mana value of 2.67 over 61 nonland cards.

## What the first attempt found

The first attempt failed, and it found two defects. Both are fixed.

- **A cold instance refused the import.** Cloud Run scales to zero, and the import answered `card database not loaded yet`. The command now waits on `HealthService.Check`, which is public, until the API holds a card snapshot. `-ready` sets the limit, and it defaults to three minutes.
- **The failed target reported success.** `Makefile` sets `.SHELLFLAGS`, and GNU Make 3.82 added that variable. This machine runs GNU Make 3.81, which ignores it. So the `tee` of the recipe hid the exit code of the command. The recipe of `make api-build` sets `set -o pipefail` itself. F-160 records that every other `tee` target keeps the fault.

## The limits

- The check proves one deck and one prompt. It does not measure deck quality. `make deck-gate` measures that, and it never calls the deployed API.
- The check signs in as the check account, and never as a reader. An account with a different invite state can read a different answer.
- The answer rules choose the first option of each closed question. A different reader answers differently, and the agent then asks a different next question.
- The run costs money, so it is not a check of every merge. It answers the question "does the deployed app still build a deck", one time, on request.
- The free tests of `go/cmd/api-build` drive the whole flow against a test server. They prove the client, the answer rules, and the bearer header. They prove nothing of the deployed project.
