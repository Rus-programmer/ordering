package products

import (
	"context"
	"github.com/jackc/pgx/v5"
	"ordering/dto"
)

func (product *product) GetProduct(ctx context.Context, id int64) (dto.ProductResponse, error) {
	products, err := product.store.GetProduct(ctx, id)
	if err != nil {
		return dto.ProductResponse{}, pgx.ErrNoRows
	}

	return dto.ProductResponse{
		ID:        products.ID,
		Name:      products.Name,
		Price:     products.Price,
		Quantity:  products.Quantity,
		CreatedAt: products.CreatedAt,
		UpdatedAt: products.UpdatedAt,
	}, nil
}
