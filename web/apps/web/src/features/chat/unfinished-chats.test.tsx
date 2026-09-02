import { screen, waitFor, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { useAppStore } from "../../lib/store";
import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const listSessions = vi.fn();
const updateSession = vi.fn();
const deleteSession = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }) },
  agentClient: {
    listSessions: (...args: unknown[]) => listSessions(...args),
    updateSession: (...args: unknown[]) => updateSession(...args),
    deleteSession: (...args: unknown[]) => deleteSession(...args),
    chat: vi.fn(),
  },
  deckClient: { listDecks: () => Promise.resolve({ decks: [], nextPageToken: "" }) },
}));

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  useAppStore.setState({ collectionId: "", sessionId: "", poolMode: "any" });
  listSessions.mockReset();
  updateSession.mockReset();
  deleteSession.mockReset();
  listSessions.mockResolvedValue({
    sessions: [
      { id: "s-open", name: "", firstMessage: "Elves that go wide.", deckCount: 0, status: 1, usage: { priced: true, costUsd: 0.0012, calls: 3 }, updatedAt: { seconds: 1756800000n, nanos: 0 } },
      { id: "s-done", name: "Done one", firstMessage: "A finished deck", deckCount: 1, status: 3 },
    ],
    nextPageToken: "",
  });
});

describe("the unfinished chats", () => {
  it("renders nothing when every chat built a deck (D-438)", async () => {
    listSessions.mockResolvedValue({ sessions: [{ id: "s-done", name: "Done one", firstMessage: "A finished deck", deckCount: 1, status: 3 }], nextPageToken: "" });
    await renderAt("/session/new");
    await screen.findByRole("heading", { level: 1, name: "New deck" });
    await new Promise((r) => setTimeout(r, 50));
    expect(screen.queryByRole("heading", { name: "Unfinished chats" })).not.toBeInTheDocument();
  });

  it("sits under the message box", async () => {
    await renderAt("/session/new");
    const list = await screen.findByTestId("unfinished-chats");
    const box = screen.getByLabelText("Your message");
    expect(box.compareDocumentPosition(list) & Node.DOCUMENT_POSITION_FOLLOWING).toBeTruthy();
  });

  it("lists the chats with no deck, and none of the finished ones (D-433)", async () => {
    await renderAt("/session/new");
    const list = await screen.findByTestId("unfinished-chats");
    expect(within(list).getByRole("link", { name: "Elves that go wide." })).toHaveAttribute("href", "/session/s-open");
    expect(within(list).queryByText("Done one")).not.toBeInTheDocument();
    expect(list).toHaveTextContent("$0.0012");
  });

  it("renames a chat", async () => {
    updateSession.mockResolvedValue({ session: { id: "s-open", name: "Elves" } });
    const user = userEvent.setup();
    await renderAt("/session/new");
    await user.click(await screen.findByRole("button", { name: "Rename Elves that go wide." }));
    const field = screen.getByRole("textbox", { name: /New name for/ });
    await user.clear(field);
    await user.type(field, "Elves");
    await user.click(screen.getByRole("button", { name: "Save" }));
    await waitFor(() => expect(updateSession).toHaveBeenCalledWith({ sessionId: "s-open", name: "Elves" }));
  });

  it("asks before it deletes a chat", async () => {
    deleteSession.mockResolvedValue({});
    const user = userEvent.setup();
    await renderAt("/session/new");
    await user.click(await screen.findByRole("button", { name: "Delete Elves that go wide." }));
    await user.click(await screen.findByRole("button", { name: "Delete the chat" }));
    await waitFor(() => expect(deleteSession).toHaveBeenCalledWith({ sessionId: "s-open" }));
  });
});
