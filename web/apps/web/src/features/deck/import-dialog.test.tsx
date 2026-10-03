import { create } from "@bufbuild/protobuf";
import { Code, ConnectError } from "@connectrpc/connect";
import { UnreadableFileSchema, UnresolvedReason } from "@mtg/api-client/mtg/v1/collection_pb";
import { FeedbackKind, ImportPage } from "@mtg/api-client/mtg/v1/feedback_service_pb";
import { FormatId } from "@mtg/api-client/mtg/v1/format_pb";
import { screen, waitFor, within } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { beforeEach, describe, expect, it, vi } from "vitest";

import { fakeUser, state } from "../../test-auth-state";
import { renderAt } from "../../test-utils";
import { nameFromFile } from "./import-dialog";

vi.mock("firebase/app");
vi.mock("firebase/auth");

const importDeck = vi.fn();
const fetchDeckList = vi.fn();
const getDeck = vi.fn();
const listCollections = vi.fn();
const submitFeedback = vi.fn();
vi.mock("../../lib/api", () => ({
  feedbackClient: { submitFeedback: (...a: unknown[]) => submitFeedback(...a) },
  healthClient: { check: () => Promise.resolve({ status: "ok", version: "test", cardSnapshot: "none" }) },
  collectionClient: { listCollections: (...a: unknown[]) => listCollections(...a) },
  agentClient: { getSession: vi.fn(), importDeck: (...a: unknown[]) => importDeck(...a), fetchDeckList: (...a: unknown[]) => fetchDeckList(...a), readImportBracket: vi.fn() },
  cardClient: { getCards: () => Promise.resolve({ cards: [], missingOracleIds: [] }) },
  deckClient: {
    listDecks: () => Promise.resolve({ decks: [], nextPageToken: "" }),
    getDeck: (...a: unknown[]) => getDeck(...a),
  },
}));

const stored = { id: "d9", sessionId: "", name: "Rats", format: { id: FormatId.MODERN }, cards: [], sideboard: [], upgrades: [], commanderOracleIds: [], commanders: [], imported: true };

beforeEach(() => {
  state.user = fakeUser;
  localStorage.clear();
  importDeck.mockReset();
  fetchDeckList.mockReset();
  getDeck.mockReset();
  listCollections.mockReset();
  submitFeedback.mockReset();
  listCollections.mockResolvedValue({ collections: [] });
  getDeck.mockResolvedValue({ deck: stored });
});

async function openAndPaste(list: string) {
  await renderAt("/decks");
  const user = userEvent.setup();
  await user.click(await screen.findByRole("button", { name: "Import a deck" }));
  await user.type(screen.getByRole("textbox", { name: "Deck list" }), list);
  await user.click(screen.getByRole("button", { name: "Import" }));
  return user;
}

describe("nameFromFile", () => {
  it("drops the extension and reads underscores as spaces", () => {
    expect(nameFromFile("living_weapon.txt")).toBe("living weapon");
    expect(nameFromFile("Turtle_power.txt")).toBe("Turtle power");
  });
});

