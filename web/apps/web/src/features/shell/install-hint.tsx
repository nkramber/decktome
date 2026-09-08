import { XIcon } from "lucide-react";
import { useEffect, useState } from "react";

import { Button } from "../../components/ui/button";
import { useAppStore } from "../../lib/store";

// The install hint of PR-25, Stage A. It shows once, on a phone, after
// the reader has a deck, and never again after a dismissal.
//
// Two paths. A browser that fires `beforeinstallprompt` installs on a
// tap, which is Chrome on Android. iOS fires no such event, so the hint
// names the two taps of the share sheet instead.

// BeforeInstallPromptEvent is the Chrome event. No lib.dom type covers
// it, so this names the two members the hint reads.
type BeforeInstallPromptEvent = Event & {
  prompt: () => Promise<void>;
  userChoice: Promise<{ outcome: "accepted" | "dismissed" }>;
};

// onPhone reports a narrow screen with a coarse pointer. The hint is
// about a Home Screen, so a desktop browser never reads it.
function onPhone(): boolean {
  if (typeof window === "undefined" || !window.matchMedia) return false;
  return window.matchMedia("(max-width: 640px) and (pointer: coarse)").matches;
}

// installed reports a window the reader already added. A standalone
// display mode is the web answer, and `navigator.standalone` is the iOS
// one.
function installed(): boolean {
  if (typeof window === "undefined") return false;
  const iosStandalone = (window.navigator as Navigator & { standalone?: boolean }).standalone === true;
  return iosStandalone || (window.matchMedia?.("(display-mode: standalone)").matches ?? false);
}

export function InstallHint() {
  const hadDecks = useAppStore((s) => s.hadDecks);
  const dismissed = useAppStore((s) => s.installHintDismissed);
  const dismiss = useAppStore((s) => s.dismissInstallHint);
  const [prompt, setPrompt] = useState<BeforeInstallPromptEvent | null>(null);

  useEffect(() => {
    const onPromptEvent = (e: Event) => {
      // The browser shows a bar of its own without this, and the hint is
      // the one place the app asks.
      e.preventDefault();
      setPrompt(e as BeforeInstallPromptEvent);
    };
    window.addEventListener("beforeinstallprompt", onPromptEvent);
    return () => window.removeEventListener("beforeinstallprompt", onPromptEvent);
  }, []);

  if (!hadDecks || dismissed || !onPhone() || installed()) return null;

  return (
    <aside
      aria-label="Add Deck Tome to your Home Screen"
      className="flex shrink-0 items-start gap-3 border-t border-border bg-muted px-4 py-3 text-sm print:hidden"
    >
      <div className="flex grow flex-col gap-2">
        <p className="font-display">Keep Deck Tome on your Home Screen.</p>
        {prompt ? (
          <Button
            type="button"
            size="sm"
            className="self-start"
            onClick={() => {
              void prompt.prompt();
              dismiss();
            }}
          >
            Add to Home Screen
          </Button>
        ) : (
          <p className="text-muted-foreground">Tap Share, then Add to Home Screen. It opens like any other app, and it starts with no signal.</p>
        )}
      </div>
      <Button type="button" variant="ghost" size="icon" aria-label="Dismiss the install hint" onClick={dismiss}>
        <XIcon className="size-4" aria-hidden="true" />
      </Button>
    </aside>
  );
}
