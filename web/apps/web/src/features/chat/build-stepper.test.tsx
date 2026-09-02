import { BuildPhase } from "@mtg/api-client/mtg/v1/agent_service_pb";
import { render, screen } from "@testing-library/react";
import { describe, expect, it } from "vitest";

import { BuildStepper } from "./build-stepper";

// The stepper lights the phase the server streamed (D-435).
describe("the build stepper", () => {
  it("shows nothing between turns", () => {
    render(<BuildStepper phase={BuildPhase.UNSPECIFIED} repaired={false} />);
    expect(screen.queryByTestId("build-stepper")).not.toBeInTheDocument();
  });

  it("marks the current step, the ones before it, and hides the repair step until one runs", () => {
    const { rerender } = render(<BuildStepper phase={BuildPhase.CHECKING} repaired={false} />);
    const items = screen.getAllByRole("listitem");
    expect(items.map((li) => li.textContent?.replace("·", "").trim())).toEqual(["Understand", "Shortlist", "Build", "Check"]);
    expect(screen.getByText("Check").closest("li")).toHaveAttribute("aria-current", "step");
    rerender(<BuildStepper phase={BuildPhase.REPAIRING} repaired />);
    expect(screen.getByText("Repair").closest("li")).toHaveAttribute("aria-current", "step");
    rerender(<BuildStepper phase={BuildPhase.DONE} repaired />);
    expect(screen.queryByRole("listitem", { current: "step" })).not.toBeInTheDocument();
  });
});
