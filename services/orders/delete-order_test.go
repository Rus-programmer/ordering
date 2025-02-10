package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/token"
	"ordering/util"
	"testing"
)

func TestDeleteOrder(t *testing.T) {
	customerID := util.RandomInt(1, 100)
	adminID := util.RandomInt(101, 200)
	orderID := util.RandomInt(1000, 2000)
	payloadUser := &token.Payload{CustomerID: customerID, Role: "user"}
	payloadAdmin := &token.Payload{CustomerID: adminID, Role: "admin"}

	order := db.Order{
		ID:         orderID,
		CustomerID: customerID,
	}

	testCases := []struct {
		name          string
		req           DeleteOrderParams
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(err error)
	}{
		{
			name: "OK",
			req: DeleteOrderParams{
				ID:      orderID,
				Payload: payloadUser,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetOrder(gomock.Any(), gomock.Eq(db.GetOrderParams{
					CustomerID: customerID,
					ID:         orderID,
				})).Times(1).Return(order, nil)

				store.EXPECT().SoftDeleteOrder(gomock.Any(), gomock.Eq(db.SoftDeleteOrderParams{
					CustomerID: customerID,
					ID:         orderID,
				})).Times(1).Return(order, nil)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "AdminCanDeleteAnyOrder",
			req: DeleteOrderParams{
				ID:      orderID,
				Payload: payloadAdmin,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Times(1).Return(order, nil)
				store.EXPECT().SoftDeleteOrder(gomock.Any(), gomock.Any()).Times(1).Return(order, nil)
			},
			checkResponse: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "UnauthorizedUserCannotDeleteOtherOrder",
			req: DeleteOrderParams{
				ID:      orderID,
				Payload: &token.Payload{CustomerID: customerID + 1, Role: "user"},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Times(1).Return(order, nil)
			},
			checkResponse: func(err error) {
				require.Error(t, err)
				require.Equal(t, util.ErrRecordNotFound, err)
			},
		},
		{
			name: "OrderNotFound",
			req: DeleteOrderParams{
				ID:      orderID,
				Payload: payloadUser,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Times(1).Return(db.Order{}, errors.New("no rows in result set"))
			},
			checkResponse: func(err error) {
				require.Error(t, err)
			},
		},
		{
			name: "DBError",
			req: DeleteOrderParams{
				ID:      orderID,
				Payload: payloadUser,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Times(1).Return(order, nil)
				store.EXPECT().SoftDeleteOrder(gomock.Any(), gomock.Any()).Times(1).Return(db.Order{}, errors.New("DB error"))
			},
			checkResponse: func(err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testOrder, store := newTestOrder(t)

			tc.buildStubs(store)
			err := testOrder.DeleteOrder(context.Background(), tc.req)
			tc.checkResponse(err)
		})
	}
}
