package server

import (
	"context"
	"net/http"
)

type Http struct {
	*http.Server
}

func NewHttp(handler http.Handler, addr string) *Http {
	return &Http{
		Server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (h *Http) Run(_ context.Context) error {
	if err := h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (h Http) Stop(ctx context.Context) error {
	return h.Shutdown(ctx)
}
