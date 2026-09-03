import type { DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";

// The sample hand (D-318): draw seven from the exact main deck, mulligan
// to six and to five, and draw one. No turn simulation. The library is
// the main deck without the command zone, one entry per copy, so a deck
// of 60 cards yields 60 draws and then runs dry.

export type HandCard = { oracleId: string; name: string };

// mulberry32 is a small seeded generator, so a test draws the same
// hand twice and a reader can share a seed.
export function seededRandom(seed: number): () => number {
  let a = seed >>> 0;
  return () => {
    a = (a + 0x6d2b79f5) >>> 0;
    let t = a;
    t = Math.imul(t ^ (t >>> 15), t | 1);
    t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
    return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
  };
}

// libraryOf expands the main deck into one card per copy, in deck order.
// A commander stays in the command zone and never in the library.
export function libraryOf(cards: DeckCard[], commanderIds: string[]): HandCard[] {
  const commanders = new Set(commanderIds);
  const out: HandCard[] = [];
  for (const dc of cards) {
    if (commanders.has(dc.oracleId)) continue;
    for (let i = 0; i < dc.count; i++) out.push({ oracleId: dc.oracleId, name: dc.name });
  }
  return out;
}

export function shuffle<T>(list: T[], random: () => number): T[] {
  const out = [...list];
  for (let i = out.length - 1; i > 0; i--) {
    const j = Math.floor(random() * (i + 1));
    [out[i], out[j]] = [out[j], out[i]];
  }
  return out;
}

// Hand is the state of one sample: the library with its top at index 0,
// the cards in hand, the mulligans taken, and how many cards the London
// rule still asks the player to put on the bottom.
export type Hand = {
  library: HandCard[];
  hand: HandCard[];
  mulligans: number;
  toBottom: number;
};

export const handSize = 7;

// maxMulligans is two: the sample goes to six and then to five (D-318).
export const maxMulligans = 2;

// openingHand shuffles the library and draws seven.
export function openingHand(library: HandCard[], random: () => number): Hand {
  const shuffled = shuffle(library, random);
  return { library: shuffled.slice(handSize), hand: shuffled.slice(0, handSize), mulligans: 0, toBottom: 0 };
}

// mulligan is the London rule: the hand goes back, the library shuffles,
// seven come again, and the player puts one card per mulligan on the
// bottom. The third mulligan is refused (D-318).
export function mulligan(state: Hand, random: () => number): Hand {
  if (state.mulligans >= maxMulligans || state.toBottom > 0) return state;
  const shuffled = shuffle([...state.library, ...state.hand], random);
  const mulligans = state.mulligans + 1;
  return { library: shuffled.slice(handSize), hand: shuffled.slice(0, handSize), mulligans, toBottom: mulligans };
}

// bottom puts one card of the hand on the bottom of the library.
export function bottom(state: Hand, index: number): Hand {
  if (state.toBottom <= 0 || index < 0 || index >= state.hand.length) return state;
  const card = state.hand[index];
  const hand = state.hand.filter((_, i) => i !== index);
  return { ...state, hand, library: [...state.library, card], toBottom: state.toBottom - 1 };
}

// draw takes the top card of the library. An empty library changes
// nothing, and the panel says so.
export function draw(state: Hand): Hand {
  if (state.toBottom > 0 || state.library.length === 0) return state;
  const [top, ...rest] = state.library;
  return { ...state, hand: [...state.hand, top], library: rest };
}
