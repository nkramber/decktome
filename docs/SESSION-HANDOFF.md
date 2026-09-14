# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-297, `docs/owner-questions.md`, and `docs/open-questions.md`.

This file holds the ten most recent sessions. Every older record sits in `docs/reference/session-handoff-archive.md`, word for word. The first narratives sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. Put `~/.nvm/versions/node/v22.23.2/bin` on the PATH before `make verify`. On this machine `nvm use` reports the change and does not make it, so prepend the path yourself.

## RESUME HERE (2026-09-14)

**The checkout.** `main` is `bce26d5`, which is pull request #174, or a later merge. The branch `plan-m17-mana-wincon` holds this hand-off and the plan of M-17, PR-52, and PR-53, as an open pull request. Run `make where` before you touch anything. Never commit on `main` (D-583). Answer the review of `gitar-bot` before you ask for a merge, and a pull request of documents alone waits for it too (D-637, D-679).

**The owner answered OQ-82: a bracket promises its power in both directions** (D-693 to D-695). M-14 ran bracket gate run 2 on `320fcbb` for $1.37 (D-694). Five of the six decks at brackets 4 and 5 miss the floor of the mana on turn four. With the Game Changer flags of PR-47, the judge agrees on 3 of 15 decks, and it still names combos from memory (F-126). #161 merged PR-46, #162 merged PR-47, and #163 merged M-15 (D-697 to D-699). PR-45 splits into PR-45a and PR-45b, and #164 merged the plan (D-701 to D-704).

**#165 merged PR-45a: the build cuts a card its bracket forbids** (F-125, D-702). The content check names each forbidden combo, mass land denial card, and extra-turn card past the limit. The build cuts one card of each forbidden combo and each other forbidden card, and a basic land fills each slot. A commander, a locked card, and a card that a revision keeps never leave. For a combo, the cut prefers the card in the most combos, then a card outside the precon, then the lower shortlist score.

**Bracket gate run 3 reads no content violation on 9 of 9 decks.** It ran prompts 1 to 9 on `6d61700` for $0.7355 over 538 seconds. The cut removed Polyraptor from deck 4 and Akki Battle Squad from deck 9. The verdict reads FAIL on the judge bar alone, at 3 of 9 and 2 of 3 at bracket 3. The first finding text said that the bracket forbids the card itself, so the finding now names the combo. The review found F-128, and #166 merged its fix (D-705).

**#166 merged PR-48: the mana pass counts the command zone as held** (F-128, D-705). The spell steps of the pass read the 99 alone, and the pool holds each commander. On the old pass a unit test reads three steps that add the commander, and a build test reads 2 copies of it.

**#167 merged PR-49: every tool builds the shortlist of the app** (F-129, D-706, D-708). The gates passed no top-list rate, so they measured shortlists that the app never builds. A parity test fails on the old tools at five calls. The dry comparison on model `20260910T012734Z` finds the rate bringing in up to 45 cards, with more fast mana and Game Changers.

**#169 merged PR-45b: the build reaches the power of brackets 4 and 5** (F-130 to F-134, D-709 to D-715). Brackets 4 and 5 take floors for tutors, fast mana, and Game Changers. The prompt names each floor, and each shortlist line marks the power cards. At brackets 4 and 5 the shortlist pins each power card at a top-list rate of 0.3. A pinned card skips the cap of its role and adds to the total. A deck that misses a floor names the gap and the cards that close it in its summary.

**The session found three defects on the way, and PR-45b fixes each.** The role caps blocked the power cards (F-131). The first pin took places under the total, and 52 fixing lands left the lists (F-132). The mana pass step of a cheaper card read no job (F-133). The review of `gitar-bot` found a fourth: the land cap read no pin, and `2ed2b06` sets the pinned lands aside.

**The paid runs of PR-45b cost $2.04, and the judge moved from 0 of 6 to between 4 and 6 of 6.** Bracket gate run 6 on the final code reads 4 of 6, and its second judge lane reads the same decks at 6 of 6. Neither miss of run 6 held on the second read, so the owner closed F-134 as judge noise (D-715). Deck gate runs 20 to 22 read PASS on prompt 3.

**Quality gate run 22 reads PASS on the local meta store of 2026-09-14** (D-499, D-565). The fit took 124 seconds and called no provider. The Commander precon bar reads 0.96 of 784, and the synergy axis reads 0.85 of 228. Built decks graded bad rise from 8 to 9 of 25, and the judge agreement stays at 9 of 25. Deck 21, the Hobbit family, moves from baseline to bad. Deck 7, Modern burn, moves from baseline to typical.

**The run is no quality tuning item, so its numbers stand as information** (D-689). The low EDHREC count needs no finding. `MergeLists` answers the count of new keys, and a re-read replaces a list already stored under its key. So the weekly read counted 3 new lists on 2026-09-07 (D-566) and 4 new lists on 2026-09-14.

**#172 merged PR-50: a card name matches without its accent** (F-135, D-716). The owner typed "Grima" as a commander, and no card matched, because the card is Gríma, Saruman's Footman. The package `cardname` holds an exact key and a folded key, and every name match reads the exact key first. A folded key that two card names share finds nothing, and none of the 39,967 names of the snapshot share one. Guardrail 4 reads the folded key as exact now. Quality gate run 23 matches run 22 in the three numbers and in every deck grade.

**Gitar found one performance defect in #172, and `6394d16` fixed it.** The fold ran NFKD on every row of a binder search, at about 270 ns against 39 ns for the old key. A name of ASCII bytes alone now skips the decomposition, at 50 to 88 ns, and a test holds the equality.

