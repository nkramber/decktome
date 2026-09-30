import { readFileSync } from "node:fs";
import { resolve } from "node:path";

import { beforeEach, describe, expect, it, vi } from "vitest";

// The push handler runs in the service worker, so the test runs the file
// against a fake worker scope (D-1005).
const source = readFileSync(resolve(__dirname, "../../../public/push-handler.js"), "utf8");

type Listener = (event: unknown) => void;

function load() {
  const listeners: Record<string, Listener> = {};
  const scope = {
    location: { origin: "https://decktome.com" },
    registration: { showNotification: vi.fn(() => Promise.resolve()) },
    clients: { matchAll: vi.fn(() => Promise.resolve([] as unknown[])), openWindow: vi.fn(() => Promise.resolve(null)) },
    addEventListener: (type: string, fn: Listener) => {
      listeners[type] = fn;
    },
  };
  new Function("self", source)(scope);
  return { scope, listeners };
}

function push(listeners: Record<string, Listener>, body: unknown) {
  let wait: Promise<unknown> = Promise.resolve();
  listeners.push({ data: { json: () => body }, waitUntil: (p: Promise<unknown>) => (wait = p) });
  return wait;
}

describe("push-handler.js", () => {
  let env: ReturnType<typeof load>;
  beforeEach(() => {
    env = load();
  });

  it("shows the data message of the API, and the tap opens the deck", async () => {
    await push(env.listeners, { data: { title: "Your deck is ready", body: "Karlov lifegain", url: "/decks/d1" } });
    expect(env.scope.registration.showNotification).toHaveBeenCalledWith(
      "Your deck is ready",
      expect.objectContaining({ body: "Karlov lifegain", data: { url: "/decks/d1" } }),
    );
    let wait: Promise<unknown> = Promise.resolve();
    const close = vi.fn();
    env.listeners.notificationclick({ notification: { close, data: { url: "/decks/d1" } }, waitUntil: (p: Promise<unknown>) => (wait = p) });
    await wait;
    expect(close).toHaveBeenCalled();
    expect(env.scope.clients.openWindow).toHaveBeenCalledWith("https://decktome.com/decks/d1");
  });

  it("keeps a tap on the app for a link to another site", async () => {
    for (const url of ["https://evil.example/x", "//evil.example/x", "javascript:alert(1)"]) {
      env.scope.registration.showNotification.mockClear();
      await push(env.listeners, { data: { title: "t", body: "b", url } });
      expect(env.scope.registration.showNotification).toHaveBeenCalledWith("t", expect.objectContaining({ data: { url: "/decks" } }));
    }
  });

  it("shows a notification for a push it can not read", async () => {
    let wait: Promise<unknown> = Promise.resolve();
    env.listeners.push({
      data: {
        json: () => {
          throw new Error("not json");
        },
      },
      waitUntil: (p: Promise<unknown>) => (wait = p),
    });
    await wait;
    expect(env.scope.registration.showNotification).toHaveBeenCalledWith("decktome", expect.objectContaining({ data: { url: "/decks" } }));
  });

  it("moves an open window of the app to the deck", async () => {
    const navigate = vi.fn(() => Promise.resolve(null));
    const win = { url: "https://decktome.com/collection", focus: vi.fn(() => Promise.resolve(null)), navigate };
    env.scope.clients.matchAll.mockResolvedValue([win]);
    let wait: Promise<unknown> = Promise.resolve();
    env.listeners.notificationclick({ notification: { close: vi.fn(), data: { url: "/decks/d2" } }, waitUntil: (p: Promise<unknown>) => (wait = p) });
    await wait;
    expect(win.focus).toHaveBeenCalled();
    expect(navigate).toHaveBeenCalledWith("https://decktome.com/decks/d2");
    expect(env.scope.clients.openWindow).not.toHaveBeenCalled();
  });
});
