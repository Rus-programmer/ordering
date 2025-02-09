package order

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
)

type Order interface {
	GetOrder(ctx context.Context, req GetOrder) (dto.OrderResponse, error)
	ListOrders(ctx context.Context, req ListOrders) ([]dto.OrderResponse, error)
}

type order struct {
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewOrder(config util.Config, store db.Store, tokenMaker token.Maker) Order {
	return &order{
		tokenMaker: tokenMaker,
		store:      store,
		config:     config,
	}
}
