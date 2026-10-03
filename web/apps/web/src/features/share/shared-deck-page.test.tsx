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
      oracleId: "o-elf", name: "Llanowar Elves", count: 1, role: CardRole.RAMP, reason: "Turn-one mana.", priceUsd: 0.25,
      printing: { artist: "LOTR Artist", imageUris: img("elf-ltr") },
      card: { oracleId: "o-elf", name: "Llanowar Elves", typeLine: "Creature — Elf Druid", cardTypes: ["Creature"], colorIdentity: [], colors: [], faces: [], defaultPrinting: { artist: "Anson Maddocks", imageUris: img("elf") } },
    },
    { oracleId: "o-forest", name: "Forest", count: 1, role: CardRole.LAND, reason: "", priceUsd: 0 },
  ],
  sideboard: [],
  commanders: [],
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
    // No owned mark, no owner, no chat.
    expect(screen.queryByTestId("owned-mark")).not.toBeInTheDocument();
    expect(screen.queryByTestId("buy-mark")).not.toBeInTheDocument();
    expect(screen.queryByText(/nate@example.com/)).not.toBeInTheDocument();
    expect(screen.queryByRole("navigation")).not.toBeInTheDocument();
    // REV-028: the terms of TopDeck.gg ask for a visible credit (D-417).
    expect(screen.getByRole("link", { name: "Tournament data by TopDeck.gg" })).toHaveAttribute("href", "https://topdeck.gg");
  });

  // D-1063: a link can not prove who owns a card, so each tile shows the
  // price of the cheapest printing, and a card with none says so.
  it("shows the price of every card", async () => {
    await renderAt(`/d/${token}`);
    const ramp = await screen.findByRole("region", { name: "Ramp (1)" });
    expect(within(ramp).getByTestId("card-price")).toHaveTextContent("$0.25");
    const lands = screen.getByRole("region", { name: "Lands (1)" });
    expect(within(lands).getByTestId("card-price")).toHaveTextContent("Price unknown");
  });

  // D-1062: the page shows the art of the deck view, the owned printing.
  it("shows the art of the printing the deck view shows", async () => {
    await renderAt(`/d/${token}`);
    const ramp = await screen.findByRole("region", { name: "Ramp (1)" });
    expect(within(ramp).getByAltText("Llanowar Elves (card)")).toHaveAttribute("src", img("elf-ltr").normal);
    const commander = screen.getByRole("region", { name: "Commander (1)" });
    expect(within(commander).getByAltText("Ezuri, Renegade Leader (card)")).toHaveAttribute("src", img("ezuri").normal);
  });

  // D-1064: the four stats and the sample hand of the deck view.
  it("shows the stats and the sample hand", async () => {
    const user = userEvent.setup();
    await renderAt(`/d/${token}`);
    await screen.findByRole("heading", { name: "Elf Ball" });
    for (const caption of ["Mana curve, lands excluded", "Mana sources, cards that make each of the deck's colors", "Card types", "Average mana value, lands excluded"]) {
      expect(screen.getByRole("table", { name: caption })).toBeInTheDocument();
    }
    expect(screen.getByRole("heading", { name: "Sample hand" })).toBeInTheDocument();
    await user.click(screen.getByRole("button", { name: "Draw seven" }));
    // The library holds the main deck alone: the commander stays out.
    expect(screen.getAllByTestId("hand-card")).toHaveLength(2);
  });

  // D-1064: the filters and the sort of the deck view, with no owned filter.
  it("filters and sorts the cards, with no owned filter", async () => {
    const user = userEvent.setup();
    await renderAt(`/d/${token}`);
    await screen.findByRole("heading", { name: "Elf Ball" });
    expect(screen.queryByRole("combobox", { name: "Owned" })).not.toBeInTheDocument();
    await user.selectOptions(screen.getByRole("combobox", { name: "Role" }), String(CardRole.RAMP));
    expect(screen.getByTestId("filter-count")).toHaveTextContent("Showing 1 of 2 cards.");
    expect(screen.queryByRole("region", { name: "Lands (1)" })).not.toBeInTheDocument();
    expect(screen.queryByRole("region", { name: "Commander (1)" })).not.toBeInTheDocument();
    await user.click(screen.getByRole("button", { name: "Clear filters" }));
    await user.selectOptions(screen.getByRole("combobox", { name: "Sort" }), "price");
    const list = screen.getByRole("region", { name: "Cards by price (2)" });
    expect(within(list).getAllByTestId("card-price").map((p) => p.textContent)).toEqual(["$0.25", "Price unknown"]);
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

  // REV-021: a build keeps the commander out of cards since F-124, and the
  // page showed no commander at all.
  it("shows the commander that the server sends in its own list", async () => {
    const [ezuri, ...rest] = shared.cards;
    getSharedDeck.mockResolvedValue({ deck: { ...shared, cards: rest, commanders: [ezuri] } });
    await renderAt(`/d/${token}`);
    const commander = await screen.findByRole("region", { name: "Commander (1)" });
    expect(within(commander).getByAltText("Ezuri, Renegade Leader (card)")).toBeInTheDocument();
  });

  it("links to the Archidekt deck of an import, and to no other link (D-1101)", async () => {
    getSharedDeck.mockResolvedValue({ deck: { ...shared, sourceUrl: "https://archidekt.com/decks/42" } });
    await renderAt(`/d/${token}`);
    expect(await screen.findByRole("link", { name: "Archidekt" })).toHaveAttribute("href", "https://archidekt.com/decks/42");
  });

  it("shows no source line for a link of another site", async () => {
    getSharedDeck.mockResolvedValue({ deck: { ...shared, sourceUrl: "https://evil.example/decks/42" } });
    await renderAt(`/d/${token}`);
    await screen.findByRole("heading", { name: shared.name });
    expect(screen.queryByTestId("import-source")).not.toBeInTheDocument();
  });

  // REV-022: a cold start answers Unavailable until the card index loads,
  // and the page had retry off, so a first visit stayed on the error.
  it("retries a cold start and then shows the deck", async () => {
    getSharedDeck.mockRejectedValueOnce(new ConnectError("card database not loaded yet", Code.Unavailable));
    await renderAt(`/d/${token}`);
    expect(await screen.findByRole("heading", { name: "Elf Ball" }, { timeout: 3000 })).toBeInTheDocument();
    expect(getSharedDeck).toHaveBeenCalledTimes(2);
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
