// Package push keeps the devices that take a web push, and it sends the
// two events of PR-26. The API sends a finished build for a user who
// left the page (D-1004, D-1005). The snapshot job sends a legality
// change that made decks of the user illegal (D-1087, D-1088).
//
// A device is one browser install, named by its Firebase Installation
// ID. The devices of a user sit under users/<uid>/push_devices. A user
// with no device gets no push, so a registered device is the opt-in.
//
// The ID belongs to the browser and not to an account, so one ID has one
// owner at a time. The document push_owners/<id> names the owner, and a
// registration by a second account moves the ID to that account.
package push

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mtgv1 "github.com/nkramber/decktome/go/gen/mtg/v1"
)

// Collection is the subcollection of users/<uid> that holds the devices.
const Collection = "push_devices"

// Owners is the top-level collection that names the owner of each ID.
const Owners = "push_owners"

// MaxDevices caps the devices of one user. A new device past the cap
// removes the device of the oldest registration, so one user can never
// make the send fan out.
const MaxDevices = 10

// maxIDBytes caps an installation id. A Firebase Installation ID has 22
// characters, and the cap leaves room for a change of that shape.
const maxIDBytes = 256

// TTL is how long Cloud Messaging holds a push for a device that is
// offline. A deck that waited a day is no news.
const TTL = 24 * time.Hour

// ErrBadID is the answer for an installation id that is empty, too long,
// or holds a character outside the base64url set.
var ErrBadID = errors.New("push: the installation id is not valid")

// ValidID says whether id can be an installation id. The id is also the
// document id, so the check keeps a slash and a dot out of a path.
func ValidID(id string) bool {
	if id == "" || len(id) > maxIDBytes {
		return false
	}
	for _, r := range id {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '-', r == '_':
		default:
			return false
		}
	}
	return true
}

// Device is one stored device.
type Device struct {
	InstallationID string    `firestore:"installation_id"`
	RegisteredAt   time.Time `firestore:"registered_at"`
}

// Repo stores the devices in Firestore.
type Repo struct {
	client *firestore.Client
}

// NewRepo wraps a Firestore client.
func NewRepo(client *firestore.Client) *Repo { return &Repo{client: client} }

func (r *Repo) col(uid string) *firestore.CollectionRef {
	return r.client.Collection("users").Doc(uid).Collection(Collection)
}

func (r *Repo) owner(id string) *firestore.DocumentRef {
	return r.client.Collection(Owners).Doc(id)
}

// ownerOf reads the owner of the ID in the transaction, or "" for none.
func (r *Repo) ownerOf(tx *firestore.Transaction, id string) (string, error) {
	snap, err := tx.Get(r.owner(id))
	if status.Code(err) == codes.NotFound {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	uid, _ := snap.Data()["uid"].(string)
	return uid, nil
}

// Add stores the device, or moves the time of a device the user holds.
// An ID that another account holds leaves that account in the same
// transaction, so a push of one account never reaches the next account
// on the browser. Then Add removes each device past MaxDevices, the
// oldest first.
func (r *Repo) Add(ctx context.Context, uid, id string, at time.Time) error {
	if !ValidID(id) {
		return ErrBadID
	}
	err := r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		prev, err := r.ownerOf(tx, id)
		if err != nil {
			return err
		}
		if prev != "" && prev != uid {
			if err := tx.Delete(r.col(prev).Doc(id)); err != nil {
				return err
			}
		}
		if err := tx.Set(r.owner(id), map[string]any{"uid": uid, "registered_at": at.UTC()}); err != nil {
			return err
		}
		return tx.Set(r.col(uid).Doc(id), Device{InstallationID: id, RegisteredAt: at.UTC()})
	})
	if err != nil {
		return fmt.Errorf("push: add device: %w", err)
	}
	it := r.col(uid).OrderBy("registered_at", firestore.Desc).Offset(MaxDevices).Documents(ctx)
	defer it.Stop()
	for {
		snap, err := it.Next()
		if errors.Is(err, iterator.Done) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("push: read devices past the cap: %w", err)
		}
		if err := r.remove(ctx, uid, snap.Ref.ID); err != nil {
			return fmt.Errorf("push: remove a device past the cap: %w", err)
		}
	}
}

