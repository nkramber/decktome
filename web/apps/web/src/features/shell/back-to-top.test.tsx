import { act, fireEvent, render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useRef } from "react";
import { afterEach, describe, expect, it, vi } from "vitest";

import { BackToTop } from "./back-to-top";

function media(phone: boolean) {
  vi.stubGlobal(
    "matchMedia",
    vi.fn((query: string) => ({ matches: phone && query === "(max-width: 640px)", media: query, addEventListener: vi.fn(), removeEventListener: vi.fn() })),
  );
}

function Page() {
  const ref = useRef<HTMLDivElement>(null);
  return (
    <>
      <div ref={ref} data-testid="scroller" />
      <BackToTop scroller={ref} />
    </>
  );
}

function scrollTo(el: HTMLElement, top: number) {
  act(() => {
    el.scrollTop = top;
    fireEvent.scroll(el);
  });
}

afterEach(() => vi.unstubAllGlobals());

describe("BackToTop (D-1050)", () => {
  it("shows on a phone past 300 pixels, hides above them, and scrolls the region to its top", async () => {
    media(true);
    render(<Page />);
    const el = screen.getByTestId("scroller");
    const move = vi.fn();
    el.scrollTo = move as unknown as typeof el.scrollTo;
    expect(screen.queryByRole("button", { name: "Back to top" })).not.toBeInTheDocument();
    scrollTo(el, 300);
    expect(screen.queryByRole("button", { name: "Back to top" })).not.toBeInTheDocument();
    scrollTo(el, 301);
    await userEvent.setup().click(screen.getByRole("button", { name: "Back to top" }));
    expect(move).toHaveBeenCalledWith({ top: 0, behavior: "smooth" });
    scrollTo(el, 0);
    expect(screen.queryByRole("button", { name: "Back to top" })).not.toBeInTheDocument();
  });

  it("never shows on a wide screen", () => {
    media(false);
    render(<Page />);
    scrollTo(screen.getByTestId("scroller"), 2000);
    expect(screen.queryByRole("button", { name: "Back to top" })).not.toBeInTheDocument();
  });
});
