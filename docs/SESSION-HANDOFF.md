# Session hand-off

`CLAUDE.md` is the entry point. It sends you here. Then read `docs/decisions.md` from D-297, `docs/owner-questions.md`, and `docs/open-questions.md`.

This file holds the ten most recent sessions. Every older record sits in `docs/reference/session-handoff-archive.md`, word for word. The first narratives sit in `docs/reference/session-log-2026-08-23-to-26.md` and `docs/reference/session-log-2026-08-27-to-28.md`. The audit of 2026-08-29 is `docs/audit-2026-08-29.md`.

CAUTION: the web tests need Node 22.23.2 (`.nvmrc`). Under Node 20 every test file fails at start with `ERR_REQUIRE_ESM` from jsdom 30. Put `~/.nvm/versions/node/v22.23.2/bin` on the PATH before `make verify`. On this machine `nvm use` reports the change and does not make it, so prepend the path yourself.

## RESUME HERE (2026-09-16)

**Pull request #184 holds the one-pr-one-session policy, and it waits for the owner's merge** (F-154, D-746 to D-748, guardrail 16). A session works on one pull request now. The pull request carries its code, tests, decisions, documents, review answers, and this hand-off. No pull request exists to record an earlier merge or deploy.

**The next step.** Start a new clean session, and load the `one-pr-one-session` skill. Run `make where`, and read the state of #184 with `gh pr view 184`. Then do next step 1. This session is bound to #184 and does no other pull request.

**#182 merged PR-53 on 2026-09-16**, as `372d912` on `main`. Git and GitHub hold that merge, and Cloud Build holds its deploy (D-747). This session read no deploy result, so no line here states one. The notes of PR-53 below still hold.

- **The count.** `profile.FinisherSet` reads the nine parent finisher tags of M-17 with no child tag, and the child tag `blood-artist-ability`. `profile.FinisherIDs` adds the evasive creatures of power 5 or more. The snapshot holds 1,857 such cards.
- **The role.** A Commander shortlist gives the role wincon to the best finishers, up to the target of the bracket, and it pins each one (D-741). An owned mode promotes the owned finishers first (D-742).
- **The floor and the target.** Every Commander bracket carries a finisher floor of 2, and bracket 5 carries 1. The job targets ask for 3 finishers, and 1 at bracket 5.
- **The words.** The finisher note and the finisher finding name the deck plan, and not the bracket (F-151). Only the commander leads the deck (F-153). The prompt version reads 15.
- **The repair.** The shortfall carries the code `finisher_short`, and no repair input holds it (F-152). An empty answer never replaces a deck that holds cards.
- `docs/reference/pr53-wincon-2026-09-16.md` holds every count.

**Four defects came out of the paid runs, and each has a test that fails on the old code** (D-743, D-744).

- **F-150**: `featureWords` held no finisher row, so the finding read " is 0, and bracket 3 wants 2 or more".
- **F-151**: the gap note claimed the bracket wants finishers. The bracket system counts Game Changers, mass land denial, extra turns, banned cards, and combos, and it counts no finisher.
- **F-152**: a pool with too few finishers made the repair answer 0 cards on deck gate prompt 5.
- **F-153**: the summary called Angel of Serenity the leader of the deck. A rejudge held that verdict, so it is no judge noise.

**The paid runs of the session, $9.00 in all.**

| Run | What | Verdict | Cost |
|---|---|---|---|
| Deck gate 24 | the first whole run | FAIL, 3 defects | $2.7876 |
| Deck gate 25 | after the fixes of D-743 | FAIL, 1 defect | $2.5900 |
| Summary judge | a rejudge of the run 25 summaries | the verdict held | $0.1045 |
| Deck gate 26 | prompt 5, the first prompt guard | PASS, 1 rule claim | $0.1292 |
| Deck gate 27 | prompt 5, the reworded guard | PASS, 0 rule claims | $0.1018 |
| Deck gate 28 | the whole run on every fix | **PASS** | $2.7527 |
| Bracket gate 8 | brackets 4 and 5 | judge 5 of 6 | $0.5318 |

**Deck gate run 28 reads PASS**: 25 of 25 decks with no block finding, 0 invented names, and 0 summaries that state a false rule. `eval-check` moves no gate row against baseline run 19.

**Bracket gate run 8 reads the judge at 5 of 6**, 83 percent against a bar of 80. Every one of its six decks meets its finisher floor: 3, 5, 4, 2, 1, and 2, against floors of 2, 2, 2, 1, 1, and 1. The verdict reads FAIL on the band bar, as runs 6 and 7 do. Najeela sits off band on `color_sources` on all three runs. Run 8 adds Kinnan, which misses `avg_mana_value` by 0.13 and `hands_two_to_four_lands` by 0.01. Each run builds new decks, so one deck of six is one sample.

