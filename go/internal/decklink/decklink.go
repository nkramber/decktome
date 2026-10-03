// Package decklink sorts a deck link that a user pasted into the site it
// names, and holds the steps of an export for each site that the app can
// not read (PR-121, D-1103).
//
// A site is one of three kinds. The app reads an Archidekt deck itself
// (D-1100). Moxfield refuses every client that is not a browser, so a
// Moxfield link gets the exact steps of its export (D-493, D-1103). Any
// other site gets the general steps, and the page files a report of its
// host, so the owner can add the site (D-1104).
package decklink

import (
	"errors"
	"net/url"
	"strings"

	"github.com/nkramber/decktome/go/internal/archidekt"
)

// Site is the site of a link.
type Site int

// The sites.
const (
	// Unknown is a site with no reader and no steps of its own.
	Unknown Site = iota
	Archidekt
	Moxfield
)

// maxBytes caps the link that a user pastes.
const maxBytes = 2048

// ErrNotLink refuses text that is no link of a deck page. It names no
// site (D-889).
var ErrNotLink = errors.New("url: paste the link of a deck page")

// Link is one sorted link.
type Link struct {
	Site Site
	// Host is the host of the link, lower case, with no "www.".
	Host string
	// DeckID is the id of an Archidekt deck.
	DeckID int64
}

// names are the names of the known sites.
var names = map[Site]string{Archidekt: "Archidekt", Moxfield: "Moxfield"}

// hosts maps the host of a known site to the site.
var hosts = map[string]Site{"archidekt.com": Archidekt, "moxfield.com": Moxfield}

// Parse sorts a link. It takes a link with or without the scheme, and
// with or without "www.". A link of a known site that names no deck, such
// as an Archidekt folder, is no link of a deck page.
func Parse(raw string) (Link, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" || len(raw) > maxBytes || strings.ContainsAny(raw, " \t\r\n") {
		return Link{}, ErrNotLink
	}
	full := raw
	if !strings.Contains(full, "://") {
		full = "https://" + full
	}
	u, err := url.Parse(full)
	if err != nil || (u.Scheme != "https" && u.Scheme != "http") || u.User != nil {
		return Link{}, ErrNotLink
	}
	host := strings.TrimPrefix(strings.ToLower(u.Hostname()), "www.")
	if !strings.Contains(host, ".") || strings.HasPrefix(host, ".") || strings.HasSuffix(host, ".") {
		return Link{}, ErrNotLink
	}
	link := Link{Site: hosts[host], Host: host}
	switch link.Site {
	case Archidekt:
		if link.DeckID, err = archidekt.ParseURL(raw); err != nil {
			return Link{}, ErrNotLink
		}
	case Moxfield:
		if parts := strings.Split(strings.Trim(u.Path, "/"), "/"); len(parts) < 2 || parts[0] != "decks" || parts[1] == "" {
			return Link{}, ErrNotLink
		}
	}
	return link, nil
}

// Name is the name of the site of a link: the name of a known site, or
// the host.
func (l Link) Name() string {
	if n, ok := names[l.Site]; ok {
		return n
	}
	return l.Host
}

// Steps are the steps of an export for a site that the app can not read.
// A known site has its exact steps, and any other site the general steps.
// The app reads Archidekt itself, so Archidekt has none.
func (l Link) Steps() []string {
	switch l.Site {
	case Archidekt:
		return nil
	case Moxfield:
		// The owner read the menu on 2026-10-03 (D-1103).
		return []string{
			"Open the deck on Moxfield.",
			"Select More, then Export.",
			"Select Copy for Arena.",
			"Paste the list in the box below, then select Import.",
		}
	}
	return []string{
		"Open the deck on its site.",
		"Find the export of the deck. It often has the name Export or Download.",
		"Copy the list as text. A list for Arena or for MTGO reads best.",
		"Paste the list in the box below, then select Import.",
	}
}
