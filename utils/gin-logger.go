package util

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"time"
)

var (
	DurationFieldName   = "elapsed"
	MethodFieldName     = "method"
	PathFieldName       = "path"
	PayloadFieldName    = "payload"
	statusCodeFieldName = "status_code"
	BodyFieldName       = "body"
)

// GinLogger is a gin middleware which use zerolog
func GinLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get zerolog
		z := &log.Logger

		// return if zerolog is disabled
		if z.GetLevel() == zerolog.Disabled {
			ctx.Next()
			return
		}

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

		// after executing the handlers
		duration := time.Since(begin)
		statusCode := ctx.Writer.Status()

		//
		var event *zerolog.Event
		var eventError bool
		var eventWarn bool

		// set message level
		if statusCode >= 400 && statusCode < 500 {
			eventWarn = true
			event = z.Warn()
		} else if statusCode >= 500 {
			eventError = true
			event = z.Error()
		} else {
			event = z.Info()
		}

		// Path field
		if len(path) > 0 {
			event.Str(PathFieldName, path)
		}

		// Method field
		event.Str(MethodFieldName, ctx.Request.Method)

		// statusCode field
		event.Int(statusCodeFieldName, statusCode)

		// Payload field
		if len(payload) > 0 {
			event.Str(PayloadFieldName, string(payload))
		}

		// Duration field
		var durationFieldName string
		switch zerolog.DurationFieldUnit {
		case time.Nanosecond:
			durationFieldName = DurationFieldName + "_ns"
		case time.Microsecond:
			durationFieldName = DurationFieldName + "_us"
		case time.Millisecond:
			durationFieldName = DurationFieldName + "_ms"
		case time.Second:
			durationFieldName = DurationFieldName + "_sec"
		case time.Minute:
			durationFieldName = DurationFieldName + "_min"
		case time.Hour:
			durationFieldName = DurationFieldName + "_hr"
		default:
			z.Error().Interface("zerolog.DurationFieldUnit", zerolog.DurationFieldUnit).Msg("unknown value for DurationFieldUnit")
			durationFieldName = DurationFieldName
		}
		event.Dur(durationFieldName, duration)

		// Body field
		if len(w.body.String()) > 0 {
			event.Str(BodyFieldName, w.body.String())
		}

		// Message
		message := ctx.Errors.String()
		if message == "" {
			message = "received a HTTP request"
		}

		// post the message
		if eventError {
			event.Msg(message)
		} else if eventWarn {
			event.Msg(message)
		} else {
			event.Msg(message)
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r responseBodyWriter) WriteString(s string) (n int, err error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}
