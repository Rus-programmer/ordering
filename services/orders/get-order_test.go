package order

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
	"testing"
	"time"
)

func TestGetOrder(t *testing.T) {
	customerID2 := util.RandomInt(1, 100)
	adminID := util.RandomInt(101, 200)

	expectedOrder := randomOrder()

	orderProducts := []db.OrderProduct{
		{
			OrderID:       expectedOrder.ID,
			ProductID:     util.RandomInt(1, 1000),
			OrderedAmount: 2,
		},
	}
	product1 := randomProduct(orderProducts[0].ProductID)
	product2 := randomProduct(orderProducts[0].ProductID)
	products := []db.Product{
		product1, product2,
	}

	totalPrice := int64(1000)

	testCases := []struct {
		name          string
		req           GetOrder
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(order dto.OrderResponse, err error)
	}{
		{
			name: "OK",
			req: GetOrder{
				ID: expectedOrder.ID,
				Payload: &token.Payload{
					CustomerID: expectedOrder.CustomerID,
					Role:       "customer",
				},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Times(1).Return(expectedOrder, nil)
				store.EXPECT().GetOrderProducts(gomock.Any(), expectedOrder.ID).Times(1).Return(orderProducts, nil)
				store.EXPECT().GetProduct(gomock.Any(), products[0].ID).Times(1).Return(products[0], nil)
				store.EXPECT().GetTotalPrice(gomock.Any(), expectedOrder.ID).Times(1).Return(totalPrice, nil)
			},
			checkResponse: func(order dto.OrderResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, expectedOrder.ID, order.ID)
				require.Equal(t, expectedOrder.CustomerID, order.CustomerID)
				require.Equal(t, totalPrice, order.TotalPrice)
			},
		},
		{
			name: "Admin",
			req: GetOrder{
				ID: expectedOrder.ID,
				Payload: &token.Payload{
					CustomerID: adminID,
					Role:       "admin",
				},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Times(1).Return(expectedOrder, nil)
				store.EXPECT().GetOrderProducts(gomock.Any(), expectedOrder.ID).Times(1).Return(orderProducts, nil)
				store.EXPECT().GetProduct(gomock.Any(), products[0].ID).Times(1).Return(products[0], nil)
				store.EXPECT().GetTotalPrice(gomock.Any(), expectedOrder.ID).Times(1).Return(totalPrice, nil)
			},
			checkResponse: func(order dto.OrderResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, expectedOrder.ID, order.ID)
				require.Equal(t, expectedOrder.CustomerID, order.CustomerID)
				require.Equal(t, totalPrice, order.TotalPrice)
			},
		},
		{
			name: "AnotherCustomer",
			req: GetOrder{
				ID: expectedOrder.ID,
				Payload: &token.Payload{
					CustomerID: customerID2,
					Role:       "user",
				},
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetOrder(gomock.Any(), gomock.Any()).Times(1).Return(db.Order{}, util.ErrRecordNotFound)
			},
			checkResponse: func(order dto.OrderResponse, err error) {
				require.Error(t, err)
				require.True(t, errors.Is(err, util.ErrRecordNotFound))
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			service, store := newTestOrder(t)

			tc.buildStubs(store)
			order, err := service.GetOrder(context.Background(), tc.req)

			tc.checkResponse(order, err)
		})
	}
}

func randomProduct(ID int64) db.Product {
	return db.Product{
		ID:        ID,
		Name:      util.RandomString(10),
		Price:     util.RandomInt(10, 1000),
		Quantity:  util.RandomInt(1, 100),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func randomOrder() db.Order {
	return db.Order{
		ID:         util.RandomInt(1, 100),
		CustomerID: util.RandomInt(1, 100),
		IsDeleted:  false,
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
