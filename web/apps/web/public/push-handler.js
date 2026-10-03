// The push handler of PR-26 (D-1004, D-1005). The service worker of
// PR-25 loads this file through workbox importScripts, so the app keeps
// one worker. An iPhone takes push in the app on the Home Screen alone.
//
// The API sends a data message through Cloud Messaging: a title, a
// body, and a path of the app. Each push shows a notification, because
// Safari ends the subscription of a worker that shows none.

function pushPayload(event) {
  let raw = {};
  try {
    raw = event.data ? event.data.json() : {};
  } catch {
    raw = {};
  }
  const data = raw.data || raw.notification || raw;
  const url = typeof data.url === "string" && data.url.startsWith("/") && !data.url.startsWith("//") ? data.url : "/decks";
  return {
    title: typeof data.title === "string" && data.title ? data.title : "decktome",
    body: typeof data.body === "string" ? data.body : "",
    url,
  };
}

self.addEventListener("push", (event) => {
  const p = pushPayload(event);
  event.waitUntil(
    self.registration.showNotification(p.title, {
      body: p.body,
      icon: "/icon-192.png",
      badge: "/icon-192.png",
      tag: p.url,
      data: { url: p.url },
    }),
  );
});

// A tap opens the path of the push: a deck, or the deck list (D-1088).
// An open window of the app moves to that path, and with none open, a new
// window opens.
self.addEventListener("notificationclick", (event) => {
  event.notification.close();
  const url = new URL((event.notification.data && event.notification.data.url) || "/decks", self.location.origin).href;
  event.waitUntil(
    self.clients.matchAll({ type: "window", includeUncontrolled: true }).then((windows) => {
      const open = windows.find((w) => new URL(w.url).origin === self.location.origin);
      // navigate fails on a window the worker does not control, and a
      // new window opens then.
      if (open) return open.focus().then((w) => (w || open).navigate(url)).catch(() => self.clients.openWindow(url));
      return self.clients.openWindow(url);
    }),
  );
});
