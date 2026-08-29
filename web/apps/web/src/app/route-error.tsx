import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { Link, useRouteError } from "react-router";

import { Button } from "../components/ui/button";
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
    <div className="mx-auto flex max-w-3xl flex-col items-start gap-3 p-4 md:p-6">
      <h1 className="text-2xl font-semibold tracking-tight">Something went wrong</h1>
      <p role="alert" className="wrap-anywhere text-danger">
        {errorMessage(error)}
      </p>
      <div className="flex flex-wrap gap-2">
        <Button asChild variant="outline">
          <Link to="/collection" reloadDocument>
            Go to your collection
          </Link>
        </Button>
        <Button variant="ghost" onClick={() => void onSignOut()}>
          Sign out
        </Button>
      </div>
      {signOutError && (
        <p role="alert" className="text-sm text-danger">
          Sign-out failed: {signOutError}
        </p>
      )}
    </div>
  );
}
