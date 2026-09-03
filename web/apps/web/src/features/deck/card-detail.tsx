import { type Card, LegalityStatus, type Printing } from "@mtg/api-client/mtg/v1/card_pb";
import type { DeckCard } from "@mtg/api-client/mtg/v1/deck_pb";
import { useQuery } from "@tanstack/react-query";

import { Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle } from "../../components/ui/sheet";
import { cardClient } from "../../lib/api";
import { errorMessage } from "../../lib/errors";
import { FaceImage, facesOf } from "./card-tile";
import { priceText } from "./deck-stats";

// The card detail (PR-20, D-318): a Sheet with the full image and both
// faces, the Oracle text, the type line, the mana cost, the rulings with
// their dates, the legalities, the printings with prices, the deck's
// reason line, and a link to Scryfall. The rulings and the printings
// load when the sheet opens and never before.

// legalityFormats are the formats the table shows, in this order, when
// the card data names them. The app builds the first three.
const legalityFormats: { key: string; name: string }[] = [
  { key: "commander", name: "Commander" },
  { key: "standard", name: "Standard" },
  { key: "modern", name: "Modern" },
  { key: "pioneer", name: "Pioneer" },
  { key: "legacy", name: "Legacy" },
  { key: "vintage", name: "Vintage" },
  { key: "pauper", name: "Pauper" },
  { key: "brawl", name: "Brawl" },
];

export function legalityLabel(status: LegalityStatus | undefined): string {
  switch (status) {
    case LegalityStatus.LEGAL:
      return "Legal";
    case LegalityStatus.NOT_LEGAL:
      return "Not legal";
    case LegalityStatus.BANNED:
      return "Banned";
    case LegalityStatus.RESTRICTED:
      return "Restricted";
    default:
      return "Unknown";
  }
}

// scryfallURL is the card page of a printing. Scryfall answers the set
// and the collector number with a redirect to the full page (read
// 2026-09-03).
export function scryfallURL(printing: Printing | undefined): string {
  if (!printing?.setCode || !printing.collectorNumber) return "";
  return `https://scryfall.com/card/${encodeURIComponent(printing.setCode)}/${encodeURIComponent(printing.collectorNumber)}`;
}

