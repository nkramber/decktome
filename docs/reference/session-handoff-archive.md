# Session hand-off archive

This file holds every session record that `docs/SESSION-HANDOFF.md` no
longer carries. The hand-off keeps the ten most recent sessions. Each
older record moves here, word for word.

Read this file for the detail behind a decision. Read `docs/decisions.md`
for the decision itself. Read `docs/SESSION-HANDOFF.md` to resume the work.

The records run newest first. The oldest narratives sit in
`docs/reference/session-log-2026-08-23-to-26.md` and
`docs/reference/session-log-2026-08-27-to-28.md`.

## 2026-09-10g: F-111, the offer waits for the theme and the colors

Merged as #139.

**The owner reported a turn that makes no sense.** Session `zJPjCkVR44jcSqxrlkmb` asked "Build a bracket-5 commander deck". One turn asked the theme and the colors and offered Vivi Ornitier, Aang, and Hapatra.

**The cause is the wait of the pick row.** D-630 made the offer wait for the power alone. The first message filled the power, and the classifier read a request for a suggestion. So the offer took the third place of the turn beside two questions it ranks on.

**The first shape made the offer wait for both answers, and question gate run 45 read FAIL** ($0.19). Conversation 14 skipped the color question and asked for a suggestion, and its first offer came on turn 4, too late to pick. Six conversations that got an offer in run 44 got none. No scripted conversation had the shape of the owner's session, so neither run ever offered beside a preference question.

**The owner chose the rule of the turn** (D-669). A new catalog field, `not_beside`, keeps the offer out of a turn that asks the theme or the colors. The offer waits for no answer. The owner chose to leave the plain commander row as it is.

**Question gate run 46 read FAIL on one classifier read** ($0.19). Conversation 14 passes, and no run offered a commander beside a theme or colors question. The miss is conversation 100, a Modern deck with no commander row, where the classifier kept the theme "team event". Runs 45 and 46 asked the same questions on every turn of it. The owner chose run 47.

**Run 47 read FAIL on another single conversation** ($0.14). Conversation 103 built after turn 2, before the reader named the bracket and the budget. Runs 44 to 46 asked the same questions on its first two turns and reached turn 3. Conversations 14 and 100 pass, and 16 conversations got an offer (F-112).

**The owner asked how to fix the check properly.** Each miss traced to a real gap the classifier trips on some runs. "You pick the commander" let the classifier decline the bracket and the budget (F-113). "Team event" filled the theme, so the theme question never went out (F-114). Both guards joined #139 with turn tests that fail without them (D-670). The gate keeps a bar of zero misses, and PR-42 reruns a missed conversation (D-671).

**The effort measurement found no miss at either effort** (D-672, $0.04). Each conversation ran 8 times at `none` and 8 times at `low`. `low` cost 1.5 times as much per run, so the classify role keeps `none`.

**Question gate run 48 reads PASS on the code with both guards** ($0.19). Every expectation holds, conversations 14, 100, and 103 pass, and no run offered a commander beside a preference question. `make eval-check` reads PASS again.

**The first turn test passed without the fix.** An any-card pool asked the budget, and three rows filled the turn. The test copies the owned-only pool of the session now, and it fails with the old catalog. A second test walks the skip of conversation 14.

## 2026-09-10f: the plan of M-11, the detector

Merged as #138.

**The owner asked for a plan of the detector fix.** The fit prints no weights, so a throwaway patch printed them. The cause of D-659 is the population of the fit. The fit breaks the precons and the average decks alone, and the top lists stand among the negatives. So `card_rate` weighs -1.08 in Commander, and no break moves it by more than 0.18 standard deviations (F-106).

**The colors break reads by its signature.** `source_spread` weighs +0.81 and `tapped_share` -0.79, and `color_sources` weighs -0.04 (F-107). So a splash and a mana base of basics read broken, and the real shortfall of deck 22 reads sound.

**Three owner answers.** M-11 measures fitted variants first (D-666). The change reads Commander alone (D-667). The new-set features change in their own item, PR-41 (D-668).

**Two more findings from the review.** The ladder reads the power norms of the top lists as faults (F-109). A bracket 4 request built a baseline deck, and OQ-82 asks what a bracket promises (F-110).

## 2026-09-10e: gate run 19, PR-39 on the quality model

Merged as #137.

**The owner asked what comes next.** The hand-off named the bracket walk, the detector, and the meta log. A read of the code found one more gap (F-105). The quality feature `color_sources` reads the source count PR-39 changed, and no gate read the model. The meta job of 2026-09-11 refits the stored model on the new count.

**Gate run 19 cost nothing and reads PASS** (D-663). A control fit on `4d8635c`, the parent of PR-39, matches run 18 in every number. So the fit repeats, and the count alone moves run 19. Every bar holds. One built deck moves: deck 7, which the judge reads as bad, grades bad now. So graded bad rises from 15 to 16 of 25, and the judge agreement from 8 to 9.

**Two owner answers.** An item that changes a feature of the quality model reports the three numbers now, as guardrail 14 (D-664). PR-38 waits on the detector fix (D-665).

**The deployed app.** The deploys of #134 and #135 succeeded. No reader session reached the API after the deploy of #134. The ERROR lines of the API log are `/readyz` answers of 503 in a cold start, before the card index loads. They go back to 2026-09-08 at least. One `GetCards` call read the same 503 at 19:02 UTC.

## 2026-09-10d: F-104, the offer reads the bracket of its own turn

Merged as #134. PR-39 followed and merged as #135. Dependabot's #131 to #133 merged, and #133 took a fix for vitest 5 and protoc-gen-es 2.14.1 first.

**The owner asked whether the commander options were the expected quality.** Session `vDzEKDPnRZntlyhz8Lqg` asked for the best possible deck at bracket 5 from an owned-only pool. It read Lotho, Corrupt Shirriff, Ghalta, Primal Hunger, and Peregrin Took. They were not. The bracket 5 signal of OQ-48 exists since PR-14B, and the turn never handed it the bracket.

**The cause is the turn order.** The turn builds the hints from the slots as restored, and the picked power option lands inside the turn. `readFacts` refreshed the colors, the pool, the sets, and the pair flag, and never the bracket. The pick row fires on the turn the power answer arrives (D-631), so the reorder ran on almost no deployed session.

**The proof cost nothing.** A local run of the pool over both fixture exports reads the same three names with no bracket. With bracket 5 it reads Vivi Ornitier, Aang, at the Crossroads, and Hapatra. The deployed model of that morning holds a signal for 964 commanders. The log of the session shows the power pick and the offer 84 milliseconds apart on turn 2.

**The fix is one interface.** `BracketAware` joins `SlotAware`, `SetAware`, and `PairAware`, and `readFacts` hands the bracket over. The turn test fails without the call and passes with it.

**The session before this one ended without a hand-off.** Its work sits on `builder-basics-by-pips`: the source count of F-103, the balance phase of F-102, and D-659 to D-661. The measurement of D-661 ran in this session, for nothing. With the fixed count the pass moves four decks of run 18, and the worst color rises on all four. Deck 22 gains two Forests, where the old count cut two. PR-39 is open.

**The `decktome` gcloud configuration pointed at the Wallabee account and project.** The session read the deployed project through an environment override and changed no configuration.

## 2026-09-10c: PR-37, the synergy check

Merged as #127.

**The owner asked what the merit is, when no reader-facing number moved.** The answer: the bar can now judge a change, and the gate prints the three numbers for free. The model grades the built decks as before, and the judge agreement of 8 of 25 is still the gap a reader feels.

**A free test then cut the Commander casual lists to a half and a quarter** (F-97). The built decks graded bad read 15, 15, and 16, and the judge agreement 8, 8, and 7. The owner chose M-9 before PR-35 (D-654).

**M-9 refuted the premise of D-650** (F-98). With the casual features of PR-29, a larger casual corpus read worse on both reader-facing numbers. The owner parked PR-35 and chose PR-30 next (D-655).

**PR-30 failed in both forms** (F-99). With no self-match no built deck moved, and with it graded bad rose from 15 to 20. The owner dropped PR-30 and chose a structural change to the model, which M-10 measures first (D-656).

**M-10 measured a softmax scorer, and the detector held the built decks** (F-100). The owner asked for the ten disputed decks as a review sheet, to grade them by hand (D-657).

**The check measures the move after the break** (D-652). In each fold the fit reads the synergy feature of every synergy copy, and it drops a copy whose fall sits under the floor. A request that passes no check makes no copy.

**The owner chose the floor on six measured runs** (D-653). At 0.10 the model won the dropped pairs 47 percent of the time before the check, a coin flip. Each higher step drops pairs it won 72 to 90 percent of the time.

**The premise was half right.** The check drops 157 of 389 Commander pairs from every band, and not the 54 weak pairs alone. **In Modern the break does not move the synergy feature in most pairs** (F-96), so the check reads Commander alone.

**Gate run 18 reads PASS, and no reader-facing number moved.** The bar reads a new population, and the model is the same model. The gate now reports the three numbers of D-648 for free, and judge lane run 5 joined `main` for them.

**The CI step "fake gcs tests" runs its test now** (F-101, D-658). Its filter matched no test, so it passed as a no-op. `scripts/gcs-check.sh` seeds the fake GCS from the trimmed snapshot and fails unless `TestLiveFakeGCS` passes by name.

## 2026-09-10b: PR-29 closed, and M-8 read the bar

Merged as #126.

**The casual corpus worked and the trade did not.** A control fitted the same day over the same data reads the numbers apart. The Commander synergy axis read 0.70 and 0.81. The built decks graded bad read 15 of 25 and 17, and the judge agreement 8 of 25 and 6. The judge comparison holds no judge noise: one run answered every deck once, and both models read those same answers.

**Two further features failed and went back.** `casual_unseen_share` moved the target axis by nothing and cost three other bars. `casual_pair_share` won one pair of 389, because the fit split it against `casual_synergy` with opposite signs.

**M-8 found the target was never reachable.** A weak precon's cards do not pair in the corpus, so the synergy break removes half of nothing. 54 of 389 Commander pairs carry no signal and read 0.43, so a perfect model reads 0.93.

**The repository already held the rule.** D-485 says a break must be material, and it checks the copies axis and the colors axis. Synergy is the fallback, and it checks nothing (F-95).

**The owner closed PR-29 and answered the three questions of PR-37.** The evidence document stays, because two findings rest on it.

## 2026-09-10: PR-34, the collection formats

Merged as #123.

**The app reads the format out of the file** (D-647). `collections.Detect` reads the header row, and the reader never names the app their file came from. The app refuses a tie between two formats and names both. It refuses a file no format claims with the list it does read.

**F-92 was live, and nobody met it.** The web sent `MANABOX_CSV` on every upload, hardcoded. The app has read an Arena list since D-15, and every one of them failed every row of itself.

**Moxfield reads, on a real export.** The owner exported their own collection, 3193 rows. The documentation of that format disagreed with itself on `Edition` and on `Condition`, and the file settled both.

**The language column was the trap.** `Resolve` refuses a row that is not `en` (D-23), and Moxfield writes `English`. A pass-through reports all 3192 English cards of the export as non-English. The reader then reads that their whole collection failed.

**The measurement.** 3193 rows parsed, none unread, 125 sets. The snapshot of 2026-09-04 resolves it into 3035 entries and 5884 cards. One row stays unresolved, a Japanese printing, and D-23 refuses that one by design.

**One CSV walker serves every format now.** A platform is one signature entry and one row builder.

**OQ-80 is open.** A Moxfield export carries a `Proxy` column, and no other format this app reads carries one. The owner parked the question rather than decide it on no evidence: every row of the one export on record reads False.

## 2026-09-09c: PR-28c, the fix cycle

Merged as #122. The feedback loop is complete: PR-27 harvests, PR-28a writes the files, PR-28b writes the cases, and PR-28c fixes them.

**The cycle.** `scripts/feedback-loop.sh` reads a harvest and writes one case per thumbs down. It proves each case fails, hands the failures to a fixer agent, and proves each case passes (D-645).

**The case and its fix ride in one pull request.** A case is a failing test by design. A case merged on its own turns the gate red on `main` until a fix lands, and the owner chose one pull request over two.

**The cycle confirms before it fixes.** A case that already passes never measured the reader's fault. It stays as a case a change must not flip, and the fixer never sees it. That is D-234 applied to a case.

**The accept rule is sharper than the tuning loop's.** No noise margin and no ratio: every case goes from fail to pass, `make eval-check` shows no flip, and the tree builds. `cmd/case-check` reads the same bars the gate's own verdict reads.

**The cycle answers the review.** `scripts/feedback-review.sh` hands every open finding of `gitar-bot` to the fixer, pushes, and replies on each thread, over three rounds (D-637).

**The frozen list drifted, and a test caught it.** `TestTheFixerPromptAndTheCycleAgreeOnTheFrozenList` found nine paths the cycle enforces that the prompt never named.

**PR-34 joined the roadmap** (D-646, F-91): the app reads two collection formats, and a reader on any other platform can not upload at all.

## 2026-09-09b: PR-28b, the triage and the cases

Merged as #121.

**The triage.** The reason keys of the dialog name the class with no model call (D-643). PR-27 already asks the reader to name the fault, and the keys map onto the classes one to one. The judge reads three cases alone. They are free text with no reason, reasons that cross two classes, and a reader who argues with a rule the owner set. Eight of the ten fixture items cost nothing.

**The classes grew from ten to 22** (F-90). The design note of 2026-09-06 held ten, and the store offers 22 reason keys. Twelve of them named no class, and a verdict of any of the twelve fell through in silence. The chat kind of D-594 alone added five, a day after the note.

**The expectations.** A "must not ask" expectation names a catalog row that must not fire, and the question gate reads it as one more miss (D-644). The row comes from the question id, which reads `q<n>-<row>`. The deck gate reads two case assertions: a card the build must not pick again, and a deck the reader owns whole.

**The store.** A verdict about a deck keeps the session that built it (D-643), because the theme, the pool rule, and the budget live on the session. Two snapshots that pass the room of one Firestore document drop the session and keep the deck.

**The repository.** D-640 is gone (D-642). The owner permits a reader's exact words in a committed file, so the harvest commits and a case joins the gate file that owns it. The pull request diff is the accept step.

## 2026-09-09: the review process, the user record, and a public repository

Seven pull requests merged, #113 to #119.

**The question workflow.** Gate run 43 read FAIL and found a regression of D-630. `Skip` fills a key and `CloseStalled` does not. So the commander offer waited for a power the net closed. A `Requires` slot reads the skipped keys now (D-631). Run 44 reads PASS.

**The decks.** The upgrade probe measured D-628 for $0.32 and refuted the estimate of D-629: the bands answered 1 of the 8 upgrade findings, not 7. The Turtle Power precon leaves 18 free slots and the role minimums need 17, so an in-band upgrade adds no card of the theme. The owner parked F-85 for want of a second precon (D-632).

**The money.** The PR-22 gate closed on a measured cost at idle of $9.54 a month gross (D-633). The snapshot cron held 80 percent of it, and every tick paid a container start to learn there was nothing to do (F-86). The tick runs hourly now, which leaves $3.81 (D-634).

**PR-28a.** PR-28 split into three (D-636). A verdict keeps the object it names, because a reader deletes the deck or the chat they complained about (D-635). `make feedback-harvest` writes a dated document and a JSONL file, and the watermark comes from those files.

**The user record.** `users/<uid>` holds the email, the dates, and six counters (D-638). The owner ran the backfill.

**The process and the repository.** `gitar-bot` reviews every pull request, and the session answers each finding before it asks for a merge (D-637). The repository went public: every pull request runs the whole workflow (D-639), and no harvest file enters the repository (D-640). The first pull request on Actions found two broken checks. No pull request ran either one before that day (D-641, F-89).

## 2026-09-08: PR-32, PR-33, PR-25, and four walks

Sixteen pull requests merged, #97 to #112. PR-32 closed F-75 and F-76, the commander a reader names and the ownership of a commander (D-606 to D-608). PR-33 landed a build in band with one model call, where the same request had read 3 minutes 59 seconds and an error (F-77, F-78, D-613). Deck gate run 18 measured it and became the decks baseline: repair turns fell from 12 of 25 to 3 of 25 (D-617). PR-25 made the app installable on a phone, and the owner walked it from the Home Screen (D-621 to D-627). The walks found F-79 to F-83, and every one is fixed.

## 2026-09-07: the deploy, the invite gate, and the dead conversation

The app went live on `decktome.com` (D-574). A merge to `main` deploys itself on Cloud Build, through two triggers that read the diff (D-584, D-586, D-588). An uninvited reader had read the whole app, and an invite gate holds every page now (D-589, D-590, D-592). The classifier wrote over the card pool the reader chose, and the reader's choice wins now (D-591). A picked option sets its slot as data, and no turn is silent (D-597 to D-599). M-7 landed the honest bars and the axis diagnostic (D-570 to D-573).

## 2026-09-06: PR-22, PR-23, and PR-27

PR-22 landed the invite list, the spend cap, and the hosting block (D-550). PR-23 landed the Playwright smoke flow, where the fake serves a build (D-551 to D-553). PR-27 landed the feedback harvest: the thumbs, the dialog, the store, and the service (D-561). The mobile and engagement proposal set Stage A as PR-25 (D-546 to D-548).

## 2026-09-05: the terse conversations and the trimmed snapshot

The 47 terse conversations joined the bar, with four classifier rules and the GCP deploy guide (D-522, D-534 to D-539). The precon check ignores basic lands, and the M-5 sheet is complete (F-35, D-523, D-529 to D-532). A trimmed card snapshot of at most 10 MB serves the free dry lane of the deck gate (D-521, D-542).

## 2026-09-04: PR-15, the eval harness

