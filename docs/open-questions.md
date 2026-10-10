# Open questions

Questions not yet asked, or asked and not yet answered. Move each answered question to `decisions.md`. Ask in small batches when the design reaches the point that needs the answer.

## Not yet asked

D-1188 answered OQ-91.

| # | Question | Why it matters | Ask when |
|---|---|---|---|
| OQ-96 | The gap question read "Your collection holds 0 planeswalkers cards, and I want 30 or more." Should the sentence name the theme in its singular form, and say that the count reads cards that fit the theme? | The live eval of deck `dDD9Iav9aGhgGF9dDLNQ` read it. The set of the request holds 10 planeswalkers, so a reader can read 30 as 30 planeswalkers. The count reads each card with a theme signal, and the set holds more than 30 such cards. `Gap.Sentence` in `go/internal/questions/pool_gap.go` puts the theme before "cards" as the reader wrote it. | When a live eval or a verdict names the sentence again, or with the next change of the gap question. |
| OQ-99 | The request "the best possible commander deck" put the power words "best possible" in the theme slot, in each replay of deck `v-8k83YGt2Dn6IddKn3CBT`. Should a word that fills the power step stay out of the theme, as the filler words of D-1187 do? | The theme reads no card for "best possible", so the build has no real theme. The superlative already fills bracket 5 (D-1154). | When the next item changes the theme words. |
| OQ-100 | A theme that matches most of the pool fills the total cap of the shortlist, and the score cut drops each nonbasic land past the mana half. The Jodah replays of F-229 offered 27 nonbasic lands in place of 40, and the decks held 22 to 27 fixing lands in place of 33 to 35. Should the land cap keep its full count before the score cut? | A five-color deck wants fixing. The floor of 17 still holds, so no check failed. | When the next item changes the shortlist caps. |
| OQ-101 | The replays of F-229 held 45 to 47 legendary cards from the collection, and the quality tier fell from typical to baseline in two of three. The grade said that few cards are ones the top lists play. Is that the grade of OQ-97, or a fault of the build? | The reader wants a Jodah deck of legends, and the grade reads a lower deck. | With OQ-97. |
| OQ-102 | An irregular verb such as "gave", and a verb that no stop list holds such as "wishing", still reach the theme row of D-725 when no card holds them (F-230). Should the classifier write the theme words alone, or should the matcher read a dictionary of verbs? | D-1222 reads the regular forms of a stop word alone. A list of each English verb grows one chat at a time, which D-1183 forbids. | When the next live eval or gate run names such a word. |
| OQ-103 | The theme "treasure" puts 361 cards on theme, and "treasure tokens" puts 3,914 on theme on the snapshot of 2026-09-04. The word "tokens" finds the go-wide tokens row. Should "treasure tokens" read the treasure row alone? | A treasure deck can fill with creature-token makers. No deck of the eval of F-230 proves the effect. | With the next item that changes the theme rows. |
| OQ-104 | A filler word that card text holds widens the theme in silence. "treasure tokens using" puts 3,943 cards on theme, and "treasure tokens win" puts 4,120, against 3,914 for "treasure tokens". Should a generic text needle need a minimum share of on-theme cards? | D-1188 names this half of the filler class, and its lists do not hold "using" or "win". | With OQ-102. |
| OQ-105 | The theme row writes the theme with no quote marks: "matches "making" in the theme making lots of treasure tokens and using them to win." Should the row quote the theme? | The sentence is hard to read. A change of the row needs a question gate run (D-1116). | With the next question gate run. |
| OQ-106 | The ask role wrote "Which format would you like?" and "Which Commander bracket should the deck target?" with no options, in three of the six fix replays of F-230. Do the buttons of the chat screen show the options in that case? | A reader with no buttons can not see the five brackets. Unverified on the screen. | With the classifier item of D-1211. |

OQ-97 waits in `docs/owner-questions.md`, because a change of the quality model needs the owner.

