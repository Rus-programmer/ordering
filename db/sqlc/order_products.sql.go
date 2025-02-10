// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order_products.sql

package db

import (
	"context"
)

type CreateOrderProductsParams struct {
	OrderID       int64 `json:"order_id"`
	ProductID     int64 `json:"product_id"`
	OrderedAmount int64 `json:"ordered_amount"`
}

const deleteOrderProduct = `-- name: DeleteOrderProduct :exec
DELETE FROM order_products WHERE order_id = $1 AND product_id = $2
`

type DeleteOrderProductParams struct {
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
}

func (q *Queries) DeleteOrderProduct(ctx context.Context, arg DeleteOrderProductParams) error {
	_, err := q.db.Exec(ctx, deleteOrderProduct, arg.OrderID, arg.ProductID)
	return err
}

const getOrderProducts = `-- name: GetOrderProducts :many
SELECT order_id, product_id, ordered_amount
FROM order_products
where order_id = $1
ORDER BY product_id
`

func (q *Queries) GetOrderProducts(ctx context.Context, orderID int64) ([]OrderProduct, error) {
	rows, err := q.db.Query(ctx, getOrderProducts, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderProduct{}
	for rows.Next() {
		var i OrderProduct
		if err := rows.Scan(&i.OrderID, &i.ProductID, &i.OrderedAmount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTotalPrice = `-- name: GetTotalPrice :one
SELECT SUM(p.price * op.ordered_amount)
FROM order_products op
         JOIN products p ON op.product_id = p.id
WHERE op.order_id = $1
`

func (q *Queries) GetTotalPrice(ctx context.Context, orderID int64) (int64, error) {
	row := q.db.QueryRow(ctx, getTotalPrice, orderID)
	var sum int64
	err := row.Scan(&sum)
	return sum, err
}

const updateOrderProduct = `-- name: UpdateOrderProduct :one
UPDATE order_products op
SET ordered_amount = $1
WHERE order_id = $2 AND product_id = $3
RETURNING order_id, product_id, ordered_amount
`

type UpdateOrderProductParams struct {
	OrderedAmount int64 `json:"ordered_amount"`
	OrderID       int64 `json:"order_id"`
	ProductID     int64 `json:"product_id"`
}

func (q *Queries) UpdateOrderProduct(ctx context.Context, arg UpdateOrderProductParams) (OrderProduct, error) {
	row := q.db.QueryRow(ctx, updateOrderProduct, arg.OrderedAmount, arg.OrderID, arg.ProductID)
	var i OrderProduct
	err := row.Scan(&i.OrderID, &i.ProductID, &i.OrderedAmount)
	return i, err
}
