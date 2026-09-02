import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { PencilIcon, Trash2Icon } from "lucide-react";
import { useState } from "react";
import { Link } from "react-router";

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
import { Input } from "../../components/ui/input";
import { Label } from "../../components/ui/label";
import { agentClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";

// The unfinished chats sit under the message box of a new chat (D-433,
// D-438): the conversations with no deck yet, with resume, rename, and
// delete. A finished chat lives on its deck. With none to show, the
// section does not render at all: an empty frame under the box says
// nothing a reader needs.

export function UnfinishedChats() {
  const queryClient = useQueryClient();
  const list = useQuery({
    queryKey: ["sessions"],
    queryFn: () => agentClient.listSessions({ pageSize: 50 }),
  });
  const unfinished = (list.data?.sessions ?? []).filter((s) => s.deckCount === 0);
  const [renaming, setRenaming] = useState("");
  const [renameTo, setRenameTo] = useState("");
  const rename = useMutation({
    mutationFn: (v: { id: string; name: string }) => agentClient.updateSession({ sessionId: v.id, name: v.name }),
    onSuccess: () => {
      setRenaming("");
      void queryClient.invalidateQueries({ queryKey: ["sessions"] });
    },
    onError: (err) => void notify("error", "Could not rename the chat", errorMessage(err)),
  });
  const remove = useMutation({
    mutationFn: (id: string) => agentClient.deleteSession({ sessionId: id }),
    onSuccess: () => {
      void notify("success", "Chat deleted", "");
      void queryClient.invalidateQueries({ queryKey: ["sessions"] });
    },
    onError: (err) => void notify("error", "Could not delete the chat", errorMessage(err)),
  });

  // Nothing renders while the list loads or when it holds no chat. A
  // failed list still shows, because a silent failure hides a defect.
  if (list.isPending || (list.isSuccess && unfinished.length === 0)) return null;

  return (
    <section aria-labelledby="unfinished-title" className="flex flex-col gap-3">
      <h2 id="unfinished-title" className="font-display text-lg font-semibold">
        Unfinished chats
      </h2>
      {list.isError && <ErrorState title="Could not list the chats" message={errorMessage(list.error)} onRetry={() => void list.refetch()} />}
      {unfinished.length > 0 && (
        <ul className="flex flex-col gap-2" data-testid="unfinished-chats">
          {unfinished.map((s) => {
            const title = s.name || s.firstMessage || "Untitled chat";
            return (
              <li key={s.id} className="flex flex-wrap items-center gap-3 rounded-card border border-border bg-card px-3 py-2">
                <Link to={`/session/${s.id}`} className="min-w-0 grow truncate font-medium hover:underline">
                  {title}
                </Link>
                <span className="font-mono text-[11px] text-muted-foreground">
                  {s.updatedAt?.seconds ? new Date(Number(s.updatedAt.seconds) * 1000).toLocaleDateString() : ""}
                  {s.usage?.priced ? ` · $${s.usage.costUsd.toFixed(4)}` : ""}
                </span>
                <Button
                  variant="ghost"
                  size="icon"
                  aria-label={`Rename ${title}`}
                  className="size-7"
                  onClick={() => {
                    setRenaming(s.id);
                    setRenameTo(s.name || s.firstMessage);
                  }}
                >
                  <PencilIcon aria-hidden="true" />
                </Button>
                <AlertDialog>
                  <AlertDialogTrigger asChild>
                    <Button variant="ghost" size="icon" aria-label={`Delete ${title}`} className="size-7 text-danger hover:text-danger">
                      <Trash2Icon aria-hidden="true" />
                    </Button>
                  </AlertDialogTrigger>
                  <AlertDialogContent>
                    <AlertDialogHeader>
                      <AlertDialogTitle>Delete this chat?</AlertDialogTitle>
                      <AlertDialogDescription>This can not be undone. The chat built no deck, so nothing else changes.</AlertDialogDescription>
                    </AlertDialogHeader>
                    <AlertDialogFooter>
                      <AlertDialogCancel>Keep it</AlertDialogCancel>
                      <AlertDialogAction onClick={() => remove.mutate(s.id)}>Delete the chat</AlertDialogAction>
                    </AlertDialogFooter>
                  </AlertDialogContent>
                </AlertDialog>
                {renaming === s.id && (
                  <form
                    className="flex w-full flex-wrap items-center gap-2"
                    onSubmit={(e) => {
                      e.preventDefault();
                      const next = renameTo.trim();
                      if (next !== "") rename.mutate({ id: s.id, name: next });
                    }}
                  >
                    <Label htmlFor={`rename-${s.id}`} className="sr-only">
                      New name for {title}
                    </Label>
                    <Input id={`rename-${s.id}`} ref={(el) => el?.focus()} value={renameTo} onChange={(e) => setRenameTo(e.target.value)} className="h-8 max-w-64" />
                    <Button type="submit" size="sm" disabled={renameTo.trim() === "" || rename.isPending}>
                      Save
                    </Button>
                    <Button type="button" size="sm" variant="outline" onClick={() => setRenaming("")}>
                      Cancel
                    </Button>
                  </form>
                )}
              </li>
            );
          })}
        </ul>
      )}
    </section>
  );
}