**The review of #182 found one more defect, and the session fixed it.** `promoteUpTo` stepped over a card that already read wincon and pinned, and it counted none of it. So the promotion added the whole target on top of it. At bracket 4 a test reads 4 pinned wincon rows against a target of 3. A pre-pinned wincon finisher counts toward the target now. The fix moves one deck gate pool: prompt 3 falls from 320 cards to 319, and its finisher count holds at 10. Deck gate run 28 measured a shortlist one card wider on that one prompt.

**PR-53 took three wrong readings before the current code.** The first gave the role to every card of the curated count, and the wincon cap of 15 dropped win routes (D-741). The second read the whole scored pool, and an owned mode drops the unowned cards after that step (D-742). The third named the bracket as the source of the finisher floor (D-743).

**#181 merged PR-52** (F-139, F-146 to F-148, D-733 to D-738). A Commander land reads its class at every bracket, and the class reads the mana abilities of the land. At brackets 4 and 5 the mana pass trades a basic land for a better land of the pool. An unowned land that costs more than the whole budget ranks last. `docs/reference/pr52-land-rank-2026-09-15.md` holds every count.

**The free checks of PR-52 hold.** The replay of the Gríma request puts Fabled Passage, Evolving Wilds, and Terramorphic Expanse on the shortlist. The swap takes Fabled Passage and Sunken Hollow, and the reader owns both. 9 of the 25 deck gate pools change, and every one of them is a Commander pool. Quality gate run 24 reads PASS, and it holds the precon bar of a control fit of `main` (D-689).

**Deck gate run 23 reads PASS, and `eval-check` moves no gate row against run 19** (D-738). It cost $2.6509 over 2,156 seconds. The information rows moved both ways, and the plan judge read 34 rows better and 21 worse. The cost of the whole decks rose from $12,893 to $15,539 over the 25 prompts. Prompt 3, an any-card deck with no budget, rose from $4,301 to $7,772. F-149 records a deck of run 23 that holds no nonbasic land.

**Bracket gate run 7 reads the judge at 5 of 6, against 4 of 6 on run 6** (D-738). The bar is 80 percent, so run 7 passes that bar and run 6 failed it. Both runs read FAIL on the band bar, and the same deck sits off band: Najeela on `color_sources`. The basic lands fell on four decks of six, and run 7 cost $0.5235.

**#181 merged PR-52, and Cloud Build deployed it.** `main` is `6ff0244`, and `deploy-api` succeeded on 2026-09-16. Revision `mtg-api-00051-sc6` serves the merge, and `/readyz` answered ok with a card snapshot of 2026-09-15 21:01 UTC. No web build ran, because #181 changed no file under `web/**`.

**The checkout.** `main` is `372d912`, which is pull request #182, or a later merge. The branch `one-pr-one-session` holds #184. Run `make where` before you touch anything. Never commit on `main` (D-583). Answer the review of `gitar-bot` before you ask for a merge, and a pull request of documents alone waits for it too (D-637, D-679).

**#179 merged PR-54** (F-141 to F-145, D-728 to D-732). A theme word finds its row through an alias or a word form, and two new rows cover superfriends and land destruction. When no word of a theme matches a card, the theme row asks for the theme before the build. The replay of the Gríma request reads 14 of the 14 owned mill cards on the shortlist, against 1. The dry deck gate gave the same output before and after the change. `docs/reference/pr54-theme-words-2026-09-14.md` holds every count.

**Question gate run 51 read PASS, and it found F-145** (D-730, D-732). It asked the theme row in 6 conversations that named no theme, such as "the strongest Modern deck". PR-54 adds format names, jank words, and "stuff" to the stop words. Question gate run 52 on the fix read PASS for $0.1951, and the row asked in conversations 108 and 110 alone.

**The owner answered four questions** (D-728 to D-731). The row asks when zero cards match, and not under a floor of 10. Superfriends reads the card type Planeswalker. Question gate run 51 ran after `make verify` passed. A singular creature type keeps the generic rule, because its type row misses type cards such as Field of the Dead (F-144).

**#177 merged M-17, and the owner chose PR-54 first** (D-723 to D-727). The replay of the Gríma request found that the theme "opponent milling cards" matched no card. `themes.json` holds a row for "mill" and none for "milling", so the owned-only shortlist held staple roles alone (F-142). The shortlist dropped the mill cards, and not the model. No reader and no model learns of such a miss (F-143). `docs/reference/m17-finishers-lands-2026-09-14.md` holds every count.

