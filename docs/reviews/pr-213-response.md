# Pull request 213: the answer to the review

The author answers each finding of `docs/reviews/pr-213.md` here, with the procedure of `.claude/skills/pr-review/references/answer-review.md`.

## P2-1: The Gitar gate reads only the first 100 review threads

- Result: full merit.
- Evidence: the query of `docs/tools/codex_review.py` read `reviewThreads(first: 100)` and no `pageInfo`. So a thread after the first 100 never reached `gitar_problems`. The ruleset of `main` still refuses a merge with an open thread (D-828). So the fault reached the start of a review, and not the merge.
- Correction: `review_threads` asks for `pageInfo` and takes `$endCursor`. It runs `gh api graphql --paginate --slurp`, and it joins the threads of each page (D-832).
- Regression check: `test_an_open_thread_on_a_later_page_fails_the_pass` gives 100 resolved threads on page 1 and one open thread on page 2. It fails on the old code, and it passes now. A live read of #213 through `review_threads` gave 2 threads and 0 open.
