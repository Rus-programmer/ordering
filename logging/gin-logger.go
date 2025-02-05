package logging

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

		ginInfo := extractInfoFromGinContext(ctx)

		var event *zerolog.Event
		var eventError bool
		var eventWarn bool

		// set message level
		if ginInfo.StatusCode >= 400 && ginInfo.StatusCode < 500 {
			eventWarn = true
			event = z.Warn()
		} else if ginInfo.StatusCode >= 500 {
			eventError = true
			event = z.Error()
		} else {
			event = z.Info()
		}

		// Path field
		if len(ginInfo.Path) > 0 {
			event.Str(PathFieldName, ginInfo.Path)
		}

		// Method field
		event.Str(MethodFieldName, ctx.Request.Method)

		// statusCode field
		event.Int(statusCodeFieldName, ginInfo.StatusCode)

		// Payload field
		if len(ginInfo.Payload) > 0 {
			event.Str(PayloadFieldName, string(ginInfo.Payload))
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
		event.Dur(durationFieldName, ginInfo.Duration)

		// Body field
		if len(ginInfo.Body) > 0 {
			event.Str(BodyFieldName, ginInfo.Body)
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
