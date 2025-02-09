-- name: CreateOrder :one
INSERT INTO orders (customer_id, total_price)
VALUES ($1, $2)
RETURNING *;

-- name: ListOrders :many
SELECT id
FROM orders
WHERE (customer_id = $1 OR EXISTS (SELECT 1 FROM customers WHERE id = $1 AND role = 'admin'))
    AND (sqlc.narg(status)::order_status IS NULL OR status = sqlc.narg(status))
    AND (sqlc.narg(min_price)::bigint IS NULL OR total_price >= sqlc.narg(min_price))
    AND (sqlc.narg(max_price)::bigint IS NULL OR total_price <= sqlc.narg(max_price));

-- name: GetOrder :one
SELECT *
FROM orders o
WHERE o.id = $1
  and (customer_id = $2 OR EXISTS (SELECT 1 FROM customers WHERE id = $2 AND role = 'admin'));

-- name: UpdateOrderStatus :one
UPDATE orders o
SET status     = $3,
    updated_at = NOW()
WHERE o.id = $1
  AND (customer_id = $2 OR EXISTS (SELECT 1 FROM customers WHERE id = $2 AND role = 'admin'))
RETURNING *;

-- name: SoftDeleteOrder :one
UPDATE orders o
SET is_deleted = TRUE
WHERE o.id = $1
  AND (customer_id = $2 OR EXISTS (SELECT 1 FROM customers WHERE id = $2 AND role = 'admin'))
RETURNING *;

-- name: GetOrderProducts :many
SELECT *
FROM order_products
where order_id = $1;

-- name: DeleteOrder :exec
DELETE
FROM orders
WHERE id = $1;

-- name: GetTotalPrice :one
SELECT SUM(p.price * op.ordered_amount)
FROM order_products op
         JOIN products p ON op.product_id = p.id
WHERE op.order_id = $1;

-- name: CreateOrderProducts :copyfrom
INSERT INTO order_products (order_id, product_id, ordered_amount)
VALUES ($1, $2, $3);
