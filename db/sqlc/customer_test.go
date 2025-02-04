package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	util "ordering/utils"
	"testing"
	"time"
)

func TestCreateCustomer(t *testing.T) {
	customer := createRandomCustomer(t)

	deleteTestCustomer(t, customer.ID)
}

func createRandomCustomer(t *testing.T) Customer {
	arg := CreateCustomerParams{
		Username:       util.RandomOwner(),
		HashedPassword: util.RandomString(6),
		Role:           UserRoleUser,
	}

	customer, err := testStore.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.Username, customer.Username)
	require.Equal(t, arg.HashedPassword, customer.HashedPassword)
	require.NotZero(t, customer.CreatedAt)
	require.NotZero(t, customer.UpdatedAt)
	require.Equal(t, customer.Role, UserRoleUser)

	return customer
}

func deleteTestCustomer(t *testing.T, customerID int64) {
	err := testStore.DeleteCustomer(context.Background(), customerID)
	require.NoError(t, err)
}

func TestGetCustomerByUsername(t *testing.T) {
	customer := createRandomCustomer(t)

	created, err := testStore.GetCustomerByUsername(context.Background(), customer.Username)
	require.NoError(t, err)

	require.Equal(t, customer.ID, created.ID)
	require.Equal(t, customer.Role, created.Role)
	require.Equal(t, customer.HashedPassword, created.HashedPassword)
	require.WithinDuration(t, customer.CreatedAt, created.CreatedAt, time.Second)
	require.WithinDuration(t, customer.UpdatedAt, created.UpdatedAt, time.Second)

	deleteTestCustomer(t, customer.ID)
}

func TestListCustomers(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2 := createRandomCustomer(t)
	customer3 := createRandomCustomer(t)

	params := ListCustomersParams{
		Limit:  10,
		Offset: 0,
	}
	customers, err := testStore.ListCustomers(context.Background(), params)
	assert.NoError(t, err)

	assert.Equal(t, len(customers), 3)

	deleteTestCustomer(t, customer1.ID)
	deleteTestCustomer(t, customer2.ID)
	deleteTestCustomer(t, customer3.ID)
}