**The owner chose the plan after M-17.** PR-54 added aliases and a word-form rule to the theme match, and a question when no card matches (D-724, D-725). PR-52 and PR-53 follow. PR-53 takes a finisher target of 3 and a floor of 2 at brackets 1 to 4, and 1 and 1 at bracket 5 (D-726). F-140 stays a record until a replay after PR-54 (D-727).

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

**#175 merged the plan, and every document keeps the review.** `docs/reference/owner-review-grima-2026-09-14.md` holds the review word for word. OQ-83 to OQ-86 hold its four suggestions that no decision took. They cover the payoff shape of a commander, evasion that the commander has, caps per effect class, and a power estimate.

**The app is live on `decktome.com`.** A merge to `main` deploys itself on Cloud Build (D-584, D-586). #158 is the newest web deploy, and `deploy-web` released it at 18:52 UTC on 2026-09-13. #179 is the newest API deploy: `deploy-api` finished at 04:13 UTC on 2026-09-15, and revision `mtg-api-00050-w9z` serves `api:4dbbc20`. Both jobs run `worker:4dbbc20`, and `/readyz` answered ok. The deploy triggers read `go/**`, `docker/**`, and `web/**` alone, so a merge of documents alone starts no build.

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

- The merge of #184, after the review of `gitar-bot`.
- The next item of the roadmap (next step 1).
- A deployed session with a theme that matches no card, such as "anime" (next step 2).
- A look at the first commander question after a load of the app, in a session that names no commander (next step 3).
- OQ-67, the Stage B channels.
- OQ-77, the blocking function of Identity Platform.
- OQ-80, a proxy and the pool.
- OQ-83 to OQ-86, four suggestions of the review of the Gríma deck.

**CAUTION: a local `make verify` is not the whole story.** It read green for weeks while shellcheck failed (F-89, D-641). Every pull request runs the workflow now.

## How to resume

1. Load the `one-pr-one-session` skill, and do its start gate. A session works on one pull request (D-746).
2. Run `make where`. It prints the branch, the tree, and the state of the branch's pull request.
3. Run `ps aux | grep autotune` before any write. The loop resets the tree when it rejects an iteration.
4. Make a branch from `main`. Never commit on `main`, and never push to it (D-583). The owner merges (D-585). Run `make hooks` one time in a fresh checkout.
5. Load the skills. Load `ste-writing` before you write any `.md`. Load `design-doc-style` before you edit the roadmap. Load `mtg-corpus` before you reason about a format, a legality, or a card term.
6. Run `make verify`. It runs every check the verify workflow runs, on this machine, for nothing. Put Node 22.23.2 on the PATH first. The pull request runs the same jobs on Actions (D-639).
7. Do "Next steps, in order" below. Ask questions as they come up. Record each owner answer in `docs/decisions.md`, and delete the row from `docs/owner-questions.md`.
8. Open the pull request with the sections of `.github/pull_request_template.md`, and run `make pr-check`. Load the `gitar-review` skill, and follow it after each push (D-637, D-745). Gitar is the only review this repo asks for. A pull request of documents alone waits for the review too (D-679).
9. Update this file inside the pull request, before you call it ready (D-747). Move the oldest session to the archive when the count passes ten.
10. When the pull request is ready, end the session. The next pull request starts in a new clean session.

A second Mac: `docs/setup-second-mac.md` holds what to carry, what to install, and how to prove the machine.

Seventeen things a fresh session gets wrong without this file.

