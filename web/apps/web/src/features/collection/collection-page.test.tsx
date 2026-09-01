import { ImportSource, UnresolvedReason } from "@mtg/api-client/mtg/v1/collection_pb";
import { screen, waitFor, within } from "@testing-library/react";
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
const deleteCollection = vi.fn();
const updateCollection = vi.fn();
const diffCollections = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  cardClient: { getCards: (...args: unknown[]) => getCards(...args) },
  collectionClient: {
    importCollection: (...args: unknown[]) => importCollection(...args),
    listCollections: (...args: unknown[]) => listCollections(...args),
    getCollection: (...args: unknown[]) => getCollection(...args),
    deleteCollection: (...args: unknown[]) => deleteCollection(...args),
    updateCollection: (...args: unknown[]) => updateCollection(...args),
    diffCollections: (...args: unknown[]) => diffCollections(...args),
  },
}));

const earlier = [
  { id: "c-old", name: "binder-july.csv", cardCount: 4317, importedAt: { seconds: 1756000000n, nanos: 0 } },
  { id: "c-older", name: "binder-june.csv", cardCount: 12, importedAt: undefined },
];

// The upload lives in a dialog (roadmap PR-18). openUpload opens it and
// gives back the file input inside it.
async function openUpload(user: ReturnType<typeof userEvent.setup>) {
  await user.click(await screen.findByRole("button", { name: "Upload a collection" }));
  return await screen.findByLabelText("ManaBox CSV file");
}

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  useAppStore.setState({ collectionId: "", sessionId: "", poolMode: "any" });
  importCollection.mockReset();
  listCollections.mockReset();
  getCollection.mockReset();
  getCards.mockReset();
  deleteCollection.mockReset();
  updateCollection.mockReset();
  diffCollections.mockReset();
  listCollections.mockResolvedValue({ collections: earlier });
  // The head reads the summary the import stored, and never an entry
  // (D-392). A paged read carries the rows the binder grid shows.
  getCollection.mockImplementation((req: { entriesOmitted?: boolean }) =>
    Promise.resolve(
      req.entriesOmitted
        ? {
            collection: {
              id: "c-old",
              name: "binder-july.csv",
              cardCount: 7,
              importedAt: { seconds: 1756000000n, nanos: 0 },
              entries: [],
              summary: {
                rowCount: 3,
                uniqueCards: 2,
                byRarity: { common: 6, mythic: 1 },
                byColor: {},
                topSets: [{ setCode: "lea", setName: "Alpha", count: 6 }],
                artOracleIds: ["o-jace", "o-bolt"],
              },
            },
          }
        : {
            collection: {
              id: "c-old",
              name: "binder-july.csv",
              cardCount: 7,
              importedAt: { seconds: 1756000000n, nanos: 0 },
              entries: [
                { oracleId: "o-bolt", name: "Lightning Bolt", quantity: 4, rarity: "common", setName: "Alpha", setCode: "lea", scryfallId: "p1", collectorNumber: "1" },
                { oracleId: "o-jace", name: "Jace, the Mind Sculptor", quantity: 1, rarity: "mythic", setName: "Worldwake", setCode: "wwk", scryfallId: "p2", collectorNumber: "2" },
                { oracleId: "o-bolt", name: "Lightning Bolt", quantity: 2, rarity: "common", setName: "Alpha", setCode: "lea", scryfallId: "p3", collectorNumber: "3" },
              ],
            },
            nextPageToken: "",
          },
    ),
  );
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

  it("reads the stored summary and asks for no entry (D-392)", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned_only" });
    await renderAt("/collection");
    // 4 + 1 + 2 = 7 cards over two Oracle ids, all from the summary.
    expect(await screen.findByText("7")).toBeInTheDocument();
    expect(screen.getByText("Unique cards")).toBeInTheDocument();
    expect(screen.getByText("2")).toBeInTheDocument();
    expect(screen.getByText(/Mythic/)).toBeInTheDocument();
    expect(screen.getByText(/Common/)).toBeInTheDocument();
    // The head asks for the collection without its entries. Before
    // D-392 it read every row and counted them here.
    expect(getCollection).toHaveBeenCalledWith({ collectionId: "c-old", entriesOmitted: true });
  });

  it("asks for the art the summary names, rarest first", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned_only" });
    await renderAt("/collection");
    await screen.findByText("7");
    // The import picked them, and the mythic sorts before the common.
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
    expect(useAppStore.getState().poolMode).toBe("owned_first");
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
    await user.upload(await openUpload(user), file);
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
    const input = await openUpload(user);
    await user.type(screen.getByLabelText(/Collection name/), "Mine");
    await user.upload(input, new File(["x"], "export.csv"));
    await user.click(screen.getByRole("button", { name: "Upload" }));
    await screen.findByTestId("card-count");
    expect((importCollection.mock.calls[0][0] as { name: string }).name).toBe("Mine");
  });

  it("shows the RPC error when the upload fails", async () => {
    importCollection.mockRejectedValue(new Error("[unauthenticated] no token"));
    await renderAt("/collection");
    const user = userEvent.setup();
    await user.upload(await openUpload(user), new File(["x"], "export.csv"));
    await user.click(screen.getByRole("button", { name: "Upload" }));
    expect(await screen.findByRole("alert")).toHaveTextContent("Upload failed: [unauthenticated] no token");
  });

  it("refuses a file over the 5 MiB upload cap before the upload", async () => {
    await renderAt("/collection");
    const user = userEvent.setup();
    const big = new File([new Uint8Array((5 << 20) + 1)], "big.csv", { type: "text/csv" });
    await user.upload(await openUpload(user), big);
    expect(await screen.findByRole("alert")).toHaveTextContent("The file is 5.0 MiB. The limit is 5 MiB.");
    expect(screen.getByRole("button", { name: "Upload" })).toBeDisabled();
    expect(importCollection).not.toHaveBeenCalled();
  });

  it("drops the picked file when the dialog closes, and Skip leaves the page", async () => {
    const { router } = await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    const user = userEvent.setup();
    await user.upload(await openUpload(user), new File(["x"], "export.csv"));
    expect(screen.getByRole("button", { name: "Upload" })).toBeEnabled();

    // Cancel closes the dialog. A second open starts clean, so no file
    // of an abandoned upload reaches the server.
    await user.click(screen.getByRole("button", { name: "Cancel" }));
    await waitFor(() => expect(screen.queryByLabelText("ManaBox CSV file")).not.toBeInTheDocument());
    const again = await openUpload(user);
    expect((again as HTMLInputElement).files).toHaveLength(0);
    expect(screen.getByRole("button", { name: "Upload" })).toBeDisabled();
    expect(importCollection).not.toHaveBeenCalled();

    await user.keyboard("{Escape}");
    await waitFor(() => expect(screen.queryByLabelText("ManaBox CSV file")).not.toBeInTheDocument());
    await user.click(screen.getByRole("button", { name: "Skip, build from any card" }));
    expect(router.state.location.pathname).toBe("/session/new");
  });

  it("skip clears the active collection and goes to the chat (D-37)", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned_only" });
    const { router } = await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    await userEvent.setup().click(screen.getByRole("button", { name: "Skip, build from any card" }));
    expect(useAppStore.getState().collectionId).toBe("");
    expect(useAppStore.getState().poolMode).toBe("any");
    expect(router.state.location.pathname).toBe("/session/new");
    expect(await screen.findByRole("heading", { level: 1, name: "New deck" })).toBeInTheDocument();
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

describe("the file dialog on arrival", () => {
  it("stays shut when the reader came here on their own", async () => {
    const clicks: string[] = [];
    const realClick = HTMLInputElement.prototype.click;
    HTMLInputElement.prototype.click = function click(this: HTMLInputElement) {
      clicks.push(this.type);
    };
    try {
      await renderAt("/collection");
      await screen.findByRole("button", { name: "binder-july.csv" });
      expect(clicks).not.toContain("file");
    } finally {
      HTMLInputElement.prototype.click = realClick;
    }
  });
});

// Deleting a collection is for good (D-347). The decks it built keep
// their cards, and their chats fall back to the whole card database.
describe("deleting a collection", () => {
  it("asks first, then removes it and clears the active choice", async () => {
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned_only" });
    deleteCollection.mockResolvedValue({});
    await renderAt("/collection");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Delete binder-july.csv" }));
    expect(await screen.findByText(/Every deck you built from it keeps all of its cards/)).toBeInTheDocument();
    await user.click(screen.getByRole("button", { name: "Delete the collection" }));
    await waitFor(() => expect(deleteCollection).toHaveBeenCalledWith({ collectionId: "c-old" }));
    await waitFor(() => expect(useAppStore.getState().collectionId).toBe(""));
    expect(useAppStore.getState().poolMode).toBe("any");
  });

  it("keeps the collection when the reader backs out", async () => {
    await renderAt("/collection");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Delete binder-july.csv" }));
    await user.click(await screen.findByRole("button", { name: "Keep it" }));
    expect(deleteCollection).not.toHaveBeenCalled();
  });
});

describe("renaming a collection", () => {
  it("writes the new name and closes the form", async () => {
    updateCollection.mockResolvedValue({ collection: { id: "c-old", name: "My binder" } });
    const user = userEvent.setup();
    await renderAt("/collection");
    await user.click(await screen.findByRole("button", { name: "Rename binder-july.csv" }));

    const field = screen.getByRole("textbox", { name: "New name for binder-july.csv" });
    expect(field).toHaveValue("binder-july.csv");
    await user.clear(field);
    await user.type(field, "My binder");
    await user.click(screen.getByRole("button", { name: "Save" }));

    await waitFor(() => expect(updateCollection).toHaveBeenCalledWith({ collectionId: "c-old", name: "My binder" }));
    await waitFor(() => expect(screen.queryByRole("textbox", { name: /New name/ })).not.toBeInTheDocument());
  });

  it("refuses an empty name, and Cancel writes nothing", async () => {
    const user = userEvent.setup();
    await renderAt("/collection");
    await user.click(await screen.findByRole("button", { name: "Rename binder-july.csv" }));
    const field = screen.getByRole("textbox", { name: "New name for binder-july.csv" });
    await user.clear(field);
    expect(screen.getByRole("button", { name: "Save" })).toBeDisabled();
    await user.click(screen.getByRole("button", { name: "Cancel" }));
    expect(updateCollection).not.toHaveBeenCalled();
  });
});

describe("a re-upload over an active collection", () => {
  const file = () => new File(["Name,Set code\nBolt,LEA\n"], "binder-august.csv", { type: "text/csv" });

  it("shows what changes before it replaces anything (D-393)", async () => {
    diffCollections.mockResolvedValue({
      diff: {
        added: [{ name: "Path to Exile", quantity: 2 }],
        removed: [],
        changed: [{ entry: { name: "Lightning Bolt" }, from: 4, to: 2 }],
        addedCards: 2, removedCards: 0, changedCards: 2, identical: false,
      },
      report: { unresolved: [], resolvedCount: 1, unresolvedByReason: {} },
    });
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned_only" });
    const user = userEvent.setup();
    await renderAt("/collection");

    await user.upload(await openUpload(user), file());
    await user.click(screen.getByRole("button", { name: "Upload" }));

    // The diff reads, and nothing is imported yet.
    expect(await screen.findByTestId("collection-diff")).toBeInTheDocument();
    expect(screen.getByTestId("diff-added")).toHaveTextContent("1 row");
    expect(screen.getByTestId("diff-changed")).toHaveTextContent("1 row");
    expect(screen.getByText("Lightning Bolt: 4 → 2")).toBeInTheDocument();
    expect(importCollection).not.toHaveBeenCalled();

    // Replace names the collection it replaces, so the id survives.
    importCollection.mockResolvedValue({ collection: { id: "c-old", name: "binder-july.csv", cardCount: 5 }, report: { unresolved: [], resolvedCount: 1, unresolvedByReason: {} } });
    await user.click(screen.getByRole("button", { name: "Replace the collection" }));
    await waitFor(() =>
      expect(importCollection).toHaveBeenCalledWith(expect.objectContaining({ replaceCollectionId: "c-old" })),
    );
  });

  it("says so when nothing changed, and offers no replace", async () => {
    diffCollections.mockResolvedValue({
      diff: { added: [], removed: [], changed: [], addedCards: 0, removedCards: 0, changedCards: 0, identical: true },
      report: { unresolved: [], resolvedCount: 1, unresolvedByReason: {} },
    });
    useAppStore.setState({ collectionId: "c-old", poolMode: "owned_only" });
    const user = userEvent.setup();
    await renderAt("/collection");
    await user.upload(await openUpload(user), file());
    await user.click(screen.getByRole("button", { name: "Upload" }));

    expect(await screen.findByText("Nothing changed")).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Replace the collection" })).not.toBeInTheDocument();
  });

  it("imports straight away when no collection is active", async () => {
    importCollection.mockResolvedValue({ collection: { id: "c-new", name: "binder-august.csv", cardCount: 1 }, report: { unresolved: [], resolvedCount: 1, unresolvedByReason: {} } });
    const user = userEvent.setup();
    await renderAt("/collection");
    await user.upload(await openUpload(user), file());
    await user.click(screen.getByRole("button", { name: "Upload" }));
    await waitFor(() => expect(importCollection).toHaveBeenCalled());
    expect(diffCollections).not.toHaveBeenCalled();
    expect(importCollection).toHaveBeenCalledWith(expect.objectContaining({ replaceCollectionId: "" }));
  });
});

// The upload dialog holds the file, the progress, and the report, so the
// page behind it never changes shape while a reader uploads (PR-18).
describe("the upload dialog", () => {
  const file = () => new File(["Name,Set code\nBolt,LEA\n"], "export.csv", { type: "text/csv" });

  it("shows a progress bar while the upload runs, and takes it away at the end", async () => {
    let release: (v: unknown) => void = () => {};
    importCollection.mockImplementation(
      () =>
        new Promise((resolve) => {
          release = resolve;
        }),
    );
    const user = userEvent.setup();
    await renderAt("/collection");
    await user.upload(await openUpload(user), file());
    await user.click(screen.getByRole("button", { name: "Upload" }));

    const bar = await screen.findByRole("progressbar");
    expect(bar).toHaveAccessibleName("Uploading and resolving cards...");
    expect(screen.getByRole("button", { name: "Working..." })).toBeDisabled();

    release({ collection: { id: "c-new", name: "export.csv", cardCount: 1 }, report: { unresolved: [], resolvedCount: 1, unresolvedByReason: {} } });
    await screen.findByTestId("card-count");
    expect(screen.queryByTestId("upload-progress")).not.toBeInTheDocument();
  });

  it("keeps the report on the page after the dialog closes", async () => {
    importCollection.mockResolvedValue({
      collection: { id: "c-new", name: "export.csv", cardCount: 1 },
      report: { unresolved: [{ line: 4, raw: "x", reason: UnresolvedReason.UNKNOWN_CARD }], resolvedCount: 1, unresolvedByReason: { UNRESOLVED_REASON_UNKNOWN_CARD: 1 } },
    });
    const user = userEvent.setup();
    await renderAt("/collection");
    await user.upload(await openUpload(user), file());
    await user.click(screen.getByRole("button", { name: "Upload" }));
    await screen.findByTestId("card-count");

    await user.click(screen.getByRole("button", { name: "Done" }));
    // One report at a time: the dialog holds it, then the page does.
    await waitFor(() => expect(screen.getAllByTestId("card-count")).toHaveLength(1));
    expect(screen.getByTestId("card-count")).toHaveTextContent("export.csv: 1 cards, 1 rows resolved, 1 unresolved.");
    expect(screen.getByRole("heading", { name: "Import result" })).toBeInTheDocument();
  });

  it("the open dialog raises no accessibility violation", async () => {
    const { container } = await renderAt("/collection");
    await screen.findByRole("button", { name: "binder-july.csv" });
    await openUpload(userEvent.setup());
    expect(await axe(container.ownerDocument.body)).toHaveNoViolations();
  });
});
