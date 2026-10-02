import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { afterEach, describe, expect, it, vi } from "vitest";

import { type Face, FaceImage } from "./card-tile";
import { ZoomFace } from "./card-zoom";

const face: Face = {
  name: "Sol Ring",
  artist: "Mark Tedin",
  imageUris: { small: "https://cards.scryfall.io/small/sol.jpg", normal: "https://cards.scryfall.io/normal/sol.jpg", large: "https://cards.scryfall.io/large/sol.jpg", artCrop: "" },
  typeLine: "Artifact",
  oracleText: "{T}: Add {C}{C}.",
  manaCost: "{1}",
} as Face;

function pointer(coarse: boolean) {
  vi.stubGlobal(
    "matchMedia",
    vi.fn((query: string) => ({ matches: coarse && query === "(pointer: coarse)", media: query, addEventListener: vi.fn(), removeEventListener: vi.fn() })),
  );
}

function renderFace() {
  return render(
    <ZoomFace face={face}>
      <FaceImage face={face} />
    </ZoomFace>,
  );
}

afterEach(() => vi.unstubAllGlobals());

describe("ZoomFace (D-1049)", () => {
  it("opens the large image on a tap, keeps it on a tap of the image, and closes on a tap beside it", async () => {
    pointer(true);
    const user = userEvent.setup();
    renderFace();
    await user.click(screen.getByRole("button", { name: "Show Sol Ring larger" }));
    const large = await screen.findByAltText("Sol Ring (card, large)");
    expect(large).toHaveAttribute("src", "https://cards.scryfall.io/large/sol.jpg");
    await user.click(large);
    expect(screen.getByTestId("card-zoom")).toBeInTheDocument();
    await user.click(screen.getByTestId("card-zoom"));
    await waitFor(() => expect(screen.queryByTestId("card-zoom")).not.toBeInTheDocument());
  });

  it("closes on the Escape key", async () => {
    pointer(true);
    const user = userEvent.setup();
    renderFace();
    await user.click(screen.getByRole("button", { name: "Show Sol Ring larger" }));
    await screen.findByTestId("card-zoom");
    await user.keyboard("{Escape}");
    await waitFor(() => expect(screen.queryByTestId("card-zoom")).not.toBeInTheDocument());
  });

  it("leaves the art a plain image under a mouse", () => {
    pointer(false);
    renderFace();
    expect(screen.queryByRole("button")).not.toBeInTheDocument();
    expect(screen.getByAltText("Sol Ring (card)")).toBeInTheDocument();
  });
});
