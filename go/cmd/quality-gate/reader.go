package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/evalrun"
	"github.com/nkramber/decktome/go/internal/generate"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/quality"
)

// The reader-facing numbers of D-648. The precon bar is a proxy, so the
// gate document of every quality item reads two more numbers beside it:
// the built decks the fitted model grades bad, and the model's agreement
// with one judge run over the same decks. An item that moves none of the
// three is dropped.
//
// Both are free. The decks come from a deck gate document, the grades from
// the model this run fitted, and the judge's tiers from a judge lane
// document that already ran. One judge run serves every model, so the
// agreement of two models holds no judge noise.

var (
	judgedSource = regexp.MustCompile(`Source: (.+)\.$`)
	judgedRow    = regexp.MustCompile(`^\| (\d+) \| (.+) \| [A-Za-z]+ \| [a-z]* \| [a-z]* \| ([a-z]+) \| (yes|no|) \|$`)
)

// judgeAnswer is one deck of a judge lane document: its title and the
// tier the judge read.
type judgeAnswer struct {
	title, tier string
}

// readerRead is the two numbers over one deck gate document.
type readerRead struct {
	decksPath, judgedPath string
	decks                 []judged
	answers               map[int]judgeAnswer
}

// readJudged reads the deck gate document a judge lane document judged,
// and the tier the judge read for each deck. A deck the judge failed on
// has no tier, and it takes no part in the agreement.
func readJudged(r io.Reader) (source string, answers map[int]judgeAnswer, err error) {
	answers = map[int]judgeAnswer{}
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		line := sc.Text()
		if source == "" && strings.HasPrefix(line, "Run date:") {
			if m := judgedSource.FindStringSubmatch(line); m != nil {
				source = m[1]
			}
			continue
		}
		m := judgedRow.FindStringSubmatch(line)
		if m == nil || meta.TierLevel(m[3]) < 0 {
			continue
		}
		id, _ := strconv.Atoi(m[1])
		answers[id] = judgeAnswer{title: m[2], tier: m[3]}
	}
	if err := sc.Err(); err != nil {
		return "", nil, err
	}
	if source == "" {
		return "", nil, fmt.Errorf("the judge document names no source deck gate document")
	}
	if len(answers) == 0 {
		return "", nil, fmt.Errorf("the judge document holds no judged deck")
	}
	return source, answers, nil
}

// gradeReader grades the decks of a deck gate document with the fitted
// model, and reads the judge's tiers when a judge document is named.
func gradeReader(idx *cards.Index, model *quality.Model, decksPath, judgedPath, promptsPath string) (*readerRead, error) {
	commanders, err := promptCommanders(promptsPath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(decksPath) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	decks, err := readDeckGate(f, idx, commanders)
	if err != nil {
		return nil, fmt.Errorf("decks %s: %w", decksPath, err)
	}
	if err := gradeDecks(decks, idx, quality.NewScorer(model)); err != nil {
		return nil, err
	}
	rr := &readerRead{decksPath: decksPath, judgedPath: judgedPath, decks: decks}
	if judgedPath == "" {
		return rr, nil
	}
	jf, err := os.Open(judgedPath) // #nosec G304 -- the operator names the file.
	if err != nil {
		return nil, err
	}
	defer func() { _ = jf.Close() }()
	source, answers, err := readJudged(jf)
	if err != nil {
		return nil, fmt.Errorf("judged %s: %w", judgedPath, err)
	}
	if err := sameDecks(source, decksPath, decks, answers); err != nil {
		return nil, err
	}
	rr.answers = answers
	return rr, nil
}

// sameDecks refuses a judge document that read other decks. The agreement
// compares a grade with a judge answer, and an answer about another deck
// compares nothing.
func sameDecks(source, decksPath string, decks []judged, answers map[int]judgeAnswer) error {
	if filepath.Base(source) != filepath.Base(decksPath) {
		return fmt.Errorf("the judge document judged %s, and -decks names %s", filepath.Base(source), filepath.Base(decksPath))
	}
	for _, d := range decks {
		if a, ok := answers[d.id]; ok && a.title != d.title {
			return fmt.Errorf("deck %d reads %q in the judge document and %q in the deck gate document", d.id, a.title, d.title)
		}
	}
	return nil
}

// counts answers the decks graded bad, the decks the judge answered, and
// the ones among those where the judge agreed with the grade.
func (r *readerRead) counts() (bad, judgedCount, agreed int) {
	for _, d := range r.decks {
		if d.modelGrade == meta.TierBad {
			bad++
		}
		a, ok := r.answers[d.id]
		if !ok {
			continue
		}
		judgedCount++
		if a.tier == d.modelGrade {
			agreed++
		}
	}
	return bad, judgedCount, agreed
}

// readerSection writes the three numbers of D-648 in one place, and the
// grade and the judge's tier of every deck under them.
func readerSection(p func(string, ...any), rep *quality.FitReport, r *readerRead, run *evalrun.Run) {
	bad, judgedCount, agreed := r.counts()
	p("## The three numbers (D-648)\n\n")
	p("The precon bar is a proxy, so every quality item reads two more numbers beside it. The model this run fitted grades the built decks of `%s`.", filepath.Base(r.decksPath))
	if r.answers != nil {
		p(" The judge's tiers come from `%s`. One judge run serves every model, so the agreement holds no judge noise.", filepath.Base(r.judgedPath))
	}
	p(" No bar reads the last two, and an item that moves none of the three is dropped.\n\n")
	p("| Number | Read |\n|---|---|\n")
	if fr := rep.Formats[meta.FormatCommander]; fr != nil {
		own := fr.Holdout.OwnByDefect[quality.DefectSynergy]
		word := "no pair"
		if own.Pairs > 0 {
			word = fmt.Sprintf("%.2f of %d", own.Share(), own.Pairs)
		}
		p("| Commander synergy axis, precon over own copy | %s |\n", word)
	}
	p("| Built decks graded bad | %d of %d |\n", bad, len(r.decks))
	run.Info("suite", "decks_graded_bad", float64(bad), fmt.Sprintf("%d of %d built decks of %s", bad, len(r.decks), filepath.Base(r.decksPath)))
	if r.answers != nil {
		p("| Judge agreement | %d of %d |\n", agreed, judgedCount)
		run.Info("suite", "judge_agreement", float64(agreed), fmt.Sprintf("%d of %d decks of %s", agreed, judgedCount, filepath.Base(r.judgedPath)))
	}
	p("\n| # | Deck | Format | Grade | Judge | Agree |\n|---|---|---|---|---|---|\n")
	for _, d := range r.decks {
		judge, agree := "-", "-"
		if a, ok := r.answers[d.id]; ok {
			judge, agree = a.tier, "no"
			if a.tier == d.modelGrade {
				agree = "yes"
			}
		}
		p("| %d | %s | %s | %s | %s | %s |\n", d.id, d.title, generate.FormatWord(d.format), d.modelGrade, judge, agree)
	}
	p("\n")
}