// Remove deletes the device. A device the user does not hold is not an
// error.
func (r *Repo) Remove(ctx context.Context, uid, id string) error {
	if !ValidID(id) {
		return ErrBadID
	}
	if err := r.remove(ctx, uid, id); err != nil {
		return fmt.Errorf("push: remove device: %w", err)
	}
	return nil
}

// remove deletes the device of the user, and the owner document when it
// names the user. An owner document of another account stays.
func (r *Repo) remove(ctx context.Context, uid, id string) error {
	return r.client.RunTransaction(ctx, func(_ context.Context, tx *firestore.Transaction) error {
		owner, err := r.ownerOf(tx, id)
		if err != nil {
			return err
		}
		if owner == uid {
			if err := tx.Delete(r.owner(id)); err != nil {
				return err
			}
		}
		return tx.Delete(r.col(uid).Doc(id))
	})
}

// Has says whether the user holds the device.
func (r *Repo) Has(ctx context.Context, uid, id string) (bool, error) {
	if !ValidID(id) {
		return false, ErrBadID
	}
	_, err := r.col(uid).Doc(id).Get(ctx)
	if status.Code(err) == codes.NotFound {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("push: read device: %w", err)
	}
	return true, nil
}

// IDs returns the installation ids of the user, at most MaxDevices.
func (r *Repo) IDs(ctx context.Context, uid string) ([]string, error) {
	snaps, err := r.col(uid).Limit(MaxDevices).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("push: read devices: %w", err)
	}
	ids := make([]string, 0, len(snaps))
	for _, s := range snaps {
		ids = append(ids, s.Ref.ID)
	}
	return ids, nil
}

// Message is what one push shows. URL is a path of the app.
type Message struct {
	Title string
	Body  string
	URL   string
}

// Sender delivers one message to each device. It returns the devices
// that Cloud Messaging says are gone, so the caller can remove them.
type Sender interface {
	Send(ctx context.Context, ids []string, m Message) (gone []string, err error)
}

// FCM sends through Firebase Cloud Messaging, with the credentials of
// the service. The service account needs the role
// roles/firebasecloudmessaging.admin (D-1005).
type FCM struct {
	client *messaging.Client
}

// NewFCM builds the sender for one project.
func NewFCM(ctx context.Context, projectID string) (*FCM, error) {
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: projectID})
	if err != nil {
		return nil, fmt.Errorf("push: firebase app: %w", err)
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("push: messaging client: %w", err)
	}
	return &FCM{client: client}, nil
}

// Send implements Sender. Each push is a data message, and the service
// worker of the app shows it (`web/apps/web/public/push-handler.js`).
func (f *FCM) Send(ctx context.Context, ids []string, m Message) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	msgs := make([]*messaging.Message, 0, len(ids))
	for _, id := range ids {
		msgs = append(msgs, &messaging.Message{
			Fid:  id,
			Data: map[string]string{"title": m.Title, "body": m.Body, "url": m.URL},
			Webpush: &messaging.WebpushConfig{Headers: map[string]string{
				"Urgency": "high",
				"TTL":     fmt.Sprint(int(TTL.Seconds())),
			}},
		})
	}
	res, err := f.client.SendEach(ctx, msgs)
	if err != nil {
		return nil, fmt.Errorf("push: send: %w", err)
	}
	var gone []string
	var failed error
	for i, r := range res.Responses {
		if r.Success {
			continue
		}
		if messaging.IsUnregistered(r.Error) {
			gone = append(gone, ids[i])
			continue
		}
		failed = errors.Join(failed, r.Error)
	}
	return gone, failed
}

// Store is what the notifier reads and prunes.
type Store interface {
	IDs(ctx context.Context, uid string) ([]string, error)
	Remove(ctx context.Context, uid, id string) error
}

// Notifier tells a user that a deck is ready, or that decks are no
// longer legal.
type Notifier struct {
	store  Store
	sender Sender
	log    *slog.Logger
}