PR-15 landed the eval harness (#65). The paid gate ran with it (D-514, D-515). The corpus took the MTGJSON skip and the precon exclusion prompt. Three set fixes closed F-40 to F-42: the Marvel set question, the group set request, and the partial run.

## 2026-09-03: four product pull requests

PR-14B landed the deck quality model (#58). PR-24 landed the precon exclusion (#59). PR-14C landed MTGTop8 and the casual 60-card decks (#60). PR-20 landed the deck view and the card detail (#61), and PR-21 landed the share link and the print view (#62). Two deck gate fixes followed: the precon share reads the pool, and a failed repair keeps the deck.

## 2026-09-02: PR-19 and PR-14A

PR-19 landed the chat and build experience, the counted land swap, and the split land bucket (#55). Its follow-up made a deck count its commander (#56). PR-14A landed the bracket profile (#57). Bracket gate run 1 and deck gate run 12 ran against them.

## 2026-09-01: PR-17B and PR-18

PR-17B landed the set filter and the fixes three gate runs found (#50). PR-18 landed collection management (#53), and its review fixes made the chat refuse a turn with no card index (#54). The question-quality pass ran the same day (D-387 to D-389).

## 2026-08-31: PR-17, the deck library

PR-17 landed the deck library, the one deck screen, and the reference design (#49). The paid runs of the day and the dead conversation of D-351 to D-354 sit below.

## The resume section of 2026-09-08

**The checkout.** `main` is `77a0ae1`, which is PR #99, merged on 2026-09-08. Branch `pr-32-named-commander` is merged and gone. Run `make where` before you touch anything.

**What #99 holds.** The second half of PR-32, and it closes F-75 and F-76.

1. F-75, the commander name the card index does not hold (D-606, D-607). Two catalog rows ask which card the reader means. `commander_unresolved` names the best three cards that hold the name and offers "None of these". `commander_unknown` says a name no card holds matched nothing. The name reaches no build until the reader answers. `CandidateHints.ResolveCommander` reads whole words, so "Aragorn" and "King of Gondor" both find "Aragorn, King of Gondor". A picked option sets the commander with no model in the path (D-597).
2. The second half of F-76 (D-608). `Deck.commanders` carries one `DeckCard` per commander, with the ownership, the count, and the price. The build fills it, and the deck screen marks it. The two buy lists read it, in the web app and in the export package. An unowned commander counts in `buy_cost_usd`, which `BuyCostWith` did already. **A deck built before #99 carries no entry**, so its commander shows no mark and reaches no buy list.
3. The smoke gap of #92. The fixture deck comes from the owned-first shortlist of the fixture export now, so the flow builds from the collection. Karlov is the one card the export does not hold. So the lane reads the buy mark and the buy list of F-76.

**The deploy of #99 passed.** It changed `go/`, `web/`, and `proto/`, so both triggers ran, and both read SUCCESS at `77a0ae1`. Cloud Run revision `mtg-api-00016-mkm` is the API of that build, ready at 14:21 UTC, and `decktome.com` answers 200. Read a build with `gcloud builds list --project decktome-prod --region us-central1 --limit 4`.

**The commander offer waits for the power** (2026-09-08, D-630). Session `XnY2uVuQkAL0VN9AeNCj` asked the power bracket and offered three commanders in one turn. The offer ranks on the bracket, so it ranked for no power at all. `commander_pick` and `commander_unresolved` require the power now. **This costs every reader one turn**, and the owner chose it after the session named the cost. Eight test fixtures held the old order, and each names the power now.

**PR-25 holds every free part, and the phone walk is what is left** (2026-09-08, D-621 to D-623). Branch `pr-25-installable`. The app carries a manifest, a service worker that updates itself, and the icons of the tome mark. It holds an install hint and a paste box for the CSV. A touch target reaches 44 pixels on a coarse pointer. **No response of the API is ever cached**: it answers a POST per user, and the navigation fallback denies its routes. `src/lib/pwa.test.ts` is the free gate of the install, and it reads the icon files themselves.

**An upgrade reads the bands now** (2026-09-08, D-628, D-629). Deck gate run 18 left 8 of its 10 off-band findings on the two precon upgrades. The read found the cause. `input` nested the deck shape block inside the job target block, and D-249 takes the targets off an upgrade. So an upgrade read no band at all. The mana pass also moves an upgrade's lands now, and it never drops a precon card and never moves a spell. **The pass answers 1 of the 8**, the `colorless_land` count, and the other seven rest on the bands reaching the model. No free lane measures that, so **the next whole deck gate run is the measure** (D-629).

**The app carries no health bar** (2026-09-08, D-626). It read "API: ok, version dev" and the age of the card snapshot, and no decision asked for it. Every deck prints the card data it rests on, on the deck. The document holds a fixed height and refuses the bounce, so a drag at the foot of a phone reveals no empty ground.

**The phone read of 2026-09-08 found no overflow, and it forbade the pinch** (D-624). `make smoke phone.spec.ts` reads every screen at 390 by 844, the deck screen and the card sheet among them, and no document passes its viewport. So the sideways scroll the owner read was a pan after a pinch. The viewport carries `user-scalable=no` now, which costs the axe meta-viewport rule of WCAG 1.4.4. **The owner chose that after the session named the cost**, and the PR-25 gate carries the exception.

**What the owner does for the rest of the phone walk.** Open `decktome.com` on the phone. Add it to the Home Screen from the share sheet. Open it from the Home Screen, turn the network off, and read the shell. Then walk the path: upload, build, revise, and export. Two things need a real phone, the install and the chat input above the keyboard.

**Three of the four PR-22 gate items hold** (2026-09-08, D-620). The build walk passed on `decktome.com`, session `72b2IAMZqaC6myfuKAg4`: upload, build, revise, and export. The refusal of an email off the invite list passed the same day, and guardrail 9 passed on 2026-09-06. **The measured monthly cost at idle is the fourth**, and it waits for a quiet week. The week of the deploy held the walks and the deploys of PR-32 and PR-33. `docs/reference/pr22-gate-2026-09-06.md` holds the table.

**The owner decides what the open cost figure blocks.** PR-25 and PR-28 wait on the PR-22 gate (D-547, D-555, D-557), and hard rule 1 holds an item until the gate of the item before it holds. The three items that test behavior all pass, and the fourth is a number to write down.

**Deck gate run 18 is the decks baseline, and it measures PR-33** (2026-09-08, D-617). It ran at the shipped configuration on the pinned snapshot of run 16, so the code of PR-33 is the one variable the pin holds. It reads PASS. **Repair turns fell from 12 of 25 to 3 of 25**, and a block bought each of the three. The wall clock fell 27 percent and the cost 28 percent, and the calls fell from 93 to 78. On the decks the mana pass runs on, the off-band findings fell from 4 to 2. **The two precon upgrades hold 8 of the 10 that are left**, and the pass skips an upgrade by design (D-249). So the next question goes to the owner. Does the mana pass move the mana base of an upgrade?

`make eval-check` reads run 18 as the baseline now, and it lists run 17 as an experiment it never compares.

**PR-33 merged as #102, and the owner walked the deployed app** (2026-09-08). The same request that read 3 minutes 59 seconds and an error read **37 seconds and one model call**. The log line "the deck is legal and its bands are what the pool allows, so no repair turn runs" is the new path, and the mana pass made 10 steps and left `off_band` at 0.01. **The walk found three faults, and branch `f79-manapass-ownership` fixes them** (F-79, F-80, F-81, D-614 to D-616).

- F-79: the mana pass named seven cards the reader owns as cards to buy. `addOne` wrote no owned count and no price.
- F-80: a revision added no card the deck does not hold. "Include The Arkenstone from my collection" changed nothing.
- F-81: a revision that changed nothing stored a new version of the deck.
- D-614: a set-limited deck shows a printing of the sets the reader named.

**The old PR-33 entry follows.** PR-33 held four parts, and it waits for a merge (2026-09-08, F-77, F-78, D-609, D-613). Branch `pr-33-on-band`, from `main` at `ea5e7f7`. Two deployed builds ran four minutes and gave the reader nothing, and `docs/reference/on-band-plan-2026-09-08.md` holds the read and the plan.

The four parts:

1. Every band the check reads reaches the model. `promptKeys` named five of the fifteen features `bands.json` holds. The four silent ones are lines the model can act on now. `TestEveryBandReachesTheModel` fails when a band reaches no prompt.
2. The job block carries the band beside the target: "land: 31 to 36, and 33 is the middle".
3. `Builder.fixMana` moves the mana base with no model call, one step at a time. It keeps a step only when the deck moves closer to its bands. It holds five levers. An untapped land replaces a tapped one. A basic replaces a basic of another color. A land replaces a spell, or the reverse. A cheaper card of the same job replaces the costliest one. Ramp replaces the costliest card. **`make manapass-check` is the free lane**, 9.6 seconds over the 25 decks of run 16.
4. `MaxRepairs` is 1, and a deck whose findings are profile findings alone reaches the reader with no repair call. A failed repair call leaves the last legal deck standing. A repair turn starts only when the build has time for it. This supersedes the second pass of D-461.

**Read the free lane before you trust the pass.** Its first shape moved the lands alone and closed 2 of 10 off-band features. The curve levers went in after that read, and it closes 3 of the 4 features of the decks it runs on. Six of the ten sit on precon upgrades, which the pass skips by design (D-249).

**Deck gate run 17 is an experiment now** (D-612). A run header names each role field it changed from `roles.json`. `make eval-check` never takes such a run for the newest run of its suite, and it lists it instead. Without that rule run 17 failed every `make verify`.

**Deck gate run 17 is the effort measurement, and it is done** (2026-09-08, D-611, D-612). It ran `LLM_GENERATE_EFFORT=low` against the run 16 baseline, on the pinned card snapshot of run 16, for $3.33 and 47 minutes. Low effort saves 7.8 seconds of the 120.7 a prompt takes. It reads FAIL, because prompt 11 answered a Commander deck of two copies of Skullport Merchant. **The generate role stays at medium**, and run 16 stays the decks baseline. Run 17 is an experiment and never a baseline. **Read the caution in the plan document before you compare the two runs.** The plan_rubric prompt and the quality model both moved between them. So the `grade` and `plan_score` rows do not compare. `make eval-check` skips run 17 as an experiment (D-612), so it never stands for the decks suite.

**What comes next after that.** The build walk of the PR-22 gate comes first, and it blocks PR-25 and PR-28. PR-29, the casual corpus, stands on the weak-axes plan (D-567). The next EDHREC read falls on 2026-09-14.

**A walk to make after the deploy.** Write "Aragorn as commander" to the deployed app, and read the row that asks which card. Then build a deck of an owned-first pool, and read the commander mark on the deck screen.

**What waits on the owner.** OQ-67, the Stage B channels. OQ-77, the blocking function of Identity Platform. It costs nothing to 50,000 monthly active users, so the cost is not the question. The seven-second answer limit on every sign-up is the question. The build walk of the PR-22 gate.

**What changed on the deployed project on 2026-09-08, outside the repo.**

- `gh-deployer` gained `roles/firebaserules.admin` and `roles/datastore.indexAdmin` (D-605). Both widen it, and the second one went in without a second question.
- The `deploy-web` trigger reads `firestore.indexes.json` now (D-603).
- The build deployed the indexes of `firestore.indexes.json`, and `make feedback-list` works.
- `an invited reader` is on the invite list.
- `the owner's second account` still exists in Firebase Auth, from the walk of 2026-09-07. No command here removes an account.

**Three habits this session paid for.** Read the local date and never the UTC date (D-595). Read a gate document before a row says a run is due. Put Node 22.23.2 on the PATH before `make verify`, or every web test file fails at start.

## Where things stand (2026-09-07)

- **The app is live on `decktome.com`** (2026-09-07, D-574). The deploy of PR-22 ran on the owner's account by `docs/setup-gcp.md`, on project `decktome-prod` in `us-central1`. The API is a Cloud Run service, and both jobs and both schedules run. Firestore holds the rules and the invite list. The bucket holds a 113.9 MB card snapshot and the meta store. The quality model and the precon table load, so a deck grades. The owner signs in, and the invite list refuses every other email. **The deploy half of the PR-22 gate stays open on the build walk alone.** A reader must upload a ManaBox export, then build a deck, revise it, and export it. Nothing else of PR-22 waits.
- **A merge to `main` deploys itself on Cloud Build** (2026-09-07, D-584, D-586). Two triggers of `us-central1` read the diff: `deploy-api` on `go/` and `docker/`, and `deploy-web` on `web/` and `firebase.json`. A merge of documents alone builds nothing. The connection is `decktome-repository`, and the repository is `nkramber-decktome`. A trigger of a 2nd-gen repository needs `--service-account`, and a project binding needs `--condition=None`. Both triggers name `gh-deployer@decktome-prod.iam.gserviceaccount.com`, because Cloud Build refuses the account Google manages (2026-09-07, D-588, #89). That account gained `roles/logging.logWriter`, and one identity deploys on both paths. The rerun of the merge of #88 passed every step in 2 minutes 1 second. The workflow `deploy` of GitHub Actions keeps a dispatch as the way back. No pull request spends an Actions minute now (D-578), and `make verify` runs every check on the Mac.
- **Every change starts on a branch, and `main` alone reaches production** (2026-09-07, D-579, D-583, D-585). The session broke both rules on the deploy day. `make where` prints the branch, the tree, whether `main` is current, and the state of the branch's pull request. `make hooks` installs a pre-commit hook that refuses a commit on `main`, refuses one on a merged branch, and runs `make ste-check` on a staged document. **Run `make hooks` once in a fresh checkout.**
- **Eleven defects came out of the deploy and the first walk** (2026-09-07, F-57 to F-67). Ten of them are fixed:

  - F-57, the Hosting header rule that cached every route for an hour.
  - F-58, the card index that never loaded at an idle instance. The API reads `--no-cpu-throttling` now.
  - F-59, the uninvited reader that reached the whole shell (2026-09-07, D-589, and D-590 after the first fix failed).
  - F-61, the meta job that stopped on 1 GiB. It reads 8 GiB and 2 vCPU now.
  - F-62, the thread that lumped every answer of a turn into one block, and dropped a decline (#80).
  - F-63, the Build entry that cleared the reader's collection with no word (#80, D-580).
  - F-64, the set phrase that read no abbreviation (D-581).
  - F-65, the setup command that wrote its own placeholder over the Firebase key.
  - F-66, the three commits stranded on a merged branch.
  - F-67, the deck placeholders that a reader with no deck never fills.

  F-59 and F-60 are fixed (2026-09-07, D-590, D-591). Every defect of the deploy and the first walk is closed.
- **The set matcher is a role** (2026-09-07, D-581, F-64). `ResolveAll` splits a phrase on its connectors, and only when the whole phrase settles nothing. The `setmatch` role reads what the table cannot, on the cost tier, after the table fails. Every code it answers meets the set table before it reaches a deck. The prompt is version 19, and the classify prompt did not change.
- **The product a reader sees is "Deck Tome"** (2026-09-07, D-577, corrects D-556). Two words in the browser tab, the header, and every string a reader reads. `decktome` stays for the domain, the repository, the module path, and the project id. `TomeMark` is the brand mark, and `public/favicon.svg` holds the same tome. The chat keeps `SparkMark` for the agent.
- **The spend cap takes a per-email override** (2026-09-07, D-576). `SPEND_CAP_OVERRIDES` reads `email:usd` pairs, and `0` turns the cap off for that email. The owner's account has no cap, and every other caller keeps $5.
- **`www.decktome.com` waits for its certificate.** The apex serves the app. The owner set `www` as a second custom domain that redirects to the apex, and its CNAME at GoDaddy names `decktome-prod.web.app`. Read it with `curl -sI https://www.decktome.com/`.
- **The next items.** The build walk of the PR-22 gate comes first, then PR-25 (D-547, D-555). PR-28 follows the first real feedback (D-557). PR-29, the casual corpus, still stands on the weak-axes plan (D-567). The next EDHREC read falls on 2026-09-14.

- **Feedback stays in Firestore, and one query reads every user** (2026-09-07, D-596). The owner asked to read every negative verdict without a walk of the user list, and asked about BigQuery. A collection group query over `feedback` reads every document under every user at once, so no caller needs the walk. `firestore.indexes.json` holds the two indexes it needs. `make feedback-list` prints the newest verdicts, and `VERDICT=up` and `LIMIT=` change what it reads. An emulator test proves two users answer one query. BigQuery is a valid store, and it earns its place at a volume this app is far from. Every verdict carries the uid as a field, because a group query answers documents and not paths. It carries `question_text` and `answer_text` too, filled from the stored session by the server. The stored schema is version 2, and a document of version 1 reads its user from the path.
- **F-76 is read, and the cause was not the pool** (2026-09-08, D-604). The runtime read came first, as PR-32 asked. Deck `u8FV7fc98qzNvRfsuJ5q` holds 77 cards, every one owned, and a buy cost of zero. The engine was right the whole time. A deck carries `commander_oracle_ids` and no ownership fact for the commander. `buyRows` of the web app wrote one when it found no entry: `owned: false`, price zero. Every deck named its commander as a card to buy. The web app invents nothing now. **A commander the reader truly lacks reaches no buy list**, because the deck holds no fact either way. That half needs a field on the deck, and PR-32 holds it. `KIND=decks ./scripts/read-session.sh <deck-id>` reads a deck now.
- **PR-32 holds the remaining work of this session** (2026-09-08). It runs right after PR #97. It fixes F-75, the commander name the index does not know. The row asks which card the reader means, with the best three by commander quality at the power they picked. It finds the cause of F-76 first, and no fix goes in without that read. It closes the smoke gap of #92 with a fixture deck of an owned-first shortlist. The roadmap holds the entry and sequence item 27.
- **The roadmap audit of 2026-09-08 found four stale rows.** F-74 read planned, and #97 fixes it. F-76 read "needs owner input", and the owner answered OQ-79. F-31 and F-32 read a gate run due since 2026-09-02, and revise gate run 9 with deck gate run 16 met both on 2026-09-04. Read a gate document before a row says a run is due.
- **Every deck carries its build metrics now** (2026-09-07, D-602, answers OQ-80). `Deck.build` holds the wall time and the count of repair turns. It holds the misses and findings of each turn, the shortlist size, the model spend of the build alone, and the commander source. The spend is the accumulator after the build less the accumulator before, because the question calls of the same turn are not this build. **No deck built before this carries it**, so the first read of "almost every deck goes to repair" needs new builds.
- **The `deploy-web` trigger reads `firestore.indexes.json`** (2026-09-07, D-603). `gcloud builds triggers update github` refuses a trigger of a 2nd-gen repository, and so does a PATCH of `includedFiles` alone. A PATCH of the whole trigger body to the Cloud Build API works.
- **F-76 needs one more read, and no fix goes in without it** (2026-09-07, OQ-79). The owner chose the hard filter, and it applies once the reader answers the pool question. `CommanderPool` already refuses an unowned commander under `POOL_RULE_OWNED_ONLY`, and `FromListOwned` reads the owned count of a commander from the collection. So the deck of session `23rplEQAMA0mtJ3QFtKO` passed both and still marks its commander unowned. The cause is not yet known.
- **F-75 waits for its row** (2026-09-07). The owner chose to ask which card the reader means. More than three names match. The row then offers the best three by commander quality, at the power the reader picked. Nobody built the row or its wording yet.
- **The first feedback of the deployed app found four faults** (2026-09-07, D-600, D-601, F-73 to F-76). Two are fixed here. A declared Firestore index reached no deployed project. `make feedback-list` failed on `decktome-prod` while every test passed. The web build deploys `hosting,firestore:indexes` now, and the trigger of `deploy-web` needs `firestore.indexes.json` in its included paths (F-73). A set group runs no model, so "LOTR" settled nothing. The row then asked a question the reader answered already. The group path reads the `setmatch` role now, as the name path does (F-74). The sample hand drew a 146 pixel image at about 340 pixels, and it asks for the 488 pixel image now (D-601).
- **Two faults of the same feedback wait on the owner** (2026-09-07, F-75, F-76, OQ-79). The reader wrote "Aragorn as commander", and no card carries that name alone. The build dropped it, took the delegation path of D-232, and picked Thranduil. No question asked which Aragorn. The set flow asks that question and the commander flow has no row for it (F-75). The same deck ran an owned-only pool and named a commander the reader does not own. The buy list holds that commander, with no price (F-76). D-63 makes the pool a weak signal for the commander. The reader read it as a promise, and OQ-79 asks which rule wins.
- **No store holds the time of a build or the count of its repair turns** (2026-09-07, OQ-80). The owner reads almost every deck through a repair turn, and every build as slow. `Usage` holds the calls, the tokens, and the cost per session. The repair turns reach the log alone, and nothing holds the time. OQ-80 asks which numbers to store.
- **`make feedback-list` reads one verdict by id now** (2026-09-07, D-600). `PROJECT_ID=decktome-prod go run ./cmd/feedback -id <id>` walks the users and needs no index. It reads `DocumentRefs` and not `Documents`, because `users/<uid>` is a missing document.
- **A stalled question goes out once more** (2026-09-07, D-599, answers OQ-78). The owner chose re-ask once over a plain skip, because a skip puts a default in the place of an answer the reader gave. `ReaskStalled` re-opens the key when the planner has nothing to ask and the session is not ready. It runs in the same turn, so the reader hears the question again at once. A key re-opens once, and the second stall falls to `CloseStalled`. **OQ-77 waits on the owner alone now.** Firebase reads "Authentication with Identity Platform" as no-cost up to 50,000 monthly active users (firebase.google.com/pricing, read 2026-09-07), and the app has three. The cost is not the question. The work is: a `beforeCreate` function to deploy, and a seven-second answer limit that fails a sign-up when it passes.
- **A picked option sets its slot as data, and a turn is never silent** (2026-09-07, D-597, D-598, F-72). The deeper read of the dead conversation. The client sends a question id and an option index, which name one value exactly. The engine threw that away. It made the option text, and it waited for the classifier to write a string the slot reads. Seventeen of the eighteen rows with options carry a typed slot, so every one held that failure. A row carries `option_values` now, and `applyOptionAnswers` sets the slot from the index after the classifier. Eight rows carry values, and a test names the nine that do not, so the count only goes down. **`CloseStalled` was already the net** (D-351, D-386): a session waits two turns and the next message closes the key, so no session died. The reader read no word of that, and a turn that asks nothing and is not ready says "I did not read your last answer" now. A test fails a turn that asks nothing, builds nothing, and says nothing.
- **A chat stopped for good on the power answer, and a reader can report one now** (2026-09-07, D-593, D-594, F-70, F-71). Session `oUZMC0F2vHe7GGl24LIP` answered the power row with "4 optimized", the option the row offers, and the slot stayed asked. The readiness gate waits on an asked key, and the agent never repeats a question, so nothing came after. `withAnswers` folds an option index into the message as the option text, and the prompt asks for "bracket 4". None of the five options of the row parsed. `power` reads a leading number of 1 to 5 now. **The stuck session stays stuck**, because the question never comes again. A reader unsticks one with the words "bracket 4" in the message box. `FEEDBACK_KIND_CHAT` and "Report a problem" under the message box give a reader a way to name a chat that stops (D-594). **This is the dead conversation of D-351 through another door.** The class is open: any answer the classifier can not read leaves its slot asked for good. OQ-78 asks the owner for the rule.
- **Read the local date for a record, and never the UTC date** (2026-09-07, D-595). The session wrote D-590 to D-592, F-68, F-69, and four correction passes with the date 2026-09-08. The machine read 2026-09-07 in the owner's zone and 2026-09-08 in UTC. Twenty-two dates moved back one day.
- **An email off the invite list never becomes an account** (2026-09-07, D-592, F-69). The owner read the refusal screen of D-590 and asked for less: no account at all. `InviteService.CheckInvite` answers a bare yes or no, and it needs no sign-in. The create-account form asks it first. An email off the list reads one red line, "Your email has not been authorized for beta access at this time.", and the form stays where it is. A limit of ten a minute per client address bounds the check, because an unbounded one reads the list with a run of addresses. A check that can not answer allows the attempt, and the API refuses the call after it. A sign-in of an account that exists meets the invite gate of D-590, so the check runs on the create path alone. **This stops the form, and not the Firebase API.** The web API key is public, so a caller who drives that API directly still makes an account. OQ-77 asks the owner about a blocking function of Identity Platform. **The account `the owner's second account` exists from the walk of 2026-09-07**, and no command here removes it.
- **The reader's card-pool choice wins over the classifier** (2026-09-07, D-591, F-68). The owner chose a collection and "Only cards I own", asked for the LOTR and Hobbit sets, and read "Pool: any card". `Context.PoolFromReader` marks a rule the request carried, and `apply` never writes over a marked rule. A collection that leaves drops the mark. D-371 still holds where the reader chose nothing. The flip also turned the buy list on, so this was not a display fault. F-60 is fixed beside it. `CollectionEntry.image_uris` carries the art of the printing the reader owns. The index of the day fills it, as it fills the colors and the price (D-396). `scripts/read-session.sh` reads one session of the deployed project and found both. **The smoke flow passed because of F-68.** It uploads a collection, so the reader's rule was owned_first, and the classify fixture answered any_card over it. The flow picks "Any card" now, which is the choice of a reader who wants one. A build from a collection needs a fixture deck of its own shortlist, and that is a gate of its own. `make read-session SESSION=<id>` runs it, and `use_decktome` puts the shell on the right project. **The shell of the owner exports `PROJECT_ID=wallabee-dev`**, so a decktome tool that reads that variable talks to the wrong project. The script reads `SESSION_PROJECT` and the active gcloud configuration, and it never reads `PROJECT_ID`.
- **The first fix of F-59 failed, and the second one is on branch `f59-invite-guard-cors`** (2026-09-07, D-590). The owner read the whole Build page from an account off the list, one day after #90 merged. `decktome.com` and the API are two origins. A browser reads no response header that the API does not expose, and `Access-Control-Expose-Headers` named `Connect-Protocol-Version` alone. So `Deck-Tome-Refusal` never reached the browser, and the gate took every refusal for an invitation. The gate reads the code now. `ListCollections` takes no argument and never answers `PermissionDenied` of its own, so that code has one cause. The header joins the exposed list. The gate also draws the app for a cleared reader alone. An error that is not the refusal closes the app and names the state, and the navigation waits for the same answer. **Read this before you trust a test of a browser path.** The tests of #90 built the error by hand with the header attached, which no browser can do. The Go test used a Go client, which has no CORS. `make smoke` runs same-origin behind the Vite proxy, so it never meets this fault. Three regression tests hold it, and each one fails against the code of #90.
- Branch `f59-invite-guard`, from `main` at 3af98a6, holds the fix of F-59 (2026-09-07, D-589). The interceptor sets `Deck-Tome-Refusal: not-invited` on the refusal, and a Go test reads that header over a real round trip. `InviteGateRoute` sits under the route guard and over the four protected pages. It asks the API once, holds every page until the answer, and shows a screen with the account and a sign-out. The gate is a deferred chunk, so the shell bundle keeps its size (D-320). The answer latches in `invite-state.ts`, because the pages share the probe's query key, and without the latch a refetch takes the page off the screen. Sign-out forgets the answer, and `signOutAndClear` moved to the auth feature for it. The header draws no navigation for a refused reader. The tree is green on Go build, vet, and the race tests of `auth` and `cmd/api`. `make lint-go` reads 0 issues. The web lint, the typecheck, and 284 tests pass. `make smoke` passes in 5.3 seconds, and `make ste-check` reads no finding. Six new tests hold the gate. **The refusal path has no smoke case.** That flow needs the allowlist on in the emulator. The branch also fixes three record gaps. D-580 had no row, and commit `cdd571f` cites it. The table read F-62 and F-63 as open, and PR #80 fixed both. This file named nine defects, and eleven came out.
- (2026-09-07, earlier the same day) Branch `m7-honest-bars`, from `main` at d2e19f5, held M-7, the axis diagnostic and the honest bars (2026-09-07, D-570, D-571). The pair bars sample evenly under the cap (F-51). The fit runs five folds, so every list is holdout once (F-52). A copy keeps its own key (F-55, found in the build). The gate document adds the misses per axis by kind and each precon against its own copy with the least-moved features. It names the ten precons that lose most. Four tests hold it, the tree is green on Go build, vet, the quality package under race, and golangci-lint. Quality gate run 14 is its gate (D-572), and the section "M-7, the read of run 14" below holds the read. The owner answered OQ-76 the same day (D-573): the bar reads each precon against its own copies, and the cross pairs stand as information. Under it run 14 reads FAIL on Commander synergy alone, and Standard and Modern pass. PR-31 parks. PR #79 holds M-7, and the owner reads the checks and merges it. **PR-29, the casual corpus, is next**, on a branch from `main` after the merge.
- The weak-axes plan is `docs/reference/weak-axes-2026-09-07.md` (2026-09-07, F-51 to F-54, D-567). Run 13 fails for four causes on record. One of them holds the synergy axis, the weights against sense, and the judge bar: the corpus weighs the casual rungs zero. The plan is M-7, the honest bars and the axis diagnostic, then PR-29 the casual corpus, PR-30 the commander reference, and PR-31 the 60-card breaks. The owner answered OQ-73 to OQ-75 the same day (D-568): M-7 first, the commander reference as one feature, and the work runs now. PR #77 merged, and the plan sits on branch `weak-axes-plan` from `main` as PR #78 (D-569). The owner reads the checks and merges it. **M-7 is next**, on a branch from `main` after the plan merges.
- PR #76 merged on 2026-09-07 (D-564), `main` is 402ec6f, and branch `records-pr27` is gone. It held the records of D-562 and the rename of D-563. The rename covers `go.mod`, the 215 Go files, the 11 proto options, the generated code, the doctor script, and the docs. The User-Agent strings keep the product name `mtg-deck-builder/0.1` with the new URL, and the product rename goes with PR-25 (D-556). **The corpus step of 2026-09-07 ran on branch `corpus-refresh`** (D-566). `make meta-refresh` took 96 minutes, and quality gate run 13 reads FAIL on the precon bar of every format. The section "The corpus step of 2026-09-07" below holds the read, and the F-50 fix sits beside it. PR #77 merged it on 2026-09-07 (D-569).
- The weekly EDHREC skip never held (2026-09-07, F-50, D-565), found when `make meta-refresh` read EDHREC whole four days after 2026-09-03. The lane took the newest commanders day for the last read, and the tournament lane writes that day's file first. The skip reads the last completed read now, `TestEDHRECSkipsInsideTheWeek` holds it, and the fix sits on branch `corpus-refresh` beside the records. The read of 2026-09-07 ran on under the old code, so the next read falls on 2026-09-14.
- The repository is `decktome` now, on GitHub and on disk (2026-09-07, D-563). Open the folder `~/Repos/decktome`. GitHub redirects the old name, the origin remote is `git@github.com:nkramber/decktome.git`, and the memory directory has a copy under the new key. The Go module path is `github.com/nkramber/decktome/go` now. The web storage key and the web package name keep the old name until the owner asks (D-563).
- PR #75 merged on 2026-09-07 (D-562), `main` is f27e4d1, and branch `pr-27` is gone. The owner ran the workflow `smoke` on the branch first, and it passed in 2 minutes 32 seconds. **The next code items wait on the owner.** The deploy of PR-22 by `docs/setup-gcp.md` unblocks PR-25 (D-555), and the first real feedback that PR-28 needs comes from the deployed app (D-557). Until then the free items stand: `make meta-refresh` daily with `make quality-gate` after it, and the weekly EDHREC read on 2026-09-14 (D-499, D-565).
- Branch `pr-27`, from `main` at 9e91959, held PR-27, the feedback harvest (2026-09-06, D-561), merged as #75 (D-562). `FeedbackService.SubmitFeedback` writes one document per verdict under `users/<uid>/feedback/<id>`, with the reason keys, the text, and the prompt versions of the moment. The API refuses another user's session or deck with `PermissionDenied`. The web feature `feedback` holds `Thumbs` and `FeedbackDialog`, and the pair sits in the five places of the design note. The smoke flow rates the built deck and reads the toast. The tree is green: Go build, vet, the whole race suite, golangci-lint, the web lint, typecheck, and 272 tests, `make store-check` on six stores, `make smoke` in 15.8 seconds, and `make ste-check`. The free gate is `docs/reference/pr27-gate-2026-09-06.md`. **The owner triggers the workflow `smoke` before the merge (D-313).** PR-28, the loop, follows the merge and the first real feedback. The section "PR-27, what the branch holds" below holds the moving parts.
- PR #74 merged on 2026-09-06 (D-560), and branch `records-decktome` is gone. It held the records D-554 to D-559, the domain note, the GoDaddy path of the deploy guide, and the feedback plan. **PR-27, the feedback harvest, is next**, on branch `pr-27` from `main`, with nothing left open.
- The feedback system is the immediate item (2026-09-06, D-557 to D-559). `docs/reference/feedback-2026-09-06.md` holds the design of PR-27, the harvest, and PR-28, the loop. The owner answered OQ-70 to OQ-72 the same day (D-559). The card thumbs sit on the tile foot and in the sheet. The harvest file holds the user's words, the deck list, the session id, and the uid, never the email. The loop runs on demand at $2 a cycle. The section "PR-27, what to build" below holds the order of work.
- The product is Decktome, and the owner bought `decktome.com` at GoDaddy on 2026-09-06 (D-556, answers OQ-69). `docs/setup-gcp.md` section 3 holds the GoDaddy path alone now, at the owner's word, and section 14 adds the Firebase records in the GoDaddy DNS tab. The .app ending waits until the app opens beyond invites. The shell title and the mark still read "MtG Deck Builder", and the rename goes with the manifest of PR-25.
- The owner asked for candidate product names and domains on 2026-09-06. `docs/reference/domain-names-2026-09-06.md` holds the method, the constraints, 136 names with their status on .com and .app, and the recommendation: curvebrew first, then sigilwright and binderscribe. Thirty-one names are free on both .com and .app as of 2026-09-06. The "binder" family is wide open, and every "deck" compound is gone. The shelf of real fantasy words is empty, so the identity names that survive are coinages on the root "sigil". The owner chose decktome the same day (D-556). Two facts stand on record. "Deck Tome" is a book-shaped deck box of Gamegenic and of FidgetThings, and Decktamer is a 2025 video game one letter away.
- PR #73 merged on 2026-09-06 (D-554), `main` is b6aa093, and branch `pr-23` is gone. The first run of the workflow `smoke` on `main` passed in 4 minutes 42 seconds of Actions time, with the browser cache empty (D-554). PR-25, the installable web app, is next on branch `pr-25` from `main` (D-547), and it waits for the deploy of PR-22 (D-555). Its phone walk needs the deployed URL, because a service worker and an install need HTTPS.
- Branch `pr-23`, from `main` at af8e75d, held PR-23, the Playwright smoke flow (2026-09-06, D-551 to D-553), merged as #73. The fake provider serves a whole build from schema-keyed fixtures (D-552). `classify.slot_fill.json` fills the slots of a lifegain Commander deck led by Karlov of the Ghost Council at bracket 3. `generate.json` and `repair.json` name deck 1 of deck gate run 16 with one swap. `make smoke` starts the emulators empty, the API over the trimmed snapshot with every role on the fake, and the web app. Then it runs one Playwright flow (D-553). The flow creates an account, uploads the fixture export, and asks for the deck with no spending limit. It reads "Legal" and 100 cards on the deck screen, and it downloads the deck list. It passes in about 12 seconds on the owner's Mac. `.github/workflows/smoke.yml` runs it on `workflow_dispatch` alone. The tree is green: Go vet, the race suite of the touched packages, golangci-lint, the web lint, typecheck, and 262 tests, shellcheck, and `make ste-check`. The free gate is `docs/reference/pr23-gate-2026-09-06.md`. The section "PR-23, what the branch holds" below holds the moving parts.
- PR #72 merged on 2026-09-06 (D-551), and branch `pr-22` can go. It held the four code parts of PR-22. The deploy half of its gate stays open in `docs/reference/pr22-gate-2026-09-06.md` until the owner runs the Google steps of `docs/setup-gcp.md`. PR-23 started before that half holds, on the owner's word (D-551).
- Branch `mobile-stage-a`, from `main` at 999c9dd, holds the mobile and engagement proposal and its decisions (2026-09-06, D-546 to D-548). `docs/reference/mobile-and-engagement-2026-09-05.md` proposes three stages. Stage A is PR-25, its own PR outside PR-22, after PR-23 (D-547). Stage B is PR-26 and waits on OQ-67, the channels and the events. Stage C, the stores, waits on request (D-548). The roadmap holds the PR-25 and PR-26 entries and sequence item 24. Engagement reads as six loops around the events of a player, with six counts and no tracking. Commit 8ba5b79 holds it all, and PR #71 is open against `main`.
- Branch `pr-22`, from `main` at dca456b, held the four code parts of PR-22 (2026-09-06, D-549, D-550), merged as #72 (D-551). The verifier answers an identity with the email, and the interceptor refuses an email off `config/allowlist` with one sentence (D-314, D-420). The agent service reads the spend ledger `users/<uid>/usage/<YYYY-MM>` before every turn, and it refuses a turn at the cap with the reset day (D-421). It writes the turn's cost after. `make allow` and `make disallow` write the list. The web app reads the four `VITE_FIREBASE_` variables, and `firebase.json` holds the hosting block (D-544). The tree is green: Go build, vet, the race suite, golangci-lint, the web lint, typecheck, and 262 tests, and `make ste-check`. Commit 95b77ea holds it all, and PR #72 is open against `main`. The deploy itself runs on the owner's account by `docs/setup-gcp.md`, and the PR-22 gate runs on the deployed URL. The ledger and the list have emulator tests, and `make store-check` runs them with the other three stores. All five passed on 2026-09-06. The free gate is `docs/reference/pr22-gate-2026-09-06.md`, and its second half waits for the deployed URL.
- PR #70 merged on 2026-09-06 (D-546), and branch `pr22-decisions` can go. It held the records of D-543 to D-545: the merge of #69, the answer to OQ-65, and the close of the reword guard question. PR-22 builds the direct Cloud Run origin, and the guard stays at 0.60.
- PR #69 merged on 2026-09-05 (D-543), and branch `tier0-snapshot` can go. It held the third item of D-531 (D-521, D-542): the trimmed card snapshot of 3.9 MB at `go/cmd/deck-gate/testdata/snapshot`. The free dry lane of the deck gate runs over it in the Go job of CI. `make deck-gate-dry` runs the lane by hand, and `make deck-gate-trim` rewrites the fixture. The three items of D-531 are done.
- PR #68 merged on 2026-09-05 (D-541), and branch `terse-bar` can go. It held the second item of D-531: the 47 terse conversations in the gate with expectations, and four classifier rules (D-535 to D-538). It held run 42 as the questions baseline (D-539), `docs/setup-gcp.md`, and the bracket rejudge (D-540). The third item is next on branch `tier0-snapshot`: the trimmed card snapshot for Tier 0 (D-521).
- The bracket rejudge ran on 2026-09-05 (D-540), on branch `terse-bar`, $0.28. `pr14a-bracket-gate-run1-judge2.md` reads 8 of 15 decks in agreement, the same count as the lane of 2026-09-02, and FAIL on the bar of 80 percent. Kinnan moved up to 5 and Yuriko down to 4, and Najeela reads 3 in both lanes. The bracket 5 misses stay the power signal of PR-14B. The suite `bracket-judge` has no baseline, and a FAIL run needs `-force` to become one. The document prints a JSON escape for an em dash in three reasons, as the first lane did in two (F-48, cosmetic).
- `docs/setup-gcp.md` is the guide from the domain to the deployed app (2026-09-05, on branch `terse-bar`). It reads the repo, the Google Cloud free-tier document, and dated third-party price pages, and section 17 names every source. Its estimate covers five users and three decks a week each. The month costs $8 to $20 in model calls, $0 in Google Cloud inside the free tiers, and $1 to $2 for the domain. It raised OQ-65, and the owner answered it on 2026-09-05 (D-544): PR-22 builds the direct Cloud Run origin, and Hosting serves the static app alone.
- PR #67 merged on 2026-09-05 (D-533), `main` is 4ee191c, and branch `precon-basic-lands` can go. It held the first of the three items of D-531: the whole-precon check reads the nonbasic printings alone (D-523, D-532, F-35). It held the Avengers fixture as the true export again, and the M-5 documents of D-529 to D-531. The second item is next on branch `terse-bar`: the 47 terse conversations join the catalog-only bar (D-522).
- The M-5 sheet is complete (2026-09-04, D-524, D-529, D-530): 60 of 60 items scored. The owner scored Items 33 to 41 by phone, one per exchange, and took every recommendation from Item 39 on. Then the owner asked the session to score Items 42 to 60 on its own recommendation (D-529). Each of those rows carries its reason in the free text. `make m5-report` reads: no threshold meets the 80 percent floor, at 38 to 44 percent precision, so the fit threshold stays at 0.35 (D-423, D-530). 7 of 29 scored refusals are better than the row they replaced, so the reword guard at 0.60 is too tight. The catalog list of the report is stale in part. The house_rules row dropped its Vintage option, the locked row left under D-260, and the colors row dropped its game fact. The owner reviewed the 16 rows of D-529 on 2026-09-05 and accepted them all. The sheet waits for the owner's commit.
- The set fixes merged as #66 on 2026-09-04 (D-528), `main` is 9843493, and branch `set-fixes` can go. The branch held the Marvel set fix (D-518, F-40, F-41). A bonus sheet never leads a set family beside a product. The set and precon questions join their choices with "or", and "Marvel" resolves to Marvel Super Heroes alone. It held the group set request too (D-525). Its parts: `set_groups` in the classify role at prompt version 18, `ResolveGroup` in the set table, `GATE_ARGS`, and probe 109. Probe 109 ran as question gate runs 36 and 37 for about a tenth of a cent each. The sets expectation held on both: msc, msh, pspm, spe, and spm. Run 36 exposed F-42: a partial run read FAIL on the two whole-set bars, and `make eval-check` took it as the run of record. The fix is D-526, a partial run never stands for a suite, and D-527 makes run 37 the probe of record. The prompt version moved, so the next full question gate run re-baselines (D-66). The M-5 scoring session is next (D-524).
- PR-15, the eval harness, is merged (2026-09-04, #65, D-517). `main` is c46d97d, the broken commit dbbcb4e is gone, and the tree is clean. Branch `pr-15` can go. PR-22 is next, on a new branch from `main` (D-494).
- The PR-15 paid gate ran on 2026-09-04 (D-514, D-515), $4.53 for three steps. Deck gate run 16 passes 25 of 25 with no gate flip against run 14 with 14b, and it is the decks baseline now. Tier judge run 4 reads 7 of 25 on the open bar, and the sweep stopped there. Question gate run 34 read 21 misses on the fourth bar. The read fixed the renderer, three expectations, and the run file of a FAIL on the branch. The two runs that complete the gate ran on the owner's word (D-516). Question gate run 35 passes with no miss, and revise gate run 9 passes 11 of 11. Every suite has a baseline: runs 35, 16, and 9. The gate document is `docs/reference/pr15-paid-gate-2026-09-04.md`, and it holds F-36 to F-39. The section "The PR-15 paid gate, run" below holds the read and the two commands. The owner committed the session's work on `pr-15` and merged it as #65. A second Mac follows `docs/setup-second-mac.md`.
- PR-15, the eval harness, merged as #65 on 2026-09-04, from `main` at f299289 (D-511, D-512, D-517), before PR-22 and PR-23. The free gate is `docs/reference/pr15-gate-2026-09-04.md`. The owner said go to the paid gate on 2026-09-04 (D-513), and it ran the same day (D-514). The branch holds all six slices: the run files, the compare, Tier 0 in CI, the golden expectations, the plan judge, and the sweep. The section "PR-15, what the branch holds" below holds the moving parts, and "PR-15, the eval harness, prepared" holds the plan and the four owner questions, OQ-57 to OQ-60.
- The two corpus items of 2026-09-04 merged as #64 (D-510). Branches `corpus-items` and `deck-gate-fixes` can go. The meta job skips the MTGJSON deck files when the deck list names the products of the stored table. Deck gate prompt 25 covers the precon exclusion of PR-24 for the first time, and it waits for a paid run. The section "The corpus items of 2026-09-04" below holds the read, and one owner question came out of it (OQ-56).
- The paid sweep after PR-21 ran on 2026-09-03 and 2026-09-04, $4.60 in all. Question gate 33 and its eval pass. Tier judge run 3 reads 8 of 24, with the bar open on the corpus. Deck gate 14 with 14b reads 24 of 24, and revise gate 8 passes. The one defect it found, F-34, merged fixed as #63 (D-509). The section "The paid runs of 2026-09-03, after PR-21" below holds the read. Branch `deck-gate-fixes` can go.
- PR-21, the share link and the print view, is merged (2026-09-03, #62, D-508). The free gate passes, `docs/reference/pr21-gate-2026-09-03.md`. The tree is green on `main`: Go build, vet, the `-race` suite, golangci-lint, `make ste-check`, the web lint, the web typecheck, and 262 web tests. Branches `pr-24`, `pr-14c`, `pr-20`, and `pr-21` can go. The section "PR-21, the share link and the print view, merged" below holds the moving parts. PR-22 is next, on a new branch from `main` (D-494).
- PR-20, the deck view and the card detail, is merged (2026-09-03, #61, D-507). The free gate passes, `docs/reference/pr20-gate-2026-09-03.md`. The tree is green on `main`: Go build, vet, the `-race` suite, golangci-lint, `make ste-check`, the web lint, the web typecheck, and 256 web tests. Branches `pr-24`, `pr-14c`, and `pr-20` can go. The section "PR-20, the deck view and the card detail, merged" below holds the moving parts. PR-21 is next, on a new branch from `main` (D-494).
- PR-14C, MTGTop8 and the casual 60-card decks, is merged (2026-09-03, #60, D-502 to D-506). Both lanes read with zero failures over 500 pages, and the Modern typical rung exists, 149 lists. The gate document is `docs/reference/pr14c-gate-2026-09-03.md`. Quality gate run 12 reads FAIL on the same two bars as run 11: Commander 0.78 and Modern 0.88 on the precon bar. The tree is green on `main`. Branches `pr-24` and `pr-14c` can go. The section "PR-14C, MTGTop8 and the casual 60-card decks, merged" below holds the moving parts. PR-20 is next, on a new branch from `main` (D-494).
- PR-24, the precon exclusion, is merged (2026-09-03, #59, D-496 to D-498, D-500, D-501). The free gate passes, `docs/reference/pr24-precon-gate-2026-09-03.md`. The tree is green on `main`: Go build, vet, the `-race` suite, golangci-lint, `make ste-check`, the proto check, and the web typecheck and 237 tests. Branches `pr-24` and `pr-14c` can go. The section "PR-24, the precon exclusion, merged" below holds the moving parts. PR-14C is next, on a new branch from `main` (D-494).
- The corpus step of 2026-09-03 is done (D-494 to D-499). `make meta-refresh` ran whole. Quality gate run 11 reads FAIL on two bars: the Commander precon bar at 0.78, and the Modern precon bar at 0.9465 against 0.95. The wider EDHREC read did not move the built decks: 19 of 24 grade bad under the refit. The section "The corpus step of 2026-09-03" below holds the read.
- The owner set the order on 2026-09-03: the corpus work, then PR-24, then PR-14C, then the paid runs (D-494, D-495). No paid run goes before the PR-14C code lands.
- PR-14B, the deck quality model, is merged (2026-09-03, #58, D-470 to D-493). Merged before it: PR-0a to PR-8, PR-7B, PR-10 to PR-13, the audits, the Phase 3B roadmap (#46), PR-16 (#47), PR-16B (#48), PR-17 (#49), PR-17B (#50), PR-18 (#53), the review fixes of PR-18 (#54), PR-19 (#55), its follow-ups (#56), and PR-14A (#57).
- The tree is green on `main`: Go build, vet, `-race` tests, golangci-lint, the proto check, `make ste-check`, and the web lint, typecheck, and 237 tests. Branches `pr-14a`, `pr-14b`, `pr-19`, and `tile-fixes` can go. The section "PR-14B, the deck quality model, built" below holds the moving parts, and no owner row waits.
- PR-14A, the bracket profile, is merged (2026-09-02, #57, D-459 to D-469). The current branch is `pr-14b`, and branches `pr-14a`, `pr-19`, and `tile-fixes` can go. The tree is green: Go build, vet, `-race` tests, golangci-lint, the web typecheck, and the proto check. Bracket gate run 1 reads FAIL on the band bar and the judge bar. Deck gate run 12 and its rerun 12b together pass all 24 prompts with no regression. The sections below hold the moving parts.
- PR-19, the chat and build experience, is merged (D-432 to D-458, #55 and #56). It holds the land-swap fix of the revision turn (F-31, D-448), the split land bucket of the shortlist (F-32, D-450), and the one-chat-one-deck delete (D-456). Branches `pr-19` and `tile-fixes` can go.
- The tree is green on `nits-and-fixes`: Go build, vet, `-race` tests, golangci-lint, web lint, typecheck, and 219 web tests. The emulator tests of the collection store pass, and `make lint` reports zero findings.
- Question gate run 31 passes every bar. The set deck gate passes 6 of 6.
- Every gate stands and passes: question gate 33, the set deck gate, deck gate 14 with 14b, and revise gate 8. The tier judge bar of PR-14B stays open at 8 of 24 (judge run 3).
- PR-9 is out of the MVP (D-256). Phase 3B comes before Phase 4 (D-316).

CAUTION: branch `pr-17` carries eight concerns. They are the contract, the Go side, the reference design, and the layout of D-331. They are also the Build menu, the one deck screen, the pool picker, and the speed of the app. Guardrail 10 asks for one. The owner chose to ship it whole (D-344).

## M-7, the read of run 14 (2026-09-07, D-570 to D-572)

`docs/reference/pr14b-quality-gate-run14.md` is the gate of M-7, free, 98 seconds under five folds. The profiles run once for the five, so the gate did not grow to five minutes. The bars read every precon now, and each precon also meets its own broken copies.

| Format | Cross bar, run 13 | Cross bar, run 14 | Own copy, run 14 |
|---|---|---|---|
| Commander | 0.78 of 20,000 | 0.86 of 99,933 | 0.88 of 945 |
| Standard | 0.88 of 50 | 0.87 of 730 | 0.97 of 40 |
| Modern | 0.86 of 20,000 | 0.85 of 99,754 | 0.99 of 1,635 |

The cross bar mixes two questions (F-56). Each precon beats its own copies almost always, and the cross bar fails, because the precon rung spans bad to good. In Modern, 7,264 of the 9,226 colors misses are pairs the detector passed on both sides. A weak product graded bad loses to the broken copy of a Challenger deck.

The ten precons that lose most in Modern are theme and intro decks of 2004 to 2014. In Commander they are the products of 2011 to 2015. The owner chose the own copies as the bar (D-573).

Commander synergy is the one axis weak on its own copy, 271 of 389. The break moves `card_rate` by 0.08 standard deviations, and `synergy` and `unseen_share` by 0.9 each. The weights on those two are +0.17 and -0.23, and the weight on `card_rate` is +0.98 (F-53, measured). Of the 118 own misses, 54 are the ladder's and 64 involve a precon the detector flags. The flagged products date from 2011 to 2015.

The 60-card breaks are visible on their own copies. Modern reads colors at 74 of 75 and copies at 61 of 61, and Standard colors at 3 of 4. PR-31 waits on OQ-76, because it fixes a break that works on its own copy. F-55, found in the build: 322 of the 1,150 bad holdout rows of Commander read the profile of another copy in run 13.

## The corpus step of 2026-09-07 (D-564 to D-566)

`make meta-refresh` ran from 01:36 to 03:12 local, 96 minutes, on branch `corpus-refresh`, and every source read with zero parse failures. The report per source:

- MTGO: 200 pages, 38 fetch errors, and 6,029 lists. The 38 are retired event pages, and more pages wait behind the cap.
- MTGJSON: skipped. The deck list of version `5.3.0+20260906` names the products of the stored table `5.3.0+20260903` (D-510).
- The cEDH database: 1 page.
- Topdeck.gg: 1 page and 1,144 lists, after two 429 answers with an 8-second wait each. The run of 2026-09-03 gave 31.
- EDHREC: 2,077 pages and 3 new lists. The weekly skip never held, and F-50 fixes it (D-565).
- MTGTop8: 255 pages and 215 lists. MTGGoldfish: 200 pages and 146 lists, and more decks wait behind the cap.
- The fit stored model `20260907T081149Z` over 48,766 lists and 1,562 commanders.

Quality gate run 13 is `docs/reference/pr14b-quality-gate-run13.md`, free, 58 seconds. It reads FAIL on the precon bar of every format, and it is the first quality run with a run file.

| Format | Great over precon | Precon over bad | Run 12 precon over bad |
|---|---|---|---|
| Commander | 0.97 | 0.78 | 0.78 |
| Standard | 0.98 | 0.88 of 50 | 1.00 of 5 |
| Modern | 0.99 | 0.86 | 0.88 |

Commander did not move. Every bar and every axis sits within 0.02 of run 12: lands 0.97, curve 0.96, colors 0.92, and synergy 0.78. The two sign changes are under 0.03 in size, `color_sources` and `empty_roles`. The four weights against the sense of their feature persist: `draw` -0.23, `wipe` -0.15, `land` -0.26, and `commander_decks` -0.13. The great and the good rungs hold 4,000 lists each, so the 1,144 tournament lists replaced others inside the cap.

Standard is a different model now. The typical rung exists, 68 casual lists from MTGGoldfish, 57 in train and 9 in the holdout. The synthetic bad rung grew from 40 to 370 lists, five breaks of each baseline and typical list. The precon bar reads 50 pairs instead of 5. The colors axis reads 0.38 of 8, and synergy 0.92 of 13. The other three axes read 1.00.

The Standard weights moved most. `card_rate` fell from +1.89 to +1.40, `synergy` from +1.44 to +0.98, `color_sources` from +0.82 to +0.35, and `avg_mana_value` from -0.21 to -0.87. Four weights changed sign: `land`, `curve_high`, `draw`, and `empty_roles`, each under 0.11 in size after the change. A model fit over 35 bad lists became a model fit over 320, so run 12 and run 13 do not compare on Standard.

Modern moved little. The great rung grew from 1,388 to 1,805 train lists and the bad rung from 1,935 to 2,205. The weak axes are the same: colors 0.51 and copies 0.89, with synergy at 0.86. Four small signs changed, each under 0.04 in size: `mana_turn_four`, `wipe`, `interaction`, and `fast_mana`.

The bracket 5 offer names The Jolly Balloon Man at 0.57, Niv-Mizzet, Parun at 0.50, and Etali, Primal Conqueror at 0.50. Etali holds 76 top cuts in 411 entries. No weight changed by hand (D-486). The judge bar stays open on the corpus (D-488, D-491).

## PR-27, what the branch held (2026-09-06, merged as #75)

`docs/reference/feedback-2026-09-06.md` is the design, D-557 to D-559 are the decisions, and D-561 holds the build calls. The seven steps of the plan are done, in order.

- `proto/mtg/v1/feedback_service.proto`: `FeedbackService.SubmitFeedback`, the `Feedback` message, `FeedbackKind`, and `FeedbackVerdict`. `make proto` wrote the Go and the TypeScript files, and they are new on the branch.
- `go/internal/feedback`: `Repo.Add` and `Repo.Get` over `users/<uid>/feedback/<id>`, the short names of the enums, the reason keys per kind, and the emulator test. `make store-check` runs six stores now.
- `go/internal/feedbacksvc`: the handler. `itemOf` reads the fields and refuses a bad one, `checkOwner` reads the object under the caller, and `Prompts` names the two prompt versions. `MaxTextBytes` equals `agentsvc.MaxMessageBytes`, and a test holds them equal.
- `go/cmd/api/main.go`: one session repo serves the deck service and the feedback service, and the feedback handler sits behind the auth interceptor.
- `web/apps/web/src/features/feedback`: `reasons.ts` with the keys and the labels, `use-feedback.ts` with the mutation, `feedback-dialog.tsx`, and `thumbs.tsx` with `Target`, `feedbackOf`, and `thanks`. `lib/api.ts` exports `feedbackClient`.
- `eslint.config.js`: `feedback` and `share` join the boundary table. `share` imports `deck` and `export`, as it did before with no rule.
- The placements: `QuestionCard` and `ThreadLine` take `sessionId`. `CardTile` and `CardGroup` take `feedbackDeckId`, and `CardDetail` takes `deckId`. `DeckView` mounts the summary pair and the deck pair, and `data-testid="deck-summary"` names the summary paragraph.
- `e2e/smoke.spec.ts`: the flow clicks "This helped" under "Rate this deck" and reads the toast.

CAUTION: `go build ./cmd/api` from the `go` directory writes a 99 MB binary at `go/api`, and `TestNoCommandBinaryIsPresent` fails on it. Build with `go build ./...` from `go`, or remove the file.

CAUTION: the reason keys live in two places, `go/internal/feedback/feedback.go` and `features/feedback/reasons.ts`. A new key goes in both, and a test on each side pins the list.

## PR-23, what the branch holds (2026-09-06)

- `go/internal/llm/fake.go`: the fixture of role R is `R.json`, and a role with several schemas has `R.<schema>.json`, which wins (D-552). `Roles()` names the role of a schema fixture once.
- `go/internal/llm/fixtures/`: `classify.slot_fill.json`, `classify.score_questions.json` (no score, so a planned row goes out as written), `ask.json` (no phrasing, so the catalog text stands), `generate.json`, and `repair.json`. The two deck fixtures are the same bytes.
- `go/cmd/deck-gate/smoke_test.go`: the deck of the fixture resolves whole in the pool of its slots over the trimmed snapshot. It holds 100 cards, and it passes the rules engine. `trim.go` keeps the cards of the fixture on every `make deck-gate-trim`.
- `web/apps/web/e2e/smoke.spec.ts` and `web/apps/web/playwright.config.ts`: one flow, one worker, Chromium. The vitest exclude keeps the flow out of `pnpm test`, and the tsconfig includes it in the typecheck. `@types/node` 26.3.0 joins the dev dependencies for the file reads of the spec.
- `scripts/smoke.sh` and `make smoke`: the stack on the ports of `make dev`, so the script refuses to start over a running stack. The keys leave the environment, and `LLM_<ROLE>_PROVIDER=fake` names the fake for every role. The dev server binds 127.0.0.1, because plain localhost resolves to ::1 on macOS and the port wait failed on the first run.
- `.github/workflows/smoke.yml`: `workflow_dispatch` alone, one job, the Playwright browser cached on its version, and the report as an artifact on failure.
- The CI runner has shellcheck 0.9.0, and the Mac has 0.11.0. Both mark a function that only a trap calls as unreachable when the script ends with `exit`. The code is SC2317 per line in 0.9.0, and SC2329 on the function in 0.11.0. `scripts/smoke.sh` disables both codes on `cleanup` with the reason. `docker run --rm -v "$PWD:/mnt" koalaman/shellcheck:v0.9.0 scripts/smoke.sh` proves a script under the runner's version.
- The smoke message refuses a spending limit on purpose. The budget row fires on an any-card pool (D-168). A decline works only on a question that is out (D-93), so the word rule is what closes the slot without a number. Run 2 of the flow found that: it asked the budget question and stopped.

CAUTION: a no-key `make dev` builds the fixture deck for any first message now (D-552). A local read of the question flow needs real keys in `.env`, as before.

## The numbers, and why none of them compare with `main` now

The audit changed two prompts (D-302). The generate prompt is at version 10 and the classify prompt at version 14. The shortlist omits the commander, the house format has a sideboard sentence, and the dead `two_plans` and `owned_mode` facts left the classify schema. So the baselines below measure the code before the audit, and the next run re-baselines (D-66).

| Measure | Baseline | Where measured |
|---|---|---|
| Question gate | PASS 25 of 27 counted, 2 invented, 0 premature, 0 lint, $0.16 | run 27, 2026-08-28, ask prompt 13, catalog of D-290 and D-294 |
| Question eval | 19 bad of 382, holdout 8 of 111 (7.2 percent) | eval of run 25, 5 conversations unjudged |
| Deck gate | PASS 18 of 18, 2 repairs, $1.09 | run 8, 2026-08-28, generate prompt 9 |
| Revise gate | PASS 6 of 6, $0.29 | run 2, 2026-08-28 (D-296) |
| Loop | off since 2026-08-26 | seven starts, nothing kept |

CAUTION: `tune-check` paired zero questions between run 24 and run 25, because the catalog and the prompt changed. The paired guard says nothing across that line, and the whole-run margins carry the verdict. The eval leaves a conversation unjudged when the judge returns fewer verdicts than questions. So 382 is the honest count, not a drop from 435.

## The audit of 2026-08-29, what changed

`docs/audit-2026-08-29.md` holds the report. The points a session needs first:

- A turn during a build gets `CodeAborted` with "a build is in progress". The build and its store writes run detached from the client, so a disconnect keeps the deck (D-303). The web app warns before it leaves a page mid-build.
- Every session, deck, and collection id passes `gzstore.ValidID`. An id with a slash is `CodeInvalidArgument`.
- `go/internal/gzstore` holds the gzip and JSON helpers of the three Firestore repos, with one inflate limit.
- The buy cost and the deck cost count the commanders and sum copies per Oracle id. A locked, kept, or commander card above the mana cap stays in the pool (`Revision.Exempt`).
- The house format offers paper cards only (D-306).
- `deck-gate` and `revise-gate` exit 1 on FAIL. A nil cost prints `unpriced` in every gate. `revise-gate` has `-only` and `-dry`. `questions-eval` and `candidates-review` refuse an existing output (D-65).
- The Makefile runs bash with `pipefail`. `candidates-review` guards on a filled score cell.
- `autotune.sh` builds `tune-check` after the branch switch, never switches in a dry run, and restores the start branch on exit.
- The go job checks out with `fetch-depth: 0`, so `make llm-defaults-check` finds a merge base. A weekly `vuln` job runs govulncheck alone on Monday 06:00 UTC (D-305).
- The STE checker flags passive voice, modals and perfect tenses, -ing forms, and the 20-word step limit (D-304). Dated records are exempt: `docs/reference/pr[0-9]*`, the session logs, `docs/audit-*`, and testdata.
- Every code comment states a rule and cites a decision id. No comment carries a date, a session id, a run number, or "the owner".

## PR-13, what it holds (2026-08-29)

- `DeckService.ExportDeck(deck_id, format)` returns the text and a file name. `EXPORT_FORMAT_ARENA_TEXT` is the ManaBox shape, and `EXPORT_FORMAT_BUY_LIST_TEXT` is one "count name" line per card to buy (D-309).
- `go/internal/export`: `ArenaText`, `BuyList`, `BuyListText`, `FileName`, `Render`. The Arena line names the owned printing when the card is owned, else the default paper printing (D-307). The line carries the full card name, and the index resolves it without ambiguity. `TestArenaTextRoundTrip` is the gate.
- The buy list is the shortfall of the commander, the main deck, and the sideboard, summed per Oracle id, and the upgrades apart (D-308). A commander with no entry in the card list counts as one card to buy.
- `web/apps/web/src/features/export`: `export-panel.tsx` (four buttons, the status line, the buy list with a Scryfall link per card) and `buy-list.ts` (the same rules as the Go package, for the screen). The deck view mounts the panel under its header. `export` is a leaf feature, and `deck` imports it.
- The web app reads the text from the API, so the copy and the file match what the API renders.

## PR-16, what it holds (2026-08-29)

CAUTION: three parts of this slice changed on 2026-08-30. The sidebar became a top bar (D-328), the light theme went (D-330), and the color identity of a deck went (D-329). The primitives, the route split, and the test helper below still hold.

- Nine primitives in `web/apps/web/src/components/ui`: Button, Input, Label, Textarea, Checkbox, Card, Skeleton, DropdownMenu, and the toast. Each one is a hand-written shadcn shape on Radix, with relative imports. The shadcn command line installs a `@` alias, and the import boundary of this repo reads relative paths only.
- `src/styles/tokens.css` holds the tokens. The dark theme overrides the neutrals, the surfaces, and the link. The six mana colors do not change. A script in `index.html` paints the class before the first paint.
- The shell is a sidebar on a desktop and a bottom tab bar on a phone. A media query picks one of the two. Two navigations with one name fail the axe landmark-unique rule.
- Sign-out moved into the account menu. A failed sign-out reports as a toast.
- Five things leave the first paint: `firebase/auth`, the Connect client, the five pages, the two shell menus, and the toast host.
- `renderAt` in `src/test-utils.tsx` is async now, and every test awaits it. It warms the page modules and the auth SDK, then flushes one act.
- `src/test-setup.ts` adds four jsdom stubs that the Radix menus need.

## PR-17, what the branch holds (2026-08-30)

The web side, in one list:

- The library grid at `/decks`: the name, the commander, the mana pips, the format, the power, the count, the cost, the date, and a favorite star. Search by name and commander, and filter by format, power, and favorites. Every filter runs on the server.
- `/decks/<id>` is the one screen of a deck (D-335). It holds the deck, the actions, and the conversation that built it.
- `/session/<id>` holds a build with no deck. It hands the reader to the deck's address the moment a turn ends with a deck.
- `src/features/workspace` is the one feature with a path to both chat and deck.
- Build in the header is a menu of the collections (D-332), and a signed-in reader lands there (D-334).
- The look of the reference design: three faces, the navy and gold palette, one top bar, dark alone.
- A "Build from" picker sits under the title of a new chat (D-336). It names the pool, and it changes it without a trip to another page.
- The deck screen carries the version history and the compare (D-343). The decks of one chat are the versions of one deck.

CAUTION: no chunk of this app loads behind a Suspense boundary (D-338). React holds a committed fallback for 300 ms, and it holds every later reveal with it. `src/app/deferred.tsx` replaces `React.lazy` everywhere. Do not put `React.lazy` back.


The Go side, in one list:

- `DeckService.UpdateDeck` writes the name and the favorite mark, and `DeleteDeck` removes a deck for good. Both are additive, and `buf breaking` passes.
- `Deck.favorite` and `Deck.card_count` are new. The list view carries no cards, so it sets `card_count`. The deck list on screen reads `cards.length` today and always shows zero.
- `ListDecks` takes `page_size`, `page_token`, `format`, `favorite`, `query`, `power_bracket`, `power_sixty_step` (D-324), and `session_id` (D-343).
- The filter runs in Go over the rows Firestore returns, not as a query. One read serves every filter, and no composite index has to exist. The scan cap is 500 rows.
- The stored document gains six flat fields. A deck written before PR-17 reads them as zero, and a rename fills them.
- `internal/decks/filter_test.go` is new, and the emulator tests cover Update, Delete, and the filter. `make store-check` passes.

CAUTION: a `t.Cleanup` can not delete from Firestore. Go cancels `t.Context` before a cleanup runs, so the delete fails and the next run reads the leftovers. Each emulator test takes a fresh user id instead.

## The look of 2026-08-30 (D-328 to D-331)

The owner gave a reference design, and it settles the look. Our terms win over its terms: the app is the MtG Deck Builder, and its entries are Build, Decks, and Collection.

- Three faces, three jobs. Cinzel engraves a heading, a label, and a button. Crimson Pro reads a paragraph. JetBrains Mono carries an id, a count, and a date. Each one ships with the build.
- The palette is navy, gold, and purple, with parchment for text. The radius is 4 px, and every panel carries a faint gold hatch.
- The shell is one top bar (D-328). The sidebar, the phone tab bar, and the theme toggle are gone.
- Dark is the only theme (D-330). `src/lib/theme.ts` and the no-flash script left with it.
- A deck has one screen and one address (D-335). `/decks/<id>` holds the deck, its actions, and its conversation. `/session/<id>` holds a build with no deck, and it hands over the moment a deck exists.
- `src/features/workspace` is the one feature with a path to both chat and deck. The lint carries that rule.
- A deck owns the whole page, and the conversation docks at the bottom left (D-331). A History control opens the thread over the composer. Before a deck exists, the conversation is the page.
- The identity wash left (D-329). Color of the game shows in a mana pip and a rarity dot.
- Build in the header is a menu of the collections (D-332). It sets the pool and opens a chat. With no collection it offers the way to the upload screen (D-334).
- A signed-in reader lands on Build, not on the collection (D-334).

CAUTION: a trigger of a Radix menu must pass on every prop it takes. The ref is among them, and Radix measures the trigger through it to place the panel.

A trigger that keeps only the props it names drops the ref. The panel then lands at the top left corner, outside the window, and a click never reaches it. Both menus of the shell carried this defect. jsdom has no layout, so no test there sees it. Playwright found it in one run.
- The binder head sits under the controls, the buy list opens on request, and a page holds 75 rem (D-333).

A session reads its own work with Playwright, and it measures rather than looks. Four checks ran on 2026-08-30.

- The deck page holds 80 percent of a 1440-pixel screen.
- The buy list carries no `open` attribute.
- The Continue-to-chat control does not move when a reader picks an upload.
- The Build menu lists every collection.

## The palette of D-327 (2026-08-29), now amended

D-311 kept every surface neutral and let the card art carry the color. The app read as boring, and the owner said so twice. The five colors of the game are the palette now, and dark leads.

- `src/styles/tokens.css` holds the dark base and a `.light` class. Dark is the class-free base, so a light reader carries the class.
- `src/features/deck/color-identity.ts` reads the color identity of a deck from its commander. `identityVars` sets `--identity-a` and `--identity-b` on one element, and the `identity-wash` and `identity-rule` utilities read them.
- Each card role carries its own hue through `roleToken`.
- `ManaPips` shows an identity as pips, and its screen-reader label names the colors in words.
- The ground of every page carries two soft lights and a fine grain. A panel takes a hairline of its own light along its top edge.
- The collection screen shows the binder: the count, the unique cards, the rarity spread, and the art of the rarest ten cards. It came forward from PR-18.

CAUTION: the binder head calls `GetCollection`, and the answer carries every entry. The owner's export holds 2,657 rows and 4,952 cards, so one page load moves about one megabyte. PR-18 stores the summary on the collection, and the head then reads no entry at all (D-392).

## Reading the app without the owner (2026-08-30)

`@playwright/test` is a dev dependency of `web/apps/web` now, and Chromium sits in the local cache. A session can read its own work.

- Start the Auth emulator and the Vite dev server. Then drive Chromium with a script that answers every `/mtg.v1.*` call from canned JSON, so no backend has to run.
- A Connect Timestamp is an RFC 3339 string in JSON, not a `{seconds}` object. A `{seconds}` stub fails with "cannot decode message google.protobuf.Timestamp".
- The script lives outside the repo, in the session scratchpad. It must run from `web/apps/web`, or the bare import of `@playwright/test` does not resolve.
- PR-23 holds the real smoke flow (D-313). This is a reading tool, not that.

## PR-17B, the set filter (F-29, D-373 to D-383)

Branch `pr-17b` holds it. The app never applied a set as a constraint, and a deck asked for one set held cards of any set.

The eight decisions, in short:

- D-376: a set name resolves to a whole set family, through the Scryfall parent link. "The Hobbit" gives `hob` and `hoc`. A phrase that names two base sets asks.
- D-377: the snapshot carries a fourth file, `sets.json.gz`, from the `/sets` endpoint. No bulk card file holds `parent_set_code`.
- D-378: a set limit never filters basic lands.
- D-379: inside a set limit the theme ranks the shortlist, and neither the theme cut nor the role caps apply.
- D-380: a family under 70 nonbasic cards in the deck colors builds no deck. The floor is 35 for a 60-card format.
- D-381: a card the reader named by name beats the set limit, and the deck marks it.
- D-382: ramp cards and nonbasic lands come from outside the sets only after the reader says yes.
- D-383: `DeckCard.outside_requested_sets` carries the mark, and the deck screen shows it in red.

CAUTION: a card does not have one set. The snapshot of 2026-08-31 holds 988 paper sets over 34,599 playable Oracle cards, and 16,247 of them hold printings in two or more. `Card.set_codes` is a list, 80,193 pairs in all, about two megabytes.

CAUTION: a name prefix can not build a set family. `ltc` is "Tales of Middle-earth Commander" and its base set is "The Lord of the Rings: Tales of Middle-earth". The two names share no prefix, and only `parent_set_code` links them.

CAUTION: the roadmap gate line named two cards that do not exist. The snapshot holds "Thranduil, the Elvenking" and "Smaug, Wicked Worm", both with a comma. Smaug the Impenetrable is in `hoc`, not in `hob`, so only the family rule of D-376 satisfies that line.

What the branch changed, in one list:

- `internal/cards/sets.go` is new: the set table, the family walk, and the resolver. `internal/cards/index.go` fills `Card.set_codes` in the printings walk it already runs, and it shares one string per set code.
- `scryfall.Client.Sets` reads `/sets`. `cards.Refresh` stores the file, and `cards.BackfillSets` fills a snapshot stored before the file existed. The worker calls it every cycle.
- `candidates.Request` takes `SetCodes` and `OutsideRoles`. `CountInSets` and `SetFloor` answer the viability floor, and `CountManaInSets` answers the mana row.
- `generate.Request.SetCodes` marks every deck card the sets do not hold, and one warning counts them.
- `agentsvc` runs the floor twice, before the commander pool and after the colors settle. `ErrThinSet` carries the counts.
- The classify prompt is at version 15 and reports `set_names`. The catalog holds two new rows. The snapshot is at version 3.
- `web/apps/web/src/features/deck/card-tile.tsx` shows the red mark.

CAUTION: the role caps were the second cut, and the first dry run found it. The "other" cap of 10 cut 40 of the 128 cards the Hobbit family offers in black-red. The shortlist reached 61, and a Commander deck needs 99. Both cuts lift under a set limit now.

## The paid runs of PR-17B (2026-08-31)

| Run | File | Result |
|---|---|---|
| Set deck gate 1 | `pr17b-set-gate-run1.md` | PASS. 6 of 6 decks, 0 blocks, 0 notes, 0 repairs. $0.3539 over 306 seconds. |
| Question gate 29 | `pr7-question-gate-run29.md` | FAIL. 29 of 30 catalog-only, over the bar of 25. Two failures, and neither comes from the set filter. $0.16. |
| Question eval 29 | `pr7-question-eval-run29.md` | 11.2 percent bad on the holdout, from 6.8. Read the caution below before you act on that number. $0.09. |
| Question gate 30 | `pr7-question-gate-run30.md` | FAIL. 27 of 27 catalog-only, 0 lint, 0 dead ends. 7 premature, which found two more defects. $0.16. |
| Question gate 31 | `pr7-question-gate-run31.md` | **PASS.** 27 of 27 catalog-only, 0 lint, 0 dead ends, 0 premature. $0.16. |
| Question eval 31 | `pr7-question-eval-run31.md` | 7.9 percent bad on the holdout, from 6.8 on run 28. $0.09. |
| Question gate 32 | `pr7-question-gate-run32.md` | **PASS.** 27 of 27 catalog-only, 0 lint, 0 dead ends, 0 premature. Prompt version 16. $0.16. |
| Question eval 32 | `pr7-question-eval-run32.md` | 22 bad of 393, from 27 of 413. The three changed rows fall from 11 findings to 4. $0.09. |
| Revise gate 3 | `pr12b-revise-gate-run3.md` | FAIL, 7 of 8. It found D-391. $0.57. |
| Revise gate 4 | `pr12b-revise-gate-run4.md` | **PASS, 8 of 8.** $0.54. |
| Revise gate 5 | `pr12b-revise-gate-run5.md` | FAIL, 10 of 11. It found F-32. $0.81, 9 minutes, 12 turns. |
| Revise gate 6 | `pr12b-revise-gate-run6.md` | FAIL, 9 of 10. Revision 9 passed with 14 basics swapped and no repair turn. Revision 1 skipped its question, a variance the prompt now closes. $0.67. |
| Revise gate 7 | `pr12b-revise-gate-run7.md` | **PASS, 11 of 11.** Every land turn asked, then swapped. $0.74, 12.5 minutes. |
| Deck gate 11 | `pr8-deck-gate-run11.md` | **PASS, 24 of 24.** Prompt version 11, the split land bucket. 3 repair turns. $1.46, 22 minutes. It found F-33. |

The revise gate gained a set-limited base on 2026-09-01. A revision reads the same slots as the build, so base 3 is the one run that proves the set filter survives a revision. It does. The revised deck holds Arcane Signet, Delighted Halfling, and Elvish Mystic, and all three are in The Hobbit Eternal. It holds no Sol Ring, which is in neither set.

Run 3 failed on one revision, and the failure had nothing to do with the sets. The model removed six cards and added seven under a mana cap, and the engine blocked the whole deck for one card. D-391 adds the trim, and run 4 passed.

## The question-quality pass of 2026-09-01 (D-387 to D-389)

Eval run 31 found 10 warranted defects over 413 questions, and 6 sat on the budget, the colors, and the format rows. Seven changes answer them, and run 32 measures the result.

| Measure | Run 31 | Run 32 |
|---|---|---|
| Questions scored | 413 | 393 |
| Not warranted | 27 | 22 |
| Detailed findings on budget, colors, and format | 11 | 4 |
| budget | 6 | 1 |
| format | 2 | 0 |
| colors | 3 | 3 |

Gate run 32 passes every bar: 27 of 27 catalog-only, 0 lint, 0 dead ends, 0 premature.

The three colors findings that remain are ones this repo refuses, and each cites a rule that predates them.

- "Whatever is winning" is not a color delegation. The corpus routes "whatever" nowhere, after gate runs 11 to 13 read it as house rules six times (D-111).
- The Atraxa question fires on turn 1 of a conversation whose whole point is a misspelling. The reader wrote "Atraxa, Praetor's Voice", and the card is "Atraxa, Praetors' Voice". The index does not hold the first spelling, so the rule of D-388 claims nothing about it. Guardrail 4 and D-140 both say the engine states nothing it can not prove. The reader corrects the name on turn 2, and the row does not fire again.
- The duplicate color question of conversation 7 follows a format decline. The reader named Pioneer, the decline row retired every open question (D-125), and a retired row asks once more (D-195). That is by design.

CAUTION: the whole-run bad count drifts. Run 28 found 14 of 391, run 29 found 22 of 398, and run 31 found 27 of 413. The step from run 28 to run 31 is 13 questions. D-230 measures the judge noise at up to 9 between two runs of identical code. Each single step sits inside that band, and the sum does not. Watch it on the next run.

The drift belongs to no part of this branch. Not one bad question sits on `set_unresolved`, on `set_outside_mana`, or on `power_sixty`. Every one sits on `budget`, `colors`, `format`, `power_commander`, `theme`, `pool_thin`, `named_card_role`, or `format_unsupported_open`. `budget` leads every eval: 3 of 6 detailed findings on run 28, and 6 of 17 on run 31. The gate is the contract, and the eval is advisory (D-136).

Three runs measure the net of D-351, because run 29 ran without it and the other two ran with it. The table below is the whole case for D-386.

| Run | The net | Conversations that ended before the script did | Questions the reader answered | Turns the net healed |
|---|---|---|---|---|
| 29 | none | 6 | 316 | 0 |
| 30 | closes on the reader's first reply | 45 | 264 | 82 |
| 31 | closes after two turns (D-386) | 11 | 308 | 34 |

The grace period returned 44 of the 52 answers run 30 threw away. It also kept every gain: run 31 reports no dead end.

Run 30 found two more defects beyond the eager net. The gate asked for the `commander` key while `commander_pick` answered it, which called 6 conversations premature. Each of the six holds a commander. `required` reads all four commander keys now.

CAUTION: the eval numbers do not compare cleanly. Run 28 scored 391 questions and found 14 not warranted. Run 29 scored 398 and found 22. The margin is 8 questions, and D-230 measures the judge noise at up to 9 between two runs of identical code. `tune-check` paired zero questions across the line, because the catalog and the prompt both changed. So the run is neither proof of a regression nor proof of none.

The three set conversations reach neither the bad-question list nor the missing-question list of the eval. Every one of the 8 extra findings sits on a `commander_pick`, `colors`, or `budget` row of a conversation that predates this branch.

The set gate proves every deck line. A session read every card of all six decks against the raw snapshot. Only two decks hold a card outside their family, and the app marked every one. Prompt 21 holds 8 lands of the mana fill, and prompt 22 holds Sol Ring. Prompts 19, 20, 23, and 24 hold none.

A free dry run proves the thin-set refusal of D-380. A mono-black Hobbit commander gives 68 cards, the floor is 70, and the turn ends with the count and no deck.

The owner settled every failure of run 29 and run 30 (D-384, D-385, D-386). Run 31 passes every bar.

- The 60-card power row is fixed now, so the ask role sends it word for word. The acronym rule reads the options beside the text, because a reader meets both at once.
- The gate runs the D-351 net, as `agentsvc.Chat` does. A test reads both turn loops and fails when they guard the net on different conditions.

The three defects below are closed. This file still records them, because a session that reads a run 29 or run 30 document needs to know what it holds.

CAUTION: question gate run 29 fails on two things, and PR-17B causes neither. Run 28 and run 29 built the identical conversation for every named case, so a session compared them line for line.

- The linter found 2 questions that say FNM with no expansion. D-374 changed the `power_sixty` row on 2026-08-31 and added that rule, both after run 28. Run 29 is the first run under them. The ask role kept the expansion 4 times and dropped it 2 times, out of 17 firings of the row.
- The dead-end check of D-357 reports 10 conversations, and 8 of them wait on the budget. Run 28 held the same conversations and the same open keys. The check did not exist then, so run 29 is its first full run.

CAUTION: the root cause of those 10 was the gate harness, not the agent. `agentsvc.Chat` closes the open questions on a turn that asks nothing new, and the gate did not. All 10 sat on the last scripted turn, and production heals every one. D-385 puts the net in the gate.

A session marked the three set conversations as probes after this run, so they now sit outside the count. They ran as gate conversations, which raised the counted set from 27 to 30. The bar holds either way, and the next run counts 27 again.

Two defects of the set filter reached the gate, and both are fixed:

- The role caps were a second cut beside the theme cut. The "other" cap of 10 dropped 40 of the 128 cards the Hobbit family offers in black-red. A free dry run found it before any paid call.
- The set row asked about Tarkir on two turns in a row, because nothing recorded the phrase it had already named. It follows the D-210 rule now.

CAUTION: the eight snapshots on the owner's disk held no set file. A session wrote one into `20260831T090157` by hand, so the tests and the gates read a real table. Every other version falls back to a derived table with no family link, and the worker fills the newest one on its next cycle.

## PR-18, and the review fixes on `nits-and-fixes` (2026-09-01)

PR-18 merged as #53. A review the same day found three defects and four gaps, and the owner approved every fix (D-398 to D-404). The fixes sit on branch `nits-and-fixes`, which starts at `main`.

- The binder filter, the sort, and the search run on the server (D-398). `GetCollectionRequest` takes a `BinderFilter` and a `BinderSort`, the page token names both, and the answer carries `matched_rows`. `collections.Filter` and `collections.Apply` hold the rule, and `binder_test.go` proves it.
- The summary lists every set and counts every card type (D-398). `CollectionSummary.sets` replaces `top_sets`, `by_type` is new, and `by_color` left (D-402). The repo reads the card index for the summary of a document stored before the field existed, through `Repo.WithIndex`.
- A derived document id belongs to its hash alone (D-399). `Repo.Put` creates the derived document, and when it exists with another hash the upload takes a fresh id. An emulator test and a service test both walk the sequence: upload X, replace with Y, upload X again.
- The upload dialog offers a choice when a collection is active (D-400): replace it after a diff, or add a new one. The name field shows for a new collection alone.
- The art loads in buckets of 100 rows (D-401). One scroll to the foot of 2,471 rows made 12 page calls and 24 art calls.
- A replacement whose file resolved no row is FailedPrecondition (D-403), and the page names no collection when an answer has no id.
- `ImportCollection` reads its file through `parseUpload`, as `DiffCollections` does. Delete invalidates the `["collection"]` prefix. The head and the binder show an error state with a retry. `statsOf` and `artIds` left the web app.
- The decline control of the budget question reads "No budget" (D-404). A declined budget stores no cap.
- The chat refuses a turn with no card index (D-405). The web app shows the failure as it shows any other.
- A decline applies its rules on both paths (D-406). The "You decide" control on the format question set no format before. No power row or commander row fired, and the build started with no format. `State.DeclineKey` holds the format default and the commander delegation, and the classifier path calls it too.

CAUTION: `make dev` runs the API binary of its start. A Go change needs a restart, or a browser read proves the old server. A session read the merged binary first, and the run reproduced the D-399 overwrite and ignored every filter. The second read built the working tree into the scratchpad and ran it on port 8091 with the environment of `scripts/dev.sh`. The script rerouted the collection and card calls to it with `page.route` and `route.fetch`. Every number below comes from that second read.

A session read the screen in a real browser with Playwright, against the real API and the 2,547-row export in `go/internal/collections/testdata`. It measured rather than looked.

| What | Measured |
|---|---|
| The import | 4,316 cards, 2,547 rows resolved, 1 unresolved, 2,471 rows in the binder |
| The set menu and the type menu | 126 sets, and 8 types with their counts, from the summary |
| Search "sol ring" | 7 of 2,471 rows, one request, the first tile Sol Ring |
| Filters chained: red, then Creature, then 4 or more | 498, then 295, then 15 of 2,471 rows |
| Sort by price | $92.74 first, and the six dearest in order |
| Scroll over 120 frames | median 16.4 ms, p95 19.8 ms, 47 frames over 16.7 ms, with the RPC reroute in the path |
| Art calls over those 120 frames | 0 |
| Scroll to the foot | 12 page calls, 24 art calls, the last virtual row 801 |
| The upload dialog with a collection active | two choices, and the name field under "Add" alone |
| The diff of the synthetic file over the export | 2,602 rows added, 2,438 removed, 21 changed |
| Replace | the request named the collection, and the list still held one, at 6,605 cards |
| The old file uploaded as new after the replace | two collections, and the replaced one kept its 6,605 cards |
| The budget question | "No budget", and "You decide" on the other two |
| Console errors, horizontal overflow | none |

CAUTION: the frame numbers come from headless Chromium against the dev server, with every RPC rerouted through Node. They say the grid is in the right range. They are not the gate line, which is the owner's own measurement on a production build.

CAUTION: `Repo.Get` returns a collection the caller owns. A page is a slice of the entry list, taken in place. A repo that shares one object across calls hands the next reader a collection the last page truncated. The test fake answers a clone.

CAUTION: the active collection is a choice of one visit, and the store keeps only the session id (D-345). A Playwright run can not seed it through localStorage. The script clicks the collection, as a reader does.

## PR-19, the chat and build experience (2026-09-02)

It merged as #55 and #56 on 2026-09-02. The decisions are D-432 to D-458.

CAUTION: the branch built a start form at `/build` first (D-432, D-434). The owner read it and refused it the same day: the app is chat, and the pool picker is the one control outside it (D-436). The form, its `GetCatalog` rows, `questions/form.go`, and the six form conversations of the gate left. Do not bring a form back.

- The landing is a new chat at `/session/new`, as D-334 said. The unfinished chats sit under its message box, only when there is one, with resume, rename, and delete (D-433, D-438). `ListSessions`, `UpdateSession`, and `DeleteSession` serve it. The store's `List` inflates each session, because a session is a few kilobytes and the summary needs its first message.
- A deck tile of the library carries a delete, behind the same question as the deck screen (D-439).
- Every clickable control shows the pointer cursor, through one base rule in `index.css` (D-440). Tailwind 4 dropped the pointer from buttons.
- The wordmark's shimmer no longer runs forever, and a dialog overlay fades with no backdrop filter (D-441). The page idled at 18 ms a frame under the endless shimmer, and the delete dialogs stuttered on top of it.

CAUTION: measure the idle page before a dialog. The delete dialogs read as the fault, and the fault was a four-second animation on the wordmark that never stopped. `page.addStyleTag` with one style off at a time found it in one run.
- The gap over the message box of a new chat equals the gap under the top bar (D-442). That is 56 pixels on a desktop and 32 on a phone.
- A single card option never zooms, and a half of a commander pair zooms by two to a single card's size (D-443).
- The card art of an option picks it, as the name button does (D-444). The box over the art takes the click, and the image stays readable.
- A commander tile lifts under the pointer with the gold light of a deck tile (D-445).
- After the merge of #55, the owner read the app once more. A deck and its chat are one thing, and a delete of either takes both (D-456). An earlier upload sits on one line (D-457), and the collection page holds two cards of one frame with a compact empty state (D-458). These merged with the tile fixes as #56.
- The owner read the PR on 2026-09-02: the stepper works, and the gate line of PR-19 holds. Two tile fixes followed. A deck's count includes the commander, so a Commander deck reads 100 (D-454). The favorite star sits with the delete button at the right edge (D-455).
- A land swap is a counted change (F-31, D-448). Session `vAvg4eteJhmuPEuJwBul` asked for better lands in place of the basics, answered "a mix", and got one Plains moved to one Island. The brief holds `swap_basics` and `land_kinds`. `generate.FitSwapBasics` fits the count to the base deck and the pool. `CheckRevision` blocks a deck that holds fewer new nonbasic lands. The generate prompt is at version 11. The revise gate answers its own questions now (`answer` in `prompts.json`) and holds a clear land-swap row, revision 9. The next run has 12 turns, not 8, and it is due before the merge.
- The land bucket of the shortlist splits, half mana and half theme (F-32, D-450). Revise gate run 5 played the land ask on the Karlov deck, and the shortlist offered 40 lands that gain life and no untapped dual. `capLands` takes the mana half by `Candidate.Fix`, capped at two, then `Pop`, and the theme half by `Themed`. Probe both decks with a throwaway `cmd` before you touch it: Karlov must show Godless Shrine and Isolated Chapel, and Éowyn the three shock lands.
- A turn stores the brief it acted on in `Turn.revision_brief` (D-449). Read it before you guess what the revise role wrote.
- The session spend holds the build (D-447). A built session showed $0.0023 for 9 calls, and the build's own calls never reached the total.
- A stored question closes with its slot, and the turn that builds a deck stales the cached session first (D-446). The docked chat of a fresh deck drew the answered commander question with its art before.
- The commander offer serves an empty theme and reads the set limit (D-437). Session `DrNPxSaisYj2QVlUkTvB` declined the theme under a Hobbit limit and got a bare pick row, then a commander the build chose. `CandidateHints.SetCodes` and `UseSets` carry the sets, and the hint no longer refuses an empty theme.
- `ChatResponse.phase` streams reading, shortlist, building, checking, repairing, and done (D-435). The generator reports its three through `generate.Request.OnPhase`. The stepper in the working row lights them.
- A failed turn offers "Try again" on a retryable failure and "Reload the session" once a session exists. The reload leaves the live panel and reads the stored session again.

A session read the branch in a real browser with Playwright, against the owner's running stack. A canned stream answered the paid Chat call, and every other RPC reached the API. It measured rather than looked.

| What | Measured |
|---|---|
| The landing of a signed-in reader | `/session/new`, and the Build entry of the top bar points there |
| A form on the page | none |
| The unfinished chats under the message box | two real rows, and a rename that stuck |
| A retryable failure | the reason with its code, "Try again", and "Reload the session" |
| Try again | a third request with the same answers as the failed one |
| The pointer cursor | pointer on a nav link, an enabled button, a select, the drop zone, and the dialog buttons, and the arrow on a disabled button |
| The delete dialog, open, before and after D-441 | 8 to 12 frames of 40 over budget, then 6, and the p95 from 27 ms to 17 |
| The idle page, before and after D-441 | 33 frames of 40 over budget, then 10 |
| The gap under the top bar and over the message box | 56 and 56 pixels on a desktop, 32 and 32 on a phone |
| Console errors, horizontal overflow | none |

CAUTION: React mounts a component twice in development, and the cleanup of the first mount aborts a send that already left. A send that starts from an effect must wait one tick and cancel on cleanup, so a double mount sends once. The form's first send hit it before the form left, and no jsdom test finds it: the test renderer mounts once.

CAUTION: a Connect stream request carries a 5-byte envelope before its JSON. A Playwright route that reads the body must skip it.

CAUTION: the owner's Firestore emulator refused every transaction on 2026-09-02, and an untouched sessions test timed out after 60 seconds. A private emulator on port 8282, through `firebase emulators:exec` with a scratch config, ran every store test green. Restart the owner's emulator before `make store-check`.

## PR-14A, the bracket profile, built (2026-09-02)

The branch holds the whole of PR-14A but its gate run. Read `docs/reference/bracket-profile-2026-09-02.md` for every source, and D-459 to D-464 for the calls.

- `go/internal/spellbook` calls `estimate-bracket` at 90 requests a minute with a named agent (D-459). The limiter spaces calls with no burst, and a 429 retries once.
- `go/internal/profile` is the library. `bands.json` holds a band per feature per bracket, and per power step for a 60-card deck. `profile.go` measures 16 features, `karsten.go` holds the two source tables, `goldfish.go` deals 10,000 hands, and the content check reads the endpoint.
- `go/internal/rules/brackets.json` gained `mass_land_denial`, `max_extra_turn_cards`, and `max_combo_speed` per bracket. The prose note of F-11 left `checkBracket`.
- `candidates.Build` drops mass land denial through bracket 3 and extra turns at bracket 1, by the tags (D-462). `themes.json` names the two slugs under `roles`. `generate.Builder.cutShortlist` then sends the pool to the endpoint and drops what it flags, the commanders and the locked cards excepted (D-468). Every build makes two endpoint calls now.
- `generate.TargetsFor` reads the band midpoints for Commander. The prompt carries a deck shape block, and prompt version 12 reads a profile finding in the repair turn. `MaxRepairs` is 2, and the second pass runs for a profile finding alone (D-461).
- `Deck.profile` is field 23, additive. `make proto` regenerated the TypeScript, and no screen reads it yet: PR-20 shows it.
- `cmd/bracket-gate` and `make bracket-gate` are the gate, 15 prompts, with `-dry`, `-only`, and `-no-judge`. The dry run resolves every commander and builds every shortlist for free, and it ran clean on 2026-09-02.
- `generate.JudgeBracket` asks the judge role for the bracket of a card list. The judge never sees the bracket of the build.

A live smoke read ran on 2026-09-02 with the real snapshot and the real endpoint. It read the Karlov deck of deck gate run 11 at bracket 3. Every reader answered. The deck holds 36 lands and 0 tapped lands, and its sources sit at 0.95 of the Karsten requirement. The average mana value is 3.43, the mana on turn four 4.31, and the endpoint tag E. The role counts read 0 there, because the scratch list carried no roles.

CAUTION: the bands are first values (D-463). The first bracket gate run will find off-band features, and that is the point of the run. Read its off-band table, move a band with a decision id, and never widen a band to make the verdict pass.

CAUTION: every build now makes one call to Commander Spellbook after the engine check. A failed call leaves an `content_unchecked` info finding and no content finding, so the build never waits on the endpoint's health.

The owner answered OQ-52 and OQ-53 on 2026-09-02 (D-467 to D-469). The session calls stand. Two things changed. The endpoint reads the shortlist before the build, and the builder drops what it flags, with an info finding `shortlist_cut` on the deck (D-468). The engine's Commander land range is 27 to 41 (D-469).

## PR-14B, the deck quality model, built (2026-09-02)

The branch holds the whole of PR-14B but its gate run and the live reads. Read `docs/reference/deck-quality-model-2026-09-02.md` for every source fact and every call, and D-470 to D-478 for the session calls.

- `go/internal/meta` reads the five sources: MTGO, MTGJSON, EDHREC, the cEDH database, and the Topdeck.gg API. MTGO gives the month and event pages, and EDHREC the commander and average deck pages. Each reader is a parser over one page, and `testdata/` holds trimmed real pages of 2026-09-02. `refresh.go` is the job: every source on its own, a failed source logs, and the report counts pages and parse failures per source (M-6).
- The store is `meta/` in the card bucket: `raw/<source>/<key>.gz`, `lists/<format>/<YYYY-MM>.jsonl.gz`, `precons/<MTGJSON version>/precons.jsonl.gz`, `commanders/<day>.jsonl.gz`, and `model/<version>/quality.json.gz` with a `complete` marker. On the local stack it is the parent folder of `CARDS_SNAPSHOT_DIR`, so `.local/gcs/mtg-local-cards/meta/`.
- `go/internal/quality` is the model. `resolve.go` reads a list against the index, and `features.go` measures it through the profile of PR-14A. `synthetic.go` breaks one axis per real list. `fit.go` fits a proportional odds scorer per format, and `score.go` grades a deck. `refit.go` reads the store, fits, and writes a model version.
- The grade lands in five places. They are `Deck.quality` (field 24), the summary's last sentence, the revision note on a drop, the shortlist's `MetaBoost`, and the bracket 4 or 5 commander offer. `Card.quality` (field 34) rides on `GetCards` alone. The generate prompt gained a "Format shape" block.
- `worker -meta` is the job, and `make meta-refresh` runs it on the local stack over the network, for free. `cmd/quality-gate` and `make quality-gate` write the gate document for free, and `-write` stores the fitted model.
- `cmd/api` loads the newest model beside the snapshot and polls for a newer one on the snapshot cadence. No model grades nothing, and every path runs as before.

A short live run of `make meta-refresh` ran on 2026-09-02, with `-meta-months 1 -meta-pages 5`, in ten minutes. MTGO read 5 pages into 133 lists, EDHREC read 13 pages into 5 lists, and the database read its page, all with zero parse failures. MTGJSON read 557 deck files, and then one connection reset ended the source before it wrote the table. The fetcher retries a transport error once since. The raw files are in the store, so the next run reads the rest and writes the table.

The fit ran over the lists it had and stored model `20260902T194142Z` on the local stack. It had no baseline tier, so its pair bars read 0.

A second short run read the last 146 MTGJSON files and stored the table, 701 products. One MTGO page timed out on both tries, and the job counts a fetch error and reads on since (the first run ended the source on it). The fit then read 194 Commander, 468 Standard, and 539 Modern lists. It found 60-card products from 1996 in the baseline, so a product older than the format's pool serves no baseline now (D-478).

`docs/reference/pr14b-quality-gate-smoke-2026-09-02.md` is the free gate over that store: a smoke read of seven MTGO pages, and not run 1. It reads FAIL. Modern puts every great list over the precons (183 pairs) and the precons over the bad lists in 0.65 of 5,185 pairs. Commander has no great tier without the Topdeck.gg key. The bracket 5 offer named Niv-Mizzet, Etali, and Urza at 0.50 each, from the database tier alone.

The owner set the windows on 2026-09-02 (D-479, D-480), and the full run went. It read 200 MTGO pages into 5,288 lists, 21,146 Topdeck.gg lists over 90 days, and 403 EDHREC pages. It exposed four defects, and each one got its fix the same day:

- The MTGO site answers a 302 to the month listing for a retired event page, and the client followed it.
- The Topdeck.gg bulk endpoint answers 429 twice in a row.
- The EDHREC read hit the page cap.
- A month of cEDH lists inflated past the 16 MB limit of `gzstore`.
 `docs/reference/deck-quality-model-2026-09-02.md` holds the read. The fit keeps 4,000 lists per tier since (D-483).

Gate runs 1 to 8 ran on the full store the same day, free. Run 1 read the precons over their broken copies in 0.16 to 0.46 of the pairs. The owner moved the bad rung under the precons alone (D-484). A defect detector joined the ladder over runs 2 to 8 (D-485).

Run 8 reads 0.96, 1.00, and 1.00 on the top-list bar for Commander, Standard, and Modern, against 0.90. It reads 0.94, 1.00, and 0.95 on the precon bar, against 0.95. The stored model is the run 8 fit, and the API and every gate load it.

The owner kept the bar at 0.95 and merges PR-14B with run 8 on record (D-486). Runs 9 and 10 followed the judge lane: the breaks draw from seen cards and the rungs weigh the same (D-488). Run 10 reads 0.97 to 1.00 on the top-list bar and 0.88, 1.00, and 0.95 on the precon bar. The stored model is run 10.

Deck gate 13b passed 24 of 24 ($2.62). The tier judge lane over it read 5 of 24 with the run 8 model and 4 of 24 with run 9 (`pr14b-quality-judge-run1.md`, `-run2.md`, $0.31 each). The model grades most built decks below the precon baseline, and the judge reads them as typical. `go run ./cmd/quality-gate -explain <deck gate document>` prints the detector probability and the six largest contributions per deck, free. It named the cause: the ladder holds 4,000 tournament lists per tier against 195 community decks, and none for the 60-card formats (D-488).

CAUTION: the judge bar is open on the corpus, not on the judge. The owner widened the EDHREC read to 1,100 commanders (D-489). PR-14C gains the user decks of Aetherhub and MTGGoldfish for the 60-card typical rung (D-490). Read the explain output before you touch a weight.

Deck gate run 13 (`docs/reference/pr8-deck-gate-run13.md`, $2.64, 66 calls) read FAIL: 22 of 24 decks passed the block checks, and 6 invented names reached the user. The six were curly apostrophes the model wrote, "Commander’s Sphere", which the name match did not fold. Fifteen `summary_rules_claim` warnings came from the lint over the grade sentence the code appends. The fold sits in `cards.normName` and `candidates.FoldName` since, and the sentence goes on after the lint. The two blocks were a deck six cards short of those misses and a model slip, two copies of Mirkwood Nurturer.

The grade's sentence names its frame since, "against the top lists of the format" (D-487). `cmd/quality-gate -judge <deck gate document>` is the tier judge lane, `make quality-judge` runs it, and the deck gate document carries a `Commander:` line and a `Grade:` line since.

CAUTION: runs 7 and 8 sit a hundredth under the precon bar in Commander and at it in Modern. The per-axis table in `docs/reference/pr14b-quality-gate-run8.md` names the weak axes, curve and synergy in Commander and colors in Modern. No bar moved. Read that table before you touch the model, and change no bar.

CAUTION: the first `make meta-refresh` takes about 40 minutes. It reads 200 MTGO pages at one a second, about 720 MTGJSON deck files, and about 300 EDHREC commanders with their average decks. `META_ARGS="-meta-months 3 -meta-pages 50"` makes a short first run.

CAUTION: Moxfield answers 403 from Cloudflare to a plain client on the deck page, the v2 API, and the v3 API. PR-14B reads no Moxfield list, and the database gives the tier and the commander alone (D-470). Do not spoof a browser to get past it.

CAUTION: the Topdeck.gg reader follows the docs alone. No session has read a live answer, because no key exists (OQ-54). The `deckObj` shape is the documented sketch, and the text list is the fallback. Read the first live answer before you trust the counts.

CAUTION: the quality gate has three bars in the code and one outside it. The pair bars and the bracket 5 offer bar are in `cmd/quality-gate`. The judge bar over the golden decks reads the next deck gate run, whose summaries carry the tier. That run costs about $2.24, so ask the owner first.

## The paid runs of 2026-09-03, after PR-21

The owner opened the paid runs once PR-21 merged, in the order of the eval list. A result that does not make sense, or is not good, stops the sweep. Three ran, $0.60 together.

- Question gate run 33, `docs/reference/pr7-question-gate-run33.md`, $0.18, 20 minutes. PASS: 27 of 27 counted conversations catalog-only, 119 questions, 98 closed a slot, no invented question, no dead end. Run 32 asked 114 and closed 90. The differences sit in 19 of 107 conversations, one question each, both ways. The one precon conversation still takes the upgrade path, and no exclusion sentence appears, so classifier version 17 reads "upgrade my precon" as before (D-496).
- Question eval run 33, `docs/reference/pr7-question-eval-run33.md`, $0.10, 13 minutes. The bad-question ratio reads 10.3 percent on a holdout of 117, against 8.8 percent on 125 in run 32. Both runs hold the same 22 unwarranted questions. `tune-check` accepts it inside the margin of nine, and no counter fell. The colors row leads the bad questions in both runs.
- Tier judge run 3, `docs/reference/pr14b-quality-judge-run3.md`, $0.32, three minutes, with the run 12 model. FAIL on the bar: the judge agreed on 8 of 24, against 4 of 24 in run 2, with 5 off by one rung. The model grades five decks typical now: the blink deck, the Modern tempo deck, the two precon upgrades, and the two-family set deck. The judge agrees on all five. The other 19 the model grades bad, and the judge reads them as typical, baseline, or good, the corpus finding of D-488. The judge itself moved on four decks between the runs.

The owner chose the deck gate next and left the bracket rejudge for later.

- Deck gate run 14, `docs/reference/pr8-deck-gate-run14.md`, $2.54, 35 minutes, 62 calls. FAIL on one deck of 24: the Goblin Storm precon upgrade of prompt 17 came back with no cards and a `deck_size` block. No invented name reached the user, and no summary stated a false rule. The other 23 decks passed every block check, with 9 repair turns against 13 in run 13b, and the six set prompts read as before. The grade line reads 18 bad, 2 baseline, and 4 typical under the run 12 model.
- The cause is a defect and not the model's variance alone (D-509). The bracket cut of D-468 drops a forbidden shortlist card and keeps only the commanders and the locked cards. So precon cards leave the pool too. The share rule of D-218 still asked for 67 of 78 names. The repair turn read fewer marked names, gave up, and answered no cards. The builder then replaced the legal first deck with the empty one. The two fixes merged as #63 on 2026-09-04. The share counts the precon names the pool holds. A repair with a block or a miss after a clean pass leaves the clean pass in place, with the note `repair_kept_earlier`. `generate.TestPreconShareCountsThePoolAlone` and `generate.TestBuildKeepsTheLegalDeckWhenTheRepairFails` pin both.
- Deck gate run 14b, `docs/reference/pr8-deck-gate-run14b.md`, $0.13, two minutes, reran prompt 17 alone with the fix. PASS: 99 cards and no block. The repair turn ran on the two profile findings and the combo, and it kept the share. The model grades the deck typical. Run 14 with 14b reads 24 of 24, as runs 12 and 13 did with their reruns.
- Revise gate run 8, `docs/reference/pr12b-revise-gate-run8.md`, $1.23, 16 minutes, 33 calls. PASS: 11 of 11 turns met their bar. The revised decks kept 97 to 100 percent of their base, against 96 to 100 in run 7. The cost rose from $0.74 on 25 calls. Five turns bought a repair turn against two in run 7, four of them for a profile finding of PR-14A. On revision 9, the land swap of the Karlov deck, the repair answered a worse deck. The deck before it stood, with the note `repair_kept_earlier` (D-509), and the turn passed with 100 percent kept.

The sweep is complete. Every row of the eval list ran once, for $4.60 together, and one defect came out of it, F-34, fixed and merged as #63.

## PR-15, what it holds (2026-09-04, merged as #65)

PR-15 holds every slice of the plan in the section after this one, merged as #65. The tree is green: Go build, vet, the `-race` suite, golangci-lint, and `make ste-check`.

- `internal/evalrun` (slice 1): `Header`, `Row`, and `Run`, the JSONL writer and reader, `Compare`, and `Merge`. A header holds the suite, the run id, the date, the commit, and the resolved model and effort per role. It holds the snapshot date, the prompt versions, the stored versions, the cost, and the verdict. A row holds one measurement of one item, as a gate row or an information row. The header names the metrics where lower is better, so a committed file describes its own senses.
- Every gate writes its run file beside its document, and the "## Run" block of every document is one fingerprint (`Run.Markdown`). The writers: `deck-gate`, `bracket-gate` and its `-rejudge` lane, `revise-gate`, `questions-gate`, `questions-eval`, `quality-gate`, and its `-judge` lane. Each one takes `-run-out`, and it refuses an existing file before the first provider call (D-65). The Makefile names the file after the document, under `docs/reference/eval/`: `DECK_GATE_RUN`, `BRACKET_GATE_RUN`, `REVISE_GATE_RUN`, `GATE_RUN`, `EVAL_ROWS`, `QUALITY_GATE_RUN`, and `QUALITY_JUDGE_RUN`.
- `cmd/eval` (slice 2), four free modes. `compare` reads two runs of one suite and names the flips. `check` reads every suite of `baselines.json` against its newest run. `baseline` records the accepted run files of a suite, a run with its rerun as one. `import` reads a deck gate document from before the run files existed.
- The verdict of a compare. A gate row that moved against its sense past the margin is a regression, and the verdict fails. A suite with no gate row, or no verdict of its own, reads NOT EVALUATED and never PASS. The information rows fold into one count, and `-info` lists them, because the cost rows move with the daily prices on every run.
- `docs/reference/eval/` holds the imported runs 13b, 14, and 14b, and `baselines.json` names 14 with 14b as the decks baseline. The compare of 13b against 14 names deck 17, `blocks` 0 to 1 with `deck_size`, as the one gate flip, which is F-34. That is the gate line of slice 2.
- Tier 0 (slice 3): `make eval-check` is free. The job `verify:eval` runs it on a pull request that touches the run files, the harness code, the Makefile, or the workflow. It takes about one minute of runner time.
- The plan judge (slice 5): `generate.JudgePlan` grades a built deck on four fields, each no, partly, or yes, with one sentence of reason per field. The fields are `plan_coherent`, `theme_fit`, `useful_as_built`, and `summary_honest`, the proposal of OQ-58, at `PlanRubricVersion` 1. The deck gate calls it after the F-26 judge, unless `-no-judge`. Its rows are information, `plan_<field>` and `plan_score`, and a failure of the lane is a row and never a verdict. The document prints one `- PLAN field=grade: why` line per field, and the summary table counts the decks read and the mean score.
- The golden expectations (slice 4): a conversation of `conversations.json` can carry `expect`, the values its slots must end with. The gate renders the settled slots as words (`slotValues`) and reads each expected key as a row, `slot_<key>`. A counted conversation with a miss fails the gate, the fourth bar, and the document names every miss. The words: `format` (commander, standard, modern, house), `colors` in WUBRG order, `power` (bracket N, casual, fnm, tournament), `pool_rule`, `commander` (a name, `*` for any pick, or `delegated` for a skipped slot), `budget` (a number, with `to buy` or `whole deck` when the scope matters), `theme` (words inside the slot text), `locked`, and `sets`. The 27 counted conversations carry expectations, written from their messages on 2026-09-04. The three that start after a build carry none. A probe can carry one, and its rows are information: probe 109 does (D-525).
- The sweep (slice 6): `go run ./cmd/eval sweep -cap <USD>` runs the paid suites through their Makefile targets, in the order of the eval list. The steps: questions, question-eval, decks, tier-judge, and revise. It numbers each document after the highest run of its family, and it names the run file beside it. `-dry` prints the plan with an estimate per step, the cost of the suite's last run file. It needs `EVAL_SWEEP=1`, each target keeps its own guard, and it stops before a step that crosses the cap. A FAIL stops it, as the owner's rule of 2026-09-03 says, and `-continue` goes on. `-suites` picks the steps, and a step whose input step did not run reads the newest document of that family.

CAUTION: question gate run 34 read the 27 expectations through a renderer that read the wrong fields, and named 21 misses. Run 35 read them through the fixed renderer, and every one met. A miss on a later run is a wrong expectation or a wrong slot, and the session decides which before it changes either.

CAUTION: D-427 names 14 terse conversations, and the file holds 47 with "terse:" in the name. The owner ruled on 2026-09-04 that all 47 join the bar at 83 percent (D-522), and the change waits for its branch. Until it lands every terse conversation stays a probe, and its expectation rows are information.

CAUTION: the plan judge costs one judge call a deck, about $0.02 on Opus 5. Deck gate run 16 cost $3.79 for 25 prompts with it, against $2.54 for run 14, and 12 decks needed the repair turn. The sweep's estimate reads the last run file, so set the cap from run 16 and not from run 14. `-no-judge` skips both judge lanes together.

CAUTION: an imported run names no model. Deck gate run 16 and tier judge run 4 carry the first complete fingerprints. The questions suite has no run file yet, because run 34 failed before the gate wrote it and the fix came after. The revise suite has none, because its run waits (OQ-63).

CAUTION: `check` compares the baseline with the newest run of its suite that is newer than every run of the baseline. A run older than the baseline is history, and the check never reads it. Newest reads the date first, then the number the run id ends with.

What comes next: the work of D-519 to D-525, and the first compare of each suite on its next paid run. Slice 3 built the compare lane of Tier 0 without the trimmed snapshot of OQ-60. Slice 6 built the sweep for the owner's machine, the recommendation of OQ-57.

## The PR-15 paid gate, run (2026-09-04, D-514, D-515)

The owner confirmed the go, and the three steps of D-513 ran on 2026-09-04 for $4.53. `docs/reference/pr15-paid-gate-2026-09-04.md` holds the whole read. The short form:

- Step 2, the question gate and its eval, $0.28. Run 34 reads FAIL on the fourth bar with 21 misses over 15 conversations, and the three old bars pass. 17 misses came from the renderer. It read the commander and the locked cards from the id fields of the proto, and the classifier fills the name fields of the state. A reader with no collection read no pool rule, and the build reads any-card for them. The renderer reads the names and the any-card default now. 3 misses were expectations that named a fourth message the net of D-351 never sends, and they changed. 1 miss is a wrong slot: the classifier replaced `infect` with `the best deck under budget` (F-36). The gate wrote no run file for run 34, because it returned the FAIL before the write, and `writeRunThen` fixes the order.
- Step 7, the deck gate on prompt 25, $0.11. The plan judge schema went through the Anthropic API with no error, and the run block names every role. The build answered the Avengers request with a Syr Konrad list and no superhero identity, and the summary hid it (F-37).
- Step 9, $4.15. Deck gate run 16 passes 25 of 25, and the compare with run 14 with 14b names no gate flip. 76 information rows moved, 41 worse. Run 16 is the decks baseline. Tier judge run 4 reads 7 of 25 on the open bar, and the sweep stopped there. The revise gate did not run.
- The plan judge wrote "placeholder" as its reason on `summary_honest` in 10 of 25 decks (F-38). It graded legality from stale card knowledge on decks 2 and 10 (F-39). The schema asks for the reason first, the document counts the empty reasons, the instructions say that code checked the legality, and `PlanRubricVersion` reads 2.
- `eval check` reads NOT EVALUATED on a baseline with no newer run, and never PASS.

The two runs that complete the gate ran on the owner's word (D-516). Question gate run 35 reads PASS, 27 of 27 catalog-only, and every one of the 27 expectations met. Its eval reads 7.3 percent bad on the holdout. Revise gate run 9 reads PASS, 11 of 11 turns, $1.07, and its run block names generate, repair, and revise on `gpt-5.6-terra`. Runs 35 and 9 are the baselines of their suites. The commands, for the record:

1. Question gate run 35 with its eval, about $0.28 (OQ-62):

   ```
   EVAL_SWEEP=1 go -C go run ./cmd/eval sweep -cap 0.50 -suites questions,question-eval -continue
   ```

2. Revise gate run 9, about $1.23 (OQ-63):

   ```
   EVAL_SWEEP=1 go -C go run ./cmd/eval sweep -cap 1.50 -suites revise
   ```

The owner committed the documents and the run files, reset `main` to f299289, and merged #65 on 2026-09-04 (D-517). The reset, for the record:

   ```
   git branch -f main f299289
   git push --force-with-lease=main:dbbcb4e origin main
   ```

CAUTION: a miss of run 35 on a commander name is a name the classifier wrote in other words than the expectation. Read the turn line before you change either.

CAUTION: the sweep stops on the tier judge, which reads FAIL on the open bar of D-488 until the corpus moves. A full sweep needs `-continue` past it, and the cap must hold the plan judge: read the estimate of the deck gate from run 16.

## PR-15, the eval harness, prepared (2026-09-04, D-511)

The owner moved PR-15 ahead of PR-22 and PR-23 on 2026-09-04 (D-511). It merged as #65 on 2026-09-04, from `main` at f299289 (#64). This section holds what a session found on 2026-09-04, the plan in slices, and the four owner questions. The roadmap entry is the authority on the goal, and D-39, D-423, D-427, D-428, and D-430 bind it.

What exists. Eight commands write a gate document, and each one defines its own result type in `package main`. Only `internal/tune` holds a shared shape, `Verdict` and `Summary`, and only `questions-eval` and `tune-check` read it. The fingerprint differs per document. The run date is in every one, the snapshot date in four of eight, the prompt version in three, and the model name in two. `deck-gate`, `bracket-gate`, and `revise-gate` print no model name.

The overwrite guard of D-65 is one shell line, repeated seven times in the Makefile, and `gatekit.RefuseExisting` has one caller. `tune-check` pairs two question runs by conversation and question, and it reads a noise margin (D-230, D-258). CI runs no gate, no probe, and no eval, and it sets no `CARDS_SNAPSHOT_DIR`. The free coverage is `go test -race`, which holds the report tests of each gate, and `make llm-defaults-check`, which warns and never fails. The paid sweep of 2026-09-03 cost $4.60 and about 84 minutes of provider time. A nightly sweep in Actions costs about $140 a month and about 42 hours of runner time, so the "Tier 1 nightly" line is a question (OQ-57).

The plan, in slices. Each slice is one concern, with a free gate.

1. `internal/evalrun`: one `Run` header and long-format `Row` records. The header holds the suite, the run id, the date, and the resolved model and effort per role. It holds the snapshot date, the prompt versions, the quality model version, the precon table version, and the commit. A row holds the item, the metric, the value, the kind (gate, info, or lenient), and a detail. Every gate writes its document as today, a JSONL file beside it under `docs/reference/eval/`, and the same "## Run" fingerprint block. Gate: every writer's test reads the block back.
2. `cmd/eval compare`: two runs of one suite, paired by item and metric, with the named flips, the margins, and one verdict. A suite with no gate metric reads "not evaluated", never PASS. `docs/reference/eval/baselines.json` names the accepted run per suite, and `tune-check` folds into it. Gate: the compare over deck gate runs 13b and 14b names F-34 as the one flip.
3. Tier 0 in CI, $0: `eval compare` over the committed baselines and the newest committed run per suite. A fingerprint check replaces `llm-defaults-check`. About one minute of Actions. A trimmed card snapshot in the repo adds the free dry runs of the deck gate. OQ-60 asks the owner about its size.
4. The golden expectations. Each conversation of `conversations.json` gains `expect`: the slot values at the end. Each deck gate prompt gains `expect` as well. It names the commander, the count, the pool rule, the locked cards, the sets, and the excluded product. The gate reads them as rows, so a wrong slot is a named flip and not a passed count. The 14 terse conversations join the bar here (D-427), and every number rebases at once.
5. The plan judge. One judge call per built deck on Opus 5 (D-430), with a rubric of four fields on three-point scales. D-66 shaped the M-5 rubric the same way. The rows are `info` until a baseline exists. A bar reads the number and never the prose (lesson 11). OQ-58 asks the owner to confirm the fields.
6. `cmd/eval sweep`: the paid suites in the order of the eval list, under a cost cap per run (D-4 shape). Each suite keeps its spend guard. It runs on the owner's machine, on the owner's word, as every paid target does today (OQ-57).

What PR-15 does not settle. D-423 sets the fit threshold from the M-5 scores, and the sheet holds 30 scored items of 60 on 2026-09-04. D-66 asks for 50, and the owner scores all 30 by phone in question and answer form (D-524). The bake-off of D-39 is a paid run the harness makes possible, and it is not part of PR-15.

The free gate of PR-15: every slice's tests, `eval compare` over the committed runs, and Tier 0 green in CI. The paid gate: one sweep through the new command, about $4.60, that reproduces the verdicts of 2026-09-03 and names no flip but F-34.

## The corpus items of 2026-09-04 (D-510)

PR #64 merged the two items the hand-off of 2026-09-04 named, with the F-34 record of #63. The tree is green: Go build, vet, the `-race` suite, golangci-lint, and `make ste-check`.

The MTGJSON skip. The version stamp of the deck list carries the build day, `5.3.0+20260903`, so the table of yesterday never matched today's stamp. The job fetched 702 deck files every day. `meta.SameProducts` compares the kept products of two deck lists: file name, set code, name, release date, and type. The run reads the stored deck list of the newest table, and it reads no deck file when the products match. The report then names the stored table under `Skipped`.

A new product reads every file again under the new stamp. A stored table older than `meta.PreconsMaxAge`, 30 days, reads whole again. MTGJSON corrects a deck file now and then with no change to the deck list. The two stored deck lists of 2026-09-02 and 2026-09-03 hold the same 3,029 entries, so the skip holds on the next run. `TestMTGJSONSkipsAnUnchangedDeckList` pins the four cases.

Deck gate prompt 25. The prompt asks for an Avengers superheroes Commander deck from the library first, with no card of the Avengers Assemble precon, and a delegated commander. Two prompt fields are new. `collection_file` names a ManaBox export beside the file of the `-collection` flag, and `exclude_precons` names the products. The gate resolves each name through `precons.Table`, and it checks that the binder holds one product of the name whole. Then it runs `precons.Exclude` before both pools, as the chat does.

The document gains the block "The precon exclusion (PR-24)": the products, the excluded cards, the spare cards, the excluded cards in the deck, and the blocks. `DECK_GATE_ARGS="-only 25"` runs it alone, for about $0.13, the cost of run 14b. The dry run of 2026-09-04 reads a pool of 152, 57 cards excluded, and 27 cards with a spare copy. The 27 are reprints the binder holds from other sets, Arcane Denial and Talisman of Conviction among them, and D-408 keeps them usable. The report column names them, so a reader can tell a spare copy from a card that slipped through.

CAUTION: the fixture `manabox_collection_avengers.csv` is `manabox_collection.csv`, the export of 2026-08-24, byte for byte, plus three rows of basic lands. The file the session of 2026-09-03 left held 91 generated rows, with ManaBox ids from 990001 and the stamp `2026-09-04T00:00:00.000Z`. Only three of them filled a gap of the export. The other 88 gave every card of the product a second copy, so the exclusion of D-408 excluded nothing. The first dry run read 0 cards excluded. A fixture for the exclusion must hold the product once.

CAUTION: the export of 2026-08-24 does not hold Avengers Assemble whole (F-35). Its ManaBox deck binder of 2026-06-27 lacks three basic-land printings, six cards: Plains 288, Island 290, and Mountain 294 of MSH. Under D-408 that binder owns no precon, and "not from my precons" excludes nothing for it. The check counts basic lands, and the exclusion never removes one (D-37). The owner ruled on 2026-09-04 that the check ignores basic lands (D-523), and the fix waits for its branch.

## PR-21, the share link and the print view, merged (2026-09-03, #62)

PR-21 merged as #62 on 2026-09-03. The call is D-508, and the free gate is `docs/reference/pr21-gate-2026-09-03.md`.

- `DeckService.ShareDeck` makes a 32-byte token, shown once, and `RevokeShare` ends the link. `GetSharedDeck` and `ExportSharedDeck` need no sign-in. The store keeps the hex SHA-256 of the token at `shares/<hash>`, one document with the owner and the deck id. The deck document keeps the hash too. A second share replaces the link, and a rename keeps it.
- `SharedDeck` and `SharedCard` are the public message. They hold the name, the format, the power, the summary, and the cards by role with the card data inline, and no user field. `decksvc.TestSharedDeckHoldsNoUserField` reads the proto text and proves it (guardrail 13).
- `internal/ratelimit` is a fixed window of 60 calls a minute per client address, from the first address of `X-Forwarded-For` and else the remote host. `cmd/api` mounts it with the auth interceptor's new public list on the deck service, for the two public procedures alone.
- The web: a Share dialog on the deck page with the link shown once, a Copy button, and Revoke. The public page at `/d/<token>` sits outside the sign-in guard and pulls in no auth module. It shows the deck with its art and an export, and it says when a link opens nothing.
- The print view: the print variant hides the header, the chat, the images, the stats, the filters, the sample hand, and the export. The stylesheet `index.css` drops the dark ground, the shadows, and the frame.

CAUTION: the dialog shows the token once. The owner who closes the dialog without a copy makes a new link, and the old one dies. The deck carries a `shared` mark, so the button reads "Shared" after a reload.

CAUTION: the local stack sends no `X-Forwarded-For`, so the remote host is the key there. Behind Firebase Hosting and Cloud Run the header carries the visitor, which PR-22 proves live.

## PR-20, the deck view and the card detail, merged (2026-09-03, #61)

PR-20 merged as #61 on 2026-09-03. The call is D-507, and the free gate is `docs/reference/pr20-gate-2026-09-03.md`.

- `cards/rulings.go` reads the Scryfall rulings bulk file, the fourth snapshot file. `SnapshotFiles` holds it, so the worker downloads it with the daily snapshot, and `LoadIndex` treats it as optional, as it treats the set file. The index keeps the playable printing ids per Oracle id too.
- `CardService.GetRulings` answers the rulings of one card, oldest first, with the snapshot date and a flag that says whether the snapshot carried the file. `CardService.GetPrintings` answers every playable printing with its price, newest set first. Both are additive (D-507).
- `components/ui/sheet.tsx` is the Radix dialog with side styling. `card-detail.tsx` is the panel. It shows both faces, the Oracle text, the type line, the mana cost, and the reason line. It also shows the legalities, the rulings with their dates, the printings with prices, and "Open on Scryfall". The rulings and the printings load when the sheet opens.
- `deck-view.tsx` gains the filters by role, color, mana value, type, and owned, and the sorts by mana value, name, and price. It shows five stats with a text table each: the curve, the sources, the type counts, the average mana value, and the cards to buy. A card's name is a button that opens the detail.
- `sample-hand.ts` holds the hand rules of D-318 with a seeded shuffle, and `sample-hand-panel.tsx` the panel. The panel draws seven, mulligans to six and to five by the London rule with the reader's bottom choice, and draws one.

CAUTION: the three stored snapshots carry no rulings file. The index logs the missing file once per load, and the panel says "The card data holds no rulings file yet" until the worker stores a new snapshot. `make dev-seed` on a running local stack fetches one, about 110 MB.

CAUTION: the local stack was down on 2026-09-03, so no session saw the sheet or the charts on a screen. The owner reads them before the merge.

## PR-14C, MTGTop8 and the casual 60-card decks, merged (2026-09-03, #60)

PR-14C merged as #60 on 2026-09-03. The calls are D-502 to D-505, and `docs/reference/pr14c-sources-2026-09-03.md` holds every verified fact about the three sites. The owner dropped the Aetherhub lane (D-502), so the slice has two lanes.

- `meta/mtgtop8.go` reads the paper events of Modern, Standard, and cEDH (D-504). It walks the format pages with their later pages. Then it reads the event pages, with the placements and the field, and the text export of each deck. A top-8 finish in a field of 32 or more is great, and the rest good (D-505). The reader stores and skips an event page that names mtgo.com as its source, because the MTGO lane holds the same lists.
- `meta/mtggoldfish.go` reads the user decks of Modern and Standard as the typical rung (D-490, D-503). It walks the listing pages, newest first, then the deck page of each. The deck page embeds the whole list in a form field. The reader never touches the download endpoint the robots file disallows.
- `refresh.go` runs the two lanes after EDHREC, with caps of 300 and 200 requests a run and 100 listing pages a format. The raw pages sit under `raw/mtgtop8/` and `raw/mtggoldfish/`, a page that did not parse under the `-failed` prefix, and `-meta-reparse` covers both lanes (M-6).
- The fixtures under `meta/testdata/` hold trimmed real pages of 2026-09-03, and `TestJobRun` reads both lanes twice and once more in reparse mode.

CAUTION: Aetherhub answers a Cloudflare challenge page to the app's agent on every page, as Moxfield does. The house rule stands: no spoofed browser (D-470). The app asks the site for nothing.

CAUTION: the robots files of MTGGoldfish and Aetherhub carry `Content-Signal: search=yes,ai-train=no,use=reference`. The owner read it and kept the legal check of D-5 (D-503). Do not reopen that call without new facts.

CAUTION: a user deck is what a person uploaded. Some are short, some hold a card the index does not know, and the fit counts those as unusable. Read the `unusable` count of the fit log per format after each refresh.

The first live run read 300 MTGTop8 pages into 259 lists of 35 paper events, and 200 MTGGoldfish pages into 158 lists, with zero failures. The whole MTGGoldfish cap went to Modern, so the cap is per format since (D-506), and the Standard rung fills on the next run. Quality gate run 12 reads FAIL on the Commander and the Modern precon bars, 0.78 and 0.88.

The Modern bar fell from 0.9465 with the typical rung, as the Commander bar fell in run 11. The user decks hold unseen cards and odd playsets, and the model reads those as less of a defect since. The explain mode over deck gate 13b moved the three casual Modern decks out of the detector's flag. No deck left the bad rung. `docs/reference/pr14c-gate-2026-09-03.md` holds the read.

CAUTION: the gate of PR-14C reads FAIL on record, as PR-14A and PR-14B merged. The bars did not move toward a pass with the typical rungs, and the judge bar of PR-14B stays open (D-488). Read the explain output before you touch a weight, and change no bar.

## PR-24, the precon exclusion, merged (2026-09-03, #59)

PR-24 merged as #59 on 2026-09-03, with the corpus-step docs of the same day in the same commit. The calls are D-496 to D-498 and D-500, and the free gate is `docs/reference/pr24-precon-gate-2026-09-03.md`. A reader asks for a deck that uses no card of a precon, by name or as "not from my precons", and the build leaves those cards out.

- `precons.Table` indexes the MTGJSON table of the meta store: 701 products, by key and by name. `Resolve` maps the reader's words onto products. `Owned` lists the products a collection holds whole (D-408), and `Exclude` takes the products' copies off the owned counts (D-500).
- `cmd/api` loads the newest table beside the quality model and swaps it on the snapshot cadence. `agentsvc.WithPreconTable` wires it. No table excludes nothing, and the turn says so.
- The classifier is version 17: `precon_names` and `facts.exclude_precons` (D-496). `applyPrecons` resolves the names through the turn's hints, and "my precons" reads the printing counts of the collection once, on that turn alone.
- `Slots.exclude_precon_keys` (field 14) holds the product keys. The snapshot is version 4. The catalog row `precon_unresolved` asks about a name the table can not settle, and the corpus section 11 holds it.
- The build subtracts the copies per Oracle id, and it passes the excluded ids to both candidate pools and to the generator. The rules check blocks a card that slips through, with the code `excluded_precon_card`.
- The chat says what left: "I will use no card of X", the partial note of D-497, "holds no whole precon", or "no precon table loaded yet".
- The nine embedded lists stay (D-498). `TestEmbeddedListsMatchTheTable` compares each with its row and found one wrong card, now corrected (D-501).
- Deck gate prompt 25 covers the exclusion since 2026-09-04 (D-510), with the fixture `manabox_collection_avengers.csv`. Runs 15 and 16 read it the same day: 57 cards excluded, none in the deck.

The classify model saw version 17 in question gate runs 33 to 35, and the unit tests script its output. Deck gate runs 15 and 16 read a live exclusion on prompt 25 (2026-09-04).

CAUTION: the MTGJSON names are the product names, and 21 names belong to two or more products with different cards, "Deck A" among them. Such a name opens the precon row with the set codes as options. Ninety-nine products carry a one-word name, and one word resolves only as the whole name.

CAUTION: the exclusion resolves once, at classify time, into keys on the session. A later import of the collection does not move it. The reader names the precons again after an import.

## The corpus step of 2026-09-03 (D-494 to D-499)

The owner set the order on 2026-09-03: the corpus work first, then PR-24, then PR-14C, then the paid runs (D-494, D-495). Every new branch starts from `main`. The branch `pr-14c` holds no commit and waits.

`make meta-refresh` ran from 11:46 to 13:38 local, 112 minutes, and every source read with zero parse failures. The report per source:

- MTGO: 200 pages, 54 fetch errors, and 6,188 lists in the months it touched. The 54 are retired event pages that answer 302, and more pages wait behind the cap.
- MTGJSON: 702 deck files, and the table version `5.3.0+20260903` with 701 products.
- The cEDH database: 1 page.
- Topdeck.gg: 1 page and 31 lists, after two 429 answers with a 7-second wait each.
- EDHREC: 2,079 pages and 838 lists. The commanders file of 2026-09-03 holds 1,563 rows.
- The fit stored model `20260903T183812Z`, and the API loads it.

CAUTION: the EDHREC read ran although the last read was 2026-09-02. The skip reads a `read` marker under the day's raw prefix, and the run of 2026-09-02 wrote none. So the weekly stamp starts on 2026-09-03, and the next read falls on 2026-09-10 (D-499). D-495 rests on a wrong premise, the skip of the read, and its order of the paid runs stands.

That stamp never held either (F-50, D-565). The tournament lane writes the day's commanders file first, so the lane read the last read as today on every run. The read of 2026-09-07 sets the next read on 2026-09-14.

The MTGJSON version stamp carries the day, so the skip on an existing table never held before 2026-09-04. The job now compares the products of the deck list with the stored deck list of the newest table. It reads no deck file when they match (D-510), and a table older than 30 days reads whole again.

Quality gate run 11 is `docs/reference/pr14b-quality-gate-run11.md`, free. It reads FAIL on two bars. The gate fits its own model over the store, `20260903T183943Z`. The stored model is the job's fit of a minute before, over the same lists.

| Format | Great over precon | Precon over bad | Run 10 precon over bad |
|---|---|---|---|
| Commander | 0.98 | 0.78 | 0.88 |
| Standard | 0.92 | 1.00 | 1.00 |
| Modern | 0.99 | 0.9465 | 0.95 |

The Commander typical rung grew from 159 to 851 train lists, and the bad rung from 1,520 to 4,980. The breaks come from the baseline and the typical lists (D-484). The precons still grade bad: 28 of 44 holdout precons read bad, against 33 of 44 in run 10. The typical rung reads itself right in 100 of 186. The per-axis table reads lands 0.96, curve 0.96, colors 0.92, and synergy 0.77, against 0.98, 0.99, 0.98, and 0.75 in run 10. The bar fell because the model ranks a broken average deck above a precon more often than before.

The Modern weights moved little, and `card_rate` moved most, by 0.14. The great tier grew from 852 to 1,351 lists with the new MTGO pages, and the bar slipped under 0.95 by 0.0035. The weak axes are the same as in run 10: colors 0.61 and copies 0.80.

Three Commander weights changed sign, all small: `color_sources` from +0.02 to -0.02, `high_bracket_share` from -0.15 to +0.11, and `ramp` from -0.01 to +0.04. Four weights against the sense of their feature persist: `draw` -0.24, `wipe` -0.17, `land` -0.25, and `commander_decks` -0.18. The cEDH lists set the great and the good rungs, so the ladder reads a casual shape as worse. No weight changed by hand (D-486).

The explain mode over deck gate 13b ran free with the refit. It grades 19 of 24 built decks bad, 4 typical, and 1 good, and the detector flags 13. The run 10 model read 18 of 24 bad (D-488), so the wider typical rung did not move the built decks. The largest contributions against every Commander deck are `card_rate`, at 0.003 to 0.048 against a mean of 0.21, then `source_spread`, `fast_mana`, and `tapped_share`. Those are cEDH features. The judge bar stays open on the corpus (D-488, D-491), and the 60-card typical rung of PR-14C comes before the judge lane runs (D-495).

## Bracket gate run 1 (2026-09-02)

`docs/reference/pr14a-bracket-gate-run1.md` holds the builds, and `pr14a-bracket-gate-run1-judge.md` holds the judge lane. The builds cost $2.08 over 46 calls, above the $1.50 estimate, because nine decks took a repair turn and four took two. The judge lane cost $0.26.

| Bar | Result |
|---|---|
| Block checks | 15 of 15 |
| In every band | 8 of 15 |
| No content violation | 15 of 15 |
| Judge agrees | 8 of 15, 53 percent, the bar is 80 |

The off-band features: `mana_turn_four` 7, `avg_mana_value` 4, `color_sources` 2, `commander_turn_over_mv` 2, `tapped_land` 1. Brackets 1 and 2 sat in every band. The misses are all at brackets 3 to 5, and the bracket 5 decks miss most. A five-color warrior deck at an average mana value of 3.42 with no fast mana is not a cEDH deck. The profile said so.

The judge read every bracket 1 deck as a 2 or a 3, and read Kinnan as a 4 and Najeela as a 3. It named a Heliod and Archangel of Thune combo that the endpoint does not list, so its reasons are not facts.

No band moved on this run. The bracket 5 misses are the missing power signal of PR-14B, and the bracket 1 reads are the generator at its strongest on-theme list. Both are product findings, not band findings. OQ-53 records the one mixed message: the engine's land warning against the band.

CAUTION: the first judge lane failed on every call. The schema bounded an integer, and the Anthropic structured output refuses that (D-465). The re-judge mode of `bracket-gate -rejudge <document>` reads the decks back and judges them, for a quarter of the build cost.

## Deck gate run 12 (2026-09-02)

`docs/reference/pr8-deck-gate-run12.md` is the regression run under PR-14A. It cost $2.24 over 63 calls and took 44 minutes.

| Measure | Run 11 | Run 12 |
|---|---|---|
| Decks returned | 24 | 22 |
| Decks with no block finding | 24 | 22 |
| Invented names | 0 | 0 |
| False rules in a summary | 0 | 0 |
| Repair turns | 3 | 11 |
| Errors | 0 | 2 |
| Cost | $1.46 | $2.24 |

The two errors are prompt 3, the bracket 4 artifact deck, and prompt 23, two set families at once. On each, the generate model passed the three-minute deadline of the client twice, so the gate got no deck. Run 11 had no such error, and the other 22 prompts with pools as large answered in time.

Run 12b (`pr8-deck-gate-run12b.md`) reran the two alone: PASS, both decks clean, 266 seconds, $0.26. The errors were provider latency. The bracket 4 artifact deck of run 12b costs $3,974.63 to buy with no budget set. That is Mishra's Workshop and its friends in any-card mode, and no rule reads it.

Every repair turn ran for a profile finding. Eight decks ended off band. The sources of a color missed twice. Six other features missed once each. They are the average mana value, the mana on turn four, the draw count, the removal count, the wipe count, and the colorless land cap. One bracket 3 deck holds a near two-card combo the endpoint flagged, Storm-Kiln Artist with Haze of Rage, and the repair turn left it.

The F-33 read. Every Commander land count sat in its band, and the sources band caught two decks short of a color. The share of nonbasic lands still swings. The locked-card deck went from 33 nonbasic lands to 0, and the precon upgrade from 20 to 36, on the same prompt. No band reads the basic-to-nonbasic composition, so F-33 stays open on that point. The owned-first decks cost nothing to buy, and the tight budget deck costs $20.84 under its $25 cap, from $16.51.

## The bracket profile, decided (2026-09-02)

The owner asked how a bracket 3 deck can play like a true 3 (D-451 to D-453). PR-14 splits. PR-14A is the bracket profile. It holds the content rules per bracket, a feature vector per built deck with a band per bracket, a goldfish simulation, and a bracket gate. It comes right after PR-19. PR-14B is the learned scorer of D-413, after Phase 3B.

Commander Spellbook's `estimate-bracket` endpoint was verified on 2026-09-02. An anonymous `POST` with a text deck list returned a bracket tag. It also returned a flag per card for Game Changer, mass land denial, and extra turn, and a flag per combo for two-card and speed. The OpenAPI schema is at `backend.commanderspellbook.com/schema/?format=json`. OQ-50 holds the terms check. OQ-51 holds the Moxfield bracket field check for PR-14B.

## The commander offer for a request with no theme (2026-09-01)

Session `t8o1nGGquK6UdTQkfY3V` asked for the best deck, Commander, bracket 5, no colors, no budget, any card. The offer was Toski, Kutzil, and Mondrak. A free test over the snapshot showed why. The theme words were "best", "you", and "can", and each unknown word becomes a text needle. The words "you" and "can" sit in the text of most commanders, so 3,029 of them scored the same. The offer was the three most popular of those.

D-411 drops a needle that more than a tenth of the cards hold, and it adds the words of a request to the stop words. The same request now reaches the unthemed pool, which ranks on EDHREC popularity.

CAUTION: the app holds no power signal for a commander. The bracket drops Game Changers under bracket 3 and nothing else. A bracket 5 request gets the most popular commanders, not the strongest, until PR-14 brings a meta source. OQ-48 records it.

CAUTION: `buf breaking` runs against `main`, and PR-18 merged before the review renamed two summary fields. D-412 reserves the merged numbers. Never rename a field of a merged PR in place, whatever the deployment state.

## PR-14 is the deck quality model now (2026-09-01)

The owner asked for a model of what makes a deck good, bad, and great, over Standard, Modern, and Commander at every quality. Five decisions settle it (D-413 to D-417), and the PR-14 entry of the roadmap holds the plan. `docs/reference/deck-quality-sources-2026-09-01.md` holds the verified facts.

- An MTGO event page embeds every list and the standings in `window.MTGO.decklists.data`. A raw fetch reads it, and the fetch tool of a session does not, because the page renders on the client.
- The owner declined Topdeck.gg, then added it the same day (D-417). cEDH standings and decklists come from its API. The app shows the credit line "Tournament data by TopDeck.gg". The key is a secret in `.env` and Secret Manager.
- The database hosts its lists on Moxfield, and a fetch of the Moxfield terms answered 403. OQ-49 waits on the owner.

## The precon exclusion, decided (2026-09-01)

The owner asked for a deck that uses no card of a precon they own, for any set. Three decisions settle the slice (D-407 to D-409), and it waits in Phase 4 as PR-24. `docs/reference/precon-data-2026-09-01.md` holds the verified facts. MTGJSON lists 3,013 deck products with a Scryfall id per card, the four Marvel Super Heroes Commander decks among them.

CAUTION: the MTGJSON deck endpoint answers 403 to a request with no user agent. A session hit it with Python's default client and passed with curl and a named agent.

## The dev stack of 2026-09-01: no card index, and why

The owner's `make dev` log wrote "snapshot version check failed: bucket doesn't exist" every 15 seconds, and session `YvyZtBJyUiEcGMNnxOwm` failed on it. The API held no card index, so the commander question offered no name, the agent said it chose one, and the build ended with "no card index is loaded".

The cause was one file. fake-gcs-server 1.56.1 keeps the metadata of an object in an extended attribute, `user.metadata`. A session wrote `sets.json.gz` into the snapshot folder by hand on 2026-08-31 (PR-17B), and that file carried no attribute. Every listing of the bucket then answered 404, and a read of one object still answered 200.

The fix was to write the same bytes through the fake GCS API, so the attribute exists. The API loaded the snapshot on its next cycle.

```
curl -X POST "http://127.0.0.1:4443/upload/storage/v1/b/mtg-local-cards/o?uploadType=media&name=scryfall%2F<version>%2Fsets.json.gz" \
  -H "Content-Type: application/gzip" --data-binary @sets.json.gz
```

CAUTION: never write a file into `.local/gcs` by hand. Upload it through the API on port 4443.

The chat ran a turn with no card index before D-405. The commander question then offered no name, and the status line said "I have no more commanders that fit this deck, so I chose one for you", which was not true. `Chat` answers Unavailable before the turn now, as `ImportCollection` does.

## Next steps, in order

1. The three work items of 2026-09-04 are done (D-531). The precon check merged as #67, the terse conversations as #68, and the trimmed snapshot as #69 (D-533, D-541, D-543). The precon check ignores basic lands (D-523, F-35). The 47 terse conversations join the bar (D-522). One paid run of about $0.28 rebases the gate with the new threshold. A trimmed card snapshot of at most 10 MB serves Tier 0 (D-521).
2. The reword guard of D-88 stays at 0.60 (D-545). The M-5 evidence against it was the D-116 class. Four of the five live refusals scored better were the resolved row with an ask-role clause removed. The fifth was the format row with a clause added. Closed with no run.
3. The M-5 sheet is complete (D-529, D-530), and the fit threshold stays at 0.35: no candidate meets the 80 percent floor.
4. PR-27, the feedback harvest, is merged as #75 (D-561, D-562). PR #76 merged the records of D-562 and the rename of D-563 (D-564). PR-28, the feedback loop, follows the first real feedback (D-557), by `docs/reference/feedback-2026-09-06.md`, with OQ-70 to OQ-72 answered (D-559). The first real feedback needs the deployed app. PR-22 is merged as #72 (D-551). The deploy follows `docs/setup-gcp.md` on the owner's account, and the deploy half of `docs/reference/pr22-gate-2026-09-06.md` fills in after it. PR-23 is merged as #73 (D-554). PR-25 waits for the deploy (D-555). Then PR-25, the installable web app (D-547), and PR-26, the return channels, once the owner answers OQ-67. OQ-45 held the store of the allowlist, and D-420 answered it: one Firestore document, `config/allowlist`, written by `make allow EMAIL=...`. Then PR-23.
5. `make meta-refresh` daily, and `make quality-gate` to a new `QUALITY_GATE_OUT` after each one. Read the pair bars per format and the per-axis table. A weight against the sense of its feature is a defect in the feature or the labels. Do not tune it. The next weekly EDHREC read falls on 2026-09-14 (D-499, D-565).
6. Every paid run goes through `eval sweep` or its Makefile target. Each one writes its run file beside the document. The bracket rejudge of 2026-09-03 ran on 2026-09-05 (D-540). Ask the owner before each one.
7. Deploy the meta job (D-492). It is one Cloud Run job on `worker -meta`, with a Scheduler cron at 06:00 UTC daily. `TOPDECK_API_KEY` goes to Secret Manager. No infra file in this repo holds the worker's schedule. So the deployment is by hand, as the snapshot worker's is.
8. PR-22 and PR-23 came in order after PR-15, one gate each (D-511). The owner triggers the workflow `smoke` before a merge that touches the user path (D-313).
9. The weak-axes plan (D-567, D-568): M-7 is done, then PR-29 and PR-30, and PR-31 parks (D-573). Read `docs/reference/weak-axes-2026-09-07.md` first. No bar moves (D-486), and no weight changes by hand.

Deck gate run 12 ran on 2026-09-02 under the profile and passed 24 of 24 with its rerun 12b. The read of every mana base is F-33. The land count and the color sources sit in band now, and the nonbasic share still swings from 0 to 36 on the same prompt. No band reads the composition, and F-33 stays open on that point.

CAUTION: `make revise-gate | tee` hides the exit code. Read the verdict line of the document, never the exit code of a pipe.

CAUTION: the CI step "fake gcs tests" filters on `LiveStore`, and the only live test is `TestLiveFakeGCS`. The step matches no test and passes as a no-op. With the right name it needs a seeded snapshot bucket, which the CI server lacks. A separate change fixes it, on the owner's word.

Deck gate run 10 is done. It ran on 2026-08-31, and `CLAUDE.md` recorded it while this file still asked for it. A session that reads only the prose here spends $1.09 on a run that exists. Read `docs/reference/` before you plan a paid run.

A ruleset that requires the `verify` check on `main` is not possible. The repo is private on the free plan, and the rulesets API answers 403 (checked 2026-08-28). The owner reads the checks before a merge.

OQ-67 waits in `docs/owner-questions.md`: the channels and the events of PR-26. The owner answered OQ-54 on 2026-09-02, the Topdeck.gg key is in `.env`, and OQ-55 closed on 2026-09-03 (D-493). No row waits in `docs/open-questions.md`. The owner answered every open question of 2026-09-01 (D-419 to D-431). The judge role runs on Opus 5 now (D-430), and `make llm-defaults-check` reports the change on the next run.

## The dead conversation of 2026-08-31 (D-351 to D-354)

Session `0EqqY19J6A4BxCVgsmAE` stopped for good. The reader answered "No" to two yes-or-no questions, and both keys stayed in the asked state. Readiness needs no key in that state, so no build started. The agent never repeats a question it asked, so no new question came. Every later turn ended with no question, no message, and no deck.

Four changes close the hole:

- `State.DeclineNegative` reads a bare negative and closes the key (D-352). It pairs the answer with its question by id, so it matches no text.
- `Answer.declined` and the "You decide" control on a question card say the same thing outright (D-353).
- The classifier prompt now names a negative answer as a decline (D-354). Its "none closes no key" rule reads as commander_pick alone.
- A turn that asks nothing new and is not ready closes what is out and builds (D-351). That is the net under the three above.

CAUTION: D-352 and D-354 change what the questions engine does with an answer. Read the run 28 gate and its eval against run 27 before you trust a comparison with an older run.

## The paid runs of 2026-08-31

The re-baseline of D-302 is done for the question gate and the deck gate. Total spend: $1.31.

| Run | File | Result |
|---|---|---|
| Question gate 28 | `pr7-question-gate-run28.md` | PASS. 26 of 27 catalog-only, over the bar of 25. Run 27 sat on it at 25. $0.1589. |
| Question eval 28 | `pr7-question-eval-run28.md` | 6.8 percent bad on the holdout, from 7.2. The tune split reads 2.2 percent, from 4.1. $0.0916. |
| Deck gate 9 | `pr8-deck-gate-run9.md` | PASS, and it found a defect the verdict can not see. $1.0450. |

Deck gate run 9 ran with the unbounded owned-first fill. Four decks that cost nothing to buy on run 8 cost $39.81, $80.24, $60.77, and $167.98 on run 9. D-362 bounds the fill at 150 names.

Run 10 proves the bound. Every one of those four decks costs nothing again, and the two other owned-first prompts hold their shape.

| # | Prompt | Run 8 | Run 9 | Run 10 |
|---|---|---|---|---|
| 2 | aristocrats, owned first | $0.00 | $39.81 | $0.00 |
| 5 | blink, owned first | $0.00 | $80.24 | $0.00 |
| 13 | the commander is not owned | $0.00 | $60.77 | $0.00 |
| 15 | delegated commander | $0.00 | $167.98 | $0.00 |
| 16 | a tight budget | $6.17 | $6.83 | $23.03 |
| 17 | upgrade a precon | $73.55 | $77.12 | $78.47 |

Run 10: PASS, 18 of 18, one repair turn, $1.0787.

CAUTION: read the buy cost of each owned-first prompt, not the verdict. All three bars of this gate read legality, never cost. Prompt 16 rose from $6.17 to $23.03 over two runs. Its budget is $25, so no finding fires. Read it again on the next run.

## The dead-end check of the question gate (D-357)

`cmd/questions-gate/stall.go` fails a dead-end conversation. A dead end is a last turn that sent no question and did not report ready. Such a turn also left every slot as it found it, with a question still out. The gate reports a stall that a later turn recovers, and it fails neither.

The check has 11 unit cases and one live smoke run of 8 conversations, which found no dead end. The other 96 conversations of the set are unproven. Read the first full run: a false failure means one line goes, `len(deadEnds) == 0` in the pass expression.

## Speed, measured on 2026-08-30

Every number below comes from Playwright over the built app on `vite preview`, with each RPC stubbed. `scripts` in the scratchpad hold the runs. The same measurement on the dev server gives larger numbers, because Vite serves each module on its own there.

| What | Before | After |
|---|---|---|
| Content of `/session/new`, from navigation start | 347 ms | 44 ms |
| First open of the Build menu | 323 ms | 16 ms |
| Second open of the Build menu | 8 ms | 12 ms |
| Hop to a page it already read | 1 call | 0 calls |
| First-paint chunk, gzipped | 115.93 kB | 116.13 kB |

Three changes give that:

- `src/app/deferred.tsx` replaces `React.lazy` (D-338). A deferred unit starts a download and mounts the component the moment the code is here. Nothing suspends, so React throttles nothing.
- `src/app/chunks.ts` lists every deferred chunk, and the layout warms them all in the idle time after the first paint (D-339).
- The query client keeps server state fresh for 30 seconds (D-340).

CAUTION: the first-paint bar of D-323 is 130 kB gzipped. Read the Vite build report after any change to `src/app/chunks.ts`. A chunk that moves into the entry chunk spends that budget.
