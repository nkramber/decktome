import { deferred } from "./deferred";

// Every chunk the shell defers, in one place. The router and the layout
// render these, and the idle warm-up below brings them in before the
// first click asks for one.
export const buildMenuChunk = deferred(async () => ({ default: (await import("./components/shell-menus")).BuildMenuContent }));
export const accountMenuChunk = deferred(async () => ({ default: (await import("./components/shell-menus")).AccountMenuContent }));
export const healthFooterChunk = deferred(async () => ({ default: (await import("./components/health-footer")).HealthFooter }));
export const toasterChunk = deferred(async () => ({ default: (await import("../components/ui/toaster")).Toaster }));

export const signInChunk = deferred(async () => ({ default: (await import("../features/auth/sign-in-page")).SignInPage }));
export const collectionChunk = deferred(async () => ({ default: (await import("../features/collection/collection-page")).CollectionPage }));
export const sessionChunk = deferred(async () => ({ default: (await import("../features/chat/session-page")).SessionPage }));
export const decksChunk = deferred(async () => ({ default: (await import("../features/deck/decks-page")).DecksPage }));
export const deckScreenChunk = deferred(async () => ({ default: (await import("../features/workspace/deck-screen")).DeckScreen }));

const all = [buildMenuChunk, accountMenuChunk, healthFooterChunk, toasterChunk, signInChunk, collectionChunk, sessionChunk, decksChunk, deckScreenChunk];

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
