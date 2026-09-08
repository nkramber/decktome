import type { ReactNode } from "react";

import { deferred } from "./deferred";

// Every chunk the shell defers, in one place. The router and the layout
// render these, and the idle warm-up below brings them in before the
// first click asks for one.
export const accountMenuChunk = deferred(async () => ({ default: (await import("./components/shell-menus")).AccountMenuContent }));
export const toasterChunk = deferred(async () => ({ default: (await import("../components/ui/toaster")).Toaster }));

export const signInChunk = deferred(async () => ({ default: (await import("../features/auth/sign-in-page")).SignInPage }));
// The invite gate asks the API whether the reader is on the list (F-59).
// It holds the Connect client, so the route guard stays out of the shell
// bundle and a redirect still needs no download (D-320).
export const inviteGateChunk = deferred<{ children: ReactNode }>(async () => ({ default: (await import("../features/auth/invite-gate")).default }));
export const collectionChunk = deferred(async () => ({ default: (await import("../features/collection/collection-page")).CollectionPage }));
export const sessionChunk = deferred(async () => ({ default: (await import("../features/chat/session-page")).SessionPage }));
export const decksChunk = deferred(async () => ({ default: (await import("../features/deck/decks-page")).DecksPage }));
export const deckScreenChunk = deferred(async () => ({ default: (await import("../features/workspace/deck-screen")).DeckScreen }));
// The public deck page of a share link (D-315). It pulls in no auth
// module, so a visitor pays for none.
export const sharedDeckChunk = deferred(async () => ({ default: (await import("../features/share/shared-deck-page")).SharedDeckPage }));

const all = [accountMenuChunk, toasterChunk, signInChunk, inviteGateChunk, collectionChunk, sessionChunk, decksChunk, deckScreenChunk, sharedDeckChunk];

// warmChunks brings in every deferred chunk. A menu that mounts on the
// click costs about 320 ms of that click, measured on 2026-08-30, and a
// menu that is already mounted opens in under 20 ms.
export function warmChunks() {
  for (const chunk of all) chunk.preload();
}

// scheduleWarm waits for idle time, and it returns the way to cancel.
export function scheduleWarm(): () => void {
  if (typeof window.requestIdleCallback === "function") {
    const id = window.requestIdleCallback(warmChunks, { timeout: 2000 });
    return () => window.cancelIdleCallback?.(id);
  }
  const timer = window.setTimeout(warmChunks, 300);
  return () => window.clearTimeout(timer);
}