## Asked, waiting

No question waits here now. D-1004 answered OQ-67, and D-990 answered OQ-77.

## Deferred by the owner

These questions wait for work, and not for an answer of the owner.

| # | Question | Why it matters | Ask when |
|---|---|---|---|
| OQ-85 | A cap for each effect class of the shortlist, weighted by what the commander cares about. | D-773 defers it: the class weights need the payoff shape, and the shortlist has no shape score. `docs/reference/grima-payoff-shape-2026-09-20.md` holds the count. | When the shortlist reads a shape score. |


## Answered (moved to decisions.md)

- OQ-1 to OQ-12: answered 2026-08-23 (D-15 to D-25).
- OQ-13 to OQ-17: answered 2026-08-23 (D-26 to D-30). OQ-17 closed 2026-08-24 (D-43).
- OQ-19: answered 2026-08-24 (D-66).
- OQ-21: answered 2026-08-26 (D-218).
- OQ-22: answered 2026-08-26 (D-146). Historic and Timeless name no nearest format.
- OQ-24 to OQ-27: answered 2026-08-26 (D-135 to D-138).
- OQ-32 to OQ-35: answered 2026-08-26 (D-131, D-129, D-130, D-132).
- OQ-36: answered 2026-08-27 (D-226, D-232).
- OQ-38: answered 2026-08-27 (D-239).
- Two questions carry the id OQ-40, and neither gets a new number. D-139 closed the first OQ-40 on 2026-08-26 (the loop starts from `main`). D-247 closed the second on 2026-08-28 (the precon data source, opened by D-240).
- OQ-41 to OQ-43 answered 2026-08-26 (D-151, D-152, D-154).
- Two questions carry the id OQ-47, and neither gets a new number, as with the two OQ-40 rows. D-324 closed the first on 2026-08-29 (the deck grid filters by power). D-376 closed the second on 2026-08-31 (a set name resolves to a whole set family).
- OQ-45, OQ-46, and OQ-49 answered 2026-09-01 (D-420, D-421, D-419).
- OQ-44 answered 2026-09-01 (D-429). OQ-23, OQ-28 to OQ-31, OQ-37, and OQ-39 answered the same day (D-422 to D-428), and their rows left `docs/owner-questions.md`.
- OQ-20 closed 2026-09-01 (D-431): the owner's export and a generator from the snapshot.

- OQ-50 answered 2026-09-02 (D-459): the Commander Spellbook terms allow the bracket endpoint, at 90 requests a minute on the app's side.
- OQ-52 answered 2026-09-02 (D-467, D-468): the session calls of PR-14A stand, and Spellbook reads the shortlist too.
- OQ-53 answered 2026-09-02 (D-469): the engine's Commander land range is 27 to 41.
- OQ-48 answered 2026-09-02 (D-476): a bracket 4 or 5 offer ranks on the cEDH signal of the quality model.
- OQ-55 closed 2026-09-03 (D-493): the app asks Moxfield for nothing, and PR-14C drops the lane.
- OQ-54 answered 2026-09-02 (D-479): the owner created the Topdeck.gg key, and a live probe of the API passed the same day.
- OQ-51 closed 2026-09-02 (D-470): no session can check the Moxfield bracket field. The deck page, the v2 API, and the v3 API answer 403 to a plain client, so PR-14B reads no Moxfield list.
- OQ-81 answered 2026-09-08 (D-611): the low-effort measurement of the generate role runs before PR-33, against the run 16 baseline.
- OQ-91 answered 2026-10-05 (D-1188): the word "set" matched card text and widened the shortlist in silence. It now stays out of the theme. The set limit reads the set name of the classifier.
- OQ-90 answered 2026-10-01 (D-1033): the page tells the server when it goes to the background.
- OQ-18 answered 2026-09-30 (D-1008): a ban of the commander, of a win condition, or of 10% of the nonland slots causes a full rebuild.
