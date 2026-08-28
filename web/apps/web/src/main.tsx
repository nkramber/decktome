import "./index.css";

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { RouterProvider } from "react-router";

import { Providers } from "./app/providers";
import { createAppRouter } from "./app/router";

const root = document.getElementById("root");
if (!root) {
  throw new Error("root element not found");
}
createRoot(root).render(
  <StrictMode>
    <Providers>
      <RouterProvider router={createAppRouter()} />
    </Providers>
  </StrictMode>,
);
