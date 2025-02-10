package order

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
)

type QueryParams struct {
	MinPrice int64
	MaxPrice int64
	Status   db.OrderStatus
}

type ListOrders struct {
	QueryParams QueryParams
	Payload     *token.Payload
}

func (o *order) ListOrders(ctx context.Context, req ListOrders) ([]dto.OrderResponse, error) {
	arg := db.ListOrdersParams{
		CustomerID: req.Payload.CustomerID,
		Status: db.NullOrderStatus{
			OrderStatus: req.QueryParams.Status,
			Valid:       req.QueryParams.Status != "",
		},
		MinPrice: pgtype.Int8{
			Int64: req.QueryParams.MinPrice,
			Valid: req.QueryParams.MinPrice > 0,
		},
		MaxPrice: pgtype.Int8{
			Int64: req.QueryParams.MaxPrice,
			Valid: req.QueryParams.MaxPrice > 0,
		},
	}

	orderIDs, err := o.store.ListOrders(ctx, arg)
	if err != nil {
		return []dto.OrderResponse{}, err
	}

	var orders []dto.OrderResponse
	for _, id := range orderIDs {
		arg := GetOrder{
			ID:      id,
			Payload: req.Payload,
		}
		gotOrder, err := o.GetOrder(ctx, arg)
		if err != nil {
			return []dto.OrderResponse{}, err
		}

		orders = append(orders, gotOrder)
	}

	return orders, nil
}