**Deployed session `z1hshyY6Npig1FN2NuV7` proves PR-50 and PR-45b.** The owner wrote "Grima as commander", and the which-card row offered Gríma, Saruman's Footman and Gríma Wormtongue. The deck `sFLbEUuKzI0QyLPft0zI` stores the bracket 5 floors: 4 tutors, 6 fast mana, and 8 Game Changers. It holds 0, 3, and 1, and the summary names the cards that close each gap, each one to buy. The curve line reads 2.39 over 66 nonland cards, as the list does, with 33 lands. The session cost $0.057 over 6 calls.

**The same session found F-136 and F-137.** Turn 1 asked whether "that card" leads the deck or sits in the 99, and the first message already named it as commander. `applyFacts` set the named-card fact again after the which-card row took the name (F-136). The owner chose PR-51, a catalog condition, with a question gate run (D-717). The summary also names "Gríma's opponent-milling plan", and no card of the deck makes an opponent mill (F-137). The owner chose a record (D-718).

**#173 merged PR-51, and deployed session `l9x5bFTcFgpZ4DOcQpxU` proves it** (F-136, D-717). The role row reads `"commander_unresolved": false`, so it waits while the which-card row holds a name. Question gate run 50 read PASS at 73 of 74, as run 49 did, for $0.1938. The owner wrote "Grima as commander" on revision `mtg-api-00049-4lj`. Turn 1 asked the power and the colors, and no role question. Turn 2 offered both Gríma cards.

**The owner gave a written review of the Gríma deck, and the session checked each claim** (D-719 to D-722). The deck holds no win condition, a weak mana base, and cards whose conditions it does not meet (F-138 to F-140). The claim of a pump shelf is mostly wrong, and the missing staples come from the owned pool. The deck also holds none of the 14 owned cards tagged mill-opponent on a theme of opponent mill (F-141). The owner chose M-17 first, then PR-52 for mana, then PR-53 for win conditions. `docs/reference/deck-review-grima-2026-09-14.md` holds every claim and its evidence.

**The app is live on `decktome.com`.** A merge to `main` deploys itself on Cloud Build (D-584, D-586). #158 is the newest web deploy, and `deploy-web` released it at 18:52 UTC on 2026-09-13. #173 is the newest API deploy: `deploy-api` finished at 20:50 UTC on 2026-09-14, and revision `mtg-api-00049-4lj` serves `api:54b07dd`. Both jobs run `worker:54b07dd`, and `/readyz` answered ok. The deploy triggers read `go/**`, `docker/**`, and `web/**` alone, so a merge of documents alone starts no build.

**Session `X4JfbXzMw4U5A4gEeOaE` reads right on every check, and the owner closed OQ-79** (D-691). The owner built an owned-only lifegain Commander deck at bracket 3 on 2026-09-13 at 04:16 UTC. The first message named Sidequest: Catch a Fish. The offer named Aerith Gainsborough, Hope Estheim, and Aerith, Last Ancient, as a free local run of the offer code predicted. The reader owns one copy of each. The reader owns none of the three commanders that an offer with any card allowed names first. The deck marks the commander and every card owned.

**The same deck proves PR-43 and PR-44 on the deployed app.** The curve line reads 3.03 over 63 nonland cards, and the stored list gives 3.0317 over 63, with 36 lands. Sidequest: Catch a Fish is the one card with more than one face, and it counts as a nonland card. The deck grades baseline on model `20260912T061545Z`. The session cost $0.056 over 6 calls, and the build took 35 seconds in one call.

**F-121: the commander row showed "Suggest one" beside "You decide"** (D-690). The two controls do different things. "Suggest one" asks for three names, and "You decide" lets the build pick with no offer. The owner wants "Suggest one" alone on that row. The catalog row carries `no_decline: true`, `Question.no_decline` carries it to the UI, and the question card shows no decline control. The pick row keeps "You decide". Old sessions keep the control, and no catalog text changed, so no question gate run is due. #157 merged it. Deployed session `vY1lCRtl64uwFObznCZ9` stores the commander question with `noDecline: true`, and the pick row without it.

**F-122: the installed app ran the old shell for one load after the deploy of #157** (D-692). The owner's first load still showed "You decide" beside "Suggest one", while the API stored the flag. `registerType: "autoUpdate"` makes the new service worker skip waiting and claim the page. The injected `registerSW.js` registers the worker and never reloads, so the old shell ran that load. The owner chose a reload. `src/lib/pwa-register.ts` imports the plugin's register module, which reloads the page when an updated worker activates. An unsent draft lives in memory, so the reload drops it.

**#158 merged the fix of F-122, and `deploy-web` released it at 18:52 UTC.** The live `index.html` loads no `registerSW.js`, and `sw.js` precaches the `workbox-window` chunk. The main chunk holds the reload.

CAUTION: the fix of F-122 reloads only a page that already runs it. The first load after its own deploy still runs the old shell. The next deploy after that is the first one that reloads the page by itself.

CAUTION: the application default credentials of this Mac failed on 2026-09-13 with `invalid_rapt`. `make feedback-list` fails, and so do the harvest and the backfill, which read the same credentials. `scripts/read-session.sh` and `gcloud logging read` still work with `CLOUDSDK_CORE_ACCOUNT`. The owner runs `gcloud auth application-default login` to repair them.

