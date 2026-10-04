// Command user-case replays the shortlist builds of a real user session
// against the user's own collection, and reads each count against a bar
// (D-1124).
//
// The free lane builds each step's shortlist with candidates.Build and
// costs nothing. It prints counts alone, never a card name, so the output
// can be pasted in a public place (D-639). The paid lane is chat-probe,
// and the case file holds its messages. chat-probe writes each deck that
// reached the user to a local file, and -decks reads the deck bars from
// it.
//
// The collection sits in a private object. The command downloads it with
// gcloud, refuses it when the sha256 differs from the case, and keeps it
// under .local alone. The card snapshot and the quality model are pinned
// to the versions the session ran on, so a change of the numbers is a
// change of the builder.
//
// Usage:
//
//	CLOUDSDK_CORE_ACCOUNT=<account> go -C go run ./cmd/user-case
//	go -C go run ./cmd/user-case -print messages
//	go -C go run ./cmd/user-case -decks ../.local/probes/<case>.decks.jsonl
package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"google.golang.org/protobuf/encoding/protojson"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
	"github.com/nkramber/decktome/go/internal/candidates"
	"github.com/nkramber/decktome/go/internal/cards"
	"github.com/nkramber/decktome/go/internal/gatekit"
	"github.com/nkramber/decktome/go/internal/meta"
	"github.com/nkramber/decktome/go/internal/quality"
)

// errFail is a bar that failed. It exits 1, and a setup fault exits 2.
var errFail = errors.New("a bar failed")

func main() {
	err := run()
	switch {
	case errors.Is(err, errFail):
		fmt.Fprintln(os.Stderr, "user-case:", err)
		os.Exit(1)
	case err != nil:
		fmt.Fprintln(os.Stderr, "user-case:", err)
		os.Exit(2)
	}
}

type opts struct {
	caseArg, local, store, snapshot, model, print, decks string
	noFetch                                              bool
}

func run() error {
	var o opts
	flag.StringVar(&o.caseArg, "case", "first-user-ms-marvel", "a case name of cases/, or the path of a case file")
	flag.StringVar(&o.local, "local", "../.local/user-cases", "the local root of the cases, ignored by git")
	flag.StringVar(&o.store, "store", "", "a store root that holds scryfall/ and meta/ (default <local>/store)")
	flag.StringVar(&o.snapshot, "snapshot", "", "a snapshot version, or newest (default: the version of the case)")
	flag.StringVar(&o.model, "model", "", "a quality model version, or none (default: the version of the case)")
	flag.StringVar(&o.print, "print", "", "print one value and exit: messages, or collection (fetch the pinned files, print the CSV path)")
	flag.BoolVar(&o.noFetch, "no-fetch", false, "refuse a missing file in place of a download")
	flag.StringVar(&o.decks, "decks", "", "the decks file of chat-probe -decks-out: read the deck bars of each step")
	flag.Parse()

	c, err := LoadCase(o.caseArg)
	if err != nil {
		return err
	}
	if o.store == "" {
		o.store = filepath.Join(o.local, "store")
	}
	csvPath := filepath.Join(o.local, c.Name, "collection.csv")
	switch o.print {
	case "":
	case "messages":
		fmt.Println(strings.Join(c.Paid.Messages, "|"))
		return nil
	case "collection":
		if err := fetch(c, o, csvPath); err != nil {
			return err
		}
		abs, err := filepath.Abs(csvPath)
		if err != nil {
			return err
		}
		fmt.Println(abs)
		return nil
	default:
		return fmt.Errorf("-print takes messages or collection, not %q", o.print)
	}
	return replay(c, o, csvPath)
}

