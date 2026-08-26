# Open questions

Questions not yet asked, or asked and not yet answered. Move each answered question to `decisions.md`. Ask in small batches when the design reaches the point that needs the answer.

## Not yet asked

| # | Question | Why it matters | Ask when |
|---|---|---|---|
| OQ-18 | Rerun depth rule: when does a ban trigger a full rebuild instead of a patch? | D-29 asks for a rerun scoped to the nature of the change. A threshold is needed (for example: commander or win condition banned means full rebuild). | Before I-1 ships. |
| OQ-20 | Where can we find public anonymized ManaBox exports for the import fixture set? | D-43. One real export is the gate today. More files widen column and value coverage. | Before PR-11, when time allows. |
| OQ-21 | How much of a named precon must the built deck keep? | D-113 names the precon in the card-pool question. It does not say what "upgrade my precon" means to the generator. A share of the precon cards must survive the build, and no number exists. | Before PR-8 builds a deck. |
| OQ-23 | What should the agent do when the user asks a question back, such as "What is a bracket?" | Probe 36 asks it. The agent answered with silence in every run, because the catalog holds questions and no answers. A row that explains a term, or a separate answer path, is a design decision. | Before PR-12 builds the chat UI. |
| OQ-22 | Which supported format is nearest to Historic, and which to Timeless? | D-112 names the nearest format we build. Brawl, Oathbreaker, Duel Commander, and Canadian Highlander map to Commander, and Alchemy maps to Standard. Historic maps to Modern and Timeless maps to Legacy, and both are unverified. | Before the first user sees the row. |

## Asked, waiting

None.


## Answered (moved to decisions.md)

OQ-1 to OQ-12 were answered 2026-08-23. See D-15 to D-25. OQ-13 to OQ-17 were answered 2026-08-23. See D-26 to D-30. OQ-17 closed 2026-08-24 (D-43). OQ-19 answered 2026-08-24 (D-66).
