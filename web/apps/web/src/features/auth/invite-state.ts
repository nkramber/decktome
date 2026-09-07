import { useSyncExternalStore } from "react";

// The invite state of the signed-in reader (F-59). The API is the only
// authority on the list (D-314, D-420), so this holds what it answered.
// "unknown" is the state on a page that asks nothing, "checking" is the
// state while the gate waits for the answer, and "unavailable" is the
// state when the API gave none. The app draws for "invited" alone
// (D-590).
export type InviteState = "unknown" | "checking" | "invited" | "refused" | "unavailable";

// The state lives outside React, because the header and the route guard
// read it and neither owns it. It stays in this module and not in the
// persisted store: an invite is the API's word, not the browser's.
let state: InviteState = "unknown";
const listeners = new Set<() => void>();

export function setInviteState(next: InviteState) {
  if (next === state) return;
  state = next;
  for (const listener of listeners) listener();
}

// Sign-out forgets the answer, because the next reader is another person.
export function resetInviteState() {
  setInviteState("unknown");
}

function subscribe(listener: () => void) {
  listeners.add(listener);
  return () => listeners.delete(listener);
}

export function useInviteState(): InviteState {
  return useSyncExternalStore(
    subscribe,
    () => state,
    () => "unknown" as InviteState,
  );
}
