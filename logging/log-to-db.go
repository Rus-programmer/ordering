package logging

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	db "ordering/db/sqlc"
)

func LogToDB(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ginInfo := extractInfoFromGinContext(ctx)

		arg := db.CreateLogParams{
			Method:      ginInfo.Method,
			Path:        ginInfo.Path,
			StatusCode:  int32(ginInfo.StatusCode),
			ElapsedTime: ginInfo.Duration.String(),
			Time:        ginInfo.BeginTime,
		}

		go func() {
			err := store.CreateLog(context.Background(), arg)
			if err != nil {
				log.Error().Err(err).Msg("Cannot create log")
			}
		}()
	}
}
