import * as DialogPrimitive from "@radix-ui/react-dialog";
import { type ReactNode, useState, useSyncExternalStore } from "react";

import type { Face } from "./card-tile";

// A touch screen opens the large view with a tap on the art (D-1049). A
// mouse keeps the art as an image, and the name opens the card detail.
export const coarsePointerQuery = "(pointer: coarse)";

function subscribe(onChange: () => void): () => void {
  if (typeof window === "undefined" || !window.matchMedia) return () => {};
  const list = window.matchMedia(coarsePointerQuery);
  list.addEventListener("change", onChange);
  return () => list.removeEventListener("change", onChange);
}

function coarsePointer(): boolean {
  if (typeof window === "undefined" || !window.matchMedia) return false;
  return window.matchMedia(coarsePointerQuery).matches;
}

// ZoomFace wraps the image of one face. On a touch screen the image is a
// button, and a tap shows the large image over the whole screen. A tap
// beside the image closes the view, and so does the Escape key.
//
// The large image is 672 by 936. The view never draws it larger than the
// file, so the rules text stays sharp (D-601).
export function ZoomFace({ face, children }: { face: Face; children: ReactNode }) {
  const coarse = useSyncExternalStore(subscribe, coarsePointer, () => false);
  const [open, setOpen] = useState(false);
  const src = face.imageUris?.large || face.imageUris?.normal || "";
  if (!coarse || !src) return children;
  return (
    <DialogPrimitive.Root open={open} onOpenChange={setOpen}>
      <DialogPrimitive.Trigger asChild>
        <button type="button" aria-label={`Show ${face.name} larger`} className="block w-full rounded" data-testid="card-zoom-open">
          {children}
        </button>
      </DialogPrimitive.Trigger>
      <DialogPrimitive.Portal>
        <DialogPrimitive.Overlay className="fixed inset-0 z-50 bg-black/85 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=closed]:animate-out data-[state=closed]:fade-out-0" />
        <DialogPrimitive.Content
          aria-describedby={undefined}
          data-testid="card-zoom"
          className="fixed inset-0 z-50 flex items-center justify-center p-4 pt-[calc(1rem+var(--edge-top,0px))] pb-[calc(1rem+var(--edge-bottom,0px))]"
          // The content box fills the screen, so a tap beside the image
          // lands on the box itself and not on the overlay.
          onClick={(e) => {
            if (e.target === e.currentTarget) setOpen(false);
          }}
        >
          <DialogPrimitive.Title className="sr-only">{face.name}</DialogPrimitive.Title>
          <img src={src} alt={`${face.name} (card, large)`} width={672} height={936} className="h-auto max-h-full w-auto max-w-full rounded-[4.75%/3.5%] object-contain" />
        </DialogPrimitive.Content>
      </DialogPrimitive.Portal>
    </DialogPrimitive.Root>
  );
}
