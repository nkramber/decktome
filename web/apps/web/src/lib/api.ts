import { createClient, type Interceptor } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { AdminService } from "@mtg/api-client/mtg/v1/admin_service_pb";
import { AgentService } from "@mtg/api-client/mtg/v1/agent_service_pb";
import { CardService } from "@mtg/api-client/mtg/v1/card_service_pb";
import { CollectionService } from "@mtg/api-client/mtg/v1/collection_service_pb";
import { DeckService } from "@mtg/api-client/mtg/v1/deck_service_pb";
import { FeedbackService } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { HealthService } from "@mtg/api-client/mtg/v1/health_pb";
import { InviteService } from "@mtg/api-client/mtg/v1/invite_service_pb";
import { PushService } from "@mtg/api-client/mtg/v1/push_service_pb";

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

// The page report of D-1033 goes out as the page leaves view, and a phone
// can suspend the page at once after. A keepalive request outlives the
// page, so the report still reaches the API.
const keepaliveTransport = createConnectTransport({
  baseUrl: import.meta.env.VITE_API_BASE_URL ?? "",
  interceptors: [bearerInterceptor],
  fetch: (input, init) => globalThis.fetch(input, { ...init, keepalive: true }),
});

export const healthClient = createClient(HealthService, transport);
export const collectionClient = createClient(CollectionService, transport);
export const deckClient = createClient(DeckService, transport);
export const agentClient = createClient(AgentService, transport);
export const pageClient = createClient(AgentService, keepaliveTransport);
export const cardClient = createClient(CardService, transport);
export const feedbackClient = createClient(FeedbackService, transport);
export const pushClient = createClient(PushService, transport);
// The invite check runs before an account exists, so it carries no token
// (D-592). The bearer interceptor adds none when nobody is signed in.
export const inviteClient = createClient(InviteService, transport);
// The admin screen needs the admin claim on the token (D-1076).
export const adminClient = createClient(AdminService, transport);
