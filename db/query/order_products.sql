-- name: GetOrderProducts :many
SELECT *
FROM order_products
where order_id = $1;

-- name: GetTotalPrice :one
SELECT SUM(p.price * op.ordered_amount)
FROM order_products op
         JOIN products p ON op.product_id = p.id
WHERE op.order_id = $1;

-- name: CreateOrderProducts :copyfrom
INSERT INTO order_products (order_id, product_id, ordered_amount)
VALUES ($1, $2, $3);

-- name: UpdateOrderProduct :one
UPDATE order_products op
SET ordered_amount = $1
WHERE order_id = $2 AND product_id = $3
RETURNING *;

-- name: DeleteOrderProduct :exec
DELETE FROM order_products WHERE order_id = $1 AND product_id = $2;
