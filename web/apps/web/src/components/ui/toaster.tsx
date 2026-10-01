import { Toaster as Sonner } from "sonner";

// One toast host for the app. Every mutation reports its result here
// (D-311). Dark is the only theme (D-330).
export function Toaster() {
  return (
    <Sonner
      theme="dark"
      position="bottom-right"
      // The installed app adds its bottom buffer to the default offsets
      // of sonner, 24px and 16px on a phone (F-187, D-1010).
      offset={{ bottom: "calc(24px + var(--edge-bottom))" }}
      mobileOffset={{ bottom: "calc(16px + var(--edge-bottom))" }}
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
