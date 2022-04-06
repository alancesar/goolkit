package tracing

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

const (
	_                      = iota
	traceCtxKey contextKey = iota
)

const (
	ReqIDHeader = "x-request-id"

	traceIDHeader      = "x-b3-traceid"
	spanIDHeader       = "x-b3-spanid"
	parentSpanIDHeader = "x-b3-parentspanid"
	sampledIDHeader    = "x-b3-sampled"
)

type (
	contextKey int

	Data struct {
		RequestID string
		B3Tracing
	}

	B3Tracing struct {
		TraceID      string
		SpanID       string
		ParentSpanID string
		SampleID     string
	}
)

func CreateContext(r *http.Request) context.Context {
	var requestID = r.Header.Get(ReqIDHeader)
	if requestID == "" {
		requestID = uuid.New().String()
	}

	data := Data{
		RequestID: requestID,
		B3Tracing: B3Tracing{
			TraceID:      r.Header.Get(traceIDHeader),
			SpanID:       r.Header.Get(spanIDHeader),
			ParentSpanID: r.Header.Get(parentSpanIDHeader),
			SampleID:     r.Header.Get(sampledIDHeader),
		},
	}

	return context.WithValue(r.Context(), traceCtxKey, data)
}

func RetrieveData(ctx context.Context) Data {
	if value := ctx.Value(traceCtxKey); value == nil {
		return Data{}
	} else {
		return value.(Data)
	}
}

func RetrieveRequestID(ctx context.Context) string {
	if trace, ok := ctx.Value(traceCtxKey).(Data); !ok {
		return ""
	} else {
		return trace.RequestID
	}
}
