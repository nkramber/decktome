import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from "firebase/auth";
import { type FormEvent, useState } from "react";
import { type Location, Navigate, useLocation } from "react-router";

import { signInErrorMessage } from "../../lib/errors";
import { auth } from "../../lib/firebase";
import { useAuth } from "./auth-context";

// One form for sign-in and sign-up (D-275). The emulator accepts any email
// and any password of six or more characters. A visit the route guard
// redirected goes back to the page it wanted after the sign-in.
export function SignInPage() {
  const { user, ready, error: authError } = useAuth();
  const location = useLocation();
  const from = (location.state as { from?: Location } | null)?.from?.pathname ?? "/collection";
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
        await createUserWithEmailAndPassword(auth, email, password);
      } else {
        await signInWithEmailAndPassword(auth, email, password);
      }
    } catch (err) {
      setError(signInErrorMessage(err));
    } finally {
      setBusy(false);
    }
  }

  const creating = mode === "sign-up";

  return (
    <div className="mx-auto max-w-sm p-6">
      <h1 className="mb-4 text-2xl font-semibold">{creating ? "Create account" : "Sign in"}</h1>
      <form onSubmit={onSubmit} className="flex flex-col gap-3">
        <label className="flex flex-col gap-1">
          <span>Email</span>
          <input
            type="email"
            name="email"
            autoComplete="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="rounded border border-neutral-400 bg-white px-2 py-1 text-neutral-900"
          />
        </label>
        <label className="flex flex-col gap-1">
          <span>Password</span>
          <input
            type="password"
            name="password"
            autoComplete={creating ? "new-password" : "current-password"}
            required
            minLength={6}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="rounded border border-neutral-400 bg-white px-2 py-1 text-neutral-900"
          />
        </label>
        <button
          type="submit"
          disabled={busy}
          className="rounded bg-neutral-900 px-3 py-2 text-white disabled:bg-neutral-300 disabled:text-neutral-600"
        >
          {creating ? "Create account" : "Sign in"}
        </button>
        <div role="alert" className="min-h-6 text-red-700">
          {error || authError}
        </div>
      </form>
      <button
        type="button"
        onClick={() => {
          setMode(creating ? "sign-in" : "sign-up");
          setError("");
        }}
        className="mt-2 underline"
      >
        {creating ? "I have an account. Sign in." : "New here? Create account"}
      </button>
    </div>
  );
}
