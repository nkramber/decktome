import { signOutOfApp } from "../../lib/firebase";
import { releasePushOnSignOut } from "../push/push";
import { resetInviteState } from "./invite-state";

// clearAccountState clears the persisted ids, the query cache, and the
// invite answer, so the next account on this browser starts with nothing
// of the last one. Sign-out and a change of user both call it (REV-039).
export function clearAccountState(reset: () => void, clear: () => void) {
  reset();
  clear();
  resetInviteState();
}

// signOutAndClear ends the session, then clears the account state. The
// push device of this browser leaves the account first, while the token
// still holds (D-1005).
export async function signOutAndClear(reset: () => void, clear: () => void) {
  await releasePushOnSignOut();
  await signOutOfApp();
  clearAccountState(reset, clear);
}