**Deck gate run 19 and question gate run 49 read PASS** (D-686). Run 19 passes all 25 decks with one repair turn, and it is the decks baseline now. Its grade rows fell on most decks, because the local stored model `20260910T012734Z` graded them. Every curve line matches its list, so PR-43 holds on real builds. Run 49 met every expectation, and no conversation missed, so PR-42 ran no rerun.

**#153 merged PR-44, and Cloud Build deployed it: a card's types follow the rules of its layout** (F-120, D-687 to D-689). The index merged the types of every face, so Legion's Landing counted as a land. A card with more than one face takes its front face now, and split and modal double-faced cards keep every face. Quality gate run 21 reads PASS: 8 built decks grade bad against 9, and the judge agreement reads 9 against 10. The owner reads guardrail 15 as a rule for quality tuning items, so this correctness fix reports the numbers as information (D-689).

**The owner's Commander build reads right** (D-678). Session `OFMnk7Tv2zkK8xfAwXxB` built a bracket 5 treasure deck led by Smaug the Magnificent, from an owned-only pool. The grade reads typical, and its three reasons name the ladder. No rules check flags the deck: 33 lands, an average mana value of 2.29, and red sources at 1.72 of their need. The reader owns the commander. Session `X4JfbXzMw4U5A4gEeOaE` later gave OQ-79 its harder case, and the owner closed it (D-691).

**The read found F-119, and #149 fixed it** (D-684). The curve line of that deck reads 2.79 over 68 nonland cards, and the stored list reads 2.29 over 66. `generate.go` ran the rules engine before the mana pass of PR-33, so every rules finding read the list before the pass. #149 merged PR-43, which runs the pass first, and Cloud Build deployed it at 06:25 UTC. The pass adds only shortlist cards, so the fault is stale text alone.

**#148 merged PR-42, and Cloud Build deployed it** (F-112, D-671). A counted conversation with a miss plays two more times inside the question gate run. The document and the run file carry its miss rate. The first play still sets the verdict. No whole run shows a rerun yet.

**M-13 measured the date cut of PR-41, and the owner closed PR-41** (F-118, D-682, D-683). A free count found that the top lists hold Hobbit cards at the rate of older sets. No cut of 30 or 90 days moves a built deck to another tier, and each cut makes one number a little worse. `docs/reference/m13-date-cut-fits-2026-09-11.md` holds the count and the fits.

**#144 merged PR-40, and Cloud Build deployed it** (D-678). The build finished at 18:44 UTC on 2026-09-11. Revision `mtg-api-00036-bkn` serves `api:38ddde6`, and both jobs run `worker:38ddde6`. `/readyz` answered OK. Gate run 20 reproduces the pick of M-12: 9 built decks graded bad against 16, and the judge agreement at 10 against 9.

**PR-40 grades a Commander deck by three rules checks** (M-12, D-676 to D-678). A deck grades bad when one check fails:

- The lands fall 10 or more short of the need of Karsten's formula.
- The spells average over 4.4 mana.
- The sources of the worst color sit under 0.75 of their need.

Any other deck takes the tier of the fitted ladder. The score keeps the fitted model, so every bar reads as gate run 19. `go/internal/quality/rules.go` holds the checks, and `docs/reference/m12-rules-detector-fits-2026-09-11.md` holds the fits.

CAUTION: a grade stored on a deck built before the deploy keeps its old tier until a rebuild or a revision. The rules checks are code, and the stored model holds none of them.

**M-10 measured the softmax scorer again, and the owner closed PR-38** (D-680). The scorer drops the Commander precon bar to 748 of 788, one pair short. It grades 14 built decks bad against 9, and the owner's grades match on 2 decks against 5. **A quality item now passes only when neither reader number gets worse** (D-681, guardrail 15).

**The deployed model.** The meta job refits the model on every run (`go/cmd/worker/meta.go`). Its run of 2026-09-11 at 06:00 UTC finished at 06:41 UTC, before the deploy of PR-40, and stored `20260911T063437Z`. The run of 2026-09-12 at 06:00 UTC ran on `worker:6d9dacf`, the first run after the deploy of PR-40. It succeeded at 06:22 UTC and stored `20260912T061545Z`, fitted on 38,395 lists and 909 commanders. The Commander accuracy fell from 0.644 to 0.542, as gate run 20 predicted. The holdout accuracy reads the tier, and the tier reads the rules checks now.

**The meta run of 2026-09-13 is the first fit on the types of PR-44.** It ran on `worker:b53fbb6`, succeeded at 06:17 UTC, and stored `20260913T061119Z`. It fits 39,209 lists and 915 commanders. The Commander accuracy reads 0.547, against 0.542 the day before. The mtggoldfish source failed on a storage write with a connection reset. It read 56 lists, against 128 to 148 on the three runs before. The job still fitted and stored the model.

CAUTION: the deployed Standard fit reads its cross share of baseline over bad at 0.685 on 2026-09-13. The three days before read 0.690, 0.750, and 0.737. Local gate run 20 reads 0.92. The share is information and no bar, and the rules checks of PR-40 read Commander alone. The deployed store holds 1,935 synthetic Standard copies, against 370 in the local store.

**M-11 is done, and no detector variant meets its gate** (F-115, D-675). Every disputed deck grades bad in all fifteen fits. The grade adds the detector probability to the bad rung of every deck, so the flag never decides the tier. The owner chose the rules detector first, and M-12 fitted it. `docs/reference/m11-detector-variants-2026-09-11.md` holds every fit, and `.local/m11/` holds the patches.

