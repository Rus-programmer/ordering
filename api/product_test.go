package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"ordering/common"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	test_utils "ordering/test-utils"
	"ordering/token"
	util "ordering/util"
	"testing"
	"time"
)

func TestListProductsAPI(t *testing.T) {
	user, _ := randomUser(t)

	n := 5
	products := make([]db.Product, n)
	for i := 0; i < n; i++ {
		products[i] = randomProduct()
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, common.AuthorizationTypeBearer, user.ID, user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListProductsParams{
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().
					ListProducts(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(products, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProducts(t, recorder.Body, products)
			},
		},
		{
			name: "NoAuthorization",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, common.AuthorizationTypeBearer, user.ID, user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, common.AuthorizationTypeBearer, user.ID, user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			query: Query{
				pageID:   1,
				pageSize: 100000,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, common.AuthorizationTypeBearer, user.ID, user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), gomock.Any()).
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

			url := "/products"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomProduct() db.Product {
	return db.Product{
		ID:       util.RandomInt(1, 1000),
		Name:     util.RandomString(6),
		Price:    util.RandomInt(50, 100),
		Quantity: util.RandomInt(1, 10),
	}
}

func requireBodyMatchProducts(t *testing.T, body *bytes.Buffer, products []db.Product) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProducts []db.Product
	err = json.Unmarshal(data, &gotProducts)
	require.NoError(t, err)
	require.Equal(t, products, gotProducts)
}

func TestGetProductAPI(t *testing.T) {
	user, _ := randomUser(t)
	product := randomProduct()

	testCases := []struct {
		name          string
		productID     int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productID: product.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, common.AuthorizationTypeBearer, user.ID, user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProduct(t, recorder.Body, product)
			},
		},
		{
			name:      "NoAuthorization",
			productID: product.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:      "NotFound",
			productID: product.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, common.AuthorizationTypeBearer, user.ID, user.Role, time.Minute)
			},

			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(db.Product{}, db.ErrRecordNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			productID: product.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, common.AuthorizationTypeBearer, user.ID, user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(db.Product{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			productID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, common.AuthorizationTypeBearer, user.ID, user.Role, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
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

			url := fmt.Sprintf("/products/%d", tc.productID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchProduct(t *testing.T, body *bytes.Buffer, product db.Product) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProduct db.Product
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, product, gotProduct)
}
