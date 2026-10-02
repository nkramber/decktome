import type { Card, ImageUris } from "@mtg/api-client/mtg/v1/card_pb";
import { describe, expect, it } from "vitest";

import { facesOf, sharedFront } from "./card-tile";

const img = (n: string) => ({ small: `https://cards.scryfall.io/small/${n}.jpg`, normal: `https://cards.scryfall.io/normal/${n}.jpg` }) as unknown as ImageUris;

// The card data gives each half of an adventure the image of the whole
// card (go/internal/cards/parse.go), so both halves carry one URL.
const adventure = {
  oracleId: "o-giant",
  name: "Bonecrusher Giant // Stomp",
  typeLine: "Creature — Giant // Instant — Adventure",
  manaCost: "{2}{R} // {1}{R}",
  oracleText: "",
  layout: "adventure",
  faces: [
    { name: "Bonecrusher Giant", artist: "Victor Adame Minguez", imageUris: img("giant"), typeLine: "Creature — Giant", oracleText: "Whenever Bonecrusher Giant becomes the target of a spell, Bonecrusher Giant deals 2 damage to that spell's controller.", manaCost: "{2}{R}" },
    { name: "Stomp", artist: "Victor Adame Minguez", imageUris: img("giant"), typeLine: "Instant — Adventure", oracleText: "Damage can't be prevented this turn. Stomp deals 2 damage to any target.", manaCost: "{1}{R}" },
  ],
  defaultPrinting: { artist: "Victor Adame Minguez", imageUris: img("giant") },
} as unknown as Card;

const doubleFaced = {
  oracleId: "o-dfc",
  name: "Delver of Secrets // Insectile Aberration",
  typeLine: "Creature — Human Wizard // Creature — Human Insect",
  layout: "transform",
  faces: [
    { name: "Delver of Secrets", artist: "Nils Hamm", imageUris: img("delver-a"), typeLine: "Creature — Human Wizard", oracleText: "", manaCost: "{U}" },
    { name: "Insectile Aberration", artist: "Nils Hamm", imageUris: img("delver-b"), typeLine: "Creature — Human Insect", oracleText: "Flying", manaCost: "" },
  ],
} as unknown as Card;

describe("facesOf (D-1048)", () => {
  it("shows one front for a card whose halves share it, with the text of each half", () => {
    expect(sharedFront(adventure)).toBe(true);
    const faces = facesOf(adventure);
    expect(faces).toHaveLength(1);
    expect(faces[0].name).toBe("Bonecrusher Giant // Stomp");
    expect(faces[0].imageUris?.normal).toBe("https://cards.scryfall.io/normal/giant.jpg");
    expect(faces[0].halves?.map((h) => h.name)).toEqual(["Bonecrusher Giant", "Stomp"]);
  });

  it("shows the owned printing of a shared front (D-299, D-1041)", () => {
    const faces = facesOf(adventure, { imageUris: img("giant-owned"), artist: "Owned Artist" });
    expect(faces).toHaveLength(1);
    expect(faces[0].imageUris?.normal).toBe("https://cards.scryfall.io/normal/giant-owned.jpg");
    expect(faces[0].artist).toBe("Owned Artist");
  });

  it("keeps both faces of a double-faced card (F-9)", () => {
    expect(sharedFront(doubleFaced)).toBe(false);
    const faces = facesOf(doubleFaced, { imageUris: img("delver-owned"), artist: "x" });
    expect(faces.map((f) => f.imageUris?.normal)).toEqual(["https://cards.scryfall.io/normal/delver-a.jpg", "https://cards.scryfall.io/normal/delver-b.jpg"]);
    expect(faces[0].halves).toBeUndefined();
  });
});
