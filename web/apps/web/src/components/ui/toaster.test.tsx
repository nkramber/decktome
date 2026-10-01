import { act, render } from "@testing-library/react";
import { afterEach, describe, expect, it, vi } from "vitest";

import { phoneToastQuery, toast, Toaster } from "./toaster";

function stubWidth(phone: boolean) {
  vi.stubGlobal("matchMedia", (q: string) => ({
    matches: q === phoneToastQuery ? phone : false,
    media: q,
    addEventListener: vi.fn(),
    removeEventListener: vi.fn(),
  }));
}

async function shownToaster(): Promise<HTMLElement> {
  render(<Toaster />);
  act(() => {
    toast("Chat deleted");
  });
  const host = await vi.waitFor(() => {
    const el = document.querySelector<HTMLElement>("[data-sonner-toaster]");
    if (!el) throw new Error("no toaster");
    return el;
  });
  return host;
}

afterEach(() => {
  vi.unstubAllGlobals();
  toast.dismiss();
});

// F-194 and D-1030: a toast at the bottom of a phone covered the last
// button of a page, so on a phone the toast shows under the header.
describe("Toaster", () => {
  it("shows a toast at the top, under the header, on a phone", async () => {
    stubWidth(true);
    const host = await shownToaster();
    expect(host.dataset.yPosition).toBe("top");
    expect(host.style.getPropertyValue("--mobile-offset-top")).toBe("calc(var(--header-height, 0px) + 8px)");
  });

  it("keeps a toast at the bottom right on a wider screen", async () => {
    stubWidth(false);
    const host = await shownToaster();
    expect(host.dataset.yPosition).toBe("bottom");
    expect(host.dataset.xPosition).toBe("right");
    expect(host.style.getPropertyValue("--offset-bottom")).toBe("calc(24px + var(--edge-bottom))");
  });
});
