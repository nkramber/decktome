import { render, screen } from "@testing-library/react";
import { axe } from "jest-axe";
import { describe, expect, it, vi } from "vitest";

import { App } from "./App";

vi.mock("./api", () => ({
  healthClient: {
    check: vi.fn().mockResolvedValue({
      status: "ok",
      version: "test",
      cardSnapshot: "2026-08-24T09:01:52Z",
      cardSnapshotAgeHours: 2.5,
    }),
  },
}));

describe("App", () => {
  it("renders the heading and the health result", async () => {
    render(<App />);
    expect(screen.getByRole("heading", { level: 1, name: "MtG Deck Builder" })).toBeInTheDocument();
    expect(await screen.findByText(/"status": "ok"/)).toBeInTheDocument();
    expect(screen.getByTestId("freshness")).toHaveTextContent("Card data as of 2026-08-24T09:01:52Z (2.5 h old)");
  });

  it("has no axe violations", async () => {
    const { container } = render(<App />);
    await screen.findByText(/"status": "ok"/);
    expect(await axe(container)).toHaveNoViolations();
  });
});
