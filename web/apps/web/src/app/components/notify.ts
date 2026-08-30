// notify reports the result of a mutation as a toast (D-311). The toast
// module loads on the first call, so it stays off the first paint (D-320).
type Kind = "success" | "error" | "info";

export async function notify(kind: Kind, title: string, description?: string) {
  const { toast } = await import("../../components/ui/toaster");
  const options = description ? { description } : undefined;
  if (kind === "success") toast.success(title, options);
  else if (kind === "error") toast.error(title, options);
  else toast(title, options);
}
