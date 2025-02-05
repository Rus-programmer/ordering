package logging

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type LogInfo struct {
	Duration   time.Duration
	Method     string
	Path       string
	Payload    []byte
	StatusCode int
	Body       string
	BeginTime  time.Time
}

func extractInfoFromGinContext(ctx *gin.Context) *LogInfo {
	// before executing the next handlers
	begin := time.Now()
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}

	// Get payload from request
	payload, _ := io.ReadAll(ctx.Request.Body)
	ctx.Request.Body = io.NopCloser(bytes.NewReader(payload))

	// Get a copy of the body
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
	ctx.Writer = w

	// executes the pending handlers
	ctx.Next()

	return &LogInfo{
		Duration:   time.Since(begin),
		StatusCode: ctx.Writer.Status(),
		Method:     ctx.Request.Method,
		Path:       path,
		Payload:    payload,
		Body:       w.body.String(),
		BeginTime:  begin,
	}
}