- A test card index with no Oracle text and no tag matches no theme. The theme row of D-725 then asks, and the build never starts. Give each fixture card its real text.
- A singular creature-type word keeps the generic rule, and its plural reads the type row (D-731). So "zombie" and "zombies" read two different lists, and F-144 records why.
- A stored deck records the size of its shortlist and no card of it. So a card that never reached the shortlist and a card that the model dropped look the same. Replay the shortlist for free before a prompt fix (M-17). `.local/m17/zz_scratch_m17_test.go` holds the method, and `list.Theme` names the theme words that matched no card.
- `make bracket-gate` runs its tool in `go/`, through `go -C go`. A relative `-rejudge` path then points inside `go/`, so pass an absolute path.
- Twelve targets and two loop scripts spend money: `make questions-gate`, `make questions-eval`, `make eval-calibrate`, `make deck-gate`, `make bracket-gate`, `make revise-gate`, `make chat-probe`, `make generate-probe`, `make summary-judge`, `make quality-judge`, `make test-smoke`, `make feedback-triage`, `scripts/autotune.sh`, and `scripts/feedback-loop.sh`. Ask the owner before each run. `make autotune`, `make feedback-loop`, `make feedback-loop-dry`, `make feedback-triage-dry`, `eval sweep -dry`, and `go run ./cmd/bracket-gate -sweep` are free.
- A rerun writes to a new file. Every `*_OUT` variable refuses a document that holds a result (D-65).
- A gate run takes about 20 minutes and an eval about 13. A foreground command stops at 10 minutes, so run both in the background.
- `docker compose up` needs provider keys in `.env` now, and fails fast without them (D-267). `make dev` still starts with no keys, and the fake builds the fixture deck of PR-23 for any first message (D-552).
- The eval and the agent share a model. Every ratio it reports is a floor, not a measurement (D-136).
- The judge is noisy: two runs of identical code move up to nine bad questions (D-230). One run proves nothing on its own.
- The sweep's estimate of a step is the cost of its last run file. The deck gate estimate read $2.54 from run 14, and run 16 cost $3.79 with the plan judge and 12 repair turns. A cap set from the estimate stops the sweep before its last step.
- Many local fits and verify runs fill the disk. On 2026-09-11 `make verify` failed at link time with 258 MiB free, when the Go and Docker build caches held about 50 GB. Run `df -h /System/Volumes/Data` before a long run, and ask the owner before you clear a cache.
- A Gitar review can be stale: a paused Gitar keeps the review of an older commit, and the comment names no commit. The `gitar-review` skill proves that a review is current, and it holds every Gitar trap (D-745).
- A background shell command starts in the directory the session left. On 2026-09-12 a `make` target ran in `go/` and found no rule, so head every command with an absolute `cd`.
- A live check of a web change reads the stored session and the deployed chunk, and not the screen alone. The service worker served the old shell for one load after a deploy (F-122). The protobuf-es code holds each field name in base64, so search a chunk for a property name such as `noDecline`.
- The hook `.claude/hooks/session_bind.py` binds a session to the first branch that it creates, pushes, or opens a pull request for. A command on a second branch exits with "Blocked". Start a new clean session. The owner alone removes a binding under `.git/decktome-session-bind/`.
- `Pool.Names` sorts the pool by the alphabet, and the shortlist groups its cards by role. Neither order ranks a card. `Pool.Score` holds the shortlist score (D-702).

## Facts that expire

