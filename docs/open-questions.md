# Open questions

Questions not yet asked, or asked and not yet answered. Move each answered question to `decisions.md`. Ask in small batches when the design reaches the point that needs the answer.

## Not yet asked

| # | Question | Why it matters | Ask when |
|---|---|---|---|
| OQ-18 | Rerun depth rule: when does a ban trigger a full rebuild instead of a patch? | D-29 asks for a rerun scoped to the nature of the change. The rule needs a threshold (for example: a banned commander or win condition means a full rebuild). | Before I-1 ships. |
| OQ-48 | Which source ranks a commander by power for a bracket 4 or 5 request? | The app holds no power signal. The bracket drops Game Changers under bracket 3 and nothing else, so a bracket 5 request gets the most popular commanders (D-411). PR-14 answers it with the commander's cEDH signal (D-413). | PR-14 closes it. |

## Asked, waiting

OQ-23, OQ-28 to OQ-31, OQ-37, and OQ-39 sit in `docs/owner-questions.md`, the decision queue. This file does not repeat them.


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
