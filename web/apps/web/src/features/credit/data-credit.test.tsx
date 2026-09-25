import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";

import { ArtCredit, DataCredit } from "./data-credit";

// REV-028: the terms of TopDeck.gg ask for a visible credit with a link
// (D-417), and no page showed one.
describe("DataCredit", () => {
  it("credits Scryfall and links TopDeck.gg", () => {
    render(<DataCredit />);
    expect(screen.getByText(/come from Scryfall/)).toBeInTheDocument();
    const link = screen.getByRole("link", { name: "Tournament data by TopDeck.gg" });
    expect(link).toHaveAttribute("href", "https://topdeck.gg");
  });
});

describe("ArtCredit", () => {
  it("names each artist once, and shows nothing with none", () => {
    const { container, rerender } = render(<ArtCredit artists={["Rebecca Guay", "", "Rebecca Guay", "Terese Nielsen"]} />);
    expect(screen.getByText("Art: Rebecca Guay, Terese Nielsen. ™ & © Wizards of the Coast.")).toBeInTheDocument();
    rerender(<ArtCredit artists={[""]} />);
    expect(container).toBeEmptyDOMElement();
  });
});