- Comprehensive Rules: the current file is 2026-08-19 (D-272). The rule citations in the corpus match the 2026-08-07 text. The session of 2026-09-12 read rules 709.4c, 710.2, 712.8a, 712.12, 715.2, and 722.2a in that file for PR-44.
- Ban-list snapshot: 2026-08-24. Next announcement 2026-10-12, in `internal/cards/announcement_dates.json`. A test fails when that file holds no future date. This is by design.
- Commander brackets: the 2025-10-21 revision. Game Changers: 53 cards, list of 2026-02-09. Lutri is banned as a companion only, per the 2026-02-09 announcement (`companion_bans.json` holds the link). The content rules per bracket in `brackets.json` and the Spellbook thresholds were read 2026-09-02, and the Karsten tables are the 2022 articles, read 2026-09-02 (`docs/reference/bracket-profile-2026-09-02.md`).
- Standard: 18 sets, Wilds of Eldraine to The Hobbit. Six sets leave at the first 2027 set. Verified 2026-08-24.
- Card snapshot on disk: `.local/gcs/mtg-local-cards/scryfall/20260904T210157`, the newest of 13. Read 2026-09-13, when it flagged 53 Game Changers. Check every card fact against it.
- The theme table `themes.json` reads `verified_at` 2026-09-14. `make themes-check` read every slug against the snapshot of 2026-09-04 on 2026-09-14. The question gate set holds 109 conversations: 78 counted and 31 probes (D-730).
- LLM model ids and prices: `roles.json` and `prices.json`, verified 2026-08-24. The max output per provider in `llm/client.go`, verified 2026-08-29. The OpenAI rows are unverified by anyone but the owner. The judge role runs on Opus 5 (D-430).
- The deck gate upgrade probe cost $0.32 for 2 prompts, 7 calls, and 266 seconds. A partial run reads its own item bars and never stands as the gate (D-526).
- Run cost, read from the run files on 2026-09-09. The question gate cost $0.190 on run 42, $0.193 on run 43, and $0.194 on run 44, over 18 to 21 minutes. Runs 45 to 48 of 2026-09-11 cost $0.1925, $0.1920, $0.1380, and $0.1923, over 20, 22, 19, and 21 minutes. Run 49 of 2026-09-12 cost $0.1935, over 18 minutes. Run 50 of 2026-09-14 cost $0.1938, over 22 minutes. Run 51 of 2026-09-15 cost $0.1958, over 21 minutes. Run 52 of 2026-09-15 cost $0.1951, over 21 minutes. The deck gate upgrade probe cost $0.32 for 2 prompts. The eval cost $0.092 on run 34 and $0.096 on run 35, over 12 to 14 minutes. The deck gate cost $2.71 on run 18, over 37 minutes, and $2.75 on run 19 of 2026-09-12, over 36 minutes. The revise gate cost $1.07 on run 9. The bracket gate judge lane cost $0.28. Bracket gate run 2 of 2026-09-13 cost $1.3720 over 925 seconds, for 15 builds and the judge. Its second judge lane cost $0.2634 over 128 seconds. The M-15 calibration lane cost $0.3283 over 149 seconds, for 21 decks. Bracket gate run 3 of 2026-09-14 cost $0.7355 over 538 seconds, for 9 builds and the judge. Bracket gate runs 4, 5, and 6 over prompts 10 to 15 cost $0.4979, $0.5247, and $0.5619, over 424, 442, and 564 seconds. The second judge lane of run 6 cost $0.0856 over 35 seconds. Deck gate runs 20, 21, and 22 over prompt 3 cost $0.1321, $0.1174, and $0.1159, over 94, 97, and 106 seconds. Deck gate run 23 of 2026-09-16 cost $2.6509 over 2,156 seconds, for 25 prompts. Bracket gate run 7 of 2026-09-16 cost $0.5235, for 6 builds over prompts 10 to 15 and the judge.
- The backfill of 2026-09-09 read one user with a record to seed: 1 collection, 1 thumbs up, and 2 thumbs down. It counted no deck and no chat, because the reader deleted both (D-635).
- The feedback store holds 3 verdicts on 2026-09-09, and every one predates the snapshot of D-635. `make feedback-list VERDICT=` reads both verdicts now (F-87). A collection group query over it needs an index for its shape, and the store holds "verdict ascending, created_at descending" alone.
- The deployed schedules, read 2026-09-09 and again on 2026-09-11: `mtg-snapshot-schedule` at `0 * * * *` (D-634) and `mtg-meta-schedule` at `0 6 * * *`. Both read ENABLED. The API service holds minScale 0, so it scales to zero. No billing export exists, so no command reads the billed spend.
- The deployed API, read 2026-09-15 at 04:14 UTC: revision `mtg-api-00050-w9z` on image `api:4dbbc20`, from #179. Both jobs run `worker:4dbbc20`, and `/readyz` answered ok with a card snapshot of 2026-09-14 21:01 UTC. `deploy-api` ran from 04:09 to 04:13 UTC. No web build ran, because #179 changed no file under `web/**`.
- The deployed web app, read 2026-09-13 at 18:54 UTC: the release of #158, from `deploy-web` at 18:52 UTC. `index.html` loads `assets/index-DYfo4m8I.js` and no `registerSW.js`, and `sw.js` precaches the `workbox-window` chunk.
- The deployed quality model, read 2026-09-14: `20260914T070904Z`, from the meta job that started at 06:02 UTC and ended at 07:13 UTC. It fits 43,182 lists and 1,517 commanders. Its Commander fit reads `immaterial` 239 and accuracy 0.554, and its Standard fit reads a cross share of 0.667. The job read the weekly EDHREC pass, 2,077 pages and 4 lists. The mtgo source read 3,094 lists with 58 fetch errors. The mtggoldfish source read 155 lists with no failure. The mtgjson source read no list, because its deck list version differs from the stored table, as on 2026-09-12 and 2026-09-13.
- The Karsten land article of 2022-07-29, read 2026-09-11 through `infinite-api.tcgplayer.com/content/article/<id>/`, because the page draws its text in the browser. `docs/reference/m12-rules-diagnostic-2026-09-11.md` holds the formula, the error, and the cheap rules.
- Baselines, in `docs/reference/eval/baselines.json`: questions is run 42, decks is run 19, revise is run 9. The generate prompt reads version 13 since PR-45b, and run 19 read version 12. So the next whole deck gate run reads a new prompt against run 19. Run 52 is the newest whole questions run, and it reads PASS. Runs 45 to 47 read FAIL (D-669, F-112), and run 43 records the regression of F-84. `make eval-check` compares the newest whole run of a suite against its baseline. The quality gate has no baseline row, and run 23 is its newest run.
- Toolchain on this Mac, read 2026-09-11: Go 1.27.1, Node 22.23.2, pnpm 9.2.0, firebase-tools 14.14.0, and Java 17.0.20.1. Playwright is 1.63.0 with its Chromium headless shell (`playwright install chromium`). `go.mod` asks Go 1.27.0 or newer, and vitest is 5.0.0 since #133.
- The local meta store, read 2026-09-14: 17,686 TopDeck good and 6,431 top-cut Commander lists from 2026-06-04 to 2026-09-14. It also holds 1,045 EDHREC average decks and 192 MTGJSON precons. M-17 read the TopDeck lists from 2026-08-01 alone.

