import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, expect, it, vi } from "vitest";

import { CardOption } from "./card-options";
// A single card never zooms, and a half of a commander pair zooms by two
// from its outer corner (D-443).
describe("the card zoom", () => {
  const card = {
    oracleId: "o1",
    name: "Thrasios, Triton Hero",
    faces: [{ name: "Thrasios, Triton Hero", imageUris: { normal: "https://x/t.jpg" } }],
  } as never;

  it("gives a single card no zoom", () => {
    render(<CardOption card={card} name="Thrasios" />);
    const span = screen.getByRole("img").closest("span");
    expect(span?.className).not.toContain("scale-[2]");
    expect(span).not.toHaveAttribute("data-zoom");
  });

  it("zooms each half of a pair from its outer corner", () => {
    render(
      <>
        <CardOption card={card} name="Thrasios" zoom="left" />
        <CardOption card={card} name="Tymna" zoom="right" />
      </>,
    );
    const [left, right] = screen.getAllByRole("img").map((img) => img.closest("span"));
    expect(left?.className).toContain("group-hover/card:scale-[2]");
    expect(left?.className).toContain("origin-top-left");
    expect(right?.className).toContain("origin-top-right");
  });
});

// The art picks the option too (D-444). It stays out of the tab order,
// because the name button is the control a keyboard uses.
describe("a pick from the art", () => {
  const card = { oracleId: "o1", name: "Toski", faces: [{ name: "Toski", imageUris: { normal: "https://x/t.jpg" } }] } as never;

  it("calls the pick on a click of the image, and keeps the image readable", async () => {
    const onPick = vi.fn();
    render(<CardOption card={card} name="Toski" onPick={onPick} />);
    const art = screen.getByTestId("card-art-pick");
    // The image stays an image for assistive technology.
    expect(screen.getByRole("img", { name: /Toski/ })).toBeInTheDocument();
    expect(art).not.toHaveAttribute("tabindex");
    await userEvent.setup().click(screen.getByRole("img", { name: /Toski/ }));
    expect(onPick).toHaveBeenCalledTimes(1);
  });

  it("is plain art with no pick", () => {
    render(<CardOption card={card} name="Toski" />);
    expect(screen.queryByTestId("card-art-pick")).not.toBeInTheDocument();
  });
});
