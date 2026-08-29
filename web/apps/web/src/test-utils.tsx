import { render } from "@testing-library/react";
import { type RouteObject, RouterProvider } from "react-router";

import { Providers } from "./app/providers";
import { createTestRouter } from "./app/router";

// renderAt mounts the real providers and router at one path.
// Auth comes from the firebase/auth mock in each test file.
export function renderAt(path: string, extra: RouteObject[] = []) {
  const router = createTestRouter(path, extra);
  const view = render(
    <Providers>
      <RouterProvider router={router} />
    </Providers>,
  );
  return { ...view, router };
}
