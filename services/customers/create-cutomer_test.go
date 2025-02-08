package customers

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	mockdb "ordering/db/mock"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/util"
	"testing"
	"time"
)

func TestCreateCustomer(t *testing.T) {
	username := util.RandomString(6)
	password := util.RandomString(6)
	testCases := []struct {
		name          string
		request       CreateCustomer
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(resp dto.CustomerResponse, err error)
	}{
		{
			name: "OK",
			request: CreateCustomer{
				Username: username,
				Password: password,
				Role:     db.UserRoleUser,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).Times(1).Return(db.Customer{
					ID:             1,
					Username:       username,
					HashedPassword: "hashedpassword",
					Role:           db.UserRoleUser,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
			checkResponse: func(resp dto.CustomerResponse, err error) {
				require.NoError(t, err)
				require.Equal(t, username, resp.Username)
			},
		},
		{
			name: "DBError",
			request: CreateCustomer{
				Username: username,
				Password: password,
				Role:     db.UserRoleUser,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateCustomer(gomock.Any(), gomock.Any()).Times(1).Return(db.Customer{}, errors.New("db error"))
			},
			checkResponse: func(resp dto.CustomerResponse, err error) {
				require.Error(t, err)
				require.Contains(t, err.Error(), "db error")
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			customer, store := newTestCustomer(t)
			tc.buildStubs(store)
			resp, err := customer.CreateCustomer(context.Background(), tc.request)
			tc.checkResponse(resp, err)
		})
	}
}
