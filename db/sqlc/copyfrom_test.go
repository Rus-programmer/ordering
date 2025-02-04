package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateOrderProducts(t *testing.T) {
	product1 := createRandomProduct(t)
	product2 := createRandomProduct(t)
	product3 := createRandomProduct(t)
	order1 := createRandomOrder(t)
	order2 := createRandomOrder(t)

	orderProducts := []CreateOrderProductsParams{
		{OrderID: order1.ID, ProductID: product1.ID, Quantity: 2},
		{OrderID: order1.ID, ProductID: product2.ID, Quantity: 3},
		{OrderID: order2.ID, ProductID: product3.ID, Quantity: 4},
	}

	rowsAffected, err := testStore.CreateOrderProducts(context.Background(), orderProducts)
	assert.NoError(t, err, "expected no error while inserting order products")

	assert.Equal(t, int64(3), rowsAffected)

	deleteTestOrder(t, order1)
	deleteTestOrder(t, order2)
	deleteTestProduct(t, product1.ID)
	deleteTestProduct(t, product2.ID)
	deleteTestProduct(t, product3.ID)
}
