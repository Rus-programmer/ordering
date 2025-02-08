package products

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/util"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	ID := util.RandomInt(1, 100)
	testCases := []struct {
		name          string
		productID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(err error)
	}{
		{
			name:      "OK",
			productID: ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), ID).Times(1).Return(db.Product{}, nil)
				store.EXPECT().DeleteProduct(gomock.Any(), ID).Times(1).Return(nil)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name:      "NotFound",
			productID: ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), ID).Times(1).Return(db.Product{}, util.ErrRecordNotFound)
			},
			checkResponse: func(err error) {
				require.Contains(t, err.Error(), "no rows")
			},
		},
		{
			name:      "DeleteError",
			productID: ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProduct(gomock.Any(), ID).Times(1).Return(db.Product{}, nil)
				store.EXPECT().DeleteProduct(gomock.Any(), ID).Times(1).Return(errors.New("db error"))
			},
			checkResponse: func(err error) {
				require.Error(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			service, store := newTestProduct(t)

			tc.buildStubs(store)
			err := service.DeleteProduct(context.Background(), tc.productID)

			tc.checkResponse(err)
		})
	}
}
