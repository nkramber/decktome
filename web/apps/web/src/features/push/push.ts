import { pushClient } from "../../lib/api";
import { loadAuth } from "../../lib/firebase";

// The web push of PR-26 (D-1004, D-1005). A build that ends after the
// user left the page reaches each device that took push. A device is one
// browser install, named by its Firebase Installation ID, and the API
// keeps it under the user. The push handler sits in the service worker
// of PR-25 (`public/push-handler.js`).
//
// The SDK loads on the first call alone, so it stays off the first paint
// (D-320).

export type PushState = "unavailable" | "blocked" | "off" | "on";

// The flag says that this browser turned push on. It saves a network
// read for each browser that never did. The API is the source of truth.
export const pushFlagKey = "decktome.push";

// The release marker says that the end of a registration with Cloud
// Messaging failed, for example offline. The next start tries again.
export const pushReleaseKey = "decktome.push.release";

// A build with no service worker never resolves `ready`, so the wait
// ends here.
const workerWaitMs = 10_000;

export function pushConfigured(): boolean {
  return Boolean(import.meta.env.VITE_FIREBASE_MESSAGING_SENDER_ID);
}

function readFlag(): boolean {
  try {
    return localStorage.getItem(pushFlagKey) === "1";
  } catch {
    return false;
  }
}

function writeKey(key: string, on: boolean) {
  try {
    if (on) localStorage.setItem(key, "1");
    else localStorage.removeItem(key);
  } catch {
    // A browser with no storage asks the API on each menu open.
  }
}

function writeFlag(on: boolean) {
  writeKey(pushFlagKey, on);
}

function releasePending(): boolean {
  try {
    return localStorage.getItem(pushReleaseKey) === "1";
  } catch {
    return false;
  }
}

// endRegistration ends the registration with Cloud Messaging. After it,
// no push reaches this browser, whatever the API still holds, and the
// next send removes the device as gone.
async function endRegistration(): Promise<void> {
  try {
    const { mod, messaging: m } = await messaging();
    await mod.unregister(m);
    writeKey(pushReleaseKey, false);
  } catch {
    writeKey(pushReleaseKey, true);
  }
}

async function messaging() {
  const [{ app }, mod] = await Promise.all([loadAuth(), import("firebase/messaging")]);
  return { mod, messaging: mod.getMessaging(app) };
}

async function installationId(): Promise<string> {
  const [{ app }, mod] = await Promise.all([loadAuth(), import("firebase/installations")]);
  return mod.getId(mod.getInstallations(app));
}

async function browserSupports(): Promise<boolean> {
  if (!pushConfigured() || typeof navigator === "undefined" || !("serviceWorker" in navigator) || typeof Notification === "undefined") {
    return false;
  }
  const { isSupported } = await import("firebase/messaging");
  return isSupported();
}

function workerReady(): Promise<ServiceWorkerRegistration> {
  return Promise.race([
    navigator.serviceWorker.ready,
    new Promise<never>((_, reject) => setTimeout(() => reject(new Error("no service worker")), workerWaitMs)),
  ]);
}

// pushState reads the state of this browser. An iPhone shows push only
// in the app on the Home Screen, so Safari in a tab reads unavailable.
export async function pushState(): Promise<PushState> {
  if (!(await browserSupports())) return "unavailable";
  if (Notification.permission === "denied") return "blocked";
  if (!readFlag() || Notification.permission !== "granted") return "off";
  const res = await pushClient.getDevice({ installationId: await installationId() });
  if (!res.registered) {
    writeFlag(false);
    return "off";
  }
  return "on";
}

// enablePush asks for the permission when the browser has no answer yet,
// registers with Cloud Messaging, and stores the device. Call it from a
// click. Safari asks for the permission only inside the gesture, so the
// ask comes before the first await.
export async function enablePush(): Promise<PushState> {
  const asked = Notification.permission === "default" ? Notification.requestPermission() : Promise.resolve(Notification.permission);
  const permission = await asked;
  if (permission === "denied") return "blocked";
  if (permission !== "granted") return "off";
  const { mod, messaging: m } = await messaging();
  let stop = () => {};
  const registered = new Promise<string>((resolve) => {
    stop = mod.onRegistered(m, resolve);
  });
  try {
    await mod.register(m, { serviceWorkerRegistration: await workerReady() });
  } catch (err) {
    stop();
    if (Notification.permission === "denied") return "blocked";
    if (Notification.permission !== "granted") return "off";
    throw err;
  }
  const id = await registered;
  stop();
  await pushClient.registerDevice({ installationId: id });
  writeFlag(true);
  return "on";
}

// disablePush removes the device from the API first, then ends the
// registration with Cloud Messaging. The registration ends also when the
// API call fails, so no push can reach this browser after it.
export async function disablePush(): Promise<PushState> {
  try {
    await pushClient.unregisterDevice({ installationId: await installationId() });
  } finally {
    writeFlag(false);
    await endRegistration();
  }
  return "off";
}

// releasePushOnSignOut runs before a sign-out. The next account on this
// browser must never see a push with the deck of the last one. So the
// registration with Cloud Messaging ends when the API call fails, and the
// API gives the ID to the next account that turns push on. A failure
// never stops the sign-out.
//
// The flag saves a read, but a browser can refuse to store it. So with no
// flag, a browser that granted the permission asks the API, which is the
// source of truth. A read that fails releases the device all the same.
export async function releasePushOnSignOut(): Promise<void> {
  if (!readFlag()) {
    if (!pushConfigured() || typeof Notification === "undefined" || Notification.permission !== "granted") return;
    try {
      const res = await pushClient.getDevice({ installationId: await installationId() });
      if (!res.registered) return;
    } catch {
      // The state is not known, so the release goes on.
    }
  }
  try {
    await disablePush();
  } catch {
    // disablePush ended the registration, or it left the release marker.
  }
}

// refreshPush registers the device again on each start of a browser that
// turned push on. The SDK keeps a registration fresh this way, and a
// device the API removed as gone comes back with its new registration.
export async function refreshPush(): Promise<void> {
  if (releasePending()) await endRegistration();
  if (!readFlag()) return;
  try {
    if (!(await browserSupports()) || Notification.permission !== "granted") {
      writeFlag(false);
      return;
    }
    await enablePush();
  } catch {
    // The next start tries again.
  }
}
