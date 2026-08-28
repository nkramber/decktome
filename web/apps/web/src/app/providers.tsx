import { QueryClientProvider } from "@tanstack/react-query";
import { type ReactNode, useState } from "react";

import { AuthProvider } from "../features/auth/auth-context";
import { makeQueryClient } from "../lib/query-client";

// The three state tiers: TanStack Query (server), Zustand (client, no provider), context (auth).
export function Providers({ children }: { children: ReactNode }) {
  const [queryClient] = useState(makeQueryClient);
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>{children}</AuthProvider>
    </QueryClientProvider>
  );
}
