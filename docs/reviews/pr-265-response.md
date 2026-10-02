# Response to the review of pull request 265

The author answers the Codex record of `docs/reviews/pr-265.md`. That record read `Changes required` at head `673c86f`, with one finding.

## P2-1: A hidden mount never reports that the page is hidden

- The result: full merit.
- The evidence: `usePageHidden` of `web/apps/web/src/features/chat/use-page-hidden.ts` started `reported` from the visibility of the page at mount, and it sent no request then. A mount in the background sent nothing until a later change of visibility. In the path of F-190, the earlier mount reports the leave first, so the record holds it. A first mount in the background has no such report, for example after a reload or a remount after the push of an update. Then `pageLeft` reads no record, and the build sends no push (D-1033).
- The correction: the hook starts `reported` at false, and it runs its change handler once at mount. A mount in the background then sends `hidden:true` at once, and the return to view sends `hidden:false`. A mount in view sends nothing. A remount in the background writes the same state again, with a later time, which changes no push.
- The regression check: "a page that mounts in the background reports the leave and the return (D-1033)" of `web/apps/web/src/features/chat/session-page.test.tsx` reads the two reports in order. It fails with the call at mount removed, and passes with it. The chat tests pass, 113 of 113.

## The repeat review: Blocked on the live check

- The result: the record names no finding. It reads `Blocked` at `e9e2844`, because the gate of PR-108 listed the live check on the iPhone of the owner.
- The evidence: that check needs the new API and web on Cloud Run, and production deploys from `main` alone (D-579). So no pull request can hold it before its merge. D-1007 gave the same order to the push check of PR-26: the check runs after the deploy, and it does not stop the merge.
- The correction: the PR-108 entry of `docs/design-roadmap.md` moves the live check out of the gate list. It now reads as the live check after the deploy, in the words of PR-26, and it cites D-1007. The hand-off keeps it as the first next step.
