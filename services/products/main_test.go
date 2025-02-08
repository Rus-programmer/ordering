package products

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	"ordering/token"
	"ordering/util"
	"testing"
	"time"
)

func newTestProduct(t *testing.T) (Product, *mockdb.MockStore) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	mockTokenMaker, err := token.NewPasetoMaker("12345678912345678912345678900032")
	require.NoError(t, err)
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	return &product{
		tokenMaker: mockTokenMaker,
		store:      store,
		config:     config,
	}, store
}
