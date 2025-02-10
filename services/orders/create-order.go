package order

import (
	"context"
	"fmt"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"ordering/util"
)

type CreateOrderParams struct {
	Payload  *token.Payload
	Products []CreateOrderItem
}

type CreateOrderItem struct {
	ProductID     int64
	OrderedAmount int64
}

func (o *order) CreateOrder(ctx context.Context, req CreateOrderParams) (dto.OrderResponse, error) {
	var orderProductResponses []dto.OrderProductResponse
	var newOrder db.Order

	// execute transaction first
	err := o.store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		totalPrice := int64(0)

		var createOrderProductsArg []db.CreateOrderProductsParams
		// loop over products request to get all the products, validate existence, calculate the total price,
		// and construct the order product response
		for _, p := range req.Products {
			product, err := o.store.GetProduct(ctx, p.ProductID)
			if err != nil {
				return util.ErrRecordNotFound
			}

			totalPrice += product.Price * p.OrderedAmount

			createOrderProductsArg = append(createOrderProductsArg, db.CreateOrderProductsParams{
				OrderID:       0,
				ProductID:     p.ProductID,
				OrderedAmount: p.OrderedAmount,
			})

			orderProductResponses = append(orderProductResponses, dto.OrderProductResponse{
				Product: dto.ProductResponse{
					ID:        product.ID,
					Name:      product.Name,
					Price:     product.Price,
					Quantity:  product.Quantity,
					CreatedAt: product.CreatedAt,
					UpdatedAt: product.UpdatedAt,
				},
				OrderedAmount: p.OrderedAmount,
			})
		}
		// insert data to orders table
		newOrder, err = q.CreateOrder(ctx, db.CreateOrderParams{
			CustomerID: req.Payload.CustomerID,
			TotalPrice: totalPrice,
		})
		if err != nil {
			return err
		}

		// update OrderIDs
		for i := range createOrderProductsArg {
			createOrderProductsArg[i].OrderID = newOrder.ID
		}

		if len(createOrderProductsArg) == 0 {
			return fmt.Errorf("%w: %v", util.ErrRecordNotFound, "order must contain at least one product")
		}

		// insert all records to order_products table
		_, err = q.CreateOrderProducts(ctx, createOrderProductsArg)

		return err
	})
	// after the transaction ends, check for errors
	if err != nil {
		return dto.OrderResponse{}, err
	}

	return dto.OrderResponse{
		ID:         newOrder.ID,
		CustomerID: newOrder.CustomerID,
		IsDeleted:  newOrder.IsDeleted,
		TotalPrice: newOrder.TotalPrice,
		Status:     newOrder.Status,
		CreatedAt:  newOrder.CreatedAt,
		UpdatedAt:  newOrder.UpdatedAt,
		Products:   orderProductResponses,
	}, nil
}
