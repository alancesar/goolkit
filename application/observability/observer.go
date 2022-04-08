package observability

import (
	"context"
	"time"
)

type (
	Observer interface {
		Start(ctx context.Context)
		Success(ctx context.Context, response any)
		Error(ctx context.Context, err error)
		Finish(ctx context.Context, duration time.Duration)
	}

	Wrapper struct {
		observers []Observer
	}
)

func NewWrapper[Req, Res any](observers ...Observer) *Wrapper {
	return &Wrapper{
		observers: observers,
	}
}

func (o Wrapper) Start(ctx context.Context) {
	for _, observer := range o.observers {
		observer.Start(ctx)
	}
}

func (o Wrapper) Success(ctx context.Context, response any) {
	for _, observer := range o.observers {
		observer.Success(ctx, response)
	}
}

func (o Wrapper) Error(ctx context.Context, err error) {
	for _, observer := range o.observers {
		observer.Error(ctx, err)
	}
}

func (o Wrapper) Finish(ctx context.Context, duration time.Duration) {
	for _, observer := range o.observers {
		observer.Finish(ctx, duration)
	}
}
