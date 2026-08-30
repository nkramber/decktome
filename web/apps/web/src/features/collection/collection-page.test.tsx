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
const getCollection = vi.fn();
const getCards = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  cardClient: { getCards: (...args: unknown[]) => getCards(...args) },
  collectionClient: {
    importCollection: (...args: unknown[]) => importCollection(...args),
    listCollections: (...args: unknown[]) => listCollections(...args),
    getCollection: (...args: unknown[]) => getCollection(...args),
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
  getCollection.mockReset();
  getCards.mockReset();
  listCollections.mockResolvedValue({ collections: earlier });
  getCollection.mockResolvedValue({
    collection: {
      id: "c-old",
      name: "binder-july.csv",
      cardCount: 4317,
      importedAt: { seconds: 1756000000n, nanos: 0 },
      entries: [
        { oracleId: "o-bolt", name: "Lightning Bolt", quantity: 4, rarity: "common", setName: "Alpha", setCode: "lea" },
        { oracleId: "o-jace", name: "Jace, the Mind Sculptor", quantity: 1, rarity: "mythic", setName: "Worldwake", setCode: "wwk" },
        { oracleId: "o-bolt", name: "Lightning Bolt", quantity: 2, rarity: "common", setName: "Alpha", setCode: "lea" },
      ],
    },
  });
  getCards.mockResolvedValue({
    cards: [{ oracleId: "o-jace", name: "Jace, the Mind Sculptor", faces: [], defaultPrinting: { artist: "A", imageUris: { artCrop: "https://x/a.jpg" } } }],
    missingOracleIds: [],
  });
});

describe("the binder head", () => {
  it("shows nothing until a collection is active", async () => {
    await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    expect(getCollection).not.toHaveBeenCalled();
  });

  it("counts the cards, the unique cards, and the rarity of the active collection", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned" });
    await renderAt("/collection");
    // 4 + 1 + 2 = 7 cards over two Oracle ids.
    expect(await screen.findByText("7")).toBeInTheDocument();
    expect(screen.getByText("Unique cards")).toBeInTheDocument();
    expect(screen.getByText("2")).toBeInTheDocument();
    expect(screen.getByText(/Mythic/)).toBeInTheDocument();
    expect(screen.getByText(/Common/)).toBeInTheDocument();
    expect(getCollection).toHaveBeenCalledWith({ collectionId: "c-old" });
  });

  it("asks for the art of the rarest cards first", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned" });
    await renderAt("/collection");
    await screen.findByText("7");
    // The mythic sorts before the common.
    expect(getCards).toHaveBeenCalledWith({ oracleIds: ["o-jace", "o-bolt"] });
  });
});

