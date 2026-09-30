import { type Deck, RerunCase } from "@mtg/api-client/mtg/v1/deck_pb";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { axe } from "jest-axe";
import { describe, expect, it, vi } from "vitest";
import { rerunLabel, StaleBanner, staleNames } from "./stale-banner";

const deck = {
  id: "d1",
  stale: true,
  staleOracleIds: ["o-welcome", "o-karlov"],
  staleReason: "The rerun replaces Ajani's Welcome alone, and keeps the rest of the deck.",
  rerunCase: RerunCase.PATCH,
  commanderOracleIds: ["o-karlov"],
  commanders: [{ oracleId: "o-karlov", name: "Karlov of the Ghost Council" }],
  cards: [
    { oracleId: "o-plains", name: "Plains" },
    { oracleId: "o-welcome", name: "Ajani's Welcome" },
  ],
  sideboard: [],
} as unknown as Deck;

describe("StaleBanner", () => {
  it("names the cards, the reason, and the rerun (I-1)", async () => {
    const onRerun = vi.fn();
    const { container } = render(<StaleBanner deck={deck} onRerun={onRerun} />);
    expect(screen.getByRole("alert")).toHaveTextContent("A rule change made this deck illegal");
    expect(screen.getByText("No longer legal: Karlov of the Ghost Council, Ajani's Welcome.")).toBeInTheDocument();
    expect(screen.getByText(deck.staleReason)).toBeInTheDocument();
    await userEvent.click(screen.getByRole("button", { name: "Replace the banned cards" }));
    expect(onRerun).toHaveBeenCalledOnce();
    expect(await axe(container)).toHaveNoViolations();
  });

  it("shows nothing on a legal deck", () => {
    const { container } = render(<StaleBanner deck={{ ...deck, stale: false } as Deck} onRerun={() => {}} />);
    expect(container).toBeEmptyDOMElement();
  });

  it("has no button without a chat, and none while a turn runs", () => {
    const { rerender } = render(<StaleBanner deck={deck} />);
    expect(screen.queryByRole("button")).toBeNull();
    rerender(<StaleBanner deck={deck} busy onRerun={() => {}} />);
    expect(screen.getByRole("button")).toBeDisabled();
  });
});

describe("rerunLabel", () => {
  it("names the action of each case (D-1008, D-1021)", () => {
    expect(rerunLabel(deck)).toBe("Replace the banned cards");
    expect(rerunLabel({ ...deck, rerunCase: RerunCase.REBUILD } as Deck)).toBe("Pick a new commander");
    expect(rerunLabel({ ...deck, rerunCase: RerunCase.REBUILD, staleOracleIds: ["o-welcome"] } as Deck)).toBe("Rebuild the deck");
  });

  it("reads names in the order of the deck, the commander first", () => {
    expect(staleNames(deck)).toEqual(["Karlov of the Ghost Council", "Ajani's Welcome"]);
  });
});
