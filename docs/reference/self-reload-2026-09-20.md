# The first self-reload of the installed app, 2026-09-20

This document records the check of D-692 on the release of #196. F-122 is the finding that D-692 answers. The check ran on 2026-09-20 and it cost nothing.

## The rule under test

D-692 says the installed app reloads itself when a new service worker takes over. D-621 says a reader must never run an old shell against a new API. #158 merged the answer on 2026-09-13, and `web/apps/web/src/lib/pwa-register.ts` holds it.

#196 is the first merge after #158 that changes `web/`. So the release of #196 is the first release that an installed app of the release of #158 can see.

## The deploy of #196

| Fact | Value |
|---|---|
| Commit | `8eb5720` |
| Cloud Build | `32676ba4-ad8a-47c5-96a6-a6d9b1c3442a`, SUCCESS |
| Build window | 2026-09-20, 22:41:59 to 22:44:58 UTC |
| Hosting release | `60c80b8c4ad2a689`, 22:44:09 UTC, 53 files |
| The release before it | `b786ceb7f89bc4fc`, 2026-09-13 18:51:13 UTC, the release of #158 |
| Cloud Run revision | `mtg-api-00057-gbl`, from `d3d5d29` |

The API revision did not move, and that is correct. The job `deploy:api` of `.github/workflows/deploy.yml` runs only when a Go file changes, and #196 changed none.

Every chunk of the deck page changed:

| Chunk | The release of #196 | The release of #158 |
|---|---|---|
| Entry | `index-Fi2SZrwg.js` | `index-DYfo4m8I.js` |
| Deck view | `deck-view-BQeBlu0i.js` | `deck-view-C8k_znHH.js` |
| Deck stats | `use-cards-DNxWsKmD.js` | `use-cards-D0s18MpM.js` |

## The method

Firebase Hosting serves the current release alone. The assets of the release of #158 answer no request on `decktome.com` after 22:44 UTC. A live browser test of the old shell is therefore impossible after the deploy.

The check builds both releases on this machine, and it serves them from one static server:

1. Build `f75f806` (#158) in a detached worktree.
2. Build `8eb5720` (#196) in the checkout.
3. Serve the old build on `127.0.0.1:8123`, with the cache headers of the Hosting service.
4. Load the page in Chromium, and wait for the service worker to control it.
5. Swap the server to the new build. This step is the deploy.
6. Reload the page one time. This step is the reader who opens the app.
7. Wait, and give no other command.

`scripts/self-reload-check.mjs` holds the check. Playwright 1.63.0 drives Chromium, and Node 22.23.2 runs the script.

## The result: PASS

| Step | What the page showed |
|---|---|
| 1. The first load | The old shell `index-B0x9Olo1.js`, and the worker controls the page |
| 2. The deploy | The server holds the new release |
| 3. The reader opens the app | The old shell again, which is the stale load of F-122 |
| 4. After the wait, with no command | The new shell `index-DYqnsnFz.js` |
| 5. The navigations after the deploy | 2: the reader's own load, and the self-reload |

The bundle of #158 holds the code that decides this. It reads `addEventListener("activated", e => (e.isUpdate || e.isExternal) && window.location.reload())`, and it passes no refresh handler. The deployed `sw.js` of #196 holds `skipWaiting` and `clientsClaim`.

## The limits of the check

- The local build of a commit and the Cloud Build build of the same commit carry different content hashes. The build environment differs. So this check proves the source of the two commits, and it does not prove the bytes of the release of #158.
- The check runs on one Chromium version. A reader on Safari or on an older browser can behave in a different way.
- The check makes one reload. A reader who leaves the app open for hours sees the self-reload only after the browser asks for `sw.js` again.

## How to run it again

`make self-reload-check` runs the check against the commit before the newest commit that changes `web/`. The target is free, and it calls no provider. `node scripts/self-reload-check.mjs <ref>` names a different old release.
