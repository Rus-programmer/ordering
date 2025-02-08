package customers

import (
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	"ordering/util"
	"testing"
	"time"
)

func newTestCustomer(t *testing.T) (Customer, *mockdb.MockStore) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	return &customer{
		store:  store,
		config: config,
	}, store
}
