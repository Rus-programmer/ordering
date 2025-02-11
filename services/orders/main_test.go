package order

import (
	lru "github.com/hashicorp/golang-lru/v2"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	"ordering/util"
	"testing"
	"time"
)

func newTestOrder(t *testing.T) (Order, *mockdb.MockStore) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	cache, _ := lru.New[int64, any](50)

	t.Cleanup(func() {
		cache.Purge()
	})

	return &order{
		store:  store,
		config: config,
		cache:  cache,
	}, store
}
