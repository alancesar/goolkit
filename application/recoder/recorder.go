package recoder

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type GinRecorder struct {
	gin.ResponseWriter
	body *bytes.Buffer
	req  *http.Request
}

func NewGinRecorder(ctx *gin.Context) *GinRecorder {
	return &GinRecorder{
		ResponseWriter: ctx.Writer,
		req:            ctx.Request,
		body:           new(bytes.Buffer),
	}
}

func (w GinRecorder) Request() *http.Request {
	return w.req
}

func (w GinRecorder) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w GinRecorder) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (w GinRecorder) Body() io.Reader {
	return w.body
}

func (w GinRecorder) DataLength() int {
	return w.Size()
}
