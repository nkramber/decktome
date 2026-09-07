import { type Card, LegalityStatus } from "@mtg/api-client/mtg/v1/card_pb";
import { CardRole, type DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import { FeedbackKind, FeedbackVerdict } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { makeQueryClient } from "../../lib/query-client";
import { CardDetail, legalityLabel, scryfallURL } from "./card-detail";

const getRulings = vi.fn();
const getPrintings = vi.fn();
const submitFeedback = vi.fn();
vi.mock("../../lib/api", () => ({
  cardClient: { getRulings: (...args: unknown[]) => getRulings(...args), getPrintings: (...args: unknown[]) => getPrintings(...args) },
  feedbackClient: { submitFeedback: (...args: unknown[]) => submitFeedback(...args) },
}));

const img = (n: string) => ({ small: `https://cards.scryfall.io/small/${n}.jpg`, normal: `https://cards.scryfall.io/normal/${n}.jpg` });

const solRing = {
  oracleId: "o-sol",
  name: "Sol Ring",
  typeLine: "Artifact",
  manaCost: "{1}",
  manaValue: 1,
  cardTypes: ["Artifact"],
  colors: [],
  oracleText: "{T}: Add {C}{C}.",
  faces: [],
  legalities: { commander: LegalityStatus.LEGAL, modern: LegalityStatus.NOT_LEGAL, vintage: LegalityStatus.RESTRICTED, legacy: LegalityStatus.BANNED },
  defaultPrinting: { scryfallId: "p-c21", setCode: "c21", setName: "Commander 2021", collectorNumber: "263", artist: "Mike Bierek", imageUris: img("sol") },
} as unknown as Card;

const entry = { oracleId: "o-sol", name: "Sol Ring", count: 1, role: CardRole.RAMP, owned: true, ownedCount: 1, priceUsd: 0, reason: "Turn-one ramp." } as DeckCard;

function renderDetail(props: Partial<Parameters<typeof CardDetail>[0]> = {}) {
  const onOpenChange = vi.fn();
  const view = render(
    <QueryClientProvider client={makeQueryClient()}>
      <CardDetail entry={entry} card={solRing} open onOpenChange={onOpenChange} {...props} />
    </QueryClientProvider>,
  );
  return { ...view, onOpenChange };
}

beforeEach(() => {
  getRulings.mockReset();
  getPrintings.mockReset();
  getRulings.mockResolvedValue({
    rulings: [
      { publishedAt: "2004-10-04", comment: "Sol Ring is an artifact.", source: "wotc" },
      { publishedAt: "2020-11-10", comment: "It taps for two.", source: "scryfall" },
    ],
    asOf: "2026-09-03",
    hasRulings: true,
  });
  getPrintings.mockResolvedValue({
    printings: [
      { scryfallId: "p-msc", setCode: "msc", setName: "Marvel Super Heroes Commander", collectorNumber: "9", rarity: "uncommon", priceUsd: 2.5, digital: false },
      { scryfallId: "p-c21", setCode: "c21", setName: "Commander 2021", collectorNumber: "263", rarity: "uncommon", priceUsd: 1.5, digital: false },
      { scryfallId: "p-mtga", setCode: "mtga", setName: "Arena", collectorNumber: "1", rarity: "uncommon", priceUsd: 0, digital: true },
    ],
    priceAsOf: "2026-09-03",
  });
});

describe("CardDetail", () => {
  it("carries the thumbs of the card when it knows the deck, and none on a page with no deck (PR-27, D-559)", async () => {
    const user = userEvent.setup();
    submitFeedback.mockReset();
    submitFeedback.mockResolvedValue({ feedbackId: "fb1" });
    const first = renderDetail({ deckId: "d1" });
    const dialog = await screen.findByRole("dialog", { name: "Sol Ring" });
    const thumbs = within(dialog).getByRole("group", { name: "Rate this card" });
    await user.click(within(thumbs).getByRole("button", { name: "This helped" }));
    await waitFor(() => expect(submitFeedback).toHaveBeenCalledTimes(1));
    expect(submitFeedback.mock.calls[0][0]).toEqual({ feedback: { kind: FeedbackKind.CARD, verdict: FeedbackVerdict.UP, reasons: [], text: "", deckId: "d1", oracleId: "o-sol" } });
    first.unmount();
    renderDetail();
    const plain = await screen.findByRole("dialog", { name: "Sol Ring" });
    expect(within(plain).queryByRole("group", { name: "Rate this card" })).not.toBeInTheDocument();
  });

  it("shows the card, the reason line, the legalities, and the Scryfall link", async () => {
    renderDetail();
    const dialog = await screen.findByRole("dialog", { name: "Sol Ring" });
    // The type line names the description of the sheet and the face.
    expect(within(dialog).getAllByText("Artifact")).toHaveLength(2);
    expect(within(dialog).getByText("{T}: Add {C}{C}.")).toBeInTheDocument();
    expect(within(dialog).getByText("{1}")).toBeInTheDocument();
    expect(within(dialog).getByTestId("detail-reason")).toHaveTextContent("In this deck: Turn-one ramp.");
    expect(within(dialog).getByAltText("Sol Ring (card)")).toHaveAttribute("src", "https://cards.scryfall.io/normal/sol.jpg");
    const legalities = within(dialog).getByRole("table", { name: "Legalities" });
    expect(within(legalities).getByRole("row", { name: /Commander/ })).toHaveTextContent("Legal");
    expect(within(legalities).getByRole("row", { name: /Modern/ })).toHaveTextContent("Not legal");
    expect(within(legalities).getByRole("row", { name: /Vintage/ })).toHaveTextContent("Restricted");
    expect(within(legalities).getByRole("row", { name: /Legacy/ })).toHaveTextContent("Banned");
    expect(within(legalities).queryByRole("row", { name: /Pauper/ })).not.toBeInTheDocument();
    expect(within(dialog).getByRole("link", { name: "Open on Scryfall" })).toHaveAttribute("href", "https://scryfall.com/card/c21/263");
  });

  it("shows the rulings with their dates and the snapshot date (PR-20 gate)", async () => {
    renderDetail();
    const rulings = await screen.findByTestId("rulings");
    const items = within(rulings).getAllByRole("listitem");
    expect(items).toHaveLength(2);
    expect(items[0]).toHaveTextContent("2004-10-04");
    expect(items[0]).toHaveTextContent("Sol Ring is an artifact.");
    expect(items[0].querySelector("time")).toHaveAttribute("datetime", "2004-10-04");
    expect(screen.getByTestId("rulings-as-of")).toHaveTextContent("Rulings from the card data of 2026-09-03.");
    expect(getRulings).toHaveBeenCalledWith({ oracleId: "o-sol" });
  });

  it("shows every printing with its price, and marks a digital one", async () => {
    renderDetail();
    const rows = await screen.findAllByTestId("printing-row");
    expect(rows).toHaveLength(3);
    expect(rows[0]).toHaveTextContent("Marvel Super Heroes Commander");
    expect(rows[0]).toHaveTextContent("$2.50");
    expect(rows[2]).toHaveTextContent("(digital)");
    expect(rows[2]).toHaveTextContent("no price");
    expect(screen.getByText("Prices of 2026-09-03.")).toBeInTheDocument();
  });

  it("says when the card data holds no rulings file, and when a card has none", async () => {
    getRulings.mockResolvedValue({ rulings: [], asOf: "2026-08-24", hasRulings: false });
    renderDetail();
    expect(await screen.findByText("The card data holds no rulings file yet.")).toBeInTheDocument();
    getRulings.mockResolvedValue({ rulings: [], asOf: "2026-09-03", hasRulings: true });
    renderDetail({ entry: { ...entry, oracleId: "o-other" } as DeckCard });
    expect(await screen.findByText("No rulings.")).toBeInTheDocument();
  });

  it("reports a failed rulings call and loads nothing while closed", async () => {
    getRulings.mockRejectedValue(new Error("[unavailable] card database not loaded yet"));
    const first = renderDetail();
    expect(await screen.findByRole("alert")).toHaveTextContent("Could not load the rulings: [unavailable] card database not loaded yet");
    first.unmount();
    getRulings.mockClear();
    renderDetail({ open: false });
    expect(getRulings).not.toHaveBeenCalled();
    expect(screen.queryByRole("dialog", { name: "Sol Ring" })).not.toBeInTheDocument();
  });

  it("labels the legality statuses and builds the Scryfall link", () => {
    expect(legalityLabel(LegalityStatus.LEGAL)).toBe("Legal");
    expect(legalityLabel(undefined)).toBe("Unknown");
    expect(scryfallURL(undefined)).toBe("");
    expect(scryfallURL({ setCode: "msc", collectorNumber: "5" } as never)).toBe("https://scryfall.com/card/msc/5");
  });

  it("has no axe violations with the sheet open", async () => {
    const { container } = renderDetail();
    await screen.findByTestId("rulings");
    // The sheet renders in a portal outside the container.
    expect(await axe(container.ownerDocument.body)).toHaveNoViolations();
  });
});