## Next steps, in order

1. **Ask the owner for the next item, in a new clean session** (D-746). #182 merged PR-53, and the sequence of the roadmap ends at step 42. OQ-83 to OQ-86 hold the other suggestions of the review of the Gríma deck. The power pass after the build still waits (D-704). F-137 stays a record (D-718), and a replay after PR-54 reads the conditions of F-140 again (D-727).
2. **Read one deployed session whose theme matches no card, such as "anime"** (PR-54). The theme row must ask before the build. The owner builds it, and a session reads it with `scripts/read-session.sh`. A new session holds no copy of the collection export. The collection sits at `users/<uid>/collections/<id>` in `decktome-prod`. `scripts/read-session.sh z1hshyY6Npig1FN2NuV7` prints the user id and the collection id. Its field `entries_gz` holds gzip JSON of the entries. Keep the exported collection out of git.
3. **Read the first commander question on the app after one more load** (D-690). It waits for the owner. The server half holds: session `vY1lCRtl64uwFObznCZ9` stores the flag. The question must show "Suggest one" and no "You decide", and the pick row must still show "You decide". Session `z1hshyY6Npig1FN2NuV7` named its commander, so it showed no such row.
4. **Watch the first self-reload on the next web deploy** (D-692). The live release of #158 passed its check on 2026-09-13 at 18:54 UTC. An installed app that loaded that release must reload by itself when the next release activates. Read it on the next merge that changes `web/**`.
5. **The next collection platform, when the owner names one** (F-91). Five are left: Archidekt, Deckbox, Delver Lens, TCGplayer, and Helvault. Each one takes a real export and never a column list. `docs/reference/pr34-collection-formats-2026-09-10.md` holds the shape.
6. **The live half of the feedback loop has no run yet.** Three things want a measurement: the judge lane of the triage, one live fix cycle, and one review round. All three need the owner's word, and the cycle also needs `AUTOTUNE_FIXER_CMD` and a harvest whose verdicts carry a snapshot.
7. **Deck gate run 19 is the decks baseline** (D-686). Its grades read the local stored model `20260910T012734Z`, and the deployed app reads a newer one. The next whole run compares against run 19, and it costs about $2.75, so ask the owner first.
8. **The owner parked PR-35 and dropped PR-30** (D-655, D-656). M-9, F-97, and F-99 found no case for either one.
9. **PR-36, the reader's verdict as a quality signal** (D-651). It waits for verdicts.
10. **PR-26, the return channels**, waits on OQ-67.
11. **PR-42 is merged as #148** (D-671). Question gate run 49 missed no conversation, so it ran no rerun. Any miss still fails the run, and each miss joins the finding register.

CAUTION: `make revise-gate | tee` hides the exit code. Read the verdict line of the document, never the exit code of a pipe.

CAUTION: `make verify` runs `eval-check`, and `eval-check` reads the newest whole run of each suite. A committed FAIL run turns verify red until a newer whole run passes.

The CI step "fake gcs tests" ran no test until 2026-09-10 (D-658). Its filter matched no test name, so it passed as a no-op. It runs `scripts/gcs-check.sh` now: the script seeds the fake GCS from the trimmed snapshot and fails unless `TestLiveFakeGCS` passes by name. `make gcs-check` runs the same script.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

## The ten most recent sessions

### 2026-09-16c: one pull request, one clean session

**The owner asked for hard lifecycle boundaries, to cut the context cost of each session.** A session binds to one pull request. The pull request carries its documents and its hand-off, and no pull request records an earlier merge (D-746, D-747).

**The first-parent log of `main` held 12 pull requests of documents alone that only record an earlier state** (F-154). Each one read a merge, a deploy, or the state before a context reset. D-42 made a ✅ mean "merged on `main`", so each item merged with 🔧, and a later session wrote the ✅.

