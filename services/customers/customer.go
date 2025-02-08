package customers

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
)

type Customer interface {
	CreateCustomer(ctx context.Context, req CreateCustomer) (dto.CustomerResponse, error)
}

type customer struct {
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewCustomer(config util.Config, store db.Store, tokenMaker token.Maker) Customer {
	return &customer{
		tokenMaker: tokenMaker,
		store:      store,
		config:     config,
	}
}