describe("CollectionPage", () => {
  it("lists earlier uploads and picks one as the active collection", async () => {
    await renderAt("/collection");
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
        // The server keys the map by the full enum name.
        unresolvedByReason: { UNRESOLVED_REASON_UNKNOWN_CARD: 1, UNRESOLVED_REASON_NOT_PLAYABLE: 1 },
        unresolved: [
          { line: 4, raw: "Not A Card,XYZ,1", reason: UnresolvedReason.UNKNOWN_CARD },
          { line: 5, raw: "Pawpatch Recruit,TBLB,1", reason: UnresolvedReason.NOT_PLAYABLE },
        ],
      },
    });
    await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    const user = userEvent.setup();
    const file = new File(["Name,Set code\nLightning Bolt,LEA\n"], "export.csv", { type: "text/csv" });
    await user.upload(screen.getByLabelText("ManaBox CSV file"), file);
    await user.click(screen.getByRole("button", { name: "Upload" }));

    expect(await screen.findByTestId("card-count")).toHaveTextContent(
      "export.csv: 3 cards, 2 rows resolved, 2 unresolved.",
    );
    const req = importCollection.mock.calls[0][0] as { name: string; source: ImportSource; content: Uint8Array };
    expect(req.name).toBe("export.csv");
    expect(req.source).toBe(ImportSource.MANABOX_CSV);
    expect(new TextDecoder().decode(req.content)).toContain("Lightning Bolt");

    const table = screen.getByRole("table", { name: /open your file at this line/ });
    const row = within(table).getAllByRole("row")[1];
    expect(row).toHaveTextContent("4");
    // The raw row stays out of the table. A user reads the line in their own file.
    expect(row).not.toHaveTextContent("Not A Card,XYZ,1");
    expect(row).toHaveTextContent("Unknown card");
    expect(row).not.toHaveTextContent("UNRESOLVED_REASON");
    expect(screen.getByText(/^Unknown card: .*: 1$/)).toBeInTheDocument();
    // A token row gets its own words, and no extra note.
    expect(within(table).getAllByRole("row")[2]).toHaveTextContent("Not a playable card: a token, emblem, or art card");
    expect(screen.queryByTestId("token-note")).not.toBeInTheDocument();
    expect(useAppStore.getState().collectionId).toBe("c-new");
  });

  it("uses the typed name over the file name", async () => {
    importCollection.mockResolvedValue({ collection: { id: "c-new", name: "Mine", cardCount: 1 }, report: {} });
    await renderAt("/collection");
    const user = userEvent.setup();
    await user.type(screen.getByLabelText(/Collection name/), "Mine");
    await user.upload(screen.getByLabelText("ManaBox CSV file"), new File(["x"], "export.csv"));
    await user.click(screen.getByRole("button", { name: "Upload" }));
    await screen.findByTestId("card-count");
    expect((importCollection.mock.calls[0][0] as { name: string }).name).toBe("Mine");
  });

  it("shows the RPC error when the upload fails", async () => {
    importCollection.mockRejectedValue(new Error("[unauthenticated] no token"));
    await renderAt("/collection");
    const user = userEvent.setup();
    await user.upload(screen.getByLabelText("ManaBox CSV file"), new File(["x"], "export.csv"));
    await user.click(screen.getByRole("button", { name: "Upload" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("Upload failed: [unauthenticated] no token");
  });

  it("refuses a file over the 5 MiB upload cap before the upload", async () => {
    await renderAt("/collection");
    const user = userEvent.setup();
    const big = new File([new Uint8Array((5 << 20) + 1)], "big.csv", { type: "text/csv" });
    await user.upload(screen.getByLabelText("ManaBox CSV file"), big);
    expect(await screen.findByRole("alert")).toHaveTextContent("The file is 5.0 MiB. The limit is 5 MiB.");
    expect(screen.getByRole("button", { name: "Upload" })).toBeDisabled();
    expect(importCollection).not.toHaveBeenCalled();
  });

  it("clears the picked file on Skip and on a click on an earlier upload", async () => {
    const { router } = await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    const user = userEvent.setup();
    await user.upload(screen.getByLabelText("ManaBox CSV file"), new File(["x"], "export.csv"));
    expect(screen.getByRole("button", { name: "Upload" })).toBeEnabled();
    await user.click(screen.getByRole("button", { name: "binder-july.csv" }));
    expect(screen.getByRole("button", { name: "Upload" })).toBeDisabled();
    expect((screen.getByLabelText("ManaBox CSV file") as HTMLInputElement).files).toHaveLength(0);
    await user.upload(screen.getByLabelText("ManaBox CSV file"), new File(["x"], "export.csv"));
    expect(screen.getByRole("button", { name: "Upload" })).toBeEnabled();
    await user.click(screen.getByRole("button", { name: "Skip, build from any card" }));
    expect(router.state.location.pathname).toBe("/session/new");
  });

  it("skip clears the active collection and goes to the chat (D-37)", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned" });
    const { router } = await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    await userEvent.setup().click(screen.getByRole("button", { name: "Skip, build from any card" }));
    expect(useAppStore.getState().collectionId).toBe("");
    expect(useAppStore.getState().poolMode).toBe("any");
    expect(router.state.location.pathname).toBe("/session/new");
    expect(await screen.findByTestId("session-id")).toHaveTextContent("No session yet.");
  });

  it("continue to chat goes to /session/new", async () => {
    const { router } = await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    await userEvent.setup().click(screen.getByRole("button", { name: "Continue to chat" }));
    expect(router.state.location.pathname).toBe("/session/new");
  });

  it("has no axe violations", async () => {
    const { container } = await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    expect(await axe(container)).toHaveNoViolations();
  });
});
