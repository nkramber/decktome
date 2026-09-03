package meta

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

// CEDHDBURL is the cEDH Decklist Database. One page holds every entry,
// each with a section word, the commanders, and links to Moxfield (read
// 2026-09-02).
const CEDHDBURL = "https://cedh-decklist-database.com/"

// The section words of the database.
const (
	SectionCompetitive = "COMPETITIVE"
	SectionBrew        = "BREW"
	SectionOutdated    = "OUTDATED"
)

// Entry is one database entry.
type Entry struct {
	Title      string   `json:"title"`
	Section    string   `json:"section"`
	Commanders []string `json:"commanders"`
	// Date is the entry's own date, YYYY-MM-DD.
	Date  string   `json:"date,omitempty"`
	Links []string `json:"links,omitempty"`
}

// Competitive says whether the entry sits in the competitive section,
// which is the great tier of cEDH (D-414).
func (e *Entry) Competitive() bool { return e.Section == SectionCompetitive }

var (
	cedhEntryStart = regexp.MustCompile(`<div class="ddb-colors`)
	cedhTitle      = regexp.MustCompile(`(?s)class="ddb-title">\s*(.*?)\s*</div>`)
	cedhSection    = regexp.MustCompile(`(?s)class="ddb-section[^"]*">\s*(.*?)\s*</div>`)
	cedhCommander  = regexp.MustCompile(`<li class="btn">([^<]+)</li>`)
	cedhDate       = regexp.MustCompile(`(?s)class="ddb-date[^"]*">\s*(\d{4}-\d{2}-\d{2})`)
	cedhLink       = regexp.MustCompile(`href="(https://(?:www\.)?moxfield\.com/decks/[^"]+)"`)
	cedhTag        = regexp.MustCompile(`<[^>]+>`)
)

// ParseCEDHDB reads every entry of the database page. A page with no
// entry is an error, because that is the markup change M-6 counts.
func ParseCEDHDB(page []byte) ([]Entry, error) {
	s := string(page)
	if i := strings.Index(s, `id="decks"`); i >= 0 {
		s = s[i:]
	}
	starts := cedhEntryStart.FindAllStringIndex(s, -1)
	if len(starts) == 0 {
		return nil, fmt.Errorf("meta: cedh database page holds no entry")
	}
	out := make([]Entry, 0, len(starts))
	for i, loc := range starts {
		end := len(s)
		if i+1 < len(starts) {
			end = starts[i+1][0]
		}
		block := s[loc[0]:end]
		e := Entry{}
		if m := cedhTitle.FindStringSubmatch(block); m != nil {
			e.Title = html.UnescapeString(strings.TrimSpace(m[1]))
		}
		if m := cedhSection.FindStringSubmatch(block); m != nil {
			e.Section = strings.TrimSpace(cedhTag.ReplaceAllString(m[1], ""))
		}
		for _, m := range cedhCommander.FindAllStringSubmatch(block, -1) {
			name := html.UnescapeString(strings.TrimSpace(m[1]))
			if name != "" {
				e.Commanders = append(e.Commanders, name)
			}
		}
		if m := cedhDate.FindStringSubmatch(block); m != nil {
			e.Date = m[1]
		}
		seen := map[string]bool{}
		for _, m := range cedhLink.FindAllStringSubmatch(block, -1) {
			if !seen[m[1]] {
				seen[m[1]] = true
				e.Links = append(e.Links, m[1])
			}
		}
		if e.Title == "" || len(e.Commanders) == 0 {
			continue
		}
		out = append(out, e)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("meta: cedh database page holds no readable entry")
	}
	return out, nil
}
