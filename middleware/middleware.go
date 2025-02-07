package middleware

import (
	"github.com/gin-gonic/gin"
	db "ordering/db/sqlc"
	"ordering/token"
)

type Middleware interface {
	Auth() gin.HandlerFunc
	LogDB() gin.HandlerFunc
}

type middlewareImpl struct {
	store      db.Store
	tokenMaker token.Maker
}

func NewMiddleware(store db.Store, tokenMaker token.Maker) Middleware {
	return &middlewareImpl{
		tokenMaker: tokenMaker,
		store:      store,
	}
}
