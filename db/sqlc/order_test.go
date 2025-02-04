package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomOrder(t *testing.T) Order {
	customer := createRandomCustomer(t)

	order, err := testStore.CreateOrder(context.Background(), customer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, customer.ID, order.CustomerID)
	require.Equal(t, OrderStatusPending, order.Status)
	require.NotZero(t, order.CreatedAt)
	require.NotZero(t, order.UpdatedAt)

	return order
}

func deleteTestOrder(t *testing.T, order Order) {
	err := testStore.DeleteOrder(context.Background(), order.ID)
	require.NoError(t, err)

	err = testStore.DeleteCustomer(context.Background(), order.CustomerID)
	require.NoError(t, err)
}
