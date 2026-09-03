import { Code, ConnectError } from "@connectrpc/connect";
import { CardRole } from "@mtg/api-client/mtg/v1/deck_pb";
import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import { screen, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const getSharedDeck = vi.fn();
const exportSharedDeck = vi.fn();
vi.mock("../../lib/api", () => ({
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: { listCollections: () => Promise.resolve({ collections: [] }) },
  agentClient: { listSessions: () => Promise.resolve({ sessions: [], nextPageToken: "" }) },
  cardClient: { getCards: vi.fn() },
  deckClient: {
    getSharedDeck: (...a: unknown[]) => getSharedDeck(...a),
    exportSharedDeck: (...a: unknown[]) => exportSharedDeck(...a),
    listDecks: () => Promise.resolve({ decks: [], nextPageToken: "" }),
  },
}));

const img = (n: string) => ({ small: `https://cards.scryfall.io/small/${n}.jpg`, normal: `https://cards.scryfall.io/normal/${n}.jpg` });
const token = "a".repeat(43);

const shared = {
  name: "Elf Ball",
  format: { id: FormatId.COMMANDER, houseRules: "" },
  power: { level: { case: "bracket", value: 3 } },
  summary: "Elves that make mana and draw cards.",
  commanderOracleIds: ["o-cmd"],
  legalityAsOf: "2026-09-03",
  cardCount: 3,
  cards: [
    {
      oracleId: "o-cmd", name: "Ezuri, Renegade Leader", count: 1, role: CardRole.UNSPECIFIED, reason: "",
      card: { oracleId: "o-cmd", name: "Ezuri, Renegade Leader", typeLine: "Legendary Creature — Elf Warrior", cardTypes: ["Creature"], colorIdentity: [], colors: [], faces: [], defaultPrinting: { artist: "Eric Deschamps", imageUris: img("ezuri") } },
    },
    {
      oracleId: "o-elf", name: "Llanowar Elves", count: 1, role: CardRole.RAMP, reason: "Turn-one mana.",
      card: { oracleId: "o-elf", name: "Llanowar Elves", typeLine: "Creature — Elf Druid", cardTypes: ["Creature"], colorIdentity: [], colors: [], faces: [], defaultPrinting: { artist: "Anson Maddocks", imageUris: img("elf") } },
    },
    { oracleId: "o-forest", name: "Forest", count: 1, role: CardRole.LAND, reason: "" },
  ],
  sideboard: [],
};

beforeEach(() => {
  // A visitor, signed out.
  state.user = null;
  getSharedDeck.mockReset();
  getSharedDeck.mockResolvedValue({ deck: shared });
  exportSharedDeck.mockReset();
  exportSharedDeck.mockResolvedValue({ text: "Commander\n1 Ezuri, Renegade Leader\n\nDeck\n1 Llanowar Elves\n1 Forest\n", fileName: "elf-ball.txt" });
});

describe("SharedDeckPage", () => {
  it("reads the deck by its token with no sign-in and shows the cards by role", async () => {
    await renderAt(`/d/${token}`);
    expect(await screen.findByRole("heading", { name: "Elf Ball" })).toBeInTheDocument();
    expect(getSharedDeck).toHaveBeenCalledWith({ token });
    expect(screen.getByText("Commander · Bracket 3 · 3 cards")).toBeInTheDocument();
    expect(screen.getByText("Elves that make mana and draw cards.")).toBeInTheDocument();
    const commander = screen.getByRole("region", { name: "Commander (1)" });
    expect(within(commander).getByAltText("Ezuri, Renegade Leader (card)")).toBeInTheDocument();
    const ramp = screen.getByRole("region", { name: "Ramp (1)" });
    expect(within(ramp).getByText("Turn-one mana.")).toBeInTheDocument();
    expect(screen.getByRole("region", { name: "Lands (1)" })).toHaveTextContent("No card data for this entry.");
    // No owned mark, no price, no owner, no chat.
    expect(screen.queryByTestId("owned-mark")).not.toBeInTheDocument();
    expect(screen.queryByTestId("buy-mark")).not.toBeInTheDocument();
    expect(screen.queryByText(/nate@example.com/)).not.toBeInTheDocument();
    expect(screen.queryByRole("navigation")).not.toBeInTheDocument();
  });

  it("exports the deck list through the public call", async () => {
    const user = userEvent.setup();
    const write = vi.fn().mockResolvedValue(undefined);
    Object.defineProperty(navigator, "clipboard", { value: { writeText: write }, configurable: true });
    await renderAt(`/d/${token}`);
    await screen.findByRole("heading", { name: "Elf Ball" });
    await user.click(screen.getByRole("button", { name: "Copy deck list" }));
    expect(await screen.findByRole("status")).toHaveTextContent("Copied the deck list: 5 lines.");
    expect(exportSharedDeck).toHaveBeenCalledWith({ token });
    expect(write).toHaveBeenCalledWith(expect.stringContaining("1 Llanowar Elves"));
  });

  it("says when a link opens nothing", async () => {
    getSharedDeck.mockRejectedValue(new ConnectError("not found", Code.NotFound));
    await renderAt(`/d/${token}`);
    expect(await screen.findByRole("alert")).toHaveTextContent("This link does not open a deck. It was revoked, or it never existed.");
  });

  it("reports another failure as it is", async () => {
    getSharedDeck.mockRejectedValue(new ConnectError("too many calls from this address", Code.ResourceExhausted));
    await renderAt(`/d/${"b".repeat(43)}`);
    expect(await screen.findByRole("alert")).toHaveTextContent(/Could not load the deck: .*too many calls from this address/);
  });

  it("has no axe violations", async () => {
    const { container } = await renderAt(`/d/${token}`);
    await screen.findByRole("heading", { name: "Elf Ball" });
    expect(await axe(container)).toHaveNoViolations();
  });
});
