import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

// Client state (wallabee tier 2). Server state lives in TanStack Query and
// the auth user in React context. localStorage keeps the session id, so a
// reload returns to the chat.
export type PoolMode = "owned" | "any";

export type AppState = {
  // The active collection, or empty in any-card mode (D-37). It lives
  // for one page load: a new load builds from any card until the reader
  // names a collection (D-345).
  collectionId: string;
  // The current chat session, or empty before the first message.
  sessionId: string;
  // owned: only cards in the collection. any: build from any card (D-37).
  poolMode: PoolMode;
  setCollection: (collectionId: string) => void;
  clearCollection: () => void;
  setSessionId: (sessionId: string) => void;
  setPoolMode: (poolMode: PoolMode) => void;
  // reset forgets every id. Sign-out calls it.
  reset: () => void;
};

export const useAppStore = create<AppState>()(
  persist(
    (set) => ({
      collectionId: "",
      sessionId: "",
      poolMode: "any",
      setCollection: (collectionId) => set({ collectionId, poolMode: "owned" }),
      clearCollection: () => set({ collectionId: "", poolMode: "any" }),
      setSessionId: (sessionId) => set({ sessionId }),
      setPoolMode: (poolMode) => set({ poolMode }),
      reset: () => set({ collectionId: "", sessionId: "", poolMode: "any" }),
    }),
    {
      name: "mtg-deck-builder",
      storage: createJSONStorage(() => localStorage),
      // Only the session id survives a reload. The pool is a choice of
      // one chat, not a setting of the app, so every load starts at any
      // card (D-345). A browser that stored the old shape drops it.
      version: 1,
      migrate: (persisted) => ({ sessionId: (persisted as { sessionId?: string } | null)?.sessionId ?? "" }) as Partial<AppState>,
      partialize: (s) => ({ sessionId: s.sessionId }),
    },
  ),
);
