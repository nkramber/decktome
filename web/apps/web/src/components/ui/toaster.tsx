import { Toaster as Sonner } from "sonner";

// One toast host for the app. Every mutation reports its result here
// (D-311). Dark is the only theme (D-330).
export function Toaster() {
  return (
    <Sonner
      theme="dark"
      position="bottom-right"
      toastOptions={{
        classNames: {
          toast: "rounded-card border border-border bg-card text-card-foreground",
          description: "text-muted-foreground",
        },
      }}
    />
  );
}

export { toast } from "sonner";
