package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/middleware"
	"ordering/services/customers"
	mockService "ordering/services/mock"
	"ordering/services/products"
	test_utils "ordering/test-utils"
	"ordering/token"
	"ordering/util"
	"testing"
	"time"
)

func TestListProductsAPI(t *testing.T) {
	customer, _ := randomCustomer()
	ID := util.RandomInt(1, 1000)

	n := 5
	listProducts := make([]dto.ProductResponse, n)
	for i := 0; i < n; i++ {
		listProducts[i] = randomProduct()
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(service *mockService.MockService)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, ID, customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				arg := products.ListProductRequest{
					Limit:  int32(n),
					Offset: 0,
				}

				service.EXPECT().
					ListProducts(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(listProducts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProducts(t, recorder.Body, listProducts)
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
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
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
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, ID, customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					ListProducts(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]dto.ProductResponse{}, sql.ErrConnDone)
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
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, ID, customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
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
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, ID, customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
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
			server, service := newTestServer(t)
			recorder := httptest.NewRecorder()
			tc.buildStubs(service)

			url := "/products"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// Add query parameters to request URL
			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()
			tc.setupAuth(t, request, service.GetTokenMaker())

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomProduct() dto.ProductResponse {
	return dto.ProductResponse{
		ID:       util.RandomInt(1, 1000),
		Name:     util.RandomString(6),
		Price:    util.RandomInt(50, 100),
		Quantity: util.RandomInt(1, 10),
	}
}

func requireBodyMatchProducts(t *testing.T, body *bytes.Buffer, products []dto.ProductResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProducts []dto.ProductResponse
	err = json.Unmarshal(data, &gotProducts)
	require.NoError(t, err)
	require.Equal(t, products, gotProducts)
}

func TestGetProductAPI(t *testing.T) {
	customer, _ := randomCustomer()
	product := randomProduct()
	ID := util.RandomInt(1, 1000)

	testCases := []struct {
		name          string
		productID     int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(service *mockService.MockService)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productID: product.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, ID, customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
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
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
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
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, ID, customer.Role, time.Minute)
			},

			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(dto.ProductResponse{}, util.ErrRecordNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			productID: product.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, ID, customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(dto.ProductResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			productID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, ID, customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
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
			server, service := newTestServer(t)
			recorder := httptest.NewRecorder()
			tc.buildStubs(service)

			url := fmt.Sprintf("/products/%d", tc.productID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, service.GetTokenMaker())
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchProduct(t *testing.T, body *bytes.Buffer, product dto.ProductResponse) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotProduct dto.ProductResponse
	err = json.Unmarshal(data, &gotProduct)
	require.NoError(t, err)
	require.Equal(t, product, gotProduct)
}

func randomCustomer() (customer customers.CreateCustomer, password string) {
	password = util.RandomString(6)

	customer = customers.CreateCustomer{
		Username: util.RandomCustomer(),
		Password: password,
		Role:     db.UserRoleUser,
	}
	return
}

func TestDeleteProductAPI(t *testing.T) {
	customer, _ := randomCustomer()
	testCases := []struct {
		name          string
		productID     string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(service *mockService.MockService)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productID: "1",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(1), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					DeleteProduct(gomock.Any(), int64(1)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "NotFound",
			productID: "999",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(999), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					DeleteProduct(gomock.Any(), int64(999)).
					Times(1).
					Return(util.ErrRecordNotFound)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			productID: "abc",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(999), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			productID: "2",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(2), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					DeleteProduct(gomock.Any(), int64(2)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "Unauthorized",
			productID: "2",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					DeleteProduct(gomock.Any(), int64(2)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
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

			url := "/products/" + tc.productID
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)
			tc.setupAuth(t, request, service.GetTokenMaker())

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestUpdateProductAPI(t *testing.T) {
	product := randomProduct()
	customer, _ := randomCustomer()

	updatedProduct := product
	updatedProduct.Name = "Updated Name"
	updatedProduct.Price = product.Price + 100
	updatedProduct.Quantity = product.Quantity + 10

	testCases := []struct {
		name          string
		productID     int64
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(service *mockService.MockService)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productID: product.ID,
			body: gin.H{
				"name":     updatedProduct.Name,
				"price":    updatedProduct.Price,
				"quantity": updatedProduct.Quantity,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(2), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					UpdateProduct(gomock.Any(), product.ID, products.UpdateProduct{
						Name:     updatedProduct.Name,
						Price:    updatedProduct.Price,
						Quantity: updatedProduct.Quantity,
					}).
					Times(1).
					Return(updatedProduct, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			productID: product.ID,
			body: gin.H{
				"name": "New Name",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(2), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					UpdateProduct(gomock.Any(), product.ID, gomock.Any()).
					Times(1).
					Return(dto.ProductResponse{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "NotFound",
			productID: 99999,
			body: gin.H{
				"name": "New Name",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(2), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					UpdateProduct(gomock.Any(), int64(99999), gomock.Any()).
					Times(1).
					Return(dto.ProductResponse{}, util.ErrRecordNotFound)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InvalidProductID",
			productID: -1,
			body: gin.H{
				"name": "New Name",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(2), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "InvalidPrice",
			productID: product.ID,
			body: gin.H{
				"price": -10,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				test_utils.AddAuthorization(t, request, tokenMaker, middleware.AuthorizationTypeBearer, int64(2), customer.Role, time.Minute)
			},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "Unauthorized",
			productID: product.ID,
			body: gin.H{
				"price": 10,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {},
			buildStubs: func(service *mockService.MockService) {
				service.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
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

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := fmt.Sprintf("/products/%d", tc.productID)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, service.GetTokenMaker())
			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(recorder)
		})
	}
}
