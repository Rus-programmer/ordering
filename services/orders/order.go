package order

import (
	"context"
	lru "github.com/hashicorp/golang-lru/v2"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/util"
)

type Order interface {
	GetOrder(ctx context.Context, req GetOrder) (dto.OrderResponse, error)
	ListOrders(ctx context.Context, req ListOrders) ([]dto.OrderResponse, error)
	DeleteOrder(ctx context.Context, req DeleteOrderParams) error
	CreateOrder(ctx context.Context, req CreateOrderParams) (dto.OrderResponse, error)
	UpdateOrder(ctx context.Context, req UpdateOrderParams) (dto.OrderResponse, error)
}

type order struct {
	store  db.Store
	config util.Config
	cache  *lru.Cache[int64, any]
}

func NewOrder(config util.Config, store db.Store, cache *lru.Cache[int64, any]) Order {
	return &order{
		store:  store,
		config: config,
		cache:  cache,
	}
}
