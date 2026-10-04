import { useEffect, useState } from "react";

import { useAuth } from "../../features/auth/auth-context";
import { isAdmin } from "../../lib/firebase";

// useAdmin answers whether the signed-in token carries the admin claim
// (D-1076). The answer belongs to one user, so it reads false until the
// token of the current user answers. The API checks the claim again on
// each admin call, and it sends the model spend to the admin alone
// (D-1148).
export function useAdmin(): boolean {
  const { user } = useAuth();
  const [answer, setAnswer] = useState<{ user: unknown; admin: boolean } | null>(null);
  useEffect(() => {
    let live = true;
    void isAdmin().then((admin) => live && setAnswer({ user, admin }));
    return () => {
      live = false;
    };
  }, [user]);
  return answer !== null && answer.user === user && answer.admin;
}
