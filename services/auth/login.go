package auth

import (
	"context"
	"fmt"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/services/customers"
	"ordering/util"
)

type LoginRequest struct {
	Username string
	Password string
}

func (auth *auth) Login(ctx context.Context, req LoginRequest, clientIp string, userAgent string) (dto.LoginResponse, error) {
	customer, err := auth.store.GetCustomerByUsername(ctx, req.Username)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	err = util.CheckPassword(req.Password, customer.HashedPassword)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("%w: %v", util.ErrInvalidPassword, err)
	}

	tokens, err := auth.getTokens(customer)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	session, err := auth.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           tokens.refreshTokenPayload.ID,
		CustomerID:   customer.ID,
		RefreshToken: tokens.refreshToken,
		UserAgent:    userAgent,
		ClientIp:     clientIp,
		IsBlocked:    false,
		ExpiresAt:    tokens.refreshTokenPayload.ExpiredAt,
	})
	if err != nil {
		return dto.LoginResponse{}, err
	}

	rsp := dto.LoginResponse{
		SessionID:             session.ID,
		AccessToken:           tokens.accessToken,
		AccessTokenExpiresAt:  tokens.accessTokenPayload.ExpiredAt,
		RefreshToken:          tokens.refreshToken,
		RefreshTokenExpiresAt: tokens.refreshTokenPayload.ExpiredAt,
		Customer:              customers.NewCustomerResponse(customer),
	}
	return rsp, err
}