CAUTION: the `decktome` gcloud configuration named the Wallabee account and project on 2026-09-10, so `use_decktome` put the shell on the wrong project. `scripts/read-session.sh` reads `SESSION_PROJECT` and the account of `CLOUDSDK_CORE_ACCOUNT`, so an environment override reads `decktome-prod` with no change to the configuration. The owner has the commands to repair the configuration. A read on 2026-09-12 at 19:24 UTC found the configuration unchanged.

**The commander offer never shares a turn with a theme or colors question** (F-111, D-669). One turn asked both and offered three commanders, so the offer ignored the answers it ranks on. A wait for the answers failed question gate run 45, because a reader who skips a question got no offer. #139 merged it, and session `ze0Im17gZlFIBRyn7k7Z` on the deployed app asked the theme and the colors alone on turn 1.

**Gate runs 46 and 47 tripped two latent gaps, and #139 closes both** (F-112 to F-114, D-670). A delegation that names its slot declines no other key, and an occasion does not fill the theme. The gate keeps a bar of zero misses, and PR-42 reruns a missed conversation and reports its miss rate (D-671, #148). Gate run 48 reads PASS with both guards.

**The commander offer ignored the bracket of its own turn, and #134 fixed it** (F-104, D-662). The turn built the hints from the restored slots, and the picked power option landed after that. The pick row fires on the turn the power answer arrives (D-631), so the bracket 5 signal ordered almost no deployed offer. Session `vDzEKDPnRZntlyhz8Lqg` read Lotho, Ghalta, and Peregrin Took. The hints take the bracket after the answers now, and a test holds it. Session `ze0Im17gZlFIBRyn7k7Z` of 2026-09-11 confirms it: the offer named Vivi Ornitier, Aang, and Hapatra.

**PR-39 is merged as #135** (D-659 to D-661, F-102, F-103). The source count reads a reliable tap ability alone, and the balance phase trades basics toward the color that falls short. The measurement of D-661 is done. Over run 18 the pass moves four decks. The worst color rises on all four and falls on none, and deck 22 gains two Forests. `docs/reference/pr39-manapass-run18.md` holds the lane.

**Gate run 19 measured PR-39 on the quality model, for nothing** (D-663, F-105). The feature `color_sources` reads the count PR-39 changed, and no gate read the model before. Every bar holds, and the judge agreement rises from 8 to 9 of 25. A control fit on the parent commit matches run 18 in every number. **An item that changes a feature of the quality model reports the three numbers now** (guardrail 14, D-664).

**PR-37 is merged as #127, and gate run 18 reads PASS** (D-653). The synergy check drops a Commander copy whose break lowered the synergy feature by less than 0.10 standard deviations. The Commander synergy axis reads 0.84 of 232, and the precon bar 0.95 of 788. No earlier run of the quality gate in the repository reads PASS.

**CAUTION: the PASS of run 18 came from the population of the bar, and not from a better model.** Over run 18 the built decks graded bad stayed at 15 of 25, and the judge agreement at 8 of 25. Run 19 reads 16 and 9, from the source count of PR-39 (D-663). Run 20 reads 9 and 10, from the rules checks of PR-40 (D-678). 36 of the 232 synergy pairs still lose on run 19, and those are the true gap of the axis.

**CAUTION: the check reads Commander alone** (D-653, F-96). In Modern the synergy break does not lower the synergy feature in 825 of 845 pairs. A check there dropped 78 percent of the Modern synergy copies, and three decks the judge reads as bad graded higher.

**The premise of D-652 was half right.** The check drops 157 of 389 Commander pairs from every band of precon strength, and not only the 54 weak pairs M-8 named. `docs/reference/pr37-synergy-materiality-2026-09-10.md` holds all six floors.

**The gate reports the three numbers of D-648 for free.** `make quality-gate` grades the decks of deck gate run 16 with the model it fits, and it reads the judge's tiers from judge lane run 5. It stores no model.

**The owner graded the ten disputed decks** (D-659). The judge bar is the right target, and the detector reads built decks wrong. PR-40 merged the rules detector of M-12 (D-678), and the owner closed PR-38 after M-10 ran again (D-680). The review sheet is `.local/review/disputed-decks-2026-09-10.md`, beside a CSV, and it stays out of git. Its notes on decks 11 and 15 read the fit of 2026-09-07, and the ladder of gate run 19 reads both decks baseline. `docs/reference/m10-softmax-scorer-2026-09-10.md` holds the fits.

**Three items aimed at the reader gap failed first**: PR-29, M-9, and PR-30 (F-94, F-98, F-99). The owner parked PR-35 and dropped PR-30.

**CAUTION: two features never get built again** (F-94). `casual_unseen_share` moved the target axis by nothing and cost three other bars. `casual_pair_share` won one pair of 389, because the fit split it against `casual_synergy` with opposite signs. `docs/reference/pr29-casual-corpus-2026-09-10.md` says why for each, and PR-35 builds neither one.

**Every quality item reports three numbers** (D-648): the precon bar, the count of built decks graded bad, and the tier judge agreement. **The owner drops an item that moves none of the three.** PR-29 moved the first up and the other two down, and that is why it closed.

**The whole feedback loop stands**, as #117, #121, and #122. **No live cycle ran yet.** It needs the owner's word, `AUTOTUNE_FIXER_CMD`, and a harvest whose verdicts carry a snapshot.

**PR-34 is merged as #123.** The app reads the format of an upload out of the file, and Moxfield reads. OQ-80 asks whether a proxy counts as a card the reader owns.