func replay(c *Case, o opts, csvPath string) error {
	ctx := context.Background()
	if err := fetch(c, o, csvPath); err != nil {
		return err
	}
	idx, version, err := snapshot(ctx, c, o)
	if err != nil {
		return err
	}
	scorer, modelVersion, err := scorer(ctx, c, o)
	if err != nil {
		return err
	}
	owned, note, err := gatekit.LoadOwned(csvPath, idx)
	if err != nil {
		return err
	}
	cmdr, ok := idx.ByName(c.Slots.Commander)
	if !ok {
		return fmt.Errorf("snapshot %s holds no card named %q", version, c.Slots.Commander)
	}
	req, err := request(c, cmdr, owned)
	if err != nil {
		return err
	}
	if scorer != nil {
		req.MetaBoost = scorer.MetaBoost(req.Format)
	}
	cb, err := candidates.New()
	if err != nil {
		return err
	}
	tags := idx.Tags()
	sets := map[string]map[string]bool{}
	tagged := func(tag, oracleID string) bool {
		if sets[tag] == nil {
			sets[tag] = map[string]bool{}
			for _, id := range tags.Resolve(tag) {
				sets[tag][id] = true
			}
		}
		return sets[tag][oracleID]
	}

	fmt.Printf("case %s (%s). Snapshot %s. Model %s.\n", c.Name, c.Decision, version, modelVersion)
	fmt.Printf("collection: %s, %d oracle ids.\n", note, len(owned))
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(tw, "step\tmeasure\tbaseline\tnow\tbar\tverdict")
	failed, todo := 0, 0
	for _, step := range c.Steps {
		req.Theme, req.Avoid = step.Theme, step.Avoid
		list, err := cb.Build(idx, req)
		if err != nil {
			return fmt.Errorf("step %s: %w", step.Name, err)
		}
		picked := make([]*mtgv1.Card, 0, len(list.Candidates))
		for _, cand := range list.Candidates {
			picked = append(picked, cand.Card)
		}
		for _, r := range Judge(step, Count(picked, list.Theme.Unmatched, tagged)) {
			switch r.Verdict {
			case Fail:
				failed++
			case Todo:
				todo++
			}
			_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\t%s\n", r.Step, r.ID, num(r.Baseline), num(r.Now), r.Bar, r.Verdict)
		}
	}
	if o.decks == "" {
		if err := tw.Flush(); err != nil {
			return err
		}
		fmt.Printf("deck measures: the paid lane (%s). Not measured here.\n", c.Paid.Target)
	} else {
		decks, err := readDecks(o.decks)
		if err != nil {
			return err
		}
		var prev *mtgv1.Deck
		for i, step := range c.Steps {
			d := deckAt(decks, step.AfterTurn)
			if i > 0 && d == deckAt(decks, c.Steps[i-1].AfterTurn) {
				d = nil
			}
			if d == nil {
				return fmt.Errorf("%w: no deck reached the user by turn %d of step %s", errFail, step.AfterTurn, step.Name)
			}
			for _, r := range Judge(Step{Name: step.Name + "-deck", Measures: step.Deck}, DeckCount(d, prev, idx.ByOracleID, tagged)) {
				switch r.Verdict {
				case Fail:
					failed++
				case Todo:
					todo++
				}
				_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\t%s\n", r.Step, r.ID, num(r.Baseline), num(r.Now), r.Bar, r.Verdict)
			}
			prev = d
		}
		if err := tw.Flush(); err != nil {
			return err
		}
	}
	fmt.Printf("bars: %d failed, %d not set.\n", failed, todo)
	if failed > 0 {
		return fmt.Errorf("%w: %d of the measures", errFail, failed)
	}
	return nil
}

// request maps the slots of the case onto a shortlist request.
func request(c *Case, cmdr *mtgv1.Card, owned map[string]int32) (candidates.Request, error) {
	req := candidates.Request{
		Format:             gatekit.FormatID(c.Slots.Format),
		CommanderOracleIDs: []string{cmdr.GetOracleId()},
		Bracket:            c.Slots.Bracket,
		Owned:              owned,
	}
	if req.Format == mtgv1.FormatId_FORMAT_ID_UNSPECIFIED {
		return req, fmt.Errorf("case format %q is unknown", c.Slots.Format)
	}
	for _, s := range c.Slots.Colors {
		col, ok := colors[strings.ToUpper(s)]
		if !ok {
			return req, fmt.Errorf("case color %q is unknown", s)
		}
		req.Colors = append(req.Colors, col)
	}
	pool, ok := pools[c.Slots.Pool]
	if !ok {
		return req, fmt.Errorf("case pool %q is unknown", c.Slots.Pool)
	}
	req.PoolRule = pool
	return req, nil
}

var colors = map[string]mtgv1.Color{
	"W": mtgv1.Color_COLOR_W, "U": mtgv1.Color_COLOR_U, "B": mtgv1.Color_COLOR_B,
	"R": mtgv1.Color_COLOR_R, "G": mtgv1.Color_COLOR_G,
}

var pools = map[string]mtgv1.PoolRule{
	"owned-only":  mtgv1.PoolRule_POOL_RULE_OWNED_ONLY,
	"owned-first": mtgv1.PoolRule_POOL_RULE_OWNED_FIRST,
	"any-card":    mtgv1.PoolRule_POOL_RULE_ANY_CARD,
}

// fetch makes sure the pinned collection, snapshot, and model are local.
// A version named by a flag is never downloaded.
func fetch(c *Case, o opts, csvPath string) error {
	if err := collection(c, csvPath, o.noFetch); err != nil {
		return err
	}
	root := filepath.Join(o.store, "scryfall")
	if o.snapshot == "" && absent(filepath.Join(root, c.Snapshot.Version, "complete")) && !o.noFetch {
		if err := gcloudCopy(c.Snapshot.Object, root, true); err != nil {
			return err
		}
	}
	models := filepath.Join(o.store, meta.ModelPrefix)
	if o.model == "" && c.Model.Version != "" && absent(filepath.Join(models, c.Model.Version, meta.ModelFile)) && !o.noFetch {
		if err := gcloudCopy(c.Model.Object, models, true); err != nil {
			return err
		}
	}
	return nil
}

func absent(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, os.ErrNotExist)
}

