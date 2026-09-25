// DataCredit names the sources of the data on a deck page (guardrail 7).
// The TopDeck.gg terms ask for a visible credit with a link, and the
// build reads its tournament data (D-417, REV-028).
export const topdeckCredit = "Tournament data by TopDeck.gg";
export const topdeckURL = "https://topdeck.gg";

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
