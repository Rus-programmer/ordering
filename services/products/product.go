package products

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
)

type Product interface {
	ListProducts(ctx context.Context, req ListProductRequest) ([]dto.ProductResponse, error)
	GetProduct(ctx context.Context, id int64) (dto.ProductResponse, error)
}

type product struct {
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewProduct(config util.Config, store db.Store, tokenMaker token.Maker) Product {
	return &product{
		tokenMaker: tokenMaker,
		store:      store,
		config:     config,
	}
}
