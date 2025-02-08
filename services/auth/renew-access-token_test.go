package auth

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
	"testing"
	"time"
)

func TestRenewAccessToken(t *testing.T) {
	customerID := util.RandomInt(1, 1000)
	refreshToken := "valid_refresh_token"
	expiredRefreshToken := "expired_refresh_token"
	invalidRefreshToken := "invalid_refresh_token"

	refreshPayload := &token.Payload{
		ID:         uuid.New(),
		CustomerID: customerID,
		Role:       "user",
		ExpiredAt:  time.Now().Add(time.Hour),
	}

	expiredPayload := *refreshPayload
	expiredPayload.ExpiredAt = time.Now().Add(-time.Hour)

	testCases := []struct {
		name          string
		token         string
		buildStubs    func(store *mockdb.MockStore, tokenMaker *token.MockMaker)
		checkResponse func(resp dto.RenewAccessTokenResponse, err error)
	}{
		{
			name:  "OK",
			token: refreshToken,
			buildStubs: func(store *mockdb.MockStore, tokenMaker *token.MockMaker) {
				tokenMaker.EXPECT().VerifyToken(refreshToken).Times(1).Return(refreshPayload, nil)
				store.EXPECT().GetSession(gomock.Any(), refreshPayload.ID).Times(1).Return(db.Session{
					ID:           refreshPayload.ID,
					CustomerID:   refreshPayload.CustomerID,
					RefreshToken: refreshToken,
					ExpiresAt:    time.Now().Add(time.Hour),
				}, nil)
				tokenMaker.EXPECT().CreateToken(refreshPayload.CustomerID, refreshPayload.Role, gomock.Any()).Times(1).
					Return("new_access_token", &token.Payload{ExpiredAt: time.Now().Add(time.Minute)}, nil)
			},
			checkResponse: func(resp dto.RenewAccessTokenResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, resp.AccessToken)
			},
		},
		{
			name:  "InvalidToken",
			token: invalidRefreshToken,
			buildStubs: func(store *mockdb.MockStore, tokenMaker *token.MockMaker) {
				tokenMaker.EXPECT().VerifyToken(invalidRefreshToken).Times(1).Return(nil, token.ErrInvalidToken)
			},
			checkResponse: func(resp dto.RenewAccessTokenResponse, err error) {
				require.Error(t, err)
				require.Equal(t, token.ErrInvalidToken, err)
			},
		},
		{
			name:  "SessionNotFound",
			token: refreshToken,
			buildStubs: func(store *mockdb.MockStore, tokenMaker *token.MockMaker) {
				tokenMaker.EXPECT().VerifyToken(refreshToken).Times(1).Return(refreshPayload, nil)
				store.EXPECT().GetSession(gomock.Any(), refreshPayload.ID).Times(1).Return(db.Session{}, util.ErrBlockedSession)
			},
			checkResponse: func(resp dto.RenewAccessTokenResponse, err error) {
				require.Error(t, err)
				require.Equal(t, util.ErrBlockedSession, err)
			},
		},
		{
			name:  "BlockedSession",
			token: refreshToken,
			buildStubs: func(store *mockdb.MockStore, tokenMaker *token.MockMaker) {
				tokenMaker.EXPECT().VerifyToken(refreshToken).Times(1).Return(refreshPayload, nil)
				store.EXPECT().GetSession(gomock.Any(), refreshPayload.ID).Times(1).Return(db.Session{
					ID:         refreshPayload.ID,
					CustomerID: refreshPayload.CustomerID,
					IsBlocked:  true,
				}, nil)
			},
			checkResponse: func(resp dto.RenewAccessTokenResponse, err error) {
				require.Error(t, err)
				require.Equal(t, util.ErrBlockedSession, err)
			},
		},
		{
			name:  "MismatchedToken",
			token: refreshToken,
			buildStubs: func(store *mockdb.MockStore, tokenMaker *token.MockMaker) {
				tokenMaker.EXPECT().VerifyToken(refreshToken).Times(1).Return(refreshPayload, nil)
				store.EXPECT().GetSession(gomock.Any(), refreshPayload.ID).Times(1).Return(db.Session{
					ID:           refreshPayload.ID,
					CustomerID:   refreshPayload.CustomerID,
					RefreshToken: "another_refresh_token",
					ExpiresAt:    time.Now().Add(time.Hour),
				}, nil)
			},
			checkResponse: func(resp dto.RenewAccessTokenResponse, err error) {
				require.Error(t, err)
				require.Equal(t, util.ErrMismatchedSessionToken, err)
			},
		},
		{
			name:  "ExpiredSession",
			token: expiredRefreshToken,
			buildStubs: func(store *mockdb.MockStore, tokenMaker *token.MockMaker) {
				tokenMaker.EXPECT().VerifyToken(expiredRefreshToken).Times(1).Return(&expiredPayload, nil)
				store.EXPECT().GetSession(gomock.Any(), expiredPayload.ID).Times(1).Return(db.Session{
					ID:           expiredPayload.ID,
					CustomerID:   expiredPayload.CustomerID,
					RefreshToken: expiredRefreshToken,
					ExpiresAt:    time.Now().Add(-time.Hour),
				}, nil)
			},
			checkResponse: func(resp dto.RenewAccessTokenResponse, err error) {
				require.Error(t, err)
				require.Equal(t, util.ErrSessionExpired, err)
			},
		},
		{
			name:  "CreateTokenError",
			token: refreshToken,
			buildStubs: func(store *mockdb.MockStore, tokenMaker *token.MockMaker) {
				tokenMaker.EXPECT().VerifyToken(refreshToken).Times(1).Return(refreshPayload, nil)
				store.EXPECT().GetSession(gomock.Any(), refreshPayload.ID).Times(1).Return(db.Session{
					ID:           refreshPayload.ID,
					CustomerID:   refreshPayload.CustomerID,
					RefreshToken: refreshToken,
					ExpiresAt:    time.Now().Add(time.Hour),
				}, nil)
				tokenMaker.EXPECT().CreateToken(refreshPayload.CustomerID, refreshPayload.Role, gomock.Any()).Times(1).
					Return("", nil, errors.New("token error"))
			},
			checkResponse: func(resp dto.RenewAccessTokenResponse, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "token error")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			auth, store, tokenMaker := newTestAuth(t)

			tc.buildStubs(store, tokenMaker)
			resp, err := auth.RenewAccessToken(context.Background(), tc.token)

			tc.checkResponse(resp, err)
		})
	}
}
