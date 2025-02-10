package order

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	db "ordering/db/sqlc"
	"ordering/dto"
	"ordering/token"
	"sort"
)

type UpdateOrderParams struct {
	OrderID  int64
	Status   db.OrderStatus
	Payload  *token.Payload
	Products []UpdateOrderItem
}

type UpdateOrderItem CreateOrderItem

func (o *order) UpdateOrder(ctx context.Context, req UpdateOrderParams) (dto.OrderResponse, error) {
	var orderProductResponses []dto.OrderProductResponse
	var newOrder db.Order

	// execute transaction first
	err := o.store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		existingOrder, err := o.store.GetOrder(ctx, db.GetOrderParams{
			ID:         req.OrderID,
			CustomerID: req.Payload.CustomerID,
		})
		if err != nil {
			return err
		}
		if existingOrder.Status == db.OrderStatusCancelled || existingOrder.Status == db.OrderStatusConfirmed {
			return fmt.Errorf("you can't update an order that has already been cancelled or confirmed")
		}
		if existingOrder.IsDeleted {
			return fmt.Errorf("you can't update deleted order")
		}

		// getting already existed products in order_products table
		existingOrderProducts, err := q.GetOrderProducts(ctx, req.OrderID)
		if err != nil {
			return err
		}

		existingMap := make(map[int64]int64) // product_id -> ordered_amount
		for _, p := range existingOrderProducts {
			existingMap[p.ProductID] = p.OrderedAmount
		}

		var toInsert []db.CreateOrderProductsParams
		var toUpdate []db.UpdateOrderProductParams
		var toDelete []int64

		newMap := make(map[int64]int64) // product_id -> ordered_amount
		for _, p := range req.Products {
			newMap[p.ProductID] = p.OrderedAmount
		}

		// identify inconsistency between existingMap and newMap
		for productID, newAmount := range newMap {
			if oldAmount, exists := existingMap[productID]; exists {
				if oldAmount != newAmount {
					toUpdate = append(toUpdate, db.UpdateOrderProductParams{
						OrderID:       req.OrderID,
						ProductID:     productID,
						OrderedAmount: newAmount,
					})
				}
				delete(existingMap, productID)
			} else {
				toInsert = append(toInsert, db.CreateOrderProductsParams{
					OrderID:       req.OrderID,
					ProductID:     productID,
					OrderedAmount: newAmount,
				})
			}
		}

		// left productIDs in existingMap should be deleted from db
		for productID := range existingMap {
			toDelete = append(toDelete, productID)
		}

		if len(toDelete) > 0 {
			for _, productID := range toDelete {
				err = q.DeleteOrderProduct(ctx, db.DeleteOrderProductParams{
					OrderID:   req.OrderID,
					ProductID: productID,
				})
				if err != nil {
					return err
				}
			}
		}
		if len(toInsert) > 0 {
			_, err = q.CreateOrderProducts(ctx, toInsert)
			if err != nil {
				return err
			}
		}
		for _, upd := range toUpdate {
			_, err := q.UpdateOrderProduct(ctx, upd)
			if err != nil {
				return err
			}
		}

		// sorting for prevent deadlocks
		req.SortProductsByID()

		// getting already existed products in order_products table
		updatedOrderProducts, err := q.GetOrderProducts(ctx, req.OrderID)
		if err != nil {
			return err
		}

		totalPrice := int64(0)
		for _, orderProduct := range updatedOrderProducts {
			var product db.Product
			product, err = q.GetProductForUpdate(ctx, orderProduct.ProductID)
			if err != nil {
				return err
			}
			//updating product quantity only for confirmed orders
			if req.Status == db.OrderStatusConfirmed {
				if product.Quantity-orderProduct.OrderedAmount < 0 {
					return fmt.Errorf("not enough stock available")
				}
				product, err = q.UpdateProduct(ctx, db.UpdateProductParams{
					Quantity: pgtype.Int8{
						Int64: product.Quantity - orderProduct.OrderedAmount,
						Valid: true,
					},
					ID: orderProduct.ProductID,
				})
				if err != nil {
					return err
				}
			}

			totalPrice += product.Price * orderProduct.OrderedAmount
			orderProductResponses = append(orderProductResponses, dto.OrderProductResponse{
				Product: dto.ProductResponse{
					ID:        product.ID,
					Name:      product.Name,
					Price:     product.Price,
					Quantity:  product.Quantity,
					CreatedAt: product.CreatedAt,
					UpdatedAt: product.UpdatedAt,
				},
				OrderedAmount: orderProduct.OrderedAmount,
			})
		}

		newOrder, err = q.UpdateOrder(ctx, db.UpdateOrderParams{
			ID:         req.OrderID,
			CustomerID: req.Payload.CustomerID,
			Status: db.NullOrderStatus{
				OrderStatus: req.Status,
				Valid:       req.Status != "",
			},
			TotalPrice: pgtype.Int8{
				Int64: totalPrice,
				Valid: totalPrice != 0,
			},
		})
		if err != nil {
			return err
		}

		// update order status if received
		if req.Status != "" {
			log.Info().Msgf(
				"status for order id %d changed. old status %s changed to %s",
				req.OrderID, existingOrder.Status, newOrder.Status,
			)
		}

		return nil
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

func (o *UpdateOrderParams) SortProductsByID() {
	sort.Slice(o.Products, func(i, j int) bool {
		return o.Products[i].ProductID < o.Products[j].ProductID
	})
}
