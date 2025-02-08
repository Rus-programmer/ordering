package auth

import (
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	"ordering/token"
	"ordering/util"
	"testing"
	"time"
)

func newTestAuth(t *testing.T) (Auth, *mockdb.MockStore, *token.MockMaker) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	tokenMaker := token.NewMockMaker(ctrl)
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	return &auth{
		tokenMaker: tokenMaker,
		store:      store,
		config:     config,
	}, store, tokenMaker
}
