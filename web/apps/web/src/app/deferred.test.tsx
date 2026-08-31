import { render, screen } from "@testing-library/react";
import { describe, expect, it, vi } from "vitest";

import { deferred } from "./deferred";

function host(unit: ReturnType<typeof deferred<object>>) {
  const Mount = unit.Mount;
  return function Host() {
    unit.preload();
    const failure = unit.useError();
    if (failure) return <p>{`failed: ${String(failure)}`}</p>;
    return <Mount fallback={<p>waiting</p>} />;
  };
}

describe("a deferred chunk", () => {
  it("shows the component once the code arrives, with no Suspense", async () => {
    const unit = deferred(() => Promise.resolve({ default: () => <p>here</p> }));
    const Host = host(unit);
    render(<Host />);
    expect(screen.getByText("waiting")).toBeInTheDocument();
    expect(await screen.findByText("here")).toBeInTheDocument();
  });

  it("renders at once when the code is already in hand", async () => {
    const unit = deferred(() => Promise.resolve({ default: () => <p>here</p> }));
    unit.preload();
    const Host = host(unit);
    // Let the download settle before the first render of the host.
    await new Promise((resolve) => setTimeout(resolve, 0));
    const { container } = render(<Host />);
    // No pass through the waiting state: the first paint holds the page.
    expect(container.textContent).toBe("here");
  });

  it("loads once for many callers", async () => {
    const load = vi.fn(() => Promise.resolve({ default: () => <p>here</p> }));
    const unit = deferred(load);
    const Host = host(unit);
    render(
      <>
        <Host />
        <Host />
      </>,
    );
    await screen.findAllByText("here");
    expect(load).toHaveBeenCalledTimes(1);
  });

  // A chunk that is gone after a new release never arrives, so the unit
  // asks once. The caller shows the failure, and the reader reloads.
  it("gives the failure to its caller, and asks no more", async () => {
    const load = vi.fn().mockRejectedValue(new Error("chunk is gone"));
    const unit = deferred(load);
    const Host = host(unit);
    render(<Host />);
    expect(await screen.findByText("failed: Error: chunk is gone")).toBeInTheDocument();
    expect(load).toHaveBeenCalledTimes(1);
  });
});
