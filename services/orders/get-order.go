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
	if cachedOrder, ok := o.cache.Get(req.ID); ok {
		if orderResponse, valid := cachedOrder.(dto.OrderResponse); valid {
			if err := validateOrderAccess(orderResponse.CustomerID, req.Payload); err != nil {
				return dto.OrderResponse{}, util.ErrRecordNotFound
			}
			return orderResponse, nil
		}
	}
	arg := db.GetOrderParams{
		CustomerID: req.Payload.CustomerID,
		ID:         req.ID,
	}
	order, err := o.store.GetOrder(ctx, arg)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	if err := validateOrderAccess(order.CustomerID, req.Payload); err != nil {
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

	if totalPrice != order.TotalPrice {
		return dto.OrderResponse{}, util.ErrMismatchedData
	}

	orderResponse := dto.OrderResponse{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		IsDeleted:  order.IsDeleted,
		TotalPrice: totalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		Products:   orderProductResponses,
	}

	o.cache.Add(req.ID, orderResponse)
	return orderResponse, nil
}

func validateOrderAccess(customerID int64, payload *token.Payload) error {
	if customerID != payload.CustomerID && payload.Role != string(db.UserRoleAdmin) {
		return util.ErrRecordNotFound
	}
	return nil
}
