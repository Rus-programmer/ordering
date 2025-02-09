package products

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
)

type ProductInput struct {
	Name     string
	Price    int64
	Quantity int64
}

type CreateProduct ProductInput

func (product *product) CreateProduct(ctx context.Context, body CreateProduct) (dto.ProductResponse, error) {
	newProduct, err := product.store.CreateProduct(ctx, db.CreateProductParams{
		Name:     body.Name,
		Price:    body.Price,
		Quantity: body.Quantity,
	})
	if err != nil {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:        newProduct.ID,
		Name:      newProduct.Name,
		Price:     newProduct.Price,
		Quantity:  newProduct.Quantity,
		CreatedAt: newProduct.CreatedAt,
		UpdatedAt: newProduct.UpdatedAt,
	}, nil
}
