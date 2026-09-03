import type { Deck } from "@mtg/api-client/mtg/v1/deck_pb";
import { ArrowLeftIcon, CopyIcon, PencilIcon, Share2Icon, StarIcon, Trash2Icon } from "lucide-react";
import { type FormEvent, useState } from "react";
import { Link, useNavigate } from "react-router";

import { notify } from "../../app/components/notify";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "../../components/ui/alert-dialog";
import { Button } from "../../components/ui/button";
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "../../components/ui/dialog";
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { useDeckWrites } from "../deck/use-decks";

// The actions a user owns over a deck (D-335): the way back, the
// favorite mark, the name, the share link (D-315), and the delete. Each
// one keeps its own state, so the screen around it needs to know none of
// it.
export function DeckActions({ deck }: { deck: Deck }) {
  const id = deck.id;
  const navigate = useNavigate();
  const { rename, setFavorite, remove, share, revokeShare } = useDeckWrites();
  const [renameOpen, setRenameOpen] = useState(false);
  const [draftName, setDraftName] = useState("");
  const [shareOpen, setShareOpen] = useState(false);
  // The link shows once, after the share, because the store keeps a
  // hash of the token and never the token (D-315).
  const [link, setLink] = useState("");
  const title = deck.name || "Untitled deck";

  async function onShare() {
    try {
      const res = await share.mutateAsync({ deckId: id });
      setLink(`${window.location.origin}/d/${res.token}`);
    } catch (err) {
      await notify("error", "Could not make the link", errorMessage(err));
    }
  }

  async function onCopyLink() {
    try {
      await navigator.clipboard.writeText(link);
      await notify("success", "Link copied", "Anyone who holds it can read the deck.");
    } catch (err) {
      await notify("error", "Could not copy the link", errorMessage(err));
    }
  }

  async function onRevoke() {
    try {
      await revokeShare.mutateAsync({ deckId: id });
      setLink("");
      await notify("success", "Link revoked", "The old link opens nothing now.");
    } catch (err) {
      await notify("error", "Could not revoke the link", errorMessage(err));
    }
  }

  async function onRename(e: FormEvent) {
    e.preventDefault();
    const name = draftName.trim();
    if (!name) return;
    try {
      await rename.mutateAsync({ deckId: id, name });
      setRenameOpen(false);
      await notify("success", "Deck renamed", name);
    } catch (err) {
      await notify("error", "Could not rename the deck", errorMessage(err));
    }
  }

  async function onFavorite() {
    const next = !deck.favorite;
    try {
      await setFavorite.mutateAsync({ deckId: id, favorite: next });
    } catch (err) {
      await notify("error", next ? "Could not add the favorite" : "Could not remove the favorite", errorMessage(err));
    }
  }

  async function onDelete() {
    try {
      await remove.mutateAsync({ deckId: id });
      await notify("success", "Deck deleted", title);
      void navigate("/decks");
    } catch (err) {
      await notify("error", "Could not delete the deck", errorMessage(err));
    }
  }

  return (
<div className="mb-5 flex flex-wrap items-center gap-2">
        <Button asChild variant="ghost" size="sm">
          <Link to="/decks">
            <ArrowLeftIcon aria-hidden="true" />
            Your decks
          </Link>
        </Button>
        <span className="grow" />

        <Button variant="outline" size="sm" aria-pressed={deck.favorite} aria-label={deck.favorite ? "Remove from your favorites" : "Add to your favorites"} onClick={() => void onFavorite()}>
          <StarIcon className={cn("size-4", deck.favorite && "fill-warning text-warning")} aria-hidden="true" />
          {deck.favorite ? "Favorite" : "Add favorite"}
        </Button>

        <Dialog
          open={renameOpen}
          onOpenChange={(open) => {
            setRenameOpen(open);
            if (open) setDraftName(deck.name);
          }}
        >
          <DialogTrigger asChild>
            <Button variant="outline" size="sm">
              <PencilIcon aria-hidden="true" />
              Rename
            </Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Rename the deck</DialogTitle>
              <DialogDescription>The name is yours. It changes nothing the agent built.</DialogDescription>
            </DialogHeader>
            <form onSubmit={onRename} className="flex flex-col gap-4">
              <div className="flex flex-col gap-1.5">
                <Label htmlFor="deck-name">Deck name</Label>
                <Input id="deck-name" value={draftName} onChange={(e) => setDraftName(e.target.value)} maxLength={200} />
              </div>
              <DialogFooter>
                <DialogClose asChild>
                  <Button type="button" variant="outline">
                    Cancel
                  </Button>
                </DialogClose>
                <Button type="submit" disabled={!draftName.trim() || rename.isPending}>
                  Save the name
                </Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>

        <Dialog
          open={shareOpen}
          onOpenChange={(open) => {
            setShareOpen(open);
            if (!open) setLink("");
          }}
        >
          <DialogTrigger asChild>
            <Button variant="outline" size="sm" aria-pressed={deck.shared}>
              <Share2Icon aria-hidden="true" />
              {deck.shared ? "Shared" : "Share"}
            </Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Share the deck</DialogTitle>
              <DialogDescription>
                Anyone who holds the link reads the deck: the cards, the summary, and nothing about you. A new link replaces the old one.
              </DialogDescription>
            </DialogHeader>
            {link ? (
              <div className="flex flex-col gap-1.5">
                <Label htmlFor="share-link">The link, shown once</Label>
                <div className="flex gap-2">
                  <Input id="share-link" value={link} readOnly onFocus={(e) => e.currentTarget.select()} />
                  <Button type="button" variant="outline" size="icon" aria-label="Copy the link" onClick={() => void onCopyLink()}>
                    <CopyIcon aria-hidden="true" />
                  </Button>
                </div>
              </div>
            ) : (
              <p className="text-sm" data-testid="share-state">
                {deck.shared ? "This deck has a link. Make a new one to see it, and the old one dies." : "This deck has no link yet."}
              </p>
            )}
            <DialogFooter>
              {deck.shared && (
                <Button type="button" variant="ghost" className="text-danger hover:bg-danger/10" disabled={revokeShare.isPending} onClick={() => void onRevoke()}>
                  Revoke the link
                </Button>
              )}
              <Button type="button" disabled={share.isPending} onClick={() => void onShare()}>
                {deck.shared || link ? "Make a new link" : "Make a link"}
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>

        <AlertDialog>
          <AlertDialogTrigger asChild>
            <Button variant="ghost" size="sm" className="text-danger hover:bg-danger/10">
              <Trash2Icon aria-hidden="true" />
              Delete
            </Button>
          </AlertDialogTrigger>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Delete {title}?</AlertDialogTitle>
              <AlertDialogDescription>The deck and the chat that built it go for good.</AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>Keep the deck</AlertDialogCancel>
              <AlertDialogAction onClick={() => void onDelete()}>Delete the deck</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
  );
}
