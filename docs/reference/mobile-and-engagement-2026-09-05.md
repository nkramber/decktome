# Mobile and engagement, the proposal of 2026-09-05

This note answers the owner's request of 2026-09-05: a whole solution for the phone and for engagement, before any cloud PR. It reads the product as it stands, and it names the platform facts of 2026-09-05 with their sources. It proposes three stages with a cost for each. The owner decides the stages through OQ-66 to OQ-68.

## 1. Where the product stands on a phone

The web app is responsive by design (D-311), and one shell serves a phone and a desktop (D-328). Eighteen files carry phone breakpoints, the page has a viewport tag, and the deck view opens its card detail as a sheet (PR-20). The collection import is a file input with a drop zone (`upload-dialog.tsx`), and the share page at `/d/<token>` opens for anyone with the link (D-315). A print stylesheet renders a deck as a list in black on white (PR-21).

Four things are absent. The app has no web manifest, no icons, and no service worker. A phone can not install it, and it shows nothing without a network. The Phase 3B gate still holds one open item: "The owner walks the whole path on a desktop and on a phone" waits (PR-16). No channel reaches a user who left the page: no push, no email. Nothing counts use: the `Usage` message counts model tokens per session and nothing else.

## 2. The phone is the first device, not the second

ManaBox is a phone app on iOS and Android. It exports the collection as a CSV file from the collection tab or from one binder. The collection therefore lives on the phone. A player builds at the table, on a couch, or in a shop, with the phone in hand. A game night puts the phone beside the deck: the list by role, the sample hand, and the bracket are the play aids. The desktop is the place for a long chat and a big binder view, and the phone is the place the product gets opened.

Two platform facts shape the phone plan, both read on 2026-09-05:

- A site added to the Home Screen on iOS 26 opens as a web app by default. A Home Screen web app on iOS 16.4 or later can receive web push. Push does not work inside Safari itself on iOS. Apple withdrew its EU removal of Home Screen web apps on 2024-03-01, so the EU runs the same path.
- iOS does not implement the Web Share Target API, so a web app can not appear in the ManaBox share sheet on an iPhone. Android does, through an installed web app. On an iPhone the export goes to Files, and the user picks it from the upload dialog.

## 3. The solution, in three stages

### Stage A, the installable web app, PR-25

The proposal of 2026-09-05 put Stage A inside PR-22. The owner made it its own PR on 2026-09-06 (D-547), after PR-23 in the phase list. The deploy of PR-22 makes the phone reachable, and PR-25 makes the app installable.

1. A web manifest with the name, the mark as icons in the required sizes, `display: standalone`, and the dark theme color. A service worker caches the app shell. A cold open with no network then shows the shell and the last deck list. `vite-plugin-pwa` writes both from the Vite build.
2. An install hint. On iOS the user taps Share, then Add to Home Screen. The app shows the hint once, on a phone, after the first deck, and never again after a dismissal.
3. The import for a phone. The upload dialog accepts `.csv` from the Files picker. A paste box takes the CSV text for a user whose share sheet offers Copy. The dialog says where ManaBox puts the export.
4. The phone gate becomes explicit. Every route at 390 pixels wide, and every touch target at 44 pixels or more. The chat input above the keyboard, the card sheet, and the share page. The owner's walk closes the PR-16 item.

Cost: $0 in services, about one day of work. It changes no contract.

### Stage B, the return channels, after PR-23

Push and email carry three events, and each user opts in per event. Every event comes from data the product already keeps, so the channels add no model spend on their own.

| Event | Source in the product | Message |
|---|---|---|
| A ban or a rules change touches a deck of yours | The announcement calendar of PR-3 and the legality diff of the snapshot, then the `stale` flag of the ban rerun plan | "Wizards banned X. Your Karlov deck holds it. Rerun?" |
| A new set holds cards for a deck of yours | The set table of the snapshot with its release dates, and the theme search of PR-6 over the new cards against each saved deck | "Three new cards fit your dinosaur deck." |
| Your build finished | The session state, for a build the user left | "Your Modern burn deck is ready." |

Web push runs through Firebase Cloud Messaging, which is free, with a service worker and a VAPID key from the Firebase console. Email runs through one provider with a free tier: Brevo gives 300 emails a day, and Resend gives 3,000 a month. Five invited users at one weekly digest each is 20 emails a month. The digest is one email a week: the decks touched by a legality change, the new cards per deck, and one suggested revision. A tap on the revision runs one revise turn, about $0.10, and only on the tap.

Cost: $0 in services at this scale. The work is one Cloud Run job for the weekly read and the service worker. One Firestore document per user holds the opt-ins and the device tokens. The ban rerun itself is the roadmap item after PR-22, and the channel rides on it.

### Stage C, the stores, only on request

Android accepts a web app in Google Play as a Trusted Web Activity, packaged with Bubblewrap or PWABuilder, with a $25 one-time developer fee. It adds nothing a Home Screen install lacks, so it waits until a user asks for the store.

Apple accepts no web app as such, and a wrapper with no native value fails App Store guideline 4.2. A native shell earns its place with three things the web can not do on iOS. They are a share extension that receives the ManaBox CSV from the share sheet, and a widget with a deck of the week. The third is push through APNs without the Home Screen step. Apple's Developer Program costs $99 a year. This stage waits until the invited users are more than a handful and ask for it.

## 4. Engagement, what it means here

