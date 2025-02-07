package dto

import (
	"github.com/google/uuid"
	"time"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	SessionID             uuid.UUID        `json:"session_id"`
	AccessToken           string           `json:"access_token"`
	AccessTokenExpiresAt  time.Time        `json:"access_token_expires_at"`
	RefreshToken          string           `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time        `json:"refresh_token_expires_at"`
	Customer              CustomerResponse `json:"customer"`
}

type RenewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}
