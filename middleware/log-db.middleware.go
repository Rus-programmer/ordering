package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	db "ordering/db/sqlc"
	"ordering/logging"
)

// LogDB creates a gin middleware for inserting logs to db
func (middleware *middlewareImpl) LogDB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ginInfo := logging.ExtractInfoFromGinContext(ctx)

		arg := db.CreateLogParams{
			Method:      ginInfo.Method,
			Path:        ctx.Request.URL.Path,
			StatusCode:  int32(ginInfo.StatusCode),
			ElapsedTime: ginInfo.Duration.String(),
			Time:        ginInfo.BeginTime,
		}

		go func() {
			err := middleware.store.CreateLog(context.Background(), arg)
			if err != nil {
				log.Error().Err(err).Msg("Cannot create log")
			}
		}()
	}
}
