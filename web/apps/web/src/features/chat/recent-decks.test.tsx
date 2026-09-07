import { render, screen, waitFor } from "@testing-library/react";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { RecentDecks } from "./recent-decks";
import { useAppStore } from "../../lib/store";

vi.mock("firebase/app");
vi.mock("firebase/auth");
vi.mock("../deck/use-decks", () => ({
  useCommanderCards: () => ({}),
  useDeckWrites: () => ({ setFavorite: { mutateAsync: vi.fn() } }),
}));

beforeEach(() => {
  localStorage.clear();
  useAppStore.setState({ hadDecks: false });
});

// F-67: a reader with no deck saw three boxes appear and go. The
// placeholder belongs on the page only when a deck follows it.
describe("the placeholder of the newest decks", () => {
  it("shows nothing while the list loads for a reader with no deck", () => {
    const { container } = render(<RecentDecks decks={[]} isPending={true} />);
    expect(screen.queryByRole("status")).not.toBeInTheDocument();
    expect(container).toBeEmptyDOMElement();
  });

  it("shows the placeholder while the list loads for a reader who owns a deck", async () => {
    useAppStore.setState({ hadDecks: true });
    render(<RecentDecks decks={[]} isPending={true} />);
    await waitFor(() => expect(screen.getByRole("status")).toBeInTheDocument());
    expect(screen.getByText("Loading your decks...")).toBeInTheDocument();
  });

  it("shows nothing once the list answers with no deck", () => {
    useAppStore.setState({ hadDecks: true });
    const { container } = render(<RecentDecks decks={[]} isPending={false} />);
    expect(container).toBeEmptyDOMElement();
  });
});
