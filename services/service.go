package services

import (
	db "ordering/db/sqlc"
	"ordering/services/auth"
	"ordering/services/customers"
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
	GetTokenMaker() token.Maker
}

type service struct {
	customers.Customer
	products.Product
	auth.Auth
	order.Order
	tokenMaker token.Maker
}

func NewService(config util.Config, store db.Store, tokenMaker token.Maker) Service {
	return &service{
		Customer:   customers.NewCustomer(config, store, tokenMaker),
		Product:    products.NewProduct(config, store, tokenMaker),
		Auth:       auth.NewAuth(config, store, tokenMaker),
		Order:      order.NewOrder(config, store, tokenMaker),
		tokenMaker: tokenMaker,
	}
}

func (s *service) GetTokenMaker() token.Maker {
	return s.tokenMaker
}
