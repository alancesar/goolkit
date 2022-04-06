package main

import (
	"context"
	"go-web-toolkit/application"
	"go-web-toolkit/examples/app/server"
)

func main() {
	rest := server.NewHttp()
	app := application.New(rest)

	// Keep alive until receive syscall.SIGINT or syscall.SIGTERM
	app.Start(context.Background())
}
