package observability

import (
	"context"
	"time"
)

type (
	Runner[Req, Res any] func(ctx context.Context, req Req) (Res, error)

	Proxy[Req, Res any] struct {
		observer Observer[Req, Res]
	}
)

func NewProxy[Req, Res any](observer Observer[Req, Res]) Proxy[Req, Res] {
	return Proxy[Req, Res]{
		observer: observer,
	}
}

func (p Proxy[Req, Res]) Run(ctx context.Context, req Req, runner Runner[Req, Res]) (res Res, err error) {
	start := time.Now()
	p.observer.Start(ctx, req)

	defer func() {
		duration := time.Since(start)
		p.observer.Finish(ctx, duration)
	}()

	if res, err = runner(ctx, req); err != nil {
		p.observer.Error(ctx, err)
		return res, err
	}

	p.observer.Success(ctx, res)
	return res, nil
}
