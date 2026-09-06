# MtG Deck Builder - Design Roadmap

Status: **approved (D-31), living document.** The owner approved draft 1 on 2026-08-23. Each PR entry carries its merge status. A ✅ means the code is merged on `main` (D-42). The doc follows the structure of `connector-syncer-docs/docs/document-summary-roadmap.md`. The text follows ASD-STE100.

External facts were verified 2026-08-23, with 2026-08-24 re-passes noted inline. Sources: Scryfall (API and bulk data), the Wizards of the Coast announcements of 2026-08-10 and 2026-02-09, mtgcommander.net, and the local checkouts of connector-syncer and wallabee-ui. The Comprehensive Rules file is `MagicCompRules 20260819.txt`, verified 2026-08-28 (D-272). The pass of 2026-08-23 checked every rules citation here against the 2026-08-07 text, and the sub-letters are unverified against the new file. MtG rules and ban lists change. Re-verify every dated fact before you cite it in a PR.

Owner decisions live in `docs/decisions.md` (D-#). Open questions live in `docs/open-questions.md` (OQ-#). The decision queue lives in `docs/owner-questions.md`. Research notes live in `docs/reference/`. The MtG knowledge base lives in `.claude/skills/mtg-corpus/`.

2026-09-05 correction pass 81 (the bracket rejudge, D-540): the second judge lane read the 15 decks of bracket gate run 1. It agrees on 8 of 15, the same as the first lane. The bar of 80 percent reads FAIL again. Two bracket 5 decks moved between the lanes: Kinnan up to 5, Yuriko down to 4. Najeela reads 3 in both. The bracket 5 misses stay the power signal of PR-14B, and no band moves. The lane cost $0.28. The document prints a JSON escape for an em dash in three reasons (F-48). Changes: D-540, F-48, `docs/reference/pr14a-bracket-gate-run1-judge2.md`.
2026-09-05 correction pass 80 (run 42 passes, the GCP guide, D-539, OQ-65): question gate run 42 read PASS on every bar, with no expectation miss. It read 73 of 74 counted conversations catalog-only against 64. It is the questions baseline now, in place of run 35. `docs/setup-gcp.md` is the guide from the domain to the deployed app. Its cost estimate covers five users and three decks a week each. The month costs $8 to $20 in model calls, $0 in Google Cloud inside the free tiers, and $1 to $2 for the domain. It found one open question for PR-22 (OQ-65). The Hosting rewrite to Cloud Run has no documented timeout, and the Cloud Functions rewrite has one of 60 seconds. Changes: D-539, OQ-65, `docs/setup-gcp.md`, `docs/reference/eval/baselines.json`.
2026-09-05 correction pass 79 (run 41 read, D-538): question gate run 41 read FAIL on one miss. Every other bar passed: 73 of 74 counted conversations catalog-only against the bar of 64. The miss was conversation 82 again. The classifier returned both spellings of Atraxa in one turn, the known one first. The D-535 rule set the right name and then let the misspelled one join it. An unknown name never joins a known commander now (F-43). Run 42 rebases every number, on the owner's word. Changes: D-538, `internal/questions/agent.go`.
2026-09-05 correction pass 78 (run 40 read, D-537, F-47): question gate run 40 read FAIL on one miss. Every other bar passed: 74 of 74 counted conversations catalog-only against the bar of 64. The miss was conversation 25. The classifier answered the open budget question with 2 on a message with no number, and the session closed before the user named 400. A budget applies only when the message names it now, the D-125 rule for the budget. Three of the four slots the classifier repeats carry that rule: the format (D-125), the colors (D-536), and the budget (D-537). Run 41 rebases every number, on the owner's word. Changes: D-537, F-47, `internal/questions/agent.go`, `internal/questions/words.go`.
2026-09-05 correction pass 77 (run 39 read, D-536): question gate run 39 read FAIL on one miss. Every other bar passed: 74 of 74 counted conversations catalog-only against the bar of 64. The miss was conversation 74 again. The colorless rule closed the slot on turn 2, and the classifier answered "Bracket 3." with all five colors over it. A closed color slot changes only on a message that names a color now (F-44, the D-125 rule for colors). Run 40 rebases every number, on the owner's word. Changes: D-536, `internal/questions/agent.go`, `internal/questions/words.go`.
2026-09-05 correction pass 76 (run 38 read, D-535, F-43 to F-46): question gate run 38 read FAIL on 13 expectation misses. Every other bar passed: 73 of 74 counted conversations catalog-only against the bar of 64, no premature build, no dead end, and no lint finding. Nine misses were rows. Four picks after a delegation now expect "delegated" (D-147). Four answers sat in turns the agent never sends, or the user never gave them. The colorless rows expect the new word "none". Four misses were defects, and each has a deterministic rule and a test now. They are the corrected commander name (F-43), the colorless answer (F-44), the budget scope words (F-45), and the superlative theme (F-46). The bar stays strict: a miss becomes a row fix or a rule, never a tolerance. Run 39 rebases every number, on the owner's word. Changes: D-535, F-43 to F-46, `internal/questions/agent.go`, `cmd/questions-gate/expect.go`, `cmd/questions-gate/conversations.json`.
2026-09-05 correction pass 75 (the terse conversations join the bar, D-522, D-534): the 47 terse conversations of ids 57 to 104 are gate conversations now. Each carries a golden expectation written from the words of its user. The live gate holds 77 conversations, and the catalog-only bar is 64, the share of 25 in 30. `questions.MinGateSize` stays at 30 for the offline half. The set is dated 2026-09-05, so the first run after the change rebases every number of the question gate, on the owner's word. Branch `terse-bar`, the second of the three items of D-531. Changes: D-534, `cmd/questions-gate/conversations.json`, `cmd/questions-gate/main.go`, `internal/questions/metrics.go`, the PR-7 gate text.
2026-09-05 correction pass 74 (the precon check ignores basic lands, D-523, D-532, F-35): `OwnedWhole` skips the printings of basic lands, with the same Oracle-id test the exclusion uses. The chat hints, the deck gate, and the build share `candidates.BasicLandByOracle`. The three basic-land rows D-510 added to the Avengers fixture are gone. The dry run of prompt 25 reads the binder as the whole product. Branch `precon-basic-lands`, the first of the three items of D-531. Changes: D-532, F-35, `internal/precons/table.go`, `internal/candidates/candidates.go`, `internal/questions/hints_candidates.go`, `cmd/deck-gate/main.go`, `internal/collections/testdata/manabox_collection_avengers.csv`.
2026-09-04 correction pass 73 (the M-5 sheet complete, D-529, D-530): the owner scored Items 33 to 41 by phone and took every recommendation from Item 39 on. The session scored Items 42 to 60 on its own recommendation, with the reason in each row (D-529). `make m5-report` reads 60 of 60 scored, no threshold at the 80 percent floor, and the fit threshold stays at 0.35 (D-530). 7 of 29 scored refusals were better than their row, so the reword guard of D-88 is too tight. Changes: D-529, D-530, `docs/reference/pr7-m5-scoring.md`.
2026-09-04 correction pass 72 (probe 109 ran, F-42, D-526): the group set probe ran as question gate run 36 for $0.0016. The sets expectation held: msc, msh, pspm, spe, and spm. The run exposed a gap of the harness. A partial run under `-only` read FAIL on the two bars over the whole set. Then `make eval-check` compared it with the full baseline as the run of record. The fix merged as #66 (D-528). The run header records the `-only` flag. The question gate reads a partial run on its item bars alone. The check lists a partial run and compares none. Run 37 is the probe of record, and the run 36 header carries the flag by hand (D-527). Changes: F-42, D-526, D-527, `internal/evalrun/evalrun.go`, `cmd/eval/baselines.go`, `cmd/eval/main.go`, `cmd/questions-gate/main.go`, the three deck gates.
2026-09-04 correction pass 71 (the Marvel set question, D-518, F-40, F-41): a Marvel-only request asked "Which one: Marvel Super Heroes and Marvel Universe?". Marvel Universe is a bonus sheet, Scryfall type masterpiece, and it never leads a set family beside a product now. The choices of the set question and the precon question read with "or". "Marvel" resolves to Marvel Super Heroes alone, and Marvel's Spider-Man stays out by the owner's word. The same branch holds the group request of D-525. Its parts: `set_groups` in the classify role at prompt version 18, `ResolveGroup` in the set table, and probe 109 of the question gate. The probe waits for a paid run of about $0.20 on the owner's word. Changes: D-518, D-525, F-40, F-41, `internal/cards/sets.go`, `internal/questions/resolve.go`, `internal/questions/prompts.go`, `cmd/questions-gate/conversations.json`.
2026-09-04 correction pass 70 (PR-15 merged, D-517): the owner merged #65 after the paid gate passed on every lane. A reset of `main` from a broken commit to #64 came first. The three suites have baselines, `make eval-check` runs in CI as Tier 0, and the next paid run of each suite is its first compare. `docs/setup-second-mac.md` holds the move to another Mac. Changes: D-517, the PR-15 entry, sequence item 21.
2026-09-04 correction pass 69 (the PR-15 paid gate, D-514 to D-516, F-36 to F-39): the three steps of D-513 and the two runs of D-516 ran for $5.89. Deck gate run 16 passes 25 of 25, and the compare names no gate flip against run 14 with 14b. Run 16 is the decks baseline with the first complete fingerprint. Tier judge run 4 reads 7 of 25 on the open bar, and the sweep stopped there. Revise gate run 9 ran after it and passes 11 of 11 (D-516). Question gate run 34 read 21 misses on the fourth bar. 17 came from the renderer, which read id fields the classifier never fills. 3 came from expectations that named a message that never goes out, and 1 is a wrong slot (F-36). The renderer, the expectations, and the run file of a FAIL are fixed on the branch, and run 35 passes with no miss (D-516). The plan judge wrote "placeholder" as a reason in 10 of 25 decks (F-38) and graded legality from stale card knowledge (F-39). The Avengers prompt with its precon excluded builds a deck with no superhero identity, and the summary hides it (F-37). The gate document is `docs/reference/pr15-paid-gate-2026-09-04.md`. Changes: D-514 to D-516, F-36 to F-39, `docs/setup-second-mac.md`, `cmd/questions-gate/expect.go`, `internal/generate/planjudge.go`, `docs/reference/eval/`.
2026-09-04 correction pass 68 (PR-15 built, D-512, OQ-61): branch `pr-15` holds all six slices of the harness. Every gate writes a run file beside its document, and the "## Run" block of every document is one fingerprint. `cmd/eval` compares, checks, records a baseline, imports a deck gate document, and sweeps under a cap. The question gate gains a fourth bar, the golden expectations of its 27 counted conversations. The compare of deck gate runs 13b and 14 names deck 17 as the one flip, which is F-34. Tier 0 runs in CI as `verify:eval`. The free gate is `docs/reference/pr15-gate-2026-09-04.md`. The owner said go to the paid gate the same day (D-513), and the next session runs it in three steps. Changes: PR-15, D-512, D-513, OQ-61, `docs/reference/eval/`.
2026-09-04 correction pass 67 (PR-15 moves ahead, D-511, OQ-57 to OQ-60): the owner chose the eval harness before PR-22 and PR-23. A session read the eight gate writers, and the hand-off holds the plan in six slices. The "Tier 1 nightly" line of the entry is a question now. The sweep of 2026-09-03 cost $4.60, so a nightly run is about $140 a month (OQ-57). Changes: D-511, OQ-57 to OQ-60, the PR-15 entry, sequence item 21.
2026-09-04 correction pass 66 (the corpus items, #64, D-510, F-35, OQ-56): the meta job skips the MTGJSON deck files when the deck list names the products of the stored table. A table older than 30 days reads whole again. Deck gate prompt 25 covers the precon exclusion of PR-24 with a fixture that holds Avengers Assemble once. The first fixture held every card of the product twice, so the exclusion excluded nothing. The export of 2026-08-24 lacks six basic lands of the product, so D-408 says that binder owns no precon (F-35). OQ-56 asks the owner about basic lands in the check. Changes: D-510, F-35, OQ-56, `cmd/deck-gate/prompts.json`, `internal/collections/testdata/manabox_collection_avengers.csv`.
2026-09-04 correction pass 65 (the paid runs after PR-21, F-34, D-509): question gate run 33 passes, and its eval accepts it. Tier judge run 3 reads 8 of 24, with the bar open on the corpus. Deck gate run 14 blocks one deck of 24. The block is a defect: the bracket cut and the precon share fought, and a failed repair replaced a legal deck (F-34). Both fixes merged as #63 (D-509), and run 14b passed the prompt with them. Revise gate run 8 passes 11 of 11. Changes: F-34, D-509, `docs/reference/pr7-question-gate-run33.md`, `docs/reference/pr7-question-eval-run33.md`, `docs/reference/pr14b-quality-judge-run3.md`, `docs/reference/pr8-deck-gate-run14.md`, `pr8-deck-gate-run14b.md`, `docs/reference/pr12b-revise-gate-run8.md`.
2026-09-03 correction pass 64 (PR-21 merged as #62, D-508): the share link, the public page at `/d/<token>`, and the print view. The shared message is its own type with the card data inline and no user field. A test reads the proto text to prove it. The store keeps a hash of the token in its own document, so a lookup needs no index. A limiter of 60 calls a minute per forwarded address bounds the two public reads. Changes: PR-21, D-508, `docs/reference/pr21-gate-2026-09-03.md`.
2026-09-03 correction pass 63 (PR-20 merged as #61, D-507): the deck view gains the card detail as a Sheet. It gains the filters and the sorts, five stats with a text table each, and the sample hand of D-318. The Scryfall rulings file joins the snapshot as a fourth file, and `GetPrintings` joins `GetRulings`, because the index held one printing per card. The real rulings file of 2026-09-03 parsed into 78,949 rulings of 19,938 cards. Changes: PR-20, D-507, `docs/reference/pr20-gate-2026-09-03.md`.
2026-09-03 correction pass 62 (PR-14C merged as #60, D-502 to D-506): Aetherhub answers a Cloudflare challenge page and leaves the plan (D-502). The owner kept the legal check of D-5 over the content signal of the robots files (D-503). The MTGTop8 reader takes the paper events of three formats and skips the MTGO copies (D-504), and its tier reads the field (D-505). The MTGTop8 format page holds its event links in plain HTML, so the D-477 sentence carries a dated correction. The first live run read 500 pages with zero failures. Quality gate run 12 reads FAIL on the Commander and the Modern precon bars, 0.78 and 0.88 (D-506). Changes: PR-14C, D-502 to D-506, `docs/reference/pr14c-sources-2026-09-03.md`, `docs/reference/pr14c-gate-2026-09-03.md`, `docs/reference/pr14b-quality-gate-run12.md`.
2026-09-03 correction pass 61 (the corpus step and the PR-24 calls, D-494 to D-499): the owner set the order: the corpus work, then PR-24, then PR-14C, then the paid runs. The refresh read 2,079 EDHREC pages into 838 lists, because the run of 2026-09-02 left no completion marker. Quality gate run 11 reads 0.78 on the Commander precon bar and 0.9465 on the Modern one, both FAIL. Under the refit, 19 of 24 built decks grade bad. The classifier reads the precon exclusion (D-496). The build excludes a named precon the collection does not hold whole (D-497). The nine embedded lists stay beside a cross-check test (D-498). PR-24 merged as #59 with a free gate, and the cross-check found one wrong card in the Limit Break list (D-500, D-501). Changes: PR-24, D-494 to D-501, `docs/reference/pr14b-quality-gate-run11.md`, `docs/reference/pr24-precon-gate-2026-09-03.md`.
2026-09-02 correction pass 60 (deck gate 13b and the judge lane, D-486 to D-488): deck gate run 13 read FAIL on two defects of the branch, both fixed. Run 13b passed 24 of 24. The tier judge lane over it read 5 of 24, then 4 of 24 after the breaks drew from seen cards. The explain mode named the cause: the ladder holds 4,000 tournament lists per tier against 195 community decks, and none for the 60-card formats. Gate run 10 reads 0.97 to 1.00 on the top-list bar and 0.88, 1.00, and 0.95 on the precon bar. The judge bar stays open on that finding, and PR-14B merges with it on record (D-491). The owner widened the EDHREC read to 1,100 commanders (D-489). PR-14C gains the casual 60-card decks (D-490), and the meta job runs daily on Cloud Run (D-492). Changes: PR-14B, PR-14C, D-486 to D-492, `docs/reference/pr14b-quality-judge-run2.md`.
2026-09-02 correction pass 59 (the full read and gate runs 1 to 8, D-479 to D-485): the owner set the windows and the Topdeck.gg key arrived. The full read holds 6,187 MTGO lists, 21,146 cEDH lists, 701 precons, and 1,007 commander reads. Gate run 8 puts every top list over the precons (0.96 to 1.00). It puts the precons over their broken copies at 0.94, 1.00, and 0.95 against the 0.95 bar. A defect detector sits beside the ladder (D-485). Deck gate run 13 started under the stored model. PR-14C holds MTGTop8 and Moxfield (D-482). Changes: PR-14B, PR-14C, OQ-54, OQ-55, `docs/reference/pr14b-quality-gate-run8.md`.
2026-09-02 correction pass 58 (PR-14B built, D-470 to D-478): the session built the deck quality model on branch `pr-14b`. It holds the source readers, the store, the precon table, the fit, the scorer, the wiring, and a free gate. A short live read passed on MTGO, EDHREC, and the database, and no paid run happened. Moxfield answers 403 to a plain client, so the database gives the tier and the commander alone (D-470). The Topdeck.gg key waits on the owner (OQ-54). Changes: PR-14B, F-30, OQ-48, OQ-51, `docs/reference/deck-quality-model-2026-09-02.md`.
2026-09-02 correction pass 57 (deck gate run 12, F-33): 22 of 22 decks pass every block check, no invented name, no false rule, $2.24. Two prompts got no deck, because the generate model passed its three-minute deadline twice on each, and the verdict reads FAIL on those errors alone. The repair turn ran on 11 decks, against 3 in run 11, and every reason was a profile finding. Run 12b reran the two prompts alone and passed both in 266 seconds, $0.26, so the errors were provider latency. The count of nonbasic lands still swings, 0 to 36 on the same prompts. Changes: PR-14A, F-33.
2026-09-02 correction pass 56 (the owner's answers, D-467 to D-469): the session calls of PR-14A stand, with one change. Commander Spellbook reads the shortlist too, before the build, and the builder drops what it flags for the bracket (D-468). The engine's Commander land range is 27 to 41, the union of the bands (D-469, amends D-60). Changes: PR-14A, OQ-52, OQ-53.
2026-09-02 correction pass 55 (bracket gate run 1, D-465 and D-466): 15 of 15 decks pass the block checks, and none holds a content violation. 8 of 15 sit in every band, and the judge agrees on 8 of 15. The verdict is FAIL on the band bar and the judge bar. The judge lane ran twice, because the first schema refused every call (D-465). The near two-card combo joins the check (D-466). The bracket 5 decks are theme decks with no power signal, which PR-14B answers. Changes: PR-14A, OQ-53.
2026-09-02 correction pass 54 (PR-14A built, D-461 to D-464): the session built the bracket profile on branch `pr-14a`, and its gate did not run yet. A profile finding warns and buys the repair turn, and the tags cut the shortlist. Commander Spellbook checks the deck, and the bands start as estimates the gate tunes. The session calls wait for the owner as OQ-52. Changes: PR-14A, F-11, F-33, `docs/reference/bracket-profile-2026-09-02.md`.
2026-09-02 correction pass 53 (OQ-50 closed, PR-14B moved, D-459 and D-460): the owner read the Commander Spellbook terms, and they allow the bracket endpoint. The app holds itself to 90 requests a minute. The owner asked why PR-14B sat after Phase 3B, and no code dependency held it there. PR-14B now comes right after PR-14A, and it builds the precon table of D-407 that its baseline tier needs. PR-24 follows it. Changes: PR-14A, PR-14B, PR-24, sequencing steps 19 and 22.
2026-09-02 correction pass 52 (deck gate run 11, F-33): 24 of 24 decks pass, $1.46. A read of every mana base shows the count of nonbasic lands swings from 0 to 33 on the same shortlist, run to run. D-450 cut the colorless lands by two thirds. The mana base is a band for PR-14A, not a prompt line. Changes: F-33, PR-14A gate.
2026-09-02 correction pass 51 (the bracket profile, D-451 to D-453): the owner asked how a bracket 3 plays like a true 3. PR-14 splits into PR-14A, the deterministic bracket profile with a goldfish simulation, and PR-14B, the learned scorer. PR-14A comes right after PR-19. Commander Spellbook's bracket endpoint was verified the same day. Changes: PR-14, PR-14A, sequencing steps 19 and 22, OQ-50 and OQ-51.
2026-09-02 correction pass 50 (the theme fills the land bucket, F-32, D-450): revise gate run 5 passed 10 of 11 turns. It failed the owner's own land ask on the Karlov deck, because the shortlist offered no untapped dual. The land cap splits: half the mana order, half the theme. Changes: F-32, PR-6, PR-12B. Run 5 stays as the record, run 6 and deck gate run 11 are due.
2026-09-02 correction pass 49 (a land upgrade rebuilt one basic, F-31, D-448 and D-449): the owner asked for better lands in place of the basics, and the rebuild moved one Plains to one Island. The brief counts the basic lands to replace, and the check holds the deck to the count. The generator's revision rule names group changes, and the revise gate plays the turn after its question. The generate prompt is at version 11. A turn stores its brief. Changes: F-31, PR-12B, the cost model line of the revise gate. Revise gate run 5 is due.
2026-09-02 correction pass 48 (the owner's read of PR-19, D-438 to D-447): the unfinished chats moved under the message box and hide when there is none. A deck tile carries a delete, and every control shows the pointer cursor. The delete dialogs stuttered, and an elimination run found the cause in the wordmark's endless shimmer, not in the dialog. The gap over the message box equals the gap under the top bar. Changes: PR-19.
2026-09-02 correction pass 47 (PR-19 built, then the form refused, D-432 to D-437): PR-19 built a start form, the unfinished chats, the stepper, and the retry. The owner read the form and refused it: the app is chat, and the pool picker is the one control outside it (D-436). The form left the same day. A session the owner read found the commander offer bare under a declined theme. The offer now serves an empty theme and reads the set limit (D-437). Changes: PR-19.
2026-09-01 correction pass 46 (Topdeck.gg joins the sources, D-417): the owner added the Topdeck.gg API and its credit line the same day. It amends D-415. cEDH standings and decklists come from the API. The app shows "Tournament data by TopDeck.gg" with a link. Changes: the system map, PR-14, guardrail 7 unchanged.
2026-09-01 correction pass 45 (PR-14 becomes the deck quality model, D-413 to D-416): the owner asked for a model of what makes a deck good, bad, and great. It covers Standard, Modern, and Commander at every quality. PR-14 was a meta feed. It is a scorer now. It reads the lists that MTGO, MTGTop8, the cEDH database, and EDHREC publish. It labels them by placement, and it fits a weight per feature per format. The owner declined Topdeck.gg, then added it (D-417). The owner chose synthetic bad decks as the bottom of the ladder, and one PR. Changes: the system map, the cost model, F-30, PR-14, Phase 5, sequencing step 22.
2026-09-01 correction pass 44 (three findings from the owner's reads, D-410 to D-412): the deck name carried the format twice. The commander offer for "the best deck you can" ranked on the words "you" and "can", which sit in the text of most cards. The pool now drops a generic text needle that more than a tenth of the cards hold. The words of a request are stop words. The merged PR-18 summary fields stay reserved, because the wire guard refused their rename. Changes: PR-6, PR-8, PR-18.
2026-09-01 correction pass 43 (the precon exclusion, D-407 to D-409): a reader asks for a deck that uses no card of a precon they own, for any set. MTGJSON publishes every WotC deck product with a Scryfall id per card, 3,013 products on 2026-09-01. The owner chose the Commander decks and the 60-card constructed decks. Ownership reads the printings and their counts, the surplus copies stay usable, and the slice sits in Phase 4 beside PR-14. Changes: Phase 4 gains PR-24, sequencing step 22.
2026-09-01 correction pass 42 (the review of PR-18, D-398 to D-404): PR-17 merged (#49), PR-17B merged (#50), and PR-18 merged (#53). A review of PR-18 found three defects and four gaps. The binder filtered the loaded pages alone, so a search for a card past row 200 found nothing. A Replace kept an id whose hash moved on, so a later upload of the old file overwrote the replaced collection. Every upload over an active collection forced a replacement. The fixes ship on branch `nits-and-fixes`. The filter, the sort, and the search run on the server now. The summary lists every set and every type, and the upload dialog offers a choice. The chat refuses a turn with no card index (D-405). A decline of the format through the "You decide" control takes the corpus default, as a decline in words does (D-406). Changes: PR-17, PR-17B, PR-18, sequencing steps 19 and 20.
2026-08-30 correction pass 35 (the reference design and the one deck screen, D-328 to D-335): the owner gave a reference design, and it settles the look. Three faces, a navy and gold palette, one top bar, and dark alone. The identity wash of D-327 left, and the color of the game now shows in a mana pip and a rarity dot. A deck has one screen and one address. Changes: Phase 3B, PR-16, PR-16B, PR-17, sequencing step 19.
2026-08-30 correction pass 36 (the collection picker and the speed of the app, D-336 to D-340): the pool lived in the header menu alone, and the owner did not find it. A "Build from" picker now sits on the chat screen. No screen shows a decision id or a Firestore id. No chunk loads behind a Suspense boundary. React holds a committed fallback for 300 ms. Content on the landing page went from 347 ms to 44 ms. The first open of the Build menu went from 323 ms to 16 ms. Changes: PR-17.
2026-08-30 correction pass 37 (the version history and the compare, D-343): the last two items of PR-17. The decks of one chat are the versions of one deck, so `ListDecks` takes a session id and the deck screen reads its own history. Any two versions compare with the diff the revision note already used. Changes: PR-17, sequencing step 19.
2026-08-31 correction pass 38 (the dead conversation, D-351 to D-355): a bare "No" to a yes-or-no question left its key in the asked state. The agent never repeats a question it asked, so no new question came, and the session was never ready. Every later turn did nothing. Three fixes now close such a key. The code reads the negative. A question card carries a "You decide" control. The classifier prompt names the shape. A turn that asks nothing and is not ready closes what is out and builds. Changes: PR-17, PR-7.
2026-08-31 correction pass 39 (the gate dead end, the pool floor, and the fixed frame, D-357 to D-364): the question gate now fails a conversation that can not move. A commander pair option carries two card ids, because one id can not hold two cards. The owned-first fill reaches a viable floor and not the shortlist cap. A fill to the cap turned four free decks into decks that cost $40 to $168 on deck gate run 9. The shell is a fixed frame, and the docked chat always fits it. Changes: PR-7, PR-8, PR-17.
2026-08-31 correction pass 40 (the commander pool and the deck tile, D-365 to D-367): `CommanderPool` dropped every commander with no theme signal. A theme the tag table does not know left three names or fewer, and a reader who refused those had nothing left to read. The theme still leads, and under a floor of 12 the pool fills from the whole format. The agent now says so when it takes the commander choice. The whole deck tile opens the deck. Changes: PR-6, PR-7, PR-17.
2026-08-31 correction pass 41 (the set filter, F-29, D-371 to D-383): the app never applied a set as a constraint. The request carried no set, and the words reached the theme alone. A deck asked for one set then held cards of any set. PR-17B adds `Card.set_codes` and the filter over it. A card holds printings in 2.4 sets on average, so the field is a list. A set name is a product, and it resolves to a whole set family through the Scryfall parent link. The classifier also read "build only from the Hobbit set" as an ownership rule. An owned rule with no collection then emptied every pool. Changes: PR-17B, F-29, sequencing step 20.

2026-08-29 correction pass 34 (the palette of D-327): the five colors of the game are the app's palette, and dark leads. D-311 kept every surface neutral, and the result read as boring. A deck now carries its own color identity, and a card role carries its own hue. Changes: PR-17, guardrail 13 unchanged.

2026-08-29 correction pass 33 (PR-17 started on branch `pr-17`): the proto and the Go side of the deck library. `UpdateDeck`, `DeleteDeck`, `Deck.favorite`, `Deck.card_count`, and the paged, filtered `ListDecks`. D-324 puts the power filter on the server. Changes: PR-17, sequencing step 19, open question 7.

2026-08-29 correction pass 32 (PR-16B, the visual pass, D-325 and D-326): the screens of PR-16 kept the composition of PR-11, so the app read as a test bench. The pass adds a type face, an elevation scale, a docked composer, and a thread with hierarchy. Changes: PR-16B, sequencing step 19.

2026-08-29 correction pass 30 (PR-16 built on branch `pr-16`, D-323): the design system, the two themes, and the app shell. The first paint is 116.31 kB gzipped, from 179.81 kB. The bar of D-320 was out of reach, and D-323 amends it. Changes: PR-16, sequencing step 19.

2026-08-29 correction pass 29b (D-320 to D-322): the PR-16 bundle gate reads raw bytes, not gzipped bytes, and the slice adds only the primitives that it uses. The PR-22 entry no longer decides OQ-45. The PR-21 rate limit reads the forwarded address. Changes: PR-16, PR-21, PR-22, `docs/open-questions.md` table.

2026-08-29 correction pass 29 (Phase 3B, the product UI, D-310 to D-319): the live-test UI of PR-11 to PR-13 is not a product (F-28). A new phase, PR-16 to PR-23, builds one. It holds the design system, the four flows of D-312, the share link, the deploy for invited users, and a Playwright smoke flow. `docs/reference/ui-phase-plan-2026-08-29.md` holds the detail. Changes: F-28, guardrail 13, the system map row of `web`, Phase 3B, Phase 5 (D-318), sequencing steps 18 to 23, open questions 5 and 6.

2026-08-29 correction pass 28 (PR-13 built on branch `pr-13`, D-307 to D-309): `DeckService.ExportDeck`, the export panel, and the buy list with a Scryfall link per card. The round-trip gate holds in `go/internal/export`. Changes: PR-13, sequencing step 18.

2026-08-29 correction pass 27 (quality audit, `docs/audit-2026-08-29.md`, D-302 to D-306): PR-12B merged (#41). The generate prompt is at version 10, and the classify prompt version rose. The gate baselines of question run 27 and deck run 8 do not compare with `main` until the owner re-runs them (D-302). Consistency fixes of this pass: F-27 and step 18 show PR-12B merged, and the PR-1 layout names the packages the code uses. The eval cost is $0.092 to $0.104, and the cost table gains the deck gate. The pass strikes the PR-12 Oracle-text hover (D-291). The grammar pass of D-304 rewrote passive sentences, modal verbs, and -ing forms across the doc, with no change of fact.

2026-08-28 correction pass 26 (PR-11 merged, #38): the web shell, sign-in over the Auth emulator, and the collection screen. The gate held in the browser the same day. Changes: PR-11, sequencing step 18.

2026-08-28 correction pass 26 (the first live conversation): the agent dropped a message after a build, and it rebuilt the deck from the first message (F-27). PR-12B adds the revision turn (D-283).
2026-08-28 correction pass 25 (UI plan, `docs/reference/ui-plan-2026-08-28.md`): the owner scoped the live test (D-273 to D-276). It covers the whole user path, on the roadmap stack, with real sign-in over the Auth emulator, locally. PR-12 gains `CardService.GetCards`, and PR-13 gains `DeckService.ExportDeck`.

2026-08-28 correction pass 24 (full audit, `docs/audit-2026-08-28.md`): PR-8 merged (#15), and #16 ignores every command binary (D-255). PR-9 leaves the MVP (D-256). The audit found the deployable API was unable to build a deck, and the owner ruled it a defect (D-257). Decisions D-247 to D-272 recorded, five catalog rows retired (D-260), and the gate set changed with no run (D-263). The Comprehensive Rules file is 2026-08-19 (D-272). The register table is one table again, the freeze and the restricted check carry the retired mark, and the STE check runs in `make lint` (D-264).

2026-08-27 correction pass 23: PR-8 built, and its gate runs began. The build-run freeze is retired (D-241). The gate harness can start a conversation after a build (D-239). The judge lane answers F-26 (D-229), and D-227 measures the cache saving. Decisions D-226 to D-243 recorded.

2026-08-25 correction pass 22: the owner's scoring found a seventh fault and a hole in the reword guard (D-103). The guard now refuses a truncation.

2026-08-26 correction pass 21 (owner directive): the app builds three formats, Commander, Standard, and Modern (D-155). D-155 removes Pioneer, Legacy, Vintage, and Pauper. The proto keeps no dead ids: the proto deletes and reserves the four enum values. `buf.yaml` moved from FILE to WIRE_JSON, which still refuses every wire-breaking and JSON-breaking change. No deployment exists, so no session needed migration. Two rules died with Vintage and left the code: the restricted-card check had no format left that carries a restricted list. "Anything goes" is not a format the user can select. It stays a house-rules layer on one of the three (D-3).

2026-08-26 correction pass 20: PR-7B added, the automated eval lane. The owner's hand scoring does not scale past 32 items, and a run asks about 250 questions. A new `eval` role scores every one for about eleven cents (D-133). Every third conversation is a holdout (D-134). The owner answered OQ-24, OQ-26, and OQ-27 the same day (D-135 to D-137). PR-7 gained sixteen fixes from the batch sweep of all 66 conversations (D-117 to D-132).

2026-08-25 correction pass 19: the owner scored items 1 to 32 of the M-5 sheet, and the scores asked for 16 rewords and one deletion. The correction session that followed recorded D-104 to D-116. The sheet held only 60 of 793 questions, which is where most defects hid (D-104).

2026-08-25 correction pass 18: two gaps closed before PR-7 is committed. The session store now runs against the Firestore emulator (D-101), and `cmd/m5-report` reads the scored sheet and computes the thresholds (D-102).

2026-08-25 correction pass 17: runs 12 and 13 confirm D-98 and D-99 under live conditions. The M-5 sheet holds 60 items and explains its own fields (D-100). PR-7 has no open item but the owner's scoring.

2026-08-25 correction pass 16: 22 probe conversations joined the gate file (D-97). They found two defects on their first run, D-98 and D-99, and three catalog gaps. The M-5 sheet reads engine-current runs alone (D-96).

2026-08-25 correction pass 15: D-94 corrects the PR-6 commander helper, which never ranked commanders by score. Run 10 passes both bars. The weak-commander-pool trigger is wired, and PR-7 has no open item left but the owner's scoring.

2026-08-25 correction pass 14: run 9 adds the decline (D-93) and passes both bars. `State.Skip` was the last dead method in the package. The M-5 sheet holds 52 items, which passes the 50 the rubric asks for.

2026-08-25 correction pass 13: runs 7 and 8 found the root cause under D-83. A schema enum suppressed the format field, and the model answered "unknown" for a message that named the format outright (D-92). Both enum fields are free strings now.

2026-08-25 correction pass 12: gate runs 5 and 6. The runs meet the catalog bar at 30 of 30, and a second bar now applies. D-91 fails a session that calls itself complete with a needed slot unanswered, which run 6 did five times. D-83 reached its third and final form, and D-90 makes a field the only record of an answer.

2026-08-25 correction pass 11: gate run 3 measured the fault checklist and found a regression in D-83. Changes: the PR-7 entry. D-86 corrects a side effect of D-79. D-87 deletes the acquisition row on the owner's challenge: the app can not act on a delivery date.

2026-08-25 correction pass 10: the PR-7 gate ran twice more. Run 2 scored 23 of 30, and it proved two fixes and exposed two deeper defects. Decisions D-83 to D-85 record them. The proto gained `Question.catalog_text`, which M-5 needs in the UI.

2026-08-25 correction pass 9: the PR-7 live gate ran and failed at 21 of 30. Changes: the PR-7 entry. Decisions D-78 to D-82 record the five defects the run exposed and their fixes. The worst one offered commanders outside the deck color identity.

2026-08-24 correction pass 8: PR-7 completed except for the live gate run. Changes: the PR-7 entry. Decisions D-70 to D-74 record the three items of the second live run, the "none" rule, and the session store. The service wiring found four more dead rows.

2026-08-24 correction pass 7: PR-7 built in two halves. Changes: the PR-7 entry, PR-6 gains two helpers for PR-7, and D-69 records the gap-score call shape and its first threshold. One live conversation measured the cost of a turn.

2026-08-24 correction pass 6: PR-6 merged (#11) after a second gate run. Changes: PR-6 status and text, M-5 gains the OQ-19 rubric (D-66), section 9 open questions, and three engine defects recorded in the PR-6 entry. PR-7 gains the pool-question timing (D-67) and the slot freeze (D-68), both from eight dogfood conversations. The corpus question catalog changed the same day.

2026-08-24 correction pass 5 (full audit, `docs/audit-2026-08-24.md`): every PR status set to its merge state (#1 to #9). Register rows F-21 to F-25 added. PR-1b (contract amendment) added before PR-6. PR-3 added to section 8. Decisions D-42 to D-60 recorded. The fixes ship on branch `audit-fixes`.

2026-08-24 correction pass 4 (owner directive): the collection is optional (D-37, amends D-2). Changes: thesis fact 3, guardrail 5, PR-6, PR-7, PR-8, PR-11, PR-12. No proto change: `PoolRule` and an optional `collection_id` existed since PR-1.

2026-08-23 correction pass 3: OQ-13 to OQ-17 answered (D-26 to D-30). Changes: M-5 added (manual scoring lane for invented questions), I-1 rewritten as the stale-deck banner and rerun, I-2 threshold fixed, section 9 updated.

2026-08-23 correction pass 2: the owner answered OQ-1 to OQ-12 (D-15 to D-25). Changes in this pass: F-4 resolved, F-14 rewritten, F-16 and F-17 added, PR-4 storage decided. Also: PR-7 gains the gap score and M-4, PR-8 gains the deck summary, I-2 has a price spec, and the GCP project ids are set.

House rule (from connector-syncer): no PR, branch name, commit message, comment, or other artifact carries AI-attribution text.

---

## 1. Thesis

A deck builder is useful only when three things are true at the same time. The deck is legal on the day the user asks. The deck fits what the user owns and what the user wants. The deck works as a plan, not as a pile of good cards. A language model can do the third thing well. It can not do the first two things reliably without deterministic checks around it.

The system is therefore built as a **thin agent over a strict engine**. The engine owns card data, legality, ownership, and validation. The agent owns the conversation, the plan, and the card choices inside the engine's limits. Every card the model names passes through the engine before the user sees it. This is the same shape connector-syncer uses for citations: the model proposes, the code verifies, and hallucinations die at the boundary.

We sequence the program so that each layer is testable before the next one exists. **Data and legality** come first (deterministic, cheap to test). Then the **agent loop** (measurable against the engine). Then the **UI** (thin over a streaming API). Then **meta and quality** (the expensive, judgment-heavy part). Each phase waits for the one before it.

> *In plain English:* the AI is good at the creative question, such as a fun lifegain deck. It is bad at the exact question, such as a ban check. So the code answers the boring exact questions. The AI answers the creative ones. The AI never gets the last word on a card. The code checks every card before the user sees it.

## 2. Lessons learned (carried in from connector-syncer, and from this research)

1. **One concern per PR.** The reference roadmap lost a full attempt (#724) to a bundled change.
2. **Verify platform claims before you build on them.** Lesson 7 in the reference doc. Here, we checked Scryfall rate limits, ManaBox column names, and emulator behavior against the real thing, not the docs alone.
3. **Model churn is a standing tax.** The role-to-model layer (D-1) and a bake-off protocol are the only durable answer. Record the resolved model in every eval row.
4. **Normalize model output against ground truth after every call.** A card name that does not exist is our hallucinated citation.
5. **Evaluate on the payload you change.** A prompt tested only on Commander tells nothing about Standard.
6. **Commit the evidence.** A/B scripts and golden sets live in the repo.
7. **Instrument cost on day one.** Every cost claim in the reference doc was an estimate until M-1 landed there.
8. **The user's words are not a spec.** "Anything goes" has no fixed meaning (D-3). The agent asks. This is a product principle, not a fallback.
9. **Facts have dates.** The ban list changed four times in 2026. Store facts with a verification date and a source.
10. **A structured-output schema is a provider contract.** Bracket gate run 1 (2026-09-02) lost every judge call to a bounded integer the Anthropic schema refuses. Make one call before a paid run.
11. **A judge's reasons are for a reader.** The same run's judge named a combo Commander Spellbook does not list. A bar reads the judge's number and never its prose.

> *In plain English:* these are the mistakes the sister project already paid for. We do not pay for them twice.

## 3. System map

| Component | Language | Owns | Reads | Writes | Sensitivity |
|---|---|---|---|---|---|
| `cards` service (card database) | Go | Scryfall snapshot, legalities, Oracle tags, images URIs | Scryfall bulk daily | Firestore `cards/`, GCS snapshot | High - every legality answer comes from here |
| `collections` service | Go | ManaBox import, ownership counts per Oracle ID | User CSV upload | Firestore `users/{uid}/collections/` | High - PII-adjacent, user data |
| `rules` engine (library) | Go | Format rules, deck validation, bracket rules, color identity | `cards` | none | Total - the last gate before the user |
| `profile` (bracket profile, library) | Go | The bands per bracket, the feature vector, the goldfish simulation, the content check (PR-14A) | `cards`, `rules`, Commander Spellbook | none | High - it says what a bracket means |
| `agent` service | Go | Turn-based chat, question workflow, deck generation, LLM role layer | `cards`, `collections`, `rules`, `meta` | Firestore `users/{uid}/sessions/`, `decks/` | High - the product |
| `meta` service | Go | Deck quality model per format: the labeled lists, the fitted weights, the scorer (PR-14) | MTGO decklists, the Topdeck.gg API, MTGTop8, the cEDH database, EDHREC (D-5, D-417) | GCS raw pages and normalized lists, Firestore `meta/` for the weights | Medium - advisory input to the agent, and the code keeps the final say |
| `worker` | Go | Scheduled jobs: Scryfall refresh, meta refresh, ban-list watch | Cloud Scheduler, Cloud Tasks | see above | Medium |
| `web` (UI) | TypeScript, React, Vite | Chat, deck view with card art, collection upload, export, the deck library, the binder, the share page (Phase 3B) | Connect-RPC API | none | Medium |
| `proto` | Protobuf | The one contract between Go and TypeScript | - | generated code, committed | High - a schema change is a cross-stack change |
| `eval` | Go + fixtures | Golden decks, deterministic checks, judge runs | all services in-process | BigQuery `evals.*` | Medium |

Three structural facts drive the plan:

- **The card database is small.** About 38,600 Oracle cards and 34,500 playable ones. It fits in memory in Go. We do not need a vector store to find cards. Scryfall Oracle tags (4,522 community tags, for example `lifegain` with 3,374 cards) plus keywords and type lines give a structured theme index. A vector index is a later option (I-3), not a foundation.
- **Legality is a lookup, not a judgment.** Scryfall updates `legalities` within a day of each announcement. The engine reads them. No prompt carries a ban list. This makes "legal as of the query date" a data-freshness problem, which is measurable (M-2).
- **Ownership is a differentiator, not a requirement (D-37, 2026-08-24).** With a collection, every suggestion is one the user can play tonight. Without one, or with the library toggle off, the app builds a truly optimized deck from the whole legal pool. The pool mode is a first-class slot (`PoolRule`): owned-first, owned-only, any-card. The import must be robust to ManaBox column variance (F-2).

> *In plain English:* the whole card list of Magic fits in memory. The rules for "is this legal?" are a table lookup. The one thing no other tool has is the user's actual binder. So we build around the binder.

## 4. Cost model (what we expect, what we do not know)

- **LLM.** One deck-build session runs 2 to 4 question turns on the small model, at about 2k tokens each. It then runs 1 to 3 generation turns on the strong model, at about 15k input with the candidate card list and 3k output. Estimate: under $0.10 per session on 2026 list prices. Unknown until M-1 measures it. Prompt caching of the format rules and the candidate list cuts the input cost. The role layer must expose the provider's caching knob (D-1, D-21).
- **Card data.** Scryfall bulk: 24.5 MB compressed per day for Oracle cards, 77.5 MB for all English printings. Free. Images hotlinked (D-6), zero storage. GCS: one snapshot per day, about 100 MB, cheap lifecycle to 30 days.
- **Meta data.** The source terms passed the legal check (D-5). MTGO decklists are official and free, and one event page is about 330 KB. A year holds about 3,600 events, so the raw pages are about 1.2 GB, and the normalized lists are a few megabytes. A fit runs in seconds in Go and costs no LLM call.
- **Firestore.** Per user: one collection doc set (PR-4 decided one gzip document per collection, about 500 KB for a 5,000-card binder, D-16), sessions, decks. Low.
- **Cloud Run.** Two services plus a worker, scale to zero. Low until users exist.
- **Eval.** Deterministic checks are free. Judge runs cost per deck. Cap per run as connector-syncer does ($5 cap in its bake-off).
- **Bracket profile (PR-14A, 2026-09-02).** Commander Spellbook is free at 90 calls a minute, two calls per build. The profile's repair passes raise a build from about one model call to two. Deck gate run 12 cost $2.24 against $1.46 for run 11.
- **Unknowns to measure first:** tokens per session (M-1), Scryfall refresh lag after an announcement (M-2), ManaBox import failure rate on real files (M-3).

> *In plain English:* the AI is the only real cost, and the target is cents per deck. Card data is free. Images are free because Scryfall lets us link to them. We will measure instead of guess.

## 5. Defect and finding register

Status: ✅ resolved · 🔧 planned (item listed) · 🅿 parked · ⏸ out of scope · ⚠ constraint · ❓ needs owner input.

| # | Finding | Status |
|---|---|---|
| F-1 | **Ban lists drift fast.** Four B&R announcements in 2026 so far (03-23, 05-18, 06-29, 08-10). Next 2026-10-12. Commander changed 2026-02-09 with a new category, "banned as a companion". Any cached legality older than one day can be wrong on announcement day. | ✅ PR-3 (#6). M-2 logs the lag. The audit replaced the calendar test with a legality diff (F-23, D-47). |
| F-2 | **ManaBox CSV columns vary.** The official guide does not list the columns. The verified set is 18 columns, from the owner's real export of 2026-08-24. A single-list export can drop the binder columns. Column order and presence can change with app versions. | ✅ PR-4 (#7): header-driven import, Scryfall ID first, unknown columns logged once. |
| F-3 | **Scryfall API rate limits are hard.** 2 requests per second on `/cards/named`, `/cards/search`, `/cards/random`, `/cards/collection`. 10 per minute on `/cards/manifest`. 10 per second elsewhere (verified 2026-08-24). A 429 blocks for 30 seconds. Repeated overload gets a ban. Bulk files have no limit. | ⚠ binds PR-2: all card lookups go to the local snapshot. The live API is for single-card fallback only, behind a client-side limiter. |
| F-4 | **Aggregator terms of use unknown.** MTGGoldfish, MTGTop8, Aetherhub, and EDHREC have no public API and their terms were unchecked. | ✅ 2026-08-23: the owner confirmed the legal check passed (D-5). All five sources may be used. PR-14 still starts with MTGO because it is the only structured source. |
| F-5 | **Oracle tags are community data.** Scryfall Tagger tags are volunteer-made. Coverage is uneven. `lifegain` is rich (3,374 cards). Niche themes may have few tags. Weights are `median` style, not scores. | ⚠ binds PR-6: tags seed the candidate list. They never gate a card. Keywords and type lines are the second signal. The model is the third. |
| F-31 | **A land upgrade rebuilt one basic land.** Session `vAvg4eteJhmuPEuJwBul`, 2026-09-01: the owner asked for better lands in place of the basics on a three-color Commander deck. The agent asked what kind, the owner answered a mix, and the rebuild moved one Plains to one Island. The shortlist held 40 nonbasic lands, among them Command Tower, the three shock lands of the colors, the fetch lands, and Cavern of Souls. The generator's revision rule asks for the smallest change, the brief named no count, and no check reads a free-text change. Revise gate runs 1 to 4 stopped at the question on both land rows, so the turn after the answer never ran under a bar. | 🔧 fixed on `pr-19` (D-448, D-449), revise gate run 5 due. |
| F-32 | **The theme fills the land bucket.** Revise gate run 5, 2026-09-02: the brief counted 12 basic lands to replace on the Karlov lifegain deck, and the rebuild added one. The land bucket holds 40 lands ranked by theme text first, and "lifegain" sits in the text of dozens of lands. All 40 slots went to tapped lands that gain life, and Godless Shrine, Isolated Chapel, Caves of Koilos, Marsh Flats, and Command Tower were absent. The model was asked for untapped duals and had none to name. The Éowyn deck of F-31 had its staples because "Human" matches few lands. | 🔧 fixed on `pr-19` (D-450), revise gate run 6 and deck gate run 11 due. |
| F-33 | **The mana base is model variance.** Deck gate runs 10 and 11, 2026-09-02, read for the lands of every deck. The lifegain Commander prompt got 12 nonbasic lands in run 10, 0 in run 11, and 30 in the base of revise gate run 7, on the same theme and pool. Five two-color Commander decks of run 11 hold 36 basics and no fixing at all, and one holds 3 basics and 33 nonbasic lands. A Standard deck holds 16 tapped lands of 24, and a tournament Modern deck 8 of 24. No rule states what a mana base should be, so the model decides each time. D-450 cut the colorless lands of the 13 shared Commander decks from 91 to 33 and left the tapped count alone. | 🔧 PR-14A built the land band, the tapped cap by power, and the Karsten sources on 2026-09-02 (D-463). Deck gate run 12 read them: every land count sat in band, and the sources band caught two decks short of a color. The share of nonbasic lands still swings, 0 to 36 on the same prompt, because no band reads the composition. Open: a basic-to-nonbasic band, on evidence. |
| F-34 | **The bracket cut and the precon share fought, and a failed repair replaced a legal deck.** Deck gate run 14, 2026-09-04: the Goblin Storm upgrade of prompt 17 came back with no cards and a `deck_size` block. The bracket cut of D-468 drops a forbidden shortlist card and keeps only the commanders and the locked cards, so precon cards left the pool. The share rule of D-218 still asked for 67 of 78 names. The repair turn read fewer marked names than the rule demanded, answered no cards, and the builder replaced the legal first deck with the empty one. Run 13b passed the same prompt because its repair did not give up. | ✅ resolved 2026-09-04 (#63, D-509): the share counts the precon names the pool holds, and a repair with a block or a miss after a clean pass leaves the clean pass in place with the note `repair_kept_earlier`. Two tests pin it, and deck gate run 14b passed the prompt with the fix. |
| F-35 | **The whole-precon check counts basic lands, and the export lacks six of them.** The ManaBox deck binder of Avengers Assemble in the export of 2026-08-24 holds 84 of the 90 printings of the product. The six absent cards are basic lands of three printings, Plains 288, Island 290, and Mountain 294 of MSH. Under D-408 that binder owns no precon, so "not from my precons" excludes nothing for it, and a named product gets the partial note of D-497. The exclusion never removes a basic land (D-37), so the six cards change no deck. | ✅ fixed 2026-09-05 (D-523, D-532): the check reads the nonbasic printings alone, in the chat and in the deck gate. The Avengers fixture is the true export again, with no hand-added row, and the dry run of prompt 25 still excludes 57 cards with 27 spare. |
| F-36 | **The classifier replaces a named theme with a vague later phrase.** Question gate run 34, 2026-09-04: "poison in a tournament" opens with "An infect deck for a Modern tournament", and the second message says "Modern. The best deck under budget." The theme slot ended as `the best deck under budget`, and the word infect was gone. The reader never changed the theme. The deck would be built on the phrase and not on infect. | 🔧 open (2026-09-04): run 35 read `infect`, so the classifier keeps the theme in one read of two, and the miss is not deterministic. The expectation stays at `infect`, and every run reads it. The fix is in the classifier: a phrase with no theme word does not replace a named theme. |
| F-37 | **A theme the exclusion starves builds a different deck, and the summary hides it.** Deck gate runs 15 and 16, prompt 25: the reader asks for Avengers superheroes from the library first, with no card of the Avengers Assemble precon. The Avengers cards of that library live in the precon, so the pool held 152 names and few superheroes. The build answered with a mono-black Syr Konrad list, the plan judge read theme_fit no twice, and the summary never said that the theme could not be served. The partial note of D-497 covers a product owned in part, not a theme the exclusion empties. | 🔧 open (2026-09-04): the build says when the pool holds too few cards of the theme after an exclusion, and offers the whole pool or another theme. The plan judge found it, which is the job of PR-15. |
| F-38 | **The plan judge writes "placeholder" as the reason of its last field.** Deck gate run 16, 2026-09-04: the judge on Opus 5 wrote the one word "placeholder" as the reason of `summary_honest` in 10 of 25 decks, and never on another field. The grade landed each time, and a bar reads the number (lesson 11), but the prose of that field is worthless in those decks. | ✅ fixed 2026-09-04 (#65, D-515): the schema asks for the reason before the grade in each field, the document counts the empty reasons, `plan_reasons_empty` is a row per deck, and `PlanRubricVersion` reads 2. The next deck gate run proves it. |
| F-39 | **The plan judge grades legality from its own card knowledge, and gets it wrong.** Deck gate run 16, 2026-09-04: on deck 2 the judge read Denethor, Ruling Steward as mono-black and half the deck outside its identity, and Scryfall reads the card white-black. On deck 10 it read Boros Charm, Inspiring Vantage, and Sacred Foundry as outside Standard and a four-of as black, and Scryfall reads all three legal in Standard and Redcap Gutter-Dweller red, on 2026-09-04. The deterministic checks passed both decks against the snapshot. Both decks read `useful_as_built` no on false facts. | 🔧 open in part (2026-09-04, D-515): the instructions tell the judge that code checked the legality, the identity, and the counts, and that its card pool can be older than the data. The deck text it reads holds no mana cost and no color, and the fix that grounds it with the card facts waits on the owner. |
| F-40 | **A bonus sheet stood as a set candidate.** 2026-09-04, the chat: "Build the best possible Marvel-only deck" asked which set, Marvel Super Heroes or Marvel Universe. Marvel Universe is a Scryfall set of type masterpiece, 94 reprints of existing cards with Marvel art from 2025-09-26, with no parent set. The resolver of D-376 reads every parentless set as a candidate, so the sheet stood beside the expansion. | ✅ fixed 2026-09-04 (D-518): a masterpiece set never leads a family when a phrase matches several sets. An exact code or an exact name still resolves it. |
| F-41 | **A question joined its choices with "and".** The set question read "Which one: Marvel Super Heroes and Marvel Universe?", and the precon question has the same shape. The row joined the options with the list joiner of the locked cards. | ✅ fixed 2026-09-04 (D-518): the two option lists join with "or", and the lists of cards and sets keep "and". |
| F-42 | **A partial run stood as the run of record.** Question gate run 36, 2026-09-04, ran probe 109 alone under `-only`. The document read FAIL on the catalog-only bar and the gate size bar, 0 of 0 against 25 and 30, which one conversation can never meet. `make eval-check` then took run 36 as the newest run of the suite and read FAIL against baseline run 35, with 107 items gone. The probe itself passed its sets expectation. | ✅ fixed 2026-09-04 (D-526): the run header records the `-only` flag. A partial question gate run reads the item bars alone, and its two whole-set rows are information. The check lists a partial run under its suite and compares none. Run 37 reads PASS as a partial run, and the run 36 header carries the flag by hand (D-527). |
| F-43 | **A corrected commander name joined the misspelled one.** Question gate run 38, conversation 82: "Build around Atraxa, Praetor's Voice", then "I meant Atraxa, Praetors' Voice". The state merged the names, and both led the deck. | ✅ fixed 2026-09-05 (D-535, D-538): a commander name the index knows replaces an earlier one it does not know, and the misspelled name leaves the named cards too. Run 41 then showed both spellings in one turn, the classifier's repeat, so an unknown name never joins a known commander either. |
| F-44 | **"Colorless" answered the color question with all five colors.** Run 38, conversation 74: the color question was out, the user wrote "An artifact deck, colorless", and the classifier answered W, U, B, R, and G. The colorless rule of D-165 saw a filled slot and stood aside. | ✅ fixed 2026-09-05 (D-535, D-536): a colorless message drops the classifier's colors, and the rule closes the slot with none. Run 39 then showed the second half: "Bracket 3." came back with all five colors over the closed slot. A closed color slot changes only on a message that names a color now, the D-125 rule for colors. The expectation word "none" reads that end. |
| F-45 | **The user's words for the budget scope lost to an inferred scope.** Run 38, conversation 70: the classifier inferred the whole deck on an earlier turn, and "The 100 caps the cards I buy" could not change it. The scope rule returned on any filled scope. | ✅ fixed 2026-09-05 (D-535): the buy-list words and the whole-deck words settle the scope on any turn, and only the inference waits for an empty scope. |
| F-46 | **A superlative phrase replaced a named theme.** Run 38, conversation 25: "An infect deck for a Modern tournament", then "Modern. The best deck under budget." The second message answers the budget question, and the prompt reads a superlative phrase as a theme, so infect became "the best deck under budget". Run 35 kept infect: the read is noise the code must absorb. Conversation 77 shows a kin: "a Background commander pair" sat in the theme slot. | ✅ fixed 2026-09-05 (D-535) for the superlative: it is a theme for a deck with none, and it never replaces a named theme. The structural phrase of 77 stays an observation, and its theme expectation is gone. |
| F-47 | **The classifier answered the budget question with a number the message never held.** Question gate run 40, conversation 25: the budget question was out, the user wrote "Modern. The best deck under budget.", and the classifier reported a budget of 2. The session called itself complete, and the turn with "400 dollars" never went out. Run 39 read the same conversation right. | ✅ fixed 2026-09-05 (D-537): a budget applies only when the message names it. A message with digits must hold the number, a message in number words stands, and a message with neither names no budget. |
| F-48 | **The judge lane document prints a JSON escape in place of an em dash.** Both judge lanes of bracket gate run 1 hold the six-character escape `\u2014` where the judge wrote an em dash: two in the lane of 2026-09-02, three in the lane of 2026-09-05. The reason text keeps an escape the renderer never decodes. | 🔧 open (2026-09-05): cosmetic, in the gate document alone. |
| F-30 | **No signal of deck quality exists.** The pool ranks on theme fit and EDHREC popularity, and the bracket drops Game Changers under bracket 3 and nothing else. A bracket 5 request got the three most popular legends whose text held "you" and "can" (session t8o1nGGquK6UdTQkfY3V, D-411, 2026-09-01). | 🔧 PR-14B built the scorer on 2026-09-02 (D-470 to D-478). Its gate waits on the first meta read and the Topdeck.gg key (OQ-54). |
| F-6 | **No Cloud Tasks emulator.** Local mode can not run real Cloud Tasks. | ✅ PR-0c (#3): a `Dispatcher` interface with a local in-process implementation. |
| F-7 | **Docker absent on the dev machine.** | ✅ PR-0b (#2): Docker 29.7.2 installed. Native `make dev` does not need it. |
| F-8 | **The mtgcommander.net banned-list page reads "last updated September 2024."** It does not show the 2026-02-09 changes. It is not a reliable source for the current list. | ✅ Scryfall `legalities.commander` is the source of truth. The page is for philosophy text only. |
| F-9 | **Double-faced and split cards have no top-level `image_uris`.** Images live in `card_faces[]`. Layouts `transform`, `modal_dfc`, `split`, `adventure` (about 800 Oracle cards). | ✅ PR-2 (#5): the card model normalizes faces. The UI shows both faces (PR-12). |
| F-10 | **"Anything goes" and other user phrases are ambiguous.** The owner confirmed this is by design (D-3). | ✅ product principle. PR-7's question catalog handles it. |
| F-11 | **Commander brackets are "beta" and change.** The 2026-02-09 update changed the Game Changers list. The bracket rules are prose, not data. | ✅ PR-5 (#8) encodes brackets as data with a version date. The `game_changer` flag comes from Scryfall. PR-14A (2026-09-02) encodes the prose rules as data too, and Commander Spellbook checks each built deck against them (D-462). The info finding left. |
| F-12 | **Legality is per Oracle card, but ownership is per printing.** A user may own a printing that is not legal in a format where the card is legal (for example a gold-bordered or Alchemy-rebalanced version). Scryfall marks these on the printing. | ✅ PR-4 (#7) keeps the printing id. PR-5 checks legality on the Oracle card. The printing exception is an info finding when the printing data carries it (audit fix, 2026-08-24). |
| F-13 | **Model output can name a card that exists but is not the card meant.** Example: "Ajani's Pridemate" versus "Ajani's Welcome". Fuzzy matching hides this. | ✅ done in PR-8 (D-222): the shortlist is the whole contract, the normalizer matches an exact name only, and a near name is reported to the repair turn and never substituted. Deck gate runs 4, 5, and 6 reached the user with no invented name. Earlier: ⚠ binds PR-8: exact name match only, with the model asked for exact names. Fuzzy match is a suggestion to the user, never a silent substitution. |
| F-14 | **Variance versus determinism.** The owner wants variance between decks. A first answer to OQ-4 asked for identical output on identical input. A second answer the same day withdrew that: random variance stays (D-18). Note for the record: an LLM is not deterministic even at temperature 0, so identical output was never a guarantee. | ✅ resolved by D-18. Updated 2026-08-28: PR-9 left the MVP (D-256), the seed and the plan variant left the contract, and the deck store keeps every build (D-245). |
| F-15 | **Standard has no rotation in 2026.** Rotation moves to the first set of 2027. Any hardcoded "September rotation" logic is wrong. | ✅ no rotation logic in code. Scryfall legalities carry it. |
| F-16 | **Scryfall prices have no condition tiers.** `prices.usd`, `usd_foil`, `usd_etched` are TCGplayer near-mint market estimates, updated once per day. The owner asked for lightly-played prices (OQ-3). No free source gives them. | ✅ D-17: show the NM estimate with a 7-day rolling average and outlier rejection, labeled as such. A condition-tiered source is a later option. |
| F-17 | **A fixed question catalog can not cover every prompt.** The owner wants catalog questions first, model-invented questions when needed, and a metric that says which case applies (D-25). Without the metric, the agent either asks nothing new or bypasses the catalog. | ✅ PR-7 (gap score) merged 2026-08-26. 🔧 M-4 (catalog coverage metric) and PR-15 (catalog-change proposals from evals) stay open. |
| F-18 | **Scryfall legalities can not say "banned as a companion".** The 2026-02-09 Commander update unbanned Lutri, the Spellchaser but banned it as a companion. The Scryfall commander legality reads "legal". The open legalities map (guardrail 2) inherits this blind spot. | ⚠ binds PR-5: the rules engine owns the companion check. `Deck.companion_oracle_id` exists so the check has a target. Found in the 2026-08-24 proto re-pass. |
| F-19 | **The Go storage SDK's default download path 404s on fake-gcs-server.** The SDK reads objects through the XML API with percent-encoded names. The fake-gcs filesystem backend serves only the JSON paths for names with slashes. Listing works, reads fail. Found 2026-08-24 in the PR-2 smoke test. | ✅ fixed: every storage client passes `storage.WithJSONReads()`. JSON reads work on fake-gcs and on real GCS. |
| F-20 | **A snapshot version was listable before its files finished uploading.** The API loaded a mid-download snapshot and failed with "object doesn't exist". This is the reference project's F15 lesson: completion must imply artifacts. | ✅ fixed: the store writes a `complete` marker last. `LatestVersion` returns only marked versions. A test pins it. |
| F-21 | **A ManaBox token row resolved to the real card.** Row 186 of the owner's export is the Bloomburrow token `Pawpatch Recruit`. The resolver dropped tokens from the index, then fell through to the name lookup, which found the real creature. The PR-4 "2,548 of 2,548" gate counted it as resolved. Found in the 2026-08-24 audit. | ✅ audit fix: the index remembers dropped printings. A token, emblem, or art-card row reports `NOT_PLAYABLE` (D-44). |
| F-22 | **Five commander-eligibility gaps.** `Partner—[text]` variants collapsed into plain Partner. A lone Background passed. "Up to seven" and "up to nine" cards were blocked. A double-faced card qualified on its back face. Vehicles and Spacecraft (CR 903.3, checked against the 2026-08-07 text, current file 2026-08-19, D-272) were refused. Found in the 2026-08-24 audit against the Comprehensive Rules. | ✅ audit fix: `partner_text` and `max_copies_override` on the card, front-face rule, Vehicle and Spacecraft support (D-49, D-50). The golden gate grew to 41 good and 53 bad. |
| F-23 | **The announcement-day model used the calendar, not the data.** A snapshot from 09:00 UTC on announcement day counted as covered, hours before Wizards posted. The fast path never ran. | ✅ audit fix: coverage is a legality diff between snapshots (D-47). The calendar only sets the poll cadence. |
| F-24 | **The LLM layer under-counted the judge.** Anthropic thinking tokens were dropped, cache writes were priced at 1x instead of 1.25x, and `LLM_JUDGE_PROVIDER=fake` passed under `LLM_REQUIRE_KEYS=1`. Found in the 2026-08-24 audit of PR-10. | ✅ audit fix: thinking tokens counted, `cache_write` price column, keys required by default (D-51). |
| F-25 | **The proto lacked fields PR-6 to PR-9 need.** No upgrade list, no slot state, no question id, no structured answer, no seed override, no `Usage`, no per-face artist. | ✅ PR-1b (audit branch): all fields added in one contract amendment (D-46). `buf breaking` guards it from now on. |
| F-26 | **The phrasing role invents claims about the game, and no gate catches them.** Gate run 14 of 2026-08-26 passed the gate with zero linter findings and told one user two false things. It said "Grist, the Hunger Tide can not lead a deck", which the rules contradict (Scryfall ruling, 2021-06-18). It asked "Do you want to use any colors beyond Grist's color identity?", which the rules allow no answer to. The catalog rows say neither. The ask role added both, and the ask prompt already said to state no fact about the game. | ✅ answered for PR-8 by the judge lane (D-229): the judge role reads every deck summary for a rule of the game and for the truth of it, on another provider than the generator, and a false rule fails the deck gate. Deck gate run 6 found none in 16 summaries, at $0.0034 a deck. The deterministic net reads the shape of a claim and never its truth, which is why the judge decides (D-224). Earlier: 🔧 partly fixed: D-140 silences the commander row for a legendary card the engine can not confirm, and D-144 puts the color-identity rule in the prompt and the shape in the linter. ⚠ binds PR-8: the generate role writes a deck summary in prose, and the same failure has more room there. The eval lane found both, and the deterministic linter found neither. |
| F-27 | **A message after a build is dropped, and the deck is rebuilt from the first message.** Session `eIrL12hRY2YNTTCo3iS4`, 2026-08-28: the user wrote "Replace some lands with better options if possible. Also tune the mana curve lower - no 6 or 7 mana cards needed". The agent sent no reply, `plan()` read turn 1 and the slots only, and the generator built a second deck 2.5 minutes after the first with 24 Plains and the same four 6- and 7-mana cards. Three cards changed, all by variance (D-18). The turn cost a full build and answered nothing. | ✅ PR-12B merged 2026-08-29 (#41). |
| F-28 | **The UI is a live-test UI, not a product.** PR-11 to PR-13 built four screens on plain Tailwind for the owner's browser test (D-273). No design system, one theme, no navigation on a phone, no rename or delete of any object, no card detail, no share, and `firebase/auth` on every route. Found 2026-08-29 when the owner asked for a user-facing UI (D-310). Binds PR-16 to PR-23. | 🔧 Phase 3B |
| F-29 | **A set is not a constraint the app can apply.** Session `WJbs7FP2csZCULi4SVJu`, 2026-08-31: the owner asked for a tier-5 Commander deck from the Hobbit set. `candidates.Request` carries the format, the colors, the theme, the commander ids, the pool rule, the owned counts, and the bracket. It carries no set. The words reached the theme, the theme matched Scryfall tags, and no card was ever tested for its set. The deck could hold any card of any set. The snapshot of 2026-08-31 holds 988 paper sets over 34,599 playable Oracle cards. | 🔧 PR-17B |

> *In plain English:* these are the traps we found before we wrote code. The biggest ones: ban lists change every few weeks. The collection file format has no documentation. The AI can name a card that sounds right but is not. Each one has a planned fix or a rule that prevents it.

## 6. Guardrails (the safety contract for every PR)

1. **No card reaches the user before the rules engine checks it.** The engine validates every generated list for size, copies, legality on the query date, color identity, bracket, and ownership. A failed check blocks the response or marks the card, never silently drops it.
2. **No ban list, rotation date, or Game Changers list in any prompt or code constant.** Legality comes from the card database, which comes from Scryfall daily. Prompts can say "the engine will check legality".
3. **No model id at a call site.** All models come from the role layer (D-1). CI warns on a default change, as in connector-syncer.
4. **Exact card names only.** The model returns exact Oracle names. The normalizer does an exact match. Anything else becomes a user-visible suggestion, never a substitution (F-13).
5. **The pool mode is explicit, and ownership is always visible (D-2, D-37).** Every deck records its pool mode. When the session has a collection, every card carries an `owned` flag with the count, in every mode. Owned-first is the default with a library, any-card without one. The engine never silently narrows or widens the pool.
6. **The agent asks before it assumes** on format, power level, and house rules. Max three questions per turn. A default applies only when the user says "you decide" (D-3).
7. **Attribution on every image.** The full card image shows its printed artist and copyright. The app does not crop or alter an image, and no paywall sits in front of card data (Scryfall terms, D-6, D-291). Only an art crop needs a separate credit line, and the app never shows one.
8. **Generated proto code is committed and CI diffs it.** No hand edits. Pinned buf and plugin versions.
9. **Local mode has no cloud dependency.** Every service starts with emulators or fakes (D-9). A new cloud dependency must ship with its local fake in the same PR.
10. **One concern per PR. Evidence committed.** Golden decks and A/B outputs live in the repo.
11. **No raw user prompts in analytics.** Hashes and ids only, as in connector-syncer's event registry.
12. **Dated facts.** Every rules or format fact in the corpus carries a source and a verification date.
13. **The allowlist gates every deployed request, and a public page carries no user data (D-314, D-315).** On GCP the interceptor refuses a uid whose email is not on the allowlist. The shared deck message holds no user field, and a test proves it.

> *In plain English:* thirteen promises every change must keep. The most important: the code, not the AI, has the final say on every card. And the app never quietly swaps a card the AI got wrong for one it guessed.

---

## 7. Roadmap

Ids: PR-# code, M-# measurement, I-# integration, D-# decisions (in `decisions.md`). Each entry has a gate and ends with a plain-English paragraph.

### Phase 0 - Foundations (no product code)

**PR-0a: Monorepo scaffold.** ✅ merged 2026-08-24 (#1, branch `pr-0a`). Deviations from the plan, recorded in D-35: Vite 7 instead of 8, dev port 5180, buf built into `.bin/` from a `go tool` directive. Node moved to 22.12 LTS in the audit (D-52).

Layout: `proto/` (buf module), `go/` (Go workspace with `cmd/api`, `cmd/worker`, `internal/cards`, `internal/collections`, `internal/rules`, `internal/agent`, `internal/meta`) (planned names, 2026-08-23, the code uses `internal/agentsvc`, and nothing built `internal/meta`), `web/` (pnpm workspace: `apps/web`, `packages/api-client` for generated TypeScript), `docs/`, `.claude/`. Makefile as the single entry point: `proto`, `lint`, `test`, `test-repeat`, `cover`, `dev`, `dev-seed`. Pinned versions: Go, buf, protoc-gen-go, protoc-gen-connect-go, protoc-gen-es, pnpm, Node, golangci-lint. CI: `verify:*` matrix with a fan-in job, a proto-diff gate, and a `buf breaking` gate (audit). The plan had path filters, and D-56 struck them: a skipped required check blocks a merge. AGENTS.md with the commands and never-edit rules.

Gate: `make dev` starts an empty API and an empty UI.

> *In plain English:* the empty house with plumbing. One folder for the shared contract, one for Go, one for the web app. One command starts everything, and the checks that stop bad changes are wired before there is anything to check.

**PR-0b: Developer machine setup, including Docker (D-10).** ✅ merged 2026-08-24 (#2). The owner installed Docker 29.7.2, and `make doctor` reports all ok (12 checks).
A `docs/setup.md` procedure: install Homebrew, Git, Go, Node and pnpm via corepack, the firebase CLI, Java 17, Docker Desktop, and gcloud. A `make doctor` target checks each tool against the pinned version and prints the fix command. buf comes from `go/go.mod`, not from a separate install.
Docker is used for the Compose file (PR-0c) and for local Cloud Run parity. Gate: `make doctor` passes on the owner's machine.
> *In plain English:* a checklist to set up a laptop, and a command that tells you which tools are absent. The owner asked for the Docker install to be a tracked step, so it is one.

**PR-0c: Local stack (D-9).** ✅ merged 2026-08-24 (#3). Both variants verified: native (`make dev`, all services up in 8 seconds, clean teardown) and containers (`make dev-docker`, Compose).

New in this PR is the port map D-36. The Wallabee stack owns 8080, 8181, 4000, and 5173 on this machine. This PR also adds `internal/dispatch` (the Cloud Tasks stand-in, F-6), and `internal/llm` with the `Fake` provider.

`firebase.json` with Firestore and Auth emulators. `fake-gcs-server` for storage. A `Dispatcher` interface with a local in-process implementation (F-6). A `fake` LLM provider with fixture responses.

`make dev` runs all of it under one process supervisor. A Compose file gives the same stack in containers once Docker exists. GCP projects are `mtg-dev` and `mtg-prod` (D-24). No domain and no hosting yet. Local testing has priority. Gate: a developer with no GCP credentials runs the full stack and the UI loads.

Container note: the firebase emulator binds 127.0.0.1 from `firebase.json`. The emulator image rewrites the host to 0.0.0.0, or the published ports stay dead.
> *In plain English:* everything runs on the laptop with no cloud account: a fake database, fake file storage, a fake AI that returns canned answers. One command starts it all. The details are in `docs/reference/local-dev-environment.md`.

### Phase 1 - Data and rules (deterministic, fully testable)

**PR-1: Proto contract, v1.** ✅ merged 2026-08-24 (#4). Nine files under `proto/mtg/v1/`. Generated Go and TypeScript compile, and the CI diff gate passes. Contract notes: `legalities` is an open map keyed by Scryfall format keys (guardrail 2). RPC names are service-scoped (`GetDeck`, `GetCollection`) because message names share one proto package. The stream message is `ChatResponse` with a oneof event. `Question` carries `invented` and `gap_score` (D-25). `Deck` carries `seed` (D-18), `stale` (D-29), and `legality_as_of`.

The owner asked for a re-pass on 2026-08-24, against the MtG corpus. It added these fields: `COLOR_C` for produced mana, parsed `supertypes`/`card_types`/`subtypes`, `any_count_in_deck` (Relentless Rats class), commander eligibility (`can_be_commander`, `PartnerKind`, `partner_with_name`, `is_background`, `is_companion`), `Deck.sideboard` and `companion_oracle_id`, `CollectionEntry.rarity` (D-16), and `Printing.image_uris` plus `digital` (F-12, D-17).

The re-pass also produced F-18.
Messages: `Card`, `CardFace`, `Legality`, `Collection`, `CollectionEntry`, `Deck`, `DeckCard` (with `owned`, `owned_count`, `role`, `reason`), `Format`, `PowerLevel` (bracket or 60-card step, D-8), `Session`, `Turn`, `Question`, `Answer`, `ValidationResult`. Services: `CardService`, `CollectionService`, `DeckService`, `AgentService` (with a server-streaming `Chat` RPC). Connect-RPC with buf (D-7). Gate: generated Go and TypeScript compile. CI diff gate is green.
> *In plain English:* one document says what a card, a deck, and a chat message look like. Both the Go code and the web app read it. Change it in one place, and both sides update.

**PR-1b: Contract amendment (audit, D-46).** ✅ merged 2026-08-24 (#10). One proto change carries every field PR-6 to PR-9 need (F-25). The fields:
- Deck: `upgrades`, `buy_cost_usd`, and `DeckCard.price_usd`.
- Validation: `legality_as_of`, `pool_rule`, and `format` on the result. `pool_rule` and `collection_id` on the request.
- Session: `slot_states` with a `SlotState` enum, `Question.id`, `Answer`, `Turn.answers`, `status`, and `usage` with a `Usage` message (M-1).
- Chat: `answers`, `seed`, `keep_oracle_ids`, and an `AgentError` `failure` event. The string `error` event stays, deprecated, so `buf breaking` holds.
- Card: `CardFace.artist` (D-6), `partner_text`, and `max_copies_override` (F-22).
- Collection: `language`, `set_name`, `unresolved_by_reason` (M-3), and two new `UnresolvedReason` values.
- Health: `card_snapshot` and its age.

`buf breaking` now runs in CI against `main`. Gate: `buf lint`, generated code committed, every service builds.
> *In plain English:* the shared contract gains every field the next four steps need. One change now, so each later step touches only code.

**PR-2: Card database from Scryfall bulk.** ✅ merged 2026-08-24 (#5). The 200 tricky names resolve 200/200. The committed fixture covers split cards, DFCs, face names, apostrophes, and Aether spellings. End-to-end verified on the local stack. The worker downloads the three bulk files through fake-gcs and writes the completion marker (F-20). The API loads the index (about 34,000 cards) and answers Lookup and Search. Derivation facts learned from the data: the "Choose a Background" keyword has a lowercase b, and "Doctor's companion" sits on the companion card, not on the Doctor. Both are pinned by tests. This PR found and fixed F-19 and F-20.
A worker job downloads `oracle_cards`, `default_cards`, and `oracle_tags` daily (F-3: bulk only). It writes a versioned snapshot to GCS and an in-memory index in the `cards` service. The index holds the name, Oracle ID, printing ID, legalities, color identity, keywords, type line, MV, produced mana, Oracle tags, `game_changer`, `edhrec_rank`, and the image URIs per face. The index normalizes faces (F-9). The `oracle_tags` file loads into a tag tree. Rulings load on demand.

A `CardService.Lookup` by exact name, by Scryfall ID, and by Oracle ID. A `CardService.Search` with structured filters (colors, types, keywords, tags, format-legal). Gate: 100% of a fixed list of 200 tricky names resolve (split, DFC, "Aether" spelling, commas, apostrophes). A metric shows the snapshot age.
> *In plain English:* every night we download the whole card list, keep a copy, and load it into memory. Anyone can ask "which green cards with lifelink are legal in Modern?" and get a fast exact answer with no AI involved.

**PR-3: Legality freshness and announcement-day fast path (F-1).** ✅ merged 2026-08-24 (#6). The calendar is `announcement_dates.json`, embedded, with a verification date (next date: 2026-10-12). The worker checks hourly, and every 15 minutes from an announcement date until a snapshot with a legality change lands (F-23, D-47). On that snapshot, the worker logs `legality_lag` with the hours (M-2). The previous version comes from the store, not from process memory, so a restart keeps the metric. In production the worker is a Cloud Run job under Cloud Scheduler (D-48). The UI shows "Card data as of" from `/healthz`. The real M-2 number arrives with the 2026-10-12 announcement.
The worker checks the Scryfall bulk `updated_at` every hour. On and after a B&R announcement day (a committed calendar, next 2026-10-12), it checks every 15 minutes until a snapshot from that day or later lands. Every deck response carries `legality_as_of` (the snapshot date). The UI shows it. Gate: M-2 shows the lag between an announcement and the snapshot that reflects it.
> *In plain English:* ban announcements come on known dates. On those days we check more often. Every deck says which day's rules the engine checked it against, so the user knows.

**PR-4: ManaBox import (F-2, F-12).** ✅ merged 2026-08-24 (#7). Audit note: the gate count included one token row (F-21). The fix reports such rows as `NOT_PLAYABLE`. The owner's real export (2,548 rows, 4,317 cards, 18 columns) is the committed gate fixture at `go/internal/collections/testdata/`. End-to-end through the API against the full snapshot: 2,548 of 2,548 rows resolve, with zero unresolved. An identical re-upload updates the same document. Get and List work. Storage per D-16: one Firestore document with gzip entry and count payloads. Auth debt: a debug user id stands in until PR-11. M-3 rides the `ImportReport` counts until the analytics phase adds events.
CSV parser driven by the header row, not by column position. Required: `Scryfall ID`, or `Set code` plus `Collector number`, or `Name` plus `Set name`. Optional: `Quantity`, `Foil`, `Condition`, `Language`, binder name. The parser ignores unknown columns and logs them once. It returns rows that do not resolve to the user as a list, and it does not drop them silently.

The result is a `Collection` with counts per Oracle ID and per printing. Also accepts the Arena text format (`4 Lightning Bolt (STA) 42`). Storage: the store keeps the full collection (D-16). One document per collection holds a compressed entry array (printing id, quantity, finish, condition, language). 

A per-Oracle-ID count map sits beside it for fast ownership checks. A content hash of the upload detects an identical re-upload. The parser reports non-English rows to the user and skips them (D-23). Gate: a fixture set of real exports (owner-provided, anonymized) imports with zero silent drops. M-3 counts unresolved rows.
> *In plain English:* upload the file ManaBox gives you. We match every line to a real card and count how many you own. We show you the lines we did not match. We do not hide them.

**PR-5: Rules engine.** ✅ merged 2026-08-24 (#8). Audit 2026-08-24: five eligibility gaps fixed (F-22). The engine now checks the companion for legality, color identity, and singleton. `banned_as_companion` is Commander-only. Ownership aggregates per Oracle id. The golden gate is 41 good and 53 bad decks, with a test that enforces at least 30 of each. `DeckService.Validate` accepts `pool_rule` and `collection_id`, reads the owned counts from the stored collection, and returns `legality_as_of`. A pure library in `internal/rules` with embedded, dated data files: `formats.json`, `brackets.json`, and `companion_bans.json` (the F-18 list Scryfall can not express). Checks: size, copies (basics and any-count exempt), legality, commander eligibility, and all five partner mechanics. Also: color identity, Game Changers per bracket, companion (Lutri blocked as companion, legal in the 99), ownership per pool mode (D-37), and land-count and curve advisories. The golden gate ran 27 good and 31 bad decks at merge. The audit padded it to 41 and 53. `DeckService.Validate` is wired and smoke-tested end to end. Still open from F-11: the bracket prose rules (mass land denial, extra turns, combos) emit an info finding, not a check. Fixture lesson: Scryfall Oracle data contains token objects that share a real card's name, and the fixture builder now prefers real layouts.
A pure Go library. Inputs: a deck, a format, a power level, a collection, a card snapshot. Checks:
- deck size and copy limits (4, singleton). The restricted check left with Vintage on 2026-08-26 (D-155),
- legality per card on the snapshot date,
- Commander eligibility and color identity,
- Game Changer count per bracket, with the bracket prose rules encoded as data with a version date (F-11),
- ownership counts,
- land count and color-source ranges per archetype, curve summary, and role coverage. Output: a `ValidationResult` with one finding per problem, each with a severity (`block`, `warn`, `info`). Formats and brackets are data files with `verified_at` dates. Gate: table-driven tests for every rule. A golden set of 30 known-good and 30 known-bad decks.
> *In plain English:* the referee. Give it a deck and the rules. It lists every problem: too few cards, a banned card, a card outside the commander's colors, a card you do not own. It never guesses. Every rule has a test.

**M-2: Legality freshness metric.** Snapshot age, and the lag between each announcement and the first snapshot that reflects it. **M-3: Import quality metric.** Unresolved rows per import, by reason.

### Phase 2 - The agent (gated on Phase 1)

**PR-6: Candidate-list builder.** ✅ merged 2026-08-24 (#11). The human gate failed on run 1 (12 of 20) and held on run 2 (20 of 20, bar 18). Both documents stay: `docs/reference/pr6-candidate-review.md` is run 1, and `pr6-candidate-review-run2.md` is run 2 (D-65). `internal/candidates` filters the index by legality, color identity, and the commander, then scores each card. Signals come in two kinds (D-62). Payoffs reward the theme: a payoff tag such as `lifegain-matters` (1.5) or a payoff needle such as "whenever you gain life" (1.2). Enablers do the thing: a tag such as `lifegain` (1.0), a subtype (0.8), a keyword such as Lifelink (0.5), a text needle (0.4). A kind counts once, so overlapping tags do not stack. EDHREC rank adds 0.3. A staple with no theme signal keeps half its score, so theme leads. `themes.json` maps 55 theme words to Tagger slugs, payoff and enabler apart. Run 1 exposed 16 slugs that Tagger does not have, over 11 rows, which the matcher dropped without a message. `make themes-check` now fails on an unknown slug. Run 1 also showed that a parent tag carries its children. `death-trigger`, `anthem`, `flicker`, and `counters-matter` each pulled in the wrong half of a theme. Payoffs are narrow from run 2 on. An unknown word falls back to a generic rule. Roles come from the tags first (`ramp`, `draw`, `removal`, `sweeper`, `counterspell`, `protection`, `alternate-win-condition`), then from the type line and text. Brackets 1 and 2 drop Game Changers from the list. Pool modes per D-37: any-card returns about 300 by role. Owned-first returns the owned cards plus up to 50 upgrades. An upgrade must beat the weakest owned card of the same role. Owned-only returns the owned cards. Basic lands are not candidates: the generator adds them. `cmd/candidates-review` writes the gate document from a local snapshot and a ManaBox export. Two helpers serve PR-7: `ThemeColors` names the colors a theme is strongest in, and `Commanders` names commander-eligible candidates. PR-7 puts both into its questions, so the agent states a fact and does not ask the user for it. The owner scores `docs/reference/pr6-candidate-review.md` (20 prompts, 10 with the owner's collection). Meta input (PR-14) has a hook and no data. `Stats.ThinTheme` marks an owned mode with under 30 on-theme owned cards. PR-7 asks the pool-mode question again on that flag (D-63).

2026-09-02 correction (F-32, D-450): the land cap splits in two. Twenty slots go to the mana order. The lands that make two or more of the deck colors come first, and the most played lead among them. The theme ranks the other twenty, and a theme that matches fewer lands hands the rest to the mana order. Theme leads everywhere else. The mana base is not a theme matter, as D-251 said of the precon.

Given a format, colors, a theme, a power level, the pool mode, and the collection (optional, D-37), build a ranked candidate list from the engine. 

Owned-first: owned candidates plus a bounded unowned-upgrade list. Owned-only: owned candidates alone. Any-card: the whole legal pool, ranked by theme fit and `edhrec_rank`, with meta input at competitive power (PR-14). Signals: Oracle tags (theme), keywords and type lines, `edhrec_rank` (popularity), legality, ownership. Output in owned-first mode: about 150 to 300 owned candidates by role, plus about 50 unowned upgrade candidates. Output in any-card mode: about 300 candidates by role from the full pool. 

This list, not the whole database, is what the model sees. Gate: for 20 theme prompts, a human confirms the top 40 candidates are on-theme in at least 18. Ten of the 20 run with no collection.
> *In plain English:* before we ask the AI to build, the code shortlists the cards that fit: your cards, the right colors, on theme, legal. The AI picks from that list. It can not pick a card that is not there.

**PR-7: Question workflow.** ✅ merged 2026-08-26 (#12). `internal/questions` holds the whole turn loop, `AgentService.Chat` serves it, and the session store runs against Firestore.
The turn-based core. A `Session` holds filled slots (format, commander, power, colors, theme, pool rule, budget, house rules, locked cards). Each turn: a small model classifies the prompt and fills slots it can. The code decides which slots are still empty and picks up to three questions from the catalog (`mtg-corpus` skill, section 11). The model phrases them.

The user answers in free text. The small model maps answers to slots.

The agent stores, summarizes, and carries slots to the next turn, as connector-syncer's schema agent does. A slot stays open through the question phase, and a later answer replaces an earlier one. The build-run freeze of D-68 was retired on 2026-08-27 (D-241): a change after a build starts a new build. The card-pool question waits for the format, the colors, and the theme, because PR-6 needs those three before it can count on-theme owned cards (D-67).

The catalog is the first source of questions (D-25). A **gap score** decides when the catalog is not enough. It is the best catalog match between the empty slot and the user's words, from a small classifier. Below a threshold (D-27, set by M-5 with the OQ-19 rubric), the model can propose a question of its own. It gives a reason and the gap score with it. 

Every invented question is logged with its slot and outcome. M-4 reports how often this happens. Repeated invented questions become catalog candidates (PR-15). "Anything goes" and similar phrases route to the house-rules question (D-3). The pool-mode slot: with a library, the agent asks or defaults to owned-first. Without one, it defaults to any-card and does not ask (D-37).

Gate: 77 scripted conversations reach a complete slot set in at most four turns, with no repeated question (30 of D-105, and the 47 terse ones of D-522). At least 64 of the counted gate conversations use catalog questions only, the share of 25 in 30 (`has_deck` rows leave the count, D-263). The gap-score threshold is set by M-5, not by this PR.

Built on 2026-08-24, in two halves. The deterministic half holds the catalog. `catalog.json` carries the rows of corpus section 11 as data. The planner picks the questions for one turn by ask order. It asks at most three, one per proto slot, never a repeat. The freeze of D-68 is retired (D-241), so the planner silences nothing after a build.

A row carries a `slot` and a `key`. The slot is the proto field the answer informs, and the key is the row's own state. A refinement question such as table tolerance therefore survives a filled power slot.

The model half: three calls per turn (D-69). `classify` fills slots from free text. A second `classify` call scores the catalog fit and can offer a replacement. `ask` then phrases what the agent chose. The agent decides, never the model: a replacement counts only when the fit is under 0.35.

Two guards sit between the model and the user. The agent resolves a placeholder before any model sees a row, and it drops a clause with no value. A placeholder-free fallback stands in when the first sentence does not survive. A phrasing goes back to the resolved catalog text when it comes back wrong. The faults are a brace, two question marks, none at all, or a length far over the row.

Eight dogfood conversations ran on 2026-08-24, before any code. None of the eight was catalog-only, and only 5 of 26 catalog questions survived without a rewrite. One cause gave three of the invented questions: the catalog offered commander suggestions and held no question to close the slot. Section 11 of the corpus went from 11 rows to 26 from those runs. It gained 11 new rows, two rows split in two, an ask order, and a word-routing rule.

One live conversation ran on 2026-08-24 with the owner's approval, to measure a turn. Four calls, 1,875 input and 350 output tokens, $0.000795, both roles on `gpt-5.6-luna`.

It found four defects that every offline test missed. Two were mine and are fixed. The agent trusted the classifier's list of closed slots, which ended a session with three slots empty. The agent also shipped raw placeholders to the model. The model turned the agent's own statement into a second question for the user.

Two more came out of the fixes. A surviving trailing clause read as a dangling question. `Ready` also called a session complete while its questions still had no answer. The no-repeat rule empties the plan as soon as a question goes out.

A second live run on 2026-08-24 confirmed every fix. It also named three items that come before the rest of PR-7. First, the locked-cards row asks the user to keep or cut a list they never gave. It fires whenever a card is named, the commander included. The scorer rated it 0.15 and the model replaced it, which is the first real M-5 row.

Second, the commander row scored 0.05, because it asks whether the user has a commander and names three in the same breath. Its follow-up row can not fire at all: The code declares and reads `Context.Suggested`, and nothing ever sets it. Third, PR-7 passes no cache key, which costs one line when the session id exists.

The three items are fixed. The locked-cards row now fires only for a named card that is not the commander, and its text no longer presumes a list (D-70). The commander row splits in two. The base row asks whether the user has a commander, and the pick row carries the three names.

The classify schema gained `facts.wants_suggestion`, which sets the fact the pick row needs (D-71). A name the user gives as the commander closes all three commander rows. Every model call carries the session id as the provider cache key (D-72). On the owner's call, the pick row repeats after a "none" answer, and each round names three commanders the agent did not offer before (D-73).

`AgentService.Chat` and `GetSession` exist. `internal/sessions` stores one conversation as two Firestore documents in one transaction (D-74). The first is the proto session, which `GetSession` returns. The second is a private state document with the asked rows, the user's words, the card names, and the planner triggers. Chat streams the session id, one event per question, the slots, and the usage total. A model failure ends the turn with a failure event, and the store keeps the slots that the turn already filled.

The M-4 report is data, not a log line. Each question leaves a record with its row, slot, source, gap score, threshold, and the fact that its key closed. `Coverage` sums the records over one session or over many. `cmd/questions-gate` runs the 30 gate conversations against the real providers and writes the gate document.

The service wiring found four more defects, each one a row that never fired. First, the agent never offered the classifier the key of the question it just asked. An advisory key then never closed, and the session was never ready.

Second, a commander the user named left the color slot open, which blocked the card-pool question forever (D-67). The commander now fills that slot, because its color identity is the deck's color identity.

Third, `Context.AfterBuild` had no source, so the variance row was dead. The service reads it from the stored deck ids. Fourth, `ThinTheme` and `CommanderNotOwned` had no source. The service answers both from PR-6 and from the collection.

Gate status: the offline half holds. Thirty scripted conversations reach a complete slot set in at most four turns, with no repeated question. Every catalog row fires in at least one of them.

Run 1 of the live half failed on 2026-08-25 (`docs/reference/pr7-question-gate.md`). It scored 21 of 30 catalog-only, and the bar is 25. The run cost $0.0439 over 212 calls and 365 seconds. It asked 108 questions, and the model replaced 11 of them. Two replacements filled a slot other than the one the agent gave them, which is the `right_slot` field of the D-66 rubric. Two more lost the owned-only pool mode, so they were worse than the row they replaced (D-37).

The run exposed five defects. The worst one is D-82: the resolver asked the PR-6 hint source for every row, and it cached the answer under the theme alone. A commander list built before the user named their colors therefore survived the whole conversation.

Conversation 22 asked for a blue-red deck and got Lotho, Corrupt Shirriff (white and black), Peregrin Took (green), and Massacre Girl, Known Killer (black). None of the three is legal in that deck. Conversation 24 proves the cause. Turn 3 named three commanders outside the colors, and turn 4 named three inside them. A longer skip list missed the stale cache entry.

The other four are smaller. "Casual" alone triggered the house-rules row (D-78). The competitive theme row never closed its key (D-79). The pick row named three other commanders on every turn (D-80). Two rows fired before their context existed (D-81). The owner approved all five fixes on 2026-08-25.

The gap-score threshold stays at 0.35. The 11 replacements scored 0.02, 0.05 four times, 0.10, 0.18, 0.20 twice, 0.22, and 0.30. A threshold near 0.15 blocks five of them and probably passes the gate. M-5 sets that number from the D-66 rubric, and a change made to pass a gate makes the number meaningless. The full sheet of 2026-09-04 set none: no candidate meets the 80 percent floor, so 0.35 stands (D-530).

Run 2 scored 23 of 30 (`docs/reference/pr7-question-gate-run2.md`). It cost $0.0411 over 192 calls, and it asked 91 questions with 7 replacements. Two fixes hold under live conditions. Conversation 24 named three commanders inside the color identity, against three outside it in run 1. The pick row kept its three names when the user answered another question, against three new names in run 1.

Run 2 also exposed two deeper defects. The first is D-83. The classifier was able to retire a slot by name, and it retired power and the pool rule from "Brago blink deck from my library". That session called itself complete after one question, and the deck then gets a power level nobody chose. A slot with a typed value now closes only on that value.

The second is D-84, and it binds the gate itself. The gap score was not reproducible. One question scored 0.02 and 0.98 in two turns of one conversation, and 0.98 then 0.02 across the two runs. Six conversations improved between the runs and four regressed. The bar of 25 sits inside that noise band, so the number measured the scorer more than the catalog.

The score prompt asked a question of taste. It now names four faults, and the score follows the fault count.

Five of the seven run-2 replacements were not improvements. One restated the catalog row almost word for word. One dropped the bracket definitions, and one dropped the owned-only pool mode again. One asked for a budget the user gave a turn earlier. The last replaced three named commanders with an open question, after the user asked the agent to choose.

Run 3 scored 23 of 30 again (`docs/reference/pr7-question-gate-run3.md`), and it changed the instrument. Every fit landed on the fault scale: 0.90 for 67 questions, 0.20 for 7, and 0.05 for 12. Runs 1 and 2 scattered over nine values between 0.02 and 1.00.

The score now reports a fault count, so the threshold has a meaning it did not have before. At 0.35 the agent invents on any clear fault. At 0.10 it invents only on two faults or more, which in run 3 was 5 replacements instead of 10. M-5 still owns the number (D-27).

Run 3 also found a regression in D-83. The first rule blocked a close by name on a typed slot, on the theory that the classifier always returns a typed value. It does not. It reports the answer in the free-text list and leaves the field unknown. The format slot therefore stayed open after a user answered "Pioneer", and the agent asked for the format again.

The rule is now the question, not the field type: a key closes by name only while its question is out. That still blocks every case the rule exists for, because none of those keys had a question out.

Three of the run-3 replacements hit the color row. Each one replaced a statement that read as nonsense: "the best deck under budget is strongest in white, blue, black, and green". D-79 made that phrase a theme value, and the color clause used it as a subject. D-86 drops the clause when the theme names no archetype, and when the answer holds more than two colors.

The owner deleted the acquisition row on 2026-08-25 (D-87). It asked where the user buys and by what date they need the cards. The app holds no store stock and no delivery times. Scryfall gives a price estimate rather than availability (D-17).

No reader was able to act on the answer. The row also asked two things in one sentence, and every run replaced it. The catalog holds 25 rows.

Run 5 scored 29 of 30 and run 6 scored 30 of 30 on the catalog bar. Neither number stands on its own. The reword guard of D-88 refused 15 of the 16 replacements the model offered in run 5. The guard therefore decides the count, and its 0.6 threshold came from run-4 data rather than from the rubric. The M-5 sheet now carries refused rewords for that reason.

Run 6 also carried a second bar for the first time (D-91). It scored 30 of 30 and still failed, because five sessions called themselves complete with no power level. The old bar alone called that run perfect, which is the point: a conversation that stops its questions looks the same as one that finished.

The cause was the second form of D-83. The classifier closed the format by name, the format value stayed empty, and every row that triggers on the format went silent. D-83 now stands in its third form. A typed slot closes on its value alone, because a name says "answered" and never says what the answer was. A second question is the safe failure, and a deck with no format is not. D-90 puts the same rule in the classify prompt.

Run 7 confirmed the D-83 fix and exposed what the fix hid. It passed both bars, at 28 of 30 with no premature session, and the pass was hollow. Twenty-six of the thirty conversations ended with a slot unanswered, and twenty-five of those sat on "format (asked, no answer)". The agent asked the format, the user answered it, and the answer never landed. Fewer slots filled means fewer rows fire, fewer questions go out, and fewer chances to invent one. Both bars improved while the product got worse.

D-92 is the cause, and it sits under the whole D-83 history. The classify schema constrained `format` and `pool_rule` to an enum that held "unknown". The model answered "unknown" for a message that named the format outright, while it filled every free-text field in the same reply. It even guessed colors and a power step from "I want a Modern burn deck" and still left the format empty.

Four samples per variant measured it. The enum extracted 1 of 8, an enum with an empty member 3 of 8, and a free string 7 of 8. With the field freed the same probe reads 16 of 16, including the negative case where no format is named.

That explains the earlier symptom, and it does not excuse it. The model reported the format through `closed_keys` because the field itself stayed suppressed. Every version of D-83 argued about which channel to trust, and none asked why the field was empty. A live probe of about ninety calls, for roughly two cents, answered in minutes what four gate runs did not.

Run 8 is the first honest pass, at 26 of 30 with no premature session. It asked 135 questions, the most of any run, and the format-stuck count fell from 25 to 3.

Run 9 closes the last dead method. The code declared `State.Skip` and never called it, so nothing let a user hand a choice back. Run 8 measured the cost: conversation 4 answered "any colors are fine" and still ended with the color slot open. D-93 adds `declined_keys`. A decline closes any key and names no value. The generator applies the default the corpus lists, and the skipped state is what tells it to.

Run 9 scores 29 of 30 with no premature session. The count of conversations holding an unanswered slot fell from 26 in run 7 to 14.

A live check caught one precision fault before the run. The first wording let "any colors are fine" decline the bracket as well, which skips a slot the user never mentioned. Two probes and one prompt sentence fixed it.

Run 10 adds D-94, a PR-6 correction found through PR-7. `Commanders` read the 99-card shortlist and took the first legends it met. That list ends in `capByRole`, which emits one role bucket after another with lands first, so it carries no score order at all. A blink request answered with three Ojer modal double-faced cards, which the theme scorer rates 0.16. Eighty-five on-theme blink commanders existed, and Emiel the Blessed rated 0.56.

`CommanderPool` now walks the index itself. It requires a theme signal, drops the staple-role fallback, and applies no role cap, because one card fills no role quota. The 99-card pipeline is untouched, so the PR-6 gate holds.

The same pass closes the last open PR-7 item. The count of on-theme commanders in an owned pool is the weak-commander-pool signal (D-63), so nobody had to invent a score bar. Run 10 asked 138 questions and closed 92, both the best of any run. The agent offers a lifegain user with no library Vito, Thorn of the Dusk Rose and Heliod, Sun-Crowned. It offers the same user with the owner's library the best lifegain commanders that library holds, because owned-first orders them first (D-37).

Run 11 added 22 probe conversations (D-97). Every one of the 30 gate conversations holds a cooperative user who answers what the agent asks. A probe does not. It changes its mind, contradicts itself, asks a question back, or wants something the app can not build. A probe runs beside the gate and feeds the M-5 sheet. It does not count toward the catalog-only bar, because adding conversations to a bar moves the bar.

The gate set passed run 11 at 27 of 30, with no premature session. The probes found two defects at once. A declined format left the planner with nothing to route on, so no power row fired (D-98). A Yu-Gi-Oh request drew the answer "Which Yu-Gi-Oh format would you like?", because the catalog had no row that declines (D-99). They also surfaced three catalog gaps with no row at all: an out-of-scope request, a sideboard-only request, and a request for two decks at once.

The M-5 sheet now reads engine-current runs alone (D-96). A run made before a defect was fixed measures the defect. Run 3 is the clearest case. The format enum threw the user's answer away, the format row fired again, and the model invented a repair question. A score on that item pushes the D-27 threshold up on evidence about a bug that no longer exists.

Runs 12 and 13 confirm both fixes. Neither holds a premature session, across the 30 gate conversations and the 22 probes, and each passes the catalog bar at 27 of 30. The probes asked 88 and 84 questions and drew 6 replacements each, against 6 from 78 in run 11. The new scope row absorbed the request that forced the model to improvise.

Two gaps closed before the commit. The session store never executed. Every Firestore path in `internal/sessions` sat at zero coverage. `Put` ran only far enough to reject a session with no id. The code that persists every conversation was unverified. Seven emulator tests now drive it, and coverage went from 31.9% to 78.0% (D-101).

A round trip alone was not enough. `Put` and `GetState` agree with each other whatever path they use, and a rename of the private path passed. Two tests read the literal document path and hold the D-74 layout: the private state stays out of the document `GetSession` returns.

The second gap was the arithmetic. Nothing read a scored sheet, so the threshold count was manual. `cmd/m5-report` reads it and reports the fit threshold, the reword guard, and what the catalog needs (D-102). It reads the first word of a field, so the owner's free text survives. A refused reword is left out of the fit threshold, because it never reached a user.

The scoring pays for itself before it finishes. Item 8 named a fault the six of D-66 did not hold. A replacement dropped the sentence that says what the answer is for, so the user did not know whether to name decks, colors, or archetypes. `vague` is the seventh fault, and the fit does not change, because `faults` never enters the threshold calculation.

The same item exposed the reword guard. Word overlap is symmetric, so a replacement that deletes half the row scores low and passes, although it says strictly less. Item 8 scored 0.44 against a 0.60 bar. The guard now refuses a truncation as well. That is a replacement that borrows 0.8 or more of its words from the row and keeps under 0.6 of it. Measured against all 60 items, the rule refuses item 8 and nothing else (D-103).

Open in PR-7: the M-5 scoring, then the threshold and one confirming run.

> *In plain English:* the chat. "Build me a lifegain deck" fills in "theme: lifegain" and leaves format, power, and colors empty. The app asks those three, remembers the answers, and never asks twice. If the user says something vague, the app asks what they mean. It does not guess.

**PR-7B: Automated eval lane.** ✅ merged 2026-08-26 (#13). A container branch the owner names holds what the loop writes (D-142). `pr-7c` was that branch until it merged as #14. The first three evals ran on 2026-08-26, and the lane earned its place at once. It found two false rules claims that gate run 14 passed with zero linter findings (F-26, D-140, D-144). It also exposed two defects of its own. The eval judged a question against answers the user gave later, which read the ratio as 39.5 percent instead of 17.4 (D-141). It did not see that the session had a collection (D-143). The calibration agreed 80 percent, and the ten disagreements are where the value sat: two real defects, and two card facts `claude-sonnet-5` invented. Gate run 24 was the baseline on the set of D-145, D-155, and D-230. Gate run 27 of 2026-08-28 is the baseline on the set of D-263. D-302 changed the prompts on 2026-08-29, so run 27 does not compare with `main` until the owner re-runs it. PR-7 proved that reading transcripts finds defects and that scoring replacements does not. The version-1 M-5 sheet held 60 of 793 questions, because it held only a question the model offered to replace (D-104). The batch sweep of 2026-08-26 read all 66 conversations by hand and found sixteen more defects, none of which had a path to that sheet. PR-7B makes that reading automatic.

A sixth role, `eval`, scores every question of a gate run against the rubric the owner applied by hand (D-133). It runs on `gpt-5.6-luna`, the cost tier, so a 104-conversation run costs $0.092 to $0.104 (runs 14 to 25, first measured 2026-08-26). It is not the judge role: D-4 gives the judge a deck, and D-22 keeps the judge off the generator's provider. The eval role shares a model with the classify and ask roles, which the owner accepted with the risk named (D-136). Every ratio it reports is a floor.

`internal/tune` reads a gate document back and holds the accept rules. `cmd/questions-eval` writes a report a person reads and a summary a script reads. `cmd/tune-check` decides whether the loop keeps one iteration, and it costs nothing. `scripts/autotune.sh` is the loop, and it refuses to start without `AUTOTUNE_ALLOW_UNATTENDED=1`.

Three evals run first, in this order.

| Eval | What it measures | Cost |
|---|---|---|
| Gate run | The transcript. 104 conversations, 30 gate and 74 probe, read the count from `conversations.json` (D-145, D-155, D-263). | $0.152 to $0.165, measured 2026-08-26 |
| Question eval | Every question, scored for whether it deserved to be asked. | $0.092 to $0.104, runs 14 to 25, first measured 2026-08-26 |
| Deck gate | 16 to 18 golden prompts through the generator, the engine, and the judge. | $1.09, run 8 with 18 prompts, measured 2026-08-28 |
| Eval calibration | The cost-tier eval against the reference judge on 12 conversations. Sonnet 5 until 2026-09-01, Opus 5 since (D-428). | $0.25 to $0.30 on Sonnet 5. About 1.7 times that on Opus 5, unmeasured. |

The calibration answers the one question the cost tier raises: how gently does a model score work its own model produced? It scores the same 12 conversations twice, once on the cost tier and once on the reference judge. `cmd/tune-check -agree` compares the two question by question.

The reference was `claude-sonnet-5` until 2026-09-01 and is `claude-opus-5` since (D-428). It reports how often they agree, and how many questions each one refused. A cost-tier eval that refuses four where the stronger model refuses twelve is not measuring the agent. It reports a floor, and the real number sits above it. D-428 answered OQ-39 on 2026-09-01: ten points of agreement, and under 90 percent the loop's ratio is advisory only.

Four counters guard the ratio, because a run that asks less scores better and serves the user worse. Gate run 7 of 2026-08-25 passed both bars with 26 of 30 conversations unanswered. The counters are the questions asked, the questions that closed a slot, the premature sessions, and the linter findings. They come from the transcript and not from the M-4 table. The table counts the 30 gate conversations alone, and the terse set of D-105 is where the hard cases live.

Every third conversation is a holdout (D-134). The eval scores it, the report never names its failures, and the loop reads its ratio. A ratio that falls on the two thirds the fixer saw, and stands still on the holdout, is a reworded test set.

The loop that consumes these evals comes second, and it starts only when the gate below holds. The owner settled its four terms on 2026-08-26. It can change the catalog inside an approved run (D-135). It can share a model with the agent it scores (D-136). It stops at a 5 percent holdout ratio (D-137). It works on a branch of its own and pushes nothing unless the owner passes `--push` (D-138).

`docs/reference/autotune-design.md` is the authority on how it runs: the guards, the accept rules, and the owner's duties run by run. It also names the one cost the loop can not see, which is the fixer agent's own tokens.

PR-7B is the instrument and not the repair. Its scope is the eval role, `internal/tune`, the three commands, the loop, and the guards. A change the loop makes is a change to PR-7. The catalog rows, the planner, the word rules, and the prompts all belong to the question workflow. The loop writes PR-7 corrections, and the PR-7 entry records them. This keeps one concern per PR, and it keeps the measuring device apart from the thing it measures.

Each accepted iteration is one commit on a branch of its own, with the ratio and the counters in its message. It appends one row to `docs/decisions.md`. A rejected iteration keeps its evidence under `.local/tune/rejected/` (D-180), and every commit on the branch passed the accept rules.

The branches work in three layers (D-142). `main` holds the merged work. A branch the owner keeps is the container for everything the loop writes, and the owner names it. Each night cuts `auto-tune/<stamp>` off that container, and it pushes nothing. The owner reads the night in the morning and fast-forwards the container, or deletes the night.

The container matters more than it looks. A loop that always starts from `main` gives the second night none of the first night's accepted work. Two nights then change the same catalog rows from the same starting point, and the owner merges two branches that disagree. A container makes the nights add up, and it makes one reviewable pull request out of many nights.

Gate: the three evals run. The report names every defect class the batch sweep of 2026-08-26 found by hand. The calibration reports the agreement between the two eval models, and the owner accepts that number or names a stronger model (OQ-39).

> *In plain English:* the owner spent hours scoring 32 questions by hand, and one test run asks 250. This adds a second model that reads every question and says whether the agent had reason to ask it. It writes a short report instead of a spreadsheet. The loop hides a third of the conversations from anything that tries to fix the code. That way we can tell a real gain from a reworded test. The scorer is cheap, and cheap scorers are kind, so we measure how kind before we trust the number.

**PR-8: Deck generator and normalizer (F-13).** ✅ merged 2026-08-28. Deck gate run 6 passed all three bars on 16 golden prompts. Every deck passed the block checks, no invented name reached the user, and no summary stated a false rule of the game. It cost $0.8776 and it ran at prompt version 4.

The shortlist is the whole contract with the model. A card outside it is a miss even when the card index knows it. The normalizer reports a near name and never substitutes it (D-222). Two real calls copied 198 names exactly. One repair turn covers a miss, a block finding, and a deck over budget. A name that misses twice reaches the user as a note (D-223, D-244).

Five faults came out of the gate runs, and each one was a thing the code gathered and never used. The shortlist leaves basic lands out on purpose, so no deck held one (D-225). No prompt named no commander, so the generator never ran the path a user reaches with the words "you pick" (D-232). The generator handed a 60-card deck a commander (D-233). The locked card the user asked to keep never reached the build (D-242). The API streamed the deck itself and dropped it, so nothing read it again (D-245).

`TestEverySlotIsReadOrNamed` now checks every slot against the build, so a test catches the sixth (D-243).

The prompt-cache saving is a measurement, not an assumption. A cache read costs a tenth of a fresh read, which is true of the input alone. A generate call is output-heavy, so the whole call falls 19.7 percent and not 90 (D-227). The judge lane answers F-26, at $0.0034 a deck, and it is a bar of the gate (D-229).

The precon share of D-218 exists in code. The nine precon decklists live in the repo (D-247). The share is a count over the nonbasic names, and it is a ceiling (D-248, D-259). An upgrade keeps the precon's own shape and mana base (D-249 to D-251).

Prompt caching pays here, not in PR-7. The generate role runs on the middle model, and its prompt carries a candidate list of about 300 cards. That list is the same across a repair retry, and a cache read costs a tenth of a fresh read. Two rules protect the lever: keep the stable text first and the session text last, and hold one cache key per session.

The PR-7 measurements of 2026-08-24 read zero cached tokens, because each call sat near 470 tokens and the OpenAI cache starts above 1,024. Anthropic caching is opt-in, and it charges 1.25 times input to write. The judge role therefore pays for a cache only when it reads the same prefix more than twice. The adapter has no `cache_control` wiring today.

With all slots filled, the strong model gets four inputs. They are the rules summary for the format, the candidate list with roles, the role targets, and the plan request. It returns a structured deck (D-19). First, one summary paragraph on the deck's style and purpose. Then cards with exact names, counts, roles, and one line each. 

The normalizer exact-matches every name to the candidate list. The normalizer returns a miss to the model once as a tool error. A second miss becomes a user-visible note.

The engine validates (PR-5). 

The ownership check runs only in the owned modes. In any-card mode, ownership marks are information, never findings (D-37). A `block` finding triggers one repair turn with the findings as input. Then the deck goes to the user with the `ValidationResult` attached. Gate: on the golden prompts, 100% of returned decks pass `block` checks. Zero invented names reach the user.
> *In plain English:* the AI writes the deck from the shortlist, with a plan and a reason for each card. The code checks every name and every rule. If something is wrong, the AI gets one chance to fix it. What the user sees already passed the referee.

**PR-9: Designed variance (F-14, D-18).** ⏸ out of MVP scope, 2026-08-28 (D-256, amends D-18). Three reasons. Phase 3 waits on PR-8's gate and not on this one, and that gate held on 2026-08-28. The gate below asks for 30 percent, which this document already calls a placeholder until PR-15 measures it. Nobody can hold the work to that bar. No user asked for it: D-18 is a preference recorded before a user existed.

The variance row is retired with the item, because it asked which kind of variance the user wanted and nothing read the answer. `Deck.seed` and `Slots.plan_variant` leave the contract, with the field numbers and the names reserved. A later PR-9 must choose new numbers, and no stored message can be misread.

What a user has instead: the deck store keeps every build (D-245). The generate role is not deterministic, so a second ask gives a different deck. That is not variance by design, and it is not nothing.

Original text, superseded by D-256 on 2026-08-28 and kept for the record: Random variance is a feature (D-18). Variance comes from three levers, not from temperature alone. Lever 1: a seeded shuffle within each role tier of the candidate list. 

Lever 2: a "plan variant" slot (for example "lifegain aristocrats" versus "lifegain go-wide"). Lever 3: a "keep these, change the rest" re-roll. The store keeps the seed with the deck, so the app can reproduce a build on request. Identical output on identical input is not a requirement (D-18). Gate: two builds of the same prompt differ in at least 30% of nonland cards and both pass validation.
> *In plain English:* ask twice, get two different but sensible decks. Each deck remembers the dice roll that made it, so you can get the same deck back.

**PR-10: LLM role layer (D-1).** ✅ merged 2026-08-24 (#9). The live smoke passed on both adapters. Classify ran on `gpt-5.6-luna` (69 in, 21 out, 2.7 s), and judge on `claude-sonnet-5` (316 in, 20 out, 3.0 s). The pair cost $0.0013 at list price.

`internal/llm` is the one door. `roles.json` is the frozen map, dated, with an owner note per default (D-38, D-39). The baseline: `classify` and `ask` on `gpt-5.6-luna`, `generate` and `repair` on `gpt-5.6-terra`, `judge` on `claude-sonnet-5` at effort medium with thinking on (D-45). D-430 moved the judge to `claude-opus-5` on 2026-09-01. Config validation refuses a judge on the generator's provider (D-22).

Adapters: OpenAI Responses API and Anthropic Messages API through the official Go SDKs (D-40), plus the fixture `Fake` and a scripted fake for tests. Both adapters send a strict JSON Schema. The client validates the output again locally. One `Budget` per logical call: four attempts, three minutes. Truncation retries once at a higher cap, bounded to min(65,536, max(8 x cap, 8,192)). Transient errors back off with a 30-second cap and jitter.

Refusal, schema, and terminal errors return at once.

`Accumulator` reports tokens and USD per session from a dated `prices.json`. It reports null when a call gave no usage or an unpriced model.

Both providers count thinking tokens, and Anthropic prices cache writes at 1.25x (F-24). `make llm-defaults-check` warns in CI when `roles.json` or `prices.json` changes. Keys live in `.env` (D-41). The app requires keys by default. `LLM_REQUIRE_KEYS=0` (set by `make dev`) lets the fixture fake stand in (D-51).

Gate: unit tests over the fakes and `httptest` adapters pass in CI. The live smoke (`make test-smoke`) proves both adapters against the real APIs.

> *In plain English:* the AI plug. Every place that calls an AI calls it through one door with a named job. Swap the vendor in one file. Count every token.

**M-1: Token and cost accounting per session.** ✅ landed with PR-10 (#9). D-228 verified the accounting on 2026-08-27: `input_tokens` counts the cached tokens too, and the cost matched the cache. Every cost claim in this doc that carries no measured date is an estimate.

**M-5: Manual scoring lane for invented questions (D-27, F-17).** ✅ scored 2026-09-04 (D-524, D-529, D-530): 60 of 60 items, 6 unsure. No threshold meets the 80 percent floor, at 38 to 44 percent precision, so 0.35 stands (D-423). The catalog rows the model replaced are the work, and the reword guard of D-88 reads too tight.
The owner uses the product on a fixed set of prompts. For each model-invented question, a review page shows four things. The question, the gap score, the slot, and the top three catalog questions that were possible instead. The owner then scores six fields (D-66, closes OQ-19):

- `catalog_enough`: yes, no, or unsure. Was one of the three catalog questions good enough?
- `invented_better`: worse, same, or better than the best catalog question.
- `right_slot`: yes or no. Did the question target the correct empty slot?
- `filled_slot`: yes, partly, or no. The session fills this field. The owner only corrects it.
- `faults`: mandatory on every row that is not clean. One or more of duplicate, two questions in one, jargon, assumes an answer, unanswerable, out of scope.
- `catalog_action`: none, add, or reword an existing entry. This feeds D-25 and PR-15.

Scales hold three points, because a five-point scale drifts between sessions and makes the rows hard to compare. An invention has warrant when `catalog_enough` is no and `invented_better` is not worse. The report shows an `unsure` row and leaves it out of the fit. 

Scores go to the eval store with the prompt version, the model id, and the rubric version. The threshold is the gap score that best separates a warranted invention from an unwarranted one. A precision floor binds the choice: of the inventions the threshold allows, at least 80% must have warrant (D-66).

Every tenth item repeats an earlier one, which measures self-consistency across sessions. A change to the rubric invalidates the fit. The threshold is re-checked after each catalog change (D-28: the owner approves changes). Gate: at least 50 scored invented questions before the threshold is set.
> *In plain English:* the app sometimes has to make up a question. The owner will use the app and see each made-up question next to the fixed questions it had available. The owner grades it on six fixed fields. Those grades decide how eager the app is to make up questions. The rule is strict: at least four of every five made-up questions must be ones the fixed list does not cover.

**M-4: Catalog coverage metric (F-17).** Per session: catalog questions asked, invented questions asked, gap scores, and whether the invented question filled its slot. A weekly report lists invented questions by frequency. This is the input for catalog changes (D-25).
> *In plain English:* we count how often the app had to invent a question. If the same invented question appears again and again, it belongs in the fixed list.

### Phase 3 - UI (gated on PR-8)

**PR-11: Web app shell.** ✅ merged 2026-08-28 (#38). The gate held in the browser on 2026-08-28. The owner signed in over the Auth emulator, uploaded the real export and saw the count, and took the skip path to the chat placeholder. A curl call proves the token path too. The unresolved rows show the line and a reason in plain words, and not the raw row. 23 web tests, axe on each page. The bundle is 546 kB, almost all `firebase/auth`, and a code split is later polish.
React 19, Vite, TypeScript, Tailwind, the wallabee-ui patterns (TanStack Query, Zustand, lint-enforced import boundaries). Firebase Auth (D-11) with the emulator in local mode. Generated Connect client in `packages/api-client`. Gate: sign-in, then either upload a collection and see the count, or skip the upload and still reach the chat (D-37).

Detail of 2026-08-28: `docs/reference/ui-plan-2026-08-28.md` holds the user path, the architecture, and the live-test procedure (D-273 to D-276). The sign-in is real, over the Auth emulator, and the token reaches the API through the interceptor of D-268.
> *In plain English:* the website skeleton: log in, upload your binder, see how many cards we recognized.

**PR-12: Chat and deck view.** ✅ merged 2026-08-28 (#40). The test gate held. Axe passes on the session page and the deck view, and a test asserts the full card image, uncropped, on every card (D-291). The browser gate waits for the owner (README section 6).
A streaming chat thread over the `Chat` RPC. The deck view groups cards by role. It shows card art from Scryfall image URIs with artist and copyright (D-6, guardrail 7). It shows both faces for DFCs (F-9). 

It marks owned versus to-buy when the session has a collection. It shows the pool-mode toggle ("use only cards in my library") with the session's mode (D-37). In any-card mode, the buy list can be the whole deck.

It shows the mana curve, the color sources, the `ValidationResult` findings, and `legality_as_of`. ~~Hover or tap shows Oracle text.~~ (Struck 2026-08-29, D-290 and D-291: a card tile shows no caption and no Oracle-text hover.) Gate: a11y checks pass. Every image has attribution in the DOM.

Contract addition of 2026-08-28: `CardService.GetCards` returns up to 120 cards by Oracle id in one call. `DeckCard` carries only the id and the name, and one `Lookup` per card is 100 calls per deck (ui plan, section 4). Landed 2026-08-28 (D-277). The chat holds an open question across a turn that asks nothing (D-278). The attribution line is the one the Scryfall docs ask for (D-279).
> *In plain English:* the main screen. The conversation on one side, the deck on the other with real card pictures, grouped by what each card does, with your own cards marked.

**PR-12B: Deck revision turns.** ✅ merged 2026-08-29 (#41). Built 2026-08-28 on branch `pr-12b` (F-27, D-283 to D-285). Gate run 1 on 2026-08-28: 5 of 6 revisions met their bar. The sixth failed the gate's own bar on a mixed message, and not the model (D-296). Run 2 the same day, with the bar fixed: 6 of 6, $0.29. The gate held.
After a build, every message is a request to change the latest deck. The turn has three parts. First, the classify call runs as today for a slot change. A changed slot means a full rebuild with a note that says so (D-241). Second, with no slot change, a new revise call reads the message and the current deck list. It returns a revision brief with four parts:

- the cards to remove,
- the cards to keep,
- the numeric limits it found, for example a top mana value,
- one clarifying question when the request is unclear. The question passes the same lint as every other question,
- the requests it declines, each with a reason.

A request has three outcomes: a question, a change, or a decline with a reason (D-284). "Replace some lands with better options" is unclear on its own, so the agent asks what the user means, for example faster mana, utility lands, or more colors. When no change helps, the reply says so in plain words: "for a casual mono-white deck, all basic lands is fine". A decline is not silence, and the deck stays as it was.

Third, the generator gets the brief and the base deck in `generate.Request`, and its instructions say to keep every card the brief does not touch.

A deterministic check reads the brief after the build. A removed card that is still present is a finding. So is a kept card that is absent, and so is a card over the mana limit. The repair turn reads these findings like any other.

2026-09-02 correction (F-31, D-448, D-449): a land swap is a counted change. The brief holds the number of basic lands to replace and the kind. The count fits the base deck and the pool. A deck that holds fewer new nonbasic lands gets a block finding. The repair turn reads the block. The generator's revision rule says a change that names a group and a number touches that many cards. The revise prompt says the earlier message still stands after the answer to its question. The gate answers its own questions and scores the rebuild, and it holds a clear land-swap row on the Karlov base. A turn stores its brief. Revise gate run 5 ran on 2026-09-02, and 10 of 11 turns passed. The owner's own land ask failed on the Karlov deck, because the shortlist held no untapped dual (F-32, D-450). Run 6 ran the same day: 9 of 10. Revision 9 swapped 14 basic lands for the fixing D-450 put on the list, with no repair turn. Revision 1, the mixed message, skipped its question and acted on a guess, which the prompt's two land rules had left open. The prompt closes it. Run 7 ran the same day: 11 of 11, $0.74. Every land turn asked first, then swapped. Run 7 also showed the other side of D-450. The Karlov base deck took 30 nonbasic lands and 6 basics, about 10 of them colorless utility lands, and run 5 took 24 basics. Deck gate run 11 measures whether that holds across 18 decks before a rule follows.

The reply is prose from the diff, not from the model, for example "I removed four cards over 5 mana, added four, and replaced six Plains". The server stores it in `Turn.agent_message`, which existed and had no writer, and streams it as `text_delta`. The new deck gets `revised_from_deck_id` and `revision_note`, both additive, and the deck view shows the diff. Gate: six revision prompts over two stored decks. Each result keeps every untouched card, holds every limit of the brief, and passes the engine. A paid run, so the owner says when.
> *In plain English:* today, the app throws away anything you type after the deck appears, and the app quietly builds the same deck again. After this change, "cut the 7-drops and fix the lands" gives you a short answer and a deck that did those two things, or a question when the request is unclear.

**PR-13: Export and share.**
Export as ManaBox text first (D-15). Other formats later. A buy list with Scryfall purchase links. Gate: a round trip ManaBox export to import loses nothing.

Contract addition of 2026-08-28: `DeckService.ExportDeck` lands with this PR, with `EXPORT_FORMAT_ARENA_TEXT` first. ManaBox imports the Arena text shape, and `collections.ParseArenaText` reads it, so the gate runs against our own parser (ui plan, section 4). The Export RPC of PR-0a left the contract on 2026-08-28 because nothing implemented it (D-266).

Built 2026-08-29 on branch `pr-13` (D-307 to D-309). The Arena line names the owned printing when the user owns the card, else the default paper printing (D-307). The buy list holds the shortfall of the commander, the main deck, and the sideboard, with a Scryfall link to the printing (D-308). The upgrades sit under their own heading. A second format, `EXPORT_FORMAT_BUY_LIST_TEXT`, exports the buy list as "count name" lines (D-309). The export line carries the full card name, because the index resolves a full name without ambiguity.

`go/internal/export` holds the renderer, and `TestArenaTextRoundTrip` is the gate: it held on 2026-08-29. The panel sits in the deck view, and `export` is a leaf feature that `deck` imports.
> *In plain English:* get the deck out of the app and into ManaBox or Arena with one click, plus a shopping list.

### Phase 3B - The product UI (gated on PR-13)

The live-test UI of Phase 3 served one purpose: the owner tests the agent in a browser (D-273, F-28). This phase builds the product (D-310). It runs for the owner locally, and for invited users on GCP at the end of the phase.

The look comes from a reference design the owner gave on 2026-08-30 (D-328 to D-330). Three faces carry three jobs. Cinzel engraves a heading, Crimson Pro reads a paragraph, and JetBrains Mono carries an id or a count.

The palette is navy, gold, and purple, with parchment for text, and the radius is 4 px. The shell is one top bar, and dark is the only theme. Radix primitives still carry the behavior (D-311), and the app owns each file.

The test bar is Vitest with axe per pull request and one Playwright smoke flow on a manual trigger (D-313). A session also reads its own work with Playwright from 2026-08-30, and it measures rather than looks. `docs/reference/ui-phase-plan-2026-08-29.md` holds the screens, the components, the contract changes, and the deploy shape.

The four flows of D-312 come in this order:

- PR-16, the design system and the shell.
- PR-17, the deck library.
- PR-18, the collection management.
- PR-19, the chat and build experience.
- PR-20, the deck view and the card detail.
- PR-21, the share link and the print view.
- PR-22, the deploy for invited users.
- PR-23, the Playwright smoke flow.

> *In plain English:* what exists today is a test bench with a browser on it. This phase makes it an app a person can use every day, on a laptop or a phone, and later from anywhere with an invitation.

**PR-16: Design system and app shell (D-311, D-317).** ✅ merged 2026-08-29 (#47).

CAUTION: the palette, the shell, and the theme of this slice all changed on 2026-08-30. D-328 replaced the sidebar with a top bar, D-329 removed the color identity of a deck, and D-330 removed the light theme. The primitives and the route split of this slice stand.
Nine shadcn primitives on Radix, written by hand into `src/components/ui`: Button, Input, Label, Textarea, Checkbox, Card, Skeleton, DropdownMenu, and the toast (D-321). Each later slice adds its own.

Tailwind 4 tokens in `src/styles/tokens.css`: a neutral scale, one accent, a link color, the six mana colors, and the semantic roles. The dark theme overrides the neutrals, the surfaces, and the link only. The theme follows the system by default, and a menu in the sidebar stores a choice. A script in `index.html` paints the class before the first paint.

The shell is a sidebar on a desktop and a bottom tab bar on a phone, with Build, Decks, and Collection. A media query picks one of the two, because two navigations with one name fail the axe landmark-unique rule. One `PageHeader`, one `EmptyState`, one `ErrorState`, and a toast for every mutation. Every existing screen moves onto the primitives with no new feature. The import boundary of the lint gains `src/components/ui`, and a primitive imports no feature and no app code.

Five things leave the first paint. They are `firebase/auth`, the Connect client, the five feature pages, the two shell menus, and the toast host. Each one loads when the app first needs it.

The baseline of 2026-08-29 is one chunk of 584.07 kB raw and 179.81 kB gzipped. It also holds 14.24 kB of CSS and 118 web tests.

Gate:

- axe passes on every route in both themes. ✅ four routes, two themes.
- The 118 web tests hold. ✅ 135 tests pass in 17 files, and the 118 hold.
- The first paint holds under 130 kB of gzipped JavaScript (D-323). ✅ one file of 116.31 kB gzipped, 363.25 kB raw.
- The owner walks the whole path on a desktop and on a phone. ⏳ waits for the owner.

CAUTION: D-320 set this bar at 200 kB of raw JavaScript, and a measurement showed that no build can reach it. React, the router, and TanStack Query are 101.51 kB gzipped and 319.51 kB raw together. D-323 amends the bar to 130 kB gzipped.

> *In plain English:* the look and the bones. Buttons, dialogs, menus, and a dark mode that all match, on a layout that works on a phone. Nothing new to do yet, but everything looks and feels like one app.

**PR-16B: The visual pass (D-325, D-326).** ✅ merged 2026-08-29 (#48).
PR-16 moved every screen onto the primitives and kept each composition, per D-317. So the system changed and the screens did not, and the app still read as the test bench of PR-11. The owner said so after the merge.

The pass gives the app one type face, Geist, that ships with the build and waits on no network. It adds an elevation scale of three shadows, a wider radius scale, and a reading measure of about 68 characters.

The chat thread carries hierarchy. A turn of the user reads as a block on its own surface. A turn of the agent reads as plain text at the measure. The notes of a turn stay quiet. A bubble on both sides reads as a messenger, and this is a tool.

The composer docks at the foot of the column with its control inside it, and it leaves while the agent works (D-325). The session id, the pool line, and the spend line leave the reading column. The sidebar takes a mark and an active bar, and the bottom bar takes one too.

Gate:

- The 135 web tests hold. ✅
- axe passes on every route in both themes. ✅
- The first paint holds under 130 kB of gzipped JavaScript (D-323). ✅ 116.57 kB, from 116.31 kB. The font is a separate asset of 29.4 kB.
- The owner reads the app in the browser. ⏳ the owner merged the slice and reads it next.

> *In plain English:* the app looked like a test bench with a dark mode on it. This makes it look like a product. A real typeface, depth, a chat that reads like a conversation, and a message box where you expect it.

**PR-17: Deck library.** ✅ merged 2026-08-31 (#49).
A grid of decks with the name, the commander, the mana pips, the format, the power, the count, the buy cost, and the date. Search by name and commander, filter by format, power, and favorites, and a favorite star. Every filter runs on the server, so a match on a later page still shows.

Contract, additive: `DeckService.UpdateDeck(name, favorite)`, `DeleteDeck`, `Deck.favorite`, `Deck.card_count`, and paging with filters on `ListDecks` (D-245). The power filter has two fields, because `PowerLevel` is a oneof (D-324). `ListDecks` also takes a `session_id` (D-343).

A deck carries its version history (D-343). A revision turn writes a new deck in the same chat, so the decks of one chat are the versions of one deck. The screen reads them with one `ListDecks` call on the session id, oldest first, and every version keeps its own address. Compare takes any two of them and reads the diff that the revision note already used.

A deck has one screen and one address (D-335). `/decks/<id>` holds the deck, the actions the user owns, and the conversation that built it. `/session/<id>` holds a build with no deck yet, and it hands the reader over the moment a turn ends with a deck. The deck fills the page, and the conversation docks at the bottom left with a History control (D-331). `src/features/workspace` is the one feature with a path to both chat and deck, and the import boundary of the lint carries that rule.

Build in the header asks which cards the deck draws on, and then opens a chat (D-332). A signed-in reader lands there (D-334). The chat screen carries the same choice as a "Build from" picker under its title (D-336). A control in a menu alone is a control a reader does not find.

No chunk of this app loads behind a Suspense boundary (D-338). React holds a committed fallback for 300 ms, and it holds every later reveal with it. A chunk that is already in the browser therefore costs the reader a third of a second. `src/app/deferred.tsx` starts a download and mounts the component the moment the code is here. `src/app/chunks.ts` lists every deferred chunk, and the shell warms them all in the idle time after the first paint (D-339).

The listing filter runs in Go over the rows Firestore returns, not as a Firestore query. One read serves every filter, and no composite index has to exist. A scan cap of 500 rows bounds the read, and a user beyond it needs a search index. The page token carries the offset and a fingerprint of the filter, so a token of another filter is an invalid argument.

Gate:

- Each action round-trips through the API and shows in the grid with no reload. ✅
- A deleted deck answers `NotFound`. ✅
- A grid of 100 decks renders under one second. ⏳ the owner reads it with real decks.
- The version history lists every deck of one chat, and any two compare. ✅
- The first paint holds under 130 kB gzipped (D-323). ✅ 116.13 kB, 362.42 kB raw.
- Content of the landing page shows under 100 ms, measured over the built app. ✅ 44 ms, from 347 ms.
- The first open of a shell menu costs under 50 ms. ✅ 16 ms, from 323 ms.

CAUTION: this branch carries eight concerns. They are the contract, the Go side, the reference design, and the layout of D-331. They are also the Build menu, the one deck screen, the pool picker, and the speed of the app. Guardrail 10 asks for one. The owner chose to ship it whole (D-344).

> *In plain English:* a home for your decks. Find one fast, name it, star it, throw one away, and talk to the agent about it on the same page. Every revision keeps its own copy, so you can read an earlier one and see what changed.

**PR-17B: The set filter.** ✅ merged 2026-09-01 (#50).
A deck can be limited to one set or to several. The request carries the set codes, and every stage reads them: the 99-card shortlist, the commander pool, and the deck check. A card passes when it holds a paper printing in one of the named sets.

The card data carries the sets. `Card.set_codes` is every paper set the card has a printing in, lowercase and sorted. The index already reads every printing to build `bySetNo`. It collects the codes in that same walk, and it shares one string per set code.

The snapshot of 2026-08-31 holds 117,608 printing rows, and 9,345 of them are digital. It holds 34,599 playable Oracle cards over 988 paper sets. A card carries 2.4 sets on average, and the card with the most carries 225. In all, 16,247 cards hold printings in more than one set, and the whole field costs about two megabytes.

CAUTION: a card does not have one set. Nearly half of them hold printings in two or more, so the field is a list and never a value. A filter that reads one set per card drops a reprint the user owns.

The name a reader gives is a product, not a code. A set family is a base set and every product Scryfall names under it (D-376). "The Hobbit" gives `hob` and `hoc` The Hobbit Eternal. A token, memorabilia, or minigame product stays out.

CAUTION: a name prefix can not build a family. `ltc` is "Tales of Middle-earth Commander", and its base set is "The Lord of the Rings: Tales of Middle-earth". The two names share no prefix. `parent_set_code` links them, and no bulk card file carries it, so the snapshot gains a fourth file from the `/sets` endpoint (D-377).

The resolver reads an exact set code, then an exact set name, then a whole-word phrase. It drops a candidate whose parent is also a candidate. One base set left resolves to its family, and two or more ask. "Tarkir" and "Ravnica" ask. "The Hobbit", "Final Fantasy", "Lord of the Rings", and "Bloomburrow" resolve.

A set limit lifts two cuts (D-379). The theme cut drops a card with no theme signal that fills no staple role. The role caps then drop more. Together they turned the 128 cards of the Hobbit family in black-red into 40 to 50, and a Commander deck needs 99. Inside a set the theme ranks the list and cuts nothing, and the staple penalty still puts every on-theme card first.

A set family too thin for the format builds no deck (D-380). The floor is 70 distinct nonbasic cards in the deck colors, and 35 for a 60-card format. Only 147 of 717 live sets offer 99 cards in their best two-color identity, so a refusal is the common case. The check runs before the commander pool. A deck that can not exist then costs no model call.

Three things pass the filter and carry a mark. A card the reader named by name is an instruction (D-381). A ramp card or a nonbasic land comes from outside only when the reader allows it, and the agent asks first (D-382). No set limit filters a basic land (D-378). `DeckCard.outside_requested_sets` marks the first two, the deck screen shows a red mark, and the deck holds one warning that counts them.

Contract, additive: `Slots.set_codes`, `Card.set_codes`, and `DeckCard.outside_requested_sets`. `CardService.ListSets` and `ChatRequest.set_codes` move to PR-18 and PR-19. Nothing in this slice calls them, and the agent resolves a set name from the reader's own words. The catalog gains two rows. One asks which set a name means. The other asks whether the mana base reaches outside the sets.

Gate:

- A deck asked for one set holds cards of that set family alone, and the check reads `set_codes` of every card of the deck.
- A commander offer for one set names commanders of that set. The Hobbit family offers Smaug the Impenetrable, Thranduil, the Elvenking, and Smaug, Wicked Worm, which the owner named on 2026-08-31.
- A set too thin to build a legal deck ends the turn with a reason, never a deck of another set.
- A set name the app can not resolve asks the user, and it names the sets it does hold.
- Every card the sets do not hold carries the mark, on the deck screen and in the deck.
- The index builds in the same time, plus or minus one second, and it holds under 40 MB more.

CAUTION: the gate line above named two cards that do not exist. The snapshot holds "Thranduil, the Elvenking" and "Smaug, Wicked Worm", both with a comma. Smaug the Impenetrable is in `hoc` and not in `hob`, so only the family rule of D-376 can satisfy that line.

> *In plain English:* today you can ask for a deck from one set and get cards from anywhere. Nothing checked the set, because the app never recorded which sets a card is in. This adds that record and the filter over it. A card can be in many sets, so the app keeps them all: a reprint still counts. It also learns set names, because you say "the Hobbit set" and the data says "hob". A set that is too small to build a deck gets a plain answer, not a bad deck.

**PR-18: Collection management.** ✅ merged 2026-09-01 (#53). **The review fixes (D-398 to D-406)** 🔧 built 2026-09-01 on branch `nits-and-fixes`. The owner reads them in the browser, then merges.
The list of collections shows the name, the count, the date, and the active mark, with rename and delete. An upload dialog holds the file, the progress, the diff, and the import report, one step at a time (D-395). A re-upload over the active collection shows the diff first: the added, removed, and changed rows with their card counts, then "Replace" (D-393). A reader with an active collection chooses between a replacement and a new collection (D-400). Replace keeps the collection id, so every deck and chat that names it still works.

The binder is a virtualized grid of the cards with the art, the count, the finish, the condition, and the price (D-394). It has a search, four filters, and four sorts. The filters are the set, the color, the card type, and the count. The sorts are the name, the count, the set, and the price. Every choice runs on the server, so a match on a later page still shows (D-398). The page token names the filter and the sort, as a deck page token names its filter.

A collection carries its own summary (D-392). The import computes the row count, the unique cards, and the rarity spread. It also counts every set and every card type, and it picks the ten rarest cards. The head reads the summary and asks for no entry, and the two filter menus read it too. The colors, the card types, and the price ride on an entry, and no document stores them. `GetCollection` fills the three from the card index of the day (D-396).

Contract, additive: `CollectionService.UpdateCollection` and `DiffCollections`, `ImportCollectionRequest.replace_collection_id`, `Collection.summary`, and the three display fields of `CollectionEntry`. `GetCollectionRequest` gains `page_size`, `page_token`, `entries_omitted`, `filter`, and `sort`, and the answer gains `next_page_token` and `matched_rows`. `DeleteCollection` came in PR-17 (D-347).

CAUTION: a first upload lands on the document id its content hash derives. A Replace keeps that id and moves the hash on. Before D-399, a later upload of the old file derived the same id and overwrote the replaced collection. `Put` creates the derived document now, and when it exists with another hash the upload takes a fresh id.

Gate:

- The same file uploaded twice diffs empty. ✅ held in `TestTheSameFileDiffsEmpty` and in the browser on 2026-09-01.
- The binder of the owner's export (2,657 rows) scrolls at 60 frames per second on the owner's laptop. ✅ the owner closed it on 2026-09-01. The measurement of a session the same day, over the 2,547-row test export in headless Chromium, was not that measurement.
- A session that names a deleted collection falls back to any-card mode with a notice (D-37). ✅ held since D-347.

> *In plain English:* your binder, on screen. Several uploads, a name on each, and a clean "what changed since last time" when you upload a new export. Browse it like a real binder, with the pictures, and search the whole binder rather than the part on your screen.

**PR-19: Chat and build experience (D-432 to D-458).** ✅ merged 2026-09-02 (#55), with its follow-ups the same day (#56). The owner read the stepper on a real build, and it held. The follow-ups are five. The count includes the commander (D-454). The favorite star sits at the tile's edge (D-455). A deck and its chat are one thing (D-456). An upload sits on one line (D-457). The collection page holds two cards of one frame (D-458).
The chat is the whole start (D-436). The owner read a start form on 2026-09-02 and refused it. The one control outside the conversation is the pool picker, the collection or any card. The form, its rows in the contract, and the gate's form conversations left the same day. D-432 and D-434 record the form, and D-436 amends both.

The thread already holds the primitives of D-295: option buttons, art tiles for a card option, and no field on a closed question. A stepper lights the phase the server streams: understand, shortlist, build, check, and repair when one ran (D-435).

Error recovery has three parts. A failed turn shows the reason and a "Try again" that sends the same turn. A build in progress shows the D-303 notice. A lost stream offers "Reload the session", which reads the stored session again.

A new chat lists the unfinished chats under its message box, and only when there is one (D-433, D-438). Those are the conversations with no deck yet, with resume, rename, and delete. A finished chat lives on its deck, and the top bar keeps three entries.

Contract, additive: `AgentService.ListSessions`, `UpdateSession`, and `DeleteSession`, with `SessionSummary`. `Session.name` and `ChatResponse.phase` are new fields, and `generate.Request` carries a phase callback.

The slice also carries five findings of the owner's read. A deck tile carries a delete (D-439), and every control shows the pointer cursor (D-440). The wordmark's shimmer no longer runs forever, and a dialog overlay fades with no backdrop filter, so a delete dialog opens with no stutter (D-441).

The gap over the message box equals the gap under the top bar (D-442).

A half of a commander pair zooms to a single card's size, and a single card never zooms (D-443). The art of an option picks it, as its name does (D-444), and the tile lifts under the pointer (D-445). A stored question closes with its slot, so the docked chat of a fresh deck shows the pick and not the offer (D-446).

The session spend holds the build, which it never did (D-447). The commander offer is fixed (D-437). A declined theme answered no name, so the pick row went out bare and the build chose a commander with no word to the reader. The offer now serves an empty theme on popularity, and it reads the set limit, so a Hobbit-only request offers Hobbit commanders.

The second half of 2026-09-02 fixed the revision turn and the shortlist. A land swap is a counted change. The brief counts the basic lands to replace, the check holds the deck to the count, and the generator's rule names group changes (F-31, D-448). A turn stores its brief (D-449). The land bucket of the shortlist splits, half mana and half theme (F-32, D-450).

Revise gate runs 5 to 7 and deck gate run 11 are the records, and both gates pass. The read of every mana base is F-33, which binds PR-14A. The same day the owner split PR-14 and put the bracket profile right after this PR (D-451 to D-453).

Gate:

- The stepper shows every phase of a real build. The phase events come from the generator itself, so a step the stepper shows is a step that ran.
- A declined theme under a set limit offers three commanders of the sets. ✅ held in `TestCommandersWithNoThemeInsideTheSets` on 2026-09-02.
- The owner builds one deck from the chat and reads the stepper.
- Revise gate run 7: ✅ 11 of 11 on 2026-09-02, $0.74. Deck gate run 11: ✅ 24 of 24 on 2026-09-02, $1.46.

CAUTION: the form path of the first plan sent the `Q:`/`A:` shape to the classify call (D-280). The form left (D-436), so only the browser sends that shape, as before.
> *In plain English:* the chat is the whole start, and it shows the build move through its steps. A turn that fails offers a retry, and a lost connection offers a reload. A chat you left before a deck waits under the message box, and a deck you no longer want goes from its tile.

**PR-20: Deck view and card detail (D-318, D-507).** ✅ merged 2026-09-03 (#61). The free gate is `docs/reference/pr20-gate-2026-09-03.md`.
A click on a card opens a detail panel. It shows the full image and both faces, the Oracle text, the type line, the mana cost, and the rulings with dates. It also shows the legalities, the printings with prices, the deck's reason line, and "Open on Scryfall". The deck view gains filters by role, color, mana value, type, and owned, and sort by mana value, name, and price. Stats show as small charts with a text table under each one.

The stats are the curve, the color sources, the type counts, the average mana value, and the buy cost. A sample hand draws seven from the exact main deck, mulligans to six and five, and draws one. It simulates no turn (D-318, amends D-20).

Contract, additive: `CardService.GetRulings`, from the Scryfall rulings bulk file the worker downloads with the daily snapshot. 2026-09-03: `CardService.GetPrintings` joins it, because the index held one printing per card and the panel shows every printing with its price (D-507). The rulings file is the fourth snapshot file, optional on load, so an older snapshot still loads.

Gate:

- axe passes on the panel.
- A test draws every card of a 60-card deck through the sample hand.
- The rulings show their dates and the snapshot date.

> *In plain English:* tap a card and read everything about it. Filter the deck the way you think about it. Shuffle up and look at a seven-card hand.

**PR-21: Share link and print view (D-315, D-508).** ✅ merged 2026-09-03 (#62). The free gate is `docs/reference/pr21-gate-2026-09-03.md`.
"Share" on the deck page makes an unguessable token and shows the link, and "Revoke" ends it. A public read-only page at `/d/<token>` shows the name, the format, the power, the summary, the cards by role with art, and the export button. It shows no owner name, no collection, no session, and no owned printing. A print stylesheet renders the deck page as the list by role in black on white, with no images.

Contract: `DeckService.ShareDeck`, `RevokeShare`, and `GetSharedDeck`. The last one needs no sign-in, and a rate limit per IP bounds it. The deck stores a hash of the token, never the token. 2026-09-03: `ExportSharedDeck` joins them for the export button of the public page, under the same limit. The shared message carries the card data inline, so the page makes one call (D-508).

Gate:

- A revoked link answers `NotFound`.
- A test reads the shared message and finds no user field (guardrail 13).
- The rate limit refuses the 61st call in a minute, and it reads the client address from `X-Forwarded-For`.

CAUTION: behind Firebase Hosting and Cloud Run, `RemoteAddr` holds the address of the proxy. A limiter that reads it puts every visitor in one bucket, and the gate then passes for the wrong reason.

> *In plain English:* send a deck to a friend with one link, and print it for the table. The link shows the deck and nothing about you.

**PR-22: Deploy to GCP for invited users (D-310, D-314).**
The API and the worker run on Cloud Run from the Dockerfiles of PR-0c, with min instances at zero. The web app runs on Firebase Hosting, with a rewrite of `/mtg.v1.*` to the API. Real Firebase Auth with email and password. The interceptor reads the allowlist of D-314, and a uid off the list gets `CodePermissionDenied` with one sentence. OQ-45 holds the store of the list, and the owner answers it before the slice starts. Firestore runs in Native mode with the deny-all rules of the repo.

The worker refreshes the card snapshot in a GCS bucket on Cloud Scheduler, and Secret Manager holds the provider keys. A per-user monthly spend cap reads `Usage`, and a budget alert sits on the project.

Gate:

- An allowlisted user signs in on the deployed URL, uploads a collection, builds a deck, revises it, and exports it.
- The first RPC refuses a user off the list.
- The roadmap records the measured monthly cost at idle.
- Guardrail 9 holds: `make dev` still runs with no cloud dependency.

> *In plain English:* the app on the internet, for the people you invite and nobody else. A cap limits what any one person can spend.

**PR-23: Playwright smoke flow (D-313).**
One flow on `workflow_dispatch` only. It signs in over the emulator and uploads the fixture export. Then it starts a session from the form with the fake provider, opens the deck, and exports it. The fake provider serves canned answers for the classify, ask, and generate roles, so the flow costs nothing. One run takes about 5 minutes of Actions time, and the owner triggers it before a merge that touches the user path. Gate: the flow passes on the emulators.
> *In plain English:* a robot that clicks through the whole app once, on demand. A change that breaks the path shows up before it ships.

### Phase 4 - Meta and quality (gated on Phase 3B, D-316)

**PR-14A: The bracket profile (D-451 to D-453, D-459 to D-469).** ✅ merged 2026-09-02 (#57). Bracket gate run 1 reads FAIL on two bars, and the owner merged with that on record. Deck gate runs 12 and 12b together pass all 24 prompts with no regression.
Today a bracket reaches the build as one prose line, "Commander bracket: 3", and one cut: no Game Changers under bracket 3. The role targets are one table for every bracket. The engine checks legality and the Game Changer count, and it notes that the prose rules of the bracket are not machine-checkable. No check reads power after the build. A 3 is whatever the model believes a 3 is.

The profile is a specification per bracket, in data and not in prose. It has four parts.

The content rules. What a bracket forbids, as card flags: Game Changers, mass land denial, extra turns, tutors, fast mana, and two-card infinite combos. Scryfall flags the first. Commander Spellbook classifies the rest through its `estimate-bracket` endpoint. It takes a deck list and returns a flag per card for Game Changer, mass land denial, and extra turn. It also returns a flag per combo for two-card, speed, lock, and extra turn, with its own bracket tag.

It answered an anonymous call on 2026-09-02, and its backend is MIT-licensed. The owner read the terms on 2026-09-02, and they allow the call (D-459, closes OQ-50). The client holds itself to 90 requests a minute, and it sends a named agent.

The shortlist drops what the bracket forbids twice. The tags `mass-land-denial` and `extra-turn` cut it for free (D-462). Then the endpoint reads it, and the builder drops every card it flags for the bracket (D-468). The endpoint checks the built deck, and the repair turn fixes a miss. The note in `checkBracket` left.

The deck bands. A feature vector per built deck. It holds the average mana value and the curve, and the ramp, draw, removal, wipe, and interaction counts. It holds the tutor and fast-mana counts, the untapped share of the lands, and the color sources against the pips by the Karsten tables. It holds the combo count and the Game Changer count.

Each bracket holds a band per feature. The generator reads the bands as its targets, in place of the one table of `TargetsFor`. The check runs after the build, an off-band feature is a finding, and a finding buys the repair turn as a budget miss does (D-244).

The goldfish simulation (D-453). It deals ten thousand opening hands with the London mulligan, and it plays the lands and the rocks on curve. It reports three numbers. They are the turn the deck casts the commander, the mana available on turn four, and the share of hands with two to four lands. The bracket table counts turns, and this is the one number a card list can not give. No opponent model.

The proof. A bracket gate builds three commanders at each bracket, and every deck must sit in band with no content violation. The judge of PR-15 reads each deck with the bracket definitions in hand and must agree with the bracket in eight of ten.

Contract, additive: `Deck.profile` with the features, the bands, and the off-band findings.

Gate: the bracket gate above, and the deck gate re-run with no regression. The land band of F-33 is part of the profile. It holds the basic and nonbasic counts, a cap on colorless lands, and the color sources against the pips. It also holds a cap on tapped lands that falls with the bracket and the power step. The repair turn gains a second pass or a bounded re-roll when the profile stays off band. That is the re-roll of PR-9 with a score to pick by (D-256).

What the build holds (2026-09-02). `internal/spellbook` is the client, at 90 requests a minute with a named agent (D-459). `internal/profile` is the library: the bands in `bands.json`, the feature vector, the Karsten tables, the goldfish simulation, and the content check. `rules/brackets.json` gained the content rules of each bracket, and the prose note of F-11 left.

The generator reads the band midpoints as its job targets, and the prompt carries a deck shape block with the rest. A profile finding is a warning that buys the repair turn, and a profile finding alone buys one more pass (D-461). `cmd/bracket-gate` and `make bracket-gate` are the gate, with the spend guard and the overwrite guard. `docs/reference/bracket-profile-2026-09-02.md` holds every source.

CAUTION: five band groups have no published source (D-463). They are the average mana value, the tapped and colorless caps, the tutor and fast mana caps, and the goldfish floors. The first gate run measures them. Read the off-band table of the document before you move a band, and give each move a decision id.

Bracket gate run 1 (2026-09-02, `docs/reference/pr14a-bracket-gate-run1.md` and its judge lane `-run1-judge.md`). The builds cost $2.08 over 46 calls, and the judge lane $0.26. Every deck passed the block checks, and none held a content violation. Seven decks sat off band, all at brackets 3 to 5, and the mana on turn four was the feature that missed most, seven times.

The judge agreed with the bracket on 8 of 15. It read every bracket 1 deck as a 2 or a 3. It read two of the three bracket 5 decks as a 3 or a 4.

The bracket 5 decks are the finding that matters. The shortlist ranks on theme and popularity. So a "cEDH" prompt gets a warrior deck at an average mana value of 3.4 and no fast mana. The bands and the judge both saw it, and no band move fixes it. That is the power signal of PR-14B (F-30, D-413).

The bracket 1 decks are the other side. The generator builds the strongest on-theme list the bands allow, and the judge reads that as a 2. No band moved on this run.

The judge invented one fact. It named Heliod, Sun-Crowned with Archangel of Thune as an infinite combo in three decks, and the endpoint lists no such pair. The judge bar reads the bracket alone, and its reasons are for a reader, never for a rule.

Deck gate run 12 (2026-09-02, `docs/reference/pr8-deck-gate-run12.md`). The regression run under the profile: 22 of 22 decks pass every block check, no invented name, no false rule, 63 calls, $2.24. Two prompts got no deck. The generate model passed its three-minute deadline twice on prompt 3 and twice on prompt 23, and the verdict reads FAIL on those errors alone. The repair turn ran on 11 decks, against 3 in run 11, and every reason was a profile finding. The owned-first decks still cost nothing to buy, and the tight budget deck rose from $16.51 to $20.84 under its $25 cap.

Run 12b (`pr8-deck-gate-run12b.md`) reran prompts 3 and 23 alone and passed both, in 266 seconds for $0.26. The errors were provider latency and not the prompt.

> *In plain English:* a bracket becomes a set of numbers the app builds to and checks, not a word it hopes the model understands. The app also deals ten thousand opening hands to see how fast the deck really gets going. The brackets are about how many turns a game lasts, so that number matters.

**PR-14B: The deck quality model (D-413 to D-417, split from PR-14 by D-451).** ✅ merged 2026-09-03 (#58, D-470 to D-493). Gate run 10 passes the top-list bar in every format and the precon bar in two of three. Deck gate 13b passes. The tier judge bar stays open on a corpus finding, and the owner merged with that on record (D-488, D-491).
The app holds no signal of what makes a deck good (F-30). The pool ranks on theme fit and popularity, and the bracket drops Game Changers under bracket 3 and nothing else. PR-14 builds a scorer that reads a deck and answers a quality tier and the named reasons. It covers Standard, Modern, and Commander, from bracket 1 to cEDH.

The data is every published list the sources of D-5 hold, back to 2015. Brackets 1 to 4 have no tournament data. The ladder gains the decks their owners tagged with a bracket on Moxfield, read through the public deck endpoint of D-419 (OQ-51 holds the check of that field and the terms).

The profile of PR-14A supplies the shape features.

The MTGO event pages embed the lists and the standings. The worker reads each League and Challenge per day, with the placement of each player.

The Topdeck.gg API gives the cEDH tournaments with each player's placement, wins, losses, draws, and win rate, and the decklists the organizers show (D-417). The app credits it with "Tournament data by TopDeck.gg" and a link, as it credits Scryfall (guardrail 7).

MTGTop8 gives the events with placements for Standard, Modern, cEDH, and Duel Commander, and the archetype shares, through a page reader. The cEDH Decklist Database gives the competitive tiers and the commanders. EDHREC gives the deck count per commander and the average deck.

PR-14B builds the precon table of D-407, and that table gives the bracket 2 baseline. PR-24 then reads the same table for the ownership check and the pool exclusion (D-460). The worker stores the raw page in GCS, with a normalized list per format per month beside it. A new parse then needs no new fetch.

The labels are the ladder of D-414: great, good, typical, baseline, and bad. A great list is a top-8 finish or a competitive-tier cEDH list. A good list is a league finish or the rest of a challenge. A typical list is the EDHREC average deck, and a baseline list is a precon.

A bad list is synthetic. The engine takes a real list and breaks one axis at a time: the lands, the curve, the colors, the copies, or the synergy. Each defect then carries its own label.

The features are what a person reads a deck by. Card quality is the inclusion rate of a card in the lists of its format and archetype, smoothed and weighted by placement. Synergy is the lift of a pair of cards in the great and good lists over chance. Shape is the curve, the land count, and the color sources against the pips. Roles are the counts of interaction, ramp, draw, and win conditions, and redundancy is how many cards fill each role.

A commander carries its own signal: the top-cut share in cEDH events at bracket 5, and the deck count at a lower bracket.

The model is one scorer per format, fitted in Go over the ladder by ordinal regression, with the weights stored beside the card snapshot. The scorer answers a tier and the three strongest signals in words, so the agent can say why. It runs on the deterministic side, and no prompt holds a weight. An LLM fine-tune on the same lists is a Phase 5 item, and the eval harness of PR-15 must justify it (D-413).

The score lands in five places. The commander pool ranks a bracket 4 or 5 request by the commander's cEDH signal, which closes OQ-48. The shortlist takes card quality as a signal beside theme and popularity. The generate prompt names the archetype shapes of the format. The deck summary names the tier and the three reasons. The revision turn says so when a change lowers the tier.

Contract, additive: `Deck.quality` with the tier, the score, and the reasons, and `CardService.GetCards` gains a quality field per card.

Gate:

- The scorer separates the ladder on a holdout of each format. A great list scores above a precon in 90 percent of the pairs, and a precon above a synthetic bad deck in 95 percent.
- A bracket 5 request offers three commanders from the top cuts of the Topdeck.gg cEDH tournaments of the last 90 days.
- The deck summary names the tier and three reasons for every golden deck. The judge of PR-15 agrees with the tier in 8 of 10.
- M-6: the parse failure rate per source per week stays under 1 percent, measured over the first month of the worker.

CAUTION: a page reader of MTGO or MTGTop8 breaks when the markup changes. The raw pages stay in GCS, so a fix re-parses and never re-fetches. M-6 reads the failure rate.

The cEDH database hosts its lists on Moxfield. The owner read the Moxfield terms on 2026-09-01, and they allow the fetch (D-419, closes OQ-49). The worker reads each linked list through the public deck endpoint, with a named agent and a slow rate. It stores the raw answer as it stores a page.

CAUTION: the Topdeck.gg key is a secret. It lives in `.env` locally and in Secret Manager on GCP, never in the repo. The worker refuses to start its Topdeck job without one.

What the build holds (2026-09-02). `internal/meta` reads the five sources: MTGO, MTGJSON, EDHREC, the cEDH database, and the Topdeck.gg API. Each reader is a parser over one page, and the tests read trimmed real pages.

The store keeps everything under `meta/` in the card bucket. It holds the raw pages, the lists per format and month, and the precon table per MTGJSON version. It also holds the commander reads per day and the model per version. `worker -meta` is the job, and `make meta-refresh` runs it on the local stack.

`internal/quality` is the model. It resolves a list against the card index and measures it through the profile of PR-14A. It breaks one axis per real list for the bad tier, and it fits a proportional odds scorer per format.

The scorer grades a deck, boosts the shortlist, and ranks a bracket 4 or 5 commander offer. It also writes the format shape block of the prompt and ends the summary with the tier and the reasons. `cmd/quality-gate` and `make quality-gate` write the gate document for free. `docs/reference/deck-quality-model-2026-09-02.md` holds every fact and every call.

A short live read of 2026-09-02 parsed 5 MTGO pages, 13 EDHREC pages, and the database page with no failure. It read 557 MTGJSON files before one connection reset. The fetcher retries a transport error once since.

CAUTION: the Moxfield lists are out of reach, so the database gives the tier and the commander alone (D-470). MTGTop8 waits (D-477). The Topdeck.gg read needs the key of OQ-54, and the bracket 5 bar of the gate reads FAIL until then.

Gate runs 1 to 8 (2026-09-02, `docs/reference/pr14b-quality-gate-run1.md` to `-run8.md`, free). The store held 6,187 MTGO lists, 21,146 cEDH lists over 90 days, and 701 precons. Run 1 put every top list over the precons. It put the precons over their broken copies in 0.16 to 0.46 of the pairs, because a broken top list still holds the staples. The owner moved the bad rung under the precons alone (D-484), and a defect detector joined the ladder (D-485).

Run 8 reads 0.96, 1.00, and 1.00 on the top-list bar and 0.94, 1.00, and 0.95 on the precon bar for Commander, Standard, and Modern. The bracket 5 offer names commanders with top cuts in the window.

CAUTION: the precon bar stands at 0.95, and runs 7 and 8 sit a hundredth under it in Commander and at it in Modern. The per-axis table of the document names the weak axes. No bar moved, and the next session reads that table before it touches the model.

Deck gate run 13 (2026-09-02, $2.64) read FAIL on two defects of the branch. Six curly apostrophes missed the name match, and fifteen rule-claim warnings came from the grade sentence. Both got their fix the same day (D-487). Run 13b (`pr8-deck-gate-run13b.md`, $2.62) passed 24 of 24 with no invented name.

The tier judge lane (`make quality-judge`, `pr14b-quality-judge-run1.md` and `-run2.md`, $0.31 each) read 5 of 24, then 4 of 24. The model grades most built decks below the precon baseline, and the judge reads them as typical. The explain mode of `quality-gate` shows why (D-488). The ladder holds 4,000 tournament lists per tier against 195 community decks, and the 60-card formats hold none. So a themed casual deck reads as off the ladder.

CAUTION: the judge bar is open on the corpus and not on the judge. The next session widens the casual middle of the ladder before it touches a weight. That means more EDHREC average decks, and a typical source for Standard and Modern, which no allowed source gives today. Read `quality-gate -explain <deck gate document>` first, for free.

> *In plain English:* the app learns what a good deck looks like from tens of thousands of real decks and how they placed. It learns the bad side too, from decks we break on purpose. Every deck it builds then gets a grade and three reasons, and a request for a top-power deck gets top-power commanders.

**PR-14C: MTGTop8 and the casual 60-card decks (D-482, D-490, D-493, D-502 to D-506).** ✅ merged 2026-09-03 (#60). `docs/reference/pr14c-sources-2026-09-03.md` holds every verified fact.
PR-14B dropped two sources the roadmap named. Moxfield answers 403 from Cloudflare to every plain client (D-470), and the MTGTop8 format page hides its event links behind a script (D-477). The owner let Moxfield go (D-493): Topdeck.gg holds the cEDH lists with placements, and the app asks the site for nothing. MTGTop8 comes back. 2026-09-03 correction: the format page holds its event links in plain HTML, in two tables of one row shape. D-477 read the wrong part of the page, and the reader takes the links from there.

The typical rung of Standard and Modern comes from the user decks of MTGGoldfish, a page reader, verified first (D-490). Aetherhub answers a Cloudflare challenge page to the app's agent, as Moxfield does, and the owner dropped that lane (D-502). The robots files of both sites carry a content signal against AI training, and the owner kept the legal check of D-5 over it (D-503). The ladder of the 60-card formats held tournament lists and precons and nothing between, and the judge bar fell on that gap (D-488).

The MTGTop8 lane reads the paper events of Modern, Standard, and cEDH (D-504). It walks the format pages, then the event pages with their placements and their field, then the text export of each deck. An MTGO event on the site duplicates the MTGO lane, and the reader skips it. A top-8 finish in a field of 32 or more is great, and the rest good (D-505). The raw store takes every page, and M-6 counts the failures.

Gate: each lane adds its lists to the store, the fit reads them, and the free quality gate shows the pair bars with them. The tier judge lane over deck gate 13b reads the bar of PR-14B that stayed open (D-491). No band and no bar moves for a source. 2026-09-03: the first run read 500 pages with zero failures, and quality gate run 12 reads FAIL on the two precon bars, on record, `docs/reference/pr14c-gate-2026-09-03.md`.
> *In plain English:* two sites the plan named are closed to a plain program today. We let them go, because other sites hold the same lists. The tournament site turned out open after all, and it brings the paper events with their placements. One more site brings the decks people build at home.

**PR-15: Eval harness.** ✅ merged 2026-09-04 (#65, D-511 to D-517), all six slices. The free gate is `docs/reference/pr15-gate-2026-09-04.md`. The paid gate ran on 2026-09-04 (D-513 to D-516). Every lane passes, and deck gate run 16, question gate run 35, and revise gate run 9 are the baselines. `docs/reference/pr15-paid-gate-2026-09-04.md` holds the read with F-36 to F-39. The owner answered OQ-57 to OQ-61 on 2026-09-04 (D-519 to D-522). The sweep runs on the owner's machine alone, and CI runs Tier 0. The four plan judge fields stand. A trimmed snapshot of at most 10 MB can join the repo, and all 47 terse conversations join the bar. The "Tier 1 nightly" line below is retired by D-519. A partial run under `-only` reads its item bars alone, and the check lists it and compares none (D-526, F-42).
Golden prompts with expected slot sets and expected validation outcomes. Deterministic checks are the gate (legality, ownership, size, curve, names). A judge role scores plan quality and usefulness on a fixed rubric. Long-format results table, corpus fingerprint per run (model, effort, snapshot date, prompt version), suffix rows for informational metrics, "observe-only is not pass". Tier 0 in CI ($0).

Tier 1 nightly. Label-gated full sweep on PRs. Cost cap per run. Gate: the harness runs on PR-8's output and reports named regressions.
> *In plain English:* the test bench. Fixed questions, expected answers, a score every night. Any change that makes decks worse is named, not averaged away.

**PR-24: Precon exclusion (D-407 to D-409).** ✅ merged 2026-09-03 (#59, D-496 to D-498, D-500, D-501). The free gate is `docs/reference/pr24-precon-gate-2026-09-03.md`. Deck gate prompt 25 covers the gate line since 2026-09-04 (D-510), and it waits for a paid run.
A reader asks for a deck that uses no card of a precon they own, for any set. The worker reads the MTGJSON deck list once per MTGJSON version and stores a precon table beside the set file (D-377 shape). A row holds the product name, the set code, the type, the release date, and the cards as printing ids with counts. The table holds the Commander decks and the 60-card constructed decks, and leaves out Jumpstart packs, Welcome decks, Secret Lair drops, and land packs (D-407).

A reader owns a precon when the collection holds every printing of it with its count (D-408). The check runs at build time from the collection. The exclusion subtracts the precon counts from the owned counts, so the surplus copies stay usable. `candidates.Request` gains an exclusion list, the commander pool and the 99 both drop it, and the rules check blocks any card that slips through.

A classify field and a classify fact read "not from precon X" and "not from my precons", as the set filter reads its names (D-496). A catalog row asks about a product name the table does not hold. The chat names the precons it excluded (D-390). The build excludes a named precon the collection does not hold whole, and the chat says so (D-497). The nine embedded lists of D-247 stay as the upgrade source, and a test compares each one with the table (D-498).

`docs/reference/precon-data-2026-09-01.md` holds the verified facts. Gate: a build for a collection that holds Avengers Assemble whole uses none of its cards. A build for a collection that holds 99 of its 100 is free to use any of them.
> *In plain English:* the app learns every deck Wizards ever sold. When you own one whole, you can ask for a deck that leaves it untouched. The app checks your binder to know which ones you own.

**I-1: Ban-list watch, stale-deck banner, and scoped rerun (D-29).**
A job reads the Wizards announcement feed and detects the Scryfall snapshot that reflects it. It then re-validates every stored deck in the affected formats. 

A deck with a now-illegal card gets a `stale` flag with the list of affected cards. The UI shows a banner on that deck with a "rerun" button.

An impact classifier scopes the rerun. Its inputs: how many cards the change touches, which roles they filled, and whether the commander or a win condition is among them. 

Low impact: a patch turn that replaces only the affected cards from the same candidate list. 

High impact (threshold OQ-18): a full rebuild with the original slots and a new seed. The banner states which case applies and why. Gate: on the golden decks, every synthetic ban produces the correct case and a legal deck.
> *In plain English:* when Wizards bans a card, every deck we built that uses it gets a warning and a rerun button. If the ban only touches one filler card, we swap that card. If it guts the deck, we rebuild it from your answers.

**I-2: Price-aware buy list** (D-17, F-16). USD. Each card carries the lowest Scryfall NM market price across legal printings and finishes. It is a 7-day rolling average, and the average rejects an outlier day. An outlier is a day more than 2x the 7-day median (D-26). Digital-only and gold-bordered printings excluded. The UI labels it "NM market estimate" with the price date.

**I-3: Semantic card search** over Oracle text as a fourth candidate signal, only if PR-6's gate shows tags are not enough.

### Phase 5 - Parked (product decisions required)

- ~~Sample-hand and~~ goldfish simulator (D-20: later, not at launch). The sample hand left the lot on 2026-08-29 and sits in PR-20 (D-318). The goldfish simulator stays here.
- Per-card explanations longer than one line (D-19 gives one line per card).
- An LLM fine-tune on the labeled deck lists of PR-14 (D-413). The eval harness of PR-15 must show the scorer falls short first.
- Non-English collections (D-23: English only for now).
- Brawl, Oathbreaker, Pauper Commander, Duel Commander, Canadian Highlander.
- Pioneer, Legacy, Vintage, and Pauper. The app built these until 2026-08-26, and D-155 removed them. The agent declines each one by name and offers no substitute.
- Sideboard builder for 60-card competitive play against a named meta.
- Collection sync from ManaBox without a file (no API exists on 2026-08-23).
- A public corpus API that serves the `mtg-corpus` content to the app's own prompts.

## 8. Sequencing - strict order, single owner

1. PR-0a scaffold.
2. PR-0b machine setup, Docker install (owner executes).
3. PR-0c local stack.
4. PR-1 proto v1.
5. PR-2 card database.
6. PR-3 legality freshness, with M-2.
7. PR-4 ManaBox import. Then M-3.
8. PR-5 rules engine.
9. **GATE.** Phase 2 starts only when the golden decks pass PR-5. Held 2026-08-24.
10. PR-10 LLM role layer, with M-1.
11. PR-1b contract amendment (audit branch, D-46).
12. PR-6 candidates.
13. PR-7 questions.
14. PR-7B automated eval lane. It runs beside PR-8 once its three evals hold.
15. PR-8 generator.
16. PR-9 variance. ⏸ out of MVP scope (D-256). It blocks nothing: the Phase 3 gate below reads PR-8's gate.
17. **GATE.** Phase 3 starts only when PR-8's gate holds on the golden prompts. ✅ held on 2026-08-28, deck gate run 6.
18. PR-11 ✅ merged 2026-08-28 (#38). PR-12 ✅ merged 2026-08-28 (#40). PR-12B ✅ merged 2026-08-29 (#41). PR-13 ✅ merged 2026-08-29 (#45). Then Phase 3B.
19. **Phase 3B** (D-316, D-317): PR-16 to PR-23 in the order of the phase list. PR-14A and PR-14B sit between PR-19 and PR-20 (D-452, D-460). PR-14A ✅ merged 2026-09-02 (#57). PR-14B ✅ merged 2026-09-03 (#58). PR-19 ✅ merged 2026-09-02 (#55, #56). PR-20 ✅ merged 2026-09-03 (#61). PR-21 ✅ merged 2026-09-03 (#62). Each gate holds before the next slice starts. PR-16 ✅ merged 2026-08-29 (#47). PR-16B ✅ merged 2026-08-29 (#48). PR-17 ✅ merged 2026-08-31 (#49). The paid re-baseline of D-302 ran on 2026-08-31.
20. **PR-17B** the set filter (F-29, D-373 to D-383). ✅ merged 2026-09-01 (#50). **PR-18** ✅ merged 2026-09-01 (#53). The review fixes of PR-18 (D-398 to D-406) 🔧 built 2026-09-01 on branch `nits-and-fixes`. The owner reads them, then merges. Then PR-19.
21. PR-15 eval harness. ✅ merged 2026-09-04 (#65), after #64 (D-511), before PR-22 and PR-23. The paid gate ran the same day and passes on every lane (D-516). M-5 manual scoring runs on the first UI build (after PR-12).
22. PR-24 precon exclusion (D-409, D-460) ✅ merged 2026-09-03 (#59). PR-14C (D-482) ✅ merged 2026-09-03 (#60). Then I-1, I-2, I-3 on evidence. PR-14B moved into step 19 (D-460).
23. Phase 5 stays parked.

## 9. Open questions

See `docs/open-questions.md` for the full list with "ask when" dates. The ones that gate a phase:

1. **OQ-19 scoring rubric** answered 2026-08-24 (D-66). M-5 is no longer gated on it.
2. **OQ-18 rerun depth rule** gates I-1.
3. **OQ-20 public anonymized ManaBox exports** closed 2026-09-01 (D-431). The fixture set is the owner's export and a generator from the card snapshot.
4. PR-9's 30% variance number is a placeholder until PR-15 measures it. PR-9 is out of the MVP (D-256), so nothing waits on it.
5. **OQ-45 the allowlist store** answered 2026-09-01 (D-420): one Firestore document, written by a make target. The old text stays below. D-314 allows one env var or one Firestore document. An env var needs a deploy per change, and a document needs an admin write path. PR-22 decides, and the owner confirms. Ask before PR-22.
6. **OQ-46 the spend cap number** answered 2026-09-01 (D-421): $5 per user per month. The old text stays below. PR-22 sets a per-user monthly cap from `Usage`. The number is the owner's. Ask before PR-22.
7. **OQ-47** answered 2026-08-29 (D-324). The deck grid filters by power, and the filter runs on the server. A second question took the same id. D-376 answered it on 2026-08-31: a set name resolves to a whole set family.
8. **OQ-49 the Moxfield terms** answered 2026-09-01 (D-419). The owner read them, and they allow the fetch.
