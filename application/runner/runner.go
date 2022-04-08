package runner

import "context"

type Runner interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context) error
}
