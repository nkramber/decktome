import type { ReactNode } from "react";
import { createBrowserRouter, createMemoryRouter, type RouteObject } from "react-router";

import { RequireAuth, RootRedirect } from "../features/auth/require-auth";
import { collectionChunk, decksChunk, deckScreenChunk, sessionChunk, signInChunk } from "./chunks";
import { PageFallback } from "./components/page-fallback";
import type { Deferred } from "./deferred";
import { Layout } from "./layout";
import { RouteError } from "./route-error";

// Every page loads on its own route (D-320), so the first paint carries the
// shell alone. The Connect clients and the generated code ride with the
// first page that needs them, not with the shell. The shell and the route
// guard stay eager, so a redirect needs no download.
//
// A page renders with no Suspense boundary. React holds a committed
// fallback for 300 ms, and a page whose chunk is already here must not
// pay it. See `deferred.ts`.
function page(chunk: Deferred<object>): ReactNode {
  const Mount = chunk.Mount;
  function Page() {
    // The download starts with the first render of the route, not at the
    // idle warm-up, so a link typed into the address bar waits for
    // nothing else.
    chunk.preload();
    // A download that fails leaves the reader with a skeleton and no way
    // out, so the failure goes to the error element of the layout.
    const failure = chunk.useError();
    if (failure) throw failure;
    return <Mount fallback={<PageFallback />} />;
  }
  return <Page />;
}

export function appRoutes(extra: RouteObject[] = []): RouteObject[] {
  return [
    {
      element: <Layout />,
      errorElement: <RouteError />,
      children: [
        { path: "/", element: <RootRedirect /> },
        { path: "/sign-in", element: page(signInChunk) },
        {
          element: <RequireAuth />,
          children: [
            { path: "/collection", element: page(collectionChunk) },
            { path: "/session/:id", element: page(sessionChunk) },
            { path: "/decks", element: page(decksChunk) },
            { path: "/decks/:id", element: page(deckScreenChunk) },
            ...extra,
          ],
        },
      ],
    },
  ];
}

export const routes: RouteObject[] = appRoutes();

export function createAppRouter() {
  return createBrowserRouter(routes);
}

// Tests start at one path in memory. extra adds routes under RequireAuth,
// for example one that throws.
export function createTestRouter(initialPath: string, extra: RouteObject[] = []) {
  return createMemoryRouter(appRoutes(extra), { initialEntries: [initialPath] });
}
