package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-web-toolkit/application"
	"go-web-toolkit/application/logger"
	"go-web-toolkit/application/middleware"
	"go-web-toolkit/application/observability"
	"go-web-toolkit/application/server"
	"go-web-toolkit/examples/service"
	"net/http"
	"time"
)

type Service interface {
	ToTime(context.Context, string) (time.Time, error)
}

func main() {
	observer := NewMyObserver("to_date")
	proxy := observability.NewProxy[string, time.Time](observer)
	myService := service.NewMyProxyService(service.NewMyService(), proxy)

	engine := buildGinEngine(myService)
	rest := server.NewHttp(engine, ":8080")
	metrics := server.NewHttp(promhttp.Handler(), ":7777")
	app := application.New(rest, metrics)

	// Keep alive until receive syscall.SIGINT or syscall.SIGTERM
	app.Start(context.Background())
}

func buildGinEngine(s Service) *gin.Engine {
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
	engine.Handle(http.MethodPost, "/time", func(c *gin.Context) {
		_, _ = s.ToTime(c.Request.Context(), "2020-01-05")

		c.JSON(http.StatusCreated, gin.H{
			"message": "data received successfully",
		})
	})

	return engine
}
