package auth

import (
	db "ordering/db/sqlc"
	"ordering/token"
)

func (auth *auth) getAccessToken(customer db.Customer) (accessToken string, payload *token.Payload, err error) {
	accessToken, payload, err = auth.tokenMaker.CreateToken(
		customer.ID,
		string(customer.Role),
		auth.config.AccessTokenDuration,
	)
	if err != nil {
		return "", nil, err
	}

	return accessToken, payload, err
}

func (auth *auth) getRefreshToken(customer db.Customer) (accessToken string, payload *token.Payload, err error) {
	accessToken, payload, err = auth.tokenMaker.CreateToken(
		customer.ID,
		string(customer.Role),
		auth.config.RefreshTokenDuration,
	)
	if err != nil {
		return "", nil, err
	}

	return accessToken, payload, err
}

type Tokens struct {
	accessToken         string
	accessTokenPayload  *token.Payload
	refreshToken        string
	refreshTokenPayload *token.Payload
}

func (auth *auth) getTokens(customer db.Customer) (Tokens, error) {
	accessToken, accessTokenPayload, err := auth.getAccessToken(customer)
	if err != nil {
		return Tokens{}, err
	}

	refreshToken, refreshTokenPayload, err := auth.getRefreshToken(customer)
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		accessToken:         accessToken,
		accessTokenPayload:  accessTokenPayload,
		refreshToken:        refreshToken,
		refreshTokenPayload: refreshTokenPayload,
	}, nil
}
