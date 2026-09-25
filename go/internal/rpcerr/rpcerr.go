// Package rpcerr keeps internal error text out of an RPC answer.
package rpcerr

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
)

// Message is the one sentence a caller reads for an Internal error.
const Message = "an internal error occurred"

// Interceptor returns a handler interceptor that replaces the message of
// an Internal error with Message and logs the original error with the
// procedure. A store or decode detail stays in the server log, and never
// reaches a reader, the anonymous share page included. Other codes pass
// unchanged. The client side is untouched.
func Interceptor(log *slog.Logger) connect.Interceptor {
	return &interceptor{log: log}
}

type interceptor struct {
	log *slog.Logger
}

func (i *interceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		res, err := next(ctx, req)
		if req.Spec().IsClient {
			return res, err
		}
		return res, i.mask(ctx, req.Spec().Procedure, err)
	}
}

func (i *interceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return next
}

func (i *interceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		return i.mask(ctx, conn.Spec().Procedure, next(ctx, conn))
	}
}

// mask returns err, or a fixed Internal error in place of an Internal
// one.
func (i *interceptor) mask(ctx context.Context, procedure string, err error) error {
	var ce *connect.Error
	if !errors.As(err, &ce) || ce.Code() != connect.CodeInternal {
		return err
	}
	i.log.ErrorContext(ctx, "internal error", "procedure", procedure, "err", err)
	return connect.NewError(connect.CodeInternal, errors.New(Message))
}
