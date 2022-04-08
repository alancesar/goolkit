package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-web-toolkit/application"
	"go-web-toolkit/application/logger"
	"go-web-toolkit/application/middleware"
	"go-web-toolkit/application/server"
	"net/http"
)

func main() {
	engine := buildGinEngine()
	rest := server.NewHttp(engine, ":8080")
	app := application.New(rest)

	// Keep alive until receive syscall.SIGINT or syscall.SIGTERM
	app.Start(context.Background())
}

func buildGinEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Tracing, middleware.GinLogger(logger.HTTPRequest, logger.HTTPResponse))

	engine.Handle(http.MethodGet, "/app", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world 🚀",
		})
	})
	engine.Handle(http.MethodPost, "/app", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "data received successfully",
		})
	})

	return engine
}
