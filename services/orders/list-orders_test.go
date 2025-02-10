package order

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
)

func TestListOrders(t *testing.T) {
	customerID := util.RandomInt(1, 100)
	payload := &token.Payload{
		CustomerID: customerID,
		Role:       "admin",
	}
	order1 := randomOrder()

	orderProducts := []db.OrderProduct{
		{
			OrderID:       order1.ID,
			ProductID:     util.RandomInt(1, 1000),
			OrderedAmount: 2,
		},
	}
	product1 := randomProduct(orderProducts[0].ProductID)
	testCases := []struct {
		name          string
		req           ListOrders
		buildStubs    func(store *mockdb.MockStore, req ListOrders)
		checkResponse func(orders []dto.OrderResponse, err error)
	}{
		{
			name: "OK",
			req: ListOrders{
				QueryParams: QueryParams{
					MinPrice: 1,
					MaxPrice: 1500,
					Status:   "pending",
				},
				Payload: payload,
			},
			buildStubs: func(store *mockdb.MockStore, req ListOrders) {
				store.EXPECT().ListOrders(gomock.Any(), gomock.Any()).Times(1).Return([]int64{order1.ID}, nil)
				store.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Times(1).Return(order1, nil)
				store.EXPECT().GetOrderProducts(gomock.Any(), order1.ID).Times(1).Return(orderProducts, nil)
				store.EXPECT().GetProduct(gomock.Any(), product1.ID).Times(1).Return(product1, nil)
				store.EXPECT().GetTotalPrice(gomock.Any(), order1.ID).Times(1).Return(order1.TotalPrice, nil)
			},
			checkResponse: func(orders []dto.OrderResponse, err error) {
				require.NoError(t, err)
				require.Len(t, orders, 1)
			},
		},
		{
			name: "NoOrdersFound",
			req: ListOrders{
				QueryParams: QueryParams{
					MinPrice: 5000,
					MaxPrice: 10000,
					Status:   "cancelled",
				},
				Payload: &token.Payload{
					CustomerID: customerID,
					Role:       "user",
				},
			},
			buildStubs: func(store *mockdb.MockStore, req ListOrders) {
				store.EXPECT().ListOrders(gomock.Any(), gomock.Any()).Times(1).Return([]int64{}, nil)
			},
			checkResponse: func(orders []dto.OrderResponse, err error) {
				require.NoError(t, err)
				require.Len(t, orders, 0)
			},
		},
		{
			name: "DBError",
			req: ListOrders{
				QueryParams: QueryParams{
					MinPrice: 50,
					MaxPrice: 1500,
					Status:   "pending",
				},
				Payload: &token.Payload{
					CustomerID: customerID,
					Role:       "user",
				},
			},
			buildStubs: func(store *mockdb.MockStore, req ListOrders) {
				store.EXPECT().ListOrders(gomock.Any(), gomock.Any()).Times(1).Return(nil, errors.New("DB error"))
			},
			checkResponse: func(orders []dto.OrderResponse, err error) {
				require.Error(t, err)
				require.Equal(t, orders, []dto.OrderResponse{})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testOrder, store := newTestOrder(t)

			tc.buildStubs(store, tc.req)
			orders, err := testOrder.ListOrders(context.Background(), tc.req)
			tc.checkResponse(orders, err)
		})
	}
}
