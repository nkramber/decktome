import { createBrowserRouter, createMemoryRouter, type RouteObject } from "react-router";

import { RequireAuth, RootRedirect } from "../features/auth/require-auth";
import { SignInPage } from "../features/auth/sign-in-page";
import { SessionPage } from "../features/chat/session-page";
import { CollectionPage } from "../features/collection/collection-page";
import { DecksPage } from "../features/deck/decks-page";
import { Layout } from "./layout";

// Four routes (ui plan, section 5). Every route except /sign-in sits under RequireAuth.
export const routes: RouteObject[] = [
  {
    element: <Layout />,
    children: [
      { path: "/", element: <RootRedirect /> },
      { path: "/sign-in", element: <SignInPage /> },
      {
        element: <RequireAuth />,
        children: [
          { path: "/collection", element: <CollectionPage /> },
          { path: "/session/:id", element: <SessionPage /> },
          { path: "/decks", element: <DecksPage /> },
        ],
      },
    ],
  },
];

export function createAppRouter() {
  return createBrowserRouter(routes);
}

// Tests start at one path in memory.
export function createTestRouter(initialPath: string) {
  return createMemoryRouter(routes, { initialEntries: [initialPath] });
}
