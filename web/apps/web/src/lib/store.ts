import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

// Client state (wallabee tier 2). Server state lives in TanStack Query and
// the auth user in React context. localStorage keeps the session id, so a
// reload returns to the chat.
// The card pool of the next chat (D-359).
// owned_only: only cards in the collection.
// owned_first: the collection leads, and the database fills a gap.
// any: the whole database, with no collection.
export type PoolMode = "owned_only" | "owned_first" | "any";

export type AppState = {
  // The active collection, or empty in any-card mode (D-37). It lives
  // for one page load: a new load builds from any card until the reader
  // names a collection (D-345).
  collectionId: string;
  // The current chat session, or empty before the first message.
  sessionId: string;
  poolMode: PoolMode;
  // hadDecks says the reader owned at least one deck the last time the
  // app read the list. The new chat shows a placeholder for the deck row
  // only when it is true, so a reader with no deck never sees a box
  // appear and go (F-67).
  hadDecks: boolean;
  // installHintDismissed says the reader closed the install hint on this
  // device. The hint shows once, and never again after that (PR-25).
  installHintDismissed: boolean;
  setCollection: (collectionId: string) => void;
  clearCollection: () => void;
  setSessionId: (sessionId: string) => void;
  setPoolMode: (poolMode: PoolMode) => void;
  setHadDecks: (hadDecks: boolean) => void;
  dismissInstallHint: () => void;
  // reset forgets every id. Sign-out calls it.
  reset: () => void;
};

export const useAppStore = create<AppState>()(
  persist(
    (set) => ({
      collectionId: "",
      sessionId: "",
      poolMode: "any",
      hadDecks: false,
      installHintDismissed: false,
      // A named collection leads by default, and the database fills a
      // gap. A reader who wants no fill checks "Only cards I own".
      setCollection: (collectionId) => set({ collectionId, poolMode: "owned_first" }),
      clearCollection: () => set({ collectionId: "", poolMode: "any" }),
      setSessionId: (sessionId) => set({ sessionId }),
      setPoolMode: (poolMode) => set({ poolMode }),
      setHadDecks: (hadDecks) => set({ hadDecks }),
      dismissInstallHint: () => set({ installHintDismissed: true }),
      // Sign-out forgets the deck count with the ids: the next reader is
      // another person.
      // The install hint is not reset here. It answers a question about
      // this device's Home Screen, and the next reader of the device
      // already read it (PR-25).
      reset: () => set({ collectionId: "", sessionId: "", poolMode: "any", hadDecks: false }),
    }),
    {
      name: "mtg-deck-builder",
      storage: createJSONStorage(() => localStorage),
      // The session id and the deck mark survive a reload. The pool is a
      // choice of one chat, not a setting of the app, so every load
      // starts at any card (D-345). A browser that stored the old shape
      // drops it. A store with no hadDecks reads the default, false, so
      // the mark needs no new version.
      version: 1,
      migrate: (persisted) => ({ sessionId: (persisted as { sessionId?: string } | null)?.sessionId ?? "" }) as Partial<AppState>,
      partialize: (s) => ({ sessionId: s.sessionId, hadDecks: s.hadDecks, installHintDismissed: s.installHintDismissed }),
    },
  ),
);
