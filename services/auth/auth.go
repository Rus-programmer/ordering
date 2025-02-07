package auth

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
)

type Auth interface {
	Login(ctx context.Context, req LoginRequest, clientIp string, userAgent string) (dto.LoginResponse, error)
	RenewAccessToken(ctx context.Context, refreshToken string) (dto.RenewAccessTokenResponse, error)
}

type auth struct {
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewAuth(config util.Config, store db.Store, tokenMaker token.Maker) Auth {
	return &auth{
		tokenMaker: tokenMaker,
		store:      store,
		config:     config,
	}
}
