# Response to the review of pull request 267

The author answers the Codex record of `docs/reviews/pr-267.md`. That record read `Changes required` at head `3ccc917`, with one finding.

## P2-1: A stopped reading turn is described as a deck build

- The result: full merit.
- The evidence: PR-110 replaced the line of `server-build` with the line of `server-build-stopped`. The old line chose a reply message when `serverReply` was true. The new line lost that choice. So a Stop during `READING` promised a deck, and the turn can return questions (D-375).
- The correction: `web/apps/web/src/features/chat/session-page.tsx` selects the line from `serverReply` again. A reading turn shows the reply message, and a build shows the deck message.
- The regression check: the new test of `session-page.test.tsx` stops a first turn in `READING`. It checks the reply message and the absence of the word "deck". The test fails without the correction, and the 81 tests of the file pass with it.
