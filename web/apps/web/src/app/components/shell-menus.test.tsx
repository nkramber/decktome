import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { beforeEach, describe, expect, it, vi } from "vitest";

const push = vi.hoisted(() => ({
  pushConfigured: vi.fn(() => true),
  pushState: vi.fn(),
  enablePush: vi.fn(),
  disablePush: vi.fn(),
}));
vi.mock("../../features/push/push", () => push);
const notify = vi.hoisted(() => vi.fn());
vi.mock("./notify", () => ({ notify }));

import { AccountMenuContent } from "./shell-menus";

function renderMenu() {
  return render(<AccountMenuContent trigger={<button type="button">Account</button>} email="reader@example.com" onSignOut={() => {}} open onOpenChange={() => {}} />);
}

beforeEach(() => {
  push.pushConfigured.mockReturnValue(true);
  push.pushState.mockReset();
  push.enablePush.mockReset();
  push.disablePush.mockReset();
  notify.mockReset();
});

// The push toggle of PR-26 (D-1005).
describe("the push item of the account menu", () => {
  it("turns push on from the off state", async () => {
    push.pushState.mockResolvedValue("off");
    push.enablePush.mockResolvedValue("on");
    renderMenu();
    await userEvent.setup().click(await screen.findByRole("menuitem", { name: "Notify me when a deck is ready" }));
    expect(push.enablePush).toHaveBeenCalledTimes(1);
    await waitFor(() => expect(notify).toHaveBeenCalledWith("success", "Notifications on", expect.any(String)));
  });

  it("turns push off from the on state", async () => {
    push.pushState.mockResolvedValue("on");
    push.disablePush.mockResolvedValue("off");
    renderMenu();
    await userEvent.setup().click(await screen.findByRole("menuitem", { name: "Stop deck notifications" }));
    expect(push.disablePush).toHaveBeenCalledTimes(1);
  });

  it("says so when the browser blocks the notifications", async () => {
    push.pushState.mockResolvedValue("blocked");
    renderMenu();
    expect(await screen.findByRole("menuitem", { name: "Notifications are blocked" })).toHaveAttribute("aria-disabled", "true");
  });

  it("shows no item where the browser takes no push", async () => {
    push.pushState.mockResolvedValue("unavailable");
    renderMenu();
    await screen.findByRole("menuitem", { name: "Sign out" });
    await waitFor(() => expect(push.pushState).toHaveBeenCalled());
    expect(screen.queryByRole("menuitem", { name: /notif/i })).not.toBeInTheDocument();
  });

  it("shows no item in a build with no sender id", async () => {
    push.pushConfigured.mockReturnValue(false);
    renderMenu();
    await screen.findByRole("menuitem", { name: "Sign out" });
    expect(push.pushState).not.toHaveBeenCalled();
  });

  it("reports a failure to turn push on", async () => {
    push.pushState.mockResolvedValue("off");
    push.enablePush.mockRejectedValue(new Error("offline"));
    renderMenu();
    await userEvent.setup().click(await screen.findByRole("menuitem", { name: "Notify me when a deck is ready" }));
    await waitFor(() => expect(notify).toHaveBeenCalledWith("error", "Notifications could not be turned on", expect.any(String)));
  });
});
