package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-web-toolkit/application/logger"
	"go-web-toolkit/application/middleware"
	"net/http"
)

type Http struct {
	server *http.Server
}

func NewHttp() *Http {
	return &Http{}
}

func (h *Http) Run(_ context.Context) error {
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

	h.server = &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	if err := h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (h Http) Stop(ctx context.Context) error {
	return h.Shutdown(ctx)
}
