package candidates

// englishStopWords is the English stop list of the Snowball project,
// read from snowballstem.org/algorithms/english/stop.txt on 2026-10-05.
// The tokenizer splits a word at an apostrophe, so the contractions of
// the source are left out. No word of it names a theme row, a card type,
// or a subtype of a playable card, and TestStopWordsNameNoTheme reads
// the rows and the types (D-1188).
var englishStopWords = map[string]bool{
	"i": true, "me": true, "my": true, "myself": true, "we": true, "our": true, "ours": true,
	"ourselves": true, "you": true, "your": true, "yours": true, "yourself": true, "yourselves": true,
	"he": true, "him": true, "his": true, "himself": true, "she": true, "her": true, "hers": true,
	"herself": true, "it": true, "its": true, "itself": true, "they": true, "them": true, "their": true,
	"theirs": true, "themselves": true, "what": true, "which": true, "who": true, "whom": true, "this": true, "that": true, "these": true, "those": true, "am": true, "is": true, "are": true, "was": true,
	"were": true, "be": true, "been": true, "being": true, "have": true, "has": true, "had": true,
	"having": true, "do": true, "does": true, "did": true, "doing": true, "would": true, "should": true,
	"could": true, "ought": true, "cannot": true, "a": true, "an": true, "the": true, "and": true,
	"but": true, "if": true, "or": true, "because": true, "as": true, "until": true, "while": true,
	"of": true, "at": true, "by": true, "for": true, "with": true, "about": true, "against": true,
	"between": true, "into": true, "through": true, "during": true, "before": true, "after": true,
	"above": true, "below": true, "to": true, "from": true, "up": true, "down": true, "in": true, "out": true, "on": true, "off": true, "over": true, "under": true, "again": true, "further": true, "then": true, "once": true, "here": true, "there": true, "when": true, "where": true, "why": true, "how": true, "all": true, "any": true, "both": true, "each": true, "few": true, "more": true, "most": true,
	"other": true, "some": true, "such": true, "no": true, "nor": true, "not": true, "only": true,
	"own": true, "same": true, "so": true, "than": true, "too": true, "very": true,
}

// requestFiller are the words a reader writes around a theme, and no
// theme row, card type, or subtype holds one. A list that grew one chat
// at a time missed each new phrase, so this one names the classes
// (D-1188): words about a deck or a set, hedges, vague nouns, and
// indefinite pronouns.
var requestFiller = map[string]bool{
	"mechanic": true, "mechanics": true, "synergy": true, "synergies": true, "strategy": true,
	"strategies": true, "archetype": true, "archetypes": true, "gameplan": true, "signature": true,
	"set": true, "sets": true, "expansion": true, "expansions": true, "block": true,
	"sideboard": true, "mainboard": true, "maindeck": true, "decklist": true,
	"maybe": true, "perhaps": true, "possibly": true, "probably": true, "mostly": true, "mainly": true,
	"primarily": true, "especially": true, "also": true, "plus": true, "etc": true, "else": true,
	"whatever": true, "whichever": true, "everything": true, "similar": true, "related": true,
	"additional": true, "along": true, "together": true, "various": true, "different": true,
	"main": true, "general": true, "overall": true, "including": true, "include": true, "centered": true,
	"kind": true, "kinds": true, "sort": true, "thing": true, "things": true,
	"nobody": true, "anybody": true, "somebody": true, "everybody": true, "everyone": true, "nothing": true,
	"fit": true, "fits": true, "work": true, "works": true, "goes": true, "well": true,
}