**The whole PR-22 gate holds** (D-633). The measured cost at idle reads $9.54 a month gross, and the hourly snapshot tick of D-634 leaves $3.81.

**What waits on the owner.**

- The merge of this documents pull request, after the review of `gitar-bot`.
- A look at the first commander question after a load of the app, in a session that names no commander (next step 2).
- OQ-67, the Stage B channels.
- OQ-77, the blocking function of Identity Platform.
- OQ-80, a proxy and the pool.

**CAUTION: a local `make verify` is not the whole story.** It read green for weeks while shellcheck failed (F-89, D-641). Every pull request runs the workflow now.

## How to resume

1. Run `make where`. It prints the branch, the tree, and the state of the branch's pull request.
2. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
3. Make a branch from `main`. Never commit on `main`, and never push to it (D-583). The owner merges (D-585). Run `make hooks` one time in a fresh checkout.
4. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
5. Run `make verify`. It runs every check the verify workflow runs, on this machine, for nothing. Put Node 22.23.2 on the PATH first. The pull request runs the same jobs on Actions (D-639).
6. Do "Next steps, in order" below. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
7. Open the pull request, then wait for the review of `gitar-bot` and answer it (D-637). A finding with merit takes a change, a commit, a push, and a reply. A finding with no merit takes a reply that says why, and you resolve it. Tell the owner when the pull request is ready to merge. Gitar is the only review this repo asks for. A pull request of documents alone waits for the review too (D-679).
8. Before you end, update this file. Move the oldest session to the archive when the count passes ten.

A second Mac: `docs/setup-second-mac.md` holds what to carry, what to install, and how to prove the machine.

Fifteen things a fresh session gets wrong without this file.

