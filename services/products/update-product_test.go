package products

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/util"
	"testing"
	"time"
)

func TestUpdateProduct(t *testing.T) {
	ID := util.RandomInt(1, 100)
	updatedName := "Updated Product"
	updatedPrice := util.RandomInt(10, 1000)
	updatedQuantity := util.RandomInt(1, 100)

	updateReq := UpdateProduct{
		Name:     updatedName,
		Price:    updatedPrice,
		Quantity: updatedQuantity,
	}

	expectedProduct := db.Product{
		ID:        ID,
		Name:      updatedName,
		Price:     updatedPrice,
		Quantity:  updatedQuantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testCases := []struct {
		name          string
		productID     int64
		body          UpdateProduct
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(product dto.ProductResponse, err error)
	}{
		{
			name:      "OK",
			productID: ID,
			body:      updateReq,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), ID).Times(1).Return(db.Product{}, nil)
				store.EXPECT().
					UpdateProduct(gomock.Any(), db.UpdateProductParams{
						Name: pgtype.Text{
							String: updatedName,
							Valid:  true,
						},
						Price: pgtype.Int8{
							Int64: updatedPrice,
							Valid: true,
						},
						Quantity: pgtype.Int8{
							Int64: updatedQuantity,
							Valid: true,
						},
						ID: ID,
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
			name:      "NotFound",
			productID: ID,
			body:      updateReq,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), ID).Times(1).Return(db.Product{}, util.ErrRecordNotFound)
			},
			checkResponse: func(product dto.ProductResponse, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "no rows")
			},
		},
		{
			name:      "UpdateError",
			productID: ID,
			body:      updateReq,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), ID).Times(1).Return(db.Product{}, nil)
				store.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Any()).
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
			product, err := service.UpdateProduct(context.Background(), tc.productID, tc.body)

			tc.checkResponse(product, err)
		})
	}
}
