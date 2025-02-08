package auth

import (
	"context"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
	"time"
)

func (auth *auth) RenewAccessToken(ctx context.Context, refreshToken string) (dto.RenewAccessTokenResponse, error) {
	refreshPayload, err := auth.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return dto.RenewAccessTokenResponse{}, token.ErrInvalidToken
	}

	session, err := auth.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		return dto.RenewAccessTokenResponse{}, util.ErrBlockedSession
	}

	if session.IsBlocked {
		return dto.RenewAccessTokenResponse{}, util.ErrBlockedSession
	}

	if session.CustomerID != refreshPayload.CustomerID {
		return dto.RenewAccessTokenResponse{}, util.ErrBlockedSession
	}

	if session.RefreshToken != refreshToken {
		return dto.RenewAccessTokenResponse{}, util.ErrMismatchedSessionToken
	}

	if time.Now().After(session.ExpiresAt) {
		return dto.RenewAccessTokenResponse{}, util.ErrSessionExpired
	}

	accessToken, accessPayload, err := auth.tokenMaker.CreateToken(
		refreshPayload.CustomerID,
		refreshPayload.Role,
		auth.config.AccessTokenDuration,
	)
	if err != nil {
		return dto.RenewAccessTokenResponse{}, err
	}

	return dto.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}, nil
}
