package db

import (
	"context"
	"github.com/stretchr/testify/require"
	util "ordering/utils"
	"testing"
)

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Name:     util.RandomString(6),
		Price:    util.RandomInt(6, 1000),
		Quantity: util.RandomInt(6, 50),
	}

	product, err := testStore.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Price, product.Price)
	require.NotZero(t, product.CreatedAt)
	require.NotZero(t, product.UpdatedAt)

	return product
}

func deleteTestProduct(t *testing.T, productID int64) {
	err := testStore.DeleteProduct(context.Background(), productID)
	require.NoError(t, err)
}