// collection makes sure the local CSV holds the bytes the case pins. The
// gzip object is the pinned file, and the CSV beside it is its content.
func collection(c *Case, csvPath string, noFetch bool) error {
	gz := csvPath + ".gz"
	if _, err := os.Stat(gz); errors.Is(err, os.ErrNotExist) {
		if noFetch {
			return fmt.Errorf("%s is absent and -no-fetch is set", gz)
		}
		if err := gcloudCopy(c.Collection.Object, gz, false); err != nil {
			return err
		}
	}
	raw, err := os.ReadFile(gz) // #nosec G304 -- the path is built from the case name.
	if err != nil {
		return err
	}
	if err := checkSHA(raw, c.Collection.SHA256); err != nil {
		return fmt.Errorf("%s: %w. Delete it to download it again", gz, err)
	}
	return gunzipTo(raw, csvPath)
}

// checkSHA refuses bytes whose sha256 differs from the pinned one.
func checkSHA(raw []byte, want string) error {
	sum := sha256.Sum256(raw)
	if got := hex.EncodeToString(sum[:]); !strings.EqualFold(got, want) {
		return fmt.Errorf("sha256 %s, and the case pins %s", got, want)
	}
	return nil
}

func gunzipTo(raw []byte, path string) error {
	r, err := gzip.NewReader(bytes.NewReader(raw))
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600) // #nosec G304 -- the path is built from the case name.
	if err != nil {
		return err
	}
	if _, err := io.Copy(f, r); err != nil { // #nosec G110 -- the hash pins the object.
		_ = f.Close()
		return err
	}
	return f.Close()
}

// snapshot loads the pinned card snapshot, or the version a flag names.
func snapshot(ctx context.Context, c *Case, o opts) (*cards.Index, string, error) {
	store := cards.DirStore{Root: filepath.Join(o.store, "scryfall")}
	version := c.Snapshot.Version
	switch o.snapshot {
	case "newest":
		idx, v, err := cards.LoadNewest(ctx, store, gatekit.Quiet())
		return idx, v, err
	case "":
	default:
		version = o.snapshot
	}
	if absent(filepath.Join(store.Root, version, "complete")) {
		return nil, "", fmt.Errorf("snapshot %s is absent under %s", version, store.Root)
	}
	idx, err := cards.LoadVersion(ctx, store, version, gatekit.Quiet())
	return idx, version, err
}

// scorer loads the pinned quality model, whose meta signal ranks the
// shortlist the way the session ranked it (PR-14B).
func scorer(ctx context.Context, c *Case, o opts) (*quality.Scorer, string, error) {
	version := c.Model.Version
	switch o.model {
	case "none":
		return nil, "none", nil
	case "":
	default:
		version = o.model
	}
	if version == "" {
		return nil, "none", nil
	}
	data, err := meta.ReadModel(ctx, meta.DirObjects{Root: o.store}, version)
	if err != nil {
		return nil, "", err
	}
	model, err := quality.Decode(data)
	if err != nil {
		return nil, "", err
	}
	return quality.NewScorer(model), version, nil
}

// gcloudCopy downloads one object, or one prefix with recursive. gcloud
// reads the account from CLOUDSDK_CORE_ACCOUNT, so no config changes.
func gcloudCopy(src, dst string, recursive bool) error {
	dir := filepath.Dir(dst)
	if recursive {
		dir = dst
	}
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return err
	}
	args := []string{"storage", "cp", "--quiet"}
	if recursive {
		args = append(args, "-r")
	}
	args = append(args, src, dst)
	fmt.Fprintf(os.Stderr, "download %s\n", src)
	cmd := exec.Command("gcloud", args...) // #nosec G204 -- the case file names the object.
	cmd.Stdout, cmd.Stderr = os.Stderr, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gcloud storage cp %s: %w", src, err)
	}
	return nil
}

// turnDeck is one line of the decks file of chat-probe: the turn, and
// the deck that reached the user on it.
type turnDeck struct {
	Turn int             `json:"turn"`
	Deck json.RawMessage `json:"deck"`
}

// readDecks reads the decks file of chat-probe, by turn.
func readDecks(path string) (map[int]*mtgv1.Deck, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	out := map[int]*mtgv1.Deck{}
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 1<<20), 16<<20)
	for sc.Scan() {
		if len(bytes.TrimSpace(sc.Bytes())) == 0 {
			continue
		}
		var td turnDeck
		if err := json.Unmarshal(sc.Bytes(), &td); err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		d := &mtgv1.Deck{}
		if err := protojson.Unmarshal(td.Deck, d); err != nil {
			return nil, fmt.Errorf("%s turn %d: %w", path, td.Turn, err)
		}
		out[td.Turn] = d
	}
	return out, sc.Err()
}

// deckAt answers the newest deck that reached the user by the turn.
func deckAt(decks map[int]*mtgv1.Deck, turn int) *mtgv1.Deck {
	for t := turn; t > 0; t-- {
		if d, ok := decks[t]; ok {
			return d
		}
	}
	return nil
}
