package services

import (
	lru "github.com/hashicorp/golang-lru/v2"
	db "ordering/db/sqlc"
	"ordering/services/auth"
	"ordering/services/customers"
	"ordering/services/metrics"
	order "ordering/services/orders"
	"ordering/services/products"
	"ordering/token"
	"ordering/util"
)

type Service interface {
	auth.Auth
	customers.Customer
	products.Product
	order.Order
	metrics.Metric
	GetTokenMaker() token.Maker
}

type service struct {
	customers.Customer
	products.Product
	auth.Auth
	order.Order
	metrics.Metric
	tokenMaker token.Maker
}

func NewService(config util.Config, store db.Store, tokenMaker token.Maker) Service {
	cache, _ := lru.New[int64, any](10000)
	return &service{
		Customer:   customers.NewCustomer(config, store, tokenMaker),
		Product:    products.NewProduct(config, store, tokenMaker),
		Auth:       auth.NewAuth(config, store, tokenMaker),
		Order:      order.NewOrder(config, store, cache),
		Metric:     metrics.NewMetric(store),
		tokenMaker: tokenMaker,
	}
}

func (s *service) GetTokenMaker() token.Maker {
	return s.tokenMaker
}