export function CardDetail({ entry, card, open, onOpenChange }: { entry: DeckCard | undefined; card: Card | undefined; open: boolean; onOpenChange: (open: boolean) => void }) {
  const oracleId = entry?.oracleId ?? "";
  const rulings = useQuery({
    queryKey: ["rulings", oracleId],
    queryFn: () => cardClient.getRulings({ oracleId }),
    enabled: open && oracleId !== "",
    staleTime: Infinity,
  });
  const printings = useQuery({
    queryKey: ["printings", oracleId],
    queryFn: () => cardClient.getPrintings({ oracleId }),
    enabled: open && oracleId !== "",
    staleTime: Infinity,
  });
  const faces = facesOf(card, entry?.ownedPrinting);
  const name = card?.name || entry?.name || "Card";
  const link = scryfallURL(entry?.ownedPrinting ?? card?.defaultPrinting);
  const legalities = legalityFormats.filter((f) => card?.legalities && f.key in card.legalities);
  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent aria-describedby={`card-detail-type-${oracleId}`}>
        {entry && (
          <>
            <SheetHeader>
              <SheetTitle>{name}</SheetTitle>
              <SheetDescription id={`card-detail-type-${oracleId}`}>{card?.typeLine || "Card detail"}</SheetDescription>
            </SheetHeader>
            {faces.length === 0 && <p className="text-sm text-muted-foreground">No card data for this entry.</p>}
            {faces.map((face, i) => (
              <section key={`${oracleId}-${i}`} aria-label={faces.length > 1 ? `${face.name} (face ${i + 1} of ${faces.length})` : face.name} className="grid gap-3 sm:grid-cols-[minmax(0,12rem)_1fr]">
                <FaceImage face={face} />
                <div className="flex flex-col gap-2 text-sm">
                  {faces.length > 1 && <h3 className="font-medium">{face.name}</h3>}
                  {face.manaCost && (
                    <p>
                      <span className="text-muted-foreground">Mana cost </span>
                      <span className="font-mono">{face.manaCost}</span>
                    </p>
                  )}
                  <p className="text-muted-foreground">{face.typeLine}</p>
                  {face.oracleText && <p className="whitespace-pre-line leading-relaxed">{face.oracleText}</p>}
                </div>
              </section>
            ))}
            {entry.reason && (
              <p className="rounded-card border border-accent/40 bg-accent/5 p-3 text-sm" data-testid="detail-reason">
                <span className="font-medium">In this deck: </span>
                {entry.reason}
              </p>
            )}
            {legalities.length > 0 && (
              <table className="w-full text-sm">
                <caption className="mb-2 border-b border-border pb-1 text-left text-xs font-semibold tracking-wide uppercase">Legalities</caption>
                <thead className="sr-only">
                  <tr>
                    <th scope="col">Format</th>
                    <th scope="col">Status</th>
                  </tr>
                </thead>
                <tbody>
                  {legalities.map((f) => (
                    <tr key={f.key}>
                      <th scope="row" className="py-0.5 pr-3 text-left font-normal text-muted-foreground">
                        {f.name}
                      </th>
                      <td className="py-0.5">{legalityLabel(card?.legalities[f.key])}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            )}
            <section aria-labelledby={`rulings-title-${oracleId}`} className="text-sm">
              <h3 id={`rulings-title-${oracleId}`} className="mb-2 border-b border-border pb-1 text-xs font-semibold tracking-wide uppercase">
                Rulings
              </h3>
              {rulings.isPending && <p role="status">Loading rulings...</p>}
              {rulings.isError && (
                <p role="alert" className="text-danger">
                  Could not load the rulings: {errorMessage(rulings.error)}
                </p>
              )}
              {rulings.data && !rulings.data.hasRulings && <p className="text-muted-foreground">The card data holds no rulings file yet.</p>}
              {rulings.data?.hasRulings && rulings.data.rulings.length === 0 && <p className="text-muted-foreground">No rulings.</p>}
              {rulings.data && rulings.data.rulings.length > 0 && (
                <ul className="flex flex-col gap-2" data-testid="rulings">
                  {rulings.data.rulings.map((r, i) => (
                    <li key={i} className="flex flex-col gap-0.5">
                      <time dateTime={r.publishedAt} className="text-xs text-muted-foreground">
                        {r.publishedAt}
                      </time>
                      <span className="leading-relaxed">{r.comment}</span>
                    </li>
                  ))}
                </ul>
              )}
              {rulings.data?.asOf && (
                <p className="mt-2 text-xs text-muted-foreground" data-testid="rulings-as-of">
                  Rulings from the card data of {rulings.data.asOf}.
                </p>
              )}
            </section>
            <section aria-labelledby={`printings-title-${oracleId}`} className="text-sm">
              <h3 id={`printings-title-${oracleId}`} className="mb-2 border-b border-border pb-1 text-xs font-semibold tracking-wide uppercase">
                Printings
              </h3>
              {printings.isPending && <p role="status">Loading printings...</p>}
              {printings.isError && (
                <p role="alert" className="text-danger">
                  Could not load the printings: {errorMessage(printings.error)}
                </p>
              )}
              {printings.data && printings.data.printings.length === 0 && <p className="text-muted-foreground">No printing is known.</p>}
              {printings.data && printings.data.printings.length > 0 && (
                <table className="w-full text-sm">
                  <caption className="sr-only">Printings with prices of {printings.data.priceAsOf}</caption>
                  <thead className="text-left text-xs text-muted-foreground">
                    <tr>
                      <th scope="col" className="py-0.5 pr-3 font-normal">
                        Set
                      </th>
                      <th scope="col" className="py-0.5 pr-3 font-normal">
                        Number
                      </th>
                      <th scope="col" className="py-0.5 pr-3 font-normal">
                        Rarity
                      </th>
                      <th scope="col" className="py-0.5 text-right font-normal">
                        Price
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    {printings.data.printings.map((p) => (
                      <tr key={p.scryfallId} data-testid="printing-row">
                        <td className="py-0.5 pr-3">
                          {p.setName || p.setCode.toUpperCase()}
                          {p.digital && <span className="ml-1 text-xs text-muted-foreground">(digital)</span>}
                        </td>
                        <td className="py-0.5 pr-3 tabular-nums">{p.collectorNumber}</td>
                        <td className="py-0.5 pr-3 capitalize">{p.rarity}</td>
                        <td className="py-0.5 text-right tabular-nums">{priceText(p.priceUsd)}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              )}
              {printings.data?.priceAsOf && <p className="mt-2 text-xs text-muted-foreground">Prices of {printings.data.priceAsOf}.</p>}
            </section>
            {link && (
              <p className="text-sm">
                <a href={link} target="_blank" rel="noreferrer" className="text-primary underline underline-offset-4 hover:no-underline">
                  Open on Scryfall
                </a>
              </p>
            )}
          </>
        )}
      </SheetContent>
    </Sheet>
  );
}