describe("ImportDialog", () => {
  it("asks for the format of a list that is not Commander, then opens the deck (D-857)", async () => {
    importDeck.mockResolvedValueOnce({ needsFormat: true, commanderOptions: [], unresolved: [] });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    const user = await openAndPaste("4 Lightning Bolt");
    await user.click(await screen.findByRole("radio", { name: "Modern" }));
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenLastCalledWith(expect.objectContaining({ text: "4 Lightning Bolt", format: FormatId.MODERN })));
    await waitFor(() => expect(getDeck).toHaveBeenCalledWith({ deckId: "d9" }));
  });

  // REV-027: a Commander list of 99 or 101 cards reads no format of its
  // own, and the choice offered no Commander.
  it("offers Commander in the format choice, then asks for the leader", async () => {
    importDeck.mockResolvedValueOnce({ needsFormat: true, commanderOptions: [], unresolved: [] });
    importDeck.mockResolvedValueOnce({ needsFormat: false, commanderOptions: [{ oracleId: "o-k", name: "Karlov of the Ghost Council" }], unresolved: [] });
    const user = await openAndPaste("1 Karlov of the Ghost Council");
    await user.click(await screen.findByRole("radio", { name: "Commander" }));
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenLastCalledWith(expect.objectContaining({ format: FormatId.COMMANDER })));
    expect(await screen.findByRole("radio", { name: "Karlov of the Ghost Council" })).toBeInTheDocument();
  });

  it("offers neither as the house format", async () => {
    importDeck.mockResolvedValueOnce({ needsFormat: true, commanderOptions: [], unresolved: [] });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    const user = await openAndPaste("4 Lightning Bolt");
    await user.click(await screen.findByRole("radio", { name: /Neither/ }));
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenLastCalledWith(expect.objectContaining({ format: FormatId.HOUSE })));
  });

  it("asks which card leads a list with no commander mark (D-847)", async () => {
    importDeck.mockResolvedValueOnce({
      needsFormat: false,
      commanderOptions: [
        { oracleId: "o-a", name: "Karlov of the Ghost Council" },
        { oracleId: "o-b", name: "Marrow-Gnawer" },
      ],
      unresolved: [],
    });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    const user = await openAndPaste("1 Karlov of the Ghost Council");
    await user.click(await screen.findByRole("radio", { name: "Marrow-Gnawer" }));
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenLastCalledWith(expect.objectContaining({ commanderOracleIds: ["o-b"] })));
  });

  it("lists each line that matched no card before the deck opens (D-846)", async () => {
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [{ line: 4, raw: "1x No Such Card" }] });
    const user = await openAndPaste("1x No Such Card");
    expect(await screen.findByText("Line 4: 1x No Such Card")).toBeInTheDocument();
    expect(getDeck).not.toHaveBeenCalled();
    await user.click(screen.getByRole("button", { name: "Open the deck" }));
    await waitFor(() => expect(getDeck).toHaveBeenCalledWith({ deckId: "d9" }));
  });

  it("sends the collection the reader picked to show owned cards (D-849)", async () => {
    listCollections.mockResolvedValue({ collections: [{ id: "c1", name: "Binder", cardCount: 900 }] });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    await renderAt("/decks");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Import a deck" }));
    await user.type(screen.getByRole("textbox", { name: "Deck list" }), "4 Lightning Bolt");
    await user.selectOptions(await screen.findByRole("combobox", { name: "Show owned cards from this collection" }), "c1");
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenCalledWith(expect.objectContaining({ collectionId: "c1" })));
  });
});

