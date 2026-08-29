import { Color } from "@mtg/api-client/mtg/v1/card_pb";
import { CardRole, type Deck, Severity } from "@mtg/api-client/mtg/v1/deck_pb";
import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import { QueryClientProvider } from "@tanstack/react-query";
import { render, screen, within } from "@testing-library/react";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { makeQueryClient } from "../../lib/query-client";
import { DeckView } from "./deck-view";

const getCards = vi.fn();
vi.mock("../../lib/api", () => ({
  cardClient: { getCards: (...args: unknown[]) => getCards(...args) },
  deckClient: { exportDeck: vi.fn() },
}));

const img = (n: string) => ({ small: `https://cards.scryfall.io/small/${n}.jpg`, normal: `https://cards.scryfall.io/normal/${n}.jpg` });

const cards = [
  {
    oracleId: "o-forest",
    name: "Forest",
    typeLine: "Basic Land — Forest",
    cardTypes: ["Land"],
    manaValue: 0,
    producedMana: [Color.G],
    oracleText: "({T}: Add {G}.)",
    faces: [],
    defaultPrinting: { artist: "John Avon", imageUris: img("forest") },
  },
  {
    oracleId: "o-elf",
    name: "Llanowar Elves",
    typeLine: "Creature — Elf Druid",
    cardTypes: ["Creature"],
    colors: [Color.G],
    manaValue: 1,
    manaCost: "{G}",
    producedMana: [Color.G],
    oracleText: "{T}: Add {G}.",
    faces: [],
    defaultPrinting: { artist: "Anson Maddocks", imageUris: img("elf") },
  },
  {
    oracleId: "o-dfc",
    name: "Delver of Secrets // Insectile Aberration",
    typeLine: "Creature — Human Wizard // Creature — Human Insect",
    cardTypes: ["Creature"],
    colors: [Color.U],
    manaValue: 1,
    producedMana: [],
    faces: [
      { name: "Delver of Secrets", artist: "Nils Hamm", imageUris: img("delver-a"), oracleText: "At the beginning of your upkeep, look at the top card of your library.", typeLine: "Creature — Human Wizard", manaCost: "{U}" },
      { name: "Insectile Aberration", artist: "Nils Hamm", imageUris: img("delver-b"), oracleText: "Flying", typeLine: "Creature — Human Insect", manaCost: "" },
    ],
  },
];

const deck = {
  id: "d1",
  name: "Elf test",
  format: { id: FormatId.MODERN, houseRules: "" },
  power: { level: { case: "sixtyStep", value: 1 } },
  summary: "A small test deck.",
  legalityAsOf: "2026-08-24",
  buyCostUsd: 0.5,
  commanderOracleIds: [],
  sideboard: [],
  upgrades: [],
  cards: [
    { oracleId: "o-forest", name: "Forest", count: 20, role: CardRole.LAND, owned: true, ownedCount: 40, priceUsd: 0 },
    { oracleId: "o-elf", name: "Llanowar Elves", count: 4, role: CardRole.RAMP, owned: false, ownedCount: 0, priceUsd: 0.5, reason: "Turn-one mana." },
    { oracleId: "o-dfc", name: "Delver of Secrets // Insectile Aberration", count: 2, role: CardRole.THREAT, owned: true, ownedCount: 2, priceUsd: 0 },
    { oracleId: "o-gone", name: "Missing Card", count: 1, role: CardRole.OTHER, owned: false, ownedCount: 0, priceUsd: 0 },
  ],
  validation: {
    passed: false,
    legalityAsOf: "2026-08-24",
    findings: [
      { code: "deck_size", severity: Severity.BLOCK, message: "27 cards, the format needs 60", oracleId: "" },
      { code: "off_color", severity: Severity.WARN, message: "a blue card in a green deck", oracleId: "o-dfc" },
    ],
  },
} as unknown as Deck;

function renderDeck(d: Deck = deck) {
  return render(
    <QueryClientProvider client={makeQueryClient()}>
      <DeckView deck={d} />
    </QueryClientProvider>,
  );
}

