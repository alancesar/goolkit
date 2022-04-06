package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-toolkit/application/logger"
	"go-web-toolkit/application/recoder"
	"go-web-toolkit/application/tracing"
	"net/http"
	"os"
	"time"
)

type (
	RequestLoggerFn func(r *http.Request, additionalFields ...logger.Field)
	ResponseLogger  func(recorder logger.Recorder, additionalFields ...logger.Field)
)

func GinLogger(req RequestLoggerFn, res ResponseLogger) gin.HandlerFunc {
	hostname := logger.Field{
		Key:   "hostname",
		Value: getHostname(),
	}

	return func(ctx *gin.Context) {
		start := time.Now()

		requestID := logger.Field{
			Key:   "request_id",
			Value: tracing.RetrieveRequestID(ctx.Request.Context()),
		}

		req(ctx.Request, hostname, requestID)

		recorder := recoder.NewGinRecorder(ctx)
		ctx.Writer = recorder
		ctx.Next()

		latency := logger.Field{
			Key:   "latency",
			Value: time.Since(start).Milliseconds(),
		}

		res(recorder, latency, requestID)
	}
}

func getHostname() string {
	if hostname, err := os.Hostname(); err != nil {
		return "unknown"
	} else {
		return hostname
	}
}
