---
name: live-test
description: Check a change on the deployed app as the test account, with no help from the owner. Three lanes - make api-build for the server path, make live-web for one turn on the screen, and make live-sweep for every screen on a desktop and a phone. Load before any check that needs a deployed chat turn, question, deck, or screen.
---

# Live test

The owner decisions are D-959 to D-961. The test account is the check account of D-779. `.env` holds its address and its password as `API_BUILD_EMAIL` and `API_BUILD_PASSWORD`.

## The rules

- Never write the address or the password into a file, a commit, a pull request, or a comment. The repository is public (D-639).
- A screenshot shows the address in the page header. Keep each screenshot under `.local`, and never commit one.
- CAUTION: A deck that a lane builds starts a live eval, because every account gets one (D-1135, D-1194). That eval can spend $3.00 and open a pull request. Name both costs in the question to the owner.
- A chat turn calls the deployed API, and the deployed API calls the real providers. So each turn costs money. A sweep with `LIVE_SWEEP_BUILD=0` sends no turn.
- The "anime" turn of 2026-09-26 cost $0.0007. A whole build costs $0.05 to $0.20. The first sweep with a build cost $0.0545.
- Ask the owner before every run, as `CLAUDE.md` tells. One answer of the owner can approve the runs of one check.
- Run a lane only after the deploy of the change. Read the build and `/readyz` first, as `docs/deploy-and-rollback.md` tells.

## Select the lane

| The check reads | Lane |
|---|---|
| The questions of a turn, the slots, or the stored session | `make api-build`, or `make live-web` |
| The text of a rendered question, its buttons, or the page layout | `make live-web` |
| A whole deck, the deck in storage, or a repair turn | `make api-build` |
| The web resend after a cold start (D-952) | `make live-web` |
| A fault that only a deploy shows, on any screen | `make live-sweep` |
| A web change after its deploy, or on a Hosting preview through `LIVE_BASE_URL` | `make live-sweep LIVE_SWEEP_BUILD=0`, for nothing |

`make api-build` and the web send the same `Chat` request. The first turn carries the collection, the message, and the pool rule. Each later turn carries the session and the answers. `make api-build` renders no screen, and it waits on `/readyz` in place of the web resend.

## The server lane: `make api-build`

1. Set `API_BUILD_OUT` to a new file under `.local`.
2. Set `API_BUILD_PROMPT` to the message of the check.
3. For one turn, set `API_BUILD_ARGS="-max-turns 1"`.
4. Run `make api-build`.
5. Read the lines `session`, `Q [slot]`, and `A [slot]` of the output.

A run of one turn ends with exit 2 and the line "no deck after 1 turns". That exit is the normal end of such a run. The run imports the ManaBox test CSV each time, and it keeps the chat and the collection.

## The screen lane: `make live-web`

1. Put Node 22 on `PATH`, then add the Node 20 bin folder of nvm for `pnpm`.
2. Run `pnpm install --frozen-lockfile` in `web/` of a new worktree.
3. Set `LIVE_WEB_OUT` to a new folder under `.local`.
4. Set `LIVE_WEB_PROMPT` to the message of the check.
5. Run `make live-web`.
6. Read `result.json` and `turn.png` in the folder.

The lane signs in on `decktome.com`, picks the first collection of the account, and sends the message. It waits until the turn shows a question, a deck, or a failure. `result.json` holds the session, the outcome, each question, and the name of each button of each question. The lane fails on a failure outcome alone.

A turn shows at most three questions, in the order of the catalog. So the message fills each slot before the slot of the check. The commander row of D-690 showed on "Build me a bracket 3 black and white tokens Commander deck from my collection with no budget. I have a commander in mind." The same message without its last sentence showed the pick row (PR-88).

`LIVE_BASE_URL` points the lane at another origin, for example a Hosting preview channel. The lane keeps no trace, because a trace records the typed password.

## The sweep: `make live-sweep`

1. Do steps 1 and 2 of the screen lane.
2. Set `LIVE_SWEEP_OUT` to a new folder under `.local`.
3. For a free run, set `LIVE_SWEEP_BUILD=0`.
4. Run `make live-sweep`.
5. Read `report.md` in the folder.
6. Open the screenshot of each screen with a fault.

The sweep walks the sign-in screens, the collection, the binder, the chat, the deck list, the deck screen, a share link, and the error states. Each screen reads five faults: a console error, an uncaught error, a failed request, sideways overflow (D-624), and a broken image.

A run with a build sends one vague message, so the agent asks its rows. The sweep answers each question with the first option, and the build ends on one deck. A free run sends no message, and it reads the newest deck of the account.

The sweep never clicks a thumb or "Report a problem", because each one writes a reader verdict (D-635). It cancels each rename and delete dialog. The sweep reads the stored share link of the deck, or makes one when the deck has none. The app has no revoke (D-1069), so that link stays live until its deck ends. With `LIVE_SWEEP_BUILD=0`, the sweep can make a public link to an existing deck. `LIVE_SWEEP_DELETE=1` deletes the deck of the run at the end.

A `flow` fault names a step that did not run, for example a control with a new name. Correct `web/apps/web/live/sweep.spec.ts` in the pull request that changed the control.

## Read the stored session

`make api-build` and `make live-web` print the session id, and `report.md` of the sweep names its deck. Read the stored session with these commands:

```bash
CLOUDSDK_CORE_ACCOUNT=<owner account> SESSION_PROJECT=decktome-prod \
  scripts/read-session.sh <session-id> <uid>
```

`make api-build` prints the uid on its `signed in` line. Without the uid, the script reads the user list, and that read takes longer.

## Record the check

- Record the session id, the lane, the time, and the outcome in the record of the pull request.
- Record the cost from the stored session. The test account is not the admin, so its screen shows no spend line (D-1148).
- A lane that fails before the send costs nothing. Record the failure and its cause.
