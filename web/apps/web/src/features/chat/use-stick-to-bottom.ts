import { type RefObject, useEffect, useRef } from "react";

// footSlackPx is how far above the foot of the chat a reader can stand
// and still count as at the bottom. The page pads the chat, so the foot
// never meets the edge of the frame exactly.
export const footSlackPx = 80;

// gapBelow is how far the foot sits under the bottom edge of its
// scroller. It is negative when the foot is in view above that edge.
export function gapBelow(foot: Element, scroller: Element): number {
  return foot.getBoundingClientRect().bottom - scroller.getBoundingClientRect().bottom;
}

// scrollsItself reports a box whose content overflows it, as the docked
// thread does on a wide screen. On a phone that box grows with its
// content, and the page scrolls in its place.
function scrollsItself(el: HTMLElement | null): el is HTMLElement {
  return el !== null && el.scrollHeight > el.clientHeight;
}

// pageOf is the scroller of the page: the main region of the shell, and
// not the window (D-364).
function pageOf(foot: HTMLElement): HTMLElement | null {
  return foot.closest("main") ?? (document.scrollingElement as HTMLElement | null);
}

// toFoot moves the thread box and the page so that the foot is in view.
// given records each scrollTop it sets.
function toFoot(b: HTMLElement | null, f: HTMLElement | null, given: WeakMap<Element, number>) {
  if (scrollsItself(b)) {
    b.scrollTop = b.scrollHeight;
    given.set(b, b.scrollTop);
  }
  // A hidden chat has no box at all, and the page keeps its place.
  if (!f || f.getClientRects().length === 0) return;
  const page = pageOf(f);
  if (!page) return;
  const gap = gapBelow(f, page);
  if (gap > 0) {
    page.scrollTop += gap;
    given.set(page, page.scrollTop);
  }
}

// topSlackPx is the space the view leaves above the reader's message when
// it puts that message at the top (D-1066).
export const topSlackPx = 8;

// toTop moves the thread box and the page so that the top of el sits at
// the top of each, with a little space above it. given records each
// scrollTop it sets.
function toTop(b: HTMLElement | null, el: HTMLElement, given: WeakMap<Element, number>) {
  if (el.getClientRects().length === 0) return;
  if (scrollsItself(b) && b.contains(el)) {
    b.scrollTop += el.getBoundingClientRect().top - b.getBoundingClientRect().top - topSlackPx;
    given.set(b, b.scrollTop);
    return;
  }
  const page = pageOf(el);
  if (!page) return;
  page.scrollTop += el.getBoundingClientRect().top - page.getBoundingClientRect().top - topSlackPx;
  given.set(page, page.scrollTop);
}

// useStickToBottom keeps the chat at its foot (D-1056). A send pins the
// view, and each change of the thread then moves the view to the foot.
// A reader who scrolls up unpins it, and a scroll back to the foot pins it
// again. The view moves the thread box and the page alone. scrollIntoView
// moves every ancestor, and that pushed the docked chat off the frame
// (D-370). change is a key that changes with each change of the thread.
//
// A turn that ends with open questions puts the reader's latest message,
// top, at the top of the view, and the questions read down from it. The
// view then stays there until a send (D-1066). holdTop says the turn
// asks.
export function useStickToBottom(box: RefObject<HTMLElement | null>, foot: RefObject<HTMLElement | null>, change: string, top?: RefObject<HTMLElement | null>, holdTop = false) {
  const pinned = useRef(true);
  // set holds the scrollTop that this hook last gave each scroller. A
  // scroll event that lands at or under it comes from this hook or from
  // new content, and only a scroll above it is the reader's own.
  const set = useRef(new WeakMap<Element, number>());

  // The layout can swap the thread box and the foot for new ones, so
  // each change attaches the listeners again.
  useEffect(() => {
    const f = foot.current;
    const b = box.current;
    const page = f ? pageOf(f) : null;
    function onScroll(e: Event) {
      const el = e.currentTarget as HTMLElement;
      // The page under a docked thread does not decide where the chat is.
      if (el !== b && scrollsItself(b)) return;
      const atFoot = el === b ? b.scrollHeight - b.scrollTop - b.clientHeight <= footSlackPx : f !== null && gapBelow(f, el) <= footSlackPx;
      if (atFoot) {
        pinned.current = true;
        return;
      }
      const given = set.current.get(el);
      if (given === undefined || el.scrollTop < given - 1) pinned.current = false;
    }
    const targets = [b, page].filter((el): el is HTMLElement => el !== null);
    for (const el of targets) el.addEventListener("scroll", onScroll, { passive: true });
    const mine = top?.current;
    if (pinned.current && holdTop && mine) {
      toTop(b, mine, set.current);
      pinned.current = false;
    } else if (pinned.current) {
      toFoot(b, f, set.current);
    }
    return () => {
      for (const el of targets) el.removeEventListener("scroll", onScroll);
    };
  }, [box, foot, change, top, holdTop]);

  // pin sticks the view to the foot again. A send calls it, so the reader
  // sees the new turn and the button that answers it.
  return () => {
    pinned.current = true;
    toFoot(box.current, foot.current, set.current);
  };
}
