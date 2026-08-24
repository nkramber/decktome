import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { HealthService } from "@mtg/api-client/mtg/v1/health_pb";

// Same-origin in dev (Vite proxy) and in prod (one host). Override with VITE_API_BASE_URL.
const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_API_BASE_URL ?? "",
});

export const healthClient = createClient(HealthService, transport);
