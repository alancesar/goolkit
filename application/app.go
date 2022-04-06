package application

import (
	"context"
	"fmt"
	"go-web-toolkit/application/runner"
	"log"
	"os/signal"
	"syscall"
)

type Application struct {
	servers []runner.Runner
}

func New(servers ...runner.Runner) *Application {
	return &Application{
		servers: servers,
	}
}

func (a Application) Start(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for _, s := range a.servers {
		go func(s runner.Runner) {
			if err := s.Run(ctx); err != nil {
				log.Fatalln(err)
			}
		}(s)
	}

	fmt.Println("all systems go!")

	<-ctx.Done()
	stop()

	fmt.Println("shutting down...")

	for _, s := range a.servers {
		if err := s.Stop(ctx); err != nil {
			log.Println(err)
		}
	}
}
