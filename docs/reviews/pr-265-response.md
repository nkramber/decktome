# Response to the review of pull request 265

The author answers the Codex record of `docs/reviews/pr-265.md`. That record read `Changes required` at head `673c86f`, with one finding.

## P2-1: A hidden mount never reports that the page is hidden

- The result: full merit.
- The evidence: `usePageHidden` of `web/apps/web/src/features/chat/use-page-hidden.ts` started `reported` from the visibility of the page at mount, and it sent no request then. A mount in the background sent nothing until a later change of visibility. In the path of F-190, the earlier mount reports the leave first, so the record holds it. A first mount in the background has no such report, for example after a reload or a remount after the push of an update. Then `pageLeft` reads no record, and the build sends no push (D-1033).
- The correction: the hook starts `reported` at false, and it runs its change handler once at mount. A mount in the background then sends `hidden:true` at once, and the return to view sends `hidden:false`. A mount in view sends nothing. A remount in the background writes the same state again, with a later time, which changes no push.
- The regression check: "a page that mounts in the background reports the leave and the return (D-1033)" of `web/apps/web/src/features/chat/session-page.test.tsx` reads the two reports in order. It fails with the call at mount removed, and passes with it. The chat tests pass, 113 of 113.