- `make bracket-gate` runs its tool in `go/`, through `go -C go`. A relative `-rejudge` path then points inside `go/`, so pass an absolute path.
- Twelve targets and two loop scripts spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before each run. `make autotune`, `make feedback-loop`, `make feedback-loop-dry`, `make feedback-triage-dry`, `eval sweep -dry`, and `go run ./cmd/bracket-gate -sweep` are free.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docker compose up` needs provider keys in `.env` now, and fails fast without them (D-267). `make dev` still starts with no keys, and the fake builds the fixture deck of PR-23 for any first message (D-552).
- The eval and the agent share a model. Every ratio it reports is a floor, not a measurement (D-136).
- The judge is noisy: two runs of identical code move up to nine bad questions (D-230). One run proves nothing on its own.
- The sweep's estimate of a step is the cost of its last run file. The deck gate estimate read $2.54 from run 14, and run 16 cost $3.79 with the plan judge and 12 repair turns. A cap set from the estimate stops the sweep before its last step.
- Many local fits and verify runs fill the disk. On 2026-09-11 `make verify` failed at link time with 258 MiB free, when the Go and Docker build caches held about 50 GB. Run `df -h /System/Volumes/Data` before a long run, and ask the owner before you clear a cache.
- Gitar sometimes deletes its summary comment and posts a new one with a new id. Find its newest summary by author, and read the review threads. A green Gitar check does not prove that no finding is open.
- Gitar can post that it paused automatic reviews, because the trial's processing ran out for the period. On #150 to #155 the note came beside a full review each time, and no trigger ran. When the note comes with no review, comment `Gitar review` on the pull request. Answer the manual review as usual (D-685).
- The owner can merge a pull request before its Gitar finding gets an answer (#154). The pre-commit hook then refuses a commit on that branch. Carry the fix to a new branch from `main`, and reply on the old thread with the new pull request.
- A background shell command starts in the directory the session left. On 2026-09-12 a `make` target ran in `go/` and found no rule, so head every command with an absolute `cd`.
- A live check of a web change reads the stored session and the deployed chunk, and not the screen alone. The service worker served the old shell for one load after a deploy (F-122). The protobuf-es code holds each field name in base64, so search a chunk for a property name such as `noDecline`.
- `Pool.Names` sorts the pool by the alphabet, and the shortlist groups its cards by role. Neither order ranks a card. `Pool.Score` holds the shortlist score (D-702).

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus match the 2026-08-07 text. The session of 2026-09-12 read rules 709.4c, 710.2, 712.8a, 712.12, 715.2, and 722.2a in that file for PR-44.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when that file holds no future date. This is by design.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link). The content rules per bracket in `brackets.json` and the Spellbook thresholds were read 2026-09-02, and the Karsten tables are the 2022 articles, read 2026-09-02 (`docs/reference/bracket-profile-2026-09-02.md`).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260904T210157`, the newest of 13. Read 2026-09-13, when it flagged 53 Game Changers. Check every card fact against it.
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The max output per provider in `llm/client.go`, verified 2026-08-29. The OpenAI rows are unverified by anyone but the owner. The judge role runs on Opus 5 (D-430).
- The deck gate upgrade probe cost $0.32 for 2 prompts, 7 calls, and 266 seconds. A partial run reads its own item bars and never stands as the gate (D-526).
- Run cost, read from the run files on 2026-09-09. The question gate cost $0.190 on run 42, $0.193 on run 43, and $0.194 on run 44, over 18 to 21 minutes. Runs 45 to 48 of 2026-09-11 cost $0.1925, $0.1920, $0.1380, and $0.1923, over 20, 22, 19, and 21 minutes. Run 49 of 2026-09-12 cost $0.1935, over 18 minutes. Run 50 of 2026-09-14 cost $0.1938, over 22 minutes. The deck gate upgrade probe cost $0.32 for 2 prompts. The eval cost $0.092 on run 34 and $0.096 on run 35, over 12 to 14 minutes. The deck gate cost $2.71 on run 18, over 37 minutes, and $2.75 on run 19 of 2026-09-12, over 36 minutes. The revise gate cost $1.07 on run 9. The bracket gate judge lane cost $0.28. Bracket gate run 2 of 2026-09-13 cost $1.3720 over 925 seconds, for 15 builds and the judge. Its second judge lane cost $0.2634 over 128 seconds. The M-15 calibration lane cost $0.3283 over 149 seconds, for 21 decks. Bracket gate run 3 of 2026-09-14 cost $0.7355 over 538 seconds, for 9 builds and the judge. Bracket gate runs 4, 5, and 6 over prompts 10 to 15 cost $0.4979, $0.5247, and $0.5619, over 424, 442, and 564 seconds. The second judge lane of run 6 cost $0.0856 over 35 seconds. Deck gate runs 20, 21, and 22 over prompt 3 cost $0.1321, $0.1174, and $0.1159, over 94, 97, and 106 seconds.
- The backfill of 2026-09-09 read one user with a record to seed: 1 collection, 1 thumbs up, and 2 thumbs down. It counted no deck and no chat, because the reader deleted both (D-635).
- The feedback store holds 3 verdicts on 2026-09-09, and every one predates the snapshot of D-635. `make feedback-list VERDICT=` reads both verdicts now (F-87). A collection group query over it needs an index for its shape, and the store holds "verdict ascending, created_at descending" alone.
- The deployed schedules, read 2026-09-09 and again on 2026-09-11: `mtg-snapshot-schedule` at `0 * * * *` (D-634) and `mtg-meta-schedule` at `0 6 * * *`. Both read ENABLED. The API service holds minScale 0, so it scales to zero. No billing export exists, so no command reads the billed spend.
- The deployed API, read 2026-09-14 at 20:50 UTC: revision `mtg-api-00049-4lj` on image `api:54b07dd`, from #173. Both jobs run `worker:54b07dd`, and `/readyz` answered ok with a card snapshot of 2026-09-14 09:01 UTC. `deploy-api` ran from 20:45 to 20:50 UTC.
- The deployed web app, read 2026-09-13 at 18:54 UTC: the release of #158, from `deploy-web` at 18:52 UTC. `index.html` loads `assets/index-DYfo4m8I.js` and no `registerSW.js`, and `sw.js` precaches the `workbox-window` chunk.
- The deployed quality model, read 2026-09-14: `20260914T070904Z`, from the meta job that started at 06:02 UTC and ended at 07:13 UTC. It fits 43,182 lists and 1,517 commanders. Its Commander fit reads `immaterial` 239 and accuracy 0.554, and its Standard fit reads a cross share of 0.667. The job read the weekly EDHREC pass, 2,077 pages and 4 lists. The mtgo source read 3,094 lists with 58 fetch errors. The mtggoldfish source read 155 lists with no failure. The mtgjson source read no list, because its deck list version differs from the stored table, as on 2026-09-12 and 2026-09-13.
- The Karsten land article of 2022-07-29, read 2026-09-11 through `infinite-api.tcgplayer.com/content/article/<id>/`, because the page draws its text in the browser. `docs/reference/m12-rules-diagnostic-2026-09-11.md` holds the formula, the error, and the cheap rules.
- Baselines, in `docs/reference/eval/baselines.json`: questions is run 42, decks is run 19, revise is run 9. The generate prompt reads version 13 since PR-45b, and run 19 read version 12. So the next whole deck gate run reads a new prompt against run 19. Run 50 is the newest whole questions run, and it reads PASS. Runs 45 to 47 read FAIL (D-669, F-112), and run 43 records the regression of F-84. `make eval-check` compares the newest whole run of a suite against its baseline. The quality gate has no baseline row, and run 23 is its newest run.
- Toolchain on this Mac, read 2026-09-11: Go 1.27.1, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, and Java 17.0.20.1. Playwright is 1.63.0 with its Chromium headless shell (`playwright install chromium`). `go.mod` asks Go 1.27.0 or newer, and vitest is 5.0.0 since #133.

## Next steps, in order

1. **Run M-17, then build PR-52, then PR-53** (D-719 to D-722). M-17 costs nothing, and it sets the floors of both items. It counts finisher tags and land classes in real lists. It also replays the Gríma request with the collection of the owner. Keep the exported collection out of git. PR-52 ranks lands by quality, and PR-53 adds a win-condition target. The power pass after the build still waits (D-704), and F-137 stays a record (D-718).
2. **Read the first commander question on the app after one more load** (D-690). It waits for the owner. The server half holds: session `vY1lCRtl64uwFObznCZ9` stores the flag. The question must show "Suggest one" and no "You decide", and the pick row must still show "You decide". Session `z1hshyY6Npig1FN2NuV7` named its commander, so it showed no such row.
3. **Watch the first self-reload on the next web deploy** (D-692). The live release of #158 passed its check on 2026-09-13 at 18:54 UTC. An installed app that loaded that release must reload by itself when the next release activates. Read it on the next merge that changes `web/**`.
4. **The next collection platform, when the owner names one** (F-91). Five are left: Archidekt, Deckbox, Delver Lens, TCGplayer, and Helvault. Each one takes a real export and never a column list. `docs/reference/pr34-collection-formats-2026-09-10.md` holds the shape.
5. **The live half of the feedback loop has no run yet.** Three things want a measurement: the judge lane of the triage, one live fix cycle, and one review round. All three need the owner's word, and the cycle also needs `AUTOTUNE_FIXER_CMD` and a harvest whose verdicts carry a snapshot.
6. **Deck gate run 19 is the decks baseline** (D-686). Its grades read the local stored model `20260910T012734Z`, and the deployed app reads a newer one. The next whole run compares against run 19, and it costs about $2.75, so ask the owner first.
7. **The owner parked PR-35 and dropped PR-30** (D-655, D-656). M-9, F-97, and F-99 found no case for either one.
8. **PR-36, the reader's verdict as a quality signal** (D-651). It waits for verdicts.
9. **PR-26, the return channels**, waits on OQ-67.
10. **PR-42 is merged as #148** (D-671). Question gate run 49 missed no conversation, so it ran no rerun. Any miss still fails the run, and each miss joins the finding register.

