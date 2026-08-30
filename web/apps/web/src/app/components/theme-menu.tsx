import { MonitorIcon, MoonIcon, SunIcon } from "lucide-react";
import { lazy, Suspense, useState } from "react";

import { Button } from "../../components/ui/button";
import { useThemeStore } from "../../lib/theme";

// The menu itself loads on the first open (D-320). The trigger below is
// the same element in both states, so nothing moves when it arrives.
const ThemeMenuContent = lazy(async () => ({ default: (await import("./shell-menus")).ThemeMenuContent }));

function ThemeTrigger({ onClick }: { onClick?: () => void }) {
  const choice = useThemeStore((s) => s.choice);
  const theme = useThemeStore((s) => s.theme);
  const Icon = choice === "system" ? MonitorIcon : theme === "dark" ? MoonIcon : SunIcon;
  return (
    <Button variant="ghost" className="w-full justify-start" aria-haspopup="menu" aria-label={`Theme: ${choice}`} onClick={onClick}>
      <Icon aria-hidden="true" />
      Theme
    </Button>
  );
}

export function ThemeMenu() {
  const [opened, setOpened] = useState(false);
  if (!opened) return <ThemeTrigger onClick={() => setOpened(true)} />;
  return (
    <Suspense fallback={<ThemeTrigger />}>
      <ThemeMenuContent trigger={<ThemeTrigger />} />
    </Suspense>
  );
}
