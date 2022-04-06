package tracing

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type contextKey string

const (
	ReqIDHeader                   = "x-request-id"
	TraceIDHeader                 = "x-b3-traceid"
	SpanIDHeader                  = "x-b3-spanid"
	ParentSpanIDHeader            = "x-b3-parentspanid"
	SampledIDHeader               = "x-b3-sampled"
	TraceCTXKey        contextKey = "tracing-context"
)

type Tracing struct {
	RequestID string
	B3Tracing
}

type B3Tracing struct {
	TraceID      string
	SpanID       string
	ParentSpanID string
	SampleID     string
}

func WithTracing(ctx context.Context, tracing Tracing) context.Context {
	return context.WithValue(ctx, TraceCTXKey, tracing)
}

func FromContext(ctx context.Context) Tracing {
	v := ctx.Value(TraceCTXKey)
	if v == nil {
		return Tracing{}
	}
	return v.(Tracing)
}

func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = BindRequestContext(c.Request)
		c.Next()
	}
}

func BindRequestContext(r *http.Request) *http.Request {
	var requestID = r.Header.Get(ReqIDHeader)
	if requestID == "" {
		requestID = uuid.New().String()
	}

	t := Tracing{
		RequestID: requestID,
		B3Tracing: B3Tracing{
			TraceID:      r.Header.Get(TraceIDHeader),
			SpanID:       r.Header.Get(SpanIDHeader),
			ParentSpanID: r.Header.Get(ParentSpanIDHeader),
			SampleID:     r.Header.Get(SampledIDHeader),
		},
	}
	return r.WithContext(WithTracing(r.Context(), t))
}

func GetRequestIDFromContext(ctx context.Context) string {
	if trace, ok := ctx.Value(TraceCTXKey).(Tracing); !ok {
		return ""
	} else {
		return trace.RequestID
	}
}
