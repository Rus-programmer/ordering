package products

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/dto"
	"testing"
)

func TestListProducts(t *testing.T) {
	limit := int32(5)
	offset := int32(5)

	product1 := randomProduct()
	product2 := randomProduct()

	expectedProducts := []db.Product{
		product1,
		product2,
	}

	testCases := []struct {
		name          string
		req           ListProductRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(products []dto.ProductResponse, err error)
	}{
		{
			name: "OK",
			req:  ListProductRequest{Limit: limit, Offset: offset},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), db.ListProductsParams{Limit: limit, Offset: offset}).
					Times(1).
					Return(expectedProducts, nil)
			},
			checkResponse: func(products []dto.ProductResponse, err error) {
				require.NoError(t, err)
				require.Len(t, products, len(expectedProducts))

				for i, p := range products {
					require.Equal(t, expectedProducts[i].ID, p.ID)
					require.Equal(t, expectedProducts[i].Name, p.Name)
					require.Equal(t, expectedProducts[i].Price, p.Price)
					require.Equal(t, expectedProducts[i].Quantity, p.Quantity)
				}
			},
		},
		{
			name: "EmptyList",
			req:  ListProductRequest{Limit: limit, Offset: offset},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), db.ListProductsParams{Limit: limit, Offset: offset}).
					Times(1).
					Return([]db.Product{}, nil)
			},
			checkResponse: func(products []dto.ProductResponse, err error) {
				require.NoError(t, err)
				require.Empty(t, products)
			},
		},
		{
			name: "DBError",
			req:  ListProductRequest{Limit: limit, Offset: offset},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), db.ListProductsParams{Limit: limit, Offset: offset}).
					Times(1).
					Return([]db.Product{}, errors.New("db error"))
			},
			checkResponse: func(products []dto.ProductResponse, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "db error")
			},
		},
		{
			name: "WrongOffset",
			req:  ListProductRequest{Limit: limit, Offset: -offset},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), db.ListProductsParams{Limit: limit, Offset: -offset}).
					Times(1).
					Return([]db.Product{}, errors.New(""))
			},
			checkResponse: func(products []dto.ProductResponse, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "WrongLimit",
			req:  ListProductRequest{Limit: -limit, Offset: offset},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListProducts(gomock.Any(), db.ListProductsParams{Limit: -limit, Offset: offset}).
					Times(1).
					Return([]db.Product{}, errors.New("LIMIT must not be negative"))
			},
			checkResponse: func(products []dto.ProductResponse, err error) {
				require.Error(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			service, store := newTestProduct(t)

			tc.buildStubs(store)
			products, err := service.ListProducts(context.Background(), tc.req)

			tc.checkResponse(products, err)
		})
	}
}
