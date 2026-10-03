import { create } from "@bufbuild/protobuf";
import { timestampFromDate } from "@bufbuild/protobuf/wkt";
import { Code, ConnectError } from "@connectrpc/connect";
import { AccessRequestSchema, AccessStatus } from "@mtg/api-client/mtg/v1/admin_service_pb";
import { QueryClientProvider } from "@tanstack/react-query";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { makeQueryClient } from "../../lib/query-client";
import { AdminPage, approvedSent, approvedUnsent } from "./admin-page";

const listAccessRequests = vi.fn();
const approveAccessRequest = vi.fn();
const dismissAccessRequest = vi.fn();
vi.mock("../../lib/api", () => ({
  adminClient: {
    listAccessRequests: (...a: unknown[]) => listAccessRequests(...a),
    approveAccessRequest: (...a: unknown[]) => approveAccessRequest(...a),
    dismissAccessRequest: (...a: unknown[]) => dismissAccessRequest(...a),
  },
}));
const notify = vi.fn();
vi.mock("../../app/components/notify", () => ({ notify: (...a: unknown[]) => notify(...a) }));

const bob = create(AccessRequestSchema, {
  email: "bob@example.com",
  note: "Elves",
  status: AccessStatus.PENDING,
  count: 2n,
  createdAt: timestampFromDate(new Date("2026-10-02T12:00:00Z")),
});

function renderPage() {
  return render(
    <QueryClientProvider client={makeQueryClient()}>
      <AdminPage />
    </QueryClientProvider>,
  );
}

beforeEach(() => {
  listAccessRequests.mockReset();
  listAccessRequests.mockResolvedValue({ requests: [bob] });
  approveAccessRequest.mockReset();
  approveAccessRequest.mockResolvedValue({ emailSent: true });
  dismissAccessRequest.mockReset();
  dismissAccessRequest.mockResolvedValue({});
  notify.mockReset();
  notify.mockResolvedValue(undefined);
});

// D-1076: the admin screen lists the pending requests and decides each.
describe("AdminPage", () => {
  it("lists the pending requests with the note and the count", async () => {
    renderPage();
    expect(await screen.findByText("bob@example.com")).toBeInTheDocument();
    expect(screen.getByText("Elves")).toBeInTheDocument();
    expect(screen.getByText(/2 requests/)).toBeInTheDocument();
    expect(listAccessRequests).toHaveBeenCalledWith({ status: AccessStatus.PENDING });
  });

  it("approves a request and says the email went out", async () => {
    const user = userEvent.setup();
    renderPage();
    await user.click(await screen.findByRole("button", { name: "Approve" }));
    expect(approveAccessRequest).toHaveBeenCalledWith({ email: "bob@example.com" });
    expect(notify).toHaveBeenCalledWith("success", approvedSent);
  });

  it("says so when the approval email did not send", async () => {
    approveAccessRequest.mockResolvedValue({ emailSent: false });
    const user = userEvent.setup();
    renderPage();
    await user.click(await screen.findByRole("button", { name: "Approve" }));
    expect(notify).toHaveBeenCalledWith("success", approvedUnsent);
  });

  it("dismisses a request", async () => {
    const user = userEvent.setup();
    renderPage();
    await user.click(await screen.findByRole("button", { name: "Dismiss" }));
    expect(dismissAccessRequest).toHaveBeenCalledWith({ email: "bob@example.com" });
  });

  it("reads another status", async () => {
    const user = userEvent.setup();
    renderPage();
    await screen.findByText("bob@example.com");
    listAccessRequests.mockResolvedValue({ requests: [] });
    await user.click(screen.getByRole("button", { name: "Approved" }));
    expect(listAccessRequests).toHaveBeenLastCalledWith({ status: AccessStatus.APPROVED });
    expect(await screen.findByText("No requests here")).toBeInTheDocument();
  });

  it("shows the refusal to a reader with no admin claim", async () => {
    listAccessRequests.mockRejectedValue(new ConnectError("this screen is open to the admin alone", Code.PermissionDenied));
    renderPage();
    expect(await screen.findByText(/open to the admin alone/)).toBeInTheDocument();
  });

  it("has no axe violations", async () => {
    const { container } = renderPage();
    await screen.findByText("bob@example.com");
    expect(await axe(container)).toHaveNoViolations();
  });
});
