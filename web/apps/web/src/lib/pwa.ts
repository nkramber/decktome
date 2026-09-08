// The web manifest of PR-25, Stage A. It lives here and not in the Vite
// config so a test reads it: "Lighthouse reports the app installable" is
// a gate, and the fields it rests on are checked for nothing.
//
// The rules a browser applies to install a web app: a name, a start url,
// a display mode away from `browser`, and an icon of 192 pixels or more.
// Android crops to a circle, so one icon is maskable.

// pageBackground is the --background token of styles/tokens.css. The
// browser paints it behind the splash screen and in the status bar of a
// standalone window.
export const pageBackground = "#0a0c16";

export const webManifest = {
  name: "Deck Tome",
  short_name: "Deck Tome",
  description: "An agentic Magic: The Gathering deck builder that reads the cards you own.",
  start_url: "/",
  scope: "/",
  display: "standalone",
  orientation: "portrait",
  theme_color: pageBackground,
  background_color: pageBackground,
  icons: [
    { src: "/icon-192.png", sizes: "192x192", type: "image/png" },
    { src: "/icon-512.png", sizes: "512x512", type: "image/png" },
    { src: "/icon-maskable-512.png", sizes: "512x512", type: "image/png", purpose: "maskable" },
  ],
} as const;
