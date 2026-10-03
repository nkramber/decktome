# Response to the review of pull request 271

The author answers the Codex record of `docs/reviews/pr-271.md`. That record read `Changes required` at head `b86717d`, with one finding.

## P2-1: The live-test instructions understate share-link creation

- The result: full merit for the text, and no merit for a new test.
- The evidence: the share step of `web/apps/web/live/sweep.spec.ts` clicks "Make a link" when the deck has no stored link, and no step ends the link. The skill line named the reuse alone, and the file header still said that the sweep revokes the link.
- The correction: `.claude/skills/live-test/SKILL.md` now says that the sweep reads the stored link or makes one. The link stays live until its deck ends (D-1069). It also names the case of `LIVE_SWEEP_BUILD=0`, which can make a public link to an existing deck. The header of `web/apps/web/live/sweep.spec.ts` states the same behavior.
- The regression check: a search for "revok" in `web/apps/web/live` and `.claude/skills/live-test` finds only the two lines that say the app has no revoke. No unit test runs the live sweep, because the sweep needs the deployed app and the check account. `make live-sweep` is a paid target. So the next sweep of the owner tests the free-run case, and this pull request runs no sweep.
