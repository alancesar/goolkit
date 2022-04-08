package observability

import (
	"context"
	"time"
)

type (
	Observer[Req, Res any] interface {
		Start(ctx context.Context, req Req)
		Success(ctx context.Context, res Res)
		Error(ctx context.Context, err error)
		Finish(ctx context.Context, duration time.Duration)
	}

	Wrapper[Req, Res any] struct {
		observers []Observer[Req, Res]
	}
)

func NewWrapper[Req, Res any](observers ...Observer[Req, Res]) *Wrapper[Req, Res] {
	return &Wrapper[Req, Res]{
		observers: observers,
	}
}

func (o Wrapper[Req, Res]) Start(ctx context.Context, req Req) {
	for _, observer := range o.observers {
		observer.Start(ctx, req)
	}
}

func (o Wrapper[Req, Res]) Success(ctx context.Context, res Res) {
	for _, observer := range o.observers {
		observer.Success(ctx, res)
	}
}

func (o Wrapper[Req, Res]) Error(ctx context.Context, err error) {
	for _, observer := range o.observers {
		observer.Error(ctx, err)
	}
}

func (o Wrapper[Req, Res]) Finish(ctx context.Context, duration time.Duration) {
	for _, observer := range o.observers {
		observer.Finish(ctx, duration)
	}
}
