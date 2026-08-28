import { ImportSource, UnresolvedReason } from "@mtg/api-client/mtg/v1/collection_pb";
import { screen, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { useAppStore } from "../../lib/store";
import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const importCollection = vi.fn();
const listCollections = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: {
    importCollection: (...args: unknown[]) => importCollection(...args),
    listCollections: (...args: unknown[]) => listCollections(...args),
  },
}));

const earlier = [
  { id: "c-old", name: "binder-july.csv", cardCount: 4317, importedAt: { seconds: 1756000000n, nanos: 0 } },
  { id: "c-older", name: "binder-june.csv", cardCount: 12, importedAt: undefined },
];

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  useAppStore.setState({ collectionId: "", sessionId: "", poolMode: "any" });
  importCollection.mockReset();
  listCollections.mockReset();
  listCollections.mockResolvedValue({ collections: earlier });
});

describe("CollectionPage", () => {
  it("lists earlier uploads and picks one as the active collection", async () => {
    renderAt("/collection");
    expect(await screen.findByRole("button", { name: "binder-july.csv" })).toBeInTheDocument();
    expect(screen.getByText(/4317 cards/)).toBeInTheDocument();
    await userEvent.setup().click(screen.getByRole("button", { name: "binder-july.csv" }));
    expect(useAppStore.getState().collectionId).toBe("c-old");
    expect(useAppStore.getState().poolMode).toBe("owned");
    expect(screen.getByTestId("active-collection")).toHaveTextContent("Active collection: binder-july.csv (4317 cards)");
    expect(screen.getByRole("button", { name: "binder-july.csv" })).toHaveAttribute("aria-pressed", "true");
  });

  it("uploads a CSV and shows the count and the unresolved rows", async () => {
    importCollection.mockResolvedValue({
      collection: { id: "c-new", name: "export.csv", cardCount: 3 },
      report: {
        resolvedCount: 2,
        unresolvedByReason: { UNKNOWN_CARD: 1 },
        unresolved: [{ line: 4, raw: "Not A Card,XYZ,1", reason: UnresolvedReason.UNKNOWN_CARD }],
      },
    });
    renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    const user = userEvent.setup();
    const file = new File(["Name,Set code\nLightning Bolt,LEA\n"], "export.csv", { type: "text/csv" });
    await user.upload(screen.getByLabelText("ManaBox CSV file"), file);
    await user.click(screen.getByRole("button", { name: "Upload" }));

    expect(await screen.findByTestId("card-count")).toHaveTextContent(
      "export.csv: 3 cards, 2 rows resolved, 1 unresolved.",
    );
    const req = importCollection.mock.calls[0][0] as { name: string; source: ImportSource; content: Uint8Array };
    expect(req.name).toBe("export.csv");
    expect(req.source).toBe(ImportSource.MANABOX_CSV);
    expect(new TextDecoder().decode(req.content)).toContain("Lightning Bolt");

    const table = screen.getByRole("table", { name: "Unresolved rows" });
    const row = within(table).getAllByRole("row")[1];
    expect(row).toHaveTextContent("4");
    expect(row).toHaveTextContent("Not A Card,XYZ,1");
    expect(row).toHaveTextContent("UNKNOWN_CARD");
    expect(screen.getByText("UNKNOWN_CARD: 1")).toBeInTheDocument();
    expect(useAppStore.getState().collectionId).toBe("c-new");
  });

  it("uses the typed name over the file name", async () => {
    importCollection.mockResolvedValue({ collection: { id: "c-new", name: "Mine", cardCount: 1 }, report: {} });
    renderAt("/collection");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText(/Collection name/), "Mine");
    await user.upload(screen.getByLabelText("ManaBox CSV file"), new File(["x"], "export.csv"));
    await user.click(screen.getByRole("button", { name: "Upload" }));
    await screen.findByTestId("card-count");
    expect((importCollection.mock.calls[0][0] as { name: string }).name).toBe("Mine");
  });

  it("shows the RPC error when the upload fails", async () => {
    importCollection.mockRejectedValue(new Error("[unauthenticated] no token"));
    renderAt("/collection");
    const user = userEvent.setup();
    await user.upload(screen.getByLabelText("ManaBox CSV file"), new File(["x"], "export.csv"));
    await user.click(screen.getByRole("button", { name: "Upload" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("Upload failed: [unauthenticated] no token");
  });

  it("skip clears the active collection and goes to the chat (D-37)", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned" });
    const { router } = renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    await userEvent.setup().click(screen.getByRole("button", { name: "Skip, build from any card" }));
    expect(useAppStore.getState().collectionId).toBe("");
    expect(useAppStore.getState().poolMode).toBe("any");
    expect(router.state.location.pathname).toBe("/session/new");
    expect(await screen.findByTestId("session-id")).toHaveTextContent("No session yet.");
  });

  it("continue to chat goes to /session/new", async () => {
    const { router } = renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    await userEvent.setup().click(screen.getByRole("button", { name: "Continue to chat" }));
    expect(router.state.location.pathname).toBe("/session/new");
  });

  it("has no axe violations", async () => {
    const { container } = renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    expect(await axe(container)).toHaveNoViolations();
  });
});
