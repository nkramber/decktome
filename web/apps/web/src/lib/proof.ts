import { Code, ConnectError } from "@connectrpc/connect";

import { proofClient } from "./api";
import { sendFirebaseProof, signInWithToken } from "./firebase";

// sendProof sends the email that proves the address of the signed-in
// account (D-1081). The API sends it through Resend, with the text of
// the owner and a short link. When the API sends no email, the email of
// Firebase goes out in its place, so a new account always gets a link:
// a server with no Resend secret, a send that failed, or a server with no
// route. A send within the cooldown of the API is the one refusal that
// stays, because a second email would arrive at once.
export async function sendProof(): Promise<void> {
  try {
    await proofClient.sendLink({});
  } catch (err) {
    if (ConnectError.from(err).code === Code.ResourceExhausted) throw err;
    await sendFirebaseProof();
  }
}

// ProofOutcome is the result of one open of a proof link. "signed-in"
// means the custom token signed in this browser. "proved" means the
// email is proved and the answer carries no token. "already-proved"
// means the link worked before and the email is proved (D-1119).
export type ProofOutcome = "signed-in" | "proved" | "already-proved";

// openProofLink takes the code of a proof link. The API proves the
// email, and the custom token of its answer signs in this browser
// (D-1082). With no token, the person signs in with the password.
export async function openProofLink(code: string): Promise<ProofOutcome> {
  const res = await proofClient.openLink({ code });
  if (res.alreadyProved) return "already-proved";
  if (!res.customToken) return "proved";
  await signInWithToken(res.customToken);
  return "signed-in";
}
