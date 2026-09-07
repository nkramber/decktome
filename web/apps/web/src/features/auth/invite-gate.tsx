import { Code, ConnectError } from "@connectrpc/connect";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { type ReactNode, useEffect } from "react";

import { Button } from "../../components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "../../components/ui/card";
import { collectionClient } from "../../lib/api";
import { useAppStore } from "../../lib/store";
import { useAuth } from "./auth-context";
import { type InviteState, setInviteState, useInviteState } from "./invite-state";
import { LoadingSession } from "./require-auth";
import { signOutAndClear } from "./sign-out";

// The header the API sets on a refusal, and the one value it takes
// (F-59). It matches auth.RefusalHeader and auth.RefusalNotInvited on the
// Go side. The state reads from the header alone, never from the
// sentence, so the wording of the message stays free to change.
const refusalHeader = "deck-tome-refusal";
const refusalNotInvited = "not-invited";

// isNotInvited says whether the API refused this call because the email
// is off the invite list (D-314). Every other denial reads false.
export function isNotInvited(err: unknown): boolean {
  return err instanceof ConnectError && err.code === Code.PermissionDenied && err.metadata.get(refusalHeader) === refusalNotInvited;
}

// useInviteProbe asks the API whether this reader is on the list. The
// list of collections is the probe, because the app reads it anyway: the
// pool picker and the collection screen hold the same key, so one call
// pays for all three. Any answer but the refusal reads as invited, so an
// API that is down never locks an invited reader out.
//
// The answer latches, and this is what makes the shared key safe. The
// pages under the gate read the same key, so a page that mounts its own
// reader starts a refetch. Without the latch the gate would read
// "unknown" again, take the page off the screen, end that reader, and
// start over. A refusal still wins over a latched answer, because an
// invite the owner takes back must reach the reader.
function useInviteProbe(): InviteState {
  const known = useInviteState();
  const { data, error } = useQuery({
    queryKey: ["collections"],
    queryFn: () => collectionClient.listCollections({}),
  });
  const answer: InviteState = isNotInvited(error) ? "refused" : data !== undefined || error ? "invited" : "unknown";
  useEffect(() => {
    if (answer !== "unknown") setInviteState(answer);
  }, [answer]);
  if (answer === "refused") return "refused";
  return known === "unknown" ? answer : known;
}

// The gate holds the protected pages until the API names the state
// (F-59). Before the fix a reader off the list read the whole shell and
// a Build button that always failed.
export default function InviteGate({ children }: { children: ReactNode }) {
  const state = useInviteProbe();
  if (state === "unknown") return <LoadingSession />;
  if (state === "refused") return <NotInvited />;
  return <>{children}</>;
}

// NotInvited names the state and offers the one way out. The app is open
// to invited users alone (D-314), and the reader holds an account that
// the project never invited.
export function NotInvited() {
  const { user } = useAuth();
  const queryClient = useQueryClient();
  const reset = useAppStore((s) => s.reset);
  // The same sign-out as the header. It forgets the latched answer, so
  // the next account on this browser is asked again.
  const onSignOut = () => void signOutAndClear(reset, () => queryClient.clear());
  return (
    <div className="mx-auto flex min-h-[80vh] w-full max-w-md flex-col justify-center p-4 md:p-6">
      <Card className="shadow-raised">
        <CardHeader>
          <CardTitle asChild className="text-2xl tracking-tight">
            <h1>You are not on the invite list</h1>
          </CardTitle>
        </CardHeader>
        <CardContent className="flex flex-col gap-4 text-sm text-muted-foreground">
          <p>
            Deck Tome is open to invited users alone. Your account{" "}
            <span className="font-mono text-foreground">{user?.email ?? ""}</span> is signed in, but the list does not hold it.
          </p>
          <p>Ask the owner for an invite for this email address. A new invite needs no new account.</p>
          <Button variant="outline" className="self-start" onClick={onSignOut}>
            Sign out
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
