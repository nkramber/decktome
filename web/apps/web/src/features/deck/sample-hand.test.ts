import { CardRole, type DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import { describe, expect, it } from "vitest";

import { bottom, draw, libraryOf, maxMulligans, mulligan, openingHand, seededRandom } from "./sample-hand";

const entry = (oracleId: string, name: string, count: number): DeckCard =>
  ({ oracleId, name, count, role: CardRole.OTHER, owned: true, ownedCount: count, priceUsd: 0, reason: "" }) as DeckCard;

// A 60-card deck: 24 lands and nine spells in playsets.
const sixty: DeckCard[] = [
  entry("o-forest", "Forest", 12),
  entry("o-island", "Island", 12),
  entry("o-elf", "Llanowar Elves", 4),
  entry("o-bolt", "Lightning Bolt", 4),
  entry("o-counter", "Counterspell", 4),
  entry("o-bear", "Grizzly Bears", 4),
  entry("o-growth", "Giant Growth", 4),
  entry("o-opt", "Opt", 4),
  entry("o-snake", "Snapcaster Mage", 4),
  entry("o-tarmo", "Tarmogoyf", 4),
  entry("o-delver", "Delver of Secrets", 4),
];

function tally(names: string[]): Map<string, number> {
  const out = new Map<string, number>();
  for (const n of names) out.set(n, (out.get(n) ?? 0) + 1);
  return out;
}

describe("sample hand", () => {
  it("expands the main deck into one card per copy and keeps the commander out", () => {
    const library = libraryOf([entry("o-cmd", "Commander", 1), entry("o-forest", "Forest", 3)], ["o-cmd"]);
    expect(library.map((c) => c.name)).toEqual(["Forest", "Forest", "Forest"]);
  });

  it("draws every card of a 60-card deck through the hand (PR-20 gate)", () => {
    const library = libraryOf(sixty, []);
    expect(library).toHaveLength(60);
    let state = openingHand(library, seededRandom(7));
    expect(state.hand).toHaveLength(7);
    expect(state.library).toHaveLength(53);
    while (state.library.length > 0) {
      const before = state.library.length;
      state = draw(state);
      expect(state.library).toHaveLength(before - 1);
    }
    expect(state.hand).toHaveLength(60);
    expect(tally(state.hand.map((c) => c.name))).toEqual(tally(library.map((c) => c.name)));
    // An empty library changes nothing.
    expect(draw(state)).toBe(state);
  });

  it("mulligans by the London rule to six and to five, and refuses a third", () => {
    const random = seededRandom(3);
    let state = openingHand(libraryOf(sixty, []), random);
    state = mulligan(state, random);
    expect(state.mulligans).toBe(1);
    expect(state.hand).toHaveLength(7);
    expect(state.toBottom).toBe(1);
    // A draw and a second mulligan wait for the bottom choice.
    expect(draw(state)).toBe(state);
    expect(mulligan(state, random)).toBe(state);
    const bottomed = state.hand[2];
    state = bottom(state, 2);
    expect(state.hand).toHaveLength(6);
    expect(state.toBottom).toBe(0);
    expect(state.library.at(-1)).toEqual(bottomed);
    expect(state.library).toHaveLength(54);
    state = mulligan(state, random);
    expect(state.mulligans).toBe(2);
    expect(state.toBottom).toBe(2);
    state = bottom(bottom(state, 0), 0);
    expect(state.hand).toHaveLength(5);
    expect(state.library).toHaveLength(55);
    expect(mulligan(state, random)).toBe(state);
    expect(maxMulligans).toBe(2);
    state = draw(state);
    expect(state.hand).toHaveLength(6);
  });

  it("draws the same hand twice from one seed", () => {
    const library = libraryOf(sixty, []);
    const a = openingHand(library, seededRandom(42));
    const b = openingHand(library, seededRandom(42));
    expect(a.hand).toEqual(b.hand);
    expect(bottom(a, 9)).toBe(a);
  });
});
