import { Toaster as Sonner } from "sonner";

import { useThemeStore } from "../../lib/theme";

// One toast host for the app. Every mutation reports its result here
// (D-311). The theme follows the resolved theme of the store.
export function Toaster() {
  const theme = useThemeStore((s) => s.theme);
  return (
    <Sonner
      theme={theme}
      position="bottom-right"
      toastOptions={{
        classNames: {
          toast: "rounded-md border border-border bg-surface text-surface-foreground",
          description: "text-muted-foreground",
        },
      }}
    />
  );
}

export { toast } from "sonner";
