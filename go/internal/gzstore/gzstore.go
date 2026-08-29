// Package gzstore holds what the Firestore repos share: one gzip encoding
// for a JSON or protojson payload, one inflate limit, and one document id
// rule. A payload that inflates past the limit is refused, so a stored
// document can not exhaust the process.
package gzstore

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// MaxInflatedBytes bounds the inflated read of one stored payload.
const MaxInflatedBytes = 16 << 20

// MaxStoredBytes bounds one gzip payload. Firestore caps a document at
// 1 MiB, and the other fields use the rest.
const MaxStoredBytes = 900 << 10

// MaxIDBytes is the Firestore limit on one document id.
const MaxIDBytes = 1500

// ErrTooLarge reports a payload that inflates past MaxInflatedBytes.
var ErrTooLarge = fmt.Errorf("gzstore: stored payload inflates past %d bytes", MaxInflatedBytes)

// Marshal gzips raw bytes.
func Marshal(raw []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(raw); err != nil {
		return nil, fmt.Errorf("gzstore: compress: %w", err)
	}
	if err := zw.Close(); err != nil {
		return nil, fmt.Errorf("gzstore: compress: %w", err)
	}
	return buf.Bytes(), nil
}

// Unmarshal inflates a payload, bounded by MaxInflatedBytes. An empty
// payload reads as an empty object, so a document from before the field
// existed opens as a zero value.
func Unmarshal(payload []byte) ([]byte, error) {
	if len(payload) == 0 {
		return []byte("{}"), nil
	}
	zr, err := gzip.NewReader(bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("gzstore: read: %w", err)
	}
	defer func() { _ = zr.Close() }()
	raw, err := io.ReadAll(io.LimitReader(zr, MaxInflatedBytes+1))
	if err != nil {
		return nil, fmt.Errorf("gzstore: read: %w", err)
	}
	if len(raw) > MaxInflatedBytes {
		return nil, ErrTooLarge
	}
	return raw, nil
}

// MarshalJSON encodes v as gzip JSON.
func MarshalJSON(v any) ([]byte, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("gzstore: encode: %w", err)
	}
	return Marshal(raw)
}

// UnmarshalJSON decodes a gzip JSON payload into v.
func UnmarshalJSON(payload []byte, v any) error {
	raw, err := Unmarshal(payload)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, v)
}

// MarshalProto encodes m as gzip protojson, so a proto change reads back
// without a schema migration.
func MarshalProto(m proto.Message) ([]byte, error) {
	raw, err := protojson.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("gzstore: encode: %w", err)
	}
	return Marshal(raw)
}

// UnmarshalProto decodes a gzip protojson payload into m.
func UnmarshalProto(payload []byte, m proto.Message) error {
	raw, err := Unmarshal(payload)
	if err != nil {
		return err
	}
	return protojson.Unmarshal(raw, m)
}

// ErrBadID reports an id that can not name one Firestore document.
var ErrBadID = errors.New("the id is empty, holds a slash, is a dot path, or is longer than 1500 bytes")

// ValidID reports whether id names one document and nothing else: it is
// not empty, holds no slash, is not "." or "..", and fits the Firestore
// limit. An id with a slash addresses another document path.
func ValidID(id string) bool {
	if id == "" || id == "." || id == ".." || len(id) > MaxIDBytes {
		return false
	}
	return !strings.Contains(id, "/")
}
