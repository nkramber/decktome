import { Code, ConnectError } from "@connectrpc/connect";

// The API listens before its card snapshot loads, and CardService reads
// Unavailable until the first index lands (F-164). A cold start of the
// Cloud Run instance holds that state for about 90 seconds, so the art of
// a deck, a binder, or a card option came back empty and stayed empty: no
// query of the app retries by itself (`lib/query-client.ts`), and each
// card query holds its data forever. The reader saw blank tiles until a
// reload.
//
// Every card query now retries through that window, and the art arrives
// with no reload (D-801). The retry takes Unavailable alone. Any other
// code, such as Unauthenticated or InvalidArgument, still fails at once.

// cardRetryDelays back off to 8 seconds. Thirteen retries at these delays
// wait 87 seconds in total, under the 90 seconds of the cold start.
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
