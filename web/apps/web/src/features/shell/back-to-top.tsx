import { ArrowUpIcon } from "lucide-react";
import { type RefObject, useEffect, useState, useSyncExternalStore } from "react";

// A phone shows the button once the page scrolls this far (D-1050).
export const backToTopOffset = 300;
export const phoneQuery = "(max-width: 640px)";

function subscribe(onChange: () => void): () => void {
  if (typeof window === "undefined" || !window.matchMedia) return () => {};
  const list = window.matchMedia(phoneQuery);
  list.addEventListener("change", onChange);
  return () => list.removeEventListener("change", onChange);
}

function onPhone(): boolean {
  if (typeof window === "undefined" || !window.matchMedia) return false;
  return window.matchMedia(phoneQuery).matches;
}

// BackToTop is the phone button under the header that scrolls the page
// back to its top. The main region of the shell scrolls, and the window
// does not (D-364), so the button reads and moves that region.
export function BackToTop({ scroller }: { scroller: RefObject<HTMLElement | null> }) {
  const phone = useSyncExternalStore(subscribe, onPhone, () => false);
  const [past, setPast] = useState(false);
  useEffect(() => {
    const el = scroller.current;
    if (!el || !phone) return;
    const read = () => setPast(el.scrollTop > backToTopOffset);
    el.addEventListener("scroll", read, { passive: true });
    return () => el.removeEventListener("scroll", read);
  }, [scroller, phone]);
  if (!phone || !past) return null;
  function onClick() {
    const reduced = typeof window.matchMedia === "function" && window.matchMedia("(prefers-reduced-motion: reduce)").matches;
    scroller.current?.scrollTo({ top: 0, behavior: reduced ? "auto" : "smooth" });
  }
  return (
    <button
      type="button"
      onClick={onClick}
      className="fixed top-[calc(var(--header-height,0px)+0.5rem)] left-1/2 z-20 flex min-h-11 -translate-x-1/2 items-center gap-1.5 rounded-full border border-border bg-card/95 px-4 text-sm font-medium print:hidden"
    >
      <ArrowUpIcon className="size-4" aria-hidden="true" />
      Back to top
    </button>
  );
}