CAUTION: `make revise-gate | tee` hides the exit code. Read the verdict line of the document, never the exit code of a pipe.

CAUTION: `make verify` runs `eval-check`, and `eval-check` reads the newest whole run of each suite. A committed FAIL run turns verify red until a newer whole run passes.

The CI step "fake gcs tests" ran no test until 2026-09-10 (D-658). Its filter matched no test name, so it passed as a no-op. It runs `scripts/gcs-check.sh` now: the script seeds the fake GCS from the trimmed snapshot and fails unless `TestLiveFakeGCS` passes by name. `make gcs-check` runs the same script.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

## The ten most recent sessions

### 2026-09-14j: #174 merged, the review of the Gríma deck, and the plan of M-17

**The owner merged #174.** It changed documents alone, so no build ran.

**The owner gave a written review of the Gríma deck, and the session checked each claim.** The session read the stored deck, the card snapshot, the stored collection of the owner, and two research passes over the code. The review holds on the win condition, the mana base, and the unmet conditions. It fails on the pump shelf, and the missing staples come from the owned pool. The check also found F-141: the deck holds none of the 14 owned mill cards.

**The owner chose the plan** (D-719 to D-722). M-17 measures first, then PR-52 improves the mana, and then PR-53 adds a win-condition target. The session changed no code, and it ran no paid target after #173.

### 2026-09-14i: #173 merged, and the deployed app asks no role question

**The owner merged #173, and Cloud Build deployed it.** `deploy-api` ran from 20:45 to 20:50 UTC, and revision `mtg-api-00049-4lj` serves `api:54b07dd`. Both jobs run `worker:54b07dd`, and `/readyz` answered ok.

**Deployed session `l9x5bFTcFgpZ4DOcQpxU` proves PR-51** (F-136, D-717). The owner wrote "Grima as commander", and the lookup found two cards. Turn 1 asked the power and the colors, and no role question. Turn 2 offered both Gríma cards.

### 2026-09-14h: #171 and #172 merged, the live Grima check, and PR-51

**The owner typed "Grima" as a commander, and no card matched.** The card is Gríma, Saruman's Footman. The owner chose a fold in every name match, with the exact key first (D-716). The session built PR-50 on the branch `fix-name-accent-fold`, and #172 merged it. Gitar found one performance defect, and `6394d16` added a fast path for ASCII names.

**The owner merged #171 and #172, and Cloud Build deployed #172.** `deploy-api` finished at 19:19 UTC, and revision `mtg-api-00048-5zm` serves `api:284b324`. Both jobs run `worker:284b324`, and `/readyz` answered ok. #171 changed documents alone.

**Deployed session `z1hshyY6Npig1FN2NuV7` read "Grima" right, and it proves PR-45b.** The which-card row offered both Gríma cards. The deck at bracket 5 misses all three power floors, and its summary names the cards that close each gap. The read found F-136 and F-137. The owner chose PR-51 with question gate run 50 (D-717), and a record of F-137 (D-718).

### 2026-09-14g: #170 merged, quality gate run 22, and the app check of PR-45b

**The owner merged #170.** It changed documents alone, so no build ran. Gitar approved it beside the note that it paused automatic reviews, and it named no finding.

**Quality gate run 22 read PASS on the local refresh of 2026-09-14, for no cost.** Built decks graded bad rose from 8 to 9 of 25, and the judge agreement stayed at 9. The EDHREC count of 4 lists looked low, and `MergeLists` shows that it counts new keys alone.

**The owner chose the app check of PR-45b.** The other choices were a whole deck gate run, a plan of the power pass, and a queued question. The owner builds a bracket 4 or 5 Commander deck, and the session reads the stored session and the deck.

### 2026-09-14f: #169 merged, and the documents read the state before a context wipe

**The owner merged #169 and asked for every document to read the current state** before a context wipe. The session changed no code, and it ran no paid target.

**Cloud Build deployed #169.** `deploy-api` finished at 17:50 UTC, and revision `mtg-api-00047-76c` serves `api:7d997f1`. Both jobs run `worker:7d997f1`, and `/readyz` answered ok. No web build ran, because #169 changed no file under `web/**`.

**The refresh.** The roadmap, the resume section, the next steps, and `CLAUDE.md` read the merge. The record of 2026-09-13 moved to the archive.

### 2026-09-14e: PR-45b, the pin, and six paid runs

Merged as #169.

**The owner merged #168 and asked what comes next.** The session built PR-45b on the branch `pr45b-power-floors`. The owner answered the gap note (D-709), the pin (D-710), and the default of the rate (D-711). The owner also answered two defects on the way (D-712, D-713). The local meta refresh finished and stored `20260914T154223Z`.

