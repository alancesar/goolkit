package logger

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
	"strings"
)

type (
	HttpResponse interface {
		Status() int
		Body() io.Reader
		DataLength() int
	}

	Recorder interface {
		HttpResponse
		http.ResponseWriter
		Request() *http.Request
	}

	Field struct {
		Key   string
		Value interface{}
	}
)

func Request(r *http.Request, additionalFields ...Field) {
	path := r.URL.Path
	method := r.Method

	body := map[string]interface{}{}
	if isJSON(r.Header) {
		body = parseBodyToMap(r.Body)
	}

	logrus.
		WithFields(buildLoggerField(additionalFields)).
		WithFields(logrus.Fields{
			"client_ip":    getIPAddress(r),
			"method":       method,
			"path":         path,
			"referer":      r.Referer(),
			"user_agent":   r.UserAgent(),
			"request_body": body,
		}).
		Infof("request received: [%v] %v", method, path)
}

func Response(recorder Recorder, additionalFields ...Field) {
	body := map[string]interface{}{}
	if isJSON(recorder.Header()) {
		body = parseBodyToMap(recorder.Body())
	}

	method := recorder.Request().Method
	path := recorder.Request().URL.Path

	logrus.
		WithFields(buildLoggerField(additionalFields)).
		WithFields(logrus.Fields{
			"status_code":   recorder.Status(),
			"data_length":   recorder.DataLength(),
			"response_body": body,
		}).
		Infof("response sent:    [%v] %v", method, path)
}

func getIPAddress(r *http.Request) string {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return ""
	}

	remoteIP := net.ParseIP(ip)
	if remoteIP == nil {
		return ""
	}

	return remoteIP.String()
}

func buildLoggerField(additionalFields []Field) logrus.Fields {
	fields := logrus.Fields{}
	for _, field := range additionalFields {
		fields[field.Key] = field.Value
	}
	return fields
}

func parseBodyToMap(r io.Reader) map[string]interface{} {
	var raw map[string]interface{}
	_ = json.NewDecoder(r).Decode(&raw)
	return raw
}

func isJSON(header http.Header) bool {
	return strings.Contains(header.Get("Content-Type"), "application/json")
}
