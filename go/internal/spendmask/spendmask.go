// Package spendmask removes the model spend from each answer of the API
// to a caller that is not the admin (D-1148). The spend is the message
// mtg.v1.Usage: the total of a session, the total of a chat turn, and the
// spend of one build inside a deck. The admin claim of D-1076 keeps it.
//
// The interceptor clears a copy and never the message of the service. A
// service can store the message that it sends, for example the deck of
// a chat turn, so a change in place would remove the spend from storage.
package spendmask

import (
	"context"
	"errors"
	"reflect"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nkramber/decktome/go/internal/auth"
)

// usageName is the full name of the spend message.
const usageName protoreflect.FullName = "mtg.v1.Usage"

// Interceptor clears the spend of each answer unless the context holds
// the admin mark. Put it after the auth interceptor, which sets the mark.
func Interceptor() connect.Interceptor { return interceptor{} }

type interceptor struct{}

func (interceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		res, err := next(ctx, req)
		if err != nil || res == nil || req.Spec().IsClient || auth.IsAdmin(ctx) {
			return res, err
		}
		m, ok := res.Any().(proto.Message)
		if !ok || !Has(m) {
			return res, err
		}
		// A Response holds its message in the exported field Msg. The
		// copy goes there, so the headers and the trailers stay.
		field := reflect.ValueOf(res).Elem().FieldByName("Msg")
		if !field.IsValid() || !field.CanSet() {
			return nil, connect.NewError(connect.CodeInternal, errNoField)
		}
		field.Set(reflect.ValueOf(Masked(m)))
		return res, nil
	}
}

func (interceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (interceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		if auth.IsAdmin(ctx) {
			return next(ctx, conn)
		}
		return next(ctx, maskedConn{conn})
	}
}

type maskedConn struct{ connect.StreamingHandlerConn }

// Send sends a masked copy. A chat event that held the spend alone is
// empty after the mask, so the stream skips it.
func (c maskedConn) Send(msg any) error {
	if m, ok := msg.(proto.Message); ok && Has(m) {
		masked := Masked(m)
		if proto.Size(masked) == 0 {
			return nil
		}
		return c.StreamingHandlerConn.Send(masked)
	}
	return c.StreamingHandlerConn.Send(msg)
}

var errNoField = errors.New("spendmask: the response holds no field Msg")

// Has reports whether m holds a set spend message at any depth.
func Has(m proto.Message) bool {
	return walk(m.ProtoReflect(), false)
}

// Masked answers a copy of m with each spend message cleared.
func Masked(m proto.Message) proto.Message {
	c := proto.Clone(m)
	walk(c.ProtoReflect(), true)
	return c
}

// walk finds each set field of the spend type under m. With wipe, it
// clears each one. It answers whether it found one.
func walk(m protoreflect.Message, wipe bool) bool {
	found := false
	var spent []protoreflect.FieldDescriptor
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Message() == nil {
			return true
		}
		switch {
		case fd.IsList():
			l := v.List()
			for i := 0; i < l.Len(); i++ {
				if walk(l.Get(i).Message(), wipe) {
					found = true
				}
			}
		case fd.IsMap():
			if fd.MapValue().Message() == nil {
				return true
			}
			v.Map().Range(func(_ protoreflect.MapKey, mv protoreflect.Value) bool {
				if walk(mv.Message(), wipe) {
					found = true
				}
				return true
			})
		case fd.Message().FullName() == usageName:
			found = true
			spent = append(spent, fd)
		default:
			if walk(v.Message(), wipe) {
				found = true
			}
		}
		// A search without wipe stops at the first find.
		return wipe || !found
	})
	if wipe {
		for _, fd := range spent {
			m.Clear(fd)
		}
	}
	return found
}
