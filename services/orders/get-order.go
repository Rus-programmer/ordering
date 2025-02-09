package order

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
)

type GetOrder struct {
	ID      int64
	Payload *token.Payload
}

func (o *order) GetOrder(ctx context.Context, req GetOrder) (dto.OrderResponse, error) {
	arg := db.GetOrderParams{
		CustomerID: req.Payload.CustomerID,
		ID:         req.ID,
	}
	order, err := o.store.GetOrder(ctx, arg)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	if order.CustomerID != req.Payload.CustomerID && req.Payload.Role != string(db.UserRoleAdmin) {
		return dto.OrderResponse{}, util.ErrRecordNotFound
	}

	orderProducts, err := o.store.GetOrderProducts(ctx, order.ID)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	var orderProductResponses []dto.OrderProductResponse

	for _, op := range orderProducts {
		product, err := o.store.GetProduct(ctx, op.ProductID)
		if err != nil {
			return dto.OrderResponse{}, err
		}

		orderProductResponses = append(orderProductResponses, dto.OrderProductResponse{
			Product: dto.ProductResponse{
				ID:        product.ID,
				Name:      product.Name,
				Price:     product.Price,
				Quantity:  product.Quantity,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
			},
			OrderedAmount: op.OrderedAmount,
		})
	}

	totalPrice, err := o.store.GetTotalPrice(ctx, order.ID)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	return dto.OrderResponse{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		IsDeleted:  order.IsDeleted,
		TotalPrice: totalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		Products:   orderProductResponses,
	}, nil
}
