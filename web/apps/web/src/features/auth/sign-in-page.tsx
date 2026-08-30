import { type FormEvent, useState } from "react";
import { type Location, Navigate, useLocation } from "react-router";

import { Button } from "../../components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "../../components/ui/card";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { signInErrorMessage } from "../../lib/errors";
import { createAccount, signIn } from "../../lib/firebase";
import { useAuth } from "./auth-context";

// One form for sign-in and sign-up (D-275). The emulator accepts any email
// and any password of six or more characters. A visit the route guard
// redirected goes back to the page it wanted after the sign-in.
export function SignInPage() {
  const { user, ready, error: authError } = useAuth();
  const location = useLocation();
  const from = (location.state as { from?: Location } | null)?.from?.pathname ?? "/session/new";
  const [mode, setMode] = useState<"sign-in" | "sign-up">("sign-in");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [busy, setBusy] = useState(false);

  if (!ready) return null;
  if (user) {
    return <Navigate to={from} replace />;
  }

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    setBusy(true);
    setError("");
    try {
      if (mode === "sign-up") {
        await createAccount(email, password);
      } else {
        await signIn(email, password);
      }
    } catch (err) {
      setError(signInErrorMessage(err));
    } finally {
      setBusy(false);
    }
  }

  const creating = mode === "sign-up";

  return (
    <div className="mx-auto flex min-h-[80vh] w-full max-w-sm flex-col justify-center p-4 md:p-6">
      <Card className="shadow-raised">
        <CardHeader>
          <CardTitle asChild className="text-2xl tracking-tight">
            <h1>{creating ? "Create account" : "Sign in"}</h1>
          </CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={onSubmit} className="flex flex-col gap-3">
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="email">Email</Label>
              <Input id="email" type="email" name="email" autoComplete="email" required value={email} onChange={(e) => setEmail(e.target.value)} />
            </div>
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="password">Password</Label>
              <Input
                id="password"
                type="password"
                name="password"
                autoComplete={creating ? "new-password" : "current-password"}
                required
                minLength={6}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <Button type="submit" disabled={busy}>
              {creating ? "Create account" : "Sign in"}
            </Button>
            <div role="alert" className="min-h-6 text-sm text-danger">
              {error || authError}
            </div>
          </form>
          <Button
            variant="link"
            className="mt-2 px-0"
            onClick={() => {
              setMode(creating ? "sign-in" : "sign-up");
              setError("");
            }}
          >
            {creating ? "I have an account. Sign in." : "New here? Create account"}
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
