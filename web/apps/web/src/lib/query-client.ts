import { QueryClient } from "@tanstack/react-query";

// Server state (wallabee tier 1). One client for the app. A failed RPC does
// not retry by itself, so an auth error shows at once instead of after three tries.
export function makeQueryClient(): QueryClient {
  return new QueryClient({
    defaultOptions: { queries: { retry: false, refetchOnWindowFocus: false } },
  });
}
