import { screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";

import { renderAt } from "../../test-utils";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const check = vi.fn();
vi.mock("../../lib/api", () => ({ healthClient: { check: (...args: unknown[]) => check(...args) } }));

describe("HealthFooter", () => {
  it("shows the error and says the card data is unknown when the API fails", async () => {
    check.mockRejectedValue(new Error("connect ECONNREFUSED"));
    renderAt("/sign-in");
    expect(await screen.findByText(/ECONNREFUSED/)).toBeInTheDocument();
    expect(screen.getByTestId("freshness")).toHaveTextContent("Card data: unknown");
  });

  it("says the card data is not loaded when the snapshot is none", async () => {
    check.mockResolvedValue({ status: "ok", version: "test", cardSnapshot: "none", cardSnapshotAgeHours: -1 });
    renderAt("/sign-in");
    expect(await screen.findByText("API: ok, version test")).toBeInTheDocument();
    expect(screen.getByTestId("freshness")).toHaveTextContent("Card data: not loaded yet.");
  });
});