This is a tool a player opens around events, not a game a player opens every day. Engagement means that the product is there at each event with something useful, and that the user comes back to it on their own. The events: a purchase, a new set, a ban, a game night, and a friend's request.

| Loop | The trigger | What the product does | What exists |
|---|---|---|---|
| The collection loop | The user buys cards and re-exports ManaBox | "Since your last import: 40 new cards, and two decks improve." One tap re-imports, and the new commander offers read the delta. | The import and the summary (D-392). The delta needs the previous import's counts, one small document. |
| The deck health loop | Wizards changes a ban list, about four times a year | The banner and the rerun button of the ban rerun plan, with the push or the email of Stage B | The calendar of PR-3, the legality diff, the `stale` flag in the plan |
| The new set loop | A set releases, about every six weeks | Three new cards per saved deck, and a one-tap revision | The set table of the snapshot, the theme search, the revise turn of PR-12B |
| The table loop | A game night | The deck view as a play aid: the list by role, the sample hand, the bracket, the game changers and the combos of the Spellbook check, and the share link for the table | The deck view of PR-20, the share page of PR-21, the Web Share API on iOS |
| The meta loop, 60-card | The weekly meta refresh of PR-14C | "Your Modern deck reads typical this week, down from good." | The quality model and its tier |
| The friend loop | A friend asks for the list | The share link, and the export in Arena text | PR-21 |

Two things stay out. No streaks, badges, or daily rewards. They fit a game and not a tool, and they push a user to open an app with nothing new in it. No third-party analytics or tracking pixel. The users are friends with an invitation, and the counts below suffice.

## 5. What to measure

Six counts, one Firestore document per user per week, written by the API and read by a `make engagement-report` target:

- Weekly active builders, and decks per active builder. The estimate of `docs/setup-gcp.md` says three a week.
- Re-imports per user per month, and the delta they carried.
- Revisions per deck.
- Opt-ins per channel, and taps per notification sent.
- Share links made, and opens per link. The rate-limited public read counts them.
- Installs: the `display-mode: standalone` media query tells the app it runs from the Home Screen.

The counts hold no card data and no message text, so the public page rule of guardrail 13 stays whole.

## 6. Costs

| Stage | Services a month | Once |
|---|---|---|
| A, the installable web app | $0 | About a day of work |
| B, push and email | $0 at five users: Cloud Messaging is free, and the email digest sits inside a free tier | One Cloud Run job and a service worker |
| C, Android store | $0 | $25 developer fee, a packaged build |
| C, iOS native shell | $0 | $99 a year, and a native project to keep |

The model spend of Stage B is one revise turn per tap, about $0.10, inside the $5 cap of D-421.

## 7. The order

1. PR-22, the deploy alone (D-547).
2. PR-23, the Playwright smoke flow, as planned.
3. PR-25, Stage A: the manifest, the icons, the service worker, the install hint, the phone import, and the phone gate.
4. The ban rerun item of the roadmap, with the `stale` flag and the banner.
5. PR-26, Stage B: the opt-ins, the weekly job, push through Cloud Messaging, and the email digest.
6. Stage C on request (D-548).

## 8. The questions for the owner

- OQ-66: Stage A inside PR-22, or a PR of its own after PR-22? Answered 2026-09-06 (D-547): its own PR, PR-25, after PR-23.
- OQ-67: Stage B channels: push, email, or both? And the three events, or fewer? Open.
- OQ-68: Stage C: never, on request, or planned? Answered 2026-09-06 (D-548): on request.

## 9. Sources, read on 2026-09-05

| Fact | Source |
|---|---|
| Web push for Home Screen web apps since iOS 16.4, none inside Safari on iOS, and iOS 26 opens a Home Screen site as a web app by default | https://www.magicbell.com/blog/pwa-ios-limitations-safari-support-complete-guide and https://www.idownloadblog.com/2025/06/17/apple-ios-26-safari-web-apps-home-screen-bookmarks/ |
| Apple withdrew the EU removal of Home Screen web apps on 2024-03-01 | https://techcrunch.com/2024/03/01/apple-reverses-decision-about-blocking-web-apps-on-iphones-in-the-eu/ |
| The Web Share Target API: Android and Windows, not iOS | https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps/Manifest/Reference/share_target |
| ManaBox exports the collection or one binder as CSV from the top right menu | https://www.manabox.app/guides/collection/import-export/ |
| Firebase Cloud Messaging is free, with a VAPID key and a service worker for the web | https://firebase.google.com/pricing and https://dev.to/pooyagolchian/firebase-cloud-messaging-in-2026-web-push-notifications-with-vue-3-sdk-v10-4l0l |
| Email free tiers: Brevo 300 a day, Resend 3,000 a month, Postmark 100 a month | https://www.brevo.com/blog/best-transactional-email-services/ and https://www.buildmvpfast.com/api-costs/email (third party) |
| Google Play accepts a web app as a Trusted Web Activity | https://developer.android.com/develop/ui/views/layout/webapps/trusted-web-activities |
| Apple Developer Program $99 a year, and guideline 4.2 on minimum functionality | https://developer.apple.com/app-store/review/guidelines/ |
| The product facts | `docs/design-roadmap.md` (PR-3, PR-6, PR-12B, PR-14C, PR-16, PR-20, PR-21, the ban rerun plan), D-311, D-315, D-328, D-392, D-421 |
