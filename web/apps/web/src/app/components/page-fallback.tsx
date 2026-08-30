import { Skeleton } from "../../components/ui/skeleton";

// The page of a route loads on its own (D-320). This fills the frame while
// it arrives, so the shell never shows an empty column.
export function PageFallback() {
  return (
    <div className="mx-auto flex w-full max-w-5xl flex-col gap-4 p-4 md:p-6" aria-busy="true">
      <Skeleton className="h-8 w-56" />
      <Skeleton className="h-4 w-80" />
      <Skeleton className="h-40 w-full" />
    </div>
  );
}
