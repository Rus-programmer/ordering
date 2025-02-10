package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"ordering/dto"
	mockService "ordering/services/mock"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"ordering/util"
)

func TestCreateCustomerAPI(t *testing.T) {
	customer, password := randomCustomer()

	customerRes := dto.CustomerResponse{
		ID:        util.RandomInt(1, 100),
		Username:  customer.Username,
		Role:      customer.Role,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(service *mockService.MockService)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": customer.Username,
				"password": password,
				"role":     customer.Role,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					CreateCustomer(gomock.Any(), customer).
					Times(1).
					Return(customerRes, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"username": customer.Username,
				"password": password,
				"role":     customer.Role,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(1).
					Return(dto.CustomerResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateUsername",
			body: gin.H{
				"username": customer.Username,
				"password": password,
				"role":     customer.Role,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(1).
					Return(dto.CustomerResponse{}, util.ErrUniqueViolation)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
			},
		},
		{
			name: "InvalidUsername",
			body: gin.H{
				"username": "invalid-user#1",
				"password": password,
				"role":     customer.Role,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"username": customer.Username,
				"password": "123",
				"role":     customer.Role,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server, service := newTestServer(t)
			recorder := httptest.NewRecorder()
			tc.buildStubs(service)

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/customers"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(recorder)
		})
	}
}

func TestLoginAPI(t *testing.T) {
	customer, password := randomCustomer()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(service *mockService.MockService)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": customer.Username,
				"password": password,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					Login(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "UserNotFound",
			body: gin.H{
				"username": "NotFound",
				"password": password,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					Login(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(dto.LoginResponse{}, util.ErrRecordNotFound)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "IncorrectPassword",
			body: gin.H{
				"username": customer.Username,
				"password": "incorrect",
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					Login(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(dto.LoginResponse{}, util.ErrInvalidPassword)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"username": customer.Username,
				"password": password,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					Login(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(dto.LoginResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidUsername",
			body: gin.H{
				"username": "invalid-user#1",
				"password": password,
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					Login(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server, service := newTestServer(t)
			recorder := httptest.NewRecorder()
			tc.buildStubs(service)

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(recorder)
		})
	}
}
