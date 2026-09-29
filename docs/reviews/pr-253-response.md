# Response to the review of pull request 253

The author answers the Codex record of `docs/reviews/pr-253.md`. That record read `Changes required` at head `f7f4f7e`, with one finding.

## P2-1: Cold token and invite reads can exceed the seven-second limit

- The result: full merit.
- The evidence: the certificate client of `go/cmd/api/main.go` waits up to 3 seconds. The list read had a deadline of its own of 5 seconds. The two waits come one after the other, so a cold answer took up to 8 seconds. Identity Platform fails the sign-up after 7 seconds (D-990).
- The correction: `authblock.Budget` of 5 seconds is one deadline for the whole answer. The token check and the list read both read that context. The key fetch reads the same context, so the deadline also stops a fetch that waits behind the key cache lock.
- The regression check: `TestOneDeadlineBoundsTheKeyAndTheList` sets a budget of 150 ms. Its key takes 100 ms, and its list waits for the deadline. The answer is 503 inside the budget. With the list read on a deadline of its own of 5 seconds, the same test fails after 5.4 seconds. `TestTheBudgetLeavesTimeInsideTheLimit` keeps `Budget` at 6 seconds or less.

## P2-2: The Gitar gate accepts a dashboard that says a review is in progress

- The result: full merit.
- The evidence: the `gitar-review` skill names two marks of a review in progress, the spinner and the line "Responding to your feedback". The check read the spinner image alone.
- The correction: `RESPONDING` reads the status line at the start of a line or of a `<kbd>` element. A quoted finding does not start the line with it, so a quote does not count. D-992 names the rule.
- The regression check: `test_a_status_line_with_no_spinner_fails` fails on `d2c0e00` and passes with the fix. `test_a_finding_that_quotes_the_status_line_passes` keeps a quote out. The finished dashboard of #253 reads as not in progress.

## P2-3: The Gitar gate misses valid spinner image tags

- The result: full merit.
- The evidence: the spinner pattern needed a `src` value in double quotes. An image tag with single quotes or with no quotes did not match. Gitar used double quotes on each dashboard of #253, but D-992 names each image tag of the spinner.
- The correction: `SPINNER` reads any `<img>` tag that names `gitar-spin.svg`. A finding that names the file in prose has no image tag, so it does not count.
- The regression check: `test_the_spinner_fails_with_any_quote_style` fails on `8c30fa2` and passes with the fix. `test_a_finding_that_names_the_spinner_in_prose_passes` still passes.
