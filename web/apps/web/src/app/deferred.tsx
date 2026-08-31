import { type ComponentType, type ReactNode, useSyncExternalStore } from "react";

// A chunk that the shell does not need for the first paint loads on its
// own (D-320). React holds a Suspense fallback on screen for 300 ms after
// the fallback commits, and it holds every later reveal with it, so a
// chunk that is already in hand still costs the reader a third of a
// second. Measured on 2026-08-30: content at 347 ms, with every file in
// the browser by 42 ms.
//
// A deferred unit therefore uses no Suspense. It starts its download and
// it mounts the component the moment the code is here. Nothing suspends,
// so nothing is throttled.
export type Deferred<P> = {
  preload: () => void;
  // Mount renders the component once the code is here, and the fallback
  // until then. It is one component for the life of the page.
  Mount: (props: P & { fallback?: ReactNode }) => ReactNode;
  // useError gives the failure of the download, for a caller that must
  // show it. A caller that can wait ignores it and shows the fallback.
  useError: () => unknown;
};

export function deferred<P extends object>(load: () => Promise<{ default: ComponentType<P> }>): Deferred<P> {
  let component: ComponentType<P> | null = null;
  let failure: unknown = null;
  let started = false;
  const listeners = new Set<() => void>();

  function announce() {
    for (const listener of listeners) listener();
  }

  function preload() {
    if (started) return;
    started = true;
    void load().then(
      (mod) => {
        component = mod.default;
        announce();
      },
      (err: unknown) => {
        // A chunk that is gone after a new release never arrives, and a
        // retry in the same document asks the same dead address. The
        // caller shows the failure, and the way out is a new document.
        failure = err;
        announce();
      },
    );
  }

  function subscribe(listener: () => void) {
    listeners.add(listener);
    return () => listeners.delete(listener);
  }

  function Mount({ fallback = null, ...rest }: P & { fallback?: ReactNode }) {
    const Loaded = useSyncExternalStore(
      subscribe,
      () => component,
      () => null,
    );
    // The component comes from a module-level cache, so it is the same
    // component on every render and no state of it resets.
    return Loaded ? <Loaded {...(rest as unknown as P)} /> : fallback;
  }

  return {
    preload,
    Mount,
    useError: () =>
      useSyncExternalStore(
        subscribe,
        () => failure,
        () => null,
      ),
  };
}