// A list the app could not read shows the short form of D-882, and a
// list with lines that did not parse offers it beside the skipped lines
// (D-887). The dialog names no service (D-889).
describe("a deck link", () => {
  async function openAndRead(url: string) {
    await renderAt("/decks");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Import a deck" }));
    await user.type(screen.getByRole("textbox", { name: "Deck link" }), url);
    await user.click(screen.getByRole("button", { name: "Read" }));
    return user;
  }

  const moxSteps = ["Open the deck on Moxfield.", "Select More, then Export.", "Select Copy for Arena.", "Paste the list in the box below, then select Import."];

  it("fills the list from a deck the app reads, and sends the link with the import (D-1100, D-1101)", async () => {
    const text = "Commander\n1 Karlov of the Ghost Council (mkm) 1\n\nDeck\n99 Plains (neo) 294\n";
    fetchDeckList.mockResolvedValueOnce({ text, name: "Karlov Lifegain", sourceUrl: "https://archidekt.com/decks/42", leftOut: 2, site: "Archidekt", exportSteps: [], knownSite: true });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    const user = await openAndRead("archidekt.com/decks/42/karlov");

    expect(await screen.findByText(/Read Karlov Lifegain from Archidekt\. 2 cards outside the deck, such as the maybeboard, stay out\./)).toBeInTheDocument();
    expect(fetchDeckList).toHaveBeenCalledWith({ url: "archidekt.com/decks/42/karlov" });
    expect(screen.getByRole("textbox", { name: "Deck list" })).toHaveValue(text);
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() =>
      expect(importDeck).toHaveBeenCalledWith(expect.objectContaining({ text, name: "Karlov Lifegain", sourceUrl: "https://archidekt.com/decks/42" })),
    );
    expect(submitFeedback).not.toHaveBeenCalled();
  });

  it("shows the exact steps of a known site, and files no report (D-1103)", async () => {
    fetchDeckList.mockResolvedValueOnce({ text: "", name: "", sourceUrl: "", leftOut: 0, site: "Moxfield", exportSteps: moxSteps, knownSite: true });
    await openAndRead("https://moxfield.com/decks/abc");

    const steps = await screen.findByTestId("export-steps");
    expect(steps).toHaveTextContent("Moxfield does not let the app read its decks.");
    expect(within(steps).getAllByRole("listitem").map((li) => li.textContent)).toEqual(moxSteps);
    expect(screen.getByRole("textbox", { name: "Deck list" })).toHaveValue("");
    expect(submitFeedback).not.toHaveBeenCalled();
  });

  it("shows the general steps of another site, and files a report of its link (D-1104)", async () => {
    submitFeedback.mockResolvedValueOnce({ feedbackId: "fb1" });
    fetchDeckList.mockResolvedValueOnce({ text: "", name: "", sourceUrl: "", leftOut: 0, site: "tappedout.net", exportSteps: ["Open the deck on its site."], knownSite: false });
    await openAndRead("https://tappedout.net/mtg-decks/x/");

    expect(await screen.findByTestId("export-steps")).toHaveTextContent("The app can not read decks from tappedout.net yet.");
    expect(await screen.findByText("The app sent a report of this site, so that it can be added.")).toBeInTheDocument();
    expect(submitFeedback).toHaveBeenCalledTimes(1);
    const sent = submitFeedback.mock.calls[0]?.[0] as { feedback: { kind: FeedbackKind; importPage: ImportPage; importContent: Uint8Array } };
    expect(sent.feedback.kind).toBe(FeedbackKind.IMPORT);
    expect(sent.feedback.importPage).toBe(ImportPage.DECK_LINK);
    expect(new TextDecoder().decode(sent.feedback.importContent)).toBe("https://tappedout.net/mtg-decks/x/");
  });

  it("shows the answer of the server to text that is no link, with no code", async () => {
    fetchDeckList.mockRejectedValueOnce(new ConnectError("url: paste the link of a deck page", Code.InvalidArgument));
    await openAndRead("hello");

    expect(await screen.findByRole("alert")).toHaveTextContent(/^paste the link of a deck page$/);
  });

  it("sends no link for a list that the reader edited after the read (D-1101)", async () => {
    const text = "Deck\n4 Lightning Bolt (2x2) 117\n";
    fetchDeckList.mockResolvedValueOnce({ text, name: "Burn", sourceUrl: "https://archidekt.com/decks/42", leftOut: 0, site: "Archidekt", exportSteps: [], knownSite: true });
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    const user = await openAndRead("archidekt.com/decks/42");

    const list = screen.getByRole("textbox", { name: "Deck list" });
    await waitFor(() => expect(list).toHaveValue(text));
    await user.type(list, "1 Shock");
    await user.click(screen.getByRole("button", { name: "Import" }));
    await waitFor(() => expect(importDeck).toHaveBeenCalledWith(expect.objectContaining({ sourceUrl: "" })));
  });

  it("sends no link for a list that the reader pasted", async () => {
    importDeck.mockResolvedValueOnce({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [] });
    await openAndPaste("4 Lightning Bolt");
    await waitFor(() => expect(importDeck).toHaveBeenCalledWith(expect.objectContaining({ sourceUrl: "" })));
    expect(fetchDeckList).not.toHaveBeenCalled();
  });
});

describe("the report of a list the app could not read", () => {
  it("names no service on the pick step", async () => {
    await renderAt("/decks");
    const user = userEvent.setup();
    await user.click(await screen.findByRole("button", { name: "Import a deck" }));
    expect(screen.getByRole("dialog")).not.toHaveTextContent(/ManaBox|Moxfield|Arena|Archidekt/);
  });

  it("asks where the list came from, and sends it as a deck report", async () => {
    importDeck.mockRejectedValue(new ConnectError("no line of the list reads as a card", Code.InvalidArgument, undefined, [{ desc: UnreadableFileSchema, value: create(UnreadableFileSchema) }]));
    submitFeedback.mockResolvedValue({ feedbackId: "fb1" });
    const user = await openAndPaste("hello");

    expect(await screen.findByText("Something went wrong. The app could not read this file.")).toBeInTheDocument();
    expect(screen.queryByText(/Import failed/)).not.toBeInTheDocument();
    await user.click(screen.getByRole("button", { name: "Send a report" }));
    await screen.findByText(/Your report will be reviewed so that we can resolve the issue/);
    expect(submitFeedback).toHaveBeenCalledWith({ feedback: expect.objectContaining({ kind: FeedbackKind.IMPORT, importPage: ImportPage.DECK, text: "" }) });
  });

  it("offers the form beside the lines that did not parse", async () => {
    importDeck.mockResolvedValue({ deck: stored, sessionId: "s9", commanderOptions: [], unresolved: [{ line: 2, raw: "hello", reason: UnresolvedReason.BAD_ROW }] });
    await openAndPaste("4 Lightning Bolt{enter}hello");

    expect(await screen.findByText("The app could not read 1 row of this file.")).toBeInTheDocument();
  });
});
