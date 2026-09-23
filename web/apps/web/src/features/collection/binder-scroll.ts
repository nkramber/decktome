// The binder scrolls with the page. `<main>` of the layout is the one
// scroll area (D-364, D-806), so the grid, the reset after a new choice,
// and the button to the binder top all read it.

// offsetIn is the distance from the top of the scroll content of `main`
// to the top of `el`. The binder sits below the page header, the upload
// cards, and the hero, and this is where it starts.
export function offsetIn(el: HTMLElement, main: HTMLElement): number {
  return el.getBoundingClientRect().top - main.getBoundingClientRect().top + main.scrollTop;
}

// resetTop is where a new filter or sort moves the page. A reader below
// the binder top returns to it, so they read the new list from its first
// row. A reader above the binder stays where they are.
export function resetTop(scrollTop: number, binderTop: number): number | undefined {
  return scrollTop > binderTop ? binderTop : undefined;
}

// showTopButton says when the button to the binder top shows: after the
// reader scrolls one screen below the binder top (D-808).
export function showTopButton(scrollTop: number, binderTop: number, screen: number): boolean {
  return scrollTop > binderTop + screen;
}

// moveBehavior honors the reader's motion setting. A reader who asks for
// less motion gets a jump, and every other reader gets a smooth scroll.
export function moveBehavior(): ScrollBehavior {
  const reduced = typeof window.matchMedia === "function" && window.matchMedia("(prefers-reduced-motion: reduce)").matches;
  return reduced ? "auto" : "smooth";
}
