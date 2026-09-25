// DataCredit names the sources of the data on a deck page (guardrail 7).
// The TopDeck.gg terms ask for a visible credit with a link, and the
// build reads its tournament data (D-417, REV-028).
export const topdeckCredit = "Tournament data by TopDeck.gg";
export const topdeckURL = "https://topdeck.gg";

// ArtCredit names the artist of each art crop on a page. The Scryfall
// guidelines ask for the artist and the copyright beside an art crop,
// and the full card image carries its own (D-291, guardrail 7, REV-042).
export function ArtCredit({ artists }: { artists: string[] }) {
  const names = [...new Set(artists.filter(Boolean))];
  if (names.length === 0) return null;
  return <p className="text-[10px] text-muted-foreground">Art: {names.join(", ")}. ™ & © Wizards of the Coast.</p>;
}

export function DataCredit() {
  return (
    <p className="text-xs text-muted-foreground">
      Card images and card text are unofficial Fan Content permitted under the Wizards of the Coast Fan Content Policy. They are
      copyright Wizards of the Coast, LLC, and come from Scryfall.{" "}
      <a href={topdeckURL} target="_blank" rel="noreferrer" className="underline underline-offset-4 hover:no-underline">
        {topdeckCredit}
      </a>
      .
    </p>
  );
}
