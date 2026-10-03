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

// openProofLink takes the code of a proof link. The API proves the
// email, and the custom token of its answer signs in this browser
// (D-1082). It answers false when the answer carries no token: the email
// is proved, and the person signs in with the password.
export async function openProofLink(code: string): Promise<boolean> {
  const res = await proofClient.openLink({ code });
  if (!res.customToken) return false;
  await signInWithToken(res.customToken);
  return true;
}

// standaloneApp is true inside an installed app, where the link of an
// email never opens on iOS (D-1083).
export function standaloneApp(): boolean {
  const nav = navigator as Navigator & { standalone?: boolean };
  return nav.standalone === true || (typeof window.matchMedia === "function" && window.matchMedia("(display-mode: standalone)").matches);
}

// iosBrowser is true in a browser of an iPhone or an iPad. iOS opens the
// link of an email in the browser, and the installed app keeps its own
// sign-in (D-1083). An iPad can name itself a Mac, so the touch points
// tell the two apart.
export function iosBrowser(): boolean {
  const ua = navigator.userAgent;
  const ipad = /Macintosh/.test(ua) && navigator.maxTouchPoints > 1;
  return (/iPhone|iPad|iPod/.test(ua) || ipad) && !standaloneApp();
}
