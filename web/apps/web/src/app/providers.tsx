import { QueryClientProvider } from "@tanstack/react-query";
import { type ReactNode, useEffect, useState } from "react";

import { AuthProvider } from "../features/auth/auth-context";
import { makeQueryClient } from "../lib/query-client";
import { watchSystemTheme } from "../lib/theme";

// The three state tiers: TanStack Query (server), Zustand (client, no provider), context (auth).
export function Providers({ children }: { children: ReactNode }) {
  const [queryClient] = useState(makeQueryClient);

  // The theme follows the system while the choice is "system" (D-311).
  useEffect(() => watchSystemTheme(), []);

  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>{children}</AuthProvider>
    </QueryClientProvider>
  );
}
