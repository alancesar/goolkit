package application

import (
	"context"
	"fmt"
	"go-web-toolkit/application/runner"
	"log"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	runners []runner.Runner
}

func New(servers ...runner.Runner) *Application {
	return &Application{
		runners: servers,
	}
}

func (a Application) Start(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for _, s := range a.runners {
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, s := range a.runners {
		if err := s.Stop(ctx); err != nil {
			log.Println("banco", err)
		}
	}

	fmt.Println("good bye")
}
