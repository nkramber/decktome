import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { Link, useRouteError } from "react-router";

import { errorMessage } from "../lib/errors";
import { useAppStore } from "../lib/store";
import { signOutAndClear } from "./layout";

// RouteError is the error element of the layout route. A render error in
// any page lands here with the message, a way back, and a way out.
export function RouteError() {
  const error = useRouteError();
  const queryClient = useQueryClient();
  const reset = useAppStore((s) => s.reset);
  const [signOutError, setSignOutError] = useState("");

  async function onSignOut() {
    try {
      await signOutAndClear(reset, () => queryClient.clear());
      window.location.assign("/sign-in");
    } catch (err) {
      setSignOutError(errorMessage(err));
    }
  }

  return (
    <div className="mx-auto flex max-w-3xl flex-col gap-3 p-6">
      <h1 className="text-2xl font-semibold">Something went wrong</h1>
      <p role="alert" className="wrap-anywhere text-red-700">
        {errorMessage(error)}
      </p>
      <p className="flex gap-4">
        <Link to="/collection" className="underline" reloadDocument>
          Go to your collection
        </Link>
        <button type="button" onClick={() => void onSignOut()} className="underline">
          Sign out
        </button>
      </p>
      {signOutError && (
        <p role="alert" className="text-sm text-red-700">
          Sign-out failed: {signOutError}
        </p>
      )}
    </div>
  );
}
