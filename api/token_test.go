package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"ordering/dto"
	mockService "ordering/services/mock"
	"ordering/util"
	"testing"
	"time"
)

func TestRenewAccessTokenAPI(t *testing.T) {
	refreshToken := util.RandomString(4)

	res := dto.RenewAccessTokenResponse{
		AccessToken:          util.RandomString(4),
		AccessTokenExpiresAt: time.Now().Add(time.Minute),
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(service *mockService.MockService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					RenewAccessToken(gomock.Any(), gomock.Eq(refreshToken)).
					Times(1).
					Return(res, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatch(t, recorder.Body, res)
			},
		},
		{
			name: "RefreshTokenRequired",
			body: gin.H{},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					RenewAccessToken(gomock.Any(), gomock.Eq(refreshToken)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ErrExpiredToken",
			body: gin.H{
				"refresh_token": refreshToken,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					RenewAccessToken(gomock.Any(), gomock.Eq(refreshToken)).
					Times(1).
					Return(dto.RenewAccessTokenResponse{}, util.ErrExpiredToken)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server, service := newTestServer(t)
			recorder := httptest.NewRecorder()
			tc.buildStubs(service)

			data, err := json.Marshal(tc.body)
			request, err := http.NewRequest(http.MethodPost, "/tokens/renew_access", bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatch(t *testing.T, body *bytes.Buffer, res dto.RenewAccessTokenResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotResponse dto.RenewAccessTokenResponse
	err = json.Unmarshal(data, &gotResponse)
	require.NoError(t, err)
	require.Equal(t, res.AccessToken, gotResponse.AccessToken)
	require.WithinDuration(t, res.AccessTokenExpiresAt, gotResponse.AccessTokenExpiresAt, time.Second)
}