**The sweep found F-131, and the owner chose a pin.** The role caps blocked the power cards, so only weight 1 met every floor, and 138 on-theme cards left the lists. A pin lets a power card at the keep rate skip the cap of its role. At weight 0.1 and a keep rate of 0.3, all 18 floors hold.

**Bracket gate run 4 found F-132.** The first pin took places under the total, and 52 fixing lands left the lists. The Korvold and Najeela decks of run 4 held mana bases of basic lands. A pinned card now adds to the total, and the sweep counts the lands.

**Deck gate run 21 found F-133, and the owner chose the fix inside PR-45b.** The mana pass step of a cheaper card read no job, so it added Mox Amber with the job removal. The step now offers a card of the job it replaces.

**Bracket gate run 6 on the final code reads the judge at 4 of 6.** A second judge lane reads the same decks at 6 of 6. The owner closed F-134 as judge noise, over a plan of a bracket 4 ceiling (D-714, D-715). The counts of power cards do not separate the reads of 4 and 5. The paid runs cost $2.04.

**The review of `gitar-bot` found one defect, and `2ed2b06` fixed it.** The land cap read no pin. A new test fails on the old cap, which drops a pinned Game Changer land.

### 2026-09-14d: #167 merged, and the documents read the state before a context wipe

**The owner merged #167 and asked whether the context can go.** The hand-off still named PR-49 as an open pull request, so this pull request records the merge. It changes no code, and it runs no paid target.

**Cloud Build deployed #167.** `deploy-api` finished at 14:58 UTC, and revision `mtg-api-00046-j29` serves `api:bbd5c5a`. Both jobs run `worker:bbd5c5a`, and `/readyz` answered 200.

**The local meta refresh had not finished at 14:54 UTC.** It started at 14:13 UTC and read mtgo events of February 2026, with 31 pages that did not fetch. A context wipe can stop it, so the next session runs it again.

### 2026-09-14c: #166 merged, the meta run, and PR-49

**The owner merged #166, and Cloud Build deployed it.** `deploy-api` finished at 14:13 UTC, and revision `mtg-api-00045-csf` serves `api:5dbf7d7`. Both jobs run `worker:5dbf7d7`, and `/readyz` answered 200.

**The meta run of 2026-09-14 needs no finding.** The mtggoldfish source read 155 lists with no failure. The job ran 71 minutes, because it read the weekly EDHREC pass and a large mtgo pass. The mtgjson skip on a version mismatch also shows on 2026-09-12 and 2026-09-13.

**The plan read of PR-45b found F-129 and F-130.** The gates passed no top-list rate to the shortlist, so they measured shortlists that the app never builds. Off-theme tutors and many Game Changers never reach a shortlist. The owner chose PR-49 first for every tool, and a rate threshold for PR-45b (D-706 to D-708).

**The session built PR-49.** The parity test fails on the old tools at five calls, and it passes on the new. The dry comparison on model `20260910T012734Z` finds the rate bringing in 0 to 45 cards of a shortlist, with more fast mana and Game Changers.

### 2026-09-14b: #165 merged, and PR-48 fixes F-128

**The owner merged #165, and the session built PR-48** (F-128, D-705). The spell steps of the mana pass read the 99 alone, and the pool holds each commander. `heldIDs` adds the command zone to the set that `manaCandidates` and `cheapestSpell` read.

**Both new tests fail on the old pass.** A unit test reads three steps that add the commander. One drops a spare basic, and two drop the costliest spell. A build test reads "Mana Legend: 2 copies, the limit is 1". Both pass on the new pass.

**Cloud Build deployed #165.** `deploy-api` finished at 13:41 UTC, and revision `mtg-api-00044-gs5` serves `api:05ef44e`. Both jobs run `worker:05ef44e`, and `/readyz` answered 200. A read at 13:39 UTC, while the build ran, answered 503, and its cause is unverified.

### 2026-09-14: PR-45a, the bracket cut, and bracket gate run 3

**The owner merged #164, and the session built PR-45a** (D-702). The content check names each card that the bracket forbids. The build cuts one card of each forbidden combo and each other forbidden card, and basic lands fill the slots. The finding `bracket_cut` names each cut card and the content it held.

**The build refuted one sentence of the plan.** The plan said that the pool keeps the rank order of the shortlist. `NewPool` sorts the pool names by the alphabet, and the shortlist groups its cards by role. So the pool records the shortlist score, and the roadmap carries a dated correction.

**The owner approved bracket gate run 3 over prompts 1 to 9.** It cost $0.7355 over 538 seconds, and 9 of 9 decks hold no content violation. The cut removed Polyraptor from deck 4, beside Marauding Raptor, and Akki Battle Squad from deck 9, beside Kiki-Jiki, Mirror Breaker. The judge agrees on 3 of 9. The first finding text said that the bracket forbids the card itself, so the finding now names the combo.

**Gitar found one defect in #165, and the session fixed it.** A cut of more than five cards left the deck short, because `MaxPad` caps the pad of basic lands at five. The cut now fills each slot it opens, and a gap that the model left stays a block finding. The test of the fix found F-128: the mana pass can add the commander to the 99. The owner chose its own pull request, PR-48, after #165 (D-705).

## The archive

`docs/reference/session-handoff-archive.md` holds every record this file no longer carries. It holds the resume section of 2026-09-08, the records of 2026-08-31 to 2026-09-13, and 42 more sections, word for word. Read it for the detail behind a decision.