**The owner answered three questions** (D-748). A project hook binds each `session_id` to one branch. `CLAUDE.md` drops its merge narrative. A loop pull request meets the same contract.

**The session built the skill, the hook, the check, and the template.** `make pr-check` and the `pr-contract` workflow read the body and the diff. `make lifecycle-check` tests the checker and the hook, and it runs in `make lint` and the verify workflow. The session ran no paid target.

### 2026-09-16b: #181 merged PR-52, and PR-53 built to a context reset

**The owner merged #181, and Cloud Build deployed it.** Revision `mtg-api-00051-sc6` serves the merge, and `/readyz` answered ok.

**The session built PR-53 on the branch `pr53-wincon-target`.** The owner answered three questions (D-739 to D-741). The finisher count reads the curated tags of M-17, the role goes to the target count alone, and every Commander bracket carries a finisher floor.

**The first reading of the role inverted the item.** Every card of the curated count took the role wincon, and the wincon cap of 15 then dropped finishers. 9 of the 25 deck gate pools held fewer finishers than `main`, and prompt 1 fell from 49 to 20. The owner chose the target-count promotion, and no pool holds fewer finishers now.

**`make verify` stopped two paid runs, and no money went out.** The first stop read a document sentence of 27 words, and the second read the constant `roleWincon` as dead code. The session removed it and started the chain again.

**The context reset came before the paid runs landed.** Nothing of PR-53 is committed.

### 2026-09-16: PR-52 built, and two paid gates measured it

**The owner merged #180 and asked what comes next.** Next step 1 named PR-52, so the session built it on the branch `pr52-land-rank`.

**The owner answered three questions** (D-736 to D-738). An unowned land over the whole budget ranks last. The typal lands that leave a typal shortlist are a record. A deck gate run and a bracket gate run measure the item.

**The free replays found five defects of the first land class, and each fix has a test.** Scryfall lists every color for Lotus Vale and Gemstone Mine, so the class reads the mana abilities of a land now. Ash Barrens and Demolition Field read as fetch lands. A typal condition read as easy to meet. The swap dropped the basic land of the scarcer color. It also left one deck 2 basic lands beside two lands that read them.

**Deck gate run 23 reads PASS for $2.6509, and bracket gate run 7 reads the judge at 5 of 6 for $0.5235.** Run 7 passes the judge bar of 80 percent, and run 6 failed it. Both bracket runs read FAIL on the band bar, and Najeela sits off band in each one.

**F-146 to F-149 joined the register.** The mana pass read the pool by the alphabet, and the profile counted a fetch land as a source of every deck color. F-148 records the typal lands, and F-149 records a deck of run 23 with no nonbasic land.

### 2026-09-14n: #179 merged, and the documents read the state before a context reset

**The owner merged #179 and asked for every document to read the current state** before a context reset. The session changed no code, and it ran no paid target.

**Cloud Build deployed #179.** `deploy-api` ran from 04:09 to 04:13 UTC on 2026-09-15, and revision `mtg-api-00050-w9z` serves `api:4dbbc20`. Both jobs run `worker:4dbbc20`, and `/readyz` answered ok. No web build ran, because #179 changed no file under `web/**`.

**The refresh.** The resume section, the next steps, the facts that expire, the roadmap, `CLAUDE.md`, and the note of PR-54 read the merge. The record of 2026-09-14d moved to the archive.

### 2026-09-14m: #178 merged, and PR-54 built

Merged as #179.

**The owner merged #178 and asked what comes next.** Next step 1 named PR-54, so the session built it on the branch `pr54-theme-words`. It ran no paid target before `make verify`.

**The owner answered four questions** (D-728 to D-731). The theme row asks when zero cards match. Superfriends reads the card type Planeswalker. Question gate run 51 runs after `make verify`. A singular creature type keeps the generic rule.

**The free checks.** A count found that the plural rule cost singular type words real type cards, and the owner kept the generic rule for them (D-731, F-144). The replay with the collection of the owner reads 14 of the 14 owned mill cards on the shortlist. The shortlists of the 43 gate prompts and the output of the dry deck gate match `main`.

**Run 51 found F-145.** The run read PASS for $0.1958. It asked the theme row about requests that named no theme, and PR-54 adds stop words. The owner chose run 52 on the fix, over the recommendation to skip it (D-732). Gitar found one performance issue, and the session now keeps the answer of the theme fact.

**Run 52 read PASS on the fix.** The theme row asked in conversations 108 and 110 alone. Gitar approved #179 after the fix, and its one finding reads resolved.

### 2026-09-14l: #177 merged, and the documents read the state before a context reset

