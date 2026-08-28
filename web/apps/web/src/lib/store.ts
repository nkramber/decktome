import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

// Client state (wallabee tier 2). Server state lives in TanStack Query and
// the auth user in React context. localStorage keeps the session id, so a
// reload returns to the chat.
export type PoolMode = "owned" | "any";

export type AppState = {
  // The active collection, or empty in any-card mode (D-37).
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
      partialize: (s) => ({ collectionId: s.collectionId, sessionId: s.sessionId, poolMode: s.poolMode }),
    },
  ),
);
