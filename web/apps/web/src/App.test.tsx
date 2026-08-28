import { render, screen } from "@testing-library/react";
import { axe } from "jest-axe";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { App } from "./App";

const check = vi.fn();
vi.mock("./api", () => ({ healthClient: { check: (...args: unknown[]) => check(...args) } }));

beforeEach(() => {
  check.mockReset();
  check.mockResolvedValue({
    status: "ok",
    version: "test",
    cardSnapshot: "2026-08-24T09:01:52Z",
    cardSnapshotAgeHours: 2.5,
  });
});

describe("App", () => {
  it("shows the error and says the card data is unknown when the API fails", async () => {
    check.mockRejectedValue(new Error("connect ECONNREFUSED"));
    render(<App />);
    expect(await screen.findByText(/ECONNREFUSED/)).toBeInTheDocument();
    expect(screen.getByTestId("freshness")).toHaveTextContent("Card data: unknown");
  });

  it("says the card data is not loaded when the snapshot is none", async () => {
    check.mockResolvedValue({ status: "ok", version: "test", cardSnapshot: "none", cardSnapshotAgeHours: -1 });
    render(<App />);
    await screen.findByText(/"status": "ok"/);
    expect(screen.getByTestId("freshness")).toHaveTextContent("Card data: not loaded yet.");
  });

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
