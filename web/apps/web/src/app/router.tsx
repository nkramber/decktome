import { createBrowserRouter, createMemoryRouter, type RouteObject } from "react-router";

import { RequireAuth, RootRedirect } from "../features/auth/require-auth";
import { SignInPage } from "../features/auth/sign-in-page";
import { SessionPage } from "../features/chat/session-page";
import { CollectionPage } from "../features/collection/collection-page";
import { DecksPage } from "../features/deck/decks-page";
import { Layout } from "./layout";
import { RouteError } from "./route-error";

// Four routes (ui plan, section 5). Every route except /sign-in sits under
// RequireAuth. The layout route catches a render error from any page.
export function appRoutes(extra: RouteObject[] = []): RouteObject[] {
  return [
    {
      element: <Layout />,
      errorElement: <RouteError />,
      children: [
        { path: "/", element: <RootRedirect /> },
        { path: "/sign-in", element: <SignInPage /> },
        {
          element: <RequireAuth />,
          children: [{ path: "/collection", element: <CollectionPage /> }, { path: "/session/:id", element: <SessionPage /> }, { path: "/decks", element: <DecksPage /> }, ...extra],
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