beforeEach(() => {
  getCards.mockReset();
  getCards.mockResolvedValue({ cards, missingOracleIds: ["o-gone"] });
});

describe("DeckView", () => {
  it("loads the cards in one GetCards call and groups them by role", async () => {
    renderDeck();
    expect(await screen.findByRole("region", { name: "Lands (20)" })).toBeInTheDocument();
    await screen.findByAltText("Forest (card)");
    expect(getCards).toHaveBeenCalledTimes(1);
    expect((getCards.mock.calls[0][0] as { oracleIds: string[] }).oracleIds).toEqual(["o-forest", "o-elf", "o-dfc", "o-gone"]);
    expect(screen.getByRole("region", { name: "Ramp (4)" })).toBeInTheDocument();
    expect(screen.getByRole("region", { name: "Threats (2)" })).toBeInTheDocument();
    expect(screen.getByRole("region", { name: "Other (1)" })).toHaveTextContent("No card data for this entry.");
    expect(screen.getByText(/not in the card database/)).toHaveTextContent("1 card of this deck are not in the card database: Missing Card.");
  });

  it("gives every image the full-width class and the 488 by 680 size, and no artist caption (D-6, D-291)", async () => {
    renderDeck();
    await screen.findByAltText("Forest (card)");
    const images = screen.getAllByRole("img");
    expect(images).toHaveLength(4);
    for (const image of images) {
      expect(image).toHaveClass("h-auto", "w-full");
      expect(image).toHaveAttribute("width", "488");
      expect(image).toHaveAttribute("height", "680");
    }
    expect(screen.queryByText(/Illustrated by/)).not.toBeInTheDocument();
    // The one Fan Content notice of the view stays.
    expect(screen.getByText(/unofficial Fan Content/)).toBeInTheDocument();
  });

  it("shows both faces of a double-faced card (F-9)", async () => {
    renderDeck();
    await screen.findByAltText("Forest (card)");
    expect(screen.getByAltText("Delver of Secrets (card)")).toHaveAttribute("src", "https://cards.scryfall.io/normal/delver-a.jpg");
    expect(screen.getByAltText("Insectile Aberration (card)")).toHaveAttribute("src", "https://cards.scryfall.io/normal/delver-b.jpg");
    expect(screen.getByText("Delver of Secrets (face 1 of 2)")).toBeInTheDocument();
  });

  it("marks owned cards and prices the rest", async () => {
    renderDeck();
    await screen.findByAltText("Forest (card)");
    const lands = screen.getByRole("region", { name: "Lands (20)" });
    expect(within(lands).getByTestId("owned-mark")).toHaveTextContent("Owned (40)");
    const ramp = screen.getByRole("region", { name: "Ramp (4)" });
    expect(within(ramp).getByTestId("buy-mark")).toHaveTextContent("To buy: $0.50");
    expect(within(ramp).getByText("Turn-one mana.")).toBeInTheDocument();
    expect(screen.getByTestId("buy-cost")).toHaveTextContent("To buy: $0.50");
  });

  it("shows the findings, the legality date, the curve, and the sources", async () => {
    renderDeck();
    await screen.findByAltText("Forest (card)");
    expect(screen.getByTestId("legality-line")).toHaveTextContent("Not legal, checked against the card data of 2026-08-24.");
    expect(screen.getByTestId("buy-cost")).toHaveTextContent("To buy: $0.50");
    const findings = screen.getByRole("region", { name: "Findings" });
    expect(within(findings).getAllByRole("listitem")).toHaveLength(2);
    expect(findings).toHaveTextContent("Block (deck_size): 27 cards, the format needs 60");
    expect(findings).toHaveTextContent("Warning (off_color): a blue card in a green deck — Delver of Secrets // Insectile Aberration");
    expect(screen.getByText("Modern · Casual · 27 cards")).toBeInTheDocument();

    const curve = screen.getByRole("table", { name: /Mana curve, lands excluded/ });
    const one = within(curve).getByRole("row", { name: /^1 / });
    expect(one).toHaveTextContent("6");
    // The deck's cards are green and blue, so those two rows show and the rest do not.
    const sources = screen.getByRole("table", { name: /Mana sources/ });
    expect(within(sources).getByRole("row", { name: /Green/ })).toHaveTextContent("24");
    expect(within(sources).getByRole("row", { name: /Blue/ })).toHaveTextContent("0");
    expect(within(sources).queryByRole("row", { name: /White/ })).not.toBeInTheDocument();
    expect(within(sources).queryByRole("row", { name: /Colorless/ })).not.toBeInTheDocument();
  });

  it("falls back to the small image when the normal one fails", async () => {
    renderDeck();
    const image = await screen.findByAltText("Forest (card)");
    image.dispatchEvent(new Event("error"));
    expect(await screen.findByAltText("Forest (card)")).toHaveAttribute("src", "https://cards.scryfall.io/small/forest.jpg");
  });

  it("shows the commander from commander_oracle_ids, which the card list does not hold (D-289)", async () => {
    getCards.mockResolvedValue({ cards, missingOracleIds: [] });
    renderDeck({
      ...deck,
      format: { id: FormatId.COMMANDER, houseRules: "" },
      power: { level: { case: "bracket", value: 2 } },
      commanderOracleIds: ["o-elf"],
      cards: deck.cards.slice(0, 1),
    } as unknown as Deck);
    const commander = await screen.findByRole("region", { name: "Commander (1)" });
    expect(within(commander).getByTestId("commander-mark")).toBeInTheDocument();
    expect(await within(commander).findByAltText("Llanowar Elves (card)")).toBeInTheDocument();
    expect(within(commander).queryByTestId("owned-mark")).not.toBeInTheDocument();
    expect(within(commander).queryByTestId("buy-mark")).not.toBeInTheDocument();
    expect((getCards.mock.calls[0][0] as { oracleIds: string[] }).oracleIds).toContain("o-elf");
    expect(screen.queryByRole("region", { name: /^Ramp/ })).not.toBeInTheDocument();
    expect(screen.getByText("Commander · Bracket 2 · 20 cards + 1 commander")).toBeInTheDocument();
  });

  it("counts the main deck without a commander that sits in cards (D-289)", async () => {
    getCards.mockResolvedValue({ cards, missingOracleIds: [] });
    renderDeck({
      ...deck,
      format: { id: FormatId.COMMANDER, houseRules: "" },
      power: undefined,
      commanderOracleIds: ["o-elf"],
      cards: deck.cards.slice(0, 2),
    } as unknown as Deck);
    await screen.findByRole("region", { name: "Commander (1)" });
    expect(screen.getByText("Commander · 20 cards + 1 commander")).toBeInTheDocument();
    expect(screen.queryByRole("region", { name: /^Ramp/ })).not.toBeInTheDocument();
  });

  it("keeps the last cards on screen while a new deck loads, with no stale missing list", async () => {
    getCards.mockResolvedValue({ cards, missingOracleIds: ["o-gone"] });
    const client = makeQueryClient();
    const view = render(
      <QueryClientProvider client={client}>
        <DeckView deck={deck} />
      </QueryClientProvider>,
    );
    await screen.findByAltText("Forest (card)");
    expect(screen.getByText(/not in the card database/)).toBeInTheDocument();
    getCards.mockReturnValue(new Promise(() => {}));
    view.rerender(
      <QueryClientProvider client={client}>
        <DeckView deck={{ ...deck, id: "d2", cards: [deck.cards[0], { ...deck.cards[1], oracleId: "o-new", name: "New Card" }] } as unknown as Deck} />
      </QueryClientProvider>,
    );
    expect(await screen.findByRole("status")).toHaveTextContent("Refreshing card data...");
    expect(screen.getByAltText("Forest (card)")).toBeInTheDocument();
    expect(screen.queryByText(/not in the card database/)).not.toBeInTheDocument();
  });

  it("shows the revision note and the diff against the base deck (PR-12B)", async () => {
    getCards.mockResolvedValue({ cards, missingOracleIds: [] });
    const base = { ...deck, id: "d0", cards: deck.cards.slice(0, 2) } as unknown as Deck;
    const revised = {
      ...deck,
      id: "d1",
      revisedFromDeckId: "d0",
      revisionNote: "I removed 4 Llanowar Elves. I changed the count of Forest: 20 to 22.",
      cards: [{ ...deck.cards[0], count: 22 }, deck.cards[2]],
    } as unknown as Deck;
    render(
      <QueryClientProvider client={makeQueryClient()}>
        <DeckView deck={revised} base={base} />
      </QueryClientProvider>,
    );
    await screen.findByAltText("Forest (card)");
    expect(screen.getByTestId("revision-note")).toHaveTextContent("I removed 4 Llanowar Elves.");
    const diff = screen.getByTestId("revision-diff");
    expect(within(diff).getAllByRole("listitem").map((li) => li.textContent)).toEqual([
      "Removed 4 Llanowar Elves",
      "Added 2 Delver of Secrets // Insectile Aberration",
      "Count of Forest: 20 to 22",
    ]);
  });

  it("shows no diff when the base is not the deck this one revised", async () => {
    getCards.mockResolvedValue({ cards, missingOracleIds: [] });
    const revised = { ...deck, id: "d1", revisedFromDeckId: "d0", revisionNote: "note" } as unknown as Deck;
    render(
      <QueryClientProvider client={makeQueryClient()}>
        <DeckView deck={revised} base={{ ...deck, id: "d9" } as unknown as Deck} />
      </QueryClientProvider>,
    );
    await screen.findByAltText("Forest (card)");
    expect(screen.getByTestId("revision-note")).toHaveTextContent("note");
    expect(screen.queryByTestId("revision-diff")).not.toBeInTheDocument();
  });

  it("drops the not_owned warnings from the findings and keeps a not_owned block (D-300)", async () => {
    getCards.mockResolvedValue({ cards, missingOracleIds: [] });
    renderDeck({
      ...deck,
      validation: {
        passed: false,
        legalityAsOf: "2026-08-24",
        findings: [
          { code: "not_owned", severity: Severity.WARN, message: "Llanowar Elves: the deck needs 4, the collection has 0", oracleId: "o-elf" },
          { code: "not_owned", severity: Severity.BLOCK, message: "Forest: the deck needs 20, the collection has 0", oracleId: "o-forest" },
          { code: "curve_summary", severity: Severity.INFO, message: "average mana value 2.1", oracleId: "" },
        ],
      },
    } as unknown as Deck);
    await screen.findByAltText("Forest (card)");
    const findings = screen.getByRole("region", { name: "Findings" });
    const items = within(findings).getAllByRole("listitem").map((li) => li.textContent);
    expect(items).toHaveLength(2);
    expect(items[0]).toContain("Block (not_owned): Forest");
    expect(items[1]).toContain("average mana value 2.1");
  });

  it("shows the owned printing's image when the deck carries one (D-299)", async () => {
    renderDeck({
      ...deck,
      cards: [{ ...deck.cards[0], ownedPrinting: { scryfallId: "p-old", artist: "Rob Alexander", imageUris: img("forest-alpha"), priceUsd: 40 } }],
    } as unknown as Deck);
    expect(await screen.findByAltText("Forest (card)")).toHaveAttribute("src", "https://cards.scryfall.io/normal/forest-alpha.jpg");
  });

  it("reports a GetCards failure", async () => {
    getCards.mockRejectedValue(new Error("[unavailable] card database not loaded yet"));
    renderDeck();
    expect(await screen.findByRole("alert")).toHaveTextContent("Could not load the card data: [unavailable] card database not loaded yet");
  });

  it("has no axe violations", async () => {
    const { container } = renderDeck();
    await screen.findByAltText("Forest (card)");
    expect(await axe(container)).toHaveNoViolations();
  });
});
