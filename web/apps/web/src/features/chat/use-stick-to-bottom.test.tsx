import { fireEvent, render, screen } from "@testing-library/react";
import { useRef } from "react";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { backToTopClearPx } from "../shell/back-to-top";
import { footSlackPx, topSlackPx, useStickToBottom } from "./use-stick-to-bottom";

// jsdom lays nothing out, so each test gives the page, the thread box,
// and the foot their geometry by test id.
type Geometry = { scrollTop: number; scrollHeight: number; clientHeight: number; bottom: number; shown: boolean; top?: number };
let geo: Record<string, Geometry>;
const of = (el: Element) => geo[(el as HTMLElement).dataset.testid ?? ""];

function Harness({ change, asks = false }: { change: string; asks?: boolean }) {
  const box = useRef<HTMLDivElement>(null);
  const foot = useRef<HTMLDivElement>(null);
  const first = useRef<HTMLDivElement>(null);
  const pin = useStickToBottom(box, foot, change, first, asks);
  return (
    <main data-testid="page">
      <div ref={box} data-testid="box" />
      <div ref={first} data-testid="first" />
      <div ref={foot} data-testid="foot" />
      <button type="button" onClick={pin}>
        Send
      </button>
    </main>
  );
}

beforeEach(() => {
  geo = {
    page: { scrollTop: 0, scrollHeight: 3000, clientHeight: 800, bottom: 800, shown: true, top: 0 },
    first: { scrollTop: 0, scrollHeight: 0, clientHeight: 0, bottom: 1100, shown: true, top: 1000 },
    box: { scrollTop: 0, scrollHeight: 0, clientHeight: 0, bottom: 0, shown: true },
    foot: { scrollTop: 0, scrollHeight: 0, clientHeight: 0, bottom: 1500, shown: true },
  };
  vi.spyOn(HTMLElement.prototype, "scrollTop", "get").mockImplementation(function (this: HTMLElement) {
    return of(this)?.scrollTop ?? 0;
  });
  vi.spyOn(HTMLElement.prototype, "scrollTop", "set").mockImplementation(function (this: HTMLElement, v: number) {
    const g = of(this);
    if (!g) return;
    // A real scroller stops at its end, and the foot moves up with it.
    const top = Math.max(0, Math.min(v, g.scrollHeight - g.clientHeight));
    if (this.dataset.testid === "page") {
      geo.foot.bottom -= top - g.scrollTop;
      geo.first.top = (geo.first.top ?? 0) - (top - g.scrollTop);
    }
    g.scrollTop = top;
  });
  vi.spyOn(HTMLElement.prototype, "scrollHeight", "get").mockImplementation(function (this: HTMLElement) {
    return of(this)?.scrollHeight ?? 0;
  });
  vi.spyOn(HTMLElement.prototype, "clientHeight", "get").mockImplementation(function (this: HTMLElement) {
    return of(this)?.clientHeight ?? 0;
  });
  vi.spyOn(Element.prototype, "getBoundingClientRect").mockImplementation(function (this: Element) {
    return { bottom: of(this)?.bottom ?? 0, top: of(this)?.top ?? 0 } as DOMRect;
  });
  vi.spyOn(Element.prototype, "getClientRects").mockImplementation(function (this: Element) {
    return { length: of(this)?.shown === false ? 0 : 1 } as DOMRectList;
  });
});

afterEach(() => {
  vi.restoreAllMocks();
});

