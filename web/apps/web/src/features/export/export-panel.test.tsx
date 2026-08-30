import type { Card } from "@mtg/api-client/mtg/v1/card_pb";
import { ExportFormat } from "@mtg/api-client/mtg/v1/deck_service_pb";
import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { render, screen, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { ExportPanel } from "./export-panel";

const exportDeck = vi.fn();
vi.mock("../../lib/api", () => ({
  deckClient: { exportDeck: (...args: unknown[]) => exportDeck(...args) },
}));

const byId = new Map<string, Card>([["o-sw", { oracleId: "o-sw", name: "Soul Warden", defaultPrinting: { setCode: "mm3", collectorNumber: "24" } } as Card]]);
const deck = {
  id: "d1",
  commanderOracleIds: [],
  cards: [
    { oracleId: "o-sw", name: "Soul Warden", count: 4, owned: false, ownedCount: 1, priceUsd: 0.5 },
    { oracleId: "o-pl", name: "Plains", count: 20, owned: true, ownedCount: 40, priceUsd: 0 },
  ],
  sideboard: [],
  upgrades: [{ oracleId: "o-up", name: "Rhystic Study", count: 1, owned: false, ownedCount: 0, priceUsd: 40 }],
} as unknown as Deck;

beforeEach(() => {
  exportDeck.mockReset();
  exportDeck.mockResolvedValue({ text: "Deck\n3 Soul Warden (MM3) 24\n", fileName: "deck.txt" });
  Object.assign(navigator, { clipboard: { writeText: vi.fn().mockResolvedValue(undefined) } });
});

describe("ExportPanel", () => {
  it("shows the buy list with a Scryfall link per card and the upgrades apart (D-308)", async () => {
    const view = render(<ExportPanel deck={deck} byId={byId} />);
    const list = screen.getByRole("region", { name: "Buy list (3)" });
    // The list is long, so it stays shut until it is asked for. jsdom
    // keeps the rows reachable inside a closed details, and a browser
    // does not, so the check reads the element and not the rows.
    expect(list.querySelector("details")).not.toHaveAttribute("open");
    expect(list).toHaveTextContent("3 × Soul Warden · $1.50");
    expect(within(list).getByRole("link", { name: "Scryfall" })).toHaveAttribute("href", "https://scryfall.com/card/mm3/24");
    expect(screen.getByRole("region", { name: "Upgrades (1)" })).toHaveTextContent("1 × Rhystic Study · $40.00");
    expect(await axe(view.container)).toHaveNoViolations();
  });

  it("says when nothing is to buy", () => {
    render(<ExportPanel deck={{ ...deck, cards: [deck.cards[1]], upgrades: [] } as Deck} byId={byId} />);
    expect(screen.getByRole("region", { name: "Buy list (0)" })).toHaveTextContent("Nothing to buy: every card is owned.");
    expect(screen.queryByRole("region", { name: /Upgrades/ })).toBeNull();
  });

  it("copies the Arena text from ExportDeck (D-15)", async () => {
    render(<ExportPanel deck={deck} byId={byId} />);
    await userEvent.click(screen.getByRole("button", { name: "Copy Arena text" }));
    expect(exportDeck).toHaveBeenCalledWith({ deckId: "d1", format: ExportFormat.ARENA_TEXT });
    expect(navigator.clipboard.writeText).toHaveBeenCalledWith("Deck\n3 Soul Warden (MM3) 24\n");
    expect(await screen.findByRole("status")).toHaveTextContent("Copied the Arena text: 2 lines.");
  });

  it("downloads the buy list under the file name the API gives (D-309)", async () => {
    exportDeck.mockResolvedValue({ text: "3 Soul Warden\n", fileName: "deck-buy-list.txt" });
    const create = vi.fn(() => "blob:x");
    vi.stubGlobal("URL", { ...URL, createObjectURL: create, revokeObjectURL: vi.fn() });
    const click = vi.spyOn(HTMLAnchorElement.prototype, "click").mockImplementation(() => {});
    render(<ExportPanel deck={deck} byId={byId} />);
    await userEvent.click(screen.getByRole("button", { name: "Download buy list" }));
    expect(exportDeck).toHaveBeenCalledWith({ deckId: "d1", format: ExportFormat.BUY_LIST_TEXT });
    expect(await screen.findByRole("status")).toHaveTextContent("Saved deck-buy-list.txt: 1 lines.");
    expect(click).toHaveBeenCalledTimes(1);
    vi.unstubAllGlobals();
    click.mockRestore();
  });

  it("shows an export failure and enables the buttons again", async () => {
    exportDeck.mockRejectedValue(new Error("deck not found"));
    render(<ExportPanel deck={deck} byId={byId} />);
    const button = screen.getByRole("button", { name: "Copy Arena text" });
    await userEvent.click(button);
    expect(await screen.findByRole("alert")).toHaveTextContent("Export failed: deck not found");
    expect(button).toBeEnabled();
  });

  it("disables the buttons for a deck with no id", () => {
    render(<ExportPanel deck={{ ...deck, id: "" } as Deck} byId={byId} />);
    expect(screen.getByRole("button", { name: "Copy Arena text" })).toBeDisabled();
  });
});