// NewNotifier joins the store and the sender.
func NewNotifier(store Store, sender Sender, log *slog.Logger) *Notifier {
	if log == nil {
		log = slog.Default()
	}
	return &Notifier{store: store, sender: sender, log: log}
}

// DeckReady sends one push to each device of the user. A failure goes to
// the log alone, because the deck is stored and the next visit shows it.
// A device that Cloud Messaging says is gone leaves the store.
func (n *Notifier) DeckReady(ctx context.Context, uid string, d *mtgv1.Deck) {
	n.send(ctx, uid, DeckMessage(d), "deck ready", "deck", d.GetId())
}

// DecksStale sends one push to each device of the user for all the decks
// that one stale pass hit (D-1088). A failure goes to the log alone,
// because the banner of each deck holds the same news.
func (n *Notifier) DecksStale(ctx context.Context, uid string, decks []*mtgv1.Deck) {
	if len(decks) == 0 {
		return
	}
	n.send(ctx, uid, StaleMessage(decks), "decks stale", "decks", len(decks))
}

// NewCards sends one push to each device of the user for all the decks
// that one new-cards pass gave cards (D-1091). A failure goes to the log
// alone, because the panel of each deck holds the same news.
func (n *Notifier) NewCards(ctx context.Context, uid string, decks []*mtgv1.Deck) {
	if len(decks) == 0 {
		return
	}
	n.send(ctx, uid, NewCardsMessage(decks), "new cards", "decks", len(decks))
}

func (n *Notifier) send(ctx context.Context, uid string, m Message, event string, key string, val any) {
	ids, err := n.store.IDs(ctx, uid)
	if err != nil {
		n.log.ErrorContext(ctx, "push: the devices were not read", "event", event, key, val, "err", err)
		return
	}
	if len(ids) == 0 {
		return
	}
	gone, err := n.sender.Send(ctx, ids, m)
	if err != nil {
		n.log.ErrorContext(ctx, "push: a send failed", "event", event, key, val, "devices", len(ids), "err", err)
	}
	for _, id := range gone {
		if err := n.store.Remove(ctx, uid, id); err != nil {
			n.log.ErrorContext(ctx, "push: a gone device stayed", "err", err)
		}
	}
	n.log.InfoContext(ctx, "push: "+event+" sent", key, val, "devices", len(ids), "gone", len(gone))
}

// DeckMessage is the push for a stored deck. A tap opens the deck.
//
// The title is the whole message. iOS shows the line "from Deck Tome"
// from the manifest name under the title, and a body adds a third line
// (F-198, D-1040).
func DeckMessage(d *mtgv1.Deck) Message {
	title := "Your deck is ready!"
	if d.GetRevisedFromDeckId() != "" {
		title = "Your revised deck is ready!"
	}
	return Message{Title: title, URL: "/decks/" + url.PathEscape(d.GetId())}
}

// StaleMessage is the push for the decks that a legality change made
// illegal. One deck opens that deck and its banner. More decks open the
// list, which marks each stale deck (D-1088).
func StaleMessage(decks []*mtgv1.Deck) Message {
	if len(decks) > 1 {
		return Message{Title: fmt.Sprintf("%d of your decks are no longer legal.", len(decks)), URL: "/decks"}
	}
	d := decks[0]
	title := "One of your decks is no longer legal."
	if name := strings.TrimSpace(d.GetName()); name != "" {
		title = `Your deck "` + name + `" is no longer legal.`
	}
	return Message{Title: title, URL: "/decks/" + url.PathEscape(d.GetId())}
}

// NewCardsMessage is the push for the decks that new cards of a new set
// fit (D-1091). One deck opens that deck and its panel of new cards. More
// decks open the list.
func NewCardsMessage(decks []*mtgv1.Deck) Message {
	if len(decks) > 1 {
		return Message{Title: fmt.Sprintf("New cards fit %d of your decks.", len(decks)), URL: "/decks"}
	}
	d := decks[0]
	title := "New cards fit one of your decks."
	if name := strings.TrimSpace(d.GetName()); name != "" {
		title = `New cards fit your deck "` + name + `".`
	}
	return Message{Title: title, URL: "/decks/" + url.PathEscape(d.GetId())}
}