describe("useStickToBottom (D-1056)", () => {
  it("moves the page to the foot of the chat, and follows each change while the reader is there", () => {
    const view = render(<Harness change="1" />);
    expect(geo.page.scrollTop).toBe(700);
    geo.foot.bottom += 300;
    view.rerender(<Harness change="2" />);
    expect(geo.page.scrollTop).toBe(1000);
  });

  it("stays where the reader scrolled up to, and a send takes it to the foot again", () => {
    const view = render(<Harness change="1" />);
    const page = screen.getByTestId("page");
    page.scrollTop = 200;
    fireEvent.scroll(page);
    geo.foot.bottom += 300;
    view.rerender(<Harness change="2" />);
    expect(geo.page.scrollTop).toBe(200);
    fireEvent.click(screen.getByRole("button", { name: "Send" }));
    expect(geo.page.scrollTop).toBe(1000);
  });

  it("follows again after the reader scrolls back to the foot", () => {
    const view = render(<Harness change="1" />);
    const page = screen.getByTestId("page");
    page.scrollTop = 200;
    fireEvent.scroll(page);
    page.scrollTop = 700 - footSlackPx / 2;
    fireEvent.scroll(page);
    geo.foot.bottom += 300;
    view.rerender(<Harness change="2" />);
    expect(geo.page.scrollTop).toBe(1000);
  });

  it("keeps following when new content lands under a scroll it made", () => {
    const view = render(<Harness change="1" />);
    // The text grows before the scroll event of the hook arrives.
    geo.foot.bottom += 400;
    fireEvent.scroll(screen.getByTestId("page"));
    view.rerender(<Harness change="2" />);
    expect(geo.page.scrollTop).toBe(1100);
  });

  it("scrolls a docked thread box to its end, and leaves the page alone", () => {
    geo.box = { scrollTop: 0, scrollHeight: 2000, clientHeight: 500, bottom: 600, shown: true };
    geo.foot.bottom = 700;
    render(<Harness change="1" />);
    expect(geo.box.scrollTop).toBe(1500);
    expect(geo.page.scrollTop).toBe(0);
  });

  // D-1066, D-1071: a turn that asks shows the first question at the top,
  // and the view stays there.
  it("puts the first question at the top when the turn ends with questions", () => {
    const view = render(<Harness change="1" />);
    expect(geo.page.scrollTop).toBe(700);
    geo.foot.bottom += 600;
    view.rerender(<Harness change="2" asks />);
    expect(geo.page.scrollTop).toBe(1000 - topSlackPx);
    expect(geo.first.top).toBe(topSlackPx);
    // A later change of the thread does not drag the view to the foot.
    geo.foot.bottom += 100;
    view.rerender(<Harness change="3" asks />);
    expect(geo.page.scrollTop).toBe(1000 - topSlackPx);
    // The send of the answers takes the view to the foot again.
    fireEvent.click(screen.getByRole("button", { name: "Send" }));
    expect(geo.page.scrollTop).toBe(1400);
  });

  // D-1072: on a phone the "Back to top" button sits above the question
  // text, not on it.
  it("leaves the band of the Back to top button above the first question on a phone", () => {
    vi.stubGlobal("matchMedia", (query: string) => ({ matches: query.includes("max-width"), addEventListener() {}, removeEventListener() {} }));
    try {
      const view = render(<Harness change="1" />);
      geo.foot.bottom += 600;
      view.rerender(<Harness change="2" asks />);
      expect(geo.first.top).toBe(topSlackPx + backToTopClearPx);
      expect(geo.page.scrollTop).toBe(1000 - topSlackPx - backToTopClearPx);
    } finally {
      vi.unstubAllGlobals();
    }
  });

  it("takes a docked thread box to its end, and the first question to the top", () => {
    geo.box = { scrollTop: 0, scrollHeight: 2000, clientHeight: 500, bottom: 600, shown: true };
    geo.foot.bottom = 700;
    const view = render(<Harness change="1" />);
    expect(geo.box.scrollTop).toBe(1500);
    geo.box.scrollHeight = 2600;
    geo.first.top = 1000;
    view.rerender(<Harness change="2" asks />);
    expect(geo.box.scrollTop).toBe(2100);
    expect(geo.page.scrollTop).toBe(1000 - topSlackPx);
    expect(geo.first.top).toBe(topSlackPx);
  });

  it("goes to the foot as before when the turn asks nothing", () => {
    const view = render(<Harness change="1" />);
    geo.foot.bottom += 600;
    view.rerender(<Harness change="2" />);
    expect(geo.page.scrollTop).toBe(1300);
  });

  it("does not move the page for a chat that is hidden", () => {
    geo.foot.shown = false;
    render(<Harness change="1" />);
    expect(geo.page.scrollTop).toBe(0);
  });
});
