import { Code, ConnectError } from "@connectrpc/connect";
import { useQuery } from "@tanstack/react-query";
import { ArrowLeftIcon, MessageSquareIcon, PencilIcon, StarIcon, Trash2Icon } from "lucide-react";
import { type FormEvent, useState } from "react";
import { Link, useNavigate, useParams } from "react-router";

import { ErrorState } from "../../app/components/error-state";
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
import { Skeleton } from "../../components/ui/skeleton";
import { deckClient } from "../../lib/api";
import { cn } from "../../lib/cn";
import { errorMessage } from "../../lib/errors";
import { DeckView } from "./deck-view";
import { useDeckWrites } from "./use-decks";

// The page of one deck (PR-17). It holds the deck view and the actions a
// user owns: rename, favorite, and delete. The agent owns everything else.
export function DeckPage() {
  const { id = "" } = useParams();
  const navigate = useNavigate();
  const { rename, setFavorite, remove } = useDeckWrites();

  const deckQuery = useQuery({
    queryKey: ["deck", id],
    queryFn: () => deckClient.getDeck({ deckId: id }),
    enabled: id !== "",
  });
  const deck = deckQuery.data?.deck;

  const [renameOpen, setRenameOpen] = useState(false);
  const [draftName, setDraftName] = useState("");

  if (deckQuery.isPending) {
    return (
      <div className="mx-auto flex w-full max-w-7xl flex-col gap-4 p-4 md:p-6" role="status">
        <span className="sr-only">Loading the deck...</span>
        <Skeleton className="h-9 w-40" />
        <Skeleton className="h-32 w-full rounded-panel" />
        <Skeleton className="h-64 w-full rounded-panel" />
      </div>
    );
  }

  if (deckQuery.isError || !deck) {
    const gone = deckQuery.isError && ConnectError.from(deckQuery.error).code === Code.NotFound;
    return (
      <div className="mx-auto flex w-full max-w-3xl flex-col gap-4 p-4 md:p-6">
        <ErrorState
          title={gone ? "That deck is gone" : "Could not load the deck"}
          message={gone ? "Someone deleted this deck, or the link is wrong." : deckQuery.isError ? errorMessage(deckQuery.error) : "The server returned no deck."}
          onRetry={gone ? undefined : () => void deckQuery.refetch()}
        />
        <Button asChild variant="outline" className="self-center">
          <Link to="/decks">Back to your decks</Link>
        </Button>
      </div>
    );
  }

  const title = deck.name || "Untitled deck";

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
    const next = !deck?.favorite;
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
    <div className="mx-auto flex w-full max-w-7xl flex-col gap-5 p-4 md:p-6">
      <div className="flex flex-wrap items-center gap-2">
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

        {deck.sessionId && (
          <Button asChild variant="outline" size="sm">
            <Link to={`/session/${deck.sessionId}`}>
              <MessageSquareIcon aria-hidden="true" />
              Open the chat
            </Link>
          </Button>
        )}

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
              <AlertDialogDescription>The deck goes for good. The chat that built it stays, and you can build again from it.</AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>Keep the deck</AlertDialogCancel>
              <AlertDialogAction onClick={() => void onDelete()}>Delete the deck</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>

      <DeckView deck={deck} />
    </div>
  );
}
