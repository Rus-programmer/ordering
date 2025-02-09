package products

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/util"
)

type UpdateProduct ProductInput

func (product *product) UpdateProduct(ctx context.Context, id int64, body UpdateProduct) (dto.ProductResponse, error) {
	_, err := product.store.GetProduct(ctx, id)
	if err != nil {
		return dto.ProductResponse{}, fmt.Errorf("%w: %v", util.ErrRecordNotFound, err)
	}

	newProduct, err := product.store.UpdateProduct(ctx, db.UpdateProductParams{
		Name: pgtype.Text{
			String: body.Name,
			Valid:  body.Name != "",
		},
		Price: pgtype.Int8{
			Int64: body.Price,
			Valid: body.Price != 0,
		},
		Quantity: pgtype.Int8{
			Int64: body.Quantity,
			Valid: body.Quantity != 0,
		},
		ID: id,
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
