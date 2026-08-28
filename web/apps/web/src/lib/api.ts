import { createClient, type Interceptor } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { AgentService } from "@mtg/api-client/mtg/v1/agent_service_pb";
import { CardService } from "@mtg/api-client/mtg/v1/card_service_pb";
import { CollectionService } from "@mtg/api-client/mtg/v1/collection_service_pb";
import { DeckService } from "@mtg/api-client/mtg/v1/deck_service_pb";
import { HealthService } from "@mtg/api-client/mtg/v1/health_pb";

import { currentIdToken } from "./firebase";

// The bearer interceptor adds the Firebase ID token to every RPC (D-268).
// A request without a signed-in user goes out without the header, and the
// API answers Unauthenticated. The web app never reads Firestore directly.
export const bearerInterceptor: Interceptor = (next) => async (req) => {
  const token = await currentIdToken();
  if (token) {
    req.header.set("Authorization", `Bearer ${token}`);
  }
  return next(req);
};

// Same-origin in dev (Vite proxy) and in prod (one host). Override with VITE_API_BASE_URL.
export const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_API_BASE_URL ?? "",
  interceptors: [bearerInterceptor],
});

export const healthClient = createClient(HealthService, transport);
export const collectionClient = createClient(CollectionService, transport);
export const deckClient = createClient(DeckService, transport);
export const agentClient = createClient(AgentService, transport);
export const cardClient = createClient(CardService, transport);