**The owner merged #177 and asked for every document to read the current state** before a context reset. The session changed no code, and it ran no paid target.

**No build ran.** #177 changed documents alone. A read of Cloud Build at 23:43 UTC found `deploy-api` on `54b07dd` of #173 as the newest build.

**The refresh.** The resume section, the facts that expire, the roadmap, `CLAUDE.md`, and the review of the Gríma deck read the merge. A new item of the list of things a fresh session gets wrong says to replay the shortlist before a prompt fix. The record of 2026-09-14b moved to the archive.

### 2026-09-14k: #176 merged, M-17 measured, and the owner chose PR-54 first

**The owner merged #176 and asked what comes next.** Next step 1 named M-17, so the session measured it on the branch `m17-finishers-lands`. It called no provider, and it changed no code. A scratch test in `go/cmd/deck-gate` made every count, and its copy sits in `.local/m17/`.

**The replay found the cause of F-141.** The theme "opponent milling cards" matched no card, because `themes.json` holds no row for "milling" (F-142). So the owned-only shortlist held staple roles alone, and the shortlist dropped the mill cards, and not the model. Nothing tells the reader of such a miss (F-143).

**The owner chose five answers** (D-723 to D-727). PR-54 fixes the theme words first, with aliases, a word-form rule, and a question when no card matches. PR-53 takes a finisher target of 3 and a floor of 2 at brackets 1 to 4, and 1 and 1 at bracket 5. F-140 stays a record until a replay after PR-54.

**The session made one mistake in a question, and it corrected it.** The first floors question said that a floor of 2 flags the Gríma deck. The deck is bracket 5, where the chosen floor is 1. The owner read the correction and kept the choice.

### 2026-09-14j: #174 and #175 merged, the review of the Gríma deck, and the plan of M-17

**The owner merged #174.** It changed documents alone, so no build ran.

**The owner gave a written review of the Gríma deck, and the session checked each claim.** The session read the stored deck, the card snapshot, the stored collection of the owner, and two research passes over the code. The review holds on the win condition, the mana base, and the unmet conditions. It fails on the pump shelf, and the missing staples come from the owned pool. The check also found F-141: the deck holds none of the 14 owned mill cards.

**The owner chose the plan** (D-719 to D-722). M-17 measures first, then PR-52 improves the mana, and then PR-53 adds a win-condition target. The session changed no code, and it ran no paid target after #173.

**The owner merged #175 and asked every document to keep what the session learned.** `docs/reference/owner-review-grima-2026-09-14.md` holds the review word for word. The review document maps each of the seven suggestions, and OQ-83 to OQ-86 hold the four that no decision took. Next step 1 names where the stored collection sits for the replay of M-17.

### 2026-09-14i: #173 merged, and the deployed app asks no role question

**The owner merged #173, and Cloud Build deployed it.** `deploy-api` ran from 20:45 to 20:50 UTC, and revision `mtg-api-00049-4lj` serves `api:54b07dd`. Both jobs run `worker:54b07dd`, and `/readyz` answered ok.

**Deployed session `l9x5bFTcFgpZ4DOcQpxU` proves PR-51** (F-136, D-717). The owner wrote "Grima as commander", and the lookup found two cards. Turn 1 asked the power and the colors, and no role question. Turn 2 offered both Gríma cards.

### 2026-09-14h: #171 and #172 merged, the live Grima check, and PR-51

**The owner typed "Grima" as a commander, and no card matched.** The card is Gríma, Saruman's Footman. The owner chose a fold in every name match, with the exact key first (D-716). The session built PR-50 on the branch `fix-name-accent-fold`, and #172 merged it. Gitar found one performance defect, and `6394d16` added a fast path for ASCII names.

**The owner merged #171 and #172, and Cloud Build deployed #172.** `deploy-api` finished at 19:19 UTC, and revision `mtg-api-00048-5zm` serves `api:284b324`. Both jobs run `worker:284b324`, and `/readyz` answered ok. #171 changed documents alone.

**Deployed session `z1hshyY6Npig1FN2NuV7` read "Grima" right, and it proves PR-45b.** The which-card row offered both Gríma cards. The deck at bracket 5 misses all three power floors, and its summary names the cards that close each gap. The read found F-136 and F-137. The owner chose PR-51 with question gate run 50 (D-717), and a record of F-137 (D-718).

## The archive

`docs/reference/session-handoff-archive.md` holds every record this file no longer carries. It holds the resume section of 2026-09-08, the records of 2026-08-31 to 2026-09-14g, and 42 more sections, word for word. Read it for the detail behind a decision.
