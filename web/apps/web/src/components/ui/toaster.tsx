import { useSyncExternalStore } from "react";
import { Toaster as Sonner } from "sonner";

// The width at which sonner lays a toast out for a phone. Its own
// stylesheet uses the same query, so the position and the mobile offsets
// change at one width.
export const phoneToastQuery = "(max-width: 600px)";

function subscribe(onChange: () => void): () => void {
  if (typeof window === "undefined" || !window.matchMedia) return () => {};
  const list = window.matchMedia(phoneToastQuery);
  list.addEventListener("change", onChange);
  return () => list.removeEventListener("change", onChange);
}

function onPhone(): boolean {
  if (typeof window === "undefined" || !window.matchMedia) return false;
  return window.matchMedia(phoneToastQuery).matches;
}

// One toast host for the app. Every mutation reports its result here
// (D-311). Dark is the only theme (D-330).
//
// On a phone a toast at the bottom covered the last button of a page, so
// there it shows under the header (F-194, D-1030). The layout writes the
// height of the header to --header-height. The desktop keeps the bottom
// right, and the installed app adds its bottom buffer there (F-187, D-1010).
export function Toaster() {
  const phone = useSyncExternalStore(subscribe, onPhone, () => false);
  return (
    <Sonner
      theme="dark"
      position={phone ? "top-center" : "bottom-right"}
      offset={{ bottom: "calc(24px + var(--edge-bottom))" }}
      mobileOffset={{ top: "calc(var(--header-height, 0px) + 8px)" }}
      toastOptions={{
        classNames: {
          toast: "rounded-card border border-border bg-card text-card-foreground",
          description: "text-muted-foreground",
        },
      }}
    />
  );
}

export { toast } from "sonner";
