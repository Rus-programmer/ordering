package products

import (
	"context"
	"ordering/dto"
)

func (product *product) GetProduct(ctx context.Context, id int64) (dto.ProductResponse, error) {
	products, err := product.store.GetProduct(ctx, id)
	if err != nil {
		return dto.ProductResponse{}, err
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
