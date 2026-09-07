import { signOutOfApp } from "../../lib/firebase";
import { resetInviteState } from "./invite-state";

// signOutAndClear clears the persisted ids, the query cache, and the
// invite answer, so the next account on this browser starts with nothing
// of the last one.
export async function signOutAndClear(reset: () => void, clear: () => void) {
  await signOutOfApp();
  reset();
  clear();
  resetInviteState();
}
