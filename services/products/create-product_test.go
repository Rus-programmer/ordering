package products

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/util"
	"testing"
	"time"
)

func TestCreateProduct(t *testing.T) {
	ID := util.RandomInt(1, 100)
	productName := "Test Product"
	productPrice := util.RandomInt(10, 1000)
	productQuantity := util.RandomInt(1, 100)

	createReq := CreateProduct{
		Name:     productName,
		Price:    productPrice,
		Quantity: productQuantity,
	}

	expectedProduct := db.Product{
		ID:        ID,
		Name:      productName,
		Price:     productPrice,
		Quantity:  productQuantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testCases := []struct {
		name          string
		body          CreateProduct
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(product dto.ProductResponse, err error)
	}{
		{
			name: "OK",
			body: createReq,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateProduct(gomock.Any(), db.CreateProductParams{
						Name:     productName,
						Price:    productPrice,
						Quantity: productQuantity,
					}).
					Times(1).
					Return(expectedProduct, nil)
			},
			checkResponse: func(product dto.ProductResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, expectedProduct.ID, product.ID)
				require.Equal(t, expectedProduct.Name, product.Name)
				require.Equal(t, expectedProduct.Price, product.Price)
				require.Equal(t, expectedProduct.Quantity, product.Quantity)
			},
		},
		{
			name: "DBError",
			body: createReq,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateProduct(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Product{}, errors.New("db error"))
			},
			checkResponse: func(product dto.ProductResponse, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "db error")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			service, store := newTestProduct(t)

			tc.buildStubs(store)
			product, err := service.CreateProduct(context.Background(), tc.body)

			tc.checkResponse(product, err)
		})
	}
}
