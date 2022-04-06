package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-toolkit/application/tracing"
)

func Tracing(ctx *gin.Context) {
	tracingCtx := tracing.CreateContext(ctx.Request)
	ctx.Request = ctx.Request.WithContext(tracingCtx)
	requestID := tracing.RetrieveRequestID(tracingCtx)
	ctx.Writer.Header().Add(tracing.ReqIDHeader, requestID)

	ctx.Next()
}
