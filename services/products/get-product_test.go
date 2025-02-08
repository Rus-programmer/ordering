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

func TestGetProduct(t *testing.T) {
	ID := util.RandomInt(1, 100)

	expectedProduct := db.Product{
		ID:        ID,
		Name:      "Test Product",
		Price:     util.RandomInt(10, 1000),
		Quantity:  util.RandomInt(1, 100),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testCases := []struct {
		name          string
		productID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(product dto.ProductResponse, err error)
	}{
		{
			name:      "OK",
			productID: ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), ID).
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
			name:      "NotFound",
			productID: ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), ID).
					Times(1).
					Return(db.Product{}, util.ErrRecordNotFound)
			},
			checkResponse: func(product dto.ProductResponse, err error) {
				require.Error(t, err)
				require.True(t, errors.Is(err, util.ErrRecordNotFound))
			},
		},
		{
			name:      "DBError",
			productID: ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProduct(gomock.Any(), ID).
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
			product, err := service.GetProduct(context.Background(), tc.productID)

			tc.checkResponse(product, err)
		})
	}
}
