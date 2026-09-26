import { Code, ConnectError } from "@connectrpc/connect";

// The API listens before its card snapshot loads, and CardService reads
// Unavailable until the first index lands (F-164). A cold start of the
// Cloud Run instance holds that state for 10 to 21 seconds (F-176), and
// for about 90 seconds when F-164 was found. So the art of
// a deck, a binder, or a card option came back empty and stayed empty: no
// query of the app retries by itself (`lib/query-client.ts`), and each
// card query holds its data forever. The reader saw blank tiles until a
// reload.
//
// Every card query now retries through that window, and the art arrives
// with no reload (D-801). The retry takes Unavailable alone. Any other
// code, such as Unauthenticated or InvalidArgument, still fails at once.

// cardRetryDelays back off to 8 seconds. Thirteen retries at these delays
// wait 87 seconds in total, far past the load of a cold start.
const cardRetryDelays = [1000, 2000, 4000, 8000];

export const cardRetryMax = 13;

export function cardRetry(failureCount: number, error: unknown): boolean {
  return failureCount < cardRetryMax && ConnectError.from(error).code === Code.Unavailable;
}

export function cardRetryDelay(attemptIndex: number): number {
  return cardRetryDelays[Math.min(attemptIndex, cardRetryDelays.length - 1)];
}

// cardQueryRetry goes into every query that reads CardService.
export const cardQueryRetry = { retry: cardRetry, retryDelay: cardRetryDelay } as const;

// indexLoading reports the refusal of a read that met no card index:
// Unavailable, with the refusal header "index-loading" (D-954). The chat
// sends a turn again on this refusal alone, because a turn that conflicts
// with another turn also reads Unavailable, after a paid call.
export function indexLoading(error: unknown): boolean {
  const err = ConnectError.from(error);
  return err.code === Code.Unavailable && err.metadata.get("deck-tome-refusal") === "index-loading";
}
