import { lazy, Suspense, type ReactNode } from "react";
import { createBrowserRouter, createMemoryRouter, type RouteObject } from "react-router";

import { RequireAuth, RootRedirect } from "../features/auth/require-auth";
import { PageFallback } from "./components/page-fallback";
import { Layout } from "./layout";
import { RouteError } from "./route-error";

// Every page loads on its own route (D-320), so the first paint carries the
// shell alone. The Connect clients and the generated code ride with the
// first page that needs them, not with the shell. The shell and the route
// guard stay eager, so a redirect needs no download.
const SignInPage = lazy(async () => ({ default: (await import("../features/auth/sign-in-page")).SignInPage }));
const CollectionPage = lazy(async () => ({ default: (await import("../features/collection/collection-page")).CollectionPage }));
const SessionPage = lazy(async () => ({ default: (await import("../features/chat/session-page")).SessionPage }));
const DecksPage = lazy(async () => ({ default: (await import("../features/deck/decks-page")).DecksPage }));

function page(node: ReactNode) {
  return <Suspense fallback={<PageFallback />}>{node}</Suspense>;
}

export function appRoutes(extra: RouteObject[] = []): RouteObject[] {
  return [
    {
      element: <Layout />,
      errorElement: <RouteError />,
      children: [
        { path: "/", element: <RootRedirect /> },
        { path: "/sign-in", element: page(<SignInPage />) },
        {
          element: <RequireAuth />,
          children: [
            { path: "/collection", element: page(<CollectionPage />) },
            { path: "/session/:id", element: page(<SessionPage />) },
            { path: "/decks", element: page(<DecksPage />) },
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
