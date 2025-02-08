package products

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
)

type ListProductRequest struct {
	Limit  int32
	Offset int32
}

func (product *product) ListProducts(ctx context.Context, req ListProductRequest) ([]dto.ProductResponse, error) {
	products, err := product.store.ListProducts(ctx, db.ListProductsParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return []dto.ProductResponse{}, err
	}

	var productResponses []dto.ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			ID:        p.ID,
			Name:      p.Name,
			Price:     p.Price,
			Quantity:  p.Quantity,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return productResponses, nil
}
