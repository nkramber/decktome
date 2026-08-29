# Open questions

Questions not yet asked, or asked and not yet answered. Move each answered question to `decisions.md`. Ask in small batches when the design reaches the point that needs the answer.

## Not yet asked

| # | Question | Why it matters | Ask when |
|---|---|---|---|
| OQ-18 | Rerun depth rule: when does a ban trigger a full rebuild instead of a patch? | D-29 asks for a rerun scoped to the nature of the change. The rule needs a threshold (for example: a banned commander or win condition means a full rebuild). | Before I-1 ships. |
| OQ-20 | Where can we find public anonymized ManaBox exports for the import fixture set? | D-43. One real export is the gate today. More files widen column and value coverage. | Before PR-13, when time allows. PR-11 merged on 2026-08-28 without it. |
| OQ-44 | What is the ManaBox condition vocabulary? The import knows `near_mint` from one export. | `docs/audit-2026-08-24.md` section 8 left it open. One export with a played card confirms the other values. | When the owner has such an export, before PR-13. PR-11 merged on 2026-08-28 without it. |

| OQ-45 | Where does the email allowlist live: one env var, or one Firestore document with an admin write path? | D-314 allows both. An env var needs a deploy per change. | Before PR-22. |
| OQ-46 | What is the per-user monthly spend cap on GCP? | PR-22 refuses a turn over the cap, and the number is the owner's. | Before PR-22. |

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
