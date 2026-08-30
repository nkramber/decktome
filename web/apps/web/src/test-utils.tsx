import { act, render } from "@testing-library/react";
import { type RouteObject, RouterProvider } from "react-router";

import { Providers } from "./app/providers";
import { createTestRouter } from "./app/router";

// The pages and the auth SDK load on demand (D-320). A test renders the
// real router, so renderAt warms those modules first and then lets React
// mount one. After the await, a synchronous query finds the page and the
// auth listener has reported. Without the warm-up the first render of a
// file resolves one module later than the rest, and only that one fails.
let loaded: Promise<unknown> | null = null;

function loadModules() {
  loaded ??= Promise.all([
    import("./features/auth/sign-in-page"),
    import("./features/collection/collection-page"),
    import("./features/chat/session-page"),
    import("./features/deck/decks-page"),
    import("./features/deck/deck-page"),
    import("firebase/app"),
    import("firebase/auth"),
  ]);
  return loaded;
}

// renderAt mounts the real providers and router at one path.
// Auth comes from the firebase/auth mock in each test file.
export async function renderAt(path: string, extra: RouteObject[] = []) {
  await loadModules();
  const router = createTestRouter(path, extra);
  const view = render(
    <Providers>
      <RouterProvider router={router} />
    </Providers>,
  );
  await act(async () => {});
  return { ...view, router };
}
