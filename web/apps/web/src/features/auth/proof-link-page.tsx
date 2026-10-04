import { ConnectError } from "@connectrpc/connect";
import { useEffect, useRef, useState } from "react";
import { useNavigate, useParams } from "react-router";

import { Button } from "../../components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "../../components/ui/card";
import { errorMessage } from "../../lib/errors";
import { iosBrowser, openProofLink } from "../../lib/proof";

// failure reads the sentence of a refusal. The API writes each sentence
// for the user (D-1082).
function failure(err: unknown): string {
  const text = err instanceof ConnectError ? err.rawMessage : errorMessage(err);
  return text.charAt(0).toUpperCase() + text.slice(1) + (text.endsWith(".") ? "" : ".");
}

type State =
  | { kind: "ready" }
  | { kind: "opening" }
  | { kind: "ios" }
  | { kind: "proved" }
  | { kind: "already-proved" }
  | { kind: "failed"; message: string };

// ProofLinkPage opens the link of the email that proves an address
// (D-1081). The link works one time, so the page uses it on a click
// alone, and a mail scanner that loads the page leaves it unused
// (D-1119). The API proves the email, and this browser signs in (D-1082).
// A desktop browser and an installed Android app go on to the app at
// once. iOS opens the link in Safari and never in the installed app, so
// Safari tells the user where to go (D-1083). The installed app sees the
// proof when it becomes visible again.
export function ProofLinkPage() {
  const { code = "" } = useParams();
  const navigate = useNavigate();
  const [state, setState] = useState<State>({ kind: "ready" });
  const live = useRef(true);

  useEffect(() => {
    live.current = true;
    return () => {
      live.current = false;
    };
  }, []);

  function confirm() {
    if (state.kind !== "ready") return;
    setState({ kind: "opening" });
    openProofLink(code).then(
      (outcome) => {
        if (!live.current) return;
        if (outcome === "already-proved") setState({ kind: "already-proved" });
        else if (outcome === "proved") setState({ kind: "proved" });
        else if (iosBrowser()) setState({ kind: "ios" });
        else void navigate("/", { replace: true });
      },
      (err: unknown) => live.current && setState({ kind: "failed", message: failure(err) }),
    );
  }

  const toSignIn = <Button onClick={() => void navigate("/sign-in", { replace: true })}>Return to login page</Button>;

  return (
    <div className="mx-auto flex min-h-[80vh] w-full max-w-sm flex-col justify-center p-4 md:p-6">
      <Card className="shadow-raised">
        <CardHeader>
          <CardTitle asChild className="text-2xl tracking-tight">
            <h1>{state.kind === "failed" ? "This link does not work" : "Verify your email"}</h1>
          </CardTitle>
        </CardHeader>
        <CardContent className="flex flex-col gap-4 text-sm">
          {(state.kind === "ready" || state.kind === "opening") && (
            <>
              <p className="text-muted-foreground">Confirm that this is your email address to finish your sign-up.</p>
              <Button onClick={confirm} disabled={state.kind === "opening"}>
                Confirm my email
              </Button>
              {state.kind === "opening" && (
                <p role="status" className="text-muted-foreground">
                  Verifying your email…
                </p>
              )}
            </>
          )}
          {state.kind === "ios" && (
            <>
              <p role="status">Your email is verified. Open Deck Tome from your home screen to continue.</p>
              <Button variant="outline" onClick={() => void navigate("/", { replace: true })}>
                Continue in the browser
              </Button>
            </>
          )}
          {state.kind === "proved" && (
            <>
              <p role="status">Your email is verified. Sign in to continue.</p>
              {toSignIn}
            </>
          )}
          {state.kind === "already-proved" && (
            <>
              <p role="status">Your email is already verified. Sign in to continue.</p>
              {toSignIn}
            </>
          )}
          {state.kind === "failed" && (
            <>
              <p role="alert" className="text-danger">
                {state.message}
              </p>
              {toSignIn}
            </>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
