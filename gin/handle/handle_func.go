package handle

import "context"

type HandlerFunc[Req, Resp any] func(context.Context, Req) (Resp, error)

func NullHandler[Resp any](handler func(ctx context.Context) (Resp, error)) HandlerFunc[Null, Resp] {
	return func(ctx context.Context, req Null) (Resp, error) {
		return handler(ctx)
	}
}
