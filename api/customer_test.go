package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/util"
)

type eqCreateCustomerParamsMatcher struct {
	arg      db.CreateCustomerParams
	password string
}

func (e eqCreateCustomerParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateCustomerParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateCustomerParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateCustomerParams, password string) gomock.Matcher {
	return eqCreateCustomerParamsMatcher{arg, password}
}

func TestCreateCustomerAPI(t *testing.T) {
	customer, password := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": customer.Username,
				"password": password,
				"role":     customer.Role,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateCustomerParams{
					Username: customer.Username,
					Role:     customer.Role,
				}
				store.EXPECT().
					CreateCustomer(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(customer, nil)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Customer{}, sql.ErrConnDone)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Customer{}, db.ErrUniqueViolation)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(recorder)
		})
	}
}

func TestLoginAPI(t *testing.T) {
	customer, password := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": customer.Username,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomerByUsername(gomock.Any(), gomock.Eq(customer.Username)).
					Times(1).
					Return(customer, nil)
				store.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomerByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Customer{}, db.ErrRecordNotFound)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomerByUsername(gomock.Any(), gomock.Eq(customer.Username)).
					Times(1).
					Return(customer, nil)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomerByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Customer{}, sql.ErrConnDone)
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
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomerByUsername(gomock.Any(), gomock.Any()).
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

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

func randomUser(t *testing.T) (user db.Customer, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.Customer{
		Username:       util.RandomCustomer(),
		HashedPassword: hashedPassword,
		Role:           db.UserRoleUser,
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.Customer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotCustomer db.Customer
	err = json.Unmarshal(data, &gotCustomer)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotCustomer.Username)
	require.Empty(t, gotCustomer.HashedPassword)
}
