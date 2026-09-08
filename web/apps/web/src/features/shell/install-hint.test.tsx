import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { useAppStore } from "../../lib/store";
import { InstallHint } from "./install-hint";

// The hint shows once, on a phone, after the first deck (PR-25). Every
// other reader sees nothing.

// phone answers the two media queries the hint reads: a narrow screen
// with a coarse pointer, and a window that is not standalone.
function phone({ narrow = true, standalone = false }: { narrow?: boolean; standalone?: boolean } = {}) {
  vi.stubGlobal("matchMedia", (q: string) => ({
    matches: q.includes("display-mode") ? standalone : narrow,
    media: q,
    addEventListener: vi.fn(),
    removeEventListener: vi.fn(),
  }));
}

beforeEach(() => {
  localStorage.clear();
  useAppStore.setState({ hadDecks: true, installHintDismissed: false });
  phone();
});

describe("InstallHint", () => {
  it("names the two taps of the iOS share sheet, and passes axe", async () => {
    const view = render(<InstallHint />);
    expect(screen.getByRole("complementary", { name: "Add Deck Tome to your Home Screen" })).toBeInTheDocument();
    expect(screen.getByText(/Tap Share, then Add to Home Screen/)).toBeInTheDocument();
    expect(await axe(view.container)).toHaveNoViolations();
  });

  it("shows nothing before the reader has a deck", () => {
    useAppStore.setState({ hadDecks: false });
    render(<InstallHint />);
    expect(screen.queryByRole("complementary")).toBeNull();
  });

  it("shows nothing on a wide screen or in an installed window", () => {
    phone({ narrow: false });
    const wide = render(<InstallHint />);
    expect(wide.queryByRole("complementary")).toBeNull();
    wide.unmount();

    phone({ standalone: true });
    render(<InstallHint />);
    expect(screen.queryByRole("complementary")).toBeNull();
  });

  it("a dismissal holds, and the hint never comes back", async () => {
    const view = render(<InstallHint />);
    await userEvent.click(screen.getByRole("button", { name: "Dismiss the install hint" }));
    expect(view.queryByRole("complementary")).toBeNull();
    expect(useAppStore.getState().installHintDismissed).toBe(true);

    // A second mount reads the store and stays quiet.
    view.unmount();
    render(<InstallHint />);
    expect(screen.queryByRole("complementary")).toBeNull();
  });

  it("installs on a tap where the browser offers it", async () => {
    const prompt = vi.fn().mockResolvedValue(undefined);
    render(<InstallHint />);
    const ev = new Event("beforeinstallprompt");
    Object.assign(ev, { prompt, userChoice: Promise.resolve({ outcome: "accepted" }) });
    window.dispatchEvent(ev);

    await userEvent.click(await screen.findByRole("button", { name: "Add to Home Screen" }));
    expect(prompt).toHaveBeenCalled();
    // The hint asked, so it does not ask again.
    expect(useAppStore.getState().installHintDismissed).toBe(true);
  });
});
