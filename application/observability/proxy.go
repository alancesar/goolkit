package observability

import (
	"context"
	"time"
)

type (
	Runner[T any] func() (T, error)

	Proxy[T any] struct {
		observer Observer
	}
)

func NewProxy[T any](observer Observer) Proxy[T] {
	return Proxy[T]{
		observer: observer,
	}
}

func (p Proxy[T]) Run(ctx context.Context, runner Runner[T]) (t T, err error) {
	start := time.Now()
	p.observer.Start(ctx)

	defer func() {
		duration := time.Since(start)
		p.observer.Finish(ctx, duration)
	}()

	if t, err = runner(); err != nil {
		p.observer.Error(ctx, err)
		return t, err
	}

	p.observer.Success(ctx, t)
	return t, nil
}
